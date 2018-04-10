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
    EntityData types.CommonEntityData
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

func (dISMANEVENTMIB *DISMANEVENTMIB) GetEntityData() *types.CommonEntityData {
    dISMANEVENTMIB.EntityData.YFilter = dISMANEVENTMIB.YFilter
    dISMANEVENTMIB.EntityData.YangName = "DISMAN-EVENT-MIB"
    dISMANEVENTMIB.EntityData.BundleName = "cisco_ios_xe"
    dISMANEVENTMIB.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    dISMANEVENTMIB.EntityData.SegmentPath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB"
    dISMANEVENTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dISMANEVENTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dISMANEVENTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dISMANEVENTMIB.EntityData.Children = make(map[string]types.YChild)
    dISMANEVENTMIB.EntityData.Children["mteResource"] = types.YChild{"Mteresource", &dISMANEVENTMIB.Mteresource}
    dISMANEVENTMIB.EntityData.Children["mteTrigger"] = types.YChild{"Mtetrigger", &dISMANEVENTMIB.Mtetrigger}
    dISMANEVENTMIB.EntityData.Children["mteEvent"] = types.YChild{"Mteevent", &dISMANEVENTMIB.Mteevent}
    dISMANEVENTMIB.EntityData.Children["mteTriggerTable"] = types.YChild{"Mtetriggertable", &dISMANEVENTMIB.Mtetriggertable}
    dISMANEVENTMIB.EntityData.Children["mteTriggerDeltaTable"] = types.YChild{"Mtetriggerdeltatable", &dISMANEVENTMIB.Mtetriggerdeltatable}
    dISMANEVENTMIB.EntityData.Children["mteTriggerExistenceTable"] = types.YChild{"Mtetriggerexistencetable", &dISMANEVENTMIB.Mtetriggerexistencetable}
    dISMANEVENTMIB.EntityData.Children["mteTriggerBooleanTable"] = types.YChild{"Mtetriggerbooleantable", &dISMANEVENTMIB.Mtetriggerbooleantable}
    dISMANEVENTMIB.EntityData.Children["mteTriggerThresholdTable"] = types.YChild{"Mtetriggerthresholdtable", &dISMANEVENTMIB.Mtetriggerthresholdtable}
    dISMANEVENTMIB.EntityData.Children["mteObjectsTable"] = types.YChild{"Mteobjectstable", &dISMANEVENTMIB.Mteobjectstable}
    dISMANEVENTMIB.EntityData.Children["mteEventTable"] = types.YChild{"Mteeventtable", &dISMANEVENTMIB.Mteeventtable}
    dISMANEVENTMIB.EntityData.Children["mteEventNotificationTable"] = types.YChild{"Mteeventnotificationtable", &dISMANEVENTMIB.Mteeventnotificationtable}
    dISMANEVENTMIB.EntityData.Children["mteEventSetTable"] = types.YChild{"Mteeventsettable", &dISMANEVENTMIB.Mteeventsettable}
    dISMANEVENTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dISMANEVENTMIB.EntityData)
}

// DISMANEVENTMIB_Mteresource
type DISMANEVENTMIB_Mteresource struct {
    EntityData types.CommonEntityData
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

func (mteresource *DISMANEVENTMIB_Mteresource) GetEntityData() *types.CommonEntityData {
    mteresource.EntityData.YFilter = mteresource.YFilter
    mteresource.EntityData.YangName = "mteResource"
    mteresource.EntityData.BundleName = "cisco_ios_xe"
    mteresource.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteresource.EntityData.SegmentPath = "mteResource"
    mteresource.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteresource.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteresource.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteresource.EntityData.Children = make(map[string]types.YChild)
    mteresource.EntityData.Leafs = make(map[string]types.YLeaf)
    mteresource.EntityData.Leafs["mteResourceSampleMinimum"] = types.YLeaf{"Mteresourcesampleminimum", mteresource.Mteresourcesampleminimum}
    mteresource.EntityData.Leafs["mteResourceSampleInstanceMaximum"] = types.YLeaf{"Mteresourcesampleinstancemaximum", mteresource.Mteresourcesampleinstancemaximum}
    mteresource.EntityData.Leafs["mteResourceSampleInstances"] = types.YLeaf{"Mteresourcesampleinstances", mteresource.Mteresourcesampleinstances}
    mteresource.EntityData.Leafs["mteResourceSampleInstancesHigh"] = types.YLeaf{"Mteresourcesampleinstanceshigh", mteresource.Mteresourcesampleinstanceshigh}
    mteresource.EntityData.Leafs["mteResourceSampleInstanceLacks"] = types.YLeaf{"Mteresourcesampleinstancelacks", mteresource.Mteresourcesampleinstancelacks}
    return &(mteresource.EntityData)
}

// DISMANEVENTMIB_Mtetrigger
type DISMANEVENTMIB_Mtetrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times an attempt to check for a trigger condition has failed.
    // This counts individually for each attempt in a group of targets or each
    // attempt for a  wildcarded object. The type is interface{} with range:
    // 0..4294967295. Units are failures.
    Mtetriggerfailures interface{}
}

func (mtetrigger *DISMANEVENTMIB_Mtetrigger) GetEntityData() *types.CommonEntityData {
    mtetrigger.EntityData.YFilter = mtetrigger.YFilter
    mtetrigger.EntityData.YangName = "mteTrigger"
    mtetrigger.EntityData.BundleName = "cisco_ios_xe"
    mtetrigger.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mtetrigger.EntityData.SegmentPath = "mteTrigger"
    mtetrigger.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetrigger.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetrigger.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetrigger.EntityData.Children = make(map[string]types.YChild)
    mtetrigger.EntityData.Leafs = make(map[string]types.YLeaf)
    mtetrigger.EntityData.Leafs["mteTriggerFailures"] = types.YLeaf{"Mtetriggerfailures", mtetrigger.Mtetriggerfailures}
    return &(mtetrigger.EntityData)
}

