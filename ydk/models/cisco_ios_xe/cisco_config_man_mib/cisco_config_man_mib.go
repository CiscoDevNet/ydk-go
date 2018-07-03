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

    
    CcmHistory CISCOCONFIGMANMIB_CcmHistory

    
    CcmCLIHistory CISCOCONFIGMANMIB_CcmCLIHistory

    
    CcmCLICfg CISCOCONFIGMANMIB_CcmCLICfg

    
    CcmCTIDObjects CISCOCONFIGMANMIB_CcmCTIDObjects

    // A table of configuration events on this router.
    CcmHistoryEventTable CISCOCONFIGMANMIB_CcmHistoryEventTable

    // A table of CLI commands that took effect during configuration events.
    CcmCLIHistoryCommandTable CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable
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

    cISCOCONFIGMANMIB.EntityData.Children = types.NewOrderedMap()
    cISCOCONFIGMANMIB.EntityData.Children.Append("ccmHistory", types.YChild{"CcmHistory", &cISCOCONFIGMANMIB.CcmHistory})
    cISCOCONFIGMANMIB.EntityData.Children.Append("ccmCLIHistory", types.YChild{"CcmCLIHistory", &cISCOCONFIGMANMIB.CcmCLIHistory})
    cISCOCONFIGMANMIB.EntityData.Children.Append("ccmCLICfg", types.YChild{"CcmCLICfg", &cISCOCONFIGMANMIB.CcmCLICfg})
    cISCOCONFIGMANMIB.EntityData.Children.Append("ccmCTIDObjects", types.YChild{"CcmCTIDObjects", &cISCOCONFIGMANMIB.CcmCTIDObjects})
    cISCOCONFIGMANMIB.EntityData.Children.Append("ccmHistoryEventTable", types.YChild{"CcmHistoryEventTable", &cISCOCONFIGMANMIB.CcmHistoryEventTable})
    cISCOCONFIGMANMIB.EntityData.Children.Append("ccmCLIHistoryCommandTable", types.YChild{"CcmCLIHistoryCommandTable", &cISCOCONFIGMANMIB.CcmCLIHistoryCommandTable})
    cISCOCONFIGMANMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOCONFIGMANMIB.EntityData.YListKeys = []string {}

    return &(cISCOCONFIGMANMIB.EntityData)
}

// CISCOCONFIGMANMIB_CcmHistory
type CISCOCONFIGMANMIB_CcmHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime when the running configuration was last changed.    
    // If the value of ccmHistoryRunningLastChanged is         greater than
    // ccmHistoryRunningLastSaved, the          configuration has been changed but
    // not saved. The type is interface{} with range: 0..4294967295.
    CcmHistoryRunningLastChanged interface{}

    // The value of sysUpTime when the running configuration was last saved
    // (written).  If the value of ccmHistoryRunningLastChanged is  greater than
    // ccmHistoryRunningLastSaved, the  configuration has been changed but not
    // saved.  What constitutes a safe saving of the running configuration is a
    // management policy issue beyond the scope of this MIB.  For some
    // installations, writing the running configuration to a terminal may be a way
    // of capturing and saving it.  Others may use local or remote storage.  Thus
    // ANY write is considered saving for the purposes of the MIB. The type is
    // interface{} with range: 0..4294967295.
    CcmHistoryRunningLastSaved interface{}

    // The value of sysUpTime when the startup configuration was last written to. 
    // In general this is the default configuration used when cold starting the
    // system.  It may have been changed by a save of the running configuration or
    // by a copy from elsewhere. The type is interface{} with range:
    // 0..4294967295.
    CcmHistoryStartupLastChanged interface{}

    // The maximum number of entries that can be held in ccmHistoryEventTable. 
    // The recommended value for implementations is 10. The type is interface{}
    // with range: 0..2147483647.
    CcmHistoryMaxEventEntries interface{}

    // The number of times the oldest entry in ccmHistoryEventTable was deleted to
    // make room  for a new entry. The type is interface{} with range:
    // 0..4294967295.
    CcmHistoryEventEntriesBumped interface{}
}

