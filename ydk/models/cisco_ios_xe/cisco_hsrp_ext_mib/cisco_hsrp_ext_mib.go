// The Extension MIB module for the CISCO-HSRP-MIB which is
// based on RFC2281.
// 
// This MIB provides an extension to the CISCO-HSRP-MIB which 
// defines Cisco's proprietary Hot Standby Routing Protocol 
// (HSRP), defined in RFC2281. The extensions cover assigning 
// of secondary HSRP ip addresses, modifying an HSRP Group's 
// priority by tracking the operational status of interfaces, 
// etc. 
// 
// Terminology:
// HSRP is a protocol used amoung a group of routers for the 
// purpose of selecting an active router and a standby router. 
// 
// An active router is the router of choice for routing 
// packets.
// 
// A standby router is a router that takes over the routing 
// duties when an active router fails, or when preset 
// conditions have been met.
// 
// A HSRP group or a standby group is a set of routers 
// which communicate using HSRP. An HSRP group has a group 
// MAC address and a group IP address. These are the 
// designated addresses. The active router assumes  
// (i.e. inherits) these group addresses. An HSRP group is
// identified by a ( ifIndex, cHsrpGrpNumber ) pair.
// 
// BIA stands for Burned In Address.
package cisco_hsrp_ext_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_hsrp_ext_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-HSRP-EXT-MIB CISCO-HSRP-EXT-MIB}", reflect.TypeOf(CISCOHSRPEXTMIB{}))
    ydk.RegisterEntity("CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB", reflect.TypeOf(CISCOHSRPEXTMIB{}))
}

// CISCOHSRPEXTMIB
type CISCOHSRPEXTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table containing information about tracked interfaces per HSRP group.
    CHsrpExtIfTrackedTable CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable

    // A table containing information about secondary HSRP IP Addresses per
    // interface and group.
    CHsrpExtSecAddrTable CISCOHSRPEXTMIB_CHsrpExtSecAddrTable

    // A table containing information about standby interfaces per HSRP group.
    CHsrpExtIfStandbyTable CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable

    // HSRP-specific configurations for each physical interface.
    CHsrpExtIfTable CISCOHSRPEXTMIB_CHsrpExtIfTable
}

func (cISCOHSRPEXTMIB *CISCOHSRPEXTMIB) GetEntityData() *types.CommonEntityData {
    cISCOHSRPEXTMIB.EntityData.YFilter = cISCOHSRPEXTMIB.YFilter
    cISCOHSRPEXTMIB.EntityData.YangName = "CISCO-HSRP-EXT-MIB"
    cISCOHSRPEXTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOHSRPEXTMIB.EntityData.ParentYangName = "CISCO-HSRP-EXT-MIB"
    cISCOHSRPEXTMIB.EntityData.SegmentPath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB"
    cISCOHSRPEXTMIB.EntityData.AbsolutePath = cISCOHSRPEXTMIB.EntityData.SegmentPath
    cISCOHSRPEXTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOHSRPEXTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOHSRPEXTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOHSRPEXTMIB.EntityData.Children = types.NewOrderedMap()
    cISCOHSRPEXTMIB.EntityData.Children.Append("cHsrpExtIfTrackedTable", types.YChild{"CHsrpExtIfTrackedTable", &cISCOHSRPEXTMIB.CHsrpExtIfTrackedTable})
    cISCOHSRPEXTMIB.EntityData.Children.Append("cHsrpExtSecAddrTable", types.YChild{"CHsrpExtSecAddrTable", &cISCOHSRPEXTMIB.CHsrpExtSecAddrTable})
    cISCOHSRPEXTMIB.EntityData.Children.Append("cHsrpExtIfStandbyTable", types.YChild{"CHsrpExtIfStandbyTable", &cISCOHSRPEXTMIB.CHsrpExtIfStandbyTable})
    cISCOHSRPEXTMIB.EntityData.Children.Append("cHsrpExtIfTable", types.YChild{"CHsrpExtIfTable", &cISCOHSRPEXTMIB.CHsrpExtIfTable})
    cISCOHSRPEXTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOHSRPEXTMIB.EntityData.YListKeys = []string {}

    return &(cISCOHSRPEXTMIB.EntityData)
}

// CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable
// A table containing information about tracked interfaces per
// HSRP group.
type CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row of this table allows the tracking of one interface of the HSRP
    // group which is identified by the (ifIndex, cHsrpGrpNumber) values in this
    // table's INDEX clause. Weight(priority) is given to each and every interface
    // tracked.  When a tracked interface is unavailable, the HSRP priority of the
    // router is decreased. i.e cHsrpGrpPriority value assigned  to an HSRP group
    // will reduce by the value assigned to cHsrpExtIfTrackedPriority. This
    // reduces the likelihood  of a router with a failed key interface becoming
    // the  active router.  Setting cHsrpExtIfTrackedRowStatus to active starts
    // the tracking of cHsrpExtIfTracked by the HSRP group.  The value of
    // cHsrpExtIfTrackedRowStatus may be set  to destroy at any time.  Entries may
    // not be created via SNMP without explicitly setting
    // cHsrpExtIfTrackedRowStatus to either 'createAndGo'  or 'createAndWait'. 
    // Entries can be created and modified via the management protocol or by the
    // device's local management interface.  If the row is not active, and a local
    // management interface command modifies that row, the row may transition to
    // active state.  A row entry in the cHsrpExtIfTrackedTable can not be created
    // unless the corresponding row in the cHsrpGrpTable has been  created. If
    // that corresponding row in cHsrpGrpTable is  deleted, the interfaces it
    // tracks also get deleted.  A row which is not in active state will timeout
    // after a configurable period (five minutes by default). This timeout period
    // can be changed by setting cHsrpConfigTimeout. The type is slice of
    // CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable_CHsrpExtIfTrackedEntry.
    CHsrpExtIfTrackedEntry []*CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable_CHsrpExtIfTrackedEntry
}

func (cHsrpExtIfTrackedTable *CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable) GetEntityData() *types.CommonEntityData {
    cHsrpExtIfTrackedTable.EntityData.YFilter = cHsrpExtIfTrackedTable.YFilter
    cHsrpExtIfTrackedTable.EntityData.YangName = "cHsrpExtIfTrackedTable"
    cHsrpExtIfTrackedTable.EntityData.BundleName = "cisco_ios_xe"
    cHsrpExtIfTrackedTable.EntityData.ParentYangName = "CISCO-HSRP-EXT-MIB"
    cHsrpExtIfTrackedTable.EntityData.SegmentPath = "cHsrpExtIfTrackedTable"
    cHsrpExtIfTrackedTable.EntityData.AbsolutePath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB/" + cHsrpExtIfTrackedTable.EntityData.SegmentPath
    cHsrpExtIfTrackedTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpExtIfTrackedTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpExtIfTrackedTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpExtIfTrackedTable.EntityData.Children = types.NewOrderedMap()
    cHsrpExtIfTrackedTable.EntityData.Children.Append("cHsrpExtIfTrackedEntry", types.YChild{"CHsrpExtIfTrackedEntry", nil})
    for i := range cHsrpExtIfTrackedTable.CHsrpExtIfTrackedEntry {
        cHsrpExtIfTrackedTable.EntityData.Children.Append(types.GetSegmentPath(cHsrpExtIfTrackedTable.CHsrpExtIfTrackedEntry[i]), types.YChild{"CHsrpExtIfTrackedEntry", cHsrpExtIfTrackedTable.CHsrpExtIfTrackedEntry[i]})
    }
    cHsrpExtIfTrackedTable.EntityData.Leafs = types.NewOrderedMap()

    cHsrpExtIfTrackedTable.EntityData.YListKeys = []string {}

    return &(cHsrpExtIfTrackedTable.EntityData)
}

// CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable_CHsrpExtIfTrackedEntry
// Each row of this table allows the tracking of one
// interface of the HSRP group which is identified by the
// (ifIndex, cHsrpGrpNumber) values in this table's INDEX clause.
// Weight(priority) is given to each and every interface tracked. 
// When a tracked interface is unavailable, the HSRP priority of
// the router is decreased. i.e cHsrpGrpPriority value assigned 
// to an HSRP group will reduce by the value assigned to
// cHsrpExtIfTrackedPriority. This reduces the likelihood 
// of a router with a failed key interface becoming the 
// active router.
// 
// Setting cHsrpExtIfTrackedRowStatus to active starts
// the tracking of cHsrpExtIfTracked by the HSRP group. 
// The value of cHsrpExtIfTrackedRowStatus may be set 
// to destroy at any time.
// 
// Entries may not be created via SNMP without explicitly setting
// cHsrpExtIfTrackedRowStatus to either 'createAndGo' 
// or 'createAndWait'.
// 
// Entries can be created and modified via the management
// protocol or by the device's local management interface.
// 
// If the row is not active, and a local management interface
// command modifies that row, the row may transition to active
// state.
// 
// A row entry in the cHsrpExtIfTrackedTable can not be created
// unless the corresponding row in the cHsrpGrpTable has been 
// created. If that corresponding row in cHsrpGrpTable is 
// deleted, the interfaces it tracks also get deleted.
// 
// A row which is not in active state will timeout after a
// configurable period (five minutes by default). This timeout
// period can be changed by setting cHsrpConfigTimeout.
type CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable_CHsrpExtIfTrackedEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..255. Refers to
    // cisco_hsrp_mib.CISCOHSRPMIB_CHsrpGrpTable_CHsrpGrpEntry_CHsrpGrpNumber
    CHsrpGrpNumber interface{}

    // This attribute is a key. The ifIndex value of the tracked interface. The
    // type is interface{} with range: 1..2147483647.
    CHsrpExtIfTracked interface{}

    // Priority of the tracked interface for the corresponding { ifIndex,
    // cHsrpGrpNumber } pair. In the range of 0 to 255, 0 is the lowest priority
    // and 255 is the highest. When a tracked  interface is unavailable, the
    // cHsrpGrpPriority of the router  is decreased by the value of this object
    // instance (If the  cHsrpGrpPriority is less than the 
    // cHsrpExtIfTrackedPriority, then the HSRP priority  becomes 0). This allows
    // a standby router to be configured  with a priority such that if the
    // currently active router's  priority is lowered because the tracked
    // interface goes down,  the standby router can takeover. The type is
    // interface{} with range: 0..255.
    CHsrpExtIfTrackedPriority interface{}

    // The control that allows modification, creation, and deletion of entries.
    // For detailed rules see the DESCRIPTION for cHsrpExtIfTrackedEntry. The type
    // is RowStatus.
    CHsrpExtIfTrackedRowStatus interface{}

    // This object specifies the disable HSRP IPv4 virtual IP address. The type is
    // bool.
    CHsrpExtIfTrackedIpNone interface{}
}

