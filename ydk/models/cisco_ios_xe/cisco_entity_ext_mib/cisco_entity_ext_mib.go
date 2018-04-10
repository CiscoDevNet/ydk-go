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
    Ceextphysicalprocessortable CISCOENTITYEXTMIB_Ceextphysicalprocessortable

    // The ceExtConfigRegTable extends the ENTITY-MIB entPhysicalTable.
    Ceextconfigregtable CISCOENTITYEXTMIB_Ceextconfigregtable

    // A table containing information of LED on an entity.
    Ceextentityledtable CISCOENTITYEXTMIB_Ceextentityledtable
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

    cISCOENTITYEXTMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOENTITYEXTMIB.EntityData.Children["ceExtPhysicalProcessorTable"] = types.YChild{"Ceextphysicalprocessortable", &cISCOENTITYEXTMIB.Ceextphysicalprocessortable}
    cISCOENTITYEXTMIB.EntityData.Children["ceExtConfigRegTable"] = types.YChild{"Ceextconfigregtable", &cISCOENTITYEXTMIB.Ceextconfigregtable}
    cISCOENTITYEXTMIB.EntityData.Children["ceExtEntityLEDTable"] = types.YChild{"Ceextentityledtable", &cISCOENTITYEXTMIB.Ceextentityledtable}
    cISCOENTITYEXTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOENTITYEXTMIB.EntityData)
}

// CISCOENTITYEXTMIB_Ceextphysicalprocessortable
// The ceExtPhysicalProcessorTable extends
// the ENTITY-MIB entPhysicalTable for modules
// (Non FRUs(Field Replacable Units) or FRUs).
type CISCOENTITYEXTMIB_Ceextphysicalprocessortable struct {
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
    // CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry.
    Ceextphysicalprocessorentry []CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry
}

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetEntityData() *types.CommonEntityData {
    ceextphysicalprocessortable.EntityData.YFilter = ceextphysicalprocessortable.YFilter
    ceextphysicalprocessortable.EntityData.YangName = "ceExtPhysicalProcessorTable"
    ceextphysicalprocessortable.EntityData.BundleName = "cisco_ios_xe"
    ceextphysicalprocessortable.EntityData.ParentYangName = "CISCO-ENTITY-EXT-MIB"
    ceextphysicalprocessortable.EntityData.SegmentPath = "ceExtPhysicalProcessorTable"
    ceextphysicalprocessortable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceextphysicalprocessortable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceextphysicalprocessortable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceextphysicalprocessortable.EntityData.Children = make(map[string]types.YChild)
    ceextphysicalprocessortable.EntityData.Children["ceExtPhysicalProcessorEntry"] = types.YChild{"Ceextphysicalprocessorentry", nil}
    for i := range ceextphysicalprocessortable.Ceextphysicalprocessorentry {
        ceextphysicalprocessortable.EntityData.Children[types.GetSegmentPath(&ceextphysicalprocessortable.Ceextphysicalprocessorentry[i])] = types.YChild{"Ceextphysicalprocessorentry", &ceextphysicalprocessortable.Ceextphysicalprocessorentry[i]}
    }
    ceextphysicalprocessortable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceextphysicalprocessortable.EntityData)
}

// CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry
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
type CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // Total number of bytes of RAM available on the Processor. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Ceextprocessorram interface{}

    // Total number of bytes of NVRAM in the entity.  A value of 0 for this object
    // means the entity does not support NVRAM or NVRAM information  is not
    // available. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    Ceextnvramsize interface{}

    // Number of bytes of NVRAM in use. This object is irrelevant if
    // ceExtNVRAMSize is 0. The type is interface{} with range: 0..4294967295.
    // Units are bytes.
    Ceextnvramused interface{}

    // This object represents the upper 32-bit of ceExtProcessorRam. This object
    // needs to be supported only if the available RAM bytes exceeds 32-bit,
    // otherwise this object value would be set to 0. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    Ceextprocessorramoverflow interface{}

    // This object represents the total number of bytes of RAM available on the
    // Processor. This object is a 64-bit version of ceExtProcessorRam. The type
    // is interface{} with range: 0..18446744073709551615. Units are bytes.
    Ceexthcprocessorram interface{}
}

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetEntityData() *types.CommonEntityData {
    ceextphysicalprocessorentry.EntityData.YFilter = ceextphysicalprocessorentry.YFilter
    ceextphysicalprocessorentry.EntityData.YangName = "ceExtPhysicalProcessorEntry"
    ceextphysicalprocessorentry.EntityData.BundleName = "cisco_ios_xe"
    ceextphysicalprocessorentry.EntityData.ParentYangName = "ceExtPhysicalProcessorTable"
    ceextphysicalprocessorentry.EntityData.SegmentPath = "ceExtPhysicalProcessorEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceextphysicalprocessorentry.Entphysicalindex) + "']"
    ceextphysicalprocessorentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceextphysicalprocessorentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceextphysicalprocessorentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceextphysicalprocessorentry.EntityData.Children = make(map[string]types.YChild)
    ceextphysicalprocessorentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceextphysicalprocessorentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceextphysicalprocessorentry.Entphysicalindex}
    ceextphysicalprocessorentry.EntityData.Leafs["ceExtProcessorRam"] = types.YLeaf{"Ceextprocessorram", ceextphysicalprocessorentry.Ceextprocessorram}
    ceextphysicalprocessorentry.EntityData.Leafs["ceExtNVRAMSize"] = types.YLeaf{"Ceextnvramsize", ceextphysicalprocessorentry.Ceextnvramsize}
    ceextphysicalprocessorentry.EntityData.Leafs["ceExtNVRAMUsed"] = types.YLeaf{"Ceextnvramused", ceextphysicalprocessorentry.Ceextnvramused}
    ceextphysicalprocessorentry.EntityData.Leafs["ceExtProcessorRamOverflow"] = types.YLeaf{"Ceextprocessorramoverflow", ceextphysicalprocessorentry.Ceextprocessorramoverflow}
    ceextphysicalprocessorentry.EntityData.Leafs["ceExtHCProcessorRam"] = types.YLeaf{"Ceexthcprocessorram", ceextphysicalprocessorentry.Ceexthcprocessorram}
    return &(ceextphysicalprocessorentry.EntityData)
}

