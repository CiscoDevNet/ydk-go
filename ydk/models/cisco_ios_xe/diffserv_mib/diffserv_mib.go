// This MIB defines the objects necessary to manage a device that
// uses the Differentiated Services Architecture described in RFC
// 2475. The Conceptual Model of a Differentiated Services Router
// provides supporting information on how such a router is modeled.
package diffserv_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package diffserv_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DIFFSERV-MIB DIFFSERV-MIB}", reflect.TypeOf(DIFFSERVMIB{}))
    ydk.RegisterEntity("DIFFSERV-MIB:DIFFSERV-MIB", reflect.TypeOf(DIFFSERVMIB{}))
}

type DiffServTBParamSimpleTokenBucket struct {
}

func (id DiffServTBParamSimpleTokenBucket) String() string {
	return "DIFFSERV-MIB:diffServTBParamSimpleTokenBucket"
}

type DiffServTBParamAvgRate struct {
}

func (id DiffServTBParamAvgRate) String() string {
	return "DIFFSERV-MIB:diffServTBParamAvgRate"
}

type DiffServTBParamSrTCMBlind struct {
}

func (id DiffServTBParamSrTCMBlind) String() string {
	return "DIFFSERV-MIB:diffServTBParamSrTCMBlind"
}

type DiffServTBParamSrTCMAware struct {
}

func (id DiffServTBParamSrTCMAware) String() string {
	return "DIFFSERV-MIB:diffServTBParamSrTCMAware"
}

type DiffServTBParamTrTCMBlind struct {
}

func (id DiffServTBParamTrTCMBlind) String() string {
	return "DIFFSERV-MIB:diffServTBParamTrTCMBlind"
}

type DiffServTBParamTrTCMAware struct {
}

func (id DiffServTBParamTrTCMAware) String() string {
	return "DIFFSERV-MIB:diffServTBParamTrTCMAware"
}

type DiffServTBParamTswTCM struct {
}

func (id DiffServTBParamTswTCM) String() string {
	return "DIFFSERV-MIB:diffServTBParamTswTCM"
}

type DiffServSchedulerPriority struct {
}

func (id DiffServSchedulerPriority) String() string {
	return "DIFFSERV-MIB:diffServSchedulerPriority"
}

type DiffServSchedulerWRR struct {
}

func (id DiffServSchedulerWRR) String() string {
	return "DIFFSERV-MIB:diffServSchedulerWRR"
}

type DiffServSchedulerWFQ struct {
}

func (id DiffServSchedulerWFQ) String() string {
	return "DIFFSERV-MIB:diffServSchedulerWFQ"
}

// IfDirection represents transmission on the interface.
type IfDirection string

const (
    IfDirection_inbound IfDirection = "inbound"

    IfDirection_outbound IfDirection = "outbound"
)

// DIFFSERVMIB
type DIFFSERVMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    DiffServClassifier DIFFSERVMIB_DiffServClassifier

    
    DiffServMeter DIFFSERVMIB_DiffServMeter

    
    DiffServTBParam DIFFSERVMIB_DiffServTBParam

    
    DiffServAction DIFFSERVMIB_DiffServAction

    
    DiffServAlgDrop DIFFSERVMIB_DiffServAlgDrop

    
    DiffServQueue DIFFSERVMIB_DiffServQueue

    
    DiffServScheduler DIFFSERVMIB_DiffServScheduler

    // The data path table contains RowPointers indicating the start of the
    // functional data path for each interface and traffic direction in this
    // device. These may merge, or be separated into parallel data paths.
    DiffServDataPathTable DIFFSERVMIB_DiffServDataPathTable

    // This table enumerates all the diffserv classifier functional data path
    // elements of this device.  The actual classification definitions are defined
    // in diffServClfrElementTable entries belonging to each classifier.  An entry
    // in this table, pointed to by a RowPointer specifying an instance of
    // diffServClfrStatus, is frequently used as the name for a set of classifier
    // elements, which all use the index diffServClfrId. Per the semantics of the
    // classifier element table, these entries constitute one or more unordered
    // sets of tests which may be simultaneously applied to a message to   
    // classify it.  The primary function of this table is to ensure that the
    // value of diffServClfrId is unique before attempting to use it in creating a
    // diffServClfrElementEntry. Therefore, the diffServClfrEntry must be created
    // on the same SET as the diffServClfrElementEntry, or before the
    // diffServClfrElementEntry is created.
    DiffServClfrTable DIFFSERVMIB_DiffServClfrTable

    // The classifier element table enumerates the relationship between
    // classification patterns and subsequent downstream Differentiated Services
    // Functional Data Path elements. diffServClfrElementSpecific points to a
    // filter that specifies the classification parameters. A classifier may use
    // filter tables of different types together.  One example of a filter table
    // defined in this MIB is diffServMultiFieldClfrTable, for IP Multi-Field
    // Classifiers (MFCs). Such an entry might identify anything from a single
    // micro-flow (an identifiable sub-session packet stream directed from one
    // sending transport to the receiving transport or transports), or aggregates
    // of those such as the traffic from a host, traffic for an application, or
    // traffic between two hosts using an application and a given DSCP. The
    // standard Behavior Aggregate used in the Differentiated Services
    // Architecture is encoded as a degenerate case of such an aggregate - the
    // traffic using a particular DSCP value.  Filter tables for other filter
    // types may be defined elsewhere.
    DiffServClfrElementTable DIFFSERVMIB_DiffServClfrElementTable

    // A table of IP Multi-field Classifier filter entries that a    system may
    // use to identify IP traffic.
    DiffServMultiFieldClfrTable DIFFSERVMIB_DiffServMultiFieldClfrTable

    // This table enumerates specific meters that a system may use to police a
    // stream of traffic. The traffic stream to be metered is determined by the
    // Differentiated Services Functional Data Path Element(s) upstream of the
    // meter i.e. by the object(s) that point to each entry in this table.  This
    // may include all traffic on an interface.  Specific meter details are to be
    // found in table entry referenced by diffServMeterSpecific.
    DiffServMeterTable DIFFSERVMIB_DiffServMeterTable

    // This table enumerates a single set of token bucket meter parameters that a
    // system may use to police a stream of traffic. Such meters are modeled here
    // as having a single rate and a single burst size. Multiple entries are used
    // when multiple rates/burst sizes are needed.
    DiffServTBParamTable DIFFSERVMIB_DiffServTBParamTable

    // The Action Table enumerates actions that can be performed to a stream of
    // traffic. Multiple actions can be concatenated. For example, traffic exiting
    // from a meter may be counted, marked, and potentially dropped before
    // entering a queue.  Specific actions are indicated by diffServActionSpecific
    // which points to an entry of a specific action type parameterizing the
    // action in detail.
    DiffServActionTable DIFFSERVMIB_DiffServActionTable

    // This table enumerates specific DSCPs used for marking or remarking the DSCP
    // field of IP packets. The entries of this table may be referenced by a
    // diffServActionSpecific attribute.
    DiffServDscpMarkActTable DIFFSERVMIB_DiffServDscpMarkActTable

    // This table contains counters for all the traffic passing through an action
    // element.
    DiffServCountActTable DIFFSERVMIB_DiffServCountActTable

    // The algorithmic drop table contains entries describing an element that
    // drops packets according to some algorithm.
    DiffServAlgDropTable DIFFSERVMIB_DiffServAlgDropTable

    // The random drop table contains entries describing a process that drops
    // packets randomly. Entries in this table are pointed to by
    // diffServAlgDropSpecific.
    DiffServRandomDropTable DIFFSERVMIB_DiffServRandomDropTable

    // The Queue Table enumerates the individual queues.  Note that the MIB models
    // queuing systems as composed of individual queues, one per class of traffic,
    // even though they may in fact be structured as classes of traffic scheduled
    // using a common calendar queue, or in other ways.
    DiffServQTable DIFFSERVMIB_DiffServQTable

    // The Scheduler Table enumerates packet schedulers. Multiple scheduling
    // algorithms can be used on a given data path, with each algorithm described
    // by one diffServSchedulerEntry.
    DiffServSchedulerTable DIFFSERVMIB_DiffServSchedulerTable

    // The Minimum Rate Parameters Table enumerates individual sets of scheduling
    // parameter that can be used/reused by Queues and Schedulers.
    DiffServMinRateTable DIFFSERVMIB_DiffServMinRateTable

    // The Maximum Rate Parameter Table enumerates individual sets of scheduling
    // parameter that can be used/reused by Queues and Schedulers.
    DiffServMaxRateTable DIFFSERVMIB_DiffServMaxRateTable
}

