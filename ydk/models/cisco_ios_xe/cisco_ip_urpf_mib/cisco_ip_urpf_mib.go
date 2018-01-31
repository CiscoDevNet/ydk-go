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
    parent types.Entity
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

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetFilter() yfilter.YFilter { return cISCOIPURPFMIB.YFilter }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) SetFilter(yf yfilter.YFilter) { cISCOIPURPFMIB.YFilter = yf }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetGoName(yname string) string {
    if yname == "cipUrpfScalar" { return "Cipurpfscalar" }
    if yname == "cipUrpfTable" { return "Cipurpftable" }
    if yname == "cipUrpfIfMonTable" { return "Cipurpfifmontable" }
    if yname == "cipUrpfVrfIfTable" { return "Cipurpfvrfiftable" }
    if yname == "cipUrpfVrfTable" { return "Cipurpfvrftable" }
    return ""
}

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetSegmentPath() string {
    return "CISCO-IP-URPF-MIB:CISCO-IP-URPF-MIB"
}

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipUrpfScalar" {
        return &cISCOIPURPFMIB.Cipurpfscalar
    }
    if childYangName == "cipUrpfTable" {
        return &cISCOIPURPFMIB.Cipurpftable
    }
    if childYangName == "cipUrpfIfMonTable" {
        return &cISCOIPURPFMIB.Cipurpfifmontable
    }
    if childYangName == "cipUrpfVrfIfTable" {
        return &cISCOIPURPFMIB.Cipurpfvrfiftable
    }
    if childYangName == "cipUrpfVrfTable" {
        return &cISCOIPURPFMIB.Cipurpfvrftable
    }
    return nil
}

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cipUrpfScalar"] = &cISCOIPURPFMIB.Cipurpfscalar
    children["cipUrpfTable"] = &cISCOIPURPFMIB.Cipurpftable
    children["cipUrpfIfMonTable"] = &cISCOIPURPFMIB.Cipurpfifmontable
    children["cipUrpfVrfIfTable"] = &cISCOIPURPFMIB.Cipurpfvrfiftable
    children["cipUrpfVrfTable"] = &cISCOIPURPFMIB.Cipurpfvrftable
    return children
}

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetYangName() string { return "CISCO-IP-URPF-MIB" }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) SetParent(parent types.Entity) { cISCOIPURPFMIB.parent = parent }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetParent() types.Entity { return cISCOIPURPFMIB.parent }

func (cISCOIPURPFMIB *CISCOIPURPFMIB) GetParentYangName() string { return "CISCO-IP-URPF-MIB" }

// CISCOIPURPFMIB_Cipurpfscalar
type CISCOIPURPFMIB_Cipurpfscalar struct {
    parent types.Entity
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

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetFilter() yfilter.YFilter { return cipurpfscalar.YFilter }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) SetFilter(yf yfilter.YFilter) { cipurpfscalar.YFilter = yf }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetGoName(yname string) string {
    if yname == "cipUrpfDropRateWindow" { return "Cipurpfdropratewindow" }
    if yname == "cipUrpfComputeInterval" { return "Cipurpfcomputeinterval" }
    if yname == "cipUrpfDropNotifyHoldDownTime" { return "Cipurpfdropnotifyholddowntime" }
    return ""
}

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetSegmentPath() string {
    return "cipUrpfScalar"
}

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipUrpfDropRateWindow"] = cipurpfscalar.Cipurpfdropratewindow
    leafs["cipUrpfComputeInterval"] = cipurpfscalar.Cipurpfcomputeinterval
    leafs["cipUrpfDropNotifyHoldDownTime"] = cipurpfscalar.Cipurpfdropnotifyholddowntime
    return leafs
}

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetYangName() string { return "cipUrpfScalar" }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) SetParent(parent types.Entity) { cipurpfscalar.parent = parent }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetParent() types.Entity { return cipurpfscalar.parent }

