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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Chsrpglobalconfig CISCOHSRPMIB_Chsrpglobalconfig

    // A table containing information on each HSRP group for each interface.
    Chsrpgrptable CISCOHSRPMIB_Chsrpgrptable
}

func (cISCOHSRPMIB *CISCOHSRPMIB) GetEntityData() *types.CommonEntityData {
    cISCOHSRPMIB.EntityData.YFilter = cISCOHSRPMIB.YFilter
    cISCOHSRPMIB.EntityData.YangName = "CISCO-HSRP-MIB"
    cISCOHSRPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOHSRPMIB.EntityData.ParentYangName = "CISCO-HSRP-MIB"
    cISCOHSRPMIB.EntityData.SegmentPath = "CISCO-HSRP-MIB:CISCO-HSRP-MIB"
    cISCOHSRPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOHSRPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOHSRPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOHSRPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOHSRPMIB.EntityData.Children["cHsrpGlobalConfig"] = types.YChild{"Chsrpglobalconfig", &cISCOHSRPMIB.Chsrpglobalconfig}
    cISCOHSRPMIB.EntityData.Children["cHsrpGrpTable"] = types.YChild{"Chsrpgrptable", &cISCOHSRPMIB.Chsrpgrptable}
    cISCOHSRPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOHSRPMIB.EntityData)
}

// CISCOHSRPMIB_Chsrpglobalconfig
type CISCOHSRPMIB_Chsrpglobalconfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The amount of time in minutes a row in cHsrpGrpTable can remain in a state
    // other than active before being timed out. The type is interface{} with
    // range: 1..60. Units are minutes.
    Chsrpconfigtimeout interface{}
}

func (chsrpglobalconfig *CISCOHSRPMIB_Chsrpglobalconfig) GetEntityData() *types.CommonEntityData {
    chsrpglobalconfig.EntityData.YFilter = chsrpglobalconfig.YFilter
    chsrpglobalconfig.EntityData.YangName = "cHsrpGlobalConfig"
    chsrpglobalconfig.EntityData.BundleName = "cisco_ios_xe"
    chsrpglobalconfig.EntityData.ParentYangName = "CISCO-HSRP-MIB"
    chsrpglobalconfig.EntityData.SegmentPath = "cHsrpGlobalConfig"
    chsrpglobalconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    chsrpglobalconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    chsrpglobalconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    chsrpglobalconfig.EntityData.Children = make(map[string]types.YChild)
    chsrpglobalconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    chsrpglobalconfig.EntityData.Leafs["cHsrpConfigTimeout"] = types.YLeaf{"Chsrpconfigtimeout", chsrpglobalconfig.Chsrpconfigtimeout}
    return &(chsrpglobalconfig.EntityData)
}

