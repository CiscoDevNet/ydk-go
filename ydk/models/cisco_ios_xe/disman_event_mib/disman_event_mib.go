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

    
    MteResource DISMANEVENTMIB_MteResource

    
    MteTrigger DISMANEVENTMIB_MteTrigger

    
    MteEvent DISMANEVENTMIB_MteEvent

    // A table of management event trigger information.
    MteTriggerTable DISMANEVENTMIB_MteTriggerTable

    // A table of management event trigger information for delta sampling.
    MteTriggerDeltaTable DISMANEVENTMIB_MteTriggerDeltaTable

    // A table of management event trigger information for existence triggers.
    MteTriggerExistenceTable DISMANEVENTMIB_MteTriggerExistenceTable

    // A table of management event trigger information for boolean triggers.
    MteTriggerBooleanTable DISMANEVENTMIB_MteTriggerBooleanTable

    // A table of management event trigger information for threshold triggers.
    MteTriggerThresholdTable DISMANEVENTMIB_MteTriggerThresholdTable

    // A table of objects that can be added to notifications based on the trigger,
    // trigger test, or event, as pointed to by entries in those tables.
    MteObjectsTable DISMANEVENTMIB_MteObjectsTable

    // A table of management event action information.
    MteEventTable DISMANEVENTMIB_MteEventTable

    // A table of information about notifications to be sent as a consequence of
    // management events.
    MteEventNotificationTable DISMANEVENTMIB_MteEventNotificationTable

    // A table of management event action information.
    MteEventSetTable DISMANEVENTMIB_MteEventSetTable
}

func (dISMANEVENTMIB *DISMANEVENTMIB) GetEntityData() *types.CommonEntityData {
    dISMANEVENTMIB.EntityData.YFilter = dISMANEVENTMIB.YFilter
    dISMANEVENTMIB.EntityData.YangName = "DISMAN-EVENT-MIB"
    dISMANEVENTMIB.EntityData.BundleName = "cisco_ios_xe"
    dISMANEVENTMIB.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    dISMANEVENTMIB.EntityData.SegmentPath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB"
    dISMANEVENTMIB.EntityData.AbsolutePath = dISMANEVENTMIB.EntityData.SegmentPath
    dISMANEVENTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dISMANEVENTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dISMANEVENTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dISMANEVENTMIB.EntityData.Children = types.NewOrderedMap()
    dISMANEVENTMIB.EntityData.Children.Append("mteResource", types.YChild{"MteResource", &dISMANEVENTMIB.MteResource})
    dISMANEVENTMIB.EntityData.Children.Append("mteTrigger", types.YChild{"MteTrigger", &dISMANEVENTMIB.MteTrigger})
    dISMANEVENTMIB.EntityData.Children.Append("mteEvent", types.YChild{"MteEvent", &dISMANEVENTMIB.MteEvent})
    dISMANEVENTMIB.EntityData.Children.Append("mteTriggerTable", types.YChild{"MteTriggerTable", &dISMANEVENTMIB.MteTriggerTable})
    dISMANEVENTMIB.EntityData.Children.Append("mteTriggerDeltaTable", types.YChild{"MteTriggerDeltaTable", &dISMANEVENTMIB.MteTriggerDeltaTable})
    dISMANEVENTMIB.EntityData.Children.Append("mteTriggerExistenceTable", types.YChild{"MteTriggerExistenceTable", &dISMANEVENTMIB.MteTriggerExistenceTable})
    dISMANEVENTMIB.EntityData.Children.Append("mteTriggerBooleanTable", types.YChild{"MteTriggerBooleanTable", &dISMANEVENTMIB.MteTriggerBooleanTable})
    dISMANEVENTMIB.EntityData.Children.Append("mteTriggerThresholdTable", types.YChild{"MteTriggerThresholdTable", &dISMANEVENTMIB.MteTriggerThresholdTable})
    dISMANEVENTMIB.EntityData.Children.Append("mteObjectsTable", types.YChild{"MteObjectsTable", &dISMANEVENTMIB.MteObjectsTable})
    dISMANEVENTMIB.EntityData.Children.Append("mteEventTable", types.YChild{"MteEventTable", &dISMANEVENTMIB.MteEventTable})
    dISMANEVENTMIB.EntityData.Children.Append("mteEventNotificationTable", types.YChild{"MteEventNotificationTable", &dISMANEVENTMIB.MteEventNotificationTable})
    dISMANEVENTMIB.EntityData.Children.Append("mteEventSetTable", types.YChild{"MteEventSetTable", &dISMANEVENTMIB.MteEventSetTable})
    dISMANEVENTMIB.EntityData.Leafs = types.NewOrderedMap()

    dISMANEVENTMIB.EntityData.YListKeys = []string {}

    return &(dISMANEVENTMIB.EntityData)
}

// DISMANEVENTMIB_MteResource
type DISMANEVENTMIB_MteResource struct {
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
    MteResourceSampleMinimum interface{}

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
    MteResourceSampleInstanceMaximum interface{}

    // The number of currently active instance entries as defined for
    // mteResourceSampleInstanceMaximum. The type is interface{} with range:
    // 0..4294967295. Units are instances.
    MteResourceSampleInstances interface{}

    // The highest value of mteResourceSampleInstances that has occurred since
    // initialization of the management system. The type is interface{} with
    // range: 0..4294967295. Units are instances.
    MteResourceSampleInstancesHigh interface{}

    // The number of times this system could not take a new sample because that
    // allocation would have exceeded the limit set by
    // mteResourceSampleInstanceMaximum. The type is interface{} with range:
    // 0..4294967295. Units are instances.
    MteResourceSampleInstanceLacks interface{}
}

func (mteResource *DISMANEVENTMIB_MteResource) GetEntityData() *types.CommonEntityData {
    mteResource.EntityData.YFilter = mteResource.YFilter
    mteResource.EntityData.YangName = "mteResource"
    mteResource.EntityData.BundleName = "cisco_ios_xe"
    mteResource.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteResource.EntityData.SegmentPath = "mteResource"
    mteResource.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteResource.EntityData.SegmentPath
    mteResource.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteResource.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteResource.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteResource.EntityData.Children = types.NewOrderedMap()
    mteResource.EntityData.Leafs = types.NewOrderedMap()
    mteResource.EntityData.Leafs.Append("mteResourceSampleMinimum", types.YLeaf{"MteResourceSampleMinimum", mteResource.MteResourceSampleMinimum})
    mteResource.EntityData.Leafs.Append("mteResourceSampleInstanceMaximum", types.YLeaf{"MteResourceSampleInstanceMaximum", mteResource.MteResourceSampleInstanceMaximum})
    mteResource.EntityData.Leafs.Append("mteResourceSampleInstances", types.YLeaf{"MteResourceSampleInstances", mteResource.MteResourceSampleInstances})
    mteResource.EntityData.Leafs.Append("mteResourceSampleInstancesHigh", types.YLeaf{"MteResourceSampleInstancesHigh", mteResource.MteResourceSampleInstancesHigh})
    mteResource.EntityData.Leafs.Append("mteResourceSampleInstanceLacks", types.YLeaf{"MteResourceSampleInstanceLacks", mteResource.MteResourceSampleInstanceLacks})

    mteResource.EntityData.YListKeys = []string {}

    return &(mteResource.EntityData)
}

// DISMANEVENTMIB_MteTrigger
type DISMANEVENTMIB_MteTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times an attempt to check for a trigger condition has failed.
    // This counts individually for each attempt in a group of targets or each
    // attempt for a  wildcarded object. The type is interface{} with range:
    // 0..4294967295. Units are failures.
    MteTriggerFailures interface{}
}

func (mteTrigger *DISMANEVENTMIB_MteTrigger) GetEntityData() *types.CommonEntityData {
    mteTrigger.EntityData.YFilter = mteTrigger.YFilter
    mteTrigger.EntityData.YangName = "mteTrigger"
    mteTrigger.EntityData.BundleName = "cisco_ios_xe"
    mteTrigger.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteTrigger.EntityData.SegmentPath = "mteTrigger"
    mteTrigger.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteTrigger.EntityData.SegmentPath
    mteTrigger.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTrigger.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTrigger.EntityData.Children = types.NewOrderedMap()
    mteTrigger.EntityData.Leafs = types.NewOrderedMap()
    mteTrigger.EntityData.Leafs.Append("mteTriggerFailures", types.YLeaf{"MteTriggerFailures", mteTrigger.MteTriggerFailures})

    mteTrigger.EntityData.YListKeys = []string {}

    return &(mteTrigger.EntityData)
}

