// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-xtc package configuration.
// 
// This module contains definitions
// for the following management objects:
//   pce: PCE configuration data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_xtc_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_xtc_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-xtc-cfg pce}", reflect.TypeOf(Pce{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-xtc-cfg:pce", reflect.TypeOf(Pce{}))
}

// PceSegment represents Pce segment
type PceSegment string

const (
    // IPv4 Address
    PceSegment_ipv4_address PceSegment = "ipv4-address"

    // MPLS Label
    PceSegment_mpls_label PceSegment = "mpls-label"
)

// PceBindingSid represents Pce binding sid
type PceBindingSid string

const (
    // Use specified BSID MPLS label
    PceBindingSid_mpls_label_specified PceBindingSid = "mpls-label-specified"

    // Allocate BSID MPLS label
    PceBindingSid_mpls_label_any PceBindingSid = "mpls-label-any"
)

// PceExplicitPathHop represents Pce explicit path hop
type PceExplicitPathHop string

const (
    // Address
    PceExplicitPathHop_address PceExplicitPathHop = "address"

    // SID Node
    PceExplicitPathHop_sid_node PceExplicitPathHop = "sid-node"

    // SID Adjacency
    PceExplicitPathHop_sid_adjancency PceExplicitPathHop = "sid-adjancency"

    // Binding SID
    PceExplicitPathHop_binding_sid PceExplicitPathHop = "binding-sid"
)

// PcePath represents Pce path
type PcePath string

const (
    // Explicit
    PcePath_explicit PcePath = "explicit"

    // Dynamic
    PcePath_dynamic PcePath = "dynamic"
)

// PceEndPoint represents Pce end point
type PceEndPoint string

const (
    // IPv4 endpoint address
    PceEndPoint_end_point_type_ipv4 PceEndPoint = "end-point-type-ipv4"

    // IPv6 endpoint address
    PceEndPoint_end_point_type_ipv6 PceEndPoint = "end-point-type-ipv6"
)

// PcerestAuthentication represents Pcerest authentication
type PcerestAuthentication string

const (
    // Basic HTTP auth
    PcerestAuthentication_basic PcerestAuthentication = "basic"

    // MD5-Digest HTTP auth
    PcerestAuthentication_digest PcerestAuthentication = "digest"
)

// PcePathHop represents Pce path hop
type PcePathHop string

const (
    // Segment-routing MPLS
    PcePathHop_mpls PcePathHop = "mpls"

    // Segment-routing v6
    PcePathHop_srv6 PcePathHop = "srv6"
)

// PceMetric represents Pce metric
type PceMetric string

const (
    // IGP metric type
    PceMetric_igp PceMetric = "igp"

    // TE metric type
    PceMetric_te PceMetric = "te"
)

// PceDisjointPath represents Pce disjoint path
type PceDisjointPath string

const (
    // Link
    PceDisjointPath_link PceDisjointPath = "link"

    // Node
    PceDisjointPath_node PceDisjointPath = "node"

    // SRLG
    PceDisjointPath_srlg PceDisjointPath = "srlg"

    // SRLG Node
    PceDisjointPath_srlg_node PceDisjointPath = "srlg-node"
)

// Pce
// PCE configuration data
type Pce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address of PCE server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // IPv6 address of PCE server. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6ServerAddress interface{}

    // MD5 password. The type is string with pattern: (!.+)|([^!].+).
    Password interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Standby IPv6 PCE configuration.
    Ipv6StateSyncs Pce_Ipv6StateSyncs

    // Path computation client configuration.
    PccAddresses Pce_PccAddresses

    // PCE logging configuration.
    Logging Pce_Logging

    // PCE backoff configuration.
    Backoff Pce_Backoff

    // REST configuration.
    Rest Pce_Rest

    // Standby IPv4 PCE configuration.
    StateSyncs Pce_StateSyncs

    // PCE segment-routing configuration.
    SegmentRouting Pce_SegmentRouting

    // PCE Timers configuration.
    Timers Pce_Timers

    // NETCONF configuration.
    Netconf Pce_Netconf

    // Disjoint path configuration.
    DisjointPath Pce_DisjointPath

    // Explicit paths.
    ExplicitPaths Pce_ExplicitPaths
}

func (pce *Pce) GetEntityData() *types.CommonEntityData {
    pce.EntityData.YFilter = pce.YFilter
    pce.EntityData.YangName = "pce"
    pce.EntityData.BundleName = "cisco_ios_xr"
    pce.EntityData.ParentYangName = "Cisco-IOS-XR-infra-xtc-cfg"
    pce.EntityData.SegmentPath = "Cisco-IOS-XR-infra-xtc-cfg:pce"
    pce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pce.EntityData.Children = types.NewOrderedMap()
    pce.EntityData.Children.Append("ipv6-state-syncs", types.YChild{"Ipv6StateSyncs", &pce.Ipv6StateSyncs})
    pce.EntityData.Children.Append("pcc-addresses", types.YChild{"PccAddresses", &pce.PccAddresses})
    pce.EntityData.Children.Append("logging", types.YChild{"Logging", &pce.Logging})
    pce.EntityData.Children.Append("backoff", types.YChild{"Backoff", &pce.Backoff})
    pce.EntityData.Children.Append("rest", types.YChild{"Rest", &pce.Rest})
    pce.EntityData.Children.Append("state-syncs", types.YChild{"StateSyncs", &pce.StateSyncs})
    pce.EntityData.Children.Append("segment-routing", types.YChild{"SegmentRouting", &pce.SegmentRouting})
    pce.EntityData.Children.Append("timers", types.YChild{"Timers", &pce.Timers})
    pce.EntityData.Children.Append("netconf", types.YChild{"Netconf", &pce.Netconf})
    pce.EntityData.Children.Append("disjoint-path", types.YChild{"DisjointPath", &pce.DisjointPath})
    pce.EntityData.Children.Append("explicit-paths", types.YChild{"ExplicitPaths", &pce.ExplicitPaths})
    pce.EntityData.Leafs = types.NewOrderedMap()
    pce.EntityData.Leafs.Append("server-address", types.YLeaf{"ServerAddress", pce.ServerAddress})
    pce.EntityData.Leafs.Append("ipv6-server-address", types.YLeaf{"Ipv6ServerAddress", pce.Ipv6ServerAddress})
    pce.EntityData.Leafs.Append("password", types.YLeaf{"Password", pce.Password})
    pce.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pce.Enable})

    pce.EntityData.YListKeys = []string {}

    return &(pce.EntityData)
}

// Pce_Ipv6StateSyncs
// Standby IPv6 PCE configuration
type Pce_Ipv6StateSyncs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Standby PCE ipv6 address. The type is slice of
    // Pce_Ipv6StateSyncs_Ipv6StateSync.
    Ipv6StateSync []*Pce_Ipv6StateSyncs_Ipv6StateSync
}

func (ipv6StateSyncs *Pce_Ipv6StateSyncs) GetEntityData() *types.CommonEntityData {
    ipv6StateSyncs.EntityData.YFilter = ipv6StateSyncs.YFilter
    ipv6StateSyncs.EntityData.YangName = "ipv6-state-syncs"
    ipv6StateSyncs.EntityData.BundleName = "cisco_ios_xr"
    ipv6StateSyncs.EntityData.ParentYangName = "pce"
    ipv6StateSyncs.EntityData.SegmentPath = "ipv6-state-syncs"
    ipv6StateSyncs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6StateSyncs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6StateSyncs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6StateSyncs.EntityData.Children = types.NewOrderedMap()
    ipv6StateSyncs.EntityData.Children.Append("ipv6-state-sync", types.YChild{"Ipv6StateSync", nil})
    for i := range ipv6StateSyncs.Ipv6StateSync {
        ipv6StateSyncs.EntityData.Children.Append(types.GetSegmentPath(ipv6StateSyncs.Ipv6StateSync[i]), types.YChild{"Ipv6StateSync", ipv6StateSyncs.Ipv6StateSync[i]})
    }
    ipv6StateSyncs.EntityData.Leafs = types.NewOrderedMap()

    ipv6StateSyncs.EntityData.YListKeys = []string {}

    return &(ipv6StateSyncs.EntityData)
}

