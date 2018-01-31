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

type ExplicitlyDefined struct {
}

func (id ExplicitlyDefined) String() string {
	return "openconfig-mpls-te:explicitly-defined"
}

type ExternallyQueried struct {
}

func (id ExternallyQueried) String() string {
	return "openconfig-mpls-te:externally-queried"
}

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

// TeMetricType represents static value, or to track the IGP metric
type TeMetricType string

const (
    // set the LSP metric to track the underlying
    // IGP metric
    TeMetricType_IGP TeMetricType = "IGP"
)

// TeBandwidthType represents explicitly specified or automatically computed
type TeBandwidthType string

const (
    // Bandwidth is explicitly specified
    TeBandwidthType_SPECIFIED TeBandwidthType = "SPECIFIED"

    // Bandwidth is automatically computed
    TeBandwidthType_AUTO TeBandwidthType = "AUTO"
)

// Mpls
// Anchor point for mpls configuration and operational
// data
// This type is a presence type.
type Mpls struct {
    parent types.Entity
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

func (mpls *Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *Mpls) GetGoName(yname string) string {
    if yname == "global" { return "Global" }
    if yname == "te-global-attributes" { return "TeGlobalAttributes" }
    if yname == "te-interface-attributes" { return "TeInterfaceAttributes" }
    if yname == "signaling-protocols" { return "SignalingProtocols" }
    if yname == "lsps" { return "Lsps" }
    return ""
}

func (mpls *Mpls) GetSegmentPath() string {
    return "openconfig-mpls:mpls"
}

func (mpls *Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global" {
        return &mpls.Global
    }
    if childYangName == "te-global-attributes" {
        return &mpls.TeGlobalAttributes
    }
    if childYangName == "te-interface-attributes" {
        return &mpls.TeInterfaceAttributes
    }
    if childYangName == "signaling-protocols" {
        return &mpls.SignalingProtocols
    }
    if childYangName == "lsps" {
        return &mpls.Lsps
    }
    return nil
}

func (mpls *Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global"] = &mpls.Global
    children["te-global-attributes"] = &mpls.TeGlobalAttributes
    children["te-interface-attributes"] = &mpls.TeInterfaceAttributes
    children["signaling-protocols"] = &mpls.SignalingProtocols
    children["lsps"] = &mpls.Lsps
    return children
}

func (mpls *Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpls *Mpls) GetBundleName() string { return "openconfig" }

func (mpls *Mpls) GetYangName() string { return "mpls" }

func (mpls *Mpls) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (mpls *Mpls) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (mpls *Mpls) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (mpls *Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *Mpls) GetParentYangName() string { return "openconfig-mpls" }

// Mpls_Global
// general mpls configuration applicable to any
// type of LSP and signaling protocol - label ranges,
// entropy label supportmay be added here
type Mpls_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Top level global MPLS configuration.
    Config Mpls_Global_Config

    // Top level global MPLS state.
    State Mpls_Global_State

    // Parameters related to MPLS interfaces.
    MplsInterfaceAttributes Mpls_Global_MplsInterfaceAttributes
}

func (global *Mpls_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Mpls_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Mpls_Global) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "mpls-interface-attributes" { return "MplsInterfaceAttributes" }
    return ""
}

func (global *Mpls_Global) GetSegmentPath() string {
    return "global"
}

func (global *Mpls_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &global.Config
    }
    if childYangName == "state" {
        return &global.State
    }
    if childYangName == "mpls-interface-attributes" {
        return &global.MplsInterfaceAttributes
    }
    return nil
}

func (global *Mpls_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &global.Config
    children["state"] = &global.State
    children["mpls-interface-attributes"] = &global.MplsInterfaceAttributes
    return children
}

func (global *Mpls_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *Mpls_Global) GetBundleName() string { return "openconfig" }

func (global *Mpls_Global) GetYangName() string { return "global" }

func (global *Mpls_Global) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (global *Mpls_Global) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (global *Mpls_Global) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (global *Mpls_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Mpls_Global) GetParent() types.Entity { return global.parent }

func (global *Mpls_Global) GetParentYangName() string { return "mpls" }

// Mpls_Global_Config
// Top level global MPLS configuration
type Mpls_Global_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The null-label type used, implicit or explicit. The type is one of the
    // following: EXPLICITIMPLICIT. The default value is mplst:IMPLICIT.
    NullLabel interface{}
}

func (config *Mpls_Global_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Global_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Global_Config) GetGoName(yname string) string {
    if yname == "null-label" { return "NullLabel" }
    return ""
}

func (config *Mpls_Global_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Global_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Global_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Global_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["null-label"] = config.NullLabel
    return leafs
}

func (config *Mpls_Global_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Global_Config) GetYangName() string { return "config" }

func (config *Mpls_Global_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Global_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Global_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Global_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Global_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Global_Config) GetParentYangName() string { return "global" }

// Mpls_Global_State
// Top level global MPLS state
type Mpls_Global_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The null-label type used, implicit or explicit. The type is one of the
    // following: EXPLICITIMPLICIT. The default value is mplst:IMPLICIT.
    NullLabel interface{}
}

func (state *Mpls_Global_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Global_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Global_State) GetGoName(yname string) string {
    if yname == "null-label" { return "NullLabel" }
    return ""
}

func (state *Mpls_Global_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Global_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Global_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Global_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["null-label"] = state.NullLabel
    return leafs
}

func (state *Mpls_Global_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Global_State) GetYangName() string { return "state" }

func (state *Mpls_Global_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Global_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Global_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Global_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Global_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Global_State) GetParentYangName() string { return "global" }

// Mpls_Global_MplsInterfaceAttributes
// Parameters related to MPLS interfaces
type Mpls_Global_MplsInterfaceAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of TE interfaces. The type is slice of
    // Mpls_Global_MplsInterfaceAttributes_Interface.
    Interface []Mpls_Global_MplsInterfaceAttributes_Interface
}

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetFilter() yfilter.YFilter { return mplsInterfaceAttributes.YFilter }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) SetFilter(yf yfilter.YFilter) { mplsInterfaceAttributes.YFilter = yf }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetSegmentPath() string {
    return "mpls-interface-attributes"
}

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range mplsInterfaceAttributes.Interface {
            if mplsInterfaceAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Global_MplsInterfaceAttributes_Interface{}
        mplsInterfaceAttributes.Interface = append(mplsInterfaceAttributes.Interface, child)
        return &mplsInterfaceAttributes.Interface[len(mplsInterfaceAttributes.Interface)-1]
    }
    return nil
}

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsInterfaceAttributes.Interface {
        children[mplsInterfaceAttributes.Interface[i].GetSegmentPath()] = &mplsInterfaceAttributes.Interface[i]
    }
    return children
}

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetBundleName() string { return "openconfig" }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetYangName() string { return "mpls-interface-attributes" }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) SetParent(parent types.Entity) { mplsInterfaceAttributes.parent = parent }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetParent() types.Entity { return mplsInterfaceAttributes.parent }

func (mplsInterfaceAttributes *Mpls_Global_MplsInterfaceAttributes) GetParentYangName() string { return "global" }

// Mpls_Global_MplsInterfaceAttributes_Interface
// List of TE interfaces
type Mpls_Global_MplsInterfaceAttributes_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string. Refers to
    // mpls.Mpls_Global_MplsInterfaceAttributes_Interface_Config_Name
    Name interface{}

    // Configuration parameters related to MPLS interfaces:.
    Config Mpls_Global_MplsInterfaceAttributes_Interface_Config

    // State parameters related to TE interfaces.
    State Mpls_Global_MplsInterfaceAttributes_Interface_State
}

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &self.Config
    }
    if childYangName == "state" {
        return &self.State
    }
    return nil
}

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &self.Config
    children["state"] = &self.State
    return children
}

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    return leafs
}

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetBundleName() string { return "openconfig" }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetYangName() string { return "interface" }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetParent() types.Entity { return self.parent }

func (self *Mpls_Global_MplsInterfaceAttributes_Interface) GetParentYangName() string { return "mpls-interface-attributes" }

// Mpls_Global_MplsInterfaceAttributes_Interface_Config
// Configuration parameters related to MPLS interfaces:
type Mpls_Global_MplsInterfaceAttributes_Interface_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reference to interface name. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // Enable MPLS forwarding on this interfacek. The type is bool. The default
    // value is false.
    MplsEnabled interface{}
}

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "mpls-enabled" { return "MplsEnabled" }
    return ""
}

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["mpls-enabled"] = config.MplsEnabled
    return leafs
}

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetYangName() string { return "config" }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Global_MplsInterfaceAttributes_Interface_Config) GetParentYangName() string { return "interface" }

// Mpls_Global_MplsInterfaceAttributes_Interface_State
// State parameters related to TE interfaces
type Mpls_Global_MplsInterfaceAttributes_Interface_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reference to interface name. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // Enable MPLS forwarding on this interfacek. The type is bool. The default
    // value is false.
    MplsEnabled interface{}
}

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "mpls-enabled" { return "MplsEnabled" }
    return ""
}

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["mpls-enabled"] = state.MplsEnabled
    return leafs
}

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetYangName() string { return "state" }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Global_MplsInterfaceAttributes_Interface_State) GetParentYangName() string { return "interface" }

// Mpls_TeGlobalAttributes
// traffic-engineering global attributes
type Mpls_TeGlobalAttributes struct {
    parent types.Entity
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

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetFilter() yfilter.YFilter { return teGlobalAttributes.YFilter }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) SetFilter(yf yfilter.YFilter) { teGlobalAttributes.YFilter = yf }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetGoName(yname string) string {
    if yname == "srlg" { return "Srlg" }
    if yname == "igp-flooding-bandwidth" { return "IgpFloodingBandwidth" }
    if yname == "mpls-admin-groups" { return "MplsAdminGroups" }
    if yname == "te-lsp-timers" { return "TeLspTimers" }
    return ""
}

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetSegmentPath() string {
    return "te-global-attributes"
}

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg" {
        return &teGlobalAttributes.Srlg
    }
    if childYangName == "igp-flooding-bandwidth" {
        return &teGlobalAttributes.IgpFloodingBandwidth
    }
    if childYangName == "mpls-admin-groups" {
        return &teGlobalAttributes.MplsAdminGroups
    }
    if childYangName == "te-lsp-timers" {
        return &teGlobalAttributes.TeLspTimers
    }
    return nil
}

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["srlg"] = &teGlobalAttributes.Srlg
    children["igp-flooding-bandwidth"] = &teGlobalAttributes.IgpFloodingBandwidth
    children["mpls-admin-groups"] = &teGlobalAttributes.MplsAdminGroups
    children["te-lsp-timers"] = &teGlobalAttributes.TeLspTimers
    return children
}

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetBundleName() string { return "openconfig" }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetYangName() string { return "te-global-attributes" }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) SetParent(parent types.Entity) { teGlobalAttributes.parent = parent }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetParent() types.Entity { return teGlobalAttributes.parent }

func (teGlobalAttributes *Mpls_TeGlobalAttributes) GetParentYangName() string { return "mpls" }

// Mpls_TeGlobalAttributes_Srlg
// Shared risk link groups attributes
type Mpls_TeGlobalAttributes_Srlg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of shared risk link groups. The type is slice of
    // Mpls_TeGlobalAttributes_Srlg_Srlg.
    Srlg []Mpls_TeGlobalAttributes_Srlg_Srlg
}

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetFilter() yfilter.YFilter { return srlg.YFilter }

func (srlg *Mpls_TeGlobalAttributes_Srlg) SetFilter(yf yfilter.YFilter) { srlg.YFilter = yf }

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetGoName(yname string) string {
    if yname == "srlg" { return "Srlg" }
    return ""
}

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetSegmentPath() string {
    return "srlg"
}

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srlg" {
        for _, c := range srlg.Srlg {
            if srlg.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_TeGlobalAttributes_Srlg_Srlg{}
        srlg.Srlg = append(srlg.Srlg, child)
        return &srlg.Srlg[len(srlg.Srlg)-1]
    }
    return nil
}

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range srlg.Srlg {
        children[srlg.Srlg[i].GetSegmentPath()] = &srlg.Srlg[i]
    }
    return children
}

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetBundleName() string { return "openconfig" }

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetYangName() string { return "srlg" }

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (srlg *Mpls_TeGlobalAttributes_Srlg) SetParent(parent types.Entity) { srlg.parent = parent }

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetParent() types.Entity { return srlg.parent }

func (srlg *Mpls_TeGlobalAttributes_Srlg) GetParentYangName() string { return "te-global-attributes" }

// Mpls_TeGlobalAttributes_Srlg_Srlg
// List of shared risk link groups
type Mpls_TeGlobalAttributes_Srlg_Srlg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The SRLG group identifier. The type is string.
    // Refers to mpls.Mpls_TeGlobalAttributes_Srlg_Srlg_Config_Name
    Name interface{}

    // Configuration parameters related to the SRLG.
    Config Mpls_TeGlobalAttributes_Srlg_Srlg_Config

    // State parameters related to the SRLG.
    State Mpls_TeGlobalAttributes_Srlg_Srlg_State

    // SRLG members for static (not flooded) SRLGs .
    StaticSrlgMembers Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers
}

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetFilter() yfilter.YFilter { return srlg.YFilter }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) SetFilter(yf yfilter.YFilter) { srlg.YFilter = yf }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "static-srlg-members" { return "StaticSrlgMembers" }
    return ""
}

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetSegmentPath() string {
    return "srlg" + "[name='" + fmt.Sprintf("%v", srlg.Name) + "']"
}

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &srlg.Config
    }
    if childYangName == "state" {
        return &srlg.State
    }
    if childYangName == "static-srlg-members" {
        return &srlg.StaticSrlgMembers
    }
    return nil
}

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &srlg.Config
    children["state"] = &srlg.State
    children["static-srlg-members"] = &srlg.StaticSrlgMembers
    return children
}

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = srlg.Name
    return leafs
}

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetBundleName() string { return "openconfig" }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetYangName() string { return "srlg" }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) SetParent(parent types.Entity) { srlg.parent = parent }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetParent() types.Entity { return srlg.parent }

func (srlg *Mpls_TeGlobalAttributes_Srlg_Srlg) GetParentYangName() string { return "srlg" }

// Mpls_TeGlobalAttributes_Srlg_Srlg_Config
// Configuration parameters related to the SRLG
type Mpls_TeGlobalAttributes_Srlg_Srlg_Config struct {
    parent types.Entity
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

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "value" { return "Value" }
    if yname == "cost" { return "Cost" }
    if yname == "flooding-type" { return "FloodingType" }
    return ""
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["value"] = config.Value
    leafs["cost"] = config.Cost
    leafs["flooding-type"] = config.FloodingType
    return leafs
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetYangName() string { return "config" }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_Config) GetParentYangName() string { return "srlg" }

// Mpls_TeGlobalAttributes_Srlg_Srlg_State
// State parameters related to the SRLG
type Mpls_TeGlobalAttributes_Srlg_Srlg_State struct {
    parent types.Entity
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

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "value" { return "Value" }
    if yname == "cost" { return "Cost" }
    if yname == "flooding-type" { return "FloodingType" }
    return ""
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["value"] = state.Value
    leafs["cost"] = state.Cost
    leafs["flooding-type"] = state.FloodingType
    return leafs
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetYangName() string { return "state" }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_State) GetParentYangName() string { return "srlg" }

// Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers
// SRLG members for static (not flooded) SRLGs 
type Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of SRLG members, which are expressed as IP address endpoints of links
    // contained in the SRLG. The type is slice of
    // Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList.
    MembersList []Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList
}

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetFilter() yfilter.YFilter { return staticSrlgMembers.YFilter }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) SetFilter(yf yfilter.YFilter) { staticSrlgMembers.YFilter = yf }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetGoName(yname string) string {
    if yname == "members-list" { return "MembersList" }
    return ""
}

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetSegmentPath() string {
    return "static-srlg-members"
}

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "members-list" {
        for _, c := range staticSrlgMembers.MembersList {
            if staticSrlgMembers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList{}
        staticSrlgMembers.MembersList = append(staticSrlgMembers.MembersList, child)
        return &staticSrlgMembers.MembersList[len(staticSrlgMembers.MembersList)-1]
    }
    return nil
}

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticSrlgMembers.MembersList {
        children[staticSrlgMembers.MembersList[i].GetSegmentPath()] = &staticSrlgMembers.MembersList[i]
    }
    return children
}

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetBundleName() string { return "openconfig" }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetYangName() string { return "static-srlg-members" }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) SetParent(parent types.Entity) { staticSrlgMembers.parent = parent }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetParent() types.Entity { return staticSrlgMembers.parent }

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers) GetParentYangName() string { return "srlg" }

// Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList
// List of SRLG members, which are expressed
// as IP address endpoints of links contained in the
// SRLG
type Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The from address of the link in the SRLG. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    FromAddress interface{}

    // Configuration parameters relating to the SRLG members.
    Config Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config

    // State parameters relating to the SRLG members.
    State Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State
}

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetFilter() yfilter.YFilter { return membersList.YFilter }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) SetFilter(yf yfilter.YFilter) { membersList.YFilter = yf }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetGoName(yname string) string {
    if yname == "from-address" { return "FromAddress" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetSegmentPath() string {
    return "members-list" + "[from-address='" + fmt.Sprintf("%v", membersList.FromAddress) + "']"
}

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &membersList.Config
    }
    if childYangName == "state" {
        return &membersList.State
    }
    return nil
}

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &membersList.Config
    children["state"] = &membersList.State
    return children
}

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["from-address"] = membersList.FromAddress
    return leafs
}

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetBundleName() string { return "openconfig" }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetYangName() string { return "members-list" }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) SetParent(parent types.Entity) { membersList.parent = parent }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetParent() types.Entity { return membersList.parent }

func (membersList *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList) GetParentYangName() string { return "static-srlg-members" }

// Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config
// Configuration parameters relating to the
// SRLG members
type Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of the a-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    FromAddress interface{}

    // IP address of the z-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ToAddress interface{}
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetGoName(yname string) string {
    if yname == "from-address" { return "FromAddress" }
    if yname == "to-address" { return "ToAddress" }
    return ""
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["from-address"] = config.FromAddress
    leafs["to-address"] = config.ToAddress
    return leafs
}

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetYangName() string { return "config" }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_Config) GetParentYangName() string { return "members-list" }

// Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State
// State parameters relating to the SRLG
// members
type Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of the a-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    FromAddress interface{}

    // IP address of the z-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ToAddress interface{}
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetGoName(yname string) string {
    if yname == "from-address" { return "FromAddress" }
    if yname == "to-address" { return "ToAddress" }
    return ""
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["from-address"] = state.FromAddress
    leafs["to-address"] = state.ToAddress
    return leafs
}

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetYangName() string { return "state" }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_TeGlobalAttributes_Srlg_Srlg_StaticSrlgMembers_MembersList_State) GetParentYangName() string { return "members-list" }

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth
// Interface bandwidth change percentages
// that trigger update events into the IGP traffic
// engineering database (TED)
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters for TED update threshold .
    Config Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config

    // State parameters for TED update threshold .
    State Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State
}

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetFilter() yfilter.YFilter { return igpFloodingBandwidth.YFilter }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) SetFilter(yf yfilter.YFilter) { igpFloodingBandwidth.YFilter = yf }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetSegmentPath() string {
    return "igp-flooding-bandwidth"
}

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &igpFloodingBandwidth.Config
    }
    if childYangName == "state" {
        return &igpFloodingBandwidth.State
    }
    return nil
}

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &igpFloodingBandwidth.Config
    children["state"] = &igpFloodingBandwidth.State
    return children
}

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetBundleName() string { return "openconfig" }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetYangName() string { return "igp-flooding-bandwidth" }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) SetParent(parent types.Entity) { igpFloodingBandwidth.parent = parent }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetParent() types.Entity { return igpFloodingBandwidth.parent }

func (igpFloodingBandwidth *Mpls_TeGlobalAttributes_IgpFloodingBandwidth) GetParentYangName() string { return "te-global-attributes" }

// Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config
// Configuration parameters for TED
// update threshold 
type Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config struct {
    parent types.Entity
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

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetGoName(yname string) string {
    if yname == "threshold-type" { return "ThresholdType" }
    if yname == "delta-percentage" { return "DeltaPercentage" }
    if yname == "threshold-specification" { return "ThresholdSpecification" }
    if yname == "up-thresholds" { return "UpThresholds" }
    if yname == "down-thresholds" { return "DownThresholds" }
    if yname == "up-down-thresholds" { return "UpDownThresholds" }
    return ""
}

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-type"] = config.ThresholdType
    leafs["delta-percentage"] = config.DeltaPercentage
    leafs["threshold-specification"] = config.ThresholdSpecification
    leafs["up-thresholds"] = config.UpThresholds
    leafs["down-thresholds"] = config.DownThresholds
    leafs["up-down-thresholds"] = config.UpDownThresholds
    return leafs
}

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetYangName() string { return "config" }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_Config) GetParentYangName() string { return "igp-flooding-bandwidth" }

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
    parent types.Entity
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

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetGoName(yname string) string {
    if yname == "threshold-type" { return "ThresholdType" }
    if yname == "delta-percentage" { return "DeltaPercentage" }
    if yname == "threshold-specification" { return "ThresholdSpecification" }
    if yname == "up-thresholds" { return "UpThresholds" }
    if yname == "down-thresholds" { return "DownThresholds" }
    if yname == "up-down-thresholds" { return "UpDownThresholds" }
    return ""
}

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-type"] = state.ThresholdType
    leafs["delta-percentage"] = state.DeltaPercentage
    leafs["threshold-specification"] = state.ThresholdSpecification
    leafs["up-thresholds"] = state.UpThresholds
    leafs["down-thresholds"] = state.DownThresholds
    leafs["up-down-thresholds"] = state.UpDownThresholds
    return leafs
}

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetYangName() string { return "state" }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_TeGlobalAttributes_IgpFloodingBandwidth_State) GetParentYangName() string { return "igp-flooding-bandwidth" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // configuration of value to name mapping for mpls affinities/admin-groups.
    // The type is slice of Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup.
    AdminGroup []Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup
}

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetFilter() yfilter.YFilter { return mplsAdminGroups.YFilter }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) SetFilter(yf yfilter.YFilter) { mplsAdminGroups.YFilter = yf }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetGoName(yname string) string {
    if yname == "admin-group" { return "AdminGroup" }
    return ""
}

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetSegmentPath() string {
    return "mpls-admin-groups"
}

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "admin-group" {
        for _, c := range mplsAdminGroups.AdminGroup {
            if mplsAdminGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup{}
        mplsAdminGroups.AdminGroup = append(mplsAdminGroups.AdminGroup, child)
        return &mplsAdminGroups.AdminGroup[len(mplsAdminGroups.AdminGroup)-1]
    }
    return nil
}

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mplsAdminGroups.AdminGroup {
        children[mplsAdminGroups.AdminGroup[i].GetSegmentPath()] = &mplsAdminGroups.AdminGroup[i]
    }
    return children
}

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetBundleName() string { return "openconfig" }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetYangName() string { return "mpls-admin-groups" }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) SetParent(parent types.Entity) { mplsAdminGroups.parent = parent }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetParent() types.Entity { return mplsAdminGroups.parent }

func (mplsAdminGroups *Mpls_TeGlobalAttributes_MplsAdminGroups) GetParentYangName() string { return "te-global-attributes" }

// Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup
// configuration of value to name mapping
// for mpls affinities/admin-groups
type Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup struct {
    parent types.Entity
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

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetFilter() yfilter.YFilter { return adminGroup.YFilter }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) SetFilter(yf yfilter.YFilter) { adminGroup.YFilter = yf }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetGoName(yname string) string {
    if yname == "admin-group-name" { return "AdminGroupName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetSegmentPath() string {
    return "admin-group" + "[admin-group-name='" + fmt.Sprintf("%v", adminGroup.AdminGroupName) + "']"
}

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &adminGroup.Config
    }
    if childYangName == "state" {
        return &adminGroup.State
    }
    return nil
}

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &adminGroup.Config
    children["state"] = &adminGroup.State
    return children
}

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-group-name"] = adminGroup.AdminGroupName
    return leafs
}

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetBundleName() string { return "openconfig" }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetYangName() string { return "admin-group" }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) SetParent(parent types.Entity) { adminGroup.parent = parent }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetParent() types.Entity { return adminGroup.parent }

func (adminGroup *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup) GetParentYangName() string { return "mpls-admin-groups" }

// Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config
// Configurable items for admin-groups
type Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config struct {
    parent types.Entity
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

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetGoName(yname string) string {
    if yname == "admin-group-name" { return "AdminGroupName" }
    if yname == "bit-position" { return "BitPosition" }
    return ""
}

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-group-name"] = config.AdminGroupName
    leafs["bit-position"] = config.BitPosition
    return leafs
}

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetYangName() string { return "config" }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_Config) GetParentYangName() string { return "admin-group" }

// Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State
// Operational state for admin-groups
type Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State struct {
    parent types.Entity
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

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetGoName(yname string) string {
    if yname == "admin-group-name" { return "AdminGroupName" }
    if yname == "bit-position" { return "BitPosition" }
    return ""
}

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-group-name"] = state.AdminGroupName
    leafs["bit-position"] = state.BitPosition
    return leafs
}

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetYangName() string { return "state" }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup_State) GetParentYangName() string { return "admin-group" }

// Mpls_TeGlobalAttributes_TeLspTimers
// Definition for delays associated with setup
// and cleanup of TE LSPs
type Mpls_TeGlobalAttributes_TeLspTimers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters related to timers for TE LSPs.
    Config Mpls_TeGlobalAttributes_TeLspTimers_Config

    // State related to timers for TE LSPs.
    State Mpls_TeGlobalAttributes_TeLspTimers_State
}

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetFilter() yfilter.YFilter { return teLspTimers.YFilter }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) SetFilter(yf yfilter.YFilter) { teLspTimers.YFilter = yf }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetSegmentPath() string {
    return "te-lsp-timers"
}

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &teLspTimers.Config
    }
    if childYangName == "state" {
        return &teLspTimers.State
    }
    return nil
}

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &teLspTimers.Config
    children["state"] = &teLspTimers.State
    return children
}

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetBundleName() string { return "openconfig" }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetYangName() string { return "te-lsp-timers" }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) SetParent(parent types.Entity) { teLspTimers.parent = parent }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetParent() types.Entity { return teLspTimers.parent }

func (teLspTimers *Mpls_TeGlobalAttributes_TeLspTimers) GetParentYangName() string { return "te-global-attributes" }

// Mpls_TeGlobalAttributes_TeLspTimers_Config
// Configuration parameters related
// to timers for TE LSPs
type Mpls_TeGlobalAttributes_TeLspTimers_Config struct {
    parent types.Entity
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

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetGoName(yname string) string {
    if yname == "install-delay" { return "InstallDelay" }
    if yname == "cleanup-delay" { return "CleanupDelay" }
    if yname == "reoptimize-timer" { return "ReoptimizeTimer" }
    return ""
}

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["install-delay"] = config.InstallDelay
    leafs["cleanup-delay"] = config.CleanupDelay
    leafs["reoptimize-timer"] = config.ReoptimizeTimer
    return leafs
}

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetYangName() string { return "config" }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_TeGlobalAttributes_TeLspTimers_Config) GetParentYangName() string { return "te-lsp-timers" }

// Mpls_TeGlobalAttributes_TeLspTimers_State
// State related to timers for TE LSPs
type Mpls_TeGlobalAttributes_TeLspTimers_State struct {
    parent types.Entity
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

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetGoName(yname string) string {
    if yname == "install-delay" { return "InstallDelay" }
    if yname == "cleanup-delay" { return "CleanupDelay" }
    if yname == "reoptimize-timer" { return "ReoptimizeTimer" }
    return ""
}

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["install-delay"] = state.InstallDelay
    leafs["cleanup-delay"] = state.CleanupDelay
    leafs["reoptimize-timer"] = state.ReoptimizeTimer
    return leafs
}

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetYangName() string { return "state" }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_TeGlobalAttributes_TeLspTimers_State) GetParentYangName() string { return "te-lsp-timers" }

// Mpls_TeInterfaceAttributes
// traffic engineering attributes specific
// for interfaces
type Mpls_TeInterfaceAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of TE interfaces. The type is slice of
    // Mpls_TeInterfaceAttributes_Interface.
    Interface []Mpls_TeInterfaceAttributes_Interface
}

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetFilter() yfilter.YFilter { return teInterfaceAttributes.YFilter }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) SetFilter(yf yfilter.YFilter) { teInterfaceAttributes.YFilter = yf }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetSegmentPath() string {
    return "te-interface-attributes"
}

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range teInterfaceAttributes.Interface {
            if teInterfaceAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_TeInterfaceAttributes_Interface{}
        teInterfaceAttributes.Interface = append(teInterfaceAttributes.Interface, child)
        return &teInterfaceAttributes.Interface[len(teInterfaceAttributes.Interface)-1]
    }
    return nil
}

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range teInterfaceAttributes.Interface {
        children[teInterfaceAttributes.Interface[i].GetSegmentPath()] = &teInterfaceAttributes.Interface[i]
    }
    return children
}

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetBundleName() string { return "openconfig" }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetYangName() string { return "te-interface-attributes" }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) SetParent(parent types.Entity) { teInterfaceAttributes.parent = parent }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetParent() types.Entity { return teInterfaceAttributes.parent }

func (teInterfaceAttributes *Mpls_TeInterfaceAttributes) GetParentYangName() string { return "mpls" }

// Mpls_TeInterfaceAttributes_Interface
// List of TE interfaces
type Mpls_TeInterfaceAttributes_Interface struct {
    parent types.Entity
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

func (self *Mpls_TeInterfaceAttributes_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mpls_TeInterfaceAttributes_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mpls_TeInterfaceAttributes_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "igp-flooding-bandwidth" { return "IgpFloodingBandwidth" }
    return ""
}

func (self *Mpls_TeInterfaceAttributes_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Mpls_TeInterfaceAttributes_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &self.Config
    }
    if childYangName == "state" {
        return &self.State
    }
    if childYangName == "igp-flooding-bandwidth" {
        return &self.IgpFloodingBandwidth
    }
    return nil
}

func (self *Mpls_TeInterfaceAttributes_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &self.Config
    children["state"] = &self.State
    children["igp-flooding-bandwidth"] = &self.IgpFloodingBandwidth
    return children
}

func (self *Mpls_TeInterfaceAttributes_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    return leafs
}

func (self *Mpls_TeInterfaceAttributes_Interface) GetBundleName() string { return "openconfig" }

func (self *Mpls_TeInterfaceAttributes_Interface) GetYangName() string { return "interface" }

func (self *Mpls_TeInterfaceAttributes_Interface) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (self *Mpls_TeInterfaceAttributes_Interface) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (self *Mpls_TeInterfaceAttributes_Interface) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (self *Mpls_TeInterfaceAttributes_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mpls_TeInterfaceAttributes_Interface) GetParent() types.Entity { return self.parent }

func (self *Mpls_TeInterfaceAttributes_Interface) GetParentYangName() string { return "te-interface-attributes" }

// Mpls_TeInterfaceAttributes_Interface_Config
// Configuration parameters related to TE interfaces:
type Mpls_TeInterfaceAttributes_Interface_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reference to interface name. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // TE specific metric for the link. The type is interface{} with range:
    // 0..4294967295.
    TeMetric interface{}

    // list of references to named shared risk link groups that the interface
    // belongs to. The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_Srlg_Srlg_Name
    SrlgMembership []interface{}

    // list of admin groups (by name) on the interface. The type is slice of
    // string.
    AdminGroup []interface{}
}

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "srlg-membership" { return "SrlgMembership" }
    if yname == "admin-group" { return "AdminGroup" }
    return ""
}

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["te-metric"] = config.TeMetric
    leafs["srlg-membership"] = config.SrlgMembership
    leafs["admin-group"] = config.AdminGroup
    return leafs
}

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetYangName() string { return "config" }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_TeInterfaceAttributes_Interface_Config) GetParentYangName() string { return "interface" }

// Mpls_TeInterfaceAttributes_Interface_State
// State parameters related to TE interfaces
type Mpls_TeInterfaceAttributes_Interface_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // reference to interface name. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // TE specific metric for the link. The type is interface{} with range:
    // 0..4294967295.
    TeMetric interface{}

    // list of references to named shared risk link groups that the interface
    // belongs to. The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_Srlg_Srlg_Name
    SrlgMembership []interface{}

    // list of admin groups (by name) on the interface. The type is slice of
    // string.
    AdminGroup []interface{}
}

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_TeInterfaceAttributes_Interface_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "srlg-membership" { return "SrlgMembership" }
    if yname == "admin-group" { return "AdminGroup" }
    return ""
}

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["te-metric"] = state.TeMetric
    leafs["srlg-membership"] = state.SrlgMembership
    leafs["admin-group"] = state.AdminGroup
    return leafs
}

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetYangName() string { return "state" }

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_TeInterfaceAttributes_Interface_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_TeInterfaceAttributes_Interface_State) GetParentYangName() string { return "interface" }

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth
// Interface bandwidth change percentages
// that trigger update events into the IGP traffic
// engineering database (TED)
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters for TED update threshold .
    Config Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config

    // State parameters for TED update threshold .
    State Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State
}

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetFilter() yfilter.YFilter { return igpFloodingBandwidth.YFilter }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) SetFilter(yf yfilter.YFilter) { igpFloodingBandwidth.YFilter = yf }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetSegmentPath() string {
    return "igp-flooding-bandwidth"
}

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &igpFloodingBandwidth.Config
    }
    if childYangName == "state" {
        return &igpFloodingBandwidth.State
    }
    return nil
}

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &igpFloodingBandwidth.Config
    children["state"] = &igpFloodingBandwidth.State
    return children
}

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetBundleName() string { return "openconfig" }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetYangName() string { return "igp-flooding-bandwidth" }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) SetParent(parent types.Entity) { igpFloodingBandwidth.parent = parent }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetParent() types.Entity { return igpFloodingBandwidth.parent }

func (igpFloodingBandwidth *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth) GetParentYangName() string { return "interface" }

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config
// Configuration parameters for TED
// update threshold 
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config struct {
    parent types.Entity
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

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetGoName(yname string) string {
    if yname == "threshold-type" { return "ThresholdType" }
    if yname == "delta-percentage" { return "DeltaPercentage" }
    if yname == "threshold-specification" { return "ThresholdSpecification" }
    if yname == "up-thresholds" { return "UpThresholds" }
    if yname == "down-thresholds" { return "DownThresholds" }
    if yname == "up-down-thresholds" { return "UpDownThresholds" }
    return ""
}

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-type"] = config.ThresholdType
    leafs["delta-percentage"] = config.DeltaPercentage
    leafs["threshold-specification"] = config.ThresholdSpecification
    leafs["up-thresholds"] = config.UpThresholds
    leafs["down-thresholds"] = config.DownThresholds
    leafs["up-down-thresholds"] = config.UpDownThresholds
    return leafs
}

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetYangName() string { return "config" }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config) GetParentYangName() string { return "igp-flooding-bandwidth" }

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
    parent types.Entity
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

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetGoName(yname string) string {
    if yname == "threshold-type" { return "ThresholdType" }
    if yname == "delta-percentage" { return "DeltaPercentage" }
    if yname == "threshold-specification" { return "ThresholdSpecification" }
    if yname == "up-thresholds" { return "UpThresholds" }
    if yname == "down-thresholds" { return "DownThresholds" }
    if yname == "up-down-thresholds" { return "UpDownThresholds" }
    return ""
}

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-type"] = state.ThresholdType
    leafs["delta-percentage"] = state.DeltaPercentage
    leafs["threshold-specification"] = state.ThresholdSpecification
    leafs["up-thresholds"] = state.UpThresholds
    leafs["down-thresholds"] = state.DownThresholds
    leafs["up-down-thresholds"] = state.UpDownThresholds
    return leafs
}

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetYangName() string { return "state" }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State) GetParentYangName() string { return "igp-flooding-bandwidth" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP-TE global signaling protocol configuration.
    RsvpTe Mpls_SignalingProtocols_RsvpTe

    // SR global signaling config.
    SegmentRouting Mpls_SignalingProtocols_SegmentRouting

    // LDP global signaling configuration.
    Ldp Mpls_SignalingProtocols_Ldp
}

