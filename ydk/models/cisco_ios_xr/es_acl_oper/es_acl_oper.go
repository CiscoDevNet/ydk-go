// This module contains a collection of YANG definitions
// for Cisco IOS-XR es-acl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   es-acl: Root class of ES ACL Oper schema tree
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// AclAce1_ represents ACE Types
type AclAce1_ string

const (
    // This is Normal ACE
    AclAce1__normal AclAce1_ = "normal"

    // This is Remark ACE
    AclAce1__remark AclAce1_ = "remark"

    // This is ABF ACE
    AclAce1__abf AclAce1_ = "abf"
)

// EsAcl
// Root class of ES ACL Oper schema tree
type EsAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Out Of Resources, Limits to the resources allocatable.
    Active EsAcl_Active
}

func (esAcl *EsAcl) GetEntityData() *types.CommonEntityData {
    esAcl.EntityData.YFilter = esAcl.YFilter
    esAcl.EntityData.YangName = "es-acl"
    esAcl.EntityData.BundleName = "cisco_ios_xr"
    esAcl.EntityData.ParentYangName = "Cisco-IOS-XR-es-acl-oper"
    esAcl.EntityData.SegmentPath = "Cisco-IOS-XR-es-acl-oper:es-acl"
    esAcl.EntityData.AbsolutePath = esAcl.EntityData.SegmentPath
    esAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esAcl.EntityData.Children = types.NewOrderedMap()
    esAcl.EntityData.Children.Append("active", types.YChild{"Active", &esAcl.Active})
    esAcl.EntityData.Leafs = types.NewOrderedMap()

    esAcl.EntityData.YListKeys = []string {}

    return &(esAcl.EntityData)
}

// EsAcl_Active
// Out Of Resources, Limits to the resources
// allocatable
type EsAcl_Active struct {
    EntityData types.CommonEntityData
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

func (active *EsAcl_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "es-acl"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("oor", types.YChild{"Oor", &active.Oor})
    active.EntityData.Children.Append("list", types.YChild{"List", &active.List})
    active.EntityData.Children.Append("oor-acls", types.YChild{"OorAcls", &active.OorAcls})
    active.EntityData.Children.Append("usages", types.YChild{"Usages", &active.Usages})
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// EsAcl_Active_Oor
// Out Of Resources, Limits to the resources
// allocatable
type EsAcl_Active_Oor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource Limits pertaining to ACLs only.
    AclSummary EsAcl_Active_Oor_AclSummary
}

func (oor *EsAcl_Active_Oor) GetEntityData() *types.CommonEntityData {
    oor.EntityData.YFilter = oor.YFilter
    oor.EntityData.YangName = "oor"
    oor.EntityData.BundleName = "cisco_ios_xr"
    oor.EntityData.ParentYangName = "active"
    oor.EntityData.SegmentPath = "oor"
    oor.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/" + oor.EntityData.SegmentPath
    oor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oor.EntityData.Children = types.NewOrderedMap()
    oor.EntityData.Children.Append("acl-summary", types.YChild{"AclSummary", &oor.AclSummary})
    oor.EntityData.Leafs = types.NewOrderedMap()

    oor.EntityData.YListKeys = []string {}

    return &(oor.EntityData)
}

// EsAcl_Active_Oor_AclSummary
// Resource Limits pertaining to ACLs only
type EsAcl_Active_Oor_AclSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Details containing the resource limits of the ACLs.
    Details EsAcl_Active_Oor_AclSummary_Details
}

func (aclSummary *EsAcl_Active_Oor_AclSummary) GetEntityData() *types.CommonEntityData {
    aclSummary.EntityData.YFilter = aclSummary.YFilter
    aclSummary.EntityData.YangName = "acl-summary"
    aclSummary.EntityData.BundleName = "cisco_ios_xr"
    aclSummary.EntityData.ParentYangName = "oor"
    aclSummary.EntityData.SegmentPath = "acl-summary"
    aclSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/oor/" + aclSummary.EntityData.SegmentPath
    aclSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclSummary.EntityData.Children = types.NewOrderedMap()
    aclSummary.EntityData.Children.Append("details", types.YChild{"Details", &aclSummary.Details})
    aclSummary.EntityData.Leafs = types.NewOrderedMap()

    aclSummary.EntityData.YListKeys = []string {}

    return &(aclSummary.EntityData)
}

