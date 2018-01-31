// This MIB Module is a supplement to the
// CISCO-IETF-ATM2-PVCTRAP-MIB.
package cisco_atm_pvctrap_extn_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_atm_pvctrap_extn_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ATM-PVCTRAP-EXTN-MIB CISCO-ATM-PVCTRAP-EXTN-MIB}", reflect.TypeOf(CISCOATMPVCTRAPEXTNMIB{}))
    ydk.RegisterEntity("CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB", reflect.TypeOf(CISCOATMPVCTRAPEXTNMIB{}))
}

// CatmOAMRecoveryType represents                           type of recovery.
type CatmOAMRecoveryType string

const (
    CatmOAMRecoveryType_catmLoopbackOAMRecover CatmOAMRecoveryType = "catmLoopbackOAMRecover"

    CatmOAMRecoveryType_catmSegmentCCOAMRecover CatmOAMRecoveryType = "catmSegmentCCOAMRecover"

    CatmOAMRecoveryType_catmEndCCOAMRecover CatmOAMRecoveryType = "catmEndCCOAMRecover"

    CatmOAMRecoveryType_catmAISRDIOAMRecover CatmOAMRecoveryType = "catmAISRDIOAMRecover"
)

// CatmOAMFailureType represents                           type of failure.
type CatmOAMFailureType string

const (
    CatmOAMFailureType_catmLoopbackOAMFailure CatmOAMFailureType = "catmLoopbackOAMFailure"

    CatmOAMFailureType_catmSegmentCCOAMFailure CatmOAMFailureType = "catmSegmentCCOAMFailure"

    CatmOAMFailureType_catmEndCCOAMFailure CatmOAMFailureType = "catmEndCCOAMFailure"

    CatmOAMFailureType_catmAISRDIOAMFailure CatmOAMFailureType = "catmAISRDIOAMFailure"
)

