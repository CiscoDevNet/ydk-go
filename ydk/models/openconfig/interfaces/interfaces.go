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
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of named interfaces on the device. The type is slice of
    // Interfaces_Interface.
    Interface []Interfaces_Interface
}

func (interfaces *Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Interfaces) GetSegmentPath() string {
    return "openconfig-interfaces:interfaces"
}

func (interfaces *Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Interfaces) GetBundleName() string { return "openconfig" }

func (interfaces *Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Interfaces) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaces *Interfaces) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaces *Interfaces) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaces *Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Interfaces) GetParentYangName() string { return "openconfig-interfaces" }

// Interfaces_Interface
// The list of named interfaces on the device.
type Interfaces_Interface struct {
    parent types.Entity
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

func (self *Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "hold-time" { return "HoldTime" }
    if yname == "subinterfaces" { return "Subinterfaces" }
    if yname == "openconfig-if-ethernet:ethernet" { return "Ethernet" }
    if yname == "openconfig-if-aggregate:aggregation" { return "Aggregation" }
    if yname == "openconfig-vlan:routed-vlan" { return "RoutedVlan" }
    if yname == "openconfig-transport-line-common:sonet" { return "Sonet" }
    return ""
}

func (self *Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &self.Config
    }
    if childYangName == "state" {
        return &self.State
    }
    if childYangName == "hold-time" {
        return &self.HoldTime
    }
    if childYangName == "subinterfaces" {
        return &self.Subinterfaces
    }
    if childYangName == "openconfig-if-ethernet:ethernet" {
        return &self.Ethernet
    }
    if childYangName == "openconfig-if-aggregate:aggregation" {
        return &self.Aggregation
    }
    if childYangName == "openconfig-vlan:routed-vlan" {
        return &self.RoutedVlan
    }
    if childYangName == "openconfig-transport-line-common:sonet" {
        return &self.Sonet
    }
    return nil
}

func (self *Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &self.Config
    children["state"] = &self.State
    children["hold-time"] = &self.HoldTime
    children["subinterfaces"] = &self.Subinterfaces
    children["openconfig-if-ethernet:ethernet"] = &self.Ethernet
    children["openconfig-if-aggregate:aggregation"] = &self.Aggregation
    children["openconfig-vlan:routed-vlan"] = &self.RoutedVlan
    children["openconfig-transport-line-common:sonet"] = &self.Sonet
    return children
}

func (self *Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    return leafs
}

func (self *Interfaces_Interface) GetBundleName() string { return "openconfig" }

func (self *Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Interfaces_Interface) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (self *Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (self *Interfaces_Interface) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (self *Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Interfaces_Interface_Config
// Configurable items at the global, physical interface
// level
type Interfaces_Interface_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Config) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "mtu" { return "Mtu" }
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Interfaces_Interface_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = config.Type
    leafs["mtu"] = config.Mtu
    leafs["name"] = config.Name
    leafs["description"] = config.Description
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Interfaces_Interface_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Config) GetParentYangName() string { return "interface" }

// Interfaces_Interface_State
// Operational state data at the global interface level
type Interfaces_Interface_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_State) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "mtu" { return "Mtu" }
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "enabled" { return "Enabled" }
    if yname == "ifindex" { return "Ifindex" }
    if yname == "admin-status" { return "AdminStatus" }
    if yname == "oper-status" { return "OperStatus" }
    if yname == "last-change" { return "LastChange" }
    if yname == "hardware-port" { return "HardwarePort" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (state *Interfaces_Interface_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &state.Counters
    }
    return nil
}

func (state *Interfaces_Interface_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &state.Counters
    return children
}

func (state *Interfaces_Interface_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = state.Type
    leafs["mtu"] = state.Mtu
    leafs["name"] = state.Name
    leafs["description"] = state.Description
    leafs["enabled"] = state.Enabled
    leafs["ifindex"] = state.Ifindex
    leafs["admin-status"] = state.AdminStatus
    leafs["oper-status"] = state.OperStatus
    leafs["last-change"] = state.LastChange
    leafs["hardware-port"] = state.HardwarePort
    return leafs
}

func (state *Interfaces_Interface_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_State) GetParentYangName() string { return "interface" }

// Interfaces_Interface_State_Counters
// A collection of interface-related statistics objects.
type Interfaces_Interface_State_Counters struct {
    parent types.Entity
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

func (counters *Interfaces_Interface_State_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Interfaces_Interface_State_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Interfaces_Interface_State_Counters) GetGoName(yname string) string {
    if yname == "in-octets" { return "InOctets" }
    if yname == "in-unicast-pkts" { return "InUnicastPkts" }
    if yname == "in-broadcast-pkts" { return "InBroadcastPkts" }
    if yname == "in-multicast-pkts" { return "InMulticastPkts" }
    if yname == "in-discards" { return "InDiscards" }
    if yname == "in-errors" { return "InErrors" }
    if yname == "in-unknown-protos" { return "InUnknownProtos" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "out-unicast-pkts" { return "OutUnicastPkts" }
    if yname == "out-broadcast-pkts" { return "OutBroadcastPkts" }
    if yname == "out-multicast-pkts" { return "OutMulticastPkts" }
    if yname == "out-discards" { return "OutDiscards" }
    if yname == "out-errors" { return "OutErrors" }
    if yname == "last-clear" { return "LastClear" }
    return ""
}

func (counters *Interfaces_Interface_State_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Interfaces_Interface_State_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Interfaces_Interface_State_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Interfaces_Interface_State_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-octets"] = counters.InOctets
    leafs["in-unicast-pkts"] = counters.InUnicastPkts
    leafs["in-broadcast-pkts"] = counters.InBroadcastPkts
    leafs["in-multicast-pkts"] = counters.InMulticastPkts
    leafs["in-discards"] = counters.InDiscards
    leafs["in-errors"] = counters.InErrors
    leafs["in-unknown-protos"] = counters.InUnknownProtos
    leafs["out-octets"] = counters.OutOctets
    leafs["out-unicast-pkts"] = counters.OutUnicastPkts
    leafs["out-broadcast-pkts"] = counters.OutBroadcastPkts
    leafs["out-multicast-pkts"] = counters.OutMulticastPkts
    leafs["out-discards"] = counters.OutDiscards
    leafs["out-errors"] = counters.OutErrors
    leafs["last-clear"] = counters.LastClear
    return leafs
}

func (counters *Interfaces_Interface_State_Counters) GetBundleName() string { return "openconfig" }

func (counters *Interfaces_Interface_State_Counters) GetYangName() string { return "counters" }

func (counters *Interfaces_Interface_State_Counters) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (counters *Interfaces_Interface_State_Counters) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (counters *Interfaces_Interface_State_Counters) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (counters *Interfaces_Interface_State_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Interfaces_Interface_State_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Interfaces_Interface_State_Counters) GetParentYangName() string { return "state" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for interface hold-time settings.
    Config Interfaces_Interface_HoldTime_Config

    // Operational state data for interface hold-time.
    State Interfaces_Interface_HoldTime_State
}

func (holdTime *Interfaces_Interface_HoldTime) GetFilter() yfilter.YFilter { return holdTime.YFilter }

func (holdTime *Interfaces_Interface_HoldTime) SetFilter(yf yfilter.YFilter) { holdTime.YFilter = yf }

func (holdTime *Interfaces_Interface_HoldTime) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (holdTime *Interfaces_Interface_HoldTime) GetSegmentPath() string {
    return "hold-time"
}

func (holdTime *Interfaces_Interface_HoldTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &holdTime.Config
    }
    if childYangName == "state" {
        return &holdTime.State
    }
    return nil
}

func (holdTime *Interfaces_Interface_HoldTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &holdTime.Config
    children["state"] = &holdTime.State
    return children
}

func (holdTime *Interfaces_Interface_HoldTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (holdTime *Interfaces_Interface_HoldTime) GetBundleName() string { return "openconfig" }

func (holdTime *Interfaces_Interface_HoldTime) GetYangName() string { return "hold-time" }

func (holdTime *Interfaces_Interface_HoldTime) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (holdTime *Interfaces_Interface_HoldTime) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (holdTime *Interfaces_Interface_HoldTime) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (holdTime *Interfaces_Interface_HoldTime) SetParent(parent types.Entity) { holdTime.parent = parent }

func (holdTime *Interfaces_Interface_HoldTime) GetParent() types.Entity { return holdTime.parent }

func (holdTime *Interfaces_Interface_HoldTime) GetParentYangName() string { return "interface" }

// Interfaces_Interface_HoldTime_Config
// Configuration data for interface hold-time settings.
type Interfaces_Interface_HoldTime_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_HoldTime_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_HoldTime_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_HoldTime_Config) GetGoName(yname string) string {
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    return ""
}

func (config *Interfaces_Interface_HoldTime_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_HoldTime_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_HoldTime_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_HoldTime_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["up"] = config.Up
    leafs["down"] = config.Down
    return leafs
}

func (config *Interfaces_Interface_HoldTime_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_HoldTime_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_HoldTime_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_HoldTime_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_HoldTime_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_HoldTime_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_HoldTime_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_HoldTime_Config) GetParentYangName() string { return "hold-time" }

// Interfaces_Interface_HoldTime_State
// Operational state data for interface hold-time.
type Interfaces_Interface_HoldTime_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_HoldTime_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_HoldTime_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_HoldTime_State) GetGoName(yname string) string {
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    return ""
}

func (state *Interfaces_Interface_HoldTime_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_HoldTime_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_HoldTime_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_HoldTime_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["up"] = state.Up
    leafs["down"] = state.Down
    return leafs
}

func (state *Interfaces_Interface_HoldTime_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_HoldTime_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_HoldTime_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_HoldTime_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_HoldTime_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_HoldTime_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_HoldTime_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_HoldTime_State) GetParentYangName() string { return "hold-time" }

// Interfaces_Interface_Subinterfaces
// Enclosing container for the list of subinterfaces associated
// with a physical interface
type Interfaces_Interface_Subinterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of subinterfaces (logical interfaces) associated with a physical
    // interface. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface.
    Subinterface []Interfaces_Interface_Subinterfaces_Subinterface
}

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetFilter() yfilter.YFilter { return subinterfaces.YFilter }

func (subinterfaces *Interfaces_Interface_Subinterfaces) SetFilter(yf yfilter.YFilter) { subinterfaces.YFilter = yf }

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetGoName(yname string) string {
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetSegmentPath() string {
    return "subinterfaces"
}

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subinterface" {
        for _, c := range subinterfaces.Subinterface {
            if subinterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Subinterfaces_Subinterface{}
        subinterfaces.Subinterface = append(subinterfaces.Subinterface, child)
        return &subinterfaces.Subinterface[len(subinterfaces.Subinterface)-1]
    }
    return nil
}

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subinterfaces.Subinterface {
        children[subinterfaces.Subinterface[i].GetSegmentPath()] = &subinterfaces.Subinterface[i]
    }
    return children
}

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetBundleName() string { return "openconfig" }

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetYangName() string { return "subinterfaces" }

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (subinterfaces *Interfaces_Interface_Subinterfaces) SetParent(parent types.Entity) { subinterfaces.parent = parent }

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetParent() types.Entity { return subinterfaces.parent }

func (subinterfaces *Interfaces_Interface_Subinterfaces) GetParentYangName() string { return "interface" }

