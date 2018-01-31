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
    parent types.Entity
    YFilter yfilter.YFilter

    // Exclude specic hw fault .
    Exclude PlatformFaultManager_Exclude

    // Table of racks.
    Racks PlatformFaultManager_Racks
}

func (platformFaultManager *PlatformFaultManager) GetFilter() yfilter.YFilter { return platformFaultManager.YFilter }

func (platformFaultManager *PlatformFaultManager) SetFilter(yf yfilter.YFilter) { platformFaultManager.YFilter = yf }

func (platformFaultManager *PlatformFaultManager) GetGoName(yname string) string {
    if yname == "exclude" { return "Exclude" }
    if yname == "racks" { return "Racks" }
    return ""
}

func (platformFaultManager *PlatformFaultManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-pfm-oper:platform-fault-manager"
}

func (platformFaultManager *PlatformFaultManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exclude" {
        return &platformFaultManager.Exclude
    }
    if childYangName == "racks" {
        return &platformFaultManager.Racks
    }
    return nil
}

func (platformFaultManager *PlatformFaultManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["exclude"] = &platformFaultManager.Exclude
    children["racks"] = &platformFaultManager.Racks
    return children
}

func (platformFaultManager *PlatformFaultManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformFaultManager *PlatformFaultManager) GetBundleName() string { return "cisco_ios_xr" }

func (platformFaultManager *PlatformFaultManager) GetYangName() string { return "platform-fault-manager" }

func (platformFaultManager *PlatformFaultManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformFaultManager *PlatformFaultManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformFaultManager *PlatformFaultManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformFaultManager *PlatformFaultManager) SetParent(parent types.Entity) { platformFaultManager.parent = parent }

func (platformFaultManager *PlatformFaultManager) GetParent() types.Entity { return platformFaultManager.parent }

func (platformFaultManager *PlatformFaultManager) GetParentYangName() string { return "Cisco-IOS-XR-pfm-oper" }

// PlatformFaultManager_Exclude
// Exclude specic hw fault 
type PlatformFaultManager_Exclude struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device.
    FaultType1S PlatformFaultManager_Exclude_FaultType1S
}

func (exclude *PlatformFaultManager_Exclude) GetFilter() yfilter.YFilter { return exclude.YFilter }

func (exclude *PlatformFaultManager_Exclude) SetFilter(yf yfilter.YFilter) { exclude.YFilter = yf }

func (exclude *PlatformFaultManager_Exclude) GetGoName(yname string) string {
    if yname == "fault-type1s" { return "FaultType1S" }
    return ""
}

func (exclude *PlatformFaultManager_Exclude) GetSegmentPath() string {
    return "exclude"
}

func (exclude *PlatformFaultManager_Exclude) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-type1s" {
        return &exclude.FaultType1S
    }
    return nil
}

func (exclude *PlatformFaultManager_Exclude) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fault-type1s"] = &exclude.FaultType1S
    return children
}

func (exclude *PlatformFaultManager_Exclude) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (exclude *PlatformFaultManager_Exclude) GetBundleName() string { return "cisco_ios_xr" }

func (exclude *PlatformFaultManager_Exclude) GetYangName() string { return "exclude" }

func (exclude *PlatformFaultManager_Exclude) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exclude *PlatformFaultManager_Exclude) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exclude *PlatformFaultManager_Exclude) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exclude *PlatformFaultManager_Exclude) SetParent(parent types.Entity) { exclude.parent = parent }

func (exclude *PlatformFaultManager_Exclude) GetParent() types.Entity { return exclude.parent }

func (exclude *PlatformFaultManager_Exclude) GetParentYangName() string { return "platform-fault-manager" }

// PlatformFaultManager_Exclude_FaultType1S
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1.
    FaultType1 []PlatformFaultManager_Exclude_FaultType1S_FaultType1
}

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetFilter() yfilter.YFilter { return faultType1S.YFilter }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) SetFilter(yf yfilter.YFilter) { faultType1S.YFilter = yf }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetGoName(yname string) string {
    if yname == "fault-type1" { return "FaultType1" }
    return ""
}

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetSegmentPath() string {
    return "fault-type1s"
}

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-type1" {
        for _, c := range faultType1S.FaultType1 {
            if faultType1S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1{}
        faultType1S.FaultType1 = append(faultType1S.FaultType1, child)
        return &faultType1S.FaultType1[len(faultType1S.FaultType1)-1]
    }
    return nil
}

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range faultType1S.FaultType1 {
        children[faultType1S.FaultType1[i].GetSegmentPath()] = &faultType1S.FaultType1[i]
    }
    return children
}

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetBundleName() string { return "cisco_ios_xr" }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetYangName() string { return "fault-type1s" }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) SetParent(parent types.Entity) { faultType1S.parent = parent }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetParent() types.Entity { return faultType1S.parent }

func (faultType1S *PlatformFaultManager_Exclude_FaultType1S) GetParentYangName() string { return "exclude" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault 1. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    HwFaultType1 interface{}

    // Table of Hardware Failure Device.
    FaultType2S PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S

    // Table of racks.
    Racks PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks
}

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetFilter() yfilter.YFilter { return faultType1.YFilter }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) SetFilter(yf yfilter.YFilter) { faultType1.YFilter = yf }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetGoName(yname string) string {
    if yname == "hw-fault-type1" { return "HwFaultType1" }
    if yname == "fault-type2s" { return "FaultType2S" }
    if yname == "racks" { return "Racks" }
    return ""
}

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetSegmentPath() string {
    return "fault-type1" + "[hw-fault-type1='" + fmt.Sprintf("%v", faultType1.HwFaultType1) + "']"
}

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-type2s" {
        return &faultType1.FaultType2S
    }
    if childYangName == "racks" {
        return &faultType1.Racks
    }
    return nil
}

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fault-type2s"] = &faultType1.FaultType2S
    children["racks"] = &faultType1.Racks
    return children
}

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-type1"] = faultType1.HwFaultType1
    return leafs
}

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetBundleName() string { return "cisco_ios_xr" }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetYangName() string { return "fault-type1" }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) SetParent(parent types.Entity) { faultType1.parent = parent }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetParent() types.Entity { return faultType1.parent }

