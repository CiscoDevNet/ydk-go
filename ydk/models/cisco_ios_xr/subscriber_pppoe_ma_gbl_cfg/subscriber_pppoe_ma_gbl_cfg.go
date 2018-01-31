// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-pppoe-ma-gbl package configuration.
// 
// This module contains definitions
// for the following management objects:
//   pppoe-cfg: PPPOE configuration data
// 
// This YANG module augments the
//   Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_pppoe_ma_gbl_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_pppoe_ma_gbl_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg pppoe-cfg}", reflect.TypeOf(PppoeCfg{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-cfg", reflect.TypeOf(PppoeCfg{}))
}

// PppoeInvalidSessionIdBehavior represents Pppoe invalid session id behavior
type PppoeInvalidSessionIdBehavior string

const (
    // Drop packets with an invalid session ID
    PppoeInvalidSessionIdBehavior_drop PppoeInvalidSessionIdBehavior = "drop"

    // Log packets with an invalid session ID
    PppoeInvalidSessionIdBehavior_log PppoeInvalidSessionIdBehavior = "log"
)

// PppoeCfg
// PPPOE configuration data
type PppoeCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable per-parent session ID spaces. The type is interface{}.
    SessionIdSpaceFlat interface{}

    // Set the PPPoE in-flight window size. The type is interface{} with range:
    // 1..20000.
    InFlightWindow interface{}

    // PPPoE BBA-Group configuration data.
    PppoeBbaGroups PppoeCfg_PppoeBbaGroups
}

func (pppoeCfg *PppoeCfg) GetFilter() yfilter.YFilter { return pppoeCfg.YFilter }

func (pppoeCfg *PppoeCfg) SetFilter(yf yfilter.YFilter) { pppoeCfg.YFilter = yf }

func (pppoeCfg *PppoeCfg) GetGoName(yname string) string {
    if yname == "session-id-space-flat" { return "SessionIdSpaceFlat" }
    if yname == "in-flight-window" { return "InFlightWindow" }
    if yname == "pppoe-bba-groups" { return "PppoeBbaGroups" }
    return ""
}

func (pppoeCfg *PppoeCfg) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-cfg"
}

func (pppoeCfg *PppoeCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pppoe-bba-groups" {
        return &pppoeCfg.PppoeBbaGroups
    }
    return nil
}

func (pppoeCfg *PppoeCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pppoe-bba-groups"] = &pppoeCfg.PppoeBbaGroups
    return children
}

func (pppoeCfg *PppoeCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id-space-flat"] = pppoeCfg.SessionIdSpaceFlat
    leafs["in-flight-window"] = pppoeCfg.InFlightWindow
    return leafs
}

func (pppoeCfg *PppoeCfg) GetBundleName() string { return "cisco_ios_xr" }

func (pppoeCfg *PppoeCfg) GetYangName() string { return "pppoe-cfg" }

func (pppoeCfg *PppoeCfg) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppoeCfg *PppoeCfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppoeCfg *PppoeCfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppoeCfg *PppoeCfg) SetParent(parent types.Entity) { pppoeCfg.parent = parent }

func (pppoeCfg *PppoeCfg) GetParent() types.Entity { return pppoeCfg.parent }

func (pppoeCfg *PppoeCfg) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg" }

// PppoeCfg_PppoeBbaGroups
// PPPoE BBA-Group configuration data
type PppoeCfg_PppoeBbaGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPPoE BBA-Group configuration data. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup.
    PppoeBbaGroup []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup
}

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetFilter() yfilter.YFilter { return pppoeBbaGroups.YFilter }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) SetFilter(yf yfilter.YFilter) { pppoeBbaGroups.YFilter = yf }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetGoName(yname string) string {
    if yname == "pppoe-bba-group" { return "PppoeBbaGroup" }
    return ""
}

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetSegmentPath() string {
    return "pppoe-bba-groups"
}

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pppoe-bba-group" {
        for _, c := range pppoeBbaGroups.PppoeBbaGroup {
            if pppoeBbaGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeCfg_PppoeBbaGroups_PppoeBbaGroup{}
        pppoeBbaGroups.PppoeBbaGroup = append(pppoeBbaGroups.PppoeBbaGroup, child)
        return &pppoeBbaGroups.PppoeBbaGroup[len(pppoeBbaGroups.PppoeBbaGroup)-1]
    }
    return nil
}

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pppoeBbaGroups.PppoeBbaGroup {
        children[pppoeBbaGroups.PppoeBbaGroup[i].GetSegmentPath()] = &pppoeBbaGroups.PppoeBbaGroup[i]
    }
    return children
}

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetBundleName() string { return "cisco_ios_xr" }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetYangName() string { return "pppoe-bba-groups" }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) SetParent(parent types.Entity) { pppoeBbaGroups.parent = parent }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetParent() types.Entity { return pppoeBbaGroups.parent }

func (pppoeBbaGroups *PppoeCfg_PppoeBbaGroups) GetParentYangName() string { return "pppoe-cfg" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup
// PPPoE BBA-Group configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. BBA-Group name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    BbaGroup interface{}

    // PPPoE session completion timeout. The type is interface{} with range:
    // 30..600.
    CompletionTimeout interface{}

    // Invalid session-ID behavior. The type is PppoeInvalidSessionIdBehavior.
    InvalidSessionId interface{}

    // Enable padt after shutdown. The type is interface{}.
    EnablePadtAfterShutDown interface{}

    // PPPoE default MTU. The type is interface{} with range: 500..2000.
    Mtu interface{}

    // PPPoE tag configuration data.
    Tag PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag

    // PPPoE session configuration data.
    Sessions PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions

    // PPPoE control-packet configuration data.
    ControlPackets PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets

    // PPPoE PADO delay configuration data.
    PaDoDelay PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay
}

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetFilter() yfilter.YFilter { return pppoeBbaGroup.YFilter }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) SetFilter(yf yfilter.YFilter) { pppoeBbaGroup.YFilter = yf }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetGoName(yname string) string {
    if yname == "bba-group" { return "BbaGroup" }
    if yname == "completion-timeout" { return "CompletionTimeout" }
    if yname == "invalid-session-id" { return "InvalidSessionId" }
    if yname == "enable-padt-after-shut-down" { return "EnablePadtAfterShutDown" }
    if yname == "mtu" { return "Mtu" }
    if yname == "tag" { return "Tag" }
    if yname == "sessions" { return "Sessions" }
    if yname == "control-packets" { return "ControlPackets" }
    if yname == "pa-do-delay" { return "PaDoDelay" }
    return ""
}

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetSegmentPath() string {
    return "pppoe-bba-group" + "[bba-group='" + fmt.Sprintf("%v", pppoeBbaGroup.BbaGroup) + "']"
}

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tag" {
        return &pppoeBbaGroup.Tag
    }
    if childYangName == "sessions" {
        return &pppoeBbaGroup.Sessions
    }
    if childYangName == "control-packets" {
        return &pppoeBbaGroup.ControlPackets
    }
    if childYangName == "pa-do-delay" {
        return &pppoeBbaGroup.PaDoDelay
    }
    return nil
}

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tag"] = &pppoeBbaGroup.Tag
    children["sessions"] = &pppoeBbaGroup.Sessions
    children["control-packets"] = &pppoeBbaGroup.ControlPackets
    children["pa-do-delay"] = &pppoeBbaGroup.PaDoDelay
    return children
}

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bba-group"] = pppoeBbaGroup.BbaGroup
    leafs["completion-timeout"] = pppoeBbaGroup.CompletionTimeout
    leafs["invalid-session-id"] = pppoeBbaGroup.InvalidSessionId
    leafs["enable-padt-after-shut-down"] = pppoeBbaGroup.EnablePadtAfterShutDown
    leafs["mtu"] = pppoeBbaGroup.Mtu
    return leafs
}

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetBundleName() string { return "cisco_ios_xr" }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetYangName() string { return "pppoe-bba-group" }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) SetParent(parent types.Entity) { pppoeBbaGroup.parent = parent }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetParent() types.Entity { return pppoeBbaGroup.parent }

