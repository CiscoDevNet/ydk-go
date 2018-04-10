// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// Copyright(c) 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_entity_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_entity_mib"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-entity-mib ENTITY-MIB}", reflect.TypeOf(ENTITYMIB{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB", reflect.TypeOf(ENTITYMIB{}))
}

// PhysicalClass
type PhysicalClass string

const (
    PhysicalClass_other PhysicalClass = "other"

    PhysicalClass_unknown PhysicalClass = "unknown"

    PhysicalClass_chassis PhysicalClass = "chassis"

    PhysicalClass_backplane PhysicalClass = "backplane"

    PhysicalClass_container PhysicalClass = "container"

    PhysicalClass_powerSupply PhysicalClass = "powerSupply"

    PhysicalClass_fan PhysicalClass = "fan"

    PhysicalClass_sensor PhysicalClass = "sensor"

    PhysicalClass_module PhysicalClass = "module"

    PhysicalClass_port PhysicalClass = "port"

    PhysicalClass_stack PhysicalClass = "stack"
)

// ENTITYMIB
type ENTITYMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Entitygeneral ENTITYMIB_Entitygeneral

    
    Entphysicaltable ENTITYMIB_Entphysicaltable

    
    Entlogicaltable ENTITYMIB_Entlogicaltable

    
    Entlpmappingtable ENTITYMIB_Entlpmappingtable

    
    Entaliasmappingtable ENTITYMIB_Entaliasmappingtable

    
    Entphysicalcontainstable ENTITYMIB_Entphysicalcontainstable
}

func (eNTITYMIB *ENTITYMIB) GetEntityData() *types.CommonEntityData {
    eNTITYMIB.EntityData.YFilter = eNTITYMIB.YFilter
    eNTITYMIB.EntityData.YangName = "ENTITY-MIB"
    eNTITYMIB.EntityData.BundleName = "cisco_ios_xr"
    eNTITYMIB.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-entity-mib"
    eNTITYMIB.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-entity-mib:ENTITY-MIB"
    eNTITYMIB.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eNTITYMIB.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eNTITYMIB.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eNTITYMIB.EntityData.Children = make(map[string]types.YChild)
    eNTITYMIB.EntityData.Children["entityGeneral"] = types.YChild{"Entitygeneral", &eNTITYMIB.Entitygeneral}
    eNTITYMIB.EntityData.Children["entPhysicalTable"] = types.YChild{"Entphysicaltable", &eNTITYMIB.Entphysicaltable}
    eNTITYMIB.EntityData.Children["entLogicalTable"] = types.YChild{"Entlogicaltable", &eNTITYMIB.Entlogicaltable}
    eNTITYMIB.EntityData.Children["entLPMappingTable"] = types.YChild{"Entlpmappingtable", &eNTITYMIB.Entlpmappingtable}
    eNTITYMIB.EntityData.Children["entAliasMappingTable"] = types.YChild{"Entaliasmappingtable", &eNTITYMIB.Entaliasmappingtable}
    eNTITYMIB.EntityData.Children["entPhysicalContainsTable"] = types.YChild{"Entphysicalcontainstable", &eNTITYMIB.Entphysicalcontainstable}
    eNTITYMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eNTITYMIB.EntityData)
}

// ENTITYMIB_Entitygeneral
type ENTITYMIB_Entitygeneral struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    Entlastchangetime interface{}
}

func (entitygeneral *ENTITYMIB_Entitygeneral) GetEntityData() *types.CommonEntityData {
    entitygeneral.EntityData.YFilter = entitygeneral.YFilter
    entitygeneral.EntityData.YangName = "entityGeneral"
    entitygeneral.EntityData.BundleName = "cisco_ios_xr"
    entitygeneral.EntityData.ParentYangName = "ENTITY-MIB"
    entitygeneral.EntityData.SegmentPath = "entityGeneral"
    entitygeneral.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entitygeneral.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entitygeneral.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entitygeneral.EntityData.Children = make(map[string]types.YChild)
    entitygeneral.EntityData.Leafs = make(map[string]types.YLeaf)
    entitygeneral.EntityData.Leafs["entLastChangeTime"] = types.YLeaf{"Entlastchangetime", entitygeneral.Entlastchangetime}
    return &(entitygeneral.EntityData)
}

