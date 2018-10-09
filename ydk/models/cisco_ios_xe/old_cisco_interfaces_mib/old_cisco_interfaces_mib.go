package old_cisco_interfaces_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package old_cisco_interfaces_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:OLD-CISCO-INTERFACES-MIB OLD-CISCO-INTERFACES-MIB}", reflect.TypeOf(OLDCISCOINTERFACESMIB{}))
    ydk.RegisterEntity("OLD-CISCO-INTERFACES-MIB:OLD-CISCO-INTERFACES-MIB", reflect.TypeOf(OLDCISCOINTERFACESMIB{}))
}

// OLDCISCOINTERFACESMIB
type OLDCISCOINTERFACESMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of interface entries.
    LifTable OLDCISCOINTERFACESMIB_LifTable

    // A list of card entries for 4T, HSSI, Mx serial or FSIP.
    LFSIPTable OLDCISCOINTERFACESMIB_LFSIPTable
}

func (oLDCISCOINTERFACESMIB *OLDCISCOINTERFACESMIB) GetEntityData() *types.CommonEntityData {
    oLDCISCOINTERFACESMIB.EntityData.YFilter = oLDCISCOINTERFACESMIB.YFilter
    oLDCISCOINTERFACESMIB.EntityData.YangName = "OLD-CISCO-INTERFACES-MIB"
    oLDCISCOINTERFACESMIB.EntityData.BundleName = "cisco_ios_xe"
    oLDCISCOINTERFACESMIB.EntityData.ParentYangName = "OLD-CISCO-INTERFACES-MIB"
    oLDCISCOINTERFACESMIB.EntityData.SegmentPath = "OLD-CISCO-INTERFACES-MIB:OLD-CISCO-INTERFACES-MIB"
    oLDCISCOINTERFACESMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    oLDCISCOINTERFACESMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    oLDCISCOINTERFACESMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    oLDCISCOINTERFACESMIB.EntityData.Children = types.NewOrderedMap()
    oLDCISCOINTERFACESMIB.EntityData.Children.Append("lifTable", types.YChild{"LifTable", &oLDCISCOINTERFACESMIB.LifTable})
    oLDCISCOINTERFACESMIB.EntityData.Children.Append("lFSIPTable", types.YChild{"LFSIPTable", &oLDCISCOINTERFACESMIB.LFSIPTable})
    oLDCISCOINTERFACESMIB.EntityData.Leafs = types.NewOrderedMap()

    oLDCISCOINTERFACESMIB.EntityData.YListKeys = []string {}

    return &(oLDCISCOINTERFACESMIB.EntityData)
}

// OLDCISCOINTERFACESMIB_LifTable
// A list of interface entries.
type OLDCISCOINTERFACESMIB_LifTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of additional objects in the cisco interface. The type is
    // slice of OLDCISCOINTERFACESMIB_LifTable_LifEntry.
    LifEntry []*OLDCISCOINTERFACESMIB_LifTable_LifEntry
}

func (lifTable *OLDCISCOINTERFACESMIB_LifTable) GetEntityData() *types.CommonEntityData {
    lifTable.EntityData.YFilter = lifTable.YFilter
    lifTable.EntityData.YangName = "lifTable"
    lifTable.EntityData.BundleName = "cisco_ios_xe"
    lifTable.EntityData.ParentYangName = "OLD-CISCO-INTERFACES-MIB"
    lifTable.EntityData.SegmentPath = "lifTable"
    lifTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lifTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lifTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lifTable.EntityData.Children = types.NewOrderedMap()
    lifTable.EntityData.Children.Append("lifEntry", types.YChild{"LifEntry", nil})
    for i := range lifTable.LifEntry {
        lifTable.EntityData.Children.Append(types.GetSegmentPath(lifTable.LifEntry[i]), types.YChild{"LifEntry", lifTable.LifEntry[i]})
    }
    lifTable.EntityData.Leafs = types.NewOrderedMap()

    lifTable.EntityData.YListKeys = []string {}

    return &(lifTable.EntityData)
}

