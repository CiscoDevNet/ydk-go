// A MIB module for extending the POWER-ETHERNET-MIB
// (RFC3621) to add objects which provide additional
// management information about Power Sourcing Equipment
// (PSE) not available in POWER-ETHERNET-MIB.
// 
// Glossary
// 
// Power Sourcing Equipment (PSE)
//     These are devices supplying electrical power to
//     other equipment. They are, for example, inline power
//     switches, inline power daughterboards and power patch
//     panels.
// 
// Powered Device (PD)
//     These are devices receiving their electrical power
//     supply from Power Sourcing Equipment. They are, for
//     example, IP telephones and wireless access points
//     or bridges.
package cisco_power_ethernet_ext_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_power_ethernet_ext_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-POWER-ETHERNET-EXT-MIB CISCO-POWER-ETHERNET-EXT-MIB}", reflect.TypeOf(CISCOPOWERETHERNETEXTMIB{}))
    ydk.RegisterEntity("CISCO-POWER-ETHERNET-EXT-MIB:CISCO-POWER-ETHERNET-EXT-MIB", reflect.TypeOf(CISCOPOWERETHERNETEXTMIB{}))
}

// CpeExtLldpPwrType represents              or a Type 2 PD.
type CpeExtLldpPwrType string

const (
    CpeExtLldpPwrType_type1Pd CpeExtLldpPwrType = "type1Pd"

    CpeExtLldpPwrType_type1Pse CpeExtLldpPwrType = "type1Pse"

    CpeExtLldpPwrType_type2Pd CpeExtLldpPwrType = "type2Pd"

    CpeExtLldpPwrType_type2Pse CpeExtLldpPwrType = "type2Pse"
)

// CpeExtLldpPwrSrc represents                   unknown.
type CpeExtLldpPwrSrc string

const (
    CpeExtLldpPwrSrc_pseAndLocal CpeExtLldpPwrSrc = "pseAndLocal"

    CpeExtLldpPwrSrc_local CpeExtLldpPwrSrc = "local"

    CpeExtLldpPwrSrc_pse CpeExtLldpPwrSrc = "pse"

    CpeExtLldpPwrSrc_backupSrc CpeExtLldpPwrSrc = "backupSrc"

    CpeExtLldpPwrSrc_primarySrc CpeExtLldpPwrSrc = "primarySrc"

    CpeExtLldpPwrSrc_unknown CpeExtLldpPwrSrc = "unknown"
)

// CpeExtPwrPriority represents 'unknown'  - power priority level is unknown.
type CpeExtPwrPriority string

const (
    CpeExtPwrPriority_critical CpeExtPwrPriority = "critical"

    CpeExtPwrPriority_high CpeExtPwrPriority = "high"

    CpeExtPwrPriority_low CpeExtPwrPriority = "low"

    CpeExtPwrPriority_unknown CpeExtPwrPriority = "unknown"
)

// CISCOPOWERETHERNETEXTMIB
type CISCOPOWERETHERNETEXTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cpeextmibobjects CISCOPOWERETHERNETEXTMIB_Cpeextmibobjects

    
    Cpeextpdstatistics CISCOPOWERETHERNETEXTMIB_Cpeextpdstatistics

    // This table contains the additional information for the main PSE group in
    // pethMainPseTable.
    Cpeextmainpsetable CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable

    // This table contains the statistics information of the powered devices
    // fallen into different power classifications in the system.
    Cpeextpdstatstable CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable

    // A table provides the Link Layer Discovery Protocol (LLDP) based Data Link
    // Layer (DLL) power classification characteristics of PSE ports and PDs
    // attached to them.
    Cpeextpseportlldptable CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable
}

