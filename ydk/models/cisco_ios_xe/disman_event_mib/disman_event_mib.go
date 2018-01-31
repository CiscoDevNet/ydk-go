// The MIB module for defining event triggers and actions
// for network management purposes.
package disman_event_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package disman_event_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DISMAN-EVENT-MIB DISMAN-EVENT-MIB}", reflect.TypeOf(DISMANEVENTMIB{}))
    ydk.RegisterEntity("DISMAN-EVENT-MIB:DISMAN-EVENT-MIB", reflect.TypeOf(DISMANEVENTMIB{}))
}

// FailureReason represents                         the previous one completed
type FailureReason string

const (
    FailureReason_sampleOverrun FailureReason = "sampleOverrun"

    FailureReason_badType FailureReason = "badType"

    FailureReason_noResponse FailureReason = "noResponse"

    FailureReason_destinationUnreachable FailureReason = "destinationUnreachable"

    FailureReason_badDestination FailureReason = "badDestination"

    FailureReason_localResourceLack FailureReason = "localResourceLack"

    FailureReason_noError FailureReason = "noError"

    FailureReason_tooBig FailureReason = "tooBig"

    FailureReason_noSuchName FailureReason = "noSuchName"

    FailureReason_badValue FailureReason = "badValue"

    FailureReason_readOnly FailureReason = "readOnly"

    FailureReason_genErr FailureReason = "genErr"

    FailureReason_noAccess FailureReason = "noAccess"

    FailureReason_wrongType FailureReason = "wrongType"

    FailureReason_wrongLength FailureReason = "wrongLength"

    FailureReason_wrongEncoding FailureReason = "wrongEncoding"

    FailureReason_wrongValue FailureReason = "wrongValue"

    FailureReason_noCreation FailureReason = "noCreation"

    FailureReason_inconsistentValue FailureReason = "inconsistentValue"

    FailureReason_resourceUnavailable FailureReason = "resourceUnavailable"

    FailureReason_commitFailed FailureReason = "commitFailed"

    FailureReason_undoFailed FailureReason = "undoFailed"

    FailureReason_authorizationError FailureReason = "authorizationError"

    FailureReason_notWritable FailureReason = "notWritable"

    FailureReason_inconsistentName FailureReason = "inconsistentName"
)

// DISMANEVENTMIB
type DISMANEVENTMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Mteresource DISMANEVENTMIB_Mteresource

    
    Mtetrigger DISMANEVENTMIB_Mtetrigger

    
    Mteevent DISMANEVENTMIB_Mteevent

    // A table of management event trigger information.
    Mtetriggertable DISMANEVENTMIB_Mtetriggertable

    // A table of management event trigger information for delta sampling.
    Mtetriggerdeltatable DISMANEVENTMIB_Mtetriggerdeltatable

    // A table of management event trigger information for existence triggers.
    Mtetriggerexistencetable DISMANEVENTMIB_Mtetriggerexistencetable

    // A table of management event trigger information for boolean triggers.
    Mtetriggerbooleantable DISMANEVENTMIB_Mtetriggerbooleantable

    // A table of management event trigger information for threshold triggers.
    Mtetriggerthresholdtable DISMANEVENTMIB_Mtetriggerthresholdtable

    // A table of objects that can be added to notifications based on the trigger,
    // trigger test, or event, as pointed to by entries in those tables.
    Mteobjectstable DISMANEVENTMIB_Mteobjectstable

    // A table of management event action information.
    Mteeventtable DISMANEVENTMIB_Mteeventtable

    // A table of information about notifications to be sent as a consequence of
    // management events.
    Mteeventnotificationtable DISMANEVENTMIB_Mteeventnotificationtable

    // A table of management event action information.
    Mteeventsettable DISMANEVENTMIB_Mteeventsettable
}

func (dISMANEVENTMIB *DISMANEVENTMIB) GetFilter() yfilter.YFilter { return dISMANEVENTMIB.YFilter }

func (dISMANEVENTMIB *DISMANEVENTMIB) SetFilter(yf yfilter.YFilter) { dISMANEVENTMIB.YFilter = yf }

func (dISMANEVENTMIB *DISMANEVENTMIB) GetGoName(yname string) string {
    if yname == "mteResource" { return "Mteresource" }
    if yname == "mteTrigger" { return "Mtetrigger" }
    if yname == "mteEvent" { return "Mteevent" }
    if yname == "mteTriggerTable" { return "Mtetriggertable" }
    if yname == "mteTriggerDeltaTable" { return "Mtetriggerdeltatable" }
    if yname == "mteTriggerExistenceTable" { return "Mtetriggerexistencetable" }
    if yname == "mteTriggerBooleanTable" { return "Mtetriggerbooleantable" }
    if yname == "mteTriggerThresholdTable" { return "Mtetriggerthresholdtable" }
    if yname == "mteObjectsTable" { return "Mteobjectstable" }
    if yname == "mteEventTable" { return "Mteeventtable" }
    if yname == "mteEventNotificationTable" { return "Mteeventnotificationtable" }
    if yname == "mteEventSetTable" { return "Mteeventsettable" }
    return ""
}

func (dISMANEVENTMIB *DISMANEVENTMIB) GetSegmentPath() string {
    return "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB"
}

func (dISMANEVENTMIB *DISMANEVENTMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteResource" {
        return &dISMANEVENTMIB.Mteresource
    }
    if childYangName == "mteTrigger" {
        return &dISMANEVENTMIB.Mtetrigger
    }
    if childYangName == "mteEvent" {
        return &dISMANEVENTMIB.Mteevent
    }
    if childYangName == "mteTriggerTable" {
        return &dISMANEVENTMIB.Mtetriggertable
    }
    if childYangName == "mteTriggerDeltaTable" {
        return &dISMANEVENTMIB.Mtetriggerdeltatable
    }
    if childYangName == "mteTriggerExistenceTable" {
        return &dISMANEVENTMIB.Mtetriggerexistencetable
    }
    if childYangName == "mteTriggerBooleanTable" {
        return &dISMANEVENTMIB.Mtetriggerbooleantable
    }
    if childYangName == "mteTriggerThresholdTable" {
        return &dISMANEVENTMIB.Mtetriggerthresholdtable
    }
    if childYangName == "mteObjectsTable" {
        return &dISMANEVENTMIB.Mteobjectstable
    }
    if childYangName == "mteEventTable" {
        return &dISMANEVENTMIB.Mteeventtable
    }
    if childYangName == "mteEventNotificationTable" {
        return &dISMANEVENTMIB.Mteeventnotificationtable
    }
    if childYangName == "mteEventSetTable" {
        return &dISMANEVENTMIB.Mteeventsettable
    }
    return nil
}

func (dISMANEVENTMIB *DISMANEVENTMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mteResource"] = &dISMANEVENTMIB.Mteresource
    children["mteTrigger"] = &dISMANEVENTMIB.Mtetrigger
    children["mteEvent"] = &dISMANEVENTMIB.Mteevent
    children["mteTriggerTable"] = &dISMANEVENTMIB.Mtetriggertable
    children["mteTriggerDeltaTable"] = &dISMANEVENTMIB.Mtetriggerdeltatable
    children["mteTriggerExistenceTable"] = &dISMANEVENTMIB.Mtetriggerexistencetable
    children["mteTriggerBooleanTable"] = &dISMANEVENTMIB.Mtetriggerbooleantable
    children["mteTriggerThresholdTable"] = &dISMANEVENTMIB.Mtetriggerthresholdtable
    children["mteObjectsTable"] = &dISMANEVENTMIB.Mteobjectstable
    children["mteEventTable"] = &dISMANEVENTMIB.Mteeventtable
    children["mteEventNotificationTable"] = &dISMANEVENTMIB.Mteeventnotificationtable
    children["mteEventSetTable"] = &dISMANEVENTMIB.Mteeventsettable
    return children
}

func (dISMANEVENTMIB *DISMANEVENTMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dISMANEVENTMIB *DISMANEVENTMIB) GetBundleName() string { return "cisco_ios_xe" }

func (dISMANEVENTMIB *DISMANEVENTMIB) GetYangName() string { return "DISMAN-EVENT-MIB" }

func (dISMANEVENTMIB *DISMANEVENTMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dISMANEVENTMIB *DISMANEVENTMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dISMANEVENTMIB *DISMANEVENTMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dISMANEVENTMIB *DISMANEVENTMIB) SetParent(parent types.Entity) { dISMANEVENTMIB.parent = parent }

func (dISMANEVENTMIB *DISMANEVENTMIB) GetParent() types.Entity { return dISMANEVENTMIB.parent }

func (dISMANEVENTMIB *DISMANEVENTMIB) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mteresource
type DISMANEVENTMIB_Mteresource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The minimum mteTriggerFrequency this system will accept.  A system may use
    // the larger values of this minimum to lessen the impact of constant
    // sampling.  For larger sampling intervals the system samples less often and
    // suffers less overhead.  This object provides a way to enforce such lower
    // overhead for all triggers created after it is set.  Unless explicitly
    // resource limited, a system's value for this object SHOULD be 1, allowing as
    // small as a 1 second interval for ongoing trigger sampling.  Changing this
    // value will not invalidate an existing setting of mteTriggerFrequency. The
    // type is interface{} with range: 1..2147483647. Units are seconds.
    Mteresourcesampleminimum interface{}

    // The maximum number of instance entries this system will support for
    // sampling.  These are the entries that maintain state, one for each instance
    // of each sampled object as selected by mteTriggerValueID.  Note that
    // wildcarded objects result in multiple instances of this state.  A value of
    // 0 indicates no preset limit, that is, the limit is dynamic based on system
    // operation and resources.  Unless explicitly resource limited, a system's
    // value for this object SHOULD be 0.  Changing this value will not eliminate
    // or inhibit existing sample state but could prevent allocation of additional
    // state information. The type is interface{} with range: 0..4294967295. Units
    // are instances.
    Mteresourcesampleinstancemaximum interface{}

    // The number of currently active instance entries as defined for
    // mteResourceSampleInstanceMaximum. The type is interface{} with range:
    // 0..4294967295. Units are instances.
    Mteresourcesampleinstances interface{}

    // The highest value of mteResourceSampleInstances that has occurred since
    // initialization of the management system. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    Mteresourcesampleinstanceshigh interface{}

    // The number of times this system could not take a new sample because that
    // allocation would have exceeded the limit set by
    // mteResourceSampleInstanceMaximum. The type is interface{} with range:
    // 0..4294967295. Units are instances.
    Mteresourcesampleinstancelacks interface{}
}

