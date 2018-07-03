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

    
    CpeExtMIBObjects CISCOPOWERETHERNETEXTMIB_CpeExtMIBObjects

    
    CpeExtPdStatistics CISCOPOWERETHERNETEXTMIB_CpeExtPdStatistics

    // This table contains the additional information for the main PSE group in
    // pethMainPseTable.
    CpeExtMainPseTable CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable

    // This table contains the statistics information of the powered devices
    // fallen into different power classifications in the system.
    CpeExtPdStatsTable CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable

    // A table provides the Link Layer Discovery Protocol (LLDP) based Data Link
    // Layer (DLL) power classification characteristics of PSE ports and PDs
    // attached to them.
    CpeExtPsePortLldpTable CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable
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

    cISCOPOWERETHERNETEXTMIB.EntityData.Children = types.NewOrderedMap()
    cISCOPOWERETHERNETEXTMIB.EntityData.Children.Append("cpeExtMIBObjects", types.YChild{"CpeExtMIBObjects", &cISCOPOWERETHERNETEXTMIB.CpeExtMIBObjects})
    cISCOPOWERETHERNETEXTMIB.EntityData.Children.Append("cpeExtPdStatistics", types.YChild{"CpeExtPdStatistics", &cISCOPOWERETHERNETEXTMIB.CpeExtPdStatistics})
    cISCOPOWERETHERNETEXTMIB.EntityData.Children.Append("cpeExtMainPseTable", types.YChild{"CpeExtMainPseTable", &cISCOPOWERETHERNETEXTMIB.CpeExtMainPseTable})
    cISCOPOWERETHERNETEXTMIB.EntityData.Children.Append("cpeExtPdStatsTable", types.YChild{"CpeExtPdStatsTable", &cISCOPOWERETHERNETEXTMIB.CpeExtPdStatsTable})
    cISCOPOWERETHERNETEXTMIB.EntityData.Children.Append("cpeExtPsePortLldpTable", types.YChild{"CpeExtPsePortLldpTable", &cISCOPOWERETHERNETEXTMIB.CpeExtPsePortLldpTable})
    cISCOPOWERETHERNETEXTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOPOWERETHERNETEXTMIB.EntityData.YListKeys = []string {}

    return &(cISCOPOWERETHERNETEXTMIB.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_CpeExtMIBObjects
type CISCOPOWERETHERNETEXTMIB_CpeExtMIBObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the default inline power allocation per port. This is
    // a global configuration parameter that applies to all inline power capable
    // ports in the system.  The system must consider this object as well as the
    // per port configuration object, cpeExtPsePortPwrMax, when determining how
    // much power to allocate to a port. The system will use the lower of the two
    // numbers. The type is interface{} with range: 0..4294967295. Units are
    // milliwatts.
    CpeExtDefaultAllocation interface{}

    // This object is used to enable/disable the the cpeExtPolicingNotif
    // notification. The type is bool.
    CpeExtPolicingNotifEnable interface{}

    // This object is the global control of the power priority feature on the
    // device. 'true' indicates that the power priority feature is globally
    // enabled. 'false' indicates that the power priority feature is globally
    // disabled. The type is bool.
    CpeExtPowerPriorityEnable interface{}
}