// CISCOENTITYEXTMIB_Ceextconfigregtable
// The ceExtConfigRegTable extends
// the ENTITY-MIB entPhysicalTable.
type CISCOENTITYEXTMIB_Ceextconfigregtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A ceExtConfigRegTable entry extends a corresponding entPhysicalTable entry
    // of class module which has a configuration register.  Entries are created by
    // the agent at the system power-up or module insertion.  Entries are removed
    // when the module is reset or  removed. The type is slice of
    // CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry.
    Ceextconfigregentry []CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry
}

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetEntityData() *types.CommonEntityData {
    ceextconfigregtable.EntityData.YFilter = ceextconfigregtable.YFilter
    ceextconfigregtable.EntityData.YangName = "ceExtConfigRegTable"
    ceextconfigregtable.EntityData.BundleName = "cisco_ios_xe"
    ceextconfigregtable.EntityData.ParentYangName = "CISCO-ENTITY-EXT-MIB"
    ceextconfigregtable.EntityData.SegmentPath = "ceExtConfigRegTable"
    ceextconfigregtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceextconfigregtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceextconfigregtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceextconfigregtable.EntityData.Children = make(map[string]types.YChild)
    ceextconfigregtable.EntityData.Children["ceExtConfigRegEntry"] = types.YChild{"Ceextconfigregentry", nil}
    for i := range ceextconfigregtable.Ceextconfigregentry {
        ceextconfigregtable.EntityData.Children[types.GetSegmentPath(&ceextconfigregtable.Ceextconfigregentry[i])] = types.YChild{"Ceextconfigregentry", &ceextconfigregtable.Ceextconfigregentry[i]}
    }
    ceextconfigregtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceextconfigregtable.EntityData)
}

// CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry
// A ceExtConfigRegTable entry extends
// a corresponding entPhysicalTable entry of class
// module which has a configuration register.
// 
// Entries are created by the agent at the system power-up
// or module insertion.
// 
// Entries are removed when the module is reset or 
// removed.
type CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The value of configuration register with which the processor module booted.
    // The type is string.
    Ceextconfigregister interface{}

    // The value of configuration register in the processor module at next reboot.
    // Just after  the reboot this has the same value as  ceExtConfigRegister. The
    // type is string.
    Ceextconfigregnext interface{}

    // The list of system boot images which can be used for booting. The type is
    // string with length: 0..255.
    Ceextsysbootimagelist interface{}

    // The list of system kickstart images which can be used for booting. The type
    // is string with length: 0..255.
    Ceextkickstartimagelist interface{}
}

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetEntityData() *types.CommonEntityData {
    ceextconfigregentry.EntityData.YFilter = ceextconfigregentry.YFilter
    ceextconfigregentry.EntityData.YangName = "ceExtConfigRegEntry"
    ceextconfigregentry.EntityData.BundleName = "cisco_ios_xe"
    ceextconfigregentry.EntityData.ParentYangName = "ceExtConfigRegTable"
    ceextconfigregentry.EntityData.SegmentPath = "ceExtConfigRegEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceextconfigregentry.Entphysicalindex) + "']"
    ceextconfigregentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceextconfigregentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceextconfigregentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceextconfigregentry.EntityData.Children = make(map[string]types.YChild)
    ceextconfigregentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceextconfigregentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceextconfigregentry.Entphysicalindex}
    ceextconfigregentry.EntityData.Leafs["ceExtConfigRegister"] = types.YLeaf{"Ceextconfigregister", ceextconfigregentry.Ceextconfigregister}
    ceextconfigregentry.EntityData.Leafs["ceExtConfigRegNext"] = types.YLeaf{"Ceextconfigregnext", ceextconfigregentry.Ceextconfigregnext}
    ceextconfigregentry.EntityData.Leafs["ceExtSysBootImageList"] = types.YLeaf{"Ceextsysbootimagelist", ceextconfigregentry.Ceextsysbootimagelist}
    ceextconfigregentry.EntityData.Leafs["ceExtKickstartImageList"] = types.YLeaf{"Ceextkickstartimagelist", ceextconfigregentry.Ceextkickstartimagelist}
    return &(ceextconfigregentry.EntityData)
}

