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

// ThresholdSetRange represents parameters packets are randomly dropped.
type ThresholdSetRange string

const (
    ThresholdSetRange_zeroT ThresholdSetRange = "zeroT"

    ThresholdSetRange_oneT ThresholdSetRange = "oneT"

    ThresholdSetRange_twoT ThresholdSetRange = "twoT"

    ThresholdSetRange_fourT ThresholdSetRange = "fourT"

    ThresholdSetRange_eightT ThresholdSetRange = "eightT"
)

// CISCOQOSPIBMIB
type CISCOQOSPIBMIB struct {
    parent types.Entity
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

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetFilter() yfilter.YFilter { return cISCOQOSPIBMIB.YFilter }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) SetFilter(yf yfilter.YFilter) { cISCOQOSPIBMIB.YFilter = yf }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetGoName(yname string) string {
    if yname == "qosDevicePibIncarnationTable" { return "Qosdevicepibincarnationtable" }
    if yname == "qosDeviceAttributeTable" { return "Qosdeviceattributetable" }
    if yname == "qosInterfaceTypeTable" { return "Qosinterfacetypetable" }
    if yname == "qosDiffServMappingTable" { return "Qosdiffservmappingtable" }
    if yname == "qosCosToDscpTable" { return "Qoscostodscptable" }
    if yname == "qosUnmatchedPolicyTable" { return "Qosunmatchedpolicytable" }
    if yname == "qosPolicerTable" { return "Qospolicertable" }
    if yname == "qosAggregateTable" { return "Qosaggregatetable" }
    if yname == "qosMacClassificationTable" { return "Qosmacclassificationtable" }
    if yname == "qosIpAceTable" { return "Qosipacetable" }
    if yname == "qosIpAclDefinitionTable" { return "Qosipacldefinitiontable" }
    if yname == "qosIpAclActionTable" { return "Qosipaclactiontable" }
    if yname == "qosIfSchedulingPreferencesTable" { return "Qosifschedulingpreferencestable" }
    if yname == "qosIfDropPreferenceTable" { return "Qosifdroppreferencetable" }
    if yname == "qosIfDscpAssignmentTable" { return "Qosifdscpassignmenttable" }
    if yname == "qosIfRedTable" { return "Qosifredtable" }
    if yname == "qosIfTailDropTable" { return "Qosiftaildroptable" }
    if yname == "qosIfWeightsTable" { return "Qosifweightstable" }
    return ""
}

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetSegmentPath() string {
    return "CISCO-QOS-PIB-MIB:CISCO-QOS-PIB-MIB"
}

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosDevicePibIncarnationTable" {
        return &cISCOQOSPIBMIB.Qosdevicepibincarnationtable
    }
    if childYangName == "qosDeviceAttributeTable" {
        return &cISCOQOSPIBMIB.Qosdeviceattributetable
    }
    if childYangName == "qosInterfaceTypeTable" {
        return &cISCOQOSPIBMIB.Qosinterfacetypetable
    }
    if childYangName == "qosDiffServMappingTable" {
        return &cISCOQOSPIBMIB.Qosdiffservmappingtable
    }
    if childYangName == "qosCosToDscpTable" {
        return &cISCOQOSPIBMIB.Qoscostodscptable
    }
    if childYangName == "qosUnmatchedPolicyTable" {
        return &cISCOQOSPIBMIB.Qosunmatchedpolicytable
    }
    if childYangName == "qosPolicerTable" {
        return &cISCOQOSPIBMIB.Qospolicertable
    }
    if childYangName == "qosAggregateTable" {
        return &cISCOQOSPIBMIB.Qosaggregatetable
    }
    if childYangName == "qosMacClassificationTable" {
        return &cISCOQOSPIBMIB.Qosmacclassificationtable
    }
    if childYangName == "qosIpAceTable" {
        return &cISCOQOSPIBMIB.Qosipacetable
    }
    if childYangName == "qosIpAclDefinitionTable" {
        return &cISCOQOSPIBMIB.Qosipacldefinitiontable
    }
    if childYangName == "qosIpAclActionTable" {
        return &cISCOQOSPIBMIB.Qosipaclactiontable
    }
    if childYangName == "qosIfSchedulingPreferencesTable" {
        return &cISCOQOSPIBMIB.Qosifschedulingpreferencestable
    }
    if childYangName == "qosIfDropPreferenceTable" {
        return &cISCOQOSPIBMIB.Qosifdroppreferencetable
    }
    if childYangName == "qosIfDscpAssignmentTable" {
        return &cISCOQOSPIBMIB.Qosifdscpassignmenttable
    }
    if childYangName == "qosIfRedTable" {
        return &cISCOQOSPIBMIB.Qosifredtable
    }
    if childYangName == "qosIfTailDropTable" {
        return &cISCOQOSPIBMIB.Qosiftaildroptable
    }
    if childYangName == "qosIfWeightsTable" {
        return &cISCOQOSPIBMIB.Qosifweightstable
    }
    return nil
}

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["qosDevicePibIncarnationTable"] = &cISCOQOSPIBMIB.Qosdevicepibincarnationtable
    children["qosDeviceAttributeTable"] = &cISCOQOSPIBMIB.Qosdeviceattributetable
    children["qosInterfaceTypeTable"] = &cISCOQOSPIBMIB.Qosinterfacetypetable
    children["qosDiffServMappingTable"] = &cISCOQOSPIBMIB.Qosdiffservmappingtable
    children["qosCosToDscpTable"] = &cISCOQOSPIBMIB.Qoscostodscptable
    children["qosUnmatchedPolicyTable"] = &cISCOQOSPIBMIB.Qosunmatchedpolicytable
    children["qosPolicerTable"] = &cISCOQOSPIBMIB.Qospolicertable
    children["qosAggregateTable"] = &cISCOQOSPIBMIB.Qosaggregatetable
    children["qosMacClassificationTable"] = &cISCOQOSPIBMIB.Qosmacclassificationtable
    children["qosIpAceTable"] = &cISCOQOSPIBMIB.Qosipacetable
    children["qosIpAclDefinitionTable"] = &cISCOQOSPIBMIB.Qosipacldefinitiontable
    children["qosIpAclActionTable"] = &cISCOQOSPIBMIB.Qosipaclactiontable
    children["qosIfSchedulingPreferencesTable"] = &cISCOQOSPIBMIB.Qosifschedulingpreferencestable
    children["qosIfDropPreferenceTable"] = &cISCOQOSPIBMIB.Qosifdroppreferencetable
    children["qosIfDscpAssignmentTable"] = &cISCOQOSPIBMIB.Qosifdscpassignmenttable
    children["qosIfRedTable"] = &cISCOQOSPIBMIB.Qosifredtable
    children["qosIfTailDropTable"] = &cISCOQOSPIBMIB.Qosiftaildroptable
    children["qosIfWeightsTable"] = &cISCOQOSPIBMIB.Qosifweightstable
    return children
}

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetYangName() string { return "CISCO-QOS-PIB-MIB" }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) SetParent(parent types.Entity) { cISCOQOSPIBMIB.parent = parent }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetParent() types.Entity { return cISCOQOSPIBMIB.parent }

func (cISCOQOSPIBMIB *CISCOQOSPIBMIB) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosdevicepibincarnationtable
// This class contains a single policy instance that identifies
// the current incarnation of the PIB and the PDP that installed
// this incarnation.  The instance of this class is reported to
// the PDP at client connect time so that the PDP can (attempt
// to) ascertain the current state of the PIB.
type CISCOQOSPIBMIB_Qosdevicepibincarnationtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The single policy instance of this class identifies the current incarnation
    // of the PIB and the PDP that installed this incarnation. The type is slice
    // of
    // CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry.
    Qosdevicepibincarnationentry []CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry
}

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetFilter() yfilter.YFilter { return qosdevicepibincarnationtable.YFilter }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) SetFilter(yf yfilter.YFilter) { qosdevicepibincarnationtable.YFilter = yf }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetGoName(yname string) string {
    if yname == "qosDevicePibIncarnationEntry" { return "Qosdevicepibincarnationentry" }
    return ""
}

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetSegmentPath() string {
    return "qosDevicePibIncarnationTable"
}

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosDevicePibIncarnationEntry" {
        for _, c := range qosdevicepibincarnationtable.Qosdevicepibincarnationentry {
            if qosdevicepibincarnationtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry{}
        qosdevicepibincarnationtable.Qosdevicepibincarnationentry = append(qosdevicepibincarnationtable.Qosdevicepibincarnationentry, child)
        return &qosdevicepibincarnationtable.Qosdevicepibincarnationentry[len(qosdevicepibincarnationtable.Qosdevicepibincarnationentry)-1]
    }
    return nil
}

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosdevicepibincarnationtable.Qosdevicepibincarnationentry {
        children[qosdevicepibincarnationtable.Qosdevicepibincarnationentry[i].GetSegmentPath()] = &qosdevicepibincarnationtable.Qosdevicepibincarnationentry[i]
    }
    return children
}

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetBundleName() string { return "cisco_ios_xe" }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetYangName() string { return "qosDevicePibIncarnationTable" }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) SetParent(parent types.Entity) { qosdevicepibincarnationtable.parent = parent }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetParent() types.Entity { return qosdevicepibincarnationtable.parent }

