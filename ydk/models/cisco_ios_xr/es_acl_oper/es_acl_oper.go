// This module contains a collection of YANG definitions
// for Cisco IOS-XR es-acl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   es-acl: Root class of ES ACL Oper schema tree
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package es_acl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package es_acl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-es-acl-oper es-acl}", reflect.TypeOf(EsAcl{}))
    ydk.RegisterEntity("Cisco-IOS-XR-es-acl-oper:es-acl", reflect.TypeOf(EsAcl{}))
}

// AclAction represents Acl action
type AclAction string

const (
    // Deny
    AclAction_deny AclAction = "deny"

    // Permit
    AclAction_permit AclAction = "permit"

    // Encrypt
    AclAction_encrypt AclAction = "encrypt"

    // Bypass
    AclAction_bypass AclAction = "bypass"

    // Fallthrough
    AclAction_fallthrough_ AclAction = "fallthrough"

    // Invalid
    AclAction_invalid AclAction = "invalid"
)

// AclAce1 represents ACE Types
type AclAce1 string

const (
    // This is Normal ACE
    AclAce1_normal AclAce1 = "normal"

    // This is Remark ACE
    AclAce1_remark AclAce1 = "remark"

    // This is ABF ACE
    AclAce1_abf AclAce1 = "abf"
)

// AclAce1 represents ACE Types
type AclAce1 string

const (
    // This is Normal ACE
    AclAce1_normal AclAce1 = "normal"

    // This is Remark ACE
    AclAce1_remark AclAce1 = "remark"

    // This is ABF ACE
    AclAce1_abf AclAce1 = "abf"
)

// EsAcl
// Root class of ES ACL Oper schema tree
type EsAcl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Out Of Resources, Limits to the resources allocatable.
    Active EsAcl_Active
}

func (esAcl *EsAcl) GetFilter() yfilter.YFilter { return esAcl.YFilter }

func (esAcl *EsAcl) SetFilter(yf yfilter.YFilter) { esAcl.YFilter = yf }

func (esAcl *EsAcl) GetGoName(yname string) string {
    if yname == "active" { return "Active" }
    return ""
}

func (esAcl *EsAcl) GetSegmentPath() string {
    return "Cisco-IOS-XR-es-acl-oper:es-acl"
}

func (esAcl *EsAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active" {
        return &esAcl.Active
    }
    return nil
}

func (esAcl *EsAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["active"] = &esAcl.Active
    return children
}

func (esAcl *EsAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (esAcl *EsAcl) GetBundleName() string { return "cisco_ios_xr" }

func (esAcl *EsAcl) GetYangName() string { return "es-acl" }

func (esAcl *EsAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esAcl *EsAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esAcl *EsAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esAcl *EsAcl) SetParent(parent types.Entity) { esAcl.parent = parent }

func (esAcl *EsAcl) GetParent() types.Entity { return esAcl.parent }

func (esAcl *EsAcl) GetParentYangName() string { return "Cisco-IOS-XR-es-acl-oper" }

// EsAcl_Active
// Out Of Resources, Limits to the resources
// allocatable
type EsAcl_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Out Of Resources, Limits to the resources allocatable.
    Oor EsAcl_Active_Oor

    // List containing ACLs.
    List EsAcl_Active_List

    // Resource occupation details for ACLs.
    OorAcls EsAcl_Active_OorAcls

    // Table of Usage statistics of ACLs at different nodes.
    Usages EsAcl_Active_Usages
}

func (active *EsAcl_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *EsAcl_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *EsAcl_Active) GetGoName(yname string) string {
    if yname == "oor" { return "Oor" }
    if yname == "list" { return "List" }
    if yname == "oor-acls" { return "OorAcls" }
    if yname == "usages" { return "Usages" }
    return ""
}

func (active *EsAcl_Active) GetSegmentPath() string {
    return "active"
}

func (active *EsAcl_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oor" {
        return &active.Oor
    }
    if childYangName == "list" {
        return &active.List
    }
    if childYangName == "oor-acls" {
        return &active.OorAcls
    }
    if childYangName == "usages" {
        return &active.Usages
    }
    return nil
}

func (active *EsAcl_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["oor"] = &active.Oor
    children["list"] = &active.List
    children["oor-acls"] = &active.OorAcls
    children["usages"] = &active.Usages
    return children
}

