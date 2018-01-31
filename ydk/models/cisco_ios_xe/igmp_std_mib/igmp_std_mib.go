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
    parent types.Entity
    YFilter yfilter.YFilter

    // The (conceptual) table listing the interfaces on which IGMP is enabled.
    Igmpinterfacetable IGMPSTDMIB_Igmpinterfacetable

    // The (conceptual) table listing the IP multicast groups for which there are
    // members on a particular interface.
    Igmpcachetable IGMPSTDMIB_Igmpcachetable
}

func (iGMPSTDMIB *IGMPSTDMIB) GetFilter() yfilter.YFilter { return iGMPSTDMIB.YFilter }

func (iGMPSTDMIB *IGMPSTDMIB) SetFilter(yf yfilter.YFilter) { iGMPSTDMIB.YFilter = yf }

func (iGMPSTDMIB *IGMPSTDMIB) GetGoName(yname string) string {
    if yname == "igmpInterfaceTable" { return "Igmpinterfacetable" }
    if yname == "igmpCacheTable" { return "Igmpcachetable" }
    return ""
}

func (iGMPSTDMIB *IGMPSTDMIB) GetSegmentPath() string {
    return "IGMP-STD-MIB:IGMP-STD-MIB"
}

func (iGMPSTDMIB *IGMPSTDMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igmpInterfaceTable" {
        return &iGMPSTDMIB.Igmpinterfacetable
    }
    if childYangName == "igmpCacheTable" {
        return &iGMPSTDMIB.Igmpcachetable
    }
    return nil
}

func (iGMPSTDMIB *IGMPSTDMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igmpInterfaceTable"] = &iGMPSTDMIB.Igmpinterfacetable
    children["igmpCacheTable"] = &iGMPSTDMIB.Igmpcachetable
    return children
}

func (iGMPSTDMIB *IGMPSTDMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (iGMPSTDMIB *IGMPSTDMIB) GetBundleName() string { return "cisco_ios_xe" }

func (iGMPSTDMIB *IGMPSTDMIB) GetYangName() string { return "IGMP-STD-MIB" }

func (iGMPSTDMIB *IGMPSTDMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iGMPSTDMIB *IGMPSTDMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iGMPSTDMIB *IGMPSTDMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iGMPSTDMIB *IGMPSTDMIB) SetParent(parent types.Entity) { iGMPSTDMIB.parent = parent }

func (iGMPSTDMIB *IGMPSTDMIB) GetParent() types.Entity { return iGMPSTDMIB.parent }

func (iGMPSTDMIB *IGMPSTDMIB) GetParentYangName() string { return "IGMP-STD-MIB" }

// IGMPSTDMIB_Igmpinterfacetable
// The (conceptual) table listing the interfaces on which IGMP
// is enabled.
type IGMPSTDMIB_Igmpinterfacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing an interface on which IGMP is
    // enabled. The type is slice of
    // IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry.
    Igmpinterfaceentry []IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry
}

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetFilter() yfilter.YFilter { return igmpinterfacetable.YFilter }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) SetFilter(yf yfilter.YFilter) { igmpinterfacetable.YFilter = yf }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetGoName(yname string) string {
    if yname == "igmpInterfaceEntry" { return "Igmpinterfaceentry" }
    return ""
}

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetSegmentPath() string {
    return "igmpInterfaceTable"
}

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igmpInterfaceEntry" {
        for _, c := range igmpinterfacetable.Igmpinterfaceentry {
            if igmpinterfacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry{}
        igmpinterfacetable.Igmpinterfaceentry = append(igmpinterfacetable.Igmpinterfaceentry, child)
        return &igmpinterfacetable.Igmpinterfaceentry[len(igmpinterfacetable.Igmpinterfaceentry)-1]
    }
    return nil
}

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range igmpinterfacetable.Igmpinterfaceentry {
        children[igmpinterfacetable.Igmpinterfaceentry[i].GetSegmentPath()] = &igmpinterfacetable.Igmpinterfaceentry[i]
    }
    return children
}

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetBundleName() string { return "cisco_ios_xe" }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetYangName() string { return "igmpInterfaceTable" }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) SetParent(parent types.Entity) { igmpinterfacetable.parent = parent }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetParent() types.Entity { return igmpinterfacetable.parent }

func (igmpinterfacetable *IGMPSTDMIB_Igmpinterfacetable) GetParentYangName() string { return "IGMP-STD-MIB" }

// IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry
// An entry (conceptual row) representing an interface on
// which IGMP is enabled.
type IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetFilter() yfilter.YFilter { return igmpinterfaceentry.YFilter }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) SetFilter(yf yfilter.YFilter) { igmpinterfaceentry.YFilter = yf }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetGoName(yname string) string {
    if yname == "igmpInterfaceIfIndex" { return "Igmpinterfaceifindex" }
    if yname == "igmpInterfaceQueryInterval" { return "Igmpinterfacequeryinterval" }
    if yname == "igmpInterfaceStatus" { return "Igmpinterfacestatus" }
    if yname == "igmpInterfaceVersion" { return "Igmpinterfaceversion" }
    if yname == "igmpInterfaceQuerier" { return "Igmpinterfacequerier" }
    if yname == "igmpInterfaceQueryMaxResponseTime" { return "Igmpinterfacequerymaxresponsetime" }
    if yname == "igmpInterfaceQuerierUpTime" { return "Igmpinterfacequerieruptime" }
    if yname == "igmpInterfaceQuerierExpiryTime" { return "Igmpinterfacequerierexpirytime" }
    if yname == "igmpInterfaceVersion1QuerierTimer" { return "Igmpinterfaceversion1Queriertimer" }
    if yname == "igmpInterfaceWrongVersionQueries" { return "Igmpinterfacewrongversionqueries" }
    if yname == "igmpInterfaceJoins" { return "Igmpinterfacejoins" }
    if yname == "igmpInterfaceProxyIfIndex" { return "Igmpinterfaceproxyifindex" }
    if yname == "igmpInterfaceGroups" { return "Igmpinterfacegroups" }
    if yname == "igmpInterfaceRobustness" { return "Igmpinterfacerobustness" }
    if yname == "igmpInterfaceLastMembQueryIntvl" { return "Igmpinterfacelastmembqueryintvl" }
    return ""
}

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetSegmentPath() string {
    return "igmpInterfaceEntry" + "[igmpInterfaceIfIndex='" + fmt.Sprintf("%v", igmpinterfaceentry.Igmpinterfaceifindex) + "']"
}

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igmpInterfaceIfIndex"] = igmpinterfaceentry.Igmpinterfaceifindex
    leafs["igmpInterfaceQueryInterval"] = igmpinterfaceentry.Igmpinterfacequeryinterval
    leafs["igmpInterfaceStatus"] = igmpinterfaceentry.Igmpinterfacestatus
    leafs["igmpInterfaceVersion"] = igmpinterfaceentry.Igmpinterfaceversion
    leafs["igmpInterfaceQuerier"] = igmpinterfaceentry.Igmpinterfacequerier
    leafs["igmpInterfaceQueryMaxResponseTime"] = igmpinterfaceentry.Igmpinterfacequerymaxresponsetime
    leafs["igmpInterfaceQuerierUpTime"] = igmpinterfaceentry.Igmpinterfacequerieruptime
    leafs["igmpInterfaceQuerierExpiryTime"] = igmpinterfaceentry.Igmpinterfacequerierexpirytime
    leafs["igmpInterfaceVersion1QuerierTimer"] = igmpinterfaceentry.Igmpinterfaceversion1Queriertimer
    leafs["igmpInterfaceWrongVersionQueries"] = igmpinterfaceentry.Igmpinterfacewrongversionqueries
    leafs["igmpInterfaceJoins"] = igmpinterfaceentry.Igmpinterfacejoins
    leafs["igmpInterfaceProxyIfIndex"] = igmpinterfaceentry.Igmpinterfaceproxyifindex
    leafs["igmpInterfaceGroups"] = igmpinterfaceentry.Igmpinterfacegroups
    leafs["igmpInterfaceRobustness"] = igmpinterfaceentry.Igmpinterfacerobustness
    leafs["igmpInterfaceLastMembQueryIntvl"] = igmpinterfaceentry.Igmpinterfacelastmembqueryintvl
    return leafs
}

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetYangName() string { return "igmpInterfaceEntry" }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) SetParent(parent types.Entity) { igmpinterfaceentry.parent = parent }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetParent() types.Entity { return igmpinterfaceentry.parent }

func (igmpinterfaceentry *IGMPSTDMIB_Igmpinterfacetable_Igmpinterfaceentry) GetParentYangName() string { return "igmpInterfaceTable" }

