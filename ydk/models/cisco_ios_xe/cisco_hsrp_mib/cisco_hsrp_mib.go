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

    
    CHsrpGlobalConfig CISCOHSRPMIB_CHsrpGlobalConfig

    // A table containing information on each HSRP group for each interface.
    CHsrpGrpTable CISCOHSRPMIB_CHsrpGrpTable
}

func (cISCOHSRPMIB *CISCOHSRPMIB) GetEntityData() *types.CommonEntityData {
    cISCOHSRPMIB.EntityData.YFilter = cISCOHSRPMIB.YFilter
    cISCOHSRPMIB.EntityData.YangName = "CISCO-HSRP-MIB"
    cISCOHSRPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOHSRPMIB.EntityData.ParentYangName = "CISCO-HSRP-MIB"
    cISCOHSRPMIB.EntityData.SegmentPath = "CISCO-HSRP-MIB:CISCO-HSRP-MIB"
    cISCOHSRPMIB.EntityData.AbsolutePath = cISCOHSRPMIB.EntityData.SegmentPath
    cISCOHSRPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOHSRPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOHSRPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOHSRPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOHSRPMIB.EntityData.Children.Append("cHsrpGlobalConfig", types.YChild{"CHsrpGlobalConfig", &cISCOHSRPMIB.CHsrpGlobalConfig})
    cISCOHSRPMIB.EntityData.Children.Append("cHsrpGrpTable", types.YChild{"CHsrpGrpTable", &cISCOHSRPMIB.CHsrpGrpTable})
    cISCOHSRPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOHSRPMIB.EntityData.YListKeys = []string {}

    return &(cISCOHSRPMIB.EntityData)
}

// CISCOHSRPMIB_CHsrpGlobalConfig
type CISCOHSRPMIB_CHsrpGlobalConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The amount of time in minutes a row in cHsrpGrpTable can remain in a state
    // other than active before being timed out. The type is interface{} with
    // range: 1..60. Units are minutes.
    CHsrpConfigTimeout interface{}
}

func (cHsrpGlobalConfig *CISCOHSRPMIB_CHsrpGlobalConfig) GetEntityData() *types.CommonEntityData {
    cHsrpGlobalConfig.EntityData.YFilter = cHsrpGlobalConfig.YFilter
    cHsrpGlobalConfig.EntityData.YangName = "cHsrpGlobalConfig"
    cHsrpGlobalConfig.EntityData.BundleName = "cisco_ios_xe"
    cHsrpGlobalConfig.EntityData.ParentYangName = "CISCO-HSRP-MIB"
    cHsrpGlobalConfig.EntityData.SegmentPath = "cHsrpGlobalConfig"
    cHsrpGlobalConfig.EntityData.AbsolutePath = "CISCO-HSRP-MIB:CISCO-HSRP-MIB/" + cHsrpGlobalConfig.EntityData.SegmentPath
    cHsrpGlobalConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpGlobalConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpGlobalConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpGlobalConfig.EntityData.Children = types.NewOrderedMap()
    cHsrpGlobalConfig.EntityData.Leafs = types.NewOrderedMap()
    cHsrpGlobalConfig.EntityData.Leafs.Append("cHsrpConfigTimeout", types.YLeaf{"CHsrpConfigTimeout", cHsrpGlobalConfig.CHsrpConfigTimeout})

    cHsrpGlobalConfig.EntityData.YListKeys = []string {}

    return &(cHsrpGlobalConfig.EntityData)
}

// CISCOHSRPMIB_CHsrpGrpTable
// A table containing information on each HSRP group
// for each interface.
type CISCOHSRPMIB_CHsrpGrpTable struct {
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
    // CISCOHSRPMIB_CHsrpGrpTable_CHsrpGrpEntry.
    CHsrpGrpEntry []*CISCOHSRPMIB_CHsrpGrpTable_CHsrpGrpEntry
}