// ENTITYMIB_Entphysicaltable
type ENTITYMIB_Entphysicaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYMIB_Entphysicaltable_Entphysicalentry.
    Entphysicalentry []ENTITYMIB_Entphysicaltable_Entphysicalentry
}

func (entphysicaltable *ENTITYMIB_Entphysicaltable) GetEntityData() *types.CommonEntityData {
    entphysicaltable.EntityData.YFilter = entphysicaltable.YFilter
    entphysicaltable.EntityData.YangName = "entPhysicalTable"
    entphysicaltable.EntityData.BundleName = "cisco_ios_xr"
    entphysicaltable.EntityData.ParentYangName = "ENTITY-MIB"
    entphysicaltable.EntityData.SegmentPath = "entPhysicalTable"
    entphysicaltable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entphysicaltable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entphysicaltable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entphysicaltable.EntityData.Children = make(map[string]types.YChild)
    entphysicaltable.EntityData.Children["entPhysicalEntry"] = types.YChild{"Entphysicalentry", nil}
    for i := range entphysicaltable.Entphysicalentry {
        entphysicaltable.EntityData.Children[types.GetSegmentPath(&entphysicaltable.Entphysicalentry[i])] = types.YChild{"Entphysicalentry", &entphysicaltable.Entphysicalentry[i]}
    }
    entphysicaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entphysicaltable.EntityData)
}

// ENTITYMIB_Entphysicaltable_Entphysicalentry
type ENTITYMIB_Entphysicaltable_Entphysicalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entphysicalindex interface{}

    // The type is string with length: 0..255.
    Entphysicaldescr interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Entphysicalvendortype interface{}

    // The type is interface{} with range: 0..2147483647.
    Entphysicalcontainedin interface{}

    // The type is PhysicalClass.
    Entphysicalclass interface{}

    // The type is interface{} with range: -1..2147483647.
    Entphysicalparentrelpos interface{}

    // The type is string with length: 0..255.
    Entphysicalname interface{}

    // The type is string with length: 0..255.
    Entphysicalhardwarerev interface{}

    // The type is string with length: 0..255.
    Entphysicalfirmwarerev interface{}

    // The type is string with length: 0..255.
    Entphysicalsoftwarerev interface{}

    // The type is string with length: 0..32.
    Entphysicalserialnum interface{}

    // The type is string with length: 0..255.
    Entphysicalmfgname interface{}

    // The type is string with length: 0..255.
    Entphysicalmodelname interface{}

    // The type is string with length: 0..32.
    Entphysicalalias interface{}

    // The type is string with length: 0..32.
    Entphysicalassetid interface{}

    // The type is TruthValue.
    Entphysicalisfru interface{}
}

