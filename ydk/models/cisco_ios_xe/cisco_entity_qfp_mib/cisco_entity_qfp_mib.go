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
    parent types.Entity
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

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetFilter() yfilter.YFilter { return cISCOENTITYQFPMIB.YFilter }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) SetFilter(yf yfilter.YFilter) { cISCOENTITYQFPMIB.YFilter = yf }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetGoName(yname string) string {
    if yname == "ciscoEntityQfp" { return "Ciscoentityqfp" }
    if yname == "ciscoEntityQfpNotif" { return "Ciscoentityqfpnotif" }
    if yname == "ceqfpSystemTable" { return "Ceqfpsystemtable" }
    if yname == "ceqfpUtilizationTable" { return "Ceqfputilizationtable" }
    if yname == "ceqfpMemoryResourceTable" { return "Ceqfpmemoryresourcetable" }
    if yname == "ceqfpThroughputTable" { return "Ceqfpthroughputtable" }
    return ""
}

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetSegmentPath() string {
    return "CISCO-ENTITY-QFP-MIB:CISCO-ENTITY-QFP-MIB"
}

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoEntityQfp" {
        return &cISCOENTITYQFPMIB.Ciscoentityqfp
    }
    if childYangName == "ciscoEntityQfpNotif" {
        return &cISCOENTITYQFPMIB.Ciscoentityqfpnotif
    }
    if childYangName == "ceqfpSystemTable" {
        return &cISCOENTITYQFPMIB.Ceqfpsystemtable
    }
    if childYangName == "ceqfpUtilizationTable" {
        return &cISCOENTITYQFPMIB.Ceqfputilizationtable
    }
    if childYangName == "ceqfpMemoryResourceTable" {
        return &cISCOENTITYQFPMIB.Ceqfpmemoryresourcetable
    }
    if childYangName == "ceqfpThroughputTable" {
        return &cISCOENTITYQFPMIB.Ceqfpthroughputtable
    }
    return nil
}

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoEntityQfp"] = &cISCOENTITYQFPMIB.Ciscoentityqfp
    children["ciscoEntityQfpNotif"] = &cISCOENTITYQFPMIB.Ciscoentityqfpnotif
    children["ceqfpSystemTable"] = &cISCOENTITYQFPMIB.Ceqfpsystemtable
    children["ceqfpUtilizationTable"] = &cISCOENTITYQFPMIB.Ceqfputilizationtable
    children["ceqfpMemoryResourceTable"] = &cISCOENTITYQFPMIB.Ceqfpmemoryresourcetable
    children["ceqfpThroughputTable"] = &cISCOENTITYQFPMIB.Ceqfpthroughputtable
    return children
}

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetYangName() string { return "CISCO-ENTITY-QFP-MIB" }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) SetParent(parent types.Entity) { cISCOENTITYQFPMIB.parent = parent }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetParent() types.Entity { return cISCOENTITYQFPMIB.parent }

func (cISCOENTITYQFPMIB *CISCOENTITYQFPMIB) GetParentYangName() string { return "CISCO-ENTITY-QFP-MIB" }