// DISMANEVENTMIB_MteEvent
type DISMANEVENTMIB_MteEvent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of times an attempt to invoke an event has failed.  This counts
    // individually for each attempt in a group of targets or each attempt for a
    // wildcarded trigger object. The type is interface{} with range:
    // 0..4294967295.
    MteEventFailures interface{}
}

func (mteEvent *DISMANEVENTMIB_MteEvent) GetEntityData() *types.CommonEntityData {
    mteEvent.EntityData.YFilter = mteEvent.YFilter
    mteEvent.EntityData.YangName = "mteEvent"
    mteEvent.EntityData.BundleName = "cisco_ios_xe"
    mteEvent.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteEvent.EntityData.SegmentPath = "mteEvent"
    mteEvent.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteEvent.EntityData.SegmentPath
    mteEvent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteEvent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteEvent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteEvent.EntityData.Children = types.NewOrderedMap()
    mteEvent.EntityData.Leafs = types.NewOrderedMap()
    mteEvent.EntityData.Leafs.Append("mteEventFailures", types.YLeaf{"MteEventFailures", mteEvent.MteEventFailures})

    mteEvent.EntityData.YListKeys = []string {}

    return &(mteEvent.EntityData)
}

// DISMANEVENTMIB_MteTriggerTable
// A table of management event trigger information.
type DISMANEVENTMIB_MteTriggerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single trigger.  Applications create and delete entries
    // using mteTriggerEntryStatus. The type is slice of
    // DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry.
    MteTriggerEntry []*DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry
}

func (mteTriggerTable *DISMANEVENTMIB_MteTriggerTable) GetEntityData() *types.CommonEntityData {
    mteTriggerTable.EntityData.YFilter = mteTriggerTable.YFilter
    mteTriggerTable.EntityData.YangName = "mteTriggerTable"
    mteTriggerTable.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteTriggerTable.EntityData.SegmentPath = "mteTriggerTable"
    mteTriggerTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteTriggerTable.EntityData.SegmentPath
    mteTriggerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerTable.EntityData.Children = types.NewOrderedMap()
    mteTriggerTable.EntityData.Children.Append("mteTriggerEntry", types.YChild{"MteTriggerEntry", nil})
    for i := range mteTriggerTable.MteTriggerEntry {
        mteTriggerTable.EntityData.Children.Append(types.GetSegmentPath(mteTriggerTable.MteTriggerEntry[i]), types.YChild{"MteTriggerEntry", mteTriggerTable.MteTriggerEntry[i]})
    }
    mteTriggerTable.EntityData.Leafs = types.NewOrderedMap()

    mteTriggerTable.EntityData.YListKeys = []string {}

    return &(mteTriggerTable.EntityData)
}

// DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry
// Information about a single trigger.  Applications create and
// delete entries using mteTriggerEntryStatus.
type DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The owner of this entry. The exact semantics of
    // this string are subject to the security policy defined by the security
    // administrator. The type is string with length: 0..32.
    MteOwner interface{}

    // This attribute is a key. A locally-unique, administratively assigned name
    // for the trigger within the scope of mteOwner. The type is string with
    // length: 1..32.
    MteTriggerName interface{}

    // A description of the trigger's function and use. The type is string.
    MteTriggerComment interface{}

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
    MteTriggerTest interface{}

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
    // this object has no meaning. The type is MteTriggerSampleType.
    MteTriggerSampleType interface{}

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
    MteTriggerValueID interface{}

    // Control for whether mteTriggerValueID is to be treated as fully-specified
    // or wildcarded, with 'true' indicating wildcard. The type is bool.
    MteTriggerValueIDWildcard interface{}

    // The tag for the target(s) from which to obtain the condition for a trigger
    // check.  A length of 0 indicates the local system.  In this case, access to
    // the objects indicated by mteTriggerValueID is under the security
    // credentials of the requester that set mteTriggerEntryStatus to 'active'. 
    // Those credentials are the input parameters for isAccessAllowed from the
    // Architecture for Describing SNMP Management Frameworks.  Otherwise access
    // rights are checked according to the security  parameters resulting from the
    // tag. The type is string.
    MteTriggerTargetTag interface{}

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
    MteTriggerContextName interface{}

    // Control for whether mteTriggerContextName is to be treated as
    // fully-specified or wildcarded, with 'true' indicating wildcard. The type is
    // bool.
    MteTriggerContextNameWildcard interface{}

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
    MteTriggerFrequency interface{}

    // To go with mteTriggerObjects, the mteOwner of a group of objects from
    // mteObjectsTable. The type is string with length: 0..32.
    MteTriggerObjectsOwner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable.  These
    // objects are to be added to any Notification resulting from the firing of
    // this trigger.  A list of objects may also be added based on the event or on
    // the value of mteTriggerTest.  A length of 0 indicates no additional
    // objects. The type is string with length: 0..32.
    MteTriggerObjects interface{}

    // A control to allow a trigger to be configured but not used. When the value
    // is 'false' the trigger is not sampled. The type is bool.
    MteTriggerEnabled interface{}

    // The control that allows creation and deletion of entries. Once made active
    // an entry may not be modified except to delete it. The type is RowStatus.
    MteTriggerEntryStatus interface{}
}

func (mteTriggerEntry *DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry) GetEntityData() *types.CommonEntityData {
    mteTriggerEntry.EntityData.YFilter = mteTriggerEntry.YFilter
    mteTriggerEntry.EntityData.YangName = "mteTriggerEntry"
    mteTriggerEntry.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerEntry.EntityData.ParentYangName = "mteTriggerTable"
    mteTriggerEntry.EntityData.SegmentPath = "mteTriggerEntry" + types.AddKeyToken(mteTriggerEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteTriggerEntry.MteTriggerName, "mteTriggerName")
    mteTriggerEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteTriggerTable/" + mteTriggerEntry.EntityData.SegmentPath
    mteTriggerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerEntry.EntityData.Children = types.NewOrderedMap()
    mteTriggerEntry.EntityData.Leafs = types.NewOrderedMap()
    mteTriggerEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteTriggerEntry.MteOwner})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerName", types.YLeaf{"MteTriggerName", mteTriggerEntry.MteTriggerName})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerComment", types.YLeaf{"MteTriggerComment", mteTriggerEntry.MteTriggerComment})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerTest", types.YLeaf{"MteTriggerTest", mteTriggerEntry.MteTriggerTest})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerSampleType", types.YLeaf{"MteTriggerSampleType", mteTriggerEntry.MteTriggerSampleType})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerValueID", types.YLeaf{"MteTriggerValueID", mteTriggerEntry.MteTriggerValueID})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerValueIDWildcard", types.YLeaf{"MteTriggerValueIDWildcard", mteTriggerEntry.MteTriggerValueIDWildcard})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerTargetTag", types.YLeaf{"MteTriggerTargetTag", mteTriggerEntry.MteTriggerTargetTag})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerContextName", types.YLeaf{"MteTriggerContextName", mteTriggerEntry.MteTriggerContextName})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerContextNameWildcard", types.YLeaf{"MteTriggerContextNameWildcard", mteTriggerEntry.MteTriggerContextNameWildcard})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerFrequency", types.YLeaf{"MteTriggerFrequency", mteTriggerEntry.MteTriggerFrequency})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerObjectsOwner", types.YLeaf{"MteTriggerObjectsOwner", mteTriggerEntry.MteTriggerObjectsOwner})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerObjects", types.YLeaf{"MteTriggerObjects", mteTriggerEntry.MteTriggerObjects})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerEnabled", types.YLeaf{"MteTriggerEnabled", mteTriggerEntry.MteTriggerEnabled})
    mteTriggerEntry.EntityData.Leafs.Append("mteTriggerEntryStatus", types.YLeaf{"MteTriggerEntryStatus", mteTriggerEntry.MteTriggerEntryStatus})

    mteTriggerEntry.EntityData.YListKeys = []string {"MteOwner", "MteTriggerName"}

    return &(mteTriggerEntry.EntityData)
}

// DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerSampleType represents no meaning.
type DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerSampleType string

const (
    DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerSampleType_absoluteValue DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerSampleType = "absoluteValue"

    DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerSampleType_deltaValue DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerSampleType = "deltaValue"
)