func (qosdevicepibincarnationtable *CISCOQOSPIBMIB_Qosdevicepibincarnationtable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry
// The single policy instance of this class identifies the
// current incarnation of the PIB and the PDP that installed
// this incarnation.
type CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry struct {
    parent types.Entity
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

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetFilter() yfilter.YFilter { return qosdevicepibincarnationentry.YFilter }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) SetFilter(yf yfilter.YFilter) { qosdevicepibincarnationentry.YFilter = yf }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetGoName(yname string) string {
    if yname == "qosDeviceIncarnationId" { return "Qosdeviceincarnationid" }
    if yname == "qosDevicePdpName" { return "Qosdevicepdpname" }
    if yname == "qosDevicePibIncarnation" { return "Qosdevicepibincarnation" }
    if yname == "qosDevicePibTtl" { return "Qosdevicepibttl" }
    return ""
}

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetSegmentPath() string {
    return "qosDevicePibIncarnationEntry" + "[qosDeviceIncarnationId='" + fmt.Sprintf("%v", qosdevicepibincarnationentry.Qosdeviceincarnationid) + "']"
}

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosDeviceIncarnationId"] = qosdevicepibincarnationentry.Qosdeviceincarnationid
    leafs["qosDevicePdpName"] = qosdevicepibincarnationentry.Qosdevicepdpname
    leafs["qosDevicePibIncarnation"] = qosdevicepibincarnationentry.Qosdevicepibincarnation
    leafs["qosDevicePibTtl"] = qosdevicepibincarnationentry.Qosdevicepibttl
    return leafs
}

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetYangName() string { return "qosDevicePibIncarnationEntry" }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) SetParent(parent types.Entity) { qosdevicepibincarnationentry.parent = parent }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetParent() types.Entity { return qosdevicepibincarnationentry.parent }

func (qosdevicepibincarnationentry *CISCOQOSPIBMIB_Qosdevicepibincarnationtable_Qosdevicepibincarnationentry) GetParentYangName() string { return "qosDevicePibIncarnationTable" }

// CISCOQOSPIBMIB_Qosdeviceattributetable
// The single instance of this class indicates specific
// attributes of the device.  These include configuration values
// such as the configured PDP addresses, the maximum message
// size, and specific device capabilities.  The latter include
// input port-based and output port-based classification and/or
// policing, support for flow based policing, aggregate based
// policing, traffic shaping capabilities, etc.
type CISCOQOSPIBMIB_Qosdeviceattributetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The single instance of this class indicates specific attributes of the
    // device. The type is slice of
    // CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry.
    Qosdeviceattributeentry []CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry
}

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetFilter() yfilter.YFilter { return qosdeviceattributetable.YFilter }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) SetFilter(yf yfilter.YFilter) { qosdeviceattributetable.YFilter = yf }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetGoName(yname string) string {
    if yname == "qosDeviceAttributeEntry" { return "Qosdeviceattributeentry" }
    return ""
}

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetSegmentPath() string {
    return "qosDeviceAttributeTable"
}

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosDeviceAttributeEntry" {
        for _, c := range qosdeviceattributetable.Qosdeviceattributeentry {
            if qosdeviceattributetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry{}
        qosdeviceattributetable.Qosdeviceattributeentry = append(qosdeviceattributetable.Qosdeviceattributeentry, child)
        return &qosdeviceattributetable.Qosdeviceattributeentry[len(qosdeviceattributetable.Qosdeviceattributeentry)-1]
    }
    return nil
}

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosdeviceattributetable.Qosdeviceattributeentry {
        children[qosdeviceattributetable.Qosdeviceattributeentry[i].GetSegmentPath()] = &qosdeviceattributetable.Qosdeviceattributeentry[i]
    }
    return children
}

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetBundleName() string { return "cisco_ios_xe" }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetYangName() string { return "qosDeviceAttributeTable" }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) SetParent(parent types.Entity) { qosdeviceattributetable.parent = parent }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetParent() types.Entity { return qosdeviceattributetable.parent }

func (qosdeviceattributetable *CISCOQOSPIBMIB_Qosdeviceattributetable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry
// The single instance of this class indicates specific
// attributes of the device.
type CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Qosdeviceprimarypdp interface{}

    // The address of the PDP configured to be the secondary PDP for the device. 
    // An address of zero indicates no secondary is configured. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetFilter() yfilter.YFilter { return qosdeviceattributeentry.YFilter }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) SetFilter(yf yfilter.YFilter) { qosdeviceattributeentry.YFilter = yf }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetGoName(yname string) string {
    if yname == "qosDeviceAttributeId" { return "Qosdeviceattributeid" }
    if yname == "qosDevicePepDomain" { return "Qosdevicepepdomain" }
    if yname == "qosDevicePrimaryPdp" { return "Qosdeviceprimarypdp" }
    if yname == "qosDeviceSecondaryPdp" { return "Qosdevicesecondarypdp" }
    if yname == "qosDeviceMaxMessageSize" { return "Qosdevicemaxmessagesize" }
    if yname == "qosDeviceCapabilities" { return "Qosdevicecapabilities" }
    return ""
}

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetSegmentPath() string {
    return "qosDeviceAttributeEntry" + "[qosDeviceAttributeId='" + fmt.Sprintf("%v", qosdeviceattributeentry.Qosdeviceattributeid) + "']"
}

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosDeviceAttributeId"] = qosdeviceattributeentry.Qosdeviceattributeid
    leafs["qosDevicePepDomain"] = qosdeviceattributeentry.Qosdevicepepdomain
    leafs["qosDevicePrimaryPdp"] = qosdeviceattributeentry.Qosdeviceprimarypdp
    leafs["qosDeviceSecondaryPdp"] = qosdeviceattributeentry.Qosdevicesecondarypdp
    leafs["qosDeviceMaxMessageSize"] = qosdeviceattributeentry.Qosdevicemaxmessagesize
    leafs["qosDeviceCapabilities"] = qosdeviceattributeentry.Qosdevicecapabilities
    return leafs
}

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetYangName() string { return "qosDeviceAttributeEntry" }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) SetParent(parent types.Entity) { qosdeviceattributeentry.parent = parent }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetParent() types.Entity { return qosdeviceattributeentry.parent }

func (qosdeviceattributeentry *CISCOQOSPIBMIB_Qosdeviceattributetable_Qosdeviceattributeentry) GetParentYangName() string { return "qosDeviceAttributeTable" }

// CISCOQOSPIBMIB_Qosinterfacetypetable
// This class describes the interface types of the interfaces
// that exist on the device.  It includes the queue type, role
// combination and capabilities of interfaces.  The PEP does not
// report which specific interfaces have which characteristics.
type CISCOQOSPIBMIB_Qosinterfacetypetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class describes a role combination for an interface
    // type of an interface that exists on the device. The type is slice of
    // CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry.
    Qosinterfacetypeentry []CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry
}

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetFilter() yfilter.YFilter { return qosinterfacetypetable.YFilter }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) SetFilter(yf yfilter.YFilter) { qosinterfacetypetable.YFilter = yf }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetGoName(yname string) string {
    if yname == "qosInterfaceTypeEntry" { return "Qosinterfacetypeentry" }
    return ""
}

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetSegmentPath() string {
    return "qosInterfaceTypeTable"
}

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosInterfaceTypeEntry" {
        for _, c := range qosinterfacetypetable.Qosinterfacetypeentry {
            if qosinterfacetypetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry{}
        qosinterfacetypetable.Qosinterfacetypeentry = append(qosinterfacetypetable.Qosinterfacetypeentry, child)
        return &qosinterfacetypetable.Qosinterfacetypeentry[len(qosinterfacetypetable.Qosinterfacetypeentry)-1]
    }
    return nil
}

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosinterfacetypetable.Qosinterfacetypeentry {
        children[qosinterfacetypetable.Qosinterfacetypeentry[i].GetSegmentPath()] = &qosinterfacetypetable.Qosinterfacetypeentry[i]
    }
    return children
}

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetBundleName() string { return "cisco_ios_xe" }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetYangName() string { return "qosInterfaceTypeTable" }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) SetParent(parent types.Entity) { qosinterfacetypetable.parent = parent }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetParent() types.Entity { return qosinterfacetypetable.parent }

func (qosinterfacetypetable *CISCOQOSPIBMIB_Qosinterfacetypetable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry
// An instance of this class describes a role combination for
// an interface type of an interface that exists on the device.
type CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry struct {
    parent types.Entity
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

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetFilter() yfilter.YFilter { return qosinterfacetypeentry.YFilter }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) SetFilter(yf yfilter.YFilter) { qosinterfacetypeentry.YFilter = yf }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetGoName(yname string) string {
    if yname == "qosInterfaceTypeId" { return "Qosinterfacetypeid" }
    if yname == "qosInterfaceQueueType" { return "Qosinterfacequeuetype" }
    if yname == "qosInterfaceTypeRoles" { return "Qosinterfacetyperoles" }
    if yname == "qosInterfaceTypeCapabilities" { return "Qosinterfacetypecapabilities" }
    return ""
}

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetSegmentPath() string {
    return "qosInterfaceTypeEntry" + "[qosInterfaceTypeId='" + fmt.Sprintf("%v", qosinterfacetypeentry.Qosinterfacetypeid) + "']"
}

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosInterfaceTypeId"] = qosinterfacetypeentry.Qosinterfacetypeid
    leafs["qosInterfaceQueueType"] = qosinterfacetypeentry.Qosinterfacequeuetype
    leafs["qosInterfaceTypeRoles"] = qosinterfacetypeentry.Qosinterfacetyperoles
    leafs["qosInterfaceTypeCapabilities"] = qosinterfacetypeentry.Qosinterfacetypecapabilities
    return leafs
}

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetYangName() string { return "qosInterfaceTypeEntry" }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) SetParent(parent types.Entity) { qosinterfacetypeentry.parent = parent }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetParent() types.Entity { return qosinterfacetypeentry.parent }

func (qosinterfacetypeentry *CISCOQOSPIBMIB_Qosinterfacetypetable_Qosinterfacetypeentry) GetParentYangName() string { return "qosInterfaceTypeTable" }

