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
    parent types.Entity
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

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetFilter() yfilter.YFilter { return cISCONBARPROTOCOLDISCOVERYMIB.YFilter }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) SetFilter(yf yfilter.YFilter) { cISCONBARPROTOCOLDISCOVERYMIB.YFilter = yf }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetGoName(yname string) string {
    if yname == "cnpdNotificationsConfig" { return "Cnpdnotificationsconfig" }
    if yname == "cnpdStatusTable" { return "Cnpdstatustable" }
    if yname == "cnpdAllStatsTable" { return "Cnpdallstatstable" }
    if yname == "cnpdTopNConfigTable" { return "Cnpdtopnconfigtable" }
    if yname == "cnpdTopNStatsTable" { return "Cnpdtopnstatstable" }
    if yname == "cnpdThresholdConfigTable" { return "Cnpdthresholdconfigtable" }
    if yname == "cnpdThresholdHistoryTable" { return "Cnpdthresholdhistorytable" }
    if yname == "cnpdSupportedProtocolsTable" { return "Cnpdsupportedprotocolstable" }
    return ""
}

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetSegmentPath() string {
    return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB:CISCO-NBAR-PROTOCOL-DISCOVERY-MIB"
}

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnpdNotificationsConfig" {
        return &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdnotificationsconfig
    }
    if childYangName == "cnpdStatusTable" {
        return &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdstatustable
    }
    if childYangName == "cnpdAllStatsTable" {
        return &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdallstatstable
    }
    if childYangName == "cnpdTopNConfigTable" {
        return &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdtopnconfigtable
    }
    if childYangName == "cnpdTopNStatsTable" {
        return &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdtopnstatstable
    }
    if childYangName == "cnpdThresholdConfigTable" {
        return &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdthresholdconfigtable
    }
    if childYangName == "cnpdThresholdHistoryTable" {
        return &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdthresholdhistorytable
    }
    if childYangName == "cnpdSupportedProtocolsTable" {
        return &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdsupportedprotocolstable
    }
    return nil
}

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cnpdNotificationsConfig"] = &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdnotificationsconfig
    children["cnpdStatusTable"] = &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdstatustable
    children["cnpdAllStatsTable"] = &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdallstatstable
    children["cnpdTopNConfigTable"] = &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdtopnconfigtable
    children["cnpdTopNStatsTable"] = &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdtopnstatstable
    children["cnpdThresholdConfigTable"] = &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdthresholdconfigtable
    children["cnpdThresholdHistoryTable"] = &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdthresholdhistorytable
    children["cnpdSupportedProtocolsTable"] = &cISCONBARPROTOCOLDISCOVERYMIB.Cnpdsupportedprotocolstable
    return children
}

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) SetParent(parent types.Entity) { cISCONBARPROTOCOLDISCOVERYMIB.parent = parent }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetParent() types.Entity { return cISCONBARPROTOCOLDISCOVERYMIB.parent }

func (cISCONBARPROTOCOLDISCOVERYMIB *CISCONBARPROTOCOLDISCOVERYMIB) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object is used to enable or disable  Notifications on a global basis. 
    // If set to 'true' - Notifications are enabled. If set to 'false' -
    // Notifications are disabled. The type is bool.
    Cnpdnotificationsenable interface{}
}

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetFilter() yfilter.YFilter { return cnpdnotificationsconfig.YFilter }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) SetFilter(yf yfilter.YFilter) { cnpdnotificationsconfig.YFilter = yf }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetGoName(yname string) string {
    if yname == "cnpdNotificationsEnable" { return "Cnpdnotificationsenable" }
    return ""
}

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetSegmentPath() string {
    return "cnpdNotificationsConfig"
}

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnpdNotificationsEnable"] = cnpdnotificationsconfig.Cnpdnotificationsenable
    return leafs
}

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetYangName() string { return "cnpdNotificationsConfig" }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) SetParent(parent types.Entity) { cnpdnotificationsconfig.parent = parent }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetParent() types.Entity { return cnpdnotificationsconfig.parent }