// DISMANEVENTMIB_MteTriggerDeltaTable
// A table of management event trigger information for delta
// sampling.
type DISMANEVENTMIB_MteTriggerDeltaTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single trigger's delta sampling.  Entries automatically
    // exist in this this table for each mteTriggerEntry that has
    // mteTriggerSampleType set to 'deltaValue'. The type is slice of
    // DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry.
    MteTriggerDeltaEntry []*DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry
}

func (mteTriggerDeltaTable *DISMANEVENTMIB_MteTriggerDeltaTable) GetEntityData() *types.CommonEntityData {
    mteTriggerDeltaTable.EntityData.YFilter = mteTriggerDeltaTable.YFilter
    mteTriggerDeltaTable.EntityData.YangName = "mteTriggerDeltaTable"
    mteTriggerDeltaTable.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerDeltaTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteTriggerDeltaTable.EntityData.SegmentPath = "mteTriggerDeltaTable"
    mteTriggerDeltaTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteTriggerDeltaTable.EntityData.SegmentPath
    mteTriggerDeltaTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerDeltaTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerDeltaTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerDeltaTable.EntityData.Children = types.NewOrderedMap()
    mteTriggerDeltaTable.EntityData.Children.Append("mteTriggerDeltaEntry", types.YChild{"MteTriggerDeltaEntry", nil})
    for i := range mteTriggerDeltaTable.MteTriggerDeltaEntry {
        mteTriggerDeltaTable.EntityData.Children.Append(types.GetSegmentPath(mteTriggerDeltaTable.MteTriggerDeltaEntry[i]), types.YChild{"MteTriggerDeltaEntry", mteTriggerDeltaTable.MteTriggerDeltaEntry[i]})
    }
    mteTriggerDeltaTable.EntityData.Leafs = types.NewOrderedMap()

    mteTriggerDeltaTable.EntityData.YListKeys = []string {}

    return &(mteTriggerDeltaTable.EntityData)
}

// DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry
// Information about a single trigger's delta sampling.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has mteTriggerSampleType set to 'deltaValue'.
type DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteOwner
    MteOwner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerName
    MteTriggerName interface{}

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
    MteTriggerDeltaDiscontinuityID interface{}

    // Control for whether mteTriggerDeltaDiscontinuityID is to be treated as
    // fully-specified or wildcarded, with 'true' indicating wildcard. Note that
    // the value of this object will be the same as that of the corresponding
    // instance of mteTriggerValueIDWildcard when the corresponding 
    // mteTriggerSampleType is 'deltaValue'. The type is bool.
    MteTriggerDeltaDiscontinuityIDWildcard interface{}

    // The value 'timeTicks' indicates the mteTriggerDeltaDiscontinuityID of this
    // row is of syntax TimeTicks.  The value 'timeStamp' indicates syntax
    // TimeStamp. The value 'dateAndTime' indicates syntax DateAndTime. The type
    // is MteTriggerDeltaDiscontinuityIDType.
    MteTriggerDeltaDiscontinuityIDType interface{}
}

func (mteTriggerDeltaEntry *DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry) GetEntityData() *types.CommonEntityData {
    mteTriggerDeltaEntry.EntityData.YFilter = mteTriggerDeltaEntry.YFilter
    mteTriggerDeltaEntry.EntityData.YangName = "mteTriggerDeltaEntry"
    mteTriggerDeltaEntry.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerDeltaEntry.EntityData.ParentYangName = "mteTriggerDeltaTable"
    mteTriggerDeltaEntry.EntityData.SegmentPath = "mteTriggerDeltaEntry" + types.AddKeyToken(mteTriggerDeltaEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteTriggerDeltaEntry.MteTriggerName, "mteTriggerName")
    mteTriggerDeltaEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteTriggerDeltaTable/" + mteTriggerDeltaEntry.EntityData.SegmentPath
    mteTriggerDeltaEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerDeltaEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerDeltaEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerDeltaEntry.EntityData.Children = types.NewOrderedMap()
    mteTriggerDeltaEntry.EntityData.Leafs = types.NewOrderedMap()
    mteTriggerDeltaEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteTriggerDeltaEntry.MteOwner})
    mteTriggerDeltaEntry.EntityData.Leafs.Append("mteTriggerName", types.YLeaf{"MteTriggerName", mteTriggerDeltaEntry.MteTriggerName})
    mteTriggerDeltaEntry.EntityData.Leafs.Append("mteTriggerDeltaDiscontinuityID", types.YLeaf{"MteTriggerDeltaDiscontinuityID", mteTriggerDeltaEntry.MteTriggerDeltaDiscontinuityID})
    mteTriggerDeltaEntry.EntityData.Leafs.Append("mteTriggerDeltaDiscontinuityIDWildcard", types.YLeaf{"MteTriggerDeltaDiscontinuityIDWildcard", mteTriggerDeltaEntry.MteTriggerDeltaDiscontinuityIDWildcard})
    mteTriggerDeltaEntry.EntityData.Leafs.Append("mteTriggerDeltaDiscontinuityIDType", types.YLeaf{"MteTriggerDeltaDiscontinuityIDType", mteTriggerDeltaEntry.MteTriggerDeltaDiscontinuityIDType})

    mteTriggerDeltaEntry.EntityData.YListKeys = []string {"MteOwner", "MteTriggerName"}

    return &(mteTriggerDeltaEntry.EntityData)
}

// DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry_MteTriggerDeltaDiscontinuityIDType represents The value 'dateAndTime' indicates syntax DateAndTime.
type DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry_MteTriggerDeltaDiscontinuityIDType string

const (
    DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry_MteTriggerDeltaDiscontinuityIDType_timeTicks DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry_MteTriggerDeltaDiscontinuityIDType = "timeTicks"

    DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry_MteTriggerDeltaDiscontinuityIDType_timeStamp DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry_MteTriggerDeltaDiscontinuityIDType = "timeStamp"

    DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry_MteTriggerDeltaDiscontinuityIDType_dateAndTime DISMANEVENTMIB_MteTriggerDeltaTable_MteTriggerDeltaEntry_MteTriggerDeltaDiscontinuityIDType = "dateAndTime"
)

// DISMANEVENTMIB_MteTriggerExistenceTable
// A table of management event trigger information for existence
// triggers.
type DISMANEVENTMIB_MteTriggerExistenceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single existence trigger.  Entries automatically exist
    // in this this table for each mteTriggerEntry that has 'existence' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_MteTriggerExistenceTable_MteTriggerExistenceEntry.
    MteTriggerExistenceEntry []*DISMANEVENTMIB_MteTriggerExistenceTable_MteTriggerExistenceEntry
}

func (mteTriggerExistenceTable *DISMANEVENTMIB_MteTriggerExistenceTable) GetEntityData() *types.CommonEntityData {
    mteTriggerExistenceTable.EntityData.YFilter = mteTriggerExistenceTable.YFilter
    mteTriggerExistenceTable.EntityData.YangName = "mteTriggerExistenceTable"
    mteTriggerExistenceTable.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerExistenceTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteTriggerExistenceTable.EntityData.SegmentPath = "mteTriggerExistenceTable"
    mteTriggerExistenceTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteTriggerExistenceTable.EntityData.SegmentPath
    mteTriggerExistenceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerExistenceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerExistenceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerExistenceTable.EntityData.Children = types.NewOrderedMap()
    mteTriggerExistenceTable.EntityData.Children.Append("mteTriggerExistenceEntry", types.YChild{"MteTriggerExistenceEntry", nil})
    for i := range mteTriggerExistenceTable.MteTriggerExistenceEntry {
        mteTriggerExistenceTable.EntityData.Children.Append(types.GetSegmentPath(mteTriggerExistenceTable.MteTriggerExistenceEntry[i]), types.YChild{"MteTriggerExistenceEntry", mteTriggerExistenceTable.MteTriggerExistenceEntry[i]})
    }
    mteTriggerExistenceTable.EntityData.Leafs = types.NewOrderedMap()

    mteTriggerExistenceTable.EntityData.YListKeys = []string {}

    return &(mteTriggerExistenceTable.EntityData)
}

