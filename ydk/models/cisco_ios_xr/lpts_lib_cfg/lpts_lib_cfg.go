// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-lib package configuration.
// 
// This module contains definitions
// for the following management objects:
//   lpts: lpts configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lpts_lib_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lpts_lib_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lpts-lib-cfg lpts}", reflect.TypeOf(Lpts{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lpts-lib-cfg:lpts", reflect.TypeOf(Lpts{}))
}

// Lpts
// lpts configuration commands
type Lpts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pre IFiB Configuration .
    Ipolicer Lpts_Ipolicer

    // Configure penalty timeout value.
    Punt Lpts_Punt
}

func (lpts *Lpts) GetFilter() yfilter.YFilter { return lpts.YFilter }

func (lpts *Lpts) SetFilter(yf yfilter.YFilter) { lpts.YFilter = yf }

func (lpts *Lpts) GetGoName(yname string) string {
    if yname == "Cisco-IOS-XR-lpts-pre-ifib-cfg:ipolicer" { return "Ipolicer" }
    if yname == "Cisco-IOS-XR-lpts-punt-flowtrap-cfg:punt" { return "Punt" }
    return ""
}

func (lpts *Lpts) GetSegmentPath() string {
    return "Cisco-IOS-XR-lpts-lib-cfg:lpts"
}

func (lpts *Lpts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-lpts-pre-ifib-cfg:ipolicer" {
        return &lpts.Ipolicer
    }
    if childYangName == "Cisco-IOS-XR-lpts-punt-flowtrap-cfg:punt" {
        return &lpts.Punt
    }
    return nil
}

func (lpts *Lpts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-lpts-pre-ifib-cfg:ipolicer"] = &lpts.Ipolicer
    children["Cisco-IOS-XR-lpts-punt-flowtrap-cfg:punt"] = &lpts.Punt
    return children
}

func (lpts *Lpts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lpts *Lpts) GetBundleName() string { return "cisco_ios_xr" }

func (lpts *Lpts) GetYangName() string { return "lpts" }

func (lpts *Lpts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lpts *Lpts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lpts *Lpts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lpts *Lpts) SetParent(parent types.Entity) { lpts.parent = parent }

func (lpts *Lpts) GetParent() types.Entity { return lpts.parent }

func (lpts *Lpts) GetParentYangName() string { return "Cisco-IOS-XR-lpts-lib-cfg" }

// Lpts_Ipolicer
// Pre IFiB Configuration 
// This type is a presence type.
type Lpts_Ipolicer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table for ACLs.
    Ipv4Acls Lpts_Ipolicer_Ipv4Acls

    // Table for Flows.
    Flows Lpts_Ipolicer_Flows
}

func (ipolicer *Lpts_Ipolicer) GetFilter() yfilter.YFilter { return ipolicer.YFilter }

func (ipolicer *Lpts_Ipolicer) SetFilter(yf yfilter.YFilter) { ipolicer.YFilter = yf }

func (ipolicer *Lpts_Ipolicer) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "ipv4acls" { return "Ipv4Acls" }
    if yname == "flows" { return "Flows" }
    return ""
}

func (ipolicer *Lpts_Ipolicer) GetSegmentPath() string {
    return "Cisco-IOS-XR-lpts-pre-ifib-cfg:ipolicer"
}

func (ipolicer *Lpts_Ipolicer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4acls" {
        return &ipolicer.Ipv4Acls
    }
    if childYangName == "flows" {
        return &ipolicer.Flows
    }
    return nil
}

func (ipolicer *Lpts_Ipolicer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4acls"] = &ipolicer.Ipv4Acls
    children["flows"] = &ipolicer.Flows
    return children
}

func (ipolicer *Lpts_Ipolicer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ipolicer.Enable
    return leafs
}

func (ipolicer *Lpts_Ipolicer) GetBundleName() string { return "cisco_ios_xr" }

func (ipolicer *Lpts_Ipolicer) GetYangName() string { return "ipolicer" }

