// Copyright (c) 2012 IETF Trust and the persons identified
// as the document authors.  All rights reserved.
// 
// This MIB module contains generic object definitions for
// 
// 
// MPLS LSR in transport networks.
package cisco_mpls_lsr_ext_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_mpls_lsr_ext_std_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-MPLS-LSR-EXT-STD-MIB CISCO-MPLS-LSR-EXT-STD-MIB}", reflect.TypeOf(CISCOMPLSLSREXTSTDMIB{}))
    ydk.RegisterEntity("CISCO-MPLS-LSR-EXT-STD-MIB:CISCO-MPLS-LSR-EXT-STD-MIB", reflect.TypeOf(CISCOMPLSLSREXTSTDMIB{}))
}

// CISCOMPLSLSREXTSTDMIB
type CISCOMPLSLSREXTSTDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table sparse augments the mplsXCTable of MPLS-LSR-STD-MIB [RFC3813] to
    // provide MPLS-TP specific information about associated tunnel information.
    CmplsXCExtTable CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable
}

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetEntityData() *types.CommonEntityData {
    cISCOMPLSLSREXTSTDMIB.EntityData.YFilter = cISCOMPLSLSREXTSTDMIB.YFilter
    cISCOMPLSLSREXTSTDMIB.EntityData.YangName = "CISCO-MPLS-LSR-EXT-STD-MIB"
    cISCOMPLSLSREXTSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOMPLSLSREXTSTDMIB.EntityData.ParentYangName = "CISCO-MPLS-LSR-EXT-STD-MIB"
    cISCOMPLSLSREXTSTDMIB.EntityData.SegmentPath = "CISCO-MPLS-LSR-EXT-STD-MIB:CISCO-MPLS-LSR-EXT-STD-MIB"
    cISCOMPLSLSREXTSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOMPLSLSREXTSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOMPLSLSREXTSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOMPLSLSREXTSTDMIB.EntityData.Children = types.NewOrderedMap()
    cISCOMPLSLSREXTSTDMIB.EntityData.Children.Append("cmplsXCExtTable", types.YChild{"CmplsXCExtTable", &cISCOMPLSLSREXTSTDMIB.CmplsXCExtTable})
    cISCOMPLSLSREXTSTDMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOMPLSLSREXTSTDMIB.EntityData.YListKeys = []string {}

    return &(cISCOMPLSLSREXTSTDMIB.EntityData)
}

// CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable
// This table sparse augments the mplsXCTable of
// MPLS-LSR-STD-MIB [RFC3813] to provide MPLS-TP specific
// information about associated tunnel information
type CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table extends the cross connect information represented by
    // an entry in the mplsXCTable in MPLS-LSR-STD-MIB [RFC3813] through a sparse
    // augmentation.  An entry can be created by a network administrator via SNMP
    // SET commands, or in response to signaling protocol events. The type is
    // slice of CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable_CmplsXCExtEntry.
    CmplsXCExtEntry []*CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable_CmplsXCExtEntry
}

func (cmplsXCExtTable *CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable) GetEntityData() *types.CommonEntityData {
    cmplsXCExtTable.EntityData.YFilter = cmplsXCExtTable.YFilter
    cmplsXCExtTable.EntityData.YangName = "cmplsXCExtTable"
    cmplsXCExtTable.EntityData.BundleName = "cisco_ios_xe"
    cmplsXCExtTable.EntityData.ParentYangName = "CISCO-MPLS-LSR-EXT-STD-MIB"
    cmplsXCExtTable.EntityData.SegmentPath = "cmplsXCExtTable"
    cmplsXCExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsXCExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsXCExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsXCExtTable.EntityData.Children = types.NewOrderedMap()
    cmplsXCExtTable.EntityData.Children.Append("cmplsXCExtEntry", types.YChild{"CmplsXCExtEntry", nil})
    for i := range cmplsXCExtTable.CmplsXCExtEntry {
        cmplsXCExtTable.EntityData.Children.Append(types.GetSegmentPath(cmplsXCExtTable.CmplsXCExtEntry[i]), types.YChild{"CmplsXCExtEntry", cmplsXCExtTable.CmplsXCExtEntry[i]})
    }
    cmplsXCExtTable.EntityData.Leafs = types.NewOrderedMap()

    cmplsXCExtTable.EntityData.YListKeys = []string {}

    return &(cmplsXCExtTable.EntityData)
}

// CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable_CmplsXCExtEntry
// An entry in this table extends the cross connect
// information represented by an entry in
// the mplsXCTable in MPLS-LSR-STD-MIB [RFC3813] through
// a sparse augmentation.  An entry can be created by
// a network administrator via SNMP SET commands, or in
// response to signaling protocol events.
type CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable_CmplsXCExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..24. Refers to
    // mpls_lsr_std_mib.MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCIndex
    MplsXCIndex interface{}

    // This attribute is a key. The type is string with length: 1..24. Refers to
    // mpls_lsr_std_mib.MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCInSegmentIndex
    MplsXCInSegmentIndex interface{}

    // This attribute is a key. The type is string with length: 1..24. Refers to
    // mpls_lsr_std_mib.MPLSLSRSTDMIB_MplsXCTable_MplsXCEntry_MplsXCOutSegmentIndex
    MplsXCOutSegmentIndex interface{}

    // This object indicates the back pointer to the tunnel entry segment.  This
    // object cannot be modified if mplsXCRowStatus for the corresponding entry in
    // the mplsXCTable is active(1). The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CmplsXCExtTunnelPointer interface{}

    // This object indicates the pointer to the opposite direction XC entry.  This
    // object cannot be modified if mplsXCRowStatus for the corresponding entry in
    // the mplsXCTable is active(1). The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CmplsXCOppositeDirXCPtr interface{}
}

func (cmplsXCExtEntry *CISCOMPLSLSREXTSTDMIB_CmplsXCExtTable_CmplsXCExtEntry) GetEntityData() *types.CommonEntityData {
    cmplsXCExtEntry.EntityData.YFilter = cmplsXCExtEntry.YFilter
    cmplsXCExtEntry.EntityData.YangName = "cmplsXCExtEntry"
    cmplsXCExtEntry.EntityData.BundleName = "cisco_ios_xe"
    cmplsXCExtEntry.EntityData.ParentYangName = "cmplsXCExtTable"
    cmplsXCExtEntry.EntityData.SegmentPath = "cmplsXCExtEntry" + types.AddKeyToken(cmplsXCExtEntry.MplsXCIndex, "mplsXCIndex") + types.AddKeyToken(cmplsXCExtEntry.MplsXCInSegmentIndex, "mplsXCInSegmentIndex") + types.AddKeyToken(cmplsXCExtEntry.MplsXCOutSegmentIndex, "mplsXCOutSegmentIndex")
    cmplsXCExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsXCExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsXCExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsXCExtEntry.EntityData.Children = types.NewOrderedMap()
    cmplsXCExtEntry.EntityData.Leafs = types.NewOrderedMap()
    cmplsXCExtEntry.EntityData.Leafs.Append("mplsXCIndex", types.YLeaf{"MplsXCIndex", cmplsXCExtEntry.MplsXCIndex})
    cmplsXCExtEntry.EntityData.Leafs.Append("mplsXCInSegmentIndex", types.YLeaf{"MplsXCInSegmentIndex", cmplsXCExtEntry.MplsXCInSegmentIndex})
    cmplsXCExtEntry.EntityData.Leafs.Append("mplsXCOutSegmentIndex", types.YLeaf{"MplsXCOutSegmentIndex", cmplsXCExtEntry.MplsXCOutSegmentIndex})
    cmplsXCExtEntry.EntityData.Leafs.Append("cmplsXCExtTunnelPointer", types.YLeaf{"CmplsXCExtTunnelPointer", cmplsXCExtEntry.CmplsXCExtTunnelPointer})
    cmplsXCExtEntry.EntityData.Leafs.Append("cmplsXCOppositeDirXCPtr", types.YLeaf{"CmplsXCOppositeDirXCPtr", cmplsXCExtEntry.CmplsXCOppositeDirXCPtr})

    cmplsXCExtEntry.EntityData.YListKeys = []string {"MplsXCIndex", "MplsXCInSegmentIndex", "MplsXCOutSegmentIndex"}

    return &(cmplsXCExtEntry.EntityData)
}