// Interfaces_Interface_Subinterfaces_Subinterface
// The list of subinterfaces (logical interfaces) associated
// with a physical interface
type Interfaces_Interface_Subinterfaces_Subinterface struct {
    parent types.Entity
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

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetFilter() yfilter.YFilter { return subinterface.YFilter }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) SetFilter(yf yfilter.YFilter) { subinterface.YFilter = yf }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "openconfig-vlan:vlan" { return "Vlan" }
    if yname == "openconfig-if-ip:ipv4" { return "Ipv4" }
    if yname == "openconfig-if-ip:ipv6" { return "Ipv6" }
    return ""
}

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetSegmentPath() string {
    return "subinterface" + "[index='" + fmt.Sprintf("%v", subinterface.Index) + "']"
}

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &subinterface.Config
    }
    if childYangName == "state" {
        return &subinterface.State
    }
    if childYangName == "openconfig-vlan:vlan" {
        return &subinterface.Vlan
    }
    if childYangName == "openconfig-if-ip:ipv4" {
        return &subinterface.Ipv4
    }
    if childYangName == "openconfig-if-ip:ipv6" {
        return &subinterface.Ipv6
    }
    return nil
}

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &subinterface.Config
    children["state"] = &subinterface.State
    children["openconfig-vlan:vlan"] = &subinterface.Vlan
    children["openconfig-if-ip:ipv4"] = &subinterface.Ipv4
    children["openconfig-if-ip:ipv6"] = &subinterface.Ipv6
    return children
}

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = subinterface.Index
    return leafs
}

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetBundleName() string { return "openconfig" }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetYangName() string { return "subinterface" }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) SetParent(parent types.Entity) { subinterface.parent = parent }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetParent() types.Entity { return subinterface.parent }

func (subinterface *Interfaces_Interface_Subinterfaces_Subinterface) GetParentYangName() string { return "subinterfaces" }

// Interfaces_Interface_Subinterfaces_Subinterface_Config
// Configurable items at the subinterface level
type Interfaces_Interface_Subinterfaces_Subinterface_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = config.Index
    leafs["name"] = config.Name
    leafs["description"] = config.Description
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Config) GetParentYangName() string { return "subinterface" }

// Interfaces_Interface_Subinterfaces_Subinterface_State
// Operational state data for logical interfaces
type Interfaces_Interface_Subinterfaces_Subinterface_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "enabled" { return "Enabled" }
    if yname == "ifindex" { return "Ifindex" }
    if yname == "admin-status" { return "AdminStatus" }
    if yname == "oper-status" { return "OperStatus" }
    if yname == "last-change" { return "LastChange" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &state.Counters
    }
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &state.Counters
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = state.Index
    leafs["name"] = state.Name
    leafs["description"] = state.Description
    leafs["enabled"] = state.Enabled
    leafs["ifindex"] = state.Ifindex
    leafs["admin-status"] = state.AdminStatus
    leafs["oper-status"] = state.OperStatus
    leafs["last-change"] = state.LastChange
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_State) GetParentYangName() string { return "subinterface" }

// Interfaces_Interface_Subinterfaces_Subinterface_State_Counters
// A collection of interface-related statistics objects.
type Interfaces_Interface_Subinterfaces_Subinterface_State_Counters struct {
    parent types.Entity
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

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetGoName(yname string) string {
    if yname == "in-octets" { return "InOctets" }
    if yname == "in-unicast-pkts" { return "InUnicastPkts" }
    if yname == "in-broadcast-pkts" { return "InBroadcastPkts" }
    if yname == "in-multicast-pkts" { return "InMulticastPkts" }
    if yname == "in-discards" { return "InDiscards" }
    if yname == "in-errors" { return "InErrors" }
    if yname == "in-unknown-protos" { return "InUnknownProtos" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "out-unicast-pkts" { return "OutUnicastPkts" }
    if yname == "out-broadcast-pkts" { return "OutBroadcastPkts" }
    if yname == "out-multicast-pkts" { return "OutMulticastPkts" }
    if yname == "out-discards" { return "OutDiscards" }
    if yname == "out-errors" { return "OutErrors" }
    if yname == "last-clear" { return "LastClear" }
    return ""
}

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-octets"] = counters.InOctets
    leafs["in-unicast-pkts"] = counters.InUnicastPkts
    leafs["in-broadcast-pkts"] = counters.InBroadcastPkts
    leafs["in-multicast-pkts"] = counters.InMulticastPkts
    leafs["in-discards"] = counters.InDiscards
    leafs["in-errors"] = counters.InErrors
    leafs["in-unknown-protos"] = counters.InUnknownProtos
    leafs["out-octets"] = counters.OutOctets
    leafs["out-unicast-pkts"] = counters.OutUnicastPkts
    leafs["out-broadcast-pkts"] = counters.OutBroadcastPkts
    leafs["out-multicast-pkts"] = counters.OutMulticastPkts
    leafs["out-discards"] = counters.OutDiscards
    leafs["out-errors"] = counters.OutErrors
    leafs["last-clear"] = counters.LastClear
    return leafs
}

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetBundleName() string { return "openconfig" }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetYangName() string { return "counters" }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Interfaces_Interface_Subinterfaces_Subinterface_State_Counters) GetParentYangName() string { return "state" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters for VLANs.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config

    // State variables for VLANs.
    State Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State
}

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetFilter() yfilter.YFilter { return vlan.YFilter }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) SetFilter(yf yfilter.YFilter) { vlan.YFilter = yf }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetSegmentPath() string {
    return "openconfig-vlan:vlan"
}

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &vlan.Config
    }
    if childYangName == "state" {
        return &vlan.State
    }
    return nil
}

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &vlan.Config
    children["state"] = &vlan.State
    return children
}

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetBundleName() string { return "openconfig" }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetYangName() string { return "vlan" }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) SetParent(parent types.Entity) { vlan.parent = parent }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetParent() types.Entity { return vlan.parent }

func (vlan *Interfaces_Interface_Subinterfaces_Subinterface_Vlan) GetParentYangName() string { return "subinterface" }

// Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config
// Configuration parameters for VLANs
type Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN id for the subinterface -- specified inline for the case of a local
    // VLAN.  The id is scoped to the subinterface, and could be repeated on
    // different subinterfaces. The type is one of the following types: int with
    // range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    VlanId interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetGoName(yname string) string {
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-id"] = config.VlanId
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_Config) GetParentYangName() string { return "vlan" }

// Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State
// State variables for VLANs
type Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VLAN id for the subinterface -- specified inline for the case of a local
    // VLAN.  The id is scoped to the subinterface, and could be repeated on
    // different subinterfaces. The type is one of the following types: int with
    // range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    VlanId interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetGoName(yname string) string {
    if yname == "vlan-id" { return "VlanId" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan-id"] = state.VlanId
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Vlan_State) GetParentYangName() string { return "vlan" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4
// Parameters for the IPv4 address family.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4 struct {
    parent types.Entity
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

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetGoName(yname string) string {
    if yname == "addresses" { return "Addresses" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "unnumbered" { return "Unnumbered" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetSegmentPath() string {
    return "openconfig-if-ip:ipv4"
}

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &ipv4.Addresses
    }
    if childYangName == "neighbors" {
        return &ipv4.Neighbors
    }
    if childYangName == "unnumbered" {
        return &ipv4.Unnumbered
    }
    if childYangName == "config" {
        return &ipv4.Config
    }
    if childYangName == "state" {
        return &ipv4.State
    }
    return nil
}

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &ipv4.Addresses
    children["neighbors"] = &ipv4.Neighbors
    children["unnumbered"] = &ipv4.Unnumbered
    children["config"] = &ipv4.Config
    children["state"] = &ipv4.State
    return children
}

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetBundleName() string { return "openconfig" }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4) GetParentYangName() string { return "subinterface" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses
// Enclosing container for address list
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of configured IPv4 addresses on the interface. The type is slice
    // of Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address.
    Address []Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range addresses.Address {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address{}
        addresses.Address = append(addresses.Address, child)
        return &addresses.Address[len(addresses.Address)-1]
    }
    return nil
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.Address {
        children[addresses.Address[i].GetSegmentPath()] = &addresses.Address[i]
    }
    return children
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetBundleName() string { return "openconfig" }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetYangName() string { return "addresses" }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address
// The list of configured IPv4 addresses on the interface.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address struct {
    parent types.Entity
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

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "vrrp" { return "Vrrp" }
    return ""
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetSegmentPath() string {
    return "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &address.Config
    }
    if childYangName == "state" {
        return &address.State
    }
    if childYangName == "vrrp" {
        return &address.Vrrp
    }
    return nil
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &address.Config
    children["state"] = &address.State
    children["vrrp"] = &address.Vrrp
    return children
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = address.Ip
    return leafs
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetBundleName() string { return "openconfig" }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetYangName() string { return "address" }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetParent() types.Entity { return address.parent }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address) GetParentYangName() string { return "addresses" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config
// Configuration data for each configured IPv4
// address on the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv4 address on the interface.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..32.
    PrefixLength interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = config.Ip
    leafs["prefix-length"] = config.PrefixLength
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Config) GetParentYangName() string { return "address" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State
// Operational state data for each IPv4 address
// configured on the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "origin" { return "Origin" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = state.Ip
    leafs["prefix-length"] = state.PrefixLength
    leafs["origin"] = state.Origin
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_State) GetParentYangName() string { return "address" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp
// Enclosing container for VRRP groups handled by this
// IP interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of VRRP groups, keyed by virtual router id. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup.
    VrrpGroup []Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetFilter() yfilter.YFilter { return vrrp.YFilter }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) SetFilter(yf yfilter.YFilter) { vrrp.YFilter = yf }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetGoName(yname string) string {
    if yname == "vrrp-group" { return "VrrpGroup" }
    return ""
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetSegmentPath() string {
    return "vrrp"
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrrp-group" {
        for _, c := range vrrp.VrrpGroup {
            if vrrp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup{}
        vrrp.VrrpGroup = append(vrrp.VrrpGroup, child)
        return &vrrp.VrrpGroup[len(vrrp.VrrpGroup)-1]
    }
    return nil
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrrp.VrrpGroup {
        children[vrrp.VrrpGroup[i].GetSegmentPath()] = &vrrp.VrrpGroup[i]
    }
    return children
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetBundleName() string { return "openconfig" }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetYangName() string { return "vrrp" }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) SetParent(parent types.Entity) { vrrp.parent = parent }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetParent() types.Entity { return vrrp.parent }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp) GetParentYangName() string { return "address" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup
// List of VRRP groups, keyed by virtual router id
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup struct {
    parent types.Entity
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

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetFilter() yfilter.YFilter { return vrrpGroup.YFilter }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) SetFilter(yf yfilter.YFilter) { vrrpGroup.YFilter = yf }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-tracking" { return "InterfaceTracking" }
    return ""
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetSegmentPath() string {
    return "vrrp-group" + "[virtual-router-id='" + fmt.Sprintf("%v", vrrpGroup.VirtualRouterId) + "']"
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &vrrpGroup.Config
    }
    if childYangName == "state" {
        return &vrrpGroup.State
    }
    if childYangName == "interface-tracking" {
        return &vrrpGroup.InterfaceTracking
    }
    return nil
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &vrrpGroup.Config
    children["state"] = &vrrpGroup.State
    children["interface-tracking"] = &vrrpGroup.InterfaceTracking
    return children
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = vrrpGroup.VirtualRouterId
    return leafs
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetBundleName() string { return "openconfig" }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetYangName() string { return "vrrp-group" }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) SetParent(parent types.Entity) { vrrpGroup.parent = parent }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetParent() types.Entity { return vrrpGroup.parent }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetParentYangName() string { return "vrrp" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config
// Configuration data for the VRRP group
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "virtual-address" { return "VirtualAddress" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "accept-mode" { return "AcceptMode" }
    if yname == "advertisement-interval" { return "AdvertisementInterval" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = config.VirtualRouterId
    leafs["virtual-address"] = config.VirtualAddress
    leafs["priority"] = config.Priority
    leafs["preempt"] = config.Preempt
    leafs["preempt-delay"] = config.PreemptDelay
    leafs["accept-mode"] = config.AcceptMode
    leafs["advertisement-interval"] = config.AdvertisementInterval
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State
// Operational state data for the VRRP group
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "virtual-address" { return "VirtualAddress" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "accept-mode" { return "AcceptMode" }
    if yname == "advertisement-interval" { return "AdvertisementInterval" }
    if yname == "current-priority" { return "CurrentPriority" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = state.VirtualRouterId
    leafs["virtual-address"] = state.VirtualAddress
    leafs["priority"] = state.Priority
    leafs["preempt"] = state.Preempt
    leafs["preempt-delay"] = state.PreemptDelay
    leafs["accept-mode"] = state.AcceptMode
    leafs["advertisement-interval"] = state.AdvertisementInterval
    leafs["current-priority"] = state.CurrentPriority
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
// Top-level container for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for VRRP interface tracking.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config

    // Operational state data for VRRP interface tracking.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetFilter() yfilter.YFilter { return interfaceTracking.YFilter }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) SetFilter(yf yfilter.YFilter) { interfaceTracking.YFilter = yf }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetSegmentPath() string {
    return "interface-tracking"
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceTracking.Config
    }
    if childYangName == "state" {
        return &interfaceTracking.State
    }
    return nil
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceTracking.Config
    children["state"] = &interfaceTracking.State
    return children
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetBundleName() string { return "openconfig" }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetYangName() string { return "interface-tracking" }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) SetParent(parent types.Entity) { interfaceTracking.parent = parent }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetParent() types.Entity { return interfaceTracking.parent }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config
// Configuration data for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetGoName(yname string) string {
    if yname == "track-interface" { return "TrackInterface" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-interface"] = config.TrackInterface
    leafs["priority-decrement"] = config.PriorityDecrement
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetParentYangName() string { return "interface-tracking" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
// Operational state data for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetGoName(yname string) string {
    if yname == "track-interface" { return "TrackInterface" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-interface"] = state.TrackInterface
    leafs["priority-decrement"] = state.PriorityDecrement
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetParentYangName() string { return "interface-tracking" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors
// Enclosing container for neighbor list
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of mappings from IPv4 addresses to link-layer addresses.  Entries in
    // this list are used as static entries in the ARP Cache. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor.
    Neighbor []Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetBundleName() string { return "openconfig" }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor
// A list of mappings from IPv4 addresses to
// link-layer addresses.
// 
// Entries in this list are used as static entries in the
// ARP Cache.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor struct {
    parent types.Entity
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

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &neighbor.Config
    }
    if childYangName == "state" {
        return &neighbor.State
    }
    return nil
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &neighbor.Config
    children["state"] = &neighbor.State
    return children
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = neighbor.Ip
    return leafs
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetBundleName() string { return "openconfig" }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config
// Configuration data for each configured IPv4
// address on the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPv4 address of the neighbor node. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = config.Ip
    leafs["link-layer-address"] = config.LinkLayerAddress
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_Config) GetParentYangName() string { return "neighbor" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State
// Operational state data for each IPv4 address
// configured on the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    if yname == "origin" { return "Origin" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = state.Ip
    leafs["link-layer-address"] = state.LinkLayerAddress
    leafs["origin"] = state.Origin
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Neighbors_Neighbor_State) GetParentYangName() string { return "neighbor" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered
// Top-level container for setting unnumbered interfaces.
// Includes reference the interface that provides the
// address information
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for unnumbered interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config

    // Operational state data for unnumbered interfaces.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State

    // Reference to an interface or subinterface.
    InterfaceRef Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetFilter() yfilter.YFilter { return unnumbered.YFilter }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) SetFilter(yf yfilter.YFilter) { unnumbered.YFilter = yf }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-ref" { return "InterfaceRef" }
    return ""
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetSegmentPath() string {
    return "unnumbered"
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &unnumbered.Config
    }
    if childYangName == "state" {
        return &unnumbered.State
    }
    if childYangName == "interface-ref" {
        return &unnumbered.InterfaceRef
    }
    return nil
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &unnumbered.Config
    children["state"] = &unnumbered.State
    children["interface-ref"] = &unnumbered.InterfaceRef
    return children
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetBundleName() string { return "openconfig" }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetYangName() string { return "unnumbered" }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) SetParent(parent types.Entity) { unnumbered.parent = parent }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetParent() types.Entity { return unnumbered.parent }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config
// Configuration data for unnumbered interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_Config) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State
// Operational state data for unnumbered interfaces
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_State) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef
// Reference to an interface or subinterface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config

    // Operational state for interface-ref.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetFilter() yfilter.YFilter { return interfaceRef.YFilter }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) SetFilter(yf yfilter.YFilter) { interfaceRef.YFilter = yf }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetSegmentPath() string {
    return "interface-ref"
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceRef.Config
    }
    if childYangName == "state" {
        return &interfaceRef.State
    }
    return nil
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceRef.Config
    children["state"] = &interfaceRef.State
    return children
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetBundleName() string { return "openconfig" }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetYangName() string { return "interface-ref" }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) SetParent(parent types.Entity) { interfaceRef.parent = parent }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetParent() types.Entity { return interfaceRef.parent }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config
// Configured reference to interface / subinterface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = config.Interface
    leafs["subinterface"] = config.Subinterface
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_Config) GetParentYangName() string { return "interface-ref" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State
// Operational state for interface-ref
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    leafs["subinterface"] = state.Subinterface
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Unnumbered_InterfaceRef_State) GetParentYangName() string { return "interface-ref" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config
// Top-level IPv4 configuration data for the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "mtu" { return "Mtu" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["mtu"] = config.Mtu
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_Config) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State
// Top level IPv4 operational state data
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "mtu" { return "Mtu" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["mtu"] = state.Mtu
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv4_State) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6
// Parameters for the IPv6 address family.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6 struct {
    parent types.Entity
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

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetGoName(yname string) string {
    if yname == "addresses" { return "Addresses" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "unnumbered" { return "Unnumbered" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetSegmentPath() string {
    return "openconfig-if-ip:ipv6"
}

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &ipv6.Addresses
    }
    if childYangName == "neighbors" {
        return &ipv6.Neighbors
    }
    if childYangName == "unnumbered" {
        return &ipv6.Unnumbered
    }
    if childYangName == "config" {
        return &ipv6.Config
    }
    if childYangName == "state" {
        return &ipv6.State
    }
    return nil
}

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &ipv6.Addresses
    children["neighbors"] = &ipv6.Neighbors
    children["unnumbered"] = &ipv6.Unnumbered
    children["config"] = &ipv6.Config
    children["state"] = &ipv6.State
    return children
}

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetBundleName() string { return "openconfig" }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6) GetParentYangName() string { return "subinterface" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses
// Enclosing container for address list
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of configured IPv6 addresses on the interface. The type is slice
    // of Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address.
    Address []Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range addresses.Address {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address{}
        addresses.Address = append(addresses.Address, child)
        return &addresses.Address[len(addresses.Address)-1]
    }
    return nil
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.Address {
        children[addresses.Address[i].GetSegmentPath()] = &addresses.Address[i]
    }
    return children
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetBundleName() string { return "openconfig" }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetYangName() string { return "addresses" }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address
// The list of configured IPv6 addresses on the interface.
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address struct {
    parent types.Entity
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

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "vrrp" { return "Vrrp" }
    return ""
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetSegmentPath() string {
    return "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &address.Config
    }
    if childYangName == "state" {
        return &address.State
    }
    if childYangName == "vrrp" {
        return &address.Vrrp
    }
    return nil
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &address.Config
    children["state"] = &address.State
    children["vrrp"] = &address.Vrrp
    return children
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = address.Ip
    return leafs
}

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetBundleName() string { return "openconfig" }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetYangName() string { return "address" }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetParent() types.Entity { return address.parent }

func (address *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address) GetParentYangName() string { return "addresses" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config
// Configuration data for each IPv6 address on
// the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address on the interface.
    // The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..128. This attribute is mandatory.
    PrefixLength interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = config.Ip
    leafs["prefix-length"] = config.PrefixLength
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Config) GetParentYangName() string { return "address" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State
// State data for each IPv6 address on the
// interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "origin" { return "Origin" }
    if yname == "status" { return "Status" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = state.Ip
    leafs["prefix-length"] = state.PrefixLength
    leafs["origin"] = state.Origin
    leafs["status"] = state.Status
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_State) GetParentYangName() string { return "address" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of VRRP groups, keyed by virtual router id. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup.
    VrrpGroup []Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetFilter() yfilter.YFilter { return vrrp.YFilter }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) SetFilter(yf yfilter.YFilter) { vrrp.YFilter = yf }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetGoName(yname string) string {
    if yname == "vrrp-group" { return "VrrpGroup" }
    return ""
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetSegmentPath() string {
    return "vrrp"
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrrp-group" {
        for _, c := range vrrp.VrrpGroup {
            if vrrp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup{}
        vrrp.VrrpGroup = append(vrrp.VrrpGroup, child)
        return &vrrp.VrrpGroup[len(vrrp.VrrpGroup)-1]
    }
    return nil
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrrp.VrrpGroup {
        children[vrrp.VrrpGroup[i].GetSegmentPath()] = &vrrp.VrrpGroup[i]
    }
    return children
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetBundleName() string { return "openconfig" }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetYangName() string { return "vrrp" }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) SetParent(parent types.Entity) { vrrp.parent = parent }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetParent() types.Entity { return vrrp.parent }

func (vrrp *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp) GetParentYangName() string { return "address" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup
// List of VRRP groups, keyed by virtual router id
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup struct {
    parent types.Entity
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

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetFilter() yfilter.YFilter { return vrrpGroup.YFilter }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) SetFilter(yf yfilter.YFilter) { vrrpGroup.YFilter = yf }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-tracking" { return "InterfaceTracking" }
    return ""
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetSegmentPath() string {
    return "vrrp-group" + "[virtual-router-id='" + fmt.Sprintf("%v", vrrpGroup.VirtualRouterId) + "']"
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &vrrpGroup.Config
    }
    if childYangName == "state" {
        return &vrrpGroup.State
    }
    if childYangName == "interface-tracking" {
        return &vrrpGroup.InterfaceTracking
    }
    return nil
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &vrrpGroup.Config
    children["state"] = &vrrpGroup.State
    children["interface-tracking"] = &vrrpGroup.InterfaceTracking
    return children
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = vrrpGroup.VirtualRouterId
    return leafs
}

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetBundleName() string { return "openconfig" }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetYangName() string { return "vrrp-group" }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) SetParent(parent types.Entity) { vrrpGroup.parent = parent }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetParent() types.Entity { return vrrpGroup.parent }

func (vrrpGroup *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetParentYangName() string { return "vrrp" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config
// Configuration data for the VRRP group
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "virtual-address" { return "VirtualAddress" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "accept-mode" { return "AcceptMode" }
    if yname == "advertisement-interval" { return "AdvertisementInterval" }
    if yname == "virtual-link-local" { return "VirtualLinkLocal" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = config.VirtualRouterId
    leafs["virtual-address"] = config.VirtualAddress
    leafs["priority"] = config.Priority
    leafs["preempt"] = config.Preempt
    leafs["preempt-delay"] = config.PreemptDelay
    leafs["accept-mode"] = config.AcceptMode
    leafs["advertisement-interval"] = config.AdvertisementInterval
    leafs["virtual-link-local"] = config.VirtualLinkLocal
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State
// Operational state data for the VRRP group
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "virtual-address" { return "VirtualAddress" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "accept-mode" { return "AcceptMode" }
    if yname == "advertisement-interval" { return "AdvertisementInterval" }
    if yname == "current-priority" { return "CurrentPriority" }
    if yname == "virtual-link-local" { return "VirtualLinkLocal" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = state.VirtualRouterId
    leafs["virtual-address"] = state.VirtualAddress
    leafs["priority"] = state.Priority
    leafs["preempt"] = state.Preempt
    leafs["preempt-delay"] = state.PreemptDelay
    leafs["accept-mode"] = state.AcceptMode
    leafs["advertisement-interval"] = state.AdvertisementInterval
    leafs["current-priority"] = state.CurrentPriority
    leafs["virtual-link-local"] = state.VirtualLinkLocal
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
// Top-level container for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for VRRP interface tracking.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config

    // Operational state data for VRRP interface tracking.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetFilter() yfilter.YFilter { return interfaceTracking.YFilter }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) SetFilter(yf yfilter.YFilter) { interfaceTracking.YFilter = yf }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetSegmentPath() string {
    return "interface-tracking"
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceTracking.Config
    }
    if childYangName == "state" {
        return &interfaceTracking.State
    }
    return nil
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceTracking.Config
    children["state"] = &interfaceTracking.State
    return children
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetBundleName() string { return "openconfig" }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetYangName() string { return "interface-tracking" }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) SetParent(parent types.Entity) { interfaceTracking.parent = parent }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetParent() types.Entity { return interfaceTracking.parent }

func (interfaceTracking *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config
// Configuration data for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetGoName(yname string) string {
    if yname == "track-interface" { return "TrackInterface" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-interface"] = config.TrackInterface
    leafs["priority-decrement"] = config.PriorityDecrement
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetParentYangName() string { return "interface-tracking" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
// Operational state data for VRRP interface tracking
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetGoName(yname string) string {
    if yname == "track-interface" { return "TrackInterface" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-interface"] = state.TrackInterface
    leafs["priority-decrement"] = state.PriorityDecrement
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetParentYangName() string { return "interface-tracking" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors
// Enclosing container for list of IPv6 neighbors
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of IPv6 neighbors. The type is slice of
    // Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor.
    Neighbor []Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetBundleName() string { return "openconfig" }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor
// List of IPv6 neighbors
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor struct {
    parent types.Entity
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

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &neighbor.Config
    }
    if childYangName == "state" {
        return &neighbor.State
    }
    return nil
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &neighbor.Config
    children["state"] = &neighbor.State
    return children
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = neighbor.Ip
    return leafs
}

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetBundleName() string { return "openconfig" }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config
// Configuration data for each IPv6 address on
// the interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = config.Ip
    leafs["link-layer-address"] = config.LinkLayerAddress
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_Config) GetParentYangName() string { return "neighbor" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State
// State data for each IPv6 address on the
// interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    if yname == "origin" { return "Origin" }
    if yname == "is-router" { return "IsRouter" }
    if yname == "neighbor-state" { return "NeighborState" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = state.Ip
    leafs["link-layer-address"] = state.LinkLayerAddress
    leafs["origin"] = state.Origin
    leafs["is-router"] = state.IsRouter
    leafs["neighbor-state"] = state.NeighborState
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Neighbors_Neighbor_State) GetParentYangName() string { return "neighbor" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for unnumbered interface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config

    // Operational state data for unnumbered interfaces.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State

    // Reference to an interface or subinterface.
    InterfaceRef Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetFilter() yfilter.YFilter { return unnumbered.YFilter }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) SetFilter(yf yfilter.YFilter) { unnumbered.YFilter = yf }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-ref" { return "InterfaceRef" }
    return ""
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetSegmentPath() string {
    return "unnumbered"
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &unnumbered.Config
    }
    if childYangName == "state" {
        return &unnumbered.State
    }
    if childYangName == "interface-ref" {
        return &unnumbered.InterfaceRef
    }
    return nil
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &unnumbered.Config
    children["state"] = &unnumbered.State
    children["interface-ref"] = &unnumbered.InterfaceRef
    return children
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetBundleName() string { return "openconfig" }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetYangName() string { return "unnumbered" }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) SetParent(parent types.Entity) { unnumbered.parent = parent }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetParent() types.Entity { return unnumbered.parent }

func (unnumbered *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config
// Configuration data for unnumbered interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_Config) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State
// Operational state data for unnumbered interfaces
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_State) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef
// Reference to an interface or subinterface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config

    // Operational state for interface-ref.
    State Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetFilter() yfilter.YFilter { return interfaceRef.YFilter }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) SetFilter(yf yfilter.YFilter) { interfaceRef.YFilter = yf }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetSegmentPath() string {
    return "interface-ref"
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceRef.Config
    }
    if childYangName == "state" {
        return &interfaceRef.State
    }
    return nil
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceRef.Config
    children["state"] = &interfaceRef.State
    return children
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetBundleName() string { return "openconfig" }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetYangName() string { return "interface-ref" }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) SetParent(parent types.Entity) { interfaceRef.parent = parent }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetParent() types.Entity { return interfaceRef.parent }

func (interfaceRef *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config
// Configured reference to interface / subinterface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = config.Interface
    leafs["subinterface"] = config.Subinterface
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_Config) GetParentYangName() string { return "interface-ref" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State
// Operational state for interface-ref
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    leafs["subinterface"] = state.Subinterface
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Unnumbered_InterfaceRef_State) GetParentYangName() string { return "interface-ref" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config
// Top-level config data for the IPv6 interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "mtu" { return "Mtu" }
    if yname == "dup-addr-detect-transmits" { return "DupAddrDetectTransmits" }
    return ""
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["mtu"] = config.Mtu
    leafs["dup-addr-detect-transmits"] = config.DupAddrDetectTransmits
    return leafs
}

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_Config) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State
// Top-level operational state data for the IPv6 interface
type Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "mtu" { return "Mtu" }
    if yname == "dup-addr-detect-transmits" { return "DupAddrDetectTransmits" }
    return ""
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["mtu"] = state.Mtu
    leafs["dup-addr-detect-transmits"] = state.DupAddrDetectTransmits
    return leafs
}

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Subinterfaces_Subinterface_Ipv6_State) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Ethernet
// Top-level container for ethernet configuration
// and state
type Interfaces_Interface_Ethernet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for ethernet interfaces.
    Config Interfaces_Interface_Ethernet_Config

    // State variables for Ethernet interfaces.
    State Interfaces_Interface_Ethernet_State

    // Enclosing container for VLAN interface-specific data on Ethernet
    // interfaces.  These are for standard L2, switched-style VLANs.
    SwitchedVlan Interfaces_Interface_Ethernet_SwitchedVlan
}

func (ethernet *Interfaces_Interface_Ethernet) GetFilter() yfilter.YFilter { return ethernet.YFilter }

func (ethernet *Interfaces_Interface_Ethernet) SetFilter(yf yfilter.YFilter) { ethernet.YFilter = yf }

func (ethernet *Interfaces_Interface_Ethernet) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "openconfig-vlan:switched-vlan" { return "SwitchedVlan" }
    return ""
}

func (ethernet *Interfaces_Interface_Ethernet) GetSegmentPath() string {
    return "openconfig-if-ethernet:ethernet"
}

func (ethernet *Interfaces_Interface_Ethernet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &ethernet.Config
    }
    if childYangName == "state" {
        return &ethernet.State
    }
    if childYangName == "openconfig-vlan:switched-vlan" {
        return &ethernet.SwitchedVlan
    }
    return nil
}

func (ethernet *Interfaces_Interface_Ethernet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &ethernet.Config
    children["state"] = &ethernet.State
    children["openconfig-vlan:switched-vlan"] = &ethernet.SwitchedVlan
    return children
}

func (ethernet *Interfaces_Interface_Ethernet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernet *Interfaces_Interface_Ethernet) GetBundleName() string { return "openconfig" }

func (ethernet *Interfaces_Interface_Ethernet) GetYangName() string { return "ethernet" }

func (ethernet *Interfaces_Interface_Ethernet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ethernet *Interfaces_Interface_Ethernet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ethernet *Interfaces_Interface_Ethernet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ethernet *Interfaces_Interface_Ethernet) SetParent(parent types.Entity) { ethernet.parent = parent }

func (ethernet *Interfaces_Interface_Ethernet) GetParent() types.Entity { return ethernet.parent }

func (ethernet *Interfaces_Interface_Ethernet) GetParentYangName() string { return "interface" }

// Interfaces_Interface_Ethernet_Config
// Configuration data for ethernet interfaces
type Interfaces_Interface_Ethernet_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_Ethernet_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Ethernet_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Ethernet_Config) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "auto-negotiate" { return "AutoNegotiate" }
    if yname == "duplex-mode" { return "DuplexMode" }
    if yname == "port-speed" { return "PortSpeed" }
    if yname == "enable-flow-control" { return "EnableFlowControl" }
    if yname == "aggregate-id" { return "AggregateId" }
    return ""
}