func (ipolicer *Lpts_Ipolicer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipolicer *Lpts_Ipolicer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipolicer *Lpts_Ipolicer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipolicer *Lpts_Ipolicer) SetParent(parent types.Entity) { ipolicer.parent = parent }

func (ipolicer *Lpts_Ipolicer) GetParent() types.Entity { return ipolicer.parent }

func (ipolicer *Lpts_Ipolicer) GetParentYangName() string { return "lpts" }

// Lpts_Ipolicer_Ipv4Acls
// Table for ACLs
type Lpts_Ipolicer_Ipv4Acls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL name. The type is slice of Lpts_Ipolicer_Ipv4Acls_Ipv4Acl.
    Ipv4Acl []Lpts_Ipolicer_Ipv4Acls_Ipv4Acl
}

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetFilter() yfilter.YFilter { return ipv4Acls.YFilter }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) SetFilter(yf yfilter.YFilter) { ipv4Acls.YFilter = yf }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetGoName(yname string) string {
    if yname == "ipv4acl" { return "Ipv4Acl" }
    return ""
}

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetSegmentPath() string {
    return "ipv4acls"
}

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4acl" {
        for _, c := range ipv4Acls.Ipv4Acl {
            if ipv4Acls.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lpts_Ipolicer_Ipv4Acls_Ipv4Acl{}
        ipv4Acls.Ipv4Acl = append(ipv4Acls.Ipv4Acl, child)
        return &ipv4Acls.Ipv4Acl[len(ipv4Acls.Ipv4Acl)-1]
    }
    return nil
}

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4Acls.Ipv4Acl {
        children[ipv4Acls.Ipv4Acl[i].GetSegmentPath()] = &ipv4Acls.Ipv4Acl[i]
    }
    return children
}

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetYangName() string { return "ipv4acls" }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) SetParent(parent types.Entity) { ipv4Acls.parent = parent }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetParent() types.Entity { return ipv4Acls.parent }

func (ipv4Acls *Lpts_Ipolicer_Ipv4Acls) GetParentYangName() string { return "ipolicer" }

// Lpts_Ipolicer_Ipv4Acls_Ipv4Acl
// ACL name
type Lpts_Ipolicer_Ipv4Acls_Ipv4Acl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ACL name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    AclName interface{}

    // VRF list.
    Ipv4VrfNames Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames
}

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetFilter() yfilter.YFilter { return ipv4Acl.YFilter }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) SetFilter(yf yfilter.YFilter) { ipv4Acl.YFilter = yf }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetGoName(yname string) string {
    if yname == "acl-name" { return "AclName" }
    if yname == "ipv4vrf-names" { return "Ipv4VrfNames" }
    return ""
}

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetSegmentPath() string {
    return "ipv4acl" + "[acl-name='" + fmt.Sprintf("%v", ipv4Acl.AclName) + "']"
}

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4vrf-names" {
        return &ipv4Acl.Ipv4VrfNames
    }
    return nil
}

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4vrf-names"] = &ipv4Acl.Ipv4VrfNames
    return children
}

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["acl-name"] = ipv4Acl.AclName
    return leafs
}

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetYangName() string { return "ipv4acl" }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) SetParent(parent types.Entity) { ipv4Acl.parent = parent }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetParent() types.Entity { return ipv4Acl.parent }

func (ipv4Acl *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl) GetParentYangName() string { return "ipv4acls" }

// Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames
// VRF list
type Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is slice of
    // Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName.
    Ipv4VrfName []Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName
}

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetFilter() yfilter.YFilter { return ipv4VrfNames.YFilter }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) SetFilter(yf yfilter.YFilter) { ipv4VrfNames.YFilter = yf }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetGoName(yname string) string {
    if yname == "ipv4vrf-name" { return "Ipv4VrfName" }
    return ""
}

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetSegmentPath() string {
    return "ipv4vrf-names"
}

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4vrf-name" {
        for _, c := range ipv4VrfNames.Ipv4VrfName {
            if ipv4VrfNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName{}
        ipv4VrfNames.Ipv4VrfName = append(ipv4VrfNames.Ipv4VrfName, child)
        return &ipv4VrfNames.Ipv4VrfName[len(ipv4VrfNames.Ipv4VrfName)-1]
    }
    return nil
}

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4VrfNames.Ipv4VrfName {
        children[ipv4VrfNames.Ipv4VrfName[i].GetSegmentPath()] = &ipv4VrfNames.Ipv4VrfName[i]
    }
    return children
}

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetYangName() string { return "ipv4vrf-names" }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) SetParent(parent types.Entity) { ipv4VrfNames.parent = parent }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetParent() types.Entity { return ipv4VrfNames.parent }

func (ipv4VrfNames *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames) GetParentYangName() string { return "ipv4acl" }

// Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName
// VRF name
type Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // pre-ifib policer rate config commands. The type is interface{} with range:
    // 0..100000.
    AclRate interface{}
}

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetFilter() yfilter.YFilter { return ipv4VrfName.YFilter }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) SetFilter(yf yfilter.YFilter) { ipv4VrfName.YFilter = yf }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "acl-rate" { return "AclRate" }
    return ""
}

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetSegmentPath() string {
    return "ipv4vrf-name" + "[vrf-name='" + fmt.Sprintf("%v", ipv4VrfName.VrfName) + "']"
}

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = ipv4VrfName.VrfName
    leafs["acl-rate"] = ipv4VrfName.AclRate
    return leafs
}

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetYangName() string { return "ipv4vrf-name" }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) SetParent(parent types.Entity) { ipv4VrfName.parent = parent }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetParent() types.Entity { return ipv4VrfName.parent }

func (ipv4VrfName *Lpts_Ipolicer_Ipv4Acls_Ipv4Acl_Ipv4VrfNames_Ipv4VrfName) GetParentYangName() string { return "ipv4vrf-names" }

// Lpts_Ipolicer_Flows
// Table for Flows
type Lpts_Ipolicer_Flows struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // selected flow type. The type is slice of Lpts_Ipolicer_Flows_Flow.
    Flow []Lpts_Ipolicer_Flows_Flow
}

func (flows *Lpts_Ipolicer_Flows) GetFilter() yfilter.YFilter { return flows.YFilter }

func (flows *Lpts_Ipolicer_Flows) SetFilter(yf yfilter.YFilter) { flows.YFilter = yf }

func (flows *Lpts_Ipolicer_Flows) GetGoName(yname string) string {
    if yname == "flow" { return "Flow" }
    return ""
}

func (flows *Lpts_Ipolicer_Flows) GetSegmentPath() string {
    return "flows"
}

func (flows *Lpts_Ipolicer_Flows) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow" {
        for _, c := range flows.Flow {
            if flows.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lpts_Ipolicer_Flows_Flow{}
        flows.Flow = append(flows.Flow, child)
        return &flows.Flow[len(flows.Flow)-1]
    }
    return nil
}

func (flows *Lpts_Ipolicer_Flows) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flows.Flow {
        children[flows.Flow[i].GetSegmentPath()] = &flows.Flow[i]
    }
    return children
}

func (flows *Lpts_Ipolicer_Flows) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flows *Lpts_Ipolicer_Flows) GetBundleName() string { return "cisco_ios_xr" }

func (flows *Lpts_Ipolicer_Flows) GetYangName() string { return "flows" }

func (flows *Lpts_Ipolicer_Flows) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flows *Lpts_Ipolicer_Flows) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flows *Lpts_Ipolicer_Flows) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flows *Lpts_Ipolicer_Flows) SetParent(parent types.Entity) { flows.parent = parent }

func (flows *Lpts_Ipolicer_Flows) GetParent() types.Entity { return flows.parent }

