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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Interfaces IFMIB_Interfaces

    
    IfMIBObjects IFMIB_IfMIBObjects

    // A list of interface entries.  The number of entries is given by the value
    // of ifNumber.
    IfTable IFMIB_IfTable

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
    IfStackTable IFMIB_IfStackTable

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
    IfRcvAddressTable IFMIB_IfRcvAddressTable
}

func (iFMIB *IFMIB) GetEntityData() *types.CommonEntityData {
    iFMIB.EntityData.YFilter = iFMIB.YFilter
    iFMIB.EntityData.YangName = "IF-MIB"
    iFMIB.EntityData.BundleName = "cisco_ios_xe"
    iFMIB.EntityData.ParentYangName = "IF-MIB"
    iFMIB.EntityData.SegmentPath = "IF-MIB:IF-MIB"
    iFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iFMIB.EntityData.Children = types.NewOrderedMap()
    iFMIB.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &iFMIB.Interfaces})
    iFMIB.EntityData.Children.Append("ifMIBObjects", types.YChild{"IfMIBObjects", &iFMIB.IfMIBObjects})
    iFMIB.EntityData.Children.Append("ifTable", types.YChild{"IfTable", &iFMIB.IfTable})
    iFMIB.EntityData.Children.Append("ifStackTable", types.YChild{"IfStackTable", &iFMIB.IfStackTable})
    iFMIB.EntityData.Children.Append("ifRcvAddressTable", types.YChild{"IfRcvAddressTable", &iFMIB.IfRcvAddressTable})
    iFMIB.EntityData.Leafs = types.NewOrderedMap()

    iFMIB.EntityData.YListKeys = []string {}

    return &(iFMIB.EntityData)
}

// IFMIB_Interfaces
type IFMIB_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of network interfaces (regardless of their current state)
    // present on this system. The type is interface{} with range:
    // -2147483648..2147483647.
    IfNumber interface{}
}

func (interfaces *IFMIB_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xe"
    interfaces.EntityData.ParentYangName = "IF-MIB"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Leafs = types.NewOrderedMap()
    interfaces.EntityData.Leafs.Append("ifNumber", types.YLeaf{"IfNumber", interfaces.IfNumber})

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// IFMIB_IfMIBObjects
type IFMIB_IfMIBObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time of the last creation or deletion of an
    // entry in the ifTable.  If the number of entries has been unchanged since
    // the last re-initialization of the local network management subsystem, then
    // this object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    IfTableLastChange interface{}

    // The value of sysUpTime at the time of the last change of the (whole)
    // interface stack.  A change of the interface stack is defined to be any
    // creation, deletion, or change in value of any instance of ifStackStatus. 
    // If the interface stack has been unchanged since the last re-initialization
    // of the local network management subsystem, then this object contains a zero
    // value. The type is interface{} with range: 0..4294967295.
    IfStackLastChange interface{}
}

func (ifMIBObjects *IFMIB_IfMIBObjects) GetEntityData() *types.CommonEntityData {
    ifMIBObjects.EntityData.YFilter = ifMIBObjects.YFilter
    ifMIBObjects.EntityData.YangName = "ifMIBObjects"
    ifMIBObjects.EntityData.BundleName = "cisco_ios_xe"
    ifMIBObjects.EntityData.ParentYangName = "IF-MIB"
    ifMIBObjects.EntityData.SegmentPath = "ifMIBObjects"
    ifMIBObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifMIBObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifMIBObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifMIBObjects.EntityData.Children = types.NewOrderedMap()
    ifMIBObjects.EntityData.Leafs = types.NewOrderedMap()
    ifMIBObjects.EntityData.Leafs.Append("ifTableLastChange", types.YLeaf{"IfTableLastChange", ifMIBObjects.IfTableLastChange})
    ifMIBObjects.EntityData.Leafs.Append("ifStackLastChange", types.YLeaf{"IfStackLastChange", ifMIBObjects.IfStackLastChange})

    ifMIBObjects.EntityData.YListKeys = []string {}

    return &(ifMIBObjects.EntityData)
}

// IFMIB_IfTable
// A list of interface entries.  The number of entries is
// given by the value of ifNumber.
type IFMIB_IfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing management information applicable to a particular
    // interface. The type is slice of IFMIB_IfTable_IfEntry.
    IfEntry []*IFMIB_IfTable_IfEntry
}

func (ifTable *IFMIB_IfTable) GetEntityData() *types.CommonEntityData {
    ifTable.EntityData.YFilter = ifTable.YFilter
    ifTable.EntityData.YangName = "ifTable"
    ifTable.EntityData.BundleName = "cisco_ios_xe"
    ifTable.EntityData.ParentYangName = "IF-MIB"
    ifTable.EntityData.SegmentPath = "ifTable"
    ifTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifTable.EntityData.Children = types.NewOrderedMap()
    ifTable.EntityData.Children.Append("ifEntry", types.YChild{"IfEntry", nil})
    for i := range ifTable.IfEntry {
        ifTable.EntityData.Children.Append(types.GetSegmentPath(ifTable.IfEntry[i]), types.YChild{"IfEntry", ifTable.IfEntry[i]})
    }
    ifTable.EntityData.Leafs = types.NewOrderedMap()

    ifTable.EntityData.YListKeys = []string {}

    return &(ifTable.EntityData)
}

