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
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // iEdge accounting feature.
    Accounting SubscriberManager_Accounting

    // SRG specific config.
    Srg SubscriberManager_Srg
}

func (subscriberManager *SubscriberManager) GetFilter() yfilter.YFilter { return subscriberManager.YFilter }

func (subscriberManager *SubscriberManager) SetFilter(yf yfilter.YFilter) { subscriberManager.YFilter = yf }

func (subscriberManager *SubscriberManager) GetGoName(yname string) string {
    if yname == "accounting" { return "Accounting" }
    if yname == "srg" { return "Srg" }
    return ""
}

func (subscriberManager *SubscriberManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-iedge4710-cfg:subscriber-manager"
}

func (subscriberManager *SubscriberManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting" {
        return &subscriberManager.Accounting
    }
    if childYangName == "srg" {
        return &subscriberManager.Srg
    }
    return nil
}

func (subscriberManager *SubscriberManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accounting"] = &subscriberManager.Accounting
    children["srg"] = &subscriberManager.Srg
    return children
}

func (subscriberManager *SubscriberManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberManager *SubscriberManager) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberManager *SubscriberManager) GetYangName() string { return "subscriber-manager" }

func (subscriberManager *SubscriberManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberManager *SubscriberManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberManager *SubscriberManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberManager *SubscriberManager) SetParent(parent types.Entity) { subscriberManager.parent = parent }

func (subscriberManager *SubscriberManager) GetParent() types.Entity { return subscriberManager.parent }

func (subscriberManager *SubscriberManager) GetParentYangName() string { return "Cisco-IOS-XR-iedge4710-cfg" }

// SubscriberManager_Accounting
// iEdge accounting feature
type SubscriberManager_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Accounting send stop feature.
    SendStop SubscriberManager_Accounting_SendStop

    // interim accounting related.
    Interim SubscriberManager_Accounting_Interim
}

func (accounting *SubscriberManager_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *SubscriberManager_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *SubscriberManager_Accounting) GetGoName(yname string) string {
    if yname == "send-stop" { return "SendStop" }
    if yname == "interim" { return "Interim" }
    return ""
}

func (accounting *SubscriberManager_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *SubscriberManager_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "send-stop" {
        return &accounting.SendStop
    }
    if childYangName == "interim" {
        return &accounting.Interim
    }
    return nil
}

func (accounting *SubscriberManager_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["send-stop"] = &accounting.SendStop
    children["interim"] = &accounting.Interim
    return children
}

func (accounting *SubscriberManager_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accounting *SubscriberManager_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *SubscriberManager_Accounting) GetYangName() string { return "accounting" }

func (accounting *SubscriberManager_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *SubscriberManager_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *SubscriberManager_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *SubscriberManager_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *SubscriberManager_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *SubscriberManager_Accounting) GetParentYangName() string { return "subscriber-manager" }

// SubscriberManager_Accounting_SendStop
// Accounting send stop feature
type SubscriberManager_Accounting_SendStop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Setup failure feature.
    SetupFailure SubscriberManager_Accounting_SendStop_SetupFailure
}

func (sendStop *SubscriberManager_Accounting_SendStop) GetFilter() yfilter.YFilter { return sendStop.YFilter }

func (sendStop *SubscriberManager_Accounting_SendStop) SetFilter(yf yfilter.YFilter) { sendStop.YFilter = yf }

func (sendStop *SubscriberManager_Accounting_SendStop) GetGoName(yname string) string {
    if yname == "setup-failure" { return "SetupFailure" }
    return ""
}

func (sendStop *SubscriberManager_Accounting_SendStop) GetSegmentPath() string {
    return "send-stop"
}

func (sendStop *SubscriberManager_Accounting_SendStop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "setup-failure" {
        return &sendStop.SetupFailure
    }
    return nil
}

func (sendStop *SubscriberManager_Accounting_SendStop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["setup-failure"] = &sendStop.SetupFailure
    return children
}

func (sendStop *SubscriberManager_Accounting_SendStop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sendStop *SubscriberManager_Accounting_SendStop) GetBundleName() string { return "cisco_ios_xr" }

func (sendStop *SubscriberManager_Accounting_SendStop) GetYangName() string { return "send-stop" }

func (sendStop *SubscriberManager_Accounting_SendStop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sendStop *SubscriberManager_Accounting_SendStop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sendStop *SubscriberManager_Accounting_SendStop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sendStop *SubscriberManager_Accounting_SendStop) SetParent(parent types.Entity) { sendStop.parent = parent }

func (sendStop *SubscriberManager_Accounting_SendStop) GetParent() types.Entity { return sendStop.parent }

