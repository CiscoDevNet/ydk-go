// Model for managing network interfaces and subinterfaces.  This
// module also defines convenience types / groupings for other
// models to create references to interfaces:
// 
//  base-interface-ref (type) -  reference to a base interface
//  interface-ref (grouping) -  container for reference to a
//    interface + subinterface
//  interface-ref-state (grouping) - container for read-only
//    (opstate) reference to interface + subinterface
// 
// This model reuses data items defined in the IETF YANG model for
// interfaces described by RFC 7223 with an alternate structure
// (particularly for operational state data) and and with
// additional configuration items.
package interfaces

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/openconfig"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package interfaces"))
    ydk.RegisterEntity("{http://openconfig.net/yang/interfaces interfaces}", reflect.TypeOf(Interfaces{}))
    ydk.RegisterEntity("openconfig-interfaces:interfaces", reflect.TypeOf(Interfaces{}))
}

// Interfaces
// Top level container for interfaces, including configuration
// and state data.
type Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of named interfaces on the device. The type is slice of
    // Interfaces_Interface.
    Interface []*Interfaces_Interface
}

func (interfaces *Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "openconfig"
    interfaces.EntityData.ParentYangName = "openconfig-interfaces"
    interfaces.EntityData.SegmentPath = "openconfig-interfaces:interfaces"
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

// Interfaces_Interface
// The list of named interfaces on the device.
type Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured name of the interface.
    // The type is string. Refers to interfaces.Interfaces_Interface_Config_Name
    Name interface{}

    // Configurable items at the global, physical interface level.
    Config Interfaces_Interface_Config

    // Operational state data at the global interface level.
    State Interfaces_Interface_State

    // Top-level container for hold-time settings to enable dampening
    // advertisements of interface transitions.
    HoldTime Interfaces_Interface_HoldTime

    // Enclosing container for the list of subinterfaces associated with a
    // physical interface.
    Subinterfaces Interfaces_Interface_Subinterfaces

    // Top-level container for ethernet configuration and state.
    Ethernet Interfaces_Interface_Ethernet

    // Options for logical interfaces representing aggregates.
    Aggregation Interfaces_Interface_Aggregation

    // Top-level container for routed vlan interfaces.  These logical interfaces
    // are also known as SVI (switched virtual interface), IRB (integrated routing
    // and bridging), RVI (routed VLAN interface).
    RoutedVlan Interfaces_Interface_RoutedVlan

    // Data related to SONET/SDH interfaces.
    Sonet Interfaces_Interface_Sonet
}

func (self *Interfaces_Interface) GetEntityData() *types.CommonEntityData {
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
    self.EntityData.Children.Append("hold-time", types.YChild{"HoldTime", &self.HoldTime})
    self.EntityData.Children.Append("subinterfaces", types.YChild{"Subinterfaces", &self.Subinterfaces})
    self.EntityData.Children.Append("openconfig-if-ethernet:ethernet", types.YChild{"Ethernet", &self.Ethernet})
    self.EntityData.Children.Append("openconfig-if-aggregate:aggregation", types.YChild{"Aggregation", &self.Aggregation})
    self.EntityData.Children.Append("openconfig-vlan:routed-vlan", types.YChild{"RoutedVlan", &self.RoutedVlan})
    self.EntityData.Children.Append("openconfig-transport-line-common:sonet", types.YChild{"Sonet", &self.Sonet})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {"Name"}

    return &(self.EntityData)
}

// Interfaces_Interface_Config
// Configurable items at the global, physical interface
// level
type Interfaces_Interface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF interfaces model (RFC 7223)]  The type of the interface.
    // When an interface entry is created, a server MAY initialize the type leaf
    // with a valid value, e.g., if it is possible to derive the type from the
    // name of the interface.  If a client tries to set the type of an interface
    // to a value that can never be used by the system, e.g., if the type is not
    // supported or if the type does not match the name of the interface, the
    // server MUST reject the request. A NETCONF server MUST reply with an
    // rpc-error with the error-tag 'invalid-value' in this case. The type is one
    // of the following: InterfaceType. This attribute is mandatory.
    Type interface{}

    // Set the max transmission unit size in octets for the physical interface. 
    // If this is not set, the mtu is set to the operational default -- e.g., 1514
    // bytes on an Ethernet interface. The type is interface{} with range:
    // 0..65535.
    Mtu interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The name of the interface.
    // A device MAY restrict the allowed values for this leaf, possibly depending
    // on the type of the interface. For system-controlled interfaces, this leaf
    // is the device-specific name of the interface.  The 'config false' list
    // interfaces/interface[name]/state contains the currently existing interfaces
    // on the device.  If a client tries to create configuration for a
    // system-controlled interface that is not present in the corresponding state
    // list, the server MAY reject the request if the implementation does not
    // support pre-provisioning of interfaces or if the name refers to an
    // interface that can never exist in the system.  A NETCONF server MUST reply
    // with an rpc-error with the error-tag 'invalid-value' in this case.  The
    // IETF model in RFC 7223 provides YANG features for the following (i.e.,
    // pre-provisioning and arbitrary-names), however they are omitted here:   If
    // the device supports pre-provisioning of interface  configuration, the
    // 'pre-provisioning' feature is  advertised.   If the device allows
    // arbitrarily named user-controlled  interfaces, the 'arbitrary-names'
    // feature is advertised.  When a configured user-controlled interface is
    // created by the system, it is instantiated with the same name in the
    // /interfaces/interface[name]/state list. The type is string.
    Name interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  A textual description of
    // the interface.  A server implementation MAY map this leaf to the ifAlias
    // MIB object.  Such an implementation needs to use some mechanism to handle
    // the differences in size and characters allowed between this leaf and
    // ifAlias.  The definition of such a mechanism is outside the scope of this
    // document.  Since ifAlias is defined to be stored in non-volatile storage,
    // the MIB implementation MUST map ifAlias to the value of 'description' in
    // the persistently stored datastore.  Specifically, if the device supports
    // ':startup', when ifAlias is read the device MUST return the value of
    // 'description' in the 'startup' datastore, and when it is written, it MUST
    // be written to the 'running' and 'startup' datastores.  Note that it is up
    // to the implementation to  decide whether to modify this single leaf in
    // 'startup' or perform an implicit copy-config from 'running' to 'startup'. 
    // If the device does not support ':startup', ifAlias MUST be mapped to the
    // 'description' leaf in the 'running' datastore. The type is string.
    Description interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  This leaf contains the
    // configured, desired state of the interface.  Systems that implement the
    // IF-MIB use the value of this leaf in the 'running' datastore to set
    // IF-MIB.ifAdminStatus to 'up' or 'down' after an ifEntry has been
    // initialized, as described in RFC 2863.  Changes in this leaf in the
    // 'running' datastore are reflected in ifAdminStatus, but if ifAdminStatus is
    // changed over SNMP, this leaf is not affected. The type is bool. The default
    // value is true.
    Enabled interface{}
}

func (config *Interfaces_Interface_Config) GetEntityData() *types.CommonEntityData {
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
    config.EntityData.Leafs.Append("type", types.YLeaf{"Type", config.Type})
    config.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", config.Mtu})
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_State
// Operational state data at the global interface level
type Interfaces_Interface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF interfaces model (RFC 7223)]  The type of the interface.
    // When an interface entry is created, a server MAY initialize the type leaf
    // with a valid value, e.g., if it is possible to derive the type from the
    // name of the interface.  If a client tries to set the type of an interface
    // to a value that can never be used by the system, e.g., if the type is not
    // supported or if the type does not match the name of the interface, the
    // server MUST reject the request. A NETCONF server MUST reply with an
    // rpc-error with the error-tag 'invalid-value' in this case. The type is one
    // of the following: InterfaceType. This attribute is mandatory.
    Type interface{}

    // Set the max transmission unit size in octets for the physical interface. 
    // If this is not set, the mtu is set to the operational default -- e.g., 1514
    // bytes on an Ethernet interface. The type is interface{} with range:
    // 0..65535.
    Mtu interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The name of the interface.
    // A device MAY restrict the allowed values for this leaf, possibly depending
    // on the type of the interface. For system-controlled interfaces, this leaf
    // is the device-specific name of the interface.  The 'config false' list
    // interfaces/interface[name]/state contains the currently existing interfaces
    // on the device.  If a client tries to create configuration for a
    // system-controlled interface that is not present in the corresponding state
    // list, the server MAY reject the request if the implementation does not
    // support pre-provisioning of interfaces or if the name refers to an
    // interface that can never exist in the system.  A NETCONF server MUST reply
    // with an rpc-error with the error-tag 'invalid-value' in this case.  The
    // IETF model in RFC 7223 provides YANG features for the following (i.e.,
    // pre-provisioning and arbitrary-names), however they are omitted here:   If
    // the device supports pre-provisioning of interface  configuration, the
    // 'pre-provisioning' feature is  advertised.   If the device allows
    // arbitrarily named user-controlled  interfaces, the 'arbitrary-names'
    // feature is advertised.  When a configured user-controlled interface is
    // created by the system, it is instantiated with the same name in the
    // /interfaces/interface[name]/state list. The type is string.
    Name interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  A textual description of
    // the interface.  A server implementation MAY map this leaf to the ifAlias
    // MIB object.  Such an implementation needs to use some mechanism to handle
    // the differences in size and characters allowed between this leaf and
    // ifAlias.  The definition of such a mechanism is outside the scope of this
    // document.  Since ifAlias is defined to be stored in non-volatile storage,
    // the MIB implementation MUST map ifAlias to the value of 'description' in
    // the persistently stored datastore.  Specifically, if the device supports
    // ':startup', when ifAlias is read the device MUST return the value of
    // 'description' in the 'startup' datastore, and when it is written, it MUST
    // be written to the 'running' and 'startup' datastores.  Note that it is up
    // to the implementation to  decide whether to modify this single leaf in
    // 'startup' or perform an implicit copy-config from 'running' to 'startup'. 
    // If the device does not support ':startup', ifAlias MUST be mapped to the
    // 'description' leaf in the 'running' datastore. The type is string.
    Description interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  This leaf contains the
    // configured, desired state of the interface.  Systems that implement the
    // IF-MIB use the value of this leaf in the 'running' datastore to set
    // IF-MIB.ifAdminStatus to 'up' or 'down' after an ifEntry has been
    // initialized, as described in RFC 2863.  Changes in this leaf in the
    // 'running' datastore are reflected in ifAdminStatus, but if ifAdminStatus is
    // changed over SNMP, this leaf is not affected. The type is bool. The default
    // value is true.
    Enabled interface{}

    // System assigned number for each interface.  Corresponds to ifIndex object
    // in SNMP Interface MIB. The type is interface{} with range: 0..4294967295.
    Ifindex interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The desired state of the
    // interface.  In RFC 7223 this leaf has the same read semantics as
    // ifAdminStatus.  Here, it reflects the administrative state as set by
    // enabling or disabling the interface. The type is AdminStatus. This
    // attribute is mandatory.
    AdminStatus interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The current operational
    // state of the interface.  This leaf has the same semantics as ifOperStatus.
    // The type is OperStatus. This attribute is mandatory.
    OperStatus interface{}

    // Date and time of the last state change of the interface (e.g., up-to-down
    // transition).   This corresponds to the ifLastChange object in the standard
    // interface MIB. The type is interface{} with range: 0..4294967295.
    LastChange interface{}

    // References the hardware port in the device inventory. The type is string.
    // Refers to platform.Components_Component_Name
    HardwarePort interface{}

    // A collection of interface-related statistics objects.
    Counters Interfaces_Interface_State_Counters
}

func (state *Interfaces_Interface_State) GetEntityData() *types.CommonEntityData {
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
    state.EntityData.Leafs.Append("type", types.YLeaf{"Type", state.Type})
    state.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", state.Mtu})
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", state.Ifindex})
    state.EntityData.Leafs.Append("admin-status", types.YLeaf{"AdminStatus", state.AdminStatus})
    state.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", state.OperStatus})
    state.EntityData.Leafs.Append("last-change", types.YLeaf{"LastChange", state.LastChange})
    state.EntityData.Leafs.Append("hardware-port", types.YLeaf{"HardwarePort", state.HardwarePort})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_State_Counters
// A collection of interface-related statistics objects.
type Interfaces_Interface_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF interfaces model (RFC 7223)]  The total number of octets
    // received on the interface, including framing characters.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctets interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The number of packets,
    // delivered by this sub-layer to a higher (sub-)layer, that were not
    // addressed to a multicast or broadcast address at this sub-layer. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InUnicastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The number of packets,
    // delivered by this sub-layer to a higher (sub-)layer, that were addressed to
    // a broadcast address at this sub-layer.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of 'discontinuity-time'. The type
    // is interface{} with range: 0..18446744073709551615.
    InBroadcastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)]   The number of packets,
    // delivered by this sub-layer to a higher (sub-)layer, that were addressed to
    // a multicast address at this sub-layer.  For a MAC-layer protocol, this
    // includes both Group and Functional addresses.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of 'discontinuity-time'. The
    // type is interface{} with range: 0..18446744073709551615.
    InMulticastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  The number of inbound packets that were chosen to be discarded
    // even though no errors had been detected to prevent their being deliverable
    // to a higher-layer protocol.  One possible reason for discarding such a
    // packet could be to free up buffer space.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of 'discontinuity-time'. The type
    // is interface{} with range: 0..18446744073709551615.
    InDiscards interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  For packet-oriented interfaces, the number of inbound packets
    // that contained errors preventing them from being deliverable to a
    // higher-layer protocol.  For character- oriented or fixed-length interfaces,
    // the number of inbound transmission units that contained errors preventing
    // them from being deliverable to a higher-layer protocol.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InErrors interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  For packet-oriented interfaces, the number of packets received
    // via the interface that were discarded because of an unknown or unsupported
    // protocol.  For character-oriented or fixed-length interfaces that support
    // protocol multiplexing, the number of transmission units received via the
    // interface that were discarded because of an unknown or unsupported
    // protocol. For any interface that does not support protocol multiplexing,
    // this counter is not present.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..4294967295.
    InUnknownProtos interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  The total number of octets transmitted out of the interface,
    // including framing characters.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..18446744073709551615.
    OutOctets interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The total number of
    // packets that higher-level protocols requested be transmitted, and that were
    // not addressed to a multicast or broadcast address at this sub-layer,
    // including those that were discarded or not sent.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUnicastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The total number of
    // packets that higher-level protocols requested be transmitted, and that were
    // addressed to a broadcast address at this sub-layer, including those that
    // were discarded or not sent.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..18446744073709551615.
    OutBroadcastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  The total number of packets that higher-level protocols
    // requested be transmitted, and that were addressed to a multicast address at
    // this sub-layer, including those that were discarded or not sent.  For a
    // MAC-layer protocol, this includes both Group and Functional addresses. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMulticastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  The number of outbound packets that were chosen to be discarded
    // even though no errors had been detected to prevent their being transmitted.
    // One possible reason for discarding such a packet could be to free up buffer
    // space.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutDiscards interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  For packet-oriented interfaces, the number of outbound packets
    // that could not be transmitted because of errors. For character-oriented or
    // fixed-length interfaces, the number of outbound transmission units that
    // could not be transmitted because of errors.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of 'discontinuity-time'. The
    // type is interface{} with range: 0..18446744073709551615.
    OutErrors interface{}

    // Indicates the last time the interface counters were cleared. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastClear interface{}
}

func (counters *Interfaces_Interface_State_Counters) GetEntityData() *types.CommonEntityData {
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
    counters.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", counters.InOctets})
    counters.EntityData.Leafs.Append("in-unicast-pkts", types.YLeaf{"InUnicastPkts", counters.InUnicastPkts})
    counters.EntityData.Leafs.Append("in-broadcast-pkts", types.YLeaf{"InBroadcastPkts", counters.InBroadcastPkts})
    counters.EntityData.Leafs.Append("in-multicast-pkts", types.YLeaf{"InMulticastPkts", counters.InMulticastPkts})
    counters.EntityData.Leafs.Append("in-discards", types.YLeaf{"InDiscards", counters.InDiscards})
    counters.EntityData.Leafs.Append("in-errors", types.YLeaf{"InErrors", counters.InErrors})
    counters.EntityData.Leafs.Append("in-unknown-protos", types.YLeaf{"InUnknownProtos", counters.InUnknownProtos})
    counters.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", counters.OutOctets})
    counters.EntityData.Leafs.Append("out-unicast-pkts", types.YLeaf{"OutUnicastPkts", counters.OutUnicastPkts})
    counters.EntityData.Leafs.Append("out-broadcast-pkts", types.YLeaf{"OutBroadcastPkts", counters.OutBroadcastPkts})
    counters.EntityData.Leafs.Append("out-multicast-pkts", types.YLeaf{"OutMulticastPkts", counters.OutMulticastPkts})
    counters.EntityData.Leafs.Append("out-discards", types.YLeaf{"OutDiscards", counters.OutDiscards})
    counters.EntityData.Leafs.Append("out-errors", types.YLeaf{"OutErrors", counters.OutErrors})
    counters.EntityData.Leafs.Append("last-clear", types.YLeaf{"LastClear", counters.LastClear})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Interfaces_Interface_State_AdminStatus represents disabling the interface.
type Interfaces_Interface_State_AdminStatus string

const (
    // Ready to pass packets.
    Interfaces_Interface_State_AdminStatus_UP Interfaces_Interface_State_AdminStatus = "UP"

    // Not ready to pass packets and not in some test mode.
    Interfaces_Interface_State_AdminStatus_DOWN Interfaces_Interface_State_AdminStatus = "DOWN"

    // In some test mode.
    Interfaces_Interface_State_AdminStatus_TESTING Interfaces_Interface_State_AdminStatus = "TESTING"
)

// Interfaces_Interface_State_OperStatus represents This leaf has the same semantics as ifOperStatus.
type Interfaces_Interface_State_OperStatus string

const (
    // Ready to pass packets.
    Interfaces_Interface_State_OperStatus_UP Interfaces_Interface_State_OperStatus = "UP"

    // The interface does not pass any packets.
    Interfaces_Interface_State_OperStatus_DOWN Interfaces_Interface_State_OperStatus = "DOWN"

    // In some test mode.  No operational packets can
    // be passed.
    Interfaces_Interface_State_OperStatus_TESTING Interfaces_Interface_State_OperStatus = "TESTING"

    // Status cannot be determined for some reason.
    Interfaces_Interface_State_OperStatus_UNKNOWN Interfaces_Interface_State_OperStatus = "UNKNOWN"

    // Waiting for some external event.
    Interfaces_Interface_State_OperStatus_DORMANT Interfaces_Interface_State_OperStatus = "DORMANT"

    // Some component (typically hardware) is missing.
    Interfaces_Interface_State_OperStatus_NOT_PRESENT Interfaces_Interface_State_OperStatus = "NOT_PRESENT"

    // Down due to state of lower-layer interface(s).
    Interfaces_Interface_State_OperStatus_LOWER_LAYER_DOWN Interfaces_Interface_State_OperStatus = "LOWER_LAYER_DOWN"
)

// Interfaces_Interface_HoldTime
// Top-level container for hold-time settings to enable
// dampening advertisements of interface transitions.
type Interfaces_Interface_HoldTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for interface hold-time settings.
    Config Interfaces_Interface_HoldTime_Config

    // Operational state data for interface hold-time.
    State Interfaces_Interface_HoldTime_State
}