func (config *Interfaces_Interface_Ethernet_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Ethernet_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Ethernet_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Ethernet_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = config.MacAddress
    leafs["auto-negotiate"] = config.AutoNegotiate
    leafs["duplex-mode"] = config.DuplexMode
    leafs["port-speed"] = config.PortSpeed
    leafs["enable-flow-control"] = config.EnableFlowControl
    leafs["aggregate-id"] = config.AggregateId
    return leafs
}

func (config *Interfaces_Interface_Ethernet_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Ethernet_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Ethernet_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Ethernet_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Ethernet_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Ethernet_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Ethernet_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Ethernet_Config) GetParentYangName() string { return "ethernet" }

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
    parent types.Entity
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

func (state *Interfaces_Interface_Ethernet_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Ethernet_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Ethernet_State) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "auto-negotiate" { return "AutoNegotiate" }
    if yname == "duplex-mode" { return "DuplexMode" }
    if yname == "port-speed" { return "PortSpeed" }
    if yname == "enable-flow-control" { return "EnableFlowControl" }
    if yname == "hw-mac-address" { return "HwMacAddress" }
    if yname == "effective-speed" { return "EffectiveSpeed" }
    if yname == "aggregate-id" { return "AggregateId" }
    if yname == "counters" { return "Counters" }
    return ""
}

func (state *Interfaces_Interface_Ethernet_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Ethernet_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "counters" {
        return &state.Counters
    }
    return nil
}

func (state *Interfaces_Interface_Ethernet_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["counters"] = &state.Counters
    return children
}

func (state *Interfaces_Interface_Ethernet_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = state.MacAddress
    leafs["auto-negotiate"] = state.AutoNegotiate
    leafs["duplex-mode"] = state.DuplexMode
    leafs["port-speed"] = state.PortSpeed
    leafs["enable-flow-control"] = state.EnableFlowControl
    leafs["hw-mac-address"] = state.HwMacAddress
    leafs["effective-speed"] = state.EffectiveSpeed
    leafs["aggregate-id"] = state.AggregateId
    return leafs
}

func (state *Interfaces_Interface_Ethernet_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Ethernet_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Ethernet_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Ethernet_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Ethernet_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Ethernet_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Ethernet_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Ethernet_State) GetParentYangName() string { return "ethernet" }

// Interfaces_Interface_Ethernet_State_Counters
// Ethernet interface counters
type Interfaces_Interface_Ethernet_State_Counters struct {
    parent types.Entity
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
    In8021QFrames interface{}

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
    Out8021QFrames interface{}
}

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetFilter() yfilter.YFilter { return counters.YFilter }

