// The MIB module for IGMP Management.
package igmp_std_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package igmp_std_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:IGMP-STD-MIB IGMP-STD-MIB}", reflect.TypeOf(IGMPSTDMIB{}))
    ydk.RegisterEntity("IGMP-STD-MIB:IGMP-STD-MIB", reflect.TypeOf(IGMPSTDMIB{}))
}

// IGMPSTDMIB
type IGMPSTDMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The (conceptual) table listing the interfaces on which IGMP is enabled.
    IgmpInterfaceTable IGMPSTDMIB_IgmpInterfaceTable

    // The (conceptual) table listing the IP multicast groups for which there are
    // members on a particular interface.
    IgmpCacheTable IGMPSTDMIB_IgmpCacheTable
}

func (iGMPSTDMIB *IGMPSTDMIB) GetEntityData() *types.CommonEntityData {
    iGMPSTDMIB.EntityData.YFilter = iGMPSTDMIB.YFilter
    iGMPSTDMIB.EntityData.YangName = "IGMP-STD-MIB"
    iGMPSTDMIB.EntityData.BundleName = "cisco_ios_xe"
    iGMPSTDMIB.EntityData.ParentYangName = "IGMP-STD-MIB"
    iGMPSTDMIB.EntityData.SegmentPath = "IGMP-STD-MIB:IGMP-STD-MIB"
    iGMPSTDMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iGMPSTDMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iGMPSTDMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iGMPSTDMIB.EntityData.Children = types.NewOrderedMap()
    iGMPSTDMIB.EntityData.Children.Append("igmpInterfaceTable", types.YChild{"IgmpInterfaceTable", &iGMPSTDMIB.IgmpInterfaceTable})
    iGMPSTDMIB.EntityData.Children.Append("igmpCacheTable", types.YChild{"IgmpCacheTable", &iGMPSTDMIB.IgmpCacheTable})
    iGMPSTDMIB.EntityData.Leafs = types.NewOrderedMap()

    iGMPSTDMIB.EntityData.YListKeys = []string {}

    return &(iGMPSTDMIB.EntityData)
}

// IGMPSTDMIB_IgmpInterfaceTable
// The (conceptual) table listing the interfaces on which IGMP
// is enabled.
type IGMPSTDMIB_IgmpInterfaceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an interface on which IGMP is
    // enabled. The type is slice of
    // IGMPSTDMIB_IgmpInterfaceTable_IgmpInterfaceEntry.
    IgmpInterfaceEntry []*IGMPSTDMIB_IgmpInterfaceTable_IgmpInterfaceEntry
}

func (igmpInterfaceTable *IGMPSTDMIB_IgmpInterfaceTable) GetEntityData() *types.CommonEntityData {
    igmpInterfaceTable.EntityData.YFilter = igmpInterfaceTable.YFilter
    igmpInterfaceTable.EntityData.YangName = "igmpInterfaceTable"
    igmpInterfaceTable.EntityData.BundleName = "cisco_ios_xe"
    igmpInterfaceTable.EntityData.ParentYangName = "IGMP-STD-MIB"
    igmpInterfaceTable.EntityData.SegmentPath = "igmpInterfaceTable"
    igmpInterfaceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpInterfaceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpInterfaceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpInterfaceTable.EntityData.Children = types.NewOrderedMap()
    igmpInterfaceTable.EntityData.Children.Append("igmpInterfaceEntry", types.YChild{"IgmpInterfaceEntry", nil})
    for i := range igmpInterfaceTable.IgmpInterfaceEntry {
        igmpInterfaceTable.EntityData.Children.Append(types.GetSegmentPath(igmpInterfaceTable.IgmpInterfaceEntry[i]), types.YChild{"IgmpInterfaceEntry", igmpInterfaceTable.IgmpInterfaceEntry[i]})
    }
    igmpInterfaceTable.EntityData.Leafs = types.NewOrderedMap()

    igmpInterfaceTable.EntityData.YListKeys = []string {}

    return &(igmpInterfaceTable.EntityData)
}