// DISMANEVENTMIB_Mteevent
type DISMANEVENTMIB_Mteevent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times an attempt to invoke an event has failed.  This counts
    // individually for each attempt in a group of targets or each attempt for a
    // wildcarded trigger object. The type is interface{} with range:
    // 0..4294967295.
    Mteeventfailures interface{}
}

func (mteevent *DISMANEVENTMIB_Mteevent) GetEntityData() *types.CommonEntityData {
    mteevent.EntityData.YFilter = mteevent.YFilter
    mteevent.EntityData.YangName = "mteEvent"
    mteevent.EntityData.BundleName = "cisco_ios_xe"
    mteevent.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteevent.EntityData.SegmentPath = "mteEvent"
    mteevent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteevent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteevent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteevent.EntityData.Children = make(map[string]types.YChild)
    mteevent.EntityData.Leafs = make(map[string]types.YLeaf)
    mteevent.EntityData.Leafs["mteEventFailures"] = types.YLeaf{"Mteeventfailures", mteevent.Mteeventfailures}
    return &(mteevent.EntityData)
}

// DISMANEVENTMIB_Mtetriggertable
// A table of management event trigger information.
type DISMANEVENTMIB_Mtetriggertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single trigger.  Applications create and delete entries
    // using mteTriggerEntryStatus. The type is slice of
    // DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry.
    Mtetriggerentry []DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry
}

func (mtetriggertable *DISMANEVENTMIB_Mtetriggertable) GetEntityData() *types.CommonEntityData {
    mtetriggertable.EntityData.YFilter = mtetriggertable.YFilter
    mtetriggertable.EntityData.YangName = "mteTriggerTable"
    mtetriggertable.EntityData.BundleName = "cisco_ios_xe"
    mtetriggertable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mtetriggertable.EntityData.SegmentPath = "mteTriggerTable"
    mtetriggertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggertable.EntityData.Children = make(map[string]types.YChild)
    mtetriggertable.EntityData.Children["mteTriggerEntry"] = types.YChild{"Mtetriggerentry", nil}
    for i := range mtetriggertable.Mtetriggerentry {
        mtetriggertable.EntityData.Children[types.GetSegmentPath(&mtetriggertable.Mtetriggerentry[i])] = types.YChild{"Mtetriggerentry", &mtetriggertable.Mtetriggerentry[i]}
    }
    mtetriggertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mtetriggertable.EntityData)
}

// DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry
// Information about a single trigger.  Applications create and
// delete entries using mteTriggerEntryStatus.
type DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (mtetriggerentry *DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry) GetEntityData() *types.CommonEntityData {
    mtetriggerentry.EntityData.YFilter = mtetriggerentry.YFilter
    mtetriggerentry.EntityData.YangName = "mteTriggerEntry"
    mtetriggerentry.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerentry.EntityData.ParentYangName = "mteTriggerTable"
    mtetriggerentry.EntityData.SegmentPath = "mteTriggerEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerentry.Mtetriggername) + "']"
    mtetriggerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerentry.EntityData.Children = make(map[string]types.YChild)
    mtetriggerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mtetriggerentry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mtetriggerentry.Mteowner}
    mtetriggerentry.EntityData.Leafs["mteTriggerName"] = types.YLeaf{"Mtetriggername", mtetriggerentry.Mtetriggername}
    mtetriggerentry.EntityData.Leafs["mteTriggerComment"] = types.YLeaf{"Mtetriggercomment", mtetriggerentry.Mtetriggercomment}
    mtetriggerentry.EntityData.Leafs["mteTriggerTest"] = types.YLeaf{"Mtetriggertest", mtetriggerentry.Mtetriggertest}
    mtetriggerentry.EntityData.Leafs["mteTriggerSampleType"] = types.YLeaf{"Mtetriggersampletype", mtetriggerentry.Mtetriggersampletype}
    mtetriggerentry.EntityData.Leafs["mteTriggerValueID"] = types.YLeaf{"Mtetriggervalueid", mtetriggerentry.Mtetriggervalueid}
    mtetriggerentry.EntityData.Leafs["mteTriggerValueIDWildcard"] = types.YLeaf{"Mtetriggervalueidwildcard", mtetriggerentry.Mtetriggervalueidwildcard}
    mtetriggerentry.EntityData.Leafs["mteTriggerTargetTag"] = types.YLeaf{"Mtetriggertargettag", mtetriggerentry.Mtetriggertargettag}
    mtetriggerentry.EntityData.Leafs["mteTriggerContextName"] = types.YLeaf{"Mtetriggercontextname", mtetriggerentry.Mtetriggercontextname}
    mtetriggerentry.EntityData.Leafs["mteTriggerContextNameWildcard"] = types.YLeaf{"Mtetriggercontextnamewildcard", mtetriggerentry.Mtetriggercontextnamewildcard}
    mtetriggerentry.EntityData.Leafs["mteTriggerFrequency"] = types.YLeaf{"Mtetriggerfrequency", mtetriggerentry.Mtetriggerfrequency}
    mtetriggerentry.EntityData.Leafs["mteTriggerObjectsOwner"] = types.YLeaf{"Mtetriggerobjectsowner", mtetriggerentry.Mtetriggerobjectsowner}
    mtetriggerentry.EntityData.Leafs["mteTriggerObjects"] = types.YLeaf{"Mtetriggerobjects", mtetriggerentry.Mtetriggerobjects}
    mtetriggerentry.EntityData.Leafs["mteTriggerEnabled"] = types.YLeaf{"Mtetriggerenabled", mtetriggerentry.Mtetriggerenabled}
    mtetriggerentry.EntityData.Leafs["mteTriggerEntryStatus"] = types.YLeaf{"Mtetriggerentrystatus", mtetriggerentry.Mtetriggerentrystatus}
    return &(mtetriggerentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single trigger's delta sampling.  Entries automatically
    // exist in this this table for each mteTriggerEntry that has
    // mteTriggerSampleType set to 'deltaValue'. The type is slice of
    // DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry.
    Mtetriggerdeltaentry []DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry
}