func (holdTime *Interfaces_Interface_HoldTime) GetEntityData() *types.CommonEntityData {
    holdTime.EntityData.YFilter = holdTime.YFilter
    holdTime.EntityData.YangName = "hold-time"
    holdTime.EntityData.BundleName = "openconfig"
    holdTime.EntityData.ParentYangName = "interface"
    holdTime.EntityData.SegmentPath = "hold-time"
    holdTime.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    holdTime.EntityData.NamespaceTable = openconfig.GetNamespaces()
    holdTime.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    holdTime.EntityData.Children = types.NewOrderedMap()
    holdTime.EntityData.Children.Append("config", types.YChild{"Config", &holdTime.Config})
    holdTime.EntityData.Children.Append("state", types.YChild{"State", &holdTime.State})
    holdTime.EntityData.Leafs = types.NewOrderedMap()

    holdTime.EntityData.YListKeys = []string {}

    return &(holdTime.EntityData)
}

// Interfaces_Interface_HoldTime_Config
// Configuration data for interface hold-time settings.
type Interfaces_Interface_HoldTime_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampens advertisement when the interface transitions from down to up.  A
    // zero value means dampening is turned off, i.e., immediate notification. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds. The
    // default value is 0.
    Up interface{}

    // Dampens advertisement when the interface transitions from up to down.  A
    // zero value means dampening is turned off, i.e., immediate notification. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds. The
    // default value is 0.
    Down interface{}
}

func (config *Interfaces_Interface_HoldTime_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "hold-time"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("up", types.YLeaf{"Up", config.Up})
    config.EntityData.Leafs.Append("down", types.YLeaf{"Down", config.Down})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_HoldTime_State
// Operational state data for interface hold-time.
type Interfaces_Interface_HoldTime_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dampens advertisement when the interface transitions from down to up.  A
    // zero value means dampening is turned off, i.e., immediate notification. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds. The
    // default value is 0.
    Up interface{}

    // Dampens advertisement when the interface transitions from up to down.  A
    // zero value means dampening is turned off, i.e., immediate notification. The
    // type is interface{} with range: 0..4294967295. Units are milliseconds. The
    // default value is 0.
    Down interface{}
}

func (state *Interfaces_Interface_HoldTime_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "hold-time"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("up", types.YLeaf{"Up", state.Up})
    state.EntityData.Leafs.Append("down", types.YLeaf{"Down", state.Down})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces
// Enclosing container for the list of subinterfaces associated
// with a physical interface
type Interfaces_Interface_Subinterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of subinterfaces (logical interfaces) associated with a physical
    // interface. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface.
    Subinterface []*Interfaces_Interface_Subinterfaces_Subinterface
}

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetEntityData() *types.CommonEntityData {
    subinterfaces.EntityData.YFilter = subinterfaces.YFilter
    subinterfaces.EntityData.YangName = "subinterfaces"
    subinterfaces.EntityData.BundleName = "openconfig"
    subinterfaces.EntityData.ParentYangName = "interface"
    subinterfaces.EntityData.SegmentPath = "subinterfaces"
    subinterfaces.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    subinterfaces.EntityData.NamespaceTable = openconfig.GetNamespaces()
    subinterfaces.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    subinterfaces.EntityData.Children = types.NewOrderedMap()
    subinterfaces.EntityData.Children.Append("subinterface", types.YChild{"Subinterface", nil})
    for i := range subinterfaces.Subinterface {
        subinterfaces.EntityData.Children.Append(types.GetSegmentPath(subinterfaces.Subinterface[i]), types.YChild{"Subinterface", subinterfaces.Subinterface[i]})
    }
    subinterfaces.EntityData.Leafs = types.NewOrderedMap()

    subinterfaces.EntityData.YListKeys = []string {}

    return &(subinterfaces.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface
// The list of subinterfaces (logical interfaces) associated
// with a physical interface
type Interfaces_Interface_Subinterfaces_Subinterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index number of the subinterface -- used to
    // address the logical interface. The type is string with range:
    // 0..4294967295. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Config_Index
    Index interface{}

    // Configurable items at the subinterface level.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Config

    // Operational state data for logical interfaces.
    State Interfaces_Interface_Subinterfaces_Subinterface_State

    // Enclosing container for VLAN interface-specific data on subinterfaces.
    Vlan Interfaces_Interface_Subinterfaces_Subinterface_Vlan

    // Parameters for the IPv4 address family.
    Ipv4 Interfaces_Interface_Subinterfaces_Subinterface_Ipv4

    // Parameters for the IPv6 address family.
    Ipv6 Interfaces_Interface_Subinterfaces_Subinterface_Ipv6
}

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetEntityData() *types.CommonEntityData {
    subinterface.EntityData.YFilter = subinterface.YFilter
    subinterface.EntityData.YangName = "subinterface"
    subinterface.EntityData.BundleName = "openconfig"
    subinterface.EntityData.ParentYangName = "subinterfaces"
    subinterface.EntityData.SegmentPath = "subinterface" + types.AddKeyToken(subinterface.Index, "index")
    subinterface.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    subinterface.EntityData.NamespaceTable = openconfig.GetNamespaces()
    subinterface.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    subinterface.EntityData.Children = types.NewOrderedMap()
    subinterface.EntityData.Children.Append("config", types.YChild{"Config", &subinterface.Config})
    subinterface.EntityData.Children.Append("state", types.YChild{"State", &subinterface.State})
    subinterface.EntityData.Children.Append("openconfig-vlan:vlan", types.YChild{"Vlan", &subinterface.Vlan})
    subinterface.EntityData.Children.Append("openconfig-if-ip:ipv4", types.YChild{"Ipv4", &subinterface.Ipv4})
    subinterface.EntityData.Children.Append("openconfig-if-ip:ipv6", types.YChild{"Ipv6", &subinterface.Ipv6})
    subinterface.EntityData.Leafs = types.NewOrderedMap()
    subinterface.EntityData.Leafs.Append("index", types.YLeaf{"Index", subinterface.Index})

    subinterface.EntityData.YListKeys = []string {"Index"}

    return &(subinterface.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Config
// Configurable items at the subinterface level
type Interfaces_Interface_Subinterfaces_Subinterface_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The index of the subinterface, or logical interface number. On systems with
    // no support for subinterfaces, or not using subinterfaces, this value should
    // default to 0, i.e., the default subinterface. The type is interface{} with
    // range: 0..4294967295. The default value is 0.
    Index interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The name of the interface.
    // A device MAY restrict the allowed values for this leaf, possibly depending
    // on the type of the interface. For system-controlled interfaces, this leaf
    // is the device-specific name of the interface.  The 'config false' list
    // interfaces/interface[name]/state contains the currently existing interfaces
    // on the device.  If a client tries to create configuration for a
    // system-controlled interface that is not present in the corresponding state
    // list, the server MAY reject the request if the implementation does not
    // support pre-provisioning of interfaces or if the name refers to an
    // interface that can never exist in the system.  A NETCONF server MUST reply
    // with an rpc-error with the error-tag 'invalid-value' in this case.  The
    // IETF model in RFC 7223 provides YANG features for the following (i.e.,
    // pre-provisioning and arbitrary-names), however they are omitted here:   If
    // the device supports pre-provisioning of interface  configuration, the
    // 'pre-provisioning' feature is  advertised.   If the device allows
    // arbitrarily named user-controlled  interfaces, the 'arbitrary-names'
    // feature is advertised.  When a configured user-controlled interface is
    // created by the system, it is instantiated with the same name in the
    // /interfaces/interface[name]/state list. The type is string.
    Name interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  A textual description of
    // the interface.  A server implementation MAY map this leaf to the ifAlias
    // MIB object.  Such an implementation needs to use some mechanism to handle
    // the differences in size and characters allowed between this leaf and
    // ifAlias.  The definition of such a mechanism is outside the scope of this
    // document.  Since ifAlias is defined to be stored in non-volatile storage,
    // the MIB implementation MUST map ifAlias to the value of 'description' in
    // the persistently stored datastore.  Specifically, if the device supports
    // ':startup', when ifAlias is read the device MUST return the value of
    // 'description' in the 'startup' datastore, and when it is written, it MUST
    // be written to the 'running' and 'startup' datastores.  Note that it is up
    // to the implementation to  decide whether to modify this single leaf in
    // 'startup' or perform an implicit copy-config from 'running' to 'startup'. 
    // If the device does not support ':startup', ifAlias MUST be mapped to the
    // 'description' leaf in the 'running' datastore. The type is string.
    Description interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  This leaf contains the
    // configured, desired state of the interface.  Systems that implement the
    // IF-MIB use the value of this leaf in the 'running' datastore to set
    // IF-MIB.ifAdminStatus to 'up' or 'down' after an ifEntry has been
    // initialized, as described in RFC 2863.  Changes in this leaf in the
    // 'running' datastore are reflected in ifAdminStatus, but if ifAdminStatus is
    // changed over SNMP, this leaf is not affected. The type is bool. The default
    // value is true.
    Enabled interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "subinterface"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("index", types.YLeaf{"Index", config.Index})
    config.EntityData.Leafs.Append("name", types.YLeaf{"Name", config.Name})
    config.EntityData.Leafs.Append("description", types.YLeaf{"Description", config.Description})
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_State
// Operational state data for logical interfaces
type Interfaces_Interface_Subinterfaces_Subinterface_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The index of the subinterface, or logical interface number. On systems with
    // no support for subinterfaces, or not using subinterfaces, this value should
    // default to 0, i.e., the default subinterface. The type is interface{} with
    // range: 0..4294967295. The default value is 0.
    Index interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The name of the interface.
    // A device MAY restrict the allowed values for this leaf, possibly depending
    // on the type of the interface. For system-controlled interfaces, this leaf
    // is the device-specific name of the interface.  The 'config false' list
    // interfaces/interface[name]/state contains the currently existing interfaces
    // on the device.  If a client tries to create configuration for a
    // system-controlled interface that is not present in the corresponding state
    // list, the server MAY reject the request if the implementation does not
    // support pre-provisioning of interfaces or if the name refers to an
    // interface that can never exist in the system.  A NETCONF server MUST reply
    // with an rpc-error with the error-tag 'invalid-value' in this case.  The
    // IETF model in RFC 7223 provides YANG features for the following (i.e.,
    // pre-provisioning and arbitrary-names), however they are omitted here:   If
    // the device supports pre-provisioning of interface  configuration, the
    // 'pre-provisioning' feature is  advertised.   If the device allows
    // arbitrarily named user-controlled  interfaces, the 'arbitrary-names'
    // feature is advertised.  When a configured user-controlled interface is
    // created by the system, it is instantiated with the same name in the
    // /interfaces/interface[name]/state list. The type is string.
    Name interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  A textual description of
    // the interface.  A server implementation MAY map this leaf to the ifAlias
    // MIB object.  Such an implementation needs to use some mechanism to handle
    // the differences in size and characters allowed between this leaf and
    // ifAlias.  The definition of such a mechanism is outside the scope of this
    // document.  Since ifAlias is defined to be stored in non-volatile storage,
    // the MIB implementation MUST map ifAlias to the value of 'description' in
    // the persistently stored datastore.  Specifically, if the device supports
    // ':startup', when ifAlias is read the device MUST return the value of
    // 'description' in the 'startup' datastore, and when it is written, it MUST
    // be written to the 'running' and 'startup' datastores.  Note that it is up
    // to the implementation to  decide whether to modify this single leaf in
    // 'startup' or perform an implicit copy-config from 'running' to 'startup'. 
    // If the device does not support ':startup', ifAlias MUST be mapped to the
    // 'description' leaf in the 'running' datastore. The type is string.
    Description interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  This leaf contains the
    // configured, desired state of the interface.  Systems that implement the
    // IF-MIB use the value of this leaf in the 'running' datastore to set
    // IF-MIB.ifAdminStatus to 'up' or 'down' after an ifEntry has been
    // initialized, as described in RFC 2863.  Changes in this leaf in the
    // 'running' datastore are reflected in ifAdminStatus, but if ifAdminStatus is
    // changed over SNMP, this leaf is not affected. The type is bool. The default
    // value is true.
    Enabled interface{}

    // System assigned number for each interface.  Corresponds to ifIndex object
    // in SNMP Interface MIB. The type is interface{} with range: 0..4294967295.
    Ifindex interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The desired state of the
    // interface.  In RFC 7223 this leaf has the same read semantics as
    // ifAdminStatus.  Here, it reflects the administrative state as set by
    // enabling or disabling the interface. The type is AdminStatus. This
    // attribute is mandatory.
    AdminStatus interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The current operational
    // state of the interface.  This leaf has the same semantics as ifOperStatus.
    // The type is OperStatus. This attribute is mandatory.
    OperStatus interface{}

    // Date and time of the last state change of the interface (e.g., up-to-down
    // transition).   This corresponds to the ifLastChange object in the standard
    // interface MIB. The type is interface{} with range: 0..4294967295.
    LastChange interface{}

    // A collection of interface-related statistics objects.
    Counters Interfaces_Interface_Subinterfaces_Subinterface_State_Counters
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "subinterface"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("counters", types.YChild{"Counters", &state.Counters})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("index", types.YLeaf{"Index", state.Index})
    state.EntityData.Leafs.Append("name", types.YLeaf{"Name", state.Name})
    state.EntityData.Leafs.Append("description", types.YLeaf{"Description", state.Description})
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("ifindex", types.YLeaf{"Ifindex", state.Ifindex})
    state.EntityData.Leafs.Append("admin-status", types.YLeaf{"AdminStatus", state.AdminStatus})
    state.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", state.OperStatus})
    state.EntityData.Leafs.Append("last-change", types.YLeaf{"LastChange", state.LastChange})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_State_Counters
// A collection of interface-related statistics objects.
type Interfaces_Interface_Subinterfaces_Subinterface_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF interfaces model (RFC 7223)]  The total number of octets
    // received on the interface, including framing characters.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctets interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The number of packets,
    // delivered by this sub-layer to a higher (sub-)layer, that were not
    // addressed to a multicast or broadcast address at this sub-layer. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InUnicastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The number of packets,
    // delivered by this sub-layer to a higher (sub-)layer, that were addressed to
    // a broadcast address at this sub-layer.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of 'discontinuity-time'. The type
    // is interface{} with range: 0..18446744073709551615.
    InBroadcastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)]   The number of packets,
    // delivered by this sub-layer to a higher (sub-)layer, that were addressed to
    // a multicast address at this sub-layer.  For a MAC-layer protocol, this
    // includes both Group and Functional addresses.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of 'discontinuity-time'. The
    // type is interface{} with range: 0..18446744073709551615.
    InMulticastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  The number of inbound packets that were chosen to be discarded
    // even though no errors had been detected to prevent their being deliverable
    // to a higher-layer protocol.  One possible reason for discarding such a
    // packet could be to free up buffer space.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of 'discontinuity-time'. The type
    // is interface{} with range: 0..18446744073709551615.
    InDiscards interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  For packet-oriented interfaces, the number of inbound packets
    // that contained errors preventing them from being deliverable to a
    // higher-layer protocol.  For character- oriented or fixed-length interfaces,
    // the number of inbound transmission units that contained errors preventing
    // them from being deliverable to a higher-layer protocol.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InErrors interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  For packet-oriented interfaces, the number of packets received
    // via the interface that were discarded because of an unknown or unsupported
    // protocol.  For character-oriented or fixed-length interfaces that support
    // protocol multiplexing, the number of transmission units received via the
    // interface that were discarded because of an unknown or unsupported
    // protocol. For any interface that does not support protocol multiplexing,
    // this counter is not present.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..4294967295.
    InUnknownProtos interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  The total number of octets transmitted out of the interface,
    // including framing characters.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..18446744073709551615.
    OutOctets interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The total number of
    // packets that higher-level protocols requested be transmitted, and that were
    // not addressed to a multicast or broadcast address at this sub-layer,
    // including those that were discarded or not sent.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUnicastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)]  The total number of
    // packets that higher-level protocols requested be transmitted, and that were
    // addressed to a broadcast address at this sub-layer, including those that
    // were discarded or not sent.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of 'discontinuity-time'. The type is interface{}
    // with range: 0..18446744073709551615.
    OutBroadcastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  The total number of packets that higher-level protocols
    // requested be transmitted, and that were addressed to a multicast address at
    // this sub-layer, including those that were discarded or not sent.  For a
    // MAC-layer protocol, this includes both Group and Functional addresses. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMulticastPkts interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  The number of outbound packets that were chosen to be discarded
    // even though no errors had been detected to prevent their being transmitted.
    // One possible reason for discarding such a packet could be to free up buffer
    // space.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutDiscards interface{}

    // [adapted from IETF interfaces model (RFC 7223)] Changed the counter type to
    // counter64.  For packet-oriented interfaces, the number of outbound packets
    // that could not be transmitted because of errors. For character-oriented or
    // fixed-length interfaces, the number of outbound transmission units that
    // could not be transmitted because of errors.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of 'discontinuity-time'. The
    // type is interface{} with range: 0..18446744073709551615.
    OutErrors interface{}

    // Indicates the last time the interface counters were cleared. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastClear interface{}
}

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetEntityData() *types.CommonEntityData {
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
    counters.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", counters.InOctets})
    counters.EntityData.Leafs.Append("in-unicast-pkts", types.YLeaf{"InUnicastPkts", counters.InUnicastPkts})
    counters.EntityData.Leafs.Append("in-broadcast-pkts", types.YLeaf{"InBroadcastPkts", counters.InBroadcastPkts})
    counters.EntityData.Leafs.Append("in-multicast-pkts", types.YLeaf{"InMulticastPkts", counters.InMulticastPkts})
    counters.EntityData.Leafs.Append("in-discards", types.YLeaf{"InDiscards", counters.InDiscards})
    counters.EntityData.Leafs.Append("in-errors", types.YLeaf{"InErrors", counters.InErrors})
    counters.EntityData.Leafs.Append("in-unknown-protos", types.YLeaf{"InUnknownProtos", counters.InUnknownProtos})
    counters.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", counters.OutOctets})
    counters.EntityData.Leafs.Append("out-unicast-pkts", types.YLeaf{"OutUnicastPkts", counters.OutUnicastPkts})
    counters.EntityData.Leafs.Append("out-broadcast-pkts", types.YLeaf{"OutBroadcastPkts", counters.OutBroadcastPkts})
    counters.EntityData.Leafs.Append("out-multicast-pkts", types.YLeaf{"OutMulticastPkts", counters.OutMulticastPkts})
    counters.EntityData.Leafs.Append("out-discards", types.YLeaf{"OutDiscards", counters.OutDiscards})
    counters.EntityData.Leafs.Append("out-errors", types.YLeaf{"OutErrors", counters.OutErrors})
    counters.EntityData.Leafs.Append("last-clear", types.YLeaf{"LastClear", counters.LastClear})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_State_AdminStatus represents disabling the interface.
type Interfaces_Interface_Subinterfaces_Subinterface_State_AdminStatus string

const (
    // Ready to pass packets.
    Interfaces_Interface_Subinterfaces_Subinterface_State_AdminStatus_UP Interfaces_Interface_Subinterfaces_Subinterface_State_AdminStatus = "UP"

    // Not ready to pass packets and not in some test mode.
    Interfaces_Interface_Subinterfaces_Subinterface_State_AdminStatus_DOWN Interfaces_Interface_Subinterfaces_Subinterface_State_AdminStatus = "DOWN"

    // In some test mode.
    Interfaces_Interface_Subinterfaces_Subinterface_State_AdminStatus_TESTING Interfaces_Interface_Subinterfaces_Subinterface_State_AdminStatus = "TESTING"
)

// Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus represents This leaf has the same semantics as ifOperStatus.
type Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus string

const (
    // Ready to pass packets.
    Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus_UP Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus = "UP"

    // The interface does not pass any packets.
    Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus_DOWN Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus = "DOWN"

    // In some test mode.  No operational packets can
    // be passed.
    Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus_TESTING Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus = "TESTING"

    // Status cannot be determined for some reason.
    Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus_UNKNOWN Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus = "UNKNOWN"

    // Waiting for some external event.
    Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus_DORMANT Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus = "DORMANT"

    // Some component (typically hardware) is missing.
    Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus_NOT_PRESENT Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus = "NOT_PRESENT"

    // Down due to state of lower-layer interface(s).
    Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus_LOWER_LAYER_DOWN Interfaces_Interface_Subinterfaces_Subinterface_State_OperStatus = "LOWER_LAYER_DOWN"
)

// Interfaces_Interface_Subinterfaces_Subinterface_Vlan
// Enclosing container for VLAN interface-specific
// data on subinterfaces
type Interfaces_Interface_Subinterfaces_Subinterface_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters for VLANs.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config

    // State variables for VLANs.
    State Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State
}

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "openconfig"
    vlan.EntityData.ParentYangName = "subinterface"
    vlan.EntityData.SegmentPath = "openconfig-vlan:vlan"
    vlan.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vlan.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vlan.EntityData.Children = types.NewOrderedMap()
    vlan.EntityData.Children.Append("config", types.YChild{"Config", &vlan.Config})
    vlan.EntityData.Children.Append("state", types.YChild{"State", &vlan.State})
    vlan.EntityData.Leafs = types.NewOrderedMap()

    vlan.EntityData.YListKeys = []string {}

    return &(vlan.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config
// Configuration parameters for VLANs
type Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN id for the subinterface -- specified inline for the case of a local
    // VLAN.  The id is scoped to the subinterface, and could be repeated on
    // different subinterfaces. The type is one of the following types: int with
    // range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    VlanId interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "vlan"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", config.VlanId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State
// State variables for VLANs
type Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VLAN id for the subinterface -- specified inline for the case of a local
    // VLAN.  The id is scoped to the subinterface, and could be repeated on
    // different subinterfaces. The type is one of the following types: int with
    // range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    VlanId interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "vlan"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", state.VlanId})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4
// Parameters for the IPv4 address family.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for address list.
    Addresses Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses

    // Enclosing container for neighbor list.
    Neighbors Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors

    // Top-level container for setting unnumbered interfaces. Includes reference
    // the interface that provides the address information.
    Unnumbered Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered

    // Top-level IPv4 configuration data for the interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config

    // Top level IPv4 operational state data.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State
}

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "openconfig"
    ipv4.EntityData.ParentYangName = "subinterface"
    ipv4.EntityData.SegmentPath = "openconfig-if-ip:ipv4"
    ipv4.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("addresses", types.YChild{"Addresses", &ipv4.Addresses})
    ipv4.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &ipv4.Neighbors})
    ipv4.EntityData.Children.Append("unnumbered", types.YChild{"Unnumbered", &ipv4.Unnumbered})
    ipv4.EntityData.Children.Append("config", types.YChild{"Config", &ipv4.Config})
    ipv4.EntityData.Children.Append("state", types.YChild{"State", &ipv4.State})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses
// Enclosing container for address list
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of configured IPv4 addresses on the interface. The type is slice
    // of Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address.
    Address []*Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "openconfig"
    addresses.EntityData.ParentYangName = "ipv4"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    addresses.EntityData.NamespaceTable = openconfig.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range addresses.Address {
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.Address[i]), types.YChild{"Address", addresses.Address[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address
// The list of configured IPv4 addresses on the interface.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured IP address. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config_Ip
    Ip interface{}

    // Configuration data for each configured IPv4 address on the interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config

    // Operational state data for each IPv4 address configured on the interface.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State

    // Enclosing container for VRRP groups handled by this IP interface.
    Vrrp Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "openconfig"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + types.AddKeyToken(address.Ip, "ip")
    address.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    address.EntityData.NamespaceTable = openconfig.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("config", types.YChild{"Config", &address.Config})
    address.EntityData.Children.Append("state", types.YChild{"State", &address.State})
    address.EntityData.Children.Append("vrrp", types.YChild{"Vrrp", &address.Vrrp})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", address.Ip})

    address.EntityData.YListKeys = []string {"Ip"}

    return &(address.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config
// Configuration data for each configured IPv4
// address on the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv4 address on the interface.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..32.
    PrefixLength interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "address"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", config.Ip})
    config.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", config.PrefixLength})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State
// Operational state data for each IPv4 address
// configured on the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv4 address on the interface.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..32.
    PrefixLength interface{}

    // The origin of this address, e.g., statically configured, assigned by DHCP,
    // etc.. The type is IpAddressOrigin.
    Origin interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "address"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", state.Ip})
    state.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", state.PrefixLength})
    state.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", state.Origin})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp
// Enclosing container for VRRP groups handled by this
// IP interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of VRRP groups, keyed by virtual router id. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup.
    VrrpGroup []*Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetEntityData() *types.CommonEntityData {
    vrrp.EntityData.YFilter = vrrp.YFilter
    vrrp.EntityData.YangName = "vrrp"
    vrrp.EntityData.BundleName = "openconfig"
    vrrp.EntityData.ParentYangName = "address"
    vrrp.EntityData.SegmentPath = "vrrp"
    vrrp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vrrp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vrrp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vrrp.EntityData.Children = types.NewOrderedMap()
    vrrp.EntityData.Children.Append("vrrp-group", types.YChild{"VrrpGroup", nil})
    for i := range vrrp.VrrpGroup {
        vrrp.EntityData.Children.Append(types.GetSegmentPath(vrrp.VrrpGroup[i]), types.YChild{"VrrpGroup", vrrp.VrrpGroup[i]})
    }
    vrrp.EntityData.Leafs = types.NewOrderedMap()

    vrrp.EntityData.YListKeys = []string {}

    return &(vrrp.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup
// List of VRRP groups, keyed by virtual router id
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured virtual router id for
    // this VRRP group. The type is string with range: 1..255. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config_VirtualRouterId
    VirtualRouterId interface{}

    // Configuration data for the VRRP group.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config

    // Operational state data for the VRRP group.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State

    // Top-level container for VRRP interface tracking.
    InterfaceTracking Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetEntityData() *types.CommonEntityData {
    vrrpGroup.EntityData.YFilter = vrrpGroup.YFilter
    vrrpGroup.EntityData.YangName = "vrrp-group"
    vrrpGroup.EntityData.BundleName = "openconfig"
    vrrpGroup.EntityData.ParentYangName = "vrrp"
    vrrpGroup.EntityData.SegmentPath = "vrrp-group" + types.AddKeyToken(vrrpGroup.VirtualRouterId, "virtual-router-id")
    vrrpGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vrrpGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vrrpGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vrrpGroup.EntityData.Children = types.NewOrderedMap()
    vrrpGroup.EntityData.Children.Append("config", types.YChild{"Config", &vrrpGroup.Config})
    vrrpGroup.EntityData.Children.Append("state", types.YChild{"State", &vrrpGroup.State})
    vrrpGroup.EntityData.Children.Append("interface-tracking", types.YChild{"InterfaceTracking", &vrrpGroup.InterfaceTracking})
    vrrpGroup.EntityData.Leafs = types.NewOrderedMap()
    vrrpGroup.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", vrrpGroup.VirtualRouterId})

    vrrpGroup.EntityData.YListKeys = []string {"VirtualRouterId"}

    return &(vrrpGroup.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config
// Configuration data for the VRRP group
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the virtual router id for use by the VRRP group.  This usually also
    // determines the virtual MAC address that is generated for the VRRP group.
    // The type is interface{} with range: 1..255.
    VirtualRouterId interface{}

    // Configure one or more virtual addresses for the VRRP group. The type is one
    // of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualAddress []interface{}

    // Specifies the sending VRRP interface's priority for the virtual router. 
    // Higher values equal higher priority. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // When set to true, enables preemption by a higher priority backup router of
    // a lower priority master router. The type is bool. The default value is
    // true.
    Preempt interface{}

    // Set the delay the higher priority router waits before preempting. The type
    // is interface{} with range: 0..3600. The default value is 0.
    PreemptDelay interface{}

    // Configure whether packets destined for virtual addresses are accepted even
    // when the virtual address is not owned by the router interface. The type is
    // bool. The default value is false.
    AcceptMode interface{}

    // Sets the interval between successive VRRP advertisements -- RFC 5798
    // defines this as a 12-bit value expressed as 0.1 seconds, with default 100,
    // i.e., 1 second.  Several implementation express this in units of seconds.
    // The type is interface{} with range: 1..4095. Units are centiseconds. The
    // default value is 100.
    AdvertisementInterval interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "vrrp-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", config.VirtualRouterId})
    config.EntityData.Leafs.Append("virtual-address", types.YLeaf{"VirtualAddress", config.VirtualAddress})
    config.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", config.Priority})
    config.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", config.Preempt})
    config.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", config.PreemptDelay})
    config.EntityData.Leafs.Append("accept-mode", types.YLeaf{"AcceptMode", config.AcceptMode})
    config.EntityData.Leafs.Append("advertisement-interval", types.YLeaf{"AdvertisementInterval", config.AdvertisementInterval})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State
// Operational state data for the VRRP group
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the virtual router id for use by the VRRP group.  This usually also
    // determines the virtual MAC address that is generated for the VRRP group.
    // The type is interface{} with range: 1..255.
    VirtualRouterId interface{}

    // Configure one or more virtual addresses for the VRRP group. The type is one
    // of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualAddress []interface{}

    // Specifies the sending VRRP interface's priority for the virtual router. 
    // Higher values equal higher priority. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // When set to true, enables preemption by a higher priority backup router of
    // a lower priority master router. The type is bool. The default value is
    // true.
    Preempt interface{}

    // Set the delay the higher priority router waits before preempting. The type
    // is interface{} with range: 0..3600. The default value is 0.
    PreemptDelay interface{}

    // Configure whether packets destined for virtual addresses are accepted even
    // when the virtual address is not owned by the router interface. The type is
    // bool. The default value is false.
    AcceptMode interface{}

    // Sets the interval between successive VRRP advertisements -- RFC 5798
    // defines this as a 12-bit value expressed as 0.1 seconds, with default 100,
    // i.e., 1 second.  Several implementation express this in units of seconds.
    // The type is interface{} with range: 1..4095. Units are centiseconds. The
    // default value is 100.
    AdvertisementInterval interface{}

    // Operational value of the priority for the interface in the VRRP group. The
    // type is interface{} with range: 0..255.
    CurrentPriority interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "vrrp-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", state.VirtualRouterId})
    state.EntityData.Leafs.Append("virtual-address", types.YLeaf{"VirtualAddress", state.VirtualAddress})
    state.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", state.Priority})
    state.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", state.Preempt})
    state.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", state.PreemptDelay})
    state.EntityData.Leafs.Append("accept-mode", types.YLeaf{"AcceptMode", state.AcceptMode})
    state.EntityData.Leafs.Append("advertisement-interval", types.YLeaf{"AdvertisementInterval", state.AdvertisementInterval})
    state.EntityData.Leafs.Append("current-priority", types.YLeaf{"CurrentPriority", state.CurrentPriority})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
// Top-level container for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for VRRP interface tracking.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config

    // Operational state data for VRRP interface tracking.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetEntityData() *types.CommonEntityData {
    interfaceTracking.EntityData.YFilter = interfaceTracking.YFilter
    interfaceTracking.EntityData.YangName = "interface-tracking"
    interfaceTracking.EntityData.BundleName = "openconfig"
    interfaceTracking.EntityData.ParentYangName = "vrrp-group"
    interfaceTracking.EntityData.SegmentPath = "interface-tracking"
    interfaceTracking.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceTracking.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceTracking.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceTracking.EntityData.Children = types.NewOrderedMap()
    interfaceTracking.EntityData.Children.Append("config", types.YChild{"Config", &interfaceTracking.Config})
    interfaceTracking.EntityData.Children.Append("state", types.YChild{"State", &interfaceTracking.State})
    interfaceTracking.EntityData.Leafs = types.NewOrderedMap()

    interfaceTracking.EntityData.YListKeys = []string {}

    return &(interfaceTracking.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config
// Configuration data for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets an interface that should be tracked for up/down events to dynamically
    // change the priority state of the VRRP group, and potentially change the
    // mastership if the tracked interface going down lowers the priority
    // sufficiently. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    TrackInterface interface{}

    // Set the value to subtract from priority when the tracked interface goes
    // down. The type is interface{} with range: 0..254. The default value is 0.
    PriorityDecrement interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-tracking"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("track-interface", types.YLeaf{"TrackInterface", config.TrackInterface})
    config.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", config.PriorityDecrement})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
// Operational state data for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets an interface that should be tracked for up/down events to dynamically
    // change the priority state of the VRRP group, and potentially change the
    // mastership if the tracked interface going down lowers the priority
    // sufficiently. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    TrackInterface interface{}

    // Set the value to subtract from priority when the tracked interface goes
    // down. The type is interface{} with range: 0..254. The default value is 0.
    PriorityDecrement interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-tracking"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("track-interface", types.YLeaf{"TrackInterface", state.TrackInterface})
    state.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", state.PriorityDecrement})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors
// Enclosing container for neighbor list
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of mappings from IPv4 addresses to link-layer addresses.  Entries in
    // this list are used as static entries in the ARP Cache. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor.
    Neighbor []*Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "ipv4"
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

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor
// A list of mappings from IPv4 addresses to
// link-layer addresses.
// 
// Entries in this list are used as static entries in the
// ARP Cache.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured IP address. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config_Ip
    Ip interface{}

    // Configuration data for each configured IPv4 address on the interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config

    // Operational state data for each IPv4 address configured on the interface.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.Ip, "ip")
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("config", types.YChild{"Config", &neighbor.Config})
    neighbor.EntityData.Children.Append("state", types.YChild{"State", &neighbor.State})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", neighbor.Ip})

    neighbor.EntityData.YListKeys = []string {"Ip"}

    return &(neighbor.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config
// Configuration data for each configured IPv4
// address on the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IPv4 address of the neighbor node. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetEntityData() *types.CommonEntityData {
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
    config.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", config.Ip})
    config.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", config.LinkLayerAddress})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State
// Operational state data for each IPv4 address
// configured on the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IPv4 address of the neighbor node. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}

    // The origin of this neighbor entry, static or dynamic. The type is
    // NeighborOrigin.
    Origin interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetEntityData() *types.CommonEntityData {
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
    state.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", state.Ip})
    state.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", state.LinkLayerAddress})
    state.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", state.Origin})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered
// Top-level container for setting unnumbered interfaces.
// Includes reference the interface that provides the
// address information
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for unnumbered interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config

    // Operational state data for unnumbered interfaces.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State

    // Reference to an interface or subinterface.
    InterfaceRef Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetEntityData() *types.CommonEntityData {
    unnumbered.EntityData.YFilter = unnumbered.YFilter
    unnumbered.EntityData.YangName = "unnumbered"
    unnumbered.EntityData.BundleName = "openconfig"
    unnumbered.EntityData.ParentYangName = "ipv4"
    unnumbered.EntityData.SegmentPath = "unnumbered"
    unnumbered.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unnumbered.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unnumbered.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unnumbered.EntityData.Children = types.NewOrderedMap()
    unnumbered.EntityData.Children.Append("config", types.YChild{"Config", &unnumbered.Config})
    unnumbered.EntityData.Children.Append("state", types.YChild{"State", &unnumbered.State})
    unnumbered.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &unnumbered.InterfaceRef})
    unnumbered.EntityData.Leafs = types.NewOrderedMap()

    unnumbered.EntityData.YListKeys = []string {}

    return &(unnumbered.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config
// Configuration data for unnumbered interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "unnumbered"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State
// Operational state data for unnumbered interfaces
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "unnumbered"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef
// Reference to an interface or subinterface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config

    // Operational state for interface-ref.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "unnumbered"
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

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config
// Configured reference to interface / subinterface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config struct {
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
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

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State
// Operational state for interface-ref
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State struct {
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
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

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config
// Top-level IPv4 configuration data for the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controls whether IPv4 is enabled or disabled on this interface.  When IPv4
    // is enabled, this interface is connected to an IPv4 stack, and the interface
    // can send and receive IPv4 packets. The type is bool. The default value is
    // true.
    Enabled interface{}

    // The size, in octets, of the largest IPv4 packet that the interface will
    // send and receive.  The server may restrict the allowed values for this
    // leaf, depending on the interface's type.  If this leaf is not configured,
    // the operationally used MTU depends on the interface's type. The type is
    // interface{} with range: 68..65535. Units are octets.
    Mtu interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv4"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", config.Mtu})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State
// Top level IPv4 operational state data
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controls whether IPv4 is enabled or disabled on this interface.  When IPv4
    // is enabled, this interface is connected to an IPv4 stack, and the interface
    // can send and receive IPv4 packets. The type is bool. The default value is
    // true.
    Enabled interface{}

    // The size, in octets, of the largest IPv4 packet that the interface will
    // send and receive.  The server may restrict the allowed values for this
    // leaf, depending on the interface's type.  If this leaf is not configured,
    // the operationally used MTU depends on the interface's type. The type is
    // interface{} with range: 68..65535. Units are octets.
    Mtu interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv4"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", state.Mtu})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6
// Parameters for the IPv6 address family.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for address list.
    Addresses Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses

    // Enclosing container for list of IPv6 neighbors.
    Neighbors Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors

    // Top-level container for setting unnumbered interfaces. Includes reference
    // the interface that provides the address information.
    Unnumbered Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered

    // Top-level config data for the IPv6 interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config

    // Top-level operational state data for the IPv6 interface.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State
}

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "openconfig"
    ipv6.EntityData.ParentYangName = "subinterface"
    ipv6.EntityData.SegmentPath = "openconfig-if-ip:ipv6"
    ipv6.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("addresses", types.YChild{"Addresses", &ipv6.Addresses})
    ipv6.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &ipv6.Neighbors})
    ipv6.EntityData.Children.Append("unnumbered", types.YChild{"Unnumbered", &ipv6.Unnumbered})
    ipv6.EntityData.Children.Append("config", types.YChild{"Config", &ipv6.Config})
    ipv6.EntityData.Children.Append("state", types.YChild{"State", &ipv6.State})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses
// Enclosing container for address list
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of configured IPv6 addresses on the interface. The type is slice
    // of Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address.
    Address []*Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "openconfig"
    addresses.EntityData.ParentYangName = "ipv6"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    addresses.EntityData.NamespaceTable = openconfig.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range addresses.Address {
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.Address[i]), types.YChild{"Address", addresses.Address[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address
// The list of configured IPv6 addresses on the interface.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured IP address. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config_Ip
    Ip interface{}

    // Configuration data for each IPv6 address on the interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config

    // State data for each IPv6 address on the interface.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State

    // Enclosing container for VRRP groups handled by this IP interface.
    Vrrp Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "openconfig"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + types.AddKeyToken(address.Ip, "ip")
    address.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    address.EntityData.NamespaceTable = openconfig.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("config", types.YChild{"Config", &address.Config})
    address.EntityData.Children.Append("state", types.YChild{"State", &address.State})
    address.EntityData.Children.Append("vrrp", types.YChild{"Vrrp", &address.Vrrp})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", address.Ip})

    address.EntityData.YListKeys = []string {"Ip"}

    return &(address.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config
// Configuration data for each IPv6 address on
// the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address on the interface.
    // The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..128. This attribute is mandatory.
    PrefixLength interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "address"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", config.Ip})
    config.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", config.PrefixLength})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State
// State data for each IPv6 address on the
// interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address on the interface.
    // The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..128. This attribute is mandatory.
    PrefixLength interface{}

    // [adapted from IETF IP model RFC 7277]  The origin of this address, e.g.,
    // static, dhcp, etc. The type is IpAddressOrigin.
    Origin interface{}

    // [adapted from IETF IP model RFC 7277]  The status of an address.  Most of
    // the states correspond to states from the IPv6 Stateless Address
    // Autoconfiguration protocol. The type is Status.
    Status interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "address"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", state.Ip})
    state.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", state.PrefixLength})
    state.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", state.Origin})
    state.EntityData.Leafs.Append("status", types.YLeaf{"Status", state.Status})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status represents Autoconfiguration protocol.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status string

const (
    // This is a valid address that can appear as the
    // destination or source address of a packet.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status_PREFERRED Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status = "PREFERRED"

    // This is a valid but deprecated address that should
    // no longer be used as a source address in new
    // communications, but packets addressed to such an
    // address are processed as expected.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status_DEPRECATED Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status = "DEPRECATED"

    // This isn't a valid address, and it shouldn't appear
    // as the destination or source address of a packet.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status_INVALID Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status = "INVALID"

    // The address is not accessible because the interface
    // to which this address is assigned is not
    // operational.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status_INACCESSIBLE Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status = "INACCESSIBLE"

    // The status cannot be determined for some reason.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status_UNKNOWN Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status = "UNKNOWN"

    // The uniqueness of the address on the link is being
    // verified.  Addresses in this state should not be
    // used for general communication and should only be
    // used to determine the uniqueness of the address.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status_TENTATIVE Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status = "TENTATIVE"

    // The address has been determined to be non-unique on
    // the link and so must not be used.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status_DUPLICATE Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status = "DUPLICATE"

    // The address is available for use, subject to
    // restrictions, while its uniqueness on a link is
    // being verified.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status_OPTIMISTIC Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State_Status = "OPTIMISTIC"
)

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp
// Enclosing container for VRRP groups handled by this
// IP interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of VRRP groups, keyed by virtual router id. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup.
    VrrpGroup []*Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetEntityData() *types.CommonEntityData {
    vrrp.EntityData.YFilter = vrrp.YFilter
    vrrp.EntityData.YangName = "vrrp"
    vrrp.EntityData.BundleName = "openconfig"
    vrrp.EntityData.ParentYangName = "address"
    vrrp.EntityData.SegmentPath = "vrrp"
    vrrp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vrrp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vrrp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vrrp.EntityData.Children = types.NewOrderedMap()
    vrrp.EntityData.Children.Append("vrrp-group", types.YChild{"VrrpGroup", nil})
    for i := range vrrp.VrrpGroup {
        vrrp.EntityData.Children.Append(types.GetSegmentPath(vrrp.VrrpGroup[i]), types.YChild{"VrrpGroup", vrrp.VrrpGroup[i]})
    }
    vrrp.EntityData.Leafs = types.NewOrderedMap()

    vrrp.EntityData.YListKeys = []string {}

    return &(vrrp.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup
// List of VRRP groups, keyed by virtual router id
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured virtual router id for
    // this VRRP group. The type is string with range: 1..255. Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config_VirtualRouterId
    VirtualRouterId interface{}

    // Configuration data for the VRRP group.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config

    // Operational state data for the VRRP group.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State

    // Top-level container for VRRP interface tracking.
    InterfaceTracking Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetEntityData() *types.CommonEntityData {
    vrrpGroup.EntityData.YFilter = vrrpGroup.YFilter
    vrrpGroup.EntityData.YangName = "vrrp-group"
    vrrpGroup.EntityData.BundleName = "openconfig"
    vrrpGroup.EntityData.ParentYangName = "vrrp"
    vrrpGroup.EntityData.SegmentPath = "vrrp-group" + types.AddKeyToken(vrrpGroup.VirtualRouterId, "virtual-router-id")
    vrrpGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vrrpGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vrrpGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vrrpGroup.EntityData.Children = types.NewOrderedMap()
    vrrpGroup.EntityData.Children.Append("config", types.YChild{"Config", &vrrpGroup.Config})
    vrrpGroup.EntityData.Children.Append("state", types.YChild{"State", &vrrpGroup.State})
    vrrpGroup.EntityData.Children.Append("interface-tracking", types.YChild{"InterfaceTracking", &vrrpGroup.InterfaceTracking})
    vrrpGroup.EntityData.Leafs = types.NewOrderedMap()
    vrrpGroup.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", vrrpGroup.VirtualRouterId})

    vrrpGroup.EntityData.YListKeys = []string {"VirtualRouterId"}

    return &(vrrpGroup.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config
// Configuration data for the VRRP group
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the virtual router id for use by the VRRP group.  This usually also
    // determines the virtual MAC address that is generated for the VRRP group.
    // The type is interface{} with range: 1..255.
    VirtualRouterId interface{}

    // Configure one or more virtual addresses for the VRRP group. The type is one
    // of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualAddress []interface{}

    // Specifies the sending VRRP interface's priority for the virtual router. 
    // Higher values equal higher priority. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // When set to true, enables preemption by a higher priority backup router of
    // a lower priority master router. The type is bool. The default value is
    // true.
    Preempt interface{}

    // Set the delay the higher priority router waits before preempting. The type
    // is interface{} with range: 0..3600. The default value is 0.
    PreemptDelay interface{}

    // Configure whether packets destined for virtual addresses are accepted even
    // when the virtual address is not owned by the router interface. The type is
    // bool. The default value is false.
    AcceptMode interface{}

    // Sets the interval between successive VRRP advertisements -- RFC 5798
    // defines this as a 12-bit value expressed as 0.1 seconds, with default 100,
    // i.e., 1 second.  Several implementation express this in units of seconds.
    // The type is interface{} with range: 1..4095. Units are centiseconds. The
    // default value is 100.
    AdvertisementInterval interface{}

    // For VRRP on IPv6 interfaces, sets the virtual link local address. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualLinkLocal interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "vrrp-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", config.VirtualRouterId})
    config.EntityData.Leafs.Append("virtual-address", types.YLeaf{"VirtualAddress", config.VirtualAddress})
    config.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", config.Priority})
    config.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", config.Preempt})
    config.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", config.PreemptDelay})
    config.EntityData.Leafs.Append("accept-mode", types.YLeaf{"AcceptMode", config.AcceptMode})
    config.EntityData.Leafs.Append("advertisement-interval", types.YLeaf{"AdvertisementInterval", config.AdvertisementInterval})
    config.EntityData.Leafs.Append("virtual-link-local", types.YLeaf{"VirtualLinkLocal", config.VirtualLinkLocal})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State
// Operational state data for the VRRP group
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the virtual router id for use by the VRRP group.  This usually also
    // determines the virtual MAC address that is generated for the VRRP group.
    // The type is interface{} with range: 1..255.
    VirtualRouterId interface{}

    // Configure one or more virtual addresses for the VRRP group. The type is one
    // of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualAddress []interface{}

    // Specifies the sending VRRP interface's priority for the virtual router. 
    // Higher values equal higher priority. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // When set to true, enables preemption by a higher priority backup router of
    // a lower priority master router. The type is bool. The default value is
    // true.
    Preempt interface{}

    // Set the delay the higher priority router waits before preempting. The type
    // is interface{} with range: 0..3600. The default value is 0.
    PreemptDelay interface{}

    // Configure whether packets destined for virtual addresses are accepted even
    // when the virtual address is not owned by the router interface. The type is
    // bool. The default value is false.
    AcceptMode interface{}

    // Sets the interval between successive VRRP advertisements -- RFC 5798
    // defines this as a 12-bit value expressed as 0.1 seconds, with default 100,
    // i.e., 1 second.  Several implementation express this in units of seconds.
    // The type is interface{} with range: 1..4095. Units are centiseconds. The
    // default value is 100.
    AdvertisementInterval interface{}

    // Operational value of the priority for the interface in the VRRP group. The
    // type is interface{} with range: 0..255.
    CurrentPriority interface{}

    // For VRRP on IPv6 interfaces, sets the virtual link local address. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualLinkLocal interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "vrrp-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", state.VirtualRouterId})
    state.EntityData.Leafs.Append("virtual-address", types.YLeaf{"VirtualAddress", state.VirtualAddress})
    state.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", state.Priority})
    state.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", state.Preempt})
    state.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", state.PreemptDelay})
    state.EntityData.Leafs.Append("accept-mode", types.YLeaf{"AcceptMode", state.AcceptMode})
    state.EntityData.Leafs.Append("advertisement-interval", types.YLeaf{"AdvertisementInterval", state.AdvertisementInterval})
    state.EntityData.Leafs.Append("current-priority", types.YLeaf{"CurrentPriority", state.CurrentPriority})
    state.EntityData.Leafs.Append("virtual-link-local", types.YLeaf{"VirtualLinkLocal", state.VirtualLinkLocal})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
// Top-level container for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for VRRP interface tracking.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config

    // Operational state data for VRRP interface tracking.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetEntityData() *types.CommonEntityData {
    interfaceTracking.EntityData.YFilter = interfaceTracking.YFilter
    interfaceTracking.EntityData.YangName = "interface-tracking"
    interfaceTracking.EntityData.BundleName = "openconfig"
    interfaceTracking.EntityData.ParentYangName = "vrrp-group"
    interfaceTracking.EntityData.SegmentPath = "interface-tracking"
    interfaceTracking.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceTracking.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceTracking.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceTracking.EntityData.Children = types.NewOrderedMap()
    interfaceTracking.EntityData.Children.Append("config", types.YChild{"Config", &interfaceTracking.Config})
    interfaceTracking.EntityData.Children.Append("state", types.YChild{"State", &interfaceTracking.State})
    interfaceTracking.EntityData.Leafs = types.NewOrderedMap()

    interfaceTracking.EntityData.YListKeys = []string {}

    return &(interfaceTracking.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config
// Configuration data for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets an interface that should be tracked for up/down events to dynamically
    // change the priority state of the VRRP group, and potentially change the
    // mastership if the tracked interface going down lowers the priority
    // sufficiently. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    TrackInterface interface{}

    // Set the value to subtract from priority when the tracked interface goes
    // down. The type is interface{} with range: 0..254. The default value is 0.
    PriorityDecrement interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-tracking"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("track-interface", types.YLeaf{"TrackInterface", config.TrackInterface})
    config.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", config.PriorityDecrement})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
// Operational state data for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets an interface that should be tracked for up/down events to dynamically
    // change the priority state of the VRRP group, and potentially change the
    // mastership if the tracked interface going down lowers the priority
    // sufficiently. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    TrackInterface interface{}

    // Set the value to subtract from priority when the tracked interface goes
    // down. The type is interface{} with range: 0..254. The default value is 0.
    PriorityDecrement interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-tracking"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("track-interface", types.YLeaf{"TrackInterface", state.TrackInterface})
    state.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", state.PriorityDecrement})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors
// Enclosing container for list of IPv6 neighbors
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of IPv6 neighbors. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor.
    Neighbor []*Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "ipv6"
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

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor
// List of IPv6 neighbors
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured IP neighbor address. The
    // type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // Refers to
    // interfaces.Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config_Ip
    Ip interface{}

    // Configuration data for each IPv6 address on the interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config

    // State data for each IPv6 address on the interface.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.Ip, "ip")
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("config", types.YChild{"Config", &neighbor.Config})
    neighbor.EntityData.Children.Append("state", types.YChild{"State", &neighbor.State})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", neighbor.Ip})

    neighbor.EntityData.YListKeys = []string {"Ip"}

    return &(neighbor.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config
// Configuration data for each IPv6 address on
// the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address of the neighbor
    // node. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The link-layer address of the
    // neighbor node. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetEntityData() *types.CommonEntityData {
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
    config.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", config.Ip})
    config.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", config.LinkLayerAddress})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State
// State data for each IPv6 address on the
// interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address of the neighbor
    // node. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The link-layer address of the
    // neighbor node. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}

    // [adapted from IETF IP model RFC 7277]  The origin of this neighbor entry.
    // The type is NeighborOrigin.
    Origin interface{}

    // [adapted from IETF IP model RFC 7277]  Indicates that the neighbor node
    // acts as a router. The type is interface{}.
    IsRouter interface{}

    // [adapted from IETF IP model RFC 7277]  The Neighbor Unreachability
    // Detection state of this entry. The type is NeighborState.
    NeighborState interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetEntityData() *types.CommonEntityData {
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
    state.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", state.Ip})
    state.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", state.LinkLayerAddress})
    state.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", state.Origin})
    state.EntityData.Leafs.Append("is-router", types.YLeaf{"IsRouter", state.IsRouter})
    state.EntityData.Leafs.Append("neighbor-state", types.YLeaf{"NeighborState", state.NeighborState})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState represents entry.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState string

const (
    // Address resolution is in progress, and the link-layer
    //      address of the neighbor has not yet been
    //      determined.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState_INCOMPLETE Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState = "INCOMPLETE"

    // Roughly speaking, the neighbor is known to have been
    //      reachable recently (within tens of seconds ago).
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState_REACHABLE Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState = "REACHABLE"

    // The neighbor is no longer known to be reachable, but
    //      until traffic is sent to the neighbor no attempt
    //      should be made to verify its reachability.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState_STALE Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState = "STALE"

    // The neighbor is no longer known to be reachable, and
    //      traffic has recently been sent to the neighbor.
    //      Rather than probe the neighbor immediately, however,
    //      delay sending probes for a short while in order to
    //      give upper-layer protocols a chance to provide
    //      reachability confirmation.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState_DELAY Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState = "DELAY"

    // The neighbor is no longer known to be reachable, and
    //      unicast Neighbor Solicitation probes are being sent
    //      to verify reachability.
    Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState_PROBE Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State_NeighborState = "PROBE"
)

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered
// Top-level container for setting unnumbered interfaces.
// Includes reference the interface that provides the
// address information
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for unnumbered interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config

    // Operational state data for unnumbered interfaces.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State

    // Reference to an interface or subinterface.
    InterfaceRef Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetEntityData() *types.CommonEntityData {
    unnumbered.EntityData.YFilter = unnumbered.YFilter
    unnumbered.EntityData.YangName = "unnumbered"
    unnumbered.EntityData.BundleName = "openconfig"
    unnumbered.EntityData.ParentYangName = "ipv6"
    unnumbered.EntityData.SegmentPath = "unnumbered"
    unnumbered.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unnumbered.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unnumbered.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unnumbered.EntityData.Children = types.NewOrderedMap()
    unnumbered.EntityData.Children.Append("config", types.YChild{"Config", &unnumbered.Config})
    unnumbered.EntityData.Children.Append("state", types.YChild{"State", &unnumbered.State})
    unnumbered.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &unnumbered.InterfaceRef})
    unnumbered.EntityData.Leafs = types.NewOrderedMap()

    unnumbered.EntityData.YListKeys = []string {}

    return &(unnumbered.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config
// Configuration data for unnumbered interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "unnumbered"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State
// Operational state data for unnumbered interfaces
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "unnumbered"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef
// Reference to an interface or subinterface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config

    // Operational state for interface-ref.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "unnumbered"
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

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config
// Configured reference to interface / subinterface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config struct {
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
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

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State
// Operational state for interface-ref
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State struct {
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
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

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config
// Top-level config data for the IPv6 interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  Controls whether IPv6 is enabled or
    // disabled on this interface.  When IPv6 is enabled, this interface is
    // connected to an IPv6 stack, and the interface can send and receive IPv6
    // packets. The type is bool. The default value is true.
    Enabled interface{}

    // [adapted from IETF IP model RFC 7277]  The size, in octets, of the largest
    // IPv6 packet that the interface will send and receive.  The server may
    // restrict the allowed values for this leaf, depending on the interface's
    // type.  If this leaf is not configured, the operationally used MTU depends
    // on the interface's type. The type is interface{} with range:
    // 1280..4294967295. Units are octets.
    Mtu interface{}

    // [adapted from IETF IP model RFC 7277]  The number of consecutive Neighbor
    // Solicitation messages sent while performing Duplicate Address Detection on
    // a tentative address.  A value of zero indicates that Duplicate Address
    // Detection is not performed on tentative addresses.  A value of one
    // indicates a single transmission with no follow-up retransmissions. The type
    // is interface{} with range: 0..4294967295. The default value is 1.
    DupAddrDetectTransmits interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv6"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", config.Mtu})
    config.EntityData.Leafs.Append("dup-addr-detect-transmits", types.YLeaf{"DupAddrDetectTransmits", config.DupAddrDetectTransmits})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State
// Top-level operational state data for the IPv6 interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  Controls whether IPv6 is enabled or
    // disabled on this interface.  When IPv6 is enabled, this interface is
    // connected to an IPv6 stack, and the interface can send and receive IPv6
    // packets. The type is bool. The default value is true.
    Enabled interface{}

    // [adapted from IETF IP model RFC 7277]  The size, in octets, of the largest
    // IPv6 packet that the interface will send and receive.  The server may
    // restrict the allowed values for this leaf, depending on the interface's
    // type.  If this leaf is not configured, the operationally used MTU depends
    // on the interface's type. The type is interface{} with range:
    // 1280..4294967295. Units are octets.
    Mtu interface{}

    // [adapted from IETF IP model RFC 7277]  The number of consecutive Neighbor
    // Solicitation messages sent while performing Duplicate Address Detection on
    // a tentative address.  A value of zero indicates that Duplicate Address
    // Detection is not performed on tentative addresses.  A value of one
    // indicates a single transmission with no follow-up retransmissions. The type
    // is interface{} with range: 0..4294967295. The default value is 1.
    DupAddrDetectTransmits interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv6"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", state.Mtu})
    state.EntityData.Leafs.Append("dup-addr-detect-transmits", types.YLeaf{"DupAddrDetectTransmits", state.DupAddrDetectTransmits})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Ethernet
// Top-level container for ethernet configuration
// and state
type Interfaces_Interface_Ethernet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for ethernet interfaces.
    Config Interfaces_Interface_Ethernet_Config

    // State variables for Ethernet interfaces.
    State Interfaces_Interface_Ethernet_State

    // Enclosing container for VLAN interface-specific data on Ethernet
    // interfaces.  These are for standard L2, switched-style VLANs.
    SwitchedVlan Interfaces_Interface_Ethernet_SwitchedVlan
}

func (ethernet *Interfaces_Interface_Ethernet) GetEntityData() *types.CommonEntityData {
    ethernet.EntityData.YFilter = ethernet.YFilter
    ethernet.EntityData.YangName = "ethernet"
    ethernet.EntityData.BundleName = "openconfig"
    ethernet.EntityData.ParentYangName = "interface"
    ethernet.EntityData.SegmentPath = "openconfig-if-ethernet:ethernet"
    ethernet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ethernet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ethernet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ethernet.EntityData.Children = types.NewOrderedMap()
    ethernet.EntityData.Children.Append("config", types.YChild{"Config", &ethernet.Config})
    ethernet.EntityData.Children.Append("state", types.YChild{"State", &ethernet.State})
    ethernet.EntityData.Children.Append("openconfig-vlan:switched-vlan", types.YChild{"SwitchedVlan", &ethernet.SwitchedVlan})
    ethernet.EntityData.Leafs = types.NewOrderedMap()

    ethernet.EntityData.YListKeys = []string {}

    return &(ethernet.EntityData)
}

// Interfaces_Interface_Ethernet_Config
// Configuration data for ethernet interfaces
type Interfaces_Interface_Ethernet_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assigns a MAC address to the Ethernet interface.  If not specified, the
    // corresponding operational state leaf is expected to show the
    // system-assigned MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Set to TRUE to request the interface to auto-negotiate transmission
    // parameters with its peer interface.  When set to FALSE, the transmission
    // parameters are specified manually. The type is bool. The default value is
    // true.
    AutoNegotiate interface{}

    // When auto-negotiate is TRUE, this optionally sets the duplex mode that will
    // be advertised to the peer.  If unspecified, the interface should negotiate
    // the duplex mode directly (typically full-duplex).  When auto-negotiate is
    // FALSE, this sets the duplex mode on the interface directly. The type is
    // DuplexMode.
    DuplexMode interface{}

    // When auto-negotiate is TRUE, this optionally sets the port-speed mode that
    // will be advertised to the peer for negotiation.  If unspecified, it is
    // expected that the interface will select the highest speed available based
    // on negotiation.  When auto-negotiate is set to FALSE, sets the link speed
    // to a fixed value -- supported values are defined by ETHERNET_SPEED
    // identities. The type is one of the following:
    // SPEED100GBSPEED1GBSPEED25GBSPEED10GBSPEED10MBSPEED40GBSPEEDUNKNOWNSPEED50GBSPEED100MB.
    PortSpeed interface{}

    // Enable or disable flow control for this interface. Ethernet flow control is
    // a mechanism by which a receiver may send PAUSE frames to a sender to stop
    // transmission for a specified time.  This setting should override
    // auto-negotiated flow control settings.  If left unspecified, and
    // auto-negotiate is TRUE, flow control mode is negotiated with the peer
    // interface. The type is bool. The default value is false.
    EnableFlowControl interface{}

    // Specify the logical aggregate interface to which this interface belongs.
    // The type is string. Refers to interfaces.Interfaces_Interface_Name
    AggregateId interface{}
}

func (config *Interfaces_Interface_Ethernet_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ethernet"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", config.MacAddress})
    config.EntityData.Leafs.Append("auto-negotiate", types.YLeaf{"AutoNegotiate", config.AutoNegotiate})
    config.EntityData.Leafs.Append("duplex-mode", types.YLeaf{"DuplexMode", config.DuplexMode})
    config.EntityData.Leafs.Append("port-speed", types.YLeaf{"PortSpeed", config.PortSpeed})
    config.EntityData.Leafs.Append("enable-flow-control", types.YLeaf{"EnableFlowControl", config.EnableFlowControl})
    config.EntityData.Leafs.Append("aggregate-id", types.YLeaf{"AggregateId", config.AggregateId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Ethernet_Config_DuplexMode represents FALSE, this sets the duplex mode on the interface directly.
type Interfaces_Interface_Ethernet_Config_DuplexMode string

const (
    // Full duplex mode
    Interfaces_Interface_Ethernet_Config_DuplexMode_FULL Interfaces_Interface_Ethernet_Config_DuplexMode = "FULL"

    // Half duplex mode
    Interfaces_Interface_Ethernet_Config_DuplexMode_HALF Interfaces_Interface_Ethernet_Config_DuplexMode = "HALF"
)

// Interfaces_Interface_Ethernet_State
// State variables for Ethernet interfaces
type Interfaces_Interface_Ethernet_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assigns a MAC address to the Ethernet interface.  If not specified, the
    // corresponding operational state leaf is expected to show the
    // system-assigned MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Set to TRUE to request the interface to auto-negotiate transmission
    // parameters with its peer interface.  When set to FALSE, the transmission
    // parameters are specified manually. The type is bool. The default value is
    // true.
    AutoNegotiate interface{}

    // When auto-negotiate is TRUE, this optionally sets the duplex mode that will
    // be advertised to the peer.  If unspecified, the interface should negotiate
    // the duplex mode directly (typically full-duplex).  When auto-negotiate is
    // FALSE, this sets the duplex mode on the interface directly. The type is
    // DuplexMode.
    DuplexMode interface{}

    // When auto-negotiate is TRUE, this optionally sets the port-speed mode that
    // will be advertised to the peer for negotiation.  If unspecified, it is
    // expected that the interface will select the highest speed available based
    // on negotiation.  When auto-negotiate is set to FALSE, sets the link speed
    // to a fixed value -- supported values are defined by ETHERNET_SPEED
    // identities. The type is one of the following:
    // SPEED100GBSPEED1GBSPEED25GBSPEED10GBSPEED10MBSPEED40GBSPEEDUNKNOWNSPEED50GBSPEED100MB.
    PortSpeed interface{}

    // Enable or disable flow control for this interface. Ethernet flow control is
    // a mechanism by which a receiver may send PAUSE frames to a sender to stop
    // transmission for a specified time.  This setting should override
    // auto-negotiated flow control settings.  If left unspecified, and
    // auto-negotiate is TRUE, flow control mode is negotiated with the peer
    // interface. The type is bool. The default value is false.
    EnableFlowControl interface{}

    // Represenets the 'burned-in',  or system-assigned, MAC address for the
    // Ethernet interface. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    HwMacAddress interface{}

    // Reports the effective speed of the interface, e.g., the negotiated speed if
    // auto-negotiate is enabled. The type is interface{} with range:
    // 0..4294967295. Units are Mbps.
    EffectiveSpeed interface{}

    // Specify the logical aggregate interface to which this interface belongs.
    // The type is string. Refers to interfaces.Interfaces_Interface_Name
    AggregateId interface{}

    // Ethernet interface counters.
    Counters Interfaces_Interface_Ethernet_State_Counters
}

func (state *Interfaces_Interface_Ethernet_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ethernet"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("counters", types.YChild{"Counters", &state.Counters})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", state.MacAddress})
    state.EntityData.Leafs.Append("auto-negotiate", types.YLeaf{"AutoNegotiate", state.AutoNegotiate})
    state.EntityData.Leafs.Append("duplex-mode", types.YLeaf{"DuplexMode", state.DuplexMode})
    state.EntityData.Leafs.Append("port-speed", types.YLeaf{"PortSpeed", state.PortSpeed})
    state.EntityData.Leafs.Append("enable-flow-control", types.YLeaf{"EnableFlowControl", state.EnableFlowControl})
    state.EntityData.Leafs.Append("hw-mac-address", types.YLeaf{"HwMacAddress", state.HwMacAddress})
    state.EntityData.Leafs.Append("effective-speed", types.YLeaf{"EffectiveSpeed", state.EffectiveSpeed})
    state.EntityData.Leafs.Append("aggregate-id", types.YLeaf{"AggregateId", state.AggregateId})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Ethernet_State_Counters
// Ethernet interface counters
type Interfaces_Interface_Ethernet_State_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC layer control frames received on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    InMacControlFrames interface{}

    // MAC layer PAUSE frames received on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    InMacPauseFrames interface{}

    // Number of oversize frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    InOversizeFrames interface{}

    // Number of jabber frames received on the interface.  Jabber frames are
    // typically defined as oversize frames which also have a bad CRC. 
    // Implementations may use slightly different definitions of what constitutes
    // a jabber frame.  Often indicative of a NIC hardware problem. The type is
    // interface{} with range: 0..18446744073709551615.
    InJabberFrames interface{}

    // Number of fragment frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    InFragmentFrames interface{}

    // Number of 802.1q tagged frames received on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    In8021qFrames interface{}

    // Number of receive error events due to FCS/CRC check failure. The type is
    // interface{} with range: 0..18446744073709551615.
    InCrcErrors interface{}

    // MAC layer control frames sent on the interface. The type is interface{}
    // with range: 0..18446744073709551615.
    OutMacControlFrames interface{}

    // MAC layer PAUSE frames sent on the interface. The type is interface{} with
    // range: 0..18446744073709551615.
    OutMacPauseFrames interface{}

    // Number of 802.1q tagged frames sent on the interface. The type is
    // interface{} with range: 0..18446744073709551615.
    Out8021qFrames interface{}
}

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetEntityData() *types.CommonEntityData {
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
    counters.EntityData.Leafs.Append("in-mac-control-frames", types.YLeaf{"InMacControlFrames", counters.InMacControlFrames})
    counters.EntityData.Leafs.Append("in-mac-pause-frames", types.YLeaf{"InMacPauseFrames", counters.InMacPauseFrames})
    counters.EntityData.Leafs.Append("in-oversize-frames", types.YLeaf{"InOversizeFrames", counters.InOversizeFrames})
    counters.EntityData.Leafs.Append("in-jabber-frames", types.YLeaf{"InJabberFrames", counters.InJabberFrames})
    counters.EntityData.Leafs.Append("in-fragment-frames", types.YLeaf{"InFragmentFrames", counters.InFragmentFrames})
    counters.EntityData.Leafs.Append("in-8021q-frames", types.YLeaf{"In8021qFrames", counters.In8021qFrames})
    counters.EntityData.Leafs.Append("in-crc-errors", types.YLeaf{"InCrcErrors", counters.InCrcErrors})
    counters.EntityData.Leafs.Append("out-mac-control-frames", types.YLeaf{"OutMacControlFrames", counters.OutMacControlFrames})
    counters.EntityData.Leafs.Append("out-mac-pause-frames", types.YLeaf{"OutMacPauseFrames", counters.OutMacPauseFrames})
    counters.EntityData.Leafs.Append("out-8021q-frames", types.YLeaf{"Out8021qFrames", counters.Out8021qFrames})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Interfaces_Interface_Ethernet_State_DuplexMode represents FALSE, this sets the duplex mode on the interface directly.
type Interfaces_Interface_Ethernet_State_DuplexMode string

const (
    // Full duplex mode
    Interfaces_Interface_Ethernet_State_DuplexMode_FULL Interfaces_Interface_Ethernet_State_DuplexMode = "FULL"

    // Half duplex mode
    Interfaces_Interface_Ethernet_State_DuplexMode_HALF Interfaces_Interface_Ethernet_State_DuplexMode = "HALF"
)

// Interfaces_Interface_Ethernet_SwitchedVlan
// Enclosing container for VLAN interface-specific
// data on Ethernet interfaces.  These are for standard
// L2, switched-style VLANs.
type Interfaces_Interface_Ethernet_SwitchedVlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters for VLANs.
    Config Interfaces_Interface_Ethernet_SwitchedVlan_Config

    // State variables for VLANs.
    State Interfaces_Interface_Ethernet_SwitchedVlan_State
}

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetEntityData() *types.CommonEntityData {
    switchedVlan.EntityData.YFilter = switchedVlan.YFilter
    switchedVlan.EntityData.YangName = "switched-vlan"
    switchedVlan.EntityData.BundleName = "openconfig"
    switchedVlan.EntityData.ParentYangName = "ethernet"
    switchedVlan.EntityData.SegmentPath = "openconfig-vlan:switched-vlan"
    switchedVlan.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    switchedVlan.EntityData.NamespaceTable = openconfig.GetNamespaces()
    switchedVlan.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    switchedVlan.EntityData.Children = types.NewOrderedMap()
    switchedVlan.EntityData.Children.Append("config", types.YChild{"Config", &switchedVlan.Config})
    switchedVlan.EntityData.Children.Append("state", types.YChild{"State", &switchedVlan.State})
    switchedVlan.EntityData.Leafs = types.NewOrderedMap()

    switchedVlan.EntityData.YListKeys = []string {}

    return &(switchedVlan.EntityData)
}

// Interfaces_Interface_Ethernet_SwitchedVlan_Config
// Configuration parameters for VLANs
type Interfaces_Interface_Ethernet_SwitchedVlan_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the interface to access or trunk mode for VLANs. The type is
    // VlanModeType.
    InterfaceMode interface{}

    // Set the native VLAN id for untagged frames arriving on a trunk interface. 
    // This configuration is only valid on a trunk interface. The type is one of
    // the following types: int with range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    NativeVlan interface{}

    // Assign the access vlan to the access port. The type is one of the following
    // types: int with range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    AccessVlan interface{}

    // Specify VLANs, or ranges thereof, that the interface may carry when in
    // trunk mode.  If not specified, all VLANs are allowed on the interface.
    // Ranges are specified in the form x..y, where x<y - ranges are assumed to be
    // inclusive (such that the VLAN range is x <= range <= y. The type is one of
    // the following types: slice of int with range: 1..4094, or slice of string
    // with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])$,
    // or slice of string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$,
    // or slice of string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$,
    // or slice of string with pattern:
    // ^(\*|(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]))\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])$..
    TrunkVlans []interface{}
}

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "switched-vlan"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface-mode", types.YLeaf{"InterfaceMode", config.InterfaceMode})
    config.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", config.NativeVlan})
    config.EntityData.Leafs.Append("access-vlan", types.YLeaf{"AccessVlan", config.AccessVlan})
    config.EntityData.Leafs.Append("trunk-vlans", types.YLeaf{"TrunkVlans", config.TrunkVlans})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Ethernet_SwitchedVlan_State
// State variables for VLANs
type Interfaces_Interface_Ethernet_SwitchedVlan_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the interface to access or trunk mode for VLANs. The type is
    // VlanModeType.
    InterfaceMode interface{}

    // Set the native VLAN id for untagged frames arriving on a trunk interface. 
    // This configuration is only valid on a trunk interface. The type is one of
    // the following types: int with range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    NativeVlan interface{}

    // Assign the access vlan to the access port. The type is one of the following
    // types: int with range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    AccessVlan interface{}

    // Specify VLANs, or ranges thereof, that the interface may carry when in
    // trunk mode.  If not specified, all VLANs are allowed on the interface.
    // Ranges are specified in the form x..y, where x<y - ranges are assumed to be
    // inclusive (such that the VLAN range is x <= range <= y. The type is one of
    // the following types: slice of int with range: 1..4094, or slice of string
    // with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])$,
    // or slice of string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$,
    // or slice of string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$,
    // or slice of string with pattern:
    // ^(\*|(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]))\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])$..
    TrunkVlans []interface{}
}

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "switched-vlan"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface-mode", types.YLeaf{"InterfaceMode", state.InterfaceMode})
    state.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", state.NativeVlan})
    state.EntityData.Leafs.Append("access-vlan", types.YLeaf{"AccessVlan", state.AccessVlan})
    state.EntityData.Leafs.Append("trunk-vlans", types.YLeaf{"TrunkVlans", state.TrunkVlans})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Aggregation
// Options for logical interfaces representing
// aggregates
type Interfaces_Interface_Aggregation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration variables for logical aggregate / LAG interfaces.
    Config Interfaces_Interface_Aggregation_Config

    // Operational state variables for logical aggregate / LAG interfaces.
    State Interfaces_Interface_Aggregation_State

    // Enclosing container for VLAN interface-specific data on Ethernet
    // interfaces.  These are for standard L2, switched-style VLANs.
    SwitchedVlan Interfaces_Interface_Aggregation_SwitchedVlan
}

