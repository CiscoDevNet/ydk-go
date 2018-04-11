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

    
    Nlmconfig NOTIFICATIONLOGMIB_Nlmconfig

    
    Nlmstats NOTIFICATIONLOGMIB_Nlmstats

    // A table of logging control entries.
    Nlmconfiglogtable NOTIFICATIONLOGMIB_Nlmconfiglogtable

    // A table of Notification log entries.  It is an implementation-specific
    // matter whether entries in this table are preserved across initializations
    // of the management system.  In general one would expect that they are not. 
    // Note that keeping entries across initializations of the management system
    // leads to some confusion with counters and TimeStamps, since both of those
    // are based on sysUpTime, which resets on management initialization.  In this
    // situation, counters apply only after the reset and nlmLogTime for entries
    // made before the reset MUST be set to 0.
    Nlmlogtable NOTIFICATIONLOGMIB_Nlmlogtable

    // A table of variables to go with Notification log entries.
    Nlmlogvariabletable NOTIFICATIONLOGMIB_Nlmlogvariabletable
}

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetEntityData() *types.CommonEntityData {
    nOTIFICATIONLOGMIB.EntityData.YFilter = nOTIFICATIONLOGMIB.YFilter
    nOTIFICATIONLOGMIB.EntityData.YangName = "NOTIFICATION-LOG-MIB"
    nOTIFICATIONLOGMIB.EntityData.BundleName = "cisco_ios_xe"
    nOTIFICATIONLOGMIB.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nOTIFICATIONLOGMIB.EntityData.SegmentPath = "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB"
    nOTIFICATIONLOGMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nOTIFICATIONLOGMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nOTIFICATIONLOGMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nOTIFICATIONLOGMIB.EntityData.Children = make(map[string]types.YChild)
    nOTIFICATIONLOGMIB.EntityData.Children["nlmConfig"] = types.YChild{"Nlmconfig", &nOTIFICATIONLOGMIB.Nlmconfig}
    nOTIFICATIONLOGMIB.EntityData.Children["nlmStats"] = types.YChild{"Nlmstats", &nOTIFICATIONLOGMIB.Nlmstats}
    nOTIFICATIONLOGMIB.EntityData.Children["nlmConfigLogTable"] = types.YChild{"Nlmconfiglogtable", &nOTIFICATIONLOGMIB.Nlmconfiglogtable}
    nOTIFICATIONLOGMIB.EntityData.Children["nlmLogTable"] = types.YChild{"Nlmlogtable", &nOTIFICATIONLOGMIB.Nlmlogtable}
    nOTIFICATIONLOGMIB.EntityData.Children["nlmLogVariableTable"] = types.YChild{"Nlmlogvariabletable", &nOTIFICATIONLOGMIB.Nlmlogvariabletable}
    nOTIFICATIONLOGMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nOTIFICATIONLOGMIB.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmconfig
type NOTIFICATIONLOGMIB_Nlmconfig struct {
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
    Nlmconfigglobalentrylimit interface{}

    // The number of minutes a Notification SHOULD be kept in a log before it is
    // automatically removed.  If an application changes the value of
    // nlmConfigGlobalAgeOut, Notifications older than the new time MAY be
    // discarded to meet the new time.  A value of 0 means no age out.  Please be
    // aware that contention between multiple managers trying to set this object
    // to different values MAY affect the reliability and completeness of data
    // seen by each manager. The type is interface{} with range: 0..4294967295.
    // Units are minutes.
    Nlmconfigglobalageout interface{}
}

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetEntityData() *types.CommonEntityData {
    nlmconfig.EntityData.YFilter = nlmconfig.YFilter
    nlmconfig.EntityData.YangName = "nlmConfig"
    nlmconfig.EntityData.BundleName = "cisco_ios_xe"
    nlmconfig.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmconfig.EntityData.SegmentPath = "nlmConfig"
    nlmconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmconfig.EntityData.Children = make(map[string]types.YChild)
    nlmconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmconfig.EntityData.Leafs["nlmConfigGlobalEntryLimit"] = types.YLeaf{"Nlmconfigglobalentrylimit", nlmconfig.Nlmconfigglobalentrylimit}
    nlmconfig.EntityData.Leafs["nlmConfigGlobalAgeOut"] = types.YLeaf{"Nlmconfigglobalageout", nlmconfig.Nlmconfigglobalageout}
    return &(nlmconfig.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmstats
type NOTIFICATIONLOGMIB_Nlmstats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of Notifications put into the nlmLogTable.  This counts a
    // Notification once for each log entry, so a Notification  put into multiple
    // logs is counted multiple times. The type is interface{} with range:
    // 0..4294967295. Units are notifications.
    Nlmstatsglobalnotificationslogged interface{}

    // The number of log entries discarded to make room for a new entry due to
    // lack of resources or the value of nlmConfigGlobalEntryLimit or
    // nlmConfigLogEntryLimit.  This does not include entries discarded due to the
    // value of nlmConfigGlobalAgeOut. The type is interface{} with range:
    // 0..4294967295. Units are notifications.
    Nlmstatsglobalnotificationsbumped interface{}
}

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetEntityData() *types.CommonEntityData {
    nlmstats.EntityData.YFilter = nlmstats.YFilter
    nlmstats.EntityData.YangName = "nlmStats"
    nlmstats.EntityData.BundleName = "cisco_ios_xe"
    nlmstats.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmstats.EntityData.SegmentPath = "nlmStats"
    nlmstats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmstats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmstats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmstats.EntityData.Children = make(map[string]types.YChild)
    nlmstats.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmstats.EntityData.Leafs["nlmStatsGlobalNotificationsLogged"] = types.YLeaf{"Nlmstatsglobalnotificationslogged", nlmstats.Nlmstatsglobalnotificationslogged}
    nlmstats.EntityData.Leafs["nlmStatsGlobalNotificationsBumped"] = types.YLeaf{"Nlmstatsglobalnotificationsbumped", nlmstats.Nlmstatsglobalnotificationsbumped}
    return &(nlmstats.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmconfiglogtable
// A table of logging control entries.
type NOTIFICATIONLOGMIB_Nlmconfiglogtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A logging control entry.  Depending on the entry's storage type entries may
    // be supplied by the system or created and deleted by applications using
    // nlmConfigLogEntryStatus. The type is slice of
    // NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry.
    Nlmconfiglogentry []NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry
}

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetEntityData() *types.CommonEntityData {
    nlmconfiglogtable.EntityData.YFilter = nlmconfiglogtable.YFilter
    nlmconfiglogtable.EntityData.YangName = "nlmConfigLogTable"
    nlmconfiglogtable.EntityData.BundleName = "cisco_ios_xe"
    nlmconfiglogtable.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmconfiglogtable.EntityData.SegmentPath = "nlmConfigLogTable"
    nlmconfiglogtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmconfiglogtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmconfiglogtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmconfiglogtable.EntityData.Children = make(map[string]types.YChild)
    nlmconfiglogtable.EntityData.Children["nlmConfigLogEntry"] = types.YChild{"Nlmconfiglogentry", nil}
    for i := range nlmconfiglogtable.Nlmconfiglogentry {
        nlmconfiglogtable.EntityData.Children[types.GetSegmentPath(&nlmconfiglogtable.Nlmconfiglogentry[i])] = types.YChild{"Nlmconfiglogentry", &nlmconfiglogtable.Nlmconfiglogentry[i]}
    }
    nlmconfiglogtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nlmconfiglogtable.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry
// A logging control entry.  Depending on the entry's storage type
// entries may be supplied by the system or created and deleted by
// applications using nlmConfigLogEntryStatus.
type NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the log.  An implementation may allow
    // multiple named logs, up to some implementation-specific limit (which may be
    // none).  A zero-length log name is reserved for creation and deletion by the
    // managed system, and MUST be used as the default log name by systems that do
    // not support named logs. The type is string with length: 0..32.
    Nlmlogname interface{}

    // A value of snmpNotifyFilterProfileName as used as an index into the
    // snmpNotifyFilterTable in the SNMP Notification MIB, specifying the locally
    // or remotely originated Notifications to be filtered out and not logged in
    // this log.  A zero-length value or a name that does not identify an existing
    // entry in snmpNotifyFilterTable indicate no Notifications are to be logged
    // in this log. The type is string with length: 0..32.
    Nlmconfiglogfiltername interface{}

    // The maximum number of notification entries that can be held in nlmLogTable
    // for this named log.  A particular setting does not guarantee that that much
    // data can be held.  If an application changes the limit while there are
    // Notifications in the log, the oldest Notifications are discarded to bring
    // the log down to the new limit.  A value of 0 indicates no limit.  Please be
    // aware that contention between multiple managers trying to set this object
    // to different values MAY affect the reliability and completeness of data
    // seen by each manager. The type is interface{} with range: 0..4294967295.
    Nlmconfiglogentrylimit interface{}

    // Control to enable or disable the log without otherwise disturbing the log's
    // entry.  Please be aware that contention between multiple managers trying to
    // set this object to different values MAY affect the reliability and
    // completeness of data seen by each manager. The type is
    // Nlmconfiglogadminstatus.
    Nlmconfiglogadminstatus interface{}

    // The operational status of this log:  disabled  administratively disabled 
    // operational    administratively enabled and working  noFilter 
    // administratively enabled but either           nlmConfigLogFilterName is
    // zero length           or does not name an existing entry in          
    // snmpNotifyFilterTable. The type is Nlmconfiglogoperstatus.
    Nlmconfiglogoperstatus interface{}

    // The storage type of this conceptual row. The type is StorageType.
    Nlmconfiglogstoragetype interface{}

    // Control for creating and deleting entries.  Entries may be modified while
    // active.  For non-null-named logs, the managed system records the security
    // credentials from the request that sets nlmConfigLogStatus to 'active' and
    // uses that identity to apply access control to the objects in the
    // Notification to decide if that Notification may be logged. The type is
    // RowStatus.
    Nlmconfiglogentrystatus interface{}

    // The number of Notifications put in this named log. The type is interface{}
    // with range: 0..4294967295. Units are notifications.
    Nlmstatslognotificationslogged interface{}

    // The number of log entries discarded from this named log to make room for a
    // new entry due to lack of resources or the value of
    // nlmConfigGlobalEntryLimit or nlmConfigLogEntryLimit.  This does not include
    // entries discarded due to the value of nlmConfigGlobalAgeOut. The type is
    // interface{} with range: 0..4294967295. Units are notifications.
    Nlmstatslognotificationsbumped interface{}
}

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetEntityData() *types.CommonEntityData {
    nlmconfiglogentry.EntityData.YFilter = nlmconfiglogentry.YFilter
    nlmconfiglogentry.EntityData.YangName = "nlmConfigLogEntry"
    nlmconfiglogentry.EntityData.BundleName = "cisco_ios_xe"
    nlmconfiglogentry.EntityData.ParentYangName = "nlmConfigLogTable"
    nlmconfiglogentry.EntityData.SegmentPath = "nlmConfigLogEntry" + "[nlmLogName='" + fmt.Sprintf("%v", nlmconfiglogentry.Nlmlogname) + "']"
    nlmconfiglogentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmconfiglogentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmconfiglogentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmconfiglogentry.EntityData.Children = make(map[string]types.YChild)
    nlmconfiglogentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmconfiglogentry.EntityData.Leafs["nlmLogName"] = types.YLeaf{"Nlmlogname", nlmconfiglogentry.Nlmlogname}
    nlmconfiglogentry.EntityData.Leafs["nlmConfigLogFilterName"] = types.YLeaf{"Nlmconfiglogfiltername", nlmconfiglogentry.Nlmconfiglogfiltername}
    nlmconfiglogentry.EntityData.Leafs["nlmConfigLogEntryLimit"] = types.YLeaf{"Nlmconfiglogentrylimit", nlmconfiglogentry.Nlmconfiglogentrylimit}
    nlmconfiglogentry.EntityData.Leafs["nlmConfigLogAdminStatus"] = types.YLeaf{"Nlmconfiglogadminstatus", nlmconfiglogentry.Nlmconfiglogadminstatus}
    nlmconfiglogentry.EntityData.Leafs["nlmConfigLogOperStatus"] = types.YLeaf{"Nlmconfiglogoperstatus", nlmconfiglogentry.Nlmconfiglogoperstatus}
    nlmconfiglogentry.EntityData.Leafs["nlmConfigLogStorageType"] = types.YLeaf{"Nlmconfiglogstoragetype", nlmconfiglogentry.Nlmconfiglogstoragetype}
    nlmconfiglogentry.EntityData.Leafs["nlmConfigLogEntryStatus"] = types.YLeaf{"Nlmconfiglogentrystatus", nlmconfiglogentry.Nlmconfiglogentrystatus}
    nlmconfiglogentry.EntityData.Leafs["nlmStatsLogNotificationsLogged"] = types.YLeaf{"Nlmstatslognotificationslogged", nlmconfiglogentry.Nlmstatslognotificationslogged}
    nlmconfiglogentry.EntityData.Leafs["nlmStatsLogNotificationsBumped"] = types.YLeaf{"Nlmstatslognotificationsbumped", nlmconfiglogentry.Nlmstatslognotificationsbumped}
    return &(nlmconfiglogentry.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogadminstatus represents reliability and completeness of data seen by each manager.
type NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogadminstatus string

const (
    NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogadminstatus_enabled NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogadminstatus = "enabled"

    NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogadminstatus_disabled NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogadminstatus = "disabled"
)

// NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogoperstatus represents           snmpNotifyFilterTable
type NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogoperstatus string

const (
    NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogoperstatus_disabled NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogoperstatus = "disabled"

    NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogoperstatus_operational NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogoperstatus = "operational"

    NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogoperstatus_noFilter NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmconfiglogoperstatus = "noFilter"
)

// NOTIFICATIONLOGMIB_Nlmlogtable
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
type NOTIFICATIONLOGMIB_Nlmlogtable struct {
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
    // type is slice of NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry.
    Nlmlogentry []NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry
}

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetEntityData() *types.CommonEntityData {
    nlmlogtable.EntityData.YFilter = nlmlogtable.YFilter
    nlmlogtable.EntityData.YangName = "nlmLogTable"
    nlmlogtable.EntityData.BundleName = "cisco_ios_xe"
    nlmlogtable.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmlogtable.EntityData.SegmentPath = "nlmLogTable"
    nlmlogtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmlogtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmlogtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmlogtable.EntityData.Children = make(map[string]types.YChild)
    nlmlogtable.EntityData.Children["nlmLogEntry"] = types.YChild{"Nlmlogentry", nil}
    for i := range nlmlogtable.Nlmlogentry {
        nlmlogtable.EntityData.Children[types.GetSegmentPath(&nlmlogtable.Nlmlogentry[i])] = types.YChild{"Nlmlogentry", &nlmlogtable.Nlmlogentry[i]}
    }
    nlmlogtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nlmlogtable.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry
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
type NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // notification_log_mib.NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmlogname
    Nlmlogname interface{}

    // This attribute is a key. A monotonically increasing integer for the sole
    // purpose of indexing entries within the named log.  When it reaches the
    // maximum value, an extremely unlikely event, the agent wraps the value back
    // to 1. The type is interface{} with range: 1..4294967295.
    Nlmlogindex interface{}

    // The value of sysUpTime when the entry was placed in the log. If the entry
    // occurred before the most recent management system initialization this
    // object value MUST be set to zero. The type is interface{} with range:
    // 0..4294967295.
    Nlmlogtime interface{}

    // The local date and time when the entry was logged, instantiated only by
    // systems that have date and time capability. The type is string.
    Nlmlogdateandtime interface{}

    // The identification of the SNMP engine at which the Notification originated.
    // If the log can contain Notifications from only one engine or the Trap is in
    // SNMPv1 format, this object is a zero-length string. The type is string with
    // length: 5..32.
    Nlmlogengineid interface{}

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
    Nlmlogenginetaddress interface{}

    // Indicates the kind of transport service by which a Notification was
    // received from an SNMP engine. nlmLogEngineTAddress contains the transport
    // service address of the SNMP engine from which this Notification was
    // received.  Possible values for this object are presently found in the
    // Transport Mappings for SNMPv2 document (RFC 1906 [8]). The type is string
    // with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Nlmlogenginetdomain interface{}

    // If the Notification was received in a protocol which has a contextEngineID
    // element like SNMPv3, this object has that value. Otherwise its value is a
    // zero-length string. The type is string with length: 5..32.
    Nlmlogcontextengineid interface{}

    // The name of the SNMP MIB context from which the Notification came. For
    // SNMPv1 Traps this is the community string from the Trap. The type is
    // string.
    Nlmlogcontextname interface{}

    // The NOTIFICATION-TYPE object identifier of the Notification that occurred.
    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Nlmlognotificationid interface{}
}

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetEntityData() *types.CommonEntityData {
    nlmlogentry.EntityData.YFilter = nlmlogentry.YFilter
    nlmlogentry.EntityData.YangName = "nlmLogEntry"
    nlmlogentry.EntityData.BundleName = "cisco_ios_xe"
    nlmlogentry.EntityData.ParentYangName = "nlmLogTable"
    nlmlogentry.EntityData.SegmentPath = "nlmLogEntry" + "[nlmLogName='" + fmt.Sprintf("%v", nlmlogentry.Nlmlogname) + "']" + "[nlmLogIndex='" + fmt.Sprintf("%v", nlmlogentry.Nlmlogindex) + "']"
    nlmlogentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmlogentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmlogentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmlogentry.EntityData.Children = make(map[string]types.YChild)
    nlmlogentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmlogentry.EntityData.Leafs["nlmLogName"] = types.YLeaf{"Nlmlogname", nlmlogentry.Nlmlogname}
    nlmlogentry.EntityData.Leafs["nlmLogIndex"] = types.YLeaf{"Nlmlogindex", nlmlogentry.Nlmlogindex}
    nlmlogentry.EntityData.Leafs["nlmLogTime"] = types.YLeaf{"Nlmlogtime", nlmlogentry.Nlmlogtime}
    nlmlogentry.EntityData.Leafs["nlmLogDateAndTime"] = types.YLeaf{"Nlmlogdateandtime", nlmlogentry.Nlmlogdateandtime}
    nlmlogentry.EntityData.Leafs["nlmLogEngineID"] = types.YLeaf{"Nlmlogengineid", nlmlogentry.Nlmlogengineid}
    nlmlogentry.EntityData.Leafs["nlmLogEngineTAddress"] = types.YLeaf{"Nlmlogenginetaddress", nlmlogentry.Nlmlogenginetaddress}
    nlmlogentry.EntityData.Leafs["nlmLogEngineTDomain"] = types.YLeaf{"Nlmlogenginetdomain", nlmlogentry.Nlmlogenginetdomain}
    nlmlogentry.EntityData.Leafs["nlmLogContextEngineID"] = types.YLeaf{"Nlmlogcontextengineid", nlmlogentry.Nlmlogcontextengineid}
    nlmlogentry.EntityData.Leafs["nlmLogContextName"] = types.YLeaf{"Nlmlogcontextname", nlmlogentry.Nlmlogcontextname}
    nlmlogentry.EntityData.Leafs["nlmLogNotificationID"] = types.YLeaf{"Nlmlognotificationid", nlmlogentry.Nlmlognotificationid}
    return &(nlmlogentry.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmlogvariabletable
// A table of variables to go with Notification log entries.
type NOTIFICATIONLOGMIB_Nlmlogvariabletable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Notification log entry variable.  Entries appear in this table when there
    // are variables in the varbind list of a Notification in nlmLogTable. The
    // type is slice of
    // NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry.
    Nlmlogvariableentry []NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry
}

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetEntityData() *types.CommonEntityData {
    nlmlogvariabletable.EntityData.YFilter = nlmlogvariabletable.YFilter
    nlmlogvariabletable.EntityData.YangName = "nlmLogVariableTable"
    nlmlogvariabletable.EntityData.BundleName = "cisco_ios_xe"
    nlmlogvariabletable.EntityData.ParentYangName = "NOTIFICATION-LOG-MIB"
    nlmlogvariabletable.EntityData.SegmentPath = "nlmLogVariableTable"
    nlmlogvariabletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmlogvariabletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmlogvariabletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmlogvariabletable.EntityData.Children = make(map[string]types.YChild)
    nlmlogvariabletable.EntityData.Children["nlmLogVariableEntry"] = types.YChild{"Nlmlogvariableentry", nil}
    for i := range nlmlogvariabletable.Nlmlogvariableentry {
        nlmlogvariabletable.EntityData.Children[types.GetSegmentPath(&nlmlogvariabletable.Nlmlogvariableentry[i])] = types.YChild{"Nlmlogvariableentry", &nlmlogvariabletable.Nlmlogvariableentry[i]}
    }
    nlmlogvariabletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nlmlogvariabletable.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry
// A Notification log entry variable.
// 
// Entries appear in this table when there are variables in
// the varbind list of a Notification in nlmLogTable.
type NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // notification_log_mib.NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry_Nlmlogname
    Nlmlogname interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // notification_log_mib.NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry_Nlmlogindex
    Nlmlogindex interface{}

    // This attribute is a key. A monotonically increasing integer, starting at 1
    // for a given nlmLogIndex, for indexing variables within the logged
    // Notification. The type is interface{} with range: 1..4294967295.
    Nlmlogvariableindex interface{}

    // The variable's object identifier. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Nlmlogvariableid interface{}

    // The type of the value.  One and only one of the value objects that follow
    // must be instantiated, based on this type. The type is
    // Nlmlogvariablevaluetype.
    Nlmlogvariablevaluetype interface{}

    // The value when nlmLogVariableType is 'counter32'. The type is interface{}
    // with range: 0..4294967295.
    Nlmlogvariablecounter32Val interface{}

    // The value when nlmLogVariableType is 'unsigned32'. The type is interface{}
    // with range: 0..4294967295.
    Nlmlogvariableunsigned32Val interface{}

    // The value when nlmLogVariableType is 'timeTicks'. The type is interface{}
    // with range: 0..4294967295.
    Nlmlogvariabletimeticksval interface{}

    // The value when nlmLogVariableType is 'integer32'. The type is interface{}
    // with range: -2147483648..2147483647.
    Nlmlogvariableinteger32Val interface{}

    // The value when nlmLogVariableType is 'octetString'. The type is string.
    Nlmlogvariableoctetstringval interface{}

    // The value when nlmLogVariableType is 'ipAddress'. Although this seems to be
    // unfriendly for IPv6, we have to recognize that there are a number of older
    // MIBs that do contain an IPv4 format address, known as IpAddress.  IPv6
    // addresses are represented using TAddress or InetAddress, and so the
    // underlying datatype is OCTET STRING, and their value would be stored in the
    // nlmLogVariableOctetStringVal column. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Nlmlogvariableipaddressval interface{}

    // The value when nlmLogVariableType is 'objectId'. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Nlmlogvariableoidval interface{}

    // The value when nlmLogVariableType is 'counter64'. The type is interface{}
    // with range: 0..18446744073709551615.
    Nlmlogvariablecounter64Val interface{}

    // The value when nlmLogVariableType is 'opaque'. The type is string.
    Nlmlogvariableopaqueval interface{}
}

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetEntityData() *types.CommonEntityData {
    nlmlogvariableentry.EntityData.YFilter = nlmlogvariableentry.YFilter
    nlmlogvariableentry.EntityData.YangName = "nlmLogVariableEntry"
    nlmlogvariableentry.EntityData.BundleName = "cisco_ios_xe"
    nlmlogvariableentry.EntityData.ParentYangName = "nlmLogVariableTable"
    nlmlogvariableentry.EntityData.SegmentPath = "nlmLogVariableEntry" + "[nlmLogName='" + fmt.Sprintf("%v", nlmlogvariableentry.Nlmlogname) + "']" + "[nlmLogIndex='" + fmt.Sprintf("%v", nlmlogvariableentry.Nlmlogindex) + "']" + "[nlmLogVariableIndex='" + fmt.Sprintf("%v", nlmlogvariableentry.Nlmlogvariableindex) + "']"
    nlmlogvariableentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nlmlogvariableentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nlmlogvariableentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nlmlogvariableentry.EntityData.Children = make(map[string]types.YChild)
    nlmlogvariableentry.EntityData.Leafs = make(map[string]types.YLeaf)
    nlmlogvariableentry.EntityData.Leafs["nlmLogName"] = types.YLeaf{"Nlmlogname", nlmlogvariableentry.Nlmlogname}
    nlmlogvariableentry.EntityData.Leafs["nlmLogIndex"] = types.YLeaf{"Nlmlogindex", nlmlogvariableentry.Nlmlogindex}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableIndex"] = types.YLeaf{"Nlmlogvariableindex", nlmlogvariableentry.Nlmlogvariableindex}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableID"] = types.YLeaf{"Nlmlogvariableid", nlmlogvariableentry.Nlmlogvariableid}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableValueType"] = types.YLeaf{"Nlmlogvariablevaluetype", nlmlogvariableentry.Nlmlogvariablevaluetype}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableCounter32Val"] = types.YLeaf{"Nlmlogvariablecounter32Val", nlmlogvariableentry.Nlmlogvariablecounter32Val}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableUnsigned32Val"] = types.YLeaf{"Nlmlogvariableunsigned32Val", nlmlogvariableentry.Nlmlogvariableunsigned32Val}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableTimeTicksVal"] = types.YLeaf{"Nlmlogvariabletimeticksval", nlmlogvariableentry.Nlmlogvariabletimeticksval}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableInteger32Val"] = types.YLeaf{"Nlmlogvariableinteger32Val", nlmlogvariableentry.Nlmlogvariableinteger32Val}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableOctetStringVal"] = types.YLeaf{"Nlmlogvariableoctetstringval", nlmlogvariableentry.Nlmlogvariableoctetstringval}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableIpAddressVal"] = types.YLeaf{"Nlmlogvariableipaddressval", nlmlogvariableentry.Nlmlogvariableipaddressval}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableOidVal"] = types.YLeaf{"Nlmlogvariableoidval", nlmlogvariableentry.Nlmlogvariableoidval}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableCounter64Val"] = types.YLeaf{"Nlmlogvariablecounter64Val", nlmlogvariableentry.Nlmlogvariablecounter64Val}
    nlmlogvariableentry.EntityData.Leafs["nlmLogVariableOpaqueVal"] = types.YLeaf{"Nlmlogvariableopaqueval", nlmlogvariableentry.Nlmlogvariableopaqueval}
    return &(nlmlogvariableentry.EntityData)
}

// NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype represents objects that follow must be instantiated, based on this type.
type NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype string

const (
    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_counter32 NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "counter32"

    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_unsigned32 NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "unsigned32"

    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_timeTicks NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "timeTicks"

    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_integer32 NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "integer32"

    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_ipAddress NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "ipAddress"

    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_octetString NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "octetString"

    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_objectId NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "objectId"

    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_counter64 NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "counter64"

    NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype_opaque NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry_Nlmlogvariablevaluetype = "opaque"
)