func (cHsrpExtIfTrackedEntry *CISCOHSRPEXTMIB_CHsrpExtIfTrackedTable_CHsrpExtIfTrackedEntry) GetEntityData() *types.CommonEntityData {
    cHsrpExtIfTrackedEntry.EntityData.YFilter = cHsrpExtIfTrackedEntry.YFilter
    cHsrpExtIfTrackedEntry.EntityData.YangName = "cHsrpExtIfTrackedEntry"
    cHsrpExtIfTrackedEntry.EntityData.BundleName = "cisco_ios_xe"
    cHsrpExtIfTrackedEntry.EntityData.ParentYangName = "cHsrpExtIfTrackedTable"
    cHsrpExtIfTrackedEntry.EntityData.SegmentPath = "cHsrpExtIfTrackedEntry" + types.AddKeyToken(cHsrpExtIfTrackedEntry.IfIndex, "ifIndex") + types.AddKeyToken(cHsrpExtIfTrackedEntry.CHsrpGrpNumber, "cHsrpGrpNumber") + types.AddKeyToken(cHsrpExtIfTrackedEntry.CHsrpExtIfTracked, "cHsrpExtIfTracked")
    cHsrpExtIfTrackedEntry.EntityData.AbsolutePath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB/cHsrpExtIfTrackedTable/" + cHsrpExtIfTrackedEntry.EntityData.SegmentPath
    cHsrpExtIfTrackedEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpExtIfTrackedEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpExtIfTrackedEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpExtIfTrackedEntry.EntityData.Children = types.NewOrderedMap()
    cHsrpExtIfTrackedEntry.EntityData.Leafs = types.NewOrderedMap()
    cHsrpExtIfTrackedEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cHsrpExtIfTrackedEntry.IfIndex})
    cHsrpExtIfTrackedEntry.EntityData.Leafs.Append("cHsrpGrpNumber", types.YLeaf{"CHsrpGrpNumber", cHsrpExtIfTrackedEntry.CHsrpGrpNumber})
    cHsrpExtIfTrackedEntry.EntityData.Leafs.Append("cHsrpExtIfTracked", types.YLeaf{"CHsrpExtIfTracked", cHsrpExtIfTrackedEntry.CHsrpExtIfTracked})
    cHsrpExtIfTrackedEntry.EntityData.Leafs.Append("cHsrpExtIfTrackedPriority", types.YLeaf{"CHsrpExtIfTrackedPriority", cHsrpExtIfTrackedEntry.CHsrpExtIfTrackedPriority})
    cHsrpExtIfTrackedEntry.EntityData.Leafs.Append("cHsrpExtIfTrackedRowStatus", types.YLeaf{"CHsrpExtIfTrackedRowStatus", cHsrpExtIfTrackedEntry.CHsrpExtIfTrackedRowStatus})
    cHsrpExtIfTrackedEntry.EntityData.Leafs.Append("cHsrpExtIfTrackedIpNone", types.YLeaf{"CHsrpExtIfTrackedIpNone", cHsrpExtIfTrackedEntry.CHsrpExtIfTrackedIpNone})

    cHsrpExtIfTrackedEntry.EntityData.YListKeys = []string {"IfIndex", "CHsrpGrpNumber", "CHsrpExtIfTracked"}

    return &(cHsrpExtIfTrackedEntry.EntityData)
}

