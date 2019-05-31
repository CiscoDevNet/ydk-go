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
    QosDevicePibIncarnationTable CISCOQOSPIBMIB_QosDevicePibIncarnationTable

    // The single instance of this class indicates specific attributes of the
    // device.  These include configuration values such as the configured PDP
    // addresses, the maximum message size, and specific device capabilities.  The
    // latter include input port-based and output port-based classification and/or
    // policing, support for flow based policing, aggregate based policing,
    // traffic shaping capabilities, etc.
    QosDeviceAttributeTable CISCOQOSPIBMIB_QosDeviceAttributeTable

    // This class describes the interface types of the interfaces that exist on
    // the device.  It includes the queue type, role combination and capabilities
    // of interfaces.  The PEP does not report which specific interfaces have
    // which characteristics.
    QosInterfaceTypeTable CISCOQOSPIBMIB_QosInterfaceTypeTable

    // Maps each DSCP to a marked-down DSCP.  Also maps each DSCP to an IP
    // precedence and QosLayer2Cos.  When configured for the first time, all 64
    // entries of the table must be specified. Thereafter, instances may be
    // modified (with a delete and install in a single decision) but not deleted
    // unless all instances are deleted.
    QosDiffServMappingTable CISCOQOSPIBMIB_QosDiffServMappingTable

    // Maps each of eight CoS values to a DSCP.  When configured for the first
    // time, all 8 entries of the table must be specified. Thereafter, instances
    // may be modified (with a delete and install in a single decision) but not
    // deleted unless all instances are deleted.
    QosCosToDscpTable CISCOQOSPIBMIB_QosCosToDscpTable

    // A policy class that specifies what QoS to apply to a packet that does not
    // match any other policy configured for this role combination for a
    // particular direction of traffic.
    QosUnmatchedPolicyTable CISCOQOSPIBMIB_QosUnmatchedPolicyTable

    // A class specifying policing parameters for both microflows and aggregate
    // flows.  This table is designed for policing according to a token bucket
    // scheme where an average rate and burst size is specified.
    QosPolicerTable CISCOQOSPIBMIB_QosPolicerTable

    // Instances of this class identify aggregate flows and the policer to apply
    // to each.
    QosAggregateTable CISCOQOSPIBMIB_QosAggregateTable

    // A class of MAC/Vlan tuples and their associated CoS values.
    QosMacClassificationTable CISCOQOSPIBMIB_QosMacClassificationTable

    // ACE definitions.
    QosIpAceTable CISCOQOSPIBMIB_QosIpAceTable

    // A class that defines a set of ACLs each being an ordered list of ACEs.
    QosIpAclDefinitionTable CISCOQOSPIBMIB_QosIpAclDefinitionTable

    // A class that applies a set of ACLs to interfaces specifying, for each
    // interface the order of the ACL with respect to other ACLs applied to the
    // same interface and, for each ACL the action to take for a packet that
    // matches a permit ACE in that ACL.  Interfaces are specified abstractly in
    // terms of interface role combinations.
    QosIpAclActionTable CISCOQOSPIBMIB_QosIpAclActionTable

    // This class specifies the scheduling preference an interface chooses if it
    // supports multiple scheduling types.  Higher values are preferred over lower
    // values.
    QosIfSchedulingPreferencesTable CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable

    // This class specifies the preference of the drop mechanism an interface
    // chooses if it supports multiple drop mechanisms. Higher values are
    // preferred over lower values.
    QosIfDropPreferenceTable CISCOQOSPIBMIB_QosIfDropPreferenceTable

    // The assignment of each DSCP to a queue and threshold for each interface
    // queue type.
    QosIfDscpAssignmentTable CISCOQOSPIBMIB_QosIfDscpAssignmentTable

    // A class of lower and upper values for each threshold set in a queue
    // supporting WRED.  If the size of the queue for a given threshold is below
    // the lower value then packets assigned to that threshold are always accepted
    // into the queue.  If the size of the queue is above upper value then packets
    // are always dropped.  If the size of the queue is between the lower and the
    // upper then packets are randomly dropped.
    QosIfRedTable CISCOQOSPIBMIB_QosIfRedTable

    // A class for threshold sets in a queue supporting tail drop. If the size of
    // the queue for a given threshold set is at or below the specified value then
    // packets assigned to that threshold set are always accepted into the queue. 
    // If the size of the queue is above the specified value then packets are
    // always dropped.
    QosIfTailDropTable CISCOQOSPIBMIB_QosIfTailDropTable

    // A class of scheduling weights for each queue of an interface that supports
    // weighted round robin scheduling or a mix of priority queueing and weighted
    // round robin.  For a queue with N priority queues, the N highest queue
    // numbers are the priority queues with the highest queue number having the
    // highest priority.  WRR is applied to the non-priority queues.
    QosIfWeightsTable CISCOQOSPIBMIB_QosIfWeightsTable
}

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetEntityData() *types.CommonEntityData {
    cISCOQOSPIBMIB.EntityData.YFilter = cISCOQOSPIBMIB.YFilter
    cISCOQOSPIBMIB.EntityData.YangName = "CISCO-QOS-PIB-MIB"
    cISCOQOSPIBMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOQOSPIBMIB.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    cISCOQOSPIBMIB.EntityData.SegmentPath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB"
    cISCOQOSPIBMIB.EntityData.AbsolutePath = cISCOQOSPIBMIB.EntityData.SegmentPath
    cISCOQOSPIBMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOQOSPIBMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOQOSPIBMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOQOSPIBMIB.EntityData.Children = types.NewOrderedMap()
    cISCOQOSPIBMIB.EntityData.Children.Append("qosDevicePibIncarnationTable", types.YChild{"QosDevicePibIncarnationTable", &cISCOQOSPIBMIB.QosDevicePibIncarnationTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosDeviceAttributeTable", types.YChild{"QosDeviceAttributeTable", &cISCOQOSPIBMIB.QosDeviceAttributeTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosInterfaceTypeTable", types.YChild{"QosInterfaceTypeTable", &cISCOQOSPIBMIB.QosInterfaceTypeTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosDiffServMappingTable", types.YChild{"QosDiffServMappingTable", &cISCOQOSPIBMIB.QosDiffServMappingTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosCosToDscpTable", types.YChild{"QosCosToDscpTable", &cISCOQOSPIBMIB.QosCosToDscpTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosUnmatchedPolicyTable", types.YChild{"QosUnmatchedPolicyTable", &cISCOQOSPIBMIB.QosUnmatchedPolicyTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosPolicerTable", types.YChild{"QosPolicerTable", &cISCOQOSPIBMIB.QosPolicerTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosAggregateTable", types.YChild{"QosAggregateTable", &cISCOQOSPIBMIB.QosAggregateTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosMacClassificationTable", types.YChild{"QosMacClassificationTable", &cISCOQOSPIBMIB.QosMacClassificationTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIpAceTable", types.YChild{"QosIpAceTable", &cISCOQOSPIBMIB.QosIpAceTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIpAclDefinitionTable", types.YChild{"QosIpAclDefinitionTable", &cISCOQOSPIBMIB.QosIpAclDefinitionTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIpAclActionTable", types.YChild{"QosIpAclActionTable", &cISCOQOSPIBMIB.QosIpAclActionTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIfSchedulingPreferencesTable", types.YChild{"QosIfSchedulingPreferencesTable", &cISCOQOSPIBMIB.QosIfSchedulingPreferencesTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIfDropPreferenceTable", types.YChild{"QosIfDropPreferenceTable", &cISCOQOSPIBMIB.QosIfDropPreferenceTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIfDscpAssignmentTable", types.YChild{"QosIfDscpAssignmentTable", &cISCOQOSPIBMIB.QosIfDscpAssignmentTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIfRedTable", types.YChild{"QosIfRedTable", &cISCOQOSPIBMIB.QosIfRedTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIfTailDropTable", types.YChild{"QosIfTailDropTable", &cISCOQOSPIBMIB.QosIfTailDropTable})
    cISCOQOSPIBMIB.EntityData.Children.Append("qosIfWeightsTable", types.YChild{"QosIfWeightsTable", &cISCOQOSPIBMIB.QosIfWeightsTable})
    cISCOQOSPIBMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOQOSPIBMIB.EntityData.YListKeys = []string {}

    return &(cISCOQOSPIBMIB.EntityData)
}

// CISCOQOSPIBMIB_QosDevicePibIncarnationTable
// This class contains a single policy instance that identifies
// the current incarnation of the PIB and the PDP that installed
// this incarnation.  The instance of this class is reported to
// the PDP at client connect time so that the PDP can (attempt
// to) ascertain the current state of the PIB.
type CISCOQOSPIBMIB_QosDevicePibIncarnationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The single policy instance of this class identifies the current incarnation
    // of the PIB and the PDP that installed this incarnation. The type is slice
    // of
    // CISCOQOSPIBMIB_QosDevicePibIncarnationTable_QosDevicePibIncarnationEntry.
    QosDevicePibIncarnationEntry []*CISCOQOSPIBMIB_QosDevicePibIncarnationTable_QosDevicePibIncarnationEntry
}

func (qosDevicePibIncarnationTable *CISCOQOSPIBMIB_QosDevicePibIncarnationTable) GetEntityData() *types.CommonEntityData {
    qosDevicePibIncarnationTable.EntityData.YFilter = qosDevicePibIncarnationTable.YFilter
    qosDevicePibIncarnationTable.EntityData.YangName = "qosDevicePibIncarnationTable"
    qosDevicePibIncarnationTable.EntityData.BundleName = "cisco_ios_xe"
    qosDevicePibIncarnationTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosDevicePibIncarnationTable.EntityData.SegmentPath = "qosDevicePibIncarnationTable"
    qosDevicePibIncarnationTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosDevicePibIncarnationTable.EntityData.SegmentPath
    qosDevicePibIncarnationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosDevicePibIncarnationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosDevicePibIncarnationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosDevicePibIncarnationTable.EntityData.Children = types.NewOrderedMap()
    qosDevicePibIncarnationTable.EntityData.Children.Append("qosDevicePibIncarnationEntry", types.YChild{"QosDevicePibIncarnationEntry", nil})
    for i := range qosDevicePibIncarnationTable.QosDevicePibIncarnationEntry {
        qosDevicePibIncarnationTable.EntityData.Children.Append(types.GetSegmentPath(qosDevicePibIncarnationTable.QosDevicePibIncarnationEntry[i]), types.YChild{"QosDevicePibIncarnationEntry", qosDevicePibIncarnationTable.QosDevicePibIncarnationEntry[i]})
    }
    qosDevicePibIncarnationTable.EntityData.Leafs = types.NewOrderedMap()

    qosDevicePibIncarnationTable.EntityData.YListKeys = []string {}

    return &(qosDevicePibIncarnationTable.EntityData)
}

// CISCOQOSPIBMIB_QosDevicePibIncarnationTable_QosDevicePibIncarnationEntry
// The single policy instance of this class identifies the
// current incarnation of the PIB and the PDP that installed
// this incarnation.
type CISCOQOSPIBMIB_QosDevicePibIncarnationTable_QosDevicePibIncarnationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosDeviceIncarnationId interface{}

    // The name of the PDP that installed the current incarnation of the PIB into
    // the device.  By default it is the zero length string. The type is string.
    QosDevicePdpName interface{}

    // An octet string to identify the current incarnation.  It has meaning to the
    // PDP that installed the PIB and perhaps its standby PDPs. By default the
    // empty string. The type is string with length: 128..128.
    QosDevicePibIncarnation interface{}

    // The number of seconds after a client close or TCP timeout for which the PEP
    // continues to enforce the policy in the PIB. After this interval, the PIB is
    // consired expired and the device no longer enforces the policy installed in
    // the PIB. The type is interface{} with range: 0..4294967295.
    QosDevicePibTtl interface{}
}

func (qosDevicePibIncarnationEntry *CISCOQOSPIBMIB_QosDevicePibIncarnationTable_QosDevicePibIncarnationEntry) GetEntityData() *types.CommonEntityData {
    qosDevicePibIncarnationEntry.EntityData.YFilter = qosDevicePibIncarnationEntry.YFilter
    qosDevicePibIncarnationEntry.EntityData.YangName = "qosDevicePibIncarnationEntry"
    qosDevicePibIncarnationEntry.EntityData.BundleName = "cisco_ios_xe"
    qosDevicePibIncarnationEntry.EntityData.ParentYangName = "qosDevicePibIncarnationTable"
    qosDevicePibIncarnationEntry.EntityData.SegmentPath = "qosDevicePibIncarnationEntry" + types.AddKeyToken(qosDevicePibIncarnationEntry.QosDeviceIncarnationId, "qosDeviceIncarnationId")
    qosDevicePibIncarnationEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosDevicePibIncarnationTable/" + qosDevicePibIncarnationEntry.EntityData.SegmentPath
    qosDevicePibIncarnationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosDevicePibIncarnationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosDevicePibIncarnationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosDevicePibIncarnationEntry.EntityData.Children = types.NewOrderedMap()
    qosDevicePibIncarnationEntry.EntityData.Leafs = types.NewOrderedMap()
    qosDevicePibIncarnationEntry.EntityData.Leafs.Append("qosDeviceIncarnationId", types.YLeaf{"QosDeviceIncarnationId", qosDevicePibIncarnationEntry.QosDeviceIncarnationId})
    qosDevicePibIncarnationEntry.EntityData.Leafs.Append("qosDevicePdpName", types.YLeaf{"QosDevicePdpName", qosDevicePibIncarnationEntry.QosDevicePdpName})
    qosDevicePibIncarnationEntry.EntityData.Leafs.Append("qosDevicePibIncarnation", types.YLeaf{"QosDevicePibIncarnation", qosDevicePibIncarnationEntry.QosDevicePibIncarnation})
    qosDevicePibIncarnationEntry.EntityData.Leafs.Append("qosDevicePibTtl", types.YLeaf{"QosDevicePibTtl", qosDevicePibIncarnationEntry.QosDevicePibTtl})

    qosDevicePibIncarnationEntry.EntityData.YListKeys = []string {"QosDeviceIncarnationId"}

    return &(qosDevicePibIncarnationEntry.EntityData)
}

// CISCOQOSPIBMIB_QosDeviceAttributeTable
// The single instance of this class indicates specific
// attributes of the device.  These include configuration values
// such as the configured PDP addresses, the maximum message
// size, and specific device capabilities.  The latter include
// input port-based and output port-based classification and/or
// policing, support for flow based policing, aggregate based
// policing, traffic shaping capabilities, etc.
type CISCOQOSPIBMIB_QosDeviceAttributeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The single instance of this class indicates specific attributes of the
    // device. The type is slice of
    // CISCOQOSPIBMIB_QosDeviceAttributeTable_QosDeviceAttributeEntry.
    QosDeviceAttributeEntry []*CISCOQOSPIBMIB_QosDeviceAttributeTable_QosDeviceAttributeEntry
}

func (qosDeviceAttributeTable *CISCOQOSPIBMIB_QosDeviceAttributeTable) GetEntityData() *types.CommonEntityData {
    qosDeviceAttributeTable.EntityData.YFilter = qosDeviceAttributeTable.YFilter
    qosDeviceAttributeTable.EntityData.YangName = "qosDeviceAttributeTable"
    qosDeviceAttributeTable.EntityData.BundleName = "cisco_ios_xe"
    qosDeviceAttributeTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosDeviceAttributeTable.EntityData.SegmentPath = "qosDeviceAttributeTable"
    qosDeviceAttributeTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosDeviceAttributeTable.EntityData.SegmentPath
    qosDeviceAttributeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosDeviceAttributeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosDeviceAttributeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosDeviceAttributeTable.EntityData.Children = types.NewOrderedMap()
    qosDeviceAttributeTable.EntityData.Children.Append("qosDeviceAttributeEntry", types.YChild{"QosDeviceAttributeEntry", nil})
    for i := range qosDeviceAttributeTable.QosDeviceAttributeEntry {
        qosDeviceAttributeTable.EntityData.Children.Append(types.GetSegmentPath(qosDeviceAttributeTable.QosDeviceAttributeEntry[i]), types.YChild{"QosDeviceAttributeEntry", qosDeviceAttributeTable.QosDeviceAttributeEntry[i]})
    }
    qosDeviceAttributeTable.EntityData.Leafs = types.NewOrderedMap()

    qosDeviceAttributeTable.EntityData.YListKeys = []string {}

    return &(qosDeviceAttributeTable.EntityData)
}

// CISCOQOSPIBMIB_QosDeviceAttributeTable_QosDeviceAttributeEntry
// The single instance of this class indicates specific
// attributes of the device.
type CISCOQOSPIBMIB_QosDeviceAttributeTable_QosDeviceAttributeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosDeviceAttributeId interface{}

    // The QoS domain that this device belongs to.  This is configured locally on
    // the device (perhaps by some management protocol such as SNMP).  By default,
    // it is the zero-length string. The type is string.
    QosDevicePepDomain interface{}

    // The address of the PDP configured to be the primary PDP for the device. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    QosDevicePrimaryPdp interface{}

    // The address of the PDP configured to be the secondary PDP for the device. 
    // An address of zero indicates no secondary is configured. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    QosDeviceSecondaryPdp interface{}

    // The maximum size message that this PEP is capable of receiving in bytes.  A
    // value of zero means that the maximum message size is unspecified (but does
    // not mean it is unlimited).  A message greater than this maximum results in
    // a MessageTooBig error on a 'no commit' REP. The type is interface{} with
    // range: 0..4294967295.
    QosDeviceMaxMessageSize interface{}

    // An enumeration of device capabilities.  Used by the PDP to select policies
    // and configuration to push to the PEP. The type is map[string]bool.
    QosDeviceCapabilities interface{}
}

func (qosDeviceAttributeEntry *CISCOQOSPIBMIB_QosDeviceAttributeTable_QosDeviceAttributeEntry) GetEntityData() *types.CommonEntityData {
    qosDeviceAttributeEntry.EntityData.YFilter = qosDeviceAttributeEntry.YFilter
    qosDeviceAttributeEntry.EntityData.YangName = "qosDeviceAttributeEntry"
    qosDeviceAttributeEntry.EntityData.BundleName = "cisco_ios_xe"
    qosDeviceAttributeEntry.EntityData.ParentYangName = "qosDeviceAttributeTable"
    qosDeviceAttributeEntry.EntityData.SegmentPath = "qosDeviceAttributeEntry" + types.AddKeyToken(qosDeviceAttributeEntry.QosDeviceAttributeId, "qosDeviceAttributeId")
    qosDeviceAttributeEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosDeviceAttributeTable/" + qosDeviceAttributeEntry.EntityData.SegmentPath
    qosDeviceAttributeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosDeviceAttributeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosDeviceAttributeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosDeviceAttributeEntry.EntityData.Children = types.NewOrderedMap()
    qosDeviceAttributeEntry.EntityData.Leafs = types.NewOrderedMap()
    qosDeviceAttributeEntry.EntityData.Leafs.Append("qosDeviceAttributeId", types.YLeaf{"QosDeviceAttributeId", qosDeviceAttributeEntry.QosDeviceAttributeId})
    qosDeviceAttributeEntry.EntityData.Leafs.Append("qosDevicePepDomain", types.YLeaf{"QosDevicePepDomain", qosDeviceAttributeEntry.QosDevicePepDomain})
    qosDeviceAttributeEntry.EntityData.Leafs.Append("qosDevicePrimaryPdp", types.YLeaf{"QosDevicePrimaryPdp", qosDeviceAttributeEntry.QosDevicePrimaryPdp})
    qosDeviceAttributeEntry.EntityData.Leafs.Append("qosDeviceSecondaryPdp", types.YLeaf{"QosDeviceSecondaryPdp", qosDeviceAttributeEntry.QosDeviceSecondaryPdp})
    qosDeviceAttributeEntry.EntityData.Leafs.Append("qosDeviceMaxMessageSize", types.YLeaf{"QosDeviceMaxMessageSize", qosDeviceAttributeEntry.QosDeviceMaxMessageSize})
    qosDeviceAttributeEntry.EntityData.Leafs.Append("qosDeviceCapabilities", types.YLeaf{"QosDeviceCapabilities", qosDeviceAttributeEntry.QosDeviceCapabilities})

    qosDeviceAttributeEntry.EntityData.YListKeys = []string {"QosDeviceAttributeId"}

    return &(qosDeviceAttributeEntry.EntityData)
}

// CISCOQOSPIBMIB_QosInterfaceTypeTable
// This class describes the interface types of the interfaces
// that exist on the device.  It includes the queue type, role
// combination and capabilities of interfaces.  The PEP does not
// report which specific interfaces have which characteristics.
type CISCOQOSPIBMIB_QosInterfaceTypeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class describes a role combination for an interface
    // type of an interface that exists on the device. The type is slice of
    // CISCOQOSPIBMIB_QosInterfaceTypeTable_QosInterfaceTypeEntry.
    QosInterfaceTypeEntry []*CISCOQOSPIBMIB_QosInterfaceTypeTable_QosInterfaceTypeEntry
}

func (qosInterfaceTypeTable *CISCOQOSPIBMIB_QosInterfaceTypeTable) GetEntityData() *types.CommonEntityData {
    qosInterfaceTypeTable.EntityData.YFilter = qosInterfaceTypeTable.YFilter
    qosInterfaceTypeTable.EntityData.YangName = "qosInterfaceTypeTable"
    qosInterfaceTypeTable.EntityData.BundleName = "cisco_ios_xe"
    qosInterfaceTypeTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosInterfaceTypeTable.EntityData.SegmentPath = "qosInterfaceTypeTable"
    qosInterfaceTypeTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosInterfaceTypeTable.EntityData.SegmentPath
    qosInterfaceTypeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosInterfaceTypeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosInterfaceTypeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosInterfaceTypeTable.EntityData.Children = types.NewOrderedMap()
    qosInterfaceTypeTable.EntityData.Children.Append("qosInterfaceTypeEntry", types.YChild{"QosInterfaceTypeEntry", nil})
    for i := range qosInterfaceTypeTable.QosInterfaceTypeEntry {
        qosInterfaceTypeTable.EntityData.Children.Append(types.GetSegmentPath(qosInterfaceTypeTable.QosInterfaceTypeEntry[i]), types.YChild{"QosInterfaceTypeEntry", qosInterfaceTypeTable.QosInterfaceTypeEntry[i]})
    }
    qosInterfaceTypeTable.EntityData.Leafs = types.NewOrderedMap()

    qosInterfaceTypeTable.EntityData.YListKeys = []string {}

    return &(qosInterfaceTypeTable.EntityData)
}

// CISCOQOSPIBMIB_QosInterfaceTypeTable_QosInterfaceTypeEntry
// An instance of this class describes a role combination for
// an interface type of an interface that exists on the device.
type CISCOQOSPIBMIB_QosInterfaceTypeTable_QosInterfaceTypeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosInterfaceTypeId interface{}

    // The interface type in terms of number of queues and thresholds. The type is
    // QosInterfaceQueueType.
    QosInterfaceQueueType interface{}

    // A combination of roles on at least one interface of type qosInterfaceType.
    // The type is string with length: 0..255.
    QosInterfaceTypeRoles interface{}

    // An enumeration of interface capabilities.  Used by the PDP to select
    // policies and configuration to push to the PEP. The type is map[string]bool.
    QosInterfaceTypeCapabilities interface{}
}

func (qosInterfaceTypeEntry *CISCOQOSPIBMIB_QosInterfaceTypeTable_QosInterfaceTypeEntry) GetEntityData() *types.CommonEntityData {
    qosInterfaceTypeEntry.EntityData.YFilter = qosInterfaceTypeEntry.YFilter
    qosInterfaceTypeEntry.EntityData.YangName = "qosInterfaceTypeEntry"
    qosInterfaceTypeEntry.EntityData.BundleName = "cisco_ios_xe"
    qosInterfaceTypeEntry.EntityData.ParentYangName = "qosInterfaceTypeTable"
    qosInterfaceTypeEntry.EntityData.SegmentPath = "qosInterfaceTypeEntry" + types.AddKeyToken(qosInterfaceTypeEntry.QosInterfaceTypeId, "qosInterfaceTypeId")
    qosInterfaceTypeEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosInterfaceTypeTable/" + qosInterfaceTypeEntry.EntityData.SegmentPath
    qosInterfaceTypeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosInterfaceTypeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosInterfaceTypeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosInterfaceTypeEntry.EntityData.Children = types.NewOrderedMap()
    qosInterfaceTypeEntry.EntityData.Leafs = types.NewOrderedMap()
    qosInterfaceTypeEntry.EntityData.Leafs.Append("qosInterfaceTypeId", types.YLeaf{"QosInterfaceTypeId", qosInterfaceTypeEntry.QosInterfaceTypeId})
    qosInterfaceTypeEntry.EntityData.Leafs.Append("qosInterfaceQueueType", types.YLeaf{"QosInterfaceQueueType", qosInterfaceTypeEntry.QosInterfaceQueueType})
    qosInterfaceTypeEntry.EntityData.Leafs.Append("qosInterfaceTypeRoles", types.YLeaf{"QosInterfaceTypeRoles", qosInterfaceTypeEntry.QosInterfaceTypeRoles})
    qosInterfaceTypeEntry.EntityData.Leafs.Append("qosInterfaceTypeCapabilities", types.YLeaf{"QosInterfaceTypeCapabilities", qosInterfaceTypeEntry.QosInterfaceTypeCapabilities})

    qosInterfaceTypeEntry.EntityData.YListKeys = []string {"QosInterfaceTypeId"}

    return &(qosInterfaceTypeEntry.EntityData)
}

// CISCOQOSPIBMIB_QosDiffServMappingTable
// Maps each DSCP to a marked-down DSCP.  Also maps each DSCP to
// an IP precedence and QosLayer2Cos.  When configured for the
// first time, all 64 entries of the table must be
// specified. Thereafter, instances may be modified (with a
// delete and install in a single decision) but not deleted
// unless all instances are deleted.
type CISCOQOSPIBMIB_QosDiffServMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class represents mappings from a DSCP. The type is
    // slice of CISCOQOSPIBMIB_QosDiffServMappingTable_QosDiffServMappingEntry.
    QosDiffServMappingEntry []*CISCOQOSPIBMIB_QosDiffServMappingTable_QosDiffServMappingEntry
}

func (qosDiffServMappingTable *CISCOQOSPIBMIB_QosDiffServMappingTable) GetEntityData() *types.CommonEntityData {
    qosDiffServMappingTable.EntityData.YFilter = qosDiffServMappingTable.YFilter
    qosDiffServMappingTable.EntityData.YangName = "qosDiffServMappingTable"
    qosDiffServMappingTable.EntityData.BundleName = "cisco_ios_xe"
    qosDiffServMappingTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosDiffServMappingTable.EntityData.SegmentPath = "qosDiffServMappingTable"
    qosDiffServMappingTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosDiffServMappingTable.EntityData.SegmentPath
    qosDiffServMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosDiffServMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosDiffServMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosDiffServMappingTable.EntityData.Children = types.NewOrderedMap()
    qosDiffServMappingTable.EntityData.Children.Append("qosDiffServMappingEntry", types.YChild{"QosDiffServMappingEntry", nil})
    for i := range qosDiffServMappingTable.QosDiffServMappingEntry {
        qosDiffServMappingTable.EntityData.Children.Append(types.GetSegmentPath(qosDiffServMappingTable.QosDiffServMappingEntry[i]), types.YChild{"QosDiffServMappingEntry", qosDiffServMappingTable.QosDiffServMappingEntry[i]})
    }
    qosDiffServMappingTable.EntityData.Leafs = types.NewOrderedMap()

    qosDiffServMappingTable.EntityData.YListKeys = []string {}

    return &(qosDiffServMappingTable.EntityData)
}

// CISCOQOSPIBMIB_QosDiffServMappingTable_QosDiffServMappingEntry
// An instance of this class represents mappings from a DSCP.
type CISCOQOSPIBMIB_QosDiffServMappingTable_QosDiffServMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A DSCP for which this entry contains mappings. The
    // type is interface{} with range: 0..63.
    QosDscp interface{}

    // The DSCP to use instead of the qosDscp when the packet is out of profile
    // and hence marked as such. The type is interface{} with range: 0..63.
    QosMarkedDscp interface{}

    // The L2 CoS value to use when mapping this DSCP to layer 2 CoS. The type is
    // interface{} with range: 0..7.
    QosL2Cos interface{}
}