// EsAcl_Active_Oor_AclSummary_Details
// Details containing the resource limits of the
// ACLs
type EsAcl_Active_Oor_AclSummary_Details struct {
    EntityData types.CommonEntityData
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

func (details *EsAcl_Active_Oor_AclSummary_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "acl-summary"
    details.EntityData.SegmentPath = "details"
    details.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/oor/acl-summary/" + details.EntityData.SegmentPath
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = types.NewOrderedMap()
    details.EntityData.Leafs = types.NewOrderedMap()
    details.EntityData.Leafs.Append("current-configured-ac-ls", types.YLeaf{"CurrentConfiguredAcLs", details.CurrentConfiguredAcLs})
    details.EntityData.Leafs.Append("current-configured-ac-es", types.YLeaf{"CurrentConfiguredAcEs", details.CurrentConfiguredAcEs})
    details.EntityData.Leafs.Append("maximum-configurable-ac-ls", types.YLeaf{"MaximumConfigurableAcLs", details.MaximumConfigurableAcLs})
    details.EntityData.Leafs.Append("maximum-configurable-ac-es", types.YLeaf{"MaximumConfigurableAcEs", details.MaximumConfigurableAcEs})

    details.EntityData.YListKeys = []string {}

    return &(details.EntityData)
}

// EsAcl_Active_List
// List containing ACLs
type EsAcl_Active_List struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ACL class displaying Usage and Entries.
    Acls EsAcl_Active_List_Acls
}

func (list *EsAcl_Active_List) GetEntityData() *types.CommonEntityData {
    list.EntityData.YFilter = list.YFilter
    list.EntityData.YangName = "list"
    list.EntityData.BundleName = "cisco_ios_xr"
    list.EntityData.ParentYangName = "active"
    list.EntityData.SegmentPath = "list"
    list.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/" + list.EntityData.SegmentPath
    list.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    list.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    list.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    list.EntityData.Children = types.NewOrderedMap()
    list.EntityData.Children.Append("acls", types.YChild{"Acls", &list.Acls})
    list.EntityData.Leafs = types.NewOrderedMap()

    list.EntityData.YListKeys = []string {}

    return &(list.EntityData)
}

// EsAcl_Active_List_Acls
// ACL class displaying Usage and Entries
type EsAcl_Active_List_Acls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Access List. The type is slice of EsAcl_Active_List_Acls_Acl.
    Acl []*EsAcl_Active_List_Acls_Acl
}

func (acls *EsAcl_Active_List_Acls) GetEntityData() *types.CommonEntityData {
    acls.EntityData.YFilter = acls.YFilter
    acls.EntityData.YangName = "acls"
    acls.EntityData.BundleName = "cisco_ios_xr"
    acls.EntityData.ParentYangName = "list"
    acls.EntityData.SegmentPath = "acls"
    acls.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/list/" + acls.EntityData.SegmentPath
    acls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acls.EntityData.Children = types.NewOrderedMap()
    acls.EntityData.Children.Append("acl", types.YChild{"Acl", nil})
    for i := range acls.Acl {
        acls.EntityData.Children.Append(types.GetSegmentPath(acls.Acl[i]), types.YChild{"Acl", acls.Acl[i]})
    }
    acls.EntityData.Leafs = types.NewOrderedMap()

    acls.EntityData.YListKeys = []string {}

    return &(acls.EntityData)
}

// EsAcl_Active_List_Acls_Acl
// Name of the Access List
type EsAcl_Active_List_Acls_Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the Access List. The type is string with
    // length: 1..64.
    Name interface{}

    // Table of all the SequenceNumbers per ACL.
    AclSequenceNumbers EsAcl_Active_List_Acls_Acl_AclSequenceNumbers
}

func (acl *EsAcl_Active_List_Acls_Acl) GetEntityData() *types.CommonEntityData {
    acl.EntityData.YFilter = acl.YFilter
    acl.EntityData.YangName = "acl"
    acl.EntityData.BundleName = "cisco_ios_xr"
    acl.EntityData.ParentYangName = "acls"
    acl.EntityData.SegmentPath = "acl" + types.AddKeyToken(acl.Name, "name")
    acl.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/list/acls/" + acl.EntityData.SegmentPath
    acl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acl.EntityData.Children = types.NewOrderedMap()
    acl.EntityData.Children.Append("acl-sequence-numbers", types.YChild{"AclSequenceNumbers", &acl.AclSequenceNumbers})
    acl.EntityData.Leafs = types.NewOrderedMap()
    acl.EntityData.Leafs.Append("name", types.YLeaf{"Name", acl.Name})

    acl.EntityData.YListKeys = []string {"Name"}

    return &(acl.EntityData)
}