func (entphysicalentry *ENTITYMIB_Entphysicaltable_Entphysicalentry) GetEntityData() *types.CommonEntityData {
    entphysicalentry.EntityData.YFilter = entphysicalentry.YFilter
    entphysicalentry.EntityData.YangName = "entPhysicalEntry"
    entphysicalentry.EntityData.BundleName = "cisco_ios_xr"
    entphysicalentry.EntityData.ParentYangName = "entPhysicalTable"
    entphysicalentry.EntityData.SegmentPath = "entPhysicalEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entphysicalentry.Entphysicalindex) + "']"
    entphysicalentry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entphysicalentry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entphysicalentry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entphysicalentry.EntityData.Children = make(map[string]types.YChild)
    entphysicalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entphysicalentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entphysicalentry.Entphysicalindex}
    entphysicalentry.EntityData.Leafs["entPhysicalDescr"] = types.YLeaf{"Entphysicaldescr", entphysicalentry.Entphysicaldescr}
    entphysicalentry.EntityData.Leafs["entPhysicalVendorType"] = types.YLeaf{"Entphysicalvendortype", entphysicalentry.Entphysicalvendortype}
    entphysicalentry.EntityData.Leafs["entPhysicalContainedIn"] = types.YLeaf{"Entphysicalcontainedin", entphysicalentry.Entphysicalcontainedin}
    entphysicalentry.EntityData.Leafs["entPhysicalClass"] = types.YLeaf{"Entphysicalclass", entphysicalentry.Entphysicalclass}
    entphysicalentry.EntityData.Leafs["entPhysicalParentRelPos"] = types.YLeaf{"Entphysicalparentrelpos", entphysicalentry.Entphysicalparentrelpos}
    entphysicalentry.EntityData.Leafs["entPhysicalName"] = types.YLeaf{"Entphysicalname", entphysicalentry.Entphysicalname}
    entphysicalentry.EntityData.Leafs["entPhysicalHardwareRev"] = types.YLeaf{"Entphysicalhardwarerev", entphysicalentry.Entphysicalhardwarerev}
    entphysicalentry.EntityData.Leafs["entPhysicalFirmwareRev"] = types.YLeaf{"Entphysicalfirmwarerev", entphysicalentry.Entphysicalfirmwarerev}
    entphysicalentry.EntityData.Leafs["entPhysicalSoftwareRev"] = types.YLeaf{"Entphysicalsoftwarerev", entphysicalentry.Entphysicalsoftwarerev}
    entphysicalentry.EntityData.Leafs["entPhysicalSerialNum"] = types.YLeaf{"Entphysicalserialnum", entphysicalentry.Entphysicalserialnum}
    entphysicalentry.EntityData.Leafs["entPhysicalMfgName"] = types.YLeaf{"Entphysicalmfgname", entphysicalentry.Entphysicalmfgname}
    entphysicalentry.EntityData.Leafs["entPhysicalModelName"] = types.YLeaf{"Entphysicalmodelname", entphysicalentry.Entphysicalmodelname}
    entphysicalentry.EntityData.Leafs["entPhysicalAlias"] = types.YLeaf{"Entphysicalalias", entphysicalentry.Entphysicalalias}
    entphysicalentry.EntityData.Leafs["entPhysicalAssetID"] = types.YLeaf{"Entphysicalassetid", entphysicalentry.Entphysicalassetid}
    entphysicalentry.EntityData.Leafs["entPhysicalIsFRU"] = types.YLeaf{"Entphysicalisfru", entphysicalentry.Entphysicalisfru}
    return &(entphysicalentry.EntityData)
}

// ENTITYMIB_Entlogicaltable
type ENTITYMIB_Entlogicaltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYMIB_Entlogicaltable_Entlogicalentry.
    Entlogicalentry []ENTITYMIB_Entlogicaltable_Entlogicalentry
}

func (entlogicaltable *ENTITYMIB_Entlogicaltable) GetEntityData() *types.CommonEntityData {
    entlogicaltable.EntityData.YFilter = entlogicaltable.YFilter
    entlogicaltable.EntityData.YangName = "entLogicalTable"
    entlogicaltable.EntityData.BundleName = "cisco_ios_xr"
    entlogicaltable.EntityData.ParentYangName = "ENTITY-MIB"
    entlogicaltable.EntityData.SegmentPath = "entLogicalTable"
    entlogicaltable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entlogicaltable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entlogicaltable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entlogicaltable.EntityData.Children = make(map[string]types.YChild)
    entlogicaltable.EntityData.Children["entLogicalEntry"] = types.YChild{"Entlogicalentry", nil}
    for i := range entlogicaltable.Entlogicalentry {
        entlogicaltable.EntityData.Children[types.GetSegmentPath(&entlogicaltable.Entlogicalentry[i])] = types.YChild{"Entlogicalentry", &entlogicaltable.Entlogicalentry[i]}
    }
    entlogicaltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entlogicaltable.EntityData)
}

