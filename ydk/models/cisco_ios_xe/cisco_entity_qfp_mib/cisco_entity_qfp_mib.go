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

    
    Ciscoentityqfp CISCOENTITYQFPMIB_Ciscoentityqfp

    
    Ciscoentityqfpnotif CISCOENTITYQFPMIB_Ciscoentityqfpnotif

    // This table maintains the QFP system information for each QFP physical
    // entity.  An agent creates a conceptual row to this table corresponding to a
    // QFP physical entity upon detection of a physical entity supporting the QFP
    // system information.  An agent destroys a conceptual row from this table    
    // corresponding to a QFP physical entity upon removal of the QFP host
    // physical entity.
    Ceqfpsystemtable CISCOENTITYQFPMIB_Ceqfpsystemtable

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
    Ceqfputilizationtable CISCOENTITYQFPMIB_Ceqfputilizationtable

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
    Ceqfpmemoryresourcetable CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable

    // This table maintains the throughput information for each QFP physical
    // entity.          An agent creates a conceptual row to this table
    // corresponding to a QFP physical entity upon detection of a physical entity
    // supporting the QFP throughput information.          An agent destroys a
    // conceptual row from this table       corresponding to a QFP physical entity
    // upon removal of the QFP host physical entity.
    Ceqfpthroughputtable CISCOENTITYQFPMIB_Ceqfpthroughputtable
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

    cISCOENTITYQFPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOENTITYQFPMIB.EntityData.Children["ciscoEntityQfp"] = types.YChild{"Ciscoentityqfp", &cISCOENTITYQFPMIB.Ciscoentityqfp}
    cISCOENTITYQFPMIB.EntityData.Children["ciscoEntityQfpNotif"] = types.YChild{"Ciscoentityqfpnotif", &cISCOENTITYQFPMIB.Ciscoentityqfpnotif}
    cISCOENTITYQFPMIB.EntityData.Children["ceqfpSystemTable"] = types.YChild{"Ceqfpsystemtable", &cISCOENTITYQFPMIB.Ceqfpsystemtable}
    cISCOENTITYQFPMIB.EntityData.Children["ceqfpUtilizationTable"] = types.YChild{"Ceqfputilizationtable", &cISCOENTITYQFPMIB.Ceqfputilizationtable}
    cISCOENTITYQFPMIB.EntityData.Children["ceqfpMemoryResourceTable"] = types.YChild{"Ceqfpmemoryresourcetable", &cISCOENTITYQFPMIB.Ceqfpmemoryresourcetable}
    cISCOENTITYQFPMIB.EntityData.Children["ceqfpThroughputTable"] = types.YChild{"Ceqfpthroughputtable", &cISCOENTITYQFPMIB.Ceqfpthroughputtable}
    cISCOENTITYQFPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOENTITYQFPMIB.EntityData)
}

// CISCOENTITYQFPMIB_Ciscoentityqfp
type CISCOENTITYQFPMIB_Ciscoentityqfp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This objects represents the method used to calculate 5 Second interval
    // utilization data. The enumerated values are described below.  unknown      
    // (1) - The calculation method is unknown fiveSecSample (2) - The value is
    // calculated based on the last                     5 second sampling period
    // of utilization                     data. The type is
    // Ceqfpfivesecondutilalgo.
    Ceqfpfivesecondutilalgo interface{}

    // This objects represents the method used to calculate 1 Minute interval
    // utilization data. The enumerated values are described below.  unknown   
    // (1) - The calculation method is unknown fiveSecSMA (2) - The value is
    // calculated using Simple Moving                    Average of last 12 five
    // seconds utilization                  data. The type is
    // Ceqfponeminuteutilalgo.
    Ceqfponeminuteutilalgo interface{}

    // This objects represents the method used to calculate 5 Minutes interval
    // utilization data. The enumerated values are described below.  unknown   
    // (1) - The calculation method is unknown fiveSecSMA (2) - The value is
    // calculated using Simple Moving                    Average of last 60 five
    // seconds utilization                  data. The type is
    // Ceqfpfiveminutesutilalgo.
    Ceqfpfiveminutesutilalgo interface{}

    // This objects represents the method used to calculate 60 Minutes interval
    // utilization data. The enumerated values are described below.  unknown   
    // (1) - The calculation method is unknown fiveSecSMA (1) - The value is
    // calculated using Simple Moving                    Average of last 720 five
    // seconds utilization                  data. The type is
    // Ceqfpsixtyminutesutilalgo.
    Ceqfpsixtyminutesutilalgo interface{}
}

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetEntityData() *types.CommonEntityData {
    ciscoentityqfp.EntityData.YFilter = ciscoentityqfp.YFilter
    ciscoentityqfp.EntityData.YangName = "ciscoEntityQfp"
    ciscoentityqfp.EntityData.BundleName = "cisco_ios_xe"
    ciscoentityqfp.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ciscoentityqfp.EntityData.SegmentPath = "ciscoEntityQfp"
    ciscoentityqfp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoentityqfp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoentityqfp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoentityqfp.EntityData.Children = make(map[string]types.YChild)
    ciscoentityqfp.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoentityqfp.EntityData.Leafs["ceqfpFiveSecondUtilAlgo"] = types.YLeaf{"Ceqfpfivesecondutilalgo", ciscoentityqfp.Ceqfpfivesecondutilalgo}
    ciscoentityqfp.EntityData.Leafs["ceqfpOneMinuteUtilAlgo"] = types.YLeaf{"Ceqfponeminuteutilalgo", ciscoentityqfp.Ceqfponeminuteutilalgo}
    ciscoentityqfp.EntityData.Leafs["ceqfpFiveMinutesUtilAlgo"] = types.YLeaf{"Ceqfpfiveminutesutilalgo", ciscoentityqfp.Ceqfpfiveminutesutilalgo}
    ciscoentityqfp.EntityData.Leafs["ceqfpSixtyMinutesUtilAlgo"] = types.YLeaf{"Ceqfpsixtyminutesutilalgo", ciscoentityqfp.Ceqfpsixtyminutesutilalgo}
    return &(ciscoentityqfp.EntityData)
}

// CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfiveminutesutilalgo represents                  data.
type CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfiveminutesutilalgo string

const (
    CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfiveminutesutilalgo_unknown CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfiveminutesutilalgo = "unknown"

    CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfiveminutesutilalgo_fiveSecSMA CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfiveminutesutilalgo = "fiveSecSMA"
)

// CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfivesecondutilalgo represents                     data.
type CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfivesecondutilalgo string

const (
    CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfivesecondutilalgo_unknown CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfivesecondutilalgo = "unknown"

    CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfivesecondutilalgo_fiveSecSample CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpfivesecondutilalgo = "fiveSecSample"
)

// CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfponeminuteutilalgo represents                  data.
type CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfponeminuteutilalgo string

const (
    CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfponeminuteutilalgo_unknown CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfponeminuteutilalgo = "unknown"

    CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfponeminuteutilalgo_fiveSecSMA CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfponeminuteutilalgo = "fiveSecSMA"
)

// CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpsixtyminutesutilalgo represents                  data.
type CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpsixtyminutesutilalgo string

const (
    CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpsixtyminutesutilalgo_unknown CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpsixtyminutesutilalgo = "unknown"

    CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpsixtyminutesutilalgo_fiveSecSMA CISCOENTITYQFPMIB_Ciscoentityqfp_Ceqfpsixtyminutesutilalgo = "fiveSecSMA"
)

