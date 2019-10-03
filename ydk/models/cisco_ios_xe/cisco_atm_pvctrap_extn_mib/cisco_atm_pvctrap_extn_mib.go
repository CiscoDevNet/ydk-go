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
    CatmCurStatChangePVclTable CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed in the same direction
    // in the last corresponding PVC notification .
    CatmStatusChangePVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed due to segment CC 
    // failure in the same direction in the last PVC  corresponding notification .
    CatmSegCCStatusChPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed due to End CC failure
    // in the same direction in the last PVC notification  interval.
    CatmEndCCStatusChPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed due to AIS/RDI failure
    // in the same direction in the last corresponding PVC  notification.
    CatmAISRDIStatusChPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have detected as Down in the last
    // corresponding PVC notification .
    CatmDownPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable

    // A table indicating all VCLs for which there is an active row in the
    // atmVclTable having an atmVclConnKind value of `pvc' and atmVclOperStatus to
    // have changed to UP in the last corresponding PVC notification .
    CatmCurStatusUpPVclTable CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and loopback OAM status to have detected as recovered in the
    // last corresponding PVC notification .
    CatmStatusUpPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable

    // A table indicating more than one VCLs in a consecutive range and for each
    // VCL there is an active row in the atmVclTable having an atmVclConnKind
    // value of `pvc' and Segment CC OAM status to have detected as recovered in
    // the last corresponding PVC notification .
    CatmSegCCStatusUpPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable

    // A table indicating more than one VCLs in a consecutive range and for each
    // VCL there is an active row in the atmVclTable having an atmVclConnKind
    // value of `pvc' and End-to-End CC OAM status to have detected as recovered
    // in the last corresponding PVC notification .
    CatmEndCCStatusUpPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable

    // A table indicating more than one VCLs in a consecutive range and for each
    // VCL there is an active row in the atmVclTable having an atmVclConnKind
    // value of `pvc' and AISRDI OAM status to have detected as recovered in the
    // last corresponding PVC notification .
    CatmAISRDIStatusUpPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have detected as Up in the last
    // corresponding PVC notification .
    CatmUpPVclRangeTable CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable
}

func (cISCOATMPVCTRAPEXTNMIB *CISCOATMPVCTRAPEXTNMIB) GetEntityData() *types.CommonEntityData {
    cISCOATMPVCTRAPEXTNMIB.EntityData.YFilter = cISCOATMPVCTRAPEXTNMIB.YFilter
    cISCOATMPVCTRAPEXTNMIB.EntityData.YangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    cISCOATMPVCTRAPEXTNMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOATMPVCTRAPEXTNMIB.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    cISCOATMPVCTRAPEXTNMIB.EntityData.SegmentPath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB"
    cISCOATMPVCTRAPEXTNMIB.EntityData.AbsolutePath = cISCOATMPVCTRAPEXTNMIB.EntityData.SegmentPath
    cISCOATMPVCTRAPEXTNMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOATMPVCTRAPEXTNMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOATMPVCTRAPEXTNMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOATMPVCTRAPEXTNMIB.EntityData.Children = types.NewOrderedMap()
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmCurStatChangePVclTable", types.YChild{"CatmCurStatChangePVclTable", &cISCOATMPVCTRAPEXTNMIB.CatmCurStatChangePVclTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmStatusChangePVclRangeTable", types.YChild{"CatmStatusChangePVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmStatusChangePVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmSegCCStatusChPVclRangeTable", types.YChild{"CatmSegCCStatusChPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmSegCCStatusChPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmEndCCStatusChPVclRangeTable", types.YChild{"CatmEndCCStatusChPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmEndCCStatusChPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmAISRDIStatusChPVclRangeTable", types.YChild{"CatmAISRDIStatusChPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmAISRDIStatusChPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmDownPVclRangeTable", types.YChild{"CatmDownPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmDownPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmCurStatusUpPVclTable", types.YChild{"CatmCurStatusUpPVclTable", &cISCOATMPVCTRAPEXTNMIB.CatmCurStatusUpPVclTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmStatusUpPVclRangeTable", types.YChild{"CatmStatusUpPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmStatusUpPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmSegCCStatusUpPVclRangeTable", types.YChild{"CatmSegCCStatusUpPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmSegCCStatusUpPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmEndCCStatusUpPVclRangeTable", types.YChild{"CatmEndCCStatusUpPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmEndCCStatusUpPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmAISRDIStatusUpPVclRangeTable", types.YChild{"CatmAISRDIStatusUpPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmAISRDIStatusUpPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Children.Append("catmUpPVclRangeTable", types.YChild{"CatmUpPVclRangeTable", &cISCOATMPVCTRAPEXTNMIB.CatmUpPVclRangeTable})
    cISCOATMPVCTRAPEXTNMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOATMPVCTRAPEXTNMIB.EntityData.YListKeys = []string {}

    return &(cISCOATMPVCTRAPEXTNMIB.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed in the
// last corresponding PVC notification.
type CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed in the last corresponding PVC
    // notification. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable_CatmCurStatChangePVclEntry.
    CatmCurStatChangePVclEntry []*CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable_CatmCurStatChangePVclEntry
}

func (catmCurStatChangePVclTable *CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable) GetEntityData() *types.CommonEntityData {
    catmCurStatChangePVclTable.EntityData.YFilter = catmCurStatChangePVclTable.YFilter
    catmCurStatChangePVclTable.EntityData.YangName = "catmCurStatChangePVclTable"
    catmCurStatChangePVclTable.EntityData.BundleName = "cisco_ios_xe"
    catmCurStatChangePVclTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmCurStatChangePVclTable.EntityData.SegmentPath = "catmCurStatChangePVclTable"
    catmCurStatChangePVclTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmCurStatChangePVclTable.EntityData.SegmentPath
    catmCurStatChangePVclTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmCurStatChangePVclTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmCurStatChangePVclTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmCurStatChangePVclTable.EntityData.Children = types.NewOrderedMap()
    catmCurStatChangePVclTable.EntityData.Children.Append("catmCurStatChangePVclEntry", types.YChild{"CatmCurStatChangePVclEntry", nil})
    for i := range catmCurStatChangePVclTable.CatmCurStatChangePVclEntry {
        catmCurStatChangePVclTable.EntityData.Children.Append(types.GetSegmentPath(catmCurStatChangePVclTable.CatmCurStatChangePVclEntry[i]), types.YChild{"CatmCurStatChangePVclEntry", catmCurStatChangePVclTable.CatmCurStatChangePVclEntry[i]})
    }
    catmCurStatChangePVclTable.EntityData.Leafs = types.NewOrderedMap()

    catmCurStatChangePVclTable.EntityData.YListKeys = []string {}

    return &(catmCurStatChangePVclTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable_CatmCurStatChangePVclEntry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed in the last corresponding PVC notification.
type CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable_CatmCurStatChangePVclEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVci
    AtmVclVci interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // corresponding notification. The type is interface{} with range:
    // 0..4294967295.
    CatmPVclStatusTransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last corresponding notification. The type is interface{} with range:
    // 0..4294967295.
    CatmPVclStatusChangeStart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification. The type is interface{} with range:
    // 0..4294967295.
    CatmPVclStatusChangeEnd interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // corresponding notification due to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclSegCCStatusTransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last corresponding notification due to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclSegCCStatusChangeStart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification due to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclSegCCStatusChangeEnd interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // corresponding notification due to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclEndCCStatusTransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last corresponding notification due to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclEndCCStatusChangeStart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification due to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclEndCCStatusChangeEnd interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // corresponding notification due to AIS RDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclAISRDIStatusTransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last corresponding notification due to AIS RDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclAISRDIStatusChangeStart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification due to AIS RDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclAISRDIStatusChangeEnd interface{}

    // The time stamp at which this PVCL changed state to DOWN last time in the
    // last corresponding notification . The type is interface{} with range:
    // 0..4294967295.
    CatmPVclCurFailTime interface{}

    // The time stamp at which this PVCL changed state to UP last time in the
    // previous corresponding notification . The type is interface{} with range:
    // 0..4294967295.
    CatmPVclPrevRecoverTime interface{}

    // Type of OAM failure. The type is CatmOAMFailureType.
    CatmPVclFailureReason interface{}
}

func (catmCurStatChangePVclEntry *CISCOATMPVCTRAPEXTNMIB_CatmCurStatChangePVclTable_CatmCurStatChangePVclEntry) GetEntityData() *types.CommonEntityData {
    catmCurStatChangePVclEntry.EntityData.YFilter = catmCurStatChangePVclEntry.YFilter
    catmCurStatChangePVclEntry.EntityData.YangName = "catmCurStatChangePVclEntry"
    catmCurStatChangePVclEntry.EntityData.BundleName = "cisco_ios_xe"
    catmCurStatChangePVclEntry.EntityData.ParentYangName = "catmCurStatChangePVclTable"
    catmCurStatChangePVclEntry.EntityData.SegmentPath = "catmCurStatChangePVclEntry" + types.AddKeyToken(catmCurStatChangePVclEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmCurStatChangePVclEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmCurStatChangePVclEntry.AtmVclVci, "atmVclVci")
    catmCurStatChangePVclEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmCurStatChangePVclTable/" + catmCurStatChangePVclEntry.EntityData.SegmentPath
    catmCurStatChangePVclEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmCurStatChangePVclEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmCurStatChangePVclEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmCurStatChangePVclEntry.EntityData.Children = types.NewOrderedMap()
    catmCurStatChangePVclEntry.EntityData.Leafs = types.NewOrderedMap()
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmCurStatChangePVclEntry.IfIndex})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmCurStatChangePVclEntry.AtmVclVpi})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("atmVclVci", types.YLeaf{"AtmVclVci", catmCurStatChangePVclEntry.AtmVclVci})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclStatusTransition", types.YLeaf{"CatmPVclStatusTransition", catmCurStatChangePVclEntry.CatmPVclStatusTransition})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclStatusChangeStart", types.YLeaf{"CatmPVclStatusChangeStart", catmCurStatChangePVclEntry.CatmPVclStatusChangeStart})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclStatusChangeEnd", types.YLeaf{"CatmPVclStatusChangeEnd", catmCurStatChangePVclEntry.CatmPVclStatusChangeEnd})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclSegCCStatusTransition", types.YLeaf{"CatmPVclSegCCStatusTransition", catmCurStatChangePVclEntry.CatmPVclSegCCStatusTransition})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclSegCCStatusChangeStart", types.YLeaf{"CatmPVclSegCCStatusChangeStart", catmCurStatChangePVclEntry.CatmPVclSegCCStatusChangeStart})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclSegCCStatusChangeEnd", types.YLeaf{"CatmPVclSegCCStatusChangeEnd", catmCurStatChangePVclEntry.CatmPVclSegCCStatusChangeEnd})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclEndCCStatusTransition", types.YLeaf{"CatmPVclEndCCStatusTransition", catmCurStatChangePVclEntry.CatmPVclEndCCStatusTransition})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclEndCCStatusChangeStart", types.YLeaf{"CatmPVclEndCCStatusChangeStart", catmCurStatChangePVclEntry.CatmPVclEndCCStatusChangeStart})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclEndCCStatusChangeEnd", types.YLeaf{"CatmPVclEndCCStatusChangeEnd", catmCurStatChangePVclEntry.CatmPVclEndCCStatusChangeEnd})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclAISRDIStatusTransition", types.YLeaf{"CatmPVclAISRDIStatusTransition", catmCurStatChangePVclEntry.CatmPVclAISRDIStatusTransition})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclAISRDIStatusChangeStart", types.YLeaf{"CatmPVclAISRDIStatusChangeStart", catmCurStatChangePVclEntry.CatmPVclAISRDIStatusChangeStart})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclAISRDIStatusChangeEnd", types.YLeaf{"CatmPVclAISRDIStatusChangeEnd", catmCurStatChangePVclEntry.CatmPVclAISRDIStatusChangeEnd})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclCurFailTime", types.YLeaf{"CatmPVclCurFailTime", catmCurStatChangePVclEntry.CatmPVclCurFailTime})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclPrevRecoverTime", types.YLeaf{"CatmPVclPrevRecoverTime", catmCurStatChangePVclEntry.CatmPVclPrevRecoverTime})
    catmCurStatChangePVclEntry.EntityData.Leafs.Append("catmPVclFailureReason", types.YLeaf{"CatmPVclFailureReason", catmCurStatChangePVclEntry.CatmPVclFailureReason})

    catmCurStatChangePVclEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "AtmVclVci"}

    return &(catmCurStatChangePVclEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed in the same
// direction in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed in the same direction in the last
    // notification  interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry.
    CatmStatusChangePVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry
}

