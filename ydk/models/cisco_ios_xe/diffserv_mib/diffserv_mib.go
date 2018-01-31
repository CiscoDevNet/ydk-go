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

type Diffservtbparamtrtcmblind struct {
}

func (id Diffservtbparamtrtcmblind) String() string {
	return "DIFFSERV-MIB:diffServTBParamTrTCMBlind"
}

type Diffservschedulerwfq struct {
}

func (id Diffservschedulerwfq) String() string {
	return "DIFFSERV-MIB:diffServSchedulerWFQ"
}

type Diffservtbparamtswtcm struct {
}

func (id Diffservtbparamtswtcm) String() string {
	return "DIFFSERV-MIB:diffServTBParamTswTCM"
}

type Diffservtbparamavgrate struct {
}

func (id Diffservtbparamavgrate) String() string {
	return "DIFFSERV-MIB:diffServTBParamAvgRate"
}

type Diffservschedulerwrr struct {
}

func (id Diffservschedulerwrr) String() string {
	return "DIFFSERV-MIB:diffServSchedulerWRR"
}

type Diffservtbparamsrtcmaware struct {
}

func (id Diffservtbparamsrtcmaware) String() string {
	return "DIFFSERV-MIB:diffServTBParamSrTCMAware"
}

type Diffservtbparamsrtcmblind struct {
}

func (id Diffservtbparamsrtcmblind) String() string {
	return "DIFFSERV-MIB:diffServTBParamSrTCMBlind"
}

type Diffservtbparamsimpletokenbucket struct {
}

func (id Diffservtbparamsimpletokenbucket) String() string {
	return "DIFFSERV-MIB:diffServTBParamSimpleTokenBucket"
}

type Diffservschedulerpriority struct {
}

func (id Diffservschedulerpriority) String() string {
	return "DIFFSERV-MIB:diffServSchedulerPriority"
}

type Diffservtbparamtrtcmaware struct {
}

func (id Diffservtbparamtrtcmaware) String() string {
	return "DIFFSERV-MIB:diffServTBParamTrTCMAware"
}

// IfDirection represents transmission on the interface.
type IfDirection string

const (
    IfDirection_inbound IfDirection = "inbound"

    IfDirection_outbound IfDirection = "outbound"
)

// DIFFSERVMIB
type DIFFSERVMIB struct {
    parent types.Entity
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

func (dIFFSERVMIB *DIFFSERVMIB) GetFilter() yfilter.YFilter { return dIFFSERVMIB.YFilter }

func (dIFFSERVMIB *DIFFSERVMIB) SetFilter(yf yfilter.YFilter) { dIFFSERVMIB.YFilter = yf }

func (dIFFSERVMIB *DIFFSERVMIB) GetGoName(yname string) string {
    if yname == "diffServClassifier" { return "Diffservclassifier" }
    if yname == "diffServMeter" { return "Diffservmeter" }
    if yname == "diffServTBParam" { return "Diffservtbparam" }
    if yname == "diffServAction" { return "Diffservaction" }
    if yname == "diffServAlgDrop" { return "Diffservalgdrop" }
    if yname == "diffServQueue" { return "Diffservqueue" }
    if yname == "diffServScheduler" { return "Diffservscheduler" }
    if yname == "diffServDataPathTable" { return "Diffservdatapathtable" }
    if yname == "diffServClfrTable" { return "Diffservclfrtable" }
    if yname == "diffServClfrElementTable" { return "Diffservclfrelementtable" }
    if yname == "diffServMultiFieldClfrTable" { return "Diffservmultifieldclfrtable" }
    if yname == "diffServMeterTable" { return "Diffservmetertable" }
    if yname == "diffServTBParamTable" { return "Diffservtbparamtable" }
    if yname == "diffServActionTable" { return "Diffservactiontable" }
    if yname == "diffServDscpMarkActTable" { return "Diffservdscpmarkacttable" }
    if yname == "diffServCountActTable" { return "Diffservcountacttable" }
    if yname == "diffServAlgDropTable" { return "Diffservalgdroptable" }
    if yname == "diffServRandomDropTable" { return "Diffservrandomdroptable" }
    if yname == "diffServQTable" { return "Diffservqtable" }
    if yname == "diffServSchedulerTable" { return "Diffservschedulertable" }
    if yname == "diffServMinRateTable" { return "Diffservminratetable" }
    if yname == "diffServMaxRateTable" { return "Diffservmaxratetable" }
    return ""
}

func (dIFFSERVMIB *DIFFSERVMIB) GetSegmentPath() string {
    return "DIFFSERV-MIB:DIFFSERV-MIB"
}

func (dIFFSERVMIB *DIFFSERVMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServClassifier" {
        return &dIFFSERVMIB.Diffservclassifier
    }
    if childYangName == "diffServMeter" {
        return &dIFFSERVMIB.Diffservmeter
    }
    if childYangName == "diffServTBParam" {
        return &dIFFSERVMIB.Diffservtbparam
    }
    if childYangName == "diffServAction" {
        return &dIFFSERVMIB.Diffservaction
    }
    if childYangName == "diffServAlgDrop" {
        return &dIFFSERVMIB.Diffservalgdrop
    }
    if childYangName == "diffServQueue" {
        return &dIFFSERVMIB.Diffservqueue
    }
    if childYangName == "diffServScheduler" {
        return &dIFFSERVMIB.Diffservscheduler
    }
    if childYangName == "diffServDataPathTable" {
        return &dIFFSERVMIB.Diffservdatapathtable
    }
    if childYangName == "diffServClfrTable" {
        return &dIFFSERVMIB.Diffservclfrtable
    }
    if childYangName == "diffServClfrElementTable" {
        return &dIFFSERVMIB.Diffservclfrelementtable
    }
    if childYangName == "diffServMultiFieldClfrTable" {
        return &dIFFSERVMIB.Diffservmultifieldclfrtable
    }
    if childYangName == "diffServMeterTable" {
        return &dIFFSERVMIB.Diffservmetertable
    }
    if childYangName == "diffServTBParamTable" {
        return &dIFFSERVMIB.Diffservtbparamtable
    }
    if childYangName == "diffServActionTable" {
        return &dIFFSERVMIB.Diffservactiontable
    }
    if childYangName == "diffServDscpMarkActTable" {
        return &dIFFSERVMIB.Diffservdscpmarkacttable
    }
    if childYangName == "diffServCountActTable" {
        return &dIFFSERVMIB.Diffservcountacttable
    }
    if childYangName == "diffServAlgDropTable" {
        return &dIFFSERVMIB.Diffservalgdroptable
    }
    if childYangName == "diffServRandomDropTable" {
        return &dIFFSERVMIB.Diffservrandomdroptable
    }
    if childYangName == "diffServQTable" {
        return &dIFFSERVMIB.Diffservqtable
    }
    if childYangName == "diffServSchedulerTable" {
        return &dIFFSERVMIB.Diffservschedulertable
    }
    if childYangName == "diffServMinRateTable" {
        return &dIFFSERVMIB.Diffservminratetable
    }
    if childYangName == "diffServMaxRateTable" {
        return &dIFFSERVMIB.Diffservmaxratetable
    }
    return nil
}

func (dIFFSERVMIB *DIFFSERVMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["diffServClassifier"] = &dIFFSERVMIB.Diffservclassifier
    children["diffServMeter"] = &dIFFSERVMIB.Diffservmeter
    children["diffServTBParam"] = &dIFFSERVMIB.Diffservtbparam
    children["diffServAction"] = &dIFFSERVMIB.Diffservaction
    children["diffServAlgDrop"] = &dIFFSERVMIB.Diffservalgdrop
    children["diffServQueue"] = &dIFFSERVMIB.Diffservqueue
    children["diffServScheduler"] = &dIFFSERVMIB.Diffservscheduler
    children["diffServDataPathTable"] = &dIFFSERVMIB.Diffservdatapathtable
    children["diffServClfrTable"] = &dIFFSERVMIB.Diffservclfrtable
    children["diffServClfrElementTable"] = &dIFFSERVMIB.Diffservclfrelementtable
    children["diffServMultiFieldClfrTable"] = &dIFFSERVMIB.Diffservmultifieldclfrtable
    children["diffServMeterTable"] = &dIFFSERVMIB.Diffservmetertable
    children["diffServTBParamTable"] = &dIFFSERVMIB.Diffservtbparamtable
    children["diffServActionTable"] = &dIFFSERVMIB.Diffservactiontable
    children["diffServDscpMarkActTable"] = &dIFFSERVMIB.Diffservdscpmarkacttable
    children["diffServCountActTable"] = &dIFFSERVMIB.Diffservcountacttable
    children["diffServAlgDropTable"] = &dIFFSERVMIB.Diffservalgdroptable
    children["diffServRandomDropTable"] = &dIFFSERVMIB.Diffservrandomdroptable
    children["diffServQTable"] = &dIFFSERVMIB.Diffservqtable
    children["diffServSchedulerTable"] = &dIFFSERVMIB.Diffservschedulertable
    children["diffServMinRateTable"] = &dIFFSERVMIB.Diffservminratetable
    children["diffServMaxRateTable"] = &dIFFSERVMIB.Diffservmaxratetable
    return children
}

func (dIFFSERVMIB *DIFFSERVMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dIFFSERVMIB *DIFFSERVMIB) GetBundleName() string { return "cisco_ios_xe" }

func (dIFFSERVMIB *DIFFSERVMIB) GetYangName() string { return "DIFFSERV-MIB" }

func (dIFFSERVMIB *DIFFSERVMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dIFFSERVMIB *DIFFSERVMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dIFFSERVMIB *DIFFSERVMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dIFFSERVMIB *DIFFSERVMIB) SetParent(parent types.Entity) { dIFFSERVMIB.parent = parent }

func (dIFFSERVMIB *DIFFSERVMIB) GetParent() types.Entity { return dIFFSERVMIB.parent }