// CISCOENTITYQFPMIB_Ciscoentityqfp
type CISCOENTITYQFPMIB_Ciscoentityqfp struct {
    parent types.Entity
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

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetFilter() yfilter.YFilter { return ciscoentityqfp.YFilter }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) SetFilter(yf yfilter.YFilter) { ciscoentityqfp.YFilter = yf }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetGoName(yname string) string {
    if yname == "ceqfpFiveSecondUtilAlgo" { return "Ceqfpfivesecondutilalgo" }
    if yname == "ceqfpOneMinuteUtilAlgo" { return "Ceqfponeminuteutilalgo" }
    if yname == "ceqfpFiveMinutesUtilAlgo" { return "Ceqfpfiveminutesutilalgo" }
    if yname == "ceqfpSixtyMinutesUtilAlgo" { return "Ceqfpsixtyminutesutilalgo" }
    return ""
}

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetSegmentPath() string {
    return "ciscoEntityQfp"
}

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ceqfpFiveSecondUtilAlgo"] = ciscoentityqfp.Ceqfpfivesecondutilalgo
    leafs["ceqfpOneMinuteUtilAlgo"] = ciscoentityqfp.Ceqfponeminuteutilalgo
    leafs["ceqfpFiveMinutesUtilAlgo"] = ciscoentityqfp.Ceqfpfiveminutesutilalgo
    leafs["ceqfpSixtyMinutesUtilAlgo"] = ciscoentityqfp.Ceqfpsixtyminutesutilalgo
    return leafs
}

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetYangName() string { return "ciscoEntityQfp" }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) SetParent(parent types.Entity) { ciscoentityqfp.parent = parent }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetParent() types.Entity { return ciscoentityqfp.parent }

func (ciscoentityqfp *CISCOENTITYQFPMIB_Ciscoentityqfp) GetParentYangName() string { return "CISCO-ENTITY-QFP-MIB" }

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
    parent types.Entity
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

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetFilter() yfilter.YFilter { return ciscoentityqfpnotif.YFilter }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) SetFilter(yf yfilter.YFilter) { ciscoentityqfpnotif.YFilter = yf }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetGoName(yname string) string {
    if yname == "ceqfpMemoryResThreshNotifEnabled" { return "Ceqfpmemoryresthreshnotifenabled" }
    if yname == "ceqfpThroughputNotifEnabled" { return "Ceqfpthroughputnotifenabled" }
    return ""
}

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetSegmentPath() string {
    return "ciscoEntityQfpNotif"
}

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ceqfpMemoryResThreshNotifEnabled"] = ciscoentityqfpnotif.Ceqfpmemoryresthreshnotifenabled
    leafs["ceqfpThroughputNotifEnabled"] = ciscoentityqfpnotif.Ceqfpthroughputnotifenabled
    return leafs
}

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetYangName() string { return "ciscoEntityQfpNotif" }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) SetParent(parent types.Entity) { ciscoentityqfpnotif.parent = parent }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetParent() types.Entity { return ciscoentityqfpnotif.parent }

func (ciscoentityqfpnotif *CISCOENTITYQFPMIB_Ciscoentityqfpnotif) GetParentYangName() string { return "CISCO-ENTITY-QFP-MIB" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpSystemTable. There is an entry in this table
    // for each QFP entity, as defined by a value of entPhysicalIndex. The type is
    // slice of CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry.
    Ceqfpsystementry []CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry
}

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetFilter() yfilter.YFilter { return ceqfpsystemtable.YFilter }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) SetFilter(yf yfilter.YFilter) { ceqfpsystemtable.YFilter = yf }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetGoName(yname string) string {
    if yname == "ceqfpSystemEntry" { return "Ceqfpsystementry" }
    return ""
}

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetSegmentPath() string {
    return "ceqfpSystemTable"
}

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceqfpSystemEntry" {
        for _, c := range ceqfpsystemtable.Ceqfpsystementry {
            if ceqfpsystemtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry{}
        ceqfpsystemtable.Ceqfpsystementry = append(ceqfpsystemtable.Ceqfpsystementry, child)
        return &ceqfpsystemtable.Ceqfpsystementry[len(ceqfpsystemtable.Ceqfpsystementry)-1]
    }
    return nil
}

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceqfpsystemtable.Ceqfpsystementry {
        children[ceqfpsystemtable.Ceqfpsystementry[i].GetSegmentPath()] = &ceqfpsystemtable.Ceqfpsystementry[i]
    }
    return children
}

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetBundleName() string { return "cisco_ios_xe" }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetYangName() string { return "ceqfpSystemTable" }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) SetParent(parent types.Entity) { ceqfpsystemtable.parent = parent }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetParent() types.Entity { return ceqfpsystemtable.parent }