// Pce_Ipv6StateSyncs_Ipv6StateSync
// Standby PCE ipv6 address
type Pce_Ipv6StateSyncs_Ipv6StateSync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (ipv6StateSync *Pce_Ipv6StateSyncs_Ipv6StateSync) GetEntityData() *types.CommonEntityData {
    ipv6StateSync.EntityData.YFilter = ipv6StateSync.YFilter
    ipv6StateSync.EntityData.YangName = "ipv6-state-sync"
    ipv6StateSync.EntityData.BundleName = "cisco_ios_xr"
    ipv6StateSync.EntityData.ParentYangName = "ipv6-state-syncs"
    ipv6StateSync.EntityData.SegmentPath = "ipv6-state-sync" + types.AddKeyToken(ipv6StateSync.Address, "address")
    ipv6StateSync.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6StateSync.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6StateSync.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6StateSync.EntityData.Children = types.NewOrderedMap()
    ipv6StateSync.EntityData.Leafs = types.NewOrderedMap()
    ipv6StateSync.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv6StateSync.Address})

    ipv6StateSync.EntityData.YListKeys = []string {"Address"}

    return &(ipv6StateSync.EntityData)
}

// Pce_PccAddresses
// Path computation client configuration
type Pce_PccAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path computation client address. The type is slice of
    // Pce_PccAddresses_PccAddress.
    PccAddress []*Pce_PccAddresses_PccAddress
}

func (pccAddresses *Pce_PccAddresses) GetEntityData() *types.CommonEntityData {
    pccAddresses.EntityData.YFilter = pccAddresses.YFilter
    pccAddresses.EntityData.YangName = "pcc-addresses"
    pccAddresses.EntityData.BundleName = "cisco_ios_xr"
    pccAddresses.EntityData.ParentYangName = "pce"
    pccAddresses.EntityData.SegmentPath = "pcc-addresses"
    pccAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pccAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pccAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pccAddresses.EntityData.Children = types.NewOrderedMap()
    pccAddresses.EntityData.Children.Append("pcc-address", types.YChild{"PccAddress", nil})
    for i := range pccAddresses.PccAddress {
        pccAddresses.EntityData.Children.Append(types.GetSegmentPath(pccAddresses.PccAddress[i]), types.YChild{"PccAddress", pccAddresses.PccAddress[i]})
    }
    pccAddresses.EntityData.Leafs = types.NewOrderedMap()

    pccAddresses.EntityData.YListKeys = []string {}

    return &(pccAddresses.EntityData)
}

// Pce_PccAddresses_PccAddress
// Path computation client address
type Pce_PccAddresses_PccAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // MPLS label switched path.
    LspNames Pce_PccAddresses_PccAddress_LspNames
}

func (pccAddress *Pce_PccAddresses_PccAddress) GetEntityData() *types.CommonEntityData {
    pccAddress.EntityData.YFilter = pccAddress.YFilter
    pccAddress.EntityData.YangName = "pcc-address"
    pccAddress.EntityData.BundleName = "cisco_ios_xr"
    pccAddress.EntityData.ParentYangName = "pcc-addresses"
    pccAddress.EntityData.SegmentPath = "pcc-address" + types.AddKeyToken(pccAddress.Address, "address")
    pccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pccAddress.EntityData.Children = types.NewOrderedMap()
    pccAddress.EntityData.Children.Append("lsp-names", types.YChild{"LspNames", &pccAddress.LspNames})
    pccAddress.EntityData.Leafs = types.NewOrderedMap()
    pccAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", pccAddress.Address})
    pccAddress.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pccAddress.Enable})

    pccAddress.EntityData.YListKeys = []string {"Address"}

    return &(pccAddress.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames
// MPLS label switched path
type Pce_PccAddresses_PccAddress_LspNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS label switched path. The type is slice of
    // Pce_PccAddresses_PccAddress_LspNames_LspName.
    LspName []*Pce_PccAddresses_PccAddress_LspNames_LspName
}