func (qosDiffServMappingEntry *CISCOQOSPIBMIB_QosDiffServMappingTable_QosDiffServMappingEntry) GetEntityData() *types.CommonEntityData {
    qosDiffServMappingEntry.EntityData.YFilter = qosDiffServMappingEntry.YFilter
    qosDiffServMappingEntry.EntityData.YangName = "qosDiffServMappingEntry"
    qosDiffServMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    qosDiffServMappingEntry.EntityData.ParentYangName = "qosDiffServMappingTable"
    qosDiffServMappingEntry.EntityData.SegmentPath = "qosDiffServMappingEntry" + types.AddKeyToken(qosDiffServMappingEntry.QosDscp, "qosDscp")
    qosDiffServMappingEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosDiffServMappingTable/" + qosDiffServMappingEntry.EntityData.SegmentPath
    qosDiffServMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosDiffServMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosDiffServMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosDiffServMappingEntry.EntityData.Children = types.NewOrderedMap()
    qosDiffServMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    qosDiffServMappingEntry.EntityData.Leafs.Append("qosDscp", types.YLeaf{"QosDscp", qosDiffServMappingEntry.QosDscp})
    qosDiffServMappingEntry.EntityData.Leafs.Append("qosMarkedDscp", types.YLeaf{"QosMarkedDscp", qosDiffServMappingEntry.QosMarkedDscp})
    qosDiffServMappingEntry.EntityData.Leafs.Append("qosL2Cos", types.YLeaf{"QosL2Cos", qosDiffServMappingEntry.QosL2Cos})

    qosDiffServMappingEntry.EntityData.YListKeys = []string {"QosDscp"}

    return &(qosDiffServMappingEntry.EntityData)
}