// ENTITYMIB_Entlogicaltable_Entlogicalentry
type ENTITYMIB_Entlogicaltable_Entlogicalentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entlogicalindex interface{}

    // The type is string with length: 0..255.
    Entlogicaldescr interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Entlogicaltype interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    Entlogicalcommunity interface{}

    // The type is string with pattern: b'(\\d*(.\\d*)*)?'.
    Entlogicaltaddress interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Entlogicaltdomain interface{}

    // The type is string with pattern:
    // b'(([0-9a-fA-F]){2}(:([0-9a-fA-F]){2})*)?'.
    Entlogicalcontextengineid interface{}

    // The type is string with length: 0..255.
    Entlogicalcontextname interface{}
}

func (entlogicalentry *ENTITYMIB_Entlogicaltable_Entlogicalentry) GetEntityData() *types.CommonEntityData {
    entlogicalentry.EntityData.YFilter = entlogicalentry.YFilter
    entlogicalentry.EntityData.YangName = "entLogicalEntry"
    entlogicalentry.EntityData.BundleName = "cisco_ios_xr"
    entlogicalentry.EntityData.ParentYangName = "entLogicalTable"
    entlogicalentry.EntityData.SegmentPath = "entLogicalEntry" + "[entLogicalIndex='" + fmt.Sprintf("%v", entlogicalentry.Entlogicalindex) + "']"
    entlogicalentry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entlogicalentry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entlogicalentry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entlogicalentry.EntityData.Children = make(map[string]types.YChild)
    entlogicalentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entlogicalentry.EntityData.Leafs["entLogicalIndex"] = types.YLeaf{"Entlogicalindex", entlogicalentry.Entlogicalindex}
    entlogicalentry.EntityData.Leafs["entLogicalDescr"] = types.YLeaf{"Entlogicaldescr", entlogicalentry.Entlogicaldescr}
    entlogicalentry.EntityData.Leafs["entLogicalType"] = types.YLeaf{"Entlogicaltype", entlogicalentry.Entlogicaltype}
    entlogicalentry.EntityData.Leafs["entLogicalCommunity"] = types.YLeaf{"Entlogicalcommunity", entlogicalentry.Entlogicalcommunity}
    entlogicalentry.EntityData.Leafs["entLogicalTAddress"] = types.YLeaf{"Entlogicaltaddress", entlogicalentry.Entlogicaltaddress}
    entlogicalentry.EntityData.Leafs["entLogicalTDomain"] = types.YLeaf{"Entlogicaltdomain", entlogicalentry.Entlogicaltdomain}
    entlogicalentry.EntityData.Leafs["entLogicalContextEngineID"] = types.YLeaf{"Entlogicalcontextengineid", entlogicalentry.Entlogicalcontextengineid}
    entlogicalentry.EntityData.Leafs["entLogicalContextName"] = types.YLeaf{"Entlogicalcontextname", entlogicalentry.Entlogicalcontextname}
    return &(entlogicalentry.EntityData)
}

// ENTITYMIB_Entlpmappingtable
type ENTITYMIB_Entlpmappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYMIB_Entlpmappingtable_Entlpmappingentry.
    Entlpmappingentry []ENTITYMIB_Entlpmappingtable_Entlpmappingentry
}

func (entlpmappingtable *ENTITYMIB_Entlpmappingtable) GetEntityData() *types.CommonEntityData {
    entlpmappingtable.EntityData.YFilter = entlpmappingtable.YFilter
    entlpmappingtable.EntityData.YangName = "entLPMappingTable"
    entlpmappingtable.EntityData.BundleName = "cisco_ios_xr"
    entlpmappingtable.EntityData.ParentYangName = "ENTITY-MIB"
    entlpmappingtable.EntityData.SegmentPath = "entLPMappingTable"
    entlpmappingtable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entlpmappingtable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entlpmappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entlpmappingtable.EntityData.Children = make(map[string]types.YChild)
    entlpmappingtable.EntityData.Children["entLPMappingEntry"] = types.YChild{"Entlpmappingentry", nil}
    for i := range entlpmappingtable.Entlpmappingentry {
        entlpmappingtable.EntityData.Children[types.GetSegmentPath(&entlpmappingtable.Entlpmappingentry[i])] = types.YChild{"Entlpmappingentry", &entlpmappingtable.Entlpmappingentry[i]}
    }
    entlpmappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entlpmappingtable.EntityData)
}