// CISCOQOSPIBMIB_Qosdiffservmappingtable
// Maps each DSCP to a marked-down DSCP.  Also maps each DSCP to
// an IP precedence and QosLayer2Cos.  When configured for the
// first time, all 64 entries of the table must be
// specified. Thereafter, instances may be modified (with a
// delete and install in a single decision) but not deleted
// unless all instances are deleted.
type CISCOQOSPIBMIB_Qosdiffservmappingtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class represents mappings from a DSCP. The type is
    // slice of CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry.
    Qosdiffservmappingentry []CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry
}

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetFilter() yfilter.YFilter { return qosdiffservmappingtable.YFilter }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) SetFilter(yf yfilter.YFilter) { qosdiffservmappingtable.YFilter = yf }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetGoName(yname string) string {
    if yname == "qosDiffServMappingEntry" { return "Qosdiffservmappingentry" }
    return ""
}

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetSegmentPath() string {
    return "qosDiffServMappingTable"
}

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosDiffServMappingEntry" {
        for _, c := range qosdiffservmappingtable.Qosdiffservmappingentry {
            if qosdiffservmappingtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry{}
        qosdiffservmappingtable.Qosdiffservmappingentry = append(qosdiffservmappingtable.Qosdiffservmappingentry, child)
        return &qosdiffservmappingtable.Qosdiffservmappingentry[len(qosdiffservmappingtable.Qosdiffservmappingentry)-1]
    }
    return nil
}

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosdiffservmappingtable.Qosdiffservmappingentry {
        children[qosdiffservmappingtable.Qosdiffservmappingentry[i].GetSegmentPath()] = &qosdiffservmappingtable.Qosdiffservmappingentry[i]
    }
    return children
}

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetBundleName() string { return "cisco_ios_xe" }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetYangName() string { return "qosDiffServMappingTable" }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) SetParent(parent types.Entity) { qosdiffservmappingtable.parent = parent }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetParent() types.Entity { return qosdiffservmappingtable.parent }

func (qosdiffservmappingtable *CISCOQOSPIBMIB_Qosdiffservmappingtable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry
// An instance of this class represents mappings from a DSCP.
type CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry struct {
    parent types.Entity
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

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetFilter() yfilter.YFilter { return qosdiffservmappingentry.YFilter }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) SetFilter(yf yfilter.YFilter) { qosdiffservmappingentry.YFilter = yf }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetGoName(yname string) string {
    if yname == "qosDscp" { return "Qosdscp" }
    if yname == "qosMarkedDscp" { return "Qosmarkeddscp" }
    if yname == "qosL2Cos" { return "Qosl2Cos" }
    return ""
}

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetSegmentPath() string {
    return "qosDiffServMappingEntry" + "[qosDscp='" + fmt.Sprintf("%v", qosdiffservmappingentry.Qosdscp) + "']"
}

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosDscp"] = qosdiffservmappingentry.Qosdscp
    leafs["qosMarkedDscp"] = qosdiffservmappingentry.Qosmarkeddscp
    leafs["qosL2Cos"] = qosdiffservmappingentry.Qosl2Cos
    return leafs
}

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetYangName() string { return "qosDiffServMappingEntry" }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) SetParent(parent types.Entity) { qosdiffservmappingentry.parent = parent }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetParent() types.Entity { return qosdiffservmappingentry.parent }

func (qosdiffservmappingentry *CISCOQOSPIBMIB_Qosdiffservmappingtable_Qosdiffservmappingentry) GetParentYangName() string { return "qosDiffServMappingTable" }

// CISCOQOSPIBMIB_Qoscostodscptable
// Maps each of eight CoS values to a DSCP.  When configured for
// the first time, all 8 entries of the table must be
// specified. Thereafter, instances may be modified (with a
// delete and install in a single decision) but not deleted
// unless all instances are deleted.
type CISCOQOSPIBMIB_Qoscostodscptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class maps a CoS value to a DSCP. The type is slice of
    // CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry.
    Qoscostodscpentry []CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry
}

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetFilter() yfilter.YFilter { return qoscostodscptable.YFilter }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) SetFilter(yf yfilter.YFilter) { qoscostodscptable.YFilter = yf }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetGoName(yname string) string {
    if yname == "qosCosToDscpEntry" { return "Qoscostodscpentry" }
    return ""
}

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetSegmentPath() string {
    return "qosCosToDscpTable"
}

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosCosToDscpEntry" {
        for _, c := range qoscostodscptable.Qoscostodscpentry {
            if qoscostodscptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry{}
        qoscostodscptable.Qoscostodscpentry = append(qoscostodscptable.Qoscostodscpentry, child)
        return &qoscostodscptable.Qoscostodscpentry[len(qoscostodscptable.Qoscostodscpentry)-1]
    }
    return nil
}

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qoscostodscptable.Qoscostodscpentry {
        children[qoscostodscptable.Qoscostodscpentry[i].GetSegmentPath()] = &qoscostodscptable.Qoscostodscpentry[i]
    }
    return children
}

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetBundleName() string { return "cisco_ios_xe" }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetYangName() string { return "qosCosToDscpTable" }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) SetParent(parent types.Entity) { qoscostodscptable.parent = parent }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetParent() types.Entity { return qoscostodscptable.parent }

func (qoscostodscptable *CISCOQOSPIBMIB_Qoscostodscptable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry
// An instance of this class maps a CoS value to a DSCP.
type CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The L2 CoS value that is being mapped. The type is
    // interface{} with range: 0..7.
    Qoscostodscpcos interface{}

    // The DSCP value to use when mapping the L2 CoS to a DSCP. The type is
    // interface{} with range: 0..63.
    Qoscostodscpdscp interface{}
}

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetFilter() yfilter.YFilter { return qoscostodscpentry.YFilter }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) SetFilter(yf yfilter.YFilter) { qoscostodscpentry.YFilter = yf }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetGoName(yname string) string {
    if yname == "qosCosToDscpCos" { return "Qoscostodscpcos" }
    if yname == "qosCosToDscpDscp" { return "Qoscostodscpdscp" }
    return ""
}

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetSegmentPath() string {
    return "qosCosToDscpEntry" + "[qosCosToDscpCos='" + fmt.Sprintf("%v", qoscostodscpentry.Qoscostodscpcos) + "']"
}

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosCosToDscpCos"] = qoscostodscpentry.Qoscostodscpcos
    leafs["qosCosToDscpDscp"] = qoscostodscpentry.Qoscostodscpdscp
    return leafs
}

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetBundleName() string { return "cisco_ios_xe" }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetYangName() string { return "qosCosToDscpEntry" }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) SetParent(parent types.Entity) { qoscostodscpentry.parent = parent }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetParent() types.Entity { return qoscostodscpentry.parent }

func (qoscostodscpentry *CISCOQOSPIBMIB_Qoscostodscptable_Qoscostodscpentry) GetParentYangName() string { return "qosCosToDscpTable" }

// CISCOQOSPIBMIB_Qosunmatchedpolicytable
// A policy class that specifies what QoS to apply to a packet
// that does not match any other policy configured for this role
// combination for a particular direction of traffic.
type CISCOQOSPIBMIB_Qosunmatchedpolicytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies the unmatched policy for a particular
    // role combination for incoming or outgoing traffic. The type is slice of
    // CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry.
    Qosunmatchedpolicyentry []CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry
}

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetFilter() yfilter.YFilter { return qosunmatchedpolicytable.YFilter }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) SetFilter(yf yfilter.YFilter) { qosunmatchedpolicytable.YFilter = yf }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetGoName(yname string) string {
    if yname == "qosUnmatchedPolicyEntry" { return "Qosunmatchedpolicyentry" }
    return ""
}

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetSegmentPath() string {
    return "qosUnmatchedPolicyTable"
}

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosUnmatchedPolicyEntry" {
        for _, c := range qosunmatchedpolicytable.Qosunmatchedpolicyentry {
            if qosunmatchedpolicytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry{}
        qosunmatchedpolicytable.Qosunmatchedpolicyentry = append(qosunmatchedpolicytable.Qosunmatchedpolicyentry, child)
        return &qosunmatchedpolicytable.Qosunmatchedpolicyentry[len(qosunmatchedpolicytable.Qosunmatchedpolicyentry)-1]
    }
    return nil
}

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosunmatchedpolicytable.Qosunmatchedpolicyentry {
        children[qosunmatchedpolicytable.Qosunmatchedpolicyentry[i].GetSegmentPath()] = &qosunmatchedpolicytable.Qosunmatchedpolicyentry[i]
    }
    return children
}

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetBundleName() string { return "cisco_ios_xe" }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetYangName() string { return "qosUnmatchedPolicyTable" }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) SetParent(parent types.Entity) { qosunmatchedpolicytable.parent = parent }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetParent() types.Entity { return qosunmatchedpolicytable.parent }