func (cnpdnotificationsconfig *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdnotificationsconfig) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable
// The cnpdStatusTable is used to enable and
// disable Protocol Discovery on an interface.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the cnpdStatusTable contains objects for enabling or disabling
    // Protocol Discovery on a per interface basis. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry.
    Cnpdstatusentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry
}

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetFilter() yfilter.YFilter { return cnpdstatustable.YFilter }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) SetFilter(yf yfilter.YFilter) { cnpdstatustable.YFilter = yf }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetGoName(yname string) string {
    if yname == "cnpdStatusEntry" { return "Cnpdstatusentry" }
    return ""
}

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetSegmentPath() string {
    return "cnpdStatusTable"
}

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnpdStatusEntry" {
        for _, c := range cnpdstatustable.Cnpdstatusentry {
            if cnpdstatustable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry{}
        cnpdstatustable.Cnpdstatusentry = append(cnpdstatustable.Cnpdstatusentry, child)
        return &cnpdstatustable.Cnpdstatusentry[len(cnpdstatustable.Cnpdstatusentry)-1]
    }
    return nil
}

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnpdstatustable.Cnpdstatusentry {
        children[cnpdstatustable.Cnpdstatusentry[i].GetSegmentPath()] = &cnpdstatustable.Cnpdstatusentry[i]
    }
    return children
}

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetYangName() string { return "cnpdStatusTable" }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) SetParent(parent types.Entity) { cnpdstatustable.parent = parent }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetParent() types.Entity { return cnpdstatustable.parent }

func (cnpdstatustable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry
// An entry in the cnpdStatusTable contains objects
// for enabling or disabling Protocol Discovery on a
// per interface basis.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry struct {
    parent types.Entity
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

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetFilter() yfilter.YFilter { return cnpdstatusentry.YFilter }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) SetFilter(yf yfilter.YFilter) { cnpdstatusentry.YFilter = yf }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cnpdStatusPdEnable" { return "Cnpdstatuspdenable" }
    if yname == "cnpdStatusLastUpdateTime" { return "Cnpdstatuslastupdatetime" }
    return ""
}

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetSegmentPath() string {
    return "cnpdStatusEntry" + "[ifIndex='" + fmt.Sprintf("%v", cnpdstatusentry.Ifindex) + "']"
}

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cnpdstatusentry.Ifindex
    leafs["cnpdStatusPdEnable"] = cnpdstatusentry.Cnpdstatuspdenable
    leafs["cnpdStatusLastUpdateTime"] = cnpdstatusentry.Cnpdstatuslastupdatetime
    return leafs
}

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetYangName() string { return "cnpdStatusEntry" }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) SetParent(parent types.Entity) { cnpdstatusentry.parent = parent }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetParent() types.Entity { return cnpdstatusentry.parent }

func (cnpdstatusentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdstatustable_Cnpdstatusentry) GetParentYangName() string { return "cnpdStatusTable" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable
// The cnpdAllStatsTable contains all the statistics
// available for all the protocols/applications currently
// recognized by NBAR Protocol Discovery for a particular 
// interface.
// 
// In the event of an overflow, the 32 bit counters are not 
// valid. There is no overflow support.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the cnpdAllStatsTable table. This entry  contains the
    // statistics collected on all the protocols  which NBAR classifies for a
    // particular interface. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry.
    Cnpdallstatsentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry
}

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetFilter() yfilter.YFilter { return cnpdallstatstable.YFilter }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) SetFilter(yf yfilter.YFilter) { cnpdallstatstable.YFilter = yf }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetGoName(yname string) string {
    if yname == "cnpdAllStatsEntry" { return "Cnpdallstatsentry" }
    return ""
}

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetSegmentPath() string {
    return "cnpdAllStatsTable"
}

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnpdAllStatsEntry" {
        for _, c := range cnpdallstatstable.Cnpdallstatsentry {
            if cnpdallstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry{}
        cnpdallstatstable.Cnpdallstatsentry = append(cnpdallstatstable.Cnpdallstatsentry, child)
        return &cnpdallstatstable.Cnpdallstatsentry[len(cnpdallstatstable.Cnpdallstatsentry)-1]
    }
    return nil
}

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnpdallstatstable.Cnpdallstatsentry {
        children[cnpdallstatstable.Cnpdallstatsentry[i].GetSegmentPath()] = &cnpdallstatstable.Cnpdallstatsentry[i]
    }
    return children
}

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetYangName() string { return "cnpdAllStatsTable" }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) SetParent(parent types.Entity) { cnpdallstatstable.parent = parent }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetParent() types.Entity { return cnpdallstatstable.parent }

