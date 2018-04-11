// The Cisco QOS Policy PIB for provisioning QOS policy.
package cisco_qos_pib_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_qos_pib_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-QOS-PIB-MIB CISCO-QOS-PIB-MIB}", reflect.TypeOf(CISCOQOSPIBMIB{}))
    ydk.RegisterEntity("CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB", reflect.TypeOf(CISCOQOSPIBMIB{}))
}

// QueueRange represents number of codepoints.
type QueueRange string

const (
    QueueRange_oneQ QueueRange = "oneQ"

    QueueRange_twoQ QueueRange = "twoQ"

    QueueRange_threeQ QueueRange = "threeQ"

    QueueRange_fourQ QueueRange = "fourQ"

    QueueRange_eightQ QueueRange = "eightQ"

    QueueRange_sixteenQ QueueRange = "sixteenQ"

    QueueRange_thirtyTwoQ QueueRange = "thirtyTwoQ"

    QueueRange_sixtyFourQ QueueRange = "sixtyFourQ"
)

// ThresholdSetRange represents parameters packets are randomly dropped.
type ThresholdSetRange string

const (
    ThresholdSetRange_zeroT ThresholdSetRange = "zeroT"

    ThresholdSetRange_oneT ThresholdSetRange = "oneT"

    ThresholdSetRange_twoT ThresholdSetRange = "twoT"

    ThresholdSetRange_fourT ThresholdSetRange = "fourT"

    ThresholdSetRange_eightT ThresholdSetRange = "eightT"
)

// QosInterfaceQueueType represents number of DSCPs.
type QosInterfaceQueueType string

const (
    QosInterfaceQueueType_oneQ1t QosInterfaceQueueType = "oneQ1t"

    QosInterfaceQueueType_oneQ2t QosInterfaceQueueType = "oneQ2t"

    QosInterfaceQueueType_oneQ4t QosInterfaceQueueType = "oneQ4t"

    QosInterfaceQueueType_oneQ8t QosInterfaceQueueType = "oneQ8t"

    QosInterfaceQueueType_twoQ1t QosInterfaceQueueType = "twoQ1t"

    QosInterfaceQueueType_twoQ2t QosInterfaceQueueType = "twoQ2t"

    QosInterfaceQueueType_twoQ4t QosInterfaceQueueType = "twoQ4t"

    QosInterfaceQueueType_twoQ8t QosInterfaceQueueType = "twoQ8t"

    QosInterfaceQueueType_threeQ1t QosInterfaceQueueType = "threeQ1t"

    QosInterfaceQueueType_threeQ2t QosInterfaceQueueType = "threeQ2t"

    QosInterfaceQueueType_threeQ4t QosInterfaceQueueType = "threeQ4t"

    QosInterfaceQueueType_threeQ8t QosInterfaceQueueType = "threeQ8t"

    QosInterfaceQueueType_fourQ1t QosInterfaceQueueType = "fourQ1t"

    QosInterfaceQueueType_fourQ2t QosInterfaceQueueType = "fourQ2t"

    QosInterfaceQueueType_fourQ4t QosInterfaceQueueType = "fourQ4t"

    QosInterfaceQueueType_fourQ8t QosInterfaceQueueType = "fourQ8t"

    QosInterfaceQueueType_eightQ1t QosInterfaceQueueType = "eightQ1t"

    QosInterfaceQueueType_eightQ2t QosInterfaceQueueType = "eightQ2t"

    QosInterfaceQueueType_eightQ4t QosInterfaceQueueType = "eightQ4t"

    QosInterfaceQueueType_eightQ8t QosInterfaceQueueType = "eightQ8t"

    QosInterfaceQueueType_sixteenQ1t QosInterfaceQueueType = "sixteenQ1t"

    QosInterfaceQueueType_sixteenQ2t QosInterfaceQueueType = "sixteenQ2t"

    QosInterfaceQueueType_sixteenQ4t QosInterfaceQueueType = "sixteenQ4t"

    QosInterfaceQueueType_sixtyfourQ1t QosInterfaceQueueType = "sixtyfourQ1t"

    QosInterfaceQueueType_sixtyfourQ2t QosInterfaceQueueType = "sixtyfourQ2t"

    QosInterfaceQueueType_sixtyfourQ4t QosInterfaceQueueType = "sixtyfourQ4t"

    QosInterfaceQueueType_oneP1Q0t QosInterfaceQueueType = "oneP1Q0t"

    QosInterfaceQueueType_oneP1Q4t QosInterfaceQueueType = "oneP1Q4t"

    QosInterfaceQueueType_oneP1Q8t QosInterfaceQueueType = "oneP1Q8t"

    QosInterfaceQueueType_oneP2Q1t QosInterfaceQueueType = "oneP2Q1t"

    QosInterfaceQueueType_oneP2Q2t QosInterfaceQueueType = "oneP2Q2t"

    QosInterfaceQueueType_oneP3Q1t QosInterfaceQueueType = "oneP3Q1t"

    QosInterfaceQueueType_oneP7Q8t QosInterfaceQueueType = "oneP7Q8t"

    QosInterfaceQueueType_oneP3Q8t QosInterfaceQueueType = "oneP3Q8t"

    QosInterfaceQueueType_sixteenQ8t QosInterfaceQueueType = "sixteenQ8t"

    QosInterfaceQueueType_oneP15Q8t QosInterfaceQueueType = "oneP15Q8t"

    QosInterfaceQueueType_oneP15Q1t QosInterfaceQueueType = "oneP15Q1t"

    QosInterfaceQueueType_oneP7Q1t QosInterfaceQueueType = "oneP7Q1t"

    QosInterfaceQueueType_oneP31Q1t QosInterfaceQueueType = "oneP31Q1t"

    QosInterfaceQueueType_thirtytwoQ1t QosInterfaceQueueType = "thirtytwoQ1t"

    QosInterfaceQueueType_thirtytwoQ8t QosInterfaceQueueType = "thirtytwoQ8t"

    QosInterfaceQueueType_oneP31Q8t QosInterfaceQueueType = "oneP31Q8t"

    QosInterfaceQueueType_oneP7Q4t QosInterfaceQueueType = "oneP7Q4t"

    QosInterfaceQueueType_oneP3Q4t QosInterfaceQueueType = "oneP3Q4t"

    QosInterfaceQueueType_oneP7Q2t QosInterfaceQueueType = "oneP7Q2t"
)

// CISCOQOSPIBMIB
type CISCOQOSPIBMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This class contains a single policy instance that identifies the current
    // incarnation of the PIB and the PDP that installed this incarnation.  The
    // instance of this class is reported to the PDP at client connect time so
    // that the PDP can (attempt to) ascertain the current state of the PIB.
    Qosdevicepibincarnationtable CISCOQOSPIBMIB_Qosdevicepibincarnationtable

    // The single instance of this class indicates specific attributes of the
    // device.  These include configuration values such as the configured PDP
    // addresses, the maximum message size, and specific device capabilities.  The
    // latter include input port-based and output port-based classification and/or
    // policing, support for flow based policing, aggregate based policing,
    // traffic shaping capabilities, etc.
    Qosdeviceattributetable CISCOQOSPIBMIB_Qosdeviceattributetable

    // This class describes the interface types of the interfaces that exist on
    // the device.  It includes the queue type, role combination and capabilities
    // of interfaces.  The PEP does not report which specific interfaces have
    // which characteristics.
    Qosinterfacetypetable CISCOQOSPIBMIB_Qosinterfacetypetable

    // Maps each DSCP to a marked-down DSCP.  Also maps each DSCP to an IP
    // precedence and QosLayer2Cos.  When configured for the first time, all 64
    // entries of the table must be specified. Thereafter, instances may be
    // modified (with a delete and install in a single decision) but not deleted
    // unless all instances are deleted.
    Qosdiffservmappingtable CISCOQOSPIBMIB_Qosdiffservmappingtable

    // Maps each of eight CoS values to a DSCP.  When configured for the first
    // time, all 8 entries of the table must be specified. Thereafter, instances
    // may be modified (with a delete and install in a single decision) but not
    // deleted unless all instances are deleted.
    Qoscostodscptable CISCOQOSPIBMIB_Qoscostodscptable

    // A policy class that specifies what QoS to apply to a packet that does not
    // match any other policy configured for this role combination for a
    // particular direction of traffic.
    Qosunmatchedpolicytable CISCOQOSPIBMIB_Qosunmatchedpolicytable

    // A class specifying policing parameters for both microflows and aggregate
    // flows.  This table is designed for policing according to a token bucket
    // scheme where an average rate and burst size is specified.
    Qospolicertable CISCOQOSPIBMIB_Qospolicertable

    // Instances of this class identify aggregate flows and the policer to apply
    // to each.
    Qosaggregatetable CISCOQOSPIBMIB_Qosaggregatetable

    // A class of MAC/Vlan tuples and their associated CoS values.
    Qosmacclassificationtable CISCOQOSPIBMIB_Qosmacclassificationtable

    // ACE definitions.
    Qosipacetable CISCOQOSPIBMIB_Qosipacetable

    // A class that defines a set of ACLs each being an ordered list of ACEs.
    Qosipacldefinitiontable CISCOQOSPIBMIB_Qosipacldefinitiontable

    // A class that applies a set of ACLs to interfaces specifying, for each
    // interface the order of the ACL with respect to other ACLs applied to the
    // same interface and, for each ACL the action to take for a packet that
    // matches a permit ACE in that ACL.  Interfaces are specified abstractly in
    // terms of interface role combinations.
    Qosipaclactiontable CISCOQOSPIBMIB_Qosipaclactiontable

    // This class specifies the scheduling preference an interface chooses if it
    // supports multiple scheduling types.  Higher values are preferred over lower
    // values.
    Qosifschedulingpreferencestable CISCOQOSPIBMIB_Qosifschedulingpreferencestable

    // This class specifies the preference of the drop mechanism an interface
    // chooses if it supports multiple drop mechanisms. Higher values are
    // preferred over lower values.
    Qosifdroppreferencetable CISCOQOSPIBMIB_Qosifdroppreferencetable

    // The assignment of each DSCP to a queue and threshold for each interface
    // queue type.
    Qosifdscpassignmenttable CISCOQOSPIBMIB_Qosifdscpassignmenttable

    // A class of lower and upper values for each threshold set in a queue
    // supporting WRED.  If the size of the queue for a given threshold is below
    // the lower value then packets assigned to that threshold are always accepted
    // into the queue.  If the size of the queue is above upper value then packets
    // are always dropped.  If the size of the queue is between the lower and the
    // upper then packets are randomly dropped.
    Qosifredtable CISCOQOSPIBMIB_Qosifredtable

    // A class for threshold sets in a queue supporting tail drop. If the size of
    // the queue for a given threshold set is at or below the specified value then
    // packets assigned to that threshold set are always accepted into the queue. 
    // If the size of the queue is above the specified value then packets are
    // always dropped.
    Qosiftaildroptable CISCOQOSPIBMIB_Qosiftaildroptable

    // A class of scheduling weights for each queue of an interface that supports
    // weighted round robin scheduling or a mix of priority queueing and weighted
    // round robin.  For a queue with N priority queues, the N highest queue
    // numbers are the priority queues with the highest queue number having the
    // highest priority.  WRR is applied to the non-priority queues.
    Qosifweightstable CISCOQOSPIBMIB_Qosifweightstable
}

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetEntityData() *types.CommonEntityData {
    cISCOQOSPIBMIB.EntityData.YFilter = cISCOQOSPIBMIB.YFilter
    cISCOQOSPIBMIB.EntityData.YangName = "CISCO-QOS-PIB-MIB"
    cISCOQOSPIBMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOQOSPIBMIB.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    cISCOQOSPIBMIB.EntityData.SegmentPath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB"
    cISCOQOSPIBMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOQOSPIBMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOQOSPIBMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOQOSPIBMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOQOSPIBMIB.EntityData.Children["qosDevicePibIncarnationTable"] = types.YChild{"Qosdevicepibincarnationtable", &cISCOQOSPIBMIB.Qosdevicepibincarnationtable}
    cISCOQOSPIBMIB.EntityData.Children["qosDeviceAttributeTable"] = types.YChild{"Qosdeviceattributetable", &cISCOQOSPIBMIB.Qosdeviceattributetable}
    cISCOQOSPIBMIB.EntityData.Children["qosInterfaceTypeTable"] = types.YChild{"Qosinterfacetypetable", &cISCOQOSPIBMIB.Qosinterfacetypetable}
    cISCOQOSPIBMIB.EntityData.Children["qosDiffServMappingTable"] = types.YChild{"Qosdiffservmappingtable", &cISCOQOSPIBMIB.Qosdiffservmappingtable}
    cISCOQOSPIBMIB.EntityData.Children["qosCosToDscpTable"] = types.YChild{"Qoscostodscptable", &cISCOQOSPIBMIB.Qoscostodscptable}
    cISCOQOSPIBMIB.EntityData.Children["qosUnmatchedPolicyTable"] = types.YChild{"Qosunmatchedpolicytable", &cISCOQOSPIBMIB.Qosunmatchedpolicytable}
    cISCOQOSPIBMIB.EntityData.Children["qosPolicerTable"] = types.YChild{"Qospolicertable", &cISCOQOSPIBMIB.Qospolicertable}
    cISCOQOSPIBMIB.EntityData.Children["qosAggregateTable"] = types.YChild{"Qosaggregatetable", &cISCOQOSPIBMIB.Qosaggregatetable}
    cISCOQOSPIBMIB.EntityData.Children["qosMacClassificationTable"] = types.YChild{"Qosmacclassificationtable", &cISCOQOSPIBMIB.Qosmacclassificationtable}
    cISCOQOSPIBMIB.EntityData.Children["qosIpAceTable"] = types.YChild{"Qosipacetable", &cISCOQOSPIBMIB.Qosipacetable}
    cISCOQOSPIBMIB.EntityData.Children["qosIpAclDefinitionTable"] = types.YChild{"Qosipacldefinitiontable", &cISCOQOSPIBMIB.Qosipacldefinitiontable}
    cISCOQOSPIBMIB.EntityData.Children["qosIpAclActionTable"] = types.YChild{"Qosipaclactiontable", &cISCOQOSPIBMIB.Qosipaclactiontable}
    cISCOQOSPIBMIB.EntityData.Children["qosIfSchedulingPreferencesTable"] = types.YChild{"Qosifschedulingpreferencestable", &cISCOQOSPIBMIB.Qosifschedulingpreferencestable}
    cISCOQOSPIBMIB.EntityData.Children["qosIfDropPreferenceTable"] = types.YChild{"Qosifdroppreferencetable", &cISCOQOSPIBMIB.Qosifdroppreferencetable}
    cISCOQOSPIBMIB.EntityData.Children["qosIfDscpAssignmentTable"] = types.YChild{"Qosifdscpassignmenttable", &cISCOQOSPIBMIB.Qosifdscpassignmenttable}
    cISCOQOSPIBMIB.EntityData.Children["qosIfRedTable"] = types.YChild{"Qosifredtable", &cISCOQOSPIBMIB.Qosifredtable}
    cISCOQOSPIBMIB.EntityData.Children["qosIfTailDropTable"] = types.YChild{"Qosiftaildroptable", &cISCOQOSPIBMIB.Qosiftaildroptable}
    cISCOQOSPIBMIB.EntityData.Children["qosIfWeightsTable"] = types.YChild{"Qosifweightstable", &cISCOQOSPIBMIB.Qosifweightstable}
    cISCOQOSPIBMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOQOSPIBMIB.EntityData)
}

