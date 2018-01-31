// The MIB module to describe generic objects for network
// interface sub-layers.  This MIB is an updated version of
// MIB-II's ifTable, and incorporates the extensions defined in
// RFC 1229.
package if_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package if_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:IF-MIB IF-MIB}", reflect.TypeOf(IFMIB{}))
    ydk.RegisterEntity("IF-MIB:IF-MIB", reflect.TypeOf(IFMIB{}))
}

// IFMIB
type IFMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Interfaces IFMIB_Interfaces

    
    Ifmibobjects IFMIB_Ifmibobjects

    // A list of interface entries.  The number of entries is given by the value
    // of ifNumber.
    Iftable IFMIB_Iftable

    // The table containing information on the relationships between the multiple
    // sub-layers of network interfaces.  In particular, it contains information
    // on which sub-layers run 'on top of' which other sub-layers, where each
    // sub-layer corresponds to a conceptual row in the ifTable.  For example,
    // when the sub-layer with ifIndex value x runs over the sub-layer with
    // ifIndex value y, then this table contains:    ifStackStatus.x.y=active  For
    // each ifIndex value, I, which identifies an active interface, there are
    // always at least two instantiated rows in this table associated with I.  For
    // one of these rows, I is the value of ifStackHigherLayer; for the other, I
    // is the value of ifStackLowerLayer.  (If I is not involved in multiplexing,
    // then these are the only two rows associated with I.)  For example, two rows
    // exist even for an interface which has no others stacked on top or below it:
    // ifStackStatus.0.x=active   ifStackStatus.x.0=active .
    Ifstacktable IFMIB_Ifstacktable

    // This table contains an entry for each address (broadcast, multicast, or
    // uni-cast) for which the system will receive packets/frames on a particular
    // interface, except as follows:  - for an interface operating in promiscuous
    // mode, entries are only required for those addresses for which the system
    // would receive frames were it not operating in promiscuous mode.  - for
    // 802.5 functional addresses, only one entry is required, for the address
    // which has the functional address bit ANDed with the bit mask of all
    // functional addresses for which the interface will accept frames.  A system
    // is normally able to use any unicast address which corresponds to an entry
    // in this table as a source address.
    Ifrcvaddresstable IFMIB_Ifrcvaddresstable
}

func (iFMIB *IFMIB) GetFilter() yfilter.YFilter { return iFMIB.YFilter }

func (iFMIB *IFMIB) SetFilter(yf yfilter.YFilter) { iFMIB.YFilter = yf }

func (iFMIB *IFMIB) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "ifMIBObjects" { return "Ifmibobjects" }
    if yname == "ifTable" { return "Iftable" }
    if yname == "ifStackTable" { return "Ifstacktable" }
    if yname == "ifRcvAddressTable" { return "Ifrcvaddresstable" }
    return ""
}

func (iFMIB *IFMIB) GetSegmentPath() string {
    return "IF-MIB:IF-MIB"
}

func (iFMIB *IFMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &iFMIB.Interfaces
    }
    if childYangName == "ifMIBObjects" {
        return &iFMIB.Ifmibobjects
    }
    if childYangName == "ifTable" {
        return &iFMIB.Iftable
    }
    if childYangName == "ifStackTable" {
        return &iFMIB.Ifstacktable
    }
    if childYangName == "ifRcvAddressTable" {
        return &iFMIB.Ifrcvaddresstable
    }
    return nil
}

func (iFMIB *IFMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &iFMIB.Interfaces
    children["ifMIBObjects"] = &iFMIB.Ifmibobjects
    children["ifTable"] = &iFMIB.Iftable
    children["ifStackTable"] = &iFMIB.Ifstacktable
    children["ifRcvAddressTable"] = &iFMIB.Ifrcvaddresstable
    return children
}

func (iFMIB *IFMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iFMIB *IFMIB) GetBundleName() string { return "cisco_ios_xe" }

func (iFMIB *IFMIB) GetYangName() string { return "IF-MIB" }

func (iFMIB *IFMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iFMIB *IFMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iFMIB *IFMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iFMIB *IFMIB) SetParent(parent types.Entity) { iFMIB.parent = parent }

func (iFMIB *IFMIB) GetParent() types.Entity { return iFMIB.parent }

func (iFMIB *IFMIB) GetParentYangName() string { return "IF-MIB" }

// IFMIB_Interfaces
type IFMIB_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of network interfaces (regardless of their current state)
    // present on this system. The type is interface{} with range:
    // -2147483648..2147483647.
    Ifnumber interface{}
}

func (interfaces *IFMIB_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *IFMIB_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *IFMIB_Interfaces) GetGoName(yname string) string {
    if yname == "ifNumber" { return "Ifnumber" }
    return ""
}

func (interfaces *IFMIB_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *IFMIB_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaces *IFMIB_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaces *IFMIB_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifNumber"] = interfaces.Ifnumber
    return leafs
}

func (interfaces *IFMIB_Interfaces) GetBundleName() string { return "cisco_ios_xe" }

