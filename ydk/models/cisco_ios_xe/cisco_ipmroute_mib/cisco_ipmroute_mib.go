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

    
    CiscoIpMRoute CISCOIPMROUTEMIB_CiscoIpMRoute

    // The (conceptual) table listing sets of IP Multicast heartbeat parameters. 
    // If no IP Multicast heartbeat is configured, this table would be empty.
    CiscoIpMRouteHeartBeatTable CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable
}

func (cISCOIPMROUTEMIB *CISCOIPMROUTEMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPMROUTEMIB.EntityData.YFilter = cISCOIPMROUTEMIB.YFilter
    cISCOIPMROUTEMIB.EntityData.YangName = "CISCO-IPMROUTE-MIB"
    cISCOIPMROUTEMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPMROUTEMIB.EntityData.ParentYangName = "CISCO-IPMROUTE-MIB"
    cISCOIPMROUTEMIB.EntityData.SegmentPath = "CISCO-IPMROUTE-MIB:CISCO-IPMROUTE-MIB"
    cISCOIPMROUTEMIB.EntityData.AbsolutePath = cISCOIPMROUTEMIB.EntityData.SegmentPath
    cISCOIPMROUTEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPMROUTEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPMROUTEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPMROUTEMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPMROUTEMIB.EntityData.Children.Append("ciscoIpMRoute", types.YChild{"CiscoIpMRoute", &cISCOIPMROUTEMIB.CiscoIpMRoute})
    cISCOIPMROUTEMIB.EntityData.Children.Append("ciscoIpMRouteHeartBeatTable", types.YChild{"CiscoIpMRouteHeartBeatTable", &cISCOIPMROUTEMIB.CiscoIpMRouteHeartBeatTable})
    cISCOIPMROUTEMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPMROUTEMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPMROUTEMIB.EntityData)
}

// CISCOIPMROUTEMIB_CiscoIpMRoute
type CISCOIPMROUTEMIB_CiscoIpMRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maintains a count of the number of entries in the ipMRouteTable. The type
    // is interface{} with range: 0..4294967295.
    CiscoIpMRouteNumberOfEntries interface{}
}

func (ciscoIpMRoute *CISCOIPMROUTEMIB_CiscoIpMRoute) GetEntityData() *types.CommonEntityData {
    ciscoIpMRoute.EntityData.YFilter = ciscoIpMRoute.YFilter
    ciscoIpMRoute.EntityData.YangName = "ciscoIpMRoute"
    ciscoIpMRoute.EntityData.BundleName = "cisco_ios_xe"
    ciscoIpMRoute.EntityData.ParentYangName = "CISCO-IPMROUTE-MIB"
    ciscoIpMRoute.EntityData.SegmentPath = "ciscoIpMRoute"
    ciscoIpMRoute.EntityData.AbsolutePath = "CISCO-IPMROUTE-MIB:CISCO-IPMROUTE-MIB/" + ciscoIpMRoute.EntityData.SegmentPath
    ciscoIpMRoute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoIpMRoute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoIpMRoute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoIpMRoute.EntityData.Children = types.NewOrderedMap()
    ciscoIpMRoute.EntityData.Leafs = types.NewOrderedMap()
    ciscoIpMRoute.EntityData.Leafs.Append("ciscoIpMRouteNumberOfEntries", types.YLeaf{"CiscoIpMRouteNumberOfEntries", ciscoIpMRoute.CiscoIpMRouteNumberOfEntries})

    ciscoIpMRoute.EntityData.YListKeys = []string {}

    return &(ciscoIpMRoute.EntityData)
}

// CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable
// The (conceptual) table listing sets of IP Multicast
// heartbeat parameters.  If no IP Multicast heartbeat is
// configured, this table would be empty.
type CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) representing a set of IP Multicast heartbeat
    // parameters. The type is slice of
    // CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable_CiscoIpMRouteHeartBeatEntry.
    CiscoIpMRouteHeartBeatEntry []*CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable_CiscoIpMRouteHeartBeatEntry
}

func (ciscoIpMRouteHeartBeatTable *CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable) GetEntityData() *types.CommonEntityData {
    ciscoIpMRouteHeartBeatTable.EntityData.YFilter = ciscoIpMRouteHeartBeatTable.YFilter
    ciscoIpMRouteHeartBeatTable.EntityData.YangName = "ciscoIpMRouteHeartBeatTable"
    ciscoIpMRouteHeartBeatTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoIpMRouteHeartBeatTable.EntityData.ParentYangName = "CISCO-IPMROUTE-MIB"
    ciscoIpMRouteHeartBeatTable.EntityData.SegmentPath = "ciscoIpMRouteHeartBeatTable"
    ciscoIpMRouteHeartBeatTable.EntityData.AbsolutePath = "CISCO-IPMROUTE-MIB:CISCO-IPMROUTE-MIB/" + ciscoIpMRouteHeartBeatTable.EntityData.SegmentPath
    ciscoIpMRouteHeartBeatTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoIpMRouteHeartBeatTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoIpMRouteHeartBeatTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoIpMRouteHeartBeatTable.EntityData.Children = types.NewOrderedMap()
    ciscoIpMRouteHeartBeatTable.EntityData.Children.Append("ciscoIpMRouteHeartBeatEntry", types.YChild{"CiscoIpMRouteHeartBeatEntry", nil})
    for i := range ciscoIpMRouteHeartBeatTable.CiscoIpMRouteHeartBeatEntry {
        ciscoIpMRouteHeartBeatTable.EntityData.Children.Append(types.GetSegmentPath(ciscoIpMRouteHeartBeatTable.CiscoIpMRouteHeartBeatEntry[i]), types.YChild{"CiscoIpMRouteHeartBeatEntry", ciscoIpMRouteHeartBeatTable.CiscoIpMRouteHeartBeatEntry[i]})
    }
    ciscoIpMRouteHeartBeatTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoIpMRouteHeartBeatTable.EntityData.YListKeys = []string {}

    return &(ciscoIpMRouteHeartBeatTable.EntityData)
}

// CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable_CiscoIpMRouteHeartBeatEntry
// An entry (conceptual row) representing a set of IP
// Multicast heartbeat parameters.
type CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable_CiscoIpMRouteHeartBeatEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Multicast group address used to receive heartbeat
    // packets. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CiscoIpMRouteHeartBeatGroupAddr interface{}

    // Source address of the last multicast heartbeat packet received. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CiscoIpMRouteHeartBeatSourceAddr interface{}

    // Number of seconds in which a Cisco multicast router expects a valid
    // heartBeat packet from a source.  This value must be a multiple of 10. The
    // type is interface{} with range: 10..3600. Units are seconds.
    CiscoIpMRouteHeartBeatInterval interface{}

    // Number of ciscoIpMRouteHeartBeatInterval intervals a Cisco multicast router
    // waits before checking if expected number of heartbeat packets are received
    // or not. The type is interface{} with range: 1..100.
    CiscoIpMRouteHeartBeatWindowSize interface{}

    // Number of time intervals where multicast packets were received in the last
    // ciscoIpMRouteHeartBeatWindowSize intervals. The type is interface{} with
    // range: 0..4294967295.
    CiscoIpMRouteHeartBeatCount interface{}

    // The minimal number of heartbeat packets expected in the last
    // ciscoIpMRouteHeartBeatWindowSize intervals. If ciscoIpMRouteHeartBeatCount
    // falls below this value, an SNMP trap/notification, if configured, will be
    // sent to the NMS. The type is interface{} with range: 1..100.
    CiscoIpMRouteHeartBeatMinimum interface{}

    // The value of sysUpTime on the most recent occasion at which a missing IP
    // multicast heartbeat condition occured for the group address specified in
    // this entry.  If no such condition have occurred since the last
    // re-initialization of the local management subsystem, then this object
    // contains a zero value. The type is interface{} with range: 0..4294967295.
    CiscoIpMRouteHeartBeatAlertTime interface{}

    // This object is used to create a new row or delete an existing row in this
    // table. The type is RowStatus.
    CiscoIpMRouteHeartBeatStatus interface{}
}

func (ciscoIpMRouteHeartBeatEntry *CISCOIPMROUTEMIB_CiscoIpMRouteHeartBeatTable_CiscoIpMRouteHeartBeatEntry) GetEntityData() *types.CommonEntityData {
    ciscoIpMRouteHeartBeatEntry.EntityData.YFilter = ciscoIpMRouteHeartBeatEntry.YFilter
    ciscoIpMRouteHeartBeatEntry.EntityData.YangName = "ciscoIpMRouteHeartBeatEntry"
    ciscoIpMRouteHeartBeatEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoIpMRouteHeartBeatEntry.EntityData.ParentYangName = "ciscoIpMRouteHeartBeatTable"
    ciscoIpMRouteHeartBeatEntry.EntityData.SegmentPath = "ciscoIpMRouteHeartBeatEntry" + types.AddKeyToken(ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatGroupAddr, "ciscoIpMRouteHeartBeatGroupAddr")
    ciscoIpMRouteHeartBeatEntry.EntityData.AbsolutePath = "CISCO-IPMROUTE-MIB:CISCO-IPMROUTE-MIB/ciscoIpMRouteHeartBeatTable/" + ciscoIpMRouteHeartBeatEntry.EntityData.SegmentPath
    ciscoIpMRouteHeartBeatEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoIpMRouteHeartBeatEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoIpMRouteHeartBeatEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoIpMRouteHeartBeatEntry.EntityData.Children = types.NewOrderedMap()
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs.Append("ciscoIpMRouteHeartBeatGroupAddr", types.YLeaf{"CiscoIpMRouteHeartBeatGroupAddr", ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatGroupAddr})
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs.Append("ciscoIpMRouteHeartBeatSourceAddr", types.YLeaf{"CiscoIpMRouteHeartBeatSourceAddr", ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatSourceAddr})
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs.Append("ciscoIpMRouteHeartBeatInterval", types.YLeaf{"CiscoIpMRouteHeartBeatInterval", ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatInterval})
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs.Append("ciscoIpMRouteHeartBeatWindowSize", types.YLeaf{"CiscoIpMRouteHeartBeatWindowSize", ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatWindowSize})
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs.Append("ciscoIpMRouteHeartBeatCount", types.YLeaf{"CiscoIpMRouteHeartBeatCount", ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatCount})
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs.Append("ciscoIpMRouteHeartBeatMinimum", types.YLeaf{"CiscoIpMRouteHeartBeatMinimum", ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatMinimum})
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs.Append("ciscoIpMRouteHeartBeatAlertTime", types.YLeaf{"CiscoIpMRouteHeartBeatAlertTime", ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatAlertTime})
    ciscoIpMRouteHeartBeatEntry.EntityData.Leafs.Append("ciscoIpMRouteHeartBeatStatus", types.YLeaf{"CiscoIpMRouteHeartBeatStatus", ciscoIpMRouteHeartBeatEntry.CiscoIpMRouteHeartBeatStatus})

    ciscoIpMRouteHeartBeatEntry.EntityData.YListKeys = []string {"CiscoIpMRouteHeartBeatGroupAddr"}

    return &(ciscoIpMRouteHeartBeatEntry.EntityData)
}

