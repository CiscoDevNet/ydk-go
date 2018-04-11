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
    EntityData types.CommonEntityData
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

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetEntityData() *types.CommonEntityData {
    cISCOATMPVCTRAPEXTNMIB.EntityData.YFilter = cISCOATMPVCTRAPEXTNMIB.YFilter
    cISCOATMPVCTRAPEXTNMIB.EntityData.YangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    cISCOATMPVCTRAPEXTNMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOATMPVCTRAPEXTNMIB.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    cISCOATMPVCTRAPEXTNMIB.EntityData.SegmentPath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB"
    cISCOATMPVCTRAPEXTNMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOATMPVCTRAPEXTNMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOATMPVCTRAPEXTNMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOATMPVCTRAPEXTNMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmCurStatChangePVclTable"] = types.YChild{"Catmcurstatchangepvcltable", &cISCOATMPVCTRAPEXTNMIB.Catmcurstatchangepvcltable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmStatusChangePVclRangeTable"] = types.YChild{"Catmstatuschangepvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmstatuschangepvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmSegCCStatusChPVclRangeTable"] = types.YChild{"Catmsegccstatuschpvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmsegccstatuschpvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmEndCCStatusChPVclRangeTable"] = types.YChild{"Catmendccstatuschpvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmendccstatuschpvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmAISRDIStatusChPVclRangeTable"] = types.YChild{"Catmaisrdistatuschpvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmaisrdistatuschpvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmDownPVclRangeTable"] = types.YChild{"Catmdownpvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmdownpvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmCurStatusUpPVclTable"] = types.YChild{"Catmcurstatusuppvcltable", &cISCOATMPVCTRAPEXTNMIB.Catmcurstatusuppvcltable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmStatusUpPVclRangeTable"] = types.YChild{"Catmstatusuppvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmstatusuppvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmSegCCStatusUpPVclRangeTable"] = types.YChild{"Catmsegccstatusuppvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmsegccstatusuppvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmEndCCStatusUpPVclRangeTable"] = types.YChild{"Catmendccstatusuppvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmendccstatusuppvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmAISRDIStatusUpPVclRangeTable"] = types.YChild{"Catmaisrdistatusuppvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmaisrdistatusuppvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children["catmUpPVclRangeTable"] = types.YChild{"Catmuppvclrangetable", &cISCOATMPVCTRAPEXTNMIB.Catmuppvclrangetable}
    cISCOATMPVCTRAPEXTNMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOATMPVCTRAPEXTNMIB.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed in the
// last corresponding PVC notification.
type CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed in the last corresponding PVC
    // notification. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry.
    Catmcurstatchangepvclentry []CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry
}

func (catmcurstatchangepvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable) GetEntityData() *types.CommonEntityData {
    catmcurstatchangepvcltable.EntityData.YFilter = catmcurstatchangepvcltable.YFilter
    catmcurstatchangepvcltable.EntityData.YangName = "catmCurStatChangePVclTable"
    catmcurstatchangepvcltable.EntityData.BundleName = "cisco_ios_xe"
    catmcurstatchangepvcltable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmcurstatchangepvcltable.EntityData.SegmentPath = "catmCurStatChangePVclTable"
    catmcurstatchangepvcltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmcurstatchangepvcltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmcurstatchangepvcltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmcurstatchangepvcltable.EntityData.Children = make(map[string]types.YChild)
    catmcurstatchangepvcltable.EntityData.Children["catmCurStatChangePVclEntry"] = types.YChild{"Catmcurstatchangepvclentry", nil}
    for i := range catmcurstatchangepvcltable.Catmcurstatchangepvclentry {
        catmcurstatchangepvcltable.EntityData.Children[types.GetSegmentPath(&catmcurstatchangepvcltable.Catmcurstatchangepvclentry[i])] = types.YChild{"Catmcurstatchangepvclentry", &catmcurstatchangepvcltable.Catmcurstatchangepvclentry[i]}
    }
    catmcurstatchangepvcltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmcurstatchangepvcltable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed in the last corresponding PVC notification.
type CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry struct {
    EntityData types.CommonEntityData
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

func (catmcurstatchangepvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatchangepvcltable_Catmcurstatchangepvclentry) GetEntityData() *types.CommonEntityData {
    catmcurstatchangepvclentry.EntityData.YFilter = catmcurstatchangepvclentry.YFilter
    catmcurstatchangepvclentry.EntityData.YangName = "catmCurStatChangePVclEntry"
    catmcurstatchangepvclentry.EntityData.BundleName = "cisco_ios_xe"
    catmcurstatchangepvclentry.EntityData.ParentYangName = "catmCurStatChangePVclTable"
    catmcurstatchangepvclentry.EntityData.SegmentPath = "catmCurStatChangePVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmcurstatchangepvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmcurstatchangepvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", catmcurstatchangepvclentry.Atmvclvci) + "']"
    catmcurstatchangepvclentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmcurstatchangepvclentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmcurstatchangepvclentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmcurstatchangepvclentry.EntityData.Children = make(map[string]types.YChild)
    catmcurstatchangepvclentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmcurstatchangepvclentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmcurstatchangepvclentry.Ifindex}
    catmcurstatchangepvclentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmcurstatchangepvclentry.Atmvclvpi}
    catmcurstatchangepvclentry.EntityData.Leafs["atmVclVci"] = types.YLeaf{"Atmvclvci", catmcurstatchangepvclentry.Atmvclvci}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclStatusTransition"] = types.YLeaf{"Catmpvclstatustransition", catmcurstatchangepvclentry.Catmpvclstatustransition}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclStatusChangeStart"] = types.YLeaf{"Catmpvclstatuschangestart", catmcurstatchangepvclentry.Catmpvclstatuschangestart}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclStatusChangeEnd"] = types.YLeaf{"Catmpvclstatuschangeend", catmcurstatchangepvclentry.Catmpvclstatuschangeend}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclSegCCStatusTransition"] = types.YLeaf{"Catmpvclsegccstatustransition", catmcurstatchangepvclentry.Catmpvclsegccstatustransition}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclSegCCStatusChangeStart"] = types.YLeaf{"Catmpvclsegccstatuschangestart", catmcurstatchangepvclentry.Catmpvclsegccstatuschangestart}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclSegCCStatusChangeEnd"] = types.YLeaf{"Catmpvclsegccstatuschangeend", catmcurstatchangepvclentry.Catmpvclsegccstatuschangeend}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclEndCCStatusTransition"] = types.YLeaf{"Catmpvclendccstatustransition", catmcurstatchangepvclentry.Catmpvclendccstatustransition}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclEndCCStatusChangeStart"] = types.YLeaf{"Catmpvclendccstatuschangestart", catmcurstatchangepvclentry.Catmpvclendccstatuschangestart}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclEndCCStatusChangeEnd"] = types.YLeaf{"Catmpvclendccstatuschangeend", catmcurstatchangepvclentry.Catmpvclendccstatuschangeend}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclAISRDIStatusTransition"] = types.YLeaf{"Catmpvclaisrdistatustransition", catmcurstatchangepvclentry.Catmpvclaisrdistatustransition}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclAISRDIStatusChangeStart"] = types.YLeaf{"Catmpvclaisrdistatuschangestart", catmcurstatchangepvclentry.Catmpvclaisrdistatuschangestart}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclAISRDIStatusChangeEnd"] = types.YLeaf{"Catmpvclaisrdistatuschangeend", catmcurstatchangepvclentry.Catmpvclaisrdistatuschangeend}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclCurFailTime"] = types.YLeaf{"Catmpvclcurfailtime", catmcurstatchangepvclentry.Catmpvclcurfailtime}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclPrevRecoverTime"] = types.YLeaf{"Catmpvclprevrecovertime", catmcurstatchangepvclentry.Catmpvclprevrecovertime}
    catmcurstatchangepvclentry.EntityData.Leafs["catmPVclFailureReason"] = types.YLeaf{"Catmpvclfailurereason", catmcurstatchangepvclentry.Catmpvclfailurereason}
    return &(catmcurstatchangepvclentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed in the same
// direction in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed in the same direction in the last
    // notification  interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry.
    Catmstatuschangepvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry
}