func (lspNames *Pce_PccAddresses_PccAddress_LspNames) GetEntityData() *types.CommonEntityData {
    lspNames.EntityData.YFilter = lspNames.YFilter
    lspNames.EntityData.YangName = "lsp-names"
    lspNames.EntityData.BundleName = "cisco_ios_xr"
    lspNames.EntityData.ParentYangName = "pcc-address"
    lspNames.EntityData.SegmentPath = "lsp-names"
    lspNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspNames.EntityData.Children = types.NewOrderedMap()
    lspNames.EntityData.Children.Append("lsp-name", types.YChild{"LspName", nil})
    for i := range lspNames.LspName {
        lspNames.EntityData.Children.Append(types.GetSegmentPath(lspNames.LspName[i]), types.YChild{"LspName", lspNames.LspName[i]})
    }
    lspNames.EntityData.Leafs = types.NewOrderedMap()

    lspNames.EntityData.YListKeys = []string {}

    return &(lspNames.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames_LspName
// MPLS label switched path
type Pce_PccAddresses_PccAddress_LspNames_LspName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSP name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Undelegate LSP. The type is interface{}.
    Undelegate interface{}

    // Explicit-path name. The type is string. This attribute is mandatory.
    ExplicitPathName interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // RSVP-TE configuration.
    RsvpTe Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe
}

func (lspName *Pce_PccAddresses_PccAddress_LspNames_LspName) GetEntityData() *types.CommonEntityData {
    lspName.EntityData.YFilter = lspName.YFilter
    lspName.EntityData.YangName = "lsp-name"
    lspName.EntityData.BundleName = "cisco_ios_xr"
    lspName.EntityData.ParentYangName = "lsp-names"
    lspName.EntityData.SegmentPath = "lsp-name" + types.AddKeyToken(lspName.Name, "name")
    lspName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspName.EntityData.Children = types.NewOrderedMap()
    lspName.EntityData.Children.Append("rsvp-te", types.YChild{"RsvpTe", &lspName.RsvpTe})
    lspName.EntityData.Leafs = types.NewOrderedMap()
    lspName.EntityData.Leafs.Append("name", types.YLeaf{"Name", lspName.Name})
    lspName.EntityData.Leafs.Append("undelegate", types.YLeaf{"Undelegate", lspName.Undelegate})
    lspName.EntityData.Leafs.Append("explicit-path-name", types.YLeaf{"ExplicitPathName", lspName.ExplicitPathName})
    lspName.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", lspName.Enable})

    lspName.EntityData.YListKeys = []string {"Name"}

    return &(lspName.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe
// RSVP-TE configuration
// This type is a presence type.
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable fast protection. The type is interface{}.
    FastProtect interface{}

    // Bandwidth configuration. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory. Units are kbit/s.
    Bandwidth interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // LSP Affinity.
    Affinity Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity

    // Tunnel Setup and Hold Priorities.
    Priority Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority
}

func (rsvpTe *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe) GetEntityData() *types.CommonEntityData {
    rsvpTe.EntityData.YFilter = rsvpTe.YFilter
    rsvpTe.EntityData.YangName = "rsvp-te"
    rsvpTe.EntityData.BundleName = "cisco_ios_xr"
    rsvpTe.EntityData.ParentYangName = "lsp-name"
    rsvpTe.EntityData.SegmentPath = "rsvp-te"
    rsvpTe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvpTe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvpTe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvpTe.EntityData.Children = types.NewOrderedMap()
    rsvpTe.EntityData.Children.Append("affinity", types.YChild{"Affinity", &rsvpTe.Affinity})
    rsvpTe.EntityData.Children.Append("priority", types.YChild{"Priority", &rsvpTe.Priority})
    rsvpTe.EntityData.Leafs = types.NewOrderedMap()
    rsvpTe.EntityData.Leafs.Append("fast-protect", types.YLeaf{"FastProtect", rsvpTe.FastProtect})
    rsvpTe.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", rsvpTe.Bandwidth})
    rsvpTe.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rsvpTe.Enable})

    rsvpTe.EntityData.YListKeys = []string {}

    return &(rsvpTe.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity
// LSP Affinity
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Include-any affinity value. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    IncludeAny interface{}

    // Include-all affinity value. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    IncludeAll interface{}

    // Exclude-any affinity value. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    ExcludeAny interface{}
}

func (affinity *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity) GetEntityData() *types.CommonEntityData {
    affinity.EntityData.YFilter = affinity.YFilter
    affinity.EntityData.YangName = "affinity"
    affinity.EntityData.BundleName = "cisco_ios_xr"
    affinity.EntityData.ParentYangName = "rsvp-te"
    affinity.EntityData.SegmentPath = "affinity"
    affinity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinity.EntityData.Children = types.NewOrderedMap()
    affinity.EntityData.Leafs = types.NewOrderedMap()
    affinity.EntityData.Leafs.Append("include-any", types.YLeaf{"IncludeAny", affinity.IncludeAny})
    affinity.EntityData.Leafs.Append("include-all", types.YLeaf{"IncludeAll", affinity.IncludeAll})
    affinity.EntityData.Leafs.Append("exclude-any", types.YLeaf{"ExcludeAny", affinity.ExcludeAny})

    affinity.EntityData.YListKeys = []string {}

    return &(affinity.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority
// Tunnel Setup and Hold Priorities
// This type is a presence type.
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Setup Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    HoldPriority interface{}
}

func (priority *Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "rsvp-te"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", priority.SetupPriority})
    priority.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", priority.HoldPriority})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// Pce_Logging
// PCE logging configuration
type Pce_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Logging NO-PATH configuration. The type is interface{}.
    NoPath interface{}

    // Logging of received PCErr messages. The type is interface{}.
    Pcerr interface{}

    // Logging fallback configuration. The type is interface{}.
    Fallback interface{}
}

func (logging *Pce_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "pce"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("no-path", types.YLeaf{"NoPath", logging.NoPath})
    logging.EntityData.Leafs.Append("pcerr", types.YLeaf{"Pcerr", logging.Pcerr})
    logging.EntityData.Leafs.Append("fallback", types.YLeaf{"Fallback", logging.Fallback})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// Pce_Backoff
// PCE backoff configuration
// This type is a presence type.
type Pce_Backoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Backoff common ratio configuration. The type is interface{} with range:
    // 0..255. The default value is 2.
    Ratio interface{}

    // Backoff threshold configuration. The type is interface{} with range:
    // 0..3600. The default value is 0.
    Threshold interface{}

    // Backoff common difference configuration. The type is interface{} with
    // range: 0..255. The default value is 2.
    Difference interface{}
}

func (backoff *Pce_Backoff) GetEntityData() *types.CommonEntityData {
    backoff.EntityData.YFilter = backoff.YFilter
    backoff.EntityData.YangName = "backoff"
    backoff.EntityData.BundleName = "cisco_ios_xr"
    backoff.EntityData.ParentYangName = "pce"
    backoff.EntityData.SegmentPath = "backoff"
    backoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backoff.EntityData.Children = types.NewOrderedMap()
    backoff.EntityData.Leafs = types.NewOrderedMap()
    backoff.EntityData.Leafs.Append("ratio", types.YLeaf{"Ratio", backoff.Ratio})
    backoff.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", backoff.Threshold})
    backoff.EntityData.Leafs.Append("difference", types.YLeaf{"Difference", backoff.Difference})

    backoff.EntityData.YListKeys = []string {}

    return &(backoff.EntityData)
}

// Pce_Rest
// REST configuration
// This type is a presence type.
type Pce_Rest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // REST authentication type. The type is PcerestAuthentication. This attribute
    // is mandatory.
    RestAuthentication interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // REST authorized users configuration.
    RestUsers Pce_Rest_RestUsers
}

func (rest *Pce_Rest) GetEntityData() *types.CommonEntityData {
    rest.EntityData.YFilter = rest.YFilter
    rest.EntityData.YangName = "rest"
    rest.EntityData.BundleName = "cisco_ios_xr"
    rest.EntityData.ParentYangName = "pce"
    rest.EntityData.SegmentPath = "rest"
    rest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rest.EntityData.Children = types.NewOrderedMap()
    rest.EntityData.Children.Append("rest-users", types.YChild{"RestUsers", &rest.RestUsers})
    rest.EntityData.Leafs = types.NewOrderedMap()
    rest.EntityData.Leafs.Append("rest-authentication", types.YLeaf{"RestAuthentication", rest.RestAuthentication})
    rest.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rest.Enable})

    rest.EntityData.YListKeys = []string {}

    return &(rest.EntityData)
}

// Pce_Rest_RestUsers
// REST authorized users configuration
type Pce_Rest_RestUsers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // REST authorized user. The type is slice of Pce_Rest_RestUsers_RestUser.
    RestUser []*Pce_Rest_RestUsers_RestUser
}

func (restUsers *Pce_Rest_RestUsers) GetEntityData() *types.CommonEntityData {
    restUsers.EntityData.YFilter = restUsers.YFilter
    restUsers.EntityData.YangName = "rest-users"
    restUsers.EntityData.BundleName = "cisco_ios_xr"
    restUsers.EntityData.ParentYangName = "rest"
    restUsers.EntityData.SegmentPath = "rest-users"
    restUsers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    restUsers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    restUsers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    restUsers.EntityData.Children = types.NewOrderedMap()
    restUsers.EntityData.Children.Append("rest-user", types.YChild{"RestUser", nil})
    for i := range restUsers.RestUser {
        restUsers.EntityData.Children.Append(types.GetSegmentPath(restUsers.RestUser[i]), types.YChild{"RestUser", restUsers.RestUser[i]})
    }
    restUsers.EntityData.Leafs = types.NewOrderedMap()

    restUsers.EntityData.YListKeys = []string {}

    return &(restUsers.EntityData)
}

// Pce_Rest_RestUsers_RestUser
// REST authorized user
type Pce_Rest_RestUsers_RestUser struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. User name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // REST user password configuration. The type is string with pattern:
    // (!.+)|([^!].+). This attribute is mandatory.
    RestUserPassword interface{}

    // True only. The type is interface{}.
    Enable interface{}
}

func (restUser *Pce_Rest_RestUsers_RestUser) GetEntityData() *types.CommonEntityData {
    restUser.EntityData.YFilter = restUser.YFilter
    restUser.EntityData.YangName = "rest-user"
    restUser.EntityData.BundleName = "cisco_ios_xr"
    restUser.EntityData.ParentYangName = "rest-users"
    restUser.EntityData.SegmentPath = "rest-user" + types.AddKeyToken(restUser.Name, "name")
    restUser.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    restUser.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    restUser.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    restUser.EntityData.Children = types.NewOrderedMap()
    restUser.EntityData.Leafs = types.NewOrderedMap()
    restUser.EntityData.Leafs.Append("name", types.YLeaf{"Name", restUser.Name})
    restUser.EntityData.Leafs.Append("rest-user-password", types.YLeaf{"RestUserPassword", restUser.RestUserPassword})
    restUser.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", restUser.Enable})

    restUser.EntityData.YListKeys = []string {"Name"}

    return &(restUser.EntityData)
}

// Pce_StateSyncs
// Standby IPv4 PCE configuration
type Pce_StateSyncs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Standby PCE ipv4 address. The type is slice of Pce_StateSyncs_StateSync.
    StateSync []*Pce_StateSyncs_StateSync
}

func (stateSyncs *Pce_StateSyncs) GetEntityData() *types.CommonEntityData {
    stateSyncs.EntityData.YFilter = stateSyncs.YFilter
    stateSyncs.EntityData.YangName = "state-syncs"
    stateSyncs.EntityData.BundleName = "cisco_ios_xr"
    stateSyncs.EntityData.ParentYangName = "pce"
    stateSyncs.EntityData.SegmentPath = "state-syncs"
    stateSyncs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSyncs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSyncs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSyncs.EntityData.Children = types.NewOrderedMap()
    stateSyncs.EntityData.Children.Append("state-sync", types.YChild{"StateSync", nil})
    for i := range stateSyncs.StateSync {
        stateSyncs.EntityData.Children.Append(types.GetSegmentPath(stateSyncs.StateSync[i]), types.YChild{"StateSync", stateSyncs.StateSync[i]})
    }
    stateSyncs.EntityData.Leafs = types.NewOrderedMap()

    stateSyncs.EntityData.YListKeys = []string {}

    return &(stateSyncs.EntityData)
}