func (catmStatusChangePVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmStatusChangePVclRangeTable.EntityData.YFilter = catmStatusChangePVclRangeTable.YFilter
    catmStatusChangePVclRangeTable.EntityData.YangName = "catmStatusChangePVclRangeTable"
    catmStatusChangePVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmStatusChangePVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmStatusChangePVclRangeTable.EntityData.SegmentPath = "catmStatusChangePVclRangeTable"
    catmStatusChangePVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmStatusChangePVclRangeTable.EntityData.SegmentPath
    catmStatusChangePVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmStatusChangePVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmStatusChangePVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmStatusChangePVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmStatusChangePVclRangeTable.EntityData.Children.Append("catmStatusChangePVclRangeEntry", types.YChild{"CatmStatusChangePVclRangeEntry", nil})
    for i := range catmStatusChangePVclRangeTable.CatmStatusChangePVclRangeEntry {
        catmStatusChangePVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmStatusChangePVclRangeTable.CatmStatusChangePVclRangeEntry[i]), types.YChild{"CatmStatusChangePVclRangeEntry", catmStatusChangePVclRangeTable.CatmStatusChangePVclRangeEntry[i]})
    }
    catmStatusChangePVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmStatusChangePVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmStatusChangePVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed in the same direction in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. Index to represent different ranges, the first
    // range is indexed by value 0, the second by 1 and so on... The type is
    // interface{} with range: 0..65535.
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus to have
    // changed in the last  corresponding notification due to Loopback OAM
    // failure. The type is interface{} with range: 0..65536.
    CatmPVclLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the  atmOperStatus to have
    // changed in the last  corresponding notification due to Loopback OAM
    // failure. The type is interface{} with range: 0..65536.
    CatmPVclHigherRangeValue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last corresponding notification due  to Loopback OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclRangeStatusChangeStart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last corresponding notification due  to Loopback OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclRangeStatusChangeEnd interface{}
}

func (catmStatusChangePVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmStatusChangePVclRangeEntry.EntityData.YFilter = catmStatusChangePVclRangeEntry.YFilter
    catmStatusChangePVclRangeEntry.EntityData.YangName = "catmStatusChangePVclRangeEntry"
    catmStatusChangePVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmStatusChangePVclRangeEntry.EntityData.ParentYangName = "catmStatusChangePVclRangeTable"
    catmStatusChangePVclRangeEntry.EntityData.SegmentPath = "catmStatusChangePVclRangeEntry" + types.AddKeyToken(catmStatusChangePVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmStatusChangePVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmStatusChangePVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmStatusChangePVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmStatusChangePVclRangeTable/" + catmStatusChangePVclRangeEntry.EntityData.SegmentPath
    catmStatusChangePVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmStatusChangePVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmStatusChangePVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmStatusChangePVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmStatusChangePVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmStatusChangePVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmStatusChangePVclRangeEntry.IfIndex})
    catmStatusChangePVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmStatusChangePVclRangeEntry.AtmVclVpi})
    catmStatusChangePVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmStatusChangePVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmStatusChangePVclRangeEntry.EntityData.Leafs.Append("catmPVclLowerRangeValue", types.YLeaf{"CatmPVclLowerRangeValue", catmStatusChangePVclRangeEntry.CatmPVclLowerRangeValue})
    catmStatusChangePVclRangeEntry.EntityData.Leafs.Append("catmPVclHigherRangeValue", types.YLeaf{"CatmPVclHigherRangeValue", catmStatusChangePVclRangeEntry.CatmPVclHigherRangeValue})
    catmStatusChangePVclRangeEntry.EntityData.Leafs.Append("catmPVclRangeStatusChangeStart", types.YLeaf{"CatmPVclRangeStatusChangeStart", catmStatusChangePVclRangeEntry.CatmPVclRangeStatusChangeStart})
    catmStatusChangePVclRangeEntry.EntityData.Leafs.Append("catmPVclRangeStatusChangeEnd", types.YLeaf{"CatmPVclRangeStatusChangeEnd", catmStatusChangePVclRangeEntry.CatmPVclRangeStatusChangeEnd})

    catmStatusChangePVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmStatusChangePVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to segment CC 
// failure in the same direction in the last PVC 
// corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to segment CC failure in the same
    // direction  in the last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable_CatmSegCCStatusChPVclRangeEntry.
    CatmSegCCStatusChPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable_CatmSegCCStatusChPVclRangeEntry
}

