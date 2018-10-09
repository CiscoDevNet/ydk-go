// This module defines configuration and operational state data
// for the LLDP protocol.
package lldp

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lldp"))
    ydk.RegisterEntity("{http://openconfig.net/yang/lldp lldp}", reflect.TypeOf(Lldp{}))
    ydk.RegisterEntity("openconfig-lldp:lldp", reflect.TypeOf(Lldp{}))
}

// Lldp
// Top-level container for LLDP configuration and state data
type Lldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data .
    Config Lldp_Config

    // Operational state data .
    State Lldp_State

    // Enclosing container .
    Interfaces Lldp_Interfaces
}

func (lldp *Lldp) GetEntityData() *types.CommonEntityData {
    lldp.EntityData.YFilter = lldp.YFilter
    lldp.EntityData.YangName = "lldp"
    lldp.EntityData.BundleName = "openconfig"
    lldp.EntityData.ParentYangName = "openconfig-lldp"
    lldp.EntityData.SegmentPath = "openconfig-lldp:lldp"
    lldp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    lldp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    lldp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    lldp.EntityData.Children = types.NewOrderedMap()
    lldp.EntityData.Children.Append("config", types.YChild{"Config", &lldp.Config})
    lldp.EntityData.Children.Append("state", types.YChild{"State", &lldp.State})
    lldp.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &lldp.Interfaces})
    lldp.EntityData.Leafs = types.NewOrderedMap()

    lldp.EntityData.YListKeys = []string {}

    return &(lldp.EntityData)
}

// Lldp_Config
// Configuration data 
type Lldp_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // System level state of the LLDP protocol. The type is bool. The default
    // value is true.
    Enabled interface{}

    // System level hello timer for the LLDP protocol. The type is interface{}
    // with range: 0..18446744073709551615. Units are seconds.
    HelloTimer interface{}

    // Indicates whether the local system should suppress the advertisement of
    // particular TLVs with the LLDP PDUs that it transmits. Where a TLV type is
    // specified within this list, it should not be included in any LLDP PDU
    // transmitted by the local agent. The type is slice of ['SYSTEMNAME',
    // 'PORTDESCRIPTION', 'SYSTEMDESCRIPTION', 'CHASSISID', 'SYSTEMCAPABILITIES',
    // 'MANAGEMENTADDRESS', 'PORTID'].
    SuppressTlvAdvertisement []interface{}

    // The system name field shall contain an alpha-numeric string that indicates
    // the system's administratively assigned name. The system name should be the
    // system's fully qualified domain name. If implementations support IETF RFC
    // 3418, the sysName object should be used for this field. The type is string
    // with length: 0..255.
    SystemName interface{}

    // The system description field shall contain an alpha-numeric string that is
    // the textual description of the network entity. The system description
    // should include the full name and version identification of the system's
    // hardware type, software operating system, and networking software. If
    // implementations support IETF RFC 3418, the sysDescr object should be used
    // for this field. The type is string with length: 0..255.
    SystemDescription interface{}

    // The Chassis ID is a mandatory TLV which identifies the chassis component of
    // the endpoint identifier associated with the transmitting LLDP agent. The
    // type is string.
    ChassisId interface{}

    // This field identifies the format and source of the chassis identifier
    // string. It is an enumerator defined by the LldpChassisIdSubtype object from
    // IEEE 802.1AB MIB. The type is ChassisIdType.
    ChassisIdType interface{}
}

