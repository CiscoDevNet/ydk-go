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
    parent types.Entity
    YFilter yfilter.YFilter

    // The ceExtPhysicalProcessorTable extends the ENTITY-MIB entPhysicalTable for
    // modules (Non FRUs(Field Replacable Units) or FRUs).
    Ceextphysicalprocessortable CISCOENTITYEXTMIB_Ceextphysicalprocessortable

    // The ceExtConfigRegTable extends the ENTITY-MIB entPhysicalTable.
    Ceextconfigregtable CISCOENTITYEXTMIB_Ceextconfigregtable

    // A table containing information of LED on an entity.
    Ceextentityledtable CISCOENTITYEXTMIB_Ceextentityledtable
}

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetFilter() yfilter.YFilter { return cISCOENTITYEXTMIB.YFilter }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) SetFilter(yf yfilter.YFilter) { cISCOENTITYEXTMIB.YFilter = yf }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetGoName(yname string) string {
    if yname == "ceExtPhysicalProcessorTable" { return "Ceextphysicalprocessortable" }
    if yname == "ceExtConfigRegTable" { return "Ceextconfigregtable" }
    if yname == "ceExtEntityLEDTable" { return "Ceextentityledtable" }
    return ""
}

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetSegmentPath() string {
    return "CISCO-ENTITY-EXT-MIB:CISCO-ENTITY-EXT-MIB"
}

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceExtPhysicalProcessorTable" {
        return &cISCOENTITYEXTMIB.Ceextphysicalprocessortable
    }
    if childYangName == "ceExtConfigRegTable" {
        return &cISCOENTITYEXTMIB.Ceextconfigregtable
    }
    if childYangName == "ceExtEntityLEDTable" {
        return &cISCOENTITYEXTMIB.Ceextentityledtable
    }
    return nil
}

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ceExtPhysicalProcessorTable"] = &cISCOENTITYEXTMIB.Ceextphysicalprocessortable
    children["ceExtConfigRegTable"] = &cISCOENTITYEXTMIB.Ceextconfigregtable
    children["ceExtEntityLEDTable"] = &cISCOENTITYEXTMIB.Ceextentityledtable
    return children
}

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetYangName() string { return "CISCO-ENTITY-EXT-MIB" }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) SetParent(parent types.Entity) { cISCOENTITYEXTMIB.parent = parent }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetParent() types.Entity { return cISCOENTITYEXTMIB.parent }

func (cISCOENTITYEXTMIB *CISCOENTITYEXTMIB) GetParentYangName() string { return "CISCO-ENTITY-EXT-MIB" }

// CISCOENTITYEXTMIB_Ceextphysicalprocessortable
// The ceExtPhysicalProcessorTable extends
// the ENTITY-MIB entPhysicalTable for modules
// (Non FRUs(Field Replacable Units) or FRUs).
type CISCOENTITYEXTMIB_Ceextphysicalprocessortable struct {
    parent types.Entity
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

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetFilter() yfilter.YFilter { return ceextphysicalprocessortable.YFilter }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) SetFilter(yf yfilter.YFilter) { ceextphysicalprocessortable.YFilter = yf }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetGoName(yname string) string {
    if yname == "ceExtPhysicalProcessorEntry" { return "Ceextphysicalprocessorentry" }
    return ""
}

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetSegmentPath() string {
    return "ceExtPhysicalProcessorTable"
}

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceExtPhysicalProcessorEntry" {
        for _, c := range ceextphysicalprocessortable.Ceextphysicalprocessorentry {
            if ceextphysicalprocessortable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry{}
        ceextphysicalprocessortable.Ceextphysicalprocessorentry = append(ceextphysicalprocessortable.Ceextphysicalprocessorentry, child)
        return &ceextphysicalprocessortable.Ceextphysicalprocessorentry[len(ceextphysicalprocessortable.Ceextphysicalprocessorentry)-1]
    }
    return nil
}

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceextphysicalprocessortable.Ceextphysicalprocessorentry {
        children[ceextphysicalprocessortable.Ceextphysicalprocessorentry[i].GetSegmentPath()] = &ceextphysicalprocessortable.Ceextphysicalprocessorentry[i]
    }
    return children
}

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetBundleName() string { return "cisco_ios_xe" }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetYangName() string { return "ceExtPhysicalProcessorTable" }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) SetParent(parent types.Entity) { ceextphysicalprocessortable.parent = parent }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetParent() types.Entity { return ceextphysicalprocessortable.parent }