// EsAcl_Active_List_Acls_Acl_AclSequenceNumbers
// Table of all the SequenceNumbers per ACL
type EsAcl_Active_List_Acls_Acl_AclSequenceNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequence Number of an ACL entry. The type is slice of
    // EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber.
    AclSequenceNumber []*EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber
}

func (aclSequenceNumbers *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers) GetEntityData() *types.CommonEntityData {
    aclSequenceNumbers.EntityData.YFilter = aclSequenceNumbers.YFilter
    aclSequenceNumbers.EntityData.YangName = "acl-sequence-numbers"
    aclSequenceNumbers.EntityData.BundleName = "cisco_ios_xr"
    aclSequenceNumbers.EntityData.ParentYangName = "acl"
    aclSequenceNumbers.EntityData.SegmentPath = "acl-sequence-numbers"
    aclSequenceNumbers.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/list/acls/acl/" + aclSequenceNumbers.EntityData.SegmentPath
    aclSequenceNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclSequenceNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclSequenceNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclSequenceNumbers.EntityData.Children = types.NewOrderedMap()
    aclSequenceNumbers.EntityData.Children.Append("acl-sequence-number", types.YChild{"AclSequenceNumber", nil})
    for i := range aclSequenceNumbers.AclSequenceNumber {
        aclSequenceNumbers.EntityData.Children.Append(types.GetSegmentPath(aclSequenceNumbers.AclSequenceNumber[i]), types.YChild{"AclSequenceNumber", aclSequenceNumbers.AclSequenceNumber[i]})
    }
    aclSequenceNumbers.EntityData.Leafs = types.NewOrderedMap()

    aclSequenceNumbers.EntityData.YListKeys = []string {}

    return &(aclSequenceNumbers.EntityData)
}

// EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber
// Sequence Number of an ACL entry
type EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. ACLEntry Sequence Number. The type is interface{}
    // with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (acl, remark). The type is AclAce1_.
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