func (mtetriggerdeltatable *DISMANEVENTMIB_Mtetriggerdeltatable) GetEntityData() *types.CommonEntityData {
    mtetriggerdeltatable.EntityData.YFilter = mtetriggerdeltatable.YFilter
    mtetriggerdeltatable.EntityData.YangName = "mteTriggerDeltaTable"
    mtetriggerdeltatable.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerdeltatable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mtetriggerdeltatable.EntityData.SegmentPath = "mteTriggerDeltaTable"
    mtetriggerdeltatable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerdeltatable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerdeltatable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerdeltatable.EntityData.Children = make(map[string]types.YChild)
    mtetriggerdeltatable.EntityData.Children["mteTriggerDeltaEntry"] = types.YChild{"Mtetriggerdeltaentry", nil}
    for i := range mtetriggerdeltatable.Mtetriggerdeltaentry {
        mtetriggerdeltatable.EntityData.Children[types.GetSegmentPath(&mtetriggerdeltatable.Mtetriggerdeltaentry[i])] = types.YChild{"Mtetriggerdeltaentry", &mtetriggerdeltatable.Mtetriggerdeltaentry[i]}
    }
    mtetriggerdeltatable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mtetriggerdeltatable.EntityData)
}

// DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry
// Information about a single trigger's delta sampling.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has mteTriggerSampleType set to 'deltaValue'.
type DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (mtetriggerdeltaentry *DISMANEVENTMIB_Mtetriggerdeltatable_Mtetriggerdeltaentry) GetEntityData() *types.CommonEntityData {
    mtetriggerdeltaentry.EntityData.YFilter = mtetriggerdeltaentry.YFilter
    mtetriggerdeltaentry.EntityData.YangName = "mteTriggerDeltaEntry"
    mtetriggerdeltaentry.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerdeltaentry.EntityData.ParentYangName = "mteTriggerDeltaTable"
    mtetriggerdeltaentry.EntityData.SegmentPath = "mteTriggerDeltaEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerdeltaentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerdeltaentry.Mtetriggername) + "']"
    mtetriggerdeltaentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerdeltaentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerdeltaentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerdeltaentry.EntityData.Children = make(map[string]types.YChild)
    mtetriggerdeltaentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mtetriggerdeltaentry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mtetriggerdeltaentry.Mteowner}
    mtetriggerdeltaentry.EntityData.Leafs["mteTriggerName"] = types.YLeaf{"Mtetriggername", mtetriggerdeltaentry.Mtetriggername}
    mtetriggerdeltaentry.EntityData.Leafs["mteTriggerDeltaDiscontinuityID"] = types.YLeaf{"Mtetriggerdeltadiscontinuityid", mtetriggerdeltaentry.Mtetriggerdeltadiscontinuityid}
    mtetriggerdeltaentry.EntityData.Leafs["mteTriggerDeltaDiscontinuityIDWildcard"] = types.YLeaf{"Mtetriggerdeltadiscontinuityidwildcard", mtetriggerdeltaentry.Mtetriggerdeltadiscontinuityidwildcard}
    mtetriggerdeltaentry.EntityData.Leafs["mteTriggerDeltaDiscontinuityIDType"] = types.YLeaf{"Mtetriggerdeltadiscontinuityidtype", mtetriggerdeltaentry.Mtetriggerdeltadiscontinuityidtype}
    return &(mtetriggerdeltaentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single existence trigger.  Entries automatically exist
    // in this this table for each mteTriggerEntry that has 'existence' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry.
    Mtetriggerexistenceentry []DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry
}

func (mtetriggerexistencetable *DISMANEVENTMIB_Mtetriggerexistencetable) GetEntityData() *types.CommonEntityData {
    mtetriggerexistencetable.EntityData.YFilter = mtetriggerexistencetable.YFilter
    mtetriggerexistencetable.EntityData.YangName = "mteTriggerExistenceTable"
    mtetriggerexistencetable.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerexistencetable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mtetriggerexistencetable.EntityData.SegmentPath = "mteTriggerExistenceTable"
    mtetriggerexistencetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerexistencetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerexistencetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerexistencetable.EntityData.Children = make(map[string]types.YChild)
    mtetriggerexistencetable.EntityData.Children["mteTriggerExistenceEntry"] = types.YChild{"Mtetriggerexistenceentry", nil}
    for i := range mtetriggerexistencetable.Mtetriggerexistenceentry {
        mtetriggerexistencetable.EntityData.Children[types.GetSegmentPath(&mtetriggerexistencetable.Mtetriggerexistenceentry[i])] = types.YChild{"Mtetriggerexistenceentry", &mtetriggerexistencetable.Mtetriggerexistenceentry[i]}
    }
    mtetriggerexistencetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mtetriggerexistencetable.EntityData)
}

// DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry
// Information about a single existence trigger.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has 'existence' set in mteTriggerTest.
type DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry struct {
    EntityData types.CommonEntityData
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

func (mtetriggerexistenceentry *DISMANEVENTMIB_Mtetriggerexistencetable_Mtetriggerexistenceentry) GetEntityData() *types.CommonEntityData {
    mtetriggerexistenceentry.EntityData.YFilter = mtetriggerexistenceentry.YFilter
    mtetriggerexistenceentry.EntityData.YangName = "mteTriggerExistenceEntry"
    mtetriggerexistenceentry.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerexistenceentry.EntityData.ParentYangName = "mteTriggerExistenceTable"
    mtetriggerexistenceentry.EntityData.SegmentPath = "mteTriggerExistenceEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerexistenceentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerexistenceentry.Mtetriggername) + "']"
    mtetriggerexistenceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerexistenceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerexistenceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerexistenceentry.EntityData.Children = make(map[string]types.YChild)
    mtetriggerexistenceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mtetriggerexistenceentry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mtetriggerexistenceentry.Mteowner}
    mtetriggerexistenceentry.EntityData.Leafs["mteTriggerName"] = types.YLeaf{"Mtetriggername", mtetriggerexistenceentry.Mtetriggername}
    mtetriggerexistenceentry.EntityData.Leafs["mteTriggerExistenceTest"] = types.YLeaf{"Mtetriggerexistencetest", mtetriggerexistenceentry.Mtetriggerexistencetest}
    mtetriggerexistenceentry.EntityData.Leafs["mteTriggerExistenceStartup"] = types.YLeaf{"Mtetriggerexistencestartup", mtetriggerexistenceentry.Mtetriggerexistencestartup}
    mtetriggerexistenceentry.EntityData.Leafs["mteTriggerExistenceObjectsOwner"] = types.YLeaf{"Mtetriggerexistenceobjectsowner", mtetriggerexistenceentry.Mtetriggerexistenceobjectsowner}
    mtetriggerexistenceentry.EntityData.Leafs["mteTriggerExistenceObjects"] = types.YLeaf{"Mtetriggerexistenceobjects", mtetriggerexistenceentry.Mtetriggerexistenceobjects}
    mtetriggerexistenceentry.EntityData.Leafs["mteTriggerExistenceEventOwner"] = types.YLeaf{"Mtetriggerexistenceeventowner", mtetriggerexistenceentry.Mtetriggerexistenceeventowner}
    mtetriggerexistenceentry.EntityData.Leafs["mteTriggerExistenceEvent"] = types.YLeaf{"Mtetriggerexistenceevent", mtetriggerexistenceentry.Mtetriggerexistenceevent}
    return &(mtetriggerexistenceentry.EntityData)
}

// DISMANEVENTMIB_Mtetriggerbooleantable
// A table of management event trigger information for boolean
// triggers.
type DISMANEVENTMIB_Mtetriggerbooleantable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single boolean trigger.  Entries automatically exist in
    // this this table for each mteTriggerEntry that has 'boolean' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry.
    Mtetriggerbooleanentry []DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry
}

func (mtetriggerbooleantable *DISMANEVENTMIB_Mtetriggerbooleantable) GetEntityData() *types.CommonEntityData {
    mtetriggerbooleantable.EntityData.YFilter = mtetriggerbooleantable.YFilter
    mtetriggerbooleantable.EntityData.YangName = "mteTriggerBooleanTable"
    mtetriggerbooleantable.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerbooleantable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mtetriggerbooleantable.EntityData.SegmentPath = "mteTriggerBooleanTable"
    mtetriggerbooleantable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerbooleantable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerbooleantable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerbooleantable.EntityData.Children = make(map[string]types.YChild)
    mtetriggerbooleantable.EntityData.Children["mteTriggerBooleanEntry"] = types.YChild{"Mtetriggerbooleanentry", nil}
    for i := range mtetriggerbooleantable.Mtetriggerbooleanentry {
        mtetriggerbooleantable.EntityData.Children[types.GetSegmentPath(&mtetriggerbooleantable.Mtetriggerbooleanentry[i])] = types.YChild{"Mtetriggerbooleanentry", &mtetriggerbooleantable.Mtetriggerbooleanentry[i]}
    }
    mtetriggerbooleantable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mtetriggerbooleantable.EntityData)
}

// DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry
// Information about a single boolean trigger.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has 'boolean' set in mteTriggerTest.
type DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry struct {
    EntityData types.CommonEntityData
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

func (mtetriggerbooleanentry *DISMANEVENTMIB_Mtetriggerbooleantable_Mtetriggerbooleanentry) GetEntityData() *types.CommonEntityData {
    mtetriggerbooleanentry.EntityData.YFilter = mtetriggerbooleanentry.YFilter
    mtetriggerbooleanentry.EntityData.YangName = "mteTriggerBooleanEntry"
    mtetriggerbooleanentry.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerbooleanentry.EntityData.ParentYangName = "mteTriggerBooleanTable"
    mtetriggerbooleanentry.EntityData.SegmentPath = "mteTriggerBooleanEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerbooleanentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerbooleanentry.Mtetriggername) + "']"
    mtetriggerbooleanentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerbooleanentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerbooleanentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerbooleanentry.EntityData.Children = make(map[string]types.YChild)
    mtetriggerbooleanentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mtetriggerbooleanentry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mtetriggerbooleanentry.Mteowner}
    mtetriggerbooleanentry.EntityData.Leafs["mteTriggerName"] = types.YLeaf{"Mtetriggername", mtetriggerbooleanentry.Mtetriggername}
    mtetriggerbooleanentry.EntityData.Leafs["mteTriggerBooleanComparison"] = types.YLeaf{"Mtetriggerbooleancomparison", mtetriggerbooleanentry.Mtetriggerbooleancomparison}
    mtetriggerbooleanentry.EntityData.Leafs["mteTriggerBooleanValue"] = types.YLeaf{"Mtetriggerbooleanvalue", mtetriggerbooleanentry.Mtetriggerbooleanvalue}
    mtetriggerbooleanentry.EntityData.Leafs["mteTriggerBooleanStartup"] = types.YLeaf{"Mtetriggerbooleanstartup", mtetriggerbooleanentry.Mtetriggerbooleanstartup}
    mtetriggerbooleanentry.EntityData.Leafs["mteTriggerBooleanObjectsOwner"] = types.YLeaf{"Mtetriggerbooleanobjectsowner", mtetriggerbooleanentry.Mtetriggerbooleanobjectsowner}
    mtetriggerbooleanentry.EntityData.Leafs["mteTriggerBooleanObjects"] = types.YLeaf{"Mtetriggerbooleanobjects", mtetriggerbooleanentry.Mtetriggerbooleanobjects}
    mtetriggerbooleanentry.EntityData.Leafs["mteTriggerBooleanEventOwner"] = types.YLeaf{"Mtetriggerbooleaneventowner", mtetriggerbooleanentry.Mtetriggerbooleaneventowner}
    mtetriggerbooleanentry.EntityData.Leafs["mteTriggerBooleanEvent"] = types.YLeaf{"Mtetriggerbooleanevent", mtetriggerbooleanentry.Mtetriggerbooleanevent}
    return &(mtetriggerbooleanentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single threshold trigger.  Entries automatically exist
    // in this table for each mteTriggerEntry that has 'threshold' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry.
    Mtetriggerthresholdentry []DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry
}

func (mtetriggerthresholdtable *DISMANEVENTMIB_Mtetriggerthresholdtable) GetEntityData() *types.CommonEntityData {
    mtetriggerthresholdtable.EntityData.YFilter = mtetriggerthresholdtable.YFilter
    mtetriggerthresholdtable.EntityData.YangName = "mteTriggerThresholdTable"
    mtetriggerthresholdtable.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerthresholdtable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mtetriggerthresholdtable.EntityData.SegmentPath = "mteTriggerThresholdTable"
    mtetriggerthresholdtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerthresholdtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerthresholdtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerthresholdtable.EntityData.Children = make(map[string]types.YChild)
    mtetriggerthresholdtable.EntityData.Children["mteTriggerThresholdEntry"] = types.YChild{"Mtetriggerthresholdentry", nil}
    for i := range mtetriggerthresholdtable.Mtetriggerthresholdentry {
        mtetriggerthresholdtable.EntityData.Children[types.GetSegmentPath(&mtetriggerthresholdtable.Mtetriggerthresholdentry[i])] = types.YChild{"Mtetriggerthresholdentry", &mtetriggerthresholdtable.Mtetriggerthresholdentry[i]}
    }
    mtetriggerthresholdtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mtetriggerthresholdtable.EntityData)
}

// DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry
// Information about a single threshold trigger.  Entries
// automatically exist in this table for each mteTriggerEntry
// that has 'threshold' set in mteTriggerTest.
type DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry struct {
    EntityData types.CommonEntityData
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

func (mtetriggerthresholdentry *DISMANEVENTMIB_Mtetriggerthresholdtable_Mtetriggerthresholdentry) GetEntityData() *types.CommonEntityData {
    mtetriggerthresholdentry.EntityData.YFilter = mtetriggerthresholdentry.YFilter
    mtetriggerthresholdentry.EntityData.YangName = "mteTriggerThresholdEntry"
    mtetriggerthresholdentry.EntityData.BundleName = "cisco_ios_xe"
    mtetriggerthresholdentry.EntityData.ParentYangName = "mteTriggerThresholdTable"
    mtetriggerthresholdentry.EntityData.SegmentPath = "mteTriggerThresholdEntry" + "[mteOwner='" + fmt.Sprintf("%v", mtetriggerthresholdentry.Mteowner) + "']" + "[mteTriggerName='" + fmt.Sprintf("%v", mtetriggerthresholdentry.Mtetriggername) + "']"
    mtetriggerthresholdentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mtetriggerthresholdentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mtetriggerthresholdentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mtetriggerthresholdentry.EntityData.Children = make(map[string]types.YChild)
    mtetriggerthresholdentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mtetriggerthresholdentry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mtetriggerthresholdentry.Mteowner}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerName"] = types.YLeaf{"Mtetriggername", mtetriggerthresholdentry.Mtetriggername}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdStartup"] = types.YLeaf{"Mtetriggerthresholdstartup", mtetriggerthresholdentry.Mtetriggerthresholdstartup}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdRising"] = types.YLeaf{"Mtetriggerthresholdrising", mtetriggerthresholdentry.Mtetriggerthresholdrising}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdFalling"] = types.YLeaf{"Mtetriggerthresholdfalling", mtetriggerthresholdentry.Mtetriggerthresholdfalling}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdDeltaRising"] = types.YLeaf{"Mtetriggerthresholddeltarising", mtetriggerthresholdentry.Mtetriggerthresholddeltarising}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdDeltaFalling"] = types.YLeaf{"Mtetriggerthresholddeltafalling", mtetriggerthresholdentry.Mtetriggerthresholddeltafalling}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdObjectsOwner"] = types.YLeaf{"Mtetriggerthresholdobjectsowner", mtetriggerthresholdentry.Mtetriggerthresholdobjectsowner}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdObjects"] = types.YLeaf{"Mtetriggerthresholdobjects", mtetriggerthresholdentry.Mtetriggerthresholdobjects}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdRisingEventOwner"] = types.YLeaf{"Mtetriggerthresholdrisingeventowner", mtetriggerthresholdentry.Mtetriggerthresholdrisingeventowner}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdRisingEvent"] = types.YLeaf{"Mtetriggerthresholdrisingevent", mtetriggerthresholdentry.Mtetriggerthresholdrisingevent}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdFallingEventOwner"] = types.YLeaf{"Mtetriggerthresholdfallingeventowner", mtetriggerthresholdentry.Mtetriggerthresholdfallingeventowner}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdFallingEvent"] = types.YLeaf{"Mtetriggerthresholdfallingevent", mtetriggerthresholdentry.Mtetriggerthresholdfallingevent}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdDeltaRisingEventOwner"] = types.YLeaf{"Mtetriggerthresholddeltarisingeventowner", mtetriggerthresholdentry.Mtetriggerthresholddeltarisingeventowner}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdDeltaRisingEvent"] = types.YLeaf{"Mtetriggerthresholddeltarisingevent", mtetriggerthresholdentry.Mtetriggerthresholddeltarisingevent}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdDeltaFallingEventOwner"] = types.YLeaf{"Mtetriggerthresholddeltafallingeventowner", mtetriggerthresholdentry.Mtetriggerthresholddeltafallingeventowner}
    mtetriggerthresholdentry.EntityData.Leafs["mteTriggerThresholdDeltaFallingEvent"] = types.YLeaf{"Mtetriggerthresholddeltafallingevent", mtetriggerthresholdentry.Mtetriggerthresholddeltafallingevent}
    return &(mtetriggerthresholdentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A group of objects.  Applications create and delete entries using
    // mteObjectsEntryStatus.  When adding objects to a notification they are
    // added in the lexical order of their index in this table.  Those associated
    // with a trigger come first, then trigger test, then event. The type is slice
    // of DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry.
    Mteobjectsentry []DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry
}

