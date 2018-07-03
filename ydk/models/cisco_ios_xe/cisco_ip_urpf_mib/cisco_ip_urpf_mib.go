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

    
    CipUrpfScalar CISCOIPURPFMIB_CipUrpfScalar

    // This table contains summary information for the managed device on URPF
    // dropping.
    CipUrpfTable CISCOIPURPFMIB_CipUrpfTable

    // This table contains information on URPF dropping on an interface.
    CipUrpfIfMonTable CISCOIPURPFMIB_CipUrpfIfMonTable

    // This table contains statistics information for interfaces performing URPF
    // using VRF table to determine reachability.
    CipUrpfVrfIfTable CISCOIPURPFMIB_CipUrpfVrfIfTable

    // This table enables indexing URPF drop statistics by Virtual Routing and
    // Forwarding instances.
    CipUrpfVrfTable CISCOIPURPFMIB_CipUrpfVrfTable
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

    cISCOIPURPFMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPURPFMIB.EntityData.Children.Append("cipUrpfScalar", types.YChild{"CipUrpfScalar", &cISCOIPURPFMIB.CipUrpfScalar})
    cISCOIPURPFMIB.EntityData.Children.Append("cipUrpfTable", types.YChild{"CipUrpfTable", &cISCOIPURPFMIB.CipUrpfTable})
    cISCOIPURPFMIB.EntityData.Children.Append("cipUrpfIfMonTable", types.YChild{"CipUrpfIfMonTable", &cISCOIPURPFMIB.CipUrpfIfMonTable})
    cISCOIPURPFMIB.EntityData.Children.Append("cipUrpfVrfIfTable", types.YChild{"CipUrpfVrfIfTable", &cISCOIPURPFMIB.CipUrpfVrfIfTable})
    cISCOIPURPFMIB.EntityData.Children.Append("cipUrpfVrfTable", types.YChild{"CipUrpfVrfTable", &cISCOIPURPFMIB.CipUrpfVrfTable})
    cISCOIPURPFMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPURPFMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPURPFMIB.EntityData)
}

// CISCOIPURPFMIB_CipUrpfScalar
type CISCOIPURPFMIB_CipUrpfScalar struct {
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
    CipUrpfDropRateWindow interface{}

    // The time between rate computations. This global value applies for the
    // computation of all URPF rates, global and per-interface.  When the value of
    // cipUrpfComputeInterval is changed, the interval in-progress proceeds as
    // though the value had not changed. The change will apply to the length of
    // subsequent intervals.  The cipUrpfComputeInterval must be less than or
    // equal  to the cipUrpfDropRateWindow. The type is interface{} with range:
    // 1..120. Units are seconds.
    CipUrpfComputeInterval interface{}

    // The minimum time between issuance of cipUrpfIfDropRateNotify notifications
    // for a  particular interface and packet forwarding type.  Notifications are
    // generated for each interface and packet forwarding type that exceeds the
    // drop-rate.  When a Notify is sent because the drop-rate is  exceeded for a
    // particular interface and forwarding type, the time specified by this object
    // is used to  specify the minimum time that must elapse before  another
    // Notify can be sent for that interface and forwarding type. The time is
    // specified globally but  used individually. The type is interface{} with
    // range: 1..1000. Units are seconds.
    CipUrpfDropNotifyHoldDownTime interface{}
}

func (cipUrpfScalar *CISCOIPURPFMIB_CipUrpfScalar) GetEntityData() *types.CommonEntityData {
    cipUrpfScalar.EntityData.YFilter = cipUrpfScalar.YFilter
    cipUrpfScalar.EntityData.YangName = "cipUrpfScalar"
    cipUrpfScalar.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfScalar.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipUrpfScalar.EntityData.SegmentPath = "cipUrpfScalar"
    cipUrpfScalar.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfScalar.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfScalar.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfScalar.EntityData.Children = types.NewOrderedMap()
    cipUrpfScalar.EntityData.Leafs = types.NewOrderedMap()
    cipUrpfScalar.EntityData.Leafs.Append("cipUrpfDropRateWindow", types.YLeaf{"CipUrpfDropRateWindow", cipUrpfScalar.CipUrpfDropRateWindow})
    cipUrpfScalar.EntityData.Leafs.Append("cipUrpfComputeInterval", types.YLeaf{"CipUrpfComputeInterval", cipUrpfScalar.CipUrpfComputeInterval})
    cipUrpfScalar.EntityData.Leafs.Append("cipUrpfDropNotifyHoldDownTime", types.YLeaf{"CipUrpfDropNotifyHoldDownTime", cipUrpfScalar.CipUrpfDropNotifyHoldDownTime})

    cipUrpfScalar.EntityData.YListKeys = []string {}

    return &(cipUrpfScalar.EntityData)
}

