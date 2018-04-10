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
    EntityData types.CommonEntityData
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

func (cISCOIETFMPLSTEEXTSTD03MIB *CISCOIETFMPLSTEEXTSTD03MIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.YFilter = cISCOIETFMPLSTEEXTSTD03MIB.YFilter
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.YangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.SegmentPath = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB:CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children["cmplsNodeConfigTable"] = types.YChild{"Cmplsnodeconfigtable", &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeconfigtable}
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children["cmplsNodeIpMapTable"] = types.YChild{"Cmplsnodeipmaptable", &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeipmaptable}
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children["cmplsNodeIccMapTable"] = types.YChild{"Cmplsnodeiccmaptable", &cISCOIETFMPLSTEEXTSTD03MIB.Cmplsnodeiccmaptable}
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children["cmplsTunnelExtTable"] = types.YChild{"Cmplstunnelexttable", &cISCOIETFMPLSTEEXTSTD03MIB.Cmplstunnelexttable}
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children["cmplsTunnelReversePerfTable"] = types.YChild{"Cmplstunnelreverseperftable", &cISCOIETFMPLSTEEXTSTD03MIB.Cmplstunnelreverseperftable}
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFMPLSTEEXTSTD03MIB.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping identification for the operator
    // or service provider with node or LSR.  As per [RFC6370], this mapping is 
    // represented as Global_Node_ID or ICC.  Note: Each entry in this table
    // should have a unique Global_ID and Node_ID combination. The type is slice
    // of CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry.
    Cmplsnodeconfigentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry
}