func (cHsrpGrpTable *CISCOHSRPMIB_CHsrpGrpTable) GetEntityData() *types.CommonEntityData {
    cHsrpGrpTable.EntityData.YFilter = cHsrpGrpTable.YFilter
    cHsrpGrpTable.EntityData.YangName = "cHsrpGrpTable"
    cHsrpGrpTable.EntityData.BundleName = "cisco_ios_xe"
    cHsrpGrpTable.EntityData.ParentYangName = "CISCO-HSRP-MIB"
    cHsrpGrpTable.EntityData.SegmentPath = "cHsrpGrpTable"
    cHsrpGrpTable.EntityData.AbsolutePath = "CISCO-HSRP-MIB:CISCO-HSRP-MIB/" + cHsrpGrpTable.EntityData.SegmentPath
    cHsrpGrpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpGrpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpGrpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpGrpTable.EntityData.Children = types.NewOrderedMap()
    cHsrpGrpTable.EntityData.Children.Append("cHsrpGrpEntry", types.YChild{"CHsrpGrpEntry", nil})
    for i := range cHsrpGrpTable.CHsrpGrpEntry {
        cHsrpGrpTable.EntityData.Children.Append(types.GetSegmentPath(cHsrpGrpTable.CHsrpGrpEntry[i]), types.YChild{"CHsrpGrpEntry", cHsrpGrpTable.CHsrpGrpEntry[i]})
    }
    cHsrpGrpTable.EntityData.Leafs = types.NewOrderedMap()

    cHsrpGrpTable.EntityData.YListKeys = []string {}

    return &(cHsrpGrpTable.EntityData)
}

