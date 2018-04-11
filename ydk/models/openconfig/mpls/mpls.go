// This module provides data definitions for configuration of
// Multiprotocol Label Switching (MPLS) and associated protocols for
// signaling and traffic engineering.
// 
// RFC 3031: Multiprotocol Label Switching Architecture
// 
// The MPLS / TE data model consists of several modules and
// submodules as shown below.  The top-level MPLS module describes
// the overall framework.  Three types of LSPs are supported:
// 
// i) traffic-engineered (or constrained-path)
// 
// ii) IGP-congruent (LSPs that follow the IGP path)
// 
// iii) static LSPs which are not signaled
// 
// The structure of each of these LSP configurations is defined in
// corresponding submodules.  Companion modules define the relevant
// configuration and operational data specific to key signaling
// protocols used in operational practice.
// 
// 
//                          +-------+
//        +---------------->| MPLS  |<--------------+
//        |                 +-------+               |
//        |                     ^                   |
//        |                     |                   |
//   +----+-----+      +--------+-------+     +-----+-----+
//   | TE LSPs  |      | IGP-based LSPs |     |static LSPs|
//   |          |      |                |     |           |
//   +----------+      +----------------+     +-----------+
//       ^  ^                    ^  ^
//       |  +----------------+   |  +--------+
//       |                   |   |           |
//       |   +------+      +-+---+-+      +--+--+
//       +---+ RSVP |      |SEGMENT|      | LDP |
//           +------+      |ROUTING|      +-----+
//                         +-------+
// 
package mpls

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls"))
    ydk.RegisterEntity("{http://openconfig.net/yang/mpls mpls}", reflect.TypeOf(Mpls{}))
    ydk.RegisterEntity("openconfig-mpls:mpls", reflect.TypeOf(Mpls{}))
}

type PathComputationMethod struct {
}

func (id PathComputationMethod) String() string {
	return "openconfig-mpls-te:path-computation-method"
}

type LocallyComputed struct {
}

func (id LocallyComputed) String() string {
	return "openconfig-mpls-te:locally-computed"
}

type ExternallyQueried struct {
}

func (id ExternallyQueried) String() string {
	return "openconfig-mpls-te:externally-queried"
}

type ExplicitlyDefined struct {
}

func (id ExplicitlyDefined) String() string {
	return "openconfig-mpls-te:explicitly-defined"
}

// TeBandwidthType represents explicitly specified or automatically computed
type TeBandwidthType string

const (
    // Bandwidth is explicitly specified
    TeBandwidthType_SPECIFIED TeBandwidthType = "SPECIFIED"

    // Bandwidth is automatically computed
    TeBandwidthType_AUTO TeBandwidthType = "AUTO"
)

// MplsSrlgFloodingType represents Enumerated bype for specifying how the SRLG is flooded
type MplsSrlgFloodingType string

const (
    // SRLG is flooded in the IGP
    MplsSrlgFloodingType_FLOODED_SRLG MplsSrlgFloodingType = "FLOODED-SRLG"

    // SRLG is not flooded, the members are
    // statically configured
    MplsSrlgFloodingType_STATIC_SRLG MplsSrlgFloodingType = "STATIC-SRLG"
)

// MplsHopType represents paths
type MplsHopType string

const (
    // loose hop in an explicit path
    MplsHopType_LOOSE MplsHopType = "LOOSE"

    // strict hop in an explicit path
    MplsHopType_STRICT MplsHopType = "STRICT"
)

// TeMetricType represents static value, or to track the IGP metric
type TeMetricType string

const (
    // set the LSP metric to track the underlying
    // IGP metric
    TeMetricType_IGP TeMetricType = "IGP"
)

// CspfTieBreaking represents multiple equal cost paths are available
type CspfTieBreaking string

const (
    // CSPF calculation selects a random path among
    // multiple equal-cost paths to the destination
    CspfTieBreaking_RANDOM CspfTieBreaking = "RANDOM"

    // CSPF calculation selects the path with greatest
    // available bandwidth
    CspfTieBreaking_LEAST_FILL CspfTieBreaking = "LEAST_FILL"

    // CSPF calculation selects the path with the least
    // available bandwidth
    CspfTieBreaking_MOST_FILL CspfTieBreaking = "MOST_FILL"
)

// Mpls
// Anchor point for mpls configuration and operational
// data
// This type is a presence type.
type Mpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // general mpls configuration applicable to any type of LSP and signaling
    // protocol - label ranges, entropy label supportmay be added here.
    Global Mpls_Global

    // traffic-engineering global attributes.
    TeGlobalAttributes Mpls_TeGlobalAttributes

    // traffic engineering attributes specific for interfaces.
    TeInterfaceAttributes Mpls_TeInterfaceAttributes

    // top-level signaling protocol configuration.
    SignalingProtocols Mpls_SignalingProtocols

    // LSP definitions and configuration.
    Lsps Mpls_Lsps
}

func (mpls *Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "openconfig"
    mpls.EntityData.ParentYangName = "openconfig-mpls"
    mpls.EntityData.SegmentPath = "openconfig-mpls:mpls"
    mpls.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    mpls.EntityData.NamespaceTable = openconfig.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    mpls.EntityData.Children = make(map[string]types.YChild)
    mpls.EntityData.Children["global"] = types.YChild{"Global", &mpls.Global}
    mpls.EntityData.Children["te-global-attributes"] = types.YChild{"TeGlobalAttributes", &mpls.TeGlobalAttributes}
    mpls.EntityData.Children["te-interface-attributes"] = types.YChild{"TeInterfaceAttributes", &mpls.TeInterfaceAttributes}
    mpls.EntityData.Children["signaling-protocols"] = types.YChild{"SignalingProtocols", &mpls.SignalingProtocols}
    mpls.EntityData.Children["lsps"] = types.YChild{"Lsps", &mpls.Lsps}
    mpls.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mpls.EntityData)
}

// Mpls_Global
// general mpls configuration applicable to any
// type of LSP and signaling protocol - label ranges,
// entropy label supportmay be added here
type Mpls_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Top level global MPLS configuration.
    Config Mpls_Global_Config

    // Top level global MPLS state.
    State Mpls_Global_State

    // Parameters related to MPLS interfaces.
    MplsInterfaceAttributes Mpls_Global_MplsInterfaceAttributes
}

func (global *Mpls_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "openconfig"
    global.EntityData.ParentYangName = "mpls"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    global.EntityData.NamespaceTable = openconfig.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["config"] = types.YChild{"Config", &global.Config}
    global.EntityData.Children["state"] = types.YChild{"State", &global.State}
    global.EntityData.Children["mpls-interface-attributes"] = types.YChild{"MplsInterfaceAttributes", &global.MplsInterfaceAttributes}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(global.EntityData)
}

// Mpls_Global_Config
// Top level global MPLS configuration
type Mpls_Global_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The null-label type used, implicit or explicit. The type is one of the
    // following: EXPLICITIMPLICIT. The default value is mplst:IMPLICIT.
    NullLabel interface{}
}

func (config *Mpls_Global_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "global"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["null-label"] = types.YLeaf{"NullLabel", config.NullLabel}
    return &(config.EntityData)
}

// Mpls_Global_State
// Top level global MPLS state
type Mpls_Global_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The null-label type used, implicit or explicit. The type is one of the
    // following: EXPLICITIMPLICIT. The default value is mplst:IMPLICIT.
    NullLabel interface{}
}

func (state *Mpls_Global_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "global"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["null-label"] = types.YLeaf{"NullLabel", state.NullLabel}
    return &(state.EntityData)
}

// Mpls_Global_MplsInterfaceAttributes
// Parameters related to MPLS interfaces
type Mpls_Global_MplsInterfaceAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of TE interfaces. The type is slice of
    // Mpls_Global_MplsInterfaceAttributes_Interface_.
    Interface_ []Mpls_Global_MplsInterfaceAttributes_Interface
}

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetEntityData() *types.CommonEntityData {
    mplsInterfaceAttributes.EntityData.YFilter = mplsInterfaceAttributes.YFilter
    mplsInterfaceAttributes.EntityData.YangName = "mpls-interface-attributes"
    mplsInterfaceAttributes.EntityData.BundleName = "openconfig"
    mplsInterfaceAttributes.EntityData.ParentYangName = "global"
    mplsInterfaceAttributes.EntityData.SegmentPath = "mpls-interface-attributes"
    mplsInterfaceAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    mplsInterfaceAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    mplsInterfaceAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    mplsInterfaceAttributes.EntityData.Children = make(map[string]types.YChild)
    mplsInterfaceAttributes.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range mplsInterfaceAttributes.Interface_ {
        mplsInterfaceAttributes.EntityData.Children[types.GetSegmentPath(&mplsInterfaceAttributes.Interface_[i])] = types.YChild{"Interface_", &mplsInterfaceAttributes.Interface_[i]}
    }
    mplsInterfaceAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsInterfaceAttributes.EntityData)
}

// Mpls_Global_MplsInterfaceAttributes_Interface
// List of TE interfaces
type Mpls_Global_MplsInterfaceAttributes_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string. Refers to
    // mpls.Mpls_Global_MplsInterfaceAttributes_Interface_Config_Name
    Name interface{}

    // Configuration parameters related to MPLS interfaces:.
    Config Mpls_Global_MplsInterfaceAttributes_Interface_Config

    // State parameters related to TE interfaces.
    State Mpls_Global_MplsInterfaceAttributes_Interface_State
}

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "mpls-interface-attributes"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["config"] = types.YChild{"Config", &self.Config}
    self.EntityData.Children["state"] = types.YChild{"State", &self.State}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    return &(self.EntityData)
}

// Mpls_Global_MplsInterfaceAttributes_Interface_Config
// Configuration parameters related to MPLS interfaces:
type Mpls_Global_MplsInterfaceAttributes_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reference to interface name. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // Enable MPLS forwarding on this interfacek. The type is bool. The default
    // value is false.
    MplsEnabled interface{}
}

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    config.EntityData.Leafs["mpls-enabled"] = types.YLeaf{"MplsEnabled", config.MplsEnabled}
    return &(config.EntityData)
}

// Mpls_Global_MplsInterfaceAttributes_Interface_State
// State parameters related to TE interfaces
type Mpls_Global_MplsInterfaceAttributes_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reference to interface name. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // Enable MPLS forwarding on this interfacek. The type is bool. The default
    // value is false.
    MplsEnabled interface{}
}

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    state.EntityData.Leafs["mpls-enabled"] = types.YLeaf{"MplsEnabled", state.MplsEnabled}
    return &(state.EntityData)
}

// Mpls_TeGlobalAttributes
// traffic-engineering global attributes
type Mpls_TeGlobalAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shared risk link groups attributes.
    Srlg Mpls_TeGlobalAttributes_Srlg

    // Interface bandwidth change percentages that trigger update events into the
    // IGP traffic engineering database (TED).
    IgpFloodingBandwidth Mpls_TeGlobalAttributes_IgpFloodingBandwidth

    // Top-level container for admin-groups configuration and state.
    MplsAdminGroups Mpls_TeGlobalAttributes_MplsAdminGroups

    // Definition for delays associated with setup and cleanup of TE LSPs.
    TeLspTimers Mpls_TeGlobalAttributes_TeLspTimers
}

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetEntityData() *types.CommonEntityData {
    teGlobalAttributes.EntityData.YFilter = teGlobalAttributes.YFilter
    teGlobalAttributes.EntityData.YangName = "te-global-attributes"
    teGlobalAttributes.EntityData.BundleName = "openconfig"
    teGlobalAttributes.EntityData.ParentYangName = "mpls"
    teGlobalAttributes.EntityData.SegmentPath = "te-global-attributes"
    teGlobalAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    teGlobalAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    teGlobalAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    teGlobalAttributes.EntityData.Children = make(map[string]types.YChild)
    teGlobalAttributes.EntityData.Children["srlg"] = types.YChild{"Srlg", &teGlobalAttributes.Srlg}
    teGlobalAttributes.EntityData.Children["igp-flooding-bandwidth"] = types.YChild{"IgpFloodingBandwidth", &teGlobalAttributes.IgpFloodingBandwidth}
    teGlobalAttributes.EntityData.Children["mpls-admin-groups"] = types.YChild{"MplsAdminGroups", &teGlobalAttributes.MplsAdminGroups}
    teGlobalAttributes.EntityData.Children["te-lsp-timers"] = types.YChild{"TeLspTimers", &teGlobalAttributes.TeLspTimers}
    teGlobalAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(teGlobalAttributes.EntityData)
}

// Mpls_TeGlobalAttributes_Srlg
// Shared risk link groups attributes
type Mpls_TeGlobalAttributes_Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of shared risk link groups. The type is slice of
    // Mpls_TeGlobalAttributes_Srlg_Srlg.
    Srlg []Mpls_TeGlobalAttributes_Srlg_Srlg_
}

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetEntityData() *types.CommonEntityData {
    srlg.EntityData.YFilter = srlg.YFilter
    srlg.EntityData.YangName = "srlg"
    srlg.EntityData.BundleName = "openconfig"
    srlg.EntityData.ParentYangName = "te-global-attributes"
    srlg.EntityData.SegmentPath = "srlg"
    srlg.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    srlg.EntityData.NamespaceTable = openconfig.GetNamespaces()
    srlg.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    srlg.EntityData.Children = make(map[string]types.YChild)
    srlg.EntityData.Children["srlg"] = types.YChild{"Srlg", nil}
    for i := range srlg.Srlg {
        srlg.EntityData.Children[types.GetSegmentPath(&srlg.Srlg[i])] = types.YChild{"Srlg", &srlg.Srlg[i]}
    }
    srlg.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(srlg.EntityData)
}

// Mpls_TeGlobalAttributes_Srlg_Srlg_
// List of shared risk link groups
type Mpls_TeGlobalAttributes_Srlg_Srlg_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The SRLG group identifier. The type is string.
    // Refers to mpls.Mpls_TeGlobalAttributes_Srlg_Srlg__Config_Name
    Name interface{}

    // Configuration parameters related to the SRLG.
    Config Mpls_TeGlobalAttributes_Srlg_Srlg__Config

    // State parameters related to the SRLG.
    State Mpls_TeGlobalAttributes_Srlg_Srlg__State

    // SRLG members for static (not flooded) SRLGs .
    StaticSrlgMembers Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers
}

func (srlg_ *Mpls_TeGlobalAttributes_Srlg_Srlg_) GetEntityData() *types.CommonEntityData {
    srlg_.EntityData.YFilter = srlg_.YFilter
    srlg_.EntityData.YangName = "srlg"
    srlg_.EntityData.BundleName = "openconfig"
    srlg_.EntityData.ParentYangName = "srlg"
    srlg_.EntityData.SegmentPath = "srlg" + "[name='" + fmt.Sprintf("%v", srlg_.Name) + "']"
    srlg_.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    srlg_.EntityData.NamespaceTable = openconfig.GetNamespaces()
    srlg_.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    srlg_.EntityData.Children = make(map[string]types.YChild)
    srlg_.EntityData.Children["config"] = types.YChild{"Config", &srlg_.Config}
    srlg_.EntityData.Children["state"] = types.YChild{"State", &srlg_.State}
    srlg_.EntityData.Children["static-srlg-members"] = types.YChild{"StaticSrlgMembers", &srlg_.StaticSrlgMembers}
    srlg_.EntityData.Leafs = make(map[string]types.YLeaf)
    srlg_.EntityData.Leafs["name"] = types.YLeaf{"Name", srlg_.Name}
    return &(srlg_.EntityData)
}

// Mpls_TeGlobalAttributes_Srlg_Srlg__Config
// Configuration parameters related to the SRLG
type Mpls_TeGlobalAttributes_Srlg_Srlg__Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG group identifier. The type is string.
    Name interface{}

    // group ID for the SRLG. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // The cost of the SRLG to the computation algorithm. The type is interface{}
    // with range: 0..4294967295.
    Cost interface{}

    // The type of SRLG, either flooded in the IGP or statically configured. The
    // type is MplsSrlgFloodingType. The default value is FLOODED-SRLG.
    FloodingType interface{}
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg__Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "srlg"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    config.EntityData.Leafs["value"] = types.YLeaf{"Value", config.Value}
    config.EntityData.Leafs["cost"] = types.YLeaf{"Cost", config.Cost}
    config.EntityData.Leafs["flooding-type"] = types.YLeaf{"FloodingType", config.FloodingType}
    return &(config.EntityData)
}

// Mpls_TeGlobalAttributes_Srlg_Srlg__State
// State parameters related to the SRLG
type Mpls_TeGlobalAttributes_Srlg_Srlg__State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG group identifier. The type is string.
    Name interface{}

    // group ID for the SRLG. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // The cost of the SRLG to the computation algorithm. The type is interface{}
    // with range: 0..4294967295.
    Cost interface{}

    // The type of SRLG, either flooded in the IGP or statically configured. The
    // type is MplsSrlgFloodingType. The default value is FLOODED-SRLG.
    FloodingType interface{}
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg__State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "srlg"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    state.EntityData.Leafs["value"] = types.YLeaf{"Value", state.Value}
    state.EntityData.Leafs["cost"] = types.YLeaf{"Cost", state.Cost}
    state.EntityData.Leafs["flooding-type"] = types.YLeaf{"FloodingType", state.FloodingType}
    return &(state.EntityData)
}

// Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers
// SRLG members for static (not flooded) SRLGs 
type Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of SRLG members, which are expressed as IP address endpoints of links
    // contained in the SRLG. The type is slice of
    // Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList.
    MembersList []Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList
}

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers) GetEntityData() *types.CommonEntityData {
    staticSrlgMembers.EntityData.YFilter = staticSrlgMembers.YFilter
    staticSrlgMembers.EntityData.YangName = "static-srlg-members"
    staticSrlgMembers.EntityData.BundleName = "openconfig"
    staticSrlgMembers.EntityData.ParentYangName = "srlg"
    staticSrlgMembers.EntityData.SegmentPath = "static-srlg-members"
    staticSrlgMembers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    staticSrlgMembers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    staticSrlgMembers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    staticSrlgMembers.EntityData.Children = make(map[string]types.YChild)
    staticSrlgMembers.EntityData.Children["members-list"] = types.YChild{"MembersList", nil}
    for i := range staticSrlgMembers.MembersList {
        staticSrlgMembers.EntityData.Children[types.GetSegmentPath(&staticSrlgMembers.MembersList[i])] = types.YChild{"MembersList", &staticSrlgMembers.MembersList[i]}
    }
    staticSrlgMembers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticSrlgMembers.EntityData)
}

// Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList
// List of SRLG members, which are expressed
// as IP address endpoints of links contained in the
// SRLG
type Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The from address of the link in the SRLG. The type
    // is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    FromAddress interface{}

    // Configuration parameters relating to the SRLG members.
    Config Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList_Config

    // State parameters relating to the SRLG members.
    State Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList_State
}

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList) GetEntityData() *types.CommonEntityData {
    membersList.EntityData.YFilter = membersList.YFilter
    membersList.EntityData.YangName = "members-list"
    membersList.EntityData.BundleName = "openconfig"
    membersList.EntityData.ParentYangName = "static-srlg-members"
    membersList.EntityData.SegmentPath = "members-list" + "[from-address='" + fmt.Sprintf("%v", membersList.FromAddress) + "']"
    membersList.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    membersList.EntityData.NamespaceTable = openconfig.GetNamespaces()
    membersList.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    membersList.EntityData.Children = make(map[string]types.YChild)
    membersList.EntityData.Children["config"] = types.YChild{"Config", &membersList.Config}
    membersList.EntityData.Children["state"] = types.YChild{"State", &membersList.State}
    membersList.EntityData.Leafs = make(map[string]types.YLeaf)
    membersList.EntityData.Leafs["from-address"] = types.YLeaf{"FromAddress", membersList.FromAddress}
    return &(membersList.EntityData)
}

// Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList_Config
// Configuration parameters relating to the
// SRLG members
type Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the a-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    FromAddress interface{}

    // IP address of the z-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ToAddress interface{}
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "members-list"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["from-address"] = types.YLeaf{"FromAddress", config.FromAddress}
    config.EntityData.Leafs["to-address"] = types.YLeaf{"ToAddress", config.ToAddress}
    return &(config.EntityData)
}

// Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList_State
// State parameters relating to the SRLG
// members
type Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the a-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    FromAddress interface{}

    // IP address of the z-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ToAddress interface{}
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg__StaticSrlgMembers_MembersList_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "members-list"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["from-address"] = types.YLeaf{"FromAddress", state.FromAddress}
    state.EntityData.Leafs["to-address"] = types.YLeaf{"ToAddress", state.ToAddress}
    return &(state.EntityData)
}

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth
// Interface bandwidth change percentages
// that trigger update events into the IGP traffic
// engineering database (TED)
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters for TED update threshold .
    Config Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config

    // State parameters for TED update threshold .
    State Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State
}

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetEntityData() *types.CommonEntityData {
    igpFloodingBandwidth.EntityData.YFilter = igpFloodingBandwidth.YFilter
    igpFloodingBandwidth.EntityData.YangName = "igp-flooding-bandwidth"
    igpFloodingBandwidth.EntityData.BundleName = "openconfig"
    igpFloodingBandwidth.EntityData.ParentYangName = "te-global-attributes"
    igpFloodingBandwidth.EntityData.SegmentPath = "igp-flooding-bandwidth"
    igpFloodingBandwidth.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    igpFloodingBandwidth.EntityData.NamespaceTable = openconfig.GetNamespaces()
    igpFloodingBandwidth.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    igpFloodingBandwidth.EntityData.Children = make(map[string]types.YChild)
    igpFloodingBandwidth.EntityData.Children["config"] = types.YChild{"Config", &igpFloodingBandwidth.Config}
    igpFloodingBandwidth.EntityData.Children["state"] = types.YChild{"State", &igpFloodingBandwidth.State}
    igpFloodingBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(igpFloodingBandwidth.EntityData)
}

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config
// Configuration parameters for TED
// update threshold 
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of threshold that should be used to specify the values at which
    // bandwidth is flooded. DELTA indicates that the local system should flood
    // IGP updates when a change in reserved bandwidth >= the specified delta
    // occurs on the interface. Where THRESHOLD-CROSSED is specified, the local
    // system should trigger an update (and hence flood) the reserved bandwidth
    // when the reserved bandwidth changes such that it crosses, or becomes equal
    // to one of the threshold values. The type is ThresholdType.
    ThresholdType interface{}

    // The percentage of the maximum-reservable-bandwidth considered as the delta
    // that results in an IGP update being flooded. The type is interface{} with
    // range: 0..100.
    DeltaPercentage interface{}

    // This value specifies whether a single set of threshold values should be
    // used for both increasing and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded in the IGP TE extensions.
    // MIRRORED-UP-DOWN indicates that a single value (or set of values) should be
    // used for both increasing and decreasing values, where SEPARATE-UP-DOWN
    // specifies that the increasing and decreasing values will be separately
    // specified. The type is ThresholdSpecification.
    ThresholdSpecification interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth) at which bandwidth updates are to be triggered when the
    // bandwidth is increasing. The type is slice of interface{} with range:
    // 0..100.
    UpThresholds []interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth) at which bandwidth updates are to be triggered when the
    // bandwidth is decreasing. The type is slice of interface{} with range:
    // 0..100.
    DownThresholds []interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth of the interface) at which bandwidth updates are flooded - used
    // both when the bandwidth is increasing and decreasing. The type is slice of
    // interface{} with range: 0..100.
    UpDownThresholds []interface{}
}

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "igp-flooding-bandwidth"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["threshold-type"] = types.YLeaf{"ThresholdType", config.ThresholdType}
    config.EntityData.Leafs["delta-percentage"] = types.YLeaf{"DeltaPercentage", config.DeltaPercentage}
    config.EntityData.Leafs["threshold-specification"] = types.YLeaf{"ThresholdSpecification", config.ThresholdSpecification}
    config.EntityData.Leafs["up-thresholds"] = types.YLeaf{"UpThresholds", config.UpThresholds}
    config.EntityData.Leafs["down-thresholds"] = types.YLeaf{"DownThresholds", config.DownThresholds}
    config.EntityData.Leafs["up-down-thresholds"] = types.YLeaf{"UpDownThresholds", config.UpDownThresholds}
    return &(config.EntityData)
}

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdSpecification represents and decreasing values will be separately specified
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdSpecification string

const (
    // MIRRORED-UP-DOWN indicates that a single set of
    // threshold values should be used for both increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdSpecification_MIRRORED_UP_DOWN Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdSpecification = "MIRRORED-UP-DOWN"

    // SEPARATE-UP-DOWN indicates that a separate
    // threshold values should be used for the increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdSpecification_SEPARATE_UP_DOWN Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdSpecification = "SEPARATE-UP-DOWN"
)

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdType represents values
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdType string

const (
    // DELTA indicates that the local
    // system should flood IGP updates when a
    // change in reserved bandwidth >= the specified
    // delta occurs on the interface.
    Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdType_DELTA Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdType = "DELTA"

    // THRESHOLD-CROSSED indicates that
    // the local system should trigger an update (and
    // hence flood) the reserved bandwidth when the
    // reserved bandwidth changes such that it crosses,
    // or becomes equal to one of the threshold values.
    Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdType_THRESHOLD_CROSSED Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config_ThresholdType = "THRESHOLD-CROSSED"
)

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State
// State parameters for TED update threshold 
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of threshold that should be used to specify the values at which
    // bandwidth is flooded. DELTA indicates that the local system should flood
    // IGP updates when a change in reserved bandwidth >= the specified delta
    // occurs on the interface. Where THRESHOLD-CROSSED is specified, the local
    // system should trigger an update (and hence flood) the reserved bandwidth
    // when the reserved bandwidth changes such that it crosses, or becomes equal
    // to one of the threshold values. The type is ThresholdType.
    ThresholdType interface{}

    // The percentage of the maximum-reservable-bandwidth considered as the delta
    // that results in an IGP update being flooded. The type is interface{} with
    // range: 0..100.
    DeltaPercentage interface{}

    // This value specifies whether a single set of threshold values should be
    // used for both increasing and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded in the IGP TE extensions.
    // MIRRORED-UP-DOWN indicates that a single value (or set of values) should be
    // used for both increasing and decreasing values, where SEPARATE-UP-DOWN
    // specifies that the increasing and decreasing values will be separately
    // specified. The type is ThresholdSpecification.
    ThresholdSpecification interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth) at which bandwidth updates are to be triggered when the
    // bandwidth is increasing. The type is slice of interface{} with range:
    // 0..100.
    UpThresholds []interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth) at which bandwidth updates are to be triggered when the
    // bandwidth is decreasing. The type is slice of interface{} with range:
    // 0..100.
    DownThresholds []interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth of the interface) at which bandwidth updates are flooded - used
    // both when the bandwidth is increasing and decreasing. The type is slice of
    // interface{} with range: 0..100.
    UpDownThresholds []interface{}
}

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "igp-flooding-bandwidth"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["threshold-type"] = types.YLeaf{"ThresholdType", state.ThresholdType}
    state.EntityData.Leafs["delta-percentage"] = types.YLeaf{"DeltaPercentage", state.DeltaPercentage}
    state.EntityData.Leafs["threshold-specification"] = types.YLeaf{"ThresholdSpecification", state.ThresholdSpecification}
    state.EntityData.Leafs["up-thresholds"] = types.YLeaf{"UpThresholds", state.UpThresholds}
    state.EntityData.Leafs["down-thresholds"] = types.YLeaf{"DownThresholds", state.DownThresholds}
    state.EntityData.Leafs["up-down-thresholds"] = types.YLeaf{"UpDownThresholds", state.UpDownThresholds}
    return &(state.EntityData)
}

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdSpecification represents and decreasing values will be separately specified
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdSpecification string

const (
    // MIRRORED-UP-DOWN indicates that a single set of
    // threshold values should be used for both increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdSpecification_MIRRORED_UP_DOWN Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdSpecification = "MIRRORED-UP-DOWN"

    // SEPARATE-UP-DOWN indicates that a separate
    // threshold values should be used for the increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdSpecification_SEPARATE_UP_DOWN Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdSpecification = "SEPARATE-UP-DOWN"
)

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdType represents values
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdType string

const (
    // DELTA indicates that the local
    // system should flood IGP updates when a
    // change in reserved bandwidth >= the specified
    // delta occurs on the interface.
    Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdType_DELTA Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdType = "DELTA"

    // THRESHOLD-CROSSED indicates that
    // the local system should trigger an update (and
    // hence flood) the reserved bandwidth when the
    // reserved bandwidth changes such that it crosses,
    // or becomes equal to one of the threshold values.
    Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdType_THRESHOLD_CROSSED Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State_ThresholdType = "THRESHOLD-CROSSED"
)

// Mpls_TeGlobalAttributes_MplsAdminGroups
// Top-level container for admin-groups configuration
// and state
type Mpls_TeGlobalAttributes_MplsAdminGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // configuration of value to name mapping for mpls affinities/admin-groups.
    // The type is slice of Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup.
    AdminGroup []Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup
}

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetEntityData() *types.CommonEntityData {
    mplsAdminGroups.EntityData.YFilter = mplsAdminGroups.YFilter
    mplsAdminGroups.EntityData.YangName = "mpls-admin-groups"
    mplsAdminGroups.EntityData.BundleName = "openconfig"
    mplsAdminGroups.EntityData.ParentYangName = "te-global-attributes"
    mplsAdminGroups.EntityData.SegmentPath = "mpls-admin-groups"
    mplsAdminGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    mplsAdminGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    mplsAdminGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    mplsAdminGroups.EntityData.Children = make(map[string]types.YChild)
    mplsAdminGroups.EntityData.Children["admin-group"] = types.YChild{"AdminGroup", nil}
    for i := range mplsAdminGroups.AdminGroup {
        mplsAdminGroups.EntityData.Children[types.GetSegmentPath(&mplsAdminGroups.AdminGroup[i])] = types.YChild{"AdminGroup", &mplsAdminGroups.AdminGroup[i]}
    }
    mplsAdminGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mplsAdminGroups.EntityData)
}

// Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup
// configuration of value to name mapping
// for mpls affinities/admin-groups
type Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. name for mpls admin-group. The type is string.
    // Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config_AdminGroupName
    AdminGroupName interface{}

    // Configurable items for admin-groups.
    Config Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config

    // Operational state for admin-groups.
    State Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State
}

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetEntityData() *types.CommonEntityData {
    adminGroup.EntityData.YFilter = adminGroup.YFilter
    adminGroup.EntityData.YangName = "admin-group"
    adminGroup.EntityData.BundleName = "openconfig"
    adminGroup.EntityData.ParentYangName = "mpls-admin-groups"
    adminGroup.EntityData.SegmentPath = "admin-group" + "[admin-group-name='" + fmt.Sprintf("%v", adminGroup.AdminGroupName) + "']"
    adminGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adminGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adminGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adminGroup.EntityData.Children = make(map[string]types.YChild)
    adminGroup.EntityData.Children["config"] = types.YChild{"Config", &adminGroup.Config}
    adminGroup.EntityData.Children["state"] = types.YChild{"State", &adminGroup.State}
    adminGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    adminGroup.EntityData.Leafs["admin-group-name"] = types.YLeaf{"AdminGroupName", adminGroup.AdminGroupName}
    return &(adminGroup.EntityData)
}

// Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config
// Configurable items for admin-groups
type Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name for mpls admin-group. The type is string.
    AdminGroupName interface{}

    // bit-position value for mpls admin-group. The value for the admin group is
    // an integer that represents one of the bit positions in the admin-group
    // bitmask. Values between 0 and 31 are interpreted as the original limit of
    // 32 admin groups. Values >=32 are interpreted as extended admin group values
    // as per RFC7308. The type is interface{} with range: 0..4294967295.
    BitPosition interface{}
}

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "admin-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["admin-group-name"] = types.YLeaf{"AdminGroupName", config.AdminGroupName}
    config.EntityData.Leafs["bit-position"] = types.YLeaf{"BitPosition", config.BitPosition}
    return &(config.EntityData)
}

// Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State
// Operational state for admin-groups
type Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name for mpls admin-group. The type is string.
    AdminGroupName interface{}

    // bit-position value for mpls admin-group. The value for the admin group is
    // an integer that represents one of the bit positions in the admin-group
    // bitmask. Values between 0 and 31 are interpreted as the original limit of
    // 32 admin groups. Values >=32 are interpreted as extended admin group values
    // as per RFC7308. The type is interface{} with range: 0..4294967295.
    BitPosition interface{}
}

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "admin-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["admin-group-name"] = types.YLeaf{"AdminGroupName", state.AdminGroupName}
    state.EntityData.Leafs["bit-position"] = types.YLeaf{"BitPosition", state.BitPosition}
    return &(state.EntityData)
}

// Mpls_TeGlobalAttributes_TeLspTimers
// Definition for delays associated with setup
// and cleanup of TE LSPs
type Mpls_TeGlobalAttributes_TeLspTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters related to timers for TE LSPs.
    Config Mpls_TeGlobalAttributes_TeLspTimers_Config

    // State related to timers for TE LSPs.
    State Mpls_TeGlobalAttributes_TeLspTimers_State
}

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetEntityData() *types.CommonEntityData {
    teLspTimers.EntityData.YFilter = teLspTimers.YFilter
    teLspTimers.EntityData.YangName = "te-lsp-timers"
    teLspTimers.EntityData.BundleName = "openconfig"
    teLspTimers.EntityData.ParentYangName = "te-global-attributes"
    teLspTimers.EntityData.SegmentPath = "te-lsp-timers"
    teLspTimers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    teLspTimers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    teLspTimers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    teLspTimers.EntityData.Children = make(map[string]types.YChild)
    teLspTimers.EntityData.Children["config"] = types.YChild{"Config", &teLspTimers.Config}
    teLspTimers.EntityData.Children["state"] = types.YChild{"State", &teLspTimers.State}
    teLspTimers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(teLspTimers.EntityData)
}

// Mpls_TeGlobalAttributes_TeLspTimers_Config
// Configuration parameters related
// to timers for TE LSPs
type Mpls_TeGlobalAttributes_TeLspTimers_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // delay the use of newly installed te lsp for a specified amount of time. The
    // type is interface{} with range: 0..3600. Units are seconds.
    InstallDelay interface{}

    // delay the removal of old te lsp for a specified amount of time. The type is
    // interface{} with range: 0..65535. Units are seconds.
    CleanupDelay interface{}

    // frequency of reoptimization of a traffic engineered LSP. The type is
    // interface{} with range: 0..65535. Units are seconds.
    ReoptimizeTimer interface{}
}

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "te-lsp-timers"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["install-delay"] = types.YLeaf{"InstallDelay", config.InstallDelay}
    config.EntityData.Leafs["cleanup-delay"] = types.YLeaf{"CleanupDelay", config.CleanupDelay}
    config.EntityData.Leafs["reoptimize-timer"] = types.YLeaf{"ReoptimizeTimer", config.ReoptimizeTimer}
    return &(config.EntityData)
}

// Mpls_TeGlobalAttributes_TeLspTimers_State
// State related to timers for TE LSPs
type Mpls_TeGlobalAttributes_TeLspTimers_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // delay the use of newly installed te lsp for a specified amount of time. The
    // type is interface{} with range: 0..3600. Units are seconds.
    InstallDelay interface{}

    // delay the removal of old te lsp for a specified amount of time. The type is
    // interface{} with range: 0..65535. Units are seconds.
    CleanupDelay interface{}

    // frequency of reoptimization of a traffic engineered LSP. The type is
    // interface{} with range: 0..65535. Units are seconds.
    ReoptimizeTimer interface{}
}

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "te-lsp-timers"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["install-delay"] = types.YLeaf{"InstallDelay", state.InstallDelay}
    state.EntityData.Leafs["cleanup-delay"] = types.YLeaf{"CleanupDelay", state.CleanupDelay}
    state.EntityData.Leafs["reoptimize-timer"] = types.YLeaf{"ReoptimizeTimer", state.ReoptimizeTimer}
    return &(state.EntityData)
}

// Mpls_TeInterfaceAttributes
// traffic engineering attributes specific
// for interfaces
type Mpls_TeInterfaceAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of TE interfaces. The type is slice of
    // Mpls_TeInterfaceAttributes_Interface_.
    Interface_ []Mpls_TeInterfaceAttributes_Interface
}

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetEntityData() *types.CommonEntityData {
    teInterfaceAttributes.EntityData.YFilter = teInterfaceAttributes.YFilter
    teInterfaceAttributes.EntityData.YangName = "te-interface-attributes"
    teInterfaceAttributes.EntityData.BundleName = "openconfig"
    teInterfaceAttributes.EntityData.ParentYangName = "mpls"
    teInterfaceAttributes.EntityData.SegmentPath = "te-interface-attributes"
    teInterfaceAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    teInterfaceAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    teInterfaceAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    teInterfaceAttributes.EntityData.Children = make(map[string]types.YChild)
    teInterfaceAttributes.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range teInterfaceAttributes.Interface_ {
        teInterfaceAttributes.EntityData.Children[types.GetSegmentPath(&teInterfaceAttributes.Interface_[i])] = types.YChild{"Interface_", &teInterfaceAttributes.Interface_[i]}
    }
    teInterfaceAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(teInterfaceAttributes.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface
// List of TE interfaces
type Mpls_TeInterfaceAttributes_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string. Refers to
    // mpls.Mpls_TeInterfaceAttributes_Interface_Config_Name
    Name interface{}

    // Configuration parameters related to TE interfaces:.
    Config Mpls_TeInterfaceAttributes_Interface_Config

    // State parameters related to TE interfaces.
    State Mpls_TeInterfaceAttributes_Interface_State

    // Interface bandwidth change percentages that trigger update events into the
    // IGP traffic engineering database (TED).
    IgpFloodingBandwidth Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth
}

func (self *Mpls_TeInterfaceAttributes_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "te-interface-attributes"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["config"] = types.YChild{"Config", &self.Config}
    self.EntityData.Children["state"] = types.YChild{"State", &self.State}
    self.EntityData.Children["igp-flooding-bandwidth"] = types.YChild{"IgpFloodingBandwidth", &self.IgpFloodingBandwidth}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    return &(self.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_Config
// Configuration parameters related to TE interfaces:
type Mpls_TeInterfaceAttributes_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reference to interface name. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // TE specific metric for the link. The type is interface{} with range:
    // 0..4294967295.
    TeMetric interface{}

    // list of references to named shared risk link groups that the interface
    // belongs to. The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_Srlg_Srlg__Name
    SrlgMembership []interface{}

    // list of admin groups (by name) on the interface. The type is slice of
    // string.
    AdminGroup []interface{}
}

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    config.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", config.TeMetric}
    config.EntityData.Leafs["srlg-membership"] = types.YLeaf{"SrlgMembership", config.SrlgMembership}
    config.EntityData.Leafs["admin-group"] = types.YLeaf{"AdminGroup", config.AdminGroup}
    return &(config.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_State
// State parameters related to TE interfaces
type Mpls_TeInterfaceAttributes_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reference to interface name. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // TE specific metric for the link. The type is interface{} with range:
    // 0..4294967295.
    TeMetric interface{}

    // list of references to named shared risk link groups that the interface
    // belongs to. The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_Srlg_Srlg__Name
    SrlgMembership []interface{}

    // list of admin groups (by name) on the interface. The type is slice of
    // string.
    AdminGroup []interface{}
}

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    state.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", state.TeMetric}
    state.EntityData.Leafs["srlg-membership"] = types.YLeaf{"SrlgMembership", state.SrlgMembership}
    state.EntityData.Leafs["admin-group"] = types.YLeaf{"AdminGroup", state.AdminGroup}
    return &(state.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth
// Interface bandwidth change percentages
// that trigger update events into the IGP traffic
// engineering database (TED)
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters for TED update threshold .
    Config Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config

    // State parameters for TED update threshold .
    State Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State
}

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetEntityData() *types.CommonEntityData {
    igpFloodingBandwidth.EntityData.YFilter = igpFloodingBandwidth.YFilter
    igpFloodingBandwidth.EntityData.YangName = "igp-flooding-bandwidth"
    igpFloodingBandwidth.EntityData.BundleName = "openconfig"
    igpFloodingBandwidth.EntityData.ParentYangName = "interface"
    igpFloodingBandwidth.EntityData.SegmentPath = "igp-flooding-bandwidth"
    igpFloodingBandwidth.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    igpFloodingBandwidth.EntityData.NamespaceTable = openconfig.GetNamespaces()
    igpFloodingBandwidth.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    igpFloodingBandwidth.EntityData.Children = make(map[string]types.YChild)
    igpFloodingBandwidth.EntityData.Children["config"] = types.YChild{"Config", &igpFloodingBandwidth.Config}
    igpFloodingBandwidth.EntityData.Children["state"] = types.YChild{"State", &igpFloodingBandwidth.State}
    igpFloodingBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(igpFloodingBandwidth.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config
// Configuration parameters for TED
// update threshold 
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of threshold that should be used to specify the values at which
    // bandwidth is flooded. DELTA indicates that the local system should flood
    // IGP updates when a change in reserved bandwidth >= the specified delta
    // occurs on the interface. Where THRESHOLD-CROSSED is specified, the local
    // system should trigger an update (and hence flood) the reserved bandwidth
    // when the reserved bandwidth changes such that it crosses, or becomes equal
    // to one of the threshold values. The type is ThresholdType.
    ThresholdType interface{}

    // The percentage of the maximum-reservable-bandwidth considered as the delta
    // that results in an IGP update being flooded. The type is interface{} with
    // range: 0..100.
    DeltaPercentage interface{}

    // This value specifies whether a single set of threshold values should be
    // used for both increasing and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded in the IGP TE extensions.
    // MIRRORED-UP-DOWN indicates that a single value (or set of values) should be
    // used for both increasing and decreasing values, where SEPARATE-UP-DOWN
    // specifies that the increasing and decreasing values will be separately
    // specified. The type is ThresholdSpecification.
    ThresholdSpecification interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth) at which bandwidth updates are to be triggered when the
    // bandwidth is increasing. The type is slice of interface{} with range:
    // 0..100.
    UpThresholds []interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth) at which bandwidth updates are to be triggered when the
    // bandwidth is decreasing. The type is slice of interface{} with range:
    // 0..100.
    DownThresholds []interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth of the interface) at which bandwidth updates are flooded - used
    // both when the bandwidth is increasing and decreasing. The type is slice of
    // interface{} with range: 0..100.
    UpDownThresholds []interface{}
}

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "igp-flooding-bandwidth"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["threshold-type"] = types.YLeaf{"ThresholdType", config.ThresholdType}
    config.EntityData.Leafs["delta-percentage"] = types.YLeaf{"DeltaPercentage", config.DeltaPercentage}
    config.EntityData.Leafs["threshold-specification"] = types.YLeaf{"ThresholdSpecification", config.ThresholdSpecification}
    config.EntityData.Leafs["up-thresholds"] = types.YLeaf{"UpThresholds", config.UpThresholds}
    config.EntityData.Leafs["down-thresholds"] = types.YLeaf{"DownThresholds", config.DownThresholds}
    config.EntityData.Leafs["up-down-thresholds"] = types.YLeaf{"UpDownThresholds", config.UpDownThresholds}
    return &(config.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification represents and decreasing values will be separately specified
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification string

const (
    // MIRRORED-UP-DOWN indicates that a single set of
    // threshold values should be used for both increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification_MIRRORED_UP_DOWN Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification = "MIRRORED-UP-DOWN"

    // SEPARATE-UP-DOWN indicates that a separate
    // threshold values should be used for the increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification_SEPARATE_UP_DOWN Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification = "SEPARATE-UP-DOWN"
)

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdType represents values
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdType string

const (
    // DELTA indicates that the local
    // system should flood IGP updates when a
    // change in reserved bandwidth >= the specified
    // delta occurs on the interface.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdType_DELTA Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdType = "DELTA"

    // THRESHOLD-CROSSED indicates that
    // the local system should trigger an update (and
    // hence flood) the reserved bandwidth when the
    // reserved bandwidth changes such that it crosses,
    // or becomes equal to one of the threshold values.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdType_THRESHOLD_CROSSED Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdType = "THRESHOLD-CROSSED"
)

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State
// State parameters for TED update threshold 
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of threshold that should be used to specify the values at which
    // bandwidth is flooded. DELTA indicates that the local system should flood
    // IGP updates when a change in reserved bandwidth >= the specified delta
    // occurs on the interface. Where THRESHOLD-CROSSED is specified, the local
    // system should trigger an update (and hence flood) the reserved bandwidth
    // when the reserved bandwidth changes such that it crosses, or becomes equal
    // to one of the threshold values. The type is ThresholdType.
    ThresholdType interface{}

    // The percentage of the maximum-reservable-bandwidth considered as the delta
    // that results in an IGP update being flooded. The type is interface{} with
    // range: 0..100.
    DeltaPercentage interface{}

    // This value specifies whether a single set of threshold values should be
    // used for both increasing and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded in the IGP TE extensions.
    // MIRRORED-UP-DOWN indicates that a single value (or set of values) should be
    // used for both increasing and decreasing values, where SEPARATE-UP-DOWN
    // specifies that the increasing and decreasing values will be separately
    // specified. The type is ThresholdSpecification.
    ThresholdSpecification interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth) at which bandwidth updates are to be triggered when the
    // bandwidth is increasing. The type is slice of interface{} with range:
    // 0..100.
    UpThresholds []interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth) at which bandwidth updates are to be triggered when the
    // bandwidth is decreasing. The type is slice of interface{} with range:
    // 0..100.
    DownThresholds []interface{}

    // The thresholds (expressed as a percentage of the maximum reservable
    // bandwidth of the interface) at which bandwidth updates are flooded - used
    // both when the bandwidth is increasing and decreasing. The type is slice of
    // interface{} with range: 0..100.
    UpDownThresholds []interface{}
}

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "igp-flooding-bandwidth"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["threshold-type"] = types.YLeaf{"ThresholdType", state.ThresholdType}
    state.EntityData.Leafs["delta-percentage"] = types.YLeaf{"DeltaPercentage", state.DeltaPercentage}
    state.EntityData.Leafs["threshold-specification"] = types.YLeaf{"ThresholdSpecification", state.ThresholdSpecification}
    state.EntityData.Leafs["up-thresholds"] = types.YLeaf{"UpThresholds", state.UpThresholds}
    state.EntityData.Leafs["down-thresholds"] = types.YLeaf{"DownThresholds", state.DownThresholds}
    state.EntityData.Leafs["up-down-thresholds"] = types.YLeaf{"UpDownThresholds", state.UpDownThresholds}
    return &(state.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification represents and decreasing values will be separately specified
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification string

const (
    // MIRRORED-UP-DOWN indicates that a single set of
    // threshold values should be used for both increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification_MIRRORED_UP_DOWN Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification = "MIRRORED-UP-DOWN"

    // SEPARATE-UP-DOWN indicates that a separate
    // threshold values should be used for the increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification_SEPARATE_UP_DOWN Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification = "SEPARATE-UP-DOWN"
)

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdType represents values
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdType string

const (
    // DELTA indicates that the local
    // system should flood IGP updates when a
    // change in reserved bandwidth >= the specified
    // delta occurs on the interface.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdType_DELTA Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdType = "DELTA"

    // THRESHOLD-CROSSED indicates that
    // the local system should trigger an update (and
    // hence flood) the reserved bandwidth when the
    // reserved bandwidth changes such that it crosses,
    // or becomes equal to one of the threshold values.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdType_THRESHOLD_CROSSED Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdType = "THRESHOLD-CROSSED"
)

// Mpls_SignalingProtocols
// top-level signaling protocol configuration
type Mpls_SignalingProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP-TE global signaling protocol configuration.
    RsvpTe Mpls_SignalingProtocols_RsvpTe

    // SR global signaling config.
    SegmentRouting Mpls_SignalingProtocols_SegmentRouting

    // LDP global signaling configuration.
    Ldp Mpls_SignalingProtocols_Ldp
}