func (signalingProtocols *Mpls_SignalingProtocols) GetFilter() yfilter.YFilter { return signalingProtocols.YFilter }

func (signalingProtocols *Mpls_SignalingProtocols) SetFilter(yf yfilter.YFilter) { signalingProtocols.YFilter = yf }

func (signalingProtocols *Mpls_SignalingProtocols) GetGoName(yname string) string {
    if yname == "rsvp-te" { return "RsvpTe" }
    if yname == "segment-routing" { return "SegmentRouting" }
    if yname == "ldp" { return "Ldp" }
    return ""
}

func (signalingProtocols *Mpls_SignalingProtocols) GetSegmentPath() string {
    return "signaling-protocols"
}

func (signalingProtocols *Mpls_SignalingProtocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvp-te" {
        return &signalingProtocols.RsvpTe
    }
    if childYangName == "segment-routing" {
        return &signalingProtocols.SegmentRouting
    }
    if childYangName == "ldp" {
        return &signalingProtocols.Ldp
    }
    return nil
}

func (signalingProtocols *Mpls_SignalingProtocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rsvp-te"] = &signalingProtocols.RsvpTe
    children["segment-routing"] = &signalingProtocols.SegmentRouting
    children["ldp"] = &signalingProtocols.Ldp
    return children
}

func (signalingProtocols *Mpls_SignalingProtocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (signalingProtocols *Mpls_SignalingProtocols) GetBundleName() string { return "openconfig" }

func (signalingProtocols *Mpls_SignalingProtocols) GetYangName() string { return "signaling-protocols" }

func (signalingProtocols *Mpls_SignalingProtocols) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (signalingProtocols *Mpls_SignalingProtocols) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (signalingProtocols *Mpls_SignalingProtocols) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (signalingProtocols *Mpls_SignalingProtocols) SetParent(parent types.Entity) { signalingProtocols.parent = parent }

func (signalingProtocols *Mpls_SignalingProtocols) GetParent() types.Entity { return signalingProtocols.parent }

func (signalingProtocols *Mpls_SignalingProtocols) GetParentYangName() string { return "mpls" }

// Mpls_SignalingProtocols_RsvpTe
// RSVP-TE global signaling protocol configuration
type Mpls_SignalingProtocols_RsvpTe struct {
    parent types.Entity
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

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetFilter() yfilter.YFilter { return rsvpTe.YFilter }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) SetFilter(yf yfilter.YFilter) { rsvpTe.YFilter = yf }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "global" { return "Global" }
    if yname == "interface-attributes" { return "InterfaceAttributes" }
    return ""
}

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetSegmentPath() string {
    return "rsvp-te"
}

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &rsvpTe.Sessions
    }
    if childYangName == "neighbors" {
        return &rsvpTe.Neighbors
    }
    if childYangName == "global" {
        return &rsvpTe.Global
    }
    if childYangName == "interface-attributes" {
        return &rsvpTe.InterfaceAttributes
    }
    return nil
}

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &rsvpTe.Sessions
    children["neighbors"] = &rsvpTe.Neighbors
    children["global"] = &rsvpTe.Global
    children["interface-attributes"] = &rsvpTe.InterfaceAttributes
    return children
}

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetBundleName() string { return "openconfig" }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetYangName() string { return "rsvp-te" }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) SetParent(parent types.Entity) { rsvpTe.parent = parent }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetParent() types.Entity { return rsvpTe.parent }

func (rsvpTe *Mpls_SignalingProtocols_RsvpTe) GetParentYangName() string { return "signaling-protocols" }

// Mpls_SignalingProtocols_RsvpTe_Sessions
// Configuration and state of RSVP sessions
type Mpls_SignalingProtocols_RsvpTe_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of RSVP sessions on the device.
    Config Mpls_SignalingProtocols_RsvpTe_Sessions_Config

    // State information relating to RSVP sessions on the device.
    State Mpls_SignalingProtocols_RsvpTe_Sessions_State
}

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &sessions.Config
    }
    if childYangName == "state" {
        return &sessions.State
    }
    return nil
}

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &sessions.Config
    children["state"] = &sessions.State
    return children
}

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetBundleName() string { return "openconfig" }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetYangName() string { return "sessions" }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *Mpls_SignalingProtocols_RsvpTe_Sessions) GetParentYangName() string { return "rsvp-te" }

// Mpls_SignalingProtocols_RsvpTe_Sessions_Config
// Configuration of RSVP sessions on the device
type Mpls_SignalingProtocols_RsvpTe_Sessions_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetGoName(yname string) string {
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Sessions_Config) GetParentYangName() string { return "sessions" }

// Mpls_SignalingProtocols_RsvpTe_Sessions_State
// State information relating to RSVP sessions
// on the device
type Mpls_SignalingProtocols_RsvpTe_Sessions_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of RSVP sessions. The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session.
    Session []Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range state.Session {
            if state.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session{}
        state.Session = append(state.Session, child)
        return &state.Session[len(state.Session)-1]
    }
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range state.Session {
        children[state.Session[i].GetSegmentPath()] = &state.Session[i]
    }
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_State) GetParentYangName() string { return "sessions" }

// Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session
// List of RSVP sessions
type Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RSVP source port. The type is interface{} with
    // range: 0..65535.
    SourcePort interface{}

    // This attribute is a key. RSVP source port. The type is interface{} with
    // range: 0..65535.
    DestinationPort interface{}

    // This attribute is a key. Origin address of RSVP session. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Destination address of RSVP session. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Enumeration of RSVP session states. The type is Status.
    Status interface{}

    // Enumeration of possible RSVP session types. The type is Type.
    Type interface{}

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

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetGoName(yname string) string {
    if yname == "source-port" { return "SourcePort" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "status" { return "Status" }
    if yname == "type" { return "Type" }
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "label-in" { return "LabelIn" }
    if yname == "label-out" { return "LabelOut" }
    if yname == "associated-lsps" { return "AssociatedLsps" }
    return ""
}

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetSegmentPath() string {
    return "session" + "[source-port='" + fmt.Sprintf("%v", session.SourcePort) + "']" + "[destination-port='" + fmt.Sprintf("%v", session.DestinationPort) + "']" + "[source-address='" + fmt.Sprintf("%v", session.SourceAddress) + "']" + "[destination-address='" + fmt.Sprintf("%v", session.DestinationAddress) + "']"
}

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-port"] = session.SourcePort
    leafs["destination-port"] = session.DestinationPort
    leafs["source-address"] = session.SourceAddress
    leafs["destination-address"] = session.DestinationAddress
    leafs["status"] = session.Status
    leafs["type"] = session.Type
    leafs["tunnel-id"] = session.TunnelId
    leafs["label-in"] = session.LabelIn
    leafs["label-out"] = session.LabelOut
    leafs["associated-lsps"] = session.AssociatedLsps
    return leafs
}

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetBundleName() string { return "openconfig" }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetYangName() string { return "session" }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetParent() types.Entity { return session.parent }

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session) GetParentYangName() string { return "state" }

// Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status represents Enumeration of RSVP session states
type Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status string

const (
    // RSVP session is up
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status_UP Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status = "UP"

    // RSVP session is down
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status_DOWN Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Status = "DOWN"
)

// Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type represents Enumeration of possible RSVP session types
type Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type string

const (
    // RSVP session originates on this device
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type_SOURCE Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type = "SOURCE"

    // RSVP session transits this device only
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type_TRANSIT Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type = "TRANSIT"

    // RSVP session terminates on this device
    Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type_DESTINATION Mpls_SignalingProtocols_RsvpTe_Sessions_State_Session_Type = "DESTINATION"
)

// Mpls_SignalingProtocols_RsvpTe_Neighbors
// Configuration and state for RSVP neighbors connecting
// to the device
type Mpls_SignalingProtocols_RsvpTe_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of RSVP neighbor information.
    Config Mpls_SignalingProtocols_RsvpTe_Neighbors_Config

    // State information relating to RSVP neighbors.
    State Mpls_SignalingProtocols_RsvpTe_Neighbors_State
}

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &neighbors.Config
    }
    if childYangName == "state" {
        return &neighbors.State
    }
    return nil
}

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &neighbors.Config
    children["state"] = &neighbors.State
    return children
}

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetBundleName() string { return "openconfig" }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Mpls_SignalingProtocols_RsvpTe_Neighbors) GetParentYangName() string { return "rsvp-te" }

// Mpls_SignalingProtocols_RsvpTe_Neighbors_Config
// Configuration of RSVP neighbor information
type Mpls_SignalingProtocols_RsvpTe_Neighbors_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetGoName(yname string) string {
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Neighbors_Config) GetParentYangName() string { return "neighbors" }

// Mpls_SignalingProtocols_RsvpTe_Neighbors_State
// State information relating to RSVP neighbors
type Mpls_SignalingProtocols_RsvpTe_Neighbors_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of RSVP neighbors connecting to the device, keyed by neighbor address.
    // The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor.
    Neighbor []Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor
}

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range state.Neighbor {
            if state.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor{}
        state.Neighbor = append(state.Neighbor, child)
        return &state.Neighbor[len(state.Neighbor)-1]
    }
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range state.Neighbor {
        children[state.Neighbor[i].GetSegmentPath()] = &state.Neighbor[i]
    }
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_State) GetParentYangName() string { return "neighbors" }

// Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor
// List of RSVP neighbors connecting to the device,
// keyed by neighbor address
type Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address of RSVP neighbor. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface where RSVP neighbor was detected. The type is string.
    DetectedInterface interface{}

    // Enumuration of possible RSVP neighbor states. The type is NeighborStatus.
    NeighborStatus interface{}

    // Suppport of neighbor for RSVP refresh reduction. The type is bool.
    RefreshReduction interface{}
}

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "detected-interface" { return "DetectedInterface" }
    if yname == "neighbor-status" { return "NeighborStatus" }
    if yname == "refresh-reduction" { return "RefreshReduction" }
    return ""
}

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[address='" + fmt.Sprintf("%v", neighbor.Address) + "']"
}

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = neighbor.Address
    leafs["detected-interface"] = neighbor.DetectedInterface
    leafs["neighbor-status"] = neighbor.NeighborStatus
    leafs["refresh-reduction"] = neighbor.RefreshReduction
    return leafs
}

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetBundleName() string { return "openconfig" }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_State_Neighbor) GetParentYangName() string { return "state" }

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
    parent types.Entity
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

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetGoName(yname string) string {
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "soft-preemption" { return "SoftPreemption" }
    if yname == "hellos" { return "Hellos" }
    if yname == "state" { return "State" }
    return ""
}

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetSegmentPath() string {
    return "global"
}

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "graceful-restart" {
        return &global.GracefulRestart
    }
    if childYangName == "soft-preemption" {
        return &global.SoftPreemption
    }
    if childYangName == "hellos" {
        return &global.Hellos
    }
    if childYangName == "state" {
        return &global.State
    }
    return nil
}

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["graceful-restart"] = &global.GracefulRestart
    children["soft-preemption"] = &global.SoftPreemption
    children["hellos"] = &global.Hellos
    children["state"] = &global.State
    return children
}

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetBundleName() string { return "openconfig" }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetYangName() string { return "global" }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetParent() types.Entity { return global.parent }

func (global *Mpls_SignalingProtocols_RsvpTe_Global) GetParentYangName() string { return "rsvp-te" }

// Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart
// Operational state and configuration parameters relating to
// graceful-restart for RSVP
type Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to graceful-restart.
    Config Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config

    // State information associated with RSVP graceful-restart.
    State Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State
}

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &gracefulRestart.Config
    }
    if childYangName == "state" {
        return &gracefulRestart.State
    }
    return nil
}

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &gracefulRestart.Config
    children["state"] = &gracefulRestart.State
    return children
}

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetBundleName() string { return "openconfig" }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart) GetParentYangName() string { return "global" }

// Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config
// Configuration parameters relating to
// graceful-restart
type Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config struct {
    parent types.Entity
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

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "recovery-time" { return "RecoveryTime" }
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = config.Enable
    leafs["restart-time"] = config.RestartTime
    leafs["recovery-time"] = config.RecoveryTime
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_Config) GetParentYangName() string { return "graceful-restart" }

// Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State
// State information associated with
// RSVP graceful-restart
type Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State struct {
    parent types.Entity
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

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "restart-time" { return "RestartTime" }
    if yname == "recovery-time" { return "RecoveryTime" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = state.Enable
    leafs["restart-time"] = state.RestartTime
    leafs["recovery-time"] = state.RecoveryTime
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_GracefulRestart_State) GetParentYangName() string { return "graceful-restart" }

// Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption
// Protocol options relating to RSVP
// soft preemption
type Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to RSVP soft preemption support.
    Config Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config

    // State parameters relating to RSVP soft preemption support.
    State Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State
}

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetFilter() yfilter.YFilter { return softPreemption.YFilter }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) SetFilter(yf yfilter.YFilter) { softPreemption.YFilter = yf }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetSegmentPath() string {
    return "soft-preemption"
}

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &softPreemption.Config
    }
    if childYangName == "state" {
        return &softPreemption.State
    }
    return nil
}

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &softPreemption.Config
    children["state"] = &softPreemption.State
    return children
}

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetBundleName() string { return "openconfig" }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetYangName() string { return "soft-preemption" }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) SetParent(parent types.Entity) { softPreemption.parent = parent }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetParent() types.Entity { return softPreemption.parent }

func (softPreemption *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption) GetParentYangName() string { return "global" }

// Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config
// Configuration parameters relating to RSVP
// soft preemption support
type Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enables soft preemption on a node. The type is bool. The default value is
    // false.
    Enable interface{}

    // Timeout value for soft preemption to revert to hard preemption. The type is
    // interface{} with range: 0..65535. The default value is 0.
    SoftPreemptionTimeout interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "soft-preemption-timeout" { return "SoftPreemptionTimeout" }
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = config.Enable
    leafs["soft-preemption-timeout"] = config.SoftPreemptionTimeout
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_Config) GetParentYangName() string { return "soft-preemption" }

// Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State
// State parameters relating to RSVP
// soft preemption support
type Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enables soft preemption on a node. The type is bool. The default value is
    // false.
    Enable interface{}

    // Timeout value for soft preemption to revert to hard preemption. The type is
    // interface{} with range: 0..65535. The default value is 0.
    SoftPreemptionTimeout interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "soft-preemption-timeout" { return "SoftPreemptionTimeout" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = state.Enable
    leafs["soft-preemption-timeout"] = state.SoftPreemptionTimeout
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_SoftPreemption_State) GetParentYangName() string { return "soft-preemption" }

// Mpls_SignalingProtocols_RsvpTe_Global_Hellos
// Top level container for RSVP hello parameters
type Mpls_SignalingProtocols_RsvpTe_Global_Hellos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to RSVP hellos.
    Config Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config

    // State information associated with RSVP hellos.
    State Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetFilter() yfilter.YFilter { return hellos.YFilter }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) SetFilter(yf yfilter.YFilter) { hellos.YFilter = yf }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetSegmentPath() string {
    return "hellos"
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &hellos.Config
    }
    if childYangName == "state" {
        return &hellos.State
    }
    return nil
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &hellos.Config
    children["state"] = &hellos.State
    return children
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetBundleName() string { return "openconfig" }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetYangName() string { return "hellos" }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) SetParent(parent types.Entity) { hellos.parent = parent }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetParent() types.Entity { return hellos.parent }

func (hellos *Mpls_SignalingProtocols_RsvpTe_Global_Hellos) GetParentYangName() string { return "global" }

// Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config
// Configuration parameters relating to RSVP
// hellos
type Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // set the interval in ms between RSVP hello messages. The type is interface{}
    // with range: 1000..60000. Units are milliseconds. The default value is 9000.
    HelloInterval interface{}

    // enables all RSVP refresh reduction message bundling, RSVP message ID,
    // reliable message delivery and summary refresh. The type is bool. The
    // default value is true.
    RefreshReduction interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetGoName(yname string) string {
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "refresh-reduction" { return "RefreshReduction" }
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-interval"] = config.HelloInterval
    leafs["refresh-reduction"] = config.RefreshReduction
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_Config) GetParentYangName() string { return "hellos" }

// Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State
// State information associated with RSVP hellos
type Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // set the interval in ms between RSVP hello messages. The type is interface{}
    // with range: 1000..60000. Units are milliseconds. The default value is 9000.
    HelloInterval interface{}

    // enables all RSVP refresh reduction message bundling, RSVP message ID,
    // reliable message delivery and summary refresh. The type is bool. The
    // default value is true.
    RefreshReduction interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetGoName(yname string) string {
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "refresh-reduction" { return "RefreshReduction" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-interval"] = state.HelloInterval
    leafs["refresh-reduction"] = state.RefreshReduction
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_Hellos_State) GetParentYangName() string { return "hellos" }

// Mpls_SignalingProtocols_RsvpTe_Global_State
// Platform wide RSVP state, including counters
type Mpls_SignalingProtocols_RsvpTe_Global_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Platform wide RSVP statistics and counters.
    Counters Mpls_SignalingProtocols_RsvpTe_Global_State_Counters
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetGoName(yname string) string {
    if yname == "counters" { return "Counters" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &state.Counters
    }
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &state.Counters
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_Global_State) GetParentYangName() string { return "global" }

// Mpls_SignalingProtocols_RsvpTe_Global_State_Counters
// Platform wide RSVP statistics and counters
type Mpls_SignalingProtocols_RsvpTe_Global_State_Counters struct {
    parent types.Entity
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

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetGoName(yname string) string {
    if yname == "path-timeouts" { return "PathTimeouts" }
    if yname == "reservation-timeouts" { return "ReservationTimeouts" }
    if yname == "rate-limited-messages" { return "RateLimitedMessages" }
    if yname == "in-path-messages" { return "InPathMessages" }
    if yname == "in-path-error-messages" { return "InPathErrorMessages" }
    if yname == "in-path-tear-messages" { return "InPathTearMessages" }
    if yname == "in-reservation-messages" { return "InReservationMessages" }
    if yname == "in-reservation-error-messages" { return "InReservationErrorMessages" }
    if yname == "in-reservation-tear-messages" { return "InReservationTearMessages" }
    if yname == "in-hello-messages" { return "InHelloMessages" }
    if yname == "in-srefresh-messages" { return "InSrefreshMessages" }
    if yname == "in-ack-messages" { return "InAckMessages" }
    if yname == "out-path-messages" { return "OutPathMessages" }
    if yname == "out-path-error-messages" { return "OutPathErrorMessages" }
    if yname == "out-path-tear-messages" { return "OutPathTearMessages" }
    if yname == "out-reservation-messages" { return "OutReservationMessages" }
    if yname == "out-reservation-error-messages" { return "OutReservationErrorMessages" }
    if yname == "out-reservation-tear-messages" { return "OutReservationTearMessages" }
    if yname == "out-hello-messages" { return "OutHelloMessages" }
    if yname == "out-srefresh-messages" { return "OutSrefreshMessages" }
    if yname == "out-ack-messages" { return "OutAckMessages" }
    return ""
}

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-timeouts"] = counters.PathTimeouts
    leafs["reservation-timeouts"] = counters.ReservationTimeouts
    leafs["rate-limited-messages"] = counters.RateLimitedMessages
    leafs["in-path-messages"] = counters.InPathMessages
    leafs["in-path-error-messages"] = counters.InPathErrorMessages
    leafs["in-path-tear-messages"] = counters.InPathTearMessages
    leafs["in-reservation-messages"] = counters.InReservationMessages
    leafs["in-reservation-error-messages"] = counters.InReservationErrorMessages
    leafs["in-reservation-tear-messages"] = counters.InReservationTearMessages
    leafs["in-hello-messages"] = counters.InHelloMessages
    leafs["in-srefresh-messages"] = counters.InSrefreshMessages
    leafs["in-ack-messages"] = counters.InAckMessages
    leafs["out-path-messages"] = counters.OutPathMessages
    leafs["out-path-error-messages"] = counters.OutPathErrorMessages
    leafs["out-path-tear-messages"] = counters.OutPathTearMessages
    leafs["out-reservation-messages"] = counters.OutReservationMessages
    leafs["out-reservation-error-messages"] = counters.OutReservationErrorMessages
    leafs["out-reservation-tear-messages"] = counters.OutReservationTearMessages
    leafs["out-hello-messages"] = counters.OutHelloMessages
    leafs["out-srefresh-messages"] = counters.OutSrefreshMessages
    leafs["out-ack-messages"] = counters.OutAckMessages
    return leafs
}

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetBundleName() string { return "openconfig" }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetYangName() string { return "counters" }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Mpls_SignalingProtocols_RsvpTe_Global_State_Counters) GetParentYangName() string { return "state" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes
// Attributes relating to RSVP-TE enabled interfaces
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // list of per-interface RSVP configurations. The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface.
    Interface []Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface
}

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetFilter() yfilter.YFilter { return interfaceAttributes.YFilter }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) SetFilter(yf yfilter.YFilter) { interfaceAttributes.YFilter = yf }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetSegmentPath() string {
    return "interface-attributes"
}

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaceAttributes.Interface {
            if interfaceAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface{}
        interfaceAttributes.Interface = append(interfaceAttributes.Interface, child)
        return &interfaceAttributes.Interface[len(interfaceAttributes.Interface)-1]
    }
    return nil
}

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceAttributes.Interface {
        children[interfaceAttributes.Interface[i].GetSegmentPath()] = &interfaceAttributes.Interface[i]
    }
    return children
}

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetBundleName() string { return "openconfig" }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetYangName() string { return "interface-attributes" }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) SetParent(parent types.Entity) { interfaceAttributes.parent = parent }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetParent() types.Entity { return interfaceAttributes.parent }

func (interfaceAttributes *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes) GetParentYangName() string { return "rsvp-te" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface
// list of per-interface RSVP configurations
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface struct {
    parent types.Entity
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

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "hellos" { return "Hellos" }
    if yname == "authentication" { return "Authentication" }
    if yname == "subscription" { return "Subscription" }
    if yname == "protection" { return "Protection" }
    return ""
}

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &self.Config
    }
    if childYangName == "state" {
        return &self.State
    }
    if childYangName == "hellos" {
        return &self.Hellos
    }
    if childYangName == "authentication" {
        return &self.Authentication
    }
    if childYangName == "subscription" {
        return &self.Subscription
    }
    if childYangName == "protection" {
        return &self.Protection
    }
    return nil
}

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &self.Config
    children["state"] = &self.State
    children["hellos"] = &self.Hellos
    children["authentication"] = &self.Authentication
    children["subscription"] = &self.Subscription
    children["protection"] = &self.Protection
    return children
}

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetBundleName() string { return "openconfig" }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetYangName() string { return "interface" }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetParent() types.Entity { return self.parent }

func (self *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface) GetParentYangName() string { return "interface-attributes" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config
// Configuration of per-interface RSVP parameters
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of configured IP interface. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    InterfaceName interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = config.InterfaceName
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config) GetParentYangName() string { return "interface" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State
// Per-interface RSVP protocol and state information
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State struct {
    parent types.Entity
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

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetGoName(yname string) string {
    if yname == "highwater-mark" { return "HighwaterMark" }
    if yname == "active-reservation-count" { return "ActiveReservationCount" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        for _, c := range state.Bandwidth {
            if state.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth{}
        state.Bandwidth = append(state.Bandwidth, child)
        return &state.Bandwidth[len(state.Bandwidth)-1]
    }
    if childYangName == "counters" {
        return &state.Counters
    }
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range state.Bandwidth {
        children[state.Bandwidth[i].GetSegmentPath()] = &state.Bandwidth[i]
    }
    children["counters"] = &state.Counters
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["highwater-mark"] = state.HighwaterMark
    leafs["active-reservation-count"] = state.ActiveReservationCount
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State) GetParentYangName() string { return "interface" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth
// Available and reserved bandwidth by priority on
// the interface.
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth struct {
    parent types.Entity
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

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetGoName(yname string) string {
    if yname == "priority" { return "Priority" }
    if yname == "available-bandwidth" { return "AvailableBandwidth" }
    if yname == "reserved-bandwidth" { return "ReservedBandwidth" }
    return ""
}

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetSegmentPath() string {
    return "bandwidth" + "[priority='" + fmt.Sprintf("%v", bandwidth.Priority) + "']"
}

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority"] = bandwidth.Priority
    leafs["available-bandwidth"] = bandwidth.AvailableBandwidth
    leafs["reserved-bandwidth"] = bandwidth.ReservedBandwidth
    return leafs
}

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetBundleName() string { return "openconfig" }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Bandwidth) GetParentYangName() string { return "state" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters
// Interface specific RSVP statistics and counters
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters struct {
    parent types.Entity
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

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetGoName(yname string) string {
    if yname == "in-path-messages" { return "InPathMessages" }
    if yname == "in-path-error-messages" { return "InPathErrorMessages" }
    if yname == "in-path-tear-messages" { return "InPathTearMessages" }
    if yname == "in-reservation-messages" { return "InReservationMessages" }
    if yname == "in-reservation-error-messages" { return "InReservationErrorMessages" }
    if yname == "in-reservation-tear-messages" { return "InReservationTearMessages" }
    if yname == "in-hello-messages" { return "InHelloMessages" }
    if yname == "in-srefresh-messages" { return "InSrefreshMessages" }
    if yname == "in-ack-messages" { return "InAckMessages" }
    if yname == "out-path-messages" { return "OutPathMessages" }
    if yname == "out-path-error-messages" { return "OutPathErrorMessages" }
    if yname == "out-path-tear-messages" { return "OutPathTearMessages" }
    if yname == "out-reservation-messages" { return "OutReservationMessages" }
    if yname == "out-reservation-error-messages" { return "OutReservationErrorMessages" }
    if yname == "out-reservation-tear-messages" { return "OutReservationTearMessages" }
    if yname == "out-hello-messages" { return "OutHelloMessages" }
    if yname == "out-srefresh-messages" { return "OutSrefreshMessages" }
    if yname == "out-ack-messages" { return "OutAckMessages" }
    return ""
}

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-path-messages"] = counters.InPathMessages
    leafs["in-path-error-messages"] = counters.InPathErrorMessages
    leafs["in-path-tear-messages"] = counters.InPathTearMessages
    leafs["in-reservation-messages"] = counters.InReservationMessages
    leafs["in-reservation-error-messages"] = counters.InReservationErrorMessages
    leafs["in-reservation-tear-messages"] = counters.InReservationTearMessages
    leafs["in-hello-messages"] = counters.InHelloMessages
    leafs["in-srefresh-messages"] = counters.InSrefreshMessages
    leafs["in-ack-messages"] = counters.InAckMessages
    leafs["out-path-messages"] = counters.OutPathMessages
    leafs["out-path-error-messages"] = counters.OutPathErrorMessages
    leafs["out-path-tear-messages"] = counters.OutPathTearMessages
    leafs["out-reservation-messages"] = counters.OutReservationMessages
    leafs["out-reservation-error-messages"] = counters.OutReservationErrorMessages
    leafs["out-reservation-tear-messages"] = counters.OutReservationTearMessages
    leafs["out-hello-messages"] = counters.OutHelloMessages
    leafs["out-srefresh-messages"] = counters.OutSrefreshMessages
    leafs["out-ack-messages"] = counters.OutAckMessages
    return leafs
}

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetBundleName() string { return "openconfig" }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetYangName() string { return "counters" }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State_Counters) GetParentYangName() string { return "state" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos
// Top level container for RSVP hello parameters
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to RSVP hellos.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config

    // State information associated with RSVP hellos.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetFilter() yfilter.YFilter { return hellos.YFilter }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) SetFilter(yf yfilter.YFilter) { hellos.YFilter = yf }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetSegmentPath() string {
    return "hellos"
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &hellos.Config
    }
    if childYangName == "state" {
        return &hellos.State
    }
    return nil
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &hellos.Config
    children["state"] = &hellos.State
    return children
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetBundleName() string { return "openconfig" }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetYangName() string { return "hellos" }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) SetParent(parent types.Entity) { hellos.parent = parent }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetParent() types.Entity { return hellos.parent }

func (hellos *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos) GetParentYangName() string { return "interface" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config
// Configuration parameters relating to RSVP
// hellos
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // set the interval in ms between RSVP hello messages. The type is interface{}
    // with range: 1000..60000. Units are milliseconds. The default value is 9000.
    HelloInterval interface{}

    // enables all RSVP refresh reduction message bundling, RSVP message ID,
    // reliable message delivery and summary refresh. The type is bool. The
    // default value is true.
    RefreshReduction interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetGoName(yname string) string {
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "refresh-reduction" { return "RefreshReduction" }
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-interval"] = config.HelloInterval
    leafs["refresh-reduction"] = config.RefreshReduction
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_Config) GetParentYangName() string { return "hellos" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State
// State information associated with RSVP hellos
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // set the interval in ms between RSVP hello messages. The type is interface{}
    // with range: 1000..60000. Units are milliseconds. The default value is 9000.
    HelloInterval interface{}

    // enables all RSVP refresh reduction message bundling, RSVP message ID,
    // reliable message delivery and summary refresh. The type is bool. The
    // default value is true.
    RefreshReduction interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetGoName(yname string) string {
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "refresh-reduction" { return "RefreshReduction" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-interval"] = state.HelloInterval
    leafs["refresh-reduction"] = state.RefreshReduction
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Hellos_State) GetParentYangName() string { return "hellos" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication
// Configuration and state parameters relating to RSVP
// authentication as per RFC2747
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to authentication.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config

    // State information associated with authentication.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State
}

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &authentication.Config
    }
    if childYangName == "state" {
        return &authentication.State
    }
    return nil
}

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &authentication.Config
    children["state"] = &authentication.State
    return children
}

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetBundleName() string { return "openconfig" }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetYangName() string { return "authentication" }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication) GetParentYangName() string { return "interface" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config
// Configuration parameters relating
// to authentication
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enables RSVP authentication on the node. The type is bool. The default
    // value is false.
    Enable interface{}

    // authenticate RSVP signaling messages. The type is string with length:
    // 1..32.
    AuthenticationKey interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "authentication-key" { return "AuthenticationKey" }
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = config.Enable
    leafs["authentication-key"] = config.AuthenticationKey
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_Config) GetParentYangName() string { return "authentication" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State
// State information associated
// with authentication
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enables RSVP authentication on the node. The type is bool. The default
    // value is false.
    Enable interface{}

    // authenticate RSVP signaling messages. The type is string with length:
    // 1..32.
    AuthenticationKey interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "authentication-key" { return "AuthenticationKey" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = state.Enable
    leafs["authentication-key"] = state.AuthenticationKey
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Authentication_State) GetParentYangName() string { return "authentication" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription
// Bandwidth percentage reservable by RSVP
// on an interface
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to RSVP subscription options.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config

    // State parameters relating to RSVP subscription options.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State
}

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetSegmentPath() string {
    return "subscription"
}

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &subscription.Config
    }
    if childYangName == "state" {
        return &subscription.State
    }
    return nil
}

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &subscription.Config
    children["state"] = &subscription.State
    return children
}

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetBundleName() string { return "openconfig" }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetYangName() string { return "subscription" }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) SetParent(parent types.Entity) { subscription.parent = parent }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetParent() types.Entity { return subscription.parent }

func (subscription *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription) GetParentYangName() string { return "interface" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config
// Configuration parameters relating to RSVP
// subscription options
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // percentage of the interface bandwidth that RSVP can reserve. The type is
    // interface{} with range: 0..100.
    Subscription interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetGoName(yname string) string {
    if yname == "subscription" { return "Subscription" }
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription"] = config.Subscription
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_Config) GetParentYangName() string { return "subscription" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State
// State parameters relating to RSVP
// subscription options
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // percentage of the interface bandwidth that RSVP can reserve. The type is
    // interface{} with range: 0..100.
    Subscription interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetGoName(yname string) string {
    if yname == "subscription" { return "Subscription" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscription"] = state.Subscription
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Subscription_State) GetParentYangName() string { return "subscription" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection
// link-protection (NHOP) related configuration
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for link-protection.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config

    // State for link-protection.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State
}

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetFilter() yfilter.YFilter { return protection.YFilter }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) SetFilter(yf yfilter.YFilter) { protection.YFilter = yf }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetSegmentPath() string {
    return "protection"
}

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &protection.Config
    }
    if childYangName == "state" {
        return &protection.State
    }
    return nil
}

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &protection.Config
    children["state"] = &protection.State
    return children
}

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetBundleName() string { return "openconfig" }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetYangName() string { return "protection" }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) SetParent(parent types.Entity) { protection.parent = parent }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetParent() types.Entity { return protection.parent }

func (protection *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection) GetParentYangName() string { return "interface" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config
// Configuration for link-protection
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config struct {
    parent types.Entity
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

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetGoName(yname string) string {
    if yname == "link-protection-style-requested" { return "LinkProtectionStyleRequested" }
    if yname == "bypass-optimize-interval" { return "BypassOptimizeInterval" }
    return ""
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-protection-style-requested"] = config.LinkProtectionStyleRequested
    leafs["bypass-optimize-interval"] = config.BypassOptimizeInterval
    return leafs
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config) GetParentYangName() string { return "protection" }

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State
// State for link-protection
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State struct {
    parent types.Entity
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

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetGoName(yname string) string {
    if yname == "link-protection-style-requested" { return "LinkProtectionStyleRequested" }
    if yname == "bypass-optimize-interval" { return "BypassOptimizeInterval" }
    return ""
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-protection-style-requested"] = state.LinkProtectionStyleRequested
    leafs["bypass-optimize-interval"] = state.BypassOptimizeInterval
    return leafs
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State) GetParentYangName() string { return "protection" }

// Mpls_SignalingProtocols_SegmentRouting
// SR global signaling config
type Mpls_SignalingProtocols_SegmentRouting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of Segment Routing Global Block (SRGB) entries. These label blocks are
    // reserved to be allocated as domain-wide entries. The type is slice of
    // Mpls_SignalingProtocols_SegmentRouting_Srgb.
    Srgb []Mpls_SignalingProtocols_SegmentRouting_Srgb

    // List of interfaces with associated segment routing configuration. The type
    // is slice of Mpls_SignalingProtocols_SegmentRouting_Interfaces.
    Interfaces []Mpls_SignalingProtocols_SegmentRouting_Interfaces
}

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetFilter() yfilter.YFilter { return segmentRouting.YFilter }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) SetFilter(yf yfilter.YFilter) { segmentRouting.YFilter = yf }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetGoName(yname string) string {
    if yname == "srgb" { return "Srgb" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetSegmentPath() string {
    return "segment-routing"
}

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srgb" {
        for _, c := range segmentRouting.Srgb {
            if segmentRouting.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_SignalingProtocols_SegmentRouting_Srgb{}
        segmentRouting.Srgb = append(segmentRouting.Srgb, child)
        return &segmentRouting.Srgb[len(segmentRouting.Srgb)-1]
    }
    if childYangName == "interfaces" {
        for _, c := range segmentRouting.Interfaces {
            if segmentRouting.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_SignalingProtocols_SegmentRouting_Interfaces{}
        segmentRouting.Interfaces = append(segmentRouting.Interfaces, child)
        return &segmentRouting.Interfaces[len(segmentRouting.Interfaces)-1]
    }
    return nil
}

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range segmentRouting.Srgb {
        children[segmentRouting.Srgb[i].GetSegmentPath()] = &segmentRouting.Srgb[i]
    }
    for i := range segmentRouting.Interfaces {
        children[segmentRouting.Interfaces[i].GetSegmentPath()] = &segmentRouting.Interfaces[i]
    }
    return children
}

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetBundleName() string { return "openconfig" }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetYangName() string { return "segment-routing" }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) SetParent(parent types.Entity) { segmentRouting.parent = parent }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetParent() types.Entity { return segmentRouting.parent }

func (segmentRouting *Mpls_SignalingProtocols_SegmentRouting) GetParentYangName() string { return "signaling-protocols" }

// Mpls_SignalingProtocols_SegmentRouting_Srgb
// List of Segment Routing Global Block (SRGB) entries. These
// label blocks are reserved to be allocated as domain-wide
// entries.
type Mpls_SignalingProtocols_SegmentRouting_Srgb struct {
    parent types.Entity
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

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetFilter() yfilter.YFilter { return srgb.YFilter }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) SetFilter(yf yfilter.YFilter) { srgb.YFilter = yf }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetSegmentPath() string {
    return "srgb" + "[lower-bound='" + fmt.Sprintf("%v", srgb.LowerBound) + "']" + "[upper-bound='" + fmt.Sprintf("%v", srgb.UpperBound) + "']"
}

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &srgb.Config
    }
    if childYangName == "state" {
        return &srgb.State
    }
    return nil
}

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &srgb.Config
    children["state"] = &srgb.State
    return children
}

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = srgb.LowerBound
    leafs["upper-bound"] = srgb.UpperBound
    return leafs
}

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetBundleName() string { return "openconfig" }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetYangName() string { return "srgb" }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) SetParent(parent types.Entity) { srgb.parent = parent }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetParent() types.Entity { return srgb.parent }