// Pce_StateSyncs_StateSync
// Standby PCE ipv4 address
type Pce_StateSyncs_StateSync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (stateSync *Pce_StateSyncs_StateSync) GetEntityData() *types.CommonEntityData {
    stateSync.EntityData.YFilter = stateSync.YFilter
    stateSync.EntityData.YangName = "state-sync"
    stateSync.EntityData.BundleName = "cisco_ios_xr"
    stateSync.EntityData.ParentYangName = "state-syncs"
    stateSync.EntityData.SegmentPath = "state-sync" + types.AddKeyToken(stateSync.Address, "address")
    stateSync.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSync.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSync.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSync.EntityData.Children = types.NewOrderedMap()
    stateSync.EntityData.Leafs = types.NewOrderedMap()
    stateSync.EntityData.Leafs.Append("address", types.YLeaf{"Address", stateSync.Address})

    stateSync.EntityData.YListKeys = []string {"Address"}

    return &(stateSync.EntityData)
}

// Pce_SegmentRouting
// PCE segment-routing configuration
type Pce_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use te-latency algorithm configuration. The type is interface{}.
    TeLatency interface{}

    // Use strict-sid-only configuration. The type is interface{}.
    StrictSidOnly interface{}

    // Traffic Engineering configuration data.
    TrafficEngineering Pce_SegmentRouting_TrafficEngineering
}

func (segmentRouting *Pce_SegmentRouting) GetEntityData() *types.CommonEntityData {
    segmentRouting.EntityData.YFilter = segmentRouting.YFilter
    segmentRouting.EntityData.YangName = "segment-routing"
    segmentRouting.EntityData.BundleName = "cisco_ios_xr"
    segmentRouting.EntityData.ParentYangName = "pce"
    segmentRouting.EntityData.SegmentPath = "segment-routing"
    segmentRouting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRouting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRouting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRouting.EntityData.Children = types.NewOrderedMap()
    segmentRouting.EntityData.Children.Append("traffic-engineering", types.YChild{"TrafficEngineering", &segmentRouting.TrafficEngineering})
    segmentRouting.EntityData.Leafs = types.NewOrderedMap()
    segmentRouting.EntityData.Leafs.Append("te-latency", types.YLeaf{"TeLatency", segmentRouting.TeLatency})
    segmentRouting.EntityData.Leafs.Append("strict-sid-only", types.YLeaf{"StrictSidOnly", segmentRouting.StrictSidOnly})

    segmentRouting.EntityData.YListKeys = []string {}

    return &(segmentRouting.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering
// Traffic Engineering configuration data
type Pce_SegmentRouting_TrafficEngineering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}

    // Affinity Bit-map.
    AffinityBits Pce_SegmentRouting_TrafficEngineering_AffinityBits

    // Peer configuration.
    Peers Pce_SegmentRouting_TrafficEngineering_Peers

    // Segment-lists configuration.
    Segments Pce_SegmentRouting_TrafficEngineering_Segments
}