func (aggregation *Interfaces_Interface_Aggregation) GetEntityData() *types.CommonEntityData {
    aggregation.EntityData.YFilter = aggregation.YFilter
    aggregation.EntityData.YangName = "aggregation"
    aggregation.EntityData.BundleName = "openconfig"
    aggregation.EntityData.ParentYangName = "interface"
    aggregation.EntityData.SegmentPath = "openconfig-if-aggregate:aggregation"
    aggregation.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    aggregation.EntityData.NamespaceTable = openconfig.GetNamespaces()
    aggregation.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    aggregation.EntityData.Children = types.NewOrderedMap()
    aggregation.EntityData.Children.Append("config", types.YChild{"Config", &aggregation.Config})
    aggregation.EntityData.Children.Append("state", types.YChild{"State", &aggregation.State})
    aggregation.EntityData.Children.Append("openconfig-vlan:switched-vlan", types.YChild{"SwitchedVlan", &aggregation.SwitchedVlan})
    aggregation.EntityData.Leafs = types.NewOrderedMap()

    aggregation.EntityData.YListKeys = []string {}

    return &(aggregation.EntityData)
}

// Interfaces_Interface_Aggregation_Config
// Configuration variables for logical aggregate /
// LAG interfaces
type Interfaces_Interface_Aggregation_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets the type of LAG, i.e., how it is configured / maintained. The type is
    // AggregationType.
    LagType interface{}

    // Specifies the mininum number of member interfaces that must be active for
    // the aggregate interface to be available. The type is interface{} with
    // range: 0..65535.
    MinLinks interface{}
}