// CISCOQOSPIBMIB_Qosdevicepibincarnationtable
// This class contains a single policy instance that identifies
// the current incarnation of the PIB and the PDP that installed
// this incarnation.  The instance of this class is reported to
// the PDP at client connect time so that the PDP can (attempt
// to) ascertain the current state of the PIB.
type CISCOQOSPIBMIB_Qosdevicepibincarnationtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The single policy instance of this class identifies the current incarnation
    // of the PIB and the PDP that installed this incarnation. The type is slice
    // of
    // CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry.
    Qosdevicepibincarnationentry []CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry
}

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetEntityData() *types.CommonEntityData {
    qosdevicepibincarnationtable.EntityData.YFilter = qosdevicepibincarnationtable.YFilter
    qosdevicepibincarnationtable.EntityData.YangName = "qosDevicePibIncarnationTable"
    qosdevicepibincarnationtable.EntityData.BundleName = "cisco_ios_xe"
    qosdevicepibincarnationtable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosdevicepibincarnationtable.EntityData.SegmentPath = "qosDevicePibIncarnationTable"
    qosdevicepibincarnationtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosdevicepibincarnationtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosdevicepibincarnationtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosdevicepibincarnationtable.EntityData.Children = make(map[string]types.YChild)
    qosdevicepibincarnationtable.EntityData.Children["qosDevicePibIncarnationEntry"] = types.YChild{"Qosdevicepibincarnationentry", nil}
    for i := range qosdevicepibincarnationtable.Qosdevicepibincarnationentry {
        qosdevicepibincarnationtable.EntityData.Children[types.GetSegmentPath(&qosdevicepibincarnationtable.Qosdevicepibincarnationentry[i])] = types.YChild{"Qosdevicepibincarnationentry", &qosdevicepibincarnationtable.Qosdevicepibincarnationentry[i]}
    }
    qosdevicepibincarnationtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosdevicepibincarnationtable.EntityData)
}

// CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry
// The single policy instance of this class identifies the
// current incarnation of the PIB and the PDP that installed
// this incarnation.
type CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosdeviceincarnationid interface{}

    // The name of the PDP that installed the current incarnation of the PIB into
    // the device.  By default it is the zero length string. The type is string.
    Qosdevicepdpname interface{}

    // An octet string to identify the current incarnation.  It has meaning to the
    // PDP that installed the PIB and perhaps its standby PDPs. By default the
    // empty string. The type is string with length: 128.
    Qosdevicepibincarnation interface{}

    // The number of seconds after a client close or TCP timeout for which the PEP
    // continues to enforce the policy in the PIB. After this interval, the PIB is
    // consired expired and the device no longer enforces the policy installed in
    // the PIB. The type is interface{} with range: 0..4294967295.
    Qosdevicepibttl interface{}
}

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetEntityData() *types.CommonEntityData {
    qosdevicepibincarnationentry.EntityData.YFilter = qosdevicepibincarnationentry.YFilter
    qosdevicepibincarnationentry.EntityData.YangName = "qosDevicePibIncarnationEntry"
    qosdevicepibincarnationentry.EntityData.BundleName = "cisco_ios_xe"
    qosdevicepibincarnationentry.EntityData.ParentYangName = "qosDevicePibIncarnationTable"
    qosdevicepibincarnationentry.EntityData.SegmentPath = "qosDevicePibIncarnationEntry" + "[qosDeviceIncarnationId='" + fmt.Sprintf("%v", qosdevicepibincarnationentry.Qosdeviceincarnationid) + "']"
    qosdevicepibincarnationentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosdevicepibincarnationentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosdevicepibincarnationentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosdevicepibincarnationentry.EntityData.Children = make(map[string]types.YChild)
    qosdevicepibincarnationentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosdevicepibincarnationentry.EntityData.Leafs["qosDeviceIncarnationId"] = types.YLeaf{"Qosdeviceincarnationid", qosdevicepibincarnationentry.Qosdeviceincarnationid}
    qosdevicepibincarnationentry.EntityData.Leafs["qosDevicePdpName"] = types.YLeaf{"Qosdevicepdpname", qosdevicepibincarnationentry.Qosdevicepdpname}
    qosdevicepibincarnationentry.EntityData.Leafs["qosDevicePibIncarnation"] = types.YLeaf{"Qosdevicepibincarnation", qosdevicepibincarnationentry.Qosdevicepibincarnation}
    qosdevicepibincarnationentry.EntityData.Leafs["qosDevicePibTtl"] = types.YLeaf{"Qosdevicepibttl", qosdevicepibincarnationentry.Qosdevicepibttl}
    return &(qosdevicepibincarnationentry.EntityData)
}

// CISCOQOSPIBMIB_Qosdeviceattributetable
// The single instance of this class indicates specific
// attributes of the device.  These include configuration values
// such as the configured PDP addresses, the maximum message
// size, and specific device capabilities.  The latter include
// input port-based and output port-based classification and/or
// policing, support for flow based policing, aggregate based
// policing, traffic shaping capabilities, etc.
type CISCOQOSPIBMIB_Qosdeviceattributetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The single instance of this class indicates specific attributes of the
    // device. The type is slice of
    // CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry.
    Qosdeviceattributeentry []CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry
}

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetEntityData() *types.CommonEntityData {
    qosdeviceattributetable.EntityData.YFilter = qosdeviceattributetable.YFilter
    qosdeviceattributetable.EntityData.YangName = "qosDeviceAttributeTable"
    qosdeviceattributetable.EntityData.BundleName = "cisco_ios_xe"
    qosdeviceattributetable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosdeviceattributetable.EntityData.SegmentPath = "qosDeviceAttributeTable"
    qosdeviceattributetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosdeviceattributetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosdeviceattributetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosdeviceattributetable.EntityData.Children = make(map[string]types.YChild)
    qosdeviceattributetable.EntityData.Children["qosDeviceAttributeEntry"] = types.YChild{"Qosdeviceattributeentry", nil}
    for i := range qosdeviceattributetable.Qosdeviceattributeentry {
        qosdeviceattributetable.EntityData.Children[types.GetSegmentPath(&qosdeviceattributetable.Qosdeviceattributeentry[i])] = types.YChild{"Qosdeviceattributeentry", &qosdeviceattributetable.Qosdeviceattributeentry[i]}
    }
    qosdeviceattributetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosdeviceattributetable.EntityData)
}

// CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry
// The single instance of this class indicates specific
// attributes of the device.
type CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosdeviceattributeid interface{}

    // The QoS domain that this device belongs to.  This is configured locally on
    // the device (perhaps by some management protocol such as SNMP).  By default,
    // it is the zero-length string. The type is string.
    Qosdevicepepdomain interface{}

    // The address of the PDP configured to be the primary PDP for the device. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Qosdeviceprimarypdp interface{}

    // The address of the PDP configured to be the secondary PDP for the device. 
    // An address of zero indicates no secondary is configured. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Qosdevicesecondarypdp interface{}

    // The maximum size message that this PEP is capable of receiving in bytes.  A
    // value of zero means that the maximum message size is unspecified (but does
    // not mean it is unlimited).  A message greater than this maximum results in
    // a MessageTooBig error on a 'no commit' REP. The type is interface{} with
    // range: 0..4294967295.
    Qosdevicemaxmessagesize interface{}

    // An enumeration of device capabilities.  Used by the PDP to select policies
    // and configuration to push to the PEP. The type is map[string]bool.
    Qosdevicecapabilities interface{}
}

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetEntityData() *types.CommonEntityData {
    qosdeviceattributeentry.EntityData.YFilter = qosdeviceattributeentry.YFilter
    qosdeviceattributeentry.EntityData.YangName = "qosDeviceAttributeEntry"
    qosdeviceattributeentry.EntityData.BundleName = "cisco_ios_xe"
    qosdeviceattributeentry.EntityData.ParentYangName = "qosDeviceAttributeTable"
    qosdeviceattributeentry.EntityData.SegmentPath = "qosDeviceAttributeEntry" + "[qosDeviceAttributeId='" + fmt.Sprintf("%v", qosdeviceattributeentry.Qosdeviceattributeid) + "']"
    qosdeviceattributeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosdeviceattributeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosdeviceattributeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosdeviceattributeentry.EntityData.Children = make(map[string]types.YChild)
    qosdeviceattributeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosdeviceattributeentry.EntityData.Leafs["qosDeviceAttributeId"] = types.YLeaf{"Qosdeviceattributeid", qosdeviceattributeentry.Qosdeviceattributeid}
    qosdeviceattributeentry.EntityData.Leafs["qosDevicePepDomain"] = types.YLeaf{"Qosdevicepepdomain", qosdeviceattributeentry.Qosdevicepepdomain}
    qosdeviceattributeentry.EntityData.Leafs["qosDevicePrimaryPdp"] = types.YLeaf{"Qosdeviceprimarypdp", qosdeviceattributeentry.Qosdeviceprimarypdp}
    qosdeviceattributeentry.EntityData.Leafs["qosDeviceSecondaryPdp"] = types.YLeaf{"Qosdevicesecondarypdp", qosdeviceattributeentry.Qosdevicesecondarypdp}
    qosdeviceattributeentry.EntityData.Leafs["qosDeviceMaxMessageSize"] = types.YLeaf{"Qosdevicemaxmessagesize", qosdeviceattributeentry.Qosdevicemaxmessagesize}
    qosdeviceattributeentry.EntityData.Leafs["qosDeviceCapabilities"] = types.YLeaf{"Qosdevicecapabilities", qosdeviceattributeentry.Qosdevicecapabilities}
    return &(qosdeviceattributeentry.EntityData)
}