func (pppoeBbaGroup *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup) GetParentYangName() string { return "pppoe-bba-groups" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag
// PPPoE tag configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ignore the ppp-max-payload tag. The type is interface{}.
    PppMaxPayloadDeny interface{}

    // Disable advertising of unrequested service names. The type is interface{}.
    ServiceSelectionDisable interface{}

    // The name to include in the AC tag. The type is string.
    AcName interface{}

    // PADR packets.
    Padr PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr

    // Service name.
    ServiceNameConfigureds PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds

    // Minimum and maximum payloads.
    PppMaxPayload PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload
}

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetFilter() yfilter.YFilter { return tag.YFilter }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) SetFilter(yf yfilter.YFilter) { tag.YFilter = yf }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetGoName(yname string) string {
    if yname == "ppp-max-payload-deny" { return "PppMaxPayloadDeny" }
    if yname == "service-selection-disable" { return "ServiceSelectionDisable" }
    if yname == "ac-name" { return "AcName" }
    if yname == "padr" { return "Padr" }
    if yname == "service-name-configureds" { return "ServiceNameConfigureds" }
    if yname == "ppp-max-payload" { return "PppMaxPayload" }
    return ""
}

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetSegmentPath() string {
    return "tag"
}

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "padr" {
        return &tag.Padr
    }
    if childYangName == "service-name-configureds" {
        return &tag.ServiceNameConfigureds
    }
    if childYangName == "ppp-max-payload" {
        return &tag.PppMaxPayload
    }
    return nil
}

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["padr"] = &tag.Padr
    children["service-name-configureds"] = &tag.ServiceNameConfigureds
    children["ppp-max-payload"] = &tag.PppMaxPayload
    return children
}

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ppp-max-payload-deny"] = tag.PppMaxPayloadDeny
    leafs["service-selection-disable"] = tag.ServiceSelectionDisable
    leafs["ac-name"] = tag.AcName
    return leafs
}

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetBundleName() string { return "cisco_ios_xr" }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetYangName() string { return "tag" }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) SetParent(parent types.Entity) { tag.parent = parent }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetParent() types.Entity { return tag.parent }

func (tag *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag) GetParentYangName() string { return "pppoe-bba-group" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr
// PADR packets
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allow sessions to come up with unique relay-session-id in padr. The type is
    // interface{}.
    SessionUniqueRelaySessionId interface{}

    // Allow sessions to come up with invalid-payload. The type is interface{}.
    InvalidPayloadAllow interface{}
}

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetFilter() yfilter.YFilter { return padr.YFilter }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) SetFilter(yf yfilter.YFilter) { padr.YFilter = yf }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetGoName(yname string) string {
    if yname == "session-unique-relay-session-id" { return "SessionUniqueRelaySessionId" }
    if yname == "invalid-payload-allow" { return "InvalidPayloadAllow" }
    return ""
}

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetSegmentPath() string {
    return "padr"
}

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-unique-relay-session-id"] = padr.SessionUniqueRelaySessionId
    leafs["invalid-payload-allow"] = padr.InvalidPayloadAllow
    return leafs
}

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetBundleName() string { return "cisco_ios_xr" }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetYangName() string { return "padr" }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) SetParent(parent types.Entity) { padr.parent = parent }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetParent() types.Entity { return padr.parent }

func (padr *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_Padr) GetParentYangName() string { return "tag" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds
// Service name
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service name supported on this group. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured.
    ServiceNameConfigured []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured
}

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetFilter() yfilter.YFilter { return serviceNameConfigureds.YFilter }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) SetFilter(yf yfilter.YFilter) { serviceNameConfigureds.YFilter = yf }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetGoName(yname string) string {
    if yname == "service-name-configured" { return "ServiceNameConfigured" }
    return ""
}

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetSegmentPath() string {
    return "service-name-configureds"
}

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-name-configured" {
        for _, c := range serviceNameConfigureds.ServiceNameConfigured {
            if serviceNameConfigureds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured{}
        serviceNameConfigureds.ServiceNameConfigured = append(serviceNameConfigureds.ServiceNameConfigured, child)
        return &serviceNameConfigureds.ServiceNameConfigured[len(serviceNameConfigureds.ServiceNameConfigured)-1]
    }
    return nil
}

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serviceNameConfigureds.ServiceNameConfigured {
        children[serviceNameConfigureds.ServiceNameConfigured[i].GetSegmentPath()] = &serviceNameConfigureds.ServiceNameConfigured[i]
    }
    return children
}

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetBundleName() string { return "cisco_ios_xr" }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetYangName() string { return "service-name-configureds" }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) SetParent(parent types.Entity) { serviceNameConfigureds.parent = parent }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetParent() types.Entity { return serviceNameConfigureds.parent }

func (serviceNameConfigureds *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds) GetParentYangName() string { return "tag" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured
// Service name supported on this group
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Service name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetFilter() yfilter.YFilter { return serviceNameConfigured.YFilter }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) SetFilter(yf yfilter.YFilter) { serviceNameConfigured.YFilter = yf }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetSegmentPath() string {
    return "service-name-configured" + "[name='" + fmt.Sprintf("%v", serviceNameConfigured.Name) + "']"
}

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = serviceNameConfigured.Name
    return leafs
}

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetBundleName() string { return "cisco_ios_xr" }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetYangName() string { return "service-name-configured" }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) SetParent(parent types.Entity) { serviceNameConfigured.parent = parent }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetParent() types.Entity { return serviceNameConfigured.parent }

func (serviceNameConfigured *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_ServiceNameConfigureds_ServiceNameConfigured) GetParentYangName() string { return "service-name-configureds" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload
// Minimum and maximum payloads
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum payload. The type is interface{} with range: 500..2000. This
    // attribute is mandatory.
    Min interface{}

    // Maximum payload. The type is interface{} with range: 500..2000. This
    // attribute is mandatory.
    Max interface{}
}

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetFilter() yfilter.YFilter { return pppMaxPayload.YFilter }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) SetFilter(yf yfilter.YFilter) { pppMaxPayload.YFilter = yf }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "max" { return "Max" }
    return ""
}

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetSegmentPath() string {
    return "ppp-max-payload"
}

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = pppMaxPayload.Min
    leafs["max"] = pppMaxPayload.Max
    return leafs
}

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetBundleName() string { return "cisco_ios_xr" }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetYangName() string { return "ppp-max-payload" }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) SetParent(parent types.Entity) { pppMaxPayload.parent = parent }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetParent() types.Entity { return pppMaxPayload.parent }