func (cISCOPOWERETHERNETEXTMIB *CISCOPOWERETHERNETEXTMIB) GetEntityData() *types.CommonEntityData {
    cISCOPOWERETHERNETEXTMIB.EntityData.YFilter = cISCOPOWERETHERNETEXTMIB.YFilter
    cISCOPOWERETHERNETEXTMIB.EntityData.YangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cISCOPOWERETHERNETEXTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOPOWERETHERNETEXTMIB.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cISCOPOWERETHERNETEXTMIB.EntityData.SegmentPath = "CISCO-POWER-ETHERNET-EXT-MIB:CISCO-POWER-ETHERNET-EXT-MIB"
    cISCOPOWERETHERNETEXTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOPOWERETHERNETEXTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOPOWERETHERNETEXTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOPOWERETHERNETEXTMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOPOWERETHERNETEXTMIB.EntityData.Children["cpeExtMIBObjects"] = types.YChild{"Cpeextmibobjects", &cISCOPOWERETHERNETEXTMIB.Cpeextmibobjects}
    cISCOPOWERETHERNETEXTMIB.EntityData.Children["cpeExtPdStatistics"] = types.YChild{"Cpeextpdstatistics", &cISCOPOWERETHERNETEXTMIB.Cpeextpdstatistics}
    cISCOPOWERETHERNETEXTMIB.EntityData.Children["cpeExtMainPseTable"] = types.YChild{"Cpeextmainpsetable", &cISCOPOWERETHERNETEXTMIB.Cpeextmainpsetable}
    cISCOPOWERETHERNETEXTMIB.EntityData.Children["cpeExtPdStatsTable"] = types.YChild{"Cpeextpdstatstable", &cISCOPOWERETHERNETEXTMIB.Cpeextpdstatstable}
    cISCOPOWERETHERNETEXTMIB.EntityData.Children["cpeExtPsePortLldpTable"] = types.YChild{"Cpeextpseportlldptable", &cISCOPOWERETHERNETEXTMIB.Cpeextpseportlldptable}
    cISCOPOWERETHERNETEXTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOPOWERETHERNETEXTMIB.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_Cpeextmibobjects
type CISCOPOWERETHERNETEXTMIB_Cpeextmibobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the default inline power allocation per port. This is
    // a global configuration parameter that applies to all inline power capable
    // ports in the system.  The system must consider this object as well as the
    // per port configuration object, cpeExtPsePortPwrMax, when determining how
    // much power to allocate to a port. The system will use the lower of the two
    // numbers. The type is interface{} with range: 0..4294967295. Units are
    // milliwatts.
    Cpeextdefaultallocation interface{}

    // This object is used to enable/disable the the cpeExtPolicingNotif
    // notification. The type is bool.
    Cpeextpolicingnotifenable interface{}

    // This object is the global control of the power priority feature on the
    // device. 'true' indicates that the power priority feature is globally
    // enabled. 'false' indicates that the power priority feature is globally
    // disabled. The type is bool.
    Cpeextpowerpriorityenable interface{}
}

