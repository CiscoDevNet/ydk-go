// Unicast Reverse Path Forwarding (URPF) is a function that
// checks the validity of the source address of IP packets
// received on an interface. This in an attempt to prevent
// Denial of Service attacks based on IP address spoofing.
// 
// URPF checks validity of a source address by determining
// whether the packet would be successfully routed as a
// destination address. 
// Based on configuration, the check made
// can be for existence of any route for the address, or more
// strictly for a route out the interface on which the packet
// was received by the device. When a violating packet is
// detected, it can be dropped. 
// This MIB allows detection of
// spoofingevents.
package cisco_ip_urpf_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ip_urpf_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IP-URPF-MIB CISCO-IP-URPF-MIB}", reflect.TypeOf(CISCOIPURPFMIB{}))
    ydk.RegisterEntity("CISCO-IP-URPF-MIB:CISCO-IP-URPF-MIB", reflect.TypeOf(CISCOIPURPFMIB{}))
}

// UnicastRpfType represents     the interface.
type UnicastRpfType string

const (
    UnicastRpfType_strict UnicastRpfType = "strict"

    UnicastRpfType_loose UnicastRpfType = "loose"

    UnicastRpfType_disabled UnicastRpfType = "disabled"
)

// CISCOIPURPFMIB
type CISCOIPURPFMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cipurpfscalar CISCOIPURPFMIB_Cipurpfscalar

    // This table contains summary information for the managed device on URPF
    // dropping.
    Cipurpftable CISCOIPURPFMIB_Cipurpftable

    // This table contains information on URPF dropping on an interface.
    Cipurpfifmontable CISCOIPURPFMIB_Cipurpfifmontable

    // This table contains statistics information for interfaces performing URPF
    // using VRF table to determine reachability.
    Cipurpfvrfiftable CISCOIPURPFMIB_Cipurpfvrfiftable

    // This table enables indexing URPF drop statistics by Virtual Routing and
    // Forwarding instances.
    Cipurpfvrftable CISCOIPURPFMIB_Cipurpfvrftable
}

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPURPFMIB.EntityData.YFilter = cISCOIPURPFMIB.YFilter
    cISCOIPURPFMIB.EntityData.YangName = "CISCO-IP-URPF-MIB"
    cISCOIPURPFMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPURPFMIB.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cISCOIPURPFMIB.EntityData.SegmentPath = "CISCO-IP-URPF-MIB:CISCO-IP-URPF-MIB"
    cISCOIPURPFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPURPFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPURPFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPURPFMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPURPFMIB.EntityData.Children["cipUrpfScalar"] = types.YChild{"Cipurpfscalar", &cISCOIPURPFMIB.Cipurpfscalar}
    cISCOIPURPFMIB.EntityData.Children["cipUrpfTable"] = types.YChild{"Cipurpftable", &cISCOIPURPFMIB.Cipurpftable}
    cISCOIPURPFMIB.EntityData.Children["cipUrpfIfMonTable"] = types.YChild{"Cipurpfifmontable", &cISCOIPURPFMIB.Cipurpfifmontable}
    cISCOIPURPFMIB.EntityData.Children["cipUrpfVrfIfTable"] = types.YChild{"Cipurpfvrfiftable", &cISCOIPURPFMIB.Cipurpfvrfiftable}
    cISCOIPURPFMIB.EntityData.Children["cipUrpfVrfTable"] = types.YChild{"Cipurpfvrftable", &cISCOIPURPFMIB.Cipurpfvrftable}
    cISCOIPURPFMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPURPFMIB.EntityData)
}