func (pppMaxPayload *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Tag_PppMaxPayload) GetParentYangName() string { return "tag" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions
// PPPoE session configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set VLAN (inner + outer tags) session throttle.
    VlanThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle

    // Set Inner VLAN session throttle.
    InnerVlanThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle

    // Set Remote ID session limit and threshold.
    RemoteIdLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit

    // Set per-MAC/Access interface throttle for IWF sessions.
    MacIwfAccessInterfaceThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle

    // Set per-access interface limit.
    AccessInterfaceLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit

    // Set per-MAC/Access Interface throttle.
    MacAccessInterfaceThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle

    // Set Outer VLAN session limit and threshold.
    OuterVlanLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit

    // Set Circuit ID session throttle.
    CircuitIdThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle

    // Set per-MAC address session limit and threshold.
    MacLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit

    // Set Circuit ID session limit and threshold.
    CircuitIdLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit

    // Set per-MAC session limit and threshold for IWF sessions.
    MacIwfLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit

    // Set per-MAC access interface session limit and threshold for IWF sessions.
    MacIwfAccessInterfaceLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit

    // Set Inner VLAN session limit and threshold.
    InnerVlanLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit

    // Set Outer VLAN session throttle.
    OuterVlanThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle

    // Set per-MAC throttle.
    MacThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle

    // Set Circuit ID and Remote ID session limit/threshold.
    CircuitIdAndRemoteIdLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit

    // Set VLAN (inner + outer tags) session limit and threshold.
    VlanLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit

    // Set per-MAC access interface session limit and threshold.
    MacAccessInterfaceLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit

    // Set Remote ID session throttle.
    RemoteIdThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle

    // Set per-card session limit and threshold.
    MaxLimit PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit

    // Set Circuit ID and Remote ID session throttle.
    CircuitIdAndRemoteIdThrottle PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle
}

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetGoName(yname string) string {
    if yname == "vlan-throttle" { return "VlanThrottle" }
    if yname == "inner-vlan-throttle" { return "InnerVlanThrottle" }
    if yname == "remote-id-limit" { return "RemoteIdLimit" }
    if yname == "mac-iwf-access-interface-throttle" { return "MacIwfAccessInterfaceThrottle" }
    if yname == "access-interface-limit" { return "AccessInterfaceLimit" }
    if yname == "mac-access-interface-throttle" { return "MacAccessInterfaceThrottle" }
    if yname == "outer-vlan-limit" { return "OuterVlanLimit" }
    if yname == "circuit-id-throttle" { return "CircuitIdThrottle" }
    if yname == "mac-limit" { return "MacLimit" }
    if yname == "circuit-id-limit" { return "CircuitIdLimit" }
    if yname == "mac-iwf-limit" { return "MacIwfLimit" }
    if yname == "mac-iwf-access-interface-limit" { return "MacIwfAccessInterfaceLimit" }
    if yname == "inner-vlan-limit" { return "InnerVlanLimit" }
    if yname == "outer-vlan-throttle" { return "OuterVlanThrottle" }
    if yname == "mac-throttle" { return "MacThrottle" }
    if yname == "circuit-id-and-remote-id-limit" { return "CircuitIdAndRemoteIdLimit" }
    if yname == "vlan-limit" { return "VlanLimit" }
    if yname == "mac-access-interface-limit" { return "MacAccessInterfaceLimit" }
    if yname == "remote-id-throttle" { return "RemoteIdThrottle" }
    if yname == "max-limit" { return "MaxLimit" }
    if yname == "circuit-id-and-remote-id-throttle" { return "CircuitIdAndRemoteIdThrottle" }
    return ""
}

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vlan-throttle" {
        return &sessions.VlanThrottle
    }
    if childYangName == "inner-vlan-throttle" {
        return &sessions.InnerVlanThrottle
    }
    if childYangName == "remote-id-limit" {
        return &sessions.RemoteIdLimit
    }
    if childYangName == "mac-iwf-access-interface-throttle" {
        return &sessions.MacIwfAccessInterfaceThrottle
    }
    if childYangName == "access-interface-limit" {
        return &sessions.AccessInterfaceLimit
    }
    if childYangName == "mac-access-interface-throttle" {
        return &sessions.MacAccessInterfaceThrottle
    }
    if childYangName == "outer-vlan-limit" {
        return &sessions.OuterVlanLimit
    }
    if childYangName == "circuit-id-throttle" {
        return &sessions.CircuitIdThrottle
    }
    if childYangName == "mac-limit" {
        return &sessions.MacLimit
    }
    if childYangName == "circuit-id-limit" {
        return &sessions.CircuitIdLimit
    }
    if childYangName == "mac-iwf-limit" {
        return &sessions.MacIwfLimit
    }
    if childYangName == "mac-iwf-access-interface-limit" {
        return &sessions.MacIwfAccessInterfaceLimit
    }
    if childYangName == "inner-vlan-limit" {
        return &sessions.InnerVlanLimit
    }
    if childYangName == "outer-vlan-throttle" {
        return &sessions.OuterVlanThrottle
    }
    if childYangName == "mac-throttle" {
        return &sessions.MacThrottle
    }
    if childYangName == "circuit-id-and-remote-id-limit" {
        return &sessions.CircuitIdAndRemoteIdLimit
    }
    if childYangName == "vlan-limit" {
        return &sessions.VlanLimit
    }
    if childYangName == "mac-access-interface-limit" {
        return &sessions.MacAccessInterfaceLimit
    }
    if childYangName == "remote-id-throttle" {
        return &sessions.RemoteIdThrottle
    }
    if childYangName == "max-limit" {
        return &sessions.MaxLimit
    }
    if childYangName == "circuit-id-and-remote-id-throttle" {
        return &sessions.CircuitIdAndRemoteIdThrottle
    }
    return nil
}

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vlan-throttle"] = &sessions.VlanThrottle
    children["inner-vlan-throttle"] = &sessions.InnerVlanThrottle
    children["remote-id-limit"] = &sessions.RemoteIdLimit
    children["mac-iwf-access-interface-throttle"] = &sessions.MacIwfAccessInterfaceThrottle
    children["access-interface-limit"] = &sessions.AccessInterfaceLimit
    children["mac-access-interface-throttle"] = &sessions.MacAccessInterfaceThrottle
    children["outer-vlan-limit"] = &sessions.OuterVlanLimit
    children["circuit-id-throttle"] = &sessions.CircuitIdThrottle
    children["mac-limit"] = &sessions.MacLimit
    children["circuit-id-limit"] = &sessions.CircuitIdLimit
    children["mac-iwf-limit"] = &sessions.MacIwfLimit
    children["mac-iwf-access-interface-limit"] = &sessions.MacIwfAccessInterfaceLimit
    children["inner-vlan-limit"] = &sessions.InnerVlanLimit
    children["outer-vlan-throttle"] = &sessions.OuterVlanThrottle
    children["mac-throttle"] = &sessions.MacThrottle
    children["circuit-id-and-remote-id-limit"] = &sessions.CircuitIdAndRemoteIdLimit
    children["vlan-limit"] = &sessions.VlanLimit
    children["mac-access-interface-limit"] = &sessions.MacAccessInterfaceLimit
    children["remote-id-throttle"] = &sessions.RemoteIdThrottle
    children["max-limit"] = &sessions.MaxLimit
    children["circuit-id-and-remote-id-throttle"] = &sessions.CircuitIdAndRemoteIdThrottle
    return children
}

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetYangName() string { return "sessions" }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions) GetParentYangName() string { return "pppoe-bba-group" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle
// Set VLAN (inner + outer tags) session
// throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetFilter() yfilter.YFilter { return vlanThrottle.YFilter }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) SetFilter(yf yfilter.YFilter) { vlanThrottle.YFilter = yf }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetSegmentPath() string {
    return "vlan-throttle"
}

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = vlanThrottle.Throttle
    leafs["request-period"] = vlanThrottle.RequestPeriod
    leafs["blocking-period"] = vlanThrottle.BlockingPeriod
    return leafs
}

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetYangName() string { return "vlan-throttle" }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) SetParent(parent types.Entity) { vlanThrottle.parent = parent }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetParent() types.Entity { return vlanThrottle.parent }

