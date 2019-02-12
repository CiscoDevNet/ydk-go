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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    ObjectTracking []*ObjectTrackings_ObjectTracking
}

func (objectTrackings *ObjectTrackings) GetEntityData() *types.CommonEntityData {
    objectTrackings.EntityData.YFilter = objectTrackings.YFilter
    objectTrackings.EntityData.YangName = "object-trackings"
    objectTrackings.EntityData.BundleName = "cisco_ios_xr"
    objectTrackings.EntityData.ParentYangName = "Cisco-IOS-XR-manageability-object-tracking-cfg"
    objectTrackings.EntityData.SegmentPath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings"
    objectTrackings.EntityData.AbsolutePath = objectTrackings.EntityData.SegmentPath
    objectTrackings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectTrackings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectTrackings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectTrackings.EntityData.Children = types.NewOrderedMap()
    objectTrackings.EntityData.Children.Append("object-tracking", types.YChild{"ObjectTracking", nil})
    for i := range objectTrackings.ObjectTracking {
        objectTrackings.EntityData.Children.Append(types.GetSegmentPath(objectTrackings.ObjectTracking[i]), types.YChild{"ObjectTracking", objectTrackings.ObjectTracking[i]})
    }
    objectTrackings.EntityData.Leafs = types.NewOrderedMap()

    objectTrackings.EntityData.YListKeys = []string {}

    return &(objectTrackings.EntityData)
}

// ObjectTrackings_ObjectTracking
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Actions associated with track state changes.
    Action ObjectTrackings_ObjectTracking_Action

    // Track type BFD RTR (BFD Response Time Reporter).
    TypeBfdRtr ObjectTrackings_ObjectTracking_TypeBfdRtr

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
    objectTracking.EntityData.SegmentPath = "object-tracking" + types.AddKeyToken(objectTracking.TrackName, "track-name")
    objectTracking.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/" + objectTracking.EntityData.SegmentPath
    objectTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectTracking.EntityData.Children = types.NewOrderedMap()
    objectTracking.EntityData.Children.Append("action", types.YChild{"Action", &objectTracking.Action})
    objectTracking.EntityData.Children.Append("type-bfd-rtr", types.YChild{"TypeBfdRtr", &objectTracking.TypeBfdRtr})
    objectTracking.EntityData.Children.Append("type-interface", types.YChild{"TypeInterface", &objectTracking.TypeInterface})
    objectTracking.EntityData.Children.Append("type-rtr", types.YChild{"TypeRtr", &objectTracking.TypeRtr})
    objectTracking.EntityData.Children.Append("type-list", types.YChild{"TypeList", &objectTracking.TypeList})
    objectTracking.EntityData.Children.Append("type-route", types.YChild{"TypeRoute", &objectTracking.TypeRoute})
    objectTracking.EntityData.Children.Append("type-boolean-list", types.YChild{"TypeBooleanList", &objectTracking.TypeBooleanList})
    objectTracking.EntityData.Leafs = types.NewOrderedMap()
    objectTracking.EntityData.Leafs.Append("track-name", types.YLeaf{"TrackName", objectTracking.TrackName})
    objectTracking.EntityData.Leafs.Append("delay-up", types.YLeaf{"DelayUp", objectTracking.DelayUp})
    objectTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", objectTracking.Enable})
    objectTracking.EntityData.Leafs.Append("delay-down", types.YLeaf{"DelayDown", objectTracking.DelayDown})
    objectTracking.EntityData.Leafs.Append("type-interface-enable", types.YLeaf{"TypeInterfaceEnable", objectTracking.TypeInterfaceEnable})
    objectTracking.EntityData.Leafs.Append("type-route-enable", types.YLeaf{"TypeRouteEnable", objectTracking.TypeRouteEnable})
    objectTracking.EntityData.Leafs.Append("type-boolean-list-and-enable", types.YLeaf{"TypeBooleanListAndEnable", objectTracking.TypeBooleanListAndEnable})
    objectTracking.EntityData.Leafs.Append("type-boolean-list-or-enable", types.YLeaf{"TypeBooleanListOrEnable", objectTracking.TypeBooleanListOrEnable})

    objectTracking.EntityData.YListKeys = []string {"TrackName"}

    return &(objectTracking.EntityData)
}