func (srgb *Mpls_SignalingProtocols_SegmentRouting_Srgb) GetParentYangName() string { return "segment-routing" }

// Mpls_SignalingProtocols_SegmentRouting_Srgb_Config
// Configuration parameters relating to the Segment Routing
// Global Block (SRGB)
type Mpls_SignalingProtocols_SegmentRouting_Srgb_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Lower value in the block. The type is interface{} with range:
    // 0..4294967295.
    LowerBound interface{}

    // Upper value in the block. The type is interface{} with range:
    // 0..4294967295.
    UpperBound interface{}
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    return ""
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = config.LowerBound
    leafs["upper-bound"] = config.UpperBound
    return leafs
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_SegmentRouting_Srgb_Config) GetParentYangName() string { return "srgb" }

// Mpls_SignalingProtocols_SegmentRouting_Srgb_State
// State parameters relating to the Segment Routing Global
// Block (SRGB)
type Mpls_SignalingProtocols_SegmentRouting_Srgb_State struct {
    parent types.Entity
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

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    if yname == "size" { return "Size" }
    if yname == "free" { return "Free" }
    if yname == "used" { return "Used" }
    return ""
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = state.LowerBound
    leafs["upper-bound"] = state.UpperBound
    leafs["size"] = state.Size
    leafs["free"] = state.Free
    leafs["used"] = state.Used
    return leafs
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_SegmentRouting_Srgb_State) GetParentYangName() string { return "srgb" }

// Mpls_SignalingProtocols_SegmentRouting_Interfaces
// List of interfaces with associated segment routing
// configuration
type Mpls_SignalingProtocols_SegmentRouting_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the interface for which segment
    // routing configuration is to be applied. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Interface interface{}

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

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "adjacency-sid" { return "AdjacencySid" }
    return ""
}

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetSegmentPath() string {
    return "interfaces" + "[interface='" + fmt.Sprintf("%v", interfaces.Interface) + "']"
}

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaces.Config
    }
    if childYangName == "state" {
        return &interfaces.State
    }
    if childYangName == "adjacency-sid" {
        return &interfaces.AdjacencySid
    }
    return nil
}

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaces.Config
    children["state"] = &interfaces.State
    children["adjacency-sid"] = &interfaces.AdjacencySid
    return children
}

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = interfaces.Interface
    return leafs
}

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetBundleName() string { return "openconfig" }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetParentYangName() string { return "segment-routing" }

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config
// Interface configuration parameters for Segment Routing
// relating to the specified interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to the interface for which segment routing configuration is to be
    // applied. The type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = config.Interface
    return leafs
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Config) GetParentYangName() string { return "interfaces" }

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_State
// State parameters for Segment Routing features relating
// to the specified interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to the interface for which segment routing configuration is to be
    // applied. The type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    return leafs
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_State) GetParentYangName() string { return "interfaces" }

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid
// Configuration for Adjacency SIDs that are related to
// the specified interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters for the Adjacency-SIDs that are related to this
    // interface.
    Config Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config

    // State parameters for the Adjacency-SIDs that are related to this interface.
    State Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State
}

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetFilter() yfilter.YFilter { return adjacencySid.YFilter }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) SetFilter(yf yfilter.YFilter) { adjacencySid.YFilter = yf }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetSegmentPath() string {
    return "adjacency-sid"
}

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &adjacencySid.Config
    }
    if childYangName == "state" {
        return &adjacencySid.State
    }
    return nil
}

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &adjacencySid.Config
    children["state"] = &adjacencySid.State
    return children
}

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetBundleName() string { return "openconfig" }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetYangName() string { return "adjacency-sid" }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) SetParent(parent types.Entity) { adjacencySid.parent = parent }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetParent() types.Entity { return adjacencySid.parent }

func (adjacencySid *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid) GetParentYangName() string { return "interfaces" }

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config
// Configuration parameters for the Adjacency-SIDs
// that are related to this interface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config struct {
    parent types.Entity
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

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetGoName(yname string) string {
    if yname == "advertise" { return "Advertise" }
    if yname == "groups" { return "Groups" }
    return ""
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["advertise"] = config.Advertise
    leafs["groups"] = config.Groups
    return leafs
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetYangName() string { return "config" }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_Config) GetParentYangName() string { return "adjacency-sid" }

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
    parent types.Entity
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

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetGoName(yname string) string {
    if yname == "advertise" { return "Advertise" }
    if yname == "groups" { return "Groups" }
    return ""
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["advertise"] = state.Advertise
    leafs["groups"] = state.Groups
    return leafs
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetYangName() string { return "state" }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_AdjacencySid_State) GetParentYangName() string { return "adjacency-sid" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP timers.
    Timers Mpls_SignalingProtocols_Ldp_Timers
}

func (ldp *Mpls_SignalingProtocols_Ldp) GetFilter() yfilter.YFilter { return ldp.YFilter }

func (ldp *Mpls_SignalingProtocols_Ldp) SetFilter(yf yfilter.YFilter) { ldp.YFilter = yf }

func (ldp *Mpls_SignalingProtocols_Ldp) GetGoName(yname string) string {
    if yname == "timers" { return "Timers" }
    return ""
}

func (ldp *Mpls_SignalingProtocols_Ldp) GetSegmentPath() string {
    return "ldp"
}

func (ldp *Mpls_SignalingProtocols_Ldp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "timers" {
        return &ldp.Timers
    }
    return nil
}

func (ldp *Mpls_SignalingProtocols_Ldp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["timers"] = &ldp.Timers
    return children
}

func (ldp *Mpls_SignalingProtocols_Ldp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldp *Mpls_SignalingProtocols_Ldp) GetBundleName() string { return "openconfig" }

func (ldp *Mpls_SignalingProtocols_Ldp) GetYangName() string { return "ldp" }

func (ldp *Mpls_SignalingProtocols_Ldp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ldp *Mpls_SignalingProtocols_Ldp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ldp *Mpls_SignalingProtocols_Ldp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ldp *Mpls_SignalingProtocols_Ldp) SetParent(parent types.Entity) { ldp.parent = parent }

func (ldp *Mpls_SignalingProtocols_Ldp) GetParent() types.Entity { return ldp.parent }

func (ldp *Mpls_SignalingProtocols_Ldp) GetParentYangName() string { return "signaling-protocols" }

// Mpls_SignalingProtocols_Ldp_Timers
// LDP timers
type Mpls_SignalingProtocols_Ldp_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetGoName(yname string) string {
    return ""
}

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetBundleName() string { return "openconfig" }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetYangName() string { return "timers" }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Mpls_SignalingProtocols_Ldp_Timers) GetParentYangName() string { return "ldp" }

// Mpls_Lsps
// LSP definitions and configuration
type Mpls_Lsps struct {
    parent types.Entity
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

func (lsps *Mpls_Lsps) GetFilter() yfilter.YFilter { return lsps.YFilter }

func (lsps *Mpls_Lsps) SetFilter(yf yfilter.YFilter) { lsps.YFilter = yf }

func (lsps *Mpls_Lsps) GetGoName(yname string) string {
    if yname == "constrained-path" { return "ConstrainedPath" }
    if yname == "unconstrained-path" { return "UnconstrainedPath" }
    if yname == "static-lsps" { return "StaticLsps" }
    return ""
}

func (lsps *Mpls_Lsps) GetSegmentPath() string {
    return "lsps"
}

func (lsps *Mpls_Lsps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "constrained-path" {
        return &lsps.ConstrainedPath
    }
    if childYangName == "unconstrained-path" {
        return &lsps.UnconstrainedPath
    }
    if childYangName == "static-lsps" {
        return &lsps.StaticLsps
    }
    return nil
}

func (lsps *Mpls_Lsps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["constrained-path"] = &lsps.ConstrainedPath
    children["unconstrained-path"] = &lsps.UnconstrainedPath
    children["static-lsps"] = &lsps.StaticLsps
    return children
}

func (lsps *Mpls_Lsps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lsps *Mpls_Lsps) GetBundleName() string { return "openconfig" }

func (lsps *Mpls_Lsps) GetYangName() string { return "lsps" }

func (lsps *Mpls_Lsps) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (lsps *Mpls_Lsps) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (lsps *Mpls_Lsps) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (lsps *Mpls_Lsps) SetParent(parent types.Entity) { lsps.parent = parent }

func (lsps *Mpls_Lsps) GetParent() types.Entity { return lsps.parent }

func (lsps *Mpls_Lsps) GetParentYangName() string { return "mpls" }

// Mpls_Lsps_ConstrainedPath
// traffic-engineered LSPs supporting different
// path computation and signaling methods
type Mpls_Lsps_ConstrainedPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of explicit paths. The type is slice of
    // Mpls_Lsps_ConstrainedPath_NamedExplicitPaths.
    NamedExplicitPaths []Mpls_Lsps_ConstrainedPath_NamedExplicitPaths

    // List of TE tunnels. The type is slice of Mpls_Lsps_ConstrainedPath_Tunnel.
    Tunnel []Mpls_Lsps_ConstrainedPath_Tunnel
}

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetFilter() yfilter.YFilter { return constrainedPath.YFilter }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) SetFilter(yf yfilter.YFilter) { constrainedPath.YFilter = yf }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetGoName(yname string) string {
    if yname == "named-explicit-paths" { return "NamedExplicitPaths" }
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetSegmentPath() string {
    return "constrained-path"
}

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "named-explicit-paths" {
        for _, c := range constrainedPath.NamedExplicitPaths {
            if constrainedPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Lsps_ConstrainedPath_NamedExplicitPaths{}
        constrainedPath.NamedExplicitPaths = append(constrainedPath.NamedExplicitPaths, child)
        return &constrainedPath.NamedExplicitPaths[len(constrainedPath.NamedExplicitPaths)-1]
    }
    if childYangName == "tunnel" {
        for _, c := range constrainedPath.Tunnel {
            if constrainedPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Lsps_ConstrainedPath_Tunnel{}
        constrainedPath.Tunnel = append(constrainedPath.Tunnel, child)
        return &constrainedPath.Tunnel[len(constrainedPath.Tunnel)-1]
    }
    return nil
}

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range constrainedPath.NamedExplicitPaths {
        children[constrainedPath.NamedExplicitPaths[i].GetSegmentPath()] = &constrainedPath.NamedExplicitPaths[i]
    }
    for i := range constrainedPath.Tunnel {
        children[constrainedPath.Tunnel[i].GetSegmentPath()] = &constrainedPath.Tunnel[i]
    }
    return children
}

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetBundleName() string { return "openconfig" }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetYangName() string { return "constrained-path" }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) SetParent(parent types.Entity) { constrainedPath.parent = parent }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetParent() types.Entity { return constrainedPath.parent }

func (constrainedPath *Mpls_Lsps_ConstrainedPath) GetParentYangName() string { return "lsps" }

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths
// A list of explicit paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths struct {
    parent types.Entity
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

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetFilter() yfilter.YFilter { return namedExplicitPaths.YFilter }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) SetFilter(yf yfilter.YFilter) { namedExplicitPaths.YFilter = yf }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "explicit-route-objects" { return "ExplicitRouteObjects" }
    return ""
}

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetSegmentPath() string {
    return "named-explicit-paths" + "[name='" + fmt.Sprintf("%v", namedExplicitPaths.Name) + "']"
}

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &namedExplicitPaths.Config
    }
    if childYangName == "state" {
        return &namedExplicitPaths.State
    }
    if childYangName == "explicit-route-objects" {
        for _, c := range namedExplicitPaths.ExplicitRouteObjects {
            if namedExplicitPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects{}
        namedExplicitPaths.ExplicitRouteObjects = append(namedExplicitPaths.ExplicitRouteObjects, child)
        return &namedExplicitPaths.ExplicitRouteObjects[len(namedExplicitPaths.ExplicitRouteObjects)-1]
    }
    return nil
}

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &namedExplicitPaths.Config
    children["state"] = &namedExplicitPaths.State
    for i := range namedExplicitPaths.ExplicitRouteObjects {
        children[namedExplicitPaths.ExplicitRouteObjects[i].GetSegmentPath()] = &namedExplicitPaths.ExplicitRouteObjects[i]
    }
    return children
}

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = namedExplicitPaths.Name
    return leafs
}

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetBundleName() string { return "openconfig" }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetYangName() string { return "named-explicit-paths" }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) SetParent(parent types.Entity) { namedExplicitPaths.parent = parent }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetParent() types.Entity { return namedExplicitPaths.parent }

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetParentYangName() string { return "constrained-path" }

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config
// Configuration parameters relating to named explicit
// paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A string name that uniquely identifies an explicit path. The type is
    // string.
    Name interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_Config) GetParentYangName() string { return "named-explicit-paths" }

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State
// Operational state parameters relating to the named
// explicit paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A string name that uniquely identifies an explicit path. The type is
    // string.
    Name interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_State) GetParentYangName() string { return "named-explicit-paths" }

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects
// List of explicit route objects
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects struct {
    parent types.Entity
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

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetFilter() yfilter.YFilter { return explicitRouteObjects.YFilter }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) SetFilter(yf yfilter.YFilter) { explicitRouteObjects.YFilter = yf }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetSegmentPath() string {
    return "explicit-route-objects" + "[index='" + fmt.Sprintf("%v", explicitRouteObjects.Index) + "']"
}

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &explicitRouteObjects.Config
    }
    if childYangName == "state" {
        return &explicitRouteObjects.State
    }
    return nil
}

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &explicitRouteObjects.Config
    children["state"] = &explicitRouteObjects.State
    return children
}

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = explicitRouteObjects.Index
    return leafs
}

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetBundleName() string { return "openconfig" }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetYangName() string { return "explicit-route-objects" }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) SetParent(parent types.Entity) { explicitRouteObjects.parent = parent }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetParent() types.Entity { return explicitRouteObjects.parent }

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects) GetParentYangName() string { return "named-explicit-paths" }

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config
// Configuration parameters relating to an explicit
// route
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // router hop for the LSP path. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // strict or loose hop. The type is MplsHopType.
    HopType interface{}

    // Index of this explicit route object to express the order of hops in the
    // path. The type is interface{} with range: 0..255.
    Index interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "hop-type" { return "HopType" }
    if yname == "index" { return "Index" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = config.Address
    leafs["hop-type"] = config.HopType
    leafs["index"] = config.Index
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_Config) GetParentYangName() string { return "explicit-route-objects" }

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State
// State parameters relating to an explicit route
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // router hop for the LSP path. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // strict or loose hop. The type is MplsHopType.
    HopType interface{}

    // Index of this explicit route object to express the order of hops in the
    // path. The type is interface{} with range: 0..255.
    Index interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "hop-type" { return "HopType" }
    if yname == "index" { return "Index" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = state.Address
    leafs["hop-type"] = state.HopType
    leafs["index"] = state.Index
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_ExplicitRouteObjects_State) GetParentYangName() string { return "explicit-route-objects" }