func (dIFFSERVMIB *DIFFSERVMIB) GetEntityData() *types.CommonEntityData {
    dIFFSERVMIB.EntityData.YFilter = dIFFSERVMIB.YFilter
    dIFFSERVMIB.EntityData.YangName = "DIFFSERV-MIB"
    dIFFSERVMIB.EntityData.BundleName = "cisco_ios_xe"
    dIFFSERVMIB.EntityData.ParentYangName = "DIFFSERV-MIB"
    dIFFSERVMIB.EntityData.SegmentPath = "DIFFSERV-MIB:DIFFSERV-MIB"
    dIFFSERVMIB.EntityData.AbsolutePath = dIFFSERVMIB.EntityData.SegmentPath
    dIFFSERVMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dIFFSERVMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dIFFSERVMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dIFFSERVMIB.EntityData.Children = types.NewOrderedMap()
    dIFFSERVMIB.EntityData.Children.Append("diffServClassifier", types.YChild{"DiffServClassifier", &dIFFSERVMIB.DiffServClassifier})
    dIFFSERVMIB.EntityData.Children.Append("diffServMeter", types.YChild{"DiffServMeter", &dIFFSERVMIB.DiffServMeter})
    dIFFSERVMIB.EntityData.Children.Append("diffServTBParam", types.YChild{"DiffServTBParam", &dIFFSERVMIB.DiffServTBParam})
    dIFFSERVMIB.EntityData.Children.Append("diffServAction", types.YChild{"DiffServAction", &dIFFSERVMIB.DiffServAction})
    dIFFSERVMIB.EntityData.Children.Append("diffServAlgDrop", types.YChild{"DiffServAlgDrop", &dIFFSERVMIB.DiffServAlgDrop})
    dIFFSERVMIB.EntityData.Children.Append("diffServQueue", types.YChild{"DiffServQueue", &dIFFSERVMIB.DiffServQueue})
    dIFFSERVMIB.EntityData.Children.Append("diffServScheduler", types.YChild{"DiffServScheduler", &dIFFSERVMIB.DiffServScheduler})
    dIFFSERVMIB.EntityData.Children.Append("diffServDataPathTable", types.YChild{"DiffServDataPathTable", &dIFFSERVMIB.DiffServDataPathTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServClfrTable", types.YChild{"DiffServClfrTable", &dIFFSERVMIB.DiffServClfrTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServClfrElementTable", types.YChild{"DiffServClfrElementTable", &dIFFSERVMIB.DiffServClfrElementTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServMultiFieldClfrTable", types.YChild{"DiffServMultiFieldClfrTable", &dIFFSERVMIB.DiffServMultiFieldClfrTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServMeterTable", types.YChild{"DiffServMeterTable", &dIFFSERVMIB.DiffServMeterTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServTBParamTable", types.YChild{"DiffServTBParamTable", &dIFFSERVMIB.DiffServTBParamTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServActionTable", types.YChild{"DiffServActionTable", &dIFFSERVMIB.DiffServActionTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServDscpMarkActTable", types.YChild{"DiffServDscpMarkActTable", &dIFFSERVMIB.DiffServDscpMarkActTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServCountActTable", types.YChild{"DiffServCountActTable", &dIFFSERVMIB.DiffServCountActTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServAlgDropTable", types.YChild{"DiffServAlgDropTable", &dIFFSERVMIB.DiffServAlgDropTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServRandomDropTable", types.YChild{"DiffServRandomDropTable", &dIFFSERVMIB.DiffServRandomDropTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServQTable", types.YChild{"DiffServQTable", &dIFFSERVMIB.DiffServQTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServSchedulerTable", types.YChild{"DiffServSchedulerTable", &dIFFSERVMIB.DiffServSchedulerTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServMinRateTable", types.YChild{"DiffServMinRateTable", &dIFFSERVMIB.DiffServMinRateTable})
    dIFFSERVMIB.EntityData.Children.Append("diffServMaxRateTable", types.YChild{"DiffServMaxRateTable", &dIFFSERVMIB.DiffServMaxRateTable})
    dIFFSERVMIB.EntityData.Leafs = types.NewOrderedMap()

    dIFFSERVMIB.EntityData.YListKeys = []string {}

    return &(dIFFSERVMIB.EntityData)
}

// DIFFSERVMIB_DiffServClassifier
type DIFFSERVMIB_DiffServClassifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServClfrId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServClfrNextFree interface{}

    // This object contains an unused value for diffServClfrElementId, or a zero
    // to indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServClfrElementNextFree interface{}

    // This object contains an unused value for diffServMultiFieldClfrId, or a
    // zero to indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServMultiFieldClfrNextFree interface{}
}

func (diffServClassifier *DIFFSERVMIB_DiffServClassifier) GetEntityData() *types.CommonEntityData {
    diffServClassifier.EntityData.YFilter = diffServClassifier.YFilter
    diffServClassifier.EntityData.YangName = "diffServClassifier"
    diffServClassifier.EntityData.BundleName = "cisco_ios_xe"
    diffServClassifier.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServClassifier.EntityData.SegmentPath = "diffServClassifier"
    diffServClassifier.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServClassifier.EntityData.SegmentPath
    diffServClassifier.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServClassifier.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServClassifier.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServClassifier.EntityData.Children = types.NewOrderedMap()
    diffServClassifier.EntityData.Leafs = types.NewOrderedMap()
    diffServClassifier.EntityData.Leafs.Append("diffServClfrNextFree", types.YLeaf{"DiffServClfrNextFree", diffServClassifier.DiffServClfrNextFree})
    diffServClassifier.EntityData.Leafs.Append("diffServClfrElementNextFree", types.YLeaf{"DiffServClfrElementNextFree", diffServClassifier.DiffServClfrElementNextFree})
    diffServClassifier.EntityData.Leafs.Append("diffServMultiFieldClfrNextFree", types.YLeaf{"DiffServMultiFieldClfrNextFree", diffServClassifier.DiffServMultiFieldClfrNextFree})

    diffServClassifier.EntityData.YListKeys = []string {}

    return &(diffServClassifier.EntityData)
}

// DIFFSERVMIB_DiffServMeter
type DIFFSERVMIB_DiffServMeter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServMeterId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServMeterNextFree interface{}
}

func (diffServMeter *DIFFSERVMIB_DiffServMeter) GetEntityData() *types.CommonEntityData {
    diffServMeter.EntityData.YFilter = diffServMeter.YFilter
    diffServMeter.EntityData.YangName = "diffServMeter"
    diffServMeter.EntityData.BundleName = "cisco_ios_xe"
    diffServMeter.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServMeter.EntityData.SegmentPath = "diffServMeter"
    diffServMeter.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServMeter.EntityData.SegmentPath
    diffServMeter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMeter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMeter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMeter.EntityData.Children = types.NewOrderedMap()
    diffServMeter.EntityData.Leafs = types.NewOrderedMap()
    diffServMeter.EntityData.Leafs.Append("diffServMeterNextFree", types.YLeaf{"DiffServMeterNextFree", diffServMeter.DiffServMeterNextFree})

    diffServMeter.EntityData.YListKeys = []string {}

    return &(diffServMeter.EntityData)
}

// DIFFSERVMIB_DiffServTBParam
type DIFFSERVMIB_DiffServTBParam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServTBParamId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServTBParamNextFree interface{}
}

func (diffServTBParam *DIFFSERVMIB_DiffServTBParam) GetEntityData() *types.CommonEntityData {
    diffServTBParam.EntityData.YFilter = diffServTBParam.YFilter
    diffServTBParam.EntityData.YangName = "diffServTBParam"
    diffServTBParam.EntityData.BundleName = "cisco_ios_xe"
    diffServTBParam.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServTBParam.EntityData.SegmentPath = "diffServTBParam"
    diffServTBParam.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServTBParam.EntityData.SegmentPath
    diffServTBParam.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServTBParam.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServTBParam.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServTBParam.EntityData.Children = types.NewOrderedMap()
    diffServTBParam.EntityData.Leafs = types.NewOrderedMap()
    diffServTBParam.EntityData.Leafs.Append("diffServTBParamNextFree", types.YLeaf{"DiffServTBParamNextFree", diffServTBParam.DiffServTBParamNextFree})

    diffServTBParam.EntityData.YListKeys = []string {}

    return &(diffServTBParam.EntityData)
}

// DIFFSERVMIB_DiffServAction
type DIFFSERVMIB_DiffServAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServActionId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServActionNextFree interface{}

    // This object contains an unused value for diffServCountActId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServCountActNextFree interface{}
}

func (diffServAction *DIFFSERVMIB_DiffServAction) GetEntityData() *types.CommonEntityData {
    diffServAction.EntityData.YFilter = diffServAction.YFilter
    diffServAction.EntityData.YangName = "diffServAction"
    diffServAction.EntityData.BundleName = "cisco_ios_xe"
    diffServAction.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServAction.EntityData.SegmentPath = "diffServAction"
    diffServAction.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServAction.EntityData.SegmentPath
    diffServAction.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServAction.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServAction.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServAction.EntityData.Children = types.NewOrderedMap()
    diffServAction.EntityData.Leafs = types.NewOrderedMap()
    diffServAction.EntityData.Leafs.Append("diffServActionNextFree", types.YLeaf{"DiffServActionNextFree", diffServAction.DiffServActionNextFree})
    diffServAction.EntityData.Leafs.Append("diffServCountActNextFree", types.YLeaf{"DiffServCountActNextFree", diffServAction.DiffServCountActNextFree})

    diffServAction.EntityData.YListKeys = []string {}

    return &(diffServAction.EntityData)
}

// DIFFSERVMIB_DiffServAlgDrop
type DIFFSERVMIB_DiffServAlgDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServAlgDropId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServAlgDropNextFree interface{}

    // This object contains an unused value for diffServRandomDropId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServRandomDropNextFree interface{}
}

func (diffServAlgDrop *DIFFSERVMIB_DiffServAlgDrop) GetEntityData() *types.CommonEntityData {
    diffServAlgDrop.EntityData.YFilter = diffServAlgDrop.YFilter
    diffServAlgDrop.EntityData.YangName = "diffServAlgDrop"
    diffServAlgDrop.EntityData.BundleName = "cisco_ios_xe"
    diffServAlgDrop.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServAlgDrop.EntityData.SegmentPath = "diffServAlgDrop"
    diffServAlgDrop.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServAlgDrop.EntityData.SegmentPath
    diffServAlgDrop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServAlgDrop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServAlgDrop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServAlgDrop.EntityData.Children = types.NewOrderedMap()
    diffServAlgDrop.EntityData.Leafs = types.NewOrderedMap()
    diffServAlgDrop.EntityData.Leafs.Append("diffServAlgDropNextFree", types.YLeaf{"DiffServAlgDropNextFree", diffServAlgDrop.DiffServAlgDropNextFree})
    diffServAlgDrop.EntityData.Leafs.Append("diffServRandomDropNextFree", types.YLeaf{"DiffServRandomDropNextFree", diffServAlgDrop.DiffServRandomDropNextFree})

    diffServAlgDrop.EntityData.YListKeys = []string {}

    return &(diffServAlgDrop.EntityData)
}

// DIFFSERVMIB_DiffServQueue
type DIFFSERVMIB_DiffServQueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServQId, or a zero to indicate
    // that none exist. The type is interface{} with range: 0..4294967295.
    DiffServQNextFree interface{}
}

func (diffServQueue *DIFFSERVMIB_DiffServQueue) GetEntityData() *types.CommonEntityData {
    diffServQueue.EntityData.YFilter = diffServQueue.YFilter
    diffServQueue.EntityData.YangName = "diffServQueue"
    diffServQueue.EntityData.BundleName = "cisco_ios_xe"
    diffServQueue.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServQueue.EntityData.SegmentPath = "diffServQueue"
    diffServQueue.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServQueue.EntityData.SegmentPath
    diffServQueue.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServQueue.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServQueue.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServQueue.EntityData.Children = types.NewOrderedMap()
    diffServQueue.EntityData.Leafs = types.NewOrderedMap()
    diffServQueue.EntityData.Leafs.Append("diffServQNextFree", types.YLeaf{"DiffServQNextFree", diffServQueue.DiffServQNextFree})

    diffServQueue.EntityData.YListKeys = []string {}

    return &(diffServQueue.EntityData)
}

// DIFFSERVMIB_DiffServScheduler
type DIFFSERVMIB_DiffServScheduler struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServSchedulerId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServSchedulerNextFree interface{}

    // This object contains an unused value for diffServMinRateId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServMinRateNextFree interface{}

    // This object contains an unused value for diffServMaxRateId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    DiffServMaxRateNextFree interface{}
}

func (diffServScheduler *DIFFSERVMIB_DiffServScheduler) GetEntityData() *types.CommonEntityData {
    diffServScheduler.EntityData.YFilter = diffServScheduler.YFilter
    diffServScheduler.EntityData.YangName = "diffServScheduler"
    diffServScheduler.EntityData.BundleName = "cisco_ios_xe"
    diffServScheduler.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServScheduler.EntityData.SegmentPath = "diffServScheduler"
    diffServScheduler.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServScheduler.EntityData.SegmentPath
    diffServScheduler.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServScheduler.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServScheduler.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServScheduler.EntityData.Children = types.NewOrderedMap()
    diffServScheduler.EntityData.Leafs = types.NewOrderedMap()
    diffServScheduler.EntityData.Leafs.Append("diffServSchedulerNextFree", types.YLeaf{"DiffServSchedulerNextFree", diffServScheduler.DiffServSchedulerNextFree})
    diffServScheduler.EntityData.Leafs.Append("diffServMinRateNextFree", types.YLeaf{"DiffServMinRateNextFree", diffServScheduler.DiffServMinRateNextFree})
    diffServScheduler.EntityData.Leafs.Append("diffServMaxRateNextFree", types.YLeaf{"DiffServMaxRateNextFree", diffServScheduler.DiffServMaxRateNextFree})

    diffServScheduler.EntityData.YListKeys = []string {}

    return &(diffServScheduler.EntityData)
}

// DIFFSERVMIB_DiffServDataPathTable
// The data path table contains RowPointers indicating the start of
// the functional data path for each interface and traffic direction
// in this device. These may merge, or be separated into parallel
// data paths.
type DIFFSERVMIB_DiffServDataPathTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the data path table indicates the start of a single
    // Differentiated Services Functional Data Path in this device.  These are
    // associated with individual interfaces, logical or physical, and therefore
    // are instantiated by ifIndex. Therefore, the interface index must have been
    // assigned, according to the procedures applicable to that, before it can be
    // meaningfully used. Generally, this means that the interface must exist. 
    // When diffServDataPathStorage is of type nonVolatile, however, this may
    // reflect the configuration for an interface whose ifIndex has been assigned
    // but for which the supporting implementation is not currently present. The
    // type is slice of DIFFSERVMIB_DiffServDataPathTable_DiffServDataPathEntry.
    DiffServDataPathEntry []*DIFFSERVMIB_DiffServDataPathTable_DiffServDataPathEntry
}

func (diffServDataPathTable *DIFFSERVMIB_DiffServDataPathTable) GetEntityData() *types.CommonEntityData {
    diffServDataPathTable.EntityData.YFilter = diffServDataPathTable.YFilter
    diffServDataPathTable.EntityData.YangName = "diffServDataPathTable"
    diffServDataPathTable.EntityData.BundleName = "cisco_ios_xe"
    diffServDataPathTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServDataPathTable.EntityData.SegmentPath = "diffServDataPathTable"
    diffServDataPathTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServDataPathTable.EntityData.SegmentPath
    diffServDataPathTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServDataPathTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServDataPathTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServDataPathTable.EntityData.Children = types.NewOrderedMap()
    diffServDataPathTable.EntityData.Children.Append("diffServDataPathEntry", types.YChild{"DiffServDataPathEntry", nil})
    for i := range diffServDataPathTable.DiffServDataPathEntry {
        diffServDataPathTable.EntityData.Children.Append(types.GetSegmentPath(diffServDataPathTable.DiffServDataPathEntry[i]), types.YChild{"DiffServDataPathEntry", diffServDataPathTable.DiffServDataPathEntry[i]})
    }
    diffServDataPathTable.EntityData.Leafs = types.NewOrderedMap()

    diffServDataPathTable.EntityData.YListKeys = []string {}

    return &(diffServDataPathTable.EntityData)
}

// DIFFSERVMIB_DiffServDataPathTable_DiffServDataPathEntry
// An entry in the data path table indicates the start of a single
// Differentiated Services Functional Data Path in this device.
// 
// These are associated with individual interfaces, logical or
// physical, and therefore are instantiated by ifIndex. Therefore,
// the interface index must have been assigned, according to the
// procedures applicable to that, before it can be meaningfully
// used. Generally, this means that the interface must exist.
// 
// When diffServDataPathStorage is of type nonVolatile, however,
// this may reflect the configuration for an interface whose ifIndex
// has been assigned but for which the supporting implementation is
// not currently present.
type DIFFSERVMIB_DiffServDataPathTable_DiffServDataPathEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. IfDirection specifies whether the reception or
    // transmission path for this interface is in view. The type is IfDirection.
    DiffServDataPathIfDirection interface{}

    // This selects the first Differentiated Services Functional Data Path Element
    // to handle traffic for this data path. This RowPointer should point to an
    // instance of one of:   diffServClfrEntry   diffServMeterEntry  
    // diffServActionEntry   diffServAlgDropEntry   diffServQEntry  A value of
    // zeroDotZero in this attribute indicates that no Differentiated Services
    // treatment is performed on traffic of this data path. A pointer with the
    // value zeroDotZero normally terminates a functional data path.  Setting this
    // to point to a target that does not exist results in an inconsistentValue
    // error.  If the row pointed to is removed or becomes inactive by other
    // means, the treatment is as if this attribute contains a value of
    // zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServDataPathStart interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServDataPathStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. The type is RowStatus.
    DiffServDataPathStatus interface{}
}

func (diffServDataPathEntry *DIFFSERVMIB_DiffServDataPathTable_DiffServDataPathEntry) GetEntityData() *types.CommonEntityData {
    diffServDataPathEntry.EntityData.YFilter = diffServDataPathEntry.YFilter
    diffServDataPathEntry.EntityData.YangName = "diffServDataPathEntry"
    diffServDataPathEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServDataPathEntry.EntityData.ParentYangName = "diffServDataPathTable"
    diffServDataPathEntry.EntityData.SegmentPath = "diffServDataPathEntry" + types.AddKeyToken(diffServDataPathEntry.IfIndex, "ifIndex") + types.AddKeyToken(diffServDataPathEntry.DiffServDataPathIfDirection, "diffServDataPathIfDirection")
    diffServDataPathEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServDataPathTable/" + diffServDataPathEntry.EntityData.SegmentPath
    diffServDataPathEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServDataPathEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServDataPathEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServDataPathEntry.EntityData.Children = types.NewOrderedMap()
    diffServDataPathEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServDataPathEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", diffServDataPathEntry.IfIndex})
    diffServDataPathEntry.EntityData.Leafs.Append("diffServDataPathIfDirection", types.YLeaf{"DiffServDataPathIfDirection", diffServDataPathEntry.DiffServDataPathIfDirection})
    diffServDataPathEntry.EntityData.Leafs.Append("diffServDataPathStart", types.YLeaf{"DiffServDataPathStart", diffServDataPathEntry.DiffServDataPathStart})
    diffServDataPathEntry.EntityData.Leafs.Append("diffServDataPathStorage", types.YLeaf{"DiffServDataPathStorage", diffServDataPathEntry.DiffServDataPathStorage})
    diffServDataPathEntry.EntityData.Leafs.Append("diffServDataPathStatus", types.YLeaf{"DiffServDataPathStatus", diffServDataPathEntry.DiffServDataPathStatus})

    diffServDataPathEntry.EntityData.YListKeys = []string {"IfIndex", "DiffServDataPathIfDirection"}

    return &(diffServDataPathEntry.EntityData)
}

// DIFFSERVMIB_DiffServClfrTable
// This table enumerates all the diffserv classifier functional
// data path elements of this device.  The actual classification
// definitions are defined in diffServClfrElementTable entries
// belonging to each classifier.
// 
// An entry in this table, pointed to by a RowPointer specifying an
// instance of diffServClfrStatus, is frequently used as the name
// for a set of classifier elements, which all use the index
// diffServClfrId. Per the semantics of the classifier element
// table, these entries constitute one or more unordered sets of
// tests which may be simultaneously applied to a message to
// 
// 
// 
// classify it.
// 
// The primary function of this table is to ensure that the value of
// diffServClfrId is unique before attempting to use it in creating
// a diffServClfrElementEntry. Therefore, the diffServClfrEntry must
// be created on the same SET as the diffServClfrElementEntry, or
// before the diffServClfrElementEntry is created.
type DIFFSERVMIB_DiffServClfrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the classifier table describes a single classifier. All
    // classifier elements belonging to the same classifier use the classifier's
    // diffServClfrId as part of their index. The type is slice of
    // DIFFSERVMIB_DiffServClfrTable_DiffServClfrEntry.
    DiffServClfrEntry []*DIFFSERVMIB_DiffServClfrTable_DiffServClfrEntry
}

func (diffServClfrTable *DIFFSERVMIB_DiffServClfrTable) GetEntityData() *types.CommonEntityData {
    diffServClfrTable.EntityData.YFilter = diffServClfrTable.YFilter
    diffServClfrTable.EntityData.YangName = "diffServClfrTable"
    diffServClfrTable.EntityData.BundleName = "cisco_ios_xe"
    diffServClfrTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServClfrTable.EntityData.SegmentPath = "diffServClfrTable"
    diffServClfrTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServClfrTable.EntityData.SegmentPath
    diffServClfrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServClfrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServClfrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServClfrTable.EntityData.Children = types.NewOrderedMap()
    diffServClfrTable.EntityData.Children.Append("diffServClfrEntry", types.YChild{"DiffServClfrEntry", nil})
    for i := range diffServClfrTable.DiffServClfrEntry {
        diffServClfrTable.EntityData.Children.Append(types.GetSegmentPath(diffServClfrTable.DiffServClfrEntry[i]), types.YChild{"DiffServClfrEntry", diffServClfrTable.DiffServClfrEntry[i]})
    }
    diffServClfrTable.EntityData.Leafs = types.NewOrderedMap()

    diffServClfrTable.EntityData.YListKeys = []string {}

    return &(diffServClfrTable.EntityData)
}