func (cipurpfscalar *CISCOIPURPFMIB_Cipurpfscalar) GetParentYangName() string { return "CISCO-IP-URPF-MIB" }

// CISCOIPURPFMIB_Cipurpftable
// This table contains summary information for the
// managed device on URPF dropping.
type CISCOIPURPFMIB_Cipurpftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If the managed device supports URPF dropping, a row exists for each IP
    // version type (v4 and v6). A row contains summary information on URPF
    // dropping over the entire managed device. The type is slice of
    // CISCOIPURPFMIB_Cipurpftable_Cipurpfentry.
    Cipurpfentry []CISCOIPURPFMIB_Cipurpftable_Cipurpfentry
}

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetFilter() yfilter.YFilter { return cipurpftable.YFilter }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) SetFilter(yf yfilter.YFilter) { cipurpftable.YFilter = yf }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetGoName(yname string) string {
    if yname == "cipUrpfEntry" { return "Cipurpfentry" }
    return ""
}

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetSegmentPath() string {
    return "cipUrpfTable"
}

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipUrpfEntry" {
        for _, c := range cipurpftable.Cipurpfentry {
            if cipurpftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPURPFMIB_Cipurpftable_Cipurpfentry{}
        cipurpftable.Cipurpfentry = append(cipurpftable.Cipurpfentry, child)
        return &cipurpftable.Cipurpfentry[len(cipurpftable.Cipurpfentry)-1]
    }
    return nil
}

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipurpftable.Cipurpfentry {
        children[cipurpftable.Cipurpfentry[i].GetSegmentPath()] = &cipurpftable.Cipurpfentry[i]
    }
    return children
}

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetYangName() string { return "cipUrpfTable" }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) SetParent(parent types.Entity) { cipurpftable.parent = parent }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetParent() types.Entity { return cipurpftable.parent }

func (cipurpftable *CISCOIPURPFMIB_Cipurpftable) GetParentYangName() string { return "CISCO-IP-URPF-MIB" }

// CISCOIPURPFMIB_Cipurpftable_Cipurpfentry
// If the managed device supports URPF dropping,
// a row exists for each IP version type (v4 and v6).
// A row contains summary information on URPF
// dropping over the entire managed device.
type CISCOIPURPFMIB_Cipurpftable_Cipurpfentry struct {
    parent types.Entity
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

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetFilter() yfilter.YFilter { return cipurpfentry.YFilter }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) SetFilter(yf yfilter.YFilter) { cipurpfentry.YFilter = yf }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetGoName(yname string) string {
    if yname == "cipUrpfIpVersion" { return "Cipurpfipversion" }
    if yname == "cipUrpfDrops" { return "Cipurpfdrops" }
    if yname == "cipUrpfDropRate" { return "Cipurpfdroprate" }
    return ""
}

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetSegmentPath() string {
    return "cipUrpfEntry" + "[cipUrpfIpVersion='" + fmt.Sprintf("%v", cipurpfentry.Cipurpfipversion) + "']"
}

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipUrpfIpVersion"] = cipurpfentry.Cipurpfipversion
    leafs["cipUrpfDrops"] = cipurpfentry.Cipurpfdrops
    leafs["cipUrpfDropRate"] = cipurpfentry.Cipurpfdroprate
    return leafs
}

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetYangName() string { return "cipUrpfEntry" }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) SetParent(parent types.Entity) { cipurpfentry.parent = parent }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetParent() types.Entity { return cipurpfentry.parent }