func (sendStop *SubscriberManager_Accounting_SendStop) GetParentYangName() string { return "accounting" }

// SubscriberManager_Accounting_SendStop_SetupFailure
// Setup failure feature
type SubscriberManager_Accounting_SendStop_SetupFailure struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA List name either default or preconfigured. The type is string.
    MethodListName interface{}
}

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetFilter() yfilter.YFilter { return setupFailure.YFilter }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) SetFilter(yf yfilter.YFilter) { setupFailure.YFilter = yf }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetGoName(yname string) string {
    if yname == "method-list-name" { return "MethodListName" }
    return ""
}

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetSegmentPath() string {
    return "setup-failure"
}

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method-list-name"] = setupFailure.MethodListName
    return leafs
}

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetBundleName() string { return "cisco_ios_xr" }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetYangName() string { return "setup-failure" }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) SetParent(parent types.Entity) { setupFailure.parent = parent }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetParent() types.Entity { return setupFailure.parent }

func (setupFailure *SubscriberManager_Accounting_SendStop_SetupFailure) GetParentYangName() string { return "send-stop" }

// SubscriberManager_Accounting_Interim
// interim accounting related
type SubscriberManager_Accounting_Interim struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // variation of first session or service interim record from configured
    // timeout.
    Variation SubscriberManager_Accounting_Interim_Variation
}

func (interim *SubscriberManager_Accounting_Interim) GetFilter() yfilter.YFilter { return interim.YFilter }

func (interim *SubscriberManager_Accounting_Interim) SetFilter(yf yfilter.YFilter) { interim.YFilter = yf }

func (interim *SubscriberManager_Accounting_Interim) GetGoName(yname string) string {
    if yname == "variation" { return "Variation" }
    return ""
}

func (interim *SubscriberManager_Accounting_Interim) GetSegmentPath() string {
    return "interim"
}

func (interim *SubscriberManager_Accounting_Interim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "variation" {
        return &interim.Variation
    }
    return nil
}

func (interim *SubscriberManager_Accounting_Interim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["variation"] = &interim.Variation
    return children
}

func (interim *SubscriberManager_Accounting_Interim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interim *SubscriberManager_Accounting_Interim) GetBundleName() string { return "cisco_ios_xr" }

func (interim *SubscriberManager_Accounting_Interim) GetYangName() string { return "interim" }

func (interim *SubscriberManager_Accounting_Interim) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interim *SubscriberManager_Accounting_Interim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interim *SubscriberManager_Accounting_Interim) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interim *SubscriberManager_Accounting_Interim) SetParent(parent types.Entity) { interim.parent = parent }

func (interim *SubscriberManager_Accounting_Interim) GetParent() types.Entity { return interim.parent }

func (interim *SubscriberManager_Accounting_Interim) GetParentYangName() string { return "accounting" }

// SubscriberManager_Accounting_Interim_Variation
// variation of first session or service interim
// record from configured timeout
type SubscriberManager_Accounting_Interim_Variation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // maximum percentage variation (maximum absolute variation is 15 minutes).
    // The type is interface{} with range: 0..50. Units are percentage.
    MaximumPercentageVariation interface{}
}

func (variation *SubscriberManager_Accounting_Interim_Variation) GetFilter() yfilter.YFilter { return variation.YFilter }

func (variation *SubscriberManager_Accounting_Interim_Variation) SetFilter(yf yfilter.YFilter) { variation.YFilter = yf }

func (variation *SubscriberManager_Accounting_Interim_Variation) GetGoName(yname string) string {
    if yname == "maximum-percentage-variation" { return "MaximumPercentageVariation" }
    return ""
}

func (variation *SubscriberManager_Accounting_Interim_Variation) GetSegmentPath() string {
    return "variation"
}