// CISCOQOSPIBMIB_Qosinterfacetypetable
// This class describes the interface types of the interfaces
// that exist on the device.  It includes the queue type, role
// combination and capabilities of interfaces.  The PEP does not
// report which specific interfaces have which characteristics.
type CISCOQOSPIBMIB_Qosinterfacetypetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class describes a role combination for an interface
    // type of an interface that exists on the device. The type is slice of
    // CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry.
    Qosinterfacetypeentry []CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry
}

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetEntityData() *types.CommonEntityData {
    qosinterfacetypetable.EntityData.YFilter = qosinterfacetypetable.YFilter
    qosinterfacetypetable.EntityData.YangName = "qosInterfaceTypeTable"
    qosinterfacetypetable.EntityData.BundleName = "cisco_ios_xe"
    qosinterfacetypetable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosinterfacetypetable.EntityData.SegmentPath = "qosInterfaceTypeTable"
    qosinterfacetypetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosinterfacetypetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosinterfacetypetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosinterfacetypetable.EntityData.Children = make(map[string]types.YChild)
    qosinterfacetypetable.EntityData.Children["qosInterfaceTypeEntry"] = types.YChild{"Qosinterfacetypeentry", nil}
    for i := range qosinterfacetypetable.Qosinterfacetypeentry {
        qosinterfacetypetable.EntityData.Children[types.GetSegmentPath(&qosinterfacetypetable.Qosinterfacetypeentry[i])] = types.YChild{"Qosinterfacetypeentry", &qosinterfacetypetable.Qosinterfacetypeentry[i]}
    }
    qosinterfacetypetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosinterfacetypetable.EntityData)
}

// CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry
// An instance of this class describes a role combination for
// an interface type of an interface that exists on the device.
type CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosinterfacetypeid interface{}

    // The interface type in terms of number of queues and thresholds. The type is
    // QosInterfaceQueueType.
    Qosinterfacequeuetype interface{}

    // A combination of roles on at least one interface of type qosInterfaceType.
    // The type is string with length: 0..255.
    Qosinterfacetyperoles interface{}

    // An enumeration of interface capabilities.  Used by the PDP to select
    // policies and configuration to push to the PEP. The type is map[string]bool.
    Qosinterfacetypecapabilities interface{}
}

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetEntityData() *types.CommonEntityData {
    qosinterfacetypeentry.EntityData.YFilter = qosinterfacetypeentry.YFilter
    qosinterfacetypeentry.EntityData.YangName = "qosInterfaceTypeEntry"
    qosinterfacetypeentry.EntityData.BundleName = "cisco_ios_xe"
    qosinterfacetypeentry.EntityData.ParentYangName = "qosInterfaceTypeTable"
    qosinterfacetypeentry.EntityData.SegmentPath = "qosInterfaceTypeEntry" + "[qosInterfaceTypeId='" + fmt.Sprintf("%v", qosinterfacetypeentry.Qosinterfacetypeid) + "']"
    qosinterfacetypeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosinterfacetypeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosinterfacetypeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosinterfacetypeentry.EntityData.Children = make(map[string]types.YChild)
    qosinterfacetypeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosinterfacetypeentry.EntityData.Leafs["qosInterfaceTypeId"] = types.YLeaf{"Qosinterfacetypeid", qosinterfacetypeentry.Qosinterfacetypeid}
    qosinterfacetypeentry.EntityData.Leafs["qosInterfaceQueueType"] = types.YLeaf{"Qosinterfacequeuetype", qosinterfacetypeentry.Qosinterfacequeuetype}
    qosinterfacetypeentry.EntityData.Leafs["qosInterfaceTypeRoles"] = types.YLeaf{"Qosinterfacetyperoles", qosinterfacetypeentry.Qosinterfacetyperoles}
    qosinterfacetypeentry.EntityData.Leafs["qosInterfaceTypeCapabilities"] = types.YLeaf{"Qosinterfacetypecapabilities", qosinterfacetypeentry.Qosinterfacetypecapabilities}
    return &(qosinterfacetypeentry.EntityData)
}

// CISCOQOSPIBMIB_Qosdiffservmappingtable
// Maps each DSCP to a marked-down DSCP.  Also maps each DSCP to
// an IP precedence and QosLayer2Cos.  When configured for the
// first time, all 64 entries of the table must be
// specified. Thereafter, instances may be modified (with a
// delete and install in a single decision) but not deleted
// unless all instances are deleted.
type CISCOQOSPIBMIB_Qosdiffservmappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class represents mappings from a DSCP. The type is
    // slice of CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry.
    Qosdiffservmappingentry []CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry
}

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetEntityData() *types.CommonEntityData {
    qosdiffservmappingtable.EntityData.YFilter = qosdiffservmappingtable.YFilter
    qosdiffservmappingtable.EntityData.YangName = "qosDiffServMappingTable"
    qosdiffservmappingtable.EntityData.BundleName = "cisco_ios_xe"
    qosdiffservmappingtable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosdiffservmappingtable.EntityData.SegmentPath = "qosDiffServMappingTable"
    qosdiffservmappingtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosdiffservmappingtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosdiffservmappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosdiffservmappingtable.EntityData.Children = make(map[string]types.YChild)
    qosdiffservmappingtable.EntityData.Children["qosDiffServMappingEntry"] = types.YChild{"Qosdiffservmappingentry", nil}
    for i := range qosdiffservmappingtable.Qosdiffservmappingentry {
        qosdiffservmappingtable.EntityData.Children[types.GetSegmentPath(&qosdiffservmappingtable.Qosdiffservmappingentry[i])] = types.YChild{"Qosdiffservmappingentry", &qosdiffservmappingtable.Qosdiffservmappingentry[i]}
    }
    qosdiffservmappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosdiffservmappingtable.EntityData)
}

// CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry
// An instance of this class represents mappings from a DSCP.
type CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A DSCP for which this entry contains mappings. The
    // type is interface{} with range: 0..63.
    Qosdscp interface{}

    // The DSCP to use instead of the qosDscp when the packet is out of profile
    // and hence marked as such. The type is interface{} with range: 0..63.
    Qosmarkeddscp interface{}

    // The L2 CoS value to use when mapping this DSCP to layer 2 CoS. The type is
    // interface{} with range: 0..7.
    Qosl2Cos interface{}
}

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetEntityData() *types.CommonEntityData {
    qosdiffservmappingentry.EntityData.YFilter = qosdiffservmappingentry.YFilter
    qosdiffservmappingentry.EntityData.YangName = "qosDiffServMappingEntry"
    qosdiffservmappingentry.EntityData.BundleName = "cisco_ios_xe"
    qosdiffservmappingentry.EntityData.ParentYangName = "qosDiffServMappingTable"
    qosdiffservmappingentry.EntityData.SegmentPath = "qosDiffServMappingEntry" + "[qosDscp='" + fmt.Sprintf("%v", qosdiffservmappingentry.Qosdscp) + "']"
    qosdiffservmappingentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosdiffservmappingentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosdiffservmappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosdiffservmappingentry.EntityData.Children = make(map[string]types.YChild)
    qosdiffservmappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosdiffservmappingentry.EntityData.Leafs["qosDscp"] = types.YLeaf{"Qosdscp", qosdiffservmappingentry.Qosdscp}
    qosdiffservmappingentry.EntityData.Leafs["qosMarkedDscp"] = types.YLeaf{"Qosmarkeddscp", qosdiffservmappingentry.Qosmarkeddscp}
    qosdiffservmappingentry.EntityData.Leafs["qosL2Cos"] = types.YLeaf{"Qosl2Cos", qosdiffservmappingentry.Qosl2Cos}
    return &(qosdiffservmappingentry.EntityData)
}

// CISCOQOSPIBMIB_Qoscostodscptable
// Maps each of eight CoS values to a DSCP.  When configured for
// the first time, all 8 entries of the table must be
// specified. Thereafter, instances may be modified (with a
// delete and install in a single decision) but not deleted
// unless all instances are deleted.
type CISCOQOSPIBMIB_Qoscostodscptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class maps a CoS value to a DSCP. The type is slice of
    // CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry.
    Qoscostodscpentry []CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry
}

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetEntityData() *types.CommonEntityData {
    qoscostodscptable.EntityData.YFilter = qoscostodscptable.YFilter
    qoscostodscptable.EntityData.YangName = "qosCosToDscpTable"
    qoscostodscptable.EntityData.BundleName = "cisco_ios_xe"
    qoscostodscptable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qoscostodscptable.EntityData.SegmentPath = "qosCosToDscpTable"
    qoscostodscptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qoscostodscptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qoscostodscptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qoscostodscptable.EntityData.Children = make(map[string]types.YChild)
    qoscostodscptable.EntityData.Children["qosCosToDscpEntry"] = types.YChild{"Qoscostodscpentry", nil}
    for i := range qoscostodscptable.Qoscostodscpentry {
        qoscostodscptable.EntityData.Children[types.GetSegmentPath(&qoscostodscptable.Qoscostodscpentry[i])] = types.YChild{"Qoscostodscpentry", &qoscostodscptable.Qoscostodscpentry[i]}
    }
    qoscostodscptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qoscostodscptable.EntityData)
}

// CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry
// An instance of this class maps a CoS value to a DSCP.
type CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The L2 CoS value that is being mapped. The type is
    // interface{} with range: 0..7.
    Qoscostodscpcos interface{}

    // The DSCP value to use when mapping the L2 CoS to a DSCP. The type is
    // interface{} with range: 0..63.
    Qoscostodscpdscp interface{}
}

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetEntityData() *types.CommonEntityData {
    qoscostodscpentry.EntityData.YFilter = qoscostodscpentry.YFilter
    qoscostodscpentry.EntityData.YangName = "qosCosToDscpEntry"
    qoscostodscpentry.EntityData.BundleName = "cisco_ios_xe"
    qoscostodscpentry.EntityData.ParentYangName = "qosCosToDscpTable"
    qoscostodscpentry.EntityData.SegmentPath = "qosCosToDscpEntry" + "[qosCosToDscpCos='" + fmt.Sprintf("%v", qoscostodscpentry.Qoscostodscpcos) + "']"
    qoscostodscpentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qoscostodscpentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qoscostodscpentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qoscostodscpentry.EntityData.Children = make(map[string]types.YChild)
    qoscostodscpentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qoscostodscpentry.EntityData.Leafs["qosCosToDscpCos"] = types.YLeaf{"Qoscostodscpcos", qoscostodscpentry.Qoscostodscpcos}
    qoscostodscpentry.EntityData.Leafs["qosCosToDscpDscp"] = types.YLeaf{"Qoscostodscpdscp", qoscostodscpentry.Qoscostodscpdscp}
    return &(qoscostodscpentry.EntityData)
}

// CISCOQOSPIBMIB_Qosunmatchedpolicytable
// A policy class that specifies what QoS to apply to a packet
// that does not match any other policy configured for this role
// combination for a particular direction of traffic.
type CISCOQOSPIBMIB_Qosunmatchedpolicytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the unmatched policy for a particular
    // role combination for incoming or outgoing traffic. The type is slice of
    // CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry.
    Qosunmatchedpolicyentry []CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry
}

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetEntityData() *types.CommonEntityData {
    qosunmatchedpolicytable.EntityData.YFilter = qosunmatchedpolicytable.YFilter
    qosunmatchedpolicytable.EntityData.YangName = "qosUnmatchedPolicyTable"
    qosunmatchedpolicytable.EntityData.BundleName = "cisco_ios_xe"
    qosunmatchedpolicytable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosunmatchedpolicytable.EntityData.SegmentPath = "qosUnmatchedPolicyTable"
    qosunmatchedpolicytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosunmatchedpolicytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosunmatchedpolicytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosunmatchedpolicytable.EntityData.Children = make(map[string]types.YChild)
    qosunmatchedpolicytable.EntityData.Children["qosUnmatchedPolicyEntry"] = types.YChild{"Qosunmatchedpolicyentry", nil}
    for i := range qosunmatchedpolicytable.Qosunmatchedpolicyentry {
        qosunmatchedpolicytable.EntityData.Children[types.GetSegmentPath(&qosunmatchedpolicytable.Qosunmatchedpolicyentry[i])] = types.YChild{"Qosunmatchedpolicyentry", &qosunmatchedpolicytable.Qosunmatchedpolicyentry[i]}
    }
    qosunmatchedpolicytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosunmatchedpolicytable.EntityData)
}

// CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry
// An instance of this class specifies the unmatched policy
// for a particular role combination for incoming or outgoing
// traffic.
type CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosunmatchedpolicyid interface{}

    // Role combination for which this instance applies. The type is string with
    // length: 0..255.
    Qosunmatchedpolicyrole interface{}

    // The direction of packet flow at the interface in question to which this
    // instance applies. The type is Qosunmatchedpolicydirection.
    Qosunmatchedpolicydirection interface{}

    // The DSCP to classify the unmatched packet with.  This must be specified
    // even if qosUnmatchedPolicyDscpTrusted is true. The type is interface{} with
    // range: 0..63.
    Qosunmatchedpolicydscp interface{}

    // If this attribute is true, then the Dscp associated with the packet is
    // trusted, i.e., it is assumed to have already been set.  In this case, the
    // Dscp is not rewritten with qosUnmatchedPolicyDscp (qosUnmatchedPolicyDscp
    // is ignored) unless this is a non-IP packet and arrives untagged.  The
    // packet is still policed as part of its micro flow and its aggregate flow. 
    // When a trusted action is applied to an input interface, the Dscp (for an IP
    // packet) or CoS (for a non-IP packet) associated with the packet is the one
    // contained in the packet. When a trusted action is applied to an output
    // interface, the Dscp associated with the packet is the one that is the
    // result of the input classification and policing. The type is bool.
    Qosunmatchedpolicydscptrusted interface{}

    // An index identifying the instance of policer to apply to unmatched packets.
    // It must correspond to the integer index of an instance of class
    // qosPolicerTable or be zero.  If zero, the microflow is not policed. The
    // type is interface{} with range: 0..4294967295.
    Qosunmatchpolmicroflowpolicerid interface{}

    // An index identifying the aggregate that the packet belongs to.  It must
    // correspond to the integer index of an instance of class qosAggregateTable
    // or be zero.  If zero, the microflow does not belong to any aggregate and is
    // not policed as part of any aggregate. The type is interface{} with range:
    // 0..4294967295.
    Qosunmatchedpolicyaggregateid interface{}
}

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetEntityData() *types.CommonEntityData {
    qosunmatchedpolicyentry.EntityData.YFilter = qosunmatchedpolicyentry.YFilter
    qosunmatchedpolicyentry.EntityData.YangName = "qosUnmatchedPolicyEntry"
    qosunmatchedpolicyentry.EntityData.BundleName = "cisco_ios_xe"
    qosunmatchedpolicyentry.EntityData.ParentYangName = "qosUnmatchedPolicyTable"
    qosunmatchedpolicyentry.EntityData.SegmentPath = "qosUnmatchedPolicyEntry" + "[qosUnmatchedPolicyId='" + fmt.Sprintf("%v", qosunmatchedpolicyentry.Qosunmatchedpolicyid) + "']"
    qosunmatchedpolicyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosunmatchedpolicyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosunmatchedpolicyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosunmatchedpolicyentry.EntityData.Children = make(map[string]types.YChild)
    qosunmatchedpolicyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosunmatchedpolicyentry.EntityData.Leafs["qosUnmatchedPolicyId"] = types.YLeaf{"Qosunmatchedpolicyid", qosunmatchedpolicyentry.Qosunmatchedpolicyid}
    qosunmatchedpolicyentry.EntityData.Leafs["qosUnmatchedPolicyRole"] = types.YLeaf{"Qosunmatchedpolicyrole", qosunmatchedpolicyentry.Qosunmatchedpolicyrole}
    qosunmatchedpolicyentry.EntityData.Leafs["qosUnmatchedPolicyDirection"] = types.YLeaf{"Qosunmatchedpolicydirection", qosunmatchedpolicyentry.Qosunmatchedpolicydirection}
    qosunmatchedpolicyentry.EntityData.Leafs["qosUnmatchedPolicyDscp"] = types.YLeaf{"Qosunmatchedpolicydscp", qosunmatchedpolicyentry.Qosunmatchedpolicydscp}
    qosunmatchedpolicyentry.EntityData.Leafs["qosUnmatchedPolicyDscpTrusted"] = types.YLeaf{"Qosunmatchedpolicydscptrusted", qosunmatchedpolicyentry.Qosunmatchedpolicydscptrusted}
    qosunmatchedpolicyentry.EntityData.Leafs["qosUnmatchPolMicroFlowPolicerId"] = types.YLeaf{"Qosunmatchpolmicroflowpolicerid", qosunmatchedpolicyentry.Qosunmatchpolmicroflowpolicerid}
    qosunmatchedpolicyentry.EntityData.Leafs["qosUnmatchedPolicyAggregateId"] = types.YLeaf{"Qosunmatchedpolicyaggregateid", qosunmatchedpolicyentry.Qosunmatchedpolicyaggregateid}
    return &(qosunmatchedpolicyentry.EntityData)
}

// CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry_Qosunmatchedpolicydirection represents which this instance applies.
type CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry_Qosunmatchedpolicydirection string

const (
    CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry_Qosunmatchedpolicydirection_in CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry_Qosunmatchedpolicydirection = "in"

    CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry_Qosunmatchedpolicydirection_out CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry_Qosunmatchedpolicydirection = "out"
)

// CISCOQOSPIBMIB_Qospolicertable
// A class specifying policing parameters for both microflows
// and aggregate flows.  This table is designed for policing
// according to a token bucket scheme where an average rate and
// burst size is specified.
type CISCOQOSPIBMIB_Qospolicertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies a set of policing parameters. The type
    // is slice of CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry.
    Qospolicerentry []CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry
}

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetEntityData() *types.CommonEntityData {
    qospolicertable.EntityData.YFilter = qospolicertable.YFilter
    qospolicertable.EntityData.YangName = "qosPolicerTable"
    qospolicertable.EntityData.BundleName = "cisco_ios_xe"
    qospolicertable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qospolicertable.EntityData.SegmentPath = "qosPolicerTable"
    qospolicertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qospolicertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qospolicertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qospolicertable.EntityData.Children = make(map[string]types.YChild)
    qospolicertable.EntityData.Children["qosPolicerEntry"] = types.YChild{"Qospolicerentry", nil}
    for i := range qospolicertable.Qospolicerentry {
        qospolicertable.EntityData.Children[types.GetSegmentPath(&qospolicertable.Qospolicerentry[i])] = types.YChild{"Qospolicerentry", &qospolicertable.Qospolicerentry[i]}
    }
    qospolicertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qospolicertable.EntityData)
}

// CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry
// An instance of this class specifies a set of policing
// parameters.
type CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qospolicerid interface{}

    // The token rate.  It is specified in units of bit/s. A rate of zero means
    // that all packets will be out of profile.  If the qosPolicerAction is set to
    // drop then this effectively denies any service to packets policed by this
    // policer. The type is interface{} with range: 0..4294967295.
    Qospolicerrate interface{}

    // The normal size of a burst in terms of bits. The type is interface{} with
    // range: 0..4294967295.
    Qospolicernormalburst interface{}

    // The excess size of a burst in terms of bits. The type is interface{} with
    // range: 0..4294967295.
    Qospolicerexcessburst interface{}

    // An indication of how to handle out of profile packets.  When the shape
    // action is chosen then traffic is shaped to the rate specified by
    // qosPolicerRate. The type is Qospoliceraction.
    Qospoliceraction interface{}
}

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetEntityData() *types.CommonEntityData {
    qospolicerentry.EntityData.YFilter = qospolicerentry.YFilter
    qospolicerentry.EntityData.YangName = "qosPolicerEntry"
    qospolicerentry.EntityData.BundleName = "cisco_ios_xe"
    qospolicerentry.EntityData.ParentYangName = "qosPolicerTable"
    qospolicerentry.EntityData.SegmentPath = "qosPolicerEntry" + "[qosPolicerId='" + fmt.Sprintf("%v", qospolicerentry.Qospolicerid) + "']"
    qospolicerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qospolicerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qospolicerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qospolicerentry.EntityData.Children = make(map[string]types.YChild)
    qospolicerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qospolicerentry.EntityData.Leafs["qosPolicerId"] = types.YLeaf{"Qospolicerid", qospolicerentry.Qospolicerid}
    qospolicerentry.EntityData.Leafs["qosPolicerRate"] = types.YLeaf{"Qospolicerrate", qospolicerentry.Qospolicerrate}
    qospolicerentry.EntityData.Leafs["qosPolicerNormalBurst"] = types.YLeaf{"Qospolicernormalburst", qospolicerentry.Qospolicernormalburst}
    qospolicerentry.EntityData.Leafs["qosPolicerExcessBurst"] = types.YLeaf{"Qospolicerexcessburst", qospolicerentry.Qospolicerexcessburst}
    qospolicerentry.EntityData.Leafs["qosPolicerAction"] = types.YLeaf{"Qospoliceraction", qospolicerentry.Qospoliceraction}
    return &(qospolicerentry.EntityData)
}

// CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry_Qospoliceraction represents specified by qosPolicerRate.
type CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry_Qospoliceraction string

const (
    CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry_Qospoliceraction_drop CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry_Qospoliceraction = "drop"

    CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry_Qospoliceraction_mark CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry_Qospoliceraction = "mark"

    CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry_Qospoliceraction_shape CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry_Qospoliceraction = "shape"
)

// CISCOQOSPIBMIB_Qosaggregatetable
// Instances of this class identify aggregate flows and the
// policer to apply to each.
type CISCOQOSPIBMIB_Qosaggregatetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the policer to apply to an aggregate
    // flow. The type is slice of
    // CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry.
    Qosaggregateentry []CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry
}

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetEntityData() *types.CommonEntityData {
    qosaggregatetable.EntityData.YFilter = qosaggregatetable.YFilter
    qosaggregatetable.EntityData.YangName = "qosAggregateTable"
    qosaggregatetable.EntityData.BundleName = "cisco_ios_xe"
    qosaggregatetable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosaggregatetable.EntityData.SegmentPath = "qosAggregateTable"
    qosaggregatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosaggregatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosaggregatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosaggregatetable.EntityData.Children = make(map[string]types.YChild)
    qosaggregatetable.EntityData.Children["qosAggregateEntry"] = types.YChild{"Qosaggregateentry", nil}
    for i := range qosaggregatetable.Qosaggregateentry {
        qosaggregatetable.EntityData.Children[types.GetSegmentPath(&qosaggregatetable.Qosaggregateentry[i])] = types.YChild{"Qosaggregateentry", &qosaggregatetable.Qosaggregateentry[i]}
    }
    qosaggregatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosaggregatetable.EntityData)
}

// CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry
// An instance of this class specifies the policer to apply to
// an aggregate flow.
type CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosaggregateid interface{}

    // An index identifying the instance of policer to apply to the aggregate.  It
    // must correspond to the integer index of an instance of class
    // qosPolicerTable. The type is interface{} with range: 0..4294967295.
    Qosaggregatepolicerid interface{}
}

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetEntityData() *types.CommonEntityData {
    qosaggregateentry.EntityData.YFilter = qosaggregateentry.YFilter
    qosaggregateentry.EntityData.YangName = "qosAggregateEntry"
    qosaggregateentry.EntityData.BundleName = "cisco_ios_xe"
    qosaggregateentry.EntityData.ParentYangName = "qosAggregateTable"
    qosaggregateentry.EntityData.SegmentPath = "qosAggregateEntry" + "[qosAggregateId='" + fmt.Sprintf("%v", qosaggregateentry.Qosaggregateid) + "']"
    qosaggregateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosaggregateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosaggregateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosaggregateentry.EntityData.Children = make(map[string]types.YChild)
    qosaggregateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosaggregateentry.EntityData.Leafs["qosAggregateId"] = types.YLeaf{"Qosaggregateid", qosaggregateentry.Qosaggregateid}
    qosaggregateentry.EntityData.Leafs["qosAggregatePolicerId"] = types.YLeaf{"Qosaggregatepolicerid", qosaggregateentry.Qosaggregatepolicerid}
    return &(qosaggregateentry.EntityData)
}