// DIFFSERVMIB_DiffServClfrTable_DiffServClfrEntry
// An entry in the classifier table describes a single classifier.
// All classifier elements belonging to the same classifier use the
// classifier's diffServClfrId as part of their index.
type DIFFSERVMIB_DiffServClfrTable_DiffServClfrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the classifier entries. 
    // Managers should obtain new values for row creation in this table by reading
    // diffServClfrNextFree. The type is interface{} with range: 1..4294967295.
    DiffServClfrId interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServClfrStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServClfrStatus interface{}
}

func (diffServClfrEntry *DIFFSERVMIB_DiffServClfrTable_DiffServClfrEntry) GetEntityData() *types.CommonEntityData {
    diffServClfrEntry.EntityData.YFilter = diffServClfrEntry.YFilter
    diffServClfrEntry.EntityData.YangName = "diffServClfrEntry"
    diffServClfrEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServClfrEntry.EntityData.ParentYangName = "diffServClfrTable"
    diffServClfrEntry.EntityData.SegmentPath = "diffServClfrEntry" + types.AddKeyToken(diffServClfrEntry.DiffServClfrId, "diffServClfrId")
    diffServClfrEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServClfrTable/" + diffServClfrEntry.EntityData.SegmentPath
    diffServClfrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServClfrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServClfrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServClfrEntry.EntityData.Children = types.NewOrderedMap()
    diffServClfrEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServClfrEntry.EntityData.Leafs.Append("diffServClfrId", types.YLeaf{"DiffServClfrId", diffServClfrEntry.DiffServClfrId})
    diffServClfrEntry.EntityData.Leafs.Append("diffServClfrStorage", types.YLeaf{"DiffServClfrStorage", diffServClfrEntry.DiffServClfrStorage})
    diffServClfrEntry.EntityData.Leafs.Append("diffServClfrStatus", types.YLeaf{"DiffServClfrStatus", diffServClfrEntry.DiffServClfrStatus})

    diffServClfrEntry.EntityData.YListKeys = []string {"DiffServClfrId"}

    return &(diffServClfrEntry.EntityData)
}

// DIFFSERVMIB_DiffServClfrElementTable
// The classifier element table enumerates the relationship between
// classification patterns and subsequent downstream Differentiated
// Services Functional Data Path elements.
// diffServClfrElementSpecific points to a filter that specifies the
// classification parameters. A classifier may use filter tables of
// different types together.
// 
// One example of a filter table defined in this MIB is
// diffServMultiFieldClfrTable, for IP Multi-Field Classifiers
// (MFCs). Such an entry might identify anything from a single
// micro-flow (an identifiable sub-session packet stream directed
// from one sending transport to the receiving transport or
// transports), or aggregates of those such as the traffic from a
// host, traffic for an application, or traffic between two hosts
// using an application and a given DSCP. The standard Behavior
// Aggregate used in the Differentiated Services Architecture is
// encoded as a degenerate case of such an aggregate - the traffic
// using a particular DSCP value.
// 
// Filter tables for other filter types may be defined elsewhere.
type DIFFSERVMIB_DiffServClfrElementTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the classifier element table describes a single element of the
    // classifier. The type is slice of
    // DIFFSERVMIB_DiffServClfrElementTable_DiffServClfrElementEntry.
    DiffServClfrElementEntry []*DIFFSERVMIB_DiffServClfrElementTable_DiffServClfrElementEntry
}

func (diffServClfrElementTable *DIFFSERVMIB_DiffServClfrElementTable) GetEntityData() *types.CommonEntityData {
    diffServClfrElementTable.EntityData.YFilter = diffServClfrElementTable.YFilter
    diffServClfrElementTable.EntityData.YangName = "diffServClfrElementTable"
    diffServClfrElementTable.EntityData.BundleName = "cisco_ios_xe"
    diffServClfrElementTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServClfrElementTable.EntityData.SegmentPath = "diffServClfrElementTable"
    diffServClfrElementTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServClfrElementTable.EntityData.SegmentPath
    diffServClfrElementTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServClfrElementTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServClfrElementTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServClfrElementTable.EntityData.Children = types.NewOrderedMap()
    diffServClfrElementTable.EntityData.Children.Append("diffServClfrElementEntry", types.YChild{"DiffServClfrElementEntry", nil})
    for i := range diffServClfrElementTable.DiffServClfrElementEntry {
        diffServClfrElementTable.EntityData.Children.Append(types.GetSegmentPath(diffServClfrElementTable.DiffServClfrElementEntry[i]), types.YChild{"DiffServClfrElementEntry", diffServClfrElementTable.DiffServClfrElementEntry[i]})
    }
    diffServClfrElementTable.EntityData.Leafs = types.NewOrderedMap()

    diffServClfrElementTable.EntityData.YListKeys = []string {}

    return &(diffServClfrElementTable.EntityData)
}

// DIFFSERVMIB_DiffServClfrElementTable_DiffServClfrElementEntry
// An entry in the classifier element table describes a single
// element of the classifier.
type DIFFSERVMIB_DiffServClfrElementTable_DiffServClfrElementEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // diffserv_mib.DIFFSERVMIB_DiffServClfrTable_DiffServClfrEntry_DiffServClfrId
    DiffServClfrId interface{}

    // This attribute is a key. An index that enumerates the Classifier Element
    // entries. Managers obtain new values for row creation in this table by
    // reading diffServClfrElementNextFree. The type is interface{} with range:
    // 1..4294967295.
    DiffServClfrElementId interface{}

    // The relative order in which classifier elements are applied: higher numbers
    // represent classifier element with higher precedence.  Classifier elements
    // with the same order must be unambiguous i.e. they must define
    // non-overlapping patterns, and are considered to be applied simultaneously
    // to the traffic stream. Classifier elements with different order may overlap
    // in their filters:  the classifier element with the highest order that
    // matches is taken.  On a given interface, there must be a complete
    // classifier in place at all times in the ingress direction.  This means one
    // or more filters must match any possible pattern. There is no such   
    // requirement in the egress direction. The type is interface{} with range:
    // 1..4294967295.
    DiffServClfrElementPrecedence interface{}

    // This attribute provides one branch of the fan-out functionality of a
    // classifier described in the Informal Differentiated Services Model section
    // 4.1.  This selects the next Differentiated Services Functional Data Path
    // Element to handle traffic for this data path. This RowPointer should point
    // to an instance of one of:   diffServClfrEntry   diffServMeterEntry  
    // diffServActionEntry   diffServAlgDropEntry   diffServQEntry  A value of
    // zeroDotZero in this attribute indicates no further Differentiated Services
    // treatment is performed on traffic of this data path. The use of zeroDotZero
    // is the normal usage for the last functional data path element of the
    // current data path.  Setting this to point to a target that does not exist
    // results in an inconsistentValue error.  If the row pointed to is removed or
    // becomes inactive by other means, the treatment is as if this attribute
    // contains a value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServClfrElementNext interface{}

    // A pointer to a valid entry in another table, filter table, that describes
    // the applicable classification parameters, e.g. an entry in
    // diffServMultiFieldClfrTable.  The value zeroDotZero is interpreted to match
    // anything not matched by another classifier element - only one such entry
    // may exist for each classifier.  Setting this to point to a target that does
    // not exist results in an inconsistentValue error.  If the row pointed to is
    // removed or    becomes inactive by other means, the element is ignored. The
    // type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServClfrElementSpecific interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServClfrElementStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServClfrElementStatus interface{}
}

func (diffServClfrElementEntry *DIFFSERVMIB_DiffServClfrElementTable_DiffServClfrElementEntry) GetEntityData() *types.CommonEntityData {
    diffServClfrElementEntry.EntityData.YFilter = diffServClfrElementEntry.YFilter
    diffServClfrElementEntry.EntityData.YangName = "diffServClfrElementEntry"
    diffServClfrElementEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServClfrElementEntry.EntityData.ParentYangName = "diffServClfrElementTable"
    diffServClfrElementEntry.EntityData.SegmentPath = "diffServClfrElementEntry" + types.AddKeyToken(diffServClfrElementEntry.DiffServClfrId, "diffServClfrId") + types.AddKeyToken(diffServClfrElementEntry.DiffServClfrElementId, "diffServClfrElementId")
    diffServClfrElementEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServClfrElementTable/" + diffServClfrElementEntry.EntityData.SegmentPath
    diffServClfrElementEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServClfrElementEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServClfrElementEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServClfrElementEntry.EntityData.Children = types.NewOrderedMap()
    diffServClfrElementEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServClfrElementEntry.EntityData.Leafs.Append("diffServClfrId", types.YLeaf{"DiffServClfrId", diffServClfrElementEntry.DiffServClfrId})
    diffServClfrElementEntry.EntityData.Leafs.Append("diffServClfrElementId", types.YLeaf{"DiffServClfrElementId", diffServClfrElementEntry.DiffServClfrElementId})
    diffServClfrElementEntry.EntityData.Leafs.Append("diffServClfrElementPrecedence", types.YLeaf{"DiffServClfrElementPrecedence", diffServClfrElementEntry.DiffServClfrElementPrecedence})
    diffServClfrElementEntry.EntityData.Leafs.Append("diffServClfrElementNext", types.YLeaf{"DiffServClfrElementNext", diffServClfrElementEntry.DiffServClfrElementNext})
    diffServClfrElementEntry.EntityData.Leafs.Append("diffServClfrElementSpecific", types.YLeaf{"DiffServClfrElementSpecific", diffServClfrElementEntry.DiffServClfrElementSpecific})
    diffServClfrElementEntry.EntityData.Leafs.Append("diffServClfrElementStorage", types.YLeaf{"DiffServClfrElementStorage", diffServClfrElementEntry.DiffServClfrElementStorage})
    diffServClfrElementEntry.EntityData.Leafs.Append("diffServClfrElementStatus", types.YLeaf{"DiffServClfrElementStatus", diffServClfrElementEntry.DiffServClfrElementStatus})

    diffServClfrElementEntry.EntityData.YListKeys = []string {"DiffServClfrId", "DiffServClfrElementId"}

    return &(diffServClfrElementEntry.EntityData)
}

// DIFFSERVMIB_DiffServMultiFieldClfrTable
// A table of IP Multi-field Classifier filter entries that a
// 
// 
// 
// system may use to identify IP traffic.
type DIFFSERVMIB_DiffServMultiFieldClfrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An IP Multi-field Classifier entry describes a single filter. The type is
    // slice of
    // DIFFSERVMIB_DiffServMultiFieldClfrTable_DiffServMultiFieldClfrEntry.
    DiffServMultiFieldClfrEntry []*DIFFSERVMIB_DiffServMultiFieldClfrTable_DiffServMultiFieldClfrEntry
}

func (diffServMultiFieldClfrTable *DIFFSERVMIB_DiffServMultiFieldClfrTable) GetEntityData() *types.CommonEntityData {
    diffServMultiFieldClfrTable.EntityData.YFilter = diffServMultiFieldClfrTable.YFilter
    diffServMultiFieldClfrTable.EntityData.YangName = "diffServMultiFieldClfrTable"
    diffServMultiFieldClfrTable.EntityData.BundleName = "cisco_ios_xe"
    diffServMultiFieldClfrTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServMultiFieldClfrTable.EntityData.SegmentPath = "diffServMultiFieldClfrTable"
    diffServMultiFieldClfrTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServMultiFieldClfrTable.EntityData.SegmentPath
    diffServMultiFieldClfrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMultiFieldClfrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMultiFieldClfrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMultiFieldClfrTable.EntityData.Children = types.NewOrderedMap()
    diffServMultiFieldClfrTable.EntityData.Children.Append("diffServMultiFieldClfrEntry", types.YChild{"DiffServMultiFieldClfrEntry", nil})
    for i := range diffServMultiFieldClfrTable.DiffServMultiFieldClfrEntry {
        diffServMultiFieldClfrTable.EntityData.Children.Append(types.GetSegmentPath(diffServMultiFieldClfrTable.DiffServMultiFieldClfrEntry[i]), types.YChild{"DiffServMultiFieldClfrEntry", diffServMultiFieldClfrTable.DiffServMultiFieldClfrEntry[i]})
    }
    diffServMultiFieldClfrTable.EntityData.Leafs = types.NewOrderedMap()

    diffServMultiFieldClfrTable.EntityData.YListKeys = []string {}

    return &(diffServMultiFieldClfrTable.EntityData)
}

// DIFFSERVMIB_DiffServMultiFieldClfrTable_DiffServMultiFieldClfrEntry
// An IP Multi-field Classifier entry describes a single filter.
type DIFFSERVMIB_DiffServMultiFieldClfrTable_DiffServMultiFieldClfrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the MultiField Classifier
    // filter entries.  Managers obtain new values for row creation in this table
    // by reading diffServMultiFieldClfrNextFree. The type is interface{} with
    // range: 1..4294967295.
    DiffServMultiFieldClfrId interface{}

    // The type of IP address used by this classifier entry.  While other types of
    // addresses are defined in the InetAddressType    textual convention, and DNS
    // names, a classifier can only look at packets on the wire. Therefore, this
    // object is limited to IPv4 and IPv6 addresses. The type is InetAddressType.
    DiffServMultiFieldClfrAddrType interface{}

    // The IP address to match against the packet's destination IP address. This
    // may not be a DNS name, but may be an IPv4 or IPv6 prefix. 
    // diffServMultiFieldClfrDstPrefixLength indicates the number of bits that are
    // relevant. The type is string with length: 0..255.
    DiffServMultiFieldClfrDstAddr interface{}

    // The length of the CIDR Prefix carried in diffServMultiFieldClfrDstAddr. In
    // IPv4 addresses, a length of 0 indicates a match of any address; a length of
    // 32 indicates a match of a single host address, and a length between 0 and
    // 32 indicates the use of a CIDR Prefix. IPv6 is similar, except that prefix
    // lengths range from 0..128. The type is interface{} with range: 0..2040.
    // Units are bits.
    DiffServMultiFieldClfrDstPrefixLength interface{}

    // The IP address to match against the packet's source IP address. This may
    // not be a DNS name, but may be an IPv4 or IPv6 prefix.
    // diffServMultiFieldClfrSrcPrefixLength indicates the number of bits that are
    // relevant. The type is string with length: 0..255.
    DiffServMultiFieldClfrSrcAddr interface{}

    // The length of the CIDR Prefix carried in diffServMultiFieldClfrSrcAddr. In
    // IPv4 addresses, a length of 0 indicates a match of any address; a length of
    // 32 indicates a match of a single host address, and a length between 0 and
    // 32 indicates the use of a CIDR Prefix. IPv6 is similar, except that prefix
    // lengths range from 0..128. The type is interface{} with range: 0..2040.
    // Units are bits.
    DiffServMultiFieldClfrSrcPrefixLength interface{}

    // The value that the DSCP in the packet must have to match this entry. A
    // value of -1 indicates that a specific DSCP value has not been defined and
    // thus all DSCP values are considered a match. The type is interface{} with
    // range: -1..63.
    DiffServMultiFieldClfrDscp interface{}

    // The flow identifier in an IPv6 header. The type is interface{} with range:
    // 0..1048575.
    DiffServMultiFieldClfrFlowId interface{}

    // The IP protocol to match against the IPv4 protocol number or the IPv6 Next-
    // Header number in the packet. A value of 255 means match all.  Note the
    // protocol number of 255 is reserved by IANA, and Next-Header number of 0 is
    // used in IPv6. The type is interface{} with range: 0..255.
    DiffServMultiFieldClfrProtocol interface{}

    // The minimum value that the layer-4 destination port number in the packet
    // must have in order to match this classifier entry. The type is interface{}
    // with range: 0..65535.
    DiffServMultiFieldClfrDstL4PortMin interface{}

    // The maximum value that the layer-4 destination port number in the packet
    // must have in order to match this classifier entry. This value must be equal
    // to or greater than the value specified for this entry in
    // diffServMultiFieldClfrDstL4PortMin. The type is interface{} with range:
    // 0..65535.
    DiffServMultiFieldClfrDstL4PortMax interface{}

    // The minimum value that the layer-4 source port number in the packet must
    // have in order to match this classifier entry. The type is interface{} with
    // range: 0..65535.
    DiffServMultiFieldClfrSrcL4PortMin interface{}

    // The maximum value that the layer-4 source port number in the packet must
    // have in order to match this classifier entry. This value must be equal to
    // or greater than the value specified for this entry in
    // diffServMultiFieldClfrSrcL4PortMin. The type is interface{} with range:
    // 0..65535.
    DiffServMultiFieldClfrSrcL4PortMax interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServMultiFieldClfrStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServMultiFieldClfrStatus interface{}
}