func (config *Lldp_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "lldp"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("hello-timer", types.YLeaf{"HelloTimer", config.HelloTimer})
    config.EntityData.Leafs.Append("suppress-tlv-advertisement", types.YLeaf{"SuppressTlvAdvertisement", config.SuppressTlvAdvertisement})
    config.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", config.SystemName})
    config.EntityData.Leafs.Append("system-description", types.YLeaf{"SystemDescription", config.SystemDescription})
    config.EntityData.Leafs.Append("chassis-id", types.YLeaf{"ChassisId", config.ChassisId})
    config.EntityData.Leafs.Append("chassis-id-type", types.YLeaf{"ChassisIdType", config.ChassisIdType})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Lldp_State
// Operational state data 
type Lldp_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // System level state of the LLDP protocol. The type is bool. The default
    // value is true.
    Enabled interface{}

    // System level hello timer for the LLDP protocol. The type is interface{}
    // with range: 0..18446744073709551615. Units are seconds.
    HelloTimer interface{}

    // Indicates whether the local system should suppress the advertisement of
    // particular TLVs with the LLDP PDUs that it transmits. Where a TLV type is
    // specified within this list, it should not be included in any LLDP PDU
    // transmitted by the local agent. The type is slice of ['SYSTEMNAME',
    // 'PORTDESCRIPTION', 'SYSTEMDESCRIPTION', 'CHASSISID', 'SYSTEMCAPABILITIES',
    // 'MANAGEMENTADDRESS', 'PORTID'].
    SuppressTlvAdvertisement []interface{}

    // The system name field shall contain an alpha-numeric string that indicates
    // the system's administratively assigned name. The system name should be the
    // system's fully qualified domain name. If implementations support IETF RFC
    // 3418, the sysName object should be used for this field. The type is string
    // with length: 0..255.
    SystemName interface{}

    // The system description field shall contain an alpha-numeric string that is
    // the textual description of the network entity. The system description
    // should include the full name and version identification of the system's
    // hardware type, software operating system, and networking software. If
    // implementations support IETF RFC 3418, the sysDescr object should be used
    // for this field. The type is string with length: 0..255.
    SystemDescription interface{}

    // The Chassis ID is a mandatory TLV which identifies the chassis component of
    // the endpoint identifier associated with the transmitting LLDP agent. The
    // type is string.
    ChassisId interface{}

    // This field identifies the format and source of the chassis identifier
    // string. It is an enumerator defined by the LldpChassisIdSubtype object from
    // IEEE 802.1AB MIB. The type is ChassisIdType.
    ChassisIdType interface{}

    // Global LLDP counters.
    Counters Lldp_State_Counters
}

func (state *Lldp_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "lldp"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("counters", types.YChild{"Counters", &state.Counters})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("hello-timer", types.YLeaf{"HelloTimer", state.HelloTimer})
    state.EntityData.Leafs.Append("suppress-tlv-advertisement", types.YLeaf{"SuppressTlvAdvertisement", state.SuppressTlvAdvertisement})
    state.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", state.SystemName})
    state.EntityData.Leafs.Append("system-description", types.YLeaf{"SystemDescription", state.SystemDescription})
    state.EntityData.Leafs.Append("chassis-id", types.YLeaf{"ChassisId", state.ChassisId})
    state.EntityData.Leafs.Append("chassis-id-type", types.YLeaf{"ChassisIdType", state.ChassisIdType})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Lldp_State_Counters
// Global LLDP counters
type Lldp_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of lldp frames received. The type is interface{} with range:
    // 0..18446744073709551615.
    FrameIn interface{}

    // The number of frames transmitted out. The type is interface{} with range:
    // 0..18446744073709551615.
    FrameOut interface{}

    // The number of LLDP frames received with errors. The type is interface{}
    // with range: 0..18446744073709551615.
    FrameErrorIn interface{}

    // The number of LLDP frames received and discarded. The type is interface{}
    // with range: 0..18446744073709551615.
    FrameDiscard interface{}

    // The number of TLV frames received and discarded. The type is interface{}
    // with range: 0..18446744073709551615.
    TlvDiscard interface{}

    // The number of frames received with unknown TLV. The type is interface{}
    // with range: 0..18446744073709551615.
    TlvUnknown interface{}

    // Indicates the last time the counters were cleared. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastClear interface{}

    // The number of valid TLVs received. The type is interface{} with range:
    // 0..18446744073709551615.
    TlvAccepted interface{}

    // The number of entries aged out due to timeout. The type is interface{} with
    // range: 0..18446744073709551615.
    EntriesAgedOut interface{}
}