func (mteresource *DISMANEVENTMIB_Mteresource) GetFilter() yfilter.YFilter { return mteresource.YFilter }

func (mteresource *DISMANEVENTMIB_Mteresource) SetFilter(yf yfilter.YFilter) { mteresource.YFilter = yf }

func (mteresource *DISMANEVENTMIB_Mteresource) GetGoName(yname string) string {
    if yname == "mteResourceSampleMinimum" { return "Mteresourcesampleminimum" }
    if yname == "mteResourceSampleInstanceMaximum" { return "Mteresourcesampleinstancemaximum" }
    if yname == "mteResourceSampleInstances" { return "Mteresourcesampleinstances" }
    if yname == "mteResourceSampleInstancesHigh" { return "Mteresourcesampleinstanceshigh" }
    if yname == "mteResourceSampleInstanceLacks" { return "Mteresourcesampleinstancelacks" }
    return ""
}

func (mteresource *DISMANEVENTMIB_Mteresource) GetSegmentPath() string {
    return "mteResource"
}

func (mteresource *DISMANEVENTMIB_Mteresource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mteresource *DISMANEVENTMIB_Mteresource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mteresource *DISMANEVENTMIB_Mteresource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteResourceSampleMinimum"] = mteresource.Mteresourcesampleminimum
    leafs["mteResourceSampleInstanceMaximum"] = mteresource.Mteresourcesampleinstancemaximum
    leafs["mteResourceSampleInstances"] = mteresource.Mteresourcesampleinstances
    leafs["mteResourceSampleInstancesHigh"] = mteresource.Mteresourcesampleinstanceshigh
    leafs["mteResourceSampleInstanceLacks"] = mteresource.Mteresourcesampleinstancelacks
    return leafs
}

func (mteresource *DISMANEVENTMIB_Mteresource) GetBundleName() string { return "cisco_ios_xe" }

func (mteresource *DISMANEVENTMIB_Mteresource) GetYangName() string { return "mteResource" }

func (mteresource *DISMANEVENTMIB_Mteresource) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteresource *DISMANEVENTMIB_Mteresource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteresource *DISMANEVENTMIB_Mteresource) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteresource *DISMANEVENTMIB_Mteresource) SetParent(parent types.Entity) { mteresource.parent = parent }

func (mteresource *DISMANEVENTMIB_Mteresource) GetParent() types.Entity { return mteresource.parent }

func (mteresource *DISMANEVENTMIB_Mteresource) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mtetrigger
type DISMANEVENTMIB_Mtetrigger struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of times an attempt to check for a trigger condition has failed.
    // This counts individually for each attempt in a group of targets or each
    // attempt for a  wildcarded object. The type is interface{} with range:
    // 0..4294967295. Units are failures.
    Mtetriggerfailures interface{}
}

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetFilter() yfilter.YFilter { return mtetrigger.YFilter }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) SetFilter(yf yfilter.YFilter) { mtetrigger.YFilter = yf }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetGoName(yname string) string {
    if yname == "mteTriggerFailures" { return "Mtetriggerfailures" }
    return ""
}

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetSegmentPath() string {
    return "mteTrigger"
}

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteTriggerFailures"] = mtetrigger.Mtetriggerfailures
    return leafs
}

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetBundleName() string { return "cisco_ios_xe" }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetYangName() string { return "mteTrigger" }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) SetParent(parent types.Entity) { mtetrigger.parent = parent }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetParent() types.Entity { return mtetrigger.parent }

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mteevent
type DISMANEVENTMIB_Mteevent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of times an attempt to invoke an event has failed.  This counts
    // individually for each attempt in a group of targets or each attempt for a
    // wildcarded trigger object. The type is interface{} with range:
    // 0..4294967295.
    Mteeventfailures interface{}
}

func (mteevent *DISMANEVENTMIB_Mteevent) GetFilter() yfilter.YFilter { return mteevent.YFilter }

func (mteevent *DISMANEVENTMIB_Mteevent) SetFilter(yf yfilter.YFilter) { mteevent.YFilter = yf }

func (mteevent *DISMANEVENTMIB_Mteevent) GetGoName(yname string) string {
    if yname == "mteEventFailures" { return "Mteeventfailures" }
    return ""
}

func (mteevent *DISMANEVENTMIB_Mteevent) GetSegmentPath() string {
    return "mteEvent"
}

func (mteevent *DISMANEVENTMIB_Mteevent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mteevent *DISMANEVENTMIB_Mteevent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mteevent *DISMANEVENTMIB_Mteevent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteEventFailures"] = mteevent.Mteeventfailures
    return leafs
}

func (mteevent *DISMANEVENTMIB_Mteevent) GetBundleName() string { return "cisco_ios_xe" }

func (mteevent *DISMANEVENTMIB_Mteevent) GetYangName() string { return "mteEvent" }

func (mteevent *DISMANEVENTMIB_Mteevent) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteevent *DISMANEVENTMIB_Mteevent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteevent *DISMANEVENTMIB_Mteevent) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteevent *DISMANEVENTMIB_Mteevent) SetParent(parent types.Entity) { mteevent.parent = parent }

func (mteevent *DISMANEVENTMIB_Mteevent) GetParent() types.Entity { return mteevent.parent }

