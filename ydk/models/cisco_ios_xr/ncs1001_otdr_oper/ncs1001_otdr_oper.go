// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1001-otdr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   hw-module: ncs1001 hw-module command chain
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ncs1001_otdr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1001_otdr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1001-otdr-oper hw-module}", reflect.TypeOf(HwModule{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1001-otdr-oper:hw-module", reflect.TypeOf(HwModule{}))
}

// Direction represents Direction
type Direction string

const (
    // Tx
    Direction_direction_tx Direction = "direction-tx"

    // Rx
    Direction_direction_rx Direction = "direction-rx"
)

// OtdrEvent represents Otdr event
type OtdrEvent string

const (
    // Loss
    OtdrEvent_otdr_event_type_loss OtdrEvent = "otdr-event-type-loss"

    // Reflection
    OtdrEvent_otdr_event_type_reflect_ion OtdrEvent = "otdr-event-type-reflect-ion"

    // Loss And Reflection
    OtdrEvent_otdr_event_type_loss_and_reflect_ion OtdrEvent = "otdr-event-type-loss-and-reflect-ion"

    // End Of Fiber
    OtdrEvent_otdr_event_type_end_of_fiber OtdrEvent = "otdr-event-type-end-of-fiber"
)

// OtsOtdrDataTypeDetails represents Ots otdr data type details
type OtsOtdrDataTypeDetails string

const (
    // Displays Scan Details
    OtsOtdrDataTypeDetails_scan_details OtsOtdrDataTypeDetails = "scan-details"

    // Displays Baseline Details
    OtsOtdrDataTypeDetails_base_line_details OtsOtdrDataTypeDetails = "base-line-details"
)

// OtdrScanMode represents Otdr scan mode
type OtdrScanMode string

const (
    // Auto
    OtdrScanMode_otdr_scan_auto OtdrScanMode = "otdr-scan-auto"

    // Expert
    OtdrScanMode_otdr_scan_expert OtdrScanMode = "otdr-scan-expert"
)

// OtdrStatus represents Otdr status
type OtdrStatus string

const (
    // Unknown
    OtdrStatus_otdr_status_unknown OtdrStatus = "otdr-status-unknown"

    // Ok
    OtdrStatus_otdr_status_ok OtdrStatus = "otdr-status-ok"

    // ORL Progress
    OtdrStatus_otdr_status_orl_progress OtdrStatus = "otdr-status-orl-progress"

    // OTDR Progress
    OtdrStatus_otdr_status_otdr_progress OtdrStatus = "otdr-status-otdr-progress"

    // Progress
    OtdrStatus_otdr_status_progress OtdrStatus = "otdr-status-progress"

    // Failure
    OtdrStatus_otdr_status_failure OtdrStatus = "otdr-status-failure"

    // Aborted
    OtdrStatus_otdr_status_aborted OtdrStatus = "otdr-status-aborted"
)

// OtsOtdrData represents Ots otdr data
type OtsOtdrData string

const (
    // Display otdr status
    OtsOtdrData_status OtsOtdrData = "status"

    // Display the list of saved measurements
    OtsOtdrData_scan OtsOtdrData = "scan"

    // Display the list of saved baseline
    OtsOtdrData_base_line OtsOtdrData = "base-line"
)

// HwModule
// ncs1001 hw-module command chain
type HwModule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Otdr Details Root Info.
    OtdrDetails HwModule_OtdrDetails

    // Otdr Root Info.
    Otdrs HwModule_Otdrs
}

func (hwModule *HwModule) GetEntityData() *types.CommonEntityData {
    hwModule.EntityData.YFilter = hwModule.YFilter
    hwModule.EntityData.YangName = "hw-module"
    hwModule.EntityData.BundleName = "cisco_ios_xr"
    hwModule.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1001-otdr-oper"
    hwModule.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module"
    hwModule.EntityData.AbsolutePath = hwModule.EntityData.SegmentPath
    hwModule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModule.EntityData.Children = types.NewOrderedMap()
    hwModule.EntityData.Children.Append("otdr-details", types.YChild{"OtdrDetails", &hwModule.OtdrDetails})
    hwModule.EntityData.Children.Append("otdrs", types.YChild{"Otdrs", &hwModule.Otdrs})
    hwModule.EntityData.Leafs = types.NewOrderedMap()

    hwModule.EntityData.YListKeys = []string {}

    return &(hwModule.EntityData)
}

// HwModule_OtdrDetails
// Otdr Details Root Info
type HwModule_OtdrDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Otdr Details. The type is slice of HwModule_OtdrDetails_OtdrDetail.
    OtdrDetail []*HwModule_OtdrDetails_OtdrDetail
}

