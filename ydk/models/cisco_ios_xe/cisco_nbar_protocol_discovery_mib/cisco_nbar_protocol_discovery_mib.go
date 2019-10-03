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

    
    CnpdNotificationsConfig CISCONBARPROTOCOLDISCOVERYMIB_CnpdNotificationsConfig

    // The cnpdStatusTable is used to enable and disable Protocol Discovery on an
    // interface.
    CnpdStatusTable CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable

    // The cnpdAllStatsTable contains all the statistics available for all the
    // protocols/applications currently recognized by NBAR Protocol Discovery for
    // a particular  interface.  In the event of an overflow, the 32 bit counters
    // are not  valid. There is no overflow support.
    CnpdAllStatsTable CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable

    // The cnpdTopNConfigTable is used to configure cnpdTopNStatsTable's.
    CnpdTopNConfigTable CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable

    // A cnpdTopNStatsTable describes an ordered list of protocols.
    CnpdTopNStatsTable CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable

    // The cnpdThresholdConfigTable allows the management station to create
    // thresholds for the purpose of sending notifications if breached, and
    // creating a history of breached thresholds.
    CnpdThresholdConfigTable CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable

    // The Threshold History table. Notifications are unreliable so this table
    // provides a history of the last 5000 threshold breached events. A
    // notification can be traced back to its cnpdThresholdHistoryEntry.
    CnpdThresholdHistoryTable CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable

    // The Supported Protocols table lists all the  protocols and applications
    // which NBAR is currently capable of recognizing.
    CnpdSupportedProtocolsTable CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable
}

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetEntityData() *types.CommonEntityData {
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.YFilter = cISCONBARPROTOCOLDISCOVERYMIB.YFilter
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.YangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.SegmentPath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.AbsolutePath = cISCONBARPROTOCOLDISCOVERYMIB.EntityData.SegmentPath
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children = types.NewOrderedMap()
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children.Append("cnpdNotificationsConfig", types.YChild{"CnpdNotificationsConfig", &cISCONBARPROTOCOLDISCOVERYMIB.CnpdNotificationsConfig})
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children.Append("cnpdStatusTable", types.YChild{"CnpdStatusTable", &cISCONBARPROTOCOLDISCOVERYMIB.CnpdStatusTable})
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children.Append("cnpdAllStatsTable", types.YChild{"CnpdAllStatsTable", &cISCONBARPROTOCOLDISCOVERYMIB.CnpdAllStatsTable})
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children.Append("cnpdTopNConfigTable", types.YChild{"CnpdTopNConfigTable", &cISCONBARPROTOCOLDISCOVERYMIB.CnpdTopNConfigTable})
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children.Append("cnpdTopNStatsTable", types.YChild{"CnpdTopNStatsTable", &cISCONBARPROTOCOLDISCOVERYMIB.CnpdTopNStatsTable})
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children.Append("cnpdThresholdConfigTable", types.YChild{"CnpdThresholdConfigTable", &cISCONBARPROTOCOLDISCOVERYMIB.CnpdThresholdConfigTable})
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children.Append("cnpdThresholdHistoryTable", types.YChild{"CnpdThresholdHistoryTable", &cISCONBARPROTOCOLDISCOVERYMIB.CnpdThresholdHistoryTable})
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Children.Append("cnpdSupportedProtocolsTable", types.YChild{"CnpdSupportedProtocolsTable", &cISCONBARPROTOCOLDISCOVERYMIB.CnpdSupportedProtocolsTable})
    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCONBARPROTOCOLDISCOVERYMIB.EntityData.YListKeys = []string {}

    return &(cISCONBARPROTOCOLDISCOVERYMIB.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdNotificationsConfig
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdNotificationsConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object is used to enable or disable  Notifications on a global basis. 
    // If set to 'true' - Notifications are enabled. If set to 'false' -
    // Notifications are disabled. The type is bool.
    CnpdNotificationsEnable interface{}
}

func (cnpdNotificationsConfig *CISCONBARPROTOCOLDISCOVERYMIB_CnpdNotificationsConfig) GetEntityData() *types.CommonEntityData {
    cnpdNotificationsConfig.EntityData.YFilter = cnpdNotificationsConfig.YFilter
    cnpdNotificationsConfig.EntityData.YangName = "cnpdNotificationsConfig"
    cnpdNotificationsConfig.EntityData.BundleName = "cisco_ios_xe"
    cnpdNotificationsConfig.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdNotificationsConfig.EntityData.SegmentPath = "cnpdNotificationsConfig"
    cnpdNotificationsConfig.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/" + cnpdNotificationsConfig.EntityData.SegmentPath
    cnpdNotificationsConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdNotificationsConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdNotificationsConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdNotificationsConfig.EntityData.Children = types.NewOrderedMap()
    cnpdNotificationsConfig.EntityData.Leafs = types.NewOrderedMap()
    cnpdNotificationsConfig.EntityData.Leafs.Append("cnpdNotificationsEnable", types.YLeaf{"CnpdNotificationsEnable", cnpdNotificationsConfig.CnpdNotificationsEnable})

    cnpdNotificationsConfig.EntityData.YListKeys = []string {}

    return &(cnpdNotificationsConfig.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable
// The cnpdStatusTable is used to enable and
// disable Protocol Discovery on an interface.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cnpdStatusTable contains objects for enabling or disabling
    // Protocol Discovery on a per interface basis. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable_CnpdStatusEntry.
    CnpdStatusEntry []*CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable_CnpdStatusEntry
}

func (cnpdStatusTable *CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable) GetEntityData() *types.CommonEntityData {
    cnpdStatusTable.EntityData.YFilter = cnpdStatusTable.YFilter
    cnpdStatusTable.EntityData.YangName = "cnpdStatusTable"
    cnpdStatusTable.EntityData.BundleName = "cisco_ios_xe"
    cnpdStatusTable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdStatusTable.EntityData.SegmentPath = "cnpdStatusTable"
    cnpdStatusTable.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/" + cnpdStatusTable.EntityData.SegmentPath
    cnpdStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdStatusTable.EntityData.Children = types.NewOrderedMap()
    cnpdStatusTable.EntityData.Children.Append("cnpdStatusEntry", types.YChild{"CnpdStatusEntry", nil})
    for i := range cnpdStatusTable.CnpdStatusEntry {
        cnpdStatusTable.EntityData.Children.Append(types.GetSegmentPath(cnpdStatusTable.CnpdStatusEntry[i]), types.YChild{"CnpdStatusEntry", cnpdStatusTable.CnpdStatusEntry[i]})
    }
    cnpdStatusTable.EntityData.Leafs = types.NewOrderedMap()

    cnpdStatusTable.EntityData.YListKeys = []string {}

    return &(cnpdStatusTable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable_CnpdStatusEntry
// An entry in the cnpdStatusTable contains objects
// for enabling or disabling Protocol Discovery on a
// per interface basis.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable_CnpdStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object is used to enable or disable  Protocol Discovery on an
    // interface.   If set to 'true' - Protocol Discovery is  enabled on this
    // Interface.  If set to 'false' - Protocol Discovery is  disabled on this
    // Interface. The type is bool.
    CnpdStatusPdEnable interface{}

    // The value of sysUpTime at the time Protocol  Discovery was last enabled  on
    // an interface. If the interface does not have Protocol Discovery enabled
    // this value is zero. The type is interface{} with range: 0..4294967295.
    CnpdStatusLastUpdateTime interface{}
}

func (cnpdStatusEntry *CISCONBARPROTOCOLDISCOVERYMIB_CnpdStatusTable_CnpdStatusEntry) GetEntityData() *types.CommonEntityData {
    cnpdStatusEntry.EntityData.YFilter = cnpdStatusEntry.YFilter
    cnpdStatusEntry.EntityData.YangName = "cnpdStatusEntry"
    cnpdStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    cnpdStatusEntry.EntityData.ParentYangName = "cnpdStatusTable"
    cnpdStatusEntry.EntityData.SegmentPath = "cnpdStatusEntry" + types.AddKeyToken(cnpdStatusEntry.IfIndex, "ifIndex")
    cnpdStatusEntry.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/cnpdStatusTable/" + cnpdStatusEntry.EntityData.SegmentPath
    cnpdStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdStatusEntry.EntityData.Children = types.NewOrderedMap()
    cnpdStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    cnpdStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cnpdStatusEntry.IfIndex})
    cnpdStatusEntry.EntityData.Leafs.Append("cnpdStatusPdEnable", types.YLeaf{"CnpdStatusPdEnable", cnpdStatusEntry.CnpdStatusPdEnable})
    cnpdStatusEntry.EntityData.Leafs.Append("cnpdStatusLastUpdateTime", types.YLeaf{"CnpdStatusLastUpdateTime", cnpdStatusEntry.CnpdStatusLastUpdateTime})

    cnpdStatusEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(cnpdStatusEntry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable
// The cnpdAllStatsTable contains all the statistics
// available for all the protocols/applications currently
// recognized by NBAR Protocol Discovery for a particular 
// interface.
// 
// In the event of an overflow, the 32 bit counters are not 
// valid. There is no overflow support.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the cnpdAllStatsTable table. This entry  contains the
    // statistics collected on all the protocols  which NBAR classifies for a
    // particular interface. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable_CnpdAllStatsEntry.
    CnpdAllStatsEntry []*CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable_CnpdAllStatsEntry
}

func (cnpdAllStatsTable *CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable) GetEntityData() *types.CommonEntityData {
    cnpdAllStatsTable.EntityData.YFilter = cnpdAllStatsTable.YFilter
    cnpdAllStatsTable.EntityData.YangName = "cnpdAllStatsTable"
    cnpdAllStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cnpdAllStatsTable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdAllStatsTable.EntityData.SegmentPath = "cnpdAllStatsTable"
    cnpdAllStatsTable.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/" + cnpdAllStatsTable.EntityData.SegmentPath
    cnpdAllStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdAllStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdAllStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdAllStatsTable.EntityData.Children = types.NewOrderedMap()
    cnpdAllStatsTable.EntityData.Children.Append("cnpdAllStatsEntry", types.YChild{"CnpdAllStatsEntry", nil})
    for i := range cnpdAllStatsTable.CnpdAllStatsEntry {
        cnpdAllStatsTable.EntityData.Children.Append(types.GetSegmentPath(cnpdAllStatsTable.CnpdAllStatsEntry[i]), types.YChild{"CnpdAllStatsEntry", cnpdAllStatsTable.CnpdAllStatsEntry[i]})
    }
    cnpdAllStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cnpdAllStatsTable.EntityData.YListKeys = []string {}

    return &(cnpdAllStatsTable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable_CnpdAllStatsEntry
// An entry in the cnpdAllStatsTable table. This entry 
// contains the statistics collected on all the protocols 
// which NBAR classifies for a particular interface.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable_CnpdAllStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. An object which represents a unique  identifier
    // for a protocol or application  which NBAR currently recognizes.  This
    // object is an index into the  SupportedProtocolsTable where details of the
    // protocol can be found. The type is interface{} with range: 1..1024.
    CnpdAllStatsProtocolsIndex interface{}

    // Name of the application or protocol, a  unique textual string, assigned in
    // the cnpdSupportedProtocolsTable. The type is string with length: 1..255.
    CnpdAllStatsProtocolName interface{}

    // The packet count of inbound packets as  determined by Protocol Discovery.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    CnpdAllStatsInPkts interface{}

    // The packet count of outbound packets as  determined by Protocol Discovery.
    // The type is interface{} with range: 0..4294967295. Units are packets.
    CnpdAllStatsOutPkts interface{}

    // The byte count of inbound octets as  determined by Protocol Discovery. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    CnpdAllStatsInBytes interface{}

    // The byte count of outbound octets as determined by Protocol Discovery. The
    // type is interface{} with range: 0..4294967295. Units are bytes.
    CnpdAllStatsOutBytes interface{}

    // The packet count of inbound packets as  determined by Protocol Discovery.
    // This is the 64-bit (High Capacity) version of cnpdAllStatsInPkts. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    CnpdAllStatsHCInPkts interface{}

    // The packet count of outbound packets as  determined by Protocol Discovery.
    // This is the 64-bit (High Capacity) version of cnpdAllStatsOutPkts. The type
    // is interface{} with range: 0..18446744073709551615. Units are packets.
    CnpdAllStatsHCOutPkts interface{}

    // The byte count of inbound octets as  determined by Protocol Discovery. This
    // is the 64-bit (High Capacity) version of cnpdAllStatsInBytes. The type is
    // interface{} with range: 0..18446744073709551615. Units are bytes.
    CnpdAllStatsHCInBytes interface{}

    // The byte count of outbound octets as  determined by Protocol Discovery.
    // This is the 64-bit (High Capacity) version of cnpdAllStatsOutBytes. The
    // type is interface{} with range: 0..18446744073709551615. Units are bytes.
    CnpdAllStatsHCOutBytes interface{}

    // The inbound bit rate as determined  by Protocol Discovery. The type is
    // interface{} with range: 1..4294967295. Units are kilo bits per second.
    CnpdAllStatsInBitRate interface{}

    // The outbound bit rate as determined  by Protocol Discovery. The type is
    // interface{} with range: 1..4294967295. Units are kilo bits per second.
    CnpdAllStatsOutBitRate interface{}
}

func (cnpdAllStatsEntry *CISCONBARPROTOCOLDISCOVERYMIB_CnpdAllStatsTable_CnpdAllStatsEntry) GetEntityData() *types.CommonEntityData {
    cnpdAllStatsEntry.EntityData.YFilter = cnpdAllStatsEntry.YFilter
    cnpdAllStatsEntry.EntityData.YangName = "cnpdAllStatsEntry"
    cnpdAllStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cnpdAllStatsEntry.EntityData.ParentYangName = "cnpdAllStatsTable"
    cnpdAllStatsEntry.EntityData.SegmentPath = "cnpdAllStatsEntry" + types.AddKeyToken(cnpdAllStatsEntry.IfIndex, "ifIndex") + types.AddKeyToken(cnpdAllStatsEntry.CnpdAllStatsProtocolsIndex, "cnpdAllStatsProtocolsIndex")
    cnpdAllStatsEntry.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/cnpdAllStatsTable/" + cnpdAllStatsEntry.EntityData.SegmentPath
    cnpdAllStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdAllStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdAllStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdAllStatsEntry.EntityData.Children = types.NewOrderedMap()
    cnpdAllStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cnpdAllStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cnpdAllStatsEntry.IfIndex})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsProtocolsIndex", types.YLeaf{"CnpdAllStatsProtocolsIndex", cnpdAllStatsEntry.CnpdAllStatsProtocolsIndex})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsProtocolName", types.YLeaf{"CnpdAllStatsProtocolName", cnpdAllStatsEntry.CnpdAllStatsProtocolName})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsInPkts", types.YLeaf{"CnpdAllStatsInPkts", cnpdAllStatsEntry.CnpdAllStatsInPkts})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsOutPkts", types.YLeaf{"CnpdAllStatsOutPkts", cnpdAllStatsEntry.CnpdAllStatsOutPkts})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsInBytes", types.YLeaf{"CnpdAllStatsInBytes", cnpdAllStatsEntry.CnpdAllStatsInBytes})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsOutBytes", types.YLeaf{"CnpdAllStatsOutBytes", cnpdAllStatsEntry.CnpdAllStatsOutBytes})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsHCInPkts", types.YLeaf{"CnpdAllStatsHCInPkts", cnpdAllStatsEntry.CnpdAllStatsHCInPkts})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsHCOutPkts", types.YLeaf{"CnpdAllStatsHCOutPkts", cnpdAllStatsEntry.CnpdAllStatsHCOutPkts})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsHCInBytes", types.YLeaf{"CnpdAllStatsHCInBytes", cnpdAllStatsEntry.CnpdAllStatsHCInBytes})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsHCOutBytes", types.YLeaf{"CnpdAllStatsHCOutBytes", cnpdAllStatsEntry.CnpdAllStatsHCOutBytes})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsInBitRate", types.YLeaf{"CnpdAllStatsInBitRate", cnpdAllStatsEntry.CnpdAllStatsInBitRate})
    cnpdAllStatsEntry.EntityData.Leafs.Append("cnpdAllStatsOutBitRate", types.YLeaf{"CnpdAllStatsOutBitRate", cnpdAllStatsEntry.CnpdAllStatsOutBitRate})

    cnpdAllStatsEntry.EntityData.YListKeys = []string {"IfIndex", "CnpdAllStatsProtocolsIndex"}

    return &(cnpdAllStatsEntry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable
// The cnpdTopNConfigTable is used to configure
// cnpdTopNStatsTable's.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry provides the objects to configure and thus initiate the
    // generation of a cnpdTopNStatsTable.. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable_CnpdTopNConfigEntry.
    CnpdTopNConfigEntry []*CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable_CnpdTopNConfigEntry
}

