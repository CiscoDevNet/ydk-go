// The MIB module for logging SNMP Notifications, that is, Traps
// and Informs.
package notification_log_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package notification_log_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:NOTIFICATION-LOG-MIB NOTIFICATION-LOG-MIB}", reflect.TypeOf(NOTIFICATIONLOGMIB{}))
    ydk.RegisterEntity("NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB", reflect.TypeOf(NOTIFICATIONLOGMIB{}))
}

// NOTIFICATIONLOGMIB
type NOTIFICATIONLOGMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    NlmConfig NOTIFICATIONLOGMIB_NlmConfig

    
    NlmStats NOTIFICATIONLOGMIB_NlmStats

    // A table of logging control entries.
    NlmConfigLogTable NOTIFICATIONLOGMIB_NlmConfigLogTable

    // A table of Notification log entries.  It is an implementation-specific
    // matter whether entries in this table are preserved across initializations
    // of the management system.  In general one would expect that they are not. 
    // Note that keeping entries across initializations of the management system
    // leads to some confusion with counters and TimeStamps, since both of those
    // are based on sysUpTime, which resets on management initialization.  In this
    // situation, counters apply only after the reset and nlmLogTime for entries
    // made before the reset MUST be set to 0.
    NlmLogTable NOTIFICATIONLOGMIB_NlmLogTable

    // A table of variables to go with Notification log entries.
    NlmLogVariableTable NOTIFICATIONLOGMIB_NlmLogVariableTable
}

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetEntityData() *types.CommonEntityData {
    nOTIFICATIONLOGMIB.EntityData.YFilter = nOTIFICATIONLOGMIB.YFilter
    nOTIFICATIONLOGMIB.EntityData.YangName = "NOTIFICATION-LOG-MIB"
    nOTIFICATIONLOGMIB.EntityData.BundleName = "cisco_ios_xe"
    nOTIFICATIONLOGMIB.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nOTIFICATIONLOGMIB.EntityData.SegmentPath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB"
    nOTIFICATIONLOGMIB.EntityData.AbsolutePath = nOTIFICATIONLOGMIB.EntityData.SegmentPath
    nOTIFICATIONLOGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nOTIFICATIONLOGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nOTIFICATIONLOGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nOTIFICATIONLOGMIB.EntityData.Children = types.NewOrderedMap()
    nOTIFICATIONLOGMIB.EntityData.Children.Append("nlmConfig", types.YChild{"NlmConfig", &nOTIFICATIONLOGMIB.NlmConfig})
    nOTIFICATIONLOGMIB.EntityData.Children.Append("nlmStats", types.YChild{"NlmStats", &nOTIFICATIONLOGMIB.NlmStats})
    nOTIFICATIONLOGMIB.EntityData.Children.Append("nlmConfigLogTable", types.YChild{"NlmConfigLogTable", &nOTIFICATIONLOGMIB.NlmConfigLogTable})
    nOTIFICATIONLOGMIB.EntityData.Children.Append("nlmLogTable", types.YChild{"NlmLogTable", &nOTIFICATIONLOGMIB.NlmLogTable})
    nOTIFICATIONLOGMIB.EntityData.Children.Append("nlmLogVariableTable", types.YChild{"NlmLogVariableTable", &nOTIFICATIONLOGMIB.NlmLogVariableTable})
    nOTIFICATIONLOGMIB.EntityData.Leafs = types.NewOrderedMap()

    nOTIFICATIONLOGMIB.EntityData.YListKeys = []string {}

    return &(nOTIFICATIONLOGMIB.EntityData)
}