// CISCOHSRPEXTMIB_CHsrpExtSecAddrTable
// A table containing information about secondary HSRP IP
// Addresses per interface and group.
type CISCOHSRPEXTMIB_CHsrpExtSecAddrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The CHsrpExtSecAddrEntry allows creation of secondary IP Addresses for each
    // cHsrpGrpEntry row.  Secondary addresses can be added by setting 
    // cHsrpExtSecAddrRowStatus to be active. The value of
    // cHsrpExtSecAddrRowStatus may be set to destroy at any time.  Entries may
    // not be created via SNMP without explicitly setting cHsrpExtSecAddrRowStatus
    // to either 'createAndGo' or 'createAndWait'.  Entries can be created and
    // modified via the management protocol or by the device's local management
    // interface.  If the row is not active, and a local management interface
    // command modifies that row, the row may transition to active state.  A row
    // which is not in active state will timeout after a configurable period (five
    // minutes by default). This timeout period can be changed by setting
    // cHsrpConfigTimeout.  Before creation of a cHsrpExtSecAddrEntry row, either
    // cHsrpGrpConfiguredVirtualIpAddr or  cHsrpGrpLearnedVirtualIpAddr must have
    // a valid IP Address. This is because a secondary ip address cannot be
    // created unless the primary ip address has already been set.  To create a
    // new cHsrpExtSecAddrEntry row, a management  station should choose the
    // ifIndex of the interface which is to  be added as part of an HSRP group.
    // Also, an HSRP group number  and a cHsrpExtSecAddrAddress should be chosen. 
    // Deleting a {ifIndex, cHsrpGrpNumber} row in the cHsrpGrpTable will delete
    // all corresponding rows in the cHsrpExtSecAddrTable. Deleting a primary
    // address value in the cHsrpGrpEntry row will delete all secondary addresses
    // for the same {ifIndex, cHsrpGrpNumber} pair. The type is slice of
    // CISCOHSRPEXTMIB_CHsrpExtSecAddrTable_CHsrpExtSecAddrEntry.
    CHsrpExtSecAddrEntry []*CISCOHSRPEXTMIB_CHsrpExtSecAddrTable_CHsrpExtSecAddrEntry
}

func (cHsrpExtSecAddrTable *CISCOHSRPEXTMIB_CHsrpExtSecAddrTable) GetEntityData() *types.CommonEntityData {
    cHsrpExtSecAddrTable.EntityData.YFilter = cHsrpExtSecAddrTable.YFilter
    cHsrpExtSecAddrTable.EntityData.YangName = "cHsrpExtSecAddrTable"
    cHsrpExtSecAddrTable.EntityData.BundleName = "cisco_ios_xe"
    cHsrpExtSecAddrTable.EntityData.ParentYangName = "CISCO-HSRP-EXT-MIB"
    cHsrpExtSecAddrTable.EntityData.SegmentPath = "cHsrpExtSecAddrTable"
    cHsrpExtSecAddrTable.EntityData.AbsolutePath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB/" + cHsrpExtSecAddrTable.EntityData.SegmentPath
    cHsrpExtSecAddrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpExtSecAddrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpExtSecAddrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpExtSecAddrTable.EntityData.Children = types.NewOrderedMap()
    cHsrpExtSecAddrTable.EntityData.Children.Append("cHsrpExtSecAddrEntry", types.YChild{"CHsrpExtSecAddrEntry", nil})
    for i := range cHsrpExtSecAddrTable.CHsrpExtSecAddrEntry {
        cHsrpExtSecAddrTable.EntityData.Children.Append(types.GetSegmentPath(cHsrpExtSecAddrTable.CHsrpExtSecAddrEntry[i]), types.YChild{"CHsrpExtSecAddrEntry", cHsrpExtSecAddrTable.CHsrpExtSecAddrEntry[i]})
    }
    cHsrpExtSecAddrTable.EntityData.Leafs = types.NewOrderedMap()

    cHsrpExtSecAddrTable.EntityData.YListKeys = []string {}

    return &(cHsrpExtSecAddrTable.EntityData)
}