func (cnpdTopNConfigTable *CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable) GetEntityData() *types.CommonEntityData {
    cnpdTopNConfigTable.EntityData.YFilter = cnpdTopNConfigTable.YFilter
    cnpdTopNConfigTable.EntityData.YangName = "cnpdTopNConfigTable"
    cnpdTopNConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cnpdTopNConfigTable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdTopNConfigTable.EntityData.SegmentPath = "cnpdTopNConfigTable"
    cnpdTopNConfigTable.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/" + cnpdTopNConfigTable.EntityData.SegmentPath
    cnpdTopNConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdTopNConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdTopNConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdTopNConfigTable.EntityData.Children = types.NewOrderedMap()
    cnpdTopNConfigTable.EntityData.Children.Append("cnpdTopNConfigEntry", types.YChild{"CnpdTopNConfigEntry", nil})
    for i := range cnpdTopNConfigTable.CnpdTopNConfigEntry {
        cnpdTopNConfigTable.EntityData.Children.Append(types.GetSegmentPath(cnpdTopNConfigTable.CnpdTopNConfigEntry[i]), types.YChild{"CnpdTopNConfigEntry", cnpdTopNConfigTable.CnpdTopNConfigEntry[i]})
    }
    cnpdTopNConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cnpdTopNConfigTable.EntityData.YListKeys = []string {}

    return &(cnpdTopNConfigTable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable_CnpdTopNConfigEntry
// This entry provides the objects to configure and thus
// initiate the generation of a cnpdTopNStatsTable..
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable_CnpdTopNConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A monotonically increasing integer which uniquely
    // identifies a cnpdTopNConfigEntry  in the cnpdTopNConfigTable. The type is
    // interface{} with range: 1..50.
    CnpdTopNConfigIndex interface{}

    // This object allows the management station to select the interface, which
    // Protocol Discovery is running on, to be used to create this 
    // cnpdTopNConfigEntry. The type is interface{} with range: 1..2147483647.
    CnpdTopNConfigIfIndex interface{}

    // This object allows the management station to select the statistic used to
    // base the order of the top-n table on.  For example: a
    // cnpdTopNConfigStatsSelect of bitRateSum means order this table based on
    // each applications/protocols combined in and out bitrate. The type is
    // CiscoPdDataType.
    CnpdTopNConfigStatsSelect interface{}

    // If the cnpdTopNConfigStatsSelect is bitRateIn, bitRateOut or bitRateSum,
    // then this value is the interval in seconds that  the bitrate is sampled. 
    // This has no effect if the cnpdTopNConfigStatsSelect is byte or packet
    // based.  When this object is modified by the management  station, a new
    // sample period is started regardless of whether the original
    // cnpdTopNConfigSampleTime was finished. The type is interface{} with range:
    // 1..2048. Units are seconds.
    CnpdTopNConfigSampleTime interface{}

    // The requested size of the associated  cnpdTopNStatsTable entry.  For
    // example a cnpdTopNConfigRequestedSize of 20 indicates the management
    // station wants to create an associated  cnpdTopNStatsTable  entry of 20
    // protocol/application's. The type is interface{} with range: 1..500.
    CnpdTopNConfigRequestedSize interface{}

    // The actual size of the associated 	 cnpdTopNStatsTable entry.  The reason
    // this may differ from  cnpdTopNConfigRequestedSize is because a  management
    // station may request a number of  protocols that is greater than the number
    // of  protocols actually found by Protocol Discovery. The type is interface{}
    // with range: 1..500.
    CnpdTopNConfigGrantedSize interface{}

    // The value of sysUpTime when the associated cnpdTopNStatsTable entry was
    // created. The type is interface{} with range: 0..4294967295.
    CnpdTopNConfigTime interface{}

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
    CnpdTopNConfigStatus interface{}
}

func (cnpdTopNConfigEntry *CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable_CnpdTopNConfigEntry) GetEntityData() *types.CommonEntityData {
    cnpdTopNConfigEntry.EntityData.YFilter = cnpdTopNConfigEntry.YFilter
    cnpdTopNConfigEntry.EntityData.YangName = "cnpdTopNConfigEntry"
    cnpdTopNConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cnpdTopNConfigEntry.EntityData.ParentYangName = "cnpdTopNConfigTable"
    cnpdTopNConfigEntry.EntityData.SegmentPath = "cnpdTopNConfigEntry" + types.AddKeyToken(cnpdTopNConfigEntry.CnpdTopNConfigIndex, "cnpdTopNConfigIndex")
    cnpdTopNConfigEntry.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/cnpdTopNConfigTable/" + cnpdTopNConfigEntry.EntityData.SegmentPath
    cnpdTopNConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdTopNConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdTopNConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdTopNConfigEntry.EntityData.Children = types.NewOrderedMap()
    cnpdTopNConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cnpdTopNConfigEntry.EntityData.Leafs.Append("cnpdTopNConfigIndex", types.YLeaf{"CnpdTopNConfigIndex", cnpdTopNConfigEntry.CnpdTopNConfigIndex})
    cnpdTopNConfigEntry.EntityData.Leafs.Append("cnpdTopNConfigIfIndex", types.YLeaf{"CnpdTopNConfigIfIndex", cnpdTopNConfigEntry.CnpdTopNConfigIfIndex})
    cnpdTopNConfigEntry.EntityData.Leafs.Append("cnpdTopNConfigStatsSelect", types.YLeaf{"CnpdTopNConfigStatsSelect", cnpdTopNConfigEntry.CnpdTopNConfigStatsSelect})
    cnpdTopNConfigEntry.EntityData.Leafs.Append("cnpdTopNConfigSampleTime", types.YLeaf{"CnpdTopNConfigSampleTime", cnpdTopNConfigEntry.CnpdTopNConfigSampleTime})
    cnpdTopNConfigEntry.EntityData.Leafs.Append("cnpdTopNConfigRequestedSize", types.YLeaf{"CnpdTopNConfigRequestedSize", cnpdTopNConfigEntry.CnpdTopNConfigRequestedSize})
    cnpdTopNConfigEntry.EntityData.Leafs.Append("cnpdTopNConfigGrantedSize", types.YLeaf{"CnpdTopNConfigGrantedSize", cnpdTopNConfigEntry.CnpdTopNConfigGrantedSize})
    cnpdTopNConfigEntry.EntityData.Leafs.Append("cnpdTopNConfigTime", types.YLeaf{"CnpdTopNConfigTime", cnpdTopNConfigEntry.CnpdTopNConfigTime})
    cnpdTopNConfigEntry.EntityData.Leafs.Append("cnpdTopNConfigStatus", types.YLeaf{"CnpdTopNConfigStatus", cnpdTopNConfigEntry.CnpdTopNConfigStatus})

    cnpdTopNConfigEntry.EntityData.YListKeys = []string {"CnpdTopNConfigIndex"}

    return &(cnpdTopNConfigEntry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable
// A cnpdTopNStatsTable describes an ordered
// list of protocols.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry is used to store a set of objects which  describe a
    // cnpdTopNStatsTable. A cnpdTopNStatsTable  is a number of protocols and
    // statistics sorted  according to the criteria in the associated
    // cnpdTopNConfigEntry.  Therefore a cnpdTopNStatsTable can differ in content 
    // and size according to what was configured in the associated
    // cnpdTopNConfigTableEntry. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable_CnpdTopNStatsEntry.
    CnpdTopNStatsEntry []*CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable_CnpdTopNStatsEntry
}

func (cnpdTopNStatsTable *CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable) GetEntityData() *types.CommonEntityData {
    cnpdTopNStatsTable.EntityData.YFilter = cnpdTopNStatsTable.YFilter
    cnpdTopNStatsTable.EntityData.YangName = "cnpdTopNStatsTable"
    cnpdTopNStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cnpdTopNStatsTable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdTopNStatsTable.EntityData.SegmentPath = "cnpdTopNStatsTable"
    cnpdTopNStatsTable.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/" + cnpdTopNStatsTable.EntityData.SegmentPath
    cnpdTopNStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdTopNStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdTopNStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdTopNStatsTable.EntityData.Children = types.NewOrderedMap()
    cnpdTopNStatsTable.EntityData.Children.Append("cnpdTopNStatsEntry", types.YChild{"CnpdTopNStatsEntry", nil})
    for i := range cnpdTopNStatsTable.CnpdTopNStatsEntry {
        cnpdTopNStatsTable.EntityData.Children.Append(types.GetSegmentPath(cnpdTopNStatsTable.CnpdTopNStatsEntry[i]), types.YChild{"CnpdTopNStatsEntry", cnpdTopNStatsTable.CnpdTopNStatsEntry[i]})
    }
    cnpdTopNStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cnpdTopNStatsTable.EntityData.YListKeys = []string {}

    return &(cnpdTopNStatsTable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable_CnpdTopNStatsEntry
// This entry is used to store a set of objects which 
// describe a cnpdTopNStatsTable. A cnpdTopNStatsTable 
// is a number of protocols and statistics sorted 
// according to the criteria in the associated
// cnpdTopNConfigEntry.
// 
// Therefore a cnpdTopNStatsTable can differ in content 
// and size according to what was configured in the associated
// cnpdTopNConfigTableEntry.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable_CnpdTopNStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..50. Refers to
    // cisco_nbar_protocol_discovery_mib.CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNConfigTable_CnpdTopNConfigEntry_CnpdTopNConfigIndex
    CnpdTopNConfigIndex interface{}

    // This attribute is a key. A monotonically increasing integer which  uniquely
    // identifies a cnpdTopNStatsEntry  in the cnpdTopNStatsTable. The type is
    // interface{} with range: 1..500.
    CnpdTopNStatsIndex interface{}

    // Name of the application or protocol,  a unique textual string, assigned in
    // the cnpdSupportedProtocolsTable. The type is string with length: 1..255.
    CnpdTopNStatsProtocolName interface{}

    // The amount of change in the selected statistic during this sampling
    // interval. The selected statistic is the cnpdTopNConfigStatsSelect from the
    // associated cnpdTopNConfigStatsEntry. The type is interface{} with range:
    // 0..4294967295.
    CnpdTopNStatsRate interface{}

    // The amount of change in the selected statistic during this sampling
    // interval. The selected statistic is the cnpdTopNConfigStatsSelect from the
    // associated cnpdTopNConfigStatsEntry.	  This is the 64-bit (High Capacity)
    // version of  cnpdTopNStatsRate. The type is interface{} with range:
    // 0..18446744073709551615.
    CnpdTopNStatsHCRate interface{}
}

func (cnpdTopNStatsEntry *CISCONBARPROTOCOLDISCOVERYMIB_CnpdTopNStatsTable_CnpdTopNStatsEntry) GetEntityData() *types.CommonEntityData {
    cnpdTopNStatsEntry.EntityData.YFilter = cnpdTopNStatsEntry.YFilter
    cnpdTopNStatsEntry.EntityData.YangName = "cnpdTopNStatsEntry"
    cnpdTopNStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cnpdTopNStatsEntry.EntityData.ParentYangName = "cnpdTopNStatsTable"
    cnpdTopNStatsEntry.EntityData.SegmentPath = "cnpdTopNStatsEntry" + types.AddKeyToken(cnpdTopNStatsEntry.CnpdTopNConfigIndex, "cnpdTopNConfigIndex") + types.AddKeyToken(cnpdTopNStatsEntry.CnpdTopNStatsIndex, "cnpdTopNStatsIndex")
    cnpdTopNStatsEntry.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/cnpdTopNStatsTable/" + cnpdTopNStatsEntry.EntityData.SegmentPath
    cnpdTopNStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdTopNStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdTopNStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdTopNStatsEntry.EntityData.Children = types.NewOrderedMap()
    cnpdTopNStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cnpdTopNStatsEntry.EntityData.Leafs.Append("cnpdTopNConfigIndex", types.YLeaf{"CnpdTopNConfigIndex", cnpdTopNStatsEntry.CnpdTopNConfigIndex})
    cnpdTopNStatsEntry.EntityData.Leafs.Append("cnpdTopNStatsIndex", types.YLeaf{"CnpdTopNStatsIndex", cnpdTopNStatsEntry.CnpdTopNStatsIndex})
    cnpdTopNStatsEntry.EntityData.Leafs.Append("cnpdTopNStatsProtocolName", types.YLeaf{"CnpdTopNStatsProtocolName", cnpdTopNStatsEntry.CnpdTopNStatsProtocolName})
    cnpdTopNStatsEntry.EntityData.Leafs.Append("cnpdTopNStatsRate", types.YLeaf{"CnpdTopNStatsRate", cnpdTopNStatsEntry.CnpdTopNStatsRate})
    cnpdTopNStatsEntry.EntityData.Leafs.Append("cnpdTopNStatsHCRate", types.YLeaf{"CnpdTopNStatsHCRate", cnpdTopNStatsEntry.CnpdTopNStatsHCRate})

    cnpdTopNStatsEntry.EntityData.YListKeys = []string {"CnpdTopNConfigIndex", "CnpdTopNStatsIndex"}

    return &(cnpdTopNStatsEntry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable
// The cnpdThresholdConfigTable allows the management
// station to create thresholds for the purpose of
// sending notifications if breached, and creating a
// history of breached thresholds.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable struct {
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
    // CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry.
    CnpdThresholdConfigEntry []*CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry
}

func (cnpdThresholdConfigTable *CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable) GetEntityData() *types.CommonEntityData {
    cnpdThresholdConfigTable.EntityData.YFilter = cnpdThresholdConfigTable.YFilter
    cnpdThresholdConfigTable.EntityData.YangName = "cnpdThresholdConfigTable"
    cnpdThresholdConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cnpdThresholdConfigTable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdThresholdConfigTable.EntityData.SegmentPath = "cnpdThresholdConfigTable"
    cnpdThresholdConfigTable.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/" + cnpdThresholdConfigTable.EntityData.SegmentPath
    cnpdThresholdConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdThresholdConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdThresholdConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdThresholdConfigTable.EntityData.Children = types.NewOrderedMap()
    cnpdThresholdConfigTable.EntityData.Children.Append("cnpdThresholdConfigEntry", types.YChild{"CnpdThresholdConfigEntry", nil})
    for i := range cnpdThresholdConfigTable.CnpdThresholdConfigEntry {
        cnpdThresholdConfigTable.EntityData.Children.Append(types.GetSegmentPath(cnpdThresholdConfigTable.CnpdThresholdConfigEntry[i]), types.YChild{"CnpdThresholdConfigEntry", cnpdThresholdConfigTable.CnpdThresholdConfigEntry[i]})
    }
    cnpdThresholdConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cnpdThresholdConfigTable.EntityData.YListKeys = []string {}

    return &(cnpdThresholdConfigTable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry
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
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A monotonically increasing integer which  uniquely
    // identifies an entry in the  cnpdThresholdConfigTable. The type is
    // interface{} with range: 1..100.
    CnpdThresholdConfigIndex interface{}

    // This object allows the management station to  select the interface, which
    // Protocol Discovery is  running on, to be used to create this 
    // cnpdThresholdConfigTable entry. The type is interface{} with range:
    // 1..2147483647.
    CnpdThresholdConfigIfIndex interface{}

    // The interval in seconds over which the data is sampled and compared with
    // cnpdThresholdConfigRising and cnpdThresholdConfigFalling thresholds. The
    // type is interface{} with range: 1..2048. Units are seconds.
    CnpdThresholdConfigInterval interface{}

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
    // CnpdThresholdConfigSampleType.
    CnpdThresholdConfigSampleType interface{}

    // The application or protocol which the management station wishes to
    // configure a threshold on.  This object is an index into the 
    // SupportedProtocolsTable where details of the protocol can be found.  If
    // cnpdThresholdConfigProtocolAny is set to TRUE this value will be ignored.
    // If it is set to FALSE, then cnpdThresholdConfigProtocol will be the only
    // protocol that is checked to see if it has breached the threshold. The type
    // is interface{} with range: 1..1024.
    CnpdThresholdConfigProtocol interface{}

    // If set to 'true' - this threshold is configured to check for any protocol
    // which meets the threshold criteria. This means that multiple protocols can
    // generate ThresholdHistoryTable entries. Each protocol is subject to the
    // hysterisis mechanism.  If set to 'false' - this threshold is configured to
    // check for the protocol which meets the threshold criteria referred to by
    // cnpdThresholdConfigProtocol. The type is bool.
    CnpdThresholdConfigProtocolAny interface{}

    // This object allows the management station to select the statistic used to
    // base the threshold on.  For example a cnpdThresholdConfigStatsSelect of
    // bitRateSum means cnpdThresholdConfigRising and cnpdThresholdConfigFalling
    // are values based on the combined value of in and out bitrates. The type is
    // CiscoPdDataType.
    CnpdThresholdConfigStatsSelect interface{}

    // This controls the type of notification that is  sent when this threshold
    // entry is first enabled.   Because there is no previous sampling history,
    // choosing one of these options determines the type of notification generated
    // - Rising or Falling.  If the first sample after this entry is enabled  is
    // greater than or equal to cnpdThresholdConfigRising and this object is equal
    // to rising(1) or risingOrFalling(3),  then a single rising notification will
    // be generated.   If the first sample after this entry is enabled is less
    // than or equal to cnpdThresholdConfigFalling and this object is equal to
    // falling(2) or  risingOrFalling(3), then a single falling notification  will
    // be generated. The type is CnpdThresholdConfigStartup.
    CnpdThresholdConfigStartup interface{}

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
    CnpdThresholdConfigRising interface{}

    // This is the threshold object which the management  station sets to
    // determine if it gets breached. It  indicates the statistic being sampled
    // was falling.   When current sample is less than or equal  to this object,
    // and the value at the last sampling interval was greater than this object
    // (in other  words the value is falling), an entry in the 
    // cnpdThresholdHistoryTable will be created.  		 After a falling event is
    // generated, another  such event will not be generated until the sampled 
    // value rises above this object and reaches the cnpdThresholdConfigRising
    // value. The type is interface{} with range: 1..4294967295.
    CnpdThresholdConfigFalling interface{}

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
    CnpdThresholdConfigStatus interface{}
}

func (cnpdThresholdConfigEntry *CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry) GetEntityData() *types.CommonEntityData {
    cnpdThresholdConfigEntry.EntityData.YFilter = cnpdThresholdConfigEntry.YFilter
    cnpdThresholdConfigEntry.EntityData.YangName = "cnpdThresholdConfigEntry"
    cnpdThresholdConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cnpdThresholdConfigEntry.EntityData.ParentYangName = "cnpdThresholdConfigTable"
    cnpdThresholdConfigEntry.EntityData.SegmentPath = "cnpdThresholdConfigEntry" + types.AddKeyToken(cnpdThresholdConfigEntry.CnpdThresholdConfigIndex, "cnpdThresholdConfigIndex")
    cnpdThresholdConfigEntry.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/cnpdThresholdConfigTable/" + cnpdThresholdConfigEntry.EntityData.SegmentPath
    cnpdThresholdConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdThresholdConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdThresholdConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdThresholdConfigEntry.EntityData.Children = types.NewOrderedMap()
    cnpdThresholdConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigIndex", types.YLeaf{"CnpdThresholdConfigIndex", cnpdThresholdConfigEntry.CnpdThresholdConfigIndex})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigIfIndex", types.YLeaf{"CnpdThresholdConfigIfIndex", cnpdThresholdConfigEntry.CnpdThresholdConfigIfIndex})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigInterval", types.YLeaf{"CnpdThresholdConfigInterval", cnpdThresholdConfigEntry.CnpdThresholdConfigInterval})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigSampleType", types.YLeaf{"CnpdThresholdConfigSampleType", cnpdThresholdConfigEntry.CnpdThresholdConfigSampleType})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigProtocol", types.YLeaf{"CnpdThresholdConfigProtocol", cnpdThresholdConfigEntry.CnpdThresholdConfigProtocol})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigProtocolAny", types.YLeaf{"CnpdThresholdConfigProtocolAny", cnpdThresholdConfigEntry.CnpdThresholdConfigProtocolAny})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigStatsSelect", types.YLeaf{"CnpdThresholdConfigStatsSelect", cnpdThresholdConfigEntry.CnpdThresholdConfigStatsSelect})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigStartup", types.YLeaf{"CnpdThresholdConfigStartup", cnpdThresholdConfigEntry.CnpdThresholdConfigStartup})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigRising", types.YLeaf{"CnpdThresholdConfigRising", cnpdThresholdConfigEntry.CnpdThresholdConfigRising})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigFalling", types.YLeaf{"CnpdThresholdConfigFalling", cnpdThresholdConfigEntry.CnpdThresholdConfigFalling})
    cnpdThresholdConfigEntry.EntityData.Leafs.Append("cnpdThresholdConfigStatus", types.YLeaf{"CnpdThresholdConfigStatus", cnpdThresholdConfigEntry.CnpdThresholdConfigStatus})

    cnpdThresholdConfigEntry.EntityData.YListKeys = []string {"CnpdThresholdConfigIndex"}

    return &(cnpdThresholdConfigEntry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigSampleType represents cnpdThresholdConfigFalling value.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigSampleType string

const (
    CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigSampleType_absoluteValue CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigSampleType = "absoluteValue"

    CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigSampleType_deltaValue CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigSampleType = "deltaValue"
)

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigStartup represents will be generated.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigStartup string

const (
    CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigStartup_rising CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigStartup = "rising"

    CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigStartup_falling CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigStartup = "falling"

    CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigStartup_risingOrFalling CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdConfigTable_CnpdThresholdConfigEntry_CnpdThresholdConfigStartup = "risingOrFalling"
)

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable
// The Threshold History table. Notifications
// are unreliable so this table provides a
// history of the last 5000 threshold breached
// events. A notification can be traced back to
// its cnpdThresholdHistoryEntry.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry is created each time a threshold  is breached.   Thus there is
    // not necessarily a one to one  relationship to cnpdThresholdConfigTable  as
    // not every Threshold configured will  be breached. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry.
    CnpdThresholdHistoryEntry []*CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry
}

func (cnpdThresholdHistoryTable *CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable) GetEntityData() *types.CommonEntityData {
    cnpdThresholdHistoryTable.EntityData.YFilter = cnpdThresholdHistoryTable.YFilter
    cnpdThresholdHistoryTable.EntityData.YangName = "cnpdThresholdHistoryTable"
    cnpdThresholdHistoryTable.EntityData.BundleName = "cisco_ios_xe"
    cnpdThresholdHistoryTable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdThresholdHistoryTable.EntityData.SegmentPath = "cnpdThresholdHistoryTable"
    cnpdThresholdHistoryTable.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/" + cnpdThresholdHistoryTable.EntityData.SegmentPath
    cnpdThresholdHistoryTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdThresholdHistoryTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdThresholdHistoryTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdThresholdHistoryTable.EntityData.Children = types.NewOrderedMap()
    cnpdThresholdHistoryTable.EntityData.Children.Append("cnpdThresholdHistoryEntry", types.YChild{"CnpdThresholdHistoryEntry", nil})
    for i := range cnpdThresholdHistoryTable.CnpdThresholdHistoryEntry {
        cnpdThresholdHistoryTable.EntityData.Children.Append(types.GetSegmentPath(cnpdThresholdHistoryTable.CnpdThresholdHistoryEntry[i]), types.YChild{"CnpdThresholdHistoryEntry", cnpdThresholdHistoryTable.CnpdThresholdHistoryEntry[i]})
    }
    cnpdThresholdHistoryTable.EntityData.Leafs = types.NewOrderedMap()

    cnpdThresholdHistoryTable.EntityData.YListKeys = []string {}

    return &(cnpdThresholdHistoryTable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry
// This entry is created each time a threshold 
// is breached. 
// 
// Thus there is not necessarily a one to one 
// relationship to cnpdThresholdConfigTable 
// as not every Threshold configured will 
// be breached.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A monotonically increasing integer which uniquely
    // identifies this  cnpdThresholdHistoryEntry in the  cnpdThresholdHistory
    // table. The type is interface{} with range: 1..1000.
    CnpdThresholdHistoryIndex interface{}

    // The cnpdThresholdConfigTable entry  which generated this entry. Using this 
    // object the management station can backtrack  to the appropriate
    // cnpdThresholdConfigEntry. The type is interface{} with range: 1..1000.
    CnpdThresholdHistoryConfigIndex interface{}

    // The actual value of the statistic when  the sampling was made. The type is
    // interface{} with range: 1..4294967295.
    CnpdThresholdHistoryValue interface{}

    // Describes whether this is an event caused by a rising or falling threshold
    // breach. The type is CnpdThresholdHistoryType.
    CnpdThresholdHistoryType interface{}

    // The value of sysUpTime of the running  configuration when the event
    // occurred. The type is interface{} with range: 0..4294967295.
    CnpdThresholdHistoryTime interface{}

    // The application or protocol which the management station configured a
    // threshold on.  This object is an index into the  SupportedProtocolsTable
    // where details of the protocol can be found. The type is interface{} with
    // range: 1..1024.
    CnpdThresholdHistoryProtocol interface{}

    // This is the statistic used to base the threshold on. The type is
    // CiscoPdDataType.
    CnpdThresholdHistoryStatsSelect interface{}
}

func (cnpdThresholdHistoryEntry *CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry) GetEntityData() *types.CommonEntityData {
    cnpdThresholdHistoryEntry.EntityData.YFilter = cnpdThresholdHistoryEntry.YFilter
    cnpdThresholdHistoryEntry.EntityData.YangName = "cnpdThresholdHistoryEntry"
    cnpdThresholdHistoryEntry.EntityData.BundleName = "cisco_ios_xe"
    cnpdThresholdHistoryEntry.EntityData.ParentYangName = "cnpdThresholdHistoryTable"
    cnpdThresholdHistoryEntry.EntityData.SegmentPath = "cnpdThresholdHistoryEntry" + types.AddKeyToken(cnpdThresholdHistoryEntry.CnpdThresholdHistoryIndex, "cnpdThresholdHistoryIndex")
    cnpdThresholdHistoryEntry.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/cnpdThresholdHistoryTable/" + cnpdThresholdHistoryEntry.EntityData.SegmentPath
    cnpdThresholdHistoryEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdThresholdHistoryEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdThresholdHistoryEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdThresholdHistoryEntry.EntityData.Children = types.NewOrderedMap()
    cnpdThresholdHistoryEntry.EntityData.Leafs = types.NewOrderedMap()
    cnpdThresholdHistoryEntry.EntityData.Leafs.Append("cnpdThresholdHistoryIndex", types.YLeaf{"CnpdThresholdHistoryIndex", cnpdThresholdHistoryEntry.CnpdThresholdHistoryIndex})
    cnpdThresholdHistoryEntry.EntityData.Leafs.Append("cnpdThresholdHistoryConfigIndex", types.YLeaf{"CnpdThresholdHistoryConfigIndex", cnpdThresholdHistoryEntry.CnpdThresholdHistoryConfigIndex})
    cnpdThresholdHistoryEntry.EntityData.Leafs.Append("cnpdThresholdHistoryValue", types.YLeaf{"CnpdThresholdHistoryValue", cnpdThresholdHistoryEntry.CnpdThresholdHistoryValue})
    cnpdThresholdHistoryEntry.EntityData.Leafs.Append("cnpdThresholdHistoryType", types.YLeaf{"CnpdThresholdHistoryType", cnpdThresholdHistoryEntry.CnpdThresholdHistoryType})
    cnpdThresholdHistoryEntry.EntityData.Leafs.Append("cnpdThresholdHistoryTime", types.YLeaf{"CnpdThresholdHistoryTime", cnpdThresholdHistoryEntry.CnpdThresholdHistoryTime})
    cnpdThresholdHistoryEntry.EntityData.Leafs.Append("cnpdThresholdHistoryProtocol", types.YLeaf{"CnpdThresholdHistoryProtocol", cnpdThresholdHistoryEntry.CnpdThresholdHistoryProtocol})
    cnpdThresholdHistoryEntry.EntityData.Leafs.Append("cnpdThresholdHistoryStatsSelect", types.YLeaf{"CnpdThresholdHistoryStatsSelect", cnpdThresholdHistoryEntry.CnpdThresholdHistoryStatsSelect})

    cnpdThresholdHistoryEntry.EntityData.YListKeys = []string {"CnpdThresholdHistoryIndex"}

    return &(cnpdThresholdHistoryEntry.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry_CnpdThresholdHistoryType represents or falling threshold breach.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry_CnpdThresholdHistoryType string

const (
    CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry_CnpdThresholdHistoryType_risingBreach CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry_CnpdThresholdHistoryType = "risingBreach"

    CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry_CnpdThresholdHistoryType_fallingBreach CISCONBARPROTOCOLDISCOVERYMIB_CnpdThresholdHistoryTable_CnpdThresholdHistoryEntry_CnpdThresholdHistoryType = "fallingBreach"
)

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable
// The Supported Protocols table lists all the 
// protocols and applications which NBAR is currently
// capable of recognizing.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A entry in the Supported Protocols table reflecting key information about a
    // protocol. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable_CnpdSupportedProtocolsEntry.
    CnpdSupportedProtocolsEntry []*CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable_CnpdSupportedProtocolsEntry
}

func (cnpdSupportedProtocolsTable *CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable) GetEntityData() *types.CommonEntityData {
    cnpdSupportedProtocolsTable.EntityData.YFilter = cnpdSupportedProtocolsTable.YFilter
    cnpdSupportedProtocolsTable.EntityData.YangName = "cnpdSupportedProtocolsTable"
    cnpdSupportedProtocolsTable.EntityData.BundleName = "cisco_ios_xe"
    cnpdSupportedProtocolsTable.EntityData.ParentYangName = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
    cnpdSupportedProtocolsTable.EntityData.SegmentPath = "cnpdSupportedProtocolsTable"
    cnpdSupportedProtocolsTable.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/" + cnpdSupportedProtocolsTable.EntityData.SegmentPath
    cnpdSupportedProtocolsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdSupportedProtocolsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdSupportedProtocolsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdSupportedProtocolsTable.EntityData.Children = types.NewOrderedMap()
    cnpdSupportedProtocolsTable.EntityData.Children.Append("cnpdSupportedProtocolsEntry", types.YChild{"CnpdSupportedProtocolsEntry", nil})
    for i := range cnpdSupportedProtocolsTable.CnpdSupportedProtocolsEntry {
        cnpdSupportedProtocolsTable.EntityData.Children.Append(types.GetSegmentPath(cnpdSupportedProtocolsTable.CnpdSupportedProtocolsEntry[i]), types.YChild{"CnpdSupportedProtocolsEntry", cnpdSupportedProtocolsTable.CnpdSupportedProtocolsEntry[i]})
    }
    cnpdSupportedProtocolsTable.EntityData.Leafs = types.NewOrderedMap()

    cnpdSupportedProtocolsTable.EntityData.YListKeys = []string {}

    return &(cnpdSupportedProtocolsTable.EntityData)
}

// CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable_CnpdSupportedProtocolsEntry
// A entry in the Supported Protocols table reflecting
// key information about a protocol.
type CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable_CnpdSupportedProtocolsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A unique identifier of a row in this table.  Thus
    // it also represents a unique identifier for a protocol or application which
    // NBAR currently recognizes. The type is interface{} with range: 1..1024.
    CnpdSupportedProtocolsIndex interface{}

    // This object reflects the valid string of a protocol or application which
    // NBAR currently recognizes. The type is string with length: 1..255.
    CnpdSupportedProtocolsName interface{}
}

func (cnpdSupportedProtocolsEntry *CISCONBARPROTOCOLDISCOVERYMIB_CnpdSupportedProtocolsTable_CnpdSupportedProtocolsEntry) GetEntityData() *types.CommonEntityData {
    cnpdSupportedProtocolsEntry.EntityData.YFilter = cnpdSupportedProtocolsEntry.YFilter
    cnpdSupportedProtocolsEntry.EntityData.YangName = "cnpdSupportedProtocolsEntry"
    cnpdSupportedProtocolsEntry.EntityData.BundleName = "cisco_ios_xe"
    cnpdSupportedProtocolsEntry.EntityData.ParentYangName = "cnpdSupportedProtocolsTable"
    cnpdSupportedProtocolsEntry.EntityData.SegmentPath = "cnpdSupportedProtocolsEntry" + types.AddKeyToken(cnpdSupportedProtocolsEntry.CnpdSupportedProtocolsIndex, "cnpdSupportedProtocolsIndex")
    cnpdSupportedProtocolsEntry.EntityData.AbsolutePath = "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB/cnpdSupportedProtocolsTable/" + cnpdSupportedProtocolsEntry.EntityData.SegmentPath
    cnpdSupportedProtocolsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cnpdSupportedProtocolsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cnpdSupportedProtocolsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cnpdSupportedProtocolsEntry.EntityData.Children = types.NewOrderedMap()
    cnpdSupportedProtocolsEntry.EntityData.Leafs = types.NewOrderedMap()
    cnpdSupportedProtocolsEntry.EntityData.Leafs.Append("cnpdSupportedProtocolsIndex", types.YLeaf{"CnpdSupportedProtocolsIndex", cnpdSupportedProtocolsEntry.CnpdSupportedProtocolsIndex})
    cnpdSupportedProtocolsEntry.EntityData.Leafs.Append("cnpdSupportedProtocolsName", types.YLeaf{"CnpdSupportedProtocolsName", cnpdSupportedProtocolsEntry.CnpdSupportedProtocolsName})

    cnpdSupportedProtocolsEntry.EntityData.YListKeys = []string {"CnpdSupportedProtocolsIndex"}

    return &(cnpdSupportedProtocolsEntry.EntityData)
}