func (ceqfpsystemtable *CISCOENTITYQFPMIB_Ceqfpsystemtable) GetParentYangName() string { return "CISCO-ENTITY-QFP-MIB" }

// CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry
// A conceptual row in the ceqfpSystemTable. There is an entry
// in this table for each QFP entity, as defined by a value of
// entPhysicalIndex.
type CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry struct {
    parent types.Entity
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

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetFilter() yfilter.YFilter { return ceqfpsystementry.YFilter }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) SetFilter(yf yfilter.YFilter) { ceqfpsystementry.YFilter = yf }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "ceqfpSystemTrafficDirection" { return "Ceqfpsystemtrafficdirection" }
    if yname == "ceqfpSystemState" { return "Ceqfpsystemstate" }
    if yname == "ceqfpNumberSystemLoads" { return "Ceqfpnumbersystemloads" }
    if yname == "ceqfpSystemLastLoadTime" { return "Ceqfpsystemlastloadtime" }
    return ""
}

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetSegmentPath() string {
    return "ceqfpSystemEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceqfpsystementry.Entphysicalindex) + "']"
}

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceqfpsystementry.Entphysicalindex
    leafs["ceqfpSystemTrafficDirection"] = ceqfpsystementry.Ceqfpsystemtrafficdirection
    leafs["ceqfpSystemState"] = ceqfpsystementry.Ceqfpsystemstate
    leafs["ceqfpNumberSystemLoads"] = ceqfpsystementry.Ceqfpnumbersystemloads
    leafs["ceqfpSystemLastLoadTime"] = ceqfpsystementry.Ceqfpsystemlastloadtime
    return leafs
}

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetBundleName() string { return "cisco_ios_xe" }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetYangName() string { return "ceqfpSystemEntry" }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) SetParent(parent types.Entity) { ceqfpsystementry.parent = parent }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetParent() types.Entity { return ceqfpsystementry.parent }

func (ceqfpsystementry *CISCOENTITYQFPMIB_Ceqfpsystemtable_Ceqfpsystementry) GetParentYangName() string { return "ceqfpSystemTable" }

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
    parent types.Entity
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

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetFilter() yfilter.YFilter { return ceqfputilizationtable.YFilter }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) SetFilter(yf yfilter.YFilter) { ceqfputilizationtable.YFilter = yf }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetGoName(yname string) string {
    if yname == "ceqfpUtilizationEntry" { return "Ceqfputilizationentry" }
    return ""
}

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetSegmentPath() string {
    return "ceqfpUtilizationTable"
}

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceqfpUtilizationEntry" {
        for _, c := range ceqfputilizationtable.Ceqfputilizationentry {
            if ceqfputilizationtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry{}
        ceqfputilizationtable.Ceqfputilizationentry = append(ceqfputilizationtable.Ceqfputilizationentry, child)
        return &ceqfputilizationtable.Ceqfputilizationentry[len(ceqfputilizationtable.Ceqfputilizationentry)-1]
    }
    return nil
}

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceqfputilizationtable.Ceqfputilizationentry {
        children[ceqfputilizationtable.Ceqfputilizationentry[i].GetSegmentPath()] = &ceqfputilizationtable.Ceqfputilizationentry[i]
    }
    return children
}

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetBundleName() string { return "cisco_ios_xe" }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetYangName() string { return "ceqfpUtilizationTable" }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) SetParent(parent types.Entity) { ceqfputilizationtable.parent = parent }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetParent() types.Entity { return ceqfputilizationtable.parent }