// CISCOIPURPFMIB_CipUrpfTable
// This table contains summary information for the
// managed device on URPF dropping.
type CISCOIPURPFMIB_CipUrpfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the managed device supports URPF dropping, a row exists for each IP
    // version type (v4 and v6). A row contains summary information on URPF
    // dropping over the entire managed device. The type is slice of
    // CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry.
    CipUrpfEntry []*CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry
}

func (cipUrpfTable *CISCOIPURPFMIB_CipUrpfTable) GetEntityData() *types.CommonEntityData {
    cipUrpfTable.EntityData.YFilter = cipUrpfTable.YFilter
    cipUrpfTable.EntityData.YangName = "cipUrpfTable"
    cipUrpfTable.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfTable.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipUrpfTable.EntityData.SegmentPath = "cipUrpfTable"
    cipUrpfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfTable.EntityData.Children = types.NewOrderedMap()
    cipUrpfTable.EntityData.Children.Append("cipUrpfEntry", types.YChild{"CipUrpfEntry", nil})
    for i := range cipUrpfTable.CipUrpfEntry {
        cipUrpfTable.EntityData.Children.Append(types.GetSegmentPath(cipUrpfTable.CipUrpfEntry[i]), types.YChild{"CipUrpfEntry", cipUrpfTable.CipUrpfEntry[i]})
    }
    cipUrpfTable.EntityData.Leafs = types.NewOrderedMap()

    cipUrpfTable.EntityData.YListKeys = []string {}

    return &(cipUrpfTable.EntityData)
}

// CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry
// If the managed device supports URPF dropping,
// a row exists for each IP version type (v4 and v6).
// A row contains summary information on URPF
// dropping over the entire managed device.
type CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specifies the version of IP forwarding on an
    // interface to which the table row URPF counts, rates, and configuration
    // apply. The type is CipUrpfIpVersion.
    CipUrpfIpVersion interface{}

    // Sum of dropped IP version cipUrpfIpVersion packets failing a URPF check.
    // This value is the sum of drops of packets  received on all interfaces of
    // the managed device. The type is interface{} with range: 0..4294967295.
    // Units are packets.
    CipUrpfDrops interface{}

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
    CipUrpfDropRate interface{}
}

func (cipUrpfEntry *CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry) GetEntityData() *types.CommonEntityData {
    cipUrpfEntry.EntityData.YFilter = cipUrpfEntry.YFilter
    cipUrpfEntry.EntityData.YangName = "cipUrpfEntry"
    cipUrpfEntry.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfEntry.EntityData.ParentYangName = "cipUrpfTable"
    cipUrpfEntry.EntityData.SegmentPath = "cipUrpfEntry" + types.AddKeyToken(cipUrpfEntry.CipUrpfIpVersion, "cipUrpfIpVersion")
    cipUrpfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfEntry.EntityData.Children = types.NewOrderedMap()
    cipUrpfEntry.EntityData.Leafs = types.NewOrderedMap()
    cipUrpfEntry.EntityData.Leafs.Append("cipUrpfIpVersion", types.YLeaf{"CipUrpfIpVersion", cipUrpfEntry.CipUrpfIpVersion})
    cipUrpfEntry.EntityData.Leafs.Append("cipUrpfDrops", types.YLeaf{"CipUrpfDrops", cipUrpfEntry.CipUrpfDrops})
    cipUrpfEntry.EntityData.Leafs.Append("cipUrpfDropRate", types.YLeaf{"CipUrpfDropRate", cipUrpfEntry.CipUrpfDropRate})

    cipUrpfEntry.EntityData.YListKeys = []string {"CipUrpfIpVersion"}

    return &(cipUrpfEntry.EntityData)
}

// CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry_CipUrpfIpVersion represents configuration apply.
type CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry_CipUrpfIpVersion string

const (
    CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry_CipUrpfIpVersion_ipv4 CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry_CipUrpfIpVersion = "ipv4"

    CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry_CipUrpfIpVersion_ipv6 CISCOIPURPFMIB_CipUrpfTable_CipUrpfEntry_CipUrpfIpVersion = "ipv6"
)

