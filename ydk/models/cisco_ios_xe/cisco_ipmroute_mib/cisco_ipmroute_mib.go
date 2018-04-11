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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Ciscoipmroute CISCOIPMROUTEMIB_Ciscoipmroute

    // The (conceptual) table listing sets of IP Multicast heartbeat parameters. 
    // If no IP Multicast heartbeat is configured, this table would be empty.
    Ciscoipmrouteheartbeattable CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable
}

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPMROUTEMIB.EntityData.YFilter = cISCOIPMROUTEMIB.YFilter
    cISCOIPMROUTEMIB.EntityData.YangName = "CISCO-IPMROUTE-MIB"
    cISCOIPMROUTEMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPMROUTEMIB.EntityData.ParentYangName = "CISCO-IPMROUTE-MIB"
    cISCOIPMROUTEMIB.EntityData.SegmentPath = "CISCO-IPMROUTE-MIB:CISCO-IPMROUTE-MIB"
    cISCOIPMROUTEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPMROUTEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPMROUTEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPMROUTEMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPMROUTEMIB.EntityData.Children["ciscoIpMRoute"] = types.YChild{"Ciscoipmroute", &cISCOIPMROUTEMIB.Ciscoipmroute}
    cISCOIPMROUTEMIB.EntityData.Children["ciscoIpMRouteHeartBeatTable"] = types.YChild{"Ciscoipmrouteheartbeattable", &cISCOIPMROUTEMIB.Ciscoipmrouteheartbeattable}
    cISCOIPMROUTEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPMROUTEMIB.EntityData)
}

// CISCOIPMROUTEMIB_Ciscoipmroute
type CISCOIPMROUTEMIB_Ciscoipmroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maintains a count of the number of entries in the ipMRouteTable. The type
    // is interface{} with range: 0..4294967295.
    Ciscoipmroutenumberofentries interface{}
}

func (ciscoipmroute *CISCOIPMROUTEMIB_Ciscoipmroute) GetEntityData() *types.CommonEntityData {
    ciscoipmroute.EntityData.YFilter = ciscoipmroute.YFilter
    ciscoipmroute.EntityData.YangName = "ciscoIpMRoute"
    ciscoipmroute.EntityData.BundleName = "cisco_ios_xe"
    ciscoipmroute.EntityData.ParentYangName = "CISCO-IPMROUTE-MIB"
    ciscoipmroute.EntityData.SegmentPath = "ciscoIpMRoute"
    ciscoipmroute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoipmroute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoipmroute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoipmroute.EntityData.Children = make(map[string]types.YChild)
    ciscoipmroute.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoipmroute.EntityData.Leafs["ciscoIpMRouteNumberOfEntries"] = types.YLeaf{"Ciscoipmroutenumberofentries", ciscoipmroute.Ciscoipmroutenumberofentries}
    return &(ciscoipmroute.EntityData)
}

// CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable
// The (conceptual) table listing sets of IP Multicast
// heartbeat parameters.  If no IP Multicast heartbeat is
// configured, this table would be empty.
type CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing a set of IP Multicast heartbeat
    // parameters. The type is slice of
    // CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry.
    Ciscoipmrouteheartbeatentry []CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry
}

func (ciscoipmrouteheartbeattable *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable) GetEntityData() *types.CommonEntityData {
    ciscoipmrouteheartbeattable.EntityData.YFilter = ciscoipmrouteheartbeattable.YFilter
    ciscoipmrouteheartbeattable.EntityData.YangName = "ciscoIpMRouteHeartBeatTable"
    ciscoipmrouteheartbeattable.EntityData.BundleName = "cisco_ios_xe"
    ciscoipmrouteheartbeattable.EntityData.ParentYangName = "CISCO-IPMROUTE-MIB"
    ciscoipmrouteheartbeattable.EntityData.SegmentPath = "ciscoIpMRouteHeartBeatTable"
    ciscoipmrouteheartbeattable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoipmrouteheartbeattable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoipmrouteheartbeattable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoipmrouteheartbeattable.EntityData.Children = make(map[string]types.YChild)
    ciscoipmrouteheartbeattable.EntityData.Children["ciscoIpMRouteHeartBeatEntry"] = types.YChild{"Ciscoipmrouteheartbeatentry", nil}
    for i := range ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry {
        ciscoipmrouteheartbeattable.EntityData.Children[types.GetSegmentPath(&ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry[i])] = types.YChild{"Ciscoipmrouteheartbeatentry", &ciscoipmrouteheartbeattable.Ciscoipmrouteheartbeatentry[i]}
    }
    ciscoipmrouteheartbeattable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoipmrouteheartbeattable.EntityData)
}

// CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry
// An entry (conceptual row) representing a set of IP
// Multicast heartbeat parameters.
type CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Multicast group address used to receive heartbeat
    // packets. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ciscoipmrouteheartbeatgroupaddr interface{}

    // Source address of the last multicast heartbeat packet received. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ciscoipmrouteheartbeatentry *CISCOIPMROUTEMIB_Ciscoipmrouteheartbeattable_Ciscoipmrouteheartbeatentry) GetEntityData() *types.CommonEntityData {
    ciscoipmrouteheartbeatentry.EntityData.YFilter = ciscoipmrouteheartbeatentry.YFilter
    ciscoipmrouteheartbeatentry.EntityData.YangName = "ciscoIpMRouteHeartBeatEntry"
    ciscoipmrouteheartbeatentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoipmrouteheartbeatentry.EntityData.ParentYangName = "ciscoIpMRouteHeartBeatTable"
    ciscoipmrouteheartbeatentry.EntityData.SegmentPath = "ciscoIpMRouteHeartBeatEntry" + "[ciscoIpMRouteHeartBeatGroupAddr='" + fmt.Sprintf("%v", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatgroupaddr) + "']"
    ciscoipmrouteheartbeatentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoipmrouteheartbeatentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoipmrouteheartbeatentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoipmrouteheartbeatentry.EntityData.Children = make(map[string]types.YChild)
    ciscoipmrouteheartbeatentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoipmrouteheartbeatentry.EntityData.Leafs["ciscoIpMRouteHeartBeatGroupAddr"] = types.YLeaf{"Ciscoipmrouteheartbeatgroupaddr", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatgroupaddr}
    ciscoipmrouteheartbeatentry.EntityData.Leafs["ciscoIpMRouteHeartBeatSourceAddr"] = types.YLeaf{"Ciscoipmrouteheartbeatsourceaddr", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatsourceaddr}
    ciscoipmrouteheartbeatentry.EntityData.Leafs["ciscoIpMRouteHeartBeatInterval"] = types.YLeaf{"Ciscoipmrouteheartbeatinterval", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatinterval}
    ciscoipmrouteheartbeatentry.EntityData.Leafs["ciscoIpMRouteHeartBeatWindowSize"] = types.YLeaf{"Ciscoipmrouteheartbeatwindowsize", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatwindowsize}
    ciscoipmrouteheartbeatentry.EntityData.Leafs["ciscoIpMRouteHeartBeatCount"] = types.YLeaf{"Ciscoipmrouteheartbeatcount", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatcount}
    ciscoipmrouteheartbeatentry.EntityData.Leafs["ciscoIpMRouteHeartBeatMinimum"] = types.YLeaf{"Ciscoipmrouteheartbeatminimum", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatminimum}
    ciscoipmrouteheartbeatentry.EntityData.Leafs["ciscoIpMRouteHeartBeatAlertTime"] = types.YLeaf{"Ciscoipmrouteheartbeatalerttime", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatalerttime}
    ciscoipmrouteheartbeatentry.EntityData.Leafs["ciscoIpMRouteHeartBeatStatus"] = types.YLeaf{"Ciscoipmrouteheartbeatstatus", ciscoipmrouteheartbeatentry.Ciscoipmrouteheartbeatstatus}
    return &(ciscoipmrouteheartbeatentry.EntityData)
}

