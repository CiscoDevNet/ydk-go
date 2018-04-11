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
    EntityData types.CommonEntityData
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

func (cISCOCONFIGMANMIB *CISCOCONFIGMANMIB) GetEntityData() *types.CommonEntityData {
    cISCOCONFIGMANMIB.EntityData.YFilter = cISCOCONFIGMANMIB.YFilter
    cISCOCONFIGMANMIB.EntityData.YangName = "CISCO-CONFIG-MAN-MIB"
    cISCOCONFIGMANMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCONFIGMANMIB.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    cISCOCONFIGMANMIB.EntityData.SegmentPath = "CISCO-CONFIG-MAN-MIB:CISCO-CONFIG-MAN-MIB"
    cISCOCONFIGMANMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCONFIGMANMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCONFIGMANMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCONFIGMANMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOCONFIGMANMIB.EntityData.Children["ccmHistory"] = types.YChild{"Ccmhistory", &cISCOCONFIGMANMIB.Ccmhistory}
    cISCOCONFIGMANMIB.EntityData.Children["ccmCLIHistory"] = types.YChild{"Ccmclihistory", &cISCOCONFIGMANMIB.Ccmclihistory}
    cISCOCONFIGMANMIB.EntityData.Children["ccmCLICfg"] = types.YChild{"Ccmclicfg", &cISCOCONFIGMANMIB.Ccmclicfg}
    cISCOCONFIGMANMIB.EntityData.Children["ccmCTIDObjects"] = types.YChild{"Ccmctidobjects", &cISCOCONFIGMANMIB.Ccmctidobjects}
    cISCOCONFIGMANMIB.EntityData.Children["ccmHistoryEventTable"] = types.YChild{"Ccmhistoryeventtable", &cISCOCONFIGMANMIB.Ccmhistoryeventtable}
    cISCOCONFIGMANMIB.EntityData.Children["ccmCLIHistoryCommandTable"] = types.YChild{"Ccmclihistorycommandtable", &cISCOCONFIGMANMIB.Ccmclihistorycommandtable}
    cISCOCONFIGMANMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOCONFIGMANMIB.EntityData)
}