func (qosunmatchedpolicytable *CISCOQOSPIBMIB_Qosunmatchedpolicytable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry
// An instance of this class specifies the unmatched policy
// for a particular role combination for incoming or outgoing
// traffic.
type CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry struct {
    parent types.Entity
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

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetFilter() yfilter.YFilter { return qosunmatchedpolicyentry.YFilter }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) SetFilter(yf yfilter.YFilter) { qosunmatchedpolicyentry.YFilter = yf }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetGoName(yname string) string {
    if yname == "qosUnmatchedPolicyId" { return "Qosunmatchedpolicyid" }
    if yname == "qosUnmatchedPolicyRole" { return "Qosunmatchedpolicyrole" }
    if yname == "qosUnmatchedPolicyDirection" { return "Qosunmatchedpolicydirection" }
    if yname == "qosUnmatchedPolicyDscp" { return "Qosunmatchedpolicydscp" }
    if yname == "qosUnmatchedPolicyDscpTrusted" { return "Qosunmatchedpolicydscptrusted" }
    if yname == "qosUnmatchPolMicroFlowPolicerId" { return "Qosunmatchpolmicroflowpolicerid" }
    if yname == "qosUnmatchedPolicyAggregateId" { return "Qosunmatchedpolicyaggregateid" }
    return ""
}

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetSegmentPath() string {
    return "qosUnmatchedPolicyEntry" + "[qosUnmatchedPolicyId='" + fmt.Sprintf("%v", qosunmatchedpolicyentry.Qosunmatchedpolicyid) + "']"
}

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosUnmatchedPolicyId"] = qosunmatchedpolicyentry.Qosunmatchedpolicyid
    leafs["qosUnmatchedPolicyRole"] = qosunmatchedpolicyentry.Qosunmatchedpolicyrole
    leafs["qosUnmatchedPolicyDirection"] = qosunmatchedpolicyentry.Qosunmatchedpolicydirection
    leafs["qosUnmatchedPolicyDscp"] = qosunmatchedpolicyentry.Qosunmatchedpolicydscp
    leafs["qosUnmatchedPolicyDscpTrusted"] = qosunmatchedpolicyentry.Qosunmatchedpolicydscptrusted
    leafs["qosUnmatchPolMicroFlowPolicerId"] = qosunmatchedpolicyentry.Qosunmatchpolmicroflowpolicerid
    leafs["qosUnmatchedPolicyAggregateId"] = qosunmatchedpolicyentry.Qosunmatchedpolicyaggregateid
    return leafs
}

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetYangName() string { return "qosUnmatchedPolicyEntry" }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) SetParent(parent types.Entity) { qosunmatchedpolicyentry.parent = parent }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetParent() types.Entity { return qosunmatchedpolicyentry.parent }

func (qosunmatchedpolicyentry *CISCOQOSPIBMIB_Qosunmatchedpolicytable_Qosunmatchedpolicyentry) GetParentYangName() string { return "qosUnmatchedPolicyTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies a set of policing parameters. The type
    // is slice of CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry.
    Qospolicerentry []CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry
}

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetFilter() yfilter.YFilter { return qospolicertable.YFilter }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) SetFilter(yf yfilter.YFilter) { qospolicertable.YFilter = yf }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetGoName(yname string) string {
    if yname == "qosPolicerEntry" { return "Qospolicerentry" }
    return ""
}

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetSegmentPath() string {
    return "qosPolicerTable"
}

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosPolicerEntry" {
        for _, c := range qospolicertable.Qospolicerentry {
            if qospolicertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry{}
        qospolicertable.Qospolicerentry = append(qospolicertable.Qospolicerentry, child)
        return &qospolicertable.Qospolicerentry[len(qospolicertable.Qospolicerentry)-1]
    }
    return nil
}

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qospolicertable.Qospolicerentry {
        children[qospolicertable.Qospolicerentry[i].GetSegmentPath()] = &qospolicertable.Qospolicerentry[i]
    }
    return children
}

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetBundleName() string { return "cisco_ios_xe" }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetYangName() string { return "qosPolicerTable" }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) SetParent(parent types.Entity) { qospolicertable.parent = parent }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetParent() types.Entity { return qospolicertable.parent }

func (qospolicertable *CISCOQOSPIBMIB_Qospolicertable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry
// An instance of this class specifies a set of policing
// parameters.
type CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry struct {
    parent types.Entity
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

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetFilter() yfilter.YFilter { return qospolicerentry.YFilter }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) SetFilter(yf yfilter.YFilter) { qospolicerentry.YFilter = yf }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetGoName(yname string) string {
    if yname == "qosPolicerId" { return "Qospolicerid" }
    if yname == "qosPolicerRate" { return "Qospolicerrate" }
    if yname == "qosPolicerNormalBurst" { return "Qospolicernormalburst" }
    if yname == "qosPolicerExcessBurst" { return "Qospolicerexcessburst" }
    if yname == "qosPolicerAction" { return "Qospoliceraction" }
    return ""
}

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetSegmentPath() string {
    return "qosPolicerEntry" + "[qosPolicerId='" + fmt.Sprintf("%v", qospolicerentry.Qospolicerid) + "']"
}

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosPolicerId"] = qospolicerentry.Qospolicerid
    leafs["qosPolicerRate"] = qospolicerentry.Qospolicerrate
    leafs["qosPolicerNormalBurst"] = qospolicerentry.Qospolicernormalburst
    leafs["qosPolicerExcessBurst"] = qospolicerentry.Qospolicerexcessburst
    leafs["qosPolicerAction"] = qospolicerentry.Qospoliceraction
    return leafs
}

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetBundleName() string { return "cisco_ios_xe" }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetYangName() string { return "qosPolicerEntry" }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) SetParent(parent types.Entity) { qospolicerentry.parent = parent }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetParent() types.Entity { return qospolicerentry.parent }

func (qospolicerentry *CISCOQOSPIBMIB_Qospolicertable_Qospolicerentry) GetParentYangName() string { return "qosPolicerTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies the policer to apply to an aggregate
    // flow. The type is slice of
    // CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry.
    Qosaggregateentry []CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry
}

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetFilter() yfilter.YFilter { return qosaggregatetable.YFilter }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) SetFilter(yf yfilter.YFilter) { qosaggregatetable.YFilter = yf }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetGoName(yname string) string {
    if yname == "qosAggregateEntry" { return "Qosaggregateentry" }
    return ""
}

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetSegmentPath() string {
    return "qosAggregateTable"
}

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosAggregateEntry" {
        for _, c := range qosaggregatetable.Qosaggregateentry {
            if qosaggregatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry{}
        qosaggregatetable.Qosaggregateentry = append(qosaggregatetable.Qosaggregateentry, child)
        return &qosaggregatetable.Qosaggregateentry[len(qosaggregatetable.Qosaggregateentry)-1]
    }
    return nil
}

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosaggregatetable.Qosaggregateentry {
        children[qosaggregatetable.Qosaggregateentry[i].GetSegmentPath()] = &qosaggregatetable.Qosaggregateentry[i]
    }
    return children
}

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetBundleName() string { return "cisco_ios_xe" }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetYangName() string { return "qosAggregateTable" }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) SetParent(parent types.Entity) { qosaggregatetable.parent = parent }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetParent() types.Entity { return qosaggregatetable.parent }

func (qosaggregatetable *CISCOQOSPIBMIB_Qosaggregatetable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry
// An instance of this class specifies the policer to apply to
// an aggregate flow.
type CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosaggregateid interface{}

    // An index identifying the instance of policer to apply to the aggregate.  It
    // must correspond to the integer index of an instance of class
    // qosPolicerTable. The type is interface{} with range: 0..4294967295.
    Qosaggregatepolicerid interface{}
}

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetFilter() yfilter.YFilter { return qosaggregateentry.YFilter }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) SetFilter(yf yfilter.YFilter) { qosaggregateentry.YFilter = yf }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetGoName(yname string) string {
    if yname == "qosAggregateId" { return "Qosaggregateid" }
    if yname == "qosAggregatePolicerId" { return "Qosaggregatepolicerid" }
    return ""
}

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetSegmentPath() string {
    return "qosAggregateEntry" + "[qosAggregateId='" + fmt.Sprintf("%v", qosaggregateentry.Qosaggregateid) + "']"
}

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosAggregateId"] = qosaggregateentry.Qosaggregateid
    leafs["qosAggregatePolicerId"] = qosaggregateentry.Qosaggregatepolicerid
    return leafs
}

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetYangName() string { return "qosAggregateEntry" }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) SetParent(parent types.Entity) { qosaggregateentry.parent = parent }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetParent() types.Entity { return qosaggregateentry.parent }

func (qosaggregateentry *CISCOQOSPIBMIB_Qosaggregatetable_Qosaggregateentry) GetParentYangName() string { return "qosAggregateTable" }

// CISCOQOSPIBMIB_Qosmacclassificationtable
// A class of MAC/Vlan tuples and their associated CoS values.
type CISCOQOSPIBMIB_Qosmacclassificationtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies the mapping of a VLAN and a MAC address
    // to a CoS value. The type is slice of
    // CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry.
    Qosmacclassificationentry []CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry
}

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetFilter() yfilter.YFilter { return qosmacclassificationtable.YFilter }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) SetFilter(yf yfilter.YFilter) { qosmacclassificationtable.YFilter = yf }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetGoName(yname string) string {
    if yname == "qosMacClassificationEntry" { return "Qosmacclassificationentry" }
    return ""
}

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetSegmentPath() string {
    return "qosMacClassificationTable"
}

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosMacClassificationEntry" {
        for _, c := range qosmacclassificationtable.Qosmacclassificationentry {
            if qosmacclassificationtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry{}
        qosmacclassificationtable.Qosmacclassificationentry = append(qosmacclassificationtable.Qosmacclassificationentry, child)
        return &qosmacclassificationtable.Qosmacclassificationentry[len(qosmacclassificationtable.Qosmacclassificationentry)-1]
    }
    return nil
}

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosmacclassificationtable.Qosmacclassificationentry {
        children[qosmacclassificationtable.Qosmacclassificationentry[i].GetSegmentPath()] = &qosmacclassificationtable.Qosmacclassificationentry[i]
    }
    return children
}

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetBundleName() string { return "cisco_ios_xe" }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetYangName() string { return "qosMacClassificationTable" }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) SetParent(parent types.Entity) { qosmacclassificationtable.parent = parent }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetParent() types.Entity { return qosmacclassificationtable.parent }