func (ceqfputilizationtable *CISCOENTITYQFPMIB_Ceqfputilizationtable) GetParentYangName() string { return "CISCO-ENTITY-QFP-MIB" }

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
    parent types.Entity
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

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetFilter() yfilter.YFilter { return ceqfputilizationentry.YFilter }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) SetFilter(yf yfilter.YFilter) { ceqfputilizationentry.YFilter = yf }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "ceqfpUtilTimeInterval" { return "Ceqfputiltimeinterval" }
    if yname == "ceqfpUtilInputPriorityPktRate" { return "Ceqfputilinputprioritypktrate" }
    if yname == "ceqfpUtilInputPriorityBitRate" { return "Ceqfputilinputprioritybitrate" }
    if yname == "ceqfpUtilInputNonPriorityPktRate" { return "Ceqfputilinputnonprioritypktrate" }
    if yname == "ceqfpUtilInputNonPriorityBitRate" { return "Ceqfputilinputnonprioritybitrate" }
    if yname == "ceqfpUtilInputTotalPktRate" { return "Ceqfputilinputtotalpktrate" }
    if yname == "ceqfpUtilInputTotalBitRate" { return "Ceqfputilinputtotalbitrate" }
    if yname == "ceqfpUtilOutputPriorityPktRate" { return "Ceqfputiloutputprioritypktrate" }
    if yname == "ceqfpUtilOutputPriorityBitRate" { return "Ceqfputiloutputprioritybitrate" }
    if yname == "ceqfpUtilOutputNonPriorityPktRate" { return "Ceqfputiloutputnonprioritypktrate" }
    if yname == "ceqfpUtilOutputNonPriorityBitRate" { return "Ceqfputiloutputnonprioritybitrate" }
    if yname == "ceqfpUtilOutputTotalPktRate" { return "Ceqfputiloutputtotalpktrate" }
    if yname == "ceqfpUtilOutputTotalBitRate" { return "Ceqfputiloutputtotalbitrate" }
    if yname == "ceqfpUtilProcessingLoad" { return "Ceqfputilprocessingload" }
    return ""
}

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetSegmentPath() string {
    return "ceqfpUtilizationEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceqfputilizationentry.Entphysicalindex) + "']" + "[ceqfpUtilTimeInterval='" + fmt.Sprintf("%v", ceqfputilizationentry.Ceqfputiltimeinterval) + "']"
}

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceqfputilizationentry.Entphysicalindex
    leafs["ceqfpUtilTimeInterval"] = ceqfputilizationentry.Ceqfputiltimeinterval
    leafs["ceqfpUtilInputPriorityPktRate"] = ceqfputilizationentry.Ceqfputilinputprioritypktrate
    leafs["ceqfpUtilInputPriorityBitRate"] = ceqfputilizationentry.Ceqfputilinputprioritybitrate
    leafs["ceqfpUtilInputNonPriorityPktRate"] = ceqfputilizationentry.Ceqfputilinputnonprioritypktrate
    leafs["ceqfpUtilInputNonPriorityBitRate"] = ceqfputilizationentry.Ceqfputilinputnonprioritybitrate
    leafs["ceqfpUtilInputTotalPktRate"] = ceqfputilizationentry.Ceqfputilinputtotalpktrate
    leafs["ceqfpUtilInputTotalBitRate"] = ceqfputilizationentry.Ceqfputilinputtotalbitrate
    leafs["ceqfpUtilOutputPriorityPktRate"] = ceqfputilizationentry.Ceqfputiloutputprioritypktrate
    leafs["ceqfpUtilOutputPriorityBitRate"] = ceqfputilizationentry.Ceqfputiloutputprioritybitrate
    leafs["ceqfpUtilOutputNonPriorityPktRate"] = ceqfputilizationentry.Ceqfputiloutputnonprioritypktrate
    leafs["ceqfpUtilOutputNonPriorityBitRate"] = ceqfputilizationentry.Ceqfputiloutputnonprioritybitrate
    leafs["ceqfpUtilOutputTotalPktRate"] = ceqfputilizationentry.Ceqfputiloutputtotalpktrate
    leafs["ceqfpUtilOutputTotalBitRate"] = ceqfputilizationentry.Ceqfputiloutputtotalbitrate
    leafs["ceqfpUtilProcessingLoad"] = ceqfputilizationentry.Ceqfputilprocessingload
    return leafs
}

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetYangName() string { return "ceqfpUtilizationEntry" }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) SetParent(parent types.Entity) { ceqfputilizationentry.parent = parent }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetParent() types.Entity { return ceqfputilizationentry.parent }