func (counters *Interfaces_Interface_Ethernet_State_Counters) SetFilter(yf yfilter.YFilter) { counters.YFilter = yf }

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetGoName(yname string) string {
    if yname == "in-mac-control-frames" { return "InMacControlFrames" }
    if yname == "in-mac-pause-frames" { return "InMacPauseFrames" }
    if yname == "in-oversize-frames" { return "InOversizeFrames" }
    if yname == "in-jabber-frames" { return "InJabberFrames" }
    if yname == "in-fragment-frames" { return "InFragmentFrames" }
    if yname == "in-8021q-frames" { return "In8021QFrames" }
    if yname == "in-crc-errors" { return "InCrcErrors" }
    if yname == "out-mac-control-frames" { return "OutMacControlFrames" }
    if yname == "out-mac-pause-frames" { return "OutMacPauseFrames" }
    if yname == "out-8021q-frames" { return "Out8021QFrames" }
    return ""
}

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetSegmentPath() string {
    return "counters"
}

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-mac-control-frames"] = counters.InMacControlFrames
    leafs["in-mac-pause-frames"] = counters.InMacPauseFrames
    leafs["in-oversize-frames"] = counters.InOversizeFrames
    leafs["in-jabber-frames"] = counters.InJabberFrames
    leafs["in-fragment-frames"] = counters.InFragmentFrames
    leafs["in-8021q-frames"] = counters.In8021QFrames
    leafs["in-crc-errors"] = counters.InCrcErrors
    leafs["out-mac-control-frames"] = counters.OutMacControlFrames
    leafs["out-mac-pause-frames"] = counters.OutMacPauseFrames
    leafs["out-8021q-frames"] = counters.Out8021QFrames
    return leafs
}

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetBundleName() string { return "openconfig" }

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetYangName() string { return "counters" }

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (counters *Interfaces_Interface_Ethernet_State_Counters) SetParent(parent types.Entity) { counters.parent = parent }

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetParent() types.Entity { return counters.parent }

func (counters *Interfaces_Interface_Ethernet_State_Counters) GetParentYangName() string { return "state" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters for VLANs.
    Config Interfaces_Interface_Ethernet_SwitchedVlan_Config

    // State variables for VLANs.
    State Interfaces_Interface_Ethernet_SwitchedVlan_State
}

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetFilter() yfilter.YFilter { return switchedVlan.YFilter }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) SetFilter(yf yfilter.YFilter) { switchedVlan.YFilter = yf }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetSegmentPath() string {
    return "openconfig-vlan:switched-vlan"
}

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &switchedVlan.Config
    }
    if childYangName == "state" {
        return &switchedVlan.State
    }
    return nil
}

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &switchedVlan.Config
    children["state"] = &switchedVlan.State
    return children
}

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetBundleName() string { return "openconfig" }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetYangName() string { return "switched-vlan" }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) SetParent(parent types.Entity) { switchedVlan.parent = parent }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetParent() types.Entity { return switchedVlan.parent }

func (switchedVlan *Interfaces_Interface_Ethernet_SwitchedVlan) GetParentYangName() string { return "ethernet" }

// Interfaces_Interface_Ethernet_SwitchedVlan_Config
// Configuration parameters for VLANs
type Interfaces_Interface_Ethernet_SwitchedVlan_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the interface to access or trunk mode for VLANs. The type is
    // VlanModeType.
    InterfaceMode interface{}

    // Set the native VLAN id for untagged frames arriving on a trunk interface. 
    // This configuration is only valid on a trunk interface. The type is one of
    // the following types: int with range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    NativeVlan interface{}

    // Assign the access vlan to the access port. The type is one of the following
    // types: int with range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    AccessVlan interface{}

    // Specify VLANs, or ranges thereof, that the interface may carry when in
    // trunk mode.  If not specified, all VLANs are allowed on the interface.
    // Ranges are specified in the form x..y, where x<y - ranges are assumed to be
    // inclusive (such that the VLAN range is x <= range <= y. The type is one of
    // the following types: slice of int with range: 1..4094, or slice of string
    // with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]),
    // or slice of string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*),
    // or slice of string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*),
    // or slice of string with pattern:
    // (\*|(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]))\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])..
    TrunkVlans []interface{}
}

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetGoName(yname string) string {
    if yname == "interface-mode" { return "InterfaceMode" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "access-vlan" { return "AccessVlan" }
    if yname == "trunk-vlans" { return "TrunkVlans" }
    return ""
}

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-mode"] = config.InterfaceMode
    leafs["native-vlan"] = config.NativeVlan
    leafs["access-vlan"] = config.AccessVlan
    leafs["trunk-vlans"] = config.TrunkVlans
    return leafs
}

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Ethernet_SwitchedVlan_Config) GetParentYangName() string { return "switched-vlan" }

// Interfaces_Interface_Ethernet_SwitchedVlan_State
// State variables for VLANs
type Interfaces_Interface_Ethernet_SwitchedVlan_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the interface to access or trunk mode for VLANs. The type is
    // VlanModeType.
    InterfaceMode interface{}

    // Set the native VLAN id for untagged frames arriving on a trunk interface. 
    // This configuration is only valid on a trunk interface. The type is one of
    // the following types: int with range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    NativeVlan interface{}

    // Assign the access vlan to the access port. The type is one of the following
    // types: int with range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    AccessVlan interface{}

    // Specify VLANs, or ranges thereof, that the interface may carry when in
    // trunk mode.  If not specified, all VLANs are allowed on the interface.
    // Ranges are specified in the form x..y, where x<y - ranges are assumed to be
    // inclusive (such that the VLAN range is x <= range <= y. The type is one of
    // the following types: slice of int with range: 1..4094, or slice of string
    // with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]),
    // or slice of string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*),
    // or slice of string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*),
    // or slice of string with pattern:
    // (\*|(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]))\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])..
    TrunkVlans []interface{}
}

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetGoName(yname string) string {
    if yname == "interface-mode" { return "InterfaceMode" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "access-vlan" { return "AccessVlan" }
    if yname == "trunk-vlans" { return "TrunkVlans" }
    return ""
}

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-mode"] = state.InterfaceMode
    leafs["native-vlan"] = state.NativeVlan
    leafs["access-vlan"] = state.AccessVlan
    leafs["trunk-vlans"] = state.TrunkVlans
    return leafs
}

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Ethernet_SwitchedVlan_State) GetParentYangName() string { return "switched-vlan" }

// Interfaces_Interface_Aggregation
// Options for logical interfaces representing
// aggregates
type Interfaces_Interface_Aggregation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration variables for logical aggregate / LAG interfaces.
    Config Interfaces_Interface_Aggregation_Config

    // Operational state variables for logical aggregate / LAG interfaces.
    State Interfaces_Interface_Aggregation_State

    // Enclosing container for VLAN interface-specific data on Ethernet
    // interfaces.  These are for standard L2, switched-style VLANs.
    SwitchedVlan Interfaces_Interface_Aggregation_SwitchedVlan
}

func (aggregation *Interfaces_Interface_Aggregation) GetFilter() yfilter.YFilter { return aggregation.YFilter }

func (aggregation *Interfaces_Interface_Aggregation) SetFilter(yf yfilter.YFilter) { aggregation.YFilter = yf }

func (aggregation *Interfaces_Interface_Aggregation) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "openconfig-vlan:switched-vlan" { return "SwitchedVlan" }
    return ""
}

func (aggregation *Interfaces_Interface_Aggregation) GetSegmentPath() string {
    return "openconfig-if-aggregate:aggregation"
}

func (aggregation *Interfaces_Interface_Aggregation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &aggregation.Config
    }
    if childYangName == "state" {
        return &aggregation.State
    }
    if childYangName == "openconfig-vlan:switched-vlan" {
        return &aggregation.SwitchedVlan
    }
    return nil
}

func (aggregation *Interfaces_Interface_Aggregation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &aggregation.Config
    children["state"] = &aggregation.State
    children["openconfig-vlan:switched-vlan"] = &aggregation.SwitchedVlan
    return children
}

func (aggregation *Interfaces_Interface_Aggregation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aggregation *Interfaces_Interface_Aggregation) GetBundleName() string { return "openconfig" }

func (aggregation *Interfaces_Interface_Aggregation) GetYangName() string { return "aggregation" }

func (aggregation *Interfaces_Interface_Aggregation) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (aggregation *Interfaces_Interface_Aggregation) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (aggregation *Interfaces_Interface_Aggregation) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (aggregation *Interfaces_Interface_Aggregation) SetParent(parent types.Entity) { aggregation.parent = parent }

func (aggregation *Interfaces_Interface_Aggregation) GetParent() types.Entity { return aggregation.parent }

func (aggregation *Interfaces_Interface_Aggregation) GetParentYangName() string { return "interface" }

// Interfaces_Interface_Aggregation_Config
// Configuration variables for logical aggregate /
// LAG interfaces
type Interfaces_Interface_Aggregation_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sets the type of LAG, i.e., how it is configured / maintained. The type is
    // AggregationType.
    LagType interface{}

    // Specifies the mininum number of member interfaces that must be active for
    // the aggregate interface to be available. The type is interface{} with
    // range: 0..65535.
    MinLinks interface{}
}

func (config *Interfaces_Interface_Aggregation_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Aggregation_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Aggregation_Config) GetGoName(yname string) string {
    if yname == "lag-type" { return "LagType" }
    if yname == "min-links" { return "MinLinks" }
    return ""
}

func (config *Interfaces_Interface_Aggregation_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Aggregation_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Aggregation_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Aggregation_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lag-type"] = config.LagType
    leafs["min-links"] = config.MinLinks
    return leafs
}

func (config *Interfaces_Interface_Aggregation_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Aggregation_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Aggregation_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Aggregation_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Aggregation_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Aggregation_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Aggregation_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Aggregation_Config) GetParentYangName() string { return "aggregation" }

// Interfaces_Interface_Aggregation_State
// Operational state variables for logical
// aggregate / LAG interfaces
type Interfaces_Interface_Aggregation_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_Aggregation_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Aggregation_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Aggregation_State) GetGoName(yname string) string {
    if yname == "lag-type" { return "LagType" }
    if yname == "min-links" { return "MinLinks" }
    if yname == "lag-speed" { return "LagSpeed" }
    if yname == "member" { return "Member" }
    return ""
}

func (state *Interfaces_Interface_Aggregation_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Aggregation_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Aggregation_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Aggregation_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lag-type"] = state.LagType
    leafs["min-links"] = state.MinLinks
    leafs["lag-speed"] = state.LagSpeed
    leafs["member"] = state.Member
    return leafs
}

func (state *Interfaces_Interface_Aggregation_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Aggregation_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Aggregation_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Aggregation_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Aggregation_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Aggregation_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Aggregation_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Aggregation_State) GetParentYangName() string { return "aggregation" }

// Interfaces_Interface_Aggregation_SwitchedVlan
// Enclosing container for VLAN interface-specific
// data on Ethernet interfaces.  These are for standard
// L2, switched-style VLANs.
type Interfaces_Interface_Aggregation_SwitchedVlan struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration parameters for VLANs.
    Config Interfaces_Interface_Aggregation_SwitchedVlan_Config

    // State variables for VLANs.
    State Interfaces_Interface_Aggregation_SwitchedVlan_State
}

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetFilter() yfilter.YFilter { return switchedVlan.YFilter }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) SetFilter(yf yfilter.YFilter) { switchedVlan.YFilter = yf }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetSegmentPath() string {
    return "openconfig-vlan:switched-vlan"
}

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &switchedVlan.Config
    }
    if childYangName == "state" {
        return &switchedVlan.State
    }
    return nil
}

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &switchedVlan.Config
    children["state"] = &switchedVlan.State
    return children
}

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetBundleName() string { return "openconfig" }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetYangName() string { return "switched-vlan" }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) SetParent(parent types.Entity) { switchedVlan.parent = parent }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetParent() types.Entity { return switchedVlan.parent }

func (switchedVlan *Interfaces_Interface_Aggregation_SwitchedVlan) GetParentYangName() string { return "aggregation" }