func (cipurpfentry *CISCOIPURPFMIB_Cipurpftable_Cipurpfentry) GetParentYangName() string { return "cipUrpfTable" }

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
    parent types.Entity
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

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetFilter() yfilter.YFilter { return cipurpfifmontable.YFilter }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) SetFilter(yf yfilter.YFilter) { cipurpfifmontable.YFilter = yf }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetGoName(yname string) string {
    if yname == "cipUrpfIfMonEntry" { return "Cipurpfifmonentry" }
    return ""
}

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetSegmentPath() string {
    return "cipUrpfIfMonTable"
}

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipUrpfIfMonEntry" {
        for _, c := range cipurpfifmontable.Cipurpfifmonentry {
            if cipurpfifmontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry{}
        cipurpfifmontable.Cipurpfifmonentry = append(cipurpfifmontable.Cipurpfifmonentry, child)
        return &cipurpfifmontable.Cipurpfifmonentry[len(cipurpfifmontable.Cipurpfifmonentry)-1]
    }
    return nil
}

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipurpfifmontable.Cipurpfifmonentry {
        children[cipurpfifmontable.Cipurpfifmonentry[i].GetSegmentPath()] = &cipurpfifmontable.Cipurpfifmonentry[i]
    }
    return children
}

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetYangName() string { return "cipUrpfIfMonTable" }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) SetParent(parent types.Entity) { cipurpfifmontable.parent = parent }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetParent() types.Entity { return cipurpfifmontable.parent }

func (cipurpfifmontable *CISCOIPURPFMIB_Cipurpfifmontable) GetParentYangName() string { return "CISCO-IP-URPF-MIB" }

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
    parent types.Entity
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

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetFilter() yfilter.YFilter { return cipurpfifmonentry.YFilter }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) SetFilter(yf yfilter.YFilter) { cipurpfifmonentry.YFilter = yf }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cipUrpfIfIpVersion" { return "Cipurpfifipversion" }
    if yname == "cipUrpfIfDrops" { return "Cipurpfifdrops" }
    if yname == "cipUrpfIfSuppressedDrops" { return "Cipurpfifsuppresseddrops" }
    if yname == "cipUrpfIfDropRate" { return "Cipurpfifdroprate" }
    if yname == "cipUrpfIfDiscontinuityTime" { return "Cipurpfifdiscontinuitytime" }
    if yname == "cipUrpfIfDropRateNotifyEnable" { return "Cipurpfifdropratenotifyenable" }
    if yname == "cipUrpfIfNotifyDropRateThreshold" { return "Cipurpfifnotifydropratethreshold" }
    if yname == "cipUrpfIfNotifyDrHoldDownReset" { return "Cipurpfifnotifydrholddownreset" }
    if yname == "cipUrpfIfCheckStrict" { return "Cipurpfifcheckstrict" }
    if yname == "cipUrpfIfWhichRouteTableID" { return "Cipurpfifwhichroutetableid" }
    if yname == "cipUrpfIfVrfName" { return "Cipurpfifvrfname" }
    return ""
}

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetSegmentPath() string {
    return "cipUrpfIfMonEntry" + "[ifIndex='" + fmt.Sprintf("%v", cipurpfifmonentry.Ifindex) + "']" + "[cipUrpfIfIpVersion='" + fmt.Sprintf("%v", cipurpfifmonentry.Cipurpfifipversion) + "']"
}

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cipurpfifmonentry.Ifindex
    leafs["cipUrpfIfIpVersion"] = cipurpfifmonentry.Cipurpfifipversion
    leafs["cipUrpfIfDrops"] = cipurpfifmonentry.Cipurpfifdrops
    leafs["cipUrpfIfSuppressedDrops"] = cipurpfifmonentry.Cipurpfifsuppresseddrops
    leafs["cipUrpfIfDropRate"] = cipurpfifmonentry.Cipurpfifdroprate
    leafs["cipUrpfIfDiscontinuityTime"] = cipurpfifmonentry.Cipurpfifdiscontinuitytime
    leafs["cipUrpfIfDropRateNotifyEnable"] = cipurpfifmonentry.Cipurpfifdropratenotifyenable
    leafs["cipUrpfIfNotifyDropRateThreshold"] = cipurpfifmonentry.Cipurpfifnotifydropratethreshold
    leafs["cipUrpfIfNotifyDrHoldDownReset"] = cipurpfifmonentry.Cipurpfifnotifydrholddownreset
    leafs["cipUrpfIfCheckStrict"] = cipurpfifmonentry.Cipurpfifcheckstrict
    leafs["cipUrpfIfWhichRouteTableID"] = cipurpfifmonentry.Cipurpfifwhichroutetableid
    leafs["cipUrpfIfVrfName"] = cipurpfifmonentry.Cipurpfifvrfname
    return leafs
}

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetYangName() string { return "cipUrpfIfMonEntry" }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) SetParent(parent types.Entity) { cipurpfifmonentry.parent = parent }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetParent() types.Entity { return cipurpfifmonentry.parent }