func (config *Interfaces_Interface_Aggregation_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "aggregation"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("lag-type", types.YLeaf{"LagType", config.LagType})
    config.EntityData.Leafs.Append("min-links", types.YLeaf{"MinLinks", config.MinLinks})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Aggregation_State
// Operational state variables for logical
// aggregate / LAG interfaces
type Interfaces_Interface_Aggregation_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets the type of LAG, i.e., how it is configured / maintained. The type is
    // AggregationType.
    LagType interface{}

    // Specifies the mininum number of member interfaces that must be active for
    // the aggregate interface to be available. The type is interface{} with
    // range: 0..65535.
    MinLinks interface{}

    // Reports effective speed of the aggregate interface, based on speed of
    // active member interfaces. The type is interface{} with range:
    // 0..4294967295. Units are Mbps.
    LagSpeed interface{}

    // List of current member interfaces for the aggregate, expressed as
    // references to existing interfaces. The type is slice of string. Refers to
    // interfaces.Interfaces_Interface_Name
    Member []interface{}
}

func (state *Interfaces_Interface_Aggregation_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "aggregation"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("lag-type", types.YLeaf{"LagType", state.LagType})
    state.EntityData.Leafs.Append("min-links", types.YLeaf{"MinLinks", state.MinLinks})
    state.EntityData.Leafs.Append("lag-speed", types.YLeaf{"LagSpeed", state.LagSpeed})
    state.EntityData.Leafs.Append("member", types.YLeaf{"Member", state.Member})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Aggregation_SwitchedVlan
// Enclosing container for VLAN interface-specific
// data on Ethernet interfaces.  These are for standard
// L2, switched-style VLANs.
type Interfaces_Interface_Aggregation_SwitchedVlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration parameters for VLANs.
    Config Interfaces_Interface_Aggregation_SwitchedVlan_Config

    // State variables for VLANs.
    State Interfaces_Interface_Aggregation_SwitchedVlan_State
}

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetEntityData() *types.CommonEntityData {
    switchedVlan.EntityData.YFilter = switchedVlan.YFilter
    switchedVlan.EntityData.YangName = "switched-vlan"
    switchedVlan.EntityData.BundleName = "openconfig"
    switchedVlan.EntityData.ParentYangName = "aggregation"
    switchedVlan.EntityData.SegmentPath = "openconfig-vlan:switched-vlan"
    switchedVlan.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    switchedVlan.EntityData.NamespaceTable = openconfig.GetNamespaces()
    switchedVlan.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    switchedVlan.EntityData.Children = types.NewOrderedMap()
    switchedVlan.EntityData.Children.Append("config", types.YChild{"Config", &switchedVlan.Config})
    switchedVlan.EntityData.Children.Append("state", types.YChild{"State", &switchedVlan.State})
    switchedVlan.EntityData.Leafs = types.NewOrderedMap()

    switchedVlan.EntityData.YListKeys = []string {}

    return &(switchedVlan.EntityData)
}

// Interfaces_Interface_Aggregation_SwitchedVlan_Config
// Configuration parameters for VLANs
type Interfaces_Interface_Aggregation_SwitchedVlan_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the interface to access or trunk mode for VLANs. The type is
    // VlanModeType.
    InterfaceMode interface{}

    // Set the native VLAN id for untagged frames arriving on a trunk interface. 
    // This configuration is only valid on a trunk interface. The type is one of
    // the following types: int with range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    NativeVlan interface{}

    // Assign the access vlan to the access port. The type is one of the following
    // types: int with range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    AccessVlan interface{}

    // Specify VLANs, or ranges thereof, that the interface may carry when in
    // trunk mode.  If not specified, all VLANs are allowed on the interface.
    // Ranges are specified in the form x..y, where x<y - ranges are assumed to be
    // inclusive (such that the VLAN range is x <= range <= y. The type is one of
    // the following types: slice of int with range: 1..4094, or slice of string
    // with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])$,
    // or slice of string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$,
    // or slice of string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$,
    // or slice of string with pattern:
    // ^(\*|(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]))\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])$..
    TrunkVlans []interface{}
}

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "switched-vlan"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("interface-mode", types.YLeaf{"InterfaceMode", config.InterfaceMode})
    config.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", config.NativeVlan})
    config.EntityData.Leafs.Append("access-vlan", types.YLeaf{"AccessVlan", config.AccessVlan})
    config.EntityData.Leafs.Append("trunk-vlans", types.YLeaf{"TrunkVlans", config.TrunkVlans})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_Aggregation_SwitchedVlan_State
// State variables for VLANs
type Interfaces_Interface_Aggregation_SwitchedVlan_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the interface to access or trunk mode for VLANs. The type is
    // VlanModeType.
    InterfaceMode interface{}

    // Set the native VLAN id for untagged frames arriving on a trunk interface. 
    // This configuration is only valid on a trunk interface. The type is one of
    // the following types: int with range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    NativeVlan interface{}

    // Assign the access vlan to the access port. The type is one of the following
    // types: int with range: 1..4094, or string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$.
    AccessVlan interface{}

    // Specify VLANs, or ranges thereof, that the interface may carry when in
    // trunk mode.  If not specified, all VLANs are allowed on the interface.
    // Ranges are specified in the form x..y, where x<y - ranges are assumed to be
    // inclusive (such that the VLAN range is x <= range <= y. The type is one of
    // the following types: slice of int with range: 1..4094, or slice of string
    // with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])$,
    // or slice of string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$,
    // or slice of string with pattern:
    // ^(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*)$,
    // or slice of string with pattern:
    // ^(\*|(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]))\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])$..
    TrunkVlans []interface{}
}

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "switched-vlan"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("interface-mode", types.YLeaf{"InterfaceMode", state.InterfaceMode})
    state.EntityData.Leafs.Append("native-vlan", types.YLeaf{"NativeVlan", state.NativeVlan})
    state.EntityData.Leafs.Append("access-vlan", types.YLeaf{"AccessVlan", state.AccessVlan})
    state.EntityData.Leafs.Append("trunk-vlans", types.YLeaf{"TrunkVlans", state.TrunkVlans})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan
// Top-level container for routed vlan interfaces.  These
// logical interfaces are also known as SVI (switched virtual
// interface), IRB (integrated routing and bridging), RVI
// (routed VLAN interface)
type Interfaces_Interface_RoutedVlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for routed vlan interfaces.
    Config Interfaces_Interface_RoutedVlan_Config

    // Operational state data .
    State Interfaces_Interface_RoutedVlan_State

    // Parameters for the IPv4 address family.
    Ipv4 Interfaces_Interface_RoutedVlan_Ipv4

    // Parameters for the IPv6 address family.
    Ipv6 Interfaces_Interface_RoutedVlan_Ipv6
}

func (routedVlan *Interfaces_Interface_RoutedVlan) GetEntityData() *types.CommonEntityData {
    routedVlan.EntityData.YFilter = routedVlan.YFilter
    routedVlan.EntityData.YangName = "routed-vlan"
    routedVlan.EntityData.BundleName = "openconfig"
    routedVlan.EntityData.ParentYangName = "interface"
    routedVlan.EntityData.SegmentPath = "openconfig-vlan:routed-vlan"
    routedVlan.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    routedVlan.EntityData.NamespaceTable = openconfig.GetNamespaces()
    routedVlan.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    routedVlan.EntityData.Children = types.NewOrderedMap()
    routedVlan.EntityData.Children.Append("config", types.YChild{"Config", &routedVlan.Config})
    routedVlan.EntityData.Children.Append("state", types.YChild{"State", &routedVlan.State})
    routedVlan.EntityData.Children.Append("openconfig-if-ip:ipv4", types.YChild{"Ipv4", &routedVlan.Ipv4})
    routedVlan.EntityData.Children.Append("openconfig-if-ip:ipv6", types.YChild{"Ipv6", &routedVlan.Ipv6})
    routedVlan.EntityData.Leafs = types.NewOrderedMap()

    routedVlan.EntityData.YListKeys = []string {}

    return &(routedVlan.EntityData)
}

// Interfaces_Interface_RoutedVlan_Config
// Configuration data for routed vlan interfaces
type Interfaces_Interface_RoutedVlan_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References the VLAN for which this IP interface provides routing services
    // -- similar to a switch virtual interface (SVI), or integrated routing and
    // bridging interface (IRB) in some implementations. The type is one of the
    // following types: int with range: 0..65535, or string.
    Vlan interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "routed-vlan"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("vlan", types.YLeaf{"Vlan", config.Vlan})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_State
// Operational state data 
type Interfaces_Interface_RoutedVlan_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // References the VLAN for which this IP interface provides routing services
    // -- similar to a switch virtual interface (SVI), or integrated routing and
    // bridging interface (IRB) in some implementations. The type is one of the
    // following types: int with range: 0..65535, or string.
    Vlan interface{}
}

func (state *Interfaces_Interface_RoutedVlan_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "routed-vlan"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("vlan", types.YLeaf{"Vlan", state.Vlan})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4
// Parameters for the IPv4 address family.
type Interfaces_Interface_RoutedVlan_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for address list.
    Addresses Interfaces_Interface_RoutedVlan_Ipv4_Addresses

    // Enclosing container for neighbor list.
    Neighbors Interfaces_Interface_RoutedVlan_Ipv4_Neighbors

    // Top-level container for setting unnumbered interfaces. Includes reference
    // the interface that provides the address information.
    Unnumbered Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered

    // Top-level IPv4 configuration data for the interface.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Config

    // Top level IPv4 operational state data.
    State Interfaces_Interface_RoutedVlan_Ipv4_State
}

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "openconfig"
    ipv4.EntityData.ParentYangName = "routed-vlan"
    ipv4.EntityData.SegmentPath = "openconfig-if-ip:ipv4"
    ipv4.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv4.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("addresses", types.YChild{"Addresses", &ipv4.Addresses})
    ipv4.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &ipv4.Neighbors})
    ipv4.EntityData.Children.Append("unnumbered", types.YChild{"Unnumbered", &ipv4.Unnumbered})
    ipv4.EntityData.Children.Append("config", types.YChild{"Config", &ipv4.Config})
    ipv4.EntityData.Children.Append("state", types.YChild{"State", &ipv4.State})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses
// Enclosing container for address list
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of configured IPv4 addresses on the interface. The type is slice
    // of Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address.
    Address []*Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "openconfig"
    addresses.EntityData.ParentYangName = "ipv4"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    addresses.EntityData.NamespaceTable = openconfig.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range addresses.Address {
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.Address[i]), types.YChild{"Address", addresses.Address[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address
// The list of configured IPv4 addresses on the interface.
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured IP address. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // interfaces.Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config_Ip
    Ip interface{}

    // Configuration data for each configured IPv4 address on the interface.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config

    // Operational state data for each IPv4 address configured on the interface.
    State Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State

    // Enclosing container for VRRP groups handled by this IP interface.
    Vrrp Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp
}

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "openconfig"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + types.AddKeyToken(address.Ip, "ip")
    address.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    address.EntityData.NamespaceTable = openconfig.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("config", types.YChild{"Config", &address.Config})
    address.EntityData.Children.Append("state", types.YChild{"State", &address.State})
    address.EntityData.Children.Append("vrrp", types.YChild{"Vrrp", &address.Vrrp})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", address.Ip})

    address.EntityData.YListKeys = []string {"Ip"}

    return &(address.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config
// Configuration data for each configured IPv4
// address on the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv4 address on the interface.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..32.
    PrefixLength interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "address"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", config.Ip})
    config.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", config.PrefixLength})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State
// Operational state data for each IPv4 address
// configured on the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv4 address on the interface.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..32.
    PrefixLength interface{}

    // The origin of this address, e.g., statically configured, assigned by DHCP,
    // etc.. The type is IpAddressOrigin.
    Origin interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "address"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", state.Ip})
    state.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", state.PrefixLength})
    state.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", state.Origin})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp
// Enclosing container for VRRP groups handled by this
// IP interface
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of VRRP groups, keyed by virtual router id. The type is slice of
    // Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup.
    VrrpGroup []*Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetEntityData() *types.CommonEntityData {
    vrrp.EntityData.YFilter = vrrp.YFilter
    vrrp.EntityData.YangName = "vrrp"
    vrrp.EntityData.BundleName = "openconfig"
    vrrp.EntityData.ParentYangName = "address"
    vrrp.EntityData.SegmentPath = "vrrp"
    vrrp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vrrp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vrrp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vrrp.EntityData.Children = types.NewOrderedMap()
    vrrp.EntityData.Children.Append("vrrp-group", types.YChild{"VrrpGroup", nil})
    for i := range vrrp.VrrpGroup {
        vrrp.EntityData.Children.Append(types.GetSegmentPath(vrrp.VrrpGroup[i]), types.YChild{"VrrpGroup", vrrp.VrrpGroup[i]})
    }
    vrrp.EntityData.Leafs = types.NewOrderedMap()

    vrrp.EntityData.YListKeys = []string {}

    return &(vrrp.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup
// List of VRRP groups, keyed by virtual router id
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured virtual router id for
    // this VRRP group. The type is string with range: 1..255. Refers to
    // interfaces.Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config_VirtualRouterId
    VirtualRouterId interface{}

    // Configuration data for the VRRP group.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config

    // Operational state data for the VRRP group.
    State Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State

    // Top-level container for VRRP interface tracking.
    InterfaceTracking Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetEntityData() *types.CommonEntityData {
    vrrpGroup.EntityData.YFilter = vrrpGroup.YFilter
    vrrpGroup.EntityData.YangName = "vrrp-group"
    vrrpGroup.EntityData.BundleName = "openconfig"
    vrrpGroup.EntityData.ParentYangName = "vrrp"
    vrrpGroup.EntityData.SegmentPath = "vrrp-group" + types.AddKeyToken(vrrpGroup.VirtualRouterId, "virtual-router-id")
    vrrpGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vrrpGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vrrpGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vrrpGroup.EntityData.Children = types.NewOrderedMap()
    vrrpGroup.EntityData.Children.Append("config", types.YChild{"Config", &vrrpGroup.Config})
    vrrpGroup.EntityData.Children.Append("state", types.YChild{"State", &vrrpGroup.State})
    vrrpGroup.EntityData.Children.Append("interface-tracking", types.YChild{"InterfaceTracking", &vrrpGroup.InterfaceTracking})
    vrrpGroup.EntityData.Leafs = types.NewOrderedMap()
    vrrpGroup.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", vrrpGroup.VirtualRouterId})

    vrrpGroup.EntityData.YListKeys = []string {"VirtualRouterId"}

    return &(vrrpGroup.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config
// Configuration data for the VRRP group
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the virtual router id for use by the VRRP group.  This usually also
    // determines the virtual MAC address that is generated for the VRRP group.
    // The type is interface{} with range: 1..255.
    VirtualRouterId interface{}

    // Configure one or more virtual addresses for the VRRP group. The type is one
    // of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualAddress []interface{}

    // Specifies the sending VRRP interface's priority for the virtual router. 
    // Higher values equal higher priority. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // When set to true, enables preemption by a higher priority backup router of
    // a lower priority master router. The type is bool. The default value is
    // true.
    Preempt interface{}

    // Set the delay the higher priority router waits before preempting. The type
    // is interface{} with range: 0..3600. The default value is 0.
    PreemptDelay interface{}

    // Configure whether packets destined for virtual addresses are accepted even
    // when the virtual address is not owned by the router interface. The type is
    // bool. The default value is false.
    AcceptMode interface{}

    // Sets the interval between successive VRRP advertisements -- RFC 5798
    // defines this as a 12-bit value expressed as 0.1 seconds, with default 100,
    // i.e., 1 second.  Several implementation express this in units of seconds.
    // The type is interface{} with range: 1..4095. Units are centiseconds. The
    // default value is 100.
    AdvertisementInterval interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "vrrp-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", config.VirtualRouterId})
    config.EntityData.Leafs.Append("virtual-address", types.YLeaf{"VirtualAddress", config.VirtualAddress})
    config.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", config.Priority})
    config.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", config.Preempt})
    config.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", config.PreemptDelay})
    config.EntityData.Leafs.Append("accept-mode", types.YLeaf{"AcceptMode", config.AcceptMode})
    config.EntityData.Leafs.Append("advertisement-interval", types.YLeaf{"AdvertisementInterval", config.AdvertisementInterval})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State
// Operational state data for the VRRP group
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the virtual router id for use by the VRRP group.  This usually also
    // determines the virtual MAC address that is generated for the VRRP group.
    // The type is interface{} with range: 1..255.
    VirtualRouterId interface{}

    // Configure one or more virtual addresses for the VRRP group. The type is one
    // of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualAddress []interface{}

    // Specifies the sending VRRP interface's priority for the virtual router. 
    // Higher values equal higher priority. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // When set to true, enables preemption by a higher priority backup router of
    // a lower priority master router. The type is bool. The default value is
    // true.
    Preempt interface{}

    // Set the delay the higher priority router waits before preempting. The type
    // is interface{} with range: 0..3600. The default value is 0.
    PreemptDelay interface{}

    // Configure whether packets destined for virtual addresses are accepted even
    // when the virtual address is not owned by the router interface. The type is
    // bool. The default value is false.
    AcceptMode interface{}

    // Sets the interval between successive VRRP advertisements -- RFC 5798
    // defines this as a 12-bit value expressed as 0.1 seconds, with default 100,
    // i.e., 1 second.  Several implementation express this in units of seconds.
    // The type is interface{} with range: 1..4095. Units are centiseconds. The
    // default value is 100.
    AdvertisementInterval interface{}

    // Operational value of the priority for the interface in the VRRP group. The
    // type is interface{} with range: 0..255.
    CurrentPriority interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "vrrp-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", state.VirtualRouterId})
    state.EntityData.Leafs.Append("virtual-address", types.YLeaf{"VirtualAddress", state.VirtualAddress})
    state.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", state.Priority})
    state.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", state.Preempt})
    state.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", state.PreemptDelay})
    state.EntityData.Leafs.Append("accept-mode", types.YLeaf{"AcceptMode", state.AcceptMode})
    state.EntityData.Leafs.Append("advertisement-interval", types.YLeaf{"AdvertisementInterval", state.AdvertisementInterval})
    state.EntityData.Leafs.Append("current-priority", types.YLeaf{"CurrentPriority", state.CurrentPriority})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
// Top-level container for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for VRRP interface tracking.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config

    // Operational state data for VRRP interface tracking.
    State Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetEntityData() *types.CommonEntityData {
    interfaceTracking.EntityData.YFilter = interfaceTracking.YFilter
    interfaceTracking.EntityData.YangName = "interface-tracking"
    interfaceTracking.EntityData.BundleName = "openconfig"
    interfaceTracking.EntityData.ParentYangName = "vrrp-group"
    interfaceTracking.EntityData.SegmentPath = "interface-tracking"
    interfaceTracking.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceTracking.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceTracking.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceTracking.EntityData.Children = types.NewOrderedMap()
    interfaceTracking.EntityData.Children.Append("config", types.YChild{"Config", &interfaceTracking.Config})
    interfaceTracking.EntityData.Children.Append("state", types.YChild{"State", &interfaceTracking.State})
    interfaceTracking.EntityData.Leafs = types.NewOrderedMap()

    interfaceTracking.EntityData.YListKeys = []string {}

    return &(interfaceTracking.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config
// Configuration data for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets an interface that should be tracked for up/down events to dynamically
    // change the priority state of the VRRP group, and potentially change the
    // mastership if the tracked interface going down lowers the priority
    // sufficiently. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    TrackInterface interface{}

    // Set the value to subtract from priority when the tracked interface goes
    // down. The type is interface{} with range: 0..254. The default value is 0.
    PriorityDecrement interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-tracking"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("track-interface", types.YLeaf{"TrackInterface", config.TrackInterface})
    config.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", config.PriorityDecrement})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
// Operational state data for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets an interface that should be tracked for up/down events to dynamically
    // change the priority state of the VRRP group, and potentially change the
    // mastership if the tracked interface going down lowers the priority
    // sufficiently. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    TrackInterface interface{}

    // Set the value to subtract from priority when the tracked interface goes
    // down. The type is interface{} with range: 0..254. The default value is 0.
    PriorityDecrement interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-tracking"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("track-interface", types.YLeaf{"TrackInterface", state.TrackInterface})
    state.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", state.PriorityDecrement})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Neighbors
// Enclosing container for neighbor list
type Interfaces_Interface_RoutedVlan_Ipv4_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of mappings from IPv4 addresses to link-layer addresses.  Entries in
    // this list are used as static entries in the ARP Cache. The type is slice of
    // Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor.
    Neighbor []*Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "ipv4"
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

// Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor
// A list of mappings from IPv4 addresses to
// link-layer addresses.
// 
// Entries in this list are used as static entries in the
// ARP Cache.
type Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured IP address. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // Refers to
    // interfaces.Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config_Ip
    Ip interface{}

    // Configuration data for each configured IPv4 address on the interface.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config

    // Operational state data for each IPv4 address configured on the interface.
    State Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.Ip, "ip")
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("config", types.YChild{"Config", &neighbor.Config})
    neighbor.EntityData.Children.Append("state", types.YChild{"State", &neighbor.State})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", neighbor.Ip})

    neighbor.EntityData.YListKeys = []string {"Ip"}

    return &(neighbor.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config
// Configuration data for each configured IPv4
// address on the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IPv4 address of the neighbor node. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetEntityData() *types.CommonEntityData {
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
    config.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", config.Ip})
    config.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", config.LinkLayerAddress})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State
// Operational state data for each IPv4 address
// configured on the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IPv4 address of the neighbor node. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}

    // The origin of this neighbor entry, static or dynamic. The type is
    // NeighborOrigin.
    Origin interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetEntityData() *types.CommonEntityData {
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
    state.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", state.Ip})
    state.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", state.LinkLayerAddress})
    state.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", state.Origin})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered
// Top-level container for setting unnumbered interfaces.
// Includes reference the interface that provides the
// address information
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for unnumbered interface.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config

    // Operational state data for unnumbered interfaces.
    State Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State

    // Reference to an interface or subinterface.
    InterfaceRef Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetEntityData() *types.CommonEntityData {
    unnumbered.EntityData.YFilter = unnumbered.YFilter
    unnumbered.EntityData.YangName = "unnumbered"
    unnumbered.EntityData.BundleName = "openconfig"
    unnumbered.EntityData.ParentYangName = "ipv4"
    unnumbered.EntityData.SegmentPath = "unnumbered"
    unnumbered.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unnumbered.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unnumbered.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unnumbered.EntityData.Children = types.NewOrderedMap()
    unnumbered.EntityData.Children.Append("config", types.YChild{"Config", &unnumbered.Config})
    unnumbered.EntityData.Children.Append("state", types.YChild{"State", &unnumbered.State})
    unnumbered.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &unnumbered.InterfaceRef})
    unnumbered.EntityData.Leafs = types.NewOrderedMap()

    unnumbered.EntityData.YListKeys = []string {}

    return &(unnumbered.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config
// Configuration data for unnumbered interface
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "unnumbered"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State
// Operational state data for unnumbered interfaces
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "unnumbered"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
// Reference to an interface or subinterface
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config

    // Operational state for interface-ref.
    State Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "unnumbered"
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

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config
// Configured reference to interface / subinterface
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config struct {
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

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
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

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State
// Operational state for interface-ref
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State struct {
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

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
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

// Interfaces_Interface_RoutedVlan_Ipv4_Config
// Top-level IPv4 configuration data for the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controls whether IPv4 is enabled or disabled on this interface.  When IPv4
    // is enabled, this interface is connected to an IPv4 stack, and the interface
    // can send and receive IPv4 packets. The type is bool. The default value is
    // true.
    Enabled interface{}

    // The size, in octets, of the largest IPv4 packet that the interface will
    // send and receive.  The server may restrict the allowed values for this
    // leaf, depending on the interface's type.  If this leaf is not configured,
    // the operationally used MTU depends on the interface's type. The type is
    // interface{} with range: 68..65535. Units are octets.
    Mtu interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv4"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", config.Mtu})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv4_State
// Top level IPv4 operational state data
type Interfaces_Interface_RoutedVlan_Ipv4_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Controls whether IPv4 is enabled or disabled on this interface.  When IPv4
    // is enabled, this interface is connected to an IPv4 stack, and the interface
    // can send and receive IPv4 packets. The type is bool. The default value is
    // true.
    Enabled interface{}

    // The size, in octets, of the largest IPv4 packet that the interface will
    // send and receive.  The server may restrict the allowed values for this
    // leaf, depending on the interface's type.  If this leaf is not configured,
    // the operationally used MTU depends on the interface's type. The type is
    // interface{} with range: 68..65535. Units are octets.
    Mtu interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv4"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", state.Mtu})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6
// Parameters for the IPv6 address family.
type Interfaces_Interface_RoutedVlan_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enclosing container for address list.
    Addresses Interfaces_Interface_RoutedVlan_Ipv6_Addresses

    // Enclosing container for list of IPv6 neighbors.
    Neighbors Interfaces_Interface_RoutedVlan_Ipv6_Neighbors

    // Top-level container for setting unnumbered interfaces. Includes reference
    // the interface that provides the address information.
    Unnumbered Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered

    // Top-level config data for the IPv6 interface.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Config

    // Top-level operational state data for the IPv6 interface.
    State Interfaces_Interface_RoutedVlan_Ipv6_State
}

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "openconfig"
    ipv6.EntityData.ParentYangName = "routed-vlan"
    ipv6.EntityData.SegmentPath = "openconfig-if-ip:ipv6"
    ipv6.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    ipv6.EntityData.NamespaceTable = openconfig.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("addresses", types.YChild{"Addresses", &ipv6.Addresses})
    ipv6.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &ipv6.Neighbors})
    ipv6.EntityData.Children.Append("unnumbered", types.YChild{"Unnumbered", &ipv6.Unnumbered})
    ipv6.EntityData.Children.Append("config", types.YChild{"Config", &ipv6.Config})
    ipv6.EntityData.Children.Append("state", types.YChild{"State", &ipv6.State})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses
// Enclosing container for address list
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of configured IPv6 addresses on the interface. The type is slice
    // of Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address.
    Address []*Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "openconfig"
    addresses.EntityData.ParentYangName = "ipv6"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    addresses.EntityData.NamespaceTable = openconfig.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range addresses.Address {
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.Address[i]), types.YChild{"Address", addresses.Address[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address
// The list of configured IPv6 addresses on the interface.
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured IP address. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // Refers to
    // interfaces.Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config_Ip
    Ip interface{}

    // Configuration data for each IPv6 address on the interface.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config

    // State data for each IPv6 address on the interface.
    State Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State

    // Enclosing container for VRRP groups handled by this IP interface.
    Vrrp Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp
}

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "openconfig"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + types.AddKeyToken(address.Ip, "ip")
    address.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    address.EntityData.NamespaceTable = openconfig.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("config", types.YChild{"Config", &address.Config})
    address.EntityData.Children.Append("state", types.YChild{"State", &address.State})
    address.EntityData.Children.Append("vrrp", types.YChild{"Vrrp", &address.Vrrp})
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", address.Ip})

    address.EntityData.YListKeys = []string {"Ip"}

    return &(address.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config
// Configuration data for each IPv6 address on
// the interface
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address on the interface.
    // The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..128. This attribute is mandatory.
    PrefixLength interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "address"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", config.Ip})
    config.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", config.PrefixLength})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State
// State data for each IPv6 address on the
// interface
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address on the interface.
    // The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..128. This attribute is mandatory.
    PrefixLength interface{}

    // [adapted from IETF IP model RFC 7277]  The origin of this address, e.g.,
    // static, dhcp, etc. The type is IpAddressOrigin.
    Origin interface{}

    // [adapted from IETF IP model RFC 7277]  The status of an address.  Most of
    // the states correspond to states from the IPv6 Stateless Address
    // Autoconfiguration protocol. The type is Status.
    Status interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "address"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", state.Ip})
    state.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", state.PrefixLength})
    state.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", state.Origin})
    state.EntityData.Leafs.Append("status", types.YLeaf{"Status", state.Status})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status represents Autoconfiguration protocol.
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status string

const (
    // This is a valid address that can appear as the
    // destination or source address of a packet.
    Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status_PREFERRED Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status = "PREFERRED"

    // This is a valid but deprecated address that should
    // no longer be used as a source address in new
    // communications, but packets addressed to such an
    // address are processed as expected.
    Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status_DEPRECATED Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status = "DEPRECATED"

    // This isn't a valid address, and it shouldn't appear
    // as the destination or source address of a packet.
    Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status_INVALID Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status = "INVALID"

    // The address is not accessible because the interface
    // to which this address is assigned is not
    // operational.
    Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status_INACCESSIBLE Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status = "INACCESSIBLE"

    // The status cannot be determined for some reason.
    Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status_UNKNOWN Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status = "UNKNOWN"

    // The uniqueness of the address on the link is being
    // verified.  Addresses in this state should not be
    // used for general communication and should only be
    // used to determine the uniqueness of the address.
    Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status_TENTATIVE Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status = "TENTATIVE"

    // The address has been determined to be non-unique on
    // the link and so must not be used.
    Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status_DUPLICATE Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status = "DUPLICATE"

    // The address is available for use, subject to
    // restrictions, while its uniqueness on a link is
    // being verified.
    Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status_OPTIMISTIC Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State_Status = "OPTIMISTIC"
)

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp
// Enclosing container for VRRP groups handled by this
// IP interface
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of VRRP groups, keyed by virtual router id. The type is slice of
    // Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup.
    VrrpGroup []*Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetEntityData() *types.CommonEntityData {
    vrrp.EntityData.YFilter = vrrp.YFilter
    vrrp.EntityData.YangName = "vrrp"
    vrrp.EntityData.BundleName = "openconfig"
    vrrp.EntityData.ParentYangName = "address"
    vrrp.EntityData.SegmentPath = "vrrp"
    vrrp.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vrrp.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vrrp.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vrrp.EntityData.Children = types.NewOrderedMap()
    vrrp.EntityData.Children.Append("vrrp-group", types.YChild{"VrrpGroup", nil})
    for i := range vrrp.VrrpGroup {
        vrrp.EntityData.Children.Append(types.GetSegmentPath(vrrp.VrrpGroup[i]), types.YChild{"VrrpGroup", vrrp.VrrpGroup[i]})
    }
    vrrp.EntityData.Leafs = types.NewOrderedMap()

    vrrp.EntityData.YListKeys = []string {}

    return &(vrrp.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup
// List of VRRP groups, keyed by virtual router id
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured virtual router id for
    // this VRRP group. The type is string with range: 1..255. Refers to
    // interfaces.Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config_VirtualRouterId
    VirtualRouterId interface{}

    // Configuration data for the VRRP group.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config

    // Operational state data for the VRRP group.
    State Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State

    // Top-level container for VRRP interface tracking.
    InterfaceTracking Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetEntityData() *types.CommonEntityData {
    vrrpGroup.EntityData.YFilter = vrrpGroup.YFilter
    vrrpGroup.EntityData.YangName = "vrrp-group"
    vrrpGroup.EntityData.BundleName = "openconfig"
    vrrpGroup.EntityData.ParentYangName = "vrrp"
    vrrpGroup.EntityData.SegmentPath = "vrrp-group" + types.AddKeyToken(vrrpGroup.VirtualRouterId, "virtual-router-id")
    vrrpGroup.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    vrrpGroup.EntityData.NamespaceTable = openconfig.GetNamespaces()
    vrrpGroup.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    vrrpGroup.EntityData.Children = types.NewOrderedMap()
    vrrpGroup.EntityData.Children.Append("config", types.YChild{"Config", &vrrpGroup.Config})
    vrrpGroup.EntityData.Children.Append("state", types.YChild{"State", &vrrpGroup.State})
    vrrpGroup.EntityData.Children.Append("interface-tracking", types.YChild{"InterfaceTracking", &vrrpGroup.InterfaceTracking})
    vrrpGroup.EntityData.Leafs = types.NewOrderedMap()
    vrrpGroup.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", vrrpGroup.VirtualRouterId})

    vrrpGroup.EntityData.YListKeys = []string {"VirtualRouterId"}

    return &(vrrpGroup.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config
// Configuration data for the VRRP group
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the virtual router id for use by the VRRP group.  This usually also
    // determines the virtual MAC address that is generated for the VRRP group.
    // The type is interface{} with range: 1..255.
    VirtualRouterId interface{}

    // Configure one or more virtual addresses for the VRRP group. The type is one
    // of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualAddress []interface{}

    // Specifies the sending VRRP interface's priority for the virtual router. 
    // Higher values equal higher priority. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // When set to true, enables preemption by a higher priority backup router of
    // a lower priority master router. The type is bool. The default value is
    // true.
    Preempt interface{}

    // Set the delay the higher priority router waits before preempting. The type
    // is interface{} with range: 0..3600. The default value is 0.
    PreemptDelay interface{}

    // Configure whether packets destined for virtual addresses are accepted even
    // when the virtual address is not owned by the router interface. The type is
    // bool. The default value is false.
    AcceptMode interface{}

    // Sets the interval between successive VRRP advertisements -- RFC 5798
    // defines this as a 12-bit value expressed as 0.1 seconds, with default 100,
    // i.e., 1 second.  Several implementation express this in units of seconds.
    // The type is interface{} with range: 1..4095. Units are centiseconds. The
    // default value is 100.
    AdvertisementInterval interface{}

    // For VRRP on IPv6 interfaces, sets the virtual link local address. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualLinkLocal interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "vrrp-group"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", config.VirtualRouterId})
    config.EntityData.Leafs.Append("virtual-address", types.YLeaf{"VirtualAddress", config.VirtualAddress})
    config.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", config.Priority})
    config.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", config.Preempt})
    config.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", config.PreemptDelay})
    config.EntityData.Leafs.Append("accept-mode", types.YLeaf{"AcceptMode", config.AcceptMode})
    config.EntityData.Leafs.Append("advertisement-interval", types.YLeaf{"AdvertisementInterval", config.AdvertisementInterval})
    config.EntityData.Leafs.Append("virtual-link-local", types.YLeaf{"VirtualLinkLocal", config.VirtualLinkLocal})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State
// Operational state data for the VRRP group
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the virtual router id for use by the VRRP group.  This usually also
    // determines the virtual MAC address that is generated for the VRRP group.
    // The type is interface{} with range: 1..255.
    VirtualRouterId interface{}

    // Configure one or more virtual addresses for the VRRP group. The type is one
    // of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualAddress []interface{}

    // Specifies the sending VRRP interface's priority for the virtual router. 
    // Higher values equal higher priority. The type is interface{} with range:
    // 1..254. The default value is 100.
    Priority interface{}

    // When set to true, enables preemption by a higher priority backup router of
    // a lower priority master router. The type is bool. The default value is
    // true.
    Preempt interface{}

    // Set the delay the higher priority router waits before preempting. The type
    // is interface{} with range: 0..3600. The default value is 0.
    PreemptDelay interface{}

    // Configure whether packets destined for virtual addresses are accepted even
    // when the virtual address is not owned by the router interface. The type is
    // bool. The default value is false.
    AcceptMode interface{}

    // Sets the interval between successive VRRP advertisements -- RFC 5798
    // defines this as a 12-bit value expressed as 0.1 seconds, with default 100,
    // i.e., 1 second.  Several implementation express this in units of seconds.
    // The type is interface{} with range: 1..4095. Units are centiseconds. The
    // default value is 100.
    AdvertisementInterval interface{}

    // Operational value of the priority for the interface in the VRRP group. The
    // type is interface{} with range: 0..255.
    CurrentPriority interface{}

    // For VRRP on IPv6 interfaces, sets the virtual link local address. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualLinkLocal interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "vrrp-group"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("virtual-router-id", types.YLeaf{"VirtualRouterId", state.VirtualRouterId})
    state.EntityData.Leafs.Append("virtual-address", types.YLeaf{"VirtualAddress", state.VirtualAddress})
    state.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", state.Priority})
    state.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", state.Preempt})
    state.EntityData.Leafs.Append("preempt-delay", types.YLeaf{"PreemptDelay", state.PreemptDelay})
    state.EntityData.Leafs.Append("accept-mode", types.YLeaf{"AcceptMode", state.AcceptMode})
    state.EntityData.Leafs.Append("advertisement-interval", types.YLeaf{"AdvertisementInterval", state.AdvertisementInterval})
    state.EntityData.Leafs.Append("current-priority", types.YLeaf{"CurrentPriority", state.CurrentPriority})
    state.EntityData.Leafs.Append("virtual-link-local", types.YLeaf{"VirtualLinkLocal", state.VirtualLinkLocal})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
// Top-level container for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for VRRP interface tracking.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config

    // Operational state data for VRRP interface tracking.
    State Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetEntityData() *types.CommonEntityData {
    interfaceTracking.EntityData.YFilter = interfaceTracking.YFilter
    interfaceTracking.EntityData.YangName = "interface-tracking"
    interfaceTracking.EntityData.BundleName = "openconfig"
    interfaceTracking.EntityData.ParentYangName = "vrrp-group"
    interfaceTracking.EntityData.SegmentPath = "interface-tracking"
    interfaceTracking.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    interfaceTracking.EntityData.NamespaceTable = openconfig.GetNamespaces()
    interfaceTracking.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    interfaceTracking.EntityData.Children = types.NewOrderedMap()
    interfaceTracking.EntityData.Children.Append("config", types.YChild{"Config", &interfaceTracking.Config})
    interfaceTracking.EntityData.Children.Append("state", types.YChild{"State", &interfaceTracking.State})
    interfaceTracking.EntityData.Leafs = types.NewOrderedMap()

    interfaceTracking.EntityData.YListKeys = []string {}

    return &(interfaceTracking.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config
// Configuration data for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets an interface that should be tracked for up/down events to dynamically
    // change the priority state of the VRRP group, and potentially change the
    // mastership if the tracked interface going down lowers the priority
    // sufficiently. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    TrackInterface interface{}

    // Set the value to subtract from priority when the tracked interface goes
    // down. The type is interface{} with range: 0..254. The default value is 0.
    PriorityDecrement interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "interface-tracking"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("track-interface", types.YLeaf{"TrackInterface", config.TrackInterface})
    config.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", config.PriorityDecrement})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
// Operational state data for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sets an interface that should be tracked for up/down events to dynamically
    // change the priority state of the VRRP group, and potentially change the
    // mastership if the tracked interface going down lowers the priority
    // sufficiently. The type is string. Refers to
    // interfaces.Interfaces_Interface_Name
    TrackInterface interface{}

    // Set the value to subtract from priority when the tracked interface goes
    // down. The type is interface{} with range: 0..254. The default value is 0.
    PriorityDecrement interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "interface-tracking"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("track-interface", types.YLeaf{"TrackInterface", state.TrackInterface})
    state.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", state.PriorityDecrement})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors
// Enclosing container for list of IPv6 neighbors
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of IPv6 neighbors. The type is slice of
    // Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor.
    Neighbor []*Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "openconfig"
    neighbors.EntityData.ParentYangName = "ipv6"
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

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor
// List of IPv6 neighbors
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. References the configured IP neighbor address. The
    // type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // Refers to
    // interfaces.Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config_Ip
    Ip interface{}

    // Configuration data for each IPv6 address on the interface.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config

    // State data for each IPv6 address on the interface.
    State Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "openconfig"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.Ip, "ip")
    neighbor.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    neighbor.EntityData.NamespaceTable = openconfig.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("config", types.YChild{"Config", &neighbor.Config})
    neighbor.EntityData.Children.Append("state", types.YChild{"State", &neighbor.State})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", neighbor.Ip})

    neighbor.EntityData.YListKeys = []string {"Ip"}

    return &(neighbor.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config
// Configuration data for each IPv6 address on
// the interface
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address of the neighbor
    // node. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The link-layer address of the
    // neighbor node. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetEntityData() *types.CommonEntityData {
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
    config.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", config.Ip})
    config.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", config.LinkLayerAddress})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State
// State data for each IPv6 address on the
// interface
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address of the neighbor
    // node. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The link-layer address of the
    // neighbor node. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}

    // [adapted from IETF IP model RFC 7277]  The origin of this neighbor entry.
    // The type is NeighborOrigin.
    Origin interface{}

    // [adapted from IETF IP model RFC 7277]  Indicates that the neighbor node
    // acts as a router. The type is interface{}.
    IsRouter interface{}

    // [adapted from IETF IP model RFC 7277]  The Neighbor Unreachability
    // Detection state of this entry. The type is NeighborState.
    NeighborState interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetEntityData() *types.CommonEntityData {
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
    state.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", state.Ip})
    state.EntityData.Leafs.Append("link-layer-address", types.YLeaf{"LinkLayerAddress", state.LinkLayerAddress})
    state.EntityData.Leafs.Append("origin", types.YLeaf{"Origin", state.Origin})
    state.EntityData.Leafs.Append("is-router", types.YLeaf{"IsRouter", state.IsRouter})
    state.EntityData.Leafs.Append("neighbor-state", types.YLeaf{"NeighborState", state.NeighborState})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState represents entry.
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState string

const (
    // Address resolution is in progress, and the link-layer
    //      address of the neighbor has not yet been
    //      determined.
    Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState_INCOMPLETE Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState = "INCOMPLETE"

    // Roughly speaking, the neighbor is known to have been
    //      reachable recently (within tens of seconds ago).
    Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState_REACHABLE Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState = "REACHABLE"

    // The neighbor is no longer known to be reachable, but
    //      until traffic is sent to the neighbor no attempt
    //      should be made to verify its reachability.
    Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState_STALE Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState = "STALE"

    // The neighbor is no longer known to be reachable, and
    //      traffic has recently been sent to the neighbor.
    //      Rather than probe the neighbor immediately, however,
    //      delay sending probes for a short while in order to
    //      give upper-layer protocols a chance to provide
    //      reachability confirmation.
    Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState_DELAY Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState = "DELAY"

    // The neighbor is no longer known to be reachable, and
    //      unicast Neighbor Solicitation probes are being sent
    //      to verify reachability.
    Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState_PROBE Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State_NeighborState = "PROBE"
)

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered
// Top-level container for setting unnumbered interfaces.
// Includes reference the interface that provides the
// address information
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration data for unnumbered interface.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config

    // Operational state data for unnumbered interfaces.
    State Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State

    // Reference to an interface or subinterface.
    InterfaceRef Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetEntityData() *types.CommonEntityData {
    unnumbered.EntityData.YFilter = unnumbered.YFilter
    unnumbered.EntityData.YangName = "unnumbered"
    unnumbered.EntityData.BundleName = "openconfig"
    unnumbered.EntityData.ParentYangName = "ipv6"
    unnumbered.EntityData.SegmentPath = "unnumbered"
    unnumbered.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    unnumbered.EntityData.NamespaceTable = openconfig.GetNamespaces()
    unnumbered.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    unnumbered.EntityData.Children = types.NewOrderedMap()
    unnumbered.EntityData.Children.Append("config", types.YChild{"Config", &unnumbered.Config})
    unnumbered.EntityData.Children.Append("state", types.YChild{"State", &unnumbered.State})
    unnumbered.EntityData.Children.Append("interface-ref", types.YChild{"InterfaceRef", &unnumbered.InterfaceRef})
    unnumbered.EntityData.Leafs = types.NewOrderedMap()

    unnumbered.EntityData.YListKeys = []string {}

    return &(unnumbered.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config
// Configuration data for unnumbered interface
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "unnumbered"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State
// Operational state data for unnumbered interfaces
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "unnumbered"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
// Reference to an interface or subinterface
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config

    // Operational state for interface-ref.
    State Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetEntityData() *types.CommonEntityData {
    interfaceRef.EntityData.YFilter = interfaceRef.YFilter
    interfaceRef.EntityData.YangName = "interface-ref"
    interfaceRef.EntityData.BundleName = "openconfig"
    interfaceRef.EntityData.ParentYangName = "unnumbered"
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

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config
// Configured reference to interface / subinterface
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config struct {
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

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetEntityData() *types.CommonEntityData {
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

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State
// Operational state for interface-ref
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State struct {
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

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetEntityData() *types.CommonEntityData {
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

// Interfaces_Interface_RoutedVlan_Ipv6_Config
// Top-level config data for the IPv6 interface
type Interfaces_Interface_RoutedVlan_Ipv6_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  Controls whether IPv6 is enabled or
    // disabled on this interface.  When IPv6 is enabled, this interface is
    // connected to an IPv6 stack, and the interface can send and receive IPv6
    // packets. The type is bool. The default value is true.
    Enabled interface{}

    // [adapted from IETF IP model RFC 7277]  The size, in octets, of the largest
    // IPv6 packet that the interface will send and receive.  The server may
    // restrict the allowed values for this leaf, depending on the interface's
    // type.  If this leaf is not configured, the operationally used MTU depends
    // on the interface's type. The type is interface{} with range:
    // 1280..4294967295. Units are octets.
    Mtu interface{}

    // [adapted from IETF IP model RFC 7277]  The number of consecutive Neighbor
    // Solicitation messages sent while performing Duplicate Address Detection on
    // a tentative address.  A value of zero indicates that Duplicate Address
    // Detection is not performed on tentative addresses.  A value of one
    // indicates a single transmission with no follow-up retransmissions. The type
    // is interface{} with range: 0..4294967295. The default value is 1.
    DupAddrDetectTransmits interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "openconfig"
    config.EntityData.ParentYangName = "ipv6"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    config.EntityData.NamespaceTable = openconfig.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", config.Enabled})
    config.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", config.Mtu})
    config.EntityData.Leafs.Append("dup-addr-detect-transmits", types.YLeaf{"DupAddrDetectTransmits", config.DupAddrDetectTransmits})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Interfaces_Interface_RoutedVlan_Ipv6_State
// Top-level operational state data for the IPv6 interface
type Interfaces_Interface_RoutedVlan_Ipv6_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  Controls whether IPv6 is enabled or
    // disabled on this interface.  When IPv6 is enabled, this interface is
    // connected to an IPv6 stack, and the interface can send and receive IPv6
    // packets. The type is bool. The default value is true.
    Enabled interface{}

    // [adapted from IETF IP model RFC 7277]  The size, in octets, of the largest
    // IPv6 packet that the interface will send and receive.  The server may
    // restrict the allowed values for this leaf, depending on the interface's
    // type.  If this leaf is not configured, the operationally used MTU depends
    // on the interface's type. The type is interface{} with range:
    // 1280..4294967295. Units are octets.
    Mtu interface{}

    // [adapted from IETF IP model RFC 7277]  The number of consecutive Neighbor
    // Solicitation messages sent while performing Duplicate Address Detection on
    // a tentative address.  A value of zero indicates that Duplicate Address
    // Detection is not performed on tentative addresses.  A value of one
    // indicates a single transmission with no follow-up retransmissions. The type
    // is interface{} with range: 0..4294967295. The default value is 1.
    DupAddrDetectTransmits interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "openconfig"
    state.EntityData.ParentYangName = "ipv6"
    state.EntityData.SegmentPath = "state"
    state.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    state.EntityData.NamespaceTable = openconfig.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", state.Enabled})
    state.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", state.Mtu})
    state.EntityData.Leafs.Append("dup-addr-detect-transmits", types.YLeaf{"DupAddrDetectTransmits", state.DupAddrDetectTransmits})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Interfaces_Interface_Sonet
// Data related to SONET/SDH interfaces
type Interfaces_Interface_Sonet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (sonet *Interfaces_Interface_Sonet) GetEntityData() *types.CommonEntityData {
    sonet.EntityData.YFilter = sonet.YFilter
    sonet.EntityData.YangName = "sonet"
    sonet.EntityData.BundleName = "openconfig"
    sonet.EntityData.ParentYangName = "interface"
    sonet.EntityData.SegmentPath = "openconfig-transport-line-common:sonet"
    sonet.EntityData.CapabilitiesTable = openconfig.GetCapabilities()
    sonet.EntityData.NamespaceTable = openconfig.GetNamespaces()
    sonet.EntityData.BundleYangModelsLocation = openconfig.GetModelsPath()

    sonet.EntityData.Children = types.NewOrderedMap()
    sonet.EntityData.Leafs = types.NewOrderedMap()

    sonet.EntityData.YListKeys = []string {}

    return &(sonet.EntityData)
}