func (otdrDetails *HwModule_OtdrDetails) GetEntityData() *types.CommonEntityData {
    otdrDetails.EntityData.YFilter = otdrDetails.YFilter
    otdrDetails.EntityData.YangName = "otdr-details"
    otdrDetails.EntityData.BundleName = "cisco_ios_xr"
    otdrDetails.EntityData.ParentYangName = "hw-module"
    otdrDetails.EntityData.SegmentPath = "otdr-details"
    otdrDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/" + otdrDetails.EntityData.SegmentPath
    otdrDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrDetails.EntityData.Children = types.NewOrderedMap()
    otdrDetails.EntityData.Children.Append("otdr-detail", types.YChild{"OtdrDetail", nil})
    for i := range otdrDetails.OtdrDetail {
        otdrDetails.EntityData.Children.Append(types.GetSegmentPath(otdrDetails.OtdrDetail[i]), types.YChild{"OtdrDetail", otdrDetails.OtdrDetail[i]})
    }
    otdrDetails.EntityData.Leafs = types.NewOrderedMap()

    otdrDetails.EntityData.YListKeys = []string {}

    return &(otdrDetails.EntityData)
}

// HwModule_OtdrDetails_OtdrDetail
// Otdr Details
type HwModule_OtdrDetails_OtdrDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Details associated with a particular slot number.
    // The type is interface{} with range: 1..3.
    SlotId interface{}

    // Dispaly data type list. The type is slice of
    // HwModule_OtdrDetails_OtdrDetail_DataTypeDetail.
    DataTypeDetail []*HwModule_OtdrDetails_OtdrDetail_DataTypeDetail
}

func (otdrDetail *HwModule_OtdrDetails_OtdrDetail) GetEntityData() *types.CommonEntityData {
    otdrDetail.EntityData.YFilter = otdrDetail.YFilter
    otdrDetail.EntityData.YangName = "otdr-detail"
    otdrDetail.EntityData.BundleName = "cisco_ios_xr"
    otdrDetail.EntityData.ParentYangName = "otdr-details"
    otdrDetail.EntityData.SegmentPath = "otdr-detail" + types.AddKeyToken(otdrDetail.SlotId, "slot-id")
    otdrDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdr-details/" + otdrDetail.EntityData.SegmentPath
    otdrDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrDetail.EntityData.Children = types.NewOrderedMap()
    otdrDetail.EntityData.Children.Append("data-type-detail", types.YChild{"DataTypeDetail", nil})
    for i := range otdrDetail.DataTypeDetail {
        otdrDetail.EntityData.Children.Append(types.GetSegmentPath(otdrDetail.DataTypeDetail[i]), types.YChild{"DataTypeDetail", otdrDetail.DataTypeDetail[i]})
    }
    otdrDetail.EntityData.Leafs = types.NewOrderedMap()
    otdrDetail.EntityData.Leafs.Append("slot-id", types.YLeaf{"SlotId", otdrDetail.SlotId})

    otdrDetail.EntityData.YListKeys = []string {"SlotId"}

    return &(otdrDetail.EntityData)
}

// HwModule_OtdrDetails_OtdrDetail_DataTypeDetail
// Dispaly data type list
type HwModule_OtdrDetails_OtdrDetail_DataTypeDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Details associated with a particular Data type.
    // The type is OtsOtdrDataTypeDetails.
    OtdrDataTypeDetails interface{}

    // Dispaly details. The type is slice of
    // HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier.
    Identifier []*HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier
}

func (dataTypeDetail *HwModule_OtdrDetails_OtdrDetail_DataTypeDetail) GetEntityData() *types.CommonEntityData {
    dataTypeDetail.EntityData.YFilter = dataTypeDetail.YFilter
    dataTypeDetail.EntityData.YangName = "data-type-detail"
    dataTypeDetail.EntityData.BundleName = "cisco_ios_xr"
    dataTypeDetail.EntityData.ParentYangName = "otdr-detail"
    dataTypeDetail.EntityData.SegmentPath = "data-type-detail" + types.AddKeyToken(dataTypeDetail.OtdrDataTypeDetails, "otdr-data-type-details")
    dataTypeDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdr-details/otdr-detail/" + dataTypeDetail.EntityData.SegmentPath
    dataTypeDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataTypeDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataTypeDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataTypeDetail.EntityData.Children = types.NewOrderedMap()
    dataTypeDetail.EntityData.Children.Append("identifier", types.YChild{"Identifier", nil})
    for i := range dataTypeDetail.Identifier {
        dataTypeDetail.EntityData.Children.Append(types.GetSegmentPath(dataTypeDetail.Identifier[i]), types.YChild{"Identifier", dataTypeDetail.Identifier[i]})
    }
    dataTypeDetail.EntityData.Leafs = types.NewOrderedMap()
    dataTypeDetail.EntityData.Leafs.Append("otdr-data-type-details", types.YLeaf{"OtdrDataTypeDetails", dataTypeDetail.OtdrDataTypeDetails})

    dataTypeDetail.EntityData.YListKeys = []string {"OtdrDataTypeDetails"}

    return &(dataTypeDetail.EntityData)
}

// HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier
// Dispaly details
type HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Dispaly details. The type is interface{} with
    // range: 0..4294967295.
    ScanDetail interface{}

    // otdr item.
    OtdrItem HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem

    // otdr scan. The type is slice of
    // HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrScan.
    OtdrScan []*HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrScan

    // otdr status. The type is slice of
    // HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrStatus.
    OtdrStatus []*HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrStatus
}

func (identifier *HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier) GetEntityData() *types.CommonEntityData {
    identifier.EntityData.YFilter = identifier.YFilter
    identifier.EntityData.YangName = "identifier"
    identifier.EntityData.BundleName = "cisco_ios_xr"
    identifier.EntityData.ParentYangName = "data-type-detail"
    identifier.EntityData.SegmentPath = "identifier" + types.AddKeyToken(identifier.ScanDetail, "scan-detail")
    identifier.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdr-details/otdr-detail/data-type-detail/" + identifier.EntityData.SegmentPath
    identifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    identifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    identifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    identifier.EntityData.Children = types.NewOrderedMap()
    identifier.EntityData.Children.Append("otdr-item", types.YChild{"OtdrItem", &identifier.OtdrItem})
    identifier.EntityData.Children.Append("otdr-scan", types.YChild{"OtdrScan", nil})
    for i := range identifier.OtdrScan {
        types.SetYListKey(identifier.OtdrScan[i], i)
        identifier.EntityData.Children.Append(types.GetSegmentPath(identifier.OtdrScan[i]), types.YChild{"OtdrScan", identifier.OtdrScan[i]})
    }
    identifier.EntityData.Children.Append("otdr-status", types.YChild{"OtdrStatus", nil})
    for i := range identifier.OtdrStatus {
        types.SetYListKey(identifier.OtdrStatus[i], i)
        identifier.EntityData.Children.Append(types.GetSegmentPath(identifier.OtdrStatus[i]), types.YChild{"OtdrStatus", identifier.OtdrStatus[i]})
    }
    identifier.EntityData.Leafs = types.NewOrderedMap()
    identifier.EntityData.Leafs.Append("scan-detail", types.YLeaf{"ScanDetail", identifier.ScanDetail})

    identifier.EntityData.YListKeys = []string {"ScanDetail"}

    return &(identifier.EntityData)
}

// HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem
// otdr item
type HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Otdr Number. The type is interface{} with range: 0..255.
    OtdrNumber interface{}

    // Scan Direction. The type is Direction.
    Direction interface{}

    // Timestamp. The type is string with length: 0..16.
    Timestamp interface{}

    // Sor file name. The type is string with length: 0..64.
    SorFile interface{}

    // Sor file location. The type is string with length: 0..64.
    SorDirectory interface{}

    // Scan Mode. The type is OtdrScanMode.
    ScanMode interface{}

    // Distance in Km. The type is string with length: 0..16.
    Distance interface{}

    // Total Orl in dB. The type is string with length: 0..16.
    TotalOrl interface{}

    // Event list. The type is slice of
    // HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem_OtdrEvent.
    OtdrEvent []*HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem_OtdrEvent
}

func (otdrItem *HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem) GetEntityData() *types.CommonEntityData {
    otdrItem.EntityData.YFilter = otdrItem.YFilter
    otdrItem.EntityData.YangName = "otdr-item"
    otdrItem.EntityData.BundleName = "cisco_ios_xr"
    otdrItem.EntityData.ParentYangName = "identifier"
    otdrItem.EntityData.SegmentPath = "otdr-item"
    otdrItem.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdr-details/otdr-detail/data-type-detail/identifier/" + otdrItem.EntityData.SegmentPath
    otdrItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrItem.EntityData.Children = types.NewOrderedMap()
    otdrItem.EntityData.Children.Append("otdr-event", types.YChild{"OtdrEvent", nil})
    for i := range otdrItem.OtdrEvent {
        types.SetYListKey(otdrItem.OtdrEvent[i], i)
        otdrItem.EntityData.Children.Append(types.GetSegmentPath(otdrItem.OtdrEvent[i]), types.YChild{"OtdrEvent", otdrItem.OtdrEvent[i]})
    }
    otdrItem.EntityData.Leafs = types.NewOrderedMap()
    otdrItem.EntityData.Leafs.Append("otdr-number", types.YLeaf{"OtdrNumber", otdrItem.OtdrNumber})
    otdrItem.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otdrItem.Direction})
    otdrItem.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", otdrItem.Timestamp})
    otdrItem.EntityData.Leafs.Append("sor-file", types.YLeaf{"SorFile", otdrItem.SorFile})
    otdrItem.EntityData.Leafs.Append("sor-directory", types.YLeaf{"SorDirectory", otdrItem.SorDirectory})
    otdrItem.EntityData.Leafs.Append("scan-mode", types.YLeaf{"ScanMode", otdrItem.ScanMode})
    otdrItem.EntityData.Leafs.Append("distance", types.YLeaf{"Distance", otdrItem.Distance})
    otdrItem.EntityData.Leafs.Append("total-orl", types.YLeaf{"TotalOrl", otdrItem.TotalOrl})

    otdrItem.EntityData.YListKeys = []string {}

    return &(otdrItem.EntityData)
}

// HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem_OtdrEvent
// Event list
type HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem_OtdrEvent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Event Id. The type is interface{} with range: 0..65535.
    EventId interface{}

    // Event Type. The type is OtdrEvent.
    Type interface{}

    // Event Position in Km. The type is string with length: 0..16.
    Position interface{}

    // Event accuracy in meters. The type is string with length: 0..16.
    Accuracy interface{}

    // Event Magnitude in dB. The type is string with length: 0..16.
    Magnitude interface{}

    // Event Attenuation in dB. The type is string with length: 0..16.
    Attenuation interface{}

    // Event End of Fiber Confidence. The type is interface{} with range:
    // 0..4294967295.
    EofConfidence interface{}

    // Threshold Crossing. The type is interface{} with range: 0..255.
    ThresholdCrossing interface{}
}

func (otdrEvent *HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrItem_OtdrEvent) GetEntityData() *types.CommonEntityData {
    otdrEvent.EntityData.YFilter = otdrEvent.YFilter
    otdrEvent.EntityData.YangName = "otdr-event"
    otdrEvent.EntityData.BundleName = "cisco_ios_xr"
    otdrEvent.EntityData.ParentYangName = "otdr-item"
    otdrEvent.EntityData.SegmentPath = "otdr-event" + types.AddNoKeyToken(otdrEvent)
    otdrEvent.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdr-details/otdr-detail/data-type-detail/identifier/otdr-item/" + otdrEvent.EntityData.SegmentPath
    otdrEvent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrEvent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrEvent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrEvent.EntityData.Children = types.NewOrderedMap()
    otdrEvent.EntityData.Leafs = types.NewOrderedMap()
    otdrEvent.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", otdrEvent.EventId})
    otdrEvent.EntityData.Leafs.Append("type", types.YLeaf{"Type", otdrEvent.Type})
    otdrEvent.EntityData.Leafs.Append("position", types.YLeaf{"Position", otdrEvent.Position})
    otdrEvent.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", otdrEvent.Accuracy})
    otdrEvent.EntityData.Leafs.Append("magnitude", types.YLeaf{"Magnitude", otdrEvent.Magnitude})
    otdrEvent.EntityData.Leafs.Append("attenuation", types.YLeaf{"Attenuation", otdrEvent.Attenuation})
    otdrEvent.EntityData.Leafs.Append("eof-confidence", types.YLeaf{"EofConfidence", otdrEvent.EofConfidence})
    otdrEvent.EntityData.Leafs.Append("threshold-crossing", types.YLeaf{"ThresholdCrossing", otdrEvent.ThresholdCrossing})

    otdrEvent.EntityData.YListKeys = []string {}

    return &(otdrEvent.EntityData)
}

// HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrScan
// otdr scan
type HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrScan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Index. The type is interface{} with range: 0..65535.
    Index interface{}

    // Otdr Number. The type is interface{} with range: 0..255.
    OtdrNumber interface{}

    // Direction. The type is Direction.
    Direction interface{}

    // Scan Mode. The type is OtdrScanMode.
    ScanMode interface{}

    // Timestamp. The type is string with length: 0..16.
    Timestamp interface{}

    // Sor File Name. The type is string with length: 0..64.
    SorFile interface{}
}

