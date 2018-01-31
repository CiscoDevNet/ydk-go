// Configuration management MIB.
// 
// The MIB represents a model of configuration data that
// exists in various locations:
// 
// running       in use by the running system
// terminal      saved to whatever is attached as the terminal        
// local         saved locally in NVRAM or flash
// remote        saved to some server on the network
// 
// Although some of the system functions that relate here
// can be used for general file storage and transfer, this
// MIB intends to include only such operations as clearly
// relate to configuration.  Its primary emphasis is to
// track changes and saves of the running configuration.
// 
// As saved data moves further from startup use, such as
// into different local flash files or onto the network,
// tracking becomes difficult to impossible, so the MIB's
// interest and functions are confined in that area.
// 
// Information from ccmCLIHistoryCommandTable can be used
// to track the exact configuration changes that took
// place within a particular Configuration History
// event. NMS' can use this information to update 
// the related components. 
// For example:
//     If commands related only to MPLS are entered
//     then the NMS need to update only the MPLS related
//     management information rather than updating
//     all of its management information.
//     Acronyms and terms:
// 
//     CLI   Command Line Interface.
package cisco_config_man_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_config_man_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-CONFIG-MAN-MIB CISCO-CONFIG-MAN-MIB}", reflect.TypeOf(CISCOCONFIGMANMIB{}))
    ydk.RegisterEntity("CISCO-CONFIG-MAN-MIB:CISCO-CONFIG-MAN-MIB", reflect.TypeOf(CISCOCONFIGMANMIB{}))
}

// HistoryEventMedium represents networkScp       network host via Secure Copy
type HistoryEventMedium string

const (
    HistoryEventMedium_erase HistoryEventMedium = "erase"

    HistoryEventMedium_commandSource HistoryEventMedium = "commandSource"

    HistoryEventMedium_running HistoryEventMedium = "running"

    HistoryEventMedium_startup HistoryEventMedium = "startup"

    HistoryEventMedium_local HistoryEventMedium = "local"

    HistoryEventMedium_networkTftp HistoryEventMedium = "networkTftp"

    HistoryEventMedium_networkRcp HistoryEventMedium = "networkRcp"

    HistoryEventMedium_networkFtp HistoryEventMedium = "networkFtp"

    HistoryEventMedium_networkScp HistoryEventMedium = "networkScp"
)

// CISCOCONFIGMANMIB
type CISCOCONFIGMANMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ccmhistory CISCOCONFIGMANMIB_Ccmhistory

    
    Ccmclihistory CISCOCONFIGMANMIB_Ccmclihistory

    
    Ccmclicfg CISCOCONFIGMANMIB_Ccmclicfg

    
    Ccmctidobjects CISCOCONFIGMANMIB_Ccmctidobjects

    // A table of configuration events on this router.
    Ccmhistoryeventtable CISCOCONFIGMANMIB_Ccmhistoryeventtable

    // A table of CLI commands that took effect during configuration events.
    Ccmclihistorycommandtable CISCOCONFIGMANMIB_Ccmclihistorycommandtable
}

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetFilter() yfilter.YFilter { return cISCOCONFIGMANMIB.YFilter }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) SetFilter(yf yfilter.YFilter) { cISCOCONFIGMANMIB.YFilter = yf }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetGoName(yname string) string {
    if yname == "ccmHistory" { return "Ccmhistory" }
    if yname == "ccmCLIHistory" { return "Ccmclihistory" }
    if yname == "ccmCLICfg" { return "Ccmclicfg" }
    if yname == "ccmCTIDObjects" { return "Ccmctidobjects" }
    if yname == "ccmHistoryEventTable" { return "Ccmhistoryeventtable" }
    if yname == "ccmCLIHistoryCommandTable" { return "Ccmclihistorycommandtable" }
    return ""
}

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetSegmentPath() string {
    return "CISCO-CONFIG-MAN-MIB:CISCO-CONFIG-MAN-MIB"
}

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccmHistory" {
        return &cISCOCONFIGMANMIB.Ccmhistory
    }
    if childYangName == "ccmCLIHistory" {
        return &cISCOCONFIGMANMIB.Ccmclihistory
    }
    if childYangName == "ccmCLICfg" {
        return &cISCOCONFIGMANMIB.Ccmclicfg
    }
    if childYangName == "ccmCTIDObjects" {
        return &cISCOCONFIGMANMIB.Ccmctidobjects
    }
    if childYangName == "ccmHistoryEventTable" {
        return &cISCOCONFIGMANMIB.Ccmhistoryeventtable
    }
    if childYangName == "ccmCLIHistoryCommandTable" {
        return &cISCOCONFIGMANMIB.Ccmclihistorycommandtable
    }
    return nil
}

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ccmHistory"] = &cISCOCONFIGMANMIB.Ccmhistory
    children["ccmCLIHistory"] = &cISCOCONFIGMANMIB.Ccmclihistory
    children["ccmCLICfg"] = &cISCOCONFIGMANMIB.Ccmclicfg
    children["ccmCTIDObjects"] = &cISCOCONFIGMANMIB.Ccmctidobjects
    children["ccmHistoryEventTable"] = &cISCOCONFIGMANMIB.Ccmhistoryeventtable
    children["ccmCLIHistoryCommandTable"] = &cISCOCONFIGMANMIB.Ccmclihistorycommandtable
    return children
}

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetYangName() string { return "CISCO-CONFIG-MAN-MIB" }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) SetParent(parent types.Entity) { cISCOCONFIGMANMIB.parent = parent }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetParent() types.Entity { return cISCOCONFIGMANMIB.parent }

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetParentYangName() string { return "CISCO-CONFIG-MAN-MIB" }