// CISCOENTITYQFPMIB_Ciscoentityqfpnotif
type CISCOENTITYQFPMIB_Ciscoentityqfpnotif struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object controls memory resource rising and falling threshold
    // notification.  When this object contains a value 'true', then generation of
    // memory resource threshold notification is enabled. If this object contains
    // a value 'false', then generation of memory resource threshold notification
    // is disabled. The type is bool.
    Ceqfpmemoryresthreshnotifenabled interface{}

    // This object controls throughput rate notification.  When this object
    // contains a value 'true', then generation of ceqfpThroughputNotif is
    // enabled. If this object contains a value 'false', then generation of
    // ceqfpThroughputNotif is disabled. The type is interface{} with range: 0..2.
    Ceqfpthroughputnotifenabled interface{}
}

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetEntityData() *types.CommonEntityData {
    ciscoentityqfpnotif.EntityData.YFilter = ciscoentityqfpnotif.YFilter
    ciscoentityqfpnotif.EntityData.YangName = "ciscoEntityQfpNotif"
    ciscoentityqfpnotif.EntityData.BundleName = "cisco_ios_xe"
    ciscoentityqfpnotif.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ciscoentityqfpnotif.EntityData.SegmentPath = "ciscoEntityQfpNotif"
    ciscoentityqfpnotif.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoentityqfpnotif.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoentityqfpnotif.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoentityqfpnotif.EntityData.Children = make(map[string]types.YChild)
    ciscoentityqfpnotif.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoentityqfpnotif.EntityData.Leafs["ceqfpMemoryResThreshNotifEnabled"] = types.YLeaf{"Ceqfpmemoryresthreshnotifenabled", ciscoentityqfpnotif.Ceqfpmemoryresthreshnotifenabled}
    ciscoentityqfpnotif.EntityData.Leafs["ceqfpThroughputNotifEnabled"] = types.YLeaf{"Ceqfpthroughputnotifenabled", ciscoentityqfpnotif.Ceqfpthroughputnotifenabled}
    return &(ciscoentityqfpnotif.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfpsystemtable
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
type CISCOENTITYQFPMIB_Ceqfpsystemtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpSystemTable. There is an entry in this table
    // for each QFP entity, as defined by a value of entPhysicalIndex. The type is
    // slice of CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry.
    Ceqfpsystementry []CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry
}

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetEntityData() *types.CommonEntityData {
    ceqfpsystemtable.EntityData.YFilter = ceqfpsystemtable.YFilter
    ceqfpsystemtable.EntityData.YangName = "ceqfpSystemTable"
    ceqfpsystemtable.EntityData.BundleName = "cisco_ios_xe"
    ceqfpsystemtable.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ceqfpsystemtable.EntityData.SegmentPath = "ceqfpSystemTable"
    ceqfpsystemtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpsystemtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpsystemtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpsystemtable.EntityData.Children = make(map[string]types.YChild)
    ceqfpsystemtable.EntityData.Children["ceqfpSystemEntry"] = types.YChild{"Ceqfpsystementry", nil}
    for i := range ceqfpsystemtable.Ceqfpsystementry {
        ceqfpsystemtable.EntityData.Children[types.GetSegmentPath(&ceqfpsystemtable.Ceqfpsystementry[i])] = types.YChild{"Ceqfpsystementry", &ceqfpsystemtable.Ceqfpsystementry[i]}
    }
    ceqfpsystemtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceqfpsystemtable.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry
// A conceptual row in the ceqfpSystemTable. There is an entry
// in this table for each QFP entity, as defined by a value of
// entPhysicalIndex.
type CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This object represents the traffic direction that this QFP is assigned to
    // process. The enumerated values are described below.  none (1)    - The QFP
    // is not assigned to processes any traffic               yet ingress (2) -
    // The QFP processes inbound traffic egress (3)  - The QFP processes outbound
    // traffic both (4)    - The QFP processes both inbound and outbound          
    // traffic. The type is Ceqfpsystemtrafficdirection.
    Ceqfpsystemtrafficdirection interface{}

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
    // type is Ceqfpsystemstate.
    Ceqfpsystemstate interface{}

    // This object represents the number of times the QFP is loaded, since the QFP
    // host is up. The type is interface{} with range: 0..4294967295.
    Ceqfpnumbersystemloads interface{}

    // This object represents the QFP last load time. The type is string.
    Ceqfpsystemlastloadtime interface{}
}

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetEntityData() *types.CommonEntityData {
    ceqfpsystementry.EntityData.YFilter = ceqfpsystementry.YFilter
    ceqfpsystementry.EntityData.YangName = "ceqfpSystemEntry"
    ceqfpsystementry.EntityData.BundleName = "cisco_ios_xe"
    ceqfpsystementry.EntityData.ParentYangName = "ceqfpSystemTable"
    ceqfpsystementry.EntityData.SegmentPath = "ceqfpSystemEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceqfpsystementry.Entphysicalindex) + "']"
    ceqfpsystementry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpsystementry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpsystementry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpsystementry.EntityData.Children = make(map[string]types.YChild)
    ceqfpsystementry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceqfpsystementry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceqfpsystementry.Entphysicalindex}
    ceqfpsystementry.EntityData.Leafs["ceqfpSystemTrafficDirection"] = types.YLeaf{"Ceqfpsystemtrafficdirection", ceqfpsystementry.Ceqfpsystemtrafficdirection}
    ceqfpsystementry.EntityData.Leafs["ceqfpSystemState"] = types.YLeaf{"Ceqfpsystemstate", ceqfpsystementry.Ceqfpsystemstate}
    ceqfpsystementry.EntityData.Leafs["ceqfpNumberSystemLoads"] = types.YLeaf{"Ceqfpnumbersystemloads", ceqfpsystementry.Ceqfpnumbersystemloads}
    ceqfpsystementry.EntityData.Leafs["ceqfpSystemLastLoadTime"] = types.YLeaf{"Ceqfpsystemlastloadtime", ceqfpsystementry.Ceqfpsystemlastloadtime}
    return &(ceqfpsystementry.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate represents                  active QFP and standby QFP
type CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate string

const (
    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate_unknown CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate = "unknown"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate_reset CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate = "reset"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate_init CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate = "init"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate_active CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate = "active"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate_activeSolo CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate = "activeSolo"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate_standby CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate = "standby"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate_hotStandby CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemstate = "hotStandby"
)

// CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection represents               traffic
type CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection string

const (
    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection_none CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection = "none"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection_ingress CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection = "ingress"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection_egress CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection = "egress"

    CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection_both CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry_Ceqfpsystemtrafficdirection = "both"
)

// CISCOENTITYQFPMIB_Ceqfputilizationtable
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
type CISCOENTITYQFPMIB_Ceqfputilizationtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpUtilizationTable. There is an entry in this
    // table for each QFP entity by a value of entPhysicalIndex and the supported
    // time interval by a value  of ceqfpUtilTimeInterval.  The method of
    // utilization data calculation for each interval period can be identified
    // through the respective interval scalar objects. For example the utilizaiton
    // data calculation method for 'fiveSecond' interval can be identified by
    // ceqfpFiveSecondUtilAlgo. The type is slice of
    // CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry.
    Ceqfputilizationentry []CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry
}

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetEntityData() *types.CommonEntityData {
    ceqfputilizationtable.EntityData.YFilter = ceqfputilizationtable.YFilter
    ceqfputilizationtable.EntityData.YangName = "ceqfpUtilizationTable"
    ceqfputilizationtable.EntityData.BundleName = "cisco_ios_xe"
    ceqfputilizationtable.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ceqfputilizationtable.EntityData.SegmentPath = "ceqfpUtilizationTable"
    ceqfputilizationtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfputilizationtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfputilizationtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfputilizationtable.EntityData.Children = make(map[string]types.YChild)
    ceqfputilizationtable.EntityData.Children["ceqfpUtilizationEntry"] = types.YChild{"Ceqfputilizationentry", nil}
    for i := range ceqfputilizationtable.Ceqfputilizationentry {
        ceqfputilizationtable.EntityData.Children[types.GetSegmentPath(&ceqfputilizationtable.Ceqfputilizationentry[i])] = types.YChild{"Ceqfputilizationentry", &ceqfputilizationtable.Ceqfputilizationentry[i]}
    }
    ceqfputilizationtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceqfputilizationtable.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry
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
type CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. This object identifies the time interval for which
    // the utilization statistics being collected. The interval  values can be 5
    // second, 1 minute, etc. as specified in  the CiscoQfpTimeInterval. The type
    // is CiscoQfpTimeInterval.
    Ceqfputiltimeinterval interface{}

    // This object represents the QFP input channel priority packet rate during
    // this interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are packets per second.
    Ceqfputilinputprioritypktrate interface{}

    // This object represents the QFP input channel priority bit rate during this
    // interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits per second.
    Ceqfputilinputprioritybitrate interface{}

    // This object represents the QFP input channel non-priority packet rate
    // during this interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets per second.
    Ceqfputilinputnonprioritypktrate interface{}

    // This object represents the QFP input channel non-priority bit rate during
    // this interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits per second.
    Ceqfputilinputnonprioritybitrate interface{}

    // This object represents the QFP input channel total packet rate during this
    // interval, which includes both priority and non-priority input packet rate.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets per second.
    Ceqfputilinputtotalpktrate interface{}

    // This object represents the QFP input channel total bit rate during this
    // interval, which includes both priority and non-priority input bit rate. The
    // type is interface{} with range: 0..18446744073709551615. Units are bits per
    // second.
    Ceqfputilinputtotalbitrate interface{}

    // This object represents the QFP output channel priority packet rate during
    // this interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are packets per second.
    Ceqfputiloutputprioritypktrate interface{}

    // This object represents the QFP output channel priority bit rate during this
    // interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits per second.
    Ceqfputiloutputprioritybitrate interface{}

    // This object represents the QFP output channel non-priority packet rate
    // during this interval. The type is interface{} with range:
    // 0..18446744073709551615. Units are packets per second.
    Ceqfputiloutputnonprioritypktrate interface{}

    // This object represents the QFP output channel non-priority bit rate during
    // this interval. The type is interface{} with range: 0..18446744073709551615.
    // Units are bits per second.
    Ceqfputiloutputnonprioritybitrate interface{}

    // This object represents the QFP output channel total packet rate during this
    // interval, which includes both priority and non-priority output packet rate.
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // packets per second.
    Ceqfputiloutputtotalpktrate interface{}

    // This object represents the QFP output channel total bit rate during this
    // interval, which includes both priority and non-priority bit rate. The type
    // is interface{} with range: 0..18446744073709551615. Units are bits per
    // second.
    Ceqfputiloutputtotalbitrate interface{}

    // This object represents the QFP processing load during this interval. The
    // type is interface{} with range: 0..100. Units are percent.
    Ceqfputilprocessingload interface{}
}

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetEntityData() *types.CommonEntityData {
    ceqfputilizationentry.EntityData.YFilter = ceqfputilizationentry.YFilter
    ceqfputilizationentry.EntityData.YangName = "ceqfpUtilizationEntry"
    ceqfputilizationentry.EntityData.BundleName = "cisco_ios_xe"
    ceqfputilizationentry.EntityData.ParentYangName = "ceqfpUtilizationTable"
    ceqfputilizationentry.EntityData.SegmentPath = "ceqfpUtilizationEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceqfputilizationentry.Entphysicalindex) + "']" + "[ceqfpUtilTimeInterval='" + fmt.Sprintf("%v", ceqfputilizationentry.Ceqfputiltimeinterval) + "']"
    ceqfputilizationentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfputilizationentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfputilizationentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfputilizationentry.EntityData.Children = make(map[string]types.YChild)
    ceqfputilizationentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceqfputilizationentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceqfputilizationentry.Entphysicalindex}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilTimeInterval"] = types.YLeaf{"Ceqfputiltimeinterval", ceqfputilizationentry.Ceqfputiltimeinterval}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilInputPriorityPktRate"] = types.YLeaf{"Ceqfputilinputprioritypktrate", ceqfputilizationentry.Ceqfputilinputprioritypktrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilInputPriorityBitRate"] = types.YLeaf{"Ceqfputilinputprioritybitrate", ceqfputilizationentry.Ceqfputilinputprioritybitrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilInputNonPriorityPktRate"] = types.YLeaf{"Ceqfputilinputnonprioritypktrate", ceqfputilizationentry.Ceqfputilinputnonprioritypktrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilInputNonPriorityBitRate"] = types.YLeaf{"Ceqfputilinputnonprioritybitrate", ceqfputilizationentry.Ceqfputilinputnonprioritybitrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilInputTotalPktRate"] = types.YLeaf{"Ceqfputilinputtotalpktrate", ceqfputilizationentry.Ceqfputilinputtotalpktrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilInputTotalBitRate"] = types.YLeaf{"Ceqfputilinputtotalbitrate", ceqfputilizationentry.Ceqfputilinputtotalbitrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilOutputPriorityPktRate"] = types.YLeaf{"Ceqfputiloutputprioritypktrate", ceqfputilizationentry.Ceqfputiloutputprioritypktrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilOutputPriorityBitRate"] = types.YLeaf{"Ceqfputiloutputprioritybitrate", ceqfputilizationentry.Ceqfputiloutputprioritybitrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilOutputNonPriorityPktRate"] = types.YLeaf{"Ceqfputiloutputnonprioritypktrate", ceqfputilizationentry.Ceqfputiloutputnonprioritypktrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilOutputNonPriorityBitRate"] = types.YLeaf{"Ceqfputiloutputnonprioritybitrate", ceqfputilizationentry.Ceqfputiloutputnonprioritybitrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilOutputTotalPktRate"] = types.YLeaf{"Ceqfputiloutputtotalpktrate", ceqfputilizationentry.Ceqfputiloutputtotalpktrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilOutputTotalBitRate"] = types.YLeaf{"Ceqfputiloutputtotalbitrate", ceqfputilizationentry.Ceqfputiloutputtotalbitrate}
    ceqfputilizationentry.EntityData.Leafs["ceqfpUtilProcessingLoad"] = types.YLeaf{"Ceqfputilprocessingload", ceqfputilizationentry.Ceqfputilprocessingload}
    return &(ceqfputilizationentry.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable
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
type CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpMemoryResourceTable. There is an entry in this
    // table for each QFP entity by a value  of entPhysicalIndex and the supported
    // memory resource type  by a value of ceqfpMemoryResType. The type is slice
    // of CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry.
    Ceqfpmemoryresourceentry []CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry
}

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetEntityData() *types.CommonEntityData {
    ceqfpmemoryresourcetable.EntityData.YFilter = ceqfpmemoryresourcetable.YFilter
    ceqfpmemoryresourcetable.EntityData.YangName = "ceqfpMemoryResourceTable"
    ceqfpmemoryresourcetable.EntityData.BundleName = "cisco_ios_xe"
    ceqfpmemoryresourcetable.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ceqfpmemoryresourcetable.EntityData.SegmentPath = "ceqfpMemoryResourceTable"
    ceqfpmemoryresourcetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpmemoryresourcetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpmemoryresourcetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpmemoryresourcetable.EntityData.Children = make(map[string]types.YChild)
    ceqfpmemoryresourcetable.EntityData.Children["ceqfpMemoryResourceEntry"] = types.YChild{"Ceqfpmemoryresourceentry", nil}
    for i := range ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry {
        ceqfpmemoryresourcetable.EntityData.Children[types.GetSegmentPath(&ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry[i])] = types.YChild{"Ceqfpmemoryresourceentry", &ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry[i]}
    }
    ceqfpmemoryresourcetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceqfpmemoryresourcetable.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry
// A conceptual row in the ceqfpMemoryResourceTable. There
// is an entry in this table for each QFP entity by a value 
// of entPhysicalIndex and the supported memory resource type 
// by a value of ceqfpMemoryResType.
type CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. This object indicates the type of the memory
    // resource used by the QFP. This object is one of the indices to uniquely
    // identify the QFP memory resource type. The type is CiscoQfpMemoryResource.
    Ceqfpmemoryrestype interface{}

    // This object represents total memory available on this memory resource. The
    // type is interface{} with range: 0..4294967295. Units are kilo bytes.
    Ceqfpmemoryrestotal interface{}

    // This object represents the memory which is currently under use on this
    // memory resource. The type is interface{} with range: 0..4294967295. Units
    // are kilo bytes.
    Ceqfpmemoryresinuse interface{}

    // This object represents the memory which is currently free on this memory
    // resource. The type is interface{} with range: 0..4294967295. Units are kilo
    // bytes.
    Ceqfpmemoryresfree interface{}

    // This object represents lowest free water mark on this memory resource. The
    // type is interface{} with range: 0..4294967295. Units are kilo bytes.
    Ceqfpmemoryreslowfreewatermark interface{}

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
    Ceqfpmemoryresrisingthreshold interface{}

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
    Ceqfpmemoryresfallingthreshold interface{}

    // This object represents the upper 32-bit of ceqfpMemoryResTotal. This object
    // needs to be supported only when the value of ceqfpMemoryResTotal exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo bytes.
    Ceqfpmemoryrestotalovrflw interface{}

    // This object is a 64-bit version of ceqfpMemoryResTotal. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo bytes.
    Ceqfpmemoryhcrestotal interface{}

    // This object represents the upper 32-bit of ceqfpMemoryResInUse. This object
    // needs to be supported only when the value of ceqfpMemoryResInUse exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo bytes.
    Ceqfpmemoryresinuseovrflw interface{}

    // This object is a 64-bit version of ceqfpMemoryInRes. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo bytes.
    Ceqfpmemoryhcresinuse interface{}

    // This object represents the upper 32-bit of ceqfpMemoryResFree. This object
    // needs to be supported only when the value of ceqfpMemoryResFree exceeds
    // 32-bit, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are kilo bytes.
    Ceqfpmemoryresfreeovrflw interface{}

    // This object is a 64-bit version of ceqfpMemoryResFree. The type is
    // interface{} with range: 0..18446744073709551615. Units are kilo bytes.
    Ceqfpmemoryhcresfree interface{}

    // This object represents the upper 32-bit of ceqfpMemoryResLowFreeWatermark.
    // This object needs to be supported only when the value of
    // ceqfpMemoryResLowFreeWatermark exceeds 32-bit, otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are kilo bytes.
    Ceqfpmemoryreslowfreewatermarkovrflw interface{}

    // This object is a 64-bit version of ceqfpMemoryResLowFreeWatermark. The type
    // is interface{} with range: 0..18446744073709551615. Units are kilo bytes.
    Ceqfpmemoryhcreslowfreewatermark interface{}
}

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetEntityData() *types.CommonEntityData {
    ceqfpmemoryresourceentry.EntityData.YFilter = ceqfpmemoryresourceentry.YFilter
    ceqfpmemoryresourceentry.EntityData.YangName = "ceqfpMemoryResourceEntry"
    ceqfpmemoryresourceentry.EntityData.BundleName = "cisco_ios_xe"
    ceqfpmemoryresourceentry.EntityData.ParentYangName = "ceqfpMemoryResourceTable"
    ceqfpmemoryresourceentry.EntityData.SegmentPath = "ceqfpMemoryResourceEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceqfpmemoryresourceentry.Entphysicalindex) + "']" + "[ceqfpMemoryResType='" + fmt.Sprintf("%v", ceqfpmemoryresourceentry.Ceqfpmemoryrestype) + "']"
    ceqfpmemoryresourceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpmemoryresourceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpmemoryresourceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpmemoryresourceentry.EntityData.Children = make(map[string]types.YChild)
    ceqfpmemoryresourceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceqfpmemoryresourceentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceqfpmemoryresourceentry.Entphysicalindex}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResType"] = types.YLeaf{"Ceqfpmemoryrestype", ceqfpmemoryresourceentry.Ceqfpmemoryrestype}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResTotal"] = types.YLeaf{"Ceqfpmemoryrestotal", ceqfpmemoryresourceentry.Ceqfpmemoryrestotal}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResInUse"] = types.YLeaf{"Ceqfpmemoryresinuse", ceqfpmemoryresourceentry.Ceqfpmemoryresinuse}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResFree"] = types.YLeaf{"Ceqfpmemoryresfree", ceqfpmemoryresourceentry.Ceqfpmemoryresfree}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResLowFreeWatermark"] = types.YLeaf{"Ceqfpmemoryreslowfreewatermark", ceqfpmemoryresourceentry.Ceqfpmemoryreslowfreewatermark}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResRisingThreshold"] = types.YLeaf{"Ceqfpmemoryresrisingthreshold", ceqfpmemoryresourceentry.Ceqfpmemoryresrisingthreshold}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResFallingThreshold"] = types.YLeaf{"Ceqfpmemoryresfallingthreshold", ceqfpmemoryresourceentry.Ceqfpmemoryresfallingthreshold}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResTotalOvrflw"] = types.YLeaf{"Ceqfpmemoryrestotalovrflw", ceqfpmemoryresourceentry.Ceqfpmemoryrestotalovrflw}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryHCResTotal"] = types.YLeaf{"Ceqfpmemoryhcrestotal", ceqfpmemoryresourceentry.Ceqfpmemoryhcrestotal}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResInUseOvrflw"] = types.YLeaf{"Ceqfpmemoryresinuseovrflw", ceqfpmemoryresourceentry.Ceqfpmemoryresinuseovrflw}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryHCResInUse"] = types.YLeaf{"Ceqfpmemoryhcresinuse", ceqfpmemoryresourceentry.Ceqfpmemoryhcresinuse}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResFreeOvrflw"] = types.YLeaf{"Ceqfpmemoryresfreeovrflw", ceqfpmemoryresourceentry.Ceqfpmemoryresfreeovrflw}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryHCResFree"] = types.YLeaf{"Ceqfpmemoryhcresfree", ceqfpmemoryresourceentry.Ceqfpmemoryhcresfree}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryResLowFreeWatermarkOvrflw"] = types.YLeaf{"Ceqfpmemoryreslowfreewatermarkovrflw", ceqfpmemoryresourceentry.Ceqfpmemoryreslowfreewatermarkovrflw}
    ceqfpmemoryresourceentry.EntityData.Leafs["ceqfpMemoryHCResLowFreeWatermark"] = types.YLeaf{"Ceqfpmemoryhcreslowfreewatermark", ceqfpmemoryresourceentry.Ceqfpmemoryhcreslowfreewatermark}
    return &(ceqfpmemoryresourceentry.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfpthroughputtable
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
type CISCOENTITYQFPMIB_Ceqfpthroughputtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpThroughputTable. There is an entry in this
    // table for each QFP entity, as defined by a value of entPhysicalIndex. The
    // type is slice of
    // CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry.
    Ceqfpthroughputentry []CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry
}

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetEntityData() *types.CommonEntityData {
    ceqfpthroughputtable.EntityData.YFilter = ceqfpthroughputtable.YFilter
    ceqfpthroughputtable.EntityData.YangName = "ceqfpThroughputTable"
    ceqfpthroughputtable.EntityData.BundleName = "cisco_ios_xe"
    ceqfpthroughputtable.EntityData.ParentYangName = "CISCO-ENTITY-QFP-MIB"
    ceqfpthroughputtable.EntityData.SegmentPath = "ceqfpThroughputTable"
    ceqfpthroughputtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpthroughputtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpthroughputtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpthroughputtable.EntityData.Children = make(map[string]types.YChild)
    ceqfpthroughputtable.EntityData.Children["ceqfpThroughputEntry"] = types.YChild{"Ceqfpthroughputentry", nil}
    for i := range ceqfpthroughputtable.Ceqfpthroughputentry {
        ceqfpthroughputtable.EntityData.Children[types.GetSegmentPath(&ceqfpthroughputtable.Ceqfpthroughputentry[i])] = types.YChild{"Ceqfpthroughputentry", &ceqfpthroughputtable.Ceqfpthroughputentry[i]}
    }
    ceqfpthroughputtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceqfpthroughputtable.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry
// A conceptual row in the ceqfpThroughputTable. There is an entry
// in this table for each QFP entity, as defined by a value of
// entPhysicalIndex.
type CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This object represents the bandwidth for installed throughput license. The
    // type is interface{} with range: 0..18446744073709551615. Units are bits per
    // second.
    Ceqfpthroughputlicensedbw interface{}

    // This object represents the current throughput level for installed
    // throughput license.                  normal  (1) - Throughput usage is
    // normal                 warning (2) - Throughput usage has crossed the      
    // configured threshold limit                 exceed  (3) - Throughput usage
    // has exceeded the                               total licensed bandwidth.
    // The type is Ceqfpthroughputlevel.
    Ceqfpthroughputlevel interface{}

    // The object represents the configured time interval at which the
    // ceqfpThroughputLevel is checked. The type is interface{} with range:
    // 10..86400. Units are seconds.
    Ceqfpthroughputinterval interface{}

    // The object represents the configured throughput threshold. The type is
    // interface{} with range: 75..95. Units are percent.
    Ceqfpthroughputthreshold interface{}

    // The object represents the average throughput rate in the interval
    // ceqfpThroughputInterval. The type is interface{} with range:
    // 0..18446744073709551615. Units are bits per second.
    Ceqfpthroughputavgrate interface{}
}

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetEntityData() *types.CommonEntityData {
    ceqfpthroughputentry.EntityData.YFilter = ceqfpthroughputentry.YFilter
    ceqfpthroughputentry.EntityData.YangName = "ceqfpThroughputEntry"
    ceqfpthroughputentry.EntityData.BundleName = "cisco_ios_xe"
    ceqfpthroughputentry.EntityData.ParentYangName = "ceqfpThroughputTable"
    ceqfpthroughputentry.EntityData.SegmentPath = "ceqfpThroughputEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceqfpthroughputentry.Entphysicalindex) + "']"
    ceqfpthroughputentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceqfpthroughputentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceqfpthroughputentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceqfpthroughputentry.EntityData.Children = make(map[string]types.YChild)
    ceqfpthroughputentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceqfpthroughputentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceqfpthroughputentry.Entphysicalindex}
    ceqfpthroughputentry.EntityData.Leafs["ceqfpThroughputLicensedBW"] = types.YLeaf{"Ceqfpthroughputlicensedbw", ceqfpthroughputentry.Ceqfpthroughputlicensedbw}
    ceqfpthroughputentry.EntityData.Leafs["ceqfpThroughputLevel"] = types.YLeaf{"Ceqfpthroughputlevel", ceqfpthroughputentry.Ceqfpthroughputlevel}
    ceqfpthroughputentry.EntityData.Leafs["ceqfpThroughputInterval"] = types.YLeaf{"Ceqfpthroughputinterval", ceqfpthroughputentry.Ceqfpthroughputinterval}
    ceqfpthroughputentry.EntityData.Leafs["ceqfpThroughputThreshold"] = types.YLeaf{"Ceqfpthroughputthreshold", ceqfpthroughputentry.Ceqfpthroughputthreshold}
    ceqfpthroughputentry.EntityData.Leafs["ceqfpThroughputAvgRate"] = types.YLeaf{"Ceqfpthroughputavgrate", ceqfpthroughputentry.Ceqfpthroughputavgrate}
    return &(ceqfpthroughputentry.EntityData)
}

// CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel represents                               total licensed bandwidth
type CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel string

const (
    CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel_normal CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel = "normal"

    CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel_warning CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel = "warning"

    CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel_exceed CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel = "exceed"
)