func (vlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle
// Set Inner VLAN session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetFilter() yfilter.YFilter { return innerVlanThrottle.YFilter }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) SetFilter(yf yfilter.YFilter) { innerVlanThrottle.YFilter = yf }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetSegmentPath() string {
    return "inner-vlan-throttle"
}

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = innerVlanThrottle.Throttle
    leafs["request-period"] = innerVlanThrottle.RequestPeriod
    leafs["blocking-period"] = innerVlanThrottle.BlockingPeriod
    return leafs
}

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetYangName() string { return "inner-vlan-throttle" }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) SetParent(parent types.Entity) { innerVlanThrottle.parent = parent }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetParent() types.Entity { return innerVlanThrottle.parent }

func (innerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit
// Set Remote ID session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-Remote ID limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Remote ID threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetFilter() yfilter.YFilter { return remoteIdLimit.YFilter }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) SetFilter(yf yfilter.YFilter) { remoteIdLimit.YFilter = yf }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetSegmentPath() string {
    return "remote-id-limit"
}

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = remoteIdLimit.Limit
    leafs["threshold"] = remoteIdLimit.Threshold
    return leafs
}

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetBundleName() string { return "cisco_ios_xr" }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetYangName() string { return "remote-id-limit" }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) SetParent(parent types.Entity) { remoteIdLimit.parent = parent }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetParent() types.Entity { return remoteIdLimit.parent }

func (remoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle
// Set per-MAC/Access interface throttle for IWF
// sessions
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetFilter() yfilter.YFilter { return macIwfAccessInterfaceThrottle.YFilter }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) SetFilter(yf yfilter.YFilter) { macIwfAccessInterfaceThrottle.YFilter = yf }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetSegmentPath() string {
    return "mac-iwf-access-interface-throttle"
}

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = macIwfAccessInterfaceThrottle.Throttle
    leafs["request-period"] = macIwfAccessInterfaceThrottle.RequestPeriod
    leafs["blocking-period"] = macIwfAccessInterfaceThrottle.BlockingPeriod
    return leafs
}

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetYangName() string { return "mac-iwf-access-interface-throttle" }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) SetParent(parent types.Entity) { macIwfAccessInterfaceThrottle.parent = parent }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetParent() types.Entity { return macIwfAccessInterfaceThrottle.parent }

func (macIwfAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit
// Set per-access interface limit
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-access interface session limit. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-access interface session threshold. The type is interface{} with range:
    // 1..65535.
    Threshold interface{}
}

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetFilter() yfilter.YFilter { return accessInterfaceLimit.YFilter }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) SetFilter(yf yfilter.YFilter) { accessInterfaceLimit.YFilter = yf }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetSegmentPath() string {
    return "access-interface-limit"
}

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = accessInterfaceLimit.Limit
    leafs["threshold"] = accessInterfaceLimit.Threshold
    return leafs
}

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetYangName() string { return "access-interface-limit" }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) SetParent(parent types.Entity) { accessInterfaceLimit.parent = parent }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetParent() types.Entity { return accessInterfaceLimit.parent }

func (accessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_AccessInterfaceLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle
// Set per-MAC/Access Interface throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetFilter() yfilter.YFilter { return macAccessInterfaceThrottle.YFilter }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) SetFilter(yf yfilter.YFilter) { macAccessInterfaceThrottle.YFilter = yf }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetSegmentPath() string {
    return "mac-access-interface-throttle"
}

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = macAccessInterfaceThrottle.Throttle
    leafs["request-period"] = macAccessInterfaceThrottle.RequestPeriod
    leafs["blocking-period"] = macAccessInterfaceThrottle.BlockingPeriod
    return leafs
}

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetYangName() string { return "mac-access-interface-throttle" }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) SetParent(parent types.Entity) { macAccessInterfaceThrottle.parent = parent }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetParent() types.Entity { return macAccessInterfaceThrottle.parent }

func (macAccessInterfaceThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit
// Set Outer VLAN session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-Outer VLAN limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Outer VLAN threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetFilter() yfilter.YFilter { return outerVlanLimit.YFilter }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) SetFilter(yf yfilter.YFilter) { outerVlanLimit.YFilter = yf }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetSegmentPath() string {
    return "outer-vlan-limit"
}

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = outerVlanLimit.Limit
    leafs["threshold"] = outerVlanLimit.Threshold
    return leafs
}

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetBundleName() string { return "cisco_ios_xr" }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetYangName() string { return "outer-vlan-limit" }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) SetParent(parent types.Entity) { outerVlanLimit.parent = parent }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetParent() types.Entity { return outerVlanLimit.parent }

func (outerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle
// Set Circuit ID session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetFilter() yfilter.YFilter { return circuitIdThrottle.YFilter }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) SetFilter(yf yfilter.YFilter) { circuitIdThrottle.YFilter = yf }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetSegmentPath() string {
    return "circuit-id-throttle"
}

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = circuitIdThrottle.Throttle
    leafs["request-period"] = circuitIdThrottle.RequestPeriod
    leafs["blocking-period"] = circuitIdThrottle.BlockingPeriod
    return leafs
}

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetYangName() string { return "circuit-id-throttle" }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) SetParent(parent types.Entity) { circuitIdThrottle.parent = parent }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetParent() types.Entity { return circuitIdThrottle.parent }

func (circuitIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit
// Set per-MAC address session limit and
// threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-MAC session limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-MAC session threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetFilter() yfilter.YFilter { return macLimit.YFilter }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) SetFilter(yf yfilter.YFilter) { macLimit.YFilter = yf }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetSegmentPath() string {
    return "mac-limit"
}

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = macLimit.Limit
    leafs["threshold"] = macLimit.Threshold
    return leafs
}

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetBundleName() string { return "cisco_ios_xr" }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetYangName() string { return "mac-limit" }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) SetParent(parent types.Entity) { macLimit.parent = parent }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetParent() types.Entity { return macLimit.parent }

func (macLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit
// Set Circuit ID session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-Circuit ID limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Circuit ID threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetFilter() yfilter.YFilter { return circuitIdLimit.YFilter }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) SetFilter(yf yfilter.YFilter) { circuitIdLimit.YFilter = yf }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetSegmentPath() string {
    return "circuit-id-limit"
}

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = circuitIdLimit.Limit
    leafs["threshold"] = circuitIdLimit.Threshold
    return leafs
}

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetYangName() string { return "circuit-id-limit" }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) SetParent(parent types.Entity) { circuitIdLimit.parent = parent }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetParent() types.Entity { return circuitIdLimit.parent }