func (cpeextmibobjects *CISCOPOWERETHERNETEXTMIB_Cpeextmibobjects) GetEntityData() *types.CommonEntityData {
    cpeextmibobjects.EntityData.YFilter = cpeextmibobjects.YFilter
    cpeextmibobjects.EntityData.YangName = "cpeExtMIBObjects"
    cpeextmibobjects.EntityData.BundleName = "cisco_ios_xe"
    cpeextmibobjects.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeextmibobjects.EntityData.SegmentPath = "cpeExtMIBObjects"
    cpeextmibobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeextmibobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeextmibobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeextmibobjects.EntityData.Children = make(map[string]types.YChild)
    cpeextmibobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    cpeextmibobjects.EntityData.Leafs["cpeExtDefaultAllocation"] = types.YLeaf{"Cpeextdefaultallocation", cpeextmibobjects.Cpeextdefaultallocation}
    cpeextmibobjects.EntityData.Leafs["cpeExtPolicingNotifEnable"] = types.YLeaf{"Cpeextpolicingnotifenable", cpeextmibobjects.Cpeextpolicingnotifenable}
    cpeextmibobjects.EntityData.Leafs["cpeExtPowerPriorityEnable"] = types.YLeaf{"Cpeextpowerpriorityenable", cpeextmibobjects.Cpeextpowerpriorityenable}
    return &(cpeextmibobjects.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_Cpeextpdstatistics
type CISCOPOWERETHERNETEXTMIB_Cpeextpdstatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the total number of the powered devices with any
    // power classifications in the system.  Classification is a way to tag
    // different terminals on the Power over LAN network according to their power
    // consumption. Devices such as IP telephones, WLAN access points and others,
    // will be classified according to their power requirements. The type is
    // interface{} with range: 0..4294967295.
    Cpeextpdstatstotaldevices interface{}
}

func (cpeextpdstatistics *CISCOPOWERETHERNETEXTMIB_Cpeextpdstatistics) GetEntityData() *types.CommonEntityData {
    cpeextpdstatistics.EntityData.YFilter = cpeextpdstatistics.YFilter
    cpeextpdstatistics.EntityData.YangName = "cpeExtPdStatistics"
    cpeextpdstatistics.EntityData.BundleName = "cisco_ios_xe"
    cpeextpdstatistics.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeextpdstatistics.EntityData.SegmentPath = "cpeExtPdStatistics"
    cpeextpdstatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeextpdstatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeextpdstatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeextpdstatistics.EntityData.Children = make(map[string]types.YChild)
    cpeextpdstatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    cpeextpdstatistics.EntityData.Leafs["cpeExtPdStatsTotalDevices"] = types.YLeaf{"Cpeextpdstatstotaldevices", cpeextpdstatistics.Cpeextpdstatstotaldevices}
    return &(cpeextpdstatistics.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable
// This table contains the additional information for the
// main PSE group in pethMainPseTable.
type CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cpeExtMainPseEntry contains information about a particular
    // pethMainPseGroupIndex. An entry is created by the agent when a main PSE
    // group is added to the pethMainPseTable. An entry is deleted by the agent
    // when a main PSE group is removed from the pethMainPseTable. The type is
    // slice of CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable_Cpeextmainpseentry.
    Cpeextmainpseentry []CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable_Cpeextmainpseentry
}

func (cpeextmainpsetable *CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable) GetEntityData() *types.CommonEntityData {
    cpeextmainpsetable.EntityData.YFilter = cpeextmainpsetable.YFilter
    cpeextmainpsetable.EntityData.YangName = "cpeExtMainPseTable"
    cpeextmainpsetable.EntityData.BundleName = "cisco_ios_xe"
    cpeextmainpsetable.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeextmainpsetable.EntityData.SegmentPath = "cpeExtMainPseTable"
    cpeextmainpsetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeextmainpsetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeextmainpsetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeextmainpsetable.EntityData.Children = make(map[string]types.YChild)
    cpeextmainpsetable.EntityData.Children["cpeExtMainPseEntry"] = types.YChild{"Cpeextmainpseentry", nil}
    for i := range cpeextmainpsetable.Cpeextmainpseentry {
        cpeextmainpsetable.EntityData.Children[types.GetSegmentPath(&cpeextmainpsetable.Cpeextmainpseentry[i])] = types.YChild{"Cpeextmainpseentry", &cpeextmainpsetable.Cpeextmainpseentry[i]}
    }
    cpeextmainpsetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpeextmainpsetable.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable_Cpeextmainpseentry
// A cpeExtMainPseEntry contains information about
// a particular pethMainPseGroupIndex. An entry is
// created by the agent when a main PSE group is added
// to the pethMainPseTable. An entry is deleted by the
// agent when a main PSE group is removed from the
// pethMainPseTable.
type CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable_Cpeextmainpseentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // power_ethernet_mib.POWERETHERNETMIB_Pethmainpsetable_Pethmainpseentry_Pethmainpsegroupindex
    Pethmainpsegroupindex interface{}

    // The entPhysicalIndex value that uniquely identifies the main PSE group. If
    // the main PSE group does not have a corresponding physical entry in
    // entPhysicalTable or if the entPhysicalTable is not supported by the
    // management system, then this object has the value of zero. The type is
    // interface{} with range: 0..2147483647.
    Cpeextmainpseentphyindex interface{}

    // A textual string containing information about the Power Source Equipment
    // (PSE) group. The type is string.
    Cpeextmainpsedescr interface{}

    // This object indicates if the given group is capable of monitoring the power
    // consumption of the interfaces that belong to the group. The value 'true'
    // means that the group is capable. The value 'false' means that the group in
    // not capable. The type is bool.
    Cpeextmainpsepwrmonitorcapable interface{}

    // Used power expressed in miliwatts. The type is interface{} with range:
    // 0..4294967295. Units are miliwatts.
    Cpeextmainpseusedpower interface{}

    // Remaining power expressed in miliwatts, this parameter is calculated as
    // pethMainPsePower minus cpeExtMainPseUsedPower. The type is interface{} with
    // range: 0..4294967295. Units are miliwatts.
    Cpeextmainpseremainingpower interface{}
}

func (cpeextmainpseentry *CISCOPOWERETHERNETEXTMIB_Cpeextmainpsetable_Cpeextmainpseentry) GetEntityData() *types.CommonEntityData {
    cpeextmainpseentry.EntityData.YFilter = cpeextmainpseentry.YFilter
    cpeextmainpseentry.EntityData.YangName = "cpeExtMainPseEntry"
    cpeextmainpseentry.EntityData.BundleName = "cisco_ios_xe"
    cpeextmainpseentry.EntityData.ParentYangName = "cpeExtMainPseTable"
    cpeextmainpseentry.EntityData.SegmentPath = "cpeExtMainPseEntry" + "[pethMainPseGroupIndex='" + fmt.Sprintf("%v", cpeextmainpseentry.Pethmainpsegroupindex) + "']"
    cpeextmainpseentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeextmainpseentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeextmainpseentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeextmainpseentry.EntityData.Children = make(map[string]types.YChild)
    cpeextmainpseentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpeextmainpseentry.EntityData.Leafs["pethMainPseGroupIndex"] = types.YLeaf{"Pethmainpsegroupindex", cpeextmainpseentry.Pethmainpsegroupindex}
    cpeextmainpseentry.EntityData.Leafs["cpeExtMainPseEntPhyIndex"] = types.YLeaf{"Cpeextmainpseentphyindex", cpeextmainpseentry.Cpeextmainpseentphyindex}
    cpeextmainpseentry.EntityData.Leafs["cpeExtMainPseDescr"] = types.YLeaf{"Cpeextmainpsedescr", cpeextmainpseentry.Cpeextmainpsedescr}
    cpeextmainpseentry.EntityData.Leafs["cpeExtMainPsePwrMonitorCapable"] = types.YLeaf{"Cpeextmainpsepwrmonitorcapable", cpeextmainpseentry.Cpeextmainpsepwrmonitorcapable}
    cpeextmainpseentry.EntityData.Leafs["cpeExtMainPseUsedPower"] = types.YLeaf{"Cpeextmainpseusedpower", cpeextmainpseentry.Cpeextmainpseusedpower}
    cpeextmainpseentry.EntityData.Leafs["cpeExtMainPseRemainingPower"] = types.YLeaf{"Cpeextmainpseremainingpower", cpeextmainpseentry.Cpeextmainpseremainingpower}
    return &(cpeextmainpseentry.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable
// This table contains the statistics information
// of the powered devices fallen into different power
// classifications in the system.
type CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cpeExtPdStatsEntry contains the statistics information about a particular
    // power classification defined in cpeExtPdStatsIndex. The type is slice of
    // CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry.
    Cpeextpdstatsentry []CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry
}

func (cpeextpdstatstable *CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable) GetEntityData() *types.CommonEntityData {
    cpeextpdstatstable.EntityData.YFilter = cpeextpdstatstable.YFilter
    cpeextpdstatstable.EntityData.YangName = "cpeExtPdStatsTable"
    cpeextpdstatstable.EntityData.BundleName = "cisco_ios_xe"
    cpeextpdstatstable.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeextpdstatstable.EntityData.SegmentPath = "cpeExtPdStatsTable"
    cpeextpdstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeextpdstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeextpdstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeextpdstatstable.EntityData.Children = make(map[string]types.YChild)
    cpeextpdstatstable.EntityData.Children["cpeExtPdStatsEntry"] = types.YChild{"Cpeextpdstatsentry", nil}
    for i := range cpeextpdstatstable.Cpeextpdstatsentry {
        cpeextpdstatstable.EntityData.Children[types.GetSegmentPath(&cpeextpdstatstable.Cpeextpdstatsentry[i])] = types.YChild{"Cpeextpdstatsentry", &cpeextpdstatstable.Cpeextpdstatsentry[i]}
    }
    cpeextpdstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpeextpdstatstable.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry
// A cpeExtPdStatsEntry contains the statistics
// information about a particular power classification
// defined in cpeExtPdStatsIndex.
type CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The power classification as the index for the
    // statistics information for powered devices.  A value of 'cisco' indicates
    // that the powered devices are CISCO proprietary and their power 
    // classification does not fall into any class in IEEE specifications.  A
    // value of 'class0' indicates that the power  classification of the powered
    // devices falls into class 0 in IEEE specifications.  A value of 'class1'
    // indicates that the power classification of the powered devices falls into
    // class 1 in IEEE specifications.  A value of 'class2' indicates that the
    // power classification of the powered devices falls into class 2 in IEEE
    // specifications.  A value of 'class3' indicates that the power
    // classification of the powered devices falls into class 3 in IEEE
    // specifications.  A value of 'class4' indicates that the power
    // classification of the powered devices falls into class 4 in IEEE
    // specifications. The type is Cpeextpdstatsclass.
    Cpeextpdstatsclass interface{}

    // This object indicates the count of the powered devices whose power
    // classification falls into  a specific value of cpeExtPdStatsIndex. The type
    // is interface{} with range: 0..4294967295.
    Cpeextpdstatsdevicecount interface{}
}

func (cpeextpdstatsentry *CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry) GetEntityData() *types.CommonEntityData {
    cpeextpdstatsentry.EntityData.YFilter = cpeextpdstatsentry.YFilter
    cpeextpdstatsentry.EntityData.YangName = "cpeExtPdStatsEntry"
    cpeextpdstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cpeextpdstatsentry.EntityData.ParentYangName = "cpeExtPdStatsTable"
    cpeextpdstatsentry.EntityData.SegmentPath = "cpeExtPdStatsEntry" + "[cpeExtPdStatsClass='" + fmt.Sprintf("%v", cpeextpdstatsentry.Cpeextpdstatsclass) + "']"
    cpeextpdstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeextpdstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeextpdstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeextpdstatsentry.EntityData.Children = make(map[string]types.YChild)
    cpeextpdstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpeextpdstatsentry.EntityData.Leafs["cpeExtPdStatsClass"] = types.YLeaf{"Cpeextpdstatsclass", cpeextpdstatsentry.Cpeextpdstatsclass}
    cpeextpdstatsentry.EntityData.Leafs["cpeExtPdStatsDeviceCount"] = types.YLeaf{"Cpeextpdstatsdevicecount", cpeextpdstatsentry.Cpeextpdstatsdevicecount}
    return &(cpeextpdstatsentry.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass represents class 4 in IEEE specifications.
type CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass string

const (
    CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass_cisco CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass = "cisco"

    CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass_class0 CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass = "class0"

    CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass_class1 CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass = "class1"

    CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass_class2 CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass = "class2"

    CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass_class3 CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass = "class3"

    CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass_class4 CISCOPOWERETHERNETEXTMIB_Cpeextpdstatstable_Cpeextpdstatsentry_Cpeextpdstatsclass = "class4"
)

// CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable
// A table provides the Link Layer Discovery Protocol (LLDP)
// based Data Link Layer (DLL) power classification
// characteristics of PSE ports and PDs attached to them.
type CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cpeExtPseDllPowerEntry entry contains the LLDP based DLL power
    // classification characteristics for a particular PSE and the PD attached to
    // it.   A PSE has its entry here when all of the following conditions are
    // satisfied: - The LLDP power classification is globally enabled. - It has a
    // PD attached. - LLDP is the operational power negotiation protocol between  
    // the PSE and the PD attached. The type is slice of
    // CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable_Cpeextpseportlldpentry.
    Cpeextpseportlldpentry []CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable_Cpeextpseportlldpentry
}

func (cpeextpseportlldptable *CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable) GetEntityData() *types.CommonEntityData {
    cpeextpseportlldptable.EntityData.YFilter = cpeextpseportlldptable.YFilter
    cpeextpseportlldptable.EntityData.YangName = "cpeExtPsePortLldpTable"
    cpeextpseportlldptable.EntityData.BundleName = "cisco_ios_xe"
    cpeextpseportlldptable.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeextpseportlldptable.EntityData.SegmentPath = "cpeExtPsePortLldpTable"
    cpeextpseportlldptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeextpseportlldptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeextpseportlldptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeextpseportlldptable.EntityData.Children = make(map[string]types.YChild)
    cpeextpseportlldptable.EntityData.Children["cpeExtPsePortLldpEntry"] = types.YChild{"Cpeextpseportlldpentry", nil}
    for i := range cpeextpseportlldptable.Cpeextpseportlldpentry {
        cpeextpseportlldptable.EntityData.Children[types.GetSegmentPath(&cpeextpseportlldptable.Cpeextpseportlldpentry[i])] = types.YChild{"Cpeextpseportlldpentry", &cpeextpseportlldptable.Cpeextpseportlldpentry[i]}
    }
    cpeextpseportlldptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cpeextpseportlldptable.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable_Cpeextpseportlldpentry
// A cpeExtPseDllPowerEntry entry contains the LLDP
// based DLL power classification characteristics for a particular
// PSE and the PD attached to it. 
// 
// A PSE has its entry here when all of the following conditions
// are satisfied:
// - The LLDP power classification is globally enabled.
// - It has a PD attached.
// - LLDP is the operational power negotiation protocol between
//   the PSE and the PD attached.
type CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable_Cpeextpseportlldpentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // power_ethernet_mib.POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportgroupindex
    Pethpseportgroupindex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // power_ethernet_mib.POWERETHERNETMIB_Pethpseporttable_Pethpseportentry_Pethpseportindex
    Pethpseportindex interface{}

    // The DTE Power via MDI type of the local system (PSE). The type is
    // CpeExtLldpPwrType.
    Cpeextpseportlldppwrtype interface{}

    // The DTE Power via MDI type of the remote system (PD). The type is
    // CpeExtLldpPwrType.
    Cpeextpseportlldppdpwrtype interface{}

    // The power source of the local system (PSE). The type is CpeExtLldpPwrSrc.
    Cpeextpseportlldppwrsrc interface{}

    // The power source of the remote system (PD). The type is CpeExtLldpPwrSrc.
    Cpeextpseportlldppdpwrsrc interface{}

    // The power priority of the local system (PSE). The type is
    // CpeExtPwrPriority.
    Cpeextpseportlldppwrpriority interface{}

    // The power priority of the remote system (PD). The type is
    // CpeExtPwrPriority.
    Cpeextpseportlldppdpwrpriority interface{}

    // The requested PD power value that the local system (PSE) mirrors back to
    // the remote system (PD). The type is interface{} with range: 0..4294967295.
    // Units are milliwatts.
    Cpeextpseportlldppwrreq interface{}

    // The PD requested power value received from the remote system (PD). The type
    // is interface{} with range: 0..4294967295. Units are milliwatts.
    Cpeextpseportlldppdpwrreq interface{}

    // The PSE allocated power value for the remote system (PD). The type is
    // interface{} with range: 0..4294967295. Units are milliwatts.
    Cpeextpseportlldppwralloc interface{}

    // The PSE allocated power value received from the remote system (PD). The
    // type is interface{} with range: 0..4294967295. Units are milliwatts.
    Cpeextpseportlldppdpwralloc interface{}
}

func (cpeextpseportlldpentry *CISCOPOWERETHERNETEXTMIB_Cpeextpseportlldptable_Cpeextpseportlldpentry) GetEntityData() *types.CommonEntityData {
    cpeextpseportlldpentry.EntityData.YFilter = cpeextpseportlldpentry.YFilter
    cpeextpseportlldpentry.EntityData.YangName = "cpeExtPsePortLldpEntry"
    cpeextpseportlldpentry.EntityData.BundleName = "cisco_ios_xe"
    cpeextpseportlldpentry.EntityData.ParentYangName = "cpeExtPsePortLldpTable"
    cpeextpseportlldpentry.EntityData.SegmentPath = "cpeExtPsePortLldpEntry" + "[pethPsePortGroupIndex='" + fmt.Sprintf("%v", cpeextpseportlldpentry.Pethpseportgroupindex) + "']" + "[pethPsePortIndex='" + fmt.Sprintf("%v", cpeextpseportlldpentry.Pethpseportindex) + "']"
    cpeextpseportlldpentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeextpseportlldpentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeextpseportlldpentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeextpseportlldpentry.EntityData.Children = make(map[string]types.YChild)
    cpeextpseportlldpentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cpeextpseportlldpentry.EntityData.Leafs["pethPsePortGroupIndex"] = types.YLeaf{"Pethpseportgroupindex", cpeextpseportlldpentry.Pethpseportgroupindex}
    cpeextpseportlldpentry.EntityData.Leafs["pethPsePortIndex"] = types.YLeaf{"Pethpseportindex", cpeextpseportlldpentry.Pethpseportindex}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPwrType"] = types.YLeaf{"Cpeextpseportlldppwrtype", cpeextpseportlldpentry.Cpeextpseportlldppwrtype}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPdPwrType"] = types.YLeaf{"Cpeextpseportlldppdpwrtype", cpeextpseportlldpentry.Cpeextpseportlldppdpwrtype}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPwrSrc"] = types.YLeaf{"Cpeextpseportlldppwrsrc", cpeextpseportlldpentry.Cpeextpseportlldppwrsrc}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPdPwrSrc"] = types.YLeaf{"Cpeextpseportlldppdpwrsrc", cpeextpseportlldpentry.Cpeextpseportlldppdpwrsrc}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPwrPriority"] = types.YLeaf{"Cpeextpseportlldppwrpriority", cpeextpseportlldpentry.Cpeextpseportlldppwrpriority}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPdPwrPriority"] = types.YLeaf{"Cpeextpseportlldppdpwrpriority", cpeextpseportlldpentry.Cpeextpseportlldppdpwrpriority}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPwrReq"] = types.YLeaf{"Cpeextpseportlldppwrreq", cpeextpseportlldpentry.Cpeextpseportlldppwrreq}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPdPwrReq"] = types.YLeaf{"Cpeextpseportlldppdpwrreq", cpeextpseportlldpentry.Cpeextpseportlldppdpwrreq}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPwrAlloc"] = types.YLeaf{"Cpeextpseportlldppwralloc", cpeextpseportlldpentry.Cpeextpseportlldppwralloc}
    cpeextpseportlldpentry.EntityData.Leafs["cpeExtPsePortLldpPdPwrAlloc"] = types.YLeaf{"Cpeextpseportlldppdpwralloc", cpeextpseportlldpentry.Cpeextpseportlldppdpwralloc}
    return &(cpeextpseportlldpentry.EntityData)
}

