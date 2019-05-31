// This module contains a collection of YANG definitions
// for Cisco IOS-XR segment-routing-srv6 package operational data.
// 
// This module contains definitions
// for the following management objects:
//   srv6: Segment Routing with IPv6 dataplane
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package segment_routing_srv6_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package segment_routing_srv6_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-segment-routing-srv6-oper srv6}", reflect.TypeOf(Srv6{}))
    ydk.RegisterEntity("Cisco-IOS-XR-segment-routing-srv6-oper:srv6", reflect.TypeOf(Srv6{}))
}

// SidState represents SID manager SID state
type SidState string

const (
    // Unknown
    SidState_unknown SidState = "unknown"

    // In Use
    SidState_in_use SidState = "in-use"

    // Pending
    SidState_pending SidState = "pending"

    // Stale
    SidState_stale SidState = "stale"
)

// Srv6EndFunction represents SRv6 End Function Type
type Srv6EndFunction string

const (
    // Unknown
    Srv6EndFunction_unknown Srv6EndFunction = "unknown"

    // End (no PSP/USP)
    Srv6EndFunction_end Srv6EndFunction = "end"

    // End with PSP
    Srv6EndFunction_end_with_psp Srv6EndFunction = "end-with-psp"

    // End with USP
    Srv6EndFunction_end_with_usp Srv6EndFunction = "end-with-usp"

    // End with PSP/USP
    Srv6EndFunction_end_with_psp_usp Srv6EndFunction = "end-with-psp-usp"

    // End.X (no PSP/USP)
    Srv6EndFunction_end_x Srv6EndFunction = "end-x"

    // End.X with PSP
    Srv6EndFunction_end_x_with_psp Srv6EndFunction = "end-x-with-psp"

    // End.X with USP
    Srv6EndFunction_end_x_with_usp Srv6EndFunction = "end-x-with-usp"

    // End.X with PSP/USP
    Srv6EndFunction_end_x_with_psp_usp Srv6EndFunction = "end-x-with-psp-usp"

    // End.T (no PSP/USP)
    Srv6EndFunction_end_tbl Srv6EndFunction = "end-tbl"

    // End.T with PSP
    Srv6EndFunction_end_tbl_with_psp Srv6EndFunction = "end-tbl-with-psp"

    // End.T with USP
    Srv6EndFunction_end_tbl_with_usp Srv6EndFunction = "end-tbl-with-usp"

    // End.T with PSP/USP
    Srv6EndFunction_end_tbl_with_psp_usp Srv6EndFunction = "end-tbl-with-psp-usp"

    // End.B6
    Srv6EndFunction_end_b6 Srv6EndFunction = "end-b6"

    // End.B6.Encaps
    Srv6EndFunction_end_b6_encaps Srv6EndFunction = "end-b6-encaps"

    // End.BM
    Srv6EndFunction_end_bm Srv6EndFunction = "end-bm"

    // End.DX6
    Srv6EndFunction_end_dx6 Srv6EndFunction = "end-dx6"

    // End.DX4
    Srv6EndFunction_end_dx4 Srv6EndFunction = "end-dx4"

    // End.DT6
    Srv6EndFunction_end_dt6 Srv6EndFunction = "end-dt6"

    // End.DT4
    Srv6EndFunction_end_dt4 Srv6EndFunction = "end-dt4"

    // End.DT46
    Srv6EndFunction_end_dt46 Srv6EndFunction = "end-dt46"

    // End.DX2
    Srv6EndFunction_end_dx2 Srv6EndFunction = "end-dx2"

    // End.DX2V
    Srv6EndFunction_end_dx2v Srv6EndFunction = "end-dx2v"

    // End.DX2U
    Srv6EndFunction_end_dx2u Srv6EndFunction = "end-dx2u"

    // End.DX2M
    Srv6EndFunction_end_dx2m Srv6EndFunction = "end-dx2m"

    // End.OTP
    Srv6EndFunction_end_otp Srv6EndFunction = "end-otp"

    // End.S
    Srv6EndFunction_end_s Srv6EndFunction = "end-s"
)

// SidAllocation represents SID allocation type
type SidAllocation string

const (
    // Unknown
    SidAllocation_unknown SidAllocation = "unknown"

    // Dynamic
    SidAllocation_dynamic SidAllocation = "dynamic"

    // Explicit
    SidAllocation_explicit SidAllocation = "explicit"
)

// Srv6OutOfResourceState represents SRv6 Out of Resource State
type Srv6OutOfResourceState string

const (
    // Resources Available
    Srv6OutOfResourceState_oor_green Srv6OutOfResourceState = "oor-green"

    // Resources Warning. Have exceeded minor
    // threshold
    Srv6OutOfResourceState_oor_yellow Srv6OutOfResourceState = "oor-yellow"

    // Out of Resources. Have exceeded major threshold
    Srv6OutOfResourceState_oor_red Srv6OutOfResourceState = "oor-red"
)

// Srv6
// Segment Routing with IPv6 dataplane
type Srv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active SRv6 operational data.
    Active Srv6_Active

    // Standby SRv6 operational data.
    Standby Srv6_Standby
}

func (srv6 *Srv6) GetEntityData() *types.CommonEntityData {
    srv6.EntityData.YFilter = srv6.YFilter
    srv6.EntityData.YangName = "srv6"
    srv6.EntityData.BundleName = "cisco_ios_xr"
    srv6.EntityData.ParentYangName = "Cisco-IOS-XR-segment-routing-srv6-oper"
    srv6.EntityData.SegmentPath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6"
    srv6.EntityData.AbsolutePath = srv6.EntityData.SegmentPath
    srv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srv6.EntityData.Children = types.NewOrderedMap()
    srv6.EntityData.Children.Append("active", types.YChild{"Active", &srv6.Active})
    srv6.EntityData.Children.Append("standby", types.YChild{"Standby", &srv6.Standby})
    srv6.EntityData.Leafs = types.NewOrderedMap()

    srv6.EntityData.YListKeys = []string {}

    return &(srv6.EntityData)
}

// Srv6_Active
// Active SRv6 operational data
type Srv6_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational container for all Stale SIDs across all Locators.
    LocatorAllStaleSids Srv6_Active_LocatorAllStaleSids

    // SID Manager information.
    Manager Srv6_Active_Manager

    // SRv6 locators related information.
    Locators Srv6_Active_Locators

    // Operational container for all (Active and Stale) SIDs across all Locators.
    LocatorAllSids Srv6_Active_LocatorAllSids

    // Operational container for Active SIDs across all Locators.
    LocatorAllActiveSids Srv6_Active_LocatorAllActiveSids
}

func (active *Srv6_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "srv6"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("locator-all-stale-sids", types.YChild{"LocatorAllStaleSids", &active.LocatorAllStaleSids})
    active.EntityData.Children.Append("manager", types.YChild{"Manager", &active.Manager})
    active.EntityData.Children.Append("locators", types.YChild{"Locators", &active.Locators})
    active.EntityData.Children.Append("locator-all-sids", types.YChild{"LocatorAllSids", &active.LocatorAllSids})
    active.EntityData.Children.Append("locator-all-active-sids", types.YChild{"LocatorAllActiveSids", &active.LocatorAllActiveSids})
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// Srv6_Active_LocatorAllStaleSids
// Operational container for all Stale SIDs across
// all Locators
type Srv6_Active_LocatorAllStaleSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given locator and SID opcode. The type is slice of
    // Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid.
    LocatorAllStaleSid []*Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid
}

func (locatorAllStaleSids *Srv6_Active_LocatorAllStaleSids) GetEntityData() *types.CommonEntityData {
    locatorAllStaleSids.EntityData.YFilter = locatorAllStaleSids.YFilter
    locatorAllStaleSids.EntityData.YangName = "locator-all-stale-sids"
    locatorAllStaleSids.EntityData.BundleName = "cisco_ios_xr"
    locatorAllStaleSids.EntityData.ParentYangName = "active"
    locatorAllStaleSids.EntityData.SegmentPath = "locator-all-stale-sids"
    locatorAllStaleSids.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/" + locatorAllStaleSids.EntityData.SegmentPath
    locatorAllStaleSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllStaleSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllStaleSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllStaleSids.EntityData.Children = types.NewOrderedMap()
    locatorAllStaleSids.EntityData.Children.Append("locator-all-stale-sid", types.YChild{"LocatorAllStaleSid", nil})
    for i := range locatorAllStaleSids.LocatorAllStaleSid {
        locatorAllStaleSids.EntityData.Children.Append(types.GetSegmentPath(locatorAllStaleSids.LocatorAllStaleSid[i]), types.YChild{"LocatorAllStaleSid", locatorAllStaleSids.LocatorAllStaleSid[i]})
    }
    locatorAllStaleSids.EntityData.Leafs = types.NewOrderedMap()

    locatorAllStaleSids.EntityData.YListKeys = []string {}

    return &(locatorAllStaleSids.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid
// Operational data for given locator and SID
// opcode
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Locator name. The type is string with length:
    // 1..58.
    LocatorName interface{}

    // This attribute is a key. Sid opcode. The type is interface{} with range:
    // 0..4294967295.
    SidOpcode interface{}

    // SID. The type is string.
    Sid interface{}

    // Allocation Type. The type is SidAllocation.
    AllocationType interface{}

    // Function Type. The type is Srv6EndFunction.
    FunctionType interface{}

    // State. The type is SidState.
    State interface{}

    // Rewrite done or not. The type is bool.
    HasForwarding interface{}

    // Associated locator. The type is string.
    Locator interface{}

    // SID Context.
    SidContext Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext

    // Creation timestamp.
    CreateTimestamp Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_CreateTimestamp

    // Owner. The type is slice of
    // Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_Owner.
    Owner []*Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_Owner
}

func (locatorAllStaleSid *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid) GetEntityData() *types.CommonEntityData {
    locatorAllStaleSid.EntityData.YFilter = locatorAllStaleSid.YFilter
    locatorAllStaleSid.EntityData.YangName = "locator-all-stale-sid"
    locatorAllStaleSid.EntityData.BundleName = "cisco_ios_xr"
    locatorAllStaleSid.EntityData.ParentYangName = "locator-all-stale-sids"
    locatorAllStaleSid.EntityData.SegmentPath = "locator-all-stale-sid" + types.AddKeyToken(locatorAllStaleSid.LocatorName, "locator-name") + types.AddKeyToken(locatorAllStaleSid.SidOpcode, "sid-opcode")
    locatorAllStaleSid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/" + locatorAllStaleSid.EntityData.SegmentPath
    locatorAllStaleSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllStaleSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllStaleSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllStaleSid.EntityData.Children = types.NewOrderedMap()
    locatorAllStaleSid.EntityData.Children.Append("sid-context", types.YChild{"SidContext", &locatorAllStaleSid.SidContext})
    locatorAllStaleSid.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &locatorAllStaleSid.CreateTimestamp})
    locatorAllStaleSid.EntityData.Children.Append("owner", types.YChild{"Owner", nil})
    for i := range locatorAllStaleSid.Owner {
        types.SetYListKey(locatorAllStaleSid.Owner[i], i)
        locatorAllStaleSid.EntityData.Children.Append(types.GetSegmentPath(locatorAllStaleSid.Owner[i]), types.YChild{"Owner", locatorAllStaleSid.Owner[i]})
    }
    locatorAllStaleSid.EntityData.Leafs = types.NewOrderedMap()
    locatorAllStaleSid.EntityData.Leafs.Append("locator-name", types.YLeaf{"LocatorName", locatorAllStaleSid.LocatorName})
    locatorAllStaleSid.EntityData.Leafs.Append("sid-opcode", types.YLeaf{"SidOpcode", locatorAllStaleSid.SidOpcode})
    locatorAllStaleSid.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", locatorAllStaleSid.Sid})
    locatorAllStaleSid.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", locatorAllStaleSid.AllocationType})
    locatorAllStaleSid.EntityData.Leafs.Append("function-type", types.YLeaf{"FunctionType", locatorAllStaleSid.FunctionType})
    locatorAllStaleSid.EntityData.Leafs.Append("state", types.YLeaf{"State", locatorAllStaleSid.State})
    locatorAllStaleSid.EntityData.Leafs.Append("has-forwarding", types.YLeaf{"HasForwarding", locatorAllStaleSid.HasForwarding})
    locatorAllStaleSid.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", locatorAllStaleSid.Locator})

    locatorAllStaleSid.EntityData.YListKeys = []string {"LocatorName", "SidOpcode"}

    return &(locatorAllStaleSid.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext
// SID Context
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application opaque data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ApplicationData interface{}

    // SID Key.
    Key Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key
}

func (sidContext *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext) GetEntityData() *types.CommonEntityData {
    sidContext.EntityData.YFilter = sidContext.YFilter
    sidContext.EntityData.YangName = "sid-context"
    sidContext.EntityData.BundleName = "cisco_ios_xr"
    sidContext.EntityData.ParentYangName = "locator-all-stale-sid"
    sidContext.EntityData.SegmentPath = "sid-context"
    sidContext.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/locator-all-stale-sid/" + sidContext.EntityData.SegmentPath
    sidContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidContext.EntityData.Children = types.NewOrderedMap()
    sidContext.EntityData.Children.Append("key", types.YChild{"Key", &sidContext.Key})
    sidContext.EntityData.Leafs = types.NewOrderedMap()
    sidContext.EntityData.Leafs.Append("application-data", types.YLeaf{"ApplicationData", sidContext.ApplicationData})

    sidContext.EntityData.YListKeys = []string {}

    return &(sidContext.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key
// SID Key
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDContextType. The type is Srv6EndFunction.
    SidContextType interface{}

    // End (PSP) SID context.
    E Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_E

    // End.X (PSP) SID context.
    X Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_X

    // End.DX4 SID context.
    Dx4 Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_Dx4

    // End.DT4 SID context.
    Dt4 Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_Dt4
}

func (key *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "sid-context"
    key.EntityData.SegmentPath = "key"
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/locator-all-stale-sid/sid-context/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("e", types.YChild{"E", &key.E})
    key.EntityData.Children.Append("x", types.YChild{"X", &key.X})
    key.EntityData.Children.Append("dx4", types.YChild{"Dx4", &key.Dx4})
    key.EntityData.Children.Append("dt4", types.YChild{"Dt4", &key.Dt4})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("sid-context-type", types.YLeaf{"SidContextType", key.SidContextType})

    key.EntityData.YListKeys = []string {}

    return &(key.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_E
// End (PSP) SID context
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_E struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Id. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}
}

func (e *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_E) GetEntityData() *types.CommonEntityData {
    e.EntityData.YFilter = e.YFilter
    e.EntityData.YangName = "e"
    e.EntityData.BundleName = "cisco_ios_xr"
    e.EntityData.ParentYangName = "key"
    e.EntityData.SegmentPath = "e"
    e.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/locator-all-stale-sid/sid-context/key/" + e.EntityData.SegmentPath
    e.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    e.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    e.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    e.EntityData.Children = types.NewOrderedMap()
    e.EntityData.Leafs = types.NewOrderedMap()
    e.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", e.TableId})
    e.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", e.OpaqueId})

    e.EntityData.YListKeys = []string {}

    return &(e.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_X
// End.X (PSP) SID context
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_X struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is protected?. The type is bool.
    IsProtected interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}

    // Nexthop interface. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Interface interface{}

    // Nexthop IP address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}
}

func (x *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_X) GetEntityData() *types.CommonEntityData {
    x.EntityData.YFilter = x.YFilter
    x.EntityData.YangName = "x"
    x.EntityData.BundleName = "cisco_ios_xr"
    x.EntityData.ParentYangName = "key"
    x.EntityData.SegmentPath = "x"
    x.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/locator-all-stale-sid/sid-context/key/" + x.EntityData.SegmentPath
    x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    x.EntityData.Children = types.NewOrderedMap()
    x.EntityData.Leafs = types.NewOrderedMap()
    x.EntityData.Leafs.Append("is-protected", types.YLeaf{"IsProtected", x.IsProtected})
    x.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", x.OpaqueId})
    x.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", x.Interface})
    x.EntityData.Leafs.Append("nexthop-address", types.YLeaf{"NexthopAddress", x.NexthopAddress})

    x.EntityData.YListKeys = []string {}

    return &(x.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_Dx4
// End.DX4 SID context
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_Dx4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Next Hop Set ID. The type is interface{} with range: 0..4294967295.
    NextHopSetId interface{}
}

func (dx4 *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_Dx4) GetEntityData() *types.CommonEntityData {
    dx4.EntityData.YFilter = dx4.YFilter
    dx4.EntityData.YangName = "dx4"
    dx4.EntityData.BundleName = "cisco_ios_xr"
    dx4.EntityData.ParentYangName = "key"
    dx4.EntityData.SegmentPath = "dx4"
    dx4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/locator-all-stale-sid/sid-context/key/" + dx4.EntityData.SegmentPath
    dx4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dx4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dx4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dx4.EntityData.Children = types.NewOrderedMap()
    dx4.EntityData.Leafs = types.NewOrderedMap()
    dx4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dx4.TableId})
    dx4.EntityData.Leafs.Append("next-hop-set-id", types.YLeaf{"NextHopSetId", dx4.NextHopSetId})

    dx4.EntityData.YListKeys = []string {}

    return &(dx4.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_Dt4
// End.DT4 SID context
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_Dt4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}
}

func (dt4 *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_SidContext_Key_Dt4) GetEntityData() *types.CommonEntityData {
    dt4.EntityData.YFilter = dt4.YFilter
    dt4.EntityData.YangName = "dt4"
    dt4.EntityData.BundleName = "cisco_ios_xr"
    dt4.EntityData.ParentYangName = "key"
    dt4.EntityData.SegmentPath = "dt4"
    dt4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/locator-all-stale-sid/sid-context/key/" + dt4.EntityData.SegmentPath
    dt4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dt4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dt4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dt4.EntityData.Children = types.NewOrderedMap()
    dt4.EntityData.Leafs = types.NewOrderedMap()
    dt4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dt4.TableId})

    dt4.EntityData.YListKeys = []string {}

    return &(dt4.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_CreateTimestamp
// Creation timestamp
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "locator-all-stale-sid"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/locator-all-stale-sid/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_Owner
// Owner
type Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_Owner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Owner. The type is string.
    Owner interface{}
}

func (owner *Srv6_Active_LocatorAllStaleSids_LocatorAllStaleSid_Owner) GetEntityData() *types.CommonEntityData {
    owner.EntityData.YFilter = owner.YFilter
    owner.EntityData.YangName = "owner"
    owner.EntityData.BundleName = "cisco_ios_xr"
    owner.EntityData.ParentYangName = "locator-all-stale-sid"
    owner.EntityData.SegmentPath = "owner" + types.AddNoKeyToken(owner)
    owner.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-stale-sids/locator-all-stale-sid/" + owner.EntityData.SegmentPath
    owner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    owner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    owner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    owner.EntityData.Children = types.NewOrderedMap()
    owner.EntityData.Leafs = types.NewOrderedMap()
    owner.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", owner.Owner})

    owner.EntityData.YListKeys = []string {}

    return &(owner.EntityData)
}

