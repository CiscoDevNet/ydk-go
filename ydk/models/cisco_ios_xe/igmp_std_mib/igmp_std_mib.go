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
    Igmpinterfacetable IGMPSTDMIB_Igmpinterfacetable

    // The (conceptual) table listing the IP multicast groups for which there are
    // members on a particular interface.
    Igmpcachetable IGMPSTDMIB_Igmpcachetable
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

    iGMPSTDMIB.EntityData.Children = make(map[string]types.YChild)
    iGMPSTDMIB.EntityData.Children["igmpInterfaceTable"] = types.YChild{"Igmpinterfacetable", &iGMPSTDMIB.Igmpinterfacetable}
    iGMPSTDMIB.EntityData.Children["igmpCacheTable"] = types.YChild{"Igmpcachetable", &iGMPSTDMIB.Igmpcachetable}
    iGMPSTDMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iGMPSTDMIB.EntityData)
}

// IGMPSTDMIB_Igmpinterfacetable
// The (conceptual) table listing the interfaces on which IGMP
// is enabled.
type IGMPSTDMIB_Igmpinterfacetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an interface on which IGMP is
    // enabled. The type is slice of
    // IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry.
    Igmpinterfaceentry []IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry
}

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetEntityData() *types.CommonEntityData {
    igmpinterfacetable.EntityData.YFilter = igmpinterfacetable.YFilter
    igmpinterfacetable.EntityData.YangName = "igmpInterfaceTable"
    igmpinterfacetable.EntityData.BundleName = "cisco_ios_xe"
    igmpinterfacetable.EntityData.ParentYangName = "IGMP-STD-MIB"
    igmpinterfacetable.EntityData.SegmentPath = "igmpInterfaceTable"
    igmpinterfacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpinterfacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpinterfacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpinterfacetable.EntityData.Children = make(map[string]types.YChild)
    igmpinterfacetable.EntityData.Children["igmpInterfaceEntry"] = types.YChild{"Igmpinterfaceentry", nil}
    for i := range igmpinterfacetable.Igmpinterfaceentry {
        igmpinterfacetable.EntityData.Children[types.GetSegmentPath(&igmpinterfacetable.Igmpinterfaceentry[i])] = types.YChild{"Igmpinterfaceentry", &igmpinterfacetable.Igmpinterfaceentry[i]}
    }
    igmpinterfacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(igmpinterfacetable.EntityData)
}

// IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry
// An entry (conceptual row) representing an interface on
// which IGMP is enabled.
type IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The ifIndex value of the interface for which IGMP
    // is enabled. The type is interface{} with range: 1..2147483647.
    Igmpinterfaceifindex interface{}

    // The frequency at which IGMP Host-Query packets are transmitted on this
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Igmpinterfacequeryinterval interface{}

    // The activation of a row enables IGMP on the interface.  The destruction of
    // a row disables IGMP on the interface. The type is RowStatus.
    Igmpinterfacestatus interface{}

    // The version of IGMP which is running on this interface. This object can be
    // used to configure a router capable of running either value.  For IGMP to
    // function correctly, all routers on a LAN must be configured to run the same
    // version of IGMP on that LAN. The type is interface{} with range:
    // 0..4294967295.
    Igmpinterfaceversion interface{}

    // The address of the IGMP Querier on the IP subnet to which      this
    // interface is attached. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Igmpinterfacequerier interface{}

    // The maximum query response time advertised in IGMPv2 queries on this
    // interface. The type is interface{} with range: 0..255. Units are tenths of
    // seconds.
    Igmpinterfacequerymaxresponsetime interface{}

    // The time since igmpInterfaceQuerier was last changed. The type is
    // interface{} with range: 0..4294967295.
    Igmpinterfacequerieruptime interface{}

    // The amount of time remaining before the Other Querier Present Timer
    // expires.  If the local system is the querier, the value of this object is
    // zero. The type is interface{} with range: 0..4294967295.
    Igmpinterfacequerierexpirytime interface{}

    // The time remaining until the host assumes that there are no IGMPv1 routers
    // present on the interface.  While this is non- zero, the host will reply to
    // all queries with version 1 membership reports. The type is interface{} with
    // range: 0..4294967295.
    Igmpinterfaceversion1Queriertimer interface{}

    // The number of queries received whose IGMP version does not match
    // igmpInterfaceVersion, over the lifetime of the row entry.  IGMP requires
    // that all routers on a LAN be configured to run the same version of IGMP. 
    // Thus, if any queries are received with the wrong version, this indicates a
    // configuration error. The type is interface{} with range: 0..4294967295.
    Igmpinterfacewrongversionqueries interface{}

    // The number of times a group membership has been added on this interface;
    // that is, the number of times an entry for this interface has been added to
    // the Cache Table.  This object gives an indication of the amount of IGMP
    // activity over the lifetime of the row entry. The type is interface{} with
    // range: 0..4294967295.
    Igmpinterfacejoins interface{}

    // Some devices implement a form of IGMP proxying whereby memberships learned
    // on the interface represented by this row, cause IGMP Host Membership
    // Reports to be sent on the interface whose ifIndex value is given by this
    // object.  Such a device would implement the igmpV2RouterMIBGroup only on its
    // router interfaces (those interfaces with non-zero
    // igmpInterfaceProxyIfIndex).  Typically, the value of this object is 0,
    // indicating that no proxying is being done. The type is interface{} with
    // range: 0..2147483647.
    Igmpinterfaceproxyifindex interface{}

    // The current number of entries for this interface in the Cache Table. The
    // type is interface{} with range: 0..4294967295.
    Igmpinterfacegroups interface{}

    // The Robustness Variable allows tuning for the expected packet loss on a
    // subnet.  If a subnet is expected to be lossy, the Robustness Variable may
    // be increased.  IGMP is robust to (Robustness Variable-1) packet losses. The
    // type is interface{} with range: 1..255.
    Igmpinterfacerobustness interface{}

    // The Last Member Query Interval is the Max Response Time inserted into
    // Group-Specific Queries sent in response to Leave Group messages, and is
    // also the amount of time between Group-Specific Query messages.  This value
    // may be tuned to modify the leave latency of the network.  A reduced value
    // results in reduced time to detect the loss of the last member of a group. 
    // The value of this object is irrelevant if igmpInterfaceVersion is 1. The
    // type is interface{} with range: 0..255. Units are tenths of seconds.
    Igmpinterfacelastmembqueryintvl interface{}
}

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetEntityData() *types.CommonEntityData {
    igmpinterfaceentry.EntityData.YFilter = igmpinterfaceentry.YFilter
    igmpinterfaceentry.EntityData.YangName = "igmpInterfaceEntry"
    igmpinterfaceentry.EntityData.BundleName = "cisco_ios_xe"
    igmpinterfaceentry.EntityData.ParentYangName = "igmpInterfaceTable"
    igmpinterfaceentry.EntityData.SegmentPath = "igmpInterfaceEntry" + "[igmpInterfaceIfIndex='" + fmt.Sprintf("%v", igmpinterfaceentry.Igmpinterfaceifindex) + "']"
    igmpinterfaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpinterfaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpinterfaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpinterfaceentry.EntityData.Children = make(map[string]types.YChild)
    igmpinterfaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceIfIndex"] = types.YLeaf{"Igmpinterfaceifindex", igmpinterfaceentry.Igmpinterfaceifindex}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceQueryInterval"] = types.YLeaf{"Igmpinterfacequeryinterval", igmpinterfaceentry.Igmpinterfacequeryinterval}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceStatus"] = types.YLeaf{"Igmpinterfacestatus", igmpinterfaceentry.Igmpinterfacestatus}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceVersion"] = types.YLeaf{"Igmpinterfaceversion", igmpinterfaceentry.Igmpinterfaceversion}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceQuerier"] = types.YLeaf{"Igmpinterfacequerier", igmpinterfaceentry.Igmpinterfacequerier}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceQueryMaxResponseTime"] = types.YLeaf{"Igmpinterfacequerymaxresponsetime", igmpinterfaceentry.Igmpinterfacequerymaxresponsetime}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceQuerierUpTime"] = types.YLeaf{"Igmpinterfacequerieruptime", igmpinterfaceentry.Igmpinterfacequerieruptime}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceQuerierExpiryTime"] = types.YLeaf{"Igmpinterfacequerierexpirytime", igmpinterfaceentry.Igmpinterfacequerierexpirytime}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceVersion1QuerierTimer"] = types.YLeaf{"Igmpinterfaceversion1Queriertimer", igmpinterfaceentry.Igmpinterfaceversion1Queriertimer}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceWrongVersionQueries"] = types.YLeaf{"Igmpinterfacewrongversionqueries", igmpinterfaceentry.Igmpinterfacewrongversionqueries}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceJoins"] = types.YLeaf{"Igmpinterfacejoins", igmpinterfaceentry.Igmpinterfacejoins}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceProxyIfIndex"] = types.YLeaf{"Igmpinterfaceproxyifindex", igmpinterfaceentry.Igmpinterfaceproxyifindex}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceGroups"] = types.YLeaf{"Igmpinterfacegroups", igmpinterfaceentry.Igmpinterfacegroups}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceRobustness"] = types.YLeaf{"Igmpinterfacerobustness", igmpinterfaceentry.Igmpinterfacerobustness}
    igmpinterfaceentry.EntityData.Leafs["igmpInterfaceLastMembQueryIntvl"] = types.YLeaf{"Igmpinterfacelastmembqueryintvl", igmpinterfaceentry.Igmpinterfacelastmembqueryintvl}
    return &(igmpinterfaceentry.EntityData)
}