func (circuitIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit
// Set per-MAC session limit and threshold for
// IWF sessions
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-MAC session limit for IWF sessions. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-MAC session threshold for IWF sessions. The type is interface{} with
    // range: 1..65535.
    Threshold interface{}
}

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetFilter() yfilter.YFilter { return macIwfLimit.YFilter }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) SetFilter(yf yfilter.YFilter) { macIwfLimit.YFilter = yf }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetSegmentPath() string {
    return "mac-iwf-limit"
}

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = macIwfLimit.Limit
    leafs["threshold"] = macIwfLimit.Threshold
    return leafs
}

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetBundleName() string { return "cisco_ios_xr" }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetYangName() string { return "mac-iwf-limit" }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) SetParent(parent types.Entity) { macIwfLimit.parent = parent }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetParent() types.Entity { return macIwfLimit.parent }

func (macIwfLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit
// Set per-MAC access interface session limit
// and threshold for IWF sessions
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-MAC access interface session limit for IWF sessions. The type is
    // interface{} with range: 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-MAC access interface session threshold for IWF sessions. The type is
    // interface{} with range: 1..65535.
    Threshold interface{}
}

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetFilter() yfilter.YFilter { return macIwfAccessInterfaceLimit.YFilter }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) SetFilter(yf yfilter.YFilter) { macIwfAccessInterfaceLimit.YFilter = yf }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetSegmentPath() string {
    return "mac-iwf-access-interface-limit"
}

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = macIwfAccessInterfaceLimit.Limit
    leafs["threshold"] = macIwfAccessInterfaceLimit.Threshold
    return leafs
}

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetBundleName() string { return "cisco_ios_xr" }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetYangName() string { return "mac-iwf-access-interface-limit" }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) SetParent(parent types.Entity) { macIwfAccessInterfaceLimit.parent = parent }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetParent() types.Entity { return macIwfAccessInterfaceLimit.parent }

func (macIwfAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacIwfAccessInterfaceLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit
// Set Inner VLAN session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-Inner VLAN limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Inner VLAN threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetFilter() yfilter.YFilter { return innerVlanLimit.YFilter }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) SetFilter(yf yfilter.YFilter) { innerVlanLimit.YFilter = yf }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetSegmentPath() string {
    return "inner-vlan-limit"
}

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = innerVlanLimit.Limit
    leafs["threshold"] = innerVlanLimit.Threshold
    return leafs
}

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetBundleName() string { return "cisco_ios_xr" }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetYangName() string { return "inner-vlan-limit" }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) SetParent(parent types.Entity) { innerVlanLimit.parent = parent }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetParent() types.Entity { return innerVlanLimit.parent }

func (innerVlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_InnerVlanLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle
// Set Outer VLAN session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetFilter() yfilter.YFilter { return outerVlanThrottle.YFilter }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) SetFilter(yf yfilter.YFilter) { outerVlanThrottle.YFilter = yf }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetSegmentPath() string {
    return "outer-vlan-throttle"
}

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = outerVlanThrottle.Throttle
    leafs["request-period"] = outerVlanThrottle.RequestPeriod
    leafs["blocking-period"] = outerVlanThrottle.BlockingPeriod
    return leafs
}

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetYangName() string { return "outer-vlan-throttle" }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) SetParent(parent types.Entity) { outerVlanThrottle.parent = parent }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetParent() types.Entity { return outerVlanThrottle.parent }

func (outerVlanThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_OuterVlanThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle
// Set per-MAC throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetFilter() yfilter.YFilter { return macThrottle.YFilter }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) SetFilter(yf yfilter.YFilter) { macThrottle.YFilter = yf }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetSegmentPath() string {
    return "mac-throttle"
}

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = macThrottle.Throttle
    leafs["request-period"] = macThrottle.RequestPeriod
    leafs["blocking-period"] = macThrottle.BlockingPeriod
    return leafs
}

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetYangName() string { return "mac-throttle" }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) SetParent(parent types.Entity) { macThrottle.parent = parent }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetParent() types.Entity { return macThrottle.parent }

func (macThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit
// Set Circuit ID and Remote ID session
// limit/threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-Circuit ID limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-Circuit ID threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetFilter() yfilter.YFilter { return circuitIdAndRemoteIdLimit.YFilter }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) SetFilter(yf yfilter.YFilter) { circuitIdAndRemoteIdLimit.YFilter = yf }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetSegmentPath() string {
    return "circuit-id-and-remote-id-limit"
}

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = circuitIdAndRemoteIdLimit.Limit
    leafs["threshold"] = circuitIdAndRemoteIdLimit.Threshold
    return leafs
}

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetYangName() string { return "circuit-id-and-remote-id-limit" }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) SetParent(parent types.Entity) { circuitIdAndRemoteIdLimit.parent = parent }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetParent() types.Entity { return circuitIdAndRemoteIdLimit.parent }

func (circuitIdAndRemoteIdLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit
// Set VLAN (inner + outer tags) session limit
// and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-VLAN (inner + outer tags) limit. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-VLAN (inner + outer tags) threshold. The type is interface{} with
    // range: 1..65535.
    Threshold interface{}
}

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetFilter() yfilter.YFilter { return vlanLimit.YFilter }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) SetFilter(yf yfilter.YFilter) { vlanLimit.YFilter = yf }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetSegmentPath() string {
    return "vlan-limit"
}

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = vlanLimit.Limit
    leafs["threshold"] = vlanLimit.Threshold
    return leafs
}

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetBundleName() string { return "cisco_ios_xr" }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetYangName() string { return "vlan-limit" }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) SetParent(parent types.Entity) { vlanLimit.parent = parent }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetParent() types.Entity { return vlanLimit.parent }

func (vlanLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_VlanLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit
// Set per-MAC access interface session limit
// and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-MAC access interface session limit. The type is interface{} with range:
    // 1..65535. This attribute is mandatory.
    Limit interface{}

    // Per-MAC access interface session threshold. The type is interface{} with
    // range: 1..65535.
    Threshold interface{}
}

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetFilter() yfilter.YFilter { return macAccessInterfaceLimit.YFilter }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) SetFilter(yf yfilter.YFilter) { macAccessInterfaceLimit.YFilter = yf }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetSegmentPath() string {
    return "mac-access-interface-limit"
}

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = macAccessInterfaceLimit.Limit
    leafs["threshold"] = macAccessInterfaceLimit.Threshold
    return leafs
}

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetBundleName() string { return "cisco_ios_xr" }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetYangName() string { return "mac-access-interface-limit" }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) SetParent(parent types.Entity) { macAccessInterfaceLimit.parent = parent }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetParent() types.Entity { return macAccessInterfaceLimit.parent }