// Srv6_Active_Manager
// SID Manager information
type Srv6_Active_Manager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID Mgr parameters.
    SidMgrParams Srv6_Active_Manager_SidMgrParams

    // SID Mgr summary info.
    SidMgrSummary Srv6_Active_Manager_SidMgrSummary

    // Platform Capabilities.
    PlatformCapabilities Srv6_Active_Manager_PlatformCapabilities
}

func (manager *Srv6_Active_Manager) GetEntityData() *types.CommonEntityData {
    manager.EntityData.YFilter = manager.YFilter
    manager.EntityData.YangName = "manager"
    manager.EntityData.BundleName = "cisco_ios_xr"
    manager.EntityData.ParentYangName = "active"
    manager.EntityData.SegmentPath = "manager"
    manager.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/" + manager.EntityData.SegmentPath
    manager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manager.EntityData.Children = types.NewOrderedMap()
    manager.EntityData.Children.Append("sid-mgr-params", types.YChild{"SidMgrParams", &manager.SidMgrParams})
    manager.EntityData.Children.Append("sid-mgr-summary", types.YChild{"SidMgrSummary", &manager.SidMgrSummary})
    manager.EntityData.Children.Append("platform-capabilities", types.YChild{"PlatformCapabilities", &manager.PlatformCapabilities})
    manager.EntityData.Leafs = types.NewOrderedMap()

    manager.EntityData.YListKeys = []string {}

    return &(manager.EntityData)
}

// Srv6_Active_Manager_SidMgrParams
// SID Mgr parameters
type Srv6_Active_Manager_SidMgrParams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SRv6 enabled?. The type is bool.
    Srv6Enabled interface{}

    // Configured Encap Source address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ConfiguredEncapSourceAddress interface{}

    // Default Encap Source address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DefaultEncapSourceAddress interface{}

    // Is TTL propagate enabled?. The type is bool.
    EncapTtlPropagate interface{}

    // Is SID Holdtime configured?. The type is bool.
    IsSidHoldtimeConfigured interface{}

    // Configured SID Holdtime in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    SidHoldtimeMinsConfigured interface{}

    // Encap Hop-limit info.
    EncapHopLimit Srv6_Active_Manager_SidMgrParams_EncapHopLimit
}

func (sidMgrParams *Srv6_Active_Manager_SidMgrParams) GetEntityData() *types.CommonEntityData {
    sidMgrParams.EntityData.YFilter = sidMgrParams.YFilter
    sidMgrParams.EntityData.YangName = "sid-mgr-params"
    sidMgrParams.EntityData.BundleName = "cisco_ios_xr"
    sidMgrParams.EntityData.ParentYangName = "manager"
    sidMgrParams.EntityData.SegmentPath = "sid-mgr-params"
    sidMgrParams.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/" + sidMgrParams.EntityData.SegmentPath
    sidMgrParams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidMgrParams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidMgrParams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidMgrParams.EntityData.Children = types.NewOrderedMap()
    sidMgrParams.EntityData.Children.Append("encap-hop-limit", types.YChild{"EncapHopLimit", &sidMgrParams.EncapHopLimit})
    sidMgrParams.EntityData.Leafs = types.NewOrderedMap()
    sidMgrParams.EntityData.Leafs.Append("srv6-enabled", types.YLeaf{"Srv6Enabled", sidMgrParams.Srv6Enabled})
    sidMgrParams.EntityData.Leafs.Append("configured-encap-source-address", types.YLeaf{"ConfiguredEncapSourceAddress", sidMgrParams.ConfiguredEncapSourceAddress})
    sidMgrParams.EntityData.Leafs.Append("default-encap-source-address", types.YLeaf{"DefaultEncapSourceAddress", sidMgrParams.DefaultEncapSourceAddress})
    sidMgrParams.EntityData.Leafs.Append("encap-ttl-propagate", types.YLeaf{"EncapTtlPropagate", sidMgrParams.EncapTtlPropagate})
    sidMgrParams.EntityData.Leafs.Append("is-sid-holdtime-configured", types.YLeaf{"IsSidHoldtimeConfigured", sidMgrParams.IsSidHoldtimeConfigured})
    sidMgrParams.EntityData.Leafs.Append("sid-holdtime-mins-configured", types.YLeaf{"SidHoldtimeMinsConfigured", sidMgrParams.SidHoldtimeMinsConfigured})

    sidMgrParams.EntityData.YListKeys = []string {}

    return &(sidMgrParams.EntityData)
}

// Srv6_Active_Manager_SidMgrParams_EncapHopLimit
// Encap Hop-limit info
type Srv6_Active_Manager_SidMgrParams_EncapHopLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use default IPv6 hop-limit value. The type is bool.
    UseDefault interface{}

    // Propagate IP TTL to Encap IPv6 hop-limit. The type is bool.
    DoPropagate interface{}

    // Specific value set for hop-limit count. The type is interface{} with range:
    // 0..255.
    Value interface{}
}

func (encapHopLimit *Srv6_Active_Manager_SidMgrParams_EncapHopLimit) GetEntityData() *types.CommonEntityData {
    encapHopLimit.EntityData.YFilter = encapHopLimit.YFilter
    encapHopLimit.EntityData.YangName = "encap-hop-limit"
    encapHopLimit.EntityData.BundleName = "cisco_ios_xr"
    encapHopLimit.EntityData.ParentYangName = "sid-mgr-params"
    encapHopLimit.EntityData.SegmentPath = "encap-hop-limit"
    encapHopLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/sid-mgr-params/" + encapHopLimit.EntityData.SegmentPath
    encapHopLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encapHopLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encapHopLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encapHopLimit.EntityData.Children = types.NewOrderedMap()
    encapHopLimit.EntityData.Leafs = types.NewOrderedMap()
    encapHopLimit.EntityData.Leafs.Append("use-default", types.YLeaf{"UseDefault", encapHopLimit.UseDefault})
    encapHopLimit.EntityData.Leafs.Append("do-propagate", types.YLeaf{"DoPropagate", encapHopLimit.DoPropagate})
    encapHopLimit.EntityData.Leafs.Append("value", types.YLeaf{"Value", encapHopLimit.Value})

    encapHopLimit.EntityData.YListKeys = []string {}

    return &(encapHopLimit.EntityData)
}

// Srv6_Active_Manager_SidMgrSummary
// SID Mgr summary info
type Srv6_Active_Manager_SidMgrSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of locators. The type is interface{} with range: 0..65535.
    LocatorsCount interface{}

    // Number of operational locators. The type is interface{} with range:
    // 0..65535.
    OperLocatorsCount interface{}

    // Number of SIDs. The type is interface{} with range: 0..4294967295.
    SidsCount interface{}

    // Number of Stale SIDs. The type is interface{} with range: 0..4294967295.
    StaleSidsCount interface{}

    // Global Maximum number of SIDs. The type is interface{} with range:
    // 0..4294967295.
    MaximumSidsCount interface{}

    // SIDs Global Out of Resource info.
    SidsOutOfResourceSummary Srv6_Active_Manager_SidMgrSummary_SidsOutOfResourceSummary
}

func (sidMgrSummary *Srv6_Active_Manager_SidMgrSummary) GetEntityData() *types.CommonEntityData {
    sidMgrSummary.EntityData.YFilter = sidMgrSummary.YFilter
    sidMgrSummary.EntityData.YangName = "sid-mgr-summary"
    sidMgrSummary.EntityData.BundleName = "cisco_ios_xr"
    sidMgrSummary.EntityData.ParentYangName = "manager"
    sidMgrSummary.EntityData.SegmentPath = "sid-mgr-summary"
    sidMgrSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/" + sidMgrSummary.EntityData.SegmentPath
    sidMgrSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidMgrSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidMgrSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidMgrSummary.EntityData.Children = types.NewOrderedMap()
    sidMgrSummary.EntityData.Children.Append("sids-out-of-resource-summary", types.YChild{"SidsOutOfResourceSummary", &sidMgrSummary.SidsOutOfResourceSummary})
    sidMgrSummary.EntityData.Leafs = types.NewOrderedMap()
    sidMgrSummary.EntityData.Leafs.Append("locators-count", types.YLeaf{"LocatorsCount", sidMgrSummary.LocatorsCount})
    sidMgrSummary.EntityData.Leafs.Append("oper-locators-count", types.YLeaf{"OperLocatorsCount", sidMgrSummary.OperLocatorsCount})
    sidMgrSummary.EntityData.Leafs.Append("sids-count", types.YLeaf{"SidsCount", sidMgrSummary.SidsCount})
    sidMgrSummary.EntityData.Leafs.Append("stale-sids-count", types.YLeaf{"StaleSidsCount", sidMgrSummary.StaleSidsCount})
    sidMgrSummary.EntityData.Leafs.Append("maximum-sids-count", types.YLeaf{"MaximumSidsCount", sidMgrSummary.MaximumSidsCount})

    sidMgrSummary.EntityData.YListKeys = []string {}

    return &(sidMgrSummary.EntityData)
}

// Srv6_Active_Manager_SidMgrSummary_SidsOutOfResourceSummary
// SIDs Global Out of Resource info
type Srv6_Active_Manager_SidMgrSummary_SidsOutOfResourceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global Resources State for SIDs. The type is Srv6OutOfResourceState.
    OutOfResourcesState interface{}

    // Threshold for Number of Free SID below which OOR Yellow State is reached.
    // The type is interface{} with range: 0..4294967295.
    OorYellowFreeSidThreshold interface{}

    // Threshold for Number of Free SID above which OOR Green State is restored.
    // The type is interface{} with range: 0..4294967295.
    OorGreenFreeSidThreshold interface{}

    // Number of times Resources Warning or Out of Resources state has been
    // cleared. The type is interface{} with range: 0..4294967295.
    OorGreenCount interface{}

    // Number of times system went into Resources Warning state. The type is
    // interface{} with range: 0..4294967295.
    OorYellowCount interface{}

    // Number of times system went into Out of Resources state. The type is
    // interface{} with range: 0..4294967295.
    OorRedCount interface{}
}

func (sidsOutOfResourceSummary *Srv6_Active_Manager_SidMgrSummary_SidsOutOfResourceSummary) GetEntityData() *types.CommonEntityData {
    sidsOutOfResourceSummary.EntityData.YFilter = sidsOutOfResourceSummary.YFilter
    sidsOutOfResourceSummary.EntityData.YangName = "sids-out-of-resource-summary"
    sidsOutOfResourceSummary.EntityData.BundleName = "cisco_ios_xr"
    sidsOutOfResourceSummary.EntityData.ParentYangName = "sid-mgr-summary"
    sidsOutOfResourceSummary.EntityData.SegmentPath = "sids-out-of-resource-summary"
    sidsOutOfResourceSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/sid-mgr-summary/" + sidsOutOfResourceSummary.EntityData.SegmentPath
    sidsOutOfResourceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidsOutOfResourceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidsOutOfResourceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidsOutOfResourceSummary.EntityData.Children = types.NewOrderedMap()
    sidsOutOfResourceSummary.EntityData.Leafs = types.NewOrderedMap()
    sidsOutOfResourceSummary.EntityData.Leafs.Append("out-of-resources-state", types.YLeaf{"OutOfResourcesState", sidsOutOfResourceSummary.OutOfResourcesState})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-yellow-free-sid-threshold", types.YLeaf{"OorYellowFreeSidThreshold", sidsOutOfResourceSummary.OorYellowFreeSidThreshold})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-green-free-sid-threshold", types.YLeaf{"OorGreenFreeSidThreshold", sidsOutOfResourceSummary.OorGreenFreeSidThreshold})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-green-count", types.YLeaf{"OorGreenCount", sidsOutOfResourceSummary.OorGreenCount})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-yellow-count", types.YLeaf{"OorYellowCount", sidsOutOfResourceSummary.OorYellowCount})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-red-count", types.YLeaf{"OorRedCount", sidsOutOfResourceSummary.OorRedCount})

    sidsOutOfResourceSummary.EntityData.YListKeys = []string {}

    return &(sidsOutOfResourceSummary.EntityData)
}

// Srv6_Active_Manager_PlatformCapabilities
// Platform Capabilities
type Srv6_Active_Manager_PlatformCapabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum Sids. The type is interface{} with range: 0..4294967295.
    MaxSid interface{}

    // Freed SID holdtime in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    SidHoldtimeMins interface{}

    // Feature support.
    Support Srv6_Active_Manager_PlatformCapabilities_Support
}

func (platformCapabilities *Srv6_Active_Manager_PlatformCapabilities) GetEntityData() *types.CommonEntityData {
    platformCapabilities.EntityData.YFilter = platformCapabilities.YFilter
    platformCapabilities.EntityData.YangName = "platform-capabilities"
    platformCapabilities.EntityData.BundleName = "cisco_ios_xr"
    platformCapabilities.EntityData.ParentYangName = "manager"
    platformCapabilities.EntityData.SegmentPath = "platform-capabilities"
    platformCapabilities.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/" + platformCapabilities.EntityData.SegmentPath
    platformCapabilities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformCapabilities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformCapabilities.EntityData.Children = types.NewOrderedMap()
    platformCapabilities.EntityData.Children.Append("support", types.YChild{"Support", &platformCapabilities.Support})
    platformCapabilities.EntityData.Leafs = types.NewOrderedMap()
    platformCapabilities.EntityData.Leafs.Append("max-sid", types.YLeaf{"MaxSid", platformCapabilities.MaxSid})
    platformCapabilities.EntityData.Leafs.Append("sid-holdtime-mins", types.YLeaf{"SidHoldtimeMins", platformCapabilities.SidHoldtimeMins})

    platformCapabilities.EntityData.YListKeys = []string {}

    return &(platformCapabilities.EntityData)
}

// Srv6_Active_Manager_PlatformCapabilities_Support
// Feature support
type Srv6_Active_Manager_PlatformCapabilities_Support struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRv6 support. The type is bool.
    Srv6 interface{}

    // TI LFA support. The type is bool.
    Tilfa interface{}

    // Microloop-avoidance support. The type is bool.
    MicroloopAvoidance interface{}

    // Signaled Parameters.
    SignaledParameters Srv6_Active_Manager_PlatformCapabilities_Support_SignaledParameters

    // Supported end functions. The type is slice of
    // Srv6_Active_Manager_PlatformCapabilities_Support_EndFunc.
    EndFunc []*Srv6_Active_Manager_PlatformCapabilities_Support_EndFunc

    // Supported Transit functions. The type is slice of
    // Srv6_Active_Manager_PlatformCapabilities_Support_TransitFunc.
    TransitFunc []*Srv6_Active_Manager_PlatformCapabilities_Support_TransitFunc

    // Security rules. The type is slice of
    // Srv6_Active_Manager_PlatformCapabilities_Support_SecurityRule.
    SecurityRule []*Srv6_Active_Manager_PlatformCapabilities_Support_SecurityRule

    // Counters. The type is slice of
    // Srv6_Active_Manager_PlatformCapabilities_Support_Counter.
    Counter []*Srv6_Active_Manager_PlatformCapabilities_Support_Counter
}

func (support *Srv6_Active_Manager_PlatformCapabilities_Support) GetEntityData() *types.CommonEntityData {
    support.EntityData.YFilter = support.YFilter
    support.EntityData.YangName = "support"
    support.EntityData.BundleName = "cisco_ios_xr"
    support.EntityData.ParentYangName = "platform-capabilities"
    support.EntityData.SegmentPath = "support"
    support.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/platform-capabilities/" + support.EntityData.SegmentPath
    support.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    support.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    support.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    support.EntityData.Children = types.NewOrderedMap()
    support.EntityData.Children.Append("signaled-parameters", types.YChild{"SignaledParameters", &support.SignaledParameters})
    support.EntityData.Children.Append("end-func", types.YChild{"EndFunc", nil})
    for i := range support.EndFunc {
        types.SetYListKey(support.EndFunc[i], i)
        support.EntityData.Children.Append(types.GetSegmentPath(support.EndFunc[i]), types.YChild{"EndFunc", support.EndFunc[i]})
    }
    support.EntityData.Children.Append("transit-func", types.YChild{"TransitFunc", nil})
    for i := range support.TransitFunc {
        types.SetYListKey(support.TransitFunc[i], i)
        support.EntityData.Children.Append(types.GetSegmentPath(support.TransitFunc[i]), types.YChild{"TransitFunc", support.TransitFunc[i]})
    }
    support.EntityData.Children.Append("security-rule", types.YChild{"SecurityRule", nil})
    for i := range support.SecurityRule {
        types.SetYListKey(support.SecurityRule[i], i)
        support.EntityData.Children.Append(types.GetSegmentPath(support.SecurityRule[i]), types.YChild{"SecurityRule", support.SecurityRule[i]})
    }
    support.EntityData.Children.Append("counter", types.YChild{"Counter", nil})
    for i := range support.Counter {
        types.SetYListKey(support.Counter[i], i)
        support.EntityData.Children.Append(types.GetSegmentPath(support.Counter[i]), types.YChild{"Counter", support.Counter[i]})
    }
    support.EntityData.Leafs = types.NewOrderedMap()
    support.EntityData.Leafs.Append("srv6", types.YLeaf{"Srv6", support.Srv6})
    support.EntityData.Leafs.Append("tilfa", types.YLeaf{"Tilfa", support.Tilfa})
    support.EntityData.Leafs.Append("microloop-avoidance", types.YLeaf{"MicroloopAvoidance", support.MicroloopAvoidance})

    support.EntityData.YListKeys = []string {}

    return &(support.EntityData)
}

// Srv6_Active_Manager_PlatformCapabilities_Support_SignaledParameters
// Signaled Parameters
type Srv6_Active_Manager_PlatformCapabilities_Support_SignaledParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max value of SegmentLeft field in received SRH. The type is interface{}
    // with range: 0..255.
    MaxSl interface{}

    // Max num of SIDs in rcvd SRH for pop. The type is interface{} with range:
    // 0..255.
    MaxEndPopSrh interface{}

    // Max num of SIDs for T.Insert op. The type is interface{} with range:
    // 0..255.
    MaxTInsert interface{}

    // Max num of SIDs for T.Encaps op. The type is interface{} with range:
    // 0..255.
    MaxTEncap interface{}

    // Max num of SIDs in rcvd SRH for decap. The type is interface{} with range:
    // 0..255.
    MaxEndD interface{}
}