func (catmstatuschangepvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable) GetEntityData() *types.CommonEntityData {
    catmstatuschangepvclrangetable.EntityData.YFilter = catmstatuschangepvclrangetable.YFilter
    catmstatuschangepvclrangetable.EntityData.YangName = "catmStatusChangePVclRangeTable"
    catmstatuschangepvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmstatuschangepvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmstatuschangepvclrangetable.EntityData.SegmentPath = "catmStatusChangePVclRangeTable"
    catmstatuschangepvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmstatuschangepvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmstatuschangepvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmstatuschangepvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmstatuschangepvclrangetable.EntityData.Children["catmStatusChangePVclRangeEntry"] = types.YChild{"Catmstatuschangepvclrangeentry", nil}
    for i := range catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry {
        catmstatuschangepvclrangetable.EntityData.Children[types.GetSegmentPath(&catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry[i])] = types.YChild{"Catmstatuschangepvclrangeentry", &catmstatuschangepvclrangetable.Catmstatuschangepvclrangeentry[i]}
    }
    catmstatuschangepvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmstatuschangepvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed in the same direction in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmstatuschangepvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatuschangepvclrangetable_Catmstatuschangepvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmstatuschangepvclrangeentry.EntityData.YFilter = catmstatuschangepvclrangeentry.YFilter
    catmstatuschangepvclrangeentry.EntityData.YangName = "catmStatusChangePVclRangeEntry"
    catmstatuschangepvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmstatuschangepvclrangeentry.EntityData.ParentYangName = "catmStatusChangePVclRangeTable"
    catmstatuschangepvclrangeentry.EntityData.SegmentPath = "catmStatusChangePVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmstatuschangepvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmstatuschangepvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmstatuschangepvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmstatuschangepvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmstatuschangepvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmstatuschangepvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmstatuschangepvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmstatuschangepvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmstatuschangepvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmstatuschangepvclrangeentry.Ifindex}
    catmstatuschangepvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmstatuschangepvclrangeentry.Atmvclvpi}
    catmstatuschangepvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmstatuschangepvclrangeentry.Catmstatuschangepvclrangeindex}
    catmstatuschangepvclrangeentry.EntityData.Leafs["catmPVclLowerRangeValue"] = types.YLeaf{"Catmpvcllowerrangevalue", catmstatuschangepvclrangeentry.Catmpvcllowerrangevalue}
    catmstatuschangepvclrangeentry.EntityData.Leafs["catmPVclHigherRangeValue"] = types.YLeaf{"Catmpvclhigherrangevalue", catmstatuschangepvclrangeentry.Catmpvclhigherrangevalue}
    catmstatuschangepvclrangeentry.EntityData.Leafs["catmPVclRangeStatusChangeStart"] = types.YLeaf{"Catmpvclrangestatuschangestart", catmstatuschangepvclrangeentry.Catmpvclrangestatuschangestart}
    catmstatuschangepvclrangeentry.EntityData.Leafs["catmPVclRangeStatusChangeEnd"] = types.YLeaf{"Catmpvclrangestatuschangeend", catmstatuschangepvclrangeentry.Catmpvclrangestatuschangeend}
    return &(catmstatuschangepvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to segment CC 
// failure in the same direction in the last PVC 
// corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to segment CC failure in the same
    // direction  in the last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry.
    Catmsegccstatuschpvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry
}