// ObjectTrackings_ObjectTracking_Action
// Actions associated with track state changes
type ObjectTrackings_ObjectTracking_Action struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable track actions. The type is interface{}.
    ActionsEnable interface{}

    // The list of all track actions.
    ActionErrDis ObjectTrackings_ObjectTracking_Action_ActionErrDis
}

func (action *ObjectTrackings_ObjectTracking_Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "object-tracking"
    action.EntityData.SegmentPath = "action"
    action.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/" + action.EntityData.SegmentPath
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = types.NewOrderedMap()
    action.EntityData.Children.Append("action-err-dis", types.YChild{"ActionErrDis", &action.ActionErrDis})
    action.EntityData.Leafs = types.NewOrderedMap()
    action.EntityData.Leafs.Append("actions-enable", types.YLeaf{"ActionsEnable", action.ActionsEnable})

    action.EntityData.YListKeys = []string {}

    return &(action.EntityData)
}

// ObjectTrackings_ObjectTracking_Action_ActionErrDis
// The list of all track actions
type ObjectTrackings_ObjectTracking_Action_ActionErrDis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error disable track action. The type is slice of
    // ObjectTrackings_ObjectTracking_Action_ActionErrDis_ActionErrDi.
    ActionErrDi []*ObjectTrackings_ObjectTracking_Action_ActionErrDis_ActionErrDi
}

func (actionErrDis *ObjectTrackings_ObjectTracking_Action_ActionErrDis) GetEntityData() *types.CommonEntityData {
    actionErrDis.EntityData.YFilter = actionErrDis.YFilter
    actionErrDis.EntityData.YangName = "action-err-dis"
    actionErrDis.EntityData.BundleName = "cisco_ios_xr"
    actionErrDis.EntityData.ParentYangName = "action"
    actionErrDis.EntityData.SegmentPath = "action-err-dis"
    actionErrDis.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/action/" + actionErrDis.EntityData.SegmentPath
    actionErrDis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionErrDis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionErrDis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionErrDis.EntityData.Children = types.NewOrderedMap()
    actionErrDis.EntityData.Children.Append("action-err-di", types.YChild{"ActionErrDi", nil})
    for i := range actionErrDis.ActionErrDi {
        actionErrDis.EntityData.Children.Append(types.GetSegmentPath(actionErrDis.ActionErrDi[i]), types.YChild{"ActionErrDi", actionErrDis.ActionErrDi[i]})
    }
    actionErrDis.EntityData.Leafs = types.NewOrderedMap()

    actionErrDis.EntityData.YListKeys = []string {}

    return &(actionErrDis.EntityData)
}

// ObjectTrackings_ObjectTracking_Action_ActionErrDis_ActionErrDi
// Error disable track action
type ObjectTrackings_ObjectTracking_Action_ActionErrDis_ActionErrDi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Track State Type. The type is interface{} with
    // range: 0..1.
    TrackStateType interface{}

    // This attribute is a key. Interface to be error-disabled. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}
}

