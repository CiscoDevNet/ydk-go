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

func (iFMIB *IFMIB) GetEntityData() *types.CommonEntityData {
    iFMIB.EntityData.YFilter = iFMIB.YFilter
    iFMIB.EntityData.YangName = "IF-MIB"
    iFMIB.EntityData.BundleName = "cisco_ios_xe"
    iFMIB.EntityData.ParentYangName = "IF-MIB"
    iFMIB.EntityData.SegmentPath = "IF-MIB:IF-MIB"
    iFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iFMIB.EntityData.Children = make(map[string]types.YChild)
    iFMIB.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &iFMIB.Interfaces}
    iFMIB.EntityData.Children["ifMIBObjects"] = types.YChild{"Ifmibobjects", &iFMIB.Ifmibobjects}
    iFMIB.EntityData.Children["ifTable"] = types.YChild{"Iftable", &iFMIB.Iftable}
    iFMIB.EntityData.Children["ifStackTable"] = types.YChild{"Ifstacktable", &iFMIB.Ifstacktable}
    iFMIB.EntityData.Children["ifRcvAddressTable"] = types.YChild{"Ifrcvaddresstable", &iFMIB.Ifrcvaddresstable}
    iFMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iFMIB.EntityData)
}

// IFMIB_Interfaces
type IFMIB_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of network interfaces (regardless of their current state)
    // present on this system. The type is interface{} with range:
    // -2147483648..2147483647.
    Ifnumber interface{}
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

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaces.EntityData.Leafs["ifNumber"] = types.YLeaf{"Ifnumber", interfaces.Ifnumber}
    return &(interfaces.EntityData)
}

// IFMIB_Ifmibobjects
type IFMIB_Ifmibobjects struct {
    EntityData types.CommonEntityData
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

func (ifmibobjects *IFMIB_Ifmibobjects) GetEntityData() *types.CommonEntityData {
    ifmibobjects.EntityData.YFilter = ifmibobjects.YFilter
    ifmibobjects.EntityData.YangName = "ifMIBObjects"
    ifmibobjects.EntityData.BundleName = "cisco_ios_xe"
    ifmibobjects.EntityData.ParentYangName = "IF-MIB"
    ifmibobjects.EntityData.SegmentPath = "ifMIBObjects"
    ifmibobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifmibobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifmibobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifmibobjects.EntityData.Children = make(map[string]types.YChild)
    ifmibobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    ifmibobjects.EntityData.Leafs["ifTableLastChange"] = types.YLeaf{"Iftablelastchange", ifmibobjects.Iftablelastchange}
    ifmibobjects.EntityData.Leafs["ifStackLastChange"] = types.YLeaf{"Ifstacklastchange", ifmibobjects.Ifstacklastchange}
    return &(ifmibobjects.EntityData)
}

// IFMIB_Iftable
// A list of interface entries.  The number of entries is
// given by the value of ifNumber.
type IFMIB_Iftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing management information applicable to a particular
    // interface. The type is slice of IFMIB_Iftable_Ifentry.
    Ifentry []IFMIB_Iftable_Ifentry
}

func (iftable *IFMIB_Iftable) GetEntityData() *types.CommonEntityData {
    iftable.EntityData.YFilter = iftable.YFilter
    iftable.EntityData.YangName = "ifTable"
    iftable.EntityData.BundleName = "cisco_ios_xe"
    iftable.EntityData.ParentYangName = "IF-MIB"
    iftable.EntityData.SegmentPath = "ifTable"
    iftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iftable.EntityData.Children = make(map[string]types.YChild)
    iftable.EntityData.Children["ifEntry"] = types.YChild{"Ifentry", nil}
    for i := range iftable.Ifentry {
        iftable.EntityData.Children[types.GetSegmentPath(&iftable.Ifentry[i])] = types.YChild{"Ifentry", &iftable.Ifentry[i]}
    }
    iftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iftable.EntityData)
}