// CISCOIPURPFMIB_CipUrpfIfMonTable
// This table contains information on URPF dropping on
// an interface.
type CISCOIPURPFMIB_CipUrpfIfMonTable struct {
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
    // CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry.
    CipUrpfIfMonEntry []*CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry
}

func (cipUrpfIfMonTable *CISCOIPURPFMIB_CipUrpfIfMonTable) GetEntityData() *types.CommonEntityData {
    cipUrpfIfMonTable.EntityData.YFilter = cipUrpfIfMonTable.YFilter
    cipUrpfIfMonTable.EntityData.YangName = "cipUrpfIfMonTable"
    cipUrpfIfMonTable.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfIfMonTable.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipUrpfIfMonTable.EntityData.SegmentPath = "cipUrpfIfMonTable"
    cipUrpfIfMonTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfIfMonTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfIfMonTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfIfMonTable.EntityData.Children = types.NewOrderedMap()
    cipUrpfIfMonTable.EntityData.Children.Append("cipUrpfIfMonEntry", types.YChild{"CipUrpfIfMonEntry", nil})
    for i := range cipUrpfIfMonTable.CipUrpfIfMonEntry {
        cipUrpfIfMonTable.EntityData.Children.Append(types.GetSegmentPath(cipUrpfIfMonTable.CipUrpfIfMonEntry[i]), types.YChild{"CipUrpfIfMonEntry", cipUrpfIfMonTable.CipUrpfIfMonEntry[i]})
    }
    cipUrpfIfMonTable.EntityData.Leafs = types.NewOrderedMap()

    cipUrpfIfMonTable.EntityData.YListKeys = []string {}

    return &(cipUrpfIfMonTable.EntityData)
}

// CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry
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
type CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. Specifies the version of IP forwarding on an
    // interface to which the table row URPF counts, rates, and  configuration
    // apply. The type is CipUrpfIfIpVersion.
    CipUrpfIfIpVersion interface{}

    // The number of IP packets of version cipUrpfIfIpVersion failing the URPF
    // check and dropped by the managed device on a particular interface. 
    // Discontinuities in the value of this variable can occur  at
    // re-initialization of the management system, and at  other times as
    // indicated by the values of  cipUrpfIfDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295. Units are packets.
    CipUrpfIfDrops interface{}

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
    CipUrpfIfSuppressedDrops interface{}

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
    CipUrpfIfDropRate interface{}

    // The value of sysUpTime on the most recent occasion at which this
    // interface's  counters suffered  a discontinuity. If no such discontinuities
    // have occurred since the last re-initialization of the local management
    // subsystem, then this object contains a value of zero. The type is
    // interface{} with range: 0..4294967295.
    CipUrpfIfDiscontinuityTime interface{}

    // This object specifies whether the system produces the
    // cipUrpfIfDropRateNotify notification as a result of URPF  dropping of
    // version cipUrpfIfIpVersion IP packets on this  interface. A false value
    // prevents such notifications from  being generated by this system. The type
    // is bool.
    CipUrpfIfDropRateNotifyEnable interface{}

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
    CipUrpfIfNotifyDropRateThreshold interface{}

    // Setting this object to true causes the five-minute hold-down timer for
    // emitting URPF drop rate  notifications for IP version cipUrpfIfIpVersion on
    // the interface to be short-circuited.  If a notification  is due and would
    // be emitted for the interface if the  five-minutes elapsed, setting this
    // object will cause  the notification to be sent.  This is a trigger, and
    // doesn't hold information. It is set and an action is performed. Therefore a
    // get for  this object always returns false. The type is bool.
    CipUrpfIfNotifyDrHoldDownReset interface{}

    // Interface configuration indicating the strictness of the reachability check
    // performed  on the interface. - strict: check that source addr is reachable
    // via            the interface it came in on. - loose : check that source
    // addr is reachable via            some interface on the device. The type is
    // CipUrpfIfCheckStrict.
    CipUrpfIfCheckStrict interface{}

    // Interface configuration indicating the routing table consulted for the
    // reachability check: - default: the non-private routing table for of the    
    // managed system. - vrf   : a particular VPN routing table. The type is
    // CipUrpfIfWhichRouteTableID.
    CipUrpfIfWhichRouteTableID interface{}

    // If the value of cipUrpfIfWhichRouteTableID is 'vrf', the name of the VRF
    // Table. Otherwise a zero-length string. The type is string with length:
    // 0..32.
    CipUrpfIfVrfName interface{}
}