func (macAccessInterfaceLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MacAccessInterfaceLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle
// Set Remote ID session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetFilter() yfilter.YFilter { return remoteIdThrottle.YFilter }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) SetFilter(yf yfilter.YFilter) { remoteIdThrottle.YFilter = yf }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetSegmentPath() string {
    return "remote-id-throttle"
}

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = remoteIdThrottle.Throttle
    leafs["request-period"] = remoteIdThrottle.RequestPeriod
    leafs["blocking-period"] = remoteIdThrottle.BlockingPeriod
    return leafs
}

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetYangName() string { return "remote-id-throttle" }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) SetParent(parent types.Entity) { remoteIdThrottle.parent = parent }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetParent() types.Entity { return remoteIdThrottle.parent }

func (remoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_RemoteIdThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit
// Set per-card session limit and threshold
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-card session limit. The type is interface{} with range: 1..65535. This
    // attribute is mandatory.
    Limit interface{}

    // Per-card session threshold. The type is interface{} with range: 1..65535.
    Threshold interface{}
}

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetFilter() yfilter.YFilter { return maxLimit.YFilter }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) SetFilter(yf yfilter.YFilter) { maxLimit.YFilter = yf }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetSegmentPath() string {
    return "max-limit"
}

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = maxLimit.Limit
    leafs["threshold"] = maxLimit.Threshold
    return leafs
}

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetBundleName() string { return "cisco_ios_xr" }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetYangName() string { return "max-limit" }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) SetParent(parent types.Entity) { maxLimit.parent = parent }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetParent() types.Entity { return maxLimit.parent }

func (maxLimit *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_MaxLimit) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle
// Set Circuit ID and Remote ID session throttle
// This type is a presence type.
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of requests at which to throttle. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    Throttle interface{}

    // Throttle request period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    RequestPeriod interface{}

    // Throttle blocking period. The type is interface{} with range: 1..100. This
    // attribute is mandatory.
    BlockingPeriod interface{}
}

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetFilter() yfilter.YFilter { return circuitIdAndRemoteIdThrottle.YFilter }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) SetFilter(yf yfilter.YFilter) { circuitIdAndRemoteIdThrottle.YFilter = yf }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    if yname == "request-period" { return "RequestPeriod" }
    if yname == "blocking-period" { return "BlockingPeriod" }
    return ""
}

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetSegmentPath() string {
    return "circuit-id-and-remote-id-throttle"
}

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["throttle"] = circuitIdAndRemoteIdThrottle.Throttle
    leafs["request-period"] = circuitIdAndRemoteIdThrottle.RequestPeriod
    leafs["blocking-period"] = circuitIdAndRemoteIdThrottle.BlockingPeriod
    return leafs
}

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetYangName() string { return "circuit-id-and-remote-id-throttle" }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) SetParent(parent types.Entity) { circuitIdAndRemoteIdThrottle.parent = parent }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetParent() types.Entity { return circuitIdAndRemoteIdThrottle.parent }

func (circuitIdAndRemoteIdThrottle *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_Sessions_CircuitIdAndRemoteIdThrottle) GetParentYangName() string { return "sessions" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets
// PPPoE control-packet configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the Priority to use for PPP and PPPoE control packets. The type is
    // interface{} with range: 0..7.
    Priority interface{}
}

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetFilter() yfilter.YFilter { return controlPackets.YFilter }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) SetFilter(yf yfilter.YFilter) { controlPackets.YFilter = yf }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetGoName(yname string) string {
    if yname == "priority" { return "Priority" }
    return ""
}

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetSegmentPath() string {
    return "control-packets"
}

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority"] = controlPackets.Priority
    return leafs
}

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetBundleName() string { return "cisco_ios_xr" }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetYangName() string { return "control-packets" }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) SetParent(parent types.Entity) { controlPackets.parent = parent }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetParent() types.Entity { return controlPackets.parent }

func (controlPackets *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_ControlPackets) GetParentYangName() string { return "pppoe-bba-group" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay
// PPPoE PADO delay configuration data
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // Units are millisecond.
    Default interface{}

    // Configure PADO delay dependent on the received Circuit ID. The type is
    // interface{} with range: 0..10000. Units are millisecond.
    CircuitId interface{}

    // Configure PADO delay dependent on the received Remote ID. The type is
    // interface{} with range: 0..10000. Units are millisecond.
    RemoteId interface{}

    // Delay the PADO response when the received Remote ID contains the given
    // string.
    RemoteIdSubstrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings

    // Delay the PADO response when there is an exact match on the received Remote
    // ID.
    RemoteIdStrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings

    // Delay the PADO response when there is an exact match on the received
    // Service Name.
    ServiceNameStrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings

    // Delay the PADO response when the received Circuit ID contains the given
    // string.
    CircuitIdSubstrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings

    // Delay the PADO response when the received Service Name contains the given
    // string.
    ServiceNameSubstrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings

    // Delay the PADO response when there is an exact match on the received
    // Circuit ID.
    CircuitIdStrings PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings
}

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetFilter() yfilter.YFilter { return paDoDelay.YFilter }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) SetFilter(yf yfilter.YFilter) { paDoDelay.YFilter = yf }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetGoName(yname string) string {
    if yname == "default" { return "Default" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "remote-id-substrings" { return "RemoteIdSubstrings" }
    if yname == "remote-id-strings" { return "RemoteIdStrings" }
    if yname == "service-name-strings" { return "ServiceNameStrings" }
    if yname == "circuit-id-substrings" { return "CircuitIdSubstrings" }
    if yname == "service-name-substrings" { return "ServiceNameSubstrings" }
    if yname == "circuit-id-strings" { return "CircuitIdStrings" }
    return ""
}

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetSegmentPath() string {
    return "pa-do-delay"
}

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-id-substrings" {
        return &paDoDelay.RemoteIdSubstrings
    }
    if childYangName == "remote-id-strings" {
        return &paDoDelay.RemoteIdStrings
    }
    if childYangName == "service-name-strings" {
        return &paDoDelay.ServiceNameStrings
    }
    if childYangName == "circuit-id-substrings" {
        return &paDoDelay.CircuitIdSubstrings
    }
    if childYangName == "service-name-substrings" {
        return &paDoDelay.ServiceNameSubstrings
    }
    if childYangName == "circuit-id-strings" {
        return &paDoDelay.CircuitIdStrings
    }
    return nil
}

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-id-substrings"] = &paDoDelay.RemoteIdSubstrings
    children["remote-id-strings"] = &paDoDelay.RemoteIdStrings
    children["service-name-strings"] = &paDoDelay.ServiceNameStrings
    children["circuit-id-substrings"] = &paDoDelay.CircuitIdSubstrings
    children["service-name-substrings"] = &paDoDelay.ServiceNameSubstrings
    children["circuit-id-strings"] = &paDoDelay.CircuitIdStrings
    return children
}

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default"] = paDoDelay.Default
    leafs["circuit-id"] = paDoDelay.CircuitId
    leafs["remote-id"] = paDoDelay.RemoteId
    return leafs
}

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetBundleName() string { return "cisco_ios_xr" }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetYangName() string { return "pa-do-delay" }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) SetParent(parent types.Entity) { paDoDelay.parent = parent }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetParent() types.Entity { return paDoDelay.parent }

