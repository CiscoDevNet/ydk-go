// This module contains a collection of YANG definitions
// for Cisco IOS-XR manageability-object-tracking package configuration.
// 
// This module contains definitions
// for the following management objects:
//   object-trackings: Object Tracking configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package manageability_object_tracking_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package manageability_object_tracking_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-manageability-object-tracking-cfg object-trackings}", reflect.TypeOf(ObjectTrackings{}))
    ydk.RegisterEntity("Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings", reflect.TypeOf(ObjectTrackings{}))
}

// ObjectTrackings
// Object Tracking configuration
type ObjectTrackings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTrackings_ObjectTracking.
    ObjectTracking []ObjectTrackings_ObjectTracking
}

func (objectTrackings *ObjectTrackings) GetFilter() yfilter.YFilter { return objectTrackings.YFilter }

func (objectTrackings *ObjectTrackings) SetFilter(yf yfilter.YFilter) { objectTrackings.YFilter = yf }

func (objectTrackings *ObjectTrackings) GetGoName(yname string) string {
    if yname == "object-tracking" { return "ObjectTracking" }
    return ""
}

func (objectTrackings *ObjectTrackings) GetSegmentPath() string {
    return "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings"
}

func (objectTrackings *ObjectTrackings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "object-tracking" {
        for _, c := range objectTrackings.ObjectTracking {
            if objectTrackings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTrackings_ObjectTracking{}
        objectTrackings.ObjectTracking = append(objectTrackings.ObjectTracking, child)
        return &objectTrackings.ObjectTracking[len(objectTrackings.ObjectTracking)-1]
    }
    return nil
}

func (objectTrackings *ObjectTrackings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range objectTrackings.ObjectTracking {
        children[objectTrackings.ObjectTracking[i].GetSegmentPath()] = &objectTrackings.ObjectTracking[i]
    }
    return children
}

func (objectTrackings *ObjectTrackings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (objectTrackings *ObjectTrackings) GetBundleName() string { return "cisco_ios_xr" }

func (objectTrackings *ObjectTrackings) GetYangName() string { return "object-trackings" }

func (objectTrackings *ObjectTrackings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (objectTrackings *ObjectTrackings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (objectTrackings *ObjectTrackings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (objectTrackings *ObjectTrackings) SetParent(parent types.Entity) { objectTrackings.parent = parent }

func (objectTrackings *ObjectTrackings) GetParent() types.Entity { return objectTrackings.parent }

func (objectTrackings *ObjectTrackings) GetParentYangName() string { return "Cisco-IOS-XR-manageability-object-tracking-cfg" }

// ObjectTrackings_ObjectTracking
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Track name. The type is string with length: 1..32.
    TrackName interface{}

    // Delay up in seconds. The type is interface{} with range: 1..10. Units are
    // second.
    DelayUp interface{}

    // Enable the Track. The type is interface{}.
    Enable interface{}

    // Track delay down time. The type is interface{} with range: 1..10. Units are
    // second.
    DelayDown interface{}

    // Enable track type Interface. The type is interface{}.
    TypeInterfaceEnable interface{}

    // Enable track type Route. The type is interface{}.
    TypeRouteEnable interface{}

    // Enable track type boolean list and. The type is interface{}.
    TypeBooleanListAndEnable interface{}

    // Enable track type boolean list or. The type is interface{}.
    TypeBooleanListOrEnable interface{}

    // Track type line-protocol.
    TypeInterface ObjectTrackings_ObjectTracking_TypeInterface

    // Track type boolean list.
    TypeList ObjectTrackings_ObjectTracking_TypeList

    // Track type route ipv4.
    TypeRoute ObjectTrackings_ObjectTracking_TypeRoute

    // Track type boolean list.
    TypeBooleanList ObjectTrackings_ObjectTracking_TypeBooleanList
}

func (objectTracking *ObjectTrackings_ObjectTracking) GetFilter() yfilter.YFilter { return objectTracking.YFilter }

func (objectTracking *ObjectTrackings_ObjectTracking) SetFilter(yf yfilter.YFilter) { objectTracking.YFilter = yf }

func (objectTracking *ObjectTrackings_ObjectTracking) GetGoName(yname string) string {
    if yname == "track-name" { return "TrackName" }
    if yname == "delay-up" { return "DelayUp" }
    if yname == "enable" { return "Enable" }
    if yname == "delay-down" { return "DelayDown" }
    if yname == "type-interface-enable" { return "TypeInterfaceEnable" }
    if yname == "type-route-enable" { return "TypeRouteEnable" }
    if yname == "type-boolean-list-and-enable" { return "TypeBooleanListAndEnable" }
    if yname == "type-boolean-list-or-enable" { return "TypeBooleanListOrEnable" }
    if yname == "type-interface" { return "TypeInterface" }
    if yname == "type-list" { return "TypeList" }
    if yname == "type-route" { return "TypeRoute" }
    if yname == "type-boolean-list" { return "TypeBooleanList" }
    return ""
}

func (objectTracking *ObjectTrackings_ObjectTracking) GetSegmentPath() string {
    return "object-tracking" + "[track-name='" + fmt.Sprintf("%v", objectTracking.TrackName) + "']"
}

func (objectTracking *ObjectTrackings_ObjectTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "type-interface" {
        return &objectTracking.TypeInterface
    }
    if childYangName == "type-list" {
        return &objectTracking.TypeList
    }
    if childYangName == "type-route" {
        return &objectTracking.TypeRoute
    }
    if childYangName == "type-boolean-list" {
        return &objectTracking.TypeBooleanList
    }
    return nil
}

func (objectTracking *ObjectTrackings_ObjectTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["type-interface"] = &objectTracking.TypeInterface
    children["type-list"] = &objectTracking.TypeList
    children["type-route"] = &objectTracking.TypeRoute
    children["type-boolean-list"] = &objectTracking.TypeBooleanList
    return children
}

func (objectTracking *ObjectTrackings_ObjectTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-name"] = objectTracking.TrackName
    leafs["delay-up"] = objectTracking.DelayUp
    leafs["enable"] = objectTracking.Enable
    leafs["delay-down"] = objectTracking.DelayDown
    leafs["type-interface-enable"] = objectTracking.TypeInterfaceEnable
    leafs["type-route-enable"] = objectTracking.TypeRouteEnable
    leafs["type-boolean-list-and-enable"] = objectTracking.TypeBooleanListAndEnable
    leafs["type-boolean-list-or-enable"] = objectTracking.TypeBooleanListOrEnable
    return leafs
}

func (objectTracking *ObjectTrackings_ObjectTracking) GetBundleName() string { return "cisco_ios_xr" }

func (objectTracking *ObjectTrackings_ObjectTracking) GetYangName() string { return "object-tracking" }

func (objectTracking *ObjectTrackings_ObjectTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (objectTracking *ObjectTrackings_ObjectTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (objectTracking *ObjectTrackings_ObjectTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (objectTracking *ObjectTrackings_ObjectTracking) SetParent(parent types.Entity) { objectTracking.parent = parent }

func (objectTracking *ObjectTrackings_ObjectTracking) GetParent() types.Entity { return objectTracking.parent }

func (objectTracking *ObjectTrackings_ObjectTracking) GetParentYangName() string { return "object-trackings" }

// ObjectTrackings_ObjectTracking_TypeInterface
// Track type line-protocol
type ObjectTrackings_ObjectTracking_TypeInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}
}

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetFilter() yfilter.YFilter { return typeInterface.YFilter }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) SetFilter(yf yfilter.YFilter) { typeInterface.YFilter = yf }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetSegmentPath() string {
    return "type-interface"
}

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = typeInterface.Interface
    return leafs
}

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetBundleName() string { return "cisco_ios_xr" }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetYangName() string { return "type-interface" }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) SetParent(parent types.Entity) { typeInterface.parent = parent }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetParent() types.Entity { return typeInterface.parent }

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetParentYangName() string { return "object-tracking" }

// ObjectTrackings_ObjectTracking_TypeList
// Track type boolean list
type ObjectTrackings_ObjectTracking_TypeList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track type threshold weight.
    ThresholdWeight ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight

    // Track type threshold percentage.
    ThresholdPercentageObject ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject

    // Track type threshold percentage.
    ThresholdPercentage ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage

    // Track type threshold weight.
    ThresholdWeightObject ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject
}

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetFilter() yfilter.YFilter { return typeList.YFilter }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) SetFilter(yf yfilter.YFilter) { typeList.YFilter = yf }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetGoName(yname string) string {
    if yname == "threshold-weight" { return "ThresholdWeight" }
    if yname == "threshold-percentage-object" { return "ThresholdPercentageObject" }
    if yname == "threshold-percentage" { return "ThresholdPercentage" }
    if yname == "threshold-weight-object" { return "ThresholdWeightObject" }
    return ""
}

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetSegmentPath() string {
    return "type-list"
}

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-weight" {
        return &typeList.ThresholdWeight
    }
    if childYangName == "threshold-percentage-object" {
        return &typeList.ThresholdPercentageObject
    }
    if childYangName == "threshold-percentage" {
        return &typeList.ThresholdPercentage
    }
    if childYangName == "threshold-weight-object" {
        return &typeList.ThresholdWeightObject
    }
    return nil
}

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold-weight"] = &typeList.ThresholdWeight
    children["threshold-percentage-object"] = &typeList.ThresholdPercentageObject
    children["threshold-percentage"] = &typeList.ThresholdPercentage
    children["threshold-weight-object"] = &typeList.ThresholdWeightObject
    return children
}

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetBundleName() string { return "cisco_ios_xr" }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetYangName() string { return "type-list" }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) SetParent(parent types.Entity) { typeList.parent = parent }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetParent() types.Entity { return typeList.parent }

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetParentYangName() string { return "object-tracking" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight
// Track type threshold weight
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold Limits.
    ThresholdLimits ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits
}

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetFilter() yfilter.YFilter { return thresholdWeight.YFilter }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) SetFilter(yf yfilter.YFilter) { thresholdWeight.YFilter = yf }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetGoName(yname string) string {
    if yname == "threshold-limits" { return "ThresholdLimits" }
    return ""
}

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetSegmentPath() string {
    return "threshold-weight"
}

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-limits" {
        return &thresholdWeight.ThresholdLimits
    }
    return nil
}

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold-limits"] = &thresholdWeight.ThresholdLimits
    return children
}

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetYangName() string { return "threshold-weight" }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) SetParent(parent types.Entity) { thresholdWeight.parent = parent }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetParent() types.Entity { return thresholdWeight.parent }

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetParentYangName() string { return "type-list" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits
// Threshold Limits
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold limit at which track is set to UP state.
    ThresholdUpValues ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetFilter() yfilter.YFilter { return thresholdLimits.YFilter }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) SetFilter(yf yfilter.YFilter) { thresholdLimits.YFilter = yf }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetGoName(yname string) string {
    if yname == "threshold-up-values" { return "ThresholdUpValues" }
    return ""
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetSegmentPath() string {
    return "threshold-limits"
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-up-values" {
        return &thresholdLimits.ThresholdUpValues
    }
    return nil
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold-up-values"] = &thresholdLimits.ThresholdUpValues
    return children
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetYangName() string { return "threshold-limits" }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) SetParent(parent types.Entity) { thresholdLimits.parent = parent }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetParent() types.Entity { return thresholdLimits.parent }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetParentYangName() string { return "threshold-weight" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold limit at which track is set to UP state. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue.
    ThresholdUpValue []ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetFilter() yfilter.YFilter { return thresholdUpValues.YFilter }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) SetFilter(yf yfilter.YFilter) { thresholdUpValues.YFilter = yf }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetGoName(yname string) string {
    if yname == "threshold-up-value" { return "ThresholdUpValue" }
    return ""
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetSegmentPath() string {
    return "threshold-up-values"
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-up-value" {
        for _, c := range thresholdUpValues.ThresholdUpValue {
            if thresholdUpValues.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue{}
        thresholdUpValues.ThresholdUpValue = append(thresholdUpValues.ThresholdUpValue, child)
        return &thresholdUpValues.ThresholdUpValue[len(thresholdUpValues.ThresholdUpValue)-1]
    }
    return nil
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdUpValues.ThresholdUpValue {
        children[thresholdUpValues.ThresholdUpValue[i].GetSegmentPath()] = &thresholdUpValues.ThresholdUpValue[i]
    }
    return children
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetYangName() string { return "threshold-up-values" }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) SetParent(parent types.Entity) { thresholdUpValues.parent = parent }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetParent() types.Entity { return thresholdUpValues.parent }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetParentYangName() string { return "threshold-limits" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Up value. The type is interface{} with range:
    // -2147483648..2147483647.
    Up interface{}

    // Threshold limit at which track is set to Down state. The type is
    // interface{} with range: -2147483648..2147483647. The default value is 0.
    ThresholdDown interface{}
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetFilter() yfilter.YFilter { return thresholdUpValue.YFilter }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) SetFilter(yf yfilter.YFilter) { thresholdUpValue.YFilter = yf }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetGoName(yname string) string {
    if yname == "up" { return "Up" }
    if yname == "threshold-down" { return "ThresholdDown" }
    return ""
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetSegmentPath() string {
    return "threshold-up-value" + "[up='" + fmt.Sprintf("%v", thresholdUpValue.Up) + "']"
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["up"] = thresholdUpValue.Up
    leafs["threshold-down"] = thresholdUpValue.ThresholdDown
    return leafs
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetYangName() string { return "threshold-up-value" }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) SetParent(parent types.Entity) { thresholdUpValue.parent = parent }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetParent() types.Entity { return thresholdUpValue.parent }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetParentYangName() string { return "threshold-up-values" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject
// Track type threshold percentage
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track name object. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object.
    Object []ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object
}

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetFilter() yfilter.YFilter { return thresholdPercentageObject.YFilter }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) SetFilter(yf yfilter.YFilter) { thresholdPercentageObject.YFilter = yf }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetSegmentPath() string {
    return "threshold-percentage-object"
}

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "object" {
        for _, c := range thresholdPercentageObject.Object {
            if thresholdPercentageObject.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object{}
        thresholdPercentageObject.Object = append(thresholdPercentageObject.Object, child)
        return &thresholdPercentageObject.Object[len(thresholdPercentageObject.Object)-1]
    }
    return nil
}

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdPercentageObject.Object {
        children[thresholdPercentageObject.Object[i].GetSegmentPath()] = &thresholdPercentageObject.Object[i]
    }
    return children
}

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetYangName() string { return "threshold-percentage-object" }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) SetParent(parent types.Entity) { thresholdPercentageObject.parent = parent }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetParent() types.Entity { return thresholdPercentageObject.parent }

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetParentYangName() string { return "type-list" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object
// Track name object
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    Object interface{}

    // Weight of object. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 1.
    ObjectWeight interface{}
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetFilter() yfilter.YFilter { return object.YFilter }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) SetFilter(yf yfilter.YFilter) { object.YFilter = yf }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    if yname == "object-weight" { return "ObjectWeight" }
    return ""
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetSegmentPath() string {
    return "object" + "[object='" + fmt.Sprintf("%v", object.Object) + "']"
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = object.Object
    leafs["object-weight"] = object.ObjectWeight
    return leafs
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetBundleName() string { return "cisco_ios_xr" }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetYangName() string { return "object" }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) SetParent(parent types.Entity) { object.parent = parent }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetParent() types.Entity { return object.parent }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetParentYangName() string { return "threshold-percentage-object" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage
// Track type threshold percentage
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold Limits.
    ThresholdLimits ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits
}

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetFilter() yfilter.YFilter { return thresholdPercentage.YFilter }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) SetFilter(yf yfilter.YFilter) { thresholdPercentage.YFilter = yf }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetGoName(yname string) string {
    if yname == "threshold-limits" { return "ThresholdLimits" }
    return ""
}

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetSegmentPath() string {
    return "threshold-percentage"
}

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-limits" {
        return &thresholdPercentage.ThresholdLimits
    }
    return nil
}

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold-limits"] = &thresholdPercentage.ThresholdLimits
    return children
}

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetYangName() string { return "threshold-percentage" }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) SetParent(parent types.Entity) { thresholdPercentage.parent = parent }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetParent() types.Entity { return thresholdPercentage.parent }

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetParentYangName() string { return "type-list" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits
// Threshold Limits
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold limit at which track is set to UP state.
    ThresholdUpValues ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetFilter() yfilter.YFilter { return thresholdLimits.YFilter }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) SetFilter(yf yfilter.YFilter) { thresholdLimits.YFilter = yf }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetGoName(yname string) string {
    if yname == "threshold-up-values" { return "ThresholdUpValues" }
    return ""
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetSegmentPath() string {
    return "threshold-limits"
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-up-values" {
        return &thresholdLimits.ThresholdUpValues
    }
    return nil
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold-up-values"] = &thresholdLimits.ThresholdUpValues
    return children
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetYangName() string { return "threshold-limits" }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) SetParent(parent types.Entity) { thresholdLimits.parent = parent }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetParent() types.Entity { return thresholdLimits.parent }

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetParentYangName() string { return "threshold-percentage" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold limit at which track is set to UP state. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue.
    ThresholdUpValue []ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetFilter() yfilter.YFilter { return thresholdUpValues.YFilter }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) SetFilter(yf yfilter.YFilter) { thresholdUpValues.YFilter = yf }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetGoName(yname string) string {
    if yname == "threshold-up-value" { return "ThresholdUpValue" }
    return ""
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetSegmentPath() string {
    return "threshold-up-values"
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-up-value" {
        for _, c := range thresholdUpValues.ThresholdUpValue {
            if thresholdUpValues.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue{}
        thresholdUpValues.ThresholdUpValue = append(thresholdUpValues.ThresholdUpValue, child)
        return &thresholdUpValues.ThresholdUpValue[len(thresholdUpValues.ThresholdUpValue)-1]
    }
    return nil
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdUpValues.ThresholdUpValue {
        children[thresholdUpValues.ThresholdUpValue[i].GetSegmentPath()] = &thresholdUpValues.ThresholdUpValue[i]
    }
    return children
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetYangName() string { return "threshold-up-values" }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) SetParent(parent types.Entity) { thresholdUpValues.parent = parent }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetParent() types.Entity { return thresholdUpValues.parent }

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetParentYangName() string { return "threshold-limits" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Up value. The type is interface{} with range:
    // -2147483648..2147483647.
    Up interface{}

    // Threshold limit at which track is set to Down state. The type is
    // interface{} with range: -2147483648..2147483647. The default value is 0.
    ThresholdDown interface{}
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetFilter() yfilter.YFilter { return thresholdUpValue.YFilter }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) SetFilter(yf yfilter.YFilter) { thresholdUpValue.YFilter = yf }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetGoName(yname string) string {
    if yname == "up" { return "Up" }
    if yname == "threshold-down" { return "ThresholdDown" }
    return ""
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetSegmentPath() string {
    return "threshold-up-value" + "[up='" + fmt.Sprintf("%v", thresholdUpValue.Up) + "']"
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["up"] = thresholdUpValue.Up
    leafs["threshold-down"] = thresholdUpValue.ThresholdDown
    return leafs
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetYangName() string { return "threshold-up-value" }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) SetParent(parent types.Entity) { thresholdUpValue.parent = parent }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetParent() types.Entity { return thresholdUpValue.parent }

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetParentYangName() string { return "threshold-up-values" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject
// Track type threshold weight
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track name object. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object.
    Object []ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object
}

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetFilter() yfilter.YFilter { return thresholdWeightObject.YFilter }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) SetFilter(yf yfilter.YFilter) { thresholdWeightObject.YFilter = yf }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetSegmentPath() string {
    return "threshold-weight-object"
}

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "object" {
        for _, c := range thresholdWeightObject.Object {
            if thresholdWeightObject.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object{}
        thresholdWeightObject.Object = append(thresholdWeightObject.Object, child)
        return &thresholdWeightObject.Object[len(thresholdWeightObject.Object)-1]
    }
    return nil
}

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdWeightObject.Object {
        children[thresholdWeightObject.Object[i].GetSegmentPath()] = &thresholdWeightObject.Object[i]
    }
    return children
}

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetYangName() string { return "threshold-weight-object" }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) SetParent(parent types.Entity) { thresholdWeightObject.parent = parent }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetParent() types.Entity { return thresholdWeightObject.parent }

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetParentYangName() string { return "type-list" }

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object
// Track name object
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    Object interface{}

    // Weight of object. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 1.
    ObjectWeight interface{}
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetFilter() yfilter.YFilter { return object.YFilter }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) SetFilter(yf yfilter.YFilter) { object.YFilter = yf }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    if yname == "object-weight" { return "ObjectWeight" }
    return ""
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetSegmentPath() string {
    return "object" + "[object='" + fmt.Sprintf("%v", object.Object) + "']"
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = object.Object
    leafs["object-weight"] = object.ObjectWeight
    return leafs
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetBundleName() string { return "cisco_ios_xr" }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetYangName() string { return "object" }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) SetParent(parent types.Entity) { object.parent = parent }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetParent() types.Entity { return object.parent }

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetParentYangName() string { return "threshold-weight-object" }