// CISCOCONFIGMANMIB_Ccmhistory
type CISCOCONFIGMANMIB_Ccmhistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The value of sysUpTime when the running configuration was last changed.    
    // If the value of ccmHistoryRunningLastChanged is         greater than
    // ccmHistoryRunningLastSaved, the          configuration has been changed but
    // not saved. The type is interface{} with range: 0..4294967295.
    Ccmhistoryrunninglastchanged interface{}

    // The value of sysUpTime when the running configuration was last saved
    // (written).  If the value of ccmHistoryRunningLastChanged is  greater than
    // ccmHistoryRunningLastSaved, the  configuration has been changed but not
    // saved.  What constitutes a safe saving of the running configuration is a
    // management policy issue beyond the scope of this MIB.  For some
    // installations, writing the running configuration to a terminal may be a way
    // of capturing and saving it.  Others may use local or remote storage.  Thus
    // ANY write is considered saving for the purposes of the MIB. The type is
    // interface{} with range: 0..4294967295.
    Ccmhistoryrunninglastsaved interface{}

    // The value of sysUpTime when the startup configuration was last written to. 
    // In general this is the default configuration used when cold starting the
    // system.  It may have been changed by a save of the running configuration or
    // by a copy from elsewhere. The type is interface{} with range:
    // 0..4294967295.
    Ccmhistorystartuplastchanged interface{}

    // The maximum number of entries that can be held in ccmHistoryEventTable. 
    // The recommended value for implementations is 10. The type is interface{}
    // with range: 0..2147483647.
    Ccmhistorymaxevententries interface{}

    // The number of times the oldest entry in ccmHistoryEventTable was deleted to
    // make room  for a new entry. The type is interface{} with range:
    // 0..4294967295.
    Ccmhistoryevententriesbumped interface{}
}

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetFilter() yfilter.YFilter { return ccmhistory.YFilter }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) SetFilter(yf yfilter.YFilter) { ccmhistory.YFilter = yf }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetGoName(yname string) string {
    if yname == "ccmHistoryRunningLastChanged" { return "Ccmhistoryrunninglastchanged" }
    if yname == "ccmHistoryRunningLastSaved" { return "Ccmhistoryrunninglastsaved" }
    if yname == "ccmHistoryStartupLastChanged" { return "Ccmhistorystartuplastchanged" }
    if yname == "ccmHistoryMaxEventEntries" { return "Ccmhistorymaxevententries" }
    if yname == "ccmHistoryEventEntriesBumped" { return "Ccmhistoryevententriesbumped" }
    return ""
}

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetSegmentPath() string {
    return "ccmHistory"
}

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccmHistoryRunningLastChanged"] = ccmhistory.Ccmhistoryrunninglastchanged
    leafs["ccmHistoryRunningLastSaved"] = ccmhistory.Ccmhistoryrunninglastsaved
    leafs["ccmHistoryStartupLastChanged"] = ccmhistory.Ccmhistorystartuplastchanged
    leafs["ccmHistoryMaxEventEntries"] = ccmhistory.Ccmhistorymaxevententries
    leafs["ccmHistoryEventEntriesBumped"] = ccmhistory.Ccmhistoryevententriesbumped
    return leafs
}

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetBundleName() string { return "cisco_ios_xe" }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetYangName() string { return "ccmHistory" }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) SetParent(parent types.Entity) { ccmhistory.parent = parent }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetParent() types.Entity { return ccmhistory.parent }

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetParentYangName() string { return "CISCO-CONFIG-MAN-MIB" }