func (faultType1 *PlatformFaultManager_Exclude_FaultType1S_FaultType1) GetParentYangName() string { return "fault-type1s" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2.
    FaultType2 []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2
}

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetFilter() yfilter.YFilter { return faultType2S.YFilter }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) SetFilter(yf yfilter.YFilter) { faultType2S.YFilter = yf }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetGoName(yname string) string {
    if yname == "fault-type2" { return "FaultType2" }
    return ""
}

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetSegmentPath() string {
    return "fault-type2s"
}

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-type2" {
        for _, c := range faultType2S.FaultType2 {
            if faultType2S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2{}
        faultType2S.FaultType2 = append(faultType2S.FaultType2, child)
        return &faultType2S.FaultType2[len(faultType2S.FaultType2)-1]
    }
    return nil
}

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range faultType2S.FaultType2 {
        children[faultType2S.FaultType2[i].GetSegmentPath()] = &faultType2S.FaultType2[i]
    }
    return children
}

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetBundleName() string { return "cisco_ios_xr" }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetYangName() string { return "fault-type2s" }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) SetParent(parent types.Entity) { faultType2S.parent = parent }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetParent() types.Entity { return faultType2S.parent }

func (faultType2S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S) GetParentYangName() string { return "fault-type1" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault 2. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    HwFaultType2 interface{}

    // Table of Hardware Failure Device.
    FaultType3S PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S

    // Table of racks.
    Racks PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks
}

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetFilter() yfilter.YFilter { return faultType2.YFilter }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) SetFilter(yf yfilter.YFilter) { faultType2.YFilter = yf }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetGoName(yname string) string {
    if yname == "hw-fault-type2" { return "HwFaultType2" }
    if yname == "fault-type3s" { return "FaultType3S" }
    if yname == "racks" { return "Racks" }
    return ""
}

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetSegmentPath() string {
    return "fault-type2" + "[hw-fault-type2='" + fmt.Sprintf("%v", faultType2.HwFaultType2) + "']"
}

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-type3s" {
        return &faultType2.FaultType3S
    }
    if childYangName == "racks" {
        return &faultType2.Racks
    }
    return nil
}

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fault-type3s"] = &faultType2.FaultType3S
    children["racks"] = &faultType2.Racks
    return children
}

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-type2"] = faultType2.HwFaultType2
    return leafs
}

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetBundleName() string { return "cisco_ios_xr" }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetYangName() string { return "fault-type2" }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) SetParent(parent types.Entity) { faultType2.parent = parent }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetParent() types.Entity { return faultType2.parent }

func (faultType2 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2) GetParentYangName() string { return "fault-type2s" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3.
    FaultType3 []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3
}

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetFilter() yfilter.YFilter { return faultType3S.YFilter }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) SetFilter(yf yfilter.YFilter) { faultType3S.YFilter = yf }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetGoName(yname string) string {
    if yname == "fault-type3" { return "FaultType3" }
    return ""
}

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetSegmentPath() string {
    return "fault-type3s"
}

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-type3" {
        for _, c := range faultType3S.FaultType3 {
            if faultType3S.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3{}
        faultType3S.FaultType3 = append(faultType3S.FaultType3, child)
        return &faultType3S.FaultType3[len(faultType3S.FaultType3)-1]
    }
    return nil
}

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range faultType3S.FaultType3 {
        children[faultType3S.FaultType3[i].GetSegmentPath()] = &faultType3S.FaultType3[i]
    }
    return children
}

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetBundleName() string { return "cisco_ios_xr" }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetYangName() string { return "fault-type3s" }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) SetParent(parent types.Entity) { faultType3S.parent = parent }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetParent() types.Entity { return faultType3S.parent }

func (faultType3S *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S) GetParentYangName() string { return "fault-type2" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault 3. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    HwFaultType3 interface{}

    // Table of racks.
    Racks PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks
}

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetFilter() yfilter.YFilter { return faultType3.YFilter }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) SetFilter(yf yfilter.YFilter) { faultType3.YFilter = yf }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetGoName(yname string) string {
    if yname == "hw-fault-type3" { return "HwFaultType3" }
    if yname == "racks" { return "Racks" }
    return ""
}

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetSegmentPath() string {
    return "fault-type3" + "[hw-fault-type3='" + fmt.Sprintf("%v", faultType3.HwFaultType3) + "']"
}

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "racks" {
        return &faultType3.Racks
    }
    return nil
}

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["racks"] = &faultType3.Racks
    return children
}

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-type3"] = faultType3.HwFaultType3
    return leafs
}

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetBundleName() string { return "cisco_ios_xr" }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetYangName() string { return "fault-type3" }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) SetParent(parent types.Entity) { faultType3.parent = parent }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetParent() types.Entity { return faultType3.parent }