func (aclSequenceNumber *EsAcl_Active_List_Acls_Acl_AclSequenceNumbers_AclSequenceNumber) GetEntityData() *types.CommonEntityData {
    aclSequenceNumber.EntityData.YFilter = aclSequenceNumber.YFilter
    aclSequenceNumber.EntityData.YangName = "acl-sequence-number"
    aclSequenceNumber.EntityData.BundleName = "cisco_ios_xr"
    aclSequenceNumber.EntityData.ParentYangName = "acl-sequence-numbers"
    aclSequenceNumber.EntityData.SegmentPath = "acl-sequence-number" + types.AddKeyToken(aclSequenceNumber.SequenceNumber, "sequence-number")
    aclSequenceNumber.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/list/acls/acl/acl-sequence-numbers/" + aclSequenceNumber.EntityData.SegmentPath
    aclSequenceNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aclSequenceNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aclSequenceNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aclSequenceNumber.EntityData.Children = types.NewOrderedMap()
    aclSequenceNumber.EntityData.Leafs = types.NewOrderedMap()
    aclSequenceNumber.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", aclSequenceNumber.SequenceNumber})
    aclSequenceNumber.EntityData.Leafs.Append("ace-type", types.YLeaf{"AceType", aclSequenceNumber.AceType})
    aclSequenceNumber.EntityData.Leafs.Append("ace-sequence-number", types.YLeaf{"AceSequenceNumber", aclSequenceNumber.AceSequenceNumber})
    aclSequenceNumber.EntityData.Leafs.Append("hits", types.YLeaf{"Hits", aclSequenceNumber.Hits})
    aclSequenceNumber.EntityData.Leafs.Append("grant", types.YLeaf{"Grant", aclSequenceNumber.Grant})
    aclSequenceNumber.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", aclSequenceNumber.SourceAddress})
    aclSequenceNumber.EntityData.Leafs.Append("source-wild-card-bits", types.YLeaf{"SourceWildCardBits", aclSequenceNumber.SourceWildCardBits})
    aclSequenceNumber.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", aclSequenceNumber.DestinationAddress})
    aclSequenceNumber.EntityData.Leafs.Append("destination-wild-card-bits", types.YLeaf{"DestinationWildCardBits", aclSequenceNumber.DestinationWildCardBits})
    aclSequenceNumber.EntityData.Leafs.Append("ether-type-number", types.YLeaf{"EtherTypeNumber", aclSequenceNumber.EtherTypeNumber})
    aclSequenceNumber.EntityData.Leafs.Append("vlan1", types.YLeaf{"Vlan1", aclSequenceNumber.Vlan1})
    aclSequenceNumber.EntityData.Leafs.Append("vlan2", types.YLeaf{"Vlan2", aclSequenceNumber.Vlan2})
    aclSequenceNumber.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", aclSequenceNumber.Cos})
    aclSequenceNumber.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", aclSequenceNumber.Dei})
    aclSequenceNumber.EntityData.Leafs.Append("inner-header-vlan1", types.YLeaf{"InnerHeaderVlan1", aclSequenceNumber.InnerHeaderVlan1})
    aclSequenceNumber.EntityData.Leafs.Append("inner-header-vlan2", types.YLeaf{"InnerHeaderVlan2", aclSequenceNumber.InnerHeaderVlan2})
    aclSequenceNumber.EntityData.Leafs.Append("inner-header-cos", types.YLeaf{"InnerHeaderCos", aclSequenceNumber.InnerHeaderCos})
    aclSequenceNumber.EntityData.Leafs.Append("inner-header-dei", types.YLeaf{"InnerHeaderDei", aclSequenceNumber.InnerHeaderDei})
    aclSequenceNumber.EntityData.Leafs.Append("capture", types.YLeaf{"Capture", aclSequenceNumber.Capture})
    aclSequenceNumber.EntityData.Leafs.Append("log-option", types.YLeaf{"LogOption", aclSequenceNumber.LogOption})
    aclSequenceNumber.EntityData.Leafs.Append("remark", types.YLeaf{"Remark", aclSequenceNumber.Remark})
    aclSequenceNumber.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", aclSequenceNumber.AclName})
    aclSequenceNumber.EntityData.Leafs.Append("sequence-string", types.YLeaf{"SequenceString", aclSequenceNumber.SequenceString})

    aclSequenceNumber.EntityData.YListKeys = []string {"SequenceNumber"}

    return &(aclSequenceNumber.EntityData)
}

// EsAcl_Active_OorAcls
// Resource occupation details for ACLs
type EsAcl_Active_OorAcls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource occupation details for a particular ACL. The type is slice of
    // EsAcl_Active_OorAcls_OorAcl.
    OorAcl []*EsAcl_Active_OorAcls_OorAcl
}

func (oorAcls *EsAcl_Active_OorAcls) GetEntityData() *types.CommonEntityData {
    oorAcls.EntityData.YFilter = oorAcls.YFilter
    oorAcls.EntityData.YangName = "oor-acls"
    oorAcls.EntityData.BundleName = "cisco_ios_xr"
    oorAcls.EntityData.ParentYangName = "active"
    oorAcls.EntityData.SegmentPath = "oor-acls"
    oorAcls.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/" + oorAcls.EntityData.SegmentPath
    oorAcls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorAcls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorAcls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorAcls.EntityData.Children = types.NewOrderedMap()
    oorAcls.EntityData.Children.Append("oor-acl", types.YChild{"OorAcl", nil})
    for i := range oorAcls.OorAcl {
        oorAcls.EntityData.Children.Append(types.GetSegmentPath(oorAcls.OorAcl[i]), types.YChild{"OorAcl", oorAcls.OorAcl[i]})
    }
    oorAcls.EntityData.Leafs = types.NewOrderedMap()

    oorAcls.EntityData.YListKeys = []string {}

    return &(oorAcls.EntityData)
}

// EsAcl_Active_OorAcls_OorAcl
// Resource occupation details for a particular
// ACL
type EsAcl_Active_OorAcls_OorAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the Access List. The type is string with
    // length: 1..64.
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