func (diffServMultiFieldClfrEntry *DIFFSERVMIB_DiffServMultiFieldClfrTable_DiffServMultiFieldClfrEntry) GetEntityData() *types.CommonEntityData {
    diffServMultiFieldClfrEntry.EntityData.YFilter = diffServMultiFieldClfrEntry.YFilter
    diffServMultiFieldClfrEntry.EntityData.YangName = "diffServMultiFieldClfrEntry"
    diffServMultiFieldClfrEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServMultiFieldClfrEntry.EntityData.ParentYangName = "diffServMultiFieldClfrTable"
    diffServMultiFieldClfrEntry.EntityData.SegmentPath = "diffServMultiFieldClfrEntry" + types.AddKeyToken(diffServMultiFieldClfrEntry.DiffServMultiFieldClfrId, "diffServMultiFieldClfrId")
    diffServMultiFieldClfrEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServMultiFieldClfrTable/" + diffServMultiFieldClfrEntry.EntityData.SegmentPath
    diffServMultiFieldClfrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMultiFieldClfrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMultiFieldClfrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMultiFieldClfrEntry.EntityData.Children = types.NewOrderedMap()
    diffServMultiFieldClfrEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrId", types.YLeaf{"DiffServMultiFieldClfrId", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrId})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrAddrType", types.YLeaf{"DiffServMultiFieldClfrAddrType", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrAddrType})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrDstAddr", types.YLeaf{"DiffServMultiFieldClfrDstAddr", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrDstAddr})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrDstPrefixLength", types.YLeaf{"DiffServMultiFieldClfrDstPrefixLength", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrDstPrefixLength})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrSrcAddr", types.YLeaf{"DiffServMultiFieldClfrSrcAddr", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrSrcAddr})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrSrcPrefixLength", types.YLeaf{"DiffServMultiFieldClfrSrcPrefixLength", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrSrcPrefixLength})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrDscp", types.YLeaf{"DiffServMultiFieldClfrDscp", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrDscp})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrFlowId", types.YLeaf{"DiffServMultiFieldClfrFlowId", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrFlowId})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrProtocol", types.YLeaf{"DiffServMultiFieldClfrProtocol", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrProtocol})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrDstL4PortMin", types.YLeaf{"DiffServMultiFieldClfrDstL4PortMin", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrDstL4PortMin})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrDstL4PortMax", types.YLeaf{"DiffServMultiFieldClfrDstL4PortMax", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrDstL4PortMax})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrSrcL4PortMin", types.YLeaf{"DiffServMultiFieldClfrSrcL4PortMin", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrSrcL4PortMin})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrSrcL4PortMax", types.YLeaf{"DiffServMultiFieldClfrSrcL4PortMax", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrSrcL4PortMax})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrStorage", types.YLeaf{"DiffServMultiFieldClfrStorage", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrStorage})
    diffServMultiFieldClfrEntry.EntityData.Leafs.Append("diffServMultiFieldClfrStatus", types.YLeaf{"DiffServMultiFieldClfrStatus", diffServMultiFieldClfrEntry.DiffServMultiFieldClfrStatus})

    diffServMultiFieldClfrEntry.EntityData.YListKeys = []string {"DiffServMultiFieldClfrId"}

    return &(diffServMultiFieldClfrEntry.EntityData)
}

// DIFFSERVMIB_DiffServMeterTable
// This table enumerates specific meters that a system may use to
// police a stream of traffic. The traffic stream to be metered is
// determined by the Differentiated Services Functional Data Path
// Element(s) upstream of the meter i.e. by the object(s) that point
// to each entry in this table.  This may include all traffic on an
// interface.
// 
// Specific meter details are to be found in table entry referenced
// by diffServMeterSpecific.
type DIFFSERVMIB_DiffServMeterTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the meter table describes a single conformance level of a
    // meter. The type is slice of
    // DIFFSERVMIB_DiffServMeterTable_DiffServMeterEntry.
    DiffServMeterEntry []*DIFFSERVMIB_DiffServMeterTable_DiffServMeterEntry
}

func (diffServMeterTable *DIFFSERVMIB_DiffServMeterTable) GetEntityData() *types.CommonEntityData {
    diffServMeterTable.EntityData.YFilter = diffServMeterTable.YFilter
    diffServMeterTable.EntityData.YangName = "diffServMeterTable"
    diffServMeterTable.EntityData.BundleName = "cisco_ios_xe"
    diffServMeterTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServMeterTable.EntityData.SegmentPath = "diffServMeterTable"
    diffServMeterTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServMeterTable.EntityData.SegmentPath
    diffServMeterTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMeterTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMeterTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMeterTable.EntityData.Children = types.NewOrderedMap()
    diffServMeterTable.EntityData.Children.Append("diffServMeterEntry", types.YChild{"DiffServMeterEntry", nil})
    for i := range diffServMeterTable.DiffServMeterEntry {
        diffServMeterTable.EntityData.Children.Append(types.GetSegmentPath(diffServMeterTable.DiffServMeterEntry[i]), types.YChild{"DiffServMeterEntry", diffServMeterTable.DiffServMeterEntry[i]})
    }
    diffServMeterTable.EntityData.Leafs = types.NewOrderedMap()

    diffServMeterTable.EntityData.YListKeys = []string {}

    return &(diffServMeterTable.EntityData)
}

// DIFFSERVMIB_DiffServMeterTable_DiffServMeterEntry
// An entry in the meter table describes a single conformance level
// of a meter.
type DIFFSERVMIB_DiffServMeterTable_DiffServMeterEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Meter entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServMeterNextFree. The type is interface{} with range: 1..4294967295.
    DiffServMeterId interface{}

    // If the traffic does conform, this selects the next Differentiated Services
    // Functional Data Path element to handle traffic for this data path. This
    // RowPointer should point to an instance of one of:   diffServClfrEntry  
    // diffServMeterEntry   diffServActionEntry   diffServAlgDropEntry  
    // diffServQEntry  A value of zeroDotZero in this attribute indicates that no
    // further Differentiated Services treatment is performed on traffic of this
    // data path. The use of zeroDotZero is the normal usage for the last
    // functional data path element of the current data path.  Setting this to
    // point to a target that does not exist results in an inconsistentValue
    // error.  If the row pointed to is removed or becomes inactive by other
    // means, the treatment is as if this attribute contains a value of
    // zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServMeterSucceedNext interface{}

    // If the traffic does not conform, this selects the next Differentiated
    // Services Functional Data Path element to handle traffic for this data path.
    // This RowPointer should point to an instance of one of:   diffServClfrEntry 
    // diffServMeterEntry      diffServActionEntry   diffServAlgDropEntry  
    // diffServQEntry  A value of zeroDotZero in this attribute indicates no
    // further Differentiated Services treatment is performed on traffic of this
    // data path. The use of zeroDotZero is the normal usage for the last
    // functional data path element of the current data path.  Setting this to
    // point to a target that does not exist results in an inconsistentValue
    // error.  If the row pointed to is removed or becomes inactive by other
    // means, the treatment is as if this attribute contains a value of
    // zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServMeterFailNext interface{}

    // This indicates the behavior of the meter by pointing to an entry containing
    // detailed parameters. Note that entries in that specific table must be
    // managed explicitly.  For example, diffServMeterSpecific may point to an
    // entry in diffServTBParamTable, which contains an instance of a single set
    // of Token Bucket parameters.  Setting this to point to a target that does
    // not exist results in an inconsistentValue error.  If the row pointed to is
    // removed or becomes inactive by other means, the meter always succeeds. The
    // type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServMeterSpecific interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServMeterStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServMeterStatus interface{}
}

func (diffServMeterEntry *DIFFSERVMIB_DiffServMeterTable_DiffServMeterEntry) GetEntityData() *types.CommonEntityData {
    diffServMeterEntry.EntityData.YFilter = diffServMeterEntry.YFilter
    diffServMeterEntry.EntityData.YangName = "diffServMeterEntry"
    diffServMeterEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServMeterEntry.EntityData.ParentYangName = "diffServMeterTable"
    diffServMeterEntry.EntityData.SegmentPath = "diffServMeterEntry" + types.AddKeyToken(diffServMeterEntry.DiffServMeterId, "diffServMeterId")
    diffServMeterEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServMeterTable/" + diffServMeterEntry.EntityData.SegmentPath
    diffServMeterEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMeterEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMeterEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMeterEntry.EntityData.Children = types.NewOrderedMap()
    diffServMeterEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServMeterEntry.EntityData.Leafs.Append("diffServMeterId", types.YLeaf{"DiffServMeterId", diffServMeterEntry.DiffServMeterId})
    diffServMeterEntry.EntityData.Leafs.Append("diffServMeterSucceedNext", types.YLeaf{"DiffServMeterSucceedNext", diffServMeterEntry.DiffServMeterSucceedNext})
    diffServMeterEntry.EntityData.Leafs.Append("diffServMeterFailNext", types.YLeaf{"DiffServMeterFailNext", diffServMeterEntry.DiffServMeterFailNext})
    diffServMeterEntry.EntityData.Leafs.Append("diffServMeterSpecific", types.YLeaf{"DiffServMeterSpecific", diffServMeterEntry.DiffServMeterSpecific})
    diffServMeterEntry.EntityData.Leafs.Append("diffServMeterStorage", types.YLeaf{"DiffServMeterStorage", diffServMeterEntry.DiffServMeterStorage})
    diffServMeterEntry.EntityData.Leafs.Append("diffServMeterStatus", types.YLeaf{"DiffServMeterStatus", diffServMeterEntry.DiffServMeterStatus})

    diffServMeterEntry.EntityData.YListKeys = []string {"DiffServMeterId"}

    return &(diffServMeterEntry.EntityData)
}

// DIFFSERVMIB_DiffServTBParamTable
// This table enumerates a single set of token bucket meter
// parameters that a system may use to police a stream of traffic.
// Such meters are modeled here as having a single rate and a single
// burst size. Multiple entries are used when multiple rates/burst
// sizes are needed.
type DIFFSERVMIB_DiffServTBParamTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry that describes a single set of token bucket parameters. The type
    // is slice of DIFFSERVMIB_DiffServTBParamTable_DiffServTBParamEntry.
    DiffServTBParamEntry []*DIFFSERVMIB_DiffServTBParamTable_DiffServTBParamEntry
}

func (diffServTBParamTable *DIFFSERVMIB_DiffServTBParamTable) GetEntityData() *types.CommonEntityData {
    diffServTBParamTable.EntityData.YFilter = diffServTBParamTable.YFilter
    diffServTBParamTable.EntityData.YangName = "diffServTBParamTable"
    diffServTBParamTable.EntityData.BundleName = "cisco_ios_xe"
    diffServTBParamTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServTBParamTable.EntityData.SegmentPath = "diffServTBParamTable"
    diffServTBParamTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServTBParamTable.EntityData.SegmentPath
    diffServTBParamTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServTBParamTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServTBParamTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServTBParamTable.EntityData.Children = types.NewOrderedMap()
    diffServTBParamTable.EntityData.Children.Append("diffServTBParamEntry", types.YChild{"DiffServTBParamEntry", nil})
    for i := range diffServTBParamTable.DiffServTBParamEntry {
        diffServTBParamTable.EntityData.Children.Append(types.GetSegmentPath(diffServTBParamTable.DiffServTBParamEntry[i]), types.YChild{"DiffServTBParamEntry", diffServTBParamTable.DiffServTBParamEntry[i]})
    }
    diffServTBParamTable.EntityData.Leafs = types.NewOrderedMap()

    diffServTBParamTable.EntityData.YListKeys = []string {}

    return &(diffServTBParamTable.EntityData)
}

// DIFFSERVMIB_DiffServTBParamTable_DiffServTBParamEntry
// An entry that describes a single set of token bucket
// parameters.
type DIFFSERVMIB_DiffServTBParamTable_DiffServTBParamEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Token Bucket
    // Parameter entries. Managers obtain new values for row creation in this
    // table by reading diffServTBParamNextFree. The type is interface{} with
    // range: 1..4294967295.
    DiffServTBParamId interface{}

    // The Metering algorithm associated with the Token Bucket parameters. 
    // zeroDotZero indicates this is unknown.  Standard values for generic
    // algorithms: diffServTBParamSimpleTokenBucket, diffServTBParamAvgRate,
    // diffServTBParamSrTCMBlind, diffServTBParamSrTCMAware,
    // diffServTBParamTrTCMBlind, diffServTBParamTrTCMAware, and
    // diffServTBParamTswTCM are specified in this MIB as OBJECT- IDENTITYs;
    // additional values may be further specified in other MIBs. The type is
    // string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServTBParamType interface{}

    // The token-bucket rate, in kilobits per second (kbps). This attribute is
    // used for: 1. CIR in RFC 2697 for srTCM 2. CIR and PIR in RFC 2698 for trTCM
    // 3. CTR and PTR in RFC 2859 for TSWTCM 4. AverageRate in RFC 3290. The type
    // is interface{} with range: 1..4294967295. Units are kilobits per second.
    DiffServTBParamRate interface{}

    // The maximum number of bytes in a single transmission burst. This attribute
    // is used for: 1. CBS and EBS in RFC 2697 for srTCM 2. CBS and PBS in RFC
    // 2698 for trTCM 3. Burst Size in RFC 3290. The type is interface{} with
    // range: 0..2147483647. Units are Bytes.
    DiffServTBParamBurstSize interface{}

    // The time interval used with the token bucket.  For: 1. Average Rate Meter,
    // the Informal Differentiated Services Model    section 5.2.1, - Delta. 2.
    // Simple Token Bucket Meter, the Informal Differentiated    Services Model
    // section 5.1, - time interval t. 3. RFC 2859 TSWTCM, - AVG_INTERVAL. 4. RFC
    // 2697 srTCM, RFC 2698 trTCM, - token bucket update time    interval. The
    // type is interface{} with range: 1..4294967295. Units are microseconds.
    DiffServTBParamInterval interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServTBParamStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServTBParamStatus interface{}
}