func (cmplsnodeconfigtable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable) GetEntityData() *types.CommonEntityData {
    cmplsnodeconfigtable.EntityData.YFilter = cmplsnodeconfigtable.YFilter
    cmplsnodeconfigtable.EntityData.YangName = "cmplsNodeConfigTable"
    cmplsnodeconfigtable.EntityData.BundleName = "cisco_ios_xe"
    cmplsnodeconfigtable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplsnodeconfigtable.EntityData.SegmentPath = "cmplsNodeConfigTable"
    cmplsnodeconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsnodeconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsnodeconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsnodeconfigtable.EntityData.Children = make(map[string]types.YChild)
    cmplsnodeconfigtable.EntityData.Children["cmplsNodeConfigEntry"] = types.YChild{"Cmplsnodeconfigentry", nil}
    for i := range cmplsnodeconfigtable.Cmplsnodeconfigentry {
        cmplsnodeconfigtable.EntityData.Children[types.GetSegmentPath(&cmplsnodeconfigtable.Cmplsnodeconfigentry[i])] = types.YChild{"Cmplsnodeconfigentry", &cmplsnodeconfigtable.Cmplsnodeconfigentry[i]}
    }
    cmplsnodeconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplsnodeconfigtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cmplsnodeconfigentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeconfigtable_Cmplsnodeconfigentry) GetEntityData() *types.CommonEntityData {
    cmplsnodeconfigentry.EntityData.YFilter = cmplsnodeconfigentry.YFilter
    cmplsnodeconfigentry.EntityData.YangName = "cmplsNodeConfigEntry"
    cmplsnodeconfigentry.EntityData.BundleName = "cisco_ios_xe"
    cmplsnodeconfigentry.EntityData.ParentYangName = "cmplsNodeConfigTable"
    cmplsnodeconfigentry.EntityData.SegmentPath = "cmplsNodeConfigEntry" + "[cmplsNodeConfigLocalId='" + fmt.Sprintf("%v", cmplsnodeconfigentry.Cmplsnodeconfiglocalid) + "']"
    cmplsnodeconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsnodeconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsnodeconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsnodeconfigentry.EntityData.Children = make(map[string]types.YChild)
    cmplsnodeconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplsnodeconfigentry.EntityData.Leafs["cmplsNodeConfigLocalId"] = types.YLeaf{"Cmplsnodeconfiglocalid", cmplsnodeconfigentry.Cmplsnodeconfiglocalid}
    cmplsnodeconfigentry.EntityData.Leafs["cmplsNodeConfigGlobalId"] = types.YLeaf{"Cmplsnodeconfigglobalid", cmplsnodeconfigentry.Cmplsnodeconfigglobalid}
    cmplsnodeconfigentry.EntityData.Leafs["cmplsNodeConfigNodeId"] = types.YLeaf{"Cmplsnodeconfignodeid", cmplsnodeconfigentry.Cmplsnodeconfignodeid}
    cmplsnodeconfigentry.EntityData.Leafs["cmplsNodeConfigIccId"] = types.YLeaf{"Cmplsnodeconfigiccid", cmplsnodeconfigentry.Cmplsnodeconfigiccid}
    cmplsnodeconfigentry.EntityData.Leafs["cmplsNodeConfigRowStatus"] = types.YLeaf{"Cmplsnodeconfigrowstatus", cmplsnodeconfigentry.Cmplsnodeconfigrowstatus}
    cmplsnodeconfigentry.EntityData.Leafs["cmplsNodeConfigStorageType"] = types.YLeaf{"Cmplsnodeconfigstoragetype", cmplsnodeconfigentry.Cmplsnodeconfigstoragetype}
    return &(cmplsnodeconfigentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping of Global_Node_ID with the
    // local identifier.  An entry in this table is created automatically when the
    // Local identifier is associated with Global_ID and Node_Id in the
    // mplsNodeConfigTable. Note: Each entry in this table should have a unique
    // Global_ID and Node_ID combination. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry.
    Cmplsnodeipmapentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry
}

func (cmplsnodeipmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable) GetEntityData() *types.CommonEntityData {
    cmplsnodeipmaptable.EntityData.YFilter = cmplsnodeipmaptable.YFilter
    cmplsnodeipmaptable.EntityData.YangName = "cmplsNodeIpMapTable"
    cmplsnodeipmaptable.EntityData.BundleName = "cisco_ios_xe"
    cmplsnodeipmaptable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplsnodeipmaptable.EntityData.SegmentPath = "cmplsNodeIpMapTable"
    cmplsnodeipmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsnodeipmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsnodeipmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsnodeipmaptable.EntityData.Children = make(map[string]types.YChild)
    cmplsnodeipmaptable.EntityData.Children["cmplsNodeIpMapEntry"] = types.YChild{"Cmplsnodeipmapentry", nil}
    for i := range cmplsnodeipmaptable.Cmplsnodeipmapentry {
        cmplsnodeipmaptable.EntityData.Children[types.GetSegmentPath(&cmplsnodeipmaptable.Cmplsnodeipmapentry[i])] = types.YChild{"Cmplsnodeipmapentry", &cmplsnodeipmaptable.Cmplsnodeipmapentry[i]}
    }
    cmplsnodeipmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplsnodeipmaptable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cmplsnodeipmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeipmaptable_Cmplsnodeipmapentry) GetEntityData() *types.CommonEntityData {
    cmplsnodeipmapentry.EntityData.YFilter = cmplsnodeipmapentry.YFilter
    cmplsnodeipmapentry.EntityData.YangName = "cmplsNodeIpMapEntry"
    cmplsnodeipmapentry.EntityData.BundleName = "cisco_ios_xe"
    cmplsnodeipmapentry.EntityData.ParentYangName = "cmplsNodeIpMapTable"
    cmplsnodeipmapentry.EntityData.SegmentPath = "cmplsNodeIpMapEntry" + "[cmplsNodeIpMapGlobalId='" + fmt.Sprintf("%v", cmplsnodeipmapentry.Cmplsnodeipmapglobalid) + "']" + "[cmplsNodeIpMapNodeId='" + fmt.Sprintf("%v", cmplsnodeipmapentry.Cmplsnodeipmapnodeid) + "']"
    cmplsnodeipmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsnodeipmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsnodeipmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsnodeipmapentry.EntityData.Children = make(map[string]types.YChild)
    cmplsnodeipmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplsnodeipmapentry.EntityData.Leafs["cmplsNodeIpMapGlobalId"] = types.YLeaf{"Cmplsnodeipmapglobalid", cmplsnodeipmapentry.Cmplsnodeipmapglobalid}
    cmplsnodeipmapentry.EntityData.Leafs["cmplsNodeIpMapNodeId"] = types.YLeaf{"Cmplsnodeipmapnodeid", cmplsnodeipmapentry.Cmplsnodeipmapnodeid}
    cmplsnodeipmapentry.EntityData.Leafs["cmplsNodeIpMapLocalId"] = types.YLeaf{"Cmplsnodeipmaplocalid", cmplsnodeipmapentry.Cmplsnodeipmaplocalid}
    return &(cmplsnodeipmapentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping of ICC with the local
    // identifier.  An entry in this table is created automatically when the Local
    // identifier is associated with ICC in the mplsNodeConfigTable. The type is
    // slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry.
    Cmplsnodeiccmapentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry
}

func (cmplsnodeiccmaptable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable) GetEntityData() *types.CommonEntityData {
    cmplsnodeiccmaptable.EntityData.YFilter = cmplsnodeiccmaptable.YFilter
    cmplsnodeiccmaptable.EntityData.YangName = "cmplsNodeIccMapTable"
    cmplsnodeiccmaptable.EntityData.BundleName = "cisco_ios_xe"
    cmplsnodeiccmaptable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplsnodeiccmaptable.EntityData.SegmentPath = "cmplsNodeIccMapTable"
    cmplsnodeiccmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsnodeiccmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsnodeiccmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsnodeiccmaptable.EntityData.Children = make(map[string]types.YChild)
    cmplsnodeiccmaptable.EntityData.Children["cmplsNodeIccMapEntry"] = types.YChild{"Cmplsnodeiccmapentry", nil}
    for i := range cmplsnodeiccmaptable.Cmplsnodeiccmapentry {
        cmplsnodeiccmaptable.EntityData.Children[types.GetSegmentPath(&cmplsnodeiccmaptable.Cmplsnodeiccmapentry[i])] = types.YChild{"Cmplsnodeiccmapentry", &cmplsnodeiccmaptable.Cmplsnodeiccmapentry[i]}
    }
    cmplsnodeiccmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplsnodeiccmaptable.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry
// An entry in this table represents a mapping of ICC with
// the local identifier.
// 
// An entry in this table is created automatically when
// the Local identifier is associated with ICC in
// the mplsNodeConfigTable.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry struct {
    EntityData types.CommonEntityData
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

func (cmplsnodeiccmapentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplsnodeiccmaptable_Cmplsnodeiccmapentry) GetEntityData() *types.CommonEntityData {
    cmplsnodeiccmapentry.EntityData.YFilter = cmplsnodeiccmapentry.YFilter
    cmplsnodeiccmapentry.EntityData.YangName = "cmplsNodeIccMapEntry"
    cmplsnodeiccmapentry.EntityData.BundleName = "cisco_ios_xe"
    cmplsnodeiccmapentry.EntityData.ParentYangName = "cmplsNodeIccMapTable"
    cmplsnodeiccmapentry.EntityData.SegmentPath = "cmplsNodeIccMapEntry" + "[cmplsNodeIccMapIccId='" + fmt.Sprintf("%v", cmplsnodeiccmapentry.Cmplsnodeiccmapiccid) + "']"
    cmplsnodeiccmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsnodeiccmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsnodeiccmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsnodeiccmapentry.EntityData.Children = make(map[string]types.YChild)
    cmplsnodeiccmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplsnodeiccmapentry.EntityData.Leafs["cmplsNodeIccMapIccId"] = types.YLeaf{"Cmplsnodeiccmapiccid", cmplsnodeiccmapentry.Cmplsnodeiccmapiccid}
    cmplsnodeiccmapentry.EntityData.Leafs["cmplsNodeIccMapLocalId"] = types.YLeaf{"Cmplsnodeiccmaplocalid", cmplsnodeiccmapentry.Cmplsnodeiccmaplocalid}
    return &(cmplsnodeiccmapentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents MPLS-TP specific additional tunnel
    // configurations. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry.
    Cmplstunnelextentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry
}

func (cmplstunnelexttable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable) GetEntityData() *types.CommonEntityData {
    cmplstunnelexttable.EntityData.YFilter = cmplstunnelexttable.YFilter
    cmplstunnelexttable.EntityData.YangName = "cmplsTunnelExtTable"
    cmplstunnelexttable.EntityData.BundleName = "cisco_ios_xe"
    cmplstunnelexttable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplstunnelexttable.EntityData.SegmentPath = "cmplsTunnelExtTable"
    cmplstunnelexttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplstunnelexttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplstunnelexttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplstunnelexttable.EntityData.Children = make(map[string]types.YChild)
    cmplstunnelexttable.EntityData.Children["cmplsTunnelExtEntry"] = types.YChild{"Cmplstunnelextentry", nil}
    for i := range cmplstunnelexttable.Cmplstunnelextentry {
        cmplstunnelexttable.EntityData.Children[types.GetSegmentPath(&cmplstunnelexttable.Cmplstunnelextentry[i])] = types.YChild{"Cmplstunnelextentry", &cmplstunnelexttable.Cmplstunnelextentry[i]}
    }
    cmplstunnelexttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplstunnelexttable.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry
// An entry in this table represents MPLS-TP
// specific additional tunnel configurations.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (cmplstunnelextentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelexttable_Cmplstunnelextentry) GetEntityData() *types.CommonEntityData {
    cmplstunnelextentry.EntityData.YFilter = cmplstunnelextentry.YFilter
    cmplstunnelextentry.EntityData.YangName = "cmplsTunnelExtEntry"
    cmplstunnelextentry.EntityData.BundleName = "cisco_ios_xe"
    cmplstunnelextentry.EntityData.ParentYangName = "cmplsTunnelExtTable"
    cmplstunnelextentry.EntityData.SegmentPath = "cmplsTunnelExtEntry" + "[mplsTunnelIndex='" + fmt.Sprintf("%v", cmplstunnelextentry.Mplstunnelindex) + "']" + "[mplsTunnelInstance='" + fmt.Sprintf("%v", cmplstunnelextentry.Mplstunnelinstance) + "']" + "[mplsTunnelIngressLSRId='" + fmt.Sprintf("%v", cmplstunnelextentry.Mplstunnelingresslsrid) + "']" + "[mplsTunnelEgressLSRId='" + fmt.Sprintf("%v", cmplstunnelextentry.Mplstunnelegresslsrid) + "']"
    cmplstunnelextentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplstunnelextentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplstunnelextentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplstunnelextentry.EntityData.Children = make(map[string]types.YChild)
    cmplstunnelextentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplstunnelextentry.EntityData.Leafs["mplsTunnelIndex"] = types.YLeaf{"Mplstunnelindex", cmplstunnelextentry.Mplstunnelindex}
    cmplstunnelextentry.EntityData.Leafs["mplsTunnelInstance"] = types.YLeaf{"Mplstunnelinstance", cmplstunnelextentry.Mplstunnelinstance}
    cmplstunnelextentry.EntityData.Leafs["mplsTunnelIngressLSRId"] = types.YLeaf{"Mplstunnelingresslsrid", cmplstunnelextentry.Mplstunnelingresslsrid}
    cmplstunnelextentry.EntityData.Leafs["mplsTunnelEgressLSRId"] = types.YLeaf{"Mplstunnelegresslsrid", cmplstunnelextentry.Mplstunnelegresslsrid}
    cmplstunnelextentry.EntityData.Leafs["cmplsTunnelOppositeDirPtr"] = types.YLeaf{"Cmplstunneloppositedirptr", cmplstunnelextentry.Cmplstunneloppositedirptr}
    cmplstunnelextentry.EntityData.Leafs["cmplsTunnelExtOppositeDirTnlValid"] = types.YLeaf{"Cmplstunnelextoppositedirtnlvalid", cmplstunnelextentry.Cmplstunnelextoppositedirtnlvalid}
    cmplstunnelextentry.EntityData.Leafs["cmplsTunnelExtDestTnlIndex"] = types.YLeaf{"Cmplstunnelextdesttnlindex", cmplstunnelextentry.Cmplstunnelextdesttnlindex}
    cmplstunnelextentry.EntityData.Leafs["cmplsTunnelExtDestTnlLspIndex"] = types.YLeaf{"Cmplstunnelextdesttnllspindex", cmplstunnelextentry.Cmplstunnelextdesttnllspindex}
    cmplstunnelextentry.EntityData.Leafs["cmplsTunnelExtDestTnlValid"] = types.YLeaf{"Cmplstunnelextdesttnlvalid", cmplstunnelextentry.Cmplstunnelextdesttnlvalid}
    return &(cmplstunnelextentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the LSR for every bidirectional MPLS
    // tunnel where packets are visible to the LSR. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry.
    Cmplstunnelreverseperfentry []CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry
}

func (cmplstunnelreverseperftable *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable) GetEntityData() *types.CommonEntityData {
    cmplstunnelreverseperftable.EntityData.YFilter = cmplstunnelreverseperftable.YFilter
    cmplstunnelreverseperftable.EntityData.YangName = "cmplsTunnelReversePerfTable"
    cmplstunnelreverseperftable.EntityData.BundleName = "cisco_ios_xe"
    cmplstunnelreverseperftable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplstunnelreverseperftable.EntityData.SegmentPath = "cmplsTunnelReversePerfTable"
    cmplstunnelreverseperftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplstunnelreverseperftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplstunnelreverseperftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplstunnelreverseperftable.EntityData.Children = make(map[string]types.YChild)
    cmplstunnelreverseperftable.EntityData.Children["cmplsTunnelReversePerfEntry"] = types.YChild{"Cmplstunnelreverseperfentry", nil}
    for i := range cmplstunnelreverseperftable.Cmplstunnelreverseperfentry {
        cmplstunnelreverseperftable.EntityData.Children[types.GetSegmentPath(&cmplstunnelreverseperftable.Cmplstunnelreverseperfentry[i])] = types.YChild{"Cmplstunnelreverseperfentry", &cmplstunnelreverseperftable.Cmplstunnelreverseperfentry[i]}
    }
    cmplstunnelreverseperftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplstunnelreverseperftable.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry
// An entry in this table is created by the LSR for every
// bidirectional MPLS tunnel where packets are visible to the
// LSR.
type CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry struct {
    EntityData types.CommonEntityData
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

func (cmplstunnelreverseperfentry *CISCOIETFMPLSTEEXTSTD03MIB_Cmplstunnelreverseperftable_Cmplstunnelreverseperfentry) GetEntityData() *types.CommonEntityData {
    cmplstunnelreverseperfentry.EntityData.YFilter = cmplstunnelreverseperfentry.YFilter
    cmplstunnelreverseperfentry.EntityData.YangName = "cmplsTunnelReversePerfEntry"
    cmplstunnelreverseperfentry.EntityData.BundleName = "cisco_ios_xe"
    cmplstunnelreverseperfentry.EntityData.ParentYangName = "cmplsTunnelReversePerfTable"
    cmplstunnelreverseperfentry.EntityData.SegmentPath = "cmplsTunnelReversePerfEntry" + "[mplsTunnelIndex='" + fmt.Sprintf("%v", cmplstunnelreverseperfentry.Mplstunnelindex) + "']" + "[mplsTunnelInstance='" + fmt.Sprintf("%v", cmplstunnelreverseperfentry.Mplstunnelinstance) + "']" + "[mplsTunnelIngressLSRId='" + fmt.Sprintf("%v", cmplstunnelreverseperfentry.Mplstunnelingresslsrid) + "']" + "[mplsTunnelEgressLSRId='" + fmt.Sprintf("%v", cmplstunnelreverseperfentry.Mplstunnelegresslsrid) + "']"
    cmplstunnelreverseperfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplstunnelreverseperfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplstunnelreverseperfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplstunnelreverseperfentry.EntityData.Children = make(map[string]types.YChild)
    cmplstunnelreverseperfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplstunnelreverseperfentry.EntityData.Leafs["mplsTunnelIndex"] = types.YLeaf{"Mplstunnelindex", cmplstunnelreverseperfentry.Mplstunnelindex}
    cmplstunnelreverseperfentry.EntityData.Leafs["mplsTunnelInstance"] = types.YLeaf{"Mplstunnelinstance", cmplstunnelreverseperfentry.Mplstunnelinstance}
    cmplstunnelreverseperfentry.EntityData.Leafs["mplsTunnelIngressLSRId"] = types.YLeaf{"Mplstunnelingresslsrid", cmplstunnelreverseperfentry.Mplstunnelingresslsrid}
    cmplstunnelreverseperfentry.EntityData.Leafs["mplsTunnelEgressLSRId"] = types.YLeaf{"Mplstunnelegresslsrid", cmplstunnelreverseperfentry.Mplstunnelegresslsrid}
    cmplstunnelreverseperfentry.EntityData.Leafs["cmplsTunnelReversePerfPackets"] = types.YLeaf{"Cmplstunnelreverseperfpackets", cmplstunnelreverseperfentry.Cmplstunnelreverseperfpackets}
    cmplstunnelreverseperfentry.EntityData.Leafs["cmplsTunnelReversePerfHCPackets"] = types.YLeaf{"Cmplstunnelreverseperfhcpackets", cmplstunnelreverseperfentry.Cmplstunnelreverseperfhcpackets}
    cmplstunnelreverseperfentry.EntityData.Leafs["cmplsTunnelReversePerfErrors"] = types.YLeaf{"Cmplstunnelreverseperferrors", cmplstunnelreverseperfentry.Cmplstunnelreverseperferrors}
    cmplstunnelreverseperfentry.EntityData.Leafs["cmplsTunnelReversePerfBytes"] = types.YLeaf{"Cmplstunnelreverseperfbytes", cmplstunnelreverseperfentry.Cmplstunnelreverseperfbytes}
    cmplstunnelreverseperfentry.EntityData.Leafs["cmplsTunnelReversePerfHCBytes"] = types.YLeaf{"Cmplstunnelreverseperfhcbytes", cmplstunnelreverseperfentry.Cmplstunnelreverseperfhcbytes}
    return &(cmplstunnelreverseperfentry.EntityData)
}