// CISCOHSRPEXTMIB_CHsrpExtSecAddrTable_CHsrpExtSecAddrEntry
// The CHsrpExtSecAddrEntry allows creation of secondary
// IP Addresses for each cHsrpGrpEntry row.
// 
// Secondary addresses can be added by setting 
// cHsrpExtSecAddrRowStatus to be active. The value of
// cHsrpExtSecAddrRowStatus may be set to destroy at any
// time.
// 
// Entries may not be created via SNMP without explicitly setting
// cHsrpExtSecAddrRowStatus to either 'createAndGo'
// or 'createAndWait'.
// 
// Entries can be created and modified via the management
// protocol or by the device's local management interface.
// 
// If the row is not active, and a local management interface
// command modifies that row, the row may transition to active
// state.
// 
// A row which is not in active state will timeout after a
// configurable period (five minutes by default). This timeout
// period can be changed by setting cHsrpConfigTimeout.
// 
// Before creation of a cHsrpExtSecAddrEntry row,
// either cHsrpGrpConfiguredVirtualIpAddr or 
// cHsrpGrpLearnedVirtualIpAddr must have a valid IP Address.
// This is because a secondary ip address cannot be created
// unless the primary ip address has already been set.
// 
// To create a new cHsrpExtSecAddrEntry row, a management 
// station should choose the ifIndex of the interface which is to 
// be added as part of an HSRP group. Also, an HSRP group number 
// and a cHsrpExtSecAddrAddress should be chosen.
// 
// Deleting a {ifIndex, cHsrpGrpNumber} row in the
// cHsrpGrpTable will delete all corresponding
// rows in the cHsrpExtSecAddrTable.
// Deleting a primary address value in the cHsrpGrpEntry row
// will delete all secondary addresses for the same
// {ifIndex, cHsrpGrpNumber} pair.
type CISCOHSRPEXTMIB_CHsrpExtSecAddrTable_CHsrpExtSecAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..255. Refers to
    // cisco_hsrp_mib.CISCOHSRPMIB_CHsrpGrpTable_CHsrpGrpEntry_CHsrpGrpNumber
    CHsrpGrpNumber interface{}

    // This attribute is a key. A secondary IpAddress for the {ifIndex,
    // cHsrpGrpNumber} pair. As explained in the DESCRIPTION for
    // cHsrpExtSecAddrEntry, a primary address must exist before a secondary
    // address for  the same {ifIndex, cHsrpGrpNumber} pair can be created. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CHsrpExtSecAddrAddress interface{}

    // The control that allows modification, creation, and deletion of entries.
    // For detailed rules see the DESCRIPTION for cHsrpExtSecAddrEntry. The type
    // is RowStatus.
    CHsrpExtSecAddrRowStatus interface{}
}

func (cHsrpExtSecAddrEntry *CISCOHSRPEXTMIB_CHsrpExtSecAddrTable_CHsrpExtSecAddrEntry) GetEntityData() *types.CommonEntityData {
    cHsrpExtSecAddrEntry.EntityData.YFilter = cHsrpExtSecAddrEntry.YFilter
    cHsrpExtSecAddrEntry.EntityData.YangName = "cHsrpExtSecAddrEntry"
    cHsrpExtSecAddrEntry.EntityData.BundleName = "cisco_ios_xe"
    cHsrpExtSecAddrEntry.EntityData.ParentYangName = "cHsrpExtSecAddrTable"
    cHsrpExtSecAddrEntry.EntityData.SegmentPath = "cHsrpExtSecAddrEntry" + types.AddKeyToken(cHsrpExtSecAddrEntry.IfIndex, "ifIndex") + types.AddKeyToken(cHsrpExtSecAddrEntry.CHsrpGrpNumber, "cHsrpGrpNumber") + types.AddKeyToken(cHsrpExtSecAddrEntry.CHsrpExtSecAddrAddress, "cHsrpExtSecAddrAddress")
    cHsrpExtSecAddrEntry.EntityData.AbsolutePath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB/cHsrpExtSecAddrTable/" + cHsrpExtSecAddrEntry.EntityData.SegmentPath
    cHsrpExtSecAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpExtSecAddrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpExtSecAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpExtSecAddrEntry.EntityData.Children = types.NewOrderedMap()
    cHsrpExtSecAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    cHsrpExtSecAddrEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cHsrpExtSecAddrEntry.IfIndex})
    cHsrpExtSecAddrEntry.EntityData.Leafs.Append("cHsrpGrpNumber", types.YLeaf{"CHsrpGrpNumber", cHsrpExtSecAddrEntry.CHsrpGrpNumber})
    cHsrpExtSecAddrEntry.EntityData.Leafs.Append("cHsrpExtSecAddrAddress", types.YLeaf{"CHsrpExtSecAddrAddress", cHsrpExtSecAddrEntry.CHsrpExtSecAddrAddress})
    cHsrpExtSecAddrEntry.EntityData.Leafs.Append("cHsrpExtSecAddrRowStatus", types.YLeaf{"CHsrpExtSecAddrRowStatus", cHsrpExtSecAddrEntry.CHsrpExtSecAddrRowStatus})

    cHsrpExtSecAddrEntry.EntityData.YListKeys = []string {"IfIndex", "CHsrpGrpNumber", "CHsrpExtSecAddrAddress"}

    return &(cHsrpExtSecAddrEntry.EntityData)
}

// CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable
// A table containing information about standby
// interfaces per HSRP group.
type CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The cHsrpExtIfStandbyEntry allows an HSRP group interface to track one or
    // more standby interfaces.  To create a new cHsrpExtIfStandbyEntry row, a
    // management station should choose the ifIndex of the interface which is to
    // be added as part of an HSRP group. Also, an HSRP group number and a
    // cHsrpExtIfStandbyIndex should be chosen. The type is slice of
    // CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable_CHsrpExtIfStandbyEntry.
    CHsrpExtIfStandbyEntry []*CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable_CHsrpExtIfStandbyEntry
}