func (mteobjectstable *DISMANEVENTMIB_Mteobjectstable) GetEntityData() *types.CommonEntityData {
    mteobjectstable.EntityData.YFilter = mteobjectstable.YFilter
    mteobjectstable.EntityData.YangName = "mteObjectsTable"
    mteobjectstable.EntityData.BundleName = "cisco_ios_xe"
    mteobjectstable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteobjectstable.EntityData.SegmentPath = "mteObjectsTable"
    mteobjectstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteobjectstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteobjectstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteobjectstable.EntityData.Children = make(map[string]types.YChild)
    mteobjectstable.EntityData.Children["mteObjectsEntry"] = types.YChild{"Mteobjectsentry", nil}
    for i := range mteobjectstable.Mteobjectsentry {
        mteobjectstable.EntityData.Children[types.GetSegmentPath(&mteobjectstable.Mteobjectsentry[i])] = types.YChild{"Mteobjectsentry", &mteobjectstable.Mteobjectsentry[i]}
    }
    mteobjectstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mteobjectstable.EntityData)
}

// DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry
// A group of objects.  Applications create and delete entries
// using mteObjectsEntryStatus.
// 
// When adding objects to a notification they are added in the
// lexical order of their index in this table.  Those associated
// with a trigger come first, then trigger test, then event.
type DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Mteobjectsid interface{}

    // Control for whether mteObjectsID is to be treated as fully-specified or
    // wildcarded, with 'true' indicating wildcard. The type is bool.
    Mteobjectsidwildcard interface{}

    // The control that allows creation and deletion of entries. Once made active
    // an entry MAY not be modified except to delete it. The type is RowStatus.
    Mteobjectsentrystatus interface{}
}

func (mteobjectsentry *DISMANEVENTMIB_Mteobjectstable_Mteobjectsentry) GetEntityData() *types.CommonEntityData {
    mteobjectsentry.EntityData.YFilter = mteobjectsentry.YFilter
    mteobjectsentry.EntityData.YangName = "mteObjectsEntry"
    mteobjectsentry.EntityData.BundleName = "cisco_ios_xe"
    mteobjectsentry.EntityData.ParentYangName = "mteObjectsTable"
    mteobjectsentry.EntityData.SegmentPath = "mteObjectsEntry" + "[mteOwner='" + fmt.Sprintf("%v", mteobjectsentry.Mteowner) + "']" + "[mteObjectsName='" + fmt.Sprintf("%v", mteobjectsentry.Mteobjectsname) + "']" + "[mteObjectsIndex='" + fmt.Sprintf("%v", mteobjectsentry.Mteobjectsindex) + "']"
    mteobjectsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteobjectsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteobjectsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteobjectsentry.EntityData.Children = make(map[string]types.YChild)
    mteobjectsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mteobjectsentry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mteobjectsentry.Mteowner}
    mteobjectsentry.EntityData.Leafs["mteObjectsName"] = types.YLeaf{"Mteobjectsname", mteobjectsentry.Mteobjectsname}
    mteobjectsentry.EntityData.Leafs["mteObjectsIndex"] = types.YLeaf{"Mteobjectsindex", mteobjectsentry.Mteobjectsindex}
    mteobjectsentry.EntityData.Leafs["mteObjectsID"] = types.YLeaf{"Mteobjectsid", mteobjectsentry.Mteobjectsid}
    mteobjectsentry.EntityData.Leafs["mteObjectsIDWildcard"] = types.YLeaf{"Mteobjectsidwildcard", mteobjectsentry.Mteobjectsidwildcard}
    mteobjectsentry.EntityData.Leafs["mteObjectsEntryStatus"] = types.YLeaf{"Mteobjectsentrystatus", mteobjectsentry.Mteobjectsentrystatus}
    return &(mteobjectsentry.EntityData)
}