// ObjectTrackings_ObjectTracking_TypeRoute
// Track type route ipv4
type ObjectTrackings_ObjectTracking_TypeRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF tag - use 'default' for the DEFAULT vrf. The type is string with
    // length: 1..32.
    Vrf interface{}

    // set track IPv4 address.
    IpAddress ObjectTrackings_ObjectTracking_TypeRoute_IpAddress
}

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetFilter() yfilter.YFilter { return typeRoute.YFilter }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) SetFilter(yf yfilter.YFilter) { typeRoute.YFilter = yf }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetSegmentPath() string {
    return "type-route"
}

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ip-address" {
        return &typeRoute.IpAddress
    }
    return nil
}

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ip-address"] = &typeRoute.IpAddress
    return children
}

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf"] = typeRoute.Vrf
    return leafs
}

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetBundleName() string { return "cisco_ios_xr" }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetYangName() string { return "type-route" }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) SetParent(parent types.Entity) { typeRoute.parent = parent }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetParent() types.Entity { return typeRoute.parent }

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetParentYangName() string { return "object-tracking" }

// ObjectTrackings_ObjectTracking_TypeRoute_IpAddress
// set track IPv4 address
// This type is a presence type.
type ObjectTrackings_ObjectTracking_TypeRoute_IpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Mask interface{}
}

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetFilter() yfilter.YFilter { return ipAddress.YFilter }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) SetFilter(yf yfilter.YFilter) { ipAddress.YFilter = yf }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "mask" { return "Mask" }
    return ""
}

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetSegmentPath() string {
    return "ip-address"
}

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipAddress.Address
    leafs["mask"] = ipAddress.Mask
    return leafs
}

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetYangName() string { return "ip-address" }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) SetParent(parent types.Entity) { ipAddress.parent = parent }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetParent() types.Entity { return ipAddress.parent }

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetParentYangName() string { return "type-route" }

