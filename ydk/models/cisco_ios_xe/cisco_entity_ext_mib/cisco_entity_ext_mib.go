// This MIB is an extension of the ENTITY-MIB
// specified in RFC2737.
// 
// This MIB module contains  Cisco-defined  extensions 
// to the  entityPhysicalTable to represent information
// related to entities of class module(entPhysicalClass
// = 'module') which have a Processor.
// 
// A processor module is defined as a physical entity
// that has a CPU, RAM and NVRAM  so that
// it can independently
//    - load a bootable image
//    - save configuration.
// This module is the entry point for external
// applications like SNMP Manager, CLI, FTP etc.
// 
// Line card is an interface card         with at least a 
// Processor and RAM. This might be referred to as 
// Service Module in some cisco products.
// 
// A configuration register is  a 16 bit 
// software register.
// The configuration register is mainly used to 
// check for instructions on where to find the Cisco 
// Operating System software.
// Some other functions of configuration register are:
//  - To select a boot source and default boot filename.
//  - To enable or disable the Break function.
//  - To control broadcast addresses.
//  - To set the console terminal baud rate.
//  - To load operating software from Flash memory.
//  - To allow us to manually boot the system using the 
//    boot command at the bootstrap program prompt.
// 
// Booting is the process of initializing the
// hardware and starting the Operating System.
package cisco_entity_ext_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_entity_ext_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ENTITY-EXT-MIB CISCO-ENTITY-EXT-MIB}", reflect.TypeOf(CISCOENTITYEXTMIB{}))
    ydk.RegisterEntity("CISCO-ENTITY-EXT-MIB:CISCO-ENTITY-EXT-MIB", reflect.TypeOf(CISCOENTITYEXTMIB{}))
}

// CISCOENTITYEXTMIB
type CISCOENTITYEXTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The ceExtPhysicalProcessorTable extends the ENTITY-MIB entPhysicalTable for
    // modules (Non FRUs(Field Replacable Units) or FRUs).
    CeExtPhysicalProcessorTable CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable

    // The ceExtConfigRegTable extends the ENTITY-MIB entPhysicalTable.
    CeExtConfigRegTable CISCOENTITYEXTMIB_CeExtConfigRegTable

    // A table containing information of LED on an entity.
    CeExtEntityLEDTable CISCOENTITYEXTMIB_CeExtEntityLEDTable
}

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetEntityData() *types.CommonEntityData {
    cISCOENTITYEXTMIB.EntityData.YFilter = cISCOENTITYEXTMIB.YFilter
    cISCOENTITYEXTMIB.EntityData.YangName = "CISCO-ENTITY-EXT-MIB"
    cISCOENTITYEXTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOENTITYEXTMIB.EntityData.ParentYangName = "CISCO-ENTITY-EXT-MIB"
    cISCOENTITYEXTMIB.EntityData.SegmentPath = "CISCO-ENTITY-EXT-MIB:CISCO-ENTITY-EXT-MIB"
    cISCOENTITYEXTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOENTITYEXTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOENTITYEXTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOENTITYEXTMIB.EntityData.Children = types.NewOrderedMap()
    cISCOENTITYEXTMIB.EntityData.Children.Append("ceExtPhysicalProcessorTable", types.YChild{"CeExtPhysicalProcessorTable", &cISCOENTITYEXTMIB.CeExtPhysicalProcessorTable})
    cISCOENTITYEXTMIB.EntityData.Children.Append("ceExtConfigRegTable", types.YChild{"CeExtConfigRegTable", &cISCOENTITYEXTMIB.CeExtConfigRegTable})
    cISCOENTITYEXTMIB.EntityData.Children.Append("ceExtEntityLEDTable", types.YChild{"CeExtEntityLEDTable", &cISCOENTITYEXTMIB.CeExtEntityLEDTable})
    cISCOENTITYEXTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOENTITYEXTMIB.EntityData.YListKeys = []string {}

    return &(cISCOENTITYEXTMIB.EntityData)
}

// CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable
// The ceExtPhysicalProcessorTable extends
// the ENTITY-MIB entPhysicalTable for modules
// (Non FRUs(Field Replacable Units) or FRUs).
type CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A ceExtPhysicalProcessorTable entry extends a corresponding
    // entPhysicalTable entry of class module(entPhysicalClass = 'module').  A
    // processor module or line card which  has a processor will have an entry in
    // this table.  A processor module or line card having multiple processors and
    // is a SMP(Symmetric multi processor) system will have only  one entry
    // corresponding to all the processors  since the resources defined below are
    // shared.  A processor module or line card having multiple processors and is
    // not an SMP system would register the processors as separate entities. 
    // Entries are created by the agent at the system power-up or module
    // insertion.  Entries are removed when the module is reset or removed. The
    // type is slice of
    // CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable_CeExtPhysicalProcessorEntry.
    CeExtPhysicalProcessorEntry []*CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable_CeExtPhysicalProcessorEntry
}

func (ceExtPhysicalProcessorTable *CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable) GetEntityData() *types.CommonEntityData {
    ceExtPhysicalProcessorTable.EntityData.YFilter = ceExtPhysicalProcessorTable.YFilter
    ceExtPhysicalProcessorTable.EntityData.YangName = "ceExtPhysicalProcessorTable"
    ceExtPhysicalProcessorTable.EntityData.BundleName = "cisco_ios_xe"
    ceExtPhysicalProcessorTable.EntityData.ParentYangName = "CISCO-ENTITY-EXT-MIB"
    ceExtPhysicalProcessorTable.EntityData.SegmentPath = "ceExtPhysicalProcessorTable"
    ceExtPhysicalProcessorTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceExtPhysicalProcessorTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceExtPhysicalProcessorTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceExtPhysicalProcessorTable.EntityData.Children = types.NewOrderedMap()
    ceExtPhysicalProcessorTable.EntityData.Children.Append("ceExtPhysicalProcessorEntry", types.YChild{"CeExtPhysicalProcessorEntry", nil})
    for i := range ceExtPhysicalProcessorTable.CeExtPhysicalProcessorEntry {
        ceExtPhysicalProcessorTable.EntityData.Children.Append(types.GetSegmentPath(ceExtPhysicalProcessorTable.CeExtPhysicalProcessorEntry[i]), types.YChild{"CeExtPhysicalProcessorEntry", ceExtPhysicalProcessorTable.CeExtPhysicalProcessorEntry[i]})
    }
    ceExtPhysicalProcessorTable.EntityData.Leafs = types.NewOrderedMap()

    ceExtPhysicalProcessorTable.EntityData.YListKeys = []string {}

    return &(ceExtPhysicalProcessorTable.EntityData)
}

// CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable_CeExtPhysicalProcessorEntry
// A ceExtPhysicalProcessorTable entry extends
// a corresponding entPhysicalTable entry of class
// module(entPhysicalClass = 'module').
// 
// A processor module or line card which 
// has a processor will have an entry in
// this table.
// 
// A processor module or line card having
// multiple processors and is a SMP(Symmetric
// multi processor) system will have only 
// one entry corresponding to all the processors 
// since the resources defined below are shared.
// 
// A processor module or line card having
// multiple processors and is not an SMP system
// would register the processors as separate entities.
// 
// Entries are created by the agent at the system power-up
// or module insertion.
// 
// Entries are removed when the module is reset or removed.
type CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable_CeExtPhysicalProcessorEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // Total number of bytes of RAM available on the Processor. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CeExtProcessorRam interface{}

    // Total number of bytes of NVRAM in the entity.  A value of 0 for this object
    // means the entity does not support NVRAM or NVRAM information  is not
    // available. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    CeExtNVRAMSize interface{}

    // Number of bytes of NVRAM in use. This object is irrelevant if
    // ceExtNVRAMSize is 0. The type is interface{} with range: 0..4294967295.
    // Units are bytes.
    CeExtNVRAMUsed interface{}

    // This object represents the upper 32-bit of ceExtProcessorRam. This object
    // needs to be supported only if the available RAM bytes exceeds 32-bit,
    // otherwise this object value would be set to 0. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    CeExtProcessorRamOverflow interface{}

    // This object represents the total number of bytes of RAM available on the
    // Processor. This object is a 64-bit version of ceExtProcessorRam. The type
    // is interface{} with range: 0..18446744073709551615. Units are bytes.
    CeExtHCProcessorRam interface{}
}