// IFMIB_IfTable_IfEntry
// An entry containing management information applicable to a
// particular interface.
type IFMIB_IfTable_IfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique value, greater than zero, for each
    // interface.  It is recommended that values are assigned contiguously
    // starting from 1.  The value for each interface sub-layer must remain
    // constant at least from one re-initialization of the entity's network
    // management system to the next re- initialization. The type is interface{}
    // with range: 1..2147483647.
    IfIndex interface{}

    // A textual string containing information about the interface.  This string
    // should include the name of the manufacturer, the product name and the
    // version of the interface hardware/software. The type is string with length:
    // 0..255.
    IfDescr interface{}

    // The type of interface.  Additional values for ifType are assigned by the
    // Internet Assigned Numbers Authority (IANA), through updating the syntax of
    // the IANAifType textual convention. The type is IANAifType.
    IfType interface{}

    // The size of the largest packet which can be sent/received on the interface,
    // specified in octets.  For interfaces that are used for transmitting network
    // datagrams, this is the size of the largest network datagram that can be
    // sent on the interface. The type is interface{} with range:
    // -2147483648..2147483647.
    IfMtu interface{}

    // An estimate of the interface's current bandwidth in bits per second.  For
    // interfaces which do not vary in bandwidth or for those where no accurate
    // estimation can be made, this object should contain the nominal bandwidth. 
    // If the bandwidth of the interface is greater than the maximum value
    // reportable by this object then this object should report its maximum value
    // (4,294,967,295) and ifHighSpeed must be used to report the interace's
    // speed.  For a sub-layer which has no concept of bandwidth, this object
    // should be zero. The type is interface{} with range: 0..4294967295.
    IfSpeed interface{}

    // The interface's address at its protocol sub-layer.  For example, for an
    // 802.x interface, this object normally contains a MAC address.  The
    // interface's media-specific MIB must define the bit and byte ordering and
    // the format of the value of this object.  For interfaces which do not have
    // such an address (e.g., a serial line), this object should contain an octet
    // string of zero length. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    IfPhysAddress interface{}

    // The desired state of the interface.  The testing(3) state indicates that no
    // operational packets can be passed.  When a managed system initializes, all
    // interfaces start with ifAdminStatus in the down(2) state.  As a result of
    // either explicit management action or per configuration information retained
    // by the managed system, ifAdminStatus is then changed to either the up(1) or
    // testing(3) states (or remains in the down(2) state). The type is
    // IfAdminStatus.
    IfAdminStatus interface{}

    // The current operational state of the interface.  The testing(3) state
    // indicates that no operational packets can be passed.  If ifAdminStatus is
    // down(2) then ifOperStatus should be down(2).  If ifAdminStatus is changed
    // to up(1) then ifOperStatus should change to up(1) if the interface is ready
    // to transmit and receive network traffic; it should change to dormant(5) if
    // the interface is waiting for external actions (such as a serial line
    // waiting for an incoming connection); it should remain in the down(2) state
    // if and only if there is a fault that prevents it from going to the up(1)
    // state; it should remain in the notPresent(6) state if the interface has
    // missing (typically, hardware) components. The type is IfOperStatus.
    IfOperStatus interface{}

    // The value of sysUpTime at the time the interface entered its current
    // operational state.  If the current state was entered prior to the last
    // re-initialization of the local network management subsystem, then this
    // object contains a zero value. The type is interface{} with range:
    // 0..4294967295.
    IfLastChange interface{}

    // The total number of octets received on the interface, including framing
    // characters.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    IfInOctets interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were not addressed to a multicast or broadcast address at this
    // sub-layer.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    IfInUcastPkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a multicast or broadcast address at this sub-layer.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime.  This object is deprecated in favour of
    // ifInMulticastPkts and ifInBroadcastPkts. The type is interface{} with
    // range: 0..4294967295.
    IfInNUcastPkts interface{}

    // The number of inbound packets which were chosen to be discarded even though
    // no errors had been detected to prevent their being deliverable to a
    // higher-layer protocol.  One possible reason for discarding such a packet
    // could be to free up buffer space.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    IfInDiscards interface{}

    // For packet-oriented interfaces, the number of inbound packets that
    // contained errors preventing them from being deliverable to a higher-layer
    // protocol.  For character- oriented or fixed-length interfaces, the number
    // of inbound transmission units that contained errors preventing them from
    // being deliverable to a higher-layer protocol.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    IfInErrors interface{}

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
    IfInUnknownProtos interface{}

    // The total number of octets transmitted out of the interface, including
    // framing characters.  Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    IfOutOctets interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were not addressed to a multicast or broadcast
    // address at this sub-layer, including those that were discarded or not sent.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IfOutUcastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a multicast or broadcast address
    // at this sub-layer, including those that were discarded or not sent. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime.  This object is deprecated in favour of
    // ifOutMulticastPkts and ifOutBroadcastPkts. The type is interface{} with
    // range: 0..4294967295.
    IfOutNUcastPkts interface{}

    // The number of outbound packets which were chosen to be discarded even
    // though no errors had been detected to prevent their being transmitted.  One
    // possible reason for discarding such a packet could be to free up buffer
    // space.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    IfOutDiscards interface{}

    // For packet-oriented interfaces, the number of outbound packets that could
    // not be transmitted because of errors. For character-oriented or
    // fixed-length interfaces, the number of outbound transmission units that
    // could not be transmitted because of errors.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of ifCounterDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295.
    IfOutErrors interface{}

    // The length of the output packet queue (in packets). The type is interface{}
    // with range: 0..4294967295.
    IfOutQLen interface{}

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
    IfSpecific interface{}

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
    IfName interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a multicast address at this sub-layer.  For a MAC
    // layer protocol, this includes both Group and Functional addresses. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IfInMulticastPkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a broadcast address at this sub-layer. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IfInBroadcastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a multicast address at this
    // sub-layer, including those that were discarded or not sent.  For a MAC
    // layer protocol, this includes both Group and Functional addresses. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IfOutMulticastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a broadcast address at this
    // sub-layer, including those that were discarded or not sent. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    IfOutBroadcastPkts interface{}

    // The total number of octets received on the interface, including framing
    // characters.  This object is a 64-bit version of ifInOctets. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IfHCInOctets interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were not addressed to a multicast or broadcast address at this
    // sub-layer.  This object is a 64-bit version of ifInUcastPkts. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IfHCInUcastPkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a multicast address at this sub-layer.  For a MAC
    // layer protocol, this includes both Group and Functional addresses.  This
    // object is a 64-bit version of ifInMulticastPkts.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IfHCInMulticastPkts interface{}

    // The number of packets, delivered by this sub-layer to a higher (sub-)layer,
    // which were addressed to a broadcast address at this sub-layer.  This object
    // is a 64-bit version of ifInBroadcastPkts.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..18446744073709551615.
    IfHCInBroadcastPkts interface{}

    // The total number of octets transmitted out of the interface, including
    // framing characters.  This object is a 64-bit version of ifOutOctets. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IfHCOutOctets interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were not addressed to a multicast or broadcast
    // address at this sub-layer, including those that were discarded or not sent.
    // This object is a 64-bit version of ifOutUcastPkts.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IfHCOutUcastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a multicast address at this
    // sub-layer, including those that were discarded or not sent.  For a MAC
    // layer protocol, this includes both Group and Functional addresses.  This
    // object is a 64-bit version of ifOutMulticastPkts.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..18446744073709551615.
    IfHCOutMulticastPkts interface{}

    // The total number of packets that higher-level protocols requested be
    // transmitted, and which were addressed to a broadcast address at this
    // sub-layer, including those that were discarded or not sent.  This object is
    // a 64-bit version of ifOutBroadcastPkts.  Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..18446744073709551615.
    IfHCOutBroadcastPkts interface{}

    // Indicates whether linkUp/linkDown traps should be generated for this
    // interface.  By default, this object should have the value enabled(1) for
    // interfaces which do not operate on 'top' of any other interface (as defined
    // in the ifStackTable), and disabled(2) otherwise. The type is
    // IfLinkUpDownTrapEnable.
    IfLinkUpDownTrapEnable interface{}

    // An estimate of the interface's current bandwidth in units of 1,000,000 bits
    // per second.  If this object reports a value of `n' then the speed of the
    // interface is somewhere in the range of `n-500,000' to `n+499,999'.  For
    // interfaces which do not vary in bandwidth or for those where no accurate
    // estimation can be made, this object should contain the nominal bandwidth. 
    // For a sub-layer which has no concept of bandwidth, this object should be
    // zero. The type is interface{} with range: 0..4294967295.
    IfHighSpeed interface{}

    // This object has a value of false(2) if this interface only accepts
    // packets/frames that are addressed to this station. This object has a value
    // of true(1) when the station accepts all packets/frames transmitted on the
    // media.  The value true(1) is only legal on certain types of media.  If
    // legal, setting this object to a value of true(1) may require the interface
    // to be reset before becoming effective.  The value of ifPromiscuousMode does
    // not affect the reception of broadcast and multicast packets/frames by the
    // interface. The type is bool.
    IfPromiscuousMode interface{}

    // This object has the value 'true(1)' if the interface sublayer has a
    // physical connector and the value 'false(2)' otherwise. The type is bool.
    IfConnectorPresent interface{}

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
    IfAlias interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this interface's counters suffered a discontinuity.  The relevant
    // counters are the specific instances associated with this interface of any
    // Counter32 or Counter64 object contained in the ifTable or ifXTable.  If no
    // such discontinuities have occurred since the last re- initialization of the
    // local management subsystem, then this object contains a zero value. The
    // type is interface{} with range: 0..4294967295.
    IfCounterDiscontinuityTime interface{}

    // This object identifies the current invocation of the interface's test. The
    // type is interface{} with range: 0..2147483647.
    IfTestId interface{}

    // This object indicates whether or not some manager currently has the
    // necessary 'ownership' required to invoke a test on this interface.  A write
    // to this object is only successful when it changes its value from
    // 'notInUse(1)' to 'inUse(2)'. After completion of a test, the agent resets
    // the value back to 'notInUse(1)'. The type is IfTestStatus.
    IfTestStatus interface{}

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
    IfTestType interface{}

    // This object contains the result of the most recently requested test, or the
    // value none(1) if no tests have been requested since the last reset.  Note
    // that this facility provides no provision for saving the results of one test
    // when starting another, as could be required if used by multiple managers
    // concurrently. The type is IfTestResult.
    IfTestResult interface{}

    // This object contains a code which contains more specific information on the
    // test result, for example an error-code after a failed test.  Error codes
    // and other values this object may take are specific to the type of interface
    // and/or test.  The value may have the semantics of either the AutonomousType
    // or InstancePointer textual conventions as defined in RFC 2579.  The
    // identifier:      testCodeUnknown  OBJECT IDENTIFIER ::= { 0 0 }  is defined
    // for use if no additional result code is available. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    IfTestCode interface{}

    // The entity which currently has the 'ownership' required to invoke a test on
    // this interface. The type is string.
    IfTestOwner interface{}
}

