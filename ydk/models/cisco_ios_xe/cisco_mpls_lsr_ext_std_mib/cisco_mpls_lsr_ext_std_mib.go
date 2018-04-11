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
    Cmplsxcexttable CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable
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

    cISCOMPLSLSREXTSTDMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOMPLSLSREXTSTDMIB.EntityData.Children["cmplsXCExtTable"] = types.YChild{"Cmplsxcexttable", &cISCOMPLSLSREXTSTDMIB.Cmplsxcexttable}
    cISCOMPLSLSREXTSTDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOMPLSLSREXTSTDMIB.EntityData)
}

// CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable
// This table sparse augments the mplsXCTable of
// MPLS-LSR-STD-MIB [RFC3813] to provide MPLS-TP specific
// information about associated tunnel information
type CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in this table extends the cross connect information represented by
    // an entry in the mplsXCTable in MPLS-LSR-STD-MIB [RFC3813] through a sparse
    // augmentation.  An entry can be created by a network administrator via SNMP
    // SET commands, or in response to signaling protocol events. The type is
    // slice of CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry.
    Cmplsxcextentry []CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry
}

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetEntityData() *types.CommonEntityData {
    cmplsxcexttable.EntityData.YFilter = cmplsxcexttable.YFilter
    cmplsxcexttable.EntityData.YangName = "cmplsXCExtTable"
    cmplsxcexttable.EntityData.BundleName = "cisco_ios_xe"
    cmplsxcexttable.EntityData.ParentYangName = "CISCO-MPLS-LSR-EXT-STD-MIB"
    cmplsxcexttable.EntityData.SegmentPath = "cmplsXCExtTable"
    cmplsxcexttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsxcexttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsxcexttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsxcexttable.EntityData.Children = make(map[string]types.YChild)
    cmplsxcexttable.EntityData.Children["cmplsXCExtEntry"] = types.YChild{"Cmplsxcextentry", nil}
    for i := range cmplsxcexttable.Cmplsxcextentry {
        cmplsxcexttable.EntityData.Children[types.GetSegmentPath(&cmplsxcexttable.Cmplsxcextentry[i])] = types.YChild{"Cmplsxcextentry", &cmplsxcexttable.Cmplsxcextentry[i]}
    }
    cmplsxcexttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmplsxcexttable.EntityData)
}

// CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry
// An entry in this table extends the cross connect
// information represented by an entry in
// the mplsXCTable in MPLS-LSR-STD-MIB [RFC3813] through
// a sparse augmentation.  An entry can be created by
// a network administrator via SNMP SET commands, or in
// response to signaling protocol events.
type CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..24. Refers to
    // mpls_lsr_std_mib.MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcindex
    Mplsxcindex interface{}

    // This attribute is a key. The type is string with length: 1..24. Refers to
    // mpls_lsr_std_mib.MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcinsegmentindex
    Mplsxcinsegmentindex interface{}

    // This attribute is a key. The type is string with length: 1..24. Refers to
    // mpls_lsr_std_mib.MPLSLSRSTDMIB_Mplsxctable_Mplsxcentry_Mplsxcoutsegmentindex
    Mplsxcoutsegmentindex interface{}

    // This object indicates the back pointer to the tunnel entry segment.  This
    // object cannot be modified if mplsXCRowStatus for the corresponding entry in
    // the mplsXCTable is active(1). The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cmplsxcexttunnelpointer interface{}

    // This object indicates the pointer to the opposite direction XC entry.  This
    // object cannot be modified if mplsXCRowStatus for the corresponding entry in
    // the mplsXCTable is active(1). The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cmplsxcoppositedirxcptr interface{}
}

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetEntityData() *types.CommonEntityData {
    cmplsxcextentry.EntityData.YFilter = cmplsxcextentry.YFilter
    cmplsxcextentry.EntityData.YangName = "cmplsXCExtEntry"
    cmplsxcextentry.EntityData.BundleName = "cisco_ios_xe"
    cmplsxcextentry.EntityData.ParentYangName = "cmplsXCExtTable"
    cmplsxcextentry.EntityData.SegmentPath = "cmplsXCExtEntry" + "[mplsXCIndex='" + fmt.Sprintf("%v", cmplsxcextentry.Mplsxcindex) + "']" + "[mplsXCInSegmentIndex='" + fmt.Sprintf("%v", cmplsxcextentry.Mplsxcinsegmentindex) + "']" + "[mplsXCOutSegmentIndex='" + fmt.Sprintf("%v", cmplsxcextentry.Mplsxcoutsegmentindex) + "']"
    cmplsxcextentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsxcextentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsxcextentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsxcextentry.EntityData.Children = make(map[string]types.YChild)
    cmplsxcextentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cmplsxcextentry.EntityData.Leafs["mplsXCIndex"] = types.YLeaf{"Mplsxcindex", cmplsxcextentry.Mplsxcindex}
    cmplsxcextentry.EntityData.Leafs["mplsXCInSegmentIndex"] = types.YLeaf{"Mplsxcinsegmentindex", cmplsxcextentry.Mplsxcinsegmentindex}
    cmplsxcextentry.EntityData.Leafs["mplsXCOutSegmentIndex"] = types.YLeaf{"Mplsxcoutsegmentindex", cmplsxcextentry.Mplsxcoutsegmentindex}
    cmplsxcextentry.EntityData.Leafs["cmplsXCExtTunnelPointer"] = types.YLeaf{"Cmplsxcexttunnelpointer", cmplsxcextentry.Cmplsxcexttunnelpointer}
    cmplsxcextentry.EntityData.Leafs["cmplsXCOppositeDirXCPtr"] = types.YLeaf{"Cmplsxcoppositedirxcptr", cmplsxcextentry.Cmplsxcoppositedirxcptr}
    return &(cmplsxcextentry.EntityData)
}