// CISCOQOSPIBMIB_Qosmacclassificationtable
// A class of MAC/Vlan tuples and their associated CoS values.
type CISCOQOSPIBMIB_Qosmacclassificationtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the mapping of a VLAN and a MAC address
    // to a CoS value. The type is slice of
    // CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry.
    Qosmacclassificationentry []CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry
}

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetEntityData() *types.CommonEntityData {
    qosmacclassificationtable.EntityData.YFilter = qosmacclassificationtable.YFilter
    qosmacclassificationtable.EntityData.YangName = "qosMacClassificationTable"
    qosmacclassificationtable.EntityData.BundleName = "cisco_ios_xe"
    qosmacclassificationtable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosmacclassificationtable.EntityData.SegmentPath = "qosMacClassificationTable"
    qosmacclassificationtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosmacclassificationtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosmacclassificationtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosmacclassificationtable.EntityData.Children = make(map[string]types.YChild)
    qosmacclassificationtable.EntityData.Children["qosMacClassificationEntry"] = types.YChild{"Qosmacclassificationentry", nil}
    for i := range qosmacclassificationtable.Qosmacclassificationentry {
        qosmacclassificationtable.EntityData.Children[types.GetSegmentPath(&qosmacclassificationtable.Qosmacclassificationentry[i])] = types.YChild{"Qosmacclassificationentry", &qosmacclassificationtable.Qosmacclassificationentry[i]}
    }
    qosmacclassificationtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosmacclassificationtable.EntityData)
}

// CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry
// An instance of this class specifies the mapping of a VLAN
// and a MAC address to a CoS value.
type CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosmacclassificationid interface{}

    // The VLAN of the destination MAC address of the L2 frame. The type is
    // interface{} with range: 1..4095.
    Qosdstmacvlan interface{}

    // The destination MAC address of the L2 frame. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Qosdstmacaddress interface{}

    // The CoS to assign the packet with the associated MAC/VLAN tuple.  Note that
    // this CoS is overridden by the policies to classify the frame at layer 3 if
    // there are any. The type is interface{} with range: 0..7.
    Qosdstmaccos interface{}
}

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetEntityData() *types.CommonEntityData {
    qosmacclassificationentry.EntityData.YFilter = qosmacclassificationentry.YFilter
    qosmacclassificationentry.EntityData.YangName = "qosMacClassificationEntry"
    qosmacclassificationentry.EntityData.BundleName = "cisco_ios_xe"
    qosmacclassificationentry.EntityData.ParentYangName = "qosMacClassificationTable"
    qosmacclassificationentry.EntityData.SegmentPath = "qosMacClassificationEntry" + "[qosMacClassificationId='" + fmt.Sprintf("%v", qosmacclassificationentry.Qosmacclassificationid) + "']"
    qosmacclassificationentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosmacclassificationentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosmacclassificationentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosmacclassificationentry.EntityData.Children = make(map[string]types.YChild)
    qosmacclassificationentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosmacclassificationentry.EntityData.Leafs["qosMacClassificationId"] = types.YLeaf{"Qosmacclassificationid", qosmacclassificationentry.Qosmacclassificationid}
    qosmacclassificationentry.EntityData.Leafs["qosDstMacVlan"] = types.YLeaf{"Qosdstmacvlan", qosmacclassificationentry.Qosdstmacvlan}
    qosmacclassificationentry.EntityData.Leafs["qosDstMacAddress"] = types.YLeaf{"Qosdstmacaddress", qosmacclassificationentry.Qosdstmacaddress}
    qosmacclassificationentry.EntityData.Leafs["qosDstMacCos"] = types.YLeaf{"Qosdstmaccos", qosmacclassificationentry.Qosdstmaccos}
    return &(qosmacclassificationentry.EntityData)
}

// CISCOQOSPIBMIB_Qosipacetable
// ACE definitions.
type CISCOQOSPIBMIB_Qosipacetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies an ACE. The type is slice of
    // CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry.
    Qosipaceentry []CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry
}

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetEntityData() *types.CommonEntityData {
    qosipacetable.EntityData.YFilter = qosipacetable.YFilter
    qosipacetable.EntityData.YangName = "qosIpAceTable"
    qosipacetable.EntityData.BundleName = "cisco_ios_xe"
    qosipacetable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosipacetable.EntityData.SegmentPath = "qosIpAceTable"
    qosipacetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosipacetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosipacetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosipacetable.EntityData.Children = make(map[string]types.YChild)
    qosipacetable.EntityData.Children["qosIpAceEntry"] = types.YChild{"Qosipaceentry", nil}
    for i := range qosipacetable.Qosipaceentry {
        qosipacetable.EntityData.Children[types.GetSegmentPath(&qosipacetable.Qosipaceentry[i])] = types.YChild{"Qosipaceentry", &qosipacetable.Qosipaceentry[i]}
    }
    qosipacetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosipacetable.EntityData)
}

// CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry
// An instance of this class specifies an ACE.
type CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosipaceid interface{}

    // The IP address to match against the packet's destination IP address. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Qosipacedstaddr interface{}

    // A mask for the matching of the destination IP address. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Qosipacedstaddrmask interface{}

    // The IP address to match against the packet's source IP address. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Qosipacesrcaddr interface{}

    // A mask for the matching of the source IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Qosipacesrcaddrmask interface{}

    // The minimum value that the DSCP in the packet can have and match this ACE.
    // The type is interface{} with range: 0..63.
    Qosipacedscpmin interface{}

    // The maximum value that the DSCP in the packet can have and match this ACE.
    // The type is interface{} with range: 0..63.
    Qosipacedscpmax interface{}

    // The IP protocol to match against the packet's protocol. A value of zero
    // means match all. The type is interface{} with range: 0..255.
    Qosipaceprotocol interface{}

    // The minimum value that the packet's layer 4 dest port number can have and
    // match this ACE. The type is interface{} with range: 0..65535.
    Qosipacedstl4Portmin interface{}

    // The maximum value that the packet's layer 4 dest port number can have and
    // match this ACE. The type is interface{} with range: 0..65535.
    Qosipacedstl4Portmax interface{}

    // The minimum value that the packet's layer 4 source port number can have and
    // match this ACE. The type is interface{} with range: 0..65535.
    Qosipacesrcl4Portmin interface{}

    // The maximum value that the packet's layer 4 source port number can have and
    // match this ACE. The type is interface{} with range: 0..65535.
    Qosipacesrcl4Portmax interface{}

    // If the packet matches this ACE and the value of this attribute is true,
    // then the matching process terminates and the QoS associated with this ACE
    // (indirectly through the ACL) is applied to the packet.  If the value of
    // this attribute is false, then no more ACEs in this ACL are compared to this
    // packet and matching continues with the first ACE of the next ACL. The type
    // is bool.
    Qosipacepermit interface{}
}

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetEntityData() *types.CommonEntityData {
    qosipaceentry.EntityData.YFilter = qosipaceentry.YFilter
    qosipaceentry.EntityData.YangName = "qosIpAceEntry"
    qosipaceentry.EntityData.BundleName = "cisco_ios_xe"
    qosipaceentry.EntityData.ParentYangName = "qosIpAceTable"
    qosipaceentry.EntityData.SegmentPath = "qosIpAceEntry" + "[qosIpAceId='" + fmt.Sprintf("%v", qosipaceentry.Qosipaceid) + "']"
    qosipaceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosipaceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosipaceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosipaceentry.EntityData.Children = make(map[string]types.YChild)
    qosipaceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosipaceentry.EntityData.Leafs["qosIpAceId"] = types.YLeaf{"Qosipaceid", qosipaceentry.Qosipaceid}
    qosipaceentry.EntityData.Leafs["qosIpAceDstAddr"] = types.YLeaf{"Qosipacedstaddr", qosipaceentry.Qosipacedstaddr}
    qosipaceentry.EntityData.Leafs["qosIpAceDstAddrMask"] = types.YLeaf{"Qosipacedstaddrmask", qosipaceentry.Qosipacedstaddrmask}
    qosipaceentry.EntityData.Leafs["qosIpAceSrcAddr"] = types.YLeaf{"Qosipacesrcaddr", qosipaceentry.Qosipacesrcaddr}
    qosipaceentry.EntityData.Leafs["qosIpAceSrcAddrMask"] = types.YLeaf{"Qosipacesrcaddrmask", qosipaceentry.Qosipacesrcaddrmask}
    qosipaceentry.EntityData.Leafs["qosIpAceDscpMin"] = types.YLeaf{"Qosipacedscpmin", qosipaceentry.Qosipacedscpmin}
    qosipaceentry.EntityData.Leafs["qosIpAceDscpMax"] = types.YLeaf{"Qosipacedscpmax", qosipaceentry.Qosipacedscpmax}
    qosipaceentry.EntityData.Leafs["qosIpAceProtocol"] = types.YLeaf{"Qosipaceprotocol", qosipaceentry.Qosipaceprotocol}
    qosipaceentry.EntityData.Leafs["qosIpAceDstL4PortMin"] = types.YLeaf{"Qosipacedstl4Portmin", qosipaceentry.Qosipacedstl4Portmin}
    qosipaceentry.EntityData.Leafs["qosIpAceDstL4PortMax"] = types.YLeaf{"Qosipacedstl4Portmax", qosipaceentry.Qosipacedstl4Portmax}
    qosipaceentry.EntityData.Leafs["qosIpAceSrcL4PortMin"] = types.YLeaf{"Qosipacesrcl4Portmin", qosipaceentry.Qosipacesrcl4Portmin}
    qosipaceentry.EntityData.Leafs["qosIpAceSrcL4PortMax"] = types.YLeaf{"Qosipacesrcl4Portmax", qosipaceentry.Qosipacesrcl4Portmax}
    qosipaceentry.EntityData.Leafs["qosIpAcePermit"] = types.YLeaf{"Qosipacepermit", qosipaceentry.Qosipacepermit}
    return &(qosipaceentry.EntityData)
}

// CISCOQOSPIBMIB_Qosipacldefinitiontable
// A class that defines a set of ACLs each being an ordered list
// of ACEs.
type CISCOQOSPIBMIB_Qosipacldefinitiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies an ACE in an ACL and its order with
    // respect to other ACEs in the same ACL. The type is slice of
    // CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry.
    Qosipacldefinitionentry []CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry
}

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetEntityData() *types.CommonEntityData {
    qosipacldefinitiontable.EntityData.YFilter = qosipacldefinitiontable.YFilter
    qosipacldefinitiontable.EntityData.YangName = "qosIpAclDefinitionTable"
    qosipacldefinitiontable.EntityData.BundleName = "cisco_ios_xe"
    qosipacldefinitiontable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosipacldefinitiontable.EntityData.SegmentPath = "qosIpAclDefinitionTable"
    qosipacldefinitiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosipacldefinitiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosipacldefinitiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosipacldefinitiontable.EntityData.Children = make(map[string]types.YChild)
    qosipacldefinitiontable.EntityData.Children["qosIpAclDefinitionEntry"] = types.YChild{"Qosipacldefinitionentry", nil}
    for i := range qosipacldefinitiontable.Qosipacldefinitionentry {
        qosipacldefinitiontable.EntityData.Children[types.GetSegmentPath(&qosipacldefinitiontable.Qosipacldefinitionentry[i])] = types.YChild{"Qosipacldefinitionentry", &qosipacldefinitiontable.Qosipacldefinitionentry[i]}
    }
    qosipacldefinitiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosipacldefinitiontable.EntityData)
}

// CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry
// An instance of this class specifies an ACE in an ACL and its
// order with respect to other ACEs in the same ACL.
type CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosipacldefinitionid interface{}

    // An index for this ACL.  There will be one instance of policy class
    // qosIpAclDefinition with this integer index for each ACE in the ACL per role
    // combination. The type is interface{} with range: 0..4294967295.
    Qosipaclid interface{}

    // An integer that determines the position of this ACE in the ACL. An ACE with
    // a given order is positioned in the access contol list before one with a
    // higher order. The type is interface{} with range: 0..4294967295.
    Qosipaceorder interface{}

    // This attribute specifies the ACE in the qosIpAceTable that is in the ACL
    // specified by qosIpAclId at the position specified by qosIpAceOrder. The
    // type is interface{} with range: 0..4294967295.
    Qosipacldefaceid interface{}
}

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetEntityData() *types.CommonEntityData {
    qosipacldefinitionentry.EntityData.YFilter = qosipacldefinitionentry.YFilter
    qosipacldefinitionentry.EntityData.YangName = "qosIpAclDefinitionEntry"
    qosipacldefinitionentry.EntityData.BundleName = "cisco_ios_xe"
    qosipacldefinitionentry.EntityData.ParentYangName = "qosIpAclDefinitionTable"
    qosipacldefinitionentry.EntityData.SegmentPath = "qosIpAclDefinitionEntry" + "[qosIpAclDefinitionId='" + fmt.Sprintf("%v", qosipacldefinitionentry.Qosipacldefinitionid) + "']"
    qosipacldefinitionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosipacldefinitionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosipacldefinitionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosipacldefinitionentry.EntityData.Children = make(map[string]types.YChild)
    qosipacldefinitionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosipacldefinitionentry.EntityData.Leafs["qosIpAclDefinitionId"] = types.YLeaf{"Qosipacldefinitionid", qosipacldefinitionentry.Qosipacldefinitionid}
    qosipacldefinitionentry.EntityData.Leafs["qosIpAclId"] = types.YLeaf{"Qosipaclid", qosipacldefinitionentry.Qosipaclid}
    qosipacldefinitionentry.EntityData.Leafs["qosIpAceOrder"] = types.YLeaf{"Qosipaceorder", qosipacldefinitionentry.Qosipaceorder}
    qosipacldefinitionentry.EntityData.Leafs["qosIpAclDefAceId"] = types.YLeaf{"Qosipacldefaceid", qosipacldefinitionentry.Qosipacldefaceid}
    return &(qosipacldefinitionentry.EntityData)
}

// CISCOQOSPIBMIB_Qosipaclactiontable
// A class that applies a set of ACLs to interfaces specifying,
// for each interface the order of the ACL with respect to other
// ACLs applied to the same interface and, for each ACL the
// action to take for a packet that matches a permit ACE in that
// ACL.  Interfaces are specified abstractly in terms of
// interface role combinations.
type CISCOQOSPIBMIB_Qosipaclactiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class applies an ACL to traffic in a particular
    // direction on an interface with a particular role combination, and specifies
    // the action for packets which match the ACL. The type is slice of
    // CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry.
    Qosipaclactionentry []CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry
}

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetEntityData() *types.CommonEntityData {
    qosipaclactiontable.EntityData.YFilter = qosipaclactiontable.YFilter
    qosipaclactiontable.EntityData.YangName = "qosIpAclActionTable"
    qosipaclactiontable.EntityData.BundleName = "cisco_ios_xe"
    qosipaclactiontable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosipaclactiontable.EntityData.SegmentPath = "qosIpAclActionTable"
    qosipaclactiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosipaclactiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosipaclactiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosipaclactiontable.EntityData.Children = make(map[string]types.YChild)
    qosipaclactiontable.EntityData.Children["qosIpAclActionEntry"] = types.YChild{"Qosipaclactionentry", nil}
    for i := range qosipaclactiontable.Qosipaclactionentry {
        qosipaclactiontable.EntityData.Children[types.GetSegmentPath(&qosipaclactiontable.Qosipaclactionentry[i])] = types.YChild{"Qosipaclactionentry", &qosipaclactiontable.Qosipaclactionentry[i]}
    }
    qosipaclactiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosipaclactiontable.EntityData)
}

// CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry
// An instance of this class applies an ACL to traffic in a
// particular direction on an interface with a particular role
// combination, and specifies the action for packets which match
// the ACL.
type CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosipaclactionid interface{}

    // The ACL associated with this action. The type is interface{} with range:
    // 0..4294967295.
    Qosipaclactaclid interface{}

    // The interfaces to which this ACL applies specified in terms of a set of
    // roles. The type is string with length: 0..255.
    Qosipaclinterfaceroles interface{}

    // The direction of packet flow at the interface in question to which this ACL
    // applies. The type is Qosipaclinterfacedirection.
    Qosipaclinterfacedirection interface{}

    // An integer that determines the order of this ACL in the list of ACLs
    // applied to interfaces of the specified role combination. An ACL with a
    // given order is positioned in the list before one with a higher order. The
    // type is interface{} with range: 0..4294967295.
    Qosipaclorder interface{}

    // The DSCP to classify the packet with in the event that the packet matches
    // an ACE in this ACL and the ACE is a permit. The type is interface{} with
    // range: 0..63.
    Qosipacldscp interface{}

    // If this attribute is true, then the Dscp associated with the packet is
    // trusted, i.e., it is assumed to have already been set.  In this case, the
    // Dscp is not rewritten with qosIpAclDscp (qosIpAclDscp is ignored).  The
    // packet is still policed as part of its micro flow and its aggregate flow. 
    // When a trusted action is applied to an input interface, the Dscp associated
    // with the packet is the one contained in the packet.  When a trusted action
    // is applied to an output interface, the Dscp associated with the packet is
    // the one that is the result of the input classification and policing. The
    // type is bool.
    Qosipacldscptrusted interface{}

    // An index identifying the instance of policer to apply to the microflow.  It
    // must correspond to the integer index of an instance of class
    // qosPolicerTableor be zero.  If zero, the microflow is not policed. The type
    // is interface{} with range: 0..4294967295.
    Qosipaclmicroflowpolicerid interface{}

    // An index identifying the aggregate that the packet belongs to.  It must
    // correspond to the integer index of an instance of class qosAggregateTable
    // or be zero.  If zero, the microflow does not belong to any aggregate and is
    // not policed as part of any aggregate. The type is interface{} with range:
    // 0..4294967295.
    Qosipaclaggregateid interface{}
}

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetEntityData() *types.CommonEntityData {
    qosipaclactionentry.EntityData.YFilter = qosipaclactionentry.YFilter
    qosipaclactionentry.EntityData.YangName = "qosIpAclActionEntry"
    qosipaclactionentry.EntityData.BundleName = "cisco_ios_xe"
    qosipaclactionentry.EntityData.ParentYangName = "qosIpAclActionTable"
    qosipaclactionentry.EntityData.SegmentPath = "qosIpAclActionEntry" + "[qosIpAclActionId='" + fmt.Sprintf("%v", qosipaclactionentry.Qosipaclactionid) + "']"
    qosipaclactionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosipaclactionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosipaclactionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosipaclactionentry.EntityData.Children = make(map[string]types.YChild)
    qosipaclactionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosipaclactionentry.EntityData.Leafs["qosIpAclActionId"] = types.YLeaf{"Qosipaclactionid", qosipaclactionentry.Qosipaclactionid}
    qosipaclactionentry.EntityData.Leafs["qosIpAclActAclId"] = types.YLeaf{"Qosipaclactaclid", qosipaclactionentry.Qosipaclactaclid}
    qosipaclactionentry.EntityData.Leafs["qosIpAclInterfaceRoles"] = types.YLeaf{"Qosipaclinterfaceroles", qosipaclactionentry.Qosipaclinterfaceroles}
    qosipaclactionentry.EntityData.Leafs["qosIpAclInterfaceDirection"] = types.YLeaf{"Qosipaclinterfacedirection", qosipaclactionentry.Qosipaclinterfacedirection}
    qosipaclactionentry.EntityData.Leafs["qosIpAclOrder"] = types.YLeaf{"Qosipaclorder", qosipaclactionentry.Qosipaclorder}
    qosipaclactionentry.EntityData.Leafs["qosIpAclDscp"] = types.YLeaf{"Qosipacldscp", qosipaclactionentry.Qosipacldscp}
    qosipaclactionentry.EntityData.Leafs["qosIpAclDscpTrusted"] = types.YLeaf{"Qosipacldscptrusted", qosipaclactionentry.Qosipacldscptrusted}
    qosipaclactionentry.EntityData.Leafs["qosIpAclMicroFlowPolicerId"] = types.YLeaf{"Qosipaclmicroflowpolicerid", qosipaclactionentry.Qosipaclmicroflowpolicerid}
    qosipaclactionentry.EntityData.Leafs["qosIpAclAggregateId"] = types.YLeaf{"Qosipaclaggregateid", qosipaclactionentry.Qosipaclaggregateid}
    return &(qosipaclactionentry.EntityData)
}

// CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry_Qosipaclinterfacedirection represents which this ACL applies.
type CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry_Qosipaclinterfacedirection string

const (
    CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry_Qosipaclinterfacedirection_in CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry_Qosipaclinterfacedirection = "in"

    CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry_Qosipaclinterfacedirection_out CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry_Qosipaclinterfacedirection = "out"
)

// CISCOQOSPIBMIB_Qosifschedulingpreferencestable
// This class specifies the scheduling preference an interface
// chooses if it supports multiple scheduling types.  Higher
// values are preferred over lower values.
type CISCOQOSPIBMIB_Qosifschedulingpreferencestable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies a scheduling preference for a
    // queue-type on an interface with a particular role combination. The type is
    // slice of
    // CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry.
    Qosifschedulingpreferenceentry []CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry
}

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetEntityData() *types.CommonEntityData {
    qosifschedulingpreferencestable.EntityData.YFilter = qosifschedulingpreferencestable.YFilter
    qosifschedulingpreferencestable.EntityData.YangName = "qosIfSchedulingPreferencesTable"
    qosifschedulingpreferencestable.EntityData.BundleName = "cisco_ios_xe"
    qosifschedulingpreferencestable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosifschedulingpreferencestable.EntityData.SegmentPath = "qosIfSchedulingPreferencesTable"
    qosifschedulingpreferencestable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifschedulingpreferencestable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifschedulingpreferencestable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifschedulingpreferencestable.EntityData.Children = make(map[string]types.YChild)
    qosifschedulingpreferencestable.EntityData.Children["qosIfSchedulingPreferenceEntry"] = types.YChild{"Qosifschedulingpreferenceentry", nil}
    for i := range qosifschedulingpreferencestable.Qosifschedulingpreferenceentry {
        qosifschedulingpreferencestable.EntityData.Children[types.GetSegmentPath(&qosifschedulingpreferencestable.Qosifschedulingpreferenceentry[i])] = types.YChild{"Qosifschedulingpreferenceentry", &qosifschedulingpreferencestable.Qosifschedulingpreferenceentry[i]}
    }
    qosifschedulingpreferencestable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosifschedulingpreferencestable.EntityData)
}

// CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry
// An instance of this class specifies a scheduling preference
// for a queue-type on an interface with a particular role
// combination.
type CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosifschedulingpreferenceid interface{}

    // The combination of roles the interface must have for this policy instance
    // to apply to that interface. The type is string with length: 0..255.
    Qosifschedulingroles interface{}

    // The preference to use this scheduling discipline and queue type.  A higher
    // value means a higher preference.  If two disciplines have the same
    // preference the choice is a local decision. The type is interface{} with
    // range: 1..16.
    Qosifschedulingpreference interface{}

    // An enumerate type for all the known scheduling disciplines. The type is
    // Qosifschedulingdiscipline.
    Qosifschedulingdiscipline interface{}

    // The queue type of this preference. The type is QosInterfaceQueueType.
    Qosifschedulingqueuetype interface{}
}

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetEntityData() *types.CommonEntityData {
    qosifschedulingpreferenceentry.EntityData.YFilter = qosifschedulingpreferenceentry.YFilter
    qosifschedulingpreferenceentry.EntityData.YangName = "qosIfSchedulingPreferenceEntry"
    qosifschedulingpreferenceentry.EntityData.BundleName = "cisco_ios_xe"
    qosifschedulingpreferenceentry.EntityData.ParentYangName = "qosIfSchedulingPreferencesTable"
    qosifschedulingpreferenceentry.EntityData.SegmentPath = "qosIfSchedulingPreferenceEntry" + "[qosIfSchedulingPreferenceId='" + fmt.Sprintf("%v", qosifschedulingpreferenceentry.Qosifschedulingpreferenceid) + "']"
    qosifschedulingpreferenceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifschedulingpreferenceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifschedulingpreferenceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifschedulingpreferenceentry.EntityData.Children = make(map[string]types.YChild)
    qosifschedulingpreferenceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosifschedulingpreferenceentry.EntityData.Leafs["qosIfSchedulingPreferenceId"] = types.YLeaf{"Qosifschedulingpreferenceid", qosifschedulingpreferenceentry.Qosifschedulingpreferenceid}
    qosifschedulingpreferenceentry.EntityData.Leafs["qosIfSchedulingRoles"] = types.YLeaf{"Qosifschedulingroles", qosifschedulingpreferenceentry.Qosifschedulingroles}
    qosifschedulingpreferenceentry.EntityData.Leafs["qosIfSchedulingPreference"] = types.YLeaf{"Qosifschedulingpreference", qosifschedulingpreferenceentry.Qosifschedulingpreference}
    qosifschedulingpreferenceentry.EntityData.Leafs["qosIfSchedulingDiscipline"] = types.YLeaf{"Qosifschedulingdiscipline", qosifschedulingpreferenceentry.Qosifschedulingdiscipline}
    qosifschedulingpreferenceentry.EntityData.Leafs["qosIfSchedulingQueueType"] = types.YLeaf{"Qosifschedulingqueuetype", qosifschedulingpreferenceentry.Qosifschedulingqueuetype}
    return &(qosifschedulingpreferenceentry.EntityData)
}

// CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline represents An enumerate type for all the known scheduling disciplines.
type CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline string

const (
    CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline_weightedFairQueueing CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline = "weightedFairQueueing"

    CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline_weightedRoundRobin CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline = "weightedRoundRobin"

    CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline_customQueueing CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline = "customQueueing"

    CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline_priorityQueueing CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline = "priorityQueueing"

    CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline_classBasedWFQ CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline = "classBasedWFQ"

    CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline_fifo CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline = "fifo"

    CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline_pqWrr CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline = "pqWrr"

    CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline_pqCbwfq CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry_Qosifschedulingdiscipline = "pqCbwfq"
)

// CISCOQOSPIBMIB_Qosifdroppreferencetable
// This class specifies the preference of the drop mechanism an
// interface chooses if it supports multiple drop mechanisms.
// Higher values are preferred over lower values.
type CISCOQOSPIBMIB_Qosifdroppreferencetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies a drop preference for a drop mechanism
    // on an interface with a particular role combination. The type is slice of
    // CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry.
    Qosifdroppreferenceentry []CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry
}

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetEntityData() *types.CommonEntityData {
    qosifdroppreferencetable.EntityData.YFilter = qosifdroppreferencetable.YFilter
    qosifdroppreferencetable.EntityData.YangName = "qosIfDropPreferenceTable"
    qosifdroppreferencetable.EntityData.BundleName = "cisco_ios_xe"
    qosifdroppreferencetable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosifdroppreferencetable.EntityData.SegmentPath = "qosIfDropPreferenceTable"
    qosifdroppreferencetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifdroppreferencetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifdroppreferencetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifdroppreferencetable.EntityData.Children = make(map[string]types.YChild)
    qosifdroppreferencetable.EntityData.Children["qosIfDropPreferenceEntry"] = types.YChild{"Qosifdroppreferenceentry", nil}
    for i := range qosifdroppreferencetable.Qosifdroppreferenceentry {
        qosifdroppreferencetable.EntityData.Children[types.GetSegmentPath(&qosifdroppreferencetable.Qosifdroppreferenceentry[i])] = types.YChild{"Qosifdroppreferenceentry", &qosifdroppreferencetable.Qosifdroppreferenceentry[i]}
    }
    qosifdroppreferencetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosifdroppreferencetable.EntityData)
}

// CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry
// An instance of this class specifies a drop preference for
// a drop mechanism on an interface with a particular role
// combination.
type CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosifdroppreferenceid interface{}

    // The combination of roles the interface must have for this policy instance
    // to apply to that interface. The type is string with length: 0..255.
    Qosifdroproles interface{}

    // The preference to use this drop mechanism.  A higher value means a higher
    // preference.  If two mechanisms have the same preference the choice is a
    // local decision. The type is interface{} with range: 1..16.
    Qosifdroppreference interface{}

    // An enumerate type for all the known drop mechanisms. The type is
    // Qosifdropdiscipline.
    Qosifdropdiscipline interface{}
}

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetEntityData() *types.CommonEntityData {
    qosifdroppreferenceentry.EntityData.YFilter = qosifdroppreferenceentry.YFilter
    qosifdroppreferenceentry.EntityData.YangName = "qosIfDropPreferenceEntry"
    qosifdroppreferenceentry.EntityData.BundleName = "cisco_ios_xe"
    qosifdroppreferenceentry.EntityData.ParentYangName = "qosIfDropPreferenceTable"
    qosifdroppreferenceentry.EntityData.SegmentPath = "qosIfDropPreferenceEntry" + "[qosIfDropPreferenceId='" + fmt.Sprintf("%v", qosifdroppreferenceentry.Qosifdroppreferenceid) + "']"
    qosifdroppreferenceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifdroppreferenceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifdroppreferenceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifdroppreferenceentry.EntityData.Children = make(map[string]types.YChild)
    qosifdroppreferenceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosifdroppreferenceentry.EntityData.Leafs["qosIfDropPreferenceId"] = types.YLeaf{"Qosifdroppreferenceid", qosifdroppreferenceentry.Qosifdroppreferenceid}
    qosifdroppreferenceentry.EntityData.Leafs["qosIfDropRoles"] = types.YLeaf{"Qosifdroproles", qosifdroppreferenceentry.Qosifdroproles}
    qosifdroppreferenceentry.EntityData.Leafs["qosIfDropPreference"] = types.YLeaf{"Qosifdroppreference", qosifdroppreferenceentry.Qosifdroppreference}
    qosifdroppreferenceentry.EntityData.Leafs["qosIfDropDiscipline"] = types.YLeaf{"Qosifdropdiscipline", qosifdroppreferenceentry.Qosifdropdiscipline}
    return &(qosifdroppreferenceentry.EntityData)
}

// CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry_Qosifdropdiscipline represents An enumerate type for all the known drop mechanisms.
type CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry_Qosifdropdiscipline string

const (
    CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry_Qosifdropdiscipline_qosIfDropWRED CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry_Qosifdropdiscipline = "qosIfDropWRED"

    CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry_Qosifdropdiscipline_qosIfDropTailDrop CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry_Qosifdropdiscipline = "qosIfDropTailDrop"
)

// CISCOQOSPIBMIB_Qosifdscpassignmenttable
// The assignment of each DSCP to a queue and threshold for each
// interface queue type.
type CISCOQOSPIBMIB_Qosifdscpassignmenttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the queue and threshold set for a
    // packet with a particular DSCP on an interface of a particular type with a
    // particular role combination. The type is slice of
    // CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry.
    Qosifdscpassignmententry []CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry
}

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetEntityData() *types.CommonEntityData {
    qosifdscpassignmenttable.EntityData.YFilter = qosifdscpassignmenttable.YFilter
    qosifdscpassignmenttable.EntityData.YangName = "qosIfDscpAssignmentTable"
    qosifdscpassignmenttable.EntityData.BundleName = "cisco_ios_xe"
    qosifdscpassignmenttable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosifdscpassignmenttable.EntityData.SegmentPath = "qosIfDscpAssignmentTable"
    qosifdscpassignmenttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifdscpassignmenttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifdscpassignmenttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifdscpassignmenttable.EntityData.Children = make(map[string]types.YChild)
    qosifdscpassignmenttable.EntityData.Children["qosIfDscpAssignmentEntry"] = types.YChild{"Qosifdscpassignmententry", nil}
    for i := range qosifdscpassignmenttable.Qosifdscpassignmententry {
        qosifdscpassignmenttable.EntityData.Children[types.GetSegmentPath(&qosifdscpassignmenttable.Qosifdscpassignmententry[i])] = types.YChild{"Qosifdscpassignmententry", &qosifdscpassignmenttable.Qosifdscpassignmententry[i]}
    }
    qosifdscpassignmenttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosifdscpassignmenttable.EntityData)
}

// CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry
// An instance of this class specifies the queue and threshold
// set for a packet with a particular DSCP on an interface of
// a particular type with a particular role combination.
type CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosifdscpassignmentid interface{}

    // The role combination the interface must be configured with. The type is
    // string with length: 0..255.
    Qosifdscproles interface{}

    // The interface queue type to which this row applies. The type is
    // QosInterfaceQueueType.
    Qosifqueuetype interface{}

    // The DSCP to which this row applies. The type is interface{} with range:
    // 0..63.
    Qosifdscp interface{}

    // The queue to which the DSCP applies for the given interface type. The type
    // is interface{} with range: 1..64.
    Qosifqueue interface{}

    // The threshold set of the specified queue to which the DSCP applies for the
    // given interface type. The type is interface{} with range: 1..8.
    Qosifthresholdset interface{}
}

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetEntityData() *types.CommonEntityData {
    qosifdscpassignmententry.EntityData.YFilter = qosifdscpassignmententry.YFilter
    qosifdscpassignmententry.EntityData.YangName = "qosIfDscpAssignmentEntry"
    qosifdscpassignmententry.EntityData.BundleName = "cisco_ios_xe"
    qosifdscpassignmententry.EntityData.ParentYangName = "qosIfDscpAssignmentTable"
    qosifdscpassignmententry.EntityData.SegmentPath = "qosIfDscpAssignmentEntry" + "[qosIfDscpAssignmentId='" + fmt.Sprintf("%v", qosifdscpassignmententry.Qosifdscpassignmentid) + "']"
    qosifdscpassignmententry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifdscpassignmententry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifdscpassignmententry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifdscpassignmententry.EntityData.Children = make(map[string]types.YChild)
    qosifdscpassignmententry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosifdscpassignmententry.EntityData.Leafs["qosIfDscpAssignmentId"] = types.YLeaf{"Qosifdscpassignmentid", qosifdscpassignmententry.Qosifdscpassignmentid}
    qosifdscpassignmententry.EntityData.Leafs["qosIfDscpRoles"] = types.YLeaf{"Qosifdscproles", qosifdscpassignmententry.Qosifdscproles}
    qosifdscpassignmententry.EntityData.Leafs["qosIfQueueType"] = types.YLeaf{"Qosifqueuetype", qosifdscpassignmententry.Qosifqueuetype}
    qosifdscpassignmententry.EntityData.Leafs["qosIfDscp"] = types.YLeaf{"Qosifdscp", qosifdscpassignmententry.Qosifdscp}
    qosifdscpassignmententry.EntityData.Leafs["qosIfQueue"] = types.YLeaf{"Qosifqueue", qosifdscpassignmententry.Qosifqueue}
    qosifdscpassignmententry.EntityData.Leafs["qosIfThresholdSet"] = types.YLeaf{"Qosifthresholdset", qosifdscpassignmententry.Qosifthresholdset}
    return &(qosifdscpassignmententry.EntityData)
}

// CISCOQOSPIBMIB_Qosifredtable
// A class of lower and upper values for each threshold set in a
// queue supporting WRED.  If the size of the queue for a given
// threshold is below the lower value then packets assigned to
// that threshold are always accepted into the queue.  If the
// size of the queue is above upper value then packets are always
// dropped.  If the size of the queue is between the lower and
// the upper then packets are randomly dropped.
type CISCOQOSPIBMIB_Qosifredtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies threshold limits for a particular RED
    // threshold of a given threshold set on an interface and with a particular
    // role combination. The type is slice of
    // CISCOQOSPIBMIB_Qosifredtable_Qosifredentry.
    Qosifredentry []CISCOQOSPIBMIB_Qosifredtable_Qosifredentry
}

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetEntityData() *types.CommonEntityData {
    qosifredtable.EntityData.YFilter = qosifredtable.YFilter
    qosifredtable.EntityData.YangName = "qosIfRedTable"
    qosifredtable.EntityData.BundleName = "cisco_ios_xe"
    qosifredtable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosifredtable.EntityData.SegmentPath = "qosIfRedTable"
    qosifredtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifredtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifredtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifredtable.EntityData.Children = make(map[string]types.YChild)
    qosifredtable.EntityData.Children["qosIfRedEntry"] = types.YChild{"Qosifredentry", nil}
    for i := range qosifredtable.Qosifredentry {
        qosifredtable.EntityData.Children[types.GetSegmentPath(&qosifredtable.Qosifredentry[i])] = types.YChild{"Qosifredentry", &qosifredtable.Qosifredentry[i]}
    }
    qosifredtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosifredtable.EntityData)
}