// CISCOIPURPFMIB_Cipurpfscalar
type CISCOIPURPFMIB_Cipurpfscalar struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The window of time in the recent past over which the drop count used in the
    // drop rate computation is collected.  This global value applies for the
    // computation of all URPF  rates, global and per-interface.   Once the period
    // over which computations have been  performed exceeds cipUrpfDropRateWindow,
    // every time a  computation is performed, the window slides up to end  at the
    // current time and start at cipUrpfDropRateWindow  seconds before.   The
    // cipUrpfDropRateWindow must be greater than or equal to the interval between
    // computations  (cipUrpfComputeInterval).  Since the agent must save the drop
    // count values for each compute interval in order to slide the window, the
    // number of counts saved is the quotient of cipUrpfDropRateWindow divided by
    // cipUrpfComputeInterval. The type is interface{} with range: 1..600. Units
    // are seconds.
    Cipurpfdropratewindow interface{}

    // The time between rate computations. This global value applies for the
    // computation of all URPF rates, global and per-interface.  When the value of
    // cipUrpfComputeInterval is changed, the interval in-progress proceeds as
    // though the value had not changed. The change will apply to the length of
    // subsequent intervals.  The cipUrpfComputeInterval must be less than or
    // equal  to the cipUrpfDropRateWindow. The type is interface{} with range:
    // 1..120. Units are seconds.
    Cipurpfcomputeinterval interface{}

    // The minimum time between issuance of cipUrpfIfDropRateNotify notifications
    // for a  particular interface and packet forwarding type.  Notifications are
    // generated for each interface and packet forwarding type that exceeds the
    // drop-rate.  When a Notify is sent because the drop-rate is  exceeded for a
    // particular interface and forwarding type, the time specified by this object
    // is used to  specify the minimum time that must elapse before  another
    // Notify can be sent for that interface and forwarding type. The time is
    // specified globally but  used individually. The type is interface{} with
    // range: 1..1000. Units are seconds.
    Cipurpfdropnotifyholddowntime interface{}
}

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetEntityData() *types.CommonEntityData {
    cipurpfscalar.EntityData.YFilter = cipurpfscalar.YFilter
    cipurpfscalar.EntityData.YangName = "cipUrpfScalar"
    cipurpfscalar.EntityData.BundleName = "cisco_ios_xe"
    cipurpfscalar.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipurpfscalar.EntityData.SegmentPath = "cipUrpfScalar"
    cipurpfscalar.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpfscalar.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpfscalar.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpfscalar.EntityData.Children = make(map[string]types.YChild)
    cipurpfscalar.EntityData.Leafs = make(map[string]types.YLeaf)
    cipurpfscalar.EntityData.Leafs["cipUrpfDropRateWindow"] = types.YLeaf{"Cipurpfdropratewindow", cipurpfscalar.Cipurpfdropratewindow}
    cipurpfscalar.EntityData.Leafs["cipUrpfComputeInterval"] = types.YLeaf{"Cipurpfcomputeinterval", cipurpfscalar.Cipurpfcomputeinterval}
    cipurpfscalar.EntityData.Leafs["cipUrpfDropNotifyHoldDownTime"] = types.YLeaf{"Cipurpfdropnotifyholddowntime", cipurpfscalar.Cipurpfdropnotifyholddowntime}
    return &(cipurpfscalar.EntityData)
}

// CISCOIPURPFMIB_Cipurpftable
// This table contains summary information for the
// managed device on URPF dropping.
type CISCOIPURPFMIB_Cipurpftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the managed device supports URPF dropping, a row exists for each IP
    // version type (v4 and v6). A row contains summary information on URPF
    // dropping over the entire managed device. The type is slice of
    // CISCOIPURPFMIB_Cipurpftable_Cipurpfentry.
    Cipurpfentry []CISCOIPURPFMIB_Cipurpftable_Cipurpfentry
}

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetEntityData() *types.CommonEntityData {
    cipurpftable.EntityData.YFilter = cipurpftable.YFilter
    cipurpftable.EntityData.YangName = "cipUrpfTable"
    cipurpftable.EntityData.BundleName = "cisco_ios_xe"
    cipurpftable.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipurpftable.EntityData.SegmentPath = "cipUrpfTable"
    cipurpftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpftable.EntityData.Children = make(map[string]types.YChild)
    cipurpftable.EntityData.Children["cipUrpfEntry"] = types.YChild{"Cipurpfentry", nil}
    for i := range cipurpftable.Cipurpfentry {
        cipurpftable.EntityData.Children[types.GetSegmentPath(&cipurpftable.Cipurpfentry[i])] = types.YChild{"Cipurpfentry", &cipurpftable.Cipurpfentry[i]}
    }
    cipurpftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipurpftable.EntityData)
}

