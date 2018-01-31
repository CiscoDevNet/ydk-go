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
    parent types.Entity
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

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetFilter() yfilter.YFilter { return nOTIFICATIONLOGMIB.YFilter }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) SetFilter(yf yfilter.YFilter) { nOTIFICATIONLOGMIB.YFilter = yf }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetGoName(yname string) string {
    if yname == "nlmConfig" { return "Nlmconfig" }
    if yname == "nlmStats" { return "Nlmstats" }
    if yname == "nlmConfigLogTable" { return "Nlmconfiglogtable" }
    if yname == "nlmLogTable" { return "Nlmlogtable" }
    if yname == "nlmLogVariableTable" { return "Nlmlogvariabletable" }
    return ""
}

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetSegmentPath() string {
    return "NOTIFICATION-LOG-MIB:NOTIFICATION-LOG-MIB"
}

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlmConfig" {
        return &nOTIFICATIONLOGMIB.Nlmconfig
    }
    if childYangName == "nlmStats" {
        return &nOTIFICATIONLOGMIB.Nlmstats
    }
    if childYangName == "nlmConfigLogTable" {
        return &nOTIFICATIONLOGMIB.Nlmconfiglogtable
    }
    if childYangName == "nlmLogTable" {
        return &nOTIFICATIONLOGMIB.Nlmlogtable
    }
    if childYangName == "nlmLogVariableTable" {
        return &nOTIFICATIONLOGMIB.Nlmlogvariabletable
    }
    return nil
}

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nlmConfig"] = &nOTIFICATIONLOGMIB.Nlmconfig
    children["nlmStats"] = &nOTIFICATIONLOGMIB.Nlmstats
    children["nlmConfigLogTable"] = &nOTIFICATIONLOGMIB.Nlmconfiglogtable
    children["nlmLogTable"] = &nOTIFICATIONLOGMIB.Nlmlogtable
    children["nlmLogVariableTable"] = &nOTIFICATIONLOGMIB.Nlmlogvariabletable
    return children
}

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetBundleName() string { return "cisco_ios_xe" }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetYangName() string { return "NOTIFICATION-LOG-MIB" }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) SetParent(parent types.Entity) { nOTIFICATIONLOGMIB.parent = parent }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetParent() types.Entity { return nOTIFICATIONLOGMIB.parent }

func (nOTIFICATIONLOGMIB *NOTIFICATIONLOGMIB) GetParentYangName() string { return "NOTIFICATION-LOG-MIB" }

// NOTIFICATIONLOGMIB_Nlmconfig
type NOTIFICATIONLOGMIB_Nlmconfig struct {
    parent types.Entity
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

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetFilter() yfilter.YFilter { return nlmconfig.YFilter }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) SetFilter(yf yfilter.YFilter) { nlmconfig.YFilter = yf }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetGoName(yname string) string {
    if yname == "nlmConfigGlobalEntryLimit" { return "Nlmconfigglobalentrylimit" }
    if yname == "nlmConfigGlobalAgeOut" { return "Nlmconfigglobalageout" }
    return ""
}

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetSegmentPath() string {
    return "nlmConfig"
}

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nlmConfigGlobalEntryLimit"] = nlmconfig.Nlmconfigglobalentrylimit
    leafs["nlmConfigGlobalAgeOut"] = nlmconfig.Nlmconfigglobalageout
    return leafs
}

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetBundleName() string { return "cisco_ios_xe" }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetYangName() string { return "nlmConfig" }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) SetParent(parent types.Entity) { nlmconfig.parent = parent }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetParent() types.Entity { return nlmconfig.parent }

func (nlmconfig *NOTIFICATIONLOGMIB_Nlmconfig) GetParentYangName() string { return "NOTIFICATION-LOG-MIB" }