func (dIFFSERVMIB *DIFFSERVMIB) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservclassifier
type DIFFSERVMIB_Diffservclassifier struct {
    parent types.Entity
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

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetFilter() yfilter.YFilter { return diffservclassifier.YFilter }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) SetFilter(yf yfilter.YFilter) { diffservclassifier.YFilter = yf }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetGoName(yname string) string {
    if yname == "diffServClfrNextFree" { return "Diffservclfrnextfree" }
    if yname == "diffServClfrElementNextFree" { return "Diffservclfrelementnextfree" }
    if yname == "diffServMultiFieldClfrNextFree" { return "Diffservmultifieldclfrnextfree" }
    return ""
}

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetSegmentPath() string {
    return "diffServClassifier"
}

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServClfrNextFree"] = diffservclassifier.Diffservclfrnextfree
    leafs["diffServClfrElementNextFree"] = diffservclassifier.Diffservclfrelementnextfree
    leafs["diffServMultiFieldClfrNextFree"] = diffservclassifier.Diffservmultifieldclfrnextfree
    return leafs
}

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetBundleName() string { return "cisco_ios_xe" }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetYangName() string { return "diffServClassifier" }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) SetParent(parent types.Entity) { diffservclassifier.parent = parent }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetParent() types.Entity { return diffservclassifier.parent }

func (diffservclassifier *DIFFSERVMIB_Diffservclassifier) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservmeter
type DIFFSERVMIB_Diffservmeter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServMeterId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservmeternextfree interface{}
}

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetFilter() yfilter.YFilter { return diffservmeter.YFilter }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) SetFilter(yf yfilter.YFilter) { diffservmeter.YFilter = yf }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetGoName(yname string) string {
    if yname == "diffServMeterNextFree" { return "Diffservmeternextfree" }
    return ""
}

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetSegmentPath() string {
    return "diffServMeter"
}

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServMeterNextFree"] = diffservmeter.Diffservmeternextfree
    return leafs
}

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetBundleName() string { return "cisco_ios_xe" }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetYangName() string { return "diffServMeter" }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) SetParent(parent types.Entity) { diffservmeter.parent = parent }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetParent() types.Entity { return diffservmeter.parent }

func (diffservmeter *DIFFSERVMIB_Diffservmeter) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservtbparam
type DIFFSERVMIB_Diffservtbparam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServTBParamId, or a zero to
    // indicate that none exist. The type is interface{} with range:
    // 0..4294967295.
    Diffservtbparamnextfree interface{}
}

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetFilter() yfilter.YFilter { return diffservtbparam.YFilter }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) SetFilter(yf yfilter.YFilter) { diffservtbparam.YFilter = yf }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetGoName(yname string) string {
    if yname == "diffServTBParamNextFree" { return "Diffservtbparamnextfree" }
    return ""
}

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetSegmentPath() string {
    return "diffServTBParam"
}

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServTBParamNextFree"] = diffservtbparam.Diffservtbparamnextfree
    return leafs
}

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetBundleName() string { return "cisco_ios_xe" }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetYangName() string { return "diffServTBParam" }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) SetParent(parent types.Entity) { diffservtbparam.parent = parent }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetParent() types.Entity { return diffservtbparam.parent }

func (diffservtbparam *DIFFSERVMIB_Diffservtbparam) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservaction
type DIFFSERVMIB_Diffservaction struct {
    parent types.Entity
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

func (diffservaction *DIFFSERVMIB_Diffservaction) GetFilter() yfilter.YFilter { return diffservaction.YFilter }

func (diffservaction *DIFFSERVMIB_Diffservaction) SetFilter(yf yfilter.YFilter) { diffservaction.YFilter = yf }

func (diffservaction *DIFFSERVMIB_Diffservaction) GetGoName(yname string) string {
    if yname == "diffServActionNextFree" { return "Diffservactionnextfree" }
    if yname == "diffServCountActNextFree" { return "Diffservcountactnextfree" }
    return ""
}

func (diffservaction *DIFFSERVMIB_Diffservaction) GetSegmentPath() string {
    return "diffServAction"
}

func (diffservaction *DIFFSERVMIB_Diffservaction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservaction *DIFFSERVMIB_Diffservaction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservaction *DIFFSERVMIB_Diffservaction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServActionNextFree"] = diffservaction.Diffservactionnextfree
    leafs["diffServCountActNextFree"] = diffservaction.Diffservcountactnextfree
    return leafs
}

func (diffservaction *DIFFSERVMIB_Diffservaction) GetBundleName() string { return "cisco_ios_xe" }

func (diffservaction *DIFFSERVMIB_Diffservaction) GetYangName() string { return "diffServAction" }

func (diffservaction *DIFFSERVMIB_Diffservaction) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservaction *DIFFSERVMIB_Diffservaction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservaction *DIFFSERVMIB_Diffservaction) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservaction *DIFFSERVMIB_Diffservaction) SetParent(parent types.Entity) { diffservaction.parent = parent }

func (diffservaction *DIFFSERVMIB_Diffservaction) GetParent() types.Entity { return diffservaction.parent }

func (diffservaction *DIFFSERVMIB_Diffservaction) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservalgdrop
type DIFFSERVMIB_Diffservalgdrop struct {
    parent types.Entity
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

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetFilter() yfilter.YFilter { return diffservalgdrop.YFilter }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) SetFilter(yf yfilter.YFilter) { diffservalgdrop.YFilter = yf }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetGoName(yname string) string {
    if yname == "diffServAlgDropNextFree" { return "Diffservalgdropnextfree" }
    if yname == "diffServRandomDropNextFree" { return "Diffservrandomdropnextfree" }
    return ""
}

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetSegmentPath() string {
    return "diffServAlgDrop"
}

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServAlgDropNextFree"] = diffservalgdrop.Diffservalgdropnextfree
    leafs["diffServRandomDropNextFree"] = diffservalgdrop.Diffservrandomdropnextfree
    return leafs
}

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetBundleName() string { return "cisco_ios_xe" }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetYangName() string { return "diffServAlgDrop" }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) SetParent(parent types.Entity) { diffservalgdrop.parent = parent }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetParent() types.Entity { return diffservalgdrop.parent }

func (diffservalgdrop *DIFFSERVMIB_Diffservalgdrop) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservqueue
type DIFFSERVMIB_Diffservqueue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object contains an unused value for diffServQId, or a zero to indicate
    // that none exist. The type is interface{} with range: 0..4294967295.
    Diffservqnextfree interface{}
}

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetFilter() yfilter.YFilter { return diffservqueue.YFilter }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) SetFilter(yf yfilter.YFilter) { diffservqueue.YFilter = yf }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetGoName(yname string) string {
    if yname == "diffServQNextFree" { return "Diffservqnextfree" }
    return ""
}

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetSegmentPath() string {
    return "diffServQueue"
}

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServQNextFree"] = diffservqueue.Diffservqnextfree
    return leafs
}

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetBundleName() string { return "cisco_ios_xe" }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetYangName() string { return "diffServQueue" }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) SetParent(parent types.Entity) { diffservqueue.parent = parent }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetParent() types.Entity { return diffservqueue.parent }

func (diffservqueue *DIFFSERVMIB_Diffservqueue) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservscheduler
type DIFFSERVMIB_Diffservscheduler struct {
    parent types.Entity
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

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetFilter() yfilter.YFilter { return diffservscheduler.YFilter }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) SetFilter(yf yfilter.YFilter) { diffservscheduler.YFilter = yf }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetGoName(yname string) string {
    if yname == "diffServSchedulerNextFree" { return "Diffservschedulernextfree" }
    if yname == "diffServMinRateNextFree" { return "Diffservminratenextfree" }
    if yname == "diffServMaxRateNextFree" { return "Diffservmaxratenextfree" }
    return ""
}

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetSegmentPath() string {
    return "diffServScheduler"
}

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServSchedulerNextFree"] = diffservscheduler.Diffservschedulernextfree
    leafs["diffServMinRateNextFree"] = diffservscheduler.Diffservminratenextfree
    leafs["diffServMaxRateNextFree"] = diffservscheduler.Diffservmaxratenextfree
    return leafs
}

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetBundleName() string { return "cisco_ios_xe" }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetYangName() string { return "diffServScheduler" }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) SetParent(parent types.Entity) { diffservscheduler.parent = parent }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetParent() types.Entity { return diffservscheduler.parent }

func (diffservscheduler *DIFFSERVMIB_Diffservscheduler) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservdatapathtable
// The data path table contains RowPointers indicating the start of
// the functional data path for each interface and traffic direction
// in this device. These may merge, or be separated into parallel
// data paths.
type DIFFSERVMIB_Diffservdatapathtable struct {
    parent types.Entity
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

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetFilter() yfilter.YFilter { return diffservdatapathtable.YFilter }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) SetFilter(yf yfilter.YFilter) { diffservdatapathtable.YFilter = yf }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetGoName(yname string) string {
    if yname == "diffServDataPathEntry" { return "Diffservdatapathentry" }
    return ""
}

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetSegmentPath() string {
    return "diffServDataPathTable"
}

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServDataPathEntry" {
        for _, c := range diffservdatapathtable.Diffservdatapathentry {
            if diffservdatapathtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry{}
        diffservdatapathtable.Diffservdatapathentry = append(diffservdatapathtable.Diffservdatapathentry, child)
        return &diffservdatapathtable.Diffservdatapathentry[len(diffservdatapathtable.Diffservdatapathentry)-1]
    }
    return nil
}

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservdatapathtable.Diffservdatapathentry {
        children[diffservdatapathtable.Diffservdatapathentry[i].GetSegmentPath()] = &diffservdatapathtable.Diffservdatapathentry[i]
    }
    return children
}

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetYangName() string { return "diffServDataPathTable" }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) SetParent(parent types.Entity) { diffservdatapathtable.parent = parent }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetParent() types.Entity { return diffservdatapathtable.parent }