func (catmSegCCStatusChPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmSegCCStatusChPVclRangeTable.EntityData.YFilter = catmSegCCStatusChPVclRangeTable.YFilter
    catmSegCCStatusChPVclRangeTable.EntityData.YangName = "catmSegCCStatusChPVclRangeTable"
    catmSegCCStatusChPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmSegCCStatusChPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmSegCCStatusChPVclRangeTable.EntityData.SegmentPath = "catmSegCCStatusChPVclRangeTable"
    catmSegCCStatusChPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmSegCCStatusChPVclRangeTable.EntityData.SegmentPath
    catmSegCCStatusChPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmSegCCStatusChPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmSegCCStatusChPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmSegCCStatusChPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmSegCCStatusChPVclRangeTable.EntityData.Children.Append("catmSegCCStatusChPVclRangeEntry", types.YChild{"CatmSegCCStatusChPVclRangeEntry", nil})
    for i := range catmSegCCStatusChPVclRangeTable.CatmSegCCStatusChPVclRangeEntry {
        catmSegCCStatusChPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmSegCCStatusChPVclRangeTable.CatmSegCCStatusChPVclRangeEntry[i]), types.YChild{"CatmSegCCStatusChPVclRangeEntry", catmSegCCStatusChPVclRangeTable.CatmSegCCStatusChPVclRangeEntry[i]})
    }
    catmSegCCStatusChPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmSegCCStatusChPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmSegCCStatusChPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable_CatmSegCCStatusChPVclRangeEntry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to segment CC failure in the same direction 
// in the last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable_CatmSegCCStatusChPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus to have
    // changed in the last  corresponding notification due to Segment CC OAM
    // failure. The type is interface{} with range: 0..65536.
    CatmPVclSegCCLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the  atmOperStatus to have
    // changed in the last  corresponding notification due to Segment CC OAM
    // failure. The type is interface{} with range: 0..65536.
    CatmPVclSegCCHigherRangeValue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last corresponding notification due  to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclSegCCRangeStatusChStart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last corresponding notification due  to Segment CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclSegCCRangeStatusChEnd interface{}
}

func (catmSegCCStatusChPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusChPVclRangeTable_CatmSegCCStatusChPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmSegCCStatusChPVclRangeEntry.EntityData.YFilter = catmSegCCStatusChPVclRangeEntry.YFilter
    catmSegCCStatusChPVclRangeEntry.EntityData.YangName = "catmSegCCStatusChPVclRangeEntry"
    catmSegCCStatusChPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmSegCCStatusChPVclRangeEntry.EntityData.ParentYangName = "catmSegCCStatusChPVclRangeTable"
    catmSegCCStatusChPVclRangeEntry.EntityData.SegmentPath = "catmSegCCStatusChPVclRangeEntry" + types.AddKeyToken(catmSegCCStatusChPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmSegCCStatusChPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmSegCCStatusChPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmSegCCStatusChPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmSegCCStatusChPVclRangeTable/" + catmSegCCStatusChPVclRangeEntry.EntityData.SegmentPath
    catmSegCCStatusChPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmSegCCStatusChPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmSegCCStatusChPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmSegCCStatusChPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmSegCCStatusChPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmSegCCStatusChPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmSegCCStatusChPVclRangeEntry.IfIndex})
    catmSegCCStatusChPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmSegCCStatusChPVclRangeEntry.AtmVclVpi})
    catmSegCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmSegCCStatusChPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmSegCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclSegCCLowerRangeValue", types.YLeaf{"CatmPVclSegCCLowerRangeValue", catmSegCCStatusChPVclRangeEntry.CatmPVclSegCCLowerRangeValue})
    catmSegCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclSegCCHigherRangeValue", types.YLeaf{"CatmPVclSegCCHigherRangeValue", catmSegCCStatusChPVclRangeEntry.CatmPVclSegCCHigherRangeValue})
    catmSegCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclSegCCRangeStatusChStart", types.YLeaf{"CatmPVclSegCCRangeStatusChStart", catmSegCCStatusChPVclRangeEntry.CatmPVclSegCCRangeStatusChStart})
    catmSegCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclSegCCRangeStatusChEnd", types.YLeaf{"CatmPVclSegCCRangeStatusChEnd", catmSegCCStatusChPVclRangeEntry.CatmPVclSegCCRangeStatusChEnd})

    catmSegCCStatusChPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmSegCCStatusChPVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to End CC failure
// in the same direction in the last PVC notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to End CC failure in the same
    // direction in the  last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable_CatmEndCCStatusChPVclRangeEntry.
    CatmEndCCStatusChPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable_CatmEndCCStatusChPVclRangeEntry
}

func (catmEndCCStatusChPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmEndCCStatusChPVclRangeTable.EntityData.YFilter = catmEndCCStatusChPVclRangeTable.YFilter
    catmEndCCStatusChPVclRangeTable.EntityData.YangName = "catmEndCCStatusChPVclRangeTable"
    catmEndCCStatusChPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmEndCCStatusChPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmEndCCStatusChPVclRangeTable.EntityData.SegmentPath = "catmEndCCStatusChPVclRangeTable"
    catmEndCCStatusChPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmEndCCStatusChPVclRangeTable.EntityData.SegmentPath
    catmEndCCStatusChPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmEndCCStatusChPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmEndCCStatusChPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmEndCCStatusChPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmEndCCStatusChPVclRangeTable.EntityData.Children.Append("catmEndCCStatusChPVclRangeEntry", types.YChild{"CatmEndCCStatusChPVclRangeEntry", nil})
    for i := range catmEndCCStatusChPVclRangeTable.CatmEndCCStatusChPVclRangeEntry {
        catmEndCCStatusChPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmEndCCStatusChPVclRangeTable.CatmEndCCStatusChPVclRangeEntry[i]), types.YChild{"CatmEndCCStatusChPVclRangeEntry", catmEndCCStatusChPVclRangeTable.CatmEndCCStatusChPVclRangeEntry[i]})
    }
    catmEndCCStatusChPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmEndCCStatusChPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmEndCCStatusChPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable_CatmEndCCStatusChPVclRangeEntry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to End CC failure in the same direction in the 
// last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable_CatmEndCCStatusChPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus to have
    // changed in the last  corresponding notification due to End CC OAM failure.
    // The type is interface{} with range: 0..65536.
    CatmPVclEndCCLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the  atmOperStatus to have
    // changed in the last  corresponding notification due to End CC OAM failure.
    // The type is interface{} with range: 0..65536.
    CatmPVclEndCCHigherRangeValue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last corresponding notification due  to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclEndCCRangeStatusChStart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last corresponding notification due  to End CC OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclEndCCRangeStatusChEnd interface{}
}

func (catmEndCCStatusChPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusChPVclRangeTable_CatmEndCCStatusChPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmEndCCStatusChPVclRangeEntry.EntityData.YFilter = catmEndCCStatusChPVclRangeEntry.YFilter
    catmEndCCStatusChPVclRangeEntry.EntityData.YangName = "catmEndCCStatusChPVclRangeEntry"
    catmEndCCStatusChPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmEndCCStatusChPVclRangeEntry.EntityData.ParentYangName = "catmEndCCStatusChPVclRangeTable"
    catmEndCCStatusChPVclRangeEntry.EntityData.SegmentPath = "catmEndCCStatusChPVclRangeEntry" + types.AddKeyToken(catmEndCCStatusChPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmEndCCStatusChPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmEndCCStatusChPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmEndCCStatusChPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmEndCCStatusChPVclRangeTable/" + catmEndCCStatusChPVclRangeEntry.EntityData.SegmentPath
    catmEndCCStatusChPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmEndCCStatusChPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmEndCCStatusChPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmEndCCStatusChPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmEndCCStatusChPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmEndCCStatusChPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmEndCCStatusChPVclRangeEntry.IfIndex})
    catmEndCCStatusChPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmEndCCStatusChPVclRangeEntry.AtmVclVpi})
    catmEndCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmEndCCStatusChPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmEndCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclEndCCLowerRangeValue", types.YLeaf{"CatmPVclEndCCLowerRangeValue", catmEndCCStatusChPVclRangeEntry.CatmPVclEndCCLowerRangeValue})
    catmEndCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclEndCCHigherRangeValue", types.YLeaf{"CatmPVclEndCCHigherRangeValue", catmEndCCStatusChPVclRangeEntry.CatmPVclEndCCHigherRangeValue})
    catmEndCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclEndCCRangeStatusChStart", types.YLeaf{"CatmPVclEndCCRangeStatusChStart", catmEndCCStatusChPVclRangeEntry.CatmPVclEndCCRangeStatusChStart})
    catmEndCCStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclEndCCRangeStatusChEnd", types.YLeaf{"CatmPVclEndCCRangeStatusChEnd", catmEndCCStatusChPVclRangeEntry.CatmPVclEndCCRangeStatusChEnd})

    catmEndCCStatusChPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmEndCCStatusChPVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed due to AIS/RDI failure
// in the same direction in the last corresponding PVC 
// notification.
type CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed due to AIS/RDI failure in the same
    // direction in the  last corresponding notification . The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable_CatmAISRDIStatusChPVclRangeEntry.
    CatmAISRDIStatusChPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable_CatmAISRDIStatusChPVclRangeEntry
}

