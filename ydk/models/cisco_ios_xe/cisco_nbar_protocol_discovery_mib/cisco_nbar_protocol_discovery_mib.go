// Cisco NBAR Protocol Discovery MIB 
// 
// NBAR - Network Based Application Recognition is
// an intelligent classification engine that recognizes 
// applications that are static (which use fixed TCP or
// UDP port numbers), and stateful (which dynamically 
// assign TCP or UDP port numbers). 
// 
// Protocol Discovery - uses NBAR to show you the mix 
// of applications currently running on the network. 
// Key statistics are associated with each protocol. 
// These statistics can be used to define traffic 
// classes and QoS policies.
// 
// Functionality:
// 1. To enable/disable Protocol Discovery per interface.
// 2. Display the protocols/applications which NBAR
//    currently recognizes.
// 3. To display various Protocol Discovery statistics.
// 4. A configurable top N table which lists
//    protocols using user defined criteria.
// 5. To configure notifications (traps) based 
//    on configurable statistic thresholds.
// 6. To maintain a history table of all notification 
//    events.
package cisco_nbar_protocol_discovery_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_nbar_protocol_discovery_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB CISCO-NBAR-PROTOCOL-DISCOVERY-MIB}", reflect.TypeOf(CISCONBARPROTOCOLDISCOVERYMIB{}))
    ydk.RegisterEntity("CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB", reflect.TypeOf(CISCONBARPROTOCOLDISCOVERYMIB{}))
}

// CiscoPdDataType represents packetcount - unit is packets
type CiscoPdDataType string

const (
    CiscoPdDataType_bitRateIn CiscoPdDataType = "bitRateIn"

    CiscoPdDataType_bitRateOut CiscoPdDataType = "bitRateOut"

    CiscoPdDataType_bitRateSum CiscoPdDataType = "bitRateSum"

    CiscoPdDataType_byteCountIn CiscoPdDataType = "byteCountIn"

    CiscoPdDataType_byteCountOut CiscoPdDataType = "byteCountOut"

    CiscoPdDataType_byteCountSum CiscoPdDataType = "byteCountSum"

    CiscoPdDataType_packetCountIn CiscoPdDataType = "packetCountIn"

    CiscoPdDataType_packetCountOut CiscoPdDataType = "packetCountOut"

    CiscoPdDataType_packetCountSum CiscoPdDataType = "packetCountSum"
)