// CISCOQOSPIBMIB_QosCosToDscpTable
// Maps each of eight CoS values to a DSCP.  When configured for
// the first time, all 8 entries of the table must be
// specified. Thereafter, instances may be modified (with a
// delete and install in a single decision) but not deleted
// unless all instances are deleted.
type CISCOQOSPIBMIB_QosCosToDscpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class maps a CoS value to a DSCP. The type is slice of
    // CISCOQOSPIBMIB_QosCosToDscpTable_QosCosToDscpEntry.
    QosCosToDscpEntry []*CISCOQOSPIBMIB_QosCosToDscpTable_QosCosToDscpEntry
}

func (qosCosToDscpTable *CISCOQOSPIBMIB_QosCosToDscpTable) GetEntityData() *types.CommonEntityData {
    qosCosToDscpTable.EntityData.YFilter = qosCosToDscpTable.YFilter
    qosCosToDscpTable.EntityData.YangName = "qosCosToDscpTable"
    qosCosToDscpTable.EntityData.BundleName = "cisco_ios_xe"
    qosCosToDscpTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosCosToDscpTable.EntityData.SegmentPath = "qosCosToDscpTable"
    qosCosToDscpTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosCosToDscpTable.EntityData.SegmentPath
    qosCosToDscpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosCosToDscpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosCosToDscpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosCosToDscpTable.EntityData.Children = types.NewOrderedMap()
    qosCosToDscpTable.EntityData.Children.Append("qosCosToDscpEntry", types.YChild{"QosCosToDscpEntry", nil})
    for i := range qosCosToDscpTable.QosCosToDscpEntry {
        qosCosToDscpTable.EntityData.Children.Append(types.GetSegmentPath(qosCosToDscpTable.QosCosToDscpEntry[i]), types.YChild{"QosCosToDscpEntry", qosCosToDscpTable.QosCosToDscpEntry[i]})
    }
    qosCosToDscpTable.EntityData.Leafs = types.NewOrderedMap()

    qosCosToDscpTable.EntityData.YListKeys = []string {}

    return &(qosCosToDscpTable.EntityData)
}

// CISCOQOSPIBMIB_QosCosToDscpTable_QosCosToDscpEntry
// An instance of this class maps a CoS value to a DSCP.
type CISCOQOSPIBMIB_QosCosToDscpTable_QosCosToDscpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The L2 CoS value that is being mapped. The type is
    // interface{} with range: 0..7.
    QosCosToDscpCos interface{}

    // The DSCP value to use when mapping the L2 CoS to a DSCP. The type is
    // interface{} with range: 0..63.
    QosCosToDscpDscp interface{}
}

func (qosCosToDscpEntry *CISCOQOSPIBMIB_QosCosToDscpTable_QosCosToDscpEntry) GetEntityData() *types.CommonEntityData {
    qosCosToDscpEntry.EntityData.YFilter = qosCosToDscpEntry.YFilter
    qosCosToDscpEntry.EntityData.YangName = "qosCosToDscpEntry"
    qosCosToDscpEntry.EntityData.BundleName = "cisco_ios_xe"
    qosCosToDscpEntry.EntityData.ParentYangName = "qosCosToDscpTable"
    qosCosToDscpEntry.EntityData.SegmentPath = "qosCosToDscpEntry" + types.AddKeyToken(qosCosToDscpEntry.QosCosToDscpCos, "qosCosToDscpCos")
    qosCosToDscpEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosCosToDscpTable/" + qosCosToDscpEntry.EntityData.SegmentPath
    qosCosToDscpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosCosToDscpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosCosToDscpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosCosToDscpEntry.EntityData.Children = types.NewOrderedMap()
    qosCosToDscpEntry.EntityData.Leafs = types.NewOrderedMap()
    qosCosToDscpEntry.EntityData.Leafs.Append("qosCosToDscpCos", types.YLeaf{"QosCosToDscpCos", qosCosToDscpEntry.QosCosToDscpCos})
    qosCosToDscpEntry.EntityData.Leafs.Append("qosCosToDscpDscp", types.YLeaf{"QosCosToDscpDscp", qosCosToDscpEntry.QosCosToDscpDscp})

    qosCosToDscpEntry.EntityData.YListKeys = []string {"QosCosToDscpCos"}

    return &(qosCosToDscpEntry.EntityData)
}

// CISCOQOSPIBMIB_QosUnmatchedPolicyTable
// A policy class that specifies what QoS to apply to a packet
// that does not match any other policy configured for this role
// combination for a particular direction of traffic.
type CISCOQOSPIBMIB_QosUnmatchedPolicyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the unmatched policy for a particular
    // role combination for incoming or outgoing traffic. The type is slice of
    // CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry.
    QosUnmatchedPolicyEntry []*CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry
}

func (qosUnmatchedPolicyTable *CISCOQOSPIBMIB_QosUnmatchedPolicyTable) GetEntityData() *types.CommonEntityData {
    qosUnmatchedPolicyTable.EntityData.YFilter = qosUnmatchedPolicyTable.YFilter
    qosUnmatchedPolicyTable.EntityData.YangName = "qosUnmatchedPolicyTable"
    qosUnmatchedPolicyTable.EntityData.BundleName = "cisco_ios_xe"
    qosUnmatchedPolicyTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosUnmatchedPolicyTable.EntityData.SegmentPath = "qosUnmatchedPolicyTable"
    qosUnmatchedPolicyTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosUnmatchedPolicyTable.EntityData.SegmentPath
    qosUnmatchedPolicyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosUnmatchedPolicyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosUnmatchedPolicyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosUnmatchedPolicyTable.EntityData.Children = types.NewOrderedMap()
    qosUnmatchedPolicyTable.EntityData.Children.Append("qosUnmatchedPolicyEntry", types.YChild{"QosUnmatchedPolicyEntry", nil})
    for i := range qosUnmatchedPolicyTable.QosUnmatchedPolicyEntry {
        qosUnmatchedPolicyTable.EntityData.Children.Append(types.GetSegmentPath(qosUnmatchedPolicyTable.QosUnmatchedPolicyEntry[i]), types.YChild{"QosUnmatchedPolicyEntry", qosUnmatchedPolicyTable.QosUnmatchedPolicyEntry[i]})
    }
    qosUnmatchedPolicyTable.EntityData.Leafs = types.NewOrderedMap()

    qosUnmatchedPolicyTable.EntityData.YListKeys = []string {}

    return &(qosUnmatchedPolicyTable.EntityData)
}

// CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry
// An instance of this class specifies the unmatched policy
// for a particular role combination for incoming or outgoing
// traffic.
type CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosUnmatchedPolicyId interface{}

    // Role combination for which this instance applies. The type is string with
    // length: 0..255.
    QosUnmatchedPolicyRole interface{}

    // The direction of packet flow at the interface in question to which this
    // instance applies. The type is QosUnmatchedPolicyDirection.
    QosUnmatchedPolicyDirection interface{}

    // The DSCP to classify the unmatched packet with.  This must be specified
    // even if qosUnmatchedPolicyDscpTrusted is true. The type is interface{} with
    // range: 0..63.
    QosUnmatchedPolicyDscp interface{}

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
    QosUnmatchedPolicyDscpTrusted interface{}

    // An index identifying the instance of policer to apply to unmatched packets.
    // It must correspond to the integer index of an instance of class
    // qosPolicerTable or be zero.  If zero, the microflow is not policed. The
    // type is interface{} with range: 0..4294967295.
    QosUnmatchPolMicroFlowPolicerId interface{}

    // An index identifying the aggregate that the packet belongs to.  It must
    // correspond to the integer index of an instance of class qosAggregateTable
    // or be zero.  If zero, the microflow does not belong to any aggregate and is
    // not policed as part of any aggregate. The type is interface{} with range:
    // 0..4294967295.
    QosUnmatchedPolicyAggregateId interface{}
}

func (qosUnmatchedPolicyEntry *CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry) GetEntityData() *types.CommonEntityData {
    qosUnmatchedPolicyEntry.EntityData.YFilter = qosUnmatchedPolicyEntry.YFilter
    qosUnmatchedPolicyEntry.EntityData.YangName = "qosUnmatchedPolicyEntry"
    qosUnmatchedPolicyEntry.EntityData.BundleName = "cisco_ios_xe"
    qosUnmatchedPolicyEntry.EntityData.ParentYangName = "qosUnmatchedPolicyTable"
    qosUnmatchedPolicyEntry.EntityData.SegmentPath = "qosUnmatchedPolicyEntry" + types.AddKeyToken(qosUnmatchedPolicyEntry.QosUnmatchedPolicyId, "qosUnmatchedPolicyId")
    qosUnmatchedPolicyEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosUnmatchedPolicyTable/" + qosUnmatchedPolicyEntry.EntityData.SegmentPath
    qosUnmatchedPolicyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosUnmatchedPolicyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosUnmatchedPolicyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosUnmatchedPolicyEntry.EntityData.Children = types.NewOrderedMap()
    qosUnmatchedPolicyEntry.EntityData.Leafs = types.NewOrderedMap()
    qosUnmatchedPolicyEntry.EntityData.Leafs.Append("qosUnmatchedPolicyId", types.YLeaf{"QosUnmatchedPolicyId", qosUnmatchedPolicyEntry.QosUnmatchedPolicyId})
    qosUnmatchedPolicyEntry.EntityData.Leafs.Append("qosUnmatchedPolicyRole", types.YLeaf{"QosUnmatchedPolicyRole", qosUnmatchedPolicyEntry.QosUnmatchedPolicyRole})
    qosUnmatchedPolicyEntry.EntityData.Leafs.Append("qosUnmatchedPolicyDirection", types.YLeaf{"QosUnmatchedPolicyDirection", qosUnmatchedPolicyEntry.QosUnmatchedPolicyDirection})
    qosUnmatchedPolicyEntry.EntityData.Leafs.Append("qosUnmatchedPolicyDscp", types.YLeaf{"QosUnmatchedPolicyDscp", qosUnmatchedPolicyEntry.QosUnmatchedPolicyDscp})
    qosUnmatchedPolicyEntry.EntityData.Leafs.Append("qosUnmatchedPolicyDscpTrusted", types.YLeaf{"QosUnmatchedPolicyDscpTrusted", qosUnmatchedPolicyEntry.QosUnmatchedPolicyDscpTrusted})
    qosUnmatchedPolicyEntry.EntityData.Leafs.Append("qosUnmatchPolMicroFlowPolicerId", types.YLeaf{"QosUnmatchPolMicroFlowPolicerId", qosUnmatchedPolicyEntry.QosUnmatchPolMicroFlowPolicerId})
    qosUnmatchedPolicyEntry.EntityData.Leafs.Append("qosUnmatchedPolicyAggregateId", types.YLeaf{"QosUnmatchedPolicyAggregateId", qosUnmatchedPolicyEntry.QosUnmatchedPolicyAggregateId})

    qosUnmatchedPolicyEntry.EntityData.YListKeys = []string {"QosUnmatchedPolicyId"}

    return &(qosUnmatchedPolicyEntry.EntityData)
}

// CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry_QosUnmatchedPolicyDirection represents which this instance applies.
type CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry_QosUnmatchedPolicyDirection string

const (
    CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry_QosUnmatchedPolicyDirection_in CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry_QosUnmatchedPolicyDirection = "in"

    CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry_QosUnmatchedPolicyDirection_out CISCOQOSPIBMIB_QosUnmatchedPolicyTable_QosUnmatchedPolicyEntry_QosUnmatchedPolicyDirection = "out"
)

// CISCOQOSPIBMIB_QosPolicerTable
// A class specifying policing parameters for both microflows
// and aggregate flows.  This table is designed for policing
// according to a token bucket scheme where an average rate and
// burst size is specified.
type CISCOQOSPIBMIB_QosPolicerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies a set of policing parameters. The type
    // is slice of CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry.
    QosPolicerEntry []*CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry
}

func (qosPolicerTable *CISCOQOSPIBMIB_QosPolicerTable) GetEntityData() *types.CommonEntityData {
    qosPolicerTable.EntityData.YFilter = qosPolicerTable.YFilter
    qosPolicerTable.EntityData.YangName = "qosPolicerTable"
    qosPolicerTable.EntityData.BundleName = "cisco_ios_xe"
    qosPolicerTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosPolicerTable.EntityData.SegmentPath = "qosPolicerTable"
    qosPolicerTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosPolicerTable.EntityData.SegmentPath
    qosPolicerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosPolicerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosPolicerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosPolicerTable.EntityData.Children = types.NewOrderedMap()
    qosPolicerTable.EntityData.Children.Append("qosPolicerEntry", types.YChild{"QosPolicerEntry", nil})
    for i := range qosPolicerTable.QosPolicerEntry {
        qosPolicerTable.EntityData.Children.Append(types.GetSegmentPath(qosPolicerTable.QosPolicerEntry[i]), types.YChild{"QosPolicerEntry", qosPolicerTable.QosPolicerEntry[i]})
    }
    qosPolicerTable.EntityData.Leafs = types.NewOrderedMap()

    qosPolicerTable.EntityData.YListKeys = []string {}

    return &(qosPolicerTable.EntityData)
}

// CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry
// An instance of this class specifies a set of policing
// parameters.
type CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosPolicerId interface{}

    // The token rate.  It is specified in units of bit/s. A rate of zero means
    // that all packets will be out of profile.  If the qosPolicerAction is set to
    // drop then this effectively denies any service to packets policed by this
    // policer. The type is interface{} with range: 0..4294967295.
    QosPolicerRate interface{}

    // The normal size of a burst in terms of bits. The type is interface{} with
    // range: 0..4294967295.
    QosPolicerNormalBurst interface{}

    // The excess size of a burst in terms of bits. The type is interface{} with
    // range: 0..4294967295.
    QosPolicerExcessBurst interface{}

    // An indication of how to handle out of profile packets.  When the shape
    // action is chosen then traffic is shaped to the rate specified by
    // qosPolicerRate. The type is QosPolicerAction.
    QosPolicerAction interface{}
}

func (qosPolicerEntry *CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry) GetEntityData() *types.CommonEntityData {
    qosPolicerEntry.EntityData.YFilter = qosPolicerEntry.YFilter
    qosPolicerEntry.EntityData.YangName = "qosPolicerEntry"
    qosPolicerEntry.EntityData.BundleName = "cisco_ios_xe"
    qosPolicerEntry.EntityData.ParentYangName = "qosPolicerTable"
    qosPolicerEntry.EntityData.SegmentPath = "qosPolicerEntry" + types.AddKeyToken(qosPolicerEntry.QosPolicerId, "qosPolicerId")
    qosPolicerEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosPolicerTable/" + qosPolicerEntry.EntityData.SegmentPath
    qosPolicerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosPolicerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosPolicerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosPolicerEntry.EntityData.Children = types.NewOrderedMap()
    qosPolicerEntry.EntityData.Leafs = types.NewOrderedMap()
    qosPolicerEntry.EntityData.Leafs.Append("qosPolicerId", types.YLeaf{"QosPolicerId", qosPolicerEntry.QosPolicerId})
    qosPolicerEntry.EntityData.Leafs.Append("qosPolicerRate", types.YLeaf{"QosPolicerRate", qosPolicerEntry.QosPolicerRate})
    qosPolicerEntry.EntityData.Leafs.Append("qosPolicerNormalBurst", types.YLeaf{"QosPolicerNormalBurst", qosPolicerEntry.QosPolicerNormalBurst})
    qosPolicerEntry.EntityData.Leafs.Append("qosPolicerExcessBurst", types.YLeaf{"QosPolicerExcessBurst", qosPolicerEntry.QosPolicerExcessBurst})
    qosPolicerEntry.EntityData.Leafs.Append("qosPolicerAction", types.YLeaf{"QosPolicerAction", qosPolicerEntry.QosPolicerAction})

    qosPolicerEntry.EntityData.YListKeys = []string {"QosPolicerId"}

    return &(qosPolicerEntry.EntityData)
}

// CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry_QosPolicerAction represents specified by qosPolicerRate.
type CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry_QosPolicerAction string

const (
    CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry_QosPolicerAction_drop CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry_QosPolicerAction = "drop"

    CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry_QosPolicerAction_mark CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry_QosPolicerAction = "mark"

    CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry_QosPolicerAction_shape CISCOQOSPIBMIB_QosPolicerTable_QosPolicerEntry_QosPolicerAction = "shape"
)

// CISCOQOSPIBMIB_QosAggregateTable
// Instances of this class identify aggregate flows and the
// policer to apply to each.
type CISCOQOSPIBMIB_QosAggregateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the policer to apply to an aggregate
    // flow. The type is slice of
    // CISCOQOSPIBMIB_QosAggregateTable_QosAggregateEntry.
    QosAggregateEntry []*CISCOQOSPIBMIB_QosAggregateTable_QosAggregateEntry
}

func (qosAggregateTable *CISCOQOSPIBMIB_QosAggregateTable) GetEntityData() *types.CommonEntityData {
    qosAggregateTable.EntityData.YFilter = qosAggregateTable.YFilter
    qosAggregateTable.EntityData.YangName = "qosAggregateTable"
    qosAggregateTable.EntityData.BundleName = "cisco_ios_xe"
    qosAggregateTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosAggregateTable.EntityData.SegmentPath = "qosAggregateTable"
    qosAggregateTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosAggregateTable.EntityData.SegmentPath
    qosAggregateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosAggregateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosAggregateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosAggregateTable.EntityData.Children = types.NewOrderedMap()
    qosAggregateTable.EntityData.Children.Append("qosAggregateEntry", types.YChild{"QosAggregateEntry", nil})
    for i := range qosAggregateTable.QosAggregateEntry {
        qosAggregateTable.EntityData.Children.Append(types.GetSegmentPath(qosAggregateTable.QosAggregateEntry[i]), types.YChild{"QosAggregateEntry", qosAggregateTable.QosAggregateEntry[i]})
    }
    qosAggregateTable.EntityData.Leafs = types.NewOrderedMap()

    qosAggregateTable.EntityData.YListKeys = []string {}

    return &(qosAggregateTable.EntityData)
}

// CISCOQOSPIBMIB_QosAggregateTable_QosAggregateEntry
// An instance of this class specifies the policer to apply to
// an aggregate flow.
type CISCOQOSPIBMIB_QosAggregateTable_QosAggregateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosAggregateId interface{}

    // An index identifying the instance of policer to apply to the aggregate.  It
    // must correspond to the integer index of an instance of class
    // qosPolicerTable. The type is interface{} with range: 0..4294967295.
    QosAggregatePolicerId interface{}
}

func (qosAggregateEntry *CISCOQOSPIBMIB_QosAggregateTable_QosAggregateEntry) GetEntityData() *types.CommonEntityData {
    qosAggregateEntry.EntityData.YFilter = qosAggregateEntry.YFilter
    qosAggregateEntry.EntityData.YangName = "qosAggregateEntry"
    qosAggregateEntry.EntityData.BundleName = "cisco_ios_xe"
    qosAggregateEntry.EntityData.ParentYangName = "qosAggregateTable"
    qosAggregateEntry.EntityData.SegmentPath = "qosAggregateEntry" + types.AddKeyToken(qosAggregateEntry.QosAggregateId, "qosAggregateId")
    qosAggregateEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosAggregateTable/" + qosAggregateEntry.EntityData.SegmentPath
    qosAggregateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosAggregateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosAggregateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosAggregateEntry.EntityData.Children = types.NewOrderedMap()
    qosAggregateEntry.EntityData.Leafs = types.NewOrderedMap()
    qosAggregateEntry.EntityData.Leafs.Append("qosAggregateId", types.YLeaf{"QosAggregateId", qosAggregateEntry.QosAggregateId})
    qosAggregateEntry.EntityData.Leafs.Append("qosAggregatePolicerId", types.YLeaf{"QosAggregatePolicerId", qosAggregateEntry.QosAggregatePolicerId})

    qosAggregateEntry.EntityData.YListKeys = []string {"QosAggregateId"}

    return &(qosAggregateEntry.EntityData)
}

// CISCOQOSPIBMIB_QosMacClassificationTable
// A class of MAC/Vlan tuples and their associated CoS values.
type CISCOQOSPIBMIB_QosMacClassificationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the mapping of a VLAN and a MAC address
    // to a CoS value. The type is slice of
    // CISCOQOSPIBMIB_QosMacClassificationTable_QosMacClassificationEntry.
    QosMacClassificationEntry []*CISCOQOSPIBMIB_QosMacClassificationTable_QosMacClassificationEntry
}

func (qosMacClassificationTable *CISCOQOSPIBMIB_QosMacClassificationTable) GetEntityData() *types.CommonEntityData {
    qosMacClassificationTable.EntityData.YFilter = qosMacClassificationTable.YFilter
    qosMacClassificationTable.EntityData.YangName = "qosMacClassificationTable"
    qosMacClassificationTable.EntityData.BundleName = "cisco_ios_xe"
    qosMacClassificationTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosMacClassificationTable.EntityData.SegmentPath = "qosMacClassificationTable"
    qosMacClassificationTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosMacClassificationTable.EntityData.SegmentPath
    qosMacClassificationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosMacClassificationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosMacClassificationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosMacClassificationTable.EntityData.Children = types.NewOrderedMap()
    qosMacClassificationTable.EntityData.Children.Append("qosMacClassificationEntry", types.YChild{"QosMacClassificationEntry", nil})
    for i := range qosMacClassificationTable.QosMacClassificationEntry {
        qosMacClassificationTable.EntityData.Children.Append(types.GetSegmentPath(qosMacClassificationTable.QosMacClassificationEntry[i]), types.YChild{"QosMacClassificationEntry", qosMacClassificationTable.QosMacClassificationEntry[i]})
    }
    qosMacClassificationTable.EntityData.Leafs = types.NewOrderedMap()

    qosMacClassificationTable.EntityData.YListKeys = []string {}

    return &(qosMacClassificationTable.EntityData)
}

// CISCOQOSPIBMIB_QosMacClassificationTable_QosMacClassificationEntry
// An instance of this class specifies the mapping of a VLAN
// and a MAC address to a CoS value.
type CISCOQOSPIBMIB_QosMacClassificationTable_QosMacClassificationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosMacClassificationId interface{}

    // The VLAN of the destination MAC address of the L2 frame. The type is
    // interface{} with range: 1..4095.
    QosDstMacVlan interface{}

    // The destination MAC address of the L2 frame. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    QosDstMacAddress interface{}

    // The CoS to assign the packet with the associated MAC/VLAN tuple.  Note that
    // this CoS is overridden by the policies to classify the frame at layer 3 if
    // there are any. The type is interface{} with range: 0..7.
    QosDstMacCos interface{}
}