func (otdrScan *HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrScan) GetEntityData() *types.CommonEntityData {
    otdrScan.EntityData.YFilter = otdrScan.YFilter
    otdrScan.EntityData.YangName = "otdr-scan"
    otdrScan.EntityData.BundleName = "cisco_ios_xr"
    otdrScan.EntityData.ParentYangName = "identifier"
    otdrScan.EntityData.SegmentPath = "otdr-scan" + types.AddNoKeyToken(otdrScan)
    otdrScan.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdr-details/otdr-detail/data-type-detail/identifier/" + otdrScan.EntityData.SegmentPath
    otdrScan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrScan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrScan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrScan.EntityData.Children = types.NewOrderedMap()
    otdrScan.EntityData.Leafs = types.NewOrderedMap()
    otdrScan.EntityData.Leafs.Append("index", types.YLeaf{"Index", otdrScan.Index})
    otdrScan.EntityData.Leafs.Append("otdr-number", types.YLeaf{"OtdrNumber", otdrScan.OtdrNumber})
    otdrScan.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otdrScan.Direction})
    otdrScan.EntityData.Leafs.Append("scan-mode", types.YLeaf{"ScanMode", otdrScan.ScanMode})
    otdrScan.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", otdrScan.Timestamp})
    otdrScan.EntityData.Leafs.Append("sor-file", types.YLeaf{"SorFile", otdrScan.SorFile})

    otdrScan.EntityData.YListKeys = []string {}

    return &(otdrScan.EntityData)
}

// HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrStatus
// otdr status
type HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Otdr Id. The type is interface{} with range: 0..255.
    OtdrId interface{}

    // Direction. The type is Direction.
    Direction interface{}

    // Timestamp. The type is string with length: 0..16.
    Timestamp interface{}

    // Training Status. The type is OtdrStatus.
    TrainingStatus interface{}

    // Measurement Status. The type is OtdrStatus.
    MeasurementStatus interface{}

    // Progress %. The type is interface{} with range: 0..65535.
    ProgressPc interface{}
}

func (otdrStatus *HwModule_OtdrDetails_OtdrDetail_DataTypeDetail_Identifier_OtdrStatus) GetEntityData() *types.CommonEntityData {
    otdrStatus.EntityData.YFilter = otdrStatus.YFilter
    otdrStatus.EntityData.YangName = "otdr-status"
    otdrStatus.EntityData.BundleName = "cisco_ios_xr"
    otdrStatus.EntityData.ParentYangName = "identifier"
    otdrStatus.EntityData.SegmentPath = "otdr-status" + types.AddNoKeyToken(otdrStatus)
    otdrStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdr-details/otdr-detail/data-type-detail/identifier/" + otdrStatus.EntityData.SegmentPath
    otdrStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrStatus.EntityData.Children = types.NewOrderedMap()
    otdrStatus.EntityData.Leafs = types.NewOrderedMap()
    otdrStatus.EntityData.Leafs.Append("otdr-id", types.YLeaf{"OtdrId", otdrStatus.OtdrId})
    otdrStatus.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otdrStatus.Direction})
    otdrStatus.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", otdrStatus.Timestamp})
    otdrStatus.EntityData.Leafs.Append("training-status", types.YLeaf{"TrainingStatus", otdrStatus.TrainingStatus})
    otdrStatus.EntityData.Leafs.Append("measurement-status", types.YLeaf{"MeasurementStatus", otdrStatus.MeasurementStatus})
    otdrStatus.EntityData.Leafs.Append("progress-pc", types.YLeaf{"ProgressPc", otdrStatus.ProgressPc})

    otdrStatus.EntityData.YListKeys = []string {}

    return &(otdrStatus.EntityData)
}

// HwModule_Otdrs
// Otdr Root Info
type HwModule_Otdrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Otdr Info. The type is slice of HwModule_Otdrs_Otdr.
    Otdr []*HwModule_Otdrs_Otdr
}

func (otdrs *HwModule_Otdrs) GetEntityData() *types.CommonEntityData {
    otdrs.EntityData.YFilter = otdrs.YFilter
    otdrs.EntityData.YangName = "otdrs"
    otdrs.EntityData.BundleName = "cisco_ios_xr"
    otdrs.EntityData.ParentYangName = "hw-module"
    otdrs.EntityData.SegmentPath = "otdrs"
    otdrs.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/" + otdrs.EntityData.SegmentPath
    otdrs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrs.EntityData.Children = types.NewOrderedMap()
    otdrs.EntityData.Children.Append("otdr", types.YChild{"Otdr", nil})
    for i := range otdrs.Otdr {
        otdrs.EntityData.Children.Append(types.GetSegmentPath(otdrs.Otdr[i]), types.YChild{"Otdr", otdrs.Otdr[i]})
    }
    otdrs.EntityData.Leafs = types.NewOrderedMap()

    otdrs.EntityData.YListKeys = []string {}

    return &(otdrs.EntityData)
}

// HwModule_Otdrs_Otdr
// Otdr Info
type HwModule_Otdrs_Otdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Details associated with a particular slot number.
    // The type is interface{} with range: 1..3.
    SlotId interface{}

    // Dispaly data type list. The type is slice of HwModule_Otdrs_Otdr_DataType.
    DataType []*HwModule_Otdrs_Otdr_DataType
}