func (ccmHistory *CISCOCONFIGMANMIB_CcmHistory) GetEntityData() *types.CommonEntityData {
    ccmHistory.EntityData.YFilter = ccmHistory.YFilter
    ccmHistory.EntityData.YangName = "ccmHistory"
    ccmHistory.EntityData.BundleName = "cisco_ios_xe"
    ccmHistory.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmHistory.EntityData.SegmentPath = "ccmHistory"
    ccmHistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmHistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmHistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmHistory.EntityData.Children = types.NewOrderedMap()
    ccmHistory.EntityData.Leafs = types.NewOrderedMap()
    ccmHistory.EntityData.Leafs.Append("ccmHistoryRunningLastChanged", types.YLeaf{"CcmHistoryRunningLastChanged", ccmHistory.CcmHistoryRunningLastChanged})
    ccmHistory.EntityData.Leafs.Append("ccmHistoryRunningLastSaved", types.YLeaf{"CcmHistoryRunningLastSaved", ccmHistory.CcmHistoryRunningLastSaved})
    ccmHistory.EntityData.Leafs.Append("ccmHistoryStartupLastChanged", types.YLeaf{"CcmHistoryStartupLastChanged", ccmHistory.CcmHistoryStartupLastChanged})
    ccmHistory.EntityData.Leafs.Append("ccmHistoryMaxEventEntries", types.YLeaf{"CcmHistoryMaxEventEntries", ccmHistory.CcmHistoryMaxEventEntries})
    ccmHistory.EntityData.Leafs.Append("ccmHistoryEventEntriesBumped", types.YLeaf{"CcmHistoryEventEntriesBumped", ccmHistory.CcmHistoryEventEntriesBumped})

    ccmHistory.EntityData.YListKeys = []string {}

    return &(ccmHistory.EntityData)
}

// CISCOCONFIGMANMIB_CcmCLIHistory
type CISCOCONFIGMANMIB_CcmCLIHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of entries that can be held in
    // ccmCLIHistoryCommandTable.  The recommended value for implementations is
    // 100.  If the number of entries in ccmCLIHistoryCommandTable  exceeds the
    // value of this object, old entries will be  bumped to make room for new
    // entries.  The ccmCLIHistoryCommandTable will not be populated if the value
    // of this object is 0. The type is interface{} with range: 0..4294967295.
    CcmCLIHistoryMaxCmdEntries interface{}

    // The current number of entries in ccmCLIHistoryCommandTable. The type is
    // interface{} with range: 0..4294967295.
    CcmCLIHistoryCmdEntries interface{}

    // This object indicates the upper limit on the number of entries allowed in 
    // ccmCLIHistoryCommandTable by the managed system. The type is interface{}
    // with range: 0..4294967295.
    CcmCLIHistoryCmdEntriesAllowed interface{}
}

func (ccmCLIHistory *CISCOCONFIGMANMIB_CcmCLIHistory) GetEntityData() *types.CommonEntityData {
    ccmCLIHistory.EntityData.YFilter = ccmCLIHistory.YFilter
    ccmCLIHistory.EntityData.YangName = "ccmCLIHistory"
    ccmCLIHistory.EntityData.BundleName = "cisco_ios_xe"
    ccmCLIHistory.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmCLIHistory.EntityData.SegmentPath = "ccmCLIHistory"
    ccmCLIHistory.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmCLIHistory.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmCLIHistory.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmCLIHistory.EntityData.Children = types.NewOrderedMap()
    ccmCLIHistory.EntityData.Leafs = types.NewOrderedMap()
    ccmCLIHistory.EntityData.Leafs.Append("ccmCLIHistoryMaxCmdEntries", types.YLeaf{"CcmCLIHistoryMaxCmdEntries", ccmCLIHistory.CcmCLIHistoryMaxCmdEntries})
    ccmCLIHistory.EntityData.Leafs.Append("ccmCLIHistoryCmdEntries", types.YLeaf{"CcmCLIHistoryCmdEntries", ccmCLIHistory.CcmCLIHistoryCmdEntries})
    ccmCLIHistory.EntityData.Leafs.Append("ccmCLIHistoryCmdEntriesAllowed", types.YLeaf{"CcmCLIHistoryCmdEntriesAllowed", ccmCLIHistory.CcmCLIHistoryCmdEntriesAllowed})

    ccmCLIHistory.EntityData.YListKeys = []string {}

    return &(ccmCLIHistory.EntityData)
}