func (qosMacClassificationEntry *CISCOQOSPIBMIB_QosMacClassificationTable_QosMacClassificationEntry) GetEntityData() *types.CommonEntityData {
    qosMacClassificationEntry.EntityData.YFilter = qosMacClassificationEntry.YFilter
    qosMacClassificationEntry.EntityData.YangName = "qosMacClassificationEntry"
    qosMacClassificationEntry.EntityData.BundleName = "cisco_ios_xe"
    qosMacClassificationEntry.EntityData.ParentYangName = "qosMacClassificationTable"
    qosMacClassificationEntry.EntityData.SegmentPath = "qosMacClassificationEntry" + types.AddKeyToken(qosMacClassificationEntry.QosMacClassificationId, "qosMacClassificationId")
    qosMacClassificationEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosMacClassificationTable/" + qosMacClassificationEntry.EntityData.SegmentPath
    qosMacClassificationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosMacClassificationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosMacClassificationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosMacClassificationEntry.EntityData.Children = types.NewOrderedMap()
    qosMacClassificationEntry.EntityData.Leafs = types.NewOrderedMap()
    qosMacClassificationEntry.EntityData.Leafs.Append("qosMacClassificationId", types.YLeaf{"QosMacClassificationId", qosMacClassificationEntry.QosMacClassificationId})
    qosMacClassificationEntry.EntityData.Leafs.Append("qosDstMacVlan", types.YLeaf{"QosDstMacVlan", qosMacClassificationEntry.QosDstMacVlan})
    qosMacClassificationEntry.EntityData.Leafs.Append("qosDstMacAddress", types.YLeaf{"QosDstMacAddress", qosMacClassificationEntry.QosDstMacAddress})
    qosMacClassificationEntry.EntityData.Leafs.Append("qosDstMacCos", types.YLeaf{"QosDstMacCos", qosMacClassificationEntry.QosDstMacCos})

    qosMacClassificationEntry.EntityData.YListKeys = []string {"QosMacClassificationId"}

    return &(qosMacClassificationEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIpAceTable
// ACE definitions.
type CISCOQOSPIBMIB_QosIpAceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies an ACE. The type is slice of
    // CISCOQOSPIBMIB_QosIpAceTable_QosIpAceEntry.
    QosIpAceEntry []*CISCOQOSPIBMIB_QosIpAceTable_QosIpAceEntry
}

func (qosIpAceTable *CISCOQOSPIBMIB_QosIpAceTable) GetEntityData() *types.CommonEntityData {
    qosIpAceTable.EntityData.YFilter = qosIpAceTable.YFilter
    qosIpAceTable.EntityData.YangName = "qosIpAceTable"
    qosIpAceTable.EntityData.BundleName = "cisco_ios_xe"
    qosIpAceTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIpAceTable.EntityData.SegmentPath = "qosIpAceTable"
    qosIpAceTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIpAceTable.EntityData.SegmentPath
    qosIpAceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIpAceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIpAceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIpAceTable.EntityData.Children = types.NewOrderedMap()
    qosIpAceTable.EntityData.Children.Append("qosIpAceEntry", types.YChild{"QosIpAceEntry", nil})
    for i := range qosIpAceTable.QosIpAceEntry {
        qosIpAceTable.EntityData.Children.Append(types.GetSegmentPath(qosIpAceTable.QosIpAceEntry[i]), types.YChild{"QosIpAceEntry", qosIpAceTable.QosIpAceEntry[i]})
    }
    qosIpAceTable.EntityData.Leafs = types.NewOrderedMap()

    qosIpAceTable.EntityData.YListKeys = []string {}

    return &(qosIpAceTable.EntityData)
}

// CISCOQOSPIBMIB_QosIpAceTable_QosIpAceEntry
// An instance of this class specifies an ACE.
type CISCOQOSPIBMIB_QosIpAceTable_QosIpAceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIpAceId interface{}

    // The IP address to match against the packet's destination IP address. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    QosIpAceDstAddr interface{}

    // A mask for the matching of the destination IP address. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    QosIpAceDstAddrMask interface{}

    // The IP address to match against the packet's source IP address. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    QosIpAceSrcAddr interface{}

    // A mask for the matching of the source IP address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    QosIpAceSrcAddrMask interface{}

    // The minimum value that the DSCP in the packet can have and match this ACE.
    // The type is interface{} with range: 0..63.
    QosIpAceDscpMin interface{}

    // The maximum value that the DSCP in the packet can have and match this ACE.
    // The type is interface{} with range: 0..63.
    QosIpAceDscpMax interface{}

    // The IP protocol to match against the packet's protocol. A value of zero
    // means match all. The type is interface{} with range: 0..255.
    QosIpAceProtocol interface{}

    // The minimum value that the packet's layer 4 dest port number can have and
    // match this ACE. The type is interface{} with range: 0..65535.
    QosIpAceDstL4PortMin interface{}

    // The maximum value that the packet's layer 4 dest port number can have and
    // match this ACE. The type is interface{} with range: 0..65535.
    QosIpAceDstL4PortMax interface{}

    // The minimum value that the packet's layer 4 source port number can have and
    // match this ACE. The type is interface{} with range: 0..65535.
    QosIpAceSrcL4PortMin interface{}

    // The maximum value that the packet's layer 4 source port number can have and
    // match this ACE. The type is interface{} with range: 0..65535.
    QosIpAceSrcL4PortMax interface{}

    // If the packet matches this ACE and the value of this attribute is true,
    // then the matching process terminates and the QoS associated with this ACE
    // (indirectly through the ACL) is applied to the packet.  If the value of
    // this attribute is false, then no more ACEs in this ACL are compared to this
    // packet and matching continues with the first ACE of the next ACL. The type
    // is bool.
    QosIpAcePermit interface{}
}

func (qosIpAceEntry *CISCOQOSPIBMIB_QosIpAceTable_QosIpAceEntry) GetEntityData() *types.CommonEntityData {
    qosIpAceEntry.EntityData.YFilter = qosIpAceEntry.YFilter
    qosIpAceEntry.EntityData.YangName = "qosIpAceEntry"
    qosIpAceEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIpAceEntry.EntityData.ParentYangName = "qosIpAceTable"
    qosIpAceEntry.EntityData.SegmentPath = "qosIpAceEntry" + types.AddKeyToken(qosIpAceEntry.QosIpAceId, "qosIpAceId")
    qosIpAceEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIpAceTable/" + qosIpAceEntry.EntityData.SegmentPath
    qosIpAceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIpAceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIpAceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIpAceEntry.EntityData.Children = types.NewOrderedMap()
    qosIpAceEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceId", types.YLeaf{"QosIpAceId", qosIpAceEntry.QosIpAceId})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceDstAddr", types.YLeaf{"QosIpAceDstAddr", qosIpAceEntry.QosIpAceDstAddr})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceDstAddrMask", types.YLeaf{"QosIpAceDstAddrMask", qosIpAceEntry.QosIpAceDstAddrMask})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceSrcAddr", types.YLeaf{"QosIpAceSrcAddr", qosIpAceEntry.QosIpAceSrcAddr})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceSrcAddrMask", types.YLeaf{"QosIpAceSrcAddrMask", qosIpAceEntry.QosIpAceSrcAddrMask})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceDscpMin", types.YLeaf{"QosIpAceDscpMin", qosIpAceEntry.QosIpAceDscpMin})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceDscpMax", types.YLeaf{"QosIpAceDscpMax", qosIpAceEntry.QosIpAceDscpMax})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceProtocol", types.YLeaf{"QosIpAceProtocol", qosIpAceEntry.QosIpAceProtocol})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceDstL4PortMin", types.YLeaf{"QosIpAceDstL4PortMin", qosIpAceEntry.QosIpAceDstL4PortMin})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceDstL4PortMax", types.YLeaf{"QosIpAceDstL4PortMax", qosIpAceEntry.QosIpAceDstL4PortMax})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceSrcL4PortMin", types.YLeaf{"QosIpAceSrcL4PortMin", qosIpAceEntry.QosIpAceSrcL4PortMin})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAceSrcL4PortMax", types.YLeaf{"QosIpAceSrcL4PortMax", qosIpAceEntry.QosIpAceSrcL4PortMax})
    qosIpAceEntry.EntityData.Leafs.Append("qosIpAcePermit", types.YLeaf{"QosIpAcePermit", qosIpAceEntry.QosIpAcePermit})

    qosIpAceEntry.EntityData.YListKeys = []string {"QosIpAceId"}

    return &(qosIpAceEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIpAclDefinitionTable
// A class that defines a set of ACLs each being an ordered list
// of ACEs.
type CISCOQOSPIBMIB_QosIpAclDefinitionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies an ACE in an ACL and its order with
    // respect to other ACEs in the same ACL. The type is slice of
    // CISCOQOSPIBMIB_QosIpAclDefinitionTable_QosIpAclDefinitionEntry.
    QosIpAclDefinitionEntry []*CISCOQOSPIBMIB_QosIpAclDefinitionTable_QosIpAclDefinitionEntry
}

func (qosIpAclDefinitionTable *CISCOQOSPIBMIB_QosIpAclDefinitionTable) GetEntityData() *types.CommonEntityData {
    qosIpAclDefinitionTable.EntityData.YFilter = qosIpAclDefinitionTable.YFilter
    qosIpAclDefinitionTable.EntityData.YangName = "qosIpAclDefinitionTable"
    qosIpAclDefinitionTable.EntityData.BundleName = "cisco_ios_xe"
    qosIpAclDefinitionTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIpAclDefinitionTable.EntityData.SegmentPath = "qosIpAclDefinitionTable"
    qosIpAclDefinitionTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIpAclDefinitionTable.EntityData.SegmentPath
    qosIpAclDefinitionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIpAclDefinitionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIpAclDefinitionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIpAclDefinitionTable.EntityData.Children = types.NewOrderedMap()
    qosIpAclDefinitionTable.EntityData.Children.Append("qosIpAclDefinitionEntry", types.YChild{"QosIpAclDefinitionEntry", nil})
    for i := range qosIpAclDefinitionTable.QosIpAclDefinitionEntry {
        qosIpAclDefinitionTable.EntityData.Children.Append(types.GetSegmentPath(qosIpAclDefinitionTable.QosIpAclDefinitionEntry[i]), types.YChild{"QosIpAclDefinitionEntry", qosIpAclDefinitionTable.QosIpAclDefinitionEntry[i]})
    }
    qosIpAclDefinitionTable.EntityData.Leafs = types.NewOrderedMap()

    qosIpAclDefinitionTable.EntityData.YListKeys = []string {}

    return &(qosIpAclDefinitionTable.EntityData)
}

// CISCOQOSPIBMIB_QosIpAclDefinitionTable_QosIpAclDefinitionEntry
// An instance of this class specifies an ACE in an ACL and its
// order with respect to other ACEs in the same ACL.
type CISCOQOSPIBMIB_QosIpAclDefinitionTable_QosIpAclDefinitionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIpAclDefinitionId interface{}

    // An index for this ACL.  There will be one instance of policy class
    // qosIpAclDefinition with this integer index for each ACE in the ACL per role
    // combination. The type is interface{} with range: 0..4294967295.
    QosIpAclId interface{}

    // An integer that determines the position of this ACE in the ACL. An ACE with
    // a given order is positioned in the access contol list before one with a
    // higher order. The type is interface{} with range: 0..4294967295.
    QosIpAceOrder interface{}

    // This attribute specifies the ACE in the qosIpAceTable that is in the ACL
    // specified by qosIpAclId at the position specified by qosIpAceOrder. The
    // type is interface{} with range: 0..4294967295.
    QosIpAclDefAceId interface{}
}