func (cnpdallstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry
// An entry in the cnpdAllStatsTable table. This entry 
// contains the statistics collected on all the protocols 
// which NBAR classifies for a particular interface.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry struct {
    parent types.Entity
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

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetFilter() yfilter.YFilter { return cnpdallstatsentry.YFilter }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) SetFilter(yf yfilter.YFilter) { cnpdallstatsentry.YFilter = yf }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cnpdAllStatsProtocolsIndex" { return "Cnpdallstatsprotocolsindex" }
    if yname == "cnpdAllStatsProtocolName" { return "Cnpdallstatsprotocolname" }
    if yname == "cnpdAllStatsInPkts" { return "Cnpdallstatsinpkts" }
    if yname == "cnpdAllStatsOutPkts" { return "Cnpdallstatsoutpkts" }
    if yname == "cnpdAllStatsInBytes" { return "Cnpdallstatsinbytes" }
    if yname == "cnpdAllStatsOutBytes" { return "Cnpdallstatsoutbytes" }
    if yname == "cnpdAllStatsHCInPkts" { return "Cnpdallstatshcinpkts" }
    if yname == "cnpdAllStatsHCOutPkts" { return "Cnpdallstatshcoutpkts" }
    if yname == "cnpdAllStatsHCInBytes" { return "Cnpdallstatshcinbytes" }
    if yname == "cnpdAllStatsHCOutBytes" { return "Cnpdallstatshcoutbytes" }
    if yname == "cnpdAllStatsInBitRate" { return "Cnpdallstatsinbitrate" }
    if yname == "cnpdAllStatsOutBitRate" { return "Cnpdallstatsoutbitrate" }
    return ""
}

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetSegmentPath() string {
    return "cnpdAllStatsEntry" + "[ifIndex='" + fmt.Sprintf("%v", cnpdallstatsentry.Ifindex) + "']" + "[cnpdAllStatsProtocolsIndex='" + fmt.Sprintf("%v", cnpdallstatsentry.Cnpdallstatsprotocolsindex) + "']"
}

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cnpdallstatsentry.Ifindex
    leafs["cnpdAllStatsProtocolsIndex"] = cnpdallstatsentry.Cnpdallstatsprotocolsindex
    leafs["cnpdAllStatsProtocolName"] = cnpdallstatsentry.Cnpdallstatsprotocolname
    leafs["cnpdAllStatsInPkts"] = cnpdallstatsentry.Cnpdallstatsinpkts
    leafs["cnpdAllStatsOutPkts"] = cnpdallstatsentry.Cnpdallstatsoutpkts
    leafs["cnpdAllStatsInBytes"] = cnpdallstatsentry.Cnpdallstatsinbytes
    leafs["cnpdAllStatsOutBytes"] = cnpdallstatsentry.Cnpdallstatsoutbytes
    leafs["cnpdAllStatsHCInPkts"] = cnpdallstatsentry.Cnpdallstatshcinpkts
    leafs["cnpdAllStatsHCOutPkts"] = cnpdallstatsentry.Cnpdallstatshcoutpkts
    leafs["cnpdAllStatsHCInBytes"] = cnpdallstatsentry.Cnpdallstatshcinbytes
    leafs["cnpdAllStatsHCOutBytes"] = cnpdallstatsentry.Cnpdallstatshcoutbytes
    leafs["cnpdAllStatsInBitRate"] = cnpdallstatsentry.Cnpdallstatsinbitrate
    leafs["cnpdAllStatsOutBitRate"] = cnpdallstatsentry.Cnpdallstatsoutbitrate
    return leafs
}

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetYangName() string { return "cnpdAllStatsEntry" }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) SetParent(parent types.Entity) { cnpdallstatsentry.parent = parent }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetParent() types.Entity { return cnpdallstatsentry.parent }