func (signaledParameters *Srv6_Active_Manager_PlatformCapabilities_Support_SignaledParameters) GetEntityData() *types.CommonEntityData {
    signaledParameters.EntityData.YFilter = signaledParameters.YFilter
    signaledParameters.EntityData.YangName = "signaled-parameters"
    signaledParameters.EntityData.BundleName = "cisco_ios_xr"
    signaledParameters.EntityData.ParentYangName = "support"
    signaledParameters.EntityData.SegmentPath = "signaled-parameters"
    signaledParameters.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/platform-capabilities/support/" + signaledParameters.EntityData.SegmentPath
    signaledParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signaledParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signaledParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signaledParameters.EntityData.Children = types.NewOrderedMap()
    signaledParameters.EntityData.Leafs = types.NewOrderedMap()
    signaledParameters.EntityData.Leafs.Append("max-sl", types.YLeaf{"MaxSl", signaledParameters.MaxSl})
    signaledParameters.EntityData.Leafs.Append("max-end-pop-srh", types.YLeaf{"MaxEndPopSrh", signaledParameters.MaxEndPopSrh})
    signaledParameters.EntityData.Leafs.Append("max-t-insert", types.YLeaf{"MaxTInsert", signaledParameters.MaxTInsert})
    signaledParameters.EntityData.Leafs.Append("max-t-encap", types.YLeaf{"MaxTEncap", signaledParameters.MaxTEncap})
    signaledParameters.EntityData.Leafs.Append("max-end-d", types.YLeaf{"MaxEndD", signaledParameters.MaxEndD})

    signaledParameters.EntityData.YListKeys = []string {}

    return &(signaledParameters.EntityData)
}

// Srv6_Active_Manager_PlatformCapabilities_Support_EndFunc
// Supported end functions
type Srv6_Active_Manager_PlatformCapabilities_Support_EndFunc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // String. The type is string.
    String interface{}
}

func (endFunc *Srv6_Active_Manager_PlatformCapabilities_Support_EndFunc) GetEntityData() *types.CommonEntityData {
    endFunc.EntityData.YFilter = endFunc.YFilter
    endFunc.EntityData.YangName = "end-func"
    endFunc.EntityData.BundleName = "cisco_ios_xr"
    endFunc.EntityData.ParentYangName = "support"
    endFunc.EntityData.SegmentPath = "end-func" + types.AddNoKeyToken(endFunc)
    endFunc.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/platform-capabilities/support/" + endFunc.EntityData.SegmentPath
    endFunc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    endFunc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    endFunc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    endFunc.EntityData.Children = types.NewOrderedMap()
    endFunc.EntityData.Leafs = types.NewOrderedMap()
    endFunc.EntityData.Leafs.Append("string", types.YLeaf{"String", endFunc.String})

    endFunc.EntityData.YListKeys = []string {}

    return &(endFunc.EntityData)
}

// Srv6_Active_Manager_PlatformCapabilities_Support_TransitFunc
// Supported Transit functions
type Srv6_Active_Manager_PlatformCapabilities_Support_TransitFunc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // String. The type is string.
    String interface{}
}

func (transitFunc *Srv6_Active_Manager_PlatformCapabilities_Support_TransitFunc) GetEntityData() *types.CommonEntityData {
    transitFunc.EntityData.YFilter = transitFunc.YFilter
    transitFunc.EntityData.YangName = "transit-func"
    transitFunc.EntityData.BundleName = "cisco_ios_xr"
    transitFunc.EntityData.ParentYangName = "support"
    transitFunc.EntityData.SegmentPath = "transit-func" + types.AddNoKeyToken(transitFunc)
    transitFunc.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/platform-capabilities/support/" + transitFunc.EntityData.SegmentPath
    transitFunc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transitFunc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transitFunc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transitFunc.EntityData.Children = types.NewOrderedMap()
    transitFunc.EntityData.Leafs = types.NewOrderedMap()
    transitFunc.EntityData.Leafs.Append("string", types.YLeaf{"String", transitFunc.String})

    transitFunc.EntityData.YListKeys = []string {}

    return &(transitFunc.EntityData)
}

// Srv6_Active_Manager_PlatformCapabilities_Support_SecurityRule
// Security rules
type Srv6_Active_Manager_PlatformCapabilities_Support_SecurityRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // String. The type is string.
    String interface{}
}

func (securityRule *Srv6_Active_Manager_PlatformCapabilities_Support_SecurityRule) GetEntityData() *types.CommonEntityData {
    securityRule.EntityData.YFilter = securityRule.YFilter
    securityRule.EntityData.YangName = "security-rule"
    securityRule.EntityData.BundleName = "cisco_ios_xr"
    securityRule.EntityData.ParentYangName = "support"
    securityRule.EntityData.SegmentPath = "security-rule" + types.AddNoKeyToken(securityRule)
    securityRule.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/platform-capabilities/support/" + securityRule.EntityData.SegmentPath
    securityRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    securityRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    securityRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    securityRule.EntityData.Children = types.NewOrderedMap()
    securityRule.EntityData.Leafs = types.NewOrderedMap()
    securityRule.EntityData.Leafs.Append("string", types.YLeaf{"String", securityRule.String})

    securityRule.EntityData.YListKeys = []string {}

    return &(securityRule.EntityData)
}

// Srv6_Active_Manager_PlatformCapabilities_Support_Counter
// Counters
type Srv6_Active_Manager_PlatformCapabilities_Support_Counter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // String. The type is string.
    String interface{}
}

func (counter *Srv6_Active_Manager_PlatformCapabilities_Support_Counter) GetEntityData() *types.CommonEntityData {
    counter.EntityData.YFilter = counter.YFilter
    counter.EntityData.YangName = "counter"
    counter.EntityData.BundleName = "cisco_ios_xr"
    counter.EntityData.ParentYangName = "support"
    counter.EntityData.SegmentPath = "counter" + types.AddNoKeyToken(counter)
    counter.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/manager/platform-capabilities/support/" + counter.EntityData.SegmentPath
    counter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counter.EntityData.Children = types.NewOrderedMap()
    counter.EntityData.Leafs = types.NewOrderedMap()
    counter.EntityData.Leafs.Append("string", types.YLeaf{"String", counter.String})

    counter.EntityData.YListKeys = []string {}

    return &(counter.EntityData)
}

// Srv6_Active_Locators
// SRv6 locators related information
type Srv6_Active_Locators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given SRv6 locator. The type is slice of
    // Srv6_Active_Locators_Locator.
    Locator []*Srv6_Active_Locators_Locator
}

func (locators *Srv6_Active_Locators) GetEntityData() *types.CommonEntityData {
    locators.EntityData.YFilter = locators.YFilter
    locators.EntityData.YangName = "locators"
    locators.EntityData.BundleName = "cisco_ios_xr"
    locators.EntityData.ParentYangName = "active"
    locators.EntityData.SegmentPath = "locators"
    locators.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/" + locators.EntityData.SegmentPath
    locators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locators.EntityData.Children = types.NewOrderedMap()
    locators.EntityData.Children.Append("locator", types.YChild{"Locator", nil})
    for i := range locators.Locator {
        locators.EntityData.Children.Append(types.GetSegmentPath(locators.Locator[i]), types.YChild{"Locator", locators.Locator[i]})
    }
    locators.EntityData.Leafs = types.NewOrderedMap()

    locators.EntityData.YListKeys = []string {}

    return &(locators.EntityData)
}

// Srv6_Active_Locators_Locator
// Operational data for given SRv6 locator
type Srv6_Active_Locators_Locator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Locator name. The type is string with length:
    // 1..58.
    Name interface{}

    // Operational data for given SRv6 locator.
    Info Srv6_Active_Locators_Locator_Info

    // SRv6 locator SID table.
    Sids Srv6_Active_Locators_Locator_Sids
}

func (locator *Srv6_Active_Locators_Locator) GetEntityData() *types.CommonEntityData {
    locator.EntityData.YFilter = locator.YFilter
    locator.EntityData.YangName = "locator"
    locator.EntityData.BundleName = "cisco_ios_xr"
    locator.EntityData.ParentYangName = "locators"
    locator.EntityData.SegmentPath = "locator" + types.AddKeyToken(locator.Name, "name")
    locator.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/" + locator.EntityData.SegmentPath
    locator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locator.EntityData.Children = types.NewOrderedMap()
    locator.EntityData.Children.Append("info", types.YChild{"Info", &locator.Info})
    locator.EntityData.Children.Append("sids", types.YChild{"Sids", &locator.Sids})
    locator.EntityData.Leafs = types.NewOrderedMap()
    locator.EntityData.Leafs.Append("name", types.YLeaf{"Name", locator.Name})

    locator.EntityData.YListKeys = []string {"Name"}

    return &(locator.EntityData)
}

// Srv6_Active_Locators_Locator_Info
// Operational data for given SRv6 locator
type Srv6_Active_Locators_Locator_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Locator Name. The type is string.
    Name interface{}

    // Locator ID. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // Locator Prefix. The type is string.
    Prefix interface{}

    // Locator status is Up or Down. The type is bool.
    IsOperational interface{}

    // Locator is the default locator. The type is bool.
    IsDefault interface{}

    // Locator Resources State for SIDs. The type is Srv6OutOfResourceState.
    OutOfResourcesState interface{}

    // Locator IM intf info.
    Interface Srv6_Active_Locators_Locator_Info_Interface

    // Creation timestamp.
    CreateTimestamp Srv6_Active_Locators_Locator_Info_CreateTimestamp
}

func (info *Srv6_Active_Locators_Locator_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "locator"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/" + info.EntityData.SegmentPath
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("interface", types.YChild{"Interface", &info.Interface})
    info.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &info.CreateTimestamp})
    info.EntityData.Leafs = types.NewOrderedMap()
    info.EntityData.Leafs.Append("name", types.YLeaf{"Name", info.Name})
    info.EntityData.Leafs.Append("id", types.YLeaf{"Id", info.Id})
    info.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", info.Prefix})
    info.EntityData.Leafs.Append("is-operational", types.YLeaf{"IsOperational", info.IsOperational})
    info.EntityData.Leafs.Append("is-default", types.YLeaf{"IsDefault", info.IsDefault})
    info.EntityData.Leafs.Append("out-of-resources-state", types.YLeaf{"OutOfResourcesState", info.OutOfResourcesState})

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Srv6_Active_Locators_Locator_Info_Interface
// Locator IM intf info
type Srv6_Active_Locators_Locator_Info_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    Name interface{}

    // Interface handle. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    IfHandle interface{}

    // Interface prefix/addr programmed. The type is string.
    ProgrammedPrefix interface{}
}

func (self *Srv6_Active_Locators_Locator_Info_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "info"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/info/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})
    self.EntityData.Leafs.Append("if-handle", types.YLeaf{"IfHandle", self.IfHandle})
    self.EntityData.Leafs.Append("programmed-prefix", types.YLeaf{"ProgrammedPrefix", self.ProgrammedPrefix})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Srv6_Active_Locators_Locator_Info_CreateTimestamp
// Creation timestamp
type Srv6_Active_Locators_Locator_Info_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Active_Locators_Locator_Info_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "info"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/info/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Active_Locators_Locator_Sids
// SRv6 locator SID table
type Srv6_Active_Locators_Locator_Sids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given SRv6 SID. The type is slice of
    // Srv6_Active_Locators_Locator_Sids_Sid.
    Sid []*Srv6_Active_Locators_Locator_Sids_Sid
}

func (sids *Srv6_Active_Locators_Locator_Sids) GetEntityData() *types.CommonEntityData {
    sids.EntityData.YFilter = sids.YFilter
    sids.EntityData.YangName = "sids"
    sids.EntityData.BundleName = "cisco_ios_xr"
    sids.EntityData.ParentYangName = "locator"
    sids.EntityData.SegmentPath = "sids"
    sids.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/" + sids.EntityData.SegmentPath
    sids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sids.EntityData.Children = types.NewOrderedMap()
    sids.EntityData.Children.Append("sid", types.YChild{"Sid", nil})
    for i := range sids.Sid {
        sids.EntityData.Children.Append(types.GetSegmentPath(sids.Sid[i]), types.YChild{"Sid", sids.Sid[i]})
    }
    sids.EntityData.Leafs = types.NewOrderedMap()

    sids.EntityData.YListKeys = []string {}

    return &(sids.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid
// Operational data for given SRv6 SID
type Srv6_Active_Locators_Locator_Sids_Sid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // SID. The type is string.
    Sid interface{}

    // Allocation Type. The type is SidAllocation.
    AllocationType interface{}

    // Function Type. The type is Srv6EndFunction.
    FunctionType interface{}

    // State. The type is SidState.
    State interface{}

    // Rewrite done or not. The type is bool.
    HasForwarding interface{}

    // Associated locator. The type is string.
    Locator interface{}

    // SID Context.
    SidContext Srv6_Active_Locators_Locator_Sids_Sid_SidContext

    // Creation timestamp.
    CreateTimestamp Srv6_Active_Locators_Locator_Sids_Sid_CreateTimestamp

    // Owner. The type is slice of Srv6_Active_Locators_Locator_Sids_Sid_Owner.
    Owner []*Srv6_Active_Locators_Locator_Sids_Sid_Owner
}

func (sid *Srv6_Active_Locators_Locator_Sids_Sid) GetEntityData() *types.CommonEntityData {
    sid.EntityData.YFilter = sid.YFilter
    sid.EntityData.YangName = "sid"
    sid.EntityData.BundleName = "cisco_ios_xr"
    sid.EntityData.ParentYangName = "sids"
    sid.EntityData.SegmentPath = "sid" + types.AddKeyToken(sid.Address, "address")
    sid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/" + sid.EntityData.SegmentPath
    sid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sid.EntityData.Children = types.NewOrderedMap()
    sid.EntityData.Children.Append("sid-context", types.YChild{"SidContext", &sid.SidContext})
    sid.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &sid.CreateTimestamp})
    sid.EntityData.Children.Append("owner", types.YChild{"Owner", nil})
    for i := range sid.Owner {
        types.SetYListKey(sid.Owner[i], i)
        sid.EntityData.Children.Append(types.GetSegmentPath(sid.Owner[i]), types.YChild{"Owner", sid.Owner[i]})
    }
    sid.EntityData.Leafs = types.NewOrderedMap()
    sid.EntityData.Leafs.Append("address", types.YLeaf{"Address", sid.Address})
    sid.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", sid.Sid})
    sid.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", sid.AllocationType})
    sid.EntityData.Leafs.Append("function-type", types.YLeaf{"FunctionType", sid.FunctionType})
    sid.EntityData.Leafs.Append("state", types.YLeaf{"State", sid.State})
    sid.EntityData.Leafs.Append("has-forwarding", types.YLeaf{"HasForwarding", sid.HasForwarding})
    sid.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", sid.Locator})

    sid.EntityData.YListKeys = []string {"Address"}

    return &(sid.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid_SidContext
// SID Context
type Srv6_Active_Locators_Locator_Sids_Sid_SidContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application opaque data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ApplicationData interface{}

    // SID Key.
    Key Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key
}

func (sidContext *Srv6_Active_Locators_Locator_Sids_Sid_SidContext) GetEntityData() *types.CommonEntityData {
    sidContext.EntityData.YFilter = sidContext.YFilter
    sidContext.EntityData.YangName = "sid-context"
    sidContext.EntityData.BundleName = "cisco_ios_xr"
    sidContext.EntityData.ParentYangName = "sid"
    sidContext.EntityData.SegmentPath = "sid-context"
    sidContext.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/sid/" + sidContext.EntityData.SegmentPath
    sidContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidContext.EntityData.Children = types.NewOrderedMap()
    sidContext.EntityData.Children.Append("key", types.YChild{"Key", &sidContext.Key})
    sidContext.EntityData.Leafs = types.NewOrderedMap()
    sidContext.EntityData.Leafs.Append("application-data", types.YLeaf{"ApplicationData", sidContext.ApplicationData})

    sidContext.EntityData.YListKeys = []string {}

    return &(sidContext.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key
// SID Key
type Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDContextType. The type is Srv6EndFunction.
    SidContextType interface{}

    // End (PSP) SID context.
    E Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_E

    // End.X (PSP) SID context.
    X Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_X

    // End.DX4 SID context.
    Dx4 Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_Dx4

    // End.DT4 SID context.
    Dt4 Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_Dt4
}

func (key *Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "sid-context"
    key.EntityData.SegmentPath = "key"
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/sid/sid-context/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("e", types.YChild{"E", &key.E})
    key.EntityData.Children.Append("x", types.YChild{"X", &key.X})
    key.EntityData.Children.Append("dx4", types.YChild{"Dx4", &key.Dx4})
    key.EntityData.Children.Append("dt4", types.YChild{"Dt4", &key.Dt4})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("sid-context-type", types.YLeaf{"SidContextType", key.SidContextType})

    key.EntityData.YListKeys = []string {}

    return &(key.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_E
// End (PSP) SID context
type Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_E struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Id. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}
}

func (e *Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_E) GetEntityData() *types.CommonEntityData {
    e.EntityData.YFilter = e.YFilter
    e.EntityData.YangName = "e"
    e.EntityData.BundleName = "cisco_ios_xr"
    e.EntityData.ParentYangName = "key"
    e.EntityData.SegmentPath = "e"
    e.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/sid/sid-context/key/" + e.EntityData.SegmentPath
    e.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    e.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    e.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    e.EntityData.Children = types.NewOrderedMap()
    e.EntityData.Leafs = types.NewOrderedMap()
    e.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", e.TableId})
    e.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", e.OpaqueId})

    e.EntityData.YListKeys = []string {}

    return &(e.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_X
// End.X (PSP) SID context
type Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_X struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is protected?. The type is bool.
    IsProtected interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}

    // Nexthop interface. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Interface interface{}

    // Nexthop IP address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}
}

func (x *Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_X) GetEntityData() *types.CommonEntityData {
    x.EntityData.YFilter = x.YFilter
    x.EntityData.YangName = "x"
    x.EntityData.BundleName = "cisco_ios_xr"
    x.EntityData.ParentYangName = "key"
    x.EntityData.SegmentPath = "x"
    x.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/sid/sid-context/key/" + x.EntityData.SegmentPath
    x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    x.EntityData.Children = types.NewOrderedMap()
    x.EntityData.Leafs = types.NewOrderedMap()
    x.EntityData.Leafs.Append("is-protected", types.YLeaf{"IsProtected", x.IsProtected})
    x.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", x.OpaqueId})
    x.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", x.Interface})
    x.EntityData.Leafs.Append("nexthop-address", types.YLeaf{"NexthopAddress", x.NexthopAddress})

    x.EntityData.YListKeys = []string {}

    return &(x.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_Dx4
// End.DX4 SID context
type Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_Dx4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Next Hop Set ID. The type is interface{} with range: 0..4294967295.
    NextHopSetId interface{}
}

func (dx4 *Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_Dx4) GetEntityData() *types.CommonEntityData {
    dx4.EntityData.YFilter = dx4.YFilter
    dx4.EntityData.YangName = "dx4"
    dx4.EntityData.BundleName = "cisco_ios_xr"
    dx4.EntityData.ParentYangName = "key"
    dx4.EntityData.SegmentPath = "dx4"
    dx4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/sid/sid-context/key/" + dx4.EntityData.SegmentPath
    dx4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dx4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dx4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dx4.EntityData.Children = types.NewOrderedMap()
    dx4.EntityData.Leafs = types.NewOrderedMap()
    dx4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dx4.TableId})
    dx4.EntityData.Leafs.Append("next-hop-set-id", types.YLeaf{"NextHopSetId", dx4.NextHopSetId})

    dx4.EntityData.YListKeys = []string {}

    return &(dx4.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_Dt4
// End.DT4 SID context
type Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_Dt4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}
}