func (mteevent *DISMANEVENTMIB_Mteevent) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mtetriggertable
// A table of management event trigger information.
type DISMANEVENTMIB_Mtetriggertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single trigger.  Applications create and delete entries
    // using mteTriggerEntryStatus. The type is slice of
    // DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry.
    Mtetriggerentry []DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry
}

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetFilter() yfilter.YFilter { return mtetriggertable.YFilter }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) SetFilter(yf yfilter.YFilter) { mtetriggertable.YFilter = yf }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetGoName(yname string) string {
    if yname == "mteTriggerEntry" { return "Mtetriggerentry" }
    return ""
}

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetSegmentPath() string {
    return "mteTriggerTable"
}

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteTriggerEntry" {
        for _, c := range mtetriggertable.Mtetriggerentry {
            if mtetriggertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry{}
        mtetriggertable.Mtetriggerentry = append(mtetriggertable.Mtetriggerentry, child)
        return &mtetriggertable.Mtetriggerentry[len(mtetriggertable.Mtetriggerentry)-1]
    }
    return nil
}

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mtetriggertable.Mtetriggerentry {
        children[mtetriggertable.Mtetriggerentry[i].GetSegmentPath()] = &mtetriggertable.Mtetriggerentry[i]
    }
    return children
}

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetYangName() string { return "mteTriggerTable" }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) SetParent(parent types.Entity) { mtetriggertable.parent = parent }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetParent() types.Entity { return mtetriggertable.parent }

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry
// Information about a single trigger.  Applications create and
// delete entries using mteTriggerEntryStatus.
type DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The owner of this entry. The exact semantics of
    // this string are subject to the security policy defined by the security
    // administrator. The type is string with length: 0..32.
    Mteowner interface{}

    // This attribute is a key. A locally-unique, administratively assigned name
    // for the trigger within the scope of mteOwner. The type is string with
    // length: 1..32.
    Mtetriggername interface{}

    // A description of the trigger's function and use. The type is string.
    Mtetriggercomment interface{}

    // The type of trigger test to perform.  For 'boolean' and 'threshold'  tests,
    // the object at mteTriggerValueID MUST evaluate to an integer, that is,
    // anything that ends up encoded for transmission (that is, in BER, not ASN.1)
    // as an integer.  For 'existence', the specific test is as selected by
    // mteTriggerExistenceTest.  When an object appears, vanishes or changes
    // value, the trigger fires. If the object's appearance caused the trigger
    // firing, the object MUST vanish before the trigger can be fired again for
    // it, and vice versa. If the trigger fired due to a change in the object's
    // value, it will be fired again on every successive value change for that
    // object.  For 'boolean', the specific test is as selected by
    // mteTriggerBooleanTest.  If the test result is true the trigger fires.  The
    // trigger will not fire again until the value has become false and come back
    // to true.  For 'threshold' the test works as described below for 
    // mteTriggerThresholdStartup, mteTriggerThresholdRising, and
    // mteTriggerThresholdFalling.  Note that combining 'boolean' and 'threshold'
    // tests on the same object may be somewhat redundant. The type is
    // map[string]bool.
    Mtetriggertest interface{}

    // The type of sampling to perform.  An 'absoluteValue' sample requires only a
    // single sample to be meaningful, and is exactly the value of the object at
    // mteTriggerValueID at the sample time.  A 'deltaValue' requires two samples
    // to be meaningful and is thus not available for testing until the second and
    // subsequent samples after the object at mteTriggerValueID is first found to
    // exist.  It is the difference between the two samples.  For unsigned values
    // it is always positive, based on unsigned arithmetic.  For signed values it
    // can be positive or negative.  For SNMP counters to be meaningful they
    // should be sampled as a 'deltaValue'.  For 'deltaValue' mteTriggerDeltaTable
    // contains further parameters.  If only 'existence' is set in mteTriggerTest
    // this object has no meaning. The type is Mtetriggersampletype.
    Mtetriggersampletype interface{}

    // The object identifier of the MIB object to sample to see if the trigger
    // should fire.  This may be wildcarded by truncating all or part of the
    // instance portion, in which case the value is obtained as if with a GetNext
    // function, checking multiple values  if they exist.  If such wildcarding is
    // applied, mteTriggerValueIDWildcard must be 'true' and if not it must be
    // 'false'.  Bad object identifiers or a mismatch between truncating the
    // identifier and the value of mteTriggerValueIDWildcard result in operation
    // as one would expect when providing the wrong identifier to a Get or GetNext
    // operation.  The Get will fail or get the wrong object.  The GetNext will
    // indeed get whatever is next, proceeding until it runs past the initial part
    // of the identifier and perhaps many unintended objects for confusing
    // results.  If the value syntax of those objects is not usable, that results
    // in a 'badType' error that terminates the scan.  Each instance that fills
    // the wildcard is independent of any additional instances, that is,
    // wildcarded objects operate as if there were a separate table entry for each
    // instance that fills the wildcard without having to actually predict all
    // possible instances ahead of time. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Mtetriggervalueid interface{}

    // Control for whether mteTriggerValueID is to be treated as fully-specified
    // or wildcarded, with 'true' indicating wildcard. The type is bool.
    Mtetriggervalueidwildcard interface{}

    // The tag for the target(s) from which to obtain the condition for a trigger
    // check.  A length of 0 indicates the local system.  In this case, access to
    // the objects indicated by mteTriggerValueID is under the security
    // credentials of the requester that set mteTriggerEntryStatus to 'active'. 
    // Those credentials are the input parameters for isAccessAllowed from the
    // Architecture for Describing SNMP Management Frameworks.  Otherwise access
    // rights are checked according to the security  parameters resulting from the
    // tag. The type is string.
    Mtetriggertargettag interface{}

    // The management context from which to obtain mteTriggerValueID.  This may be
    // wildcarded by leaving characters off the end.  For example use 'Repeater'
    // to wildcard to 'Repeater1', 'Repeater2', 'Repeater-999.87b', and so on.  To
    // indicate such wildcarding is intended, mteTriggerContextNameWildcard must
    // be 'true'.  Each instance that fills the wildcard is independent of any
    // additional instances, that is, wildcarded objects operate as if there were
    // a separate table entry for each instance that fills the wildcard without
    // having to actually predict all possible instances ahead of time.  Operation
    // of this feature assumes that the local system has a list of available
    // contexts against which to apply the wildcard.  If the objects are being
    // read from the local system, this is clearly the system's own list of
    // contexts. For a remote system a local version of such a list is not defined
    // by any current standard and may not be available, so this function MAY not
    // be supported. The type is string.
    Mtetriggercontextname interface{}

    // Control for whether mteTriggerContextName is to be treated as
    // fully-specified or wildcarded, with 'true' indicating wildcard. The type is
    // bool.
    Mtetriggercontextnamewildcard interface{}

    // The number of seconds to wait between trigger samples.  To encourage
    // consistency in sampling, the interval is measured from the beginning of one
    // check to the beginning of the next and the timer is restarted immediately
    // when it expires, not when the check completes.  If the next sample begins
    // before the previous one completed the system may either attempt to make the
    // check or treat this as an error condition with the error 'sampleOverrun'. 
    // A frequency of 0 indicates instantaneous recognition of the condition. 
    // This is not possible in many cases, but may be supported in cases where it
    // makes sense and the system is able to do so.  This feature allows the MIB
    // to be used in implementations where such interrupt-driven behavior is
    // possible and is not likely to be supported for all MIB objects even then
    // since such sampling generally has to be tightly integrated into low-level
    // code.  Systems that can support this SHOULD document those cases where it
    // can be used.  In cases where it can not, setting this object to 0 should be
    // disallowed. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Mtetriggerfrequency interface{}

    // To go with mteTriggerObjects, the mteOwner of a group of objects from
    // mteObjectsTable. The type is string with length: 0..32.
    Mtetriggerobjectsowner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable.  These
    // objects are to be added to any Notification resulting from the firing of
    // this trigger.  A list of objects may also be added based on the event or on
    // the value of mteTriggerTest.  A length of 0 indicates no additional
    // objects. The type is string with length: 0..32.
    Mtetriggerobjects interface{}

    // A control to allow a trigger to be configured but not used. When the value
    // is 'false' the trigger is not sampled. The type is bool.
    Mtetriggerenabled interface{}

    // The control that allows creation and deletion of entries. Once made active
    // an entry may not be modified except to delete it. The type is RowStatus.
    Mtetriggerentrystatus interface{}
}

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetFilter() yfilter.YFilter { return mtetriggerentry.YFilter }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) SetFilter(yf yfilter.YFilter) { mtetriggerentry.YFilter = yf }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteTriggerName" { return "Mtetriggername" }
    if yname == "mteTriggerComment" { return "Mtetriggercomment" }
    if yname == "mteTriggerTest" { return "Mtetriggertest" }
    if yname == "mteTriggerSampleType" { return "Mtetriggersampletype" }
    if yname == "mteTriggerValueID" { return "Mtetriggervalueid" }
    if yname == "mteTriggerValueIDWildcard" { return "Mtetriggervalueidwildcard" }
    if yname == "mteTriggerTargetTag" { return "Mtetriggertargettag" }
    if yname == "mteTriggerContextName" { return "Mtetriggercontextname" }
    if yname == "mteTriggerContextNameWildcard" { return "Mtetriggercontextnamewildcard" }
    if yname == "mteTriggerFrequency" { return "Mtetriggerfrequency" }
    if yname == "mteTriggerObjectsOwner" { return "Mtetriggerobjectsowner" }
    if yname == "mteTriggerObjects" { return "Mtetriggerobjects" }
    if yname == "mteTriggerEnabled" { return "Mtetriggerenabled" }
    if yname == "mteTriggerEntryStatus" { return "Mtetriggerentrystatus" }
    return ""
}

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetSegmentPath() string {
    return "mteTriggerEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerentry.Mtetriggername) + "']"
}

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mtetriggerentry.Mteowner
    leafs["mteTriggerName"] = mtetriggerentry.Mtetriggername
    leafs["mteTriggerComment"] = mtetriggerentry.Mtetriggercomment
    leafs["mteTriggerTest"] = mtetriggerentry.Mtetriggertest
    leafs["mteTriggerSampleType"] = mtetriggerentry.Mtetriggersampletype
    leafs["mteTriggerValueID"] = mtetriggerentry.Mtetriggervalueid
    leafs["mteTriggerValueIDWildcard"] = mtetriggerentry.Mtetriggervalueidwildcard
    leafs["mteTriggerTargetTag"] = mtetriggerentry.Mtetriggertargettag
    leafs["mteTriggerContextName"] = mtetriggerentry.Mtetriggercontextname
    leafs["mteTriggerContextNameWildcard"] = mtetriggerentry.Mtetriggercontextnamewildcard
    leafs["mteTriggerFrequency"] = mtetriggerentry.Mtetriggerfrequency
    leafs["mteTriggerObjectsOwner"] = mtetriggerentry.Mtetriggerobjectsowner
    leafs["mteTriggerObjects"] = mtetriggerentry.Mtetriggerobjects
    leafs["mteTriggerEnabled"] = mtetriggerentry.Mtetriggerenabled
    leafs["mteTriggerEntryStatus"] = mtetriggerentry.Mtetriggerentrystatus
    return leafs
}

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetYangName() string { return "mteTriggerEntry" }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) SetParent(parent types.Entity) { mtetriggerentry.parent = parent }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetParent() types.Entity { return mtetriggerentry.parent }

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetParentYangName() string { return "mteTriggerTable" }

// DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggersampletype represents no meaning.
type DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggersampletype string

const (
    DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggersampletype_absoluteValue DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggersampletype = "absoluteValue"

    DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggersampletype_deltaValue DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggersampletype = "deltaValue"
)