func (qosmacclassificationtable *CISCOQOSPIBMIB_Qosmacclassificationtable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry
// An instance of this class specifies the mapping of a VLAN
// and a MAC address to a CoS value.
type CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosmacclassificationid interface{}

    // The VLAN of the destination MAC address of the L2 frame. The type is
    // interface{} with range: 1..4095.
    Qosdstmacvlan interface{}

    // The destination MAC address of the L2 frame. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Qosdstmacaddress interface{}

    // The CoS to assign the packet with the associated MAC/VLAN tuple.  Note that
    // this CoS is overridden by the policies to classify the frame at layer 3 if
    // there are any. The type is interface{} with range: 0..7.
    Qosdstmaccos interface{}
}

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetFilter() yfilter.YFilter { return qosmacclassificationentry.YFilter }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) SetFilter(yf yfilter.YFilter) { qosmacclassificationentry.YFilter = yf }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetGoName(yname string) string {
    if yname == "qosMacClassificationId" { return "Qosmacclassificationid" }
    if yname == "qosDstMacVlan" { return "Qosdstmacvlan" }
    if yname == "qosDstMacAddress" { return "Qosdstmacaddress" }
    if yname == "qosDstMacCos" { return "Qosdstmaccos" }
    return ""
}

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetSegmentPath() string {
    return "qosMacClassificationEntry" + "[qosMacClassificationId='" + fmt.Sprintf("%v", qosmacclassificationentry.Qosmacclassificationid) + "']"
}

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosMacClassificationId"] = qosmacclassificationentry.Qosmacclassificationid
    leafs["qosDstMacVlan"] = qosmacclassificationentry.Qosdstmacvlan
    leafs["qosDstMacAddress"] = qosmacclassificationentry.Qosdstmacaddress
    leafs["qosDstMacCos"] = qosmacclassificationentry.Qosdstmaccos
    return leafs
}

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetYangName() string { return "qosMacClassificationEntry" }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) SetParent(parent types.Entity) { qosmacclassificationentry.parent = parent }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetParent() types.Entity { return qosmacclassificationentry.parent }

func (qosmacclassificationentry *CISCOQOSPIBMIB_Qosmacclassificationtable_Qosmacclassificationentry) GetParentYangName() string { return "qosMacClassificationTable" }

// CISCOQOSPIBMIB_Qosipacetable
// ACE definitions.
type CISCOQOSPIBMIB_Qosipacetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies an ACE. The type is slice of
    // CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry.
    Qosipaceentry []CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry
}

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetFilter() yfilter.YFilter { return qosipacetable.YFilter }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) SetFilter(yf yfilter.YFilter) { qosipacetable.YFilter = yf }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetGoName(yname string) string {
    if yname == "qosIpAceEntry" { return "Qosipaceentry" }
    return ""
}

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetSegmentPath() string {
    return "qosIpAceTable"
}

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIpAceEntry" {
        for _, c := range qosipacetable.Qosipaceentry {
            if qosipacetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry{}
        qosipacetable.Qosipaceentry = append(qosipacetable.Qosipaceentry, child)
        return &qosipacetable.Qosipaceentry[len(qosipacetable.Qosipaceentry)-1]
    }
    return nil
}

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosipacetable.Qosipaceentry {
        children[qosipacetable.Qosipaceentry[i].GetSegmentPath()] = &qosipacetable.Qosipaceentry[i]
    }
    return children
}

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetBundleName() string { return "cisco_ios_xe" }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetYangName() string { return "qosIpAceTable" }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) SetParent(parent types.Entity) { qosipacetable.parent = parent }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetParent() types.Entity { return qosipacetable.parent }

func (qosipacetable *CISCOQOSPIBMIB_Qosipacetable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry
// An instance of this class specifies an ACE.
type CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An integer index to identify the instance of the
    // policy class. The type is interface{} with range: 0..4294967295.
    Qosipaceid interface{}

    // The IP address to match against the packet's destination IP address. The
    // type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Qosipacedstaddr interface{}

    // A mask for the matching of the destination IP address. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Qosipacedstaddrmask interface{}

    // The IP address to match against the packet's source IP address. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Qosipacesrcaddr interface{}

    // A mask for the matching of the source IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetFilter() yfilter.YFilter { return qosipaceentry.YFilter }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) SetFilter(yf yfilter.YFilter) { qosipaceentry.YFilter = yf }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetGoName(yname string) string {
    if yname == "qosIpAceId" { return "Qosipaceid" }
    if yname == "qosIpAceDstAddr" { return "Qosipacedstaddr" }
    if yname == "qosIpAceDstAddrMask" { return "Qosipacedstaddrmask" }
    if yname == "qosIpAceSrcAddr" { return "Qosipacesrcaddr" }
    if yname == "qosIpAceSrcAddrMask" { return "Qosipacesrcaddrmask" }
    if yname == "qosIpAceDscpMin" { return "Qosipacedscpmin" }
    if yname == "qosIpAceDscpMax" { return "Qosipacedscpmax" }
    if yname == "qosIpAceProtocol" { return "Qosipaceprotocol" }
    if yname == "qosIpAceDstL4PortMin" { return "Qosipacedstl4Portmin" }
    if yname == "qosIpAceDstL4PortMax" { return "Qosipacedstl4Portmax" }
    if yname == "qosIpAceSrcL4PortMin" { return "Qosipacesrcl4Portmin" }
    if yname == "qosIpAceSrcL4PortMax" { return "Qosipacesrcl4Portmax" }
    if yname == "qosIpAcePermit" { return "Qosipacepermit" }
    return ""
}

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetSegmentPath() string {
    return "qosIpAceEntry" + "[qosIpAceId='" + fmt.Sprintf("%v", qosipaceentry.Qosipaceid) + "']"
}

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIpAceId"] = qosipaceentry.Qosipaceid
    leafs["qosIpAceDstAddr"] = qosipaceentry.Qosipacedstaddr
    leafs["qosIpAceDstAddrMask"] = qosipaceentry.Qosipacedstaddrmask
    leafs["qosIpAceSrcAddr"] = qosipaceentry.Qosipacesrcaddr
    leafs["qosIpAceSrcAddrMask"] = qosipaceentry.Qosipacesrcaddrmask
    leafs["qosIpAceDscpMin"] = qosipaceentry.Qosipacedscpmin
    leafs["qosIpAceDscpMax"] = qosipaceentry.Qosipacedscpmax
    leafs["qosIpAceProtocol"] = qosipaceentry.Qosipaceprotocol
    leafs["qosIpAceDstL4PortMin"] = qosipaceentry.Qosipacedstl4Portmin
    leafs["qosIpAceDstL4PortMax"] = qosipaceentry.Qosipacedstl4Portmax
    leafs["qosIpAceSrcL4PortMin"] = qosipaceentry.Qosipacesrcl4Portmin
    leafs["qosIpAceSrcL4PortMax"] = qosipaceentry.Qosipacesrcl4Portmax
    leafs["qosIpAcePermit"] = qosipaceentry.Qosipacepermit
    return leafs
}

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetYangName() string { return "qosIpAceEntry" }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) SetParent(parent types.Entity) { qosipaceentry.parent = parent }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetParent() types.Entity { return qosipaceentry.parent }

func (qosipaceentry *CISCOQOSPIBMIB_Qosipacetable_Qosipaceentry) GetParentYangName() string { return "qosIpAceTable" }

// CISCOQOSPIBMIB_Qosipacldefinitiontable
// A class that defines a set of ACLs each being an ordered list
// of ACEs.
type CISCOQOSPIBMIB_Qosipacldefinitiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies an ACE in an ACL and its order with
    // respect to other ACEs in the same ACL. The type is slice of
    // CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry.
    Qosipacldefinitionentry []CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry
}

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetFilter() yfilter.YFilter { return qosipacldefinitiontable.YFilter }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) SetFilter(yf yfilter.YFilter) { qosipacldefinitiontable.YFilter = yf }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetGoName(yname string) string {
    if yname == "qosIpAclDefinitionEntry" { return "Qosipacldefinitionentry" }
    return ""
}

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetSegmentPath() string {
    return "qosIpAclDefinitionTable"
}

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIpAclDefinitionEntry" {
        for _, c := range qosipacldefinitiontable.Qosipacldefinitionentry {
            if qosipacldefinitiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry{}
        qosipacldefinitiontable.Qosipacldefinitionentry = append(qosipacldefinitiontable.Qosipacldefinitionentry, child)
        return &qosipacldefinitiontable.Qosipacldefinitionentry[len(qosipacldefinitiontable.Qosipacldefinitionentry)-1]
    }
    return nil
}

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosipacldefinitiontable.Qosipacldefinitionentry {
        children[qosipacldefinitiontable.Qosipacldefinitionentry[i].GetSegmentPath()] = &qosipacldefinitiontable.Qosipacldefinitionentry[i]
    }
    return children
}

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetBundleName() string { return "cisco_ios_xe" }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetYangName() string { return "qosIpAclDefinitionTable" }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) SetParent(parent types.Entity) { qosipacldefinitiontable.parent = parent }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetParent() types.Entity { return qosipacldefinitiontable.parent }

func (qosipacldefinitiontable *CISCOQOSPIBMIB_Qosipacldefinitiontable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry
// An instance of this class specifies an ACE in an ACL and its
// order with respect to other ACEs in the same ACL.
type CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry struct {
    parent types.Entity
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

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetFilter() yfilter.YFilter { return qosipacldefinitionentry.YFilter }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) SetFilter(yf yfilter.YFilter) { qosipacldefinitionentry.YFilter = yf }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetGoName(yname string) string {
    if yname == "qosIpAclDefinitionId" { return "Qosipacldefinitionid" }
    if yname == "qosIpAclId" { return "Qosipaclid" }
    if yname == "qosIpAceOrder" { return "Qosipaceorder" }
    if yname == "qosIpAclDefAceId" { return "Qosipacldefaceid" }
    return ""
}

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetSegmentPath() string {
    return "qosIpAclDefinitionEntry" + "[qosIpAclDefinitionId='" + fmt.Sprintf("%v", qosipacldefinitionentry.Qosipacldefinitionid) + "']"
}

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIpAclDefinitionId"] = qosipacldefinitionentry.Qosipacldefinitionid
    leafs["qosIpAclId"] = qosipacldefinitionentry.Qosipaclid
    leafs["qosIpAceOrder"] = qosipacldefinitionentry.Qosipaceorder
    leafs["qosIpAclDefAceId"] = qosipacldefinitionentry.Qosipacldefaceid
    return leafs
}

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetYangName() string { return "qosIpAclDefinitionEntry" }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) SetParent(parent types.Entity) { qosipacldefinitionentry.parent = parent }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetParent() types.Entity { return qosipacldefinitionentry.parent }