func (counters *Lldp_State_Counters) GetEntityData() *types.CommonEntityData {
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
    counters.EntityData.Leafs.Append("frame-in", types.YLeaf{"FrameIn", counters.FrameIn})
    counters.EntityData.Leafs.Append("frame-out", types.YLeaf{"FrameOut", counters.FrameOut})
    counters.EntityData.Leafs.Append("frame-error-in", types.YLeaf{"FrameErrorIn", counters.FrameErrorIn})
    counters.EntityData.Leafs.Append("frame-discard", types.YLeaf{"FrameDiscard", counters.FrameDiscard})
    counters.EntityData.Leafs.Append("tlv-discard", types.YLeaf{"TlvDiscard", counters.TlvDiscard})
    counters.EntityData.Leafs.Append("tlv-unknown", types.YLeaf{"TlvUnknown", counters.TlvUnknown})
    counters.EntityData.Leafs.Append("last-clear", types.YLeaf{"LastClear", counters.LastClear})
    counters.EntityData.Leafs.Append("tlv-accepted", types.YLeaf{"TlvAccepted", counters.TlvAccepted})
    counters.EntityData.Leafs.Append("entries-aged-out", types.YLeaf{"EntriesAgedOut", counters.EntriesAgedOut})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Lldp_Interfaces
// Enclosing container 
type Lldp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of interfaces on which LLDP is enabled / available. The type is slice
    // of Lldp_Interfaces_Interface.
    Interface []*Lldp_Interfaces_Interface
}

func (interfaces *Lldp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "openconfig"
    interfaces.EntityData.ParentYangName = "lldp"
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

// Lldp_Interfaces_Interface
// List of interfaces on which LLDP is enabled / available
type Lldp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to the list key. The type is string.
    // Refers to lldp.Lldp_Interfaces_Interface_Config_Name
    Name interface{}

    // Configuration data for LLDP on each interface.
    Config Lldp_Interfaces_Interface_Config

    // Operational state data .
    State Lldp_Interfaces_Interface_State

    // Enclosing container for list of LLDP neighbors on an interface.
    Neighbors Lldp_Interfaces_Interface_Neighbors
}

func (self *Lldp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "openconfig"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Name, "name")
    self.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    self.EntityData.NamespaceTable = openconfig.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("config", types.YChild{"Config", &self.Config})
    self.EntityData.Children.Append("state", types.YChild{"State", &self.State})
    self.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &self.Neighbors})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {"Name"}

    return &(self.EntityData)
}

// Lldp_Interfaces_Interface_Config
// Configuration data for LLDP on each interface
type Lldp_Interfaces_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the LLDP Ethernet interface. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // Enable or disable the LLDP protocol on the interface. The type is bool. The
    // default value is true.
    Enabled interface{}
}

func (config *Lldp_Interfaces_Interface_Config) GetEntityData() *types.CommonEntityData {
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
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Lldp_Interfaces_Interface_State
// Operational state data 
type Lldp_Interfaces_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to the LLDP Ethernet interface. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    Name interface{}

    // Enable or disable the LLDP protocol on the interface. The type is bool. The
    // default value is true.
    Enabled interface{}

    // LLDP counters on each interface.
    Counters Lldp_Interfaces_Interface_State_Counters
}