func (cipurpfifmonentry *CISCOIPURPFMIB_Cipurpfifmontable_Cipurpfifmonentry) GetParentYangName() string { return "cipUrpfIfMonTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry exists for a VRF and interface if and only if the VRF associated
    // with the interface is configured  to perform IP URPF checking using the
    // routing  table for the VRF. The type is slice of
    // CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry.
    Cipurpfvrfifentry []CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry
}

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetFilter() yfilter.YFilter { return cipurpfvrfiftable.YFilter }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) SetFilter(yf yfilter.YFilter) { cipurpfvrfiftable.YFilter = yf }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetGoName(yname string) string {
    if yname == "cipUrpfVrfIfEntry" { return "Cipurpfvrfifentry" }
    return ""
}

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetSegmentPath() string {
    return "cipUrpfVrfIfTable"
}

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipUrpfVrfIfEntry" {
        for _, c := range cipurpfvrfiftable.Cipurpfvrfifentry {
            if cipurpfvrfiftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry{}
        cipurpfvrfiftable.Cipurpfvrfifentry = append(cipurpfvrfiftable.Cipurpfvrfifentry, child)
        return &cipurpfvrfiftable.Cipurpfvrfifentry[len(cipurpfvrfiftable.Cipurpfvrfifentry)-1]
    }
    return nil
}

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipurpfvrfiftable.Cipurpfvrfifentry {
        children[cipurpfvrfiftable.Cipurpfvrfifentry[i].GetSegmentPath()] = &cipurpfvrfiftable.Cipurpfvrfifentry[i]
    }
    return children
}

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetYangName() string { return "cipUrpfVrfIfTable" }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) SetParent(parent types.Entity) { cipurpfvrfiftable.parent = parent }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetParent() types.Entity { return cipurpfvrfiftable.parent }

func (cipurpfvrfiftable *CISCOIPURPFMIB_Cipurpfvrfiftable) GetParentYangName() string { return "CISCO-IP-URPF-MIB" }

// CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry
// An entry exists for a VRF and interface if and only
// if the VRF associated with the interface is configured 
// to perform IP URPF checking using the routing 
// table for the VRF.
type CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry struct {
    parent types.Entity
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

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetFilter() yfilter.YFilter { return cipurpfvrfifentry.YFilter }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) SetFilter(yf yfilter.YFilter) { cipurpfvrfifentry.YFilter = yf }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetGoName(yname string) string {
    if yname == "cipUrpfVrfName" { return "Cipurpfvrfname" }
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cipUrpfVrfIfDrops" { return "Cipurpfvrfifdrops" }
    if yname == "cipUrpfVrfIfDiscontinuityTime" { return "Cipurpfvrfifdiscontinuitytime" }
    return ""
}

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetSegmentPath() string {
    return "cipUrpfVrfIfEntry" + "[cipUrpfVrfName='" + fmt.Sprintf("%v", cipurpfvrfifentry.Cipurpfvrfname) + "']" + "[ifIndex='" + fmt.Sprintf("%v", cipurpfvrfifentry.Ifindex) + "']"
}

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipUrpfVrfName"] = cipurpfvrfifentry.Cipurpfvrfname
    leafs["ifIndex"] = cipurpfvrfifentry.Ifindex
    leafs["cipUrpfVrfIfDrops"] = cipurpfvrfifentry.Cipurpfvrfifdrops
    leafs["cipUrpfVrfIfDiscontinuityTime"] = cipurpfvrfifentry.Cipurpfvrfifdiscontinuitytime
    return leafs
}

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetYangName() string { return "cipUrpfVrfIfEntry" }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) SetParent(parent types.Entity) { cipurpfvrfifentry.parent = parent }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetParent() types.Entity { return cipurpfvrfifentry.parent }

