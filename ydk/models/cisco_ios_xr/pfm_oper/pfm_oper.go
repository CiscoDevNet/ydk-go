// This module contains a collection of YANG definitions
// for Cisco IOS-XR pfm package operational data.
// 
// This module contains definitions
// for the following management objects:
//   platform-fault-manager: PFM data space
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package pfm_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pfm_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pfm-oper platform-fault-manager}", reflect.TypeOf(PlatformFaultManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pfm-oper:platform-fault-manager", reflect.TypeOf(PlatformFaultManager{}))
}

// PlatformFaultManager
// PFM data space
type PlatformFaultManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude specic hw fault .
    Exclude PlatformFaultManager_Exclude

    // Table of racks.
    Racks PlatformFaultManager_Racks
}

func (platformFaultManager *PlatformFaultManager) GetEntityData() *types.CommonEntityData {
    platformFaultManager.EntityData.YFilter = platformFaultManager.YFilter
    platformFaultManager.EntityData.YangName = "platform-fault-manager"
    platformFaultManager.EntityData.BundleName = "cisco_ios_xr"
    platformFaultManager.EntityData.ParentYangName = "Cisco-IOS-XR-pfm-oper"
    platformFaultManager.EntityData.SegmentPath = "Cisco-IOS-XR-pfm-oper:platform-fault-manager"
    platformFaultManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformFaultManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformFaultManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformFaultManager.EntityData.Children = make(map[string]types.YChild)
    platformFaultManager.EntityData.Children["exclude"] = types.YChild{"Exclude", &platformFaultManager.Exclude}
    platformFaultManager.EntityData.Children["racks"] = types.YChild{"Racks", &platformFaultManager.Racks}
    platformFaultManager.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(platformFaultManager.EntityData)
}

// PlatformFaultManager_Exclude
// Exclude specic hw fault 
type PlatformFaultManager_Exclude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device.
    FaultType1S PlatformFaultManager_Exclude_FaultType1S
}