func (dt4 *Srv6_Active_Locators_Locator_Sids_Sid_SidContext_Key_Dt4) GetEntityData() *types.CommonEntityData {
    dt4.EntityData.YFilter = dt4.YFilter
    dt4.EntityData.YangName = "dt4"
    dt4.EntityData.BundleName = "cisco_ios_xr"
    dt4.EntityData.ParentYangName = "key"
    dt4.EntityData.SegmentPath = "dt4"
    dt4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/sid/sid-context/key/" + dt4.EntityData.SegmentPath
    dt4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dt4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dt4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dt4.EntityData.Children = types.NewOrderedMap()
    dt4.EntityData.Leafs = types.NewOrderedMap()
    dt4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dt4.TableId})

    dt4.EntityData.YListKeys = []string {}

    return &(dt4.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid_CreateTimestamp
// Creation timestamp
type Srv6_Active_Locators_Locator_Sids_Sid_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Active_Locators_Locator_Sids_Sid_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "sid"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/sid/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Active_Locators_Locator_Sids_Sid_Owner
// Owner
type Srv6_Active_Locators_Locator_Sids_Sid_Owner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Owner. The type is string.
    Owner interface{}
}

func (owner *Srv6_Active_Locators_Locator_Sids_Sid_Owner) GetEntityData() *types.CommonEntityData {
    owner.EntityData.YFilter = owner.YFilter
    owner.EntityData.YangName = "owner"
    owner.EntityData.BundleName = "cisco_ios_xr"
    owner.EntityData.ParentYangName = "sid"
    owner.EntityData.SegmentPath = "owner" + types.AddNoKeyToken(owner)
    owner.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locators/locator/sids/sid/" + owner.EntityData.SegmentPath
    owner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    owner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    owner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    owner.EntityData.Children = types.NewOrderedMap()
    owner.EntityData.Leafs = types.NewOrderedMap()
    owner.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", owner.Owner})

    owner.EntityData.YListKeys = []string {}

    return &(owner.EntityData)
}

// Srv6_Active_LocatorAllSids
// Operational container for all (Active and Stale)
// SIDs across all Locators
type Srv6_Active_LocatorAllSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given locator and SID opcode. The type is slice of
    // Srv6_Active_LocatorAllSids_LocatorAllSid.
    LocatorAllSid []*Srv6_Active_LocatorAllSids_LocatorAllSid
}

func (locatorAllSids *Srv6_Active_LocatorAllSids) GetEntityData() *types.CommonEntityData {
    locatorAllSids.EntityData.YFilter = locatorAllSids.YFilter
    locatorAllSids.EntityData.YangName = "locator-all-sids"
    locatorAllSids.EntityData.BundleName = "cisco_ios_xr"
    locatorAllSids.EntityData.ParentYangName = "active"
    locatorAllSids.EntityData.SegmentPath = "locator-all-sids"
    locatorAllSids.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/" + locatorAllSids.EntityData.SegmentPath
    locatorAllSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllSids.EntityData.Children = types.NewOrderedMap()
    locatorAllSids.EntityData.Children.Append("locator-all-sid", types.YChild{"LocatorAllSid", nil})
    for i := range locatorAllSids.LocatorAllSid {
        locatorAllSids.EntityData.Children.Append(types.GetSegmentPath(locatorAllSids.LocatorAllSid[i]), types.YChild{"LocatorAllSid", locatorAllSids.LocatorAllSid[i]})
    }
    locatorAllSids.EntityData.Leafs = types.NewOrderedMap()

    locatorAllSids.EntityData.YListKeys = []string {}

    return &(locatorAllSids.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid
// Operational data for given locator and SID
// opcode
type Srv6_Active_LocatorAllSids_LocatorAllSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Locator name. The type is string with length:
    // 1..58.
    LocatorName interface{}

    // This attribute is a key. Sid opcode. The type is interface{} with range:
    // 0..4294967295.
    SidOpcode interface{}

    // SID. The type is string.
    Sid interface{}

    // Allocation Type. The type is SidAllocation.
    AllocationType interface{}

    // Function Type. The type is Srv6EndFunction.
    FunctionType interface{}

    // State. The type is SidState.
    State interface{}

    // Rewrite done or not. The type is bool.
    HasForwarding interface{}

    // Associated locator. The type is string.
    Locator interface{}

    // SID Context.
    SidContext Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext

    // Creation timestamp.
    CreateTimestamp Srv6_Active_LocatorAllSids_LocatorAllSid_CreateTimestamp

    // Owner. The type is slice of Srv6_Active_LocatorAllSids_LocatorAllSid_Owner.
    Owner []*Srv6_Active_LocatorAllSids_LocatorAllSid_Owner
}

func (locatorAllSid *Srv6_Active_LocatorAllSids_LocatorAllSid) GetEntityData() *types.CommonEntityData {
    locatorAllSid.EntityData.YFilter = locatorAllSid.YFilter
    locatorAllSid.EntityData.YangName = "locator-all-sid"
    locatorAllSid.EntityData.BundleName = "cisco_ios_xr"
    locatorAllSid.EntityData.ParentYangName = "locator-all-sids"
    locatorAllSid.EntityData.SegmentPath = "locator-all-sid" + types.AddKeyToken(locatorAllSid.LocatorName, "locator-name") + types.AddKeyToken(locatorAllSid.SidOpcode, "sid-opcode")
    locatorAllSid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/" + locatorAllSid.EntityData.SegmentPath
    locatorAllSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllSid.EntityData.Children = types.NewOrderedMap()
    locatorAllSid.EntityData.Children.Append("sid-context", types.YChild{"SidContext", &locatorAllSid.SidContext})
    locatorAllSid.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &locatorAllSid.CreateTimestamp})
    locatorAllSid.EntityData.Children.Append("owner", types.YChild{"Owner", nil})
    for i := range locatorAllSid.Owner {
        types.SetYListKey(locatorAllSid.Owner[i], i)
        locatorAllSid.EntityData.Children.Append(types.GetSegmentPath(locatorAllSid.Owner[i]), types.YChild{"Owner", locatorAllSid.Owner[i]})
    }
    locatorAllSid.EntityData.Leafs = types.NewOrderedMap()
    locatorAllSid.EntityData.Leafs.Append("locator-name", types.YLeaf{"LocatorName", locatorAllSid.LocatorName})
    locatorAllSid.EntityData.Leafs.Append("sid-opcode", types.YLeaf{"SidOpcode", locatorAllSid.SidOpcode})
    locatorAllSid.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", locatorAllSid.Sid})
    locatorAllSid.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", locatorAllSid.AllocationType})
    locatorAllSid.EntityData.Leafs.Append("function-type", types.YLeaf{"FunctionType", locatorAllSid.FunctionType})
    locatorAllSid.EntityData.Leafs.Append("state", types.YLeaf{"State", locatorAllSid.State})
    locatorAllSid.EntityData.Leafs.Append("has-forwarding", types.YLeaf{"HasForwarding", locatorAllSid.HasForwarding})
    locatorAllSid.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", locatorAllSid.Locator})

    locatorAllSid.EntityData.YListKeys = []string {"LocatorName", "SidOpcode"}

    return &(locatorAllSid.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext
// SID Context
type Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application opaque data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ApplicationData interface{}

    // SID Key.
    Key Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key
}

func (sidContext *Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext) GetEntityData() *types.CommonEntityData {
    sidContext.EntityData.YFilter = sidContext.YFilter
    sidContext.EntityData.YangName = "sid-context"
    sidContext.EntityData.BundleName = "cisco_ios_xr"
    sidContext.EntityData.ParentYangName = "locator-all-sid"
    sidContext.EntityData.SegmentPath = "sid-context"
    sidContext.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/locator-all-sid/" + sidContext.EntityData.SegmentPath
    sidContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidContext.EntityData.Children = types.NewOrderedMap()
    sidContext.EntityData.Children.Append("key", types.YChild{"Key", &sidContext.Key})
    sidContext.EntityData.Leafs = types.NewOrderedMap()
    sidContext.EntityData.Leafs.Append("application-data", types.YLeaf{"ApplicationData", sidContext.ApplicationData})

    sidContext.EntityData.YListKeys = []string {}

    return &(sidContext.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key
// SID Key
type Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDContextType. The type is Srv6EndFunction.
    SidContextType interface{}

    // End (PSP) SID context.
    E Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_E

    // End.X (PSP) SID context.
    X Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_X

    // End.DX4 SID context.
    Dx4 Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_Dx4

    // End.DT4 SID context.
    Dt4 Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_Dt4
}

func (key *Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "sid-context"
    key.EntityData.SegmentPath = "key"
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/locator-all-sid/sid-context/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("e", types.YChild{"E", &key.E})
    key.EntityData.Children.Append("x", types.YChild{"X", &key.X})
    key.EntityData.Children.Append("dx4", types.YChild{"Dx4", &key.Dx4})
    key.EntityData.Children.Append("dt4", types.YChild{"Dt4", &key.Dt4})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("sid-context-type", types.YLeaf{"SidContextType", key.SidContextType})

    key.EntityData.YListKeys = []string {}

    return &(key.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_E
// End (PSP) SID context
type Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_E struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Id. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}
}

func (e *Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_E) GetEntityData() *types.CommonEntityData {
    e.EntityData.YFilter = e.YFilter
    e.EntityData.YangName = "e"
    e.EntityData.BundleName = "cisco_ios_xr"
    e.EntityData.ParentYangName = "key"
    e.EntityData.SegmentPath = "e"
    e.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/locator-all-sid/sid-context/key/" + e.EntityData.SegmentPath
    e.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    e.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    e.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    e.EntityData.Children = types.NewOrderedMap()
    e.EntityData.Leafs = types.NewOrderedMap()
    e.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", e.TableId})
    e.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", e.OpaqueId})

    e.EntityData.YListKeys = []string {}

    return &(e.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_X
// End.X (PSP) SID context
type Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_X struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is protected?. The type is bool.
    IsProtected interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}

    // Nexthop interface. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Interface interface{}

    // Nexthop IP address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}
}

func (x *Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_X) GetEntityData() *types.CommonEntityData {
    x.EntityData.YFilter = x.YFilter
    x.EntityData.YangName = "x"
    x.EntityData.BundleName = "cisco_ios_xr"
    x.EntityData.ParentYangName = "key"
    x.EntityData.SegmentPath = "x"
    x.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/locator-all-sid/sid-context/key/" + x.EntityData.SegmentPath
    x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    x.EntityData.Children = types.NewOrderedMap()
    x.EntityData.Leafs = types.NewOrderedMap()
    x.EntityData.Leafs.Append("is-protected", types.YLeaf{"IsProtected", x.IsProtected})
    x.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", x.OpaqueId})
    x.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", x.Interface})
    x.EntityData.Leafs.Append("nexthop-address", types.YLeaf{"NexthopAddress", x.NexthopAddress})

    x.EntityData.YListKeys = []string {}

    return &(x.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_Dx4
// End.DX4 SID context
type Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_Dx4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Next Hop Set ID. The type is interface{} with range: 0..4294967295.
    NextHopSetId interface{}
}

func (dx4 *Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_Dx4) GetEntityData() *types.CommonEntityData {
    dx4.EntityData.YFilter = dx4.YFilter
    dx4.EntityData.YangName = "dx4"
    dx4.EntityData.BundleName = "cisco_ios_xr"
    dx4.EntityData.ParentYangName = "key"
    dx4.EntityData.SegmentPath = "dx4"
    dx4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/locator-all-sid/sid-context/key/" + dx4.EntityData.SegmentPath
    dx4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dx4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dx4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dx4.EntityData.Children = types.NewOrderedMap()
    dx4.EntityData.Leafs = types.NewOrderedMap()
    dx4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dx4.TableId})
    dx4.EntityData.Leafs.Append("next-hop-set-id", types.YLeaf{"NextHopSetId", dx4.NextHopSetId})

    dx4.EntityData.YListKeys = []string {}

    return &(dx4.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_Dt4
// End.DT4 SID context
type Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_Dt4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}
}

func (dt4 *Srv6_Active_LocatorAllSids_LocatorAllSid_SidContext_Key_Dt4) GetEntityData() *types.CommonEntityData {
    dt4.EntityData.YFilter = dt4.YFilter
    dt4.EntityData.YangName = "dt4"
    dt4.EntityData.BundleName = "cisco_ios_xr"
    dt4.EntityData.ParentYangName = "key"
    dt4.EntityData.SegmentPath = "dt4"
    dt4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/locator-all-sid/sid-context/key/" + dt4.EntityData.SegmentPath
    dt4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dt4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dt4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dt4.EntityData.Children = types.NewOrderedMap()
    dt4.EntityData.Leafs = types.NewOrderedMap()
    dt4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dt4.TableId})

    dt4.EntityData.YListKeys = []string {}

    return &(dt4.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid_CreateTimestamp
// Creation timestamp
type Srv6_Active_LocatorAllSids_LocatorAllSid_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Active_LocatorAllSids_LocatorAllSid_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "locator-all-sid"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/locator-all-sid/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Active_LocatorAllSids_LocatorAllSid_Owner
// Owner
type Srv6_Active_LocatorAllSids_LocatorAllSid_Owner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Owner. The type is string.
    Owner interface{}
}

func (owner *Srv6_Active_LocatorAllSids_LocatorAllSid_Owner) GetEntityData() *types.CommonEntityData {
    owner.EntityData.YFilter = owner.YFilter
    owner.EntityData.YangName = "owner"
    owner.EntityData.BundleName = "cisco_ios_xr"
    owner.EntityData.ParentYangName = "locator-all-sid"
    owner.EntityData.SegmentPath = "owner" + types.AddNoKeyToken(owner)
    owner.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-sids/locator-all-sid/" + owner.EntityData.SegmentPath
    owner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    owner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    owner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    owner.EntityData.Children = types.NewOrderedMap()
    owner.EntityData.Leafs = types.NewOrderedMap()
    owner.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", owner.Owner})

    owner.EntityData.YListKeys = []string {}

    return &(owner.EntityData)
}

// Srv6_Active_LocatorAllActiveSids
// Operational container for Active SIDs across all
// Locators
type Srv6_Active_LocatorAllActiveSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given locator and SID opcode. The type is slice of
    // Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid.
    LocatorAllActiveSid []*Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid
}

func (locatorAllActiveSids *Srv6_Active_LocatorAllActiveSids) GetEntityData() *types.CommonEntityData {
    locatorAllActiveSids.EntityData.YFilter = locatorAllActiveSids.YFilter
    locatorAllActiveSids.EntityData.YangName = "locator-all-active-sids"
    locatorAllActiveSids.EntityData.BundleName = "cisco_ios_xr"
    locatorAllActiveSids.EntityData.ParentYangName = "active"
    locatorAllActiveSids.EntityData.SegmentPath = "locator-all-active-sids"
    locatorAllActiveSids.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/" + locatorAllActiveSids.EntityData.SegmentPath
    locatorAllActiveSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllActiveSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllActiveSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllActiveSids.EntityData.Children = types.NewOrderedMap()
    locatorAllActiveSids.EntityData.Children.Append("locator-all-active-sid", types.YChild{"LocatorAllActiveSid", nil})
    for i := range locatorAllActiveSids.LocatorAllActiveSid {
        locatorAllActiveSids.EntityData.Children.Append(types.GetSegmentPath(locatorAllActiveSids.LocatorAllActiveSid[i]), types.YChild{"LocatorAllActiveSid", locatorAllActiveSids.LocatorAllActiveSid[i]})
    }
    locatorAllActiveSids.EntityData.Leafs = types.NewOrderedMap()

    locatorAllActiveSids.EntityData.YListKeys = []string {}

    return &(locatorAllActiveSids.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid
// Operational data for given locator and SID
// opcode
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Locator name. The type is string with length:
    // 1..58.
    LocatorName interface{}

    // This attribute is a key. Sid opcode. The type is interface{} with range:
    // 0..4294967295.
    SidOpcode interface{}

    // SID. The type is string.
    Sid interface{}

    // Allocation Type. The type is SidAllocation.
    AllocationType interface{}

    // Function Type. The type is Srv6EndFunction.
    FunctionType interface{}

    // State. The type is SidState.
    State interface{}

    // Rewrite done or not. The type is bool.
    HasForwarding interface{}

    // Associated locator. The type is string.
    Locator interface{}

    // SID Context.
    SidContext Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext

    // Creation timestamp.
    CreateTimestamp Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_CreateTimestamp

    // Owner. The type is slice of
    // Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_Owner.
    Owner []*Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_Owner
}

func (locatorAllActiveSid *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid) GetEntityData() *types.CommonEntityData {
    locatorAllActiveSid.EntityData.YFilter = locatorAllActiveSid.YFilter
    locatorAllActiveSid.EntityData.YangName = "locator-all-active-sid"
    locatorAllActiveSid.EntityData.BundleName = "cisco_ios_xr"
    locatorAllActiveSid.EntityData.ParentYangName = "locator-all-active-sids"
    locatorAllActiveSid.EntityData.SegmentPath = "locator-all-active-sid" + types.AddKeyToken(locatorAllActiveSid.LocatorName, "locator-name") + types.AddKeyToken(locatorAllActiveSid.SidOpcode, "sid-opcode")
    locatorAllActiveSid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/" + locatorAllActiveSid.EntityData.SegmentPath
    locatorAllActiveSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllActiveSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllActiveSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllActiveSid.EntityData.Children = types.NewOrderedMap()
    locatorAllActiveSid.EntityData.Children.Append("sid-context", types.YChild{"SidContext", &locatorAllActiveSid.SidContext})
    locatorAllActiveSid.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &locatorAllActiveSid.CreateTimestamp})
    locatorAllActiveSid.EntityData.Children.Append("owner", types.YChild{"Owner", nil})
    for i := range locatorAllActiveSid.Owner {
        types.SetYListKey(locatorAllActiveSid.Owner[i], i)
        locatorAllActiveSid.EntityData.Children.Append(types.GetSegmentPath(locatorAllActiveSid.Owner[i]), types.YChild{"Owner", locatorAllActiveSid.Owner[i]})
    }
    locatorAllActiveSid.EntityData.Leafs = types.NewOrderedMap()
    locatorAllActiveSid.EntityData.Leafs.Append("locator-name", types.YLeaf{"LocatorName", locatorAllActiveSid.LocatorName})
    locatorAllActiveSid.EntityData.Leafs.Append("sid-opcode", types.YLeaf{"SidOpcode", locatorAllActiveSid.SidOpcode})
    locatorAllActiveSid.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", locatorAllActiveSid.Sid})
    locatorAllActiveSid.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", locatorAllActiveSid.AllocationType})
    locatorAllActiveSid.EntityData.Leafs.Append("function-type", types.YLeaf{"FunctionType", locatorAllActiveSid.FunctionType})
    locatorAllActiveSid.EntityData.Leafs.Append("state", types.YLeaf{"State", locatorAllActiveSid.State})
    locatorAllActiveSid.EntityData.Leafs.Append("has-forwarding", types.YLeaf{"HasForwarding", locatorAllActiveSid.HasForwarding})
    locatorAllActiveSid.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", locatorAllActiveSid.Locator})

    locatorAllActiveSid.EntityData.YListKeys = []string {"LocatorName", "SidOpcode"}

    return &(locatorAllActiveSid.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext
// SID Context
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application opaque data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ApplicationData interface{}

    // SID Key.
    Key Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key
}