func (state *Lldp_Interfaces_Interface_State) GetEntityData() *types.CommonEntityData {
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
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Lldp_Interfaces_Interface_State_Counters
// LLDP counters on each interface
type Lldp_Interfaces_Interface_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of lldp frames received. The type is interface{} with range:
    // 0..18446744073709551615.
    FrameIn interface{}

    // The number of frames transmitted out. The type is interface{} with range:
    // 0..18446744073709551615.
    FrameOut interface{}

    // The number of LLDP frames received with errors. The type is interface{}
    // with range: 0..18446744073709551615.
    FrameErrorIn interface{}

    // The number of LLDP frames received and discarded. The type is interface{}
    // with range: 0..18446744073709551615.
    FrameDiscard interface{}

    // The number of TLV frames received and discarded. The type is interface{}
    // with range: 0..18446744073709551615.
    TlvDiscard interface{}

    // The number of frames received with unknown TLV. The type is interface{}
    // with range: 0..18446744073709551615.
    TlvUnknown interface{}

    // Indicates the last time the counters were cleared. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastClear interface{}

    // The number of frame transmit errors on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    FrameErrorOut interface{}
}

func (counters *Lldp_Interfaces_Interface_State_Counters) GetEntityData() *types.CommonEntityData {
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
    counters.EntityData.Leafs.Append("frame-in", types.YLeaf{"FrameIn", counters.FrameIn})
    counters.EntityData.Leafs.Append("frame-out", types.YLeaf{"FrameOut", counters.FrameOut})
    counters.EntityData.Leafs.Append("frame-error-in", types.YLeaf{"FrameErrorIn", counters.FrameErrorIn})
    counters.EntityData.Leafs.Append("frame-discard", types.YLeaf{"FrameDiscard", counters.FrameDiscard})
    counters.EntityData.Leafs.Append("tlv-discard", types.YLeaf{"TlvDiscard", counters.TlvDiscard})
    counters.EntityData.Leafs.Append("tlv-unknown", types.YLeaf{"TlvUnknown", counters.TlvUnknown})
    counters.EntityData.Leafs.Append("last-clear", types.YLeaf{"LastClear", counters.LastClear})
    counters.EntityData.Leafs.Append("frame-error-out", types.YLeaf{"FrameErrorOut", counters.FrameErrorOut})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors
// Enclosing container for list of LLDP neighbors on an
// interface
type Lldp_Interfaces_Interface_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of LLDP neighbors. The type is slice of
    // Lldp_Interfaces_Interface_Neighbors_Neighbor.
    Neighbor []*Lldp_Interfaces_Interface_Neighbors_Neighbor
}

func (neighbors *Lldp_Interfaces_Interface_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "interface"
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

// Lldp_Interfaces_Interface_Neighbors_Neighbor
// List of LLDP neighbors
type Lldp_Interfaces_Interface_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key.  . The type is string. Refers to
    // lldp.Lldp_Interfaces_Interface_Neighbors_Neighbor_State_Id
    Id interface{}

    // Configuration data .
    Config Lldp_Interfaces_Interface_Neighbors_Neighbor_Config

    // Operational state data .
    State Lldp_Interfaces_Interface_Neighbors_Neighbor_State

    // Enclosing container for list of custom TLVs from a neighbor.
    CustomTlvs Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs

    // Enclosing container for list of LLDP capabilities.
    Capabilities Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities
}