func (catmsegccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable) GetEntityData() *types.CommonEntityData {
    catmsegccstatuschpvclrangetable.EntityData.YFilter = catmsegccstatuschpvclrangetable.YFilter
    catmsegccstatuschpvclrangetable.EntityData.YangName = "catmSegCCStatusChPVclRangeTable"
    catmsegccstatuschpvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmsegccstatuschpvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmsegccstatuschpvclrangetable.EntityData.SegmentPath = "catmSegCCStatusChPVclRangeTable"
    catmsegccstatuschpvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmsegccstatuschpvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmsegccstatuschpvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmsegccstatuschpvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmsegccstatuschpvclrangetable.EntityData.Children["catmSegCCStatusChPVclRangeEntry"] = types.YChild{"Catmsegccstatuschpvclrangeentry", nil}
    for i := range catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry {
        catmsegccstatuschpvclrangetable.EntityData.Children[types.GetSegmentPath(&catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry[i])] = types.YChild{"Catmsegccstatuschpvclrangeentry", &catmsegccstatuschpvclrangetable.Catmsegccstatuschpvclrangeentry[i]}
    }
    catmsegccstatuschpvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmsegccstatuschpvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to segment CC failure in the same direction 
// in the last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmsegccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatuschpvclrangetable_Catmsegccstatuschpvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmsegccstatuschpvclrangeentry.EntityData.YFilter = catmsegccstatuschpvclrangeentry.YFilter
    catmsegccstatuschpvclrangeentry.EntityData.YangName = "catmSegCCStatusChPVclRangeEntry"
    catmsegccstatuschpvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmsegccstatuschpvclrangeentry.EntityData.ParentYangName = "catmSegCCStatusChPVclRangeTable"
    catmsegccstatuschpvclrangeentry.EntityData.SegmentPath = "catmSegCCStatusChPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmsegccstatuschpvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmsegccstatuschpvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmsegccstatuschpvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmsegccstatuschpvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmsegccstatuschpvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmsegccstatuschpvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmsegccstatuschpvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmsegccstatuschpvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmsegccstatuschpvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmsegccstatuschpvclrangeentry.Ifindex}
    catmsegccstatuschpvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmsegccstatuschpvclrangeentry.Atmvclvpi}
    catmsegccstatuschpvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmsegccstatuschpvclrangeentry.Catmstatuschangepvclrangeindex}
    catmsegccstatuschpvclrangeentry.EntityData.Leafs["catmPVclSegCCLowerRangeValue"] = types.YLeaf{"Catmpvclsegcclowerrangevalue", catmsegccstatuschpvclrangeentry.Catmpvclsegcclowerrangevalue}
    catmsegccstatuschpvclrangeentry.EntityData.Leafs["catmPVclSegCCHigherRangeValue"] = types.YLeaf{"Catmpvclsegcchigherrangevalue", catmsegccstatuschpvclrangeentry.Catmpvclsegcchigherrangevalue}
    catmsegccstatuschpvclrangeentry.EntityData.Leafs["catmPVclSegCCRangeStatusChStart"] = types.YLeaf{"Catmpvclsegccrangestatuschstart", catmsegccstatuschpvclrangeentry.Catmpvclsegccrangestatuschstart}
    catmsegccstatuschpvclrangeentry.EntityData.Leafs["catmPVclSegCCRangeStatusChEnd"] = types.YLeaf{"Catmpvclsegccrangestatuschend", catmsegccstatuschpvclrangeentry.Catmpvclsegccrangestatuschend}
    return &(catmsegccstatuschpvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to End CC failure
// in the same direction in the last PVC notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to End CC failure in the same
    // direction in the  last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry.
    Catmendccstatuschpvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry
}

func (catmendccstatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable) GetEntityData() *types.CommonEntityData {
    catmendccstatuschpvclrangetable.EntityData.YFilter = catmendccstatuschpvclrangetable.YFilter
    catmendccstatuschpvclrangetable.EntityData.YangName = "catmEndCCStatusChPVclRangeTable"
    catmendccstatuschpvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmendccstatuschpvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmendccstatuschpvclrangetable.EntityData.SegmentPath = "catmEndCCStatusChPVclRangeTable"
    catmendccstatuschpvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmendccstatuschpvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmendccstatuschpvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmendccstatuschpvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmendccstatuschpvclrangetable.EntityData.Children["catmEndCCStatusChPVclRangeEntry"] = types.YChild{"Catmendccstatuschpvclrangeentry", nil}
    for i := range catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry {
        catmendccstatuschpvclrangetable.EntityData.Children[types.GetSegmentPath(&catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry[i])] = types.YChild{"Catmendccstatuschpvclrangeentry", &catmendccstatuschpvclrangetable.Catmendccstatuschpvclrangeentry[i]}
    }
    catmendccstatuschpvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmendccstatuschpvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to End CC failure in the same direction in the 
// last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmendccstatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatuschpvclrangetable_Catmendccstatuschpvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmendccstatuschpvclrangeentry.EntityData.YFilter = catmendccstatuschpvclrangeentry.YFilter
    catmendccstatuschpvclrangeentry.EntityData.YangName = "catmEndCCStatusChPVclRangeEntry"
    catmendccstatuschpvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmendccstatuschpvclrangeentry.EntityData.ParentYangName = "catmEndCCStatusChPVclRangeTable"
    catmendccstatuschpvclrangeentry.EntityData.SegmentPath = "catmEndCCStatusChPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmendccstatuschpvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmendccstatuschpvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmendccstatuschpvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmendccstatuschpvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmendccstatuschpvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmendccstatuschpvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmendccstatuschpvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmendccstatuschpvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmendccstatuschpvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmendccstatuschpvclrangeentry.Ifindex}
    catmendccstatuschpvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmendccstatuschpvclrangeentry.Atmvclvpi}
    catmendccstatuschpvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmendccstatuschpvclrangeentry.Catmstatuschangepvclrangeindex}
    catmendccstatuschpvclrangeentry.EntityData.Leafs["catmPVclEndCCLowerRangeValue"] = types.YLeaf{"Catmpvclendcclowerrangevalue", catmendccstatuschpvclrangeentry.Catmpvclendcclowerrangevalue}
    catmendccstatuschpvclrangeentry.EntityData.Leafs["catmPVclEndCCHigherRangeValue"] = types.YLeaf{"Catmpvclendcchigherrangevalue", catmendccstatuschpvclrangeentry.Catmpvclendcchigherrangevalue}
    catmendccstatuschpvclrangeentry.EntityData.Leafs["catmPVclEndCCRangeStatusChStart"] = types.YLeaf{"Catmpvclendccrangestatuschstart", catmendccstatuschpvclrangeentry.Catmpvclendccrangestatuschstart}
    catmendccstatuschpvclrangeentry.EntityData.Leafs["catmPVclEndCCRangeStatusChEnd"] = types.YLeaf{"Catmpvclendccrangestatuschend", catmendccstatuschpvclrangeentry.Catmpvclendccrangestatuschend}
    return &(catmendccstatuschpvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to AIS/RDI failure
// in the same direction in the last corresponding PVC 
// notification.
type CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to AIS/RDI failure in the same
    // direction in the  last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry.
    Catmaisrdistatuschpvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry
}