// CISCOENTITYEXTMIB_Ceextentityledtable
// A table containing information of LED on an entity.
type CISCOENTITYEXTMIB_Ceextentityledtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ceExtEntityLEDTable, containing
    // information about an LED on an entity, identified by  entPhysicalIndex. The
    // type is slice of CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry.
    Ceextentityledentry []CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry
}

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetEntityData() *types.CommonEntityData {
    ceextentityledtable.EntityData.YFilter = ceextentityledtable.YFilter
    ceextentityledtable.EntityData.YangName = "ceExtEntityLEDTable"
    ceextentityledtable.EntityData.BundleName = "cisco_ios_xe"
    ceextentityledtable.EntityData.ParentYangName = "CISCO-ENTITY-EXT-MIB"
    ceextentityledtable.EntityData.SegmentPath = "ceExtEntityLEDTable"
    ceextentityledtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceextentityledtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceextentityledtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceextentityledtable.EntityData.Children = make(map[string]types.YChild)
    ceextentityledtable.EntityData.Children["ceExtEntityLEDEntry"] = types.YChild{"Ceextentityledentry", nil}
    for i := range ceextentityledtable.Ceextentityledentry {
        ceextentityledtable.EntityData.Children[types.GetSegmentPath(&ceextentityledtable.Ceextentityledentry[i])] = types.YChild{"Ceextentityledentry", &ceextentityledtable.Ceextentityledentry[i]}
    }
    ceextentityledtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ceextentityledtable.EntityData)
}

// CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry
// An entry (conceptual row) in the ceExtEntityLEDTable,
// containing information about an LED on an entity, identified by 
// entPhysicalIndex.
type CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type of LED on this entity. 'status' -
    // indicates the entity status. 'system' - indicates the overall system
    // status.  'active' - the redundancy status of a module, for e.g.           
    // supervisor module.  'power'  - indicates sufficient power availability for
    // all             modules. 'battery'- indicates the battery status. The type
    // is Ceextentityledtype.
    Ceextentityledtype interface{}

    // The color of the LED. The type is Ceextentityledcolor.
    Ceextentityledcolor interface{}
}

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetEntityData() *types.CommonEntityData {
    ceextentityledentry.EntityData.YFilter = ceextentityledentry.YFilter
    ceextentityledentry.EntityData.YangName = "ceExtEntityLEDEntry"
    ceextentityledentry.EntityData.BundleName = "cisco_ios_xe"
    ceextentityledentry.EntityData.ParentYangName = "ceExtEntityLEDTable"
    ceextentityledentry.EntityData.SegmentPath = "ceExtEntityLEDEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceextentityledentry.Entphysicalindex) + "']" + "[ceExtEntityLEDType='" + fmt.Sprintf("%v", ceextentityledentry.Ceextentityledtype) + "']"
    ceextentityledentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ceextentityledentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ceextentityledentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ceextentityledentry.EntityData.Children = make(map[string]types.YChild)
    ceextentityledentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ceextentityledentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", ceextentityledentry.Entphysicalindex}
    ceextentityledentry.EntityData.Leafs["ceExtEntityLEDType"] = types.YLeaf{"Ceextentityledtype", ceextentityledentry.Ceextentityledtype}
    ceextentityledentry.EntityData.Leafs["ceExtEntityLEDColor"] = types.YLeaf{"Ceextentityledcolor", ceextentityledentry.Ceextentityledcolor}
    return &(ceextentityledentry.EntityData)
}

// CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor represents The color of the LED.
type CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor string

const (
    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor_off CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor = "off"

    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor_green CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor = "green"

    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor_amber CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor = "amber"

    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor_red CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledcolor = "red"
)

// CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype represents 'battery'- indicates the battery status.
type CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype string

const (
    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype_status CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype = "status"

    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype_system CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype = "system"

    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype_active CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype = "active"

    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype_power CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype = "power"

    CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype_battery CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry_Ceextentityledtype = "battery"
)