func (active *EsAcl_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (active *EsAcl_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *EsAcl_Active) GetYangName() string { return "active" }

func (active *EsAcl_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *EsAcl_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *EsAcl_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *EsAcl_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *EsAcl_Active) GetParent() types.Entity { return active.parent }

func (active *EsAcl_Active) GetParentYangName() string { return "es-acl" }

// EsAcl_Active_Oor
// Out Of Resources, Limits to the resources
// allocatable
type EsAcl_Active_Oor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource Limits pertaining to ACLs only.
    AclSummary EsAcl_Active_Oor_AclSummary
}

func (oor *EsAcl_Active_Oor) GetFilter() yfilter.YFilter { return oor.YFilter }

func (oor *EsAcl_Active_Oor) SetFilter(yf yfilter.YFilter) { oor.YFilter = yf }

func (oor *EsAcl_Active_Oor) GetGoName(yname string) string {
    if yname == "acl-summary" { return "AclSummary" }
    return ""
}

func (oor *EsAcl_Active_Oor) GetSegmentPath() string {
    return "oor"
}

func (oor *EsAcl_Active_Oor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl-summary" {
        return &oor.AclSummary
    }
    return nil
}

func (oor *EsAcl_Active_Oor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acl-summary"] = &oor.AclSummary
    return children
}

func (oor *EsAcl_Active_Oor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oor *EsAcl_Active_Oor) GetBundleName() string { return "cisco_ios_xr" }

func (oor *EsAcl_Active_Oor) GetYangName() string { return "oor" }

func (oor *EsAcl_Active_Oor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oor *EsAcl_Active_Oor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oor *EsAcl_Active_Oor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oor *EsAcl_Active_Oor) SetParent(parent types.Entity) { oor.parent = parent }

func (oor *EsAcl_Active_Oor) GetParent() types.Entity { return oor.parent }

func (oor *EsAcl_Active_Oor) GetParentYangName() string { return "active" }

// EsAcl_Active_Oor_AclSummary
// Resource Limits pertaining to ACLs only
type EsAcl_Active_Oor_AclSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details containing the resource limits of the ACLs.
    Details EsAcl_Active_Oor_AclSummary_Details
}

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetFilter() yfilter.YFilter { return aclSummary.YFilter }

func (aclSummary *EsAcl_Active_Oor_AclSummary) SetFilter(yf yfilter.YFilter) { aclSummary.YFilter = yf }

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetSegmentPath() string {
    return "acl-summary"
}

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &aclSummary.Details
    }
    return nil
}

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &aclSummary.Details
    return children
}

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetBundleName() string { return "cisco_ios_xr" }

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetYangName() string { return "acl-summary" }

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aclSummary *EsAcl_Active_Oor_AclSummary) SetParent(parent types.Entity) { aclSummary.parent = parent }

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetParent() types.Entity { return aclSummary.parent }

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetParentYangName() string { return "oor" }

// EsAcl_Active_Oor_AclSummary_Details
// Details containing the resource limits of the
// ACLs
type EsAcl_Active_Oor_AclSummary_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcEs interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    MaximumConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    MaximumConfigurableAcEs interface{}
}

func (details *EsAcl_Active_Oor_AclSummary_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *EsAcl_Active_Oor_AclSummary_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *EsAcl_Active_Oor_AclSummary_Details) GetGoName(yname string) string {
    if yname == "current-configured-ac-ls" { return "CurrentConfiguredAcLs" }
    if yname == "current-configured-ac-es" { return "CurrentConfiguredAcEs" }
    if yname == "maximum-configurable-ac-ls" { return "MaximumConfigurableAcLs" }
    if yname == "maximum-configurable-ac-es" { return "MaximumConfigurableAcEs" }
    return ""
}

func (details *EsAcl_Active_Oor_AclSummary_Details) GetSegmentPath() string {
    return "details"
}