func (cHsrpExtIfStandbyTable *CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable) GetEntityData() *types.CommonEntityData {
    cHsrpExtIfStandbyTable.EntityData.YFilter = cHsrpExtIfStandbyTable.YFilter
    cHsrpExtIfStandbyTable.EntityData.YangName = "cHsrpExtIfStandbyTable"
    cHsrpExtIfStandbyTable.EntityData.BundleName = "cisco_ios_xe"
    cHsrpExtIfStandbyTable.EntityData.ParentYangName = "CISCO-HSRP-EXT-MIB"
    cHsrpExtIfStandbyTable.EntityData.SegmentPath = "cHsrpExtIfStandbyTable"
    cHsrpExtIfStandbyTable.EntityData.AbsolutePath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB/" + cHsrpExtIfStandbyTable.EntityData.SegmentPath
    cHsrpExtIfStandbyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpExtIfStandbyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpExtIfStandbyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpExtIfStandbyTable.EntityData.Children = types.NewOrderedMap()
    cHsrpExtIfStandbyTable.EntityData.Children.Append("cHsrpExtIfStandbyEntry", types.YChild{"CHsrpExtIfStandbyEntry", nil})
    for i := range cHsrpExtIfStandbyTable.CHsrpExtIfStandbyEntry {
        cHsrpExtIfStandbyTable.EntityData.Children.Append(types.GetSegmentPath(cHsrpExtIfStandbyTable.CHsrpExtIfStandbyEntry[i]), types.YChild{"CHsrpExtIfStandbyEntry", cHsrpExtIfStandbyTable.CHsrpExtIfStandbyEntry[i]})
    }
    cHsrpExtIfStandbyTable.EntityData.Leafs = types.NewOrderedMap()

    cHsrpExtIfStandbyTable.EntityData.YListKeys = []string {}

    return &(cHsrpExtIfStandbyTable.EntityData)
}

// CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable_CHsrpExtIfStandbyEntry
// The cHsrpExtIfStandbyEntry allows an HSRP group
// interface to track one or more standby interfaces.
// 
// To create a new cHsrpExtIfStandbyEntry row, a
// management station should choose the ifIndex of
// the interface which is to be added as part of an
// HSRP group. Also, an HSRP group number and a
// cHsrpExtIfStandbyIndex should be chosen.
type CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable_CHsrpExtIfStandbyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..255. Refers to
    // cisco_hsrp_mib.CISCOHSRPMIB_CHsrpGrpTable_CHsrpGrpEntry_CHsrpGrpNumber
    CHsrpGrpNumber interface{}

    // This attribute is a key. This object defines the index of the standby
    // table. The type is interface{} with range: 1..4.
    CHsrpExtIfStandbyIndex interface{}

    // This object specifies the type of Internet address denoted by
    // cHsrpExtIfStandbyDestAddr. The type is InetAddressType.
    CHsrpExtIfStandbyDestAddrType interface{}

    // This object specifies the destination IP address of the standby router. The
    // type is string with length: 0..255.
    CHsrpExtIfStandbyDestAddr interface{}

    // This object specifies the type of Internet address denoted by
    // cHsrpExtIfStandbySourceAddr. The type is InetAddressType.
    CHsrpExtIfStandbySourceAddrType interface{}

    // This object specifies the source IP address of the standby router. The type
    // is string with length: 0..255.
    CHsrpExtIfStandbySourceAddr interface{}

    // The control that allows modification, creation, and deletion of entries.
    // Entries may not be created via SNMP without explicitly setting
    // cHsrpExtIfStandbyRowStatus to either 'createAndGo' or 'createAndWait'. The
    // type is RowStatus.
    CHsrpExtIfStandbyRowStatus interface{}
}