// IGMPSTDMIB_IgmpInterfaceTable_IgmpInterfaceEntry
// An entry (conceptual row) representing an interface on
// which IGMP is enabled.
type IGMPSTDMIB_IgmpInterfaceTable_IgmpInterfaceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the interface for which IGMP
    // is enabled. The type is interface{} with range: 1..2147483647.
    IgmpInterfaceIfIndex interface{}

    // The frequency at which IGMP Host-Query packets are transmitted on this
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    IgmpInterfaceQueryInterval interface{}

    // The activation of a row enables IGMP on the interface.  The destruction of
    // a row disables IGMP on the interface. The type is RowStatus.
    IgmpInterfaceStatus interface{}

    // The version of IGMP which is running on this interface. This object can be
    // used to configure a router capable of running either value.  For IGMP to
    // function correctly, all routers on a LAN must be configured to run the same
    // version of IGMP on that LAN. The type is interface{} with range:
    // 0..4294967295.
    IgmpInterfaceVersion interface{}

    // The address of the IGMP Querier on the IP subnet to which      this
    // interface is attached. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IgmpInterfaceQuerier interface{}

    // The maximum query response time advertised in IGMPv2 queries on this
    // interface. The type is interface{} with range: 0..255. Units are tenths of
    // seconds.
    IgmpInterfaceQueryMaxResponseTime interface{}

    // The time since igmpInterfaceQuerier was last changed. The type is
    // interface{} with range: 0..4294967295.
    IgmpInterfaceQuerierUpTime interface{}

    // The amount of time remaining before the Other Querier Present Timer
    // expires.  If the local system is the querier, the value of this object is
    // zero. The type is interface{} with range: 0..4294967295.
    IgmpInterfaceQuerierExpiryTime interface{}

    // The time remaining until the host assumes that there are no IGMPv1 routers
    // present on the interface.  While this is non- zero, the host will reply to
    // all queries with version 1 membership reports. The type is interface{} with
    // range: 0..4294967295.
    IgmpInterfaceVersion1QuerierTimer interface{}

    // The number of queries received whose IGMP version does not match
    // igmpInterfaceVersion, over the lifetime of the row entry.  IGMP requires
    // that all routers on a LAN be configured to run the same version of IGMP. 
    // Thus, if any queries are received with the wrong version, this indicates a
    // configuration error. The type is interface{} with range: 0..4294967295.
    IgmpInterfaceWrongVersionQueries interface{}

    // The number of times a group membership has been added on this interface;
    // that is, the number of times an entry for this interface has been added to
    // the Cache Table.  This object gives an indication of the amount of IGMP
    // activity over the lifetime of the row entry. The type is interface{} with
    // range: 0..4294967295.
    IgmpInterfaceJoins interface{}

    // Some devices implement a form of IGMP proxying whereby memberships learned
    // on the interface represented by this row, cause IGMP Host Membership
    // Reports to be sent on the interface whose ifIndex value is given by this
    // object.  Such a device would implement the igmpV2RouterMIBGroup only on its
    // router interfaces (those interfaces with non-zero
    // igmpInterfaceProxyIfIndex).  Typically, the value of this object is 0,
    // indicating that no proxying is being done. The type is interface{} with
    // range: 0..2147483647.
    IgmpInterfaceProxyIfIndex interface{}

    // The current number of entries for this interface in the Cache Table. The
    // type is interface{} with range: 0..4294967295.
    IgmpInterfaceGroups interface{}

    // The Robustness Variable allows tuning for the expected packet loss on a
    // subnet.  If a subnet is expected to be lossy, the Robustness Variable may
    // be increased.  IGMP is robust to (Robustness Variable-1) packet losses. The
    // type is interface{} with range: 1..255.
    IgmpInterfaceRobustness interface{}

    // The Last Member Query Interval is the Max Response Time inserted into
    // Group-Specific Queries sent in response to Leave Group messages, and is
    // also the amount of time between Group-Specific Query messages.  This value
    // may be tuned to modify the leave latency of the network.  A reduced value
    // results in reduced time to detect the loss of the last member of a group. 
    // The value of this object is irrelevant if igmpInterfaceVersion is 1. The
    // type is interface{} with range: 0..255. Units are tenths of seconds.
    IgmpInterfaceLastMembQueryIntvl interface{}
}