// IGMPSTDMIB_Igmpcachetable
// The (conceptual) table listing the IP multicast groups for
// which there are members on a particular interface.
type IGMPSTDMIB_Igmpcachetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the igmpCacheTable. The type is slice of
    // IGMPSTDMIB_Igmpcachetable_Igmpcacheentry.
    Igmpcacheentry []IGMPSTDMIB_Igmpcachetable_Igmpcacheentry
}

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetFilter() yfilter.YFilter { return igmpcachetable.YFilter }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) SetFilter(yf yfilter.YFilter) { igmpcachetable.YFilter = yf }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetGoName(yname string) string {
    if yname == "igmpCacheEntry" { return "Igmpcacheentry" }
    return ""
}

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetSegmentPath() string {
    return "igmpCacheTable"
}

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igmpCacheEntry" {
        for _, c := range igmpcachetable.Igmpcacheentry {
            if igmpcachetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := IGMPSTDMIB_Igmpcachetable_Igmpcacheentry{}
        igmpcachetable.Igmpcacheentry = append(igmpcachetable.Igmpcacheentry, child)
        return &igmpcachetable.Igmpcacheentry[len(igmpcachetable.Igmpcacheentry)-1]
    }
    return nil
}

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range igmpcachetable.Igmpcacheentry {
        children[igmpcachetable.Igmpcacheentry[i].GetSegmentPath()] = &igmpcachetable.Igmpcacheentry[i]
    }
    return children
}

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetBundleName() string { return "cisco_ios_xe" }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetYangName() string { return "igmpCacheTable" }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) SetParent(parent types.Entity) { igmpcachetable.parent = parent }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetParent() types.Entity { return igmpcachetable.parent }

func (igmpcachetable *IGMPSTDMIB_Igmpcachetable) GetParentYangName() string { return "IGMP-STD-MIB" }

// IGMPSTDMIB_Igmpcachetable_Igmpcacheentry
// An entry (conceptual row) in the igmpCacheTable.
type IGMPSTDMIB_Igmpcachetable_Igmpcacheentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The IP multicast group address for which this
    // entry contains information. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetFilter() yfilter.YFilter { return igmpcacheentry.YFilter }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) SetFilter(yf yfilter.YFilter) { igmpcacheentry.YFilter = yf }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetGoName(yname string) string {
    if yname == "igmpCacheAddress" { return "Igmpcacheaddress" }
    if yname == "igmpCacheIfIndex" { return "Igmpcacheifindex" }
    if yname == "igmpCacheSelf" { return "Igmpcacheself" }
    if yname == "igmpCacheLastReporter" { return "Igmpcachelastreporter" }
    if yname == "igmpCacheUpTime" { return "Igmpcacheuptime" }
    if yname == "igmpCacheExpiryTime" { return "Igmpcacheexpirytime" }
    if yname == "igmpCacheStatus" { return "Igmpcachestatus" }
    if yname == "igmpCacheVersion1HostTimer" { return "Igmpcacheversion1Hosttimer" }
    return ""
}

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetSegmentPath() string {
    return "igmpCacheEntry" + "[igmpCacheAddress='" + fmt.Sprintf("%v", igmpcacheentry.Igmpcacheaddress) + "']" + "[igmpCacheIfIndex='" + fmt.Sprintf("%v", igmpcacheentry.Igmpcacheifindex) + "']"
}

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igmpCacheAddress"] = igmpcacheentry.Igmpcacheaddress
    leafs["igmpCacheIfIndex"] = igmpcacheentry.Igmpcacheifindex
    leafs["igmpCacheSelf"] = igmpcacheentry.Igmpcacheself
    leafs["igmpCacheLastReporter"] = igmpcacheentry.Igmpcachelastreporter
    leafs["igmpCacheUpTime"] = igmpcacheentry.Igmpcacheuptime
    leafs["igmpCacheExpiryTime"] = igmpcacheentry.Igmpcacheexpirytime
    leafs["igmpCacheStatus"] = igmpcacheentry.Igmpcachestatus
    leafs["igmpCacheVersion1HostTimer"] = igmpcacheentry.Igmpcacheversion1Hosttimer
    return leafs
}

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetBundleName() string { return "cisco_ios_xe" }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetYangName() string { return "igmpCacheEntry" }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) SetParent(parent types.Entity) { igmpcacheentry.parent = parent }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetParent() types.Entity { return igmpcacheentry.parent }

func (igmpcacheentry *IGMPSTDMIB_Igmpcachetable_Igmpcacheentry) GetParentYangName() string { return "igmpCacheTable" }