func (catmAISRDIStatusChPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmAISRDIStatusChPVclRangeTable.EntityData.YFilter = catmAISRDIStatusChPVclRangeTable.YFilter
    catmAISRDIStatusChPVclRangeTable.EntityData.YangName = "catmAISRDIStatusChPVclRangeTable"
    catmAISRDIStatusChPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmAISRDIStatusChPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmAISRDIStatusChPVclRangeTable.EntityData.SegmentPath = "catmAISRDIStatusChPVclRangeTable"
    catmAISRDIStatusChPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmAISRDIStatusChPVclRangeTable.EntityData.SegmentPath
    catmAISRDIStatusChPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmAISRDIStatusChPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmAISRDIStatusChPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmAISRDIStatusChPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmAISRDIStatusChPVclRangeTable.EntityData.Children.Append("catmAISRDIStatusChPVclRangeEntry", types.YChild{"CatmAISRDIStatusChPVclRangeEntry", nil})
    for i := range catmAISRDIStatusChPVclRangeTable.CatmAISRDIStatusChPVclRangeEntry {
        catmAISRDIStatusChPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmAISRDIStatusChPVclRangeTable.CatmAISRDIStatusChPVclRangeEntry[i]), types.YChild{"CatmAISRDIStatusChPVclRangeEntry", catmAISRDIStatusChPVclRangeTable.CatmAISRDIStatusChPVclRangeEntry[i]})
    }
    catmAISRDIStatusChPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmAISRDIStatusChPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmAISRDIStatusChPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable_CatmAISRDIStatusChPVclRangeEntry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed due to AIS/RDI failure in the same direction in the 
// last corresponding notification .
type CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable_CatmAISRDIStatusChPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus to have
    // changed in the last  corresponding notification due to AISRDI OAM failure.
    // The type is interface{} with range: 0..65536.
    CatmPVclAISRDILowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the  atmOperStatus to have
    // changed in the last  corresponding notification due to AISRDI OAM failure.
    // The type is interface{} with range: 0..65536.
    CatmPVclAISRDIHigherRangeValue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last corresponding notification due  to AISRDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclAISRDIRangeStatusChStart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last corresponding notification due  to AISRDI OAM failure. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclAISRDIRangeStatusChEnd interface{}
}

func (catmAISRDIStatusChPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusChPVclRangeTable_CatmAISRDIStatusChPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmAISRDIStatusChPVclRangeEntry.EntityData.YFilter = catmAISRDIStatusChPVclRangeEntry.YFilter
    catmAISRDIStatusChPVclRangeEntry.EntityData.YangName = "catmAISRDIStatusChPVclRangeEntry"
    catmAISRDIStatusChPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmAISRDIStatusChPVclRangeEntry.EntityData.ParentYangName = "catmAISRDIStatusChPVclRangeTable"
    catmAISRDIStatusChPVclRangeEntry.EntityData.SegmentPath = "catmAISRDIStatusChPVclRangeEntry" + types.AddKeyToken(catmAISRDIStatusChPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmAISRDIStatusChPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmAISRDIStatusChPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmAISRDIStatusChPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmAISRDIStatusChPVclRangeTable/" + catmAISRDIStatusChPVclRangeEntry.EntityData.SegmentPath
    catmAISRDIStatusChPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmAISRDIStatusChPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmAISRDIStatusChPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmAISRDIStatusChPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmAISRDIStatusChPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmAISRDIStatusChPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmAISRDIStatusChPVclRangeEntry.IfIndex})
    catmAISRDIStatusChPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmAISRDIStatusChPVclRangeEntry.AtmVclVpi})
    catmAISRDIStatusChPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmAISRDIStatusChPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmAISRDIStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclAISRDILowerRangeValue", types.YLeaf{"CatmPVclAISRDILowerRangeValue", catmAISRDIStatusChPVclRangeEntry.CatmPVclAISRDILowerRangeValue})
    catmAISRDIStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclAISRDIHigherRangeValue", types.YLeaf{"CatmPVclAISRDIHigherRangeValue", catmAISRDIStatusChPVclRangeEntry.CatmPVclAISRDIHigherRangeValue})
    catmAISRDIStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclAISRDIRangeStatusChStart", types.YLeaf{"CatmPVclAISRDIRangeStatusChStart", catmAISRDIStatusChPVclRangeEntry.CatmPVclAISRDIRangeStatusChStart})
    catmAISRDIStatusChPVclRangeEntry.EntityData.Leafs.Append("catmPVclAISRDIRangeStatusChEnd", types.YLeaf{"CatmPVclAISRDIRangeStatusChEnd", catmAISRDIStatusChPVclRangeEntry.CatmPVclAISRDIRangeStatusChEnd})

    catmAISRDIStatusChPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmAISRDIStatusChPVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have detected as Down
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  atmVclOperStatus to  have detected as Down in the last notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable_CatmDownPVclRangeEntry.
    CatmDownPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable_CatmDownPVclRangeEntry
}

func (catmDownPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmDownPVclRangeTable.EntityData.YFilter = catmDownPVclRangeTable.YFilter
    catmDownPVclRangeTable.EntityData.YangName = "catmDownPVclRangeTable"
    catmDownPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmDownPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmDownPVclRangeTable.EntityData.SegmentPath = "catmDownPVclRangeTable"
    catmDownPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmDownPVclRangeTable.EntityData.SegmentPath
    catmDownPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmDownPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmDownPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmDownPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmDownPVclRangeTable.EntityData.Children.Append("catmDownPVclRangeEntry", types.YChild{"CatmDownPVclRangeEntry", nil})
    for i := range catmDownPVclRangeTable.CatmDownPVclRangeEntry {
        catmDownPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmDownPVclRangeTable.CatmDownPVclRangeEntry[i]), types.YChild{"CatmDownPVclRangeEntry", catmDownPVclRangeTable.CatmDownPVclRangeEntry[i]})
    }
    catmDownPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmDownPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmDownPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable_CatmDownPVclRangeEntry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  atmVclOperStatus to 
// have detected as Down in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable_CatmDownPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus has been
    // detected as Down in the  corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmDownPVclLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the  atmVclOperStatus has been
    // detected as Down in the  corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmDownPVclHigherRangeValue interface{}

    // The time stamp at which the first atmVclOperStatus is detected as Down on
    // any of the PVCLs in the range in the corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    CatmDownPVclRangeStart interface{}

    // The time stamp at which the last atmVclOperStatus is detected as Down on
    // any of the PVCLs in the range in the corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    CatmDownPVclRangeEnd interface{}

    // The time stamp at which the first atmVclOperStatus is detected as Up on any
    // of the PVCLs in the range in the previous catmIntfPvcUp2Trap notification.
    // The type is interface{} with range: 0..4294967295.
    CatmPrevUpPVclRangeStart interface{}

    // The time stamp at which the last atmVclOperStatus is detected as Up on any
    // of the PVCLs in the range in the previous catmIntfPvcUp2Trap notification.
    // The type is interface{} with range: 0..4294967295.
    CatmPrevUpPVclRangeEnd interface{}

    // Type of OAM failure. The type is CatmOAMFailureType.
    CatmPVclRangeFailureReason interface{}
}

func (catmDownPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmDownPVclRangeTable_CatmDownPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmDownPVclRangeEntry.EntityData.YFilter = catmDownPVclRangeEntry.YFilter
    catmDownPVclRangeEntry.EntityData.YangName = "catmDownPVclRangeEntry"
    catmDownPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmDownPVclRangeEntry.EntityData.ParentYangName = "catmDownPVclRangeTable"
    catmDownPVclRangeEntry.EntityData.SegmentPath = "catmDownPVclRangeEntry" + types.AddKeyToken(catmDownPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmDownPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmDownPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmDownPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmDownPVclRangeTable/" + catmDownPVclRangeEntry.EntityData.SegmentPath
    catmDownPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmDownPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmDownPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmDownPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmDownPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmDownPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmDownPVclRangeEntry.IfIndex})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmDownPVclRangeEntry.AtmVclVpi})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmDownPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("catmDownPVclLowerRangeValue", types.YLeaf{"CatmDownPVclLowerRangeValue", catmDownPVclRangeEntry.CatmDownPVclLowerRangeValue})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("catmDownPVclHigherRangeValue", types.YLeaf{"CatmDownPVclHigherRangeValue", catmDownPVclRangeEntry.CatmDownPVclHigherRangeValue})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("catmDownPVclRangeStart", types.YLeaf{"CatmDownPVclRangeStart", catmDownPVclRangeEntry.CatmDownPVclRangeStart})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("catmDownPVclRangeEnd", types.YLeaf{"CatmDownPVclRangeEnd", catmDownPVclRangeEntry.CatmDownPVclRangeEnd})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("catmPrevUpPVclRangeStart", types.YLeaf{"CatmPrevUpPVclRangeStart", catmDownPVclRangeEntry.CatmPrevUpPVclRangeStart})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("catmPrevUpPVclRangeEnd", types.YLeaf{"CatmPrevUpPVclRangeEnd", catmDownPVclRangeEntry.CatmPrevUpPVclRangeEnd})
    catmDownPVclRangeEntry.EntityData.Leafs.Append("catmPVclRangeFailureReason", types.YLeaf{"CatmPVclRangeFailureReason", catmDownPVclRangeEntry.CatmPVclRangeFailureReason})

    catmDownPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmDownPVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed to UP
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed to UP in the last PVC notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable_CatmCurStatusUpPVclEntry.
    CatmCurStatusUpPVclEntry []*CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable_CatmCurStatusUpPVclEntry
}