func (diffServTBParamEntry *DIFFSERVMIB_DiffServTBParamTable_DiffServTBParamEntry) GetEntityData() *types.CommonEntityData {
    diffServTBParamEntry.EntityData.YFilter = diffServTBParamEntry.YFilter
    diffServTBParamEntry.EntityData.YangName = "diffServTBParamEntry"
    diffServTBParamEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServTBParamEntry.EntityData.ParentYangName = "diffServTBParamTable"
    diffServTBParamEntry.EntityData.SegmentPath = "diffServTBParamEntry" + types.AddKeyToken(diffServTBParamEntry.DiffServTBParamId, "diffServTBParamId")
    diffServTBParamEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServTBParamTable/" + diffServTBParamEntry.EntityData.SegmentPath
    diffServTBParamEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServTBParamEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServTBParamEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServTBParamEntry.EntityData.Children = types.NewOrderedMap()
    diffServTBParamEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServTBParamEntry.EntityData.Leafs.Append("diffServTBParamId", types.YLeaf{"DiffServTBParamId", diffServTBParamEntry.DiffServTBParamId})
    diffServTBParamEntry.EntityData.Leafs.Append("diffServTBParamType", types.YLeaf{"DiffServTBParamType", diffServTBParamEntry.DiffServTBParamType})
    diffServTBParamEntry.EntityData.Leafs.Append("diffServTBParamRate", types.YLeaf{"DiffServTBParamRate", diffServTBParamEntry.DiffServTBParamRate})
    diffServTBParamEntry.EntityData.Leafs.Append("diffServTBParamBurstSize", types.YLeaf{"DiffServTBParamBurstSize", diffServTBParamEntry.DiffServTBParamBurstSize})
    diffServTBParamEntry.EntityData.Leafs.Append("diffServTBParamInterval", types.YLeaf{"DiffServTBParamInterval", diffServTBParamEntry.DiffServTBParamInterval})
    diffServTBParamEntry.EntityData.Leafs.Append("diffServTBParamStorage", types.YLeaf{"DiffServTBParamStorage", diffServTBParamEntry.DiffServTBParamStorage})
    diffServTBParamEntry.EntityData.Leafs.Append("diffServTBParamStatus", types.YLeaf{"DiffServTBParamStatus", diffServTBParamEntry.DiffServTBParamStatus})

    diffServTBParamEntry.EntityData.YListKeys = []string {"DiffServTBParamId"}

    return &(diffServTBParamEntry.EntityData)
}

// DIFFSERVMIB_DiffServActionTable
// The Action Table enumerates actions that can be performed to a
// stream of traffic. Multiple actions can be concatenated. For
// example, traffic exiting from a meter may be counted, marked, and
// potentially dropped before entering a queue.
// 
// Specific actions are indicated by diffServActionSpecific which
// points to an entry of a specific action type parameterizing the
// action in detail.
type DIFFSERVMIB_DiffServActionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in the action table allows description of one specific action to
    // be applied to traffic. The type is slice of
    // DIFFSERVMIB_DiffServActionTable_DiffServActionEntry.
    DiffServActionEntry []*DIFFSERVMIB_DiffServActionTable_DiffServActionEntry
}

func (diffServActionTable *DIFFSERVMIB_DiffServActionTable) GetEntityData() *types.CommonEntityData {
    diffServActionTable.EntityData.YFilter = diffServActionTable.YFilter
    diffServActionTable.EntityData.YangName = "diffServActionTable"
    diffServActionTable.EntityData.BundleName = "cisco_ios_xe"
    diffServActionTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServActionTable.EntityData.SegmentPath = "diffServActionTable"
    diffServActionTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServActionTable.EntityData.SegmentPath
    diffServActionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServActionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServActionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServActionTable.EntityData.Children = types.NewOrderedMap()
    diffServActionTable.EntityData.Children.Append("diffServActionEntry", types.YChild{"DiffServActionEntry", nil})
    for i := range diffServActionTable.DiffServActionEntry {
        diffServActionTable.EntityData.Children.Append(types.GetSegmentPath(diffServActionTable.DiffServActionEntry[i]), types.YChild{"DiffServActionEntry", diffServActionTable.DiffServActionEntry[i]})
    }
    diffServActionTable.EntityData.Leafs = types.NewOrderedMap()

    diffServActionTable.EntityData.YListKeys = []string {}

    return &(diffServActionTable.EntityData)
}

// DIFFSERVMIB_DiffServActionTable_DiffServActionEntry
// Each entry in the action table allows description of one
// specific action to be applied to traffic.
type DIFFSERVMIB_DiffServActionTable_DiffServActionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Action entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServActionNextFree. The type is interface{} with range: 1..4294967295.
    DiffServActionId interface{}

    // The interface index (value of ifIndex) that this action occurs on. This may
    // be derived from the diffServDataPathStartEntry's index by extension through
    // the various RowPointers. However, as this may be difficult for a network
    // management station, it is placed here as well.  If this is indeterminate,
    // the value is zero.  This is of especial relevance when reporting the
    // counters which may apply to traffic crossing an interface:   
    // diffServCountActOctets,    diffServCountActPkts,    diffServAlgDropOctets, 
    // diffServAlgDropPkts,    diffServAlgRandomDropOctets, and   
    // diffServAlgRandomDropPkts.  It is also especially relevant to the queue and
    // scheduler which may be subsequently applied. The type is interface{} with
    // range: 0..2147483647.
    DiffServActionInterface interface{}

    // This selects the next Differentiated Services Functional Data Path Element
    // to handle traffic for this data path. This RowPointer should point to an
    // instance of one of:   diffServClfrEntry   diffServMeterEntry  
    // diffServActionEntry   diffServAlgDropEntry   diffServQEntry  A value of
    // zeroDotZero in this attribute indicates no further Differentiated Services
    // treatment is performed on traffic of this data path. The use of zeroDotZero
    // is the normal usage for the last functional data path element of the
    // current data path.  Setting this to point to a target that does not exist
    // results in an inconsistentValue error.  If the row pointed to is removed or
    // becomes inactive by other means, the treatment is as if this attribute
    // contains a value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServActionNext interface{}

    // A pointer to an object instance providing additional information for the
    // type of action indicated by this action table entry.  For the standard
    // actions defined by this MIB module, this should point to either a
    // diffServDscpMarkActEntry or a diffServCountActEntry. For other actions, it
    // may point to an object instance defined in some other MIB.  Setting this to
    // point to a target that does not exist results in an inconsistentValue
    // error.  If the row pointed to is removed or becomes inactive by other
    // means, the Meter should be treated as if it were not present.  This may
    // lead to incorrect policy behavior. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServActionSpecific interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServActionStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServActionStatus interface{}
}

func (diffServActionEntry *DIFFSERVMIB_DiffServActionTable_DiffServActionEntry) GetEntityData() *types.CommonEntityData {
    diffServActionEntry.EntityData.YFilter = diffServActionEntry.YFilter
    diffServActionEntry.EntityData.YangName = "diffServActionEntry"
    diffServActionEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServActionEntry.EntityData.ParentYangName = "diffServActionTable"
    diffServActionEntry.EntityData.SegmentPath = "diffServActionEntry" + types.AddKeyToken(diffServActionEntry.DiffServActionId, "diffServActionId")
    diffServActionEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServActionTable/" + diffServActionEntry.EntityData.SegmentPath
    diffServActionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServActionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServActionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServActionEntry.EntityData.Children = types.NewOrderedMap()
    diffServActionEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServActionEntry.EntityData.Leafs.Append("diffServActionId", types.YLeaf{"DiffServActionId", diffServActionEntry.DiffServActionId})
    diffServActionEntry.EntityData.Leafs.Append("diffServActionInterface", types.YLeaf{"DiffServActionInterface", diffServActionEntry.DiffServActionInterface})
    diffServActionEntry.EntityData.Leafs.Append("diffServActionNext", types.YLeaf{"DiffServActionNext", diffServActionEntry.DiffServActionNext})
    diffServActionEntry.EntityData.Leafs.Append("diffServActionSpecific", types.YLeaf{"DiffServActionSpecific", diffServActionEntry.DiffServActionSpecific})
    diffServActionEntry.EntityData.Leafs.Append("diffServActionStorage", types.YLeaf{"DiffServActionStorage", diffServActionEntry.DiffServActionStorage})
    diffServActionEntry.EntityData.Leafs.Append("diffServActionStatus", types.YLeaf{"DiffServActionStatus", diffServActionEntry.DiffServActionStatus})

    diffServActionEntry.EntityData.YListKeys = []string {"DiffServActionId"}

    return &(diffServActionEntry.EntityData)
}

// DIFFSERVMIB_DiffServDscpMarkActTable
// This table enumerates specific DSCPs used for marking or
// remarking the DSCP field of IP packets. The entries of this table
// may be referenced by a diffServActionSpecific attribute.
type DIFFSERVMIB_DiffServDscpMarkActTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DSCP mark action table that describes a single DSCP used
    // for marking. The type is slice of
    // DIFFSERVMIB_DiffServDscpMarkActTable_DiffServDscpMarkActEntry.
    DiffServDscpMarkActEntry []*DIFFSERVMIB_DiffServDscpMarkActTable_DiffServDscpMarkActEntry
}

func (diffServDscpMarkActTable *DIFFSERVMIB_DiffServDscpMarkActTable) GetEntityData() *types.CommonEntityData {
    diffServDscpMarkActTable.EntityData.YFilter = diffServDscpMarkActTable.YFilter
    diffServDscpMarkActTable.EntityData.YangName = "diffServDscpMarkActTable"
    diffServDscpMarkActTable.EntityData.BundleName = "cisco_ios_xe"
    diffServDscpMarkActTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServDscpMarkActTable.EntityData.SegmentPath = "diffServDscpMarkActTable"
    diffServDscpMarkActTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServDscpMarkActTable.EntityData.SegmentPath
    diffServDscpMarkActTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServDscpMarkActTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServDscpMarkActTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServDscpMarkActTable.EntityData.Children = types.NewOrderedMap()
    diffServDscpMarkActTable.EntityData.Children.Append("diffServDscpMarkActEntry", types.YChild{"DiffServDscpMarkActEntry", nil})
    for i := range diffServDscpMarkActTable.DiffServDscpMarkActEntry {
        diffServDscpMarkActTable.EntityData.Children.Append(types.GetSegmentPath(diffServDscpMarkActTable.DiffServDscpMarkActEntry[i]), types.YChild{"DiffServDscpMarkActEntry", diffServDscpMarkActTable.DiffServDscpMarkActEntry[i]})
    }
    diffServDscpMarkActTable.EntityData.Leafs = types.NewOrderedMap()

    diffServDscpMarkActTable.EntityData.YListKeys = []string {}

    return &(diffServDscpMarkActTable.EntityData)
}

// DIFFSERVMIB_DiffServDscpMarkActTable_DiffServDscpMarkActEntry
// An entry in the DSCP mark action table that describes a single
// DSCP used for marking.
type DIFFSERVMIB_DiffServDscpMarkActTable_DiffServDscpMarkActEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The DSCP that this Action will store into the DSCP
    // field of the subject. It is quite possible that the only packets subject to
    // this Action are already marked with this DSCP. Note also that
    // Differentiated Services processing may result in packet being marked on
    // both ingress to a network and on egress from it, and that ingress and
    // egress can occur in the same router. The type is interface{} with range:
    // 0..63.
    DiffServDscpMarkActDscp interface{}
}

func (diffServDscpMarkActEntry *DIFFSERVMIB_DiffServDscpMarkActTable_DiffServDscpMarkActEntry) GetEntityData() *types.CommonEntityData {
    diffServDscpMarkActEntry.EntityData.YFilter = diffServDscpMarkActEntry.YFilter
    diffServDscpMarkActEntry.EntityData.YangName = "diffServDscpMarkActEntry"
    diffServDscpMarkActEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServDscpMarkActEntry.EntityData.ParentYangName = "diffServDscpMarkActTable"
    diffServDscpMarkActEntry.EntityData.SegmentPath = "diffServDscpMarkActEntry" + types.AddKeyToken(diffServDscpMarkActEntry.DiffServDscpMarkActDscp, "diffServDscpMarkActDscp")
    diffServDscpMarkActEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServDscpMarkActTable/" + diffServDscpMarkActEntry.EntityData.SegmentPath
    diffServDscpMarkActEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServDscpMarkActEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServDscpMarkActEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServDscpMarkActEntry.EntityData.Children = types.NewOrderedMap()
    diffServDscpMarkActEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServDscpMarkActEntry.EntityData.Leafs.Append("diffServDscpMarkActDscp", types.YLeaf{"DiffServDscpMarkActDscp", diffServDscpMarkActEntry.DiffServDscpMarkActDscp})

    diffServDscpMarkActEntry.EntityData.YListKeys = []string {"DiffServDscpMarkActDscp"}

    return &(diffServDscpMarkActEntry.EntityData)
}

// DIFFSERVMIB_DiffServCountActTable
// This table contains counters for all the traffic passing through
// an action element.
type DIFFSERVMIB_DiffServCountActTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the count action table describes a single set of traffic
    // counters. The type is slice of
    // DIFFSERVMIB_DiffServCountActTable_DiffServCountActEntry.
    DiffServCountActEntry []*DIFFSERVMIB_DiffServCountActTable_DiffServCountActEntry
}

func (diffServCountActTable *DIFFSERVMIB_DiffServCountActTable) GetEntityData() *types.CommonEntityData {
    diffServCountActTable.EntityData.YFilter = diffServCountActTable.YFilter
    diffServCountActTable.EntityData.YangName = "diffServCountActTable"
    diffServCountActTable.EntityData.BundleName = "cisco_ios_xe"
    diffServCountActTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServCountActTable.EntityData.SegmentPath = "diffServCountActTable"
    diffServCountActTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServCountActTable.EntityData.SegmentPath
    diffServCountActTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServCountActTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServCountActTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServCountActTable.EntityData.Children = types.NewOrderedMap()
    diffServCountActTable.EntityData.Children.Append("diffServCountActEntry", types.YChild{"DiffServCountActEntry", nil})
    for i := range diffServCountActTable.DiffServCountActEntry {
        diffServCountActTable.EntityData.Children.Append(types.GetSegmentPath(diffServCountActTable.DiffServCountActEntry[i]), types.YChild{"DiffServCountActEntry", diffServCountActTable.DiffServCountActEntry[i]})
    }
    diffServCountActTable.EntityData.Leafs = types.NewOrderedMap()

    diffServCountActTable.EntityData.YListKeys = []string {}

    return &(diffServCountActTable.EntityData)
}

// DIFFSERVMIB_DiffServCountActTable_DiffServCountActEntry
// An entry in the count action table describes a single set of
// traffic counters.
type DIFFSERVMIB_DiffServCountActTable_DiffServCountActEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Count Action entries.
    // Managers obtain new values for row creation in this table by reading   
    // diffServCountActNextFree. The type is interface{} with range:
    // 1..4294967295.
    DiffServCountActId interface{}

    // The number of octets at the Action data path element.  Discontinuities in
    // the value of this counter can occur at re- initialization of the management
    // system and at other times as indicated by the value of
    // ifCounterDiscontinuityTime on the relevant interface. The type is
    // interface{} with range: 0..18446744073709551615.
    DiffServCountActOctets interface{}

    // The number of packets at the Action data path element.  Discontinuities in
    // the value of this counter can occur at re- initialization of the management
    // system and at other times as indicated by the value of
    // ifCounterDiscontinuityTime on the relevant interface. The type is
    // interface{} with range: 0..18446744073709551615.
    DiffServCountActPkts interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServCountActStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing    to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServCountActStatus interface{}
}