func (cnpdallstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdallstatstable_Cnpdallstatsentry) GetParentYangName() string { return "cnpdAllStatsTable" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable
// The cnpdTopNConfigTable is used to configure
// cnpdTopNStatsTable's.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry provides the objects to configure and thus initiate the
    // generation of a cnpdTopNStatsTable.. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry.
    Cnpdtopnconfigentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry
}

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetFilter() yfilter.YFilter { return cnpdtopnconfigtable.YFilter }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) SetFilter(yf yfilter.YFilter) { cnpdtopnconfigtable.YFilter = yf }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetGoName(yname string) string {
    if yname == "cnpdTopNConfigEntry" { return "Cnpdtopnconfigentry" }
    return ""
}

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetSegmentPath() string {
    return "cnpdTopNConfigTable"
}

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnpdTopNConfigEntry" {
        for _, c := range cnpdtopnconfigtable.Cnpdtopnconfigentry {
            if cnpdtopnconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry{}
        cnpdtopnconfigtable.Cnpdtopnconfigentry = append(cnpdtopnconfigtable.Cnpdtopnconfigentry, child)
        return &cnpdtopnconfigtable.Cnpdtopnconfigentry[len(cnpdtopnconfigtable.Cnpdtopnconfigentry)-1]
    }
    return nil
}

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnpdtopnconfigtable.Cnpdtopnconfigentry {
        children[cnpdtopnconfigtable.Cnpdtopnconfigentry[i].GetSegmentPath()] = &cnpdtopnconfigtable.Cnpdtopnconfigentry[i]
    }
    return children
}

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetYangName() string { return "cnpdTopNConfigTable" }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) SetParent(parent types.Entity) { cnpdtopnconfigtable.parent = parent }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetParent() types.Entity { return cnpdtopnconfigtable.parent }

func (cnpdtopnconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry
// This entry provides the objects to configure and thus
// initiate the generation of a cnpdTopNStatsTable..
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry struct {
    parent types.Entity
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

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetFilter() yfilter.YFilter { return cnpdtopnconfigentry.YFilter }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) SetFilter(yf yfilter.YFilter) { cnpdtopnconfigentry.YFilter = yf }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetGoName(yname string) string {
    if yname == "cnpdTopNConfigIndex" { return "Cnpdtopnconfigindex" }
    if yname == "cnpdTopNConfigIfIndex" { return "Cnpdtopnconfigifindex" }
    if yname == "cnpdTopNConfigStatsSelect" { return "Cnpdtopnconfigstatsselect" }
    if yname == "cnpdTopNConfigSampleTime" { return "Cnpdtopnconfigsampletime" }
    if yname == "cnpdTopNConfigRequestedSize" { return "Cnpdtopnconfigrequestedsize" }
    if yname == "cnpdTopNConfigGrantedSize" { return "Cnpdtopnconfiggrantedsize" }
    if yname == "cnpdTopNConfigTime" { return "Cnpdtopnconfigtime" }
    if yname == "cnpdTopNConfigStatus" { return "Cnpdtopnconfigstatus" }
    return ""
}

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetSegmentPath() string {
    return "cnpdTopNConfigEntry" + "[cnpdTopNConfigIndex='" + fmt.Sprintf("%v", cnpdtopnconfigentry.Cnpdtopnconfigindex) + "']"
}

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnpdTopNConfigIndex"] = cnpdtopnconfigentry.Cnpdtopnconfigindex
    leafs["cnpdTopNConfigIfIndex"] = cnpdtopnconfigentry.Cnpdtopnconfigifindex
    leafs["cnpdTopNConfigStatsSelect"] = cnpdtopnconfigentry.Cnpdtopnconfigstatsselect
    leafs["cnpdTopNConfigSampleTime"] = cnpdtopnconfigentry.Cnpdtopnconfigsampletime
    leafs["cnpdTopNConfigRequestedSize"] = cnpdtopnconfigentry.Cnpdtopnconfigrequestedsize
    leafs["cnpdTopNConfigGrantedSize"] = cnpdtopnconfigentry.Cnpdtopnconfiggrantedsize
    leafs["cnpdTopNConfigTime"] = cnpdtopnconfigentry.Cnpdtopnconfigtime
    leafs["cnpdTopNConfigStatus"] = cnpdtopnconfigentry.Cnpdtopnconfigstatus
    return leafs
}

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetYangName() string { return "cnpdTopNConfigEntry" }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) SetParent(parent types.Entity) { cnpdtopnconfigentry.parent = parent }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetParent() types.Entity { return cnpdtopnconfigentry.parent }

