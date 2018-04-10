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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTrackings_ObjectTracking.
    ObjectTracking []ObjectTrackings_ObjectTracking
}

func (objectTrackings *ObjectTrackings) GetEntityData() *types.CommonEntityData {
    objectTrackings.EntityData.YFilter = objectTrackings.YFilter
    objectTrackings.EntityData.YangName = "object-trackings"
    objectTrackings.EntityData.BundleName = "cisco_ios_xr"
    objectTrackings.EntityData.ParentYangName = "Cisco-IOS-XR-manageability-object-tracking-cfg"
    objectTrackings.EntityData.SegmentPath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings"
    objectTrackings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectTrackings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectTrackings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectTrackings.EntityData.Children = make(map[string]types.YChild)
    objectTrackings.EntityData.Children["object-tracking"] = types.YChild{"ObjectTracking", nil}
    for i := range objectTrackings.ObjectTracking {
        objectTrackings.EntityData.Children[types.GetSegmentPath(&objectTrackings.ObjectTracking[i])] = types.YChild{"ObjectTracking", &objectTrackings.ObjectTracking[i]}
    }
    objectTrackings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(objectTrackings.EntityData)
}

// ObjectTrackings_ObjectTracking
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Track name. The type is string with length: 1..32.
    TrackName interface{}

    // Delay up in seconds. The type is interface{} with range: 1..180. Units are
    // second.
    DelayUp interface{}

    // Enable the Track. The type is interface{}.
    Enable interface{}

    // Track delay down time. The type is interface{} with range: 1..180. Units
    // are second.
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

    // Track type RTR (Response Time Reporter - IPSLA).
    TypeRtr ObjectTrackings_ObjectTracking_TypeRtr

    // Track type boolean list.
    TypeList ObjectTrackings_ObjectTracking_TypeList

    // Track type route ipv4.
    TypeRoute ObjectTrackings_ObjectTracking_TypeRoute

    // Track type boolean list.
    TypeBooleanList ObjectTrackings_ObjectTracking_TypeBooleanList
}

