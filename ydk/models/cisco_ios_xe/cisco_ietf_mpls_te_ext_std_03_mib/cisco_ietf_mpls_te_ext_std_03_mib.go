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
    CmplsNodeConfigTable CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable

    // This read-only table allows the administrator to retrieve the local
    // identifier for a given Global_Node_ID in an IP compatible operator
    // environment.  This table MAY be used in on-demand and/or proactive OAM
    // operations to get the Ingress/Egress LSR identifier (Local Identifier) from
    // Src-Global_Node_ID or Dst-Global_Node_ID and the Ingress and Egress LSR
    // identifiers are used to retrieve the tunnel entry.  This table returns
    // nothing when the associated entry is not defined in mplsNodeConfigTable.
    CmplsNodeIpMapTable CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable

    // This read-only table allows the administrator to retrieve the local
    // identifier for a given ICC operator in an ICC operator environment.  This
    // table MAY be used in on-demand and/or proactive OAM operations to get the
    // Ingress/Egress LSR identifier (Local Identifier) from Src-ICC or Dst-ICC
    // and the Ingress and Egress LSR identifiers are used to retrieve the tunnel
    // entry. This table returns nothing when the associated entry is not defined
    // in mplsNodeConfigTable.
    CmplsNodeIccMapTable CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable

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
    CmplsTunnelExtTable CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable

    // This table extends the mplsTunnelTable to provide per-tunnel packet
    // performance information for the reverse direction of a bidirectional
    // tunnel.  It can be seen as supplementing the mplsTunnelPerfTable, which
    // augments the mplsTunnelTable.  For links that do not transport packets,
    // these packet counters cannot be maintained.  For such links, attempts to
    // read the objects in this table will return noSuchInstance.
    CmplsTunnelReversePerfTable CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable
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

    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children.Append("cmplsNodeConfigTable", types.YChild{"CmplsNodeConfigTable", &cISCOIETFMPLSTEEXTSTD03MIB.CmplsNodeConfigTable})
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children.Append("cmplsNodeIpMapTable", types.YChild{"CmplsNodeIpMapTable", &cISCOIETFMPLSTEEXTSTD03MIB.CmplsNodeIpMapTable})
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children.Append("cmplsNodeIccMapTable", types.YChild{"CmplsNodeIccMapTable", &cISCOIETFMPLSTEEXTSTD03MIB.CmplsNodeIccMapTable})
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children.Append("cmplsTunnelExtTable", types.YChild{"CmplsTunnelExtTable", &cISCOIETFMPLSTEEXTSTD03MIB.CmplsTunnelExtTable})
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Children.Append("cmplsTunnelReversePerfTable", types.YChild{"CmplsTunnelReversePerfTable", &cISCOIETFMPLSTEEXTSTD03MIB.CmplsTunnelReversePerfTable})
    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFMPLSTEEXTSTD03MIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFMPLSTEEXTSTD03MIB.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable
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
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping identification for the operator
    // or service provider with node or LSR.  As per [RFC6370], this mapping is 
    // represented as Global_Node_ID or ICC.  Note: Each entry in this table
    // should have a unique Global_ID and Node_ID combination. The type is slice
    // of CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable_CmplsNodeConfigEntry.
    CmplsNodeConfigEntry []*CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable_CmplsNodeConfigEntry
}