// CISCOHSRPMIB_Chsrpgrptable
// A table containing information on each HSRP group
// for each interface.
type CISCOHSRPMIB_Chsrpgrptable struct {
    EntityData types.CommonEntityData
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

func (chsrpgrptable *CISCOHSRPMIB_Chsrpgrptable) GetEntityData() *types.CommonEntityData {
    chsrpgrptable.EntityData.YFilter = chsrpgrptable.YFilter
    chsrpgrptable.EntityData.YangName = "cHsrpGrpTable"
    chsrpgrptable.EntityData.BundleName = "cisco_ios_xe"
    chsrpgrptable.EntityData.ParentYangName = "CISCO-HSRP-MIB"
    chsrpgrptable.EntityData.SegmentPath = "cHsrpGrpTable"
    chsrpgrptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    chsrpgrptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    chsrpgrptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    chsrpgrptable.EntityData.Children = make(map[string]types.YChild)
    chsrpgrptable.EntityData.Children["cHsrpGrpEntry"] = types.YChild{"Chsrpgrpentry", nil}
    for i := range chsrpgrptable.Chsrpgrpentry {
        chsrpgrptable.EntityData.Children[types.GetSegmentPath(&chsrpgrptable.Chsrpgrpentry[i])] = types.YChild{"Chsrpgrpentry", &chsrpgrptable.Chsrpgrpentry[i]}
    }
    chsrpgrptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(chsrpgrptable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Chsrpgrpvirtualipaddr interface{}

    // If this object is TRUE, cHsrpGrpVirtualIpAddr was a configured one.
    // Otherwise, it indicates that  cHsrpGrpVirtualIpAddr was a learned one. The
    // type is bool.
    Chsrpgrpuseconfigvirtualipaddr interface{}

    // Ip Address of the currently active router for this group. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Chsrpgrpactiverouter interface{}

    // Ip Address of the currently standby router for this group. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Chsrpgrpvirtualmacaddr interface{}

    // The control that allows modification, creation, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cHsrpGrpEntry. The type is
    // RowStatus.
    Chsrpgrpentryrowstatus interface{}

    // This object specifies the disable HSRP IPv4 virtual IP address. The type is
    // bool.
    Chsrpgrpipnone interface{}
}

func (chsrpgrpentry *CISCOHSRPMIB_Chsrpgrptable_Chsrpgrpentry) GetEntityData() *types.CommonEntityData {
    chsrpgrpentry.EntityData.YFilter = chsrpgrpentry.YFilter
    chsrpgrpentry.EntityData.YangName = "cHsrpGrpEntry"
    chsrpgrpentry.EntityData.BundleName = "cisco_ios_xe"
    chsrpgrpentry.EntityData.ParentYangName = "cHsrpGrpTable"
    chsrpgrpentry.EntityData.SegmentPath = "cHsrpGrpEntry" + "[ifIndex='" + fmt.Sprintf("%v", chsrpgrpentry.Ifindex) + "']" + "[cHsrpGrpNumber='" + fmt.Sprintf("%v", chsrpgrpentry.Chsrpgrpnumber) + "']"
    chsrpgrpentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    chsrpgrpentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    chsrpgrpentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    chsrpgrpentry.EntityData.Children = make(map[string]types.YChild)
    chsrpgrpentry.EntityData.Leafs = make(map[string]types.YLeaf)
    chsrpgrpentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", chsrpgrpentry.Ifindex}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpNumber"] = types.YLeaf{"Chsrpgrpnumber", chsrpgrpentry.Chsrpgrpnumber}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpAuth"] = types.YLeaf{"Chsrpgrpauth", chsrpgrpentry.Chsrpgrpauth}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpPriority"] = types.YLeaf{"Chsrpgrppriority", chsrpgrpentry.Chsrpgrppriority}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpPreempt"] = types.YLeaf{"Chsrpgrppreempt", chsrpgrpentry.Chsrpgrppreempt}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpPreemptDelay"] = types.YLeaf{"Chsrpgrppreemptdelay", chsrpgrpentry.Chsrpgrppreemptdelay}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpUseConfiguredTimers"] = types.YLeaf{"Chsrpgrpuseconfiguredtimers", chsrpgrpentry.Chsrpgrpuseconfiguredtimers}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpConfiguredHelloTime"] = types.YLeaf{"Chsrpgrpconfiguredhellotime", chsrpgrpentry.Chsrpgrpconfiguredhellotime}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpConfiguredHoldTime"] = types.YLeaf{"Chsrpgrpconfiguredholdtime", chsrpgrpentry.Chsrpgrpconfiguredholdtime}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpLearnedHelloTime"] = types.YLeaf{"Chsrpgrplearnedhellotime", chsrpgrpentry.Chsrpgrplearnedhellotime}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpLearnedHoldTime"] = types.YLeaf{"Chsrpgrplearnedholdtime", chsrpgrpentry.Chsrpgrplearnedholdtime}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpVirtualIpAddr"] = types.YLeaf{"Chsrpgrpvirtualipaddr", chsrpgrpentry.Chsrpgrpvirtualipaddr}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpUseConfigVirtualIpAddr"] = types.YLeaf{"Chsrpgrpuseconfigvirtualipaddr", chsrpgrpentry.Chsrpgrpuseconfigvirtualipaddr}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpActiveRouter"] = types.YLeaf{"Chsrpgrpactiverouter", chsrpgrpentry.Chsrpgrpactiverouter}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpStandbyRouter"] = types.YLeaf{"Chsrpgrpstandbyrouter", chsrpgrpentry.Chsrpgrpstandbyrouter}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpStandbyState"] = types.YLeaf{"Chsrpgrpstandbystate", chsrpgrpentry.Chsrpgrpstandbystate}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpVirtualMacAddr"] = types.YLeaf{"Chsrpgrpvirtualmacaddr", chsrpgrpentry.Chsrpgrpvirtualmacaddr}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpEntryRowStatus"] = types.YLeaf{"Chsrpgrpentryrowstatus", chsrpgrpentry.Chsrpgrpentryrowstatus}
    chsrpgrpentry.EntityData.Leafs["cHsrpGrpIpNone"] = types.YLeaf{"Chsrpgrpipnone", chsrpgrpentry.Chsrpgrpipnone}
    return &(chsrpgrpentry.EntityData)
}