// CISCOIPURPFMIB_Cipurpftable_Cipurpfentry
// If the managed device supports URPF dropping,
// a row exists for each IP version type (v4 and v6).
// A row contains summary information on URPF
// dropping over the entire managed device.
type CISCOIPURPFMIB_Cipurpftable_Cipurpfentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specifies the version of IP forwarding on an
    // interface to which the table row URPF counts, rates, and configuration
    // apply. The type is Cipurpfipversion.
    Cipurpfipversion interface{}

    // Sum of dropped IP version cipUrpfIpVersion packets failing a URPF check.
    // This value is the sum of drops of packets  received on all interfaces of
    // the managed device. The type is interface{} with range: 0..4294967295.
    // Units are packets.
    Cipurpfdrops interface{}

    // The rate of packet drops of IP version cipUrpfIpVersion packets due to URPF
    // for the managed device. The per-interface drop rate notification is issued
    // on rates exceeding a limit (rising rate). This dropping may indicate an
    // security attack on the network. To determine whether the attack/event is
    // over, the NMS must consult the managed device. This object can be polled to
    // determine the recent drop rate for the managed device as a whole, in
    // addition to querying particular interface objects.  This object is the
    // average rate of dropping over the most recent window of time. The rate is
    // computed by dividing the number of packets dropped over a window by the
    // window time in seconds. The window time is specified by
    // cipUrpfDropRateWindow. Each time the drop rate is computed, and at system
    // startup, a snapshot is taken of the latest value of cipUrpfDrops.
    // Subtracting from this the snapshot of cipUrpfDrops at the start of the
    // current window of time gives the number of packets dropped. The drop rate
    // is computed every cipUrpfComputeInterval seconds. As an example, let
    // cipUrpfDropRateWindow be 300 seconds, and cipUrpfComputeInterval 30
    // seconds. Every 30 seconds, the drop count five minutes previous is
    // subtracted from the current drop count, and the result is divided by 300 to
    // arrive at the drop rate.  At device start-up, until the device has been up
    // more than cipUrpfDropRateWindow, when drop rate is computed, the value of
    // cipUrpfDrops is divided by the time the device has been up.  After the
    // device has been up for cipUrpfDropRateWindow, when drop rate is computed,
    // the number of packet drops counted from interval start time to the
    // computation time is divided by cipUrpfDropRateWindow.  Changes to
    // cipUrpfDropRateWindow are not reflected in this object until the next
    // computation time.  The rate from the most recent computation is the value
    // fetched until the subsequent computation is performed. The type is
    // interface{} with range: 0..4294967295. Units are packets per second.
    Cipurpfdroprate interface{}
}

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetEntityData() *types.CommonEntityData {
    cipurpfentry.EntityData.YFilter = cipurpfentry.YFilter
    cipurpfentry.EntityData.YangName = "cipUrpfEntry"
    cipurpfentry.EntityData.BundleName = "cisco_ios_xe"
    cipurpfentry.EntityData.ParentYangName = "cipUrpfTable"
    cipurpfentry.EntityData.SegmentPath = "cipUrpfEntry" + "[cipUrpfIpVersion='" + fmt.Sprintf("%v", cipurpfentry.Cipurpfipversion) + "']"
    cipurpfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpfentry.EntityData.Children = make(map[string]types.YChild)
    cipurpfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipurpfentry.EntityData.Leafs["cipUrpfIpVersion"] = types.YLeaf{"Cipurpfipversion", cipurpfentry.Cipurpfipversion}
    cipurpfentry.EntityData.Leafs["cipUrpfDrops"] = types.YLeaf{"Cipurpfdrops", cipurpfentry.Cipurpfdrops}
    cipurpfentry.EntityData.Leafs["cipUrpfDropRate"] = types.YLeaf{"Cipurpfdroprate", cipurpfentry.Cipurpfdroprate}
    return &(cipurpfentry.EntityData)
}