func (qosIpAclDefinitionEntry *CISCOQOSPIBMIB_QosIpAclDefinitionTable_QosIpAclDefinitionEntry) GetEntityData() *types.CommonEntityData {
    qosIpAclDefinitionEntry.EntityData.YFilter = qosIpAclDefinitionEntry.YFilter
    qosIpAclDefinitionEntry.EntityData.YangName = "qosIpAclDefinitionEntry"
    qosIpAclDefinitionEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIpAclDefinitionEntry.EntityData.ParentYangName = "qosIpAclDefinitionTable"
    qosIpAclDefinitionEntry.EntityData.SegmentPath = "qosIpAclDefinitionEntry" + types.AddKeyToken(qosIpAclDefinitionEntry.QosIpAclDefinitionId, "qosIpAclDefinitionId")
    qosIpAclDefinitionEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIpAclDefinitionTable/" + qosIpAclDefinitionEntry.EntityData.SegmentPath
    qosIpAclDefinitionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIpAclDefinitionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIpAclDefinitionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIpAclDefinitionEntry.EntityData.Children = types.NewOrderedMap()
    qosIpAclDefinitionEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIpAclDefinitionEntry.EntityData.Leafs.Append("qosIpAclDefinitionId", types.YLeaf{"QosIpAclDefinitionId", qosIpAclDefinitionEntry.QosIpAclDefinitionId})
    qosIpAclDefinitionEntry.EntityData.Leafs.Append("qosIpAclId", types.YLeaf{"QosIpAclId", qosIpAclDefinitionEntry.QosIpAclId})
    qosIpAclDefinitionEntry.EntityData.Leafs.Append("qosIpAceOrder", types.YLeaf{"QosIpAceOrder", qosIpAclDefinitionEntry.QosIpAceOrder})
    qosIpAclDefinitionEntry.EntityData.Leafs.Append("qosIpAclDefAceId", types.YLeaf{"QosIpAclDefAceId", qosIpAclDefinitionEntry.QosIpAclDefAceId})

    qosIpAclDefinitionEntry.EntityData.YListKeys = []string {"QosIpAclDefinitionId"}

    return &(qosIpAclDefinitionEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIpAclActionTable
// A class that applies a set of ACLs to interfaces specifying,
// for each interface the order of the ACL with respect to other
// ACLs applied to the same interface and, for each ACL the
// action to take for a packet that matches a permit ACE in that
// ACL.  Interfaces are specified abstractly in terms of
// interface role combinations.
type CISCOQOSPIBMIB_QosIpAclActionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class applies an ACL to traffic in a particular
    // direction on an interface with a particular role combination, and specifies
    // the action for packets which match the ACL. The type is slice of
    // CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry.
    QosIpAclActionEntry []*CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry
}

func (qosIpAclActionTable *CISCOQOSPIBMIB_QosIpAclActionTable) GetEntityData() *types.CommonEntityData {
    qosIpAclActionTable.EntityData.YFilter = qosIpAclActionTable.YFilter
    qosIpAclActionTable.EntityData.YangName = "qosIpAclActionTable"
    qosIpAclActionTable.EntityData.BundleName = "cisco_ios_xe"
    qosIpAclActionTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIpAclActionTable.EntityData.SegmentPath = "qosIpAclActionTable"
    qosIpAclActionTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIpAclActionTable.EntityData.SegmentPath
    qosIpAclActionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIpAclActionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIpAclActionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIpAclActionTable.EntityData.Children = types.NewOrderedMap()
    qosIpAclActionTable.EntityData.Children.Append("qosIpAclActionEntry", types.YChild{"QosIpAclActionEntry", nil})
    for i := range qosIpAclActionTable.QosIpAclActionEntry {
        qosIpAclActionTable.EntityData.Children.Append(types.GetSegmentPath(qosIpAclActionTable.QosIpAclActionEntry[i]), types.YChild{"QosIpAclActionEntry", qosIpAclActionTable.QosIpAclActionEntry[i]})
    }
    qosIpAclActionTable.EntityData.Leafs = types.NewOrderedMap()

    qosIpAclActionTable.EntityData.YListKeys = []string {}

    return &(qosIpAclActionTable.EntityData)
}

// CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry
// An instance of this class applies an ACL to traffic in a
// particular direction on an interface with a particular role
// combination, and specifies the action for packets which match
// the ACL.
type CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIpAclActionId interface{}

    // The ACL associated with this action. The type is interface{} with range:
    // 0..4294967295.
    QosIpAclActAclId interface{}

    // The interfaces to which this ACL applies specified in terms of a set of
    // roles. The type is string with length: 0..255.
    QosIpAclInterfaceRoles interface{}

    // The direction of packet flow at the interface in question to which this ACL
    // applies. The type is QosIpAclInterfaceDirection.
    QosIpAclInterfaceDirection interface{}

    // An integer that determines the order of this ACL in the list of ACLs
    // applied to interfaces of the specified role combination. An ACL with a
    // given order is positioned in the list before one with a higher order. The
    // type is interface{} with range: 0..4294967295.
    QosIpAclOrder interface{}

    // The DSCP to classify the packet with in the event that the packet matches
    // an ACE in this ACL and the ACE is a permit. The type is interface{} with
    // range: 0..63.
    QosIpAclDscp interface{}

    // If this attribute is true, then the Dscp associated with the packet is
    // trusted, i.e., it is assumed to have already been set.  In this case, the
    // Dscp is not rewritten with qosIpAclDscp (qosIpAclDscp is ignored).  The
    // packet is still policed as part of its micro flow and its aggregate flow. 
    // When a trusted action is applied to an input interface, the Dscp associated
    // with the packet is the one contained in the packet.  When a trusted action
    // is applied to an output interface, the Dscp associated with the packet is
    // the one that is the result of the input classification and policing. The
    // type is bool.
    QosIpAclDscpTrusted interface{}

    // An index identifying the instance of policer to apply to the microflow.  It
    // must correspond to the integer index of an instance of class
    // qosPolicerTableor be zero.  If zero, the microflow is not policed. The type
    // is interface{} with range: 0..4294967295.
    QosIpAclMicroFlowPolicerId interface{}

    // An index identifying the aggregate that the packet belongs to.  It must
    // correspond to the integer index of an instance of class qosAggregateTable
    // or be zero.  If zero, the microflow does not belong to any aggregate and is
    // not policed as part of any aggregate. The type is interface{} with range:
    // 0..4294967295.
    QosIpAclAggregateId interface{}
}

func (qosIpAclActionEntry *CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry) GetEntityData() *types.CommonEntityData {
    qosIpAclActionEntry.EntityData.YFilter = qosIpAclActionEntry.YFilter
    qosIpAclActionEntry.EntityData.YangName = "qosIpAclActionEntry"
    qosIpAclActionEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIpAclActionEntry.EntityData.ParentYangName = "qosIpAclActionTable"
    qosIpAclActionEntry.EntityData.SegmentPath = "qosIpAclActionEntry" + types.AddKeyToken(qosIpAclActionEntry.QosIpAclActionId, "qosIpAclActionId")
    qosIpAclActionEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIpAclActionTable/" + qosIpAclActionEntry.EntityData.SegmentPath
    qosIpAclActionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIpAclActionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIpAclActionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIpAclActionEntry.EntityData.Children = types.NewOrderedMap()
    qosIpAclActionEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclActionId", types.YLeaf{"QosIpAclActionId", qosIpAclActionEntry.QosIpAclActionId})
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclActAclId", types.YLeaf{"QosIpAclActAclId", qosIpAclActionEntry.QosIpAclActAclId})
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclInterfaceRoles", types.YLeaf{"QosIpAclInterfaceRoles", qosIpAclActionEntry.QosIpAclInterfaceRoles})
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclInterfaceDirection", types.YLeaf{"QosIpAclInterfaceDirection", qosIpAclActionEntry.QosIpAclInterfaceDirection})
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclOrder", types.YLeaf{"QosIpAclOrder", qosIpAclActionEntry.QosIpAclOrder})
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclDscp", types.YLeaf{"QosIpAclDscp", qosIpAclActionEntry.QosIpAclDscp})
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclDscpTrusted", types.YLeaf{"QosIpAclDscpTrusted", qosIpAclActionEntry.QosIpAclDscpTrusted})
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclMicroFlowPolicerId", types.YLeaf{"QosIpAclMicroFlowPolicerId", qosIpAclActionEntry.QosIpAclMicroFlowPolicerId})
    qosIpAclActionEntry.EntityData.Leafs.Append("qosIpAclAggregateId", types.YLeaf{"QosIpAclAggregateId", qosIpAclActionEntry.QosIpAclAggregateId})

    qosIpAclActionEntry.EntityData.YListKeys = []string {"QosIpAclActionId"}

    return &(qosIpAclActionEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry_QosIpAclInterfaceDirection represents which this ACL applies.
type CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry_QosIpAclInterfaceDirection string

const (
    CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry_QosIpAclInterfaceDirection_in CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry_QosIpAclInterfaceDirection = "in"

    CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry_QosIpAclInterfaceDirection_out CISCOQOSPIBMIB_QosIpAclActionTable_QosIpAclActionEntry_QosIpAclInterfaceDirection = "out"
)

// CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable
// This class specifies the scheduling preference an interface
// chooses if it supports multiple scheduling types.  Higher
// values are preferred over lower values.
type CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies a scheduling preference for a
    // queue-type on an interface with a particular role combination. The type is
    // slice of
    // CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry.
    QosIfSchedulingPreferenceEntry []*CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry
}

func (qosIfSchedulingPreferencesTable *CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable) GetEntityData() *types.CommonEntityData {
    qosIfSchedulingPreferencesTable.EntityData.YFilter = qosIfSchedulingPreferencesTable.YFilter
    qosIfSchedulingPreferencesTable.EntityData.YangName = "qosIfSchedulingPreferencesTable"
    qosIfSchedulingPreferencesTable.EntityData.BundleName = "cisco_ios_xe"
    qosIfSchedulingPreferencesTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIfSchedulingPreferencesTable.EntityData.SegmentPath = "qosIfSchedulingPreferencesTable"
    qosIfSchedulingPreferencesTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIfSchedulingPreferencesTable.EntityData.SegmentPath
    qosIfSchedulingPreferencesTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfSchedulingPreferencesTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfSchedulingPreferencesTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfSchedulingPreferencesTable.EntityData.Children = types.NewOrderedMap()
    qosIfSchedulingPreferencesTable.EntityData.Children.Append("qosIfSchedulingPreferenceEntry", types.YChild{"QosIfSchedulingPreferenceEntry", nil})
    for i := range qosIfSchedulingPreferencesTable.QosIfSchedulingPreferenceEntry {
        qosIfSchedulingPreferencesTable.EntityData.Children.Append(types.GetSegmentPath(qosIfSchedulingPreferencesTable.QosIfSchedulingPreferenceEntry[i]), types.YChild{"QosIfSchedulingPreferenceEntry", qosIfSchedulingPreferencesTable.QosIfSchedulingPreferenceEntry[i]})
    }
    qosIfSchedulingPreferencesTable.EntityData.Leafs = types.NewOrderedMap()

    qosIfSchedulingPreferencesTable.EntityData.YListKeys = []string {}

    return &(qosIfSchedulingPreferencesTable.EntityData)
}

// CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry
// An instance of this class specifies a scheduling preference
// for a queue-type on an interface with a particular role
// combination.
type CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIfSchedulingPreferenceId interface{}

    // The combination of roles the interface must have for this policy instance
    // to apply to that interface. The type is string with length: 0..255.
    QosIfSchedulingRoles interface{}

    // The preference to use this scheduling discipline and queue type.  A higher
    // value means a higher preference.  If two disciplines have the same
    // preference the choice is a local decision. The type is interface{} with
    // range: 1..16.
    QosIfSchedulingPreference interface{}

    // An enumerate type for all the known scheduling disciplines. The type is
    // QosIfSchedulingDiscipline.
    QosIfSchedulingDiscipline interface{}

    // The queue type of this preference. The type is QosInterfaceQueueType.
    QosIfSchedulingQueueType interface{}
}

func (qosIfSchedulingPreferenceEntry *CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry) GetEntityData() *types.CommonEntityData {
    qosIfSchedulingPreferenceEntry.EntityData.YFilter = qosIfSchedulingPreferenceEntry.YFilter
    qosIfSchedulingPreferenceEntry.EntityData.YangName = "qosIfSchedulingPreferenceEntry"
    qosIfSchedulingPreferenceEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIfSchedulingPreferenceEntry.EntityData.ParentYangName = "qosIfSchedulingPreferencesTable"
    qosIfSchedulingPreferenceEntry.EntityData.SegmentPath = "qosIfSchedulingPreferenceEntry" + types.AddKeyToken(qosIfSchedulingPreferenceEntry.QosIfSchedulingPreferenceId, "qosIfSchedulingPreferenceId")
    qosIfSchedulingPreferenceEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIfSchedulingPreferencesTable/" + qosIfSchedulingPreferenceEntry.EntityData.SegmentPath
    qosIfSchedulingPreferenceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfSchedulingPreferenceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfSchedulingPreferenceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfSchedulingPreferenceEntry.EntityData.Children = types.NewOrderedMap()
    qosIfSchedulingPreferenceEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIfSchedulingPreferenceEntry.EntityData.Leafs.Append("qosIfSchedulingPreferenceId", types.YLeaf{"QosIfSchedulingPreferenceId", qosIfSchedulingPreferenceEntry.QosIfSchedulingPreferenceId})
    qosIfSchedulingPreferenceEntry.EntityData.Leafs.Append("qosIfSchedulingRoles", types.YLeaf{"QosIfSchedulingRoles", qosIfSchedulingPreferenceEntry.QosIfSchedulingRoles})
    qosIfSchedulingPreferenceEntry.EntityData.Leafs.Append("qosIfSchedulingPreference", types.YLeaf{"QosIfSchedulingPreference", qosIfSchedulingPreferenceEntry.QosIfSchedulingPreference})
    qosIfSchedulingPreferenceEntry.EntityData.Leafs.Append("qosIfSchedulingDiscipline", types.YLeaf{"QosIfSchedulingDiscipline", qosIfSchedulingPreferenceEntry.QosIfSchedulingDiscipline})
    qosIfSchedulingPreferenceEntry.EntityData.Leafs.Append("qosIfSchedulingQueueType", types.YLeaf{"QosIfSchedulingQueueType", qosIfSchedulingPreferenceEntry.QosIfSchedulingQueueType})

    qosIfSchedulingPreferenceEntry.EntityData.YListKeys = []string {"QosIfSchedulingPreferenceId"}

    return &(qosIfSchedulingPreferenceEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline represents An enumerate type for all the known scheduling disciplines.
type CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline string

const (
    CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline_weightedFairQueueing CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline = "weightedFairQueueing"

    CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline_weightedRoundRobin CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline = "weightedRoundRobin"

    CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline_customQueueing CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline = "customQueueing"

    CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline_priorityQueueing CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline = "priorityQueueing"

    CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline_classBasedWFQ CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline = "classBasedWFQ"

    CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline_fifo CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline = "fifo"

    CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline_pqWrr CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline = "pqWrr"

    CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline_pqCbwfq CISCOQOSPIBMIB_QosIfSchedulingPreferencesTable_QosIfSchedulingPreferenceEntry_QosIfSchedulingDiscipline = "pqCbwfq"
)

// CISCOQOSPIBMIB_QosIfDropPreferenceTable
// This class specifies the preference of the drop mechanism an
// interface chooses if it supports multiple drop mechanisms.
// Higher values are preferred over lower values.
type CISCOQOSPIBMIB_QosIfDropPreferenceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies a drop preference for a drop mechanism
    // on an interface with a particular role combination. The type is slice of
    // CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry.
    QosIfDropPreferenceEntry []*CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry
}

func (qosIfDropPreferenceTable *CISCOQOSPIBMIB_QosIfDropPreferenceTable) GetEntityData() *types.CommonEntityData {
    qosIfDropPreferenceTable.EntityData.YFilter = qosIfDropPreferenceTable.YFilter
    qosIfDropPreferenceTable.EntityData.YangName = "qosIfDropPreferenceTable"
    qosIfDropPreferenceTable.EntityData.BundleName = "cisco_ios_xe"
    qosIfDropPreferenceTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIfDropPreferenceTable.EntityData.SegmentPath = "qosIfDropPreferenceTable"
    qosIfDropPreferenceTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIfDropPreferenceTable.EntityData.SegmentPath
    qosIfDropPreferenceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfDropPreferenceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfDropPreferenceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfDropPreferenceTable.EntityData.Children = types.NewOrderedMap()
    qosIfDropPreferenceTable.EntityData.Children.Append("qosIfDropPreferenceEntry", types.YChild{"QosIfDropPreferenceEntry", nil})
    for i := range qosIfDropPreferenceTable.QosIfDropPreferenceEntry {
        qosIfDropPreferenceTable.EntityData.Children.Append(types.GetSegmentPath(qosIfDropPreferenceTable.QosIfDropPreferenceEntry[i]), types.YChild{"QosIfDropPreferenceEntry", qosIfDropPreferenceTable.QosIfDropPreferenceEntry[i]})
    }
    qosIfDropPreferenceTable.EntityData.Leafs = types.NewOrderedMap()

    qosIfDropPreferenceTable.EntityData.YListKeys = []string {}

    return &(qosIfDropPreferenceTable.EntityData)
}

// CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry
// An instance of this class specifies a drop preference for
// a drop mechanism on an interface with a particular role
// combination.
type CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIfDropPreferenceId interface{}

    // The combination of roles the interface must have for this policy instance
    // to apply to that interface. The type is string with length: 0..255.
    QosIfDropRoles interface{}

    // The preference to use this drop mechanism.  A higher value means a higher
    // preference.  If two mechanisms have the same preference the choice is a
    // local decision. The type is interface{} with range: 1..16.
    QosIfDropPreference interface{}

    // An enumerate type for all the known drop mechanisms. The type is
    // QosIfDropDiscipline.
    QosIfDropDiscipline interface{}
}