func (diffservdatapathtable *DIFFSERVMIB_Diffservdatapathtable) GetParentYangName() string { return "DIFFSERV-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Diffservdatapathstart interface{}

    // The storage type for this conceptual row.  Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    Diffservdatapathstorage interface{}

    // The status of this conceptual row. All writable objects in this row may be
    // modified at any time. The type is RowStatus.
    Diffservdatapathstatus interface{}
}

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetFilter() yfilter.YFilter { return diffservdatapathentry.YFilter }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) SetFilter(yf yfilter.YFilter) { diffservdatapathentry.YFilter = yf }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "diffServDataPathIfDirection" { return "Diffservdatapathifdirection" }
    if yname == "diffServDataPathStart" { return "Diffservdatapathstart" }
    if yname == "diffServDataPathStorage" { return "Diffservdatapathstorage" }
    if yname == "diffServDataPathStatus" { return "Diffservdatapathstatus" }
    return ""
}

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetSegmentPath() string {
    return "diffServDataPathEntry" + "[ifIndex='" + fmt.Sprintf("%v", diffservdatapathentry.Ifindex) + "']" + "[diffServDataPathIfDirection='" + fmt.Sprintf("%v", diffservdatapathentry.Diffservdatapathifdirection) + "']"
}

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = diffservdatapathentry.Ifindex
    leafs["diffServDataPathIfDirection"] = diffservdatapathentry.Diffservdatapathifdirection
    leafs["diffServDataPathStart"] = diffservdatapathentry.Diffservdatapathstart
    leafs["diffServDataPathStorage"] = diffservdatapathentry.Diffservdatapathstorage
    leafs["diffServDataPathStatus"] = diffservdatapathentry.Diffservdatapathstatus
    return leafs
}

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetYangName() string { return "diffServDataPathEntry" }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) SetParent(parent types.Entity) { diffservdatapathentry.parent = parent }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetParent() types.Entity { return diffservdatapathentry.parent }

func (diffservdatapathentry *DIFFSERVMIB_Diffservdatapathtable_Diffservdatapathentry) GetParentYangName() string { return "diffServDataPathTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the classifier table describes a single classifier. All
    // classifier elements belonging to the same classifier use the classifier's
    // diffServClfrId as part of their index. The type is slice of
    // DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry.
    Diffservclfrentry []DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry
}

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetFilter() yfilter.YFilter { return diffservclfrtable.YFilter }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) SetFilter(yf yfilter.YFilter) { diffservclfrtable.YFilter = yf }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetGoName(yname string) string {
    if yname == "diffServClfrEntry" { return "Diffservclfrentry" }
    return ""
}

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetSegmentPath() string {
    return "diffServClfrTable"
}

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServClfrEntry" {
        for _, c := range diffservclfrtable.Diffservclfrentry {
            if diffservclfrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry{}
        diffservclfrtable.Diffservclfrentry = append(diffservclfrtable.Diffservclfrentry, child)
        return &diffservclfrtable.Diffservclfrentry[len(diffservclfrtable.Diffservclfrentry)-1]
    }
    return nil
}

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservclfrtable.Diffservclfrentry {
        children[diffservclfrtable.Diffservclfrentry[i].GetSegmentPath()] = &diffservclfrtable.Diffservclfrentry[i]
    }
    return children
}

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetYangName() string { return "diffServClfrTable" }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) SetParent(parent types.Entity) { diffservclfrtable.parent = parent }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetParent() types.Entity { return diffservclfrtable.parent }

func (diffservclfrtable *DIFFSERVMIB_Diffservclfrtable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry
// An entry in the classifier table describes a single classifier.
// All classifier elements belonging to the same classifier use the
// classifier's diffServClfrId as part of their index.
type DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry struct {
    parent types.Entity
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

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetFilter() yfilter.YFilter { return diffservclfrentry.YFilter }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) SetFilter(yf yfilter.YFilter) { diffservclfrentry.YFilter = yf }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetGoName(yname string) string {
    if yname == "diffServClfrId" { return "Diffservclfrid" }
    if yname == "diffServClfrStorage" { return "Diffservclfrstorage" }
    if yname == "diffServClfrStatus" { return "Diffservclfrstatus" }
    return ""
}

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetSegmentPath() string {
    return "diffServClfrEntry" + "[diffServClfrId='" + fmt.Sprintf("%v", diffservclfrentry.Diffservclfrid) + "']"
}

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServClfrId"] = diffservclfrentry.Diffservclfrid
    leafs["diffServClfrStorage"] = diffservclfrentry.Diffservclfrstorage
    leafs["diffServClfrStatus"] = diffservclfrentry.Diffservclfrstatus
    return leafs
}

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetYangName() string { return "diffServClfrEntry" }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) SetParent(parent types.Entity) { diffservclfrentry.parent = parent }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetParent() types.Entity { return diffservclfrentry.parent }

func (diffservclfrentry *DIFFSERVMIB_Diffservclfrtable_Diffservclfrentry) GetParentYangName() string { return "diffServClfrTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the classifier element table describes a single element of the
    // classifier. The type is slice of
    // DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry.
    Diffservclfrelemententry []DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry
}

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetFilter() yfilter.YFilter { return diffservclfrelementtable.YFilter }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) SetFilter(yf yfilter.YFilter) { diffservclfrelementtable.YFilter = yf }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetGoName(yname string) string {
    if yname == "diffServClfrElementEntry" { return "Diffservclfrelemententry" }
    return ""
}

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetSegmentPath() string {
    return "diffServClfrElementTable"
}

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServClfrElementEntry" {
        for _, c := range diffservclfrelementtable.Diffservclfrelemententry {
            if diffservclfrelementtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry{}
        diffservclfrelementtable.Diffservclfrelemententry = append(diffservclfrelementtable.Diffservclfrelemententry, child)
        return &diffservclfrelementtable.Diffservclfrelemententry[len(diffservclfrelementtable.Diffservclfrelemententry)-1]
    }
    return nil
}

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservclfrelementtable.Diffservclfrelemententry {
        children[diffservclfrelementtable.Diffservclfrelemententry[i].GetSegmentPath()] = &diffservclfrelementtable.Diffservclfrelemententry[i]
    }
    return children
}

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetYangName() string { return "diffServClfrElementTable" }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) SetParent(parent types.Entity) { diffservclfrelementtable.parent = parent }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetParent() types.Entity { return diffservclfrelementtable.parent }

func (diffservclfrelementtable *DIFFSERVMIB_Diffservclfrelementtable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry
// An entry in the classifier element table describes a single
// element of the classifier.
type DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Diffservclfrelementnext interface{}

    // A pointer to a valid entry in another table, filter table, that describes
    // the applicable classification parameters, e.g. an entry in
    // diffServMultiFieldClfrTable.  The value zeroDotZero is interpreted to match
    // anything not matched by another classifier element - only one such entry
    // may exist for each classifier.  Setting this to point to a target that does
    // not exist results in an inconsistentValue error.  If the row pointed to is
    // removed or    becomes inactive by other means, the element is ignored. The
    // type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetFilter() yfilter.YFilter { return diffservclfrelemententry.YFilter }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) SetFilter(yf yfilter.YFilter) { diffservclfrelemententry.YFilter = yf }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetGoName(yname string) string {
    if yname == "diffServClfrId" { return "Diffservclfrid" }
    if yname == "diffServClfrElementId" { return "Diffservclfrelementid" }
    if yname == "diffServClfrElementPrecedence" { return "Diffservclfrelementprecedence" }
    if yname == "diffServClfrElementNext" { return "Diffservclfrelementnext" }
    if yname == "diffServClfrElementSpecific" { return "Diffservclfrelementspecific" }
    if yname == "diffServClfrElementStorage" { return "Diffservclfrelementstorage" }
    if yname == "diffServClfrElementStatus" { return "Diffservclfrelementstatus" }
    return ""
}

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetSegmentPath() string {
    return "diffServClfrElementEntry" + "[diffServClfrId='" + fmt.Sprintf("%v", diffservclfrelemententry.Diffservclfrid) + "']" + "[diffServClfrElementId='" + fmt.Sprintf("%v", diffservclfrelemententry.Diffservclfrelementid) + "']"
}

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServClfrId"] = diffservclfrelemententry.Diffservclfrid
    leafs["diffServClfrElementId"] = diffservclfrelemententry.Diffservclfrelementid
    leafs["diffServClfrElementPrecedence"] = diffservclfrelemententry.Diffservclfrelementprecedence
    leafs["diffServClfrElementNext"] = diffservclfrelemententry.Diffservclfrelementnext
    leafs["diffServClfrElementSpecific"] = diffservclfrelemententry.Diffservclfrelementspecific
    leafs["diffServClfrElementStorage"] = diffservclfrelemententry.Diffservclfrelementstorage
    leafs["diffServClfrElementStatus"] = diffservclfrelemententry.Diffservclfrelementstatus
    return leafs
}

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetYangName() string { return "diffServClfrElementEntry" }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) SetParent(parent types.Entity) { diffservclfrelemententry.parent = parent }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetParent() types.Entity { return diffservclfrelemententry.parent }

func (diffservclfrelemententry *DIFFSERVMIB_Diffservclfrelementtable_Diffservclfrelemententry) GetParentYangName() string { return "diffServClfrElementTable" }

// DIFFSERVMIB_Diffservmultifieldclfrtable
// A table of IP Multi-field Classifier filter entries that a
// 
// 
// 
// system may use to identify IP traffic.
type DIFFSERVMIB_Diffservmultifieldclfrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An IP Multi-field Classifier entry describes a single filter. The type is
    // slice of
    // DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry.
    Diffservmultifieldclfrentry []DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry
}

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetFilter() yfilter.YFilter { return diffservmultifieldclfrtable.YFilter }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) SetFilter(yf yfilter.YFilter) { diffservmultifieldclfrtable.YFilter = yf }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetGoName(yname string) string {
    if yname == "diffServMultiFieldClfrEntry" { return "Diffservmultifieldclfrentry" }
    return ""
}

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetSegmentPath() string {
    return "diffServMultiFieldClfrTable"
}

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServMultiFieldClfrEntry" {
        for _, c := range diffservmultifieldclfrtable.Diffservmultifieldclfrentry {
            if diffservmultifieldclfrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry{}
        diffservmultifieldclfrtable.Diffservmultifieldclfrentry = append(diffservmultifieldclfrtable.Diffservmultifieldclfrentry, child)
        return &diffservmultifieldclfrtable.Diffservmultifieldclfrentry[len(diffservmultifieldclfrtable.Diffservmultifieldclfrentry)-1]
    }
    return nil
}

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservmultifieldclfrtable.Diffservmultifieldclfrentry {
        children[diffservmultifieldclfrtable.Diffservmultifieldclfrentry[i].GetSegmentPath()] = &diffservmultifieldclfrtable.Diffservmultifieldclfrentry[i]
    }
    return children
}

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetYangName() string { return "diffServMultiFieldClfrTable" }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) SetParent(parent types.Entity) { diffservmultifieldclfrtable.parent = parent }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetParent() types.Entity { return diffservmultifieldclfrtable.parent }