// NOTIFICATIONLOGMIB_NlmConfig
type NOTIFICATIONLOGMIB_NlmConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of notification entries that may be held in nlmLogTable
    // for all nlmLogNames added together.  A particular setting does not
    // guarantee that much data can be held.  If an application changes the limit
    // while there are Notifications in the log, the oldest Notifications MUST be
    // discarded to bring the log down to the new limit - thus the value of
    // nlmConfigGlobalEntryLimit MUST take precedence over the values of
    // nlmConfigGlobalAgeOut and nlmConfigLogEntryLimit, even if the Notification
    // being discarded has been present for fewer minutes than the value of
    // nlmConfigGlobalAgeOut, or if the named log has fewer entries than that
    // specified in nlmConfigLogEntryLimit.  A value of 0 means no limit.  Please
    // be aware that contention between multiple managers trying to set this
    // object to different values MAY affect the reliability and completeness of
    // data seen by each manager. The type is interface{} with range:
    // 0..4294967295.
    NlmConfigGlobalEntryLimit interface{}

    // The number of minutes a Notification SHOULD be kept in a log before it is
    // automatically removed.  If an application changes the value of
    // nlmConfigGlobalAgeOut, Notifications older than the new time MAY be
    // discarded to meet the new time.  A value of 0 means no age out.  Please be
    // aware that contention between multiple managers trying to set this object
    // to different values MAY affect the reliability and completeness of data
    // seen by each manager. The type is interface{} with range: 0..4294967295.
    // Units are minutes.
    NlmConfigGlobalAgeOut interface{}
}

func (nlmConfig *NOTIFICATIONLOGMIB_NlmConfig) GetEntityData() *types.CommonEntityData {
    nlmConfig.EntityData.YFilter = nlmConfig.YFilter
    nlmConfig.EntityData.YangName = "nlmConfig"
    nlmConfig.EntityData.BundleName = "cisco_ios_xe"
    nlmConfig.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmConfig.EntityData.SegmentPath = "nlmConfig"
    nlmConfig.EntityData.AbsolutePath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB/" + nlmConfig.EntityData.SegmentPath
    nlmConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmConfig.EntityData.Children = types.NewOrderedMap()
    nlmConfig.EntityData.Leafs = types.NewOrderedMap()
    nlmConfig.EntityData.Leafs.Append("nlmConfigGlobalEntryLimit", types.YLeaf{"NlmConfigGlobalEntryLimit", nlmConfig.NlmConfigGlobalEntryLimit})
    nlmConfig.EntityData.Leafs.Append("nlmConfigGlobalAgeOut", types.YLeaf{"NlmConfigGlobalAgeOut", nlmConfig.NlmConfigGlobalAgeOut})

    nlmConfig.EntityData.YListKeys = []string {}

    return &(nlmConfig.EntityData)
}

// NOTIFICATIONLOGMIB_NlmStats
type NOTIFICATIONLOGMIB_NlmStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of Notifications put into the nlmLogTable.  This counts a
    // Notification once for each log entry, so a Notification  put into multiple
    // logs is counted multiple times. The type is interface{} with range:
    // 0..4294967295. Units are notifications.
    NlmStatsGlobalNotificationsLogged interface{}

    // The number of log entries discarded to make room for a new entry due to
    // lack of resources or the value of nlmConfigGlobalEntryLimit or
    // nlmConfigLogEntryLimit.  This does not include entries discarded due to the
    // value of nlmConfigGlobalAgeOut. The type is interface{} with range:
    // 0..4294967295. Units are notifications.
    NlmStatsGlobalNotificationsBumped interface{}
}

func (nlmStats *NOTIFICATIONLOGMIB_NlmStats) GetEntityData() *types.CommonEntityData {
    nlmStats.EntityData.YFilter = nlmStats.YFilter
    nlmStats.EntityData.YangName = "nlmStats"
    nlmStats.EntityData.BundleName = "cisco_ios_xe"
    nlmStats.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmStats.EntityData.SegmentPath = "nlmStats"
    nlmStats.EntityData.AbsolutePath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB/" + nlmStats.EntityData.SegmentPath
    nlmStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmStats.EntityData.Children = types.NewOrderedMap()
    nlmStats.EntityData.Leafs = types.NewOrderedMap()
    nlmStats.EntityData.Leafs.Append("nlmStatsGlobalNotificationsLogged", types.YLeaf{"NlmStatsGlobalNotificationsLogged", nlmStats.NlmStatsGlobalNotificationsLogged})
    nlmStats.EntityData.Leafs.Append("nlmStatsGlobalNotificationsBumped", types.YLeaf{"NlmStatsGlobalNotificationsBumped", nlmStats.NlmStatsGlobalNotificationsBumped})

    nlmStats.EntityData.YListKeys = []string {}

    return &(nlmStats.EntityData)
}

// NOTIFICATIONLOGMIB_NlmConfigLogTable
// A table of logging control entries.
type NOTIFICATIONLOGMIB_NlmConfigLogTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A logging control entry.  Depending on the entry's storage type entries may
    // be supplied by the system or created and deleted by applications using
    // nlmConfigLogEntryStatus. The type is slice of
    // NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry.
    NlmConfigLogEntry []*NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry
}