func (qosIfDropPreferenceEntry *CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry) GetEntityData() *types.CommonEntityData {
    qosIfDropPreferenceEntry.EntityData.YFilter = qosIfDropPreferenceEntry.YFilter
    qosIfDropPreferenceEntry.EntityData.YangName = "qosIfDropPreferenceEntry"
    qosIfDropPreferenceEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIfDropPreferenceEntry.EntityData.ParentYangName = "qosIfDropPreferenceTable"
    qosIfDropPreferenceEntry.EntityData.SegmentPath = "qosIfDropPreferenceEntry" + types.AddKeyToken(qosIfDropPreferenceEntry.QosIfDropPreferenceId, "qosIfDropPreferenceId")
    qosIfDropPreferenceEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIfDropPreferenceTable/" + qosIfDropPreferenceEntry.EntityData.SegmentPath
    qosIfDropPreferenceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfDropPreferenceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfDropPreferenceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfDropPreferenceEntry.EntityData.Children = types.NewOrderedMap()
    qosIfDropPreferenceEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIfDropPreferenceEntry.EntityData.Leafs.Append("qosIfDropPreferenceId", types.YLeaf{"QosIfDropPreferenceId", qosIfDropPreferenceEntry.QosIfDropPreferenceId})
    qosIfDropPreferenceEntry.EntityData.Leafs.Append("qosIfDropRoles", types.YLeaf{"QosIfDropRoles", qosIfDropPreferenceEntry.QosIfDropRoles})
    qosIfDropPreferenceEntry.EntityData.Leafs.Append("qosIfDropPreference", types.YLeaf{"QosIfDropPreference", qosIfDropPreferenceEntry.QosIfDropPreference})
    qosIfDropPreferenceEntry.EntityData.Leafs.Append("qosIfDropDiscipline", types.YLeaf{"QosIfDropDiscipline", qosIfDropPreferenceEntry.QosIfDropDiscipline})

    qosIfDropPreferenceEntry.EntityData.YListKeys = []string {"QosIfDropPreferenceId"}

    return &(qosIfDropPreferenceEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry_QosIfDropDiscipline represents An enumerate type for all the known drop mechanisms.
type CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry_QosIfDropDiscipline string

const (
    CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry_QosIfDropDiscipline_qosIfDropWRED CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry_QosIfDropDiscipline = "qosIfDropWRED"

    CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry_QosIfDropDiscipline_qosIfDropTailDrop CISCOQOSPIBMIB_QosIfDropPreferenceTable_QosIfDropPreferenceEntry_QosIfDropDiscipline = "qosIfDropTailDrop"
)

// CISCOQOSPIBMIB_QosIfDscpAssignmentTable
// The assignment of each DSCP to a queue and threshold for each
// interface queue type.
type CISCOQOSPIBMIB_QosIfDscpAssignmentTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the queue and threshold set for a
    // packet with a particular DSCP on an interface of a particular type with a
    // particular role combination. The type is slice of
    // CISCOQOSPIBMIB_QosIfDscpAssignmentTable_QosIfDscpAssignmentEntry.
    QosIfDscpAssignmentEntry []*CISCOQOSPIBMIB_QosIfDscpAssignmentTable_QosIfDscpAssignmentEntry
}

func (qosIfDscpAssignmentTable *CISCOQOSPIBMIB_QosIfDscpAssignmentTable) GetEntityData() *types.CommonEntityData {
    qosIfDscpAssignmentTable.EntityData.YFilter = qosIfDscpAssignmentTable.YFilter
    qosIfDscpAssignmentTable.EntityData.YangName = "qosIfDscpAssignmentTable"
    qosIfDscpAssignmentTable.EntityData.BundleName = "cisco_ios_xe"
    qosIfDscpAssignmentTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIfDscpAssignmentTable.EntityData.SegmentPath = "qosIfDscpAssignmentTable"
    qosIfDscpAssignmentTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIfDscpAssignmentTable.EntityData.SegmentPath
    qosIfDscpAssignmentTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfDscpAssignmentTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfDscpAssignmentTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfDscpAssignmentTable.EntityData.Children = types.NewOrderedMap()
    qosIfDscpAssignmentTable.EntityData.Children.Append("qosIfDscpAssignmentEntry", types.YChild{"QosIfDscpAssignmentEntry", nil})
    for i := range qosIfDscpAssignmentTable.QosIfDscpAssignmentEntry {
        qosIfDscpAssignmentTable.EntityData.Children.Append(types.GetSegmentPath(qosIfDscpAssignmentTable.QosIfDscpAssignmentEntry[i]), types.YChild{"QosIfDscpAssignmentEntry", qosIfDscpAssignmentTable.QosIfDscpAssignmentEntry[i]})
    }
    qosIfDscpAssignmentTable.EntityData.Leafs = types.NewOrderedMap()

    qosIfDscpAssignmentTable.EntityData.YListKeys = []string {}

    return &(qosIfDscpAssignmentTable.EntityData)
}

// CISCOQOSPIBMIB_QosIfDscpAssignmentTable_QosIfDscpAssignmentEntry
// An instance of this class specifies the queue and threshold
// set for a packet with a particular DSCP on an interface of
// a particular type with a particular role combination.
type CISCOQOSPIBMIB_QosIfDscpAssignmentTable_QosIfDscpAssignmentEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIfDscpAssignmentId interface{}

    // The role combination the interface must be configured with. The type is
    // string with length: 0..255.
    QosIfDscpRoles interface{}

    // The interface queue type to which this row applies. The type is
    // QosInterfaceQueueType.
    QosIfQueueType interface{}

    // The DSCP to which this row applies. The type is interface{} with range:
    // 0..63.
    QosIfDscp interface{}

    // The queue to which the DSCP applies for the given interface type. The type
    // is interface{} with range: 1..64.
    QosIfQueue interface{}

    // The threshold set of the specified queue to which the DSCP applies for the
    // given interface type. The type is interface{} with range: 1..8.
    QosIfThresholdSet interface{}
}

func (qosIfDscpAssignmentEntry *CISCOQOSPIBMIB_QosIfDscpAssignmentTable_QosIfDscpAssignmentEntry) GetEntityData() *types.CommonEntityData {
    qosIfDscpAssignmentEntry.EntityData.YFilter = qosIfDscpAssignmentEntry.YFilter
    qosIfDscpAssignmentEntry.EntityData.YangName = "qosIfDscpAssignmentEntry"
    qosIfDscpAssignmentEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIfDscpAssignmentEntry.EntityData.ParentYangName = "qosIfDscpAssignmentTable"
    qosIfDscpAssignmentEntry.EntityData.SegmentPath = "qosIfDscpAssignmentEntry" + types.AddKeyToken(qosIfDscpAssignmentEntry.QosIfDscpAssignmentId, "qosIfDscpAssignmentId")
    qosIfDscpAssignmentEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIfDscpAssignmentTable/" + qosIfDscpAssignmentEntry.EntityData.SegmentPath
    qosIfDscpAssignmentEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfDscpAssignmentEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfDscpAssignmentEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfDscpAssignmentEntry.EntityData.Children = types.NewOrderedMap()
    qosIfDscpAssignmentEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIfDscpAssignmentEntry.EntityData.Leafs.Append("qosIfDscpAssignmentId", types.YLeaf{"QosIfDscpAssignmentId", qosIfDscpAssignmentEntry.QosIfDscpAssignmentId})
    qosIfDscpAssignmentEntry.EntityData.Leafs.Append("qosIfDscpRoles", types.YLeaf{"QosIfDscpRoles", qosIfDscpAssignmentEntry.QosIfDscpRoles})
    qosIfDscpAssignmentEntry.EntityData.Leafs.Append("qosIfQueueType", types.YLeaf{"QosIfQueueType", qosIfDscpAssignmentEntry.QosIfQueueType})
    qosIfDscpAssignmentEntry.EntityData.Leafs.Append("qosIfDscp", types.YLeaf{"QosIfDscp", qosIfDscpAssignmentEntry.QosIfDscp})
    qosIfDscpAssignmentEntry.EntityData.Leafs.Append("qosIfQueue", types.YLeaf{"QosIfQueue", qosIfDscpAssignmentEntry.QosIfQueue})
    qosIfDscpAssignmentEntry.EntityData.Leafs.Append("qosIfThresholdSet", types.YLeaf{"QosIfThresholdSet", qosIfDscpAssignmentEntry.QosIfThresholdSet})

    qosIfDscpAssignmentEntry.EntityData.YListKeys = []string {"QosIfDscpAssignmentId"}

    return &(qosIfDscpAssignmentEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIfRedTable
// A class of lower and upper values for each threshold set in a
// queue supporting WRED.  If the size of the queue for a given
// threshold is below the lower value then packets assigned to
// that threshold are always accepted into the queue.  If the
// size of the queue is above upper value then packets are always
// dropped.  If the size of the queue is between the lower and
// the upper then packets are randomly dropped.
type CISCOQOSPIBMIB_QosIfRedTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies threshold limits for a particular RED
    // threshold of a given threshold set on an interface and with a particular
    // role combination. The type is slice of
    // CISCOQOSPIBMIB_QosIfRedTable_QosIfRedEntry.
    QosIfRedEntry []*CISCOQOSPIBMIB_QosIfRedTable_QosIfRedEntry
}

func (qosIfRedTable *CISCOQOSPIBMIB_QosIfRedTable) GetEntityData() *types.CommonEntityData {
    qosIfRedTable.EntityData.YFilter = qosIfRedTable.YFilter
    qosIfRedTable.EntityData.YangName = "qosIfRedTable"
    qosIfRedTable.EntityData.BundleName = "cisco_ios_xe"
    qosIfRedTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIfRedTable.EntityData.SegmentPath = "qosIfRedTable"
    qosIfRedTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIfRedTable.EntityData.SegmentPath
    qosIfRedTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfRedTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfRedTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfRedTable.EntityData.Children = types.NewOrderedMap()
    qosIfRedTable.EntityData.Children.Append("qosIfRedEntry", types.YChild{"QosIfRedEntry", nil})
    for i := range qosIfRedTable.QosIfRedEntry {
        qosIfRedTable.EntityData.Children.Append(types.GetSegmentPath(qosIfRedTable.QosIfRedEntry[i]), types.YChild{"QosIfRedEntry", qosIfRedTable.QosIfRedEntry[i]})
    }
    qosIfRedTable.EntityData.Leafs = types.NewOrderedMap()

    qosIfRedTable.EntityData.YListKeys = []string {}

    return &(qosIfRedTable.EntityData)
}

// CISCOQOSPIBMIB_QosIfRedTable_QosIfRedEntry
// An instance of this class specifies threshold limits for a
// particular RED threshold of a given threshold set on an
// interface and with a particular role combination.
type CISCOQOSPIBMIB_QosIfRedTable_QosIfRedEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIfRedId interface{}

    // The role combination the interface must be configured with. The type is
    // string with length: 0..255.
    QosIfRedRoles interface{}

    // The values in this entry apply only to queues with the number of thresholds
    // specified by this attribute. The type is ThresholdSetRange.
    QosIfRedNumThresholdSets interface{}

    // The threshold set to which the lower and upper values apply. It must be in
    // the range 1 through qosIfRedNumThresholdSets. There must be exactly one PRI
    // for each value in this range. The type is interface{} with range: 1..8.
    QosIfRedThresholdSet interface{}

    // The threshold value below which no packets are dropped. The type is
    // interface{} with range: 0..100.
    QosIfRedThresholdSetLower interface{}

    // The threshold value above which all packets are dropped. The type is
    // interface{} with range: 0..100.
    QosIfRedThresholdSetUpper interface{}
}