func (ceextphysicalprocessortable *CISCOENTITYEXTMIB_Ceextphysicalprocessortable) GetParentYangName() string { return "CISCO-ENTITY-EXT-MIB" }

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
    parent types.Entity
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

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetFilter() yfilter.YFilter { return ceextphysicalprocessorentry.YFilter }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) SetFilter(yf yfilter.YFilter) { ceextphysicalprocessorentry.YFilter = yf }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "ceExtProcessorRam" { return "Ceextprocessorram" }
    if yname == "ceExtNVRAMSize" { return "Ceextnvramsize" }
    if yname == "ceExtNVRAMUsed" { return "Ceextnvramused" }
    if yname == "ceExtProcessorRamOverflow" { return "Ceextprocessorramoverflow" }
    if yname == "ceExtHCProcessorRam" { return "Ceexthcprocessorram" }
    return ""
}

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetSegmentPath() string {
    return "ceExtPhysicalProcessorEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceextphysicalprocessorentry.Entphysicalindex) + "']"
}

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceextphysicalprocessorentry.Entphysicalindex
    leafs["ceExtProcessorRam"] = ceextphysicalprocessorentry.Ceextprocessorram
    leafs["ceExtNVRAMSize"] = ceextphysicalprocessorentry.Ceextnvramsize
    leafs["ceExtNVRAMUsed"] = ceextphysicalprocessorentry.Ceextnvramused
    leafs["ceExtProcessorRamOverflow"] = ceextphysicalprocessorentry.Ceextprocessorramoverflow
    leafs["ceExtHCProcessorRam"] = ceextphysicalprocessorentry.Ceexthcprocessorram
    return leafs
}

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetYangName() string { return "ceExtPhysicalProcessorEntry" }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) SetParent(parent types.Entity) { ceextphysicalprocessorentry.parent = parent }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetParent() types.Entity { return ceextphysicalprocessorentry.parent }

func (ceextphysicalprocessorentry *CISCOENTITYEXTMIB_Ceextphysicalprocessortable_Ceextphysicalprocessorentry) GetParentYangName() string { return "ceExtPhysicalProcessorTable" }

// CISCOENTITYEXTMIB_Ceextconfigregtable
// The ceExtConfigRegTable extends
// the ENTITY-MIB entPhysicalTable.
type CISCOENTITYEXTMIB_Ceextconfigregtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A ceExtConfigRegTable entry extends a corresponding entPhysicalTable entry
    // of class module which has a configuration register.  Entries are created by
    // the agent at the system power-up or module insertion.  Entries are removed
    // when the module is reset or  removed. The type is slice of
    // CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry.
    Ceextconfigregentry []CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry
}

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetFilter() yfilter.YFilter { return ceextconfigregtable.YFilter }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) SetFilter(yf yfilter.YFilter) { ceextconfigregtable.YFilter = yf }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetGoName(yname string) string {
    if yname == "ceExtConfigRegEntry" { return "Ceextconfigregentry" }
    return ""
}

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetSegmentPath() string {
    return "ceExtConfigRegTable"
}

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceExtConfigRegEntry" {
        for _, c := range ceextconfigregtable.Ceextconfigregentry {
            if ceextconfigregtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry{}
        ceextconfigregtable.Ceextconfigregentry = append(ceextconfigregtable.Ceextconfigregentry, child)
        return &ceextconfigregtable.Ceextconfigregentry[len(ceextconfigregtable.Ceextconfigregentry)-1]
    }
    return nil
}

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceextconfigregtable.Ceextconfigregentry {
        children[ceextconfigregtable.Ceextconfigregentry[i].GetSegmentPath()] = &ceextconfigregtable.Ceextconfigregentry[i]
    }
    return children
}

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetBundleName() string { return "cisco_ios_xe" }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetYangName() string { return "ceExtConfigRegTable" }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) SetParent(parent types.Entity) { ceextconfigregtable.parent = parent }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetParent() types.Entity { return ceextconfigregtable.parent }