// NOTIFICATIONLOGMIB_Nlmstats
type NOTIFICATIONLOGMIB_Nlmstats struct {
    parent types.Entity
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

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetFilter() yfilter.YFilter { return nlmstats.YFilter }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) SetFilter(yf yfilter.YFilter) { nlmstats.YFilter = yf }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetGoName(yname string) string {
    if yname == "nlmStatsGlobalNotificationsLogged" { return "Nlmstatsglobalnotificationslogged" }
    if yname == "nlmStatsGlobalNotificationsBumped" { return "Nlmstatsglobalnotificationsbumped" }
    return ""
}

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetSegmentPath() string {
    return "nlmStats"
}

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nlmStatsGlobalNotificationsLogged"] = nlmstats.Nlmstatsglobalnotificationslogged
    leafs["nlmStatsGlobalNotificationsBumped"] = nlmstats.Nlmstatsglobalnotificationsbumped
    return leafs
}

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetBundleName() string { return "cisco_ios_xe" }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetYangName() string { return "nlmStats" }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) SetParent(parent types.Entity) { nlmstats.parent = parent }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetParent() types.Entity { return nlmstats.parent }

func (nlmstats *NOTIFICATIONLOGMIB_Nlmstats) GetParentYangName() string { return "NOTIFICATION-LOG-MIB" }

// NOTIFICATIONLOGMIB_Nlmconfiglogtable
// A table of logging control entries.
type NOTIFICATIONLOGMIB_Nlmconfiglogtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A logging control entry.  Depending on the entry's storage type entries may
    // be supplied by the system or created and deleted by applications using
    // nlmConfigLogEntryStatus. The type is slice of
    // NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry.
    Nlmconfiglogentry []NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry
}

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetFilter() yfilter.YFilter { return nlmconfiglogtable.YFilter }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) SetFilter(yf yfilter.YFilter) { nlmconfiglogtable.YFilter = yf }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetGoName(yname string) string {
    if yname == "nlmConfigLogEntry" { return "Nlmconfiglogentry" }
    return ""
}

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetSegmentPath() string {
    return "nlmConfigLogTable"
}

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlmConfigLogEntry" {
        for _, c := range nlmconfiglogtable.Nlmconfiglogentry {
            if nlmconfiglogtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry{}
        nlmconfiglogtable.Nlmconfiglogentry = append(nlmconfiglogtable.Nlmconfiglogentry, child)
        return &nlmconfiglogtable.Nlmconfiglogentry[len(nlmconfiglogtable.Nlmconfiglogentry)-1]
    }
    return nil
}

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nlmconfiglogtable.Nlmconfiglogentry {
        children[nlmconfiglogtable.Nlmconfiglogentry[i].GetSegmentPath()] = &nlmconfiglogtable.Nlmconfiglogentry[i]
    }
    return children
}

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetBundleName() string { return "cisco_ios_xe" }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetYangName() string { return "nlmConfigLogTable" }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) SetParent(parent types.Entity) { nlmconfiglogtable.parent = parent }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetParent() types.Entity { return nlmconfiglogtable.parent }

func (nlmconfiglogtable *NOTIFICATIONLOGMIB_Nlmconfiglogtable) GetParentYangName() string { return "NOTIFICATION-LOG-MIB" }

// NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry
// A logging control entry.  Depending on the entry's storage type
// entries may be supplied by the system or created and deleted by
// applications using nlmConfigLogEntryStatus.
type NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry struct {
    parent types.Entity
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

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetFilter() yfilter.YFilter { return nlmconfiglogentry.YFilter }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) SetFilter(yf yfilter.YFilter) { nlmconfiglogentry.YFilter = yf }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetGoName(yname string) string {
    if yname == "nlmLogName" { return "Nlmlogname" }
    if yname == "nlmConfigLogFilterName" { return "Nlmconfiglogfiltername" }
    if yname == "nlmConfigLogEntryLimit" { return "Nlmconfiglogentrylimit" }
    if yname == "nlmConfigLogAdminStatus" { return "Nlmconfiglogadminstatus" }
    if yname == "nlmConfigLogOperStatus" { return "Nlmconfiglogoperstatus" }
    if yname == "nlmConfigLogStorageType" { return "Nlmconfiglogstoragetype" }
    if yname == "nlmConfigLogEntryStatus" { return "Nlmconfiglogentrystatus" }
    if yname == "nlmStatsLogNotificationsLogged" { return "Nlmstatslognotificationslogged" }
    if yname == "nlmStatsLogNotificationsBumped" { return "Nlmstatslognotificationsbumped" }
    return ""
}

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetSegmentPath() string {
    return "nlmConfigLogEntry" + "[nlmLogName='" + fmt.Sprintf("%v", nlmconfiglogentry.Nlmlogname) + "']"
}

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nlmLogName"] = nlmconfiglogentry.Nlmlogname
    leafs["nlmConfigLogFilterName"] = nlmconfiglogentry.Nlmconfiglogfiltername
    leafs["nlmConfigLogEntryLimit"] = nlmconfiglogentry.Nlmconfiglogentrylimit
    leafs["nlmConfigLogAdminStatus"] = nlmconfiglogentry.Nlmconfiglogadminstatus
    leafs["nlmConfigLogOperStatus"] = nlmconfiglogentry.Nlmconfiglogoperstatus
    leafs["nlmConfigLogStorageType"] = nlmconfiglogentry.Nlmconfiglogstoragetype
    leafs["nlmConfigLogEntryStatus"] = nlmconfiglogentry.Nlmconfiglogentrystatus
    leafs["nlmStatsLogNotificationsLogged"] = nlmconfiglogentry.Nlmstatslognotificationslogged
    leafs["nlmStatsLogNotificationsBumped"] = nlmconfiglogentry.Nlmstatslognotificationsbumped
    return leafs
}

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetBundleName() string { return "cisco_ios_xe" }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetYangName() string { return "nlmConfigLogEntry" }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) SetParent(parent types.Entity) { nlmconfiglogentry.parent = parent }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetParent() types.Entity { return nlmconfiglogentry.parent }

func (nlmconfiglogentry *NOTIFICATIONLOGMIB_Nlmconfiglogtable_Nlmconfiglogentry) GetParentYangName() string { return "nlmConfigLogTable" }

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
    parent types.Entity
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

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetFilter() yfilter.YFilter { return nlmlogtable.YFilter }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) SetFilter(yf yfilter.YFilter) { nlmlogtable.YFilter = yf }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetGoName(yname string) string {
    if yname == "nlmLogEntry" { return "Nlmlogentry" }
    return ""
}

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetSegmentPath() string {
    return "nlmLogTable"
}

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlmLogEntry" {
        for _, c := range nlmlogtable.Nlmlogentry {
            if nlmlogtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry{}
        nlmlogtable.Nlmlogentry = append(nlmlogtable.Nlmlogentry, child)
        return &nlmlogtable.Nlmlogentry[len(nlmlogtable.Nlmlogentry)-1]
    }
    return nil
}

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nlmlogtable.Nlmlogentry {
        children[nlmlogtable.Nlmlogentry[i].GetSegmentPath()] = &nlmlogtable.Nlmlogentry[i]
    }
    return children
}

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetBundleName() string { return "cisco_ios_xe" }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetYangName() string { return "nlmLogTable" }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) SetParent(parent types.Entity) { nlmlogtable.parent = parent }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetParent() types.Entity { return nlmlogtable.parent }