func (qosipacldefinitionentry *CISCOQOSPIBMIB_Qosipacldefinitiontable_Qosipacldefinitionentry) GetParentYangName() string { return "qosIpAclDefinitionTable" }

// CISCOQOSPIBMIB_Qosipaclactiontable
// A class that applies a set of ACLs to interfaces specifying,
// for each interface the order of the ACL with respect to other
// ACLs applied to the same interface and, for each ACL the
// action to take for a packet that matches a permit ACE in that
// ACL.  Interfaces are specified abstractly in terms of
// interface role combinations.
type CISCOQOSPIBMIB_Qosipaclactiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class applies an ACL to traffic in a particular
    // direction on an interface with a particular role combination, and specifies
    // the action for packets which match the ACL. The type is slice of
    // CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry.
    Qosipaclactionentry []CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry
}

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetFilter() yfilter.YFilter { return qosipaclactiontable.YFilter }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) SetFilter(yf yfilter.YFilter) { qosipaclactiontable.YFilter = yf }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetGoName(yname string) string {
    if yname == "qosIpAclActionEntry" { return "Qosipaclactionentry" }
    return ""
}

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetSegmentPath() string {
    return "qosIpAclActionTable"
}

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIpAclActionEntry" {
        for _, c := range qosipaclactiontable.Qosipaclactionentry {
            if qosipaclactiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry{}
        qosipaclactiontable.Qosipaclactionentry = append(qosipaclactiontable.Qosipaclactionentry, child)
        return &qosipaclactiontable.Qosipaclactionentry[len(qosipaclactiontable.Qosipaclactionentry)-1]
    }
    return nil
}

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosipaclactiontable.Qosipaclactionentry {
        children[qosipaclactiontable.Qosipaclactionentry[i].GetSegmentPath()] = &qosipaclactiontable.Qosipaclactionentry[i]
    }
    return children
}

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetBundleName() string { return "cisco_ios_xe" }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetYangName() string { return "qosIpAclActionTable" }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) SetParent(parent types.Entity) { qosipaclactiontable.parent = parent }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetParent() types.Entity { return qosipaclactiontable.parent }

func (qosipaclactiontable *CISCOQOSPIBMIB_Qosipaclactiontable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry
// An instance of this class applies an ACL to traffic in a
// particular direction on an interface with a particular role
// combination, and specifies the action for packets which match
// the ACL.
type CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry struct {
    parent types.Entity
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

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetFilter() yfilter.YFilter { return qosipaclactionentry.YFilter }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) SetFilter(yf yfilter.YFilter) { qosipaclactionentry.YFilter = yf }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetGoName(yname string) string {
    if yname == "qosIpAclActionId" { return "Qosipaclactionid" }
    if yname == "qosIpAclActAclId" { return "Qosipaclactaclid" }
    if yname == "qosIpAclInterfaceRoles" { return "Qosipaclinterfaceroles" }
    if yname == "qosIpAclInterfaceDirection" { return "Qosipaclinterfacedirection" }
    if yname == "qosIpAclOrder" { return "Qosipaclorder" }
    if yname == "qosIpAclDscp" { return "Qosipacldscp" }
    if yname == "qosIpAclDscpTrusted" { return "Qosipacldscptrusted" }
    if yname == "qosIpAclMicroFlowPolicerId" { return "Qosipaclmicroflowpolicerid" }
    if yname == "qosIpAclAggregateId" { return "Qosipaclaggregateid" }
    return ""
}

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetSegmentPath() string {
    return "qosIpAclActionEntry" + "[qosIpAclActionId='" + fmt.Sprintf("%v", qosipaclactionentry.Qosipaclactionid) + "']"
}

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIpAclActionId"] = qosipaclactionentry.Qosipaclactionid
    leafs["qosIpAclActAclId"] = qosipaclactionentry.Qosipaclactaclid
    leafs["qosIpAclInterfaceRoles"] = qosipaclactionentry.Qosipaclinterfaceroles
    leafs["qosIpAclInterfaceDirection"] = qosipaclactionentry.Qosipaclinterfacedirection
    leafs["qosIpAclOrder"] = qosipaclactionentry.Qosipaclorder
    leafs["qosIpAclDscp"] = qosipaclactionentry.Qosipacldscp
    leafs["qosIpAclDscpTrusted"] = qosipaclactionentry.Qosipacldscptrusted
    leafs["qosIpAclMicroFlowPolicerId"] = qosipaclactionentry.Qosipaclmicroflowpolicerid
    leafs["qosIpAclAggregateId"] = qosipaclactionentry.Qosipaclaggregateid
    return leafs
}

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetYangName() string { return "qosIpAclActionEntry" }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) SetParent(parent types.Entity) { qosipaclactionentry.parent = parent }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetParent() types.Entity { return qosipaclactionentry.parent }

func (qosipaclactionentry *CISCOQOSPIBMIB_Qosipaclactiontable_Qosipaclactionentry) GetParentYangName() string { return "qosIpAclActionTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies a scheduling preference for a
    // queue-type on an interface with a particular role combination. The type is
    // slice of
    // CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry.
    Qosifschedulingpreferenceentry []CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry
}

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetFilter() yfilter.YFilter { return qosifschedulingpreferencestable.YFilter }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) SetFilter(yf yfilter.YFilter) { qosifschedulingpreferencestable.YFilter = yf }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetGoName(yname string) string {
    if yname == "qosIfSchedulingPreferenceEntry" { return "Qosifschedulingpreferenceentry" }
    return ""
}

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetSegmentPath() string {
    return "qosIfSchedulingPreferencesTable"
}

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIfSchedulingPreferenceEntry" {
        for _, c := range qosifschedulingpreferencestable.Qosifschedulingpreferenceentry {
            if qosifschedulingpreferencestable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry{}
        qosifschedulingpreferencestable.Qosifschedulingpreferenceentry = append(qosifschedulingpreferencestable.Qosifschedulingpreferenceentry, child)
        return &qosifschedulingpreferencestable.Qosifschedulingpreferenceentry[len(qosifschedulingpreferencestable.Qosifschedulingpreferenceentry)-1]
    }
    return nil
}

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosifschedulingpreferencestable.Qosifschedulingpreferenceentry {
        children[qosifschedulingpreferencestable.Qosifschedulingpreferenceentry[i].GetSegmentPath()] = &qosifschedulingpreferencestable.Qosifschedulingpreferenceentry[i]
    }
    return children
}

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetBundleName() string { return "cisco_ios_xe" }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetYangName() string { return "qosIfSchedulingPreferencesTable" }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) SetParent(parent types.Entity) { qosifschedulingpreferencestable.parent = parent }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetParent() types.Entity { return qosifschedulingpreferencestable.parent }

func (qosifschedulingpreferencestable *CISCOQOSPIBMIB_Qosifschedulingpreferencestable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry
// An instance of this class specifies a scheduling preference
// for a queue-type on an interface with a particular role
// combination.
type CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry struct {
    parent types.Entity
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

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetFilter() yfilter.YFilter { return qosifschedulingpreferenceentry.YFilter }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) SetFilter(yf yfilter.YFilter) { qosifschedulingpreferenceentry.YFilter = yf }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetGoName(yname string) string {
    if yname == "qosIfSchedulingPreferenceId" { return "Qosifschedulingpreferenceid" }
    if yname == "qosIfSchedulingRoles" { return "Qosifschedulingroles" }
    if yname == "qosIfSchedulingPreference" { return "Qosifschedulingpreference" }
    if yname == "qosIfSchedulingDiscipline" { return "Qosifschedulingdiscipline" }
    if yname == "qosIfSchedulingQueueType" { return "Qosifschedulingqueuetype" }
    return ""
}

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetSegmentPath() string {
    return "qosIfSchedulingPreferenceEntry" + "[qosIfSchedulingPreferenceId='" + fmt.Sprintf("%v", qosifschedulingpreferenceentry.Qosifschedulingpreferenceid) + "']"
}

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIfSchedulingPreferenceId"] = qosifschedulingpreferenceentry.Qosifschedulingpreferenceid
    leafs["qosIfSchedulingRoles"] = qosifschedulingpreferenceentry.Qosifschedulingroles
    leafs["qosIfSchedulingPreference"] = qosifschedulingpreferenceentry.Qosifschedulingpreference
    leafs["qosIfSchedulingDiscipline"] = qosifschedulingpreferenceentry.Qosifschedulingdiscipline
    leafs["qosIfSchedulingQueueType"] = qosifschedulingpreferenceentry.Qosifschedulingqueuetype
    return leafs
}

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetYangName() string { return "qosIfSchedulingPreferenceEntry" }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) SetParent(parent types.Entity) { qosifschedulingpreferenceentry.parent = parent }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetParent() types.Entity { return qosifschedulingpreferenceentry.parent }