// OLDCISCOINTERFACESMIB_LifTable_LifEntry
// A collection of additional objects in the
// cisco interface.
type OLDCISCOINTERFACESMIB_LifTable_LifEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range:
    // -2147483648..2147483647. Refers to
    // rfc1213_mib.RFC1213MIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // Returns the type of interface. The type is string.
    LocIfHardType interface{}

    // Boolean whether interface line protocol is up or not. The type is
    // interface{} with range: -2147483648..2147483647.
    LocIfLineProt interface{}

    // Elapsed time in milliseconds since last line protocol input packet was
    // received. The type is interface{} with range: -2147483648..2147483647.
    LocIfLastIn interface{}

    // Elapsed time in milliseconds since last line protocol output packet was
    // transmitted. The type is interface{} with range: -2147483648..2147483647.
    LocIfLastOut interface{}

    // Elapsed time in milliseconds since last line protocol output packet could
    // not be successfully transmitted. The type is interface{} with range:
    // -2147483648..2147483647.
    LocIfLastOutHang interface{}

    // Five minute exponentially-decayed moving average of input bits per second.
    // The type is interface{} with range: -2147483648..2147483647.
    LocIfInBitsSec interface{}

    // Five minute exponentially-decayed moving average of input packets per
    // second. The type is interface{} with range: -2147483648..2147483647.
    LocIfInPktsSec interface{}

    // Five minute exponentially-decayed moving average of output bits per second.
    // The type is interface{} with range: -2147483648..2147483647.
    LocIfOutBitsSec interface{}

    // Five minute exponentially-decayed moving average of output packets per
    // second. The type is interface{} with range: -2147483648..2147483647.
    LocIfOutPktsSec interface{}

    // Number of packets input which were smaller then the allowable physical
    // media permitted. The type is interface{} with range:
    // -2147483648..2147483647.
    LocIfInRunts interface{}

    // Number of input packets which were larger then the physical media
    // permitted. The type is interface{} with range: -2147483648..2147483647.
    LocIfInGiants interface{}

    // Number of input packets which had cyclic redundancy checksum errors. The
    // type is interface{} with range: -2147483648..2147483647.
    LocIfInCRC interface{}

    // Number of input packet which were misaligned. The type is interface{} with
    // range: -2147483648..2147483647.
    LocIfInFrame interface{}

    // Count of input which arrived too quickly for the to hardware receive. The
    // type is interface{} with range: -2147483648..2147483647.
    LocIfInOverrun interface{}

    // Number of input packets which were simply ignored by this interface. The
    // type is interface{} with range: -2147483648..2147483647.
    LocIfInIgnored interface{}

    // Number of input packets which were aborted. The type is interface{} with
    // range: -2147483648..2147483647.
    LocIfInAbort interface{}

    // Number of times the interface internally reset. The type is interface{}
    // with range: -2147483648..2147483647.
    LocIfResets interface{}

    // Number of times interface needed to be completely restarted. The type is
    // interface{} with range: -2147483648..2147483647.
    LocIfRestarts interface{}

    // Boolean whether keepalives are enabled on this interface. The type is
    // interface{} with range: -2147483648..2147483647.
    LocIfKeep interface{}

    // Reason for interface last status change. The type is string.
    LocIfReason interface{}

    // Number of times interface saw the carrier signal transition. The type is
    // interface{} with range: -2147483648..2147483647.
    LocIfCarTrans interface{}

    // The reliability of the interface. Used by IGRP. The type is interface{}
    // with range: -2147483648..2147483647.
    LocIfReliab interface{}

    // The amount of delay in microseconds of the interface. Used by IGRP. The
    // type is interface{} with range: -2147483648..2147483647.
    LocIfDelay interface{}

    // The loading factor of the interface. Used by IGRP. The type is interface{}
    // with range: -2147483648..2147483647.
    LocIfLoad interface{}

    // The number of output collisions detected on this interface. The type is
    // interface{} with range: -2147483648..2147483647.
    LocIfCollisions interface{}

    // The number of packets dropped because the input queue was full. The type is
    // interface{} with range: -2147483648..2147483647.
    LocIfInputQueueDrops interface{}

    // The number of packets dropped because the output queue was full. The type
    // is interface{} with range: -2147483648..2147483647.
    LocIfOutputQueueDrops interface{}

    // User configurable interface description. The type is string.
    LocIfDescr interface{}

    // Packet count for Inbound traffic routed with slow switching. The type is
    // interface{} with range: 0..4294967295.
    LocIfSlowInPkts interface{}

    // Packet count for Outbound traffic routed with slow switching. The type is
    // interface{} with range: 0..4294967295.
    LocIfSlowOutPkts interface{}

    // Octet count for Inbound traffic routed with slow switching. The type is
    // interface{} with range: 0..4294967295.
    LocIfSlowInOctets interface{}

    // Octet count for Outbound traffic routed with slow switching. The type is
    // interface{} with range: 0..4294967295.
    LocIfSlowOutOctets interface{}

    // Packet count for Inbound traffic routed with fast switching. The type is
    // interface{} with range: 0..4294967295.
    LocIfFastInPkts interface{}

    // Packet count for Outbound traffic routed with fast switching. The type is
    // interface{} with range: 0..4294967295.
    LocIfFastOutPkts interface{}

    // Octet count for Inbound traffic routed with fast switching. The type is
    // interface{} with range: 0..4294967295.
    LocIfFastInOctets interface{}

    // Octet count for Outbound traffic routed with fast switching. The type is
    // interface{} with range: 0..4294967295.
    LocIfFastOutOctets interface{}

    // Other protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfotherInPkts interface{}

    // Other protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfotherOutPkts interface{}

    // Other protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfotherInOctets interface{}

    // Other protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfotherOutOctets interface{}

    // ip protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfipInPkts interface{}

    // ip protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfipOutPkts interface{}

    // ip protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfipInOctets interface{}

    // ip protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfipOutOctets interface{}

    // Decnet protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfdecnetInPkts interface{}

    // Decnet protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfdecnetOutPkts interface{}

    // Decnet protocol input byte count. The type is interface{} with range:
    // 0..4294967295.
    LocIfdecnetInOctets interface{}

    // Decnet protocol output byte count. The type is interface{} with range:
    // 0..4294967295.
    LocIfdecnetOutOctets interface{}

    // XNS protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfxnsInPkts interface{}

    // XNS protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfxnsOutPkts interface{}

    // XNS protocol input byte count. The type is interface{} with range:
    // 0..4294967295.
    LocIfxnsInOctets interface{}

    // XNS protocol output byte count. The type is interface{} with range:
    // 0..4294967295.
    LocIfxnsOutOctets interface{}

    // CLNS protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfclnsInPkts interface{}

    // CLNS protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfclnsOutPkts interface{}

    // CLNS protocol input byte count. The type is interface{} with range:
    // 0..4294967295.
    LocIfclnsInOctets interface{}

    // CLNS protocol output byte count. The type is interface{} with range:
    // 0..4294967295.
    LocIfclnsOutOctets interface{}

    // Appletalk protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfappletalkInPkts interface{}

    // Appletalk protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfappletalkOutPkts interface{}

    // Appletalk protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfappletalkInOctets interface{}

    // Appletalk protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfappletalkOutOctets interface{}

    // Novell protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfnovellInPkts interface{}

    // Novell protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfnovellOutPkts interface{}

    // Novell protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfnovellInOctets interface{}

    // Novell protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfnovellOutOctets interface{}

    // Apollo protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfapolloInPkts interface{}

    // Apollo protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfapolloOutPkts interface{}

    // Apollo protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfapolloInOctets interface{}

    // Apollo protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfapolloOutOctets interface{}

    // Vines protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfvinesInPkts interface{}

    // Vines protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfvinesOutPkts interface{}

    // Vines protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfvinesInOctets interface{}

    // Vines protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfvinesOutOctets interface{}

    // Bridged protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfbridgedInPkts interface{}

    // Bridged protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfbridgedOutPkts interface{}

    // Bridged protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfbridgedInOctets interface{}

    // Bridged protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfbridgedOutOctets interface{}

    // SRB protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfsrbInPkts interface{}

    // SRB protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfsrbOutPkts interface{}

    // SRB protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfsrbInOctets interface{}

    // SRB protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfsrbOutOctets interface{}

    // Choas protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfchaosInPkts interface{}

    // Choas protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfchaosOutPkts interface{}

    // Choas protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfchaosInOctets interface{}

    // Choas protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfchaosOutOctets interface{}

    // PUP protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfpupInPkts interface{}

    // PUP protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfpupOutPkts interface{}

    // PUP protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfpupInOctets interface{}

    // PUP protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfpupOutOctets interface{}

    // MOP protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfmopInPkts interface{}

    // MOP protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfmopOutPkts interface{}

    // MOP protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfmopInOctets interface{}

    // MOP protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfmopOutOctets interface{}

    // LanMan protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIflanmanInPkts interface{}

    // LanMan protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIflanmanOutPkts interface{}

    // LanMan protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIflanmanInOctets interface{}

    // LanMan protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIflanmanOutOctets interface{}

    // STUN protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfstunInPkts interface{}

    // STUN protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfstunOutPkts interface{}

    // STUN protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfstunInOctets interface{}

    // STUN protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfstunOutOctets interface{}

    // Spanning tree input protocol packet count. The type is interface{} with
    // range: 0..4294967295.
    LocIfspanInPkts interface{}

    // Spanning tree output protocol packet count. The type is interface{} with
    // range: 0..4294967295.
    LocIfspanOutPkts interface{}

    // Spanning tree input octet packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfspanInOctets interface{}

    // Spanning tree output octet packet count. The type is interface{} with
    // range: 0..4294967295.
    LocIfspanOutOctets interface{}

    // Arp protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfarpInPkts interface{}

    // Arp protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfarpOutPkts interface{}

    // Arp protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfarpInOctets interface{}

    // Arp protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfarpOutOctets interface{}

    // Probe protocol input packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfprobeInPkts interface{}

    // Probe protocol output packet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfprobeOutPkts interface{}

    // Probe protocol input octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfprobeInOctets interface{}

    // Probe protocol output octet count. The type is interface{} with range:
    // 0..4294967295.
    LocIfprobeOutOctets interface{}

    // The number of good packets received with the dribble condition present. The
    // type is interface{} with range: 0..4294967295.
    LocIfDribbleInputs interface{}
}