func (ceqfputilizationentry *CISCOENTITYQFPMIB_Ceqfputilizationtable_Ceqfputilizationentry) GetParentYangName() string { return "ceqfpUtilizationTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpMemoryResourceTable. There is an entry in this
    // table for each QFP entity by a value  of entPhysicalIndex and the supported
    // memory resource type  by a value of ceqfpMemoryResType. The type is slice
    // of CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry.
    Ceqfpmemoryresourceentry []CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry
}

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetFilter() yfilter.YFilter { return ceqfpmemoryresourcetable.YFilter }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) SetFilter(yf yfilter.YFilter) { ceqfpmemoryresourcetable.YFilter = yf }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetGoName(yname string) string {
    if yname == "ceqfpMemoryResourceEntry" { return "Ceqfpmemoryresourceentry" }
    return ""
}

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetSegmentPath() string {
    return "ceqfpMemoryResourceTable"
}

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceqfpMemoryResourceEntry" {
        for _, c := range ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry {
            if ceqfpmemoryresourcetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry{}
        ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry = append(ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry, child)
        return &ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry[len(ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry)-1]
    }
    return nil
}

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry {
        children[ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry[i].GetSegmentPath()] = &ceqfpmemoryresourcetable.Ceqfpmemoryresourceentry[i]
    }
    return children
}

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetBundleName() string { return "cisco_ios_xe" }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetYangName() string { return "ceqfpMemoryResourceTable" }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) SetParent(parent types.Entity) { ceqfpmemoryresourcetable.parent = parent }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetParent() types.Entity { return ceqfpmemoryresourcetable.parent }

func (ceqfpmemoryresourcetable *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable) GetParentYangName() string { return "CISCO-ENTITY-QFP-MIB" }

// CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry
// A conceptual row in the ceqfpMemoryResourceTable. There
// is an entry in this table for each QFP entity by a value 
// of entPhysicalIndex and the supported memory resource type 
// by a value of ceqfpMemoryResType.
type CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry struct {
    parent types.Entity
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

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetFilter() yfilter.YFilter { return ceqfpmemoryresourceentry.YFilter }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) SetFilter(yf yfilter.YFilter) { ceqfpmemoryresourceentry.YFilter = yf }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "ceqfpMemoryResType" { return "Ceqfpmemoryrestype" }
    if yname == "ceqfpMemoryResTotal" { return "Ceqfpmemoryrestotal" }
    if yname == "ceqfpMemoryResInUse" { return "Ceqfpmemoryresinuse" }
    if yname == "ceqfpMemoryResFree" { return "Ceqfpmemoryresfree" }
    if yname == "ceqfpMemoryResLowFreeWatermark" { return "Ceqfpmemoryreslowfreewatermark" }
    if yname == "ceqfpMemoryResRisingThreshold" { return "Ceqfpmemoryresrisingthreshold" }
    if yname == "ceqfpMemoryResFallingThreshold" { return "Ceqfpmemoryresfallingthreshold" }
    if yname == "ceqfpMemoryResTotalOvrflw" { return "Ceqfpmemoryrestotalovrflw" }
    if yname == "ceqfpMemoryHCResTotal" { return "Ceqfpmemoryhcrestotal" }
    if yname == "ceqfpMemoryResInUseOvrflw" { return "Ceqfpmemoryresinuseovrflw" }
    if yname == "ceqfpMemoryHCResInUse" { return "Ceqfpmemoryhcresinuse" }
    if yname == "ceqfpMemoryResFreeOvrflw" { return "Ceqfpmemoryresfreeovrflw" }
    if yname == "ceqfpMemoryHCResFree" { return "Ceqfpmemoryhcresfree" }
    if yname == "ceqfpMemoryResLowFreeWatermarkOvrflw" { return "Ceqfpmemoryreslowfreewatermarkovrflw" }
    if yname == "ceqfpMemoryHCResLowFreeWatermark" { return "Ceqfpmemoryhcreslowfreewatermark" }
    return ""
}

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetSegmentPath() string {
    return "ceqfpMemoryResourceEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceqfpmemoryresourceentry.Entphysicalindex) + "']" + "[ceqfpMemoryResType='" + fmt.Sprintf("%v", ceqfpmemoryresourceentry.Ceqfpmemoryrestype) + "']"
}

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceqfpmemoryresourceentry.Entphysicalindex
    leafs["ceqfpMemoryResType"] = ceqfpmemoryresourceentry.Ceqfpmemoryrestype
    leafs["ceqfpMemoryResTotal"] = ceqfpmemoryresourceentry.Ceqfpmemoryrestotal
    leafs["ceqfpMemoryResInUse"] = ceqfpmemoryresourceentry.Ceqfpmemoryresinuse
    leafs["ceqfpMemoryResFree"] = ceqfpmemoryresourceentry.Ceqfpmemoryresfree
    leafs["ceqfpMemoryResLowFreeWatermark"] = ceqfpmemoryresourceentry.Ceqfpmemoryreslowfreewatermark
    leafs["ceqfpMemoryResRisingThreshold"] = ceqfpmemoryresourceentry.Ceqfpmemoryresrisingthreshold
    leafs["ceqfpMemoryResFallingThreshold"] = ceqfpmemoryresourceentry.Ceqfpmemoryresfallingthreshold
    leafs["ceqfpMemoryResTotalOvrflw"] = ceqfpmemoryresourceentry.Ceqfpmemoryrestotalovrflw
    leafs["ceqfpMemoryHCResTotal"] = ceqfpmemoryresourceentry.Ceqfpmemoryhcrestotal
    leafs["ceqfpMemoryResInUseOvrflw"] = ceqfpmemoryresourceentry.Ceqfpmemoryresinuseovrflw
    leafs["ceqfpMemoryHCResInUse"] = ceqfpmemoryresourceentry.Ceqfpmemoryhcresinuse
    leafs["ceqfpMemoryResFreeOvrflw"] = ceqfpmemoryresourceentry.Ceqfpmemoryresfreeovrflw
    leafs["ceqfpMemoryHCResFree"] = ceqfpmemoryresourceentry.Ceqfpmemoryhcresfree
    leafs["ceqfpMemoryResLowFreeWatermarkOvrflw"] = ceqfpmemoryresourceentry.Ceqfpmemoryreslowfreewatermarkovrflw
    leafs["ceqfpMemoryHCResLowFreeWatermark"] = ceqfpmemoryresourceentry.Ceqfpmemoryhcreslowfreewatermark
    return leafs
}

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetYangName() string { return "ceqfpMemoryResourceEntry" }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) SetParent(parent types.Entity) { ceqfpmemoryresourceentry.parent = parent }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetParent() types.Entity { return ceqfpmemoryresourceentry.parent }