func (oorAcl *EsAcl_Active_OorAcls_OorAcl) GetEntityData() *types.CommonEntityData {
    oorAcl.EntityData.YFilter = oorAcl.YFilter
    oorAcl.EntityData.YangName = "oor-acl"
    oorAcl.EntityData.BundleName = "cisco_ios_xr"
    oorAcl.EntityData.ParentYangName = "oor-acls"
    oorAcl.EntityData.SegmentPath = "oor-acl" + types.AddKeyToken(oorAcl.Name, "name")
    oorAcl.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/oor-acls/" + oorAcl.EntityData.SegmentPath
    oorAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorAcl.EntityData.Children = types.NewOrderedMap()
    oorAcl.EntityData.Leafs = types.NewOrderedMap()
    oorAcl.EntityData.Leafs.Append("name", types.YLeaf{"Name", oorAcl.Name})
    oorAcl.EntityData.Leafs.Append("current-configured-ac-ls", types.YLeaf{"CurrentConfiguredAcLs", oorAcl.CurrentConfiguredAcLs})
    oorAcl.EntityData.Leafs.Append("current-configured-ac-es", types.YLeaf{"CurrentConfiguredAcEs", oorAcl.CurrentConfiguredAcEs})
    oorAcl.EntityData.Leafs.Append("maximum-configurable-ac-ls", types.YLeaf{"MaximumConfigurableAcLs", oorAcl.MaximumConfigurableAcLs})
    oorAcl.EntityData.Leafs.Append("maximum-configurable-ac-es", types.YLeaf{"MaximumConfigurableAcEs", oorAcl.MaximumConfigurableAcEs})

    oorAcl.EntityData.YListKeys = []string {"Name"}

    return &(oorAcl.EntityData)
}

// EsAcl_Active_Usages
// Table of Usage statistics of ACLs at different
// nodes
type EsAcl_Active_Usages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Usage statistics of an ACL at a node. The type is slice of
    // EsAcl_Active_Usages_Usage.
    Usage []*EsAcl_Active_Usages_Usage
}

func (usages *EsAcl_Active_Usages) GetEntityData() *types.CommonEntityData {
    usages.EntityData.YFilter = usages.YFilter
    usages.EntityData.YangName = "usages"
    usages.EntityData.BundleName = "cisco_ios_xr"
    usages.EntityData.ParentYangName = "active"
    usages.EntityData.SegmentPath = "usages"
    usages.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/" + usages.EntityData.SegmentPath
    usages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usages.EntityData.Children = types.NewOrderedMap()
    usages.EntityData.Children.Append("usage", types.YChild{"Usage", nil})
    for i := range usages.Usage {
        types.SetYListKey(usages.Usage[i], i)
        usages.EntityData.Children.Append(types.GetSegmentPath(usages.Usage[i]), types.YChild{"Usage", usages.Usage[i]})
    }
    usages.EntityData.Leafs = types.NewOrderedMap()

    usages.EntityData.YListKeys = []string {}

    return &(usages.EntityData)
}

// EsAcl_Active_Usages_Usage
// Usage statistics of an ACL at a node
type EsAcl_Active_Usages_Usage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Node where ACL is applied. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Application ID. The type is AclUsageAppIdEnum.
    ApplicationId interface{}

    // Name of the ACL. The type is string with length: 1..64.
    Name interface{}

    // Usage Statistics Details. The type is string. This attribute is mandatory.
    UsageDetails interface{}
}

func (usage *EsAcl_Active_Usages_Usage) GetEntityData() *types.CommonEntityData {
    usage.EntityData.YFilter = usage.YFilter
    usage.EntityData.YangName = "usage"
    usage.EntityData.BundleName = "cisco_ios_xr"
    usage.EntityData.ParentYangName = "usages"
    usage.EntityData.SegmentPath = "usage" + types.AddNoKeyToken(usage)
    usage.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-oper:es-acl/active/usages/" + usage.EntityData.SegmentPath
    usage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usage.EntityData.Children = types.NewOrderedMap()
    usage.EntityData.Leafs = types.NewOrderedMap()
    usage.EntityData.Leafs.Append("location", types.YLeaf{"Location", usage.Location})
    usage.EntityData.Leafs.Append("application-id", types.YLeaf{"ApplicationId", usage.ApplicationId})
    usage.EntityData.Leafs.Append("name", types.YLeaf{"Name", usage.Name})
    usage.EntityData.Leafs.Append("usage-details", types.YLeaf{"UsageDetails", usage.UsageDetails})

    usage.EntityData.YListKeys = []string {}

    return &(usage.EntityData)
}