func (cHsrpExtIfStandbyEntry *CISCOHSRPEXTMIB_CHsrpExtIfStandbyTable_CHsrpExtIfStandbyEntry) GetEntityData() *types.CommonEntityData {
    cHsrpExtIfStandbyEntry.EntityData.YFilter = cHsrpExtIfStandbyEntry.YFilter
    cHsrpExtIfStandbyEntry.EntityData.YangName = "cHsrpExtIfStandbyEntry"
    cHsrpExtIfStandbyEntry.EntityData.BundleName = "cisco_ios_xe"
    cHsrpExtIfStandbyEntry.EntityData.ParentYangName = "cHsrpExtIfStandbyTable"
    cHsrpExtIfStandbyEntry.EntityData.SegmentPath = "cHsrpExtIfStandbyEntry" + types.AddKeyToken(cHsrpExtIfStandbyEntry.IfIndex, "ifIndex") + types.AddKeyToken(cHsrpExtIfStandbyEntry.CHsrpGrpNumber, "cHsrpGrpNumber") + types.AddKeyToken(cHsrpExtIfStandbyEntry.CHsrpExtIfStandbyIndex, "cHsrpExtIfStandbyIndex")
    cHsrpExtIfStandbyEntry.EntityData.AbsolutePath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB/cHsrpExtIfStandbyTable/" + cHsrpExtIfStandbyEntry.EntityData.SegmentPath
    cHsrpExtIfStandbyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpExtIfStandbyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpExtIfStandbyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpExtIfStandbyEntry.EntityData.Children = types.NewOrderedMap()
    cHsrpExtIfStandbyEntry.EntityData.Leafs = types.NewOrderedMap()
    cHsrpExtIfStandbyEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cHsrpExtIfStandbyEntry.IfIndex})
    cHsrpExtIfStandbyEntry.EntityData.Leafs.Append("cHsrpGrpNumber", types.YLeaf{"CHsrpGrpNumber", cHsrpExtIfStandbyEntry.CHsrpGrpNumber})
    cHsrpExtIfStandbyEntry.EntityData.Leafs.Append("cHsrpExtIfStandbyIndex", types.YLeaf{"CHsrpExtIfStandbyIndex", cHsrpExtIfStandbyEntry.CHsrpExtIfStandbyIndex})
    cHsrpExtIfStandbyEntry.EntityData.Leafs.Append("cHsrpExtIfStandbyDestAddrType", types.YLeaf{"CHsrpExtIfStandbyDestAddrType", cHsrpExtIfStandbyEntry.CHsrpExtIfStandbyDestAddrType})
    cHsrpExtIfStandbyEntry.EntityData.Leafs.Append("cHsrpExtIfStandbyDestAddr", types.YLeaf{"CHsrpExtIfStandbyDestAddr", cHsrpExtIfStandbyEntry.CHsrpExtIfStandbyDestAddr})
    cHsrpExtIfStandbyEntry.EntityData.Leafs.Append("cHsrpExtIfStandbySourceAddrType", types.YLeaf{"CHsrpExtIfStandbySourceAddrType", cHsrpExtIfStandbyEntry.CHsrpExtIfStandbySourceAddrType})
    cHsrpExtIfStandbyEntry.EntityData.Leafs.Append("cHsrpExtIfStandbySourceAddr", types.YLeaf{"CHsrpExtIfStandbySourceAddr", cHsrpExtIfStandbyEntry.CHsrpExtIfStandbySourceAddr})
    cHsrpExtIfStandbyEntry.EntityData.Leafs.Append("cHsrpExtIfStandbyRowStatus", types.YLeaf{"CHsrpExtIfStandbyRowStatus", cHsrpExtIfStandbyEntry.CHsrpExtIfStandbyRowStatus})

    cHsrpExtIfStandbyEntry.EntityData.YListKeys = []string {"IfIndex", "CHsrpGrpNumber", "CHsrpExtIfStandbyIndex"}

    return &(cHsrpExtIfStandbyEntry.EntityData)
}

// CISCOHSRPEXTMIB_CHsrpExtIfTable
// HSRP-specific configurations for each physical interface.
type CISCOHSRPEXTMIB_CHsrpExtIfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If HSRP entries on this interface must use the BIA (Burned In Address),
    // there must be an entry for the interface in this  table. Entries of this
    // table are only accessible if HSRP has  been enabled i.e entries can not be
    // created if HSRP is not enabled. Also, the interfaces should be of IEEE 802
    // ones (Ethernet, Token Ring, FDDI,VLAN, LANE, TR-LANE).  Setting
    // cHsrpExtIfRowStatus to active initiates the entry with default value for
    // cHsrpExtIfUseBIA as FALSE. The value of cHsrpExtIfRowStatus may be set to
    // destroy at any time. If the row is not initiated, it is similar to having
    // cHsrpExtIfUseBIA as FALSE.  Entries may not be created via SNMP without
    // explicitly setting cHsrpExtIfRowStatus to either 'createAndGo' or
    // 'createAndWait'.  Entries can be created and modified via the management
    // protocol or by the device's local management interface.  If the row is not
    // active, and a local management interface command modifies that row, the row
    // may transition to active state.  A row which is not in active state will
    // timeout after a configurable period (five minutes by default). This timeout
    // period can be changed by setting cHsrpConfigTimeout. The type is slice of
    // CISCOHSRPEXTMIB_CHsrpExtIfTable_CHsrpExtIfEntry.
    CHsrpExtIfEntry []*CISCOHSRPEXTMIB_CHsrpExtIfTable_CHsrpExtIfEntry
}