func (variation *SubscriberManager_Accounting_Interim_Variation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (variation *SubscriberManager_Accounting_Interim_Variation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (variation *SubscriberManager_Accounting_Interim_Variation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-percentage-variation"] = variation.MaximumPercentageVariation
    return leafs
}

func (variation *SubscriberManager_Accounting_Interim_Variation) GetBundleName() string { return "cisco_ios_xr" }

func (variation *SubscriberManager_Accounting_Interim_Variation) GetYangName() string { return "variation" }

func (variation *SubscriberManager_Accounting_Interim_Variation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (variation *SubscriberManager_Accounting_Interim_Variation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (variation *SubscriberManager_Accounting_Interim_Variation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (variation *SubscriberManager_Accounting_Interim_Variation) SetParent(parent types.Entity) { variation.parent = parent }

func (variation *SubscriberManager_Accounting_Interim_Variation) GetParent() types.Entity { return variation.parent }

func (variation *SubscriberManager_Accounting_Interim_Variation) GetParentYangName() string { return "interim" }

// SubscriberManager_Srg
// SRG specific config
type SubscriberManager_Srg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // sync account session id from master to slave. The type is interface{}.
    SyncAccountSessionId interface{}
}

func (srg *SubscriberManager_Srg) GetFilter() yfilter.YFilter { return srg.YFilter }

func (srg *SubscriberManager_Srg) SetFilter(yf yfilter.YFilter) { srg.YFilter = yf }

func (srg *SubscriberManager_Srg) GetGoName(yname string) string {
    if yname == "sync-account-session-id" { return "SyncAccountSessionId" }
    return ""
}

func (srg *SubscriberManager_Srg) GetSegmentPath() string {
    return "srg"
}

func (srg *SubscriberManager_Srg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srg *SubscriberManager_Srg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srg *SubscriberManager_Srg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sync-account-session-id"] = srg.SyncAccountSessionId
    return leafs
}

func (srg *SubscriberManager_Srg) GetBundleName() string { return "cisco_ios_xr" }

func (srg *SubscriberManager_Srg) GetYangName() string { return "srg" }

func (srg *SubscriberManager_Srg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srg *SubscriberManager_Srg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srg *SubscriberManager_Srg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srg *SubscriberManager_Srg) SetParent(parent types.Entity) { srg.parent = parent }

func (srg *SubscriberManager_Srg) GetParent() types.Entity { return srg.parent }

func (srg *SubscriberManager_Srg) GetParentYangName() string { return "subscriber-manager" }

// SubscriberFeaturette
// subscriber featurette
type SubscriberFeaturette struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // enable identity change processing. The type is slice of
    // SubscriberFeaturette_IdentityChange.
    IdentityChange []SubscriberFeaturette_IdentityChange
}

func (subscriberFeaturette *SubscriberFeaturette) GetFilter() yfilter.YFilter { return subscriberFeaturette.YFilter }

func (subscriberFeaturette *SubscriberFeaturette) SetFilter(yf yfilter.YFilter) { subscriberFeaturette.YFilter = yf }

func (subscriberFeaturette *SubscriberFeaturette) GetGoName(yname string) string {
    if yname == "identity-change" { return "IdentityChange" }
    return ""
}

func (subscriberFeaturette *SubscriberFeaturette) GetSegmentPath() string {
    return "Cisco-IOS-XR-iedge4710-cfg:subscriber-featurette"
}

func (subscriberFeaturette *SubscriberFeaturette) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "identity-change" {
        for _, c := range subscriberFeaturette.IdentityChange {
            if subscriberFeaturette.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubscriberFeaturette_IdentityChange{}
        subscriberFeaturette.IdentityChange = append(subscriberFeaturette.IdentityChange, child)
        return &subscriberFeaturette.IdentityChange[len(subscriberFeaturette.IdentityChange)-1]
    }
    return nil
}

func (subscriberFeaturette *SubscriberFeaturette) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subscriberFeaturette.IdentityChange {
        children[subscriberFeaturette.IdentityChange[i].GetSegmentPath()] = &subscriberFeaturette.IdentityChange[i]
    }
    return children
}

func (subscriberFeaturette *SubscriberFeaturette) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberFeaturette *SubscriberFeaturette) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberFeaturette *SubscriberFeaturette) GetYangName() string { return "subscriber-featurette" }

func (subscriberFeaturette *SubscriberFeaturette) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberFeaturette *SubscriberFeaturette) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberFeaturette *SubscriberFeaturette) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberFeaturette *SubscriberFeaturette) SetParent(parent types.Entity) { subscriberFeaturette.parent = parent }

func (subscriberFeaturette *SubscriberFeaturette) GetParent() types.Entity { return subscriberFeaturette.parent }

func (subscriberFeaturette *SubscriberFeaturette) GetParentYangName() string { return "Cisco-IOS-XR-iedge4710-cfg" }

// SubscriberFeaturette_IdentityChange
// enable identity change processing
type SubscriberFeaturette_IdentityChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. identity change. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    IdentityChange interface{}

    // instance of identity-change. The type is interface{} with range:
    // -2147483648..2147483647.
    Enable interface{}
}

func (identityChange *SubscriberFeaturette_IdentityChange) GetFilter() yfilter.YFilter { return identityChange.YFilter }

func (identityChange *SubscriberFeaturette_IdentityChange) SetFilter(yf yfilter.YFilter) { identityChange.YFilter = yf }