// CISCOCONFIGMANMIB_Ccmhistory
type CISCOCONFIGMANMIB_Ccmhistory struct {
    EntityData types.CommonEntityData
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

func (ccmhistory *CISCOCONFIGMANMIB_Ccmhistory) GetEntityData() *types.CommonEntityData {
    ccmhistory.EntityData.YFilter = ccmhistory.YFilter
    ccmhistory.EntityData.YangName = "ccmHistory"
    ccmhistory.EntityData.BundleName = "cisco_ios_xe"
    ccmhistory.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmhistory.EntityData.SegmentPath = "ccmHistory"
    ccmhistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmhistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmhistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmhistory.EntityData.Children = make(map[string]types.YChild)
    ccmhistory.EntityData.Leafs = make(map[string]types.YLeaf)
    ccmhistory.EntityData.Leafs["ccmHistoryRunningLastChanged"] = types.YLeaf{"Ccmhistoryrunninglastchanged", ccmhistory.Ccmhistoryrunninglastchanged}
    ccmhistory.EntityData.Leafs["ccmHistoryRunningLastSaved"] = types.YLeaf{"Ccmhistoryrunninglastsaved", ccmhistory.Ccmhistoryrunninglastsaved}
    ccmhistory.EntityData.Leafs["ccmHistoryStartupLastChanged"] = types.YLeaf{"Ccmhistorystartuplastchanged", ccmhistory.Ccmhistorystartuplastchanged}
    ccmhistory.EntityData.Leafs["ccmHistoryMaxEventEntries"] = types.YLeaf{"Ccmhistorymaxevententries", ccmhistory.Ccmhistorymaxevententries}
    ccmhistory.EntityData.Leafs["ccmHistoryEventEntriesBumped"] = types.YLeaf{"Ccmhistoryevententriesbumped", ccmhistory.Ccmhistoryevententriesbumped}
    return &(ccmhistory.EntityData)
}

// CISCOCONFIGMANMIB_Ccmclihistory
type CISCOCONFIGMANMIB_Ccmclihistory struct {
    EntityData types.CommonEntityData
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

func (ccmclihistory *CISCOCONFIGMANMIB_Ccmclihistory) GetEntityData() *types.CommonEntityData {
    ccmclihistory.EntityData.YFilter = ccmclihistory.YFilter
    ccmclihistory.EntityData.YangName = "ccmCLIHistory"
    ccmclihistory.EntityData.BundleName = "cisco_ios_xe"
    ccmclihistory.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmclihistory.EntityData.SegmentPath = "ccmCLIHistory"
    ccmclihistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmclihistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmclihistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmclihistory.EntityData.Children = make(map[string]types.YChild)
    ccmclihistory.EntityData.Leafs = make(map[string]types.YLeaf)
    ccmclihistory.EntityData.Leafs["ccmCLIHistoryMaxCmdEntries"] = types.YLeaf{"Ccmclihistorymaxcmdentries", ccmclihistory.Ccmclihistorymaxcmdentries}
    ccmclihistory.EntityData.Leafs["ccmCLIHistoryCmdEntries"] = types.YLeaf{"Ccmclihistorycmdentries", ccmclihistory.Ccmclihistorycmdentries}
    ccmclihistory.EntityData.Leafs["ccmCLIHistoryCmdEntriesAllowed"] = types.YLeaf{"Ccmclihistorycmdentriesallowed", ccmclihistory.Ccmclihistorycmdentriesallowed}
    return &(ccmclihistory.EntityData)
}

// CISCOCONFIGMANMIB_Ccmclicfg
type CISCOCONFIGMANMIB_Ccmclicfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable indicates whether the system produces the
    // ccmCLIRunningConfigChanged notification. A false  value will prevent
    // notifications from being generated  by this system. The type is bool.
    Ccmclicfgrunconfnotifenable interface{}
}

func (ccmclicfg *CISCOCONFIGMANMIB_Ccmclicfg) GetEntityData() *types.CommonEntityData {
    ccmclicfg.EntityData.YFilter = ccmclicfg.YFilter
    ccmclicfg.EntityData.YangName = "ccmCLICfg"
    ccmclicfg.EntityData.BundleName = "cisco_ios_xe"
    ccmclicfg.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmclicfg.EntityData.SegmentPath = "ccmCLICfg"
    ccmclicfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmclicfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmclicfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmclicfg.EntityData.Children = make(map[string]types.YChild)
    ccmclicfg.EntityData.Leafs = make(map[string]types.YLeaf)
    ccmclicfg.EntityData.Leafs["ccmCLICfgRunConfNotifEnable"] = types.YLeaf{"Ccmclicfgrunconfnotifenable", ccmclicfg.Ccmclicfgrunconfnotifenable}
    return &(ccmclicfg.EntityData)
}

// CISCOCONFIGMANMIB_Ccmctidobjects
type CISCOCONFIGMANMIB_Ccmctidobjects struct {
    EntityData types.CommonEntityData
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

func (ccmctidobjects *CISCOCONFIGMANMIB_Ccmctidobjects) GetEntityData() *types.CommonEntityData {
    ccmctidobjects.EntityData.YFilter = ccmctidobjects.YFilter
    ccmctidobjects.EntityData.YangName = "ccmCTIDObjects"
    ccmctidobjects.EntityData.BundleName = "cisco_ios_xe"
    ccmctidobjects.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmctidobjects.EntityData.SegmentPath = "ccmCTIDObjects"
    ccmctidobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmctidobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmctidobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmctidobjects.EntityData.Children = make(map[string]types.YChild)
    ccmctidobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    ccmctidobjects.EntityData.Leafs["ccmCTID"] = types.YLeaf{"Ccmctid", ccmctidobjects.Ccmctid}
    ccmctidobjects.EntityData.Leafs["ccmCTIDLastChangeTime"] = types.YLeaf{"Ccmctidlastchangetime", ccmctidobjects.Ccmctidlastchangetime}
    ccmctidobjects.EntityData.Leafs["ccmCTIDWhoChanged"] = types.YLeaf{"Ccmctidwhochanged", ccmctidobjects.Ccmctidwhochanged}
    ccmctidobjects.EntityData.Leafs["ccmCTIDRolledOverNotifEnable"] = types.YLeaf{"Ccmctidrolledovernotifenable", ccmctidobjects.Ccmctidrolledovernotifenable}
    return &(ccmctidobjects.EntityData)
}

// CISCOCONFIGMANMIB_Ccmhistoryeventtable
// A table of configuration events on this router.
type CISCOCONFIGMANMIB_Ccmhistoryeventtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a configuration event on this router. The type is slice
    // of CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry.
    Ccmhistoryevententry []CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry
}

func (ccmhistoryeventtable *CISCOCONFIGMANMIB_Ccmhistoryeventtable) GetEntityData() *types.CommonEntityData {
    ccmhistoryeventtable.EntityData.YFilter = ccmhistoryeventtable.YFilter
    ccmhistoryeventtable.EntityData.YangName = "ccmHistoryEventTable"
    ccmhistoryeventtable.EntityData.BundleName = "cisco_ios_xe"
    ccmhistoryeventtable.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmhistoryeventtable.EntityData.SegmentPath = "ccmHistoryEventTable"
    ccmhistoryeventtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmhistoryeventtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmhistoryeventtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmhistoryeventtable.EntityData.Children = make(map[string]types.YChild)
    ccmhistoryeventtable.EntityData.Children["ccmHistoryEventEntry"] = types.YChild{"Ccmhistoryevententry", nil}
    for i := range ccmhistoryeventtable.Ccmhistoryevententry {
        ccmhistoryeventtable.EntityData.Children[types.GetSegmentPath(&ccmhistoryeventtable.Ccmhistoryevententry[i])] = types.YChild{"Ccmhistoryevententry", &ccmhistoryeventtable.Ccmhistoryevententry[i]}
    }
    ccmhistoryeventtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccmhistoryeventtable.EntityData)
}

// CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry
// Information about a configuration event on this
// router.
type CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ccmhistoryevententry *CISCOCONFIGMANMIB_Ccmhistoryeventtable_Ccmhistoryevententry) GetEntityData() *types.CommonEntityData {
    ccmhistoryevententry.EntityData.YFilter = ccmhistoryevententry.YFilter
    ccmhistoryevententry.EntityData.YangName = "ccmHistoryEventEntry"
    ccmhistoryevententry.EntityData.BundleName = "cisco_ios_xe"
    ccmhistoryevententry.EntityData.ParentYangName = "ccmHistoryEventTable"
    ccmhistoryevententry.EntityData.SegmentPath = "ccmHistoryEventEntry" + "[ccmHistoryEventIndex='" + fmt.Sprintf("%v", ccmhistoryevententry.Ccmhistoryeventindex) + "']"
    ccmhistoryevententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmhistoryevententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmhistoryevententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmhistoryevententry.EntityData.Children = make(map[string]types.YChild)
    ccmhistoryevententry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventIndex"] = types.YLeaf{"Ccmhistoryeventindex", ccmhistoryevententry.Ccmhistoryeventindex}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventTime"] = types.YLeaf{"Ccmhistoryeventtime", ccmhistoryevententry.Ccmhistoryeventtime}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventCommandSource"] = types.YLeaf{"Ccmhistoryeventcommandsource", ccmhistoryevententry.Ccmhistoryeventcommandsource}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventConfigSource"] = types.YLeaf{"Ccmhistoryeventconfigsource", ccmhistoryevententry.Ccmhistoryeventconfigsource}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventConfigDestination"] = types.YLeaf{"Ccmhistoryeventconfigdestination", ccmhistoryevententry.Ccmhistoryeventconfigdestination}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventTerminalType"] = types.YLeaf{"Ccmhistoryeventterminaltype", ccmhistoryevententry.Ccmhistoryeventterminaltype}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventTerminalNumber"] = types.YLeaf{"Ccmhistoryeventterminalnumber", ccmhistoryevententry.Ccmhistoryeventterminalnumber}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventTerminalUser"] = types.YLeaf{"Ccmhistoryeventterminaluser", ccmhistoryevententry.Ccmhistoryeventterminaluser}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventTerminalLocation"] = types.YLeaf{"Ccmhistoryeventterminallocation", ccmhistoryevententry.Ccmhistoryeventterminallocation}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventCommandSourceAddress"] = types.YLeaf{"Ccmhistoryeventcommandsourceaddress", ccmhistoryevententry.Ccmhistoryeventcommandsourceaddress}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventVirtualHostName"] = types.YLeaf{"Ccmhistoryeventvirtualhostname", ccmhistoryevententry.Ccmhistoryeventvirtualhostname}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventServerAddress"] = types.YLeaf{"Ccmhistoryeventserveraddress", ccmhistoryevententry.Ccmhistoryeventserveraddress}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventFile"] = types.YLeaf{"Ccmhistoryeventfile", ccmhistoryevententry.Ccmhistoryeventfile}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventRcpUser"] = types.YLeaf{"Ccmhistoryeventrcpuser", ccmhistoryevententry.Ccmhistoryeventrcpuser}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryCLICmdEntriesBumped"] = types.YLeaf{"Ccmhistoryclicmdentriesbumped", ccmhistoryevententry.Ccmhistoryclicmdentriesbumped}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventCommandSourceAddrType"] = types.YLeaf{"Ccmhistoryeventcommandsourceaddrtype", ccmhistoryevententry.Ccmhistoryeventcommandsourceaddrtype}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventCommandSourceAddrRev1"] = types.YLeaf{"Ccmhistoryeventcommandsourceaddrrev1", ccmhistoryevententry.Ccmhistoryeventcommandsourceaddrrev1}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventServerAddrType"] = types.YLeaf{"Ccmhistoryeventserveraddrtype", ccmhistoryevententry.Ccmhistoryeventserveraddrtype}
    ccmhistoryevententry.EntityData.Leafs["ccmHistoryEventServerAddrRev1"] = types.YLeaf{"Ccmhistoryeventserveraddrrev1", ccmhistoryevententry.Ccmhistoryeventserveraddrrev1}
    return &(ccmhistoryevententry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (ccmclihistorycommandtable *CISCOCONFIGMANMIB_Ccmclihistorycommandtable) GetEntityData() *types.CommonEntityData {
    ccmclihistorycommandtable.EntityData.YFilter = ccmclihistorycommandtable.YFilter
    ccmclihistorycommandtable.EntityData.YangName = "ccmCLIHistoryCommandTable"
    ccmclihistorycommandtable.EntityData.BundleName = "cisco_ios_xe"
    ccmclihistorycommandtable.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmclihistorycommandtable.EntityData.SegmentPath = "ccmCLIHistoryCommandTable"
    ccmclihistorycommandtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmclihistorycommandtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmclihistorycommandtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmclihistorycommandtable.EntityData.Children = make(map[string]types.YChild)
    ccmclihistorycommandtable.EntityData.Children["ccmCLIHistoryCommandEntry"] = types.YChild{"Ccmclihistorycommandentry", nil}
    for i := range ccmclihistorycommandtable.Ccmclihistorycommandentry {
        ccmclihistorycommandtable.EntityData.Children[types.GetSegmentPath(&ccmclihistorycommandtable.Ccmclihistorycommandentry[i])] = types.YChild{"Ccmclihistorycommandentry", &ccmclihistorycommandtable.Ccmclihistorycommandentry[i]}
    }
    ccmclihistorycommandtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccmclihistorycommandtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (ccmclihistorycommandentry *CISCOCONFIGMANMIB_Ccmclihistorycommandtable_Ccmclihistorycommandentry) GetEntityData() *types.CommonEntityData {
    ccmclihistorycommandentry.EntityData.YFilter = ccmclihistorycommandentry.YFilter
    ccmclihistorycommandentry.EntityData.YangName = "ccmCLIHistoryCommandEntry"
    ccmclihistorycommandentry.EntityData.BundleName = "cisco_ios_xe"
    ccmclihistorycommandentry.EntityData.ParentYangName = "ccmCLIHistoryCommandTable"
    ccmclihistorycommandentry.EntityData.SegmentPath = "ccmCLIHistoryCommandEntry" + "[ccmHistoryEventIndex='" + fmt.Sprintf("%v", ccmclihistorycommandentry.Ccmhistoryeventindex) + "']" + "[ccmCLIHistoryCommandIndex='" + fmt.Sprintf("%v", ccmclihistorycommandentry.Ccmclihistorycommandindex) + "']"
    ccmclihistorycommandentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmclihistorycommandentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmclihistorycommandentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmclihistorycommandentry.EntityData.Children = make(map[string]types.YChild)
    ccmclihistorycommandentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccmclihistorycommandentry.EntityData.Leafs["ccmHistoryEventIndex"] = types.YLeaf{"Ccmhistoryeventindex", ccmclihistorycommandentry.Ccmhistoryeventindex}
    ccmclihistorycommandentry.EntityData.Leafs["ccmCLIHistoryCommandIndex"] = types.YLeaf{"Ccmclihistorycommandindex", ccmclihistorycommandentry.Ccmclihistorycommandindex}
    ccmclihistorycommandentry.EntityData.Leafs["ccmCLIHistoryCommand"] = types.YLeaf{"Ccmclihistorycommand", ccmclihistorycommandentry.Ccmclihistorycommand}
    return &(ccmclihistorycommandentry.EntityData)
}

