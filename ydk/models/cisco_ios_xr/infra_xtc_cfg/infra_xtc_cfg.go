// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-xtc package configuration.
// 
// This module contains definitions
// for the following management objects:
//   pce: PCE configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// Pce
// PCE configuration data
type Pce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address of PCE server. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    ServerAddress interface{}

    // MD5 password. The type is string with pattern: b'(!.+)|([^!].+)'.
    Password interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Path computation client configuration.
    PccAddresses Pce_PccAddresses

    // PCE logging configuration.
    Logging Pce_Logging

    // PCE backoff configuration.
    Backoff Pce_Backoff

    // Standby PCE configuration.
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

    pce.EntityData.Children = make(map[string]types.YChild)
    pce.EntityData.Children["pcc-addresses"] = types.YChild{"PccAddresses", &pce.PccAddresses}
    pce.EntityData.Children["logging"] = types.YChild{"Logging", &pce.Logging}
    pce.EntityData.Children["backoff"] = types.YChild{"Backoff", &pce.Backoff}
    pce.EntityData.Children["state-syncs"] = types.YChild{"StateSyncs", &pce.StateSyncs}
    pce.EntityData.Children["segment-routing"] = types.YChild{"SegmentRouting", &pce.SegmentRouting}
    pce.EntityData.Children["timers"] = types.YChild{"Timers", &pce.Timers}
    pce.EntityData.Children["netconf"] = types.YChild{"Netconf", &pce.Netconf}
    pce.EntityData.Children["disjoint-path"] = types.YChild{"DisjointPath", &pce.DisjointPath}
    pce.EntityData.Children["explicit-paths"] = types.YChild{"ExplicitPaths", &pce.ExplicitPaths}
    pce.EntityData.Leafs = make(map[string]types.YLeaf)
    pce.EntityData.Leafs["server-address"] = types.YLeaf{"ServerAddress", pce.ServerAddress}
    pce.EntityData.Leafs["password"] = types.YLeaf{"Password", pce.Password}
    pce.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pce.Enable}
    return &(pce.EntityData)
}

// Pce_PccAddresses
// Path computation client configuration
type Pce_PccAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path computation client address. The type is slice of
    // Pce_PccAddresses_PccAddress.
    PccAddress []Pce_PccAddresses_PccAddress
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

    pccAddresses.EntityData.Children = make(map[string]types.YChild)
    pccAddresses.EntityData.Children["pcc-address"] = types.YChild{"PccAddress", nil}
    for i := range pccAddresses.PccAddress {
        pccAddresses.EntityData.Children[types.GetSegmentPath(&pccAddresses.PccAddress[i])] = types.YChild{"PccAddress", &pccAddresses.PccAddress[i]}
    }
    pccAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pccAddresses.EntityData)
}