func (details *EsAcl_Active_Oor_AclSummary_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (details *EsAcl_Active_Oor_AclSummary_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (details *EsAcl_Active_Oor_AclSummary_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-configured-ac-ls"] = details.CurrentConfiguredAcLs
    leafs["current-configured-ac-es"] = details.CurrentConfiguredAcEs
    leafs["maximum-configurable-ac-ls"] = details.MaximumConfigurableAcLs
    leafs["maximum-configurable-ac-es"] = details.MaximumConfigurableAcEs
    return leafs
}

func (details *EsAcl_Active_Oor_AclSummary_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *EsAcl_Active_Oor_AclSummary_Details) GetYangName() string { return "details" }

func (details *EsAcl_Active_Oor_AclSummary_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *EsAcl_Active_Oor_AclSummary_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *EsAcl_Active_Oor_AclSummary_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *EsAcl_Active_Oor_AclSummary_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *EsAcl_Active_Oor_AclSummary_Details) GetParent() types.Entity { return details.parent }

func (details *EsAcl_Active_Oor_AclSummary_Details) GetParentYangName() string { return "acl-summary" }

// EsAcl_Active_List
// List containing ACLs
type EsAcl_Active_List struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ACL class displaying Usage and Entries.
    Acls EsAcl_Active_List_Acls
}

func (list *EsAcl_Active_List) GetFilter() yfilter.YFilter { return list.YFilter }

func (list *EsAcl_Active_List) SetFilter(yf yfilter.YFilter) { list.YFilter = yf }

func (list *EsAcl_Active_List) GetGoName(yname string) string {
    if yname == "acls" { return "Acls" }
    return ""
}

func (list *EsAcl_Active_List) GetSegmentPath() string {
    return "list"
}

func (list *EsAcl_Active_List) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acls" {
        return &list.Acls
    }
    return nil
}

func (list *EsAcl_Active_List) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acls"] = &list.Acls
    return children
}

func (list *EsAcl_Active_List) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (list *EsAcl_Active_List) GetBundleName() string { return "cisco_ios_xr" }

func (list *EsAcl_Active_List) GetYangName() string { return "list" }

func (list *EsAcl_Active_List) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (list *EsAcl_Active_List) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (list *EsAcl_Active_List) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (list *EsAcl_Active_List) SetParent(parent types.Entity) { list.parent = parent }

func (list *EsAcl_Active_List) GetParent() types.Entity { return list.parent }

func (list *EsAcl_Active_List) GetParentYangName() string { return "active" }

// EsAcl_Active_List_Acls
// ACL class displaying Usage and Entries
type EsAcl_Active_List_Acls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the Access List. The type is slice of EsAcl_Active_List_Acls_Acl.
    Acl []EsAcl_Active_List_Acls_Acl
}

func (acls *EsAcl_Active_List_Acls) GetFilter() yfilter.YFilter { return acls.YFilter }

func (acls *EsAcl_Active_List_Acls) SetFilter(yf yfilter.YFilter) { acls.YFilter = yf }

func (acls *EsAcl_Active_List_Acls) GetGoName(yname string) string {
    if yname == "acl" { return "Acl" }
    return ""
}

func (acls *EsAcl_Active_List_Acls) GetSegmentPath() string {
    return "acls"
}

func (acls *EsAcl_Active_List_Acls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl" {
        for _, c := range acls.Acl {
            if acls.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EsAcl_Active_List_Acls_Acl{}
        acls.Acl = append(acls.Acl, child)
        return &acls.Acl[len(acls.Acl)-1]
    }
    return nil
}

func (acls *EsAcl_Active_List_Acls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range acls.Acl {
        children[acls.Acl[i].GetSegmentPath()] = &acls.Acl[i]
    }
    return children
}

func (acls *EsAcl_Active_List_Acls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (acls *EsAcl_Active_List_Acls) GetBundleName() string { return "cisco_ios_xr" }

func (acls *EsAcl_Active_List_Acls) GetYangName() string { return "acls" }

func (acls *EsAcl_Active_List_Acls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acls *EsAcl_Active_List_Acls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acls *EsAcl_Active_List_Acls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acls *EsAcl_Active_List_Acls) SetParent(parent types.Entity) { acls.parent = parent }

func (acls *EsAcl_Active_List_Acls) GetParent() types.Entity { return acls.parent }

func (acls *EsAcl_Active_List_Acls) GetParentYangName() string { return "list" }

// EsAcl_Active_List_Acls_Acl
// Name of the Access List
type EsAcl_Active_List_Acls_Acl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Access List. The type is string with
    // length: 1..65.
    Name interface{}

    // Table of all the SequenceNumbers per ACL.
    AclSequenceNumbers EsAcl_Active_List_Acls_Acl_AclSequenceNumbers
}

