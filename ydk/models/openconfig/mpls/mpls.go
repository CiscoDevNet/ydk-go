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

// MplsSrlgFloodingType represents Enumerated bype for specifying how the SRLG is flooded
type MplsSrlgFloodingType string

const (
    // SRLG is flooded in the IGP
    MplsSrlgFloodingType_FLOODED_SRLG MplsSrlgFloodingType = "FLOODED_SRLG"

    // SRLG is not flooded, the members are
    // statically configured
    MplsSrlgFloodingType_STATIC_SRLG MplsSrlgFloodingType = "STATIC_SRLG"
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

    mpls.EntityData.Children = types.NewOrderedMap()
    mpls.EntityData.Children.Append("global", types.YChild{"Global", &mpls.Global})
    mpls.EntityData.Children.Append("te-global-attributes", types.YChild{"TeGlobalAttributes", &mpls.TeGlobalAttributes})
    mpls.EntityData.Children.Append("te-interface-attributes", types.YChild{"TeInterfaceAttributes", &mpls.TeInterfaceAttributes})
    mpls.EntityData.Children.Append("signaling-protocols", types.YChild{"SignalingProtocols", &mpls.SignalingProtocols})
    mpls.EntityData.Children.Append("lsps", types.YChild{"Lsps", &mpls.Lsps})
    mpls.EntityData.Leafs = types.NewOrderedMap()

    mpls.EntityData.YListKeys = []string {}

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
    InterfaceAttributes Mpls_Global_InterfaceAttributes

    // A range of labels starting with the start-label and up-to and including the
    // end label that should be allocated as reserved. These labels should not be
    // utilised by any dynamic label allocation on the local system unless the
    // allocating protocol is explicitly configured to specify that allocation of
    // labels should be out of the label block specified.
    ReservedLabelBlocks Mpls_Global_ReservedLabelBlocks
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

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("config", types.YChild{"Config", &global.Config})
    global.EntityData.Children.Append("state", types.YChild{"State", &global.State})
    global.EntityData.Children.Append("interface-attributes", types.YChild{"InterfaceAttributes", &global.InterfaceAttributes})
    global.EntityData.Children.Append("reserved-label-blocks", types.YChild{"ReservedLabelBlocks", &global.ReservedLabelBlocks})
    global.EntityData.Leafs = types.NewOrderedMap()

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// Mpls_Global_Config
// Top level global MPLS configuration
type Mpls_Global_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The null-label type used, implicit or explicit. The type is one of the
    // following: EXPLICITIMPLICIT. The default value is oc-mplst:IMPLICIT.
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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("null-label", types.YLeaf{"NullLabel", config.NullLabel})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Global_State
// Top level global MPLS state
type Mpls_Global_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The null-label type used, implicit or explicit. The type is one of the
    // following: EXPLICITIMPLICIT. The default value is oc-mplst:IMPLICIT.
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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("null-label", types.YLeaf{"NullLabel", state.NullLabel})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Global_InterfaceAttributes
// Parameters related to MPLS interfaces
type Mpls_Global_InterfaceAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of TE interfaces. The type is slice of
    // Mpls_Global_InterfaceAttributes_Interface.
    Interface []*Mpls_Global_InterfaceAttributes_Interface
}

func (interfaceAttributes *Mpls_Global_InterfaceAttributes) GetEntityData() *types.CommonEntityData {
    interfaceAttributes.EntityData.YFilter = interfaceAttributes.YFilter
    interfaceAttributes.EntityData.YangName = "interface-attributes"
    interfaceAttributes.EntityData.BundleName = "openconfig"
    interfaceAttributes.EntityData.ParentYangName = "global"
    interfaceAttributes.EntityData.SegmentPath = "interface-attributes"
    interfaceAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceAttributes.EntityData.Children = types.NewOrderedMap()
    interfaceAttributes.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceAttributes.Interface {
        interfaceAttributes.EntityData.Children.Append(types.GetSegmentPath(interfaceAttributes.Interface[i]), types.YChild{"Interface", interfaceAttributes.Interface[i]})
    }
    interfaceAttributes.EntityData.Leafs = types.NewOrderedMap()

    interfaceAttributes.EntityData.YListKeys = []string {}

    return &(interfaceAttributes.EntityData)
}

// Mpls_Global_InterfaceAttributes_Interface
// List of TE interfaces
type Mpls_Global_InterfaceAttributes_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the interface id list key. The type
    // is string. Refers to
    // mpls.Mpls_Global_InterfaceAttributes_Interface_Config_InterfaceId
    InterfaceId interface{}

    // Configuration parameters related to MPLS interfaces:.
    Config Mpls_Global_InterfaceAttributes_Interface_Config

    // State parameters related to TE interfaces.
    State Mpls_Global_InterfaceAttributes_Interface_State

    // Reference to an interface or subinterface.
    InterfaceRef Mpls_Global_InterfaceAttributes_Interface_InterfaceRef
}

func (self *Mpls_Global_InterfaceAttributes_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "interface-attributes"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceId, "interface-id")
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("config", types.YChild{"Config", &self.Config})
    self.EntityData.Children.Append("state", types.YChild{"State", &self.State})
    self.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &self.InterfaceRef})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", self.InterfaceId})

    self.EntityData.YListKeys = []string {"InterfaceId"}

    return &(self.EntityData)
}

// Mpls_Global_InterfaceAttributes_Interface_Config
// Configuration parameters related to MPLS interfaces:
type Mpls_Global_InterfaceAttributes_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indentifier for the MPLS interface. The type is string.
    InterfaceId interface{}

    // Enable MPLS forwarding on this interface. The type is bool. The default
    // value is false.
    MplsEnabled interface{}
}

func (config *Mpls_Global_InterfaceAttributes_Interface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", config.InterfaceId})
    config.EntityData.Leafs.Append("mpls-enabled", types.YLeaf{"MplsEnabled", config.MplsEnabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Global_InterfaceAttributes_Interface_State
// State parameters related to TE interfaces
type Mpls_Global_InterfaceAttributes_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indentifier for the MPLS interface. The type is string.
    InterfaceId interface{}

    // Enable MPLS forwarding on this interface. The type is bool. The default
    // value is false.
    MplsEnabled interface{}
}

func (state *Mpls_Global_InterfaceAttributes_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", state.InterfaceId})
    state.EntityData.Leafs.Append("mpls-enabled", types.YLeaf{"MplsEnabled", state.MplsEnabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Global_InterfaceAttributes_Interface_InterfaceRef
// Reference to an interface or subinterface
type Mpls_Global_InterfaceAttributes_Interface_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Mpls_Global_InterfaceAttributes_Interface_InterfaceRef_Config

    // Operational state for interface-ref.
    State Mpls_Global_InterfaceAttributes_Interface_InterfaceRef_State
}

func (interfaceRef *Mpls_Global_InterfaceAttributes_Interface_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "interface"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("config", types.YChild{"Config", &interfaceRef.Config})
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Mpls_Global_InterfaceAttributes_Interface_InterfaceRef_Config
// Configured reference to interface / subinterface
type Mpls_Global_InterfaceAttributes_Interface_InterfaceRef_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (config *Mpls_Global_InterfaceAttributes_Interface_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-ref"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", config.Interface})
    config.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", config.Subinterface})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Global_InterfaceAttributes_Interface_InterfaceRef_State
// Operational state for interface-ref
type Mpls_Global_InterfaceAttributes_Interface_InterfaceRef_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (state *Mpls_Global_InterfaceAttributes_Interface_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Global_ReservedLabelBlocks
// A range of labels starting with the start-label and up-to and including
// the end label that should be allocated as reserved. These labels should
// not be utilised by any dynamic label allocation on the local system unless
// the allocating protocol is explicitly configured to specify that
// allocation of labels should be out of the label block specified.
type Mpls_Global_ReservedLabelBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A range of labels starting with the start-label up to and including the end
    // label that should be allocated for use by a specific protocol. The type is
    // slice of Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock.
    ReservedLabelBlock []*Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock
}

func (reservedLabelBlocks *Mpls_Global_ReservedLabelBlocks) GetEntityData() *types.CommonEntityData {
    reservedLabelBlocks.EntityData.YFilter = reservedLabelBlocks.YFilter
    reservedLabelBlocks.EntityData.YangName = "reserved-label-blocks"
    reservedLabelBlocks.EntityData.BundleName = "openconfig"
    reservedLabelBlocks.EntityData.ParentYangName = "global"
    reservedLabelBlocks.EntityData.SegmentPath = "reserved-label-blocks"
    reservedLabelBlocks.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    reservedLabelBlocks.EntityData.NamespaceTable = openconfig.GetNamespaces()
    reservedLabelBlocks.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    reservedLabelBlocks.EntityData.Children = types.NewOrderedMap()
    reservedLabelBlocks.EntityData.Children.Append("reserved-label-block", types.YChild{"ReservedLabelBlock", nil})
    for i := range reservedLabelBlocks.ReservedLabelBlock {
        reservedLabelBlocks.EntityData.Children.Append(types.GetSegmentPath(reservedLabelBlocks.ReservedLabelBlock[i]), types.YChild{"ReservedLabelBlock", reservedLabelBlocks.ReservedLabelBlock[i]})
    }
    reservedLabelBlocks.EntityData.Leafs = types.NewOrderedMap()

    reservedLabelBlocks.EntityData.YListKeys = []string {}

    return &(reservedLabelBlocks.EntityData)
}

// Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock
// A range of labels starting with the start-label up to and including
// the end label that should be allocated for use by a specific protocol.
type Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A reference to a unique local identifier for this
    // label block. The type is string. Refers to
    // mpls.Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_Config_LocalId
    LocalId interface{}

    // Configuration parameters relating to the label block.
    Config Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_Config

    // State parameters relating to the label block.
    State Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_State
}

func (reservedLabelBlock *Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock) GetEntityData() *types.CommonEntityData {
    reservedLabelBlock.EntityData.YFilter = reservedLabelBlock.YFilter
    reservedLabelBlock.EntityData.YangName = "reserved-label-block"
    reservedLabelBlock.EntityData.BundleName = "openconfig"
    reservedLabelBlock.EntityData.ParentYangName = "reserved-label-blocks"
    reservedLabelBlock.EntityData.SegmentPath = "reserved-label-block" + types.AddKeyToken(reservedLabelBlock.LocalId, "local-id")
    reservedLabelBlock.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    reservedLabelBlock.EntityData.NamespaceTable = openconfig.GetNamespaces()
    reservedLabelBlock.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    reservedLabelBlock.EntityData.Children = types.NewOrderedMap()
    reservedLabelBlock.EntityData.Children.Append("config", types.YChild{"Config", &reservedLabelBlock.Config})
    reservedLabelBlock.EntityData.Children.Append("state", types.YChild{"State", &reservedLabelBlock.State})
    reservedLabelBlock.EntityData.Leafs = types.NewOrderedMap()
    reservedLabelBlock.EntityData.Leafs.Append("local-id", types.YLeaf{"LocalId", reservedLabelBlock.LocalId})

    reservedLabelBlock.EntityData.YListKeys = []string {"LocalId"}

    return &(reservedLabelBlock.EntityData)
}

// Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_Config
// Configuration parameters relating to the label block.
type Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A local identifier for the global label block allocation. The type is
    // string.
    LocalId interface{}

    // Lower bound of the global label block. The block is defined to include this
    // label. The type is one of the following types: int with range: 16..1048575,
    // or enumeration MplsLabel.
    LowerBound interface{}

    // Upper bound for the global label block. The block is defined to include
    // this label. The type is one of the following types: int with range:
    // 16..1048575, or enumeration MplsLabel.
    UpperBound interface{}
}

func (config *Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "reserved-label-block"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("local-id", types.YLeaf{"LocalId", config.LocalId})
    config.EntityData.Leafs.Append("lower-bound", types.YLeaf{"LowerBound", config.LowerBound})
    config.EntityData.Leafs.Append("upper-bound", types.YLeaf{"UpperBound", config.UpperBound})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_State
// State parameters relating to the label block.
type Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A local identifier for the global label block allocation. The type is
    // string.
    LocalId interface{}

    // Lower bound of the global label block. The block is defined to include this
    // label. The type is one of the following types: int with range: 16..1048575,
    // or enumeration MplsLabel.
    LowerBound interface{}

    // Upper bound for the global label block. The block is defined to include
    // this label. The type is one of the following types: int with range:
    // 16..1048575, or enumeration MplsLabel.
    UpperBound interface{}
}

func (state *Mpls_Global_ReservedLabelBlocks_ReservedLabelBlock_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "reserved-label-block"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("local-id", types.YLeaf{"LocalId", state.LocalId})
    state.EntityData.Leafs.Append("lower-bound", types.YLeaf{"LowerBound", state.LowerBound})
    state.EntityData.Leafs.Append("upper-bound", types.YLeaf{"UpperBound", state.UpperBound})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_TeGlobalAttributes
// traffic-engineering global attributes
type Mpls_TeGlobalAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shared risk link groups attributes.
    Srlgs Mpls_TeGlobalAttributes_Srlgs

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

    teGlobalAttributes.EntityData.Children = types.NewOrderedMap()
    teGlobalAttributes.EntityData.Children.Append("srlgs", types.YChild{"Srlgs", &teGlobalAttributes.Srlgs})
    teGlobalAttributes.EntityData.Children.Append("mpls-admin-groups", types.YChild{"MplsAdminGroups", &teGlobalAttributes.MplsAdminGroups})
    teGlobalAttributes.EntityData.Children.Append("te-lsp-timers", types.YChild{"TeLspTimers", &teGlobalAttributes.TeLspTimers})
    teGlobalAttributes.EntityData.Leafs = types.NewOrderedMap()

    teGlobalAttributes.EntityData.YListKeys = []string {}

    return &(teGlobalAttributes.EntityData)
}

// Mpls_TeGlobalAttributes_Srlgs
// Shared risk link groups attributes
type Mpls_TeGlobalAttributes_Srlgs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of shared risk link groups. The type is slice of
    // Mpls_TeGlobalAttributes_Srlgs_Srlg.
    Srlg []*Mpls_TeGlobalAttributes_Srlgs_Srlg
}

func (srlgs *Mpls_TeGlobalAttributes_Srlgs) GetEntityData() *types.CommonEntityData {
    srlgs.EntityData.YFilter = srlgs.YFilter
    srlgs.EntityData.YangName = "srlgs"
    srlgs.EntityData.BundleName = "openconfig"
    srlgs.EntityData.ParentYangName = "te-global-attributes"
    srlgs.EntityData.SegmentPath = "srlgs"
    srlgs.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    srlgs.EntityData.NamespaceTable = openconfig.GetNamespaces()
    srlgs.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    srlgs.EntityData.Children = types.NewOrderedMap()
    srlgs.EntityData.Children.Append("srlg", types.YChild{"Srlg", nil})
    for i := range srlgs.Srlg {
        srlgs.EntityData.Children.Append(types.GetSegmentPath(srlgs.Srlg[i]), types.YChild{"Srlg", srlgs.Srlg[i]})
    }
    srlgs.EntityData.Leafs = types.NewOrderedMap()

    srlgs.EntityData.YListKeys = []string {}

    return &(srlgs.EntityData)
}

// Mpls_TeGlobalAttributes_Srlgs_Srlg
// List of shared risk link groups
type Mpls_TeGlobalAttributes_Srlgs_Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The SRLG group identifier. The type is string.
    // Refers to mpls.Mpls_TeGlobalAttributes_Srlgs_Srlg_Config_Name
    Name interface{}

    // Configuration parameters related to the SRLG.
    Config Mpls_TeGlobalAttributes_Srlgs_Srlg_Config

    // State parameters related to the SRLG.
    State Mpls_TeGlobalAttributes_Srlgs_Srlg_State

    // SRLG members for static (not flooded) SRLGs .
    StaticSrlgMembers Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers
}

func (srlg *Mpls_TeGlobalAttributes_Srlgs_Srlg) GetEntityData() *types.CommonEntityData {
    srlg.EntityData.YFilter = srlg.YFilter
    srlg.EntityData.YangName = "srlg"
    srlg.EntityData.BundleName = "openconfig"
    srlg.EntityData.ParentYangName = "srlgs"
    srlg.EntityData.SegmentPath = "srlg" + types.AddKeyToken(srlg.Name, "name")
    srlg.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    srlg.EntityData.NamespaceTable = openconfig.GetNamespaces()
    srlg.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    srlg.EntityData.Children = types.NewOrderedMap()
    srlg.EntityData.Children.Append("config", types.YChild{"Config", &srlg.Config})
    srlg.EntityData.Children.Append("state", types.YChild{"State", &srlg.State})
    srlg.EntityData.Children.Append("static-srlg-members", types.YChild{"StaticSrlgMembers", &srlg.StaticSrlgMembers})
    srlg.EntityData.Leafs = types.NewOrderedMap()
    srlg.EntityData.Leafs.Append("name", types.YLeaf{"Name", srlg.Name})

    srlg.EntityData.YListKeys = []string {"Name"}

    return &(srlg.EntityData)
}

// Mpls_TeGlobalAttributes_Srlgs_Srlg_Config
// Configuration parameters related to the SRLG
type Mpls_TeGlobalAttributes_Srlgs_Srlg_Config struct {
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
    // type is MplsSrlgFloodingType. The default value is FLOODED_SRLG.
    FloodingType interface{}
}

func (config *Mpls_TeGlobalAttributes_Srlgs_Srlg_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "srlg"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("value", types.YLeaf{"Value", config.Value})
    config.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", config.Cost})
    config.EntityData.Leafs.Append("flooding-type", types.YLeaf{"FloodingType", config.FloodingType})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_TeGlobalAttributes_Srlgs_Srlg_State
// State parameters related to the SRLG
type Mpls_TeGlobalAttributes_Srlgs_Srlg_State struct {
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
    // type is MplsSrlgFloodingType. The default value is FLOODED_SRLG.
    FloodingType interface{}
}

func (state *Mpls_TeGlobalAttributes_Srlgs_Srlg_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "srlg"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("value", types.YLeaf{"Value", state.Value})
    state.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", state.Cost})
    state.EntityData.Leafs.Append("flooding-type", types.YLeaf{"FloodingType", state.FloodingType})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers
// SRLG members for static (not flooded) SRLGs 
type Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of SRLG members, which are expressed as IP address endpoints of links
    // contained in the SRLG. The type is slice of
    // Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList.
    MembersList []*Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList
}

func (staticSrlgMembers *Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers) GetEntityData() *types.CommonEntityData {
    staticSrlgMembers.EntityData.YFilter = staticSrlgMembers.YFilter
    staticSrlgMembers.EntityData.YangName = "static-srlg-members"
    staticSrlgMembers.EntityData.BundleName = "openconfig"
    staticSrlgMembers.EntityData.ParentYangName = "srlg"
    staticSrlgMembers.EntityData.SegmentPath = "static-srlg-members"
    staticSrlgMembers.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    staticSrlgMembers.EntityData.NamespaceTable = openconfig.GetNamespaces()
    staticSrlgMembers.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    staticSrlgMembers.EntityData.Children = types.NewOrderedMap()
    staticSrlgMembers.EntityData.Children.Append("members-list", types.YChild{"MembersList", nil})
    for i := range staticSrlgMembers.MembersList {
        staticSrlgMembers.EntityData.Children.Append(types.GetSegmentPath(staticSrlgMembers.MembersList[i]), types.YChild{"MembersList", staticSrlgMembers.MembersList[i]})
    }
    staticSrlgMembers.EntityData.Leafs = types.NewOrderedMap()

    staticSrlgMembers.EntityData.YListKeys = []string {}

    return &(staticSrlgMembers.EntityData)
}

// Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList
// List of SRLG members, which are expressed
// as IP address endpoints of links contained in the
// SRLG
type Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The from address of the link in the SRLG. The type
    // is one of the following types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    FromAddress interface{}

    // Configuration parameters relating to the SRLG members.
    Config Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList_Config

    // State parameters relating to the SRLG members.
    State Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList_State
}