func (cpeExtMIBObjects *CISCOPOWERETHERNETEXTMIB_CpeExtMIBObjects) GetEntityData() *types.CommonEntityData {
    cpeExtMIBObjects.EntityData.YFilter = cpeExtMIBObjects.YFilter
    cpeExtMIBObjects.EntityData.YangName = "cpeExtMIBObjects"
    cpeExtMIBObjects.EntityData.BundleName = "cisco_ios_xe"
    cpeExtMIBObjects.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeExtMIBObjects.EntityData.SegmentPath = "cpeExtMIBObjects"
    cpeExtMIBObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeExtMIBObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeExtMIBObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeExtMIBObjects.EntityData.Children = types.NewOrderedMap()
    cpeExtMIBObjects.EntityData.Leafs = types.NewOrderedMap()
    cpeExtMIBObjects.EntityData.Leafs.Append("cpeExtDefaultAllocation", types.YLeaf{"CpeExtDefaultAllocation", cpeExtMIBObjects.CpeExtDefaultAllocation})
    cpeExtMIBObjects.EntityData.Leafs.Append("cpeExtPolicingNotifEnable", types.YLeaf{"CpeExtPolicingNotifEnable", cpeExtMIBObjects.CpeExtPolicingNotifEnable})
    cpeExtMIBObjects.EntityData.Leafs.Append("cpeExtPowerPriorityEnable", types.YLeaf{"CpeExtPowerPriorityEnable", cpeExtMIBObjects.CpeExtPowerPriorityEnable})

    cpeExtMIBObjects.EntityData.YListKeys = []string {}

    return &(cpeExtMIBObjects.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_CpeExtPdStatistics
type CISCOPOWERETHERNETEXTMIB_CpeExtPdStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the total number of the powered devices with any
    // power classifications in the system.  Classification is a way to tag
    // different terminals on the Power over LAN network according to their power
    // consumption. Devices such as IP telephones, WLAN access points and others,
    // will be classified according to their power requirements. The type is
    // interface{} with range: 0..4294967295.
    CpeExtPdStatsTotalDevices interface{}
}

func (cpeExtPdStatistics *CISCOPOWERETHERNETEXTMIB_CpeExtPdStatistics) GetEntityData() *types.CommonEntityData {
    cpeExtPdStatistics.EntityData.YFilter = cpeExtPdStatistics.YFilter
    cpeExtPdStatistics.EntityData.YangName = "cpeExtPdStatistics"
    cpeExtPdStatistics.EntityData.BundleName = "cisco_ios_xe"
    cpeExtPdStatistics.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeExtPdStatistics.EntityData.SegmentPath = "cpeExtPdStatistics"
    cpeExtPdStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeExtPdStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeExtPdStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeExtPdStatistics.EntityData.Children = types.NewOrderedMap()
    cpeExtPdStatistics.EntityData.Leafs = types.NewOrderedMap()
    cpeExtPdStatistics.EntityData.Leafs.Append("cpeExtPdStatsTotalDevices", types.YLeaf{"CpeExtPdStatsTotalDevices", cpeExtPdStatistics.CpeExtPdStatsTotalDevices})

    cpeExtPdStatistics.EntityData.YListKeys = []string {}

    return &(cpeExtPdStatistics.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable
// This table contains the additional information for the
// main PSE group in pethMainPseTable.
type CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cpeExtMainPseEntry contains information about a particular
    // pethMainPseGroupIndex. An entry is created by the agent when a main PSE
    // group is added to the pethMainPseTable. An entry is deleted by the agent
    // when a main PSE group is removed from the pethMainPseTable. The type is
    // slice of CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable_CpeExtMainPseEntry.
    CpeExtMainPseEntry []*CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable_CpeExtMainPseEntry
}

func (cpeExtMainPseTable *CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable) GetEntityData() *types.CommonEntityData {
    cpeExtMainPseTable.EntityData.YFilter = cpeExtMainPseTable.YFilter
    cpeExtMainPseTable.EntityData.YangName = "cpeExtMainPseTable"
    cpeExtMainPseTable.EntityData.BundleName = "cisco_ios_xe"
    cpeExtMainPseTable.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeExtMainPseTable.EntityData.SegmentPath = "cpeExtMainPseTable"
    cpeExtMainPseTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeExtMainPseTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeExtMainPseTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeExtMainPseTable.EntityData.Children = types.NewOrderedMap()
    cpeExtMainPseTable.EntityData.Children.Append("cpeExtMainPseEntry", types.YChild{"CpeExtMainPseEntry", nil})
    for i := range cpeExtMainPseTable.CpeExtMainPseEntry {
        cpeExtMainPseTable.EntityData.Children.Append(types.GetSegmentPath(cpeExtMainPseTable.CpeExtMainPseEntry[i]), types.YChild{"CpeExtMainPseEntry", cpeExtMainPseTable.CpeExtMainPseEntry[i]})
    }
    cpeExtMainPseTable.EntityData.Leafs = types.NewOrderedMap()

    cpeExtMainPseTable.EntityData.YListKeys = []string {}

    return &(cpeExtMainPseTable.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable_CpeExtMainPseEntry
// A cpeExtMainPseEntry contains information about
// a particular pethMainPseGroupIndex. An entry is
// created by the agent when a main PSE group is added
// to the pethMainPseTable. An entry is deleted by the
// agent when a main PSE group is removed from the
// pethMainPseTable.
type CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable_CpeExtMainPseEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // power_ethernet_mib.POWERETHERNETMIB_PethMainPseTable_PethMainPseEntry_PethMainPseGroupIndex
    PethMainPseGroupIndex interface{}

    // The entPhysicalIndex value that uniquely identifies the main PSE group. If
    // the main PSE group does not have a corresponding physical entry in
    // entPhysicalTable or if the entPhysicalTable is not supported by the
    // management system, then this object has the value of zero. The type is
    // interface{} with range: 0..2147483647.
    CpeExtMainPseEntPhyIndex interface{}

    // A textual string containing information about the Power Source Equipment
    // (PSE) group. The type is string.
    CpeExtMainPseDescr interface{}

    // This object indicates if the given group is capable of monitoring the power
    // consumption of the interfaces that belong to the group. The value 'true'
    // means that the group is capable. The value 'false' means that the group in
    // not capable. The type is bool.
    CpeExtMainPsePwrMonitorCapable interface{}

    // Used power expressed in miliwatts. The type is interface{} with range:
    // 0..4294967295. Units are miliwatts.
    CpeExtMainPseUsedPower interface{}

    // Remaining power expressed in miliwatts, this parameter is calculated as
    // pethMainPsePower minus cpeExtMainPseUsedPower. The type is interface{} with
    // range: 0..4294967295. Units are miliwatts.
    CpeExtMainPseRemainingPower interface{}
}

func (cpeExtMainPseEntry *CISCOPOWERETHERNETEXTMIB_CpeExtMainPseTable_CpeExtMainPseEntry) GetEntityData() *types.CommonEntityData {
    cpeExtMainPseEntry.EntityData.YFilter = cpeExtMainPseEntry.YFilter
    cpeExtMainPseEntry.EntityData.YangName = "cpeExtMainPseEntry"
    cpeExtMainPseEntry.EntityData.BundleName = "cisco_ios_xe"
    cpeExtMainPseEntry.EntityData.ParentYangName = "cpeExtMainPseTable"
    cpeExtMainPseEntry.EntityData.SegmentPath = "cpeExtMainPseEntry" + types.AddKeyToken(cpeExtMainPseEntry.PethMainPseGroupIndex, "pethMainPseGroupIndex")
    cpeExtMainPseEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeExtMainPseEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeExtMainPseEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeExtMainPseEntry.EntityData.Children = types.NewOrderedMap()
    cpeExtMainPseEntry.EntityData.Leafs = types.NewOrderedMap()
    cpeExtMainPseEntry.EntityData.Leafs.Append("pethMainPseGroupIndex", types.YLeaf{"PethMainPseGroupIndex", cpeExtMainPseEntry.PethMainPseGroupIndex})
    cpeExtMainPseEntry.EntityData.Leafs.Append("cpeExtMainPseEntPhyIndex", types.YLeaf{"CpeExtMainPseEntPhyIndex", cpeExtMainPseEntry.CpeExtMainPseEntPhyIndex})
    cpeExtMainPseEntry.EntityData.Leafs.Append("cpeExtMainPseDescr", types.YLeaf{"CpeExtMainPseDescr", cpeExtMainPseEntry.CpeExtMainPseDescr})
    cpeExtMainPseEntry.EntityData.Leafs.Append("cpeExtMainPsePwrMonitorCapable", types.YLeaf{"CpeExtMainPsePwrMonitorCapable", cpeExtMainPseEntry.CpeExtMainPsePwrMonitorCapable})
    cpeExtMainPseEntry.EntityData.Leafs.Append("cpeExtMainPseUsedPower", types.YLeaf{"CpeExtMainPseUsedPower", cpeExtMainPseEntry.CpeExtMainPseUsedPower})
    cpeExtMainPseEntry.EntityData.Leafs.Append("cpeExtMainPseRemainingPower", types.YLeaf{"CpeExtMainPseRemainingPower", cpeExtMainPseEntry.CpeExtMainPseRemainingPower})

    cpeExtMainPseEntry.EntityData.YListKeys = []string {"PethMainPseGroupIndex"}

    return &(cpeExtMainPseEntry.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable
// This table contains the statistics information
// of the powered devices fallen into different power
// classifications in the system.
type CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cpeExtPdStatsEntry contains the statistics information about a particular
    // power classification defined in cpeExtPdStatsIndex. The type is slice of
    // CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry.
    CpeExtPdStatsEntry []*CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry
}

func (cpeExtPdStatsTable *CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable) GetEntityData() *types.CommonEntityData {
    cpeExtPdStatsTable.EntityData.YFilter = cpeExtPdStatsTable.YFilter
    cpeExtPdStatsTable.EntityData.YangName = "cpeExtPdStatsTable"
    cpeExtPdStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cpeExtPdStatsTable.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeExtPdStatsTable.EntityData.SegmentPath = "cpeExtPdStatsTable"
    cpeExtPdStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeExtPdStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeExtPdStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeExtPdStatsTable.EntityData.Children = types.NewOrderedMap()
    cpeExtPdStatsTable.EntityData.Children.Append("cpeExtPdStatsEntry", types.YChild{"CpeExtPdStatsEntry", nil})
    for i := range cpeExtPdStatsTable.CpeExtPdStatsEntry {
        cpeExtPdStatsTable.EntityData.Children.Append(types.GetSegmentPath(cpeExtPdStatsTable.CpeExtPdStatsEntry[i]), types.YChild{"CpeExtPdStatsEntry", cpeExtPdStatsTable.CpeExtPdStatsEntry[i]})
    }
    cpeExtPdStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cpeExtPdStatsTable.EntityData.YListKeys = []string {}

    return &(cpeExtPdStatsTable.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry
// A cpeExtPdStatsEntry contains the statistics
// information about a particular power classification
// defined in cpeExtPdStatsIndex.
type CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry struct {
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
    // specifications. The type is CpeExtPdStatsClass.
    CpeExtPdStatsClass interface{}

    // This object indicates the count of the powered devices whose power
    // classification falls into  a specific value of cpeExtPdStatsIndex. The type
    // is interface{} with range: 0..4294967295.
    CpeExtPdStatsDeviceCount interface{}
}

func (cpeExtPdStatsEntry *CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry) GetEntityData() *types.CommonEntityData {
    cpeExtPdStatsEntry.EntityData.YFilter = cpeExtPdStatsEntry.YFilter
    cpeExtPdStatsEntry.EntityData.YangName = "cpeExtPdStatsEntry"
    cpeExtPdStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cpeExtPdStatsEntry.EntityData.ParentYangName = "cpeExtPdStatsTable"
    cpeExtPdStatsEntry.EntityData.SegmentPath = "cpeExtPdStatsEntry" + types.AddKeyToken(cpeExtPdStatsEntry.CpeExtPdStatsClass, "cpeExtPdStatsClass")
    cpeExtPdStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeExtPdStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeExtPdStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeExtPdStatsEntry.EntityData.Children = types.NewOrderedMap()
    cpeExtPdStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cpeExtPdStatsEntry.EntityData.Leafs.Append("cpeExtPdStatsClass", types.YLeaf{"CpeExtPdStatsClass", cpeExtPdStatsEntry.CpeExtPdStatsClass})
    cpeExtPdStatsEntry.EntityData.Leafs.Append("cpeExtPdStatsDeviceCount", types.YLeaf{"CpeExtPdStatsDeviceCount", cpeExtPdStatsEntry.CpeExtPdStatsDeviceCount})

    cpeExtPdStatsEntry.EntityData.YListKeys = []string {"CpeExtPdStatsClass"}

    return &(cpeExtPdStatsEntry.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass represents class 4 in IEEE specifications.
type CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass string

const (
    CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass_cisco CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass = "cisco"

    CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass_class0 CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass = "class0"

    CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass_class1 CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass = "class1"

    CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass_class2 CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass = "class2"

    CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass_class3 CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass = "class3"

    CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass_class4 CISCOPOWERETHERNETEXTMIB_CpeExtPdStatsTable_CpeExtPdStatsEntry_CpeExtPdStatsClass = "class4"
)

// CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable
// A table provides the Link Layer Discovery Protocol (LLDP)
// based Data Link Layer (DLL) power classification
// characteristics of PSE ports and PDs attached to them.
type CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A cpeExtPseDllPowerEntry entry contains the LLDP based DLL power
    // classification characteristics for a particular PSE and the PD attached to
    // it.   A PSE has its entry here when all of the following conditions are
    // satisfied: - The LLDP power classification is globally enabled. - It has a
    // PD attached. - LLDP is the operational power negotiation protocol between  
    // the PSE and the PD attached. The type is slice of
    // CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable_CpeExtPsePortLldpEntry.
    CpeExtPsePortLldpEntry []*CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable_CpeExtPsePortLldpEntry
}

func (cpeExtPsePortLldpTable *CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable) GetEntityData() *types.CommonEntityData {
    cpeExtPsePortLldpTable.EntityData.YFilter = cpeExtPsePortLldpTable.YFilter
    cpeExtPsePortLldpTable.EntityData.YangName = "cpeExtPsePortLldpTable"
    cpeExtPsePortLldpTable.EntityData.BundleName = "cisco_ios_xe"
    cpeExtPsePortLldpTable.EntityData.ParentYangName = "CISCO-POWER-ETHERNET-EXT-MIB"
    cpeExtPsePortLldpTable.EntityData.SegmentPath = "cpeExtPsePortLldpTable"
    cpeExtPsePortLldpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeExtPsePortLldpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeExtPsePortLldpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeExtPsePortLldpTable.EntityData.Children = types.NewOrderedMap()
    cpeExtPsePortLldpTable.EntityData.Children.Append("cpeExtPsePortLldpEntry", types.YChild{"CpeExtPsePortLldpEntry", nil})
    for i := range cpeExtPsePortLldpTable.CpeExtPsePortLldpEntry {
        cpeExtPsePortLldpTable.EntityData.Children.Append(types.GetSegmentPath(cpeExtPsePortLldpTable.CpeExtPsePortLldpEntry[i]), types.YChild{"CpeExtPsePortLldpEntry", cpeExtPsePortLldpTable.CpeExtPsePortLldpEntry[i]})
    }
    cpeExtPsePortLldpTable.EntityData.Leafs = types.NewOrderedMap()

    cpeExtPsePortLldpTable.EntityData.YListKeys = []string {}

    return &(cpeExtPsePortLldpTable.EntityData)
}

// CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable_CpeExtPsePortLldpEntry
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
type CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable_CpeExtPsePortLldpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // power_ethernet_mib.POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortGroupIndex
    PethPsePortGroupIndex interface{}

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // power_ethernet_mib.POWERETHERNETMIB_PethPsePortTable_PethPsePortEntry_PethPsePortIndex
    PethPsePortIndex interface{}

    // The DTE Power via MDI type of the local system (PSE). The type is
    // CpeExtLldpPwrType.
    CpeExtPsePortLldpPwrType interface{}

    // The DTE Power via MDI type of the remote system (PD). The type is
    // CpeExtLldpPwrType.
    CpeExtPsePortLldpPdPwrType interface{}

    // The power source of the local system (PSE). The type is CpeExtLldpPwrSrc.
    CpeExtPsePortLldpPwrSrc interface{}

    // The power source of the remote system (PD). The type is CpeExtLldpPwrSrc.
    CpeExtPsePortLldpPdPwrSrc interface{}

    // The power priority of the local system (PSE). The type is
    // CpeExtPwrPriority.
    CpeExtPsePortLldpPwrPriority interface{}

    // The power priority of the remote system (PD). The type is
    // CpeExtPwrPriority.
    CpeExtPsePortLldpPdPwrPriority interface{}

    // The requested PD power value that the local system (PSE) mirrors back to
    // the remote system (PD). The type is interface{} with range: 0..4294967295.
    // Units are milliwatts.
    CpeExtPsePortLldpPwrReq interface{}

    // The PD requested power value received from the remote system (PD). The type
    // is interface{} with range: 0..4294967295. Units are milliwatts.
    CpeExtPsePortLldpPdPwrReq interface{}

    // The PSE allocated power value for the remote system (PD). The type is
    // interface{} with range: 0..4294967295. Units are milliwatts.
    CpeExtPsePortLldpPwrAlloc interface{}

    // The PSE allocated power value received from the remote system (PD). The
    // type is interface{} with range: 0..4294967295. Units are milliwatts.
    CpeExtPsePortLldpPdPwrAlloc interface{}
}

func (cpeExtPsePortLldpEntry *CISCOPOWERETHERNETEXTMIB_CpeExtPsePortLldpTable_CpeExtPsePortLldpEntry) GetEntityData() *types.CommonEntityData {
    cpeExtPsePortLldpEntry.EntityData.YFilter = cpeExtPsePortLldpEntry.YFilter
    cpeExtPsePortLldpEntry.EntityData.YangName = "cpeExtPsePortLldpEntry"
    cpeExtPsePortLldpEntry.EntityData.BundleName = "cisco_ios_xe"
    cpeExtPsePortLldpEntry.EntityData.ParentYangName = "cpeExtPsePortLldpTable"
    cpeExtPsePortLldpEntry.EntityData.SegmentPath = "cpeExtPsePortLldpEntry" + types.AddKeyToken(cpeExtPsePortLldpEntry.PethPsePortGroupIndex, "pethPsePortGroupIndex") + types.AddKeyToken(cpeExtPsePortLldpEntry.PethPsePortIndex, "pethPsePortIndex")
    cpeExtPsePortLldpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpeExtPsePortLldpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpeExtPsePortLldpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpeExtPsePortLldpEntry.EntityData.Children = types.NewOrderedMap()
    cpeExtPsePortLldpEntry.EntityData.Leafs = types.NewOrderedMap()
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("pethPsePortGroupIndex", types.YLeaf{"PethPsePortGroupIndex", cpeExtPsePortLldpEntry.PethPsePortGroupIndex})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("pethPsePortIndex", types.YLeaf{"PethPsePortIndex", cpeExtPsePortLldpEntry.PethPsePortIndex})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPwrType", types.YLeaf{"CpeExtPsePortLldpPwrType", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPwrType})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPdPwrType", types.YLeaf{"CpeExtPsePortLldpPdPwrType", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPdPwrType})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPwrSrc", types.YLeaf{"CpeExtPsePortLldpPwrSrc", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPwrSrc})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPdPwrSrc", types.YLeaf{"CpeExtPsePortLldpPdPwrSrc", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPdPwrSrc})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPwrPriority", types.YLeaf{"CpeExtPsePortLldpPwrPriority", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPwrPriority})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPdPwrPriority", types.YLeaf{"CpeExtPsePortLldpPdPwrPriority", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPdPwrPriority})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPwrReq", types.YLeaf{"CpeExtPsePortLldpPwrReq", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPwrReq})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPdPwrReq", types.YLeaf{"CpeExtPsePortLldpPdPwrReq", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPdPwrReq})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPwrAlloc", types.YLeaf{"CpeExtPsePortLldpPwrAlloc", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPwrAlloc})
    cpeExtPsePortLldpEntry.EntityData.Leafs.Append("cpeExtPsePortLldpPdPwrAlloc", types.YLeaf{"CpeExtPsePortLldpPdPwrAlloc", cpeExtPsePortLldpEntry.CpeExtPsePortLldpPdPwrAlloc})

    cpeExtPsePortLldpEntry.EntityData.YListKeys = []string {"PethPsePortGroupIndex", "PethPsePortIndex"}

    return &(cpeExtPsePortLldpEntry.EntityData)
}