func (actionErrDi *ObjectTrackings_ObjectTracking_Action_ActionErrDis_ActionErrDi) GetEntityData() *types.CommonEntityData {
    actionErrDi.EntityData.YFilter = actionErrDi.YFilter
    actionErrDi.EntityData.YangName = "action-err-di"
    actionErrDi.EntityData.BundleName = "cisco_ios_xr"
    actionErrDi.EntityData.ParentYangName = "action-err-dis"
    actionErrDi.EntityData.SegmentPath = "action-err-di" + types.AddKeyToken(actionErrDi.TrackStateType, "track-state-type") + types.AddKeyToken(actionErrDi.InterfaceName, "interface-name")
    actionErrDi.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/action/action-err-dis/" + actionErrDi.EntityData.SegmentPath
    actionErrDi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionErrDi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionErrDi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionErrDi.EntityData.Children = types.NewOrderedMap()
    actionErrDi.EntityData.Leafs = types.NewOrderedMap()
    actionErrDi.EntityData.Leafs.Append("track-state-type", types.YLeaf{"TrackStateType", actionErrDi.TrackStateType})
    actionErrDi.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", actionErrDi.InterfaceName})

    actionErrDi.EntityData.YListKeys = []string {"TrackStateType", "InterfaceName"}

    return &(actionErrDi.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBfdRtr
// Track type BFD RTR (BFD Response Time Reporter)
type ObjectTrackings_ObjectTracking_TypeBfdRtr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BFD session related parameters.
    BfdRtr ObjectTrackings_ObjectTracking_TypeBfdRtr_BfdRtr
}

func (typeBfdRtr *ObjectTrackings_ObjectTracking_TypeBfdRtr) GetEntityData() *types.CommonEntityData {
    typeBfdRtr.EntityData.YFilter = typeBfdRtr.YFilter
    typeBfdRtr.EntityData.YangName = "type-bfd-rtr"
    typeBfdRtr.EntityData.BundleName = "cisco_ios_xr"
    typeBfdRtr.EntityData.ParentYangName = "object-tracking"
    typeBfdRtr.EntityData.SegmentPath = "type-bfd-rtr"
    typeBfdRtr.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/" + typeBfdRtr.EntityData.SegmentPath
    typeBfdRtr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeBfdRtr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeBfdRtr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeBfdRtr.EntityData.Children = types.NewOrderedMap()
    typeBfdRtr.EntityData.Children.Append("bfd-rtr", types.YChild{"BfdRtr", &typeBfdRtr.BfdRtr})
    typeBfdRtr.EntityData.Leafs = types.NewOrderedMap()

    typeBfdRtr.EntityData.YListKeys = []string {}

    return &(typeBfdRtr.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBfdRtr_BfdRtr
// BFD session related parameters
// This type is a presence type.
type ObjectTrackings_ObjectTracking_TypeBfdRtr_BfdRtr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Tx interval in ms. The type is interface{} with range: 1..5000. This
    // attribute is mandatory.
    Rate interface{}

    // Debounce Count. The type is interface{} with range: 1..10. This attribute
    // is mandatory.
    DebounceCount interface{}

    // Interface to be used for BFD session. The type is string with pattern:
    // [a-zA-Z0-9._/-]+. This attribute is mandatory.
    InterfaceName interface{}

    // Destination IP Address to track via BFD. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    DestAddress interface{}
}

func (bfdRtr *ObjectTrackings_ObjectTracking_TypeBfdRtr_BfdRtr) GetEntityData() *types.CommonEntityData {
    bfdRtr.EntityData.YFilter = bfdRtr.YFilter
    bfdRtr.EntityData.YangName = "bfd-rtr"
    bfdRtr.EntityData.BundleName = "cisco_ios_xr"
    bfdRtr.EntityData.ParentYangName = "type-bfd-rtr"
    bfdRtr.EntityData.SegmentPath = "bfd-rtr"
    bfdRtr.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-bfd-rtr/" + bfdRtr.EntityData.SegmentPath
    bfdRtr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdRtr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdRtr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdRtr.EntityData.Children = types.NewOrderedMap()
    bfdRtr.EntityData.Leafs = types.NewOrderedMap()
    bfdRtr.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", bfdRtr.Rate})
    bfdRtr.EntityData.Leafs.Append("debounce-count", types.YLeaf{"DebounceCount", bfdRtr.DebounceCount})
    bfdRtr.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", bfdRtr.InterfaceName})
    bfdRtr.EntityData.Leafs.Append("dest-address", types.YLeaf{"DestAddress", bfdRtr.DestAddress})

    bfdRtr.EntityData.YListKeys = []string {}

    return &(bfdRtr.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeInterface
// Track type line-protocol
type ObjectTrackings_ObjectTracking_TypeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Interface interface{}
}

func (typeInterface *ObjectTrackings_ObjectTracking_TypeInterface) GetEntityData() *types.CommonEntityData {
    typeInterface.EntityData.YFilter = typeInterface.YFilter
    typeInterface.EntityData.YangName = "type-interface"
    typeInterface.EntityData.BundleName = "cisco_ios_xr"
    typeInterface.EntityData.ParentYangName = "object-tracking"
    typeInterface.EntityData.SegmentPath = "type-interface"
    typeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/" + typeInterface.EntityData.SegmentPath
    typeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeInterface.EntityData.Children = types.NewOrderedMap()
    typeInterface.EntityData.Leafs = types.NewOrderedMap()
    typeInterface.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", typeInterface.Interface})

    typeInterface.EntityData.YListKeys = []string {}

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
    typeRtr.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/" + typeRtr.EntityData.SegmentPath
    typeRtr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeRtr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeRtr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeRtr.EntityData.Children = types.NewOrderedMap()
    typeRtr.EntityData.Leafs = types.NewOrderedMap()
    typeRtr.EntityData.Leafs.Append("rtr", types.YLeaf{"Rtr", typeRtr.Rtr})

    typeRtr.EntityData.YListKeys = []string {}

    return &(typeRtr.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList
// Track type boolean list
type ObjectTrackings_ObjectTracking_TypeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable threshold based on percentage. The type is interface{}. Units are
    // percentage.
    ThresholdPercentageObjectEnable interface{}

    // Enable threshold based on weighted sum. The type is interface{}.
    ThresholdWeightObjectEnable interface{}

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
    typeList.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/" + typeList.EntityData.SegmentPath
    typeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeList.EntityData.Children = types.NewOrderedMap()
    typeList.EntityData.Children.Append("threshold-weight", types.YChild{"ThresholdWeight", &typeList.ThresholdWeight})
    typeList.EntityData.Children.Append("threshold-percentage-object", types.YChild{"ThresholdPercentageObject", &typeList.ThresholdPercentageObject})
    typeList.EntityData.Children.Append("threshold-percentage", types.YChild{"ThresholdPercentage", &typeList.ThresholdPercentage})
    typeList.EntityData.Children.Append("threshold-weight-object", types.YChild{"ThresholdWeightObject", &typeList.ThresholdWeightObject})
    typeList.EntityData.Leafs = types.NewOrderedMap()
    typeList.EntityData.Leafs.Append("threshold-percentage-object-enable", types.YLeaf{"ThresholdPercentageObjectEnable", typeList.ThresholdPercentageObjectEnable})
    typeList.EntityData.Leafs.Append("threshold-weight-object-enable", types.YLeaf{"ThresholdWeightObjectEnable", typeList.ThresholdWeightObjectEnable})

    typeList.EntityData.YListKeys = []string {}

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
    thresholdWeight.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/" + thresholdWeight.EntityData.SegmentPath
    thresholdWeight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdWeight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdWeight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdWeight.EntityData.Children = types.NewOrderedMap()
    thresholdWeight.EntityData.Children.Append("threshold-limits", types.YChild{"ThresholdLimits", &thresholdWeight.ThresholdLimits})
    thresholdWeight.EntityData.Leafs = types.NewOrderedMap()

    thresholdWeight.EntityData.YListKeys = []string {}

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
    thresholdLimits.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/threshold-weight/" + thresholdLimits.EntityData.SegmentPath
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = types.NewOrderedMap()
    thresholdLimits.EntityData.Children.Append("threshold-up-values", types.YChild{"ThresholdUpValues", &thresholdLimits.ThresholdUpValues})
    thresholdLimits.EntityData.Leafs = types.NewOrderedMap()

    thresholdLimits.EntityData.YListKeys = []string {}

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
    ThresholdUpValue []*ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues) GetEntityData() *types.CommonEntityData {
    thresholdUpValues.EntityData.YFilter = thresholdUpValues.YFilter
    thresholdUpValues.EntityData.YangName = "threshold-up-values"
    thresholdUpValues.EntityData.BundleName = "cisco_ios_xr"
    thresholdUpValues.EntityData.ParentYangName = "threshold-limits"
    thresholdUpValues.EntityData.SegmentPath = "threshold-up-values"
    thresholdUpValues.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/threshold-weight/threshold-limits/" + thresholdUpValues.EntityData.SegmentPath
    thresholdUpValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdUpValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdUpValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdUpValues.EntityData.Children = types.NewOrderedMap()
    thresholdUpValues.EntityData.Children.Append("threshold-up-value", types.YChild{"ThresholdUpValue", nil})
    for i := range thresholdUpValues.ThresholdUpValue {
        thresholdUpValues.EntityData.Children.Append(types.GetSegmentPath(thresholdUpValues.ThresholdUpValue[i]), types.YChild{"ThresholdUpValue", thresholdUpValues.ThresholdUpValue[i]})
    }
    thresholdUpValues.EntityData.Leafs = types.NewOrderedMap()

    thresholdUpValues.EntityData.YListKeys = []string {}

    return &(thresholdUpValues.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Up value. The type is interface{} with range:
    // 0..4294967295.
    Up interface{}

    // Threshold limit at which track is set to Down state. The type is
    // interface{} with range: 0..4294967295. The default value is 0.
    ThresholdDown interface{}
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeight_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetEntityData() *types.CommonEntityData {
    thresholdUpValue.EntityData.YFilter = thresholdUpValue.YFilter
    thresholdUpValue.EntityData.YangName = "threshold-up-value"
    thresholdUpValue.EntityData.BundleName = "cisco_ios_xr"
    thresholdUpValue.EntityData.ParentYangName = "threshold-up-values"
    thresholdUpValue.EntityData.SegmentPath = "threshold-up-value" + types.AddKeyToken(thresholdUpValue.Up, "up")
    thresholdUpValue.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/threshold-weight/threshold-limits/threshold-up-values/" + thresholdUpValue.EntityData.SegmentPath
    thresholdUpValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdUpValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdUpValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdUpValue.EntityData.Children = types.NewOrderedMap()
    thresholdUpValue.EntityData.Leafs = types.NewOrderedMap()
    thresholdUpValue.EntityData.Leafs.Append("up", types.YLeaf{"Up", thresholdUpValue.Up})
    thresholdUpValue.EntityData.Leafs.Append("threshold-down", types.YLeaf{"ThresholdDown", thresholdUpValue.ThresholdDown})

    thresholdUpValue.EntityData.YListKeys = []string {"Up"}

    return &(thresholdUpValue.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject
// Track type threshold percentage
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name object. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object.
    Object []*ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object
}

func (thresholdPercentageObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject) GetEntityData() *types.CommonEntityData {
    thresholdPercentageObject.EntityData.YFilter = thresholdPercentageObject.YFilter
    thresholdPercentageObject.EntityData.YangName = "threshold-percentage-object"
    thresholdPercentageObject.EntityData.BundleName = "cisco_ios_xr"
    thresholdPercentageObject.EntityData.ParentYangName = "type-list"
    thresholdPercentageObject.EntityData.SegmentPath = "threshold-percentage-object"
    thresholdPercentageObject.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/" + thresholdPercentageObject.EntityData.SegmentPath
    thresholdPercentageObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdPercentageObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdPercentageObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdPercentageObject.EntityData.Children = types.NewOrderedMap()
    thresholdPercentageObject.EntityData.Children.Append("object", types.YChild{"Object", nil})
    for i := range thresholdPercentageObject.Object {
        thresholdPercentageObject.EntityData.Children.Append(types.GetSegmentPath(thresholdPercentageObject.Object[i]), types.YChild{"Object", thresholdPercentageObject.Object[i]})
    }
    thresholdPercentageObject.EntityData.Leafs = types.NewOrderedMap()

    thresholdPercentageObject.EntityData.YListKeys = []string {}

    return &(thresholdPercentageObject.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object
// Track name object
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    Object interface{}

    // Weight of object. The type is interface{} with range: 0..4294967295. The
    // default value is 1.
    ObjectWeight interface{}
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentageObject_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "threshold-percentage-object"
    object.EntityData.SegmentPath = "object" + types.AddKeyToken(object.Object, "object")
    object.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/threshold-percentage-object/" + object.EntityData.SegmentPath
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = types.NewOrderedMap()
    object.EntityData.Leafs = types.NewOrderedMap()
    object.EntityData.Leafs.Append("object", types.YLeaf{"Object", object.Object})
    object.EntityData.Leafs.Append("object-weight", types.YLeaf{"ObjectWeight", object.ObjectWeight})

    object.EntityData.YListKeys = []string {"Object"}

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
    thresholdPercentage.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/" + thresholdPercentage.EntityData.SegmentPath
    thresholdPercentage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdPercentage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdPercentage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdPercentage.EntityData.Children = types.NewOrderedMap()
    thresholdPercentage.EntityData.Children.Append("threshold-limits", types.YChild{"ThresholdLimits", &thresholdPercentage.ThresholdLimits})
    thresholdPercentage.EntityData.Leafs = types.NewOrderedMap()

    thresholdPercentage.EntityData.YListKeys = []string {}

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
    thresholdLimits.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/threshold-percentage/" + thresholdLimits.EntityData.SegmentPath
    thresholdLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdLimits.EntityData.Children = types.NewOrderedMap()
    thresholdLimits.EntityData.Children.Append("threshold-up-values", types.YChild{"ThresholdUpValues", &thresholdLimits.ThresholdUpValues})
    thresholdLimits.EntityData.Leafs = types.NewOrderedMap()

    thresholdLimits.EntityData.YListKeys = []string {}

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
    ThresholdUpValue []*ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
}

func (thresholdUpValues *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues) GetEntityData() *types.CommonEntityData {
    thresholdUpValues.EntityData.YFilter = thresholdUpValues.YFilter
    thresholdUpValues.EntityData.YangName = "threshold-up-values"
    thresholdUpValues.EntityData.BundleName = "cisco_ios_xr"
    thresholdUpValues.EntityData.ParentYangName = "threshold-limits"
    thresholdUpValues.EntityData.SegmentPath = "threshold-up-values"
    thresholdUpValues.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/threshold-percentage/threshold-limits/" + thresholdUpValues.EntityData.SegmentPath
    thresholdUpValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdUpValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdUpValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdUpValues.EntityData.Children = types.NewOrderedMap()
    thresholdUpValues.EntityData.Children.Append("threshold-up-value", types.YChild{"ThresholdUpValue", nil})
    for i := range thresholdUpValues.ThresholdUpValue {
        thresholdUpValues.EntityData.Children.Append(types.GetSegmentPath(thresholdUpValues.ThresholdUpValue[i]), types.YChild{"ThresholdUpValue", thresholdUpValues.ThresholdUpValue[i]})
    }
    thresholdUpValues.EntityData.Leafs = types.NewOrderedMap()

    thresholdUpValues.EntityData.YListKeys = []string {}

    return &(thresholdUpValues.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue
// Threshold limit at which track is set to UP
// state
type ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Up value. The type is interface{} with range:
    // 0..4294967295.
    Up interface{}

    // Threshold limit at which track is set to Down state. The type is
    // interface{} with range: 0..4294967295. The default value is 0.
    ThresholdDown interface{}
}

func (thresholdUpValue *ObjectTrackings_ObjectTracking_TypeList_ThresholdPercentage_ThresholdLimits_ThresholdUpValues_ThresholdUpValue) GetEntityData() *types.CommonEntityData {
    thresholdUpValue.EntityData.YFilter = thresholdUpValue.YFilter
    thresholdUpValue.EntityData.YangName = "threshold-up-value"
    thresholdUpValue.EntityData.BundleName = "cisco_ios_xr"
    thresholdUpValue.EntityData.ParentYangName = "threshold-up-values"
    thresholdUpValue.EntityData.SegmentPath = "threshold-up-value" + types.AddKeyToken(thresholdUpValue.Up, "up")
    thresholdUpValue.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/threshold-percentage/threshold-limits/threshold-up-values/" + thresholdUpValue.EntityData.SegmentPath
    thresholdUpValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdUpValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdUpValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdUpValue.EntityData.Children = types.NewOrderedMap()
    thresholdUpValue.EntityData.Leafs = types.NewOrderedMap()
    thresholdUpValue.EntityData.Leafs.Append("up", types.YLeaf{"Up", thresholdUpValue.Up})
    thresholdUpValue.EntityData.Leafs.Append("threshold-down", types.YLeaf{"ThresholdDown", thresholdUpValue.ThresholdDown})

    thresholdUpValue.EntityData.YListKeys = []string {"Up"}

    return &(thresholdUpValue.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject
// Track type threshold weight
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name object. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object.
    Object []*ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object
}

func (thresholdWeightObject *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject) GetEntityData() *types.CommonEntityData {
    thresholdWeightObject.EntityData.YFilter = thresholdWeightObject.YFilter
    thresholdWeightObject.EntityData.YangName = "threshold-weight-object"
    thresholdWeightObject.EntityData.BundleName = "cisco_ios_xr"
    thresholdWeightObject.EntityData.ParentYangName = "type-list"
    thresholdWeightObject.EntityData.SegmentPath = "threshold-weight-object"
    thresholdWeightObject.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/" + thresholdWeightObject.EntityData.SegmentPath
    thresholdWeightObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdWeightObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdWeightObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdWeightObject.EntityData.Children = types.NewOrderedMap()
    thresholdWeightObject.EntityData.Children.Append("object", types.YChild{"Object", nil})
    for i := range thresholdWeightObject.Object {
        thresholdWeightObject.EntityData.Children.Append(types.GetSegmentPath(thresholdWeightObject.Object[i]), types.YChild{"Object", thresholdWeightObject.Object[i]})
    }
    thresholdWeightObject.EntityData.Leafs = types.NewOrderedMap()

    thresholdWeightObject.EntityData.YListKeys = []string {}

    return &(thresholdWeightObject.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object
// Track name object
type ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object name. The type is string with length:
    // 1..32.
    Object interface{}

    // Weight of object. The type is interface{} with range: 0..4294967295. The
    // default value is 1.
    ObjectWeight interface{}
}

func (object *ObjectTrackings_ObjectTracking_TypeList_ThresholdWeightObject_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "threshold-weight-object"
    object.EntityData.SegmentPath = "object" + types.AddKeyToken(object.Object, "object")
    object.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-list/threshold-weight-object/" + object.EntityData.SegmentPath
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = types.NewOrderedMap()
    object.EntityData.Leafs = types.NewOrderedMap()
    object.EntityData.Leafs.Append("object", types.YLeaf{"Object", object.Object})
    object.EntityData.Leafs.Append("object-weight", types.YLeaf{"ObjectWeight", object.ObjectWeight})

    object.EntityData.YListKeys = []string {"Object"}

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
    typeRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/" + typeRoute.EntityData.SegmentPath
    typeRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeRoute.EntityData.Children = types.NewOrderedMap()
    typeRoute.EntityData.Children.Append("ip-address", types.YChild{"IpAddress", &typeRoute.IpAddress})
    typeRoute.EntityData.Leafs = types.NewOrderedMap()
    typeRoute.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", typeRoute.Vrf})

    typeRoute.EntityData.YListKeys = []string {}

    return &(typeRoute.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeRoute_IpAddress
// set track IPv4 address
// This type is a presence type.
type ObjectTrackings_ObjectTracking_TypeRoute_IpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Mask interface{}
}

func (ipAddress *ObjectTrackings_ObjectTracking_TypeRoute_IpAddress) GetEntityData() *types.CommonEntityData {
    ipAddress.EntityData.YFilter = ipAddress.YFilter
    ipAddress.EntityData.YangName = "ip-address"
    ipAddress.EntityData.BundleName = "cisco_ios_xr"
    ipAddress.EntityData.ParentYangName = "type-route"
    ipAddress.EntityData.SegmentPath = "ip-address"
    ipAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-route/" + ipAddress.EntityData.SegmentPath
    ipAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipAddress.EntityData.Children = types.NewOrderedMap()
    ipAddress.EntityData.Leafs = types.NewOrderedMap()
    ipAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipAddress.Address})
    ipAddress.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", ipAddress.Mask})

    ipAddress.EntityData.YListKeys = []string {}

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
    typeBooleanList.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/" + typeBooleanList.EntityData.SegmentPath
    typeBooleanList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    typeBooleanList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    typeBooleanList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    typeBooleanList.EntityData.Children = types.NewOrderedMap()
    typeBooleanList.EntityData.Children.Append("or-objects", types.YChild{"OrObjects", &typeBooleanList.OrObjects})
    typeBooleanList.EntityData.Children.Append("and-objects", types.YChild{"AndObjects", &typeBooleanList.AndObjects})
    typeBooleanList.EntityData.Leafs = types.NewOrderedMap()

    typeBooleanList.EntityData.YListKeys = []string {}

    return &(typeBooleanList.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects
// Track type boolean or list
type ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject.
    OrObject []*ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject
}

func (orObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects) GetEntityData() *types.CommonEntityData {
    orObjects.EntityData.YFilter = orObjects.YFilter
    orObjects.EntityData.YangName = "or-objects"
    orObjects.EntityData.BundleName = "cisco_ios_xr"
    orObjects.EntityData.ParentYangName = "type-boolean-list"
    orObjects.EntityData.SegmentPath = "or-objects"
    orObjects.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-boolean-list/" + orObjects.EntityData.SegmentPath
    orObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orObjects.EntityData.Children = types.NewOrderedMap()
    orObjects.EntityData.Children.Append("or-object", types.YChild{"OrObject", nil})
    for i := range orObjects.OrObject {
        orObjects.EntityData.Children.Append(types.GetSegmentPath(orObjects.OrObject[i]), types.YChild{"OrObject", orObjects.OrObject[i]})
    }
    orObjects.EntityData.Leafs = types.NewOrderedMap()

    orObjects.EntityData.YListKeys = []string {}

    return &(orObjects.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking_TypeBooleanList_OrObjects_OrObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    orObject.EntityData.SegmentPath = "or-object" + types.AddKeyToken(orObject.Object, "object")
    orObject.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-boolean-list/or-objects/" + orObject.EntityData.SegmentPath
    orObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orObject.EntityData.Children = types.NewOrderedMap()
    orObject.EntityData.Leafs = types.NewOrderedMap()
    orObject.EntityData.Leafs.Append("object", types.YLeaf{"Object", orObject.Object})
    orObject.EntityData.Leafs.Append("object-sign", types.YLeaf{"ObjectSign", orObject.ObjectSign})

    orObject.EntityData.YListKeys = []string {"Object"}

    return &(orObject.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects
// Track type boolean and list
type ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Track name - maximum 32 characters. The type is slice of
    // ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject.
    AndObject []*ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject
}

func (andObjects *ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects) GetEntityData() *types.CommonEntityData {
    andObjects.EntityData.YFilter = andObjects.YFilter
    andObjects.EntityData.YangName = "and-objects"
    andObjects.EntityData.BundleName = "cisco_ios_xr"
    andObjects.EntityData.ParentYangName = "type-boolean-list"
    andObjects.EntityData.SegmentPath = "and-objects"
    andObjects.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-boolean-list/" + andObjects.EntityData.SegmentPath
    andObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    andObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    andObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    andObjects.EntityData.Children = types.NewOrderedMap()
    andObjects.EntityData.Children.Append("and-object", types.YChild{"AndObject", nil})
    for i := range andObjects.AndObject {
        andObjects.EntityData.Children.Append(types.GetSegmentPath(andObjects.AndObject[i]), types.YChild{"AndObject", andObjects.AndObject[i]})
    }
    andObjects.EntityData.Leafs = types.NewOrderedMap()

    andObjects.EntityData.YListKeys = []string {}

    return &(andObjects.EntityData)
}

// ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject
// Track name - maximum 32 characters
type ObjectTrackings_ObjectTracking_TypeBooleanList_AndObjects_AndObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    andObject.EntityData.SegmentPath = "and-object" + types.AddKeyToken(andObject.ObjectName, "object-name")
    andObject.EntityData.AbsolutePath = "Cisco-IOS-XR-manageability-object-tracking-cfg:object-trackings/object-tracking/type-boolean-list/and-objects/" + andObject.EntityData.SegmentPath
    andObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    andObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    andObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    andObject.EntityData.Children = types.NewOrderedMap()
    andObject.EntityData.Leafs = types.NewOrderedMap()
    andObject.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", andObject.ObjectName})
    andObject.EntityData.Leafs.Append("object-sign", types.YLeaf{"ObjectSign", andObject.ObjectSign})

    andObject.EntityData.YListKeys = []string {"ObjectName"}

    return &(andObject.EntityData)
}