func (cnpdtopnconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnconfigtable_Cnpdtopnconfigentry) GetParentYangName() string { return "cnpdTopNConfigTable" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable
// A cnpdTopNStatsTable describes an ordered
// list of protocols.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable struct {
    parent types.Entity
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

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetFilter() yfilter.YFilter { return cnpdtopnstatstable.YFilter }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) SetFilter(yf yfilter.YFilter) { cnpdtopnstatstable.YFilter = yf }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetGoName(yname string) string {
    if yname == "cnpdTopNStatsEntry" { return "Cnpdtopnstatsentry" }
    return ""
}

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetSegmentPath() string {
    return "cnpdTopNStatsTable"
}

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnpdTopNStatsEntry" {
        for _, c := range cnpdtopnstatstable.Cnpdtopnstatsentry {
            if cnpdtopnstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry{}
        cnpdtopnstatstable.Cnpdtopnstatsentry = append(cnpdtopnstatstable.Cnpdtopnstatsentry, child)
        return &cnpdtopnstatstable.Cnpdtopnstatsentry[len(cnpdtopnstatstable.Cnpdtopnstatsentry)-1]
    }
    return nil
}

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnpdtopnstatstable.Cnpdtopnstatsentry {
        children[cnpdtopnstatstable.Cnpdtopnstatsentry[i].GetSegmentPath()] = &cnpdtopnstatstable.Cnpdtopnstatsentry[i]
    }
    return children
}

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetYangName() string { return "cnpdTopNStatsTable" }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) SetParent(parent types.Entity) { cnpdtopnstatstable.parent = parent }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetParent() types.Entity { return cnpdtopnstatstable.parent }

func (cnpdtopnstatstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

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
    parent types.Entity
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

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetFilter() yfilter.YFilter { return cnpdtopnstatsentry.YFilter }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) SetFilter(yf yfilter.YFilter) { cnpdtopnstatsentry.YFilter = yf }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetGoName(yname string) string {
    if yname == "cnpdTopNConfigIndex" { return "Cnpdtopnconfigindex" }
    if yname == "cnpdTopNStatsIndex" { return "Cnpdtopnstatsindex" }
    if yname == "cnpdTopNStatsProtocolName" { return "Cnpdtopnstatsprotocolname" }
    if yname == "cnpdTopNStatsRate" { return "Cnpdtopnstatsrate" }
    if yname == "cnpdTopNStatsHCRate" { return "Cnpdtopnstatshcrate" }
    return ""
}

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetSegmentPath() string {
    return "cnpdTopNStatsEntry" + "[cnpdTopNConfigIndex='" + fmt.Sprintf("%v", cnpdtopnstatsentry.Cnpdtopnconfigindex) + "']" + "[cnpdTopNStatsIndex='" + fmt.Sprintf("%v", cnpdtopnstatsentry.Cnpdtopnstatsindex) + "']"
}

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnpdTopNConfigIndex"] = cnpdtopnstatsentry.Cnpdtopnconfigindex
    leafs["cnpdTopNStatsIndex"] = cnpdtopnstatsentry.Cnpdtopnstatsindex
    leafs["cnpdTopNStatsProtocolName"] = cnpdtopnstatsentry.Cnpdtopnstatsprotocolname
    leafs["cnpdTopNStatsRate"] = cnpdtopnstatsentry.Cnpdtopnstatsrate
    leafs["cnpdTopNStatsHCRate"] = cnpdtopnstatsentry.Cnpdtopnstatshcrate
    return leafs
}

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetYangName() string { return "cnpdTopNStatsEntry" }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) SetParent(parent types.Entity) { cnpdtopnstatsentry.parent = parent }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetParent() types.Entity { return cnpdtopnstatsentry.parent }