// DISMANEVENTMIB_Mtetriggerdeltatable
// A table of management event trigger information for delta
// sampling.
type DISMANEVENTMIB_Mtetriggerdeltatable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single trigger's delta sampling.  Entries automatically
    // exist in this this table for each mteTriggerEntry that has
    // mteTriggerSampleType set to 'deltaValue'. The type is slice of
    // DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry.
    Mtetriggerdeltaentry []DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry
}

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetFilter() yfilter.YFilter { return mtetriggerdeltatable.YFilter }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) SetFilter(yf yfilter.YFilter) { mtetriggerdeltatable.YFilter = yf }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetGoName(yname string) string {
    if yname == "mteTriggerDeltaEntry" { return "Mtetriggerdeltaentry" }
    return ""
}

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetSegmentPath() string {
    return "mteTriggerDeltaTable"
}

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteTriggerDeltaEntry" {
        for _, c := range mtetriggerdeltatable.Mtetriggerdeltaentry {
            if mtetriggerdeltatable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry{}
        mtetriggerdeltatable.Mtetriggerdeltaentry = append(mtetriggerdeltatable.Mtetriggerdeltaentry, child)
        return &mtetriggerdeltatable.Mtetriggerdeltaentry[len(mtetriggerdeltatable.Mtetriggerdeltaentry)-1]
    }
    return nil
}

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mtetriggerdeltatable.Mtetriggerdeltaentry {
        children[mtetriggerdeltatable.Mtetriggerdeltaentry[i].GetSegmentPath()] = &mtetriggerdeltatable.Mtetriggerdeltaentry[i]
    }
    return children
}

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetYangName() string { return "mteTriggerDeltaTable" }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) SetParent(parent types.Entity) { mtetriggerdeltatable.parent = parent }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetParent() types.Entity { return mtetriggerdeltatable.parent }

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry
// Information about a single trigger's delta sampling.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has mteTriggerSampleType set to 'deltaValue'.
type DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggername
    Mtetriggername interface{}

    // The OBJECT IDENTIFIER (OID) of a TimeTicks, TimeStamp, or DateAndTime
    // object that indicates a discontinuity in the value at mteTriggerValueID. 
    // The OID may be for a leaf object (e.g. sysUpTime.0) or may be wildcarded to
    // match mteTriggerValueID.  This object supports normal checking for a
    // discontinuity in a counter.  Note that if this object does not point to
    // sysUpTime discontinuity checking MUST still check sysUpTime for an overall
    // discontinuity.  If the object identified is not accessible the sample
    // attempt is in error, with the error code as from an SNMP request.  Bad
    // object identifiers or a mismatch between truncating the identifier and the
    // value of mteDeltaDiscontinuityIDWildcard result in operation as one would
    // expect when providing the wrong identifier to a Get operation.  The Get
    // will fail or get the wrong object.  If the value syntax of those objects is
    // not usable, that results in an error that terminates the sample with a
    // 'badType' error code. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Mtetriggerdeltadiscontinuityid interface{}

    // Control for whether mteTriggerDeltaDiscontinuityID is to be treated as
    // fully-specified or wildcarded, with 'true' indicating wildcard. Note that
    // the value of this object will be the same as that of the corresponding
    // instance of mteTriggerValueIDWildcard when the corresponding 
    // mteTriggerSampleType is 'deltaValue'. The type is bool.
    Mtetriggerdeltadiscontinuityidwildcard interface{}

    // The value 'timeTicks' indicates the mteTriggerDeltaDiscontinuityID of this
    // row is of syntax TimeTicks.  The value 'timeStamp' indicates syntax
    // TimeStamp. The value 'dateAndTime' indicates syntax DateAndTime. The type
    // is Mtetriggerdeltadiscontinuityidtype.
    Mtetriggerdeltadiscontinuityidtype interface{}
}

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetFilter() yfilter.YFilter { return mtetriggerdeltaentry.YFilter }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) SetFilter(yf yfilter.YFilter) { mtetriggerdeltaentry.YFilter = yf }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteTriggerName" { return "Mtetriggername" }
    if yname == "mteTriggerDeltaDiscontinuityID" { return "Mtetriggerdeltadiscontinuityid" }
    if yname == "mteTriggerDeltaDiscontinuityIDWildcard" { return "Mtetriggerdeltadiscontinuityidwildcard" }
    if yname == "mteTriggerDeltaDiscontinuityIDType" { return "Mtetriggerdeltadiscontinuityidtype" }
    return ""
}

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetSegmentPath() string {
    return "mteTriggerDeltaEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerdeltaentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerdeltaentry.Mtetriggername) + "']"
}

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mtetriggerdeltaentry.Mteowner
    leafs["mteTriggerName"] = mtetriggerdeltaentry.Mtetriggername
    leafs["mteTriggerDeltaDiscontinuityID"] = mtetriggerdeltaentry.Mtetriggerdeltadiscontinuityid
    leafs["mteTriggerDeltaDiscontinuityIDWildcard"] = mtetriggerdeltaentry.Mtetriggerdeltadiscontinuityidwildcard
    leafs["mteTriggerDeltaDiscontinuityIDType"] = mtetriggerdeltaentry.Mtetriggerdeltadiscontinuityidtype
    return leafs
}

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetYangName() string { return "mteTriggerDeltaEntry" }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) SetParent(parent types.Entity) { mtetriggerdeltaentry.parent = parent }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetParent() types.Entity { return mtetriggerdeltaentry.parent }

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetParentYangName() string { return "mteTriggerDeltaTable" }

// DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry_Mtetriggerdeltadiscontinuityidtype represents The value 'dateAndTime' indicates syntax DateAndTime.
type DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry_Mtetriggerdeltadiscontinuityidtype string

const (
    DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry_Mtetriggerdeltadiscontinuityidtype_timeTicks DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry_Mtetriggerdeltadiscontinuityidtype = "timeTicks"

    DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry_Mtetriggerdeltadiscontinuityidtype_timeStamp DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry_Mtetriggerdeltadiscontinuityidtype = "timeStamp"

    DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry_Mtetriggerdeltadiscontinuityidtype_dateAndTime DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry_Mtetriggerdeltadiscontinuityidtype = "dateAndTime"
)

// DISMANEVENTMIB_Mtetriggerexistencetable
// A table of management event trigger information for existence
// triggers.
type DISMANEVENTMIB_Mtetriggerexistencetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single existence trigger.  Entries automatically exist
    // in this this table for each mteTriggerEntry that has 'existence' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry.
    Mtetriggerexistenceentry []DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry
}

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetFilter() yfilter.YFilter { return mtetriggerexistencetable.YFilter }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) SetFilter(yf yfilter.YFilter) { mtetriggerexistencetable.YFilter = yf }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetGoName(yname string) string {
    if yname == "mteTriggerExistenceEntry" { return "Mtetriggerexistenceentry" }
    return ""
}

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetSegmentPath() string {
    return "mteTriggerExistenceTable"
}

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteTriggerExistenceEntry" {
        for _, c := range mtetriggerexistencetable.Mtetriggerexistenceentry {
            if mtetriggerexistencetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry{}
        mtetriggerexistencetable.Mtetriggerexistenceentry = append(mtetriggerexistencetable.Mtetriggerexistenceentry, child)
        return &mtetriggerexistencetable.Mtetriggerexistenceentry[len(mtetriggerexistencetable.Mtetriggerexistenceentry)-1]
    }
    return nil
}

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mtetriggerexistencetable.Mtetriggerexistenceentry {
        children[mtetriggerexistencetable.Mtetriggerexistenceentry[i].GetSegmentPath()] = &mtetriggerexistencetable.Mtetriggerexistenceentry[i]
    }
    return children
}

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetYangName() string { return "mteTriggerExistenceTable" }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) SetParent(parent types.Entity) { mtetriggerexistencetable.parent = parent }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetParent() types.Entity { return mtetriggerexistencetable.parent }

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry
// Information about a single existence trigger.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has 'existence' set in mteTriggerTest.
type DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggername
    Mtetriggername interface{}

    // The type of existence test to perform.  The trigger fires when the object
    // at mteTriggerValueID is seen to go from present to absent, from absent to
    // present, or to have it's value changed, depending on which tests are
    // selected:  present(0) - when this test is selected, the trigger fires when
    // the mteTriggerValueID object goes from absent to present.  absent(1)  -
    // when this test is selected, the trigger fires when the mteTriggerValueID
    // object goes from present to absent. changed(2) - when this test is
    // selected, the trigger fires the mteTriggerValueID object value changes. 
    // Once the trigger has fired for either presence or absence it will not fire
    // again for that state until the object has been to the other state. . The
    // type is map[string]bool.
    Mtetriggerexistencetest interface{}

    // Control for whether an event may be triggered when this entry is first set
    // to 'active' and the test specified by mteTriggerExistenceTest is true. 
    // Setting an option causes that trigger to fire when its test is true. The
    // type is map[string]bool.
    Mtetriggerexistencestartup interface{}

    // To go with mteTriggerExistenceObjects, the mteOwner of a group of objects
    // from mteObjectsTable. The type is string with length: 0..32.
    Mtetriggerexistenceobjectsowner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable.  These
    // objects are to be added to any Notification resulting from the firing of
    // this trigger for this test.  A list of objects may also be added based on
    // the overall trigger, the event or other settings in mteTriggerTest.  A
    // length of 0 indicates no additional objects. The type is string with
    // length: 0..32.
    Mtetriggerexistenceobjects interface{}

    // To go with mteTriggerExistenceEvent, the mteOwner of an event entry from
    // the mteEventTable. The type is string with length: 0..32.
    Mtetriggerexistenceeventowner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'existence'
    // and this trigger fires.  A length of 0 indicates no event. The type is
    // string with length: 0..32.
    Mtetriggerexistenceevent interface{}
}

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetFilter() yfilter.YFilter { return mtetriggerexistenceentry.YFilter }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) SetFilter(yf yfilter.YFilter) { mtetriggerexistenceentry.YFilter = yf }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteTriggerName" { return "Mtetriggername" }
    if yname == "mteTriggerExistenceTest" { return "Mtetriggerexistencetest" }
    if yname == "mteTriggerExistenceStartup" { return "Mtetriggerexistencestartup" }
    if yname == "mteTriggerExistenceObjectsOwner" { return "Mtetriggerexistenceobjectsowner" }
    if yname == "mteTriggerExistenceObjects" { return "Mtetriggerexistenceobjects" }
    if yname == "mteTriggerExistenceEventOwner" { return "Mtetriggerexistenceeventowner" }
    if yname == "mteTriggerExistenceEvent" { return "Mtetriggerexistenceevent" }
    return ""
}

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetSegmentPath() string {
    return "mteTriggerExistenceEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerexistenceentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerexistenceentry.Mtetriggername) + "']"
}

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mtetriggerexistenceentry.Mteowner
    leafs["mteTriggerName"] = mtetriggerexistenceentry.Mtetriggername
    leafs["mteTriggerExistenceTest"] = mtetriggerexistenceentry.Mtetriggerexistencetest
    leafs["mteTriggerExistenceStartup"] = mtetriggerexistenceentry.Mtetriggerexistencestartup
    leafs["mteTriggerExistenceObjectsOwner"] = mtetriggerexistenceentry.Mtetriggerexistenceobjectsowner
    leafs["mteTriggerExistenceObjects"] = mtetriggerexistenceentry.Mtetriggerexistenceobjects
    leafs["mteTriggerExistenceEventOwner"] = mtetriggerexistenceentry.Mtetriggerexistenceeventowner
    leafs["mteTriggerExistenceEvent"] = mtetriggerexistenceentry.Mtetriggerexistenceevent
    return leafs
}

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetYangName() string { return "mteTriggerExistenceEntry" }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) SetParent(parent types.Entity) { mtetriggerexistenceentry.parent = parent }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetParent() types.Entity { return mtetriggerexistenceentry.parent }

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetParentYangName() string { return "mteTriggerExistenceTable" }