// CISCOCONFIGMANMIB_CcmCLICfg
type CISCOCONFIGMANMIB_CcmCLICfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable indicates whether the system produces the
    // ccmCLIRunningConfigChanged notification. A false  value will prevent
    // notifications from being generated  by this system. The type is bool.
    CcmCLICfgRunConfNotifEnable interface{}
}

func (ccmCLICfg *CISCOCONFIGMANMIB_CcmCLICfg) GetEntityData() *types.CommonEntityData {
    ccmCLICfg.EntityData.YFilter = ccmCLICfg.YFilter
    ccmCLICfg.EntityData.YangName = "ccmCLICfg"
    ccmCLICfg.EntityData.BundleName = "cisco_ios_xe"
    ccmCLICfg.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmCLICfg.EntityData.SegmentPath = "ccmCLICfg"
    ccmCLICfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmCLICfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmCLICfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmCLICfg.EntityData.Children = types.NewOrderedMap()
    ccmCLICfg.EntityData.Leafs = types.NewOrderedMap()
    ccmCLICfg.EntityData.Leafs.Append("ccmCLICfgRunConfNotifEnable", types.YLeaf{"CcmCLICfgRunConfNotifEnable", ccmCLICfg.CcmCLICfgRunConfNotifEnable})

    ccmCLICfg.EntityData.YListKeys = []string {}

    return &(ccmCLICfg.EntityData)
}

// CISCOCONFIGMANMIB_CcmCTIDObjects
type CISCOCONFIGMANMIB_CcmCTIDObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the Config Change Tracking ID which uniquely
    // represents version-incrementing changes to the IOS  running configuration.
    // The type is interface{} with range: 0..18446744073709551615.
    CcmCTID interface{}

    // This object indicates the time when the Config Change Tracking ID last
    // changed. The type is string.
    CcmCTIDLastChangeTime interface{}

    // This object indicates the user who last reset the Config Change Tracking
    // ID. The type is string.
    CcmCTIDWhoChanged interface{}

    // This variable indicates whether the system produces the ccmCTIDRolledOver
    // notification. A false value will prevent notifications from being generated
    // by this system. The type is bool.
    CcmCTIDRolledOverNotifEnable interface{}
}

func (ccmCTIDObjects *CISCOCONFIGMANMIB_CcmCTIDObjects) GetEntityData() *types.CommonEntityData {
    ccmCTIDObjects.EntityData.YFilter = ccmCTIDObjects.YFilter
    ccmCTIDObjects.EntityData.YangName = "ccmCTIDObjects"
    ccmCTIDObjects.EntityData.BundleName = "cisco_ios_xe"
    ccmCTIDObjects.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmCTIDObjects.EntityData.SegmentPath = "ccmCTIDObjects"
    ccmCTIDObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmCTIDObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmCTIDObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmCTIDObjects.EntityData.Children = types.NewOrderedMap()
    ccmCTIDObjects.EntityData.Leafs = types.NewOrderedMap()
    ccmCTIDObjects.EntityData.Leafs.Append("ccmCTID", types.YLeaf{"CcmCTID", ccmCTIDObjects.CcmCTID})
    ccmCTIDObjects.EntityData.Leafs.Append("ccmCTIDLastChangeTime", types.YLeaf{"CcmCTIDLastChangeTime", ccmCTIDObjects.CcmCTIDLastChangeTime})
    ccmCTIDObjects.EntityData.Leafs.Append("ccmCTIDWhoChanged", types.YLeaf{"CcmCTIDWhoChanged", ccmCTIDObjects.CcmCTIDWhoChanged})
    ccmCTIDObjects.EntityData.Leafs.Append("ccmCTIDRolledOverNotifEnable", types.YLeaf{"CcmCTIDRolledOverNotifEnable", ccmCTIDObjects.CcmCTIDRolledOverNotifEnable})

    ccmCTIDObjects.EntityData.YListKeys = []string {}

    return &(ccmCTIDObjects.EntityData)
}

// CISCOCONFIGMANMIB_CcmHistoryEventTable
// A table of configuration events on this router.
type CISCOCONFIGMANMIB_CcmHistoryEventTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a configuration event on this router. The type is slice
    // of CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry.
    CcmHistoryEventEntry []*CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry
}