func (diffservmultifieldclfrtable *DIFFSERVMIB_Diffservmultifieldclfrtable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry
// An IP Multi-field Classifier entry describes a single filter.
type DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry struct {
    parent types.Entity
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

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetFilter() yfilter.YFilter { return diffservmultifieldclfrentry.YFilter }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) SetFilter(yf yfilter.YFilter) { diffservmultifieldclfrentry.YFilter = yf }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetGoName(yname string) string {
    if yname == "diffServMultiFieldClfrId" { return "Diffservmultifieldclfrid" }
    if yname == "diffServMultiFieldClfrAddrType" { return "Diffservmultifieldclfraddrtype" }
    if yname == "diffServMultiFieldClfrDstAddr" { return "Diffservmultifieldclfrdstaddr" }
    if yname == "diffServMultiFieldClfrDstPrefixLength" { return "Diffservmultifieldclfrdstprefixlength" }
    if yname == "diffServMultiFieldClfrSrcAddr" { return "Diffservmultifieldclfrsrcaddr" }
    if yname == "diffServMultiFieldClfrSrcPrefixLength" { return "Diffservmultifieldclfrsrcprefixlength" }
    if yname == "diffServMultiFieldClfrDscp" { return "Diffservmultifieldclfrdscp" }
    if yname == "diffServMultiFieldClfrFlowId" { return "Diffservmultifieldclfrflowid" }
    if yname == "diffServMultiFieldClfrProtocol" { return "Diffservmultifieldclfrprotocol" }
    if yname == "diffServMultiFieldClfrDstL4PortMin" { return "Diffservmultifieldclfrdstl4Portmin" }
    if yname == "diffServMultiFieldClfrDstL4PortMax" { return "Diffservmultifieldclfrdstl4Portmax" }
    if yname == "diffServMultiFieldClfrSrcL4PortMin" { return "Diffservmultifieldclfrsrcl4Portmin" }
    if yname == "diffServMultiFieldClfrSrcL4PortMax" { return "Diffservmultifieldclfrsrcl4Portmax" }
    if yname == "diffServMultiFieldClfrStorage" { return "Diffservmultifieldclfrstorage" }
    if yname == "diffServMultiFieldClfrStatus" { return "Diffservmultifieldclfrstatus" }
    return ""
}

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetSegmentPath() string {
    return "diffServMultiFieldClfrEntry" + "[diffServMultiFieldClfrId='" + fmt.Sprintf("%v", diffservmultifieldclfrentry.Diffservmultifieldclfrid) + "']"
}

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServMultiFieldClfrId"] = diffservmultifieldclfrentry.Diffservmultifieldclfrid
    leafs["diffServMultiFieldClfrAddrType"] = diffservmultifieldclfrentry.Diffservmultifieldclfraddrtype
    leafs["diffServMultiFieldClfrDstAddr"] = diffservmultifieldclfrentry.Diffservmultifieldclfrdstaddr
    leafs["diffServMultiFieldClfrDstPrefixLength"] = diffservmultifieldclfrentry.Diffservmultifieldclfrdstprefixlength
    leafs["diffServMultiFieldClfrSrcAddr"] = diffservmultifieldclfrentry.Diffservmultifieldclfrsrcaddr
    leafs["diffServMultiFieldClfrSrcPrefixLength"] = diffservmultifieldclfrentry.Diffservmultifieldclfrsrcprefixlength
    leafs["diffServMultiFieldClfrDscp"] = diffservmultifieldclfrentry.Diffservmultifieldclfrdscp
    leafs["diffServMultiFieldClfrFlowId"] = diffservmultifieldclfrentry.Diffservmultifieldclfrflowid
    leafs["diffServMultiFieldClfrProtocol"] = diffservmultifieldclfrentry.Diffservmultifieldclfrprotocol
    leafs["diffServMultiFieldClfrDstL4PortMin"] = diffservmultifieldclfrentry.Diffservmultifieldclfrdstl4Portmin
    leafs["diffServMultiFieldClfrDstL4PortMax"] = diffservmultifieldclfrentry.Diffservmultifieldclfrdstl4Portmax
    leafs["diffServMultiFieldClfrSrcL4PortMin"] = diffservmultifieldclfrentry.Diffservmultifieldclfrsrcl4Portmin
    leafs["diffServMultiFieldClfrSrcL4PortMax"] = diffservmultifieldclfrentry.Diffservmultifieldclfrsrcl4Portmax
    leafs["diffServMultiFieldClfrStorage"] = diffservmultifieldclfrentry.Diffservmultifieldclfrstorage
    leafs["diffServMultiFieldClfrStatus"] = diffservmultifieldclfrentry.Diffservmultifieldclfrstatus
    return leafs
}

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetYangName() string { return "diffServMultiFieldClfrEntry" }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) SetParent(parent types.Entity) { diffservmultifieldclfrentry.parent = parent }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetParent() types.Entity { return diffservmultifieldclfrentry.parent }

func (diffservmultifieldclfrentry *DIFFSERVMIB_Diffservmultifieldclfrtable_Diffservmultifieldclfrentry) GetParentYangName() string { return "diffServMultiFieldClfrTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the meter table describes a single conformance level of a
    // meter. The type is slice of
    // DIFFSERVMIB_Diffservmetertable_Diffservmeterentry.
    Diffservmeterentry []DIFFSERVMIB_Diffservmetertable_Diffservmeterentry
}

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetFilter() yfilter.YFilter { return diffservmetertable.YFilter }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) SetFilter(yf yfilter.YFilter) { diffservmetertable.YFilter = yf }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetGoName(yname string) string {
    if yname == "diffServMeterEntry" { return "Diffservmeterentry" }
    return ""
}

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetSegmentPath() string {
    return "diffServMeterTable"
}

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServMeterEntry" {
        for _, c := range diffservmetertable.Diffservmeterentry {
            if diffservmetertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservmetertable_Diffservmeterentry{}
        diffservmetertable.Diffservmeterentry = append(diffservmetertable.Diffservmeterentry, child)
        return &diffservmetertable.Diffservmeterentry[len(diffservmetertable.Diffservmeterentry)-1]
    }
    return nil
}

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservmetertable.Diffservmeterentry {
        children[diffservmetertable.Diffservmeterentry[i].GetSegmentPath()] = &diffservmetertable.Diffservmeterentry[i]
    }
    return children
}

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetYangName() string { return "diffServMeterTable" }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) SetParent(parent types.Entity) { diffservmetertable.parent = parent }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetParent() types.Entity { return diffservmetertable.parent }

func (diffservmetertable *DIFFSERVMIB_Diffservmetertable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservmetertable_Diffservmeterentry
// An entry in the meter table describes a single conformance level
// of a meter.
type DIFFSERVMIB_Diffservmetertable_Diffservmeterentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Diffservmeterfailnext interface{}

    // This indicates the behavior of the meter by pointing to an entry containing
    // detailed parameters. Note that entries in that specific table must be
    // managed explicitly.  For example, diffServMeterSpecific may point to an
    // entry in diffServTBParamTable, which contains an instance of a single set
    // of Token Bucket parameters.  Setting this to point to a target that does
    // not exist results in an inconsistentValue error.  If the row pointed to is
    // removed or becomes inactive by other means, the meter always succeeds. The
    // type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetFilter() yfilter.YFilter { return diffservmeterentry.YFilter }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) SetFilter(yf yfilter.YFilter) { diffservmeterentry.YFilter = yf }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetGoName(yname string) string {
    if yname == "diffServMeterId" { return "Diffservmeterid" }
    if yname == "diffServMeterSucceedNext" { return "Diffservmetersucceednext" }
    if yname == "diffServMeterFailNext" { return "Diffservmeterfailnext" }
    if yname == "diffServMeterSpecific" { return "Diffservmeterspecific" }
    if yname == "diffServMeterStorage" { return "Diffservmeterstorage" }
    if yname == "diffServMeterStatus" { return "Diffservmeterstatus" }
    return ""
}

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetSegmentPath() string {
    return "diffServMeterEntry" + "[diffServMeterId='" + fmt.Sprintf("%v", diffservmeterentry.Diffservmeterid) + "']"
}

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServMeterId"] = diffservmeterentry.Diffservmeterid
    leafs["diffServMeterSucceedNext"] = diffservmeterentry.Diffservmetersucceednext
    leafs["diffServMeterFailNext"] = diffservmeterentry.Diffservmeterfailnext
    leafs["diffServMeterSpecific"] = diffservmeterentry.Diffservmeterspecific
    leafs["diffServMeterStorage"] = diffservmeterentry.Diffservmeterstorage
    leafs["diffServMeterStatus"] = diffservmeterentry.Diffservmeterstatus
    return leafs
}

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetYangName() string { return "diffServMeterEntry" }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) SetParent(parent types.Entity) { diffservmeterentry.parent = parent }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetParent() types.Entity { return diffservmeterentry.parent }

func (diffservmeterentry *DIFFSERVMIB_Diffservmetertable_Diffservmeterentry) GetParentYangName() string { return "diffServMeterTable" }