func (trafficEngineering *Pce_SegmentRouting_TrafficEngineering) GetEntityData() *types.CommonEntityData {
    trafficEngineering.EntityData.YFilter = trafficEngineering.YFilter
    trafficEngineering.EntityData.YangName = "traffic-engineering"
    trafficEngineering.EntityData.BundleName = "cisco_ios_xr"
    trafficEngineering.EntityData.ParentYangName = "segment-routing"
    trafficEngineering.EntityData.SegmentPath = "traffic-engineering"
    trafficEngineering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficEngineering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficEngineering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficEngineering.EntityData.Children = types.NewOrderedMap()
    trafficEngineering.EntityData.Children.Append("affinity-bits", types.YChild{"AffinityBits", &trafficEngineering.AffinityBits})
    trafficEngineering.EntityData.Children.Append("peers", types.YChild{"Peers", &trafficEngineering.Peers})
    trafficEngineering.EntityData.Children.Append("segments", types.YChild{"Segments", &trafficEngineering.Segments})
    trafficEngineering.EntityData.Leafs = types.NewOrderedMap()
    trafficEngineering.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", trafficEngineering.Enable})

    trafficEngineering.EntityData.YListKeys = []string {}

    return &(trafficEngineering.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_AffinityBits
// Affinity Bit-map
type Pce_SegmentRouting_TrafficEngineering_AffinityBits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Affinity Bit. The type is slice of
    // Pce_SegmentRouting_TrafficEngineering_AffinityBits_AffinityBit.
    AffinityBit []*Pce_SegmentRouting_TrafficEngineering_AffinityBits_AffinityBit
}

func (affinityBits *Pce_SegmentRouting_TrafficEngineering_AffinityBits) GetEntityData() *types.CommonEntityData {
    affinityBits.EntityData.YFilter = affinityBits.YFilter
    affinityBits.EntityData.YangName = "affinity-bits"
    affinityBits.EntityData.BundleName = "cisco_ios_xr"
    affinityBits.EntityData.ParentYangName = "traffic-engineering"
    affinityBits.EntityData.SegmentPath = "affinity-bits"
    affinityBits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityBits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityBits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityBits.EntityData.Children = types.NewOrderedMap()
    affinityBits.EntityData.Children.Append("affinity-bit", types.YChild{"AffinityBit", nil})
    for i := range affinityBits.AffinityBit {
        affinityBits.EntityData.Children.Append(types.GetSegmentPath(affinityBits.AffinityBit[i]), types.YChild{"AffinityBit", affinityBits.AffinityBit[i]})
    }
    affinityBits.EntityData.Leafs = types.NewOrderedMap()

    affinityBits.EntityData.YListKeys = []string {}

    return &(affinityBits.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_AffinityBits_AffinityBit
// Affinity Bit
type Pce_SegmentRouting_TrafficEngineering_AffinityBits_AffinityBit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Color Name. The type is string with length: 1..32.
    ColorName interface{}

    // The bit. The type is interface{} with range: 0..31. This attribute is
    // mandatory.
    Bit interface{}
}

func (affinityBit *Pce_SegmentRouting_TrafficEngineering_AffinityBits_AffinityBit) GetEntityData() *types.CommonEntityData {
    affinityBit.EntityData.YFilter = affinityBit.YFilter
    affinityBit.EntityData.YangName = "affinity-bit"
    affinityBit.EntityData.BundleName = "cisco_ios_xr"
    affinityBit.EntityData.ParentYangName = "affinity-bits"
    affinityBit.EntityData.SegmentPath = "affinity-bit" + types.AddKeyToken(affinityBit.ColorName, "color-name")
    affinityBit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityBit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityBit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityBit.EntityData.Children = types.NewOrderedMap()
    affinityBit.EntityData.Leafs = types.NewOrderedMap()
    affinityBit.EntityData.Leafs.Append("color-name", types.YLeaf{"ColorName", affinityBit.ColorName})
    affinityBit.EntityData.Leafs.Append("bit", types.YLeaf{"Bit", affinityBit.Bit})

    affinityBit.EntityData.YListKeys = []string {"ColorName"}

    return &(affinityBit.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers
// Peer configuration
type Pce_SegmentRouting_TrafficEngineering_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer configuration. The type is slice of
    // Pce_SegmentRouting_TrafficEngineering_Peers_Peer.
    Peer []*Pce_SegmentRouting_TrafficEngineering_Peers_Peer
}

func (peers *Pce_SegmentRouting_TrafficEngineering_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "traffic-engineering"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer
// Peer configuration
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddr interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Policy configuration.
    Policies Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies
}

func (peer *Pce_SegmentRouting_TrafficEngineering_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.PeerAddr, "peer-addr")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("policies", types.YChild{"Policies", &peer.Policies})
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("peer-addr", types.YLeaf{"PeerAddr", peer.PeerAddr})
    peer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", peer.Enable})

    peer.EntityData.YListKeys = []string {"PeerAddr"}

    return &(peer.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies
// Policy configuration
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration. The type is slice of
    // Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy.
    Policy []*Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy
}

func (policies *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "cisco_ios_xr"
    policies.EntityData.ParentYangName = "peer"
    policies.EntityData.SegmentPath = "policies"
    policies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policies.EntityData.Children = types.NewOrderedMap()
    policies.EntityData.Children.Append("policy", types.YChild{"Policy", nil})
    for i := range policies.Policy {
        policies.EntityData.Children.Append(types.GetSegmentPath(policies.Policy[i]), types.YChild{"Policy", policies.Policy[i]})
    }
    policies.EntityData.Leafs = types.NewOrderedMap()

    policies.EntityData.YListKeys = []string {}

    return &(policies.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy
// Policy configuration
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Policy name. The type is string with length:
    // 1..128.
    PolicyName interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Administratively shutdown policy. The type is interface{}.
    Shutdown interface{}

    // Binding Segment ID.
    BindingSid Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_BindingSid

    // Color and Endpoint.
    ColorEndpoint Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_ColorEndpoint

    // Policy candidate-paths configuration.
    CandidatePaths Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths
}

func (policy *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "policies"
    policy.EntityData.SegmentPath = "policy" + types.AddKeyToken(policy.PolicyName, "policy-name")
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = types.NewOrderedMap()
    policy.EntityData.Children.Append("binding-sid", types.YChild{"BindingSid", &policy.BindingSid})
    policy.EntityData.Children.Append("color-endpoint", types.YChild{"ColorEndpoint", &policy.ColorEndpoint})
    policy.EntityData.Children.Append("candidate-paths", types.YChild{"CandidatePaths", &policy.CandidatePaths})
    policy.EntityData.Leafs = types.NewOrderedMap()
    policy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", policy.PolicyName})
    policy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", policy.Enable})
    policy.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", policy.Shutdown})

    policy.EntityData.YListKeys = []string {"PolicyName"}

    return &(policy.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_BindingSid
// Binding Segment ID
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_BindingSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Binding SID type. The type is PceBindingSid.
    BindingSidType interface{}

    // MPLS Label. The type is interface{} with range: 16..1048575.
    MplsLabel interface{}
}

func (bindingSid *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_BindingSid) GetEntityData() *types.CommonEntityData {
    bindingSid.EntityData.YFilter = bindingSid.YFilter
    bindingSid.EntityData.YangName = "binding-sid"
    bindingSid.EntityData.BundleName = "cisco_ios_xr"
    bindingSid.EntityData.ParentYangName = "policy"
    bindingSid.EntityData.SegmentPath = "binding-sid"
    bindingSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingSid.EntityData.Children = types.NewOrderedMap()
    bindingSid.EntityData.Leafs = types.NewOrderedMap()
    bindingSid.EntityData.Leafs.Append("binding-sid-type", types.YLeaf{"BindingSidType", bindingSid.BindingSidType})
    bindingSid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", bindingSid.MplsLabel})

    bindingSid.EntityData.YListKeys = []string {}

    return &(bindingSid.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_ColorEndpoint
// Color and Endpoint
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_ColorEndpoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Color. The type is interface{} with range: 1..4294967295.
    Color interface{}

    // End point type. The type is PceEndPoint.
    EndPointType interface{}

    // End point address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    EndPointAddress interface{}
}

func (colorEndpoint *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_ColorEndpoint) GetEntityData() *types.CommonEntityData {
    colorEndpoint.EntityData.YFilter = colorEndpoint.YFilter
    colorEndpoint.EntityData.YangName = "color-endpoint"
    colorEndpoint.EntityData.BundleName = "cisco_ios_xr"
    colorEndpoint.EntityData.ParentYangName = "policy"
    colorEndpoint.EntityData.SegmentPath = "color-endpoint"
    colorEndpoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    colorEndpoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    colorEndpoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    colorEndpoint.EntityData.Children = types.NewOrderedMap()
    colorEndpoint.EntityData.Leafs = types.NewOrderedMap()
    colorEndpoint.EntityData.Leafs.Append("color", types.YLeaf{"Color", colorEndpoint.Color})
    colorEndpoint.EntityData.Leafs.Append("end-point-type", types.YLeaf{"EndPointType", colorEndpoint.EndPointType})
    colorEndpoint.EntityData.Leafs.Append("end-point-address", types.YLeaf{"EndPointAddress", colorEndpoint.EndPointAddress})

    colorEndpoint.EntityData.YListKeys = []string {}

    return &(colorEndpoint.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths
// Policy candidate-paths configuration
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}

    // Affinity rule table.
    AffinityRules Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules

    // Policy path-option preference table.
    Preferences Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences
}

func (candidatePaths *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths) GetEntityData() *types.CommonEntityData {
    candidatePaths.EntityData.YFilter = candidatePaths.YFilter
    candidatePaths.EntityData.YangName = "candidate-paths"
    candidatePaths.EntityData.BundleName = "cisco_ios_xr"
    candidatePaths.EntityData.ParentYangName = "policy"
    candidatePaths.EntityData.SegmentPath = "candidate-paths"
    candidatePaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidatePaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidatePaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidatePaths.EntityData.Children = types.NewOrderedMap()
    candidatePaths.EntityData.Children.Append("affinity-rules", types.YChild{"AffinityRules", &candidatePaths.AffinityRules})
    candidatePaths.EntityData.Children.Append("preferences", types.YChild{"Preferences", &candidatePaths.Preferences})
    candidatePaths.EntityData.Leafs = types.NewOrderedMap()
    candidatePaths.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", candidatePaths.Enable})

    candidatePaths.EntityData.YListKeys = []string {}

    return &(candidatePaths.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules
// Affinity rule table
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Affinity rule. The type is slice of
    // Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules_AffinityRule.
    AffinityRule []*Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules_AffinityRule
}

func (affinityRules *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules) GetEntityData() *types.CommonEntityData {
    affinityRules.EntityData.YFilter = affinityRules.YFilter
    affinityRules.EntityData.YangName = "affinity-rules"
    affinityRules.EntityData.BundleName = "cisco_ios_xr"
    affinityRules.EntityData.ParentYangName = "candidate-paths"
    affinityRules.EntityData.SegmentPath = "affinity-rules"
    affinityRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityRules.EntityData.Children = types.NewOrderedMap()
    affinityRules.EntityData.Children.Append("affinity-rule", types.YChild{"AffinityRule", nil})
    for i := range affinityRules.AffinityRule {
        affinityRules.EntityData.Children.Append(types.GetSegmentPath(affinityRules.AffinityRule[i]), types.YChild{"AffinityRule", affinityRules.AffinityRule[i]})
    }
    affinityRules.EntityData.Leafs = types.NewOrderedMap()

    affinityRules.EntityData.YListKeys = []string {}

    return &(affinityRules.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules_AffinityRule
// Affinity rule
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules_AffinityRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. affinity rule. The type is interface{} with range:
    // 0..2.
    Rule interface{}

    // This attribute is a key. affinity value. The type is string with length:
    // 1..32.
    AffValue interface{}
}

func (affinityRule *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_AffinityRules_AffinityRule) GetEntityData() *types.CommonEntityData {
    affinityRule.EntityData.YFilter = affinityRule.YFilter
    affinityRule.EntityData.YangName = "affinity-rule"
    affinityRule.EntityData.BundleName = "cisco_ios_xr"
    affinityRule.EntityData.ParentYangName = "affinity-rules"
    affinityRule.EntityData.SegmentPath = "affinity-rule" + types.AddKeyToken(affinityRule.Rule, "rule") + types.AddKeyToken(affinityRule.AffValue, "aff-value")
    affinityRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityRule.EntityData.Children = types.NewOrderedMap()
    affinityRule.EntityData.Leafs = types.NewOrderedMap()
    affinityRule.EntityData.Leafs.Append("rule", types.YLeaf{"Rule", affinityRule.Rule})
    affinityRule.EntityData.Leafs.Append("aff-value", types.YLeaf{"AffValue", affinityRule.AffValue})

    affinityRule.EntityData.YListKeys = []string {"Rule", "AffValue"}

    return &(affinityRule.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences
// Policy path-option preference table
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy path-option preference entry. The type is slice of
    // Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference.
    Preference []*Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference
}

func (preferences *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences) GetEntityData() *types.CommonEntityData {
    preferences.EntityData.YFilter = preferences.YFilter
    preferences.EntityData.YangName = "preferences"
    preferences.EntityData.BundleName = "cisco_ios_xr"
    preferences.EntityData.ParentYangName = "candidate-paths"
    preferences.EntityData.SegmentPath = "preferences"
    preferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preferences.EntityData.Children = types.NewOrderedMap()
    preferences.EntityData.Children.Append("preference", types.YChild{"Preference", nil})
    for i := range preferences.Preference {
        preferences.EntityData.Children.Append(types.GetSegmentPath(preferences.Preference[i]), types.YChild{"Preference", preferences.Preference[i]})
    }
    preferences.EntityData.Leafs = types.NewOrderedMap()

    preferences.EntityData.YListKeys = []string {}

    return &(preferences.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference
// Policy path-option preference entry
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path-option preference. The type is interface{}
    // with range: 1..65535.
    PathIndex interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Policy path-option preference configuration.
    PathInfos Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos
}

func (preference *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference) GetEntityData() *types.CommonEntityData {
    preference.EntityData.YFilter = preference.YFilter
    preference.EntityData.YangName = "preference"
    preference.EntityData.BundleName = "cisco_ios_xr"
    preference.EntityData.ParentYangName = "preferences"
    preference.EntityData.SegmentPath = "preference" + types.AddKeyToken(preference.PathIndex, "path-index")
    preference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preference.EntityData.Children = types.NewOrderedMap()
    preference.EntityData.Children.Append("path-infos", types.YChild{"PathInfos", &preference.PathInfos})
    preference.EntityData.Leafs = types.NewOrderedMap()
    preference.EntityData.Leafs.Append("path-index", types.YLeaf{"PathIndex", preference.PathIndex})
    preference.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", preference.Enable})

    preference.EntityData.YListKeys = []string {"PathIndex"}

    return &(preference.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos
// Policy path-option preference
// configuration
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration. The type is slice of
    // Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo.
    PathInfo []*Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo
}

func (pathInfos *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos) GetEntityData() *types.CommonEntityData {
    pathInfos.EntityData.YFilter = pathInfos.YFilter
    pathInfos.EntityData.YangName = "path-infos"
    pathInfos.EntityData.BundleName = "cisco_ios_xr"
    pathInfos.EntityData.ParentYangName = "preference"
    pathInfos.EntityData.SegmentPath = "path-infos"
    pathInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfos.EntityData.Children = types.NewOrderedMap()
    pathInfos.EntityData.Children.Append("path-info", types.YChild{"PathInfo", nil})
    for i := range pathInfos.PathInfo {
        pathInfos.EntityData.Children.Append(types.GetSegmentPath(pathInfos.PathInfo[i]), types.YChild{"PathInfo", pathInfos.PathInfo[i]})
    }
    pathInfos.EntityData.Leafs = types.NewOrderedMap()

    pathInfos.EntityData.YListKeys = []string {}

    return &(pathInfos.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo
// Policy configuration
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path-option type. The type is PcePath.
    Type interface{}

    // This attribute is a key. Type of dynamic path to be computed. The type is
    // PcePathHop.
    HopType interface{}

    // This attribute is a key. Segment-list name. The type is string with length:
    // 1..128.
    SegmentListName interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Metric configuration, valid only for dynamic path-options.
    Metric Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric
}

func (pathInfo *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo) GetEntityData() *types.CommonEntityData {
    pathInfo.EntityData.YFilter = pathInfo.YFilter
    pathInfo.EntityData.YangName = "path-info"
    pathInfo.EntityData.BundleName = "cisco_ios_xr"
    pathInfo.EntityData.ParentYangName = "path-infos"
    pathInfo.EntityData.SegmentPath = "path-info" + types.AddKeyToken(pathInfo.Type, "type") + types.AddKeyToken(pathInfo.HopType, "hop-type") + types.AddKeyToken(pathInfo.SegmentListName, "segment-list-name")
    pathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfo.EntityData.Children = types.NewOrderedMap()
    pathInfo.EntityData.Children.Append("metric", types.YChild{"Metric", &pathInfo.Metric})
    pathInfo.EntityData.Leafs = types.NewOrderedMap()
    pathInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathInfo.Type})
    pathInfo.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", pathInfo.HopType})
    pathInfo.EntityData.Leafs.Append("segment-list-name", types.YLeaf{"SegmentListName", pathInfo.SegmentListName})
    pathInfo.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathInfo.Enable})

    pathInfo.EntityData.YListKeys = []string {"Type", "HopType", "SegmentListName"}

    return &(pathInfo.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric
// Metric configuration, valid only for
// dynamic path-options
// This type is a presence type.
type Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Metric type. The type is PceMetric. This attribute is mandatory.
    MetricType interface{}
}

func (metric *Pce_SegmentRouting_TrafficEngineering_Peers_Peer_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "path-info"
    metric.EntityData.SegmentPath = "metric"
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = types.NewOrderedMap()
    metric.EntityData.Leafs = types.NewOrderedMap()
    metric.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", metric.MetricType})

    metric.EntityData.YListKeys = []string {}

    return &(metric.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Segments
// Segment-lists configuration
type Pce_SegmentRouting_TrafficEngineering_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment-list configuration. The type is slice of
    // Pce_SegmentRouting_TrafficEngineering_Segments_Segment.
    Segment []*Pce_SegmentRouting_TrafficEngineering_Segments_Segment
}

func (segments *Pce_SegmentRouting_TrafficEngineering_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "traffic-engineering"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = types.NewOrderedMap()
    segments.EntityData.Children.Append("segment", types.YChild{"Segment", nil})
    for i := range segments.Segment {
        segments.EntityData.Children.Append(types.GetSegmentPath(segments.Segment[i]), types.YChild{"Segment", segments.Segment[i]})
    }
    segments.EntityData.Leafs = types.NewOrderedMap()

    segments.EntityData.YListKeys = []string {}

    return &(segments.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Segments_Segment
// Segment-list configuration
type Pce_SegmentRouting_TrafficEngineering_Segments_Segment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Segment-list name. The type is string with length:
    // 1..128.
    PathName interface{}

    // Segments/hops configuration for given Segment-list.
    Segments Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments
}

func (segment *Pce_SegmentRouting_TrafficEngineering_Segments_Segment) GetEntityData() *types.CommonEntityData {
    segment.EntityData.YFilter = segment.YFilter
    segment.EntityData.YangName = "segment"
    segment.EntityData.BundleName = "cisco_ios_xr"
    segment.EntityData.ParentYangName = "segments"
    segment.EntityData.SegmentPath = "segment" + types.AddKeyToken(segment.PathName, "path-name")
    segment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segment.EntityData.Children = types.NewOrderedMap()
    segment.EntityData.Children.Append("segments", types.YChild{"Segments", &segment.Segments})
    segment.EntityData.Leafs = types.NewOrderedMap()
    segment.EntityData.Leafs.Append("path-name", types.YLeaf{"PathName", segment.PathName})

    segment.EntityData.YListKeys = []string {"PathName"}

    return &(segment.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments
// Segments/hops configuration for given
// Segment-list
type Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Segment/hop at the index. The type is slice of
    // Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments_Segment.
    Segment []*Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments_Segment
}

func (segments *Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "segment"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = types.NewOrderedMap()
    segments.EntityData.Children.Append("segment", types.YChild{"Segment", nil})
    for i := range segments.Segment {
        segments.EntityData.Children.Append(types.GetSegmentPath(segments.Segment[i]), types.YChild{"Segment", segments.Segment[i]})
    }
    segments.EntityData.Leafs = types.NewOrderedMap()

    segments.EntityData.YListKeys = []string {}

    return &(segments.EntityData)
}

// Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments_Segment
// Configure Segment/hop at the index
type Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments_Segment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Segment index. The type is interface{} with range:
    // 1..65535.
    SegmentIndex interface{}

    // Segment/hop type. The type is PceSegment.
    SegmentType interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // MPLS Label. The type is interface{} with range: 0..1048575.
    MplsLabel interface{}
}

func (segment *Pce_SegmentRouting_TrafficEngineering_Segments_Segment_Segments_Segment) GetEntityData() *types.CommonEntityData {
    segment.EntityData.YFilter = segment.YFilter
    segment.EntityData.YangName = "segment"
    segment.EntityData.BundleName = "cisco_ios_xr"
    segment.EntityData.ParentYangName = "segments"
    segment.EntityData.SegmentPath = "segment" + types.AddKeyToken(segment.SegmentIndex, "segment-index")
    segment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segment.EntityData.Children = types.NewOrderedMap()
    segment.EntityData.Leafs = types.NewOrderedMap()
    segment.EntityData.Leafs.Append("segment-index", types.YLeaf{"SegmentIndex", segment.SegmentIndex})
    segment.EntityData.Leafs.Append("segment-type", types.YLeaf{"SegmentType", segment.SegmentType})
    segment.EntityData.Leafs.Append("address", types.YLeaf{"Address", segment.Address})
    segment.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", segment.MplsLabel})

    segment.EntityData.YListKeys = []string {"SegmentIndex"}

    return &(segment.EntityData)
}

// Pce_Timers
// PCE Timers configuration
// This type is a presence type.
type Pce_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Topology reoptimization timer configuration. The type is interface{} with
    // range: 10..3600. Units are second. The default value is 60.
    ReoptimizationTimer interface{}

    // Keepalive interval in seconds, zero to disable. The type is interface{}
    // with range: 0..255. Units are second. The default value is 30.
    Keepalive interface{}

    // Minimum acceptable peer proposed keepalive interval. The type is
    // interface{} with range: 0..255. Units are second. The default value is 20.
    MinimumPeerKeepalive interface{}
}

func (timers *Pce_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "pce"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Leafs = types.NewOrderedMap()
    timers.EntityData.Leafs.Append("reoptimization-timer", types.YLeaf{"ReoptimizationTimer", timers.ReoptimizationTimer})
    timers.EntityData.Leafs.Append("keepalive", types.YLeaf{"Keepalive", timers.Keepalive})
    timers.EntityData.Leafs.Append("minimum-peer-keepalive", types.YLeaf{"MinimumPeerKeepalive", timers.MinimumPeerKeepalive})

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Pce_Netconf
// NETCONF configuration
type Pce_Netconf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}

    // NETCONF SSH configuration.
    NetconfSsh Pce_Netconf_NetconfSsh
}

func (netconf *Pce_Netconf) GetEntityData() *types.CommonEntityData {
    netconf.EntityData.YFilter = netconf.YFilter
    netconf.EntityData.YangName = "netconf"
    netconf.EntityData.BundleName = "cisco_ios_xr"
    netconf.EntityData.ParentYangName = "pce"
    netconf.EntityData.SegmentPath = "netconf"
    netconf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconf.EntityData.Children = types.NewOrderedMap()
    netconf.EntityData.Children.Append("netconf-ssh", types.YChild{"NetconfSsh", &netconf.NetconfSsh})
    netconf.EntityData.Leafs = types.NewOrderedMap()
    netconf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", netconf.Enable})

    netconf.EntityData.YListKeys = []string {}

    return &(netconf.EntityData)
}

// Pce_Netconf_NetconfSsh
// NETCONF SSH configuration
type Pce_Netconf_NetconfSsh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Password to use for NETCONF SSH connections. The type is string with
    // pattern: (!.+)|([^!].+).
    NetconfSshPassword interface{}

    // User name to use for NETCONF SSH connections. The type is string.
    NetconfSshUser interface{}
}

func (netconfSsh *Pce_Netconf_NetconfSsh) GetEntityData() *types.CommonEntityData {
    netconfSsh.EntityData.YFilter = netconfSsh.YFilter
    netconfSsh.EntityData.YangName = "netconf-ssh"
    netconfSsh.EntityData.BundleName = "cisco_ios_xr"
    netconfSsh.EntityData.ParentYangName = "netconf"
    netconfSsh.EntityData.SegmentPath = "netconf-ssh"
    netconfSsh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconfSsh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconfSsh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconfSsh.EntityData.Children = types.NewOrderedMap()
    netconfSsh.EntityData.Leafs = types.NewOrderedMap()
    netconfSsh.EntityData.Leafs.Append("netconf-ssh-password", types.YLeaf{"NetconfSshPassword", netconfSsh.NetconfSshPassword})
    netconfSsh.EntityData.Leafs.Append("netconf-ssh-user", types.YLeaf{"NetconfSshUser", netconfSsh.NetconfSshUser})

    netconfSsh.EntityData.YListKeys = []string {}

    return &(netconfSsh.EntityData)
}

// Pce_DisjointPath
// Disjoint path configuration
type Pce_DisjointPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}

    // Association configuration.
    Groups Pce_DisjointPath_Groups
}

func (disjointPath *Pce_DisjointPath) GetEntityData() *types.CommonEntityData {
    disjointPath.EntityData.YFilter = disjointPath.YFilter
    disjointPath.EntityData.YangName = "disjoint-path"
    disjointPath.EntityData.BundleName = "cisco_ios_xr"
    disjointPath.EntityData.ParentYangName = "pce"
    disjointPath.EntityData.SegmentPath = "disjoint-path"
    disjointPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disjointPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disjointPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disjointPath.EntityData.Children = types.NewOrderedMap()
    disjointPath.EntityData.Children.Append("groups", types.YChild{"Groups", &disjointPath.Groups})
    disjointPath.EntityData.Leafs = types.NewOrderedMap()
    disjointPath.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", disjointPath.Enable})

    disjointPath.EntityData.YListKeys = []string {}

    return &(disjointPath.EntityData)
}

// Pce_DisjointPath_Groups
// Association configuration
type Pce_DisjointPath_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association Group Configuration. The type is slice of
    // Pce_DisjointPath_Groups_Group.
    Group []*Pce_DisjointPath_Groups_Group
}

func (groups *Pce_DisjointPath_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "disjoint-path"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Pce_DisjointPath_Groups_Group
// Association Group Configuration
type Pce_DisjointPath_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 1..65535.
    GroupId interface{}

    // This attribute is a key. Disjointness type. The type is PceDisjointPath.
    DpType interface{}

    // This attribute is a key. Sub group ID, 0 to unset. The type is interface{}
    // with range: 0..65535.
    SubId interface{}

    // Disable Fallback. The type is interface{}.
    Strict interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // lsp pcc records container with in group.
    GroupLspRecords Pce_DisjointPath_Groups_Group_GroupLspRecords
}

func (group *Pce_DisjointPath_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupId, "group-id") + types.AddKeyToken(group.DpType, "dp-type") + types.AddKeyToken(group.SubId, "sub-id")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("group-lsp-records", types.YChild{"GroupLspRecords", &group.GroupLspRecords})
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", group.GroupId})
    group.EntityData.Leafs.Append("dp-type", types.YLeaf{"DpType", group.DpType})
    group.EntityData.Leafs.Append("sub-id", types.YLeaf{"SubId", group.SubId})
    group.EntityData.Leafs.Append("strict", types.YLeaf{"Strict", group.Strict})
    group.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", group.Enable})

    group.EntityData.YListKeys = []string {"GroupId", "DpType", "SubId"}

    return &(group.EntityData)
}

// Pce_DisjointPath_Groups_Group_GroupLspRecords
// lsp pcc records container with in group
type Pce_DisjointPath_Groups_Group_GroupLspRecords struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP first/second PCC record tuple containingIpAddr, LspName, DisjPath. The
    // type is slice of
    // Pce_DisjointPath_Groups_Group_GroupLspRecords_GroupLspRecord.
    GroupLspRecord []*Pce_DisjointPath_Groups_Group_GroupLspRecords_GroupLspRecord
}

func (groupLspRecords *Pce_DisjointPath_Groups_Group_GroupLspRecords) GetEntityData() *types.CommonEntityData {
    groupLspRecords.EntityData.YFilter = groupLspRecords.YFilter
    groupLspRecords.EntityData.YangName = "group-lsp-records"
    groupLspRecords.EntityData.BundleName = "cisco_ios_xr"
    groupLspRecords.EntityData.ParentYangName = "group"
    groupLspRecords.EntityData.SegmentPath = "group-lsp-records"
    groupLspRecords.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupLspRecords.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupLspRecords.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupLspRecords.EntityData.Children = types.NewOrderedMap()
    groupLspRecords.EntityData.Children.Append("group-lsp-record", types.YChild{"GroupLspRecord", nil})
    for i := range groupLspRecords.GroupLspRecord {
        groupLspRecords.EntityData.Children.Append(types.GetSegmentPath(groupLspRecords.GroupLspRecord[i]), types.YChild{"GroupLspRecord", groupLspRecords.GroupLspRecord[i]})
    }
    groupLspRecords.EntityData.Leafs = types.NewOrderedMap()

    groupLspRecords.EntityData.YListKeys = []string {}

    return &(groupLspRecords.EntityData)
}

// Pce_DisjointPath_Groups_Group_GroupLspRecords_GroupLspRecord
// LSP first/second PCC record tuple
// containingIpAddr, LspName, DisjPath
type Pce_DisjointPath_Groups_Group_GroupLspRecords_GroupLspRecord struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Lsp id. The type is interface{} with range: 1..2.
    LspId interface{}

    // IP address of PCC. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddr interface{}

    // Identifying name for LSP. The type is string.
    LspName interface{}

    // Set LSP to follow shortest-path. The type is interface{} with range:
    // 0..4294967295.
    DisjPath interface{}
}

func (groupLspRecord *Pce_DisjointPath_Groups_Group_GroupLspRecords_GroupLspRecord) GetEntityData() *types.CommonEntityData {
    groupLspRecord.EntityData.YFilter = groupLspRecord.YFilter
    groupLspRecord.EntityData.YangName = "group-lsp-record"
    groupLspRecord.EntityData.BundleName = "cisco_ios_xr"
    groupLspRecord.EntityData.ParentYangName = "group-lsp-records"
    groupLspRecord.EntityData.SegmentPath = "group-lsp-record" + types.AddKeyToken(groupLspRecord.LspId, "lsp-id")
    groupLspRecord.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupLspRecord.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupLspRecord.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupLspRecord.EntityData.Children = types.NewOrderedMap()
    groupLspRecord.EntityData.Leafs = types.NewOrderedMap()
    groupLspRecord.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", groupLspRecord.LspId})
    groupLspRecord.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", groupLspRecord.IpAddr})
    groupLspRecord.EntityData.Leafs.Append("lsp-name", types.YLeaf{"LspName", groupLspRecord.LspName})
    groupLspRecord.EntityData.Leafs.Append("disj-path", types.YLeaf{"DisjPath", groupLspRecord.DisjPath})

    groupLspRecord.EntityData.YListKeys = []string {"LspId"}

    return &(groupLspRecord.EntityData)
}

// Pce_ExplicitPaths
// Explicit paths
type Pce_ExplicitPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Explicit-path configuration. The type is slice of
    // Pce_ExplicitPaths_ExplicitPath.
    ExplicitPath []*Pce_ExplicitPaths_ExplicitPath
}

func (explicitPaths *Pce_ExplicitPaths) GetEntityData() *types.CommonEntityData {
    explicitPaths.EntityData.YFilter = explicitPaths.YFilter
    explicitPaths.EntityData.YangName = "explicit-paths"
    explicitPaths.EntityData.BundleName = "cisco_ios_xr"
    explicitPaths.EntityData.ParentYangName = "pce"
    explicitPaths.EntityData.SegmentPath = "explicit-paths"
    explicitPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitPaths.EntityData.Children = types.NewOrderedMap()
    explicitPaths.EntityData.Children.Append("explicit-path", types.YChild{"ExplicitPath", nil})
    for i := range explicitPaths.ExplicitPath {
        explicitPaths.EntityData.Children.Append(types.GetSegmentPath(explicitPaths.ExplicitPath[i]), types.YChild{"ExplicitPath", explicitPaths.ExplicitPath[i]})
    }
    explicitPaths.EntityData.Leafs = types.NewOrderedMap()

    explicitPaths.EntityData.YListKeys = []string {}

    return &(explicitPaths.EntityData)
}

// Pce_ExplicitPaths_ExplicitPath
// Explicit-path configuration
type Pce_ExplicitPaths_ExplicitPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Explicit-path name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Path Hops.
    PathHops Pce_ExplicitPaths_ExplicitPath_PathHops
}