// DISMANEVENTMIB_Mtetriggerbooleantable
// A table of management event trigger information for boolean
// triggers.
type DISMANEVENTMIB_Mtetriggerbooleantable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single boolean trigger.  Entries automatically exist in
    // this this table for each mteTriggerEntry that has 'boolean' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry.
    Mtetriggerbooleanentry []DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry
}

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetFilter() yfilter.YFilter { return mtetriggerbooleantable.YFilter }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) SetFilter(yf yfilter.YFilter) { mtetriggerbooleantable.YFilter = yf }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetGoName(yname string) string {
    if yname == "mteTriggerBooleanEntry" { return "Mtetriggerbooleanentry" }
    return ""
}

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetSegmentPath() string {
    return "mteTriggerBooleanTable"
}

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteTriggerBooleanEntry" {
        for _, c := range mtetriggerbooleantable.Mtetriggerbooleanentry {
            if mtetriggerbooleantable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry{}
        mtetriggerbooleantable.Mtetriggerbooleanentry = append(mtetriggerbooleantable.Mtetriggerbooleanentry, child)
        return &mtetriggerbooleantable.Mtetriggerbooleanentry[len(mtetriggerbooleantable.Mtetriggerbooleanentry)-1]
    }
    return nil
}

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mtetriggerbooleantable.Mtetriggerbooleanentry {
        children[mtetriggerbooleantable.Mtetriggerbooleanentry[i].GetSegmentPath()] = &mtetriggerbooleantable.Mtetriggerbooleanentry[i]
    }
    return children
}

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetYangName() string { return "mteTriggerBooleanTable" }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) SetParent(parent types.Entity) { mtetriggerbooleantable.parent = parent }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetParent() types.Entity { return mtetriggerbooleantable.parent }

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry
// Information about a single boolean trigger.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has 'boolean' set in mteTriggerTest.
type DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggername
    Mtetriggername interface{}

    // The type of boolean comparison to perform.  The value at mteTriggerValueID
    // is compared to mteTriggerBooleanValue, so for example if
    // mteTriggerBooleanComparison is 'less' the result would be true if the value
    // at mteTriggerValueID is less than the value of mteTriggerBooleanValue. The
    // type is Mtetriggerbooleancomparison.
    Mtetriggerbooleancomparison interface{}

    // The value to use for the test specified by mteTriggerBooleanTest. The type
    // is interface{} with range: -2147483648..2147483647.
    Mtetriggerbooleanvalue interface{}

    // Control for whether an event may be triggered when this entry is first set
    // to 'active' or a new instance of the object at mteTriggerValueID is found
    // and the test specified by mteTriggerBooleanComparison is true.  In that
    // case an event is triggered if mteTriggerBooleanStartup is 'true'. The type
    // is bool.
    Mtetriggerbooleanstartup interface{}

    // To go with mteTriggerBooleanObjects, the mteOwner of a group of objects
    // from mteObjectsTable. The type is string with length: 0..32.
    Mtetriggerbooleanobjectsowner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable.  These
    // objects are to be added to any Notification resulting from the firing of
    // this trigger for this test.  A list of objects may also be added based on
    // the overall trigger, the event or other settings in mteTriggerTest.  A
    // length of 0 indicates no additional objects. The type is string with
    // length: 0..32.
    Mtetriggerbooleanobjects interface{}

    // To go with mteTriggerBooleanEvent, the mteOwner of an event entry from
    // mteEventTable. The type is string with length: 0..32.
    Mtetriggerbooleaneventowner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'boolean'
    // and this trigger fires.  A length of 0 indicates no event. The type is
    // string with length: 0..32.
    Mtetriggerbooleanevent interface{}
}

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetFilter() yfilter.YFilter { return mtetriggerbooleanentry.YFilter }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) SetFilter(yf yfilter.YFilter) { mtetriggerbooleanentry.YFilter = yf }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteTriggerName" { return "Mtetriggername" }
    if yname == "mteTriggerBooleanComparison" { return "Mtetriggerbooleancomparison" }
    if yname == "mteTriggerBooleanValue" { return "Mtetriggerbooleanvalue" }
    if yname == "mteTriggerBooleanStartup" { return "Mtetriggerbooleanstartup" }
    if yname == "mteTriggerBooleanObjectsOwner" { return "Mtetriggerbooleanobjectsowner" }
    if yname == "mteTriggerBooleanObjects" { return "Mtetriggerbooleanobjects" }
    if yname == "mteTriggerBooleanEventOwner" { return "Mtetriggerbooleaneventowner" }
    if yname == "mteTriggerBooleanEvent" { return "Mtetriggerbooleanevent" }
    return ""
}

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetSegmentPath() string {
    return "mteTriggerBooleanEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerbooleanentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerbooleanentry.Mtetriggername) + "']"
}

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mtetriggerbooleanentry.Mteowner
    leafs["mteTriggerName"] = mtetriggerbooleanentry.Mtetriggername
    leafs["mteTriggerBooleanComparison"] = mtetriggerbooleanentry.Mtetriggerbooleancomparison
    leafs["mteTriggerBooleanValue"] = mtetriggerbooleanentry.Mtetriggerbooleanvalue
    leafs["mteTriggerBooleanStartup"] = mtetriggerbooleanentry.Mtetriggerbooleanstartup
    leafs["mteTriggerBooleanObjectsOwner"] = mtetriggerbooleanentry.Mtetriggerbooleanobjectsowner
    leafs["mteTriggerBooleanObjects"] = mtetriggerbooleanentry.Mtetriggerbooleanobjects
    leafs["mteTriggerBooleanEventOwner"] = mtetriggerbooleanentry.Mtetriggerbooleaneventowner
    leafs["mteTriggerBooleanEvent"] = mtetriggerbooleanentry.Mtetriggerbooleanevent
    return leafs
}

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetYangName() string { return "mteTriggerBooleanEntry" }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) SetParent(parent types.Entity) { mtetriggerbooleanentry.parent = parent }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetParent() types.Entity { return mtetriggerbooleanentry.parent }

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetParentYangName() string { return "mteTriggerBooleanTable" }

// DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison represents mteTriggerBooleanValue.
type DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison string

const (
    DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison_unequal DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison = "unequal"

    DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison_equal DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison = "equal"

    DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison_less DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison = "less"

    DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison_lessOrEqual DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison = "lessOrEqual"

    DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison_greater DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison = "greater"

    DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison_greaterOrEqual DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry_Mtetriggerbooleancomparison = "greaterOrEqual"
)