func (igmpInterfaceEntry *IGMPSTDMIB_IgmpInterfaceTable_IgmpInterfaceEntry) GetEntityData() *types.CommonEntityData {
    igmpInterfaceEntry.EntityData.YFilter = igmpInterfaceEntry.YFilter
    igmpInterfaceEntry.EntityData.YangName = "igmpInterfaceEntry"
    igmpInterfaceEntry.EntityData.BundleName = "cisco_ios_xe"
    igmpInterfaceEntry.EntityData.ParentYangName = "igmpInterfaceTable"
    igmpInterfaceEntry.EntityData.SegmentPath = "igmpInterfaceEntry" + types.AddKeyToken(igmpInterfaceEntry.IgmpInterfaceIfIndex, "igmpInterfaceIfIndex")
    igmpInterfaceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpInterfaceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpInterfaceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpInterfaceEntry.EntityData.Children = types.NewOrderedMap()
    igmpInterfaceEntry.EntityData.Leafs = types.NewOrderedMap()
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceIfIndex", types.YLeaf{"IgmpInterfaceIfIndex", igmpInterfaceEntry.IgmpInterfaceIfIndex})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceQueryInterval", types.YLeaf{"IgmpInterfaceQueryInterval", igmpInterfaceEntry.IgmpInterfaceQueryInterval})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceStatus", types.YLeaf{"IgmpInterfaceStatus", igmpInterfaceEntry.IgmpInterfaceStatus})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceVersion", types.YLeaf{"IgmpInterfaceVersion", igmpInterfaceEntry.IgmpInterfaceVersion})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceQuerier", types.YLeaf{"IgmpInterfaceQuerier", igmpInterfaceEntry.IgmpInterfaceQuerier})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceQueryMaxResponseTime", types.YLeaf{"IgmpInterfaceQueryMaxResponseTime", igmpInterfaceEntry.IgmpInterfaceQueryMaxResponseTime})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceQuerierUpTime", types.YLeaf{"IgmpInterfaceQuerierUpTime", igmpInterfaceEntry.IgmpInterfaceQuerierUpTime})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceQuerierExpiryTime", types.YLeaf{"IgmpInterfaceQuerierExpiryTime", igmpInterfaceEntry.IgmpInterfaceQuerierExpiryTime})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceVersion1QuerierTimer", types.YLeaf{"IgmpInterfaceVersion1QuerierTimer", igmpInterfaceEntry.IgmpInterfaceVersion1QuerierTimer})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceWrongVersionQueries", types.YLeaf{"IgmpInterfaceWrongVersionQueries", igmpInterfaceEntry.IgmpInterfaceWrongVersionQueries})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceJoins", types.YLeaf{"IgmpInterfaceJoins", igmpInterfaceEntry.IgmpInterfaceJoins})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceProxyIfIndex", types.YLeaf{"IgmpInterfaceProxyIfIndex", igmpInterfaceEntry.IgmpInterfaceProxyIfIndex})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceGroups", types.YLeaf{"IgmpInterfaceGroups", igmpInterfaceEntry.IgmpInterfaceGroups})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceRobustness", types.YLeaf{"IgmpInterfaceRobustness", igmpInterfaceEntry.IgmpInterfaceRobustness})
    igmpInterfaceEntry.EntityData.Leafs.Append("igmpInterfaceLastMembQueryIntvl", types.YLeaf{"IgmpInterfaceLastMembQueryIntvl", igmpInterfaceEntry.IgmpInterfaceLastMembQueryIntvl})

    igmpInterfaceEntry.EntityData.YListKeys = []string {"IgmpInterfaceIfIndex"}

    return &(igmpInterfaceEntry.EntityData)
}

// IGMPSTDMIB_IgmpCacheTable
// The (conceptual) table listing the IP multicast groups for
// which there are members on a particular interface.
type IGMPSTDMIB_IgmpCacheTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the igmpCacheTable. The type is slice of
    // IGMPSTDMIB_IgmpCacheTable_IgmpCacheEntry.
    IgmpCacheEntry []*IGMPSTDMIB_IgmpCacheTable_IgmpCacheEntry
}

func (igmpCacheTable *IGMPSTDMIB_IgmpCacheTable) GetEntityData() *types.CommonEntityData {
    igmpCacheTable.EntityData.YFilter = igmpCacheTable.YFilter
    igmpCacheTable.EntityData.YangName = "igmpCacheTable"
    igmpCacheTable.EntityData.BundleName = "cisco_ios_xe"
    igmpCacheTable.EntityData.ParentYangName = "IGMP-STD-MIB"
    igmpCacheTable.EntityData.SegmentPath = "igmpCacheTable"
    igmpCacheTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpCacheTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpCacheTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpCacheTable.EntityData.Children = types.NewOrderedMap()
    igmpCacheTable.EntityData.Children.Append("igmpCacheEntry", types.YChild{"IgmpCacheEntry", nil})
    for i := range igmpCacheTable.IgmpCacheEntry {
        igmpCacheTable.EntityData.Children.Append(types.GetSegmentPath(igmpCacheTable.IgmpCacheEntry[i]), types.YChild{"IgmpCacheEntry", igmpCacheTable.IgmpCacheEntry[i]})
    }
    igmpCacheTable.EntityData.Leafs = types.NewOrderedMap()

    igmpCacheTable.EntityData.YListKeys = []string {}

    return &(igmpCacheTable.EntityData)
}