// DIFFSERVMIB_Diffservtbparamtable
// This table enumerates a single set of token bucket meter
// parameters that a system may use to police a stream of traffic.
// Such meters are modeled here as having a single rate and a single
// burst size. Multiple entries are used when multiple rates/burst
// sizes are needed.
type DIFFSERVMIB_Diffservtbparamtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry that describes a single set of token bucket parameters. The type
    // is slice of DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry.
    Diffservtbparamentry []DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry
}

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetFilter() yfilter.YFilter { return diffservtbparamtable.YFilter }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) SetFilter(yf yfilter.YFilter) { diffservtbparamtable.YFilter = yf }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetGoName(yname string) string {
    if yname == "diffServTBParamEntry" { return "Diffservtbparamentry" }
    return ""
}

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetSegmentPath() string {
    return "diffServTBParamTable"
}

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServTBParamEntry" {
        for _, c := range diffservtbparamtable.Diffservtbparamentry {
            if diffservtbparamtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry{}
        diffservtbparamtable.Diffservtbparamentry = append(diffservtbparamtable.Diffservtbparamentry, child)
        return &diffservtbparamtable.Diffservtbparamentry[len(diffservtbparamtable.Diffservtbparamentry)-1]
    }
    return nil
}

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservtbparamtable.Diffservtbparamentry {
        children[diffservtbparamtable.Diffservtbparamentry[i].GetSegmentPath()] = &diffservtbparamtable.Diffservtbparamentry[i]
    }
    return children
}

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetYangName() string { return "diffServTBParamTable" }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) SetParent(parent types.Entity) { diffservtbparamtable.parent = parent }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetParent() types.Entity { return diffservtbparamtable.parent }

func (diffservtbparamtable *DIFFSERVMIB_Diffservtbparamtable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry
// An entry that describes a single set of token bucket
// parameters.
type DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetFilter() yfilter.YFilter { return diffservtbparamentry.YFilter }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) SetFilter(yf yfilter.YFilter) { diffservtbparamentry.YFilter = yf }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetGoName(yname string) string {
    if yname == "diffServTBParamId" { return "Diffservtbparamid" }
    if yname == "diffServTBParamType" { return "Diffservtbparamtype" }
    if yname == "diffServTBParamRate" { return "Diffservtbparamrate" }
    if yname == "diffServTBParamBurstSize" { return "Diffservtbparamburstsize" }
    if yname == "diffServTBParamInterval" { return "Diffservtbparaminterval" }
    if yname == "diffServTBParamStorage" { return "Diffservtbparamstorage" }
    if yname == "diffServTBParamStatus" { return "Diffservtbparamstatus" }
    return ""
}

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetSegmentPath() string {
    return "diffServTBParamEntry" + "[diffServTBParamId='" + fmt.Sprintf("%v", diffservtbparamentry.Diffservtbparamid) + "']"
}

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServTBParamId"] = diffservtbparamentry.Diffservtbparamid
    leafs["diffServTBParamType"] = diffservtbparamentry.Diffservtbparamtype
    leafs["diffServTBParamRate"] = diffservtbparamentry.Diffservtbparamrate
    leafs["diffServTBParamBurstSize"] = diffservtbparamentry.Diffservtbparamburstsize
    leafs["diffServTBParamInterval"] = diffservtbparamentry.Diffservtbparaminterval
    leafs["diffServTBParamStorage"] = diffservtbparamentry.Diffservtbparamstorage
    leafs["diffServTBParamStatus"] = diffservtbparamentry.Diffservtbparamstatus
    return leafs
}

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetYangName() string { return "diffServTBParamEntry" }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) SetParent(parent types.Entity) { diffservtbparamentry.parent = parent }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetParent() types.Entity { return diffservtbparamentry.parent }

func (diffservtbparamentry *DIFFSERVMIB_Diffservtbparamtable_Diffservtbparamentry) GetParentYangName() string { return "diffServTBParamTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in the action table allows description of one specific action to
    // be applied to traffic. The type is slice of
    // DIFFSERVMIB_Diffservactiontable_Diffservactionentry.
    Diffservactionentry []DIFFSERVMIB_Diffservactiontable_Diffservactionentry
}

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetFilter() yfilter.YFilter { return diffservactiontable.YFilter }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) SetFilter(yf yfilter.YFilter) { diffservactiontable.YFilter = yf }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetGoName(yname string) string {
    if yname == "diffServActionEntry" { return "Diffservactionentry" }
    return ""
}

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetSegmentPath() string {
    return "diffServActionTable"
}

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServActionEntry" {
        for _, c := range diffservactiontable.Diffservactionentry {
            if diffservactiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservactiontable_Diffservactionentry{}
        diffservactiontable.Diffservactionentry = append(diffservactiontable.Diffservactionentry, child)
        return &diffservactiontable.Diffservactionentry[len(diffservactiontable.Diffservactionentry)-1]
    }
    return nil
}

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservactiontable.Diffservactionentry {
        children[diffservactiontable.Diffservactionentry[i].GetSegmentPath()] = &diffservactiontable.Diffservactionentry[i]
    }
    return children
}

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetYangName() string { return "diffServActionTable" }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) SetParent(parent types.Entity) { diffservactiontable.parent = parent }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetParent() types.Entity { return diffservactiontable.parent }

func (diffservactiontable *DIFFSERVMIB_Diffservactiontable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservactiontable_Diffservactionentry
// Each entry in the action table allows description of one
// specific action to be applied to traffic.
type DIFFSERVMIB_Diffservactiontable_Diffservactionentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetFilter() yfilter.YFilter { return diffservactionentry.YFilter }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) SetFilter(yf yfilter.YFilter) { diffservactionentry.YFilter = yf }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetGoName(yname string) string {
    if yname == "diffServActionId" { return "Diffservactionid" }
    if yname == "diffServActionInterface" { return "Diffservactioninterface" }
    if yname == "diffServActionNext" { return "Diffservactionnext" }
    if yname == "diffServActionSpecific" { return "Diffservactionspecific" }
    if yname == "diffServActionStorage" { return "Diffservactionstorage" }
    if yname == "diffServActionStatus" { return "Diffservactionstatus" }
    return ""
}

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetSegmentPath() string {
    return "diffServActionEntry" + "[diffServActionId='" + fmt.Sprintf("%v", diffservactionentry.Diffservactionid) + "']"
}

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServActionId"] = diffservactionentry.Diffservactionid
    leafs["diffServActionInterface"] = diffservactionentry.Diffservactioninterface
    leafs["diffServActionNext"] = diffservactionentry.Diffservactionnext
    leafs["diffServActionSpecific"] = diffservactionentry.Diffservactionspecific
    leafs["diffServActionStorage"] = diffservactionentry.Diffservactionstorage
    leafs["diffServActionStatus"] = diffservactionentry.Diffservactionstatus
    return leafs
}

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetYangName() string { return "diffServActionEntry" }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) SetParent(parent types.Entity) { diffservactionentry.parent = parent }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetParent() types.Entity { return diffservactionentry.parent }

func (diffservactionentry *DIFFSERVMIB_Diffservactiontable_Diffservactionentry) GetParentYangName() string { return "diffServActionTable" }

// DIFFSERVMIB_Diffservdscpmarkacttable
// This table enumerates specific DSCPs used for marking or
// remarking the DSCP field of IP packets. The entries of this table
// may be referenced by a diffServActionSpecific attribute.
type DIFFSERVMIB_Diffservdscpmarkacttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the DSCP mark action table that describes a single DSCP used
    // for marking. The type is slice of
    // DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry.
    Diffservdscpmarkactentry []DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry
}

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetFilter() yfilter.YFilter { return diffservdscpmarkacttable.YFilter }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) SetFilter(yf yfilter.YFilter) { diffservdscpmarkacttable.YFilter = yf }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetGoName(yname string) string {
    if yname == "diffServDscpMarkActEntry" { return "Diffservdscpmarkactentry" }
    return ""
}

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetSegmentPath() string {
    return "diffServDscpMarkActTable"
}

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServDscpMarkActEntry" {
        for _, c := range diffservdscpmarkacttable.Diffservdscpmarkactentry {
            if diffservdscpmarkacttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry{}
        diffservdscpmarkacttable.Diffservdscpmarkactentry = append(diffservdscpmarkacttable.Diffservdscpmarkactentry, child)
        return &diffservdscpmarkacttable.Diffservdscpmarkactentry[len(diffservdscpmarkacttable.Diffservdscpmarkactentry)-1]
    }
    return nil
}

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservdscpmarkacttable.Diffservdscpmarkactentry {
        children[diffservdscpmarkacttable.Diffservdscpmarkactentry[i].GetSegmentPath()] = &diffservdscpmarkacttable.Diffservdscpmarkactentry[i]
    }
    return children
}

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetYangName() string { return "diffServDscpMarkActTable" }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) SetParent(parent types.Entity) { diffservdscpmarkacttable.parent = parent }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetParent() types.Entity { return diffservdscpmarkacttable.parent }

func (diffservdscpmarkacttable *DIFFSERVMIB_Diffservdscpmarkacttable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry
// An entry in the DSCP mark action table that describes a single
// DSCP used for marking.
type DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry struct {
    parent types.Entity
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

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetFilter() yfilter.YFilter { return diffservdscpmarkactentry.YFilter }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) SetFilter(yf yfilter.YFilter) { diffservdscpmarkactentry.YFilter = yf }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetGoName(yname string) string {
    if yname == "diffServDscpMarkActDscp" { return "Diffservdscpmarkactdscp" }
    return ""
}

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetSegmentPath() string {
    return "diffServDscpMarkActEntry" + "[diffServDscpMarkActDscp='" + fmt.Sprintf("%v", diffservdscpmarkactentry.Diffservdscpmarkactdscp) + "']"
}

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServDscpMarkActDscp"] = diffservdscpmarkactentry.Diffservdscpmarkactdscp
    return leafs
}

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetYangName() string { return "diffServDscpMarkActEntry" }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) SetParent(parent types.Entity) { diffservdscpmarkactentry.parent = parent }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetParent() types.Entity { return diffservdscpmarkactentry.parent }

func (diffservdscpmarkactentry *DIFFSERVMIB_Diffservdscpmarkacttable_Diffservdscpmarkactentry) GetParentYangName() string { return "diffServDscpMarkActTable" }