func (interfaces *IFMIB_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *IFMIB_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (interfaces *IFMIB_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (interfaces *IFMIB_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (interfaces *IFMIB_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *IFMIB_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *IFMIB_Interfaces) GetParentYangName() string { return "IF-MIB" }

// IFMIB_Ifmibobjects
type IFMIB_Ifmibobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time of the last creation or deletion of an
    // entry in the ifTable.  If the number of entries has been unchanged since
    // the last re-initialization of the local network management subsystem, then
    // this object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    Iftablelastchange interface{}

    // The value of sysUpTime at the time of the last change of the (whole)
    // interface stack.  A change of the interface stack is defined to be any
    // creation, deletion, or change in value of any instance of ifStackStatus. 
    // If the interface stack has been unchanged since the last re-initialization
    // of the local network management subsystem, then this object contains a zero
    // value. The type is interface{} with range: 0..4294967295.
    Ifstacklastchange interface{}
}

func (ifmibobjects *IFMIB_Ifmibobjects) GetFilter() yfilter.YFilter { return ifmibobjects.YFilter }

func (ifmibobjects *IFMIB_Ifmibobjects) SetFilter(yf yfilter.YFilter) { ifmibobjects.YFilter = yf }

func (ifmibobjects *IFMIB_Ifmibobjects) GetGoName(yname string) string {
    if yname == "ifTableLastChange" { return "Iftablelastchange" }
    if yname == "ifStackLastChange" { return "Ifstacklastchange" }
    return ""
}

func (ifmibobjects *IFMIB_Ifmibobjects) GetSegmentPath() string {
    return "ifMIBObjects"
}

func (ifmibobjects *IFMIB_Ifmibobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifmibobjects *IFMIB_Ifmibobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifmibobjects *IFMIB_Ifmibobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifTableLastChange"] = ifmibobjects.Iftablelastchange
    leafs["ifStackLastChange"] = ifmibobjects.Ifstacklastchange
    return leafs
}

func (ifmibobjects *IFMIB_Ifmibobjects) GetBundleName() string { return "cisco_ios_xe" }

func (ifmibobjects *IFMIB_Ifmibobjects) GetYangName() string { return "ifMIBObjects" }

func (ifmibobjects *IFMIB_Ifmibobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ifmibobjects *IFMIB_Ifmibobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ifmibobjects *IFMIB_Ifmibobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ifmibobjects *IFMIB_Ifmibobjects) SetParent(parent types.Entity) { ifmibobjects.parent = parent }

func (ifmibobjects *IFMIB_Ifmibobjects) GetParent() types.Entity { return ifmibobjects.parent }

func (ifmibobjects *IFMIB_Ifmibobjects) GetParentYangName() string { return "IF-MIB" }

// IFMIB_Iftable
// A list of interface entries.  The number of entries is
// given by the value of ifNumber.
type IFMIB_Iftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing management information applicable to a particular
    // interface. The type is slice of IFMIB_Iftable_Ifentry.
    Ifentry []IFMIB_Iftable_Ifentry
}

func (iftable *IFMIB_Iftable) GetFilter() yfilter.YFilter { return iftable.YFilter }

func (iftable *IFMIB_Iftable) SetFilter(yf yfilter.YFilter) { iftable.YFilter = yf }

func (iftable *IFMIB_Iftable) GetGoName(yname string) string {
    if yname == "ifEntry" { return "Ifentry" }
    return ""
}

func (iftable *IFMIB_Iftable) GetSegmentPath() string {
    return "ifTable"
}

func (iftable *IFMIB_Iftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ifEntry" {
        for _, c := range iftable.Ifentry {
            if iftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IFMIB_Iftable_Ifentry{}
        iftable.Ifentry = append(iftable.Ifentry, child)
        return &iftable.Ifentry[len(iftable.Ifentry)-1]
    }
    return nil
}

func (iftable *IFMIB_Iftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range iftable.Ifentry {
        children[iftable.Ifentry[i].GetSegmentPath()] = &iftable.Ifentry[i]
    }
    return children
}

func (iftable *IFMIB_Iftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iftable *IFMIB_Iftable) GetBundleName() string { return "cisco_ios_xe" }

func (iftable *IFMIB_Iftable) GetYangName() string { return "ifTable" }

func (iftable *IFMIB_Iftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iftable *IFMIB_Iftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iftable *IFMIB_Iftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iftable *IFMIB_Iftable) SetParent(parent types.Entity) { iftable.parent = parent }

func (iftable *IFMIB_Iftable) GetParent() types.Entity { return iftable.parent }

func (iftable *IFMIB_Iftable) GetParentYangName() string { return "IF-MIB" }

// IFMIB_Iftable_Ifentry
// An entry containing management information applicable to a
// particular interface.
type IFMIB_Iftable_Ifentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A unique value, greater than zero, for each
    // interface.  It is recommended that values are assigned contiguously
    // starting from 1.  The value for each interface sub-layer must remain
    // constant at least from one re-initialization of the entity's network
    // management system to the next re- initialization. The type is interface{}
    // with range: 1..2147483647.
    Ifindex interface{}

    // A textual string containing information about the interface.  This string
    // should include the name of the manufacturer, the product name and the
    // version of the interface hardware/software. The type is string with length:
    // 0..255.
    Ifdescr interface{}

    // The type of interface.  Additional values for ifType are assigned by the
    // Internet Assigned Numbers Authority (IANA), through updating the syntax of
    // the IANAifType textual convention. The type is IANAifType.
    Iftype interface{}

    // The size of the largest packet which can be sent/received on the interface,
    // specified in octets.  For interfaces that are used for transmitting network
    // datagrams, this is the size of the largest network datagram that can be
    // sent on the interface. The type is interface{} with range:
    // -2147483648..2147483647.
    Ifmtu interface{}

    // An estimate of the interface's current bandwidth in bits per second.  For
    // interfaces which do not vary in bandwidth or for those where no accurate
    // estimation can be made, this object should contain the nominal bandwidth. 
    // If the bandwidth of the interface is greater than the maximum value
    // reportable by this object then this object should report its maximum value
    // (4,294,967,295) and ifHighSpeed must be used to report the interace's
    // speed.  For a sub-layer which has no concept of bandwidth, this object
    // should be zero. The type is interface{} with range: 0..4294967295.
    Ifspeed interface{}

    // The interface's address at its protocol sub-layer.  For example, for an
    // 802.x interface, this object normally contains a MAC address.  The
    // interface's media-specific MIB must define the bit and byte ordering and
    // the format of the value of this object.  For interfaces which do not have
    // such an address (e.g., a serial line), this object should contain an octet
    // string of zero length. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Ifphysaddress interface{}

    // The desired state of the interface.  The testing(3) state indicates that no
    // operational packets can be passed.  When a managed system initializes, all
    // interfaces start with ifAdminStatus in the down(2) state.  As a result of
    // either explicit management action or per configuration information retained
    // by the managed system, ifAdminStatus is then changed to either the up(1) or
    // testing(3) states (or remains in the down(2) state). The type is
    // Ifadminstatus.
    Ifadminstatus interface{}

    // The current operational state of the interface.  The testing(3) state
    // indicates that no operational packets can be passed.  If ifAdminStatus is
    // down(2) then ifOperStatus should be down(2).  If ifAdminStatus is changed
    // to up(1) then ifOperStatus should change to up(1) if the interface is ready
    // to transmit and receive network traffic; it should change to dormant(5) if
    // the interface is waiting for external actions (such as a serial line
    // waiting for an incoming connection); it should remain in the down(2) state
    // if and only if there is a fault that prevents it from going to the up(1)
    // state; it should remain in the notPresent(6) state if the interface has
    // missing (typically, hardware) components. The type is Ifoperstatus.
    Ifoperstatus interface{}

    // The value of sysUpTime at the time the interface entered its current
    // operational state.  If the current state was entered prior to the last
    // re-initialization of the local network management subsystem, then this
    // object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    Iflastchange interface{}

    // The total number of octets received on the interface, including framing
    // characters.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Ifinoctets interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were not addressed to a multicast or broadcast address at this
    // sub-layer.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Ifinucastpkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a multicast or broadcast address at this sub-layer.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime.  This object is deprecated in favour of
    // ifInMulticastPkts and ifInBroadcastPkts. The type is interface{} with
    // range: 0..4294967295.
    Ifinnucastpkts interface{}

    // The number of inbound packets which were chosen to be discarded even though
    // no errors had been detected to prevent their being deliverable to a
    // higher-layer protocol.  One possible reason for discarding such a packet
    // could be to free up buffer space.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ifindiscards interface{}

    // For packet-oriented interfaces, the number of inbound packets that
    // contained errors preventing them from being deliverable to a higher-layer
    // protocol.  For character- oriented or fixed-length interfaces, the number
    // of inbound transmission units that contained errors preventing them from
    // being deliverable to a higher-layer protocol.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ifinerrors interface{}

    // For packet-oriented interfaces, the number of packets received via the
    // interface which were discarded because of an unknown or unsupported
    // protocol.  For character-oriented or fixed-length interfaces that support
    // protocol multiplexing the number of transmission units received via the
    // interface which were discarded because of an unknown or unsupported
    // protocol.  For any interface that does not support protocol multiplexing,
    // this counter will always be 0.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    Ifinunknownprotos interface{}

    // The total number of octets transmitted out of the interface, including
    // framing characters.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    Ifoutoctets interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were not addressed to a multicast or broadcast
    // address at this sub-layer, including those that were discarded or not sent.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ifoutucastpkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a multicast or broadcast address
    // at this sub-layer, including those that were discarded or not sent. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime.  This object is deprecated in favour of
    // ifOutMulticastPkts and ifOutBroadcastPkts. The type is interface{} with
    // range: 0..4294967295.
    Ifoutnucastpkts interface{}

    // The number of outbound packets which were chosen to be discarded even
    // though no errors had been detected to prevent their being transmitted.  One
    // possible reason for discarding such a packet could be to free up buffer
    // space.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    Ifoutdiscards interface{}

    // For packet-oriented interfaces, the number of outbound packets that could
    // not be transmitted because of errors. For character-oriented or
    // fixed-length interfaces, the number of outbound transmission units that
    // could not be transmitted because of errors.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    Ifouterrors interface{}

    // The length of the output packet queue (in packets). The type is interface{}
    // with range: 0..4294967295.
    Ifoutqlen interface{}

    // A reference to MIB definitions specific to the particular media being used
    // to realize the interface.  It is recommended that this value point to an
    // instance of a MIB object in the media-specific MIB, i.e., that this object
    // have the semantics associated with the InstancePointer textual convention
    // defined in RFC 2579.  In fact, it is recommended that the media-specific
    // MIB specify what value ifSpecific should/can take for values of ifType.  If
    // no MIB definitions specific to the particular media are available, the
    // value should be set to the OBJECT IDENTIFIER { 0 0 }. The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Ifspecific interface{}

    // The textual name of the interface.  The value of this object should be the
    // name of the interface as assigned by the local device and should be
    // suitable for use in commands entered at the device's `console'.  This might
    // be a text name, such as `le0' or a simple port number, such as `1',
    // depending on the interface naming syntax of the device.  If several entries
    // in the ifTable together represent a single interface as named by the
    // device, then each will have the same value of ifName.  Note that for an
    // agent which responds to SNMP queries concerning an interface on some other
    // (proxied) device, then the value of ifName for such an interface is the
    // proxied device's local name for it.  If there is no local name, or this
    // object is otherwise not applicable, then this object contains a zero-length
    // string. The type is string.
    Ifname interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a multicast address at this sub-layer.  For a MAC
    // layer protocol, this includes both Group and Functional addresses. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ifinmulticastpkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a broadcast address at this sub-layer. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ifinbroadcastpkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a multicast address at this
    // sub-layer, including those that were discarded or not sent.  For a MAC
    // layer protocol, this includes both Group and Functional addresses. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ifoutmulticastpkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a broadcast address at this
    // sub-layer, including those that were discarded or not sent. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    Ifoutbroadcastpkts interface{}

    // The total number of octets received on the interface, including framing
    // characters.  This object is a 64-bit version of ifInOctets. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ifhcinoctets interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were not addressed to a multicast or broadcast address at this
    // sub-layer.  This object is a 64-bit version of ifInUcastPkts. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ifhcinucastpkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a multicast address at this sub-layer.  For a MAC
    // layer protocol, this includes both Group and Functional addresses.  This
    // object is a 64-bit version of ifInMulticastPkts.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ifhcinmulticastpkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a broadcast address at this sub-layer.  This object
    // is a 64-bit version of ifInBroadcastPkts.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..18446744073709551615.
    Ifhcinbroadcastpkts interface{}

    // The total number of octets transmitted out of the interface, including
    // framing characters.  This object is a 64-bit version of ifOutOctets. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ifhcoutoctets interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were not addressed to a multicast or broadcast
    // address at this sub-layer, including those that were discarded or not sent.
    // This object is a 64-bit version of ifOutUcastPkts.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ifhcoutucastpkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a multicast address at this
    // sub-layer, including those that were discarded or not sent.  For a MAC
    // layer protocol, this includes both Group and Functional addresses.  This
    // object is a 64-bit version of ifOutMulticastPkts.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    Ifhcoutmulticastpkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a broadcast address at this
    // sub-layer, including those that were discarded or not sent.  This object is
    // a 64-bit version of ifOutBroadcastPkts.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..18446744073709551615.
    Ifhcoutbroadcastpkts interface{}

    // Indicates whether linkUp/linkDown traps should be generated for this
    // interface.  By default, this object should have the value enabled(1) for
    // interfaces which do not operate on 'top' of any other interface (as defined
    // in the ifStackTable), and disabled(2) otherwise. The type is
    // Iflinkupdowntrapenable.
    Iflinkupdowntrapenable interface{}

    // An estimate of the interface's current bandwidth in units of 1,000,000 bits
    // per second.  If this object reports a value of `n' then the speed of the
    // interface is somewhere in the range of `n-500,000' to `n+499,999'.  For
    // interfaces which do not vary in bandwidth or for those where no accurate
    // estimation can be made, this object should contain the nominal bandwidth. 
    // For a sub-layer which has no concept of bandwidth, this object should be
    // zero. The type is interface{} with range: 0..4294967295.
    Ifhighspeed interface{}

    // This object has a value of false(2) if this interface only accepts
    // packets/frames that are addressed to this station. This object has a value
    // of true(1) when the station accepts all packets/frames transmitted on the
    // media.  The value true(1) is only legal on certain types of media.  If
    // legal, setting this object to a value of true(1) may require the interface
    // to be reset before becoming effective.  The value of ifPromiscuousMode does
    // not affect the reception of broadcast and multicast packets/frames by the
    // interface. The type is bool.
    Ifpromiscuousmode interface{}

    // This object has the value 'true(1)' if the interface sublayer has a
    // physical connector and the value 'false(2)' otherwise. The type is bool.
    Ifconnectorpresent interface{}

    // This object is an 'alias' name for the interface as specified by a network
    // manager, and provides a non-volatile 'handle' for the interface.  On the
    // first instantiation of an interface, the value of ifAlias associated with
    // that interface is the zero-length string.  As and when a value is written
    // into an instance of ifAlias through a network management set operation,
    // then the agent must retain the supplied value in the ifAlias instance
    // associated with the same interface for as long as that interface remains
    // instantiated, including across all re- initializations/reboots of the
    // network management system, including those which result in a change of the
    // interface's ifIndex value.  An example of the value which a network manager
    // might store in this object for a WAN interface is the (Telco's) circuit
    // number/identifier of the interface.  Some agents may support write-access
    // only for interfaces having particular values of ifType.  An agent which
    // supports write access to this object is required to keep the value in
    // non-volatile storage, but it may limit the length of new values depending
    // on how much storage is already occupied by the current values for other
    // interfaces. The type is string with length: 0..64.
    Ifalias interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this interface's counters suffered a discontinuity.  The relevant
    // counters are the specific instances associated with this interface of any
    // Counter32 or Counter64 object contained in the ifTable or ifXTable.  If no
    // such discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    Ifcounterdiscontinuitytime interface{}

    // This object identifies the current invocation of the interface's test. The
    // type is interface{} with range: 0..2147483647.
    Iftestid interface{}

    // This object indicates whether or not some manager currently has the
    // necessary 'ownership' required to invoke a test on this interface.  A write
    // to this object is only successful when it changes its value from
    // 'notInUse(1)' to 'inUse(2)'. After completion of a test, the agent resets
    // the value back to 'notInUse(1)'. The type is Ifteststatus.
    Ifteststatus interface{}

    // A control variable used to start and stop operator- initiated interface
    // tests.  Most OBJECT IDENTIFIER values assigned to tests are defined
    // elsewhere, in association with specific types of interface.  However, this
    // document assigns a value for a full-duplex loopback test, and defines the
    // special meanings of the subject identifier:      noTest  OBJECT IDENTIFIER
    // ::= { 0 0 }  When the value noTest is written to this object, no action is
    // taken unless a test is in progress, in which case the test is aborted. 
    // Writing any other value to this object is only valid when no test is
    // currently in progress, in which case the indicated test is initiated.  When
    // read, this object always returns the most recent value that ifTestType was
    // set to.  If it has not been set since the last initialization of the
    // network management subsystem on the agent, a value of noTest is returned.
    // The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Iftesttype interface{}

    // This object contains the result of the most recently requested test, or the
    // value none(1) if no tests have been requested since the last reset.  Note
    // that this facility provides no provision for saving the results of one test
    // when starting another, as could be required if used by multiple managers
    // concurrently. The type is Iftestresult.
    Iftestresult interface{}

    // This object contains a code which contains more specific information on the
    // test result, for example an error-code after a failed test.  Error codes
    // and other values this object may take are specific to the type of interface
    // and/or test.  The value may have the semantics of either the AutonomousType
    // or InstancePointer textual conventions as defined in RFC 2579.  The
    // identifier:      testCodeUnknown  OBJECT IDENTIFIER ::= { 0 0 }  is defined
    // for use if no additional result code is available. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Iftestcode interface{}

    // The entity which currently has the 'ownership' required to invoke a test on
    // this interface. The type is string.
    Iftestowner interface{}
}