func (sidContext *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext) GetEntityData() *types.CommonEntityData {
    sidContext.EntityData.YFilter = sidContext.YFilter
    sidContext.EntityData.YangName = "sid-context"
    sidContext.EntityData.BundleName = "cisco_ios_xr"
    sidContext.EntityData.ParentYangName = "locator-all-active-sid"
    sidContext.EntityData.SegmentPath = "sid-context"
    sidContext.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/locator-all-active-sid/" + sidContext.EntityData.SegmentPath
    sidContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidContext.EntityData.Children = types.NewOrderedMap()
    sidContext.EntityData.Children.Append("key", types.YChild{"Key", &sidContext.Key})
    sidContext.EntityData.Leafs = types.NewOrderedMap()
    sidContext.EntityData.Leafs.Append("application-data", types.YLeaf{"ApplicationData", sidContext.ApplicationData})

    sidContext.EntityData.YListKeys = []string {}

    return &(sidContext.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key
// SID Key
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDContextType. The type is Srv6EndFunction.
    SidContextType interface{}

    // End (PSP) SID context.
    E Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_E

    // End.X (PSP) SID context.
    X Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_X

    // End.DX4 SID context.
    Dx4 Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dx4

    // End.DT4 SID context.
    Dt4 Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dt4
}

func (key *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "sid-context"
    key.EntityData.SegmentPath = "key"
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/locator-all-active-sid/sid-context/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("e", types.YChild{"E", &key.E})
    key.EntityData.Children.Append("x", types.YChild{"X", &key.X})
    key.EntityData.Children.Append("dx4", types.YChild{"Dx4", &key.Dx4})
    key.EntityData.Children.Append("dt4", types.YChild{"Dt4", &key.Dt4})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("sid-context-type", types.YLeaf{"SidContextType", key.SidContextType})

    key.EntityData.YListKeys = []string {}

    return &(key.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_E
// End (PSP) SID context
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_E struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Id. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}
}

func (e *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_E) GetEntityData() *types.CommonEntityData {
    e.EntityData.YFilter = e.YFilter
    e.EntityData.YangName = "e"
    e.EntityData.BundleName = "cisco_ios_xr"
    e.EntityData.ParentYangName = "key"
    e.EntityData.SegmentPath = "e"
    e.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/locator-all-active-sid/sid-context/key/" + e.EntityData.SegmentPath
    e.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    e.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    e.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    e.EntityData.Children = types.NewOrderedMap()
    e.EntityData.Leafs = types.NewOrderedMap()
    e.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", e.TableId})
    e.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", e.OpaqueId})

    e.EntityData.YListKeys = []string {}

    return &(e.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_X
// End.X (PSP) SID context
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_X struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is protected?. The type is bool.
    IsProtected interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}

    // Nexthop interface. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Interface interface{}

    // Nexthop IP address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}
}

func (x *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_X) GetEntityData() *types.CommonEntityData {
    x.EntityData.YFilter = x.YFilter
    x.EntityData.YangName = "x"
    x.EntityData.BundleName = "cisco_ios_xr"
    x.EntityData.ParentYangName = "key"
    x.EntityData.SegmentPath = "x"
    x.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/locator-all-active-sid/sid-context/key/" + x.EntityData.SegmentPath
    x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    x.EntityData.Children = types.NewOrderedMap()
    x.EntityData.Leafs = types.NewOrderedMap()
    x.EntityData.Leafs.Append("is-protected", types.YLeaf{"IsProtected", x.IsProtected})
    x.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", x.OpaqueId})
    x.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", x.Interface})
    x.EntityData.Leafs.Append("nexthop-address", types.YLeaf{"NexthopAddress", x.NexthopAddress})

    x.EntityData.YListKeys = []string {}

    return &(x.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dx4
// End.DX4 SID context
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dx4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Next Hop Set ID. The type is interface{} with range: 0..4294967295.
    NextHopSetId interface{}
}

func (dx4 *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dx4) GetEntityData() *types.CommonEntityData {
    dx4.EntityData.YFilter = dx4.YFilter
    dx4.EntityData.YangName = "dx4"
    dx4.EntityData.BundleName = "cisco_ios_xr"
    dx4.EntityData.ParentYangName = "key"
    dx4.EntityData.SegmentPath = "dx4"
    dx4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/locator-all-active-sid/sid-context/key/" + dx4.EntityData.SegmentPath
    dx4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dx4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dx4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dx4.EntityData.Children = types.NewOrderedMap()
    dx4.EntityData.Leafs = types.NewOrderedMap()
    dx4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dx4.TableId})
    dx4.EntityData.Leafs.Append("next-hop-set-id", types.YLeaf{"NextHopSetId", dx4.NextHopSetId})

    dx4.EntityData.YListKeys = []string {}

    return &(dx4.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dt4
// End.DT4 SID context
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dt4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}
}

func (dt4 *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dt4) GetEntityData() *types.CommonEntityData {
    dt4.EntityData.YFilter = dt4.YFilter
    dt4.EntityData.YangName = "dt4"
    dt4.EntityData.BundleName = "cisco_ios_xr"
    dt4.EntityData.ParentYangName = "key"
    dt4.EntityData.SegmentPath = "dt4"
    dt4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/locator-all-active-sid/sid-context/key/" + dt4.EntityData.SegmentPath
    dt4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dt4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dt4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dt4.EntityData.Children = types.NewOrderedMap()
    dt4.EntityData.Leafs = types.NewOrderedMap()
    dt4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dt4.TableId})

    dt4.EntityData.YListKeys = []string {}

    return &(dt4.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_CreateTimestamp
// Creation timestamp
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "locator-all-active-sid"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/locator-all-active-sid/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_Owner
// Owner
type Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_Owner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Owner. The type is string.
    Owner interface{}
}

func (owner *Srv6_Active_LocatorAllActiveSids_LocatorAllActiveSid_Owner) GetEntityData() *types.CommonEntityData {
    owner.EntityData.YFilter = owner.YFilter
    owner.EntityData.YangName = "owner"
    owner.EntityData.BundleName = "cisco_ios_xr"
    owner.EntityData.ParentYangName = "locator-all-active-sid"
    owner.EntityData.SegmentPath = "owner" + types.AddNoKeyToken(owner)
    owner.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/active/locator-all-active-sids/locator-all-active-sid/" + owner.EntityData.SegmentPath
    owner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    owner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    owner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    owner.EntityData.Children = types.NewOrderedMap()
    owner.EntityData.Leafs = types.NewOrderedMap()
    owner.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", owner.Owner})

    owner.EntityData.YListKeys = []string {}

    return &(owner.EntityData)
}

// Srv6_Standby
// Standby SRv6 operational data
type Srv6_Standby struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID Manager information.
    Manager Srv6_Standby_Manager

    // SRv6 locators related information.
    Locators Srv6_Standby_Locators

    // Operational container for all (Active and Stale) SIDs across all Locators.
    LocatorAllSids Srv6_Standby_LocatorAllSids

    // Operational container for Active SIDs across all Locators.
    LocatorAllActiveSids Srv6_Standby_LocatorAllActiveSids
}

func (standby *Srv6_Standby) GetEntityData() *types.CommonEntityData {
    standby.EntityData.YFilter = standby.YFilter
    standby.EntityData.YangName = "standby"
    standby.EntityData.BundleName = "cisco_ios_xr"
    standby.EntityData.ParentYangName = "srv6"
    standby.EntityData.SegmentPath = "standby"
    standby.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/" + standby.EntityData.SegmentPath
    standby.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standby.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standby.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standby.EntityData.Children = types.NewOrderedMap()
    standby.EntityData.Children.Append("manager", types.YChild{"Manager", &standby.Manager})
    standby.EntityData.Children.Append("locators", types.YChild{"Locators", &standby.Locators})
    standby.EntityData.Children.Append("locator-all-sids", types.YChild{"LocatorAllSids", &standby.LocatorAllSids})
    standby.EntityData.Children.Append("locator-all-active-sids", types.YChild{"LocatorAllActiveSids", &standby.LocatorAllActiveSids})
    standby.EntityData.Leafs = types.NewOrderedMap()

    standby.EntityData.YListKeys = []string {}

    return &(standby.EntityData)
}

// Srv6_Standby_Manager
// SID Manager information
type Srv6_Standby_Manager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID Mgr parameters.
    SidMgrParams Srv6_Standby_Manager_SidMgrParams

    // SID Mgr summary info.
    SidMgrSummary Srv6_Standby_Manager_SidMgrSummary

    // Platform Capabilities.
    PlatformCapabilities Srv6_Standby_Manager_PlatformCapabilities
}

func (manager *Srv6_Standby_Manager) GetEntityData() *types.CommonEntityData {
    manager.EntityData.YFilter = manager.YFilter
    manager.EntityData.YangName = "manager"
    manager.EntityData.BundleName = "cisco_ios_xr"
    manager.EntityData.ParentYangName = "standby"
    manager.EntityData.SegmentPath = "manager"
    manager.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/" + manager.EntityData.SegmentPath
    manager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manager.EntityData.Children = types.NewOrderedMap()
    manager.EntityData.Children.Append("sid-mgr-params", types.YChild{"SidMgrParams", &manager.SidMgrParams})
    manager.EntityData.Children.Append("sid-mgr-summary", types.YChild{"SidMgrSummary", &manager.SidMgrSummary})
    manager.EntityData.Children.Append("platform-capabilities", types.YChild{"PlatformCapabilities", &manager.PlatformCapabilities})
    manager.EntityData.Leafs = types.NewOrderedMap()

    manager.EntityData.YListKeys = []string {}

    return &(manager.EntityData)
}

// Srv6_Standby_Manager_SidMgrParams
// SID Mgr parameters
type Srv6_Standby_Manager_SidMgrParams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SRv6 enabled?. The type is bool.
    Srv6Enabled interface{}

    // Configured Encap Source address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ConfiguredEncapSourceAddress interface{}

    // Default Encap Source address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DefaultEncapSourceAddress interface{}

    // Is TTL propagate enabled?. The type is bool.
    EncapTtlPropagate interface{}

    // Is SID Holdtime configured?. The type is bool.
    IsSidHoldtimeConfigured interface{}

    // Configured SID Holdtime in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    SidHoldtimeMinsConfigured interface{}

    // Encap Hop-limit info.
    EncapHopLimit Srv6_Standby_Manager_SidMgrParams_EncapHopLimit
}

func (sidMgrParams *Srv6_Standby_Manager_SidMgrParams) GetEntityData() *types.CommonEntityData {
    sidMgrParams.EntityData.YFilter = sidMgrParams.YFilter
    sidMgrParams.EntityData.YangName = "sid-mgr-params"
    sidMgrParams.EntityData.BundleName = "cisco_ios_xr"
    sidMgrParams.EntityData.ParentYangName = "manager"
    sidMgrParams.EntityData.SegmentPath = "sid-mgr-params"
    sidMgrParams.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/" + sidMgrParams.EntityData.SegmentPath
    sidMgrParams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidMgrParams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidMgrParams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidMgrParams.EntityData.Children = types.NewOrderedMap()
    sidMgrParams.EntityData.Children.Append("encap-hop-limit", types.YChild{"EncapHopLimit", &sidMgrParams.EncapHopLimit})
    sidMgrParams.EntityData.Leafs = types.NewOrderedMap()
    sidMgrParams.EntityData.Leafs.Append("srv6-enabled", types.YLeaf{"Srv6Enabled", sidMgrParams.Srv6Enabled})
    sidMgrParams.EntityData.Leafs.Append("configured-encap-source-address", types.YLeaf{"ConfiguredEncapSourceAddress", sidMgrParams.ConfiguredEncapSourceAddress})
    sidMgrParams.EntityData.Leafs.Append("default-encap-source-address", types.YLeaf{"DefaultEncapSourceAddress", sidMgrParams.DefaultEncapSourceAddress})
    sidMgrParams.EntityData.Leafs.Append("encap-ttl-propagate", types.YLeaf{"EncapTtlPropagate", sidMgrParams.EncapTtlPropagate})
    sidMgrParams.EntityData.Leafs.Append("is-sid-holdtime-configured", types.YLeaf{"IsSidHoldtimeConfigured", sidMgrParams.IsSidHoldtimeConfigured})
    sidMgrParams.EntityData.Leafs.Append("sid-holdtime-mins-configured", types.YLeaf{"SidHoldtimeMinsConfigured", sidMgrParams.SidHoldtimeMinsConfigured})

    sidMgrParams.EntityData.YListKeys = []string {}

    return &(sidMgrParams.EntityData)
}

// Srv6_Standby_Manager_SidMgrParams_EncapHopLimit
// Encap Hop-limit info
type Srv6_Standby_Manager_SidMgrParams_EncapHopLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use default IPv6 hop-limit value. The type is bool.
    UseDefault interface{}

    // Propagate IP TTL to Encap IPv6 hop-limit. The type is bool.
    DoPropagate interface{}

    // Specific value set for hop-limit count. The type is interface{} with range:
    // 0..255.
    Value interface{}
}

func (encapHopLimit *Srv6_Standby_Manager_SidMgrParams_EncapHopLimit) GetEntityData() *types.CommonEntityData {
    encapHopLimit.EntityData.YFilter = encapHopLimit.YFilter
    encapHopLimit.EntityData.YangName = "encap-hop-limit"
    encapHopLimit.EntityData.BundleName = "cisco_ios_xr"
    encapHopLimit.EntityData.ParentYangName = "sid-mgr-params"
    encapHopLimit.EntityData.SegmentPath = "encap-hop-limit"
    encapHopLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/sid-mgr-params/" + encapHopLimit.EntityData.SegmentPath
    encapHopLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encapHopLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encapHopLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encapHopLimit.EntityData.Children = types.NewOrderedMap()
    encapHopLimit.EntityData.Leafs = types.NewOrderedMap()
    encapHopLimit.EntityData.Leafs.Append("use-default", types.YLeaf{"UseDefault", encapHopLimit.UseDefault})
    encapHopLimit.EntityData.Leafs.Append("do-propagate", types.YLeaf{"DoPropagate", encapHopLimit.DoPropagate})
    encapHopLimit.EntityData.Leafs.Append("value", types.YLeaf{"Value", encapHopLimit.Value})

    encapHopLimit.EntityData.YListKeys = []string {}

    return &(encapHopLimit.EntityData)
}

// Srv6_Standby_Manager_SidMgrSummary
// SID Mgr summary info
type Srv6_Standby_Manager_SidMgrSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of locators. The type is interface{} with range: 0..65535.
    LocatorsCount interface{}

    // Number of operational locators. The type is interface{} with range:
    // 0..65535.
    OperLocatorsCount interface{}

    // Number of SIDs. The type is interface{} with range: 0..4294967295.
    SidsCount interface{}

    // Number of Stale SIDs. The type is interface{} with range: 0..4294967295.
    StaleSidsCount interface{}

    // Global Maximum number of SIDs. The type is interface{} with range:
    // 0..4294967295.
    MaximumSidsCount interface{}

    // SIDs Global Out of Resource info.
    SidsOutOfResourceSummary Srv6_Standby_Manager_SidMgrSummary_SidsOutOfResourceSummary
}

func (sidMgrSummary *Srv6_Standby_Manager_SidMgrSummary) GetEntityData() *types.CommonEntityData {
    sidMgrSummary.EntityData.YFilter = sidMgrSummary.YFilter
    sidMgrSummary.EntityData.YangName = "sid-mgr-summary"
    sidMgrSummary.EntityData.BundleName = "cisco_ios_xr"
    sidMgrSummary.EntityData.ParentYangName = "manager"
    sidMgrSummary.EntityData.SegmentPath = "sid-mgr-summary"
    sidMgrSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/" + sidMgrSummary.EntityData.SegmentPath
    sidMgrSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidMgrSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidMgrSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidMgrSummary.EntityData.Children = types.NewOrderedMap()
    sidMgrSummary.EntityData.Children.Append("sids-out-of-resource-summary", types.YChild{"SidsOutOfResourceSummary", &sidMgrSummary.SidsOutOfResourceSummary})
    sidMgrSummary.EntityData.Leafs = types.NewOrderedMap()
    sidMgrSummary.EntityData.Leafs.Append("locators-count", types.YLeaf{"LocatorsCount", sidMgrSummary.LocatorsCount})
    sidMgrSummary.EntityData.Leafs.Append("oper-locators-count", types.YLeaf{"OperLocatorsCount", sidMgrSummary.OperLocatorsCount})
    sidMgrSummary.EntityData.Leafs.Append("sids-count", types.YLeaf{"SidsCount", sidMgrSummary.SidsCount})
    sidMgrSummary.EntityData.Leafs.Append("stale-sids-count", types.YLeaf{"StaleSidsCount", sidMgrSummary.StaleSidsCount})
    sidMgrSummary.EntityData.Leafs.Append("maximum-sids-count", types.YLeaf{"MaximumSidsCount", sidMgrSummary.MaximumSidsCount})

    sidMgrSummary.EntityData.YListKeys = []string {}

    return &(sidMgrSummary.EntityData)
}

// Srv6_Standby_Manager_SidMgrSummary_SidsOutOfResourceSummary
// SIDs Global Out of Resource info
type Srv6_Standby_Manager_SidMgrSummary_SidsOutOfResourceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global Resources State for SIDs. The type is Srv6OutOfResourceState.
    OutOfResourcesState interface{}

    // Threshold for Number of Free SID below which OOR Yellow State is reached.
    // The type is interface{} with range: 0..4294967295.
    OorYellowFreeSidThreshold interface{}

    // Threshold for Number of Free SID above which OOR Green State is restored.
    // The type is interface{} with range: 0..4294967295.
    OorGreenFreeSidThreshold interface{}

    // Number of times Resources Warning or Out of Resources state has been
    // cleared. The type is interface{} with range: 0..4294967295.
    OorGreenCount interface{}

    // Number of times system went into Resources Warning state. The type is
    // interface{} with range: 0..4294967295.
    OorYellowCount interface{}

    // Number of times system went into Out of Resources state. The type is
    // interface{} with range: 0..4294967295.
    OorRedCount interface{}
}