func (diffServCountActEntry *DIFFSERVMIB_DiffServCountActTable_DiffServCountActEntry) GetEntityData() *types.CommonEntityData {
    diffServCountActEntry.EntityData.YFilter = diffServCountActEntry.YFilter
    diffServCountActEntry.EntityData.YangName = "diffServCountActEntry"
    diffServCountActEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServCountActEntry.EntityData.ParentYangName = "diffServCountActTable"
    diffServCountActEntry.EntityData.SegmentPath = "diffServCountActEntry" + types.AddKeyToken(diffServCountActEntry.DiffServCountActId, "diffServCountActId")
    diffServCountActEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServCountActTable/" + diffServCountActEntry.EntityData.SegmentPath
    diffServCountActEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServCountActEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServCountActEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServCountActEntry.EntityData.Children = types.NewOrderedMap()
    diffServCountActEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServCountActEntry.EntityData.Leafs.Append("diffServCountActId", types.YLeaf{"DiffServCountActId", diffServCountActEntry.DiffServCountActId})
    diffServCountActEntry.EntityData.Leafs.Append("diffServCountActOctets", types.YLeaf{"DiffServCountActOctets", diffServCountActEntry.DiffServCountActOctets})
    diffServCountActEntry.EntityData.Leafs.Append("diffServCountActPkts", types.YLeaf{"DiffServCountActPkts", diffServCountActEntry.DiffServCountActPkts})
    diffServCountActEntry.EntityData.Leafs.Append("diffServCountActStorage", types.YLeaf{"DiffServCountActStorage", diffServCountActEntry.DiffServCountActStorage})
    diffServCountActEntry.EntityData.Leafs.Append("diffServCountActStatus", types.YLeaf{"DiffServCountActStatus", diffServCountActEntry.DiffServCountActStatus})

    diffServCountActEntry.EntityData.YListKeys = []string {"DiffServCountActId"}

    return &(diffServCountActEntry.EntityData)
}

// DIFFSERVMIB_DiffServAlgDropTable
// The algorithmic drop table contains entries describing an
// element that drops packets according to some algorithm.
type DIFFSERVMIB_DiffServAlgDropTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry describes a process that drops packets according to some
    // algorithm. Further details of the algorithm type are to be found in
    // diffServAlgDropType and with more detail parameter entry pointed to by
    // diffServAlgDropSpecific when necessary. The type is slice of
    // DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry.
    DiffServAlgDropEntry []*DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry
}

func (diffServAlgDropTable *DIFFSERVMIB_DiffServAlgDropTable) GetEntityData() *types.CommonEntityData {
    diffServAlgDropTable.EntityData.YFilter = diffServAlgDropTable.YFilter
    diffServAlgDropTable.EntityData.YangName = "diffServAlgDropTable"
    diffServAlgDropTable.EntityData.BundleName = "cisco_ios_xe"
    diffServAlgDropTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServAlgDropTable.EntityData.SegmentPath = "diffServAlgDropTable"
    diffServAlgDropTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServAlgDropTable.EntityData.SegmentPath
    diffServAlgDropTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServAlgDropTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServAlgDropTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServAlgDropTable.EntityData.Children = types.NewOrderedMap()
    diffServAlgDropTable.EntityData.Children.Append("diffServAlgDropEntry", types.YChild{"DiffServAlgDropEntry", nil})
    for i := range diffServAlgDropTable.DiffServAlgDropEntry {
        diffServAlgDropTable.EntityData.Children.Append(types.GetSegmentPath(diffServAlgDropTable.DiffServAlgDropEntry[i]), types.YChild{"DiffServAlgDropEntry", diffServAlgDropTable.DiffServAlgDropEntry[i]})
    }
    diffServAlgDropTable.EntityData.Leafs = types.NewOrderedMap()

    diffServAlgDropTable.EntityData.YListKeys = []string {}

    return &(diffServAlgDropTable.EntityData)
}

// DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry
// An entry describes a process that drops packets according to
// some algorithm. Further details of the algorithm type are to be
// found in diffServAlgDropType and with more detail parameter entry
// pointed to by diffServAlgDropSpecific when necessary.
type DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Algorithmic Dropper
    // entries. Managers obtain new values for row creation in this table by
    // reading diffServAlgDropNextFree. The type is interface{} with range:
    // 1..4294967295.
    DiffServAlgDropId interface{}

    // The type of algorithm used by this dropper. The value other(1) requires
    // further specification in some other MIB module.  In the tailDrop(2)
    // algorithm, diffServAlgDropQThreshold represents the maximum depth of the
    // queue, pointed to by diffServAlgDropQMeasure, beyond which all newly
    // arriving packets will be dropped.  In the headDrop(3) algorithm, if a
    // packet arrives when the current depth of the queue, pointed to by
    // diffServAlgDropQMeasure, is at diffServAlgDropQThreshold, packets currently
    // at the head of the queue are dropped to make room for the new packet to be
    // enqueued at the tail of the queue.  In the randomDrop(4) algorithm, on
    // packet arrival, an Active Queue Management algorithm is executed which may
    // randomly drop a packet. This algorithm may be proprietary, and it may drop
    // either the arriving packet or another packet in the queue.
    // diffServAlgDropSpecific points to a diffServRandomDropEntry that describes
    // the algorithm. For this algorithm,    diffServAlgDropQThreshold is
    // understood to be the absolute maximum size of the queue and additional
    // parameters are described in diffServRandomDropTable.  The alwaysDrop(5)
    // algorithm is as its name specifies; always drop. In this case, the other
    // configuration values in this Entry are not meaningful; There is no useful
    // 'next' processing step, there is no queue, and parameters describing the
    // queue are not useful. Therefore, diffServAlgDropNext,
    // diffServAlgDropMeasure, and diffServAlgDropSpecific are all zeroDotZero.
    // The type is DiffServAlgDropType.
    DiffServAlgDropType interface{}

    // This selects the next Differentiated Services Functional Data Path Element
    // to handle traffic for this data path. This RowPointer should point to an
    // instance of one of:   diffServClfrEntry   diffServMeterEntry  
    // diffServActionEntry   diffServQEntry  A value of zeroDotZero in this
    // attribute indicates no further Differentiated Services treatment is
    // performed on traffic of this data path. The use of zeroDotZero is the
    // normal usage for the last functional data path element of the current data
    // path.  When diffServAlgDropType is alwaysDrop(5), this object is ignored. 
    // Setting this to point to a target that does not exist results in an
    // inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServAlgDropNext interface{}

    // Points to an entry in the diffServQTable to indicate the queue that a drop
    // algorithm is to monitor when deciding whether to drop a packet. If the row
    // pointed to does not exist, the algorithmic dropper element is considered
    // inactive.    Setting this to point to a target that does not exist results
    // in an inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServAlgDropQMeasure interface{}

    // A threshold on the depth in bytes of the queue being measured at which a
    // trigger is generated to the dropping algorithm, unless diffServAlgDropType
    // is alwaysDrop(5) where this object is ignored.  For the tailDrop(2) or
    // headDrop(3) algorithms, this represents the depth of the queue, pointed to
    // by diffServAlgDropQMeasure, at which the drop action will take place. Other
    // algorithms will need to define their own semantics for this threshold. The
    // type is interface{} with range: 1..4294967295. Units are Bytes.
    DiffServAlgDropQThreshold interface{}

    // Points to a table entry that provides further detail regarding a drop
    // algorithm.  Entries with diffServAlgDropType equal to other(1) may have
    // this point to a table defined in another MIB module.  Entries with
    // diffServAlgDropType equal to randomDrop(4) must have this point to an entry
    // in diffServRandomDropTable.  For all other algorithms specified in this
    // MIB, this should take the value zeroDotZero.  The diffServAlgDropType is
    // authoritative for the type of the drop algorithm and the specific
    // parameters for the drop algorithm needs to be evaluated based on the
    // diffServAlgDropType.  Setting this to point to a target that does not exist
    // results in an inconsistentValue error.  If the row pointed to is removed or
    // becomes inactive by other means, the treatment is as if this attribute
    // contains a value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServAlgDropSpecific interface{}

    // The number of octets that have been deterministically dropped by this drop
    // process.  Discontinuities in the value of this counter can occur at re-
    // initialization of the management system and at other times as indicated by
    // the value of ifCounterDiscontinuityTime on the relevant interface. The type
    // is interface{} with range: 0..18446744073709551615.
    DiffServAlgDropOctets interface{}

    // The number of packets that have been deterministically dropped by this drop
    // process.  Discontinuities in the value of this counter can occur at re-
    // initialization of the management system and at other times as indicated by
    // the value of ifCounterDiscontinuityTime on the relevant interface. The type
    // is interface{} with range: 0..18446744073709551615.
    DiffServAlgDropPkts interface{}

    // The number of octets that have been randomly dropped by this drop process. 
    // This counter applies, therefore, only to random droppers.  Discontinuities
    // in the value of this counter can occur at re- initialization of the
    // management system and at other times as indicated by the value of
    // ifCounterDiscontinuityTime on the relevant interface. The type is
    // interface{} with range: 0..18446744073709551615.
    DiffServAlgRandomDropOctets interface{}

    // The number of packets that have been randomly dropped by this drop process.
    // This counter applies, therefore, only to random droppers.  Discontinuities
    // in the value of this counter can occur at re- initialization of the
    // management system and at other times as indicated by the value of
    // ifCounterDiscontinuityTime on the relevant interface. The type is
    // interface{} with range: 0..18446744073709551615.
    DiffServAlgRandomDropPkts interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServAlgDropStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServAlgDropStatus interface{}
}

func (diffServAlgDropEntry *DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry) GetEntityData() *types.CommonEntityData {
    diffServAlgDropEntry.EntityData.YFilter = diffServAlgDropEntry.YFilter
    diffServAlgDropEntry.EntityData.YangName = "diffServAlgDropEntry"
    diffServAlgDropEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServAlgDropEntry.EntityData.ParentYangName = "diffServAlgDropTable"
    diffServAlgDropEntry.EntityData.SegmentPath = "diffServAlgDropEntry" + types.AddKeyToken(diffServAlgDropEntry.DiffServAlgDropId, "diffServAlgDropId")
    diffServAlgDropEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServAlgDropTable/" + diffServAlgDropEntry.EntityData.SegmentPath
    diffServAlgDropEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServAlgDropEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServAlgDropEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServAlgDropEntry.EntityData.Children = types.NewOrderedMap()
    diffServAlgDropEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropId", types.YLeaf{"DiffServAlgDropId", diffServAlgDropEntry.DiffServAlgDropId})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropType", types.YLeaf{"DiffServAlgDropType", diffServAlgDropEntry.DiffServAlgDropType})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropNext", types.YLeaf{"DiffServAlgDropNext", diffServAlgDropEntry.DiffServAlgDropNext})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropQMeasure", types.YLeaf{"DiffServAlgDropQMeasure", diffServAlgDropEntry.DiffServAlgDropQMeasure})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropQThreshold", types.YLeaf{"DiffServAlgDropQThreshold", diffServAlgDropEntry.DiffServAlgDropQThreshold})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropSpecific", types.YLeaf{"DiffServAlgDropSpecific", diffServAlgDropEntry.DiffServAlgDropSpecific})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropOctets", types.YLeaf{"DiffServAlgDropOctets", diffServAlgDropEntry.DiffServAlgDropOctets})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropPkts", types.YLeaf{"DiffServAlgDropPkts", diffServAlgDropEntry.DiffServAlgDropPkts})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgRandomDropOctets", types.YLeaf{"DiffServAlgRandomDropOctets", diffServAlgDropEntry.DiffServAlgRandomDropOctets})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgRandomDropPkts", types.YLeaf{"DiffServAlgRandomDropPkts", diffServAlgDropEntry.DiffServAlgRandomDropPkts})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropStorage", types.YLeaf{"DiffServAlgDropStorage", diffServAlgDropEntry.DiffServAlgDropStorage})
    diffServAlgDropEntry.EntityData.Leafs.Append("diffServAlgDropStatus", types.YLeaf{"DiffServAlgDropStatus", diffServAlgDropEntry.DiffServAlgDropStatus})

    diffServAlgDropEntry.EntityData.YListKeys = []string {"DiffServAlgDropId"}

    return &(diffServAlgDropEntry.EntityData)
}

// DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType represents and diffServAlgDropSpecific are all zeroDotZero.
type DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType string

const (
    DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType_other DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType = "other"

    DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType_tailDrop DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType = "tailDrop"

    DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType_headDrop DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType = "headDrop"

    DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType_randomDrop DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType = "randomDrop"

    DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType_alwaysDrop DIFFSERVMIB_DiffServAlgDropTable_DiffServAlgDropEntry_DiffServAlgDropType = "alwaysDrop"
)

// DIFFSERVMIB_DiffServRandomDropTable
// The random drop table contains entries describing a process that
// drops packets randomly. Entries in this table are pointed to by
// diffServAlgDropSpecific.
type DIFFSERVMIB_DiffServRandomDropTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry describes a process that drops packets according to a random
    // algorithm. The type is slice of
    // DIFFSERVMIB_DiffServRandomDropTable_DiffServRandomDropEntry.
    DiffServRandomDropEntry []*DIFFSERVMIB_DiffServRandomDropTable_DiffServRandomDropEntry
}

func (diffServRandomDropTable *DIFFSERVMIB_DiffServRandomDropTable) GetEntityData() *types.CommonEntityData {
    diffServRandomDropTable.EntityData.YFilter = diffServRandomDropTable.YFilter
    diffServRandomDropTable.EntityData.YangName = "diffServRandomDropTable"
    diffServRandomDropTable.EntityData.BundleName = "cisco_ios_xe"
    diffServRandomDropTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServRandomDropTable.EntityData.SegmentPath = "diffServRandomDropTable"
    diffServRandomDropTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServRandomDropTable.EntityData.SegmentPath
    diffServRandomDropTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServRandomDropTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServRandomDropTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServRandomDropTable.EntityData.Children = types.NewOrderedMap()
    diffServRandomDropTable.EntityData.Children.Append("diffServRandomDropEntry", types.YChild{"DiffServRandomDropEntry", nil})
    for i := range diffServRandomDropTable.DiffServRandomDropEntry {
        diffServRandomDropTable.EntityData.Children.Append(types.GetSegmentPath(diffServRandomDropTable.DiffServRandomDropEntry[i]), types.YChild{"DiffServRandomDropEntry", diffServRandomDropTable.DiffServRandomDropEntry[i]})
    }
    diffServRandomDropTable.EntityData.Leafs = types.NewOrderedMap()

    diffServRandomDropTable.EntityData.YListKeys = []string {}

    return &(diffServRandomDropTable.EntityData)
}

// DIFFSERVMIB_DiffServRandomDropTable_DiffServRandomDropEntry
// An entry describes a process that drops packets according to a
// random algorithm.
type DIFFSERVMIB_DiffServRandomDropTable_DiffServRandomDropEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Random Drop entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServRandomDropNextFree. The type is interface{} with range:
    // 1..4294967295.
    DiffServRandomDropId interface{}

    // The average queue depth in bytes, beyond which traffic has a non-zero
    // probability of being dropped. Changes in this variable may or may not be
    // reflected in the reported value of diffServRandomDropMinThreshPkts. The
    // type is interface{} with range: 1..4294967295. Units are bytes.
    DiffServRandomDropMinThreshBytes interface{}

    // The average queue depth in packets, beyond which traffic has a non-zero
    // probability of being dropped. Changes in this variable may or may not be
    // reflected in the reported value of diffServRandomDropMinThreshBytes. The
    // type is interface{} with range: 1..4294967295. Units are packets.
    DiffServRandomDropMinThreshPkts interface{}

    // The average queue depth beyond which traffic has a probability indicated by
    // diffServRandomDropProbMax of being dropped or marked. Note that this
    // differs from the physical queue limit, which is stored in
    // diffServAlgDropQThreshold. Changes in this variable may or may not be
    // reflected in the reported value of diffServRandomDropMaxThreshPkts. The
    // type is interface{} with range: 1..4294967295. Units are bytes.
    DiffServRandomDropMaxThreshBytes interface{}

    // The average queue depth beyond which traffic has a probability indicated by
    // diffServRandomDropProbMax of being dropped or marked. Note that this
    // differs from the physical queue limit, which is stored in
    // diffServAlgDropQThreshold. Changes in this variable may or may not be
    // reflected in the reported value of diffServRandomDropMaxThreshBytes. The
    // type is interface{} with range: 1..4294967295. Units are packets.
    DiffServRandomDropMaxThreshPkts interface{}

    // The worst case random drop probability, expressed in drops per thousand
    // packets.  For example, if in the worst case every arriving packet may be
    // dropped (100%) for a period, this has the value 1000. Alternatively, if in
    // the worst case only one percent (1%) of traffic may be dropped, it has the
    // value 10. The type is interface{} with range: 0..1000.
    DiffServRandomDropProbMax interface{}

    // The weighting of past history in affecting the Exponentially Weighted
    // Moving Average function that calculates the current average queue depth. 
    // The equation uses diffServRandomDropWeight/65536 as the coefficient for the
    // new sample in the equation, and (65536 - diffServRandomDropWeight)/65536 as
    // the coefficient of the old value.  Implementations may limit the values of
    // diffServRandomDropWeight to a subset of the possible range of values, such
    // as powers of two. Doing this would facilitate implementation of the
    // Exponentially Weighted Moving Average using shift instructions or
    // registers. The type is interface{} with range: 0..65536.
    DiffServRandomDropWeight interface{}

    // The number of times per second the queue is sampled for queue average
    // calculation.  A value of zero is used to mean that the queue is sampled
    // approximately each time a packet is enqueued (or dequeued). The type is
    // interface{} with range: 0..1000000.
    DiffServRandomDropSamplingRate interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServRandomDropStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServRandomDropStatus interface{}
}