func (otdr *HwModule_Otdrs_Otdr) GetEntityData() *types.CommonEntityData {
    otdr.EntityData.YFilter = otdr.YFilter
    otdr.EntityData.YangName = "otdr"
    otdr.EntityData.BundleName = "cisco_ios_xr"
    otdr.EntityData.ParentYangName = "otdrs"
    otdr.EntityData.SegmentPath = "otdr" + types.AddKeyToken(otdr.SlotId, "slot-id")
    otdr.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdrs/" + otdr.EntityData.SegmentPath
    otdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdr.EntityData.Children = types.NewOrderedMap()
    otdr.EntityData.Children.Append("data-type", types.YChild{"DataType", nil})
    for i := range otdr.DataType {
        otdr.EntityData.Children.Append(types.GetSegmentPath(otdr.DataType[i]), types.YChild{"DataType", otdr.DataType[i]})
    }
    otdr.EntityData.Leafs = types.NewOrderedMap()
    otdr.EntityData.Leafs.Append("slot-id", types.YLeaf{"SlotId", otdr.SlotId})

    otdr.EntityData.YListKeys = []string {"SlotId"}

    return &(otdr.EntityData)
}

// HwModule_Otdrs_Otdr_DataType
// Dispaly data type list
type HwModule_Otdrs_Otdr_DataType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Details associated with a particular Data type.
    // The type is OtsOtdrData.
    OtdrDataType interface{}

    // otdr item.
    OtdrItem HwModule_Otdrs_Otdr_DataType_OtdrItem

    // otdr scan. The type is slice of HwModule_Otdrs_Otdr_DataType_OtdrScan.
    OtdrScan []*HwModule_Otdrs_Otdr_DataType_OtdrScan

    // otdr status. The type is slice of HwModule_Otdrs_Otdr_DataType_OtdrStatus.
    OtdrStatus []*HwModule_Otdrs_Otdr_DataType_OtdrStatus
}

func (dataType *HwModule_Otdrs_Otdr_DataType) GetEntityData() *types.CommonEntityData {
    dataType.EntityData.YFilter = dataType.YFilter
    dataType.EntityData.YangName = "data-type"
    dataType.EntityData.BundleName = "cisco_ios_xr"
    dataType.EntityData.ParentYangName = "otdr"
    dataType.EntityData.SegmentPath = "data-type" + types.AddKeyToken(dataType.OtdrDataType, "otdr-data-type")
    dataType.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdrs/otdr/" + dataType.EntityData.SegmentPath
    dataType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataType.EntityData.Children = types.NewOrderedMap()
    dataType.EntityData.Children.Append("otdr-item", types.YChild{"OtdrItem", &dataType.OtdrItem})
    dataType.EntityData.Children.Append("otdr-scan", types.YChild{"OtdrScan", nil})
    for i := range dataType.OtdrScan {
        types.SetYListKey(dataType.OtdrScan[i], i)
        dataType.EntityData.Children.Append(types.GetSegmentPath(dataType.OtdrScan[i]), types.YChild{"OtdrScan", dataType.OtdrScan[i]})
    }
    dataType.EntityData.Children.Append("otdr-status", types.YChild{"OtdrStatus", nil})
    for i := range dataType.OtdrStatus {
        types.SetYListKey(dataType.OtdrStatus[i], i)
        dataType.EntityData.Children.Append(types.GetSegmentPath(dataType.OtdrStatus[i]), types.YChild{"OtdrStatus", dataType.OtdrStatus[i]})
    }
    dataType.EntityData.Leafs = types.NewOrderedMap()
    dataType.EntityData.Leafs.Append("otdr-data-type", types.YLeaf{"OtdrDataType", dataType.OtdrDataType})

    dataType.EntityData.YListKeys = []string {"OtdrDataType"}

    return &(dataType.EntityData)
}

// HwModule_Otdrs_Otdr_DataType_OtdrItem
// otdr item
type HwModule_Otdrs_Otdr_DataType_OtdrItem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Otdr Number. The type is interface{} with range: 0..255.
    OtdrNumber interface{}

    // Scan Direction. The type is Direction.
    Direction interface{}

    // Timestamp. The type is string with length: 0..16.
    Timestamp interface{}

    // Sor file name. The type is string with length: 0..64.
    SorFile interface{}

    // Sor file location. The type is string with length: 0..64.
    SorDirectory interface{}

    // Scan Mode. The type is OtdrScanMode.
    ScanMode interface{}

    // Distance in Km. The type is string with length: 0..16.
    Distance interface{}

    // Total Orl in dB. The type is string with length: 0..16.
    TotalOrl interface{}

    // Event list. The type is slice of
    // HwModule_Otdrs_Otdr_DataType_OtdrItem_OtdrEvent.
    OtdrEvent []*HwModule_Otdrs_Otdr_DataType_OtdrItem_OtdrEvent
}