// DISMANEVENTMIB_MteTriggerExistenceTable_MteTriggerExistenceEntry
// Information about a single existence trigger.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has 'existence' set in mteTriggerTest.
type DISMANEVENTMIB_MteTriggerExistenceTable_MteTriggerExistenceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteOwner
    MteOwner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerName
    MteTriggerName interface{}

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
    MteTriggerExistenceTest interface{}

    // Control for whether an event may be triggered when this entry is first set
    // to 'active' and the test specified by mteTriggerExistenceTest is true. 
    // Setting an option causes that trigger to fire when its test is true. The
    // type is map[string]bool.
    MteTriggerExistenceStartup interface{}

    // To go with mteTriggerExistenceObjects, the mteOwner of a group of objects
    // from mteObjectsTable. The type is string with length: 0..32.
    MteTriggerExistenceObjectsOwner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable.  These
    // objects are to be added to any Notification resulting from the firing of
    // this trigger for this test.  A list of objects may also be added based on
    // the overall trigger, the event or other settings in mteTriggerTest.  A
    // length of 0 indicates no additional objects. The type is string with
    // length: 0..32.
    MteTriggerExistenceObjects interface{}

    // To go with mteTriggerExistenceEvent, the mteOwner of an event entry from
    // the mteEventTable. The type is string with length: 0..32.
    MteTriggerExistenceEventOwner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'existence'
    // and this trigger fires.  A length of 0 indicates no event. The type is
    // string with length: 0..32.
    MteTriggerExistenceEvent interface{}
}

func (mteTriggerExistenceEntry *DISMANEVENTMIB_MteTriggerExistenceTable_MteTriggerExistenceEntry) GetEntityData() *types.CommonEntityData {
    mteTriggerExistenceEntry.EntityData.YFilter = mteTriggerExistenceEntry.YFilter
    mteTriggerExistenceEntry.EntityData.YangName = "mteTriggerExistenceEntry"
    mteTriggerExistenceEntry.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerExistenceEntry.EntityData.ParentYangName = "mteTriggerExistenceTable"
    mteTriggerExistenceEntry.EntityData.SegmentPath = "mteTriggerExistenceEntry" + types.AddKeyToken(mteTriggerExistenceEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteTriggerExistenceEntry.MteTriggerName, "mteTriggerName")
    mteTriggerExistenceEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteTriggerExistenceTable/" + mteTriggerExistenceEntry.EntityData.SegmentPath
    mteTriggerExistenceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerExistenceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerExistenceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerExistenceEntry.EntityData.Children = types.NewOrderedMap()
    mteTriggerExistenceEntry.EntityData.Leafs = types.NewOrderedMap()
    mteTriggerExistenceEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteTriggerExistenceEntry.MteOwner})
    mteTriggerExistenceEntry.EntityData.Leafs.Append("mteTriggerName", types.YLeaf{"MteTriggerName", mteTriggerExistenceEntry.MteTriggerName})
    mteTriggerExistenceEntry.EntityData.Leafs.Append("mteTriggerExistenceTest", types.YLeaf{"MteTriggerExistenceTest", mteTriggerExistenceEntry.MteTriggerExistenceTest})
    mteTriggerExistenceEntry.EntityData.Leafs.Append("mteTriggerExistenceStartup", types.YLeaf{"MteTriggerExistenceStartup", mteTriggerExistenceEntry.MteTriggerExistenceStartup})
    mteTriggerExistenceEntry.EntityData.Leafs.Append("mteTriggerExistenceObjectsOwner", types.YLeaf{"MteTriggerExistenceObjectsOwner", mteTriggerExistenceEntry.MteTriggerExistenceObjectsOwner})
    mteTriggerExistenceEntry.EntityData.Leafs.Append("mteTriggerExistenceObjects", types.YLeaf{"MteTriggerExistenceObjects", mteTriggerExistenceEntry.MteTriggerExistenceObjects})
    mteTriggerExistenceEntry.EntityData.Leafs.Append("mteTriggerExistenceEventOwner", types.YLeaf{"MteTriggerExistenceEventOwner", mteTriggerExistenceEntry.MteTriggerExistenceEventOwner})
    mteTriggerExistenceEntry.EntityData.Leafs.Append("mteTriggerExistenceEvent", types.YLeaf{"MteTriggerExistenceEvent", mteTriggerExistenceEntry.MteTriggerExistenceEvent})

    mteTriggerExistenceEntry.EntityData.YListKeys = []string {"MteOwner", "MteTriggerName"}

    return &(mteTriggerExistenceEntry.EntityData)
}

// DISMANEVENTMIB_MteTriggerBooleanTable
// A table of management event trigger information for boolean
// triggers.
type DISMANEVENTMIB_MteTriggerBooleanTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single boolean trigger.  Entries automatically exist in
    // this this table for each mteTriggerEntry that has 'boolean' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry.
    MteTriggerBooleanEntry []*DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry
}

func (mteTriggerBooleanTable *DISMANEVENTMIB_MteTriggerBooleanTable) GetEntityData() *types.CommonEntityData {
    mteTriggerBooleanTable.EntityData.YFilter = mteTriggerBooleanTable.YFilter
    mteTriggerBooleanTable.EntityData.YangName = "mteTriggerBooleanTable"
    mteTriggerBooleanTable.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerBooleanTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteTriggerBooleanTable.EntityData.SegmentPath = "mteTriggerBooleanTable"
    mteTriggerBooleanTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteTriggerBooleanTable.EntityData.SegmentPath
    mteTriggerBooleanTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerBooleanTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerBooleanTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerBooleanTable.EntityData.Children = types.NewOrderedMap()
    mteTriggerBooleanTable.EntityData.Children.Append("mteTriggerBooleanEntry", types.YChild{"MteTriggerBooleanEntry", nil})
    for i := range mteTriggerBooleanTable.MteTriggerBooleanEntry {
        mteTriggerBooleanTable.EntityData.Children.Append(types.GetSegmentPath(mteTriggerBooleanTable.MteTriggerBooleanEntry[i]), types.YChild{"MteTriggerBooleanEntry", mteTriggerBooleanTable.MteTriggerBooleanEntry[i]})
    }
    mteTriggerBooleanTable.EntityData.Leafs = types.NewOrderedMap()

    mteTriggerBooleanTable.EntityData.YListKeys = []string {}

    return &(mteTriggerBooleanTable.EntityData)
}

// DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry
// Information about a single boolean trigger.  Entries
// automatically exist in this this table for each mteTriggerEntry
// that has 'boolean' set in mteTriggerTest.
type DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteOwner
    MteOwner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerName
    MteTriggerName interface{}

    // The type of boolean comparison to perform.  The value at mteTriggerValueID
    // is compared to mteTriggerBooleanValue, so for example if
    // mteTriggerBooleanComparison is 'less' the result would be true if the value
    // at mteTriggerValueID is less than the value of mteTriggerBooleanValue. The
    // type is MteTriggerBooleanComparison.
    MteTriggerBooleanComparison interface{}

    // The value to use for the test specified by mteTriggerBooleanTest. The type
    // is interface{} with range: -2147483648..2147483647.
    MteTriggerBooleanValue interface{}

    // Control for whether an event may be triggered when this entry is first set
    // to 'active' or a new instance of the object at mteTriggerValueID is found
    // and the test specified by mteTriggerBooleanComparison is true.  In that
    // case an event is triggered if mteTriggerBooleanStartup is 'true'. The type
    // is bool.
    MteTriggerBooleanStartup interface{}

    // To go with mteTriggerBooleanObjects, the mteOwner of a group of objects
    // from mteObjectsTable. The type is string with length: 0..32.
    MteTriggerBooleanObjectsOwner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable.  These
    // objects are to be added to any Notification resulting from the firing of
    // this trigger for this test.  A list of objects may also be added based on
    // the overall trigger, the event or other settings in mteTriggerTest.  A
    // length of 0 indicates no additional objects. The type is string with
    // length: 0..32.
    MteTriggerBooleanObjects interface{}

    // To go with mteTriggerBooleanEvent, the mteOwner of an event entry from
    // mteEventTable. The type is string with length: 0..32.
    MteTriggerBooleanEventOwner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'boolean'
    // and this trigger fires.  A length of 0 indicates no event. The type is
    // string with length: 0..32.
    MteTriggerBooleanEvent interface{}
}