// IGMPSTDMIB_Igmpcachetable
// The (conceptual) table listing the IP multicast groups for
// which there are members on a particular interface.
type IGMPSTDMIB_Igmpcachetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the igmpCacheTable. The type is slice of
    // IGMPSTDMIB_Igmpcachetable_Igmpcacheentry.
    Igmpcacheentry []IGMPSTDMIB_Igmpcachetable_Igmpcacheentry
}

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetEntityData() *types.CommonEntityData {
    igmpcachetable.EntityData.YFilter = igmpcachetable.YFilter
    igmpcachetable.EntityData.YangName = "igmpCacheTable"
    igmpcachetable.EntityData.BundleName = "cisco_ios_xe"
    igmpcachetable.EntityData.ParentYangName = "IGMP-STD-MIB"
    igmpcachetable.EntityData.SegmentPath = "igmpCacheTable"
    igmpcachetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpcachetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpcachetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpcachetable.EntityData.Children = make(map[string]types.YChild)
    igmpcachetable.EntityData.Children["igmpCacheEntry"] = types.YChild{"Igmpcacheentry", nil}
    for i := range igmpcachetable.Igmpcacheentry {
        igmpcachetable.EntityData.Children[types.GetSegmentPath(&igmpcachetable.Igmpcacheentry[i])] = types.YChild{"Igmpcacheentry", &igmpcachetable.Igmpcacheentry[i]}
    }
    igmpcachetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(igmpcachetable.EntityData)
}