func (catmaisrdistatuschpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable) GetEntityData() *types.CommonEntityData {
    catmaisrdistatuschpvclrangetable.EntityData.YFilter = catmaisrdistatuschpvclrangetable.YFilter
    catmaisrdistatuschpvclrangetable.EntityData.YangName = "catmAISRDIStatusChPVclRangeTable"
    catmaisrdistatuschpvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmaisrdistatuschpvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmaisrdistatuschpvclrangetable.EntityData.SegmentPath = "catmAISRDIStatusChPVclRangeTable"
    catmaisrdistatuschpvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmaisrdistatuschpvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmaisrdistatuschpvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmaisrdistatuschpvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmaisrdistatuschpvclrangetable.EntityData.Children["catmAISRDIStatusChPVclRangeEntry"] = types.YChild{"Catmaisrdistatuschpvclrangeentry", nil}
    for i := range catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry {
        catmaisrdistatuschpvclrangetable.EntityData.Children[types.GetSegmentPath(&catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry[i])] = types.YChild{"Catmaisrdistatuschpvclrangeentry", &catmaisrdistatuschpvclrangetable.Catmaisrdistatuschpvclrangeentry[i]}
    }
    catmaisrdistatuschpvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmaisrdistatuschpvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to AIS/RDI failure in the same direction in the 
// last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmaisrdistatuschpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatuschpvclrangetable_Catmaisrdistatuschpvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmaisrdistatuschpvclrangeentry.EntityData.YFilter = catmaisrdistatuschpvclrangeentry.YFilter
    catmaisrdistatuschpvclrangeentry.EntityData.YangName = "catmAISRDIStatusChPVclRangeEntry"
    catmaisrdistatuschpvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmaisrdistatuschpvclrangeentry.EntityData.ParentYangName = "catmAISRDIStatusChPVclRangeTable"
    catmaisrdistatuschpvclrangeentry.EntityData.SegmentPath = "catmAISRDIStatusChPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmaisrdistatuschpvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmaisrdistatuschpvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmaisrdistatuschpvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmaisrdistatuschpvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmaisrdistatuschpvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmaisrdistatuschpvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmaisrdistatuschpvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmaisrdistatuschpvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmaisrdistatuschpvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmaisrdistatuschpvclrangeentry.Ifindex}
    catmaisrdistatuschpvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmaisrdistatuschpvclrangeentry.Atmvclvpi}
    catmaisrdistatuschpvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmaisrdistatuschpvclrangeentry.Catmstatuschangepvclrangeindex}
    catmaisrdistatuschpvclrangeentry.EntityData.Leafs["catmPVclAISRDILowerRangeValue"] = types.YLeaf{"Catmpvclaisrdilowerrangevalue", catmaisrdistatuschpvclrangeentry.Catmpvclaisrdilowerrangevalue}
    catmaisrdistatuschpvclrangeentry.EntityData.Leafs["catmPVclAISRDIHigherRangeValue"] = types.YLeaf{"Catmpvclaisrdihigherrangevalue", catmaisrdistatuschpvclrangeentry.Catmpvclaisrdihigherrangevalue}
    catmaisrdistatuschpvclrangeentry.EntityData.Leafs["catmPVclAISRDIRangeStatusChStart"] = types.YLeaf{"Catmpvclaisrdirangestatuschstart", catmaisrdistatuschpvclrangeentry.Catmpvclaisrdirangestatuschstart}
    catmaisrdistatuschpvclrangeentry.EntityData.Leafs["catmPVclAISRDIRangeStatusChEnd"] = types.YLeaf{"Catmpvclaisrdirangestatuschend", catmaisrdistatuschpvclrangeentry.Catmpvclaisrdirangestatuschend}
    return &(catmaisrdistatuschpvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have detected as Down
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  atmVclOperStatus to  have detected as Down in the last notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry.
    Catmdownpvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry
}

func (catmdownpvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable) GetEntityData() *types.CommonEntityData {
    catmdownpvclrangetable.EntityData.YFilter = catmdownpvclrangetable.YFilter
    catmdownpvclrangetable.EntityData.YangName = "catmDownPVclRangeTable"
    catmdownpvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmdownpvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmdownpvclrangetable.EntityData.SegmentPath = "catmDownPVclRangeTable"
    catmdownpvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmdownpvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmdownpvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmdownpvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmdownpvclrangetable.EntityData.Children["catmDownPVclRangeEntry"] = types.YChild{"Catmdownpvclrangeentry", nil}
    for i := range catmdownpvclrangetable.Catmdownpvclrangeentry {
        catmdownpvclrangetable.EntityData.Children[types.GetSegmentPath(&catmdownpvclrangetable.Catmdownpvclrangeentry[i])] = types.YChild{"Catmdownpvclrangeentry", &catmdownpvclrangetable.Catmdownpvclrangeentry[i]}
    }
    catmdownpvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmdownpvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  atmVclOperStatus to 
// have detected as Down in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmdownpvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmdownpvclrangetable_Catmdownpvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmdownpvclrangeentry.EntityData.YFilter = catmdownpvclrangeentry.YFilter
    catmdownpvclrangeentry.EntityData.YangName = "catmDownPVclRangeEntry"
    catmdownpvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmdownpvclrangeentry.EntityData.ParentYangName = "catmDownPVclRangeTable"
    catmdownpvclrangeentry.EntityData.SegmentPath = "catmDownPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmdownpvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmdownpvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmdownpvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmdownpvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmdownpvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmdownpvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmdownpvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmdownpvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmdownpvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmdownpvclrangeentry.Ifindex}
    catmdownpvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmdownpvclrangeentry.Atmvclvpi}
    catmdownpvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmdownpvclrangeentry.Catmstatuschangepvclrangeindex}
    catmdownpvclrangeentry.EntityData.Leafs["catmDownPVclLowerRangeValue"] = types.YLeaf{"Catmdownpvcllowerrangevalue", catmdownpvclrangeentry.Catmdownpvcllowerrangevalue}
    catmdownpvclrangeentry.EntityData.Leafs["catmDownPVclHigherRangeValue"] = types.YLeaf{"Catmdownpvclhigherrangevalue", catmdownpvclrangeentry.Catmdownpvclhigherrangevalue}
    catmdownpvclrangeentry.EntityData.Leafs["catmDownPVclRangeStart"] = types.YLeaf{"Catmdownpvclrangestart", catmdownpvclrangeentry.Catmdownpvclrangestart}
    catmdownpvclrangeentry.EntityData.Leafs["catmDownPVclRangeEnd"] = types.YLeaf{"Catmdownpvclrangeend", catmdownpvclrangeentry.Catmdownpvclrangeend}
    catmdownpvclrangeentry.EntityData.Leafs["catmPrevUpPVclRangeStart"] = types.YLeaf{"Catmprevuppvclrangestart", catmdownpvclrangeentry.Catmprevuppvclrangestart}
    catmdownpvclrangeentry.EntityData.Leafs["catmPrevUpPVclRangeEnd"] = types.YLeaf{"Catmprevuppvclrangeend", catmdownpvclrangeentry.Catmprevuppvclrangeend}
    catmdownpvclrangeentry.EntityData.Leafs["catmPVclRangeFailureReason"] = types.YLeaf{"Catmpvclrangefailurereason", catmdownpvclrangeentry.Catmpvclrangefailurereason}
    return &(catmdownpvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed to UP
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed to UP in the last PVC notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry.
    Catmcurstatusuppvclentry []CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry
}