func (ceqfpmemoryresourceentry *CISCOENTITYQFPMIB_Ceqfpmemoryresourcetable_Ceqfpmemoryresourceentry) GetParentYangName() string { return "ceqfpMemoryResourceTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // A conceptual row in the ceqfpThroughputTable. There is an entry in this
    // table for each QFP entity, as defined by a value of entPhysicalIndex. The
    // type is slice of
    // CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry.
    Ceqfpthroughputentry []CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry
}

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetFilter() yfilter.YFilter { return ceqfpthroughputtable.YFilter }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) SetFilter(yf yfilter.YFilter) { ceqfpthroughputtable.YFilter = yf }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetGoName(yname string) string {
    if yname == "ceqfpThroughputEntry" { return "Ceqfpthroughputentry" }
    return ""
}

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetSegmentPath() string {
    return "ceqfpThroughputTable"
}

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceqfpThroughputEntry" {
        for _, c := range ceqfpthroughputtable.Ceqfpthroughputentry {
            if ceqfpthroughputtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry{}
        ceqfpthroughputtable.Ceqfpthroughputentry = append(ceqfpthroughputtable.Ceqfpthroughputentry, child)
        return &ceqfpthroughputtable.Ceqfpthroughputentry[len(ceqfpthroughputtable.Ceqfpthroughputentry)-1]
    }
    return nil
}

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceqfpthroughputtable.Ceqfpthroughputentry {
        children[ceqfpthroughputtable.Ceqfpthroughputentry[i].GetSegmentPath()] = &ceqfpthroughputtable.Ceqfpthroughputentry[i]
    }
    return children
}

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetBundleName() string { return "cisco_ios_xe" }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetYangName() string { return "ceqfpThroughputTable" }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) SetParent(parent types.Entity) { ceqfpthroughputtable.parent = parent }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetParent() types.Entity { return ceqfpthroughputtable.parent }

func (ceqfpthroughputtable *CISCOENTITYQFPMIB_Ceqfpthroughputtable) GetParentYangName() string { return "CISCO-ENTITY-QFP-MIB" }

// CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry
// A conceptual row in the ceqfpThroughputTable. There is an entry
// in this table for each QFP entity, as defined by a value of
// entPhysicalIndex.
type CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry struct {
    parent types.Entity
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

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetFilter() yfilter.YFilter { return ceqfpthroughputentry.YFilter }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) SetFilter(yf yfilter.YFilter) { ceqfpthroughputentry.YFilter = yf }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "ceqfpThroughputLicensedBW" { return "Ceqfpthroughputlicensedbw" }
    if yname == "ceqfpThroughputLevel" { return "Ceqfpthroughputlevel" }
    if yname == "ceqfpThroughputInterval" { return "Ceqfpthroughputinterval" }
    if yname == "ceqfpThroughputThreshold" { return "Ceqfpthroughputthreshold" }
    if yname == "ceqfpThroughputAvgRate" { return "Ceqfpthroughputavgrate" }
    return ""
}

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetSegmentPath() string {
    return "ceqfpThroughputEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceqfpthroughputentry.Entphysicalindex) + "']"
}

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceqfpthroughputentry.Entphysicalindex
    leafs["ceqfpThroughputLicensedBW"] = ceqfpthroughputentry.Ceqfpthroughputlicensedbw
    leafs["ceqfpThroughputLevel"] = ceqfpthroughputentry.Ceqfpthroughputlevel
    leafs["ceqfpThroughputInterval"] = ceqfpthroughputentry.Ceqfpthroughputinterval
    leafs["ceqfpThroughputThreshold"] = ceqfpthroughputentry.Ceqfpthroughputthreshold
    leafs["ceqfpThroughputAvgRate"] = ceqfpthroughputentry.Ceqfpthroughputavgrate
    return leafs
}

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetYangName() string { return "ceqfpThroughputEntry" }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) SetParent(parent types.Entity) { ceqfpthroughputentry.parent = parent }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetParent() types.Entity { return ceqfpthroughputentry.parent }

func (ceqfpthroughputentry *CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry) GetParentYangName() string { return "ceqfpThroughputTable" }

// CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel represents                               total licensed bandwidth
type CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel string

const (
    CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel_normal CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel = "normal"

    CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel_warning CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel = "warning"

    CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel_exceed CISCOENTITYQFPMIB_Ceqfpthroughputtable_Ceqfpthroughputentry_Ceqfpthroughputlevel = "exceed"
)