func (ifEntry *IFMIB_IfTable_IfEntry) GetEntityData() *types.CommonEntityData {
    ifEntry.EntityData.YFilter = ifEntry.YFilter
    ifEntry.EntityData.YangName = "ifEntry"
    ifEntry.EntityData.BundleName = "cisco_ios_xe"
    ifEntry.EntityData.ParentYangName = "ifTable"
    ifEntry.EntityData.SegmentPath = "ifEntry" + types.AddKeyToken(ifEntry.IfIndex, "ifIndex")
    ifEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifEntry.EntityData.Children = types.NewOrderedMap()
    ifEntry.EntityData.Leafs = types.NewOrderedMap()
    ifEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", ifEntry.IfIndex})
    ifEntry.EntityData.Leafs.Append("ifDescr", types.YLeaf{"IfDescr", ifEntry.IfDescr})
    ifEntry.EntityData.Leafs.Append("ifType", types.YLeaf{"IfType", ifEntry.IfType})
    ifEntry.EntityData.Leafs.Append("ifMtu", types.YLeaf{"IfMtu", ifEntry.IfMtu})
    ifEntry.EntityData.Leafs.Append("ifSpeed", types.YLeaf{"IfSpeed", ifEntry.IfSpeed})
    ifEntry.EntityData.Leafs.Append("ifPhysAddress", types.YLeaf{"IfPhysAddress", ifEntry.IfPhysAddress})
    ifEntry.EntityData.Leafs.Append("ifAdminStatus", types.YLeaf{"IfAdminStatus", ifEntry.IfAdminStatus})
    ifEntry.EntityData.Leafs.Append("ifOperStatus", types.YLeaf{"IfOperStatus", ifEntry.IfOperStatus})
    ifEntry.EntityData.Leafs.Append("ifLastChange", types.YLeaf{"IfLastChange", ifEntry.IfLastChange})
    ifEntry.EntityData.Leafs.Append("ifInOctets", types.YLeaf{"IfInOctets", ifEntry.IfInOctets})
    ifEntry.EntityData.Leafs.Append("ifInUcastPkts", types.YLeaf{"IfInUcastPkts", ifEntry.IfInUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifInNUcastPkts", types.YLeaf{"IfInNUcastPkts", ifEntry.IfInNUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifInDiscards", types.YLeaf{"IfInDiscards", ifEntry.IfInDiscards})
    ifEntry.EntityData.Leafs.Append("ifInErrors", types.YLeaf{"IfInErrors", ifEntry.IfInErrors})
    ifEntry.EntityData.Leafs.Append("ifInUnknownProtos", types.YLeaf{"IfInUnknownProtos", ifEntry.IfInUnknownProtos})
    ifEntry.EntityData.Leafs.Append("ifOutOctets", types.YLeaf{"IfOutOctets", ifEntry.IfOutOctets})
    ifEntry.EntityData.Leafs.Append("ifOutUcastPkts", types.YLeaf{"IfOutUcastPkts", ifEntry.IfOutUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifOutNUcastPkts", types.YLeaf{"IfOutNUcastPkts", ifEntry.IfOutNUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifOutDiscards", types.YLeaf{"IfOutDiscards", ifEntry.IfOutDiscards})
    ifEntry.EntityData.Leafs.Append("ifOutErrors", types.YLeaf{"IfOutErrors", ifEntry.IfOutErrors})
    ifEntry.EntityData.Leafs.Append("ifOutQLen", types.YLeaf{"IfOutQLen", ifEntry.IfOutQLen})
    ifEntry.EntityData.Leafs.Append("ifSpecific", types.YLeaf{"IfSpecific", ifEntry.IfSpecific})
    ifEntry.EntityData.Leafs.Append("ifName", types.YLeaf{"IfName", ifEntry.IfName})
    ifEntry.EntityData.Leafs.Append("ifInMulticastPkts", types.YLeaf{"IfInMulticastPkts", ifEntry.IfInMulticastPkts})
    ifEntry.EntityData.Leafs.Append("ifInBroadcastPkts", types.YLeaf{"IfInBroadcastPkts", ifEntry.IfInBroadcastPkts})
    ifEntry.EntityData.Leafs.Append("ifOutMulticastPkts", types.YLeaf{"IfOutMulticastPkts", ifEntry.IfOutMulticastPkts})
    ifEntry.EntityData.Leafs.Append("ifOutBroadcastPkts", types.YLeaf{"IfOutBroadcastPkts", ifEntry.IfOutBroadcastPkts})
    ifEntry.EntityData.Leafs.Append("ifHCInOctets", types.YLeaf{"IfHCInOctets", ifEntry.IfHCInOctets})
    ifEntry.EntityData.Leafs.Append("ifHCInUcastPkts", types.YLeaf{"IfHCInUcastPkts", ifEntry.IfHCInUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifHCInMulticastPkts", types.YLeaf{"IfHCInMulticastPkts", ifEntry.IfHCInMulticastPkts})
    ifEntry.EntityData.Leafs.Append("ifHCInBroadcastPkts", types.YLeaf{"IfHCInBroadcastPkts", ifEntry.IfHCInBroadcastPkts})
    ifEntry.EntityData.Leafs.Append("ifHCOutOctets", types.YLeaf{"IfHCOutOctets", ifEntry.IfHCOutOctets})
    ifEntry.EntityData.Leafs.Append("ifHCOutUcastPkts", types.YLeaf{"IfHCOutUcastPkts", ifEntry.IfHCOutUcastPkts})
    ifEntry.EntityData.Leafs.Append("ifHCOutMulticastPkts", types.YLeaf{"IfHCOutMulticastPkts", ifEntry.IfHCOutMulticastPkts})
    ifEntry.EntityData.Leafs.Append("ifHCOutBroadcastPkts", types.YLeaf{"IfHCOutBroadcastPkts", ifEntry.IfHCOutBroadcastPkts})
    ifEntry.EntityData.Leafs.Append("ifLinkUpDownTrapEnable", types.YLeaf{"IfLinkUpDownTrapEnable", ifEntry.IfLinkUpDownTrapEnable})
    ifEntry.EntityData.Leafs.Append("ifHighSpeed", types.YLeaf{"IfHighSpeed", ifEntry.IfHighSpeed})
    ifEntry.EntityData.Leafs.Append("ifPromiscuousMode", types.YLeaf{"IfPromiscuousMode", ifEntry.IfPromiscuousMode})
    ifEntry.EntityData.Leafs.Append("ifConnectorPresent", types.YLeaf{"IfConnectorPresent", ifEntry.IfConnectorPresent})
    ifEntry.EntityData.Leafs.Append("ifAlias", types.YLeaf{"IfAlias", ifEntry.IfAlias})
    ifEntry.EntityData.Leafs.Append("ifCounterDiscontinuityTime", types.YLeaf{"IfCounterDiscontinuityTime", ifEntry.IfCounterDiscontinuityTime})
    ifEntry.EntityData.Leafs.Append("ifTestId", types.YLeaf{"IfTestId", ifEntry.IfTestId})
    ifEntry.EntityData.Leafs.Append("ifTestStatus", types.YLeaf{"IfTestStatus", ifEntry.IfTestStatus})
    ifEntry.EntityData.Leafs.Append("ifTestType", types.YLeaf{"IfTestType", ifEntry.IfTestType})
    ifEntry.EntityData.Leafs.Append("ifTestResult", types.YLeaf{"IfTestResult", ifEntry.IfTestResult})
    ifEntry.EntityData.Leafs.Append("ifTestCode", types.YLeaf{"IfTestCode", ifEntry.IfTestCode})
    ifEntry.EntityData.Leafs.Append("ifTestOwner", types.YLeaf{"IfTestOwner", ifEntry.IfTestOwner})

    ifEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(ifEntry.EntityData)
}

// IFMIB_IfTable_IfEntry_IfAdminStatus represents in the down(2) state).
type IFMIB_IfTable_IfEntry_IfAdminStatus string

const (
    IFMIB_IfTable_IfEntry_IfAdminStatus_up IFMIB_IfTable_IfEntry_IfAdminStatus = "up"

    IFMIB_IfTable_IfEntry_IfAdminStatus_down IFMIB_IfTable_IfEntry_IfAdminStatus = "down"

    IFMIB_IfTable_IfEntry_IfAdminStatus_testing IFMIB_IfTable_IfEntry_IfAdminStatus = "testing"
)

// IFMIB_IfTable_IfEntry_IfLinkUpDownTrapEnable represents otherwise.
type IFMIB_IfTable_IfEntry_IfLinkUpDownTrapEnable string

const (
    IFMIB_IfTable_IfEntry_IfLinkUpDownTrapEnable_enabled IFMIB_IfTable_IfEntry_IfLinkUpDownTrapEnable = "enabled"

    IFMIB_IfTable_IfEntry_IfLinkUpDownTrapEnable_disabled IFMIB_IfTable_IfEntry_IfLinkUpDownTrapEnable = "disabled"
)

// IFMIB_IfTable_IfEntry_IfOperStatus represents components.
type IFMIB_IfTable_IfEntry_IfOperStatus string

const (
    IFMIB_IfTable_IfEntry_IfOperStatus_up IFMIB_IfTable_IfEntry_IfOperStatus = "up"

    IFMIB_IfTable_IfEntry_IfOperStatus_down IFMIB_IfTable_IfEntry_IfOperStatus = "down"

    IFMIB_IfTable_IfEntry_IfOperStatus_testing IFMIB_IfTable_IfEntry_IfOperStatus = "testing"

    IFMIB_IfTable_IfEntry_IfOperStatus_unknown IFMIB_IfTable_IfEntry_IfOperStatus = "unknown"

    IFMIB_IfTable_IfEntry_IfOperStatus_dormant IFMIB_IfTable_IfEntry_IfOperStatus = "dormant"

    IFMIB_IfTable_IfEntry_IfOperStatus_notPresent IFMIB_IfTable_IfEntry_IfOperStatus = "notPresent"

    IFMIB_IfTable_IfEntry_IfOperStatus_lowerLayerDown IFMIB_IfTable_IfEntry_IfOperStatus = "lowerLayerDown"
)

// IFMIB_IfTable_IfEntry_IfTestResult represents multiple managers concurrently.
type IFMIB_IfTable_IfEntry_IfTestResult string

const (
    IFMIB_IfTable_IfEntry_IfTestResult_none IFMIB_IfTable_IfEntry_IfTestResult = "none"

    IFMIB_IfTable_IfEntry_IfTestResult_success IFMIB_IfTable_IfEntry_IfTestResult = "success"

    IFMIB_IfTable_IfEntry_IfTestResult_inProgress IFMIB_IfTable_IfEntry_IfTestResult = "inProgress"

    IFMIB_IfTable_IfEntry_IfTestResult_notSupported IFMIB_IfTable_IfEntry_IfTestResult = "notSupported"

    IFMIB_IfTable_IfEntry_IfTestResult_unAbleToRun IFMIB_IfTable_IfEntry_IfTestResult = "unAbleToRun"

    IFMIB_IfTable_IfEntry_IfTestResult_aborted IFMIB_IfTable_IfEntry_IfTestResult = "aborted"

    IFMIB_IfTable_IfEntry_IfTestResult_failed IFMIB_IfTable_IfEntry_IfTestResult = "failed"
)

// IFMIB_IfTable_IfEntry_IfTestStatus represents to 'notInUse(1)'.
type IFMIB_IfTable_IfEntry_IfTestStatus string

const (
    IFMIB_IfTable_IfEntry_IfTestStatus_notInUse IFMIB_IfTable_IfEntry_IfTestStatus = "notInUse"

    IFMIB_IfTable_IfEntry_IfTestStatus_inUse IFMIB_IfTable_IfEntry_IfTestStatus = "inUse"
)

// IFMIB_IfStackTable
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
type IFMIB_IfStackTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on a particular relationship between two sub- layers,
    // specifying that one sub-layer runs on 'top' of the other sub-layer.  Each
    // sub-layer corresponds to a conceptual row in the ifTable. The type is slice
    // of IFMIB_IfStackTable_IfStackEntry.
    IfStackEntry []*IFMIB_IfStackTable_IfStackEntry
}

func (ifStackTable *IFMIB_IfStackTable) GetEntityData() *types.CommonEntityData {
    ifStackTable.EntityData.YFilter = ifStackTable.YFilter
    ifStackTable.EntityData.YangName = "ifStackTable"
    ifStackTable.EntityData.BundleName = "cisco_ios_xe"
    ifStackTable.EntityData.ParentYangName = "IF-MIB"
    ifStackTable.EntityData.SegmentPath = "ifStackTable"
    ifStackTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifStackTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifStackTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifStackTable.EntityData.Children = types.NewOrderedMap()
    ifStackTable.EntityData.Children.Append("ifStackEntry", types.YChild{"IfStackEntry", nil})
    for i := range ifStackTable.IfStackEntry {
        ifStackTable.EntityData.Children.Append(types.GetSegmentPath(ifStackTable.IfStackEntry[i]), types.YChild{"IfStackEntry", ifStackTable.IfStackEntry[i]})
    }
    ifStackTable.EntityData.Leafs = types.NewOrderedMap()

    ifStackTable.EntityData.YListKeys = []string {}

    return &(ifStackTable.EntityData)
}

// IFMIB_IfStackTable_IfStackEntry
// Information on a particular relationship between two sub-
// layers, specifying that one sub-layer runs on 'top' of the
// other sub-layer.  Each sub-layer corresponds to a conceptual
// row in the ifTable.
type IFMIB_IfStackTable_IfStackEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The value of ifIndex corresponding to the higher
    // sub-layer of the relationship, i.e., the sub-layer which runs on 'top' of
    // the sub-layer identified by the corresponding instance of
    // ifStackLowerLayer.  If there is no higher sub-layer (below the internetwork
    // layer), then this object has the value 0. The type is interface{} with
    // range: 0..2147483647.
    IfStackHigherLayer interface{}

    // This attribute is a key. The value of ifIndex corresponding to the lower
    // sub-layer of the relationship, i.e., the sub-layer which runs 'below' the
    // sub-layer identified by the corresponding instance of ifStackHigherLayer. 
    // If there is no lower sub-layer, then this object has the value 0. The type
    // is interface{} with range: 0..2147483647.
    IfStackLowerLayer interface{}

    // The status of the relationship between two sub-layers.  Changing the value
    // of this object from 'active' to 'notInService' or 'destroy' will likely
    // have consequences up and down the interface stack.  Thus, write access to
    // this object is likely to be inappropriate for some types of interfaces, and
    // many implementations will choose not to support write-access for any type
    // of interface. The type is RowStatus.
    IfStackStatus interface{}
}

func (ifStackEntry *IFMIB_IfStackTable_IfStackEntry) GetEntityData() *types.CommonEntityData {
    ifStackEntry.EntityData.YFilter = ifStackEntry.YFilter
    ifStackEntry.EntityData.YangName = "ifStackEntry"
    ifStackEntry.EntityData.BundleName = "cisco_ios_xe"
    ifStackEntry.EntityData.ParentYangName = "ifStackTable"
    ifStackEntry.EntityData.SegmentPath = "ifStackEntry" + types.AddKeyToken(ifStackEntry.IfStackHigherLayer, "ifStackHigherLayer") + types.AddKeyToken(ifStackEntry.IfStackLowerLayer, "ifStackLowerLayer")
    ifStackEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifStackEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifStackEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifStackEntry.EntityData.Children = types.NewOrderedMap()
    ifStackEntry.EntityData.Leafs = types.NewOrderedMap()
    ifStackEntry.EntityData.Leafs.Append("ifStackHigherLayer", types.YLeaf{"IfStackHigherLayer", ifStackEntry.IfStackHigherLayer})
    ifStackEntry.EntityData.Leafs.Append("ifStackLowerLayer", types.YLeaf{"IfStackLowerLayer", ifStackEntry.IfStackLowerLayer})
    ifStackEntry.EntityData.Leafs.Append("ifStackStatus", types.YLeaf{"IfStackStatus", ifStackEntry.IfStackStatus})

    ifStackEntry.EntityData.YListKeys = []string {"IfStackHigherLayer", "IfStackLowerLayer"}

    return &(ifStackEntry.EntityData)
}

// IFMIB_IfRcvAddressTable
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
type IFMIB_IfRcvAddressTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects identifying an address for which the system will accept
    // packets/frames on the particular interface identified by the index value
    // ifIndex. The type is slice of IFMIB_IfRcvAddressTable_IfRcvAddressEntry.
    IfRcvAddressEntry []*IFMIB_IfRcvAddressTable_IfRcvAddressEntry
}

func (ifRcvAddressTable *IFMIB_IfRcvAddressTable) GetEntityData() *types.CommonEntityData {
    ifRcvAddressTable.EntityData.YFilter = ifRcvAddressTable.YFilter
    ifRcvAddressTable.EntityData.YangName = "ifRcvAddressTable"
    ifRcvAddressTable.EntityData.BundleName = "cisco_ios_xe"
    ifRcvAddressTable.EntityData.ParentYangName = "IF-MIB"
    ifRcvAddressTable.EntityData.SegmentPath = "ifRcvAddressTable"
    ifRcvAddressTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifRcvAddressTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifRcvAddressTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifRcvAddressTable.EntityData.Children = types.NewOrderedMap()
    ifRcvAddressTable.EntityData.Children.Append("ifRcvAddressEntry", types.YChild{"IfRcvAddressEntry", nil})
    for i := range ifRcvAddressTable.IfRcvAddressEntry {
        ifRcvAddressTable.EntityData.Children.Append(types.GetSegmentPath(ifRcvAddressTable.IfRcvAddressEntry[i]), types.YChild{"IfRcvAddressEntry", ifRcvAddressTable.IfRcvAddressEntry[i]})
    }
    ifRcvAddressTable.EntityData.Leafs = types.NewOrderedMap()

    ifRcvAddressTable.EntityData.YListKeys = []string {}

    return &(ifRcvAddressTable.EntityData)
}

// IFMIB_IfRcvAddressTable_IfRcvAddressEntry
// A list of objects identifying an address for which the
// system will accept packets/frames on the particular
// interface identified by the index value ifIndex.
type IFMIB_IfRcvAddressTable_IfRcvAddressEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. An address for which the system will accept
    // packets/frames on this entry's interface. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    IfRcvAddressAddress interface{}

    // This object is used to create and delete rows in the ifRcvAddressTable. The
    // type is RowStatus.
    IfRcvAddressStatus interface{}

    // This object has the value nonVolatile(3) for those entries in the table
    // which are valid and will not be deleted by the next restart of the managed
    // system.  Entries having the value volatile(2) are valid and exist, but have
    // not been saved, so that will not exist after the next restart of the
    // managed system.  Entries having the value other(1) are valid and exist but
    // are not classified as to whether they will continue to exist after the next
    // restart. The type is IfRcvAddressType.
    IfRcvAddressType interface{}
}