// CISCOQOSPIBMIB_Qosifredtable_Qosifredentry
// An instance of this class specifies threshold limits for a
// particular RED threshold of a given threshold set on an
// interface and with a particular role combination.
type CISCOQOSPIBMIB_Qosifredtable_Qosifredentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosifredid interface{}

    // The role combination the interface must be configured with. The type is
    // string with length: 0..255.
    Qosifredroles interface{}

    // The values in this entry apply only to queues with the number of thresholds
    // specified by this attribute. The type is ThresholdSetRange.
    Qosifrednumthresholdsets interface{}

    // The threshold set to which the lower and upper values apply. It must be in
    // the range 1 through qosIfRedNumThresholdSets. There must be exactly one PRI
    // for each value in this range. The type is interface{} with range: 1..8.
    Qosifredthresholdset interface{}

    // The threshold value below which no packets are dropped. The type is
    // interface{} with range: 0..100.
    Qosifredthresholdsetlower interface{}

    // The threshold value above which all packets are dropped. The type is
    // interface{} with range: 0..100.
    Qosifredthresholdsetupper interface{}
}

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetEntityData() *types.CommonEntityData {
    qosifredentry.EntityData.YFilter = qosifredentry.YFilter
    qosifredentry.EntityData.YangName = "qosIfRedEntry"
    qosifredentry.EntityData.BundleName = "cisco_ios_xe"
    qosifredentry.EntityData.ParentYangName = "qosIfRedTable"
    qosifredentry.EntityData.SegmentPath = "qosIfRedEntry" + "[qosIfRedId='" + fmt.Sprintf("%v", qosifredentry.Qosifredid) + "']"
    qosifredentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifredentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifredentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifredentry.EntityData.Children = make(map[string]types.YChild)
    qosifredentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosifredentry.EntityData.Leafs["qosIfRedId"] = types.YLeaf{"Qosifredid", qosifredentry.Qosifredid}
    qosifredentry.EntityData.Leafs["qosIfRedRoles"] = types.YLeaf{"Qosifredroles", qosifredentry.Qosifredroles}
    qosifredentry.EntityData.Leafs["qosIfRedNumThresholdSets"] = types.YLeaf{"Qosifrednumthresholdsets", qosifredentry.Qosifrednumthresholdsets}
    qosifredentry.EntityData.Leafs["qosIfRedThresholdSet"] = types.YLeaf{"Qosifredthresholdset", qosifredentry.Qosifredthresholdset}
    qosifredentry.EntityData.Leafs["qosIfRedThresholdSetLower"] = types.YLeaf{"Qosifredthresholdsetlower", qosifredentry.Qosifredthresholdsetlower}
    qosifredentry.EntityData.Leafs["qosIfRedThresholdSetUpper"] = types.YLeaf{"Qosifredthresholdsetupper", qosifredentry.Qosifredthresholdsetupper}
    return &(qosifredentry.EntityData)
}

// CISCOQOSPIBMIB_Qosiftaildroptable
// A class for threshold sets in a queue supporting tail drop.
// If the size of the queue for a given threshold set is at or
// below the specified value then packets assigned to that
// threshold set are always accepted into the queue.  If the size
// of the queue is above the specified value then packets are
// always dropped.
type CISCOQOSPIBMIB_Qosiftaildroptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the queue depth for a particular
    // tail-drop threshold set on an interface with a particular role combination.
    // The type is slice of CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry.
    Qosiftaildropentry []CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry
}

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetEntityData() *types.CommonEntityData {
    qosiftaildroptable.EntityData.YFilter = qosiftaildroptable.YFilter
    qosiftaildroptable.EntityData.YangName = "qosIfTailDropTable"
    qosiftaildroptable.EntityData.BundleName = "cisco_ios_xe"
    qosiftaildroptable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosiftaildroptable.EntityData.SegmentPath = "qosIfTailDropTable"
    qosiftaildroptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosiftaildroptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosiftaildroptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosiftaildroptable.EntityData.Children = make(map[string]types.YChild)
    qosiftaildroptable.EntityData.Children["qosIfTailDropEntry"] = types.YChild{"Qosiftaildropentry", nil}
    for i := range qosiftaildroptable.Qosiftaildropentry {
        qosiftaildroptable.EntityData.Children[types.GetSegmentPath(&qosiftaildroptable.Qosiftaildropentry[i])] = types.YChild{"Qosiftaildropentry", &qosiftaildroptable.Qosiftaildropentry[i]}
    }
    qosiftaildroptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosiftaildroptable.EntityData)
}

// CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry
// An instance of this class specifies the queue depth for a
// particular tail-drop threshold set on an interface with a
// particular role combination.
type CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosiftaildropid interface{}

    // The role combination the interface must be configured with. The type is
    // string with length: 0..255.
    Qosiftaildroproles interface{}

    // The value in this entry applies only to queues with the number of
    // thresholds specified by this attribute. The type is ThresholdSetRange.
    Qosiftaildropnumthresholdsets interface{}

    // The threshold set to which the threshold value applies. The type is
    // interface{} with range: 1..8.
    Qosiftaildropthresholdset interface{}

    // The threshold value above which packets are dropped. The type is
    // interface{} with range: 0..100.
    Qosiftaildropthresholdsetvalue interface{}
}

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetEntityData() *types.CommonEntityData {
    qosiftaildropentry.EntityData.YFilter = qosiftaildropentry.YFilter
    qosiftaildropentry.EntityData.YangName = "qosIfTailDropEntry"
    qosiftaildropentry.EntityData.BundleName = "cisco_ios_xe"
    qosiftaildropentry.EntityData.ParentYangName = "qosIfTailDropTable"
    qosiftaildropentry.EntityData.SegmentPath = "qosIfTailDropEntry" + "[qosIfTailDropId='" + fmt.Sprintf("%v", qosiftaildropentry.Qosiftaildropid) + "']"
    qosiftaildropentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosiftaildropentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosiftaildropentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosiftaildropentry.EntityData.Children = make(map[string]types.YChild)
    qosiftaildropentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosiftaildropentry.EntityData.Leafs["qosIfTailDropId"] = types.YLeaf{"Qosiftaildropid", qosiftaildropentry.Qosiftaildropid}
    qosiftaildropentry.EntityData.Leafs["qosIfTailDropRoles"] = types.YLeaf{"Qosiftaildroproles", qosiftaildropentry.Qosiftaildroproles}
    qosiftaildropentry.EntityData.Leafs["qosIfTailDropNumThresholdSets"] = types.YLeaf{"Qosiftaildropnumthresholdsets", qosiftaildropentry.Qosiftaildropnumthresholdsets}
    qosiftaildropentry.EntityData.Leafs["qosIfTailDropThresholdSet"] = types.YLeaf{"Qosiftaildropthresholdset", qosiftaildropentry.Qosiftaildropthresholdset}
    qosiftaildropentry.EntityData.Leafs["qosIfTailDropThresholdSetValue"] = types.YLeaf{"Qosiftaildropthresholdsetvalue", qosiftaildropentry.Qosiftaildropthresholdsetvalue}
    return &(qosiftaildropentry.EntityData)
}

// CISCOQOSPIBMIB_Qosifweightstable
// A class of scheduling weights for each queue of an interface
// that supports weighted round robin scheduling or a mix of
// priority queueing and weighted round robin.  For a queue with
// N priority queues, the N highest queue numbers are the
// priority queues with the highest queue number having the
// highest priority.  WRR is applied to the non-priority queues.
type CISCOQOSPIBMIB_Qosifweightstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the scheduling weight for a particular
    // queue of an interface with a particular number of queues and with a
    // particular role combination. The type is slice of
    // CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry.
    Qosifweightsentry []CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry
}

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetEntityData() *types.CommonEntityData {
    qosifweightstable.EntityData.YFilter = qosifweightstable.YFilter
    qosifweightstable.EntityData.YangName = "qosIfWeightsTable"
    qosifweightstable.EntityData.BundleName = "cisco_ios_xe"
    qosifweightstable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosifweightstable.EntityData.SegmentPath = "qosIfWeightsTable"
    qosifweightstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifweightstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifweightstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifweightstable.EntityData.Children = make(map[string]types.YChild)
    qosifweightstable.EntityData.Children["qosIfWeightsEntry"] = types.YChild{"Qosifweightsentry", nil}
    for i := range qosifweightstable.Qosifweightsentry {
        qosifweightstable.EntityData.Children[types.GetSegmentPath(&qosifweightstable.Qosifweightsentry[i])] = types.YChild{"Qosifweightsentry", &qosifweightstable.Qosifweightsentry[i]}
    }
    qosifweightstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(qosifweightstable.EntityData)
}

// CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry
// An instance of this class specifies the scheduling weight for
// a particular queue of an interface with a particular number
// of queues and with a particular role combination.
type CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosifweightsid interface{}

    // The role combination the interface must be configured with. The type is
    // string with length: 0..255.
    Qosifweightsroles interface{}

    // The value of the weight in this instance applies only to interfaces with
    // the number of queues specified by this attribute. The type is QueueRange.
    Qosifweightsnumqueues interface{}

    // The queue to which the weight applies. The type is interface{} with range:
    // 1..64.
    Qosifweightsqueue interface{}

    // The maximum number of bytes that may be drained from the queue in one
    // cycle.  The percentage of the bandwith allocated to this queue can be
    // calculated from this attribute and the sum of the drain sizes of all the
    // non-priority queues of the interface. The type is interface{} with range:
    // 0..4294967295.
    Qosifweightsdrainsize interface{}

    // The size of the queue in bytes.  Some devices set queue size in terms of
    // packets.  These devices must calculate the queue size in packets by
    // assuming an average packet size suitable for the particular interface. 
    // Some devices have a fixed size buffer to be shared among all queues.  These
    // devices must allocate a fraction of the total buffer space to this queue
    // calculated as the the ratio of the queue size to the sum of the queue sizes
    // for the interface. The type is interface{} with range: 0..4294967295.
    Qosifweightsqueuesize interface{}
}

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetEntityData() *types.CommonEntityData {
    qosifweightsentry.EntityData.YFilter = qosifweightsentry.YFilter
    qosifweightsentry.EntityData.YangName = "qosIfWeightsEntry"
    qosifweightsentry.EntityData.BundleName = "cisco_ios_xe"
    qosifweightsentry.EntityData.ParentYangName = "qosIfWeightsTable"
    qosifweightsentry.EntityData.SegmentPath = "qosIfWeightsEntry" + "[qosIfWeightsId='" + fmt.Sprintf("%v", qosifweightsentry.Qosifweightsid) + "']"
    qosifweightsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosifweightsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosifweightsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosifweightsentry.EntityData.Children = make(map[string]types.YChild)
    qosifweightsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    qosifweightsentry.EntityData.Leafs["qosIfWeightsId"] = types.YLeaf{"Qosifweightsid", qosifweightsentry.Qosifweightsid}
    qosifweightsentry.EntityData.Leafs["qosIfWeightsRoles"] = types.YLeaf{"Qosifweightsroles", qosifweightsentry.Qosifweightsroles}
    qosifweightsentry.EntityData.Leafs["qosIfWeightsNumQueues"] = types.YLeaf{"Qosifweightsnumqueues", qosifweightsentry.Qosifweightsnumqueues}
    qosifweightsentry.EntityData.Leafs["qosIfWeightsQueue"] = types.YLeaf{"Qosifweightsqueue", qosifweightsentry.Qosifweightsqueue}
    qosifweightsentry.EntityData.Leafs["qosIfWeightsDrainSize"] = types.YLeaf{"Qosifweightsdrainsize", qosifweightsentry.Qosifweightsdrainsize}
    qosifweightsentry.EntityData.Leafs["qosIfWeightsQueueSize"] = types.YLeaf{"Qosifweightsqueuesize", qosifweightsentry.Qosifweightsqueuesize}
    return &(qosifweightsentry.EntityData)
}