func (ifentry *IFMIB_Iftable_Ifentry) GetFilter() yfilter.YFilter { return ifentry.YFilter }

func (ifentry *IFMIB_Iftable_Ifentry) SetFilter(yf yfilter.YFilter) { ifentry.YFilter = yf }

func (ifentry *IFMIB_Iftable_Ifentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "ifDescr" { return "Ifdescr" }
    if yname == "ifType" { return "Iftype" }
    if yname == "ifMtu" { return "Ifmtu" }
    if yname == "ifSpeed" { return "Ifspeed" }
    if yname == "ifPhysAddress" { return "Ifphysaddress" }
    if yname == "ifAdminStatus" { return "Ifadminstatus" }
    if yname == "ifOperStatus" { return "Ifoperstatus" }
    if yname == "ifLastChange" { return "Iflastchange" }
    if yname == "ifInOctets" { return "Ifinoctets" }
    if yname == "ifInUcastPkts" { return "Ifinucastpkts" }
    if yname == "ifInNUcastPkts" { return "Ifinnucastpkts" }
    if yname == "ifInDiscards" { return "Ifindiscards" }
    if yname == "ifInErrors" { return "Ifinerrors" }
    if yname == "ifInUnknownProtos" { return "Ifinunknownprotos" }
    if yname == "ifOutOctets" { return "Ifoutoctets" }
    if yname == "ifOutUcastPkts" { return "Ifoutucastpkts" }
    if yname == "ifOutNUcastPkts" { return "Ifoutnucastpkts" }
    if yname == "ifOutDiscards" { return "Ifoutdiscards" }
    if yname == "ifOutErrors" { return "Ifouterrors" }
    if yname == "ifOutQLen" { return "Ifoutqlen" }
    if yname == "ifSpecific" { return "Ifspecific" }
    if yname == "ifName" { return "Ifname" }
    if yname == "ifInMulticastPkts" { return "Ifinmulticastpkts" }
    if yname == "ifInBroadcastPkts" { return "Ifinbroadcastpkts" }
    if yname == "ifOutMulticastPkts" { return "Ifoutmulticastpkts" }
    if yname == "ifOutBroadcastPkts" { return "Ifoutbroadcastpkts" }
    if yname == "ifHCInOctets" { return "Ifhcinoctets" }
    if yname == "ifHCInUcastPkts" { return "Ifhcinucastpkts" }
    if yname == "ifHCInMulticastPkts" { return "Ifhcinmulticastpkts" }
    if yname == "ifHCInBroadcastPkts" { return "Ifhcinbroadcastpkts" }
    if yname == "ifHCOutOctets" { return "Ifhcoutoctets" }
    if yname == "ifHCOutUcastPkts" { return "Ifhcoutucastpkts" }
    if yname == "ifHCOutMulticastPkts" { return "Ifhcoutmulticastpkts" }
    if yname == "ifHCOutBroadcastPkts" { return "Ifhcoutbroadcastpkts" }
    if yname == "ifLinkUpDownTrapEnable" { return "Iflinkupdowntrapenable" }
    if yname == "ifHighSpeed" { return "Ifhighspeed" }
    if yname == "ifPromiscuousMode" { return "Ifpromiscuousmode" }
    if yname == "ifConnectorPresent" { return "Ifconnectorpresent" }
    if yname == "ifAlias" { return "Ifalias" }
    if yname == "ifCounterDiscontinuityTime" { return "Ifcounterdiscontinuitytime" }
    if yname == "ifTestId" { return "Iftestid" }
    if yname == "ifTestStatus" { return "Ifteststatus" }
    if yname == "ifTestType" { return "Iftesttype" }
    if yname == "ifTestResult" { return "Iftestresult" }
    if yname == "ifTestCode" { return "Iftestcode" }
    if yname == "ifTestOwner" { return "Iftestowner" }
    return ""
}