func (membersList *Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList) GetEntityData() *types.CommonEntityData {
    membersList.EntityData.YFilter = membersList.YFilter
    membersList.EntityData.YangName = "members-list"
    membersList.EntityData.BundleName = "openconfig"
    membersList.EntityData.ParentYangName = "static-srlg-members"
    membersList.EntityData.SegmentPath = "members-list" + types.AddKeyToken(membersList.FromAddress, "from-address")
    membersList.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    membersList.EntityData.NamespaceTable = openconfig.GetNamespaces()
    membersList.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    membersList.EntityData.Children = types.NewOrderedMap()
    membersList.EntityData.Children.Append("config", types.YChild{"Config", &membersList.Config})
    membersList.EntityData.Children.Append("state", types.YChild{"State", &membersList.State})
    membersList.EntityData.Leafs = types.NewOrderedMap()
    membersList.EntityData.Leafs.Append("from-address", types.YLeaf{"FromAddress", membersList.FromAddress})

    membersList.EntityData.YListKeys = []string {"FromAddress"}

    return &(membersList.EntityData)
}

// Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList_Config
// Configuration parameters relating to the
// SRLG members
type Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the a-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    FromAddress interface{}

    // IP address of the z-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    ToAddress interface{}
}

func (config *Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "members-list"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("from-address", types.YLeaf{"FromAddress", config.FromAddress})
    config.EntityData.Leafs.Append("to-address", types.YLeaf{"ToAddress", config.ToAddress})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList_State
// State parameters relating to the SRLG
// members
type Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of the a-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    FromAddress interface{}

    // IP address of the z-side of the SRLG link. The type is one of the following
    // types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    ToAddress interface{}
}

func (state *Mpls_TeGlobalAttributes_Srlgs_Srlg_StaticSrlgMembers_MembersList_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "members-list"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("from-address", types.YLeaf{"FromAddress", state.FromAddress})
    state.EntityData.Leafs.Append("to-address", types.YLeaf{"ToAddress", state.ToAddress})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_TeGlobalAttributes_MplsAdminGroups
// Top-level container for admin-groups configuration
// and state
type Mpls_TeGlobalAttributes_MplsAdminGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // configuration of value to name mapping for mpls affinities/admin-groups.
    // The type is slice of Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup.
    AdminGroup []*Mpls_TeGlobalAttributes_MplsAdminGroups_AdminGroup
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

    mplsAdminGroups.EntityData.Children = types.NewOrderedMap()
    mplsAdminGroups.EntityData.Children.Append("admin-group", types.YChild{"AdminGroup", nil})
    for i := range mplsAdminGroups.AdminGroup {
        mplsAdminGroups.EntityData.Children.Append(types.GetSegmentPath(mplsAdminGroups.AdminGroup[i]), types.YChild{"AdminGroup", mplsAdminGroups.AdminGroup[i]})
    }
    mplsAdminGroups.EntityData.Leafs = types.NewOrderedMap()

    mplsAdminGroups.EntityData.YListKeys = []string {}

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
    adminGroup.EntityData.SegmentPath = "admin-group" + types.AddKeyToken(adminGroup.AdminGroupName, "admin-group-name")
    adminGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adminGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adminGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adminGroup.EntityData.Children = types.NewOrderedMap()
    adminGroup.EntityData.Children.Append("config", types.YChild{"Config", &adminGroup.Config})
    adminGroup.EntityData.Children.Append("state", types.YChild{"State", &adminGroup.State})
    adminGroup.EntityData.Leafs = types.NewOrderedMap()
    adminGroup.EntityData.Leafs.Append("admin-group-name", types.YLeaf{"AdminGroupName", adminGroup.AdminGroupName})

    adminGroup.EntityData.YListKeys = []string {"AdminGroupName"}

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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("admin-group-name", types.YLeaf{"AdminGroupName", config.AdminGroupName})
    config.EntityData.Leafs.Append("bit-position", types.YLeaf{"BitPosition", config.BitPosition})

    config.EntityData.YListKeys = []string {}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("admin-group-name", types.YLeaf{"AdminGroupName", state.AdminGroupName})
    state.EntityData.Leafs.Append("bit-position", types.YLeaf{"BitPosition", state.BitPosition})

    state.EntityData.YListKeys = []string {}

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

    teLspTimers.EntityData.Children = types.NewOrderedMap()
    teLspTimers.EntityData.Children.Append("config", types.YChild{"Config", &teLspTimers.Config})
    teLspTimers.EntityData.Children.Append("state", types.YChild{"State", &teLspTimers.State})
    teLspTimers.EntityData.Leafs = types.NewOrderedMap()

    teLspTimers.EntityData.YListKeys = []string {}

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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("install-delay", types.YLeaf{"InstallDelay", config.InstallDelay})
    config.EntityData.Leafs.Append("cleanup-delay", types.YLeaf{"CleanupDelay", config.CleanupDelay})
    config.EntityData.Leafs.Append("reoptimize-timer", types.YLeaf{"ReoptimizeTimer", config.ReoptimizeTimer})

    config.EntityData.YListKeys = []string {}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("install-delay", types.YLeaf{"InstallDelay", state.InstallDelay})
    state.EntityData.Leafs.Append("cleanup-delay", types.YLeaf{"CleanupDelay", state.CleanupDelay})
    state.EntityData.Leafs.Append("reoptimize-timer", types.YLeaf{"ReoptimizeTimer", state.ReoptimizeTimer})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_TeInterfaceAttributes
// traffic engineering attributes specific
// for interfaces
type Mpls_TeInterfaceAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of TE interfaces. The type is slice of
    // Mpls_TeInterfaceAttributes_Interface.
    Interface []*Mpls_TeInterfaceAttributes_Interface
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

    teInterfaceAttributes.EntityData.Children = types.NewOrderedMap()
    teInterfaceAttributes.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range teInterfaceAttributes.Interface {
        teInterfaceAttributes.EntityData.Children.Append(types.GetSegmentPath(teInterfaceAttributes.Interface[i]), types.YChild{"Interface", teInterfaceAttributes.Interface[i]})
    }
    teInterfaceAttributes.EntityData.Leafs = types.NewOrderedMap()

    teInterfaceAttributes.EntityData.YListKeys = []string {}

    return &(teInterfaceAttributes.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface
// List of TE interfaces
type Mpls_TeInterfaceAttributes_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the interface id list key. The type
    // is string. Refers to
    // mpls.Mpls_TeInterfaceAttributes_Interface_Config_InterfaceId
    InterfaceId interface{}

    // Configuration parameters related to TE interfaces:.
    Config Mpls_TeInterfaceAttributes_Interface_Config

    // State parameters related to TE interfaces.
    State Mpls_TeInterfaceAttributes_Interface_State

    // Reference to an interface or subinterface.
    InterfaceRef Mpls_TeInterfaceAttributes_Interface_InterfaceRef

    // Interface bandwidth change percentages that trigger update events into the
    // IGP traffic engineering database (TED).
    IgpFloodingBandwidth Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth
}

func (self *Mpls_TeInterfaceAttributes_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "te-interface-attributes"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceId, "interface-id")
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("config", types.YChild{"Config", &self.Config})
    self.EntityData.Children.Append("state", types.YChild{"State", &self.State})
    self.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &self.InterfaceRef})
    self.EntityData.Children.Append("igp-flooding-bandwidth", types.YChild{"IgpFloodingBandwidth", &self.IgpFloodingBandwidth})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", self.InterfaceId})

    self.EntityData.YListKeys = []string {"InterfaceId"}

    return &(self.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_Config
// Configuration parameters related to TE interfaces:
type Mpls_TeInterfaceAttributes_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Id of the interface. The type is string.
    InterfaceId interface{}

    // TE specific metric for the link. The type is interface{} with range:
    // 0..4294967295.
    TeMetric interface{}

    // list of references to named shared risk link groups that the interface
    // belongs to. The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_Srlgs_Srlg_Name
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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", config.InterfaceId})
    config.EntityData.Leafs.Append("te-metric", types.YLeaf{"TeMetric", config.TeMetric})
    config.EntityData.Leafs.Append("srlg-membership", types.YLeaf{"SrlgMembership", config.SrlgMembership})
    config.EntityData.Leafs.Append("admin-group", types.YLeaf{"AdminGroup", config.AdminGroup})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_State
// State parameters related to TE interfaces
type Mpls_TeInterfaceAttributes_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Id of the interface. The type is string.
    InterfaceId interface{}

    // TE specific metric for the link. The type is interface{} with range:
    // 0..4294967295.
    TeMetric interface{}

    // list of references to named shared risk link groups that the interface
    // belongs to. The type is slice of string. Refers to
    // mpls.Mpls_TeGlobalAttributes_Srlgs_Srlg_Name
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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", state.InterfaceId})
    state.EntityData.Leafs.Append("te-metric", types.YLeaf{"TeMetric", state.TeMetric})
    state.EntityData.Leafs.Append("srlg-membership", types.YLeaf{"SrlgMembership", state.SrlgMembership})
    state.EntityData.Leafs.Append("admin-group", types.YLeaf{"AdminGroup", state.AdminGroup})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_InterfaceRef
// Reference to an interface or subinterface
type Mpls_TeInterfaceAttributes_Interface_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Mpls_TeInterfaceAttributes_Interface_InterfaceRef_Config

    // Operational state for interface-ref.
    State Mpls_TeInterfaceAttributes_Interface_InterfaceRef_State
}

func (interfaceRef *Mpls_TeInterfaceAttributes_Interface_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "interface"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("config", types.YChild{"Config", &interfaceRef.Config})
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_InterfaceRef_Config
// Configured reference to interface / subinterface
type Mpls_TeInterfaceAttributes_Interface_InterfaceRef_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (config *Mpls_TeInterfaceAttributes_Interface_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-ref"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", config.Interface})
    config.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", config.Subinterface})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_InterfaceRef_State
// Operational state for interface-ref
type Mpls_TeInterfaceAttributes_Interface_InterfaceRef_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (state *Mpls_TeInterfaceAttributes_Interface_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

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

    igpFloodingBandwidth.EntityData.Children = types.NewOrderedMap()
    igpFloodingBandwidth.EntityData.Children.Append("config", types.YChild{"Config", &igpFloodingBandwidth.Config})
    igpFloodingBandwidth.EntityData.Children.Append("state", types.YChild{"State", &igpFloodingBandwidth.State})
    igpFloodingBandwidth.EntityData.Leafs = types.NewOrderedMap()

    igpFloodingBandwidth.EntityData.YListKeys = []string {}

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
    // occurs on the interface. Where THRESHOLD_CROSSED is specified, the local
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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("threshold-type", types.YLeaf{"ThresholdType", config.ThresholdType})
    config.EntityData.Leafs.Append("delta-percentage", types.YLeaf{"DeltaPercentage", config.DeltaPercentage})
    config.EntityData.Leafs.Append("threshold-specification", types.YLeaf{"ThresholdSpecification", config.ThresholdSpecification})
    config.EntityData.Leafs.Append("up-thresholds", types.YLeaf{"UpThresholds", config.UpThresholds})
    config.EntityData.Leafs.Append("down-thresholds", types.YLeaf{"DownThresholds", config.DownThresholds})
    config.EntityData.Leafs.Append("up-down-thresholds", types.YLeaf{"UpDownThresholds", config.UpDownThresholds})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification represents and decreasing values will be separately specified
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification string

const (
    // MIRRORED_UP_DOWN indicates that a single set of
    // threshold values should be used for both increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification_MIRRORED_UP_DOWN Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification = "MIRRORED_UP_DOWN"

    // SEPARATE_UP_DOWN indicates that a separate
    // threshold values should be used for the increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification_SEPARATE_UP_DOWN Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdSpecification = "SEPARATE_UP_DOWN"
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
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdType_THRESHOLD_CROSSED Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_Config_ThresholdType = "THRESHOLD_CROSSED"
)

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State
// State parameters for TED update threshold 
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type of threshold that should be used to specify the values at which
    // bandwidth is flooded. DELTA indicates that the local system should flood
    // IGP updates when a change in reserved bandwidth >= the specified delta
    // occurs on the interface. Where THRESHOLD_CROSSED is specified, the local
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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("threshold-type", types.YLeaf{"ThresholdType", state.ThresholdType})
    state.EntityData.Leafs.Append("delta-percentage", types.YLeaf{"DeltaPercentage", state.DeltaPercentage})
    state.EntityData.Leafs.Append("threshold-specification", types.YLeaf{"ThresholdSpecification", state.ThresholdSpecification})
    state.EntityData.Leafs.Append("up-thresholds", types.YLeaf{"UpThresholds", state.UpThresholds})
    state.EntityData.Leafs.Append("down-thresholds", types.YLeaf{"DownThresholds", state.DownThresholds})
    state.EntityData.Leafs.Append("up-down-thresholds", types.YLeaf{"UpDownThresholds", state.UpDownThresholds})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification represents and decreasing values will be separately specified
type Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification string

const (
    // MIRRORED_UP_DOWN indicates that a single set of
    // threshold values should be used for both increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification_MIRRORED_UP_DOWN Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification = "MIRRORED_UP_DOWN"

    // SEPARATE_UP_DOWN indicates that a separate
    // threshold values should be used for the increasing
    // and decreasing bandwidth when determining whether
    // to trigger updated bandwidth values to be flooded
    // in the IGP TE extensions.
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification_SEPARATE_UP_DOWN Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdSpecification = "SEPARATE_UP_DOWN"
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
    Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdType_THRESHOLD_CROSSED Mpls_TeInterfaceAttributes_Interface_IgpFloodingBandwidth_State_ThresholdType = "THRESHOLD_CROSSED"
)

// Mpls_SignalingProtocols
// top-level signaling protocol configuration
type Mpls_SignalingProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP-TE global signaling protocol configuration.
    RsvpTe Mpls_SignalingProtocols_RsvpTe

    // LDP global signaling configuration.
    Ldp Mpls_SignalingProtocols_Ldp

    // MPLS-specific Segment Routing configuration and operational state
    // parameters.
    SegmentRouting Mpls_SignalingProtocols_SegmentRouting
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

    signalingProtocols.EntityData.Children = types.NewOrderedMap()
    signalingProtocols.EntityData.Children.Append("rsvp-te", types.YChild{"RsvpTe", &signalingProtocols.RsvpTe})
    signalingProtocols.EntityData.Children.Append("ldp", types.YChild{"Ldp", &signalingProtocols.Ldp})
    signalingProtocols.EntityData.Children.Append("segment-routing", types.YChild{"SegmentRouting", &signalingProtocols.SegmentRouting})
    signalingProtocols.EntityData.Leafs = types.NewOrderedMap()

    signalingProtocols.EntityData.YListKeys = []string {}

    return &(signalingProtocols.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe
// RSVP-TE global signaling protocol configuration
type Mpls_SignalingProtocols_RsvpTe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for sessions.
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

    rsvpTe.EntityData.Children = types.NewOrderedMap()
    rsvpTe.EntityData.Children.Append("sessions", types.YChild{"Sessions", &rsvpTe.Sessions})
    rsvpTe.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &rsvpTe.Neighbors})
    rsvpTe.EntityData.Children.Append("global", types.YChild{"Global", &rsvpTe.Global})
    rsvpTe.EntityData.Children.Append("interface-attributes", types.YChild{"InterfaceAttributes", &rsvpTe.InterfaceAttributes})
    rsvpTe.EntityData.Leafs = types.NewOrderedMap()

    rsvpTe.EntityData.YListKeys = []string {}

    return &(rsvpTe.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions
// Enclosing container for sessions
type Mpls_SignalingProtocols_RsvpTe_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of RSVP sessions. The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_Sessions_Session.
    Session []*Mpls_SignalingProtocols_RsvpTe_Sessions_Session
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

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_Session
// List of RSVP sessions
type Mpls_SignalingProtocols_RsvpTe_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the local index for the RSVP session.
    // The type is string with range: 0..18446744073709551615. Refers to
    // mpls.Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_LocalIndex
    LocalIndex interface{}

    // Enclosing container for MPLS RRO objects associated with the traffic
    // engineered tunnel.
    RecordRouteObjects Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects

    // Operational state parameters relating to the RSVP session.
    State Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State
}

func (session *Mpls_SignalingProtocols_RsvpTe_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "openconfig"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.LocalIndex, "local-index")
    session.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    session.EntityData.NamespaceTable = openconfig.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("record-route-objects", types.YChild{"RecordRouteObjects", &session.RecordRouteObjects})
    session.EntityData.Children.Append("state", types.YChild{"State", &session.State})
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("local-index", types.YLeaf{"LocalIndex", session.LocalIndex})

    session.EntityData.YListKeys = []string {"LocalIndex"}

    return &(session.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects
// Enclosing container for MPLS RRO objects associated with the
// traffic engineered tunnel.
type Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Read-only list of record route objects associated with the traffic
    // engineered tunnel. Each entry in the list may contain a hop IP address,
    // MPLS label allocated at the hop, and the flags associated with the entry.
    // The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject.
    RecordRouteObject []*Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject
}