// ENTITYMIB_Entlpmappingtable_Entlpmappingentry
type ENTITYMIB_Entlpmappingtable_Entlpmappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entlogicalindex interface{}

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entlpphysicalindex interface{}
}

func (entlpmappingentry *ENTITYMIB_Entlpmappingtable_Entlpmappingentry) GetEntityData() *types.CommonEntityData {
    entlpmappingentry.EntityData.YFilter = entlpmappingentry.YFilter
    entlpmappingentry.EntityData.YangName = "entLPMappingEntry"
    entlpmappingentry.EntityData.BundleName = "cisco_ios_xr"
    entlpmappingentry.EntityData.ParentYangName = "entLPMappingTable"
    entlpmappingentry.EntityData.SegmentPath = "entLPMappingEntry" + "[entLogicalIndex='" + fmt.Sprintf("%v", entlpmappingentry.Entlogicalindex) + "']" + "[entLPPhysicalIndex='" + fmt.Sprintf("%v", entlpmappingentry.Entlpphysicalindex) + "']"
    entlpmappingentry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entlpmappingentry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entlpmappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entlpmappingentry.EntityData.Children = make(map[string]types.YChild)
    entlpmappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entlpmappingentry.EntityData.Leafs["entLogicalIndex"] = types.YLeaf{"Entlogicalindex", entlpmappingentry.Entlogicalindex}
    entlpmappingentry.EntityData.Leafs["entLPPhysicalIndex"] = types.YLeaf{"Entlpphysicalindex", entlpmappingentry.Entlpphysicalindex}
    return &(entlpmappingentry.EntityData)
}

// ENTITYMIB_Entaliasmappingtable
type ENTITYMIB_Entaliasmappingtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry.
    Entaliasmappingentry []ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry
}

func (entaliasmappingtable *ENTITYMIB_Entaliasmappingtable) GetEntityData() *types.CommonEntityData {
    entaliasmappingtable.EntityData.YFilter = entaliasmappingtable.YFilter
    entaliasmappingtable.EntityData.YangName = "entAliasMappingTable"
    entaliasmappingtable.EntityData.BundleName = "cisco_ios_xr"
    entaliasmappingtable.EntityData.ParentYangName = "ENTITY-MIB"
    entaliasmappingtable.EntityData.SegmentPath = "entAliasMappingTable"
    entaliasmappingtable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entaliasmappingtable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entaliasmappingtable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entaliasmappingtable.EntityData.Children = make(map[string]types.YChild)
    entaliasmappingtable.EntityData.Children["entAliasMappingEntry"] = types.YChild{"Entaliasmappingentry", nil}
    for i := range entaliasmappingtable.Entaliasmappingentry {
        entaliasmappingtable.EntityData.Children[types.GetSegmentPath(&entaliasmappingtable.Entaliasmappingentry[i])] = types.YChild{"Entaliasmappingentry", &entaliasmappingtable.Entaliasmappingentry[i]}
    }
    entaliasmappingtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entaliasmappingtable.EntityData)
}

// ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry
type ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entphysicalindex interface{}

    // This attribute is a key. The type is interface{} with range: 0..2147483647.
    Entaliaslogicalindexorzero interface{}

    // The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Entaliasmappingidentifier interface{}
}