func (acl *EsAcl_Active_List_Acls_Acl) GetFilter() yfilter.YFilter { return acl.YFilter }

func (acl *EsAcl_Active_List_Acls_Acl) SetFilter(yf yfilter.YFilter) { acl.YFilter = yf }

func (acl *EsAcl_Active_List_Acls_Acl) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "acl-sequence-numbers" { return "AclSequenceNumbers" }
    return ""
}

func (acl *EsAcl_Active_List_Acls_Acl) GetSegmentPath() string {
    return "acl" + "[name='" + fmt.Sprintf("%v", acl.Name) + "']"
}

func (acl *EsAcl_Active_List_Acls_Acl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl-sequence-numbers" {
        return &acl.AclSequenceNumbers
    }
    return nil
}

func (acl *EsAcl_Active_List_Acls_Acl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["acl-sequence-numbers"] = &acl.AclSequenceNumbers
    return children
}

func (acl *EsAcl_Active_List_Acls_Acl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = acl.Name
    return leafs
}

func (acl *EsAcl_Active_List_Acls_Acl) GetBundleName() string { return "cisco_ios_xr" }

func (acl *EsAcl_Active_List_Acls_Acl) GetYangName() string { return "acl" }

func (acl *EsAcl_Active_List_Acls_Acl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acl *EsAcl_Active_List_Acls_Acl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acl *EsAcl_Active_List_Acls_Acl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acl *EsAcl_Active_List_Acls_Acl) SetParent(parent types.Entity) { acl.parent = parent }

func (acl *EsAcl_Active_List_Acls_Acl) GetParent() types.Entity { return acl.parent }

func (acl *EsAcl_Active_List_Acls_Acl) GetParentYangName() string { return "acls" }

// EsAcl_Active_List_Acls_Acl_AclSequenceNumbers
// Table of all the SequenceNumbers per ACL
type EsAcl_Active_List_Acls_Acl_AclSequenceNumbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sequence Number of an ACL entry. The type is slice of
    // EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber.
    AclSequenceNumber []EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber
}

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetFilter() yfilter.YFilter { return aclSequenceNumbers.YFilter }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) SetFilter(yf yfilter.YFilter) { aclSequenceNumbers.YFilter = yf }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetGoName(yname string) string {
    if yname == "acl-sequence-number" { return "AclSequenceNumber" }
    return ""
}

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetSegmentPath() string {
    return "acl-sequence-numbers"
}

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "acl-sequence-number" {
        for _, c := range aclSequenceNumbers.AclSequenceNumber {
            if aclSequenceNumbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber{}
        aclSequenceNumbers.AclSequenceNumber = append(aclSequenceNumbers.AclSequenceNumber, child)
        return &aclSequenceNumbers.AclSequenceNumber[len(aclSequenceNumbers.AclSequenceNumber)-1]
    }
    return nil
}

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range aclSequenceNumbers.AclSequenceNumber {
        children[aclSequenceNumbers.AclSequenceNumber[i].GetSegmentPath()] = &aclSequenceNumbers.AclSequenceNumber[i]
    }
    return children
}

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetBundleName() string { return "cisco_ios_xr" }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetYangName() string { return "acl-sequence-numbers" }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) SetParent(parent types.Entity) { aclSequenceNumbers.parent = parent }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetParent() types.Entity { return aclSequenceNumbers.parent }

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetParentYangName() string { return "acl" }

// EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber
// Sequence Number of an ACL entry
type EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ACLEntry Sequence Number. The type is interface{}
    // with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (acl, remark). The type is AclAce1.
    AceType interface{}

    // ACE sequence number. The type is interface{} with range: 0..4294967295.
    AceSequenceNumber interface{}

    // ACE hit number. The type is interface{} with range:
    // 0..18446744073709551615.
    Hits interface{}

    // Grant value permit/deny . The type is AclAction.
    Grant interface{}

    // Source MAC address. The type is string with pattern:
    // [a-fA-F0-9]{4}(\.[a-fA-F0-9]{4}){2}.
    SourceAddress interface{}

    // Source wild card bits. The type is string with pattern:
    // [a-fA-F0-9]{4}(\.[a-fA-F0-9]{4}){2}.
    SourceWildCardBits interface{}

    // Destination MAC address. The type is string with pattern:
    // [a-fA-F0-9]{4}(\.[a-fA-F0-9]{4}){2}.
    DestinationAddress interface{}

    // Destination wild card bits. The type is string with pattern:
    // [a-fA-F0-9]{4}(\.[a-fA-F0-9]{4}){2}.
    DestinationWildCardBits interface{}

    // Ethernet type Number. The type is interface{} with range: 0..65535.
    EtherTypeNumber interface{}

    // VLAN ID/range lower limit. The type is interface{} with range: 0..65535.
    Vlan1 interface{}

    // VLAN ID range higher limit. The type is interface{} with range: 0..65535.
    Vlan2 interface{}

    // COS value. The type is interface{} with range: 0..255.
    Cos interface{}

    // DEI bit. The type is interface{} with range: 0..255.
    Dei interface{}

    // Inner header VLAN ID/range lower limit. The type is interface{} with range:
    // 0..65535.
    InnerHeaderVlan1 interface{}

    // Inner header VLAN ID range higher limit. The type is interface{} with
    // range: 0..65535.
    InnerHeaderVlan2 interface{}

    // Inner header COS value. The type is interface{} with range: 0..255.
    InnerHeaderCos interface{}

    // Inner header DEI bit. The type is interface{} with range: 0..255.
    InnerHeaderDei interface{}

    // Capture option, TRUE if enabled. The type is bool.
    Capture interface{}

    // Log option. The type is interface{} with range: 0..255.
    LogOption interface{}

    // Remark string. The type is string.
    Remark interface{}

    // Acl Name. The type is string.
    AclName interface{}

    // Sequence Sring. The type is string.
    SequenceString interface{}
}

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetFilter() yfilter.YFilter { return aclSequenceNumber.YFilter }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) SetFilter(yf yfilter.YFilter) { aclSequenceNumber.YFilter = yf }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "ace-type" { return "AceType" }
    if yname == "ace-sequence-number" { return "AceSequenceNumber" }
    if yname == "hits" { return "Hits" }
    if yname == "grant" { return "Grant" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-wild-card-bits" { return "SourceWildCardBits" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-wild-card-bits" { return "DestinationWildCardBits" }
    if yname == "ether-type-number" { return "EtherTypeNumber" }
    if yname == "vlan1" { return "Vlan1" }
    if yname == "vlan2" { return "Vlan2" }
    if yname == "cos" { return "Cos" }
    if yname == "dei" { return "Dei" }
    if yname == "inner-header-vlan1" { return "InnerHeaderVlan1" }
    if yname == "inner-header-vlan2" { return "InnerHeaderVlan2" }
    if yname == "inner-header-cos" { return "InnerHeaderCos" }
    if yname == "inner-header-dei" { return "InnerHeaderDei" }
    if yname == "capture" { return "Capture" }
    if yname == "log-option" { return "LogOption" }
    if yname == "remark" { return "Remark" }
    if yname == "acl-name" { return "AclName" }
    if yname == "sequence-string" { return "SequenceString" }
    return ""
}

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetSegmentPath() string {
    return "acl-sequence-number" + "[sequence-number='" + fmt.Sprintf("%v", aclSequenceNumber.SequenceNumber) + "']"
}

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = aclSequenceNumber.SequenceNumber
    leafs["ace-type"] = aclSequenceNumber.AceType
    leafs["ace-sequence-number"] = aclSequenceNumber.AceSequenceNumber
    leafs["hits"] = aclSequenceNumber.Hits
    leafs["grant"] = aclSequenceNumber.Grant
    leafs["source-address"] = aclSequenceNumber.SourceAddress
    leafs["source-wild-card-bits"] = aclSequenceNumber.SourceWildCardBits
    leafs["destination-address"] = aclSequenceNumber.DestinationAddress
    leafs["destination-wild-card-bits"] = aclSequenceNumber.DestinationWildCardBits
    leafs["ether-type-number"] = aclSequenceNumber.EtherTypeNumber
    leafs["vlan1"] = aclSequenceNumber.Vlan1
    leafs["vlan2"] = aclSequenceNumber.Vlan2
    leafs["cos"] = aclSequenceNumber.Cos
    leafs["dei"] = aclSequenceNumber.Dei
    leafs["inner-header-vlan1"] = aclSequenceNumber.InnerHeaderVlan1
    leafs["inner-header-vlan2"] = aclSequenceNumber.InnerHeaderVlan2
    leafs["inner-header-cos"] = aclSequenceNumber.InnerHeaderCos
    leafs["inner-header-dei"] = aclSequenceNumber.InnerHeaderDei
    leafs["capture"] = aclSequenceNumber.Capture
    leafs["log-option"] = aclSequenceNumber.LogOption
    leafs["remark"] = aclSequenceNumber.Remark
    leafs["acl-name"] = aclSequenceNumber.AclName
    leafs["sequence-string"] = aclSequenceNumber.SequenceString
    return leafs
}

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetBundleName() string { return "cisco_ios_xr" }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetYangName() string { return "acl-sequence-number" }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) SetParent(parent types.Entity) { aclSequenceNumber.parent = parent }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetParent() types.Entity { return aclSequenceNumber.parent }

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetParentYangName() string { return "acl-sequence-numbers" }