func (nlmConfigLogTable *NOTIFICATIONLOGMIB_NlmConfigLogTable) GetEntityData() *types.CommonEntityData {
    nlmConfigLogTable.EntityData.YFilter = nlmConfigLogTable.YFilter
    nlmConfigLogTable.EntityData.YangName = "nlmConfigLogTable"
    nlmConfigLogTable.EntityData.BundleName = "cisco_ios_xe"
    nlmConfigLogTable.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmConfigLogTable.EntityData.SegmentPath = "nlmConfigLogTable"
    nlmConfigLogTable.EntityData.AbsolutePath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB/" + nlmConfigLogTable.EntityData.SegmentPath
    nlmConfigLogTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmConfigLogTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmConfigLogTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmConfigLogTable.EntityData.Children = types.NewOrderedMap()
    nlmConfigLogTable.EntityData.Children.Append("nlmConfigLogEntry", types.YChild{"NlmConfigLogEntry", nil})
    for i := range nlmConfigLogTable.NlmConfigLogEntry {
        nlmConfigLogTable.EntityData.Children.Append(types.GetSegmentPath(nlmConfigLogTable.NlmConfigLogEntry[i]), types.YChild{"NlmConfigLogEntry", nlmConfigLogTable.NlmConfigLogEntry[i]})
    }
    nlmConfigLogTable.EntityData.Leafs = types.NewOrderedMap()

    nlmConfigLogTable.EntityData.YListKeys = []string {}

    return &(nlmConfigLogTable.EntityData)
}

// NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry
// A logging control entry.  Depending on the entry's storage type
// entries may be supplied by the system or created and deleted by
// applications using nlmConfigLogEntryStatus.
type NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the log.  An implementation may allow
    // multiple named logs, up to some implementation-specific limit (which may be
    // none).  A zero-length log name is reserved for creation and deletion by the
    // managed system, and MUST be used as the default log name by systems that do
    // not support named logs. The type is string with length: 0..32.
    NlmLogName interface{}

    // A value of snmpNotifyFilterProfileName as used as an index into the
    // snmpNotifyFilterTable in the SNMP Notification MIB, specifying the locally
    // or remotely originated Notifications to be filtered out and not logged in
    // this log.  A zero-length value or a name that does not identify an existing
    // entry in snmpNotifyFilterTable indicate no Notifications are to be logged
    // in this log. The type is string with length: 0..32.
    NlmConfigLogFilterName interface{}

    // The maximum number of notification entries that can be held in nlmLogTable
    // for this named log.  A particular setting does not guarantee that that much
    // data can be held.  If an application changes the limit while there are
    // Notifications in the log, the oldest Notifications are discarded to bring
    // the log down to the new limit.  A value of 0 indicates no limit.  Please be
    // aware that contention between multiple managers trying to set this object
    // to different values MAY affect the reliability and completeness of data
    // seen by each manager. The type is interface{} with range: 0..4294967295.
    NlmConfigLogEntryLimit interface{}

    // Control to enable or disable the log without otherwise disturbing the log's
    // entry.  Please be aware that contention between multiple managers trying to
    // set this object to different values MAY affect the reliability and
    // completeness of data seen by each manager. The type is
    // NlmConfigLogAdminStatus.
    NlmConfigLogAdminStatus interface{}

    // The operational status of this log:  disabled  administratively disabled 
    // operational    administratively enabled and working  noFilter 
    // administratively enabled but either           nlmConfigLogFilterName is
    // zero length           or does not name an existing entry in          
    // snmpNotifyFilterTable. The type is NlmConfigLogOperStatus.
    NlmConfigLogOperStatus interface{}

    // The storage type of this conceptual row. The type is StorageType.
    NlmConfigLogStorageType interface{}

    // Control for creating and deleting entries.  Entries may be modified while
    // active.  For non-null-named logs, the managed system records the security
    // credentials from the request that sets nlmConfigLogStatus to 'active' and
    // uses that identity to apply access control to the objects in the
    // Notification to decide if that Notification may be logged. The type is
    // RowStatus.
    NlmConfigLogEntryStatus interface{}

    // The number of Notifications put in this named log. The type is interface{}
    // with range: 0..4294967295. Units are notifications.
    NlmStatsLogNotificationsLogged interface{}

    // The number of log entries discarded from this named log to make room for a
    // new entry due to lack of resources or the value of
    // nlmConfigGlobalEntryLimit or nlmConfigLogEntryLimit.  This does not include
    // entries discarded due to the value of nlmConfigGlobalAgeOut. The type is
    // interface{} with range: 0..4294967295. Units are notifications.
    NlmStatsLogNotificationsBumped interface{}
}