func (catmcurstatusuppvcltable *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable) GetEntityData() *types.CommonEntityData {
    catmcurstatusuppvcltable.EntityData.YFilter = catmcurstatusuppvcltable.YFilter
    catmcurstatusuppvcltable.EntityData.YangName = "catmCurStatusUpPVclTable"
    catmcurstatusuppvcltable.EntityData.BundleName = "cisco_ios_xe"
    catmcurstatusuppvcltable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmcurstatusuppvcltable.EntityData.SegmentPath = "catmCurStatusUpPVclTable"
    catmcurstatusuppvcltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmcurstatusuppvcltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmcurstatusuppvcltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmcurstatusuppvcltable.EntityData.Children = make(map[string]types.YChild)
    catmcurstatusuppvcltable.EntityData.Children["catmCurStatusUpPVclEntry"] = types.YChild{"Catmcurstatusuppvclentry", nil}
    for i := range catmcurstatusuppvcltable.Catmcurstatusuppvclentry {
        catmcurstatusuppvcltable.EntityData.Children[types.GetSegmentPath(&catmcurstatusuppvcltable.Catmcurstatusuppvclentry[i])] = types.YChild{"Catmcurstatusuppvclentry", &catmcurstatusuppvcltable.Catmcurstatusuppvclentry[i]}
    }
    catmcurstatusuppvcltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmcurstatusuppvcltable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed to UP in the last PVC notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry struct {
    EntityData types.CommonEntityData
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

func (catmcurstatusuppvclentry *CISCOATMPVCTRAPEXTNMIB_Catmcurstatusuppvcltable_Catmcurstatusuppvclentry) GetEntityData() *types.CommonEntityData {
    catmcurstatusuppvclentry.EntityData.YFilter = catmcurstatusuppvclentry.YFilter
    catmcurstatusuppvclentry.EntityData.YangName = "catmCurStatusUpPVclEntry"
    catmcurstatusuppvclentry.EntityData.BundleName = "cisco_ios_xe"
    catmcurstatusuppvclentry.EntityData.ParentYangName = "catmCurStatusUpPVclTable"
    catmcurstatusuppvclentry.EntityData.SegmentPath = "catmCurStatusUpPVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmcurstatusuppvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmcurstatusuppvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", catmcurstatusuppvclentry.Atmvclvci) + "']"
    catmcurstatusuppvclentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmcurstatusuppvclentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmcurstatusuppvclentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmcurstatusuppvclentry.EntityData.Children = make(map[string]types.YChild)
    catmcurstatusuppvclentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmcurstatusuppvclentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmcurstatusuppvclentry.Ifindex}
    catmcurstatusuppvclentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmcurstatusuppvclentry.Atmvclvpi}
    catmcurstatusuppvclentry.EntityData.Leafs["atmVclVci"] = types.YLeaf{"Atmvclvci", catmcurstatusuppvclentry.Atmvclvci}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclStatusUpTransition"] = types.YLeaf{"Catmpvclstatusuptransition", catmcurstatusuppvclentry.Catmpvclstatusuptransition}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclStatusUpStart"] = types.YLeaf{"Catmpvclstatusupstart", catmcurstatusuppvclentry.Catmpvclstatusupstart}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclStatusUpEnd"] = types.YLeaf{"Catmpvclstatusupend", catmcurstatusuppvclentry.Catmpvclstatusupend}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclSegCCStatusUpTransition"] = types.YLeaf{"Catmpvclsegccstatusuptransition", catmcurstatusuppvclentry.Catmpvclsegccstatusuptransition}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclSegCCStatusUpStart"] = types.YLeaf{"Catmpvclsegccstatusupstart", catmcurstatusuppvclentry.Catmpvclsegccstatusupstart}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclSegCCStatusUpEnd"] = types.YLeaf{"Catmpvclsegccstatusupend", catmcurstatusuppvclentry.Catmpvclsegccstatusupend}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclEndCCStatusUpTransition"] = types.YLeaf{"Catmpvclendccstatusuptransition", catmcurstatusuppvclentry.Catmpvclendccstatusuptransition}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclEndCCStatusUpStart"] = types.YLeaf{"Catmpvclendccstatusupstart", catmcurstatusuppvclentry.Catmpvclendccstatusupstart}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclEndCCStatusUpEnd"] = types.YLeaf{"Catmpvclendccstatusupend", catmcurstatusuppvclentry.Catmpvclendccstatusupend}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclAISRDIStatusUpTransition"] = types.YLeaf{"Catmpvclaisrdistatusuptransition", catmcurstatusuppvclentry.Catmpvclaisrdistatusuptransition}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclAISRDIStatusUpStart"] = types.YLeaf{"Catmpvclaisrdistatusupstart", catmcurstatusuppvclentry.Catmpvclaisrdistatusupstart}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclAISRDIStatusUpEnd"] = types.YLeaf{"Catmpvclaisrdistatusupend", catmcurstatusuppvclentry.Catmpvclaisrdistatusupend}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclCurRecoverTime"] = types.YLeaf{"Catmpvclcurrecovertime", catmcurstatusuppvclentry.Catmpvclcurrecovertime}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclPrevFailTime"] = types.YLeaf{"Catmpvclprevfailtime", catmcurstatusuppvclentry.Catmpvclprevfailtime}
    catmcurstatusuppvclentry.EntityData.Leafs["catmPVclRecoveryReason"] = types.YLeaf{"Catmpvclrecoveryreason", catmcurstatusuppvclentry.Catmpvclrecoveryreason}
    return &(catmcurstatusuppvclentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and loopback OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  loopback OAM status to  have detected as recovered in the last
    // notification  interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry.
    Catmstatusuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry
}