func (faultType3 *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3) GetParentYangName() string { return "fault-type3s" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks
// Table of racks
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack.
    Rack []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetYangName() string { return "racks" }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetParent() types.Entity { return racks.parent }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks) GetParentYangName() string { return "fault-type3" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack
// Number
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    if yname == "slots" { return "Slots" }
    return ""
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slots" {
        return &rack.Slots
    }
    return nil
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slots"] = &rack.Slots
    return children
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack"] = rack.Rack
    return leafs
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetYangName() string { return "rack" }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack) GetParentYangName() string { return "racks" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots
// Table of slots
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot.
    Slot []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot
// Name
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Slot interface{}

    // Table of Hardware Summary.
    FaultSummary PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary

    // Table of Hardware Failure.
    HardwareFaultDevices PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    if yname == "fault-summary" { return "FaultSummary" }
    if yname == "hardware-fault-devices" { return "HardwareFaultDevices" }
    return ""
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-summary" {
        return &slot.FaultSummary
    }
    if childYangName == "hardware-fault-devices" {
        return &slot.HardwareFaultDevices
    }
    return nil
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fault-summary"] = &slot.FaultSummary
    children["hardware-fault-devices"] = &slot.HardwareFaultDevices
    return children
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot"] = slot.Slot
    return leafs
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary
// Table of Hardware Summary
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary struct {
    parent types.Entity
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

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetFilter() yfilter.YFilter { return faultSummary.YFilter }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) SetFilter(yf yfilter.YFilter) { faultSummary.YFilter = yf }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetGoName(yname string) string {
    if yname == "severity-critical-count" { return "SeverityCriticalCount" }
    if yname == "severity-emergency-or-alert-count" { return "SeverityEmergencyOrAlertCount" }
    if yname == "total" { return "Total" }
    if yname == "severity-error-count" { return "SeverityErrorCount" }
    return ""
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetSegmentPath() string {
    return "fault-summary"
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity-critical-count"] = faultSummary.SeverityCriticalCount
    leafs["severity-emergency-or-alert-count"] = faultSummary.SeverityEmergencyOrAlertCount
    leafs["total"] = faultSummary.Total
    leafs["severity-error-count"] = faultSummary.SeverityErrorCount
    return leafs
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetBundleName() string { return "cisco_ios_xr" }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetYangName() string { return "fault-summary" }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) SetParent(parent types.Entity) { faultSummary.parent = parent }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetParent() types.Entity { return faultSummary.parent }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_FaultSummary) GetParentYangName() string { return "slot" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices
// Table of Hardware Failure
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice.
    HardwareFaultDevice []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetFilter() yfilter.YFilter { return hardwareFaultDevices.YFilter }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) SetFilter(yf yfilter.YFilter) { hardwareFaultDevices.YFilter = yf }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetGoName(yname string) string {
    if yname == "hardware-fault-device" { return "HardwareFaultDevice" }
    return ""
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetSegmentPath() string {
    return "hardware-fault-devices"
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-fault-device" {
        for _, c := range hardwareFaultDevices.HardwareFaultDevice {
            if hardwareFaultDevices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice{}
        hardwareFaultDevices.HardwareFaultDevice = append(hardwareFaultDevices.HardwareFaultDevice, child)
        return &hardwareFaultDevices.HardwareFaultDevice[len(hardwareFaultDevices.HardwareFaultDevice)-1]
    }
    return nil
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareFaultDevices.HardwareFaultDevice {
        children[hardwareFaultDevices.HardwareFaultDevice[i].GetSegmentPath()] = &hardwareFaultDevices.HardwareFaultDevice[i]
    }
    return children
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetYangName() string { return "hardware-fault-devices" }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) SetParent(parent types.Entity) { hardwareFaultDevices.parent = parent }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetParent() types.Entity { return hardwareFaultDevices.parent }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetParentYangName() string { return "slot" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault device list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    HwFaultDevice interface{}

    // Table of Hardware Failure Type. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType.
    HardwareFaultType []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetFilter() yfilter.YFilter { return hardwareFaultDevice.YFilter }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) SetFilter(yf yfilter.YFilter) { hardwareFaultDevice.YFilter = yf }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetGoName(yname string) string {
    if yname == "hw-fault-device" { return "HwFaultDevice" }
    if yname == "hardware-fault-type" { return "HardwareFaultType" }
    return ""
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetSegmentPath() string {
    return "hardware-fault-device" + "[hw-fault-device='" + fmt.Sprintf("%v", hardwareFaultDevice.HwFaultDevice) + "']"
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-fault-type" {
        for _, c := range hardwareFaultDevice.HardwareFaultType {
            if hardwareFaultDevice.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType{}
        hardwareFaultDevice.HardwareFaultType = append(hardwareFaultDevice.HardwareFaultType, child)
        return &hardwareFaultDevice.HardwareFaultType[len(hardwareFaultDevice.HardwareFaultType)-1]
    }
    return nil
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareFaultDevice.HardwareFaultType {
        children[hardwareFaultDevice.HardwareFaultType[i].GetSegmentPath()] = &hardwareFaultDevice.HardwareFaultType[i]
    }
    return children
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-device"] = hardwareFaultDevice.HwFaultDevice
    return leafs
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetYangName() string { return "hardware-fault-device" }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) SetParent(parent types.Entity) { hardwareFaultDevice.parent = parent }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetParent() types.Entity { return hardwareFaultDevice.parent }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetParentYangName() string { return "hardware-fault-devices" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
// Table of Hardware Failure Type
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault type list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetFilter() yfilter.YFilter { return hardwareFaultType.YFilter }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) SetFilter(yf yfilter.YFilter) { hardwareFaultType.YFilter = yf }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetGoName(yname string) string {
    if yname == "hw-fault-type" { return "HwFaultType" }
    if yname == "condition-description" { return "ConditionDescription" }
    if yname == "condition-name" { return "ConditionName" }
    if yname == "device-key" { return "DeviceKey" }
    if yname == "device-version" { return "DeviceVersion" }
    if yname == "condition-raised-timestamp" { return "ConditionRaisedTimestamp" }
    if yname == "process-id" { return "ProcessId" }
    if yname == "device-description" { return "DeviceDescription" }
    if yname == "condition-severity" { return "ConditionSeverity" }
    return ""
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetSegmentPath() string {
    return "hardware-fault-type" + "[hw-fault-type='" + fmt.Sprintf("%v", hardwareFaultType.HwFaultType) + "']"
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-type"] = hardwareFaultType.HwFaultType
    leafs["condition-description"] = hardwareFaultType.ConditionDescription
    leafs["condition-name"] = hardwareFaultType.ConditionName
    leafs["device-key"] = hardwareFaultType.DeviceKey
    leafs["device-version"] = hardwareFaultType.DeviceVersion
    leafs["condition-raised-timestamp"] = hardwareFaultType.ConditionRaisedTimestamp
    leafs["process-id"] = hardwareFaultType.ProcessId
    leafs["device-description"] = hardwareFaultType.DeviceDescription
    leafs["condition-severity"] = hardwareFaultType.ConditionSeverity
    return leafs
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetYangName() string { return "hardware-fault-type" }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) SetParent(parent types.Entity) { hardwareFaultType.parent = parent }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetParent() types.Entity { return hardwareFaultType.parent }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_FaultType3S_FaultType3_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetParentYangName() string { return "hardware-fault-device" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks
// Table of racks
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack.
    Rack []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetYangName() string { return "racks" }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetParent() types.Entity { return racks.parent }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks) GetParentYangName() string { return "fault-type2" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack
// Number
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    if yname == "slots" { return "Slots" }
    return ""
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slots" {
        return &rack.Slots
    }
    return nil
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slots"] = &rack.Slots
    return children
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack"] = rack.Rack
    return leafs
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetYangName() string { return "rack" }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack) GetParentYangName() string { return "racks" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots
// Table of slots
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot.
    Slot []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot
// Name
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Slot interface{}

    // Table of Hardware Summary.
    FaultSummary PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary

    // Table of Hardware Failure.
    HardwareFaultDevices PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    if yname == "fault-summary" { return "FaultSummary" }
    if yname == "hardware-fault-devices" { return "HardwareFaultDevices" }
    return ""
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-summary" {
        return &slot.FaultSummary
    }
    if childYangName == "hardware-fault-devices" {
        return &slot.HardwareFaultDevices
    }
    return nil
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fault-summary"] = &slot.FaultSummary
    children["hardware-fault-devices"] = &slot.HardwareFaultDevices
    return children
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot"] = slot.Slot
    return leafs
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary
// Table of Hardware Summary
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary struct {
    parent types.Entity
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

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetFilter() yfilter.YFilter { return faultSummary.YFilter }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) SetFilter(yf yfilter.YFilter) { faultSummary.YFilter = yf }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetGoName(yname string) string {
    if yname == "severity-critical-count" { return "SeverityCriticalCount" }
    if yname == "severity-emergency-or-alert-count" { return "SeverityEmergencyOrAlertCount" }
    if yname == "total" { return "Total" }
    if yname == "severity-error-count" { return "SeverityErrorCount" }
    return ""
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetSegmentPath() string {
    return "fault-summary"
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity-critical-count"] = faultSummary.SeverityCriticalCount
    leafs["severity-emergency-or-alert-count"] = faultSummary.SeverityEmergencyOrAlertCount
    leafs["total"] = faultSummary.Total
    leafs["severity-error-count"] = faultSummary.SeverityErrorCount
    return leafs
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetBundleName() string { return "cisco_ios_xr" }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetYangName() string { return "fault-summary" }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) SetParent(parent types.Entity) { faultSummary.parent = parent }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetParent() types.Entity { return faultSummary.parent }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_FaultSummary) GetParentYangName() string { return "slot" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices
// Table of Hardware Failure
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice.
    HardwareFaultDevice []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetFilter() yfilter.YFilter { return hardwareFaultDevices.YFilter }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) SetFilter(yf yfilter.YFilter) { hardwareFaultDevices.YFilter = yf }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetGoName(yname string) string {
    if yname == "hardware-fault-device" { return "HardwareFaultDevice" }
    return ""
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetSegmentPath() string {
    return "hardware-fault-devices"
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-fault-device" {
        for _, c := range hardwareFaultDevices.HardwareFaultDevice {
            if hardwareFaultDevices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice{}
        hardwareFaultDevices.HardwareFaultDevice = append(hardwareFaultDevices.HardwareFaultDevice, child)
        return &hardwareFaultDevices.HardwareFaultDevice[len(hardwareFaultDevices.HardwareFaultDevice)-1]
    }
    return nil
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareFaultDevices.HardwareFaultDevice {
        children[hardwareFaultDevices.HardwareFaultDevice[i].GetSegmentPath()] = &hardwareFaultDevices.HardwareFaultDevice[i]
    }
    return children
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetYangName() string { return "hardware-fault-devices" }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) SetParent(parent types.Entity) { hardwareFaultDevices.parent = parent }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetParent() types.Entity { return hardwareFaultDevices.parent }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetParentYangName() string { return "slot" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault device list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    HwFaultDevice interface{}

    // Table of Hardware Failure Type. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType.
    HardwareFaultType []PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetFilter() yfilter.YFilter { return hardwareFaultDevice.YFilter }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) SetFilter(yf yfilter.YFilter) { hardwareFaultDevice.YFilter = yf }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetGoName(yname string) string {
    if yname == "hw-fault-device" { return "HwFaultDevice" }
    if yname == "hardware-fault-type" { return "HardwareFaultType" }
    return ""
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetSegmentPath() string {
    return "hardware-fault-device" + "[hw-fault-device='" + fmt.Sprintf("%v", hardwareFaultDevice.HwFaultDevice) + "']"
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-fault-type" {
        for _, c := range hardwareFaultDevice.HardwareFaultType {
            if hardwareFaultDevice.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType{}
        hardwareFaultDevice.HardwareFaultType = append(hardwareFaultDevice.HardwareFaultType, child)
        return &hardwareFaultDevice.HardwareFaultType[len(hardwareFaultDevice.HardwareFaultType)-1]
    }
    return nil
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareFaultDevice.HardwareFaultType {
        children[hardwareFaultDevice.HardwareFaultType[i].GetSegmentPath()] = &hardwareFaultDevice.HardwareFaultType[i]
    }
    return children
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-device"] = hardwareFaultDevice.HwFaultDevice
    return leafs
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetYangName() string { return "hardware-fault-device" }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) SetParent(parent types.Entity) { hardwareFaultDevice.parent = parent }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetParent() types.Entity { return hardwareFaultDevice.parent }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetParentYangName() string { return "hardware-fault-devices" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
// Table of Hardware Failure Type
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault type list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetFilter() yfilter.YFilter { return hardwareFaultType.YFilter }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) SetFilter(yf yfilter.YFilter) { hardwareFaultType.YFilter = yf }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetGoName(yname string) string {
    if yname == "hw-fault-type" { return "HwFaultType" }
    if yname == "condition-description" { return "ConditionDescription" }
    if yname == "condition-name" { return "ConditionName" }
    if yname == "device-key" { return "DeviceKey" }
    if yname == "device-version" { return "DeviceVersion" }
    if yname == "condition-raised-timestamp" { return "ConditionRaisedTimestamp" }
    if yname == "process-id" { return "ProcessId" }
    if yname == "device-description" { return "DeviceDescription" }
    if yname == "condition-severity" { return "ConditionSeverity" }
    return ""
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetSegmentPath() string {
    return "hardware-fault-type" + "[hw-fault-type='" + fmt.Sprintf("%v", hardwareFaultType.HwFaultType) + "']"
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-type"] = hardwareFaultType.HwFaultType
    leafs["condition-description"] = hardwareFaultType.ConditionDescription
    leafs["condition-name"] = hardwareFaultType.ConditionName
    leafs["device-key"] = hardwareFaultType.DeviceKey
    leafs["device-version"] = hardwareFaultType.DeviceVersion
    leafs["condition-raised-timestamp"] = hardwareFaultType.ConditionRaisedTimestamp
    leafs["process-id"] = hardwareFaultType.ProcessId
    leafs["device-description"] = hardwareFaultType.DeviceDescription
    leafs["condition-severity"] = hardwareFaultType.ConditionSeverity
    return leafs
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetYangName() string { return "hardware-fault-type" }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) SetParent(parent types.Entity) { hardwareFaultType.parent = parent }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetParent() types.Entity { return hardwareFaultType.parent }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_FaultType2S_FaultType2_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetParentYangName() string { return "hardware-fault-device" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks
// Table of racks
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack.
    Rack []PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetYangName() string { return "racks" }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetParent() types.Entity { return racks.parent }

func (racks *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks) GetParentYangName() string { return "fault-type1" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack
// Number
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    if yname == "slots" { return "Slots" }
    return ""
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slots" {
        return &rack.Slots
    }
    return nil
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slots"] = &rack.Slots
    return children
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack"] = rack.Rack
    return leafs
}

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetYangName() string { return "rack" }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack) GetParentYangName() string { return "racks" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots
// Table of slots
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot.
    Slot []PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot
// Name
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Slot interface{}

    // Table of Hardware Summary.
    FaultSummary PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary

    // Table of Hardware Failure.
    HardwareFaultDevices PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    if yname == "fault-summary" { return "FaultSummary" }
    if yname == "hardware-fault-devices" { return "HardwareFaultDevices" }
    return ""
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-summary" {
        return &slot.FaultSummary
    }
    if childYangName == "hardware-fault-devices" {
        return &slot.HardwareFaultDevices
    }
    return nil
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fault-summary"] = &slot.FaultSummary
    children["hardware-fault-devices"] = &slot.HardwareFaultDevices
    return children
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot"] = slot.Slot
    return leafs
}

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary
// Table of Hardware Summary
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary struct {
    parent types.Entity
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

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetFilter() yfilter.YFilter { return faultSummary.YFilter }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) SetFilter(yf yfilter.YFilter) { faultSummary.YFilter = yf }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetGoName(yname string) string {
    if yname == "severity-critical-count" { return "SeverityCriticalCount" }
    if yname == "severity-emergency-or-alert-count" { return "SeverityEmergencyOrAlertCount" }
    if yname == "total" { return "Total" }
    if yname == "severity-error-count" { return "SeverityErrorCount" }
    return ""
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetSegmentPath() string {
    return "fault-summary"
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity-critical-count"] = faultSummary.SeverityCriticalCount
    leafs["severity-emergency-or-alert-count"] = faultSummary.SeverityEmergencyOrAlertCount
    leafs["total"] = faultSummary.Total
    leafs["severity-error-count"] = faultSummary.SeverityErrorCount
    return leafs
}

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetBundleName() string { return "cisco_ios_xr" }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetYangName() string { return "fault-summary" }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) SetParent(parent types.Entity) { faultSummary.parent = parent }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetParent() types.Entity { return faultSummary.parent }

func (faultSummary *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_FaultSummary) GetParentYangName() string { return "slot" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices
// Table of Hardware Failure
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice.
    HardwareFaultDevice []PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetFilter() yfilter.YFilter { return hardwareFaultDevices.YFilter }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) SetFilter(yf yfilter.YFilter) { hardwareFaultDevices.YFilter = yf }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetGoName(yname string) string {
    if yname == "hardware-fault-device" { return "HardwareFaultDevice" }
    return ""
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetSegmentPath() string {
    return "hardware-fault-devices"
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-fault-device" {
        for _, c := range hardwareFaultDevices.HardwareFaultDevice {
            if hardwareFaultDevices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice{}
        hardwareFaultDevices.HardwareFaultDevice = append(hardwareFaultDevices.HardwareFaultDevice, child)
        return &hardwareFaultDevices.HardwareFaultDevice[len(hardwareFaultDevices.HardwareFaultDevice)-1]
    }
    return nil
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareFaultDevices.HardwareFaultDevice {
        children[hardwareFaultDevices.HardwareFaultDevice[i].GetSegmentPath()] = &hardwareFaultDevices.HardwareFaultDevice[i]
    }
    return children
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetYangName() string { return "hardware-fault-devices" }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) SetParent(parent types.Entity) { hardwareFaultDevices.parent = parent }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetParent() types.Entity { return hardwareFaultDevices.parent }

func (hardwareFaultDevices *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetParentYangName() string { return "slot" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
// Table of Hardware Failure Device
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault device list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    HwFaultDevice interface{}

    // Table of Hardware Failure Type. The type is slice of
    // PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType.
    HardwareFaultType []PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetFilter() yfilter.YFilter { return hardwareFaultDevice.YFilter }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) SetFilter(yf yfilter.YFilter) { hardwareFaultDevice.YFilter = yf }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetGoName(yname string) string {
    if yname == "hw-fault-device" { return "HwFaultDevice" }
    if yname == "hardware-fault-type" { return "HardwareFaultType" }
    return ""
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetSegmentPath() string {
    return "hardware-fault-device" + "[hw-fault-device='" + fmt.Sprintf("%v", hardwareFaultDevice.HwFaultDevice) + "']"
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-fault-type" {
        for _, c := range hardwareFaultDevice.HardwareFaultType {
            if hardwareFaultDevice.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType{}
        hardwareFaultDevice.HardwareFaultType = append(hardwareFaultDevice.HardwareFaultType, child)
        return &hardwareFaultDevice.HardwareFaultType[len(hardwareFaultDevice.HardwareFaultType)-1]
    }
    return nil
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareFaultDevice.HardwareFaultType {
        children[hardwareFaultDevice.HardwareFaultType[i].GetSegmentPath()] = &hardwareFaultDevice.HardwareFaultType[i]
    }
    return children
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-device"] = hardwareFaultDevice.HwFaultDevice
    return leafs
}

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetYangName() string { return "hardware-fault-device" }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) SetParent(parent types.Entity) { hardwareFaultDevice.parent = parent }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetParent() types.Entity { return hardwareFaultDevice.parent }

func (hardwareFaultDevice *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetParentYangName() string { return "hardware-fault-devices" }

// PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
// Table of Hardware Failure Type
type PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault type list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetFilter() yfilter.YFilter { return hardwareFaultType.YFilter }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) SetFilter(yf yfilter.YFilter) { hardwareFaultType.YFilter = yf }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetGoName(yname string) string {
    if yname == "hw-fault-type" { return "HwFaultType" }
    if yname == "condition-description" { return "ConditionDescription" }
    if yname == "condition-name" { return "ConditionName" }
    if yname == "device-key" { return "DeviceKey" }
    if yname == "device-version" { return "DeviceVersion" }
    if yname == "condition-raised-timestamp" { return "ConditionRaisedTimestamp" }
    if yname == "process-id" { return "ProcessId" }
    if yname == "device-description" { return "DeviceDescription" }
    if yname == "condition-severity" { return "ConditionSeverity" }
    return ""
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetSegmentPath() string {
    return "hardware-fault-type" + "[hw-fault-type='" + fmt.Sprintf("%v", hardwareFaultType.HwFaultType) + "']"
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-type"] = hardwareFaultType.HwFaultType
    leafs["condition-description"] = hardwareFaultType.ConditionDescription
    leafs["condition-name"] = hardwareFaultType.ConditionName
    leafs["device-key"] = hardwareFaultType.DeviceKey
    leafs["device-version"] = hardwareFaultType.DeviceVersion
    leafs["condition-raised-timestamp"] = hardwareFaultType.ConditionRaisedTimestamp
    leafs["process-id"] = hardwareFaultType.ProcessId
    leafs["device-description"] = hardwareFaultType.DeviceDescription
    leafs["condition-severity"] = hardwareFaultType.ConditionSeverity
    return leafs
}

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetYangName() string { return "hardware-fault-type" }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) SetParent(parent types.Entity) { hardwareFaultType.parent = parent }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetParent() types.Entity { return hardwareFaultType.parent }

func (hardwareFaultType *PlatformFaultManager_Exclude_FaultType1S_FaultType1_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetParentYangName() string { return "hardware-fault-device" }

// PlatformFaultManager_Racks
// Table of racks
type PlatformFaultManager_Racks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of PlatformFaultManager_Racks_Rack.
    Rack []PlatformFaultManager_Racks_Rack
}

func (racks *PlatformFaultManager_Racks) GetFilter() yfilter.YFilter { return racks.YFilter }

func (racks *PlatformFaultManager_Racks) SetFilter(yf yfilter.YFilter) { racks.YFilter = yf }

func (racks *PlatformFaultManager_Racks) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    return ""
}

func (racks *PlatformFaultManager_Racks) GetSegmentPath() string {
    return "racks"
}

func (racks *PlatformFaultManager_Racks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rack" {
        for _, c := range racks.Rack {
            if racks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Racks_Rack{}
        racks.Rack = append(racks.Rack, child)
        return &racks.Rack[len(racks.Rack)-1]
    }
    return nil
}

func (racks *PlatformFaultManager_Racks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range racks.Rack {
        children[racks.Rack[i].GetSegmentPath()] = &racks.Rack[i]
    }
    return children
}

func (racks *PlatformFaultManager_Racks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (racks *PlatformFaultManager_Racks) GetBundleName() string { return "cisco_ios_xr" }

func (racks *PlatformFaultManager_Racks) GetYangName() string { return "racks" }

func (racks *PlatformFaultManager_Racks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (racks *PlatformFaultManager_Racks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (racks *PlatformFaultManager_Racks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (racks *PlatformFaultManager_Racks) SetParent(parent types.Entity) { racks.parent = parent }

func (racks *PlatformFaultManager_Racks) GetParent() types.Entity { return racks.parent }

func (racks *PlatformFaultManager_Racks) GetParentYangName() string { return "platform-fault-manager" }

// PlatformFaultManager_Racks_Rack
// Number
type PlatformFaultManager_Racks_Rack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rack number. The type is interface{} with range:
    // -2147483648..2147483647.
    Rack interface{}

    // Table of slots.
    Slots PlatformFaultManager_Racks_Rack_Slots
}

func (rack *PlatformFaultManager_Racks_Rack) GetFilter() yfilter.YFilter { return rack.YFilter }

func (rack *PlatformFaultManager_Racks_Rack) SetFilter(yf yfilter.YFilter) { rack.YFilter = yf }

func (rack *PlatformFaultManager_Racks_Rack) GetGoName(yname string) string {
    if yname == "rack" { return "Rack" }
    if yname == "slots" { return "Slots" }
    return ""
}

func (rack *PlatformFaultManager_Racks_Rack) GetSegmentPath() string {
    return "rack" + "[rack='" + fmt.Sprintf("%v", rack.Rack) + "']"
}

func (rack *PlatformFaultManager_Racks_Rack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slots" {
        return &rack.Slots
    }
    return nil
}

func (rack *PlatformFaultManager_Racks_Rack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slots"] = &rack.Slots
    return children
}

func (rack *PlatformFaultManager_Racks_Rack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack"] = rack.Rack
    return leafs
}

func (rack *PlatformFaultManager_Racks_Rack) GetBundleName() string { return "cisco_ios_xr" }

func (rack *PlatformFaultManager_Racks_Rack) GetYangName() string { return "rack" }

func (rack *PlatformFaultManager_Racks_Rack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rack *PlatformFaultManager_Racks_Rack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rack *PlatformFaultManager_Racks_Rack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rack *PlatformFaultManager_Racks_Rack) SetParent(parent types.Entity) { rack.parent = parent }

func (rack *PlatformFaultManager_Racks_Rack) GetParent() types.Entity { return rack.parent }

func (rack *PlatformFaultManager_Racks_Rack) GetParentYangName() string { return "racks" }

// PlatformFaultManager_Racks_Rack_Slots
// Table of slots
type PlatformFaultManager_Racks_Rack_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name. The type is slice of PlatformFaultManager_Racks_Rack_Slots_Slot.
    Slot []PlatformFaultManager_Racks_Rack_Slots_Slot
}

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *PlatformFaultManager_Racks_Rack_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Racks_Rack_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetYangName() string { return "slots" }

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *PlatformFaultManager_Racks_Rack_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetParent() types.Entity { return slots.parent }

func (slots *PlatformFaultManager_Racks_Rack_Slots) GetParentYangName() string { return "rack" }

// PlatformFaultManager_Racks_Rack_Slots_Slot
// Name
type PlatformFaultManager_Racks_Rack_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Slot interface{}

    // Table of Hardware Summary.
    FaultSummary PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary

    // Table of Hardware Failure.
    HardwareFaultDevices PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices
}

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    if yname == "fault-summary" { return "FaultSummary" }
    if yname == "hardware-fault-devices" { return "HardwareFaultDevices" }
    return ""
}

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
}

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fault-summary" {
        return &slot.FaultSummary
    }
    if childYangName == "hardware-fault-devices" {
        return &slot.HardwareFaultDevices
    }
    return nil
}

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fault-summary"] = &slot.FaultSummary
    children["hardware-fault-devices"] = &slot.HardwareFaultDevices
    return children
}

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot"] = slot.Slot
    return leafs
}

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetYangName() string { return "slot" }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *PlatformFaultManager_Racks_Rack_Slots_Slot) GetParentYangName() string { return "slots" }

// PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary
// Table of Hardware Summary
type PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary struct {
    parent types.Entity
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

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetFilter() yfilter.YFilter { return faultSummary.YFilter }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) SetFilter(yf yfilter.YFilter) { faultSummary.YFilter = yf }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetGoName(yname string) string {
    if yname == "severity-critical-count" { return "SeverityCriticalCount" }
    if yname == "severity-emergency-or-alert-count" { return "SeverityEmergencyOrAlertCount" }
    if yname == "total" { return "Total" }
    if yname == "severity-error-count" { return "SeverityErrorCount" }
    return ""
}

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetSegmentPath() string {
    return "fault-summary"
}

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity-critical-count"] = faultSummary.SeverityCriticalCount
    leafs["severity-emergency-or-alert-count"] = faultSummary.SeverityEmergencyOrAlertCount
    leafs["total"] = faultSummary.Total
    leafs["severity-error-count"] = faultSummary.SeverityErrorCount
    return leafs
}

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetBundleName() string { return "cisco_ios_xr" }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetYangName() string { return "fault-summary" }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) SetParent(parent types.Entity) { faultSummary.parent = parent }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetParent() types.Entity { return faultSummary.parent }

func (faultSummary *PlatformFaultManager_Racks_Rack_Slots_Slot_FaultSummary) GetParentYangName() string { return "slot" }

// PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices
// Table of Hardware Failure
type PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Hardware Failure Device. The type is slice of
    // PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice.
    HardwareFaultDevice []PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
}

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetFilter() yfilter.YFilter { return hardwareFaultDevices.YFilter }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) SetFilter(yf yfilter.YFilter) { hardwareFaultDevices.YFilter = yf }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetGoName(yname string) string {
    if yname == "hardware-fault-device" { return "HardwareFaultDevice" }
    return ""
}

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetSegmentPath() string {
    return "hardware-fault-devices"
}

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-fault-device" {
        for _, c := range hardwareFaultDevices.HardwareFaultDevice {
            if hardwareFaultDevices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice{}
        hardwareFaultDevices.HardwareFaultDevice = append(hardwareFaultDevices.HardwareFaultDevice, child)
        return &hardwareFaultDevices.HardwareFaultDevice[len(hardwareFaultDevices.HardwareFaultDevice)-1]
    }
    return nil
}

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareFaultDevices.HardwareFaultDevice {
        children[hardwareFaultDevices.HardwareFaultDevice[i].GetSegmentPath()] = &hardwareFaultDevices.HardwareFaultDevice[i]
    }
    return children
}

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetYangName() string { return "hardware-fault-devices" }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) SetParent(parent types.Entity) { hardwareFaultDevices.parent = parent }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetParent() types.Entity { return hardwareFaultDevices.parent }