func (lifEntry *OLDCISCOINTERFACESMIB_LifTable_LifEntry) GetEntityData() *types.CommonEntityData {
    lifEntry.EntityData.YFilter = lifEntry.YFilter
    lifEntry.EntityData.YangName = "lifEntry"
    lifEntry.EntityData.BundleName = "cisco_ios_xe"
    lifEntry.EntityData.ParentYangName = "lifTable"
    lifEntry.EntityData.SegmentPath = "lifEntry" + types.AddKeyToken(lifEntry.IfIndex, "ifIndex")
    lifEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lifEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lifEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lifEntry.EntityData.Children = types.NewOrderedMap()
    lifEntry.EntityData.Leafs = types.NewOrderedMap()
    lifEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", lifEntry.IfIndex})
    lifEntry.EntityData.Leafs.Append("locIfHardType", types.YLeaf{"LocIfHardType", lifEntry.LocIfHardType})
    lifEntry.EntityData.Leafs.Append("locIfLineProt", types.YLeaf{"LocIfLineProt", lifEntry.LocIfLineProt})
    lifEntry.EntityData.Leafs.Append("locIfLastIn", types.YLeaf{"LocIfLastIn", lifEntry.LocIfLastIn})
    lifEntry.EntityData.Leafs.Append("locIfLastOut", types.YLeaf{"LocIfLastOut", lifEntry.LocIfLastOut})
    lifEntry.EntityData.Leafs.Append("locIfLastOutHang", types.YLeaf{"LocIfLastOutHang", lifEntry.LocIfLastOutHang})
    lifEntry.EntityData.Leafs.Append("locIfInBitsSec", types.YLeaf{"LocIfInBitsSec", lifEntry.LocIfInBitsSec})
    lifEntry.EntityData.Leafs.Append("locIfInPktsSec", types.YLeaf{"LocIfInPktsSec", lifEntry.LocIfInPktsSec})
    lifEntry.EntityData.Leafs.Append("locIfOutBitsSec", types.YLeaf{"LocIfOutBitsSec", lifEntry.LocIfOutBitsSec})
    lifEntry.EntityData.Leafs.Append("locIfOutPktsSec", types.YLeaf{"LocIfOutPktsSec", lifEntry.LocIfOutPktsSec})
    lifEntry.EntityData.Leafs.Append("locIfInRunts", types.YLeaf{"LocIfInRunts", lifEntry.LocIfInRunts})
    lifEntry.EntityData.Leafs.Append("locIfInGiants", types.YLeaf{"LocIfInGiants", lifEntry.LocIfInGiants})
    lifEntry.EntityData.Leafs.Append("locIfInCRC", types.YLeaf{"LocIfInCRC", lifEntry.LocIfInCRC})
    lifEntry.EntityData.Leafs.Append("locIfInFrame", types.YLeaf{"LocIfInFrame", lifEntry.LocIfInFrame})
    lifEntry.EntityData.Leafs.Append("locIfInOverrun", types.YLeaf{"LocIfInOverrun", lifEntry.LocIfInOverrun})
    lifEntry.EntityData.Leafs.Append("locIfInIgnored", types.YLeaf{"LocIfInIgnored", lifEntry.LocIfInIgnored})
    lifEntry.EntityData.Leafs.Append("locIfInAbort", types.YLeaf{"LocIfInAbort", lifEntry.LocIfInAbort})
    lifEntry.EntityData.Leafs.Append("locIfResets", types.YLeaf{"LocIfResets", lifEntry.LocIfResets})
    lifEntry.EntityData.Leafs.Append("locIfRestarts", types.YLeaf{"LocIfRestarts", lifEntry.LocIfRestarts})
    lifEntry.EntityData.Leafs.Append("locIfKeep", types.YLeaf{"LocIfKeep", lifEntry.LocIfKeep})
    lifEntry.EntityData.Leafs.Append("locIfReason", types.YLeaf{"LocIfReason", lifEntry.LocIfReason})
    lifEntry.EntityData.Leafs.Append("locIfCarTrans", types.YLeaf{"LocIfCarTrans", lifEntry.LocIfCarTrans})
    lifEntry.EntityData.Leafs.Append("locIfReliab", types.YLeaf{"LocIfReliab", lifEntry.LocIfReliab})
    lifEntry.EntityData.Leafs.Append("locIfDelay", types.YLeaf{"LocIfDelay", lifEntry.LocIfDelay})
    lifEntry.EntityData.Leafs.Append("locIfLoad", types.YLeaf{"LocIfLoad", lifEntry.LocIfLoad})
    lifEntry.EntityData.Leafs.Append("locIfCollisions", types.YLeaf{"LocIfCollisions", lifEntry.LocIfCollisions})
    lifEntry.EntityData.Leafs.Append("locIfInputQueueDrops", types.YLeaf{"LocIfInputQueueDrops", lifEntry.LocIfInputQueueDrops})
    lifEntry.EntityData.Leafs.Append("locIfOutputQueueDrops", types.YLeaf{"LocIfOutputQueueDrops", lifEntry.LocIfOutputQueueDrops})
    lifEntry.EntityData.Leafs.Append("locIfDescr", types.YLeaf{"LocIfDescr", lifEntry.LocIfDescr})
    lifEntry.EntityData.Leafs.Append("locIfSlowInPkts", types.YLeaf{"LocIfSlowInPkts", lifEntry.LocIfSlowInPkts})
    lifEntry.EntityData.Leafs.Append("locIfSlowOutPkts", types.YLeaf{"LocIfSlowOutPkts", lifEntry.LocIfSlowOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfSlowInOctets", types.YLeaf{"LocIfSlowInOctets", lifEntry.LocIfSlowInOctets})
    lifEntry.EntityData.Leafs.Append("locIfSlowOutOctets", types.YLeaf{"LocIfSlowOutOctets", lifEntry.LocIfSlowOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfFastInPkts", types.YLeaf{"LocIfFastInPkts", lifEntry.LocIfFastInPkts})
    lifEntry.EntityData.Leafs.Append("locIfFastOutPkts", types.YLeaf{"LocIfFastOutPkts", lifEntry.LocIfFastOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfFastInOctets", types.YLeaf{"LocIfFastInOctets", lifEntry.LocIfFastInOctets})
    lifEntry.EntityData.Leafs.Append("locIfFastOutOctets", types.YLeaf{"LocIfFastOutOctets", lifEntry.LocIfFastOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfotherInPkts", types.YLeaf{"LocIfotherInPkts", lifEntry.LocIfotherInPkts})
    lifEntry.EntityData.Leafs.Append("locIfotherOutPkts", types.YLeaf{"LocIfotherOutPkts", lifEntry.LocIfotherOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfotherInOctets", types.YLeaf{"LocIfotherInOctets", lifEntry.LocIfotherInOctets})
    lifEntry.EntityData.Leafs.Append("locIfotherOutOctets", types.YLeaf{"LocIfotherOutOctets", lifEntry.LocIfotherOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfipInPkts", types.YLeaf{"LocIfipInPkts", lifEntry.LocIfipInPkts})
    lifEntry.EntityData.Leafs.Append("locIfipOutPkts", types.YLeaf{"LocIfipOutPkts", lifEntry.LocIfipOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfipInOctets", types.YLeaf{"LocIfipInOctets", lifEntry.LocIfipInOctets})
    lifEntry.EntityData.Leafs.Append("locIfipOutOctets", types.YLeaf{"LocIfipOutOctets", lifEntry.LocIfipOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfdecnetInPkts", types.YLeaf{"LocIfdecnetInPkts", lifEntry.LocIfdecnetInPkts})
    lifEntry.EntityData.Leafs.Append("locIfdecnetOutPkts", types.YLeaf{"LocIfdecnetOutPkts", lifEntry.LocIfdecnetOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfdecnetInOctets", types.YLeaf{"LocIfdecnetInOctets", lifEntry.LocIfdecnetInOctets})
    lifEntry.EntityData.Leafs.Append("locIfdecnetOutOctets", types.YLeaf{"LocIfdecnetOutOctets", lifEntry.LocIfdecnetOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfxnsInPkts", types.YLeaf{"LocIfxnsInPkts", lifEntry.LocIfxnsInPkts})
    lifEntry.EntityData.Leafs.Append("locIfxnsOutPkts", types.YLeaf{"LocIfxnsOutPkts", lifEntry.LocIfxnsOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfxnsInOctets", types.YLeaf{"LocIfxnsInOctets", lifEntry.LocIfxnsInOctets})
    lifEntry.EntityData.Leafs.Append("locIfxnsOutOctets", types.YLeaf{"LocIfxnsOutOctets", lifEntry.LocIfxnsOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfclnsInPkts", types.YLeaf{"LocIfclnsInPkts", lifEntry.LocIfclnsInPkts})
    lifEntry.EntityData.Leafs.Append("locIfclnsOutPkts", types.YLeaf{"LocIfclnsOutPkts", lifEntry.LocIfclnsOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfclnsInOctets", types.YLeaf{"LocIfclnsInOctets", lifEntry.LocIfclnsInOctets})
    lifEntry.EntityData.Leafs.Append("locIfclnsOutOctets", types.YLeaf{"LocIfclnsOutOctets", lifEntry.LocIfclnsOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfappletalkInPkts", types.YLeaf{"LocIfappletalkInPkts", lifEntry.LocIfappletalkInPkts})
    lifEntry.EntityData.Leafs.Append("locIfappletalkOutPkts", types.YLeaf{"LocIfappletalkOutPkts", lifEntry.LocIfappletalkOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfappletalkInOctets", types.YLeaf{"LocIfappletalkInOctets", lifEntry.LocIfappletalkInOctets})
    lifEntry.EntityData.Leafs.Append("locIfappletalkOutOctets", types.YLeaf{"LocIfappletalkOutOctets", lifEntry.LocIfappletalkOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfnovellInPkts", types.YLeaf{"LocIfnovellInPkts", lifEntry.LocIfnovellInPkts})
    lifEntry.EntityData.Leafs.Append("locIfnovellOutPkts", types.YLeaf{"LocIfnovellOutPkts", lifEntry.LocIfnovellOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfnovellInOctets", types.YLeaf{"LocIfnovellInOctets", lifEntry.LocIfnovellInOctets})
    lifEntry.EntityData.Leafs.Append("locIfnovellOutOctets", types.YLeaf{"LocIfnovellOutOctets", lifEntry.LocIfnovellOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfapolloInPkts", types.YLeaf{"LocIfapolloInPkts", lifEntry.LocIfapolloInPkts})
    lifEntry.EntityData.Leafs.Append("locIfapolloOutPkts", types.YLeaf{"LocIfapolloOutPkts", lifEntry.LocIfapolloOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfapolloInOctets", types.YLeaf{"LocIfapolloInOctets", lifEntry.LocIfapolloInOctets})
    lifEntry.EntityData.Leafs.Append("locIfapolloOutOctets", types.YLeaf{"LocIfapolloOutOctets", lifEntry.LocIfapolloOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfvinesInPkts", types.YLeaf{"LocIfvinesInPkts", lifEntry.LocIfvinesInPkts})
    lifEntry.EntityData.Leafs.Append("locIfvinesOutPkts", types.YLeaf{"LocIfvinesOutPkts", lifEntry.LocIfvinesOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfvinesInOctets", types.YLeaf{"LocIfvinesInOctets", lifEntry.LocIfvinesInOctets})
    lifEntry.EntityData.Leafs.Append("locIfvinesOutOctets", types.YLeaf{"LocIfvinesOutOctets", lifEntry.LocIfvinesOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfbridgedInPkts", types.YLeaf{"LocIfbridgedInPkts", lifEntry.LocIfbridgedInPkts})
    lifEntry.EntityData.Leafs.Append("locIfbridgedOutPkts", types.YLeaf{"LocIfbridgedOutPkts", lifEntry.LocIfbridgedOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfbridgedInOctets", types.YLeaf{"LocIfbridgedInOctets", lifEntry.LocIfbridgedInOctets})
    lifEntry.EntityData.Leafs.Append("locIfbridgedOutOctets", types.YLeaf{"LocIfbridgedOutOctets", lifEntry.LocIfbridgedOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfsrbInPkts", types.YLeaf{"LocIfsrbInPkts", lifEntry.LocIfsrbInPkts})
    lifEntry.EntityData.Leafs.Append("locIfsrbOutPkts", types.YLeaf{"LocIfsrbOutPkts", lifEntry.LocIfsrbOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfsrbInOctets", types.YLeaf{"LocIfsrbInOctets", lifEntry.LocIfsrbInOctets})
    lifEntry.EntityData.Leafs.Append("locIfsrbOutOctets", types.YLeaf{"LocIfsrbOutOctets", lifEntry.LocIfsrbOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfchaosInPkts", types.YLeaf{"LocIfchaosInPkts", lifEntry.LocIfchaosInPkts})
    lifEntry.EntityData.Leafs.Append("locIfchaosOutPkts", types.YLeaf{"LocIfchaosOutPkts", lifEntry.LocIfchaosOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfchaosInOctets", types.YLeaf{"LocIfchaosInOctets", lifEntry.LocIfchaosInOctets})
    lifEntry.EntityData.Leafs.Append("locIfchaosOutOctets", types.YLeaf{"LocIfchaosOutOctets", lifEntry.LocIfchaosOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfpupInPkts", types.YLeaf{"LocIfpupInPkts", lifEntry.LocIfpupInPkts})
    lifEntry.EntityData.Leafs.Append("locIfpupOutPkts", types.YLeaf{"LocIfpupOutPkts", lifEntry.LocIfpupOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfpupInOctets", types.YLeaf{"LocIfpupInOctets", lifEntry.LocIfpupInOctets})
    lifEntry.EntityData.Leafs.Append("locIfpupOutOctets", types.YLeaf{"LocIfpupOutOctets", lifEntry.LocIfpupOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfmopInPkts", types.YLeaf{"LocIfmopInPkts", lifEntry.LocIfmopInPkts})
    lifEntry.EntityData.Leafs.Append("locIfmopOutPkts", types.YLeaf{"LocIfmopOutPkts", lifEntry.LocIfmopOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfmopInOctets", types.YLeaf{"LocIfmopInOctets", lifEntry.LocIfmopInOctets})
    lifEntry.EntityData.Leafs.Append("locIfmopOutOctets", types.YLeaf{"LocIfmopOutOctets", lifEntry.LocIfmopOutOctets})
    lifEntry.EntityData.Leafs.Append("locIflanmanInPkts", types.YLeaf{"LocIflanmanInPkts", lifEntry.LocIflanmanInPkts})
    lifEntry.EntityData.Leafs.Append("locIflanmanOutPkts", types.YLeaf{"LocIflanmanOutPkts", lifEntry.LocIflanmanOutPkts})
    lifEntry.EntityData.Leafs.Append("locIflanmanInOctets", types.YLeaf{"LocIflanmanInOctets", lifEntry.LocIflanmanInOctets})
    lifEntry.EntityData.Leafs.Append("locIflanmanOutOctets", types.YLeaf{"LocIflanmanOutOctets", lifEntry.LocIflanmanOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfstunInPkts", types.YLeaf{"LocIfstunInPkts", lifEntry.LocIfstunInPkts})
    lifEntry.EntityData.Leafs.Append("locIfstunOutPkts", types.YLeaf{"LocIfstunOutPkts", lifEntry.LocIfstunOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfstunInOctets", types.YLeaf{"LocIfstunInOctets", lifEntry.LocIfstunInOctets})
    lifEntry.EntityData.Leafs.Append("locIfstunOutOctets", types.YLeaf{"LocIfstunOutOctets", lifEntry.LocIfstunOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfspanInPkts", types.YLeaf{"LocIfspanInPkts", lifEntry.LocIfspanInPkts})
    lifEntry.EntityData.Leafs.Append("locIfspanOutPkts", types.YLeaf{"LocIfspanOutPkts", lifEntry.LocIfspanOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfspanInOctets", types.YLeaf{"LocIfspanInOctets", lifEntry.LocIfspanInOctets})
    lifEntry.EntityData.Leafs.Append("locIfspanOutOctets", types.YLeaf{"LocIfspanOutOctets", lifEntry.LocIfspanOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfarpInPkts", types.YLeaf{"LocIfarpInPkts", lifEntry.LocIfarpInPkts})
    lifEntry.EntityData.Leafs.Append("locIfarpOutPkts", types.YLeaf{"LocIfarpOutPkts", lifEntry.LocIfarpOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfarpInOctets", types.YLeaf{"LocIfarpInOctets", lifEntry.LocIfarpInOctets})
    lifEntry.EntityData.Leafs.Append("locIfarpOutOctets", types.YLeaf{"LocIfarpOutOctets", lifEntry.LocIfarpOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfprobeInPkts", types.YLeaf{"LocIfprobeInPkts", lifEntry.LocIfprobeInPkts})
    lifEntry.EntityData.Leafs.Append("locIfprobeOutPkts", types.YLeaf{"LocIfprobeOutPkts", lifEntry.LocIfprobeOutPkts})
    lifEntry.EntityData.Leafs.Append("locIfprobeInOctets", types.YLeaf{"LocIfprobeInOctets", lifEntry.LocIfprobeInOctets})
    lifEntry.EntityData.Leafs.Append("locIfprobeOutOctets", types.YLeaf{"LocIfprobeOutOctets", lifEntry.LocIfprobeOutOctets})
    lifEntry.EntityData.Leafs.Append("locIfDribbleInputs", types.YLeaf{"LocIfDribbleInputs", lifEntry.LocIfDribbleInputs})

    lifEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(lifEntry.EntityData)
}

// OLDCISCOINTERFACESMIB_LFSIPTable
// A list of card entries for 4T, HSSI,
// Mx serial or FSIP.
type OLDCISCOINTERFACESMIB_LFSIPTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of objects specific to 4T, HSSI, Mx serial or FSIP. The type
    // is slice of OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry.
    LFSIPEntry []*OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry
}

func (lFSIPTable *OLDCISCOINTERFACESMIB_LFSIPTable) GetEntityData() *types.CommonEntityData {
    lFSIPTable.EntityData.YFilter = lFSIPTable.YFilter
    lFSIPTable.EntityData.YangName = "lFSIPTable"
    lFSIPTable.EntityData.BundleName = "cisco_ios_xe"
    lFSIPTable.EntityData.ParentYangName = "OLD-CISCO-INTERFACES-MIB"
    lFSIPTable.EntityData.SegmentPath = "lFSIPTable"
    lFSIPTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lFSIPTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lFSIPTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lFSIPTable.EntityData.Children = types.NewOrderedMap()
    lFSIPTable.EntityData.Children.Append("lFSIPEntry", types.YChild{"LFSIPEntry", nil})
    for i := range lFSIPTable.LFSIPEntry {
        lFSIPTable.EntityData.Children.Append(types.GetSegmentPath(lFSIPTable.LFSIPEntry[i]), types.YChild{"LFSIPEntry", lFSIPTable.LFSIPEntry[i]})
    }
    lFSIPTable.EntityData.Leafs = types.NewOrderedMap()

    lFSIPTable.EntityData.YListKeys = []string {}

    return &(lFSIPTable.EntityData)
}

// OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry
// A collection of objects specific to 4T,
// HSSI, Mx serial or FSIP.
type OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface index of this card corresponding to its
    // ifIndex. The type is interface{} with range: -2147483648..2147483647.
    LocIfFSIPIndex interface{}

    // Is this FSIP line DCE or DTE. The type is LocIfFSIPtype.
    LocIfFSIPtype interface{}

    // Is the RTS signal up or down. The type is LocIfFSIPrts.
    LocIfFSIPrts interface{}

    // Is the CTS signal up or down. The type is LocIfFSIPcts.
    LocIfFSIPcts interface{}

    // Is the DTR signal up or down. The type is LocIfFSIPdtr.
    LocIfFSIPdtr interface{}

    // Is the DCD signal up or down. The type is LocIfFSIPdcd.
    LocIfFSIPdcd interface{}

    // Is the DSR signal up or down. The type is LocIfFSIPdsr.
    LocIfFSIPdsr interface{}

    // Received clock rate. The type is interface{} with range:
    // -2147483648..2147483647.
    LocIfFSIPrxClockrate interface{}

    // Use when received clock rate  is greater than 2^32 (gigabits). The type is
    // interface{} with range: -2147483648..2147483647.
    LocIfFSIPrxClockrateHi interface{}

    // Cable Type of 4T, HSSI, Mx serial or FSIP. The type is LocIfFSIPportType.
    LocIfFSIPportType interface{}
}

func (lFSIPEntry *OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry) GetEntityData() *types.CommonEntityData {
    lFSIPEntry.EntityData.YFilter = lFSIPEntry.YFilter
    lFSIPEntry.EntityData.YangName = "lFSIPEntry"
    lFSIPEntry.EntityData.BundleName = "cisco_ios_xe"
    lFSIPEntry.EntityData.ParentYangName = "lFSIPTable"
    lFSIPEntry.EntityData.SegmentPath = "lFSIPEntry" + types.AddKeyToken(lFSIPEntry.LocIfFSIPIndex, "locIfFSIPIndex")
    lFSIPEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lFSIPEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lFSIPEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lFSIPEntry.EntityData.Children = types.NewOrderedMap()
    lFSIPEntry.EntityData.Leafs = types.NewOrderedMap()
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPIndex", types.YLeaf{"LocIfFSIPIndex", lFSIPEntry.LocIfFSIPIndex})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPtype", types.YLeaf{"LocIfFSIPtype", lFSIPEntry.LocIfFSIPtype})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPrts", types.YLeaf{"LocIfFSIPrts", lFSIPEntry.LocIfFSIPrts})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPcts", types.YLeaf{"LocIfFSIPcts", lFSIPEntry.LocIfFSIPcts})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPdtr", types.YLeaf{"LocIfFSIPdtr", lFSIPEntry.LocIfFSIPdtr})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPdcd", types.YLeaf{"LocIfFSIPdcd", lFSIPEntry.LocIfFSIPdcd})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPdsr", types.YLeaf{"LocIfFSIPdsr", lFSIPEntry.LocIfFSIPdsr})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPrxClockrate", types.YLeaf{"LocIfFSIPrxClockrate", lFSIPEntry.LocIfFSIPrxClockrate})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPrxClockrateHi", types.YLeaf{"LocIfFSIPrxClockrateHi", lFSIPEntry.LocIfFSIPrxClockrateHi})
    lFSIPEntry.EntityData.Leafs.Append("locIfFSIPportType", types.YLeaf{"LocIfFSIPportType", lFSIPEntry.LocIfFSIPportType})

    lFSIPEntry.EntityData.YListKeys = []string {"LocIfFSIPIndex"}

    return &(lFSIPEntry.EntityData)
}

// OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPcts represents Is the CTS signal up or down
type OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPcts string

const (
    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPcts_notAvailable OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPcts = "notAvailable"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPcts_up OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPcts = "up"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPcts_down OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPcts = "down"
)

// OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdcd represents Is the DCD signal up or down
type OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdcd string

const (
    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdcd_notAvailable OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdcd = "notAvailable"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdcd_up OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdcd = "up"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdcd_down OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdcd = "down"
)

// OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdsr represents Is the DSR signal up or down
type OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdsr string

const (
    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdsr_notAvailable OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdsr = "notAvailable"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdsr_up OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdsr = "up"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdsr_down OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdsr = "down"
)

// OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdtr represents Is the DTR signal up or down
type OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdtr string

const (
    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdtr_notAvailable OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdtr = "notAvailable"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdtr_up OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdtr = "up"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdtr_down OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPdtr = "down"
)

// OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType represents Cable Type of 4T, HSSI, Mx serial or FSIP
type OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType string

const (
    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_noCable OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "noCable"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_rs232 OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "rs232"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_rs422 OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "rs422"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_rs423 OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "rs423"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_v35 OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "v35"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_x21 OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "x21"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_rs449 OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "rs449"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_rs530 OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "rs530"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_hssi OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "hssi"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_g703unbal OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "g703unbal"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_g703bal OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "g703bal"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType_jt2unbal OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPportType = "jt2unbal"
)

// OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPrts represents Is the RTS signal up or down
type OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPrts string

const (
    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPrts_notAvailable OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPrts = "notAvailable"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPrts_up OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPrts = "up"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPrts_down OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPrts = "down"
)

// OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPtype represents Is this FSIP line DCE or DTE
type OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPtype string

const (
    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPtype_notAvailable OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPtype = "notAvailable"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPtype_dte OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPtype = "dte"

    OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPtype_dce OLDCISCOINTERFACESMIB_LFSIPTable_LFSIPEntry_LocIfFSIPtype = "dce"
)