func (entaliasmappingentry *ENTITYMIB_Entaliasmappingtable_Entaliasmappingentry) GetEntityData() *types.CommonEntityData {
    entaliasmappingentry.EntityData.YFilter = entaliasmappingentry.YFilter
    entaliasmappingentry.EntityData.YangName = "entAliasMappingEntry"
    entaliasmappingentry.EntityData.BundleName = "cisco_ios_xr"
    entaliasmappingentry.EntityData.ParentYangName = "entAliasMappingTable"
    entaliasmappingentry.EntityData.SegmentPath = "entAliasMappingEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entaliasmappingentry.Entphysicalindex) + "']" + "[entAliasLogicalIndexOrZero='" + fmt.Sprintf("%v", entaliasmappingentry.Entaliaslogicalindexorzero) + "']"
    entaliasmappingentry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entaliasmappingentry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entaliasmappingentry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entaliasmappingentry.EntityData.Children = make(map[string]types.YChild)
    entaliasmappingentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entaliasmappingentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entaliasmappingentry.Entphysicalindex}
    entaliasmappingentry.EntityData.Leafs["entAliasLogicalIndexOrZero"] = types.YLeaf{"Entaliaslogicalindexorzero", entaliasmappingentry.Entaliaslogicalindexorzero}
    entaliasmappingentry.EntityData.Leafs["entAliasMappingIdentifier"] = types.YLeaf{"Entaliasmappingidentifier", entaliasmappingentry.Entaliasmappingidentifier}
    return &(entaliasmappingentry.EntityData)
}

// ENTITYMIB_Entphysicalcontainstable
type ENTITYMIB_Entphysicalcontainstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry.
    Entphysicalcontainsentry []ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry
}

func (entphysicalcontainstable *ENTITYMIB_Entphysicalcontainstable) GetEntityData() *types.CommonEntityData {
    entphysicalcontainstable.EntityData.YFilter = entphysicalcontainstable.YFilter
    entphysicalcontainstable.EntityData.YangName = "entPhysicalContainsTable"
    entphysicalcontainstable.EntityData.BundleName = "cisco_ios_xr"
    entphysicalcontainstable.EntityData.ParentYangName = "ENTITY-MIB"
    entphysicalcontainstable.EntityData.SegmentPath = "entPhysicalContainsTable"
    entphysicalcontainstable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entphysicalcontainstable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entphysicalcontainstable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entphysicalcontainstable.EntityData.Children = make(map[string]types.YChild)
    entphysicalcontainstable.EntityData.Children["entPhysicalContainsEntry"] = types.YChild{"Entphysicalcontainsentry", nil}
    for i := range entphysicalcontainstable.Entphysicalcontainsentry {
        entphysicalcontainstable.EntityData.Children[types.GetSegmentPath(&entphysicalcontainstable.Entphysicalcontainsentry[i])] = types.YChild{"Entphysicalcontainsentry", &entphysicalcontainstable.Entphysicalcontainsentry[i]}
    }
    entphysicalcontainstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entphysicalcontainstable.EntityData)
}

// ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry
type ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entphysicalindex interface{}

    // This attribute is a key. The type is interface{} with range: 1..2147483647.
    Entphysicalchildindex interface{}
}

func (entphysicalcontainsentry *ENTITYMIB_Entphysicalcontainstable_Entphysicalcontainsentry) GetEntityData() *types.CommonEntityData {
    entphysicalcontainsentry.EntityData.YFilter = entphysicalcontainsentry.YFilter
    entphysicalcontainsentry.EntityData.YangName = "entPhysicalContainsEntry"
    entphysicalcontainsentry.EntityData.BundleName = "cisco_ios_xr"
    entphysicalcontainsentry.EntityData.ParentYangName = "entPhysicalContainsTable"
    entphysicalcontainsentry.EntityData.SegmentPath = "entPhysicalContainsEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entphysicalcontainsentry.Entphysicalindex) + "']" + "[entPhysicalChildIndex='" + fmt.Sprintf("%v", entphysicalcontainsentry.Entphysicalchildindex) + "']"
    entphysicalcontainsentry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entphysicalcontainsentry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entphysicalcontainsentry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entphysicalcontainsentry.EntityData.Children = make(map[string]types.YChild)
    entphysicalcontainsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    entphysicalcontainsentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", entphysicalcontainsentry.Entphysicalindex}
    entphysicalcontainsentry.EntityData.Leafs["entPhysicalChildIndex"] = types.YLeaf{"Entphysicalchildindex", entphysicalcontainsentry.Entphysicalchildindex}
    return &(entphysicalcontainsentry.EntityData)
}

