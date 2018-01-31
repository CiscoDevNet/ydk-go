// Copyright (c) 2012 IETF Trust and the persons identified
// as the document authors.  All rights reserved.
// This MIB module contains generic object definitions for
// MPLS Traffic Engineering in transport networks.This module is a
// cisco-ized version of the IETF draft:
// draft-ietf-mpls-tp-te-mib-03
package cisco_ietf_mpls_te_ext_std_03_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_mpls_te_ext_std_03_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-MPLS-TE-EXT-STD-03-MIB CISCO-IETF-MPLS-TE-EXT-STD-03-MIB}", reflect.TypeOf(CISCOIETFMPLSTEEXTSTD03MIB{}))
    ydk.RegisterEntity("CISCO-IETF-MPLS-TE-EXT-STD-03-MIB:CISCO-IETF-MPLS-TE-EXT-STD-03-MIB", reflect.TypeOf(CISCOIETFMPLSTEEXTSTD03MIB{}))
}

// CISCOIETFMPLSTEEXTSTD03MIB
type CISCOIETFMPLSTEEXTSTD03MIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table allows the administrator to map a node or LSR Identifier (IP
    // compatible [Global_Node_ID] or ICC) with a local identifier.   This table
    // is created to reuse the existing mplsTunnelTable for MPLS based transport
    // network tunnels also. Since the MPLS tunnel's Ingress/Egress LSR
    // identifiers' size (Unsigned32) value is not compatible for MPLS-TP tunnel
    // i.e. Global_Node_Id of size 8 bytes and ICC of size 6 bytes, there exists a
    // need to map the Global_Node_ID or ICC with the local identifier of size 4
    // bytes (Unsigned32) value in order to index (Ingress/Egress LSR identifier)
    // the existing mplsTunnelTable.
    Cmplsnodeconfigtable CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable

    // This read-only table allows the administrator to retrieve the local
    // identifier for a given Global_Node_ID in an IP compatible operator
    // environment.  This table MAY be used in on-demand and/or proactive OAM
    // operations to get the Ingress/Egress LSR identifier (Local Identifier) from
    // Src-Global_Node_ID or Dst-Global_Node_ID and the Ingress and Egress LSR
    // identifiers are used to retrieve the tunnel entry.  This table returns
    // nothing when the associated entry is not defined in mplsNodeConfigTable.
    Cmplsnodeipmaptable CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable

    // This read-only table allows the administrator to retrieve the local
    // identifier for a given ICC operator in an ICC operator environment.  This
    // table MAY be used in on-demand and/or proactive OAM operations to get the
    // Ingress/Egress LSR identifier (Local Identifier) from Src-ICC or Dst-ICC
    // and the Ingress and Egress LSR identifiers are used to retrieve the tunnel
    // entry. This table returns nothing when the associated entry is not defined
    // in mplsNodeConfigTable.
    Cmplsnodeiccmaptable CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable

    // This table represents MPLS-TP specific extensions to mplsTunnelTable.  As
    // per MPLS-TP Identifiers [RFC6370], LSP_ID for IP based co-routed
    // bidirectional tunnel,  A1-{Global_ID::Node_ID::Tunnel_Num}::Z9-{Global_ID::
    // Node_ID::Tunnel_Num}::LSP_Num  LSP_ID for IP based associated bidirectional
    // tunnel, A1-{Global_ID::Node_ID::Tunnel_Num::LSP_Num}::
    // Z9-{Global_ID::Node_ID::Tunnel_Num::LSP_Num}  mplsTunnelTable is reused for
    // forming the LSP_ID as follows,  Source Tunnel_Num is mapped with
    // mplsTunnelIndex, Source Node_ID is mapped with mplsTunnelIngressLSRId,
    // Destination Node_ID is mapped with mplsTunnelEgressLSRId LSP_Num is mapped
    // with mplsTunnelInstance.  Source Global_Node_ID and/or ICC and Destination
    // Global_Node_ID and/or ICC are maintained in the mplsNodeConfigTable and
    // mplsNodeConfigLocalId is used to create an entry in mplsTunnelTable.
    Cmplstunnelexttable CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable

    // This table extends the mplsTunnelTable to provide per-tunnel packet
    // performance information for the reverse direction of a bidirectional
    // tunnel.  It can be seen as supplementing the mplsTunnelPerfTable, which
    // augments the mplsTunnelTable.  For links that do not transport packets,
    // these packet counters cannot be maintained.  For such links, attempts to
    // read the objects in this table will return noSuchInstance.
    Cmplstunnelreverseperftable CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable
}

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetFilter() yfilter.YFilter { return cISCOIETFMPLSTEEXTSTD03MIB.YFilter }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) SetFilter(yf yfilter.YFilter) { cISCOIETFMPLSTEEXTSTD03MIB.YFilter = yf }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetGoName(yname string) string {
    if yname == "cmplsNodeConfigTable" { return "Cmplsnodeconfigtable" }
    if yname == "cmplsNodeIpMapTable" { return "Cmplsnodeipmaptable" }
    if yname == "cmplsNodeIccMapTable" { return "Cmplsnodeiccmaptable" }
    if yname == "cmplsTunnelExtTable" { return "Cmplstunnelexttable" }
    if yname == "cmplsTunnelReversePerfTable" { return "Cmplstunnelreverseperftable" }
    return ""
}

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetSegmentPath() string {
    return "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB:CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
}

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsNodeConfigTable" {
        return &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeconfigtable
    }
    if childYangName == "cmplsNodeIpMapTable" {
        return &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeipmaptable
    }
    if childYangName == "cmplsNodeIccMapTable" {
        return &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeiccmaptable
    }
    if childYangName == "cmplsTunnelExtTable" {
        return &cISCOIETFMPLSTEEXTSTD03MIB.Cmplstunnelexttable
    }
    if childYangName == "cmplsTunnelReversePerfTable" {
        return &cISCOIETFMPLSTEEXTSTD03MIB.Cmplstunnelreverseperftable
    }
    return nil
}

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cmplsNodeConfigTable"] = &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeconfigtable
    children["cmplsNodeIpMapTable"] = &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeipmaptable
    children["cmplsNodeIccMapTable"] = &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeiccmaptable
    children["cmplsTunnelExtTable"] = &cISCOIETFMPLSTEEXTSTD03MIB.Cmplstunnelexttable
    children["cmplsTunnelReversePerfTable"] = &cISCOIETFMPLSTEEXTSTD03MIB.Cmplstunnelreverseperftable
    return children
}

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetYangName() string { return "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB" }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) SetParent(parent types.Entity) { cISCOIETFMPLSTEEXTSTD03MIB.parent = parent }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetParent() types.Entity { return cISCOIETFMPLSTEEXTSTD03MIB.parent }

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetParentYangName() string { return "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable
// This table allows the administrator to map a node or
// LSR Identifier (IP compatible [Global_Node_ID] or ICC)
// with a local identifier.
// 
// 
// This table is created to reuse the existing
// mplsTunnelTable for MPLS based transport network
// tunnels also.
// Since the MPLS tunnel's Ingress/Egress LSR identifiers'
// size (Unsigned32) value is not compatible for
// MPLS-TP tunnel i.e. Global_Node_Id of size 8 bytes and
// ICC of size 6 bytes, there exists a need to map the
// Global_Node_ID or ICC with the local identifier of size
// 4 bytes (Unsigned32) value in order
// to index (Ingress/Egress LSR identifier)
// the existing mplsTunnelTable.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping identification for the operator
    // or service provider with node or LSR.  As per [RFC6370], this mapping is 
    // represented as Global_Node_ID or ICC.  Note: Each entry in this table
    // should have a unique Global_ID and Node_ID combination. The type is slice
    // of CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry.
    Cmplsnodeconfigentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry
}

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetFilter() yfilter.YFilter { return cmplsnodeconfigtable.YFilter }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) SetFilter(yf yfilter.YFilter) { cmplsnodeconfigtable.YFilter = yf }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetGoName(yname string) string {
    if yname == "cmplsNodeConfigEntry" { return "Cmplsnodeconfigentry" }
    return ""
}

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetSegmentPath() string {
    return "cmplsNodeConfigTable"
}

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsNodeConfigEntry" {
        for _, c := range cmplsnodeconfigtable.Cmplsnodeconfigentry {
            if cmplsnodeconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry{}
        cmplsnodeconfigtable.Cmplsnodeconfigentry = append(cmplsnodeconfigtable.Cmplsnodeconfigentry, child)
        return &cmplsnodeconfigtable.Cmplsnodeconfigentry[len(cmplsnodeconfigtable.Cmplsnodeconfigentry)-1]
    }
    return nil
}

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplsnodeconfigtable.Cmplsnodeconfigentry {
        children[cmplsnodeconfigtable.Cmplsnodeconfigentry[i].GetSegmentPath()] = &cmplsnodeconfigtable.Cmplsnodeconfigentry[i]
    }
    return children
}

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetYangName() string { return "cmplsNodeConfigTable" }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) SetParent(parent types.Entity) { cmplsnodeconfigtable.parent = parent }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetParent() types.Entity { return cmplsnodeconfigtable.parent }

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetParentYangName() string { return "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry
// An entry in this table represents a mapping
// identification for the operator or service provider
// with node or LSR.
// 
// As per [RFC6370], this mapping is
// 
// represented as Global_Node_ID or ICC.
// 
// Note: Each entry in this table should have a unique
// Global_ID and Node_ID combination.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object allows the administrator to assign a
    // unique local identifier to map Global_Node_ID or ICC. The type is
    // interface{} with range: 1..16777215.
    Cmplsnodeconfiglocalid interface{}

    // This object indicates the Global Operator Identifier. This object value
    // should be zero when mplsNodeConfigIccId is configured with non-null value.
    // The type is string with length: 4.
    Cmplsnodeconfigglobalid interface{}

    // This object indicates the Node_ID within the operator. This object value
    // should be zero when mplsNodeConfigIccId is configured with non-null value.
    // The type is interface{} with range: 0..4294967295.
    Cmplsnodeconfignodeid interface{}

    // This object allows the operator or service provider to configure a unique
    // MPLS-TP ITU-T Carrier Code (ICC) either for Ingress ID or Egress ID.  This
    // object value should be zero when mplsNodeConfigGlobalId and
    // mplsNodeConfigNodeId are assigned with non-zero value. The type is string
    // with length: 1..6.
    Cmplsnodeconfigiccid interface{}

    // This object allows the administrator to create, modify, and/or delete a row
    // in this table. The type is RowStatus.
    Cmplsnodeconfigrowstatus interface{}

    // This variable indicates the storage type for this object. Conceptual rows
    // having the value 'permanent' need not allow write-access to any columnar
    // objects in the row. The type is StorageType.
    Cmplsnodeconfigstoragetype interface{}
}

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetFilter() yfilter.YFilter { return cmplsnodeconfigentry.YFilter }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) SetFilter(yf yfilter.YFilter) { cmplsnodeconfigentry.YFilter = yf }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetGoName(yname string) string {
    if yname == "cmplsNodeConfigLocalId" { return "Cmplsnodeconfiglocalid" }
    if yname == "cmplsNodeConfigGlobalId" { return "Cmplsnodeconfigglobalid" }
    if yname == "cmplsNodeConfigNodeId" { return "Cmplsnodeconfignodeid" }
    if yname == "cmplsNodeConfigIccId" { return "Cmplsnodeconfigiccid" }
    if yname == "cmplsNodeConfigRowStatus" { return "Cmplsnodeconfigrowstatus" }
    if yname == "cmplsNodeConfigStorageType" { return "Cmplsnodeconfigstoragetype" }
    return ""
}

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetSegmentPath() string {
    return "cmplsNodeConfigEntry" + "[cmplsNodeConfigLocalId='" + fmt.Sprintf("%v", cmplsnodeconfigentry.Cmplsnodeconfiglocalid) + "']"
}

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmplsNodeConfigLocalId"] = cmplsnodeconfigentry.Cmplsnodeconfiglocalid
    leafs["cmplsNodeConfigGlobalId"] = cmplsnodeconfigentry.Cmplsnodeconfigglobalid
    leafs["cmplsNodeConfigNodeId"] = cmplsnodeconfigentry.Cmplsnodeconfignodeid
    leafs["cmplsNodeConfigIccId"] = cmplsnodeconfigentry.Cmplsnodeconfigiccid
    leafs["cmplsNodeConfigRowStatus"] = cmplsnodeconfigentry.Cmplsnodeconfigrowstatus
    leafs["cmplsNodeConfigStorageType"] = cmplsnodeconfigentry.Cmplsnodeconfigstoragetype
    return leafs
}

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetYangName() string { return "cmplsNodeConfigEntry" }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) SetParent(parent types.Entity) { cmplsnodeconfigentry.parent = parent }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetParent() types.Entity { return cmplsnodeconfigentry.parent }

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetParentYangName() string { return "cmplsNodeConfigTable" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable
// This read-only table allows the administrator to retrieve
// the local identifier for a given Global_Node_ID in an IP
// compatible operator environment.
// 
// This table MAY be used in on-demand and/or proactive
// OAM operations to get the Ingress/Egress LSR identifier
// (Local Identifier) from Src-Global_Node_ID
// or Dst-Global_Node_ID and the Ingress and Egress LSR
// identifiers are used to retrieve the tunnel entry.
// 
// This table returns nothing when the associated entry
// is not defined in mplsNodeConfigTable.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping of Global_Node_ID with the
    // local identifier.  An entry in this table is created automatically when the
    // Local identifier is associated with Global_ID and Node_Id in the
    // mplsNodeConfigTable. Note: Each entry in this table should have a unique
    // Global_ID and Node_ID combination. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry.
    Cmplsnodeipmapentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry
}

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetFilter() yfilter.YFilter { return cmplsnodeipmaptable.YFilter }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) SetFilter(yf yfilter.YFilter) { cmplsnodeipmaptable.YFilter = yf }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetGoName(yname string) string {
    if yname == "cmplsNodeIpMapEntry" { return "Cmplsnodeipmapentry" }
    return ""
}

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetSegmentPath() string {
    return "cmplsNodeIpMapTable"
}

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsNodeIpMapEntry" {
        for _, c := range cmplsnodeipmaptable.Cmplsnodeipmapentry {
            if cmplsnodeipmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry{}
        cmplsnodeipmaptable.Cmplsnodeipmapentry = append(cmplsnodeipmaptable.Cmplsnodeipmapentry, child)
        return &cmplsnodeipmaptable.Cmplsnodeipmapentry[len(cmplsnodeipmaptable.Cmplsnodeipmapentry)-1]
    }
    return nil
}

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplsnodeipmaptable.Cmplsnodeipmapentry {
        children[cmplsnodeipmaptable.Cmplsnodeipmapentry[i].GetSegmentPath()] = &cmplsnodeipmaptable.Cmplsnodeipmapentry[i]
    }
    return children
}

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetYangName() string { return "cmplsNodeIpMapTable" }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) SetParent(parent types.Entity) { cmplsnodeipmaptable.parent = parent }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetParent() types.Entity { return cmplsnodeipmaptable.parent }

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetParentYangName() string { return "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry
// An entry in this table represents a mapping of
// Global_Node_ID with the local identifier.
// 
// An entry in this table is created automatically when
// the Local identifier is associated with Global_ID and
// Node_Id in the mplsNodeConfigTable.
// Note: Each entry in this table should have a unique
// Global_ID and Node_ID combination.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates the Global_ID. The type is
    // string with length: 4.
    Cmplsnodeipmapglobalid interface{}

    // This attribute is a key. This object indicates the Node_ID within the
    // operator. The type is interface{} with range: 0..4294967295.
    Cmplsnodeipmapnodeid interface{}

    // This object contains an IP compatible local identifier which is defined in
    // mplsNodeConfigTable. The type is interface{} with range: 1..16777215.
    Cmplsnodeipmaplocalid interface{}
}

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetFilter() yfilter.YFilter { return cmplsnodeipmapentry.YFilter }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) SetFilter(yf yfilter.YFilter) { cmplsnodeipmapentry.YFilter = yf }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetGoName(yname string) string {
    if yname == "cmplsNodeIpMapGlobalId" { return "Cmplsnodeipmapglobalid" }
    if yname == "cmplsNodeIpMapNodeId" { return "Cmplsnodeipmapnodeid" }
    if yname == "cmplsNodeIpMapLocalId" { return "Cmplsnodeipmaplocalid" }
    return ""
}

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetSegmentPath() string {
    return "cmplsNodeIpMapEntry" + "[cmplsNodeIpMapGlobalId='" + fmt.Sprintf("%v", cmplsnodeipmapentry.Cmplsnodeipmapglobalid) + "']" + "[cmplsNodeIpMapNodeId='" + fmt.Sprintf("%v", cmplsnodeipmapentry.Cmplsnodeipmapnodeid) + "']"
}

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmplsNodeIpMapGlobalId"] = cmplsnodeipmapentry.Cmplsnodeipmapglobalid
    leafs["cmplsNodeIpMapNodeId"] = cmplsnodeipmapentry.Cmplsnodeipmapnodeid
    leafs["cmplsNodeIpMapLocalId"] = cmplsnodeipmapentry.Cmplsnodeipmaplocalid
    return leafs
}

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetYangName() string { return "cmplsNodeIpMapEntry" }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) SetParent(parent types.Entity) { cmplsnodeipmapentry.parent = parent }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetParent() types.Entity { return cmplsnodeipmapentry.parent }

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetParentYangName() string { return "cmplsNodeIpMapTable" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable
// This read-only table allows the administrator to retrieve
// the local identifier for a given ICC operator in an ICC
// operator environment.
// 
// This table MAY be used in on-demand and/or proactive
// OAM operations to get the Ingress/Egress LSR
// identifier (Local Identifier) from Src-ICC
// or Dst-ICC and the Ingress and Egress LSR
// identifiers are used to retrieve the tunnel entry.
// This table returns nothing when the associated entry
// is not defined in mplsNodeConfigTable.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping of ICC with the local
    // identifier.  An entry in this table is created automatically when the Local
    // identifier is associated with ICC in the mplsNodeConfigTable. The type is
    // slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry.
    Cmplsnodeiccmapentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry
}

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetFilter() yfilter.YFilter { return cmplsnodeiccmaptable.YFilter }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) SetFilter(yf yfilter.YFilter) { cmplsnodeiccmaptable.YFilter = yf }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetGoName(yname string) string {
    if yname == "cmplsNodeIccMapEntry" { return "Cmplsnodeiccmapentry" }
    return ""
}

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetSegmentPath() string {
    return "cmplsNodeIccMapTable"
}

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsNodeIccMapEntry" {
        for _, c := range cmplsnodeiccmaptable.Cmplsnodeiccmapentry {
            if cmplsnodeiccmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry{}
        cmplsnodeiccmaptable.Cmplsnodeiccmapentry = append(cmplsnodeiccmaptable.Cmplsnodeiccmapentry, child)
        return &cmplsnodeiccmaptable.Cmplsnodeiccmapentry[len(cmplsnodeiccmaptable.Cmplsnodeiccmapentry)-1]
    }
    return nil
}

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplsnodeiccmaptable.Cmplsnodeiccmapentry {
        children[cmplsnodeiccmaptable.Cmplsnodeiccmapentry[i].GetSegmentPath()] = &cmplsnodeiccmaptable.Cmplsnodeiccmapentry[i]
    }
    return children
}

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetYangName() string { return "cmplsNodeIccMapTable" }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) SetParent(parent types.Entity) { cmplsnodeiccmaptable.parent = parent }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetParent() types.Entity { return cmplsnodeiccmaptable.parent }

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetParentYangName() string { return "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry
// An entry in this table represents a mapping of ICC with
// the local identifier.
// 
// An entry in this table is created automatically when
// the Local identifier is associated with ICC in
// the mplsNodeConfigTable.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object allows the operator or service
    // provider to configure a unique MPLS-TP ITU-T Carrier Code (ICC) either for
    // Ingress or Egress LSR ID.  The ICC is a string of one to six characters,
    // each character being either alphabetic (i.e.  A-Z) or numeric (i.e. 0-9)
    // characters. Alphabetic characters in the ICC should be represented with
    // upper case letters. The type is string with length: 1..6.
    Cmplsnodeiccmapiccid interface{}

    // This object contains an ICC based local identifier which is defined in
    // mplsNodeConfigTable. The type is interface{} with range: 1..16777215.
    Cmplsnodeiccmaplocalid interface{}
}

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetFilter() yfilter.YFilter { return cmplsnodeiccmapentry.YFilter }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) SetFilter(yf yfilter.YFilter) { cmplsnodeiccmapentry.YFilter = yf }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetGoName(yname string) string {
    if yname == "cmplsNodeIccMapIccId" { return "Cmplsnodeiccmapiccid" }
    if yname == "cmplsNodeIccMapLocalId" { return "Cmplsnodeiccmaplocalid" }
    return ""
}

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetSegmentPath() string {
    return "cmplsNodeIccMapEntry" + "[cmplsNodeIccMapIccId='" + fmt.Sprintf("%v", cmplsnodeiccmapentry.Cmplsnodeiccmapiccid) + "']"
}

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmplsNodeIccMapIccId"] = cmplsnodeiccmapentry.Cmplsnodeiccmapiccid
    leafs["cmplsNodeIccMapLocalId"] = cmplsnodeiccmapentry.Cmplsnodeiccmaplocalid
    return leafs
}

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetYangName() string { return "cmplsNodeIccMapEntry" }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) SetParent(parent types.Entity) { cmplsnodeiccmapentry.parent = parent }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetParent() types.Entity { return cmplsnodeiccmapentry.parent }

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetParentYangName() string { return "cmplsNodeIccMapTable" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable
// This table represents MPLS-TP specific extensions to
// mplsTunnelTable.
// 
// As per MPLS-TP Identifiers [RFC6370], LSP_ID for IP based
// co-routed bidirectional tunnel,
// 
// A1-{Global_ID::Node_ID::Tunnel_Num}::Z9-{Global_ID::
// Node_ID::Tunnel_Num}::LSP_Num
// 
// LSP_ID for IP based associated bidirectional tunnel,
// A1-{Global_ID::Node_ID::Tunnel_Num::LSP_Num}::
// Z9-{Global_ID::Node_ID::Tunnel_Num::LSP_Num}
// 
// mplsTunnelTable is reused for forming the LSP_ID
// as follows,
// 
// Source Tunnel_Num is mapped with mplsTunnelIndex,
// Source Node_ID is mapped with
// mplsTunnelIngressLSRId, Destination Node_ID is
// mapped with mplsTunnelEgressLSRId LSP_Num is mapped with
// mplsTunnelInstance.
// 
// Source Global_Node_ID and/or ICC and Destination
// Global_Node_ID and/or ICC are maintained in the
// mplsNodeConfigTable and mplsNodeConfigLocalId is
// used to create an entry in mplsTunnelTable.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table represents MPLS-TP specific additional tunnel
    // configurations. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry.
    Cmplstunnelextentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry
}

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetFilter() yfilter.YFilter { return cmplstunnelexttable.YFilter }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) SetFilter(yf yfilter.YFilter) { cmplstunnelexttable.YFilter = yf }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetGoName(yname string) string {
    if yname == "cmplsTunnelExtEntry" { return "Cmplstunnelextentry" }
    return ""
}

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetSegmentPath() string {
    return "cmplsTunnelExtTable"
}

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsTunnelExtEntry" {
        for _, c := range cmplstunnelexttable.Cmplstunnelextentry {
            if cmplstunnelexttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry{}
        cmplstunnelexttable.Cmplstunnelextentry = append(cmplstunnelexttable.Cmplstunnelextentry, child)
        return &cmplstunnelexttable.Cmplstunnelextentry[len(cmplstunnelexttable.Cmplstunnelextentry)-1]
    }
    return nil
}

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplstunnelexttable.Cmplstunnelextentry {
        children[cmplstunnelexttable.Cmplstunnelextentry[i].GetSegmentPath()] = &cmplstunnelexttable.Cmplstunnelextentry[i]
    }
    return children
}

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetYangName() string { return "cmplsTunnelExtTable" }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) SetParent(parent types.Entity) { cmplstunnelexttable.parent = parent }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetParent() types.Entity { return cmplstunnelexttable.parent }

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetParentYangName() string { return "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry
// An entry in this table represents MPLS-TP
// specific additional tunnel configurations.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelindex
    Mplstunnelindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelinstance
    Mplstunnelinstance interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelingresslsrid
    Mplstunnelingresslsrid interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelegresslsrid
    Mplstunnelegresslsrid interface{}

    // This object is applicable only for the bidirectional tunnel that has the
    // forward and reverse LSPs in the same tunnel or in the different tunnels. 
    // This object holds the opposite direction tunnel entry if the bidirectional
    // tunnel is setup by configuring two tunnel entries in mplsTunnelTable.  The
    // value of zeroDotZero indicates single tunnel entry is used for
    // bidirectional tunnel setup. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Cmplstunneloppositedirptr interface{}

    // Denotes whether or not this tunnel uses mplsTunnelOppositeDirPtr for
    // identifying the opposite direction tunnel information. Note that if this
    // variable is set to true then the mplsTunnelOppositeDirPtr should point to
    // the first accessible row of the opposite direction tunnel. The type is
    // bool.
    Cmplstunnelextoppositedirtnlvalid interface{}

    // This object is applicable only for the bidirectional tunnel that has the
    // forward and reverse LSPs in the same tunnel or in the different tunnels. 
    // This object holds the same value as that of the mplsTunnelIndex of
    // mplsTunnelEntry if the forward and reverse LSPs are in the same tunnel.
    // Otherwise, this object holds the value of the other direction associated
    // LSP's mplsTunnelIndex from a different tunnel.  The values of this object
    // and the mplsTunnelExtDestTnlLspIndex object together can be used to
    // identify an opposite direction LSP i.e. if the mplsTunnelIndex and
    // mplsTunnelInstance hold the value for forward LSP, this object and
    // mplsTunnelExtDestTnlLspIndex can be used to retrieve the reverse direction
    // LSP and vice versa.  This object and mplsTunnelExtDestTnlLspIndex values
    // provide the first two indices of tunnel entry and the remaining indices can
    // be derived as follows, if both the forward and reverse LSPs are present in
    // the same tunnel, the opposite direction LSP's Ingress and Egress Identifier
    // will be same for both the LSPs, else the Ingress and Egress Identifiers
    // should be swapped in order to index the other direction tunnel. The type is
    // interface{} with range: 0..65535.
    Cmplstunnelextdesttnlindex interface{}

    // This object is applicable only for the bidirectional tunnel that has the
    // forward and reverse LSPs in the same tunnel or in the different tunnels. 
    // This object should contain different value if both the forward and reverse
    // LSPs present in the same tunnel.  This object can contain same value or
    // different values if the forward and reverse LSPs present in the different
    // tunnels. The type is interface{} with range: 0..4294967295.
    Cmplstunnelextdesttnllspindex interface{}

    // Denotes whether or not this tunnel uses mplsTunnelExtDestTnlIndex and
    // mplsTunnelExtDestTnlLspIndex for identifying the opposite direction tunnel
    // information. Note that if this variable is set to true then the
    // mplsTunnelExtDestTnlIndex and mplsTunnelExtDestTnlLspIndex objects should
    // have the valid opposite direction tunnel indices. The type is bool.
    Cmplstunnelextdesttnlvalid interface{}
}

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetFilter() yfilter.YFilter { return cmplstunnelextentry.YFilter }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) SetFilter(yf yfilter.YFilter) { cmplstunnelextentry.YFilter = yf }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetGoName(yname string) string {
    if yname == "mplsTunnelIndex" { return "Mplstunnelindex" }
    if yname == "mplsTunnelInstance" { return "Mplstunnelinstance" }
    if yname == "mplsTunnelIngressLSRId" { return "Mplstunnelingresslsrid" }
    if yname == "mplsTunnelEgressLSRId" { return "Mplstunnelegresslsrid" }
    if yname == "cmplsTunnelOppositeDirPtr" { return "Cmplstunneloppositedirptr" }
    if yname == "cmplsTunnelExtOppositeDirTnlValid" { return "Cmplstunnelextoppositedirtnlvalid" }
    if yname == "cmplsTunnelExtDestTnlIndex" { return "Cmplstunnelextdesttnlindex" }
    if yname == "cmplsTunnelExtDestTnlLspIndex" { return "Cmplstunnelextdesttnllspindex" }
    if yname == "cmplsTunnelExtDestTnlValid" { return "Cmplstunnelextdesttnlvalid" }
    return ""
}

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetSegmentPath() string {
    return "cmplsTunnelExtEntry" + "[mplsTunnelIndex='" + fmt.Sprintf("%v", cmplstunnelextentry.Mplstunnelindex) + "']" + "[mplsTunnelInstance='" + fmt.Sprintf("%v", cmplstunnelextentry.Mplstunnelinstance) + "']" + "[mplsTunnelIngressLSRId='" + fmt.Sprintf("%v", cmplstunnelextentry.Mplstunnelingresslsrid) + "']" + "[mplsTunnelEgressLSRId='" + fmt.Sprintf("%v", cmplstunnelextentry.Mplstunnelegresslsrid) + "']"
}

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelIndex"] = cmplstunnelextentry.Mplstunnelindex
    leafs["mplsTunnelInstance"] = cmplstunnelextentry.Mplstunnelinstance
    leafs["mplsTunnelIngressLSRId"] = cmplstunnelextentry.Mplstunnelingresslsrid
    leafs["mplsTunnelEgressLSRId"] = cmplstunnelextentry.Mplstunnelegresslsrid
    leafs["cmplsTunnelOppositeDirPtr"] = cmplstunnelextentry.Cmplstunneloppositedirptr
    leafs["cmplsTunnelExtOppositeDirTnlValid"] = cmplstunnelextentry.Cmplstunnelextoppositedirtnlvalid
    leafs["cmplsTunnelExtDestTnlIndex"] = cmplstunnelextentry.Cmplstunnelextdesttnlindex
    leafs["cmplsTunnelExtDestTnlLspIndex"] = cmplstunnelextentry.Cmplstunnelextdesttnllspindex
    leafs["cmplsTunnelExtDestTnlValid"] = cmplstunnelextentry.Cmplstunnelextdesttnlvalid
    return leafs
}

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetYangName() string { return "cmplsTunnelExtEntry" }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) SetParent(parent types.Entity) { cmplstunnelextentry.parent = parent }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetParent() types.Entity { return cmplstunnelextentry.parent }

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetParentYangName() string { return "cmplsTunnelExtTable" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable
// This table extends the mplsTunnelTable to provide
// per-tunnel packet performance information for the reverse
// direction of a bidirectional tunnel.  It can be seen as
// supplementing the mplsTunnelPerfTable, which augments the
// mplsTunnelTable.
// 
// For links that do not transport packets, these packet
// counters cannot be maintained.  For such links, attempts
// to read the objects in this table will return
// noSuchInstance.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table is created by the LSR for every bidirectional MPLS
    // tunnel where packets are visible to the LSR. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry.
    Cmplstunnelreverseperfentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry
}

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetFilter() yfilter.YFilter { return cmplstunnelreverseperftable.YFilter }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) SetFilter(yf yfilter.YFilter) { cmplstunnelreverseperftable.YFilter = yf }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetGoName(yname string) string {
    if yname == "cmplsTunnelReversePerfEntry" { return "Cmplstunnelreverseperfentry" }
    return ""
}

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetSegmentPath() string {
    return "cmplsTunnelReversePerfTable"
}

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsTunnelReversePerfEntry" {
        for _, c := range cmplstunnelreverseperftable.Cmplstunnelreverseperfentry {
            if cmplstunnelreverseperftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry{}
        cmplstunnelreverseperftable.Cmplstunnelreverseperfentry = append(cmplstunnelreverseperftable.Cmplstunnelreverseperfentry, child)
        return &cmplstunnelreverseperftable.Cmplstunnelreverseperfentry[len(cmplstunnelreverseperftable.Cmplstunnelreverseperfentry)-1]
    }
    return nil
}

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplstunnelreverseperftable.Cmplstunnelreverseperfentry {
        children[cmplstunnelreverseperftable.Cmplstunnelreverseperfentry[i].GetSegmentPath()] = &cmplstunnelreverseperftable.Cmplstunnelreverseperfentry[i]
    }
    return children
}

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetYangName() string { return "cmplsTunnelReversePerfTable" }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) SetParent(parent types.Entity) { cmplstunnelreverseperftable.parent = parent }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetParent() types.Entity { return cmplstunnelreverseperftable.parent }

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetParentYangName() string { return "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB" }

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry
// An entry in this table is created by the LSR for every
// bidirectional MPLS tunnel where packets are visible to the
// LSR.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelindex
    Mplstunnelindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelinstance
    Mplstunnelinstance interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelingresslsrid
    Mplstunnelingresslsrid interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_Mplstunneltable_Mplstunnelentry_Mplstunnelegresslsrid
    Mplstunnelegresslsrid interface{}

    // Number of packets forwarded on the tunnel in the reverse direction if it is
    // bidirectional.  This object represents the 32-bit value of the least
    // significant part of the 64-bit value if both mplsTunnelReversePerfHCPackets
    // and this object are returned.  For links that do not transport packets,
    // this packet counter cannot be maintained.  For such links, this value will
    // return noSuchInstance. The type is interface{} with range: 0..4294967295.
    Cmplstunnelreverseperfpackets interface{}

    // High-capacity counter for number of packets forwarded on the tunnel in the
    // reverse direction if it is bidirectional.  For links that do not transport
    // packets, this packet counter cannot be maintained.  For such links, this
    // value will return noSuchInstance. The type is interface{} with range:
    // 0..18446744073709551615.
    Cmplstunnelreverseperfhcpackets interface{}

    // Number of errored packets received on the tunnel in the reverse direction
    // if it is bidirectional.  For links that do not transport packets, this
    // packet counter cannot be maintained.  For such links, this value will
    // return noSuchInstance. The type is interface{} with range: 0..4294967295.
    Cmplstunnelreverseperferrors interface{}

    // Number of bytes forwarded on the tunnel in the reverse direction if it is
    // bidirectional.  This object represents the 32-bit value of the least
    // significant part of the 64-bit value if both mplsTunnelReversePerfHCBytes
    // and this object are returned.  For links that do not transport packets,
    // this packet counter cannot be maintained.  For such links, this value will
    // return noSuchInstance. The type is interface{} with range: 0..4294967295.
    Cmplstunnelreverseperfbytes interface{}

    // High-capacity counter for number of bytes forwarded on the tunnel in the
    // reverse direction if it is bidirectional.  For links that do not transport
    // packets, this packet counter cannot be maintained.  For such links, this
    // value will return noSuchInstance. The type is interface{} with range:
    // 0..18446744073709551615.
    Cmplstunnelreverseperfhcbytes interface{}
}

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetFilter() yfilter.YFilter { return cmplstunnelreverseperfentry.YFilter }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) SetFilter(yf yfilter.YFilter) { cmplstunnelreverseperfentry.YFilter = yf }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetGoName(yname string) string {
    if yname == "mplsTunnelIndex" { return "Mplstunnelindex" }
    if yname == "mplsTunnelInstance" { return "Mplstunnelinstance" }
    if yname == "mplsTunnelIngressLSRId" { return "Mplstunnelingresslsrid" }
    if yname == "mplsTunnelEgressLSRId" { return "Mplstunnelegresslsrid" }
    if yname == "cmplsTunnelReversePerfPackets" { return "Cmplstunnelreverseperfpackets" }
    if yname == "cmplsTunnelReversePerfHCPackets" { return "Cmplstunnelreverseperfhcpackets" }
    if yname == "cmplsTunnelReversePerfErrors" { return "Cmplstunnelreverseperferrors" }
    if yname == "cmplsTunnelReversePerfBytes" { return "Cmplstunnelreverseperfbytes" }
    if yname == "cmplsTunnelReversePerfHCBytes" { return "Cmplstunnelreverseperfhcbytes" }
    return ""
}

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetSegmentPath() string {
    return "cmplsTunnelReversePerfEntry" + "[mplsTunnelIndex='" + fmt.Sprintf("%v", cmplstunnelreverseperfentry.Mplstunnelindex) + "']" + "[mplsTunnelInstance='" + fmt.Sprintf("%v", cmplstunnelreverseperfentry.Mplstunnelinstance) + "']" + "[mplsTunnelIngressLSRId='" + fmt.Sprintf("%v", cmplstunnelreverseperfentry.Mplstunnelingresslsrid) + "']" + "[mplsTunnelEgressLSRId='" + fmt.Sprintf("%v", cmplstunnelreverseperfentry.Mplstunnelegresslsrid) + "']"
}

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsTunnelIndex"] = cmplstunnelreverseperfentry.Mplstunnelindex
    leafs["mplsTunnelInstance"] = cmplstunnelreverseperfentry.Mplstunnelinstance
    leafs["mplsTunnelIngressLSRId"] = cmplstunnelreverseperfentry.Mplstunnelingresslsrid
    leafs["mplsTunnelEgressLSRId"] = cmplstunnelreverseperfentry.Mplstunnelegresslsrid
    leafs["cmplsTunnelReversePerfPackets"] = cmplstunnelreverseperfentry.Cmplstunnelreverseperfpackets
    leafs["cmplsTunnelReversePerfHCPackets"] = cmplstunnelreverseperfentry.Cmplstunnelreverseperfhcpackets
    leafs["cmplsTunnelReversePerfErrors"] = cmplstunnelreverseperfentry.Cmplstunnelreverseperferrors
    leafs["cmplsTunnelReversePerfBytes"] = cmplstunnelreverseperfentry.Cmplstunnelreverseperfbytes
    leafs["cmplsTunnelReversePerfHCBytes"] = cmplstunnelreverseperfentry.Cmplstunnelreverseperfhcbytes
    return leafs
}

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetYangName() string { return "cmplsTunnelReversePerfEntry" }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) SetParent(parent types.Entity) { cmplstunnelreverseperfentry.parent = parent }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetParent() types.Entity { return cmplstunnelreverseperfentry.parent }

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetParentYangName() string { return "cmplsTunnelReversePerfTable" }