func (ceExtPhysicalProcessorEntry *CISCOENTITYEXTMIB_CeExtPhysicalProcessorTable_CeExtPhysicalProcessorEntry) GetEntityData() *types.CommonEntityData {
    ceExtPhysicalProcessorEntry.EntityData.YFilter = ceExtPhysicalProcessorEntry.YFilter
    ceExtPhysicalProcessorEntry.EntityData.YangName = "ceExtPhysicalProcessorEntry"
    ceExtPhysicalProcessorEntry.EntityData.BundleName = "cisco_ios_xe"
    ceExtPhysicalProcessorEntry.EntityData.ParentYangName = "ceExtPhysicalProcessorTable"
    ceExtPhysicalProcessorEntry.EntityData.SegmentPath = "ceExtPhysicalProcessorEntry" + types.AddKeyToken(ceExtPhysicalProcessorEntry.EntPhysicalIndex, "entPhysicalIndex")
    ceExtPhysicalProcessorEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceExtPhysicalProcessorEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceExtPhysicalProcessorEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceExtPhysicalProcessorEntry.EntityData.Children = types.NewOrderedMap()
    ceExtPhysicalProcessorEntry.EntityData.Leafs = types.NewOrderedMap()
    ceExtPhysicalProcessorEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ceExtPhysicalProcessorEntry.EntPhysicalIndex})
    ceExtPhysicalProcessorEntry.EntityData.Leafs.Append("ceExtProcessorRam", types.YLeaf{"CeExtProcessorRam", ceExtPhysicalProcessorEntry.CeExtProcessorRam})
    ceExtPhysicalProcessorEntry.EntityData.Leafs.Append("ceExtNVRAMSize", types.YLeaf{"CeExtNVRAMSize", ceExtPhysicalProcessorEntry.CeExtNVRAMSize})
    ceExtPhysicalProcessorEntry.EntityData.Leafs.Append("ceExtNVRAMUsed", types.YLeaf{"CeExtNVRAMUsed", ceExtPhysicalProcessorEntry.CeExtNVRAMUsed})
    ceExtPhysicalProcessorEntry.EntityData.Leafs.Append("ceExtProcessorRamOverflow", types.YLeaf{"CeExtProcessorRamOverflow", ceExtPhysicalProcessorEntry.CeExtProcessorRamOverflow})
    ceExtPhysicalProcessorEntry.EntityData.Leafs.Append("ceExtHCProcessorRam", types.YLeaf{"CeExtHCProcessorRam", ceExtPhysicalProcessorEntry.CeExtHCProcessorRam})

    ceExtPhysicalProcessorEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(ceExtPhysicalProcessorEntry.EntityData)
}

// CISCOENTITYEXTMIB_CeExtConfigRegTable
// The ceExtConfigRegTable extends
// the ENTITY-MIB entPhysicalTable.
type CISCOENTITYEXTMIB_CeExtConfigRegTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A ceExtConfigRegTable entry extends a corresponding entPhysicalTable entry
    // of class module which has a configuration register.  Entries are created by
    // the agent at the system power-up or module insertion.  Entries are removed
    // when the module is reset or  removed. The type is slice of
    // CISCOENTITYEXTMIB_CeExtConfigRegTable_CeExtConfigRegEntry.
    CeExtConfigRegEntry []*CISCOENTITYEXTMIB_CeExtConfigRegTable_CeExtConfigRegEntry
}

func (ceExtConfigRegTable *CISCOENTITYEXTMIB_CeExtConfigRegTable) GetEntityData() *types.CommonEntityData {
    ceExtConfigRegTable.EntityData.YFilter = ceExtConfigRegTable.YFilter
    ceExtConfigRegTable.EntityData.YangName = "ceExtConfigRegTable"
    ceExtConfigRegTable.EntityData.BundleName = "cisco_ios_xe"
    ceExtConfigRegTable.EntityData.ParentYangName = "CISCO-ENTITY-EXT-MIB"
    ceExtConfigRegTable.EntityData.SegmentPath = "ceExtConfigRegTable"
    ceExtConfigRegTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceExtConfigRegTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceExtConfigRegTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceExtConfigRegTable.EntityData.Children = types.NewOrderedMap()
    ceExtConfigRegTable.EntityData.Children.Append("ceExtConfigRegEntry", types.YChild{"CeExtConfigRegEntry", nil})
    for i := range ceExtConfigRegTable.CeExtConfigRegEntry {
        ceExtConfigRegTable.EntityData.Children.Append(types.GetSegmentPath(ceExtConfigRegTable.CeExtConfigRegEntry[i]), types.YChild{"CeExtConfigRegEntry", ceExtConfigRegTable.CeExtConfigRegEntry[i]})
    }
    ceExtConfigRegTable.EntityData.Leafs = types.NewOrderedMap()

    ceExtConfigRegTable.EntityData.YListKeys = []string {}

    return &(ceExtConfigRegTable.EntityData)
}