func (mteTriggerBooleanEntry *DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry) GetEntityData() *types.CommonEntityData {
    mteTriggerBooleanEntry.EntityData.YFilter = mteTriggerBooleanEntry.YFilter
    mteTriggerBooleanEntry.EntityData.YangName = "mteTriggerBooleanEntry"
    mteTriggerBooleanEntry.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerBooleanEntry.EntityData.ParentYangName = "mteTriggerBooleanTable"
    mteTriggerBooleanEntry.EntityData.SegmentPath = "mteTriggerBooleanEntry" + types.AddKeyToken(mteTriggerBooleanEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteTriggerBooleanEntry.MteTriggerName, "mteTriggerName")
    mteTriggerBooleanEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteTriggerBooleanTable/" + mteTriggerBooleanEntry.EntityData.SegmentPath
    mteTriggerBooleanEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerBooleanEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerBooleanEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerBooleanEntry.EntityData.Children = types.NewOrderedMap()
    mteTriggerBooleanEntry.EntityData.Leafs = types.NewOrderedMap()
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteTriggerBooleanEntry.MteOwner})
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteTriggerName", types.YLeaf{"MteTriggerName", mteTriggerBooleanEntry.MteTriggerName})
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteTriggerBooleanComparison", types.YLeaf{"MteTriggerBooleanComparison", mteTriggerBooleanEntry.MteTriggerBooleanComparison})
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteTriggerBooleanValue", types.YLeaf{"MteTriggerBooleanValue", mteTriggerBooleanEntry.MteTriggerBooleanValue})
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteTriggerBooleanStartup", types.YLeaf{"MteTriggerBooleanStartup", mteTriggerBooleanEntry.MteTriggerBooleanStartup})
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteTriggerBooleanObjectsOwner", types.YLeaf{"MteTriggerBooleanObjectsOwner", mteTriggerBooleanEntry.MteTriggerBooleanObjectsOwner})
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteTriggerBooleanObjects", types.YLeaf{"MteTriggerBooleanObjects", mteTriggerBooleanEntry.MteTriggerBooleanObjects})
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteTriggerBooleanEventOwner", types.YLeaf{"MteTriggerBooleanEventOwner", mteTriggerBooleanEntry.MteTriggerBooleanEventOwner})
    mteTriggerBooleanEntry.EntityData.Leafs.Append("mteTriggerBooleanEvent", types.YLeaf{"MteTriggerBooleanEvent", mteTriggerBooleanEntry.MteTriggerBooleanEvent})

    mteTriggerBooleanEntry.EntityData.YListKeys = []string {"MteOwner", "MteTriggerName"}

    return &(mteTriggerBooleanEntry.EntityData)
}

// DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison represents mteTriggerBooleanValue.
type DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison string

const (
    DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison_unequal DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison = "unequal"

    DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison_equal DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison = "equal"

    DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison_less DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison = "less"

    DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison_lessOrEqual DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison = "lessOrEqual"

    DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison_greater DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison = "greater"

    DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison_greaterOrEqual DISMANEVENTMIB_MteTriggerBooleanTable_MteTriggerBooleanEntry_MteTriggerBooleanComparison = "greaterOrEqual"
)

// DISMANEVENTMIB_MteTriggerThresholdTable
// A table of management event trigger information for threshold
// triggers.
type DISMANEVENTMIB_MteTriggerThresholdTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single threshold trigger.  Entries automatically exist
    // in this table for each mteTriggerEntry that has 'threshold' set in
    // mteTriggerTest. The type is slice of
    // DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry.
    MteTriggerThresholdEntry []*DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry
}

func (mteTriggerThresholdTable *DISMANEVENTMIB_MteTriggerThresholdTable) GetEntityData() *types.CommonEntityData {
    mteTriggerThresholdTable.EntityData.YFilter = mteTriggerThresholdTable.YFilter
    mteTriggerThresholdTable.EntityData.YangName = "mteTriggerThresholdTable"
    mteTriggerThresholdTable.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerThresholdTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteTriggerThresholdTable.EntityData.SegmentPath = "mteTriggerThresholdTable"
    mteTriggerThresholdTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteTriggerThresholdTable.EntityData.SegmentPath
    mteTriggerThresholdTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerThresholdTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerThresholdTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerThresholdTable.EntityData.Children = types.NewOrderedMap()
    mteTriggerThresholdTable.EntityData.Children.Append("mteTriggerThresholdEntry", types.YChild{"MteTriggerThresholdEntry", nil})
    for i := range mteTriggerThresholdTable.MteTriggerThresholdEntry {
        mteTriggerThresholdTable.EntityData.Children.Append(types.GetSegmentPath(mteTriggerThresholdTable.MteTriggerThresholdEntry[i]), types.YChild{"MteTriggerThresholdEntry", mteTriggerThresholdTable.MteTriggerThresholdEntry[i]})
    }
    mteTriggerThresholdTable.EntityData.Leafs = types.NewOrderedMap()

    mteTriggerThresholdTable.EntityData.YListKeys = []string {}

    return &(mteTriggerThresholdTable.EntityData)
}

// DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry
// Information about a single threshold trigger.  Entries
// automatically exist in this table for each mteTriggerEntry
// that has 'threshold' set in mteTriggerTest.
type DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteOwner
    MteOwner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteTriggerName
    MteTriggerName interface{}

    // The event that may be triggered when this entry is first set to 'active'
    // and a new instance of the object at mteTriggerValueID is found.  If the
    // first sample after this instance becomes active is greater than or equal to
    // mteTriggerThresholdRising and mteTriggerThresholdStartup is equal to
    // 'rising' or 'risingOrFalling', then one mteTriggerThresholdRisingEvent is
    // triggered for that instance. If the first sample after this entry becomes
    // active is less than or equal to mteTriggerThresholdFalling and
    // mteTriggerThresholdStartup is equal to 'falling' or 'risingOrFalling', then
    // one mteTriggerThresholdRisingEvent is triggered for that instance. The type
    // is MteTriggerThresholdStartup.
    MteTriggerThresholdStartup interface{}

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
    MteTriggerThresholdRising interface{}

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
    MteTriggerThresholdFalling interface{}

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
    MteTriggerThresholdDeltaRising interface{}

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
    MteTriggerThresholdDeltaFalling interface{}

    // To go with mteTriggerThresholdObjects, the mteOwner of a group of objects
    // from mteObjectsTable. The type is string with length: 0..32.
    MteTriggerThresholdObjectsOwner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable.  These
    // objects are to be added to any Notification resulting from the firing of
    // this trigger for this test.  A list of objects may also be added based on
    // the overall  trigger, the event or other settings in mteTriggerTest.  A
    // length of 0 indicates no additional objects. The type is string with
    // length: 0..32.
    MteTriggerThresholdObjects interface{}

    // To go with mteTriggerThresholdRisingEvent, the mteOwner of an event entry
    // from mteEventTable. The type is string with length: 0..32.
    MteTriggerThresholdRisingEventOwner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'threshold'
    // and this trigger fires based on mteTriggerThresholdRising.  A length of 0
    // indicates no event. The type is string with length: 0..32.
    MteTriggerThresholdRisingEvent interface{}

    // To go with mteTriggerThresholdFallingEvent, the mteOwner of an event entry
    // from mteEventTable. The type is string with length: 0..32.
    MteTriggerThresholdFallingEventOwner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'threshold'
    // and this trigger fires based on mteTriggerThresholdFalling.  A length of 0
    // indicates no event. The type is string with length: 0..32.
    MteTriggerThresholdFallingEvent interface{}

    // To go with mteTriggerThresholdDeltaRisingEvent, the mteOwner of an event
    // entry from mteEventTable. The type is string with length: 0..32.
    MteTriggerThresholdDeltaRisingEventOwner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'threshold'
    // and this trigger fires based on mteTriggerThresholdDeltaRising. A length of
    // 0 indicates no event. The type is string with length: 0..32.
    MteTriggerThresholdDeltaRisingEvent interface{}

    // To go with mteTriggerThresholdDeltaFallingEvent, the mteOwner of an event
    // entry from mteEventTable. The type is string with length: 0..32.
    MteTriggerThresholdDeltaFallingEventOwner interface{}

    // The mteEventName of the event to invoke when mteTriggerType is 'threshold'
    // and this trigger fires based on mteTriggerThresholdDeltaFalling.  A length
    // of 0 indicates no event. The type is string with length: 0..32.
    MteTriggerThresholdDeltaFallingEvent interface{}
}