func (cHsrpExtIfTable *CISCOHSRPEXTMIB_CHsrpExtIfTable) GetEntityData() *types.CommonEntityData {
    cHsrpExtIfTable.EntityData.YFilter = cHsrpExtIfTable.YFilter
    cHsrpExtIfTable.EntityData.YangName = "cHsrpExtIfTable"
    cHsrpExtIfTable.EntityData.BundleName = "cisco_ios_xe"
    cHsrpExtIfTable.EntityData.ParentYangName = "CISCO-HSRP-EXT-MIB"
    cHsrpExtIfTable.EntityData.SegmentPath = "cHsrpExtIfTable"
    cHsrpExtIfTable.EntityData.AbsolutePath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB/" + cHsrpExtIfTable.EntityData.SegmentPath
    cHsrpExtIfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpExtIfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpExtIfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpExtIfTable.EntityData.Children = types.NewOrderedMap()
    cHsrpExtIfTable.EntityData.Children.Append("cHsrpExtIfEntry", types.YChild{"CHsrpExtIfEntry", nil})
    for i := range cHsrpExtIfTable.CHsrpExtIfEntry {
        cHsrpExtIfTable.EntityData.Children.Append(types.GetSegmentPath(cHsrpExtIfTable.CHsrpExtIfEntry[i]), types.YChild{"CHsrpExtIfEntry", cHsrpExtIfTable.CHsrpExtIfEntry[i]})
    }
    cHsrpExtIfTable.EntityData.Leafs = types.NewOrderedMap()

    cHsrpExtIfTable.EntityData.YListKeys = []string {}

    return &(cHsrpExtIfTable.EntityData)
}

// CISCOHSRPEXTMIB_CHsrpExtIfTable_CHsrpExtIfEntry
// If HSRP entries on this interface must use the BIA (Burned
// In Address), there must be an entry for the interface in this 
// table. Entries of this table are only accessible if HSRP has 
// been enabled i.e entries can not be created if HSRP is not
// enabled. Also, the interfaces should be of IEEE 802 ones
// (Ethernet, Token Ring, FDDI,VLAN, LANE, TR-LANE).
// 
// Setting cHsrpExtIfRowStatus to active initiates the
// entry with default value for cHsrpExtIfUseBIA as FALSE.
// The value of cHsrpExtIfRowStatus may be set to destroy
// at any time. If the row is not initiated, it is similar to
// having cHsrpExtIfUseBIA as FALSE.
// 
// Entries may not be created via SNMP without explicitly setting
// cHsrpExtIfRowStatus to either 'createAndGo' or 'createAndWait'.
// 
// Entries can be created and modified via the management
// protocol or by the device's local management interface.
// 
// If the row is not active, and a local management interface
// command modifies that row, the row may transition to active
// state.
// 
// A row which is not in active state will timeout after a
// configurable period (five minutes by default). This timeout
// period can be changed by setting cHsrpConfigTimeout.
type CISCOHSRPEXTMIB_CHsrpExtIfTable_CHsrpExtIfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // If set to TRUE, the HSRP Group MAC Address for all groups on this 
    // interface will be the burned-in-address. Otherwise, this will be determined
    // by ciscoHsrpGroupNumber. In case of sub-interfaces, UseBIA applies to all
    // sub-interfaces on an  interface and to all groups on those sub-interfaces.
    // The type is bool.
    CHsrpExtIfUseBIA interface{}

    // The control that allows modification, creation, and deletion of entries.
    // For detailed rules see the DESCRIPTION for cHsrpExtIfEntry. The type is
    // RowStatus.
    CHsrpExtIfRowStatus interface{}
}

func (cHsrpExtIfEntry *CISCOHSRPEXTMIB_CHsrpExtIfTable_CHsrpExtIfEntry) GetEntityData() *types.CommonEntityData {
    cHsrpExtIfEntry.EntityData.YFilter = cHsrpExtIfEntry.YFilter
    cHsrpExtIfEntry.EntityData.YangName = "cHsrpExtIfEntry"
    cHsrpExtIfEntry.EntityData.BundleName = "cisco_ios_xe"
    cHsrpExtIfEntry.EntityData.ParentYangName = "cHsrpExtIfTable"
    cHsrpExtIfEntry.EntityData.SegmentPath = "cHsrpExtIfEntry" + types.AddKeyToken(cHsrpExtIfEntry.IfIndex, "ifIndex")
    cHsrpExtIfEntry.EntityData.AbsolutePath = "CISCO-HSRP-EXT-MIB:CISCO-HSRP-EXT-MIB/cHsrpExtIfTable/" + cHsrpExtIfEntry.EntityData.SegmentPath
    cHsrpExtIfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cHsrpExtIfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cHsrpExtIfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cHsrpExtIfEntry.EntityData.Children = types.NewOrderedMap()
    cHsrpExtIfEntry.EntityData.Leafs = types.NewOrderedMap()
    cHsrpExtIfEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cHsrpExtIfEntry.IfIndex})
    cHsrpExtIfEntry.EntityData.Leafs.Append("cHsrpExtIfUseBIA", types.YLeaf{"CHsrpExtIfUseBIA", cHsrpExtIfEntry.CHsrpExtIfUseBIA})
    cHsrpExtIfEntry.EntityData.Leafs.Append("cHsrpExtIfRowStatus", types.YLeaf{"CHsrpExtIfRowStatus", cHsrpExtIfEntry.CHsrpExtIfRowStatus})

    cHsrpExtIfEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cHsrpExtIfEntry.EntityData)
}