func (catmstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable) GetEntityData() *types.CommonEntityData {
    catmstatusuppvclrangetable.EntityData.YFilter = catmstatusuppvclrangetable.YFilter
    catmstatusuppvclrangetable.EntityData.YangName = "catmStatusUpPVclRangeTable"
    catmstatusuppvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmstatusuppvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmstatusuppvclrangetable.EntityData.SegmentPath = "catmStatusUpPVclRangeTable"
    catmstatusuppvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmstatusuppvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmstatusuppvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmstatusuppvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmstatusuppvclrangetable.EntityData.Children["catmStatusUpPVclRangeEntry"] = types.YChild{"Catmstatusuppvclrangeentry", nil}
    for i := range catmstatusuppvclrangetable.Catmstatusuppvclrangeentry {
        catmstatusuppvclrangetable.EntityData.Children[types.GetSegmentPath(&catmstatusuppvclrangetable.Catmstatusuppvclrangeentry[i])] = types.YChild{"Catmstatusuppvclrangeentry", &catmstatusuppvclrangetable.Catmstatusuppvclrangeentry[i]}
    }
    catmstatusuppvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmstatusuppvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  loopback OAM status to 
// have detected as recovered in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmstatusuppvclrangetable_Catmstatusuppvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmstatusuppvclrangeentry.EntityData.YFilter = catmstatusuppvclrangeentry.YFilter
    catmstatusuppvclrangeentry.EntityData.YangName = "catmStatusUpPVclRangeEntry"
    catmstatusuppvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmstatusuppvclrangeentry.EntityData.ParentYangName = "catmStatusUpPVclRangeTable"
    catmstatusuppvclrangeentry.EntityData.SegmentPath = "catmStatusUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmstatusuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmstatusuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmstatusuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmstatusuppvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmstatusuppvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmstatusuppvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmstatusuppvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmstatusuppvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmstatusuppvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmstatusuppvclrangeentry.Ifindex}
    catmstatusuppvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmstatusuppvclrangeentry.Atmvclvpi}
    catmstatusuppvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmstatusuppvclrangeentry.Catmstatuschangepvclrangeindex}
    catmstatusuppvclrangeentry.EntityData.Leafs["catmPVclUpLowerRangeValue"] = types.YLeaf{"Catmpvcluplowerrangevalue", catmstatusuppvclrangeentry.Catmpvcluplowerrangevalue}
    catmstatusuppvclrangeentry.EntityData.Leafs["catmPVclUpHigherRangeValue"] = types.YLeaf{"Catmpvcluphigherrangevalue", catmstatusuppvclrangeentry.Catmpvcluphigherrangevalue}
    catmstatusuppvclrangeentry.EntityData.Leafs["catmPVclRangeStatusUpStart"] = types.YLeaf{"Catmpvclrangestatusupstart", catmstatusuppvclrangeentry.Catmpvclrangestatusupstart}
    catmstatusuppvclrangeentry.EntityData.Leafs["catmPVclRangeStatusUpEnd"] = types.YLeaf{"Catmpvclrangestatusupend", catmstatusuppvclrangeentry.Catmpvclrangestatusupend}
    return &(catmstatusuppvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and Segment CC OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and Segment CC OAM status to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry.
    Catmsegccstatusuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry
}

func (catmsegccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable) GetEntityData() *types.CommonEntityData {
    catmsegccstatusuppvclrangetable.EntityData.YFilter = catmsegccstatusuppvclrangetable.YFilter
    catmsegccstatusuppvclrangetable.EntityData.YangName = "catmSegCCStatusUpPVclRangeTable"
    catmsegccstatusuppvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmsegccstatusuppvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmsegccstatusuppvclrangetable.EntityData.SegmentPath = "catmSegCCStatusUpPVclRangeTable"
    catmsegccstatusuppvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmsegccstatusuppvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmsegccstatusuppvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmsegccstatusuppvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmsegccstatusuppvclrangetable.EntityData.Children["catmSegCCStatusUpPVclRangeEntry"] = types.YChild{"Catmsegccstatusuppvclrangeentry", nil}
    for i := range catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry {
        catmsegccstatusuppvclrangetable.EntityData.Children[types.GetSegmentPath(&catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry[i])] = types.YChild{"Catmsegccstatusuppvclrangeentry", &catmsegccstatusuppvclrangetable.Catmsegccstatusuppvclrangeentry[i]}
    }
    catmsegccstatusuppvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmsegccstatusuppvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and Segment CC OAM status to
// have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmsegccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmsegccstatusuppvclrangetable_Catmsegccstatusuppvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmsegccstatusuppvclrangeentry.EntityData.YFilter = catmsegccstatusuppvclrangeentry.YFilter
    catmsegccstatusuppvclrangeentry.EntityData.YangName = "catmSegCCStatusUpPVclRangeEntry"
    catmsegccstatusuppvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmsegccstatusuppvclrangeentry.EntityData.ParentYangName = "catmSegCCStatusUpPVclRangeTable"
    catmsegccstatusuppvclrangeentry.EntityData.SegmentPath = "catmSegCCStatusUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmsegccstatusuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmsegccstatusuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmsegccstatusuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmsegccstatusuppvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmsegccstatusuppvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmsegccstatusuppvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmsegccstatusuppvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmsegccstatusuppvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmsegccstatusuppvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmsegccstatusuppvclrangeentry.Ifindex}
    catmsegccstatusuppvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmsegccstatusuppvclrangeentry.Atmvclvpi}
    catmsegccstatusuppvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmsegccstatusuppvclrangeentry.Catmstatuschangepvclrangeindex}
    catmsegccstatusuppvclrangeentry.EntityData.Leafs["catmPVclSegCCUpLowerRangeValue"] = types.YLeaf{"Catmpvclsegccuplowerrangevalue", catmsegccstatusuppvclrangeentry.Catmpvclsegccuplowerrangevalue}
    catmsegccstatusuppvclrangeentry.EntityData.Leafs["catmPVclSegCCUpHigherRangeValue"] = types.YLeaf{"Catmpvclsegccuphigherrangevalue", catmsegccstatusuppvclrangeentry.Catmpvclsegccuphigherrangevalue}
    catmsegccstatusuppvclrangeentry.EntityData.Leafs["catmPVclSegCCRangeStatusUpStart"] = types.YLeaf{"Catmpvclsegccrangestatusupstart", catmsegccstatusuppvclrangeentry.Catmpvclsegccrangestatusupstart}
    catmsegccstatusuppvclrangeentry.EntityData.Leafs["catmPVclSegCCRangeStatusUpEnd"] = types.YLeaf{"Catmpvclsegccrangestatusupend", catmsegccstatusuppvclrangeentry.Catmpvclsegccrangestatusupend}
    return &(catmsegccstatusuppvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and End-to-End CC OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and End-to-End CC OAM status  to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry.
    Catmendccstatusuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry
}

