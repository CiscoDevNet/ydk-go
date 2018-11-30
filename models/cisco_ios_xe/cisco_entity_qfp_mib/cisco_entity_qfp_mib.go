// This MIB module defines managed objects that facilitate the
// management of Quantum Flow Processors (QFP), which are listed
// in the ENTITY-MIB (RFC 4133) entPhysicalTable as an
// entPhysicalClass of 'cpu'.
// 
// The Quantum Flow Processors (QFP) technology is an architecture
// family developed by Cisco, which has fully integrated and
// programmable networking chip set that controls different
// functions of a system such as packet forwarding.
// 
// This module contains objects to monitor various QFP
// statistics such as system, utilization, memory, etc.
// 
// The utilization statistics can be used for future capacity
// planning. The calculation of this statistics may vary by
// devices which host QFPs, hence the user must refer to the
// specific device document of interest for the exact method
// of calculation. The utilization statistics have the following
// terminology.
// 
// o Input - Communication channel where packets arrive on the
//           QFP.
// 
// o Output - Communication channel where packets leave the QFP.
// 
// o Priority - A classification applied to packets indicating
//              they should be treated with greater urgency.
// 
// o Non-Priority - A classification applied to packets indicating
//                  they should be treated with lesser urgency.
// 
// o Processing Load - The percentage of time over some interval
//                     that the QFP has spent forwarding packets,
//                     relative to the total time spent both 
//                     forwarding packets and being idle.
package cisco_entity_qfp_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_entity_qfp_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ENTITY-QFP-MIB CISCO-ENTITY-QFP-MIB}", reflect.TypeOf(CISCOENTITYQFPMIB{}))
    ydk.RegisterEntity("CISCO-ENTITY-QFP-MIB:CISCO-ENTITY-QFP-MIB", reflect.TypeOf(CISCOENTITYQFPMIB{}))
}

// CiscoQfpTimeInterval represents                     statistics
type CiscoQfpTimeInterval string

const (
    CiscoQfpTimeInterval_fiveSeconds CiscoQfpTimeInterval = "fiveSeconds"

    CiscoQfpTimeInterval_oneMinute CiscoQfpTimeInterval = "oneMinute"

    CiscoQfpTimeInterval_fiveMinutes CiscoQfpTimeInterval = "fiveMinutes"

    CiscoQfpTimeInterval_sixtyMinutes CiscoQfpTimeInterval = "sixtyMinutes"
)

// CiscoQfpMemoryResource represents dram (1) - Dynamic Random Access Memory (DRAM) memory resource
type CiscoQfpMemoryResource string

const (
    CiscoQfpMemoryResource_dram CiscoQfpMemoryResource = "dram"
)

// CISCOENTITYQFPMIB
type CISCOENTITYQFPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CiscoEntityQfp CISCOENTITYQFPMIB_CiscoEntityQfp

    
    CiscoEntityQfpNotif CISCOENTITYQFPMIB_CiscoEntityQfpNotif

    // This table maintains the QFP system information for each QFP physical
    // entity.  An agent creates a conceptual row to this table corresponding to a
    // QFP physical entity upon detection of a physical entity supporting the QFP
    // system information.  An agent destroys a conceptual row from this table    
    // corresponding to a QFP physical entity upon removal of the QFP host
    // physical entity.
    CeqfpSystemTable CISCOENTITYQFPMIB_CeqfpSystemTable

    // This table maintains the utilization statistics collected by each QFP
    // physical entity at various time interval such as last 5 seconds, 1 minute,
    // etc.  An agent creates a conceptual row to this table corresponding to a
    // QFP physical entity for a time interval upon detection of a physical entity
    // supporting the utilization statistics for a time interval.  The agent
    // destroys a conceptual row from this table        corresponding to a QFP
    // physical entity for a time interval upon removal of the QFP host physical
    // entity or it does not receive the utilization statistics update for a
    // certain time period. The time period to wait before deleting an entry from
    // this table would be the discretion of the supporting device.
    CeqfpUtilizationTable CISCOENTITYQFPMIB_CeqfpUtilizationTable

    // This table maintains the memory resources statistics for each QFP physical
    // entity.  An agent creates a conceptual row to this table corresponding to a
    // QFP physical entity and its supported memory resource type upon detection
    // of a physical entity supporting the memory  resource statistics for a
    // memory resource type.  An agent destroys a conceptual row from this table  
    // corresponding to a QFP physical entity and its supported memory resource
    // type upon removal of the QFP host physical entity or it does not receive
    // memory resource statistics update for a certain time period. The time
    // period to wait before deleting an entry from this table would be the
    // discretion of the supporting device.
    CeqfpMemoryResourceTable CISCOENTITYQFPMIB_CeqfpMemoryResourceTable

    // This table maintains the throughput information for each QFP physical
    // entity.          An agent creates a conceptual row to this table
    // corresponding to a QFP physical entity upon detection of a physical entity
    // supporting the QFP throughput information.          An agent destroys a
    // conceptual row from this table       corresponding to a QFP physical entity
    // upon removal of the QFP host physical entity.
    CeqfpThroughputTable CISCOENTITYQFPMIB_CeqfpThroughputTable
}

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetEntityData() *types.CommonEntityData {
    cISCOENTITYQFPMIB.EntityData.YFilter = cISCOENTITYQFPMIB.YFilter
    cISCOENTITYQFPMIB.EntityData.YangName = "CISCO-ENTITY-QFP-MIB"
    cISCOENTITYQFPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOENTITYQFPMIB.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    cISCOENTITYQFPMIB.EntityData.SegmentPath = "CISCO-ENTITY-QFP-MIB:CISCO-ENTITY-QFP-MIB"
    cISCOENTITYQFPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOENTITYQFPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOENTITYQFPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOENTITYQFPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOENTITYQFPMIB.EntityData.Children.Append("ciscoEntityQfp", types.YChild{"CiscoEntityQfp", &cISCOENTITYQFPMIB.CiscoEntityQfp})
    cISCOENTITYQFPMIB.EntityData.Children.Append("ciscoEntityQfpNotif", types.YChild{"CiscoEntityQfpNotif", &cISCOENTITYQFPMIB.CiscoEntityQfpNotif})
    cISCOENTITYQFPMIB.EntityData.Children.Append("ceqfpSystemTable", types.YChild{"CeqfpSystemTable", &cISCOENTITYQFPMIB.CeqfpSystemTable})
    cISCOENTITYQFPMIB.EntityData.Children.Append("ceqfpUtilizationTable", types.YChild{"CeqfpUtilizationTable", &cISCOENTITYQFPMIB.CeqfpUtilizationTable})
    cISCOENTITYQFPMIB.EntityData.Children.Append("ceqfpMemoryResourceTable", types.YChild{"CeqfpMemoryResourceTable", &cISCOENTITYQFPMIB.CeqfpMemoryResourceTable})
    cISCOENTITYQFPMIB.EntityData.Children.Append("ceqfpThroughputTable", types.YChild{"CeqfpThroughputTable", &cISCOENTITYQFPMIB.CeqfpThroughputTable})
    cISCOENTITYQFPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOENTITYQFPMIB.EntityData.YListKeys = []string {}

    return &(cISCOENTITYQFPMIB.EntityData)
}