func (ifentry *IFMIB_Iftable_Ifentry) GetSegmentPath() string {
    return "ifEntry" + "[ifIndex='" + fmt.Sprintf("%v", ifentry.Ifindex) + "']"
}

func (ifentry *IFMIB_Iftable_Ifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifentry *IFMIB_Iftable_Ifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifentry *IFMIB_Iftable_Ifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = ifentry.Ifindex
    leafs["ifDescr"] = ifentry.Ifdescr
    leafs["ifType"] = ifentry.Iftype
    leafs["ifMtu"] = ifentry.Ifmtu
    leafs["ifSpeed"] = ifentry.Ifspeed
    leafs["ifPhysAddress"] = ifentry.Ifphysaddress
    leafs["ifAdminStatus"] = ifentry.Ifadminstatus
    leafs["ifOperStatus"] = ifentry.Ifoperstatus
    leafs["ifLastChange"] = ifentry.Iflastchange
    leafs["ifInOctets"] = ifentry.Ifinoctets
    leafs["ifInUcastPkts"] = ifentry.Ifinucastpkts
    leafs["ifInNUcastPkts"] = ifentry.Ifinnucastpkts
    leafs["ifInDiscards"] = ifentry.Ifindiscards
    leafs["ifInErrors"] = ifentry.Ifinerrors
    leafs["ifInUnknownProtos"] = ifentry.Ifinunknownprotos
    leafs["ifOutOctets"] = ifentry.Ifoutoctets
    leafs["ifOutUcastPkts"] = ifentry.Ifoutucastpkts
    leafs["ifOutNUcastPkts"] = ifentry.Ifoutnucastpkts
    leafs["ifOutDiscards"] = ifentry.Ifoutdiscards
    leafs["ifOutErrors"] = ifentry.Ifouterrors
    leafs["ifOutQLen"] = ifentry.Ifoutqlen
    leafs["ifSpecific"] = ifentry.Ifspecific
    leafs["ifName"] = ifentry.Ifname
    leafs["ifInMulticastPkts"] = ifentry.Ifinmulticastpkts
    leafs["ifInBroadcastPkts"] = ifentry.Ifinbroadcastpkts
    leafs["ifOutMulticastPkts"] = ifentry.Ifoutmulticastpkts
    leafs["ifOutBroadcastPkts"] = ifentry.Ifoutbroadcastpkts
    leafs["ifHCInOctets"] = ifentry.Ifhcinoctets
    leafs["ifHCInUcastPkts"] = ifentry.Ifhcinucastpkts
    leafs["ifHCInMulticastPkts"] = ifentry.Ifhcinmulticastpkts
    leafs["ifHCInBroadcastPkts"] = ifentry.Ifhcinbroadcastpkts
    leafs["ifHCOutOctets"] = ifentry.Ifhcoutoctets
    leafs["ifHCOutUcastPkts"] = ifentry.Ifhcoutucastpkts
    leafs["ifHCOutMulticastPkts"] = ifentry.Ifhcoutmulticastpkts
    leafs["ifHCOutBroadcastPkts"] = ifentry.Ifhcoutbroadcastpkts
    leafs["ifLinkUpDownTrapEnable"] = ifentry.Iflinkupdowntrapenable
    leafs["ifHighSpeed"] = ifentry.Ifhighspeed
    leafs["ifPromiscuousMode"] = ifentry.Ifpromiscuousmode
    leafs["ifConnectorPresent"] = ifentry.Ifconnectorpresent
    leafs["ifAlias"] = ifentry.Ifalias
    leafs["ifCounterDiscontinuityTime"] = ifentry.Ifcounterdiscontinuitytime
    leafs["ifTestId"] = ifentry.Iftestid
    leafs["ifTestStatus"] = ifentry.Ifteststatus
    leafs["ifTestType"] = ifentry.Iftesttype
    leafs["ifTestResult"] = ifentry.Iftestresult
    leafs["ifTestCode"] = ifentry.Iftestcode
    leafs["ifTestOwner"] = ifentry.Iftestowner
    return leafs
}