func (qosIfRedEntry *CISCOQOSPIBMIB_QosIfRedTable_QosIfRedEntry) GetEntityData() *types.CommonEntityData {
    qosIfRedEntry.EntityData.YFilter = qosIfRedEntry.YFilter
    qosIfRedEntry.EntityData.YangName = "qosIfRedEntry"
    qosIfRedEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIfRedEntry.EntityData.ParentYangName = "qosIfRedTable"
    qosIfRedEntry.EntityData.SegmentPath = "qosIfRedEntry" + types.AddKeyToken(qosIfRedEntry.QosIfRedId, "qosIfRedId")
    qosIfRedEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIfRedTable/" + qosIfRedEntry.EntityData.SegmentPath
    qosIfRedEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfRedEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfRedEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfRedEntry.EntityData.Children = types.NewOrderedMap()
    qosIfRedEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIfRedEntry.EntityData.Leafs.Append("qosIfRedId", types.YLeaf{"QosIfRedId", qosIfRedEntry.QosIfRedId})
    qosIfRedEntry.EntityData.Leafs.Append("qosIfRedRoles", types.YLeaf{"QosIfRedRoles", qosIfRedEntry.QosIfRedRoles})
    qosIfRedEntry.EntityData.Leafs.Append("qosIfRedNumThresholdSets", types.YLeaf{"QosIfRedNumThresholdSets", qosIfRedEntry.QosIfRedNumThresholdSets})
    qosIfRedEntry.EntityData.Leafs.Append("qosIfRedThresholdSet", types.YLeaf{"QosIfRedThresholdSet", qosIfRedEntry.QosIfRedThresholdSet})
    qosIfRedEntry.EntityData.Leafs.Append("qosIfRedThresholdSetLower", types.YLeaf{"QosIfRedThresholdSetLower", qosIfRedEntry.QosIfRedThresholdSetLower})
    qosIfRedEntry.EntityData.Leafs.Append("qosIfRedThresholdSetUpper", types.YLeaf{"QosIfRedThresholdSetUpper", qosIfRedEntry.QosIfRedThresholdSetUpper})

    qosIfRedEntry.EntityData.YListKeys = []string {"QosIfRedId"}

    return &(qosIfRedEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIfTailDropTable
// A class for threshold sets in a queue supporting tail drop.
// If the size of the queue for a given threshold set is at or
// below the specified value then packets assigned to that
// threshold set are always accepted into the queue.  If the size
// of the queue is above the specified value then packets are
// always dropped.
type CISCOQOSPIBMIB_QosIfTailDropTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the queue depth for a particular
    // tail-drop threshold set on an interface with a particular role combination.
    // The type is slice of CISCOQOSPIBMIB_QosIfTailDropTable_QosIfTailDropEntry.
    QosIfTailDropEntry []*CISCOQOSPIBMIB_QosIfTailDropTable_QosIfTailDropEntry
}

func (qosIfTailDropTable *CISCOQOSPIBMIB_QosIfTailDropTable) GetEntityData() *types.CommonEntityData {
    qosIfTailDropTable.EntityData.YFilter = qosIfTailDropTable.YFilter
    qosIfTailDropTable.EntityData.YangName = "qosIfTailDropTable"
    qosIfTailDropTable.EntityData.BundleName = "cisco_ios_xe"
    qosIfTailDropTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIfTailDropTable.EntityData.SegmentPath = "qosIfTailDropTable"
    qosIfTailDropTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIfTailDropTable.EntityData.SegmentPath
    qosIfTailDropTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfTailDropTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfTailDropTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfTailDropTable.EntityData.Children = types.NewOrderedMap()
    qosIfTailDropTable.EntityData.Children.Append("qosIfTailDropEntry", types.YChild{"QosIfTailDropEntry", nil})
    for i := range qosIfTailDropTable.QosIfTailDropEntry {
        qosIfTailDropTable.EntityData.Children.Append(types.GetSegmentPath(qosIfTailDropTable.QosIfTailDropEntry[i]), types.YChild{"QosIfTailDropEntry", qosIfTailDropTable.QosIfTailDropEntry[i]})
    }
    qosIfTailDropTable.EntityData.Leafs = types.NewOrderedMap()

    qosIfTailDropTable.EntityData.YListKeys = []string {}

    return &(qosIfTailDropTable.EntityData)
}

// CISCOQOSPIBMIB_QosIfTailDropTable_QosIfTailDropEntry
// An instance of this class specifies the queue depth for a
// particular tail-drop threshold set on an interface with a
// particular role combination.
type CISCOQOSPIBMIB_QosIfTailDropTable_QosIfTailDropEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIfTailDropId interface{}

    // The role combination the interface must be configured with. The type is
    // string with length: 0..255.
    QosIfTailDropRoles interface{}

    // The value in this entry applies only to queues with the number of
    // thresholds specified by this attribute. The type is ThresholdSetRange.
    QosIfTailDropNumThresholdSets interface{}

    // The threshold set to which the threshold value applies. The type is
    // interface{} with range: 1..8.
    QosIfTailDropThresholdSet interface{}

    // The threshold value above which packets are dropped. The type is
    // interface{} with range: 0..100.
    QosIfTailDropThresholdSetValue interface{}
}

func (qosIfTailDropEntry *CISCOQOSPIBMIB_QosIfTailDropTable_QosIfTailDropEntry) GetEntityData() *types.CommonEntityData {
    qosIfTailDropEntry.EntityData.YFilter = qosIfTailDropEntry.YFilter
    qosIfTailDropEntry.EntityData.YangName = "qosIfTailDropEntry"
    qosIfTailDropEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIfTailDropEntry.EntityData.ParentYangName = "qosIfTailDropTable"
    qosIfTailDropEntry.EntityData.SegmentPath = "qosIfTailDropEntry" + types.AddKeyToken(qosIfTailDropEntry.QosIfTailDropId, "qosIfTailDropId")
    qosIfTailDropEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIfTailDropTable/" + qosIfTailDropEntry.EntityData.SegmentPath
    qosIfTailDropEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfTailDropEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfTailDropEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfTailDropEntry.EntityData.Children = types.NewOrderedMap()
    qosIfTailDropEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIfTailDropEntry.EntityData.Leafs.Append("qosIfTailDropId", types.YLeaf{"QosIfTailDropId", qosIfTailDropEntry.QosIfTailDropId})
    qosIfTailDropEntry.EntityData.Leafs.Append("qosIfTailDropRoles", types.YLeaf{"QosIfTailDropRoles", qosIfTailDropEntry.QosIfTailDropRoles})
    qosIfTailDropEntry.EntityData.Leafs.Append("qosIfTailDropNumThresholdSets", types.YLeaf{"QosIfTailDropNumThresholdSets", qosIfTailDropEntry.QosIfTailDropNumThresholdSets})
    qosIfTailDropEntry.EntityData.Leafs.Append("qosIfTailDropThresholdSet", types.YLeaf{"QosIfTailDropThresholdSet", qosIfTailDropEntry.QosIfTailDropThresholdSet})
    qosIfTailDropEntry.EntityData.Leafs.Append("qosIfTailDropThresholdSetValue", types.YLeaf{"QosIfTailDropThresholdSetValue", qosIfTailDropEntry.QosIfTailDropThresholdSetValue})

    qosIfTailDropEntry.EntityData.YListKeys = []string {"QosIfTailDropId"}

    return &(qosIfTailDropEntry.EntityData)
}

// CISCOQOSPIBMIB_QosIfWeightsTable
// A class of scheduling weights for each queue of an interface
// that supports weighted round robin scheduling or a mix of
// priority queueing and weighted round robin.  For a queue with
// N priority queues, the N highest queue numbers are the
// priority queues with the highest queue number having the
// highest priority.  WRR is applied to the non-priority queues.
type CISCOQOSPIBMIB_QosIfWeightsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An instance of this class specifies the scheduling weight for a particular
    // queue of an interface with a particular number of queues and with a
    // particular role combination. The type is slice of
    // CISCOQOSPIBMIB_QosIfWeightsTable_QosIfWeightsEntry.
    QosIfWeightsEntry []*CISCOQOSPIBMIB_QosIfWeightsTable_QosIfWeightsEntry
}

func (qosIfWeightsTable *CISCOQOSPIBMIB_QosIfWeightsTable) GetEntityData() *types.CommonEntityData {
    qosIfWeightsTable.EntityData.YFilter = qosIfWeightsTable.YFilter
    qosIfWeightsTable.EntityData.YangName = "qosIfWeightsTable"
    qosIfWeightsTable.EntityData.BundleName = "cisco_ios_xe"
    qosIfWeightsTable.EntityData.ParentYangName = "CISCO-QOS-PIB-MIB"
    qosIfWeightsTable.EntityData.SegmentPath = "qosIfWeightsTable"
    qosIfWeightsTable.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/" + qosIfWeightsTable.EntityData.SegmentPath
    qosIfWeightsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfWeightsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfWeightsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfWeightsTable.EntityData.Children = types.NewOrderedMap()
    qosIfWeightsTable.EntityData.Children.Append("qosIfWeightsEntry", types.YChild{"QosIfWeightsEntry", nil})
    for i := range qosIfWeightsTable.QosIfWeightsEntry {
        qosIfWeightsTable.EntityData.Children.Append(types.GetSegmentPath(qosIfWeightsTable.QosIfWeightsEntry[i]), types.YChild{"QosIfWeightsEntry", qosIfWeightsTable.QosIfWeightsEntry[i]})
    }
    qosIfWeightsTable.EntityData.Leafs = types.NewOrderedMap()

    qosIfWeightsTable.EntityData.YListKeys = []string {}

    return &(qosIfWeightsTable.EntityData)
}

// CISCOQOSPIBMIB_QosIfWeightsTable_QosIfWeightsEntry
// An instance of this class specifies the scheduling weight for
// a particular queue of an interface with a particular number
// of queues and with a particular role combination.
type CISCOQOSPIBMIB_QosIfWeightsTable_QosIfWeightsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    QosIfWeightsId interface{}

    // The role combination the interface must be configured with. The type is
    // string with length: 0..255.
    QosIfWeightsRoles interface{}

    // The value of the weight in this instance applies only to interfaces with
    // the number of queues specified by this attribute. The type is QueueRange.
    QosIfWeightsNumQueues interface{}

    // The queue to which the weight applies. The type is interface{} with range:
    // 1..64.
    QosIfWeightsQueue interface{}

    // The maximum number of bytes that may be drained from the queue in one
    // cycle.  The percentage of the bandwith allocated to this queue can be
    // calculated from this attribute and the sum of the drain sizes of all the
    // non-priority queues of the interface. The type is interface{} with range:
    // 0..4294967295.
    QosIfWeightsDrainSize interface{}

    // The size of the queue in bytes.  Some devices set queue size in terms of
    // packets.  These devices must calculate the queue size in packets by
    // assuming an average packet size suitable for the particular interface. 
    // Some devices have a fixed size buffer to be shared among all queues.  These
    // devices must allocate a fraction of the total buffer space to this queue
    // calculated as the the ratio of the queue size to the sum of the queue sizes
    // for the interface. The type is interface{} with range: 0..4294967295.
    QosIfWeightsQueueSize interface{}
}

func (qosIfWeightsEntry *CISCOQOSPIBMIB_QosIfWeightsTable_QosIfWeightsEntry) GetEntityData() *types.CommonEntityData {
    qosIfWeightsEntry.EntityData.YFilter = qosIfWeightsEntry.YFilter
    qosIfWeightsEntry.EntityData.YangName = "qosIfWeightsEntry"
    qosIfWeightsEntry.EntityData.BundleName = "cisco_ios_xe"
    qosIfWeightsEntry.EntityData.ParentYangName = "qosIfWeightsTable"
    qosIfWeightsEntry.EntityData.SegmentPath = "qosIfWeightsEntry" + types.AddKeyToken(qosIfWeightsEntry.QosIfWeightsId, "qosIfWeightsId")
    qosIfWeightsEntry.EntityData.AbsolutePath = "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB/qosIfWeightsTable/" + qosIfWeightsEntry.EntityData.SegmentPath
    qosIfWeightsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    qosIfWeightsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    qosIfWeightsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    qosIfWeightsEntry.EntityData.Children = types.NewOrderedMap()
    qosIfWeightsEntry.EntityData.Leafs = types.NewOrderedMap()
    qosIfWeightsEntry.EntityData.Leafs.Append("qosIfWeightsId", types.YLeaf{"QosIfWeightsId", qosIfWeightsEntry.QosIfWeightsId})
    qosIfWeightsEntry.EntityData.Leafs.Append("qosIfWeightsRoles", types.YLeaf{"QosIfWeightsRoles", qosIfWeightsEntry.QosIfWeightsRoles})
    qosIfWeightsEntry.EntityData.Leafs.Append("qosIfWeightsNumQueues", types.YLeaf{"QosIfWeightsNumQueues", qosIfWeightsEntry.QosIfWeightsNumQueues})
    qosIfWeightsEntry.EntityData.Leafs.Append("qosIfWeightsQueue", types.YLeaf{"QosIfWeightsQueue", qosIfWeightsEntry.QosIfWeightsQueue})
    qosIfWeightsEntry.EntityData.Leafs.Append("qosIfWeightsDrainSize", types.YLeaf{"QosIfWeightsDrainSize", qosIfWeightsEntry.QosIfWeightsDrainSize})
    qosIfWeightsEntry.EntityData.Leafs.Append("qosIfWeightsQueueSize", types.YLeaf{"QosIfWeightsQueueSize", qosIfWeightsEntry.QosIfWeightsQueueSize})

    qosIfWeightsEntry.EntityData.YListKeys = []string {"QosIfWeightsId"}

    return &(qosIfWeightsEntry.EntityData)
}