func (identityChange *SubscriberFeaturette_IdentityChange) GetGoName(yname string) string {
    if yname == "identity-change" { return "IdentityChange" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (identityChange *SubscriberFeaturette_IdentityChange) GetSegmentPath() string {
    return "identity-change" + "[identity-change='" + fmt.Sprintf("%v", identityChange.IdentityChange) + "']"
}

func (identityChange *SubscriberFeaturette_IdentityChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (identityChange *SubscriberFeaturette_IdentityChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (identityChange *SubscriberFeaturette_IdentityChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["identity-change"] = identityChange.IdentityChange
    leafs["enable"] = identityChange.Enable
    return leafs
}

func (identityChange *SubscriberFeaturette_IdentityChange) GetBundleName() string { return "cisco_ios_xr" }

func (identityChange *SubscriberFeaturette_IdentityChange) GetYangName() string { return "identity-change" }

func (identityChange *SubscriberFeaturette_IdentityChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (identityChange *SubscriberFeaturette_IdentityChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (identityChange *SubscriberFeaturette_IdentityChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (identityChange *SubscriberFeaturette_IdentityChange) SetParent(parent types.Entity) { identityChange.parent = parent }

func (identityChange *SubscriberFeaturette_IdentityChange) GetParent() types.Entity { return identityChange.parent }

func (identityChange *SubscriberFeaturette_IdentityChange) GetParentYangName() string { return "subscriber-featurette" }

// IedgeLicenseManager
// iedge license manager
type IedgeLicenseManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Location. For eg., 0/1/CPU0. The type is slice of IedgeLicenseManager_Node.
    Node []IedgeLicenseManager_Node
}

func (iedgeLicenseManager *IedgeLicenseManager) GetFilter() yfilter.YFilter { return iedgeLicenseManager.YFilter }

func (iedgeLicenseManager *IedgeLicenseManager) SetFilter(yf yfilter.YFilter) { iedgeLicenseManager.YFilter = yf }

func (iedgeLicenseManager *IedgeLicenseManager) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (iedgeLicenseManager *IedgeLicenseManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-iedge4710-cfg:iedge-license-manager"
}

func (iedgeLicenseManager *IedgeLicenseManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range iedgeLicenseManager.Node {
            if iedgeLicenseManager.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IedgeLicenseManager_Node{}
        iedgeLicenseManager.Node = append(iedgeLicenseManager.Node, child)
        return &iedgeLicenseManager.Node[len(iedgeLicenseManager.Node)-1]
    }
    return nil
}

func (iedgeLicenseManager *IedgeLicenseManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range iedgeLicenseManager.Node {
        children[iedgeLicenseManager.Node[i].GetSegmentPath()] = &iedgeLicenseManager.Node[i]
    }
    return children
}

func (iedgeLicenseManager *IedgeLicenseManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iedgeLicenseManager *IedgeLicenseManager) GetBundleName() string { return "cisco_ios_xr" }

func (iedgeLicenseManager *IedgeLicenseManager) GetYangName() string { return "iedge-license-manager" }

func (iedgeLicenseManager *IedgeLicenseManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (iedgeLicenseManager *IedgeLicenseManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (iedgeLicenseManager *IedgeLicenseManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (iedgeLicenseManager *IedgeLicenseManager) SetParent(parent types.Entity) { iedgeLicenseManager.parent = parent }

func (iedgeLicenseManager *IedgeLicenseManager) GetParent() types.Entity { return iedgeLicenseManager.parent }

func (iedgeLicenseManager *IedgeLicenseManager) GetParentYangName() string { return "Cisco-IOS-XR-iedge4710-cfg" }

// IedgeLicenseManager_Node
// Location. For eg., 0/1/CPU0
type IedgeLicenseManager_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node id to filter on. For eg., 0/1/CPU0. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    NodeName interface{}

    // Session limit configured on linecard. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionLimit interface{}

    // Session threshold configured on linecard. The type is interface{} with
    // range: -2147483648..2147483647.
    SessionThreshold interface{}
}

func (node *IedgeLicenseManager_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *IedgeLicenseManager_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *IedgeLicenseManager_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "session-limit" { return "SessionLimit" }
    if yname == "session-threshold" { return "SessionThreshold" }
    return ""
}

func (node *IedgeLicenseManager_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *IedgeLicenseManager_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *IedgeLicenseManager_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *IedgeLicenseManager_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["session-limit"] = node.SessionLimit
    leafs["session-threshold"] = node.SessionThreshold
    return leafs
}

func (node *IedgeLicenseManager_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *IedgeLicenseManager_Node) GetYangName() string { return "node" }

func (node *IedgeLicenseManager_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *IedgeLicenseManager_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *IedgeLicenseManager_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *IedgeLicenseManager_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *IedgeLicenseManager_Node) GetParent() types.Entity { return node.parent }

func (node *IedgeLicenseManager_Node) GetParentYangName() string { return "iedge-license-manager" }

// SubManager
// sub manager
type SubManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Select location. The type is slice of SubManager_Location.
    Location []SubManager_Location
}

func (subManager *SubManager) GetFilter() yfilter.YFilter { return subManager.YFilter }

func (subManager *SubManager) SetFilter(yf yfilter.YFilter) { subManager.YFilter = yf }

func (subManager *SubManager) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (subManager *SubManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-iedge4710-cfg:sub-manager"
}

func (subManager *SubManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "location" {
        for _, c := range subManager.Location {
            if subManager.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SubManager_Location{}
        subManager.Location = append(subManager.Location, child)
        return &subManager.Location[len(subManager.Location)-1]
    }
    return nil
}

func (subManager *SubManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subManager.Location {
        children[subManager.Location[i].GetSegmentPath()] = &subManager.Location[i]
    }
    return children
}

func (subManager *SubManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subManager *SubManager) GetBundleName() string { return "cisco_ios_xr" }

func (subManager *SubManager) GetYangName() string { return "sub-manager" }

func (subManager *SubManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subManager *SubManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subManager *SubManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subManager *SubManager) SetParent(parent types.Entity) { subManager.parent = parent }

func (subManager *SubManager) GetParent() types.Entity { return subManager.parent }

func (subManager *SubManager) GetParentYangName() string { return "Cisco-IOS-XR-iedge4710-cfg" }

// SubManager_Location
// Select location
type SubManager_Location struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify location. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Location1 interface{}

    // Disable history for subscriber manager. The type is interface{}.
    History interface{}

    // Subscriber manager trace.
    Trace SubManager_Location_Trace
}

func (location *SubManager_Location) GetFilter() yfilter.YFilter { return location.YFilter }

func (location *SubManager_Location) SetFilter(yf yfilter.YFilter) { location.YFilter = yf }

func (location *SubManager_Location) GetGoName(yname string) string {
    if yname == "location1" { return "Location1" }
    if yname == "history" { return "History" }
    if yname == "trace" { return "Trace" }
    return ""
}

func (location *SubManager_Location) GetSegmentPath() string {
    return "location" + "[location1='" + fmt.Sprintf("%v", location.Location1) + "']"
}

func (location *SubManager_Location) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trace" {
        return &location.Trace
    }
    return nil
}

func (location *SubManager_Location) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["trace"] = &location.Trace
    return children
}

func (location *SubManager_Location) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location1"] = location.Location1
    leafs["history"] = location.History
    return leafs
}

func (location *SubManager_Location) GetBundleName() string { return "cisco_ios_xr" }

func (location *SubManager_Location) GetYangName() string { return "location" }

func (location *SubManager_Location) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (location *SubManager_Location) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (location *SubManager_Location) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (location *SubManager_Location) SetParent(parent types.Entity) { location.parent = parent }

func (location *SubManager_Location) GetParent() types.Entity { return location.parent }

func (location *SubManager_Location) GetParentYangName() string { return "sub-manager" }

// SubManager_Location_Trace
// Subscriber manager trace
type SubManager_Location_Trace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber manager trace level. The type is interface{} with range:
    // -2147483648..2147483647.
    TraceLevel interface{}
}

func (trace *SubManager_Location_Trace) GetFilter() yfilter.YFilter { return trace.YFilter }

func (trace *SubManager_Location_Trace) SetFilter(yf yfilter.YFilter) { trace.YFilter = yf }

func (trace *SubManager_Location_Trace) GetGoName(yname string) string {
    if yname == "trace-level" { return "TraceLevel" }
    return ""
}

func (trace *SubManager_Location_Trace) GetSegmentPath() string {
    return "trace"
}

func (trace *SubManager_Location_Trace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trace *SubManager_Location_Trace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trace *SubManager_Location_Trace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trace-level"] = trace.TraceLevel
    return leafs
}

func (trace *SubManager_Location_Trace) GetBundleName() string { return "cisco_ios_xr" }

func (trace *SubManager_Location_Trace) GetYangName() string { return "trace" }

func (trace *SubManager_Location_Trace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trace *SubManager_Location_Trace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trace *SubManager_Location_Trace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trace *SubManager_Location_Trace) SetParent(parent types.Entity) { trace.parent = parent }

func (trace *SubManager_Location_Trace) GetParent() types.Entity { return trace.parent }

func (trace *SubManager_Location_Trace) GetParentYangName() string { return "location" }