func (mteTriggerThresholdEntry *DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry) GetEntityData() *types.CommonEntityData {
    mteTriggerThresholdEntry.EntityData.YFilter = mteTriggerThresholdEntry.YFilter
    mteTriggerThresholdEntry.EntityData.YangName = "mteTriggerThresholdEntry"
    mteTriggerThresholdEntry.EntityData.BundleName = "cisco_ios_xe"
    mteTriggerThresholdEntry.EntityData.ParentYangName = "mteTriggerThresholdTable"
    mteTriggerThresholdEntry.EntityData.SegmentPath = "mteTriggerThresholdEntry" + types.AddKeyToken(mteTriggerThresholdEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteTriggerThresholdEntry.MteTriggerName, "mteTriggerName")
    mteTriggerThresholdEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteTriggerThresholdTable/" + mteTriggerThresholdEntry.EntityData.SegmentPath
    mteTriggerThresholdEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteTriggerThresholdEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteTriggerThresholdEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteTriggerThresholdEntry.EntityData.Children = types.NewOrderedMap()
    mteTriggerThresholdEntry.EntityData.Leafs = types.NewOrderedMap()
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteTriggerThresholdEntry.MteOwner})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerName", types.YLeaf{"MteTriggerName", mteTriggerThresholdEntry.MteTriggerName})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdStartup", types.YLeaf{"MteTriggerThresholdStartup", mteTriggerThresholdEntry.MteTriggerThresholdStartup})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdRising", types.YLeaf{"MteTriggerThresholdRising", mteTriggerThresholdEntry.MteTriggerThresholdRising})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdFalling", types.YLeaf{"MteTriggerThresholdFalling", mteTriggerThresholdEntry.MteTriggerThresholdFalling})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdDeltaRising", types.YLeaf{"MteTriggerThresholdDeltaRising", mteTriggerThresholdEntry.MteTriggerThresholdDeltaRising})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdDeltaFalling", types.YLeaf{"MteTriggerThresholdDeltaFalling", mteTriggerThresholdEntry.MteTriggerThresholdDeltaFalling})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdObjectsOwner", types.YLeaf{"MteTriggerThresholdObjectsOwner", mteTriggerThresholdEntry.MteTriggerThresholdObjectsOwner})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdObjects", types.YLeaf{"MteTriggerThresholdObjects", mteTriggerThresholdEntry.MteTriggerThresholdObjects})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdRisingEventOwner", types.YLeaf{"MteTriggerThresholdRisingEventOwner", mteTriggerThresholdEntry.MteTriggerThresholdRisingEventOwner})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdRisingEvent", types.YLeaf{"MteTriggerThresholdRisingEvent", mteTriggerThresholdEntry.MteTriggerThresholdRisingEvent})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdFallingEventOwner", types.YLeaf{"MteTriggerThresholdFallingEventOwner", mteTriggerThresholdEntry.MteTriggerThresholdFallingEventOwner})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdFallingEvent", types.YLeaf{"MteTriggerThresholdFallingEvent", mteTriggerThresholdEntry.MteTriggerThresholdFallingEvent})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdDeltaRisingEventOwner", types.YLeaf{"MteTriggerThresholdDeltaRisingEventOwner", mteTriggerThresholdEntry.MteTriggerThresholdDeltaRisingEventOwner})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdDeltaRisingEvent", types.YLeaf{"MteTriggerThresholdDeltaRisingEvent", mteTriggerThresholdEntry.MteTriggerThresholdDeltaRisingEvent})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdDeltaFallingEventOwner", types.YLeaf{"MteTriggerThresholdDeltaFallingEventOwner", mteTriggerThresholdEntry.MteTriggerThresholdDeltaFallingEventOwner})
    mteTriggerThresholdEntry.EntityData.Leafs.Append("mteTriggerThresholdDeltaFallingEvent", types.YLeaf{"MteTriggerThresholdDeltaFallingEvent", mteTriggerThresholdEntry.MteTriggerThresholdDeltaFallingEvent})

    mteTriggerThresholdEntry.EntityData.YListKeys = []string {"MteOwner", "MteTriggerName"}

    return &(mteTriggerThresholdEntry.EntityData)
}

// DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry_MteTriggerThresholdStartup represents triggered for that instance.
type DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry_MteTriggerThresholdStartup string

const (
    DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry_MteTriggerThresholdStartup_rising DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry_MteTriggerThresholdStartup = "rising"

    DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry_MteTriggerThresholdStartup_falling DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry_MteTriggerThresholdStartup = "falling"

    DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry_MteTriggerThresholdStartup_risingOrFalling DISMANEVENTMIB_MteTriggerThresholdTable_MteTriggerThresholdEntry_MteTriggerThresholdStartup = "risingOrFalling"
)

// DISMANEVENTMIB_MteObjectsTable
// A table of objects that can be added to notifications based
// on the trigger, trigger test, or event, as pointed to by
// entries in those tables.
type DISMANEVENTMIB_MteObjectsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A group of objects.  Applications create and delete entries using
    // mteObjectsEntryStatus.  When adding objects to a notification they are
    // added in the lexical order of their index in this table.  Those associated
    // with a trigger come first, then trigger test, then event. The type is slice
    // of DISMANEVENTMIB_MteObjectsTable_MteObjectsEntry.
    MteObjectsEntry []*DISMANEVENTMIB_MteObjectsTable_MteObjectsEntry
}

func (mteObjectsTable *DISMANEVENTMIB_MteObjectsTable) GetEntityData() *types.CommonEntityData {
    mteObjectsTable.EntityData.YFilter = mteObjectsTable.YFilter
    mteObjectsTable.EntityData.YangName = "mteObjectsTable"
    mteObjectsTable.EntityData.BundleName = "cisco_ios_xe"
    mteObjectsTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteObjectsTable.EntityData.SegmentPath = "mteObjectsTable"
    mteObjectsTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteObjectsTable.EntityData.SegmentPath
    mteObjectsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteObjectsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteObjectsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteObjectsTable.EntityData.Children = types.NewOrderedMap()
    mteObjectsTable.EntityData.Children.Append("mteObjectsEntry", types.YChild{"MteObjectsEntry", nil})
    for i := range mteObjectsTable.MteObjectsEntry {
        mteObjectsTable.EntityData.Children.Append(types.GetSegmentPath(mteObjectsTable.MteObjectsEntry[i]), types.YChild{"MteObjectsEntry", mteObjectsTable.MteObjectsEntry[i]})
    }
    mteObjectsTable.EntityData.Leafs = types.NewOrderedMap()

    mteObjectsTable.EntityData.YListKeys = []string {}

    return &(mteObjectsTable.EntityData)
}

// DISMANEVENTMIB_MteObjectsTable_MteObjectsEntry
// A group of objects.  Applications create and delete entries
// using mteObjectsEntryStatus.
// 
// When adding objects to a notification they are added in the
// lexical order of their index in this table.  Those associated
// with a trigger come first, then trigger test, then event.
type DISMANEVENTMIB_MteObjectsTable_MteObjectsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteOwner
    MteOwner interface{}

    // This attribute is a key. A locally-unique, administratively assigned name
    // for a group of objects. The type is string with length: 1..32.
    MteObjectsName interface{}

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
    MteObjectsIndex interface{}

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
    MteObjectsID interface{}

    // Control for whether mteObjectsID is to be treated as fully-specified or
    // wildcarded, with 'true' indicating wildcard. The type is bool.
    MteObjectsIDWildcard interface{}

    // The control that allows creation and deletion of entries. Once made active
    // an entry MAY not be modified except to delete it. The type is RowStatus.
    MteObjectsEntryStatus interface{}
}

func (mteObjectsEntry *DISMANEVENTMIB_MteObjectsTable_MteObjectsEntry) GetEntityData() *types.CommonEntityData {
    mteObjectsEntry.EntityData.YFilter = mteObjectsEntry.YFilter
    mteObjectsEntry.EntityData.YangName = "mteObjectsEntry"
    mteObjectsEntry.EntityData.BundleName = "cisco_ios_xe"
    mteObjectsEntry.EntityData.ParentYangName = "mteObjectsTable"
    mteObjectsEntry.EntityData.SegmentPath = "mteObjectsEntry" + types.AddKeyToken(mteObjectsEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteObjectsEntry.MteObjectsName, "mteObjectsName") + types.AddKeyToken(mteObjectsEntry.MteObjectsIndex, "mteObjectsIndex")
    mteObjectsEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteObjectsTable/" + mteObjectsEntry.EntityData.SegmentPath
    mteObjectsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteObjectsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteObjectsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteObjectsEntry.EntityData.Children = types.NewOrderedMap()
    mteObjectsEntry.EntityData.Leafs = types.NewOrderedMap()
    mteObjectsEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteObjectsEntry.MteOwner})
    mteObjectsEntry.EntityData.Leafs.Append("mteObjectsName", types.YLeaf{"MteObjectsName", mteObjectsEntry.MteObjectsName})
    mteObjectsEntry.EntityData.Leafs.Append("mteObjectsIndex", types.YLeaf{"MteObjectsIndex", mteObjectsEntry.MteObjectsIndex})
    mteObjectsEntry.EntityData.Leafs.Append("mteObjectsID", types.YLeaf{"MteObjectsID", mteObjectsEntry.MteObjectsID})
    mteObjectsEntry.EntityData.Leafs.Append("mteObjectsIDWildcard", types.YLeaf{"MteObjectsIDWildcard", mteObjectsEntry.MteObjectsIDWildcard})
    mteObjectsEntry.EntityData.Leafs.Append("mteObjectsEntryStatus", types.YLeaf{"MteObjectsEntryStatus", mteObjectsEntry.MteObjectsEntryStatus})

    mteObjectsEntry.EntityData.YListKeys = []string {"MteOwner", "MteObjectsName", "MteObjectsIndex"}

    return &(mteObjectsEntry.EntityData)
}