func (hardwareFaultDevices *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices) GetParentYangName() string { return "slot" }

// PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice
// Table of Hardware Failure Device
type PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault device list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    HwFaultDevice interface{}

    // Table of Hardware Failure Type. The type is slice of
    // PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType.
    HardwareFaultType []PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
}

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetFilter() yfilter.YFilter { return hardwareFaultDevice.YFilter }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) SetFilter(yf yfilter.YFilter) { hardwareFaultDevice.YFilter = yf }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetGoName(yname string) string {
    if yname == "hw-fault-device" { return "HwFaultDevice" }
    if yname == "hardware-fault-type" { return "HardwareFaultType" }
    return ""
}

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetSegmentPath() string {
    return "hardware-fault-device" + "[hw-fault-device='" + fmt.Sprintf("%v", hardwareFaultDevice.HwFaultDevice) + "']"
}

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-fault-type" {
        for _, c := range hardwareFaultDevice.HardwareFaultType {
            if hardwareFaultDevice.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType{}
        hardwareFaultDevice.HardwareFaultType = append(hardwareFaultDevice.HardwareFaultType, child)
        return &hardwareFaultDevice.HardwareFaultType[len(hardwareFaultDevice.HardwareFaultType)-1]
    }
    return nil
}

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareFaultDevice.HardwareFaultType {
        children[hardwareFaultDevice.HardwareFaultType[i].GetSegmentPath()] = &hardwareFaultDevice.HardwareFaultType[i]
    }
    return children
}

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-device"] = hardwareFaultDevice.HwFaultDevice
    return leafs
}

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetYangName() string { return "hardware-fault-device" }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) SetParent(parent types.Entity) { hardwareFaultDevice.parent = parent }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetParent() types.Entity { return hardwareFaultDevice.parent }

