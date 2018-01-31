// The MIB module provides a means to monitor and configure
// the Cisco IOS proprietary Hot Standby Router Protocol
// (HSRP). Cisco HSRP protocol is defined in RFC2281.
// 
// Terminology:
// 
// HSRP is a protocol used amoung a group of routers for the
// purpose of selecting an 'active router' and a 'standby
// router'.
// 
// An 'active router' is the router of choice for routing
// packets.
// 
// A 'standby router' is a router that takes over the routing
// duties when an active router fails, or when preset
// conditions have been met.
// 
// An 'HSRP group' or a 'standby group' is a set of routers
// which communicate using HSRP. An HSRP group has a group MAC
// address and a group Virtual IP address. These are the
// designated addresses. The active router assumes (i.e.
// inherits) these group addresses.
// 
// 'Hello' messages are sent to indicate that a router is
// running and is capable of becoming the active or standby
// router.
// 
// 'Hellotime' is the interval between successive HSRP Hello
// messages from a given router.
// 
// 'Holdtime' is the interval between the receipt of a Hello
// message and the presumption that the sending router has
// failed.
package cisco_hsrp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_hsrp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-HSRP-MIB CISCO-HSRP-MIB}", reflect.TypeOf(CISCOHSRPMIB{}))
    ydk.RegisterEntity("CISCO-HSRP-MIB:CISCO-HSRP-MIB", reflect.TypeOf(CISCOHSRPMIB{}))
}

// HsrpState represents HSRP group entry.
type HsrpState string

const (
    HsrpState_initial HsrpState = "initial"

    HsrpState_learn HsrpState = "learn"

    HsrpState_listen HsrpState = "listen"

    HsrpState_speak HsrpState = "speak"

    HsrpState_standby HsrpState = "standby"

    HsrpState_active HsrpState = "active"
)

// CISCOHSRPMIB
type CISCOHSRPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Chsrpglobalconfig CISCOHSRPMIB_Chsrpglobalconfig

    // A table containing information on each HSRP group for each interface.
    Chsrpgrptable CISCOHSRPMIB_Chsrpgrptable
}

func (cISCOHSRPMIB *CISCOHSRPMIB) GetFilter() yfilter.YFilter { return cISCOHSRPMIB.YFilter }

func (cISCOHSRPMIB *CISCOHSRPMIB) SetFilter(yf yfilter.YFilter) { cISCOHSRPMIB.YFilter = yf }

func (cISCOHSRPMIB *CISCOHSRPMIB) GetGoName(yname string) string {
    if yname == "cHsrpGlobalConfig" { return "Chsrpglobalconfig" }
    if yname == "cHsrpGrpTable" { return "Chsrpgrptable" }
    return ""
}

func (cISCOHSRPMIB *CISCOHSRPMIB) GetSegmentPath() string {
    return "CISCO-HSRP-MIB:CISCO-HSRP-MIB"
}

func (cISCOHSRPMIB *CISCOHSRPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cHsrpGlobalConfig" {
        return &cISCOHSRPMIB.Chsrpglobalconfig
    }
    if childYangName == "cHsrpGrpTable" {
        return &cISCOHSRPMIB.Chsrpgrptable
    }
    return nil
}

func (cISCOHSRPMIB *CISCOHSRPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cHsrpGlobalConfig"] = &cISCOHSRPMIB.Chsrpglobalconfig
    children["cHsrpGrpTable"] = &cISCOHSRPMIB.Chsrpgrptable
    return children
}

func (cISCOHSRPMIB *CISCOHSRPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOHSRPMIB *CISCOHSRPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOHSRPMIB *CISCOHSRPMIB) GetYangName() string { return "CISCO-HSRP-MIB" }

func (cISCOHSRPMIB *CISCOHSRPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOHSRPMIB *CISCOHSRPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOHSRPMIB *CISCOHSRPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOHSRPMIB *CISCOHSRPMIB) SetParent(parent types.Entity) { cISCOHSRPMIB.parent = parent }

func (cISCOHSRPMIB *CISCOHSRPMIB) GetParent() types.Entity { return cISCOHSRPMIB.parent }

func (cISCOHSRPMIB *CISCOHSRPMIB) GetParentYangName() string { return "CISCO-HSRP-MIB" }