// Mpls_Lsps_ConstrainedPath_Tunnel
// List of TE tunnels
type Mpls_Lsps_ConstrainedPath_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The tunnel name. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnel_Config_Name
    Name interface{}

    // This attribute is a key. The tunnel type, p2p or p2mp. The type is one of
    // the following: P2PP2MP.
    Type interface{}

    // Configuration parameters related to TE tunnels:.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Config

    // State parameters related to TE tunnels.
    State Mpls_Lsps_ConstrainedPath_Tunnel_State

    // Bandwidth configuration for TE LSPs.
    Bandwidth Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth

    // Parameters related to LSPs of type P2P.
    P2PTunnelAttributes Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes
}

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "p2p-tunnel-attributes" { return "P2PTunnelAttributes" }
    return ""
}

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetSegmentPath() string {
    return "tunnel" + "[name='" + fmt.Sprintf("%v", tunnel.Name) + "']" + "[type='" + fmt.Sprintf("%v", tunnel.Type) + "']"
}

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &tunnel.Config
    }
    if childYangName == "state" {
        return &tunnel.State
    }
    if childYangName == "bandwidth" {
        return &tunnel.Bandwidth
    }
    if childYangName == "p2p-tunnel-attributes" {
        return &tunnel.P2PTunnelAttributes
    }
    return nil
}

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &tunnel.Config
    children["state"] = &tunnel.State
    children["bandwidth"] = &tunnel.Bandwidth
    children["p2p-tunnel-attributes"] = &tunnel.P2PTunnelAttributes
    return children
}

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = tunnel.Name
    leafs["type"] = tunnel.Type
    return leafs
}

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetBundleName() string { return "openconfig" }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnel) GetParentYangName() string { return "constrained-path" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Config
// Configuration parameters related to TE tunnels:
type Mpls_Lsps_ConstrainedPath_Tunnel_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The tunnel name. The type is string.
    Name interface{}

    // Tunnel type, p2p or p2mp. The type is one of the following: P2PP2MP.
    Type interface{}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "signaling-protocol" { return "SignalingProtocol" }
    if yname == "local-id" { return "LocalId" }
    if yname == "description" { return "Description" }
    if yname == "admin-status" { return "AdminStatus" }
    if yname == "preference" { return "Preference" }
    if yname == "metric" { return "Metric" }
    if yname == "protection-style-requested" { return "ProtectionStyleRequested" }
    if yname == "reoptimize-timer" { return "ReoptimizeTimer" }
    if yname == "source" { return "Source" }
    if yname == "soft-preemption" { return "SoftPreemption" }
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["type"] = config.Type
    leafs["signaling-protocol"] = config.SignalingProtocol
    leafs["local-id"] = config.LocalId
    leafs["description"] = config.Description
    leafs["admin-status"] = config.AdminStatus
    leafs["preference"] = config.Preference
    leafs["metric"] = config.Metric
    leafs["protection-style-requested"] = config.ProtectionStyleRequested
    leafs["reoptimize-timer"] = config.ReoptimizeTimer
    leafs["source"] = config.Source
    leafs["soft-preemption"] = config.SoftPreemption
    leafs["setup-priority"] = config.SetupPriority
    leafs["hold-priority"] = config.HoldPriority
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Config) GetParentYangName() string { return "tunnel" }

// Mpls_Lsps_ConstrainedPath_Tunnel_State
// State parameters related to TE tunnels
type Mpls_Lsps_ConstrainedPath_Tunnel_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The tunnel name. The type is string.
    Name interface{}

    // Tunnel type, p2p or p2mp. The type is one of the following: P2PP2MP.
    Type interface{}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "signaling-protocol" { return "SignalingProtocol" }
    if yname == "local-id" { return "LocalId" }
    if yname == "description" { return "Description" }
    if yname == "admin-status" { return "AdminStatus" }
    if yname == "preference" { return "Preference" }
    if yname == "metric" { return "Metric" }
    if yname == "protection-style-requested" { return "ProtectionStyleRequested" }
    if yname == "reoptimize-timer" { return "ReoptimizeTimer" }
    if yname == "source" { return "Source" }
    if yname == "soft-preemption" { return "SoftPreemption" }
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    if yname == "oper-status" { return "OperStatus" }
    if yname == "role" { return "Role" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &state.Counters
    }
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &state.Counters
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["type"] = state.Type
    leafs["signaling-protocol"] = state.SignalingProtocol
    leafs["local-id"] = state.LocalId
    leafs["description"] = state.Description
    leafs["admin-status"] = state.AdminStatus
    leafs["preference"] = state.Preference
    leafs["metric"] = state.Metric
    leafs["protection-style-requested"] = state.ProtectionStyleRequested
    leafs["reoptimize-timer"] = state.ReoptimizeTimer
    leafs["source"] = state.Source
    leafs["soft-preemption"] = state.SoftPreemption
    leafs["setup-priority"] = state.SetupPriority
    leafs["hold-priority"] = state.HoldPriority
    leafs["oper-status"] = state.OperStatus
    leafs["role"] = state.Role
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_State) GetParentYangName() string { return "tunnel" }

// Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters
// State data for MPLS label switched paths. This state
// data is specific to a single label switched path.
type Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters struct {
    parent types.Entity
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
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    OnlineTime interface{}

    // Indicates the time the LSP switched onto its current path. This is reset
    // upon a LSP path change. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    CurrentPathTime interface{}

    // Indicates the next scheduled time the LSP will be reoptimized. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    NextReoptimizationTime interface{}
}

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetGoName(yname string) string {
    if yname == "bytes" { return "Bytes" }
    if yname == "packets" { return "Packets" }
    if yname == "path-changes" { return "PathChanges" }
    if yname == "state-changes" { return "StateChanges" }
    if yname == "online-time" { return "OnlineTime" }
    if yname == "current-path-time" { return "CurrentPathTime" }
    if yname == "next-reoptimization-time" { return "NextReoptimizationTime" }
    return ""
}

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bytes"] = counters.Bytes
    leafs["packets"] = counters.Packets
    leafs["path-changes"] = counters.PathChanges
    leafs["state-changes"] = counters.StateChanges
    leafs["online-time"] = counters.OnlineTime
    leafs["current-path-time"] = counters.CurrentPathTime
    leafs["next-reoptimization-time"] = counters.NextReoptimizationTime
    return leafs
}

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetBundleName() string { return "openconfig" }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetYangName() string { return "counters" }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Mpls_Lsps_ConstrainedPath_Tunnel_State_Counters) GetParentYangName() string { return "state" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth
// Bandwidth configuration for TE LSPs
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters related to bandwidth on TE tunnels:.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config

    // State parameters related to bandwidth configuration of TE tunnels.
    State Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State

    // Parameters related to auto-bandwidth.
    AutoBandwidth Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth
}

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "auto-bandwidth" { return "AutoBandwidth" }
    return ""
}

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &bandwidth.Config
    }
    if childYangName == "state" {
        return &bandwidth.State
    }
    if childYangName == "auto-bandwidth" {
        return &bandwidth.AutoBandwidth
    }
    return nil
}

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &bandwidth.Config
    children["state"] = &bandwidth.State
    children["auto-bandwidth"] = &bandwidth.AutoBandwidth
    return children
}

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetBundleName() string { return "openconfig" }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth) GetParentYangName() string { return "tunnel" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config
// Configuration parameters related to bandwidth on TE
// tunnels:
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The method used for settign the bandwidth, either explicitly specified or
    // configured. The type is TeBandwidthType. The default value is SPECIFIED.
    SpecificationType interface{}

    // set bandwidth explicitly, e.g., using offline calculation. The type is
    // interface{} with range: 0..4294967295.
    SetBandwidth interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetGoName(yname string) string {
    if yname == "specification-type" { return "SpecificationType" }
    if yname == "set-bandwidth" { return "SetBandwidth" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["specification-type"] = config.SpecificationType
    leafs["set-bandwidth"] = config.SetBandwidth
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_Config) GetParentYangName() string { return "bandwidth" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State
// State parameters related to bandwidth
// configuration of TE tunnels
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The method used for settign the bandwidth, either explicitly specified or
    // configured. The type is TeBandwidthType. The default value is SPECIFIED.
    SpecificationType interface{}

    // set bandwidth explicitly, e.g., using offline calculation. The type is
    // interface{} with range: 0..4294967295.
    SetBandwidth interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetGoName(yname string) string {
    if yname == "specification-type" { return "SpecificationType" }
    if yname == "set-bandwidth" { return "SetBandwidth" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["specification-type"] = state.SpecificationType
    leafs["set-bandwidth"] = state.SetBandwidth
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_State) GetParentYangName() string { return "bandwidth" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth
// Parameters related to auto-bandwidth
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth struct {
    parent types.Entity
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

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetFilter() yfilter.YFilter { return autoBandwidth.YFilter }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) SetFilter(yf yfilter.YFilter) { autoBandwidth.YFilter = yf }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "overflow" { return "Overflow" }
    if yname == "underflow" { return "Underflow" }
    return ""
}

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetSegmentPath() string {
    return "auto-bandwidth"
}

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &autoBandwidth.Config
    }
    if childYangName == "state" {
        return &autoBandwidth.State
    }
    if childYangName == "overflow" {
        return &autoBandwidth.Overflow
    }
    if childYangName == "underflow" {
        return &autoBandwidth.Underflow
    }
    return nil
}

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &autoBandwidth.Config
    children["state"] = &autoBandwidth.State
    children["overflow"] = &autoBandwidth.Overflow
    children["underflow"] = &autoBandwidth.Underflow
    return children
}

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetBundleName() string { return "openconfig" }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetYangName() string { return "auto-bandwidth" }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) SetParent(parent types.Entity) { autoBandwidth.parent = parent }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetParent() types.Entity { return autoBandwidth.parent }

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth) GetParentYangName() string { return "bandwidth" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config
// Configuration parameters relating to MPLS
// auto-bandwidth on the tunnel.
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config struct {
    parent types.Entity
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "min-bw" { return "MinBw" }
    if yname == "max-bw" { return "MaxBw" }
    if yname == "adjust-interval" { return "AdjustInterval" }
    if yname == "adjust-threshold" { return "AdjustThreshold" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["min-bw"] = config.MinBw
    leafs["max-bw"] = config.MaxBw
    leafs["adjust-interval"] = config.AdjustInterval
    leafs["adjust-threshold"] = config.AdjustThreshold
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Config) GetParentYangName() string { return "auto-bandwidth" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State
// State parameters relating to MPLS
// auto-bandwidth on the tunnel.
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State struct {
    parent types.Entity
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "min-bw" { return "MinBw" }
    if yname == "max-bw" { return "MaxBw" }
    if yname == "adjust-interval" { return "AdjustInterval" }
    if yname == "adjust-threshold" { return "AdjustThreshold" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["min-bw"] = state.MinBw
    leafs["max-bw"] = state.MaxBw
    leafs["adjust-interval"] = state.AdjustInterval
    leafs["adjust-threshold"] = state.AdjustThreshold
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_State) GetParentYangName() string { return "auto-bandwidth" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow
// configuration of MPLS overflow bandwidth
// adjustement for the LSP
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config information for MPLS overflow bandwidth adjustment.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config

    // Config information for MPLS overflow bandwidth adjustment.
    State Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State
}

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetFilter() yfilter.YFilter { return overflow.YFilter }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) SetFilter(yf yfilter.YFilter) { overflow.YFilter = yf }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetSegmentPath() string {
    return "overflow"
}

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &overflow.Config
    }
    if childYangName == "state" {
        return &overflow.State
    }
    return nil
}

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &overflow.Config
    children["state"] = &overflow.State
    return children
}

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetBundleName() string { return "openconfig" }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetYangName() string { return "overflow" }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) SetParent(parent types.Entity) { overflow.parent = parent }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetParent() types.Entity { return overflow.parent }

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetParentYangName() string { return "auto-bandwidth" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config
// Config information for MPLS overflow bandwidth
// adjustment
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config struct {
    parent types.Entity
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "overflow-threshold" { return "OverflowThreshold" }
    if yname == "trigger-event-count" { return "TriggerEventCount" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["overflow-threshold"] = config.OverflowThreshold
    leafs["trigger-event-count"] = config.TriggerEventCount
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetParentYangName() string { return "overflow" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State
// Config information for MPLS overflow bandwidth
// adjustment
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State struct {
    parent types.Entity
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "overflow-threshold" { return "OverflowThreshold" }
    if yname == "trigger-event-count" { return "TriggerEventCount" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["overflow-threshold"] = state.OverflowThreshold
    leafs["trigger-event-count"] = state.TriggerEventCount
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetParentYangName() string { return "overflow" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow
// configuration of MPLS underflow bandwidth
//           adjustement for the LSP
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config information for MPLS underflow bandwidth           adjustment.
    Config Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config

    // State information for MPLS underflow bandwidth           adjustment.
    State Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State
}

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetFilter() yfilter.YFilter { return underflow.YFilter }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) SetFilter(yf yfilter.YFilter) { underflow.YFilter = yf }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetSegmentPath() string {
    return "underflow"
}

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &underflow.Config
    }
    if childYangName == "state" {
        return &underflow.State
    }
    return nil
}

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &underflow.Config
    children["state"] = &underflow.State
    return children
}

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetBundleName() string { return "openconfig" }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetYangName() string { return "underflow" }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) SetParent(parent types.Entity) { underflow.parent = parent }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetParent() types.Entity { return underflow.parent }

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetParentYangName() string { return "auto-bandwidth" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config
// Config information for MPLS underflow bandwidth
//           adjustment
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config struct {
    parent types.Entity
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "underflow-threshold" { return "UnderflowThreshold" }
    if yname == "trigger-event-count" { return "TriggerEventCount" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["underflow-threshold"] = config.UnderflowThreshold
    leafs["trigger-event-count"] = config.TriggerEventCount
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetParentYangName() string { return "underflow" }

// Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State
// State information for MPLS underflow bandwidth
//           adjustment
type Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State struct {
    parent types.Entity
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "underflow-threshold" { return "UnderflowThreshold" }
    if yname == "trigger-event-count" { return "TriggerEventCount" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["underflow-threshold"] = state.UnderflowThreshold
    leafs["trigger-event-count"] = state.TriggerEventCount
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetParentYangName() string { return "underflow" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes
// Parameters related to LSPs of type P2P
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes struct {
    parent types.Entity
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

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetFilter() yfilter.YFilter { return p2PTunnelAttributes.YFilter }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) SetFilter(yf yfilter.YFilter) { p2PTunnelAttributes.YFilter = yf }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "p2p-primary-paths" { return "P2PPrimaryPaths" }
    if yname == "p2p-secondary-paths" { return "P2PSecondaryPaths" }
    return ""
}

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetSegmentPath() string {
    return "p2p-tunnel-attributes"
}

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &p2PTunnelAttributes.Config
    }
    if childYangName == "state" {
        return &p2PTunnelAttributes.State
    }
    if childYangName == "p2p-primary-paths" {
        for _, c := range p2PTunnelAttributes.P2PPrimaryPaths {
            if p2PTunnelAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths{}
        p2PTunnelAttributes.P2PPrimaryPaths = append(p2PTunnelAttributes.P2PPrimaryPaths, child)
        return &p2PTunnelAttributes.P2PPrimaryPaths[len(p2PTunnelAttributes.P2PPrimaryPaths)-1]
    }
    if childYangName == "p2p-secondary-paths" {
        for _, c := range p2PTunnelAttributes.P2PSecondaryPaths {
            if p2PTunnelAttributes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths{}
        p2PTunnelAttributes.P2PSecondaryPaths = append(p2PTunnelAttributes.P2PSecondaryPaths, child)
        return &p2PTunnelAttributes.P2PSecondaryPaths[len(p2PTunnelAttributes.P2PSecondaryPaths)-1]
    }
    return nil
}

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &p2PTunnelAttributes.Config
    children["state"] = &p2PTunnelAttributes.State
    for i := range p2PTunnelAttributes.P2PPrimaryPaths {
        children[p2PTunnelAttributes.P2PPrimaryPaths[i].GetSegmentPath()] = &p2PTunnelAttributes.P2PPrimaryPaths[i]
    }
    for i := range p2PTunnelAttributes.P2PSecondaryPaths {
        children[p2PTunnelAttributes.P2PSecondaryPaths[i].GetSegmentPath()] = &p2PTunnelAttributes.P2PSecondaryPaths[i]
    }
    return children
}

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetBundleName() string { return "openconfig" }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetYangName() string { return "p2p-tunnel-attributes" }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) SetParent(parent types.Entity) { p2PTunnelAttributes.parent = parent }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetParent() types.Entity { return p2PTunnelAttributes.parent }

func (p2PTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes) GetParentYangName() string { return "tunnel" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config
// Configuration parameters for P2P LSPs
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // P2P tunnel destination address. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Destination interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = config.Destination
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_Config) GetParentYangName() string { return "p2p-tunnel-attributes" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State
// State parameters for P2P LSPs
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // P2P tunnel destination address. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Destination interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetGoName(yname string) string {
    if yname == "destination" { return "Destination" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination"] = state.Destination
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_State) GetParentYangName() string { return "p2p-tunnel-attributes" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths
// List of p2p primary paths for a tunnel
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths struct {
    parent types.Entity
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

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetFilter() yfilter.YFilter { return p2PPrimaryPaths.YFilter }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) SetFilter(yf yfilter.YFilter) { p2PPrimaryPaths.YFilter = yf }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "candidate-secondary-paths" { return "CandidateSecondaryPaths" }
    if yname == "admin-groups" { return "AdminGroups" }
    return ""
}

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetSegmentPath() string {
    return "p2p-primary-paths" + "[name='" + fmt.Sprintf("%v", p2PPrimaryPaths.Name) + "']"
}

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &p2PPrimaryPaths.Config
    }
    if childYangName == "state" {
        return &p2PPrimaryPaths.State
    }
    if childYangName == "candidate-secondary-paths" {
        return &p2PPrimaryPaths.CandidateSecondaryPaths
    }
    if childYangName == "admin-groups" {
        return &p2PPrimaryPaths.AdminGroups
    }
    return nil
}

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &p2PPrimaryPaths.Config
    children["state"] = &p2PPrimaryPaths.State
    children["candidate-secondary-paths"] = &p2PPrimaryPaths.CandidateSecondaryPaths
    children["admin-groups"] = &p2PPrimaryPaths.AdminGroups
    return children
}

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = p2PPrimaryPaths.Name
    return leafs
}

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetBundleName() string { return "openconfig" }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetYangName() string { return "p2p-primary-paths" }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) SetParent(parent types.Entity) { p2PPrimaryPaths.parent = parent }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetParent() types.Entity { return p2PPrimaryPaths.parent }