func (neighbor *Lldp_Interfaces_Interface_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.Id, "id")
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("config", types.YChild{"Config", &neighbor.Config})
    neighbor.EntityData.Children.Append("state", types.YChild{"State", &neighbor.State})
    neighbor.EntityData.Children.Append("custom-tlvs", types.YChild{"CustomTlvs", &neighbor.CustomTlvs})
    neighbor.EntityData.Children.Append("capabilities", types.YChild{"Capabilities", &neighbor.Capabilities})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("id", types.YLeaf{"Id", neighbor.Id})

    neighbor.EntityData.YListKeys = []string {"Id"}

    return &(neighbor.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_Config
// Configuration data 
type Lldp_Interfaces_Interface_Neighbors_Neighbor_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *Lldp_Interfaces_Interface_Neighbors_Neighbor_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "neighbor"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_State
// Operational state data 
type Lldp_Interfaces_Interface_Neighbors_Neighbor_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The system name field shall contain an alpha-numeric string that indicates
    // the system's administratively assigned name. The system name should be the
    // system's fully qualified domain name. If implementations support IETF RFC
    // 3418, the sysName object should be used for this field. The type is string
    // with length: 0..255.
    SystemName interface{}

    // The system description field shall contain an alpha-numeric string that is
    // the textual description of the network entity. The system description
    // should include the full name and version identification of the system's
    // hardware type, software operating system, and networking software. If
    // implementations support IETF RFC 3418, the sysDescr object should be used
    // for this field. The type is string with length: 0..255.
    SystemDescription interface{}

    // The Chassis ID is a mandatory TLV which identifies the chassis component of
    // the endpoint identifier associated with the transmitting LLDP agent. The
    // type is string.
    ChassisId interface{}

    // This field identifies the format and source of the chassis identifier
    // string. It is an enumerator defined by the LldpChassisIdSubtype object from
    // IEEE 802.1AB MIB. The type is ChassisIdType.
    ChassisIdType interface{}

    // System generated identifier for the neighbor on the interface. The type is
    // string.
    Id interface{}

    // Age since discovery. The type is interface{} with range:
    // 0..18446744073709551615. Units are seconds.
    Age interface{}

    // Seconds since last update received. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    LastUpdate interface{}

    // The Port ID is a mandatory TLV which identifies the port component of the
    // endpoint identifier associated with the transmitting LLDP agent. If the
    // specified port is an IEEE 802.3 Repeater port, then this TLV is optional.
    // The type is string.
    PortId interface{}

    // This field identifies the format and source of the port identifier string.
    // It is an enumerator defined by the PtopoPortIdType object from RFC2922. The
    // type is PortIdType.
    PortIdType interface{}

    // The binary string containing the actual port identifier for the port which
    // this LLDP PDU was transmitted. The source and format of this field is
    // defined by PtopoPortId from RFC2922. The type is string.
    PortDescription interface{}

    // The Management Address is a mandatory TLV which identifies a network
    // address associated with the local LLDP agent, which can be used to reach
    // the agent on the port identified in the Port ID TLV. The type is string.
    ManagementAddress interface{}

    // The enumerated value for the network address type identified in this TLV.
    // This enumeration is defined in the 'Assigned Numbers' RFC [RFC3232] and the
    // ianaAddressFamilyNumbers object. The type is string.
    ManagementAddressType interface{}
}

func (state *Lldp_Interfaces_Interface_Neighbors_Neighbor_State) GetEntityData() *types.CommonEntityData {
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
    state.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", state.SystemName})
    state.EntityData.Leafs.Append("system-description", types.YLeaf{"SystemDescription", state.SystemDescription})
    state.EntityData.Leafs.Append("chassis-id", types.YLeaf{"ChassisId", state.ChassisId})
    state.EntityData.Leafs.Append("chassis-id-type", types.YLeaf{"ChassisIdType", state.ChassisIdType})
    state.EntityData.Leafs.Append("id", types.YLeaf{"Id", state.Id})
    state.EntityData.Leafs.Append("age", types.YLeaf{"Age", state.Age})
    state.EntityData.Leafs.Append("last-update", types.YLeaf{"LastUpdate", state.LastUpdate})
    state.EntityData.Leafs.Append("port-id", types.YLeaf{"PortId", state.PortId})
    state.EntityData.Leafs.Append("port-id-type", types.YLeaf{"PortIdType", state.PortIdType})
    state.EntityData.Leafs.Append("port-description", types.YLeaf{"PortDescription", state.PortDescription})
    state.EntityData.Leafs.Append("management-address", types.YLeaf{"ManagementAddress", state.ManagementAddress})
    state.EntityData.Leafs.Append("management-address-type", types.YLeaf{"ManagementAddressType", state.ManagementAddressType})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs
// Enclosing container for list of custom TLVs from a
// neighbor
type Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of custom LLDP TLVs from a neighbor. The type is slice of
    // Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv.
    Tlv []*Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv
}

