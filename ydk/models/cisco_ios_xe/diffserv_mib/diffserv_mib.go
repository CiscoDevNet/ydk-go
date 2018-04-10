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

type Diffservtbparamsimpletokenbucket struct {
}

func (id Diffservtbparamsimpletokenbucket) String() string {
	return "DIFFSERV-MIB:diffServTBParamSimpleTokenBucket"
}

type Diffservtbparamavgrate struct {
}

func (id Diffservtbparamavgrate) String() string {
	return "DIFFSERV-MIB:diffServTBParamAvgRate"
}

type Diffservtbparamsrtcmblind struct {
}

func (id Diffservtbparamsrtcmblind) String() string {
	return "DIFFSERV-MIB:diffServTBParamSrTCMBlind"
}

type Diffservtbparamsrtcmaware struct {
}

func (id Diffservtbparamsrtcmaware) String() string {
	return "DIFFSERV-MIB:diffServTBParamSrTCMAware"
}

type Diffservtbparamtrtcmblind struct {
}

func (id Diffservtbparamtrtcmblind) String() string {
	return "DIFFSERV-MIB:diffServTBParamTrTCMBlind"
}

type Diffservtbparamtrtcmaware struct {
}

func (id Diffservtbparamtrtcmaware) String() string {
	return "DIFFSERV-MIB:diffServTBParamTrTCMAware"
}

type Diffservtbparamtswtcm struct {
}

func (id Diffservtbparamtswtcm) String() string {
	return "DIFFSERV-MIB:diffServTBParamTswTCM"
}

type Diffservschedulerpriority struct {
}

func (id Diffservschedulerpriority) String() string {
	return "DIFFSERV-MIB:diffServSchedulerPriority"
}

type Diffservschedulerwrr struct {
}

func (id Diffservschedulerwrr) String() string {
	return "DIFFSERV-MIB:diffServSchedulerWRR"
}

type Diffservschedulerwfq struct {
}

func (id Diffservschedulerwfq) String() string {
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

    
    Diffservclassifier DIFFSERVMIB_Diffservclassifier

    
    Diffservmeter DIFFSERVMIB_Diffservmeter

    
    Diffservtbparam DIFFSERVMIB_Diffservtbparam

    
    Diffservaction DIFFSERVMIB_Diffservaction

    
    Diffservalgdrop DIFFSERVMIB_Diffservalgdrop

    
    Diffservqueue DIFFSERVMIB_Diffservqueue

    
    Diffservscheduler DIFFSERVMIB_Diffservscheduler

    // The data path table contains RowPointers indicating the start of the
    // functional data path for each interface and traffic direction in this
    // device. These may merge, or be separated into parallel data paths.
    Diffservdatapathtable DIFFSERVMIB_Diffservdatapathtable

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
    Diffservclfrtable DIFFSERVMIB_Diffservclfrtable

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
    Diffservclfrelementtable DIFFSERVMIB_Diffservclfrelementtable

    // A table of IP Multi-field Classifier filter entries that a    system may
    // use to identify IP traffic.
    Diffservmultifieldclfrtable DIFFSERVMIB_Diffservmultifieldclfrtable

    // This table enumerates specific meters that a system may use to police a
    // stream of traffic. The traffic stream to be metered is determined by the
    // Differentiated Services Functional Data Path Element(s) upstream of the
    // meter i.e. by the object(s) that point to each entry in this table.  This
    // may include all traffic on an interface.  Specific meter details are to be
    // found in table entry referenced by diffServMeterSpecific.
    Diffservmetertable DIFFSERVMIB_Diffservmetertable

    // This table enumerates a single set of token bucket meter parameters that a
    // system may use to police a stream of traffic. Such meters are modeled here
    // as having a single rate and a single burst size. Multiple entries are used
    // when multiple rates/burst sizes are needed.
    Diffservtbparamtable DIFFSERVMIB_Diffservtbparamtable

    // The Action Table enumerates actions that can be performed to a stream of
    // traffic. Multiple actions can be concatenated. For example, traffic exiting
    // from a meter may be counted, marked, and potentially dropped before
    // entering a queue.  Specific actions are indicated by diffServActionSpecific
    // which points to an entry of a specific action type parameterizing the
    // action in detail.
    Diffservactiontable DIFFSERVMIB_Diffservactiontable

    // This table enumerates specific DSCPs used for marking or remarking the DSCP
    // field of IP packets. The entries of this table may be referenced by a
    // diffServActionSpecific attribute.
    Diffservdscpmarkacttable DIFFSERVMIB_Diffservdscpmarkacttable

    // This table contains counters for all the traffic passing through an action
    // element.
    Diffservcountacttable DIFFSERVMIB_Diffservcountacttable

    // The algorithmic drop table contains entries describing an element that
    // drops packets according to some algorithm.
    Diffservalgdroptable DIFFSERVMIB_Diffservalgdroptable

    // The random drop table contains entries describing a process that drops
    // packets randomly. Entries in this table are pointed to by
    // diffServAlgDropSpecific.
    Diffservrandomdroptable DIFFSERVMIB_Diffservrandomdroptable

    // The Queue Table enumerates the individual queues.  Note that the MIB models
    // queuing systems as composed of individual queues, one per class of traffic,
    // even though they may in fact be structured as classes of traffic scheduled
    // using a common calendar queue, or in other ways.
    Diffservqtable DIFFSERVMIB_Diffservqtable

    // The Scheduler Table enumerates packet schedulers. Multiple scheduling
    // algorithms can be used on a given data path, with each algorithm described
    // by one diffServSchedulerEntry.
    Diffservschedulertable DIFFSERVMIB_Diffservschedulertable

    // The Minimum Rate Parameters Table enumerates individual sets of scheduling
    // parameter that can be used/reused by Queues and Schedulers.
    Diffservminratetable DIFFSERVMIB_Diffservminratetable

    // The Maximum Rate Parameter Table enumerates individual sets of scheduling
    // parameter that can be used/reused by Queues and Schedulers.
    Diffservmaxratetable DIFFSERVMIB_Diffservmaxratetable
}

func (dIFFSERVMIB *DIFFSERVMIB) GetEntityData() *types.CommonEntityData {
    dIFFSERVMIB.EntityData.YFilter = dIFFSERVMIB.YFilter
    dIFFSERVMIB.EntityData.YangName = "DIFFSERV-MIB"
    dIFFSERVMIB.EntityData.BundleName = "cisco_ios_xe"
    dIFFSERVMIB.EntityData.ParentYangName = "DIFFSERV-MIB"
    dIFFSERVMIB.EntityData.SegmentPath = "DIFFSERV-MIB:DIFFSERV-MIB"
    dIFFSERVMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dIFFSERVMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dIFFSERVMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dIFFSERVMIB.EntityData.Children = make(map[string]types.YChild)
    dIFFSERVMIB.EntityData.Children["diffServClassifier"] = types.YChild{"Diffservclassifier", &dIFFSERVMIB.Diffservclassifier}
    dIFFSERVMIB.EntityData.Children["diffServMeter"] = types.YChild{"Diffservmeter", &dIFFSERVMIB.Diffservmeter}
    dIFFSERVMIB.EntityData.Children["diffServTBParam"] = types.YChild{"Diffservtbparam", &dIFFSERVMIB.Diffservtbparam}
    dIFFSERVMIB.EntityData.Children["diffServAction"] = types.YChild{"Diffservaction", &dIFFSERVMIB.Diffservaction}
    dIFFSERVMIB.EntityData.Children["diffServAlgDrop"] = types.YChild{"Diffservalgdrop", &dIFFSERVMIB.Diffservalgdrop}
    dIFFSERVMIB.EntityData.Children["diffServQueue"] = types.YChild{"Diffservqueue", &dIFFSERVMIB.Diffservqueue}
    dIFFSERVMIB.EntityData.Children["diffServScheduler"] = types.YChild{"Diffservscheduler", &dIFFSERVMIB.Diffservscheduler}
    dIFFSERVMIB.EntityData.Children["diffServDataPathTable"] = types.YChild{"Diffservdatapathtable", &dIFFSERVMIB.Diffservdatapathtable}
    dIFFSERVMIB.EntityData.Children["diffServClfrTable"] = types.YChild{"Diffservclfrtable", &dIFFSERVMIB.Diffservclfrtable}
    dIFFSERVMIB.EntityData.Children["diffServClfrElementTable"] = types.YChild{"Diffservclfrelementtable", &dIFFSERVMIB.Diffservclfrelementtable}
    dIFFSERVMIB.EntityData.Children["diffServMultiFieldClfrTable"] = types.YChild{"Diffservmultifieldclfrtable", &dIFFSERVMIB.Diffservmultifieldclfrtable}
    dIFFSERVMIB.EntityData.Children["diffServMeterTable"] = types.YChild{"Diffservmetertable", &dIFFSERVMIB.Diffservmetertable}
    dIFFSERVMIB.EntityData.Children["diffServTBParamTable"] = types.YChild{"Diffservtbparamtable", &dIFFSERVMIB.Diffservtbparamtable}
    dIFFSERVMIB.EntityData.Children["diffServActionTable"] = types.YChild{"Diffservactiontable", &dIFFSERVMIB.Diffservactiontable}
    dIFFSERVMIB.EntityData.Children["diffServDscpMarkActTable"] = types.YChild{"Diffservdscpmarkacttable", &dIFFSERVMIB.Diffservdscpmarkacttable}
    dIFFSERVMIB.EntityData.Children["diffServCountActTable"] = types.YChild{"Diffservcountacttable", &dIFFSERVMIB.Diffservcountacttable}
    dIFFSERVMIB.EntityData.Children["diffServAlgDropTable"] = types.YChild{"Diffservalgdroptable", &dIFFSERVMIB.Diffservalgdroptable}
    dIFFSERVMIB.EntityData.Children["diffServRandomDropTable"] = types.YChild{"Diffservrandomdroptable", &dIFFSERVMIB.Diffservrandomdroptable}
    dIFFSERVMIB.EntityData.Children["diffServQTable"] = types.YChild{"Diffservqtable", &dIFFSERVMIB.Diffservqtable}
    dIFFSERVMIB.EntityData.Children["diffServSchedulerTable"] = types.YChild{"Diffservschedulertable", &dIFFSERVMIB.Diffservschedulertable}
    dIFFSERVMIB.EntityData.Children["diffServMinRateTable"] = types.YChild{"Diffservminratetable", &dIFFSERVMIB.Diffservminratetable}
    dIFFSERVMIB.EntityData.Children["diffServMaxRateTable"] = types.YChild{"Diffservmaxratetable", &dIFFSERVMIB.Diffservmaxratetable}
    dIFFSERVMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dIFFSERVMIB.EntityData)
}

// DIFFSERVMIB_Diffservclassifier
type DIFFSERVMIB_Diffservclassifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServClfrId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservclfrnextfree interface{}

    // This object contains an unused value for diffServClfrElementId, or a zero
    // to indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservclfrelementnextfree interface{}

    // This object contains an unused value for diffServMultiFieldClfrId, or a
    // zero to indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservmultifieldclfrnextfree interface{}
}

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetEntityData() *types.CommonEntityData {
    diffservclassifier.EntityData.YFilter = diffservclassifier.YFilter
    diffservclassifier.EntityData.YangName = "diffServClassifier"
    diffservclassifier.EntityData.BundleName = "cisco_ios_xe"
    diffservclassifier.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservclassifier.EntityData.SegmentPath = "diffServClassifier"
    diffservclassifier.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservclassifier.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservclassifier.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservclassifier.EntityData.Children = make(map[string]types.YChild)
    diffservclassifier.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservclassifier.EntityData.Leafs["diffServClfrNextFree"] = types.YLeaf{"Diffservclfrnextfree", diffservclassifier.Diffservclfrnextfree}
    diffservclassifier.EntityData.Leafs["diffServClfrElementNextFree"] = types.YLeaf{"Diffservclfrelementnextfree", diffservclassifier.Diffservclfrelementnextfree}
    diffservclassifier.EntityData.Leafs["diffServMultiFieldClfrNextFree"] = types.YLeaf{"Diffservmultifieldclfrnextfree", diffservclassifier.Diffservmultifieldclfrnextfree}
    return &(diffservclassifier.EntityData)
}