func (cipUrpfIfMonEntry *CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry) GetEntityData() *types.CommonEntityData {
    cipUrpfIfMonEntry.EntityData.YFilter = cipUrpfIfMonEntry.YFilter
    cipUrpfIfMonEntry.EntityData.YangName = "cipUrpfIfMonEntry"
    cipUrpfIfMonEntry.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfIfMonEntry.EntityData.ParentYangName = "cipUrpfIfMonTable"
    cipUrpfIfMonEntry.EntityData.SegmentPath = "cipUrpfIfMonEntry" + types.AddKeyToken(cipUrpfIfMonEntry.IfIndex, "ifIndex") + types.AddKeyToken(cipUrpfIfMonEntry.CipUrpfIfIpVersion, "cipUrpfIfIpVersion")
    cipUrpfIfMonEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfIfMonEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfIfMonEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfIfMonEntry.EntityData.Children = types.NewOrderedMap()
    cipUrpfIfMonEntry.EntityData.Leafs = types.NewOrderedMap()
    cipUrpfIfMonEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cipUrpfIfMonEntry.IfIndex})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfIpVersion", types.YLeaf{"CipUrpfIfIpVersion", cipUrpfIfMonEntry.CipUrpfIfIpVersion})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfDrops", types.YLeaf{"CipUrpfIfDrops", cipUrpfIfMonEntry.CipUrpfIfDrops})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfSuppressedDrops", types.YLeaf{"CipUrpfIfSuppressedDrops", cipUrpfIfMonEntry.CipUrpfIfSuppressedDrops})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfDropRate", types.YLeaf{"CipUrpfIfDropRate", cipUrpfIfMonEntry.CipUrpfIfDropRate})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfDiscontinuityTime", types.YLeaf{"CipUrpfIfDiscontinuityTime", cipUrpfIfMonEntry.CipUrpfIfDiscontinuityTime})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfDropRateNotifyEnable", types.YLeaf{"CipUrpfIfDropRateNotifyEnable", cipUrpfIfMonEntry.CipUrpfIfDropRateNotifyEnable})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfNotifyDropRateThreshold", types.YLeaf{"CipUrpfIfNotifyDropRateThreshold", cipUrpfIfMonEntry.CipUrpfIfNotifyDropRateThreshold})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfNotifyDrHoldDownReset", types.YLeaf{"CipUrpfIfNotifyDrHoldDownReset", cipUrpfIfMonEntry.CipUrpfIfNotifyDrHoldDownReset})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfCheckStrict", types.YLeaf{"CipUrpfIfCheckStrict", cipUrpfIfMonEntry.CipUrpfIfCheckStrict})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfWhichRouteTableID", types.YLeaf{"CipUrpfIfWhichRouteTableID", cipUrpfIfMonEntry.CipUrpfIfWhichRouteTableID})
    cipUrpfIfMonEntry.EntityData.Leafs.Append("cipUrpfIfVrfName", types.YLeaf{"CipUrpfIfVrfName", cipUrpfIfMonEntry.CipUrpfIfVrfName})

    cipUrpfIfMonEntry.EntityData.YListKeys = []string {"IfIndex", "CipUrpfIfIpVersion"}

    return &(cipUrpfIfMonEntry.EntityData)
}

// CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfCheckStrict represents           some interface on the device.
type CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfCheckStrict string

const (
    CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfCheckStrict_strict CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfCheckStrict = "strict"

    CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfCheckStrict_loose CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfCheckStrict = "loose"
)

// CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfIpVersion represents configuration apply.
type CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfIpVersion string

const (
    CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfIpVersion_ipv4 CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfIpVersion = "ipv4"

    CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfIpVersion_ipv6 CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfIpVersion = "ipv6"
)

// CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfWhichRouteTableID represents - vrf   : a particular VPN routing table.
type CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfWhichRouteTableID string

const (
    CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfWhichRouteTableID_default_ CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfWhichRouteTableID = "default"

    CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfWhichRouteTableID_vrf CISCOIPURPFMIB_CipUrpfIfMonTable_CipUrpfIfMonEntry_CipUrpfIfWhichRouteTableID = "vrf"
)

// CISCOIPURPFMIB_CipUrpfVrfIfTable
// This table contains statistics information for interfaces
// performing URPF using VRF table to determine reachability.
type CISCOIPURPFMIB_CipUrpfVrfIfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry exists for a VRF and interface if and only if the VRF associated
    // with the interface is configured  to perform IP URPF checking using the
    // routing  table for the VRF. The type is slice of
    // CISCOIPURPFMIB_CipUrpfVrfIfTable_CipUrpfVrfIfEntry.
    CipUrpfVrfIfEntry []*CISCOIPURPFMIB_CipUrpfVrfIfTable_CipUrpfVrfIfEntry
}