func (catmCurStatusUpPVclTable *CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable) GetEntityData() *types.CommonEntityData {
    catmCurStatusUpPVclTable.EntityData.YFilter = catmCurStatusUpPVclTable.YFilter
    catmCurStatusUpPVclTable.EntityData.YangName = "catmCurStatusUpPVclTable"
    catmCurStatusUpPVclTable.EntityData.BundleName = "cisco_ios_xe"
    catmCurStatusUpPVclTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmCurStatusUpPVclTable.EntityData.SegmentPath = "catmCurStatusUpPVclTable"
    catmCurStatusUpPVclTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmCurStatusUpPVclTable.EntityData.SegmentPath
    catmCurStatusUpPVclTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmCurStatusUpPVclTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmCurStatusUpPVclTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmCurStatusUpPVclTable.EntityData.Children = types.NewOrderedMap()
    catmCurStatusUpPVclTable.EntityData.Children.Append("catmCurStatusUpPVclEntry", types.YChild{"CatmCurStatusUpPVclEntry", nil})
    for i := range catmCurStatusUpPVclTable.CatmCurStatusUpPVclEntry {
        catmCurStatusUpPVclTable.EntityData.Children.Append(types.GetSegmentPath(catmCurStatusUpPVclTable.CatmCurStatusUpPVclEntry[i]), types.YChild{"CatmCurStatusUpPVclEntry", catmCurStatusUpPVclTable.CatmCurStatusUpPVclEntry[i]})
    }
    catmCurStatusUpPVclTable.EntityData.Leafs = types.NewOrderedMap()

    catmCurStatusUpPVclTable.EntityData.YListKeys = []string {}

    return &(catmCurStatusUpPVclTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable_CatmCurStatusUpPVclEntry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed to UP in the last PVC notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable_CatmCurStatusUpPVclEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVci
    AtmVclVci interface{}

    // The number of Down to Up state transitions due to OAM loopback recovery
    // that has happened on this PVCL  in the last corresponding notification .
    // The type is interface{} with range: 0..4294967295.
    CatmPVclStatusUpTransition interface{}

    // The time stamp at which this PVCL changed state to UP  for the first time
    // due to OAM loopback recovery in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    CatmPVclStatusUpStart interface{}

    // The time stamp at which this PVCL changed state to UP  for the last time
    // due to OAM loopback recovery in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    CatmPVclStatusUpEnd interface{}

    // The number of Down to Up state transitions that has  happened on this PVCL
    // in the last corresponding notification  due to Segment CC OAM recovery. The
    // type is interface{} with range: 0..4294967295.
    CatmPVclSegCCStatusUpTransition interface{}

    // The time stamp at which this PVCL changed state to Up for  the first time
    // in the last corresponding notification due to Segment CC OAM recovery. The
    // type is interface{} with range: 0..4294967295.
    CatmPVclSegCCStatusUpStart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // corresponding notification due to Segment CC  OAM recovery. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclSegCCStatusUpEnd interface{}

    // The number of Down to UP state transitions that has  happened on this PVCL
    // in the last notification  interval due to End CC OAM recovery. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclEndCCStatusUpTransition interface{}

    // The time stamp at which this PVCL changed state to Up for the first time in
    // the last corresponding notification  due to End CC OAM recovery. The type
    // is interface{} with range: 0..4294967295.
    CatmPVclEndCCStatusUpStart interface{}

    // The time stamp at which this PVCL changed state to Up for the last time in
    // the last corresponding notification  due to End CC OAM recovery. The type
    // is interface{} with range: 0..4294967295.
    CatmPVclEndCCStatusUpEnd interface{}

    // The number of Down to Up state transitions that  has happened on this PVCL
    // in the last notification  interval due to AIS RDI OAM recovery. The type is
    // interface{} with range: 0..4294967295.
    CatmPVclAISRDIStatusUpTransition interface{}

    // The time stamp at which this PVCL changed state to Up for the first time in
    // the last corresponding notification  due to AIS/RDI OAM recovery. The type
    // is interface{} with range: 0..4294967295.
    CatmPVclAISRDIStatusUpStart interface{}

    // The time stamp at which this PVCL changed state to Up for the last time in
    // the last corresponding notification  due to AIS/RDI OAM recovery. The type
    // is interface{} with range: 0..4294967295.
    CatmPVclAISRDIStatusUpEnd interface{}

    // The time stamp at which this PVCL changed state to UP last time in the last
    // corresponding notification . The type is interface{} with range:
    // 0..4294967295.
    CatmPVclCurRecoverTime interface{}

    // The time stamp at which this PVCL changed state to DOWN last time in the
    // previous corresponding notification . The type is interface{} with range:
    // 0..4294967295.
    CatmPVclPrevFailTime interface{}

    // Type of OAM Recovered. The type is CatmOAMRecoveryType.
    CatmPVclRecoveryReason interface{}
}

func (catmCurStatusUpPVclEntry *CISCOATMPVCTRAPEXTNMIB_CatmCurStatusUpPVclTable_CatmCurStatusUpPVclEntry) GetEntityData() *types.CommonEntityData {
    catmCurStatusUpPVclEntry.EntityData.YFilter = catmCurStatusUpPVclEntry.YFilter
    catmCurStatusUpPVclEntry.EntityData.YangName = "catmCurStatusUpPVclEntry"
    catmCurStatusUpPVclEntry.EntityData.BundleName = "cisco_ios_xe"
    catmCurStatusUpPVclEntry.EntityData.ParentYangName = "catmCurStatusUpPVclTable"
    catmCurStatusUpPVclEntry.EntityData.SegmentPath = "catmCurStatusUpPVclEntry" + types.AddKeyToken(catmCurStatusUpPVclEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmCurStatusUpPVclEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmCurStatusUpPVclEntry.AtmVclVci, "atmVclVci")
    catmCurStatusUpPVclEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmCurStatusUpPVclTable/" + catmCurStatusUpPVclEntry.EntityData.SegmentPath
    catmCurStatusUpPVclEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmCurStatusUpPVclEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmCurStatusUpPVclEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmCurStatusUpPVclEntry.EntityData.Children = types.NewOrderedMap()
    catmCurStatusUpPVclEntry.EntityData.Leafs = types.NewOrderedMap()
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmCurStatusUpPVclEntry.IfIndex})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmCurStatusUpPVclEntry.AtmVclVpi})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("atmVclVci", types.YLeaf{"AtmVclVci", catmCurStatusUpPVclEntry.AtmVclVci})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclStatusUpTransition", types.YLeaf{"CatmPVclStatusUpTransition", catmCurStatusUpPVclEntry.CatmPVclStatusUpTransition})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclStatusUpStart", types.YLeaf{"CatmPVclStatusUpStart", catmCurStatusUpPVclEntry.CatmPVclStatusUpStart})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclStatusUpEnd", types.YLeaf{"CatmPVclStatusUpEnd", catmCurStatusUpPVclEntry.CatmPVclStatusUpEnd})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclSegCCStatusUpTransition", types.YLeaf{"CatmPVclSegCCStatusUpTransition", catmCurStatusUpPVclEntry.CatmPVclSegCCStatusUpTransition})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclSegCCStatusUpStart", types.YLeaf{"CatmPVclSegCCStatusUpStart", catmCurStatusUpPVclEntry.CatmPVclSegCCStatusUpStart})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclSegCCStatusUpEnd", types.YLeaf{"CatmPVclSegCCStatusUpEnd", catmCurStatusUpPVclEntry.CatmPVclSegCCStatusUpEnd})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclEndCCStatusUpTransition", types.YLeaf{"CatmPVclEndCCStatusUpTransition", catmCurStatusUpPVclEntry.CatmPVclEndCCStatusUpTransition})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclEndCCStatusUpStart", types.YLeaf{"CatmPVclEndCCStatusUpStart", catmCurStatusUpPVclEntry.CatmPVclEndCCStatusUpStart})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclEndCCStatusUpEnd", types.YLeaf{"CatmPVclEndCCStatusUpEnd", catmCurStatusUpPVclEntry.CatmPVclEndCCStatusUpEnd})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclAISRDIStatusUpTransition", types.YLeaf{"CatmPVclAISRDIStatusUpTransition", catmCurStatusUpPVclEntry.CatmPVclAISRDIStatusUpTransition})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclAISRDIStatusUpStart", types.YLeaf{"CatmPVclAISRDIStatusUpStart", catmCurStatusUpPVclEntry.CatmPVclAISRDIStatusUpStart})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclAISRDIStatusUpEnd", types.YLeaf{"CatmPVclAISRDIStatusUpEnd", catmCurStatusUpPVclEntry.CatmPVclAISRDIStatusUpEnd})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclCurRecoverTime", types.YLeaf{"CatmPVclCurRecoverTime", catmCurStatusUpPVclEntry.CatmPVclCurRecoverTime})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclPrevFailTime", types.YLeaf{"CatmPVclPrevFailTime", catmCurStatusUpPVclEntry.CatmPVclPrevFailTime})
    catmCurStatusUpPVclEntry.EntityData.Leafs.Append("catmPVclRecoveryReason", types.YLeaf{"CatmPVclRecoveryReason", catmCurStatusUpPVclEntry.CatmPVclRecoveryReason})

    catmCurStatusUpPVclEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "AtmVclVci"}

    return &(catmCurStatusUpPVclEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and loopback OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  loopback OAM status to  have detected as recovered in the last
    // notification  interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable_CatmStatusUpPVclRangeEntry.
    CatmStatusUpPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable_CatmStatusUpPVclRangeEntry
}