func (paDoDelay *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay) GetParentYangName() string { return "pppoe-bba-group" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings
// Delay the PADO response when the received
// Remote ID contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay the PADO response when the received Remote ID contains the given
    // string. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring.
    RemoteIdSubstring []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring
}

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetFilter() yfilter.YFilter { return remoteIdSubstrings.YFilter }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) SetFilter(yf yfilter.YFilter) { remoteIdSubstrings.YFilter = yf }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetGoName(yname string) string {
    if yname == "remote-id-substring" { return "RemoteIdSubstring" }
    return ""
}

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetSegmentPath() string {
    return "remote-id-substrings"
}

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-id-substring" {
        for _, c := range remoteIdSubstrings.RemoteIdSubstring {
            if remoteIdSubstrings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring{}
        remoteIdSubstrings.RemoteIdSubstring = append(remoteIdSubstrings.RemoteIdSubstring, child)
        return &remoteIdSubstrings.RemoteIdSubstring[len(remoteIdSubstrings.RemoteIdSubstring)-1]
    }
    return nil
}

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteIdSubstrings.RemoteIdSubstring {
        children[remoteIdSubstrings.RemoteIdSubstring[i].GetSegmentPath()] = &remoteIdSubstrings.RemoteIdSubstring[i]
    }
    return children
}

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetBundleName() string { return "cisco_ios_xr" }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetYangName() string { return "remote-id-substrings" }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) SetParent(parent types.Entity) { remoteIdSubstrings.parent = parent }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetParent() types.Entity { return remoteIdSubstrings.parent }

func (remoteIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings) GetParentYangName() string { return "pa-do-delay" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring
// Delay the PADO response when the received
// Remote ID contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The string to be contained within the received
    // Remote ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetFilter() yfilter.YFilter { return remoteIdSubstring.YFilter }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) SetFilter(yf yfilter.YFilter) { remoteIdSubstring.YFilter = yf }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "delay" { return "Delay" }
    return ""
}

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetSegmentPath() string {
    return "remote-id-substring" + "[name='" + fmt.Sprintf("%v", remoteIdSubstring.Name) + "']"
}

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = remoteIdSubstring.Name
    leafs["delay"] = remoteIdSubstring.Delay
    return leafs
}

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetBundleName() string { return "cisco_ios_xr" }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetYangName() string { return "remote-id-substring" }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) SetParent(parent types.Entity) { remoteIdSubstring.parent = parent }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetParent() types.Entity { return remoteIdSubstring.parent }

func (remoteIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdSubstrings_RemoteIdSubstring) GetParentYangName() string { return "remote-id-substrings" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings
// Delay the PADO response when there is an
// exact match on the received Remote ID
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay the PADO response when there is an exact match on the received Remote
    // ID. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString.
    RemoteIdString []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString
}

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetFilter() yfilter.YFilter { return remoteIdStrings.YFilter }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) SetFilter(yf yfilter.YFilter) { remoteIdStrings.YFilter = yf }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetGoName(yname string) string {
    if yname == "remote-id-string" { return "RemoteIdString" }
    return ""
}

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetSegmentPath() string {
    return "remote-id-strings"
}

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-id-string" {
        for _, c := range remoteIdStrings.RemoteIdString {
            if remoteIdStrings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString{}
        remoteIdStrings.RemoteIdString = append(remoteIdStrings.RemoteIdString, child)
        return &remoteIdStrings.RemoteIdString[len(remoteIdStrings.RemoteIdString)-1]
    }
    return nil
}

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteIdStrings.RemoteIdString {
        children[remoteIdStrings.RemoteIdString[i].GetSegmentPath()] = &remoteIdStrings.RemoteIdString[i]
    }
    return children
}

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetBundleName() string { return "cisco_ios_xr" }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetYangName() string { return "remote-id-strings" }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) SetParent(parent types.Entity) { remoteIdStrings.parent = parent }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetParent() types.Entity { return remoteIdStrings.parent }

func (remoteIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings) GetParentYangName() string { return "pa-do-delay" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString
// Delay the PADO response when there is an
// exact match on the received Remote ID
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The string to exactly match the received Remote
    // ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetFilter() yfilter.YFilter { return remoteIdString.YFilter }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) SetFilter(yf yfilter.YFilter) { remoteIdString.YFilter = yf }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "delay" { return "Delay" }
    return ""
}

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetSegmentPath() string {
    return "remote-id-string" + "[name='" + fmt.Sprintf("%v", remoteIdString.Name) + "']"
}

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = remoteIdString.Name
    leafs["delay"] = remoteIdString.Delay
    return leafs
}

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetBundleName() string { return "cisco_ios_xr" }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetYangName() string { return "remote-id-string" }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) SetParent(parent types.Entity) { remoteIdString.parent = parent }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetParent() types.Entity { return remoteIdString.parent }

func (remoteIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_RemoteIdStrings_RemoteIdString) GetParentYangName() string { return "remote-id-strings" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings
// Delay the PADO response when there is an
// exact match on the received Service Name
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay the PADO response when there is an exact match on the received
    // Service Name. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString.
    ServiceNameString []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString
}

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetFilter() yfilter.YFilter { return serviceNameStrings.YFilter }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) SetFilter(yf yfilter.YFilter) { serviceNameStrings.YFilter = yf }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetGoName(yname string) string {
    if yname == "service-name-string" { return "ServiceNameString" }
    return ""
}

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetSegmentPath() string {
    return "service-name-strings"
}

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-name-string" {
        for _, c := range serviceNameStrings.ServiceNameString {
            if serviceNameStrings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString{}
        serviceNameStrings.ServiceNameString = append(serviceNameStrings.ServiceNameString, child)
        return &serviceNameStrings.ServiceNameString[len(serviceNameStrings.ServiceNameString)-1]
    }
    return nil
}

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serviceNameStrings.ServiceNameString {
        children[serviceNameStrings.ServiceNameString[i].GetSegmentPath()] = &serviceNameStrings.ServiceNameString[i]
    }
    return children
}

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetBundleName() string { return "cisco_ios_xr" }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetYangName() string { return "service-name-strings" }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) SetParent(parent types.Entity) { serviceNameStrings.parent = parent }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetParent() types.Entity { return serviceNameStrings.parent }

func (serviceNameStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings) GetParentYangName() string { return "pa-do-delay" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString
// Delay the PADO response when there is an
// exact match on the received Service Name
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The string to exactly match the received Service
    // Name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetFilter() yfilter.YFilter { return serviceNameString.YFilter }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) SetFilter(yf yfilter.YFilter) { serviceNameString.YFilter = yf }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "delay" { return "Delay" }
    return ""
}

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetSegmentPath() string {
    return "service-name-string" + "[name='" + fmt.Sprintf("%v", serviceNameString.Name) + "']"
}

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = serviceNameString.Name
    leafs["delay"] = serviceNameString.Delay
    return leafs
}

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetBundleName() string { return "cisco_ios_xr" }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetYangName() string { return "service-name-string" }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) SetParent(parent types.Entity) { serviceNameString.parent = parent }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetParent() types.Entity { return serviceNameString.parent }