func (cipUrpfVrfIfTable *CISCOIPURPFMIB_CipUrpfVrfIfTable) GetEntityData() *types.CommonEntityData {
    cipUrpfVrfIfTable.EntityData.YFilter = cipUrpfVrfIfTable.YFilter
    cipUrpfVrfIfTable.EntityData.YangName = "cipUrpfVrfIfTable"
    cipUrpfVrfIfTable.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfVrfIfTable.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipUrpfVrfIfTable.EntityData.SegmentPath = "cipUrpfVrfIfTable"
    cipUrpfVrfIfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfVrfIfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfVrfIfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfVrfIfTable.EntityData.Children = types.NewOrderedMap()
    cipUrpfVrfIfTable.EntityData.Children.Append("cipUrpfVrfIfEntry", types.YChild{"CipUrpfVrfIfEntry", nil})
    for i := range cipUrpfVrfIfTable.CipUrpfVrfIfEntry {
        cipUrpfVrfIfTable.EntityData.Children.Append(types.GetSegmentPath(cipUrpfVrfIfTable.CipUrpfVrfIfEntry[i]), types.YChild{"CipUrpfVrfIfEntry", cipUrpfVrfIfTable.CipUrpfVrfIfEntry[i]})
    }
    cipUrpfVrfIfTable.EntityData.Leafs = types.NewOrderedMap()

    cipUrpfVrfIfTable.EntityData.YListKeys = []string {}

    return &(cipUrpfVrfIfTable.EntityData)
}

// CISCOIPURPFMIB_CipUrpfVrfIfTable_CipUrpfVrfIfEntry
// An entry exists for a VRF and interface if and only
// if the VRF associated with the interface is configured 
// to perform IP URPF checking using the routing 
// table for the VRF.
type CISCOIPURPFMIB_CipUrpfVrfIfTable_CipUrpfVrfIfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // cisco_ip_urpf_mib.CISCOIPURPFMIB_CipUrpfVrfTable_CipUrpfVrfEntry_CipUrpfVrfName
    CipUrpfVrfName interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The number of packets failing the URPF check for a VRF on the interface and
    // dropped by the managed device.  Discontinuities in the value of this
    // variable can occur  at re-initialization of the management system, and at 
    // other times as indicated by the values of  cipUrpfVrfIfDiscontinuityTime.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    CipUrpfVrfIfDrops interface{}

    // The value of sysUpTime on the most recent occasion at which the URPF
    // counters for this VRF on this interface  suffered  a discontinuity.  If no
    // such discontinuities  have occurred since the last re-initialization of the
    // local management subsystem, then this object contains a  value of zero. The
    // type is interface{} with range: 0..4294967295.
    CipUrpfVrfIfDiscontinuityTime interface{}
}

func (cipUrpfVrfIfEntry *CISCOIPURPFMIB_CipUrpfVrfIfTable_CipUrpfVrfIfEntry) GetEntityData() *types.CommonEntityData {
    cipUrpfVrfIfEntry.EntityData.YFilter = cipUrpfVrfIfEntry.YFilter
    cipUrpfVrfIfEntry.EntityData.YangName = "cipUrpfVrfIfEntry"
    cipUrpfVrfIfEntry.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfVrfIfEntry.EntityData.ParentYangName = "cipUrpfVrfIfTable"
    cipUrpfVrfIfEntry.EntityData.SegmentPath = "cipUrpfVrfIfEntry" + types.AddKeyToken(cipUrpfVrfIfEntry.CipUrpfVrfName, "cipUrpfVrfName") + types.AddKeyToken(cipUrpfVrfIfEntry.IfIndex, "ifIndex")
    cipUrpfVrfIfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfVrfIfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfVrfIfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfVrfIfEntry.EntityData.Children = types.NewOrderedMap()
    cipUrpfVrfIfEntry.EntityData.Leafs = types.NewOrderedMap()
    cipUrpfVrfIfEntry.EntityData.Leafs.Append("cipUrpfVrfName", types.YLeaf{"CipUrpfVrfName", cipUrpfVrfIfEntry.CipUrpfVrfName})
    cipUrpfVrfIfEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cipUrpfVrfIfEntry.IfIndex})
    cipUrpfVrfIfEntry.EntityData.Leafs.Append("cipUrpfVrfIfDrops", types.YLeaf{"CipUrpfVrfIfDrops", cipUrpfVrfIfEntry.CipUrpfVrfIfDrops})
    cipUrpfVrfIfEntry.EntityData.Leafs.Append("cipUrpfVrfIfDiscontinuityTime", types.YLeaf{"CipUrpfVrfIfDiscontinuityTime", cipUrpfVrfIfEntry.CipUrpfVrfIfDiscontinuityTime})

    cipUrpfVrfIfEntry.EntityData.YListKeys = []string {"CipUrpfVrfName", "IfIndex"}

    return &(cipUrpfVrfIfEntry.EntityData)
}

