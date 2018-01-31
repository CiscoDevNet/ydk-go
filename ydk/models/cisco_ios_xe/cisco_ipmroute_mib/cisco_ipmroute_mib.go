// The MIB module for management of IP Multicast routing,
// but independent of the specific multicast routing protocol
// in use.
package cisco_ipmroute_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ipmroute_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IPMROUTE-MIB CISCO-IPMROUTE-MIB}", reflect.TypeOf(CISCOIPMROUTEMIB{}))
    ydk.RegisterEntity("CISCO-IPMROUTE-MIB:CISCO-IPMROUTE-MIB", reflect.TypeOf(CISCOIPMROUTEMIB{}))
}

// CISCOIPMROUTEMIB
type CISCOIPMROUTEMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ciscoipmroute CISCOIPMROUTEMIB_Ciscoipmroute

    // The (conceptual) table listing sets of IP Multicast heartbeat parameters. 
    // If no IP Multicast heartbeat is configured, this table would be empty.
    Ciscoipmrouteheartbeattable CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable
}

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetFilter() yfilter.YFilter { return cISCOIPMROUTEMIB.YFilter }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) SetFilter(yf yfilter.YFilter) { cISCOIPMROUTEMIB.YFilter = yf }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetGoName(yname string) string {
    if yname == "ciscoIpMRoute" { return "Ciscoipmroute" }
    if yname == "ciscoIpMRouteHeartBeatTable" { return "Ciscoipmrouteheartbeattable" }
    return ""
}

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetSegmentPath() string {
    return "CISCO-IPMROUTE-MIB:CISCO-IPMROUTE-MIB"
}

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoIpMRoute" {
        return &cISCOIPMROUTEMIB.Ciscoipmroute
    }
    if childYangName == "ciscoIpMRouteHeartBeatTable" {
        return &cISCOIPMROUTEMIB.Ciscoipmrouteheartbeattable
    }
    return nil
}

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoIpMRoute"] = &cISCOIPMROUTEMIB.Ciscoipmroute
    children["ciscoIpMRouteHeartBeatTable"] = &cISCOIPMROUTEMIB.Ciscoipmrouteheartbeattable
    return children
}

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetYangName() string { return "CISCO-IPMROUTE-MIB" }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) SetParent(parent types.Entity) { cISCOIPMROUTEMIB.parent = parent }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetParent() types.Entity { return cISCOIPMROUTEMIB.parent }

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetParentYangName() string { return "CISCO-IPMROUTE-MIB" }

// CISCOIPMROUTEMIB_Ciscoipmroute
type CISCOIPMROUTEMIB_Ciscoipmroute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maintains a count of the number of entries in the ipMRouteTable. The type
    // is interface{} with range: 0..4294967295.
    Ciscoipmroutenumberofentries interface{}
}

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetFilter() yfilter.YFilter { return ciscoipmroute.YFilter }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) SetFilter(yf yfilter.YFilter) { ciscoipmroute.YFilter = yf }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetGoName(yname string) string {
    if yname == "ciscoIpMRouteNumberOfEntries" { return "Ciscoipmroutenumberofentries" }
    return ""
}

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetSegmentPath() string {
    return "ciscoIpMRoute"
}

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoIpMRouteNumberOfEntries"] = ciscoipmroute.Ciscoipmroutenumberofentries
    return leafs
}

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetYangName() string { return "ciscoIpMRoute" }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) SetParent(parent types.Entity) { ciscoipmroute.parent = parent }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetParent() types.Entity { return ciscoipmroute.parent }

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetParentYangName() string { return "CISCO-IPMROUTE-MIB" }

// CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable
// The (conceptual) table listing sets of IP Multicast
// heartbeat parameters.  If no IP Multicast heartbeat is
// configured, this table would be empty.
type CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing a set of IP Multicast heartbeat
    // parameters. The type is slice of
    // CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry.
    Ciscoipmrouteheartbeatentry []CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry
}

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetFilter() yfilter.YFilter { return ciscoipmrouteheartbeattable.YFilter }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) SetFilter(yf yfilter.YFilter) { ciscoipmrouteheartbeattable.YFilter = yf }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetGoName(yname string) string {
    if yname == "ciscoIpMRouteHeartBeatEntry" { return "Ciscoipmrouteheartbeatentry" }
    return ""
}

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetSegmentPath() string {
    return "ciscoIpMRouteHeartBeatTable"
}

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoIpMRouteHeartBeatEntry" {
        for _, c := range ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry {
            if ciscoipmrouteheartbeattable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry{}
        ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry = append(ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry, child)
        return &ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry[len(ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry)-1]
    }
    return nil
}

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry {
        children[ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry[i].GetSegmentPath()] = &ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry[i]
    }
    return children
}

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetYangName() string { return "ciscoIpMRouteHeartBeatTable" }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) SetParent(parent types.Entity) { ciscoipmrouteheartbeattable.parent = parent }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetParent() types.Entity { return ciscoipmrouteheartbeattable.parent }

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetParentYangName() string { return "CISCO-IPMROUTE-MIB" }

// CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry
// An entry (conceptual row) representing a set of IP
// Multicast heartbeat parameters.
type CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Multicast group address used to receive heartbeat
    // packets. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ciscoipmrouteheartbeatgroupaddr interface{}

    // Source address of the last multicast heartbeat packet received. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ciscoipmrouteheartbeatsourceaddr interface{}

    // Number of seconds in which a Cisco multicast router expects a valid
    // heartBeat packet from a source.  This value must be a multiple of 10. The
    // type is interface{} with range: 10..3600. Units are seconds.
    Ciscoipmrouteheartbeatinterval interface{}

    // Number of ciscoIpMRouteHeartBeatInterval intervals a Cisco multicast router
    // waits before checking if expected number of heartbeat packets are received
    // or not. The type is interface{} with range: 1..100.
    Ciscoipmrouteheartbeatwindowsize interface{}

    // Number of time intervals where multicast packets were received in the last
    // ciscoIpMRouteHeartBeatWindowSize intervals. The type is interface{} with
    // range: 0..4294967295.
    Ciscoipmrouteheartbeatcount interface{}

    // The minimal number of heartbeat packets expected in the last
    // ciscoIpMRouteHeartBeatWindowSize intervals. If ciscoIpMRouteHeartBeatCount
    // falls below this value, an SNMP trap/notification, if configured, will be
    // sent to the NMS. The type is interface{} with range: 1..100.
    Ciscoipmrouteheartbeatminimum interface{}

    // The value of sysUpTime on the most recent occasion at which a missing IP
    // multicast heartbeat condition occured for the group address specified in
    // this entry.  If no such condition have occurred since the last
    // re-initialization of the local management subsystem, then this object
    // contains a zero value. The type is interface{} with range: 0..4294967295.
    Ciscoipmrouteheartbeatalerttime interface{}

    // This object is used to create a new row or delete an existing row in this
    // table. The type is RowStatus.
    Ciscoipmrouteheartbeatstatus interface{}
}

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetFilter() yfilter.YFilter { return ciscoipmrouteheartbeatentry.YFilter }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) SetFilter(yf yfilter.YFilter) { ciscoipmrouteheartbeatentry.YFilter = yf }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetGoName(yname string) string {
    if yname == "ciscoIpMRouteHeartBeatGroupAddr" { return "Ciscoipmrouteheartbeatgroupaddr" }
    if yname == "ciscoIpMRouteHeartBeatSourceAddr" { return "Ciscoipmrouteheartbeatsourceaddr" }
    if yname == "ciscoIpMRouteHeartBeatInterval" { return "Ciscoipmrouteheartbeatinterval" }
    if yname == "ciscoIpMRouteHeartBeatWindowSize" { return "Ciscoipmrouteheartbeatwindowsize" }
    if yname == "ciscoIpMRouteHeartBeatCount" { return "Ciscoipmrouteheartbeatcount" }
    if yname == "ciscoIpMRouteHeartBeatMinimum" { return "Ciscoipmrouteheartbeatminimum" }
    if yname == "ciscoIpMRouteHeartBeatAlertTime" { return "Ciscoipmrouteheartbeatalerttime" }
    if yname == "ciscoIpMRouteHeartBeatStatus" { return "Ciscoipmrouteheartbeatstatus" }
    return ""
}

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetSegmentPath() string {
    return "ciscoIpMRouteHeartBeatEntry" + "[ciscoIpMRouteHeartBeatGroupAddr='" + fmt.Sprintf("%v", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatgroupaddr) + "']"
}

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoIpMRouteHeartBeatGroupAddr"] = ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatgroupaddr
    leafs["ciscoIpMRouteHeartBeatSourceAddr"] = ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatsourceaddr
    leafs["ciscoIpMRouteHeartBeatInterval"] = ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatinterval
    leafs["ciscoIpMRouteHeartBeatWindowSize"] = ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatwindowsize
    leafs["ciscoIpMRouteHeartBeatCount"] = ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatcount
    leafs["ciscoIpMRouteHeartBeatMinimum"] = ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatminimum
    leafs["ciscoIpMRouteHeartBeatAlertTime"] = ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatalerttime
    leafs["ciscoIpMRouteHeartBeatStatus"] = ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatstatus
    return leafs
}

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetYangName() string { return "ciscoIpMRouteHeartBeatEntry" }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) SetParent(parent types.Entity) { ciscoipmrouteheartbeatentry.parent = parent }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetParent() types.Entity { return ciscoipmrouteheartbeatentry.parent }

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetParentYangName() string { return "ciscoIpMRouteHeartBeatTable" }