func (catmendccstatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable) GetEntityData() *types.CommonEntityData {
    catmendccstatusuppvclrangetable.EntityData.YFilter = catmendccstatusuppvclrangetable.YFilter
    catmendccstatusuppvclrangetable.EntityData.YangName = "catmEndCCStatusUpPVclRangeTable"
    catmendccstatusuppvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmendccstatusuppvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmendccstatusuppvclrangetable.EntityData.SegmentPath = "catmEndCCStatusUpPVclRangeTable"
    catmendccstatusuppvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmendccstatusuppvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmendccstatusuppvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmendccstatusuppvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmendccstatusuppvclrangetable.EntityData.Children["catmEndCCStatusUpPVclRangeEntry"] = types.YChild{"Catmendccstatusuppvclrangeentry", nil}
    for i := range catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry {
        catmendccstatusuppvclrangetable.EntityData.Children[types.GetSegmentPath(&catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry[i])] = types.YChild{"Catmendccstatusuppvclrangeentry", &catmendccstatusuppvclrangetable.Catmendccstatusuppvclrangeentry[i]}
    }
    catmendccstatusuppvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmendccstatusuppvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and End-to-End CC OAM status 
// to have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmendccstatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmendccstatusuppvclrangetable_Catmendccstatusuppvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmendccstatusuppvclrangeentry.EntityData.YFilter = catmendccstatusuppvclrangeentry.YFilter
    catmendccstatusuppvclrangeentry.EntityData.YangName = "catmEndCCStatusUpPVclRangeEntry"
    catmendccstatusuppvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmendccstatusuppvclrangeentry.EntityData.ParentYangName = "catmEndCCStatusUpPVclRangeTable"
    catmendccstatusuppvclrangeentry.EntityData.SegmentPath = "catmEndCCStatusUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmendccstatusuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmendccstatusuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmendccstatusuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmendccstatusuppvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmendccstatusuppvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmendccstatusuppvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmendccstatusuppvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmendccstatusuppvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmendccstatusuppvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmendccstatusuppvclrangeentry.Ifindex}
    catmendccstatusuppvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmendccstatusuppvclrangeentry.Atmvclvpi}
    catmendccstatusuppvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmendccstatusuppvclrangeentry.Catmstatuschangepvclrangeindex}
    catmendccstatusuppvclrangeentry.EntityData.Leafs["catmPVclEndCCUpLowerRangeValue"] = types.YLeaf{"Catmpvclendccuplowerrangevalue", catmendccstatusuppvclrangeentry.Catmpvclendccuplowerrangevalue}
    catmendccstatusuppvclrangeentry.EntityData.Leafs["catmPVclEndCCUpHigherRangeValue"] = types.YLeaf{"Catmpvclendccuphigherrangevalue", catmendccstatusuppvclrangeentry.Catmpvclendccuphigherrangevalue}
    catmendccstatusuppvclrangeentry.EntityData.Leafs["catmPVclEndCCRangeStatusUpStart"] = types.YLeaf{"Catmpvclendccrangestatusupstart", catmendccstatusuppvclrangeentry.Catmpvclendccrangestatusupstart}
    catmendccstatusuppvclrangeentry.EntityData.Leafs["catmPVclEndCCRangeStatusUpEnd"] = types.YLeaf{"Catmpvclendccrangestatusupend", catmendccstatusuppvclrangeentry.Catmpvclendccrangestatusupend}
    return &(catmendccstatusuppvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and AISRDI OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and AISRDI OAM status  to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry.
    Catmaisrdistatusuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry
}