func (sidsOutOfResourceSummary *Srv6_Standby_Manager_SidMgrSummary_SidsOutOfResourceSummary) GetEntityData() *types.CommonEntityData {
    sidsOutOfResourceSummary.EntityData.YFilter = sidsOutOfResourceSummary.YFilter
    sidsOutOfResourceSummary.EntityData.YangName = "sids-out-of-resource-summary"
    sidsOutOfResourceSummary.EntityData.BundleName = "cisco_ios_xr"
    sidsOutOfResourceSummary.EntityData.ParentYangName = "sid-mgr-summary"
    sidsOutOfResourceSummary.EntityData.SegmentPath = "sids-out-of-resource-summary"
    sidsOutOfResourceSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/sid-mgr-summary/" + sidsOutOfResourceSummary.EntityData.SegmentPath
    sidsOutOfResourceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidsOutOfResourceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidsOutOfResourceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidsOutOfResourceSummary.EntityData.Children = types.NewOrderedMap()
    sidsOutOfResourceSummary.EntityData.Leafs = types.NewOrderedMap()
    sidsOutOfResourceSummary.EntityData.Leafs.Append("out-of-resources-state", types.YLeaf{"OutOfResourcesState", sidsOutOfResourceSummary.OutOfResourcesState})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-yellow-free-sid-threshold", types.YLeaf{"OorYellowFreeSidThreshold", sidsOutOfResourceSummary.OorYellowFreeSidThreshold})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-green-free-sid-threshold", types.YLeaf{"OorGreenFreeSidThreshold", sidsOutOfResourceSummary.OorGreenFreeSidThreshold})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-green-count", types.YLeaf{"OorGreenCount", sidsOutOfResourceSummary.OorGreenCount})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-yellow-count", types.YLeaf{"OorYellowCount", sidsOutOfResourceSummary.OorYellowCount})
    sidsOutOfResourceSummary.EntityData.Leafs.Append("oor-red-count", types.YLeaf{"OorRedCount", sidsOutOfResourceSummary.OorRedCount})

    sidsOutOfResourceSummary.EntityData.YListKeys = []string {}

    return &(sidsOutOfResourceSummary.EntityData)
}

// Srv6_Standby_Manager_PlatformCapabilities
// Platform Capabilities
type Srv6_Standby_Manager_PlatformCapabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum Sids. The type is interface{} with range: 0..4294967295.
    MaxSid interface{}

    // Freed SID holdtime in mins. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    SidHoldtimeMins interface{}

    // Feature support.
    Support Srv6_Standby_Manager_PlatformCapabilities_Support
}

func (platformCapabilities *Srv6_Standby_Manager_PlatformCapabilities) GetEntityData() *types.CommonEntityData {
    platformCapabilities.EntityData.YFilter = platformCapabilities.YFilter
    platformCapabilities.EntityData.YangName = "platform-capabilities"
    platformCapabilities.EntityData.BundleName = "cisco_ios_xr"
    platformCapabilities.EntityData.ParentYangName = "manager"
    platformCapabilities.EntityData.SegmentPath = "platform-capabilities"
    platformCapabilities.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/" + platformCapabilities.EntityData.SegmentPath
    platformCapabilities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformCapabilities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformCapabilities.EntityData.Children = types.NewOrderedMap()
    platformCapabilities.EntityData.Children.Append("support", types.YChild{"Support", &platformCapabilities.Support})
    platformCapabilities.EntityData.Leafs = types.NewOrderedMap()
    platformCapabilities.EntityData.Leafs.Append("max-sid", types.YLeaf{"MaxSid", platformCapabilities.MaxSid})
    platformCapabilities.EntityData.Leafs.Append("sid-holdtime-mins", types.YLeaf{"SidHoldtimeMins", platformCapabilities.SidHoldtimeMins})

    platformCapabilities.EntityData.YListKeys = []string {}

    return &(platformCapabilities.EntityData)
}

// Srv6_Standby_Manager_PlatformCapabilities_Support
// Feature support
type Srv6_Standby_Manager_PlatformCapabilities_Support struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRv6 support. The type is bool.
    Srv6 interface{}

    // TI LFA support. The type is bool.
    Tilfa interface{}

    // Microloop-avoidance support. The type is bool.
    MicroloopAvoidance interface{}

    // Signaled Parameters.
    SignaledParameters Srv6_Standby_Manager_PlatformCapabilities_Support_SignaledParameters

    // Supported end functions. The type is slice of
    // Srv6_Standby_Manager_PlatformCapabilities_Support_EndFunc.
    EndFunc []*Srv6_Standby_Manager_PlatformCapabilities_Support_EndFunc

    // Supported Transit functions. The type is slice of
    // Srv6_Standby_Manager_PlatformCapabilities_Support_TransitFunc.
    TransitFunc []*Srv6_Standby_Manager_PlatformCapabilities_Support_TransitFunc

    // Security rules. The type is slice of
    // Srv6_Standby_Manager_PlatformCapabilities_Support_SecurityRule.
    SecurityRule []*Srv6_Standby_Manager_PlatformCapabilities_Support_SecurityRule

    // Counters. The type is slice of
    // Srv6_Standby_Manager_PlatformCapabilities_Support_Counter.
    Counter []*Srv6_Standby_Manager_PlatformCapabilities_Support_Counter
}

func (support *Srv6_Standby_Manager_PlatformCapabilities_Support) GetEntityData() *types.CommonEntityData {
    support.EntityData.YFilter = support.YFilter
    support.EntityData.YangName = "support"
    support.EntityData.BundleName = "cisco_ios_xr"
    support.EntityData.ParentYangName = "platform-capabilities"
    support.EntityData.SegmentPath = "support"
    support.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/platform-capabilities/" + support.EntityData.SegmentPath
    support.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    support.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    support.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    support.EntityData.Children = types.NewOrderedMap()
    support.EntityData.Children.Append("signaled-parameters", types.YChild{"SignaledParameters", &support.SignaledParameters})
    support.EntityData.Children.Append("end-func", types.YChild{"EndFunc", nil})
    for i := range support.EndFunc {
        types.SetYListKey(support.EndFunc[i], i)
        support.EntityData.Children.Append(types.GetSegmentPath(support.EndFunc[i]), types.YChild{"EndFunc", support.EndFunc[i]})
    }
    support.EntityData.Children.Append("transit-func", types.YChild{"TransitFunc", nil})
    for i := range support.TransitFunc {
        types.SetYListKey(support.TransitFunc[i], i)
        support.EntityData.Children.Append(types.GetSegmentPath(support.TransitFunc[i]), types.YChild{"TransitFunc", support.TransitFunc[i]})
    }
    support.EntityData.Children.Append("security-rule", types.YChild{"SecurityRule", nil})
    for i := range support.SecurityRule {
        types.SetYListKey(support.SecurityRule[i], i)
        support.EntityData.Children.Append(types.GetSegmentPath(support.SecurityRule[i]), types.YChild{"SecurityRule", support.SecurityRule[i]})
    }
    support.EntityData.Children.Append("counter", types.YChild{"Counter", nil})
    for i := range support.Counter {
        types.SetYListKey(support.Counter[i], i)
        support.EntityData.Children.Append(types.GetSegmentPath(support.Counter[i]), types.YChild{"Counter", support.Counter[i]})
    }
    support.EntityData.Leafs = types.NewOrderedMap()
    support.EntityData.Leafs.Append("srv6", types.YLeaf{"Srv6", support.Srv6})
    support.EntityData.Leafs.Append("tilfa", types.YLeaf{"Tilfa", support.Tilfa})
    support.EntityData.Leafs.Append("microloop-avoidance", types.YLeaf{"MicroloopAvoidance", support.MicroloopAvoidance})

    support.EntityData.YListKeys = []string {}

    return &(support.EntityData)
}

// Srv6_Standby_Manager_PlatformCapabilities_Support_SignaledParameters
// Signaled Parameters
type Srv6_Standby_Manager_PlatformCapabilities_Support_SignaledParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max value of SegmentLeft field in received SRH. The type is interface{}
    // with range: 0..255.
    MaxSl interface{}

    // Max num of SIDs in rcvd SRH for pop. The type is interface{} with range:
    // 0..255.
    MaxEndPopSrh interface{}

    // Max num of SIDs for T.Insert op. The type is interface{} with range:
    // 0..255.
    MaxTInsert interface{}

    // Max num of SIDs for T.Encaps op. The type is interface{} with range:
    // 0..255.
    MaxTEncap interface{}

    // Max num of SIDs in rcvd SRH for decap. The type is interface{} with range:
    // 0..255.
    MaxEndD interface{}
}

func (signaledParameters *Srv6_Standby_Manager_PlatformCapabilities_Support_SignaledParameters) GetEntityData() *types.CommonEntityData {
    signaledParameters.EntityData.YFilter = signaledParameters.YFilter
    signaledParameters.EntityData.YangName = "signaled-parameters"
    signaledParameters.EntityData.BundleName = "cisco_ios_xr"
    signaledParameters.EntityData.ParentYangName = "support"
    signaledParameters.EntityData.SegmentPath = "signaled-parameters"
    signaledParameters.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/platform-capabilities/support/" + signaledParameters.EntityData.SegmentPath
    signaledParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signaledParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signaledParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signaledParameters.EntityData.Children = types.NewOrderedMap()
    signaledParameters.EntityData.Leafs = types.NewOrderedMap()
    signaledParameters.EntityData.Leafs.Append("max-sl", types.YLeaf{"MaxSl", signaledParameters.MaxSl})
    signaledParameters.EntityData.Leafs.Append("max-end-pop-srh", types.YLeaf{"MaxEndPopSrh", signaledParameters.MaxEndPopSrh})
    signaledParameters.EntityData.Leafs.Append("max-t-insert", types.YLeaf{"MaxTInsert", signaledParameters.MaxTInsert})
    signaledParameters.EntityData.Leafs.Append("max-t-encap", types.YLeaf{"MaxTEncap", signaledParameters.MaxTEncap})
    signaledParameters.EntityData.Leafs.Append("max-end-d", types.YLeaf{"MaxEndD", signaledParameters.MaxEndD})

    signaledParameters.EntityData.YListKeys = []string {}

    return &(signaledParameters.EntityData)
}

// Srv6_Standby_Manager_PlatformCapabilities_Support_EndFunc
// Supported end functions
type Srv6_Standby_Manager_PlatformCapabilities_Support_EndFunc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // String. The type is string.
    String interface{}
}

func (endFunc *Srv6_Standby_Manager_PlatformCapabilities_Support_EndFunc) GetEntityData() *types.CommonEntityData {
    endFunc.EntityData.YFilter = endFunc.YFilter
    endFunc.EntityData.YangName = "end-func"
    endFunc.EntityData.BundleName = "cisco_ios_xr"
    endFunc.EntityData.ParentYangName = "support"
    endFunc.EntityData.SegmentPath = "end-func" + types.AddNoKeyToken(endFunc)
    endFunc.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/platform-capabilities/support/" + endFunc.EntityData.SegmentPath
    endFunc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    endFunc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    endFunc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    endFunc.EntityData.Children = types.NewOrderedMap()
    endFunc.EntityData.Leafs = types.NewOrderedMap()
    endFunc.EntityData.Leafs.Append("string", types.YLeaf{"String", endFunc.String})

    endFunc.EntityData.YListKeys = []string {}

    return &(endFunc.EntityData)
}

// Srv6_Standby_Manager_PlatformCapabilities_Support_TransitFunc
// Supported Transit functions
type Srv6_Standby_Manager_PlatformCapabilities_Support_TransitFunc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // String. The type is string.
    String interface{}
}

func (transitFunc *Srv6_Standby_Manager_PlatformCapabilities_Support_TransitFunc) GetEntityData() *types.CommonEntityData {
    transitFunc.EntityData.YFilter = transitFunc.YFilter
    transitFunc.EntityData.YangName = "transit-func"
    transitFunc.EntityData.BundleName = "cisco_ios_xr"
    transitFunc.EntityData.ParentYangName = "support"
    transitFunc.EntityData.SegmentPath = "transit-func" + types.AddNoKeyToken(transitFunc)
    transitFunc.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/platform-capabilities/support/" + transitFunc.EntityData.SegmentPath
    transitFunc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transitFunc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transitFunc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transitFunc.EntityData.Children = types.NewOrderedMap()
    transitFunc.EntityData.Leafs = types.NewOrderedMap()
    transitFunc.EntityData.Leafs.Append("string", types.YLeaf{"String", transitFunc.String})

    transitFunc.EntityData.YListKeys = []string {}

    return &(transitFunc.EntityData)
}

// Srv6_Standby_Manager_PlatformCapabilities_Support_SecurityRule
// Security rules
type Srv6_Standby_Manager_PlatformCapabilities_Support_SecurityRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // String. The type is string.
    String interface{}
}

func (securityRule *Srv6_Standby_Manager_PlatformCapabilities_Support_SecurityRule) GetEntityData() *types.CommonEntityData {
    securityRule.EntityData.YFilter = securityRule.YFilter
    securityRule.EntityData.YangName = "security-rule"
    securityRule.EntityData.BundleName = "cisco_ios_xr"
    securityRule.EntityData.ParentYangName = "support"
    securityRule.EntityData.SegmentPath = "security-rule" + types.AddNoKeyToken(securityRule)
    securityRule.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/platform-capabilities/support/" + securityRule.EntityData.SegmentPath
    securityRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    securityRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    securityRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    securityRule.EntityData.Children = types.NewOrderedMap()
    securityRule.EntityData.Leafs = types.NewOrderedMap()
    securityRule.EntityData.Leafs.Append("string", types.YLeaf{"String", securityRule.String})

    securityRule.EntityData.YListKeys = []string {}

    return &(securityRule.EntityData)
}

// Srv6_Standby_Manager_PlatformCapabilities_Support_Counter
// Counters
type Srv6_Standby_Manager_PlatformCapabilities_Support_Counter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // String. The type is string.
    String interface{}
}

func (counter *Srv6_Standby_Manager_PlatformCapabilities_Support_Counter) GetEntityData() *types.CommonEntityData {
    counter.EntityData.YFilter = counter.YFilter
    counter.EntityData.YangName = "counter"
    counter.EntityData.BundleName = "cisco_ios_xr"
    counter.EntityData.ParentYangName = "support"
    counter.EntityData.SegmentPath = "counter" + types.AddNoKeyToken(counter)
    counter.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/manager/platform-capabilities/support/" + counter.EntityData.SegmentPath
    counter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counter.EntityData.Children = types.NewOrderedMap()
    counter.EntityData.Leafs = types.NewOrderedMap()
    counter.EntityData.Leafs.Append("string", types.YLeaf{"String", counter.String})

    counter.EntityData.YListKeys = []string {}

    return &(counter.EntityData)
}

// Srv6_Standby_Locators
// SRv6 locators related information
type Srv6_Standby_Locators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given SRv6 locator. The type is slice of
    // Srv6_Standby_Locators_Locator.
    Locator []*Srv6_Standby_Locators_Locator
}

func (locators *Srv6_Standby_Locators) GetEntityData() *types.CommonEntityData {
    locators.EntityData.YFilter = locators.YFilter
    locators.EntityData.YangName = "locators"
    locators.EntityData.BundleName = "cisco_ios_xr"
    locators.EntityData.ParentYangName = "standby"
    locators.EntityData.SegmentPath = "locators"
    locators.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/" + locators.EntityData.SegmentPath
    locators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locators.EntityData.Children = types.NewOrderedMap()
    locators.EntityData.Children.Append("locator", types.YChild{"Locator", nil})
    for i := range locators.Locator {
        locators.EntityData.Children.Append(types.GetSegmentPath(locators.Locator[i]), types.YChild{"Locator", locators.Locator[i]})
    }
    locators.EntityData.Leafs = types.NewOrderedMap()

    locators.EntityData.YListKeys = []string {}

    return &(locators.EntityData)
}

// Srv6_Standby_Locators_Locator
// Operational data for given SRv6 locator
type Srv6_Standby_Locators_Locator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Locator name. The type is string with length:
    // 1..58.
    Name interface{}

    // Operational data for given SRv6 locator.
    Info Srv6_Standby_Locators_Locator_Info

    // SRv6 locator SID table.
    Sids Srv6_Standby_Locators_Locator_Sids
}

func (locator *Srv6_Standby_Locators_Locator) GetEntityData() *types.CommonEntityData {
    locator.EntityData.YFilter = locator.YFilter
    locator.EntityData.YangName = "locator"
    locator.EntityData.BundleName = "cisco_ios_xr"
    locator.EntityData.ParentYangName = "locators"
    locator.EntityData.SegmentPath = "locator" + types.AddKeyToken(locator.Name, "name")
    locator.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/" + locator.EntityData.SegmentPath
    locator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locator.EntityData.Children = types.NewOrderedMap()
    locator.EntityData.Children.Append("info", types.YChild{"Info", &locator.Info})
    locator.EntityData.Children.Append("sids", types.YChild{"Sids", &locator.Sids})
    locator.EntityData.Leafs = types.NewOrderedMap()
    locator.EntityData.Leafs.Append("name", types.YLeaf{"Name", locator.Name})

    locator.EntityData.YListKeys = []string {"Name"}

    return &(locator.EntityData)
}

// Srv6_Standby_Locators_Locator_Info
// Operational data for given SRv6 locator
type Srv6_Standby_Locators_Locator_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Locator Name. The type is string.
    Name interface{}

    // Locator ID. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // Locator Prefix. The type is string.
    Prefix interface{}

    // Locator status is Up or Down. The type is bool.
    IsOperational interface{}

    // Locator is the default locator. The type is bool.
    IsDefault interface{}

    // Locator Resources State for SIDs. The type is Srv6OutOfResourceState.
    OutOfResourcesState interface{}

    // Locator IM intf info.
    Interface Srv6_Standby_Locators_Locator_Info_Interface

    // Creation timestamp.
    CreateTimestamp Srv6_Standby_Locators_Locator_Info_CreateTimestamp
}