func (ifentry *IFMIB_Iftable_Ifentry) GetBundleName() string { return "cisco_ios_xe" }

func (ifentry *IFMIB_Iftable_Ifentry) GetYangName() string { return "ifEntry" }

func (ifentry *IFMIB_Iftable_Ifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ifentry *IFMIB_Iftable_Ifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ifentry *IFMIB_Iftable_Ifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ifentry *IFMIB_Iftable_Ifentry) SetParent(parent types.Entity) { ifentry.parent = parent }

func (ifentry *IFMIB_Iftable_Ifentry) GetParent() types.Entity { return ifentry.parent }

func (ifentry *IFMIB_Iftable_Ifentry) GetParentYangName() string { return "ifTable" }

// IFMIB_Iftable_Ifentry_Ifadminstatus represents in the down(2) state).
type IFMIB_Iftable_Ifentry_Ifadminstatus string

const (
    IFMIB_Iftable_Ifentry_Ifadminstatus_up IFMIB_Iftable_Ifentry_Ifadminstatus = "up"

    IFMIB_Iftable_Ifentry_Ifadminstatus_down IFMIB_Iftable_Ifentry_Ifadminstatus = "down"

    IFMIB_Iftable_Ifentry_Ifadminstatus_testing IFMIB_Iftable_Ifentry_Ifadminstatus = "testing"
)