func (ceextconfigregtable *CISCOENTITYEXTMIB_Ceextconfigregtable) GetParentYangName() string { return "CISCO-ENTITY-EXT-MIB" }

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
    parent types.Entity
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

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetFilter() yfilter.YFilter { return ceextconfigregentry.YFilter }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) SetFilter(yf yfilter.YFilter) { ceextconfigregentry.YFilter = yf }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "ceExtConfigRegister" { return "Ceextconfigregister" }
    if yname == "ceExtConfigRegNext" { return "Ceextconfigregnext" }
    if yname == "ceExtSysBootImageList" { return "Ceextsysbootimagelist" }
    if yname == "ceExtKickstartImageList" { return "Ceextkickstartimagelist" }
    return ""
}

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetSegmentPath() string {
    return "ceExtConfigRegEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceextconfigregentry.Entphysicalindex) + "']"
}

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceextconfigregentry.Entphysicalindex
    leafs["ceExtConfigRegister"] = ceextconfigregentry.Ceextconfigregister
    leafs["ceExtConfigRegNext"] = ceextconfigregentry.Ceextconfigregnext
    leafs["ceExtSysBootImageList"] = ceextconfigregentry.Ceextsysbootimagelist
    leafs["ceExtKickstartImageList"] = ceextconfigregentry.Ceextkickstartimagelist
    return leafs
}

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetYangName() string { return "ceExtConfigRegEntry" }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) SetParent(parent types.Entity) { ceextconfigregentry.parent = parent }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetParent() types.Entity { return ceextconfigregentry.parent }

func (ceextconfigregentry *CISCOENTITYEXTMIB_Ceextconfigregtable_Ceextconfigregentry) GetParentYangName() string { return "ceExtConfigRegTable" }

// CISCOENTITYEXTMIB_Ceextentityledtable
// A table containing information of LED on an entity.
type CISCOENTITYEXTMIB_Ceextentityledtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry (conceptual row) in the ceExtEntityLEDTable, containing
    // information about an LED on an entity, identified by  entPhysicalIndex. The
    // type is slice of CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry.
    Ceextentityledentry []CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry
}

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetFilter() yfilter.YFilter { return ceextentityledtable.YFilter }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) SetFilter(yf yfilter.YFilter) { ceextentityledtable.YFilter = yf }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetGoName(yname string) string {
    if yname == "ceExtEntityLEDEntry" { return "Ceextentityledentry" }
    return ""
}

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetSegmentPath() string {
    return "ceExtEntityLEDTable"
}

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ceExtEntityLEDEntry" {
        for _, c := range ceextentityledtable.Ceextentityledentry {
            if ceextentityledtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry{}
        ceextentityledtable.Ceextentityledentry = append(ceextentityledtable.Ceextentityledentry, child)
        return &ceextentityledtable.Ceextentityledentry[len(ceextentityledtable.Ceextentityledentry)-1]
    }
    return nil
}

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ceextentityledtable.Ceextentityledentry {
        children[ceextentityledtable.Ceextentityledentry[i].GetSegmentPath()] = &ceextentityledtable.Ceextentityledentry[i]
    }
    return children
}

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetBundleName() string { return "cisco_ios_xe" }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetYangName() string { return "ceExtEntityLEDTable" }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) SetParent(parent types.Entity) { ceextentityledtable.parent = parent }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetParent() types.Entity { return ceextentityledtable.parent }

func (ceextentityledtable *CISCOENTITYEXTMIB_Ceextentityledtable) GetParentYangName() string { return "CISCO-ENTITY-EXT-MIB" }

// CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry
// An entry (conceptual row) in the ceExtEntityLEDTable,
// containing information about an LED on an entity, identified by 
// entPhysicalIndex.
type CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry struct {
    parent types.Entity
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

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetFilter() yfilter.YFilter { return ceextentityledentry.YFilter }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) SetFilter(yf yfilter.YFilter) { ceextentityledentry.YFilter = yf }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "ceExtEntityLEDType" { return "Ceextentityledtype" }
    if yname == "ceExtEntityLEDColor" { return "Ceextentityledcolor" }
    return ""
}

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetSegmentPath() string {
    return "ceExtEntityLEDEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", ceextentityledentry.Entphysicalindex) + "']" + "[ceExtEntityLEDType='" + fmt.Sprintf("%v", ceextentityledentry.Ceextentityledtype) + "']"
}

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = ceextentityledentry.Entphysicalindex
    leafs["ceExtEntityLEDType"] = ceextentityledentry.Ceextentityledtype
    leafs["ceExtEntityLEDColor"] = ceextentityledentry.Ceextentityledcolor
    return leafs
}

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetBundleName() string { return "cisco_ios_xe" }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetYangName() string { return "ceExtEntityLEDEntry" }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) SetParent(parent types.Entity) { ceextentityledentry.parent = parent }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetParent() types.Entity { return ceextentityledentry.parent }

func (ceextentityledentry *CISCOENTITYEXTMIB_Ceextentityledtable_Ceextentityledentry) GetParentYangName() string { return "ceExtEntityLEDTable" }

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