func (cnpdtopnstatsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdtopnstatstable_Cnpdtopnstatsentry) GetParentYangName() string { return "cnpdTopNStatsTable" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable
// The cnpdThresholdConfigTable allows the management
// station to create thresholds for the purpose of
// sending notifications if breached, and creating a
// history of breached thresholds.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable struct {
    parent types.Entity
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

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetFilter() yfilter.YFilter { return cnpdthresholdconfigtable.YFilter }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) SetFilter(yf yfilter.YFilter) { cnpdthresholdconfigtable.YFilter = yf }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetGoName(yname string) string {
    if yname == "cnpdThresholdConfigEntry" { return "Cnpdthresholdconfigentry" }
    return ""
}

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetSegmentPath() string {
    return "cnpdThresholdConfigTable"
}

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnpdThresholdConfigEntry" {
        for _, c := range cnpdthresholdconfigtable.Cnpdthresholdconfigentry {
            if cnpdthresholdconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry{}
        cnpdthresholdconfigtable.Cnpdthresholdconfigentry = append(cnpdthresholdconfigtable.Cnpdthresholdconfigentry, child)
        return &cnpdthresholdconfigtable.Cnpdthresholdconfigentry[len(cnpdthresholdconfigtable.Cnpdthresholdconfigentry)-1]
    }
    return nil
}

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnpdthresholdconfigtable.Cnpdthresholdconfigentry {
        children[cnpdthresholdconfigtable.Cnpdthresholdconfigentry[i].GetSegmentPath()] = &cnpdthresholdconfigtable.Cnpdthresholdconfigentry[i]
    }
    return children
}

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetYangName() string { return "cnpdThresholdConfigTable" }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) SetParent(parent types.Entity) { cnpdthresholdconfigtable.parent = parent }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetParent() types.Entity { return cnpdthresholdconfigtable.parent }

func (cnpdthresholdconfigtable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

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
    parent types.Entity
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

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetFilter() yfilter.YFilter { return cnpdthresholdconfigentry.YFilter }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) SetFilter(yf yfilter.YFilter) { cnpdthresholdconfigentry.YFilter = yf }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetGoName(yname string) string {
    if yname == "cnpdThresholdConfigIndex" { return "Cnpdthresholdconfigindex" }
    if yname == "cnpdThresholdConfigIfIndex" { return "Cnpdthresholdconfigifindex" }
    if yname == "cnpdThresholdConfigInterval" { return "Cnpdthresholdconfiginterval" }
    if yname == "cnpdThresholdConfigSampleType" { return "Cnpdthresholdconfigsampletype" }
    if yname == "cnpdThresholdConfigProtocol" { return "Cnpdthresholdconfigprotocol" }
    if yname == "cnpdThresholdConfigProtocolAny" { return "Cnpdthresholdconfigprotocolany" }
    if yname == "cnpdThresholdConfigStatsSelect" { return "Cnpdthresholdconfigstatsselect" }
    if yname == "cnpdThresholdConfigStartup" { return "Cnpdthresholdconfigstartup" }
    if yname == "cnpdThresholdConfigRising" { return "Cnpdthresholdconfigrising" }
    if yname == "cnpdThresholdConfigFalling" { return "Cnpdthresholdconfigfalling" }
    if yname == "cnpdThresholdConfigStatus" { return "Cnpdthresholdconfigstatus" }
    return ""
}

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetSegmentPath() string {
    return "cnpdThresholdConfigEntry" + "[cnpdThresholdConfigIndex='" + fmt.Sprintf("%v", cnpdthresholdconfigentry.Cnpdthresholdconfigindex) + "']"
}

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnpdThresholdConfigIndex"] = cnpdthresholdconfigentry.Cnpdthresholdconfigindex
    leafs["cnpdThresholdConfigIfIndex"] = cnpdthresholdconfigentry.Cnpdthresholdconfigifindex
    leafs["cnpdThresholdConfigInterval"] = cnpdthresholdconfigentry.Cnpdthresholdconfiginterval
    leafs["cnpdThresholdConfigSampleType"] = cnpdthresholdconfigentry.Cnpdthresholdconfigsampletype
    leafs["cnpdThresholdConfigProtocol"] = cnpdthresholdconfigentry.Cnpdthresholdconfigprotocol
    leafs["cnpdThresholdConfigProtocolAny"] = cnpdthresholdconfigentry.Cnpdthresholdconfigprotocolany
    leafs["cnpdThresholdConfigStatsSelect"] = cnpdthresholdconfigentry.Cnpdthresholdconfigstatsselect
    leafs["cnpdThresholdConfigStartup"] = cnpdthresholdconfigentry.Cnpdthresholdconfigstartup
    leafs["cnpdThresholdConfigRising"] = cnpdthresholdconfigentry.Cnpdthresholdconfigrising
    leafs["cnpdThresholdConfigFalling"] = cnpdthresholdconfigentry.Cnpdthresholdconfigfalling
    leafs["cnpdThresholdConfigStatus"] = cnpdthresholdconfigentry.Cnpdthresholdconfigstatus
    return leafs
}

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetYangName() string { return "cnpdThresholdConfigEntry" }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) SetParent(parent types.Entity) { cnpdthresholdconfigentry.parent = parent }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetParent() types.Entity { return cnpdthresholdconfigentry.parent }