// Interfaces_Interface_Aggregation_SwitchedVlan_Config
// Configuration parameters for VLANs
type Interfaces_Interface_Aggregation_SwitchedVlan_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the interface to access or trunk mode for VLANs. The type is
    // VlanModeType.
    InterfaceMode interface{}

    // Set the native VLAN id for untagged frames arriving on a trunk interface. 
    // This configuration is only valid on a trunk interface. The type is one of
    // the following types: int with range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    NativeVlan interface{}

    // Assign the access vlan to the access port. The type is one of the following
    // types: int with range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    AccessVlan interface{}

    // Specify VLANs, or ranges thereof, that the interface may carry when in
    // trunk mode.  If not specified, all VLANs are allowed on the interface.
    // Ranges are specified in the form x..y, where x<y - ranges are assumed to be
    // inclusive (such that the VLAN range is x <= range <= y. The type is one of
    // the following types: slice of int with range: 1..4094, or slice of string
    // with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]),
    // or slice of string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*),
    // or slice of string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*),
    // or slice of string with pattern:
    // (\*|(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]))\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])..
    TrunkVlans []interface{}
}

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetGoName(yname string) string {
    if yname == "interface-mode" { return "InterfaceMode" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "access-vlan" { return "AccessVlan" }
    if yname == "trunk-vlans" { return "TrunkVlans" }
    return ""
}

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-mode"] = config.InterfaceMode
    leafs["native-vlan"] = config.NativeVlan
    leafs["access-vlan"] = config.AccessVlan
    leafs["trunk-vlans"] = config.TrunkVlans
    return leafs
}

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_Aggregation_SwitchedVlan_Config) GetParentYangName() string { return "switched-vlan" }

// Interfaces_Interface_Aggregation_SwitchedVlan_State
// State variables for VLANs
type Interfaces_Interface_Aggregation_SwitchedVlan_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the interface to access or trunk mode for VLANs. The type is
    // VlanModeType.
    InterfaceMode interface{}

    // Set the native VLAN id for untagged frames arriving on a trunk interface. 
    // This configuration is only valid on a trunk interface. The type is one of
    // the following types: int with range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    NativeVlan interface{}

    // Assign the access vlan to the access port. The type is one of the following
    // types: int with range: 1..4094, or string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*).
    AccessVlan interface{}

    // Specify VLANs, or ranges thereof, that the interface may carry when in
    // trunk mode.  If not specified, all VLANs are allowed on the interface.
    // Ranges are specified in the form x..y, where x<y - ranges are assumed to be
    // inclusive (such that the VLAN range is x <= range <= y. The type is one of
    // the following types: slice of int with range: 1..4094, or slice of string
    // with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]),
    // or slice of string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*),
    // or slice of string with pattern:
    // (409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.((409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])|\*),
    // or slice of string with pattern:
    // (\*|(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9]))\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])\.\.(409[0-4]|40[0-8][0-9]|[1-3][0-9]{3}|[1-9][0-9]{1,2}|[1-9])..
    TrunkVlans []interface{}
}

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetGoName(yname string) string {
    if yname == "interface-mode" { return "InterfaceMode" }
    if yname == "native-vlan" { return "NativeVlan" }
    if yname == "access-vlan" { return "AccessVlan" }
    if yname == "trunk-vlans" { return "TrunkVlans" }
    return ""
}

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-mode"] = state.InterfaceMode
    leafs["native-vlan"] = state.NativeVlan
    leafs["access-vlan"] = state.AccessVlan
    leafs["trunk-vlans"] = state.TrunkVlans
    return leafs
}

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_Aggregation_SwitchedVlan_State) GetParentYangName() string { return "switched-vlan" }

// Interfaces_Interface_RoutedVlan
// Top-level container for routed vlan interfaces.  These
// logical interfaces are also known as SVI (switched virtual
// interface), IRB (integrated routing and bridging), RVI
// (routed VLAN interface)
type Interfaces_Interface_RoutedVlan struct {
    parent types.Entity
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

func (routedVlan *Interfaces_Interface_RoutedVlan) GetFilter() yfilter.YFilter { return routedVlan.YFilter }

func (routedVlan *Interfaces_Interface_RoutedVlan) SetFilter(yf yfilter.YFilter) { routedVlan.YFilter = yf }

func (routedVlan *Interfaces_Interface_RoutedVlan) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "openconfig-if-ip:ipv4" { return "Ipv4" }
    if yname == "openconfig-if-ip:ipv6" { return "Ipv6" }
    return ""
}

func (routedVlan *Interfaces_Interface_RoutedVlan) GetSegmentPath() string {
    return "openconfig-vlan:routed-vlan"
}

func (routedVlan *Interfaces_Interface_RoutedVlan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &routedVlan.Config
    }
    if childYangName == "state" {
        return &routedVlan.State
    }
    if childYangName == "openconfig-if-ip:ipv4" {
        return &routedVlan.Ipv4
    }
    if childYangName == "openconfig-if-ip:ipv6" {
        return &routedVlan.Ipv6
    }
    return nil
}

func (routedVlan *Interfaces_Interface_RoutedVlan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &routedVlan.Config
    children["state"] = &routedVlan.State
    children["openconfig-if-ip:ipv4"] = &routedVlan.Ipv4
    children["openconfig-if-ip:ipv6"] = &routedVlan.Ipv6
    return children
}

func (routedVlan *Interfaces_Interface_RoutedVlan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routedVlan *Interfaces_Interface_RoutedVlan) GetBundleName() string { return "openconfig" }

func (routedVlan *Interfaces_Interface_RoutedVlan) GetYangName() string { return "routed-vlan" }

func (routedVlan *Interfaces_Interface_RoutedVlan) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (routedVlan *Interfaces_Interface_RoutedVlan) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (routedVlan *Interfaces_Interface_RoutedVlan) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (routedVlan *Interfaces_Interface_RoutedVlan) SetParent(parent types.Entity) { routedVlan.parent = parent }

func (routedVlan *Interfaces_Interface_RoutedVlan) GetParent() types.Entity { return routedVlan.parent }

func (routedVlan *Interfaces_Interface_RoutedVlan) GetParentYangName() string { return "interface" }

// Interfaces_Interface_RoutedVlan_Config
// Configuration data for routed vlan interfaces
type Interfaces_Interface_RoutedVlan_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References the VLAN for which this IP interface provides routing services
    // -- similar to a switch virtual interface (SVI), or integrated routing and
    // bridging interface (IRB) in some implementations. The type is one of the
    // following types: int with range: 0..65535, or string.
    Vlan interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Config) GetGoName(yname string) string {
    if yname == "vlan" { return "Vlan" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan"] = config.Vlan
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Config) GetParentYangName() string { return "routed-vlan" }

// Interfaces_Interface_RoutedVlan_State
// Operational state data 
type Interfaces_Interface_RoutedVlan_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // References the VLAN for which this IP interface provides routing services
    // -- similar to a switch virtual interface (SVI), or integrated routing and
    // bridging interface (IRB) in some implementations. The type is one of the
    // following types: int with range: 0..65535, or string.
    Vlan interface{}
}

func (state *Interfaces_Interface_RoutedVlan_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_State) GetGoName(yname string) string {
    if yname == "vlan" { return "Vlan" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vlan"] = state.Vlan
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_State) GetParentYangName() string { return "routed-vlan" }

// Interfaces_Interface_RoutedVlan_Ipv4
// Parameters for the IPv4 address family.
type Interfaces_Interface_RoutedVlan_Ipv4 struct {
    parent types.Entity
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

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetGoName(yname string) string {
    if yname == "addresses" { return "Addresses" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "unnumbered" { return "Unnumbered" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetSegmentPath() string {
    return "openconfig-if-ip:ipv4"
}

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &ipv4.Addresses
    }
    if childYangName == "neighbors" {
        return &ipv4.Neighbors
    }
    if childYangName == "unnumbered" {
        return &ipv4.Unnumbered
    }
    if childYangName == "config" {
        return &ipv4.Config
    }
    if childYangName == "state" {
        return &ipv4.State
    }
    return nil
}

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &ipv4.Addresses
    children["neighbors"] = &ipv4.Neighbors
    children["unnumbered"] = &ipv4.Unnumbered
    children["config"] = &ipv4.Config
    children["state"] = &ipv4.State
    return children
}

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetBundleName() string { return "openconfig" }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Interfaces_Interface_RoutedVlan_Ipv4) GetParentYangName() string { return "routed-vlan" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses
// Enclosing container for address list
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of configured IPv4 addresses on the interface. The type is slice
    // of Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address.
    Address []Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range addresses.Address {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address{}
        addresses.Address = append(addresses.Address, child)
        return &addresses.Address[len(addresses.Address)-1]
    }
    return nil
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.Address {
        children[addresses.Address[i].GetSegmentPath()] = &addresses.Address[i]
    }
    return children
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetBundleName() string { return "openconfig" }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetYangName() string { return "addresses" }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv4_Addresses) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address
// The list of configured IPv4 addresses on the interface.
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address struct {
    parent types.Entity
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

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "vrrp" { return "Vrrp" }
    return ""
}

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetSegmentPath() string {
    return "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
}

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &address.Config
    }
    if childYangName == "state" {
        return &address.State
    }
    if childYangName == "vrrp" {
        return &address.Vrrp
    }
    return nil
}

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &address.Config
    children["state"] = &address.State
    children["vrrp"] = &address.Vrrp
    return children
}

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = address.Ip
    return leafs
}

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetBundleName() string { return "openconfig" }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetYangName() string { return "address" }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetParent() types.Entity { return address.parent }