// IFMIB_Iftable_Ifentry_Iflinkupdowntrapenable represents otherwise.
type IFMIB_Iftable_Ifentry_Iflinkupdowntrapenable string

const (
    IFMIB_Iftable_Ifentry_Iflinkupdowntrapenable_enabled IFMIB_Iftable_Ifentry_Iflinkupdowntrapenable = "enabled"

    IFMIB_Iftable_Ifentry_Iflinkupdowntrapenable_disabled IFMIB_Iftable_Ifentry_Iflinkupdowntrapenable = "disabled"
)

// IFMIB_Iftable_Ifentry_Ifoperstatus represents components.
type IFMIB_Iftable_Ifentry_Ifoperstatus string

const (
    IFMIB_Iftable_Ifentry_Ifoperstatus_up IFMIB_Iftable_Ifentry_Ifoperstatus = "up"

    IFMIB_Iftable_Ifentry_Ifoperstatus_down IFMIB_Iftable_Ifentry_Ifoperstatus = "down"

    IFMIB_Iftable_Ifentry_Ifoperstatus_testing IFMIB_Iftable_Ifentry_Ifoperstatus = "testing"

    IFMIB_Iftable_Ifentry_Ifoperstatus_unknown IFMIB_Iftable_Ifentry_Ifoperstatus = "unknown"

    IFMIB_Iftable_Ifentry_Ifoperstatus_dormant IFMIB_Iftable_Ifentry_Ifoperstatus = "dormant"

    IFMIB_Iftable_Ifentry_Ifoperstatus_notPresent IFMIB_Iftable_Ifentry_Ifoperstatus = "notPresent"

    IFMIB_Iftable_Ifentry_Ifoperstatus_lowerLayerDown IFMIB_Iftable_Ifentry_Ifoperstatus = "lowerLayerDown"
)

// IFMIB_Iftable_Ifentry_Iftestresult represents multiple managers concurrently.
type IFMIB_Iftable_Ifentry_Iftestresult string

const (
    IFMIB_Iftable_Ifentry_Iftestresult_none IFMIB_Iftable_Ifentry_Iftestresult = "none"

    IFMIB_Iftable_Ifentry_Iftestresult_success IFMIB_Iftable_Ifentry_Iftestresult = "success"

    IFMIB_Iftable_Ifentry_Iftestresult_inProgress IFMIB_Iftable_Ifentry_Iftestresult = "inProgress"

    IFMIB_Iftable_Ifentry_Iftestresult_notSupported IFMIB_Iftable_Ifentry_Iftestresult = "notSupported"

    IFMIB_Iftable_Ifentry_Iftestresult_unAbleToRun IFMIB_Iftable_Ifentry_Iftestresult = "unAbleToRun"

    IFMIB_Iftable_Ifentry_Iftestresult_aborted IFMIB_Iftable_Ifentry_Iftestresult = "aborted"

    IFMIB_Iftable_Ifentry_Iftestresult_failed IFMIB_Iftable_Ifentry_Iftestresult = "failed"
)