func (cnpdthresholdconfigentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdconfigtable_Cnpdthresholdconfigentry) GetParentYangName() string { return "cnpdThresholdConfigTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry is created each time a threshold  is breached.   Thus there is
    // not necessarily a one to one  relationship to cnpdThresholdConfigTable  as
    // not every Threshold configured will  be breached. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry.
    Cnpdthresholdhistoryentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry
}

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetFilter() yfilter.YFilter { return cnpdthresholdhistorytable.YFilter }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) SetFilter(yf yfilter.YFilter) { cnpdthresholdhistorytable.YFilter = yf }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetGoName(yname string) string {
    if yname == "cnpdThresholdHistoryEntry" { return "Cnpdthresholdhistoryentry" }
    return ""
}

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetSegmentPath() string {
    return "cnpdThresholdHistoryTable"
}

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnpdThresholdHistoryEntry" {
        for _, c := range cnpdthresholdhistorytable.Cnpdthresholdhistoryentry {
            if cnpdthresholdhistorytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry{}
        cnpdthresholdhistorytable.Cnpdthresholdhistoryentry = append(cnpdthresholdhistorytable.Cnpdthresholdhistoryentry, child)
        return &cnpdthresholdhistorytable.Cnpdthresholdhistoryentry[len(cnpdthresholdhistorytable.Cnpdthresholdhistoryentry)-1]
    }
    return nil
}

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnpdthresholdhistorytable.Cnpdthresholdhistoryentry {
        children[cnpdthresholdhistorytable.Cnpdthresholdhistoryentry[i].GetSegmentPath()] = &cnpdthresholdhistorytable.Cnpdthresholdhistoryentry[i]
    }
    return children
}

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetYangName() string { return "cnpdThresholdHistoryTable" }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) SetParent(parent types.Entity) { cnpdthresholdhistorytable.parent = parent }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetParent() types.Entity { return cnpdthresholdhistorytable.parent }

func (cnpdthresholdhistorytable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry
// This entry is created each time a threshold 
// is breached. 
// 
// Thus there is not necessarily a one to one 
// relationship to cnpdThresholdConfigTable 
// as not every Threshold configured will 
// be breached.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry struct {
    parent types.Entity
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

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetFilter() yfilter.YFilter { return cnpdthresholdhistoryentry.YFilter }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) SetFilter(yf yfilter.YFilter) { cnpdthresholdhistoryentry.YFilter = yf }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetGoName(yname string) string {
    if yname == "cnpdThresholdHistoryIndex" { return "Cnpdthresholdhistoryindex" }
    if yname == "cnpdThresholdHistoryConfigIndex" { return "Cnpdthresholdhistoryconfigindex" }
    if yname == "cnpdThresholdHistoryValue" { return "Cnpdthresholdhistoryvalue" }
    if yname == "cnpdThresholdHistoryType" { return "Cnpdthresholdhistorytype" }
    if yname == "cnpdThresholdHistoryTime" { return "Cnpdthresholdhistorytime" }
    if yname == "cnpdThresholdHistoryProtocol" { return "Cnpdthresholdhistoryprotocol" }
    if yname == "cnpdThresholdHistoryStatsSelect" { return "Cnpdthresholdhistorystatsselect" }
    return ""
}

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetSegmentPath() string {
    return "cnpdThresholdHistoryEntry" + "[cnpdThresholdHistoryIndex='" + fmt.Sprintf("%v", cnpdthresholdhistoryentry.Cnpdthresholdhistoryindex) + "']"
}

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnpdThresholdHistoryIndex"] = cnpdthresholdhistoryentry.Cnpdthresholdhistoryindex
    leafs["cnpdThresholdHistoryConfigIndex"] = cnpdthresholdhistoryentry.Cnpdthresholdhistoryconfigindex
    leafs["cnpdThresholdHistoryValue"] = cnpdthresholdhistoryentry.Cnpdthresholdhistoryvalue
    leafs["cnpdThresholdHistoryType"] = cnpdthresholdhistoryentry.Cnpdthresholdhistorytype
    leafs["cnpdThresholdHistoryTime"] = cnpdthresholdhistoryentry.Cnpdthresholdhistorytime
    leafs["cnpdThresholdHistoryProtocol"] = cnpdthresholdhistoryentry.Cnpdthresholdhistoryprotocol
    leafs["cnpdThresholdHistoryStatsSelect"] = cnpdthresholdhistoryentry.Cnpdthresholdhistorystatsselect
    return leafs
}

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetYangName() string { return "cnpdThresholdHistoryEntry" }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) SetParent(parent types.Entity) { cnpdthresholdhistoryentry.parent = parent }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetParent() types.Entity { return cnpdthresholdhistoryentry.parent }

func (cnpdthresholdhistoryentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdthresholdhistorytable_Cnpdthresholdhistoryentry) GetParentYangName() string { return "cnpdThresholdHistoryTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A entry in the Supported Protocols table reflecting key information about a
    // protocol. The type is slice of
    // CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry.
    Cnpdsupportedprotocolsentry []CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry
}

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetFilter() yfilter.YFilter { return cnpdsupportedprotocolstable.YFilter }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) SetFilter(yf yfilter.YFilter) { cnpdsupportedprotocolstable.YFilter = yf }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetGoName(yname string) string {
    if yname == "cnpdSupportedProtocolsEntry" { return "Cnpdsupportedprotocolsentry" }
    return ""
}

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetSegmentPath() string {
    return "cnpdSupportedProtocolsTable"
}

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cnpdSupportedProtocolsEntry" {
        for _, c := range cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry {
            if cnpdsupportedprotocolstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry{}
        cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry = append(cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry, child)
        return &cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry[len(cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry)-1]
    }
    return nil
}

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry {
        children[cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry[i].GetSegmentPath()] = &cnpdsupportedprotocolstable.Cnpdsupportedprotocolsentry[i]
    }
    return children
}

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetYangName() string { return "cnpdSupportedProtocolsTable" }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) SetParent(parent types.Entity) { cnpdsupportedprotocolstable.parent = parent }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetParent() types.Entity { return cnpdsupportedprotocolstable.parent }

func (cnpdsupportedprotocolstable *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable) GetParentYangName() string { return "CISCO-NBAR-PROTOCOL-DISCOVERY-MIB" }

// CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry
// A entry in the Supported Protocols table reflecting
// key information about a protocol.
type CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A unique identifier of a row in this table.  Thus
    // it also represents a unique identifier for a protocol or application which
    // NBAR currently recognizes. The type is interface{} with range: 1..1024.
    Cnpdsupportedprotocolsindex interface{}

    // This object reflects the valid string of a protocol or application which
    // NBAR currently recognizes. The type is string with length: 1..255.
    Cnpdsupportedprotocolsname interface{}
}

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetFilter() yfilter.YFilter { return cnpdsupportedprotocolsentry.YFilter }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) SetFilter(yf yfilter.YFilter) { cnpdsupportedprotocolsentry.YFilter = yf }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetGoName(yname string) string {
    if yname == "cnpdSupportedProtocolsIndex" { return "Cnpdsupportedprotocolsindex" }
    if yname == "cnpdSupportedProtocolsName" { return "Cnpdsupportedprotocolsname" }
    return ""
}

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetSegmentPath() string {
    return "cnpdSupportedProtocolsEntry" + "[cnpdSupportedProtocolsIndex='" + fmt.Sprintf("%v", cnpdsupportedprotocolsentry.Cnpdsupportedprotocolsindex) + "']"
}

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cnpdSupportedProtocolsIndex"] = cnpdsupportedprotocolsentry.Cnpdsupportedprotocolsindex
    leafs["cnpdSupportedProtocolsName"] = cnpdsupportedprotocolsentry.Cnpdsupportedprotocolsname
    return leafs
}

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetYangName() string { return "cnpdSupportedProtocolsEntry" }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) SetParent(parent types.Entity) { cnpdsupportedprotocolsentry.parent = parent }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetParent() types.Entity { return cnpdsupportedprotocolsentry.parent }

func (cnpdsupportedprotocolsentry *CISCONBARPROTOCOLDISCOVERYMIB_Cnpdsupportedprotocolstable_Cnpdsupportedprotocolsentry) GetParentYangName() string { return "cnpdSupportedProtocolsTable" }