// CISCOIPURPFMIB_CipUrpfVrfTable
// This table enables indexing URPF drop statistics
// by Virtual Routing and Forwarding instances.
type CISCOIPURPFMIB_CipUrpfVrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry exists for a VRF if and only if the VRF is associated with an
    // interface that is configured to perform IP URPF checking using the routing
    // table  for that VRF. The type is slice of
    // CISCOIPURPFMIB_CipUrpfVrfTable_CipUrpfVrfEntry.
    CipUrpfVrfEntry []*CISCOIPURPFMIB_CipUrpfVrfTable_CipUrpfVrfEntry
}

func (cipUrpfVrfTable *CISCOIPURPFMIB_CipUrpfVrfTable) GetEntityData() *types.CommonEntityData {
    cipUrpfVrfTable.EntityData.YFilter = cipUrpfVrfTable.YFilter
    cipUrpfVrfTable.EntityData.YangName = "cipUrpfVrfTable"
    cipUrpfVrfTable.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfVrfTable.EntityData.ParentYangName = "CISCO-IP-URPF-MIB"
    cipUrpfVrfTable.EntityData.SegmentPath = "cipUrpfVrfTable"
    cipUrpfVrfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfVrfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfVrfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfVrfTable.EntityData.Children = types.NewOrderedMap()
    cipUrpfVrfTable.EntityData.Children.Append("cipUrpfVrfEntry", types.YChild{"CipUrpfVrfEntry", nil})
    for i := range cipUrpfVrfTable.CipUrpfVrfEntry {
        cipUrpfVrfTable.EntityData.Children.Append(types.GetSegmentPath(cipUrpfVrfTable.CipUrpfVrfEntry[i]), types.YChild{"CipUrpfVrfEntry", cipUrpfVrfTable.CipUrpfVrfEntry[i]})
    }
    cipUrpfVrfTable.EntityData.Leafs = types.NewOrderedMap()

    cipUrpfVrfTable.EntityData.YListKeys = []string {}

    return &(cipUrpfVrfTable.EntityData)
}

// CISCOIPURPFMIB_CipUrpfVrfTable_CipUrpfVrfEntry
// An entry exists for a VRF if and only if the VRF
// is associated with an interface that is configured
// to perform IP URPF checking using the routing table 
// for that VRF.
type CISCOIPURPFMIB_CipUrpfVrfTable_CipUrpfVrfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This field is used to specify the VRF Table name.
    // The type is string with length: 0..32.
    CipUrpfVrfName interface{}
}

func (cipUrpfVrfEntry *CISCOIPURPFMIB_CipUrpfVrfTable_CipUrpfVrfEntry) GetEntityData() *types.CommonEntityData {
    cipUrpfVrfEntry.EntityData.YFilter = cipUrpfVrfEntry.YFilter
    cipUrpfVrfEntry.EntityData.YangName = "cipUrpfVrfEntry"
    cipUrpfVrfEntry.EntityData.BundleName = "cisco_ios_xe"
    cipUrpfVrfEntry.EntityData.ParentYangName = "cipUrpfVrfTable"
    cipUrpfVrfEntry.EntityData.SegmentPath = "cipUrpfVrfEntry" + types.AddKeyToken(cipUrpfVrfEntry.CipUrpfVrfName, "cipUrpfVrfName")
    cipUrpfVrfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipUrpfVrfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipUrpfVrfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipUrpfVrfEntry.EntityData.Children = types.NewOrderedMap()
    cipUrpfVrfEntry.EntityData.Leafs = types.NewOrderedMap()
    cipUrpfVrfEntry.EntityData.Leafs.Append("cipUrpfVrfName", types.YLeaf{"CipUrpfVrfName", cipUrpfVrfEntry.CipUrpfVrfName})

    cipUrpfVrfEntry.EntityData.YListKeys = []string {"CipUrpfVrfName"}

    return &(cipUrpfVrfEntry.EntityData)
}