// EsAcl_Active_OorAcls
// Resource occupation details for ACLs
type EsAcl_Active_OorAcls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource occupation details for a particular ACL. The type is slice of
    // EsAcl_Active_OorAcls_OorAcl.
    OorAcl []EsAcl_Active_OorAcls_OorAcl
}

func (oorAcls *EsAcl_Active_OorAcls) GetFilter() yfilter.YFilter { return oorAcls.YFilter }

func (oorAcls *EsAcl_Active_OorAcls) SetFilter(yf yfilter.YFilter) { oorAcls.YFilter = yf }

func (oorAcls *EsAcl_Active_OorAcls) GetGoName(yname string) string {
    if yname == "oor-acl" { return "OorAcl" }
    return ""
}

func (oorAcls *EsAcl_Active_OorAcls) GetSegmentPath() string {
    return "oor-acls"
}

func (oorAcls *EsAcl_Active_OorAcls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oor-acl" {
        for _, c := range oorAcls.OorAcl {
            if oorAcls.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EsAcl_Active_OorAcls_OorAcl{}
        oorAcls.OorAcl = append(oorAcls.OorAcl, child)
        return &oorAcls.OorAcl[len(oorAcls.OorAcl)-1]
    }
    return nil
}

func (oorAcls *EsAcl_Active_OorAcls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range oorAcls.OorAcl {
        children[oorAcls.OorAcl[i].GetSegmentPath()] = &oorAcls.OorAcl[i]
    }
    return children
}

func (oorAcls *EsAcl_Active_OorAcls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oorAcls *EsAcl_Active_OorAcls) GetBundleName() string { return "cisco_ios_xr" }

func (oorAcls *EsAcl_Active_OorAcls) GetYangName() string { return "oor-acls" }

func (oorAcls *EsAcl_Active_OorAcls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorAcls *EsAcl_Active_OorAcls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorAcls *EsAcl_Active_OorAcls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorAcls *EsAcl_Active_OorAcls) SetParent(parent types.Entity) { oorAcls.parent = parent }

func (oorAcls *EsAcl_Active_OorAcls) GetParent() types.Entity { return oorAcls.parent }

func (oorAcls *EsAcl_Active_OorAcls) GetParentYangName() string { return "active" }

// EsAcl_Active_OorAcls_OorAcl
// Resource occupation details for a particular
// ACL
type EsAcl_Active_OorAcls_OorAcl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Access List. The type is string with
    // length: 1..65.
    Name interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcEs interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    MaximumConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    MaximumConfigurableAcEs interface{}
}

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetFilter() yfilter.YFilter { return oorAcl.YFilter }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) SetFilter(yf yfilter.YFilter) { oorAcl.YFilter = yf }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "current-configured-ac-ls" { return "CurrentConfiguredAcLs" }
    if yname == "current-configured-ac-es" { return "CurrentConfiguredAcEs" }
    if yname == "maximum-configurable-ac-ls" { return "MaximumConfigurableAcLs" }
    if yname == "maximum-configurable-ac-es" { return "MaximumConfigurableAcEs" }
    return ""
}

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetSegmentPath() string {
    return "oor-acl" + "[name='" + fmt.Sprintf("%v", oorAcl.Name) + "']"
}

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = oorAcl.Name
    leafs["current-configured-ac-ls"] = oorAcl.CurrentConfiguredAcLs
    leafs["current-configured-ac-es"] = oorAcl.CurrentConfiguredAcEs
    leafs["maximum-configurable-ac-ls"] = oorAcl.MaximumConfigurableAcLs
    leafs["maximum-configurable-ac-es"] = oorAcl.MaximumConfigurableAcEs
    return leafs
}

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetBundleName() string { return "cisco_ios_xr" }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetYangName() string { return "oor-acl" }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) SetParent(parent types.Entity) { oorAcl.parent = parent }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetParent() types.Entity { return oorAcl.parent }

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetParentYangName() string { return "oor-acls" }