// IFMIB_Iftable_Ifentry
// An entry containing management information applicable to a
// particular interface.
type IFMIB_Iftable_Ifentry struct {
    EntityData types.CommonEntityData
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
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Iftestcode interface{}

    // The entity which currently has the 'ownership' required to invoke a test on
    // this interface. The type is string.
    Iftestowner interface{}
}

func (ifentry *IFMIB_Iftable_Ifentry) GetEntityData() *types.CommonEntityData {
    ifentry.EntityData.YFilter = ifentry.YFilter
    ifentry.EntityData.YangName = "ifEntry"
    ifentry.EntityData.BundleName = "cisco_ios_xe"
    ifentry.EntityData.ParentYangName = "ifTable"
    ifentry.EntityData.SegmentPath = "ifEntry" + "[ifIndex='" + fmt.Sprintf("%v", ifentry.Ifindex) + "']"
    ifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifentry.EntityData.Children = make(map[string]types.YChild)
    ifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ifentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", ifentry.Ifindex}
    ifentry.EntityData.Leafs["ifDescr"] = types.YLeaf{"Ifdescr", ifentry.Ifdescr}
    ifentry.EntityData.Leafs["ifType"] = types.YLeaf{"Iftype", ifentry.Iftype}
    ifentry.EntityData.Leafs["ifMtu"] = types.YLeaf{"Ifmtu", ifentry.Ifmtu}
    ifentry.EntityData.Leafs["ifSpeed"] = types.YLeaf{"Ifspeed", ifentry.Ifspeed}
    ifentry.EntityData.Leafs["ifPhysAddress"] = types.YLeaf{"Ifphysaddress", ifentry.Ifphysaddress}
    ifentry.EntityData.Leafs["ifAdminStatus"] = types.YLeaf{"Ifadminstatus", ifentry.Ifadminstatus}
    ifentry.EntityData.Leafs["ifOperStatus"] = types.YLeaf{"Ifoperstatus", ifentry.Ifoperstatus}
    ifentry.EntityData.Leafs["ifLastChange"] = types.YLeaf{"Iflastchange", ifentry.Iflastchange}
    ifentry.EntityData.Leafs["ifInOctets"] = types.YLeaf{"Ifinoctets", ifentry.Ifinoctets}
    ifentry.EntityData.Leafs["ifInUcastPkts"] = types.YLeaf{"Ifinucastpkts", ifentry.Ifinucastpkts}
    ifentry.EntityData.Leafs["ifInNUcastPkts"] = types.YLeaf{"Ifinnucastpkts", ifentry.Ifinnucastpkts}
    ifentry.EntityData.Leafs["ifInDiscards"] = types.YLeaf{"Ifindiscards", ifentry.Ifindiscards}
    ifentry.EntityData.Leafs["ifInErrors"] = types.YLeaf{"Ifinerrors", ifentry.Ifinerrors}
    ifentry.EntityData.Leafs["ifInUnknownProtos"] = types.YLeaf{"Ifinunknownprotos", ifentry.Ifinunknownprotos}
    ifentry.EntityData.Leafs["ifOutOctets"] = types.YLeaf{"Ifoutoctets", ifentry.Ifoutoctets}
    ifentry.EntityData.Leafs["ifOutUcastPkts"] = types.YLeaf{"Ifoutucastpkts", ifentry.Ifoutucastpkts}
    ifentry.EntityData.Leafs["ifOutNUcastPkts"] = types.YLeaf{"Ifoutnucastpkts", ifentry.Ifoutnucastpkts}
    ifentry.EntityData.Leafs["ifOutDiscards"] = types.YLeaf{"Ifoutdiscards", ifentry.Ifoutdiscards}
    ifentry.EntityData.Leafs["ifOutErrors"] = types.YLeaf{"Ifouterrors", ifentry.Ifouterrors}
    ifentry.EntityData.Leafs["ifOutQLen"] = types.YLeaf{"Ifoutqlen", ifentry.Ifoutqlen}
    ifentry.EntityData.Leafs["ifSpecific"] = types.YLeaf{"Ifspecific", ifentry.Ifspecific}
    ifentry.EntityData.Leafs["ifName"] = types.YLeaf{"Ifname", ifentry.Ifname}
    ifentry.EntityData.Leafs["ifInMulticastPkts"] = types.YLeaf{"Ifinmulticastpkts", ifentry.Ifinmulticastpkts}
    ifentry.EntityData.Leafs["ifInBroadcastPkts"] = types.YLeaf{"Ifinbroadcastpkts", ifentry.Ifinbroadcastpkts}
    ifentry.EntityData.Leafs["ifOutMulticastPkts"] = types.YLeaf{"Ifoutmulticastpkts", ifentry.Ifoutmulticastpkts}
    ifentry.EntityData.Leafs["ifOutBroadcastPkts"] = types.YLeaf{"Ifoutbroadcastpkts", ifentry.Ifoutbroadcastpkts}
    ifentry.EntityData.Leafs["ifHCInOctets"] = types.YLeaf{"Ifhcinoctets", ifentry.Ifhcinoctets}
    ifentry.EntityData.Leafs["ifHCInUcastPkts"] = types.YLeaf{"Ifhcinucastpkts", ifentry.Ifhcinucastpkts}
    ifentry.EntityData.Leafs["ifHCInMulticastPkts"] = types.YLeaf{"Ifhcinmulticastpkts", ifentry.Ifhcinmulticastpkts}
    ifentry.EntityData.Leafs["ifHCInBroadcastPkts"] = types.YLeaf{"Ifhcinbroadcastpkts", ifentry.Ifhcinbroadcastpkts}
    ifentry.EntityData.Leafs["ifHCOutOctets"] = types.YLeaf{"Ifhcoutoctets", ifentry.Ifhcoutoctets}
    ifentry.EntityData.Leafs["ifHCOutUcastPkts"] = types.YLeaf{"Ifhcoutucastpkts", ifentry.Ifhcoutucastpkts}
    ifentry.EntityData.Leafs["ifHCOutMulticastPkts"] = types.YLeaf{"Ifhcoutmulticastpkts", ifentry.Ifhcoutmulticastpkts}
    ifentry.EntityData.Leafs["ifHCOutBroadcastPkts"] = types.YLeaf{"Ifhcoutbroadcastpkts", ifentry.Ifhcoutbroadcastpkts}
    ifentry.EntityData.Leafs["ifLinkUpDownTrapEnable"] = types.YLeaf{"Iflinkupdowntrapenable", ifentry.Iflinkupdowntrapenable}
    ifentry.EntityData.Leafs["ifHighSpeed"] = types.YLeaf{"Ifhighspeed", ifentry.Ifhighspeed}
    ifentry.EntityData.Leafs["ifPromiscuousMode"] = types.YLeaf{"Ifpromiscuousmode", ifentry.Ifpromiscuousmode}
    ifentry.EntityData.Leafs["ifConnectorPresent"] = types.YLeaf{"Ifconnectorpresent", ifentry.Ifconnectorpresent}
    ifentry.EntityData.Leafs["ifAlias"] = types.YLeaf{"Ifalias", ifentry.Ifalias}
    ifentry.EntityData.Leafs["ifCounterDiscontinuityTime"] = types.YLeaf{"Ifcounterdiscontinuitytime", ifentry.Ifcounterdiscontinuitytime}
    ifentry.EntityData.Leafs["ifTestId"] = types.YLeaf{"Iftestid", ifentry.Iftestid}
    ifentry.EntityData.Leafs["ifTestStatus"] = types.YLeaf{"Ifteststatus", ifentry.Ifteststatus}
    ifentry.EntityData.Leafs["ifTestType"] = types.YLeaf{"Iftesttype", ifentry.Iftesttype}
    ifentry.EntityData.Leafs["ifTestResult"] = types.YLeaf{"Iftestresult", ifentry.Iftestresult}
    ifentry.EntityData.Leafs["ifTestCode"] = types.YLeaf{"Iftestcode", ifentry.Iftestcode}
    ifentry.EntityData.Leafs["ifTestOwner"] = types.YLeaf{"Iftestowner", ifentry.Iftestowner}
    return &(ifentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on a particular relationship between two sub- layers,
    // specifying that one sub-layer runs on 'top' of the other sub-layer.  Each
    // sub-layer corresponds to a conceptual row in the ifTable. The type is slice
    // of IFMIB_Ifstacktable_Ifstackentry.
    Ifstackentry []IFMIB_Ifstacktable_Ifstackentry
}

func (ifstacktable *IFMIB_Ifstacktable) GetEntityData() *types.CommonEntityData {
    ifstacktable.EntityData.YFilter = ifstacktable.YFilter
    ifstacktable.EntityData.YangName = "ifStackTable"
    ifstacktable.EntityData.BundleName = "cisco_ios_xe"
    ifstacktable.EntityData.ParentYangName = "IF-MIB"
    ifstacktable.EntityData.SegmentPath = "ifStackTable"
    ifstacktable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifstacktable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifstacktable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifstacktable.EntityData.Children = make(map[string]types.YChild)
    ifstacktable.EntityData.Children["ifStackEntry"] = types.YChild{"Ifstackentry", nil}
    for i := range ifstacktable.Ifstackentry {
        ifstacktable.EntityData.Children[types.GetSegmentPath(&ifstacktable.Ifstackentry[i])] = types.YChild{"Ifstackentry", &ifstacktable.Ifstackentry[i]}
    }
    ifstacktable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ifstacktable.EntityData)
}