// DIFFSERVMIB_Diffservmeter
type DIFFSERVMIB_Diffservmeter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServMeterId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservmeternextfree interface{}
}

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetEntityData() *types.CommonEntityData {
    diffservmeter.EntityData.YFilter = diffservmeter.YFilter
    diffservmeter.EntityData.YangName = "diffServMeter"
    diffservmeter.EntityData.BundleName = "cisco_ios_xe"
    diffservmeter.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservmeter.EntityData.SegmentPath = "diffServMeter"
    diffservmeter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservmeter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservmeter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservmeter.EntityData.Children = make(map[string]types.YChild)
    diffservmeter.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservmeter.EntityData.Leafs["diffServMeterNextFree"] = types.YLeaf{"Diffservmeternextfree", diffservmeter.Diffservmeternextfree}
    return &(diffservmeter.EntityData)
}

// DIFFSERVMIB_Diffservtbparam
type DIFFSERVMIB_Diffservtbparam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServTBParamId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservtbparamnextfree interface{}
}

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetEntityData() *types.CommonEntityData {
    diffservtbparam.EntityData.YFilter = diffservtbparam.YFilter
    diffservtbparam.EntityData.YangName = "diffServTBParam"
    diffservtbparam.EntityData.BundleName = "cisco_ios_xe"
    diffservtbparam.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservtbparam.EntityData.SegmentPath = "diffServTBParam"
    diffservtbparam.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservtbparam.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservtbparam.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservtbparam.EntityData.Children = make(map[string]types.YChild)
    diffservtbparam.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservtbparam.EntityData.Leafs["diffServTBParamNextFree"] = types.YLeaf{"Diffservtbparamnextfree", diffservtbparam.Diffservtbparamnextfree}
    return &(diffservtbparam.EntityData)
}

// DIFFSERVMIB_Diffservaction
type DIFFSERVMIB_Diffservaction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServActionId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservactionnextfree interface{}

    // This object contains an unused value for diffServCountActId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservcountactnextfree interface{}
}

func (diffservaction *DIFFSERVMIB_Diffservaction) GetEntityData() *types.CommonEntityData {
    diffservaction.EntityData.YFilter = diffservaction.YFilter
    diffservaction.EntityData.YangName = "diffServAction"
    diffservaction.EntityData.BundleName = "cisco_ios_xe"
    diffservaction.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservaction.EntityData.SegmentPath = "diffServAction"
    diffservaction.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservaction.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservaction.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservaction.EntityData.Children = make(map[string]types.YChild)
    diffservaction.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservaction.EntityData.Leafs["diffServActionNextFree"] = types.YLeaf{"Diffservactionnextfree", diffservaction.Diffservactionnextfree}
    diffservaction.EntityData.Leafs["diffServCountActNextFree"] = types.YLeaf{"Diffservcountactnextfree", diffservaction.Diffservcountactnextfree}
    return &(diffservaction.EntityData)
}

// DIFFSERVMIB_Diffservalgdrop
type DIFFSERVMIB_Diffservalgdrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServAlgDropId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservalgdropnextfree interface{}

    // This object contains an unused value for diffServRandomDropId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservrandomdropnextfree interface{}
}

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetEntityData() *types.CommonEntityData {
    diffservalgdrop.EntityData.YFilter = diffservalgdrop.YFilter
    diffservalgdrop.EntityData.YangName = "diffServAlgDrop"
    diffservalgdrop.EntityData.BundleName = "cisco_ios_xe"
    diffservalgdrop.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservalgdrop.EntityData.SegmentPath = "diffServAlgDrop"
    diffservalgdrop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservalgdrop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservalgdrop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservalgdrop.EntityData.Children = make(map[string]types.YChild)
    diffservalgdrop.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservalgdrop.EntityData.Leafs["diffServAlgDropNextFree"] = types.YLeaf{"Diffservalgdropnextfree", diffservalgdrop.Diffservalgdropnextfree}
    diffservalgdrop.EntityData.Leafs["diffServRandomDropNextFree"] = types.YLeaf{"Diffservrandomdropnextfree", diffservalgdrop.Diffservrandomdropnextfree}
    return &(diffservalgdrop.EntityData)
}

// DIFFSERVMIB_Diffservqueue
type DIFFSERVMIB_Diffservqueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServQId, or a zero to indicate
    // that none exist. The type is interface{} with range: 0..4294967295.
    Diffservqnextfree interface{}
}

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetEntityData() *types.CommonEntityData {
    diffservqueue.EntityData.YFilter = diffservqueue.YFilter
    diffservqueue.EntityData.YangName = "diffServQueue"
    diffservqueue.EntityData.BundleName = "cisco_ios_xe"
    diffservqueue.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservqueue.EntityData.SegmentPath = "diffServQueue"
    diffservqueue.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservqueue.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservqueue.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservqueue.EntityData.Children = make(map[string]types.YChild)
    diffservqueue.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservqueue.EntityData.Leafs["diffServQNextFree"] = types.YLeaf{"Diffservqnextfree", diffservqueue.Diffservqnextfree}
    return &(diffservqueue.EntityData)
}

// DIFFSERVMIB_Diffservscheduler
type DIFFSERVMIB_Diffservscheduler struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServSchedulerId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservschedulernextfree interface{}

    // This object contains an unused value for diffServMinRateId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservminratenextfree interface{}

    // This object contains an unused value for diffServMaxRateId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservmaxratenextfree interface{}
}

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetEntityData() *types.CommonEntityData {
    diffservscheduler.EntityData.YFilter = diffservscheduler.YFilter
    diffservscheduler.EntityData.YangName = "diffServScheduler"
    diffservscheduler.EntityData.BundleName = "cisco_ios_xe"
    diffservscheduler.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservscheduler.EntityData.SegmentPath = "diffServScheduler"
    diffservscheduler.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservscheduler.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservscheduler.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservscheduler.EntityData.Children = make(map[string]types.YChild)
    diffservscheduler.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservscheduler.EntityData.Leafs["diffServSchedulerNextFree"] = types.YLeaf{"Diffservschedulernextfree", diffservscheduler.Diffservschedulernextfree}
    diffservscheduler.EntityData.Leafs["diffServMinRateNextFree"] = types.YLeaf{"Diffservminratenextfree", diffservscheduler.Diffservminratenextfree}
    diffservscheduler.EntityData.Leafs["diffServMaxRateNextFree"] = types.YLeaf{"Diffservmaxratenextfree", diffservscheduler.Diffservmaxratenextfree}
    return &(diffservscheduler.EntityData)
}

// DIFFSERVMIB_Diffservdatapathtable
// The data path table contains RowPointers indicating the start of
// the functional data path for each interface and traffic direction
// in this device. These may merge, or be separated into parallel
// data paths.
type DIFFSERVMIB_Diffservdatapathtable struct {
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
    // type is slice of DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry.
    Diffservdatapathentry []DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry
}

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetEntityData() *types.CommonEntityData {
    diffservdatapathtable.EntityData.YFilter = diffservdatapathtable.YFilter
    diffservdatapathtable.EntityData.YangName = "diffServDataPathTable"
    diffservdatapathtable.EntityData.BundleName = "cisco_ios_xe"
    diffservdatapathtable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservdatapathtable.EntityData.SegmentPath = "diffServDataPathTable"
    diffservdatapathtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservdatapathtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservdatapathtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservdatapathtable.EntityData.Children = make(map[string]types.YChild)
    diffservdatapathtable.EntityData.Children["diffServDataPathEntry"] = types.YChild{"Diffservdatapathentry", nil}
    for i := range diffservdatapathtable.Diffservdatapathentry {
        diffservdatapathtable.EntityData.Children[types.GetSegmentPath(&diffservdatapathtable.Diffservdatapathentry[i])] = types.YChild{"Diffservdatapathentry", &diffservdatapathtable.Diffservdatapathentry[i]}
    }
    diffservdatapathtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservdatapathtable.EntityData)
}

// DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry
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
type DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. IfDirection specifies whether the reception or
    // transmission path for this interface is in view. The type is IfDirection.
    Diffservdatapathifdirection interface{}

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
    Diffservdatapathstart interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservdatapathstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. The type is RowStatus.
    Diffservdatapathstatus interface{}
}

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetEntityData() *types.CommonEntityData {
    diffservdatapathentry.EntityData.YFilter = diffservdatapathentry.YFilter
    diffservdatapathentry.EntityData.YangName = "diffServDataPathEntry"
    diffservdatapathentry.EntityData.BundleName = "cisco_ios_xe"
    diffservdatapathentry.EntityData.ParentYangName = "diffServDataPathTable"
    diffservdatapathentry.EntityData.SegmentPath = "diffServDataPathEntry" + "[ifIndex='" + fmt.Sprintf("%v", diffservdatapathentry.Ifindex) + "']" + "[diffServDataPathIfDirection='" + fmt.Sprintf("%v", diffservdatapathentry.Diffservdatapathifdirection) + "']"
    diffservdatapathentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservdatapathentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservdatapathentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservdatapathentry.EntityData.Children = make(map[string]types.YChild)
    diffservdatapathentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservdatapathentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", diffservdatapathentry.Ifindex}
    diffservdatapathentry.EntityData.Leafs["diffServDataPathIfDirection"] = types.YLeaf{"Diffservdatapathifdirection", diffservdatapathentry.Diffservdatapathifdirection}
    diffservdatapathentry.EntityData.Leafs["diffServDataPathStart"] = types.YLeaf{"Diffservdatapathstart", diffservdatapathentry.Diffservdatapathstart}
    diffservdatapathentry.EntityData.Leafs["diffServDataPathStorage"] = types.YLeaf{"Diffservdatapathstorage", diffservdatapathentry.Diffservdatapathstorage}
    diffservdatapathentry.EntityData.Leafs["diffServDataPathStatus"] = types.YLeaf{"Diffservdatapathstatus", diffservdatapathentry.Diffservdatapathstatus}
    return &(diffservdatapathentry.EntityData)
}

// DIFFSERVMIB_Diffservclfrtable
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
type DIFFSERVMIB_Diffservclfrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the classifier table describes a single classifier. All
    // classifier elements belonging to the same classifier use the classifier's
    // diffServClfrId as part of their index. The type is slice of
    // DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry.
    Diffservclfrentry []DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry
}

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetEntityData() *types.CommonEntityData {
    diffservclfrtable.EntityData.YFilter = diffservclfrtable.YFilter
    diffservclfrtable.EntityData.YangName = "diffServClfrTable"
    diffservclfrtable.EntityData.BundleName = "cisco_ios_xe"
    diffservclfrtable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservclfrtable.EntityData.SegmentPath = "diffServClfrTable"
    diffservclfrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservclfrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservclfrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservclfrtable.EntityData.Children = make(map[string]types.YChild)
    diffservclfrtable.EntityData.Children["diffServClfrEntry"] = types.YChild{"Diffservclfrentry", nil}
    for i := range diffservclfrtable.Diffservclfrentry {
        diffservclfrtable.EntityData.Children[types.GetSegmentPath(&diffservclfrtable.Diffservclfrentry[i])] = types.YChild{"Diffservclfrentry", &diffservclfrtable.Diffservclfrentry[i]}
    }
    diffservclfrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservclfrtable.EntityData)
}

// DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry
// An entry in the classifier table describes a single classifier.
// All classifier elements belonging to the same classifier use the
// classifier's diffServClfrId as part of their index.
type DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the classifier entries. 
    // Managers should obtain new values for row creation in this table by reading
    // diffServClfrNextFree. The type is interface{} with range: 1..4294967295.
    Diffservclfrid interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservclfrstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservclfrstatus interface{}
}

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetEntityData() *types.CommonEntityData {
    diffservclfrentry.EntityData.YFilter = diffservclfrentry.YFilter
    diffservclfrentry.EntityData.YangName = "diffServClfrEntry"
    diffservclfrentry.EntityData.BundleName = "cisco_ios_xe"
    diffservclfrentry.EntityData.ParentYangName = "diffServClfrTable"
    diffservclfrentry.EntityData.SegmentPath = "diffServClfrEntry" + "[diffServClfrId='" + fmt.Sprintf("%v", diffservclfrentry.Diffservclfrid) + "']"
    diffservclfrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservclfrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservclfrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservclfrentry.EntityData.Children = make(map[string]types.YChild)
    diffservclfrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservclfrentry.EntityData.Leafs["diffServClfrId"] = types.YLeaf{"Diffservclfrid", diffservclfrentry.Diffservclfrid}
    diffservclfrentry.EntityData.Leafs["diffServClfrStorage"] = types.YLeaf{"Diffservclfrstorage", diffservclfrentry.Diffservclfrstorage}
    diffservclfrentry.EntityData.Leafs["diffServClfrStatus"] = types.YLeaf{"Diffservclfrstatus", diffservclfrentry.Diffservclfrstatus}
    return &(diffservclfrentry.EntityData)
}

// DIFFSERVMIB_Diffservclfrelementtable
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
type DIFFSERVMIB_Diffservclfrelementtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the classifier element table describes a single element of the
    // classifier. The type is slice of
    // DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry.
    Diffservclfrelemententry []DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry
}

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetEntityData() *types.CommonEntityData {
    diffservclfrelementtable.EntityData.YFilter = diffservclfrelementtable.YFilter
    diffservclfrelementtable.EntityData.YangName = "diffServClfrElementTable"
    diffservclfrelementtable.EntityData.BundleName = "cisco_ios_xe"
    diffservclfrelementtable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservclfrelementtable.EntityData.SegmentPath = "diffServClfrElementTable"
    diffservclfrelementtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservclfrelementtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservclfrelementtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservclfrelementtable.EntityData.Children = make(map[string]types.YChild)
    diffservclfrelementtable.EntityData.Children["diffServClfrElementEntry"] = types.YChild{"Diffservclfrelemententry", nil}
    for i := range diffservclfrelementtable.Diffservclfrelemententry {
        diffservclfrelementtable.EntityData.Children[types.GetSegmentPath(&diffservclfrelementtable.Diffservclfrelemententry[i])] = types.YChild{"Diffservclfrelemententry", &diffservclfrelementtable.Diffservclfrelemententry[i]}
    }
    diffservclfrelementtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservclfrelementtable.EntityData)
}

// DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry
// An entry in the classifier element table describes a single
// element of the classifier.
type DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // diffserv_mib.DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry_Diffservclfrid
    Diffservclfrid interface{}

    // This attribute is a key. An index that enumerates the Classifier Element
    // entries. Managers obtain new values for row creation in this table by
    // reading diffServClfrElementNextFree. The type is interface{} with range:
    // 1..4294967295.
    Diffservclfrelementid interface{}

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
    Diffservclfrelementprecedence interface{}

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
    Diffservclfrelementnext interface{}

    // A pointer to a valid entry in another table, filter table, that describes
    // the applicable classification parameters, e.g. an entry in
    // diffServMultiFieldClfrTable.  The value zeroDotZero is interpreted to match
    // anything not matched by another classifier element - only one such entry
    // may exist for each classifier.  Setting this to point to a target that does
    // not exist results in an inconsistentValue error.  If the row pointed to is
    // removed or    becomes inactive by other means, the element is ignored. The
    // type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservclfrelementspecific interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservclfrelementstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservclfrelementstatus interface{}
}

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetEntityData() *types.CommonEntityData {
    diffservclfrelemententry.EntityData.YFilter = diffservclfrelemententry.YFilter
    diffservclfrelemententry.EntityData.YangName = "diffServClfrElementEntry"
    diffservclfrelemententry.EntityData.BundleName = "cisco_ios_xe"
    diffservclfrelemententry.EntityData.ParentYangName = "diffServClfrElementTable"
    diffservclfrelemententry.EntityData.SegmentPath = "diffServClfrElementEntry" + "[diffServClfrId='" + fmt.Sprintf("%v", diffservclfrelemententry.Diffservclfrid) + "']" + "[diffServClfrElementId='" + fmt.Sprintf("%v", diffservclfrelemententry.Diffservclfrelementid) + "']"
    diffservclfrelemententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservclfrelemententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservclfrelemententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservclfrelemententry.EntityData.Children = make(map[string]types.YChild)
    diffservclfrelemententry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservclfrelemententry.EntityData.Leafs["diffServClfrId"] = types.YLeaf{"Diffservclfrid", diffservclfrelemententry.Diffservclfrid}
    diffservclfrelemententry.EntityData.Leafs["diffServClfrElementId"] = types.YLeaf{"Diffservclfrelementid", diffservclfrelemententry.Diffservclfrelementid}
    diffservclfrelemententry.EntityData.Leafs["diffServClfrElementPrecedence"] = types.YLeaf{"Diffservclfrelementprecedence", diffservclfrelemententry.Diffservclfrelementprecedence}
    diffservclfrelemententry.EntityData.Leafs["diffServClfrElementNext"] = types.YLeaf{"Diffservclfrelementnext", diffservclfrelemententry.Diffservclfrelementnext}
    diffservclfrelemententry.EntityData.Leafs["diffServClfrElementSpecific"] = types.YLeaf{"Diffservclfrelementspecific", diffservclfrelemententry.Diffservclfrelementspecific}
    diffservclfrelemententry.EntityData.Leafs["diffServClfrElementStorage"] = types.YLeaf{"Diffservclfrelementstorage", diffservclfrelemententry.Diffservclfrelementstorage}
    diffservclfrelemententry.EntityData.Leafs["diffServClfrElementStatus"] = types.YLeaf{"Diffservclfrelementstatus", diffservclfrelemententry.Diffservclfrelementstatus}
    return &(diffservclfrelemententry.EntityData)
}

// DIFFSERVMIB_Diffservmultifieldclfrtable
// A table of IP Multi-field Classifier filter entries that a
// 
// 
// 
// system may use to identify IP traffic.
type DIFFSERVMIB_Diffservmultifieldclfrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An IP Multi-field Classifier entry describes a single filter. The type is
    // slice of
    // DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry.
    Diffservmultifieldclfrentry []DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry
}

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetEntityData() *types.CommonEntityData {
    diffservmultifieldclfrtable.EntityData.YFilter = diffservmultifieldclfrtable.YFilter
    diffservmultifieldclfrtable.EntityData.YangName = "diffServMultiFieldClfrTable"
    diffservmultifieldclfrtable.EntityData.BundleName = "cisco_ios_xe"
    diffservmultifieldclfrtable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservmultifieldclfrtable.EntityData.SegmentPath = "diffServMultiFieldClfrTable"
    diffservmultifieldclfrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservmultifieldclfrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservmultifieldclfrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservmultifieldclfrtable.EntityData.Children = make(map[string]types.YChild)
    diffservmultifieldclfrtable.EntityData.Children["diffServMultiFieldClfrEntry"] = types.YChild{"Diffservmultifieldclfrentry", nil}
    for i := range diffservmultifieldclfrtable.Diffservmultifieldclfrentry {
        diffservmultifieldclfrtable.EntityData.Children[types.GetSegmentPath(&diffservmultifieldclfrtable.Diffservmultifieldclfrentry[i])] = types.YChild{"Diffservmultifieldclfrentry", &diffservmultifieldclfrtable.Diffservmultifieldclfrentry[i]}
    }
    diffservmultifieldclfrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservmultifieldclfrtable.EntityData)
}

// DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry
// An IP Multi-field Classifier entry describes a single filter.
type DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the MultiField Classifier
    // filter entries.  Managers obtain new values for row creation in this table
    // by reading diffServMultiFieldClfrNextFree. The type is interface{} with
    // range: 1..4294967295.
    Diffservmultifieldclfrid interface{}

    // The type of IP address used by this classifier entry.  While other types of
    // addresses are defined in the InetAddressType    textual convention, and DNS
    // names, a classifier can only look at packets on the wire. Therefore, this
    // object is limited to IPv4 and IPv6 addresses. The type is InetAddressType.
    Diffservmultifieldclfraddrtype interface{}

    // The IP address to match against the packet's destination IP address. This
    // may not be a DNS name, but may be an IPv4 or IPv6 prefix. 
    // diffServMultiFieldClfrDstPrefixLength indicates the number of bits that are
    // relevant. The type is string with length: 0..255.
    Diffservmultifieldclfrdstaddr interface{}

    // The length of the CIDR Prefix carried in diffServMultiFieldClfrDstAddr. In
    // IPv4 addresses, a length of 0 indicates a match of any address; a length of
    // 32 indicates a match of a single host address, and a length between 0 and
    // 32 indicates the use of a CIDR Prefix. IPv6 is similar, except that prefix
    // lengths range from 0..128. The type is interface{} with range: 0..2040.
    // Units are bits.
    Diffservmultifieldclfrdstprefixlength interface{}

    // The IP address to match against the packet's source IP address. This may
    // not be a DNS name, but may be an IPv4 or IPv6 prefix.
    // diffServMultiFieldClfrSrcPrefixLength indicates the number of bits that are
    // relevant. The type is string with length: 0..255.
    Diffservmultifieldclfrsrcaddr interface{}

    // The length of the CIDR Prefix carried in diffServMultiFieldClfrSrcAddr. In
    // IPv4 addresses, a length of 0 indicates a match of any address; a length of
    // 32 indicates a match of a single host address, and a length between 0 and
    // 32 indicates the use of a CIDR Prefix. IPv6 is similar, except that prefix
    // lengths range from 0..128. The type is interface{} with range: 0..2040.
    // Units are bits.
    Diffservmultifieldclfrsrcprefixlength interface{}

    // The value that the DSCP in the packet must have to match this entry. A
    // value of -1 indicates that a specific DSCP value has not been defined and
    // thus all DSCP values are considered a match. The type is interface{} with
    // range: -1..63.
    Diffservmultifieldclfrdscp interface{}

    // The flow identifier in an IPv6 header. The type is interface{} with range:
    // 0..1048575.
    Diffservmultifieldclfrflowid interface{}

    // The IP protocol to match against the IPv4 protocol number or the IPv6 Next-
    // Header number in the packet. A value of 255 means match all.  Note the
    // protocol number of 255 is reserved by IANA, and Next-Header number of 0 is
    // used in IPv6. The type is interface{} with range: 0..255.
    Diffservmultifieldclfrprotocol interface{}

    // The minimum value that the layer-4 destination port number in the packet
    // must have in order to match this classifier entry. The type is interface{}
    // with range: 0..65535.
    Diffservmultifieldclfrdstl4Portmin interface{}

    // The maximum value that the layer-4 destination port number in the packet
    // must have in order to match this classifier entry. This value must be equal
    // to or greater than the value specified for this entry in
    // diffServMultiFieldClfrDstL4PortMin. The type is interface{} with range:
    // 0..65535.
    Diffservmultifieldclfrdstl4Portmax interface{}

    // The minimum value that the layer-4 source port number in the packet must
    // have in order to match this classifier entry. The type is interface{} with
    // range: 0..65535.
    Diffservmultifieldclfrsrcl4Portmin interface{}

    // The maximum value that the layer-4 source port number in the packet must
    // have in order to match this classifier entry. This value must be equal to
    // or greater than the value specified for this entry in
    // diffServMultiFieldClfrSrcL4PortMin. The type is interface{} with range:
    // 0..65535.
    Diffservmultifieldclfrsrcl4Portmax interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservmultifieldclfrstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservmultifieldclfrstatus interface{}
}

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetEntityData() *types.CommonEntityData {
    diffservmultifieldclfrentry.EntityData.YFilter = diffservmultifieldclfrentry.YFilter
    diffservmultifieldclfrentry.EntityData.YangName = "diffServMultiFieldClfrEntry"
    diffservmultifieldclfrentry.EntityData.BundleName = "cisco_ios_xe"
    diffservmultifieldclfrentry.EntityData.ParentYangName = "diffServMultiFieldClfrTable"
    diffservmultifieldclfrentry.EntityData.SegmentPath = "diffServMultiFieldClfrEntry" + "[diffServMultiFieldClfrId='" + fmt.Sprintf("%v", diffservmultifieldclfrentry.Diffservmultifieldclfrid) + "']"
    diffservmultifieldclfrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservmultifieldclfrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservmultifieldclfrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservmultifieldclfrentry.EntityData.Children = make(map[string]types.YChild)
    diffservmultifieldclfrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrId"] = types.YLeaf{"Diffservmultifieldclfrid", diffservmultifieldclfrentry.Diffservmultifieldclfrid}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrAddrType"] = types.YLeaf{"Diffservmultifieldclfraddrtype", diffservmultifieldclfrentry.Diffservmultifieldclfraddrtype}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrDstAddr"] = types.YLeaf{"Diffservmultifieldclfrdstaddr", diffservmultifieldclfrentry.Diffservmultifieldclfrdstaddr}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrDstPrefixLength"] = types.YLeaf{"Diffservmultifieldclfrdstprefixlength", diffservmultifieldclfrentry.Diffservmultifieldclfrdstprefixlength}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrSrcAddr"] = types.YLeaf{"Diffservmultifieldclfrsrcaddr", diffservmultifieldclfrentry.Diffservmultifieldclfrsrcaddr}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrSrcPrefixLength"] = types.YLeaf{"Diffservmultifieldclfrsrcprefixlength", diffservmultifieldclfrentry.Diffservmultifieldclfrsrcprefixlength}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrDscp"] = types.YLeaf{"Diffservmultifieldclfrdscp", diffservmultifieldclfrentry.Diffservmultifieldclfrdscp}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrFlowId"] = types.YLeaf{"Diffservmultifieldclfrflowid", diffservmultifieldclfrentry.Diffservmultifieldclfrflowid}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrProtocol"] = types.YLeaf{"Diffservmultifieldclfrprotocol", diffservmultifieldclfrentry.Diffservmultifieldclfrprotocol}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrDstL4PortMin"] = types.YLeaf{"Diffservmultifieldclfrdstl4Portmin", diffservmultifieldclfrentry.Diffservmultifieldclfrdstl4Portmin}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrDstL4PortMax"] = types.YLeaf{"Diffservmultifieldclfrdstl4Portmax", diffservmultifieldclfrentry.Diffservmultifieldclfrdstl4Portmax}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrSrcL4PortMin"] = types.YLeaf{"Diffservmultifieldclfrsrcl4Portmin", diffservmultifieldclfrentry.Diffservmultifieldclfrsrcl4Portmin}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrSrcL4PortMax"] = types.YLeaf{"Diffservmultifieldclfrsrcl4Portmax", diffservmultifieldclfrentry.Diffservmultifieldclfrsrcl4Portmax}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrStorage"] = types.YLeaf{"Diffservmultifieldclfrstorage", diffservmultifieldclfrentry.Diffservmultifieldclfrstorage}
    diffservmultifieldclfrentry.EntityData.Leafs["diffServMultiFieldClfrStatus"] = types.YLeaf{"Diffservmultifieldclfrstatus", diffservmultifieldclfrentry.Diffservmultifieldclfrstatus}
    return &(diffservmultifieldclfrentry.EntityData)
}

// DIFFSERVMIB_Diffservmetertable
// This table enumerates specific meters that a system may use to
// police a stream of traffic. The traffic stream to be metered is
// determined by the Differentiated Services Functional Data Path
// Element(s) upstream of the meter i.e. by the object(s) that point
// to each entry in this table.  This may include all traffic on an
// interface.
// 
// Specific meter details are to be found in table entry referenced
// by diffServMeterSpecific.
type DIFFSERVMIB_Diffservmetertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the meter table describes a single conformance level of a
    // meter. The type is slice of
    // DIFFSERVMIB_Diffservmetertable_Diffservmeterentry.
    Diffservmeterentry []DIFFSERVMIB_Diffservmetertable_Diffservmeterentry
}

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetEntityData() *types.CommonEntityData {
    diffservmetertable.EntityData.YFilter = diffservmetertable.YFilter
    diffservmetertable.EntityData.YangName = "diffServMeterTable"
    diffservmetertable.EntityData.BundleName = "cisco_ios_xe"
    diffservmetertable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservmetertable.EntityData.SegmentPath = "diffServMeterTable"
    diffservmetertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservmetertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservmetertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservmetertable.EntityData.Children = make(map[string]types.YChild)
    diffservmetertable.EntityData.Children["diffServMeterEntry"] = types.YChild{"Diffservmeterentry", nil}
    for i := range diffservmetertable.Diffservmeterentry {
        diffservmetertable.EntityData.Children[types.GetSegmentPath(&diffservmetertable.Diffservmeterentry[i])] = types.YChild{"Diffservmeterentry", &diffservmetertable.Diffservmeterentry[i]}
    }
    diffservmetertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservmetertable.EntityData)
}

// DIFFSERVMIB_Diffservmetertable_Diffservmeterentry
// An entry in the meter table describes a single conformance level
// of a meter.
type DIFFSERVMIB_Diffservmetertable_Diffservmeterentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Meter entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServMeterNextFree. The type is interface{} with range: 1..4294967295.
    Diffservmeterid interface{}

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
    Diffservmetersucceednext interface{}

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
    Diffservmeterfailnext interface{}

    // This indicates the behavior of the meter by pointing to an entry containing
    // detailed parameters. Note that entries in that specific table must be
    // managed explicitly.  For example, diffServMeterSpecific may point to an
    // entry in diffServTBParamTable, which contains an instance of a single set
    // of Token Bucket parameters.  Setting this to point to a target that does
    // not exist results in an inconsistentValue error.  If the row pointed to is
    // removed or becomes inactive by other means, the meter always succeeds. The
    // type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservmeterspecific interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservmeterstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservmeterstatus interface{}
}

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetEntityData() *types.CommonEntityData {
    diffservmeterentry.EntityData.YFilter = diffservmeterentry.YFilter
    diffservmeterentry.EntityData.YangName = "diffServMeterEntry"
    diffservmeterentry.EntityData.BundleName = "cisco_ios_xe"
    diffservmeterentry.EntityData.ParentYangName = "diffServMeterTable"
    diffservmeterentry.EntityData.SegmentPath = "diffServMeterEntry" + "[diffServMeterId='" + fmt.Sprintf("%v", diffservmeterentry.Diffservmeterid) + "']"
    diffservmeterentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservmeterentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservmeterentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservmeterentry.EntityData.Children = make(map[string]types.YChild)
    diffservmeterentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservmeterentry.EntityData.Leafs["diffServMeterId"] = types.YLeaf{"Diffservmeterid", diffservmeterentry.Diffservmeterid}
    diffservmeterentry.EntityData.Leafs["diffServMeterSucceedNext"] = types.YLeaf{"Diffservmetersucceednext", diffservmeterentry.Diffservmetersucceednext}
    diffservmeterentry.EntityData.Leafs["diffServMeterFailNext"] = types.YLeaf{"Diffservmeterfailnext", diffservmeterentry.Diffservmeterfailnext}
    diffservmeterentry.EntityData.Leafs["diffServMeterSpecific"] = types.YLeaf{"Diffservmeterspecific", diffservmeterentry.Diffservmeterspecific}
    diffservmeterentry.EntityData.Leafs["diffServMeterStorage"] = types.YLeaf{"Diffservmeterstorage", diffservmeterentry.Diffservmeterstorage}
    diffservmeterentry.EntityData.Leafs["diffServMeterStatus"] = types.YLeaf{"Diffservmeterstatus", diffservmeterentry.Diffservmeterstatus}
    return &(diffservmeterentry.EntityData)
}

// DIFFSERVMIB_Diffservtbparamtable
// This table enumerates a single set of token bucket meter
// parameters that a system may use to police a stream of traffic.
// Such meters are modeled here as having a single rate and a single
// burst size. Multiple entries are used when multiple rates/burst
// sizes are needed.
type DIFFSERVMIB_Diffservtbparamtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry that describes a single set of token bucket parameters. The type
    // is slice of DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry.
    Diffservtbparamentry []DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry
}

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetEntityData() *types.CommonEntityData {
    diffservtbparamtable.EntityData.YFilter = diffservtbparamtable.YFilter
    diffservtbparamtable.EntityData.YangName = "diffServTBParamTable"
    diffservtbparamtable.EntityData.BundleName = "cisco_ios_xe"
    diffservtbparamtable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservtbparamtable.EntityData.SegmentPath = "diffServTBParamTable"
    diffservtbparamtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservtbparamtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservtbparamtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservtbparamtable.EntityData.Children = make(map[string]types.YChild)
    diffservtbparamtable.EntityData.Children["diffServTBParamEntry"] = types.YChild{"Diffservtbparamentry", nil}
    for i := range diffservtbparamtable.Diffservtbparamentry {
        diffservtbparamtable.EntityData.Children[types.GetSegmentPath(&diffservtbparamtable.Diffservtbparamentry[i])] = types.YChild{"Diffservtbparamentry", &diffservtbparamtable.Diffservtbparamentry[i]}
    }
    diffservtbparamtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservtbparamtable.EntityData)
}

// DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry
// An entry that describes a single set of token bucket
// parameters.
type DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Token Bucket
    // Parameter entries. Managers obtain new values for row creation in this
    // table by reading diffServTBParamNextFree. The type is interface{} with
    // range: 1..4294967295.
    Diffservtbparamid interface{}

    // The Metering algorithm associated with the Token Bucket parameters. 
    // zeroDotZero indicates this is unknown.  Standard values for generic
    // algorithms: diffServTBParamSimpleTokenBucket, diffServTBParamAvgRate,
    // diffServTBParamSrTCMBlind, diffServTBParamSrTCMAware,
    // diffServTBParamTrTCMBlind, diffServTBParamTrTCMAware, and
    // diffServTBParamTswTCM are specified in this MIB as OBJECT- IDENTITYs;
    // additional values may be further specified in other MIBs. The type is
    // string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservtbparamtype interface{}

    // The token-bucket rate, in kilobits per second (kbps). This attribute is
    // used for: 1. CIR in RFC 2697 for srTCM 2. CIR and PIR in RFC 2698 for trTCM
    // 3. CTR and PTR in RFC 2859 for TSWTCM 4. AverageRate in RFC 3290. The type
    // is interface{} with range: 1..4294967295. Units are kilobits per second.
    Diffservtbparamrate interface{}

    // The maximum number of bytes in a single transmission burst. This attribute
    // is used for: 1. CBS and EBS in RFC 2697 for srTCM 2. CBS and PBS in RFC
    // 2698 for trTCM 3. Burst Size in RFC 3290. The type is interface{} with
    // range: 0..2147483647. Units are Bytes.
    Diffservtbparamburstsize interface{}

    // The time interval used with the token bucket.  For: 1. Average Rate Meter,
    // the Informal Differentiated Services Model    section 5.2.1, - Delta. 2.
    // Simple Token Bucket Meter, the Informal Differentiated    Services Model
    // section 5.1, - time interval t. 3. RFC 2859 TSWTCM, - AVG_INTERVAL. 4. RFC
    // 2697 srTCM, RFC 2698 trTCM, - token bucket update time    interval. The
    // type is interface{} with range: 1..4294967295. Units are microseconds.
    Diffservtbparaminterval interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservtbparamstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservtbparamstatus interface{}
}

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetEntityData() *types.CommonEntityData {
    diffservtbparamentry.EntityData.YFilter = diffservtbparamentry.YFilter
    diffservtbparamentry.EntityData.YangName = "diffServTBParamEntry"
    diffservtbparamentry.EntityData.BundleName = "cisco_ios_xe"
    diffservtbparamentry.EntityData.ParentYangName = "diffServTBParamTable"
    diffservtbparamentry.EntityData.SegmentPath = "diffServTBParamEntry" + "[diffServTBParamId='" + fmt.Sprintf("%v", diffservtbparamentry.Diffservtbparamid) + "']"
    diffservtbparamentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservtbparamentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservtbparamentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservtbparamentry.EntityData.Children = make(map[string]types.YChild)
    diffservtbparamentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservtbparamentry.EntityData.Leafs["diffServTBParamId"] = types.YLeaf{"Diffservtbparamid", diffservtbparamentry.Diffservtbparamid}
    diffservtbparamentry.EntityData.Leafs["diffServTBParamType"] = types.YLeaf{"Diffservtbparamtype", diffservtbparamentry.Diffservtbparamtype}
    diffservtbparamentry.EntityData.Leafs["diffServTBParamRate"] = types.YLeaf{"Diffservtbparamrate", diffservtbparamentry.Diffservtbparamrate}
    diffservtbparamentry.EntityData.Leafs["diffServTBParamBurstSize"] = types.YLeaf{"Diffservtbparamburstsize", diffservtbparamentry.Diffservtbparamburstsize}
    diffservtbparamentry.EntityData.Leafs["diffServTBParamInterval"] = types.YLeaf{"Diffservtbparaminterval", diffservtbparamentry.Diffservtbparaminterval}
    diffservtbparamentry.EntityData.Leafs["diffServTBParamStorage"] = types.YLeaf{"Diffservtbparamstorage", diffservtbparamentry.Diffservtbparamstorage}
    diffservtbparamentry.EntityData.Leafs["diffServTBParamStatus"] = types.YLeaf{"Diffservtbparamstatus", diffservtbparamentry.Diffservtbparamstatus}
    return &(diffservtbparamentry.EntityData)
}

// DIFFSERVMIB_Diffservactiontable
// The Action Table enumerates actions that can be performed to a
// stream of traffic. Multiple actions can be concatenated. For
// example, traffic exiting from a meter may be counted, marked, and
// potentially dropped before entering a queue.
// 
// Specific actions are indicated by diffServActionSpecific which
// points to an entry of a specific action type parameterizing the
// action in detail.
type DIFFSERVMIB_Diffservactiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in the action table allows description of one specific action to
    // be applied to traffic. The type is slice of
    // DIFFSERVMIB_Diffservactiontable_Diffservactionentry.
    Diffservactionentry []DIFFSERVMIB_Diffservactiontable_Diffservactionentry
}

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetEntityData() *types.CommonEntityData {
    diffservactiontable.EntityData.YFilter = diffservactiontable.YFilter
    diffservactiontable.EntityData.YangName = "diffServActionTable"
    diffservactiontable.EntityData.BundleName = "cisco_ios_xe"
    diffservactiontable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservactiontable.EntityData.SegmentPath = "diffServActionTable"
    diffservactiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservactiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservactiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservactiontable.EntityData.Children = make(map[string]types.YChild)
    diffservactiontable.EntityData.Children["diffServActionEntry"] = types.YChild{"Diffservactionentry", nil}
    for i := range diffservactiontable.Diffservactionentry {
        diffservactiontable.EntityData.Children[types.GetSegmentPath(&diffservactiontable.Diffservactionentry[i])] = types.YChild{"Diffservactionentry", &diffservactiontable.Diffservactionentry[i]}
    }
    diffservactiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservactiontable.EntityData)
}

// DIFFSERVMIB_Diffservactiontable_Diffservactionentry
// Each entry in the action table allows description of one
// specific action to be applied to traffic.
type DIFFSERVMIB_Diffservactiontable_Diffservactionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Action entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServActionNextFree. The type is interface{} with range: 1..4294967295.
    Diffservactionid interface{}

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
    Diffservactioninterface interface{}

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
    Diffservactionnext interface{}

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
    Diffservactionspecific interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservactionstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservactionstatus interface{}
}

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetEntityData() *types.CommonEntityData {
    diffservactionentry.EntityData.YFilter = diffservactionentry.YFilter
    diffservactionentry.EntityData.YangName = "diffServActionEntry"
    diffservactionentry.EntityData.BundleName = "cisco_ios_xe"
    diffservactionentry.EntityData.ParentYangName = "diffServActionTable"
    diffservactionentry.EntityData.SegmentPath = "diffServActionEntry" + "[diffServActionId='" + fmt.Sprintf("%v", diffservactionentry.Diffservactionid) + "']"
    diffservactionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservactionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservactionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservactionentry.EntityData.Children = make(map[string]types.YChild)
    diffservactionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservactionentry.EntityData.Leafs["diffServActionId"] = types.YLeaf{"Diffservactionid", diffservactionentry.Diffservactionid}
    diffservactionentry.EntityData.Leafs["diffServActionInterface"] = types.YLeaf{"Diffservactioninterface", diffservactionentry.Diffservactioninterface}
    diffservactionentry.EntityData.Leafs["diffServActionNext"] = types.YLeaf{"Diffservactionnext", diffservactionentry.Diffservactionnext}
    diffservactionentry.EntityData.Leafs["diffServActionSpecific"] = types.YLeaf{"Diffservactionspecific", diffservactionentry.Diffservactionspecific}
    diffservactionentry.EntityData.Leafs["diffServActionStorage"] = types.YLeaf{"Diffservactionstorage", diffservactionentry.Diffservactionstorage}
    diffservactionentry.EntityData.Leafs["diffServActionStatus"] = types.YLeaf{"Diffservactionstatus", diffservactionentry.Diffservactionstatus}
    return &(diffservactionentry.EntityData)
}

// DIFFSERVMIB_Diffservdscpmarkacttable
// This table enumerates specific DSCPs used for marking or
// remarking the DSCP field of IP packets. The entries of this table
// may be referenced by a diffServActionSpecific attribute.
type DIFFSERVMIB_Diffservdscpmarkacttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the DSCP mark action table that describes a single DSCP used
    // for marking. The type is slice of
    // DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry.
    Diffservdscpmarkactentry []DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry
}

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetEntityData() *types.CommonEntityData {
    diffservdscpmarkacttable.EntityData.YFilter = diffservdscpmarkacttable.YFilter
    diffservdscpmarkacttable.EntityData.YangName = "diffServDscpMarkActTable"
    diffservdscpmarkacttable.EntityData.BundleName = "cisco_ios_xe"
    diffservdscpmarkacttable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservdscpmarkacttable.EntityData.SegmentPath = "diffServDscpMarkActTable"
    diffservdscpmarkacttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservdscpmarkacttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservdscpmarkacttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservdscpmarkacttable.EntityData.Children = make(map[string]types.YChild)
    diffservdscpmarkacttable.EntityData.Children["diffServDscpMarkActEntry"] = types.YChild{"Diffservdscpmarkactentry", nil}
    for i := range diffservdscpmarkacttable.Diffservdscpmarkactentry {
        diffservdscpmarkacttable.EntityData.Children[types.GetSegmentPath(&diffservdscpmarkacttable.Diffservdscpmarkactentry[i])] = types.YChild{"Diffservdscpmarkactentry", &diffservdscpmarkacttable.Diffservdscpmarkactentry[i]}
    }
    diffservdscpmarkacttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservdscpmarkacttable.EntityData)
}

// DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry
// An entry in the DSCP mark action table that describes a single
// DSCP used for marking.
type DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The DSCP that this Action will store into the DSCP
    // field of the subject. It is quite possible that the only packets subject to
    // this Action are already marked with this DSCP. Note also that
    // Differentiated Services processing may result in packet being marked on
    // both ingress to a network and on egress from it, and that ingress and
    // egress can occur in the same router. The type is interface{} with range:
    // 0..63.
    Diffservdscpmarkactdscp interface{}
}

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetEntityData() *types.CommonEntityData {
    diffservdscpmarkactentry.EntityData.YFilter = diffservdscpmarkactentry.YFilter
    diffservdscpmarkactentry.EntityData.YangName = "diffServDscpMarkActEntry"
    diffservdscpmarkactentry.EntityData.BundleName = "cisco_ios_xe"
    diffservdscpmarkactentry.EntityData.ParentYangName = "diffServDscpMarkActTable"
    diffservdscpmarkactentry.EntityData.SegmentPath = "diffServDscpMarkActEntry" + "[diffServDscpMarkActDscp='" + fmt.Sprintf("%v", diffservdscpmarkactentry.Diffservdscpmarkactdscp) + "']"
    diffservdscpmarkactentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservdscpmarkactentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservdscpmarkactentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservdscpmarkactentry.EntityData.Children = make(map[string]types.YChild)
    diffservdscpmarkactentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservdscpmarkactentry.EntityData.Leafs["diffServDscpMarkActDscp"] = types.YLeaf{"Diffservdscpmarkactdscp", diffservdscpmarkactentry.Diffservdscpmarkactdscp}
    return &(diffservdscpmarkactentry.EntityData)
}

// DIFFSERVMIB_Diffservcountacttable
// This table contains counters for all the traffic passing through
// an action element.
type DIFFSERVMIB_Diffservcountacttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the count action table describes a single set of traffic
    // counters. The type is slice of
    // DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry.
    Diffservcountactentry []DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry
}

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetEntityData() *types.CommonEntityData {
    diffservcountacttable.EntityData.YFilter = diffservcountacttable.YFilter
    diffservcountacttable.EntityData.YangName = "diffServCountActTable"
    diffservcountacttable.EntityData.BundleName = "cisco_ios_xe"
    diffservcountacttable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservcountacttable.EntityData.SegmentPath = "diffServCountActTable"
    diffservcountacttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservcountacttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservcountacttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservcountacttable.EntityData.Children = make(map[string]types.YChild)
    diffservcountacttable.EntityData.Children["diffServCountActEntry"] = types.YChild{"Diffservcountactentry", nil}
    for i := range diffservcountacttable.Diffservcountactentry {
        diffservcountacttable.EntityData.Children[types.GetSegmentPath(&diffservcountacttable.Diffservcountactentry[i])] = types.YChild{"Diffservcountactentry", &diffservcountacttable.Diffservcountactentry[i]}
    }
    diffservcountacttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservcountacttable.EntityData)
}

// DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry
// An entry in the count action table describes a single set of
// traffic counters.
type DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Count Action entries.
    // Managers obtain new values for row creation in this table by reading   
    // diffServCountActNextFree. The type is interface{} with range:
    // 1..4294967295.
    Diffservcountactid interface{}

    // The number of octets at the Action data path element.  Discontinuities in
    // the value of this counter can occur at re- initialization of the management
    // system and at other times as indicated by the value of
    // ifCounterDiscontinuityTime on the relevant interface. The type is
    // interface{} with range: 0..18446744073709551615.
    Diffservcountactoctets interface{}

    // The number of packets at the Action data path element.  Discontinuities in
    // the value of this counter can occur at re- initialization of the management
    // system and at other times as indicated by the value of
    // ifCounterDiscontinuityTime on the relevant interface. The type is
    // interface{} with range: 0..18446744073709551615.
    Diffservcountactpkts interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservcountactstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing    to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservcountactstatus interface{}
}

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetEntityData() *types.CommonEntityData {
    diffservcountactentry.EntityData.YFilter = diffservcountactentry.YFilter
    diffservcountactentry.EntityData.YangName = "diffServCountActEntry"
    diffservcountactentry.EntityData.BundleName = "cisco_ios_xe"
    diffservcountactentry.EntityData.ParentYangName = "diffServCountActTable"
    diffservcountactentry.EntityData.SegmentPath = "diffServCountActEntry" + "[diffServCountActId='" + fmt.Sprintf("%v", diffservcountactentry.Diffservcountactid) + "']"
    diffservcountactentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservcountactentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservcountactentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservcountactentry.EntityData.Children = make(map[string]types.YChild)
    diffservcountactentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservcountactentry.EntityData.Leafs["diffServCountActId"] = types.YLeaf{"Diffservcountactid", diffservcountactentry.Diffservcountactid}
    diffservcountactentry.EntityData.Leafs["diffServCountActOctets"] = types.YLeaf{"Diffservcountactoctets", diffservcountactentry.Diffservcountactoctets}
    diffservcountactentry.EntityData.Leafs["diffServCountActPkts"] = types.YLeaf{"Diffservcountactpkts", diffservcountactentry.Diffservcountactpkts}
    diffservcountactentry.EntityData.Leafs["diffServCountActStorage"] = types.YLeaf{"Diffservcountactstorage", diffservcountactentry.Diffservcountactstorage}
    diffservcountactentry.EntityData.Leafs["diffServCountActStatus"] = types.YLeaf{"Diffservcountactstatus", diffservcountactentry.Diffservcountactstatus}
    return &(diffservcountactentry.EntityData)
}

// DIFFSERVMIB_Diffservalgdroptable
// The algorithmic drop table contains entries describing an
// element that drops packets according to some algorithm.
type DIFFSERVMIB_Diffservalgdroptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry describes a process that drops packets according to some
    // algorithm. Further details of the algorithm type are to be found in
    // diffServAlgDropType and with more detail parameter entry pointed to by
    // diffServAlgDropSpecific when necessary. The type is slice of
    // DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry.
    Diffservalgdropentry []DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry
}

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetEntityData() *types.CommonEntityData {
    diffservalgdroptable.EntityData.YFilter = diffservalgdroptable.YFilter
    diffservalgdroptable.EntityData.YangName = "diffServAlgDropTable"
    diffservalgdroptable.EntityData.BundleName = "cisco_ios_xe"
    diffservalgdroptable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservalgdroptable.EntityData.SegmentPath = "diffServAlgDropTable"
    diffservalgdroptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservalgdroptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservalgdroptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservalgdroptable.EntityData.Children = make(map[string]types.YChild)
    diffservalgdroptable.EntityData.Children["diffServAlgDropEntry"] = types.YChild{"Diffservalgdropentry", nil}
    for i := range diffservalgdroptable.Diffservalgdropentry {
        diffservalgdroptable.EntityData.Children[types.GetSegmentPath(&diffservalgdroptable.Diffservalgdropentry[i])] = types.YChild{"Diffservalgdropentry", &diffservalgdroptable.Diffservalgdropentry[i]}
    }
    diffservalgdroptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservalgdroptable.EntityData)
}

// DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry
// An entry describes a process that drops packets according to
// some algorithm. Further details of the algorithm type are to be
// found in diffServAlgDropType and with more detail parameter entry
// pointed to by diffServAlgDropSpecific when necessary.
type DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Algorithmic Dropper
    // entries. Managers obtain new values for row creation in this table by
    // reading diffServAlgDropNextFree. The type is interface{} with range:
    // 1..4294967295.
    Diffservalgdropid interface{}

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
    // The type is Diffservalgdroptype.
    Diffservalgdroptype interface{}

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
    Diffservalgdropnext interface{}

    // Points to an entry in the diffServQTable to indicate the queue that a drop
    // algorithm is to monitor when deciding whether to drop a packet. If the row
    // pointed to does not exist, the algorithmic dropper element is considered
    // inactive.    Setting this to point to a target that does not exist results
    // in an inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservalgdropqmeasure interface{}

    // A threshold on the depth in bytes of the queue being measured at which a
    // trigger is generated to the dropping algorithm, unless diffServAlgDropType
    // is alwaysDrop(5) where this object is ignored.  For the tailDrop(2) or
    // headDrop(3) algorithms, this represents the depth of the queue, pointed to
    // by diffServAlgDropQMeasure, at which the drop action will take place. Other
    // algorithms will need to define their own semantics for this threshold. The
    // type is interface{} with range: 1..4294967295. Units are Bytes.
    Diffservalgdropqthreshold interface{}

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
    Diffservalgdropspecific interface{}

    // The number of octets that have been deterministically dropped by this drop
    // process.  Discontinuities in the value of this counter can occur at re-
    // initialization of the management system and at other times as indicated by
    // the value of ifCounterDiscontinuityTime on the relevant interface. The type
    // is interface{} with range: 0..18446744073709551615.
    Diffservalgdropoctets interface{}

    // The number of packets that have been deterministically dropped by this drop
    // process.  Discontinuities in the value of this counter can occur at re-
    // initialization of the management system and at other times as indicated by
    // the value of ifCounterDiscontinuityTime on the relevant interface. The type
    // is interface{} with range: 0..18446744073709551615.
    Diffservalgdroppkts interface{}

    // The number of octets that have been randomly dropped by this drop process. 
    // This counter applies, therefore, only to random droppers.  Discontinuities
    // in the value of this counter can occur at re- initialization of the
    // management system and at other times as indicated by the value of
    // ifCounterDiscontinuityTime on the relevant interface. The type is
    // interface{} with range: 0..18446744073709551615.
    Diffservalgrandomdropoctets interface{}

    // The number of packets that have been randomly dropped by this drop process.
    // This counter applies, therefore, only to random droppers.  Discontinuities
    // in the value of this counter can occur at re- initialization of the
    // management system and at other times as indicated by the value of
    // ifCounterDiscontinuityTime on the relevant interface. The type is
    // interface{} with range: 0..18446744073709551615.
    Diffservalgrandomdroppkts interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservalgdropstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservalgdropstatus interface{}
}

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetEntityData() *types.CommonEntityData {
    diffservalgdropentry.EntityData.YFilter = diffservalgdropentry.YFilter
    diffservalgdropentry.EntityData.YangName = "diffServAlgDropEntry"
    diffservalgdropentry.EntityData.BundleName = "cisco_ios_xe"
    diffservalgdropentry.EntityData.ParentYangName = "diffServAlgDropTable"
    diffservalgdropentry.EntityData.SegmentPath = "diffServAlgDropEntry" + "[diffServAlgDropId='" + fmt.Sprintf("%v", diffservalgdropentry.Diffservalgdropid) + "']"
    diffservalgdropentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservalgdropentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservalgdropentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservalgdropentry.EntityData.Children = make(map[string]types.YChild)
    diffservalgdropentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropId"] = types.YLeaf{"Diffservalgdropid", diffservalgdropentry.Diffservalgdropid}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropType"] = types.YLeaf{"Diffservalgdroptype", diffservalgdropentry.Diffservalgdroptype}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropNext"] = types.YLeaf{"Diffservalgdropnext", diffservalgdropentry.Diffservalgdropnext}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropQMeasure"] = types.YLeaf{"Diffservalgdropqmeasure", diffservalgdropentry.Diffservalgdropqmeasure}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropQThreshold"] = types.YLeaf{"Diffservalgdropqthreshold", diffservalgdropentry.Diffservalgdropqthreshold}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropSpecific"] = types.YLeaf{"Diffservalgdropspecific", diffservalgdropentry.Diffservalgdropspecific}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropOctets"] = types.YLeaf{"Diffservalgdropoctets", diffservalgdropentry.Diffservalgdropoctets}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropPkts"] = types.YLeaf{"Diffservalgdroppkts", diffservalgdropentry.Diffservalgdroppkts}
    diffservalgdropentry.EntityData.Leafs["diffServAlgRandomDropOctets"] = types.YLeaf{"Diffservalgrandomdropoctets", diffservalgdropentry.Diffservalgrandomdropoctets}
    diffservalgdropentry.EntityData.Leafs["diffServAlgRandomDropPkts"] = types.YLeaf{"Diffservalgrandomdroppkts", diffservalgdropentry.Diffservalgrandomdroppkts}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropStorage"] = types.YLeaf{"Diffservalgdropstorage", diffservalgdropentry.Diffservalgdropstorage}
    diffservalgdropentry.EntityData.Leafs["diffServAlgDropStatus"] = types.YLeaf{"Diffservalgdropstatus", diffservalgdropentry.Diffservalgdropstatus}
    return &(diffservalgdropentry.EntityData)
}

// DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype represents and diffServAlgDropSpecific are all zeroDotZero.
type DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype string

const (
    DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype_other DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype = "other"

    DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype_tailDrop DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype = "tailDrop"

    DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype_headDrop DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype = "headDrop"

    DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype_randomDrop DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype = "randomDrop"

    DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype_alwaysDrop DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry_Diffservalgdroptype = "alwaysDrop"
)

// DIFFSERVMIB_Diffservrandomdroptable
// The random drop table contains entries describing a process that
// drops packets randomly. Entries in this table are pointed to by
// diffServAlgDropSpecific.
type DIFFSERVMIB_Diffservrandomdroptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry describes a process that drops packets according to a random
    // algorithm. The type is slice of
    // DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry.
    Diffservrandomdropentry []DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry
}

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetEntityData() *types.CommonEntityData {
    diffservrandomdroptable.EntityData.YFilter = diffservrandomdroptable.YFilter
    diffservrandomdroptable.EntityData.YangName = "diffServRandomDropTable"
    diffservrandomdroptable.EntityData.BundleName = "cisco_ios_xe"
    diffservrandomdroptable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservrandomdroptable.EntityData.SegmentPath = "diffServRandomDropTable"
    diffservrandomdroptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservrandomdroptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservrandomdroptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservrandomdroptable.EntityData.Children = make(map[string]types.YChild)
    diffservrandomdroptable.EntityData.Children["diffServRandomDropEntry"] = types.YChild{"Diffservrandomdropentry", nil}
    for i := range diffservrandomdroptable.Diffservrandomdropentry {
        diffservrandomdroptable.EntityData.Children[types.GetSegmentPath(&diffservrandomdroptable.Diffservrandomdropentry[i])] = types.YChild{"Diffservrandomdropentry", &diffservrandomdroptable.Diffservrandomdropentry[i]}
    }
    diffservrandomdroptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservrandomdroptable.EntityData)
}

// DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry
// An entry describes a process that drops packets according to a
// random algorithm.
type DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Random Drop entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServRandomDropNextFree. The type is interface{} with range:
    // 1..4294967295.
    Diffservrandomdropid interface{}

    // The average queue depth in bytes, beyond which traffic has a non-zero
    // probability of being dropped. Changes in this variable may or may not be
    // reflected in the reported value of diffServRandomDropMinThreshPkts. The
    // type is interface{} with range: 1..4294967295. Units are bytes.
    Diffservrandomdropminthreshbytes interface{}

    // The average queue depth in packets, beyond which traffic has a non-zero
    // probability of being dropped. Changes in this variable may or may not be
    // reflected in the reported value of diffServRandomDropMinThreshBytes. The
    // type is interface{} with range: 1..4294967295. Units are packets.
    Diffservrandomdropminthreshpkts interface{}

    // The average queue depth beyond which traffic has a probability indicated by
    // diffServRandomDropProbMax of being dropped or marked. Note that this
    // differs from the physical queue limit, which is stored in
    // diffServAlgDropQThreshold. Changes in this variable may or may not be
    // reflected in the reported value of diffServRandomDropMaxThreshPkts. The
    // type is interface{} with range: 1..4294967295. Units are bytes.
    Diffservrandomdropmaxthreshbytes interface{}

    // The average queue depth beyond which traffic has a probability indicated by
    // diffServRandomDropProbMax of being dropped or marked. Note that this
    // differs from the physical queue limit, which is stored in
    // diffServAlgDropQThreshold. Changes in this variable may or may not be
    // reflected in the reported value of diffServRandomDropMaxThreshBytes. The
    // type is interface{} with range: 1..4294967295. Units are packets.
    Diffservrandomdropmaxthreshpkts interface{}

    // The worst case random drop probability, expressed in drops per thousand
    // packets.  For example, if in the worst case every arriving packet may be
    // dropped (100%) for a period, this has the value 1000. Alternatively, if in
    // the worst case only one percent (1%) of traffic may be dropped, it has the
    // value 10. The type is interface{} with range: 0..1000.
    Diffservrandomdropprobmax interface{}

    // The weighting of past history in affecting the Exponentially Weighted
    // Moving Average function that calculates the current average queue depth. 
    // The equation uses diffServRandomDropWeight/65536 as the coefficient for the
    // new sample in the equation, and (65536 - diffServRandomDropWeight)/65536 as
    // the coefficient of the old value.  Implementations may limit the values of
    // diffServRandomDropWeight to a subset of the possible range of values, such
    // as powers of two. Doing this would facilitate implementation of the
    // Exponentially Weighted Moving Average using shift instructions or
    // registers. The type is interface{} with range: 0..65536.
    Diffservrandomdropweight interface{}

    // The number of times per second the queue is sampled for queue average
    // calculation.  A value of zero is used to mean that the queue is sampled
    // approximately each time a packet is enqueued (or dequeued). The type is
    // interface{} with range: 0..1000000.
    Diffservrandomdropsamplingrate interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservrandomdropstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservrandomdropstatus interface{}
}

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetEntityData() *types.CommonEntityData {
    diffservrandomdropentry.EntityData.YFilter = diffservrandomdropentry.YFilter
    diffservrandomdropentry.EntityData.YangName = "diffServRandomDropEntry"
    diffservrandomdropentry.EntityData.BundleName = "cisco_ios_xe"
    diffservrandomdropentry.EntityData.ParentYangName = "diffServRandomDropTable"
    diffservrandomdropentry.EntityData.SegmentPath = "diffServRandomDropEntry" + "[diffServRandomDropId='" + fmt.Sprintf("%v", diffservrandomdropentry.Diffservrandomdropid) + "']"
    diffservrandomdropentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservrandomdropentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservrandomdropentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservrandomdropentry.EntityData.Children = make(map[string]types.YChild)
    diffservrandomdropentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropId"] = types.YLeaf{"Diffservrandomdropid", diffservrandomdropentry.Diffservrandomdropid}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropMinThreshBytes"] = types.YLeaf{"Diffservrandomdropminthreshbytes", diffservrandomdropentry.Diffservrandomdropminthreshbytes}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropMinThreshPkts"] = types.YLeaf{"Diffservrandomdropminthreshpkts", diffservrandomdropentry.Diffservrandomdropminthreshpkts}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropMaxThreshBytes"] = types.YLeaf{"Diffservrandomdropmaxthreshbytes", diffservrandomdropentry.Diffservrandomdropmaxthreshbytes}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropMaxThreshPkts"] = types.YLeaf{"Diffservrandomdropmaxthreshpkts", diffservrandomdropentry.Diffservrandomdropmaxthreshpkts}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropProbMax"] = types.YLeaf{"Diffservrandomdropprobmax", diffservrandomdropentry.Diffservrandomdropprobmax}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropWeight"] = types.YLeaf{"Diffservrandomdropweight", diffservrandomdropentry.Diffservrandomdropweight}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropSamplingRate"] = types.YLeaf{"Diffservrandomdropsamplingrate", diffservrandomdropentry.Diffservrandomdropsamplingrate}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropStorage"] = types.YLeaf{"Diffservrandomdropstorage", diffservrandomdropentry.Diffservrandomdropstorage}
    diffservrandomdropentry.EntityData.Leafs["diffServRandomDropStatus"] = types.YLeaf{"Diffservrandomdropstatus", diffservrandomdropentry.Diffservrandomdropstatus}
    return &(diffservrandomdropentry.EntityData)
}

// DIFFSERVMIB_Diffservqtable
// The Queue Table enumerates the individual queues.  Note that the
// MIB models queuing systems as composed of individual queues, one
// per class of traffic, even though they may in fact be structured
// as classes of traffic scheduled using a common calendar queue, or
// in other ways.
type DIFFSERVMIB_Diffservqtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Queue Table describes a single queue or class of traffic.
    // The type is slice of DIFFSERVMIB_Diffservqtable_Diffservqentry.
    Diffservqentry []DIFFSERVMIB_Diffservqtable_Diffservqentry
}

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetEntityData() *types.CommonEntityData {
    diffservqtable.EntityData.YFilter = diffservqtable.YFilter
    diffservqtable.EntityData.YangName = "diffServQTable"
    diffservqtable.EntityData.BundleName = "cisco_ios_xe"
    diffservqtable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservqtable.EntityData.SegmentPath = "diffServQTable"
    diffservqtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservqtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservqtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservqtable.EntityData.Children = make(map[string]types.YChild)
    diffservqtable.EntityData.Children["diffServQEntry"] = types.YChild{"Diffservqentry", nil}
    for i := range diffservqtable.Diffservqentry {
        diffservqtable.EntityData.Children[types.GetSegmentPath(&diffservqtable.Diffservqentry[i])] = types.YChild{"Diffservqentry", &diffservqtable.Diffservqentry[i]}
    }
    diffservqtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservqtable.EntityData)
}

// DIFFSERVMIB_Diffservqtable_Diffservqentry
// An entry in the Queue Table describes a single queue or class of
// traffic.
type DIFFSERVMIB_Diffservqtable_Diffservqentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Queue entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServQNextFree. The type is interface{} with range: 1..4294967295.
    Diffservqid interface{}

    // This selects the next Differentiated Services Scheduler.  The RowPointer
    // must point to a diffServSchedulerEntry.  A value of zeroDotZero in this
    // attribute indicates an incomplete diffServQEntry instance. In such a case,
    // the entry has no operational effect, since it has no parameters to give it
    // meaning.  Setting this to point to a target that does not exist results in
    // an inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservqnext interface{}

    // This RowPointer indicates the diffServMinRateEntry that the scheduler,
    // pointed to by diffServQNext, should use to service this queue.  If the row
    // pointed to is zeroDotZero, the minimum rate and priority is unspecified. 
    // Setting this to point to a target that does not exist results in an
    // inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservqminrate interface{}

    // This RowPointer indicates the diffServMaxRateEntry that the scheduler,
    // pointed to by diffServQNext, should use to service this queue.  If the row
    // pointed to is zeroDotZero, the maximum rate is the line speed of the
    // interface.     Setting this to point to a target that does not exist
    // results in an inconsistentValue error.  If the row pointed to is removed or
    // becomes inactive by other means, the treatment is as if this attribute
    // contains a value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservqmaxrate interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservqstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservqstatus interface{}
}

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetEntityData() *types.CommonEntityData {
    diffservqentry.EntityData.YFilter = diffservqentry.YFilter
    diffservqentry.EntityData.YangName = "diffServQEntry"
    diffservqentry.EntityData.BundleName = "cisco_ios_xe"
    diffservqentry.EntityData.ParentYangName = "diffServQTable"
    diffservqentry.EntityData.SegmentPath = "diffServQEntry" + "[diffServQId='" + fmt.Sprintf("%v", diffservqentry.Diffservqid) + "']"
    diffservqentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservqentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservqentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservqentry.EntityData.Children = make(map[string]types.YChild)
    diffservqentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservqentry.EntityData.Leafs["diffServQId"] = types.YLeaf{"Diffservqid", diffservqentry.Diffservqid}
    diffservqentry.EntityData.Leafs["diffServQNext"] = types.YLeaf{"Diffservqnext", diffservqentry.Diffservqnext}
    diffservqentry.EntityData.Leafs["diffServQMinRate"] = types.YLeaf{"Diffservqminrate", diffservqentry.Diffservqminrate}
    diffservqentry.EntityData.Leafs["diffServQMaxRate"] = types.YLeaf{"Diffservqmaxrate", diffservqentry.Diffservqmaxrate}
    diffservqentry.EntityData.Leafs["diffServQStorage"] = types.YLeaf{"Diffservqstorage", diffservqentry.Diffservqstorage}
    diffservqentry.EntityData.Leafs["diffServQStatus"] = types.YLeaf{"Diffservqstatus", diffservqentry.Diffservqstatus}
    return &(diffservqentry.EntityData)
}

// DIFFSERVMIB_Diffservschedulertable
// The Scheduler Table enumerates packet schedulers. Multiple
// scheduling algorithms can be used on a given data path, with each
// algorithm described by one diffServSchedulerEntry.
type DIFFSERVMIB_Diffservschedulertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Scheduler Table describing a single instance of a
    // scheduling algorithm. The type is slice of
    // DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry.
    Diffservschedulerentry []DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry
}

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetEntityData() *types.CommonEntityData {
    diffservschedulertable.EntityData.YFilter = diffservschedulertable.YFilter
    diffservschedulertable.EntityData.YangName = "diffServSchedulerTable"
    diffservschedulertable.EntityData.BundleName = "cisco_ios_xe"
    diffservschedulertable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservschedulertable.EntityData.SegmentPath = "diffServSchedulerTable"
    diffservschedulertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservschedulertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservschedulertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservschedulertable.EntityData.Children = make(map[string]types.YChild)
    diffservschedulertable.EntityData.Children["diffServSchedulerEntry"] = types.YChild{"Diffservschedulerentry", nil}
    for i := range diffservschedulertable.Diffservschedulerentry {
        diffservschedulertable.EntityData.Children[types.GetSegmentPath(&diffservschedulertable.Diffservschedulerentry[i])] = types.YChild{"Diffservschedulerentry", &diffservschedulertable.Diffservschedulerentry[i]}
    }
    diffservschedulertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservschedulertable.EntityData)
}

// DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry
// An entry in the Scheduler Table describing a single instance of
// a scheduling algorithm.
type DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Scheduler entries. 
    // Managers obtain new values for row creation in this table by reading
    // diffServSchedulerNextFree. The type is interface{} with range:
    // 1..4294967295.
    Diffservschedulerid interface{}

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
    Diffservschedulernext interface{}

    // The scheduling algorithm used by this Scheduler. zeroDotZero indicates that
    // this is unknown.  Standard values for generic algorithms:
    // diffServSchedulerPriority, diffServSchedulerWRR, and diffServSchedulerWFQ
    // are specified in this MIB; additional values    may be further specified in
    // other MIBs. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservschedulermethod interface{}

    // This RowPointer indicates the entry in diffServMinRateTable which indicates
    // the priority or minimum output rate from this scheduler. This attribute is
    // used only when there is more than one level of scheduler.  When it has the
    // value zeroDotZero, it indicates that no minimum rate or priority is
    // imposed.  Setting this to point to a target that does not exist results in
    // an inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Diffservschedulerminrate interface{}

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
    Diffservschedulermaxrate interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservschedulerstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservschedulerstatus interface{}
}

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetEntityData() *types.CommonEntityData {
    diffservschedulerentry.EntityData.YFilter = diffservschedulerentry.YFilter
    diffservschedulerentry.EntityData.YangName = "diffServSchedulerEntry"
    diffservschedulerentry.EntityData.BundleName = "cisco_ios_xe"
    diffservschedulerentry.EntityData.ParentYangName = "diffServSchedulerTable"
    diffservschedulerentry.EntityData.SegmentPath = "diffServSchedulerEntry" + "[diffServSchedulerId='" + fmt.Sprintf("%v", diffservschedulerentry.Diffservschedulerid) + "']"
    diffservschedulerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservschedulerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservschedulerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservschedulerentry.EntityData.Children = make(map[string]types.YChild)
    diffservschedulerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservschedulerentry.EntityData.Leafs["diffServSchedulerId"] = types.YLeaf{"Diffservschedulerid", diffservschedulerentry.Diffservschedulerid}
    diffservschedulerentry.EntityData.Leafs["diffServSchedulerNext"] = types.YLeaf{"Diffservschedulernext", diffservschedulerentry.Diffservschedulernext}
    diffservschedulerentry.EntityData.Leafs["diffServSchedulerMethod"] = types.YLeaf{"Diffservschedulermethod", diffservschedulerentry.Diffservschedulermethod}
    diffservschedulerentry.EntityData.Leafs["diffServSchedulerMinRate"] = types.YLeaf{"Diffservschedulerminrate", diffservschedulerentry.Diffservschedulerminrate}
    diffservschedulerentry.EntityData.Leafs["diffServSchedulerMaxRate"] = types.YLeaf{"Diffservschedulermaxrate", diffservschedulerentry.Diffservschedulermaxrate}
    diffservschedulerentry.EntityData.Leafs["diffServSchedulerStorage"] = types.YLeaf{"Diffservschedulerstorage", diffservschedulerentry.Diffservschedulerstorage}
    diffservschedulerentry.EntityData.Leafs["diffServSchedulerStatus"] = types.YLeaf{"Diffservschedulerstatus", diffservschedulerentry.Diffservschedulerstatus}
    return &(diffservschedulerentry.EntityData)
}

// DIFFSERVMIB_Diffservminratetable
// The Minimum Rate Parameters Table enumerates individual sets of
// scheduling parameter that can be used/reused by Queues and
// Schedulers.
type DIFFSERVMIB_Diffservminratetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Minimum Rate Parameters Table describes a single set of
    // scheduling parameters for use by one or more queues or schedulers. The type
    // is slice of DIFFSERVMIB_Diffservminratetable_Diffservminrateentry.
    Diffservminrateentry []DIFFSERVMIB_Diffservminratetable_Diffservminrateentry
}

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetEntityData() *types.CommonEntityData {
    diffservminratetable.EntityData.YFilter = diffservminratetable.YFilter
    diffservminratetable.EntityData.YangName = "diffServMinRateTable"
    diffservminratetable.EntityData.BundleName = "cisco_ios_xe"
    diffservminratetable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservminratetable.EntityData.SegmentPath = "diffServMinRateTable"
    diffservminratetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservminratetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservminratetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservminratetable.EntityData.Children = make(map[string]types.YChild)
    diffservminratetable.EntityData.Children["diffServMinRateEntry"] = types.YChild{"Diffservminrateentry", nil}
    for i := range diffservminratetable.Diffservminrateentry {
        diffservminratetable.EntityData.Children[types.GetSegmentPath(&diffservminratetable.Diffservminrateentry[i])] = types.YChild{"Diffservminrateentry", &diffservminratetable.Diffservminrateentry[i]}
    }
    diffservminratetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservminratetable.EntityData)
}

// DIFFSERVMIB_Diffservminratetable_Diffservminrateentry
// An entry in the Minimum Rate Parameters Table describes a single
// set of scheduling parameters for use by one or more queues or
// schedulers.
type DIFFSERVMIB_Diffservminratetable_Diffservminrateentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Scheduler Parameter
    // entries. Managers obtain new values for row creation in this table by
    // reading diffServMinRateNextFree. The type is interface{} with range:
    // 1..4294967295.
    Diffservminrateid interface{}

    // The priority of this input to the associated scheduler, relative    to the
    // scheduler's other inputs. A queue or scheduler with a larger numeric value
    // will be served before another with a smaller numeric value. The type is
    // interface{} with range: 1..4294967295.
    Diffservminratepriority interface{}

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
    Diffservminrateabsolute interface{}

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
    Diffservminraterelative interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservminratestorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservminratestatus interface{}
}

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetEntityData() *types.CommonEntityData {
    diffservminrateentry.EntityData.YFilter = diffservminrateentry.YFilter
    diffservminrateentry.EntityData.YangName = "diffServMinRateEntry"
    diffservminrateentry.EntityData.BundleName = "cisco_ios_xe"
    diffservminrateentry.EntityData.ParentYangName = "diffServMinRateTable"
    diffservminrateentry.EntityData.SegmentPath = "diffServMinRateEntry" + "[diffServMinRateId='" + fmt.Sprintf("%v", diffservminrateentry.Diffservminrateid) + "']"
    diffservminrateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservminrateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservminrateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservminrateentry.EntityData.Children = make(map[string]types.YChild)
    diffservminrateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservminrateentry.EntityData.Leafs["diffServMinRateId"] = types.YLeaf{"Diffservminrateid", diffservminrateentry.Diffservminrateid}
    diffservminrateentry.EntityData.Leafs["diffServMinRatePriority"] = types.YLeaf{"Diffservminratepriority", diffservminrateentry.Diffservminratepriority}
    diffservminrateentry.EntityData.Leafs["diffServMinRateAbsolute"] = types.YLeaf{"Diffservminrateabsolute", diffservminrateentry.Diffservminrateabsolute}
    diffservminrateentry.EntityData.Leafs["diffServMinRateRelative"] = types.YLeaf{"Diffservminraterelative", diffservminrateentry.Diffservminraterelative}
    diffservminrateentry.EntityData.Leafs["diffServMinRateStorage"] = types.YLeaf{"Diffservminratestorage", diffservminrateentry.Diffservminratestorage}
    diffservminrateentry.EntityData.Leafs["diffServMinRateStatus"] = types.YLeaf{"Diffservminratestatus", diffservminrateentry.Diffservminratestatus}
    return &(diffservminrateentry.EntityData)
}

// DIFFSERVMIB_Diffservmaxratetable
// The Maximum Rate Parameter Table enumerates individual sets of
// scheduling parameter that can be used/reused by Queues and
// Schedulers.
type DIFFSERVMIB_Diffservmaxratetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the Maximum Rate Parameter Table describes a single set of
    // scheduling parameters for use by one or more queues or schedulers. The type
    // is slice of DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry.
    Diffservmaxrateentry []DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry
}

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetEntityData() *types.CommonEntityData {
    diffservmaxratetable.EntityData.YFilter = diffservmaxratetable.YFilter
    diffservmaxratetable.EntityData.YangName = "diffServMaxRateTable"
    diffservmaxratetable.EntityData.BundleName = "cisco_ios_xe"
    diffservmaxratetable.EntityData.ParentYangName = "DIFFSERV-MIB"
    diffservmaxratetable.EntityData.SegmentPath = "diffServMaxRateTable"
    diffservmaxratetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservmaxratetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservmaxratetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservmaxratetable.EntityData.Children = make(map[string]types.YChild)
    diffservmaxratetable.EntityData.Children["diffServMaxRateEntry"] = types.YChild{"Diffservmaxrateentry", nil}
    for i := range diffservmaxratetable.Diffservmaxrateentry {
        diffservmaxratetable.EntityData.Children[types.GetSegmentPath(&diffservmaxratetable.Diffservmaxrateentry[i])] = types.YChild{"Diffservmaxrateentry", &diffservmaxratetable.Diffservmaxrateentry[i]}
    }
    diffservmaxratetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservmaxratetable.EntityData)
}

// DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry
// An entry in the Maximum Rate Parameter Table describes a single
// set of scheduling parameters for use by one or more queues or
// schedulers.
type DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An index that enumerates the Maximum Rate
    // Parameter entries. Managers obtain new values for row creation in this
    // table by reading diffServMaxRateNextFree. The type is interface{} with
    // range: 1..4294967295.
    Diffservmaxrateid interface{}

    // This attribute is a key. An index that indicates which level of a
    // multi-rate shaper is being given its parameters. A multi-rate shaper has
    // some number of rate levels. Frame Relay's dual rate specification refers to
    // a 'committed' and an 'excess' rate; ATM's dual rate specification refers to
    // a 'mean' and a 'peak' rate. This table is generalized to support an
    // arbitrary number of rates. The committed or mean rate is level 1, the peak
    // rate (if any) is the highest level rate configured, and if there are other
    // rates they are distributed in monotonically increasing order between them.
    // The type is interface{} with range: 1..32.
    Diffservmaxratelevel interface{}

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
    Diffservmaxrateabsolute interface{}

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
    Diffservmaxraterelative interface{}

    // The number of bytes of queue depth at which the rate of a    multi-rate
    // scheduler will increase to the next output rate. In the last conceptual row
    // for such a shaper, this threshold is ignored and by convention is zero. The
    // type is interface{} with range: 0..2147483647. Units are Bytes.
    Diffservmaxratethreshold interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservmaxratestorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. Setting this variable to 'destroy' when the MIB
    // contains one or more RowPointers pointing to it results in destruction
    // being delayed until the row is no longer used. The type is RowStatus.
    Diffservmaxratestatus interface{}
}

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetEntityData() *types.CommonEntityData {
    diffservmaxrateentry.EntityData.YFilter = diffservmaxrateentry.YFilter
    diffservmaxrateentry.EntityData.YangName = "diffServMaxRateEntry"
    diffservmaxrateentry.EntityData.BundleName = "cisco_ios_xe"
    diffservmaxrateentry.EntityData.ParentYangName = "diffServMaxRateTable"
    diffservmaxrateentry.EntityData.SegmentPath = "diffServMaxRateEntry" + "[diffServMaxRateId='" + fmt.Sprintf("%v", diffservmaxrateentry.Diffservmaxrateid) + "']" + "[diffServMaxRateLevel='" + fmt.Sprintf("%v", diffservmaxrateentry.Diffservmaxratelevel) + "']"
    diffservmaxrateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservmaxrateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservmaxrateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservmaxrateentry.EntityData.Children = make(map[string]types.YChild)
    diffservmaxrateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservmaxrateentry.EntityData.Leafs["diffServMaxRateId"] = types.YLeaf{"Diffservmaxrateid", diffservmaxrateentry.Diffservmaxrateid}
    diffservmaxrateentry.EntityData.Leafs["diffServMaxRateLevel"] = types.YLeaf{"Diffservmaxratelevel", diffservmaxrateentry.Diffservmaxratelevel}
    diffservmaxrateentry.EntityData.Leafs["diffServMaxRateAbsolute"] = types.YLeaf{"Diffservmaxrateabsolute", diffservmaxrateentry.Diffservmaxrateabsolute}
    diffservmaxrateentry.EntityData.Leafs["diffServMaxRateRelative"] = types.YLeaf{"Diffservmaxraterelative", diffservmaxrateentry.Diffservmaxraterelative}
    diffservmaxrateentry.EntityData.Leafs["diffServMaxRateThreshold"] = types.YLeaf{"Diffservmaxratethreshold", diffservmaxrateentry.Diffservmaxratethreshold}
    diffservmaxrateentry.EntityData.Leafs["diffServMaxRateStorage"] = types.YLeaf{"Diffservmaxratestorage", diffservmaxrateentry.Diffservmaxratestorage}
    diffservmaxrateentry.EntityData.Leafs["diffServMaxRateStatus"] = types.YLeaf{"Diffservmaxratestatus", diffservmaxrateentry.Diffservmaxratestatus}
    return &(diffservmaxrateentry.EntityData)
}