func (explicitPath *Pce_ExplicitPaths_ExplicitPath) GetEntityData() *types.CommonEntityData {
    explicitPath.EntityData.YFilter = explicitPath.YFilter
    explicitPath.EntityData.YangName = "explicit-path"
    explicitPath.EntityData.BundleName = "cisco_ios_xr"
    explicitPath.EntityData.ParentYangName = "explicit-paths"
    explicitPath.EntityData.SegmentPath = "explicit-path" + types.AddKeyToken(explicitPath.Name, "name")
    explicitPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitPath.EntityData.Children = types.NewOrderedMap()
    explicitPath.EntityData.Children.Append("path-hops", types.YChild{"PathHops", &explicitPath.PathHops})
    explicitPath.EntityData.Leafs = types.NewOrderedMap()
    explicitPath.EntityData.Leafs.Append("name", types.YLeaf{"Name", explicitPath.Name})
    explicitPath.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitPath.Enable})

    explicitPath.EntityData.YListKeys = []string {"Name"}

    return &(explicitPath.EntityData)
}

// Pce_ExplicitPaths_ExplicitPath_PathHops
// Path Hops
type Pce_ExplicitPaths_ExplicitPath_PathHops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Explicit path hop configuration. The type is slice of
    // Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop.
    PathHop []*Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop
}

