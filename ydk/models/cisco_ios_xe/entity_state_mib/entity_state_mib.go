// This MIB defines a state extension to the Entity MIB.
// 
// Copyright (C) The Internet Society 2005.  This version
// of this MIB module is part of RFC 4268; see the RFC
// itself for full legal notices.
package entity_state_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package entity_state_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:ENTITY-STATE-MIB ENTITY-STATE-MIB}", reflect.TypeOf(ENTITYSTATEMIB{}))
    ydk.RegisterEntity("ENTITY-STATE-MIB:ENTITY-STATE-MIB", reflect.TypeOf(ENTITYSTATEMIB{}))
}

// ENTITYSTATEMIB
type ENTITYSTATEMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table of information about state/status of entities. This is a sparse
    // augment of the entPhysicalTable.  Entries appear in this table for values
    // of entPhysicalClass [RFC4133] that in this implementation are able to
    // report any of the state or status stored in this table.
    Entstatetable ENTITYSTATEMIB_Entstatetable
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetFilter() yfilter.YFilter { return eNTITYSTATEMIB.YFilter }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) SetFilter(yf yfilter.YFilter) { eNTITYSTATEMIB.YFilter = yf }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetGoName(yname string) string {
    if yname == "entStateTable" { return "Entstatetable" }
    return ""
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetSegmentPath() string {
    return "ENTITY-STATE-MIB:ENTITY-STATE-MIB"
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entStateTable" {
        return &eNTITYSTATEMIB.Entstatetable
    }
    return nil
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["entStateTable"] = &eNTITYSTATEMIB.Entstatetable
    return children
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetYangName() string { return "ENTITY-STATE-MIB" }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) SetParent(parent types.Entity) { eNTITYSTATEMIB.parent = parent }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetParent() types.Entity { return eNTITYSTATEMIB.parent }

func (eNTITYSTATEMIB *ENTITYSTATEMIB) GetParentYangName() string { return "ENTITY-STATE-MIB" }

// ENTITYSTATEMIB_Entstatetable
// A table of information about state/status of entities.
// This is a sparse augment of the entPhysicalTable.  Entries
// appear in this table for values of
// entPhysicalClass [RFC4133] that in this implementation
// are able to report any of the state or status stored in
// this table.
type ENTITYSTATEMIB_Entstatetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // State information about this physical entity. The type is slice of
    // ENTITYSTATEMIB_Entstatetable_Entstateentry.
    Entstateentry []ENTITYSTATEMIB_Entstatetable_Entstateentry
}

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetFilter() yfilter.YFilter { return entstatetable.YFilter }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) SetFilter(yf yfilter.YFilter) { entstatetable.YFilter = yf }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetGoName(yname string) string {
    if yname == "entStateEntry" { return "Entstateentry" }
    return ""
}

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetSegmentPath() string {
    return "entStateTable"
}

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entStateEntry" {
        for _, c := range entstatetable.Entstateentry {
            if entstatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ENTITYSTATEMIB_Entstatetable_Entstateentry{}
        entstatetable.Entstateentry = append(entstatetable.Entstateentry, child)
        return &entstatetable.Entstateentry[len(entstatetable.Entstateentry)-1]
    }
    return nil
}

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entstatetable.Entstateentry {
        children[entstatetable.Entstateentry[i].GetSegmentPath()] = &entstatetable.Entstateentry[i]
    }
    return children
}

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetBundleName() string { return "cisco_ios_xe" }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetYangName() string { return "entStateTable" }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) SetParent(parent types.Entity) { entstatetable.parent = parent }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetParent() types.Entity { return entstatetable.parent }

func (entstatetable *ENTITYSTATEMIB_Entstatetable) GetParentYangName() string { return "ENTITY-STATE-MIB" }