// DISMANEVENTMIB_MteEventTable
// A table of management event action information.
type DISMANEVENTMIB_MteEventTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single event.  Applications create and delete entries
    // using mteEventEntryStatus. The type is slice of
    // DISMANEVENTMIB_MteEventTable_MteEventEntry.
    MteEventEntry []*DISMANEVENTMIB_MteEventTable_MteEventEntry
}

func (mteEventTable *DISMANEVENTMIB_MteEventTable) GetEntityData() *types.CommonEntityData {
    mteEventTable.EntityData.YFilter = mteEventTable.YFilter
    mteEventTable.EntityData.YangName = "mteEventTable"
    mteEventTable.EntityData.BundleName = "cisco_ios_xe"
    mteEventTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteEventTable.EntityData.SegmentPath = "mteEventTable"
    mteEventTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteEventTable.EntityData.SegmentPath
    mteEventTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteEventTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteEventTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteEventTable.EntityData.Children = types.NewOrderedMap()
    mteEventTable.EntityData.Children.Append("mteEventEntry", types.YChild{"MteEventEntry", nil})
    for i := range mteEventTable.MteEventEntry {
        mteEventTable.EntityData.Children.Append(types.GetSegmentPath(mteEventTable.MteEventEntry[i]), types.YChild{"MteEventEntry", mteEventTable.MteEventEntry[i]})
    }
    mteEventTable.EntityData.Leafs = types.NewOrderedMap()

    mteEventTable.EntityData.YListKeys = []string {}

    return &(mteEventTable.EntityData)
}

// DISMANEVENTMIB_MteEventTable_MteEventEntry
// Information about a single event.  Applications create and
// delete entries using mteEventEntryStatus.
type DISMANEVENTMIB_MteEventTable_MteEventEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteOwner
    MteOwner interface{}

    // This attribute is a key. A locally-unique, administratively assigned name
    // for the event. The type is string with length: 1..32.
    MteEventName interface{}

    // A description of the event's function and use. The type is string.
    MteEventComment interface{}

    // The actions to perform when this event occurs.  For 'notification', Traps
    // and/or Informs are sent according to the configuration in the SNMP
    // Notification MIB.  For 'set', an SNMP Set operation is performed according
    // to control values in this entry. The type is map[string]bool.
    MteEventActions interface{}

    // A control to allow an event to be configured but not used. When the value
    // is 'false' the event does not execute even if  triggered. The type is bool.
    MteEventEnabled interface{}

    // The control that allows creation and deletion of entries. Once made active
    // an entry MAY not be modified except to delete it. The type is RowStatus.
    MteEventEntryStatus interface{}
}

func (mteEventEntry *DISMANEVENTMIB_MteEventTable_MteEventEntry) GetEntityData() *types.CommonEntityData {
    mteEventEntry.EntityData.YFilter = mteEventEntry.YFilter
    mteEventEntry.EntityData.YangName = "mteEventEntry"
    mteEventEntry.EntityData.BundleName = "cisco_ios_xe"
    mteEventEntry.EntityData.ParentYangName = "mteEventTable"
    mteEventEntry.EntityData.SegmentPath = "mteEventEntry" + types.AddKeyToken(mteEventEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteEventEntry.MteEventName, "mteEventName")
    mteEventEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteEventTable/" + mteEventEntry.EntityData.SegmentPath
    mteEventEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteEventEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteEventEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteEventEntry.EntityData.Children = types.NewOrderedMap()
    mteEventEntry.EntityData.Leafs = types.NewOrderedMap()
    mteEventEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteEventEntry.MteOwner})
    mteEventEntry.EntityData.Leafs.Append("mteEventName", types.YLeaf{"MteEventName", mteEventEntry.MteEventName})
    mteEventEntry.EntityData.Leafs.Append("mteEventComment", types.YLeaf{"MteEventComment", mteEventEntry.MteEventComment})
    mteEventEntry.EntityData.Leafs.Append("mteEventActions", types.YLeaf{"MteEventActions", mteEventEntry.MteEventActions})
    mteEventEntry.EntityData.Leafs.Append("mteEventEnabled", types.YLeaf{"MteEventEnabled", mteEventEntry.MteEventEnabled})
    mteEventEntry.EntityData.Leafs.Append("mteEventEntryStatus", types.YLeaf{"MteEventEntryStatus", mteEventEntry.MteEventEntryStatus})

    mteEventEntry.EntityData.YListKeys = []string {"MteOwner", "MteEventName"}

    return &(mteEventEntry.EntityData)
}

// DISMANEVENTMIB_MteEventNotificationTable
// A table of information about notifications to be sent as a
// consequence of management events.
type DISMANEVENTMIB_MteEventNotificationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single event's notification.  Entries automatically
    // exist in this this table for each mteEventEntry that has 'notification' set
    // in mteEventActions. The type is slice of
    // DISMANEVENTMIB_MteEventNotificationTable_MteEventNotificationEntry.
    MteEventNotificationEntry []*DISMANEVENTMIB_MteEventNotificationTable_MteEventNotificationEntry
}

func (mteEventNotificationTable *DISMANEVENTMIB_MteEventNotificationTable) GetEntityData() *types.CommonEntityData {
    mteEventNotificationTable.EntityData.YFilter = mteEventNotificationTable.YFilter
    mteEventNotificationTable.EntityData.YangName = "mteEventNotificationTable"
    mteEventNotificationTable.EntityData.BundleName = "cisco_ios_xe"
    mteEventNotificationTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteEventNotificationTable.EntityData.SegmentPath = "mteEventNotificationTable"
    mteEventNotificationTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteEventNotificationTable.EntityData.SegmentPath
    mteEventNotificationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteEventNotificationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteEventNotificationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteEventNotificationTable.EntityData.Children = types.NewOrderedMap()
    mteEventNotificationTable.EntityData.Children.Append("mteEventNotificationEntry", types.YChild{"MteEventNotificationEntry", nil})
    for i := range mteEventNotificationTable.MteEventNotificationEntry {
        mteEventNotificationTable.EntityData.Children.Append(types.GetSegmentPath(mteEventNotificationTable.MteEventNotificationEntry[i]), types.YChild{"MteEventNotificationEntry", mteEventNotificationTable.MteEventNotificationEntry[i]})
    }
    mteEventNotificationTable.EntityData.Leafs = types.NewOrderedMap()

    mteEventNotificationTable.EntityData.YListKeys = []string {}

    return &(mteEventNotificationTable.EntityData)
}

// DISMANEVENTMIB_MteEventNotificationTable_MteEventNotificationEntry
// Information about a single event's notification.  Entries
// automatically exist in this this table for each mteEventEntry
// that has 'notification' set in mteEventActions.
type DISMANEVENTMIB_MteEventNotificationTable_MteEventNotificationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteOwner
    MteOwner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteEventTable_MteEventEntry_MteEventName
    MteEventName interface{}

    // The object identifier from the NOTIFICATION-TYPE for the notification to
    // use if metEventActions has 'notification' set. The type is string with
    // pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    MteEventNotification interface{}

    // To go with mteEventNotificationObjects, the mteOwner of a group of objects
    // from mteObjectsTable. The type is string with length: 0..32.
    MteEventNotificationObjectsOwner interface{}

    // The mteObjectsName of a group of objects from mteObjectsTable if
    // mteEventActions has 'notification' set. These objects are to be added to
    // any Notification generated by this event.  Objects may also be added based
    // on the trigger that stimulated the event.  A length of 0 indicates no
    // additional objects. The type is string with length: 0..32.
    MteEventNotificationObjects interface{}
}