func (ccmHistoryEventTable *CISCOCONFIGMANMIB_CcmHistoryEventTable) GetEntityData() *types.CommonEntityData {
    ccmHistoryEventTable.EntityData.YFilter = ccmHistoryEventTable.YFilter
    ccmHistoryEventTable.EntityData.YangName = "ccmHistoryEventTable"
    ccmHistoryEventTable.EntityData.BundleName = "cisco_ios_xe"
    ccmHistoryEventTable.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmHistoryEventTable.EntityData.SegmentPath = "ccmHistoryEventTable"
    ccmHistoryEventTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmHistoryEventTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmHistoryEventTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmHistoryEventTable.EntityData.Children = types.NewOrderedMap()
    ccmHistoryEventTable.EntityData.Children.Append("ccmHistoryEventEntry", types.YChild{"CcmHistoryEventEntry", nil})
    for i := range ccmHistoryEventTable.CcmHistoryEventEntry {
        ccmHistoryEventTable.EntityData.Children.Append(types.GetSegmentPath(ccmHistoryEventTable.CcmHistoryEventEntry[i]), types.YChild{"CcmHistoryEventEntry", ccmHistoryEventTable.CcmHistoryEventEntry[i]})
    }
    ccmHistoryEventTable.EntityData.Leafs = types.NewOrderedMap()

    ccmHistoryEventTable.EntityData.YListKeys = []string {}

    return &(ccmHistoryEventTable.EntityData)
}

// CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry
// Information about a configuration event on this
// router.
type CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A monotonically increasing integer for the sole
    // purpose of indexing events.  When it reaches the  maximum value, an
    // extremely unlikely event, the agent  wraps the value back to 1 and may
    // flush existing  entries. The type is interface{} with range: 1..2147483647.
    CcmHistoryEventIndex interface{}

    // The value of sysUpTime when the event occurred. The type is interface{}
    // with range: 0..4294967295.
    CcmHistoryEventTime interface{}

    // The source of the command that instigated the event. The type is
    // CcmHistoryEventCommandSource.
    CcmHistoryEventCommandSource interface{}

    // The configuration data source for the event. The type is
    // HistoryEventMedium.
    CcmHistoryEventConfigSource interface{}

    // The configuration data destination for the event. The type is
    // HistoryEventMedium.
    CcmHistoryEventConfigDestination interface{}

    // If ccmHistoryEventCommandSource is 'commandLine', the terminal type,
    // otherwise 'notApplicable'. The type is CcmHistoryEventTerminalType.
    CcmHistoryEventTerminalType interface{}

    // If ccmHistoryEventCommandSource is 'commandLine', the terminal number.  The
    // value is -1 if not available or not applicable. The type is interface{}
    // with range: -2147483648..2147483647.
    CcmHistoryEventTerminalNumber interface{}

    // If ccmHistoryEventCommandSource is 'commandLine', the name of the logged in
    // user.  The length is zero if not available or not applicable. The type is
    // string with length: 0..64.
    CcmHistoryEventTerminalUser interface{}

    // If ccmHistoryEventCommandSource is 'commandLine', the hard-wired location
    // of the terminal or the remote  host for an incoming connection.  The length
    // is zero  if not available or not applicable. The type is string with
    // length: 0..64.
    CcmHistoryEventTerminalLocation interface{}

    // If ccmHistoryEventTerminalType is 'virtual', the internet address of the
    // connected system.  If ccmHistoryEventCommandSource is 'snmp', the internet
    // address of the requester.  The value is 0.0.0.0 if not available or not 
    // applicable.  This object is deprecated by
    // ccmHistoryEventCommandSourceAddrRev1. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CcmHistoryEventCommandSourceAddress interface{}

    // If ccmHistoryEventTerminalType is 'virtual', the host name of the connected
    // system.  The length is zero if not available or not applicable. The type is
    // string with length: 0..64.
    CcmHistoryEventVirtualHostName interface{}

    // If ccmHistoryEventConfigSource or ccmHistoryEventConfigDestination is
    // 'networkTftp' or 'networkRcp', the internet address of the storage file
    // server.  The value is 0.0.0.0 if not applicable or not         available.  
    // This object is deprecated by         ccmHistoryEventServerAddrRev1. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CcmHistoryEventServerAddress interface{}

    // If ccmHistoryEventConfigSource or ccmHistoryEventConfigDestination is
    // 'networkTftp' or 'networkRcp', the configuration file name at the storage
    // file server.  The length is zero if not available or not applicable. The
    // type is string with length: 0..64.
    CcmHistoryEventFile interface{}

    // If ccmHistoryEventConfigSource or ccmHistoryEventConfigDestination is
    // 'networkRcp', the remote user name.  The length is zero if not applicable
    // or not available. The type is string with length: 0..64.
    CcmHistoryEventRcpUser interface{}

    // The number of times the oldest entry in ccmCLIHistoryCommandTable with
    // first index as  ccmHistoryEventIndex was deleted to make  room for a new
    // entry.  This object is applicable only if  ccmHistoryEventCommandSource has
    // a value  of 'commandLine'. The type is interface{} with range:
    // 0..4294967295.
    CcmHistoryCLICmdEntriesBumped interface{}

    // This object indicates the transport type of the address contained in
    // ccmHistoryEventCommandSourceAddrRev1.  The value will be zero if not
    // available or not applicable. The type is InetAddressType.
    CcmHistoryEventCommandSourceAddrType interface{}

    // If ccmHistoryEventTerminalType is 'virtual', the internet address of the
    // connected system.  If ccmHistoryEventCommandSource is 'snmp', the internet
    // address of the requester.  The value of all bit's is zero  if not available
    // or not applicable.  The Format of this address depends on the value of the
    // ccmHistoryEventCommandSourceAddrType object.  This object deprecates
    // ccmHistoryEventCommandSourceAddress. The type is string with length:
    // 0..255.
    CcmHistoryEventCommandSourceAddrRev1 interface{}

    // This object indicates the transport type of the address contained in
    // ccmHistoryEventServerAddrRev1.  The value will be zero if not available or
    // not aplicable. The type is InetAddressType.
    CcmHistoryEventServerAddrType interface{}

    // If ccmHistoryEventConfigSource or ccmHistoryEventConfigDestination is
    // 'networkTftp' or 'networkRcp', the internet address of the storage file
    // server.   The value of all bits is 0s if not applicable or not available. 
    // The Format of this address depends on the value of the
    // ccmHistoryEventServerAddrType object.  This object deprecates
    // ccmHistoryEventServerAddress. The type is string with length: 0..255.
    CcmHistoryEventServerAddrRev1 interface{}
}