func (flows *Lpts_Ipolicer_Flows) GetParentYangName() string { return "ipolicer" }

// Lpts_Ipolicer_Flows_Flow
// selected flow type
type Lpts_Ipolicer_Flows_Flow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LPTS Flow Type. The type is LptsFlow.
    FlowType interface{}

    // Configured rate value. The type is interface{} with range:
    // -2147483648..2147483647.
    Rate interface{}

    // TOS Precedence value(s).
    Precedences Lpts_Ipolicer_Flows_Flow_Precedences
}

func (flow *Lpts_Ipolicer_Flows_Flow) GetFilter() yfilter.YFilter { return flow.YFilter }

func (flow *Lpts_Ipolicer_Flows_Flow) SetFilter(yf yfilter.YFilter) { flow.YFilter = yf }

func (flow *Lpts_Ipolicer_Flows_Flow) GetGoName(yname string) string {
    if yname == "flow-type" { return "FlowType" }
    if yname == "rate" { return "Rate" }
    if yname == "precedences" { return "Precedences" }
    return ""
}

func (flow *Lpts_Ipolicer_Flows_Flow) GetSegmentPath() string {
    return "flow" + "[flow-type='" + fmt.Sprintf("%v", flow.FlowType) + "']"
}

func (flow *Lpts_Ipolicer_Flows_Flow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "precedences" {
        return &flow.Precedences
    }
    return nil
}

func (flow *Lpts_Ipolicer_Flows_Flow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["precedences"] = &flow.Precedences
    return children
}

func (flow *Lpts_Ipolicer_Flows_Flow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-type"] = flow.FlowType
    leafs["rate"] = flow.Rate
    return leafs
}

func (flow *Lpts_Ipolicer_Flows_Flow) GetBundleName() string { return "cisco_ios_xr" }

func (flow *Lpts_Ipolicer_Flows_Flow) GetYangName() string { return "flow" }

func (flow *Lpts_Ipolicer_Flows_Flow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flow *Lpts_Ipolicer_Flows_Flow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flow *Lpts_Ipolicer_Flows_Flow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flow *Lpts_Ipolicer_Flows_Flow) SetParent(parent types.Entity) { flow.parent = parent }

func (flow *Lpts_Ipolicer_Flows_Flow) GetParent() types.Entity { return flow.parent }

func (flow *Lpts_Ipolicer_Flows_Flow) GetParentYangName() string { return "flows" }

// Lpts_Ipolicer_Flows_Flow_Precedences
// TOS Precedence value(s)
type Lpts_Ipolicer_Flows_Flow_Precedences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Precedence values. The type is one of the following types: slice of  
    // :go:struct:`LptsPreIFibPrecedenceNumber
    // <ydk/models/lpts_pre_ifib_cfg/LptsPreIFibPrecedenceNumber>`, or slice of
    // int with range: 0..7.
    Precedence []interface{}
}

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetFilter() yfilter.YFilter { return precedences.YFilter }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) SetFilter(yf yfilter.YFilter) { precedences.YFilter = yf }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetGoName(yname string) string {
    if yname == "precedence" { return "Precedence" }
    return ""
}

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetSegmentPath() string {
    return "precedences"
}

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["precedence"] = precedences.Precedence
    return leafs
}

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetBundleName() string { return "cisco_ios_xr" }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetYangName() string { return "precedences" }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) SetParent(parent types.Entity) { precedences.parent = parent }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetParent() types.Entity { return precedences.parent }

func (precedences *Lpts_Ipolicer_Flows_Flow_Precedences) GetParentYangName() string { return "flow" }

// Lpts_Punt
// Configure penalty timeout value
type Lpts_Punt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // excessive punt flow trap configuration commands.
    Flowtrap Lpts_Punt_Flowtrap
}

func (punt *Lpts_Punt) GetFilter() yfilter.YFilter { return punt.YFilter }

func (punt *Lpts_Punt) SetFilter(yf yfilter.YFilter) { punt.YFilter = yf }