// CISCOIPURPFMIB_Cipurpftable_Cipurpfentry_Cipurpfipversion represents configuration apply.
type CISCOIPURPFMIB_Cipurpftable_Cipurpfentry_Cipurpfipversion string

const (
    CISCOIPURPFMIB_Cipurpftable_Cipurpfentry_Cipurpfipversion_ipv4 CISCOIPURPFMIB_Cipurpftable_Cipurpfentry_Cipurpfipversion = "ipv4"

    CISCOIPURPFMIB_Cipurpftable_Cipurpfentry_Cipurpfipversion_ipv6 CISCOIPURPFMIB_Cipurpftable_Cipurpfentry_Cipurpfipversion = "ipv6"
)

// CISCOIPURPFMIB_Cipurpfifmontable
// This table contains information on URPF dropping on
// an interface.
type CISCOIPURPFMIB_Cipurpfifmontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If IPv4 packet forwarding is configured on an interface, and is configured
    // to perform URPF checking, a row appears in this table with indices
    // [ifIndex][ipv4]. If IPv4 packet forwarding is deconfigured, or URPF
    // checking is deconfigured, the row disappears.  If IPv6 packet forwarding is
    // configured on an interface, and is configured to perform URPF checking, a
    // row appears in the table with indices [ifIndex][ipv6].  If IPv6 packet
    // forwarding is deconfigured, or URPF checking is deconfigured, the row
    // disappears. The type is slice of
    // CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry.
    Cipurpfifmonentry []CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry
}

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetEntityData() *types.CommonEntityData {
    cipurpfifmontable.EntityData.YFilter = cipurpfifmontable.YFilter
    cipurpfifmontable.EntityData.YangName = "cipUrpfIfMonTable"
    cipurpfifmontable.EntityData.BundleName = "cisco_ios_xe"
    cipurpfifmontable.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipurpfifmontable.EntityData.SegmentPath = "cipUrpfIfMonTable"
    cipurpfifmontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpfifmontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpfifmontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpfifmontable.EntityData.Children = make(map[string]types.YChild)
    cipurpfifmontable.EntityData.Children["cipUrpfIfMonEntry"] = types.YChild{"Cipurpfifmonentry", nil}
    for i := range cipurpfifmontable.Cipurpfifmonentry {
        cipurpfifmontable.EntityData.Children[types.GetSegmentPath(&cipurpfifmontable.Cipurpfifmonentry[i])] = types.YChild{"Cipurpfifmonentry", &cipurpfifmontable.Cipurpfifmonentry[i]}
    }
    cipurpfifmontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipurpfifmontable.EntityData)
}

// CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry
// If IPv4 packet forwarding is configured on an interface,
// and is configured to perform URPF checking, a row appears
// in this table with indices [ifIndex][ipv4]. If IPv4
// packet forwarding is deconfigured, or URPF checking
// is deconfigured, the row disappears.
// 
// If IPv6 packet forwarding is configured on an interface,
// and is configured to perform URPF checking, a row appears
// in the table with indices [ifIndex][ipv6].  If IPv6
// packet forwarding is deconfigured, or URPF checking
// is deconfigured, the row disappears.
type CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. Specifies the version of IP forwarding on an
    // interface to which the table row URPF counts, rates, and  configuration
    // apply. The type is Cipurpfifipversion.
    Cipurpfifipversion interface{}

    // The number of IP packets of version cipUrpfIfIpVersion failing the URPF
    // check and dropped by the managed device on a particular interface. 
    // Discontinuities in the value of this variable can occur  at
    // re-initialization of the management system, and at  other times as
    // indicated by the values of  cipUrpfIfDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295. Units are packets.
    Cipurpfifdrops interface{}

    // The number of IP packets of version cipUrpfIfIpVersion failing the URPF
    // check but given a reprieve and not  dropped by the managed device.
    // Depending on the  device configuration and capabilities, the following 
    // cases may cause incrementing of the counter:  - if the managed device is
    // configured to allow self-pings    and the managed device pings itself. - if
    // the managed device is configured for loose URPF (if any   interface has a
    // route to the source), and the strict   case fails while the loose case
    // passes. - DHCP Request packets (src 0.0.0.0 dst 255.255.255.255)    will
    // pass after initially being marked for drop. - RIP routing on unnumbered
    // interfaces will pass after    initially being marked for drop. - multicast
    // packets will pass after initially being marked    for drop - ACL's can be
    // applied to permit packets after initially    being marked for drop. 
    // Discontinuities in the value of this variable can occur  at
    // re-initialization of the management system, and at  other times as
    // indicated by the values of  cipUrpfIfDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295. Units are packets.
    Cipurpfifsuppresseddrops interface{}

    // The rate of packet drops of IP version cipUrpfIfIpVersion packets due to
    // URPF on the interface.   This object is the average rate of dropping over
    // the most  recent interval of time. The rate is computed by dividing the
    // number of packets dropped over an interval by the  interval time in
    // seconds. Each time the drop rate is computed, and at system startup, a
    // snapshot is taken of the latest value of cipUrpfIfDrops. Subtracting from
    // this the snapshot of cipUrpfIfDrops at the start of the current interval of
    // time gives the number of packets dropped. The drop rate is computed every
    // cipUrpfComputeInterval seconds.  When drop rate is computed, if time since
    // the creation of  a row in cipUrpfIfMonTable is less than 
    // cipUrpfDropRateWindow, the value of cipUrpfIfDrops is  divided by the time
    // since row was created.  After the row has been in existence for 
    // cipUrpfDropRateWindow, when drop rate is computed, the  number of packet
    // drops counted on the interface from  interval start time to the computation
    // time is divided  by cipUrpfDropRateWindow.  Changes to
    // cipUrpfDropRateWindow are not reflected in this object until the next
    // computation time.  The rate from the  most recent computation is the value 
    // fetched until the subsequent computation is performed. The type is
    // interface{} with range: 0..4294967295. Units are packets/second.
    Cipurpfifdroprate interface{}

    // The value of sysUpTime on the most recent occasion at which this
    // interface's  counters suffered  a discontinuity. If no such discontinuities
    // have occurred since the last re-initialization of the local management
    // subsystem, then this object contains a value of zero. The type is
    // interface{} with range: 0..4294967295.
    Cipurpfifdiscontinuitytime interface{}

    // This object specifies whether the system produces the
    // cipUrpfIfDropRateNotify notification as a result of URPF  dropping of
    // version cipUrpfIfIpVersion IP packets on this  interface. A false value
    // prevents such notifications from  being generated by this system. The type
    // is bool.
    Cipurpfifdropratenotifyenable interface{}

    // When the calculated rate of URPF packet drops (cipUrpfIfDropRate) meets or
    // exceeds the value  specified by this object, a cipUrpfIfDropRateNotify 
    // notification is sent if cipUrpfIfDropRateNotifyEnable  is set to true, and
    // no such notification for the IP version has been sent for this interface
    // for the  hold-down period.  Note that due to the calculation used for drop
    // rate,  if there are less than n drop events in an n-second period the
    // notification will not be generated. To allow for the detection of a small
    // number of drop events, the value 0 (zero) is used to indicate that if any
    // drop events occur during the interval, a notification is generated. The
    // type is interface{} with range: 0..4294967295. Units are packets/second.
    Cipurpfifnotifydropratethreshold interface{}

    // Setting this object to true causes the five-minute hold-down timer for
    // emitting URPF drop rate  notifications for IP version cipUrpfIfIpVersion on
    // the interface to be short-circuited.  If a notification  is due and would
    // be emitted for the interface if the  five-minutes elapsed, setting this
    // object will cause  the notification to be sent.  This is a trigger, and
    // doesn't hold information. It is set and an action is performed. Therefore a
    // get for  this object always returns false. The type is bool.
    Cipurpfifnotifydrholddownreset interface{}

    // Interface configuration indicating the strictness of the reachability check
    // performed  on the interface. - strict: check that source addr is reachable
    // via            the interface it came in on. - loose : check that source
    // addr is reachable via            some interface on the device. The type is
    // Cipurpfifcheckstrict.
    Cipurpfifcheckstrict interface{}

    // Interface configuration indicating the routing table consulted for the
    // reachability check: - default: the non-private routing table for of the    
    // managed system. - vrf   : a particular VPN routing table. The type is
    // Cipurpfifwhichroutetableid.
    Cipurpfifwhichroutetableid interface{}

    // If the value of cipUrpfIfWhichRouteTableID is 'vrf', the name of the VRF
    // Table. Otherwise a zero-length string. The type is string with length:
    // 0..32.
    Cipurpfifvrfname interface{}
}

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetEntityData() *types.CommonEntityData {
    cipurpfifmonentry.EntityData.YFilter = cipurpfifmonentry.YFilter
    cipurpfifmonentry.EntityData.YangName = "cipUrpfIfMonEntry"
    cipurpfifmonentry.EntityData.BundleName = "cisco_ios_xe"
    cipurpfifmonentry.EntityData.ParentYangName = "cipUrpfIfMonTable"
    cipurpfifmonentry.EntityData.SegmentPath = "cipUrpfIfMonEntry" + "[ifIndex='" + fmt.Sprintf("%v", cipurpfifmonentry.Ifindex) + "']" + "[cipUrpfIfIpVersion='" + fmt.Sprintf("%v", cipurpfifmonentry.Cipurpfifipversion) + "']"
    cipurpfifmonentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpfifmonentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpfifmonentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpfifmonentry.EntityData.Children = make(map[string]types.YChild)
    cipurpfifmonentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipurpfifmonentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cipurpfifmonentry.Ifindex}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfIpVersion"] = types.YLeaf{"Cipurpfifipversion", cipurpfifmonentry.Cipurpfifipversion}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfDrops"] = types.YLeaf{"Cipurpfifdrops", cipurpfifmonentry.Cipurpfifdrops}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfSuppressedDrops"] = types.YLeaf{"Cipurpfifsuppresseddrops", cipurpfifmonentry.Cipurpfifsuppresseddrops}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfDropRate"] = types.YLeaf{"Cipurpfifdroprate", cipurpfifmonentry.Cipurpfifdroprate}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfDiscontinuityTime"] = types.YLeaf{"Cipurpfifdiscontinuitytime", cipurpfifmonentry.Cipurpfifdiscontinuitytime}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfDropRateNotifyEnable"] = types.YLeaf{"Cipurpfifdropratenotifyenable", cipurpfifmonentry.Cipurpfifdropratenotifyenable}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfNotifyDropRateThreshold"] = types.YLeaf{"Cipurpfifnotifydropratethreshold", cipurpfifmonentry.Cipurpfifnotifydropratethreshold}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfNotifyDrHoldDownReset"] = types.YLeaf{"Cipurpfifnotifydrholddownreset", cipurpfifmonentry.Cipurpfifnotifydrholddownreset}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfCheckStrict"] = types.YLeaf{"Cipurpfifcheckstrict", cipurpfifmonentry.Cipurpfifcheckstrict}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfWhichRouteTableID"] = types.YLeaf{"Cipurpfifwhichroutetableid", cipurpfifmonentry.Cipurpfifwhichroutetableid}
    cipurpfifmonentry.EntityData.Leafs["cipUrpfIfVrfName"] = types.YLeaf{"Cipurpfifvrfname", cipurpfifmonentry.Cipurpfifvrfname}
    return &(cipurpfifmonentry.EntityData)
}

// CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifcheckstrict represents           some interface on the device.
type CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifcheckstrict string

const (
    CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifcheckstrict_strict CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifcheckstrict = "strict"

    CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifcheckstrict_loose CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifcheckstrict = "loose"
)

// CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifipversion represents configuration apply.
type CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifipversion string

const (
    CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifipversion_ipv4 CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifipversion = "ipv4"

    CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifipversion_ipv6 CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifipversion = "ipv6"
)

// CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifwhichroutetableid represents - vrf   : a particular VPN routing table.
type CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifwhichroutetableid string

const (
    CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifwhichroutetableid_default_ CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifwhichroutetableid = "default"

    CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifwhichroutetableid_vrf CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry_Cipurpfifwhichroutetableid = "vrf"
)

// CISCOIPURPFMIB_Cipurpfvrfiftable
// This table contains statistics information for interfaces
// performing URPF using VRF table to determine reachability.
type CISCOIPURPFMIB_Cipurpfvrfiftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry exists for a VRF and interface if and only if the VRF associated
    // with the interface is configured  to perform IP URPF checking using the
    // routing  table for the VRF. The type is slice of
    // CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry.
    Cipurpfvrfifentry []CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry
}

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetEntityData() *types.CommonEntityData {
    cipurpfvrfiftable.EntityData.YFilter = cipurpfvrfiftable.YFilter
    cipurpfvrfiftable.EntityData.YangName = "cipUrpfVrfIfTable"
    cipurpfvrfiftable.EntityData.BundleName = "cisco_ios_xe"
    cipurpfvrfiftable.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipurpfvrfiftable.EntityData.SegmentPath = "cipUrpfVrfIfTable"
    cipurpfvrfiftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpfvrfiftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpfvrfiftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpfvrfiftable.EntityData.Children = make(map[string]types.YChild)
    cipurpfvrfiftable.EntityData.Children["cipUrpfVrfIfEntry"] = types.YChild{"Cipurpfvrfifentry", nil}
    for i := range cipurpfvrfiftable.Cipurpfvrfifentry {
        cipurpfvrfiftable.EntityData.Children[types.GetSegmentPath(&cipurpfvrfiftable.Cipurpfvrfifentry[i])] = types.YChild{"Cipurpfvrfifentry", &cipurpfvrfiftable.Cipurpfvrfifentry[i]}
    }
    cipurpfvrfiftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipurpfvrfiftable.EntityData)
}

// CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry
// An entry exists for a VRF and interface if and only
// if the VRF associated with the interface is configured 
// to perform IP URPF checking using the routing 
// table for the VRF.
type CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // cisco_ip_urpf_mib.CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry_Cipurpfvrfname
    Cipurpfvrfname interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // The number of packets failing the URPF check for a VRF on the interface and
    // dropped by the managed device.  Discontinuities in the value of this
    // variable can occur  at re-initialization of the management system, and at 
    // other times as indicated by the values of  cipUrpfVrfIfDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    Cipurpfvrfifdrops interface{}

    // The value of sysUpTime on the most recent occasion at which the URPF
    // counters for this VRF on this interface  suffered  a discontinuity.  If no
    // such discontinuities  have occurred since the last re-initialization of the
    // local management subsystem, then this object contains a  value of zero. The
    // type is interface{} with range: 0..4294967295.
    Cipurpfvrfifdiscontinuitytime interface{}
}

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetEntityData() *types.CommonEntityData {
    cipurpfvrfifentry.EntityData.YFilter = cipurpfvrfifentry.YFilter
    cipurpfvrfifentry.EntityData.YangName = "cipUrpfVrfIfEntry"
    cipurpfvrfifentry.EntityData.BundleName = "cisco_ios_xe"
    cipurpfvrfifentry.EntityData.ParentYangName = "cipUrpfVrfIfTable"
    cipurpfvrfifentry.EntityData.SegmentPath = "cipUrpfVrfIfEntry" + "[cipUrpfVrfName='" + fmt.Sprintf("%v", cipurpfvrfifentry.Cipurpfvrfname) + "']" + "[ifIndex='" + fmt.Sprintf("%v", cipurpfvrfifentry.Ifindex) + "']"
    cipurpfvrfifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpfvrfifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpfvrfifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpfvrfifentry.EntityData.Children = make(map[string]types.YChild)
    cipurpfvrfifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipurpfvrfifentry.EntityData.Leafs["cipUrpfVrfName"] = types.YLeaf{"Cipurpfvrfname", cipurpfvrfifentry.Cipurpfvrfname}
    cipurpfvrfifentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cipurpfvrfifentry.Ifindex}
    cipurpfvrfifentry.EntityData.Leafs["cipUrpfVrfIfDrops"] = types.YLeaf{"Cipurpfvrfifdrops", cipurpfvrfifentry.Cipurpfvrfifdrops}
    cipurpfvrfifentry.EntityData.Leafs["cipUrpfVrfIfDiscontinuityTime"] = types.YLeaf{"Cipurpfvrfifdiscontinuitytime", cipurpfvrfifentry.Cipurpfvrfifdiscontinuitytime}
    return &(cipurpfvrfifentry.EntityData)
}