func (diffServRandomDropEntry *DIFFSERVMIB_DiffServRandomDropTable_DiffServRandomDropEntry) GetEntityData() *types.CommonEntityData {
    diffServRandomDropEntry.EntityData.YFilter = diffServRandomDropEntry.YFilter
    diffServRandomDropEntry.EntityData.YangName = "diffServRandomDropEntry"
    diffServRandomDropEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServRandomDropEntry.EntityData.ParentYangName = "diffServRandomDropTable"
    diffServRandomDropEntry.EntityData.SegmentPath = "diffServRandomDropEntry" + types.AddKeyToken(diffServRandomDropEntry.DiffServRandomDropId, "diffServRandomDropId")
    diffServRandomDropEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServRandomDropTable/" + diffServRandomDropEntry.EntityData.SegmentPath
    diffServRandomDropEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServRandomDropEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServRandomDropEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServRandomDropEntry.EntityData.Children = types.NewOrderedMap()
    diffServRandomDropEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropId", types.YLeaf{"DiffServRandomDropId", diffServRandomDropEntry.DiffServRandomDropId})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropMinThreshBytes", types.YLeaf{"DiffServRandomDropMinThreshBytes", diffServRandomDropEntry.DiffServRandomDropMinThreshBytes})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropMinThreshPkts", types.YLeaf{"DiffServRandomDropMinThreshPkts", diffServRandomDropEntry.DiffServRandomDropMinThreshPkts})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropMaxThreshBytes", types.YLeaf{"DiffServRandomDropMaxThreshBytes", diffServRandomDropEntry.DiffServRandomDropMaxThreshBytes})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropMaxThreshPkts", types.YLeaf{"DiffServRandomDropMaxThreshPkts", diffServRandomDropEntry.DiffServRandomDropMaxThreshPkts})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropProbMax", types.YLeaf{"DiffServRandomDropProbMax", diffServRandomDropEntry.DiffServRandomDropProbMax})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropWeight", types.YLeaf{"DiffServRandomDropWeight", diffServRandomDropEntry.DiffServRandomDropWeight})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropSamplingRate", types.YLeaf{"DiffServRandomDropSamplingRate", diffServRandomDropEntry.DiffServRandomDropSamplingRate})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropStorage", types.YLeaf{"DiffServRandomDropStorage", diffServRandomDropEntry.DiffServRandomDropStorage})
    diffServRandomDropEntry.EntityData.Leafs.Append("diffServRandomDropStatus", types.YLeaf{"DiffServRandomDropStatus", diffServRandomDropEntry.DiffServRandomDropStatus})

    diffServRandomDropEntry.EntityData.YListKeys = []string {"DiffServRandomDropId"}

    return &(diffServRandomDropEntry.EntityData)
}

// DIFFSERVMIB_DiffServQTable
// The Queue Table enumerates the individual queues.  Note that the
// MIB models queuing systems as composed of individual queues, one
// per class of traffic, even though they may in fact be structured
// as classes of traffic scheduled using a common calendar queue, or
// in other ways.
type DIFFSERVMIB_DiffServQTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Queue Table describes a single queue or class of traffic.
    // The type is slice of DIFFSERVMIB_DiffServQTable_DiffServQEntry.
    DiffServQEntry []*DIFFSERVMIB_DiffServQTable_DiffServQEntry
}

func (diffServQTable *DIFFSERVMIB_DiffServQTable) GetEntityData() *types.CommonEntityData {
    diffServQTable.EntityData.YFilter = diffServQTable.YFilter
    diffServQTable.EntityData.YangName = "diffServQTable"
    diffServQTable.EntityData.BundleName = "cisco_ios_xe"
    diffServQTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServQTable.EntityData.SegmentPath = "diffServQTable"
    diffServQTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServQTable.EntityData.SegmentPath
    diffServQTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServQTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServQTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServQTable.EntityData.Children = types.NewOrderedMap()
    diffServQTable.EntityData.Children.Append("diffServQEntry", types.YChild{"DiffServQEntry", nil})
    for i := range diffServQTable.DiffServQEntry {
        diffServQTable.EntityData.Children.Append(types.GetSegmentPath(diffServQTable.DiffServQEntry[i]), types.YChild{"DiffServQEntry", diffServQTable.DiffServQEntry[i]})
    }
    diffServQTable.EntityData.Leafs = types.NewOrderedMap()

    diffServQTable.EntityData.YListKeys = []string {}

    return &(diffServQTable.EntityData)
}

// DIFFSERVMIB_DiffServQTable_DiffServQEntry
// An entry in the Queue Table describes a single queue or class of
// traffic.
type DIFFSERVMIB_DiffServQTable_DiffServQEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Queue entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServQNextFree. The type is interface{} with range: 1..4294967295.
    DiffServQId interface{}

    // This selects the next Differentiated Services Scheduler.  The RowPointer
    // must point to a diffServSchedulerEntry.  A value of zeroDotZero in this
    // attribute indicates an incomplete diffServQEntry instance. In such a case,
    // the entry has no operational effect, since it has no parameters to give it
    // meaning.  Setting this to point to a target that does not exist results in
    // an inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServQNext interface{}

    // This RowPointer indicates the diffServMinRateEntry that the scheduler,
    // pointed to by diffServQNext, should use to service this queue.  If the row
    // pointed to is zeroDotZero, the minimum rate and priority is unspecified. 
    // Setting this to point to a target that does not exist results in an
    // inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServQMinRate interface{}

    // This RowPointer indicates the diffServMaxRateEntry that the scheduler,
    // pointed to by diffServQNext, should use to service this queue.  If the row
    // pointed to is zeroDotZero, the maximum rate is the line speed of the
    // interface.     Setting this to point to a target that does not exist
    // results in an inconsistentValue error.  If the row pointed to is removed or
    // becomes inactive by other means, the treatment is as if this attribute
    // contains a value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServQMaxRate interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServQStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServQStatus interface{}
}

func (diffServQEntry *DIFFSERVMIB_DiffServQTable_DiffServQEntry) GetEntityData() *types.CommonEntityData {
    diffServQEntry.EntityData.YFilter = diffServQEntry.YFilter
    diffServQEntry.EntityData.YangName = "diffServQEntry"
    diffServQEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServQEntry.EntityData.ParentYangName = "diffServQTable"
    diffServQEntry.EntityData.SegmentPath = "diffServQEntry" + types.AddKeyToken(diffServQEntry.DiffServQId, "diffServQId")
    diffServQEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServQTable/" + diffServQEntry.EntityData.SegmentPath
    diffServQEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServQEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServQEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServQEntry.EntityData.Children = types.NewOrderedMap()
    diffServQEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServQEntry.EntityData.Leafs.Append("diffServQId", types.YLeaf{"DiffServQId", diffServQEntry.DiffServQId})
    diffServQEntry.EntityData.Leafs.Append("diffServQNext", types.YLeaf{"DiffServQNext", diffServQEntry.DiffServQNext})
    diffServQEntry.EntityData.Leafs.Append("diffServQMinRate", types.YLeaf{"DiffServQMinRate", diffServQEntry.DiffServQMinRate})
    diffServQEntry.EntityData.Leafs.Append("diffServQMaxRate", types.YLeaf{"DiffServQMaxRate", diffServQEntry.DiffServQMaxRate})
    diffServQEntry.EntityData.Leafs.Append("diffServQStorage", types.YLeaf{"DiffServQStorage", diffServQEntry.DiffServQStorage})
    diffServQEntry.EntityData.Leafs.Append("diffServQStatus", types.YLeaf{"DiffServQStatus", diffServQEntry.DiffServQStatus})

    diffServQEntry.EntityData.YListKeys = []string {"DiffServQId"}

    return &(diffServQEntry.EntityData)
}

// DIFFSERVMIB_DiffServSchedulerTable
// The Scheduler Table enumerates packet schedulers. Multiple
// scheduling algorithms can be used on a given data path, with each
// algorithm described by one diffServSchedulerEntry.
type DIFFSERVMIB_DiffServSchedulerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Scheduler Table describing a single instance of a
    // scheduling algorithm. The type is slice of
    // DIFFSERVMIB_DiffServSchedulerTable_DiffServSchedulerEntry.
    DiffServSchedulerEntry []*DIFFSERVMIB_DiffServSchedulerTable_DiffServSchedulerEntry
}

func (diffServSchedulerTable *DIFFSERVMIB_DiffServSchedulerTable) GetEntityData() *types.CommonEntityData {
    diffServSchedulerTable.EntityData.YFilter = diffServSchedulerTable.YFilter
    diffServSchedulerTable.EntityData.YangName = "diffServSchedulerTable"
    diffServSchedulerTable.EntityData.BundleName = "cisco_ios_xe"
    diffServSchedulerTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServSchedulerTable.EntityData.SegmentPath = "diffServSchedulerTable"
    diffServSchedulerTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServSchedulerTable.EntityData.SegmentPath
    diffServSchedulerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServSchedulerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServSchedulerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServSchedulerTable.EntityData.Children = types.NewOrderedMap()
    diffServSchedulerTable.EntityData.Children.Append("diffServSchedulerEntry", types.YChild{"DiffServSchedulerEntry", nil})
    for i := range diffServSchedulerTable.DiffServSchedulerEntry {
        diffServSchedulerTable.EntityData.Children.Append(types.GetSegmentPath(diffServSchedulerTable.DiffServSchedulerEntry[i]), types.YChild{"DiffServSchedulerEntry", diffServSchedulerTable.DiffServSchedulerEntry[i]})
    }
    diffServSchedulerTable.EntityData.Leafs = types.NewOrderedMap()

    diffServSchedulerTable.EntityData.YListKeys = []string {}

    return &(diffServSchedulerTable.EntityData)
}

// DIFFSERVMIB_DiffServSchedulerTable_DiffServSchedulerEntry
// An entry in the Scheduler Table describing a single instance of
// a scheduling algorithm.
type DIFFSERVMIB_DiffServSchedulerTable_DiffServSchedulerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Scheduler entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServSchedulerNextFree. The type is interface{} with range:
    // 1..4294967295.
    DiffServSchedulerId interface{}

    // This selects the next Differentiated Services Functional Data Path Element
    // to handle traffic for this data path. This normally is null (zeroDotZero),
    // or points to a diffServSchedulerEntry or a diffServQEntry.  However, this
    // RowPointer may also point to an instance of:   diffServClfrEntry,  
    // diffServMeterEntry,   diffServActionEntry,   diffServAlgDropEntry.  It
    // would point another diffServSchedulerEntry when implementing multiple
    // scheduler methods for the same data path, such as having one set of queues
    // scheduled by WRR and that group participating in a priority scheduling
    // system in which other queues compete with it in that way.  It might also
    // point to a second scheduler in a hierarchical scheduling system.  If the
    // row pointed to is zeroDotZero, no further Differentiated Services treatment
    // is performed on traffic of this data path.  Setting this to point to a
    // target that does not exist results in an inconsistentValue error.  If the
    // row pointed to is removed or becomes inactive by other means, the treatment
    // is as if this attribute contains a value of zeroDotZero. The type is string
    // with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServSchedulerNext interface{}

    // The scheduling algorithm used by this Scheduler. zeroDotZero indicates that
    // this is unknown.  Standard values for generic algorithms:
    // diffServSchedulerPriority, diffServSchedulerWRR, and diffServSchedulerWFQ
    // are specified in this MIB; additional values    may be further specified in
    // other MIBs. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServSchedulerMethod interface{}

    // This RowPointer indicates the entry in diffServMinRateTable which indicates
    // the priority or minimum output rate from this scheduler. This attribute is
    // used only when there is more than one level of scheduler.  When it has the
    // value zeroDotZero, it indicates that no minimum rate or priority is
    // imposed.  Setting this to point to a target that does not exist results in
    // an inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServSchedulerMinRate interface{}

    // This RowPointer indicates the entry in diffServMaxRateTable which indicates
    // the maximum output rate from this scheduler. When more than one maximum
    // rate applies (eg, when a multi-rate shaper is in view), it points to the
    // first of those rate entries. This attribute is used only when there is more
    // than one level of scheduler.  When it has the value zeroDotZero, it
    // indicates that no maximum rate is imposed.  Setting this to point to a
    // target that does not exist results in an inconsistentValue error.  If the
    // row pointed to is removed or becomes inactive by other means, the treatment
    // is as if this attribute contains a value of zeroDotZero. The type is string
    // with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    DiffServSchedulerMaxRate interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServSchedulerStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServSchedulerStatus interface{}
}

func (diffServSchedulerEntry *DIFFSERVMIB_DiffServSchedulerTable_DiffServSchedulerEntry) GetEntityData() *types.CommonEntityData {
    diffServSchedulerEntry.EntityData.YFilter = diffServSchedulerEntry.YFilter
    diffServSchedulerEntry.EntityData.YangName = "diffServSchedulerEntry"
    diffServSchedulerEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServSchedulerEntry.EntityData.ParentYangName = "diffServSchedulerTable"
    diffServSchedulerEntry.EntityData.SegmentPath = "diffServSchedulerEntry" + types.AddKeyToken(diffServSchedulerEntry.DiffServSchedulerId, "diffServSchedulerId")
    diffServSchedulerEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServSchedulerTable/" + diffServSchedulerEntry.EntityData.SegmentPath
    diffServSchedulerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServSchedulerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServSchedulerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServSchedulerEntry.EntityData.Children = types.NewOrderedMap()
    diffServSchedulerEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServSchedulerEntry.EntityData.Leafs.Append("diffServSchedulerId", types.YLeaf{"DiffServSchedulerId", diffServSchedulerEntry.DiffServSchedulerId})
    diffServSchedulerEntry.EntityData.Leafs.Append("diffServSchedulerNext", types.YLeaf{"DiffServSchedulerNext", diffServSchedulerEntry.DiffServSchedulerNext})
    diffServSchedulerEntry.EntityData.Leafs.Append("diffServSchedulerMethod", types.YLeaf{"DiffServSchedulerMethod", diffServSchedulerEntry.DiffServSchedulerMethod})
    diffServSchedulerEntry.EntityData.Leafs.Append("diffServSchedulerMinRate", types.YLeaf{"DiffServSchedulerMinRate", diffServSchedulerEntry.DiffServSchedulerMinRate})
    diffServSchedulerEntry.EntityData.Leafs.Append("diffServSchedulerMaxRate", types.YLeaf{"DiffServSchedulerMaxRate", diffServSchedulerEntry.DiffServSchedulerMaxRate})
    diffServSchedulerEntry.EntityData.Leafs.Append("diffServSchedulerStorage", types.YLeaf{"DiffServSchedulerStorage", diffServSchedulerEntry.DiffServSchedulerStorage})
    diffServSchedulerEntry.EntityData.Leafs.Append("diffServSchedulerStatus", types.YLeaf{"DiffServSchedulerStatus", diffServSchedulerEntry.DiffServSchedulerStatus})

    diffServSchedulerEntry.EntityData.YListKeys = []string {"DiffServSchedulerId"}

    return &(diffServSchedulerEntry.EntityData)
}