func (punt *Lpts_Punt) GetGoName(yname string) string {
    if yname == "flowtrap" { return "Flowtrap" }
    return ""
}

func (punt *Lpts_Punt) GetSegmentPath() string {
    return "Cisco-IOS-XR-lpts-punt-flowtrap-cfg:punt"
}

func (punt *Lpts_Punt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flowtrap" {
        return &punt.Flowtrap
    }
    return nil
}

func (punt *Lpts_Punt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flowtrap"] = &punt.Flowtrap
    return children
}

func (punt *Lpts_Punt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (punt *Lpts_Punt) GetBundleName() string { return "cisco_ios_xr" }

func (punt *Lpts_Punt) GetYangName() string { return "punt" }

func (punt *Lpts_Punt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (punt *Lpts_Punt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (punt *Lpts_Punt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (punt *Lpts_Punt) SetParent(parent types.Entity) { punt.parent = parent }

func (punt *Lpts_Punt) GetParent() types.Entity { return punt.parent }

func (punt *Lpts_Punt) GetParentYangName() string { return "lpts" }

// Lpts_Punt_Flowtrap
// excessive punt flow trap configuration commands
type Lpts_Punt_Flowtrap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum flow gap in milliseconds. The type is interface{} with range:
    // 1..60000.
    MaxFlowGap interface{}

    // Should be power of 2. Any one of 1,2,4,8,16,32 ,64,128. The type is
    // interface{} with range: 1..128.
    EtSize interface{}

    // Eviction threshold, should be less than report-threshold. The type is
    // interface{} with range: 1..65535.
    EvictionThreshold interface{}

    // Threshold to cross for a flow to be considered as bad actor flow. The type
    // is interface{} with range: 1..65535.
    ReportThreshold interface{}

    // Enable trap based on source mac on non-subscriber interface. The type is
    // interface{} with range: -2147483648..2147483647.
    NonSubscriberInterfaces interface{}

    // Probability of packets to be sampled. The type is string with length:
    // 1..32.
    SampleProb interface{}

    // Eviction search limit, should be less than trap-size. The type is
    // interface{} with range: 1..128.
    EvictionSearchLimit interface{}

    // Allow routing protocols to pass through copp sampler. The type is bool.
    RoutingProtocolsEnable interface{}

    // Enable the trap on subscriber interfaces. The type is bool.
    SubscriberInterfaces interface{}

    // Identify flow based on interface and flowtype. The type is bool.
    InterfaceBasedFlow interface{}

    // Dampening period for a bad actor flow in milliseconds. The type is
    // interface{} with range: 5000..60000.
    Dampening interface{}

    // Configure penalty policing rate.
    PenaltyRates Lpts_Punt_Flowtrap_PenaltyRates

    // Configure penalty timeout value.
    PenaltyTimeouts Lpts_Punt_Flowtrap_PenaltyTimeouts

    // Exclude an item from all traps.
    Exclude Lpts_Punt_Flowtrap_Exclude
}

func (flowtrap *Lpts_Punt_Flowtrap) GetFilter() yfilter.YFilter { return flowtrap.YFilter }

func (flowtrap *Lpts_Punt_Flowtrap) SetFilter(yf yfilter.YFilter) { flowtrap.YFilter = yf }

func (flowtrap *Lpts_Punt_Flowtrap) GetGoName(yname string) string {
    if yname == "max-flow-gap" { return "MaxFlowGap" }
    if yname == "et-size" { return "EtSize" }
    if yname == "eviction-threshold" { return "EvictionThreshold" }
    if yname == "report-threshold" { return "ReportThreshold" }
    if yname == "non-subscriber-interfaces" { return "NonSubscriberInterfaces" }
    if yname == "sample-prob" { return "SampleProb" }
    if yname == "eviction-search-limit" { return "EvictionSearchLimit" }
    if yname == "routing-protocols-enable" { return "RoutingProtocolsEnable" }
    if yname == "subscriber-interfaces" { return "SubscriberInterfaces" }
    if yname == "interface-based-flow" { return "InterfaceBasedFlow" }
    if yname == "dampening" { return "Dampening" }
    if yname == "penalty-rates" { return "PenaltyRates" }
    if yname == "penalty-timeouts" { return "PenaltyTimeouts" }
    if yname == "exclude" { return "Exclude" }
    return ""
}

func (flowtrap *Lpts_Punt_Flowtrap) GetSegmentPath() string {
    return "flowtrap"
}

func (flowtrap *Lpts_Punt_Flowtrap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "penalty-rates" {
        return &flowtrap.PenaltyRates
    }
    if childYangName == "penalty-timeouts" {
        return &flowtrap.PenaltyTimeouts
    }
    if childYangName == "exclude" {
        return &flowtrap.Exclude
    }
    return nil
}

func (flowtrap *Lpts_Punt_Flowtrap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["penalty-rates"] = &flowtrap.PenaltyRates
    children["penalty-timeouts"] = &flowtrap.PenaltyTimeouts
    children["exclude"] = &flowtrap.Exclude
    return children
}

func (flowtrap *Lpts_Punt_Flowtrap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-flow-gap"] = flowtrap.MaxFlowGap
    leafs["et-size"] = flowtrap.EtSize
    leafs["eviction-threshold"] = flowtrap.EvictionThreshold
    leafs["report-threshold"] = flowtrap.ReportThreshold
    leafs["non-subscriber-interfaces"] = flowtrap.NonSubscriberInterfaces
    leafs["sample-prob"] = flowtrap.SampleProb
    leafs["eviction-search-limit"] = flowtrap.EvictionSearchLimit
    leafs["routing-protocols-enable"] = flowtrap.RoutingProtocolsEnable
    leafs["subscriber-interfaces"] = flowtrap.SubscriberInterfaces
    leafs["interface-based-flow"] = flowtrap.InterfaceBasedFlow
    leafs["dampening"] = flowtrap.Dampening
    return leafs
}

func (flowtrap *Lpts_Punt_Flowtrap) GetBundleName() string { return "cisco_ios_xr" }

func (flowtrap *Lpts_Punt_Flowtrap) GetYangName() string { return "flowtrap" }

func (flowtrap *Lpts_Punt_Flowtrap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowtrap *Lpts_Punt_Flowtrap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowtrap *Lpts_Punt_Flowtrap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowtrap *Lpts_Punt_Flowtrap) SetParent(parent types.Entity) { flowtrap.parent = parent }

func (flowtrap *Lpts_Punt_Flowtrap) GetParent() types.Entity { return flowtrap.parent }

func (flowtrap *Lpts_Punt_Flowtrap) GetParentYangName() string { return "punt" }

// Lpts_Punt_Flowtrap_PenaltyRates
// Configure penalty policing rate
type Lpts_Punt_Flowtrap_PenaltyRates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is slice of Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate.
    PenaltyRate []Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate
}

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetFilter() yfilter.YFilter { return penaltyRates.YFilter }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) SetFilter(yf yfilter.YFilter) { penaltyRates.YFilter = yf }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetGoName(yname string) string {
    if yname == "penalty-rate" { return "PenaltyRate" }
    return ""
}

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetSegmentPath() string {
    return "penalty-rates"
}

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "penalty-rate" {
        for _, c := range penaltyRates.PenaltyRate {
            if penaltyRates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate{}
        penaltyRates.PenaltyRate = append(penaltyRates.PenaltyRate, child)
        return &penaltyRates.PenaltyRate[len(penaltyRates.PenaltyRate)-1]
    }
    return nil
}

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range penaltyRates.PenaltyRate {
        children[penaltyRates.PenaltyRate[i].GetSegmentPath()] = &penaltyRates.PenaltyRate[i]
    }
    return children
}

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetBundleName() string { return "cisco_ios_xr" }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetYangName() string { return "penalty-rates" }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) SetParent(parent types.Entity) { penaltyRates.parent = parent }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetParent() types.Entity { return penaltyRates.parent }