func (info *Srv6_Standby_Locators_Locator_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "locator"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/" + info.EntityData.SegmentPath
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("interface", types.YChild{"Interface", &info.Interface})
    info.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &info.CreateTimestamp})
    info.EntityData.Leafs = types.NewOrderedMap()
    info.EntityData.Leafs.Append("name", types.YLeaf{"Name", info.Name})
    info.EntityData.Leafs.Append("id", types.YLeaf{"Id", info.Id})
    info.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", info.Prefix})
    info.EntityData.Leafs.Append("is-operational", types.YLeaf{"IsOperational", info.IsOperational})
    info.EntityData.Leafs.Append("is-default", types.YLeaf{"IsDefault", info.IsDefault})
    info.EntityData.Leafs.Append("out-of-resources-state", types.YLeaf{"OutOfResourcesState", info.OutOfResourcesState})

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Srv6_Standby_Locators_Locator_Info_Interface
// Locator IM intf info
type Srv6_Standby_Locators_Locator_Info_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string.
    Name interface{}

    // Interface handle. The type is string with pattern: b'[0-9a-fA-F]{1,8}'.
    IfHandle interface{}

    // Interface prefix/addr programmed. The type is string.
    ProgrammedPrefix interface{}
}

func (self *Srv6_Standby_Locators_Locator_Info_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "info"
    self.EntityData.SegmentPath = "interface"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/info/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})
    self.EntityData.Leafs.Append("if-handle", types.YLeaf{"IfHandle", self.IfHandle})
    self.EntityData.Leafs.Append("programmed-prefix", types.YLeaf{"ProgrammedPrefix", self.ProgrammedPrefix})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Srv6_Standby_Locators_Locator_Info_CreateTimestamp
// Creation timestamp
type Srv6_Standby_Locators_Locator_Info_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Standby_Locators_Locator_Info_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "info"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/info/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids
// SRv6 locator SID table
type Srv6_Standby_Locators_Locator_Sids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given SRv6 SID. The type is slice of
    // Srv6_Standby_Locators_Locator_Sids_Sid.
    Sid []*Srv6_Standby_Locators_Locator_Sids_Sid
}

func (sids *Srv6_Standby_Locators_Locator_Sids) GetEntityData() *types.CommonEntityData {
    sids.EntityData.YFilter = sids.YFilter
    sids.EntityData.YangName = "sids"
    sids.EntityData.BundleName = "cisco_ios_xr"
    sids.EntityData.ParentYangName = "locator"
    sids.EntityData.SegmentPath = "sids"
    sids.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/" + sids.EntityData.SegmentPath
    sids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sids.EntityData.Children = types.NewOrderedMap()
    sids.EntityData.Children.Append("sid", types.YChild{"Sid", nil})
    for i := range sids.Sid {
        sids.EntityData.Children.Append(types.GetSegmentPath(sids.Sid[i]), types.YChild{"Sid", sids.Sid[i]})
    }
    sids.EntityData.Leafs = types.NewOrderedMap()

    sids.EntityData.YListKeys = []string {}

    return &(sids.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid
// Operational data for given SRv6 SID
type Srv6_Standby_Locators_Locator_Sids_Sid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // SID. The type is string.
    Sid interface{}

    // Allocation Type. The type is SidAllocation.
    AllocationType interface{}

    // Function Type. The type is Srv6EndFunction.
    FunctionType interface{}

    // State. The type is SidState.
    State interface{}

    // Rewrite done or not. The type is bool.
    HasForwarding interface{}

    // Associated locator. The type is string.
    Locator interface{}

    // SID Context.
    SidContext Srv6_Standby_Locators_Locator_Sids_Sid_SidContext

    // Creation timestamp.
    CreateTimestamp Srv6_Standby_Locators_Locator_Sids_Sid_CreateTimestamp

    // Owner. The type is slice of Srv6_Standby_Locators_Locator_Sids_Sid_Owner.
    Owner []*Srv6_Standby_Locators_Locator_Sids_Sid_Owner
}

func (sid *Srv6_Standby_Locators_Locator_Sids_Sid) GetEntityData() *types.CommonEntityData {
    sid.EntityData.YFilter = sid.YFilter
    sid.EntityData.YangName = "sid"
    sid.EntityData.BundleName = "cisco_ios_xr"
    sid.EntityData.ParentYangName = "sids"
    sid.EntityData.SegmentPath = "sid" + types.AddKeyToken(sid.Address, "address")
    sid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/" + sid.EntityData.SegmentPath
    sid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sid.EntityData.Children = types.NewOrderedMap()
    sid.EntityData.Children.Append("sid-context", types.YChild{"SidContext", &sid.SidContext})
    sid.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &sid.CreateTimestamp})
    sid.EntityData.Children.Append("owner", types.YChild{"Owner", nil})
    for i := range sid.Owner {
        types.SetYListKey(sid.Owner[i], i)
        sid.EntityData.Children.Append(types.GetSegmentPath(sid.Owner[i]), types.YChild{"Owner", sid.Owner[i]})
    }
    sid.EntityData.Leafs = types.NewOrderedMap()
    sid.EntityData.Leafs.Append("address", types.YLeaf{"Address", sid.Address})
    sid.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", sid.Sid})
    sid.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", sid.AllocationType})
    sid.EntityData.Leafs.Append("function-type", types.YLeaf{"FunctionType", sid.FunctionType})
    sid.EntityData.Leafs.Append("state", types.YLeaf{"State", sid.State})
    sid.EntityData.Leafs.Append("has-forwarding", types.YLeaf{"HasForwarding", sid.HasForwarding})
    sid.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", sid.Locator})

    sid.EntityData.YListKeys = []string {"Address"}

    return &(sid.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid_SidContext
// SID Context
type Srv6_Standby_Locators_Locator_Sids_Sid_SidContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application opaque data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ApplicationData interface{}

    // SID Key.
    Key Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key
}

func (sidContext *Srv6_Standby_Locators_Locator_Sids_Sid_SidContext) GetEntityData() *types.CommonEntityData {
    sidContext.EntityData.YFilter = sidContext.YFilter
    sidContext.EntityData.YangName = "sid-context"
    sidContext.EntityData.BundleName = "cisco_ios_xr"
    sidContext.EntityData.ParentYangName = "sid"
    sidContext.EntityData.SegmentPath = "sid-context"
    sidContext.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/sid/" + sidContext.EntityData.SegmentPath
    sidContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidContext.EntityData.Children = types.NewOrderedMap()
    sidContext.EntityData.Children.Append("key", types.YChild{"Key", &sidContext.Key})
    sidContext.EntityData.Leafs = types.NewOrderedMap()
    sidContext.EntityData.Leafs.Append("application-data", types.YLeaf{"ApplicationData", sidContext.ApplicationData})

    sidContext.EntityData.YListKeys = []string {}

    return &(sidContext.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key
// SID Key
type Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDContextType. The type is Srv6EndFunction.
    SidContextType interface{}

    // End (PSP) SID context.
    E Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_E

    // End.X (PSP) SID context.
    X Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_X

    // End.DX4 SID context.
    Dx4 Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_Dx4

    // End.DT4 SID context.
    Dt4 Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_Dt4
}

func (key *Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "sid-context"
    key.EntityData.SegmentPath = "key"
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/sid/sid-context/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("e", types.YChild{"E", &key.E})
    key.EntityData.Children.Append("x", types.YChild{"X", &key.X})
    key.EntityData.Children.Append("dx4", types.YChild{"Dx4", &key.Dx4})
    key.EntityData.Children.Append("dt4", types.YChild{"Dt4", &key.Dt4})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("sid-context-type", types.YLeaf{"SidContextType", key.SidContextType})

    key.EntityData.YListKeys = []string {}

    return &(key.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_E
// End (PSP) SID context
type Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_E struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Id. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}
}

func (e *Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_E) GetEntityData() *types.CommonEntityData {
    e.EntityData.YFilter = e.YFilter
    e.EntityData.YangName = "e"
    e.EntityData.BundleName = "cisco_ios_xr"
    e.EntityData.ParentYangName = "key"
    e.EntityData.SegmentPath = "e"
    e.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/sid/sid-context/key/" + e.EntityData.SegmentPath
    e.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    e.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    e.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    e.EntityData.Children = types.NewOrderedMap()
    e.EntityData.Leafs = types.NewOrderedMap()
    e.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", e.TableId})
    e.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", e.OpaqueId})

    e.EntityData.YListKeys = []string {}

    return &(e.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_X
// End.X (PSP) SID context
type Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_X struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is protected?. The type is bool.
    IsProtected interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}

    // Nexthop interface. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Interface interface{}

    // Nexthop IP address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}
}

func (x *Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_X) GetEntityData() *types.CommonEntityData {
    x.EntityData.YFilter = x.YFilter
    x.EntityData.YangName = "x"
    x.EntityData.BundleName = "cisco_ios_xr"
    x.EntityData.ParentYangName = "key"
    x.EntityData.SegmentPath = "x"
    x.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/sid/sid-context/key/" + x.EntityData.SegmentPath
    x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    x.EntityData.Children = types.NewOrderedMap()
    x.EntityData.Leafs = types.NewOrderedMap()
    x.EntityData.Leafs.Append("is-protected", types.YLeaf{"IsProtected", x.IsProtected})
    x.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", x.OpaqueId})
    x.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", x.Interface})
    x.EntityData.Leafs.Append("nexthop-address", types.YLeaf{"NexthopAddress", x.NexthopAddress})

    x.EntityData.YListKeys = []string {}

    return &(x.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_Dx4
// End.DX4 SID context
type Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_Dx4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Next Hop Set ID. The type is interface{} with range: 0..4294967295.
    NextHopSetId interface{}
}

func (dx4 *Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_Dx4) GetEntityData() *types.CommonEntityData {
    dx4.EntityData.YFilter = dx4.YFilter
    dx4.EntityData.YangName = "dx4"
    dx4.EntityData.BundleName = "cisco_ios_xr"
    dx4.EntityData.ParentYangName = "key"
    dx4.EntityData.SegmentPath = "dx4"
    dx4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/sid/sid-context/key/" + dx4.EntityData.SegmentPath
    dx4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dx4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dx4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dx4.EntityData.Children = types.NewOrderedMap()
    dx4.EntityData.Leafs = types.NewOrderedMap()
    dx4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dx4.TableId})
    dx4.EntityData.Leafs.Append("next-hop-set-id", types.YLeaf{"NextHopSetId", dx4.NextHopSetId})

    dx4.EntityData.YListKeys = []string {}

    return &(dx4.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_Dt4
// End.DT4 SID context
type Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_Dt4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}
}

func (dt4 *Srv6_Standby_Locators_Locator_Sids_Sid_SidContext_Key_Dt4) GetEntityData() *types.CommonEntityData {
    dt4.EntityData.YFilter = dt4.YFilter
    dt4.EntityData.YangName = "dt4"
    dt4.EntityData.BundleName = "cisco_ios_xr"
    dt4.EntityData.ParentYangName = "key"
    dt4.EntityData.SegmentPath = "dt4"
    dt4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/sid/sid-context/key/" + dt4.EntityData.SegmentPath
    dt4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dt4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dt4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dt4.EntityData.Children = types.NewOrderedMap()
    dt4.EntityData.Leafs = types.NewOrderedMap()
    dt4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dt4.TableId})

    dt4.EntityData.YListKeys = []string {}

    return &(dt4.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid_CreateTimestamp
// Creation timestamp
type Srv6_Standby_Locators_Locator_Sids_Sid_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Standby_Locators_Locator_Sids_Sid_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "sid"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/sid/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Standby_Locators_Locator_Sids_Sid_Owner
// Owner
type Srv6_Standby_Locators_Locator_Sids_Sid_Owner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Owner. The type is string.
    Owner interface{}
}

func (owner *Srv6_Standby_Locators_Locator_Sids_Sid_Owner) GetEntityData() *types.CommonEntityData {
    owner.EntityData.YFilter = owner.YFilter
    owner.EntityData.YangName = "owner"
    owner.EntityData.BundleName = "cisco_ios_xr"
    owner.EntityData.ParentYangName = "sid"
    owner.EntityData.SegmentPath = "owner" + types.AddNoKeyToken(owner)
    owner.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locators/locator/sids/sid/" + owner.EntityData.SegmentPath
    owner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    owner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    owner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    owner.EntityData.Children = types.NewOrderedMap()
    owner.EntityData.Leafs = types.NewOrderedMap()
    owner.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", owner.Owner})

    owner.EntityData.YListKeys = []string {}

    return &(owner.EntityData)
}

// Srv6_Standby_LocatorAllSids
// Operational container for all (Active and Stale)
// SIDs across all Locators
type Srv6_Standby_LocatorAllSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given locator and SID opcode. The type is slice of
    // Srv6_Standby_LocatorAllSids_LocatorAllSid.
    LocatorAllSid []*Srv6_Standby_LocatorAllSids_LocatorAllSid
}

func (locatorAllSids *Srv6_Standby_LocatorAllSids) GetEntityData() *types.CommonEntityData {
    locatorAllSids.EntityData.YFilter = locatorAllSids.YFilter
    locatorAllSids.EntityData.YangName = "locator-all-sids"
    locatorAllSids.EntityData.BundleName = "cisco_ios_xr"
    locatorAllSids.EntityData.ParentYangName = "standby"
    locatorAllSids.EntityData.SegmentPath = "locator-all-sids"
    locatorAllSids.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/" + locatorAllSids.EntityData.SegmentPath
    locatorAllSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllSids.EntityData.Children = types.NewOrderedMap()
    locatorAllSids.EntityData.Children.Append("locator-all-sid", types.YChild{"LocatorAllSid", nil})
    for i := range locatorAllSids.LocatorAllSid {
        locatorAllSids.EntityData.Children.Append(types.GetSegmentPath(locatorAllSids.LocatorAllSid[i]), types.YChild{"LocatorAllSid", locatorAllSids.LocatorAllSid[i]})
    }
    locatorAllSids.EntityData.Leafs = types.NewOrderedMap()

    locatorAllSids.EntityData.YListKeys = []string {}

    return &(locatorAllSids.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid
// Operational data for given locator and SID
// opcode
type Srv6_Standby_LocatorAllSids_LocatorAllSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Locator name. The type is string with length:
    // 1..58.
    LocatorName interface{}

    // This attribute is a key. Sid opcode. The type is interface{} with range:
    // 0..4294967295.
    SidOpcode interface{}

    // SID. The type is string.
    Sid interface{}

    // Allocation Type. The type is SidAllocation.
    AllocationType interface{}

    // Function Type. The type is Srv6EndFunction.
    FunctionType interface{}

    // State. The type is SidState.
    State interface{}

    // Rewrite done or not. The type is bool.
    HasForwarding interface{}

    // Associated locator. The type is string.
    Locator interface{}

    // SID Context.
    SidContext Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext

    // Creation timestamp.
    CreateTimestamp Srv6_Standby_LocatorAllSids_LocatorAllSid_CreateTimestamp

    // Owner. The type is slice of
    // Srv6_Standby_LocatorAllSids_LocatorAllSid_Owner.
    Owner []*Srv6_Standby_LocatorAllSids_LocatorAllSid_Owner
}

func (locatorAllSid *Srv6_Standby_LocatorAllSids_LocatorAllSid) GetEntityData() *types.CommonEntityData {
    locatorAllSid.EntityData.YFilter = locatorAllSid.YFilter
    locatorAllSid.EntityData.YangName = "locator-all-sid"
    locatorAllSid.EntityData.BundleName = "cisco_ios_xr"
    locatorAllSid.EntityData.ParentYangName = "locator-all-sids"
    locatorAllSid.EntityData.SegmentPath = "locator-all-sid" + types.AddKeyToken(locatorAllSid.LocatorName, "locator-name") + types.AddKeyToken(locatorAllSid.SidOpcode, "sid-opcode")
    locatorAllSid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/" + locatorAllSid.EntityData.SegmentPath
    locatorAllSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllSid.EntityData.Children = types.NewOrderedMap()
    locatorAllSid.EntityData.Children.Append("sid-context", types.YChild{"SidContext", &locatorAllSid.SidContext})
    locatorAllSid.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &locatorAllSid.CreateTimestamp})
    locatorAllSid.EntityData.Children.Append("owner", types.YChild{"Owner", nil})
    for i := range locatorAllSid.Owner {
        types.SetYListKey(locatorAllSid.Owner[i], i)
        locatorAllSid.EntityData.Children.Append(types.GetSegmentPath(locatorAllSid.Owner[i]), types.YChild{"Owner", locatorAllSid.Owner[i]})
    }
    locatorAllSid.EntityData.Leafs = types.NewOrderedMap()
    locatorAllSid.EntityData.Leafs.Append("locator-name", types.YLeaf{"LocatorName", locatorAllSid.LocatorName})
    locatorAllSid.EntityData.Leafs.Append("sid-opcode", types.YLeaf{"SidOpcode", locatorAllSid.SidOpcode})
    locatorAllSid.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", locatorAllSid.Sid})
    locatorAllSid.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", locatorAllSid.AllocationType})
    locatorAllSid.EntityData.Leafs.Append("function-type", types.YLeaf{"FunctionType", locatorAllSid.FunctionType})
    locatorAllSid.EntityData.Leafs.Append("state", types.YLeaf{"State", locatorAllSid.State})
    locatorAllSid.EntityData.Leafs.Append("has-forwarding", types.YLeaf{"HasForwarding", locatorAllSid.HasForwarding})
    locatorAllSid.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", locatorAllSid.Locator})

    locatorAllSid.EntityData.YListKeys = []string {"LocatorName", "SidOpcode"}

    return &(locatorAllSid.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext
// SID Context
type Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application opaque data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ApplicationData interface{}

    // SID Key.
    Key Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key
}