func (ccmHistoryEventEntry *CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry) GetEntityData() *types.CommonEntityData {
    ccmHistoryEventEntry.EntityData.YFilter = ccmHistoryEventEntry.YFilter
    ccmHistoryEventEntry.EntityData.YangName = "ccmHistoryEventEntry"
    ccmHistoryEventEntry.EntityData.BundleName = "cisco_ios_xe"
    ccmHistoryEventEntry.EntityData.ParentYangName = "ccmHistoryEventTable"
    ccmHistoryEventEntry.EntityData.SegmentPath = "ccmHistoryEventEntry" + types.AddKeyToken(ccmHistoryEventEntry.CcmHistoryEventIndex, "ccmHistoryEventIndex")
    ccmHistoryEventEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmHistoryEventEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmHistoryEventEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmHistoryEventEntry.EntityData.Children = types.NewOrderedMap()
    ccmHistoryEventEntry.EntityData.Leafs = types.NewOrderedMap()
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventIndex", types.YLeaf{"CcmHistoryEventIndex", ccmHistoryEventEntry.CcmHistoryEventIndex})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventTime", types.YLeaf{"CcmHistoryEventTime", ccmHistoryEventEntry.CcmHistoryEventTime})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventCommandSource", types.YLeaf{"CcmHistoryEventCommandSource", ccmHistoryEventEntry.CcmHistoryEventCommandSource})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventConfigSource", types.YLeaf{"CcmHistoryEventConfigSource", ccmHistoryEventEntry.CcmHistoryEventConfigSource})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventConfigDestination", types.YLeaf{"CcmHistoryEventConfigDestination", ccmHistoryEventEntry.CcmHistoryEventConfigDestination})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventTerminalType", types.YLeaf{"CcmHistoryEventTerminalType", ccmHistoryEventEntry.CcmHistoryEventTerminalType})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventTerminalNumber", types.YLeaf{"CcmHistoryEventTerminalNumber", ccmHistoryEventEntry.CcmHistoryEventTerminalNumber})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventTerminalUser", types.YLeaf{"CcmHistoryEventTerminalUser", ccmHistoryEventEntry.CcmHistoryEventTerminalUser})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventTerminalLocation", types.YLeaf{"CcmHistoryEventTerminalLocation", ccmHistoryEventEntry.CcmHistoryEventTerminalLocation})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventCommandSourceAddress", types.YLeaf{"CcmHistoryEventCommandSourceAddress", ccmHistoryEventEntry.CcmHistoryEventCommandSourceAddress})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventVirtualHostName", types.YLeaf{"CcmHistoryEventVirtualHostName", ccmHistoryEventEntry.CcmHistoryEventVirtualHostName})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventServerAddress", types.YLeaf{"CcmHistoryEventServerAddress", ccmHistoryEventEntry.CcmHistoryEventServerAddress})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventFile", types.YLeaf{"CcmHistoryEventFile", ccmHistoryEventEntry.CcmHistoryEventFile})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventRcpUser", types.YLeaf{"CcmHistoryEventRcpUser", ccmHistoryEventEntry.CcmHistoryEventRcpUser})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryCLICmdEntriesBumped", types.YLeaf{"CcmHistoryCLICmdEntriesBumped", ccmHistoryEventEntry.CcmHistoryCLICmdEntriesBumped})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventCommandSourceAddrType", types.YLeaf{"CcmHistoryEventCommandSourceAddrType", ccmHistoryEventEntry.CcmHistoryEventCommandSourceAddrType})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventCommandSourceAddrRev1", types.YLeaf{"CcmHistoryEventCommandSourceAddrRev1", ccmHistoryEventEntry.CcmHistoryEventCommandSourceAddrRev1})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventServerAddrType", types.YLeaf{"CcmHistoryEventServerAddrType", ccmHistoryEventEntry.CcmHistoryEventServerAddrType})
    ccmHistoryEventEntry.EntityData.Leafs.Append("ccmHistoryEventServerAddrRev1", types.YLeaf{"CcmHistoryEventServerAddrRev1", ccmHistoryEventEntry.CcmHistoryEventServerAddrRev1})

    ccmHistoryEventEntry.EntityData.YListKeys = []string {"CcmHistoryEventIndex"}

    return &(ccmHistoryEventEntry.EntityData)
}

// CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventCommandSource represents The source of the command that instigated the event.
type CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventCommandSource string

const (
    CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventCommandSource_commandLine CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventCommandSource = "commandLine"

    CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventCommandSource_snmp CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventCommandSource = "snmp"
)

// CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType represents the terminal type, otherwise 'notApplicable'.
type CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType string

const (
    CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType_notApplicable CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType = "notApplicable"

    CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType_unknown CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType = "unknown"

    CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType_console CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType = "console"

    CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType_terminal CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType = "terminal"

    CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType_virtual CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType = "virtual"

    CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType_auxiliary CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventTerminalType = "auxiliary"
)

// CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable
// A table of CLI commands that took effect during
// configuration events.
type CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about the CLI commands that took effect during the
    // configuration event pointed by  ccmCLIHistoryEventIndex.  A set of rows in
    // this table having the first index as ccmHistoryEventIndex will store the
    // CLI commands entered during the corresponding  configuration event in
    // ccmHistoryEventTable.  An entry will be created in this table only if  the
    // corresponding entry in ccmHistoryEventTable has  a value of 'commandLine'
    // for  ccmHistoryEventCommandSource. The type is slice of
    // CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable_CcmCLIHistoryCommandEntry.
    CcmCLIHistoryCommandEntry []*CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable_CcmCLIHistoryCommandEntry
}