// CISCOCONFIGMANMIB_Ccmclihistory
type CISCOCONFIGMANMIB_Ccmclihistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The maximum number of entries that can be held in
    // ccmCLIHistoryCommandTable.  The recommended value for implementations is
    // 100.  If the number of entries in ccmCLIHistoryCommandTable  exceeds the
    // value of this object, old entries will be  bumped to make room for new
    // entries.  The ccmCLIHistoryCommandTable will not be populated if the value
    // of this object is 0. The type is interface{} with range: 0..4294967295.
    Ccmclihistorymaxcmdentries interface{}

    // The current number of entries in ccmCLIHistoryCommandTable. The type is
    // interface{} with range: 0..4294967295.
    Ccmclihistorycmdentries interface{}

    // This object indicates the upper limit on the number of entries allowed in 
    // ccmCLIHistoryCommandTable by the managed system. The type is interface{}
    // with range: 0..4294967295.
    Ccmclihistorycmdentriesallowed interface{}
}

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetFilter() yfilter.YFilter { return ccmclihistory.YFilter }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) SetFilter(yf yfilter.YFilter) { ccmclihistory.YFilter = yf }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetGoName(yname string) string {
    if yname == "ccmCLIHistoryMaxCmdEntries" { return "Ccmclihistorymaxcmdentries" }
    if yname == "ccmCLIHistoryCmdEntries" { return "Ccmclihistorycmdentries" }
    if yname == "ccmCLIHistoryCmdEntriesAllowed" { return "Ccmclihistorycmdentriesallowed" }
    return ""
}

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetSegmentPath() string {
    return "ccmCLIHistory"
}

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccmCLIHistoryMaxCmdEntries"] = ccmclihistory.Ccmclihistorymaxcmdentries
    leafs["ccmCLIHistoryCmdEntries"] = ccmclihistory.Ccmclihistorycmdentries
    leafs["ccmCLIHistoryCmdEntriesAllowed"] = ccmclihistory.Ccmclihistorycmdentriesallowed
    return leafs
}

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetBundleName() string { return "cisco_ios_xe" }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetYangName() string { return "ccmCLIHistory" }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) SetParent(parent types.Entity) { ccmclihistory.parent = parent }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetParent() types.Entity { return ccmclihistory.parent }

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetParentYangName() string { return "CISCO-CONFIG-MAN-MIB" }

// CISCOCONFIGMANMIB_Ccmclicfg
type CISCOCONFIGMANMIB_Ccmclicfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This variable indicates whether the system produces the
    // ccmCLIRunningConfigChanged notification. A false  value will prevent
    // notifications from being generated  by this system. The type is bool.
    Ccmclicfgrunconfnotifenable interface{}
}

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetFilter() yfilter.YFilter { return ccmclicfg.YFilter }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) SetFilter(yf yfilter.YFilter) { ccmclicfg.YFilter = yf }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetGoName(yname string) string {
    if yname == "ccmCLICfgRunConfNotifEnable" { return "Ccmclicfgrunconfnotifenable" }
    return ""
}

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetSegmentPath() string {
    return "ccmCLICfg"
}

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccmCLICfgRunConfNotifEnable"] = ccmclicfg.Ccmclicfgrunconfnotifenable
    return leafs
}

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetBundleName() string { return "cisco_ios_xe" }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetYangName() string { return "ccmCLICfg" }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) SetParent(parent types.Entity) { ccmclicfg.parent = parent }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetParent() types.Entity { return ccmclicfg.parent }

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetParentYangName() string { return "CISCO-CONFIG-MAN-MIB" }