func (sidContext *Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext) GetEntityData() *types.CommonEntityData {
    sidContext.EntityData.YFilter = sidContext.YFilter
    sidContext.EntityData.YangName = "sid-context"
    sidContext.EntityData.BundleName = "cisco_ios_xr"
    sidContext.EntityData.ParentYangName = "locator-all-sid"
    sidContext.EntityData.SegmentPath = "sid-context"
    sidContext.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/locator-all-sid/" + sidContext.EntityData.SegmentPath
    sidContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidContext.EntityData.Children = types.NewOrderedMap()
    sidContext.EntityData.Children.Append("key", types.YChild{"Key", &sidContext.Key})
    sidContext.EntityData.Leafs = types.NewOrderedMap()
    sidContext.EntityData.Leafs.Append("application-data", types.YLeaf{"ApplicationData", sidContext.ApplicationData})

    sidContext.EntityData.YListKeys = []string {}

    return &(sidContext.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key
// SID Key
type Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDContextType. The type is Srv6EndFunction.
    SidContextType interface{}

    // End (PSP) SID context.
    E Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_E

    // End.X (PSP) SID context.
    X Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_X

    // End.DX4 SID context.
    Dx4 Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_Dx4

    // End.DT4 SID context.
    Dt4 Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_Dt4
}

func (key *Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "sid-context"
    key.EntityData.SegmentPath = "key"
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/locator-all-sid/sid-context/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("e", types.YChild{"E", &key.E})
    key.EntityData.Children.Append("x", types.YChild{"X", &key.X})
    key.EntityData.Children.Append("dx4", types.YChild{"Dx4", &key.Dx4})
    key.EntityData.Children.Append("dt4", types.YChild{"Dt4", &key.Dt4})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("sid-context-type", types.YLeaf{"SidContextType", key.SidContextType})

    key.EntityData.YListKeys = []string {}

    return &(key.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_E
// End (PSP) SID context
type Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_E struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Id. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}
}

func (e *Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_E) GetEntityData() *types.CommonEntityData {
    e.EntityData.YFilter = e.YFilter
    e.EntityData.YangName = "e"
    e.EntityData.BundleName = "cisco_ios_xr"
    e.EntityData.ParentYangName = "key"
    e.EntityData.SegmentPath = "e"
    e.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/locator-all-sid/sid-context/key/" + e.EntityData.SegmentPath
    e.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    e.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    e.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    e.EntityData.Children = types.NewOrderedMap()
    e.EntityData.Leafs = types.NewOrderedMap()
    e.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", e.TableId})
    e.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", e.OpaqueId})

    e.EntityData.YListKeys = []string {}

    return &(e.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_X
// End.X (PSP) SID context
type Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_X struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is protected?. The type is bool.
    IsProtected interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}

    // Nexthop interface. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Interface interface{}

    // Nexthop IP address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}
}

func (x *Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_X) GetEntityData() *types.CommonEntityData {
    x.EntityData.YFilter = x.YFilter
    x.EntityData.YangName = "x"
    x.EntityData.BundleName = "cisco_ios_xr"
    x.EntityData.ParentYangName = "key"
    x.EntityData.SegmentPath = "x"
    x.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/locator-all-sid/sid-context/key/" + x.EntityData.SegmentPath
    x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    x.EntityData.Children = types.NewOrderedMap()
    x.EntityData.Leafs = types.NewOrderedMap()
    x.EntityData.Leafs.Append("is-protected", types.YLeaf{"IsProtected", x.IsProtected})
    x.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", x.OpaqueId})
    x.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", x.Interface})
    x.EntityData.Leafs.Append("nexthop-address", types.YLeaf{"NexthopAddress", x.NexthopAddress})

    x.EntityData.YListKeys = []string {}

    return &(x.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_Dx4
// End.DX4 SID context
type Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_Dx4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Next Hop Set ID. The type is interface{} with range: 0..4294967295.
    NextHopSetId interface{}
}

func (dx4 *Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_Dx4) GetEntityData() *types.CommonEntityData {
    dx4.EntityData.YFilter = dx4.YFilter
    dx4.EntityData.YangName = "dx4"
    dx4.EntityData.BundleName = "cisco_ios_xr"
    dx4.EntityData.ParentYangName = "key"
    dx4.EntityData.SegmentPath = "dx4"
    dx4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/locator-all-sid/sid-context/key/" + dx4.EntityData.SegmentPath
    dx4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dx4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dx4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dx4.EntityData.Children = types.NewOrderedMap()
    dx4.EntityData.Leafs = types.NewOrderedMap()
    dx4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dx4.TableId})
    dx4.EntityData.Leafs.Append("next-hop-set-id", types.YLeaf{"NextHopSetId", dx4.NextHopSetId})

    dx4.EntityData.YListKeys = []string {}

    return &(dx4.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_Dt4
// End.DT4 SID context
type Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_Dt4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}
}

func (dt4 *Srv6_Standby_LocatorAllSids_LocatorAllSid_SidContext_Key_Dt4) GetEntityData() *types.CommonEntityData {
    dt4.EntityData.YFilter = dt4.YFilter
    dt4.EntityData.YangName = "dt4"
    dt4.EntityData.BundleName = "cisco_ios_xr"
    dt4.EntityData.ParentYangName = "key"
    dt4.EntityData.SegmentPath = "dt4"
    dt4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/locator-all-sid/sid-context/key/" + dt4.EntityData.SegmentPath
    dt4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dt4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dt4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dt4.EntityData.Children = types.NewOrderedMap()
    dt4.EntityData.Leafs = types.NewOrderedMap()
    dt4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dt4.TableId})

    dt4.EntityData.YListKeys = []string {}

    return &(dt4.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid_CreateTimestamp
// Creation timestamp
type Srv6_Standby_LocatorAllSids_LocatorAllSid_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Standby_LocatorAllSids_LocatorAllSid_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "locator-all-sid"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/locator-all-sid/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Standby_LocatorAllSids_LocatorAllSid_Owner
// Owner
type Srv6_Standby_LocatorAllSids_LocatorAllSid_Owner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Owner. The type is string.
    Owner interface{}
}

func (owner *Srv6_Standby_LocatorAllSids_LocatorAllSid_Owner) GetEntityData() *types.CommonEntityData {
    owner.EntityData.YFilter = owner.YFilter
    owner.EntityData.YangName = "owner"
    owner.EntityData.BundleName = "cisco_ios_xr"
    owner.EntityData.ParentYangName = "locator-all-sid"
    owner.EntityData.SegmentPath = "owner" + types.AddNoKeyToken(owner)
    owner.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-sids/locator-all-sid/" + owner.EntityData.SegmentPath
    owner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    owner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    owner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    owner.EntityData.Children = types.NewOrderedMap()
    owner.EntityData.Leafs = types.NewOrderedMap()
    owner.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", owner.Owner})

    owner.EntityData.YListKeys = []string {}

    return &(owner.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids
// Operational container for Active SIDs across all
// Locators
type Srv6_Standby_LocatorAllActiveSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for given locator and SID opcode. The type is slice of
    // Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid.
    LocatorAllActiveSid []*Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid
}

func (locatorAllActiveSids *Srv6_Standby_LocatorAllActiveSids) GetEntityData() *types.CommonEntityData {
    locatorAllActiveSids.EntityData.YFilter = locatorAllActiveSids.YFilter
    locatorAllActiveSids.EntityData.YangName = "locator-all-active-sids"
    locatorAllActiveSids.EntityData.BundleName = "cisco_ios_xr"
    locatorAllActiveSids.EntityData.ParentYangName = "standby"
    locatorAllActiveSids.EntityData.SegmentPath = "locator-all-active-sids"
    locatorAllActiveSids.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/" + locatorAllActiveSids.EntityData.SegmentPath
    locatorAllActiveSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllActiveSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllActiveSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllActiveSids.EntityData.Children = types.NewOrderedMap()
    locatorAllActiveSids.EntityData.Children.Append("locator-all-active-sid", types.YChild{"LocatorAllActiveSid", nil})
    for i := range locatorAllActiveSids.LocatorAllActiveSid {
        locatorAllActiveSids.EntityData.Children.Append(types.GetSegmentPath(locatorAllActiveSids.LocatorAllActiveSid[i]), types.YChild{"LocatorAllActiveSid", locatorAllActiveSids.LocatorAllActiveSid[i]})
    }
    locatorAllActiveSids.EntityData.Leafs = types.NewOrderedMap()

    locatorAllActiveSids.EntityData.YListKeys = []string {}

    return &(locatorAllActiveSids.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid
// Operational data for given locator and SID
// opcode
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Locator name. The type is string with length:
    // 1..58.
    LocatorName interface{}

    // This attribute is a key. Sid opcode. The type is interface{} with range:
    // 0..4294967295.
    SidOpcode interface{}

    // SID. The type is string.
    Sid interface{}

    // Allocation Type. The type is SidAllocation.
    AllocationType interface{}

    // Function Type. The type is Srv6EndFunction.
    FunctionType interface{}

    // State. The type is SidState.
    State interface{}

    // Rewrite done or not. The type is bool.
    HasForwarding interface{}

    // Associated locator. The type is string.
    Locator interface{}

    // SID Context.
    SidContext Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext

    // Creation timestamp.
    CreateTimestamp Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_CreateTimestamp

    // Owner. The type is slice of
    // Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_Owner.
    Owner []*Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_Owner
}

func (locatorAllActiveSid *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid) GetEntityData() *types.CommonEntityData {
    locatorAllActiveSid.EntityData.YFilter = locatorAllActiveSid.YFilter
    locatorAllActiveSid.EntityData.YangName = "locator-all-active-sid"
    locatorAllActiveSid.EntityData.BundleName = "cisco_ios_xr"
    locatorAllActiveSid.EntityData.ParentYangName = "locator-all-active-sids"
    locatorAllActiveSid.EntityData.SegmentPath = "locator-all-active-sid" + types.AddKeyToken(locatorAllActiveSid.LocatorName, "locator-name") + types.AddKeyToken(locatorAllActiveSid.SidOpcode, "sid-opcode")
    locatorAllActiveSid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/" + locatorAllActiveSid.EntityData.SegmentPath
    locatorAllActiveSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locatorAllActiveSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locatorAllActiveSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locatorAllActiveSid.EntityData.Children = types.NewOrderedMap()
    locatorAllActiveSid.EntityData.Children.Append("sid-context", types.YChild{"SidContext", &locatorAllActiveSid.SidContext})
    locatorAllActiveSid.EntityData.Children.Append("create-timestamp", types.YChild{"CreateTimestamp", &locatorAllActiveSid.CreateTimestamp})
    locatorAllActiveSid.EntityData.Children.Append("owner", types.YChild{"Owner", nil})
    for i := range locatorAllActiveSid.Owner {
        types.SetYListKey(locatorAllActiveSid.Owner[i], i)
        locatorAllActiveSid.EntityData.Children.Append(types.GetSegmentPath(locatorAllActiveSid.Owner[i]), types.YChild{"Owner", locatorAllActiveSid.Owner[i]})
    }
    locatorAllActiveSid.EntityData.Leafs = types.NewOrderedMap()
    locatorAllActiveSid.EntityData.Leafs.Append("locator-name", types.YLeaf{"LocatorName", locatorAllActiveSid.LocatorName})
    locatorAllActiveSid.EntityData.Leafs.Append("sid-opcode", types.YLeaf{"SidOpcode", locatorAllActiveSid.SidOpcode})
    locatorAllActiveSid.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", locatorAllActiveSid.Sid})
    locatorAllActiveSid.EntityData.Leafs.Append("allocation-type", types.YLeaf{"AllocationType", locatorAllActiveSid.AllocationType})
    locatorAllActiveSid.EntityData.Leafs.Append("function-type", types.YLeaf{"FunctionType", locatorAllActiveSid.FunctionType})
    locatorAllActiveSid.EntityData.Leafs.Append("state", types.YLeaf{"State", locatorAllActiveSid.State})
    locatorAllActiveSid.EntityData.Leafs.Append("has-forwarding", types.YLeaf{"HasForwarding", locatorAllActiveSid.HasForwarding})
    locatorAllActiveSid.EntityData.Leafs.Append("locator", types.YLeaf{"Locator", locatorAllActiveSid.Locator})

    locatorAllActiveSid.EntityData.YListKeys = []string {"LocatorName", "SidOpcode"}

    return &(locatorAllActiveSid.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext
// SID Context
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application opaque data. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    ApplicationData interface{}

    // SID Key.
    Key Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key
}

func (sidContext *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext) GetEntityData() *types.CommonEntityData {
    sidContext.EntityData.YFilter = sidContext.YFilter
    sidContext.EntityData.YangName = "sid-context"
    sidContext.EntityData.BundleName = "cisco_ios_xr"
    sidContext.EntityData.ParentYangName = "locator-all-active-sid"
    sidContext.EntityData.SegmentPath = "sid-context"
    sidContext.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/locator-all-active-sid/" + sidContext.EntityData.SegmentPath
    sidContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidContext.EntityData.Children = types.NewOrderedMap()
    sidContext.EntityData.Children.Append("key", types.YChild{"Key", &sidContext.Key})
    sidContext.EntityData.Leafs = types.NewOrderedMap()
    sidContext.EntityData.Leafs.Append("application-data", types.YLeaf{"ApplicationData", sidContext.ApplicationData})

    sidContext.EntityData.YListKeys = []string {}

    return &(sidContext.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key
// SID Key
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDContextType. The type is Srv6EndFunction.
    SidContextType interface{}

    // End (PSP) SID context.
    E Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_E

    // End.X (PSP) SID context.
    X Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_X

    // End.DX4 SID context.
    Dx4 Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dx4

    // End.DT4 SID context.
    Dt4 Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dt4
}

func (key *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "cisco_ios_xr"
    key.EntityData.ParentYangName = "sid-context"
    key.EntityData.SegmentPath = "key"
    key.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/locator-all-active-sid/sid-context/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    key.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("e", types.YChild{"E", &key.E})
    key.EntityData.Children.Append("x", types.YChild{"X", &key.X})
    key.EntityData.Children.Append("dx4", types.YChild{"Dx4", &key.Dx4})
    key.EntityData.Children.Append("dt4", types.YChild{"Dt4", &key.Dt4})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("sid-context-type", types.YLeaf{"SidContextType", key.SidContextType})

    key.EntityData.YListKeys = []string {}

    return &(key.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_E
// End (PSP) SID context
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_E struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table Id. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}
}

func (e *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_E) GetEntityData() *types.CommonEntityData {
    e.EntityData.YFilter = e.YFilter
    e.EntityData.YangName = "e"
    e.EntityData.BundleName = "cisco_ios_xr"
    e.EntityData.ParentYangName = "key"
    e.EntityData.SegmentPath = "e"
    e.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/locator-all-active-sid/sid-context/key/" + e.EntityData.SegmentPath
    e.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    e.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    e.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    e.EntityData.Children = types.NewOrderedMap()
    e.EntityData.Leafs = types.NewOrderedMap()
    e.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", e.TableId})
    e.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", e.OpaqueId})

    e.EntityData.YListKeys = []string {}

    return &(e.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_X
// End.X (PSP) SID context
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_X struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is protected?. The type is bool.
    IsProtected interface{}

    // Additional differentiator - opaque to SIDMgr. The type is interface{} with
    // range: 0..255.
    OpaqueId interface{}

    // Nexthop interface. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    Interface interface{}

    // Nexthop IP address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NexthopAddress interface{}
}

func (x *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_X) GetEntityData() *types.CommonEntityData {
    x.EntityData.YFilter = x.YFilter
    x.EntityData.YangName = "x"
    x.EntityData.BundleName = "cisco_ios_xr"
    x.EntityData.ParentYangName = "key"
    x.EntityData.SegmentPath = "x"
    x.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/locator-all-active-sid/sid-context/key/" + x.EntityData.SegmentPath
    x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    x.EntityData.Children = types.NewOrderedMap()
    x.EntityData.Leafs = types.NewOrderedMap()
    x.EntityData.Leafs.Append("is-protected", types.YLeaf{"IsProtected", x.IsProtected})
    x.EntityData.Leafs.Append("opaque-id", types.YLeaf{"OpaqueId", x.OpaqueId})
    x.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", x.Interface})
    x.EntityData.Leafs.Append("nexthop-address", types.YLeaf{"NexthopAddress", x.NexthopAddress})

    x.EntityData.YListKeys = []string {}

    return &(x.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dx4
// End.DX4 SID context
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dx4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Next Hop Set ID. The type is interface{} with range: 0..4294967295.
    NextHopSetId interface{}
}

func (dx4 *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dx4) GetEntityData() *types.CommonEntityData {
    dx4.EntityData.YFilter = dx4.YFilter
    dx4.EntityData.YangName = "dx4"
    dx4.EntityData.BundleName = "cisco_ios_xr"
    dx4.EntityData.ParentYangName = "key"
    dx4.EntityData.SegmentPath = "dx4"
    dx4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/locator-all-active-sid/sid-context/key/" + dx4.EntityData.SegmentPath
    dx4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dx4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dx4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dx4.EntityData.Children = types.NewOrderedMap()
    dx4.EntityData.Leafs = types.NewOrderedMap()
    dx4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dx4.TableId})
    dx4.EntityData.Leafs.Append("next-hop-set-id", types.YLeaf{"NextHopSetId", dx4.NextHopSetId})

    dx4.EntityData.YListKeys = []string {}

    return &(dx4.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dt4
// End.DT4 SID context
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dt4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}
}

func (dt4 *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_SidContext_Key_Dt4) GetEntityData() *types.CommonEntityData {
    dt4.EntityData.YFilter = dt4.YFilter
    dt4.EntityData.YangName = "dt4"
    dt4.EntityData.BundleName = "cisco_ios_xr"
    dt4.EntityData.ParentYangName = "key"
    dt4.EntityData.SegmentPath = "dt4"
    dt4.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/locator-all-active-sid/sid-context/key/" + dt4.EntityData.SegmentPath
    dt4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dt4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dt4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dt4.EntityData.Children = types.NewOrderedMap()
    dt4.EntityData.Leafs = types.NewOrderedMap()
    dt4.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", dt4.TableId})

    dt4.EntityData.YListKeys = []string {}

    return &(dt4.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_CreateTimestamp
// Creation timestamp
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_CreateTimestamp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TimeInNanoSeconds interface{}

    // Age in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    AgeInNanoSeconds interface{}
}

func (createTimestamp *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_CreateTimestamp) GetEntityData() *types.CommonEntityData {
    createTimestamp.EntityData.YFilter = createTimestamp.YFilter
    createTimestamp.EntityData.YangName = "create-timestamp"
    createTimestamp.EntityData.BundleName = "cisco_ios_xr"
    createTimestamp.EntityData.ParentYangName = "locator-all-active-sid"
    createTimestamp.EntityData.SegmentPath = "create-timestamp"
    createTimestamp.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/locator-all-active-sid/" + createTimestamp.EntityData.SegmentPath
    createTimestamp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    createTimestamp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    createTimestamp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    createTimestamp.EntityData.Children = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs = types.NewOrderedMap()
    createTimestamp.EntityData.Leafs.Append("time-in-nano-seconds", types.YLeaf{"TimeInNanoSeconds", createTimestamp.TimeInNanoSeconds})
    createTimestamp.EntityData.Leafs.Append("age-in-nano-seconds", types.YLeaf{"AgeInNanoSeconds", createTimestamp.AgeInNanoSeconds})

    createTimestamp.EntityData.YListKeys = []string {}

    return &(createTimestamp.EntityData)
}

// Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_Owner
// Owner
type Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_Owner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Owner. The type is string.
    Owner interface{}
}

func (owner *Srv6_Standby_LocatorAllActiveSids_LocatorAllActiveSid_Owner) GetEntityData() *types.CommonEntityData {
    owner.EntityData.YFilter = owner.YFilter
    owner.EntityData.YangName = "owner"
    owner.EntityData.BundleName = "cisco_ios_xr"
    owner.EntityData.ParentYangName = "locator-all-active-sid"
    owner.EntityData.SegmentPath = "owner" + types.AddNoKeyToken(owner)
    owner.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-srv6-oper:srv6/standby/locator-all-active-sids/locator-all-active-sid/" + owner.EntityData.SegmentPath
    owner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    owner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    owner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    owner.EntityData.Children = types.NewOrderedMap()
    owner.EntityData.Leafs = types.NewOrderedMap()
    owner.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", owner.Owner})

    owner.EntityData.YListKeys = []string {}

    return &(owner.EntityData)
}