// CISCOENTITYEXTMIB_CeExtConfigRegTable_CeExtConfigRegEntry
// A ceExtConfigRegTable entry extends
// a corresponding entPhysicalTable entry of class
// module which has a configuration register.
// 
// Entries are created by the agent at the system power-up
// or module insertion.
// 
// Entries are removed when the module is reset or 
// removed.
type CISCOENTITYEXTMIB_CeExtConfigRegTable_CeExtConfigRegEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // The value of configuration register with which the processor module booted.
    // The type is string.
    CeExtConfigRegister interface{}

    // The value of configuration register in the processor module at next reboot.
    // Just after  the reboot this has the same value as  ceExtConfigRegister. The
    // type is string.
    CeExtConfigRegNext interface{}

    // The list of system boot images which can be used for booting. The type is
    // string with length: 0..255.
    CeExtSysBootImageList interface{}

    // The list of system kickstart images which can be used for booting. The type
    // is string with length: 0..255.
    CeExtKickstartImageList interface{}
}

func (ceExtConfigRegEntry *CISCOENTITYEXTMIB_CeExtConfigRegTable_CeExtConfigRegEntry) GetEntityData() *types.CommonEntityData {
    ceExtConfigRegEntry.EntityData.YFilter = ceExtConfigRegEntry.YFilter
    ceExtConfigRegEntry.EntityData.YangName = "ceExtConfigRegEntry"
    ceExtConfigRegEntry.EntityData.BundleName = "cisco_ios_xe"
    ceExtConfigRegEntry.EntityData.ParentYangName = "ceExtConfigRegTable"
    ceExtConfigRegEntry.EntityData.SegmentPath = "ceExtConfigRegEntry" + types.AddKeyToken(ceExtConfigRegEntry.EntPhysicalIndex, "entPhysicalIndex")
    ceExtConfigRegEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceExtConfigRegEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceExtConfigRegEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceExtConfigRegEntry.EntityData.Children = types.NewOrderedMap()
    ceExtConfigRegEntry.EntityData.Leafs = types.NewOrderedMap()
    ceExtConfigRegEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ceExtConfigRegEntry.EntPhysicalIndex})
    ceExtConfigRegEntry.EntityData.Leafs.Append("ceExtConfigRegister", types.YLeaf{"CeExtConfigRegister", ceExtConfigRegEntry.CeExtConfigRegister})
    ceExtConfigRegEntry.EntityData.Leafs.Append("ceExtConfigRegNext", types.YLeaf{"CeExtConfigRegNext", ceExtConfigRegEntry.CeExtConfigRegNext})
    ceExtConfigRegEntry.EntityData.Leafs.Append("ceExtSysBootImageList", types.YLeaf{"CeExtSysBootImageList", ceExtConfigRegEntry.CeExtSysBootImageList})
    ceExtConfigRegEntry.EntityData.Leafs.Append("ceExtKickstartImageList", types.YLeaf{"CeExtKickstartImageList", ceExtConfigRegEntry.CeExtKickstartImageList})

    ceExtConfigRegEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(ceExtConfigRegEntry.EntityData)
}

// CISCOENTITYEXTMIB_CeExtEntityLEDTable
// A table containing information of LED on an entity.
type CISCOENTITYEXTMIB_CeExtEntityLEDTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ceExtEntityLEDTable, containing
    // information about an LED on an entity, identified by  entPhysicalIndex. The
    // type is slice of CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry.
    CeExtEntityLEDEntry []*CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry
}

