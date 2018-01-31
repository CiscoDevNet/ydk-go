// This MIB Module is a supplement to the
// CISCO-IETF-ATM2-PVCTRAP-MIB.
package cisco_ietf_atm2_pvctrap_mib_extn

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_atm2_pvctrap_mib_extn"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN}", reflect.TypeOf(CISCOIETFATM2PVCTRAPMIBEXTN{}))
    ydk.RegisterEntity("CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN:CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN", reflect.TypeOf(CISCOIETFATM2PVCTRAPMIBEXTN{}))
}

// CISCOIETFATM2PVCTRAPMIBEXTN
type CISCOIETFATM2PVCTRAPMIBEXTN struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table indicating all VCLs for which there is an active row in the
    // atmVclTable having an atmVclConnKind value of `pvc' and atmVclOperStatus to
    // have changed in the last PVC notification interval.
    Atmcurrentstatuschangepvcltable CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed in the same direction
    // in the last PVC notification interval.
    Atmstatuschangepvclrangetable CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable
}

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetFilter() yfilter.YFilter { return cISCOIETFATM2PVCTRAPMIBEXTN.YFilter }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) SetFilter(yf yfilter.YFilter) { cISCOIETFATM2PVCTRAPMIBEXTN.YFilter = yf }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetGoName(yname string) string {
    if yname == "atmCurrentStatusChangePVclTable" { return "Atmcurrentstatuschangepvcltable" }
    if yname == "atmStatusChangePVclRangeTable" { return "Atmstatuschangepvclrangetable" }
    return ""
}

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetSegmentPath() string {
    return "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN:CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN"
}

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmCurrentStatusChangePVclTable" {
        return &cISCOIETFATM2PVCTRAPMIBEXTN.Atmcurrentstatuschangepvcltable
    }
    if childYangName == "atmStatusChangePVclRangeTable" {
        return &cISCOIETFATM2PVCTRAPMIBEXTN.Atmstatuschangepvclrangetable
    }
    return nil
}

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["atmCurrentStatusChangePVclTable"] = &cISCOIETFATM2PVCTRAPMIBEXTN.Atmcurrentstatuschangepvcltable
    children["atmStatusChangePVclRangeTable"] = &cISCOIETFATM2PVCTRAPMIBEXTN.Atmstatuschangepvclrangetable
    return children
}

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetYangName() string { return "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN" }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) SetParent(parent types.Entity) { cISCOIETFATM2PVCTRAPMIBEXTN.parent = parent }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetParent() types.Entity { return cISCOIETFATM2PVCTRAPMIBEXTN.parent }

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetParentYangName() string { return "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN" }

// CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed in the
// last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed in the last PVC notification interval. The
    // type is slice of
    // CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry.
    Atmcurrentstatuschangepvclentry []CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry
}

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetFilter() yfilter.YFilter { return atmcurrentstatuschangepvcltable.YFilter }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) SetFilter(yf yfilter.YFilter) { atmcurrentstatuschangepvcltable.YFilter = yf }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetGoName(yname string) string {
    if yname == "atmCurrentStatusChangePVclEntry" { return "Atmcurrentstatuschangepvclentry" }
    return ""
}

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetSegmentPath() string {
    return "atmCurrentStatusChangePVclTable"
}

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmCurrentStatusChangePVclEntry" {
        for _, c := range atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry {
            if atmcurrentstatuschangepvcltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry{}
        atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry = append(atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry, child)
        return &atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry[len(atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry)-1]
    }
    return nil
}

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry {
        children[atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry[i].GetSegmentPath()] = &atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry[i]
    }
    return children
}

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetBundleName() string { return "cisco_ios_xe" }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetYangName() string { return "atmCurrentStatusChangePVclTable" }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) SetParent(parent types.Entity) { atmcurrentstatuschangepvcltable.parent = parent }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetParent() types.Entity { return atmcurrentstatuschangepvcltable.parent }

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetParentYangName() string { return "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN" }

// CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed in the last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvci
    Atmvclvci interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Atmpvclstatustransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Atmpvclstatuschangestart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Atmpvclstatuschangeend interface{}
}

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetFilter() yfilter.YFilter { return atmcurrentstatuschangepvclentry.YFilter }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) SetFilter(yf yfilter.YFilter) { atmcurrentstatuschangepvclentry.YFilter = yf }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "atmVclVci" { return "Atmvclvci" }
    if yname == "atmPVclStatusTransition" { return "Atmpvclstatustransition" }
    if yname == "atmPVclStatusChangeStart" { return "Atmpvclstatuschangestart" }
    if yname == "atmPVclStatusChangeEnd" { return "Atmpvclstatuschangeend" }
    return ""
}

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetSegmentPath() string {
    return "atmCurrentStatusChangePVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmcurrentstatuschangepvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", atmcurrentstatuschangepvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", atmcurrentstatuschangepvclentry.Atmvclvci) + "']"
}

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = atmcurrentstatuschangepvclentry.Ifindex
    leafs["atmVclVpi"] = atmcurrentstatuschangepvclentry.Atmvclvpi
    leafs["atmVclVci"] = atmcurrentstatuschangepvclentry.Atmvclvci
    leafs["atmPVclStatusTransition"] = atmcurrentstatuschangepvclentry.Atmpvclstatustransition
    leafs["atmPVclStatusChangeStart"] = atmcurrentstatuschangepvclentry.Atmpvclstatuschangestart
    leafs["atmPVclStatusChangeEnd"] = atmcurrentstatuschangepvclentry.Atmpvclstatuschangeend
    return leafs
}

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetBundleName() string { return "cisco_ios_xe" }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetYangName() string { return "atmCurrentStatusChangePVclEntry" }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) SetParent(parent types.Entity) { atmcurrentstatuschangepvclentry.parent = parent }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetParent() types.Entity { return atmcurrentstatuschangepvclentry.parent }

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetParentYangName() string { return "atmCurrentStatusChangePVclTable" }

// CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed in the same
// direction in the last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed in the same direction in the last
    // notification  interval. The type is slice of
    // CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry.
    Atmstatuschangepvclrangeentry []CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry
}

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetFilter() yfilter.YFilter { return atmstatuschangepvclrangetable.YFilter }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) SetFilter(yf yfilter.YFilter) { atmstatuschangepvclrangetable.YFilter = yf }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetGoName(yname string) string {
    if yname == "atmStatusChangePVclRangeEntry" { return "Atmstatuschangepvclrangeentry" }
    return ""
}

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetSegmentPath() string {
    return "atmStatusChangePVclRangeTable"
}

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmStatusChangePVclRangeEntry" {
        for _, c := range atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry {
            if atmstatuschangepvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry{}
        atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry = append(atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry, child)
        return &atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry[len(atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry)-1]
    }
    return nil
}

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry {
        children[atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry[i].GetSegmentPath()] = &atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry[i]
    }
    return children
}

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetYangName() string { return "atmStatusChangePVclRangeTable" }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) SetParent(parent types.Entity) { atmstatuschangepvclrangetable.parent = parent }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetParent() types.Entity { return atmstatuschangepvclrangetable.parent }

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetParentYangName() string { return "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN" }

// CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed in the same direction in the last notification 
// interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. Index to represent different ranges, the first
    // range is  indexed by value 0, the second by 1 and so on... The type is
    // interface{} with range: 0..4294967295.
    Rangeindex interface{}

    // The first PVCL in a range of PVcls for which the  atmVclOperStatus to have
    // changed in the last  notification interval. The type is interface{} with
    // range: 0..65536.
    Atmpvcllowerrangevalue interface{}

    // The last PVCL in a range of PVcls for which the  atmOperStatus to have
    // changed in the last  notification interval. The type is interface{} with
    // range: 0..65536.
    Atmpvclhigherrangevalue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Atmpvclrangestatuschangestart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Atmpvclrangestatuschangeend interface{}
}

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetFilter() yfilter.YFilter { return atmstatuschangepvclrangeentry.YFilter }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) SetFilter(yf yfilter.YFilter) { atmstatuschangepvclrangeentry.YFilter = yf }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "rangeIndex" { return "Rangeindex" }
    if yname == "atmPVclLowerRangeValue" { return "Atmpvcllowerrangevalue" }
    if yname == "atmPVclHigherRangeValue" { return "Atmpvclhigherrangevalue" }
    if yname == "atmPVclRangeStatusChangeStart" { return "Atmpvclrangestatuschangestart" }
    if yname == "atmPVclRangeStatusChangeEnd" { return "Atmpvclrangestatuschangeend" }
    return ""
}

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetSegmentPath() string {
    return "atmStatusChangePVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmstatuschangepvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", atmstatuschangepvclrangeentry.Atmvclvpi) + "']" + "[rangeIndex='" + fmt.Sprintf("%v", atmstatuschangepvclrangeentry.Rangeindex) + "']"
}

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = atmstatuschangepvclrangeentry.Ifindex
    leafs["atmVclVpi"] = atmstatuschangepvclrangeentry.Atmvclvpi
    leafs["rangeIndex"] = atmstatuschangepvclrangeentry.Rangeindex
    leafs["atmPVclLowerRangeValue"] = atmstatuschangepvclrangeentry.Atmpvcllowerrangevalue
    leafs["atmPVclHigherRangeValue"] = atmstatuschangepvclrangeentry.Atmpvclhigherrangevalue
    leafs["atmPVclRangeStatusChangeStart"] = atmstatuschangepvclrangeentry.Atmpvclrangestatuschangestart
    leafs["atmPVclRangeStatusChangeEnd"] = atmstatuschangepvclrangeentry.Atmpvclrangestatuschangeend
    return leafs
}

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetYangName() string { return "atmStatusChangePVclRangeEntry" }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) SetParent(parent types.Entity) { atmstatuschangepvclrangeentry.parent = parent }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetParent() types.Entity { return atmstatuschangepvclrangeentry.parent }

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetParentYangName() string { return "atmStatusChangePVclRangeTable" }