// CISCOHSRPMIB_Chsrpglobalconfig
type CISCOHSRPMIB_Chsrpglobalconfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The amount of time in minutes a row in cHsrpGrpTable can remain in a state
    // other than active before being timed out. The type is interface{} with
    // range: 1..60. Units are minutes.
    Chsrpconfigtimeout interface{}
}

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetFilter() yfilter.YFilter { return chsrpglobalconfig.YFilter }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) SetFilter(yf yfilter.YFilter) { chsrpglobalconfig.YFilter = yf }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetGoName(yname string) string {
    if yname == "cHsrpConfigTimeout" { return "Chsrpconfigtimeout" }
    return ""
}

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetSegmentPath() string {
    return "cHsrpGlobalConfig"
}

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cHsrpConfigTimeout"] = chsrpglobalconfig.Chsrpconfigtimeout
    return leafs
}

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetYangName() string { return "cHsrpGlobalConfig" }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) SetParent(parent types.Entity) { chsrpglobalconfig.parent = parent }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetParent() types.Entity { return chsrpglobalconfig.parent }

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetParentYangName() string { return "CISCO-HSRP-MIB" }

// CISCOHSRPMIB_Chsrpgrptable
// A table containing information on each HSRP group
// for each interface.
type CISCOHSRPMIB_Chsrpgrptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about an HSRP group. Management applications use
    // cHsrpGrpRowStatus to control entry modification, creation and deletion. 
    // Setting cHsrpGrpRowStatus to 'active' causes the router to communicate
    // using HSRP.  The value of cHsrpGrpRowStatus may be set to 'destroy' at any
    // time.  Entries may not be created via SNMP without explicitly  setting
    // cHsrpGrpRowStatus to either 'createAndGo' or  'createAndWait'.  Entries can
    // be created and modified via the management  protocol or by the device's
    // local management interface.  A management application wishing to create an
    // entry should choose the ifIndex of the interface which is to be added as
    // part of an HSRP group. Also, a cHsrpGrpNumber should be chosen. A group
    // number is unique only amongst the groups  on a particular interface. The
    // value of the group number appears in packets which are transmitted and
    // received on a  LAN segment to which the router is connected. The
    // application must select the group number as explained in the description
    // for cHsrpGrpNumber.  If the row is not active, and a local management
    // interface command modifies that row, the row may transition to active
    // state.  A row which is not in active state will timeout after a
    // configurable period (five minutes by default). This timeout  period can be
    // changed by setting cHsrpConfigTimeout. The type is slice of
    // CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry.
    Chsrpgrpentry []CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry
}

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetFilter() yfilter.YFilter { return chsrpgrptable.YFilter }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) SetFilter(yf yfilter.YFilter) { chsrpgrptable.YFilter = yf }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetGoName(yname string) string {
    if yname == "cHsrpGrpEntry" { return "Chsrpgrpentry" }
    return ""
}

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetSegmentPath() string {
    return "cHsrpGrpTable"
}

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cHsrpGrpEntry" {
        for _, c := range chsrpgrptable.Chsrpgrpentry {
            if chsrpgrptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry{}
        chsrpgrptable.Chsrpgrpentry = append(chsrpgrptable.Chsrpgrpentry, child)
        return &chsrpgrptable.Chsrpgrpentry[len(chsrpgrptable.Chsrpgrpentry)-1]
    }
    return nil
}

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range chsrpgrptable.Chsrpgrpentry {
        children[chsrpgrptable.Chsrpgrpentry[i].GetSegmentPath()] = &chsrpgrptable.Chsrpgrpentry[i]
    }
    return children
}

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetYangName() string { return "cHsrpGrpTable" }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) SetParent(parent types.Entity) { chsrpgrptable.parent = parent }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetParent() types.Entity { return chsrpgrptable.parent }

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetParentYangName() string { return "CISCO-HSRP-MIB" }

// CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry
// Information about an HSRP group. Management applications
// use cHsrpGrpRowStatus to control entry modification,
// creation and deletion.
// 
// Setting cHsrpGrpRowStatus to 'active' causes the router to
// communicate using HSRP.
// 
// The value of cHsrpGrpRowStatus may be set to 'destroy' at
// any time.
// 
// Entries may not be created via SNMP without explicitly 
// setting cHsrpGrpRowStatus to either 'createAndGo' or 
// 'createAndWait'.
// 
// Entries can be created and modified via the management 
// protocol or by the device's local management interface.
// 
// A management application wishing to create an entry should
// choose the ifIndex of the interface which is to be added
// as part of an HSRP group. Also, a cHsrpGrpNumber should
// be chosen. A group number is unique only amongst the groups 
// on a particular interface. The value of the group number
// appears in packets which are transmitted and received on a 
// LAN segment to which the router is connected. The application
// must select the group number as explained in the description
// for cHsrpGrpNumber.
// 
// If the row is not active, and a local management interface
// command modifies that row, the row may transition to active
// state.
// 
// A row which is not in active state will timeout after a
// configurable period (five minutes by default). This timeout 
// period can be changed by setting cHsrpConfigTimeout.
type CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. This object along with the ifIndex of a particular
    // interface uniquely identifies an HSRP group.  Group numbers 0,1 and 2 are
    // the only valid group numbers for TokenRing interfaces. For other media
    // types, numbers range from 0 to 255. Each interface has its own set of group
    // numbers. There's no relationship between the groups configured on different
    // interfaces. Using a group number on one interface doesn't preclude using
    // the same group number on a different interface. For example, there can be a
    // group 1 on an Ethernet and a group 1 on Token Ring. More details can be
    // found from RFC 2281. The type is interface{} with range: 0..255.
    Chsrpgrpnumber interface{}

    // This is an unencrypted authentication string which is carried in all HSRP
    // messages. An authentication string mismatch prevents a router interface
    // from learning the designated IP address or HSRP timer values from other
    // HSRP-enabled routers with the same group number.  The function of this
    // object is not to supply any sort of security-like authentication but rather
    // to confirm that what's happening is what's intended. In other words, this
    // is meant for sanity checking only. The type is string with length: 0..8.
    Chsrpgrpauth interface{}

    // The cHsrpGrpPriority helps to select the active and the standby routers.
    // The router with the highest priority is selected as the active router. In
    // the priority range of 0 to 255, 0 is the lowest priority and 255 is the
    // highest priority.  If two (or more) routers in a group have the same
    // priority, the one with the highest ip address of the interface is the
    // active router. When the active router fails to send a Hello message within
    // a configurable period of time, the standby router with the highest priority
    // becomes the active router.  A router with highest priority will only
    // attempt to overthrow a lower priority active router if it is configured to
    // preempt.  But, if there is more than one router which is not active, the
    // highest priority non-active router becomes the standby router. The type is
    // interface{} with range: 0..255.
    Chsrpgrppriority interface{}

    // This object, if TRUE, indicates that the current router should attempt to
    // overthrow a lower priority active router and attempt to become the active
    // router. If this object is FALSE, the router will become the active router
    // only if there is no such router (or if an active router fails). The type is
    // bool.
    Chsrpgrppreempt interface{}

    // This delay is the time difference between a router power up and the time it
    // can actually start preempting the currently active router.  When a router
    // first comes up, it doesn't have a complete routing table. If it's
    // configured to preempt, then it will become the Active router, but it will
    // not be able to provide adequate routing services. The solution to this is
    // to allow for a configurable delay before the router actually preempts the
    // currently active router. The type is interface{} with range: 0..3600. Units
    // are seconds.
    Chsrpgrppreemptdelay interface{}

    // HSRP routers learn a group's Hellotime or Holdtime from hello messages. 
    // The Hellotime is used to determine the frequency of generating hello
    // messages when this router becomes the active or standby router. The
    // Holdtime is the interval between the receipt of a Hello message and the
    // presumption that the sending router has failed.  If this object is TRUE,
    // the cHsrpGrpConfiguredHelloTime and cHsrpGrpConfiguredHoldTime will be
    // used. If it is FALSE, the Hellotime and Holdtime values are learned. The
    // type is bool.
    Chsrpgrpuseconfiguredtimers interface{}

    // If cHsrpGrpUseConfiguredTimers is true, cHsrpGrpConfiguredHelloTime is used
    // when this router is an active router. Otherwise, the Hellotime learned from
    // the current active router is used. All routers on a particular LAN segment
    // must use the same Hellotime. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    Chsrpgrpconfiguredhellotime interface{}

    // If cHsrpGrpUseConfiguredTimers is true, cHsrpGrpConfiguredHoldTime is used
    // when this router is an active router. Otherwise, the Holdtime learned from
    // the current active router is used. All routers on a particular LAN segment
    // should use the same Holdtime. Also, the Holdtime should be at least three
    // times the value of the Hellotime and must be greater than the Hellotime.
    // The type is interface{} with range: 0..4294967295. Units are milliseconds.
    Chsrpgrpconfiguredholdtime interface{}

    // If the Hellotime is not configured on a router, it can be learned from the
    // Hello messages from active router, provided the Hello message is
    // authenticated. If the Hellotime is not learned from a Hello message from
    // the active router and it is not manually configured, a default value of 3
    // seconds is recommended. The type is interface{} with range: 0..4294967295.
    // Units are milliseconds.
    Chsrpgrplearnedhellotime interface{}

    // If the Holdtime is not configured on a router, it can be learned from the
    // Hello message from the active router. Holdtime should be learned only if
    // the Hello message is authenticated. If the Holdtime is not learned and it
    // is not manually configured, a default value of 10 seconds is  recommended.
    // The type is interface{} with range: 0..4294967295. Units are milliseconds.
    Chsrpgrplearnedholdtime interface{}

    // This is the primary virtual IP address used by this group.  If this address
    // is configured (i.e a non zero ip address), this value is used. Otherwise,
    // the agent will attempt to discover the virtual address through a discovery
    // process (which scans the hello messages). The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Chsrpgrpvirtualipaddr interface{}

    // If this object is TRUE, cHsrpGrpVirtualIpAddr was a configured one.
    // Otherwise, it indicates that  cHsrpGrpVirtualIpAddr was a learned one. The
    // type is bool.
    Chsrpgrpuseconfigvirtualipaddr interface{}

    // Ip Address of the currently active router for this group. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Chsrpgrpactiverouter interface{}

    // Ip Address of the currently standby router for this group. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Chsrpgrpstandbyrouter interface{}

    // The current HSRP state of this group on this interface. The type is
    // HsrpState.
    Chsrpgrpstandbystate interface{}

    // Mac Addresses used are as specified in RFC 2281. For ethernet and fddi
    // interfaces, a MAC address will be in the range 00:00:0c:07:ac:00 through
    // 00:00:0c:07:ac:ff. The last octet is the hexadecimal equivalent of
    // cHsrpGrpNumber (0-255).  Some Ethernet and FDDI interfaces allow a unicast
    // MAC address for each HSRP group. Certain Ethernet chipsets(LANCE Ethernet,
    // VGANYLAN and QUICC Ethernet) only support a single Unicast Mac Address. In
    // this case, only one HSRP group is allowed.  For TokenRing interfaces, the
    // following three MAC  addresses are permitted (functional addresses):       
    // C0:00:00:01:00:00              C0:00:00:02:00:00             
    // C0:00:00:04:00:00. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Chsrpgrpvirtualmacaddr interface{}

    // The control that allows modification, creation, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cHsrpGrpEntry. The type is
    // RowStatus.
    Chsrpgrpentryrowstatus interface{}

    // This object specifies the disable HSRP IPv4 virtual IP address. The type is
    // bool.
    Chsrpgrpipnone interface{}
}

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetFilter() yfilter.YFilter { return chsrpgrpentry.YFilter }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) SetFilter(yf yfilter.YFilter) { chsrpgrpentry.YFilter = yf }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cHsrpGrpNumber" { return "Chsrpgrpnumber" }
    if yname == "cHsrpGrpAuth" { return "Chsrpgrpauth" }
    if yname == "cHsrpGrpPriority" { return "Chsrpgrppriority" }
    if yname == "cHsrpGrpPreempt" { return "Chsrpgrppreempt" }
    if yname == "cHsrpGrpPreemptDelay" { return "Chsrpgrppreemptdelay" }
    if yname == "cHsrpGrpUseConfiguredTimers" { return "Chsrpgrpuseconfiguredtimers" }
    if yname == "cHsrpGrpConfiguredHelloTime" { return "Chsrpgrpconfiguredhellotime" }
    if yname == "cHsrpGrpConfiguredHoldTime" { return "Chsrpgrpconfiguredholdtime" }
    if yname == "cHsrpGrpLearnedHelloTime" { return "Chsrpgrplearnedhellotime" }
    if yname == "cHsrpGrpLearnedHoldTime" { return "Chsrpgrplearnedholdtime" }
    if yname == "cHsrpGrpVirtualIpAddr" { return "Chsrpgrpvirtualipaddr" }
    if yname == "cHsrpGrpUseConfigVirtualIpAddr" { return "Chsrpgrpuseconfigvirtualipaddr" }
    if yname == "cHsrpGrpActiveRouter" { return "Chsrpgrpactiverouter" }
    if yname == "cHsrpGrpStandbyRouter" { return "Chsrpgrpstandbyrouter" }
    if yname == "cHsrpGrpStandbyState" { return "Chsrpgrpstandbystate" }
    if yname == "cHsrpGrpVirtualMacAddr" { return "Chsrpgrpvirtualmacaddr" }
    if yname == "cHsrpGrpEntryRowStatus" { return "Chsrpgrpentryrowstatus" }
    if yname == "cHsrpGrpIpNone" { return "Chsrpgrpipnone" }
    return ""
}

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetSegmentPath() string {
    return "cHsrpGrpEntry" + "[ifIndex='" + fmt.Sprintf("%v", chsrpgrpentry.Ifindex) + "']" + "[cHsrpGrpNumber='" + fmt.Sprintf("%v", chsrpgrpentry.Chsrpgrpnumber) + "']"
}

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = chsrpgrpentry.Ifindex
    leafs["cHsrpGrpNumber"] = chsrpgrpentry.Chsrpgrpnumber
    leafs["cHsrpGrpAuth"] = chsrpgrpentry.Chsrpgrpauth
    leafs["cHsrpGrpPriority"] = chsrpgrpentry.Chsrpgrppriority
    leafs["cHsrpGrpPreempt"] = chsrpgrpentry.Chsrpgrppreempt
    leafs["cHsrpGrpPreemptDelay"] = chsrpgrpentry.Chsrpgrppreemptdelay
    leafs["cHsrpGrpUseConfiguredTimers"] = chsrpgrpentry.Chsrpgrpuseconfiguredtimers
    leafs["cHsrpGrpConfiguredHelloTime"] = chsrpgrpentry.Chsrpgrpconfiguredhellotime
    leafs["cHsrpGrpConfiguredHoldTime"] = chsrpgrpentry.Chsrpgrpconfiguredholdtime
    leafs["cHsrpGrpLearnedHelloTime"] = chsrpgrpentry.Chsrpgrplearnedhellotime
    leafs["cHsrpGrpLearnedHoldTime"] = chsrpgrpentry.Chsrpgrplearnedholdtime
    leafs["cHsrpGrpVirtualIpAddr"] = chsrpgrpentry.Chsrpgrpvirtualipaddr
    leafs["cHsrpGrpUseConfigVirtualIpAddr"] = chsrpgrpentry.Chsrpgrpuseconfigvirtualipaddr
    leafs["cHsrpGrpActiveRouter"] = chsrpgrpentry.Chsrpgrpactiverouter
    leafs["cHsrpGrpStandbyRouter"] = chsrpgrpentry.Chsrpgrpstandbyrouter
    leafs["cHsrpGrpStandbyState"] = chsrpgrpentry.Chsrpgrpstandbystate
    leafs["cHsrpGrpVirtualMacAddr"] = chsrpgrpentry.Chsrpgrpvirtualmacaddr
    leafs["cHsrpGrpEntryRowStatus"] = chsrpgrpentry.Chsrpgrpentryrowstatus
    leafs["cHsrpGrpIpNone"] = chsrpgrpentry.Chsrpgrpipnone
    return leafs
}

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetBundleName() string { return "cisco_ios_xe" }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetYangName() string { return "cHsrpGrpEntry" }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) SetParent(parent types.Entity) { chsrpgrpentry.parent = parent }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetParent() types.Entity { return chsrpgrpentry.parent }

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetParentYangName() string { return "cHsrpGrpTable" }

