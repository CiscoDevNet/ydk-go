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
    parent types.Entity
    YFilter yfilter.YFilter

    // This table sparse augments the mplsXCTable of MPLS-LSR-STD-MIB [RFC3813] to
    // provide MPLS-TP specific information about associated tunnel information.
    Cmplsxcexttable CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable
}

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetFilter() yfilter.YFilter { return cISCOMPLSLSREXTSTDMIB.YFilter }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) SetFilter(yf yfilter.YFilter) { cISCOMPLSLSREXTSTDMIB.YFilter = yf }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetGoName(yname string) string {
    if yname == "cmplsXCExtTable" { return "Cmplsxcexttable" }
    return ""
}

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetSegmentPath() string {
    return "CISCO-MPLS-LSR-EXT-STD-MIB:CISCO-MPLS-LSR-EXT-STD-MIB"
}

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsXCExtTable" {
        return &cISCOMPLSLSREXTSTDMIB.Cmplsxcexttable
    }
    return nil
}

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cmplsXCExtTable"] = &cISCOMPLSLSREXTSTDMIB.Cmplsxcexttable
    return children
}

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetYangName() string { return "CISCO-MPLS-LSR-EXT-STD-MIB" }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) SetParent(parent types.Entity) { cISCOMPLSLSREXTSTDMIB.parent = parent }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetParent() types.Entity { return cISCOMPLSLSREXTSTDMIB.parent }

func (cISCOMPLSLSREXTSTDMIB *CISCOMPLSLSREXTSTDMIB) GetParentYangName() string { return "CISCO-MPLS-LSR-EXT-STD-MIB" }

// CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable
// This table sparse augments the mplsXCTable of
// MPLS-LSR-STD-MIB [RFC3813] to provide MPLS-TP specific
// information about associated tunnel information
type CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in this table extends the cross connect information represented by
    // an entry in the mplsXCTable in MPLS-LSR-STD-MIB [RFC3813] through a sparse
    // augmentation.  An entry can be created by a network administrator via SNMP
    // SET commands, or in response to signaling protocol events. The type is
    // slice of CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry.
    Cmplsxcextentry []CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry
}

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetFilter() yfilter.YFilter { return cmplsxcexttable.YFilter }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) SetFilter(yf yfilter.YFilter) { cmplsxcexttable.YFilter = yf }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetGoName(yname string) string {
    if yname == "cmplsXCExtEntry" { return "Cmplsxcextentry" }
    return ""
}

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetSegmentPath() string {
    return "cmplsXCExtTable"
}

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsXCExtEntry" {
        for _, c := range cmplsxcexttable.Cmplsxcextentry {
            if cmplsxcexttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry{}
        cmplsxcexttable.Cmplsxcextentry = append(cmplsxcexttable.Cmplsxcextentry, child)
        return &cmplsxcexttable.Cmplsxcextentry[len(cmplsxcexttable.Cmplsxcextentry)-1]
    }
    return nil
}

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cmplsxcexttable.Cmplsxcextentry {
        children[cmplsxcexttable.Cmplsxcextentry[i].GetSegmentPath()] = &cmplsxcexttable.Cmplsxcextentry[i]
    }
    return children
}

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetYangName() string { return "cmplsXCExtTable" }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) SetParent(parent types.Entity) { cmplsxcexttable.parent = parent }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetParent() types.Entity { return cmplsxcexttable.parent }

func (cmplsxcexttable *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable) GetParentYangName() string { return "CISCO-MPLS-LSR-EXT-STD-MIB" }

// CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry
// An entry in this table extends the cross connect
// information represented by an entry in
// the mplsXCTable in MPLS-LSR-STD-MIB [RFC3813] through
// a sparse augmentation.  An entry can be created by
// a network administrator via SNMP SET commands, or in
// response to signaling protocol events.
type CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Cmplsxcexttunnelpointer interface{}

    // This object indicates the pointer to the opposite direction XC entry.  This
    // object cannot be modified if mplsXCRowStatus for the corresponding entry in
    // the mplsXCTable is active(1). The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Cmplsxcoppositedirxcptr interface{}
}

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetFilter() yfilter.YFilter { return cmplsxcextentry.YFilter }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) SetFilter(yf yfilter.YFilter) { cmplsxcextentry.YFilter = yf }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetGoName(yname string) string {
    if yname == "mplsXCIndex" { return "Mplsxcindex" }
    if yname == "mplsXCInSegmentIndex" { return "Mplsxcinsegmentindex" }
    if yname == "mplsXCOutSegmentIndex" { return "Mplsxcoutsegmentindex" }
    if yname == "cmplsXCExtTunnelPointer" { return "Cmplsxcexttunnelpointer" }
    if yname == "cmplsXCOppositeDirXCPtr" { return "Cmplsxcoppositedirxcptr" }
    return ""
}

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetSegmentPath() string {
    return "cmplsXCExtEntry" + "[mplsXCIndex='" + fmt.Sprintf("%v", cmplsxcextentry.Mplsxcindex) + "']" + "[mplsXCInSegmentIndex='" + fmt.Sprintf("%v", cmplsxcextentry.Mplsxcinsegmentindex) + "']" + "[mplsXCOutSegmentIndex='" + fmt.Sprintf("%v", cmplsxcextentry.Mplsxcoutsegmentindex) + "']"
}

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mplsXCIndex"] = cmplsxcextentry.Mplsxcindex
    leafs["mplsXCInSegmentIndex"] = cmplsxcextentry.Mplsxcinsegmentindex
    leafs["mplsXCOutSegmentIndex"] = cmplsxcextentry.Mplsxcoutsegmentindex
    leafs["cmplsXCExtTunnelPointer"] = cmplsxcextentry.Cmplsxcexttunnelpointer
    leafs["cmplsXCOppositeDirXCPtr"] = cmplsxcextentry.Cmplsxcoppositedirxcptr
    return leafs
}

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetYangName() string { return "cmplsXCExtEntry" }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) SetParent(parent types.Entity) { cmplsxcextentry.parent = parent }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetParent() types.Entity { return cmplsxcextentry.parent }

func (cmplsxcextentry *CISCOMPLSLSREXTSTDMIB_Cmplsxcexttable_Cmplsxcextentry) GetParentYangName() string { return "cmplsXCExtTable" }