func (penaltyRates *Lpts_Punt_Flowtrap_PenaltyRates) GetParentYangName() string { return "flowtrap" }

// Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate
// none
type Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is LptsPuntFlowtrapProtoId.
    ProtocolName interface{}

    // Penalty policer rate in packets-per-second. The type is interface{} with
    // range: 2..100. This attribute is mandatory.
    Rate interface{}
}

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetFilter() yfilter.YFilter { return penaltyRate.YFilter }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) SetFilter(yf yfilter.YFilter) { penaltyRate.YFilter = yf }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "rate" { return "Rate" }
    return ""
}

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetSegmentPath() string {
    return "penalty-rate" + "[protocol-name='" + fmt.Sprintf("%v", penaltyRate.ProtocolName) + "']"
}

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = penaltyRate.ProtocolName
    leafs["rate"] = penaltyRate.Rate
    return leafs
}

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetBundleName() string { return "cisco_ios_xr" }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetYangName() string { return "penalty-rate" }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) SetParent(parent types.Entity) { penaltyRate.parent = parent }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetParent() types.Entity { return penaltyRate.parent }

func (penaltyRate *Lpts_Punt_Flowtrap_PenaltyRates_PenaltyRate) GetParentYangName() string { return "penalty-rates" }

// Lpts_Punt_Flowtrap_PenaltyTimeouts
// Configure penalty timeout value
type Lpts_Punt_Flowtrap_PenaltyTimeouts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is slice of
    // Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout.
    PenaltyTimeout []Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout
}

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetFilter() yfilter.YFilter { return penaltyTimeouts.YFilter }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) SetFilter(yf yfilter.YFilter) { penaltyTimeouts.YFilter = yf }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetGoName(yname string) string {
    if yname == "penalty-timeout" { return "PenaltyTimeout" }
    return ""
}

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetSegmentPath() string {
    return "penalty-timeouts"
}

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "penalty-timeout" {
        for _, c := range penaltyTimeouts.PenaltyTimeout {
            if penaltyTimeouts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout{}
        penaltyTimeouts.PenaltyTimeout = append(penaltyTimeouts.PenaltyTimeout, child)
        return &penaltyTimeouts.PenaltyTimeout[len(penaltyTimeouts.PenaltyTimeout)-1]
    }
    return nil
}

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range penaltyTimeouts.PenaltyTimeout {
        children[penaltyTimeouts.PenaltyTimeout[i].GetSegmentPath()] = &penaltyTimeouts.PenaltyTimeout[i]
    }
    return children
}

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetBundleName() string { return "cisco_ios_xr" }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetYangName() string { return "penalty-timeouts" }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) SetParent(parent types.Entity) { penaltyTimeouts.parent = parent }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetParent() types.Entity { return penaltyTimeouts.parent }