func (cipurpfvrfifentry *CISCOIPURPFMIB_Cipurpfvrfiftable_Cipurpfvrfifentry) GetParentYangName() string { return "cipUrpfVrfIfTable" }

// CISCOIPURPFMIB_Cipurpfvrftable
// This table enables indexing URPF drop statistics
// by Virtual Routing and Forwarding instances.
type CISCOIPURPFMIB_Cipurpfvrftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry exists for a VRF if and only if the VRF is associated with an
    // interface that is configured to perform IP URPF checking using the routing
    // table  for that VRF. The type is slice of
    // CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry.
    Cipurpfvrfentry []CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry
}

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetFilter() yfilter.YFilter { return cipurpfvrftable.YFilter }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) SetFilter(yf yfilter.YFilter) { cipurpfvrftable.YFilter = yf }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetGoName(yname string) string {
    if yname == "cipUrpfVrfEntry" { return "Cipurpfvrfentry" }
    return ""
}

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetSegmentPath() string {
    return "cipUrpfVrfTable"
}

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipUrpfVrfEntry" {
        for _, c := range cipurpfvrftable.Cipurpfvrfentry {
            if cipurpfvrftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry{}
        cipurpfvrftable.Cipurpfvrfentry = append(cipurpfvrftable.Cipurpfvrfentry, child)
        return &cipurpfvrftable.Cipurpfvrfentry[len(cipurpfvrftable.Cipurpfvrfentry)-1]
    }
    return nil
}

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipurpfvrftable.Cipurpfvrfentry {
        children[cipurpfvrftable.Cipurpfvrfentry[i].GetSegmentPath()] = &cipurpfvrftable.Cipurpfvrfentry[i]
    }
    return children
}

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetYangName() string { return "cipUrpfVrfTable" }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) SetParent(parent types.Entity) { cipurpfvrftable.parent = parent }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetParent() types.Entity { return cipurpfvrftable.parent }

func (cipurpfvrftable *CISCOIPURPFMIB_Cipurpfvrftable) GetParentYangName() string { return "CISCO-IP-URPF-MIB" }

// CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry
// An entry exists for a VRF if and only if the VRF
// is associated with an interface that is configured
// to perform IP URPF checking using the routing table 
// for that VRF.
type CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This field is used to specify the VRF Table name.
    // The type is string with length: 0..32.
    Cipurpfvrfname interface{}
}

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetFilter() yfilter.YFilter { return cipurpfvrfentry.YFilter }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) SetFilter(yf yfilter.YFilter) { cipurpfvrfentry.YFilter = yf }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetGoName(yname string) string {
    if yname == "cipUrpfVrfName" { return "Cipurpfvrfname" }
    return ""
}

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetSegmentPath() string {
    return "cipUrpfVrfEntry" + "[cipUrpfVrfName='" + fmt.Sprintf("%v", cipurpfvrfentry.Cipurpfvrfname) + "']"
}

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipUrpfVrfName"] = cipurpfvrfentry.Cipurpfvrfname
    return leafs
}

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetYangName() string { return "cipUrpfVrfEntry" }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) SetParent(parent types.Entity) { cipurpfvrfentry.parent = parent }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetParent() types.Entity { return cipurpfvrfentry.parent }

func (cipurpfvrfentry *CISCOIPURPFMIB_Cipurpfvrftable_Cipurpfvrfentry) GetParentYangName() string { return "cipUrpfVrfTable" }