func (signalingProtocols *Mpls_SignalingProtocols) GetEntityData() *types.CommonEntityData {
    signalingProtocols.EntityData.YFilter = signalingProtocols.YFilter
    signalingProtocols.EntityData.YangName = "signaling-protocols"
    signalingProtocols.EntityData.BundleName = "openconfig"
    signalingProtocols.EntityData.ParentYangName = "mpls"
    signalingProtocols.EntityData.SegmentPath = "signaling-protocols"
    signalingProtocols.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    signalingProtocols.EntityData.NamespaceTable = openconfig.GetNamespaces()
    signalingProtocols.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    signalingProtocols.EntityData.Children = make(map[string]types.YChild)
    signalingProtocols.EntityData.Children["rsvp-te"] = types.YChild{"RsvpTe", &signalingProtocols.RsvpTe}
    signalingProtocols.EntityData.Children["segment-routing"] = types.YChild{"SegmentRouting", &signalingProtocols.SegmentRouting}
    signalingProtocols.EntityData.Children["ldp"] = types.YChild{"Ldp", &signalingProtocols.Ldp}
    signalingProtocols.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(signalingProtocols.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe
// RSVP-TE global signaling protocol configuration
type Mpls_SignalingProtocols_RsvpTe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration and state of RSVP sessions.
    Sessions Mpls_SignalingProtocols_RsvpTe_Sessions

    // Configuration and state for RSVP neighbors connecting to the device.
    Neighbors Mpls_SignalingProtocols_RsvpTe_Neighbors

    // Platform wide RSVP configuration and state.
    Global Mpls_SignalingProtocols_RsvpTe_Global

    // Attributes relating to RSVP-TE enabled interfaces.
    InterfaceAttributes Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes
}

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetEntityData() *types.CommonEntityData {
    rsvpTe.EntityData.YFilter = rsvpTe.YFilter
    rsvpTe.EntityData.YangName = "rsvp-te"
    rsvpTe.EntityData.BundleName = "openconfig"
    rsvpTe.EntityData.ParentYangName = "signaling-protocols"
    rsvpTe.EntityData.SegmentPath = "rsvp-te"
    rsvpTe.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    rsvpTe.EntityData.NamespaceTable = openconfig.GetNamespaces()
    rsvpTe.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    rsvpTe.EntityData.Children = make(map[string]types.YChild)
    rsvpTe.EntityData.Children["sessions"] = types.YChild{"Sessions", &rsvpTe.Sessions}
    rsvpTe.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &rsvpTe.Neighbors}
    rsvpTe.EntityData.Children["global"] = types.YChild{"Global", &rsvpTe.Global}
    rsvpTe.EntityData.Children["interface-attributes"] = types.YChild{"InterfaceAttributes", &rsvpTe.InterfaceAttributes}
    rsvpTe.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rsvpTe.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions
// Configuration and state of RSVP sessions
type Mpls_SignalingProtocols_RsvpTe_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of RSVP sessions on the device.
    Config Mpls_SignalingProtocols_RsvpTe_Sessions_Config

    // State information relating to RSVP sessions on the device.
    State Mpls_SignalingProtocols_RsvpTe_Sessions_State
}

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "openconfig"
    sessions.EntityData.ParentYangName = "rsvp-te"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sessions.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["config"] = types.YChild{"Config", &sessions.Config}
    sessions.EntityData.Children["state"] = types.YChild{"State", &sessions.State}
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_Config
// Configuration of RSVP sessions on the device
type Mpls_SignalingProtocols_RsvpTe_Sessions_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "sessions"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_State
// State information relating to RSVP sessions
// on the device
type Mpls_SignalingProtocols_RsvpTe_Sessions_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of RSVP sessions. The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session.
    Session []Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "sessions"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["session"] = types.YChild{"Session", nil}
    for i := range state.Session {
        state.EntityData.Children[types.GetSegmentPath(&state.Session[i])] = types.YChild{"Session", &state.Session[i]}
    }
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session
// List of RSVP sessions
type Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RSVP source port. The type is interface{} with
    // range: 0..65535.
    SourcePort interface{}

    // This attribute is a key. RSVP source port. The type is interface{} with
    // range: 0..65535.
    DestinationPort interface{}

    // This attribute is a key. Origin address of RSVP session. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // This attribute is a key. Destination address of RSVP session. The type is
    // one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Enumeration of RSVP session states. The type is Status.
    Status interface{}

    // Enumeration of possible RSVP session types. The type is Type_.
    Type_ interface{}

    // Unique identifier of RSVP session. The type is interface{} with range:
    // 0..65535.
    TunnelId interface{}

    // Incoming MPLS label associated with this RSVP session. The type is one of
    // the following types: int with range: 16..1048575, or enumeration MplsLabel.
    LabelIn interface{}

    // Outgoing MPLS label associated with this RSVP session. The type is one of
    // the following types: int with range: 16..1048575, or enumeration MplsLabel.
    LabelOut interface{}

    // List of label switched paths associated with this RSVP session. The type is
    // slice of string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnel_Config_Name
    AssociatedLsps []interface{}
}

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "openconfig"
    session.EntityData.ParentYangName = "state"
    session.EntityData.SegmentPath = "session" + "[source-port='" + fmt.Sprintf("%v", session.SourcePort) + "']" + "[destination-port='" + fmt.Sprintf("%v", session.DestinationPort) + "']" + "[source-address='" + fmt.Sprintf("%v", session.SourceAddress) + "']" + "[destination-address='" + fmt.Sprintf("%v", session.DestinationAddress) + "']"
    session.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    session.EntityData.NamespaceTable = openconfig.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", session.SourcePort}
    session.EntityData.Leafs["destination-port"] = types.YLeaf{"DestinationPort", session.DestinationPort}
    session.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", session.SourceAddress}
    session.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", session.DestinationAddress}
    session.EntityData.Leafs["status"] = types.YLeaf{"Status", session.Status}
    session.EntityData.Leafs["type"] = types.YLeaf{"Type_", session.Type_}
    session.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", session.TunnelId}
    session.EntityData.Leafs["label-in"] = types.YLeaf{"LabelIn", session.LabelIn}
    session.EntityData.Leafs["label-out"] = types.YLeaf{"LabelOut", session.LabelOut}
    session.EntityData.Leafs["associated-lsps"] = types.YLeaf{"AssociatedLsps", session.AssociatedLsps}
    return &(session.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status represents Enumeration of RSVP session states
type Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status string

const (
    // RSVP session is up
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status_UP Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status = "UP"

    // RSVP session is down
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status_DOWN Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status = "DOWN"
)

// Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type_ represents Enumeration of possible RSVP session types
type Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type_ string

const (
    // RSVP session originates on this device
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type__SOURCE Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type_ = "SOURCE"

    // RSVP session transits this device only
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type__TRANSIT Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type_ = "TRANSIT"

    // RSVP session terminates on this device
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type__DESTINATION Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type_ = "DESTINATION"
)

// Mpls_SignalingProtocols_RsvpTe_Neighbors
// Configuration and state for RSVP neighbors connecting
// to the device
type Mpls_SignalingProtocols_RsvpTe_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of RSVP neighbor information.
    Config Mpls_SignalingProtocols_RsvpTe_Neighbors_Config

    // State information relating to RSVP neighbors.
    State Mpls_SignalingProtocols_RsvpTe_Neighbors_State
}

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "rsvp-te"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbors.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["config"] = types.YChild{"Config", &neighbors.Config}
    neighbors.EntityData.Children["state"] = types.YChild{"State", &neighbors.State}
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Neighbors_Config
// Configuration of RSVP neighbor information
type Mpls_SignalingProtocols_RsvpTe_Neighbors_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "neighbors"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Neighbors_State
// State information relating to RSVP neighbors
type Mpls_SignalingProtocols_RsvpTe_Neighbors_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of RSVP neighbors connecting to the device, keyed by neighbor address.
    // The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor.
    Neighbor []Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor
}

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "neighbors"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range state.Neighbor {
        state.EntityData.Children[types.GetSegmentPath(&state.Neighbor[i])] = types.YChild{"Neighbor", &state.Neighbor[i]}
    }
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor
// List of RSVP neighbors connecting to the device,
// keyed by neighbor address
type Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address of RSVP neighbor. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface where RSVP neighbor was detected. The type is string.
    DetectedInterface interface{}

    // Enumuration of possible RSVP neighbor states. The type is NeighborStatus.
    NeighborStatus interface{}

    // Suppport of neighbor for RSVP refresh reduction. The type is bool.
    RefreshReduction interface{}
}

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "state"
    neighbor.EntityData.SegmentPath = "neighbor" + "[address='" + fmt.Sprintf("%v", neighbor.Address) + "']"
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["address"] = types.YLeaf{"Address", neighbor.Address}
    neighbor.EntityData.Leafs["detected-interface"] = types.YLeaf{"DetectedInterface", neighbor.DetectedInterface}
    neighbor.EntityData.Leafs["neighbor-status"] = types.YLeaf{"NeighborStatus", neighbor.NeighborStatus}
    neighbor.EntityData.Leafs["refresh-reduction"] = types.YLeaf{"RefreshReduction", neighbor.RefreshReduction}
    return &(neighbor.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor_NeighborStatus represents Enumuration of possible RSVP neighbor states
type Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor_NeighborStatus string

const (
    // RSVP hello messages are detected from the neighbor
    Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor_NeighborStatus_UP Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor_NeighborStatus = "UP"

    // RSVP neighbor not detected as up, due to a
    // communication failure or IGP notification
    // the neighbor is unavailable
    Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor_NeighborStatus_DOWN Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor_NeighborStatus = "DOWN"
)

// Mpls_SignalingProtocols_RsvpTe_Global
// Platform wide RSVP configuration and state
type Mpls_SignalingProtocols_RsvpTe_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational state and configuration parameters relating to graceful-restart
    // for RSVP.
    GracefulRestart Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart

    // Protocol options relating to RSVP soft preemption.
    SoftPreemption Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption

    // Top level container for RSVP hello parameters.
    Hellos Mpls_SignalingProtocols_RsvpTe_Global_Hellos

    // Platform wide RSVP state, including counters.
    State Mpls_SignalingProtocols_RsvpTe_Global_State
}

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "openconfig"
    global.EntityData.ParentYangName = "rsvp-te"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    global.EntityData.NamespaceTable = openconfig.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["graceful-restart"] = types.YChild{"GracefulRestart", &global.GracefulRestart}
    global.EntityData.Children["soft-preemption"] = types.YChild{"SoftPreemption", &global.SoftPreemption}
    global.EntityData.Children["hellos"] = types.YChild{"Hellos", &global.Hellos}
    global.EntityData.Children["state"] = types.YChild{"State", &global.State}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(global.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart
// Operational state and configuration parameters relating to
// graceful-restart for RSVP
type Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to graceful-restart.
    Config Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config

    // State information associated with RSVP graceful-restart.
    State Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State
}

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "openconfig"
    gracefulRestart.EntityData.ParentYangName = "global"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = openconfig.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    gracefulRestart.EntityData.Children = make(map[string]types.YChild)
    gracefulRestart.EntityData.Children["config"] = types.YChild{"Config", &gracefulRestart.Config}
    gracefulRestart.EntityData.Children["state"] = types.YChild{"State", &gracefulRestart.State}
    gracefulRestart.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(gracefulRestart.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config
// Configuration parameters relating to
// graceful-restart
type Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables graceful restart on the node. The type is bool. The default value
    // is false.
    Enable interface{}

    // Graceful restart time (seconds). The type is interface{} with range:
    // 0..4294967295.
    RestartTime interface{}

    // RSVP state recovery time. The type is interface{} with range:
    // 0..4294967295.
    RecoveryTime interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "graceful-restart"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enable"] = types.YLeaf{"Enable", config.Enable}
    config.EntityData.Leafs["restart-time"] = types.YLeaf{"RestartTime", config.RestartTime}
    config.EntityData.Leafs["recovery-time"] = types.YLeaf{"RecoveryTime", config.RecoveryTime}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State
// State information associated with
// RSVP graceful-restart
type Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables graceful restart on the node. The type is bool. The default value
    // is false.
    Enable interface{}

    // Graceful restart time (seconds). The type is interface{} with range:
    // 0..4294967295.
    RestartTime interface{}

    // RSVP state recovery time. The type is interface{} with range:
    // 0..4294967295.
    RecoveryTime interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "graceful-restart"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enable"] = types.YLeaf{"Enable", state.Enable}
    state.EntityData.Leafs["restart-time"] = types.YLeaf{"RestartTime", state.RestartTime}
    state.EntityData.Leafs["recovery-time"] = types.YLeaf{"RecoveryTime", state.RecoveryTime}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption
// Protocol options relating to RSVP
// soft preemption
type Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to RSVP soft preemption support.
    Config Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config

    // State parameters relating to RSVP soft preemption support.
    State Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State
}

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetEntityData() *types.CommonEntityData {
    softPreemption.EntityData.YFilter = softPreemption.YFilter
    softPreemption.EntityData.YangName = "soft-preemption"
    softPreemption.EntityData.BundleName = "openconfig"
    softPreemption.EntityData.ParentYangName = "global"
    softPreemption.EntityData.SegmentPath = "soft-preemption"
    softPreemption.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    softPreemption.EntityData.NamespaceTable = openconfig.GetNamespaces()
    softPreemption.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    softPreemption.EntityData.Children = make(map[string]types.YChild)
    softPreemption.EntityData.Children["config"] = types.YChild{"Config", &softPreemption.Config}
    softPreemption.EntityData.Children["state"] = types.YChild{"State", &softPreemption.State}
    softPreemption.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(softPreemption.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config
// Configuration parameters relating to RSVP
// soft preemption support
type Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables soft preemption on a node. The type is bool. The default value is
    // false.
    Enable interface{}

    // Timeout value for soft preemption to revert to hard preemption. The type is
    // interface{} with range: 0..65535. The default value is 0.
    SoftPreemptionTimeout interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "soft-preemption"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enable"] = types.YLeaf{"Enable", config.Enable}
    config.EntityData.Leafs["soft-preemption-timeout"] = types.YLeaf{"SoftPreemptionTimeout", config.SoftPreemptionTimeout}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State
// State parameters relating to RSVP
// soft preemption support
type Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables soft preemption on a node. The type is bool. The default value is
    // false.
    Enable interface{}

    // Timeout value for soft preemption to revert to hard preemption. The type is
    // interface{} with range: 0..65535. The default value is 0.
    SoftPreemptionTimeout interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "soft-preemption"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enable"] = types.YLeaf{"Enable", state.Enable}
    state.EntityData.Leafs["soft-preemption-timeout"] = types.YLeaf{"SoftPreemptionTimeout", state.SoftPreemptionTimeout}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_Hellos
// Top level container for RSVP hello parameters
type Mpls_SignalingProtocols_RsvpTe_Global_Hellos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to RSVP hellos.
    Config Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config

    // State information associated with RSVP hellos.
    State Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetEntityData() *types.CommonEntityData {
    hellos.EntityData.YFilter = hellos.YFilter
    hellos.EntityData.YangName = "hellos"
    hellos.EntityData.BundleName = "openconfig"
    hellos.EntityData.ParentYangName = "global"
    hellos.EntityData.SegmentPath = "hellos"
    hellos.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    hellos.EntityData.NamespaceTable = openconfig.GetNamespaces()
    hellos.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    hellos.EntityData.Children = make(map[string]types.YChild)
    hellos.EntityData.Children["config"] = types.YChild{"Config", &hellos.Config}
    hellos.EntityData.Children["state"] = types.YChild{"State", &hellos.State}
    hellos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hellos.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config
// Configuration parameters relating to RSVP
// hellos
type Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // set the interval in ms between RSVP hello messages. The type is interface{}
    // with range: 1000..60000. Units are milliseconds. The default value is 9000.
    HelloInterval interface{}

    // enables all RSVP refresh reduction message bundling, RSVP message ID,
    // reliable message delivery and summary refresh. The type is bool. The
    // default value is true.
    RefreshReduction interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "hellos"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", config.HelloInterval}
    config.EntityData.Leafs["refresh-reduction"] = types.YLeaf{"RefreshReduction", config.RefreshReduction}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State
// State information associated with RSVP hellos
type Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // set the interval in ms between RSVP hello messages. The type is interface{}
    // with range: 1000..60000. Units are milliseconds. The default value is 9000.
    HelloInterval interface{}

    // enables all RSVP refresh reduction message bundling, RSVP message ID,
    // reliable message delivery and summary refresh. The type is bool. The
    // default value is true.
    RefreshReduction interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "hellos"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", state.HelloInterval}
    state.EntityData.Leafs["refresh-reduction"] = types.YLeaf{"RefreshReduction", state.RefreshReduction}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_State
// Platform wide RSVP state, including counters
type Mpls_SignalingProtocols_RsvpTe_Global_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Platform wide RSVP statistics and counters.
    Counters Mpls_SignalingProtocols_RsvpTe_Global_State_Counters
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "global"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["counters"] = types.YChild{"Counters", &state.Counters}
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Global_State_Counters
// Platform wide RSVP statistics and counters
type Mpls_SignalingProtocols_RsvpTe_Global_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TODO. The type is interface{} with range: 0..18446744073709551615.
    PathTimeouts interface{}

    // TODO. The type is interface{} with range: 0..18446744073709551615.
    ReservationTimeouts interface{}

    // RSVP messages dropped due to rate limiting. The type is interface{} with
    // range: 0..18446744073709551615.
    RateLimitedMessages interface{}

    // Number of received RSVP Path messages. The type is interface{} with range:
    // 0..18446744073709551615.
    InPathMessages interface{}

    // Number of received RSVP Path Error messages. The type is interface{} with
    // range: 0..18446744073709551615.
    InPathErrorMessages interface{}

    // Number of received RSVP Path Tear messages. The type is interface{} with
    // range: 0..18446744073709551615.
    InPathTearMessages interface{}

    // Number of received RSVP Resv messages. The type is interface{} with range:
    // 0..18446744073709551615.
    InReservationMessages interface{}

    // Number of received RSVP Resv Error messages. The type is interface{} with
    // range: 0..18446744073709551615.
    InReservationErrorMessages interface{}

    // Number of received RSVP Resv Tear messages. The type is interface{} with
    // range: 0..18446744073709551615.
    InReservationTearMessages interface{}

    // Number of received RSVP hello messages. The type is interface{} with range:
    // 0..18446744073709551615.
    InHelloMessages interface{}

    // Number of received RSVP summary refresh messages. The type is interface{}
    // with range: 0..18446744073709551615.
    InSrefreshMessages interface{}

    // Number of received RSVP refresh reduction ack messages. The type is
    // interface{} with range: 0..18446744073709551615.
    InAckMessages interface{}

    // Number of sent RSVP PATH messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPathMessages interface{}

    // Number of sent RSVP Path Error messages. The type is interface{} with
    // range: 0..18446744073709551615.
    OutPathErrorMessages interface{}

    // Number of sent RSVP Path Tear messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPathTearMessages interface{}

    // Number of sent RSVP Resv messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutReservationMessages interface{}

    // Number of sent RSVP Resv Error messages. The type is interface{} with
    // range: 0..18446744073709551615.
    OutReservationErrorMessages interface{}

    // Number of sent RSVP Resv Tear messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutReservationTearMessages interface{}

    // Number of sent RSVP hello messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutHelloMessages interface{}

    // Number of sent RSVP summary refresh messages. The type is interface{} with
    // range: 0..18446744073709551615.
    OutSrefreshMessages interface{}

    // Number of sent RSVP refresh reduction ack messages. The type is interface{}
    // with range: 0..18446744073709551615.
    OutAckMessages interface{}
}

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "openconfig"
    counters.EntityData.ParentYangName = "state"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    counters.EntityData.NamespaceTable = openconfig.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    counters.EntityData.Leafs["path-timeouts"] = types.YLeaf{"PathTimeouts", counters.PathTimeouts}
    counters.EntityData.Leafs["reservation-timeouts"] = types.YLeaf{"ReservationTimeouts", counters.ReservationTimeouts}
    counters.EntityData.Leafs["rate-limited-messages"] = types.YLeaf{"RateLimitedMessages", counters.RateLimitedMessages}
    counters.EntityData.Leafs["in-path-messages"] = types.YLeaf{"InPathMessages", counters.InPathMessages}
    counters.EntityData.Leafs["in-path-error-messages"] = types.YLeaf{"InPathErrorMessages", counters.InPathErrorMessages}
    counters.EntityData.Leafs["in-path-tear-messages"] = types.YLeaf{"InPathTearMessages", counters.InPathTearMessages}
    counters.EntityData.Leafs["in-reservation-messages"] = types.YLeaf{"InReservationMessages", counters.InReservationMessages}
    counters.EntityData.Leafs["in-reservation-error-messages"] = types.YLeaf{"InReservationErrorMessages", counters.InReservationErrorMessages}
    counters.EntityData.Leafs["in-reservation-tear-messages"] = types.YLeaf{"InReservationTearMessages", counters.InReservationTearMessages}
    counters.EntityData.Leafs["in-hello-messages"] = types.YLeaf{"InHelloMessages", counters.InHelloMessages}
    counters.EntityData.Leafs["in-srefresh-messages"] = types.YLeaf{"InSrefreshMessages", counters.InSrefreshMessages}
    counters.EntityData.Leafs["in-ack-messages"] = types.YLeaf{"InAckMessages", counters.InAckMessages}
    counters.EntityData.Leafs["out-path-messages"] = types.YLeaf{"OutPathMessages", counters.OutPathMessages}
    counters.EntityData.Leafs["out-path-error-messages"] = types.YLeaf{"OutPathErrorMessages", counters.OutPathErrorMessages}
    counters.EntityData.Leafs["out-path-tear-messages"] = types.YLeaf{"OutPathTearMessages", counters.OutPathTearMessages}
    counters.EntityData.Leafs["out-reservation-messages"] = types.YLeaf{"OutReservationMessages", counters.OutReservationMessages}
    counters.EntityData.Leafs["out-reservation-error-messages"] = types.YLeaf{"OutReservationErrorMessages", counters.OutReservationErrorMessages}
    counters.EntityData.Leafs["out-reservation-tear-messages"] = types.YLeaf{"OutReservationTearMessages", counters.OutReservationTearMessages}
    counters.EntityData.Leafs["out-hello-messages"] = types.YLeaf{"OutHelloMessages", counters.OutHelloMessages}
    counters.EntityData.Leafs["out-srefresh-messages"] = types.YLeaf{"OutSrefreshMessages", counters.OutSrefreshMessages}
    counters.EntityData.Leafs["out-ack-messages"] = types.YLeaf{"OutAckMessages", counters.OutAckMessages}
    return &(counters.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes
// Attributes relating to RSVP-TE enabled interfaces
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of per-interface RSVP configurations. The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_.
    Interface_ []Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface
}

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetEntityData() *types.CommonEntityData {
    interfaceAttributes.EntityData.YFilter = interfaceAttributes.YFilter
    interfaceAttributes.EntityData.YangName = "interface-attributes"
    interfaceAttributes.EntityData.BundleName = "openconfig"
    interfaceAttributes.EntityData.ParentYangName = "rsvp-te"
    interfaceAttributes.EntityData.SegmentPath = "interface-attributes"
    interfaceAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceAttributes.EntityData.Children = make(map[string]types.YChild)
    interfaceAttributes.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaceAttributes.Interface_ {
        interfaceAttributes.EntityData.Children[types.GetSegmentPath(&interfaceAttributes.Interface_[i])] = types.YChild{"Interface_", &interfaceAttributes.Interface_[i]}
    }
    interfaceAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceAttributes.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface
// list of per-interface RSVP configurations
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. references a configured IP interface. The type is
    // string. Refers to
    // mpls.Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config_InterfaceName
    InterfaceName interface{}

    // Configuration of per-interface RSVP parameters.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config

    // Per-interface RSVP protocol and state information.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State

    // Top level container for RSVP hello parameters.
    Hellos Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos

    // Configuration and state parameters relating to RSVP authentication as per
    // RFC2747.
    Authentication Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication

    // Bandwidth percentage reservable by RSVP on an interface.
    Subscription Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription

    // link-protection (NHOP) related configuration.
    Protection Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection
}

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "interface-attributes"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["config"] = types.YChild{"Config", &self.Config}
    self.EntityData.Children["state"] = types.YChild{"State", &self.State}
    self.EntityData.Children["hellos"] = types.YChild{"Hellos", &self.Hellos}
    self.EntityData.Children["authentication"] = types.YChild{"Authentication", &self.Authentication}
    self.EntityData.Children["subscription"] = types.YChild{"Subscription", &self.Subscription}
    self.EntityData.Children["protection"] = types.YChild{"Protection", &self.Protection}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config
// Configuration of per-interface RSVP parameters
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of configured IP interface. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    InterfaceName interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", config.InterfaceName}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State
// Per-interface RSVP protocol and state information
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum bandwidth ever reserved. The type is interface{} with range:
    // 0..18446744073709551615.
    HighwaterMark interface{}

    // Number of active RSVP reservations. The type is interface{} with range:
    // 0..18446744073709551615.
    ActiveReservationCount interface{}

    // Available and reserved bandwidth by priority on the interface. The type is
    // slice of
    // Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth.
    Bandwidth []Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth

    // Interface specific RSVP statistics and counters.
    Counters Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", nil}
    for i := range state.Bandwidth {
        state.EntityData.Children[types.GetSegmentPath(&state.Bandwidth[i])] = types.YChild{"Bandwidth", &state.Bandwidth[i]}
    }
    state.EntityData.Children["counters"] = types.YChild{"Counters", &state.Counters}
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["highwater-mark"] = types.YLeaf{"HighwaterMark", state.HighwaterMark}
    state.EntityData.Leafs["active-reservation-count"] = types.YLeaf{"ActiveReservationCount", state.ActiveReservationCount}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth
// Available and reserved bandwidth by priority on
// the interface.
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. RSVP priority level for LSPs traversing the
    // interface. The type is interface{} with range: 0..7.
    Priority interface{}

    // Bandwidth currently available. The type is interface{} with range:
    // 0..18446744073709551615.
    AvailableBandwidth interface{}

    // Bandwidth currently reserved. The type is interface{} with range:
    // 0..18446744073709551615.
    ReservedBandwidth interface{}
}

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "openconfig"
    bandwidth.EntityData.ParentYangName = "state"
    bandwidth.EntityData.SegmentPath = "bandwidth" + "[priority='" + fmt.Sprintf("%v", bandwidth.Priority) + "']"
    bandwidth.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    bandwidth.EntityData.Leafs["priority"] = types.YLeaf{"Priority", bandwidth.Priority}
    bandwidth.EntityData.Leafs["available-bandwidth"] = types.YLeaf{"AvailableBandwidth", bandwidth.AvailableBandwidth}
    bandwidth.EntityData.Leafs["reserved-bandwidth"] = types.YLeaf{"ReservedBandwidth", bandwidth.ReservedBandwidth}
    return &(bandwidth.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters
// Interface specific RSVP statistics and counters
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of received RSVP Path messages. The type is interface{} with range:
    // 0..18446744073709551615.
    InPathMessages interface{}

    // Number of received RSVP Path Error messages. The type is interface{} with
    // range: 0..18446744073709551615.
    InPathErrorMessages interface{}

    // Number of received RSVP Path Tear messages. The type is interface{} with
    // range: 0..18446744073709551615.
    InPathTearMessages interface{}

    // Number of received RSVP Resv messages. The type is interface{} with range:
    // 0..18446744073709551615.
    InReservationMessages interface{}

    // Number of received RSVP Resv Error messages. The type is interface{} with
    // range: 0..18446744073709551615.
    InReservationErrorMessages interface{}

    // Number of received RSVP Resv Tear messages. The type is interface{} with
    // range: 0..18446744073709551615.
    InReservationTearMessages interface{}

    // Number of received RSVP hello messages. The type is interface{} with range:
    // 0..18446744073709551615.
    InHelloMessages interface{}

    // Number of received RSVP summary refresh messages. The type is interface{}
    // with range: 0..18446744073709551615.
    InSrefreshMessages interface{}

    // Number of received RSVP refresh reduction ack messages. The type is
    // interface{} with range: 0..18446744073709551615.
    InAckMessages interface{}

    // Number of sent RSVP PATH messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPathMessages interface{}

    // Number of sent RSVP Path Error messages. The type is interface{} with
    // range: 0..18446744073709551615.
    OutPathErrorMessages interface{}

    // Number of sent RSVP Path Tear messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPathTearMessages interface{}

    // Number of sent RSVP Resv messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutReservationMessages interface{}

    // Number of sent RSVP Resv Error messages. The type is interface{} with
    // range: 0..18446744073709551615.
    OutReservationErrorMessages interface{}

    // Number of sent RSVP Resv Tear messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutReservationTearMessages interface{}

    // Number of sent RSVP hello messages. The type is interface{} with range:
    // 0..18446744073709551615.
    OutHelloMessages interface{}

    // Number of sent RSVP summary refresh messages. The type is interface{} with
    // range: 0..18446744073709551615.
    OutSrefreshMessages interface{}

    // Number of sent RSVP refresh reduction ack messages. The type is interface{}
    // with range: 0..18446744073709551615.
    OutAckMessages interface{}
}

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "openconfig"
    counters.EntityData.ParentYangName = "state"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    counters.EntityData.NamespaceTable = openconfig.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    counters.EntityData.Leafs["in-path-messages"] = types.YLeaf{"InPathMessages", counters.InPathMessages}
    counters.EntityData.Leafs["in-path-error-messages"] = types.YLeaf{"InPathErrorMessages", counters.InPathErrorMessages}
    counters.EntityData.Leafs["in-path-tear-messages"] = types.YLeaf{"InPathTearMessages", counters.InPathTearMessages}
    counters.EntityData.Leafs["in-reservation-messages"] = types.YLeaf{"InReservationMessages", counters.InReservationMessages}
    counters.EntityData.Leafs["in-reservation-error-messages"] = types.YLeaf{"InReservationErrorMessages", counters.InReservationErrorMessages}
    counters.EntityData.Leafs["in-reservation-tear-messages"] = types.YLeaf{"InReservationTearMessages", counters.InReservationTearMessages}
    counters.EntityData.Leafs["in-hello-messages"] = types.YLeaf{"InHelloMessages", counters.InHelloMessages}
    counters.EntityData.Leafs["in-srefresh-messages"] = types.YLeaf{"InSrefreshMessages", counters.InSrefreshMessages}
    counters.EntityData.Leafs["in-ack-messages"] = types.YLeaf{"InAckMessages", counters.InAckMessages}
    counters.EntityData.Leafs["out-path-messages"] = types.YLeaf{"OutPathMessages", counters.OutPathMessages}
    counters.EntityData.Leafs["out-path-error-messages"] = types.YLeaf{"OutPathErrorMessages", counters.OutPathErrorMessages}
    counters.EntityData.Leafs["out-path-tear-messages"] = types.YLeaf{"OutPathTearMessages", counters.OutPathTearMessages}
    counters.EntityData.Leafs["out-reservation-messages"] = types.YLeaf{"OutReservationMessages", counters.OutReservationMessages}
    counters.EntityData.Leafs["out-reservation-error-messages"] = types.YLeaf{"OutReservationErrorMessages", counters.OutReservationErrorMessages}
    counters.EntityData.Leafs["out-reservation-tear-messages"] = types.YLeaf{"OutReservationTearMessages", counters.OutReservationTearMessages}
    counters.EntityData.Leafs["out-hello-messages"] = types.YLeaf{"OutHelloMessages", counters.OutHelloMessages}
    counters.EntityData.Leafs["out-srefresh-messages"] = types.YLeaf{"OutSrefreshMessages", counters.OutSrefreshMessages}
    counters.EntityData.Leafs["out-ack-messages"] = types.YLeaf{"OutAckMessages", counters.OutAckMessages}
    return &(counters.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos
// Top level container for RSVP hello parameters
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to RSVP hellos.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config

    // State information associated with RSVP hellos.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetEntityData() *types.CommonEntityData {
    hellos.EntityData.YFilter = hellos.YFilter
    hellos.EntityData.YangName = "hellos"
    hellos.EntityData.BundleName = "openconfig"
    hellos.EntityData.ParentYangName = "interface"
    hellos.EntityData.SegmentPath = "hellos"
    hellos.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    hellos.EntityData.NamespaceTable = openconfig.GetNamespaces()
    hellos.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    hellos.EntityData.Children = make(map[string]types.YChild)
    hellos.EntityData.Children["config"] = types.YChild{"Config", &hellos.Config}
    hellos.EntityData.Children["state"] = types.YChild{"State", &hellos.State}
    hellos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hellos.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config
// Configuration parameters relating to RSVP
// hellos
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // set the interval in ms between RSVP hello messages. The type is interface{}
    // with range: 1000..60000. Units are milliseconds. The default value is 9000.
    HelloInterval interface{}

    // enables all RSVP refresh reduction message bundling, RSVP message ID,
    // reliable message delivery and summary refresh. The type is bool. The
    // default value is true.
    RefreshReduction interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "hellos"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", config.HelloInterval}
    config.EntityData.Leafs["refresh-reduction"] = types.YLeaf{"RefreshReduction", config.RefreshReduction}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State
// State information associated with RSVP hellos
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // set the interval in ms between RSVP hello messages. The type is interface{}
    // with range: 1000..60000. Units are milliseconds. The default value is 9000.
    HelloInterval interface{}

    // enables all RSVP refresh reduction message bundling, RSVP message ID,
    // reliable message delivery and summary refresh. The type is bool. The
    // default value is true.
    RefreshReduction interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "hellos"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", state.HelloInterval}
    state.EntityData.Leafs["refresh-reduction"] = types.YLeaf{"RefreshReduction", state.RefreshReduction}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication
// Configuration and state parameters relating to RSVP
// authentication as per RFC2747
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to authentication.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config

    // State information associated with authentication.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State
}

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "openconfig"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    authentication.EntityData.NamespaceTable = openconfig.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["config"] = types.YChild{"Config", &authentication.Config}
    authentication.EntityData.Children["state"] = types.YChild{"State", &authentication.State}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authentication.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config
// Configuration parameters relating
// to authentication
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables RSVP authentication on the node. The type is bool. The default
    // value is false.
    Enable interface{}

    // authenticate RSVP signaling messages. The type is string with length:
    // 1..32.
    AuthenticationKey interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "authentication"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enable"] = types.YLeaf{"Enable", config.Enable}
    config.EntityData.Leafs["authentication-key"] = types.YLeaf{"AuthenticationKey", config.AuthenticationKey}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State
// State information associated
// with authentication
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables RSVP authentication on the node. The type is bool. The default
    // value is false.
    Enable interface{}

    // authenticate RSVP signaling messages. The type is string with length:
    // 1..32.
    AuthenticationKey interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "authentication"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enable"] = types.YLeaf{"Enable", state.Enable}
    state.EntityData.Leafs["authentication-key"] = types.YLeaf{"AuthenticationKey", state.AuthenticationKey}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription
// Bandwidth percentage reservable by RSVP
// on an interface
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to RSVP subscription options.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config

    // State parameters relating to RSVP subscription options.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State
}

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetEntityData() *types.CommonEntityData {
    subscription.EntityData.YFilter = subscription.YFilter
    subscription.EntityData.YangName = "subscription"
    subscription.EntityData.BundleName = "openconfig"
    subscription.EntityData.ParentYangName = "interface"
    subscription.EntityData.SegmentPath = "subscription"
    subscription.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    subscription.EntityData.NamespaceTable = openconfig.GetNamespaces()
    subscription.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    subscription.EntityData.Children = make(map[string]types.YChild)
    subscription.EntityData.Children["config"] = types.YChild{"Config", &subscription.Config}
    subscription.EntityData.Children["state"] = types.YChild{"State", &subscription.State}
    subscription.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(subscription.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config
// Configuration parameters relating to RSVP
// subscription options
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // percentage of the interface bandwidth that RSVP can reserve. The type is
    // interface{} with range: 0..100.
    Subscription interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "subscription"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["subscription"] = types.YLeaf{"Subscription", config.Subscription}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State
// State parameters relating to RSVP
// subscription options
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // percentage of the interface bandwidth that RSVP can reserve. The type is
    // interface{} with range: 0..100.
    Subscription interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "subscription"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["subscription"] = types.YLeaf{"Subscription", state.Subscription}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection
// link-protection (NHOP) related configuration
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for link-protection.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config

    // State for link-protection.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State
}

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetEntityData() *types.CommonEntityData {
    protection.EntityData.YFilter = protection.YFilter
    protection.EntityData.YangName = "protection"
    protection.EntityData.BundleName = "openconfig"
    protection.EntityData.ParentYangName = "interface"
    protection.EntityData.SegmentPath = "protection"
    protection.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    protection.EntityData.NamespaceTable = openconfig.GetNamespaces()
    protection.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    protection.EntityData.Children = make(map[string]types.YChild)
    protection.EntityData.Children["config"] = types.YChild{"Config", &protection.Config}
    protection.EntityData.Children["state"] = types.YChild{"State", &protection.State}
    protection.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(protection.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config
// Configuration for link-protection
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Style of mpls frr protection desired: link, link-node, or unprotected. The
    // type is one of the following:
    // UnprotectedLinkProtectionRequestedLinkNodeProtectionRequested. The default
    // value is mplst:link-node-protection-requested.
    LinkProtectionStyleRequested interface{}

    // interval between periodic optimization of the bypass LSPs. The type is
    // interface{} with range: 0..65535. Units are seconds.
    BypassOptimizeInterval interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "protection"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["link-protection-style-requested"] = types.YLeaf{"LinkProtectionStyleRequested", config.LinkProtectionStyleRequested}
    config.EntityData.Leafs["bypass-optimize-interval"] = types.YLeaf{"BypassOptimizeInterval", config.BypassOptimizeInterval}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State
// State for link-protection
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Style of mpls frr protection desired: link, link-node, or unprotected. The
    // type is one of the following:
    // UnprotectedLinkProtectionRequestedLinkNodeProtectionRequested. The default
    // value is mplst:link-node-protection-requested.
    LinkProtectionStyleRequested interface{}

    // interval between periodic optimization of the bypass LSPs. The type is
    // interface{} with range: 0..65535. Units are seconds.
    BypassOptimizeInterval interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "protection"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["link-protection-style-requested"] = types.YLeaf{"LinkProtectionStyleRequested", state.LinkProtectionStyleRequested}
    state.EntityData.Leafs["bypass-optimize-interval"] = types.YLeaf{"BypassOptimizeInterval", state.BypassOptimizeInterval}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting
// SR global signaling config
type Mpls_SignalingProtocols_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Segment Routing Global Block (SRGB) entries. These label blocks are
    // reserved to be allocated as domain-wide entries. The type is slice of
    // Mpls_SignalingProtocols_SegmentRouting_Srgb.
    Srgb []Mpls_SignalingProtocols_SegmentRouting_Srgb

    // List of interfaces with associated segment routing configuration. The type
    // is slice of Mpls_SignalingProtocols_SegmentRouting_Interfaces.
    Interfaces []Mpls_SignalingProtocols_SegmentRouting_Interfaces
}

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetEntityData() *types.CommonEntityData {
    segmentRouting.EntityData.YFilter = segmentRouting.YFilter
    segmentRouting.EntityData.YangName = "segment-routing"
    segmentRouting.EntityData.BundleName = "openconfig"
    segmentRouting.EntityData.ParentYangName = "signaling-protocols"
    segmentRouting.EntityData.SegmentPath = "segment-routing"
    segmentRouting.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    segmentRouting.EntityData.NamespaceTable = openconfig.GetNamespaces()
    segmentRouting.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    segmentRouting.EntityData.Children = make(map[string]types.YChild)
    segmentRouting.EntityData.Children["srgb"] = types.YChild{"Srgb", nil}
    for i := range segmentRouting.Srgb {
        segmentRouting.EntityData.Children[types.GetSegmentPath(&segmentRouting.Srgb[i])] = types.YChild{"Srgb", &segmentRouting.Srgb[i]}
    }
    segmentRouting.EntityData.Children["interfaces"] = types.YChild{"Interfaces", nil}
    for i := range segmentRouting.Interfaces {
        segmentRouting.EntityData.Children[types.GetSegmentPath(&segmentRouting.Interfaces[i])] = types.YChild{"Interfaces", &segmentRouting.Interfaces[i]}
    }
    segmentRouting.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(segmentRouting.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Srgb
// List of Segment Routing Global Block (SRGB) entries. These
// label blocks are reserved to be allocated as domain-wide
// entries.
type Mpls_SignalingProtocols_SegmentRouting_Srgb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Lower value in the block. The type is interface{}
    // with range: 0..4294967295.
    LowerBound interface{}

    // This attribute is a key. Upper value in the block. The type is interface{}
    // with range: 0..4294967295.
    UpperBound interface{}

    // Configuration parameters relating to the Segment Routing Global Block
    // (SRGB).
    Config Mpls_SignalingProtocols_SegmentRouting_Srgb_Config

    // State parameters relating to the Segment Routing Global Block (SRGB).
    State Mpls_SignalingProtocols_SegmentRouting_Srgb_State
}

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetEntityData() *types.CommonEntityData {
    srgb.EntityData.YFilter = srgb.YFilter
    srgb.EntityData.YangName = "srgb"
    srgb.EntityData.BundleName = "openconfig"
    srgb.EntityData.ParentYangName = "segment-routing"
    srgb.EntityData.SegmentPath = "srgb" + "[lower-bound='" + fmt.Sprintf("%v", srgb.LowerBound) + "']" + "[upper-bound='" + fmt.Sprintf("%v", srgb.UpperBound) + "']"
    srgb.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    srgb.EntityData.NamespaceTable = openconfig.GetNamespaces()
    srgb.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    srgb.EntityData.Children = make(map[string]types.YChild)
    srgb.EntityData.Children["config"] = types.YChild{"Config", &srgb.Config}
    srgb.EntityData.Children["state"] = types.YChild{"State", &srgb.State}
    srgb.EntityData.Leafs = make(map[string]types.YLeaf)
    srgb.EntityData.Leafs["lower-bound"] = types.YLeaf{"LowerBound", srgb.LowerBound}
    srgb.EntityData.Leafs["upper-bound"] = types.YLeaf{"UpperBound", srgb.UpperBound}
    return &(srgb.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Srgb_Config
// Configuration parameters relating to the Segment Routing
// Global Block (SRGB)
type Mpls_SignalingProtocols_SegmentRouting_Srgb_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower value in the block. The type is interface{} with range:
    // 0..4294967295.
    LowerBound interface{}

    // Upper value in the block. The type is interface{} with range:
    // 0..4294967295.
    UpperBound interface{}
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "srgb"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["lower-bound"] = types.YLeaf{"LowerBound", config.LowerBound}
    config.EntityData.Leafs["upper-bound"] = types.YLeaf{"UpperBound", config.UpperBound}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Srgb_State
// State parameters relating to the Segment Routing Global
// Block (SRGB)
type Mpls_SignalingProtocols_SegmentRouting_Srgb_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Lower value in the block. The type is interface{} with range:
    // 0..4294967295.
    LowerBound interface{}

    // Upper value in the block. The type is interface{} with range:
    // 0..4294967295.
    UpperBound interface{}

    // Number of indexes in the SRGB block. The type is interface{} with range:
    // 0..4294967295.
    Size interface{}

    // Number of SRGB indexes that have not yet been allocated. The type is
    // interface{} with range: 0..4294967295.
    Free interface{}

    // Number of SRGB indexes that are currently allocated. The type is
    // interface{} with range: 0..4294967295.
    Used interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "srgb"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["lower-bound"] = types.YLeaf{"LowerBound", state.LowerBound}
    state.EntityData.Leafs["upper-bound"] = types.YLeaf{"UpperBound", state.UpperBound}
    state.EntityData.Leafs["size"] = types.YLeaf{"Size", state.Size}
    state.EntityData.Leafs["free"] = types.YLeaf{"Free", state.Free}
    state.EntityData.Leafs["used"] = types.YLeaf{"Used", state.Used}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces
// List of interfaces with associated segment routing
// configuration
type Mpls_SignalingProtocols_SegmentRouting_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the interface for which segment
    // routing configuration is to be applied. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Interface_ interface{}

    // Interface configuration parameters for Segment Routing relating to the
    // specified interface.
    Config Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config

    // State parameters for Segment Routing features relating to the specified
    // interface.
    State Mpls_SignalingProtocols_SegmentRouting_Interfaces_State

    // Configuration for Adjacency SIDs that are related to the specified
    // interface.
    AdjacencySid Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid
}

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "openconfig"
    interfaces.EntityData.ParentYangName = "segment-routing"
    interfaces.EntityData.SegmentPath = "interfaces" + "[interface='" + fmt.Sprintf("%v", interfaces.Interface_) + "']"
    interfaces.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaces.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["config"] = types.YChild{"Config", &interfaces.Config}
    interfaces.EntityData.Children["state"] = types.YChild{"State", &interfaces.State}
    interfaces.EntityData.Children["adjacency-sid"] = types.YChild{"AdjacencySid", &interfaces.AdjacencySid}
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaces.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", interfaces.Interface_}
    return &(interfaces.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config
// Interface configuration parameters for Segment Routing
// relating to the specified interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the interface for which segment routing configuration is to be
    // applied. The type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface_ interface{}
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interfaces"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", config.Interface_}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_State
// State parameters for Segment Routing features relating
// to the specified interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the interface for which segment routing configuration is to be
    // applied. The type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface_ interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interfaces"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", state.Interface_}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid
// Configuration for Adjacency SIDs that are related to
// the specified interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters for the Adjacency-SIDs that are related to this
    // interface.
    Config Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config

    // State parameters for the Adjacency-SIDs that are related to this interface.
    State Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State
}

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "openconfig"
    adjacencySid.EntityData.ParentYangName = "interfaces"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adjacencySid.EntityData.Children = make(map[string]types.YChild)
    adjacencySid.EntityData.Children["config"] = types.YChild{"Config", &adjacencySid.Config}
    adjacencySid.EntityData.Children["state"] = types.YChild{"State", &adjacencySid.State}
    adjacencySid.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adjacencySid.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config
// Configuration parameters for the Adjacency-SIDs
// that are related to this interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies the type of adjacency SID which should be advertised for the
    // specified entity. The type is slice of Advertise.
    Advertise []interface{}

    // Specifies the groups to which this interface belongs. Setting a value in
    // this list results in an additional AdjSID being advertised, with the S-bit
    // set to 1. The AdjSID is assumed to be protected. The type is slice of
    // interface{} with range: 0..4294967295.
    Groups []interface{}
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "adjacency-sid"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["advertise"] = types.YLeaf{"Advertise", config.Advertise}
    config.EntityData.Leafs["groups"] = types.YLeaf{"Groups", config.Groups}
    return &(config.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config_Advertise represents advertised for the specified entity.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config_Advertise string

const (
    // Advertise an Adjacency-SID for this interface, which is
    // eligible to be protected using a local protection
    // mechanism on the local LSR. The local protection
    // mechanism selected is dependent upon the configuration
    // of RSVP-TE FRR or LFA elsewhere on the system
    Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config_Advertise_PROTECTED Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config_Advertise = "PROTECTED"

    // Advertise an Adajcency-SID for this interface, which is
    // explicitly excluded from being protected by any local
    // protection mechanism
    Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config_Advertise_UNPROTECTED Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config_Advertise = "UNPROTECTED"
)

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State
// State parameters for the Adjacency-SIDs that are
// related to this interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies the type of adjacency SID which should be advertised for the
    // specified entity. The type is slice of Advertise.
    Advertise []interface{}

    // Specifies the groups to which this interface belongs. Setting a value in
    // this list results in an additional AdjSID being advertised, with the S-bit
    // set to 1. The AdjSID is assumed to be protected. The type is slice of
    // interface{} with range: 0..4294967295.
    Groups []interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "adjacency-sid"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["advertise"] = types.YLeaf{"Advertise", state.Advertise}
    state.EntityData.Leafs["groups"] = types.YLeaf{"Groups", state.Groups}
    return &(state.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State_Advertise represents advertised for the specified entity.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State_Advertise string

const (
    // Advertise an Adjacency-SID for this interface, which is
    // eligible to be protected using a local protection
    // mechanism on the local LSR. The local protection
    // mechanism selected is dependent upon the configuration
    // of RSVP-TE FRR or LFA elsewhere on the system
    Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State_Advertise_PROTECTED Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State_Advertise = "PROTECTED"

    // Advertise an Adajcency-SID for this interface, which is
    // explicitly excluded from being protected by any local
    // protection mechanism
    Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State_Advertise_UNPROTECTED Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State_Advertise = "UNPROTECTED"
)

// Mpls_SignalingProtocols_Ldp
// LDP global signaling configuration
type Mpls_SignalingProtocols_Ldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP timers.
    Timers Mpls_SignalingProtocols_Ldp_Timers
}

func (ldp *Mpls_SignalingProtocols_Ldp) GetEntityData() *types.CommonEntityData {
    ldp.EntityData.YFilter = ldp.YFilter
    ldp.EntityData.YangName = "ldp"
    ldp.EntityData.BundleName = "openconfig"
    ldp.EntityData.ParentYangName = "signaling-protocols"
    ldp.EntityData.SegmentPath = "ldp"
    ldp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ldp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ldp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ldp.EntityData.Children = make(map[string]types.YChild)
    ldp.EntityData.Children["timers"] = types.YChild{"Timers", &ldp.Timers}
    ldp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ldp.EntityData)
}

// Mpls_SignalingProtocols_Ldp_Timers
// LDP timers
type Mpls_SignalingProtocols_Ldp_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "openconfig"
    timers.EntityData.ParentYangName = "ldp"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    timers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timers.EntityData)
}

// Mpls_Lsps
// LSP definitions and configuration
type Mpls_Lsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // traffic-engineered LSPs supporting different path computation and signaling
    // methods.
    ConstrainedPath Mpls_Lsps_ConstrainedPath

    // LSPs that use the IGP-determined path, i.e., non traffic-engineered, or non
    // constrained-path.
    UnconstrainedPath Mpls_Lsps_UnconstrainedPath

    // statically configured LSPs, without dynamic signaling.
    StaticLsps Mpls_Lsps_StaticLsps
}