// DISMANEVENTMIB_Mtetriggerthresholdtable
// A table of management event trigger information for threshold
// triggers.
type DISMANEVENTMIB_Mtetriggerthresholdtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single threshold trigger.  Entries automatically exist
    // in this table for each mteTriggerEntry that has 'threshold' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry.
    Mtetriggerthresholdentry []DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry
}

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetFilter() yfilter.YFilter { return mtetriggerthresholdtable.YFilter }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) SetFilter(yf yfilter.YFilter) { mtetriggerthresholdtable.YFilter = yf }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetGoName(yname string) string {
    if yname == "mteTriggerThresholdEntry" { return "Mtetriggerthresholdentry" }
    return ""
}

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetSegmentPath() string {
    return "mteTriggerThresholdTable"
}

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteTriggerThresholdEntry" {
        for _, c := range mtetriggerthresholdtable.Mtetriggerthresholdentry {
            if mtetriggerthresholdtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry{}
        mtetriggerthresholdtable.Mtetriggerthresholdentry = append(mtetriggerthresholdtable.Mtetriggerthresholdentry, child)
        return &mtetriggerthresholdtable.Mtetriggerthresholdentry[len(mtetriggerthresholdtable.Mtetriggerthresholdentry)-1]
    }
    return nil
}

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mtetriggerthresholdtable.Mtetriggerthresholdentry {
        children[mtetriggerthresholdtable.Mtetriggerthresholdentry[i].GetSegmentPath()] = &mtetriggerthresholdtable.Mtetriggerthresholdentry[i]
    }
    return children
}

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetYangName() string { return "mteTriggerThresholdTable" }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) SetParent(parent types.Entity) { mtetriggerthresholdtable.parent = parent }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetParent() types.Entity { return mtetriggerthresholdtable.parent }

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry
// Information about a single threshold trigger.  Entries
// automatically exist in this table for each mteTriggerEntry
// that has 'threshold' set in mteTriggerTest.
type DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mtetriggername
    Mtetriggername interface{}

    // The event that may be triggered when this entry is first set to 'active'
    // and a new instance of the object at mteTriggerValueID is found.  If the
    // first sample after this instance becomes active is greater than or equal to
    // mteTriggerThresholdRising and mteTriggerThresholdStartup is equal to
    // 'rising' or 'risingOrFalling', then one mteTriggerThresholdRisingEvent is
    // triggered for that instance. If the first sample after this entry becomes
    // active is less than or equal to mteTriggerThresholdFalling and
    // mteTriggerThresholdStartup is equal to 'falling' or 'risingOrFalling', then
    // one mteTriggerThresholdRisingEvent is triggered for that instance. The type
    // is Mtetriggerthresholdstartup.
    Mtetriggerthresholdstartup interface{}

    // A threshold value to check against if mteTriggerType is 'threshold'.  When
    // the current sampled value is greater than or equal to this threshold, and
    // the value at the last sampling interval was less than this threshold, one
    // mteTriggerThresholdRisingEvent is triggered.  That event is also triggered
    // if the first sample after this entry becomes active is greater than or
    // equal to this threshold and mteTriggerThresholdStartup is equal to 'rising'
    // or 'risingOrFalling'.  After a rising event is generated, another such
    // event is not triggered until the sampled value falls below this threshold
    // and reaches mteTriggerThresholdFalling. The type is interface{} with range:
    // -2147483648..2147483647.
    Mtetriggerthresholdrising interface{}

    // A threshold value to check against if mteTriggerType is 'threshold'.  When
    // the current sampled value is less than or equal to this threshold, and the
    // value at the last sampling interval was greater than this threshold, one
    // mteTriggerThresholdFallingEvent is triggered.  That event is also triggered
    // if the first sample after this entry becomes active is less than or equal
    // to this threshold and mteTriggerThresholdStartup is equal to 'falling' or
    // 'risingOrFalling'.  After a falling event is generated, another such event
    // is not triggered until the sampled value rises above this threshold and
    // reaches mteTriggerThresholdRising. The type is interface{} with range:
    // -2147483648..2147483647.
    Mtetriggerthresholdfalling interface{}

    // A threshold value to check against if mteTriggerType is 'threshold'.  When
    // the delta value (difference) between the current sampled value (value(n))
    // and the previous sampled value (value(n-1)) is greater than or equal to
    // this threshold, and the delta value calculated at the last sampling
    // interval (i.e. value(n-1) - value(n-2)) was less than this threshold, one
    // mteTriggerThresholdDeltaRisingEvent is triggered. That event is also
    // triggered if the first delta value calculated after this entry becomes
    // active, i.e. value(2) - value(1), where value(1) is the first sample taken
    // of that instance, is greater than or equal to this threshold.  After a
    // rising event is generated, another such event is not triggered until the
    // delta value falls below this threshold and reaches
    // mteTriggerThresholdDeltaFalling. The type is interface{} with range:
    // -2147483648..2147483647.
    Mtetriggerthresholddeltarising interface{}

    // A threshold value to check against if mteTriggerType is 'threshold'.  When
    // the delta value (difference) between the current sampled value (value(n))
    // and the previous sampled value (value(n-1)) is less than or equal to this
    // threshold, and the delta value calculated at the last sampling interval
    // (i.e. value(n-1) - value(n-2)) was greater than this threshold, one
    // mteTriggerThresholdDeltaFallingEvent is triggered. That event is also
    // triggered if the first delta value calculated after this entry becomes
    // active, i.e. value(2) - value(1), where value(1) is the first sample taken
    // of that instance, is less than or equal to this threshold.  After a falling
    // event is generated, another such event is not triggered until the delta
    // value falls below this threshold and reaches
    // mteTriggerThresholdDeltaRising. The type is interface{} with range:
    // -2147483648..2147483647.
    Mtetriggerthresholddeltafalling interface{}

    // To go with mteTriggerThresholdObjects, the mteOwner of a group of objects
    // from mteObjectsTable. The type is string with length: 0..32.
    Mtetriggerthresholdobjectsowner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable.  These
    // objects are to be added to any Notification resulting from the firing of
    // this trigger for this test.  A list of objects may also be added based on
    // the overall  trigger, the event or other settings in mteTriggerTest.  A
    // length of 0 indicates no additional objects. The type is string with
    // length: 0..32.
    Mtetriggerthresholdobjects interface{}

    // To go with mteTriggerThresholdRisingEvent, the mteOwner of an event entry
    // from mteEventTable. The type is string with length: 0..32.
    Mtetriggerthresholdrisingeventowner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'threshold'
    // and this trigger fires based on mteTriggerThresholdRising.  A length of 0
    // indicates no event. The type is string with length: 0..32.
    Mtetriggerthresholdrisingevent interface{}

    // To go with mteTriggerThresholdFallingEvent, the mteOwner of an event entry
    // from mteEventTable. The type is string with length: 0..32.
    Mtetriggerthresholdfallingeventowner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'threshold'
    // and this trigger fires based on mteTriggerThresholdFalling.  A length of 0
    // indicates no event. The type is string with length: 0..32.
    Mtetriggerthresholdfallingevent interface{}

    // To go with mteTriggerThresholdDeltaRisingEvent, the mteOwner of an event
    // entry from mteEventTable. The type is string with length: 0..32.
    Mtetriggerthresholddeltarisingeventowner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'threshold'
    // and this trigger fires based on mteTriggerThresholdDeltaRising. A length of
    // 0 indicates no event. The type is string with length: 0..32.
    Mtetriggerthresholddeltarisingevent interface{}

    // To go with mteTriggerThresholdDeltaFallingEvent, the mteOwner of an event
    // entry from mteEventTable. The type is string with length: 0..32.
    Mtetriggerthresholddeltafallingeventowner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'threshold'
    // and this trigger fires based on mteTriggerThresholdDeltaFalling.  A length
    // of 0 indicates no event. The type is string with length: 0..32.
    Mtetriggerthresholddeltafallingevent interface{}
}

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetFilter() yfilter.YFilter { return mtetriggerthresholdentry.YFilter }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) SetFilter(yf yfilter.YFilter) { mtetriggerthresholdentry.YFilter = yf }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteTriggerName" { return "Mtetriggername" }
    if yname == "mteTriggerThresholdStartup" { return "Mtetriggerthresholdstartup" }
    if yname == "mteTriggerThresholdRising" { return "Mtetriggerthresholdrising" }
    if yname == "mteTriggerThresholdFalling" { return "Mtetriggerthresholdfalling" }
    if yname == "mteTriggerThresholdDeltaRising" { return "Mtetriggerthresholddeltarising" }
    if yname == "mteTriggerThresholdDeltaFalling" { return "Mtetriggerthresholddeltafalling" }
    if yname == "mteTriggerThresholdObjectsOwner" { return "Mtetriggerthresholdobjectsowner" }
    if yname == "mteTriggerThresholdObjects" { return "Mtetriggerthresholdobjects" }
    if yname == "mteTriggerThresholdRisingEventOwner" { return "Mtetriggerthresholdrisingeventowner" }
    if yname == "mteTriggerThresholdRisingEvent" { return "Mtetriggerthresholdrisingevent" }
    if yname == "mteTriggerThresholdFallingEventOwner" { return "Mtetriggerthresholdfallingeventowner" }
    if yname == "mteTriggerThresholdFallingEvent" { return "Mtetriggerthresholdfallingevent" }
    if yname == "mteTriggerThresholdDeltaRisingEventOwner" { return "Mtetriggerthresholddeltarisingeventowner" }
    if yname == "mteTriggerThresholdDeltaRisingEvent" { return "Mtetriggerthresholddeltarisingevent" }
    if yname == "mteTriggerThresholdDeltaFallingEventOwner" { return "Mtetriggerthresholddeltafallingeventowner" }
    if yname == "mteTriggerThresholdDeltaFallingEvent" { return "Mtetriggerthresholddeltafallingevent" }
    return ""
}

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetSegmentPath() string {
    return "mteTriggerThresholdEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerthresholdentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerthresholdentry.Mtetriggername) + "']"
}

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mtetriggerthresholdentry.Mteowner
    leafs["mteTriggerName"] = mtetriggerthresholdentry.Mtetriggername
    leafs["mteTriggerThresholdStartup"] = mtetriggerthresholdentry.Mtetriggerthresholdstartup
    leafs["mteTriggerThresholdRising"] = mtetriggerthresholdentry.Mtetriggerthresholdrising
    leafs["mteTriggerThresholdFalling"] = mtetriggerthresholdentry.Mtetriggerthresholdfalling
    leafs["mteTriggerThresholdDeltaRising"] = mtetriggerthresholdentry.Mtetriggerthresholddeltarising
    leafs["mteTriggerThresholdDeltaFalling"] = mtetriggerthresholdentry.Mtetriggerthresholddeltafalling
    leafs["mteTriggerThresholdObjectsOwner"] = mtetriggerthresholdentry.Mtetriggerthresholdobjectsowner
    leafs["mteTriggerThresholdObjects"] = mtetriggerthresholdentry.Mtetriggerthresholdobjects
    leafs["mteTriggerThresholdRisingEventOwner"] = mtetriggerthresholdentry.Mtetriggerthresholdrisingeventowner
    leafs["mteTriggerThresholdRisingEvent"] = mtetriggerthresholdentry.Mtetriggerthresholdrisingevent
    leafs["mteTriggerThresholdFallingEventOwner"] = mtetriggerthresholdentry.Mtetriggerthresholdfallingeventowner
    leafs["mteTriggerThresholdFallingEvent"] = mtetriggerthresholdentry.Mtetriggerthresholdfallingevent
    leafs["mteTriggerThresholdDeltaRisingEventOwner"] = mtetriggerthresholdentry.Mtetriggerthresholddeltarisingeventowner
    leafs["mteTriggerThresholdDeltaRisingEvent"] = mtetriggerthresholdentry.Mtetriggerthresholddeltarisingevent
    leafs["mteTriggerThresholdDeltaFallingEventOwner"] = mtetriggerthresholdentry.Mtetriggerthresholddeltafallingeventowner
    leafs["mteTriggerThresholdDeltaFallingEvent"] = mtetriggerthresholdentry.Mtetriggerthresholddeltafallingevent
    return leafs
}

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetBundleName() string { return "cisco_ios_xe" }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetYangName() string { return "mteTriggerThresholdEntry" }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) SetParent(parent types.Entity) { mtetriggerthresholdentry.parent = parent }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetParent() types.Entity { return mtetriggerthresholdentry.parent }

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetParentYangName() string { return "mteTriggerThresholdTable" }

// DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry_Mtetriggerthresholdstartup represents triggered for that instance.
type DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry_Mtetriggerthresholdstartup string

const (
    DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry_Mtetriggerthresholdstartup_rising DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry_Mtetriggerthresholdstartup = "rising"

    DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry_Mtetriggerthresholdstartup_falling DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry_Mtetriggerthresholdstartup = "falling"

    DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry_Mtetriggerthresholdstartup_risingOrFalling DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry_Mtetriggerthresholdstartup = "risingOrFalling"
)