// CISCOCONFIGMANMIB_Ccmctidobjects
type CISCOCONFIGMANMIB_Ccmctidobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object indicates the Config Change Tracking ID which uniquely
    // represents version-incrementing changes to the IOS  running configuration.
    // The type is interface{} with range: 0..18446744073709551615.
    Ccmctid interface{}

    // This object indicates the time when the Config Change Tracking ID last
    // changed. The type is string.
    Ccmctidlastchangetime interface{}

    // This object indicates the user who last reset the Config Change Tracking
    // ID. The type is string.
    Ccmctidwhochanged interface{}

    // This variable indicates whether the system produces the ccmCTIDRolledOver
    // notification. A false value will prevent notifications from being generated
    // by this system. The type is bool.
    Ccmctidrolledovernotifenable interface{}
}

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetFilter() yfilter.YFilter { return ccmctidobjects.YFilter }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) SetFilter(yf yfilter.YFilter) { ccmctidobjects.YFilter = yf }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetGoName(yname string) string {
    if yname == "ccmCTID" { return "Ccmctid" }
    if yname == "ccmCTIDLastChangeTime" { return "Ccmctidlastchangetime" }
    if yname == "ccmCTIDWhoChanged" { return "Ccmctidwhochanged" }
    if yname == "ccmCTIDRolledOverNotifEnable" { return "Ccmctidrolledovernotifenable" }
    return ""
}

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetSegmentPath() string {
    return "ccmCTIDObjects"
}

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccmCTID"] = ccmctidobjects.Ccmctid
    leafs["ccmCTIDLastChangeTime"] = ccmctidobjects.Ccmctidlastchangetime
    leafs["ccmCTIDWhoChanged"] = ccmctidobjects.Ccmctidwhochanged
    leafs["ccmCTIDRolledOverNotifEnable"] = ccmctidobjects.Ccmctidrolledovernotifenable
    return leafs
}

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetBundleName() string { return "cisco_ios_xe" }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetYangName() string { return "ccmCTIDObjects" }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) SetParent(parent types.Entity) { ccmctidobjects.parent = parent }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetParent() types.Entity { return ccmctidobjects.parent }

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetParentYangName() string { return "CISCO-CONFIG-MAN-MIB" }

// CISCOCONFIGMANMIB_Ccmhistoryeventtable
// A table of configuration events on this router.
type CISCOCONFIGMANMIB_Ccmhistoryeventtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a configuration event on this router. The type is slice
    // of CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry.
    Ccmhistoryevententry []CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry
}

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetFilter() yfilter.YFilter { return ccmhistoryeventtable.YFilter }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) SetFilter(yf yfilter.YFilter) { ccmhistoryeventtable.YFilter = yf }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetGoName(yname string) string {
    if yname == "ccmHistoryEventEntry" { return "Ccmhistoryevententry" }
    return ""
}

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetSegmentPath() string {
    return "ccmHistoryEventTable"
}

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccmHistoryEventEntry" {
        for _, c := range ccmhistoryeventtable.Ccmhistoryevententry {
            if ccmhistoryeventtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry{}
        ccmhistoryeventtable.Ccmhistoryevententry = append(ccmhistoryeventtable.Ccmhistoryevententry, child)
        return &ccmhistoryeventtable.Ccmhistoryevententry[len(ccmhistoryeventtable.Ccmhistoryevententry)-1]
    }
    return nil
}

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccmhistoryeventtable.Ccmhistoryevententry {
        children[ccmhistoryeventtable.Ccmhistoryevententry[i].GetSegmentPath()] = &ccmhistoryeventtable.Ccmhistoryevententry[i]
    }
    return children
}

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetBundleName() string { return "cisco_ios_xe" }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetYangName() string { return "ccmHistoryEventTable" }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) SetParent(parent types.Entity) { ccmhistoryeventtable.parent = parent }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetParent() types.Entity { return ccmhistoryeventtable.parent }

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetParentYangName() string { return "CISCO-CONFIG-MAN-MIB" }

// CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry
// Information about a configuration event on this
// router.
type CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A monotonically increasing integer for the sole
    // purpose of indexing events.  When it reaches the  maximum value, an
    // extremely unlikely event, the agent  wraps the value back to 1 and may
    // flush existing  entries. The type is interface{} with range: 1..2147483647.
    Ccmhistoryeventindex interface{}

    // The value of sysUpTime when the event occurred. The type is interface{}
    // with range: 0..4294967295.
    Ccmhistoryeventtime interface{}

    // The source of the command that instigated the event. The type is
    // Ccmhistoryeventcommandsource.
    Ccmhistoryeventcommandsource interface{}

    // The configuration data source for the event. The type is
    // HistoryEventMedium.
    Ccmhistoryeventconfigsource interface{}

    // The configuration data destination for the event. The type is
    // HistoryEventMedium.
    Ccmhistoryeventconfigdestination interface{}

    // If ccmHistoryEventCommandSource is 'commandLine', the terminal type,
    // otherwise 'notApplicable'. The type is Ccmhistoryeventterminaltype.
    Ccmhistoryeventterminaltype interface{}

    // If ccmHistoryEventCommandSource is 'commandLine', the terminal number.  The
    // value is -1 if not available or not applicable. The type is interface{}
    // with range: -2147483648..2147483647.
    Ccmhistoryeventterminalnumber interface{}

    // If ccmHistoryEventCommandSource is 'commandLine', the name of the logged in
    // user.  The length is zero if not available or not applicable. The type is
    // string with length: 0..64.
    Ccmhistoryeventterminaluser interface{}

    // If ccmHistoryEventCommandSource is 'commandLine', the hard-wired location
    // of the terminal or the remote  host for an incoming connection.  The length
    // is zero  if not available or not applicable. The type is string with
    // length: 0..64.
    Ccmhistoryeventterminallocation interface{}

    // If ccmHistoryEventTerminalType is 'virtual', the internet address of the
    // connected system.  If ccmHistoryEventCommandSource is 'snmp', the internet
    // address of the requester.  The value is 0.0.0.0 if not available or not 
    // applicable.  This object is deprecated by
    // ccmHistoryEventCommandSourceAddrRev1. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ccmhistoryeventcommandsourceaddress interface{}

    // If ccmHistoryEventTerminalType is 'virtual', the host name of the connected
    // system.  The length is zero if not available or not applicable. The type is
    // string with length: 0..64.
    Ccmhistoryeventvirtualhostname interface{}

    // If ccmHistoryEventConfigSource or ccmHistoryEventConfigDestination is
    // 'networkTftp' or 'networkRcp', the internet address of the storage file
    // server.  The value is 0.0.0.0 if not applicable or not         available.  
    // This object is deprecated by         ccmHistoryEventServerAddrRev1. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ccmhistoryeventserveraddress interface{}

    // If ccmHistoryEventConfigSource or ccmHistoryEventConfigDestination is
    // 'networkTftp' or 'networkRcp', the configuration file name at the storage
    // file server.  The length is zero if not available or not applicable. The
    // type is string with length: 0..64.
    Ccmhistoryeventfile interface{}

    // If ccmHistoryEventConfigSource or ccmHistoryEventConfigDestination is
    // 'networkRcp', the remote user name.  The length is zero if not applicable
    // or not available. The type is string with length: 0..64.
    Ccmhistoryeventrcpuser interface{}

    // The number of times the oldest entry in ccmCLIHistoryCommandTable with
    // first index as  ccmHistoryEventIndex was deleted to make  room for a new
    // entry.  This object is applicable only if  ccmHistoryEventCommandSource has
    // a value  of 'commandLine'. The type is interface{} with range:
    // 0..4294967295.
    Ccmhistoryclicmdentriesbumped interface{}

    // This object indicates the transport type of the address contained in
    // ccmHistoryEventCommandSourceAddrRev1.  The value will be zero if not
    // available or not applicable. The type is InetAddressType.
    Ccmhistoryeventcommandsourceaddrtype interface{}

    // If ccmHistoryEventTerminalType is 'virtual', the internet address of the
    // connected system.  If ccmHistoryEventCommandSource is 'snmp', the internet
    // address of the requester.  The value of all bit's is zero  if not available
    // or not applicable.  The Format of this address depends on the value of the
    // ccmHistoryEventCommandSourceAddrType object.  This object deprecates
    // ccmHistoryEventCommandSourceAddress. The type is string with length:
    // 0..255.
    Ccmhistoryeventcommandsourceaddrrev1 interface{}

    // This object indicates the transport type of the address contained in
    // ccmHistoryEventServerAddrRev1.  The value will be zero if not available or
    // not aplicable. The type is InetAddressType.
    Ccmhistoryeventserveraddrtype interface{}

    // If ccmHistoryEventConfigSource or ccmHistoryEventConfigDestination is
    // 'networkTftp' or 'networkRcp', the internet address of the storage file
    // server.   The value of all bits is 0s if not applicable or not available. 
    // The Format of this address depends on the value of the
    // ccmHistoryEventServerAddrType object.  This object deprecates
    // ccmHistoryEventServerAddress. The type is string with length: 0..255.
    Ccmhistoryeventserveraddrrev1 interface{}
}

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetFilter() yfilter.YFilter { return ccmhistoryevententry.YFilter }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) SetFilter(yf yfilter.YFilter) { ccmhistoryevententry.YFilter = yf }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetGoName(yname string) string {
    if yname == "ccmHistoryEventIndex" { return "Ccmhistoryeventindex" }
    if yname == "ccmHistoryEventTime" { return "Ccmhistoryeventtime" }
    if yname == "ccmHistoryEventCommandSource" { return "Ccmhistoryeventcommandsource" }
    if yname == "ccmHistoryEventConfigSource" { return "Ccmhistoryeventconfigsource" }
    if yname == "ccmHistoryEventConfigDestination" { return "Ccmhistoryeventconfigdestination" }
    if yname == "ccmHistoryEventTerminalType" { return "Ccmhistoryeventterminaltype" }
    if yname == "ccmHistoryEventTerminalNumber" { return "Ccmhistoryeventterminalnumber" }
    if yname == "ccmHistoryEventTerminalUser" { return "Ccmhistoryeventterminaluser" }
    if yname == "ccmHistoryEventTerminalLocation" { return "Ccmhistoryeventterminallocation" }
    if yname == "ccmHistoryEventCommandSourceAddress" { return "Ccmhistoryeventcommandsourceaddress" }
    if yname == "ccmHistoryEventVirtualHostName" { return "Ccmhistoryeventvirtualhostname" }
    if yname == "ccmHistoryEventServerAddress" { return "Ccmhistoryeventserveraddress" }
    if yname == "ccmHistoryEventFile" { return "Ccmhistoryeventfile" }
    if yname == "ccmHistoryEventRcpUser" { return "Ccmhistoryeventrcpuser" }
    if yname == "ccmHistoryCLICmdEntriesBumped" { return "Ccmhistoryclicmdentriesbumped" }
    if yname == "ccmHistoryEventCommandSourceAddrType" { return "Ccmhistoryeventcommandsourceaddrtype" }
    if yname == "ccmHistoryEventCommandSourceAddrRev1" { return "Ccmhistoryeventcommandsourceaddrrev1" }
    if yname == "ccmHistoryEventServerAddrType" { return "Ccmhistoryeventserveraddrtype" }
    if yname == "ccmHistoryEventServerAddrRev1" { return "Ccmhistoryeventserveraddrrev1" }
    return ""
}

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetSegmentPath() string {
    return "ccmHistoryEventEntry" + "[ccmHistoryEventIndex='" + fmt.Sprintf("%v", ccmhistoryevententry.Ccmhistoryeventindex) + "']"
}

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccmHistoryEventIndex"] = ccmhistoryevententry.Ccmhistoryeventindex
    leafs["ccmHistoryEventTime"] = ccmhistoryevententry.Ccmhistoryeventtime
    leafs["ccmHistoryEventCommandSource"] = ccmhistoryevententry.Ccmhistoryeventcommandsource
    leafs["ccmHistoryEventConfigSource"] = ccmhistoryevententry.Ccmhistoryeventconfigsource
    leafs["ccmHistoryEventConfigDestination"] = ccmhistoryevententry.Ccmhistoryeventconfigdestination
    leafs["ccmHistoryEventTerminalType"] = ccmhistoryevententry.Ccmhistoryeventterminaltype
    leafs["ccmHistoryEventTerminalNumber"] = ccmhistoryevententry.Ccmhistoryeventterminalnumber
    leafs["ccmHistoryEventTerminalUser"] = ccmhistoryevententry.Ccmhistoryeventterminaluser
    leafs["ccmHistoryEventTerminalLocation"] = ccmhistoryevententry.Ccmhistoryeventterminallocation
    leafs["ccmHistoryEventCommandSourceAddress"] = ccmhistoryevententry.Ccmhistoryeventcommandsourceaddress
    leafs["ccmHistoryEventVirtualHostName"] = ccmhistoryevententry.Ccmhistoryeventvirtualhostname
    leafs["ccmHistoryEventServerAddress"] = ccmhistoryevententry.Ccmhistoryeventserveraddress
    leafs["ccmHistoryEventFile"] = ccmhistoryevententry.Ccmhistoryeventfile
    leafs["ccmHistoryEventRcpUser"] = ccmhistoryevententry.Ccmhistoryeventrcpuser
    leafs["ccmHistoryCLICmdEntriesBumped"] = ccmhistoryevententry.Ccmhistoryclicmdentriesbumped
    leafs["ccmHistoryEventCommandSourceAddrType"] = ccmhistoryevententry.Ccmhistoryeventcommandsourceaddrtype
    leafs["ccmHistoryEventCommandSourceAddrRev1"] = ccmhistoryevententry.Ccmhistoryeventcommandsourceaddrrev1
    leafs["ccmHistoryEventServerAddrType"] = ccmhistoryevententry.Ccmhistoryeventserveraddrtype
    leafs["ccmHistoryEventServerAddrRev1"] = ccmhistoryevententry.Ccmhistoryeventserveraddrrev1
    return leafs
}

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetBundleName() string { return "cisco_ios_xe" }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetYangName() string { return "ccmHistoryEventEntry" }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) SetParent(parent types.Entity) { ccmhistoryevententry.parent = parent }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetParent() types.Entity { return ccmhistoryevententry.parent }

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetParentYangName() string { return "ccmHistoryEventTable" }

// CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventcommandsource represents The source of the command that instigated the event.
type CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventcommandsource string

const (
    CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventcommandsource_commandLine CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventcommandsource = "commandLine"

    CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventcommandsource_snmp CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventcommandsource = "snmp"
)

// CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype represents the terminal type, otherwise 'notApplicable'.
type CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype string

const (
    CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype_notApplicable CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype = "notApplicable"

    CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype_unknown CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype = "unknown"

    CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype_console CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype = "console"

    CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype_terminal CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype = "terminal"

    CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype_virtual CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype = "virtual"

    CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype_auxiliary CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventterminaltype = "auxiliary"
)

// CISCOCONFIGMANMIB_Ccmclihistorycommandtable
// A table of CLI commands that took effect during
// configuration events.
type CISCOCONFIGMANMIB_Ccmclihistorycommandtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about the CLI commands that took effect during the
    // configuration event pointed by  ccmCLIHistoryEventIndex.  A set of rows in
    // this table having the first index as ccmHistoryEventIndex will store the
    // CLI commands entered during the corresponding  configuration event in
    // ccmHistoryEventTable.  An entry will be created in this table only if  the
    // corresponding entry in ccmHistoryEventTable has  a value of 'commandLine'
    // for  ccmHistoryEventCommandSource. The type is slice of
    // CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry.
    Ccmclihistorycommandentry []CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry
}

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetFilter() yfilter.YFilter { return ccmclihistorycommandtable.YFilter }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) SetFilter(yf yfilter.YFilter) { ccmclihistorycommandtable.YFilter = yf }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetGoName(yname string) string {
    if yname == "ccmCLIHistoryCommandEntry" { return "Ccmclihistorycommandentry" }
    return ""
}

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetSegmentPath() string {
    return "ccmCLIHistoryCommandTable"
}

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccmCLIHistoryCommandEntry" {
        for _, c := range ccmclihistorycommandtable.Ccmclihistorycommandentry {
            if ccmclihistorycommandtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry{}
        ccmclihistorycommandtable.Ccmclihistorycommandentry = append(ccmclihistorycommandtable.Ccmclihistorycommandentry, child)
        return &ccmclihistorycommandtable.Ccmclihistorycommandentry[len(ccmclihistorycommandtable.Ccmclihistorycommandentry)-1]
    }
    return nil
}

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccmclihistorycommandtable.Ccmclihistorycommandentry {
        children[ccmclihistorycommandtable.Ccmclihistorycommandentry[i].GetSegmentPath()] = &ccmclihistorycommandtable.Ccmclihistorycommandentry[i]
    }
    return children
}

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetBundleName() string { return "cisco_ios_xe" }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetYangName() string { return "ccmCLIHistoryCommandTable" }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) SetParent(parent types.Entity) { ccmclihistorycommandtable.parent = parent }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetParent() types.Entity { return ccmclihistorycommandtable.parent }

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetParentYangName() string { return "CISCO-CONFIG-MAN-MIB" }

// CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry
// Information about the CLI commands that took effect
// during the configuration event pointed by 
// ccmCLIHistoryEventIndex.
// 
// A set of rows in this table having the first
// index as ccmHistoryEventIndex will store the
// CLI commands entered during the corresponding 
// configuration event in ccmHistoryEventTable.
// 
// An entry will be created in this table only if 
// the corresponding entry in ccmHistoryEventTable has 
// a value of 'commandLine' for 
// ccmHistoryEventCommandSource.
type CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_config_man_mib.CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry_Ccmhistoryeventindex
    Ccmhistoryeventindex interface{}

    // This attribute is a key. A monotonically increasing integer for the purpose
    // of indexing CLI commands which took effect during a configuration event.
    // The type is interface{} with range: 1..4294967295.
    Ccmclihistorycommandindex interface{}

    // The CLI command entered which took effect during the configuration event
    // pointed by  ccmHistoryEventIndex. The type is string.
    Ccmclihistorycommand interface{}
}

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetFilter() yfilter.YFilter { return ccmclihistorycommandentry.YFilter }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) SetFilter(yf yfilter.YFilter) { ccmclihistorycommandentry.YFilter = yf }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetGoName(yname string) string {
    if yname == "ccmHistoryEventIndex" { return "Ccmhistoryeventindex" }
    if yname == "ccmCLIHistoryCommandIndex" { return "Ccmclihistorycommandindex" }
    if yname == "ccmCLIHistoryCommand" { return "Ccmclihistorycommand" }
    return ""
}

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetSegmentPath() string {
    return "ccmCLIHistoryCommandEntry" + "[ccmHistoryEventIndex='" + fmt.Sprintf("%v", ccmclihistorycommandentry.Ccmhistoryeventindex) + "']" + "[ccmCLIHistoryCommandIndex='" + fmt.Sprintf("%v", ccmclihistorycommandentry.Ccmclihistorycommandindex) + "']"
}

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccmHistoryEventIndex"] = ccmclihistorycommandentry.Ccmhistoryeventindex
    leafs["ccmCLIHistoryCommandIndex"] = ccmclihistorycommandentry.Ccmclihistorycommandindex
    leafs["ccmCLIHistoryCommand"] = ccmclihistorycommandentry.Ccmclihistorycommand
    return leafs
}

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetBundleName() string { return "cisco_ios_xe" }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetYangName() string { return "ccmCLIHistoryCommandEntry" }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) SetParent(parent types.Entity) { ccmclihistorycommandentry.parent = parent }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetParent() types.Entity { return ccmclihistorycommandentry.parent }

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetParentYangName() string { return "ccmCLIHistoryCommandTable" }