func (serviceNameString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameStrings_ServiceNameString) GetParentYangName() string { return "service-name-strings" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings
// Delay the PADO response when the received
// Circuit ID contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay the PADO response when the received Circuit ID contains the given
    // string. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring.
    CircuitIdSubstring []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring
}

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetFilter() yfilter.YFilter { return circuitIdSubstrings.YFilter }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) SetFilter(yf yfilter.YFilter) { circuitIdSubstrings.YFilter = yf }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetGoName(yname string) string {
    if yname == "circuit-id-substring" { return "CircuitIdSubstring" }
    return ""
}

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetSegmentPath() string {
    return "circuit-id-substrings"
}

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "circuit-id-substring" {
        for _, c := range circuitIdSubstrings.CircuitIdSubstring {
            if circuitIdSubstrings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring{}
        circuitIdSubstrings.CircuitIdSubstring = append(circuitIdSubstrings.CircuitIdSubstring, child)
        return &circuitIdSubstrings.CircuitIdSubstring[len(circuitIdSubstrings.CircuitIdSubstring)-1]
    }
    return nil
}

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range circuitIdSubstrings.CircuitIdSubstring {
        children[circuitIdSubstrings.CircuitIdSubstring[i].GetSegmentPath()] = &circuitIdSubstrings.CircuitIdSubstring[i]
    }
    return children
}

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetYangName() string { return "circuit-id-substrings" }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) SetParent(parent types.Entity) { circuitIdSubstrings.parent = parent }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetParent() types.Entity { return circuitIdSubstrings.parent }

func (circuitIdSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings) GetParentYangName() string { return "pa-do-delay" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring
// Delay the PADO response when the received
// Circuit ID contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The string to be contained within the received
    // Circuit ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetFilter() yfilter.YFilter { return circuitIdSubstring.YFilter }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) SetFilter(yf yfilter.YFilter) { circuitIdSubstring.YFilter = yf }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "delay" { return "Delay" }
    return ""
}

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetSegmentPath() string {
    return "circuit-id-substring" + "[name='" + fmt.Sprintf("%v", circuitIdSubstring.Name) + "']"
}

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = circuitIdSubstring.Name
    leafs["delay"] = circuitIdSubstring.Delay
    return leafs
}

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetYangName() string { return "circuit-id-substring" }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) SetParent(parent types.Entity) { circuitIdSubstring.parent = parent }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetParent() types.Entity { return circuitIdSubstring.parent }

func (circuitIdSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdSubstrings_CircuitIdSubstring) GetParentYangName() string { return "circuit-id-substrings" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings
// Delay the PADO response when the received
// Service Name contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay the PADO response when the received Service Name contains the given
    // string. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring.
    ServiceNameSubstring []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring
}

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetFilter() yfilter.YFilter { return serviceNameSubstrings.YFilter }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) SetFilter(yf yfilter.YFilter) { serviceNameSubstrings.YFilter = yf }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetGoName(yname string) string {
    if yname == "service-name-substring" { return "ServiceNameSubstring" }
    return ""
}

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetSegmentPath() string {
    return "service-name-substrings"
}

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-name-substring" {
        for _, c := range serviceNameSubstrings.ServiceNameSubstring {
            if serviceNameSubstrings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring{}
        serviceNameSubstrings.ServiceNameSubstring = append(serviceNameSubstrings.ServiceNameSubstring, child)
        return &serviceNameSubstrings.ServiceNameSubstring[len(serviceNameSubstrings.ServiceNameSubstring)-1]
    }
    return nil
}

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serviceNameSubstrings.ServiceNameSubstring {
        children[serviceNameSubstrings.ServiceNameSubstring[i].GetSegmentPath()] = &serviceNameSubstrings.ServiceNameSubstring[i]
    }
    return children
}

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetBundleName() string { return "cisco_ios_xr" }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetYangName() string { return "service-name-substrings" }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) SetParent(parent types.Entity) { serviceNameSubstrings.parent = parent }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetParent() types.Entity { return serviceNameSubstrings.parent }

func (serviceNameSubstrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings) GetParentYangName() string { return "pa-do-delay" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring
// Delay the PADO response when the received
// Service Name contains the given string
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The string to be contained within the received
    // Service Name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetFilter() yfilter.YFilter { return serviceNameSubstring.YFilter }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) SetFilter(yf yfilter.YFilter) { serviceNameSubstring.YFilter = yf }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "delay" { return "Delay" }
    return ""
}

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetSegmentPath() string {
    return "service-name-substring" + "[name='" + fmt.Sprintf("%v", serviceNameSubstring.Name) + "']"
}

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = serviceNameSubstring.Name
    leafs["delay"] = serviceNameSubstring.Delay
    return leafs
}

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetBundleName() string { return "cisco_ios_xr" }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetYangName() string { return "service-name-substring" }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) SetParent(parent types.Entity) { serviceNameSubstring.parent = parent }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetParent() types.Entity { return serviceNameSubstring.parent }

func (serviceNameSubstring *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_ServiceNameSubstrings_ServiceNameSubstring) GetParentYangName() string { return "service-name-substrings" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings
// Delay the PADO response when there is an
// exact match on the received Circuit ID
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay the PADO response when there is an exact match on the received
    // Circuit ID. The type is slice of
    // PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString.
    CircuitIdString []PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString
}

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetFilter() yfilter.YFilter { return circuitIdStrings.YFilter }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) SetFilter(yf yfilter.YFilter) { circuitIdStrings.YFilter = yf }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetGoName(yname string) string {
    if yname == "circuit-id-string" { return "CircuitIdString" }
    return ""
}

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetSegmentPath() string {
    return "circuit-id-strings"
}

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "circuit-id-string" {
        for _, c := range circuitIdStrings.CircuitIdString {
            if circuitIdStrings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString{}
        circuitIdStrings.CircuitIdString = append(circuitIdStrings.CircuitIdString, child)
        return &circuitIdStrings.CircuitIdString[len(circuitIdStrings.CircuitIdString)-1]
    }
    return nil
}

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range circuitIdStrings.CircuitIdString {
        children[circuitIdStrings.CircuitIdString[i].GetSegmentPath()] = &circuitIdStrings.CircuitIdString[i]
    }
    return children
}

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetYangName() string { return "circuit-id-strings" }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) SetParent(parent types.Entity) { circuitIdStrings.parent = parent }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetParent() types.Entity { return circuitIdStrings.parent }

func (circuitIdStrings *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings) GetParentYangName() string { return "pa-do-delay" }

// PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString
// Delay the PADO response when there is an
// exact match on the received Circuit ID
type PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The string to exactly match the received Circuit
    // ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // PADO delay (in milliseconds). The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are millisecond.
    Delay interface{}
}

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetFilter() yfilter.YFilter { return circuitIdString.YFilter }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) SetFilter(yf yfilter.YFilter) { circuitIdString.YFilter = yf }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "delay" { return "Delay" }
    return ""
}

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetSegmentPath() string {
    return "circuit-id-string" + "[name='" + fmt.Sprintf("%v", circuitIdString.Name) + "']"
}

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = circuitIdString.Name
    leafs["delay"] = circuitIdString.Delay
    return leafs
}

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetBundleName() string { return "cisco_ios_xr" }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetYangName() string { return "circuit-id-string" }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) SetParent(parent types.Entity) { circuitIdString.parent = parent }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetParent() types.Entity { return circuitIdString.parent }

func (circuitIdString *PppoeCfg_PppoeBbaGroups_PppoeBbaGroup_PaDoDelay_CircuitIdStrings_CircuitIdString) GetParentYangName() string { return "circuit-id-strings" }