func (otdrItem *HwModule_Otdrs_Otdr_DataType_OtdrItem) GetEntityData() *types.CommonEntityData {
    otdrItem.EntityData.YFilter = otdrItem.YFilter
    otdrItem.EntityData.YangName = "otdr-item"
    otdrItem.EntityData.BundleName = "cisco_ios_xr"
    otdrItem.EntityData.ParentYangName = "data-type"
    otdrItem.EntityData.SegmentPath = "otdr-item"
    otdrItem.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdrs/otdr/data-type/" + otdrItem.EntityData.SegmentPath
    otdrItem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrItem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrItem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrItem.EntityData.Children = types.NewOrderedMap()
    otdrItem.EntityData.Children.Append("otdr-event", types.YChild{"OtdrEvent", nil})
    for i := range otdrItem.OtdrEvent {
        types.SetYListKey(otdrItem.OtdrEvent[i], i)
        otdrItem.EntityData.Children.Append(types.GetSegmentPath(otdrItem.OtdrEvent[i]), types.YChild{"OtdrEvent", otdrItem.OtdrEvent[i]})
    }
    otdrItem.EntityData.Leafs = types.NewOrderedMap()
    otdrItem.EntityData.Leafs.Append("otdr-number", types.YLeaf{"OtdrNumber", otdrItem.OtdrNumber})
    otdrItem.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otdrItem.Direction})
    otdrItem.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", otdrItem.Timestamp})
    otdrItem.EntityData.Leafs.Append("sor-file", types.YLeaf{"SorFile", otdrItem.SorFile})
    otdrItem.EntityData.Leafs.Append("sor-directory", types.YLeaf{"SorDirectory", otdrItem.SorDirectory})
    otdrItem.EntityData.Leafs.Append("scan-mode", types.YLeaf{"ScanMode", otdrItem.ScanMode})
    otdrItem.EntityData.Leafs.Append("distance", types.YLeaf{"Distance", otdrItem.Distance})
    otdrItem.EntityData.Leafs.Append("total-orl", types.YLeaf{"TotalOrl", otdrItem.TotalOrl})

    otdrItem.EntityData.YListKeys = []string {}

    return &(otdrItem.EntityData)
}

// HwModule_Otdrs_Otdr_DataType_OtdrItem_OtdrEvent
// Event list
type HwModule_Otdrs_Otdr_DataType_OtdrItem_OtdrEvent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Event Id. The type is interface{} with range: 0..65535.
    EventId interface{}

    // Event Type. The type is OtdrEvent.
    Type interface{}

    // Event Position in Km. The type is string with length: 0..16.
    Position interface{}

    // Event accuracy in meters. The type is string with length: 0..16.
    Accuracy interface{}

    // Event Magnitude in dB. The type is string with length: 0..16.
    Magnitude interface{}

    // Event Attenuation in dB. The type is string with length: 0..16.
    Attenuation interface{}

    // Event End of Fiber Confidence. The type is interface{} with range:
    // 0..4294967295.
    EofConfidence interface{}

    // Threshold Crossing. The type is interface{} with range: 0..255.
    ThresholdCrossing interface{}
}

func (otdrEvent *HwModule_Otdrs_Otdr_DataType_OtdrItem_OtdrEvent) GetEntityData() *types.CommonEntityData {
    otdrEvent.EntityData.YFilter = otdrEvent.YFilter
    otdrEvent.EntityData.YangName = "otdr-event"
    otdrEvent.EntityData.BundleName = "cisco_ios_xr"
    otdrEvent.EntityData.ParentYangName = "otdr-item"
    otdrEvent.EntityData.SegmentPath = "otdr-event" + types.AddNoKeyToken(otdrEvent)
    otdrEvent.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdrs/otdr/data-type/otdr-item/" + otdrEvent.EntityData.SegmentPath
    otdrEvent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrEvent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrEvent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrEvent.EntityData.Children = types.NewOrderedMap()
    otdrEvent.EntityData.Leafs = types.NewOrderedMap()
    otdrEvent.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", otdrEvent.EventId})
    otdrEvent.EntityData.Leafs.Append("type", types.YLeaf{"Type", otdrEvent.Type})
    otdrEvent.EntityData.Leafs.Append("position", types.YLeaf{"Position", otdrEvent.Position})
    otdrEvent.EntityData.Leafs.Append("accuracy", types.YLeaf{"Accuracy", otdrEvent.Accuracy})
    otdrEvent.EntityData.Leafs.Append("magnitude", types.YLeaf{"Magnitude", otdrEvent.Magnitude})
    otdrEvent.EntityData.Leafs.Append("attenuation", types.YLeaf{"Attenuation", otdrEvent.Attenuation})
    otdrEvent.EntityData.Leafs.Append("eof-confidence", types.YLeaf{"EofConfidence", otdrEvent.EofConfidence})
    otdrEvent.EntityData.Leafs.Append("threshold-crossing", types.YLeaf{"ThresholdCrossing", otdrEvent.ThresholdCrossing})

    otdrEvent.EntityData.YListKeys = []string {}

    return &(otdrEvent.EntityData)
}