func (penaltyTimeouts *Lpts_Punt_Flowtrap_PenaltyTimeouts) GetParentYangName() string { return "flowtrap" }

// Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout
// none
type Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is LptsPuntFlowtrapProtoId.
    ProtocolName interface{}

    // Timeout value in minutes. The type is interface{} with range: 1..1000. This
    // attribute is mandatory.
    Timeout interface{}
}

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetFilter() yfilter.YFilter { return penaltyTimeout.YFilter }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) SetFilter(yf yfilter.YFilter) { penaltyTimeout.YFilter = yf }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetSegmentPath() string {
    return "penalty-timeout" + "[protocol-name='" + fmt.Sprintf("%v", penaltyTimeout.ProtocolName) + "']"
}

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = penaltyTimeout.ProtocolName
    leafs["timeout"] = penaltyTimeout.Timeout
    return leafs
}

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetBundleName() string { return "cisco_ios_xr" }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetYangName() string { return "penalty-timeout" }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) SetParent(parent types.Entity) { penaltyTimeout.parent = parent }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetParent() types.Entity { return penaltyTimeout.parent }

func (penaltyTimeout *Lpts_Punt_Flowtrap_PenaltyTimeouts_PenaltyTimeout) GetParentYangName() string { return "penalty-timeouts" }

// Lpts_Punt_Flowtrap_Exclude
// Exclude an item from all traps
type Lpts_Punt_Flowtrap_Exclude struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none.
    InterfaceNames Lpts_Punt_Flowtrap_Exclude_InterfaceNames
}

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetFilter() yfilter.YFilter { return exclude.YFilter }