// CISCOHSRPMIB_CHsrpGrpTable_CHsrpGrpEntry
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
type CISCOHSRPMIB_CHsrpGrpTable_CHsrpGrpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This object along with the ifIndex of a particular
    // interface uniquely identifies an HSRP group.  Group numbers 0,1 and 2 are
    // the only valid group numbers for TokenRing interfaces. For other media
    // types, numbers range from 0 to 255. Each interface has its own set of group
    // numbers. There's no relationship between the groups configured on different
    // interfaces. Using a group number on one interface doesn't preclude using
    // the same group number on a different interface. For example, there can be a
    // group 1 on an Ethernet and a group 1 on Token Ring. More details can be
    // found from RFC 2281. The type is interface{} with range: 0..255.
    CHsrpGrpNumber interface{}

    // This is an unencrypted authentication string which is carried in all HSRP
    // messages. An authentication string mismatch prevents a router interface
    // from learning the designated IP address or HSRP timer values from other
    // HSRP-enabled routers with the same group number.  The function of this
    // object is not to supply any sort of security-like authentication but rather
    // to confirm that what's happening is what's intended. In other words, this
    // is meant for sanity checking only. The type is string with length: 0..8.
    CHsrpGrpAuth interface{}

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
    CHsrpGrpPriority interface{}

    // This object, if TRUE, indicates that the current router should attempt to
    // overthrow a lower priority active router and attempt to become the active
    // router. If this object is FALSE, the router will become the active router
    // only if there is no such router (or if an active router fails). The type is
    // bool.
    CHsrpGrpPreempt interface{}

    // This delay is the time difference between a router power up and the time it
    // can actually start preempting the currently active router.  When a router
    // first comes up, it doesn't have a complete routing table. If it's
    // configured to preempt, then it will become the Active router, but it will
    // not be able to provide adequate routing services. The solution to this is
    // to allow for a configurable delay before the router actually preempts the
    // currently active router. The type is interface{} with range: 0..3600. Units
    // are seconds.
    CHsrpGrpPreemptDelay interface{}

    // HSRP routers learn a group's Hellotime or Holdtime from hello messages. 
    // The Hellotime is used to determine the frequency of generating hello
    // messages when this router becomes the active or standby router. The
    // Holdtime is the interval between the receipt of a Hello message and the
    // presumption that the sending router has failed.  If this object is TRUE,
    // the cHsrpGrpConfiguredHelloTime and cHsrpGrpConfiguredHoldTime will be
    // used. If it is FALSE, the Hellotime and Holdtime values are learned. The
    // type is bool.
    CHsrpGrpUseConfiguredTimers interface{}

    // If cHsrpGrpUseConfiguredTimers is true, cHsrpGrpConfiguredHelloTime is used
    // when this router is an active router. Otherwise, the Hellotime learned from
    // the current active router is used. All routers on a particular LAN segment
    // must use the same Hellotime. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    CHsrpGrpConfiguredHelloTime interface{}

    // If cHsrpGrpUseConfiguredTimers is true, cHsrpGrpConfiguredHoldTime is used
    // when this router is an active router. Otherwise, the Holdtime learned from
    // the current active router is used. All routers on a particular LAN segment
    // should use the same Holdtime. Also, the Holdtime should be at least three
    // times the value of the Hellotime and must be greater than the Hellotime.
    // The type is interface{} with range: 0..4294967295. Units are milliseconds.
    CHsrpGrpConfiguredHoldTime interface{}

    // If the Hellotime is not configured on a router, it can be learned from the
    // Hello messages from active router, provided the Hello message is
    // authenticated. If the Hellotime is not learned from a Hello message from
    // the active router and it is not manually configured, a default value of 3
    // seconds is recommended. The type is interface{} with range: 0..4294967295.
    // Units are milliseconds.
    CHsrpGrpLearnedHelloTime interface{}

    // If the Holdtime is not configured on a router, it can be learned from the
    // Hello message from the active router. Holdtime should be learned only if
    // the Hello message is authenticated. If the Holdtime is not learned and it
    // is not manually configured, a default value of 10 seconds is  recommended.
    // The type is interface{} with range: 0..4294967295. Units are milliseconds.
    CHsrpGrpLearnedHoldTime interface{}

    // This is the primary virtual IP address used by this group.  If this address
    // is configured (i.e a non zero ip address), this value is used. Otherwise,
    // the agent will attempt to discover the virtual address through a discovery
    // process (which scans the hello messages). The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CHsrpGrpVirtualIpAddr interface{}

    // If this object is TRUE, cHsrpGrpVirtualIpAddr was a configured one.
    // Otherwise, it indicates that  cHsrpGrpVirtualIpAddr was a learned one. The
    // type is bool.
    CHsrpGrpUseConfigVirtualIpAddr interface{}

    // Ip Address of the currently active router for this group. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CHsrpGrpActiveRouter interface{}

    // Ip Address of the currently standby router for this group. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CHsrpGrpStandbyRouter interface{}

    // The current HSRP state of this group on this interface. The type is
    // HsrpState.
    CHsrpGrpStandbyState interface{}

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
    CHsrpGrpVirtualMacAddr interface{}

    // The control that allows modification, creation, and deletion of entries. 
    // For detailed rules see the DESCRIPTION for cHsrpGrpEntry. The type is
    // RowStatus.
    CHsrpGrpEntryRowStatus interface{}

    // This object specifies the disable HSRP IPv4 virtual IP address. The type is
    // bool.
    CHsrpGrpIpNone interface{}
}