func (nlmConfigLogEntry *NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry) GetEntityData() *types.CommonEntityData {
    nlmConfigLogEntry.EntityData.YFilter = nlmConfigLogEntry.YFilter
    nlmConfigLogEntry.EntityData.YangName = "nlmConfigLogEntry"
    nlmConfigLogEntry.EntityData.BundleName = "cisco_ios_xe"
    nlmConfigLogEntry.EntityData.ParentYangName = "nlmConfigLogTable"
    nlmConfigLogEntry.EntityData.SegmentPath = "nlmConfigLogEntry" + types.AddKeyToken(nlmConfigLogEntry.NlmLogName, "nlmLogName")
    nlmConfigLogEntry.EntityData.AbsolutePath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB/nlmConfigLogTable/" + nlmConfigLogEntry.EntityData.SegmentPath
    nlmConfigLogEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmConfigLogEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmConfigLogEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmConfigLogEntry.EntityData.Children = types.NewOrderedMap()
    nlmConfigLogEntry.EntityData.Leafs = types.NewOrderedMap()
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmLogName", types.YLeaf{"NlmLogName", nlmConfigLogEntry.NlmLogName})
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmConfigLogFilterName", types.YLeaf{"NlmConfigLogFilterName", nlmConfigLogEntry.NlmConfigLogFilterName})
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmConfigLogEntryLimit", types.YLeaf{"NlmConfigLogEntryLimit", nlmConfigLogEntry.NlmConfigLogEntryLimit})
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmConfigLogAdminStatus", types.YLeaf{"NlmConfigLogAdminStatus", nlmConfigLogEntry.NlmConfigLogAdminStatus})
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmConfigLogOperStatus", types.YLeaf{"NlmConfigLogOperStatus", nlmConfigLogEntry.NlmConfigLogOperStatus})
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmConfigLogStorageType", types.YLeaf{"NlmConfigLogStorageType", nlmConfigLogEntry.NlmConfigLogStorageType})
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmConfigLogEntryStatus", types.YLeaf{"NlmConfigLogEntryStatus", nlmConfigLogEntry.NlmConfigLogEntryStatus})
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmStatsLogNotificationsLogged", types.YLeaf{"NlmStatsLogNotificationsLogged", nlmConfigLogEntry.NlmStatsLogNotificationsLogged})
    nlmConfigLogEntry.EntityData.Leafs.Append("nlmStatsLogNotificationsBumped", types.YLeaf{"NlmStatsLogNotificationsBumped", nlmConfigLogEntry.NlmStatsLogNotificationsBumped})

    nlmConfigLogEntry.EntityData.YListKeys = []string {"NlmLogName"}

    return &(nlmConfigLogEntry.EntityData)
}

// NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogAdminStatus represents reliability and completeness of data seen by each manager.
type NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogAdminStatus string

const (
    NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogAdminStatus_enabled NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogAdminStatus = "enabled"

    NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogAdminStatus_disabled NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogAdminStatus = "disabled"
)

// NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogOperStatus represents           snmpNotifyFilterTable
type NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogOperStatus string

const (
    NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogOperStatus_disabled NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogOperStatus = "disabled"

    NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogOperStatus_operational NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogOperStatus = "operational"

    NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogOperStatus_noFilter NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmConfigLogOperStatus = "noFilter"
)