// Pce_PccAddresses_PccAddress
// Path computation client address
type Pce_PccAddresses_PccAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    pccAddress.EntityData.SegmentPath = "pcc-address" + "[address='" + fmt.Sprintf("%v", pccAddress.Address) + "']"
    pccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pccAddress.EntityData.Children = make(map[string]types.YChild)
    pccAddress.EntityData.Children["lsp-names"] = types.YChild{"LspNames", &pccAddress.LspNames}
    pccAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    pccAddress.EntityData.Leafs["address"] = types.YLeaf{"Address", pccAddress.Address}
    pccAddress.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pccAddress.Enable}
    return &(pccAddress.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames
// MPLS label switched path
type Pce_PccAddresses_PccAddress_LspNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS label switched path. The type is slice of
    // Pce_PccAddresses_PccAddress_LspNames_LspName.
    LspName []Pce_PccAddresses_PccAddress_LspNames_LspName
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

    lspNames.EntityData.Children = make(map[string]types.YChild)
    lspNames.EntityData.Children["lsp-name"] = types.YChild{"LspName", nil}
    for i := range lspNames.LspName {
        lspNames.EntityData.Children[types.GetSegmentPath(&lspNames.LspName[i])] = types.YChild{"LspName", &lspNames.LspName[i]}
    }
    lspNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspNames.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames_LspName
// MPLS label switched path
type Pce_PccAddresses_PccAddress_LspNames_LspName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSP name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Undelegate LSP. The type is interface{}.
    Undelegate interface{}

    // Explicit-path name. The type is string.
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
    lspName.EntityData.SegmentPath = "lsp-name" + "[name='" + fmt.Sprintf("%v", lspName.Name) + "']"
    lspName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspName.EntityData.Children = make(map[string]types.YChild)
    lspName.EntityData.Children["rsvp-te"] = types.YChild{"RsvpTe", &lspName.RsvpTe}
    lspName.EntityData.Leafs = make(map[string]types.YLeaf)
    lspName.EntityData.Leafs["name"] = types.YLeaf{"Name", lspName.Name}
    lspName.EntityData.Leafs["undelegate"] = types.YLeaf{"Undelegate", lspName.Undelegate}
    lspName.EntityData.Leafs["explicit-path-name"] = types.YLeaf{"ExplicitPathName", lspName.ExplicitPathName}
    lspName.EntityData.Leafs["enable"] = types.YLeaf{"Enable", lspName.Enable}
    return &(lspName.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe
// RSVP-TE configuration
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable fast protection. The type is interface{}.
    FastProtect interface{}

    // Bandwidth configuration. The type is interface{} with range:
    // -2147483648..2147483647. Units are kbit/s.
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

    rsvpTe.EntityData.Children = make(map[string]types.YChild)
    rsvpTe.EntityData.Children["affinity"] = types.YChild{"Affinity", &rsvpTe.Affinity}
    rsvpTe.EntityData.Children["priority"] = types.YChild{"Priority", &rsvpTe.Priority}
    rsvpTe.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpTe.EntityData.Leafs["fast-protect"] = types.YLeaf{"FastProtect", rsvpTe.FastProtect}
    rsvpTe.EntityData.Leafs["bandwidth"] = types.YLeaf{"Bandwidth", rsvpTe.Bandwidth}
    rsvpTe.EntityData.Leafs["enable"] = types.YLeaf{"Enable", rsvpTe.Enable}
    return &(rsvpTe.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity
// LSP Affinity
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Affinity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Include-any affinity value. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    IncludeAny interface{}

    // Include-all affinity value. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    IncludeAll interface{}

    // Exclude-any affinity value. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
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

    affinity.EntityData.Children = make(map[string]types.YChild)
    affinity.EntityData.Leafs = make(map[string]types.YLeaf)
    affinity.EntityData.Leafs["include-any"] = types.YLeaf{"IncludeAny", affinity.IncludeAny}
    affinity.EntityData.Leafs["include-all"] = types.YLeaf{"IncludeAll", affinity.IncludeAll}
    affinity.EntityData.Leafs["exclude-any"] = types.YLeaf{"ExcludeAny", affinity.ExcludeAny}
    return &(affinity.EntityData)
}

// Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority
// Tunnel Setup and Hold Priorities
// This type is a presence type.
type Pce_PccAddresses_PccAddress_LspNames_LspName_RsvpTe_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    priority.EntityData.Children = make(map[string]types.YChild)
    priority.EntityData.Leafs = make(map[string]types.YLeaf)
    priority.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", priority.SetupPriority}
    priority.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", priority.HoldPriority}
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

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
    logging.EntityData.Leafs["no-path"] = types.YLeaf{"NoPath", logging.NoPath}
    logging.EntityData.Leafs["pcerr"] = types.YLeaf{"Pcerr", logging.Pcerr}
    logging.EntityData.Leafs["fallback"] = types.YLeaf{"Fallback", logging.Fallback}
    return &(logging.EntityData)
}

// Pce_Backoff
// PCE backoff configuration
type Pce_Backoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    backoff.EntityData.Children = make(map[string]types.YChild)
    backoff.EntityData.Leafs = make(map[string]types.YLeaf)
    backoff.EntityData.Leafs["ratio"] = types.YLeaf{"Ratio", backoff.Ratio}
    backoff.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", backoff.Threshold}
    backoff.EntityData.Leafs["difference"] = types.YLeaf{"Difference", backoff.Difference}
    return &(backoff.EntityData)
}

// Pce_StateSyncs
// Standby PCE configuration
type Pce_StateSyncs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Standby PCE ipv4 address. The type is slice of Pce_StateSyncs_StateSync.
    StateSync []Pce_StateSyncs_StateSync
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

    stateSyncs.EntityData.Children = make(map[string]types.YChild)
    stateSyncs.EntityData.Children["state-sync"] = types.YChild{"StateSync", nil}
    for i := range stateSyncs.StateSync {
        stateSyncs.EntityData.Children[types.GetSegmentPath(&stateSyncs.StateSync[i])] = types.YChild{"StateSync", &stateSyncs.StateSync[i]}
    }
    stateSyncs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateSyncs.EntityData)
}