func (recordRouteObjects *Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects) GetEntityData() *types.CommonEntityData {
    recordRouteObjects.EntityData.YFilter = recordRouteObjects.YFilter
    recordRouteObjects.EntityData.YangName = "record-route-objects"
    recordRouteObjects.EntityData.BundleName = "openconfig"
    recordRouteObjects.EntityData.ParentYangName = "session"
    recordRouteObjects.EntityData.SegmentPath = "record-route-objects"
    recordRouteObjects.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    recordRouteObjects.EntityData.NamespaceTable = openconfig.GetNamespaces()
    recordRouteObjects.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    recordRouteObjects.EntityData.Children = types.NewOrderedMap()
    recordRouteObjects.EntityData.Children.Append("record-route-object", types.YChild{"RecordRouteObject", nil})
    for i := range recordRouteObjects.RecordRouteObject {
        recordRouteObjects.EntityData.Children.Append(types.GetSegmentPath(recordRouteObjects.RecordRouteObject[i]), types.YChild{"RecordRouteObject", recordRouteObjects.RecordRouteObject[i]})
    }
    recordRouteObjects.EntityData.Leafs = types.NewOrderedMap()

    recordRouteObjects.EntityData.YListKeys = []string {}

    return &(recordRouteObjects.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject
// Read-only list of record route objects associated with the
// traffic engineered tunnel. Each entry in the list
// may contain a hop IP address, MPLS label allocated
// at the hop, and the flags associated with the entry.
type Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the index of the record route object.
    // The index is used to indicate the ordering of hops in the path. The type is
    // string with range: 0..255. Refers to
    // mpls.Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject_State_Index
    Index interface{}

    // Information related to RRO objects. The hop, label, and optional flags are
    // present for each entry in the list.
    State Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject_State
}

func (recordRouteObject *Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject) GetEntityData() *types.CommonEntityData {
    recordRouteObject.EntityData.YFilter = recordRouteObject.YFilter
    recordRouteObject.EntityData.YangName = "record-route-object"
    recordRouteObject.EntityData.BundleName = "openconfig"
    recordRouteObject.EntityData.ParentYangName = "record-route-objects"
    recordRouteObject.EntityData.SegmentPath = "record-route-object" + types.AddKeyToken(recordRouteObject.Index, "index")
    recordRouteObject.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    recordRouteObject.EntityData.NamespaceTable = openconfig.GetNamespaces()
    recordRouteObject.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    recordRouteObject.EntityData.Children = types.NewOrderedMap()
    recordRouteObject.EntityData.Children.Append("state", types.YChild{"State", &recordRouteObject.State})
    recordRouteObject.EntityData.Leafs = types.NewOrderedMap()
    recordRouteObject.EntityData.Leafs.Append("index", types.YLeaf{"Index", recordRouteObject.Index})

    recordRouteObject.EntityData.YListKeys = []string {"Index"}

    return &(recordRouteObject.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject_State
// Information related to RRO objects. The hop, label, and
// optional flags are present for each entry in the list.
type Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index of object in the list. Used for ordering. The type is interface{}
    // with range: 0..255.
    Index interface{}

    // IP router hop for RRO entry. The type is one of the following types: string
    // with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    Address interface{}

    // Label reported for RRO hop. The type is one of the following types: int
    // with range: 16..1048575, or enumeration MplsLabel.
    ReportedLabel interface{}

    // Subobject flags for MPLS label. The type is interface{} with range: 0..255.
    ReportedFlags interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_Session_RecordRouteObjects_RecordRouteObject_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "record-route-object"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})
    state.EntityData.Leafs.Append("address", types.YLeaf{"Address", state.Address})
    state.EntityData.Leafs.Append("reported-label", types.YLeaf{"ReportedLabel", state.ReportedLabel})
    state.EntityData.Leafs.Append("reported-flags", types.YLeaf{"ReportedFlags", state.ReportedFlags})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State
// Operational state parameters relating to the
// RSVP session
type Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The index used to identify the RSVP session on the local network element.
    // This index is generated by the device and is unique only to the local
    // network element. The type is interface{} with range:
    // 0..18446744073709551615.
    LocalIndex interface{}

    // Origin address of RSVP session. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    SourceAddress interface{}

    // Destination address of RSVP session. The type is one of the following
    // types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    DestinationAddress interface{}

    // The tunnel ID is an identifier used in the RSVP session, which remains
    // constant over the life of the tunnel. The type is interface{} with range:
    // 0..65535.
    TunnelId interface{}

    // The LSP ID distinguishes between two LSPs originated from the same headend,
    // and is commonly used to distinguish RSVP sessions during make before break
    // operations. The type is interface{} with range: 0..65535.
    LspId interface{}

    // The signaled name of this RSVP session. The type is string.
    SessionName interface{}

    // Enumeration of RSVP session states. The type is Status.
    Status interface{}

    // The type/role of the RSVP session, signifing the session's role on the
    // current device, such as a transit session vs. an ingress session. The type
    // is one of the following: INGRESSEGRESSTRANSIT.
    Type interface{}

    // The type of protection requested for the RSVP session. The type is one of
    // the following:
    // LINKPROTECTIONREQUIREDLINKNODEPROTECTIONREQUESTEDUNPROTECTED.
    ProtectionRequested interface{}

    // Incoming MPLS label associated with this RSVP session. The type is one of
    // the following types: int with range: 16..1048575, or enumeration MplsLabel.
    LabelIn interface{}

    // Outgoing MPLS label associated with this RSVP session. The type is one of
    // the following types: int with range: 16..1048575, or enumeration MplsLabel.
    LabelOut interface{}

    // Operational state statistics relating to the SENDER_TSPEC received for the
    // RSVP session.
    SenderTspec Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_SenderTspec
}

func (state *Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "session"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("sender-tspec", types.YChild{"SenderTspec", &state.SenderTspec})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("local-index", types.YLeaf{"LocalIndex", state.LocalIndex})
    state.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", state.SourceAddress})
    state.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", state.DestinationAddress})
    state.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", state.TunnelId})
    state.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", state.LspId})
    state.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", state.SessionName})
    state.EntityData.Leafs.Append("status", types.YLeaf{"Status", state.Status})
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})
    state.EntityData.Leafs.Append("protection-requested", types.YLeaf{"ProtectionRequested", state.ProtectionRequested})
    state.EntityData.Leafs.Append("label-in", types.YLeaf{"LabelIn", state.LabelIn})
    state.EntityData.Leafs.Append("label-out", types.YLeaf{"LabelOut", state.LabelOut})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_SenderTspec
// Operational state statistics relating to the SENDER_TSPEC
// received for the RSVP session
type Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_SenderTspec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The rate at which the head-end device generates traffic, expressed in bytes
    // per second. The type is string with length: 32. Units are Bps.
    Rate interface{}

    // The size of the token bucket that is used to determine the rate at which
    // the head-end device generates traffic, expressed in bytes per second. The
    // type is string with length: 32. Units are bytes per second.
    Size interface{}

    // The maximum traffic generation rate that the head-end device sends traffic
    // at. The type is one of the following types: string with length: 32 Units
    // are bytes per second., or enumeration
    // NetworkInstances.NetworkInstance.Mpls.SignalingProtocols.RsvpTe.Sessions.Session.State.SenderTspec.PeakDataRate
    // Units are bytes per second..
    PeakDataRate interface{}
}

func (senderTspec *Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_SenderTspec) GetEntityData() *types.CommonEntityData {
    senderTspec.EntityData.YFilter = senderTspec.YFilter
    senderTspec.EntityData.YangName = "sender-tspec"
    senderTspec.EntityData.BundleName = "openconfig"
    senderTspec.EntityData.ParentYangName = "state"
    senderTspec.EntityData.SegmentPath = "sender-tspec"
    senderTspec.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    senderTspec.EntityData.NamespaceTable = openconfig.GetNamespaces()
    senderTspec.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    senderTspec.EntityData.Children = types.NewOrderedMap()
    senderTspec.EntityData.Leafs = types.NewOrderedMap()
    senderTspec.EntityData.Leafs.Append("rate", types.YLeaf{"Rate", senderTspec.Rate})
    senderTspec.EntityData.Leafs.Append("size", types.YLeaf{"Size", senderTspec.Size})
    senderTspec.EntityData.Leafs.Append("peak-data-rate", types.YLeaf{"PeakDataRate", senderTspec.PeakDataRate})

    senderTspec.EntityData.YListKeys = []string {}

    return &(senderTspec.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_SenderTspec_PeakDataRate represents device sends traffic at.
type Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_SenderTspec_PeakDataRate string

const (
    // The head-end device has no maximum data rate.
    Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_SenderTspec_PeakDataRate_INFINITY Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_SenderTspec_PeakDataRate = "INFINITY"
)

// Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_Status represents Enumeration of RSVP session states
type Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_Status string

const (
    // RSVP session is up
    Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_Status_UP Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_Status = "UP"

    // RSVP session is down
    Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_Status_DOWN Mpls_SignalingProtocols_RsvpTe_Sessions_Session_State_Status = "DOWN"
)

// Mpls_SignalingProtocols_RsvpTe_Neighbors
// Configuration and state for RSVP neighbors connecting
// to the device
type Mpls_SignalingProtocols_RsvpTe_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of RSVP neighbors of the local system. The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor.
    Neighbor []*Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor
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

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor
// List of RSVP neighbors of the local system
type Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the address of the RSVP neighbor. The
    // type is one of the following types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    Address interface{}

    // Operational state parameters relating to the RSVP neighbor.
    State Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State
}

func (neighbor *Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.Address, "address")
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("state", types.YChild{"State", &neighbor.State})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("address", types.YLeaf{"Address", neighbor.Address})

    neighbor.EntityData.YListKeys = []string {"Address"}

    return &(neighbor.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State
// Operational state parameters relating to the
// RSVP neighbor
type Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address of RSVP neighbor. The type is one of the following types: string
    // with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    Address interface{}

    // Interface where RSVP neighbor was detected. The type is string.
    DetectedInterface interface{}

    // Enumuration of possible RSVP neighbor states. The type is NeighborStatus.
    NeighborStatus interface{}

    // Suppport of neighbor for RSVP refresh reduction. The type is bool.
    RefreshReduction interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "neighbor"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("address", types.YLeaf{"Address", state.Address})
    state.EntityData.Leafs.Append("detected-interface", types.YLeaf{"DetectedInterface", state.DetectedInterface})
    state.EntityData.Leafs.Append("neighbor-status", types.YLeaf{"NeighborStatus", state.NeighborStatus})
    state.EntityData.Leafs.Append("refresh-reduction", types.YLeaf{"RefreshReduction", state.RefreshReduction})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State_NeighborStatus represents Enumuration of possible RSVP neighbor states
type Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State_NeighborStatus string

const (
    // RSVP hello messages are detected from the neighbor
    Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State_NeighborStatus_UP Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State_NeighborStatus = "UP"

    // RSVP neighbor not detected as up, due to a
    // communication failure or IGP notification
    // the neighbor is unavailable
    Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State_NeighborStatus_DOWN Mpls_SignalingProtocols_RsvpTe_Neighbors_Neighbor_State_NeighborStatus = "DOWN"
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

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &global.GracefulRestart})
    global.EntityData.Children.Append("soft-preemption", types.YChild{"SoftPreemption", &global.SoftPreemption})
    global.EntityData.Children.Append("hellos", types.YChild{"Hellos", &global.Hellos})
    global.EntityData.Children.Append("state", types.YChild{"State", &global.State})
    global.EntityData.Leafs = types.NewOrderedMap()

    global.EntityData.YListKeys = []string {}

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

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("config", types.YChild{"Config", &gracefulRestart.Config})
    gracefulRestart.EntityData.Children.Append("state", types.YChild{"State", &gracefulRestart.State})
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()

    gracefulRestart.EntityData.YListKeys = []string {}

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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", config.Enable})
    config.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", config.RestartTime})
    config.EntityData.Leafs.Append("recovery-time", types.YLeaf{"RecoveryTime", config.RecoveryTime})

    config.EntityData.YListKeys = []string {}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", state.Enable})
    state.EntityData.Leafs.Append("restart-time", types.YLeaf{"RestartTime", state.RestartTime})
    state.EntityData.Leafs.Append("recovery-time", types.YLeaf{"RecoveryTime", state.RecoveryTime})

    state.EntityData.YListKeys = []string {}

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

    softPreemption.EntityData.Children = types.NewOrderedMap()
    softPreemption.EntityData.Children.Append("config", types.YChild{"Config", &softPreemption.Config})
    softPreemption.EntityData.Children.Append("state", types.YChild{"State", &softPreemption.State})
    softPreemption.EntityData.Leafs = types.NewOrderedMap()

    softPreemption.EntityData.YListKeys = []string {}

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

    // Timeout value for soft preemption to revert to hard preemption. The default
    // timeout for soft-preemption is 30 seconds - after which the local system
    // reverts to hard pre-emption. The type is interface{} with range: 0..65535.
    // The default value is 30.
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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", config.Enable})
    config.EntityData.Leafs.Append("soft-preemption-timeout", types.YLeaf{"SoftPreemptionTimeout", config.SoftPreemptionTimeout})

    config.EntityData.YListKeys = []string {}

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

    // Timeout value for soft preemption to revert to hard preemption. The default
    // timeout for soft-preemption is 30 seconds - after which the local system
    // reverts to hard pre-emption. The type is interface{} with range: 0..65535.
    // The default value is 30.
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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", state.Enable})
    state.EntityData.Leafs.Append("soft-preemption-timeout", types.YLeaf{"SoftPreemptionTimeout", state.SoftPreemptionTimeout})

    state.EntityData.YListKeys = []string {}

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

    hellos.EntityData.Children = types.NewOrderedMap()
    hellos.EntityData.Children.Append("config", types.YChild{"Config", &hellos.Config})
    hellos.EntityData.Children.Append("state", types.YChild{"State", &hellos.State})
    hellos.EntityData.Leafs = types.NewOrderedMap()

    hellos.EntityData.YListKeys = []string {}

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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", config.HelloInterval})
    config.EntityData.Leafs.Append("refresh-reduction", types.YLeaf{"RefreshReduction", config.RefreshReduction})

    config.EntityData.YListKeys = []string {}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", state.HelloInterval})
    state.EntityData.Leafs.Append("refresh-reduction", types.YLeaf{"RefreshReduction", state.RefreshReduction})

    state.EntityData.YListKeys = []string {}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("counters", types.YChild{"Counters", &state.Counters})
    state.EntityData.Leafs = types.NewOrderedMap()

    state.EntityData.YListKeys = []string {}

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

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Leafs = types.NewOrderedMap()
    counters.EntityData.Leafs.Append("path-timeouts", types.YLeaf{"PathTimeouts", counters.PathTimeouts})
    counters.EntityData.Leafs.Append("reservation-timeouts", types.YLeaf{"ReservationTimeouts", counters.ReservationTimeouts})
    counters.EntityData.Leafs.Append("rate-limited-messages", types.YLeaf{"RateLimitedMessages", counters.RateLimitedMessages})
    counters.EntityData.Leafs.Append("in-path-messages", types.YLeaf{"InPathMessages", counters.InPathMessages})
    counters.EntityData.Leafs.Append("in-path-error-messages", types.YLeaf{"InPathErrorMessages", counters.InPathErrorMessages})
    counters.EntityData.Leafs.Append("in-path-tear-messages", types.YLeaf{"InPathTearMessages", counters.InPathTearMessages})
    counters.EntityData.Leafs.Append("in-reservation-messages", types.YLeaf{"InReservationMessages", counters.InReservationMessages})
    counters.EntityData.Leafs.Append("in-reservation-error-messages", types.YLeaf{"InReservationErrorMessages", counters.InReservationErrorMessages})
    counters.EntityData.Leafs.Append("in-reservation-tear-messages", types.YLeaf{"InReservationTearMessages", counters.InReservationTearMessages})
    counters.EntityData.Leafs.Append("in-hello-messages", types.YLeaf{"InHelloMessages", counters.InHelloMessages})
    counters.EntityData.Leafs.Append("in-srefresh-messages", types.YLeaf{"InSrefreshMessages", counters.InSrefreshMessages})
    counters.EntityData.Leafs.Append("in-ack-messages", types.YLeaf{"InAckMessages", counters.InAckMessages})
    counters.EntityData.Leafs.Append("out-path-messages", types.YLeaf{"OutPathMessages", counters.OutPathMessages})
    counters.EntityData.Leafs.Append("out-path-error-messages", types.YLeaf{"OutPathErrorMessages", counters.OutPathErrorMessages})
    counters.EntityData.Leafs.Append("out-path-tear-messages", types.YLeaf{"OutPathTearMessages", counters.OutPathTearMessages})
    counters.EntityData.Leafs.Append("out-reservation-messages", types.YLeaf{"OutReservationMessages", counters.OutReservationMessages})
    counters.EntityData.Leafs.Append("out-reservation-error-messages", types.YLeaf{"OutReservationErrorMessages", counters.OutReservationErrorMessages})
    counters.EntityData.Leafs.Append("out-reservation-tear-messages", types.YLeaf{"OutReservationTearMessages", counters.OutReservationTearMessages})
    counters.EntityData.Leafs.Append("out-hello-messages", types.YLeaf{"OutHelloMessages", counters.OutHelloMessages})
    counters.EntityData.Leafs.Append("out-srefresh-messages", types.YLeaf{"OutSrefreshMessages", counters.OutSrefreshMessages})
    counters.EntityData.Leafs.Append("out-ack-messages", types.YLeaf{"OutAckMessages", counters.OutAckMessages})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes
// Attributes relating to RSVP-TE enabled interfaces
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of per-interface RSVP configurations. The type is slice of
    // Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface.
    Interface []*Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface
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

    interfaceAttributes.EntityData.Children = types.NewOrderedMap()
    interfaceAttributes.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaceAttributes.Interface {
        interfaceAttributes.EntityData.Children.Append(types.GetSegmentPath(interfaceAttributes.Interface[i]), types.YChild{"Interface", interfaceAttributes.Interface[i]})
    }
    interfaceAttributes.EntityData.Leafs = types.NewOrderedMap()

    interfaceAttributes.EntityData.YListKeys = []string {}

    return &(interfaceAttributes.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface
// list of per-interface RSVP configurations
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. reference to the interface-id data. The type is
    // string. Refers to
    // mpls.Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config_InterfaceId
    InterfaceId interface{}

    // Configuration of per-interface RSVP parameters.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config

    // Per-interface RSVP protocol and state information.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State

    // Reference to an interface or subinterface.
    InterfaceRef Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef

    // Enclosing container for bandwidth reservation.
    BandwidthReservations Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations

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
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceId, "interface-id")
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("config", types.YChild{"Config", &self.Config})
    self.EntityData.Children.Append("state", types.YChild{"State", &self.State})
    self.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &self.InterfaceRef})
    self.EntityData.Children.Append("bandwidth-reservations", types.YChild{"BandwidthReservations", &self.BandwidthReservations})
    self.EntityData.Children.Append("hellos", types.YChild{"Hellos", &self.Hellos})
    self.EntityData.Children.Append("authentication", types.YChild{"Authentication", &self.Authentication})
    self.EntityData.Children.Append("subscription", types.YChild{"Subscription", &self.Subscription})
    self.EntityData.Children.Append("protection", types.YChild{"Protection", &self.Protection})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", self.InterfaceId})

    self.EntityData.YListKeys = []string {"InterfaceId"}

    return &(self.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config
// Configuration of per-interface RSVP parameters
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifier for the interface. The type is string.
    InterfaceId interface{}
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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", config.InterfaceId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State
// Per-interface RSVP protocol and state information
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identifier for the interface. The type is string.
    InterfaceId interface{}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("counters", types.YChild{"Counters", &state.Counters})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", state.InterfaceId})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
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

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Leafs = types.NewOrderedMap()
    counters.EntityData.Leafs.Append("in-path-messages", types.YLeaf{"InPathMessages", counters.InPathMessages})
    counters.EntityData.Leafs.Append("in-path-error-messages", types.YLeaf{"InPathErrorMessages", counters.InPathErrorMessages})
    counters.EntityData.Leafs.Append("in-path-tear-messages", types.YLeaf{"InPathTearMessages", counters.InPathTearMessages})
    counters.EntityData.Leafs.Append("in-reservation-messages", types.YLeaf{"InReservationMessages", counters.InReservationMessages})
    counters.EntityData.Leafs.Append("in-reservation-error-messages", types.YLeaf{"InReservationErrorMessages", counters.InReservationErrorMessages})
    counters.EntityData.Leafs.Append("in-reservation-tear-messages", types.YLeaf{"InReservationTearMessages", counters.InReservationTearMessages})
    counters.EntityData.Leafs.Append("in-hello-messages", types.YLeaf{"InHelloMessages", counters.InHelloMessages})
    counters.EntityData.Leafs.Append("in-srefresh-messages", types.YLeaf{"InSrefreshMessages", counters.InSrefreshMessages})
    counters.EntityData.Leafs.Append("in-ack-messages", types.YLeaf{"InAckMessages", counters.InAckMessages})
    counters.EntityData.Leafs.Append("out-path-messages", types.YLeaf{"OutPathMessages", counters.OutPathMessages})
    counters.EntityData.Leafs.Append("out-path-error-messages", types.YLeaf{"OutPathErrorMessages", counters.OutPathErrorMessages})
    counters.EntityData.Leafs.Append("out-path-tear-messages", types.YLeaf{"OutPathTearMessages", counters.OutPathTearMessages})
    counters.EntityData.Leafs.Append("out-reservation-messages", types.YLeaf{"OutReservationMessages", counters.OutReservationMessages})
    counters.EntityData.Leafs.Append("out-reservation-error-messages", types.YLeaf{"OutReservationErrorMessages", counters.OutReservationErrorMessages})
    counters.EntityData.Leafs.Append("out-reservation-tear-messages", types.YLeaf{"OutReservationTearMessages", counters.OutReservationTearMessages})
    counters.EntityData.Leafs.Append("out-hello-messages", types.YLeaf{"OutHelloMessages", counters.OutHelloMessages})
    counters.EntityData.Leafs.Append("out-srefresh-messages", types.YLeaf{"OutSrefreshMessages", counters.OutSrefreshMessages})
    counters.EntityData.Leafs.Append("out-ack-messages", types.YLeaf{"OutAckMessages", counters.OutAckMessages})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef
// Reference to an interface or subinterface
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef_Config

    // Operational state for interface-ref.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef_State
}