// EsAcl_Active_Usages
// Table of Usage statistics of ACLs at different
// nodes
type EsAcl_Active_Usages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Usage statistics of an ACL at a node. The type is slice of
    // EsAcl_Active_Usages_Usage.
    Usage []EsAcl_Active_Usages_Usage
}

func (usages *EsAcl_Active_Usages) GetFilter() yfilter.YFilter { return usages.YFilter }

func (usages *EsAcl_Active_Usages) SetFilter(yf yfilter.YFilter) { usages.YFilter = yf }

func (usages *EsAcl_Active_Usages) GetGoName(yname string) string {
    if yname == "usage" { return "Usage" }
    return ""
}

func (usages *EsAcl_Active_Usages) GetSegmentPath() string {
    return "usages"
}

func (usages *EsAcl_Active_Usages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usage" {
        for _, c := range usages.Usage {
            if usages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EsAcl_Active_Usages_Usage{}
        usages.Usage = append(usages.Usage, child)
        return &usages.Usage[len(usages.Usage)-1]
    }
    return nil
}

func (usages *EsAcl_Active_Usages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usages.Usage {
        children[usages.Usage[i].GetSegmentPath()] = &usages.Usage[i]
    }
    return children
}

func (usages *EsAcl_Active_Usages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usages *EsAcl_Active_Usages) GetBundleName() string { return "cisco_ios_xr" }

func (usages *EsAcl_Active_Usages) GetYangName() string { return "usages" }

func (usages *EsAcl_Active_Usages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usages *EsAcl_Active_Usages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usages *EsAcl_Active_Usages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usages *EsAcl_Active_Usages) SetParent(parent types.Entity) { usages.parent = parent }

func (usages *EsAcl_Active_Usages) GetParent() types.Entity { return usages.parent }

func (usages *EsAcl_Active_Usages) GetParentYangName() string { return "active" }

// EsAcl_Active_Usages_Usage
// Usage statistics of an ACL at a node
type EsAcl_Active_Usages_Usage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node where ACL is applied. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Application ID. The type is AclUsageAppIdEnum.
    ApplicationId interface{}

    // Name of the ACL. The type is string with length: 1..65.
    Name interface{}

    // Usage Statistics Details. The type is string. This attribute is mandatory.
    UsageDetails interface{}
}

func (usage *EsAcl_Active_Usages_Usage) GetFilter() yfilter.YFilter { return usage.YFilter }

func (usage *EsAcl_Active_Usages_Usage) SetFilter(yf yfilter.YFilter) { usage.YFilter = yf }

func (usage *EsAcl_Active_Usages_Usage) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "application-id" { return "ApplicationId" }
    if yname == "name" { return "Name" }
    if yname == "usage-details" { return "UsageDetails" }
    return ""
}

func (usage *EsAcl_Active_Usages_Usage) GetSegmentPath() string {
    return "usage"
}

func (usage *EsAcl_Active_Usages_Usage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usage *EsAcl_Active_Usages_Usage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usage *EsAcl_Active_Usages_Usage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = usage.Location
    leafs["application-id"] = usage.ApplicationId
    leafs["name"] = usage.Name
    leafs["usage-details"] = usage.UsageDetails
    return leafs
}

func (usage *EsAcl_Active_Usages_Usage) GetBundleName() string { return "cisco_ios_xr" }

func (usage *EsAcl_Active_Usages_Usage) GetYangName() string { return "usage" }

func (usage *EsAcl_Active_Usages_Usage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usage *EsAcl_Active_Usages_Usage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usage *EsAcl_Active_Usages_Usage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usage *EsAcl_Active_Usages_Usage) SetParent(parent types.Entity) { usage.parent = parent }

func (usage *EsAcl_Active_Usages_Usage) GetParent() types.Entity { return usage.parent }

func (usage *EsAcl_Active_Usages_Usage) GetParentYangName() string { return "usages" }