// IFMIB_Ifstacktable_Ifstackentry
// Information on a particular relationship between two sub-
// layers, specifying that one sub-layer runs on 'top' of the
// other sub-layer.  Each sub-layer corresponds to a conceptual
// row in the ifTable.
type IFMIB_Ifstacktable_Ifstackentry struct {
    EntityData types.CommonEntityData
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

func (ifstackentry *IFMIB_Ifstacktable_Ifstackentry) GetEntityData() *types.CommonEntityData {
    ifstackentry.EntityData.YFilter = ifstackentry.YFilter
    ifstackentry.EntityData.YangName = "ifStackEntry"
    ifstackentry.EntityData.BundleName = "cisco_ios_xe"
    ifstackentry.EntityData.ParentYangName = "ifStackTable"
    ifstackentry.EntityData.SegmentPath = "ifStackEntry" + "[ifStackHigherLayer='" + fmt.Sprintf("%v", ifstackentry.Ifstackhigherlayer) + "']" + "[ifStackLowerLayer='" + fmt.Sprintf("%v", ifstackentry.Ifstacklowerlayer) + "']"
    ifstackentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifstackentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifstackentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifstackentry.EntityData.Children = make(map[string]types.YChild)
    ifstackentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ifstackentry.EntityData.Leafs["ifStackHigherLayer"] = types.YLeaf{"Ifstackhigherlayer", ifstackentry.Ifstackhigherlayer}
    ifstackentry.EntityData.Leafs["ifStackLowerLayer"] = types.YLeaf{"Ifstacklowerlayer", ifstackentry.Ifstacklowerlayer}
    ifstackentry.EntityData.Leafs["ifStackStatus"] = types.YLeaf{"Ifstackstatus", ifstackentry.Ifstackstatus}
    return &(ifstackentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects identifying an address for which the system will accept
    // packets/frames on the particular interface identified by the index value
    // ifIndex. The type is slice of IFMIB_Ifrcvaddresstable_Ifrcvaddressentry.
    Ifrcvaddressentry []IFMIB_Ifrcvaddresstable_Ifrcvaddressentry
}

func (ifrcvaddresstable *IFMIB_Ifrcvaddresstable) GetEntityData() *types.CommonEntityData {
    ifrcvaddresstable.EntityData.YFilter = ifrcvaddresstable.YFilter
    ifrcvaddresstable.EntityData.YangName = "ifRcvAddressTable"
    ifrcvaddresstable.EntityData.BundleName = "cisco_ios_xe"
    ifrcvaddresstable.EntityData.ParentYangName = "IF-MIB"
    ifrcvaddresstable.EntityData.SegmentPath = "ifRcvAddressTable"
    ifrcvaddresstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifrcvaddresstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifrcvaddresstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifrcvaddresstable.EntityData.Children = make(map[string]types.YChild)
    ifrcvaddresstable.EntityData.Children["ifRcvAddressEntry"] = types.YChild{"Ifrcvaddressentry", nil}
    for i := range ifrcvaddresstable.Ifrcvaddressentry {
        ifrcvaddresstable.EntityData.Children[types.GetSegmentPath(&ifrcvaddresstable.Ifrcvaddressentry[i])] = types.YChild{"Ifrcvaddressentry", &ifrcvaddresstable.Ifrcvaddressentry[i]}
    }
    ifrcvaddresstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ifrcvaddresstable.EntityData)
}

// IFMIB_Ifrcvaddresstable_Ifrcvaddressentry
// A list of objects identifying an address for which the
// system will accept packets/frames on the particular
// interface identified by the index value ifIndex.
type IFMIB_Ifrcvaddresstable_Ifrcvaddressentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. An address for which the system will accept
    // packets/frames on this entry's interface. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
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

func (ifrcvaddressentry *IFMIB_Ifrcvaddresstable_Ifrcvaddressentry) GetEntityData() *types.CommonEntityData {
    ifrcvaddressentry.EntityData.YFilter = ifrcvaddressentry.YFilter
    ifrcvaddressentry.EntityData.YangName = "ifRcvAddressEntry"
    ifrcvaddressentry.EntityData.BundleName = "cisco_ios_xe"
    ifrcvaddressentry.EntityData.ParentYangName = "ifRcvAddressTable"
    ifrcvaddressentry.EntityData.SegmentPath = "ifRcvAddressEntry" + "[ifIndex='" + fmt.Sprintf("%v", ifrcvaddressentry.Ifindex) + "']" + "[ifRcvAddressAddress='" + fmt.Sprintf("%v", ifrcvaddressentry.Ifrcvaddressaddress) + "']"
    ifrcvaddressentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ifrcvaddressentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ifrcvaddressentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ifrcvaddressentry.EntityData.Children = make(map[string]types.YChild)
    ifrcvaddressentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ifrcvaddressentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", ifrcvaddressentry.Ifindex}
    ifrcvaddressentry.EntityData.Leafs["ifRcvAddressAddress"] = types.YLeaf{"Ifrcvaddressaddress", ifrcvaddressentry.Ifrcvaddressaddress}
    ifrcvaddressentry.EntityData.Leafs["ifRcvAddressStatus"] = types.YLeaf{"Ifrcvaddressstatus", ifrcvaddressentry.Ifrcvaddressstatus}
    ifrcvaddressentry.EntityData.Leafs["ifRcvAddressType"] = types.YLeaf{"Ifrcvaddresstype", ifrcvaddressentry.Ifrcvaddresstype}
    return &(ifrcvaddressentry.EntityData)
}

// IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype represents continue to exist after the next restart.
type IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype string

const (
    IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype_other IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype = "other"

    IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype_volatile IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype = "volatile"

    IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype_nonVolatile IFMIB_Ifrcvaddresstable_Ifrcvaddressentry_Ifrcvaddresstype = "nonVolatile"
)