func (qosifschedulingpreferenceentry *CISCOQOSPIBMIB_Qosifschedulingpreferencestable_Qosifschedulingpreferenceentry) GetParentYangName() string { return "qosIfSchedulingPreferencesTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies a drop preference for a drop mechanism
    // on an interface with a particular role combination. The type is slice of
    // CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry.
    Qosifdroppreferenceentry []CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry
}

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetFilter() yfilter.YFilter { return qosifdroppreferencetable.YFilter }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) SetFilter(yf yfilter.YFilter) { qosifdroppreferencetable.YFilter = yf }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetGoName(yname string) string {
    if yname == "qosIfDropPreferenceEntry" { return "Qosifdroppreferenceentry" }
    return ""
}

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetSegmentPath() string {
    return "qosIfDropPreferenceTable"
}

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIfDropPreferenceEntry" {
        for _, c := range qosifdroppreferencetable.Qosifdroppreferenceentry {
            if qosifdroppreferencetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry{}
        qosifdroppreferencetable.Qosifdroppreferenceentry = append(qosifdroppreferencetable.Qosifdroppreferenceentry, child)
        return &qosifdroppreferencetable.Qosifdroppreferenceentry[len(qosifdroppreferencetable.Qosifdroppreferenceentry)-1]
    }
    return nil
}

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosifdroppreferencetable.Qosifdroppreferenceentry {
        children[qosifdroppreferencetable.Qosifdroppreferenceentry[i].GetSegmentPath()] = &qosifdroppreferencetable.Qosifdroppreferenceentry[i]
    }
    return children
}

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetBundleName() string { return "cisco_ios_xe" }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetYangName() string { return "qosIfDropPreferenceTable" }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) SetParent(parent types.Entity) { qosifdroppreferencetable.parent = parent }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetParent() types.Entity { return qosifdroppreferencetable.parent }

func (qosifdroppreferencetable *CISCOQOSPIBMIB_Qosifdroppreferencetable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry
// An instance of this class specifies a drop preference for
// a drop mechanism on an interface with a particular role
// combination.
type CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry struct {
    parent types.Entity
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

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetFilter() yfilter.YFilter { return qosifdroppreferenceentry.YFilter }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) SetFilter(yf yfilter.YFilter) { qosifdroppreferenceentry.YFilter = yf }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetGoName(yname string) string {
    if yname == "qosIfDropPreferenceId" { return "Qosifdroppreferenceid" }
    if yname == "qosIfDropRoles" { return "Qosifdroproles" }
    if yname == "qosIfDropPreference" { return "Qosifdroppreference" }
    if yname == "qosIfDropDiscipline" { return "Qosifdropdiscipline" }
    return ""
}

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetSegmentPath() string {
    return "qosIfDropPreferenceEntry" + "[qosIfDropPreferenceId='" + fmt.Sprintf("%v", qosifdroppreferenceentry.Qosifdroppreferenceid) + "']"
}

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIfDropPreferenceId"] = qosifdroppreferenceentry.Qosifdroppreferenceid
    leafs["qosIfDropRoles"] = qosifdroppreferenceentry.Qosifdroproles
    leafs["qosIfDropPreference"] = qosifdroppreferenceentry.Qosifdroppreference
    leafs["qosIfDropDiscipline"] = qosifdroppreferenceentry.Qosifdropdiscipline
    return leafs
}

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetYangName() string { return "qosIfDropPreferenceEntry" }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) SetParent(parent types.Entity) { qosifdroppreferenceentry.parent = parent }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetParent() types.Entity { return qosifdroppreferenceentry.parent }

func (qosifdroppreferenceentry *CISCOQOSPIBMIB_Qosifdroppreferencetable_Qosifdroppreferenceentry) GetParentYangName() string { return "qosIfDropPreferenceTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies the queue and threshold set for a
    // packet with a particular DSCP on an interface of a particular type with a
    // particular role combination. The type is slice of
    // CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry.
    Qosifdscpassignmententry []CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry
}

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetFilter() yfilter.YFilter { return qosifdscpassignmenttable.YFilter }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) SetFilter(yf yfilter.YFilter) { qosifdscpassignmenttable.YFilter = yf }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetGoName(yname string) string {
    if yname == "qosIfDscpAssignmentEntry" { return "Qosifdscpassignmententry" }
    return ""
}

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetSegmentPath() string {
    return "qosIfDscpAssignmentTable"
}

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIfDscpAssignmentEntry" {
        for _, c := range qosifdscpassignmenttable.Qosifdscpassignmententry {
            if qosifdscpassignmenttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry{}
        qosifdscpassignmenttable.Qosifdscpassignmententry = append(qosifdscpassignmenttable.Qosifdscpassignmententry, child)
        return &qosifdscpassignmenttable.Qosifdscpassignmententry[len(qosifdscpassignmenttable.Qosifdscpassignmententry)-1]
    }
    return nil
}

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosifdscpassignmenttable.Qosifdscpassignmententry {
        children[qosifdscpassignmenttable.Qosifdscpassignmententry[i].GetSegmentPath()] = &qosifdscpassignmenttable.Qosifdscpassignmententry[i]
    }
    return children
}

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetBundleName() string { return "cisco_ios_xe" }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetYangName() string { return "qosIfDscpAssignmentTable" }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) SetParent(parent types.Entity) { qosifdscpassignmenttable.parent = parent }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetParent() types.Entity { return qosifdscpassignmenttable.parent }

func (qosifdscpassignmenttable *CISCOQOSPIBMIB_Qosifdscpassignmenttable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry
// An instance of this class specifies the queue and threshold
// set for a packet with a particular DSCP on an interface of
// a particular type with a particular role combination.
type CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry struct {
    parent types.Entity
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

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetFilter() yfilter.YFilter { return qosifdscpassignmententry.YFilter }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) SetFilter(yf yfilter.YFilter) { qosifdscpassignmententry.YFilter = yf }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetGoName(yname string) string {
    if yname == "qosIfDscpAssignmentId" { return "Qosifdscpassignmentid" }
    if yname == "qosIfDscpRoles" { return "Qosifdscproles" }
    if yname == "qosIfQueueType" { return "Qosifqueuetype" }
    if yname == "qosIfDscp" { return "Qosifdscp" }
    if yname == "qosIfQueue" { return "Qosifqueue" }
    if yname == "qosIfThresholdSet" { return "Qosifthresholdset" }
    return ""
}

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetSegmentPath() string {
    return "qosIfDscpAssignmentEntry" + "[qosIfDscpAssignmentId='" + fmt.Sprintf("%v", qosifdscpassignmententry.Qosifdscpassignmentid) + "']"
}

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIfDscpAssignmentId"] = qosifdscpassignmententry.Qosifdscpassignmentid
    leafs["qosIfDscpRoles"] = qosifdscpassignmententry.Qosifdscproles
    leafs["qosIfQueueType"] = qosifdscpassignmententry.Qosifqueuetype
    leafs["qosIfDscp"] = qosifdscpassignmententry.Qosifdscp
    leafs["qosIfQueue"] = qosifdscpassignmententry.Qosifqueue
    leafs["qosIfThresholdSet"] = qosifdscpassignmententry.Qosifthresholdset
    return leafs
}

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetBundleName() string { return "cisco_ios_xe" }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetYangName() string { return "qosIfDscpAssignmentEntry" }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) SetParent(parent types.Entity) { qosifdscpassignmententry.parent = parent }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetParent() types.Entity { return qosifdscpassignmententry.parent }

func (qosifdscpassignmententry *CISCOQOSPIBMIB_Qosifdscpassignmenttable_Qosifdscpassignmententry) GetParentYangName() string { return "qosIfDscpAssignmentTable" }

// CISCOQOSPIBMIB_Qosifredtable
// A class of lower and upper values for each threshold set in a
// queue supporting WRED.  If the size of the queue for a given
// threshold is below the lower value then packets assigned to
// that threshold are always accepted into the queue.  If the
// size of the queue is above upper value then packets are always
// dropped.  If the size of the queue is between the lower and
// the upper then packets are randomly dropped.
type CISCOQOSPIBMIB_Qosifredtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies threshold limits for a particular RED
    // threshold of a given threshold set on an interface and with a particular
    // role combination. The type is slice of
    // CISCOQOSPIBMIB_Qosifredtable_Qosifredentry.
    Qosifredentry []CISCOQOSPIBMIB_Qosifredtable_Qosifredentry
}

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetFilter() yfilter.YFilter { return qosifredtable.YFilter }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) SetFilter(yf yfilter.YFilter) { qosifredtable.YFilter = yf }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetGoName(yname string) string {
    if yname == "qosIfRedEntry" { return "Qosifredentry" }
    return ""
}

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetSegmentPath() string {
    return "qosIfRedTable"
}

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIfRedEntry" {
        for _, c := range qosifredtable.Qosifredentry {
            if qosifredtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosifredtable_Qosifredentry{}
        qosifredtable.Qosifredentry = append(qosifredtable.Qosifredentry, child)
        return &qosifredtable.Qosifredentry[len(qosifredtable.Qosifredentry)-1]
    }
    return nil
}

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosifredtable.Qosifredentry {
        children[qosifredtable.Qosifredentry[i].GetSegmentPath()] = &qosifredtable.Qosifredentry[i]
    }
    return children
}

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetBundleName() string { return "cisco_ios_xe" }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetYangName() string { return "qosIfRedTable" }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) SetParent(parent types.Entity) { qosifredtable.parent = parent }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetParent() types.Entity { return qosifredtable.parent }