func (lsps *Mpls_Lsps) GetEntityData() *types.CommonEntityData {
    lsps.EntityData.YFilter = lsps.YFilter
    lsps.EntityData.YangName = "lsps"
    lsps.EntityData.BundleName = "openconfig"
    lsps.EntityData.ParentYangName = "mpls"
    lsps.EntityData.SegmentPath = "lsps"
    lsps.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    lsps.EntityData.NamespaceTable = openconfig.GetNamespaces()
    lsps.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    lsps.EntityData.Children = make(map[string]types.YChild)
    lsps.EntityData.Children["constrained-path"] = types.YChild{"ConstrainedPath", &lsps.ConstrainedPath}
    lsps.EntityData.Children["unconstrained-path"] = types.YChild{"UnconstrainedPath", &lsps.UnconstrainedPath}
    lsps.EntityData.Children["static-lsps"] = types.YChild{"StaticLsps", &lsps.StaticLsps}
    lsps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lsps.EntityData)
}

// Mpls_Lsps_ConstrainedPath
// traffic-engineered LSPs supporting different
// path computation and signaling methods
type Mpls_Lsps_ConstrainedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of explicit paths. The type is slice of
    // Mpls_Lsps_ConstrainedPath_NamedExplicitPaths.
    NamedExplicitPaths []Mpls_Lsps_ConstrainedPath_NamedExplicitPaths

    // List of TE tunnels. The type is slice of Mpls_Lsps_ConstrainedPath_Tunnel.
    Tunnel []Mpls_Lsps_ConstrainedPath_Tunnel
}

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetEntityData() *types.CommonEntityData {
    constrainedPath.EntityData.YFilter = constrainedPath.YFilter
    constrainedPath.EntityData.YangName = "constrained-path"
    constrainedPath.EntityData.BundleName = "openconfig"
    constrainedPath.EntityData.ParentYangName = "lsps"
    constrainedPath.EntityData.SegmentPath = "constrained-path"
    constrainedPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    constrainedPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    constrainedPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    constrainedPath.EntityData.Children = make(map[string]types.YChild)
    constrainedPath.EntityData.Children["named-explicit-paths"] = types.YChild{"NamedExplicitPaths", nil}
    for i := range constrainedPath.NamedExplicitPaths {
        constrainedPath.EntityData.Children[types.GetSegmentPath(&constrainedPath.NamedExplicitPaths[i])] = types.YChild{"NamedExplicitPaths", &constrainedPath.NamedExplicitPaths[i]}
    }
    constrainedPath.EntityData.Children["tunnel"] = types.YChild{"Tunnel", nil}
    for i := range constrainedPath.Tunnel {
        constrainedPath.EntityData.Children[types.GetSegmentPath(&constrainedPath.Tunnel[i])] = types.YChild{"Tunnel", &constrainedPath.Tunnel[i]}
    }
    constrainedPath.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(constrainedPath.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths
// A list of explicit paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A string name that uniquely identifies an explicit
    // path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config_Name
    Name interface{}

    // Configuration parameters relating to named explicit paths.
    Config Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config

    // Operational state parameters relating to the named explicit paths.
    State Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State

    // List of explicit route objects. The type is slice of
    // Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects.
    ExplicitRouteObjects []Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects
}

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetEntityData() *types.CommonEntityData {
    namedExplicitPaths.EntityData.YFilter = namedExplicitPaths.YFilter
    namedExplicitPaths.EntityData.YangName = "named-explicit-paths"
    namedExplicitPaths.EntityData.BundleName = "openconfig"
    namedExplicitPaths.EntityData.ParentYangName = "constrained-path"
    namedExplicitPaths.EntityData.SegmentPath = "named-explicit-paths" + "[name='" + fmt.Sprintf("%v", namedExplicitPaths.Name) + "']"
    namedExplicitPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    namedExplicitPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    namedExplicitPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    namedExplicitPaths.EntityData.Children = make(map[string]types.YChild)
    namedExplicitPaths.EntityData.Children["config"] = types.YChild{"Config", &namedExplicitPaths.Config}
    namedExplicitPaths.EntityData.Children["state"] = types.YChild{"State", &namedExplicitPaths.State}
    namedExplicitPaths.EntityData.Children["explicit-route-objects"] = types.YChild{"ExplicitRouteObjects", nil}
    for i := range namedExplicitPaths.ExplicitRouteObjects {
        namedExplicitPaths.EntityData.Children[types.GetSegmentPath(&namedExplicitPaths.ExplicitRouteObjects[i])] = types.YChild{"ExplicitRouteObjects", &namedExplicitPaths.ExplicitRouteObjects[i]}
    }
    namedExplicitPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    namedExplicitPaths.EntityData.Leafs["name"] = types.YLeaf{"Name", namedExplicitPaths.Name}
    return &(namedExplicitPaths.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config
// Configuration parameters relating to named explicit
// paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A string name that uniquely identifies an explicit path. The type is
    // string.
    Name interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "named-explicit-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State
// Operational state parameters relating to the named
// explicit paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A string name that uniquely identifies an explicit path. The type is
    // string.
    Name interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "named-explicit-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects
// List of explicit route objects
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of this explicit route object, to express
    // the order of hops in path. The type is string with range: 0..255. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config_Index
    Index interface{}

    // Configuration parameters relating to an explicit route.
    Config Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config

    // State parameters relating to an explicit route.
    State Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State
}

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetEntityData() *types.CommonEntityData {
    explicitRouteObjects.EntityData.YFilter = explicitRouteObjects.YFilter
    explicitRouteObjects.EntityData.YangName = "explicit-route-objects"
    explicitRouteObjects.EntityData.BundleName = "openconfig"
    explicitRouteObjects.EntityData.ParentYangName = "named-explicit-paths"
    explicitRouteObjects.EntityData.SegmentPath = "explicit-route-objects" + "[index='" + fmt.Sprintf("%v", explicitRouteObjects.Index) + "']"
    explicitRouteObjects.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    explicitRouteObjects.EntityData.NamespaceTable = openconfig.GetNamespaces()
    explicitRouteObjects.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    explicitRouteObjects.EntityData.Children = make(map[string]types.YChild)
    explicitRouteObjects.EntityData.Children["config"] = types.YChild{"Config", &explicitRouteObjects.Config}
    explicitRouteObjects.EntityData.Children["state"] = types.YChild{"State", &explicitRouteObjects.State}
    explicitRouteObjects.EntityData.Leafs = make(map[string]types.YLeaf)
    explicitRouteObjects.EntityData.Leafs["index"] = types.YLeaf{"Index", explicitRouteObjects.Index}
    return &(explicitRouteObjects.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config
// Configuration parameters relating to an explicit
// route
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // router hop for the LSP path. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // strict or loose hop. The type is MplsHopType.
    HopType interface{}

    // Index of this explicit route object to express the order of hops in the
    // path. The type is interface{} with range: 0..255.
    Index interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "explicit-route-objects"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["address"] = types.YLeaf{"Address", config.Address}
    config.EntityData.Leafs["hop-type"] = types.YLeaf{"HopType", config.HopType}
    config.EntityData.Leafs["index"] = types.YLeaf{"Index", config.Index}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State
// State parameters relating to an explicit route
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // router hop for the LSP path. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // strict or loose hop. The type is MplsHopType.
    HopType interface{}

    // Index of this explicit route object to express the order of hops in the
    // path. The type is interface{} with range: 0..255.
    Index interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "explicit-route-objects"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["address"] = types.YLeaf{"Address", state.Address}
    state.EntityData.Leafs["hop-type"] = types.YLeaf{"HopType", state.HopType}
    state.EntityData.Leafs["index"] = types.YLeaf{"Index", state.Index}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel
// List of TE tunnels
type Mpls_Lsps_ConstrainedPath_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The tunnel name. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnel_Config_Name
    Name interface{}

    // This attribute is a key. The tunnel type, p2p or p2mp. The type is one of
    // the following: P2PP2MP.
    Type_ interface{}

    // Configuration parameters related to TE tunnels:.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Config

    // State parameters related to TE tunnels.
    State Mpls_Lsps_ConstrainedPath_Tunnel_State

    // Bandwidth configuration for TE LSPs.
    Bandwidth Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth

    // Parameters related to LSPs of type P2P.
    P2PTunnelAttributes Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes
}

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "openconfig"
    tunnel.EntityData.ParentYangName = "constrained-path"
    tunnel.EntityData.SegmentPath = "tunnel" + "[name='" + fmt.Sprintf("%v", tunnel.Name) + "']" + "[type='" + fmt.Sprintf("%v", tunnel.Type_) + "']"
    tunnel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    tunnel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Children["config"] = types.YChild{"Config", &tunnel.Config}
    tunnel.EntityData.Children["state"] = types.YChild{"State", &tunnel.State}
    tunnel.EntityData.Children["bandwidth"] = types.YChild{"Bandwidth", &tunnel.Bandwidth}
    tunnel.EntityData.Children["p2p-tunnel-attributes"] = types.YChild{"P2PTunnelAttributes", &tunnel.P2PTunnelAttributes}
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["name"] = types.YLeaf{"Name", tunnel.Name}
    tunnel.EntityData.Leafs["type"] = types.YLeaf{"Type_", tunnel.Type_}
    return &(tunnel.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Config
// Configuration parameters related to TE tunnels:
type Mpls_Lsps_ConstrainedPath_Tunnel_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The tunnel name. The type is string.
    Name interface{}

    // Tunnel type, p2p or p2mp. The type is one of the following: P2PP2MP.
    Type_ interface{}

    // Signaling protocol used to set up this tunnel. The type is one of the
    // following: P2PP2MP.
    SignalingProtocol interface{}

    // locally signficant optional identifier for the tunnel; may be a numerical
    // or string value. The type is one of the following types: int with range:
    // 0..4294967295, or string.
    LocalId interface{}

    // optional text description for the tunnel. The type is string.
    Description interface{}

    // TE tunnel administrative state. The type is one of the following:
    // ADMINDOWNADMINUP. The default value is mplst:ADMIN_UP.
    AdminStatus interface{}

    // Specifies a preference for this tunnel. A lower number signifies a better
    // preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // LSP metric, either explicit or IGP. The type is one of the following types:
    // enumeration TeMetricType, or int with range: 0..4294967295.
    Metric interface{}

    // style of mpls frr protection desired: can be link, link-node or
    // unprotected. The type is one of the following:
    // UnprotectedLinkProtectionRequestedLinkNodeProtectionRequested. The default
    // value is mplst:unprotected.
    ProtectionStyleRequested interface{}

    // frequency of reoptimization of a traffic engineered LSP. The type is
    // interface{} with range: 0..65535. Units are seconds.
    ReoptimizeTimer interface{}

    // RSVP-TE tunnel source address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Source interface{}

    // Enables RSVP soft-preemption on this LSP. The type is bool. The default
    // value is false.
    SoftPreemption interface{}

    // RSVP-TE preemption priority during LSP setup, lower is higher priority;
    // default 7 indicates that LSP will not preempt established LSPs during
    // setup. The type is interface{} with range: 0..7. The default value is 7.
    SetupPriority interface{}

    // preemption priority once the LSP is established, lower is higher priority;
    // default 0 indicates other LSPs will not preempt the LSPs once established.
    // The type is interface{} with range: 0..7. The default value is 0.
    HoldPriority interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "tunnel"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    config.EntityData.Leafs["type"] = types.YLeaf{"Type_", config.Type_}
    config.EntityData.Leafs["signaling-protocol"] = types.YLeaf{"SignalingProtocol", config.SignalingProtocol}
    config.EntityData.Leafs["local-id"] = types.YLeaf{"LocalId", config.LocalId}
    config.EntityData.Leafs["description"] = types.YLeaf{"Description", config.Description}
    config.EntityData.Leafs["admin-status"] = types.YLeaf{"AdminStatus", config.AdminStatus}
    config.EntityData.Leafs["preference"] = types.YLeaf{"Preference", config.Preference}
    config.EntityData.Leafs["metric"] = types.YLeaf{"Metric", config.Metric}
    config.EntityData.Leafs["protection-style-requested"] = types.YLeaf{"ProtectionStyleRequested", config.ProtectionStyleRequested}
    config.EntityData.Leafs["reoptimize-timer"] = types.YLeaf{"ReoptimizeTimer", config.ReoptimizeTimer}
    config.EntityData.Leafs["source"] = types.YLeaf{"Source", config.Source}
    config.EntityData.Leafs["soft-preemption"] = types.YLeaf{"SoftPreemption", config.SoftPreemption}
    config.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", config.SetupPriority}
    config.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", config.HoldPriority}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_State
// State parameters related to TE tunnels
type Mpls_Lsps_ConstrainedPath_Tunnel_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The tunnel name. The type is string.
    Name interface{}

    // Tunnel type, p2p or p2mp. The type is one of the following: P2PP2MP.
    Type_ interface{}

    // Signaling protocol used to set up this tunnel. The type is one of the
    // following: P2PP2MP.
    SignalingProtocol interface{}

    // locally signficant optional identifier for the tunnel; may be a numerical
    // or string value. The type is one of the following types: int with range:
    // 0..4294967295, or string.
    LocalId interface{}

    // optional text description for the tunnel. The type is string.
    Description interface{}

    // TE tunnel administrative state. The type is one of the following:
    // ADMINDOWNADMINUP. The default value is mplst:ADMIN_UP.
    AdminStatus interface{}

    // Specifies a preference for this tunnel. A lower number signifies a better
    // preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // LSP metric, either explicit or IGP. The type is one of the following types:
    // enumeration TeMetricType, or int with range: 0..4294967295.
    Metric interface{}

    // style of mpls frr protection desired: can be link, link-node or
    // unprotected. The type is one of the following:
    // UnprotectedLinkProtectionRequestedLinkNodeProtectionRequested. The default
    // value is mplst:unprotected.
    ProtectionStyleRequested interface{}

    // frequency of reoptimization of a traffic engineered LSP. The type is
    // interface{} with range: 0..65535. Units are seconds.
    ReoptimizeTimer interface{}

    // RSVP-TE tunnel source address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Source interface{}

    // Enables RSVP soft-preemption on this LSP. The type is bool. The default
    // value is false.
    SoftPreemption interface{}

    // RSVP-TE preemption priority during LSP setup, lower is higher priority;
    // default 7 indicates that LSP will not preempt established LSPs during
    // setup. The type is interface{} with range: 0..7. The default value is 7.
    SetupPriority interface{}

    // preemption priority once the LSP is established, lower is higher priority;
    // default 0 indicates other LSPs will not preempt the LSPs once established.
    // The type is interface{} with range: 0..7. The default value is 0.
    HoldPriority interface{}

    // The operational status of the TE tunnel. The type is one of the following:
    // DOWNUP.
    OperStatus interface{}

    // The lsp role at the current node, whether it is headend, transit or
    // tailend. The type is one of the following: INGRESSEGRESSTRANSIT.
    Role interface{}

    // State data for MPLS label switched paths. This state data is specific to a
    // single label switched path.
    Counters Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "tunnel"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Children["counters"] = types.YChild{"Counters", &state.Counters}
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    state.EntityData.Leafs["type"] = types.YLeaf{"Type_", state.Type_}
    state.EntityData.Leafs["signaling-protocol"] = types.YLeaf{"SignalingProtocol", state.SignalingProtocol}
    state.EntityData.Leafs["local-id"] = types.YLeaf{"LocalId", state.LocalId}
    state.EntityData.Leafs["description"] = types.YLeaf{"Description", state.Description}
    state.EntityData.Leafs["admin-status"] = types.YLeaf{"AdminStatus", state.AdminStatus}
    state.EntityData.Leafs["preference"] = types.YLeaf{"Preference", state.Preference}
    state.EntityData.Leafs["metric"] = types.YLeaf{"Metric", state.Metric}
    state.EntityData.Leafs["protection-style-requested"] = types.YLeaf{"ProtectionStyleRequested", state.ProtectionStyleRequested}
    state.EntityData.Leafs["reoptimize-timer"] = types.YLeaf{"ReoptimizeTimer", state.ReoptimizeTimer}
    state.EntityData.Leafs["source"] = types.YLeaf{"Source", state.Source}
    state.EntityData.Leafs["soft-preemption"] = types.YLeaf{"SoftPreemption", state.SoftPreemption}
    state.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", state.SetupPriority}
    state.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", state.HoldPriority}
    state.EntityData.Leafs["oper-status"] = types.YLeaf{"OperStatus", state.OperStatus}
    state.EntityData.Leafs["role"] = types.YLeaf{"Role", state.Role}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters
// State data for MPLS label switched paths. This state
// data is specific to a single label switched path.
type Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of bytes that have been forwarded over the label switched path. The
    // type is interface{} with range: 0..18446744073709551615.
    Bytes interface{}

    // Number of pacets that have been forwarded over the label switched path. The
    // type is interface{} with range: 0..18446744073709551615.
    Packets interface{}

    // Number of path changes for the label switched path. The type is interface{}
    // with range: 0..18446744073709551615.
    PathChanges interface{}

    // Number of state changes for the label switched path. The type is
    // interface{} with range: 0..18446744073709551615.
    StateChanges interface{}

    // Indication of the time the label switched path transitioned to an Oper Up
    // or in-service state. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    OnlineTime interface{}

    // Indicates the time the LSP switched onto its current path. This is reset
    // upon a LSP path change. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    CurrentPathTime interface{}

    // Indicates the next scheduled time the LSP will be reoptimized. The type is
    // string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    NextReoptimizationTime interface{}
}

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "openconfig"
    counters.EntityData.ParentYangName = "state"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    counters.EntityData.NamespaceTable = openconfig.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    counters.EntityData.Leafs["bytes"] = types.YLeaf{"Bytes", counters.Bytes}
    counters.EntityData.Leafs["packets"] = types.YLeaf{"Packets", counters.Packets}
    counters.EntityData.Leafs["path-changes"] = types.YLeaf{"PathChanges", counters.PathChanges}
    counters.EntityData.Leafs["state-changes"] = types.YLeaf{"StateChanges", counters.StateChanges}
    counters.EntityData.Leafs["online-time"] = types.YLeaf{"OnlineTime", counters.OnlineTime}
    counters.EntityData.Leafs["current-path-time"] = types.YLeaf{"CurrentPathTime", counters.CurrentPathTime}
    counters.EntityData.Leafs["next-reoptimization-time"] = types.YLeaf{"NextReoptimizationTime", counters.NextReoptimizationTime}
    return &(counters.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth
// Bandwidth configuration for TE LSPs
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters related to bandwidth on TE tunnels:.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config

    // State parameters related to bandwidth configuration of TE tunnels.
    State Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State

    // Parameters related to auto-bandwidth.
    AutoBandwidth Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth
}

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "openconfig"
    bandwidth.EntityData.ParentYangName = "tunnel"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bandwidth.EntityData.Children = make(map[string]types.YChild)
    bandwidth.EntityData.Children["config"] = types.YChild{"Config", &bandwidth.Config}
    bandwidth.EntityData.Children["state"] = types.YChild{"State", &bandwidth.State}
    bandwidth.EntityData.Children["auto-bandwidth"] = types.YChild{"AutoBandwidth", &bandwidth.AutoBandwidth}
    bandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bandwidth.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config
// Configuration parameters related to bandwidth on TE
// tunnels:
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The method used for settign the bandwidth, either explicitly specified or
    // configured. The type is TeBandwidthType. The default value is SPECIFIED.
    SpecificationType interface{}

    // set bandwidth explicitly, e.g., using offline calculation. The type is
    // interface{} with range: 0..4294967295.
    SetBandwidth interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "bandwidth"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["specification-type"] = types.YLeaf{"SpecificationType", config.SpecificationType}
    config.EntityData.Leafs["set-bandwidth"] = types.YLeaf{"SetBandwidth", config.SetBandwidth}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State
// State parameters related to bandwidth
// configuration of TE tunnels
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The method used for settign the bandwidth, either explicitly specified or
    // configured. The type is TeBandwidthType. The default value is SPECIFIED.
    SpecificationType interface{}

    // set bandwidth explicitly, e.g., using offline calculation. The type is
    // interface{} with range: 0..4294967295.
    SetBandwidth interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "bandwidth"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["specification-type"] = types.YLeaf{"SpecificationType", state.SpecificationType}
    state.EntityData.Leafs["set-bandwidth"] = types.YLeaf{"SetBandwidth", state.SetBandwidth}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth
// Parameters related to auto-bandwidth
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to MPLS auto-bandwidth on the tunnel.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config

    // State parameters relating to MPLS auto-bandwidth on the tunnel.
    State Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State

    // configuration of MPLS overflow bandwidth adjustement for the LSP.
    Overflow Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow

    // configuration of MPLS underflow bandwidth           adjustement for the
    // LSP.
    Underflow Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow
}

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetEntityData() *types.CommonEntityData {
    autoBandwidth.EntityData.YFilter = autoBandwidth.YFilter
    autoBandwidth.EntityData.YangName = "auto-bandwidth"
    autoBandwidth.EntityData.BundleName = "openconfig"
    autoBandwidth.EntityData.ParentYangName = "bandwidth"
    autoBandwidth.EntityData.SegmentPath = "auto-bandwidth"
    autoBandwidth.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    autoBandwidth.EntityData.NamespaceTable = openconfig.GetNamespaces()
    autoBandwidth.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    autoBandwidth.EntityData.Children = make(map[string]types.YChild)
    autoBandwidth.EntityData.Children["config"] = types.YChild{"Config", &autoBandwidth.Config}
    autoBandwidth.EntityData.Children["state"] = types.YChild{"State", &autoBandwidth.State}
    autoBandwidth.EntityData.Children["overflow"] = types.YChild{"Overflow", &autoBandwidth.Overflow}
    autoBandwidth.EntityData.Children["underflow"] = types.YChild{"Underflow", &autoBandwidth.Underflow}
    autoBandwidth.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(autoBandwidth.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config
// Configuration parameters relating to MPLS
// auto-bandwidth on the tunnel.
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enables mpls auto-bandwidth on the lsp. The type is bool. The default value
    // is false.
    Enabled interface{}

    // set the minimum bandwidth in Mbps for an auto-bandwidth LSP. The type is
    // interface{} with range: 0..4294967295.
    MinBw interface{}

    // set the maximum bandwidth in Mbps for an auto-bandwidth LSP. The type is
    // interface{} with range: 0..4294967295.
    MaxBw interface{}

    // time in seconds between adjustments to LSP bandwidth. The type is
    // interface{} with range: 0..4294967295.
    AdjustInterval interface{}

    // percentage difference between the LSP's specified bandwidth and its current
    // bandwidth allocation -- if the difference is greater than the specified
    // percentage, auto-bandwidth adjustment is triggered. The type is interface{}
    // with range: 0..100.
    AdjustThreshold interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "auto-bandwidth"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["min-bw"] = types.YLeaf{"MinBw", config.MinBw}
    config.EntityData.Leafs["max-bw"] = types.YLeaf{"MaxBw", config.MaxBw}
    config.EntityData.Leafs["adjust-interval"] = types.YLeaf{"AdjustInterval", config.AdjustInterval}
    config.EntityData.Leafs["adjust-threshold"] = types.YLeaf{"AdjustThreshold", config.AdjustThreshold}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State
// State parameters relating to MPLS
// auto-bandwidth on the tunnel.
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enables mpls auto-bandwidth on the lsp. The type is bool. The default value
    // is false.
    Enabled interface{}

    // set the minimum bandwidth in Mbps for an auto-bandwidth LSP. The type is
    // interface{} with range: 0..4294967295.
    MinBw interface{}

    // set the maximum bandwidth in Mbps for an auto-bandwidth LSP. The type is
    // interface{} with range: 0..4294967295.
    MaxBw interface{}

    // time in seconds between adjustments to LSP bandwidth. The type is
    // interface{} with range: 0..4294967295.
    AdjustInterval interface{}

    // percentage difference between the LSP's specified bandwidth and its current
    // bandwidth allocation -- if the difference is greater than the specified
    // percentage, auto-bandwidth adjustment is triggered. The type is interface{}
    // with range: 0..100.
    AdjustThreshold interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "auto-bandwidth"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["min-bw"] = types.YLeaf{"MinBw", state.MinBw}
    state.EntityData.Leafs["max-bw"] = types.YLeaf{"MaxBw", state.MaxBw}
    state.EntityData.Leafs["adjust-interval"] = types.YLeaf{"AdjustInterval", state.AdjustInterval}
    state.EntityData.Leafs["adjust-threshold"] = types.YLeaf{"AdjustThreshold", state.AdjustThreshold}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow
// configuration of MPLS overflow bandwidth
// adjustement for the LSP
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config information for MPLS overflow bandwidth adjustment.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config

    // Config information for MPLS overflow bandwidth adjustment.
    State Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State
}

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetEntityData() *types.CommonEntityData {
    overflow.EntityData.YFilter = overflow.YFilter
    overflow.EntityData.YangName = "overflow"
    overflow.EntityData.BundleName = "openconfig"
    overflow.EntityData.ParentYangName = "auto-bandwidth"
    overflow.EntityData.SegmentPath = "overflow"
    overflow.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    overflow.EntityData.NamespaceTable = openconfig.GetNamespaces()
    overflow.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    overflow.EntityData.Children = make(map[string]types.YChild)
    overflow.EntityData.Children["config"] = types.YChild{"Config", &overflow.Config}
    overflow.EntityData.Children["state"] = types.YChild{"State", &overflow.State}
    overflow.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(overflow.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config
// Config information for MPLS overflow bandwidth
// adjustment
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enables mpls lsp bandwidth overflow adjustment on the lsp. The type is
    // bool. The default value is false.
    Enabled interface{}

    // bandwidth percentage change to trigger an overflow event. The type is
    // interface{} with range: 0..100.
    OverflowThreshold interface{}

    // number of consecutive overflow sample events needed to trigger an overflow
    // adjustment. The type is interface{} with range: 0..65535.
    TriggerEventCount interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "overflow"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["overflow-threshold"] = types.YLeaf{"OverflowThreshold", config.OverflowThreshold}
    config.EntityData.Leafs["trigger-event-count"] = types.YLeaf{"TriggerEventCount", config.TriggerEventCount}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State
// Config information for MPLS overflow bandwidth
// adjustment
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enables mpls lsp bandwidth overflow adjustment on the lsp. The type is
    // bool. The default value is false.
    Enabled interface{}

    // bandwidth percentage change to trigger an overflow event. The type is
    // interface{} with range: 0..100.
    OverflowThreshold interface{}

    // number of consecutive overflow sample events needed to trigger an overflow
    // adjustment. The type is interface{} with range: 0..65535.
    TriggerEventCount interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "overflow"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["overflow-threshold"] = types.YLeaf{"OverflowThreshold", state.OverflowThreshold}
    state.EntityData.Leafs["trigger-event-count"] = types.YLeaf{"TriggerEventCount", state.TriggerEventCount}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow
// configuration of MPLS underflow bandwidth
//           adjustement for the LSP
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config information for MPLS underflow bandwidth           adjustment.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config

    // State information for MPLS underflow bandwidth           adjustment.
    State Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State
}

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetEntityData() *types.CommonEntityData {
    underflow.EntityData.YFilter = underflow.YFilter
    underflow.EntityData.YangName = "underflow"
    underflow.EntityData.BundleName = "openconfig"
    underflow.EntityData.ParentYangName = "auto-bandwidth"
    underflow.EntityData.SegmentPath = "underflow"
    underflow.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    underflow.EntityData.NamespaceTable = openconfig.GetNamespaces()
    underflow.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    underflow.EntityData.Children = make(map[string]types.YChild)
    underflow.EntityData.Children["config"] = types.YChild{"Config", &underflow.Config}
    underflow.EntityData.Children["state"] = types.YChild{"State", &underflow.State}
    underflow.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(underflow.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config
// Config information for MPLS underflow bandwidth
//           adjustment
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enables bandwidth underflow adjustment on the lsp. The type is bool. The
    // default value is false.
    Enabled interface{}

    // bandwidth percentage change to trigger and underflow event. The type is
    // interface{} with range: 0..100.
    UnderflowThreshold interface{}

    // number of consecutive underflow sample events needed to trigger an
    // underflow adjustment. The type is interface{} with range: 0..65535.
    TriggerEventCount interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "underflow"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", config.Enabled}
    config.EntityData.Leafs["underflow-threshold"] = types.YLeaf{"UnderflowThreshold", config.UnderflowThreshold}
    config.EntityData.Leafs["trigger-event-count"] = types.YLeaf{"TriggerEventCount", config.TriggerEventCount}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State
// State information for MPLS underflow bandwidth
//           adjustment
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enables bandwidth underflow adjustment on the lsp. The type is bool. The
    // default value is false.
    Enabled interface{}

    // bandwidth percentage change to trigger and underflow event. The type is
    // interface{} with range: 0..100.
    UnderflowThreshold interface{}

    // number of consecutive underflow sample events needed to trigger an
    // underflow adjustment. The type is interface{} with range: 0..65535.
    TriggerEventCount interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "underflow"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", state.Enabled}
    state.EntityData.Leafs["underflow-threshold"] = types.YLeaf{"UnderflowThreshold", state.UnderflowThreshold}
    state.EntityData.Leafs["trigger-event-count"] = types.YLeaf{"TriggerEventCount", state.TriggerEventCount}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes
// Parameters related to LSPs of type P2P
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters for P2P LSPs.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config

    // State parameters for P2P LSPs.
    State Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State

    // List of p2p primary paths for a tunnel. The type is slice of
    // Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths.
    P2PPrimaryPaths []Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths

    // List of p2p primary paths for a tunnel. The type is slice of
    // Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths.
    P2PSecondaryPaths []Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths
}

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetEntityData() *types.CommonEntityData {
    p2PTunnelAttributes.EntityData.YFilter = p2PTunnelAttributes.YFilter
    p2PTunnelAttributes.EntityData.YangName = "p2p-tunnel-attributes"
    p2PTunnelAttributes.EntityData.BundleName = "openconfig"
    p2PTunnelAttributes.EntityData.ParentYangName = "tunnel"
    p2PTunnelAttributes.EntityData.SegmentPath = "p2p-tunnel-attributes"
    p2PTunnelAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2PTunnelAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2PTunnelAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2PTunnelAttributes.EntityData.Children = make(map[string]types.YChild)
    p2PTunnelAttributes.EntityData.Children["config"] = types.YChild{"Config", &p2PTunnelAttributes.Config}
    p2PTunnelAttributes.EntityData.Children["state"] = types.YChild{"State", &p2PTunnelAttributes.State}
    p2PTunnelAttributes.EntityData.Children["p2p-primary-paths"] = types.YChild{"P2PPrimaryPaths", nil}
    for i := range p2PTunnelAttributes.P2PPrimaryPaths {
        p2PTunnelAttributes.EntityData.Children[types.GetSegmentPath(&p2PTunnelAttributes.P2PPrimaryPaths[i])] = types.YChild{"P2PPrimaryPaths", &p2PTunnelAttributes.P2PPrimaryPaths[i]}
    }
    p2PTunnelAttributes.EntityData.Children["p2p-secondary-paths"] = types.YChild{"P2PSecondaryPaths", nil}
    for i := range p2PTunnelAttributes.P2PSecondaryPaths {
        p2PTunnelAttributes.EntityData.Children[types.GetSegmentPath(&p2PTunnelAttributes.P2PSecondaryPaths[i])] = types.YChild{"P2PSecondaryPaths", &p2PTunnelAttributes.P2PSecondaryPaths[i]}
    }
    p2PTunnelAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(p2PTunnelAttributes.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config
// Configuration parameters for P2P LSPs
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // P2P tunnel destination address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Destination interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "p2p-tunnel-attributes"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["destination"] = types.YLeaf{"Destination", config.Destination}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State
// State parameters for P2P LSPs
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // P2P tunnel destination address. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Destination interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "p2p-tunnel-attributes"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["destination"] = types.YLeaf{"Destination", state.Destination}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths
// List of p2p primary paths for a tunnel
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path name. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config_Name
    Name interface{}

    // Configuration parameters related to paths.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config

    // State parameters related to paths.
    State Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State

    // The set of candidate secondary paths which may be used for this primary
    // path. When secondary paths are specified in the list the path of the
    // secondary LSP in use must be restricted to those path options referenced.
    // The priority of the secondary paths is specified within the list. Higher
    // priority values are less preferred - that is to say that a path with
    // priority 0 is the most preferred path. In the case that the list is empty,
    // any secondary path option may be utilised when the current primary path is
    // in use.
    CandidateSecondaryPaths Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths

    // Top-level container for include/exclude constraints for link affinities.
    AdminGroups Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups
}

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetEntityData() *types.CommonEntityData {
    p2PPrimaryPaths.EntityData.YFilter = p2PPrimaryPaths.YFilter
    p2PPrimaryPaths.EntityData.YangName = "p2p-primary-paths"
    p2PPrimaryPaths.EntityData.BundleName = "openconfig"
    p2PPrimaryPaths.EntityData.ParentYangName = "p2p-tunnel-attributes"
    p2PPrimaryPaths.EntityData.SegmentPath = "p2p-primary-paths" + "[name='" + fmt.Sprintf("%v", p2PPrimaryPaths.Name) + "']"
    p2PPrimaryPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2PPrimaryPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2PPrimaryPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2PPrimaryPaths.EntityData.Children = make(map[string]types.YChild)
    p2PPrimaryPaths.EntityData.Children["config"] = types.YChild{"Config", &p2PPrimaryPaths.Config}
    p2PPrimaryPaths.EntityData.Children["state"] = types.YChild{"State", &p2PPrimaryPaths.State}
    p2PPrimaryPaths.EntityData.Children["candidate-secondary-paths"] = types.YChild{"CandidateSecondaryPaths", &p2PPrimaryPaths.CandidateSecondaryPaths}
    p2PPrimaryPaths.EntityData.Children["admin-groups"] = types.YChild{"AdminGroups", &p2PPrimaryPaths.AdminGroups}
    p2PPrimaryPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    p2PPrimaryPaths.EntityData.Leafs["name"] = types.YLeaf{"Name", p2PPrimaryPaths.Name}
    return &(p2PPrimaryPaths.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config
// Configuration parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LocallyComputedExternallyQueriedExplicitlyDefined.
    // The default value is locally-computed.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PathComputationServer interface{}

    // reference to a defined path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config_Name
    ExplicitPathName interface{}

    // Specifies a preference for this path. The lower the number higher the
    // preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // RSVP-TE preemption priority during LSP setup, lower is higher priority;
    // default 7 indicates that LSP will not preempt established LSPs during
    // setup. The type is interface{} with range: 0..7. The default value is 7.
    SetupPriority interface{}

    // preemption priority once the LSP is established, lower is higher priority;
    // default 0 indicates other LSPs will not preempt the LSPs once established.
    // The type is interface{} with range: 0..7. The default value is 0.
    HoldPriority interface{}

    // sets the time between attempts to establish the LSP. The type is
    // interface{} with range: 1..600. Units are seconds.
    RetryTimer interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "p2p-primary-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    config.EntityData.Leafs["path-computation-method"] = types.YLeaf{"PathComputationMethod", config.PathComputationMethod}
    config.EntityData.Leafs["use-cspf"] = types.YLeaf{"UseCspf", config.UseCspf}
    config.EntityData.Leafs["cspf-tiebreaker"] = types.YLeaf{"CspfTiebreaker", config.CspfTiebreaker}
    config.EntityData.Leafs["path-computation-server"] = types.YLeaf{"PathComputationServer", config.PathComputationServer}
    config.EntityData.Leafs["explicit-path-name"] = types.YLeaf{"ExplicitPathName", config.ExplicitPathName}
    config.EntityData.Leafs["preference"] = types.YLeaf{"Preference", config.Preference}
    config.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", config.SetupPriority}
    config.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", config.HoldPriority}
    config.EntityData.Leafs["retry-timer"] = types.YLeaf{"RetryTimer", config.RetryTimer}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State
// State parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LocallyComputedExternallyQueriedExplicitlyDefined.
    // The default value is locally-computed.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PathComputationServer interface{}

    // reference to a defined path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config_Name
    ExplicitPathName interface{}

    // Specifies a preference for this path. The lower the number higher the
    // preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // RSVP-TE preemption priority during LSP setup, lower is higher priority;
    // default 7 indicates that LSP will not preempt established LSPs during
    // setup. The type is interface{} with range: 0..7. The default value is 7.
    SetupPriority interface{}

    // preemption priority once the LSP is established, lower is higher priority;
    // default 0 indicates other LSPs will not preempt the LSPs once established.
    // The type is interface{} with range: 0..7. The default value is 0.
    HoldPriority interface{}

    // sets the time between attempts to establish the LSP. The type is
    // interface{} with range: 1..600. Units are seconds.
    RetryTimer interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "p2p-primary-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    state.EntityData.Leafs["path-computation-method"] = types.YLeaf{"PathComputationMethod", state.PathComputationMethod}
    state.EntityData.Leafs["use-cspf"] = types.YLeaf{"UseCspf", state.UseCspf}
    state.EntityData.Leafs["cspf-tiebreaker"] = types.YLeaf{"CspfTiebreaker", state.CspfTiebreaker}
    state.EntityData.Leafs["path-computation-server"] = types.YLeaf{"PathComputationServer", state.PathComputationServer}
    state.EntityData.Leafs["explicit-path-name"] = types.YLeaf{"ExplicitPathName", state.ExplicitPathName}
    state.EntityData.Leafs["preference"] = types.YLeaf{"Preference", state.Preference}
    state.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", state.SetupPriority}
    state.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", state.HoldPriority}
    state.EntityData.Leafs["retry-timer"] = types.YLeaf{"RetryTimer", state.RetryTimer}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths
// The set of candidate secondary paths which may be used
// for this primary path. When secondary paths are specified
// in the list the path of the secondary LSP in use must be
// restricted to those path options referenced. The
// priority of the secondary paths is specified within the
// list. Higher priority values are less preferred - that is
// to say that a path with priority 0 is the most preferred
// path. In the case that the list is empty, any secondary
// path option may be utilised when the current primary path
// is in use.
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of secondary paths which may be utilised when the current primary path
    // is in use. The type is slice of
    // Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath.
    CandidateSecondaryPath []Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath
}

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetEntityData() *types.CommonEntityData {
    candidateSecondaryPaths.EntityData.YFilter = candidateSecondaryPaths.YFilter
    candidateSecondaryPaths.EntityData.YangName = "candidate-secondary-paths"
    candidateSecondaryPaths.EntityData.BundleName = "openconfig"
    candidateSecondaryPaths.EntityData.ParentYangName = "p2p-primary-paths"
    candidateSecondaryPaths.EntityData.SegmentPath = "candidate-secondary-paths"
    candidateSecondaryPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    candidateSecondaryPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    candidateSecondaryPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    candidateSecondaryPaths.EntityData.Children = make(map[string]types.YChild)
    candidateSecondaryPaths.EntityData.Children["candidate-secondary-path"] = types.YChild{"CandidateSecondaryPath", nil}
    for i := range candidateSecondaryPaths.CandidateSecondaryPath {
        candidateSecondaryPaths.EntityData.Children[types.GetSegmentPath(&candidateSecondaryPaths.CandidateSecondaryPath[i])] = types.YChild{"CandidateSecondaryPath", &candidateSecondaryPaths.CandidateSecondaryPath[i]}
    }
    candidateSecondaryPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(candidateSecondaryPaths.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath
// List of secondary paths which may be utilised when the
// current primary path is in use
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A reference to the secondary path option reference
    // which acts as the key of the candidate-secondary-path list. The type is
    // string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config_SecondaryPath
    SecondaryPath interface{}

    // Configuration parameters relating to the candidate secondary path.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config

    // Operational state parameters relating to the candidate secondary path.
    State Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State
}

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetEntityData() *types.CommonEntityData {
    candidateSecondaryPath.EntityData.YFilter = candidateSecondaryPath.YFilter
    candidateSecondaryPath.EntityData.YangName = "candidate-secondary-path"
    candidateSecondaryPath.EntityData.BundleName = "openconfig"
    candidateSecondaryPath.EntityData.ParentYangName = "candidate-secondary-paths"
    candidateSecondaryPath.EntityData.SegmentPath = "candidate-secondary-path" + "[secondary-path='" + fmt.Sprintf("%v", candidateSecondaryPath.SecondaryPath) + "']"
    candidateSecondaryPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    candidateSecondaryPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    candidateSecondaryPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    candidateSecondaryPath.EntityData.Children = make(map[string]types.YChild)
    candidateSecondaryPath.EntityData.Children["config"] = types.YChild{"Config", &candidateSecondaryPath.Config}
    candidateSecondaryPath.EntityData.Children["state"] = types.YChild{"State", &candidateSecondaryPath.State}
    candidateSecondaryPath.EntityData.Leafs = make(map[string]types.YLeaf)
    candidateSecondaryPath.EntityData.Leafs["secondary-path"] = types.YLeaf{"SecondaryPath", candidateSecondaryPath.SecondaryPath}
    return &(candidateSecondaryPath.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config
// Configuration parameters relating to the candidate
// secondary path
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A reference to the secondary path that should be utilised when the
    // containing primary path option is in use. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config_Name
    SecondaryPath interface{}

    // The priority of the specified secondary path option. Higher priority
    // options are less preferable - such that a secondary path reference with a
    // priority of 0 is the most preferred. The type is interface{} with range:
    // 0..65535.
    Priority interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "candidate-secondary-path"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["secondary-path"] = types.YLeaf{"SecondaryPath", config.SecondaryPath}
    config.EntityData.Leafs["priority"] = types.YLeaf{"Priority", config.Priority}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State
// Operational state parameters relating to the candidate
// secondary path
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A reference to the secondary path that should be utilised when the
    // containing primary path option is in use. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config_Name
    SecondaryPath interface{}

    // The priority of the specified secondary path option. Higher priority
    // options are less preferable - such that a secondary path reference with a
    // priority of 0 is the most preferred. The type is interface{} with range:
    // 0..65535.
    Priority interface{}

    // Indicates the current active path option that has been selected of the
    // candidate secondary paths. The type is bool.
    Active interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "candidate-secondary-path"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["secondary-path"] = types.YLeaf{"SecondaryPath", state.SecondaryPath}
    state.EntityData.Leafs["priority"] = types.YLeaf{"Priority", state.Priority}
    state.EntityData.Leafs["active"] = types.YLeaf{"Active", state.Active}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups
// Top-level container for include/exclude constraints for
// link affinities
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data .
    Config Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config

    // Operational state data .
    State Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetEntityData() *types.CommonEntityData {
    adminGroups.EntityData.YFilter = adminGroups.YFilter
    adminGroups.EntityData.YangName = "admin-groups"
    adminGroups.EntityData.BundleName = "openconfig"
    adminGroups.EntityData.ParentYangName = "p2p-primary-paths"
    adminGroups.EntityData.SegmentPath = "admin-groups"
    adminGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adminGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adminGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adminGroups.EntityData.Children = make(map[string]types.YChild)
    adminGroups.EntityData.Children["config"] = types.YChild{"Config", &adminGroups.Config}
    adminGroups.EntityData.Children["state"] = types.YChild{"State", &adminGroups.State}
    adminGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adminGroups.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config
// Configuration data 
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of references to named admin-groups to exclude in path calculation.
    // The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    ExcludeGroup []interface{}

    // list of references to named admin-groups of which all must be included. The
    // type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    IncludeAllGroup []interface{}

    // list of references to named admin-groups of which one must be included. The
    // type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    IncludeAnyGroup []interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "admin-groups"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["exclude-group"] = types.YLeaf{"ExcludeGroup", config.ExcludeGroup}
    config.EntityData.Leafs["include-all-group"] = types.YLeaf{"IncludeAllGroup", config.IncludeAllGroup}
    config.EntityData.Leafs["include-any-group"] = types.YLeaf{"IncludeAnyGroup", config.IncludeAnyGroup}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State
// Operational state data 
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of references to named admin-groups to exclude in path calculation.
    // The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    ExcludeGroup []interface{}

    // list of references to named admin-groups of which all must be included. The
    // type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    IncludeAllGroup []interface{}

    // list of references to named admin-groups of which one must be included. The
    // type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    IncludeAnyGroup []interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "admin-groups"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["exclude-group"] = types.YLeaf{"ExcludeGroup", state.ExcludeGroup}
    state.EntityData.Leafs["include-all-group"] = types.YLeaf{"IncludeAllGroup", state.IncludeAllGroup}
    state.EntityData.Leafs["include-any-group"] = types.YLeaf{"IncludeAnyGroup", state.IncludeAnyGroup}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths
// List of p2p primary paths for a tunnel
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path name. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config_Name
    Name interface{}

    // Configuration parameters related to paths.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config

    // State parameters related to paths.
    State Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State

    // Top-level container for include/exclude constraints for link affinities.
    AdminGroups Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups
}

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetEntityData() *types.CommonEntityData {
    p2PSecondaryPaths.EntityData.YFilter = p2PSecondaryPaths.YFilter
    p2PSecondaryPaths.EntityData.YangName = "p2p-secondary-paths"
    p2PSecondaryPaths.EntityData.BundleName = "openconfig"
    p2PSecondaryPaths.EntityData.ParentYangName = "p2p-tunnel-attributes"
    p2PSecondaryPaths.EntityData.SegmentPath = "p2p-secondary-paths" + "[name='" + fmt.Sprintf("%v", p2PSecondaryPaths.Name) + "']"
    p2PSecondaryPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2PSecondaryPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2PSecondaryPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2PSecondaryPaths.EntityData.Children = make(map[string]types.YChild)
    p2PSecondaryPaths.EntityData.Children["config"] = types.YChild{"Config", &p2PSecondaryPaths.Config}
    p2PSecondaryPaths.EntityData.Children["state"] = types.YChild{"State", &p2PSecondaryPaths.State}
    p2PSecondaryPaths.EntityData.Children["admin-groups"] = types.YChild{"AdminGroups", &p2PSecondaryPaths.AdminGroups}
    p2PSecondaryPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    p2PSecondaryPaths.EntityData.Leafs["name"] = types.YLeaf{"Name", p2PSecondaryPaths.Name}
    return &(p2PSecondaryPaths.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config
// Configuration parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LocallyComputedExternallyQueriedExplicitlyDefined.
    // The default value is locally-computed.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PathComputationServer interface{}

    // reference to a defined path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config_Name
    ExplicitPathName interface{}

    // Specifies a preference for this path. The lower the number higher the
    // preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // RSVP-TE preemption priority during LSP setup, lower is higher priority;
    // default 7 indicates that LSP will not preempt established LSPs during
    // setup. The type is interface{} with range: 0..7. The default value is 7.
    SetupPriority interface{}

    // preemption priority once the LSP is established, lower is higher priority;
    // default 0 indicates other LSPs will not preempt the LSPs once established.
    // The type is interface{} with range: 0..7. The default value is 0.
    HoldPriority interface{}

    // sets the time between attempts to establish the LSP. The type is
    // interface{} with range: 1..600. Units are seconds.
    RetryTimer interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "p2p-secondary-paths"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["name"] = types.YLeaf{"Name", config.Name}
    config.EntityData.Leafs["path-computation-method"] = types.YLeaf{"PathComputationMethod", config.PathComputationMethod}
    config.EntityData.Leafs["use-cspf"] = types.YLeaf{"UseCspf", config.UseCspf}
    config.EntityData.Leafs["cspf-tiebreaker"] = types.YLeaf{"CspfTiebreaker", config.CspfTiebreaker}
    config.EntityData.Leafs["path-computation-server"] = types.YLeaf{"PathComputationServer", config.PathComputationServer}
    config.EntityData.Leafs["explicit-path-name"] = types.YLeaf{"ExplicitPathName", config.ExplicitPathName}
    config.EntityData.Leafs["preference"] = types.YLeaf{"Preference", config.Preference}
    config.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", config.SetupPriority}
    config.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", config.HoldPriority}
    config.EntityData.Leafs["retry-timer"] = types.YLeaf{"RetryTimer", config.RetryTimer}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State
// State parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LocallyComputedExternallyQueriedExplicitlyDefined.
    // The default value is locally-computed.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PathComputationServer interface{}

    // reference to a defined path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config_Name
    ExplicitPathName interface{}

    // Specifies a preference for this path. The lower the number higher the
    // preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // RSVP-TE preemption priority during LSP setup, lower is higher priority;
    // default 7 indicates that LSP will not preempt established LSPs during
    // setup. The type is interface{} with range: 0..7. The default value is 7.
    SetupPriority interface{}

    // preemption priority once the LSP is established, lower is higher priority;
    // default 0 indicates other LSPs will not preempt the LSPs once established.
    // The type is interface{} with range: 0..7. The default value is 0.
    HoldPriority interface{}

    // sets the time between attempts to establish the LSP. The type is
    // interface{} with range: 1..600. Units are seconds.
    RetryTimer interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "p2p-secondary-paths"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["name"] = types.YLeaf{"Name", state.Name}
    state.EntityData.Leafs["path-computation-method"] = types.YLeaf{"PathComputationMethod", state.PathComputationMethod}
    state.EntityData.Leafs["use-cspf"] = types.YLeaf{"UseCspf", state.UseCspf}
    state.EntityData.Leafs["cspf-tiebreaker"] = types.YLeaf{"CspfTiebreaker", state.CspfTiebreaker}
    state.EntityData.Leafs["path-computation-server"] = types.YLeaf{"PathComputationServer", state.PathComputationServer}
    state.EntityData.Leafs["explicit-path-name"] = types.YLeaf{"ExplicitPathName", state.ExplicitPathName}
    state.EntityData.Leafs["preference"] = types.YLeaf{"Preference", state.Preference}
    state.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", state.SetupPriority}
    state.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", state.HoldPriority}
    state.EntityData.Leafs["retry-timer"] = types.YLeaf{"RetryTimer", state.RetryTimer}
    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups
// Top-level container for include/exclude constraints for
// link affinities
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data .
    Config Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config

    // Operational state data .
    State Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetEntityData() *types.CommonEntityData {
    adminGroups.EntityData.YFilter = adminGroups.YFilter
    adminGroups.EntityData.YangName = "admin-groups"
    adminGroups.EntityData.BundleName = "openconfig"
    adminGroups.EntityData.ParentYangName = "p2p-secondary-paths"
    adminGroups.EntityData.SegmentPath = "admin-groups"
    adminGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adminGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adminGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adminGroups.EntityData.Children = make(map[string]types.YChild)
    adminGroups.EntityData.Children["config"] = types.YChild{"Config", &adminGroups.Config}
    adminGroups.EntityData.Children["state"] = types.YChild{"State", &adminGroups.State}
    adminGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adminGroups.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config
// Configuration data 
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of references to named admin-groups to exclude in path calculation.
    // The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    ExcludeGroup []interface{}

    // list of references to named admin-groups of which all must be included. The
    // type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    IncludeAllGroup []interface{}

    // list of references to named admin-groups of which one must be included. The
    // type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    IncludeAnyGroup []interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "admin-groups"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["exclude-group"] = types.YLeaf{"ExcludeGroup", config.ExcludeGroup}
    config.EntityData.Leafs["include-all-group"] = types.YLeaf{"IncludeAllGroup", config.IncludeAllGroup}
    config.EntityData.Leafs["include-any-group"] = types.YLeaf{"IncludeAnyGroup", config.IncludeAnyGroup}
    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State
// Operational state data 
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of references to named admin-groups to exclude in path calculation.
    // The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    ExcludeGroup []interface{}

    // list of references to named admin-groups of which all must be included. The
    // type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    IncludeAllGroup []interface{}

    // list of references to named admin-groups of which one must be included. The
    // type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_AdminGroupName
    IncludeAnyGroup []interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "admin-groups"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["exclude-group"] = types.YLeaf{"ExcludeGroup", state.ExcludeGroup}
    state.EntityData.Leafs["include-all-group"] = types.YLeaf{"IncludeAllGroup", state.IncludeAllGroup}
    state.EntityData.Leafs["include-any-group"] = types.YLeaf{"IncludeAnyGroup", state.IncludeAnyGroup}
    return &(state.EntityData)
}

// Mpls_Lsps_UnconstrainedPath
// LSPs that use the IGP-determined path, i.e., non
// traffic-engineered, or non constrained-path
type Mpls_Lsps_UnconstrainedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // select and configure the signaling method for  the LSP.
    PathSetupProtocol Mpls_Lsps_UnconstrainedPath_PathSetupProtocol
}

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetEntityData() *types.CommonEntityData {
    unconstrainedPath.EntityData.YFilter = unconstrainedPath.YFilter
    unconstrainedPath.EntityData.YangName = "unconstrained-path"
    unconstrainedPath.EntityData.BundleName = "openconfig"
    unconstrainedPath.EntityData.ParentYangName = "lsps"
    unconstrainedPath.EntityData.SegmentPath = "unconstrained-path"
    unconstrainedPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unconstrainedPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unconstrainedPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unconstrainedPath.EntityData.Children = make(map[string]types.YChild)
    unconstrainedPath.EntityData.Children["path-setup-protocol"] = types.YChild{"PathSetupProtocol", &unconstrainedPath.PathSetupProtocol}
    unconstrainedPath.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(unconstrainedPath.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol
// select and configure the signaling method for
//  the LSP
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP signaling setup for IGP-congruent LSPs.
    Ldp Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp

    // segment routing signaling extensions for IGP-confgruent LSPs.
    SegmentRouting Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting
}

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetEntityData() *types.CommonEntityData {
    pathSetupProtocol.EntityData.YFilter = pathSetupProtocol.YFilter
    pathSetupProtocol.EntityData.YangName = "path-setup-protocol"
    pathSetupProtocol.EntityData.BundleName = "openconfig"
    pathSetupProtocol.EntityData.ParentYangName = "unconstrained-path"
    pathSetupProtocol.EntityData.SegmentPath = "path-setup-protocol"
    pathSetupProtocol.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    pathSetupProtocol.EntityData.NamespaceTable = openconfig.GetNamespaces()
    pathSetupProtocol.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    pathSetupProtocol.EntityData.Children = make(map[string]types.YChild)
    pathSetupProtocol.EntityData.Children["ldp"] = types.YChild{"Ldp", &pathSetupProtocol.Ldp}
    pathSetupProtocol.EntityData.Children["segment-routing"] = types.YChild{"SegmentRouting", &pathSetupProtocol.SegmentRouting}
    pathSetupProtocol.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pathSetupProtocol.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp
// LDP signaling setup for IGP-congruent LSPs
// This type is a presence type.
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // contains configuration stanzas for different LSP tunnel types (P2P, P2MP,
    // etc.).
    Tunnel Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel
}

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetEntityData() *types.CommonEntityData {
    ldp.EntityData.YFilter = ldp.YFilter
    ldp.EntityData.YangName = "ldp"
    ldp.EntityData.BundleName = "openconfig"
    ldp.EntityData.ParentYangName = "path-setup-protocol"
    ldp.EntityData.SegmentPath = "ldp"
    ldp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ldp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ldp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ldp.EntityData.Children = make(map[string]types.YChild)
    ldp.EntityData.Children["tunnel"] = types.YChild{"Tunnel", &ldp.Tunnel}
    ldp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ldp.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel
// contains configuration stanzas for different LSP
// tunnel types (P2P, P2MP, etc.)
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // specifies the type of LSP, e.g., P2P or P2MP. The type is TunnelType_.
    TunnelType interface{}

    // specify basic or targeted LDP LSP. The type is LdpType.
    LdpType interface{}

    // properties of point-to-point tunnels.
    P2PLsp Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp

    // properties of point-to-multipoint tunnels.
    P2MpLsp Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp

    // properties of multipoint-to-multipoint tunnels.
    Mp2MpLsp Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "openconfig"
    tunnel.EntityData.ParentYangName = "ldp"
    tunnel.EntityData.SegmentPath = "tunnel"
    tunnel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    tunnel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Children["p2p-lsp"] = types.YChild{"P2PLsp", &tunnel.P2PLsp}
    tunnel.EntityData.Children["p2mp-lsp"] = types.YChild{"P2MpLsp", &tunnel.P2MpLsp}
    tunnel.EntityData.Children["mp2mp-lsp"] = types.YChild{"Mp2MpLsp", &tunnel.Mp2MpLsp}
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["tunnel-type"] = types.YLeaf{"TunnelType", tunnel.TunnelType}
    tunnel.EntityData.Leafs["ldp-type"] = types.YLeaf{"LdpType", tunnel.LdpType}
    return &(tunnel.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp
// properties of point-to-point tunnels
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address prefix for packets sharing the same forwarding equivalence class
    // for the IGP-based LSP. The type is one of the following types: slice of
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or slice of string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    FecAddress []interface{}
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetEntityData() *types.CommonEntityData {
    p2PLsp.EntityData.YFilter = p2PLsp.YFilter
    p2PLsp.EntityData.YangName = "p2p-lsp"
    p2PLsp.EntityData.BundleName = "openconfig"
    p2PLsp.EntityData.ParentYangName = "tunnel"
    p2PLsp.EntityData.SegmentPath = "p2p-lsp"
    p2PLsp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2PLsp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2PLsp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2PLsp.EntityData.Children = make(map[string]types.YChild)
    p2PLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    p2PLsp.EntityData.Leafs["fec-address"] = types.YLeaf{"FecAddress", p2PLsp.FecAddress}
    return &(p2PLsp.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp
// properties of point-to-multipoint tunnels
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetEntityData() *types.CommonEntityData {
    p2MpLsp.EntityData.YFilter = p2MpLsp.YFilter
    p2MpLsp.EntityData.YangName = "p2mp-lsp"
    p2MpLsp.EntityData.BundleName = "openconfig"
    p2MpLsp.EntityData.ParentYangName = "tunnel"
    p2MpLsp.EntityData.SegmentPath = "p2mp-lsp"
    p2MpLsp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2MpLsp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2MpLsp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2MpLsp.EntityData.Children = make(map[string]types.YChild)
    p2MpLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(p2MpLsp.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp
// properties of multipoint-to-multipoint tunnels
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetEntityData() *types.CommonEntityData {
    mp2MpLsp.EntityData.YFilter = mp2MpLsp.YFilter
    mp2MpLsp.EntityData.YangName = "mp2mp-lsp"
    mp2MpLsp.EntityData.BundleName = "openconfig"
    mp2MpLsp.EntityData.ParentYangName = "tunnel"
    mp2MpLsp.EntityData.SegmentPath = "mp2mp-lsp"
    mp2MpLsp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    mp2MpLsp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    mp2MpLsp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    mp2MpLsp.EntityData.Children = make(map[string]types.YChild)
    mp2MpLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mp2MpLsp.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_LdpType represents specify basic or targeted LDP LSP
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_LdpType string

const (
    // basic hop-by-hop LSP
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_LdpType_BASIC Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_LdpType = "BASIC"

    // tLDP LSP
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_LdpType_TARGETED Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_LdpType = "TARGETED"
)

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting
// segment routing signaling extensions for
// IGP-confgruent LSPs
// This type is a presence type.
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // contains configuration stanzas for different LSP tunnel types (P2P, P2MP,
    // etc.).
    Tunnel Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel
}

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetEntityData() *types.CommonEntityData {
    segmentRouting.EntityData.YFilter = segmentRouting.YFilter
    segmentRouting.EntityData.YangName = "segment-routing"
    segmentRouting.EntityData.BundleName = "openconfig"
    segmentRouting.EntityData.ParentYangName = "path-setup-protocol"
    segmentRouting.EntityData.SegmentPath = "segment-routing"
    segmentRouting.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    segmentRouting.EntityData.NamespaceTable = openconfig.GetNamespaces()
    segmentRouting.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    segmentRouting.EntityData.Children = make(map[string]types.YChild)
    segmentRouting.EntityData.Children["tunnel"] = types.YChild{"Tunnel", &segmentRouting.Tunnel}
    segmentRouting.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(segmentRouting.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel
// contains configuration stanzas for different LSP
// tunnel types (P2P, P2MP, etc.)
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // specifies the type of LSP, e.g., P2P or P2MP. The type is TunnelType_.
    TunnelType interface{}

    // properties of point-to-point tunnels.
    P2PLsp Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "openconfig"
    tunnel.EntityData.ParentYangName = "segment-routing"
    tunnel.EntityData.SegmentPath = "tunnel"
    tunnel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    tunnel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Children["p2p-lsp"] = types.YChild{"P2PLsp", &tunnel.P2PLsp}
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["tunnel-type"] = types.YLeaf{"TunnelType", tunnel.TunnelType}
    return &(tunnel.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp
// properties of point-to-point tunnels
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of FECs that are to be originated as SR LSPs. The type is slice of
    // Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec.
    Fec []Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetEntityData() *types.CommonEntityData {
    p2PLsp.EntityData.YFilter = p2PLsp.YFilter
    p2PLsp.EntityData.YangName = "p2p-lsp"
    p2PLsp.EntityData.BundleName = "openconfig"
    p2PLsp.EntityData.ParentYangName = "tunnel"
    p2PLsp.EntityData.SegmentPath = "p2p-lsp"
    p2PLsp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2PLsp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2PLsp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2PLsp.EntityData.Children = make(map[string]types.YChild)
    p2PLsp.EntityData.Children["fec"] = types.YChild{"Fec", nil}
    for i := range p2PLsp.Fec {
        p2PLsp.EntityData.Children[types.GetSegmentPath(&p2PLsp.Fec[i])] = types.YChild{"Fec", &p2PLsp.Fec[i]}
    }
    p2PLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(p2PLsp.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec
// List of FECs that are to be originated as SR LSPs
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. FEC that is to be advertised as part of the
    // Prefix-SID. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    FecAddress interface{}

    // Configuration parameters relating to the FEC to be advertised by SR.
    Config Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config

    // Operational state relating to a FEC advertised by SR.
    State Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State

    // Parameters relating to the Prefix-SID used for the originated FEC.
    PrefixSid Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid
}

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetEntityData() *types.CommonEntityData {
    fec.EntityData.YFilter = fec.YFilter
    fec.EntityData.YangName = "fec"
    fec.EntityData.BundleName = "openconfig"
    fec.EntityData.ParentYangName = "p2p-lsp"
    fec.EntityData.SegmentPath = "fec" + "[fec-address='" + fmt.Sprintf("%v", fec.FecAddress) + "']"
    fec.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    fec.EntityData.NamespaceTable = openconfig.GetNamespaces()
    fec.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    fec.EntityData.Children = make(map[string]types.YChild)
    fec.EntityData.Children["config"] = types.YChild{"Config", &fec.Config}
    fec.EntityData.Children["state"] = types.YChild{"State", &fec.State}
    fec.EntityData.Children["prefix-sid"] = types.YChild{"PrefixSid", &fec.PrefixSid}
    fec.EntityData.Leafs = make(map[string]types.YLeaf)
    fec.EntityData.Leafs["fec-address"] = types.YLeaf{"FecAddress", fec.FecAddress}
    return &(fec.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config
// Configuration parameters relating to the FEC to be
// advertised by SR
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FEC that is to be advertised as part of the Prefix-SID. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    FecAddress interface{}
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "fec"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["fec-address"] = types.YLeaf{"FecAddress", config.FecAddress}
    return &(config.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State
// Operational state relating to a FEC advertised by SR
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FEC that is to be advertised as part of the Prefix-SID. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    FecAddress interface{}
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "fec"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["fec-address"] = types.YLeaf{"FecAddress", state.FecAddress}
    return &(state.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid
// Parameters relating to the Prefix-SID
// used for the originated FEC
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to the Prefix-SID used for the originated
    // FEC.
    Config Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config

    // Operational state parameters relating to the Prefix-SID used for the
    // originated FEC.
    State Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State
}

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetEntityData() *types.CommonEntityData {
    prefixSid.EntityData.YFilter = prefixSid.YFilter
    prefixSid.EntityData.YangName = "prefix-sid"
    prefixSid.EntityData.BundleName = "openconfig"
    prefixSid.EntityData.ParentYangName = "fec"
    prefixSid.EntityData.SegmentPath = "prefix-sid"
    prefixSid.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    prefixSid.EntityData.NamespaceTable = openconfig.GetNamespaces()
    prefixSid.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    prefixSid.EntityData.Children = make(map[string]types.YChild)
    prefixSid.EntityData.Children["config"] = types.YChild{"Config", &prefixSid.Config}
    prefixSid.EntityData.Children["state"] = types.YChild{"State", &prefixSid.State}
    prefixSid.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixSid.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config
// Configuration parameters relating to the Prefix-SID
// used for the originated FEC
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies how the value of the Prefix-SID should be interpreted - whether
    // as an offset to the SRGB, or as an absolute value. The type is Type_. The
    // default value is INDEX.
    Type_ interface{}

    // Specifies that the Prefix-SID is to be treated as a Node-SID by setting the
    // N-flag in the advertised Prefix-SID TLV in the IGP. The type is bool.
    NodeFlag interface{}

    // Configuration relating to the LFIB actions for the Prefix-SID to be used by
    // the penultimate-hop. The type is LastHopBehavior.
    LastHopBehavior interface{}
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "prefix-sid"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["type"] = types.YLeaf{"Type_", config.Type_}
    config.EntityData.Leafs["node-flag"] = types.YLeaf{"NodeFlag", config.NodeFlag}
    config.EntityData.Leafs["last-hop-behavior"] = types.YLeaf{"LastHopBehavior", config.LastHopBehavior}
    return &(config.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_LastHopBehavior represents Prefix-SID to be used by the penultimate-hop
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_LastHopBehavior string

const (
    // Specifies that the explicit null label is to be used
    // when the penultimate hop forwards a labelled packet to
    // this Prefix-SID
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_LastHopBehavior_EXPLICIT_NULL Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_LastHopBehavior = "EXPLICIT-NULL"

    // Specicies that the Prefix-SID's label value is to be
    // left in place when the penultimate hop forwards to this
    // Prefix-SID
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_LastHopBehavior_UNCHANGED Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_LastHopBehavior = "UNCHANGED"

    // Specicies that the penultimate hop should pop the
    // Prefix-SID label before forwarding to the eLER
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_LastHopBehavior_PHP Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_LastHopBehavior = "PHP"
)

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type_ represents absolute value
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type_ string

const (
    // Set when the value of the prefix SID should be specified
    // as an off-set from the SRGB's zero-value. When multiple
    // SRGBs are specified, the zero-value is the minimum
    // of their lower bounds
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type__INDEX Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type_ = "INDEX"

    // Set when the value of a prefix SID is specified as the
    // absolute value within an SRGB. It is an error to specify
    // an absolute value outside of a specified SRGB
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type__ABSOLUTE Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type_ = "ABSOLUTE"
)

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State
// Operational state parameters relating to the
// Prefix-SID used for the originated FEC
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies how the value of the Prefix-SID should be interpreted - whether
    // as an offset to the SRGB, or as an absolute value. The type is Type_. The
    // default value is INDEX.
    Type_ interface{}

    // Specifies that the Prefix-SID is to be treated as a Node-SID by setting the
    // N-flag in the advertised Prefix-SID TLV in the IGP. The type is bool.
    NodeFlag interface{}

    // Configuration relating to the LFIB actions for the Prefix-SID to be used by
    // the penultimate-hop. The type is LastHopBehavior.
    LastHopBehavior interface{}
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "prefix-sid"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["type"] = types.YLeaf{"Type_", state.Type_}
    state.EntityData.Leafs["node-flag"] = types.YLeaf{"NodeFlag", state.NodeFlag}
    state.EntityData.Leafs["last-hop-behavior"] = types.YLeaf{"LastHopBehavior", state.LastHopBehavior}
    return &(state.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_LastHopBehavior represents Prefix-SID to be used by the penultimate-hop
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_LastHopBehavior string

const (
    // Specifies that the explicit null label is to be used
    // when the penultimate hop forwards a labelled packet to
    // this Prefix-SID
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_LastHopBehavior_EXPLICIT_NULL Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_LastHopBehavior = "EXPLICIT-NULL"

    // Specicies that the Prefix-SID's label value is to be
    // left in place when the penultimate hop forwards to this
    // Prefix-SID
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_LastHopBehavior_UNCHANGED Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_LastHopBehavior = "UNCHANGED"

    // Specicies that the penultimate hop should pop the
    // Prefix-SID label before forwarding to the eLER
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_LastHopBehavior_PHP Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_LastHopBehavior = "PHP"
)

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type_ represents absolute value
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type_ string

const (
    // Set when the value of the prefix SID should be specified
    // as an off-set from the SRGB's zero-value. When multiple
    // SRGBs are specified, the zero-value is the minimum
    // of their lower bounds
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type__INDEX Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type_ = "INDEX"

    // Set when the value of a prefix SID is specified as the
    // absolute value within an SRGB. It is an error to specify
    // an absolute value outside of a specified SRGB
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type__ABSOLUTE Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type_ = "ABSOLUTE"
)

// Mpls_Lsps_StaticLsps
// statically configured LSPs, without dynamic
// signaling
type Mpls_Lsps_StaticLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of defined static LSPs. The type is slice of
    // Mpls_Lsps_StaticLsps_LabelSwitchedPath.
    LabelSwitchedPath []Mpls_Lsps_StaticLsps_LabelSwitchedPath
}

func (staticLsps *Mpls_Lsps_StaticLsps) GetEntityData() *types.CommonEntityData {
    staticLsps.EntityData.YFilter = staticLsps.YFilter
    staticLsps.EntityData.YangName = "static-lsps"
    staticLsps.EntityData.BundleName = "openconfig"
    staticLsps.EntityData.ParentYangName = "lsps"
    staticLsps.EntityData.SegmentPath = "static-lsps"
    staticLsps.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    staticLsps.EntityData.NamespaceTable = openconfig.GetNamespaces()
    staticLsps.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    staticLsps.EntityData.Children = make(map[string]types.YChild)
    staticLsps.EntityData.Children["label-switched-path"] = types.YChild{"LabelSwitchedPath", nil}
    for i := range staticLsps.LabelSwitchedPath {
        staticLsps.EntityData.Children[types.GetSegmentPath(&staticLsps.LabelSwitchedPath[i])] = types.YChild{"LabelSwitchedPath", &staticLsps.LabelSwitchedPath[i]}
    }
    staticLsps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticLsps.EntityData)
}

// Mpls_Lsps_StaticLsps_LabelSwitchedPath
// list of defined static LSPs
type Mpls_Lsps_StaticLsps_LabelSwitchedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. name to identify the LSP. The type is string.
    Name interface{}

    // Static LSPs for which the router is an ingress node.
    Ingress Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress

    // static LSPs for which the router is a transit node.
    Transit Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit

    // static LSPs for which the router is a egress  node.
    Egress Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress
}

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetEntityData() *types.CommonEntityData {
    labelSwitchedPath.EntityData.YFilter = labelSwitchedPath.YFilter
    labelSwitchedPath.EntityData.YangName = "label-switched-path"
    labelSwitchedPath.EntityData.BundleName = "openconfig"
    labelSwitchedPath.EntityData.ParentYangName = "static-lsps"
    labelSwitchedPath.EntityData.SegmentPath = "label-switched-path" + "[name='" + fmt.Sprintf("%v", labelSwitchedPath.Name) + "']"
    labelSwitchedPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    labelSwitchedPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    labelSwitchedPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    labelSwitchedPath.EntityData.Children = make(map[string]types.YChild)
    labelSwitchedPath.EntityData.Children["ingress"] = types.YChild{"Ingress", &labelSwitchedPath.Ingress}
    labelSwitchedPath.EntityData.Children["transit"] = types.YChild{"Transit", &labelSwitchedPath.Transit}
    labelSwitchedPath.EntityData.Children["egress"] = types.YChild{"Egress", &labelSwitchedPath.Egress}
    labelSwitchedPath.EntityData.Leafs = make(map[string]types.YLeaf)
    labelSwitchedPath.EntityData.Leafs["name"] = types.YLeaf{"Name", labelSwitchedPath.Name}
    return &(labelSwitchedPath.EntityData)
}

// Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress
// Static LSPs for which the router is an
// ingress node
type Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetEntityData() *types.CommonEntityData {
    ingress.EntityData.YFilter = ingress.YFilter
    ingress.EntityData.YangName = "ingress"
    ingress.EntityData.BundleName = "openconfig"
    ingress.EntityData.ParentYangName = "label-switched-path"
    ingress.EntityData.SegmentPath = "ingress"
    ingress.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ingress.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ingress.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ingress.EntityData.Children = make(map[string]types.YChild)
    ingress.EntityData.Leafs = make(map[string]types.YLeaf)
    ingress.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", ingress.NextHop}
    ingress.EntityData.Leafs["incoming-label"] = types.YLeaf{"IncomingLabel", ingress.IncomingLabel}
    ingress.EntityData.Leafs["push-label"] = types.YLeaf{"PushLabel", ingress.PushLabel}
    return &(ingress.EntityData)
}

// Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit
// static LSPs for which the router is a
// transit node
type Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetEntityData() *types.CommonEntityData {
    transit.EntityData.YFilter = transit.YFilter
    transit.EntityData.YangName = "transit"
    transit.EntityData.BundleName = "openconfig"
    transit.EntityData.ParentYangName = "label-switched-path"
    transit.EntityData.SegmentPath = "transit"
    transit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    transit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    transit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    transit.EntityData.Children = make(map[string]types.YChild)
    transit.EntityData.Leafs = make(map[string]types.YLeaf)
    transit.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", transit.NextHop}
    transit.EntityData.Leafs["incoming-label"] = types.YLeaf{"IncomingLabel", transit.IncomingLabel}
    transit.EntityData.Leafs["push-label"] = types.YLeaf{"PushLabel", transit.PushLabel}
    return &(transit.EntityData)
}

// Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress
// static LSPs for which the router is a
// egress  node
type Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetEntityData() *types.CommonEntityData {
    egress.EntityData.YFilter = egress.YFilter
    egress.EntityData.YangName = "egress"
    egress.EntityData.BundleName = "openconfig"
    egress.EntityData.ParentYangName = "label-switched-path"
    egress.EntityData.SegmentPath = "egress"
    egress.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    egress.EntityData.NamespaceTable = openconfig.GetNamespaces()
    egress.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    egress.EntityData.Children = make(map[string]types.YChild)
    egress.EntityData.Leafs = make(map[string]types.YLeaf)
    egress.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", egress.NextHop}
    egress.EntityData.Leafs["incoming-label"] = types.YLeaf{"IncomingLabel", egress.IncomingLabel}
    egress.EntityData.Leafs["push-label"] = types.YLeaf{"PushLabel", egress.PushLabel}
    return &(egress.EntityData)
}