func (ceExtEntityLEDTable *CISCOENTITYEXTMIB_CeExtEntityLEDTable) GetEntityData() *types.CommonEntityData {
    ceExtEntityLEDTable.EntityData.YFilter = ceExtEntityLEDTable.YFilter
    ceExtEntityLEDTable.EntityData.YangName = "ceExtEntityLEDTable"
    ceExtEntityLEDTable.EntityData.BundleName = "cisco_ios_xe"
    ceExtEntityLEDTable.EntityData.ParentYangName = "CISCO-ENTITY-EXT-MIB"
    ceExtEntityLEDTable.EntityData.SegmentPath = "ceExtEntityLEDTable"
    ceExtEntityLEDTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceExtEntityLEDTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceExtEntityLEDTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceExtEntityLEDTable.EntityData.Children = types.NewOrderedMap()
    ceExtEntityLEDTable.EntityData.Children.Append("ceExtEntityLEDEntry", types.YChild{"CeExtEntityLEDEntry", nil})
    for i := range ceExtEntityLEDTable.CeExtEntityLEDEntry {
        ceExtEntityLEDTable.EntityData.Children.Append(types.GetSegmentPath(ceExtEntityLEDTable.CeExtEntityLEDEntry[i]), types.YChild{"CeExtEntityLEDEntry", ceExtEntityLEDTable.CeExtEntityLEDEntry[i]})
    }
    ceExtEntityLEDTable.EntityData.Leafs = types.NewOrderedMap()

    ceExtEntityLEDTable.EntityData.YListKeys = []string {}

    return &(ceExtEntityLEDTable.EntityData)
}

// CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry
// An entry (conceptual row) in the ceExtEntityLEDTable,
// containing information about an LED on an entity, identified by 
// entPhysicalIndex.
type CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type of LED on this entity. 'status' -
    // indicates the entity status. 'system' - indicates the overall system
    // status.  'active' - the redundancy status of a module, for e.g.           
    // supervisor module.  'power'  - indicates sufficient power availability for
    // all             modules. 'battery'- indicates the battery status. The type
    // is CeExtEntityLEDType.
    CeExtEntityLEDType interface{}

    // The color of the LED. The type is CeExtEntityLEDColor.
    CeExtEntityLEDColor interface{}
}

func (ceExtEntityLEDEntry *CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry) GetEntityData() *types.CommonEntityData {
    ceExtEntityLEDEntry.EntityData.YFilter = ceExtEntityLEDEntry.YFilter
    ceExtEntityLEDEntry.EntityData.YangName = "ceExtEntityLEDEntry"
    ceExtEntityLEDEntry.EntityData.BundleName = "cisco_ios_xe"
    ceExtEntityLEDEntry.EntityData.ParentYangName = "ceExtEntityLEDTable"
    ceExtEntityLEDEntry.EntityData.SegmentPath = "ceExtEntityLEDEntry" + types.AddKeyToken(ceExtEntityLEDEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(ceExtEntityLEDEntry.CeExtEntityLEDType, "ceExtEntityLEDType")
    ceExtEntityLEDEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceExtEntityLEDEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceExtEntityLEDEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceExtEntityLEDEntry.EntityData.Children = types.NewOrderedMap()
    ceExtEntityLEDEntry.EntityData.Leafs = types.NewOrderedMap()
    ceExtEntityLEDEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", ceExtEntityLEDEntry.EntPhysicalIndex})
    ceExtEntityLEDEntry.EntityData.Leafs.Append("ceExtEntityLEDType", types.YLeaf{"CeExtEntityLEDType", ceExtEntityLEDEntry.CeExtEntityLEDType})
    ceExtEntityLEDEntry.EntityData.Leafs.Append("ceExtEntityLEDColor", types.YLeaf{"CeExtEntityLEDColor", ceExtEntityLEDEntry.CeExtEntityLEDColor})

    ceExtEntityLEDEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CeExtEntityLEDType"}

    return &(ceExtEntityLEDEntry.EntityData)
}

// CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor represents The color of the LED.
type CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor string

const (
    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor_off CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor = "off"

    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor_green CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor = "green"

    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor_amber CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor = "amber"

    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor_red CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDColor = "red"
)

// CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType represents 'battery'- indicates the battery status.
type CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType string

const (
    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType_status CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType = "status"

    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType_system CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType = "system"

    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType_active CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType = "active"

    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType_power CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType = "power"

    CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType_battery CISCOENTITYEXTMIB_CeExtEntityLEDTable_CeExtEntityLEDEntry_CeExtEntityLEDType = "battery"
)