// NOTIFICATIONLOGMIB_NlmLogTable
// A table of Notification log entries.
// 
// It is an implementation-specific matter whether entries in this
// table are preserved across initializations of the management
// system.  In general one would expect that they are not.
// 
// Note that keeping entries across initializations of the
// management system leads to some confusion with counters and
// TimeStamps, since both of those are based on sysUpTime, which
// resets on management initialization.  In this situation,
// counters apply only after the reset and nlmLogTime for entries
// made before the reset MUST be set to 0.
type NOTIFICATIONLOGMIB_NlmLogTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Notification log entry.  Entries appear in this table when Notifications
    // occur and pass filtering by nlmConfigLogFilterName and access control. 
    // They are removed to make way for new entries due to lack of resources or
    // the values of nlmConfigGlobalEntryLimit, nlmConfigGlobalAgeOut, or
    // nlmConfigLogEntryLimit.  If adding an entry would exceed
    // nlmConfigGlobalEntryLimit or system resources in general, the oldest entry
    // in any log SHOULD be removed to make room for the new one.  If adding an
    // entry would exceed nlmConfigLogEntryLimit the oldest entry in that log
    // SHOULD be removed to make room for the new one.  Before the managed system
    // puts a locally-generated Notification into a non-null-named log it assures
    // that the creator of the log has access to the information in the
    // Notification.  If not it does not log that Notification in that log. The
    // type is slice of NOTIFICATIONLOGMIB_NlmLogTable_NlmLogEntry.
    NlmLogEntry []*NOTIFICATIONLOGMIB_NlmLogTable_NlmLogEntry
}

func (nlmLogTable *NOTIFICATIONLOGMIB_NlmLogTable) GetEntityData() *types.CommonEntityData {
    nlmLogTable.EntityData.YFilter = nlmLogTable.YFilter
    nlmLogTable.EntityData.YangName = "nlmLogTable"
    nlmLogTable.EntityData.BundleName = "cisco_ios_xe"
    nlmLogTable.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmLogTable.EntityData.SegmentPath = "nlmLogTable"
    nlmLogTable.EntityData.AbsolutePath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB/" + nlmLogTable.EntityData.SegmentPath
    nlmLogTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmLogTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmLogTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmLogTable.EntityData.Children = types.NewOrderedMap()
    nlmLogTable.EntityData.Children.Append("nlmLogEntry", types.YChild{"NlmLogEntry", nil})
    for i := range nlmLogTable.NlmLogEntry {
        nlmLogTable.EntityData.Children.Append(types.GetSegmentPath(nlmLogTable.NlmLogEntry[i]), types.YChild{"NlmLogEntry", nlmLogTable.NlmLogEntry[i]})
    }
    nlmLogTable.EntityData.Leafs = types.NewOrderedMap()

    nlmLogTable.EntityData.YListKeys = []string {}

    return &(nlmLogTable.EntityData)
}

// NOTIFICATIONLOGMIB_NlmLogTable_NlmLogEntry
// A Notification log entry.
// 
// Entries appear in this table when Notifications occur and pass
// filtering by nlmConfigLogFilterName and access control.  They are
// removed to make way for new entries due to lack of resources or
// the values of nlmConfigGlobalEntryLimit, nlmConfigGlobalAgeOut, or
// nlmConfigLogEntryLimit.
// 
// If adding an entry would exceed nlmConfigGlobalEntryLimit or system
// resources in general, the oldest entry in any log SHOULD be removed
// to make room for the new one.
// 
// If adding an entry would exceed nlmConfigLogEntryLimit the oldest
// entry in that log SHOULD be removed to make room for the new one.
// 
// Before the managed system puts a locally-generated Notification
// into a non-null-named log it assures that the creator of the log
// has access to the information in the Notification.  If not it
// does not log that Notification in that log.
type NOTIFICATIONLOGMIB_NlmLogTable_NlmLogEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // notification_log_mib.NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmLogName
    NlmLogName interface{}

    // This attribute is a key. A monotonically increasing integer for the sole
    // purpose of indexing entries within the named log.  When it reaches the
    // maximum value, an extremely unlikely event, the agent wraps the value back
    // to 1. The type is interface{} with range: 1..4294967295.
    NlmLogIndex interface{}

    // The value of sysUpTime when the entry was placed in the log. If the entry
    // occurred before the most recent management system initialization this
    // object value MUST be set to zero. The type is interface{} with range:
    // 0..4294967295.
    NlmLogTime interface{}

    // The local date and time when the entry was logged, instantiated only by
    // systems that have date and time capability. The type is string.
    NlmLogDateAndTime interface{}

    // The identification of the SNMP engine at which the Notification originated.
    // If the log can contain Notifications from only one engine or the Trap is in
    // SNMPv1 format, this object is a zero-length string. The type is string with
    // length: 5..32.
    NlmLogEngineID interface{}

    // The transport service address of the SNMP engine from which the
    // Notification was received, formatted according to the corresponding value
    // of nlmLogEngineTDomain. This is used to identify the source of an SNMPv1
    // trap, since an nlmLogEngineId cannot be extracted from the SNMPv1 trap pdu.
    // This object MUST always be instantiated, even if the log can contain
    // Notifications from only one engine.  Please be aware that the
    // nlmLogEngineTAddress may not uniquely identify the SNMP engine from which
    // the Notification was received. For example, if an SNMP engine uses DHCP or
    // NAT to obtain ip addresses, the address it uses may be shared with other
    // network devices, and hence will not uniquely identify the SNMP engine. The
    // type is string with length: 1..255.
    NlmLogEngineTAddress interface{}

    // Indicates the kind of transport service by which a Notification was
    // received from an SNMP engine. nlmLogEngineTAddress contains the transport
    // service address of the SNMP engine from which this Notification was
    // received.  Possible values for this object are presently found in the
    // Transport Mappings for SNMPv2 document (RFC 1906 [8]). The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    NlmLogEngineTDomain interface{}

    // If the Notification was received in a protocol which has a contextEngineID
    // element like SNMPv3, this object has that value. Otherwise its value is a
    // zero-length string. The type is string with length: 5..32.
    NlmLogContextEngineID interface{}

    // The name of the SNMP MIB context from which the Notification came. For
    // SNMPv1 Traps this is the community string from the Trap. The type is
    // string.
    NlmLogContextName interface{}

    // The NOTIFICATION-TYPE object identifier of the Notification that occurred.
    // The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    NlmLogNotificationID interface{}
}