func (hardwareFaultDevice *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice) GetParentYangName() string { return "hardware-fault-devices" }

// PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType
// Table of Hardware Failure Type
type PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. hw fault type list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetFilter() yfilter.YFilter { return hardwareFaultType.YFilter }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) SetFilter(yf yfilter.YFilter) { hardwareFaultType.YFilter = yf }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetGoName(yname string) string {
    if yname == "hw-fault-type" { return "HwFaultType" }
    if yname == "condition-description" { return "ConditionDescription" }
    if yname == "condition-name" { return "ConditionName" }
    if yname == "device-key" { return "DeviceKey" }
    if yname == "device-version" { return "DeviceVersion" }
    if yname == "condition-raised-timestamp" { return "ConditionRaisedTimestamp" }
    if yname == "process-id" { return "ProcessId" }
    if yname == "device-description" { return "DeviceDescription" }
    if yname == "condition-severity" { return "ConditionSeverity" }
    return ""
}

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetSegmentPath() string {
    return "hardware-fault-type" + "[hw-fault-type='" + fmt.Sprintf("%v", hardwareFaultType.HwFaultType) + "']"
}

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hw-fault-type"] = hardwareFaultType.HwFaultType
    leafs["condition-description"] = hardwareFaultType.ConditionDescription
    leafs["condition-name"] = hardwareFaultType.ConditionName
    leafs["device-key"] = hardwareFaultType.DeviceKey
    leafs["device-version"] = hardwareFaultType.DeviceVersion
    leafs["condition-raised-timestamp"] = hardwareFaultType.ConditionRaisedTimestamp
    leafs["process-id"] = hardwareFaultType.ProcessId
    leafs["device-description"] = hardwareFaultType.DeviceDescription
    leafs["condition-severity"] = hardwareFaultType.ConditionSeverity
    return leafs
}

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetYangName() string { return "hardware-fault-type" }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) SetParent(parent types.Entity) { hardwareFaultType.parent = parent }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetParent() types.Entity { return hardwareFaultType.parent }

func (hardwareFaultType *PlatformFaultManager_Racks_Rack_Slots_Slot_HardwareFaultDevices_HardwareFaultDevice_HardwareFaultType) GetParentYangName() string { return "hardware-fault-device" }