func (ifRcvAddressEntry *IFMIB_IfRcvAddressTable_IfRcvAddressEntry) GetEntityData() *types.CommonEntityData {
    ifRcvAddressEntry.EntityData.YFilter = ifRcvAddressEntry.YFilter
    ifRcvAddressEntry.EntityData.YangName = "ifRcvAddressEntry"
    ifRcvAddressEntry.EntityData.BundleName = "cisco_ios_xe"
    ifRcvAddressEntry.EntityData.ParentYangName = "ifRcvAddressTable"
    ifRcvAddressEntry.EntityData.SegmentPath = "ifRcvAddressEntry" + types.AddKeyToken(ifRcvAddressEntry.IfIndex, "ifIndex") + types.AddKeyToken(ifRcvAddressEntry.IfRcvAddressAddress, "ifRcvAddressAddress")
    ifRcvAddressEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifRcvAddressEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifRcvAddressEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifRcvAddressEntry.EntityData.Children = types.NewOrderedMap()
    ifRcvAddressEntry.EntityData.Leafs = types.NewOrderedMap()
    ifRcvAddressEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", ifRcvAddressEntry.IfIndex})
    ifRcvAddressEntry.EntityData.Leafs.Append("ifRcvAddressAddress", types.YLeaf{"IfRcvAddressAddress", ifRcvAddressEntry.IfRcvAddressAddress})
    ifRcvAddressEntry.EntityData.Leafs.Append("ifRcvAddressStatus", types.YLeaf{"IfRcvAddressStatus", ifRcvAddressEntry.IfRcvAddressStatus})
    ifRcvAddressEntry.EntityData.Leafs.Append("ifRcvAddressType", types.YLeaf{"IfRcvAddressType", ifRcvAddressEntry.IfRcvAddressType})

    ifRcvAddressEntry.EntityData.YListKeys = []string {"IfIndex", "IfRcvAddressAddress"}

    return &(ifRcvAddressEntry.EntityData)
}

// IFMIB_IfRcvAddressTable_IfRcvAddressEntry_IfRcvAddressType represents continue to exist after the next restart.
type IFMIB_IfRcvAddressTable_IfRcvAddressEntry_IfRcvAddressType string

const (
    IFMIB_IfRcvAddressTable_IfRcvAddressEntry_IfRcvAddressType_other IFMIB_IfRcvAddressTable_IfRcvAddressEntry_IfRcvAddressType = "other"

    IFMIB_IfRcvAddressTable_IfRcvAddressEntry_IfRcvAddressType_volatile IFMIB_IfRcvAddressTable_IfRcvAddressEntry_IfRcvAddressType = "volatile"

    IFMIB_IfRcvAddressTable_IfRcvAddressEntry_IfRcvAddressType_nonVolatile IFMIB_IfRcvAddressTable_IfRcvAddressEntry_IfRcvAddressType = "nonVolatile"
)