func (nlmLogEntry *NOTIFICATIONLOGMIB_NlmLogTable_NlmLogEntry) GetEntityData() *types.CommonEntityData {
    nlmLogEntry.EntityData.YFilter = nlmLogEntry.YFilter
    nlmLogEntry.EntityData.YangName = "nlmLogEntry"
    nlmLogEntry.EntityData.BundleName = "cisco_ios_xe"
    nlmLogEntry.EntityData.ParentYangName = "nlmLogTable"
    nlmLogEntry.EntityData.SegmentPath = "nlmLogEntry" + types.AddKeyToken(nlmLogEntry.NlmLogName, "nlmLogName") + types.AddKeyToken(nlmLogEntry.NlmLogIndex, "nlmLogIndex")
    nlmLogEntry.EntityData.AbsolutePath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB/nlmLogTable/" + nlmLogEntry.EntityData.SegmentPath
    nlmLogEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmLogEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmLogEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmLogEntry.EntityData.Children = types.NewOrderedMap()
    nlmLogEntry.EntityData.Leafs = types.NewOrderedMap()
    nlmLogEntry.EntityData.Leafs.Append("nlmLogName", types.YLeaf{"NlmLogName", nlmLogEntry.NlmLogName})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogIndex", types.YLeaf{"NlmLogIndex", nlmLogEntry.NlmLogIndex})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogTime", types.YLeaf{"NlmLogTime", nlmLogEntry.NlmLogTime})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogDateAndTime", types.YLeaf{"NlmLogDateAndTime", nlmLogEntry.NlmLogDateAndTime})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogEngineID", types.YLeaf{"NlmLogEngineID", nlmLogEntry.NlmLogEngineID})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogEngineTAddress", types.YLeaf{"NlmLogEngineTAddress", nlmLogEntry.NlmLogEngineTAddress})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogEngineTDomain", types.YLeaf{"NlmLogEngineTDomain", nlmLogEntry.NlmLogEngineTDomain})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogContextEngineID", types.YLeaf{"NlmLogContextEngineID", nlmLogEntry.NlmLogContextEngineID})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogContextName", types.YLeaf{"NlmLogContextName", nlmLogEntry.NlmLogContextName})
    nlmLogEntry.EntityData.Leafs.Append("nlmLogNotificationID", types.YLeaf{"NlmLogNotificationID", nlmLogEntry.NlmLogNotificationID})

    nlmLogEntry.EntityData.YListKeys = []string {"NlmLogName", "NlmLogIndex"}

    return &(nlmLogEntry.EntityData)
}