func (interfaceRef *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "interface"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("config", types.YChild{"Config", &interfaceRef.Config})
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef_Config
// Configured reference to interface / subinterface
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (config *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-ref"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", config.Interface})
    config.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", config.Subinterface})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef_State
// Operational state for interface-ref
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations
// Enclosing container for bandwidth reservation
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Available and reserved bandwidth by priority on the interface. The type is
    // slice of
    // Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation.
    BandwidthReservation []*Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation
}

func (bandwidthReservations *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations) GetEntityData() *types.CommonEntityData {
    bandwidthReservations.EntityData.YFilter = bandwidthReservations.YFilter
    bandwidthReservations.EntityData.YangName = "bandwidth-reservations"
    bandwidthReservations.EntityData.BundleName = "openconfig"
    bandwidthReservations.EntityData.ParentYangName = "interface"
    bandwidthReservations.EntityData.SegmentPath = "bandwidth-reservations"
    bandwidthReservations.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bandwidthReservations.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bandwidthReservations.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bandwidthReservations.EntityData.Children = types.NewOrderedMap()
    bandwidthReservations.EntityData.Children.Append("bandwidth-reservation", types.YChild{"BandwidthReservation", nil})
    for i := range bandwidthReservations.BandwidthReservation {
        bandwidthReservations.EntityData.Children.Append(types.GetSegmentPath(bandwidthReservations.BandwidthReservation[i]), types.YChild{"BandwidthReservation", bandwidthReservations.BandwidthReservation[i]})
    }
    bandwidthReservations.EntityData.Leafs = types.NewOrderedMap()

    bandwidthReservations.EntityData.YListKeys = []string {}

    return &(bandwidthReservations.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation
// Available and reserved bandwidth by priority on
// the interface.
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the RSVP priority level. The type is
    // one of the following types: int with range: 0..7, or
    // :go:struct:`NetworkInstances_NetworkInstance_Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State_Priority
    // <ydk/models/network_instance/NetworkInstances_NetworkInstance_Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State_Priority>`.
    Priority interface{}

    // Operational state parameters relating to a bandwidth reservation at a
    // certain priority.
    State Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State
}

func (bandwidthReservation *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation) GetEntityData() *types.CommonEntityData {
    bandwidthReservation.EntityData.YFilter = bandwidthReservation.YFilter
    bandwidthReservation.EntityData.YangName = "bandwidth-reservation"
    bandwidthReservation.EntityData.BundleName = "openconfig"
    bandwidthReservation.EntityData.ParentYangName = "bandwidth-reservations"
    bandwidthReservation.EntityData.SegmentPath = "bandwidth-reservation" + types.AddKeyToken(bandwidthReservation.Priority, "priority")
    bandwidthReservation.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bandwidthReservation.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bandwidthReservation.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bandwidthReservation.EntityData.Children = types.NewOrderedMap()
    bandwidthReservation.EntityData.Children.Append("state", types.YChild{"State", &bandwidthReservation.State})
    bandwidthReservation.EntityData.Leafs = types.NewOrderedMap()
    bandwidthReservation.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", bandwidthReservation.Priority})

    bandwidthReservation.EntityData.YListKeys = []string {"Priority"}

    return &(bandwidthReservation.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State
// Operational state parameters relating to a
// bandwidth reservation at a certain priority
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP priority level for LSPs traversing the interface. The type is one of
    // the following types: int with range: 0..7, or enumeration
    // NetworkInstances.NetworkInstance.Mpls.SignalingProtocols.RsvpTe.InterfaceAttributes.Interface.BandwidthReservations.BandwidthReservation.State.Priority.
    Priority interface{}

    // Bandwidth currently available with the priority level, or for the entire
    // interface when the priority is set to ALL. The type is interface{} with
    // range: 0..18446744073709551615.
    AvailableBandwidth interface{}

    // Bandwidth currently reserved within the priority level, or the sum of all
    // priority levels when the keyword is set to ALL. The type is interface{}
    // with range: 0..18446744073709551615.
    ReservedBandwidth interface{}

    // Number of active RSVP reservations in the associated priority, or the sum
    // of all reservations when the priority level is set to ALL. The type is
    // interface{} with range: 0..18446744073709551615.
    ActiveReservationsCount interface{}

    // Maximum bandwidth reserved on the interface within the priority, or across
    // all priorities in the case that the priority level is set to ALL. The type
    // is interface{} with range: 0..18446744073709551615.
    HighwaterMark interface{}
}

func (state *Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "bandwidth-reservation"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", state.Priority})
    state.EntityData.Leafs.Append("available-bandwidth", types.YLeaf{"AvailableBandwidth", state.AvailableBandwidth})
    state.EntityData.Leafs.Append("reserved-bandwidth", types.YLeaf{"ReservedBandwidth", state.ReservedBandwidth})
    state.EntityData.Leafs.Append("active-reservations-count", types.YLeaf{"ActiveReservationsCount", state.ActiveReservationsCount})
    state.EntityData.Leafs.Append("highwater-mark", types.YLeaf{"HighwaterMark", state.HighwaterMark})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State_Priority represents RSVP priority level for LSPs traversing the interface
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State_Priority string

const (
    // The ALL keyword represents the overall
    // state of the interface - i.e., the union
    // of all of the priority levels
    Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State_Priority_ALL Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_BandwidthReservations_BandwidthReservation_State_Priority = "ALL"
)

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

    hellos.EntityData.Children = types.NewOrderedMap()
    hellos.EntityData.Children.Append("config", types.YChild{"Config", &hellos.Config})
    hellos.EntityData.Children.Append("state", types.YChild{"State", &hellos.State})
    hellos.EntityData.Leafs = types.NewOrderedMap()

    hellos.EntityData.YListKeys = []string {}

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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", config.HelloInterval})
    config.EntityData.Leafs.Append("refresh-reduction", types.YLeaf{"RefreshReduction", config.RefreshReduction})

    config.EntityData.YListKeys = []string {}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", state.HelloInterval})
    state.EntityData.Leafs.Append("refresh-reduction", types.YLeaf{"RefreshReduction", state.RefreshReduction})

    state.EntityData.YListKeys = []string {}

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

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Children.Append("config", types.YChild{"Config", &authentication.Config})
    authentication.EntityData.Children.Append("state", types.YChild{"State", &authentication.State})
    authentication.EntityData.Leafs = types.NewOrderedMap()

    authentication.EntityData.YListKeys = []string {}

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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", config.Enable})
    config.EntityData.Leafs.Append("authentication-key", types.YLeaf{"AuthenticationKey", config.AuthenticationKey})

    config.EntityData.YListKeys = []string {}

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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", state.Enable})
    state.EntityData.Leafs.Append("authentication-key", types.YLeaf{"AuthenticationKey", state.AuthenticationKey})

    state.EntityData.YListKeys = []string {}

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

    subscription.EntityData.Children = types.NewOrderedMap()
    subscription.EntityData.Children.Append("config", types.YChild{"Config", &subscription.Config})
    subscription.EntityData.Children.Append("state", types.YChild{"State", &subscription.State})
    subscription.EntityData.Leafs = types.NewOrderedMap()

    subscription.EntityData.YListKeys = []string {}

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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("subscription", types.YLeaf{"Subscription", config.Subscription})

    config.EntityData.YListKeys = []string {}

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

    // The calculated absolute value of the bandwidth which is reservable to
    // RSVP-TE on the interface prior to any adjustments that may be made from
    // external sources. The type is interface{} with range:
    // 0..18446744073709551615. Units are kbps.
    CalculatedAbsoluteSubscriptionBw interface{}
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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("subscription", types.YLeaf{"Subscription", state.Subscription})
    state.EntityData.Leafs.Append("calculated-absolute-subscription-bw", types.YLeaf{"CalculatedAbsoluteSubscriptionBw", state.CalculatedAbsoluteSubscriptionBw})

    state.EntityData.YListKeys = []string {}

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

    protection.EntityData.Children = types.NewOrderedMap()
    protection.EntityData.Children.Append("config", types.YChild{"Config", &protection.Config})
    protection.EntityData.Children.Append("state", types.YChild{"State", &protection.State})
    protection.EntityData.Leafs = types.NewOrderedMap()

    protection.EntityData.YListKeys = []string {}

    return &(protection.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config
// Configuration for link-protection
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Style of mpls frr protection desired: link, link-node, or unprotected. The
    // type is one of the following:
    // LINKPROTECTIONREQUIREDLINKNODEPROTECTIONREQUESTEDUNPROTECTED. The default
    // value is oc-mplst:LINK_NODE_PROTECTION_REQUESTED.
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

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("link-protection-style-requested", types.YLeaf{"LinkProtectionStyleRequested", config.LinkProtectionStyleRequested})
    config.EntityData.Leafs.Append("bypass-optimize-interval", types.YLeaf{"BypassOptimizeInterval", config.BypassOptimizeInterval})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State
// State for link-protection
type Mpls_SignalingProtocols_RsvpTe_InterfaceAttributes_Interface_Protection_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Style of mpls frr protection desired: link, link-node, or unprotected. The
    // type is one of the following:
    // LINKPROTECTIONREQUIREDLINKNODEPROTECTIONREQUESTEDUNPROTECTED. The default
    // value is oc-mplst:LINK_NODE_PROTECTION_REQUESTED.
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

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("link-protection-style-requested", types.YLeaf{"LinkProtectionStyleRequested", state.LinkProtectionStyleRequested})
    state.EntityData.Leafs.Append("bypass-optimize-interval", types.YLeaf{"BypassOptimizeInterval", state.BypassOptimizeInterval})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_Ldp
// LDP global signaling configuration
type Mpls_SignalingProtocols_Ldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
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

    ldp.EntityData.Children = types.NewOrderedMap()
    ldp.EntityData.Leafs = types.NewOrderedMap()

    ldp.EntityData.YListKeys = []string {}

    return &(ldp.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting
// MPLS-specific Segment Routing configuration and operational state
// parameters
type Mpls_SignalingProtocols_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-SID counters aggregated across all interfaces on the local system.
    AggregateSidCounters Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters

    // Interface related Segment Routing parameters.
    Interfaces Mpls_SignalingProtocols_SegmentRouting_Interfaces
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

    segmentRouting.EntityData.Children = types.NewOrderedMap()
    segmentRouting.EntityData.Children.Append("aggregate-sid-counters", types.YChild{"AggregateSidCounters", &segmentRouting.AggregateSidCounters})
    segmentRouting.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &segmentRouting.Interfaces})
    segmentRouting.EntityData.Leafs = types.NewOrderedMap()

    segmentRouting.EntityData.YListKeys = []string {}

    return &(segmentRouting.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters
// Per-SID counters aggregated across all interfaces on the local system
type Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Counters aggregated across all of the interfaces of the local system
    // corresponding to traffic received or forwarded with a particular SID. The
    // type is slice of
    // Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter.
    AggregateSidCounter []*Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter
}

func (aggregateSidCounters *Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters) GetEntityData() *types.CommonEntityData {
    aggregateSidCounters.EntityData.YFilter = aggregateSidCounters.YFilter
    aggregateSidCounters.EntityData.YangName = "aggregate-sid-counters"
    aggregateSidCounters.EntityData.BundleName = "openconfig"
    aggregateSidCounters.EntityData.ParentYangName = "segment-routing"
    aggregateSidCounters.EntityData.SegmentPath = "aggregate-sid-counters"
    aggregateSidCounters.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregateSidCounters.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregateSidCounters.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregateSidCounters.EntityData.Children = types.NewOrderedMap()
    aggregateSidCounters.EntityData.Children.Append("aggregate-sid-counter", types.YChild{"AggregateSidCounter", nil})
    for i := range aggregateSidCounters.AggregateSidCounter {
        aggregateSidCounters.EntityData.Children.Append(types.GetSegmentPath(aggregateSidCounters.AggregateSidCounter[i]), types.YChild{"AggregateSidCounter", aggregateSidCounters.AggregateSidCounter[i]})
    }
    aggregateSidCounters.EntityData.Leafs = types.NewOrderedMap()

    aggregateSidCounters.EntityData.YListKeys = []string {}

    return &(aggregateSidCounters.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter
// Counters aggregated across all of the interfaces of the local
// system corresponding to traffic received or forwarded with a
// particular SID
type Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MPLS label representing the segment
    // identifier. The type is one of the following types: int with range:
    // 16..1048575, or :go:struct:`MplsLabel
    // <ydk/models/segment_routing/MplsLabel>`.
    MplsLabel interface{}

    // State parameters for per-SID statistics.
    State Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter_State
}

func (aggregateSidCounter *Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter) GetEntityData() *types.CommonEntityData {
    aggregateSidCounter.EntityData.YFilter = aggregateSidCounter.YFilter
    aggregateSidCounter.EntityData.YangName = "aggregate-sid-counter"
    aggregateSidCounter.EntityData.BundleName = "openconfig"
    aggregateSidCounter.EntityData.ParentYangName = "aggregate-sid-counters"
    aggregateSidCounter.EntityData.SegmentPath = "aggregate-sid-counter" + types.AddKeyToken(aggregateSidCounter.MplsLabel, "mpls-label")
    aggregateSidCounter.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregateSidCounter.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregateSidCounter.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregateSidCounter.EntityData.Children = types.NewOrderedMap()
    aggregateSidCounter.EntityData.Children.Append("state", types.YChild{"State", &aggregateSidCounter.State})
    aggregateSidCounter.EntityData.Leafs = types.NewOrderedMap()
    aggregateSidCounter.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", aggregateSidCounter.MplsLabel})

    aggregateSidCounter.EntityData.YListKeys = []string {"MplsLabel"}

    return &(aggregateSidCounter.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter_State
// State parameters for per-SID statistics
type Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The MPLS label used for the segment identifier. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    MplsLabel interface{}

    // A cumulative counter of the packets received within the context which have
    // matched a label corresponding to an SR Segment Identifier. The type is
    // interface{} with range: 0..18446744073709551615.
    InPkts interface{}

    // The cumulative counter of the total bytes received within the context which
    // have matched a label corresponding to an SR Segment Identifier. The type is
    // interface{} with range: 0..18446744073709551615.
    InOctets interface{}

    // A cumulative counter of the total number of packets transmitted by the
    // local system within the context which have a label imposed that corresponds
    // to an Segment Identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPkts interface{}

    // A cumulative counter of the total bytes transmitted by the local system
    // within the context which have a label imported that corresponds to an SR
    // Segment Identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctets interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_AggregateSidCounters_AggregateSidCounter_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "aggregate-sid-counter"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", state.MplsLabel})
    state.EntityData.Leafs.Append("in-pkts", types.YLeaf{"InPkts", state.InPkts})
    state.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", state.InOctets})
    state.EntityData.Leafs.Append("out-pkts", types.YLeaf{"OutPkts", state.OutPkts})
    state.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", state.OutOctets})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces
// Interface related Segment Routing parameters.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Parameters and MPLS-specific configuration relating to Segment Routing on
    // an interface. The type is slice of
    // Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface.
    Interface []*Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface
}

func (interfaces *Mpls_SignalingProtocols_SegmentRouting_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "openconfig"
    interfaces.EntityData.ParentYangName = "segment-routing"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaces.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface
// Parameters and MPLS-specific configuration relating to Segment
// Routing on an interface.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A reference to the ID for the interface for which
    // SR is configured. The type is string. Refers to
    // mpls.Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_Config_InterfaceId
    InterfaceId interface{}

    // MPLS-specific Segment Routing configuration parameters related to an
    // interface.
    Config Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_Config

    // MPLS-specific Segment Routing operational state parameters related to an
    // interface.
    State Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_State

    // Per-SID statistics for MPLS.
    SidCounters Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters

    // Reference to an interface or subinterface.
    InterfaceRef Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef
}

func (self *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceId, "interface-id")
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("config", types.YChild{"Config", &self.Config})
    self.EntityData.Children.Append("state", types.YChild{"State", &self.State})
    self.EntityData.Children.Append("sid-counters", types.YChild{"SidCounters", &self.SidCounters})
    self.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &self.InterfaceRef})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", self.InterfaceId})

    self.EntityData.YListKeys = []string {"InterfaceId"}

    return &(self.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_Config
// MPLS-specific Segment Routing configuration parameters
// related to an interface.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A unique identifier for the interface. The type is string.
    InterfaceId interface{}
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", config.InterfaceId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_State
// MPLS-specific Segment Routing operational state parameters
// related to an interface.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A unique identifier for the interface. The type is string.
    InterfaceId interface{}

    // A cumulative counter of the packets received within the context which have
    // matched a label corresponding to an SR Segment Identifier. The type is
    // interface{} with range: 0..18446744073709551615.
    InPkts interface{}

    // The cumulative counter of the total bytes received within the context which
    // have matched a label corresponding to an SR Segment Identifier. The type is
    // interface{} with range: 0..18446744073709551615.
    InOctets interface{}

    // A cumulative counter of the total number of packets transmitted by the
    // local system within the context which have a label imposed that corresponds
    // to an Segment Identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPkts interface{}

    // A cumulative counter of the total bytes transmitted by the local system
    // within the context which have a label imported that corresponds to an SR
    // Segment Identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctets interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface-id", types.YLeaf{"InterfaceId", state.InterfaceId})
    state.EntityData.Leafs.Append("in-pkts", types.YLeaf{"InPkts", state.InPkts})
    state.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", state.InOctets})
    state.EntityData.Leafs.Append("out-pkts", types.YLeaf{"OutPkts", state.OutPkts})
    state.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", state.OutOctets})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters
// Per-SID statistics for MPLS
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per segment identifier counters for the MPLS dataplane. The type is slice
    // of
    // Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter.
    SidCounter []*Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter
}

func (sidCounters *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters) GetEntityData() *types.CommonEntityData {
    sidCounters.EntityData.YFilter = sidCounters.YFilter
    sidCounters.EntityData.YangName = "sid-counters"
    sidCounters.EntityData.BundleName = "openconfig"
    sidCounters.EntityData.ParentYangName = "interface"
    sidCounters.EntityData.SegmentPath = "sid-counters"
    sidCounters.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sidCounters.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sidCounters.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sidCounters.EntityData.Children = types.NewOrderedMap()
    sidCounters.EntityData.Children.Append("sid-counter", types.YChild{"SidCounter", nil})
    for i := range sidCounters.SidCounter {
        sidCounters.EntityData.Children.Append(types.GetSegmentPath(sidCounters.SidCounter[i]), types.YChild{"SidCounter", sidCounters.SidCounter[i]})
    }
    sidCounters.EntityData.Leafs = types.NewOrderedMap()

    sidCounters.EntityData.YListKeys = []string {}

    return &(sidCounters.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter
// Per segment identifier counters for the MPLS dataplane.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The MPLS label representing the segment
    // identifier. The type is one of the following types: int with range:
    // 16..1048575, or :go:struct:`MplsLabel
    // <ydk/models/segment_routing/MplsLabel>`.
    MplsLabel interface{}

    // State parameters for per-SID statistics.
    State Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_State

    // Per-SID per-forwarding class counters for Segment Routing.
    ForwardingClasses Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses
}

func (sidCounter *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter) GetEntityData() *types.CommonEntityData {
    sidCounter.EntityData.YFilter = sidCounter.YFilter
    sidCounter.EntityData.YangName = "sid-counter"
    sidCounter.EntityData.BundleName = "openconfig"
    sidCounter.EntityData.ParentYangName = "sid-counters"
    sidCounter.EntityData.SegmentPath = "sid-counter" + types.AddKeyToken(sidCounter.MplsLabel, "mpls-label")
    sidCounter.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sidCounter.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sidCounter.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sidCounter.EntityData.Children = types.NewOrderedMap()
    sidCounter.EntityData.Children.Append("state", types.YChild{"State", &sidCounter.State})
    sidCounter.EntityData.Children.Append("forwarding-classes", types.YChild{"ForwardingClasses", &sidCounter.ForwardingClasses})
    sidCounter.EntityData.Leafs = types.NewOrderedMap()
    sidCounter.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", sidCounter.MplsLabel})

    sidCounter.EntityData.YListKeys = []string {"MplsLabel"}

    return &(sidCounter.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_State
// State parameters for per-SID statistics
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The MPLS label used for the segment identifier. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    MplsLabel interface{}

    // A cumulative counter of the packets received within the context which have
    // matched a label corresponding to an SR Segment Identifier. The type is
    // interface{} with range: 0..18446744073709551615.
    InPkts interface{}

    // The cumulative counter of the total bytes received within the context which
    // have matched a label corresponding to an SR Segment Identifier. The type is
    // interface{} with range: 0..18446744073709551615.
    InOctets interface{}

    // A cumulative counter of the total number of packets transmitted by the
    // local system within the context which have a label imposed that corresponds
    // to an Segment Identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPkts interface{}

    // A cumulative counter of the total bytes transmitted by the local system
    // within the context which have a label imported that corresponds to an SR
    // Segment Identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctets interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "sid-counter"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", state.MplsLabel})
    state.EntityData.Leafs.Append("in-pkts", types.YLeaf{"InPkts", state.InPkts})
    state.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", state.InOctets})
    state.EntityData.Leafs.Append("out-pkts", types.YLeaf{"OutPkts", state.OutPkts})
    state.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", state.OutOctets})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses
// Per-SID per-forwarding class counters for Segment Routing.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID entries for the forwarding class associated with the referenced MPLS
    // EXP. The type is slice of
    // Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass.
    ForwardingClass []*Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass
}

func (forwardingClasses *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses) GetEntityData() *types.CommonEntityData {
    forwardingClasses.EntityData.YFilter = forwardingClasses.YFilter
    forwardingClasses.EntityData.YangName = "forwarding-classes"
    forwardingClasses.EntityData.BundleName = "openconfig"
    forwardingClasses.EntityData.ParentYangName = "sid-counter"
    forwardingClasses.EntityData.SegmentPath = "forwarding-classes"
    forwardingClasses.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    forwardingClasses.EntityData.NamespaceTable = openconfig.GetNamespaces()
    forwardingClasses.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    forwardingClasses.EntityData.Children = types.NewOrderedMap()
    forwardingClasses.EntityData.Children.Append("forwarding-class", types.YChild{"ForwardingClass", nil})
    for i := range forwardingClasses.ForwardingClass {
        forwardingClasses.EntityData.Children.Append(types.GetSegmentPath(forwardingClasses.ForwardingClass[i]), types.YChild{"ForwardingClass", forwardingClasses.ForwardingClass[i]})
    }
    forwardingClasses.EntityData.Leafs = types.NewOrderedMap()

    forwardingClasses.EntityData.YListKeys = []string {}

    return &(forwardingClasses.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass
// SID entries for the forwarding class associated with the
// referenced MPLS EXP.
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the value of the EXP bits of the
    // segment identifier. The type is string with range: 0..7. Refers to
    // mpls.Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass_State_Exp
    Exp interface{}

    // Per-SID, per forwarding class counters for Segment Routing with the MPLS
    // dataplane.
    State Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass_State
}

func (forwardingClass *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass) GetEntityData() *types.CommonEntityData {
    forwardingClass.EntityData.YFilter = forwardingClass.YFilter
    forwardingClass.EntityData.YangName = "forwarding-class"
    forwardingClass.EntityData.BundleName = "openconfig"
    forwardingClass.EntityData.ParentYangName = "forwarding-classes"
    forwardingClass.EntityData.SegmentPath = "forwarding-class" + types.AddKeyToken(forwardingClass.Exp, "exp")
    forwardingClass.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    forwardingClass.EntityData.NamespaceTable = openconfig.GetNamespaces()
    forwardingClass.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    forwardingClass.EntityData.Children = types.NewOrderedMap()
    forwardingClass.EntityData.Children.Append("state", types.YChild{"State", &forwardingClass.State})
    forwardingClass.EntityData.Leafs = types.NewOrderedMap()
    forwardingClass.EntityData.Leafs.Append("exp", types.YLeaf{"Exp", forwardingClass.Exp})

    forwardingClass.EntityData.YListKeys = []string {"Exp"}

    return &(forwardingClass.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass_State
// Per-SID, per forwarding class counters for Segment Routing
// with the MPLS dataplane
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of the MPLS EXP (experimental) or Traffic Class bits that the SID
    // statistics relate to. Packets received with a MPLS label value equal to the
    // SID's MPLS label and EXP bits equal to the this value should be counted
    // towards the associated ingress statistics. Packets that are forwarded to
    // the destination MPLS label corresponding to the SID should be counted
    // towards this value. In the egress direction, where forwarding follows a SID
    // value that requires PHP at the local node, packets should still be counted
    // towards the egress counters. The type is interface{} with range: 0..7.
    Exp interface{}

    // A cumulative counter of the packets received within the context which have
    // matched a label corresponding to an SR Segment Identifier. The type is
    // interface{} with range: 0..18446744073709551615.
    InPkts interface{}

    // The cumulative counter of the total bytes received within the context which
    // have matched a label corresponding to an SR Segment Identifier. The type is
    // interface{} with range: 0..18446744073709551615.
    InOctets interface{}

    // A cumulative counter of the total number of packets transmitted by the
    // local system within the context which have a label imposed that corresponds
    // to an Segment Identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPkts interface{}

    // A cumulative counter of the total bytes transmitted by the local system
    // within the context which have a label imported that corresponds to an SR
    // Segment Identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctets interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_SidCounters_SidCounter_ForwardingClasses_ForwardingClass_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "forwarding-class"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("exp", types.YLeaf{"Exp", state.Exp})
    state.EntityData.Leafs.Append("in-pkts", types.YLeaf{"InPkts", state.InPkts})
    state.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", state.InOctets})
    state.EntityData.Leafs.Append("out-pkts", types.YLeaf{"OutPkts", state.OutPkts})
    state.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", state.OutOctets})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef
// Reference to an interface or subinterface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef_Config

    // Operational state for interface-ref.
    State Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef_State
}

func (interfaceRef *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "interface"
    interfaceRef.EntityData.SegmentPath = "interface-ref"
    interfaceRef.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceRef.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceRef.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceRef.EntityData.Children = types.NewOrderedMap()
    interfaceRef.EntityData.Children.Append("config", types.YChild{"Config", &interfaceRef.Config})
    interfaceRef.EntityData.Children.Append("state", types.YChild{"State", &interfaceRef.State})
    interfaceRef.EntityData.Leafs = types.NewOrderedMap()

    interfaceRef.EntityData.YListKeys = []string {}

    return &(interfaceRef.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef_Config
// Configured reference to interface / subinterface
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (config *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-ref"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", config.Interface})
    config.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", config.Subinterface})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef_State
// Operational state for interface-ref
type Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a base interface.  If a reference to a subinterface is
    // required, this leaf must be specified to indicate the base interface. The
    // type is string. Refers to interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Reference to a subinterface -- this requires the base interface to be
    // specified using the interface leaf in this container.  If only a reference
    // to a base interface is requuired, this leaf should not be set. The type is
    // string with range: 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Index
    Subinterface interface{}
}

func (state *Mpls_SignalingProtocols_SegmentRouting_Interfaces_Interface_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-ref"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", state.Interface})
    state.EntityData.Leafs.Append("subinterface", types.YLeaf{"Subinterface", state.Subinterface})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
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

    lsps.EntityData.Children = types.NewOrderedMap()
    lsps.EntityData.Children.Append("constrained-path", types.YChild{"ConstrainedPath", &lsps.ConstrainedPath})
    lsps.EntityData.Children.Append("unconstrained-path", types.YChild{"UnconstrainedPath", &lsps.UnconstrainedPath})
    lsps.EntityData.Children.Append("static-lsps", types.YChild{"StaticLsps", &lsps.StaticLsps})
    lsps.EntityData.Leafs = types.NewOrderedMap()

    lsps.EntityData.YListKeys = []string {}

    return &(lsps.EntityData)
}

// Mpls_Lsps_ConstrainedPath
// traffic-engineered LSPs supporting different
// path computation and signaling methods
type Mpls_Lsps_ConstrainedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for the named explicit paths.
    NamedExplicitPaths Mpls_Lsps_ConstrainedPath_NamedExplicitPaths

    // Enclosing container for tunnels.
    Tunnels Mpls_Lsps_ConstrainedPath_Tunnels
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

    constrainedPath.EntityData.Children = types.NewOrderedMap()
    constrainedPath.EntityData.Children.Append("named-explicit-paths", types.YChild{"NamedExplicitPaths", &constrainedPath.NamedExplicitPaths})
    constrainedPath.EntityData.Children.Append("tunnels", types.YChild{"Tunnels", &constrainedPath.Tunnels})
    constrainedPath.EntityData.Leafs = types.NewOrderedMap()

    constrainedPath.EntityData.YListKeys = []string {}

    return &(constrainedPath.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths
// Enclosing container for the named explicit paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of explicit paths. The type is slice of
    // Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath.
    NamedExplicitPath []*Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath
}

func (namedExplicitPaths *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths) GetEntityData() *types.CommonEntityData {
    namedExplicitPaths.EntityData.YFilter = namedExplicitPaths.YFilter
    namedExplicitPaths.EntityData.YangName = "named-explicit-paths"
    namedExplicitPaths.EntityData.BundleName = "openconfig"
    namedExplicitPaths.EntityData.ParentYangName = "constrained-path"
    namedExplicitPaths.EntityData.SegmentPath = "named-explicit-paths"
    namedExplicitPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    namedExplicitPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    namedExplicitPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    namedExplicitPaths.EntityData.Children = types.NewOrderedMap()
    namedExplicitPaths.EntityData.Children.Append("named-explicit-path", types.YChild{"NamedExplicitPath", nil})
    for i := range namedExplicitPaths.NamedExplicitPath {
        namedExplicitPaths.EntityData.Children.Append(types.GetSegmentPath(namedExplicitPaths.NamedExplicitPath[i]), types.YChild{"NamedExplicitPath", namedExplicitPaths.NamedExplicitPath[i]})
    }
    namedExplicitPaths.EntityData.Leafs = types.NewOrderedMap()

    namedExplicitPaths.EntityData.YListKeys = []string {}

    return &(namedExplicitPaths.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath
// A list of explicit paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A string name that uniquely identifies an explicit
    // path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_Name
    Name interface{}

    // Configuration parameters relating to named explicit paths.
    Config Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config

    // Operational state parameters relating to the named explicit paths.
    State Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State

    // Enclosing container for EROs.
    ExplicitRouteObjects Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects
}

func (namedExplicitPath *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath) GetEntityData() *types.CommonEntityData {
    namedExplicitPath.EntityData.YFilter = namedExplicitPath.YFilter
    namedExplicitPath.EntityData.YangName = "named-explicit-path"
    namedExplicitPath.EntityData.BundleName = "openconfig"
    namedExplicitPath.EntityData.ParentYangName = "named-explicit-paths"
    namedExplicitPath.EntityData.SegmentPath = "named-explicit-path" + types.AddKeyToken(namedExplicitPath.Name, "name")
    namedExplicitPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    namedExplicitPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    namedExplicitPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    namedExplicitPath.EntityData.Children = types.NewOrderedMap()
    namedExplicitPath.EntityData.Children.Append("config", types.YChild{"Config", &namedExplicitPath.Config})
    namedExplicitPath.EntityData.Children.Append("state", types.YChild{"State", &namedExplicitPath.State})
    namedExplicitPath.EntityData.Children.Append("explicit-route-objects", types.YChild{"ExplicitRouteObjects", &namedExplicitPath.ExplicitRouteObjects})
    namedExplicitPath.EntityData.Leafs = types.NewOrderedMap()
    namedExplicitPath.EntityData.Leafs.Append("name", types.YLeaf{"Name", namedExplicitPath.Name})

    namedExplicitPath.EntityData.YListKeys = []string {"Name"}

    return &(namedExplicitPath.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config
// Configuration parameters relating to named explicit
// paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A string name that uniquely identifies an explicit path. The type is
    // string.
    Name interface{}

    // The restrictions placed on the SIDs to be selected by the calculation
    // method for the explicit path when it is instantiated for a SR-TE LSP. The
    // type is SidSelectionMode. The default value is MIXED_MODE.
    SidSelectionMode interface{}

    // When this value is set to true, only SIDs that are protected are to be
    // selected by the calculating method when the explicit path is instantiated
    // by a SR-TE LSP. The type is bool. The default value is false.
    SidProtectionRequired interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "named-explicit-path"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("sid-selection-mode", types.YLeaf{"SidSelectionMode", config.SidSelectionMode})
    config.EntityData.Leafs.Append("sid-protection-required", types.YLeaf{"SidProtectionRequired", config.SidProtectionRequired})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_SidSelectionMode represents instantiated for a SR-TE LSP
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_SidSelectionMode string

const (
    // The SR-TE tunnel should only use adjacency SIDs
    // to build the SID stack to be pushed for the LSP
    Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_SidSelectionMode_ADJ_SID_ONLY Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_SidSelectionMode = "ADJ_SID_ONLY"

    // The SR-TE tunnel can use a mix of adjacency
    // and prefix SIDs to build the SID stack to be pushed
    // to the LSP
    Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_SidSelectionMode_MIXED_MODE Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_SidSelectionMode = "MIXED_MODE"
)

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State
// Operational state parameters relating to the named
// explicit paths
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A string name that uniquely identifies an explicit path. The type is
    // string.
    Name interface{}

    // The restrictions placed on the SIDs to be selected by the calculation
    // method for the explicit path when it is instantiated for a SR-TE LSP. The
    // type is SidSelectionMode. The default value is MIXED_MODE.
    SidSelectionMode interface{}

    // When this value is set to true, only SIDs that are protected are to be
    // selected by the calculating method when the explicit path is instantiated
    // by a SR-TE LSP. The type is bool. The default value is false.
    SidProtectionRequired interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "named-explicit-path"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("sid-selection-mode", types.YLeaf{"SidSelectionMode", state.SidSelectionMode})
    state.EntityData.Leafs.Append("sid-protection-required", types.YLeaf{"SidProtectionRequired", state.SidProtectionRequired})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State_SidSelectionMode represents instantiated for a SR-TE LSP
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State_SidSelectionMode string

const (
    // The SR-TE tunnel should only use adjacency SIDs
    // to build the SID stack to be pushed for the LSP
    Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State_SidSelectionMode_ADJ_SID_ONLY Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State_SidSelectionMode = "ADJ_SID_ONLY"

    // The SR-TE tunnel can use a mix of adjacency
    // and prefix SIDs to build the SID stack to be pushed
    // to the LSP
    Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State_SidSelectionMode_MIXED_MODE Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_State_SidSelectionMode = "MIXED_MODE"
)

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects
// Enclosing container for EROs
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of explicit route objects. The type is slice of
    // Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject.
    ExplicitRouteObject []*Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject
}

func (explicitRouteObjects *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects) GetEntityData() *types.CommonEntityData {
    explicitRouteObjects.EntityData.YFilter = explicitRouteObjects.YFilter
    explicitRouteObjects.EntityData.YangName = "explicit-route-objects"
    explicitRouteObjects.EntityData.BundleName = "openconfig"
    explicitRouteObjects.EntityData.ParentYangName = "named-explicit-path"
    explicitRouteObjects.EntityData.SegmentPath = "explicit-route-objects"
    explicitRouteObjects.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    explicitRouteObjects.EntityData.NamespaceTable = openconfig.GetNamespaces()
    explicitRouteObjects.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    explicitRouteObjects.EntityData.Children = types.NewOrderedMap()
    explicitRouteObjects.EntityData.Children.Append("explicit-route-object", types.YChild{"ExplicitRouteObject", nil})
    for i := range explicitRouteObjects.ExplicitRouteObject {
        explicitRouteObjects.EntityData.Children.Append(types.GetSegmentPath(explicitRouteObjects.ExplicitRouteObject[i]), types.YChild{"ExplicitRouteObject", explicitRouteObjects.ExplicitRouteObject[i]})
    }
    explicitRouteObjects.EntityData.Leafs = types.NewOrderedMap()

    explicitRouteObjects.EntityData.YListKeys = []string {}

    return &(explicitRouteObjects.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject
// List of explicit route objects
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of this explicit route object, to express
    // the order of hops in path. The type is string with range: 0..255. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_Config_Index
    Index interface{}

    // Configuration parameters relating to an explicit route.
    Config Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_Config

    // State parameters relating to an explicit route.
    State Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_State
}

func (explicitRouteObject *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject) GetEntityData() *types.CommonEntityData {
    explicitRouteObject.EntityData.YFilter = explicitRouteObject.YFilter
    explicitRouteObject.EntityData.YangName = "explicit-route-object"
    explicitRouteObject.EntityData.BundleName = "openconfig"
    explicitRouteObject.EntityData.ParentYangName = "explicit-route-objects"
    explicitRouteObject.EntityData.SegmentPath = "explicit-route-object" + types.AddKeyToken(explicitRouteObject.Index, "index")
    explicitRouteObject.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    explicitRouteObject.EntityData.NamespaceTable = openconfig.GetNamespaces()
    explicitRouteObject.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    explicitRouteObject.EntityData.Children = types.NewOrderedMap()
    explicitRouteObject.EntityData.Children.Append("config", types.YChild{"Config", &explicitRouteObject.Config})
    explicitRouteObject.EntityData.Children.Append("state", types.YChild{"State", &explicitRouteObject.State})
    explicitRouteObject.EntityData.Leafs = types.NewOrderedMap()
    explicitRouteObject.EntityData.Leafs.Append("index", types.YLeaf{"Index", explicitRouteObject.Index})

    explicitRouteObject.EntityData.YListKeys = []string {"Index"}

    return &(explicitRouteObject.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_Config
// Configuration parameters relating to an explicit
// route
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // router hop for the LSP path. The type is one of the following types: string
    // with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    Address interface{}

    // strict or loose hop. The type is MplsHopType.
    HopType interface{}

    // Index of this explicit route object to express the order of hops in the
    // path. The type is interface{} with range: 0..255.
    Index interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "explicit-route-object"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("address", types.YLeaf{"Address", config.Address})
    config.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", config.HopType})
    config.EntityData.Leafs.Append("index", types.YLeaf{"Index", config.Index})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_State
// State parameters relating to an explicit route
type Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // router hop for the LSP path. The type is one of the following types: string
    // with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    Address interface{}

    // strict or loose hop. The type is MplsHopType.
    HopType interface{}

    // Index of this explicit route object to express the order of hops in the
    // path. The type is interface{} with range: 0..255.
    Index interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_ExplicitRouteObjects_ExplicitRouteObject_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "explicit-route-object"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("address", types.YLeaf{"Address", state.Address})
    state.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", state.HopType})
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels
// Enclosing container for tunnels
type Mpls_Lsps_ConstrainedPath_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of TE tunnels. This list contains only the LSPs that the current
    // device originates (i.e., for which it is the head-end). Where the signaling
    // protocol utilised for an LSP allows a mid-point or tail device to be aware
    // of the LSP (e.g., RSVP-TE), then the associated sessions are maintained per
    // protocol. The type is slice of Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel.
    Tunnel []*Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel
}

func (tunnels *Mpls_Lsps_ConstrainedPath_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "openconfig"
    tunnels.EntityData.ParentYangName = "constrained-path"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    tunnels.EntityData.NamespaceTable = openconfig.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    tunnels.EntityData.Children = types.NewOrderedMap()
    tunnels.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", nil})
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children.Append(types.GetSegmentPath(tunnels.Tunnel[i]), types.YChild{"Tunnel", tunnels.Tunnel[i]})
    }
    tunnels.EntityData.Leafs = types.NewOrderedMap()

    tunnels.EntityData.YListKeys = []string {}

    return &(tunnels.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel
// List of TE tunnels. This list contains only the LSPs that the
// current device originates (i.e., for which it is the head-end).
// Where the signaling protocol utilised for an LSP allows a mid-point
// or tail device to be aware of the LSP (e.g., RSVP-TE), then the
// associated sessions are maintained per protocol
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The tunnel name. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Config_Name
    Name interface{}

    // Configuration parameters related to TE tunnels:.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Config

    // State parameters related to TE tunnels.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_State

    // Bandwidth configuration for TE LSPs.
    Bandwidth Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth

    // Parameters related to LSPs of type P2P.
    P2pTunnelAttributes Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes
}

func (tunnel *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "openconfig"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + types.AddKeyToken(tunnel.Name, "name")
    tunnel.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    tunnel.EntityData.NamespaceTable = openconfig.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("config", types.YChild{"Config", &tunnel.Config})
    tunnel.EntityData.Children.Append("state", types.YChild{"State", &tunnel.State})
    tunnel.EntityData.Children.Append("bandwidth", types.YChild{"Bandwidth", &tunnel.Bandwidth})
    tunnel.EntityData.Children.Append("p2p-tunnel-attributes", types.YChild{"P2pTunnelAttributes", &tunnel.P2pTunnelAttributes})
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("name", types.YLeaf{"Name", tunnel.Name})

    tunnel.EntityData.YListKeys = []string {"Name"}

    return &(tunnel.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Config
// Configuration parameters related to TE tunnels:
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The tunnel name. The type is string.
    Name interface{}

    // Tunnel type, p2p or p2mp. The type is one of the following: P2PP2MP.
    Type interface{}

    // Signaling protocol used to set up this tunnel. The type is one of the
    // following: PATHSETUPSRPATHSETUPRSVPPATHSETUPLDP.
    SignalingProtocol interface{}

    // optional text description for the tunnel. The type is string.
    Description interface{}

    // TE tunnel administrative state. The type is one of the following:
    // ADMINUPADMINDOWN. The default value is oc-mplst:ADMIN_UP.
    AdminStatus interface{}

    // Specifies a preference for this tunnel. A lower number signifies a better
    // preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // The type of metric specification that should be used to set the LSP(s)
    // metric. The type is one of the following:
    // LSPMETRICINHERITEDLSPMETRICABSOLUTELSPMETRICRELATIVE. The default value is
    // oc-mplst:LSP_METRIC_INHERITED.
    MetricType interface{}

    // The value of the metric that should be specified. The value supplied in
    // this leaf is used in conjunction with the metric type to determine the
    // value of the metric used by the system. Where the metric-type is set to
    // LSP_METRIC_ABSOLUTE - the value of this leaf is used directly; where it is
    // set to LSP_METRIC_RELATIVE, the relevant (positive or negative) offset is
    // used to formulate the metric; where metric-type is LSP_METRIC_INHERITED,
    // the value of this leaf is not utilised. The type is interface{} with range:
    // -2147483648..2147483647.
    Metric interface{}

    // Whether this LSP is considered to be eligible for us as a shortcut in the
    // IGP. In the case that this leaf is set to true, the IGP SPF calculation
    // uses the metric specified to determine whether traffic should be carried
    // over this LSP. The type is bool. The default value is true.
    ShortcutEligible interface{}

    // style of mpls frr protection desired: can be link, link-node or
    // unprotected. The type is one of the following:
    // LINKPROTECTIONREQUIREDLINKNODEPROTECTIONREQUESTEDUNPROTECTED. The default
    // value is oc-mplst:UNPROTECTED.
    ProtectionStyleRequested interface{}

    // frequency of reoptimization of a traffic engineered LSP. The type is
    // interface{} with range: 0..65535. Units are seconds.
    ReoptimizeTimer interface{}

    // RSVP-TE tunnel source address. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "tunnel"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("type", types.YLeaf{"Type", config.Type})
    config.EntityData.Leafs.Append("signaling-protocol", types.YLeaf{"SignalingProtocol", config.SignalingProtocol})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})
    config.EntityData.Leafs.Append("admin-status", types.YLeaf{"AdminStatus", config.AdminStatus})
    config.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", config.Preference})
    config.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", config.MetricType})
    config.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", config.Metric})
    config.EntityData.Leafs.Append("shortcut-eligible", types.YLeaf{"ShortcutEligible", config.ShortcutEligible})
    config.EntityData.Leafs.Append("protection-style-requested", types.YLeaf{"ProtectionStyleRequested", config.ProtectionStyleRequested})
    config.EntityData.Leafs.Append("reoptimize-timer", types.YLeaf{"ReoptimizeTimer", config.ReoptimizeTimer})
    config.EntityData.Leafs.Append("source", types.YLeaf{"Source", config.Source})
    config.EntityData.Leafs.Append("soft-preemption", types.YLeaf{"SoftPreemption", config.SoftPreemption})
    config.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", config.SetupPriority})
    config.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", config.HoldPriority})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_State
// State parameters related to TE tunnels
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The tunnel name. The type is string.
    Name interface{}

    // Tunnel type, p2p or p2mp. The type is one of the following: P2PP2MP.
    Type interface{}

    // Signaling protocol used to set up this tunnel. The type is one of the
    // following: PATHSETUPSRPATHSETUPRSVPPATHSETUPLDP.
    SignalingProtocol interface{}

    // optional text description for the tunnel. The type is string.
    Description interface{}

    // TE tunnel administrative state. The type is one of the following:
    // ADMINUPADMINDOWN. The default value is oc-mplst:ADMIN_UP.
    AdminStatus interface{}

    // Specifies a preference for this tunnel. A lower number signifies a better
    // preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // The type of metric specification that should be used to set the LSP(s)
    // metric. The type is one of the following:
    // LSPMETRICINHERITEDLSPMETRICABSOLUTELSPMETRICRELATIVE. The default value is
    // oc-mplst:LSP_METRIC_INHERITED.
    MetricType interface{}

    // The value of the metric that should be specified. The value supplied in
    // this leaf is used in conjunction with the metric type to determine the
    // value of the metric used by the system. Where the metric-type is set to
    // LSP_METRIC_ABSOLUTE - the value of this leaf is used directly; where it is
    // set to LSP_METRIC_RELATIVE, the relevant (positive or negative) offset is
    // used to formulate the metric; where metric-type is LSP_METRIC_INHERITED,
    // the value of this leaf is not utilised. The type is interface{} with range:
    // -2147483648..2147483647.
    Metric interface{}

    // Whether this LSP is considered to be eligible for us as a shortcut in the
    // IGP. In the case that this leaf is set to true, the IGP SPF calculation
    // uses the metric specified to determine whether traffic should be carried
    // over this LSP. The type is bool. The default value is true.
    ShortcutEligible interface{}

    // style of mpls frr protection desired: can be link, link-node or
    // unprotected. The type is one of the following:
    // LINKPROTECTIONREQUIREDLINKNODEPROTECTIONREQUESTEDUNPROTECTED. The default
    // value is oc-mplst:UNPROTECTED.
    ProtectionStyleRequested interface{}

    // frequency of reoptimization of a traffic engineered LSP. The type is
    // interface{} with range: 0..65535. Units are seconds.
    ReoptimizeTimer interface{}

    // RSVP-TE tunnel source address. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
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
    Counters Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_State_Counters
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "tunnel"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("counters", types.YChild{"Counters", &state.Counters})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})
    state.EntityData.Leafs.Append("signaling-protocol", types.YLeaf{"SignalingProtocol", state.SignalingProtocol})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("admin-status", types.YLeaf{"AdminStatus", state.AdminStatus})
    state.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", state.Preference})
    state.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", state.MetricType})
    state.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", state.Metric})
    state.EntityData.Leafs.Append("shortcut-eligible", types.YLeaf{"ShortcutEligible", state.ShortcutEligible})
    state.EntityData.Leafs.Append("protection-style-requested", types.YLeaf{"ProtectionStyleRequested", state.ProtectionStyleRequested})
    state.EntityData.Leafs.Append("reoptimize-timer", types.YLeaf{"ReoptimizeTimer", state.ReoptimizeTimer})
    state.EntityData.Leafs.Append("source", types.YLeaf{"Source", state.Source})
    state.EntityData.Leafs.Append("soft-preemption", types.YLeaf{"SoftPreemption", state.SoftPreemption})
    state.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", state.SetupPriority})
    state.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", state.HoldPriority})
    state.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", state.OperStatus})
    state.EntityData.Leafs.Append("role", types.YLeaf{"Role", state.Role})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_State_Counters
// State data for MPLS label switched paths. This state
// data is specific to a single label switched path.
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_State_Counters struct {
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
    // ^[0-9]{4}\-[0-9]{2}\-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}(\.[0-9]+)?Z[+-][0-9]{2}:[0-9]{2}$.
    OnlineTime interface{}

    // Indicates the time the LSP switched onto its current path. This is reset
    // upon a LSP path change. The type is string with pattern:
    // ^[0-9]{4}\-[0-9]{2}\-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}(\.[0-9]+)?Z[+-][0-9]{2}:[0-9]{2}$.
    CurrentPathTime interface{}

    // Indicates the next scheduled time the LSP will be reoptimized. The type is
    // string with pattern:
    // ^[0-9]{4}\-[0-9]{2}\-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}(\.[0-9]+)?Z[+-][0-9]{2}:[0-9]{2}$.
    NextReoptimizationTime interface{}
}

func (counters *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_State_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "openconfig"
    counters.EntityData.ParentYangName = "state"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    counters.EntityData.NamespaceTable = openconfig.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Leafs = types.NewOrderedMap()
    counters.EntityData.Leafs.Append("bytes", types.YLeaf{"Bytes", counters.Bytes})
    counters.EntityData.Leafs.Append("packets", types.YLeaf{"Packets", counters.Packets})
    counters.EntityData.Leafs.Append("path-changes", types.YLeaf{"PathChanges", counters.PathChanges})
    counters.EntityData.Leafs.Append("state-changes", types.YLeaf{"StateChanges", counters.StateChanges})
    counters.EntityData.Leafs.Append("online-time", types.YLeaf{"OnlineTime", counters.OnlineTime})
    counters.EntityData.Leafs.Append("current-path-time", types.YLeaf{"CurrentPathTime", counters.CurrentPathTime})
    counters.EntityData.Leafs.Append("next-reoptimization-time", types.YLeaf{"NextReoptimizationTime", counters.NextReoptimizationTime})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth
// Bandwidth configuration for TE LSPs
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters related to bandwidth on TE tunnels:.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_Config

    // State parameters related to bandwidth configuration of TE tunnels.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_State

    // Parameters related to auto-bandwidth.
    AutoBandwidth Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth
}

func (bandwidth *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "openconfig"
    bandwidth.EntityData.ParentYangName = "tunnel"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = openconfig.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    bandwidth.EntityData.Children = types.NewOrderedMap()
    bandwidth.EntityData.Children.Append("config", types.YChild{"Config", &bandwidth.Config})
    bandwidth.EntityData.Children.Append("state", types.YChild{"State", &bandwidth.State})
    bandwidth.EntityData.Children.Append("auto-bandwidth", types.YChild{"AutoBandwidth", &bandwidth.AutoBandwidth})
    bandwidth.EntityData.Leafs = types.NewOrderedMap()

    bandwidth.EntityData.YListKeys = []string {}

    return &(bandwidth.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_Config
// Configuration parameters related to bandwidth on TE
// tunnels:
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The method used for settign the bandwidth, either explicitly specified or
    // configured. The type is TeBandwidthType. The default value is SPECIFIED.
    SpecificationType interface{}

    // set bandwidth explicitly, e.g., using offline calculation. The type is
    // interface{} with range: 0..18446744073709551615.
    SetBandwidth interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "bandwidth"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("specification-type", types.YLeaf{"SpecificationType", config.SpecificationType})
    config.EntityData.Leafs.Append("set-bandwidth", types.YLeaf{"SetBandwidth", config.SetBandwidth})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_State
// State parameters related to bandwidth
// configuration of TE tunnels
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The method used for settign the bandwidth, either explicitly specified or
    // configured. The type is TeBandwidthType. The default value is SPECIFIED.
    SpecificationType interface{}

    // set bandwidth explicitly, e.g., using offline calculation. The type is
    // interface{} with range: 0..18446744073709551615.
    SetBandwidth interface{}

    // The currently signaled bandwidth of the LSP. In the case where the
    // bandwidth is specified explicitly, then this will match the value of the
    // set-bandwidth leaf; in cases where the bandwidth is dynamically computed by
    // the system, the current value of the bandwidth should be reflected. The
    // type is interface{} with range: 0..18446744073709551615.
    SignaledBandwidth interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "bandwidth"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("specification-type", types.YLeaf{"SpecificationType", state.SpecificationType})
    state.EntityData.Leafs.Append("set-bandwidth", types.YLeaf{"SetBandwidth", state.SetBandwidth})
    state.EntityData.Leafs.Append("signaled-bandwidth", types.YLeaf{"SignaledBandwidth", state.SignaledBandwidth})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth
// Parameters related to auto-bandwidth
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters relating to MPLS auto-bandwidth on the tunnel.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Config

    // State parameters relating to MPLS auto-bandwidth on the tunnel.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_State

    // configuration of MPLS overflow bandwidth adjustement for the LSP.
    Overflow Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow

    // configuration of MPLS underflow bandwidth adjustement for the LSP.
    Underflow Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow
}

func (autoBandwidth *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth) GetEntityData() *types.CommonEntityData {
    autoBandwidth.EntityData.YFilter = autoBandwidth.YFilter
    autoBandwidth.EntityData.YangName = "auto-bandwidth"
    autoBandwidth.EntityData.BundleName = "openconfig"
    autoBandwidth.EntityData.ParentYangName = "bandwidth"
    autoBandwidth.EntityData.SegmentPath = "auto-bandwidth"
    autoBandwidth.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    autoBandwidth.EntityData.NamespaceTable = openconfig.GetNamespaces()
    autoBandwidth.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    autoBandwidth.EntityData.Children = types.NewOrderedMap()
    autoBandwidth.EntityData.Children.Append("config", types.YChild{"Config", &autoBandwidth.Config})
    autoBandwidth.EntityData.Children.Append("state", types.YChild{"State", &autoBandwidth.State})
    autoBandwidth.EntityData.Children.Append("overflow", types.YChild{"Overflow", &autoBandwidth.Overflow})
    autoBandwidth.EntityData.Children.Append("underflow", types.YChild{"Underflow", &autoBandwidth.Underflow})
    autoBandwidth.EntityData.Leafs = types.NewOrderedMap()

    autoBandwidth.EntityData.YListKeys = []string {}

    return &(autoBandwidth.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Config
// Configuration parameters relating to MPLS
// auto-bandwidth on the tunnel.
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enables mpls auto-bandwidth on the lsp. The type is bool. The default value
    // is false.
    Enabled interface{}

    // set the minimum bandwidth in Kbps for an auto-bandwidth LSP. The type is
    // interface{} with range: 0..18446744073709551615.
    MinBw interface{}

    // set the maximum bandwidth in Kbps for an auto-bandwidth LSP. The type is
    // interface{} with range: 0..18446744073709551615.
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "auto-bandwidth"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("min-bw", types.YLeaf{"MinBw", config.MinBw})
    config.EntityData.Leafs.Append("max-bw", types.YLeaf{"MaxBw", config.MaxBw})
    config.EntityData.Leafs.Append("adjust-interval", types.YLeaf{"AdjustInterval", config.AdjustInterval})
    config.EntityData.Leafs.Append("adjust-threshold", types.YLeaf{"AdjustThreshold", config.AdjustThreshold})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_State
// State parameters relating to MPLS
// auto-bandwidth on the tunnel.
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enables mpls auto-bandwidth on the lsp. The type is bool. The default value
    // is false.
    Enabled interface{}

    // set the minimum bandwidth in Kbps for an auto-bandwidth LSP. The type is
    // interface{} with range: 0..18446744073709551615.
    MinBw interface{}

    // set the maximum bandwidth in Kbps for an auto-bandwidth LSP. The type is
    // interface{} with range: 0..18446744073709551615.
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "auto-bandwidth"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("min-bw", types.YLeaf{"MinBw", state.MinBw})
    state.EntityData.Leafs.Append("max-bw", types.YLeaf{"MaxBw", state.MaxBw})
    state.EntityData.Leafs.Append("adjust-interval", types.YLeaf{"AdjustInterval", state.AdjustInterval})
    state.EntityData.Leafs.Append("adjust-threshold", types.YLeaf{"AdjustThreshold", state.AdjustThreshold})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow
// configuration of MPLS overflow bandwidth
// adjustement for the LSP
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config information for MPLS overflow bandwidth adjustment.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config

    // Config information for MPLS overflow bandwidth adjustment.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow_State
}

func (overflow *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow) GetEntityData() *types.CommonEntityData {
    overflow.EntityData.YFilter = overflow.YFilter
    overflow.EntityData.YangName = "overflow"
    overflow.EntityData.BundleName = "openconfig"
    overflow.EntityData.ParentYangName = "auto-bandwidth"
    overflow.EntityData.SegmentPath = "overflow"
    overflow.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    overflow.EntityData.NamespaceTable = openconfig.GetNamespaces()
    overflow.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    overflow.EntityData.Children = types.NewOrderedMap()
    overflow.EntityData.Children.Append("config", types.YChild{"Config", &overflow.Config})
    overflow.EntityData.Children.Append("state", types.YChild{"State", &overflow.State})
    overflow.EntityData.Leafs = types.NewOrderedMap()

    overflow.EntityData.YListKeys = []string {}

    return &(overflow.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config
// Config information for MPLS overflow bandwidth
// adjustment
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config struct {
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "overflow"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("overflow-threshold", types.YLeaf{"OverflowThreshold", config.OverflowThreshold})
    config.EntityData.Leafs.Append("trigger-event-count", types.YLeaf{"TriggerEventCount", config.TriggerEventCount})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow_State
// Config information for MPLS overflow bandwidth
// adjustment
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow_State struct {
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Overflow_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "overflow"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("overflow-threshold", types.YLeaf{"OverflowThreshold", state.OverflowThreshold})
    state.EntityData.Leafs.Append("trigger-event-count", types.YLeaf{"TriggerEventCount", state.TriggerEventCount})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow
// configuration of MPLS underflow bandwidth
// adjustement for the LSP
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Config information for MPLS underflow bandwidth adjustment.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config

    // State information for MPLS underflow bandwidth adjustment.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow_State
}

func (underflow *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow) GetEntityData() *types.CommonEntityData {
    underflow.EntityData.YFilter = underflow.YFilter
    underflow.EntityData.YangName = "underflow"
    underflow.EntityData.BundleName = "openconfig"
    underflow.EntityData.ParentYangName = "auto-bandwidth"
    underflow.EntityData.SegmentPath = "underflow"
    underflow.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    underflow.EntityData.NamespaceTable = openconfig.GetNamespaces()
    underflow.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    underflow.EntityData.Children = types.NewOrderedMap()
    underflow.EntityData.Children.Append("config", types.YChild{"Config", &underflow.Config})
    underflow.EntityData.Children.Append("state", types.YChild{"State", &underflow.State})
    underflow.EntityData.Leafs = types.NewOrderedMap()

    underflow.EntityData.YListKeys = []string {}

    return &(underflow.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config
// Config information for MPLS underflow bandwidth
// adjustment
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config struct {
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "underflow"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("underflow-threshold", types.YLeaf{"UnderflowThreshold", config.UnderflowThreshold})
    config.EntityData.Leafs.Append("trigger-event-count", types.YLeaf{"TriggerEventCount", config.TriggerEventCount})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow_State
// State information for MPLS underflow bandwidth
// adjustment
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow_State struct {
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_Bandwidth_AutoBandwidth_Underflow_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "underflow"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("underflow-threshold", types.YLeaf{"UnderflowThreshold", state.UnderflowThreshold})
    state.EntityData.Leafs.Append("trigger-event-count", types.YLeaf{"TriggerEventCount", state.TriggerEventCount})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes
// Parameters related to LSPs of type P2P
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters for P2P LSPs.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_Config

    // State parameters for P2P LSPs.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_State

    // Primary paths associated with the LSP.
    P2pPrimaryPath Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath

    // Secondary paths for the LSP.
    P2pSecondaryPaths Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths
}

func (p2pTunnelAttributes *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes) GetEntityData() *types.CommonEntityData {
    p2pTunnelAttributes.EntityData.YFilter = p2pTunnelAttributes.YFilter
    p2pTunnelAttributes.EntityData.YangName = "p2p-tunnel-attributes"
    p2pTunnelAttributes.EntityData.BundleName = "openconfig"
    p2pTunnelAttributes.EntityData.ParentYangName = "tunnel"
    p2pTunnelAttributes.EntityData.SegmentPath = "p2p-tunnel-attributes"
    p2pTunnelAttributes.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2pTunnelAttributes.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2pTunnelAttributes.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2pTunnelAttributes.EntityData.Children = types.NewOrderedMap()
    p2pTunnelAttributes.EntityData.Children.Append("config", types.YChild{"Config", &p2pTunnelAttributes.Config})
    p2pTunnelAttributes.EntityData.Children.Append("state", types.YChild{"State", &p2pTunnelAttributes.State})
    p2pTunnelAttributes.EntityData.Children.Append("p2p-primary-path", types.YChild{"P2pPrimaryPath", &p2pTunnelAttributes.P2pPrimaryPath})
    p2pTunnelAttributes.EntityData.Children.Append("p2p-secondary-paths", types.YChild{"P2pSecondaryPaths", &p2pTunnelAttributes.P2pSecondaryPaths})
    p2pTunnelAttributes.EntityData.Leafs = types.NewOrderedMap()

    p2pTunnelAttributes.EntityData.YListKeys = []string {}

    return &(p2pTunnelAttributes.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_Config
// Configuration parameters for P2P LSPs
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // P2P tunnel destination address. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    Destination interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "p2p-tunnel-attributes"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", config.Destination})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_State
// State parameters for P2P LSPs
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // P2P tunnel destination address. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    Destination interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "p2p-tunnel-attributes"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", state.Destination})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath
// Primary paths associated with the LSP
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of p2p primary paths for a tunnel. The type is slice of
    // Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath.
    P2pPrimaryPath []*Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath
}

func (p2pPrimaryPath *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath) GetEntityData() *types.CommonEntityData {
    p2pPrimaryPath.EntityData.YFilter = p2pPrimaryPath.YFilter
    p2pPrimaryPath.EntityData.YangName = "p2p-primary-path"
    p2pPrimaryPath.EntityData.BundleName = "openconfig"
    p2pPrimaryPath.EntityData.ParentYangName = "p2p-tunnel-attributes"
    p2pPrimaryPath.EntityData.SegmentPath = "p2p-primary-path"
    p2pPrimaryPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2pPrimaryPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2pPrimaryPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2pPrimaryPath.EntityData.Children = types.NewOrderedMap()
    p2pPrimaryPath.EntityData.Children.Append("p2p-primary-path", types.YChild{"P2pPrimaryPath", nil})
    for i := range p2pPrimaryPath.P2pPrimaryPath {
        p2pPrimaryPath.EntityData.Children.Append(types.GetSegmentPath(p2pPrimaryPath.P2pPrimaryPath[i]), types.YChild{"P2pPrimaryPath", p2pPrimaryPath.P2pPrimaryPath[i]})
    }
    p2pPrimaryPath.EntityData.Leafs = types.NewOrderedMap()

    p2pPrimaryPath.EntityData.YListKeys = []string {}

    return &(p2pPrimaryPath.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath
// List of p2p primary paths for a tunnel
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path name. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_Config_Name
    Name interface{}

    // Configuration parameters related to paths.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_Config

    // State parameters related to paths.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_State

    // The set of candidate secondary paths which may be used for this primary
    // path. When secondary paths are specified in the list the path of the
    // secondary LSP in use must be restricted to those path options referenced.
    // The priority of the secondary paths is specified within the list. Higher
    // priority values are less preferred - that is to say that a path with
    // priority 0 is the most preferred path. In the case that the list is empty,
    // any secondary path option may be utilised when the current primary path is
    // in use.
    CandidateSecondaryPaths Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths

    // Top-level container for include/exclude constraints for link affinities.
    AdminGroups Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups
}

func (p2pPrimaryPath *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath) GetEntityData() *types.CommonEntityData {
    p2pPrimaryPath.EntityData.YFilter = p2pPrimaryPath.YFilter
    p2pPrimaryPath.EntityData.YangName = "p2p-primary-path"
    p2pPrimaryPath.EntityData.BundleName = "openconfig"
    p2pPrimaryPath.EntityData.ParentYangName = "p2p-primary-path"
    p2pPrimaryPath.EntityData.SegmentPath = "p2p-primary-path" + types.AddKeyToken(p2pPrimaryPath.Name, "name")
    p2pPrimaryPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2pPrimaryPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2pPrimaryPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2pPrimaryPath.EntityData.Children = types.NewOrderedMap()
    p2pPrimaryPath.EntityData.Children.Append("config", types.YChild{"Config", &p2pPrimaryPath.Config})
    p2pPrimaryPath.EntityData.Children.Append("state", types.YChild{"State", &p2pPrimaryPath.State})
    p2pPrimaryPath.EntityData.Children.Append("candidate-secondary-paths", types.YChild{"CandidateSecondaryPaths", &p2pPrimaryPath.CandidateSecondaryPaths})
    p2pPrimaryPath.EntityData.Children.Append("admin-groups", types.YChild{"AdminGroups", &p2pPrimaryPath.AdminGroups})
    p2pPrimaryPath.EntityData.Leafs = types.NewOrderedMap()
    p2pPrimaryPath.EntityData.Leafs.Append("name", types.YLeaf{"Name", p2pPrimaryPath.Name})

    p2pPrimaryPath.EntityData.YListKeys = []string {"Name"}

    return &(p2pPrimaryPath.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_Config
// Configuration parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LOCALLYCOMPUTEDEXTERNALLYQUERIEDEXPLICITLYDEFINED.
    // The default value is oc-mplst:LOCALLY_COMPUTED.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    PathComputationServer interface{}

    // reference to a defined path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_Name
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "p2p-primary-path"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("path-computation-method", types.YLeaf{"PathComputationMethod", config.PathComputationMethod})
    config.EntityData.Leafs.Append("use-cspf", types.YLeaf{"UseCspf", config.UseCspf})
    config.EntityData.Leafs.Append("cspf-tiebreaker", types.YLeaf{"CspfTiebreaker", config.CspfTiebreaker})
    config.EntityData.Leafs.Append("path-computation-server", types.YLeaf{"PathComputationServer", config.PathComputationServer})
    config.EntityData.Leafs.Append("explicit-path-name", types.YLeaf{"ExplicitPathName", config.ExplicitPathName})
    config.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", config.Preference})
    config.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", config.SetupPriority})
    config.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", config.HoldPriority})
    config.EntityData.Leafs.Append("retry-timer", types.YLeaf{"RetryTimer", config.RetryTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_State
// State parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LOCALLYCOMPUTEDEXTERNALLYQUERIEDEXPLICITLYDEFINED.
    // The default value is oc-mplst:LOCALLY_COMPUTED.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    PathComputationServer interface{}

    // reference to a defined path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_Name
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

    // If the signalling protocol specified for this path is RSVP-TE, this leaf
    // provides a reference to the associated session within the RSVP-TE protocol
    // sessions list, such that details of the signaling can be retrieved. The
    // type is string with range: 0..18446744073709551615. Refers to
    // mpls.Mpls_SignalingProtocols_RsvpTe_Sessions_Session_LocalIndex
    AssociatedRsvpSession interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "p2p-primary-path"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("path-computation-method", types.YLeaf{"PathComputationMethod", state.PathComputationMethod})
    state.EntityData.Leafs.Append("use-cspf", types.YLeaf{"UseCspf", state.UseCspf})
    state.EntityData.Leafs.Append("cspf-tiebreaker", types.YLeaf{"CspfTiebreaker", state.CspfTiebreaker})
    state.EntityData.Leafs.Append("path-computation-server", types.YLeaf{"PathComputationServer", state.PathComputationServer})
    state.EntityData.Leafs.Append("explicit-path-name", types.YLeaf{"ExplicitPathName", state.ExplicitPathName})
    state.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", state.Preference})
    state.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", state.SetupPriority})
    state.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", state.HoldPriority})
    state.EntityData.Leafs.Append("retry-timer", types.YLeaf{"RetryTimer", state.RetryTimer})
    state.EntityData.Leafs.Append("associated-rsvp-session", types.YLeaf{"AssociatedRsvpSession", state.AssociatedRsvpSession})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths
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
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of secondary paths which may be utilised when the current primary path
    // is in use. The type is slice of
    // Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath.
    CandidateSecondaryPath []*Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath
}

func (candidateSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths) GetEntityData() *types.CommonEntityData {
    candidateSecondaryPaths.EntityData.YFilter = candidateSecondaryPaths.YFilter
    candidateSecondaryPaths.EntityData.YangName = "candidate-secondary-paths"
    candidateSecondaryPaths.EntityData.BundleName = "openconfig"
    candidateSecondaryPaths.EntityData.ParentYangName = "p2p-primary-path"
    candidateSecondaryPaths.EntityData.SegmentPath = "candidate-secondary-paths"
    candidateSecondaryPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    candidateSecondaryPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    candidateSecondaryPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    candidateSecondaryPaths.EntityData.Children = types.NewOrderedMap()
    candidateSecondaryPaths.EntityData.Children.Append("candidate-secondary-path", types.YChild{"CandidateSecondaryPath", nil})
    for i := range candidateSecondaryPaths.CandidateSecondaryPath {
        candidateSecondaryPaths.EntityData.Children.Append(types.GetSegmentPath(candidateSecondaryPaths.CandidateSecondaryPath[i]), types.YChild{"CandidateSecondaryPath", candidateSecondaryPaths.CandidateSecondaryPath[i]})
    }
    candidateSecondaryPaths.EntityData.Leafs = types.NewOrderedMap()

    candidateSecondaryPaths.EntityData.YListKeys = []string {}

    return &(candidateSecondaryPaths.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath
// List of secondary paths which may be utilised when the
// current primary path is in use
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A reference to the secondary path option reference
    // which acts as the key of the candidate-secondary-path list. The type is
    // string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_Config_SecondaryPath
    SecondaryPath interface{}

    // Configuration parameters relating to the candidate secondary path.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_Config

    // Operational state parameters relating to the candidate secondary path.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_State
}

func (candidateSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath) GetEntityData() *types.CommonEntityData {
    candidateSecondaryPath.EntityData.YFilter = candidateSecondaryPath.YFilter
    candidateSecondaryPath.EntityData.YangName = "candidate-secondary-path"
    candidateSecondaryPath.EntityData.BundleName = "openconfig"
    candidateSecondaryPath.EntityData.ParentYangName = "candidate-secondary-paths"
    candidateSecondaryPath.EntityData.SegmentPath = "candidate-secondary-path" + types.AddKeyToken(candidateSecondaryPath.SecondaryPath, "secondary-path")
    candidateSecondaryPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    candidateSecondaryPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    candidateSecondaryPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    candidateSecondaryPath.EntityData.Children = types.NewOrderedMap()
    candidateSecondaryPath.EntityData.Children.Append("config", types.YChild{"Config", &candidateSecondaryPath.Config})
    candidateSecondaryPath.EntityData.Children.Append("state", types.YChild{"State", &candidateSecondaryPath.State})
    candidateSecondaryPath.EntityData.Leafs = types.NewOrderedMap()
    candidateSecondaryPath.EntityData.Leafs.Append("secondary-path", types.YLeaf{"SecondaryPath", candidateSecondaryPath.SecondaryPath})

    candidateSecondaryPath.EntityData.YListKeys = []string {"SecondaryPath"}

    return &(candidateSecondaryPath.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_Config
// Configuration parameters relating to the candidate
// secondary path
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A reference to the secondary path that should be utilised when the
    // containing primary path option is in use. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_Config_Name
    SecondaryPath interface{}

    // The priority of the specified secondary path option. Higher priority
    // options are less preferable - such that a secondary path reference with a
    // priority of 0 is the most preferred. The type is interface{} with range:
    // 0..65535.
    Priority interface{}
}

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "candidate-secondary-path"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("secondary-path", types.YLeaf{"SecondaryPath", config.SecondaryPath})
    config.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", config.Priority})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_State
// Operational state parameters relating to the candidate
// secondary path
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A reference to the secondary path that should be utilised when the
    // containing primary path option is in use. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_Config_Name
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_CandidateSecondaryPaths_CandidateSecondaryPath_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "candidate-secondary-path"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("secondary-path", types.YLeaf{"SecondaryPath", state.SecondaryPath})
    state.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", state.Priority})
    state.EntityData.Leafs.Append("active", types.YLeaf{"Active", state.Active})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups
// Top-level container for include/exclude constraints for
// link affinities
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data .
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups_Config

    // Operational state data .
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups_State
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups) GetEntityData() *types.CommonEntityData {
    adminGroups.EntityData.YFilter = adminGroups.YFilter
    adminGroups.EntityData.YangName = "admin-groups"
    adminGroups.EntityData.BundleName = "openconfig"
    adminGroups.EntityData.ParentYangName = "p2p-primary-path"
    adminGroups.EntityData.SegmentPath = "admin-groups"
    adminGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adminGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adminGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adminGroups.EntityData.Children = types.NewOrderedMap()
    adminGroups.EntityData.Children.Append("config", types.YChild{"Config", &adminGroups.Config})
    adminGroups.EntityData.Children.Append("state", types.YChild{"State", &adminGroups.State})
    adminGroups.EntityData.Leafs = types.NewOrderedMap()

    adminGroups.EntityData.YListKeys = []string {}

    return &(adminGroups.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups_Config
// Configuration data 
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups_Config struct {
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "admin-groups"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("exclude-group", types.YLeaf{"ExcludeGroup", config.ExcludeGroup})
    config.EntityData.Leafs.Append("include-all-group", types.YLeaf{"IncludeAllGroup", config.IncludeAllGroup})
    config.EntityData.Leafs.Append("include-any-group", types.YLeaf{"IncludeAnyGroup", config.IncludeAnyGroup})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups_State
// Operational state data 
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups_State struct {
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pPrimaryPath_P2pPrimaryPath_AdminGroups_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "admin-groups"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("exclude-group", types.YLeaf{"ExcludeGroup", state.ExcludeGroup})
    state.EntityData.Leafs.Append("include-all-group", types.YLeaf{"IncludeAllGroup", state.IncludeAllGroup})
    state.EntityData.Leafs.Append("include-any-group", types.YLeaf{"IncludeAnyGroup", state.IncludeAnyGroup})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths
// Secondary paths for the LSP
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of p2p primary paths for a tunnel. The type is slice of
    // Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath.
    P2pSecondaryPath []*Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath
}

func (p2pSecondaryPaths *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths) GetEntityData() *types.CommonEntityData {
    p2pSecondaryPaths.EntityData.YFilter = p2pSecondaryPaths.YFilter
    p2pSecondaryPaths.EntityData.YangName = "p2p-secondary-paths"
    p2pSecondaryPaths.EntityData.BundleName = "openconfig"
    p2pSecondaryPaths.EntityData.ParentYangName = "p2p-tunnel-attributes"
    p2pSecondaryPaths.EntityData.SegmentPath = "p2p-secondary-paths"
    p2pSecondaryPaths.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2pSecondaryPaths.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2pSecondaryPaths.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2pSecondaryPaths.EntityData.Children = types.NewOrderedMap()
    p2pSecondaryPaths.EntityData.Children.Append("p2p-secondary-path", types.YChild{"P2pSecondaryPath", nil})
    for i := range p2pSecondaryPaths.P2pSecondaryPath {
        p2pSecondaryPaths.EntityData.Children.Append(types.GetSegmentPath(p2pSecondaryPaths.P2pSecondaryPath[i]), types.YChild{"P2pSecondaryPath", p2pSecondaryPaths.P2pSecondaryPath[i]})
    }
    p2pSecondaryPaths.EntityData.Leafs = types.NewOrderedMap()

    p2pSecondaryPaths.EntityData.YListKeys = []string {}

    return &(p2pSecondaryPaths.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath
// List of p2p primary paths for a tunnel
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path name. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_Config_Name
    Name interface{}

    // Configuration parameters related to paths.
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_Config

    // State parameters related to paths.
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_State

    // Top-level container for include/exclude constraints for link affinities.
    AdminGroups Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups
}

func (p2pSecondaryPath *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath) GetEntityData() *types.CommonEntityData {
    p2pSecondaryPath.EntityData.YFilter = p2pSecondaryPath.YFilter
    p2pSecondaryPath.EntityData.YangName = "p2p-secondary-path"
    p2pSecondaryPath.EntityData.BundleName = "openconfig"
    p2pSecondaryPath.EntityData.ParentYangName = "p2p-secondary-paths"
    p2pSecondaryPath.EntityData.SegmentPath = "p2p-secondary-path" + types.AddKeyToken(p2pSecondaryPath.Name, "name")
    p2pSecondaryPath.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    p2pSecondaryPath.EntityData.NamespaceTable = openconfig.GetNamespaces()
    p2pSecondaryPath.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    p2pSecondaryPath.EntityData.Children = types.NewOrderedMap()
    p2pSecondaryPath.EntityData.Children.Append("config", types.YChild{"Config", &p2pSecondaryPath.Config})
    p2pSecondaryPath.EntityData.Children.Append("state", types.YChild{"State", &p2pSecondaryPath.State})
    p2pSecondaryPath.EntityData.Children.Append("admin-groups", types.YChild{"AdminGroups", &p2pSecondaryPath.AdminGroups})
    p2pSecondaryPath.EntityData.Leafs = types.NewOrderedMap()
    p2pSecondaryPath.EntityData.Leafs.Append("name", types.YLeaf{"Name", p2pSecondaryPath.Name})

    p2pSecondaryPath.EntityData.YListKeys = []string {"Name"}

    return &(p2pSecondaryPath.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_Config
// Configuration parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LOCALLYCOMPUTEDEXTERNALLYQUERIEDEXPLICITLYDEFINED.
    // The default value is oc-mplst:LOCALLY_COMPUTED.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    PathComputationServer interface{}

    // reference to a defined path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_Name
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "p2p-secondary-path"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("path-computation-method", types.YLeaf{"PathComputationMethod", config.PathComputationMethod})
    config.EntityData.Leafs.Append("use-cspf", types.YLeaf{"UseCspf", config.UseCspf})
    config.EntityData.Leafs.Append("cspf-tiebreaker", types.YLeaf{"CspfTiebreaker", config.CspfTiebreaker})
    config.EntityData.Leafs.Append("path-computation-server", types.YLeaf{"PathComputationServer", config.PathComputationServer})
    config.EntityData.Leafs.Append("explicit-path-name", types.YLeaf{"ExplicitPathName", config.ExplicitPathName})
    config.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", config.Preference})
    config.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", config.SetupPriority})
    config.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", config.HoldPriority})
    config.EntityData.Leafs.Append("retry-timer", types.YLeaf{"RetryTimer", config.RetryTimer})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_State
// State parameters related to paths
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path name. The type is string.
    Name interface{}

    // The method used for computing the path, either locally computed, queried
    // from a server or not computed at all (explicitly configured). The type is
    // one of the following: LOCALLYCOMPUTEDEXTERNALLYQUERIEDEXPLICITLYDEFINED.
    // The default value is oc-mplst:LOCALLY_COMPUTED.
    PathComputationMethod interface{}

    // Flag to enable CSPF for locally computed LSPs. The type is bool.
    UseCspf interface{}

    // Determine the tie-breaking method to choose between equally desirable paths
    // during CSFP computation. The type is CspfTieBreaking.
    CspfTiebreaker interface{}

    // Address of the external path computation server. The type is one of the
    // following types: string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    PathComputationServer interface{}

    // reference to a defined path. The type is string. Refers to
    // mpls.Mpls_Lsps_ConstrainedPath_NamedExplicitPaths_NamedExplicitPath_Config_Name
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

    // If the signalling protocol specified for this path is RSVP-TE, this leaf
    // provides a reference to the associated session within the RSVP-TE protocol
    // sessions list, such that details of the signaling can be retrieved. The
    // type is string with range: 0..18446744073709551615. Refers to
    // mpls.Mpls_SignalingProtocols_RsvpTe_Sessions_Session_LocalIndex
    AssociatedRsvpSession interface{}
}

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "p2p-secondary-path"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("path-computation-method", types.YLeaf{"PathComputationMethod", state.PathComputationMethod})
    state.EntityData.Leafs.Append("use-cspf", types.YLeaf{"UseCspf", state.UseCspf})
    state.EntityData.Leafs.Append("cspf-tiebreaker", types.YLeaf{"CspfTiebreaker", state.CspfTiebreaker})
    state.EntityData.Leafs.Append("path-computation-server", types.YLeaf{"PathComputationServer", state.PathComputationServer})
    state.EntityData.Leafs.Append("explicit-path-name", types.YLeaf{"ExplicitPathName", state.ExplicitPathName})
    state.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", state.Preference})
    state.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", state.SetupPriority})
    state.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", state.HoldPriority})
    state.EntityData.Leafs.Append("retry-timer", types.YLeaf{"RetryTimer", state.RetryTimer})
    state.EntityData.Leafs.Append("associated-rsvp-session", types.YLeaf{"AssociatedRsvpSession", state.AssociatedRsvpSession})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups
// Top-level container for include/exclude constraints for
// link affinities
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data .
    Config Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups_Config

    // Operational state data .
    State Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups_State
}

func (adminGroups *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups) GetEntityData() *types.CommonEntityData {
    adminGroups.EntityData.YFilter = adminGroups.YFilter
    adminGroups.EntityData.YangName = "admin-groups"
    adminGroups.EntityData.BundleName = "openconfig"
    adminGroups.EntityData.ParentYangName = "p2p-secondary-path"
    adminGroups.EntityData.SegmentPath = "admin-groups"
    adminGroups.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    adminGroups.EntityData.NamespaceTable = openconfig.GetNamespaces()
    adminGroups.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    adminGroups.EntityData.Children = types.NewOrderedMap()
    adminGroups.EntityData.Children.Append("config", types.YChild{"Config", &adminGroups.Config})
    adminGroups.EntityData.Children.Append("state", types.YChild{"State", &adminGroups.State})
    adminGroups.EntityData.Leafs = types.NewOrderedMap()

    adminGroups.EntityData.YListKeys = []string {}

    return &(adminGroups.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups_Config
// Configuration data 
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups_Config struct {
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

func (config *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "admin-groups"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("exclude-group", types.YLeaf{"ExcludeGroup", config.ExcludeGroup})
    config.EntityData.Leafs.Append("include-all-group", types.YLeaf{"IncludeAllGroup", config.IncludeAllGroup})
    config.EntityData.Leafs.Append("include-any-group", types.YLeaf{"IncludeAnyGroup", config.IncludeAnyGroup})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups_State
// Operational state data 
type Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups_State struct {
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

func (state *Mpls_Lsps_ConstrainedPath_Tunnels_Tunnel_P2pTunnelAttributes_P2pSecondaryPaths_P2pSecondaryPath_AdminGroups_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "admin-groups"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("exclude-group", types.YLeaf{"ExcludeGroup", state.ExcludeGroup})
    state.EntityData.Leafs.Append("include-all-group", types.YLeaf{"IncludeAllGroup", state.IncludeAllGroup})
    state.EntityData.Leafs.Append("include-any-group", types.YLeaf{"IncludeAnyGroup", state.IncludeAnyGroup})

    state.EntityData.YListKeys = []string {}

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

    unconstrainedPath.EntityData.Children = types.NewOrderedMap()
    unconstrainedPath.EntityData.Children.Append("path-setup-protocol", types.YChild{"PathSetupProtocol", &unconstrainedPath.PathSetupProtocol})
    unconstrainedPath.EntityData.Leafs = types.NewOrderedMap()

    unconstrainedPath.EntityData.YListKeys = []string {}

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

    pathSetupProtocol.EntityData.Children = types.NewOrderedMap()
    pathSetupProtocol.EntityData.Children.Append("ldp", types.YChild{"Ldp", &pathSetupProtocol.Ldp})
    pathSetupProtocol.EntityData.Leafs = types.NewOrderedMap()

    pathSetupProtocol.EntityData.YListKeys = []string {}

    return &(pathSetupProtocol.EntityData)
}

// Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp
// LDP signaling setup for IGP-congruent LSPs
type Mpls_Lsps_UnconstrainedPath_PathSetupProtocol_Ldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
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

    ldp.EntityData.Children = types.NewOrderedMap()
    ldp.EntityData.Leafs = types.NewOrderedMap()

    ldp.EntityData.YListKeys = []string {}

    return &(ldp.EntityData)
}

// Mpls_Lsps_StaticLsps
// statically configured LSPs, without dynamic
// signaling
type Mpls_Lsps_StaticLsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // list of defined static LSPs. The type is slice of
    // Mpls_Lsps_StaticLsps_StaticLsp.
    StaticLsp []*Mpls_Lsps_StaticLsps_StaticLsp
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

    staticLsps.EntityData.Children = types.NewOrderedMap()
    staticLsps.EntityData.Children.Append("static-lsp", types.YChild{"StaticLsp", nil})
    for i := range staticLsps.StaticLsp {
        staticLsps.EntityData.Children.Append(types.GetSegmentPath(staticLsps.StaticLsp[i]), types.YChild{"StaticLsp", staticLsps.StaticLsp[i]})
    }
    staticLsps.EntityData.Leafs = types.NewOrderedMap()

    staticLsps.EntityData.YListKeys = []string {}

    return &(staticLsps.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp
// list of defined static LSPs
type Mpls_Lsps_StaticLsps_StaticLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference the name list key. The type is string.
    // Refers to mpls.Mpls_Lsps_StaticLsps_StaticLsp_Config_Name
    Name interface{}

    // Configuration data for the static lsp.
    Config Mpls_Lsps_StaticLsps_StaticLsp_Config

    // Operational state data for the static lsp.
    State Mpls_Lsps_StaticLsps_StaticLsp_State

    // Static LSPs for which the router is an  ingress node.
    Ingress Mpls_Lsps_StaticLsps_StaticLsp_Ingress

    // Static LSPs for which the router is an  transit node.
    Transit Mpls_Lsps_StaticLsps_StaticLsp_Transit

    // Static LSPs for which the router is an  egress node.
    Egress Mpls_Lsps_StaticLsps_StaticLsp_Egress
}

func (staticLsp *Mpls_Lsps_StaticLsps_StaticLsp) GetEntityData() *types.CommonEntityData {
    staticLsp.EntityData.YFilter = staticLsp.YFilter
    staticLsp.EntityData.YangName = "static-lsp"
    staticLsp.EntityData.BundleName = "openconfig"
    staticLsp.EntityData.ParentYangName = "static-lsps"
    staticLsp.EntityData.SegmentPath = "static-lsp" + types.AddKeyToken(staticLsp.Name, "name")
    staticLsp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    staticLsp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    staticLsp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    staticLsp.EntityData.Children = types.NewOrderedMap()
    staticLsp.EntityData.Children.Append("config", types.YChild{"Config", &staticLsp.Config})
    staticLsp.EntityData.Children.Append("state", types.YChild{"State", &staticLsp.State})
    staticLsp.EntityData.Children.Append("ingress", types.YChild{"Ingress", &staticLsp.Ingress})
    staticLsp.EntityData.Children.Append("transit", types.YChild{"Transit", &staticLsp.Transit})
    staticLsp.EntityData.Children.Append("egress", types.YChild{"Egress", &staticLsp.Egress})
    staticLsp.EntityData.Leafs = types.NewOrderedMap()
    staticLsp.EntityData.Leafs.Append("name", types.YLeaf{"Name", staticLsp.Name})

    staticLsp.EntityData.YListKeys = []string {"Name"}

    return &(staticLsp.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Config
// Configuration data for the static lsp
type Mpls_Lsps_StaticLsps_StaticLsp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name to identify the LSP. The type is string.
    Name interface{}
}

func (config *Mpls_Lsps_StaticLsps_StaticLsp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "static-lsp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_State
// Operational state data for the static lsp
type Mpls_Lsps_StaticLsps_StaticLsp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name to identify the LSP. The type is string.
    Name interface{}
}

func (state *Mpls_Lsps_StaticLsps_StaticLsp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "static-lsp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Ingress
// Static LSPs for which the router is an
//  ingress node
type Mpls_Lsps_StaticLsps_StaticLsp_Ingress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for ingress LSPs.
    Config Mpls_Lsps_StaticLsps_StaticLsp_Ingress_Config

    // Operational state data for ingress LSPs.
    State Mpls_Lsps_StaticLsps_StaticLsp_Ingress_State
}

func (ingress *Mpls_Lsps_StaticLsps_StaticLsp_Ingress) GetEntityData() *types.CommonEntityData {
    ingress.EntityData.YFilter = ingress.YFilter
    ingress.EntityData.YangName = "ingress"
    ingress.EntityData.BundleName = "openconfig"
    ingress.EntityData.ParentYangName = "static-lsp"
    ingress.EntityData.SegmentPath = "ingress"
    ingress.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ingress.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ingress.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ingress.EntityData.Children = types.NewOrderedMap()
    ingress.EntityData.Children.Append("config", types.YChild{"Config", &ingress.Config})
    ingress.EntityData.Children.Append("state", types.YChild{"State", &ingress.State})
    ingress.EntityData.Leafs = types.NewOrderedMap()

    ingress.EntityData.YListKeys = []string {}

    return &(ingress.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Ingress_Config
// Configuration data for ingress LSPs
type Mpls_Lsps_StaticLsps_StaticLsp_Ingress_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (config *Mpls_Lsps_StaticLsps_StaticLsp_Ingress_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ingress"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", config.NextHop})
    config.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", config.IncomingLabel})
    config.EntityData.Leafs.Append("push-label", types.YLeaf{"PushLabel", config.PushLabel})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Ingress_State
// Operational state data for ingress LSPs
type Mpls_Lsps_StaticLsps_StaticLsp_Ingress_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (state *Mpls_Lsps_StaticLsps_StaticLsp_Ingress_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ingress"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", state.NextHop})
    state.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", state.IncomingLabel})
    state.EntityData.Leafs.Append("push-label", types.YLeaf{"PushLabel", state.PushLabel})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Transit
// Static LSPs for which the router is an
//  transit node
type Mpls_Lsps_StaticLsps_StaticLsp_Transit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for transit LSPs.
    Config Mpls_Lsps_StaticLsps_StaticLsp_Transit_Config

    // Operational state data for transit LSPs.
    State Mpls_Lsps_StaticLsps_StaticLsp_Transit_State
}

func (transit *Mpls_Lsps_StaticLsps_StaticLsp_Transit) GetEntityData() *types.CommonEntityData {
    transit.EntityData.YFilter = transit.YFilter
    transit.EntityData.YangName = "transit"
    transit.EntityData.BundleName = "openconfig"
    transit.EntityData.ParentYangName = "static-lsp"
    transit.EntityData.SegmentPath = "transit"
    transit.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    transit.EntityData.NamespaceTable = openconfig.GetNamespaces()
    transit.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    transit.EntityData.Children = types.NewOrderedMap()
    transit.EntityData.Children.Append("config", types.YChild{"Config", &transit.Config})
    transit.EntityData.Children.Append("state", types.YChild{"State", &transit.State})
    transit.EntityData.Leafs = types.NewOrderedMap()

    transit.EntityData.YListKeys = []string {}

    return &(transit.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Transit_Config
// Configuration data for transit LSPs
type Mpls_Lsps_StaticLsps_StaticLsp_Transit_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (config *Mpls_Lsps_StaticLsps_StaticLsp_Transit_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "transit"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", config.NextHop})
    config.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", config.IncomingLabel})
    config.EntityData.Leafs.Append("push-label", types.YLeaf{"PushLabel", config.PushLabel})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Transit_State
// Operational state data for transit LSPs
type Mpls_Lsps_StaticLsps_StaticLsp_Transit_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (state *Mpls_Lsps_StaticLsps_StaticLsp_Transit_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "transit"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", state.NextHop})
    state.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", state.IncomingLabel})
    state.EntityData.Leafs.Append("push-label", types.YLeaf{"PushLabel", state.PushLabel})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Egress
// Static LSPs for which the router is an
//  egress node
type Mpls_Lsps_StaticLsps_StaticLsp_Egress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for egress LSPs.
    Config Mpls_Lsps_StaticLsps_StaticLsp_Egress_Config

    // Operational state data for egress LSPs.
    State Mpls_Lsps_StaticLsps_StaticLsp_Egress_State
}

func (egress *Mpls_Lsps_StaticLsps_StaticLsp_Egress) GetEntityData() *types.CommonEntityData {
    egress.EntityData.YFilter = egress.YFilter
    egress.EntityData.YangName = "egress"
    egress.EntityData.BundleName = "openconfig"
    egress.EntityData.ParentYangName = "static-lsp"
    egress.EntityData.SegmentPath = "egress"
    egress.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    egress.EntityData.NamespaceTable = openconfig.GetNamespaces()
    egress.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    egress.EntityData.Children = types.NewOrderedMap()
    egress.EntityData.Children.Append("config", types.YChild{"Config", &egress.Config})
    egress.EntityData.Children.Append("state", types.YChild{"State", &egress.State})
    egress.EntityData.Leafs = types.NewOrderedMap()

    egress.EntityData.YListKeys = []string {}

    return &(egress.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Egress_Config
// Configuration data for egress LSPs
type Mpls_Lsps_StaticLsps_StaticLsp_Egress_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (config *Mpls_Lsps_StaticLsps_StaticLsp_Egress_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "egress"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", config.NextHop})
    config.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", config.IncomingLabel})
    config.EntityData.Leafs.Append("push-label", types.YLeaf{"PushLabel", config.PushLabel})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mpls_Lsps_StaticLsps_StaticLsp_Egress_State
// Operational state data for egress LSPs
type Mpls_Lsps_StaticLsps_StaticLsp_Egress_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // next hop IP address for the LSP. The type is one of the following types:
    // string with pattern:
    // ^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$,
    // or string with pattern:
    // ^(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))$.
    NextHop interface{}

    // label value on the incoming packet. The type is one of the following types:
    // int with range: 16..1048575, or enumeration MplsLabel.
    IncomingLabel interface{}

    // label value to push at the current hop for the LSP. The type is one of the
    // following types: int with range: 16..1048575, or enumeration MplsLabel.
    PushLabel interface{}
}

func (state *Mpls_Lsps_StaticLsps_StaticLsp_Egress_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "egress"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", state.NextHop})
    state.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", state.IncomingLabel})
    state.EntityData.Leafs.Append("push-label", types.YLeaf{"PushLabel", state.PushLabel})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