// IFMIB_Iftable_Ifentry_Ifteststatus represents to 'notInUse(1)'.
type IFMIB_Iftable_Ifentry_Ifteststatus string

const (
    IFMIB_Iftable_Ifentry_Ifteststatus_notInUse IFMIB_Iftable_Ifentry_Ifteststatus = "notInUse"

    IFMIB_Iftable_Ifentry_Ifteststatus_inUse IFMIB_Iftable_Ifentry_Ifteststatus = "inUse"
)

// IFMIB_Ifstacktable
// The table containing information on the relationships
// between the multiple sub-layers of network interfaces.  In
// particular, it contains information on which sub-layers run
// 'on top of' which other sub-layers, where each sub-layer
// corresponds to a conceptual row in the ifTable.  For
// example, when the sub-layer with ifIndex value x runs over
// the sub-layer with ifIndex value y, then this table
// contains:
// 
//   ifStackStatus.x.y=active
// 
// For each ifIndex value, I, which identifies an active
// interface, there are always at least two instantiated rows
// in this table associated with I.  For one of these rows, I
// is the value of ifStackHigherLayer; for the other, I is the
// value of ifStackLowerLayer.  (If I is not involved in
// multiplexing, then these are the only two rows associated
// with I.)
// 
// For example, two rows exist even for an interface which has
// no others stacked on top or below it:
// 
//   ifStackStatus.0.x=active
//   ifStackStatus.x.0=active 
type IFMIB_Ifstacktable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information on a particular relationship between two sub- layers,
    // specifying that one sub-layer runs on 'top' of the other sub-layer.  Each
    // sub-layer corresponds to a conceptual row in the ifTable. The type is slice
    // of IFMIB_Ifstacktable_Ifstackentry.
    Ifstackentry []IFMIB_Ifstacktable_Ifstackentry
}

func (ifstacktable *IFMIB_Ifstacktable) GetFilter() yfilter.YFilter { return ifstacktable.YFilter }

func (ifstacktable *IFMIB_Ifstacktable) SetFilter(yf yfilter.YFilter) { ifstacktable.YFilter = yf }

func (ifstacktable *IFMIB_Ifstacktable) GetGoName(yname string) string {
    if yname == "ifStackEntry" { return "Ifstackentry" }
    return ""
}

func (ifstacktable *IFMIB_Ifstacktable) GetSegmentPath() string {
    return "ifStackTable"
}

func (ifstacktable *IFMIB_Ifstacktable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ifStackEntry" {
        for _, c := range ifstacktable.Ifstackentry {
            if ifstacktable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IFMIB_Ifstacktable_Ifstackentry{}
        ifstacktable.Ifstackentry = append(ifstacktable.Ifstackentry, child)
        return &ifstacktable.Ifstackentry[len(ifstacktable.Ifstackentry)-1]
    }
    return nil
}

func (ifstacktable *IFMIB_Ifstacktable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ifstacktable.Ifstackentry {
        children[ifstacktable.Ifstackentry[i].GetSegmentPath()] = &ifstacktable.Ifstackentry[i]
    }
    return children
}

func (ifstacktable *IFMIB_Ifstacktable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ifstacktable *IFMIB_Ifstacktable) GetBundleName() string { return "cisco_ios_xe" }

func (ifstacktable *IFMIB_Ifstacktable) GetYangName() string { return "ifStackTable" }

func (ifstacktable *IFMIB_Ifstacktable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ifstacktable *IFMIB_Ifstacktable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ifstacktable *IFMIB_Ifstacktable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ifstacktable *IFMIB_Ifstacktable) SetParent(parent types.Entity) { ifstacktable.parent = parent }

func (ifstacktable *IFMIB_Ifstacktable) GetParent() types.Entity { return ifstacktable.parent }

func (ifstacktable *IFMIB_Ifstacktable) GetParentYangName() string { return "IF-MIB" }

// IFMIB_Ifstacktable_Ifstackentry
// Information on a particular relationship between two sub-
// layers, specifying that one sub-layer runs on 'top' of the
// other sub-layer.  Each sub-layer corresponds to a conceptual
// row in the ifTable.
type IFMIB_Ifstacktable_Ifstackentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The value of ifIndex corresponding to the higher
    // sub-layer of the relationship, i.e., the sub-layer which runs on 'top' of
    // the sub-layer identified by the corresponding instance of
    // ifStackLowerLayer.  If there is no higher sub-layer (below the internetwork
    // layer), then this object has the value 0. The type is interface{} with
    // range: 0..2147483647.
    Ifstackhigherlayer interface{}

    // This attribute is a key. The value of ifIndex corresponding to the lower
    // sub-layer of the relationship, i.e., the sub-layer which runs 'below' the
    // sub-layer identified by the corresponding instance of ifStackHigherLayer. 
    // If there is no lower sub-layer, then this object has the value 0. The type
    // is interface{} with range: 0..2147483647.
    Ifstacklowerlayer interface{}

    // The status of the relationship between two sub-layers.  Changing the value
    // of this object from 'active' to 'notInService' or 'destroy' will likely
    // have consequences up and down the interface stack.  Thus, write access to
    // this object is likely to be inappropriate for some types of interfaces, and
    // many implementations will choose not to support write-access for any type
    // of interface. The type is RowStatus.
    Ifstackstatus interface{}
}

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetFilter() yfilter.YFilter { return ifstackentry.YFilter }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) SetFilter(yf yfilter.YFilter) { ifstackentry.YFilter = yf }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetGoName(yname string) string {
    if yname == "ifStackHigherLayer" { return "Ifstackhigherlayer" }
    if yname == "ifStackLowerLayer" { return "Ifstacklowerlayer" }
    if yname == "ifStackStatus" { return "Ifstackstatus" }
    return ""
}

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetSegmentPath() string {
    return "ifStackEntry" + "[ifStackHigherLayer='" + fmt.Sprintf("%v", ifstackentry.Ifstackhigherlayer) + "']" + "[ifStackLowerLayer='" + fmt.Sprintf("%v", ifstackentry.Ifstacklowerlayer) + "']"
}

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifStackHigherLayer"] = ifstackentry.Ifstackhigherlayer
    leafs["ifStackLowerLayer"] = ifstackentry.Ifstacklowerlayer
    leafs["ifStackStatus"] = ifstackentry.Ifstackstatus
    return leafs
}

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetBundleName() string { return "cisco_ios_xe" }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetYangName() string { return "ifStackEntry" }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) SetParent(parent types.Entity) { ifstackentry.parent = parent }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetParent() types.Entity { return ifstackentry.parent }

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetParentYangName() string { return "ifStackTable" }