func (exclude *PlatformFaultManager_Exclude) GetEntityData() *types.CommonEntityData {
    exclude.EntityData.YFilter = exclude.YFilter
    exclude.EntityData.YangName = "exclude"
    exclude.EntityData.BundleName = "cisco_ios_xr"
    exclude.EntityData.ParentYangName = "platform-fault-manager"
    exclude.EntityData.SegmentPath = "exclude"
    exclude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exclude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exclude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exclude.EntityData.Children = make(map[string]types.YChild)
    exclude.EntityData.Children["fault-type1s"] = types.YChild{"FaultType1S", &exclude.FaultType1S}
    exclude.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(exclude.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1.
    FaultType1 []PlatformFaultManager_Exclude_FaultType1S_FaultType1
}

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetEntityData() *types.CommonEntityData {
    faultType1S.EntityData.YFilter = faultType1S.YFilter
    faultType1S.EntityData.YangName = "fault-type1s"
    faultType1S.EntityData.BundleName = "cisco_ios_xr"
    faultType1S.EntityData.ParentYangName = "exclude"
    faultType1S.EntityData.SegmentPath = "fault-type1s"
    faultType1S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultType1S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultType1S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultType1S.EntityData.Children = make(map[string]types.YChild)
    faultType1S.EntityData.Children["fault-type1"] = types.YChild{"FaultType1", nil}
    for i := range faultType1S.FaultType1 {
        faultType1S.EntityData.Children[types.GetSegmentPath(&faultType1S.FaultType1[i])] = types.YChild{"FaultType1", &faultType1S.FaultType1[i]}
    }
    faultType1S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(faultType1S.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault 1. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultType1 interface{}

    // Table of Hardware Failure Device.
    FaultType2S PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S

    // Table of racks.
    Racks PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks
}

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetEntityData() *types.CommonEntityData {
    faultType1.EntityData.YFilter = faultType1.YFilter
    faultType1.EntityData.YangName = "fault-type1"
    faultType1.EntityData.BundleName = "cisco_ios_xr"
    faultType1.EntityData.ParentYangName = "fault-type1s"
    faultType1.EntityData.SegmentPath = "fault-type1" + "[hw-fault-type1='" + fmt.Sprintf("%v", faultType1.HwFaultType1) + "']"
    faultType1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultType1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultType1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultType1.EntityData.Children = make(map[string]types.YChild)
    faultType1.EntityData.Children["fault-type2s"] = types.YChild{"FaultType2S", &faultType1.FaultType2S}
    faultType1.EntityData.Children["racks"] = types.YChild{"Racks", &faultType1.Racks}
    faultType1.EntityData.Leafs = make(map[string]types.YLeaf)
    faultType1.EntityData.Leafs["hw-fault-type1"] = types.YLeaf{"HwFaultType1", faultType1.HwFaultType1}
    return &(faultType1.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2.
    FaultType2 []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2
}

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetEntityData() *types.CommonEntityData {
    faultType2S.EntityData.YFilter = faultType2S.YFilter
    faultType2S.EntityData.YangName = "fault-type2s"
    faultType2S.EntityData.BundleName = "cisco_ios_xr"
    faultType2S.EntityData.ParentYangName = "fault-type1"
    faultType2S.EntityData.SegmentPath = "fault-type2s"
    faultType2S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultType2S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultType2S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultType2S.EntityData.Children = make(map[string]types.YChild)
    faultType2S.EntityData.Children["fault-type2"] = types.YChild{"FaultType2", nil}
    for i := range faultType2S.FaultType2 {
        faultType2S.EntityData.Children[types.GetSegmentPath(&faultType2S.FaultType2[i])] = types.YChild{"FaultType2", &faultType2S.FaultType2[i]}
    }
    faultType2S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(faultType2S.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault 2. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultType2 interface{}

    // Table of Hardware Failure Device.
    FaultType3S PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S

    // Table of racks.
    Racks PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks
}

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetEntityData() *types.CommonEntityData {
    faultType2.EntityData.YFilter = faultType2.YFilter
    faultType2.EntityData.YangName = "fault-type2"
    faultType2.EntityData.BundleName = "cisco_ios_xr"
    faultType2.EntityData.ParentYangName = "fault-type2s"
    faultType2.EntityData.SegmentPath = "fault-type2" + "[hw-fault-type2='" + fmt.Sprintf("%v", faultType2.HwFaultType2) + "']"
    faultType2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultType2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultType2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultType2.EntityData.Children = make(map[string]types.YChild)
    faultType2.EntityData.Children["fault-type3s"] = types.YChild{"FaultType3S", &faultType2.FaultType3S}
    faultType2.EntityData.Children["racks"] = types.YChild{"Racks", &faultType2.Racks}
    faultType2.EntityData.Leafs = make(map[string]types.YLeaf)
    faultType2.EntityData.Leafs["hw-fault-type2"] = types.YLeaf{"HwFaultType2", faultType2.HwFaultType2}
    return &(faultType2.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3.
    FaultType3 []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3
}

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetEntityData() *types.CommonEntityData {
    faultType3S.EntityData.YFilter = faultType3S.YFilter
    faultType3S.EntityData.YangName = "fault-type3s"
    faultType3S.EntityData.BundleName = "cisco_ios_xr"
    faultType3S.EntityData.ParentYangName = "fault-type2"
    faultType3S.EntityData.SegmentPath = "fault-type3s"
    faultType3S.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultType3S.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultType3S.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultType3S.EntityData.Children = make(map[string]types.YChild)
    faultType3S.EntityData.Children["fault-type3"] = types.YChild{"FaultType3", nil}
    for i := range faultType3S.FaultType3 {
        faultType3S.EntityData.Children[types.GetSegmentPath(&faultType3S.FaultType3[i])] = types.YChild{"FaultType3", &faultType3S.FaultType3[i]}
    }
    faultType3S.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(faultType3S.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault 3. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultType3 interface{}

    // Table of racks.
    Racks PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks
}

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetEntityData() *types.CommonEntityData {
    faultType3.EntityData.YFilter = faultType3.YFilter
    faultType3.EntityData.YangName = "fault-type3"
    faultType3.EntityData.BundleName = "cisco_ios_xr"
    faultType3.EntityData.ParentYangName = "fault-type3s"
    faultType3.EntityData.SegmentPath = "fault-type3" + "[hw-fault-type3='" + fmt.Sprintf("%v", faultType3.HwFaultType3) + "']"
    faultType3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultType3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultType3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultType3.EntityData.Children = make(map[string]types.YChild)
    faultType3.EntityData.Children["racks"] = types.YChild{"Racks", &faultType3.Racks}
    faultType3.EntityData.Leafs = make(map[string]types.YLeaf)
    faultType3.EntityData.Leafs["hw-fault-type3"] = types.YLeaf{"HwFaultType3", faultType3.HwFaultType3}
    return &(faultType3.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks
// Table of racks
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack.
    Rack []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "fault-type3"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = make(map[string]types.YChild)
    racks.EntityData.Children["rack"] = types.YChild{"Rack", nil}
    for i := range racks.Rack {
        racks.EntityData.Children[types.GetSegmentPath(&racks.Rack[i])] = types.YChild{"Rack", &racks.Rack[i]}
    }
    racks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(racks.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack
// Number
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["slots"] = types.YChild{"Slots", &rack.Slots}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["rack"] = types.YLeaf{"Rack", rack.Rack}
    return &(rack.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots
// Table of slots
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot.
    Slot []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = make(map[string]types.YChild)
    slots.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range slots.Slot {
        slots.EntityData.Children[types.GetSegmentPath(&slots.Slot[i])] = types.YChild{"Slot", &slots.Slot[i]}
    }
    slots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slots.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot
// Name
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Slot interface{}

    // Table of Hardware Summary.
    FaultSummary PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary

    // Table of Hardware Failure.
    HardwareFaultDevices PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["fault-summary"] = types.YChild{"FaultSummary", &slot.FaultSummary}
    slot.EntityData.Children["hardware-fault-devices"] = types.YChild{"HardwareFaultDevices", &slot.HardwareFaultDevices}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["slot"] = types.YLeaf{"Slot", slot.Slot}
    return &(slot.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary
// Table of Hardware Summary
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fault Severity Critical count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityCriticalCount interface{}

    // Fault Severity Emergency count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityEmergencyOrAlertCount interface{}

    // Faulty Hardware total count. The type is interface{} with range:
    // -2147483648..2147483647.
    Total interface{}

    // Fault Severity Error count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityErrorCount interface{}
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetEntityData() *types.CommonEntityData {
    faultSummary.EntityData.YFilter = faultSummary.YFilter
    faultSummary.EntityData.YangName = "fault-summary"
    faultSummary.EntityData.BundleName = "cisco_ios_xr"
    faultSummary.EntityData.ParentYangName = "slot"
    faultSummary.EntityData.SegmentPath = "fault-summary"
    faultSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultSummary.EntityData.Children = make(map[string]types.YChild)
    faultSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    faultSummary.EntityData.Leafs["severity-critical-count"] = types.YLeaf{"SeverityCriticalCount", faultSummary.SeverityCriticalCount}
    faultSummary.EntityData.Leafs["severity-emergency-or-alert-count"] = types.YLeaf{"SeverityEmergencyOrAlertCount", faultSummary.SeverityEmergencyOrAlertCount}
    faultSummary.EntityData.Leafs["total"] = types.YLeaf{"Total", faultSummary.Total}
    faultSummary.EntityData.Leafs["severity-error-count"] = types.YLeaf{"SeverityErrorCount", faultSummary.SeverityErrorCount}
    return &(faultSummary.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices
// Table of Hardware Failure
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice.
    HardwareFaultDevice []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetEntityData() *types.CommonEntityData {
    hardwareFaultDevices.EntityData.YFilter = hardwareFaultDevices.YFilter
    hardwareFaultDevices.EntityData.YangName = "hardware-fault-devices"
    hardwareFaultDevices.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultDevices.EntityData.ParentYangName = "slot"
    hardwareFaultDevices.EntityData.SegmentPath = "hardware-fault-devices"
    hardwareFaultDevices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultDevices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultDevices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultDevices.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultDevices.EntityData.Children["hardware-fault-device"] = types.YChild{"HardwareFaultDevice", nil}
    for i := range hardwareFaultDevices.HardwareFaultDevice {
        hardwareFaultDevices.EntityData.Children[types.GetSegmentPath(&hardwareFaultDevices.HardwareFaultDevice[i])] = types.YChild{"HardwareFaultDevice", &hardwareFaultDevices.HardwareFaultDevice[i]}
    }
    hardwareFaultDevices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareFaultDevices.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault device list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultDevice interface{}

    // Table of Hardware Failure Type. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType.
    HardwareFaultType []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetEntityData() *types.CommonEntityData {
    hardwareFaultDevice.EntityData.YFilter = hardwareFaultDevice.YFilter
    hardwareFaultDevice.EntityData.YangName = "hardware-fault-device"
    hardwareFaultDevice.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultDevice.EntityData.ParentYangName = "hardware-fault-devices"
    hardwareFaultDevice.EntityData.SegmentPath = "hardware-fault-device" + "[hw-fault-device='" + fmt.Sprintf("%v", hardwareFaultDevice.HwFaultDevice) + "']"
    hardwareFaultDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultDevice.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultDevice.EntityData.Children["hardware-fault-type"] = types.YChild{"HardwareFaultType", nil}
    for i := range hardwareFaultDevice.HardwareFaultType {
        hardwareFaultDevice.EntityData.Children[types.GetSegmentPath(&hardwareFaultDevice.HardwareFaultType[i])] = types.YChild{"HardwareFaultType", &hardwareFaultDevice.HardwareFaultType[i]}
    }
    hardwareFaultDevice.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareFaultDevice.EntityData.Leafs["hw-fault-device"] = types.YLeaf{"HwFaultDevice", hardwareFaultDevice.HwFaultDevice}
    return &(hardwareFaultDevice.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
// Table of Hardware Failure Type
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault type list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultType interface{}

    // Faulty Hardware Condition Description. The type is string.
    ConditionDescription interface{}

    // Faulty Hardware Condition Name. The type is string.
    ConditionName interface{}

    // Faulty Hardware Device Key. The type is string.
    DeviceKey interface{}

    // Faulty Hardware Device Version. The type is interface{} with range:
    // -2147483648..2147483647.
    DeviceVersion interface{}

    // Fault Raised Timestamp. The type is string.
    ConditionRaisedTimestamp interface{}

    // Faulty Hardware Process ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Faulty Hardware Device Description. The type is string.
    DeviceDescription interface{}

    // Faulty Hardware Condition Severity. The type is string.
    ConditionSeverity interface{}
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetEntityData() *types.CommonEntityData {
    hardwareFaultType.EntityData.YFilter = hardwareFaultType.YFilter
    hardwareFaultType.EntityData.YangName = "hardware-fault-type"
    hardwareFaultType.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultType.EntityData.ParentYangName = "hardware-fault-device"
    hardwareFaultType.EntityData.SegmentPath = "hardware-fault-type" + "[hw-fault-type='" + fmt.Sprintf("%v", hardwareFaultType.HwFaultType) + "']"
    hardwareFaultType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultType.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultType.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareFaultType.EntityData.Leafs["hw-fault-type"] = types.YLeaf{"HwFaultType", hardwareFaultType.HwFaultType}
    hardwareFaultType.EntityData.Leafs["condition-description"] = types.YLeaf{"ConditionDescription", hardwareFaultType.ConditionDescription}
    hardwareFaultType.EntityData.Leafs["condition-name"] = types.YLeaf{"ConditionName", hardwareFaultType.ConditionName}
    hardwareFaultType.EntityData.Leafs["device-key"] = types.YLeaf{"DeviceKey", hardwareFaultType.DeviceKey}
    hardwareFaultType.EntityData.Leafs["device-version"] = types.YLeaf{"DeviceVersion", hardwareFaultType.DeviceVersion}
    hardwareFaultType.EntityData.Leafs["condition-raised-timestamp"] = types.YLeaf{"ConditionRaisedTimestamp", hardwareFaultType.ConditionRaisedTimestamp}
    hardwareFaultType.EntityData.Leafs["process-id"] = types.YLeaf{"ProcessId", hardwareFaultType.ProcessId}
    hardwareFaultType.EntityData.Leafs["device-description"] = types.YLeaf{"DeviceDescription", hardwareFaultType.DeviceDescription}
    hardwareFaultType.EntityData.Leafs["condition-severity"] = types.YLeaf{"ConditionSeverity", hardwareFaultType.ConditionSeverity}
    return &(hardwareFaultType.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks
// Table of racks
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack.
    Rack []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "fault-type2"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = make(map[string]types.YChild)
    racks.EntityData.Children["rack"] = types.YChild{"Rack", nil}
    for i := range racks.Rack {
        racks.EntityData.Children[types.GetSegmentPath(&racks.Rack[i])] = types.YChild{"Rack", &racks.Rack[i]}
    }
    racks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(racks.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack
// Number
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["slots"] = types.YChild{"Slots", &rack.Slots}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["rack"] = types.YLeaf{"Rack", rack.Rack}
    return &(rack.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots
// Table of slots
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot.
    Slot []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = make(map[string]types.YChild)
    slots.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range slots.Slot {
        slots.EntityData.Children[types.GetSegmentPath(&slots.Slot[i])] = types.YChild{"Slot", &slots.Slot[i]}
    }
    slots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slots.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot
// Name
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Slot interface{}

    // Table of Hardware Summary.
    FaultSummary PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary

    // Table of Hardware Failure.
    HardwareFaultDevices PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["fault-summary"] = types.YChild{"FaultSummary", &slot.FaultSummary}
    slot.EntityData.Children["hardware-fault-devices"] = types.YChild{"HardwareFaultDevices", &slot.HardwareFaultDevices}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["slot"] = types.YLeaf{"Slot", slot.Slot}
    return &(slot.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary
// Table of Hardware Summary
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fault Severity Critical count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityCriticalCount interface{}

    // Fault Severity Emergency count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityEmergencyOrAlertCount interface{}

    // Faulty Hardware total count. The type is interface{} with range:
    // -2147483648..2147483647.
    Total interface{}

    // Fault Severity Error count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityErrorCount interface{}
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetEntityData() *types.CommonEntityData {
    faultSummary.EntityData.YFilter = faultSummary.YFilter
    faultSummary.EntityData.YangName = "fault-summary"
    faultSummary.EntityData.BundleName = "cisco_ios_xr"
    faultSummary.EntityData.ParentYangName = "slot"
    faultSummary.EntityData.SegmentPath = "fault-summary"
    faultSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultSummary.EntityData.Children = make(map[string]types.YChild)
    faultSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    faultSummary.EntityData.Leafs["severity-critical-count"] = types.YLeaf{"SeverityCriticalCount", faultSummary.SeverityCriticalCount}
    faultSummary.EntityData.Leafs["severity-emergency-or-alert-count"] = types.YLeaf{"SeverityEmergencyOrAlertCount", faultSummary.SeverityEmergencyOrAlertCount}
    faultSummary.EntityData.Leafs["total"] = types.YLeaf{"Total", faultSummary.Total}
    faultSummary.EntityData.Leafs["severity-error-count"] = types.YLeaf{"SeverityErrorCount", faultSummary.SeverityErrorCount}
    return &(faultSummary.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices
// Table of Hardware Failure
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice.
    HardwareFaultDevice []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetEntityData() *types.CommonEntityData {
    hardwareFaultDevices.EntityData.YFilter = hardwareFaultDevices.YFilter
    hardwareFaultDevices.EntityData.YangName = "hardware-fault-devices"
    hardwareFaultDevices.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultDevices.EntityData.ParentYangName = "slot"
    hardwareFaultDevices.EntityData.SegmentPath = "hardware-fault-devices"
    hardwareFaultDevices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultDevices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultDevices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultDevices.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultDevices.EntityData.Children["hardware-fault-device"] = types.YChild{"HardwareFaultDevice", nil}
    for i := range hardwareFaultDevices.HardwareFaultDevice {
        hardwareFaultDevices.EntityData.Children[types.GetSegmentPath(&hardwareFaultDevices.HardwareFaultDevice[i])] = types.YChild{"HardwareFaultDevice", &hardwareFaultDevices.HardwareFaultDevice[i]}
    }
    hardwareFaultDevices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareFaultDevices.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault device list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultDevice interface{}

    // Table of Hardware Failure Type. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType.
    HardwareFaultType []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetEntityData() *types.CommonEntityData {
    hardwareFaultDevice.EntityData.YFilter = hardwareFaultDevice.YFilter
    hardwareFaultDevice.EntityData.YangName = "hardware-fault-device"
    hardwareFaultDevice.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultDevice.EntityData.ParentYangName = "hardware-fault-devices"
    hardwareFaultDevice.EntityData.SegmentPath = "hardware-fault-device" + "[hw-fault-device='" + fmt.Sprintf("%v", hardwareFaultDevice.HwFaultDevice) + "']"
    hardwareFaultDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultDevice.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultDevice.EntityData.Children["hardware-fault-type"] = types.YChild{"HardwareFaultType", nil}
    for i := range hardwareFaultDevice.HardwareFaultType {
        hardwareFaultDevice.EntityData.Children[types.GetSegmentPath(&hardwareFaultDevice.HardwareFaultType[i])] = types.YChild{"HardwareFaultType", &hardwareFaultDevice.HardwareFaultType[i]}
    }
    hardwareFaultDevice.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareFaultDevice.EntityData.Leafs["hw-fault-device"] = types.YLeaf{"HwFaultDevice", hardwareFaultDevice.HwFaultDevice}
    return &(hardwareFaultDevice.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
// Table of Hardware Failure Type
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault type list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultType interface{}

    // Faulty Hardware Condition Description. The type is string.
    ConditionDescription interface{}

    // Faulty Hardware Condition Name. The type is string.
    ConditionName interface{}

    // Faulty Hardware Device Key. The type is string.
    DeviceKey interface{}

    // Faulty Hardware Device Version. The type is interface{} with range:
    // -2147483648..2147483647.
    DeviceVersion interface{}

    // Fault Raised Timestamp. The type is string.
    ConditionRaisedTimestamp interface{}

    // Faulty Hardware Process ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Faulty Hardware Device Description. The type is string.
    DeviceDescription interface{}

    // Faulty Hardware Condition Severity. The type is string.
    ConditionSeverity interface{}
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetEntityData() *types.CommonEntityData {
    hardwareFaultType.EntityData.YFilter = hardwareFaultType.YFilter
    hardwareFaultType.EntityData.YangName = "hardware-fault-type"
    hardwareFaultType.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultType.EntityData.ParentYangName = "hardware-fault-device"
    hardwareFaultType.EntityData.SegmentPath = "hardware-fault-type" + "[hw-fault-type='" + fmt.Sprintf("%v", hardwareFaultType.HwFaultType) + "']"
    hardwareFaultType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultType.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultType.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareFaultType.EntityData.Leafs["hw-fault-type"] = types.YLeaf{"HwFaultType", hardwareFaultType.HwFaultType}
    hardwareFaultType.EntityData.Leafs["condition-description"] = types.YLeaf{"ConditionDescription", hardwareFaultType.ConditionDescription}
    hardwareFaultType.EntityData.Leafs["condition-name"] = types.YLeaf{"ConditionName", hardwareFaultType.ConditionName}
    hardwareFaultType.EntityData.Leafs["device-key"] = types.YLeaf{"DeviceKey", hardwareFaultType.DeviceKey}
    hardwareFaultType.EntityData.Leafs["device-version"] = types.YLeaf{"DeviceVersion", hardwareFaultType.DeviceVersion}
    hardwareFaultType.EntityData.Leafs["condition-raised-timestamp"] = types.YLeaf{"ConditionRaisedTimestamp", hardwareFaultType.ConditionRaisedTimestamp}
    hardwareFaultType.EntityData.Leafs["process-id"] = types.YLeaf{"ProcessId", hardwareFaultType.ProcessId}
    hardwareFaultType.EntityData.Leafs["device-description"] = types.YLeaf{"DeviceDescription", hardwareFaultType.DeviceDescription}
    hardwareFaultType.EntityData.Leafs["condition-severity"] = types.YLeaf{"ConditionSeverity", hardwareFaultType.ConditionSeverity}
    return &(hardwareFaultType.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks
// Table of racks
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack.
    Rack []PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "fault-type1"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = make(map[string]types.YChild)
    racks.EntityData.Children["rack"] = types.YChild{"Rack", nil}
    for i := range racks.Rack {
        racks.EntityData.Children[types.GetSegmentPath(&racks.Rack[i])] = types.YChild{"Rack", &racks.Rack[i]}
    }
    racks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(racks.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack
// Number
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["slots"] = types.YChild{"Slots", &rack.Slots}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["rack"] = types.YLeaf{"Rack", rack.Rack}
    return &(rack.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots
// Table of slots
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot.
    Slot []PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = make(map[string]types.YChild)
    slots.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range slots.Slot {
        slots.EntityData.Children[types.GetSegmentPath(&slots.Slot[i])] = types.YChild{"Slot", &slots.Slot[i]}
    }
    slots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slots.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot
// Name
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Slot interface{}

    // Table of Hardware Summary.
    FaultSummary PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary

    // Table of Hardware Failure.
    HardwareFaultDevices PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["fault-summary"] = types.YChild{"FaultSummary", &slot.FaultSummary}
    slot.EntityData.Children["hardware-fault-devices"] = types.YChild{"HardwareFaultDevices", &slot.HardwareFaultDevices}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["slot"] = types.YLeaf{"Slot", slot.Slot}
    return &(slot.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary
// Table of Hardware Summary
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fault Severity Critical count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityCriticalCount interface{}

    // Fault Severity Emergency count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityEmergencyOrAlertCount interface{}

    // Faulty Hardware total count. The type is interface{} with range:
    // -2147483648..2147483647.
    Total interface{}

    // Fault Severity Error count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityErrorCount interface{}
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetEntityData() *types.CommonEntityData {
    faultSummary.EntityData.YFilter = faultSummary.YFilter
    faultSummary.EntityData.YangName = "fault-summary"
    faultSummary.EntityData.BundleName = "cisco_ios_xr"
    faultSummary.EntityData.ParentYangName = "slot"
    faultSummary.EntityData.SegmentPath = "fault-summary"
    faultSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultSummary.EntityData.Children = make(map[string]types.YChild)
    faultSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    faultSummary.EntityData.Leafs["severity-critical-count"] = types.YLeaf{"SeverityCriticalCount", faultSummary.SeverityCriticalCount}
    faultSummary.EntityData.Leafs["severity-emergency-or-alert-count"] = types.YLeaf{"SeverityEmergencyOrAlertCount", faultSummary.SeverityEmergencyOrAlertCount}
    faultSummary.EntityData.Leafs["total"] = types.YLeaf{"Total", faultSummary.Total}
    faultSummary.EntityData.Leafs["severity-error-count"] = types.YLeaf{"SeverityErrorCount", faultSummary.SeverityErrorCount}
    return &(faultSummary.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices
// Table of Hardware Failure
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice.
    HardwareFaultDevice []PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetEntityData() *types.CommonEntityData {
    hardwareFaultDevices.EntityData.YFilter = hardwareFaultDevices.YFilter
    hardwareFaultDevices.EntityData.YangName = "hardware-fault-devices"
    hardwareFaultDevices.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultDevices.EntityData.ParentYangName = "slot"
    hardwareFaultDevices.EntityData.SegmentPath = "hardware-fault-devices"
    hardwareFaultDevices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultDevices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultDevices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultDevices.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultDevices.EntityData.Children["hardware-fault-device"] = types.YChild{"HardwareFaultDevice", nil}
    for i := range hardwareFaultDevices.HardwareFaultDevice {
        hardwareFaultDevices.EntityData.Children[types.GetSegmentPath(&hardwareFaultDevices.HardwareFaultDevice[i])] = types.YChild{"HardwareFaultDevice", &hardwareFaultDevices.HardwareFaultDevice[i]}
    }
    hardwareFaultDevices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareFaultDevices.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault device list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultDevice interface{}

    // Table of Hardware Failure Type. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType.
    HardwareFaultType []PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetEntityData() *types.CommonEntityData {
    hardwareFaultDevice.EntityData.YFilter = hardwareFaultDevice.YFilter
    hardwareFaultDevice.EntityData.YangName = "hardware-fault-device"
    hardwareFaultDevice.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultDevice.EntityData.ParentYangName = "hardware-fault-devices"
    hardwareFaultDevice.EntityData.SegmentPath = "hardware-fault-device" + "[hw-fault-device='" + fmt.Sprintf("%v", hardwareFaultDevice.HwFaultDevice) + "']"
    hardwareFaultDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultDevice.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultDevice.EntityData.Children["hardware-fault-type"] = types.YChild{"HardwareFaultType", nil}
    for i := range hardwareFaultDevice.HardwareFaultType {
        hardwareFaultDevice.EntityData.Children[types.GetSegmentPath(&hardwareFaultDevice.HardwareFaultType[i])] = types.YChild{"HardwareFaultType", &hardwareFaultDevice.HardwareFaultType[i]}
    }
    hardwareFaultDevice.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareFaultDevice.EntityData.Leafs["hw-fault-device"] = types.YLeaf{"HwFaultDevice", hardwareFaultDevice.HwFaultDevice}
    return &(hardwareFaultDevice.EntityData)
}

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
// Table of Hardware Failure Type
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault type list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultType interface{}

    // Faulty Hardware Condition Description. The type is string.
    ConditionDescription interface{}

    // Faulty Hardware Condition Name. The type is string.
    ConditionName interface{}

    // Faulty Hardware Device Key. The type is string.
    DeviceKey interface{}

    // Faulty Hardware Device Version. The type is interface{} with range:
    // -2147483648..2147483647.
    DeviceVersion interface{}

    // Fault Raised Timestamp. The type is string.
    ConditionRaisedTimestamp interface{}

    // Faulty Hardware Process ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Faulty Hardware Device Description. The type is string.
    DeviceDescription interface{}

    // Faulty Hardware Condition Severity. The type is string.
    ConditionSeverity interface{}
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetEntityData() *types.CommonEntityData {
    hardwareFaultType.EntityData.YFilter = hardwareFaultType.YFilter
    hardwareFaultType.EntityData.YangName = "hardware-fault-type"
    hardwareFaultType.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultType.EntityData.ParentYangName = "hardware-fault-device"
    hardwareFaultType.EntityData.SegmentPath = "hardware-fault-type" + "[hw-fault-type='" + fmt.Sprintf("%v", hardwareFaultType.HwFaultType) + "']"
    hardwareFaultType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultType.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultType.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareFaultType.EntityData.Leafs["hw-fault-type"] = types.YLeaf{"HwFaultType", hardwareFaultType.HwFaultType}
    hardwareFaultType.EntityData.Leafs["condition-description"] = types.YLeaf{"ConditionDescription", hardwareFaultType.ConditionDescription}
    hardwareFaultType.EntityData.Leafs["condition-name"] = types.YLeaf{"ConditionName", hardwareFaultType.ConditionName}
    hardwareFaultType.EntityData.Leafs["device-key"] = types.YLeaf{"DeviceKey", hardwareFaultType.DeviceKey}
    hardwareFaultType.EntityData.Leafs["device-version"] = types.YLeaf{"DeviceVersion", hardwareFaultType.DeviceVersion}
    hardwareFaultType.EntityData.Leafs["condition-raised-timestamp"] = types.YLeaf{"ConditionRaisedTimestamp", hardwareFaultType.ConditionRaisedTimestamp}
    hardwareFaultType.EntityData.Leafs["process-id"] = types.YLeaf{"ProcessId", hardwareFaultType.ProcessId}
    hardwareFaultType.EntityData.Leafs["device-description"] = types.YLeaf{"DeviceDescription", hardwareFaultType.DeviceDescription}
    hardwareFaultType.EntityData.Leafs["condition-severity"] = types.YLeaf{"ConditionSeverity", hardwareFaultType.ConditionSeverity}
    return &(hardwareFaultType.EntityData)
}

// PlatformFaultManager_Racks
// Table of racks
type PlatformFaultManager_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of PlatformFaultManager_Racks_Rack.
    Rack []PlatformFaultManager_Racks_Rack
}

func (racks *PlatformFaultManager_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "platform-fault-manager"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = make(map[string]types.YChild)
    racks.EntityData.Children["rack"] = types.YChild{"Rack", nil}
    for i := range racks.Rack {
        racks.EntityData.Children[types.GetSegmentPath(&racks.Rack[i])] = types.YChild{"Rack", &racks.Rack[i]}
    }
    racks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(racks.EntityData)
}

// PlatformFaultManager_Racks_Rack
// Number
type PlatformFaultManager_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots PlatformFaultManager_Racks_Rack_Slots
}

func (rack *PlatformFaultManager_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["slots"] = types.YChild{"Slots", &rack.Slots}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["rack"] = types.YLeaf{"Rack", rack.Rack}
    return &(rack.EntityData)
}

// PlatformFaultManager_Racks_Rack_Slots
// Table of slots
type PlatformFaultManager_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is slice of PlatformFaultManager_Racks_Rack_Slots_Slot.
    Slot []PlatformFaultManager_Racks_Rack_Slots_Slot
}

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = make(map[string]types.YChild)
    slots.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range slots.Slot {
        slots.EntityData.Children[types.GetSegmentPath(&slots.Slot[i])] = types.YChild{"Slot", &slots.Slot[i]}
    }
    slots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slots.EntityData)
}

// PlatformFaultManager_Racks_Rack_Slots_Slot
// Name
type PlatformFaultManager_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Slot interface{}

    // Table of Hardware Summary.
    FaultSummary PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary

    // Table of Hardware Failure.
    HardwareFaultDevices PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices
}

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["fault-summary"] = types.YChild{"FaultSummary", &slot.FaultSummary}
    slot.EntityData.Children["hardware-fault-devices"] = types.YChild{"HardwareFaultDevices", &slot.HardwareFaultDevices}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["slot"] = types.YLeaf{"Slot", slot.Slot}
    return &(slot.EntityData)
}

// PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary
// Table of Hardware Summary
type PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fault Severity Critical count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityCriticalCount interface{}

    // Fault Severity Emergency count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityEmergencyOrAlertCount interface{}

    // Faulty Hardware total count. The type is interface{} with range:
    // -2147483648..2147483647.
    Total interface{}

    // Fault Severity Error count. The type is interface{} with range:
    // -2147483648..2147483647.
    SeverityErrorCount interface{}
}

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetEntityData() *types.CommonEntityData {
    faultSummary.EntityData.YFilter = faultSummary.YFilter
    faultSummary.EntityData.YangName = "fault-summary"
    faultSummary.EntityData.BundleName = "cisco_ios_xr"
    faultSummary.EntityData.ParentYangName = "slot"
    faultSummary.EntityData.SegmentPath = "fault-summary"
    faultSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    faultSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    faultSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    faultSummary.EntityData.Children = make(map[string]types.YChild)
    faultSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    faultSummary.EntityData.Leafs["severity-critical-count"] = types.YLeaf{"SeverityCriticalCount", faultSummary.SeverityCriticalCount}
    faultSummary.EntityData.Leafs["severity-emergency-or-alert-count"] = types.YLeaf{"SeverityEmergencyOrAlertCount", faultSummary.SeverityEmergencyOrAlertCount}
    faultSummary.EntityData.Leafs["total"] = types.YLeaf{"Total", faultSummary.Total}
    faultSummary.EntityData.Leafs["severity-error-count"] = types.YLeaf{"SeverityErrorCount", faultSummary.SeverityErrorCount}
    return &(faultSummary.EntityData)
}

// PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices
// Table of Hardware Failure
type PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice.
    HardwareFaultDevice []PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
}

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetEntityData() *types.CommonEntityData {
    hardwareFaultDevices.EntityData.YFilter = hardwareFaultDevices.YFilter
    hardwareFaultDevices.EntityData.YangName = "hardware-fault-devices"
    hardwareFaultDevices.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultDevices.EntityData.ParentYangName = "slot"
    hardwareFaultDevices.EntityData.SegmentPath = "hardware-fault-devices"
    hardwareFaultDevices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultDevices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultDevices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultDevices.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultDevices.EntityData.Children["hardware-fault-device"] = types.YChild{"HardwareFaultDevice", nil}
    for i := range hardwareFaultDevices.HardwareFaultDevice {
        hardwareFaultDevices.EntityData.Children[types.GetSegmentPath(&hardwareFaultDevices.HardwareFaultDevice[i])] = types.YChild{"HardwareFaultDevice", &hardwareFaultDevices.HardwareFaultDevice[i]}
    }
    hardwareFaultDevices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareFaultDevices.EntityData)
}

// PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
// Table of Hardware Failure Device
type PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault device list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultDevice interface{}

    // Table of Hardware Failure Type. The type is slice of
    // PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType.
    HardwareFaultType []PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
}

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetEntityData() *types.CommonEntityData {
    hardwareFaultDevice.EntityData.YFilter = hardwareFaultDevice.YFilter
    hardwareFaultDevice.EntityData.YangName = "hardware-fault-device"
    hardwareFaultDevice.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultDevice.EntityData.ParentYangName = "hardware-fault-devices"
    hardwareFaultDevice.EntityData.SegmentPath = "hardware-fault-device" + "[hw-fault-device='" + fmt.Sprintf("%v", hardwareFaultDevice.HwFaultDevice) + "']"
    hardwareFaultDevice.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultDevice.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultDevice.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultDevice.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultDevice.EntityData.Children["hardware-fault-type"] = types.YChild{"HardwareFaultType", nil}
    for i := range hardwareFaultDevice.HardwareFaultType {
        hardwareFaultDevice.EntityData.Children[types.GetSegmentPath(&hardwareFaultDevice.HardwareFaultType[i])] = types.YChild{"HardwareFaultType", &hardwareFaultDevice.HardwareFaultType[i]}
    }
    hardwareFaultDevice.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareFaultDevice.EntityData.Leafs["hw-fault-device"] = types.YLeaf{"HwFaultDevice", hardwareFaultDevice.HwFaultDevice}
    return &(hardwareFaultDevice.EntityData)
}

// PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
// Table of Hardware Failure Type
type PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault type list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    HwFaultType interface{}

    // Faulty Hardware Condition Description. The type is string.
    ConditionDescription interface{}

    // Faulty Hardware Condition Name. The type is string.
    ConditionName interface{}

    // Faulty Hardware Device Key. The type is string.
    DeviceKey interface{}

    // Faulty Hardware Device Version. The type is interface{} with range:
    // -2147483648..2147483647.
    DeviceVersion interface{}

    // Fault Raised Timestamp. The type is string.
    ConditionRaisedTimestamp interface{}

    // Faulty Hardware Process ID. The type is interface{} with range:
    // -2147483648..2147483647.
    ProcessId interface{}

    // Faulty Hardware Device Description. The type is string.
    DeviceDescription interface{}

    // Faulty Hardware Condition Severity. The type is string.
    ConditionSeverity interface{}
}

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetEntityData() *types.CommonEntityData {
    hardwareFaultType.EntityData.YFilter = hardwareFaultType.YFilter
    hardwareFaultType.EntityData.YangName = "hardware-fault-type"
    hardwareFaultType.EntityData.BundleName = "cisco_ios_xr"
    hardwareFaultType.EntityData.ParentYangName = "hardware-fault-device"
    hardwareFaultType.EntityData.SegmentPath = "hardware-fault-type" + "[hw-fault-type='" + fmt.Sprintf("%v", hardwareFaultType.HwFaultType) + "']"
    hardwareFaultType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareFaultType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareFaultType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareFaultType.EntityData.Children = make(map[string]types.YChild)
    hardwareFaultType.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareFaultType.EntityData.Leafs["hw-fault-type"] = types.YLeaf{"HwFaultType", hardwareFaultType.HwFaultType}
    hardwareFaultType.EntityData.Leafs["condition-description"] = types.YLeaf{"ConditionDescription", hardwareFaultType.ConditionDescription}
    hardwareFaultType.EntityData.Leafs["condition-name"] = types.YLeaf{"ConditionName", hardwareFaultType.ConditionName}
    hardwareFaultType.EntityData.Leafs["device-key"] = types.YLeaf{"DeviceKey", hardwareFaultType.DeviceKey}
    hardwareFaultType.EntityData.Leafs["device-version"] = types.YLeaf{"DeviceVersion", hardwareFaultType.DeviceVersion}
    hardwareFaultType.EntityData.Leafs["condition-raised-timestamp"] = types.YLeaf{"ConditionRaisedTimestamp", hardwareFaultType.ConditionRaisedTimestamp}
    hardwareFaultType.EntityData.Leafs["process-id"] = types.YLeaf{"ProcessId", hardwareFaultType.ProcessId}
    hardwareFaultType.EntityData.Leafs["device-description"] = types.YLeaf{"DeviceDescription", hardwareFaultType.DeviceDescription}
    hardwareFaultType.EntityData.Leafs["condition-severity"] = types.YLeaf{"ConditionSeverity", hardwareFaultType.ConditionSeverity}
    return &(hardwareFaultType.EntityData)
}