func (customTlvs *Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs) GetEntityData() *types.CommonEntityData {
    customTlvs.EntityData.YFilter = customTlvs.YFilter
    customTlvs.EntityData.YangName = "custom-tlvs"
    customTlvs.EntityData.BundleName = "openconfig"
    customTlvs.EntityData.ParentYangName = "neighbor"
    customTlvs.EntityData.SegmentPath = "custom-tlvs"
    customTlvs.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    customTlvs.EntityData.NamespaceTable = openconfig.GetNamespaces()
    customTlvs.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    customTlvs.EntityData.Children = types.NewOrderedMap()
    customTlvs.EntityData.Children.Append("tlv", types.YChild{"Tlv", nil})
    for i := range customTlvs.Tlv {
        customTlvs.EntityData.Children.Append(types.GetSegmentPath(customTlvs.Tlv[i]), types.YChild{"Tlv", customTlvs.Tlv[i]})
    }
    customTlvs.EntityData.Leafs = types.NewOrderedMap()

    customTlvs.EntityData.YListKeys = []string {}

    return &(customTlvs.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv
// List of custom LLDP TLVs from a neighbor
type Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to type list key. The type is string
    // with range: -2147483648..2147483647. Refers to
    // lldp.Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_State_Type
    Type interface{}

    // This attribute is a key. Reference to oui list key. The type is string.
    // Refers to
    // lldp.Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_State_Oui
    Oui interface{}

    // This attribute is a key. Reference to oui-subtype list key. The type is
    // string. Refers to
    // lldp.Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_State_OuiSubtype
    OuiSubtype interface{}

    // Configuration data .
    Config Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_Config

    // Operational state data .
    State Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_State
}

func (tlv *Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv) GetEntityData() *types.CommonEntityData {
    tlv.EntityData.YFilter = tlv.YFilter
    tlv.EntityData.YangName = "tlv"
    tlv.EntityData.BundleName = "openconfig"
    tlv.EntityData.ParentYangName = "custom-tlvs"
    tlv.EntityData.SegmentPath = "tlv" + types.AddKeyToken(tlv.Type, "type") + types.AddKeyToken(tlv.Oui, "oui") + types.AddKeyToken(tlv.OuiSubtype, "oui-subtype")
    tlv.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    tlv.EntityData.NamespaceTable = openconfig.GetNamespaces()
    tlv.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    tlv.EntityData.Children = types.NewOrderedMap()
    tlv.EntityData.Children.Append("config", types.YChild{"Config", &tlv.Config})
    tlv.EntityData.Children.Append("state", types.YChild{"State", &tlv.State})
    tlv.EntityData.Leafs = types.NewOrderedMap()
    tlv.EntityData.Leafs.Append("type", types.YLeaf{"Type", tlv.Type})
    tlv.EntityData.Leafs.Append("oui", types.YLeaf{"Oui", tlv.Oui})
    tlv.EntityData.Leafs.Append("oui-subtype", types.YLeaf{"OuiSubtype", tlv.OuiSubtype})

    tlv.EntityData.YListKeys = []string {"Type", "Oui", "OuiSubtype"}

    return &(tlv.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_Config
// Configuration data 
type Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "tlv"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_State
// Operational state data 
type Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The integer value identifying the type of information contained in the
    // value field. The type is interface{} with range: -2147483648..2147483647.
    Type interface{}

    // The organizationally unique identifier field shall contain the
    // organization's OUI as defined in Clause 9 of IEEE Std 802. The high-order
    // octet is 0 and the low-order 3 octets are the SMI Network Management
    // Private Enterprise Code of the Vendor in network byte order, as defined in
    // the 'Assigned Numbers' RFC [RFC3232]. The type is string.
    Oui interface{}

    // The organizationally defined subtype field shall contain a unique subtype
    // value assigned by the defining organization. The type is string.
    OuiSubtype interface{}

    // A variable-length octet-string containing the instance-specific information
    // for this TLV. The type is string.
    Value interface{}
}

func (state *Lldp_Interfaces_Interface_Neighbors_Neighbor_CustomTlvs_Tlv_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "tlv"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})
    state.EntityData.Leafs.Append("oui", types.YLeaf{"Oui", state.Oui})
    state.EntityData.Leafs.Append("oui-subtype", types.YLeaf{"OuiSubtype", state.OuiSubtype})
    state.EntityData.Leafs.Append("value", types.YLeaf{"Value", state.Value})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities
// Enclosing container for list of LLDP capabilities
type Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of LLDP system capabilities advertised by the neighbor. The type is
    // slice of
    // Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability.
    Capability []*Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability
}