// ENTITYSTATEMIB_Entstatetable_Entstateentry
// State information about this physical entity.
type ENTITYSTATEMIB_Entstatetable_Entstateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // The value of this object is the date and time when the value of any of
    // entStateAdmin, entStateOper, entStateUsage, entStateAlarm, or
    // entStateStandby changed for this entity.  If there has been no change since
    // the last re-initialization of the local system, this object contains the
    // date and time of local system initialization.  If there has been no change
    // since the entity was added to the local system, this object contains the
    // date and time of the insertion. The type is string.
    Entstatelastchanged interface{}

    // The administrative state for this entity.  This object refers to an
    // entities administrative permission to service both other entities within
    // its containment hierarchy as well other users of its services defined by
    // means outside the scope of this MIB.  Setting this object to 'notSupported'
    // will result in an 'inconsistentValue' error.  For entities that do not
    // support administrative state, all set operations will result in an
    // 'inconsistentValue' error.  Some physical entities exhibit only a subset of
    // the remaining administrative state values.  Some entities cannot be locked,
    // and hence this object exhibits only the 'unlocked' state.  Other entities
    // cannot be shutdown gracefully, and hence this object does not exhibit the
    // 'shuttingDown' state.  A value of 'inconsistentValue' will be returned if
    // attempts are made to set this object to values not supported by its
    // administrative model. The type is EntityAdminState.
    Entstateadmin interface{}

    // The operational state for this entity.  Note that unlike the state model
    // used within the Interfaces MIB [RFC2863], this object does not follow the
    // administrative state.  An administrative state of down does not predict an
    // operational state of disabled.  A value of 'testing' means that entity
    // currently being tested and cannot therefore report whether it is
    // operational or not.  A value of 'disabled' means that an entity is totally
    // inoperable and unable to provide service both to entities within its
    // containment hierarchy, or to other receivers of its service as defined in
    // ways outside the scope of this MIB.  A value of 'enabled' means that an
    // entity is fully or partially operable and able to provide service both to 
    // entities within its containment hierarchy, or to other receivers of its
    // service as defined in ways outside the scope of this MIB.  Note that some
    // implementations may not be able to accurately report entStateOper while the
    // entStateAdmin object has a value other than 'unlocked'. In these cases,
    // this object MUST have a value of 'unknown'. The type is EntityOperState.
    Entstateoper interface{}

    // The usage state for this entity.  This object refers to an entity's ability
    // to service more physical entities in a containment hierarchy.  A value of
    // 'idle' means this entity is able to contain other entities but that no
    // other entity is currently contained within this entity.  A value of
    // 'active' means that at least one entity is contained within this entity,
    // but that it could handle more.  A value of 'busy' means that the entity is
    // unable to handle any additional entities being contained in it.  Some
    // entities will exhibit only a subset of the usage state values.  Entities
    // that are unable to ever service any entities within a containment hierarchy
    // will always have a usage state of 'busy'.  Some entities will only ever be
    // able to support one entity within its containment hierarchy and will
    // therefore only exhibit values of 'idle' and 'busy'. The type is
    // EntityUsageState.
    Entstateusage interface{}

    // The alarm status for this entity.  It does not include the alarms raised on
    // child components within its containment hierarchy.  A value of 'unknown'
    // means that this entity is  unable to report alarm state.  Note that this
    // differs from 'indeterminate', which means that alarm state is supported and
    // there are alarms against this entity, but the severity of some of the
    // alarms is not known.  If no bits are set, then this entity supports
    // reporting of alarms, but there are currently no active alarms against this
    // entity. The type is map[string]bool.
    Entstatealarm interface{}

    // The standby status for this entity.  Some entities will exhibit only a
    // subset of the remaining standby state values.  If this entity cannot
    // operate in a standby role, the value of this object will always be
    // 'providingService'. The type is EntityStandbyStatus.
    Entstatestandby interface{}
}

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetFilter() yfilter.YFilter { return entstateentry.YFilter }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) SetFilter(yf yfilter.YFilter) { entstateentry.YFilter = yf }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "entStateLastChanged" { return "Entstatelastchanged" }
    if yname == "entStateAdmin" { return "Entstateadmin" }
    if yname == "entStateOper" { return "Entstateoper" }
    if yname == "entStateUsage" { return "Entstateusage" }
    if yname == "entStateAlarm" { return "Entstatealarm" }
    if yname == "entStateStandby" { return "Entstatestandby" }
    return ""
}

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetSegmentPath() string {
    return "entStateEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", entstateentry.Entphysicalindex) + "']"
}

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = entstateentry.Entphysicalindex
    leafs["entStateLastChanged"] = entstateentry.Entstatelastchanged
    leafs["entStateAdmin"] = entstateentry.Entstateadmin
    leafs["entStateOper"] = entstateentry.Entstateoper
    leafs["entStateUsage"] = entstateentry.Entstateusage
    leafs["entStateAlarm"] = entstateentry.Entstatealarm
    leafs["entStateStandby"] = entstateentry.Entstatestandby
    return leafs
}

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetBundleName() string { return "cisco_ios_xe" }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetYangName() string { return "entStateEntry" }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) SetParent(parent types.Entity) { entstateentry.parent = parent }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetParent() types.Entity { return entstateentry.parent }

func (entstateentry *ENTITYSTATEMIB_Entstatetable_Entstateentry) GetParentYangName() string { return "entStateTable" }

