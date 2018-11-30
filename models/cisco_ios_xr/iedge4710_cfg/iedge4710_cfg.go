// This module contains a collection of YANG definitions
// for Cisco IOS-XR iedge4710 package configuration.
// 
// This module contains definitions
// for the following management objects:
//   subscriber-manager: iEdge subscriber manager configuration
//   subscriber-featurette: subscriber featurette
//   iedge-license-manager: iedge license manager
//   sub-manager: sub manager
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package iedge4710_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package iedge4710_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-iedge4710-cfg subscriber-manager}", reflect.TypeOf(SubscriberManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-iedge4710-cfg:subscriber-manager", reflect.TypeOf(SubscriberManager{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-iedge4710-cfg subscriber-featurette}", reflect.TypeOf(SubscriberFeaturette{}))
    ydk.RegisterEntity("Cisco-IOS-XR-iedge4710-cfg:subscriber-featurette", reflect.TypeOf(SubscriberFeaturette{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-iedge4710-cfg iedge-license-manager}", reflect.TypeOf(IedgeLicenseManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-iedge4710-cfg:iedge-license-manager", reflect.TypeOf(IedgeLicenseManager{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-iedge4710-cfg sub-manager}", reflect.TypeOf(SubManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-iedge4710-cfg:sub-manager", reflect.TypeOf(SubManager{}))
}

// SubscriberManager
// iEdge subscriber manager configuration
type SubscriberManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // iEdge accounting feature.
    Accounting SubscriberManager_Accounting

    // SRG specific config.
    Srg SubscriberManager_Srg
}

func (subscriberManager *SubscriberManager) GetEntityData() *types.CommonEntityData {
    subscriberManager.EntityData.YFilter = subscriberManager.YFilter
    subscriberManager.EntityData.YangName = "subscriber-manager"
    subscriberManager.EntityData.BundleName = "cisco_ios_xr"
    subscriberManager.EntityData.ParentYangName = "Cisco-IOS-XR-iedge4710-cfg"
    subscriberManager.EntityData.SegmentPath = "Cisco-IOS-XR-iedge4710-cfg:subscriber-manager"
    subscriberManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberManager.EntityData.Children = types.NewOrderedMap()
    subscriberManager.EntityData.Children.Append("accounting", types.YChild{"Accounting", &subscriberManager.Accounting})
    subscriberManager.EntityData.Children.Append("srg", types.YChild{"Srg", &subscriberManager.Srg})
    subscriberManager.EntityData.Leafs = types.NewOrderedMap()

    subscriberManager.EntityData.YListKeys = []string {}

    return &(subscriberManager.EntityData)
}

// SubscriberManager_Accounting
// iEdge accounting feature
type SubscriberManager_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accounting send stop feature.
    SendStop SubscriberManager_Accounting_SendStop

    // interim accounting related.
    Interim SubscriberManager_Accounting_Interim
}

func (accounting *SubscriberManager_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "subscriber-manager"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("send-stop", types.YChild{"SendStop", &accounting.SendStop})
    accounting.EntityData.Children.Append("interim", types.YChild{"Interim", &accounting.Interim})
    accounting.EntityData.Leafs = types.NewOrderedMap()

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// SubscriberManager_Accounting_SendStop
// Accounting send stop feature
type SubscriberManager_Accounting_SendStop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Setup failure feature.
    SetupFailure SubscriberManager_Accounting_SendStop_SetupFailure
}

func (sendStop *SubscriberManager_Accounting_SendStop) GetEntityData() *types.CommonEntityData {
    sendStop.EntityData.YFilter = sendStop.YFilter
    sendStop.EntityData.YangName = "send-stop"
    sendStop.EntityData.BundleName = "cisco_ios_xr"
    sendStop.EntityData.ParentYangName = "accounting"
    sendStop.EntityData.SegmentPath = "send-stop"
    sendStop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendStop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendStop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendStop.EntityData.Children = types.NewOrderedMap()
    sendStop.EntityData.Children.Append("setup-failure", types.YChild{"SetupFailure", &sendStop.SetupFailure})
    sendStop.EntityData.Leafs = types.NewOrderedMap()

    sendStop.EntityData.YListKeys = []string {}

    return &(sendStop.EntityData)
}

// SubscriberManager_Accounting_SendStop_SetupFailure
// Setup failure feature
type SubscriberManager_Accounting_SendStop_SetupFailure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA List name either default or preconfigured. The type is string.
    MethodListName interface{}
}

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetEntityData() *types.CommonEntityData {
    setupFailure.EntityData.YFilter = setupFailure.YFilter
    setupFailure.EntityData.YangName = "setup-failure"
    setupFailure.EntityData.BundleName = "cisco_ios_xr"
    setupFailure.EntityData.ParentYangName = "send-stop"
    setupFailure.EntityData.SegmentPath = "setup-failure"
    setupFailure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    setupFailure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    setupFailure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    setupFailure.EntityData.Children = types.NewOrderedMap()
    setupFailure.EntityData.Leafs = types.NewOrderedMap()
    setupFailure.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", setupFailure.MethodListName})

    setupFailure.EntityData.YListKeys = []string {}

    return &(setupFailure.EntityData)
}