func (capabilities *Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities) GetEntityData() *types.CommonEntityData {
    capabilities.EntityData.YFilter = capabilities.YFilter
    capabilities.EntityData.YangName = "capabilities"
    capabilities.EntityData.BundleName = "openconfig"
    capabilities.EntityData.ParentYangName = "neighbor"
    capabilities.EntityData.SegmentPath = "capabilities"
    capabilities.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    capabilities.EntityData.NamespaceTable = openconfig.GetNamespaces()
    capabilities.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    capabilities.EntityData.Children = types.NewOrderedMap()
    capabilities.EntityData.Children.Append("capability", types.YChild{"Capability", nil})
    for i := range capabilities.Capability {
        capabilities.EntityData.Children.Append(types.GetSegmentPath(capabilities.Capability[i]), types.YChild{"Capability", capabilities.Capability[i]})
    }
    capabilities.EntityData.Leafs = types.NewOrderedMap()

    capabilities.EntityData.YListKeys = []string {}

    return &(capabilities.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability
// List of LLDP system capabilities advertised by the
// neighbor
type Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to capabilities list key. The type is
    // one of the following:
    // MACBRIDGEDOCSISCABLEDEVICEREPEATERCVLANTWOPORTMACRELAYSVLANTELEPHONEOTHERROUTERSTATIONONLYWLANACCESSPOINT.
    Name interface{}

    // Configuration data for LLDP capabilities.
    Config Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability_Config

    // Operational state data for LLDP capabilities.
    State Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability_State
}

func (capability *Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability) GetEntityData() *types.CommonEntityData {
    capability.EntityData.YFilter = capability.YFilter
    capability.EntityData.YangName = "capability"
    capability.EntityData.BundleName = "openconfig"
    capability.EntityData.ParentYangName = "capabilities"
    capability.EntityData.SegmentPath = "capability" + types.AddKeyToken(capability.Name, "name")
    capability.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    capability.EntityData.NamespaceTable = openconfig.GetNamespaces()
    capability.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    capability.EntityData.Children = types.NewOrderedMap()
    capability.EntityData.Children.Append("config", types.YChild{"Config", &capability.Config})
    capability.EntityData.Children.Append("state", types.YChild{"State", &capability.State})
    capability.EntityData.Leafs = types.NewOrderedMap()
    capability.EntityData.Leafs.Append("name", types.YLeaf{"Name", capability.Name})

    capability.EntityData.YListKeys = []string {"Name"}

    return &(capability.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability_Config
// Configuration data for LLDP capabilities
type Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (config *Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "capability"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability_State
// Operational state data for LLDP capabilities
type Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the system capability advertised by the neighbor. Capabilities are
    // represented in a bitmap that defines the primary functions of the system.
    // The capabilities are defined in IEEE 802.1AB. The type is one of the
    // following:
    // MACBRIDGEDOCSISCABLEDEVICEREPEATERCVLANTWOPORTMACRELAYSVLANTELEPHONEOTHERROUTERSTATIONONLYWLANACCESSPOINT.
    Name interface{}

    // Indicates whether the corresponding system capability is enabled on the
    // neighbor. The type is bool.
    Enabled interface{}
}

func (state *Lldp_Interfaces_Interface_Neighbors_Neighbor_Capabilities_Capability_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "capability"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