func (exclude *Lpts_Punt_Flowtrap_Exclude) SetFilter(yf yfilter.YFilter) { exclude.YFilter = yf }

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetGoName(yname string) string {
    if yname == "interface-names" { return "InterfaceNames" }
    return ""
}

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetSegmentPath() string {
    return "exclude"
}

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-names" {
        return &exclude.InterfaceNames
    }
    return nil
}

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-names"] = &exclude.InterfaceNames
    return children
}

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetBundleName() string { return "cisco_ios_xr" }

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetYangName() string { return "exclude" }

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exclude *Lpts_Punt_Flowtrap_Exclude) SetParent(parent types.Entity) { exclude.parent = parent }

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetParent() types.Entity { return exclude.parent }

func (exclude *Lpts_Punt_Flowtrap_Exclude) GetParentYangName() string { return "flowtrap" }

// Lpts_Punt_Flowtrap_Exclude_InterfaceNames
// none
type Lpts_Punt_Flowtrap_Exclude_InterfaceNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of interface to exclude from all traps. The type is slice of
    // Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName.
    InterfaceName []Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName
}

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetFilter() yfilter.YFilter { return interfaceNames.YFilter }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) SetFilter(yf yfilter.YFilter) { interfaceNames.YFilter = yf }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetSegmentPath() string {
    return "interface-names"
}

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-name" {
        for _, c := range interfaceNames.InterfaceName {
            if interfaceNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName{}
        interfaceNames.InterfaceName = append(interfaceNames.InterfaceName, child)
        return &interfaceNames.InterfaceName[len(interfaceNames.InterfaceName)-1]
    }
    return nil
}

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceNames.InterfaceName {
        children[interfaceNames.InterfaceName[i].GetSegmentPath()] = &interfaceNames.InterfaceName[i]
    }
    return children
}

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetYangName() string { return "interface-names" }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) SetParent(parent types.Entity) { interfaceNames.parent = parent }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetParent() types.Entity { return interfaceNames.parent }

func (interfaceNames *Lpts_Punt_Flowtrap_Exclude_InterfaceNames) GetParentYangName() string { return "exclude" }

// Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName
// Name of interface to exclude from all traps
type Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of interface to exclude from all traps. The
    // type is string with pattern: [a-zA-Z0-9./-]+.
    Ifname interface{}

    // Enabled or disabled. The type is bool. This attribute is mandatory.
    Id1 interface{}
}

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetFilter() yfilter.YFilter { return interfaceName.YFilter }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) SetFilter(yf yfilter.YFilter) { interfaceName.YFilter = yf }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetGoName(yname string) string {
    if yname == "ifname" { return "Ifname" }
    if yname == "id1" { return "Id1" }
    return ""
}

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetSegmentPath() string {
    return "interface-name" + "[ifname='" + fmt.Sprintf("%v", interfaceName.Ifname) + "']"
}

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifname"] = interfaceName.Ifname
    leafs["id1"] = interfaceName.Id1
    return leafs
}

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetYangName() string { return "interface-name" }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) SetParent(parent types.Entity) { interfaceName.parent = parent }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetParent() types.Entity { return interfaceName.parent }

func (interfaceName *Lpts_Punt_Flowtrap_Exclude_InterfaceNames_InterfaceName) GetParentYangName() string { return "interface-names" }