// CISCOENTITYQFPMIB_CiscoEntityQfp
type CISCOENTITYQFPMIB_CiscoEntityQfp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This objects represents the method used to calculate 5 Second interval
    // utilization data. The enumerated values are described below.  unknown      
    // (1) - The calculation method is unknown fiveSecSample (2) - The value is
    // calculated based on the last                     5 second sampling period
    // of utilization                     data. The type is
    // CeqfpFiveSecondUtilAlgo.
    CeqfpFiveSecondUtilAlgo interface{}

    // This objects represents the method used to calculate 1 Minute interval
    // utilization data. The enumerated values are described below.  unknown   
    // (1) - The calculation method is unknown fiveSecSMA (2) - The value is
    // calculated using Simple Moving                    Average of last 12 five
    // seconds utilization                  data. The type is
    // CeqfpOneMinuteUtilAlgo.
    CeqfpOneMinuteUtilAlgo interface{}

    // This objects represents the method used to calculate 5 Minutes interval
    // utilization data. The enumerated values are described below.  unknown   
    // (1) - The calculation method is unknown fiveSecSMA (2) - The value is
    // calculated using Simple Moving                    Average of last 60 five
    // seconds utilization                  data. The type is
    // CeqfpFiveMinutesUtilAlgo.
    CeqfpFiveMinutesUtilAlgo interface{}

    // This objects represents the method used to calculate 60 Minutes interval
    // utilization data. The enumerated values are described below.  unknown   
    // (1) - The calculation method is unknown fiveSecSMA (1) - The value is
    // calculated using Simple Moving                    Average of last 720 five
    // seconds utilization                  data. The type is
    // CeqfpSixtyMinutesUtilAlgo.
    CeqfpSixtyMinutesUtilAlgo interface{}
}

func (ciscoEntityQfp *CISCOENTITYQFPMIB_CiscoEntityQfp) GetEntityData() *types.CommonEntityData {
    ciscoEntityQfp.EntityData.YFilter = ciscoEntityQfp.YFilter
    ciscoEntityQfp.EntityData.YangName = "ciscoEntityQfp"
    ciscoEntityQfp.EntityData.BundleName = "cisco_ios_xe"
    ciscoEntityQfp.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ciscoEntityQfp.EntityData.SegmentPath = "ciscoEntityQfp"
    ciscoEntityQfp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEntityQfp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEntityQfp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEntityQfp.EntityData.Children = types.NewOrderedMap()
    ciscoEntityQfp.EntityData.Leafs = types.NewOrderedMap()
    ciscoEntityQfp.EntityData.Leafs.Append("ceqfpFiveSecondUtilAlgo", types.YLeaf{"CeqfpFiveSecondUtilAlgo", ciscoEntityQfp.CeqfpFiveSecondUtilAlgo})
    ciscoEntityQfp.EntityData.Leafs.Append("ceqfpOneMinuteUtilAlgo", types.YLeaf{"CeqfpOneMinuteUtilAlgo", ciscoEntityQfp.CeqfpOneMinuteUtilAlgo})
    ciscoEntityQfp.EntityData.Leafs.Append("ceqfpFiveMinutesUtilAlgo", types.YLeaf{"CeqfpFiveMinutesUtilAlgo", ciscoEntityQfp.CeqfpFiveMinutesUtilAlgo})
    ciscoEntityQfp.EntityData.Leafs.Append("ceqfpSixtyMinutesUtilAlgo", types.YLeaf{"CeqfpSixtyMinutesUtilAlgo", ciscoEntityQfp.CeqfpSixtyMinutesUtilAlgo})

    ciscoEntityQfp.EntityData.YListKeys = []string {}

    return &(ciscoEntityQfp.EntityData)
}

// CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveMinutesUtilAlgo represents                  data.
type CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveMinutesUtilAlgo string

const (
    CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveMinutesUtilAlgo_unknown CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveMinutesUtilAlgo = "unknown"

    CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveMinutesUtilAlgo_fiveSecSMA CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveMinutesUtilAlgo = "fiveSecSMA"
)

// CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveSecondUtilAlgo represents                     data.
type CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveSecondUtilAlgo string

const (
    CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveSecondUtilAlgo_unknown CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveSecondUtilAlgo = "unknown"

    CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveSecondUtilAlgo_fiveSecSample CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpFiveSecondUtilAlgo = "fiveSecSample"
)

// CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpOneMinuteUtilAlgo represents                  data.
type CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpOneMinuteUtilAlgo string

const (
    CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpOneMinuteUtilAlgo_unknown CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpOneMinuteUtilAlgo = "unknown"

    CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpOneMinuteUtilAlgo_fiveSecSMA CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpOneMinuteUtilAlgo = "fiveSecSMA"
)

// CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpSixtyMinutesUtilAlgo represents                  data.
type CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpSixtyMinutesUtilAlgo string

const (
    CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpSixtyMinutesUtilAlgo_unknown CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpSixtyMinutesUtilAlgo = "unknown"

    CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpSixtyMinutesUtilAlgo_fiveSecSMA CISCOENTITYQFPMIB_CiscoEntityQfp_CeqfpSixtyMinutesUtilAlgo = "fiveSecSMA"
)

// CISCOENTITYQFPMIB_CiscoEntityQfpNotif
type CISCOENTITYQFPMIB_CiscoEntityQfpNotif struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object controls memory resource rising and falling threshold
    // notification.  When this object contains a value 'true', then generation of
    // memory resource threshold notification is enabled. If this object contains
    // a value 'false', then generation of memory resource threshold notification
    // is disabled. The type is bool.
    CeqfpMemoryResThreshNotifEnabled interface{}

    // This object controls throughput rate notification.  When this object
    // contains a value 'true', then generation of ceqfpThroughputNotif is
    // enabled. If this object contains a value 'false', then generation of
    // ceqfpThroughputNotif is disabled. The type is interface{} with range: 0..2.
    CeqfpThroughputNotifEnabled interface{}
}

func (ciscoEntityQfpNotif *CISCOENTITYQFPMIB_CiscoEntityQfpNotif) GetEntityData() *types.CommonEntityData {
    ciscoEntityQfpNotif.EntityData.YFilter = ciscoEntityQfpNotif.YFilter
    ciscoEntityQfpNotif.EntityData.YangName = "ciscoEntityQfpNotif"
    ciscoEntityQfpNotif.EntityData.BundleName = "cisco_ios_xe"
    ciscoEntityQfpNotif.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ciscoEntityQfpNotif.EntityData.SegmentPath = "ciscoEntityQfpNotif"
    ciscoEntityQfpNotif.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoEntityQfpNotif.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoEntityQfpNotif.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoEntityQfpNotif.EntityData.Children = types.NewOrderedMap()
    ciscoEntityQfpNotif.EntityData.Leafs = types.NewOrderedMap()
    ciscoEntityQfpNotif.EntityData.Leafs.Append("ceqfpMemoryResThreshNotifEnabled", types.YLeaf{"CeqfpMemoryResThreshNotifEnabled", ciscoEntityQfpNotif.CeqfpMemoryResThreshNotifEnabled})
    ciscoEntityQfpNotif.EntityData.Leafs.Append("ceqfpThroughputNotifEnabled", types.YLeaf{"CeqfpThroughputNotifEnabled", ciscoEntityQfpNotif.CeqfpThroughputNotifEnabled})

    ciscoEntityQfpNotif.EntityData.YListKeys = []string {}

    return &(ciscoEntityQfpNotif.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpSystemTable
// This table maintains the QFP system information for each QFP
// physical entity.
// 
// An agent creates a conceptual row to this table corresponding
// to a QFP physical entity upon detection of a physical entity
// supporting the QFP system information.
// 
// An agent destroys a conceptual row from this table       
// corresponding to a QFP physical entity upon removal
// of the QFP host physical entity.
type CISCOENTITYQFPMIB_CeqfpSystemTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpSystemTable. There is an entry in this table
    // for each QFP entity, as defined by a value of entPhysicalIndex. The type is
    // slice of CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry.
    CeqfpSystemEntry []*CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry
}

func (ceqfpSystemTable *CISCOENTITYQFPMIB_CeqfpSystemTable) GetEntityData() *types.CommonEntityData {
    ceqfpSystemTable.EntityData.YFilter = ceqfpSystemTable.YFilter
    ceqfpSystemTable.EntityData.YangName = "ceqfpSystemTable"
    ceqfpSystemTable.EntityData.BundleName = "cisco_ios_xe"
    ceqfpSystemTable.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ceqfpSystemTable.EntityData.SegmentPath = "ceqfpSystemTable"
    ceqfpSystemTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpSystemTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpSystemTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpSystemTable.EntityData.Children = types.NewOrderedMap()
    ceqfpSystemTable.EntityData.Children.Append("ceqfpSystemEntry", types.YChild{"CeqfpSystemEntry", nil})
    for i := range ceqfpSystemTable.CeqfpSystemEntry {
        ceqfpSystemTable.EntityData.Children.Append(types.GetSegmentPath(ceqfpSystemTable.CeqfpSystemEntry[i]), types.YChild{"CeqfpSystemEntry", ceqfpSystemTable.CeqfpSystemEntry[i]})
    }
    ceqfpSystemTable.EntityData.Leafs = types.NewOrderedMap()

    ceqfpSystemTable.EntityData.YListKeys = []string {}

    return &(ceqfpSystemTable.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry
// A conceptual row in the ceqfpSystemTable. There is an entry
// in this table for each QFP entity, as defined by a value of
// entPhysicalIndex.
type CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This object represents the traffic direction that this QFP is assigned to
    // process. The enumerated values are described below.  none (1)    - The QFP
    // is not assigned to processes any traffic               yet ingress (2) -
    // The QFP processes inbound traffic egress (3)  - The QFP processes outbound
    // traffic both (4)    - The QFP processes both inbound and outbound          
    // traffic. The type is CeqfpSystemTrafficDirection.
    CeqfpSystemTrafficDirection interface{}

    // This object represents the current QFP state. The enumerated values are
    // described below.  unknown (1)    - The state of the QFP is unknown reset
    // (2)      - The QFP is reset init (3)       - The QFP is being initialized
    // active (4)     - The QFP is active in a system with redundant              
    // QFP activeSolo (5) - The QFP is active and there is no redundant           
    // QFP in the system standby (6)    - The QFP is standby in a redundant
    // system. hotStandby (7) - The QFP is standby and synchronized with          
    // active, so that a switchover in this state                  will preserve
    // state of the active. Stateful                   datapath features are
    // synchronized between the                  active QFP and standby QFP. The
    // type is CeqfpSystemState.
    CeqfpSystemState interface{}

    // This object represents the number of times the QFP is loaded, since the QFP
    // host is up. The type is interface{} with range: 0..4294967295.
    CeqfpNumberSystemLoads interface{}

    // This object represents the QFP last load time. The type is string.
    CeqfpSystemLastLoadTime interface{}
}

func (ceqfpSystemEntry *CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry) GetEntityData() *types.CommonEntityData {
    ceqfpSystemEntry.EntityData.YFilter = ceqfpSystemEntry.YFilter
    ceqfpSystemEntry.EntityData.YangName = "ceqfpSystemEntry"
    ceqfpSystemEntry.EntityData.BundleName = "cisco_ios_xe"
    ceqfpSystemEntry.EntityData.ParentYangName = "ceqfpSystemTable"
    ceqfpSystemEntry.EntityData.SegmentPath = "ceqfpSystemEntry" + types.AddKeyToken(ceqfpSystemEntry.EntPhysicalIndex, "entPhysicalIndex")
    ceqfpSystemEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpSystemEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpSystemEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpSystemEntry.EntityData.Children = types.NewOrderedMap()
    ceqfpSystemEntry.EntityData.Leafs = types.NewOrderedMap()
    ceqfpSystemEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ceqfpSystemEntry.EntPhysicalIndex})
    ceqfpSystemEntry.EntityData.Leafs.Append("ceqfpSystemTrafficDirection", types.YLeaf{"CeqfpSystemTrafficDirection", ceqfpSystemEntry.CeqfpSystemTrafficDirection})
    ceqfpSystemEntry.EntityData.Leafs.Append("ceqfpSystemState", types.YLeaf{"CeqfpSystemState", ceqfpSystemEntry.CeqfpSystemState})
    ceqfpSystemEntry.EntityData.Leafs.Append("ceqfpNumberSystemLoads", types.YLeaf{"CeqfpNumberSystemLoads", ceqfpSystemEntry.CeqfpNumberSystemLoads})
    ceqfpSystemEntry.EntityData.Leafs.Append("ceqfpSystemLastLoadTime", types.YLeaf{"CeqfpSystemLastLoadTime", ceqfpSystemEntry.CeqfpSystemLastLoadTime})

    ceqfpSystemEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(ceqfpSystemEntry.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState represents                  active QFP and standby QFP
type CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState string

const (
    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState_unknown CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState = "unknown"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState_reset CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState = "reset"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState_init CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState = "init"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState_active CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState = "active"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState_activeSolo CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState = "activeSolo"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState_standby CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState = "standby"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState_hotStandby CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemState = "hotStandby"
)

// CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection represents               traffic
type CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection string

const (
    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection_none CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection = "none"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection_ingress CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection = "ingress"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection_egress CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection = "egress"

    CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection_both CISCOENTITYQFPMIB_CeqfpSystemTable_CeqfpSystemEntry_CeqfpSystemTrafficDirection = "both"
)

// CISCOENTITYQFPMIB_CeqfpUtilizationTable
// This table maintains the utilization statistics collected
// by each QFP physical entity at various time interval such as
// last 5 seconds, 1 minute, etc.
// 
// An agent creates a conceptual row to this table corresponding
// to a QFP physical entity for a time interval upon detection of
// a physical entity supporting the utilization statistics for a
// time interval.
// 
// The agent destroys a conceptual row from this table       
// corresponding to a QFP physical entity for a time interval
// upon removal of the QFP host physical entity or it does not
// receive the utilization statistics update for a certain time
// period. The time period to wait before deleting an entry from
// this table would be the discretion of the supporting device.
type CISCOENTITYQFPMIB_CeqfpUtilizationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpUtilizationTable. There is an entry in this
    // table for each QFP entity by a value of entPhysicalIndex and the supported
    // time interval by a value  of ceqfpUtilTimeInterval.  The method of
    // utilization data calculation for each interval period can be identified
    // through the respective interval scalar objects. For example the utilizaiton
    // data calculation method for 'fiveSecond' interval can be identified by
    // ceqfpFiveSecondUtilAlgo. The type is slice of
    // CISCOENTITYQFPMIB_CeqfpUtilizationTable_CeqfpUtilizationEntry.
    CeqfpUtilizationEntry []*CISCOENTITYQFPMIB_CeqfpUtilizationTable_CeqfpUtilizationEntry
}

func (ceqfpUtilizationTable *CISCOENTITYQFPMIB_CeqfpUtilizationTable) GetEntityData() *types.CommonEntityData {
    ceqfpUtilizationTable.EntityData.YFilter = ceqfpUtilizationTable.YFilter
    ceqfpUtilizationTable.EntityData.YangName = "ceqfpUtilizationTable"
    ceqfpUtilizationTable.EntityData.BundleName = "cisco_ios_xe"
    ceqfpUtilizationTable.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ceqfpUtilizationTable.EntityData.SegmentPath = "ceqfpUtilizationTable"
    ceqfpUtilizationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpUtilizationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpUtilizationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpUtilizationTable.EntityData.Children = types.NewOrderedMap()
    ceqfpUtilizationTable.EntityData.Children.Append("ceqfpUtilizationEntry", types.YChild{"CeqfpUtilizationEntry", nil})
    for i := range ceqfpUtilizationTable.CeqfpUtilizationEntry {
        ceqfpUtilizationTable.EntityData.Children.Append(types.GetSegmentPath(ceqfpUtilizationTable.CeqfpUtilizationEntry[i]), types.YChild{"CeqfpUtilizationEntry", ceqfpUtilizationTable.CeqfpUtilizationEntry[i]})
    }
    ceqfpUtilizationTable.EntityData.Leafs = types.NewOrderedMap()

    ceqfpUtilizationTable.EntityData.YListKeys = []string {}

    return &(ceqfpUtilizationTable.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpUtilizationTable_CeqfpUtilizationEntry
// A conceptual row in the ceqfpUtilizationTable. There is
// an entry in this table for each QFP entity by a value of
// entPhysicalIndex and the supported time interval by a value 
// of ceqfpUtilTimeInterval.
// 
// The method of utilization data calculation for each interval
// period can be identified through the respective interval
// scalar objects. For example the utilizaiton data calculation
// method for 'fiveSecond' interval can be identified by
// ceqfpFiveSecondUtilAlgo.
type CISCOENTITYQFPMIB_CeqfpUtilizationTable_CeqfpUtilizationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. This object identifies the time interval for which
    // the utilization statistics being collected. The interval  values can be 5
    // second, 1 minute, etc. as specified in  the CiscoQfpTimeInterval. The type
    // is CiscoQfpTimeInterval.
    CeqfpUtilTimeInterval interface{}

    // This object represents the QFP input channel priority packet rate during
    // this interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are packets per second.
    CeqfpUtilInputPriorityPktRate interface{}

    // This object represents the QFP input channel priority bit rate during this
    // interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits per second.
    CeqfpUtilInputPriorityBitRate interface{}

    // This object represents the QFP input channel non-priority packet rate
    // during this interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets per second.
    CeqfpUtilInputNonPriorityPktRate interface{}

    // This object represents the QFP input channel non-priority bit rate during
    // this interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits per second.
    CeqfpUtilInputNonPriorityBitRate interface{}

    // This object represents the QFP input channel total packet rate during this
    // interval, which includes both priority and non-priority input packet rate.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets per second.
    CeqfpUtilInputTotalPktRate interface{}

    // This object represents the QFP input channel total bit rate during this
    // interval, which includes both priority and non-priority input bit rate. The
    // type is interface{} with range: 0..18446744073709551615. Units are bits per
    // second.
    CeqfpUtilInputTotalBitRate interface{}

    // This object represents the QFP output channel priority packet rate during
    // this interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are packets per second.
    CeqfpUtilOutputPriorityPktRate interface{}

    // This object represents the QFP output channel priority bit rate during this
    // interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits per second.
    CeqfpUtilOutputPriorityBitRate interface{}

    // This object represents the QFP output channel non-priority packet rate
    // during this interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets per second.
    CeqfpUtilOutputNonPriorityPktRate interface{}

    // This object represents the QFP output channel non-priority bit rate during
    // this interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits per second.
    CeqfpUtilOutputNonPriorityBitRate interface{}

    // This object represents the QFP output channel total packet rate during this
    // interval, which includes both priority and non-priority output packet rate.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets per second.
    CeqfpUtilOutputTotalPktRate interface{}

    // This object represents the QFP output channel total bit rate during this
    // interval, which includes both priority and non-priority bit rate. The type
    // is interface{} with range: 0..18446744073709551615. Units are bits per
    // second.
    CeqfpUtilOutputTotalBitRate interface{}

    // This object represents the QFP processing load during this interval. The
    // type is interface{} with range: 0..100. Units are percent.
    CeqfpUtilProcessingLoad interface{}
}

func (ceqfpUtilizationEntry *CISCOENTITYQFPMIB_CeqfpUtilizationTable_CeqfpUtilizationEntry) GetEntityData() *types.CommonEntityData {
    ceqfpUtilizationEntry.EntityData.YFilter = ceqfpUtilizationEntry.YFilter
    ceqfpUtilizationEntry.EntityData.YangName = "ceqfpUtilizationEntry"
    ceqfpUtilizationEntry.EntityData.BundleName = "cisco_ios_xe"
    ceqfpUtilizationEntry.EntityData.ParentYangName = "ceqfpUtilizationTable"
    ceqfpUtilizationEntry.EntityData.SegmentPath = "ceqfpUtilizationEntry" + types.AddKeyToken(ceqfpUtilizationEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(ceqfpUtilizationEntry.CeqfpUtilTimeInterval, "ceqfpUtilTimeInterval")
    ceqfpUtilizationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpUtilizationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpUtilizationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpUtilizationEntry.EntityData.Children = types.NewOrderedMap()
    ceqfpUtilizationEntry.EntityData.Leafs = types.NewOrderedMap()
    ceqfpUtilizationEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ceqfpUtilizationEntry.EntPhysicalIndex})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilTimeInterval", types.YLeaf{"CeqfpUtilTimeInterval", ceqfpUtilizationEntry.CeqfpUtilTimeInterval})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilInputPriorityPktRate", types.YLeaf{"CeqfpUtilInputPriorityPktRate", ceqfpUtilizationEntry.CeqfpUtilInputPriorityPktRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilInputPriorityBitRate", types.YLeaf{"CeqfpUtilInputPriorityBitRate", ceqfpUtilizationEntry.CeqfpUtilInputPriorityBitRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilInputNonPriorityPktRate", types.YLeaf{"CeqfpUtilInputNonPriorityPktRate", ceqfpUtilizationEntry.CeqfpUtilInputNonPriorityPktRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilInputNonPriorityBitRate", types.YLeaf{"CeqfpUtilInputNonPriorityBitRate", ceqfpUtilizationEntry.CeqfpUtilInputNonPriorityBitRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilInputTotalPktRate", types.YLeaf{"CeqfpUtilInputTotalPktRate", ceqfpUtilizationEntry.CeqfpUtilInputTotalPktRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilInputTotalBitRate", types.YLeaf{"CeqfpUtilInputTotalBitRate", ceqfpUtilizationEntry.CeqfpUtilInputTotalBitRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilOutputPriorityPktRate", types.YLeaf{"CeqfpUtilOutputPriorityPktRate", ceqfpUtilizationEntry.CeqfpUtilOutputPriorityPktRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilOutputPriorityBitRate", types.YLeaf{"CeqfpUtilOutputPriorityBitRate", ceqfpUtilizationEntry.CeqfpUtilOutputPriorityBitRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilOutputNonPriorityPktRate", types.YLeaf{"CeqfpUtilOutputNonPriorityPktRate", ceqfpUtilizationEntry.CeqfpUtilOutputNonPriorityPktRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilOutputNonPriorityBitRate", types.YLeaf{"CeqfpUtilOutputNonPriorityBitRate", ceqfpUtilizationEntry.CeqfpUtilOutputNonPriorityBitRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilOutputTotalPktRate", types.YLeaf{"CeqfpUtilOutputTotalPktRate", ceqfpUtilizationEntry.CeqfpUtilOutputTotalPktRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilOutputTotalBitRate", types.YLeaf{"CeqfpUtilOutputTotalBitRate", ceqfpUtilizationEntry.CeqfpUtilOutputTotalBitRate})
    ceqfpUtilizationEntry.EntityData.Leafs.Append("ceqfpUtilProcessingLoad", types.YLeaf{"CeqfpUtilProcessingLoad", ceqfpUtilizationEntry.CeqfpUtilProcessingLoad})

    ceqfpUtilizationEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CeqfpUtilTimeInterval"}

    return &(ceqfpUtilizationEntry.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpMemoryResourceTable
// This table maintains the memory resources statistics for
// each QFP physical entity.
// 
// An agent creates a conceptual row to this table corresponding
// to a QFP physical entity and its supported memory resource type
// upon detection of a physical entity supporting the memory 
// resource statistics for a memory resource type.
// 
// An agent destroys a conceptual row from this table       
// corresponding to a QFP physical entity and its supported
// memory resource type upon removal of the QFP host physical
// entity or it does not receive memory resource statistics
// update for a certain time period. The time period to wait
// before deleting an entry from this table would be the
// discretion of the supporting device.
type CISCOENTITYQFPMIB_CeqfpMemoryResourceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpMemoryResourceTable. There is an entry in this
    // table for each QFP entity by a value  of entPhysicalIndex and the supported
    // memory resource type  by a value of ceqfpMemoryResType. The type is slice
    // of CISCOENTITYQFPMIB_CeqfpMemoryResourceTable_CeqfpMemoryResourceEntry.
    CeqfpMemoryResourceEntry []*CISCOENTITYQFPMIB_CeqfpMemoryResourceTable_CeqfpMemoryResourceEntry
}

func (ceqfpMemoryResourceTable *CISCOENTITYQFPMIB_CeqfpMemoryResourceTable) GetEntityData() *types.CommonEntityData {
    ceqfpMemoryResourceTable.EntityData.YFilter = ceqfpMemoryResourceTable.YFilter
    ceqfpMemoryResourceTable.EntityData.YangName = "ceqfpMemoryResourceTable"
    ceqfpMemoryResourceTable.EntityData.BundleName = "cisco_ios_xe"
    ceqfpMemoryResourceTable.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ceqfpMemoryResourceTable.EntityData.SegmentPath = "ceqfpMemoryResourceTable"
    ceqfpMemoryResourceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpMemoryResourceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpMemoryResourceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpMemoryResourceTable.EntityData.Children = types.NewOrderedMap()
    ceqfpMemoryResourceTable.EntityData.Children.Append("ceqfpMemoryResourceEntry", types.YChild{"CeqfpMemoryResourceEntry", nil})
    for i := range ceqfpMemoryResourceTable.CeqfpMemoryResourceEntry {
        ceqfpMemoryResourceTable.EntityData.Children.Append(types.GetSegmentPath(ceqfpMemoryResourceTable.CeqfpMemoryResourceEntry[i]), types.YChild{"CeqfpMemoryResourceEntry", ceqfpMemoryResourceTable.CeqfpMemoryResourceEntry[i]})
    }
    ceqfpMemoryResourceTable.EntityData.Leafs = types.NewOrderedMap()

    ceqfpMemoryResourceTable.EntityData.YListKeys = []string {}

    return &(ceqfpMemoryResourceTable.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpMemoryResourceTable_CeqfpMemoryResourceEntry
// A conceptual row in the ceqfpMemoryResourceTable. There
// is an entry in this table for each QFP entity by a value 
// of entPhysicalIndex and the supported memory resource type 
// by a value of ceqfpMemoryResType.
type CISCOENTITYQFPMIB_CeqfpMemoryResourceTable_CeqfpMemoryResourceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. This object indicates the type of the memory
    // resource used by the QFP. This object is one of the indices to uniquely
    // identify the QFP memory resource type. The type is CiscoQfpMemoryResource.
    CeqfpMemoryResType interface{}

    // This object represents total memory available on this memory resource. The
    // type is interface{} with range: 0..4294967295. Units are kilo bytes.
    CeqfpMemoryResTotal interface{}

    // This object represents the memory which is currently under use on this
    // memory resource. The type is interface{} with range: 0..4294967295. Units
    // are kilo bytes.
    CeqfpMemoryResInUse interface{}

    // This object represents the memory which is currently free on this memory
    // resource. The type is interface{} with range: 0..4294967295. Units are kilo
    // bytes.
    CeqfpMemoryResFree interface{}

    // This object represents lowest free water mark on this memory resource. The
    // type is interface{} with range: 0..4294967295. Units are kilo bytes.
    CeqfpMemoryResLowFreeWatermark interface{}

    // This object represents the rising threshold value for this memory resource.
    // A value of zero means that the rising threshold is not supported for this
    // memory resource.  The value of this object can not be set to lower than or
    // equal to ceqfpMemoryResFallingThreshold.  A rising
    // (ceqfpMemoryResRisingThreshNotif) notification will be generated, whenever
    // the memory resource usage (ceqfpMemoryHCResInUse) is equal to or greater
    // than this value.  After a rising notification is generated, another such 
    // notification will not be generated until the  ceqfpMemoryResInUse falls
    // below this value and reaches  the ceqfpMemoryResFallingThreshold. The type
    // is interface{} with range: 0..100. Units are percent.
    CeqfpMemoryResRisingThreshold interface{}

    // This object represents the falling threshold value for this memory
    // resource. A value of zero means that the falling threshold is not supported
    // for this memory resource.  The value of this object can not be set to
    // higher than or equal to ceqfpMemoryResRisingThreshold.  A falling
    // (ceqfpMemoryResRisingThreshNotif) notification  will be generated, whenever
    // the memory resource usage (ceqfpMemoryHCResInUse) is equal to or lesser
    // than this value.  After a falling notification is generated, another  such
    // notification will not be generated until the  ceqfpMemoryResInUse rises
    // above this value and reaches  the ceqfpMemoryResRisingThreshold. The type
    // is interface{} with range: 0..100. Units are percent.
    CeqfpMemoryResFallingThreshold interface{}

    // This object represents the upper 32-bit of ceqfpMemoryResTotal. This object
    // needs to be supported only when the value of ceqfpMemoryResTotal exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo bytes.
    CeqfpMemoryResTotalOvrflw interface{}

    // This object is a 64-bit version of ceqfpMemoryResTotal. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo bytes.
    CeqfpMemoryHCResTotal interface{}

    // This object represents the upper 32-bit of ceqfpMemoryResInUse. This object
    // needs to be supported only when the value of ceqfpMemoryResInUse exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo bytes.
    CeqfpMemoryResInUseOvrflw interface{}

    // This object is a 64-bit version of ceqfpMemoryInRes. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo bytes.
    CeqfpMemoryHCResInUse interface{}

    // This object represents the upper 32-bit of ceqfpMemoryResFree. This object
    // needs to be supported only when the value of ceqfpMemoryResFree exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo bytes.
    CeqfpMemoryResFreeOvrflw interface{}

    // This object is a 64-bit version of ceqfpMemoryResFree. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo bytes.
    CeqfpMemoryHCResFree interface{}

    // This object represents the upper 32-bit of ceqfpMemoryResLowFreeWatermark.
    // This object needs to be supported only when the value of
    // ceqfpMemoryResLowFreeWatermark exceeds 32-bit, otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo bytes.
    CeqfpMemoryResLowFreeWatermarkOvrflw interface{}

    // This object is a 64-bit version of ceqfpMemoryResLowFreeWatermark. The type
    // is interface{} with range: 0..18446744073709551615. Units are kilo bytes.
    CeqfpMemoryHCResLowFreeWatermark interface{}
}

func (ceqfpMemoryResourceEntry *CISCOENTITYQFPMIB_CeqfpMemoryResourceTable_CeqfpMemoryResourceEntry) GetEntityData() *types.CommonEntityData {
    ceqfpMemoryResourceEntry.EntityData.YFilter = ceqfpMemoryResourceEntry.YFilter
    ceqfpMemoryResourceEntry.EntityData.YangName = "ceqfpMemoryResourceEntry"
    ceqfpMemoryResourceEntry.EntityData.BundleName = "cisco_ios_xe"
    ceqfpMemoryResourceEntry.EntityData.ParentYangName = "ceqfpMemoryResourceTable"
    ceqfpMemoryResourceEntry.EntityData.SegmentPath = "ceqfpMemoryResourceEntry" + types.AddKeyToken(ceqfpMemoryResourceEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(ceqfpMemoryResourceEntry.CeqfpMemoryResType, "ceqfpMemoryResType")
    ceqfpMemoryResourceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpMemoryResourceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpMemoryResourceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpMemoryResourceEntry.EntityData.Children = types.NewOrderedMap()
    ceqfpMemoryResourceEntry.EntityData.Leafs = types.NewOrderedMap()
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ceqfpMemoryResourceEntry.EntPhysicalIndex})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResType", types.YLeaf{"CeqfpMemoryResType", ceqfpMemoryResourceEntry.CeqfpMemoryResType})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResTotal", types.YLeaf{"CeqfpMemoryResTotal", ceqfpMemoryResourceEntry.CeqfpMemoryResTotal})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResInUse", types.YLeaf{"CeqfpMemoryResInUse", ceqfpMemoryResourceEntry.CeqfpMemoryResInUse})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResFree", types.YLeaf{"CeqfpMemoryResFree", ceqfpMemoryResourceEntry.CeqfpMemoryResFree})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResLowFreeWatermark", types.YLeaf{"CeqfpMemoryResLowFreeWatermark", ceqfpMemoryResourceEntry.CeqfpMemoryResLowFreeWatermark})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResRisingThreshold", types.YLeaf{"CeqfpMemoryResRisingThreshold", ceqfpMemoryResourceEntry.CeqfpMemoryResRisingThreshold})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResFallingThreshold", types.YLeaf{"CeqfpMemoryResFallingThreshold", ceqfpMemoryResourceEntry.CeqfpMemoryResFallingThreshold})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResTotalOvrflw", types.YLeaf{"CeqfpMemoryResTotalOvrflw", ceqfpMemoryResourceEntry.CeqfpMemoryResTotalOvrflw})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryHCResTotal", types.YLeaf{"CeqfpMemoryHCResTotal", ceqfpMemoryResourceEntry.CeqfpMemoryHCResTotal})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResInUseOvrflw", types.YLeaf{"CeqfpMemoryResInUseOvrflw", ceqfpMemoryResourceEntry.CeqfpMemoryResInUseOvrflw})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryHCResInUse", types.YLeaf{"CeqfpMemoryHCResInUse", ceqfpMemoryResourceEntry.CeqfpMemoryHCResInUse})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResFreeOvrflw", types.YLeaf{"CeqfpMemoryResFreeOvrflw", ceqfpMemoryResourceEntry.CeqfpMemoryResFreeOvrflw})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryHCResFree", types.YLeaf{"CeqfpMemoryHCResFree", ceqfpMemoryResourceEntry.CeqfpMemoryHCResFree})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryResLowFreeWatermarkOvrflw", types.YLeaf{"CeqfpMemoryResLowFreeWatermarkOvrflw", ceqfpMemoryResourceEntry.CeqfpMemoryResLowFreeWatermarkOvrflw})
    ceqfpMemoryResourceEntry.EntityData.Leafs.Append("ceqfpMemoryHCResLowFreeWatermark", types.YLeaf{"CeqfpMemoryHCResLowFreeWatermark", ceqfpMemoryResourceEntry.CeqfpMemoryHCResLowFreeWatermark})

    ceqfpMemoryResourceEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CeqfpMemoryResType"}

    return &(ceqfpMemoryResourceEntry.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpThroughputTable
// This table maintains the throughput information for each
// QFP physical entity.
// 
//         An agent creates a conceptual row to this table
// corresponding to a QFP physical entity upon detection of a
// physical entity supporting the QFP throughput information.
// 
//         An agent destroys a conceptual row from this table     
// 
// corresponding to a QFP physical entity upon removal of the QFP
// host physical entity.
type CISCOENTITYQFPMIB_CeqfpThroughputTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpThroughputTable. There is an entry in this
    // table for each QFP entity, as defined by a value of entPhysicalIndex. The
    // type is slice of
    // CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry.
    CeqfpThroughputEntry []*CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry
}

func (ceqfpThroughputTable *CISCOENTITYQFPMIB_CeqfpThroughputTable) GetEntityData() *types.CommonEntityData {
    ceqfpThroughputTable.EntityData.YFilter = ceqfpThroughputTable.YFilter
    ceqfpThroughputTable.EntityData.YangName = "ceqfpThroughputTable"
    ceqfpThroughputTable.EntityData.BundleName = "cisco_ios_xe"
    ceqfpThroughputTable.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ceqfpThroughputTable.EntityData.SegmentPath = "ceqfpThroughputTable"
    ceqfpThroughputTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpThroughputTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpThroughputTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpThroughputTable.EntityData.Children = types.NewOrderedMap()
    ceqfpThroughputTable.EntityData.Children.Append("ceqfpThroughputEntry", types.YChild{"CeqfpThroughputEntry", nil})
    for i := range ceqfpThroughputTable.CeqfpThroughputEntry {
        ceqfpThroughputTable.EntityData.Children.Append(types.GetSegmentPath(ceqfpThroughputTable.CeqfpThroughputEntry[i]), types.YChild{"CeqfpThroughputEntry", ceqfpThroughputTable.CeqfpThroughputEntry[i]})
    }
    ceqfpThroughputTable.EntityData.Leafs = types.NewOrderedMap()

    ceqfpThroughputTable.EntityData.YListKeys = []string {}

    return &(ceqfpThroughputTable.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry
// A conceptual row in the ceqfpThroughputTable. There is an entry
// in this table for each QFP entity, as defined by a value of
// entPhysicalIndex.
type CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This object represents the bandwidth for installed throughput license. The
    // type is interface{} with range: 0..18446744073709551615. Units are bits per
    // second.
    CeqfpThroughputLicensedBW interface{}

    // This object represents the current throughput level for installed
    // throughput license.                  normal  (1) - Throughput usage is
    // normal                 warning (2) - Throughput usage has crossed the      
    // configured threshold limit                 exceed  (3) - Throughput usage
    // has exceeded the                               total licensed bandwidth.
    // The type is CeqfpThroughputLevel.
    CeqfpThroughputLevel interface{}

    // The object represents the configured time interval at which the
    // ceqfpThroughputLevel is checked. The type is interface{} with range:
    // 10..86400. Units are seconds.
    CeqfpThroughputInterval interface{}

    // The object represents the configured throughput threshold. The type is
    // interface{} with range: 75..95. Units are percent.
    CeqfpThroughputThreshold interface{}

    // The object represents the average throughput rate in the interval
    // ceqfpThroughputInterval. The type is interface{} with range:
    // 0..18446744073709551615. Units are bits per second.
    CeqfpThroughputAvgRate interface{}
}

func (ceqfpThroughputEntry *CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry) GetEntityData() *types.CommonEntityData {
    ceqfpThroughputEntry.EntityData.YFilter = ceqfpThroughputEntry.YFilter
    ceqfpThroughputEntry.EntityData.YangName = "ceqfpThroughputEntry"
    ceqfpThroughputEntry.EntityData.BundleName = "cisco_ios_xe"
    ceqfpThroughputEntry.EntityData.ParentYangName = "ceqfpThroughputTable"
    ceqfpThroughputEntry.EntityData.SegmentPath = "ceqfpThroughputEntry" + types.AddKeyToken(ceqfpThroughputEntry.EntPhysicalIndex, "entPhysicalIndex")
    ceqfpThroughputEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpThroughputEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpThroughputEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpThroughputEntry.EntityData.Children = types.NewOrderedMap()
    ceqfpThroughputEntry.EntityData.Leafs = types.NewOrderedMap()
    ceqfpThroughputEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ceqfpThroughputEntry.EntPhysicalIndex})
    ceqfpThroughputEntry.EntityData.Leafs.Append("ceqfpThroughputLicensedBW", types.YLeaf{"CeqfpThroughputLicensedBW", ceqfpThroughputEntry.CeqfpThroughputLicensedBW})
    ceqfpThroughputEntry.EntityData.Leafs.Append("ceqfpThroughputLevel", types.YLeaf{"CeqfpThroughputLevel", ceqfpThroughputEntry.CeqfpThroughputLevel})
    ceqfpThroughputEntry.EntityData.Leafs.Append("ceqfpThroughputInterval", types.YLeaf{"CeqfpThroughputInterval", ceqfpThroughputEntry.CeqfpThroughputInterval})
    ceqfpThroughputEntry.EntityData.Leafs.Append("ceqfpThroughputThreshold", types.YLeaf{"CeqfpThroughputThreshold", ceqfpThroughputEntry.CeqfpThroughputThreshold})
    ceqfpThroughputEntry.EntityData.Leafs.Append("ceqfpThroughputAvgRate", types.YLeaf{"CeqfpThroughputAvgRate", ceqfpThroughputEntry.CeqfpThroughputAvgRate})

    ceqfpThroughputEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(ceqfpThroughputEntry.EntityData)
}

// CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry_CeqfpThroughputLevel represents                               total licensed bandwidth
type CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry_CeqfpThroughputLevel string

const (
    CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry_CeqfpThroughputLevel_normal CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry_CeqfpThroughputLevel = "normal"

    CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry_CeqfpThroughputLevel_warning CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry_CeqfpThroughputLevel = "warning"

    CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry_CeqfpThroughputLevel_exceed CISCOENTITYQFPMIB_CeqfpThroughputTable_CeqfpThroughputEntry_CeqfpThroughputLevel = "exceed"
)