// NOTIFICATIONLOGMIB_NlmLogVariableTable
// A table of variables to go with Notification log entries.
type NOTIFICATIONLOGMIB_NlmLogVariableTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Notification log entry variable.  Entries appear in this table when there
    // are variables in the varbind list of a Notification in nlmLogTable. The
    // type is slice of
    // NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry.
    NlmLogVariableEntry []*NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry
}

func (nlmLogVariableTable *NOTIFICATIONLOGMIB_NlmLogVariableTable) GetEntityData() *types.CommonEntityData {
    nlmLogVariableTable.EntityData.YFilter = nlmLogVariableTable.YFilter
    nlmLogVariableTable.EntityData.YangName = "nlmLogVariableTable"
    nlmLogVariableTable.EntityData.BundleName = "cisco_ios_xe"
    nlmLogVariableTable.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmLogVariableTable.EntityData.SegmentPath = "nlmLogVariableTable"
    nlmLogVariableTable.EntityData.AbsolutePath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB/" + nlmLogVariableTable.EntityData.SegmentPath
    nlmLogVariableTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmLogVariableTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmLogVariableTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmLogVariableTable.EntityData.Children = types.NewOrderedMap()
    nlmLogVariableTable.EntityData.Children.Append("nlmLogVariableEntry", types.YChild{"NlmLogVariableEntry", nil})
    for i := range nlmLogVariableTable.NlmLogVariableEntry {
        nlmLogVariableTable.EntityData.Children.Append(types.GetSegmentPath(nlmLogVariableTable.NlmLogVariableEntry[i]), types.YChild{"NlmLogVariableEntry", nlmLogVariableTable.NlmLogVariableEntry[i]})
    }
    nlmLogVariableTable.EntityData.Leafs = types.NewOrderedMap()

    nlmLogVariableTable.EntityData.YListKeys = []string {}

    return &(nlmLogVariableTable.EntityData)
}

// NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry
// A Notification log entry variable.
// 
// Entries appear in this table when there are variables in
// the varbind list of a Notification in nlmLogTable.
type NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // notification_log_mib.NOTIFICATIONLOGMIB_NlmConfigLogTable_NlmConfigLogEntry_NlmLogName
    NlmLogName interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // notification_log_mib.NOTIFICATIONLOGMIB_NlmLogTable_NlmLogEntry_NlmLogIndex
    NlmLogIndex interface{}

    // This attribute is a key. A monotonically increasing integer, starting at 1
    // for a given nlmLogIndex, for indexing variables within the logged
    // Notification. The type is interface{} with range: 1..4294967295.
    NlmLogVariableIndex interface{}

    // The variable's object identifier. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    NlmLogVariableID interface{}

    // The type of the value.  One and only one of the value objects that follow
    // must be instantiated, based on this type. The type is
    // NlmLogVariableValueType.
    NlmLogVariableValueType interface{}

    // The value when nlmLogVariableType is 'counter32'. The type is interface{}
    // with range: 0..4294967295.
    NlmLogVariableCounter32Val interface{}

    // The value when nlmLogVariableType is 'unsigned32'. The type is interface{}
    // with range: 0..4294967295.
    NlmLogVariableUnsigned32Val interface{}

    // The value when nlmLogVariableType is 'timeTicks'. The type is interface{}
    // with range: 0..4294967295.
    NlmLogVariableTimeTicksVal interface{}

    // The value when nlmLogVariableType is 'integer32'. The type is interface{}
    // with range: -2147483648..2147483647.
    NlmLogVariableInteger32Val interface{}

    // The value when nlmLogVariableType is 'octetString'. The type is string.
    NlmLogVariableOctetStringVal interface{}

    // The value when nlmLogVariableType is 'ipAddress'. Although this seems to be
    // unfriendly for IPv6, we have to recognize that there are a number of older
    // MIBs that do contain an IPv4 format address, known as IpAddress.  IPv6
    // addresses are represented using TAddress or InetAddress, and so the
    // underlying datatype is OCTET STRING, and their value would be stored in the
    // nlmLogVariableOctetStringVal column. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NlmLogVariableIpAddressVal interface{}

    // The value when nlmLogVariableType is 'objectId'. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    NlmLogVariableOidVal interface{}

    // The value when nlmLogVariableType is 'counter64'. The type is interface{}
    // with range: 0..18446744073709551615.
    NlmLogVariableCounter64Val interface{}

    // The value when nlmLogVariableType is 'opaque'. The type is string.
    NlmLogVariableOpaqueVal interface{}
}