// IGMPSTDMIB_IgmpCacheTable_IgmpCacheEntry
// An entry (conceptual row) in the igmpCacheTable.
type IGMPSTDMIB_IgmpCacheTable_IgmpCacheEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address for which this
    // entry contains information. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IgmpCacheAddress interface{}

    // This attribute is a key. The interface for which this entry contains
    // information for an IP multicast group address. The type is interface{} with
    // range: 1..2147483647.
    IgmpCacheIfIndex interface{}

    // An indication of whether the local system is a member of this group address
    // on this interface. The type is bool.
    IgmpCacheSelf interface{}

    // The IP address of the source of the last membership report received for
    // this IP Multicast group address on this interface.  If no membership report
    // has been received, this object has the value 0.0.0.0. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IgmpCacheLastReporter interface{}

    // The time elapsed since this entry was created. The type is interface{} with
    // range: 0..4294967295.
    IgmpCacheUpTime interface{}

    // The minimum amount of time remaining before this entry will be aged out.  A
    // value of 0 indicates that the entry is only present because igmpCacheSelf
    // is true and that if the router left the group, this entry would be aged out
    // immediately. Note that some implementations may process membership reports
    // from the local system in the same way as reports from other hosts, so a
    // value of 0 is not required. The type is interface{} with range:
    // 0..4294967295.
    IgmpCacheExpiryTime interface{}

    // The status of this entry. The type is RowStatus.
    IgmpCacheStatus interface{}

    // The time remaining until the local router will assume that there are no
    // longer any IGMP version 1 members on the IP subnet attached to this
    // interface.  Upon hearing any IGMPv1 Membership Report, this value is reset
    // to the group membership timer.  While this time remaining is non-zero, the
    // local router ignores any IGMPv2 Leave messages for this group that it
    // receives on this interface. The type is interface{} with range:
    // 0..4294967295.
    IgmpCacheVersion1HostTimer interface{}
}

func (igmpCacheEntry *IGMPSTDMIB_IgmpCacheTable_IgmpCacheEntry) GetEntityData() *types.CommonEntityData {
    igmpCacheEntry.EntityData.YFilter = igmpCacheEntry.YFilter
    igmpCacheEntry.EntityData.YangName = "igmpCacheEntry"
    igmpCacheEntry.EntityData.BundleName = "cisco_ios_xe"
    igmpCacheEntry.EntityData.ParentYangName = "igmpCacheTable"
    igmpCacheEntry.EntityData.SegmentPath = "igmpCacheEntry" + types.AddKeyToken(igmpCacheEntry.IgmpCacheAddress, "igmpCacheAddress") + types.AddKeyToken(igmpCacheEntry.IgmpCacheIfIndex, "igmpCacheIfIndex")
    igmpCacheEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpCacheEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpCacheEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpCacheEntry.EntityData.Children = types.NewOrderedMap()
    igmpCacheEntry.EntityData.Leafs = types.NewOrderedMap()
    igmpCacheEntry.EntityData.Leafs.Append("igmpCacheAddress", types.YLeaf{"IgmpCacheAddress", igmpCacheEntry.IgmpCacheAddress})
    igmpCacheEntry.EntityData.Leafs.Append("igmpCacheIfIndex", types.YLeaf{"IgmpCacheIfIndex", igmpCacheEntry.IgmpCacheIfIndex})
    igmpCacheEntry.EntityData.Leafs.Append("igmpCacheSelf", types.YLeaf{"IgmpCacheSelf", igmpCacheEntry.IgmpCacheSelf})
    igmpCacheEntry.EntityData.Leafs.Append("igmpCacheLastReporter", types.YLeaf{"IgmpCacheLastReporter", igmpCacheEntry.IgmpCacheLastReporter})
    igmpCacheEntry.EntityData.Leafs.Append("igmpCacheUpTime", types.YLeaf{"IgmpCacheUpTime", igmpCacheEntry.IgmpCacheUpTime})
    igmpCacheEntry.EntityData.Leafs.Append("igmpCacheExpiryTime", types.YLeaf{"IgmpCacheExpiryTime", igmpCacheEntry.IgmpCacheExpiryTime})
    igmpCacheEntry.EntityData.Leafs.Append("igmpCacheStatus", types.YLeaf{"IgmpCacheStatus", igmpCacheEntry.IgmpCacheStatus})
    igmpCacheEntry.EntityData.Leafs.Append("igmpCacheVersion1HostTimer", types.YLeaf{"IgmpCacheVersion1HostTimer", igmpCacheEntry.IgmpCacheVersion1HostTimer})

    igmpCacheEntry.EntityData.YListKeys = []string {"IgmpCacheAddress", "IgmpCacheIfIndex"}

    return &(igmpCacheEntry.EntityData)
}