func (address *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address) GetParentYangName() string { return "addresses" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config
// Configuration data for each configured IPv4
// address on the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv4 address on the interface.
    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..32.
    PrefixLength interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = config.Ip
    leafs["prefix-length"] = config.PrefixLength
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Config) GetParentYangName() string { return "address" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State
// Operational state data for each IPv4 address
// configured on the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "origin" { return "Origin" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = state.Ip
    leafs["prefix-length"] = state.PrefixLength
    leafs["origin"] = state.Origin
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_State) GetParentYangName() string { return "address" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp
// Enclosing container for VRRP groups handled by this
// IP interface
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of VRRP groups, keyed by virtual router id. The type is slice of
    // Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup.
    VrrpGroup []Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetFilter() yfilter.YFilter { return vrrp.YFilter }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) SetFilter(yf yfilter.YFilter) { vrrp.YFilter = yf }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetGoName(yname string) string {
    if yname == "vrrp-group" { return "VrrpGroup" }
    return ""
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetSegmentPath() string {
    return "vrrp"
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrrp-group" {
        for _, c := range vrrp.VrrpGroup {
            if vrrp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup{}
        vrrp.VrrpGroup = append(vrrp.VrrpGroup, child)
        return &vrrp.VrrpGroup[len(vrrp.VrrpGroup)-1]
    }
    return nil
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrrp.VrrpGroup {
        children[vrrp.VrrpGroup[i].GetSegmentPath()] = &vrrp.VrrpGroup[i]
    }
    return children
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetBundleName() string { return "openconfig" }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetYangName() string { return "vrrp" }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) SetParent(parent types.Entity) { vrrp.parent = parent }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetParent() types.Entity { return vrrp.parent }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp) GetParentYangName() string { return "address" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup
// List of VRRP groups, keyed by virtual router id
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup struct {
    parent types.Entity
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

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetFilter() yfilter.YFilter { return vrrpGroup.YFilter }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) SetFilter(yf yfilter.YFilter) { vrrpGroup.YFilter = yf }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-tracking" { return "InterfaceTracking" }
    return ""
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetSegmentPath() string {
    return "vrrp-group" + "[virtual-router-id='" + fmt.Sprintf("%v", vrrpGroup.VirtualRouterId) + "']"
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &vrrpGroup.Config
    }
    if childYangName == "state" {
        return &vrrpGroup.State
    }
    if childYangName == "interface-tracking" {
        return &vrrpGroup.InterfaceTracking
    }
    return nil
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &vrrpGroup.Config
    children["state"] = &vrrpGroup.State
    children["interface-tracking"] = &vrrpGroup.InterfaceTracking
    return children
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = vrrpGroup.VirtualRouterId
    return leafs
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetBundleName() string { return "openconfig" }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetYangName() string { return "vrrp-group" }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) SetParent(parent types.Entity) { vrrpGroup.parent = parent }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetParent() types.Entity { return vrrpGroup.parent }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup) GetParentYangName() string { return "vrrp" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config
// Configuration data for the VRRP group
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "virtual-address" { return "VirtualAddress" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "accept-mode" { return "AcceptMode" }
    if yname == "advertisement-interval" { return "AdvertisementInterval" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = config.VirtualRouterId
    leafs["virtual-address"] = config.VirtualAddress
    leafs["priority"] = config.Priority
    leafs["preempt"] = config.Preempt
    leafs["preempt-delay"] = config.PreemptDelay
    leafs["accept-mode"] = config.AcceptMode
    leafs["advertisement-interval"] = config.AdvertisementInterval
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_Config) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State
// Operational state data for the VRRP group
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "virtual-address" { return "VirtualAddress" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "accept-mode" { return "AcceptMode" }
    if yname == "advertisement-interval" { return "AdvertisementInterval" }
    if yname == "current-priority" { return "CurrentPriority" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = state.VirtualRouterId
    leafs["virtual-address"] = state.VirtualAddress
    leafs["priority"] = state.Priority
    leafs["preempt"] = state.Preempt
    leafs["preempt-delay"] = state.PreemptDelay
    leafs["accept-mode"] = state.AcceptMode
    leafs["advertisement-interval"] = state.AdvertisementInterval
    leafs["current-priority"] = state.CurrentPriority
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_State) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
// Top-level container for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for VRRP interface tracking.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config

    // Operational state data for VRRP interface tracking.
    State Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetFilter() yfilter.YFilter { return interfaceTracking.YFilter }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) SetFilter(yf yfilter.YFilter) { interfaceTracking.YFilter = yf }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetSegmentPath() string {
    return "interface-tracking"
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceTracking.Config
    }
    if childYangName == "state" {
        return &interfaceTracking.State
    }
    return nil
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceTracking.Config
    children["state"] = &interfaceTracking.State
    return children
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetBundleName() string { return "openconfig" }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetYangName() string { return "interface-tracking" }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) SetParent(parent types.Entity) { interfaceTracking.parent = parent }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetParent() types.Entity { return interfaceTracking.parent }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config
// Configuration data for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetGoName(yname string) string {
    if yname == "track-interface" { return "TrackInterface" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-interface"] = config.TrackInterface
    leafs["priority-decrement"] = config.PriorityDecrement
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetParentYangName() string { return "interface-tracking" }

// Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
// Operational state data for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetGoName(yname string) string {
    if yname == "track-interface" { return "TrackInterface" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-interface"] = state.TrackInterface
    leafs["priority-decrement"] = state.PriorityDecrement
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetParentYangName() string { return "interface-tracking" }

// Interfaces_Interface_RoutedVlan_Ipv4_Neighbors
// Enclosing container for neighbor list
type Interfaces_Interface_RoutedVlan_Ipv4_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of mappings from IPv4 addresses to link-layer addresses.  Entries in
    // this list are used as static entries in the ARP Cache. The type is slice of
    // Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor.
    Neighbor []Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetBundleName() string { return "openconfig" }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor
// A list of mappings from IPv4 addresses to
// link-layer addresses.
// 
// Entries in this list are used as static entries in the
// ARP Cache.
type Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor struct {
    parent types.Entity
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

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &neighbor.Config
    }
    if childYangName == "state" {
        return &neighbor.State
    }
    return nil
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &neighbor.Config
    children["state"] = &neighbor.State
    return children
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = neighbor.Ip
    return leafs
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetBundleName() string { return "openconfig" }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config
// Configuration data for each configured IPv4
// address on the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPv4 address of the neighbor node. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The link-layer address of the neighbor node. The type is string with
    // pattern: ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?. This attribute is mandatory.
    LinkLayerAddress interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = config.Ip
    leafs["link-layer-address"] = config.LinkLayerAddress
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_Config) GetParentYangName() string { return "neighbor" }

// Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State
// Operational state data for each IPv4 address
// configured on the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    if yname == "origin" { return "Origin" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = state.Ip
    leafs["link-layer-address"] = state.LinkLayerAddress
    leafs["origin"] = state.Origin
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Neighbors_Neighbor_State) GetParentYangName() string { return "neighbor" }

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered
// Top-level container for setting unnumbered interfaces.
// Includes reference the interface that provides the
// address information
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for unnumbered interface.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config

    // Operational state data for unnumbered interfaces.
    State Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State

    // Reference to an interface or subinterface.
    InterfaceRef Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetFilter() yfilter.YFilter { return unnumbered.YFilter }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) SetFilter(yf yfilter.YFilter) { unnumbered.YFilter = yf }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-ref" { return "InterfaceRef" }
    return ""
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetSegmentPath() string {
    return "unnumbered"
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &unnumbered.Config
    }
    if childYangName == "state" {
        return &unnumbered.State
    }
    if childYangName == "interface-ref" {
        return &unnumbered.InterfaceRef
    }
    return nil
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &unnumbered.Config
    children["state"] = &unnumbered.State
    children["interface-ref"] = &unnumbered.InterfaceRef
    return children
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetBundleName() string { return "openconfig" }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetYangName() string { return "unnumbered" }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) SetParent(parent types.Entity) { unnumbered.parent = parent }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetParent() types.Entity { return unnumbered.parent }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config
// Configuration data for unnumbered interface
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_Config) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State
// Operational state data for unnumbered interfaces
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_State) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
// Reference to an interface or subinterface
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config

    // Operational state for interface-ref.
    State Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetFilter() yfilter.YFilter { return interfaceRef.YFilter }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) SetFilter(yf yfilter.YFilter) { interfaceRef.YFilter = yf }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetSegmentPath() string {
    return "interface-ref"
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceRef.Config
    }
    if childYangName == "state" {
        return &interfaceRef.State
    }
    return nil
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceRef.Config
    children["state"] = &interfaceRef.State
    return children
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetBundleName() string { return "openconfig" }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetYangName() string { return "interface-ref" }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) SetParent(parent types.Entity) { interfaceRef.parent = parent }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetParent() types.Entity { return interfaceRef.parent }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config
// Configured reference to interface / subinterface
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = config.Interface
    leafs["subinterface"] = config.Subinterface
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_Config) GetParentYangName() string { return "interface-ref" }

// Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State
// Operational state for interface-ref
type Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    leafs["subinterface"] = state.Subinterface
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_State) GetParentYangName() string { return "interface-ref" }

// Interfaces_Interface_RoutedVlan_Ipv4_Config
// Top-level IPv4 configuration data for the interface
type Interfaces_Interface_RoutedVlan_Ipv4_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "mtu" { return "Mtu" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["mtu"] = config.Mtu
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv4_Config) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_RoutedVlan_Ipv4_State
// Top level IPv4 operational state data
type Interfaces_Interface_RoutedVlan_Ipv4_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "mtu" { return "Mtu" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["mtu"] = state.Mtu
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv4_State) GetParentYangName() string { return "ipv4" }

// Interfaces_Interface_RoutedVlan_Ipv6
// Parameters for the IPv6 address family.
type Interfaces_Interface_RoutedVlan_Ipv6 struct {
    parent types.Entity
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

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetGoName(yname string) string {
    if yname == "addresses" { return "Addresses" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "unnumbered" { return "Unnumbered" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetSegmentPath() string {
    return "openconfig-if-ip:ipv6"
}

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &ipv6.Addresses
    }
    if childYangName == "neighbors" {
        return &ipv6.Neighbors
    }
    if childYangName == "unnumbered" {
        return &ipv6.Unnumbered
    }
    if childYangName == "config" {
        return &ipv6.Config
    }
    if childYangName == "state" {
        return &ipv6.State
    }
    return nil
}

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &ipv6.Addresses
    children["neighbors"] = &ipv6.Neighbors
    children["unnumbered"] = &ipv6.Unnumbered
    children["config"] = &ipv6.Config
    children["state"] = &ipv6.State
    return children
}

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetBundleName() string { return "openconfig" }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Interfaces_Interface_RoutedVlan_Ipv6) GetParentYangName() string { return "routed-vlan" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses
// Enclosing container for address list
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of configured IPv6 addresses on the interface. The type is slice
    // of Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address.
    Address []Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range addresses.Address {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address{}
        addresses.Address = append(addresses.Address, child)
        return &addresses.Address[len(addresses.Address)-1]
    }
    return nil
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.Address {
        children[addresses.Address[i].GetSegmentPath()] = &addresses.Address[i]
    }
    return children
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetBundleName() string { return "openconfig" }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetYangName() string { return "addresses" }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *Interfaces_Interface_RoutedVlan_Ipv6_Addresses) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address
// The list of configured IPv6 addresses on the interface.
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address struct {
    parent types.Entity
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

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "vrrp" { return "Vrrp" }
    return ""
}

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetSegmentPath() string {
    return "address" + "[ip='" + fmt.Sprintf("%v", address.Ip) + "']"
}

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &address.Config
    }
    if childYangName == "state" {
        return &address.State
    }
    if childYangName == "vrrp" {
        return &address.Vrrp
    }
    return nil
}

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &address.Config
    children["state"] = &address.State
    children["vrrp"] = &address.Vrrp
    return children
}

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = address.Ip
    return leafs
}

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetBundleName() string { return "openconfig" }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetYangName() string { return "address" }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetParent() types.Entity { return address.parent }