// DISMANEVENTMIB_Mteobjectstable
// A table of objects that can be added to notifications based
// on the trigger, trigger test, or event, as pointed to by
// entries in those tables.
type DISMANEVENTMIB_Mteobjectstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A group of objects.  Applications create and delete entries using
    // mteObjectsEntryStatus.  When adding objects to a notification they are
    // added in the lexical order of their index in this table.  Those associated
    // with a trigger come first, then trigger test, then event. The type is slice
    // of DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry.
    Mteobjectsentry []DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry
}

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetFilter() yfilter.YFilter { return mteobjectstable.YFilter }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) SetFilter(yf yfilter.YFilter) { mteobjectstable.YFilter = yf }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetGoName(yname string) string {
    if yname == "mteObjectsEntry" { return "Mteobjectsentry" }
    return ""
}

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetSegmentPath() string {
    return "mteObjectsTable"
}

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteObjectsEntry" {
        for _, c := range mteobjectstable.Mteobjectsentry {
            if mteobjectstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry{}
        mteobjectstable.Mteobjectsentry = append(mteobjectstable.Mteobjectsentry, child)
        return &mteobjectstable.Mteobjectsentry[len(mteobjectstable.Mteobjectsentry)-1]
    }
    return nil
}

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mteobjectstable.Mteobjectsentry {
        children[mteobjectstable.Mteobjectsentry[i].GetSegmentPath()] = &mteobjectstable.Mteobjectsentry[i]
    }
    return children
}

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetBundleName() string { return "cisco_ios_xe" }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetYangName() string { return "mteObjectsTable" }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) SetParent(parent types.Entity) { mteobjectstable.parent = parent }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetParent() types.Entity { return mteobjectstable.parent }

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry
// A group of objects.  Applications create and delete entries
// using mteObjectsEntryStatus.
// 
// When adding objects to a notification they are added in the
// lexical order of their index in this table.  Those associated
// with a trigger come first, then trigger test, then event.
type DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. A locally-unique, administratively assigned name
    // for a group of objects. The type is string with length: 1..32.
    Mteobjectsname interface{}

    // This attribute is a key. An arbitrary integer for the purpose of
    // identifying individual objects within a mteObjectsName group.  Objects
    // within a group are placed in the notification in the numerical order of
    // this index.  Groups are placed in the notification in the order of the
    // selections for overall trigger, trigger test, and event. Within trigger
    // test they are in the same order as the numerical values of the bits defined
    // for mteTriggerTest.  Bad object identifiers or a mismatch between
    // truncating the identifier and the value of mteDeltaDiscontinuityIDWildcard
    // result in operation as one would expect when providing the wrong identifier
    // to a Get operation.  The Get will fail or get the wrong object.  If the
    // object is not available it is omitted from the notification. The type is
    // interface{} with range: 1..4294967295.
    Mteobjectsindex interface{}

    // The object identifier of a MIB object to add to a Notification that results
    // from the firing of a trigger.  This may be wildcarded by truncating all or
    // part of the instance portion, in which case the instance portion of the OID
    // for obtaining this object will be the same as that used in obtaining the
    // mteTriggerValueID that fired.  If such wildcarding is applied,
    // mteObjectsIDWildcard must be 'true' and if not it must be 'false'.  Each
    // instance that fills the wildcard is independent of any additional
    // instances, that is, wildcarded objects operate as if there were a separate
    // table entry for each instance that fills the wildcard without having to
    // actually predict all possible instances ahead of time. The type is string
    // with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Mteobjectsid interface{}

    // Control for whether mteObjectsID is to be treated as fully-specified or
    // wildcarded, with 'true' indicating wildcard. The type is bool.
    Mteobjectsidwildcard interface{}

    // The control that allows creation and deletion of entries. Once made active
    // an entry MAY not be modified except to delete it. The type is RowStatus.
    Mteobjectsentrystatus interface{}
}

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetFilter() yfilter.YFilter { return mteobjectsentry.YFilter }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) SetFilter(yf yfilter.YFilter) { mteobjectsentry.YFilter = yf }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteObjectsName" { return "Mteobjectsname" }
    if yname == "mteObjectsIndex" { return "Mteobjectsindex" }
    if yname == "mteObjectsID" { return "Mteobjectsid" }
    if yname == "mteObjectsIDWildcard" { return "Mteobjectsidwildcard" }
    if yname == "mteObjectsEntryStatus" { return "Mteobjectsentrystatus" }
    return ""
}

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetSegmentPath() string {
    return "mteObjectsEntry" + "[mteOwner='" + fmt.Sprintf("%v", mteobjectsentry.Mteowner) + "']" + "[mteObjectsName='" + fmt.Sprintf("%v", mteobjectsentry.Mteobjectsname) + "']" + "[mteObjectsIndex='" + fmt.Sprintf("%v", mteobjectsentry.Mteobjectsindex) + "']"
}

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mteobjectsentry.Mteowner
    leafs["mteObjectsName"] = mteobjectsentry.Mteobjectsname
    leafs["mteObjectsIndex"] = mteobjectsentry.Mteobjectsindex
    leafs["mteObjectsID"] = mteobjectsentry.Mteobjectsid
    leafs["mteObjectsIDWildcard"] = mteobjectsentry.Mteobjectsidwildcard
    leafs["mteObjectsEntryStatus"] = mteobjectsentry.Mteobjectsentrystatus
    return leafs
}

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetBundleName() string { return "cisco_ios_xe" }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetYangName() string { return "mteObjectsEntry" }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) SetParent(parent types.Entity) { mteobjectsentry.parent = parent }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetParent() types.Entity { return mteobjectsentry.parent }

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetParentYangName() string { return "mteObjectsTable" }

// DISMANEVENTMIB_Mteeventtable
// A table of management event action information.
type DISMANEVENTMIB_Mteeventtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single event.  Applications create and delete entries
    // using mteEventEntryStatus. The type is slice of
    // DISMANEVENTMIB_Mteeventtable_Mteevententry.
    Mteevententry []DISMANEVENTMIB_Mteeventtable_Mteevententry
}

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetFilter() yfilter.YFilter { return mteeventtable.YFilter }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) SetFilter(yf yfilter.YFilter) { mteeventtable.YFilter = yf }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetGoName(yname string) string {
    if yname == "mteEventEntry" { return "Mteevententry" }
    return ""
}

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetSegmentPath() string {
    return "mteEventTable"
}

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteEventEntry" {
        for _, c := range mteeventtable.Mteevententry {
            if mteeventtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mteeventtable_Mteevententry{}
        mteeventtable.Mteevententry = append(mteeventtable.Mteevententry, child)
        return &mteeventtable.Mteevententry[len(mteeventtable.Mteevententry)-1]
    }
    return nil
}

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mteeventtable.Mteevententry {
        children[mteeventtable.Mteevententry[i].GetSegmentPath()] = &mteeventtable.Mteevententry[i]
    }
    return children
}

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetBundleName() string { return "cisco_ios_xe" }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetYangName() string { return "mteEventTable" }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) SetParent(parent types.Entity) { mteeventtable.parent = parent }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetParent() types.Entity { return mteeventtable.parent }

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mteeventtable_Mteevententry
// Information about a single event.  Applications create and
// delete entries using mteEventEntryStatus.
type DISMANEVENTMIB_Mteeventtable_Mteevententry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. A locally-unique, administratively assigned name
    // for the event. The type is string with length: 1..32.
    Mteeventname interface{}

    // A description of the event's function and use. The type is string.
    Mteeventcomment interface{}

    // The actions to perform when this event occurs.  For 'notification', Traps
    // and/or Informs are sent according to the configuration in the SNMP
    // Notification MIB.  For 'set', an SNMP Set operation is performed according
    // to control values in this entry. The type is map[string]bool.
    Mteeventactions interface{}

    // A control to allow an event to be configured but not used. When the value
    // is 'false' the event does not execute even if  triggered. The type is bool.
    Mteeventenabled interface{}

    // The control that allows creation and deletion of entries. Once made active
    // an entry MAY not be modified except to delete it. The type is RowStatus.
    Mteevententrystatus interface{}
}

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetFilter() yfilter.YFilter { return mteevententry.YFilter }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) SetFilter(yf yfilter.YFilter) { mteevententry.YFilter = yf }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteEventName" { return "Mteeventname" }
    if yname == "mteEventComment" { return "Mteeventcomment" }
    if yname == "mteEventActions" { return "Mteeventactions" }
    if yname == "mteEventEnabled" { return "Mteeventenabled" }
    if yname == "mteEventEntryStatus" { return "Mteevententrystatus" }
    return ""
}

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetSegmentPath() string {
    return "mteEventEntry" + "[mteOwner='" + fmt.Sprintf("%v", mteevententry.Mteowner) + "']" + "[mteEventName='" + fmt.Sprintf("%v", mteevententry.Mteeventname) + "']"
}

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mteevententry.Mteowner
    leafs["mteEventName"] = mteevententry.Mteeventname
    leafs["mteEventComment"] = mteevententry.Mteeventcomment
    leafs["mteEventActions"] = mteevententry.Mteeventactions
    leafs["mteEventEnabled"] = mteevententry.Mteeventenabled
    leafs["mteEventEntryStatus"] = mteevententry.Mteevententrystatus
    return leafs
}

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetBundleName() string { return "cisco_ios_xe" }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetYangName() string { return "mteEventEntry" }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) SetParent(parent types.Entity) { mteevententry.parent = parent }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetParent() types.Entity { return mteevententry.parent }

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetParentYangName() string { return "mteEventTable" }