func (catmStatusUpPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmStatusUpPVclRangeTable.EntityData.YFilter = catmStatusUpPVclRangeTable.YFilter
    catmStatusUpPVclRangeTable.EntityData.YangName = "catmStatusUpPVclRangeTable"
    catmStatusUpPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmStatusUpPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmStatusUpPVclRangeTable.EntityData.SegmentPath = "catmStatusUpPVclRangeTable"
    catmStatusUpPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmStatusUpPVclRangeTable.EntityData.SegmentPath
    catmStatusUpPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmStatusUpPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmStatusUpPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmStatusUpPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmStatusUpPVclRangeTable.EntityData.Children.Append("catmStatusUpPVclRangeEntry", types.YChild{"CatmStatusUpPVclRangeEntry", nil})
    for i := range catmStatusUpPVclRangeTable.CatmStatusUpPVclRangeEntry {
        catmStatusUpPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmStatusUpPVclRangeTable.CatmStatusUpPVclRangeEntry[i]), types.YChild{"CatmStatusUpPVclRangeEntry", catmStatusUpPVclRangeTable.CatmStatusUpPVclRangeEntry[i]})
    }
    catmStatusUpPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmStatusUpPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmStatusUpPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable_CatmStatusUpPVclRangeEntry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  loopback OAM status to 
// have detected as recovered in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable_CatmStatusUpPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the  Loopback OAM recovery has
    // been detected in the last  corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmPVclUpLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the  Loopback OAM recovery has
    // been detected in the last  corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmPVclUpHigherRangeValue interface{}

    // The time stamp at which the first Loopback OAM recovery is detected on any
    // of the PVCLs in the range in the last corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    CatmPVclRangeStatusUpStart interface{}

    // The time stamp at which the last Loopback OAM recovery is detected on any
    // of the PVCLs in the range in the last corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    CatmPVclRangeStatusUpEnd interface{}
}

func (catmStatusUpPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmStatusUpPVclRangeTable_CatmStatusUpPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmStatusUpPVclRangeEntry.EntityData.YFilter = catmStatusUpPVclRangeEntry.YFilter
    catmStatusUpPVclRangeEntry.EntityData.YangName = "catmStatusUpPVclRangeEntry"
    catmStatusUpPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmStatusUpPVclRangeEntry.EntityData.ParentYangName = "catmStatusUpPVclRangeTable"
    catmStatusUpPVclRangeEntry.EntityData.SegmentPath = "catmStatusUpPVclRangeEntry" + types.AddKeyToken(catmStatusUpPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmStatusUpPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmStatusUpPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmStatusUpPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmStatusUpPVclRangeTable/" + catmStatusUpPVclRangeEntry.EntityData.SegmentPath
    catmStatusUpPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmStatusUpPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmStatusUpPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmStatusUpPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmStatusUpPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmStatusUpPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmStatusUpPVclRangeEntry.IfIndex})
    catmStatusUpPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmStatusUpPVclRangeEntry.AtmVclVpi})
    catmStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmStatusUpPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclUpLowerRangeValue", types.YLeaf{"CatmPVclUpLowerRangeValue", catmStatusUpPVclRangeEntry.CatmPVclUpLowerRangeValue})
    catmStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclUpHigherRangeValue", types.YLeaf{"CatmPVclUpHigherRangeValue", catmStatusUpPVclRangeEntry.CatmPVclUpHigherRangeValue})
    catmStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclRangeStatusUpStart", types.YLeaf{"CatmPVclRangeStatusUpStart", catmStatusUpPVclRangeEntry.CatmPVclRangeStatusUpStart})
    catmStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclRangeStatusUpEnd", types.YLeaf{"CatmPVclRangeStatusUpEnd", catmStatusUpPVclRangeEntry.CatmPVclRangeStatusUpEnd})

    catmStatusUpPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmStatusUpPVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and Segment CC OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and Segment CC OAM status to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable_CatmSegCCStatusUpPVclRangeEntry.
    CatmSegCCStatusUpPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable_CatmSegCCStatusUpPVclRangeEntry
}

func (catmSegCCStatusUpPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmSegCCStatusUpPVclRangeTable.EntityData.YFilter = catmSegCCStatusUpPVclRangeTable.YFilter
    catmSegCCStatusUpPVclRangeTable.EntityData.YangName = "catmSegCCStatusUpPVclRangeTable"
    catmSegCCStatusUpPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmSegCCStatusUpPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmSegCCStatusUpPVclRangeTable.EntityData.SegmentPath = "catmSegCCStatusUpPVclRangeTable"
    catmSegCCStatusUpPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmSegCCStatusUpPVclRangeTable.EntityData.SegmentPath
    catmSegCCStatusUpPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmSegCCStatusUpPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmSegCCStatusUpPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmSegCCStatusUpPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmSegCCStatusUpPVclRangeTable.EntityData.Children.Append("catmSegCCStatusUpPVclRangeEntry", types.YChild{"CatmSegCCStatusUpPVclRangeEntry", nil})
    for i := range catmSegCCStatusUpPVclRangeTable.CatmSegCCStatusUpPVclRangeEntry {
        catmSegCCStatusUpPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmSegCCStatusUpPVclRangeTable.CatmSegCCStatusUpPVclRangeEntry[i]), types.YChild{"CatmSegCCStatusUpPVclRangeEntry", catmSegCCStatusUpPVclRangeTable.CatmSegCCStatusUpPVclRangeEntry[i]})
    }
    catmSegCCStatusUpPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmSegCCStatusUpPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmSegCCStatusUpPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable_CatmSegCCStatusUpPVclRangeEntry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and Segment CC OAM status to
// have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable_CatmSegCCStatusUpPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the Segment CC OAM recovery
    // has been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmPVclSegCCUpLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the Segment CC OAM recovery has
    // been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmPVclSegCCUpHigherRangeValue interface{}

    // The time stamp at which the first Segment CC OAM recovery is detected on
    // any of the PVCLs in the range in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    CatmPVclSegCCRangeStatusUpStart interface{}

    // The time stamp at which the last Segment CC OAM recovery is detected on any
    // of the PVCLs in the range in the last corresponding notification . The type
    // is interface{} with range: 0..4294967295.
    CatmPVclSegCCRangeStatusUpEnd interface{}
}

func (catmSegCCStatusUpPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmSegCCStatusUpPVclRangeTable_CatmSegCCStatusUpPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmSegCCStatusUpPVclRangeEntry.EntityData.YFilter = catmSegCCStatusUpPVclRangeEntry.YFilter
    catmSegCCStatusUpPVclRangeEntry.EntityData.YangName = "catmSegCCStatusUpPVclRangeEntry"
    catmSegCCStatusUpPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmSegCCStatusUpPVclRangeEntry.EntityData.ParentYangName = "catmSegCCStatusUpPVclRangeTable"
    catmSegCCStatusUpPVclRangeEntry.EntityData.SegmentPath = "catmSegCCStatusUpPVclRangeEntry" + types.AddKeyToken(catmSegCCStatusUpPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmSegCCStatusUpPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmSegCCStatusUpPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmSegCCStatusUpPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmSegCCStatusUpPVclRangeTable/" + catmSegCCStatusUpPVclRangeEntry.EntityData.SegmentPath
    catmSegCCStatusUpPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmSegCCStatusUpPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmSegCCStatusUpPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmSegCCStatusUpPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmSegCCStatusUpPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmSegCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmSegCCStatusUpPVclRangeEntry.IfIndex})
    catmSegCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmSegCCStatusUpPVclRangeEntry.AtmVclVpi})
    catmSegCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmSegCCStatusUpPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmSegCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclSegCCUpLowerRangeValue", types.YLeaf{"CatmPVclSegCCUpLowerRangeValue", catmSegCCStatusUpPVclRangeEntry.CatmPVclSegCCUpLowerRangeValue})
    catmSegCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclSegCCUpHigherRangeValue", types.YLeaf{"CatmPVclSegCCUpHigherRangeValue", catmSegCCStatusUpPVclRangeEntry.CatmPVclSegCCUpHigherRangeValue})
    catmSegCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclSegCCRangeStatusUpStart", types.YLeaf{"CatmPVclSegCCRangeStatusUpStart", catmSegCCStatusUpPVclRangeEntry.CatmPVclSegCCRangeStatusUpStart})
    catmSegCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclSegCCRangeStatusUpEnd", types.YLeaf{"CatmPVclSegCCRangeStatusUpEnd", catmSegCCStatusUpPVclRangeEntry.CatmPVclSegCCRangeStatusUpEnd})

    catmSegCCStatusUpPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmSegCCStatusUpPVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and End-to-End CC OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and End-to-End CC OAM status  to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable_CatmEndCCStatusUpPVclRangeEntry.
    CatmEndCCStatusUpPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable_CatmEndCCStatusUpPVclRangeEntry
}