// DIFFSERVMIB_Diffservcountacttable
// This table contains counters for all the traffic passing through
// an action element.
type DIFFSERVMIB_Diffservcountacttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the count action table describes a single set of traffic
    // counters. The type is slice of
    // DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry.
    Diffservcountactentry []DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry
}

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetFilter() yfilter.YFilter { return diffservcountacttable.YFilter }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) SetFilter(yf yfilter.YFilter) { diffservcountacttable.YFilter = yf }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetGoName(yname string) string {
    if yname == "diffServCountActEntry" { return "Diffservcountactentry" }
    return ""
}

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetSegmentPath() string {
    return "diffServCountActTable"
}

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServCountActEntry" {
        for _, c := range diffservcountacttable.Diffservcountactentry {
            if diffservcountacttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry{}
        diffservcountacttable.Diffservcountactentry = append(diffservcountacttable.Diffservcountactentry, child)
        return &diffservcountacttable.Diffservcountactentry[len(diffservcountacttable.Diffservcountactentry)-1]
    }
    return nil
}

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservcountacttable.Diffservcountactentry {
        children[diffservcountacttable.Diffservcountactentry[i].GetSegmentPath()] = &diffservcountacttable.Diffservcountactentry[i]
    }
    return children
}

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetYangName() string { return "diffServCountActTable" }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) SetParent(parent types.Entity) { diffservcountacttable.parent = parent }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetParent() types.Entity { return diffservcountacttable.parent }

func (diffservcountacttable *DIFFSERVMIB_Diffservcountacttable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry
// An entry in the count action table describes a single set of
// traffic counters.
type DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry struct {
    parent types.Entity
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

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetFilter() yfilter.YFilter { return diffservcountactentry.YFilter }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) SetFilter(yf yfilter.YFilter) { diffservcountactentry.YFilter = yf }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetGoName(yname string) string {
    if yname == "diffServCountActId" { return "Diffservcountactid" }
    if yname == "diffServCountActOctets" { return "Diffservcountactoctets" }
    if yname == "diffServCountActPkts" { return "Diffservcountactpkts" }
    if yname == "diffServCountActStorage" { return "Diffservcountactstorage" }
    if yname == "diffServCountActStatus" { return "Diffservcountactstatus" }
    return ""
}

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetSegmentPath() string {
    return "diffServCountActEntry" + "[diffServCountActId='" + fmt.Sprintf("%v", diffservcountactentry.Diffservcountactid) + "']"
}

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServCountActId"] = diffservcountactentry.Diffservcountactid
    leafs["diffServCountActOctets"] = diffservcountactentry.Diffservcountactoctets
    leafs["diffServCountActPkts"] = diffservcountactentry.Diffservcountactpkts
    leafs["diffServCountActStorage"] = diffservcountactentry.Diffservcountactstorage
    leafs["diffServCountActStatus"] = diffservcountactentry.Diffservcountactstatus
    return leafs
}

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetYangName() string { return "diffServCountActEntry" }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) SetParent(parent types.Entity) { diffservcountactentry.parent = parent }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetParent() types.Entity { return diffservcountactentry.parent }

func (diffservcountactentry *DIFFSERVMIB_Diffservcountacttable_Diffservcountactentry) GetParentYangName() string { return "diffServCountActTable" }

// DIFFSERVMIB_Diffservalgdroptable
// The algorithmic drop table contains entries describing an
// element that drops packets according to some algorithm.
type DIFFSERVMIB_Diffservalgdroptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describes a process that drops packets according to some
    // algorithm. Further details of the algorithm type are to be found in
    // diffServAlgDropType and with more detail parameter entry pointed to by
    // diffServAlgDropSpecific when necessary. The type is slice of
    // DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry.
    Diffservalgdropentry []DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry
}

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetFilter() yfilter.YFilter { return diffservalgdroptable.YFilter }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) SetFilter(yf yfilter.YFilter) { diffservalgdroptable.YFilter = yf }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetGoName(yname string) string {
    if yname == "diffServAlgDropEntry" { return "Diffservalgdropentry" }
    return ""
}

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetSegmentPath() string {
    return "diffServAlgDropTable"
}

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServAlgDropEntry" {
        for _, c := range diffservalgdroptable.Diffservalgdropentry {
            if diffservalgdroptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry{}
        diffservalgdroptable.Diffservalgdropentry = append(diffservalgdroptable.Diffservalgdropentry, child)
        return &diffservalgdroptable.Diffservalgdropentry[len(diffservalgdroptable.Diffservalgdropentry)-1]
    }
    return nil
}

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservalgdroptable.Diffservalgdropentry {
        children[diffservalgdroptable.Diffservalgdropentry[i].GetSegmentPath()] = &diffservalgdroptable.Diffservalgdropentry[i]
    }
    return children
}

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetYangName() string { return "diffServAlgDropTable" }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) SetParent(parent types.Entity) { diffservalgdroptable.parent = parent }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetParent() types.Entity { return diffservalgdroptable.parent }

func (diffservalgdroptable *DIFFSERVMIB_Diffservalgdroptable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry
// An entry describes a process that drops packets according to
// some algorithm. Further details of the algorithm type are to be
// found in diffServAlgDropType and with more detail parameter entry
// pointed to by diffServAlgDropSpecific when necessary.
type DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Diffservalgdropnext interface{}

    // Points to an entry in the diffServQTable to indicate the queue that a drop
    // algorithm is to monitor when deciding whether to drop a packet. If the row
    // pointed to does not exist, the algorithmic dropper element is considered
    // inactive.    Setting this to point to a target that does not exist results
    // in an inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetFilter() yfilter.YFilter { return diffservalgdropentry.YFilter }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) SetFilter(yf yfilter.YFilter) { diffservalgdropentry.YFilter = yf }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetGoName(yname string) string {
    if yname == "diffServAlgDropId" { return "Diffservalgdropid" }
    if yname == "diffServAlgDropType" { return "Diffservalgdroptype" }
    if yname == "diffServAlgDropNext" { return "Diffservalgdropnext" }
    if yname == "diffServAlgDropQMeasure" { return "Diffservalgdropqmeasure" }
    if yname == "diffServAlgDropQThreshold" { return "Diffservalgdropqthreshold" }
    if yname == "diffServAlgDropSpecific" { return "Diffservalgdropspecific" }
    if yname == "diffServAlgDropOctets" { return "Diffservalgdropoctets" }
    if yname == "diffServAlgDropPkts" { return "Diffservalgdroppkts" }
    if yname == "diffServAlgRandomDropOctets" { return "Diffservalgrandomdropoctets" }
    if yname == "diffServAlgRandomDropPkts" { return "Diffservalgrandomdroppkts" }
    if yname == "diffServAlgDropStorage" { return "Diffservalgdropstorage" }
    if yname == "diffServAlgDropStatus" { return "Diffservalgdropstatus" }
    return ""
}

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetSegmentPath() string {
    return "diffServAlgDropEntry" + "[diffServAlgDropId='" + fmt.Sprintf("%v", diffservalgdropentry.Diffservalgdropid) + "']"
}

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServAlgDropId"] = diffservalgdropentry.Diffservalgdropid
    leafs["diffServAlgDropType"] = diffservalgdropentry.Diffservalgdroptype
    leafs["diffServAlgDropNext"] = diffservalgdropentry.Diffservalgdropnext
    leafs["diffServAlgDropQMeasure"] = diffservalgdropentry.Diffservalgdropqmeasure
    leafs["diffServAlgDropQThreshold"] = diffservalgdropentry.Diffservalgdropqthreshold
    leafs["diffServAlgDropSpecific"] = diffservalgdropentry.Diffservalgdropspecific
    leafs["diffServAlgDropOctets"] = diffservalgdropentry.Diffservalgdropoctets
    leafs["diffServAlgDropPkts"] = diffservalgdropentry.Diffservalgdroppkts
    leafs["diffServAlgRandomDropOctets"] = diffservalgdropentry.Diffservalgrandomdropoctets
    leafs["diffServAlgRandomDropPkts"] = diffservalgdropentry.Diffservalgrandomdroppkts
    leafs["diffServAlgDropStorage"] = diffservalgdropentry.Diffservalgdropstorage
    leafs["diffServAlgDropStatus"] = diffservalgdropentry.Diffservalgdropstatus
    return leafs
}

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetYangName() string { return "diffServAlgDropEntry" }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) SetParent(parent types.Entity) { diffservalgdropentry.parent = parent }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetParent() types.Entity { return diffservalgdropentry.parent }

func (diffservalgdropentry *DIFFSERVMIB_Diffservalgdroptable_Diffservalgdropentry) GetParentYangName() string { return "diffServAlgDropTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describes a process that drops packets according to a random
    // algorithm. The type is slice of
    // DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry.
    Diffservrandomdropentry []DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry
}

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetFilter() yfilter.YFilter { return diffservrandomdroptable.YFilter }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) SetFilter(yf yfilter.YFilter) { diffservrandomdroptable.YFilter = yf }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetGoName(yname string) string {
    if yname == "diffServRandomDropEntry" { return "Diffservrandomdropentry" }
    return ""
}

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetSegmentPath() string {
    return "diffServRandomDropTable"
}

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServRandomDropEntry" {
        for _, c := range diffservrandomdroptable.Diffservrandomdropentry {
            if diffservrandomdroptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry{}
        diffservrandomdroptable.Diffservrandomdropentry = append(diffservrandomdroptable.Diffservrandomdropentry, child)
        return &diffservrandomdroptable.Diffservrandomdropentry[len(diffservrandomdroptable.Diffservrandomdropentry)-1]
    }
    return nil
}

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservrandomdroptable.Diffservrandomdropentry {
        children[diffservrandomdroptable.Diffservrandomdropentry[i].GetSegmentPath()] = &diffservrandomdroptable.Diffservrandomdropentry[i]
    }
    return children
}

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetYangName() string { return "diffServRandomDropTable" }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) SetParent(parent types.Entity) { diffservrandomdroptable.parent = parent }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetParent() types.Entity { return diffservrandomdroptable.parent }