func (ccmCLIHistoryCommandTable *CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable) GetEntityData() *types.CommonEntityData {
    ccmCLIHistoryCommandTable.EntityData.YFilter = ccmCLIHistoryCommandTable.YFilter
    ccmCLIHistoryCommandTable.EntityData.YangName = "ccmCLIHistoryCommandTable"
    ccmCLIHistoryCommandTable.EntityData.BundleName = "cisco_ios_xe"
    ccmCLIHistoryCommandTable.EntityData.ParentYangName = "CISCO-CONFIG-MAN-MIB"
    ccmCLIHistoryCommandTable.EntityData.SegmentPath = "ccmCLIHistoryCommandTable"
    ccmCLIHistoryCommandTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmCLIHistoryCommandTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmCLIHistoryCommandTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmCLIHistoryCommandTable.EntityData.Children = types.NewOrderedMap()
    ccmCLIHistoryCommandTable.EntityData.Children.Append("ccmCLIHistoryCommandEntry", types.YChild{"CcmCLIHistoryCommandEntry", nil})
    for i := range ccmCLIHistoryCommandTable.CcmCLIHistoryCommandEntry {
        ccmCLIHistoryCommandTable.EntityData.Children.Append(types.GetSegmentPath(ccmCLIHistoryCommandTable.CcmCLIHistoryCommandEntry[i]), types.YChild{"CcmCLIHistoryCommandEntry", ccmCLIHistoryCommandTable.CcmCLIHistoryCommandEntry[i]})
    }
    ccmCLIHistoryCommandTable.EntityData.Leafs = types.NewOrderedMap()

    ccmCLIHistoryCommandTable.EntityData.YListKeys = []string {}

    return &(ccmCLIHistoryCommandTable.EntityData)
}

// CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable_CcmCLIHistoryCommandEntry
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
type CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable_CcmCLIHistoryCommandEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_config_man_mib.CISCOCONFIGMANMIB_CcmHistoryEventTable_CcmHistoryEventEntry_CcmHistoryEventIndex
    CcmHistoryEventIndex interface{}

    // This attribute is a key. A monotonically increasing integer for the purpose
    // of indexing CLI commands which took effect during a configuration event.
    // The type is interface{} with range: 1..4294967295.
    CcmCLIHistoryCommandIndex interface{}

    // The CLI command entered which took effect during the configuration event
    // pointed by  ccmHistoryEventIndex. The type is string.
    CcmCLIHistoryCommand interface{}
}

func (ccmCLIHistoryCommandEntry *CISCOCONFIGMANMIB_CcmCLIHistoryCommandTable_CcmCLIHistoryCommandEntry) GetEntityData() *types.CommonEntityData {
    ccmCLIHistoryCommandEntry.EntityData.YFilter = ccmCLIHistoryCommandEntry.YFilter
    ccmCLIHistoryCommandEntry.EntityData.YangName = "ccmCLIHistoryCommandEntry"
    ccmCLIHistoryCommandEntry.EntityData.BundleName = "cisco_ios_xe"
    ccmCLIHistoryCommandEntry.EntityData.ParentYangName = "ccmCLIHistoryCommandTable"
    ccmCLIHistoryCommandEntry.EntityData.SegmentPath = "ccmCLIHistoryCommandEntry" + types.AddKeyToken(ccmCLIHistoryCommandEntry.CcmHistoryEventIndex, "ccmHistoryEventIndex") + types.AddKeyToken(ccmCLIHistoryCommandEntry.CcmCLIHistoryCommandIndex, "ccmCLIHistoryCommandIndex")
    ccmCLIHistoryCommandEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccmCLIHistoryCommandEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccmCLIHistoryCommandEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccmCLIHistoryCommandEntry.EntityData.Children = types.NewOrderedMap()
    ccmCLIHistoryCommandEntry.EntityData.Leafs = types.NewOrderedMap()
    ccmCLIHistoryCommandEntry.EntityData.Leafs.Append("ccmHistoryEventIndex", types.YLeaf{"CcmHistoryEventIndex", ccmCLIHistoryCommandEntry.CcmHistoryEventIndex})
    ccmCLIHistoryCommandEntry.EntityData.Leafs.Append("ccmCLIHistoryCommandIndex", types.YLeaf{"CcmCLIHistoryCommandIndex", ccmCLIHistoryCommandEntry.CcmCLIHistoryCommandIndex})
    ccmCLIHistoryCommandEntry.EntityData.Leafs.Append("ccmCLIHistoryCommand", types.YLeaf{"CcmCLIHistoryCommand", ccmCLIHistoryCommandEntry.CcmCLIHistoryCommand})

    ccmCLIHistoryCommandEntry.EntityData.YListKeys = []string {"CcmHistoryEventIndex", "CcmCLIHistoryCommandIndex"}

    return &(ccmCLIHistoryCommandEntry.EntityData)
}