func (catmEndCCStatusUpPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmEndCCStatusUpPVclRangeTable.EntityData.YFilter = catmEndCCStatusUpPVclRangeTable.YFilter
    catmEndCCStatusUpPVclRangeTable.EntityData.YangName = "catmEndCCStatusUpPVclRangeTable"
    catmEndCCStatusUpPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmEndCCStatusUpPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmEndCCStatusUpPVclRangeTable.EntityData.SegmentPath = "catmEndCCStatusUpPVclRangeTable"
    catmEndCCStatusUpPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmEndCCStatusUpPVclRangeTable.EntityData.SegmentPath
    catmEndCCStatusUpPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmEndCCStatusUpPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmEndCCStatusUpPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmEndCCStatusUpPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmEndCCStatusUpPVclRangeTable.EntityData.Children.Append("catmEndCCStatusUpPVclRangeEntry", types.YChild{"CatmEndCCStatusUpPVclRangeEntry", nil})
    for i := range catmEndCCStatusUpPVclRangeTable.CatmEndCCStatusUpPVclRangeEntry {
        catmEndCCStatusUpPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmEndCCStatusUpPVclRangeTable.CatmEndCCStatusUpPVclRangeEntry[i]), types.YChild{"CatmEndCCStatusUpPVclRangeEntry", catmEndCCStatusUpPVclRangeTable.CatmEndCCStatusUpPVclRangeEntry[i]})
    }
    catmEndCCStatusUpPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmEndCCStatusUpPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmEndCCStatusUpPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable_CatmEndCCStatusUpPVclRangeEntry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and End-to-End CC OAM status 
// to have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable_CatmEndCCStatusUpPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the End-to-End CC OAM recovery
    // has been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmPVclEndCCUpLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the End-to-End CC OAM recovery
    // has been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmPVclEndCCUpHigherRangeValue interface{}

    // The time stamp at which the first End-to-End CC OAM recovery is detected on
    // any of the PVCLs in the range in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    CatmPVclEndCCRangeStatusUpStart interface{}

    // The time stamp at which the last End-to-End CC OAM recovery is detected on
    // any of the PVCLs in the range in the last corresponding notification . The
    // type is interface{} with range: 0..4294967295.
    CatmPVclEndCCRangeStatusUpEnd interface{}
}

func (catmEndCCStatusUpPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmEndCCStatusUpPVclRangeTable_CatmEndCCStatusUpPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmEndCCStatusUpPVclRangeEntry.EntityData.YFilter = catmEndCCStatusUpPVclRangeEntry.YFilter
    catmEndCCStatusUpPVclRangeEntry.EntityData.YangName = "catmEndCCStatusUpPVclRangeEntry"
    catmEndCCStatusUpPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmEndCCStatusUpPVclRangeEntry.EntityData.ParentYangName = "catmEndCCStatusUpPVclRangeTable"
    catmEndCCStatusUpPVclRangeEntry.EntityData.SegmentPath = "catmEndCCStatusUpPVclRangeEntry" + types.AddKeyToken(catmEndCCStatusUpPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmEndCCStatusUpPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmEndCCStatusUpPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmEndCCStatusUpPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmEndCCStatusUpPVclRangeTable/" + catmEndCCStatusUpPVclRangeEntry.EntityData.SegmentPath
    catmEndCCStatusUpPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmEndCCStatusUpPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmEndCCStatusUpPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmEndCCStatusUpPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmEndCCStatusUpPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmEndCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmEndCCStatusUpPVclRangeEntry.IfIndex})
    catmEndCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmEndCCStatusUpPVclRangeEntry.AtmVclVpi})
    catmEndCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmEndCCStatusUpPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmEndCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclEndCCUpLowerRangeValue", types.YLeaf{"CatmPVclEndCCUpLowerRangeValue", catmEndCCStatusUpPVclRangeEntry.CatmPVclEndCCUpLowerRangeValue})
    catmEndCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclEndCCUpHigherRangeValue", types.YLeaf{"CatmPVclEndCCUpHigherRangeValue", catmEndCCStatusUpPVclRangeEntry.CatmPVclEndCCUpHigherRangeValue})
    catmEndCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclEndCCRangeStatusUpStart", types.YLeaf{"CatmPVclEndCCRangeStatusUpStart", catmEndCCStatusUpPVclRangeEntry.CatmPVclEndCCRangeStatusUpStart})
    catmEndCCStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclEndCCRangeStatusUpEnd", types.YLeaf{"CatmPVclEndCCRangeStatusUpEnd", catmEndCCStatusUpPVclRangeEntry.CatmPVclEndCCRangeStatusUpEnd})

    catmEndCCStatusUpPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmEndCCStatusUpPVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable
// A table indicating more than one VCLs in a consecutive
// range and for each VCL there is an active row in the
// atmVclTable having an atmVclConnKind value of `pvc'
// and AISRDI OAM status to have detected as recovered
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and AISRDI OAM status  to have detected as recovered in the last
    // notification interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable_CatmAISRDIStatusUpPVclRangeEntry.
    CatmAISRDIStatusUpPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable_CatmAISRDIStatusUpPVclRangeEntry
}

func (catmAISRDIStatusUpPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmAISRDIStatusUpPVclRangeTable.EntityData.YFilter = catmAISRDIStatusUpPVclRangeTable.YFilter
    catmAISRDIStatusUpPVclRangeTable.EntityData.YangName = "catmAISRDIStatusUpPVclRangeTable"
    catmAISRDIStatusUpPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmAISRDIStatusUpPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmAISRDIStatusUpPVclRangeTable.EntityData.SegmentPath = "catmAISRDIStatusUpPVclRangeTable"
    catmAISRDIStatusUpPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmAISRDIStatusUpPVclRangeTable.EntityData.SegmentPath
    catmAISRDIStatusUpPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmAISRDIStatusUpPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmAISRDIStatusUpPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmAISRDIStatusUpPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmAISRDIStatusUpPVclRangeTable.EntityData.Children.Append("catmAISRDIStatusUpPVclRangeEntry", types.YChild{"CatmAISRDIStatusUpPVclRangeEntry", nil})
    for i := range catmAISRDIStatusUpPVclRangeTable.CatmAISRDIStatusUpPVclRangeEntry {
        catmAISRDIStatusUpPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmAISRDIStatusUpPVclRangeTable.CatmAISRDIStatusUpPVclRangeEntry[i]), types.YChild{"CatmAISRDIStatusUpPVclRangeEntry", catmAISRDIStatusUpPVclRangeTable.CatmAISRDIStatusUpPVclRangeEntry[i]})
    }
    catmAISRDIStatusUpPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmAISRDIStatusUpPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmAISRDIStatusUpPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable_CatmAISRDIStatusUpPVclRangeEntry
// Each entry in this table represents a range of VCLs and
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and AISRDI OAM status 
// to have detected as recovered in the last notification
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable_CatmAISRDIStatusUpPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the AISRDI OAM recovery has
    // been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmPVclAISRDIUpLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the AISRDI OAM recovery has
    // been detected in the last corresponding notification . The type is
    // interface{} with range: 0..65536.
    CatmPVclAISRDIUpHigherRangeValue interface{}

    // The time stamp at which the first AISRDI OAM recovery is detected on any of
    // the PVCLs in the range in the last corresponding notification . The type is
    // interface{} with range: 0..4294967295.
    CatmPVclAISRDIRangeStatusUpStart interface{}

    // The time stamp at which the last AISRDI OAM recovery is detected on any of
    // the PVCLs in the range in the last corresponding notification . The type is
    // interface{} with range: 0..4294967295.
    CatmPVclAISRDIRangeStatusUpEnd interface{}
}

func (catmAISRDIStatusUpPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmAISRDIStatusUpPVclRangeTable_CatmAISRDIStatusUpPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmAISRDIStatusUpPVclRangeEntry.EntityData.YFilter = catmAISRDIStatusUpPVclRangeEntry.YFilter
    catmAISRDIStatusUpPVclRangeEntry.EntityData.YangName = "catmAISRDIStatusUpPVclRangeEntry"
    catmAISRDIStatusUpPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmAISRDIStatusUpPVclRangeEntry.EntityData.ParentYangName = "catmAISRDIStatusUpPVclRangeTable"
    catmAISRDIStatusUpPVclRangeEntry.EntityData.SegmentPath = "catmAISRDIStatusUpPVclRangeEntry" + types.AddKeyToken(catmAISRDIStatusUpPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmAISRDIStatusUpPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmAISRDIStatusUpPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmAISRDIStatusUpPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmAISRDIStatusUpPVclRangeTable/" + catmAISRDIStatusUpPVclRangeEntry.EntityData.SegmentPath
    catmAISRDIStatusUpPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmAISRDIStatusUpPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmAISRDIStatusUpPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmAISRDIStatusUpPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmAISRDIStatusUpPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmAISRDIStatusUpPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmAISRDIStatusUpPVclRangeEntry.IfIndex})
    catmAISRDIStatusUpPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmAISRDIStatusUpPVclRangeEntry.AtmVclVpi})
    catmAISRDIStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmAISRDIStatusUpPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmAISRDIStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclAISRDIUpLowerRangeValue", types.YLeaf{"CatmPVclAISRDIUpLowerRangeValue", catmAISRDIStatusUpPVclRangeEntry.CatmPVclAISRDIUpLowerRangeValue})
    catmAISRDIStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclAISRDIUpHigherRangeValue", types.YLeaf{"CatmPVclAISRDIUpHigherRangeValue", catmAISRDIStatusUpPVclRangeEntry.CatmPVclAISRDIUpHigherRangeValue})
    catmAISRDIStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclAISRDIRangeStatusUpStart", types.YLeaf{"CatmPVclAISRDIRangeStatusUpStart", catmAISRDIStatusUpPVclRangeEntry.CatmPVclAISRDIRangeStatusUpStart})
    catmAISRDIStatusUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclAISRDIRangeStatusUpEnd", types.YLeaf{"CatmPVclAISRDIRangeStatusUpEnd", catmAISRDIStatusUpPVclRangeEntry.CatmPVclAISRDIRangeStatusUpEnd})

    catmAISRDIStatusUpPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmAISRDIStatusUpPVclRangeEntry.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have detected as Up
// in the last corresponding PVC notification .
type CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and  atmVclOperStatus to  have detected as Up in the last notification 
    // interval. The type is slice of
    // CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable_CatmUpPVclRangeEntry.
    CatmUpPVclRangeEntry []*CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable_CatmUpPVclRangeEntry
}

func (catmUpPVclRangeTable *CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable) GetEntityData() *types.CommonEntityData {
    catmUpPVclRangeTable.EntityData.YFilter = catmUpPVclRangeTable.YFilter
    catmUpPVclRangeTable.EntityData.YangName = "catmUpPVclRangeTable"
    catmUpPVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    catmUpPVclRangeTable.EntityData.ParentYangName = "CISCO-ATM-PVCTRAP-EXTN-MIB"
    catmUpPVclRangeTable.EntityData.SegmentPath = "catmUpPVclRangeTable"
    catmUpPVclRangeTable.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/" + catmUpPVclRangeTable.EntityData.SegmentPath
    catmUpPVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmUpPVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmUpPVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmUpPVclRangeTable.EntityData.Children = types.NewOrderedMap()
    catmUpPVclRangeTable.EntityData.Children.Append("catmUpPVclRangeEntry", types.YChild{"CatmUpPVclRangeEntry", nil})
    for i := range catmUpPVclRangeTable.CatmUpPVclRangeEntry {
        catmUpPVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(catmUpPVclRangeTable.CatmUpPVclRangeEntry[i]), types.YChild{"CatmUpPVclRangeEntry", catmUpPVclRangeTable.CatmUpPVclRangeEntry[i]})
    }
    catmUpPVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    catmUpPVclRangeTable.EntityData.YListKeys = []string {}

    return &(catmUpPVclRangeTable.EntityData)
}

// CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable_CatmUpPVclRangeEntry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and  atmVclOperStatus to 
// have detected as Up in the last notification 
// interval.
type CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable_CatmUpPVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // cisco_atm_pvctrap_extn_mib.CISCOATMPVCTRAPEXTNMIB_CatmStatusChangePVclRangeTable_CatmStatusChangePVclRangeEntry_CatmStatusChangePVclRangeIndex
    CatmStatusChangePVclRangeIndex interface{}

    // The first PVCL in a range of PVCLs for which the  atmVclOperStatus has been
    // detected as Up in the  corresponding notification . The type is interface{}
    // with range: 0..65536.
    CatmUpPVclLowerRangeValue interface{}

    // The last PVCL in a range of PVCLs for which the  atmVclOperStatus has been
    // detected as Up in the  corresponding notification . The type is interface{}
    // with range: 0..65536.
    CatmUpPVclHigherRangeValue interface{}

    // The time stamp at which the first atmVclOperStatus is detected as Up on any
    // of the PVCLs in the range in the corresponding notification . The type is
    // interface{} with range: 0..4294967295.
    CatmUpPVclRangeStart interface{}

    // The time stamp at which the last atmVclOperStatus is detected as Up on any
    // of the PVCLs in the range in the corresponding notification . The type is
    // interface{} with range: 0..4294967295.
    CatmUpPVclRangeEnd interface{}

    // The time stamp at which the first atmVclOperStatus is detected as Down on
    // any of the PVCLs in the range in the previous catmIntfPvcDownTrap
    // notification. The type is interface{} with range: 0..4294967295.
    CatmPrevDownPVclRangeStart interface{}

    // The time stamp at which the last atmVclOperStatus is detected as Down on
    // any of the PVCLs in the range in the previous catmIntfPvcDownTrap
    // notification. The type is interface{} with range: 0..4294967295.
    CatmPrevDownPVclRangeEnd interface{}

    // Type of OAM Recovered. The type is CatmOAMRecoveryType.
    CatmPVclRangeRecoveryReason interface{}
}

func (catmUpPVclRangeEntry *CISCOATMPVCTRAPEXTNMIB_CatmUpPVclRangeTable_CatmUpPVclRangeEntry) GetEntityData() *types.CommonEntityData {
    catmUpPVclRangeEntry.EntityData.YFilter = catmUpPVclRangeEntry.YFilter
    catmUpPVclRangeEntry.EntityData.YangName = "catmUpPVclRangeEntry"
    catmUpPVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    catmUpPVclRangeEntry.EntityData.ParentYangName = "catmUpPVclRangeTable"
    catmUpPVclRangeEntry.EntityData.SegmentPath = "catmUpPVclRangeEntry" + types.AddKeyToken(catmUpPVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(catmUpPVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(catmUpPVclRangeEntry.CatmStatusChangePVclRangeIndex, "catmStatusChangePVclRangeIndex")
    catmUpPVclRangeEntry.EntityData.AbsolutePath = "CISCO-ATM-PVCTRAP-EXTN-MIB:CISCO-ATM-PVCTRAP-EXTN-MIB/catmUpPVclRangeTable/" + catmUpPVclRangeEntry.EntityData.SegmentPath
    catmUpPVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    catmUpPVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    catmUpPVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    catmUpPVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    catmUpPVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    catmUpPVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", catmUpPVclRangeEntry.IfIndex})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", catmUpPVclRangeEntry.AtmVclVpi})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("catmStatusChangePVclRangeIndex", types.YLeaf{"CatmStatusChangePVclRangeIndex", catmUpPVclRangeEntry.CatmStatusChangePVclRangeIndex})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("catmUpPVclLowerRangeValue", types.YLeaf{"CatmUpPVclLowerRangeValue", catmUpPVclRangeEntry.CatmUpPVclLowerRangeValue})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("catmUpPVclHigherRangeValue", types.YLeaf{"CatmUpPVclHigherRangeValue", catmUpPVclRangeEntry.CatmUpPVclHigherRangeValue})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("catmUpPVclRangeStart", types.YLeaf{"CatmUpPVclRangeStart", catmUpPVclRangeEntry.CatmUpPVclRangeStart})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("catmUpPVclRangeEnd", types.YLeaf{"CatmUpPVclRangeEnd", catmUpPVclRangeEntry.CatmUpPVclRangeEnd})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("catmPrevDownPVclRangeStart", types.YLeaf{"CatmPrevDownPVclRangeStart", catmUpPVclRangeEntry.CatmPrevDownPVclRangeStart})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("catmPrevDownPVclRangeEnd", types.YLeaf{"CatmPrevDownPVclRangeEnd", catmUpPVclRangeEntry.CatmPrevDownPVclRangeEnd})
    catmUpPVclRangeEntry.EntityData.Leafs.Append("catmPVclRangeRecoveryReason", types.YLeaf{"CatmPVclRangeRecoveryReason", catmUpPVclRangeEntry.CatmPVclRangeRecoveryReason})

    catmUpPVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "CatmStatusChangePVclRangeIndex"}

    return &(catmUpPVclRangeEntry.EntityData)
}