func (nlmLogVariableEntry *NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry) GetEntityData() *types.CommonEntityData {
    nlmLogVariableEntry.EntityData.YFilter = nlmLogVariableEntry.YFilter
    nlmLogVariableEntry.EntityData.YangName = "nlmLogVariableEntry"
    nlmLogVariableEntry.EntityData.BundleName = "cisco_ios_xe"
    nlmLogVariableEntry.EntityData.ParentYangName = "nlmLogVariableTable"
    nlmLogVariableEntry.EntityData.SegmentPath = "nlmLogVariableEntry" + types.AddKeyToken(nlmLogVariableEntry.NlmLogName, "nlmLogName") + types.AddKeyToken(nlmLogVariableEntry.NlmLogIndex, "nlmLogIndex") + types.AddKeyToken(nlmLogVariableEntry.NlmLogVariableIndex, "nlmLogVariableIndex")
    nlmLogVariableEntry.EntityData.AbsolutePath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB/nlmLogVariableTable/" + nlmLogVariableEntry.EntityData.SegmentPath
    nlmLogVariableEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmLogVariableEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmLogVariableEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmLogVariableEntry.EntityData.Children = types.NewOrderedMap()
    nlmLogVariableEntry.EntityData.Leafs = types.NewOrderedMap()
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogName", types.YLeaf{"NlmLogName", nlmLogVariableEntry.NlmLogName})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogIndex", types.YLeaf{"NlmLogIndex", nlmLogVariableEntry.NlmLogIndex})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableIndex", types.YLeaf{"NlmLogVariableIndex", nlmLogVariableEntry.NlmLogVariableIndex})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableID", types.YLeaf{"NlmLogVariableID", nlmLogVariableEntry.NlmLogVariableID})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableValueType", types.YLeaf{"NlmLogVariableValueType", nlmLogVariableEntry.NlmLogVariableValueType})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableCounter32Val", types.YLeaf{"NlmLogVariableCounter32Val", nlmLogVariableEntry.NlmLogVariableCounter32Val})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableUnsigned32Val", types.YLeaf{"NlmLogVariableUnsigned32Val", nlmLogVariableEntry.NlmLogVariableUnsigned32Val})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableTimeTicksVal", types.YLeaf{"NlmLogVariableTimeTicksVal", nlmLogVariableEntry.NlmLogVariableTimeTicksVal})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableInteger32Val", types.YLeaf{"NlmLogVariableInteger32Val", nlmLogVariableEntry.NlmLogVariableInteger32Val})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableOctetStringVal", types.YLeaf{"NlmLogVariableOctetStringVal", nlmLogVariableEntry.NlmLogVariableOctetStringVal})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableIpAddressVal", types.YLeaf{"NlmLogVariableIpAddressVal", nlmLogVariableEntry.NlmLogVariableIpAddressVal})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableOidVal", types.YLeaf{"NlmLogVariableOidVal", nlmLogVariableEntry.NlmLogVariableOidVal})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableCounter64Val", types.YLeaf{"NlmLogVariableCounter64Val", nlmLogVariableEntry.NlmLogVariableCounter64Val})
    nlmLogVariableEntry.EntityData.Leafs.Append("nlmLogVariableOpaqueVal", types.YLeaf{"NlmLogVariableOpaqueVal", nlmLogVariableEntry.NlmLogVariableOpaqueVal})

    nlmLogVariableEntry.EntityData.YListKeys = []string {"NlmLogName", "NlmLogIndex", "NlmLogVariableIndex"}

    return &(nlmLogVariableEntry.EntityData)
}

// NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType represents objects that follow must be instantiated, based on this type.
type NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType string

const (
    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_counter32 NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "counter32"

    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_unsigned32 NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "unsigned32"

    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_timeTicks NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "timeTicks"

    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_integer32 NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "integer32"

    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_ipAddress NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "ipAddress"

    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_octetString NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "octetString"

    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_objectId NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "objectId"

    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_counter64 NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "counter64"

    NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType_opaque NOTIFICATIONLOGMIB_NlmLogVariableTable_NlmLogVariableEntry_NlmLogVariableValueType = "opaque"
)