func (pathHops *Pce_ExplicitPaths_ExplicitPath_PathHops) GetEntityData() *types.CommonEntityData {
    pathHops.EntityData.YFilter = pathHops.YFilter
    pathHops.EntityData.YangName = "path-hops"
    pathHops.EntityData.BundleName = "cisco_ios_xr"
    pathHops.EntityData.ParentYangName = "explicit-path"
    pathHops.EntityData.SegmentPath = "path-hops"
    pathHops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathHops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathHops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathHops.EntityData.Children = types.NewOrderedMap()
    pathHops.EntityData.Children.Append("path-hop", types.YChild{"PathHop", nil})
    for i := range pathHops.PathHop {
        pathHops.EntityData.Children.Append(types.GetSegmentPath(pathHops.PathHop[i]), types.YChild{"PathHop", pathHops.PathHop[i]})
    }
    pathHops.EntityData.Leafs = types.NewOrderedMap()

    pathHops.EntityData.YListKeys = []string {}

    return &(pathHops.EntityData)
}

// Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop
// Explicit path hop configuration
type Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Hop Index. The type is interface{} with range:
    // 1..65535.
    Index interface{}

    // Path hop type. The type is PceExplicitPathHop.
    HopType interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // The default value is 0.0.0.0.
    Address interface{}

    // Remote IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // The default value is 0.0.0.0.
    RemoteAddress interface{}

    // MPLS Label. The type is interface{} with range: 0..1048575. The default
    // value is 0.
    MplsLabel interface{}
}

func (pathHop *Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop) GetEntityData() *types.CommonEntityData {
    pathHop.EntityData.YFilter = pathHop.YFilter
    pathHop.EntityData.YangName = "path-hop"
    pathHop.EntityData.BundleName = "cisco_ios_xr"
    pathHop.EntityData.ParentYangName = "path-hops"
    pathHop.EntityData.SegmentPath = "path-hop" + types.AddKeyToken(pathHop.Index, "index")
    pathHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathHop.EntityData.Children = types.NewOrderedMap()
    pathHop.EntityData.Leafs = types.NewOrderedMap()
    pathHop.EntityData.Leafs.Append("index", types.YLeaf{"Index", pathHop.Index})
    pathHop.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", pathHop.HopType})
    pathHop.EntityData.Leafs.Append("address", types.YLeaf{"Address", pathHop.Address})
    pathHop.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", pathHop.RemoteAddress})
    pathHop.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", pathHop.MplsLabel})

    pathHop.EntityData.YListKeys = []string {"Index"}

    return &(pathHop.EntityData)
}