// HwModule_Otdrs_Otdr_DataType_OtdrScan
// otdr scan
type HwModule_Otdrs_Otdr_DataType_OtdrScan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Index. The type is interface{} with range: 0..65535.
    Index interface{}

    // Otdr Number. The type is interface{} with range: 0..255.
    OtdrNumber interface{}

    // Direction. The type is Direction.
    Direction interface{}

    // Scan Mode. The type is OtdrScanMode.
    ScanMode interface{}

    // Timestamp. The type is string with length: 0..16.
    Timestamp interface{}

    // Sor File Name. The type is string with length: 0..64.
    SorFile interface{}
}

func (otdrScan *HwModule_Otdrs_Otdr_DataType_OtdrScan) GetEntityData() *types.CommonEntityData {
    otdrScan.EntityData.YFilter = otdrScan.YFilter
    otdrScan.EntityData.YangName = "otdr-scan"
    otdrScan.EntityData.BundleName = "cisco_ios_xr"
    otdrScan.EntityData.ParentYangName = "data-type"
    otdrScan.EntityData.SegmentPath = "otdr-scan" + types.AddNoKeyToken(otdrScan)
    otdrScan.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdrs/otdr/data-type/" + otdrScan.EntityData.SegmentPath
    otdrScan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrScan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrScan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrScan.EntityData.Children = types.NewOrderedMap()
    otdrScan.EntityData.Leafs = types.NewOrderedMap()
    otdrScan.EntityData.Leafs.Append("index", types.YLeaf{"Index", otdrScan.Index})
    otdrScan.EntityData.Leafs.Append("otdr-number", types.YLeaf{"OtdrNumber", otdrScan.OtdrNumber})
    otdrScan.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otdrScan.Direction})
    otdrScan.EntityData.Leafs.Append("scan-mode", types.YLeaf{"ScanMode", otdrScan.ScanMode})
    otdrScan.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", otdrScan.Timestamp})
    otdrScan.EntityData.Leafs.Append("sor-file", types.YLeaf{"SorFile", otdrScan.SorFile})

    otdrScan.EntityData.YListKeys = []string {}

    return &(otdrScan.EntityData)
}

// HwModule_Otdrs_Otdr_DataType_OtdrStatus
// otdr status
type HwModule_Otdrs_Otdr_DataType_OtdrStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Otdr Id. The type is interface{} with range: 0..255.
    OtdrId interface{}

    // Direction. The type is Direction.
    Direction interface{}

    // Timestamp. The type is string with length: 0..16.
    Timestamp interface{}

    // Training Status. The type is OtdrStatus.
    TrainingStatus interface{}

    // Measurement Status. The type is OtdrStatus.
    MeasurementStatus interface{}

    // Progress %. The type is interface{} with range: 0..65535.
    ProgressPc interface{}
}

func (otdrStatus *HwModule_Otdrs_Otdr_DataType_OtdrStatus) GetEntityData() *types.CommonEntityData {
    otdrStatus.EntityData.YFilter = otdrStatus.YFilter
    otdrStatus.EntityData.YangName = "otdr-status"
    otdrStatus.EntityData.BundleName = "cisco_ios_xr"
    otdrStatus.EntityData.ParentYangName = "data-type"
    otdrStatus.EntityData.SegmentPath = "otdr-status" + types.AddNoKeyToken(otdrStatus)
    otdrStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1001-otdr-oper:hw-module/otdrs/otdr/data-type/" + otdrStatus.EntityData.SegmentPath
    otdrStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otdrStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otdrStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otdrStatus.EntityData.Children = types.NewOrderedMap()
    otdrStatus.EntityData.Leafs = types.NewOrderedMap()
    otdrStatus.EntityData.Leafs.Append("otdr-id", types.YLeaf{"OtdrId", otdrStatus.OtdrId})
    otdrStatus.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", otdrStatus.Direction})
    otdrStatus.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", otdrStatus.Timestamp})
    otdrStatus.EntityData.Leafs.Append("training-status", types.YLeaf{"TrainingStatus", otdrStatus.TrainingStatus})
    otdrStatus.EntityData.Leafs.Append("measurement-status", types.YLeaf{"MeasurementStatus", otdrStatus.MeasurementStatus})
    otdrStatus.EntityData.Leafs.Append("progress-pc", types.YLeaf{"ProgressPc", otdrStatus.ProgressPc})

    otdrStatus.EntityData.YListKeys = []string {}

    return &(otdrStatus.EntityData)
}