func (cmplsNodeConfigTable *CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable) GetEntityData() *types.CommonEntityData {
    cmplsNodeConfigTable.EntityData.YFilter = cmplsNodeConfigTable.YFilter
    cmplsNodeConfigTable.EntityData.YangName = "cmplsNodeConfigTable"
    cmplsNodeConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsNodeConfigTable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplsNodeConfigTable.EntityData.SegmentPath = "cmplsNodeConfigTable"
    cmplsNodeConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsNodeConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsNodeConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsNodeConfigTable.EntityData.Children = types.NewOrderedMap()
    cmplsNodeConfigTable.EntityData.Children.Append("cmplsNodeConfigEntry", types.YChild{"CmplsNodeConfigEntry", nil})
    for i := range cmplsNodeConfigTable.CmplsNodeConfigEntry {
        cmplsNodeConfigTable.EntityData.Children.Append(types.GetSegmentPath(cmplsNodeConfigTable.CmplsNodeConfigEntry[i]), types.YChild{"CmplsNodeConfigEntry", cmplsNodeConfigTable.CmplsNodeConfigEntry[i]})
    }
    cmplsNodeConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsNodeConfigTable.EntityData.YListKeys = []string {}

    return &(cmplsNodeConfigTable.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable_CmplsNodeConfigEntry
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
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable_CmplsNodeConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object allows the administrator to assign a
    // unique local identifier to map Global_Node_ID or ICC. The type is
    // interface{} with range: 1..16777215.
    CmplsNodeConfigLocalId interface{}

    // This object indicates the Global Operator Identifier. This object value
    // should be zero when mplsNodeConfigIccId is configured with non-null value.
    // The type is string with length: 4.
    CmplsNodeConfigGlobalId interface{}

    // This object indicates the Node_ID within the operator. This object value
    // should be zero when mplsNodeConfigIccId is configured with non-null value.
    // The type is interface{} with range: 0..4294967295.
    CmplsNodeConfigNodeId interface{}

    // This object allows the operator or service provider to configure a unique
    // MPLS-TP ITU-T Carrier Code (ICC) either for Ingress ID or Egress ID.  This
    // object value should be zero when mplsNodeConfigGlobalId and
    // mplsNodeConfigNodeId are assigned with non-zero value. The type is string
    // with length: 1..6.
    CmplsNodeConfigIccId interface{}

    // This object allows the administrator to create, modify, and/or delete a row
    // in this table. The type is RowStatus.
    CmplsNodeConfigRowStatus interface{}

    // This variable indicates the storage type for this object. Conceptual rows
    // having the value 'permanent' need not allow write-access to any columnar
    // objects in the row. The type is StorageType.
    CmplsNodeConfigStorageType interface{}
}

func (cmplsNodeConfigEntry *CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeConfigTable_CmplsNodeConfigEntry) GetEntityData() *types.CommonEntityData {
    cmplsNodeConfigEntry.EntityData.YFilter = cmplsNodeConfigEntry.YFilter
    cmplsNodeConfigEntry.EntityData.YangName = "cmplsNodeConfigEntry"
    cmplsNodeConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsNodeConfigEntry.EntityData.ParentYangName = "cmplsNodeConfigTable"
    cmplsNodeConfigEntry.EntityData.SegmentPath = "cmplsNodeConfigEntry" + types.AddKeyToken(cmplsNodeConfigEntry.CmplsNodeConfigLocalId, "cmplsNodeConfigLocalId")
    cmplsNodeConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsNodeConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsNodeConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsNodeConfigEntry.EntityData.Children = types.NewOrderedMap()
    cmplsNodeConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsNodeConfigEntry.EntityData.Leafs.Append("cmplsNodeConfigLocalId", types.YLeaf{"CmplsNodeConfigLocalId", cmplsNodeConfigEntry.CmplsNodeConfigLocalId})
    cmplsNodeConfigEntry.EntityData.Leafs.Append("cmplsNodeConfigGlobalId", types.YLeaf{"CmplsNodeConfigGlobalId", cmplsNodeConfigEntry.CmplsNodeConfigGlobalId})
    cmplsNodeConfigEntry.EntityData.Leafs.Append("cmplsNodeConfigNodeId", types.YLeaf{"CmplsNodeConfigNodeId", cmplsNodeConfigEntry.CmplsNodeConfigNodeId})
    cmplsNodeConfigEntry.EntityData.Leafs.Append("cmplsNodeConfigIccId", types.YLeaf{"CmplsNodeConfigIccId", cmplsNodeConfigEntry.CmplsNodeConfigIccId})
    cmplsNodeConfigEntry.EntityData.Leafs.Append("cmplsNodeConfigRowStatus", types.YLeaf{"CmplsNodeConfigRowStatus", cmplsNodeConfigEntry.CmplsNodeConfigRowStatus})
    cmplsNodeConfigEntry.EntityData.Leafs.Append("cmplsNodeConfigStorageType", types.YLeaf{"CmplsNodeConfigStorageType", cmplsNodeConfigEntry.CmplsNodeConfigStorageType})

    cmplsNodeConfigEntry.EntityData.YListKeys = []string {"CmplsNodeConfigLocalId"}

    return &(cmplsNodeConfigEntry.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable
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
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping of Global_Node_ID with the
    // local identifier.  An entry in this table is created automatically when the
    // Local identifier is associated with Global_ID and Node_Id in the
    // mplsNodeConfigTable. Note: Each entry in this table should have a unique
    // Global_ID and Node_ID combination. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable_CmplsNodeIpMapEntry.
    CmplsNodeIpMapEntry []*CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable_CmplsNodeIpMapEntry
}

func (cmplsNodeIpMapTable *CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable) GetEntityData() *types.CommonEntityData {
    cmplsNodeIpMapTable.EntityData.YFilter = cmplsNodeIpMapTable.YFilter
    cmplsNodeIpMapTable.EntityData.YangName = "cmplsNodeIpMapTable"
    cmplsNodeIpMapTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsNodeIpMapTable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplsNodeIpMapTable.EntityData.SegmentPath = "cmplsNodeIpMapTable"
    cmplsNodeIpMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsNodeIpMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsNodeIpMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsNodeIpMapTable.EntityData.Children = types.NewOrderedMap()
    cmplsNodeIpMapTable.EntityData.Children.Append("cmplsNodeIpMapEntry", types.YChild{"CmplsNodeIpMapEntry", nil})
    for i := range cmplsNodeIpMapTable.CmplsNodeIpMapEntry {
        cmplsNodeIpMapTable.EntityData.Children.Append(types.GetSegmentPath(cmplsNodeIpMapTable.CmplsNodeIpMapEntry[i]), types.YChild{"CmplsNodeIpMapEntry", cmplsNodeIpMapTable.CmplsNodeIpMapEntry[i]})
    }
    cmplsNodeIpMapTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsNodeIpMapTable.EntityData.YListKeys = []string {}

    return &(cmplsNodeIpMapTable.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable_CmplsNodeIpMapEntry
// An entry in this table represents a mapping of
// Global_Node_ID with the local identifier.
// 
// An entry in this table is created automatically when
// the Local identifier is associated with Global_ID and
// Node_Id in the mplsNodeConfigTable.
// Note: Each entry in this table should have a unique
// Global_ID and Node_ID combination.
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable_CmplsNodeIpMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates the Global_ID. The type is
    // string with length: 4.
    CmplsNodeIpMapGlobalId interface{}

    // This attribute is a key. This object indicates the Node_ID within the
    // operator. The type is interface{} with range: 0..4294967295.
    CmplsNodeIpMapNodeId interface{}

    // This object contains an IP compatible local identifier which is defined in
    // mplsNodeConfigTable. The type is interface{} with range: 1..16777215.
    CmplsNodeIpMapLocalId interface{}
}

func (cmplsNodeIpMapEntry *CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIpMapTable_CmplsNodeIpMapEntry) GetEntityData() *types.CommonEntityData {
    cmplsNodeIpMapEntry.EntityData.YFilter = cmplsNodeIpMapEntry.YFilter
    cmplsNodeIpMapEntry.EntityData.YangName = "cmplsNodeIpMapEntry"
    cmplsNodeIpMapEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsNodeIpMapEntry.EntityData.ParentYangName = "cmplsNodeIpMapTable"
    cmplsNodeIpMapEntry.EntityData.SegmentPath = "cmplsNodeIpMapEntry" + types.AddKeyToken(cmplsNodeIpMapEntry.CmplsNodeIpMapGlobalId, "cmplsNodeIpMapGlobalId") + types.AddKeyToken(cmplsNodeIpMapEntry.CmplsNodeIpMapNodeId, "cmplsNodeIpMapNodeId")
    cmplsNodeIpMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsNodeIpMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsNodeIpMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsNodeIpMapEntry.EntityData.Children = types.NewOrderedMap()
    cmplsNodeIpMapEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsNodeIpMapEntry.EntityData.Leafs.Append("cmplsNodeIpMapGlobalId", types.YLeaf{"CmplsNodeIpMapGlobalId", cmplsNodeIpMapEntry.CmplsNodeIpMapGlobalId})
    cmplsNodeIpMapEntry.EntityData.Leafs.Append("cmplsNodeIpMapNodeId", types.YLeaf{"CmplsNodeIpMapNodeId", cmplsNodeIpMapEntry.CmplsNodeIpMapNodeId})
    cmplsNodeIpMapEntry.EntityData.Leafs.Append("cmplsNodeIpMapLocalId", types.YLeaf{"CmplsNodeIpMapLocalId", cmplsNodeIpMapEntry.CmplsNodeIpMapLocalId})

    cmplsNodeIpMapEntry.EntityData.YListKeys = []string {"CmplsNodeIpMapGlobalId", "CmplsNodeIpMapNodeId"}

    return &(cmplsNodeIpMapEntry.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable
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
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents a mapping of ICC with the local
    // identifier.  An entry in this table is created automatically when the Local
    // identifier is associated with ICC in the mplsNodeConfigTable. The type is
    // slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable_CmplsNodeIccMapEntry.
    CmplsNodeIccMapEntry []*CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable_CmplsNodeIccMapEntry
}

func (cmplsNodeIccMapTable *CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable) GetEntityData() *types.CommonEntityData {
    cmplsNodeIccMapTable.EntityData.YFilter = cmplsNodeIccMapTable.YFilter
    cmplsNodeIccMapTable.EntityData.YangName = "cmplsNodeIccMapTable"
    cmplsNodeIccMapTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsNodeIccMapTable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplsNodeIccMapTable.EntityData.SegmentPath = "cmplsNodeIccMapTable"
    cmplsNodeIccMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsNodeIccMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsNodeIccMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsNodeIccMapTable.EntityData.Children = types.NewOrderedMap()
    cmplsNodeIccMapTable.EntityData.Children.Append("cmplsNodeIccMapEntry", types.YChild{"CmplsNodeIccMapEntry", nil})
    for i := range cmplsNodeIccMapTable.CmplsNodeIccMapEntry {
        cmplsNodeIccMapTable.EntityData.Children.Append(types.GetSegmentPath(cmplsNodeIccMapTable.CmplsNodeIccMapEntry[i]), types.YChild{"CmplsNodeIccMapEntry", cmplsNodeIccMapTable.CmplsNodeIccMapEntry[i]})
    }
    cmplsNodeIccMapTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsNodeIccMapTable.EntityData.YListKeys = []string {}

    return &(cmplsNodeIccMapTable.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable_CmplsNodeIccMapEntry
// An entry in this table represents a mapping of ICC with
// the local identifier.
// 
// An entry in this table is created automatically when
// the Local identifier is associated with ICC in
// the mplsNodeConfigTable.
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable_CmplsNodeIccMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object allows the operator or service
    // provider to configure a unique MPLS-TP ITU-T Carrier Code (ICC) either for
    // Ingress or Egress LSR ID.  The ICC is a string of one to six characters,
    // each character being either alphabetic (i.e.  A-Z) or numeric (i.e. 0-9)
    // characters. Alphabetic characters in the ICC should be represented with
    // upper case letters. The type is string with length: 1..6.
    CmplsNodeIccMapIccId interface{}

    // This object contains an ICC based local identifier which is defined in
    // mplsNodeConfigTable. The type is interface{} with range: 1..16777215.
    CmplsNodeIccMapLocalId interface{}
}

func (cmplsNodeIccMapEntry *CISCOIETFMPLSTEEXTSTD03MIB_CmplsNodeIccMapTable_CmplsNodeIccMapEntry) GetEntityData() *types.CommonEntityData {
    cmplsNodeIccMapEntry.EntityData.YFilter = cmplsNodeIccMapEntry.YFilter
    cmplsNodeIccMapEntry.EntityData.YangName = "cmplsNodeIccMapEntry"
    cmplsNodeIccMapEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsNodeIccMapEntry.EntityData.ParentYangName = "cmplsNodeIccMapTable"
    cmplsNodeIccMapEntry.EntityData.SegmentPath = "cmplsNodeIccMapEntry" + types.AddKeyToken(cmplsNodeIccMapEntry.CmplsNodeIccMapIccId, "cmplsNodeIccMapIccId")
    cmplsNodeIccMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsNodeIccMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsNodeIccMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsNodeIccMapEntry.EntityData.Children = types.NewOrderedMap()
    cmplsNodeIccMapEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsNodeIccMapEntry.EntityData.Leafs.Append("cmplsNodeIccMapIccId", types.YLeaf{"CmplsNodeIccMapIccId", cmplsNodeIccMapEntry.CmplsNodeIccMapIccId})
    cmplsNodeIccMapEntry.EntityData.Leafs.Append("cmplsNodeIccMapLocalId", types.YLeaf{"CmplsNodeIccMapLocalId", cmplsNodeIccMapEntry.CmplsNodeIccMapLocalId})

    cmplsNodeIccMapEntry.EntityData.YListKeys = []string {"CmplsNodeIccMapIccId"}

    return &(cmplsNodeIccMapEntry.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable
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
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table represents MPLS-TP specific additional tunnel
    // configurations. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable_CmplsTunnelExtEntry.
    CmplsTunnelExtEntry []*CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable_CmplsTunnelExtEntry
}

func (cmplsTunnelExtTable *CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable) GetEntityData() *types.CommonEntityData {
    cmplsTunnelExtTable.EntityData.YFilter = cmplsTunnelExtTable.YFilter
    cmplsTunnelExtTable.EntityData.YangName = "cmplsTunnelExtTable"
    cmplsTunnelExtTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsTunnelExtTable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplsTunnelExtTable.EntityData.SegmentPath = "cmplsTunnelExtTable"
    cmplsTunnelExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsTunnelExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsTunnelExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsTunnelExtTable.EntityData.Children = types.NewOrderedMap()
    cmplsTunnelExtTable.EntityData.Children.Append("cmplsTunnelExtEntry", types.YChild{"CmplsTunnelExtEntry", nil})
    for i := range cmplsTunnelExtTable.CmplsTunnelExtEntry {
        cmplsTunnelExtTable.EntityData.Children.Append(types.GetSegmentPath(cmplsTunnelExtTable.CmplsTunnelExtEntry[i]), types.YChild{"CmplsTunnelExtEntry", cmplsTunnelExtTable.CmplsTunnelExtEntry[i]})
    }
    cmplsTunnelExtTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsTunnelExtTable.EntityData.YListKeys = []string {}

    return &(cmplsTunnelExtTable.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable_CmplsTunnelExtEntry
// An entry in this table represents MPLS-TP
// specific additional tunnel configurations.
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable_CmplsTunnelExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelIndex
    MplsTunnelIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelInstance
    MplsTunnelInstance interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelIngressLSRId
    MplsTunnelIngressLSRId interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelEgressLSRId
    MplsTunnelEgressLSRId interface{}

    // This object is applicable only for the bidirectional tunnel that has the
    // forward and reverse LSPs in the same tunnel or in the different tunnels. 
    // This object holds the opposite direction tunnel entry if the bidirectional
    // tunnel is setup by configuring two tunnel entries in mplsTunnelTable.  The
    // value of zeroDotZero indicates single tunnel entry is used for
    // bidirectional tunnel setup. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CmplsTunnelOppositeDirPtr interface{}

    // Denotes whether or not this tunnel uses mplsTunnelOppositeDirPtr for
    // identifying the opposite direction tunnel information. Note that if this
    // variable is set to true then the mplsTunnelOppositeDirPtr should point to
    // the first accessible row of the opposite direction tunnel. The type is
    // bool.
    CmplsTunnelExtOppositeDirTnlValid interface{}

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
    CmplsTunnelExtDestTnlIndex interface{}

    // This object is applicable only for the bidirectional tunnel that has the
    // forward and reverse LSPs in the same tunnel or in the different tunnels. 
    // This object should contain different value if both the forward and reverse
    // LSPs present in the same tunnel.  This object can contain same value or
    // different values if the forward and reverse LSPs present in the different
    // tunnels. The type is interface{} with range: 0..4294967295.
    CmplsTunnelExtDestTnlLspIndex interface{}

    // Denotes whether or not this tunnel uses mplsTunnelExtDestTnlIndex and
    // mplsTunnelExtDestTnlLspIndex for identifying the opposite direction tunnel
    // information. Note that if this variable is set to true then the
    // mplsTunnelExtDestTnlIndex and mplsTunnelExtDestTnlLspIndex objects should
    // have the valid opposite direction tunnel indices. The type is bool.
    CmplsTunnelExtDestTnlValid interface{}
}

func (cmplsTunnelExtEntry *CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelExtTable_CmplsTunnelExtEntry) GetEntityData() *types.CommonEntityData {
    cmplsTunnelExtEntry.EntityData.YFilter = cmplsTunnelExtEntry.YFilter
    cmplsTunnelExtEntry.EntityData.YangName = "cmplsTunnelExtEntry"
    cmplsTunnelExtEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsTunnelExtEntry.EntityData.ParentYangName = "cmplsTunnelExtTable"
    cmplsTunnelExtEntry.EntityData.SegmentPath = "cmplsTunnelExtEntry" + types.AddKeyToken(cmplsTunnelExtEntry.MplsTunnelIndex, "mplsTunnelIndex") + types.AddKeyToken(cmplsTunnelExtEntry.MplsTunnelInstance, "mplsTunnelInstance") + types.AddKeyToken(cmplsTunnelExtEntry.MplsTunnelIngressLSRId, "mplsTunnelIngressLSRId") + types.AddKeyToken(cmplsTunnelExtEntry.MplsTunnelEgressLSRId, "mplsTunnelEgressLSRId")
    cmplsTunnelExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsTunnelExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsTunnelExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsTunnelExtEntry.EntityData.Children = types.NewOrderedMap()
    cmplsTunnelExtEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsTunnelExtEntry.EntityData.Leafs.Append("mplsTunnelIndex", types.YLeaf{"MplsTunnelIndex", cmplsTunnelExtEntry.MplsTunnelIndex})
    cmplsTunnelExtEntry.EntityData.Leafs.Append("mplsTunnelInstance", types.YLeaf{"MplsTunnelInstance", cmplsTunnelExtEntry.MplsTunnelInstance})
    cmplsTunnelExtEntry.EntityData.Leafs.Append("mplsTunnelIngressLSRId", types.YLeaf{"MplsTunnelIngressLSRId", cmplsTunnelExtEntry.MplsTunnelIngressLSRId})
    cmplsTunnelExtEntry.EntityData.Leafs.Append("mplsTunnelEgressLSRId", types.YLeaf{"MplsTunnelEgressLSRId", cmplsTunnelExtEntry.MplsTunnelEgressLSRId})
    cmplsTunnelExtEntry.EntityData.Leafs.Append("cmplsTunnelOppositeDirPtr", types.YLeaf{"CmplsTunnelOppositeDirPtr", cmplsTunnelExtEntry.CmplsTunnelOppositeDirPtr})
    cmplsTunnelExtEntry.EntityData.Leafs.Append("cmplsTunnelExtOppositeDirTnlValid", types.YLeaf{"CmplsTunnelExtOppositeDirTnlValid", cmplsTunnelExtEntry.CmplsTunnelExtOppositeDirTnlValid})
    cmplsTunnelExtEntry.EntityData.Leafs.Append("cmplsTunnelExtDestTnlIndex", types.YLeaf{"CmplsTunnelExtDestTnlIndex", cmplsTunnelExtEntry.CmplsTunnelExtDestTnlIndex})
    cmplsTunnelExtEntry.EntityData.Leafs.Append("cmplsTunnelExtDestTnlLspIndex", types.YLeaf{"CmplsTunnelExtDestTnlLspIndex", cmplsTunnelExtEntry.CmplsTunnelExtDestTnlLspIndex})
    cmplsTunnelExtEntry.EntityData.Leafs.Append("cmplsTunnelExtDestTnlValid", types.YLeaf{"CmplsTunnelExtDestTnlValid", cmplsTunnelExtEntry.CmplsTunnelExtDestTnlValid})

    cmplsTunnelExtEntry.EntityData.YListKeys = []string {"MplsTunnelIndex", "MplsTunnelInstance", "MplsTunnelIngressLSRId", "MplsTunnelEgressLSRId"}

    return &(cmplsTunnelExtEntry.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable
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
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table is created by the LSR for every bidirectional MPLS
    // tunnel where packets are visible to the LSR. The type is slice of
    // CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable_CmplsTunnelReversePerfEntry.
    CmplsTunnelReversePerfEntry []*CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable_CmplsTunnelReversePerfEntry
}

func (cmplsTunnelReversePerfTable *CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable) GetEntityData() *types.CommonEntityData {
    cmplsTunnelReversePerfTable.EntityData.YFilter = cmplsTunnelReversePerfTable.YFilter
    cmplsTunnelReversePerfTable.EntityData.YangName = "cmplsTunnelReversePerfTable"
    cmplsTunnelReversePerfTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsTunnelReversePerfTable.EntityData.ParentYangName = "CISCO-IETF-MPLS-TE-EXT-STD-03-MIB"
    cmplsTunnelReversePerfTable.EntityData.SegmentPath = "cmplsTunnelReversePerfTable"
    cmplsTunnelReversePerfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsTunnelReversePerfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsTunnelReversePerfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsTunnelReversePerfTable.EntityData.Children = types.NewOrderedMap()
    cmplsTunnelReversePerfTable.EntityData.Children.Append("cmplsTunnelReversePerfEntry", types.YChild{"CmplsTunnelReversePerfEntry", nil})
    for i := range cmplsTunnelReversePerfTable.CmplsTunnelReversePerfEntry {
        cmplsTunnelReversePerfTable.EntityData.Children.Append(types.GetSegmentPath(cmplsTunnelReversePerfTable.CmplsTunnelReversePerfEntry[i]), types.YChild{"CmplsTunnelReversePerfEntry", cmplsTunnelReversePerfTable.CmplsTunnelReversePerfEntry[i]})
    }
    cmplsTunnelReversePerfTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsTunnelReversePerfTable.EntityData.YListKeys = []string {}

    return &(cmplsTunnelReversePerfTable.EntityData)
}

// CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable_CmplsTunnelReversePerfEntry
// An entry in this table is created by the LSR for every
// bidirectional MPLS tunnel where packets are visible to the
// LSR.
type CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable_CmplsTunnelReversePerfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelIndex
    MplsTunnelIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelInstance
    MplsTunnelInstance interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelIngressLSRId
    MplsTunnelIngressLSRId interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // mpls_te_std_mib.MPLSTESTDMIB_MplsTunnelTable_MplsTunnelEntry_MplsTunnelEgressLSRId
    MplsTunnelEgressLSRId interface{}

    // Number of packets forwarded on the tunnel in the reverse direction if it is
    // bidirectional.  This object represents the 32-bit value of the least
    // significant part of the 64-bit value if both mplsTunnelReversePerfHCPackets
    // and this object are returned.  For links that do not transport packets,
    // this packet counter cannot be maintained.  For such links, this value will
    // return noSuchInstance. The type is interface{} with range: 0..4294967295.
    CmplsTunnelReversePerfPackets interface{}

    // High-capacity counter for number of packets forwarded on the tunnel in the
    // reverse direction if it is bidirectional.  For links that do not transport
    // packets, this packet counter cannot be maintained.  For such links, this
    // value will return noSuchInstance. The type is interface{} with range:
    // 0..18446744073709551615.
    CmplsTunnelReversePerfHCPackets interface{}

    // Number of errored packets received on the tunnel in the reverse direction
    // if it is bidirectional.  For links that do not transport packets, this
    // packet counter cannot be maintained.  For such links, this value will
    // return noSuchInstance. The type is interface{} with range: 0..4294967295.
    CmplsTunnelReversePerfErrors interface{}

    // Number of bytes forwarded on the tunnel in the reverse direction if it is
    // bidirectional.  This object represents the 32-bit value of the least
    // significant part of the 64-bit value if both mplsTunnelReversePerfHCBytes
    // and this object are returned.  For links that do not transport packets,
    // this packet counter cannot be maintained.  For such links, this value will
    // return noSuchInstance. The type is interface{} with range: 0..4294967295.
    CmplsTunnelReversePerfBytes interface{}

    // High-capacity counter for number of bytes forwarded on the tunnel in the
    // reverse direction if it is bidirectional.  For links that do not transport
    // packets, this packet counter cannot be maintained.  For such links, this
    // value will return noSuchInstance. The type is interface{} with range:
    // 0..18446744073709551615.
    CmplsTunnelReversePerfHCBytes interface{}
}

func (cmplsTunnelReversePerfEntry *CISCOIETFMPLSTEEXTSTD03MIB_CmplsTunnelReversePerfTable_CmplsTunnelReversePerfEntry) GetEntityData() *types.CommonEntityData {
    cmplsTunnelReversePerfEntry.EntityData.YFilter = cmplsTunnelReversePerfEntry.YFilter
    cmplsTunnelReversePerfEntry.EntityData.YangName = "cmplsTunnelReversePerfEntry"
    cmplsTunnelReversePerfEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsTunnelReversePerfEntry.EntityData.ParentYangName = "cmplsTunnelReversePerfTable"
    cmplsTunnelReversePerfEntry.EntityData.SegmentPath = "cmplsTunnelReversePerfEntry" + types.AddKeyToken(cmplsTunnelReversePerfEntry.MplsTunnelIndex, "mplsTunnelIndex") + types.AddKeyToken(cmplsTunnelReversePerfEntry.MplsTunnelInstance, "mplsTunnelInstance") + types.AddKeyToken(cmplsTunnelReversePerfEntry.MplsTunnelIngressLSRId, "mplsTunnelIngressLSRId") + types.AddKeyToken(cmplsTunnelReversePerfEntry.MplsTunnelEgressLSRId, "mplsTunnelEgressLSRId")
    cmplsTunnelReversePerfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsTunnelReversePerfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsTunnelReversePerfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsTunnelReversePerfEntry.EntityData.Children = types.NewOrderedMap()
    cmplsTunnelReversePerfEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("mplsTunnelIndex", types.YLeaf{"MplsTunnelIndex", cmplsTunnelReversePerfEntry.MplsTunnelIndex})
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("mplsTunnelInstance", types.YLeaf{"MplsTunnelInstance", cmplsTunnelReversePerfEntry.MplsTunnelInstance})
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("mplsTunnelIngressLSRId", types.YLeaf{"MplsTunnelIngressLSRId", cmplsTunnelReversePerfEntry.MplsTunnelIngressLSRId})
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("mplsTunnelEgressLSRId", types.YLeaf{"MplsTunnelEgressLSRId", cmplsTunnelReversePerfEntry.MplsTunnelEgressLSRId})
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("cmplsTunnelReversePerfPackets", types.YLeaf{"CmplsTunnelReversePerfPackets", cmplsTunnelReversePerfEntry.CmplsTunnelReversePerfPackets})
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("cmplsTunnelReversePerfHCPackets", types.YLeaf{"CmplsTunnelReversePerfHCPackets", cmplsTunnelReversePerfEntry.CmplsTunnelReversePerfHCPackets})
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("cmplsTunnelReversePerfErrors", types.YLeaf{"CmplsTunnelReversePerfErrors", cmplsTunnelReversePerfEntry.CmplsTunnelReversePerfErrors})
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("cmplsTunnelReversePerfBytes", types.YLeaf{"CmplsTunnelReversePerfBytes", cmplsTunnelReversePerfEntry.CmplsTunnelReversePerfBytes})
    cmplsTunnelReversePerfEntry.EntityData.Leafs.Append("cmplsTunnelReversePerfHCBytes", types.YLeaf{"CmplsTunnelReversePerfHCBytes", cmplsTunnelReversePerfEntry.CmplsTunnelReversePerfHCBytes})

    cmplsTunnelReversePerfEntry.EntityData.YListKeys = []string {"MplsTunnelIndex", "MplsTunnelInstance", "MplsTunnelIngressLSRId", "MplsTunnelEgressLSRId"}

    return &(cmplsTunnelReversePerfEntry.EntityData)
}