// DISMANEVENTMIB_Mteeventtable
// A table of management event action information.
type DISMANEVENTMIB_Mteeventtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single event.  Applications create and delete entries
    // using mteEventEntryStatus. The type is slice of
    // DISMANEVENTMIB_Mteeventtable_Mteevententry.
    Mteevententry []DISMANEVENTMIB_Mteeventtable_Mteevententry
}

func (mteeventtable *DISMANEVENTMIB_Mteeventtable) GetEntityData() *types.CommonEntityData {
    mteeventtable.EntityData.YFilter = mteeventtable.YFilter
    mteeventtable.EntityData.YangName = "mteEventTable"
    mteeventtable.EntityData.BundleName = "cisco_ios_xe"
    mteeventtable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteeventtable.EntityData.SegmentPath = "mteEventTable"
    mteeventtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteeventtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteeventtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteeventtable.EntityData.Children = make(map[string]types.YChild)
    mteeventtable.EntityData.Children["mteEventEntry"] = types.YChild{"Mteevententry", nil}
    for i := range mteeventtable.Mteevententry {
        mteeventtable.EntityData.Children[types.GetSegmentPath(&mteeventtable.Mteevententry[i])] = types.YChild{"Mteevententry", &mteeventtable.Mteevententry[i]}
    }
    mteeventtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mteeventtable.EntityData)
}

// DISMANEVENTMIB_Mteeventtable_Mteevententry
// Information about a single event.  Applications create and
// delete entries using mteEventEntryStatus.
type DISMANEVENTMIB_Mteeventtable_Mteevententry struct {
    EntityData types.CommonEntityData
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

func (mteevententry *DISMANEVENTMIB_Mteeventtable_Mteevententry) GetEntityData() *types.CommonEntityData {
    mteevententry.EntityData.YFilter = mteevententry.YFilter
    mteevententry.EntityData.YangName = "mteEventEntry"
    mteevententry.EntityData.BundleName = "cisco_ios_xe"
    mteevententry.EntityData.ParentYangName = "mteEventTable"
    mteevententry.EntityData.SegmentPath = "mteEventEntry" + "[mteOwner='" + fmt.Sprintf("%v", mteevententry.Mteowner) + "']" + "[mteEventName='" + fmt.Sprintf("%v", mteevententry.Mteeventname) + "']"
    mteevententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteevententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteevententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteevententry.EntityData.Children = make(map[string]types.YChild)
    mteevententry.EntityData.Leafs = make(map[string]types.YLeaf)
    mteevententry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mteevententry.Mteowner}
    mteevententry.EntityData.Leafs["mteEventName"] = types.YLeaf{"Mteeventname", mteevententry.Mteeventname}
    mteevententry.EntityData.Leafs["mteEventComment"] = types.YLeaf{"Mteeventcomment", mteevententry.Mteeventcomment}
    mteevententry.EntityData.Leafs["mteEventActions"] = types.YLeaf{"Mteeventactions", mteevententry.Mteeventactions}
    mteevententry.EntityData.Leafs["mteEventEnabled"] = types.YLeaf{"Mteeventenabled", mteevententry.Mteeventenabled}
    mteevententry.EntityData.Leafs["mteEventEntryStatus"] = types.YLeaf{"Mteevententrystatus", mteevententry.Mteevententrystatus}
    return &(mteevententry.EntityData)
}

// DISMANEVENTMIB_Mteeventnotificationtable
// A table of information about notifications to be sent as a
// consequence of management events.
type DISMANEVENTMIB_Mteeventnotificationtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single event's notification.  Entries automatically
    // exist in this this table for each mteEventEntry that has 'notification' set
    // in mteEventActions. The type is slice of
    // DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry.
    Mteeventnotificationentry []DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry
}

func (mteeventnotificationtable *DISMANEVENTMIB_Mteeventnotificationtable) GetEntityData() *types.CommonEntityData {
    mteeventnotificationtable.EntityData.YFilter = mteeventnotificationtable.YFilter
    mteeventnotificationtable.EntityData.YangName = "mteEventNotificationTable"
    mteeventnotificationtable.EntityData.BundleName = "cisco_ios_xe"
    mteeventnotificationtable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteeventnotificationtable.EntityData.SegmentPath = "mteEventNotificationTable"
    mteeventnotificationtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteeventnotificationtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteeventnotificationtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteeventnotificationtable.EntityData.Children = make(map[string]types.YChild)
    mteeventnotificationtable.EntityData.Children["mteEventNotificationEntry"] = types.YChild{"Mteeventnotificationentry", nil}
    for i := range mteeventnotificationtable.Mteeventnotificationentry {
        mteeventnotificationtable.EntityData.Children[types.GetSegmentPath(&mteeventnotificationtable.Mteeventnotificationentry[i])] = types.YChild{"Mteeventnotificationentry", &mteeventnotificationtable.Mteeventnotificationentry[i]}
    }
    mteeventnotificationtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mteeventnotificationtable.EntityData)
}

// DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry
// Information about a single event's notification.  Entries
// automatically exist in this this table for each mteEventEntry
// that has 'notification' set in mteEventActions.
type DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mtetriggertable_Mtetriggerentry_Mteowner
    Mteowner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_Mteeventtable_Mteevententry_Mteeventname
    Mteeventname interface{}

    // The object identifier from the NOTIFICATION-TYPE for the notification to
    // use if metEventActions has 'notification' set. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (mteeventnotificationentry *DISMANEVENTMIB_Mteeventnotificationtable_Mteeventnotificationentry) GetEntityData() *types.CommonEntityData {
    mteeventnotificationentry.EntityData.YFilter = mteeventnotificationentry.YFilter
    mteeventnotificationentry.EntityData.YangName = "mteEventNotificationEntry"
    mteeventnotificationentry.EntityData.BundleName = "cisco_ios_xe"
    mteeventnotificationentry.EntityData.ParentYangName = "mteEventNotificationTable"
    mteeventnotificationentry.EntityData.SegmentPath = "mteEventNotificationEntry" + "[mteOwner='" + fmt.Sprintf("%v", mteeventnotificationentry.Mteowner) + "']" + "[mteEventName='" + fmt.Sprintf("%v", mteeventnotificationentry.Mteeventname) + "']"
    mteeventnotificationentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteeventnotificationentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteeventnotificationentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteeventnotificationentry.EntityData.Children = make(map[string]types.YChild)
    mteeventnotificationentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mteeventnotificationentry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mteeventnotificationentry.Mteowner}
    mteeventnotificationentry.EntityData.Leafs["mteEventName"] = types.YLeaf{"Mteeventname", mteeventnotificationentry.Mteeventname}
    mteeventnotificationentry.EntityData.Leafs["mteEventNotification"] = types.YLeaf{"Mteeventnotification", mteeventnotificationentry.Mteeventnotification}
    mteeventnotificationentry.EntityData.Leafs["mteEventNotificationObjectsOwner"] = types.YLeaf{"Mteeventnotificationobjectsowner", mteeventnotificationentry.Mteeventnotificationobjectsowner}
    mteeventnotificationentry.EntityData.Leafs["mteEventNotificationObjects"] = types.YLeaf{"Mteeventnotificationobjects", mteeventnotificationentry.Mteeventnotificationobjects}
    return &(mteeventnotificationentry.EntityData)
}

// DISMANEVENTMIB_Mteeventsettable
// A table of management event action information.
type DISMANEVENTMIB_Mteeventsettable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single event's set option.  Entries automatically exist
    // in this this table for each mteEventEntry that has 'set' set in
    // mteEventActions. The type is slice of
    // DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry.
    Mteeventsetentry []DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry
}

func (mteeventsettable *DISMANEVENTMIB_Mteeventsettable) GetEntityData() *types.CommonEntityData {
    mteeventsettable.EntityData.YFilter = mteeventsettable.YFilter
    mteeventsettable.EntityData.YangName = "mteEventSetTable"
    mteeventsettable.EntityData.BundleName = "cisco_ios_xe"
    mteeventsettable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteeventsettable.EntityData.SegmentPath = "mteEventSetTable"
    mteeventsettable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteeventsettable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteeventsettable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteeventsettable.EntityData.Children = make(map[string]types.YChild)
    mteeventsettable.EntityData.Children["mteEventSetEntry"] = types.YChild{"Mteeventsetentry", nil}
    for i := range mteeventsettable.Mteeventsetentry {
        mteeventsettable.EntityData.Children[types.GetSegmentPath(&mteeventsettable.Mteeventsetentry[i])] = types.YChild{"Mteeventsetentry", &mteeventsettable.Mteeventsetentry[i]}
    }
    mteeventsettable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mteeventsettable.EntityData)
}

// DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry
// Information about a single event's set option.  Entries
// automatically exist in this this table for each mteEventEntry
// that has 'set' set in mteEventActions.
type DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry struct {
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (mteeventsetentry *DISMANEVENTMIB_Mteeventsettable_Mteeventsetentry) GetEntityData() *types.CommonEntityData {
    mteeventsetentry.EntityData.YFilter = mteeventsetentry.YFilter
    mteeventsetentry.EntityData.YangName = "mteEventSetEntry"
    mteeventsetentry.EntityData.BundleName = "cisco_ios_xe"
    mteeventsetentry.EntityData.ParentYangName = "mteEventSetTable"
    mteeventsetentry.EntityData.SegmentPath = "mteEventSetEntry" + "[mteOwner='" + fmt.Sprintf("%v", mteeventsetentry.Mteowner) + "']" + "[mteEventName='" + fmt.Sprintf("%v", mteeventsetentry.Mteeventname) + "']"
    mteeventsetentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteeventsetentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteeventsetentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteeventsetentry.EntityData.Children = make(map[string]types.YChild)
    mteeventsetentry.EntityData.Leafs = make(map[string]types.YLeaf)
    mteeventsetentry.EntityData.Leafs["mteOwner"] = types.YLeaf{"Mteowner", mteeventsetentry.Mteowner}
    mteeventsetentry.EntityData.Leafs["mteEventName"] = types.YLeaf{"Mteeventname", mteeventsetentry.Mteeventname}
    mteeventsetentry.EntityData.Leafs["mteEventSetObject"] = types.YLeaf{"Mteeventsetobject", mteeventsetentry.Mteeventsetobject}
    mteeventsetentry.EntityData.Leafs["mteEventSetObjectWildcard"] = types.YLeaf{"Mteeventsetobjectwildcard", mteeventsetentry.Mteeventsetobjectwildcard}
    mteeventsetentry.EntityData.Leafs["mteEventSetValue"] = types.YLeaf{"Mteeventsetvalue", mteeventsetentry.Mteeventsetvalue}
    mteeventsetentry.EntityData.Leafs["mteEventSetTargetTag"] = types.YLeaf{"Mteeventsettargettag", mteeventsetentry.Mteeventsettargettag}
    mteeventsetentry.EntityData.Leafs["mteEventSetContextName"] = types.YLeaf{"Mteeventsetcontextname", mteeventsetentry.Mteeventsetcontextname}
    mteeventsetentry.EntityData.Leafs["mteEventSetContextNameWildcard"] = types.YLeaf{"Mteeventsetcontextnamewildcard", mteeventsetentry.Mteeventsetcontextnamewildcard}
    return &(mteeventsetentry.EntityData)
}