// CISCOATMPVCTRAPEXTNMIB
type CISCOATMPVCTRAPEXTNMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table indicating all VCLs for which there is an active row in the
    // atmVclTable having an atmVclConnKind value of `pvc' and atmVclOperStatus to
    // have changed in the last corresponding PVC notification.
    Catmcurstatchangepvcltable CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed in the same direction
    // in the last corresponding PVC notification .
    Catmstatuschangepvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed due to segment CC 
    // failure in the same direction in the last PVC  corresponding notification .
    Catmsegccstatuschpvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed due to End CC failure
    // in the same direction in the last PVC notification  interval.
    Catmendccstatuschpvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed due to AIS/RDI failure
    // in the same direction in the last corresponding PVC  notification.
    Catmaisrdistatuschpvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have detected as Down in the last
    // corresponding PVC notification .
    Catmdownpvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable

    // A table indicating all VCLs for which there is an active row in the
    // atmVclTable having an atmVclConnKind value of `pvc' and atmVclOperStatus to
    // have changed to UP in the last corresponding PVC notification .
    Catmcurstatusuppvcltable CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and loopback OAM status to have detected as recovered in the
    // last corresponding PVC notification .
    Catmstatusuppvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable

    // A table indicating more than one VCLs in a consecutive range and for each
    // VCL there is an active row in the atmVclTable having an atmVclConnKind
    // value of `pvc' and Segment CC OAM status to have detected as recovered in
    // the last corresponding PVC notification .
    Catmsegccstatusuppvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable

    // A table indicating more than one VCLs in a consecutive range and for each
    // VCL there is an active row in the atmVclTable having an atmVclConnKind
    // value of `pvc' and End-to-End CC OAM status to have detected as recovered
    // in the last corresponding PVC notification .
    Catmendccstatusuppvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable

    // A table indicating more than one VCLs in a consecutive range and for each
    // VCL there is an active row in the atmVclTable having an atmVclConnKind
    // value of `pvc' and AISRDI OAM status to have detected as recovered in the
    // last corresponding PVC notification .
    Catmaisrdistatusuppvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have detected as Up in the last
    // corresponding PVC notification .
    Catmuppvclrangetable CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable
}

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetFilter() yfilter.YFilter { return cISCOATMPVCTRAPEXTNMIB.YFilter }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) SetFilter(yf yfilter.YFilter) { cISCOATMPVCTRAPEXTNMIB.YFilter = yf }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetGoName(yname string) string {
    if yname == "catmCurStatChangePVclTable" { return "Catmcurstatchangepvcltable" }
    if yname == "catmStatusChangePVclRangeTable" { return "Catmstatuschangepvclrangetable" }
    if yname == "catmSegCCStatusChPVclRangeTable" { return "Catmsegccstatuschpvclrangetable" }
    if yname == "catmEndCCStatusChPVclRangeTable" { return "Catmendccstatuschpvclrangetable" }
    if yname == "catmAISRDIStatusChPVclRangeTable" { return "Catmaisrdistatuschpvclrangetable" }
    if yname == "catmDownPVclRangeTable" { return "Catmdownpvclrangetable" }
    if yname == "catmCurStatusUpPVclTable" { return "Catmcurstatusuppvcltable" }
    if yname == "catmStatusUpPVclRangeTable" { return "Catmstatusuppvclrangetable" }
    if yname == "catmSegCCStatusUpPVclRangeTable" { return "Catmsegccstatusuppvclrangetable" }
    if yname == "catmEndCCStatusUpPVclRangeTable" { return "Catmendccstatusuppvclrangetable" }
    if yname == "catmAISRDIStatusUpPVclRangeTable" { return "Catmaisrdistatusuppvclrangetable" }
    if yname == "catmUpPVclRangeTable" { return "Catmuppvclrangetable" }
    return ""
}

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetSegmentPath() string {
    return "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB"
}

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmCurStatChangePVclTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmcurstatchangepvcltable
    }
    if childYangName == "catmStatusChangePVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmstatuschangepvclrangetable
    }
    if childYangName == "catmSegCCStatusChPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmsegccstatuschpvclrangetable
    }
    if childYangName == "catmEndCCStatusChPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmendccstatuschpvclrangetable
    }
    if childYangName == "catmAISRDIStatusChPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmaisrdistatuschpvclrangetable
    }
    if childYangName == "catmDownPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmdownpvclrangetable
    }
    if childYangName == "catmCurStatusUpPVclTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmcurstatusuppvcltable
    }
    if childYangName == "catmStatusUpPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmstatusuppvclrangetable
    }
    if childYangName == "catmSegCCStatusUpPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmsegccstatusuppvclrangetable
    }
    if childYangName == "catmEndCCStatusUpPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmendccstatusuppvclrangetable
    }
    if childYangName == "catmAISRDIStatusUpPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmaisrdistatusuppvclrangetable
    }
    if childYangName == "catmUpPVclRangeTable" {
        return &cISCOATMPVCTRAPEXTNMIB.Catmuppvclrangetable
    }
    return nil
}

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["catmCurStatChangePVclTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmcurstatchangepvcltable
    children["catmStatusChangePVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmstatuschangepvclrangetable
    children["catmSegCCStatusChPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmsegccstatuschpvclrangetable
    children["catmEndCCStatusChPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmendccstatuschpvclrangetable
    children["catmAISRDIStatusChPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmaisrdistatuschpvclrangetable
    children["catmDownPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmdownpvclrangetable
    children["catmCurStatusUpPVclTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmcurstatusuppvcltable
    children["catmStatusUpPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmstatusuppvclrangetable
    children["catmSegCCStatusUpPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmsegccstatusuppvclrangetable
    children["catmEndCCStatusUpPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmendccstatusuppvclrangetable
    children["catmAISRDIStatusUpPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmaisrdistatusuppvclrangetable
    children["catmUpPVclRangeTable"] = &cISCOATMPVCTRAPEXTNMIB.Catmuppvclrangetable
    return children
}

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) SetParent(parent types.Entity) { cISCOATMPVCTRAPEXTNMIB.parent = parent }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetParent() types.Entity { return cISCOATMPVCTRAPEXTNMIB.parent }

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed in the
// last corresponding PVC notification.
type CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed in the last corresponding PVC
    // notification. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry.
    Catmcurstatchangepvclentry []CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry
}

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetFilter() yfilter.YFilter { return catmcurstatchangepvcltable.YFilter }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) SetFilter(yf yfilter.YFilter) { catmcurstatchangepvcltable.YFilter = yf }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetGoName(yname string) string {
    if yname == "catmCurStatChangePVclEntry" { return "Catmcurstatchangepvclentry" }
    return ""
}

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetSegmentPath() string {
    return "catmCurStatChangePVclTable"
}

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmCurStatChangePVclEntry" {
        for _, c := range catmcurstatchangepvcltable.Catmcurstatchangepvclentry {
            if catmcurstatchangepvcltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry{}
        catmcurstatchangepvcltable.Catmcurstatchangepvclentry = append(catmcurstatchangepvcltable.Catmcurstatchangepvclentry, child)
        return &catmcurstatchangepvcltable.Catmcurstatchangepvclentry[len(catmcurstatchangepvcltable.Catmcurstatchangepvclentry)-1]
    }
    return nil
}

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmcurstatchangepvcltable.Catmcurstatchangepvclentry {
        children[catmcurstatchangepvcltable.Catmcurstatchangepvclentry[i].GetSegmentPath()] = &catmcurstatchangepvcltable.Catmcurstatchangepvclentry[i]
    }
    return children
}

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetBundleName() string { return "cisco_ios_xe" }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetYangName() string { return "catmCurStatChangePVclTable" }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) SetParent(parent types.Entity) { catmcurstatchangepvcltable.parent = parent }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetParent() types.Entity { return catmcurstatchangepvcltable.parent }

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed in the last corresponding PVC notification.
type CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry struct {
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
    // corresponding notification. The type is interface{} with range:
    // 0..4294967295.
    Catmpvclstatustransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last corresponding notification. The type is interface{} with range:
    // 0..4294967295.
    Catmpvclstatuschangestart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification. The type is interface{} with range:
    // 0..4294967295.
    Catmpvclstatuschangeend interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // corresponding notification due to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclsegccstatustransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last corresponding notification due to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclsegccstatuschangestart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification due to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclsegccstatuschangeend interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // corresponding notification due to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclendccstatustransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last corresponding notification due to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclendccstatuschangestart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification due to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclendccstatuschangeend interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // corresponding notification due to AIS RDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclaisrdistatustransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last corresponding notification due to AIS RDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclaisrdistatuschangestart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification due to AIS RDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclaisrdistatuschangeend interface{}

    // The time stamp at which this PVCL changed state to DOWN last time in the
    // last corresponding notification . The type is interface{} with range:
    // 0..4294967295.
    Catmpvclcurfailtime interface{}

    // The time stamp at which this PVCL changed state to UP last time in the
    // previous corresponding notification . The type is interface{} with range:
    // 0..4294967295.
    Catmpvclprevrecovertime interface{}

    // Type of OAM failure. The type is CatmOAMFailureType.
    Catmpvclfailurereason interface{}
}

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetFilter() yfilter.YFilter { return catmcurstatchangepvclentry.YFilter }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) SetFilter(yf yfilter.YFilter) { catmcurstatchangepvclentry.YFilter = yf }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "atmVclVci" { return "Atmvclvci" }
    if yname == "catmPVclStatusTransition" { return "Catmpvclstatustransition" }
    if yname == "catmPVclStatusChangeStart" { return "Catmpvclstatuschangestart" }
    if yname == "catmPVclStatusChangeEnd" { return "Catmpvclstatuschangeend" }
    if yname == "catmPVclSegCCStatusTransition" { return "Catmpvclsegccstatustransition" }
    if yname == "catmPVclSegCCStatusChangeStart" { return "Catmpvclsegccstatuschangestart" }
    if yname == "catmPVclSegCCStatusChangeEnd" { return "Catmpvclsegccstatuschangeend" }
    if yname == "catmPVclEndCCStatusTransition" { return "Catmpvclendccstatustransition" }
    if yname == "catmPVclEndCCStatusChangeStart" { return "Catmpvclendccstatuschangestart" }
    if yname == "catmPVclEndCCStatusChangeEnd" { return "Catmpvclendccstatuschangeend" }
    if yname == "catmPVclAISRDIStatusTransition" { return "Catmpvclaisrdistatustransition" }
    if yname == "catmPVclAISRDIStatusChangeStart" { return "Catmpvclaisrdistatuschangestart" }
    if yname == "catmPVclAISRDIStatusChangeEnd" { return "Catmpvclaisrdistatuschangeend" }
    if yname == "catmPVclCurFailTime" { return "Catmpvclcurfailtime" }
    if yname == "catmPVclPrevRecoverTime" { return "Catmpvclprevrecovertime" }
    if yname == "catmPVclFailureReason" { return "Catmpvclfailurereason" }
    return ""
}

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetSegmentPath() string {
    return "catmCurStatChangePVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmcurstatchangepvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmcurstatchangepvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", catmcurstatchangepvclentry.Atmvclvci) + "']"
}

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmcurstatchangepvclentry.Ifindex
    leafs["atmVclVpi"] = catmcurstatchangepvclentry.Atmvclvpi
    leafs["atmVclVci"] = catmcurstatchangepvclentry.Atmvclvci
    leafs["catmPVclStatusTransition"] = catmcurstatchangepvclentry.Catmpvclstatustransition
    leafs["catmPVclStatusChangeStart"] = catmcurstatchangepvclentry.Catmpvclstatuschangestart
    leafs["catmPVclStatusChangeEnd"] = catmcurstatchangepvclentry.Catmpvclstatuschangeend
    leafs["catmPVclSegCCStatusTransition"] = catmcurstatchangepvclentry.Catmpvclsegccstatustransition
    leafs["catmPVclSegCCStatusChangeStart"] = catmcurstatchangepvclentry.Catmpvclsegccstatuschangestart
    leafs["catmPVclSegCCStatusChangeEnd"] = catmcurstatchangepvclentry.Catmpvclsegccstatuschangeend
    leafs["catmPVclEndCCStatusTransition"] = catmcurstatchangepvclentry.Catmpvclendccstatustransition
    leafs["catmPVclEndCCStatusChangeStart"] = catmcurstatchangepvclentry.Catmpvclendccstatuschangestart
    leafs["catmPVclEndCCStatusChangeEnd"] = catmcurstatchangepvclentry.Catmpvclendccstatuschangeend
    leafs["catmPVclAISRDIStatusTransition"] = catmcurstatchangepvclentry.Catmpvclaisrdistatustransition
    leafs["catmPVclAISRDIStatusChangeStart"] = catmcurstatchangepvclentry.Catmpvclaisrdistatuschangestart
    leafs["catmPVclAISRDIStatusChangeEnd"] = catmcurstatchangepvclentry.Catmpvclaisrdistatuschangeend
    leafs["catmPVclCurFailTime"] = catmcurstatchangepvclentry.Catmpvclcurfailtime
    leafs["catmPVclPrevRecoverTime"] = catmcurstatchangepvclentry.Catmpvclprevrecovertime
    leafs["catmPVclFailureReason"] = catmcurstatchangepvclentry.Catmpvclfailurereason
    return leafs
}

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetYangName() string { return "catmCurStatChangePVclEntry" }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) SetParent(parent types.Entity) { catmcurstatchangepvclentry.parent = parent }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetParent() types.Entity { return catmcurstatchangepvclentry.parent }

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetParentYangName() string { return "catmCurStatChangePVclTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed in the same
// direction in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed in the same direction in the last
    // notification  interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry.
    Catmstatuschangepvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry
}

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetFilter() yfilter.YFilter { return catmstatuschangepvclrangetable.YFilter }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) SetFilter(yf yfilter.YFilter) { catmstatuschangepvclrangetable.YFilter = yf }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetGoName(yname string) string {
    if yname == "catmStatusChangePVclRangeEntry" { return "Catmstatuschangepvclrangeentry" }
    return ""
}

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetSegmentPath() string {
    return "catmStatusChangePVclRangeTable"
}

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmStatusChangePVclRangeEntry" {
        for _, c := range catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry {
            if catmstatuschangepvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry{}
        catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry = append(catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry, child)
        return &catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry[len(catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry)-1]
    }
    return nil
}

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry {
        children[catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry[i].GetSegmentPath()] = &catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry[i]
    }
    return children
}

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetYangName() string { return "catmStatusChangePVclRangeTable" }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) SetParent(parent types.Entity) { catmstatuschangepvclrangetable.parent = parent }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetParent() types.Entity { return catmstatuschangepvclrangetable.parent }

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed in the same direction in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. Index to represent different ranges, the first
    // range is indexed by value 0, the second by 1 and so on... The type is
    // interface{} with range: 0..65535.
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus to have
    // changed in the last  corresponding notification due to Loopback OAM
    // failure. The type is interface{} with range: 0..65536.
    Catmpvcllowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the  atmOperStatus to have
    // changed in the last  corresponding notification due to Loopback OAM
    // failure. The type is interface{} with range: 0..65536.
    Catmpvclhigherrangevalue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last corresponding notification due  to Loopback OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclrangestatuschangestart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last corresponding notification due  to Loopback OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclrangestatuschangeend interface{}
}

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetFilter() yfilter.YFilter { return catmstatuschangepvclrangeentry.YFilter }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) SetFilter(yf yfilter.YFilter) { catmstatuschangepvclrangeentry.YFilter = yf }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmPVclLowerRangeValue" { return "Catmpvcllowerrangevalue" }
    if yname == "catmPVclHigherRangeValue" { return "Catmpvclhigherrangevalue" }
    if yname == "catmPVclRangeStatusChangeStart" { return "Catmpvclrangestatuschangestart" }
    if yname == "catmPVclRangeStatusChangeEnd" { return "Catmpvclrangestatuschangeend" }
    return ""
}

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetSegmentPath() string {
    return "catmStatusChangePVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmstatuschangepvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmstatuschangepvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmstatuschangepvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmstatuschangepvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmstatuschangepvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmstatuschangepvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmPVclLowerRangeValue"] = catmstatuschangepvclrangeentry.Catmpvcllowerrangevalue
    leafs["catmPVclHigherRangeValue"] = catmstatuschangepvclrangeentry.Catmpvclhigherrangevalue
    leafs["catmPVclRangeStatusChangeStart"] = catmstatuschangepvclrangeentry.Catmpvclrangestatuschangestart
    leafs["catmPVclRangeStatusChangeEnd"] = catmstatuschangepvclrangeentry.Catmpvclrangestatuschangeend
    return leafs
}

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetYangName() string { return "catmStatusChangePVclRangeEntry" }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) SetParent(parent types.Entity) { catmstatuschangepvclrangeentry.parent = parent }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetParent() types.Entity { return catmstatuschangepvclrangeentry.parent }

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetParentYangName() string { return "catmStatusChangePVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to segment CC 
// failure in the same direction in the last PVC 
// corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to segment CC failure in the same
    // direction  in the last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry.
    Catmsegccstatuschpvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry
}

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetFilter() yfilter.YFilter { return catmsegccstatuschpvclrangetable.YFilter }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) SetFilter(yf yfilter.YFilter) { catmsegccstatuschpvclrangetable.YFilter = yf }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetGoName(yname string) string {
    if yname == "catmSegCCStatusChPVclRangeEntry" { return "Catmsegccstatuschpvclrangeentry" }
    return ""
}

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetSegmentPath() string {
    return "catmSegCCStatusChPVclRangeTable"
}

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmSegCCStatusChPVclRangeEntry" {
        for _, c := range catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry {
            if catmsegccstatuschpvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry{}
        catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry = append(catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry, child)
        return &catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry[len(catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry)-1]
    }
    return nil
}

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry {
        children[catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry[i].GetSegmentPath()] = &catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry[i]
    }
    return children
}

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetYangName() string { return "catmSegCCStatusChPVclRangeTable" }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) SetParent(parent types.Entity) { catmsegccstatuschpvclrangetable.parent = parent }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetParent() types.Entity { return catmsegccstatuschpvclrangetable.parent }

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to segment CC failure in the same direction 
// in the last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus to have
    // changed in the last  corresponding notification due to Segment CC OAM
    // failure. The type is interface{} with range: 0..65536.
    Catmpvclsegcclowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the  atmOperStatus to have
    // changed in the last  corresponding notification due to Segment CC OAM
    // failure. The type is interface{} with range: 0..65536.
    Catmpvclsegcchigherrangevalue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last corresponding notification due  to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclsegccrangestatuschstart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last corresponding notification due  to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclsegccrangestatuschend interface{}
}

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetFilter() yfilter.YFilter { return catmsegccstatuschpvclrangeentry.YFilter }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) SetFilter(yf yfilter.YFilter) { catmsegccstatuschpvclrangeentry.YFilter = yf }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmPVclSegCCLowerRangeValue" { return "Catmpvclsegcclowerrangevalue" }
    if yname == "catmPVclSegCCHigherRangeValue" { return "Catmpvclsegcchigherrangevalue" }
    if yname == "catmPVclSegCCRangeStatusChStart" { return "Catmpvclsegccrangestatuschstart" }
    if yname == "catmPVclSegCCRangeStatusChEnd" { return "Catmpvclsegccrangestatuschend" }
    return ""
}

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetSegmentPath() string {
    return "catmSegCCStatusChPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmsegccstatuschpvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmsegccstatuschpvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmsegccstatuschpvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmsegccstatuschpvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmsegccstatuschpvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmsegccstatuschpvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmPVclSegCCLowerRangeValue"] = catmsegccstatuschpvclrangeentry.Catmpvclsegcclowerrangevalue
    leafs["catmPVclSegCCHigherRangeValue"] = catmsegccstatuschpvclrangeentry.Catmpvclsegcchigherrangevalue
    leafs["catmPVclSegCCRangeStatusChStart"] = catmsegccstatuschpvclrangeentry.Catmpvclsegccrangestatuschstart
    leafs["catmPVclSegCCRangeStatusChEnd"] = catmsegccstatuschpvclrangeentry.Catmpvclsegccrangestatuschend
    return leafs
}

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetYangName() string { return "catmSegCCStatusChPVclRangeEntry" }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) SetParent(parent types.Entity) { catmsegccstatuschpvclrangeentry.parent = parent }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetParent() types.Entity { return catmsegccstatuschpvclrangeentry.parent }

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetParentYangName() string { return "catmSegCCStatusChPVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to End CC failure
// in the same direction in the last PVC notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to End CC failure in the same
    // direction in the  last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry.
    Catmendccstatuschpvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry
}

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetFilter() yfilter.YFilter { return catmendccstatuschpvclrangetable.YFilter }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) SetFilter(yf yfilter.YFilter) { catmendccstatuschpvclrangetable.YFilter = yf }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetGoName(yname string) string {
    if yname == "catmEndCCStatusChPVclRangeEntry" { return "Catmendccstatuschpvclrangeentry" }
    return ""
}

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetSegmentPath() string {
    return "catmEndCCStatusChPVclRangeTable"
}

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmEndCCStatusChPVclRangeEntry" {
        for _, c := range catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry {
            if catmendccstatuschpvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry{}
        catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry = append(catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry, child)
        return &catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry[len(catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry)-1]
    }
    return nil
}

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry {
        children[catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry[i].GetSegmentPath()] = &catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry[i]
    }
    return children
}

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetYangName() string { return "catmEndCCStatusChPVclRangeTable" }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) SetParent(parent types.Entity) { catmendccstatuschpvclrangetable.parent = parent }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetParent() types.Entity { return catmendccstatuschpvclrangetable.parent }

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to End CC failure in the same direction in the 
// last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus to have
    // changed in the last  corresponding notification due to End CC OAM failure.
    // The type is interface{} with range: 0..65536.
    Catmpvclendcclowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the  atmOperStatus to have
    // changed in the last  corresponding notification due to End CC OAM failure.
    // The type is interface{} with range: 0..65536.
    Catmpvclendcchigherrangevalue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last corresponding notification due  to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclendccrangestatuschstart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last corresponding notification due  to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclendccrangestatuschend interface{}
}

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetFilter() yfilter.YFilter { return catmendccstatuschpvclrangeentry.YFilter }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) SetFilter(yf yfilter.YFilter) { catmendccstatuschpvclrangeentry.YFilter = yf }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmPVclEndCCLowerRangeValue" { return "Catmpvclendcclowerrangevalue" }
    if yname == "catmPVclEndCCHigherRangeValue" { return "Catmpvclendcchigherrangevalue" }
    if yname == "catmPVclEndCCRangeStatusChStart" { return "Catmpvclendccrangestatuschstart" }
    if yname == "catmPVclEndCCRangeStatusChEnd" { return "Catmpvclendccrangestatuschend" }
    return ""
}

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetSegmentPath() string {
    return "catmEndCCStatusChPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmendccstatuschpvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmendccstatuschpvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmendccstatuschpvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmendccstatuschpvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmendccstatuschpvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmendccstatuschpvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmPVclEndCCLowerRangeValue"] = catmendccstatuschpvclrangeentry.Catmpvclendcclowerrangevalue
    leafs["catmPVclEndCCHigherRangeValue"] = catmendccstatuschpvclrangeentry.Catmpvclendcchigherrangevalue
    leafs["catmPVclEndCCRangeStatusChStart"] = catmendccstatuschpvclrangeentry.Catmpvclendccrangestatuschstart
    leafs["catmPVclEndCCRangeStatusChEnd"] = catmendccstatuschpvclrangeentry.Catmpvclendccrangestatuschend
    return leafs
}

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetYangName() string { return "catmEndCCStatusChPVclRangeEntry" }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) SetParent(parent types.Entity) { catmendccstatuschpvclrangeentry.parent = parent }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetParent() types.Entity { return catmendccstatuschpvclrangeentry.parent }

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetParentYangName() string { return "catmEndCCStatusChPVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to AIS/RDI failure
// in the same direction in the last corresponding PVC 
// notification.
type CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to AIS/RDI failure in the same
    // direction in the  last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry.
    Catmaisrdistatuschpvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry
}

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetFilter() yfilter.YFilter { return catmaisrdistatuschpvclrangetable.YFilter }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) SetFilter(yf yfilter.YFilter) { catmaisrdistatuschpvclrangetable.YFilter = yf }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetGoName(yname string) string {
    if yname == "catmAISRDIStatusChPVclRangeEntry" { return "Catmaisrdistatuschpvclrangeentry" }
    return ""
}

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetSegmentPath() string {
    return "catmAISRDIStatusChPVclRangeTable"
}

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmAISRDIStatusChPVclRangeEntry" {
        for _, c := range catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry {
            if catmaisrdistatuschpvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry{}
        catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry = append(catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry, child)
        return &catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry[len(catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry)-1]
    }
    return nil
}

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry {
        children[catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry[i].GetSegmentPath()] = &catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry[i]
    }
    return children
}

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetYangName() string { return "catmAISRDIStatusChPVclRangeTable" }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) SetParent(parent types.Entity) { catmaisrdistatuschpvclrangetable.parent = parent }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetParent() types.Entity { return catmaisrdistatuschpvclrangetable.parent }

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to AIS/RDI failure in the same direction in the 
// last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus to have
    // changed in the last  corresponding notification due to AISRDI OAM failure.
    // The type is interface{} with range: 0..65536.
    Catmpvclaisrdilowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the  atmOperStatus to have
    // changed in the last  corresponding notification due to AISRDI OAM failure.
    // The type is interface{} with range: 0..65536.
    Catmpvclaisrdihigherrangevalue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last corresponding notification due  to AISRDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclaisrdirangestatuschstart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last corresponding notification due  to AISRDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclaisrdirangestatuschend interface{}
}

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetFilter() yfilter.YFilter { return catmaisrdistatuschpvclrangeentry.YFilter }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) SetFilter(yf yfilter.YFilter) { catmaisrdistatuschpvclrangeentry.YFilter = yf }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmPVclAISRDILowerRangeValue" { return "Catmpvclaisrdilowerrangevalue" }
    if yname == "catmPVclAISRDIHigherRangeValue" { return "Catmpvclaisrdihigherrangevalue" }
    if yname == "catmPVclAISRDIRangeStatusChStart" { return "Catmpvclaisrdirangestatuschstart" }
    if yname == "catmPVclAISRDIRangeStatusChEnd" { return "Catmpvclaisrdirangestatuschend" }
    return ""
}

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetSegmentPath() string {
    return "catmAISRDIStatusChPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmaisrdistatuschpvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmaisrdistatuschpvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmaisrdistatuschpvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmaisrdistatuschpvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmaisrdistatuschpvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmaisrdistatuschpvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmPVclAISRDILowerRangeValue"] = catmaisrdistatuschpvclrangeentry.Catmpvclaisrdilowerrangevalue
    leafs["catmPVclAISRDIHigherRangeValue"] = catmaisrdistatuschpvclrangeentry.Catmpvclaisrdihigherrangevalue
    leafs["catmPVclAISRDIRangeStatusChStart"] = catmaisrdistatuschpvclrangeentry.Catmpvclaisrdirangestatuschstart
    leafs["catmPVclAISRDIRangeStatusChEnd"] = catmaisrdistatuschpvclrangeentry.Catmpvclaisrdirangestatuschend
    return leafs
}

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetYangName() string { return "catmAISRDIStatusChPVclRangeEntry" }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) SetParent(parent types.Entity) { catmaisrdistatuschpvclrangeentry.parent = parent }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetParent() types.Entity { return catmaisrdistatuschpvclrangeentry.parent }

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetParentYangName() string { return "catmAISRDIStatusChPVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have detected as Down
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  atmVclOperStatus to  have detected as Down in the last notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry.
    Catmdownpvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry
}

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetFilter() yfilter.YFilter { return catmdownpvclrangetable.YFilter }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) SetFilter(yf yfilter.YFilter) { catmdownpvclrangetable.YFilter = yf }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetGoName(yname string) string {
    if yname == "catmDownPVclRangeEntry" { return "Catmdownpvclrangeentry" }
    return ""
}

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetSegmentPath() string {
    return "catmDownPVclRangeTable"
}

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmDownPVclRangeEntry" {
        for _, c := range catmdownpvclrangetable.Catmdownpvclrangeentry {
            if catmdownpvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry{}
        catmdownpvclrangetable.Catmdownpvclrangeentry = append(catmdownpvclrangetable.Catmdownpvclrangeentry, child)
        return &catmdownpvclrangetable.Catmdownpvclrangeentry[len(catmdownpvclrangetable.Catmdownpvclrangeentry)-1]
    }
    return nil
}

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmdownpvclrangetable.Catmdownpvclrangeentry {
        children[catmdownpvclrangetable.Catmdownpvclrangeentry[i].GetSegmentPath()] = &catmdownpvclrangetable.Catmdownpvclrangeentry[i]
    }
    return children
}

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetYangName() string { return "catmDownPVclRangeTable" }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) SetParent(parent types.Entity) { catmdownpvclrangetable.parent = parent }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetParent() types.Entity { return catmdownpvclrangetable.parent }

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  atmVclOperStatus to 
// have detected as Down in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus has been
    // detected as Down in the  corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmdownpvcllowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the  atmVclOperStatus has been
    // detected as Down in the  corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmdownpvclhigherrangevalue interface{}

    // The time stamp at which the first atmVclOperStatus is detected as Down on
    // any of the PVCLs in the range in the corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    Catmdownpvclrangestart interface{}

    // The time stamp at which the last atmVclOperStatus is detected as Down on
    // any of the PVCLs in the range in the corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    Catmdownpvclrangeend interface{}

    // The time stamp at which the first atmVclOperStatus is detected as Up on any
    // of the PVCLs in the range in the previous catmIntfPvcUp2Trap notification.
    // The type is interface{} with range: 0..4294967295.
    Catmprevuppvclrangestart interface{}

    // The time stamp at which the last atmVclOperStatus is detected as Up on any
    // of the PVCLs in the range in the previous catmIntfPvcUp2Trap notification.
    // The type is interface{} with range: 0..4294967295.
    Catmprevuppvclrangeend interface{}

    // Type of OAM failure. The type is CatmOAMFailureType.
    Catmpvclrangefailurereason interface{}
}

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetFilter() yfilter.YFilter { return catmdownpvclrangeentry.YFilter }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) SetFilter(yf yfilter.YFilter) { catmdownpvclrangeentry.YFilter = yf }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmDownPVclLowerRangeValue" { return "Catmdownpvcllowerrangevalue" }
    if yname == "catmDownPVclHigherRangeValue" { return "Catmdownpvclhigherrangevalue" }
    if yname == "catmDownPVclRangeStart" { return "Catmdownpvclrangestart" }
    if yname == "catmDownPVclRangeEnd" { return "Catmdownpvclrangeend" }
    if yname == "catmPrevUpPVclRangeStart" { return "Catmprevuppvclrangestart" }
    if yname == "catmPrevUpPVclRangeEnd" { return "Catmprevuppvclrangeend" }
    if yname == "catmPVclRangeFailureReason" { return "Catmpvclrangefailurereason" }
    return ""
}

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetSegmentPath() string {
    return "catmDownPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmdownpvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmdownpvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmdownpvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmdownpvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmdownpvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmdownpvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmDownPVclLowerRangeValue"] = catmdownpvclrangeentry.Catmdownpvcllowerrangevalue
    leafs["catmDownPVclHigherRangeValue"] = catmdownpvclrangeentry.Catmdownpvclhigherrangevalue
    leafs["catmDownPVclRangeStart"] = catmdownpvclrangeentry.Catmdownpvclrangestart
    leafs["catmDownPVclRangeEnd"] = catmdownpvclrangeentry.Catmdownpvclrangeend
    leafs["catmPrevUpPVclRangeStart"] = catmdownpvclrangeentry.Catmprevuppvclrangestart
    leafs["catmPrevUpPVclRangeEnd"] = catmdownpvclrangeentry.Catmprevuppvclrangeend
    leafs["catmPVclRangeFailureReason"] = catmdownpvclrangeentry.Catmpvclrangefailurereason
    return leafs
}

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetYangName() string { return "catmDownPVclRangeEntry" }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) SetParent(parent types.Entity) { catmdownpvclrangeentry.parent = parent }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetParent() types.Entity { return catmdownpvclrangeentry.parent }

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetParentYangName() string { return "catmDownPVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed to UP
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed to UP in the last PVC notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry.
    Catmcurstatusuppvclentry []CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry
}

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetFilter() yfilter.YFilter { return catmcurstatusuppvcltable.YFilter }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) SetFilter(yf yfilter.YFilter) { catmcurstatusuppvcltable.YFilter = yf }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetGoName(yname string) string {
    if yname == "catmCurStatusUpPVclEntry" { return "Catmcurstatusuppvclentry" }
    return ""
}

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetSegmentPath() string {
    return "catmCurStatusUpPVclTable"
}

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmCurStatusUpPVclEntry" {
        for _, c := range catmcurstatusuppvcltable.Catmcurstatusuppvclentry {
            if catmcurstatusuppvcltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry{}
        catmcurstatusuppvcltable.Catmcurstatusuppvclentry = append(catmcurstatusuppvcltable.Catmcurstatusuppvclentry, child)
        return &catmcurstatusuppvcltable.Catmcurstatusuppvclentry[len(catmcurstatusuppvcltable.Catmcurstatusuppvclentry)-1]
    }
    return nil
}

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmcurstatusuppvcltable.Catmcurstatusuppvclentry {
        children[catmcurstatusuppvcltable.Catmcurstatusuppvclentry[i].GetSegmentPath()] = &catmcurstatusuppvcltable.Catmcurstatusuppvclentry[i]
    }
    return children
}

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetBundleName() string { return "cisco_ios_xe" }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetYangName() string { return "catmCurStatusUpPVclTable" }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) SetParent(parent types.Entity) { catmcurstatusuppvcltable.parent = parent }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetParent() types.Entity { return catmcurstatusuppvcltable.parent }

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed to UP in the last PVC notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry struct {
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

    // The number of Down to Up state transitions due to OAM loopback recovery
    // that has happened on this PVCL  in the last corresponding notification .
    // The type is interface{} with range: 0..4294967295.
    Catmpvclstatusuptransition interface{}

    // The time stamp at which this PVCL changed state to UP  for the first time
    // due to OAM loopback recovery in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    Catmpvclstatusupstart interface{}

    // The time stamp at which this PVCL changed state to UP  for the last time
    // due to OAM loopback recovery in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    Catmpvclstatusupend interface{}

    // The number of Down to Up state transitions that has  happened on this PVCL
    // in the last corresponding notification  due to Segment CC OAM recovery. The
    // type is interface{} with range: 0..4294967295.
    Catmpvclsegccstatusuptransition interface{}

    // The time stamp at which this PVCL changed state to Up for  the first time
    // in the last corresponding notification due to Segment CC OAM recovery. The
    // type is interface{} with range: 0..4294967295.
    Catmpvclsegccstatusupstart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification due to Segment CC  OAM recovery. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclsegccstatusupend interface{}

    // The number of Down to UP state transitions that has  happened on this PVCL
    // in the last notification  interval due to End CC OAM recovery. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclendccstatusuptransition interface{}

    // The time stamp at which this PVCL changed state to Up for the first time in
    // the last corresponding notification  due to End CC OAM recovery. The type
    // is interface{} with range: 0..4294967295.
    Catmpvclendccstatusupstart interface{}

    // The time stamp at which this PVCL changed state to Up for the last time in
    // the last corresponding notification  due to End CC OAM recovery. The type
    // is interface{} with range: 0..4294967295.
    Catmpvclendccstatusupend interface{}

    // The number of Down to Up state transitions that  has happened on this PVCL
    // in the last notification  interval due to AIS RDI OAM recovery. The type is
    // interface{} with range: 0..4294967295.
    Catmpvclaisrdistatusuptransition interface{}

    // The time stamp at which this PVCL changed state to Up for the first time in
    // the last corresponding notification  due to AIS/RDI OAM recovery. The type
    // is interface{} with range: 0..4294967295.
    Catmpvclaisrdistatusupstart interface{}

    // The time stamp at which this PVCL changed state to Up for the last time in
    // the last corresponding notification  due to AIS/RDI OAM recovery. The type
    // is interface{} with range: 0..4294967295.
    Catmpvclaisrdistatusupend interface{}

    // The time stamp at which this PVCL changed state to UP last time in the last
    // corresponding notification . The type is interface{} with range:
    // 0..4294967295.
    Catmpvclcurrecovertime interface{}

    // The time stamp at which this PVCL changed state to DOWN last time in the
    // previous corresponding notification . The type is interface{} with range:
    // 0..4294967295.
    Catmpvclprevfailtime interface{}

    // Type of OAM Recovered. The type is CatmOAMRecoveryType.
    Catmpvclrecoveryreason interface{}
}

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetFilter() yfilter.YFilter { return catmcurstatusuppvclentry.YFilter }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) SetFilter(yf yfilter.YFilter) { catmcurstatusuppvclentry.YFilter = yf }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "atmVclVci" { return "Atmvclvci" }
    if yname == "catmPVclStatusUpTransition" { return "Catmpvclstatusuptransition" }
    if yname == "catmPVclStatusUpStart" { return "Catmpvclstatusupstart" }
    if yname == "catmPVclStatusUpEnd" { return "Catmpvclstatusupend" }
    if yname == "catmPVclSegCCStatusUpTransition" { return "Catmpvclsegccstatusuptransition" }
    if yname == "catmPVclSegCCStatusUpStart" { return "Catmpvclsegccstatusupstart" }
    if yname == "catmPVclSegCCStatusUpEnd" { return "Catmpvclsegccstatusupend" }
    if yname == "catmPVclEndCCStatusUpTransition" { return "Catmpvclendccstatusuptransition" }
    if yname == "catmPVclEndCCStatusUpStart" { return "Catmpvclendccstatusupstart" }
    if yname == "catmPVclEndCCStatusUpEnd" { return "Catmpvclendccstatusupend" }
    if yname == "catmPVclAISRDIStatusUpTransition" { return "Catmpvclaisrdistatusuptransition" }
    if yname == "catmPVclAISRDIStatusUpStart" { return "Catmpvclaisrdistatusupstart" }
    if yname == "catmPVclAISRDIStatusUpEnd" { return "Catmpvclaisrdistatusupend" }
    if yname == "catmPVclCurRecoverTime" { return "Catmpvclcurrecovertime" }
    if yname == "catmPVclPrevFailTime" { return "Catmpvclprevfailtime" }
    if yname == "catmPVclRecoveryReason" { return "Catmpvclrecoveryreason" }
    return ""
}

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetSegmentPath() string {
    return "catmCurStatusUpPVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmcurstatusuppvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmcurstatusuppvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", catmcurstatusuppvclentry.Atmvclvci) + "']"
}

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmcurstatusuppvclentry.Ifindex
    leafs["atmVclVpi"] = catmcurstatusuppvclentry.Atmvclvpi
    leafs["atmVclVci"] = catmcurstatusuppvclentry.Atmvclvci
    leafs["catmPVclStatusUpTransition"] = catmcurstatusuppvclentry.Catmpvclstatusuptransition
    leafs["catmPVclStatusUpStart"] = catmcurstatusuppvclentry.Catmpvclstatusupstart
    leafs["catmPVclStatusUpEnd"] = catmcurstatusuppvclentry.Catmpvclstatusupend
    leafs["catmPVclSegCCStatusUpTransition"] = catmcurstatusuppvclentry.Catmpvclsegccstatusuptransition
    leafs["catmPVclSegCCStatusUpStart"] = catmcurstatusuppvclentry.Catmpvclsegccstatusupstart
    leafs["catmPVclSegCCStatusUpEnd"] = catmcurstatusuppvclentry.Catmpvclsegccstatusupend
    leafs["catmPVclEndCCStatusUpTransition"] = catmcurstatusuppvclentry.Catmpvclendccstatusuptransition
    leafs["catmPVclEndCCStatusUpStart"] = catmcurstatusuppvclentry.Catmpvclendccstatusupstart
    leafs["catmPVclEndCCStatusUpEnd"] = catmcurstatusuppvclentry.Catmpvclendccstatusupend
    leafs["catmPVclAISRDIStatusUpTransition"] = catmcurstatusuppvclentry.Catmpvclaisrdistatusuptransition
    leafs["catmPVclAISRDIStatusUpStart"] = catmcurstatusuppvclentry.Catmpvclaisrdistatusupstart
    leafs["catmPVclAISRDIStatusUpEnd"] = catmcurstatusuppvclentry.Catmpvclaisrdistatusupend
    leafs["catmPVclCurRecoverTime"] = catmcurstatusuppvclentry.Catmpvclcurrecovertime
    leafs["catmPVclPrevFailTime"] = catmcurstatusuppvclentry.Catmpvclprevfailtime
    leafs["catmPVclRecoveryReason"] = catmcurstatusuppvclentry.Catmpvclrecoveryreason
    return leafs
}

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetYangName() string { return "catmCurStatusUpPVclEntry" }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) SetParent(parent types.Entity) { catmcurstatusuppvclentry.parent = parent }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetParent() types.Entity { return catmcurstatusuppvclentry.parent }

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetParentYangName() string { return "catmCurStatusUpPVclTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and loopback OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  loopback OAM status to  have detected as recovered in the last
    // notification  interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry.
    Catmstatusuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry
}

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetFilter() yfilter.YFilter { return catmstatusuppvclrangetable.YFilter }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) SetFilter(yf yfilter.YFilter) { catmstatusuppvclrangetable.YFilter = yf }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetGoName(yname string) string {
    if yname == "catmStatusUpPVclRangeEntry" { return "Catmstatusuppvclrangeentry" }
    return ""
}

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetSegmentPath() string {
    return "catmStatusUpPVclRangeTable"
}

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmStatusUpPVclRangeEntry" {
        for _, c := range catmstatusuppvclrangetable.Catmstatusuppvclrangeentry {
            if catmstatusuppvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry{}
        catmstatusuppvclrangetable.Catmstatusuppvclrangeentry = append(catmstatusuppvclrangetable.Catmstatusuppvclrangeentry, child)
        return &catmstatusuppvclrangetable.Catmstatusuppvclrangeentry[len(catmstatusuppvclrangetable.Catmstatusuppvclrangeentry)-1]
    }
    return nil
}

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmstatusuppvclrangetable.Catmstatusuppvclrangeentry {
        children[catmstatusuppvclrangetable.Catmstatusuppvclrangeentry[i].GetSegmentPath()] = &catmstatusuppvclrangetable.Catmstatusuppvclrangeentry[i]
    }
    return children
}

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetYangName() string { return "catmStatusUpPVclRangeTable" }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) SetParent(parent types.Entity) { catmstatusuppvclrangetable.parent = parent }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetParent() types.Entity { return catmstatusuppvclrangetable.parent }

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  loopback OAM status to 
// have detected as recovered in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the  Loopback OAM recovery has
    // been detected in the last  corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmpvcluplowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the  Loopback OAM recovery has
    // been detected in the last  corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmpvcluphigherrangevalue interface{}

    // The time stamp at which the first Loopback OAM recovery is detected on any
    // of the PVCLs in the range in the last corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    Catmpvclrangestatusupstart interface{}

    // The time stamp at which the last Loopback OAM recovery is detected on any
    // of the PVCLs in the range in the last corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    Catmpvclrangestatusupend interface{}
}

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetFilter() yfilter.YFilter { return catmstatusuppvclrangeentry.YFilter }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) SetFilter(yf yfilter.YFilter) { catmstatusuppvclrangeentry.YFilter = yf }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmPVclUpLowerRangeValue" { return "Catmpvcluplowerrangevalue" }
    if yname == "catmPVclUpHigherRangeValue" { return "Catmpvcluphigherrangevalue" }
    if yname == "catmPVclRangeStatusUpStart" { return "Catmpvclrangestatusupstart" }
    if yname == "catmPVclRangeStatusUpEnd" { return "Catmpvclrangestatusupend" }
    return ""
}

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetSegmentPath() string {
    return "catmStatusUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmstatusuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmstatusuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmstatusuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmstatusuppvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmstatusuppvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmstatusuppvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmPVclUpLowerRangeValue"] = catmstatusuppvclrangeentry.Catmpvcluplowerrangevalue
    leafs["catmPVclUpHigherRangeValue"] = catmstatusuppvclrangeentry.Catmpvcluphigherrangevalue
    leafs["catmPVclRangeStatusUpStart"] = catmstatusuppvclrangeentry.Catmpvclrangestatusupstart
    leafs["catmPVclRangeStatusUpEnd"] = catmstatusuppvclrangeentry.Catmpvclrangestatusupend
    return leafs
}

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetYangName() string { return "catmStatusUpPVclRangeEntry" }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) SetParent(parent types.Entity) { catmstatusuppvclrangeentry.parent = parent }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetParent() types.Entity { return catmstatusuppvclrangeentry.parent }

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetParentYangName() string { return "catmStatusUpPVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and Segment CC OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and Segment CC OAM status to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry.
    Catmsegccstatusuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry
}

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetFilter() yfilter.YFilter { return catmsegccstatusuppvclrangetable.YFilter }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) SetFilter(yf yfilter.YFilter) { catmsegccstatusuppvclrangetable.YFilter = yf }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetGoName(yname string) string {
    if yname == "catmSegCCStatusUpPVclRangeEntry" { return "Catmsegccstatusuppvclrangeentry" }
    return ""
}

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetSegmentPath() string {
    return "catmSegCCStatusUpPVclRangeTable"
}

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmSegCCStatusUpPVclRangeEntry" {
        for _, c := range catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry {
            if catmsegccstatusuppvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry{}
        catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry = append(catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry, child)
        return &catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry[len(catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry)-1]
    }
    return nil
}

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry {
        children[catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry[i].GetSegmentPath()] = &catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry[i]
    }
    return children
}

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetYangName() string { return "catmSegCCStatusUpPVclRangeTable" }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) SetParent(parent types.Entity) { catmsegccstatusuppvclrangetable.parent = parent }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetParent() types.Entity { return catmsegccstatusuppvclrangetable.parent }

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and Segment CC OAM status to
// have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the Segment CC OAM recovery
    // has been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmpvclsegccuplowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the Segment CC OAM recovery has
    // been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmpvclsegccuphigherrangevalue interface{}

    // The time stamp at which the first Segment CC OAM recovery is detected on
    // any of the PVCLs in the range in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    Catmpvclsegccrangestatusupstart interface{}

    // The time stamp at which the last Segment CC OAM recovery is detected on any
    // of the PVCLs in the range in the last corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    Catmpvclsegccrangestatusupend interface{}
}

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetFilter() yfilter.YFilter { return catmsegccstatusuppvclrangeentry.YFilter }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) SetFilter(yf yfilter.YFilter) { catmsegccstatusuppvclrangeentry.YFilter = yf }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmPVclSegCCUpLowerRangeValue" { return "Catmpvclsegccuplowerrangevalue" }
    if yname == "catmPVclSegCCUpHigherRangeValue" { return "Catmpvclsegccuphigherrangevalue" }
    if yname == "catmPVclSegCCRangeStatusUpStart" { return "Catmpvclsegccrangestatusupstart" }
    if yname == "catmPVclSegCCRangeStatusUpEnd" { return "Catmpvclsegccrangestatusupend" }
    return ""
}

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetSegmentPath() string {
    return "catmSegCCStatusUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmsegccstatusuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmsegccstatusuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmsegccstatusuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmsegccstatusuppvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmsegccstatusuppvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmsegccstatusuppvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmPVclSegCCUpLowerRangeValue"] = catmsegccstatusuppvclrangeentry.Catmpvclsegccuplowerrangevalue
    leafs["catmPVclSegCCUpHigherRangeValue"] = catmsegccstatusuppvclrangeentry.Catmpvclsegccuphigherrangevalue
    leafs["catmPVclSegCCRangeStatusUpStart"] = catmsegccstatusuppvclrangeentry.Catmpvclsegccrangestatusupstart
    leafs["catmPVclSegCCRangeStatusUpEnd"] = catmsegccstatusuppvclrangeentry.Catmpvclsegccrangestatusupend
    return leafs
}

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetYangName() string { return "catmSegCCStatusUpPVclRangeEntry" }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) SetParent(parent types.Entity) { catmsegccstatusuppvclrangeentry.parent = parent }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetParent() types.Entity { return catmsegccstatusuppvclrangeentry.parent }

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetParentYangName() string { return "catmSegCCStatusUpPVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and End-to-End CC OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and End-to-End CC OAM status  to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry.
    Catmendccstatusuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry
}

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetFilter() yfilter.YFilter { return catmendccstatusuppvclrangetable.YFilter }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) SetFilter(yf yfilter.YFilter) { catmendccstatusuppvclrangetable.YFilter = yf }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetGoName(yname string) string {
    if yname == "catmEndCCStatusUpPVclRangeEntry" { return "Catmendccstatusuppvclrangeentry" }
    return ""
}

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetSegmentPath() string {
    return "catmEndCCStatusUpPVclRangeTable"
}

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmEndCCStatusUpPVclRangeEntry" {
        for _, c := range catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry {
            if catmendccstatusuppvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry{}
        catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry = append(catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry, child)
        return &catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry[len(catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry)-1]
    }
    return nil
}

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry {
        children[catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry[i].GetSegmentPath()] = &catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry[i]
    }
    return children
}

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetYangName() string { return "catmEndCCStatusUpPVclRangeTable" }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) SetParent(parent types.Entity) { catmendccstatusuppvclrangetable.parent = parent }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetParent() types.Entity { return catmendccstatusuppvclrangetable.parent }

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and End-to-End CC OAM status 
// to have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the End-to-End CC OAM recovery
    // has been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmpvclendccuplowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the End-to-End CC OAM recovery
    // has been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmpvclendccuphigherrangevalue interface{}

    // The time stamp at which the first End-to-End CC OAM recovery is detected on
    // any of the PVCLs in the range in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    Catmpvclendccrangestatusupstart interface{}

    // The time stamp at which the last End-to-End CC OAM recovery is detected on
    // any of the PVCLs in the range in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    Catmpvclendccrangestatusupend interface{}
}

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetFilter() yfilter.YFilter { return catmendccstatusuppvclrangeentry.YFilter }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) SetFilter(yf yfilter.YFilter) { catmendccstatusuppvclrangeentry.YFilter = yf }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmPVclEndCCUpLowerRangeValue" { return "Catmpvclendccuplowerrangevalue" }
    if yname == "catmPVclEndCCUpHigherRangeValue" { return "Catmpvclendccuphigherrangevalue" }
    if yname == "catmPVclEndCCRangeStatusUpStart" { return "Catmpvclendccrangestatusupstart" }
    if yname == "catmPVclEndCCRangeStatusUpEnd" { return "Catmpvclendccrangestatusupend" }
    return ""
}

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetSegmentPath() string {
    return "catmEndCCStatusUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmendccstatusuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmendccstatusuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmendccstatusuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmendccstatusuppvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmendccstatusuppvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmendccstatusuppvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmPVclEndCCUpLowerRangeValue"] = catmendccstatusuppvclrangeentry.Catmpvclendccuplowerrangevalue
    leafs["catmPVclEndCCUpHigherRangeValue"] = catmendccstatusuppvclrangeentry.Catmpvclendccuphigherrangevalue
    leafs["catmPVclEndCCRangeStatusUpStart"] = catmendccstatusuppvclrangeentry.Catmpvclendccrangestatusupstart
    leafs["catmPVclEndCCRangeStatusUpEnd"] = catmendccstatusuppvclrangeentry.Catmpvclendccrangestatusupend
    return leafs
}

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetYangName() string { return "catmEndCCStatusUpPVclRangeEntry" }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) SetParent(parent types.Entity) { catmendccstatusuppvclrangeentry.parent = parent }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetParent() types.Entity { return catmendccstatusuppvclrangeentry.parent }

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetParentYangName() string { return "catmEndCCStatusUpPVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and AISRDI OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and AISRDI OAM status  to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry.
    Catmaisrdistatusuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry
}

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetFilter() yfilter.YFilter { return catmaisrdistatusuppvclrangetable.YFilter }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) SetFilter(yf yfilter.YFilter) { catmaisrdistatusuppvclrangetable.YFilter = yf }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetGoName(yname string) string {
    if yname == "catmAISRDIStatusUpPVclRangeEntry" { return "Catmaisrdistatusuppvclrangeentry" }
    return ""
}

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetSegmentPath() string {
    return "catmAISRDIStatusUpPVclRangeTable"
}

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmAISRDIStatusUpPVclRangeEntry" {
        for _, c := range catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry {
            if catmaisrdistatusuppvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry{}
        catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry = append(catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry, child)
        return &catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry[len(catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry)-1]
    }
    return nil
}

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry {
        children[catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry[i].GetSegmentPath()] = &catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry[i]
    }
    return children
}

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetYangName() string { return "catmAISRDIStatusUpPVclRangeTable" }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) SetParent(parent types.Entity) { catmaisrdistatusuppvclrangetable.parent = parent }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetParent() types.Entity { return catmaisrdistatusuppvclrangetable.parent }

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and AISRDI OAM status 
// to have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the AISRDI OAM recovery has
    // been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmpvclaisrdiuplowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the AISRDI OAM recovery has
    // been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    Catmpvclaisrdiuphigherrangevalue interface{}

    // The time stamp at which the first AISRDI OAM recovery is detected on any of
    // the PVCLs in the range in the last corresponding notification . The type is
    // interface{} with range: 0..4294967295.
    Catmpvclaisrdirangestatusupstart interface{}

    // The time stamp at which the last AISRDI OAM recovery is detected on any of
    // the PVCLs in the range in the last corresponding notification . The type is
    // interface{} with range: 0..4294967295.
    Catmpvclaisrdirangestatusupend interface{}
}

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetFilter() yfilter.YFilter { return catmaisrdistatusuppvclrangeentry.YFilter }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) SetFilter(yf yfilter.YFilter) { catmaisrdistatusuppvclrangeentry.YFilter = yf }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmPVclAISRDIUpLowerRangeValue" { return "Catmpvclaisrdiuplowerrangevalue" }
    if yname == "catmPVclAISRDIUpHigherRangeValue" { return "Catmpvclaisrdiuphigherrangevalue" }
    if yname == "catmPVclAISRDIRangeStatusUpStart" { return "Catmpvclaisrdirangestatusupstart" }
    if yname == "catmPVclAISRDIRangeStatusUpEnd" { return "Catmpvclaisrdirangestatusupend" }
    return ""
}

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetSegmentPath() string {
    return "catmAISRDIStatusUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmaisrdistatusuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmaisrdistatusuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmaisrdistatusuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmaisrdistatusuppvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmaisrdistatusuppvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmaisrdistatusuppvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmPVclAISRDIUpLowerRangeValue"] = catmaisrdistatusuppvclrangeentry.Catmpvclaisrdiuplowerrangevalue
    leafs["catmPVclAISRDIUpHigherRangeValue"] = catmaisrdistatusuppvclrangeentry.Catmpvclaisrdiuphigherrangevalue
    leafs["catmPVclAISRDIRangeStatusUpStart"] = catmaisrdistatusuppvclrangeentry.Catmpvclaisrdirangestatusupstart
    leafs["catmPVclAISRDIRangeStatusUpEnd"] = catmaisrdistatusuppvclrangeentry.Catmpvclaisrdirangestatusupend
    return leafs
}

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetYangName() string { return "catmAISRDIStatusUpPVclRangeEntry" }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) SetParent(parent types.Entity) { catmaisrdistatusuppvclrangeentry.parent = parent }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetParent() types.Entity { return catmaisrdistatusuppvclrangeentry.parent }

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetParentYangName() string { return "catmAISRDIStatusUpPVclRangeTable" }

// CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have detected as Up
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  atmVclOperStatus to  have detected as Up in the last notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry.
    Catmuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry
}

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetFilter() yfilter.YFilter { return catmuppvclrangetable.YFilter }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) SetFilter(yf yfilter.YFilter) { catmuppvclrangetable.YFilter = yf }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetGoName(yname string) string {
    if yname == "catmUpPVclRangeEntry" { return "Catmuppvclrangeentry" }
    return ""
}

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetSegmentPath() string {
    return "catmUpPVclRangeTable"
}

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "catmUpPVclRangeEntry" {
        for _, c := range catmuppvclrangetable.Catmuppvclrangeentry {
            if catmuppvclrangetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry{}
        catmuppvclrangetable.Catmuppvclrangeentry = append(catmuppvclrangetable.Catmuppvclrangeentry, child)
        return &catmuppvclrangetable.Catmuppvclrangeentry[len(catmuppvclrangetable.Catmuppvclrangeentry)-1]
    }
    return nil
}

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range catmuppvclrangetable.Catmuppvclrangeentry {
        children[catmuppvclrangetable.Catmuppvclrangeentry[i].GetSegmentPath()] = &catmuppvclrangetable.Catmuppvclrangeentry[i]
    }
    return children
}

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetBundleName() string { return "cisco_ios_xe" }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetYangName() string { return "catmUpPVclRangeTable" }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) SetParent(parent types.Entity) { catmuppvclrangetable.parent = parent }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetParent() types.Entity { return catmuppvclrangetable.parent }

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetParentYangName() string { return "CISCO-ATM-PVCTRAP-EXTN-MIB" }

// CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  atmVclOperStatus to 
// have detected as Up in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry_Catmstatuschangepvclrangeindex
    Catmstatuschangepvclrangeindex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus has been
    // detected as Up in the  corresponding notification . The type is interface{}
    // with range: 0..65536.
    Catmuppvcllowerrangevalue interface{}

    // The last PVCL in a range of PVCLs for which the  atmVclOperStatus has been
    // detected as Up in the  corresponding notification . The type is interface{}
    // with range: 0..65536.
    Catmuppvclhigherrangevalue interface{}

    // The time stamp at which the first atmVclOperStatus is detected as Up on any
    // of the PVCLs in the range in the corresponding notification . The type is
    // interface{} with range: 0..4294967295.
    Catmuppvclrangestart interface{}

    // The time stamp at which the last atmVclOperStatus is detected as Up on any
    // of the PVCLs in the range in the corresponding notification . The type is
    // interface{} with range: 0..4294967295.
    Catmuppvclrangeend interface{}

    // The time stamp at which the first atmVclOperStatus is detected as Down on
    // any of the PVCLs in the range in the previous catmIntfPvcDownTrap
    // notification. The type is interface{} with range: 0..4294967295.
    Catmprevdownpvclrangestart interface{}

    // The time stamp at which the last atmVclOperStatus is detected as Down on
    // any of the PVCLs in the range in the previous catmIntfPvcDownTrap
    // notification. The type is interface{} with range: 0..4294967295.
    Catmprevdownpvclrangeend interface{}

    // Type of OAM Recovered. The type is CatmOAMRecoveryType.
    Catmpvclrangerecoveryreason interface{}
}

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetFilter() yfilter.YFilter { return catmuppvclrangeentry.YFilter }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) SetFilter(yf yfilter.YFilter) { catmuppvclrangeentry.YFilter = yf }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "catmStatusChangePVclRangeIndex" { return "Catmstatuschangepvclrangeindex" }
    if yname == "catmUpPVclLowerRangeValue" { return "Catmuppvcllowerrangevalue" }
    if yname == "catmUpPVclHigherRangeValue" { return "Catmuppvclhigherrangevalue" }
    if yname == "catmUpPVclRangeStart" { return "Catmuppvclrangestart" }
    if yname == "catmUpPVclRangeEnd" { return "Catmuppvclrangeend" }
    if yname == "catmPrevDownPVclRangeStart" { return "Catmprevdownpvclrangestart" }
    if yname == "catmPrevDownPVclRangeEnd" { return "Catmprevdownpvclrangeend" }
    if yname == "catmPVclRangeRecoveryReason" { return "Catmpvclrangerecoveryreason" }
    return ""
}

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetSegmentPath() string {
    return "catmUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
}

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = catmuppvclrangeentry.Ifindex
    leafs["atmVclVpi"] = catmuppvclrangeentry.Atmvclvpi
    leafs["catmStatusChangePVclRangeIndex"] = catmuppvclrangeentry.Catmstatuschangepvclrangeindex
    leafs["catmUpPVclLowerRangeValue"] = catmuppvclrangeentry.Catmuppvcllowerrangevalue
    leafs["catmUpPVclHigherRangeValue"] = catmuppvclrangeentry.Catmuppvclhigherrangevalue
    leafs["catmUpPVclRangeStart"] = catmuppvclrangeentry.Catmuppvclrangestart
    leafs["catmUpPVclRangeEnd"] = catmuppvclrangeentry.Catmuppvclrangeend
    leafs["catmPrevDownPVclRangeStart"] = catmuppvclrangeentry.Catmprevdownpvclrangestart
    leafs["catmPrevDownPVclRangeEnd"] = catmuppvclrangeentry.Catmprevdownpvclrangeend
    leafs["catmPVclRangeRecoveryReason"] = catmuppvclrangeentry.Catmpvclrangerecoveryreason
    return leafs
}

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetBundleName() string { return "cisco_ios_xe" }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetYangName() string { return "catmUpPVclRangeEntry" }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) SetParent(parent types.Entity) { catmuppvclrangeentry.parent = parent }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetParent() types.Entity { return catmuppvclrangeentry.parent }

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetParentYangName() string { return "catmUpPVclRangeTable" }