// DIFFSERVMIB_DiffServMinRateTable
// The Minimum Rate Parameters Table enumerates individual sets of
// scheduling parameter that can be used/reused by Queues and
// Schedulers.
type DIFFSERVMIB_DiffServMinRateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Minimum Rate Parameters Table describes a single set of
    // scheduling parameters for use by one or more queues or schedulers. The type
    // is slice of DIFFSERVMIB_DiffServMinRateTable_DiffServMinRateEntry.
    DiffServMinRateEntry []*DIFFSERVMIB_DiffServMinRateTable_DiffServMinRateEntry
}

func (diffServMinRateTable *DIFFSERVMIB_DiffServMinRateTable) GetEntityData() *types.CommonEntityData {
    diffServMinRateTable.EntityData.YFilter = diffServMinRateTable.YFilter
    diffServMinRateTable.EntityData.YangName = "diffServMinRateTable"
    diffServMinRateTable.EntityData.BundleName = "cisco_ios_xe"
    diffServMinRateTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServMinRateTable.EntityData.SegmentPath = "diffServMinRateTable"
    diffServMinRateTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServMinRateTable.EntityData.SegmentPath
    diffServMinRateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMinRateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMinRateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMinRateTable.EntityData.Children = types.NewOrderedMap()
    diffServMinRateTable.EntityData.Children.Append("diffServMinRateEntry", types.YChild{"DiffServMinRateEntry", nil})
    for i := range diffServMinRateTable.DiffServMinRateEntry {
        diffServMinRateTable.EntityData.Children.Append(types.GetSegmentPath(diffServMinRateTable.DiffServMinRateEntry[i]), types.YChild{"DiffServMinRateEntry", diffServMinRateTable.DiffServMinRateEntry[i]})
    }
    diffServMinRateTable.EntityData.Leafs = types.NewOrderedMap()

    diffServMinRateTable.EntityData.YListKeys = []string {}

    return &(diffServMinRateTable.EntityData)
}

// DIFFSERVMIB_DiffServMinRateTable_DiffServMinRateEntry
// An entry in the Minimum Rate Parameters Table describes a single
// set of scheduling parameters for use by one or more queues or
// schedulers.
type DIFFSERVMIB_DiffServMinRateTable_DiffServMinRateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Scheduler Parameter
    // entries. Managers obtain new values for row creation in this table by
    // reading diffServMinRateNextFree. The type is interface{} with range:
    // 1..4294967295.
    DiffServMinRateId interface{}

    // The priority of this input to the associated scheduler, relative    to the
    // scheduler's other inputs. A queue or scheduler with a larger numeric value
    // will be served before another with a smaller numeric value. The type is
    // interface{} with range: 1..4294967295.
    DiffServMinRatePriority interface{}

    // The minimum absolute rate, in kilobits/sec, that a downstream scheduler
    // element should allocate to this queue. If the value is zero, then there is
    // effectively no minimum rate guarantee. If the value is non-zero, the
    // scheduler will assure the servicing of this queue to at least this rate. 
    // Note that this attribute value and that of diffServMinRateRelative are
    // coupled: changes to one will affect the value of the other. They are linked
    // by the following equation, in that setting one will change the other:   
    // diffServMinRateRelative =          
    // (diffServMinRateAbsolute*1000000)/ifSpeed  or, if appropriate:   
    // diffServMinRateRelative = diffServMinRateAbsolute/ifHighSpeed. The type is
    // interface{} with range: 1..4294967295. Units are kilobits per second.
    DiffServMinRateAbsolute interface{}

    // The minimum rate that a downstream scheduler element should allocate to
    // this queue, relative to the maximum rate of the interface as reported by
    // ifSpeed or ifHighSpeed, in units of 1/1000 of 1. If the value is zero, then
    // there is effectively no minimum rate guarantee. If the value is non-zero,
    // the scheduler will assure the servicing of this queue to at least this
    // rate.  Note that this attribute value and that of diffServMinRateAbsolute
    // are coupled: changes to one will affect the value of the other. They are
    // linked by the following equation, in that setting one will change the
    // other:      diffServMinRateRelative =          
    // (diffServMinRateAbsolute*1000000)/ifSpeed  or, if appropriate:   
    // diffServMinRateRelative = diffServMinRateAbsolute/ifHighSpeed. The type is
    // interface{} with range: 1..4294967295.
    DiffServMinRateRelative interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServMinRateStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServMinRateStatus interface{}
}

func (diffServMinRateEntry *DIFFSERVMIB_DiffServMinRateTable_DiffServMinRateEntry) GetEntityData() *types.CommonEntityData {
    diffServMinRateEntry.EntityData.YFilter = diffServMinRateEntry.YFilter
    diffServMinRateEntry.EntityData.YangName = "diffServMinRateEntry"
    diffServMinRateEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServMinRateEntry.EntityData.ParentYangName = "diffServMinRateTable"
    diffServMinRateEntry.EntityData.SegmentPath = "diffServMinRateEntry" + types.AddKeyToken(diffServMinRateEntry.DiffServMinRateId, "diffServMinRateId")
    diffServMinRateEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServMinRateTable/" + diffServMinRateEntry.EntityData.SegmentPath
    diffServMinRateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMinRateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMinRateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMinRateEntry.EntityData.Children = types.NewOrderedMap()
    diffServMinRateEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServMinRateEntry.EntityData.Leafs.Append("diffServMinRateId", types.YLeaf{"DiffServMinRateId", diffServMinRateEntry.DiffServMinRateId})
    diffServMinRateEntry.EntityData.Leafs.Append("diffServMinRatePriority", types.YLeaf{"DiffServMinRatePriority", diffServMinRateEntry.DiffServMinRatePriority})
    diffServMinRateEntry.EntityData.Leafs.Append("diffServMinRateAbsolute", types.YLeaf{"DiffServMinRateAbsolute", diffServMinRateEntry.DiffServMinRateAbsolute})
    diffServMinRateEntry.EntityData.Leafs.Append("diffServMinRateRelative", types.YLeaf{"DiffServMinRateRelative", diffServMinRateEntry.DiffServMinRateRelative})
    diffServMinRateEntry.EntityData.Leafs.Append("diffServMinRateStorage", types.YLeaf{"DiffServMinRateStorage", diffServMinRateEntry.DiffServMinRateStorage})
    diffServMinRateEntry.EntityData.Leafs.Append("diffServMinRateStatus", types.YLeaf{"DiffServMinRateStatus", diffServMinRateEntry.DiffServMinRateStatus})

    diffServMinRateEntry.EntityData.YListKeys = []string {"DiffServMinRateId"}

    return &(diffServMinRateEntry.EntityData)
}

// DIFFSERVMIB_DiffServMaxRateTable
// The Maximum Rate Parameter Table enumerates individual sets of
// scheduling parameter that can be used/reused by Queues and
// Schedulers.
type DIFFSERVMIB_DiffServMaxRateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Maximum Rate Parameter Table describes a single set of
    // scheduling parameters for use by one or more queues or schedulers. The type
    // is slice of DIFFSERVMIB_DiffServMaxRateTable_DiffServMaxRateEntry.
    DiffServMaxRateEntry []*DIFFSERVMIB_DiffServMaxRateTable_DiffServMaxRateEntry
}

func (diffServMaxRateTable *DIFFSERVMIB_DiffServMaxRateTable) GetEntityData() *types.CommonEntityData {
    diffServMaxRateTable.EntityData.YFilter = diffServMaxRateTable.YFilter
    diffServMaxRateTable.EntityData.YangName = "diffServMaxRateTable"
    diffServMaxRateTable.EntityData.BundleName = "cisco_ios_xe"
    diffServMaxRateTable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffServMaxRateTable.EntityData.SegmentPath = "diffServMaxRateTable"
    diffServMaxRateTable.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/" + diffServMaxRateTable.EntityData.SegmentPath
    diffServMaxRateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMaxRateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMaxRateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMaxRateTable.EntityData.Children = types.NewOrderedMap()
    diffServMaxRateTable.EntityData.Children.Append("diffServMaxRateEntry", types.YChild{"DiffServMaxRateEntry", nil})
    for i := range diffServMaxRateTable.DiffServMaxRateEntry {
        diffServMaxRateTable.EntityData.Children.Append(types.GetSegmentPath(diffServMaxRateTable.DiffServMaxRateEntry[i]), types.YChild{"DiffServMaxRateEntry", diffServMaxRateTable.DiffServMaxRateEntry[i]})
    }
    diffServMaxRateTable.EntityData.Leafs = types.NewOrderedMap()

    diffServMaxRateTable.EntityData.YListKeys = []string {}

    return &(diffServMaxRateTable.EntityData)
}

// DIFFSERVMIB_DiffServMaxRateTable_DiffServMaxRateEntry
// An entry in the Maximum Rate Parameter Table describes a single
// set of scheduling parameters for use by one or more queues or
// schedulers.
type DIFFSERVMIB_DiffServMaxRateTable_DiffServMaxRateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An index that enumerates the Maximum Rate
    // Parameter entries. Managers obtain new values for row creation in this
    // table by reading diffServMaxRateNextFree. The type is interface{} with
    // range: 1..4294967295.
    DiffServMaxRateId interface{}

    // This attribute is a key. An index that indicates which level of a
    // multi-rate shaper is being given its parameters. A multi-rate shaper has
    // some number of rate levels. Frame Relay's dual rate specification refers to
    // a 'committed' and an 'excess' rate; ATM's dual rate specification refers to
    // a 'mean' and a 'peak' rate. This table is generalized to support an
    // arbitrary number of rates. The committed or mean rate is level 1, the peak
    // rate (if any) is the highest level rate configured, and if there are other
    // rates they are distributed in monotonically increasing order between them.
    // The type is interface{} with range: 1..32.
    DiffServMaxRateLevel interface{}

    // The maximum rate in kilobits/sec that a downstream scheduler element should
    // allocate to this queue. If the value is zero, then there is effectively no
    // maximum rate limit and that the scheduler should attempt to be work
    // conserving for this queue. If the value is non-zero, the scheduler will
    // limit the servicing of this queue to, at most, this rate in a
    // non-work-conserving manner.  Note that this attribute value and that of
    // diffServMaxRateRelative are coupled: changes to one will affect the value
    // of the other. They are linked by the following    equation, in that setting
    // one will change the other:    diffServMaxRateRelative =          
    // (diffServMaxRateAbsolute*1000000)/ifSpeed  or, if appropriate:   
    // diffServMaxRateRelative = diffServMaxRateAbsolute/ifHighSpeed. The type is
    // interface{} with range: 1..4294967295. Units are kilobits per second.
    DiffServMaxRateAbsolute interface{}

    // The maximum rate that a downstream scheduler element should allocate to
    // this queue, relative to the maximum rate of the interface as reported by
    // ifSpeed or ifHighSpeed, in units of 1/1000 of 1. If the value is zero, then
    // there is effectively no maximum rate limit and the scheduler should attempt
    // to be work conserving for this queue. If the value is non-zero, the
    // scheduler will limit the servicing of this queue to, at most, this rate in
    // a non-work-conserving manner.  Note that this attribute value and that of
    // diffServMaxRateAbsolute are coupled: changes to one will affect the value
    // of the other. They are linked by the following equation, in that setting
    // one will change the other:    diffServMaxRateRelative =          
    // (diffServMaxRateAbsolute*1000000)/ifSpeed  or, if appropriate:   
    // diffServMaxRateRelative = diffServMaxRateAbsolute/ifHighSpeed. The type is
    // interface{} with range: 1..4294967295.
    DiffServMaxRateRelative interface{}

    // The number of bytes of queue depth at which the rate of a    multi-rate
    // scheduler will increase to the next output rate. In the last conceptual row
    // for such a shaper, this threshold is ignored and by convention is zero. The
    // type is interface{} with range: 0..2147483647. Units are Bytes.
    DiffServMaxRateThreshold interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DiffServMaxRateStorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    DiffServMaxRateStatus interface{}
}

func (diffServMaxRateEntry *DIFFSERVMIB_DiffServMaxRateTable_DiffServMaxRateEntry) GetEntityData() *types.CommonEntityData {
    diffServMaxRateEntry.EntityData.YFilter = diffServMaxRateEntry.YFilter
    diffServMaxRateEntry.EntityData.YangName = "diffServMaxRateEntry"
    diffServMaxRateEntry.EntityData.BundleName = "cisco_ios_xe"
    diffServMaxRateEntry.EntityData.ParentYangName = "diffServMaxRateTable"
    diffServMaxRateEntry.EntityData.SegmentPath = "diffServMaxRateEntry" + types.AddKeyToken(diffServMaxRateEntry.DiffServMaxRateId, "diffServMaxRateId") + types.AddKeyToken(diffServMaxRateEntry.DiffServMaxRateLevel, "diffServMaxRateLevel")
    diffServMaxRateEntry.EntityData.AbsolutePath = "DIFFSERV-MIB:DIFFSERV-MIB/diffServMaxRateTable/" + diffServMaxRateEntry.EntityData.SegmentPath
    diffServMaxRateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffServMaxRateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffServMaxRateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffServMaxRateEntry.EntityData.Children = types.NewOrderedMap()
    diffServMaxRateEntry.EntityData.Leafs = types.NewOrderedMap()
    diffServMaxRateEntry.EntityData.Leafs.Append("diffServMaxRateId", types.YLeaf{"DiffServMaxRateId", diffServMaxRateEntry.DiffServMaxRateId})
    diffServMaxRateEntry.EntityData.Leafs.Append("diffServMaxRateLevel", types.YLeaf{"DiffServMaxRateLevel", diffServMaxRateEntry.DiffServMaxRateLevel})
    diffServMaxRateEntry.EntityData.Leafs.Append("diffServMaxRateAbsolute", types.YLeaf{"DiffServMaxRateAbsolute", diffServMaxRateEntry.DiffServMaxRateAbsolute})
    diffServMaxRateEntry.EntityData.Leafs.Append("diffServMaxRateRelative", types.YLeaf{"DiffServMaxRateRelative", diffServMaxRateEntry.DiffServMaxRateRelative})
    diffServMaxRateEntry.EntityData.Leafs.Append("diffServMaxRateThreshold", types.YLeaf{"DiffServMaxRateThreshold", diffServMaxRateEntry.DiffServMaxRateThreshold})
    diffServMaxRateEntry.EntityData.Leafs.Append("diffServMaxRateStorage", types.YLeaf{"DiffServMaxRateStorage", diffServMaxRateEntry.DiffServMaxRateStorage})
    diffServMaxRateEntry.EntityData.Leafs.Append("diffServMaxRateStatus", types.YLeaf{"DiffServMaxRateStatus", diffServMaxRateEntry.DiffServMaxRateStatus})

    diffServMaxRateEntry.EntityData.YListKeys = []string {"DiffServMaxRateId", "DiffServMaxRateLevel"}

    return &(diffServMaxRateEntry.EntityData)
}