func (p2PPrimaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths) GetParentYangName() string { return "p2p-tunnel-attributes" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config
// Configuration parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LocallyComputedExplicitlyDefinedExternallyQueried.
    // The default value is locally-computed.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "path-computation-method" { return "PathComputationMethod" }
    if yname == "use-cspf" { return "UseCspf" }
    if yname == "cspf-tiebreaker" { return "CspfTiebreaker" }
    if yname == "path-computation-server" { return "PathComputationServer" }
    if yname == "explicit-path-name" { return "ExplicitPathName" }
    if yname == "preference" { return "Preference" }
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    if yname == "retry-timer" { return "RetryTimer" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["path-computation-method"] = config.PathComputationMethod
    leafs["use-cspf"] = config.UseCspf
    leafs["cspf-tiebreaker"] = config.CspfTiebreaker
    leafs["path-computation-server"] = config.PathComputationServer
    leafs["explicit-path-name"] = config.ExplicitPathName
    leafs["preference"] = config.Preference
    leafs["setup-priority"] = config.SetupPriority
    leafs["hold-priority"] = config.HoldPriority
    leafs["retry-timer"] = config.RetryTimer
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_Config) GetParentYangName() string { return "p2p-primary-paths" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State
// State parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LocallyComputedExplicitlyDefinedExternallyQueried.
    // The default value is locally-computed.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "path-computation-method" { return "PathComputationMethod" }
    if yname == "use-cspf" { return "UseCspf" }
    if yname == "cspf-tiebreaker" { return "CspfTiebreaker" }
    if yname == "path-computation-server" { return "PathComputationServer" }
    if yname == "explicit-path-name" { return "ExplicitPathName" }
    if yname == "preference" { return "Preference" }
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    if yname == "retry-timer" { return "RetryTimer" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["path-computation-method"] = state.PathComputationMethod
    leafs["use-cspf"] = state.UseCspf
    leafs["cspf-tiebreaker"] = state.CspfTiebreaker
    leafs["path-computation-server"] = state.PathComputationServer
    leafs["explicit-path-name"] = state.ExplicitPathName
    leafs["preference"] = state.Preference
    leafs["setup-priority"] = state.SetupPriority
    leafs["hold-priority"] = state.HoldPriority
    leafs["retry-timer"] = state.RetryTimer
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_State) GetParentYangName() string { return "p2p-primary-paths" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of secondary paths which may be utilised when the current primary path
    // is in use. The type is slice of
    // Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath.
    CandidateSecondaryPath []Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath
}

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetFilter() yfilter.YFilter { return candidateSecondaryPaths.YFilter }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) SetFilter(yf yfilter.YFilter) { candidateSecondaryPaths.YFilter = yf }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetGoName(yname string) string {
    if yname == "candidate-secondary-path" { return "CandidateSecondaryPath" }
    return ""
}

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetSegmentPath() string {
    return "candidate-secondary-paths"
}

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-secondary-path" {
        for _, c := range candidateSecondaryPaths.CandidateSecondaryPath {
            if candidateSecondaryPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath{}
        candidateSecondaryPaths.CandidateSecondaryPath = append(candidateSecondaryPaths.CandidateSecondaryPath, child)
        return &candidateSecondaryPaths.CandidateSecondaryPath[len(candidateSecondaryPaths.CandidateSecondaryPath)-1]
    }
    return nil
}

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range candidateSecondaryPaths.CandidateSecondaryPath {
        children[candidateSecondaryPaths.CandidateSecondaryPath[i].GetSegmentPath()] = &candidateSecondaryPaths.CandidateSecondaryPath[i]
    }
    return children
}

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetBundleName() string { return "openconfig" }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetYangName() string { return "candidate-secondary-paths" }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) SetParent(parent types.Entity) { candidateSecondaryPaths.parent = parent }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetParent() types.Entity { return candidateSecondaryPaths.parent }

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths) GetParentYangName() string { return "p2p-primary-paths" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath
// List of secondary paths which may be utilised when the
// current primary path is in use
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath struct {
    parent types.Entity
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

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetFilter() yfilter.YFilter { return candidateSecondaryPath.YFilter }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) SetFilter(yf yfilter.YFilter) { candidateSecondaryPath.YFilter = yf }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetGoName(yname string) string {
    if yname == "secondary-path" { return "SecondaryPath" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetSegmentPath() string {
    return "candidate-secondary-path" + "[secondary-path='" + fmt.Sprintf("%v", candidateSecondaryPath.SecondaryPath) + "']"
}

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &candidateSecondaryPath.Config
    }
    if childYangName == "state" {
        return &candidateSecondaryPath.State
    }
    return nil
}

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &candidateSecondaryPath.Config
    children["state"] = &candidateSecondaryPath.State
    return children
}

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["secondary-path"] = candidateSecondaryPath.SecondaryPath
    return leafs
}

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetBundleName() string { return "openconfig" }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetYangName() string { return "candidate-secondary-path" }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) SetParent(parent types.Entity) { candidateSecondaryPath.parent = parent }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetParent() types.Entity { return candidateSecondaryPath.parent }

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath) GetParentYangName() string { return "candidate-secondary-paths" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config
// Configuration parameters relating to the candidate
// secondary path
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config struct {
    parent types.Entity
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetGoName(yname string) string {
    if yname == "secondary-path" { return "SecondaryPath" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["secondary-path"] = config.SecondaryPath
    leafs["priority"] = config.Priority
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetParentYangName() string { return "candidate-secondary-path" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State
// Operational state parameters relating to the candidate
// secondary path
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State struct {
    parent types.Entity
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetGoName(yname string) string {
    if yname == "secondary-path" { return "SecondaryPath" }
    if yname == "priority" { return "Priority" }
    if yname == "active" { return "Active" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["secondary-path"] = state.SecondaryPath
    leafs["priority"] = state.Priority
    leafs["active"] = state.Active
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetParentYangName() string { return "candidate-secondary-path" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups
// Top-level container for include/exclude constraints for
// link affinities
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data .
    Config Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config

    // Operational state data .
    State Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetFilter() yfilter.YFilter { return adminGroups.YFilter }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) SetFilter(yf yfilter.YFilter) { adminGroups.YFilter = yf }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetSegmentPath() string {
    return "admin-groups"
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &adminGroups.Config
    }
    if childYangName == "state" {
        return &adminGroups.State
    }
    return nil
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &adminGroups.Config
    children["state"] = &adminGroups.State
    return children
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetBundleName() string { return "openconfig" }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetYangName() string { return "admin-groups" }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) SetParent(parent types.Entity) { adminGroups.parent = parent }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetParent() types.Entity { return adminGroups.parent }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups) GetParentYangName() string { return "p2p-primary-paths" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config
// Configuration data 
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config struct {
    parent types.Entity
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetGoName(yname string) string {
    if yname == "exclude-group" { return "ExcludeGroup" }
    if yname == "include-all-group" { return "IncludeAllGroup" }
    if yname == "include-any-group" { return "IncludeAnyGroup" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exclude-group"] = config.ExcludeGroup
    leafs["include-all-group"] = config.IncludeAllGroup
    leafs["include-any-group"] = config.IncludeAnyGroup
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_Config) GetParentYangName() string { return "admin-groups" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State
// Operational state data 
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State struct {
    parent types.Entity
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetGoName(yname string) string {
    if yname == "exclude-group" { return "ExcludeGroup" }
    if yname == "include-all-group" { return "IncludeAllGroup" }
    if yname == "include-any-group" { return "IncludeAnyGroup" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exclude-group"] = state.ExcludeGroup
    leafs["include-all-group"] = state.IncludeAllGroup
    leafs["include-any-group"] = state.IncludeAnyGroup
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PPrimaryPaths_AdminGroups_State) GetParentYangName() string { return "admin-groups" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths
// List of p2p primary paths for a tunnel
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths struct {
    parent types.Entity
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

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetFilter() yfilter.YFilter { return p2PSecondaryPaths.YFilter }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) SetFilter(yf yfilter.YFilter) { p2PSecondaryPaths.YFilter = yf }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "admin-groups" { return "AdminGroups" }
    return ""
}

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetSegmentPath() string {
    return "p2p-secondary-paths" + "[name='" + fmt.Sprintf("%v", p2PSecondaryPaths.Name) + "']"
}

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &p2PSecondaryPaths.Config
    }
    if childYangName == "state" {
        return &p2PSecondaryPaths.State
    }
    if childYangName == "admin-groups" {
        return &p2PSecondaryPaths.AdminGroups
    }
    return nil
}

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &p2PSecondaryPaths.Config
    children["state"] = &p2PSecondaryPaths.State
    children["admin-groups"] = &p2PSecondaryPaths.AdminGroups
    return children
}

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = p2PSecondaryPaths.Name
    return leafs
}

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetBundleName() string { return "openconfig" }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetYangName() string { return "p2p-secondary-paths" }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) SetParent(parent types.Entity) { p2PSecondaryPaths.parent = parent }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetParent() types.Entity { return p2PSecondaryPaths.parent }

func (p2PSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths) GetParentYangName() string { return "p2p-tunnel-attributes" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config
// Configuration parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LocallyComputedExplicitlyDefinedExternallyQueried.
    // The default value is locally-computed.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "path-computation-method" { return "PathComputationMethod" }
    if yname == "use-cspf" { return "UseCspf" }
    if yname == "cspf-tiebreaker" { return "CspfTiebreaker" }
    if yname == "path-computation-server" { return "PathComputationServer" }
    if yname == "explicit-path-name" { return "ExplicitPathName" }
    if yname == "preference" { return "Preference" }
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    if yname == "retry-timer" { return "RetryTimer" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = config.Name
    leafs["path-computation-method"] = config.PathComputationMethod
    leafs["use-cspf"] = config.UseCspf
    leafs["cspf-tiebreaker"] = config.CspfTiebreaker
    leafs["path-computation-server"] = config.PathComputationServer
    leafs["explicit-path-name"] = config.ExplicitPathName
    leafs["preference"] = config.Preference
    leafs["setup-priority"] = config.SetupPriority
    leafs["hold-priority"] = config.HoldPriority
    leafs["retry-timer"] = config.RetryTimer
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_Config) GetParentYangName() string { return "p2p-secondary-paths" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State
// State parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LocallyComputedExplicitlyDefinedExternallyQueried.
    // The default value is locally-computed.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "path-computation-method" { return "PathComputationMethod" }
    if yname == "use-cspf" { return "UseCspf" }
    if yname == "cspf-tiebreaker" { return "CspfTiebreaker" }
    if yname == "path-computation-server" { return "PathComputationServer" }
    if yname == "explicit-path-name" { return "ExplicitPathName" }
    if yname == "preference" { return "Preference" }
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    if yname == "retry-timer" { return "RetryTimer" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = state.Name
    leafs["path-computation-method"] = state.PathComputationMethod
    leafs["use-cspf"] = state.UseCspf
    leafs["cspf-tiebreaker"] = state.CspfTiebreaker
    leafs["path-computation-server"] = state.PathComputationServer
    leafs["explicit-path-name"] = state.ExplicitPathName
    leafs["preference"] = state.Preference
    leafs["setup-priority"] = state.SetupPriority
    leafs["hold-priority"] = state.HoldPriority
    leafs["retry-timer"] = state.RetryTimer
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_State) GetParentYangName() string { return "p2p-secondary-paths" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups
// Top-level container for include/exclude constraints for
// link affinities
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data .
    Config Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config

    // Operational state data .
    State Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetFilter() yfilter.YFilter { return adminGroups.YFilter }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) SetFilter(yf yfilter.YFilter) { adminGroups.YFilter = yf }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetSegmentPath() string {
    return "admin-groups"
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &adminGroups.Config
    }
    if childYangName == "state" {
        return &adminGroups.State
    }
    return nil
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &adminGroups.Config
    children["state"] = &adminGroups.State
    return children
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetBundleName() string { return "openconfig" }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetYangName() string { return "admin-groups" }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) SetParent(parent types.Entity) { adminGroups.parent = parent }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetParent() types.Entity { return adminGroups.parent }

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups) GetParentYangName() string { return "p2p-secondary-paths" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config
// Configuration data 
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config struct {
    parent types.Entity
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetGoName(yname string) string {
    if yname == "exclude-group" { return "ExcludeGroup" }
    if yname == "include-all-group" { return "IncludeAllGroup" }
    if yname == "include-any-group" { return "IncludeAnyGroup" }
    return ""
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exclude-group"] = config.ExcludeGroup
    leafs["include-all-group"] = config.IncludeAllGroup
    leafs["include-any-group"] = config.IncludeAnyGroup
    return leafs
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_Config) GetParentYangName() string { return "admin-groups" }

// Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State
// Operational state data 
type Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State struct {
    parent types.Entity
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetGoName(yname string) string {
    if yname == "exclude-group" { return "ExcludeGroup" }
    if yname == "include-all-group" { return "IncludeAllGroup" }
    if yname == "include-any-group" { return "IncludeAnyGroup" }
    return ""
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exclude-group"] = state.ExcludeGroup
    leafs["include-all-group"] = state.IncludeAllGroup
    leafs["include-any-group"] = state.IncludeAnyGroup
    return leafs
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_ConstrainedPath_Tunnel_P2PTunnelAttributes_P2PSecondaryPaths_AdminGroups_State) GetParentYangName() string { return "admin-groups" }

// Mpls_Lsps_UnconstrainedPath
// LSPs that use the IGP-determined path, i.e., non
// traffic-engineered, or non constrained-path
type Mpls_Lsps_UnconstrainedPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // select and configure the signaling method for  the LSP.
    PathSetupProtocol Mpls_Lsps_UnconstrainedPath_PathSetupProtocol
}

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetFilter() yfilter.YFilter { return unconstrainedPath.YFilter }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) SetFilter(yf yfilter.YFilter) { unconstrainedPath.YFilter = yf }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetGoName(yname string) string {
    if yname == "path-setup-protocol" { return "PathSetupProtocol" }
    return ""
}

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetSegmentPath() string {
    return "unconstrained-path"
}

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path-setup-protocol" {
        return &unconstrainedPath.PathSetupProtocol
    }
    return nil
}

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["path-setup-protocol"] = &unconstrainedPath.PathSetupProtocol
    return children
}

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetBundleName() string { return "openconfig" }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetYangName() string { return "unconstrained-path" }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) SetParent(parent types.Entity) { unconstrainedPath.parent = parent }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetParent() types.Entity { return unconstrainedPath.parent }

func (unconstrainedPath *Mpls_Lsps_UnconstrainedPath) GetParentYangName() string { return "lsps" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol
// select and configure the signaling method for
//  the LSP
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP signaling setup for IGP-congruent LSPs.
    Ldp Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp

    // segment routing signaling extensions for IGP-confgruent LSPs.
    SegmentRouting Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting
}

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetFilter() yfilter.YFilter { return pathSetupProtocol.YFilter }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) SetFilter(yf yfilter.YFilter) { pathSetupProtocol.YFilter = yf }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetGoName(yname string) string {
    if yname == "ldp" { return "Ldp" }
    if yname == "segment-routing" { return "SegmentRouting" }
    return ""
}

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetSegmentPath() string {
    return "path-setup-protocol"
}

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ldp" {
        return &pathSetupProtocol.Ldp
    }
    if childYangName == "segment-routing" {
        return &pathSetupProtocol.SegmentRouting
    }
    return nil
}

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ldp"] = &pathSetupProtocol.Ldp
    children["segment-routing"] = &pathSetupProtocol.SegmentRouting
    return children
}

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetBundleName() string { return "openconfig" }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetYangName() string { return "path-setup-protocol" }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) SetParent(parent types.Entity) { pathSetupProtocol.parent = parent }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetParent() types.Entity { return pathSetupProtocol.parent }

func (pathSetupProtocol *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol) GetParentYangName() string { return "unconstrained-path" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp
// LDP signaling setup for IGP-congruent LSPs
// This type is a presence type.
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // contains configuration stanzas for different LSP tunnel types (P2P, P2MP,
    // etc.).
    Tunnel Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel
}

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetFilter() yfilter.YFilter { return ldp.YFilter }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) SetFilter(yf yfilter.YFilter) { ldp.YFilter = yf }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetSegmentPath() string {
    return "ldp"
}

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel" {
        return &ldp.Tunnel
    }
    return nil
}

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tunnel"] = &ldp.Tunnel
    return children
}

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetBundleName() string { return "openconfig" }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetYangName() string { return "ldp" }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) SetParent(parent types.Entity) { ldp.parent = parent }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetParent() types.Entity { return ldp.parent }