// SubscriberManager_Accounting_Interim
// interim accounting related
type SubscriberManager_Accounting_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // variation of first session or service interim record from configured
    // timeout.
    Variation SubscriberManager_Accounting_Interim_Variation
}

func (interim *SubscriberManager_Accounting_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "accounting"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = types.NewOrderedMap()
    interim.EntityData.Children.Append("variation", types.YChild{"Variation", &interim.Variation})
    interim.EntityData.Leafs = types.NewOrderedMap()

    interim.EntityData.YListKeys = []string {}

    return &(interim.EntityData)
}

// SubscriberManager_Accounting_Interim_Variation
// variation of first session or service interim
// record from configured timeout
type SubscriberManager_Accounting_Interim_Variation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // maximum percentage variation (maximum absolute variation is 15 minutes).
    // The type is interface{} with range: 0..50. Units are percentage.
    MaximumPercentageVariation interface{}
}

func (variation *SubscriberManager_Accounting_Interim_Variation) GetEntityData() *types.CommonEntityData {
    variation.EntityData.YFilter = variation.YFilter
    variation.EntityData.YangName = "variation"
    variation.EntityData.BundleName = "cisco_ios_xr"
    variation.EntityData.ParentYangName = "interim"
    variation.EntityData.SegmentPath = "variation"
    variation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    variation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    variation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    variation.EntityData.Children = types.NewOrderedMap()
    variation.EntityData.Leafs = types.NewOrderedMap()
    variation.EntityData.Leafs.Append("maximum-percentage-variation", types.YLeaf{"MaximumPercentageVariation", variation.MaximumPercentageVariation})

    variation.EntityData.YListKeys = []string {}

    return &(variation.EntityData)
}

// SubscriberManager_Srg
// SRG specific config
type SubscriberManager_Srg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sync account session id from master to slave. The type is interface{}.
    SyncAccountSessionId interface{}
}

func (srg *SubscriberManager_Srg) GetEntityData() *types.CommonEntityData {
    srg.EntityData.YFilter = srg.YFilter
    srg.EntityData.YangName = "srg"
    srg.EntityData.BundleName = "cisco_ios_xr"
    srg.EntityData.ParentYangName = "subscriber-manager"
    srg.EntityData.SegmentPath = "srg"
    srg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srg.EntityData.Children = types.NewOrderedMap()
    srg.EntityData.Leafs = types.NewOrderedMap()
    srg.EntityData.Leafs.Append("sync-account-session-id", types.YLeaf{"SyncAccountSessionId", srg.SyncAccountSessionId})

    srg.EntityData.YListKeys = []string {}

    return &(srg.EntityData)
}

// SubscriberFeaturette
// subscriber featurette
type SubscriberFeaturette struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enable featurette processing. The type is slice of
    // SubscriberFeaturette_FeaturetteName.
    FeaturetteName []*SubscriberFeaturette_FeaturetteName
}

func (subscriberFeaturette *SubscriberFeaturette) GetEntityData() *types.CommonEntityData {
    subscriberFeaturette.EntityData.YFilter = subscriberFeaturette.YFilter
    subscriberFeaturette.EntityData.YangName = "subscriber-featurette"
    subscriberFeaturette.EntityData.BundleName = "cisco_ios_xr"
    subscriberFeaturette.EntityData.ParentYangName = "Cisco-IOS-XR-iedge4710-cfg"
    subscriberFeaturette.EntityData.SegmentPath = "Cisco-IOS-XR-iedge4710-cfg:subscriber-featurette"
    subscriberFeaturette.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberFeaturette.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberFeaturette.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberFeaturette.EntityData.Children = types.NewOrderedMap()
    subscriberFeaturette.EntityData.Children.Append("featurette-name", types.YChild{"FeaturetteName", nil})
    for i := range subscriberFeaturette.FeaturetteName {
        subscriberFeaturette.EntityData.Children.Append(types.GetSegmentPath(subscriberFeaturette.FeaturetteName[i]), types.YChild{"FeaturetteName", subscriberFeaturette.FeaturetteName[i]})
    }
    subscriberFeaturette.EntityData.Leafs = types.NewOrderedMap()

    subscriberFeaturette.EntityData.YListKeys = []string {}

    return &(subscriberFeaturette.EntityData)
}