func (catmaisrdistatusuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable) GetEntityData() *types.CommonEntityData {
    catmaisrdistatusuppvclrangetable.EntityData.YFilter = catmaisrdistatusuppvclrangetable.YFilter
    catmaisrdistatusuppvclrangetable.EntityData.YangName = "catmAISRDIStatusUpPVclRangeTable"
    catmaisrdistatusuppvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmaisrdistatusuppvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmaisrdistatusuppvclrangetable.EntityData.SegmentPath = "catmAISRDIStatusUpPVclRangeTable"
    catmaisrdistatusuppvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmaisrdistatusuppvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmaisrdistatusuppvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmaisrdistatusuppvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmaisrdistatusuppvclrangetable.EntityData.Children["catmAISRDIStatusUpPVclRangeEntry"] = types.YChild{"Catmaisrdistatusuppvclrangeentry", nil}
    for i := range catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry {
        catmaisrdistatusuppvclrangetable.EntityData.Children[types.GetSegmentPath(&catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry[i])] = types.YChild{"Catmaisrdistatusuppvclrangeentry", &catmaisrdistatusuppvclrangetable.Catmaisrdistatusuppvclrangeentry[i]}
    }
    catmaisrdistatusuppvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmaisrdistatusuppvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and AISRDI OAM status 
// to have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmaisrdistatusuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmaisrdistatusuppvclrangetable_Catmaisrdistatusuppvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmaisrdistatusuppvclrangeentry.EntityData.YFilter = catmaisrdistatusuppvclrangeentry.YFilter
    catmaisrdistatusuppvclrangeentry.EntityData.YangName = "catmAISRDIStatusUpPVclRangeEntry"
    catmaisrdistatusuppvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmaisrdistatusuppvclrangeentry.EntityData.ParentYangName = "catmAISRDIStatusUpPVclRangeTable"
    catmaisrdistatusuppvclrangeentry.EntityData.SegmentPath = "catmAISRDIStatusUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmaisrdistatusuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmaisrdistatusuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmaisrdistatusuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmaisrdistatusuppvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmaisrdistatusuppvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmaisrdistatusuppvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmaisrdistatusuppvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmaisrdistatusuppvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmaisrdistatusuppvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmaisrdistatusuppvclrangeentry.Ifindex}
    catmaisrdistatusuppvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmaisrdistatusuppvclrangeentry.Atmvclvpi}
    catmaisrdistatusuppvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmaisrdistatusuppvclrangeentry.Catmstatuschangepvclrangeindex}
    catmaisrdistatusuppvclrangeentry.EntityData.Leafs["catmPVclAISRDIUpLowerRangeValue"] = types.YLeaf{"Catmpvclaisrdiuplowerrangevalue", catmaisrdistatusuppvclrangeentry.Catmpvclaisrdiuplowerrangevalue}
    catmaisrdistatusuppvclrangeentry.EntityData.Leafs["catmPVclAISRDIUpHigherRangeValue"] = types.YLeaf{"Catmpvclaisrdiuphigherrangevalue", catmaisrdistatusuppvclrangeentry.Catmpvclaisrdiuphigherrangevalue}
    catmaisrdistatusuppvclrangeentry.EntityData.Leafs["catmPVclAISRDIRangeStatusUpStart"] = types.YLeaf{"Catmpvclaisrdirangestatusupstart", catmaisrdistatusuppvclrangeentry.Catmpvclaisrdirangestatusupstart}
    catmaisrdistatusuppvclrangeentry.EntityData.Leafs["catmPVclAISRDIRangeStatusUpEnd"] = types.YLeaf{"Catmpvclaisrdirangestatusupend", catmaisrdistatusuppvclrangeentry.Catmpvclaisrdirangestatusupend}
    return &(catmaisrdistatusuppvclrangeentry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have detected as Up
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  atmVclOperStatus to  have detected as Up in the last notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry.
    Catmuppvclrangeentry []CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry
}

func (catmuppvclrangetable *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable) GetEntityData() *types.CommonEntityData {
    catmuppvclrangetable.EntityData.YFilter = catmuppvclrangetable.YFilter
    catmuppvclrangetable.EntityData.YangName = "catmUpPVclRangeTable"
    catmuppvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    catmuppvclrangetable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmuppvclrangetable.EntityData.SegmentPath = "catmUpPVclRangeTable"
    catmuppvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmuppvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmuppvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmuppvclrangetable.EntityData.Children = make(map[string]types.YChild)
    catmuppvclrangetable.EntityData.Children["catmUpPVclRangeEntry"] = types.YChild{"Catmuppvclrangeentry", nil}
    for i := range catmuppvclrangetable.Catmuppvclrangeentry {
        catmuppvclrangetable.EntityData.Children[types.GetSegmentPath(&catmuppvclrangetable.Catmuppvclrangeentry[i])] = types.YChild{"Catmuppvclrangeentry", &catmuppvclrangetable.Catmuppvclrangeentry[i]}
    }
    catmuppvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(catmuppvclrangetable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  atmVclOperStatus to 
// have detected as Up in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry struct {
    EntityData types.CommonEntityData
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

func (catmuppvclrangeentry *CISCOATMPVCTRAPEXTNMIB_Catmuppvclrangetable_Catmuppvclrangeentry) GetEntityData() *types.CommonEntityData {
    catmuppvclrangeentry.EntityData.YFilter = catmuppvclrangeentry.YFilter
    catmuppvclrangeentry.EntityData.YangName = "catmUpPVclRangeEntry"
    catmuppvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    catmuppvclrangeentry.EntityData.ParentYangName = "catmUpPVclRangeTable"
    catmuppvclrangeentry.EntityData.SegmentPath = "catmUpPVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", catmuppvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", catmuppvclrangeentry.Atmvclvpi) + "']" + "[catmStatusChangePVclRangeIndex='" + fmt.Sprintf("%v", catmuppvclrangeentry.Catmstatuschangepvclrangeindex) + "']"
    catmuppvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmuppvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmuppvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmuppvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    catmuppvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    catmuppvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", catmuppvclrangeentry.Ifindex}
    catmuppvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", catmuppvclrangeentry.Atmvclvpi}
    catmuppvclrangeentry.EntityData.Leafs["catmStatusChangePVclRangeIndex"] = types.YLeaf{"Catmstatuschangepvclrangeindex", catmuppvclrangeentry.Catmstatuschangepvclrangeindex}
    catmuppvclrangeentry.EntityData.Leafs["catmUpPVclLowerRangeValue"] = types.YLeaf{"Catmuppvcllowerrangevalue", catmuppvclrangeentry.Catmuppvcllowerrangevalue}
    catmuppvclrangeentry.EntityData.Leafs["catmUpPVclHigherRangeValue"] = types.YLeaf{"Catmuppvclhigherrangevalue", catmuppvclrangeentry.Catmuppvclhigherrangevalue}
    catmuppvclrangeentry.EntityData.Leafs["catmUpPVclRangeStart"] = types.YLeaf{"Catmuppvclrangestart", catmuppvclrangeentry.Catmuppvclrangestart}
    catmuppvclrangeentry.EntityData.Leafs["catmUpPVclRangeEnd"] = types.YLeaf{"Catmuppvclrangeend", catmuppvclrangeentry.Catmuppvclrangeend}
    catmuppvclrangeentry.EntityData.Leafs["catmPrevDownPVclRangeStart"] = types.YLeaf{"Catmprevdownpvclrangestart", catmuppvclrangeentry.Catmprevdownpvclrangestart}
    catmuppvclrangeentry.EntityData.Leafs["catmPrevDownPVclRangeEnd"] = types.YLeaf{"Catmprevdownpvclrangeend", catmuppvclrangeentry.Catmprevdownpvclrangeend}
    catmuppvclrangeentry.EntityData.Leafs["catmPVclRangeRecoveryReason"] = types.YLeaf{"Catmpvclrangerecoveryreason", catmuppvclrangeentry.Catmpvclrangerecoveryreason}
    return &(catmuppvclrangeentry.EntityData)
}