// Pce_StateSyncs_StateSync
// Standby PCE ipv4 address
type Pce_StateSyncs_StateSync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (stateSync *Pce_StateSyncs_StateSync) GetEntityData() *types.CommonEntityData {
    stateSync.EntityData.YFilter = stateSync.YFilter
    stateSync.EntityData.YangName = "state-sync"
    stateSync.EntityData.BundleName = "cisco_ios_xr"
    stateSync.EntityData.ParentYangName = "state-syncs"
    stateSync.EntityData.SegmentPath = "state-sync" + "[address='" + fmt.Sprintf("%v", stateSync.Address) + "']"
    stateSync.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSync.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSync.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSync.EntityData.Children = make(map[string]types.YChild)
    stateSync.EntityData.Leafs = make(map[string]types.YLeaf)
    stateSync.EntityData.Leafs["address"] = types.YLeaf{"Address", stateSync.Address}
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

    segmentRouting.EntityData.Children = make(map[string]types.YChild)
    segmentRouting.EntityData.Leafs = make(map[string]types.YLeaf)
    segmentRouting.EntityData.Leafs["te-latency"] = types.YLeaf{"TeLatency", segmentRouting.TeLatency}
    segmentRouting.EntityData.Leafs["strict-sid-only"] = types.YLeaf{"StrictSidOnly", segmentRouting.StrictSidOnly}
    return &(segmentRouting.EntityData)
}

// Pce_Timers
// PCE Timers configuration
type Pce_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
    timers.EntityData.Leafs["reoptimization-timer"] = types.YLeaf{"ReoptimizationTimer", timers.ReoptimizationTimer}
    timers.EntityData.Leafs["keepalive"] = types.YLeaf{"Keepalive", timers.Keepalive}
    timers.EntityData.Leafs["minimum-peer-keepalive"] = types.YLeaf{"MinimumPeerKeepalive", timers.MinimumPeerKeepalive}
    return &(timers.EntityData)
}

// Pce_Netconf
// NETCONF configuration
type Pce_Netconf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    netconf.EntityData.Children = make(map[string]types.YChild)
    netconf.EntityData.Children["netconf-ssh"] = types.YChild{"NetconfSsh", &netconf.NetconfSsh}
    netconf.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(netconf.EntityData)
}

// Pce_Netconf_NetconfSsh
// NETCONF SSH configuration
type Pce_Netconf_NetconfSsh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Password to use for NETCONF SSH connections. The type is string with
    // pattern: b'(!.+)|([^!].+)'.
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

    netconfSsh.EntityData.Children = make(map[string]types.YChild)
    netconfSsh.EntityData.Leafs = make(map[string]types.YLeaf)
    netconfSsh.EntityData.Leafs["netconf-ssh-password"] = types.YLeaf{"NetconfSshPassword", netconfSsh.NetconfSshPassword}
    netconfSsh.EntityData.Leafs["netconf-ssh-user"] = types.YLeaf{"NetconfSshUser", netconfSsh.NetconfSshUser}
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

    disjointPath.EntityData.Children = make(map[string]types.YChild)
    disjointPath.EntityData.Children["groups"] = types.YChild{"Groups", &disjointPath.Groups}
    disjointPath.EntityData.Leafs = make(map[string]types.YLeaf)
    disjointPath.EntityData.Leafs["enable"] = types.YLeaf{"Enable", disjointPath.Enable}
    return &(disjointPath.EntityData)
}

// Pce_DisjointPath_Groups
// Association configuration
type Pce_DisjointPath_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association Group Configuration. The type is slice of
    // Pce_DisjointPath_Groups_Group.
    Group []Pce_DisjointPath_Groups_Group
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

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
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
    group.EntityData.SegmentPath = "group" + "[group-id='" + fmt.Sprintf("%v", group.GroupId) + "']" + "[dp-type='" + fmt.Sprintf("%v", group.DpType) + "']" + "[sub-id='" + fmt.Sprintf("%v", group.SubId) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Children["group-lsp-records"] = types.YChild{"GroupLspRecords", &group.GroupLspRecords}
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", group.GroupId}
    group.EntityData.Leafs["dp-type"] = types.YLeaf{"DpType", group.DpType}
    group.EntityData.Leafs["sub-id"] = types.YLeaf{"SubId", group.SubId}
    group.EntityData.Leafs["strict"] = types.YLeaf{"Strict", group.Strict}
    group.EntityData.Leafs["enable"] = types.YLeaf{"Enable", group.Enable}
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
    GroupLspRecord []Pce_DisjointPath_Groups_Group_GroupLspRecords_GroupLspRecord
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

    groupLspRecords.EntityData.Children = make(map[string]types.YChild)
    groupLspRecords.EntityData.Children["group-lsp-record"] = types.YChild{"GroupLspRecord", nil}
    for i := range groupLspRecords.GroupLspRecord {
        groupLspRecords.EntityData.Children[types.GetSegmentPath(&groupLspRecords.GroupLspRecord[i])] = types.YChild{"GroupLspRecord", &groupLspRecords.GroupLspRecord[i]}
    }
    groupLspRecords.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpAddr interface{}

    // Identifying name for LSP. The type is string.
    LspName interface{}

    // Set LSP to follow shortest-path. The type is interface{} with range:
    // -2147483648..2147483647.
    DisjPath interface{}
}