// SubscriberFeaturette_FeaturetteName
// enable featurette processing
type SubscriberFeaturette_FeaturetteName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. subscriber feature. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Featurette interface{}

    // instance of featurette. The type is interface{} with range: 0..4294967295.
    Enable interface{}
}

func (featuretteName *SubscriberFeaturette_FeaturetteName) GetEntityData() *types.CommonEntityData {
    featuretteName.EntityData.YFilter = featuretteName.YFilter
    featuretteName.EntityData.YangName = "featurette-name"
    featuretteName.EntityData.BundleName = "cisco_ios_xr"
    featuretteName.EntityData.ParentYangName = "subscriber-featurette"
    featuretteName.EntityData.SegmentPath = "featurette-name" + types.AddKeyToken(featuretteName.Featurette, "featurette")
    featuretteName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    featuretteName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    featuretteName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    featuretteName.EntityData.Children = types.NewOrderedMap()
    featuretteName.EntityData.Leafs = types.NewOrderedMap()
    featuretteName.EntityData.Leafs.Append("featurette", types.YLeaf{"Featurette", featuretteName.Featurette})
    featuretteName.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", featuretteName.Enable})

    featuretteName.EntityData.YListKeys = []string {"Featurette"}

    return &(featuretteName.EntityData)
}

// IedgeLicenseManager
// iedge license manager
type IedgeLicenseManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session limit configured on linecard. The type is interface{} with range:
    // 1..200000.
    SessionLimit interface{}
}

func (iedgeLicenseManager *IedgeLicenseManager) GetEntityData() *types.CommonEntityData {
    iedgeLicenseManager.EntityData.YFilter = iedgeLicenseManager.YFilter
    iedgeLicenseManager.EntityData.YangName = "iedge-license-manager"
    iedgeLicenseManager.EntityData.BundleName = "cisco_ios_xr"
    iedgeLicenseManager.EntityData.ParentYangName = "Cisco-IOS-XR-iedge4710-cfg"
    iedgeLicenseManager.EntityData.SegmentPath = "Cisco-IOS-XR-iedge4710-cfg:iedge-license-manager"
    iedgeLicenseManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iedgeLicenseManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iedgeLicenseManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iedgeLicenseManager.EntityData.Children = types.NewOrderedMap()
    iedgeLicenseManager.EntityData.Leafs = types.NewOrderedMap()
    iedgeLicenseManager.EntityData.Leafs.Append("session-limit", types.YLeaf{"SessionLimit", iedgeLicenseManager.SessionLimit})

    iedgeLicenseManager.EntityData.YListKeys = []string {}

    return &(iedgeLicenseManager.EntityData)
}

// SubManager
// sub manager
type SubManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select location. The type is slice of SubManager_Location.
    Location []*SubManager_Location
}

func (subManager *SubManager) GetEntityData() *types.CommonEntityData {
    subManager.EntityData.YFilter = subManager.YFilter
    subManager.EntityData.YangName = "sub-manager"
    subManager.EntityData.BundleName = "cisco_ios_xr"
    subManager.EntityData.ParentYangName = "Cisco-IOS-XR-iedge4710-cfg"
    subManager.EntityData.SegmentPath = "Cisco-IOS-XR-iedge4710-cfg:sub-manager"
    subManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subManager.EntityData.Children = types.NewOrderedMap()
    subManager.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range subManager.Location {
        subManager.EntityData.Children.Append(types.GetSegmentPath(subManager.Location[i]), types.YChild{"Location", subManager.Location[i]})
    }
    subManager.EntityData.Leafs = types.NewOrderedMap()

    subManager.EntityData.YListKeys = []string {}

    return &(subManager.EntityData)
}

// SubManager_Location
// Select location
type SubManager_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify location. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Location1 interface{}

    // Disable history for subscriber manager. The type is interface{}.
    History interface{}

    // Subscriber manager trace.
    Trace SubManager_Location_Trace
}

func (location *SubManager_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "sub-manager"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.Location1, "location1")
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("trace", types.YChild{"Trace", &location.Trace})
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location1", types.YLeaf{"Location1", location.Location1})
    location.EntityData.Leafs.Append("history", types.YLeaf{"History", location.History})

    location.EntityData.YListKeys = []string {"Location1"}

    return &(location.EntityData)
}

// SubManager_Location_Trace
// Subscriber manager trace
type SubManager_Location_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber manager trace level. The type is interface{} with range:
    // 0..4294967295.
    TraceLevel interface{}
}

func (trace *SubManager_Location_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "location"
    trace.EntityData.SegmentPath = "trace"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("trace-level", types.YLeaf{"TraceLevel", trace.TraceLevel})

    trace.EntityData.YListKeys = []string {}

    return &(trace.EntityData)
}