// IFMIB_Ifrcvaddresstable
// This table contains an entry for each address (broadcast,
// multicast, or uni-cast) for which the system will receive
// packets/frames on a particular interface, except as follows:
// 
// - for an interface operating in promiscuous mode, entries
// are only required for those addresses for which the system
// would receive frames were it not operating in promiscuous
// mode.
// 
// - for 802.5 functional addresses, only one entry is
// required, for the address which has the functional address
// bit ANDed with the bit mask of all functional addresses for
// which the interface will accept frames.
// 
// A system is normally able to use any unicast address which
// corresponds to an entry in this table as a source address.
type IFMIB_Ifrcvaddresstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects identifying an address for which the system will accept
    // packets/frames on the particular interface identified by the index value
    // ifIndex. The type is slice of IFMIB_Ifrcvaddresstable_Ifrcvaddressentry.
    Ifrcvaddressentry []IFMIB_Ifrcvaddresstable_Ifrcvaddressentry
}

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetFilter() yfilter.YFilter { return ifrcvaddresstable.YFilter }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) SetFilter(yf yfilter.YFilter) { ifrcvaddresstable.YFilter = yf }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetGoName(yname string) string {
    if yname == "ifRcvAddressEntry" { return "Ifrcvaddressentry" }
    return ""
}

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetSegmentPath() string {
    return "ifRcvAddressTable"
}

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ifRcvAddressEntry" {
        for _, c := range ifrcvaddresstable.Ifrcvaddressentry {
            if ifrcvaddresstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IFMIB_Ifrcvaddresstable_Ifrcvaddressentry{}
        ifrcvaddresstable.Ifrcvaddressentry = append(ifrcvaddresstable.Ifrcvaddressentry, child)
        return &ifrcvaddresstable.Ifrcvaddressentry[len(ifrcvaddresstable.Ifrcvaddressentry)-1]
    }
    return nil
}

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ifrcvaddresstable.Ifrcvaddressentry {
        children[ifrcvaddresstable.Ifrcvaddressentry[i].GetSegmentPath()] = &ifrcvaddresstable.Ifrcvaddressentry[i]
    }
    return children
}

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetBundleName() string { return "cisco_ios_xe" }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetYangName() string { return "ifRcvAddressTable" }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) SetParent(parent types.Entity) { ifrcvaddresstable.parent = parent }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetParent() types.Entity { return ifrcvaddresstable.parent }

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetParentYangName() string { return "IF-MIB" }

// IFMIB_Ifrcvaddresstable_Ifrcvaddressentry
// A list of objects identifying an address for which the
// system will accept packets/frames on the particular
// interface identified by the index value ifIndex.
type IFMIB_Ifrcvaddresstable_Ifrcvaddressentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. An address for which the system will accept
    // packets/frames on this entry's interface. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    Ifrcvaddressaddress interface{}

    // This object is used to create and delete rows in the ifRcvAddressTable. The
    // type is RowStatus.
    Ifrcvaddressstatus interface{}

    // This object has the value nonVolatile(3) for those entries in the table
    // which are valid and will not be deleted by the next restart of the managed
    // system.  Entries having the value volatile(2) are valid and exist, but have
    // not been saved, so that will not exist after the next restart of the
    // managed system.  Entries having the value other(1) are valid and exist but
    // are not classified as to whether they will continue to exist after the next
    // restart. The type is Ifrcvaddresstype.
    Ifrcvaddresstype interface{}
}

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetFilter() yfilter.YFilter { return ifrcvaddressentry.YFilter }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) SetFilter(yf yfilter.YFilter) { ifrcvaddressentry.YFilter = yf }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "ifRcvAddressAddress" { return "Ifrcvaddressaddress" }
    if yname == "ifRcvAddressStatus" { return "Ifrcvaddressstatus" }
    if yname == "ifRcvAddressType" { return "Ifrcvaddresstype" }
    return ""
}

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetSegmentPath() string {
    return "ifRcvAddressEntry" + "[ifIndex='" + fmt.Sprintf("%v", ifrcvaddressentry.Ifindex) + "']" + "[ifRcvAddressAddress='" + fmt.Sprintf("%v", ifrcvaddressentry.Ifrcvaddressaddress) + "']"
}

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = ifrcvaddressentry.Ifindex
    leafs["ifRcvAddressAddress"] = ifrcvaddressentry.Ifrcvaddressaddress
    leafs["ifRcvAddressStatus"] = ifrcvaddressentry.Ifrcvaddressstatus
    leafs["ifRcvAddressType"] = ifrcvaddressentry.Ifrcvaddresstype
    return leafs
}

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetBundleName() string { return "cisco_ios_xe" }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetYangName() string { return "ifRcvAddressEntry" }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) SetParent(parent types.Entity) { ifrcvaddressentry.parent = parent }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetParent() types.Entity { return ifrcvaddressentry.parent }

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetParentYangName() string { return "ifRcvAddressTable" }

// IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype represents continue to exist after the next restart.
type IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype string

const (
    IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype_other IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype = "other"

    IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype_volatile IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype = "volatile"

    IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype_nonVolatile IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype = "nonVolatile"
)