// IGMPSTDMIB_Igmpcachetable_Igmpcacheentry
// An entry (conceptual row) in the igmpCacheTable.
type IGMPSTDMIB_Igmpcachetable_Igmpcacheentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address for which this
    // entry contains information. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Igmpcacheaddress interface{}

    // This attribute is a key. The interface for which this entry contains
    // information for an IP multicast group address. The type is interface{} with
    // range: 1..2147483647.
    Igmpcacheifindex interface{}

    // An indication of whether the local system is a member of this group address
    // on this interface. The type is bool.
    Igmpcacheself interface{}

    // The IP address of the source of the last membership report received for
    // this IP Multicast group address on this interface.  If no membership report
    // has been received, this object has the value 0.0.0.0. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Igmpcachelastreporter interface{}

    // The time elapsed since this entry was created. The type is interface{} with
    // range: 0..4294967295.
    Igmpcacheuptime interface{}

    // The minimum amount of time remaining before this entry will be aged out.  A
    // value of 0 indicates that the entry is only present because igmpCacheSelf
    // is true and that if the router left the group, this entry would be aged out
    // immediately. Note that some implementations may process membership reports
    // from the local system in the same way as reports from other hosts, so a
    // value of 0 is not required. The type is interface{} with range:
    // 0..4294967295.
    Igmpcacheexpirytime interface{}

    // The status of this entry. The type is RowStatus.
    Igmpcachestatus interface{}

    // The time remaining until the local router will assume that there are no
    // longer any IGMP version 1 members on the IP subnet attached to this
    // interface.  Upon hearing any IGMPv1 Membership Report, this value is reset
    // to the group membership timer.  While this time remaining is non-zero, the
    // local router ignores any IGMPv2 Leave messages for this group that it
    // receives on this interface. The type is interface{} with range:
    // 0..4294967295.
    Igmpcacheversion1Hosttimer interface{}
}

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetEntityData() *types.CommonEntityData {
    igmpcacheentry.EntityData.YFilter = igmpcacheentry.YFilter
    igmpcacheentry.EntityData.YangName = "igmpCacheEntry"
    igmpcacheentry.EntityData.BundleName = "cisco_ios_xe"
    igmpcacheentry.EntityData.ParentYangName = "igmpCacheTable"
    igmpcacheentry.EntityData.SegmentPath = "igmpCacheEntry" + "[igmpCacheAddress='" + fmt.Sprintf("%v", igmpcacheentry.Igmpcacheaddress) + "']" + "[igmpCacheIfIndex='" + fmt.Sprintf("%v", igmpcacheentry.Igmpcacheifindex) + "']"
    igmpcacheentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpcacheentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpcacheentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpcacheentry.EntityData.Children = make(map[string]types.YChild)
    igmpcacheentry.EntityData.Leafs = make(map[string]types.YLeaf)
    igmpcacheentry.EntityData.Leafs["igmpCacheAddress"] = types.YLeaf{"Igmpcacheaddress", igmpcacheentry.Igmpcacheaddress}
    igmpcacheentry.EntityData.Leafs["igmpCacheIfIndex"] = types.YLeaf{"Igmpcacheifindex", igmpcacheentry.Igmpcacheifindex}
    igmpcacheentry.EntityData.Leafs["igmpCacheSelf"] = types.YLeaf{"Igmpcacheself", igmpcacheentry.Igmpcacheself}
    igmpcacheentry.EntityData.Leafs["igmpCacheLastReporter"] = types.YLeaf{"Igmpcachelastreporter", igmpcacheentry.Igmpcachelastreporter}
    igmpcacheentry.EntityData.Leafs["igmpCacheUpTime"] = types.YLeaf{"Igmpcacheuptime", igmpcacheentry.Igmpcacheuptime}
    igmpcacheentry.EntityData.Leafs["igmpCacheExpiryTime"] = types.YLeaf{"Igmpcacheexpirytime", igmpcacheentry.Igmpcacheexpirytime}
    igmpcacheentry.EntityData.Leafs["igmpCacheStatus"] = types.YLeaf{"Igmpcachestatus", igmpcacheentry.Igmpcachestatus}
    igmpcacheentry.EntityData.Leafs["igmpCacheVersion1HostTimer"] = types.YLeaf{"Igmpcacheversion1Hosttimer", igmpcacheentry.Igmpcacheversion1Hosttimer}
    return &(igmpcacheentry.EntityData)
}