// ObjectTrackings_ObjectTracking_TypeBooleanList
// Track type boolean list
type ObjectTrackings_ObjectTracking_TypeBooleanList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track type boolean or list.
    OrObjects ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects

    // Track type boolean and list.
    AndObjects ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects
}

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetFilter() yfilter.YFilter { return typeBooleanList.YFilter }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) SetFilter(yf yfilter.YFilter) { typeBooleanList.YFilter = yf }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetGoName(yname string) string {
    if yname == "or-objects" { return "OrObjects" }
    if yname == "and-objects" { return "AndObjects" }
    return ""
}

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetSegmentPath() string {
    return "type-boolean-list"
}

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "or-objects" {
        return &typeBooleanList.OrObjects
    }
    if childYangName == "and-objects" {
        return &typeBooleanList.AndObjects
    }
    return nil
}

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["or-objects"] = &typeBooleanList.OrObjects
    children["and-objects"] = &typeBooleanList.AndObjects
    return children
}

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetBundleName() string { return "cisco_ios_xr" }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetYangName() string { return "type-boolean-list" }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) SetParent(parent types.Entity) { typeBooleanList.parent = parent }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetParent() types.Entity { return typeBooleanList.parent }

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetParentYangName() string { return "object-tracking" }

// ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects
// Track type boolean or list
type ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject.
    OrObject []ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject
}

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetFilter() yfilter.YFilter { return orObjects.YFilter }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) SetFilter(yf yfilter.YFilter) { orObjects.YFilter = yf }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetGoName(yname string) string {
    if yname == "or-object" { return "OrObject" }
    return ""
}

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetSegmentPath() string {
    return "or-objects"
}

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "or-object" {
        for _, c := range orObjects.OrObject {
            if orObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject{}
        orObjects.OrObject = append(orObjects.OrObject, child)
        return &orObjects.OrObject[len(orObjects.OrObject)-1]
    }
    return nil
}

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range orObjects.OrObject {
        children[orObjects.OrObject[i].GetSegmentPath()] = &orObjects.OrObject[i]
    }
    return children
}

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetBundleName() string { return "cisco_ios_xr" }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetYangName() string { return "or-objects" }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) SetParent(parent types.Entity) { orObjects.parent = parent }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetParent() types.Entity { return orObjects.parent }

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetParentYangName() string { return "type-boolean-list" }

// ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    Object interface{}

    // Tracked Object sign (with or without not). The type is
    // ObjectTrackingBooleanSign.
    ObjectSign interface{}
}

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetFilter() yfilter.YFilter { return orObject.YFilter }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) SetFilter(yf yfilter.YFilter) { orObject.YFilter = yf }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    if yname == "object-sign" { return "ObjectSign" }
    return ""
}

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetSegmentPath() string {
    return "or-object" + "[object='" + fmt.Sprintf("%v", orObject.Object) + "']"
}

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object"] = orObject.Object
    leafs["object-sign"] = orObject.ObjectSign
    return leafs
}

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetBundleName() string { return "cisco_ios_xr" }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetYangName() string { return "or-object" }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) SetParent(parent types.Entity) { orObject.parent = parent }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetParent() types.Entity { return orObject.parent }

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetParentYangName() string { return "or-objects" }

// ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects
// Track type boolean and list
type ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject.
    AndObject []ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject
}

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetFilter() yfilter.YFilter { return andObjects.YFilter }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) SetFilter(yf yfilter.YFilter) { andObjects.YFilter = yf }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetGoName(yname string) string {
    if yname == "and-object" { return "AndObject" }
    return ""
}

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetSegmentPath() string {
    return "and-objects"
}

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "and-object" {
        for _, c := range andObjects.AndObject {
            if andObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject{}
        andObjects.AndObject = append(andObjects.AndObject, child)
        return &andObjects.AndObject[len(andObjects.AndObject)-1]
    }
    return nil
}

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range andObjects.AndObject {
        children[andObjects.AndObject[i].GetSegmentPath()] = &andObjects.AndObject[i]
    }
    return children
}

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetBundleName() string { return "cisco_ios_xr" }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetYangName() string { return "and-objects" }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) SetParent(parent types.Entity) { andObjects.parent = parent }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetParent() types.Entity { return andObjects.parent }

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetParentYangName() string { return "type-boolean-list" }

// ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    ObjectName interface{}

    // Tracked Object sign (with or without not). The type is
    // ObjectTrackingBooleanSign.
    ObjectSign interface{}
}

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetFilter() yfilter.YFilter { return andObject.YFilter }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) SetFilter(yf yfilter.YFilter) { andObject.YFilter = yf }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "object-sign" { return "ObjectSign" }
    return ""
}

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetSegmentPath() string {
    return "and-object" + "[object-name='" + fmt.Sprintf("%v", andObject.ObjectName) + "']"
}

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = andObject.ObjectName
    leafs["object-sign"] = andObject.ObjectSign
    return leafs
}

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetBundleName() string { return "cisco_ios_xr" }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetYangName() string { return "and-object" }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) SetParent(parent types.Entity) { andObject.parent = parent }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetParent() types.Entity { return andObject.parent }

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetParentYangName() string { return "and-objects" }