func (diffservrandomdroptable *DIFFSERVMIB_Diffservrandomdroptable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry
// An entry describes a process that drops packets according to a
// random algorithm.
type DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry struct {
    parent types.Entity
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

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetFilter() yfilter.YFilter { return diffservrandomdropentry.YFilter }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) SetFilter(yf yfilter.YFilter) { diffservrandomdropentry.YFilter = yf }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetGoName(yname string) string {
    if yname == "diffServRandomDropId" { return "Diffservrandomdropid" }
    if yname == "diffServRandomDropMinThreshBytes" { return "Diffservrandomdropminthreshbytes" }
    if yname == "diffServRandomDropMinThreshPkts" { return "Diffservrandomdropminthreshpkts" }
    if yname == "diffServRandomDropMaxThreshBytes" { return "Diffservrandomdropmaxthreshbytes" }
    if yname == "diffServRandomDropMaxThreshPkts" { return "Diffservrandomdropmaxthreshpkts" }
    if yname == "diffServRandomDropProbMax" { return "Diffservrandomdropprobmax" }
    if yname == "diffServRandomDropWeight" { return "Diffservrandomdropweight" }
    if yname == "diffServRandomDropSamplingRate" { return "Diffservrandomdropsamplingrate" }
    if yname == "diffServRandomDropStorage" { return "Diffservrandomdropstorage" }
    if yname == "diffServRandomDropStatus" { return "Diffservrandomdropstatus" }
    return ""
}

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetSegmentPath() string {
    return "diffServRandomDropEntry" + "[diffServRandomDropId='" + fmt.Sprintf("%v", diffservrandomdropentry.Diffservrandomdropid) + "']"
}

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServRandomDropId"] = diffservrandomdropentry.Diffservrandomdropid
    leafs["diffServRandomDropMinThreshBytes"] = diffservrandomdropentry.Diffservrandomdropminthreshbytes
    leafs["diffServRandomDropMinThreshPkts"] = diffservrandomdropentry.Diffservrandomdropminthreshpkts
    leafs["diffServRandomDropMaxThreshBytes"] = diffservrandomdropentry.Diffservrandomdropmaxthreshbytes
    leafs["diffServRandomDropMaxThreshPkts"] = diffservrandomdropentry.Diffservrandomdropmaxthreshpkts
    leafs["diffServRandomDropProbMax"] = diffservrandomdropentry.Diffservrandomdropprobmax
    leafs["diffServRandomDropWeight"] = diffservrandomdropentry.Diffservrandomdropweight
    leafs["diffServRandomDropSamplingRate"] = diffservrandomdropentry.Diffservrandomdropsamplingrate
    leafs["diffServRandomDropStorage"] = diffservrandomdropentry.Diffservrandomdropstorage
    leafs["diffServRandomDropStatus"] = diffservrandomdropentry.Diffservrandomdropstatus
    return leafs
}

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetYangName() string { return "diffServRandomDropEntry" }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) SetParent(parent types.Entity) { diffservrandomdropentry.parent = parent }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetParent() types.Entity { return diffservrandomdropentry.parent }

func (diffservrandomdropentry *DIFFSERVMIB_Diffservrandomdroptable_Diffservrandomdropentry) GetParentYangName() string { return "diffServRandomDropTable" }

// DIFFSERVMIB_Diffservqtable
// The Queue Table enumerates the individual queues.  Note that the
// MIB models queuing systems as composed of individual queues, one
// per class of traffic, even though they may in fact be structured
// as classes of traffic scheduled using a common calendar queue, or
// in other ways.
type DIFFSERVMIB_Diffservqtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the Queue Table describes a single queue or class of traffic.
    // The type is slice of DIFFSERVMIB_Diffservqtable_Diffservqentry.
    Diffservqentry []DIFFSERVMIB_Diffservqtable_Diffservqentry
}

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetFilter() yfilter.YFilter { return diffservqtable.YFilter }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) SetFilter(yf yfilter.YFilter) { diffservqtable.YFilter = yf }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetGoName(yname string) string {
    if yname == "diffServQEntry" { return "Diffservqentry" }
    return ""
}

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetSegmentPath() string {
    return "diffServQTable"
}

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServQEntry" {
        for _, c := range diffservqtable.Diffservqentry {
            if diffservqtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservqtable_Diffservqentry{}
        diffservqtable.Diffservqentry = append(diffservqtable.Diffservqentry, child)
        return &diffservqtable.Diffservqentry[len(diffservqtable.Diffservqentry)-1]
    }
    return nil
}

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservqtable.Diffservqentry {
        children[diffservqtable.Diffservqentry[i].GetSegmentPath()] = &diffservqtable.Diffservqentry[i]
    }
    return children
}

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetYangName() string { return "diffServQTable" }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) SetParent(parent types.Entity) { diffservqtable.parent = parent }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetParent() types.Entity { return diffservqtable.parent }

func (diffservqtable *DIFFSERVMIB_Diffservqtable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservqtable_Diffservqentry
// An entry in the Queue Table describes a single queue or class of
// traffic.
type DIFFSERVMIB_Diffservqtable_Diffservqentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Diffservqnext interface{}

    // This RowPointer indicates the diffServMinRateEntry that the scheduler,
    // pointed to by diffServQNext, should use to service this queue.  If the row
    // pointed to is zeroDotZero, the minimum rate and priority is unspecified. 
    // Setting this to point to a target that does not exist results in an
    // inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Diffservqminrate interface{}

    // This RowPointer indicates the diffServMaxRateEntry that the scheduler,
    // pointed to by diffServQNext, should use to service this queue.  If the row
    // pointed to is zeroDotZero, the maximum rate is the line speed of the
    // interface.     Setting this to point to a target that does not exist
    // results in an inconsistentValue error.  If the row pointed to is removed or
    // becomes inactive by other means, the treatment is as if this attribute
    // contains a value of zeroDotZero. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetFilter() yfilter.YFilter { return diffservqentry.YFilter }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) SetFilter(yf yfilter.YFilter) { diffservqentry.YFilter = yf }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetGoName(yname string) string {
    if yname == "diffServQId" { return "Diffservqid" }
    if yname == "diffServQNext" { return "Diffservqnext" }
    if yname == "diffServQMinRate" { return "Diffservqminrate" }
    if yname == "diffServQMaxRate" { return "Diffservqmaxrate" }
    if yname == "diffServQStorage" { return "Diffservqstorage" }
    if yname == "diffServQStatus" { return "Diffservqstatus" }
    return ""
}

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetSegmentPath() string {
    return "diffServQEntry" + "[diffServQId='" + fmt.Sprintf("%v", diffservqentry.Diffservqid) + "']"
}

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServQId"] = diffservqentry.Diffservqid
    leafs["diffServQNext"] = diffservqentry.Diffservqnext
    leafs["diffServQMinRate"] = diffservqentry.Diffservqminrate
    leafs["diffServQMaxRate"] = diffservqentry.Diffservqmaxrate
    leafs["diffServQStorage"] = diffservqentry.Diffservqstorage
    leafs["diffServQStatus"] = diffservqentry.Diffservqstatus
    return leafs
}

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetYangName() string { return "diffServQEntry" }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) SetParent(parent types.Entity) { diffservqentry.parent = parent }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetParent() types.Entity { return diffservqentry.parent }

func (diffservqentry *DIFFSERVMIB_Diffservqtable_Diffservqentry) GetParentYangName() string { return "diffServQTable" }

// DIFFSERVMIB_Diffservschedulertable
// The Scheduler Table enumerates packet schedulers. Multiple
// scheduling algorithms can be used on a given data path, with each
// algorithm described by one diffServSchedulerEntry.
type DIFFSERVMIB_Diffservschedulertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the Scheduler Table describing a single instance of a
    // scheduling algorithm. The type is slice of
    // DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry.
    Diffservschedulerentry []DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry
}

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetFilter() yfilter.YFilter { return diffservschedulertable.YFilter }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) SetFilter(yf yfilter.YFilter) { diffservschedulertable.YFilter = yf }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetGoName(yname string) string {
    if yname == "diffServSchedulerEntry" { return "Diffservschedulerentry" }
    return ""
}

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetSegmentPath() string {
    return "diffServSchedulerTable"
}

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServSchedulerEntry" {
        for _, c := range diffservschedulertable.Diffservschedulerentry {
            if diffservschedulertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry{}
        diffservschedulertable.Diffservschedulerentry = append(diffservschedulertable.Diffservschedulerentry, child)
        return &diffservschedulertable.Diffservschedulerentry[len(diffservschedulertable.Diffservschedulerentry)-1]
    }
    return nil
}

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservschedulertable.Diffservschedulerentry {
        children[diffservschedulertable.Diffservschedulerentry[i].GetSegmentPath()] = &diffservschedulertable.Diffservschedulerentry[i]
    }
    return children
}

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetYangName() string { return "diffServSchedulerTable" }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) SetParent(parent types.Entity) { diffservschedulertable.parent = parent }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetParent() types.Entity { return diffservschedulertable.parent }