func (ldp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp) GetParentYangName() string { return "path-setup-protocol" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel
// contains configuration stanzas for different LSP
// tunnel types (P2P, P2MP, etc.)
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // specifies the type of LSP, e.g., P2P or P2MP. The type is TunnelType.
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

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetGoName(yname string) string {
    if yname == "tunnel-type" { return "TunnelType" }
    if yname == "ldp-type" { return "LdpType" }
    if yname == "p2p-lsp" { return "P2PLsp" }
    if yname == "p2mp-lsp" { return "P2MpLsp" }
    if yname == "mp2mp-lsp" { return "Mp2MpLsp" }
    return ""
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetSegmentPath() string {
    return "tunnel"
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "p2p-lsp" {
        return &tunnel.P2PLsp
    }
    if childYangName == "p2mp-lsp" {
        return &tunnel.P2MpLsp
    }
    if childYangName == "mp2mp-lsp" {
        return &tunnel.Mp2MpLsp
    }
    return nil
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["p2p-lsp"] = &tunnel.P2PLsp
    children["p2mp-lsp"] = &tunnel.P2MpLsp
    children["mp2mp-lsp"] = &tunnel.Mp2MpLsp
    return children
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnel-type"] = tunnel.TunnelType
    leafs["ldp-type"] = tunnel.LdpType
    return leafs
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetBundleName() string { return "openconfig" }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel) GetParentYangName() string { return "ldp" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp
// properties of point-to-point tunnels
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address prefix for packets sharing the same forwarding equivalence class
    // for the IGP-based LSP. The type is one of the following types: slice of
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    FecAddress []interface{}
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetFilter() yfilter.YFilter { return p2PLsp.YFilter }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) SetFilter(yf yfilter.YFilter) { p2PLsp.YFilter = yf }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetGoName(yname string) string {
    if yname == "fec-address" { return "FecAddress" }
    return ""
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetSegmentPath() string {
    return "p2p-lsp"
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fec-address"] = p2PLsp.FecAddress
    return leafs
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetBundleName() string { return "openconfig" }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetYangName() string { return "p2p-lsp" }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) SetParent(parent types.Entity) { p2PLsp.parent = parent }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetParent() types.Entity { return p2PLsp.parent }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2PLsp) GetParentYangName() string { return "tunnel" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp
// properties of point-to-multipoint tunnels
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetFilter() yfilter.YFilter { return p2MpLsp.YFilter }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) SetFilter(yf yfilter.YFilter) { p2MpLsp.YFilter = yf }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetGoName(yname string) string {
    return ""
}

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetSegmentPath() string {
    return "p2mp-lsp"
}

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetBundleName() string { return "openconfig" }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetYangName() string { return "p2mp-lsp" }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) SetParent(parent types.Entity) { p2MpLsp.parent = parent }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetParent() types.Entity { return p2MpLsp.parent }

func (p2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_P2MpLsp) GetParentYangName() string { return "tunnel" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp
// properties of multipoint-to-multipoint tunnels
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetFilter() yfilter.YFilter { return mp2MpLsp.YFilter }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) SetFilter(yf yfilter.YFilter) { mp2MpLsp.YFilter = yf }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetGoName(yname string) string {
    return ""
}

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetSegmentPath() string {
    return "mp2mp-lsp"
}

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetBundleName() string { return "openconfig" }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetYangName() string { return "mp2mp-lsp" }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) SetParent(parent types.Entity) { mp2MpLsp.parent = parent }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetParent() types.Entity { return mp2MpLsp.parent }

func (mp2MpLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp_Tunnel_Mp2MpLsp) GetParentYangName() string { return "tunnel" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // contains configuration stanzas for different LSP tunnel types (P2P, P2MP,
    // etc.).
    Tunnel Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel
}

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetFilter() yfilter.YFilter { return segmentRouting.YFilter }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) SetFilter(yf yfilter.YFilter) { segmentRouting.YFilter = yf }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetGoName(yname string) string {
    if yname == "tunnel" { return "Tunnel" }
    return ""
}

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetSegmentPath() string {
    return "segment-routing"
}

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel" {
        return &segmentRouting.Tunnel
    }
    return nil
}

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tunnel"] = &segmentRouting.Tunnel
    return children
}

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetBundleName() string { return "openconfig" }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetYangName() string { return "segment-routing" }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) SetParent(parent types.Entity) { segmentRouting.parent = parent }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetParent() types.Entity { return segmentRouting.parent }

func (segmentRouting *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting) GetParentYangName() string { return "path-setup-protocol" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel
// contains configuration stanzas for different LSP
// tunnel types (P2P, P2MP, etc.)
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // specifies the type of LSP, e.g., P2P or P2MP. The type is TunnelType.
    TunnelType interface{}

    // properties of point-to-point tunnels.
    P2PLsp Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetGoName(yname string) string {
    if yname == "tunnel-type" { return "TunnelType" }
    if yname == "p2p-lsp" { return "P2PLsp" }
    return ""
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetSegmentPath() string {
    return "tunnel"
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "p2p-lsp" {
        return &tunnel.P2PLsp
    }
    return nil
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["p2p-lsp"] = &tunnel.P2PLsp
    return children
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnel-type"] = tunnel.TunnelType
    return leafs
}

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetBundleName() string { return "openconfig" }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel) GetParentYangName() string { return "segment-routing" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp
// properties of point-to-point tunnels
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of FECs that are to be originated as SR LSPs. The type is slice of
    // Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec.
    Fec []Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetFilter() yfilter.YFilter { return p2PLsp.YFilter }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) SetFilter(yf yfilter.YFilter) { p2PLsp.YFilter = yf }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetGoName(yname string) string {
    if yname == "fec" { return "Fec" }
    return ""
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetSegmentPath() string {
    return "p2p-lsp"
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fec" {
        for _, c := range p2PLsp.Fec {
            if p2PLsp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec{}
        p2PLsp.Fec = append(p2PLsp.Fec, child)
        return &p2PLsp.Fec[len(p2PLsp.Fec)-1]
    }
    return nil
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range p2PLsp.Fec {
        children[p2PLsp.Fec[i].GetSegmentPath()] = &p2PLsp.Fec[i]
    }
    return children
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetBundleName() string { return "openconfig" }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetYangName() string { return "p2p-lsp" }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) SetParent(parent types.Entity) { p2PLsp.parent = parent }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetParent() types.Entity { return p2PLsp.parent }

func (p2PLsp *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp) GetParentYangName() string { return "tunnel" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec
// List of FECs that are to be originated as SR LSPs
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. FEC that is to be advertised as part of the
    // Prefix-SID. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    FecAddress interface{}

    // Configuration parameters relating to the FEC to be advertised by SR.
    Config Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config

    // Operational state relating to a FEC advertised by SR.
    State Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State

    // Parameters relating to the Prefix-SID used for the originated FEC.
    PrefixSid Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid
}

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetFilter() yfilter.YFilter { return fec.YFilter }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) SetFilter(yf yfilter.YFilter) { fec.YFilter = yf }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetGoName(yname string) string {
    if yname == "fec-address" { return "FecAddress" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "prefix-sid" { return "PrefixSid" }
    return ""
}

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetSegmentPath() string {
    return "fec" + "[fec-address='" + fmt.Sprintf("%v", fec.FecAddress) + "']"
}

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &fec.Config
    }
    if childYangName == "state" {
        return &fec.State
    }
    if childYangName == "prefix-sid" {
        return &fec.PrefixSid
    }
    return nil
}

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &fec.Config
    children["state"] = &fec.State
    children["prefix-sid"] = &fec.PrefixSid
    return children
}

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fec-address"] = fec.FecAddress
    return leafs
}

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetBundleName() string { return "openconfig" }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetYangName() string { return "fec" }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) SetParent(parent types.Entity) { fec.parent = parent }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetParent() types.Entity { return fec.parent }

func (fec *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec) GetParentYangName() string { return "p2p-lsp" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config
// Configuration parameters relating to the FEC to be
// advertised by SR
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FEC that is to be advertised as part of the Prefix-SID. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    FecAddress interface{}
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetGoName(yname string) string {
    if yname == "fec-address" { return "FecAddress" }
    return ""
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fec-address"] = config.FecAddress
    return leafs
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_Config) GetParentYangName() string { return "fec" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State
// Operational state relating to a FEC advertised by SR
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FEC that is to be advertised as part of the Prefix-SID. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    FecAddress interface{}
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetGoName(yname string) string {
    if yname == "fec-address" { return "FecAddress" }
    return ""
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fec-address"] = state.FecAddress
    return leafs
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_State) GetParentYangName() string { return "fec" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid
// Parameters relating to the Prefix-SID
// used for the originated FEC
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters relating to the Prefix-SID used for the originated
    // FEC.
    Config Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config

    // Operational state parameters relating to the Prefix-SID used for the
    // originated FEC.
    State Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State
}

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetFilter() yfilter.YFilter { return prefixSid.YFilter }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) SetFilter(yf yfilter.YFilter) { prefixSid.YFilter = yf }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetSegmentPath() string {
    return "prefix-sid"
}

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &prefixSid.Config
    }
    if childYangName == "state" {
        return &prefixSid.State
    }
    return nil
}

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &prefixSid.Config
    children["state"] = &prefixSid.State
    return children
}

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetBundleName() string { return "openconfig" }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetYangName() string { return "prefix-sid" }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) SetParent(parent types.Entity) { prefixSid.parent = parent }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetParent() types.Entity { return prefixSid.parent }

func (prefixSid *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid) GetParentYangName() string { return "fec" }

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config
// Configuration parameters relating to the Prefix-SID
// used for the originated FEC
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifies how the value of the Prefix-SID should be interpreted - whether
    // as an offset to the SRGB, or as an absolute value. The type is Type. The
    // default value is INDEX.
    Type interface{}

    // Specifies that the Prefix-SID is to be treated as a Node-SID by setting the
    // N-flag in the advertised Prefix-SID TLV in the IGP. The type is bool.
    NodeFlag interface{}

    // Configuration relating to the LFIB actions for the Prefix-SID to be used by
    // the penultimate-hop. The type is LastHopBehavior.
    LastHopBehavior interface{}
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "node-flag" { return "NodeFlag" }
    if yname == "last-hop-behavior" { return "LastHopBehavior" }
    return ""
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetSegmentPath() string {
    return "config"
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = config.Type
    leafs["node-flag"] = config.NodeFlag
    leafs["last-hop-behavior"] = config.LastHopBehavior
    return leafs
}

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetBundleName() string { return "openconfig" }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetYangName() string { return "config" }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetParent() types.Entity { return config.parent }

func (config *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config) GetParentYangName() string { return "prefix-sid" }

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

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type represents absolute value
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type string

const (
    // Set when the value of the prefix SID should be specified
    // as an off-set from the SRGB's zero-value. When multiple
    // SRGBs are specified, the zero-value is the minimum
    // of their lower bounds
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type_INDEX Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type = "INDEX"

    // Set when the value of a prefix SID is specified as the
    // absolute value within an SRGB. It is an error to specify
    // an absolute value outside of a specified SRGB
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type_ABSOLUTE Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_Config_Type = "ABSOLUTE"
)

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State
// Operational state parameters relating to the
// Prefix-SID used for the originated FEC
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifies how the value of the Prefix-SID should be interpreted - whether
    // as an offset to the SRGB, or as an absolute value. The type is Type. The
    // default value is INDEX.
    Type interface{}

    // Specifies that the Prefix-SID is to be treated as a Node-SID by setting the
    // N-flag in the advertised Prefix-SID TLV in the IGP. The type is bool.
    NodeFlag interface{}

    // Configuration relating to the LFIB actions for the Prefix-SID to be used by
    // the penultimate-hop. The type is LastHopBehavior.
    LastHopBehavior interface{}
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "node-flag" { return "NodeFlag" }
    if yname == "last-hop-behavior" { return "LastHopBehavior" }
    return ""
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetSegmentPath() string {
    return "state"
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = state.Type
    leafs["node-flag"] = state.NodeFlag
    leafs["last-hop-behavior"] = state.LastHopBehavior
    return leafs
}

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetBundleName() string { return "openconfig" }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetYangName() string { return "state" }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetParent() types.Entity { return state.parent }

func (state *Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State) GetParentYangName() string { return "prefix-sid" }

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

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type represents absolute value
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type string

const (
    // Set when the value of the prefix SID should be specified
    // as an off-set from the SRGB's zero-value. When multiple
    // SRGBs are specified, the zero-value is the minimum
    // of their lower bounds
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type_INDEX Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type = "INDEX"

    // Set when the value of a prefix SID is specified as the
    // absolute value within an SRGB. It is an error to specify
    // an absolute value outside of a specified SRGB
    Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type_ABSOLUTE Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_SegmentRouting_Tunnel_P2PLsp_Fec_PrefixSid_State_Type = "ABSOLUTE"
)

// Mpls_Lsps_StaticLsps
// statically configured LSPs, without dynamic
// signaling
type Mpls_Lsps_StaticLsps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // list of defined static LSPs. The type is slice of
    // Mpls_Lsps_StaticLsps_LabelSwitchedPath.
    LabelSwitchedPath []Mpls_Lsps_StaticLsps_LabelSwitchedPath
}

func (staticLsps *Mpls_Lsps_StaticLsps) GetFilter() yfilter.YFilter { return staticLsps.YFilter }

func (staticLsps *Mpls_Lsps_StaticLsps) SetFilter(yf yfilter.YFilter) { staticLsps.YFilter = yf }

func (staticLsps *Mpls_Lsps_StaticLsps) GetGoName(yname string) string {
    if yname == "label-switched-path" { return "LabelSwitchedPath" }
    return ""
}

func (staticLsps *Mpls_Lsps_StaticLsps) GetSegmentPath() string {
    return "static-lsps"
}

func (staticLsps *Mpls_Lsps_StaticLsps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-switched-path" {
        for _, c := range staticLsps.LabelSwitchedPath {
            if staticLsps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpls_Lsps_StaticLsps_LabelSwitchedPath{}
        staticLsps.LabelSwitchedPath = append(staticLsps.LabelSwitchedPath, child)
        return &staticLsps.LabelSwitchedPath[len(staticLsps.LabelSwitchedPath)-1]
    }
    return nil
}

func (staticLsps *Mpls_Lsps_StaticLsps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticLsps.LabelSwitchedPath {
        children[staticLsps.LabelSwitchedPath[i].GetSegmentPath()] = &staticLsps.LabelSwitchedPath[i]
    }
    return children
}

func (staticLsps *Mpls_Lsps_StaticLsps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticLsps *Mpls_Lsps_StaticLsps) GetBundleName() string { return "openconfig" }

func (staticLsps *Mpls_Lsps_StaticLsps) GetYangName() string { return "static-lsps" }

func (staticLsps *Mpls_Lsps_StaticLsps) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (staticLsps *Mpls_Lsps_StaticLsps) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (staticLsps *Mpls_Lsps_StaticLsps) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (staticLsps *Mpls_Lsps_StaticLsps) SetParent(parent types.Entity) { staticLsps.parent = parent }

func (staticLsps *Mpls_Lsps_StaticLsps) GetParent() types.Entity { return staticLsps.parent }

func (staticLsps *Mpls_Lsps_StaticLsps) GetParentYangName() string { return "lsps" }

// Mpls_Lsps_StaticLsps_LabelSwitchedPath
// list of defined static LSPs
type Mpls_Lsps_StaticLsps_LabelSwitchedPath struct {
    parent types.Entity
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

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetFilter() yfilter.YFilter { return labelSwitchedPath.YFilter }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) SetFilter(yf yfilter.YFilter) { labelSwitchedPath.YFilter = yf }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "ingress" { return "Ingress" }
    if yname == "transit" { return "Transit" }
    if yname == "egress" { return "Egress" }
    return ""
}

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetSegmentPath() string {
    return "label-switched-path" + "[name='" + fmt.Sprintf("%v", labelSwitchedPath.Name) + "']"
}

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ingress" {
        return &labelSwitchedPath.Ingress
    }
    if childYangName == "transit" {
        return &labelSwitchedPath.Transit
    }
    if childYangName == "egress" {
        return &labelSwitchedPath.Egress
    }
    return nil
}

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ingress"] = &labelSwitchedPath.Ingress
    children["transit"] = &labelSwitchedPath.Transit
    children["egress"] = &labelSwitchedPath.Egress
    return children
}

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = labelSwitchedPath.Name
    return leafs
}

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetBundleName() string { return "openconfig" }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetYangName() string { return "label-switched-path" }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) SetParent(parent types.Entity) { labelSwitchedPath.parent = parent }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetParent() types.Entity { return labelSwitchedPath.parent }

func (labelSwitchedPath *Mpls_Lsps_StaticLsps_LabelSwitchedPath) GetParentYangName() string { return "static-lsps" }

// Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress
// Static LSPs for which the router is an
// ingress node
type Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetFilter() yfilter.YFilter { return ingress.YFilter }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) SetFilter(yf yfilter.YFilter) { ingress.YFilter = yf }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "incoming-label" { return "IncomingLabel" }
    if yname == "push-label" { return "PushLabel" }
    return ""
}

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetSegmentPath() string {
    return "ingress"
}

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = ingress.NextHop
    leafs["incoming-label"] = ingress.IncomingLabel
    leafs["push-label"] = ingress.PushLabel
    return leafs
}

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetBundleName() string { return "openconfig" }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetYangName() string { return "ingress" }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) SetParent(parent types.Entity) { ingress.parent = parent }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetParent() types.Entity { return ingress.parent }

func (ingress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Ingress) GetParentYangName() string { return "label-switched-path" }

// Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit
// static LSPs for which the router is a
// transit node
type Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetFilter() yfilter.YFilter { return transit.YFilter }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) SetFilter(yf yfilter.YFilter) { transit.YFilter = yf }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "incoming-label" { return "IncomingLabel" }
    if yname == "push-label" { return "PushLabel" }
    return ""
}

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetSegmentPath() string {
    return "transit"
}

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = transit.NextHop
    leafs["incoming-label"] = transit.IncomingLabel
    leafs["push-label"] = transit.PushLabel
    return leafs
}

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetBundleName() string { return "openconfig" }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetYangName() string { return "transit" }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) SetParent(parent types.Entity) { transit.parent = parent }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetParent() types.Entity { return transit.parent }

func (transit *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Transit) GetParentYangName() string { return "label-switched-path" }

// Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress
// static LSPs for which the router is a
// egress  node
type Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetFilter() yfilter.YFilter { return egress.YFilter }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) SetFilter(yf yfilter.YFilter) { egress.YFilter = yf }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "incoming-label" { return "IncomingLabel" }
    if yname == "push-label" { return "PushLabel" }
    return ""
}

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetSegmentPath() string {
    return "egress"
}

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = egress.NextHop
    leafs["incoming-label"] = egress.IncomingLabel
    leafs["push-label"] = egress.PushLabel
    return leafs
}

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetBundleName() string { return "openconfig" }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetYangName() string { return "egress" }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) SetParent(parent types.Entity) { egress.parent = parent }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetParent() types.Entity { return egress.parent }

func (egress *Mpls_Lsps_StaticLsps_LabelSwitchedPath_Egress) GetParentYangName() string { return "label-switched-path" }