func (objectTracking *ObjectTrackings_ObjectTracking) GetEntityData() *types.CommonEntityData {
    objectTracking.EntityData.YFilter = objectTracking.YFilter
    objectTracking.EntityData.YangName = "object-tracking"
    objectTracking.EntityData.BundleName = "cisco_ios_xr"
    objectTracking.EntityData.ParentYangName = "object-trackings"
    objectTracking.EntityData.SegmentPath = "object-tracking" + "[track-name='" + fmt.Sprintf("%v", objectTracking.TrackName) + "']"
    objectTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectTracking.EntityData.Children = make(map[string]types.YChild)
    objectTracking.EntityData.Children["type-interface"] = types.YChild{"TypeInterface", &objectTracking.TypeInterface}
    objectTracking.EntityData.Children["type-rtr"] = types.YChild{"TypeRtr", &objectTracking.TypeRtr}
    objectTracking.EntityData.Children["type-list"] = types.YChild{"TypeList", &objectTracking.TypeList}
    objectTracking.EntityData.Children["type-route"] = types.YChild{"TypeRoute", &objectTracking.TypeRoute}
    objectTracking.EntityData.Children["type-boolean-list"] = types.YChild{"TypeBooleanList", &objectTracking.TypeBooleanList}
    objectTracking.EntityData.Leafs = make(map[string]types.YLeaf)
    objectTracking.EntityData.Leafs["track-name"] = types.YLeaf{"TrackName", objectTracking.TrackName}
    objectTracking.EntityData.Leafs["delay-up"] = types.YLeaf{"DelayUp", objectTracking.DelayUp}
    objectTracking.EntityData.Leafs["enable"] = types.YLeaf{"Enable", objectTracking.Enable}
    objectTracking.EntityData.Leafs["delay-down"] = types.YLeaf{"DelayDown", objectTracking.DelayDown}
    objectTracking.EntityData.Leafs["type-interface-enable"] = types.YLeaf{"TypeInterfaceEnable", objectTracking.TypeInterfaceEnable}
    objectTracking.EntityData.Leafs["type-route-enable"] = types.YLeaf{"TypeRouteEnable", objectTracking.TypeRouteEnable}
    objectTracking.EntityData.Leafs["type-boolean-list-and-enable"] = types.YLeaf{"TypeBooleanListAndEnable", objectTracking.TypeBooleanListAndEnable}
    objectTracking.EntityData.Leafs["type-boolean-list-or-enable"] = types.YLeaf{"TypeBooleanListOrEnable", objectTracking.TypeBooleanListOrEnable}
    return &(objectTracking.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeInterface
// Track type line-protocol
type ObjectTrackings_ObjectTracking_TypeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}
}

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetEntityData() *types.CommonEntityData {
    typeInterface.EntityData.YFilter = typeInterface.YFilter
    typeInterface.EntityData.YangName = "type-interface"
    typeInterface.EntityData.BundleName = "cisco_ios_xr"
    typeInterface.EntityData.ParentYangName = "object-tracking"
    typeInterface.EntityData.SegmentPath = "type-interface"
    typeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeInterface.EntityData.Children = make(map[string]types.YChild)
    typeInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    typeInterface.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", typeInterface.Interface_}
    return &(typeInterface.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeRtr
// Track type RTR (Response Time Reporter - IPSLA)
type ObjectTrackings_ObjectTracking_TypeRtr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPSLA Operation ID. The type is interface{} with range: 1..2048.
    Rtr interface{}
}

func (typeRtr *ObjectTrackings_ObjectTracking_TypeRtr) GetEntityData() *types.CommonEntityData {
    typeRtr.EntityData.YFilter = typeRtr.YFilter
    typeRtr.EntityData.YangName = "type-rtr"
    typeRtr.EntityData.BundleName = "cisco_ios_xr"
    typeRtr.EntityData.ParentYangName = "object-tracking"
    typeRtr.EntityData.SegmentPath = "type-rtr"
    typeRtr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeRtr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeRtr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeRtr.EntityData.Children = make(map[string]types.YChild)
    typeRtr.EntityData.Leafs = make(map[string]types.YLeaf)
    typeRtr.EntityData.Leafs["rtr"] = types.YLeaf{"Rtr", typeRtr.Rtr}
    return &(typeRtr.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList
// Track type boolean list
type ObjectTrackings_ObjectTracking_TypeList struct {
    EntityData types.CommonEntityData
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

func (typeList *ObjectTrackings_ObjectTracking_TypeList) GetEntityData() *types.CommonEntityData {
    typeList.EntityData.YFilter = typeList.YFilter
    typeList.EntityData.YangName = "type-list"
    typeList.EntityData.BundleName = "cisco_ios_xr"
    typeList.EntityData.ParentYangName = "object-tracking"
    typeList.EntityData.SegmentPath = "type-list"
    typeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeList.EntityData.Children = make(map[string]types.YChild)
    typeList.EntityData.Children["threshold-weight"] = types.YChild{"ThresholdWeight", &typeList.ThresholdWeight}
    typeList.EntityData.Children["threshold-percentage-object"] = types.YChild{"ThresholdPercentageObject", &typeList.ThresholdPercentageObject}
    typeList.EntityData.Children["threshold-percentage"] = types.YChild{"ThresholdPercentage", &typeList.ThresholdPercentage}
    typeList.EntityData.Children["threshold-weight-object"] = types.YChild{"ThresholdWeightObject", &typeList.ThresholdWeightObject}
    typeList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(typeList.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight
// Track type threshold weight
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold Limits.
    ThresholdLimits ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits
}

func (thresholdWeight *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight) GetEntityData() *types.CommonEntityData {
    thresholdWeight.EntityData.YFilter = thresholdWeight.YFilter
    thresholdWeight.EntityData.YangName = "threshold-weight"
    thresholdWeight.EntityData.BundleName = "cisco_ios_xr"
    thresholdWeight.EntityData.ParentYangName = "type-list"
    thresholdWeight.EntityData.SegmentPath = "threshold-weight"
    thresholdWeight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdWeight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdWeight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdWeight.EntityData.Children = make(map[string]types.YChild)
    thresholdWeight.EntityData.Children["threshold-limits"] = types.YChild{"ThresholdLimits", &thresholdWeight.ThresholdLimits}
    thresholdWeight.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdWeight.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits
// Threshold Limits
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold limit at which track is set to UP state.
    ThresholdUpValues ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits) GetEntityData() *types.CommonEntityData {
    thresholdLimits.EntityData.YFilter = thresholdLimits.YFilter
    thresholdLimits.EntityData.YangName = "threshold-limits"
    thresholdLimits.EntityData.BundleName = "cisco_ios_xr"
    thresholdLimits.EntityData.ParentYangName = "threshold-weight"
    thresholdLimits.EntityData.SegmentPath = "threshold-limits"
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = make(map[string]types.YChild)
    thresholdLimits.EntityData.Children["threshold-up-values"] = types.YChild{"ThresholdUpValues", &thresholdLimits.ThresholdUpValues}
    thresholdLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdLimits.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold limit at which track is set to UP state. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue.
    ThresholdUpValue []ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetEntityData() *types.CommonEntityData {
    thresholdUpValues.EntityData.YFilter = thresholdUpValues.YFilter
    thresholdUpValues.EntityData.YangName = "threshold-up-values"
    thresholdUpValues.EntityData.BundleName = "cisco_ios_xr"
    thresholdUpValues.EntityData.ParentYangName = "threshold-limits"
    thresholdUpValues.EntityData.SegmentPath = "threshold-up-values"
    thresholdUpValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdUpValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdUpValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdUpValues.EntityData.Children = make(map[string]types.YChild)
    thresholdUpValues.EntityData.Children["threshold-up-value"] = types.YChild{"ThresholdUpValue", nil}
    for i := range thresholdUpValues.ThresholdUpValue {
        thresholdUpValues.EntityData.Children[types.GetSegmentPath(&thresholdUpValues.ThresholdUpValue[i])] = types.YChild{"ThresholdUpValue", &thresholdUpValues.ThresholdUpValue[i]}
    }
    thresholdUpValues.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdUpValues.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Up value. The type is interface{} with range:
    // -2147483648..2147483647.
    Up interface{}

    // Threshold limit at which track is set to Down state. The type is
    // interface{} with range: -2147483648..2147483647. The default value is 0.
    ThresholdDown interface{}
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetEntityData() *types.CommonEntityData {
    thresholdUpValue.EntityData.YFilter = thresholdUpValue.YFilter
    thresholdUpValue.EntityData.YangName = "threshold-up-value"
    thresholdUpValue.EntityData.BundleName = "cisco_ios_xr"
    thresholdUpValue.EntityData.ParentYangName = "threshold-up-values"
    thresholdUpValue.EntityData.SegmentPath = "threshold-up-value" + "[up='" + fmt.Sprintf("%v", thresholdUpValue.Up) + "']"
    thresholdUpValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdUpValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdUpValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdUpValue.EntityData.Children = make(map[string]types.YChild)
    thresholdUpValue.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdUpValue.EntityData.Leafs["up"] = types.YLeaf{"Up", thresholdUpValue.Up}
    thresholdUpValue.EntityData.Leafs["threshold-down"] = types.YLeaf{"ThresholdDown", thresholdUpValue.ThresholdDown}
    return &(thresholdUpValue.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject
// Track type threshold percentage
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name object. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object.
    Object []ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object
}

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetEntityData() *types.CommonEntityData {
    thresholdPercentageObject.EntityData.YFilter = thresholdPercentageObject.YFilter
    thresholdPercentageObject.EntityData.YangName = "threshold-percentage-object"
    thresholdPercentageObject.EntityData.BundleName = "cisco_ios_xr"
    thresholdPercentageObject.EntityData.ParentYangName = "type-list"
    thresholdPercentageObject.EntityData.SegmentPath = "threshold-percentage-object"
    thresholdPercentageObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdPercentageObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdPercentageObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdPercentageObject.EntityData.Children = make(map[string]types.YChild)
    thresholdPercentageObject.EntityData.Children["object"] = types.YChild{"Object", nil}
    for i := range thresholdPercentageObject.Object {
        thresholdPercentageObject.EntityData.Children[types.GetSegmentPath(&thresholdPercentageObject.Object[i])] = types.YChild{"Object", &thresholdPercentageObject.Object[i]}
    }
    thresholdPercentageObject.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdPercentageObject.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object
// Track name object
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    Object interface{}

    // Weight of object. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 1.
    ObjectWeight interface{}
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "threshold-percentage-object"
    object.EntityData.SegmentPath = "object" + "[object='" + fmt.Sprintf("%v", object.Object) + "']"
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = make(map[string]types.YChild)
    object.EntityData.Leafs = make(map[string]types.YLeaf)
    object.EntityData.Leafs["object"] = types.YLeaf{"Object", object.Object}
    object.EntityData.Leafs["object-weight"] = types.YLeaf{"ObjectWeight", object.ObjectWeight}
    return &(object.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage
// Track type threshold percentage
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold Limits.
    ThresholdLimits ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits
}

func (thresholdPercentage *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage) GetEntityData() *types.CommonEntityData {
    thresholdPercentage.EntityData.YFilter = thresholdPercentage.YFilter
    thresholdPercentage.EntityData.YangName = "threshold-percentage"
    thresholdPercentage.EntityData.BundleName = "cisco_ios_xr"
    thresholdPercentage.EntityData.ParentYangName = "type-list"
    thresholdPercentage.EntityData.SegmentPath = "threshold-percentage"
    thresholdPercentage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdPercentage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdPercentage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdPercentage.EntityData.Children = make(map[string]types.YChild)
    thresholdPercentage.EntityData.Children["threshold-limits"] = types.YChild{"ThresholdLimits", &thresholdPercentage.ThresholdLimits}
    thresholdPercentage.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdPercentage.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits
// Threshold Limits
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold limit at which track is set to UP state.
    ThresholdUpValues ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues
}

func (thresholdLimits *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits) GetEntityData() *types.CommonEntityData {
    thresholdLimits.EntityData.YFilter = thresholdLimits.YFilter
    thresholdLimits.EntityData.YangName = "threshold-limits"
    thresholdLimits.EntityData.BundleName = "cisco_ios_xr"
    thresholdLimits.EntityData.ParentYangName = "threshold-percentage"
    thresholdLimits.EntityData.SegmentPath = "threshold-limits"
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = make(map[string]types.YChild)
    thresholdLimits.EntityData.Children["threshold-up-values"] = types.YChild{"ThresholdUpValues", &thresholdLimits.ThresholdUpValues}
    thresholdLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdLimits.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold limit at which track is set to UP state. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue.
    ThresholdUpValue []ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetEntityData() *types.CommonEntityData {
    thresholdUpValues.EntityData.YFilter = thresholdUpValues.YFilter
    thresholdUpValues.EntityData.YangName = "threshold-up-values"
    thresholdUpValues.EntityData.BundleName = "cisco_ios_xr"
    thresholdUpValues.EntityData.ParentYangName = "threshold-limits"
    thresholdUpValues.EntityData.SegmentPath = "threshold-up-values"
    thresholdUpValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdUpValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdUpValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdUpValues.EntityData.Children = make(map[string]types.YChild)
    thresholdUpValues.EntityData.Children["threshold-up-value"] = types.YChild{"ThresholdUpValue", nil}
    for i := range thresholdUpValues.ThresholdUpValue {
        thresholdUpValues.EntityData.Children[types.GetSegmentPath(&thresholdUpValues.ThresholdUpValue[i])] = types.YChild{"ThresholdUpValue", &thresholdUpValues.ThresholdUpValue[i]}
    }
    thresholdUpValues.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdUpValues.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Up value. The type is interface{} with range:
    // -2147483648..2147483647.
    Up interface{}

    // Threshold limit at which track is set to Down state. The type is
    // interface{} with range: -2147483648..2147483647. The default value is 0.
    ThresholdDown interface{}
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetEntityData() *types.CommonEntityData {
    thresholdUpValue.EntityData.YFilter = thresholdUpValue.YFilter
    thresholdUpValue.EntityData.YangName = "threshold-up-value"
    thresholdUpValue.EntityData.BundleName = "cisco_ios_xr"
    thresholdUpValue.EntityData.ParentYangName = "threshold-up-values"
    thresholdUpValue.EntityData.SegmentPath = "threshold-up-value" + "[up='" + fmt.Sprintf("%v", thresholdUpValue.Up) + "']"
    thresholdUpValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdUpValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdUpValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdUpValue.EntityData.Children = make(map[string]types.YChild)
    thresholdUpValue.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdUpValue.EntityData.Leafs["up"] = types.YLeaf{"Up", thresholdUpValue.Up}
    thresholdUpValue.EntityData.Leafs["threshold-down"] = types.YLeaf{"ThresholdDown", thresholdUpValue.ThresholdDown}
    return &(thresholdUpValue.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject
// Track type threshold weight
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name object. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object.
    Object []ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object
}

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetEntityData() *types.CommonEntityData {
    thresholdWeightObject.EntityData.YFilter = thresholdWeightObject.YFilter
    thresholdWeightObject.EntityData.YangName = "threshold-weight-object"
    thresholdWeightObject.EntityData.BundleName = "cisco_ios_xr"
    thresholdWeightObject.EntityData.ParentYangName = "type-list"
    thresholdWeightObject.EntityData.SegmentPath = "threshold-weight-object"
    thresholdWeightObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdWeightObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdWeightObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdWeightObject.EntityData.Children = make(map[string]types.YChild)
    thresholdWeightObject.EntityData.Children["object"] = types.YChild{"Object", nil}
    for i := range thresholdWeightObject.Object {
        thresholdWeightObject.EntityData.Children[types.GetSegmentPath(&thresholdWeightObject.Object[i])] = types.YChild{"Object", &thresholdWeightObject.Object[i]}
    }
    thresholdWeightObject.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdWeightObject.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object
// Track name object
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    Object interface{}

    // Weight of object. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 1.
    ObjectWeight interface{}
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "threshold-weight-object"
    object.EntityData.SegmentPath = "object" + "[object='" + fmt.Sprintf("%v", object.Object) + "']"
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = make(map[string]types.YChild)
    object.EntityData.Leafs = make(map[string]types.YLeaf)
    object.EntityData.Leafs["object"] = types.YLeaf{"Object", object.Object}
    object.EntityData.Leafs["object-weight"] = types.YLeaf{"ObjectWeight", object.ObjectWeight}
    return &(object.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeRoute
// Track type route ipv4
type ObjectTrackings_ObjectTracking_TypeRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF tag - use 'default' for the DEFAULT vrf. The type is string with
    // length: 1..32.
    Vrf interface{}

    // set track IPv4 address.
    IpAddress ObjectTrackings_ObjectTracking_TypeRoute_IpAddress
}

func (typeRoute *ObjectTrackings_ObjectTracking_TypeRoute) GetEntityData() *types.CommonEntityData {
    typeRoute.EntityData.YFilter = typeRoute.YFilter
    typeRoute.EntityData.YangName = "type-route"
    typeRoute.EntityData.BundleName = "cisco_ios_xr"
    typeRoute.EntityData.ParentYangName = "object-tracking"
    typeRoute.EntityData.SegmentPath = "type-route"
    typeRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeRoute.EntityData.Children = make(map[string]types.YChild)
    typeRoute.EntityData.Children["ip-address"] = types.YChild{"IpAddress", &typeRoute.IpAddress}
    typeRoute.EntityData.Leafs = make(map[string]types.YLeaf)
    typeRoute.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", typeRoute.Vrf}
    return &(typeRoute.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeRoute_IpAddress
// set track IPv4 address
// This type is a presence type.
type ObjectTrackings_ObjectTracking_TypeRoute_IpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Mask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Mask interface{}
}

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetEntityData() *types.CommonEntityData {
    ipAddress.EntityData.YFilter = ipAddress.YFilter
    ipAddress.EntityData.YangName = "ip-address"
    ipAddress.EntityData.BundleName = "cisco_ios_xr"
    ipAddress.EntityData.ParentYangName = "type-route"
    ipAddress.EntityData.SegmentPath = "ip-address"
    ipAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipAddress.EntityData.Children = make(map[string]types.YChild)
    ipAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    ipAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", ipAddress.Address}
    ipAddress.EntityData.Leafs["mask"] = types.YLeaf{"Mask", ipAddress.Mask}
    return &(ipAddress.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList
// Track type boolean list
type ObjectTrackings_ObjectTracking_TypeBooleanList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track type boolean or list.
    OrObjects ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects

    // Track type boolean and list.
    AndObjects ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects
}

func (typeBooleanList *ObjectTrackings_ObjectTracking_TypeBooleanList) GetEntityData() *types.CommonEntityData {
    typeBooleanList.EntityData.YFilter = typeBooleanList.YFilter
    typeBooleanList.EntityData.YangName = "type-boolean-list"
    typeBooleanList.EntityData.BundleName = "cisco_ios_xr"
    typeBooleanList.EntityData.ParentYangName = "object-tracking"
    typeBooleanList.EntityData.SegmentPath = "type-boolean-list"
    typeBooleanList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeBooleanList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeBooleanList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeBooleanList.EntityData.Children = make(map[string]types.YChild)
    typeBooleanList.EntityData.Children["or-objects"] = types.YChild{"OrObjects", &typeBooleanList.OrObjects}
    typeBooleanList.EntityData.Children["and-objects"] = types.YChild{"AndObjects", &typeBooleanList.AndObjects}
    typeBooleanList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(typeBooleanList.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects
// Track type boolean or list
type ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject.
    OrObject []ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject
}

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetEntityData() *types.CommonEntityData {
    orObjects.EntityData.YFilter = orObjects.YFilter
    orObjects.EntityData.YangName = "or-objects"
    orObjects.EntityData.BundleName = "cisco_ios_xr"
    orObjects.EntityData.ParentYangName = "type-boolean-list"
    orObjects.EntityData.SegmentPath = "or-objects"
    orObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orObjects.EntityData.Children = make(map[string]types.YChild)
    orObjects.EntityData.Children["or-object"] = types.YChild{"OrObject", nil}
    for i := range orObjects.OrObject {
        orObjects.EntityData.Children[types.GetSegmentPath(&orObjects.OrObject[i])] = types.YChild{"OrObject", &orObjects.OrObject[i]}
    }
    orObjects.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(orObjects.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    Object interface{}

    // Tracked Object sign (with or without not). The type is
    // ObjectTrackingBooleanSign.
    ObjectSign interface{}
}

func (orObject *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject) GetEntityData() *types.CommonEntityData {
    orObject.EntityData.YFilter = orObject.YFilter
    orObject.EntityData.YangName = "or-object"
    orObject.EntityData.BundleName = "cisco_ios_xr"
    orObject.EntityData.ParentYangName = "or-objects"
    orObject.EntityData.SegmentPath = "or-object" + "[object='" + fmt.Sprintf("%v", orObject.Object) + "']"
    orObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orObject.EntityData.Children = make(map[string]types.YChild)
    orObject.EntityData.Leafs = make(map[string]types.YLeaf)
    orObject.EntityData.Leafs["object"] = types.YLeaf{"Object", orObject.Object}
    orObject.EntityData.Leafs["object-sign"] = types.YLeaf{"ObjectSign", orObject.ObjectSign}
    return &(orObject.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects
// Track type boolean and list
type ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject.
    AndObject []ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject
}

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetEntityData() *types.CommonEntityData {
    andObjects.EntityData.YFilter = andObjects.YFilter
    andObjects.EntityData.YangName = "and-objects"
    andObjects.EntityData.BundleName = "cisco_ios_xr"
    andObjects.EntityData.ParentYangName = "type-boolean-list"
    andObjects.EntityData.SegmentPath = "and-objects"
    andObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    andObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    andObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    andObjects.EntityData.Children = make(map[string]types.YChild)
    andObjects.EntityData.Children["and-object"] = types.YChild{"AndObject", nil}
    for i := range andObjects.AndObject {
        andObjects.EntityData.Children[types.GetSegmentPath(&andObjects.AndObject[i])] = types.YChild{"AndObject", &andObjects.AndObject[i]}
    }
    andObjects.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(andObjects.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    ObjectName interface{}

    // Tracked Object sign (with or without not). The type is
    // ObjectTrackingBooleanSign.
    ObjectSign interface{}
}

func (andObject *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject) GetEntityData() *types.CommonEntityData {
    andObject.EntityData.YFilter = andObject.YFilter
    andObject.EntityData.YangName = "and-object"
    andObject.EntityData.BundleName = "cisco_ios_xr"
    andObject.EntityData.ParentYangName = "and-objects"
    andObject.EntityData.SegmentPath = "and-object" + "[object-name='" + fmt.Sprintf("%v", andObject.ObjectName) + "']"
    andObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    andObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    andObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    andObject.EntityData.Children = make(map[string]types.YChild)
    andObject.EntityData.Leafs = make(map[string]types.YLeaf)
    andObject.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", andObject.ObjectName}
    andObject.EntityData.Leafs["object-sign"] = types.YLeaf{"ObjectSign", andObject.ObjectSign}
    return &(andObject.EntityData)
}