func (qosifredtable *CISCOQOSPIBMIB_Qosifredtable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosifredtable_Qosifredentry
// An instance of this class specifies threshold limits for a
// particular RED threshold of a given threshold set on an
// interface and with a particular role combination.
type CISCOQOSPIBMIB_Qosifredtable_Qosifredentry struct {
    parent types.Entity
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

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetFilter() yfilter.YFilter { return qosifredentry.YFilter }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) SetFilter(yf yfilter.YFilter) { qosifredentry.YFilter = yf }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetGoName(yname string) string {
    if yname == "qosIfRedId" { return "Qosifredid" }
    if yname == "qosIfRedRoles" { return "Qosifredroles" }
    if yname == "qosIfRedNumThresholdSets" { return "Qosifrednumthresholdsets" }
    if yname == "qosIfRedThresholdSet" { return "Qosifredthresholdset" }
    if yname == "qosIfRedThresholdSetLower" { return "Qosifredthresholdsetlower" }
    if yname == "qosIfRedThresholdSetUpper" { return "Qosifredthresholdsetupper" }
    return ""
}

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetSegmentPath() string {
    return "qosIfRedEntry" + "[qosIfRedId='" + fmt.Sprintf("%v", qosifredentry.Qosifredid) + "']"
}

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIfRedId"] = qosifredentry.Qosifredid
    leafs["qosIfRedRoles"] = qosifredentry.Qosifredroles
    leafs["qosIfRedNumThresholdSets"] = qosifredentry.Qosifrednumthresholdsets
    leafs["qosIfRedThresholdSet"] = qosifredentry.Qosifredthresholdset
    leafs["qosIfRedThresholdSetLower"] = qosifredentry.Qosifredthresholdsetlower
    leafs["qosIfRedThresholdSetUpper"] = qosifredentry.Qosifredthresholdsetupper
    return leafs
}

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetYangName() string { return "qosIfRedEntry" }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) SetParent(parent types.Entity) { qosifredentry.parent = parent }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetParent() types.Entity { return qosifredentry.parent }

func (qosifredentry *CISCOQOSPIBMIB_Qosifredtable_Qosifredentry) GetParentYangName() string { return "qosIfRedTable" }

// CISCOQOSPIBMIB_Qosiftaildroptable
// A class for threshold sets in a queue supporting tail drop.
// If the size of the queue for a given threshold set is at or
// below the specified value then packets assigned to that
// threshold set are always accepted into the queue.  If the size
// of the queue is above the specified value then packets are
// always dropped.
type CISCOQOSPIBMIB_Qosiftaildroptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies the queue depth for a particular
    // tail-drop threshold set on an interface with a particular role combination.
    // The type is slice of CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry.
    Qosiftaildropentry []CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry
}

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetFilter() yfilter.YFilter { return qosiftaildroptable.YFilter }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) SetFilter(yf yfilter.YFilter) { qosiftaildroptable.YFilter = yf }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetGoName(yname string) string {
    if yname == "qosIfTailDropEntry" { return "Qosiftaildropentry" }
    return ""
}

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetSegmentPath() string {
    return "qosIfTailDropTable"
}

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIfTailDropEntry" {
        for _, c := range qosiftaildroptable.Qosiftaildropentry {
            if qosiftaildroptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry{}
        qosiftaildroptable.Qosiftaildropentry = append(qosiftaildroptable.Qosiftaildropentry, child)
        return &qosiftaildroptable.Qosiftaildropentry[len(qosiftaildroptable.Qosiftaildropentry)-1]
    }
    return nil
}

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosiftaildroptable.Qosiftaildropentry {
        children[qosiftaildroptable.Qosiftaildropentry[i].GetSegmentPath()] = &qosiftaildroptable.Qosiftaildropentry[i]
    }
    return children
}

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetBundleName() string { return "cisco_ios_xe" }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetYangName() string { return "qosIfTailDropTable" }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) SetParent(parent types.Entity) { qosiftaildroptable.parent = parent }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetParent() types.Entity { return qosiftaildroptable.parent }

func (qosiftaildroptable *CISCOQOSPIBMIB_Qosiftaildroptable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry
// An instance of this class specifies the queue depth for a
// particular tail-drop threshold set on an interface with a
// particular role combination.
type CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry struct {
    parent types.Entity
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

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetFilter() yfilter.YFilter { return qosiftaildropentry.YFilter }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) SetFilter(yf yfilter.YFilter) { qosiftaildropentry.YFilter = yf }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetGoName(yname string) string {
    if yname == "qosIfTailDropId" { return "Qosiftaildropid" }
    if yname == "qosIfTailDropRoles" { return "Qosiftaildroproles" }
    if yname == "qosIfTailDropNumThresholdSets" { return "Qosiftaildropnumthresholdsets" }
    if yname == "qosIfTailDropThresholdSet" { return "Qosiftaildropthresholdset" }
    if yname == "qosIfTailDropThresholdSetValue" { return "Qosiftaildropthresholdsetvalue" }
    return ""
}

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetSegmentPath() string {
    return "qosIfTailDropEntry" + "[qosIfTailDropId='" + fmt.Sprintf("%v", qosiftaildropentry.Qosiftaildropid) + "']"
}

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIfTailDropId"] = qosiftaildropentry.Qosiftaildropid
    leafs["qosIfTailDropRoles"] = qosiftaildropentry.Qosiftaildroproles
    leafs["qosIfTailDropNumThresholdSets"] = qosiftaildropentry.Qosiftaildropnumthresholdsets
    leafs["qosIfTailDropThresholdSet"] = qosiftaildropentry.Qosiftaildropthresholdset
    leafs["qosIfTailDropThresholdSetValue"] = qosiftaildropentry.Qosiftaildropthresholdsetvalue
    return leafs
}

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetYangName() string { return "qosIfTailDropEntry" }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) SetParent(parent types.Entity) { qosiftaildropentry.parent = parent }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetParent() types.Entity { return qosiftaildropentry.parent }

func (qosiftaildropentry *CISCOQOSPIBMIB_Qosiftaildroptable_Qosiftaildropentry) GetParentYangName() string { return "qosIfTailDropTable" }

// CISCOQOSPIBMIB_Qosifweightstable
// A class of scheduling weights for each queue of an interface
// that supports weighted round robin scheduling or a mix of
// priority queueing and weighted round robin.  For a queue with
// N priority queues, the N highest queue numbers are the
// priority queues with the highest queue number having the
// highest priority.  WRR is applied to the non-priority queues.
type CISCOQOSPIBMIB_Qosifweightstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An instance of this class specifies the scheduling weight for a particular
    // queue of an interface with a particular number of queues and with a
    // particular role combination. The type is slice of
    // CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry.
    Qosifweightsentry []CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry
}

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetFilter() yfilter.YFilter { return qosifweightstable.YFilter }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) SetFilter(yf yfilter.YFilter) { qosifweightstable.YFilter = yf }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetGoName(yname string) string {
    if yname == "qosIfWeightsEntry" { return "Qosifweightsentry" }
    return ""
}

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetSegmentPath() string {
    return "qosIfWeightsTable"
}

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qosIfWeightsEntry" {
        for _, c := range qosifweightstable.Qosifweightsentry {
            if qosifweightstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry{}
        qosifweightstable.Qosifweightsentry = append(qosifweightstable.Qosifweightsentry, child)
        return &qosifweightstable.Qosifweightsentry[len(qosifweightstable.Qosifweightsentry)-1]
    }
    return nil
}

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range qosifweightstable.Qosifweightsentry {
        children[qosifweightstable.Qosifweightsentry[i].GetSegmentPath()] = &qosifweightstable.Qosifweightsentry[i]
    }
    return children
}

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetBundleName() string { return "cisco_ios_xe" }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetYangName() string { return "qosIfWeightsTable" }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) SetParent(parent types.Entity) { qosifweightstable.parent = parent }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetParent() types.Entity { return qosifweightstable.parent }

func (qosifweightstable *CISCOQOSPIBMIB_Qosifweightstable) GetParentYangName() string { return "CISCO-QOS-PIB-MIB" }

// CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry
// An instance of this class specifies the scheduling weight for
// a particular queue of an interface with a particular number
// of queues and with a particular role combination.
type CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry struct {
    parent types.Entity
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

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetFilter() yfilter.YFilter { return qosifweightsentry.YFilter }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) SetFilter(yf yfilter.YFilter) { qosifweightsentry.YFilter = yf }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetGoName(yname string) string {
    if yname == "qosIfWeightsId" { return "Qosifweightsid" }
    if yname == "qosIfWeightsRoles" { return "Qosifweightsroles" }
    if yname == "qosIfWeightsNumQueues" { return "Qosifweightsnumqueues" }
    if yname == "qosIfWeightsQueue" { return "Qosifweightsqueue" }
    if yname == "qosIfWeightsDrainSize" { return "Qosifweightsdrainsize" }
    if yname == "qosIfWeightsQueueSize" { return "Qosifweightsqueuesize" }
    return ""
}

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetSegmentPath() string {
    return "qosIfWeightsEntry" + "[qosIfWeightsId='" + fmt.Sprintf("%v", qosifweightsentry.Qosifweightsid) + "']"
}

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qosIfWeightsId"] = qosifweightsentry.Qosifweightsid
    leafs["qosIfWeightsRoles"] = qosifweightsentry.Qosifweightsroles
    leafs["qosIfWeightsNumQueues"] = qosifweightsentry.Qosifweightsnumqueues
    leafs["qosIfWeightsQueue"] = qosifweightsentry.Qosifweightsqueue
    leafs["qosIfWeightsDrainSize"] = qosifweightsentry.Qosifweightsdrainsize
    leafs["qosIfWeightsQueueSize"] = qosifweightsentry.Qosifweightsqueuesize
    return leafs
}

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetBundleName() string { return "cisco_ios_xe" }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetYangName() string { return "qosIfWeightsEntry" }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) SetParent(parent types.Entity) { qosifweightsentry.parent = parent }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetParent() types.Entity { return qosifweightsentry.parent }

func (qosifweightsentry *CISCOQOSPIBMIB_Qosifweightstable_Qosifweightsentry) GetParentYangName() string { return "qosIfWeightsTable" }