func (diffservschedulertable *DIFFSERVMIB_Diffservschedulertable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry
// An entry in the Scheduler Table describing a single instance of
// a scheduling algorithm.
type DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Diffservschedulernext interface{}

    // The scheduling algorithm used by this Scheduler. zeroDotZero indicates that
    // this is unknown.  Standard values for generic algorithms:
    // diffServSchedulerPriority, diffServSchedulerWRR, and diffServSchedulerWFQ
    // are specified in this MIB; additional values    may be further specified in
    // other MIBs. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Diffservschedulermethod interface{}

    // This RowPointer indicates the entry in diffServMinRateTable which indicates
    // the priority or minimum output rate from this scheduler. This attribute is
    // used only when there is more than one level of scheduler.  When it has the
    // value zeroDotZero, it indicates that no minimum rate or priority is
    // imposed.  Setting this to point to a target that does not exist results in
    // an inconsistentValue error.  If the row pointed to is removed or becomes
    // inactive by other means, the treatment is as if this attribute contains a
    // value of zeroDotZero. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetFilter() yfilter.YFilter { return diffservschedulerentry.YFilter }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) SetFilter(yf yfilter.YFilter) { diffservschedulerentry.YFilter = yf }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetGoName(yname string) string {
    if yname == "diffServSchedulerId" { return "Diffservschedulerid" }
    if yname == "diffServSchedulerNext" { return "Diffservschedulernext" }
    if yname == "diffServSchedulerMethod" { return "Diffservschedulermethod" }
    if yname == "diffServSchedulerMinRate" { return "Diffservschedulerminrate" }
    if yname == "diffServSchedulerMaxRate" { return "Diffservschedulermaxrate" }
    if yname == "diffServSchedulerStorage" { return "Diffservschedulerstorage" }
    if yname == "diffServSchedulerStatus" { return "Diffservschedulerstatus" }
    return ""
}

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetSegmentPath() string {
    return "diffServSchedulerEntry" + "[diffServSchedulerId='" + fmt.Sprintf("%v", diffservschedulerentry.Diffservschedulerid) + "']"
}

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServSchedulerId"] = diffservschedulerentry.Diffservschedulerid
    leafs["diffServSchedulerNext"] = diffservschedulerentry.Diffservschedulernext
    leafs["diffServSchedulerMethod"] = diffservschedulerentry.Diffservschedulermethod
    leafs["diffServSchedulerMinRate"] = diffservschedulerentry.Diffservschedulerminrate
    leafs["diffServSchedulerMaxRate"] = diffservschedulerentry.Diffservschedulermaxrate
    leafs["diffServSchedulerStorage"] = diffservschedulerentry.Diffservschedulerstorage
    leafs["diffServSchedulerStatus"] = diffservschedulerentry.Diffservschedulerstatus
    return leafs
}

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetYangName() string { return "diffServSchedulerEntry" }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) SetParent(parent types.Entity) { diffservschedulerentry.parent = parent }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetParent() types.Entity { return diffservschedulerentry.parent }

func (diffservschedulerentry *DIFFSERVMIB_Diffservschedulertable_Diffservschedulerentry) GetParentYangName() string { return "diffServSchedulerTable" }

// DIFFSERVMIB_Diffservminratetable
// The Minimum Rate Parameters Table enumerates individual sets of
// scheduling parameter that can be used/reused by Queues and
// Schedulers.
type DIFFSERVMIB_Diffservminratetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the Minimum Rate Parameters Table describes a single set of
    // scheduling parameters for use by one or more queues or schedulers. The type
    // is slice of DIFFSERVMIB_Diffservminratetable_Diffservminrateentry.
    Diffservminrateentry []DIFFSERVMIB_Diffservminratetable_Diffservminrateentry
}

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetFilter() yfilter.YFilter { return diffservminratetable.YFilter }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) SetFilter(yf yfilter.YFilter) { diffservminratetable.YFilter = yf }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetGoName(yname string) string {
    if yname == "diffServMinRateEntry" { return "Diffservminrateentry" }
    return ""
}

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetSegmentPath() string {
    return "diffServMinRateTable"
}

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServMinRateEntry" {
        for _, c := range diffservminratetable.Diffservminrateentry {
            if diffservminratetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservminratetable_Diffservminrateentry{}
        diffservminratetable.Diffservminrateentry = append(diffservminratetable.Diffservminrateentry, child)
        return &diffservminratetable.Diffservminrateentry[len(diffservminratetable.Diffservminrateentry)-1]
    }
    return nil
}

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservminratetable.Diffservminrateentry {
        children[diffservminratetable.Diffservminrateentry[i].GetSegmentPath()] = &diffservminratetable.Diffservminrateentry[i]
    }
    return children
}

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetYangName() string { return "diffServMinRateTable" }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) SetParent(parent types.Entity) { diffservminratetable.parent = parent }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetParent() types.Entity { return diffservminratetable.parent }

func (diffservminratetable *DIFFSERVMIB_Diffservminratetable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservminratetable_Diffservminrateentry
// An entry in the Minimum Rate Parameters Table describes a single
// set of scheduling parameters for use by one or more queues or
// schedulers.
type DIFFSERVMIB_Diffservminratetable_Diffservminrateentry struct {
    parent types.Entity
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

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetFilter() yfilter.YFilter { return diffservminrateentry.YFilter }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) SetFilter(yf yfilter.YFilter) { diffservminrateentry.YFilter = yf }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetGoName(yname string) string {
    if yname == "diffServMinRateId" { return "Diffservminrateid" }
    if yname == "diffServMinRatePriority" { return "Diffservminratepriority" }
    if yname == "diffServMinRateAbsolute" { return "Diffservminrateabsolute" }
    if yname == "diffServMinRateRelative" { return "Diffservminraterelative" }
    if yname == "diffServMinRateStorage" { return "Diffservminratestorage" }
    if yname == "diffServMinRateStatus" { return "Diffservminratestatus" }
    return ""
}

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetSegmentPath() string {
    return "diffServMinRateEntry" + "[diffServMinRateId='" + fmt.Sprintf("%v", diffservminrateentry.Diffservminrateid) + "']"
}

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServMinRateId"] = diffservminrateentry.Diffservminrateid
    leafs["diffServMinRatePriority"] = diffservminrateentry.Diffservminratepriority
    leafs["diffServMinRateAbsolute"] = diffservminrateentry.Diffservminrateabsolute
    leafs["diffServMinRateRelative"] = diffservminrateentry.Diffservminraterelative
    leafs["diffServMinRateStorage"] = diffservminrateentry.Diffservminratestorage
    leafs["diffServMinRateStatus"] = diffservminrateentry.Diffservminratestatus
    return leafs
}

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetYangName() string { return "diffServMinRateEntry" }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) SetParent(parent types.Entity) { diffservminrateentry.parent = parent }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetParent() types.Entity { return diffservminrateentry.parent }

func (diffservminrateentry *DIFFSERVMIB_Diffservminratetable_Diffservminrateentry) GetParentYangName() string { return "diffServMinRateTable" }

// DIFFSERVMIB_Diffservmaxratetable
// The Maximum Rate Parameter Table enumerates individual sets of
// scheduling parameter that can be used/reused by Queues and
// Schedulers.
type DIFFSERVMIB_Diffservmaxratetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the Maximum Rate Parameter Table describes a single set of
    // scheduling parameters for use by one or more queues or schedulers. The type
    // is slice of DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry.
    Diffservmaxrateentry []DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry
}

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetFilter() yfilter.YFilter { return diffservmaxratetable.YFilter }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) SetFilter(yf yfilter.YFilter) { diffservmaxratetable.YFilter = yf }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetGoName(yname string) string {
    if yname == "diffServMaxRateEntry" { return "Diffservmaxrateentry" }
    return ""
}

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetSegmentPath() string {
    return "diffServMaxRateTable"
}

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffServMaxRateEntry" {
        for _, c := range diffservmaxratetable.Diffservmaxrateentry {
            if diffservmaxratetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry{}
        diffservmaxratetable.Diffservmaxrateentry = append(diffservmaxratetable.Diffservmaxrateentry, child)
        return &diffservmaxratetable.Diffservmaxrateentry[len(diffservmaxratetable.Diffservmaxrateentry)-1]
    }
    return nil
}

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservmaxratetable.Diffservmaxrateentry {
        children[diffservmaxratetable.Diffservmaxrateentry[i].GetSegmentPath()] = &diffservmaxratetable.Diffservmaxrateentry[i]
    }
    return children
}

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetBundleName() string { return "cisco_ios_xe" }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetYangName() string { return "diffServMaxRateTable" }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) SetParent(parent types.Entity) { diffservmaxratetable.parent = parent }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetParent() types.Entity { return diffservmaxratetable.parent }

func (diffservmaxratetable *DIFFSERVMIB_Diffservmaxratetable) GetParentYangName() string { return "DIFFSERV-MIB" }

// DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry
// An entry in the Maximum Rate Parameter Table describes a single
// set of scheduling parameters for use by one or more queues or
// schedulers.
type DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry struct {
    parent types.Entity
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

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetFilter() yfilter.YFilter { return diffservmaxrateentry.YFilter }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) SetFilter(yf yfilter.YFilter) { diffservmaxrateentry.YFilter = yf }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetGoName(yname string) string {
    if yname == "diffServMaxRateId" { return "Diffservmaxrateid" }
    if yname == "diffServMaxRateLevel" { return "Diffservmaxratelevel" }
    if yname == "diffServMaxRateAbsolute" { return "Diffservmaxrateabsolute" }
    if yname == "diffServMaxRateRelative" { return "Diffservmaxraterelative" }
    if yname == "diffServMaxRateThreshold" { return "Diffservmaxratethreshold" }
    if yname == "diffServMaxRateStorage" { return "Diffservmaxratestorage" }
    if yname == "diffServMaxRateStatus" { return "Diffservmaxratestatus" }
    return ""
}

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetSegmentPath() string {
    return "diffServMaxRateEntry" + "[diffServMaxRateId='" + fmt.Sprintf("%v", diffservmaxrateentry.Diffservmaxrateid) + "']" + "[diffServMaxRateLevel='" + fmt.Sprintf("%v", diffservmaxrateentry.Diffservmaxratelevel) + "']"
}

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["diffServMaxRateId"] = diffservmaxrateentry.Diffservmaxrateid
    leafs["diffServMaxRateLevel"] = diffservmaxrateentry.Diffservmaxratelevel
    leafs["diffServMaxRateAbsolute"] = diffservmaxrateentry.Diffservmaxrateabsolute
    leafs["diffServMaxRateRelative"] = diffservmaxrateentry.Diffservmaxraterelative
    leafs["diffServMaxRateThreshold"] = diffservmaxrateentry.Diffservmaxratethreshold
    leafs["diffServMaxRateStorage"] = diffservmaxrateentry.Diffservmaxratestorage
    leafs["diffServMaxRateStatus"] = diffservmaxrateentry.Diffservmaxratestatus
    return leafs
}

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetYangName() string { return "diffServMaxRateEntry" }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) SetParent(parent types.Entity) { diffservmaxrateentry.parent = parent }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetParent() types.Entity { return diffservmaxrateentry.parent }

func (diffservmaxrateentry *DIFFSERVMIB_Diffservmaxratetable_Diffservmaxrateentry) GetParentYangName() string { return "diffServMaxRateTable" }