// CISCOIPURPFMIB_Cipurpfvrftable
// This table enables indexing URPF drop statistics
// by Virtual Routing and Forwarding instances.
type CISCOIPURPFMIB_Cipurpfvrftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry exists for a VRF if and only if the VRF is associated with an
    // interface that is configured to perform IP URPF checking using the routing
    // table  for that VRF. The type is slice of
    // CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry.
    Cipurpfvrfentry []CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry
}

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetEntityData() *types.CommonEntityData {
    cipurpfvrftable.EntityData.YFilter = cipurpfvrftable.YFilter
    cipurpfvrftable.EntityData.YangName = "cipUrpfVrfTable"
    cipurpfvrftable.EntityData.BundleName = "cisco_ios_xe"
    cipurpfvrftable.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipurpfvrftable.EntityData.SegmentPath = "cipUrpfVrfTable"
    cipurpfvrftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpfvrftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpfvrftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpfvrftable.EntityData.Children = make(map[string]types.YChild)
    cipurpfvrftable.EntityData.Children["cipUrpfVrfEntry"] = types.YChild{"Cipurpfvrfentry", nil}
    for i := range cipurpfvrftable.Cipurpfvrfentry {
        cipurpfvrftable.EntityData.Children[types.GetSegmentPath(&cipurpfvrftable.Cipurpfvrfentry[i])] = types.YChild{"Cipurpfvrfentry", &cipurpfvrftable.Cipurpfvrfentry[i]}
    }
    cipurpfvrftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipurpfvrftable.EntityData)
}

// CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry
// An entry exists for a VRF if and only if the VRF
// is associated with an interface that is configured
// to perform IP URPF checking using the routing table 
// for that VRF.
type CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This field is used to specify the VRF Table name.
    // The type is string with length: 0..32.
    Cipurpfvrfname interface{}
}

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetEntityData() *types.CommonEntityData {
    cipurpfvrfentry.EntityData.YFilter = cipurpfvrfentry.YFilter
    cipurpfvrfentry.EntityData.YangName = "cipUrpfVrfEntry"
    cipurpfvrfentry.EntityData.BundleName = "cisco_ios_xe"
    cipurpfvrfentry.EntityData.ParentYangName = "cipUrpfVrfTable"
    cipurpfvrfentry.EntityData.SegmentPath = "cipUrpfVrfEntry" + "[cipUrpfVrfName='" + fmt.Sprintf("%v", cipurpfvrfentry.Cipurpfvrfname) + "']"
    cipurpfvrfentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipurpfvrfentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipurpfvrfentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipurpfvrfentry.EntityData.Children = make(map[string]types.YChild)
    cipurpfvrfentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipurpfvrfentry.EntityData.Leafs["cipUrpfVrfName"] = types.YLeaf{"Cipurpfvrfname", cipurpfvrfentry.Cipurpfvrfname}
    return &(cipurpfvrfentry.EntityData)
}