// DISMANEVENTMIB_Mteeventnotificationtable
// A table of information about notifications to be sent as a
// consequence of management events.
type DISMANEVENTMIB_Mteeventnotificationtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single event's notification.  Entries automatically
    // exist in this this table for each mteEventEntry that has 'notification' set
    // in mteEventActions. The type is slice of
    // DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry.
    Mteeventnotificationentry []DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry
}

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetFilter() yfilter.YFilter { return mteeventnotificationtable.YFilter }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) SetFilter(yf yfilter.YFilter) { mteeventnotificationtable.YFilter = yf }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetGoName(yname string) string {
    if yname == "mteEventNotificationEntry" { return "Mteeventnotificationentry" }
    return ""
}

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetSegmentPath() string {
    return "mteEventNotificationTable"
}

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteEventNotificationEntry" {
        for _, c := range mteeventnotificationtable.Mteeventnotificationentry {
            if mteeventnotificationtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry{}
        mteeventnotificationtable.Mteeventnotificationentry = append(mteeventnotificationtable.Mteeventnotificationentry, child)
        return &mteeventnotificationtable.Mteeventnotificationentry[len(mteeventnotificationtable.Mteeventnotificationentry)-1]
    }
    return nil
}

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mteeventnotificationtable.Mteeventnotificationentry {
        children[mteeventnotificationtable.Mteeventnotificationentry[i].GetSegmentPath()] = &mteeventnotificationtable.Mteeventnotificationentry[i]
    }
    return children
}

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetBundleName() string { return "cisco_ios_xe" }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetYangName() string { return "mteEventNotificationTable" }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) SetParent(parent types.Entity) { mteeventnotificationtable.parent = parent }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetParent() types.Entity { return mteeventnotificationtable.parent }

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry
// Information about a single event's notification.  Entries
// automatically exist in this this table for each mteEventEntry
// that has 'notification' set in mteEventActions.
type DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mteeventtable_Mteevententry_Mteeventname
    Mteeventname interface{}

    // The object identifier from the NOTIFICATION-TYPE for the notification to
    // use if metEventActions has 'notification' set. The type is string with
    // pattern: (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Mteeventnotification interface{}

    // To go with mteEventNotificationObjects, the mteOwner of a group of objects
    // from mteObjectsTable. The type is string with length: 0..32.
    Mteeventnotificationobjectsowner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable if
    // mteEventActions has 'notification' set. These objects are to be added to
    // any Notification generated by this event.  Objects may also be added based
    // on the trigger that stimulated the event.  A length of 0 indicates no
    // additional objects. The type is string with length: 0..32.
    Mteeventnotificationobjects interface{}
}

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetFilter() yfilter.YFilter { return mteeventnotificationentry.YFilter }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) SetFilter(yf yfilter.YFilter) { mteeventnotificationentry.YFilter = yf }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteEventName" { return "Mteeventname" }
    if yname == "mteEventNotification" { return "Mteeventnotification" }
    if yname == "mteEventNotificationObjectsOwner" { return "Mteeventnotificationobjectsowner" }
    if yname == "mteEventNotificationObjects" { return "Mteeventnotificationobjects" }
    return ""
}

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetSegmentPath() string {
    return "mteEventNotificationEntry" + "[mteOwner='" + fmt.Sprintf("%v", mteeventnotificationentry.Mteowner) + "']" + "[mteEventName='" + fmt.Sprintf("%v", mteeventnotificationentry.Mteeventname) + "']"
}

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mteeventnotificationentry.Mteowner
    leafs["mteEventName"] = mteeventnotificationentry.Mteeventname
    leafs["mteEventNotification"] = mteeventnotificationentry.Mteeventnotification
    leafs["mteEventNotificationObjectsOwner"] = mteeventnotificationentry.Mteeventnotificationobjectsowner
    leafs["mteEventNotificationObjects"] = mteeventnotificationentry.Mteeventnotificationobjects
    return leafs
}

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetBundleName() string { return "cisco_ios_xe" }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetYangName() string { return "mteEventNotificationEntry" }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) SetParent(parent types.Entity) { mteeventnotificationentry.parent = parent }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetParent() types.Entity { return mteeventnotificationentry.parent }

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetParentYangName() string { return "mteEventNotificationTable" }

// DISMANEVENTMIB_Mteeventsettable
// A table of management event action information.
type DISMANEVENTMIB_Mteeventsettable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single event's set option.  Entries automatically exist
    // in this this table for each mteEventEntry that has 'set' set in
    // mteEventActions. The type is slice of
    // DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry.
    Mteeventsetentry []DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry
}

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetFilter() yfilter.YFilter { return mteeventsettable.YFilter }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) SetFilter(yf yfilter.YFilter) { mteeventsettable.YFilter = yf }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetGoName(yname string) string {
    if yname == "mteEventSetEntry" { return "Mteeventsetentry" }
    return ""
}

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetSegmentPath() string {
    return "mteEventSetTable"
}

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mteEventSetEntry" {
        for _, c := range mteeventsettable.Mteeventsetentry {
            if mteeventsettable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry{}
        mteeventsettable.Mteeventsetentry = append(mteeventsettable.Mteeventsetentry, child)
        return &mteeventsettable.Mteeventsetentry[len(mteeventsettable.Mteeventsetentry)-1]
    }
    return nil
}

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mteeventsettable.Mteeventsetentry {
        children[mteeventsettable.Mteeventsetentry[i].GetSegmentPath()] = &mteeventsettable.Mteeventsetentry[i]
    }
    return children
}

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetBundleName() string { return "cisco_ios_xe" }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetYangName() string { return "mteEventSetTable" }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) SetParent(parent types.Entity) { mteeventsettable.parent = parent }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetParent() types.Entity { return mteeventsettable.parent }

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetParentYangName() string { return "DISMAN-EVENT-MIB" }

// DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry
// Information about a single event's set option.  Entries
// automatically exist in this this table for each mteEventEntry
// that has 'set' set in mteEventActions.
type DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mteeventtable_Mteevententry_Mteeventname
    Mteeventname interface{}

    // The object identifier from the MIB object to set if mteEventActions has
    // 'set' set.  This object identifier may be wildcarded by leaving
    // sub-identifiers off the end, in which case nteEventSetObjectWildCard must
    // be 'true'.  If mteEventSetObject is wildcarded the instance used to set the
    // object to which it points is the same as the instance from the value of
    // mteTriggerValueID that triggered the event.  Each instance that fills the
    // wildcard is independent of any additional instances, that is, wildcarded
    // objects operate as if there were a separate table entry for each instance
    // that fills the wildcard without having to actually predict all possible
    // instances ahead of time.  Bad object identifiers or a mismatch between
    // truncating the identifier and the value of mteSetObjectWildcard result in
    // operation as one would expect when providing the wrong identifier to a Set
    // operation.  The Set will fail or set the wrong object.  If the value syntax
    // of the destination object is not correct, the Set fails with the normal
    // SNMP error code. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Mteeventsetobject interface{}

    // Control over whether mteEventSetObject is to be treated as fully-specified
    // or wildcarded, with 'true' indicating wildcard if mteEventActions has 'set'
    // set. The type is bool.
    Mteeventsetobjectwildcard interface{}

    // The value to which to set the object at mteEventSetObject if
    // mteEventActions has 'set' set. The type is interface{} with range:
    // -2147483648..2147483647.
    Mteeventsetvalue interface{}

    // The tag for the target(s) at which to set the object at mteEventSetObject
    // to mteEventSetValue if mteEventActions has 'set' set.  Systems limited to
    // self management MAY reject a non-zero length for the value of this object. 
    // A length of 0 indicates the local system.  In this case, access to the
    // objects indicated by mteEventSetObject is under the security credentials of
    // the requester that set mteTriggerEntryStatus to 'active'.  Those
    // credentials are the input parameters for isAccessAllowed from the
    // Architecture for Describing SNMP Management Frameworks.  Otherwise access
    // rights are checked according to the security parameters resulting from the
    // tag. The type is string.
    Mteeventsettargettag interface{}

    // The management context in which to set mteEventObjectID. if mteEventActions
    // has 'set' set.  This may be wildcarded by leaving characters off the end. 
    // To indicate such wildcarding mteEventSetContextNameWildcard must be 'true'.
    // If this context name is wildcarded the value used to complete the
    // wildcarding of mteTriggerContextName will be appended. The type is string.
    Mteeventsetcontextname interface{}

    // Control for whether mteEventSetContextName is to be treated as
    // fully-specified or wildcarded, with 'true' indicating wildcard if
    // mteEventActions has 'set' set. The type is bool.
    Mteeventsetcontextnamewildcard interface{}
}

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetFilter() yfilter.YFilter { return mteeventsetentry.YFilter }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) SetFilter(yf yfilter.YFilter) { mteeventsetentry.YFilter = yf }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetGoName(yname string) string {
    if yname == "mteOwner" { return "Mteowner" }
    if yname == "mteEventName" { return "Mteeventname" }
    if yname == "mteEventSetObject" { return "Mteeventsetobject" }
    if yname == "mteEventSetObjectWildcard" { return "Mteeventsetobjectwildcard" }
    if yname == "mteEventSetValue" { return "Mteeventsetvalue" }
    if yname == "mteEventSetTargetTag" { return "Mteeventsettargettag" }
    if yname == "mteEventSetContextName" { return "Mteeventsetcontextname" }
    if yname == "mteEventSetContextNameWildcard" { return "Mteeventsetcontextnamewildcard" }
    return ""
}

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetSegmentPath() string {
    return "mteEventSetEntry" + "[mteOwner='" + fmt.Sprintf("%v", mteeventsetentry.Mteowner) + "']" + "[mteEventName='" + fmt.Sprintf("%v", mteeventsetentry.Mteeventname) + "']"
}

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mteOwner"] = mteeventsetentry.Mteowner
    leafs["mteEventName"] = mteeventsetentry.Mteeventname
    leafs["mteEventSetObject"] = mteeventsetentry.Mteeventsetobject
    leafs["mteEventSetObjectWildcard"] = mteeventsetentry.Mteeventsetobjectwildcard
    leafs["mteEventSetValue"] = mteeventsetentry.Mteeventsetvalue
    leafs["mteEventSetTargetTag"] = mteeventsetentry.Mteeventsettargettag
    leafs["mteEventSetContextName"] = mteeventsetentry.Mteeventsetcontextname
    leafs["mteEventSetContextNameWildcard"] = mteeventsetentry.Mteeventsetcontextnamewildcard
    return leafs
}

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetBundleName() string { return "cisco_ios_xe" }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetYangName() string { return "mteEventSetEntry" }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) SetParent(parent types.Entity) { mteeventsetentry.parent = parent }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetParent() types.Entity { return mteeventsetentry.parent }

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetParentYangName() string { return "mteEventSetTable" }