func (address *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address) GetParentYangName() string { return "addresses" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config
// Configuration data for each IPv6 address on
// the interface
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // [adapted from IETF IP model RFC 7277]  The IPv6 address on the interface.
    // The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // [adapted from IETF IP model RFC 7277]  The length of the subnet prefix. The
    // type is interface{} with range: 0..128. This attribute is mandatory.
    PrefixLength interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = config.Ip
    leafs["prefix-length"] = config.PrefixLength
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Config) GetParentYangName() string { return "address" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State
// State data for each IPv6 address on the
// interface
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "origin" { return "Origin" }
    if yname == "status" { return "Status" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = state.Ip
    leafs["prefix-length"] = state.PrefixLength
    leafs["origin"] = state.Origin
    leafs["status"] = state.Status
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_State) GetParentYangName() string { return "address" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // List of VRRP groups, keyed by virtual router id. The type is slice of
    // Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup.
    VrrpGroup []Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetFilter() yfilter.YFilter { return vrrp.YFilter }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) SetFilter(yf yfilter.YFilter) { vrrp.YFilter = yf }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetGoName(yname string) string {
    if yname == "vrrp-group" { return "VrrpGroup" }
    return ""
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetSegmentPath() string {
    return "vrrp"
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrrp-group" {
        for _, c := range vrrp.VrrpGroup {
            if vrrp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup{}
        vrrp.VrrpGroup = append(vrrp.VrrpGroup, child)
        return &vrrp.VrrpGroup[len(vrrp.VrrpGroup)-1]
    }
    return nil
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrrp.VrrpGroup {
        children[vrrp.VrrpGroup[i].GetSegmentPath()] = &vrrp.VrrpGroup[i]
    }
    return children
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetBundleName() string { return "openconfig" }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetYangName() string { return "vrrp" }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) SetParent(parent types.Entity) { vrrp.parent = parent }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetParent() types.Entity { return vrrp.parent }

func (vrrp *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp) GetParentYangName() string { return "address" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup
// List of VRRP groups, keyed by virtual router id
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup struct {
    parent types.Entity
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

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetFilter() yfilter.YFilter { return vrrpGroup.YFilter }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) SetFilter(yf yfilter.YFilter) { vrrpGroup.YFilter = yf }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-tracking" { return "InterfaceTracking" }
    return ""
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetSegmentPath() string {
    return "vrrp-group" + "[virtual-router-id='" + fmt.Sprintf("%v", vrrpGroup.VirtualRouterId) + "']"
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &vrrpGroup.Config
    }
    if childYangName == "state" {
        return &vrrpGroup.State
    }
    if childYangName == "interface-tracking" {
        return &vrrpGroup.InterfaceTracking
    }
    return nil
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &vrrpGroup.Config
    children["state"] = &vrrpGroup.State
    children["interface-tracking"] = &vrrpGroup.InterfaceTracking
    return children
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = vrrpGroup.VirtualRouterId
    return leafs
}

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetBundleName() string { return "openconfig" }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetYangName() string { return "vrrp-group" }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) SetParent(parent types.Entity) { vrrpGroup.parent = parent }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetParent() types.Entity { return vrrpGroup.parent }

func (vrrpGroup *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup) GetParentYangName() string { return "vrrp" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config
// Configuration data for the VRRP group
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "virtual-address" { return "VirtualAddress" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "accept-mode" { return "AcceptMode" }
    if yname == "advertisement-interval" { return "AdvertisementInterval" }
    if yname == "virtual-link-local" { return "VirtualLinkLocal" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = config.VirtualRouterId
    leafs["virtual-address"] = config.VirtualAddress
    leafs["priority"] = config.Priority
    leafs["preempt"] = config.Preempt
    leafs["preempt-delay"] = config.PreemptDelay
    leafs["accept-mode"] = config.AcceptMode
    leafs["advertisement-interval"] = config.AdvertisementInterval
    leafs["virtual-link-local"] = config.VirtualLinkLocal
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_Config) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State
// Operational state data for the VRRP group
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetGoName(yname string) string {
    if yname == "virtual-router-id" { return "VirtualRouterId" }
    if yname == "virtual-address" { return "VirtualAddress" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "preempt-delay" { return "PreemptDelay" }
    if yname == "accept-mode" { return "AcceptMode" }
    if yname == "advertisement-interval" { return "AdvertisementInterval" }
    if yname == "current-priority" { return "CurrentPriority" }
    if yname == "virtual-link-local" { return "VirtualLinkLocal" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-router-id"] = state.VirtualRouterId
    leafs["virtual-address"] = state.VirtualAddress
    leafs["priority"] = state.Priority
    leafs["preempt"] = state.Preempt
    leafs["preempt-delay"] = state.PreemptDelay
    leafs["accept-mode"] = state.AcceptMode
    leafs["advertisement-interval"] = state.AdvertisementInterval
    leafs["current-priority"] = state.CurrentPriority
    leafs["virtual-link-local"] = state.VirtualLinkLocal
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_State) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking
// Top-level container for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for VRRP interface tracking.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config

    // Operational state data for VRRP interface tracking.
    State Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetFilter() yfilter.YFilter { return interfaceTracking.YFilter }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) SetFilter(yf yfilter.YFilter) { interfaceTracking.YFilter = yf }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetSegmentPath() string {
    return "interface-tracking"
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceTracking.Config
    }
    if childYangName == "state" {
        return &interfaceTracking.State
    }
    return nil
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceTracking.Config
    children["state"] = &interfaceTracking.State
    return children
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetBundleName() string { return "openconfig" }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetYangName() string { return "interface-tracking" }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) SetParent(parent types.Entity) { interfaceTracking.parent = parent }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetParent() types.Entity { return interfaceTracking.parent }

func (interfaceTracking *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking) GetParentYangName() string { return "vrrp-group" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config
// Configuration data for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetGoName(yname string) string {
    if yname == "track-interface" { return "TrackInterface" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-interface"] = config.TrackInterface
    leafs["priority-decrement"] = config.PriorityDecrement
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_Config) GetParentYangName() string { return "interface-tracking" }

// Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State
// Operational state data for VRRP interface tracking
type Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetGoName(yname string) string {
    if yname == "track-interface" { return "TrackInterface" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["track-interface"] = state.TrackInterface
    leafs["priority-decrement"] = state.PriorityDecrement
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Addresses_Address_Vrrp_VrrpGroup_InterfaceTracking_State) GetParentYangName() string { return "interface-tracking" }

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors
// Enclosing container for list of IPv6 neighbors
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of IPv6 neighbors. The type is slice of
    // Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor.
    Neighbor []Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    return nil
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    return children
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetBundleName() string { return "openconfig" }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor
// List of IPv6 neighbors
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor struct {
    parent types.Entity
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

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[ip='" + fmt.Sprintf("%v", neighbor.Ip) + "']"
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &neighbor.Config
    }
    if childYangName == "state" {
        return &neighbor.State
    }
    return nil
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &neighbor.Config
    children["state"] = &neighbor.State
    return children
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = neighbor.Ip
    return leafs
}

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetBundleName() string { return "openconfig" }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config
// Configuration data for each IPv6 address on
// the interface
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = config.Ip
    leafs["link-layer-address"] = config.LinkLayerAddress
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_Config) GetParentYangName() string { return "neighbor" }

// Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State
// State data for each IPv6 address on the
// interface
type Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "link-layer-address" { return "LinkLayerAddress" }
    if yname == "origin" { return "Origin" }
    if yname == "is-router" { return "IsRouter" }
    if yname == "neighbor-state" { return "NeighborState" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = state.Ip
    leafs["link-layer-address"] = state.LinkLayerAddress
    leafs["origin"] = state.Origin
    leafs["is-router"] = state.IsRouter
    leafs["neighbor-state"] = state.NeighborState
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Neighbors_Neighbor_State) GetParentYangName() string { return "neighbor" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration data for unnumbered interface.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config

    // Operational state data for unnumbered interfaces.
    State Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State

    // Reference to an interface or subinterface.
    InterfaceRef Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetFilter() yfilter.YFilter { return unnumbered.YFilter }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) SetFilter(yf yfilter.YFilter) { unnumbered.YFilter = yf }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    if yname == "interface-ref" { return "InterfaceRef" }
    return ""
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetSegmentPath() string {
    return "unnumbered"
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &unnumbered.Config
    }
    if childYangName == "state" {
        return &unnumbered.State
    }
    if childYangName == "interface-ref" {
        return &unnumbered.InterfaceRef
    }
    return nil
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &unnumbered.Config
    children["state"] = &unnumbered.State
    children["interface-ref"] = &unnumbered.InterfaceRef
    return children
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetBundleName() string { return "openconfig" }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetYangName() string { return "unnumbered" }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) SetParent(parent types.Entity) { unnumbered.parent = parent }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetParent() types.Entity { return unnumbered.parent }

func (unnumbered *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config
// Configuration data for unnumbered interface
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_Config) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State
// Operational state data for unnumbered interfaces
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Indicates that the subinterface is unnumbered.  By default the subinterface
    // is numbered, i.e., expected to have an IP address configuration. The type
    // is bool. The default value is false.
    Enabled interface{}
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_State) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef
// Reference to an interface or subinterface
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configured reference to interface / subinterface.
    Config Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config

    // Operational state for interface-ref.
    State Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetFilter() yfilter.YFilter { return interfaceRef.YFilter }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) SetFilter(yf yfilter.YFilter) { interfaceRef.YFilter = yf }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetGoName(yname string) string {
    if yname == "config" { return "Config" }
    if yname == "state" { return "State" }
    return ""
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetSegmentPath() string {
    return "interface-ref"
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &interfaceRef.Config
    }
    if childYangName == "state" {
        return &interfaceRef.State
    }
    return nil
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &interfaceRef.Config
    children["state"] = &interfaceRef.State
    return children
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetBundleName() string { return "openconfig" }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetYangName() string { return "interface-ref" }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) SetParent(parent types.Entity) { interfaceRef.parent = parent }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetParent() types.Entity { return interfaceRef.parent }

func (interfaceRef *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef) GetParentYangName() string { return "unnumbered" }

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config
// Configured reference to interface / subinterface
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = config.Interface
    leafs["subinterface"] = config.Subinterface
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_Config) GetParentYangName() string { return "interface-ref" }

// Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State
// Operational state for interface-ref
type Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "subinterface" { return "Subinterface" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = state.Interface
    leafs["subinterface"] = state.Subinterface
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_Unnumbered_InterfaceRef_State) GetParentYangName() string { return "interface-ref" }

// Interfaces_Interface_RoutedVlan_Ipv6_Config
// Top-level config data for the IPv6 interface
type Interfaces_Interface_RoutedVlan_Ipv6_Config struct {
    parent types.Entity
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

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "mtu" { return "Mtu" }
    if yname == "dup-addr-detect-transmits" { return "DupAddrDetectTransmits" }
    return ""
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetSegmentPath() string {
    return "config"
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = config.Enabled
    leafs["mtu"] = config.Mtu
    leafs["dup-addr-detect-transmits"] = config.DupAddrDetectTransmits
    return leafs
}

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetBundleName() string { return "openconfig" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetYangName() string { return "config" }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetParent() types.Entity { return config.parent }

func (config *Interfaces_Interface_RoutedVlan_Ipv6_Config) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_RoutedVlan_Ipv6_State
// Top-level operational state data for the IPv6 interface
type Interfaces_Interface_RoutedVlan_Ipv6_State struct {
    parent types.Entity
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

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "mtu" { return "Mtu" }
    if yname == "dup-addr-detect-transmits" { return "DupAddrDetectTransmits" }
    return ""
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetSegmentPath() string {
    return "state"
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = state.Enabled
    leafs["mtu"] = state.Mtu
    leafs["dup-addr-detect-transmits"] = state.DupAddrDetectTransmits
    return leafs
}

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetBundleName() string { return "openconfig" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetYangName() string { return "state" }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetParent() types.Entity { return state.parent }

func (state *Interfaces_Interface_RoutedVlan_Ipv6_State) GetParentYangName() string { return "ipv6" }

// Interfaces_Interface_Sonet
// Data related to SONET/SDH interfaces
type Interfaces_Interface_Sonet struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (sonet *Interfaces_Interface_Sonet) GetFilter() yfilter.YFilter { return sonet.YFilter }

func (sonet *Interfaces_Interface_Sonet) SetFilter(yf yfilter.YFilter) { sonet.YFilter = yf }

func (sonet *Interfaces_Interface_Sonet) GetGoName(yname string) string {
    return ""
}

func (sonet *Interfaces_Interface_Sonet) GetSegmentPath() string {
    return "openconfig-transport-line-common:sonet"
}

func (sonet *Interfaces_Interface_Sonet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sonet *Interfaces_Interface_Sonet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sonet *Interfaces_Interface_Sonet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sonet *Interfaces_Interface_Sonet) GetBundleName() string { return "openconfig" }

func (sonet *Interfaces_Interface_Sonet) GetYangName() string { return "sonet" }

func (sonet *Interfaces_Interface_Sonet) GetBundleYangModelsLocation() string { return openconfig.GetModelsPath() }

func (sonet *Interfaces_Interface_Sonet) GetCapabilitiesTable() map[string]string {
    return openconfig.GetCapabilities() }

func (sonet *Interfaces_Interface_Sonet) GetNamespaceTable() map[string]string {
    return openconfig.GetNamespaces() }

func (sonet *Interfaces_Interface_Sonet) SetParent(parent types.Entity) { sonet.parent = parent }

func (sonet *Interfaces_Interface_Sonet) GetParent() types.Entity { return sonet.parent }

func (sonet *Interfaces_Interface_Sonet) GetParentYangName() string { return "interface" }