func (groupLspRecord *Pce_DisjointPath_Groups_Group_GroupLspRecords_GroupLspRecord) GetEntityData() *types.CommonEntityData {
    groupLspRecord.EntityData.YFilter = groupLspRecord.YFilter
    groupLspRecord.EntityData.YangName = "group-lsp-record"
    groupLspRecord.EntityData.BundleName = "cisco_ios_xr"
    groupLspRecord.EntityData.ParentYangName = "group-lsp-records"
    groupLspRecord.EntityData.SegmentPath = "group-lsp-record" + "[lsp-id='" + fmt.Sprintf("%v", groupLspRecord.LspId) + "']"
    groupLspRecord.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupLspRecord.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupLspRecord.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupLspRecord.EntityData.Children = make(map[string]types.YChild)
    groupLspRecord.EntityData.Leafs = make(map[string]types.YLeaf)
    groupLspRecord.EntityData.Leafs["lsp-id"] = types.YLeaf{"LspId", groupLspRecord.LspId}
    groupLspRecord.EntityData.Leafs["ip-addr"] = types.YLeaf{"IpAddr", groupLspRecord.IpAddr}
    groupLspRecord.EntityData.Leafs["lsp-name"] = types.YLeaf{"LspName", groupLspRecord.LspName}
    groupLspRecord.EntityData.Leafs["disj-path"] = types.YLeaf{"DisjPath", groupLspRecord.DisjPath}
    return &(groupLspRecord.EntityData)
}

// Pce_ExplicitPaths
// Explicit paths
type Pce_ExplicitPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Explicit-path configuration. The type is slice of
    // Pce_ExplicitPaths_ExplicitPath.
    ExplicitPath []Pce_ExplicitPaths_ExplicitPath
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

    explicitPaths.EntityData.Children = make(map[string]types.YChild)
    explicitPaths.EntityData.Children["explicit-path"] = types.YChild{"ExplicitPath", nil}
    for i := range explicitPaths.ExplicitPath {
        explicitPaths.EntityData.Children[types.GetSegmentPath(&explicitPaths.ExplicitPath[i])] = types.YChild{"ExplicitPath", &explicitPaths.ExplicitPath[i]}
    }
    explicitPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(explicitPaths.EntityData)
}

// Pce_ExplicitPaths_ExplicitPath
// Explicit-path configuration
type Pce_ExplicitPaths_ExplicitPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Explicit-path name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    explicitPath.EntityData.SegmentPath = "explicit-path" + "[name='" + fmt.Sprintf("%v", explicitPath.Name) + "']"
    explicitPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitPath.EntityData.Children = make(map[string]types.YChild)
    explicitPath.EntityData.Children["path-hops"] = types.YChild{"PathHops", &explicitPath.PathHops}
    explicitPath.EntityData.Leafs = make(map[string]types.YLeaf)
    explicitPath.EntityData.Leafs["name"] = types.YLeaf{"Name", explicitPath.Name}
    explicitPath.EntityData.Leafs["enable"] = types.YLeaf{"Enable", explicitPath.Enable}
    return &(explicitPath.EntityData)
}

// Pce_ExplicitPaths_ExplicitPath_PathHops
// Path Hops
type Pce_ExplicitPaths_ExplicitPath_PathHops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Explicit path hop configuration. The type is slice of
    // Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop.
    PathHop []Pce_ExplicitPaths_ExplicitPath_PathHops_PathHop
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

    pathHops.EntityData.Children = make(map[string]types.YChild)
    pathHops.EntityData.Children["path-hop"] = types.YChild{"PathHop", nil}
    for i := range pathHops.PathHop {
        pathHops.EntityData.Children[types.GetSegmentPath(&pathHops.PathHop[i])] = types.YChild{"PathHop", &pathHops.PathHop[i]}
    }
    pathHops.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    // The default value is 0.0.0.0.
    Address interface{}

    // Remote IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    pathHop.EntityData.SegmentPath = "path-hop" + "[index='" + fmt.Sprintf("%v", pathHop.Index) + "']"
    pathHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathHop.EntityData.Children = make(map[string]types.YChild)
    pathHop.EntityData.Leafs = make(map[string]types.YLeaf)
    pathHop.EntityData.Leafs["index"] = types.YLeaf{"Index", pathHop.Index}
    pathHop.EntityData.Leafs["hop-type"] = types.YLeaf{"HopType", pathHop.HopType}
    pathHop.EntityData.Leafs["address"] = types.YLeaf{"Address", pathHop.Address}
    pathHop.EntityData.Leafs["remote-address"] = types.YLeaf{"RemoteAddress", pathHop.RemoteAddress}
    pathHop.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", pathHop.MplsLabel}
    return &(pathHop.EntityData)
}