func (mteEventNotificationEntry *DISMANEVENTMIB_MteEventNotificationTable_MteEventNotificationEntry) GetEntityData() *types.CommonEntityData {
    mteEventNotificationEntry.EntityData.YFilter = mteEventNotificationEntry.YFilter
    mteEventNotificationEntry.EntityData.YangName = "mteEventNotificationEntry"
    mteEventNotificationEntry.EntityData.BundleName = "cisco_ios_xe"
    mteEventNotificationEntry.EntityData.ParentYangName = "mteEventNotificationTable"
    mteEventNotificationEntry.EntityData.SegmentPath = "mteEventNotificationEntry" + types.AddKeyToken(mteEventNotificationEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteEventNotificationEntry.MteEventName, "mteEventName")
    mteEventNotificationEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteEventNotificationTable/" + mteEventNotificationEntry.EntityData.SegmentPath
    mteEventNotificationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteEventNotificationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteEventNotificationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteEventNotificationEntry.EntityData.Children = types.NewOrderedMap()
    mteEventNotificationEntry.EntityData.Leafs = types.NewOrderedMap()
    mteEventNotificationEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteEventNotificationEntry.MteOwner})
    mteEventNotificationEntry.EntityData.Leafs.Append("mteEventName", types.YLeaf{"MteEventName", mteEventNotificationEntry.MteEventName})
    mteEventNotificationEntry.EntityData.Leafs.Append("mteEventNotification", types.YLeaf{"MteEventNotification", mteEventNotificationEntry.MteEventNotification})
    mteEventNotificationEntry.EntityData.Leafs.Append("mteEventNotificationObjectsOwner", types.YLeaf{"MteEventNotificationObjectsOwner", mteEventNotificationEntry.MteEventNotificationObjectsOwner})
    mteEventNotificationEntry.EntityData.Leafs.Append("mteEventNotificationObjects", types.YLeaf{"MteEventNotificationObjects", mteEventNotificationEntry.MteEventNotificationObjects})

    mteEventNotificationEntry.EntityData.YListKeys = []string {"MteOwner", "MteEventName"}

    return &(mteEventNotificationEntry.EntityData)
}

// DISMANEVENTMIB_MteEventSetTable
// A table of management event action information.
type DISMANEVENTMIB_MteEventSetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single event's set option.  Entries automatically exist
    // in this this table for each mteEventEntry that has 'set' set in
    // mteEventActions. The type is slice of
    // DISMANEVENTMIB_MteEventSetTable_MteEventSetEntry.
    MteEventSetEntry []*DISMANEVENTMIB_MteEventSetTable_MteEventSetEntry
}

func (mteEventSetTable *DISMANEVENTMIB_MteEventSetTable) GetEntityData() *types.CommonEntityData {
    mteEventSetTable.EntityData.YFilter = mteEventSetTable.YFilter
    mteEventSetTable.EntityData.YangName = "mteEventSetTable"
    mteEventSetTable.EntityData.BundleName = "cisco_ios_xe"
    mteEventSetTable.EntityData.ParentYangName = "DISMAN-EVENT-MIB"
    mteEventSetTable.EntityData.SegmentPath = "mteEventSetTable"
    mteEventSetTable.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/" + mteEventSetTable.EntityData.SegmentPath
    mteEventSetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteEventSetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteEventSetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteEventSetTable.EntityData.Children = types.NewOrderedMap()
    mteEventSetTable.EntityData.Children.Append("mteEventSetEntry", types.YChild{"MteEventSetEntry", nil})
    for i := range mteEventSetTable.MteEventSetEntry {
        mteEventSetTable.EntityData.Children.Append(types.GetSegmentPath(mteEventSetTable.MteEventSetEntry[i]), types.YChild{"MteEventSetEntry", mteEventSetTable.MteEventSetEntry[i]})
    }
    mteEventSetTable.EntityData.Leafs = types.NewOrderedMap()

    mteEventSetTable.EntityData.YListKeys = []string {}

    return &(mteEventSetTable.EntityData)
}

// DISMANEVENTMIB_MteEventSetTable_MteEventSetEntry
// Information about a single event's set option.  Entries
// automatically exist in this this table for each mteEventEntry
// that has 'set' set in mteEventActions.
type DISMANEVENTMIB_MteEventSetTable_MteEventSetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 0..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteTriggerTable_MteTriggerEntry_MteOwner
    MteOwner interface{}

    // This attribute is a key. The type is string with length: 1..32. Refers to
    // disman_event_mib.DISMANEVENTMIB_MteEventTable_MteEventEntry_MteEventName
    MteEventName interface{}

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
    MteEventSetObject interface{}

    // Control over whether mteEventSetObject is to be treated as fully-specified
    // or wildcarded, with 'true' indicating wildcard if mteEventActions has 'set'
    // set. The type is bool.
    MteEventSetObjectWildcard interface{}

    // The value to which to set the object at mteEventSetObject if
    // mteEventActions has 'set' set. The type is interface{} with range:
    // -2147483648..2147483647.
    MteEventSetValue interface{}

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
    MteEventSetTargetTag interface{}

    // The management context in which to set mteEventObjectID. if mteEventActions
    // has 'set' set.  This may be wildcarded by leaving characters off the end. 
    // To indicate such wildcarding mteEventSetContextNameWildcard must be 'true'.
    // If this context name is wildcarded the value used to complete the
    // wildcarding of mteTriggerContextName will be appended. The type is string.
    MteEventSetContextName interface{}

    // Control for whether mteEventSetContextName is to be treated as
    // fully-specified or wildcarded, with 'true' indicating wildcard if
    // mteEventActions has 'set' set. The type is bool.
    MteEventSetContextNameWildcard interface{}
}

func (mteEventSetEntry *DISMANEVENTMIB_MteEventSetTable_MteEventSetEntry) GetEntityData() *types.CommonEntityData {
    mteEventSetEntry.EntityData.YFilter = mteEventSetEntry.YFilter
    mteEventSetEntry.EntityData.YangName = "mteEventSetEntry"
    mteEventSetEntry.EntityData.BundleName = "cisco_ios_xe"
    mteEventSetEntry.EntityData.ParentYangName = "mteEventSetTable"
    mteEventSetEntry.EntityData.SegmentPath = "mteEventSetEntry" + types.AddKeyToken(mteEventSetEntry.MteOwner, "mteOwner") + types.AddKeyToken(mteEventSetEntry.MteEventName, "mteEventName")
    mteEventSetEntry.EntityData.AbsolutePath = "DISMAN-EVENT-MIB:DISMAN-EVENT-MIB/mteEventSetTable/" + mteEventSetEntry.EntityData.SegmentPath
    mteEventSetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mteEventSetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mteEventSetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mteEventSetEntry.EntityData.Children = types.NewOrderedMap()
    mteEventSetEntry.EntityData.Leafs = types.NewOrderedMap()
    mteEventSetEntry.EntityData.Leafs.Append("mteOwner", types.YLeaf{"MteOwner", mteEventSetEntry.MteOwner})
    mteEventSetEntry.EntityData.Leafs.Append("mteEventName", types.YLeaf{"MteEventName", mteEventSetEntry.MteEventName})
    mteEventSetEntry.EntityData.Leafs.Append("mteEventSetObject", types.YLeaf{"MteEventSetObject", mteEventSetEntry.MteEventSetObject})
    mteEventSetEntry.EntityData.Leafs.Append("mteEventSetObjectWildcard", types.YLeaf{"MteEventSetObjectWildcard", mteEventSetEntry.MteEventSetObjectWildcard})
    mteEventSetEntry.EntityData.Leafs.Append("mteEventSetValue", types.YLeaf{"MteEventSetValue", mteEventSetEntry.MteEventSetValue})
    mteEventSetEntry.EntityData.Leafs.Append("mteEventSetTargetTag", types.YLeaf{"MteEventSetTargetTag", mteEventSetEntry.MteEventSetTargetTag})
    mteEventSetEntry.EntityData.Leafs.Append("mteEventSetContextName", types.YLeaf{"MteEventSetContextName", mteEventSetEntry.MteEventSetContextName})
    mteEventSetEntry.EntityData.Leafs.Append("mteEventSetContextNameWildcard", types.YLeaf{"MteEventSetContextNameWildcard", mteEventSetEntry.MteEventSetContextNameWildcard})

    mteEventSetEntry.EntityData.YListKeys = []string {"MteOwner", "MteEventName"}

    return &(mteEventSetEntry.EntityData)
}