func (nlmlogtable *NOTIFICATIONLOGMIB_Nlmlogtable) GetParentYangName() string { return "NOTIFICATION-LOG-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Nlmlognotificationid interface{}
}

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetFilter() yfilter.YFilter { return nlmlogentry.YFilter }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) SetFilter(yf yfilter.YFilter) { nlmlogentry.YFilter = yf }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetGoName(yname string) string {
    if yname == "nlmLogName" { return "Nlmlogname" }
    if yname == "nlmLogIndex" { return "Nlmlogindex" }
    if yname == "nlmLogTime" { return "Nlmlogtime" }
    if yname == "nlmLogDateAndTime" { return "Nlmlogdateandtime" }
    if yname == "nlmLogEngineID" { return "Nlmlogengineid" }
    if yname == "nlmLogEngineTAddress" { return "Nlmlogenginetaddress" }
    if yname == "nlmLogEngineTDomain" { return "Nlmlogenginetdomain" }
    if yname == "nlmLogContextEngineID" { return "Nlmlogcontextengineid" }
    if yname == "nlmLogContextName" { return "Nlmlogcontextname" }
    if yname == "nlmLogNotificationID" { return "Nlmlognotificationid" }
    return ""
}

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetSegmentPath() string {
    return "nlmLogEntry" + "[nlmLogName='" + fmt.Sprintf("%v", nlmlogentry.Nlmlogname) + "']" + "[nlmLogIndex='" + fmt.Sprintf("%v", nlmlogentry.Nlmlogindex) + "']"
}

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nlmLogName"] = nlmlogentry.Nlmlogname
    leafs["nlmLogIndex"] = nlmlogentry.Nlmlogindex
    leafs["nlmLogTime"] = nlmlogentry.Nlmlogtime
    leafs["nlmLogDateAndTime"] = nlmlogentry.Nlmlogdateandtime
    leafs["nlmLogEngineID"] = nlmlogentry.Nlmlogengineid
    leafs["nlmLogEngineTAddress"] = nlmlogentry.Nlmlogenginetaddress
    leafs["nlmLogEngineTDomain"] = nlmlogentry.Nlmlogenginetdomain
    leafs["nlmLogContextEngineID"] = nlmlogentry.Nlmlogcontextengineid
    leafs["nlmLogContextName"] = nlmlogentry.Nlmlogcontextname
    leafs["nlmLogNotificationID"] = nlmlogentry.Nlmlognotificationid
    return leafs
}

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetBundleName() string { return "cisco_ios_xe" }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetYangName() string { return "nlmLogEntry" }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) SetParent(parent types.Entity) { nlmlogentry.parent = parent }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetParent() types.Entity { return nlmlogentry.parent }

func (nlmlogentry *NOTIFICATIONLOGMIB_Nlmlogtable_Nlmlogentry) GetParentYangName() string { return "nlmLogTable" }

// NOTIFICATIONLOGMIB_Nlmlogvariabletable
// A table of variables to go with Notification log entries.
type NOTIFICATIONLOGMIB_Nlmlogvariabletable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Notification log entry variable.  Entries appear in this table when there
    // are variables in the varbind list of a Notification in nlmLogTable. The
    // type is slice of
    // NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry.
    Nlmlogvariableentry []NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry
}

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetFilter() yfilter.YFilter { return nlmlogvariabletable.YFilter }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) SetFilter(yf yfilter.YFilter) { nlmlogvariabletable.YFilter = yf }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetGoName(yname string) string {
    if yname == "nlmLogVariableEntry" { return "Nlmlogvariableentry" }
    return ""
}

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetSegmentPath() string {
    return "nlmLogVariableTable"
}

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nlmLogVariableEntry" {
        for _, c := range nlmlogvariabletable.Nlmlogvariableentry {
            if nlmlogvariabletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry{}
        nlmlogvariabletable.Nlmlogvariableentry = append(nlmlogvariabletable.Nlmlogvariableentry, child)
        return &nlmlogvariabletable.Nlmlogvariableentry[len(nlmlogvariabletable.Nlmlogvariableentry)-1]
    }
    return nil
}

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nlmlogvariabletable.Nlmlogvariableentry {
        children[nlmlogvariabletable.Nlmlogvariableentry[i].GetSegmentPath()] = &nlmlogvariabletable.Nlmlogvariableentry[i]
    }
    return children
}

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetBundleName() string { return "cisco_ios_xe" }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetYangName() string { return "nlmLogVariableTable" }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) SetParent(parent types.Entity) { nlmlogvariabletable.parent = parent }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetParent() types.Entity { return nlmlogvariabletable.parent }

func (nlmlogvariabletable *NOTIFICATIONLOGMIB_Nlmlogvariabletable) GetParentYangName() string { return "NOTIFICATION-LOG-MIB" }

// NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry
// A Notification log entry variable.
// 
// Entries appear in this table when there are variables in
// the varbind list of a Notification in nlmLogTable.
type NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Nlmlogvariableipaddressval interface{}

    // The value when nlmLogVariableType is 'objectId'. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Nlmlogvariableoidval interface{}

    // The value when nlmLogVariableType is 'counter64'. The type is interface{}
    // with range: 0..18446744073709551615.
    Nlmlogvariablecounter64Val interface{}

    // The value when nlmLogVariableType is 'opaque'. The type is string.
    Nlmlogvariableopaqueval interface{}
}

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetFilter() yfilter.YFilter { return nlmlogvariableentry.YFilter }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) SetFilter(yf yfilter.YFilter) { nlmlogvariableentry.YFilter = yf }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetGoName(yname string) string {
    if yname == "nlmLogName" { return "Nlmlogname" }
    if yname == "nlmLogIndex" { return "Nlmlogindex" }
    if yname == "nlmLogVariableIndex" { return "Nlmlogvariableindex" }
    if yname == "nlmLogVariableID" { return "Nlmlogvariableid" }
    if yname == "nlmLogVariableValueType" { return "Nlmlogvariablevaluetype" }
    if yname == "nlmLogVariableCounter32Val" { return "Nlmlogvariablecounter32Val" }
    if yname == "nlmLogVariableUnsigned32Val" { return "Nlmlogvariableunsigned32Val" }
    if yname == "nlmLogVariableTimeTicksVal" { return "Nlmlogvariabletimeticksval" }
    if yname == "nlmLogVariableInteger32Val" { return "Nlmlogvariableinteger32Val" }
    if yname == "nlmLogVariableOctetStringVal" { return "Nlmlogvariableoctetstringval" }
    if yname == "nlmLogVariableIpAddressVal" { return "Nlmlogvariableipaddressval" }
    if yname == "nlmLogVariableOidVal" { return "Nlmlogvariableoidval" }
    if yname == "nlmLogVariableCounter64Val" { return "Nlmlogvariablecounter64Val" }
    if yname == "nlmLogVariableOpaqueVal" { return "Nlmlogvariableopaqueval" }
    return ""
}

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetSegmentPath() string {
    return "nlmLogVariableEntry" + "[nlmLogName='" + fmt.Sprintf("%v", nlmlogvariableentry.Nlmlogname) + "']" + "[nlmLogIndex='" + fmt.Sprintf("%v", nlmlogvariableentry.Nlmlogindex) + "']" + "[nlmLogVariableIndex='" + fmt.Sprintf("%v", nlmlogvariableentry.Nlmlogvariableindex) + "']"
}

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nlmLogName"] = nlmlogvariableentry.Nlmlogname
    leafs["nlmLogIndex"] = nlmlogvariableentry.Nlmlogindex
    leafs["nlmLogVariableIndex"] = nlmlogvariableentry.Nlmlogvariableindex
    leafs["nlmLogVariableID"] = nlmlogvariableentry.Nlmlogvariableid
    leafs["nlmLogVariableValueType"] = nlmlogvariableentry.Nlmlogvariablevaluetype
    leafs["nlmLogVariableCounter32Val"] = nlmlogvariableentry.Nlmlogvariablecounter32Val
    leafs["nlmLogVariableUnsigned32Val"] = nlmlogvariableentry.Nlmlogvariableunsigned32Val
    leafs["nlmLogVariableTimeTicksVal"] = nlmlogvariableentry.Nlmlogvariabletimeticksval
    leafs["nlmLogVariableInteger32Val"] = nlmlogvariableentry.Nlmlogvariableinteger32Val
    leafs["nlmLogVariableOctetStringVal"] = nlmlogvariableentry.Nlmlogvariableoctetstringval
    leafs["nlmLogVariableIpAddressVal"] = nlmlogvariableentry.Nlmlogvariableipaddressval
    leafs["nlmLogVariableOidVal"] = nlmlogvariableentry.Nlmlogvariableoidval
    leafs["nlmLogVariableCounter64Val"] = nlmlogvariableentry.Nlmlogvariablecounter64Val
    leafs["nlmLogVariableOpaqueVal"] = nlmlogvariableentry.Nlmlogvariableopaqueval
    return leafs
}

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetBundleName() string { return "cisco_ios_xe" }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetYangName() string { return "nlmLogVariableEntry" }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) SetParent(parent types.Entity) { nlmlogvariableentry.parent = parent }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetParent() types.Entity { return nlmlogvariableentry.parent }

func (nlmlogvariableentry *NOTIFICATIONLOGMIB_Nlmlogvariabletable_Nlmlogvariableentry) GetParentYangName() string { return "nlmLogVariableTable" }

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