func (cHsrpGrpEntry *CISCOHSRPMIB_CHsrpGrpTable_CHsrpGrpEntry) GetEntityData() *types.CommonEntityData {
    cHsrpGrpEntry.EntityData.YFilter = cHsrpGrpEntry.YFilter
    cHsrpGrpEntry.EntityData.YangName = "cHsrpGrpEntry"
    cHsrpGrpEntry.EntityData.BundleName = "cisco_ios_xe"
    cHsrpGrpEntry.EntityData.ParentYangName = "cHsrpGrpTable"
    cHsrpGrpEntry.EntityData.SegmentPath = "cHsrpGrpEntry" + types.AddKeyToken(cHsrpGrpEntry.IfIndex, "ifIndex") + types.AddKeyToken(cHsrpGrpEntry.CHsrpGrpNumber, "cHsrpGrpNumber")
    cHsrpGrpEntry.EntityData.AbsolutePath = "CISCO-HSRP-MIB:CISCO-HSRP-MIB/cHsrpGrpTable/" + cHsrpGrpEntry.EntityData.SegmentPath
    cHsrpGrpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpGrpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpGrpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpGrpEntry.EntityData.Children = types.NewOrderedMap()
    cHsrpGrpEntry.EntityData.Leafs = types.NewOrderedMap()
    cHsrpGrpEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cHsrpGrpEntry.IfIndex})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpNumber", types.YLeaf{"CHsrpGrpNumber", cHsrpGrpEntry.CHsrpGrpNumber})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpAuth", types.YLeaf{"CHsrpGrpAuth", cHsrpGrpEntry.CHsrpGrpAuth})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpPriority", types.YLeaf{"CHsrpGrpPriority", cHsrpGrpEntry.CHsrpGrpPriority})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpPreempt", types.YLeaf{"CHsrpGrpPreempt", cHsrpGrpEntry.CHsrpGrpPreempt})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpPreemptDelay", types.YLeaf{"CHsrpGrpPreemptDelay", cHsrpGrpEntry.CHsrpGrpPreemptDelay})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpUseConfiguredTimers", types.YLeaf{"CHsrpGrpUseConfiguredTimers", cHsrpGrpEntry.CHsrpGrpUseConfiguredTimers})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpConfiguredHelloTime", types.YLeaf{"CHsrpGrpConfiguredHelloTime", cHsrpGrpEntry.CHsrpGrpConfiguredHelloTime})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpConfiguredHoldTime", types.YLeaf{"CHsrpGrpConfiguredHoldTime", cHsrpGrpEntry.CHsrpGrpConfiguredHoldTime})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpLearnedHelloTime", types.YLeaf{"CHsrpGrpLearnedHelloTime", cHsrpGrpEntry.CHsrpGrpLearnedHelloTime})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpLearnedHoldTime", types.YLeaf{"CHsrpGrpLearnedHoldTime", cHsrpGrpEntry.CHsrpGrpLearnedHoldTime})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpVirtualIpAddr", types.YLeaf{"CHsrpGrpVirtualIpAddr", cHsrpGrpEntry.CHsrpGrpVirtualIpAddr})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpUseConfigVirtualIpAddr", types.YLeaf{"CHsrpGrpUseConfigVirtualIpAddr", cHsrpGrpEntry.CHsrpGrpUseConfigVirtualIpAddr})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpActiveRouter", types.YLeaf{"CHsrpGrpActiveRouter", cHsrpGrpEntry.CHsrpGrpActiveRouter})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpStandbyRouter", types.YLeaf{"CHsrpGrpStandbyRouter", cHsrpGrpEntry.CHsrpGrpStandbyRouter})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpStandbyState", types.YLeaf{"CHsrpGrpStandbyState", cHsrpGrpEntry.CHsrpGrpStandbyState})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpVirtualMacAddr", types.YLeaf{"CHsrpGrpVirtualMacAddr", cHsrpGrpEntry.CHsrpGrpVirtualMacAddr})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpEntryRowStatus", types.YLeaf{"CHsrpGrpEntryRowStatus", cHsrpGrpEntry.CHsrpGrpEntryRowStatus})
    cHsrpGrpEntry.EntityData.Leafs.Append("cHsrpGrpIpNone", types.YLeaf{"CHsrpGrpIpNone", cHsrpGrpEntry.CHsrpGrpIpNone})

    cHsrpGrpEntry.EntityData.YListKeys = []string {"IfIndex", "CHsrpGrpNumber"}

    return &(cHsrpGrpEntry.EntityData)
}