// CISCONBARPROTOCOLDISCOVERYMIB
type CISCONBARPROTOCOLDISCOVERYMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cnpdnotificationsconfig CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig

    // The cnpdStatusTable is used to enable and disable Protocol Discovery on an
    // interface.
    Cnpdstatustable CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable

    // The cnpdAllStatsTable contains all the statistics available for all the
    // protocols/applications currently recognized by NBAR Protocol Discovery for
    // a particular  interface.  In the event of an overflow, the 32 bit counters
    // are not  valid. There is no overflow support.
    Cnpdallstatstable CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable

    // The cnpdTopNConfigTable is used to configure cnpdTopNStatsTable's.
    Cnpdtopnconfigtable CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable

    // A cnpdTopNStatsTable describes an ordered list of protocols.
    Cnpdtopnstatstable CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable

    // The cnpdThresholdConfigTable allows the management station to create
    // thresholds for the purpose of sending notifications if breached, and
    // creating a history of breached thresholds.
    Cnpdthresholdconfigtable CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable

    // The Threshold History table. Notifications are unreliable so this table
    // provides a history of the last 5000 threshold breached events. A
    // notification can be traced back to its cnpdThresholdHistoryEntry.
    Cnpdthresholdhistorytable CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable

    // The Supported Protocols table lists all the  protocols and applications
    // which NBAR is currently capable of recognizing.
    Cnpdsupportedprotocolstable CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable
}

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetEntityData() *types.CommonEntityData {
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.YFilter = cISCONBARPROTOCOLDISCOVERYMIB.YFilter
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.YangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.SegmentPath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children = make(map[string]types.YChild)
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children["cnpdNotificationsConfig"] = types.YChild{"Cnpdnotificationsconfig", &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdnotificationsconfig}
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children["cnpdStatusTable"] = types.YChild{"Cnpdstatustable", &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdstatustable}
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children["cnpdAllStatsTable"] = types.YChild{"Cnpdallstatstable", &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdallstatstable}
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children["cnpdTopNConfigTable"] = types.YChild{"Cnpdtopnconfigtable", &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdtopnconfigtable}
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children["cnpdTopNStatsTable"] = types.YChild{"Cnpdtopnstatstable", &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdtopnstatstable}
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children["cnpdThresholdConfigTable"] = types.YChild{"Cnpdthresholdconfigtable", &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdthresholdconfigtable}
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children["cnpdThresholdHistoryTable"] = types.YChild{"Cnpdthresholdhistorytable", &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdthresholdhistorytable}
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children["cnpdSupportedProtocolsTable"] = types.YChild{"Cnpdsupportedprotocolstable", &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdsupportedprotocolstable}
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCONBARPROTOCOLDISCOVERYMIB.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object is used to enable or disable  Notifications on a global basis. 
    // If set to 'true' - Notifications are enabled. If set to 'false' -
    // Notifications are disabled. The type is bool.
    Cnpdnotificationsenable interface{}
}

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetEntityData() *types.CommonEntityData {
    cnpdnotificationsconfig.EntityData.YFilter = cnpdnotificationsconfig.YFilter
    cnpdnotificationsconfig.EntityData.YangName = "cnpdNotificationsConfig"
    cnpdnotificationsconfig.EntityData.BundleName = "cisco_ios_xe"
    cnpdnotificationsconfig.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdnotificationsconfig.EntityData.SegmentPath = "cnpdNotificationsConfig"
    cnpdnotificationsconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdnotificationsconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdnotificationsconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdnotificationsconfig.EntityData.Children = make(map[string]types.YChild)
    cnpdnotificationsconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    cnpdnotificationsconfig.EntityData.Leafs["cnpdNotificationsEnable"] = types.YLeaf{"Cnpdnotificationsenable", cnpdnotificationsconfig.Cnpdnotificationsenable}
    return &(cnpdnotificationsconfig.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable
// The cnpdStatusTable is used to enable and
// disable Protocol Discovery on an interface.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cnpdStatusTable contains objects for enabling or disabling
    // Protocol Discovery on a per interface basis. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry.
    Cnpdstatusentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry
}

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetEntityData() *types.CommonEntityData {
    cnpdstatustable.EntityData.YFilter = cnpdstatustable.YFilter
    cnpdstatustable.EntityData.YangName = "cnpdStatusTable"
    cnpdstatustable.EntityData.BundleName = "cisco_ios_xe"
    cnpdstatustable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdstatustable.EntityData.SegmentPath = "cnpdStatusTable"
    cnpdstatustable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdstatustable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdstatustable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdstatustable.EntityData.Children = make(map[string]types.YChild)
    cnpdstatustable.EntityData.Children["cnpdStatusEntry"] = types.YChild{"Cnpdstatusentry", nil}
    for i := range cnpdstatustable.Cnpdstatusentry {
        cnpdstatustable.EntityData.Children[types.GetSegmentPath(&cnpdstatustable.Cnpdstatusentry[i])] = types.YChild{"Cnpdstatusentry", &cnpdstatustable.Cnpdstatusentry[i]}
    }
    cnpdstatustable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnpdstatustable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry
// An entry in the cnpdStatusTable contains objects
// for enabling or disabling Protocol Discovery on a
// per interface basis.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This object is used to enable or disable  Protocol Discovery on an
    // interface.   If set to 'true' - Protocol Discovery is  enabled on this
    // Interface.  If set to 'false' - Protocol Discovery is  disabled on this
    // Interface. The type is bool.
    Cnpdstatuspdenable interface{}

    // The value of sysUpTime at the time Protocol  Discovery was last enabled  on
    // an interface. If the interface does not have Protocol Discovery enabled
    // this value is zero. The type is interface{} with range: 0..4294967295.
    Cnpdstatuslastupdatetime interface{}
}

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetEntityData() *types.CommonEntityData {
    cnpdstatusentry.EntityData.YFilter = cnpdstatusentry.YFilter
    cnpdstatusentry.EntityData.YangName = "cnpdStatusEntry"
    cnpdstatusentry.EntityData.BundleName = "cisco_ios_xe"
    cnpdstatusentry.EntityData.ParentYangName = "cnpdStatusTable"
    cnpdstatusentry.EntityData.SegmentPath = "cnpdStatusEntry" + "[ifIndex='" + fmt.Sprintf("%v", cnpdstatusentry.Ifindex) + "']"
    cnpdstatusentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdstatusentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdstatusentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdstatusentry.EntityData.Children = make(map[string]types.YChild)
    cnpdstatusentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnpdstatusentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cnpdstatusentry.Ifindex}
    cnpdstatusentry.EntityData.Leafs["cnpdStatusPdEnable"] = types.YLeaf{"Cnpdstatuspdenable", cnpdstatusentry.Cnpdstatuspdenable}
    cnpdstatusentry.EntityData.Leafs["cnpdStatusLastUpdateTime"] = types.YLeaf{"Cnpdstatuslastupdatetime", cnpdstatusentry.Cnpdstatuslastupdatetime}
    return &(cnpdstatusentry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable
// The cnpdAllStatsTable contains all the statistics
// available for all the protocols/applications currently
// recognized by NBAR Protocol Discovery for a particular 
// interface.
// 
// In the event of an overflow, the 32 bit counters are not 
// valid. There is no overflow support.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cnpdAllStatsTable table. This entry  contains the
    // statistics collected on all the protocols  which NBAR classifies for a
    // particular interface. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry.
    Cnpdallstatsentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry
}

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetEntityData() *types.CommonEntityData {
    cnpdallstatstable.EntityData.YFilter = cnpdallstatstable.YFilter
    cnpdallstatstable.EntityData.YangName = "cnpdAllStatsTable"
    cnpdallstatstable.EntityData.BundleName = "cisco_ios_xe"
    cnpdallstatstable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdallstatstable.EntityData.SegmentPath = "cnpdAllStatsTable"
    cnpdallstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdallstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdallstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdallstatstable.EntityData.Children = make(map[string]types.YChild)
    cnpdallstatstable.EntityData.Children["cnpdAllStatsEntry"] = types.YChild{"Cnpdallstatsentry", nil}
    for i := range cnpdallstatstable.Cnpdallstatsentry {
        cnpdallstatstable.EntityData.Children[types.GetSegmentPath(&cnpdallstatstable.Cnpdallstatsentry[i])] = types.YChild{"Cnpdallstatsentry", &cnpdallstatstable.Cnpdallstatsentry[i]}
    }
    cnpdallstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnpdallstatstable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry
// An entry in the cnpdAllStatsTable table. This entry 
// contains the statistics collected on all the protocols 
// which NBAR classifies for a particular interface.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. An object which represents a unique  identifier
    // for a protocol or application  which NBAR currently recognizes.  This
    // object is an index into the  SupportedProtocolsTable where details of the
    // protocol can be found. The type is interface{} with range: 1..1024.
    Cnpdallstatsprotocolsindex interface{}

    // Name of the application or protocol, a  unique textual string, assigned in
    // the cnpdSupportedProtocolsTable. The type is string with length: 1..255.
    Cnpdallstatsprotocolname interface{}

    // The packet count of inbound packets as  determined by Protocol Discovery.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    Cnpdallstatsinpkts interface{}

    // The packet count of outbound packets as  determined by Protocol Discovery.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    Cnpdallstatsoutpkts interface{}

    // The byte count of inbound octets as  determined by Protocol Discovery. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    Cnpdallstatsinbytes interface{}

    // The byte count of outbound octets as determined by Protocol Discovery. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    Cnpdallstatsoutbytes interface{}

    // The packet count of inbound packets as  determined by Protocol Discovery.
    // This is the 64-bit (High Capacity) version of cnpdAllStatsInPkts. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    Cnpdallstatshcinpkts interface{}

    // The packet count of outbound packets as  determined by Protocol Discovery.
    // This is the 64-bit (High Capacity) version of cnpdAllStatsOutPkts. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    Cnpdallstatshcoutpkts interface{}

    // The byte count of inbound octets as  determined by Protocol Discovery. This
    // is the 64-bit (High Capacity) version of cnpdAllStatsInBytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    Cnpdallstatshcinbytes interface{}

    // The byte count of outbound octets as  determined by Protocol Discovery.
    // This is the 64-bit (High Capacity) version of cnpdAllStatsOutBytes. The
    // type is interface{} with range: 0..18446744073709551615. Units are bytes.
    Cnpdallstatshcoutbytes interface{}

    // The inbound bit rate as determined  by Protocol Discovery. The type is
    // interface{} with range: 1..4294967295. Units are kilo bits per second.
    Cnpdallstatsinbitrate interface{}

    // The outbound bit rate as determined  by Protocol Discovery. The type is
    // interface{} with range: 1..4294967295. Units are kilo bits per second.
    Cnpdallstatsoutbitrate interface{}
}

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetEntityData() *types.CommonEntityData {
    cnpdallstatsentry.EntityData.YFilter = cnpdallstatsentry.YFilter
    cnpdallstatsentry.EntityData.YangName = "cnpdAllStatsEntry"
    cnpdallstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cnpdallstatsentry.EntityData.ParentYangName = "cnpdAllStatsTable"
    cnpdallstatsentry.EntityData.SegmentPath = "cnpdAllStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", cnpdallstatsentry.Ifindex) + "']" + "[cnpdAllStatsProtocolsIndex='" + fmt.Sprintf("%v", cnpdallstatsentry.Cnpdallstatsprotocolsindex) + "']"
    cnpdallstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdallstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdallstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdallstatsentry.EntityData.Children = make(map[string]types.YChild)
    cnpdallstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnpdallstatsentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cnpdallstatsentry.Ifindex}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsProtocolsIndex"] = types.YLeaf{"Cnpdallstatsprotocolsindex", cnpdallstatsentry.Cnpdallstatsprotocolsindex}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsProtocolName"] = types.YLeaf{"Cnpdallstatsprotocolname", cnpdallstatsentry.Cnpdallstatsprotocolname}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsInPkts"] = types.YLeaf{"Cnpdallstatsinpkts", cnpdallstatsentry.Cnpdallstatsinpkts}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsOutPkts"] = types.YLeaf{"Cnpdallstatsoutpkts", cnpdallstatsentry.Cnpdallstatsoutpkts}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsInBytes"] = types.YLeaf{"Cnpdallstatsinbytes", cnpdallstatsentry.Cnpdallstatsinbytes}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsOutBytes"] = types.YLeaf{"Cnpdallstatsoutbytes", cnpdallstatsentry.Cnpdallstatsoutbytes}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsHCInPkts"] = types.YLeaf{"Cnpdallstatshcinpkts", cnpdallstatsentry.Cnpdallstatshcinpkts}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsHCOutPkts"] = types.YLeaf{"Cnpdallstatshcoutpkts", cnpdallstatsentry.Cnpdallstatshcoutpkts}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsHCInBytes"] = types.YLeaf{"Cnpdallstatshcinbytes", cnpdallstatsentry.Cnpdallstatshcinbytes}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsHCOutBytes"] = types.YLeaf{"Cnpdallstatshcoutbytes", cnpdallstatsentry.Cnpdallstatshcoutbytes}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsInBitRate"] = types.YLeaf{"Cnpdallstatsinbitrate", cnpdallstatsentry.Cnpdallstatsinbitrate}
    cnpdallstatsentry.EntityData.Leafs["cnpdAllStatsOutBitRate"] = types.YLeaf{"Cnpdallstatsoutbitrate", cnpdallstatsentry.Cnpdallstatsoutbitrate}
    return &(cnpdallstatsentry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable
// The cnpdTopNConfigTable is used to configure
// cnpdTopNStatsTable's.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry provides the objects to configure and thus initiate the
    // generation of a cnpdTopNStatsTable.. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry.
    Cnpdtopnconfigentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry
}

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetEntityData() *types.CommonEntityData {
    cnpdtopnconfigtable.EntityData.YFilter = cnpdtopnconfigtable.YFilter
    cnpdtopnconfigtable.EntityData.YangName = "cnpdTopNConfigTable"
    cnpdtopnconfigtable.EntityData.BundleName = "cisco_ios_xe"
    cnpdtopnconfigtable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdtopnconfigtable.EntityData.SegmentPath = "cnpdTopNConfigTable"
    cnpdtopnconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdtopnconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdtopnconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdtopnconfigtable.EntityData.Children = make(map[string]types.YChild)
    cnpdtopnconfigtable.EntityData.Children["cnpdTopNConfigEntry"] = types.YChild{"Cnpdtopnconfigentry", nil}
    for i := range cnpdtopnconfigtable.Cnpdtopnconfigentry {
        cnpdtopnconfigtable.EntityData.Children[types.GetSegmentPath(&cnpdtopnconfigtable.Cnpdtopnconfigentry[i])] = types.YChild{"Cnpdtopnconfigentry", &cnpdtopnconfigtable.Cnpdtopnconfigentry[i]}
    }
    cnpdtopnconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnpdtopnconfigtable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry
// This entry provides the objects to configure and thus
// initiate the generation of a cnpdTopNStatsTable..
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A monotonically increasing integer which uniquely
    // identifies a cnpdTopNConfigEntry  in the cnpdTopNConfigTable. The type is
    // interface{} with range: 1..50.
    Cnpdtopnconfigindex interface{}

    // This object allows the management station to select the interface, which
    // Protocol Discovery is running on, to be used to create this 
    // cnpdTopNConfigEntry. The type is interface{} with range: 1..2147483647.
    Cnpdtopnconfigifindex interface{}

    // This object allows the management station to select the statistic used to
    // base the order of the top-n table on.  For example: a
    // cnpdTopNConfigStatsSelect of bitRateSum means order this table based on
    // each applications/protocols combined in and out bitrate. The type is
    // CiscoPdDataType.
    Cnpdtopnconfigstatsselect interface{}

    // If the cnpdTopNConfigStatsSelect is bitRateIn, bitRateOut or bitRateSum,
    // then this value is the interval in seconds that  the bitrate is sampled. 
    // This has no effect if the cnpdTopNConfigStatsSelect is byte or packet
    // based.  When this object is modified by the management  station, a new
    // sample period is started regardless of whether the original
    // cnpdTopNConfigSampleTime was finished. The type is interface{} with range:
    // 1..2048. Units are seconds.
    Cnpdtopnconfigsampletime interface{}

    // The requested size of the associated  cnpdTopNStatsTable entry.  For
    // example a cnpdTopNConfigRequestedSize of 20 indicates the management
    // station wants to create an associated  cnpdTopNStatsTable  entry of 20
    // protocol/application's. The type is interface{} with range: 1..500.
    Cnpdtopnconfigrequestedsize interface{}

    // The actual size of the associated 	 cnpdTopNStatsTable entry.  The reason
    // this may differ from  cnpdTopNConfigRequestedSize is because a  management
    // station may request a number of  protocols that is greater than the number
    // of  protocols actually found by Protocol Discovery. The type is interface{}
    // with range: 1..500.
    Cnpdtopnconfiggrantedsize interface{}

    // The value of sysUpTime when the associated cnpdTopNStatsTable entry was
    // created. The type is interface{} with range: 0..4294967295.
    Cnpdtopnconfigtime interface{}

    // This object is used to create or delete  the row entry in
    // cnpdTopNConfigTable.  When creating a row entry the management station is
    // required to specify a value for cnpdTopNConfigIfIndex only.  'notReady'
    // means that a row exists but  either it has no valid IfIndex or it has  not
    // been set to createAndGo or active.  'active' means that a createAndGo or
    // active  has been issued, AND a valid ifIndex exists.  Therefore if a row is
    // 'active' it means a  TopNStats entry has been generated.  If you set an
    // 'active' row to createAndWait  it will get the status 'notReady'.   If you
    // set any row to 'notReady' - it will go  to the 'notReadystate'.  If you set
    // any row to 'notInService' - it will  go to the 'notInService' state and the
    // corresponding  TopNStatsEntry will be deleted.  The same TopNConfig entry
    // can be re-used without  changes by setting it to 'active'. The
    // corresponding  TopStatsTable entry will be regenerated. This can  be used
    // by the NMS to poll a particular TopNConfig  Entry.  Changes to an existing
    // TopNConfig entry can be made by setting the status to 'createAndWait' and
    // changing the necessary objects. Setting it to 'createAndGo' or 'active'
    // will cause the corresponding TopNStats entry to be regenerated. The type is
    // RowStatus.
    Cnpdtopnconfigstatus interface{}
}

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetEntityData() *types.CommonEntityData {
    cnpdtopnconfigentry.EntityData.YFilter = cnpdtopnconfigentry.YFilter
    cnpdtopnconfigentry.EntityData.YangName = "cnpdTopNConfigEntry"
    cnpdtopnconfigentry.EntityData.BundleName = "cisco_ios_xe"
    cnpdtopnconfigentry.EntityData.ParentYangName = "cnpdTopNConfigTable"
    cnpdtopnconfigentry.EntityData.SegmentPath = "cnpdTopNConfigEntry" + "[cnpdTopNConfigIndex='" + fmt.Sprintf("%v", cnpdtopnconfigentry.Cnpdtopnconfigindex) + "']"
    cnpdtopnconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdtopnconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdtopnconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdtopnconfigentry.EntityData.Children = make(map[string]types.YChild)
    cnpdtopnconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnpdtopnconfigentry.EntityData.Leafs["cnpdTopNConfigIndex"] = types.YLeaf{"Cnpdtopnconfigindex", cnpdtopnconfigentry.Cnpdtopnconfigindex}
    cnpdtopnconfigentry.EntityData.Leafs["cnpdTopNConfigIfIndex"] = types.YLeaf{"Cnpdtopnconfigifindex", cnpdtopnconfigentry.Cnpdtopnconfigifindex}
    cnpdtopnconfigentry.EntityData.Leafs["cnpdTopNConfigStatsSelect"] = types.YLeaf{"Cnpdtopnconfigstatsselect", cnpdtopnconfigentry.Cnpdtopnconfigstatsselect}
    cnpdtopnconfigentry.EntityData.Leafs["cnpdTopNConfigSampleTime"] = types.YLeaf{"Cnpdtopnconfigsampletime", cnpdtopnconfigentry.Cnpdtopnconfigsampletime}
    cnpdtopnconfigentry.EntityData.Leafs["cnpdTopNConfigRequestedSize"] = types.YLeaf{"Cnpdtopnconfigrequestedsize", cnpdtopnconfigentry.Cnpdtopnconfigrequestedsize}
    cnpdtopnconfigentry.EntityData.Leafs["cnpdTopNConfigGrantedSize"] = types.YLeaf{"Cnpdtopnconfiggrantedsize", cnpdtopnconfigentry.Cnpdtopnconfiggrantedsize}
    cnpdtopnconfigentry.EntityData.Leafs["cnpdTopNConfigTime"] = types.YLeaf{"Cnpdtopnconfigtime", cnpdtopnconfigentry.Cnpdtopnconfigtime}
    cnpdtopnconfigentry.EntityData.Leafs["cnpdTopNConfigStatus"] = types.YLeaf{"Cnpdtopnconfigstatus", cnpdtopnconfigentry.Cnpdtopnconfigstatus}
    return &(cnpdtopnconfigentry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable
// A cnpdTopNStatsTable describes an ordered
// list of protocols.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry is used to store a set of objects which  describe a
    // cnpdTopNStatsTable. A cnpdTopNStatsTable  is a number of protocols and
    // statistics sorted  according to the criteria in the associated
    // cnpdTopNConfigEntry.  Therefore a cnpdTopNStatsTable can differ in content 
    // and size according to what was configured in the associated
    // cnpdTopNConfigTableEntry. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry.
    Cnpdtopnstatsentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry
}

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetEntityData() *types.CommonEntityData {
    cnpdtopnstatstable.EntityData.YFilter = cnpdtopnstatstable.YFilter
    cnpdtopnstatstable.EntityData.YangName = "cnpdTopNStatsTable"
    cnpdtopnstatstable.EntityData.BundleName = "cisco_ios_xe"
    cnpdtopnstatstable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdtopnstatstable.EntityData.SegmentPath = "cnpdTopNStatsTable"
    cnpdtopnstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdtopnstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdtopnstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdtopnstatstable.EntityData.Children = make(map[string]types.YChild)
    cnpdtopnstatstable.EntityData.Children["cnpdTopNStatsEntry"] = types.YChild{"Cnpdtopnstatsentry", nil}
    for i := range cnpdtopnstatstable.Cnpdtopnstatsentry {
        cnpdtopnstatstable.EntityData.Children[types.GetSegmentPath(&cnpdtopnstatstable.Cnpdtopnstatsentry[i])] = types.YChild{"Cnpdtopnstatsentry", &cnpdtopnstatstable.Cnpdtopnstatsentry[i]}
    }
    cnpdtopnstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnpdtopnstatstable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry
// This entry is used to store a set of objects which 
// describe a cnpdTopNStatsTable. A cnpdTopNStatsTable 
// is a number of protocols and statistics sorted 
// according to the criteria in the associated
// cnpdTopNConfigEntry.
// 
// Therefore a cnpdTopNStatsTable can differ in content 
// and size according to what was configured in the associated
// cnpdTopNConfigTableEntry.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..50. Refers to
    // cisco_nbar_protocol_discovery_mib.CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry_Cnpdtopnconfigindex
    Cnpdtopnconfigindex interface{}

    // This attribute is a key. A monotonically increasing integer which  uniquely
    // identifies a cnpdTopNStatsEntry  in the cnpdTopNStatsTable. The type is
    // interface{} with range: 1..500.
    Cnpdtopnstatsindex interface{}

    // Name of the application or protocol,  a unique textual string, assigned in
    // the cnpdSupportedProtocolsTable. The type is string with length: 1..255.
    Cnpdtopnstatsprotocolname interface{}

    // The amount of change in the selected statistic during this sampling
    // interval. The selected statistic is the cnpdTopNConfigStatsSelect from the
    // associated cnpdTopNConfigStatsEntry. The type is interface{} with range:
    // 0..4294967295.
    Cnpdtopnstatsrate interface{}

    // The amount of change in the selected statistic during this sampling
    // interval. The selected statistic is the cnpdTopNConfigStatsSelect from the
    // associated cnpdTopNConfigStatsEntry.	  This is the 64-bit (High Capacity)
    // version of  cnpdTopNStatsRate. The type is interface{} with range:
    // 0..18446744073709551615.
    Cnpdtopnstatshcrate interface{}
}

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetEntityData() *types.CommonEntityData {
    cnpdtopnstatsentry.EntityData.YFilter = cnpdtopnstatsentry.YFilter
    cnpdtopnstatsentry.EntityData.YangName = "cnpdTopNStatsEntry"
    cnpdtopnstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cnpdtopnstatsentry.EntityData.ParentYangName = "cnpdTopNStatsTable"
    cnpdtopnstatsentry.EntityData.SegmentPath = "cnpdTopNStatsEntry" + "[cnpdTopNConfigIndex='" + fmt.Sprintf("%v", cnpdtopnstatsentry.Cnpdtopnconfigindex) + "']" + "[cnpdTopNStatsIndex='" + fmt.Sprintf("%v", cnpdtopnstatsentry.Cnpdtopnstatsindex) + "']"
    cnpdtopnstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdtopnstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdtopnstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdtopnstatsentry.EntityData.Children = make(map[string]types.YChild)
    cnpdtopnstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnpdtopnstatsentry.EntityData.Leafs["cnpdTopNConfigIndex"] = types.YLeaf{"Cnpdtopnconfigindex", cnpdtopnstatsentry.Cnpdtopnconfigindex}
    cnpdtopnstatsentry.EntityData.Leafs["cnpdTopNStatsIndex"] = types.YLeaf{"Cnpdtopnstatsindex", cnpdtopnstatsentry.Cnpdtopnstatsindex}
    cnpdtopnstatsentry.EntityData.Leafs["cnpdTopNStatsProtocolName"] = types.YLeaf{"Cnpdtopnstatsprotocolname", cnpdtopnstatsentry.Cnpdtopnstatsprotocolname}
    cnpdtopnstatsentry.EntityData.Leafs["cnpdTopNStatsRate"] = types.YLeaf{"Cnpdtopnstatsrate", cnpdtopnstatsentry.Cnpdtopnstatsrate}
    cnpdtopnstatsentry.EntityData.Leafs["cnpdTopNStatsHCRate"] = types.YLeaf{"Cnpdtopnstatshcrate", cnpdtopnstatsentry.Cnpdtopnstatshcrate}
    return &(cnpdtopnstatsentry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable
// The cnpdThresholdConfigTable allows the management
// station to create thresholds for the purpose of
// sending notifications if breached, and creating a
// history of breached thresholds.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry contains configuration information to  set thresholds for the
    // purpose of notifications.  The management station is allowed to set
    // thresholds on individual statistics for individual protocols on an
    // interface. If the threshold is breached by the protocol statistic, a new
    // event is written to the cnpdThresholdHistoryTable, which in turn will 
    // generate a Notification Event.  This function has a hysteresis mechanism to
    // limit the generation of events. This mechanism generates one event as a
    // threshold is crossed in the appropriate direction. No more events are
    // generated for that threshold until the opposite threshold is crossed. This
    // stops repeated Notification events being generated each time the value is
    // sampled, when the value is above the threshold. Instead one notification is
    // sent when the threshold is breached and one notification when the statistic
    // drops below the threshold value again. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry.
    Cnpdthresholdconfigentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry
}

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetEntityData() *types.CommonEntityData {
    cnpdthresholdconfigtable.EntityData.YFilter = cnpdthresholdconfigtable.YFilter
    cnpdthresholdconfigtable.EntityData.YangName = "cnpdThresholdConfigTable"
    cnpdthresholdconfigtable.EntityData.BundleName = "cisco_ios_xe"
    cnpdthresholdconfigtable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdthresholdconfigtable.EntityData.SegmentPath = "cnpdThresholdConfigTable"
    cnpdthresholdconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdthresholdconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdthresholdconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdthresholdconfigtable.EntityData.Children = make(map[string]types.YChild)
    cnpdthresholdconfigtable.EntityData.Children["cnpdThresholdConfigEntry"] = types.YChild{"Cnpdthresholdconfigentry", nil}
    for i := range cnpdthresholdconfigtable.Cnpdthresholdconfigentry {
        cnpdthresholdconfigtable.EntityData.Children[types.GetSegmentPath(&cnpdthresholdconfigtable.Cnpdthresholdconfigentry[i])] = types.YChild{"Cnpdthresholdconfigentry", &cnpdthresholdconfigtable.Cnpdthresholdconfigentry[i]}
    }
    cnpdthresholdconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnpdthresholdconfigtable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry
// This entry contains configuration information to 
// set thresholds for the purpose of notifications.
// 
// The management station is allowed to set thresholds
// on individual statistics for individual protocols
// on an interface. If the threshold is breached by
// the protocol statistic, a new event is written to
// the cnpdThresholdHistoryTable, which in turn will 
// generate a Notification Event.
// 
// This function has a hysteresis mechanism to limit
// the generation of events. This mechanism generates
// one event as a threshold is crossed in the
// appropriate direction. No more events are generated
// for that threshold until the opposite threshold is
// crossed. This stops repeated Notification events
// being generated each time the value is sampled,
// when the value is above the threshold. Instead one
// notification is sent when the threshold is breached
// and one notification when the statistic drops below
// the threshold value again.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A monotonically increasing integer which  uniquely
    // identifies an entry in the  cnpdThresholdConfigTable. The type is
    // interface{} with range: 1..100.
    Cnpdthresholdconfigindex interface{}

    // This object allows the management station to  select the interface, which
    // Protocol Discovery is  running on, to be used to create this 
    // cnpdThresholdConfigTable entry. The type is interface{} with range:
    // 1..2147483647.
    Cnpdthresholdconfigifindex interface{}

    // The interval in seconds over which the data is sampled and compared with
    // cnpdThresholdConfigRising and cnpdThresholdConfigFalling thresholds. The
    // type is interface{} with range: 1..2048. Units are seconds.
    Cnpdthresholdconfiginterval interface{}

    // The method of sampling the selected statistic and calculating the value to
    // be compared against  cnpdThresholdConfigRising or 
    // cnpdThresholdConfigFalling thresholds.  		 If the value of this object is
    // absoluteValue(1),  the value at the end of the sampling interval 
    // cnpdThresholdConfigInterval, will be compared  with the
    // cnpdThresholdConfigRising and  cnpdThresholdConfigFalling thresholds.   In
    // this mode, when cnpdThresholdConfigStatsSelect is byte or packet based, a
    // maximum of two  cnpdThresholdHistory entries will be created per
    // application, as these byte and packet counts  monotonically increase from
    // zero. 		 If the value of this object is deltaValue(2),  the difference
    // between the samples at the  beginning and end of the
    // cnpdThresholdConfigInterval  will be compared with the
    // cnpdThresholdConfigRising  and cnpdThresholdConfigFalling thresholds. 		
    // Because the difference in the previous and current samples are compared
    // over the sample period cnpdThresholdConfigInterval, this mode provides 
    // more granularity to the thresholds because the NMS  is now provided with
    // the gradient or change in the  cnpdThresholdConfigStatsSelect.  Note that
    // even though the sample value is monotonically increasing for byte and
    // packet counts,  cnpdThresholdConfigSampleType set to deltaValue, can 
    // generate falling cnpdThresholdHistory entries, because the gradient can be
    // lower than the  cnpdThresholdConfigFalling value. The type is
    // Cnpdthresholdconfigsampletype.
    Cnpdthresholdconfigsampletype interface{}

    // The application or protocol which the management station wishes to
    // configure a threshold on.  This object is an index into the 
    // SupportedProtocolsTable where details of the protocol can be found.  If
    // cnpdThresholdConfigProtocolAny is set to TRUE this value will be ignored.
    // If it is set to FALSE, then cnpdThresholdConfigProtocol will be the only
    // protocol that is checked to see if it has breached the threshold. The type
    // is interface{} with range: 1..1024.
    Cnpdthresholdconfigprotocol interface{}

    // If set to 'true' - this threshold is configured to check for any protocol
    // which meets the threshold criteria. This means that multiple protocols can
    // generate ThresholdHistoryTable entries. Each protocol is subject to the
    // hysterisis mechanism.  If set to 'false' - this threshold is configured to
    // check for the protocol which meets the threshold criteria referred to by
    // cnpdThresholdConfigProtocol. The type is bool.
    Cnpdthresholdconfigprotocolany interface{}

    // This object allows the management station to select the statistic used to
    // base the threshold on.  For example a cnpdThresholdConfigStatsSelect of
    // bitRateSum means cnpdThresholdConfigRising and cnpdThresholdConfigFalling
    // are values based on the combined value of in and out bitrates. The type is
    // CiscoPdDataType.
    Cnpdthresholdconfigstatsselect interface{}

    // This controls the type of notification that is  sent when this threshold
    // entry is first enabled.   Because there is no previous sampling history,
    // choosing one of these options determines the type of notification generated
    // - Rising or Falling.  If the first sample after this entry is enabled  is
    // greater than or equal to cnpdThresholdConfigRising and this object is equal
    // to rising(1) or risingOrFalling(3),  then a single rising notification will
    // be generated.   If the first sample after this entry is enabled is less
    // than or equal to cnpdThresholdConfigFalling and this object is equal to
    // falling(2) or  risingOrFalling(3), then a single falling notification  will
    // be generated. The type is Cnpdthresholdconfigstartup.
    Cnpdthresholdconfigstartup interface{}

    // This is the threshold object which the managment station sets to determine
    // if it gets breached. It  indicates the statistic being sampled was rising. 
    // When the current sample is greater than or  equal to this object, and the
    // value at the last  sampling interval was less than this object (in  other
    // words the value is rising), an entry in the  cnpdThresholdHistoryTable will
    // be created.  After a rising event is generated, another such  event will
    // not be generated until the sampled value falls below this threshold and
    // reaches the cnpdThresholdConfigFalling value.  This ensures that samples
    // which are taken after a cnpdThresholdConfigRising threshold event has been
    // created, do not create further thresholds and therefore notifications,
    // until the  cnpdThresholdConfigFalling threshold has been met.  Thus a very
    // short cnpdThresholdConfigInterval can be chosen without risk of multiple
    // notifications for the same threshold breach condition. The type is
    // interface{} with range: 1..4294967295.
    Cnpdthresholdconfigrising interface{}

    // This is the threshold object which the management  station sets to
    // determine if it gets breached. It  indicates the statistic being sampled
    // was falling.   When current sample is less than or equal  to this object,
    // and the value at the last sampling interval was greater than this object
    // (in other  words the value is falling), an entry in the 
    // cnpdThresholdHistoryTable will be created.  		 After a falling event is
    // generated, another  such event will not be generated until the sampled 
    // value rises above this object and reaches the cnpdThresholdConfigRising
    // value. The type is interface{} with range: 1..4294967295.
    Cnpdthresholdconfigfalling interface{}

    // This object is used to create or delete  the row entry in
    // cnpdThresholdConfigTable.   When creating a row entry the management
    // station  is required to specify a value for  cnpdThresholdConfigIfIndex,
    // cnpdThresholdConfigRising  and cnpdThresholdConfigFalling.  'active' means
    // that a createAndGo or active has  been issued, AND a valid ifIndex exists.
    // And therefore  if a row is 'active' it means a ThresholdHistory entry  may
    // have been generated if the value was breached.  If you set an 'active' row
    // to 'createAndWait' - it will  in fact get the status 'notReady'.   Likewise
    // if you set any row to 'notInService' or 'notReady'  it will go to the
    // 'notReady' state. The type is RowStatus.
    Cnpdthresholdconfigstatus interface{}
}

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetEntityData() *types.CommonEntityData {
    cnpdthresholdconfigentry.EntityData.YFilter = cnpdthresholdconfigentry.YFilter
    cnpdthresholdconfigentry.EntityData.YangName = "cnpdThresholdConfigEntry"
    cnpdthresholdconfigentry.EntityData.BundleName = "cisco_ios_xe"
    cnpdthresholdconfigentry.EntityData.ParentYangName = "cnpdThresholdConfigTable"
    cnpdthresholdconfigentry.EntityData.SegmentPath = "cnpdThresholdConfigEntry" + "[cnpdThresholdConfigIndex='" + fmt.Sprintf("%v", cnpdthresholdconfigentry.Cnpdthresholdconfigindex) + "']"
    cnpdthresholdconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdthresholdconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdthresholdconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdthresholdconfigentry.EntityData.Children = make(map[string]types.YChild)
    cnpdthresholdconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigIndex"] = types.YLeaf{"Cnpdthresholdconfigindex", cnpdthresholdconfigentry.Cnpdthresholdconfigindex}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigIfIndex"] = types.YLeaf{"Cnpdthresholdconfigifindex", cnpdthresholdconfigentry.Cnpdthresholdconfigifindex}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigInterval"] = types.YLeaf{"Cnpdthresholdconfiginterval", cnpdthresholdconfigentry.Cnpdthresholdconfiginterval}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigSampleType"] = types.YLeaf{"Cnpdthresholdconfigsampletype", cnpdthresholdconfigentry.Cnpdthresholdconfigsampletype}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigProtocol"] = types.YLeaf{"Cnpdthresholdconfigprotocol", cnpdthresholdconfigentry.Cnpdthresholdconfigprotocol}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigProtocolAny"] = types.YLeaf{"Cnpdthresholdconfigprotocolany", cnpdthresholdconfigentry.Cnpdthresholdconfigprotocolany}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigStatsSelect"] = types.YLeaf{"Cnpdthresholdconfigstatsselect", cnpdthresholdconfigentry.Cnpdthresholdconfigstatsselect}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigStartup"] = types.YLeaf{"Cnpdthresholdconfigstartup", cnpdthresholdconfigentry.Cnpdthresholdconfigstartup}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigRising"] = types.YLeaf{"Cnpdthresholdconfigrising", cnpdthresholdconfigentry.Cnpdthresholdconfigrising}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigFalling"] = types.YLeaf{"Cnpdthresholdconfigfalling", cnpdthresholdconfigentry.Cnpdthresholdconfigfalling}
    cnpdthresholdconfigentry.EntityData.Leafs["cnpdThresholdConfigStatus"] = types.YLeaf{"Cnpdthresholdconfigstatus", cnpdthresholdconfigentry.Cnpdthresholdconfigstatus}
    return &(cnpdthresholdconfigentry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigsampletype represents cnpdThresholdConfigFalling value.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigsampletype string

const (
    CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigsampletype_absoluteValue CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigsampletype = "absoluteValue"

    CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigsampletype_deltaValue CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigsampletype = "deltaValue"
)

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigstartup represents will be generated.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigstartup string

const (
    CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigstartup_rising CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigstartup = "rising"

    CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigstartup_falling CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigstartup = "falling"

    CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigstartup_risingOrFalling CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry_Cnpdthresholdconfigstartup = "risingOrFalling"
)

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable
// The Threshold History table. Notifications
// are unreliable so this table provides a
// history of the last 5000 threshold breached
// events. A notification can be traced back to
// its cnpdThresholdHistoryEntry.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry is created each time a threshold  is breached.   Thus there is
    // not necessarily a one to one  relationship to cnpdThresholdConfigTable  as
    // not every Threshold configured will  be breached. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry.
    Cnpdthresholdhistoryentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry
}

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetEntityData() *types.CommonEntityData {
    cnpdthresholdhistorytable.EntityData.YFilter = cnpdthresholdhistorytable.YFilter
    cnpdthresholdhistorytable.EntityData.YangName = "cnpdThresholdHistoryTable"
    cnpdthresholdhistorytable.EntityData.BundleName = "cisco_ios_xe"
    cnpdthresholdhistorytable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdthresholdhistorytable.EntityData.SegmentPath = "cnpdThresholdHistoryTable"
    cnpdthresholdhistorytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdthresholdhistorytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdthresholdhistorytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdthresholdhistorytable.EntityData.Children = make(map[string]types.YChild)
    cnpdthresholdhistorytable.EntityData.Children["cnpdThresholdHistoryEntry"] = types.YChild{"Cnpdthresholdhistoryentry", nil}
    for i := range cnpdthresholdhistorytable.Cnpdthresholdhistoryentry {
        cnpdthresholdhistorytable.EntityData.Children[types.GetSegmentPath(&cnpdthresholdhistorytable.Cnpdthresholdhistoryentry[i])] = types.YChild{"Cnpdthresholdhistoryentry", &cnpdthresholdhistorytable.Cnpdthresholdhistoryentry[i]}
    }
    cnpdthresholdhistorytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnpdthresholdhistorytable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry
// This entry is created each time a threshold 
// is breached. 
// 
// Thus there is not necessarily a one to one 
// relationship to cnpdThresholdConfigTable 
// as not every Threshold configured will 
// be breached.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A monotonically increasing integer which uniquely
    // identifies this  cnpdThresholdHistoryEntry in the  cnpdThresholdHistory
    // table. The type is interface{} with range: 1..1000.
    Cnpdthresholdhistoryindex interface{}

    // The cnpdThresholdConfigTable entry  which generated this entry. Using this 
    // object the management station can backtrack  to the appropriate
    // cnpdThresholdConfigEntry. The type is interface{} with range: 1..1000.
    Cnpdthresholdhistoryconfigindex interface{}

    // The actual value of the statistic when  the sampling was made. The type is
    // interface{} with range: 1..4294967295.
    Cnpdthresholdhistoryvalue interface{}

    // Describes whether this is an event caused by a rising or falling threshold
    // breach. The type is Cnpdthresholdhistorytype.
    Cnpdthresholdhistorytype interface{}

    // The value of sysUpTime of the running  configuration when the event
    // occurred. The type is interface{} with range: 0..4294967295.
    Cnpdthresholdhistorytime interface{}

    // The application or protocol which the management station configured a
    // threshold on.  This object is an index into the  SupportedProtocolsTable
    // where details of the protocol can be found. The type is interface{} with
    // range: 1..1024.
    Cnpdthresholdhistoryprotocol interface{}

    // This is the statistic used to base the threshold on. The type is
    // CiscoPdDataType.
    Cnpdthresholdhistorystatsselect interface{}
}

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetEntityData() *types.CommonEntityData {
    cnpdthresholdhistoryentry.EntityData.YFilter = cnpdthresholdhistoryentry.YFilter
    cnpdthresholdhistoryentry.EntityData.YangName = "cnpdThresholdHistoryEntry"
    cnpdthresholdhistoryentry.EntityData.BundleName = "cisco_ios_xe"
    cnpdthresholdhistoryentry.EntityData.ParentYangName = "cnpdThresholdHistoryTable"
    cnpdthresholdhistoryentry.EntityData.SegmentPath = "cnpdThresholdHistoryEntry" + "[cnpdThresholdHistoryIndex='" + fmt.Sprintf("%v", cnpdthresholdhistoryentry.Cnpdthresholdhistoryindex) + "']"
    cnpdthresholdhistoryentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdthresholdhistoryentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdthresholdhistoryentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdthresholdhistoryentry.EntityData.Children = make(map[string]types.YChild)
    cnpdthresholdhistoryentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnpdthresholdhistoryentry.EntityData.Leafs["cnpdThresholdHistoryIndex"] = types.YLeaf{"Cnpdthresholdhistoryindex", cnpdthresholdhistoryentry.Cnpdthresholdhistoryindex}
    cnpdthresholdhistoryentry.EntityData.Leafs["cnpdThresholdHistoryConfigIndex"] = types.YLeaf{"Cnpdthresholdhistoryconfigindex", cnpdthresholdhistoryentry.Cnpdthresholdhistoryconfigindex}
    cnpdthresholdhistoryentry.EntityData.Leafs["cnpdThresholdHistoryValue"] = types.YLeaf{"Cnpdthresholdhistoryvalue", cnpdthresholdhistoryentry.Cnpdthresholdhistoryvalue}
    cnpdthresholdhistoryentry.EntityData.Leafs["cnpdThresholdHistoryType"] = types.YLeaf{"Cnpdthresholdhistorytype", cnpdthresholdhistoryentry.Cnpdthresholdhistorytype}
    cnpdthresholdhistoryentry.EntityData.Leafs["cnpdThresholdHistoryTime"] = types.YLeaf{"Cnpdthresholdhistorytime", cnpdthresholdhistoryentry.Cnpdthresholdhistorytime}
    cnpdthresholdhistoryentry.EntityData.Leafs["cnpdThresholdHistoryProtocol"] = types.YLeaf{"Cnpdthresholdhistoryprotocol", cnpdthresholdhistoryentry.Cnpdthresholdhistoryprotocol}
    cnpdthresholdhistoryentry.EntityData.Leafs["cnpdThresholdHistoryStatsSelect"] = types.YLeaf{"Cnpdthresholdhistorystatsselect", cnpdthresholdhistoryentry.Cnpdthresholdhistorystatsselect}
    return &(cnpdthresholdhistoryentry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry_Cnpdthresholdhistorytype represents or falling threshold breach.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry_Cnpdthresholdhistorytype string

const (
    CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry_Cnpdthresholdhistorytype_risingBreach CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry_Cnpdthresholdhistorytype = "risingBreach"

    CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry_Cnpdthresholdhistorytype_fallingBreach CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry_Cnpdthresholdhistorytype = "fallingBreach"
)

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable
// The Supported Protocols table lists all the 
// protocols and applications which NBAR is currently
// capable of recognizing.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A entry in the Supported Protocols table reflecting key information about a
    // protocol. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry.
    Cnpdsupportedprotocolsentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry
}

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetEntityData() *types.CommonEntityData {
    cnpdsupportedprotocolstable.EntityData.YFilter = cnpdsupportedprotocolstable.YFilter
    cnpdsupportedprotocolstable.EntityData.YangName = "cnpdSupportedProtocolsTable"
    cnpdsupportedprotocolstable.EntityData.BundleName = "cisco_ios_xe"
    cnpdsupportedprotocolstable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdsupportedprotocolstable.EntityData.SegmentPath = "cnpdSupportedProtocolsTable"
    cnpdsupportedprotocolstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdsupportedprotocolstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdsupportedprotocolstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdsupportedprotocolstable.EntityData.Children = make(map[string]types.YChild)
    cnpdsupportedprotocolstable.EntityData.Children["cnpdSupportedProtocolsEntry"] = types.YChild{"Cnpdsupportedprotocolsentry", nil}
    for i := range cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry {
        cnpdsupportedprotocolstable.EntityData.Children[types.GetSegmentPath(&cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry[i])] = types.YChild{"Cnpdsupportedprotocolsentry", &cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry[i]}
    }
    cnpdsupportedprotocolstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cnpdsupportedprotocolstable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry
// A entry in the Supported Protocols table reflecting
// key information about a protocol.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique identifier of a row in this table.  Thus
    // it also represents a unique identifier for a protocol or application which
    // NBAR currently recognizes. The type is interface{} with range: 1..1024.
    Cnpdsupportedprotocolsindex interface{}

    // This object reflects the valid string of a protocol or application which
    // NBAR currently recognizes. The type is string with length: 1..255.
    Cnpdsupportedprotocolsname interface{}
}

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetEntityData() *types.CommonEntityData {
    cnpdsupportedprotocolsentry.EntityData.YFilter = cnpdsupportedprotocolsentry.YFilter
    cnpdsupportedprotocolsentry.EntityData.YangName = "cnpdSupportedProtocolsEntry"
    cnpdsupportedprotocolsentry.EntityData.BundleName = "cisco_ios_xe"
    cnpdsupportedprotocolsentry.EntityData.ParentYangName = "cnpdSupportedProtocolsTable"
    cnpdsupportedprotocolsentry.EntityData.SegmentPath = "cnpdSupportedProtocolsEntry" + "[cnpdSupportedProtocolsIndex='" + fmt.Sprintf("%v", cnpdsupportedprotocolsentry.Cnpdsupportedprotocolsindex) + "']"
    cnpdsupportedprotocolsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdsupportedprotocolsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdsupportedprotocolsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdsupportedprotocolsentry.EntityData.Children = make(map[string]types.YChild)
    cnpdsupportedprotocolsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cnpdsupportedprotocolsentry.EntityData.Leafs["cnpdSupportedProtocolsIndex"] = types.YLeaf{"Cnpdsupportedprotocolsindex", cnpdsupportedprotocolsentry.Cnpdsupportedprotocolsindex}
    cnpdsupportedprotocolsentry.EntityData.Leafs["cnpdSupportedProtocolsName"] = types.YLeaf{"Cnpdsupportedprotocolsname", cnpdsupportedprotocolsentry.Cnpdsupportedprotocolsname}
    return &(cnpdsupportedprotocolsentry.EntityData)
}

