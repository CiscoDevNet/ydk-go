// This MIB defines objects that describe dynamic templates.  A
// dynamic template is a set of configuration attributes that a
// system can dynamically apply to a target.
// 
// The target of a dynamic template is the entity configured by the
// system using the configuration attributes contained by a 
// template. At the time of this writing, an interface represents
// the only valid target of a dynamic template.  However, the
// architecture does not prevent other target types, and hence, the
// MIB definition generalizes the notion of a target to allow for
// this.
// 
// Generally, the system does not directly apply the attributes
// contained by a dynamic template to an associated
// target.  The system might generate a derived configuration from
// the set of dynamic templates associated with the target.  A
// derived configuration represents the union of the configuration
// attributes contained by each associated dynamic template.  In
// the case of a conflict (i.e., more than one dynamic template
// specifies the same configuration attribute), the system accepts
// the configuration attribute from the dynamic template with the
// highest precedence.
package cisco_dynamic_template_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_dynamic_template_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-DYNAMIC-TEMPLATE-MIB CISCO-DYNAMIC-TEMPLATE-MIB}", reflect.TypeOf(CISCODYNAMICTEMPLATEMIB{}))
    ydk.RegisterEntity("CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB", reflect.TypeOf(CISCODYNAMICTEMPLATEMIB{}))
}

// CISCODYNAMICTEMPLATEMIB
type CISCODYNAMICTEMPLATEMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table lists the dynamic templates maintained by the system, including
    // those that have been locally-configured on the system and those pushed to
    // the system by external policy servers.
    Cdttemplatetable CISCODYNAMICTEMPLATEMIB_Cdttemplatetable

    // This table contains a list of targets associated with one or more dynamic
    // templates.
    Cdttemplatetargettable CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable

    // This table contains a list of templates associated with each target.  This
    // table has an expansion dependent relationship on the
    // cdtTemplateTargetTable, containing zero or more rows for each target.
    Cdttemplateassociationtable CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable

    // This table contains a list of targets using each dynamic template.  This
    // table has an expansion dependent relationship on the cdtTemplateTable,
    // containing zero or more rows for each dynamic template.
    Cdttemplateusagetable CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable

    // This table contains attributes relating to all dynamic templates.  Observe
    // that the type of dynamic templates containing these attributes may change
    // with the addition of new dynamic template types.  This table has a
    // sparse-dependent relationship on the cdtTemplateTable, containing a row for
    // each dynamic template having a cdtTemplateType of one of the following
    // values:      'derived'     'ppp'     'ethernet'     'ipSubscriber'    
    // 'service'.
    Cdttemplatecommontable CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable

    // This table contains attributes relating to interface configuration.  This
    // table has a sparse-dependent relationship on the cdtTemplateTable,
    // containing a row for each dynamic template having a cdtTemplateType of one
    // of the following values:      'derived'     'ppp'     'ethernet'    
    // 'ipSubscriber'.
    Cdtiftemplatetable CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable

    // This table contains attributes relating to PPP connection configuration. 
    // This table has a sparse-dependent relationship on the cdtTemplateTable,
    // containing a row for each dynamic template having a cdtTemplateType of one
    // of the following values:      'derived'     'ppp'.
    Cdtppptemplatetable CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable

    // This table contains a prioritized list of named pools for each PPP
    // template.  A list corresponding to a PPP template specifies the pools the
    // system searches in an attempt to assign an IP address to the peer of a
    // target PPP connection. The system searches the pools in the order of their
    // priority.  This table has an expansion dependent relationship on the
    // cdtPppTemplateTable, containing zero or more rows for each PPP template.
    Cdtppppeeripaddrpooltable CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable

    // This table contains attributes relating to dynamic interfaces initiated on
    // Ethernet virtual interfaces (e.g., EoMPLS) or automatically created VLANs. 
    // This table has a sparse-dependent relationship on the cdtTemplateTable,
    // containing a row for each dynamic template having a cdtTemplateType of one
    // of the following values:      'derived'     'ethernet'.
    Cdtethernettemplatetable CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable

    // This table contains attributes relating to a service.  This table has a
    // sparse-dependent relationship on the cdtTemplateTable, containing a row for
    // each dynamic template having a cdtTemplateType of one of the following
    // values:      'derived'     'service'.
    Cdtsrvtemplatetable CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable
}

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetFilter() yfilter.YFilter { return cISCODYNAMICTEMPLATEMIB.YFilter }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) SetFilter(yf yfilter.YFilter) { cISCODYNAMICTEMPLATEMIB.YFilter = yf }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetGoName(yname string) string {
    if yname == "cdtTemplateTable" { return "Cdttemplatetable" }
    if yname == "cdtTemplateTargetTable" { return "Cdttemplatetargettable" }
    if yname == "cdtTemplateAssociationTable" { return "Cdttemplateassociationtable" }
    if yname == "cdtTemplateUsageTable" { return "Cdttemplateusagetable" }
    if yname == "cdtTemplateCommonTable" { return "Cdttemplatecommontable" }
    if yname == "cdtIfTemplateTable" { return "Cdtiftemplatetable" }
    if yname == "cdtPppTemplateTable" { return "Cdtppptemplatetable" }
    if yname == "cdtPppPeerIpAddrPoolTable" { return "Cdtppppeeripaddrpooltable" }
    if yname == "cdtEthernetTemplateTable" { return "Cdtethernettemplatetable" }
    if yname == "cdtSrvTemplateTable" { return "Cdtsrvtemplatetable" }
    return ""
}

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetSegmentPath() string {
    return "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB"
}

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtTemplateTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdttemplatetable
    }
    if childYangName == "cdtTemplateTargetTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdttemplatetargettable
    }
    if childYangName == "cdtTemplateAssociationTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdttemplateassociationtable
    }
    if childYangName == "cdtTemplateUsageTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdttemplateusagetable
    }
    if childYangName == "cdtTemplateCommonTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdttemplatecommontable
    }
    if childYangName == "cdtIfTemplateTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdtiftemplatetable
    }
    if childYangName == "cdtPppTemplateTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdtppptemplatetable
    }
    if childYangName == "cdtPppPeerIpAddrPoolTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdtppppeeripaddrpooltable
    }
    if childYangName == "cdtEthernetTemplateTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdtethernettemplatetable
    }
    if childYangName == "cdtSrvTemplateTable" {
        return &cISCODYNAMICTEMPLATEMIB.Cdtsrvtemplatetable
    }
    return nil
}

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cdtTemplateTable"] = &cISCODYNAMICTEMPLATEMIB.Cdttemplatetable
    children["cdtTemplateTargetTable"] = &cISCODYNAMICTEMPLATEMIB.Cdttemplatetargettable
    children["cdtTemplateAssociationTable"] = &cISCODYNAMICTEMPLATEMIB.Cdttemplateassociationtable
    children["cdtTemplateUsageTable"] = &cISCODYNAMICTEMPLATEMIB.Cdttemplateusagetable
    children["cdtTemplateCommonTable"] = &cISCODYNAMICTEMPLATEMIB.Cdttemplatecommontable
    children["cdtIfTemplateTable"] = &cISCODYNAMICTEMPLATEMIB.Cdtiftemplatetable
    children["cdtPppTemplateTable"] = &cISCODYNAMICTEMPLATEMIB.Cdtppptemplatetable
    children["cdtPppPeerIpAddrPoolTable"] = &cISCODYNAMICTEMPLATEMIB.Cdtppppeeripaddrpooltable
    children["cdtEthernetTemplateTable"] = &cISCODYNAMICTEMPLATEMIB.Cdtethernettemplatetable
    children["cdtSrvTemplateTable"] = &cISCODYNAMICTEMPLATEMIB.Cdtsrvtemplatetable
    return children
}

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) SetParent(parent types.Entity) { cISCODYNAMICTEMPLATEMIB.parent = parent }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetParent() types.Entity { return cISCODYNAMICTEMPLATEMIB.parent }

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplatetable
// This table lists the dynamic templates maintained by the
// system, including those that have been locally-configured on the
// system and those pushed to the system by external policy
// servers.
type CISCODYNAMICTEMPLATEMIB_Cdttemplatetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describes a dynamic template, which serves as a container for
    // configuration attributes.  The configuration attributes contained by a
    // dynamic template depends on its type.  The system automatically creates a
    // corresponding entry when a dynamic template has been created through
    // another management entity (e.g., the system's local console).  Likewise,
    // the system automatically destroys a corresponding entry when a dynamic
    // template has been destroyed through another management entity.  The system
    // automatically creates a corresponding entry when an external policy server
    // pushes a user profile or a service profile to the system.  The system
    // automatically creates a corresponding entry when it generates a derived
    // configuration. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry.
    Cdttemplateentry []CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry
}

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetFilter() yfilter.YFilter { return cdttemplatetable.YFilter }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) SetFilter(yf yfilter.YFilter) { cdttemplatetable.YFilter = yf }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetGoName(yname string) string {
    if yname == "cdtTemplateEntry" { return "Cdttemplateentry" }
    return ""
}

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetSegmentPath() string {
    return "cdtTemplateTable"
}

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtTemplateEntry" {
        for _, c := range cdttemplatetable.Cdttemplateentry {
            if cdttemplatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry{}
        cdttemplatetable.Cdttemplateentry = append(cdttemplatetable.Cdttemplateentry, child)
        return &cdttemplatetable.Cdttemplateentry[len(cdttemplatetable.Cdttemplateentry)-1]
    }
    return nil
}

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdttemplatetable.Cdttemplateentry {
        children[cdttemplatetable.Cdttemplateentry[i].GetSegmentPath()] = &cdttemplatetable.Cdttemplateentry[i]
    }
    return children
}

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetYangName() string { return "cdtTemplateTable" }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) SetParent(parent types.Entity) { cdttemplatetable.parent = parent }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetParent() types.Entity { return cdttemplatetable.parent }

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry
// An entry describes a dynamic template, which serves as a
// container for configuration attributes.  The configuration
// attributes contained by a dynamic template depends on its type.
// 
// The system automatically creates a corresponding entry when a
// dynamic template has been created through another management
// entity (e.g., the system's local console).  Likewise, the system
// automatically destroys a corresponding entry when a dynamic
// template has been destroyed through another management entity.
// 
// The system automatically creates a corresponding entry when an
// external policy server pushes a user profile or a service
// profile to the system.
// 
// The system automatically creates a corresponding entry when it
// generates a derived configuration.
type CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates a string-value that uniquely
    // identifies the dynamic template.  If the corresponding instance of
    // cdtTemplateSrc is not 'local', then the system automatically generates the
    // name identifying the dynamic template. The type is string with length:
    // 1..64.
    Cdttemplatename interface{}

    // This object specifies the status of the dynamic template.  The following
    // columns must be valid before activating a dynamic template:      -
    // cdtTemplateStorage     - cdtTemplateType  However, these objects specify a
    // default value.  Thus, it is possible to use create-and-go semantics without
    // setting any additional columns.  An implementation must allow the EMS/NMS
    // to modify any column when this column is 'active', including columns
    // defined in tables that have a one-to-one or sparse dependent relationship
    // on this table. The type is RowStatus.
    Cdttemplatestatus interface{}

    // This object specifies what happens to the dynamic template upon restart. 
    // If the corresponding instance of cdtTemplateSrc is not 'local', then this
    // column must be 'volatile'. The type is StorageType.
    Cdttemplatestorage interface{}

    // This object indicates the types of dynamic template. The type is
    // DynamicTemplateType.
    Cdttemplatetype interface{}

    // This object specifies the source of the dynamic template:  'other'     The
    // implementation of the MIB module does not recognize the     source of the
    // dynamic template.  'derived'     The system created the set of attributes
    // from one or     more dynamic templates.  'local'     The dynamic template
    // was locally configured through a     management entity, such as the local
    // console or a SNMP     entity.  'aaaUserProfile'     The dynamic template
    // originated from a user profile     pushed from an external policy server. 
    // 'aaaServiceProfile'     The dynamic template originated from a service
    // profile     pushed from an external policy server. The type is
    // Cdttemplatesrc.
    Cdttemplatesrc interface{}

    // This object specifies the number of targets using a dynamic template. The
    // type is interface{} with range: 0..4294967295.
    Cdttemplateusagecount interface{}
}

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetFilter() yfilter.YFilter { return cdttemplateentry.YFilter }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) SetFilter(yf yfilter.YFilter) { cdttemplateentry.YFilter = yf }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetGoName(yname string) string {
    if yname == "cdtTemplateName" { return "Cdttemplatename" }
    if yname == "cdtTemplateStatus" { return "Cdttemplatestatus" }
    if yname == "cdtTemplateStorage" { return "Cdttemplatestorage" }
    if yname == "cdtTemplateType" { return "Cdttemplatetype" }
    if yname == "cdtTemplateSrc" { return "Cdttemplatesrc" }
    if yname == "cdtTemplateUsageCount" { return "Cdttemplateusagecount" }
    return ""
}

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetSegmentPath() string {
    return "cdtTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdttemplateentry.Cdttemplatename) + "']"
}

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateName"] = cdttemplateentry.Cdttemplatename
    leafs["cdtTemplateStatus"] = cdttemplateentry.Cdttemplatestatus
    leafs["cdtTemplateStorage"] = cdttemplateentry.Cdttemplatestorage
    leafs["cdtTemplateType"] = cdttemplateentry.Cdttemplatetype
    leafs["cdtTemplateSrc"] = cdttemplateentry.Cdttemplatesrc
    leafs["cdtTemplateUsageCount"] = cdttemplateentry.Cdttemplateusagecount
    return leafs
}

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetYangName() string { return "cdtTemplateEntry" }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) SetParent(parent types.Entity) { cdttemplateentry.parent = parent }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetParent() types.Entity { return cdttemplateentry.parent }

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetParentYangName() string { return "cdtTemplateTable" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc represents     pushed from an external policy server.
type CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc string

const (
    CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc_other CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc = "other"

    CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc_derived CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc = "derived"

    CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc_local CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc = "local"

    CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc_aaaUserProfile CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc = "aaaUserProfile"

    CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc_aaaServiceProfile CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatesrc = "aaaServiceProfile"
)

// CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable
// This table contains a list of targets associated with
// one or more dynamic templates.
type CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry describes a target associated with one or more dynamic templates. 
    // The system automatically creates an entry when it associates a dynamic
    // template to a target.  Likewise, the system automatically destroys an entry
    // when a target no longer has any associated dynamic templates. The type is
    // slice of
    // CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry.
    Cdttemplatetargetentry []CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry
}

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetFilter() yfilter.YFilter { return cdttemplatetargettable.YFilter }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) SetFilter(yf yfilter.YFilter) { cdttemplatetargettable.YFilter = yf }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetGoName(yname string) string {
    if yname == "cdtTemplateTargetEntry" { return "Cdttemplatetargetentry" }
    return ""
}

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetSegmentPath() string {
    return "cdtTemplateTargetTable"
}

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtTemplateTargetEntry" {
        for _, c := range cdttemplatetargettable.Cdttemplatetargetentry {
            if cdttemplatetargettable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry{}
        cdttemplatetargettable.Cdttemplatetargetentry = append(cdttemplatetargettable.Cdttemplatetargetentry, child)
        return &cdttemplatetargettable.Cdttemplatetargetentry[len(cdttemplatetargettable.Cdttemplatetargetentry)-1]
    }
    return nil
}

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdttemplatetargettable.Cdttemplatetargetentry {
        children[cdttemplatetargettable.Cdttemplatetargetentry[i].GetSegmentPath()] = &cdttemplatetargettable.Cdttemplatetargetentry[i]
    }
    return children
}

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetYangName() string { return "cdtTemplateTargetTable" }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) SetParent(parent types.Entity) { cdttemplatetargettable.parent = parent }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetParent() types.Entity { return cdttemplatetargettable.parent }

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry
// An entry describes a target associated with one or more
// dynamic templates.
// 
// The system automatically creates an entry when it associates a
// dynamic template to a target.  Likewise, the system
// automatically destroys an entry when a target no longer has any
// associated dynamic templates.
type CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object indicates the type of target. The type
    // is DynamicTemplateTargetType.
    Cdttemplatetargettype interface{}

    // This attribute is a key. This object uniquely identifies the target within
    // the scope of its type. The type is string with length: 1..20.
    Cdttemplatetargetid interface{}

    // This object specifies the status of the dynamic template target.  The
    // following columns must be valid before activating a subscriber access
    // profile:      - cdtTemplateTargetStorage  However, these objects specify a
    // default value.  Thus, it is possible to use create-and-go semantics without
    // setting any additional columns.  An implementation must allow the EMS/NMS
    // to modify any column when this column is 'active', including columns
    // defined in tables that have a one-to-one or sparse dependent relationship
    // on this table. The type is RowStatus.
    Cdttemplatetargetstatus interface{}

    // This object specifies what happens to the dynamic template target upon
    // restart. The type is StorageType.
    Cdttemplatetargetstorage interface{}
}

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetFilter() yfilter.YFilter { return cdttemplatetargetentry.YFilter }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) SetFilter(yf yfilter.YFilter) { cdttemplatetargetentry.YFilter = yf }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetGoName(yname string) string {
    if yname == "cdtTemplateTargetType" { return "Cdttemplatetargettype" }
    if yname == "cdtTemplateTargetId" { return "Cdttemplatetargetid" }
    if yname == "cdtTemplateTargetStatus" { return "Cdttemplatetargetstatus" }
    if yname == "cdtTemplateTargetStorage" { return "Cdttemplatetargetstorage" }
    return ""
}

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetSegmentPath() string {
    return "cdtTemplateTargetEntry" + "[cdtTemplateTargetType='" + fmt.Sprintf("%v", cdttemplatetargetentry.Cdttemplatetargettype) + "']" + "[cdtTemplateTargetId='" + fmt.Sprintf("%v", cdttemplatetargetentry.Cdttemplatetargetid) + "']"
}

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateTargetType"] = cdttemplatetargetentry.Cdttemplatetargettype
    leafs["cdtTemplateTargetId"] = cdttemplatetargetentry.Cdttemplatetargetid
    leafs["cdtTemplateTargetStatus"] = cdttemplatetargetentry.Cdttemplatetargetstatus
    leafs["cdtTemplateTargetStorage"] = cdttemplatetargetentry.Cdttemplatetargetstorage
    return leafs
}

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetYangName() string { return "cdtTemplateTargetEntry" }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) SetParent(parent types.Entity) { cdttemplatetargetentry.parent = parent }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetParent() types.Entity { return cdttemplatetargetentry.parent }

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetParentYangName() string { return "cdtTemplateTargetTable" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable
// This table contains a list of templates associated with each
// target.
// 
// This table has an expansion dependent relationship on the
// cdtTemplateTargetTable, containing zero or more rows for each
// target.
type CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry indicates an association of a dynamic template with a target.  The
    // system automatically creates an entry when it associates a dynamic template
    // to a target.  The system also creates an entry when it derives the
    // configuration that it applies to a target.  The system automatically
    // destroys an entry under the following circumstances:  1)  The target
    // indicated by the entry no longer exists.  2)  The association between the
    // target and the dynamic template     no longer exists. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry.
    Cdttemplateassociationentry []CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry
}

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetFilter() yfilter.YFilter { return cdttemplateassociationtable.YFilter }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) SetFilter(yf yfilter.YFilter) { cdttemplateassociationtable.YFilter = yf }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetGoName(yname string) string {
    if yname == "cdtTemplateAssociationEntry" { return "Cdttemplateassociationentry" }
    return ""
}

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetSegmentPath() string {
    return "cdtTemplateAssociationTable"
}

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtTemplateAssociationEntry" {
        for _, c := range cdttemplateassociationtable.Cdttemplateassociationentry {
            if cdttemplateassociationtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry{}
        cdttemplateassociationtable.Cdttemplateassociationentry = append(cdttemplateassociationtable.Cdttemplateassociationentry, child)
        return &cdttemplateassociationtable.Cdttemplateassociationentry[len(cdttemplateassociationtable.Cdttemplateassociationentry)-1]
    }
    return nil
}

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdttemplateassociationtable.Cdttemplateassociationentry {
        children[cdttemplateassociationtable.Cdttemplateassociationentry[i].GetSegmentPath()] = &cdttemplateassociationtable.Cdttemplateassociationentry[i]
    }
    return children
}

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetYangName() string { return "cdtTemplateAssociationTable" }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) SetParent(parent types.Entity) { cdttemplateassociationtable.parent = parent }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetParent() types.Entity { return cdttemplateassociationtable.parent }

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry
// An entry indicates an association of a dynamic template with a
// target.
// 
// The system automatically creates an entry when it associates a
// dynamic template to a target.
// 
// The system also creates an entry when it derives the
// configuration that it applies to a target.
// 
// The system automatically destroys an entry under the following
// circumstances:
// 
// 1)  The target indicated by the entry no longer exists.
// 
// 2)  The association between the target and the dynamic template
//     no longer exists.
type CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is DynamicTemplateTargetType. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry_Cdttemplatetargettype
    Cdttemplatetargettype interface{}

    // This attribute is a key. The type is string with length: 1..20. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry_Cdttemplatetargetid
    Cdttemplatetargetid interface{}

    // This attribute is a key. This object indicates the name of the template
    // associated with the target. The type is string with length: 1..64.
    Cdttemplateassociationname interface{}

    // This object indicates the relative precedence of the associated dynamic
    // template.  Lower values have higher precedence than higher values. The type
    // is interface{} with range: 0..4294967295.
    Cdttemplateassociationprecedence interface{}
}

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetFilter() yfilter.YFilter { return cdttemplateassociationentry.YFilter }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) SetFilter(yf yfilter.YFilter) { cdttemplateassociationentry.YFilter = yf }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetGoName(yname string) string {
    if yname == "cdtTemplateTargetType" { return "Cdttemplatetargettype" }
    if yname == "cdtTemplateTargetId" { return "Cdttemplatetargetid" }
    if yname == "cdtTemplateAssociationName" { return "Cdttemplateassociationname" }
    if yname == "cdtTemplateAssociationPrecedence" { return "Cdttemplateassociationprecedence" }
    return ""
}

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetSegmentPath() string {
    return "cdtTemplateAssociationEntry" + "[cdtTemplateTargetType='" + fmt.Sprintf("%v", cdttemplateassociationentry.Cdttemplatetargettype) + "']" + "[cdtTemplateTargetId='" + fmt.Sprintf("%v", cdttemplateassociationentry.Cdttemplatetargetid) + "']" + "[cdtTemplateAssociationName='" + fmt.Sprintf("%v", cdttemplateassociationentry.Cdttemplateassociationname) + "']"
}

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateTargetType"] = cdttemplateassociationentry.Cdttemplatetargettype
    leafs["cdtTemplateTargetId"] = cdttemplateassociationentry.Cdttemplatetargetid
    leafs["cdtTemplateAssociationName"] = cdttemplateassociationentry.Cdttemplateassociationname
    leafs["cdtTemplateAssociationPrecedence"] = cdttemplateassociationentry.Cdttemplateassociationprecedence
    return leafs
}

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetYangName() string { return "cdtTemplateAssociationEntry" }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) SetParent(parent types.Entity) { cdttemplateassociationentry.parent = parent }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetParent() types.Entity { return cdttemplateassociationentry.parent }

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetParentYangName() string { return "cdtTemplateAssociationTable" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable
// This table contains a list of targets using each dynamic
// template.
// 
// This table has an expansion dependent relationship on the
// cdtTemplateTable, containing zero or more rows for each
// dynamic template.
type CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry indicates a target using the dynamic template.  The system
    // automatically creates an entry when it associates a dynamic template to a
    // target.  The system also creates an entry when it derives the configuration
    // that it applies to a target.  The system automatically destroys an entry
    // under the following circumstances:  1)  The target indicated by the entry
    // no longer exists.  2)  The association between the target and the dynamic
    // template     no longer exists. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry.
    Cdttemplateusageentry []CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry
}

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetFilter() yfilter.YFilter { return cdttemplateusagetable.YFilter }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) SetFilter(yf yfilter.YFilter) { cdttemplateusagetable.YFilter = yf }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetGoName(yname string) string {
    if yname == "cdtTemplateUsageEntry" { return "Cdttemplateusageentry" }
    return ""
}

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetSegmentPath() string {
    return "cdtTemplateUsageTable"
}

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtTemplateUsageEntry" {
        for _, c := range cdttemplateusagetable.Cdttemplateusageentry {
            if cdttemplateusagetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry{}
        cdttemplateusagetable.Cdttemplateusageentry = append(cdttemplateusagetable.Cdttemplateusageentry, child)
        return &cdttemplateusagetable.Cdttemplateusageentry[len(cdttemplateusagetable.Cdttemplateusageentry)-1]
    }
    return nil
}

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdttemplateusagetable.Cdttemplateusageentry {
        children[cdttemplateusagetable.Cdttemplateusageentry[i].GetSegmentPath()] = &cdttemplateusagetable.Cdttemplateusageentry[i]
    }
    return children
}

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetYangName() string { return "cdtTemplateUsageTable" }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) SetParent(parent types.Entity) { cdttemplateusagetable.parent = parent }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetParent() types.Entity { return cdttemplateusagetable.parent }

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry
// An entry indicates a target using the dynamic template.
// 
// The system automatically creates an entry when it associates a
// dynamic template to a target.
// 
// The system also creates an entry when it derives the
// configuration that it applies to a target.
// 
// The system automatically destroys an entry under the following
// circumstances:
// 
// 1)  The target indicated by the entry no longer exists.
// 
// 2)  The association between the target and the dynamic template
//     no longer exists.
type CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatename
    Cdttemplatename interface{}

    // This attribute is a key. This object indicates the type of target using the
    // dynamic template. The type is DynamicTemplateTargetType.
    Cdttemplateusagetargettype interface{}

    // This attribute is a key. This object indicates the name of the target using
    // the dynamic template. The type is string with length: 1..20.
    Cdttemplateusagetargetid interface{}
}

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetFilter() yfilter.YFilter { return cdttemplateusageentry.YFilter }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) SetFilter(yf yfilter.YFilter) { cdttemplateusageentry.YFilter = yf }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetGoName(yname string) string {
    if yname == "cdtTemplateName" { return "Cdttemplatename" }
    if yname == "cdtTemplateUsageTargetType" { return "Cdttemplateusagetargettype" }
    if yname == "cdtTemplateUsageTargetId" { return "Cdttemplateusagetargetid" }
    return ""
}

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetSegmentPath() string {
    return "cdtTemplateUsageEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdttemplateusageentry.Cdttemplatename) + "']" + "[cdtTemplateUsageTargetType='" + fmt.Sprintf("%v", cdttemplateusageentry.Cdttemplateusagetargettype) + "']" + "[cdtTemplateUsageTargetId='" + fmt.Sprintf("%v", cdttemplateusageentry.Cdttemplateusagetargetid) + "']"
}

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateName"] = cdttemplateusageentry.Cdttemplatename
    leafs["cdtTemplateUsageTargetType"] = cdttemplateusageentry.Cdttemplateusagetargettype
    leafs["cdtTemplateUsageTargetId"] = cdttemplateusageentry.Cdttemplateusagetargetid
    return leafs
}

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetYangName() string { return "cdtTemplateUsageEntry" }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) SetParent(parent types.Entity) { cdttemplateusageentry.parent = parent }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetParent() types.Entity { return cdttemplateusageentry.parent }

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetParentYangName() string { return "cdtTemplateUsageTable" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable
// This table contains attributes relating to all dynamic
// templates.  Observe that the type of dynamic templates
// containing these attributes may change with the addition of new
// dynamic template types.
// 
// This table has a sparse-dependent relationship on the
// cdtTemplateTable, containing a row for each dynamic template
// having a cdtTemplateType of one of the following values:
// 
//     'derived'
//     'ppp'
//     'ethernet'
//     'ipSubscriber'
//     'service'
type CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing attributes relating to any target.  The system
    // automatically creates an entry when the system or the EMS/NMS creates a row
    // in the cdtTemplateTable with a cdtTemplateType of on the following values: 
    // 'derived'     'ppp'     'ethernet'     'ipSubscriber'     'service' 
    // Likewise, the system automatically destroys an entry when the system or the
    // EMS/NMS destroys the corresponding row in the cdtTemplateTable. The type is
    // slice of
    // CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry.
    Cdttemplatecommonentry []CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry
}

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetFilter() yfilter.YFilter { return cdttemplatecommontable.YFilter }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) SetFilter(yf yfilter.YFilter) { cdttemplatecommontable.YFilter = yf }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetGoName(yname string) string {
    if yname == "cdtTemplateCommonEntry" { return "Cdttemplatecommonentry" }
    return ""
}

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetSegmentPath() string {
    return "cdtTemplateCommonTable"
}

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtTemplateCommonEntry" {
        for _, c := range cdttemplatecommontable.Cdttemplatecommonentry {
            if cdttemplatecommontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry{}
        cdttemplatecommontable.Cdttemplatecommonentry = append(cdttemplatecommontable.Cdttemplatecommonentry, child)
        return &cdttemplatecommontable.Cdttemplatecommonentry[len(cdttemplatecommontable.Cdttemplatecommonentry)-1]
    }
    return nil
}

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdttemplatecommontable.Cdttemplatecommonentry {
        children[cdttemplatecommontable.Cdttemplatecommonentry[i].GetSegmentPath()] = &cdttemplatecommontable.Cdttemplatecommonentry[i]
    }
    return children
}

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetYangName() string { return "cdtTemplateCommonTable" }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) SetParent(parent types.Entity) { cdttemplatecommontable.parent = parent }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetParent() types.Entity { return cdttemplatecommontable.parent }

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry
// An entry containing attributes relating to any target.
// 
// The system automatically creates an entry when the system or the
// EMS/NMS creates a row in the cdtTemplateTable with a
// cdtTemplateType of on the following values:
// 
//     'derived'
//     'ppp'
//     'ethernet'
//     'ipSubscriber'
//     'service'
// 
// Likewise, the system automatically destroys an entry when the
// system or the EMS/NMS destroys the corresponding row in the
// cdtTemplateTable.
type CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatename
    Cdttemplatename interface{}

    // This object specifies which attributes in the dynamic template have been
    // configured to valid values.  Each bit in this bit string corresponds to a
    // column in this table.  If the bit is '0', then the value of the
    // corresponding column is not valid.  If the bit is '1', then the value of
    // the corresponding column has been configured to a valid value.  The
    // following list specifies the mappings between bits and the columns:     
    // 'descr'             => cdtCommonDescr     'keepaliveInt'      =>
    // cdtCommonKeepaliveInt     'keepaliveRetries'  => cdtCommonKeepaliveRetries 
    // 'vrf'               => cdtCommonVrf     'addrPool'          =>
    // cdtCommonAddrPool     'ipv4AccessGroup'   => cdtCommonIpv4AccessGroup    
    // 'ipv4Unreachables'  => cdtCommonIpv4Unreachables     'ipv6AccessGroup'   =>
    // cdtCommonIpv6AccessGroup     'ipv6Unreachables'  =>
    // cdtCommonIpv6Unreachables     'srvSubControl'     => cdtCommonSrvSubControl
    // 'srvRedirect'       => cdtCommonSrvRedirect     'srvAcct'           =>
    // cdtCommonSrvAcct     'srvQos'            => cdtCommonSrvQos    
    // 'srvNetflow'        => cdtCommonSrvNetflow. The type is map[string]bool.
    Cdtcommonvalid interface{}

    // This object specifies a human-readable description for the dynamic
    // template.  This column is valid only if the 'descr' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is string with
    // length: 0..255.
    Cdtcommondescr interface{}

    // This object specifies the interval that the system sends keepalive messages
    // to a target.  This column is valid only if the 'keepaliveInterval' bit of
    // the corresponding instance of cdtCommonValid is '1'. The type is
    // interface{} with range: 1..4294967295. Units are seconds.
    Cdtcommonkeepaliveint interface{}

    // This object specifies the number of times the system will resend a
    // keepalive message without a response before it transitions a target to an
    // operationally down state.  This column is valid only if the
    // 'keepaliveRetries' bit of the corresponding instance of cdtCommonValid is
    // '1'. The type is interface{} with range: 1..255. Units are retries.
    Cdtcommonkeepaliveretries interface{}

    // This object specifies the name of the VRF with which a target has an
    // association.  This column is valid only if the 'vrf' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is string with
    // length: 0..32.
    Cdtcommonvrf interface{}

    // This object specifies the name of the IP address pool the system will use
    // to assign an IP address to a peer of a target.  This column is valid only
    // if the 'addrPool' bit of the corresponding instance of cdtCommonValid is
    // '1'. The type is string with length: 0..255.
    Cdtcommonaddrpool interface{}

    // This object specifies the name (or number) of the IPv4 ACL applied to a
    // target.  This column is valid only if the 'ipv4AccessGroup' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is string with
    // length: 0..255.
    Cdtcommonipv4Accessgroup interface{}

    // This object specifies whether a target generates ICMPv4 unreachable
    // messages.  This column is valid only if the 'ipv4Unreachables' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is bool.
    Cdtcommonipv4Unreachables interface{}

    // This object specifies the name (or number) of the IPv4 ACL applied to a
    // target.  This column is valid only if the 'ipv6AccessGroup' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is string with
    // length: 0..255.
    Cdtcommonipv6Accessgroup interface{}

    // This object specifies whether a target generates ICMPv6 unreachable
    // messages.  This column is valid only if the 'ipv6Unreachables' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is bool.
    Cdtcommonipv6Unreachables interface{}

    // This object specifies the name of the subscriber control policy applied to
    // a target.  The system should assume that the cbpPolicyMapType (defined by
    // the CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtControlSubscriber
    // (defined by the CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the
    // 'srvSubControl' bit of the corresponding instance of cdtCommonValid is '1'.
    // The type is string.
    Cdtcommonsrvsubcontrol interface{}

    // This object specifies the name of the traffic redirect policy applied to a
    // target.  The system should assume that the cbpPolicyMapType (defined by the
    // CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtTrafficRedirect (defined by
    // the CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the
    // 'srvRedirect' bit of the corresponding instance of cdtCommonValid is '1'.
    // The type is string.
    Cdtcommonsrvredirect interface{}

    // This object specifies the name of the traffic accounting policy applied to
    // a target.  The system should assume that the cbpPolicyMapType (defined by
    // the CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtTrafficAccounting
    // (defined by the CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the
    // 'srvAcct' bit of the corresponding instance of cdtCommonValid is '1'. The
    // type is string.
    Cdtcommonsrvacct interface{}

    // This object specifies the name of the traffic QoS policy applied to a
    // target.  The system should assume that the cbpPolicyMapType (defined by the
    // CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtQos (defined by the
    // CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the 'srvQos' bit of
    // the corresponding instance of cdtCommonValid is '1'. The type is string.
    Cdtcommonsrvqos interface{}

    // This object specifies the name of the NetFlow policy applied to a target. 
    // The system should assume that the cbpPolicyMapType (defined by the
    // CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtNetflow (defined by the
    // CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the 'srvNetflow' bit
    // of the corresponding instance of cdtCommonValid is '1'. The type is string.
    Cdtcommonsrvnetflow interface{}
}

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetFilter() yfilter.YFilter { return cdttemplatecommonentry.YFilter }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) SetFilter(yf yfilter.YFilter) { cdttemplatecommonentry.YFilter = yf }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetGoName(yname string) string {
    if yname == "cdtTemplateName" { return "Cdttemplatename" }
    if yname == "cdtCommonValid" { return "Cdtcommonvalid" }
    if yname == "cdtCommonDescr" { return "Cdtcommondescr" }
    if yname == "cdtCommonKeepaliveInt" { return "Cdtcommonkeepaliveint" }
    if yname == "cdtCommonKeepaliveRetries" { return "Cdtcommonkeepaliveretries" }
    if yname == "cdtCommonVrf" { return "Cdtcommonvrf" }
    if yname == "cdtCommonAddrPool" { return "Cdtcommonaddrpool" }
    if yname == "cdtCommonIpv4AccessGroup" { return "Cdtcommonipv4Accessgroup" }
    if yname == "cdtCommonIpv4Unreachables" { return "Cdtcommonipv4Unreachables" }
    if yname == "cdtCommonIpv6AccessGroup" { return "Cdtcommonipv6Accessgroup" }
    if yname == "cdtCommonIpv6Unreachables" { return "Cdtcommonipv6Unreachables" }
    if yname == "cdtCommonSrvSubControl" { return "Cdtcommonsrvsubcontrol" }
    if yname == "cdtCommonSrvRedirect" { return "Cdtcommonsrvredirect" }
    if yname == "cdtCommonSrvAcct" { return "Cdtcommonsrvacct" }
    if yname == "cdtCommonSrvQos" { return "Cdtcommonsrvqos" }
    if yname == "cdtCommonSrvNetflow" { return "Cdtcommonsrvnetflow" }
    return ""
}

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetSegmentPath() string {
    return "cdtTemplateCommonEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdttemplatecommonentry.Cdttemplatename) + "']"
}

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateName"] = cdttemplatecommonentry.Cdttemplatename
    leafs["cdtCommonValid"] = cdttemplatecommonentry.Cdtcommonvalid
    leafs["cdtCommonDescr"] = cdttemplatecommonentry.Cdtcommondescr
    leafs["cdtCommonKeepaliveInt"] = cdttemplatecommonentry.Cdtcommonkeepaliveint
    leafs["cdtCommonKeepaliveRetries"] = cdttemplatecommonentry.Cdtcommonkeepaliveretries
    leafs["cdtCommonVrf"] = cdttemplatecommonentry.Cdtcommonvrf
    leafs["cdtCommonAddrPool"] = cdttemplatecommonentry.Cdtcommonaddrpool
    leafs["cdtCommonIpv4AccessGroup"] = cdttemplatecommonentry.Cdtcommonipv4Accessgroup
    leafs["cdtCommonIpv4Unreachables"] = cdttemplatecommonentry.Cdtcommonipv4Unreachables
    leafs["cdtCommonIpv6AccessGroup"] = cdttemplatecommonentry.Cdtcommonipv6Accessgroup
    leafs["cdtCommonIpv6Unreachables"] = cdttemplatecommonentry.Cdtcommonipv6Unreachables
    leafs["cdtCommonSrvSubControl"] = cdttemplatecommonentry.Cdtcommonsrvsubcontrol
    leafs["cdtCommonSrvRedirect"] = cdttemplatecommonentry.Cdtcommonsrvredirect
    leafs["cdtCommonSrvAcct"] = cdttemplatecommonentry.Cdtcommonsrvacct
    leafs["cdtCommonSrvQos"] = cdttemplatecommonentry.Cdtcommonsrvqos
    leafs["cdtCommonSrvNetflow"] = cdttemplatecommonentry.Cdtcommonsrvnetflow
    return leafs
}

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetYangName() string { return "cdtTemplateCommonEntry" }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) SetParent(parent types.Entity) { cdttemplatecommonentry.parent = parent }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetParent() types.Entity { return cdttemplatecommonentry.parent }

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetParentYangName() string { return "cdtTemplateCommonTable" }

// CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable
// This table contains attributes relating to interface
// configuration.
// 
// This table has a sparse-dependent relationship on the
// cdtTemplateTable, containing a row for each dynamic template
// having a cdtTemplateType of one of the following values:
// 
//     'derived'
//     'ppp'
//     'ethernet'
//     'ipSubscriber'
type CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing attributes relating to interface configuration.  The
    // system automatically creates an entry when the system or the EMS/NMS
    // creates a row in the cdtTemplateTable with a cdtTemplateType of one of the
    // following values:      'derived'     'ppp'     'ethernet'    
    // 'ipSubscriber'  Likewise, the system automatically destroys an entry when
    // the system or the EMS/NMS destroys the corresponding row in the
    // cdtTemplateTable. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry.
    Cdtiftemplateentry []CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry
}

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetFilter() yfilter.YFilter { return cdtiftemplatetable.YFilter }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) SetFilter(yf yfilter.YFilter) { cdtiftemplatetable.YFilter = yf }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetGoName(yname string) string {
    if yname == "cdtIfTemplateEntry" { return "Cdtiftemplateentry" }
    return ""
}

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetSegmentPath() string {
    return "cdtIfTemplateTable"
}

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtIfTemplateEntry" {
        for _, c := range cdtiftemplatetable.Cdtiftemplateentry {
            if cdtiftemplatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry{}
        cdtiftemplatetable.Cdtiftemplateentry = append(cdtiftemplatetable.Cdtiftemplateentry, child)
        return &cdtiftemplatetable.Cdtiftemplateentry[len(cdtiftemplatetable.Cdtiftemplateentry)-1]
    }
    return nil
}

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdtiftemplatetable.Cdtiftemplateentry {
        children[cdtiftemplatetable.Cdtiftemplateentry[i].GetSegmentPath()] = &cdtiftemplatetable.Cdtiftemplateentry[i]
    }
    return children
}

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetBundleName() string { return "cisco_ios_xe" }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetYangName() string { return "cdtIfTemplateTable" }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) SetParent(parent types.Entity) { cdtiftemplatetable.parent = parent }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetParent() types.Entity { return cdtiftemplatetable.parent }

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry
// An entry containing attributes relating to interface
// configuration.
// 
// The system automatically creates an entry when the system or the
// EMS/NMS creates a row in the cdtTemplateTable with a
// cdtTemplateType of one of the following values:
// 
//     'derived'
//     'ppp'
//     'ethernet'
//     'ipSubscriber'
// 
// Likewise, the system automatically destroys an entry when the
// system or the EMS/NMS destroys the corresponding row in the
// cdtTemplateTable.
type CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatename
    Cdttemplatename interface{}

    // This object specifies which attributes in the dynamic template have been
    // configured to valid values.  Each bit in this bit string corresponds to a
    // column in this table.  If the bit is '0', then the value of the
    // corresponding column is not valid.  If the bit is '1', then the value of
    // the corresponding column has been configured to a valid value.  The
    // following list specifies the mappings between bits and the columns:     
    // 'mtu'                     => cdtIfMtu     'cdpEnable'               =>
    // cdtIfCdpEnable     'flowMonitor'             => cdtIfFlowMonitor    
    // 'ipv4Unnumbered'          => cdtIfIpv4Unnumbered     'ipv4SubEnable'       
    // => cdtIfIpv4SubEnable     'ipv4Mtu'                 => cdtIfIpv4Mtu    
    // 'ipv4TcpMssAdjust'        => cdtIfIpv4TcpMssAdjust     'ipv4VerifyUniRpf'  
    // => cdtIfIpv4VerifyUniRpf     'ipv4VerifyUniRpfAcl'     =>
    // cdtIfIpv4VerifyUniRpfAcl     'ipv4VerifyUniRpfOpts'    =>
    // cdtIfIpv4VerifyUniRpfOpts     'ipv6Enable'              => cdtIfIpv6Enable 
    // 'ipv6SubEnable'           => cdtIfIpv6SubEnable     'ipv6TcpMssAdjust'     
    // => cdtIfIpv6TcpMssAdjust     'ipv6VerifyUniRpf'        =>
    // cdtIfIpv6VerifyUniRpf     'ipv6VerifyUniRpfAcl'     =>
    // cdtIfIpv6VerifyUniRpfAcl     'ipv6VerifyUniRpfOpts'    =>
    // cdtIfIpv6VerifyUniRpfOpts     'ipv6NdPrefix'            =>
    // cdtIfIpv6NdPrefix,                                  cdtIfIpv6NdPrefixLength
    // 'ipv6NdValidLife'         => cdtIfIpv6NdValidLife     'ipv6NdPreferredLife'
    // => cdtIfIpv6NdPreferredLife     'ipv6NdOpts'              =>
    // cdtIfIpv6NdOpts     'ipv6NdDadAttempts'       => cdtIfIpv6NdDadAttempts    
    // 'ipv6NdNsInterval'        => cdtIfIpv6NdNsInterval     'ipv6NdReacableTime'
    // => cdtIfIpv6NdReacableTime     'ipv6NdRaIntervalMax'     =>
    // cdtIfIpv6NdRaIntervalUnits,                                 
    // cdtIfIpv6NdRaIntervalMax     'ipv6NdRaIntervalMin'     =>
    // cdtIfIpv6NdRaIntervalMin     'ipv6NdRaLife'            => cdtIfIpv6NdRaLife
    // 'ipv6NdRouterPreference'' => cdtIfIpv6NdRouterPreference. The type is
    // map[string]bool.
    Cdtifvalid interface{}

    // This object specifies the Maximum Transfer Unit (MTU) size for all packets
    // sent on the target interface.  The value '0' cannot be written to an
    // instance of this object. However, it serves as a convenient value when the
    // column is not valid.  This column is valid only if the 'mtu' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 0..None | 64..65535. Units are octets.
    Cdtifmtu interface{}

    // This object specifies whether the target interface participates in the
    // Cisco Discovery Protocol (CDP).  This column is valid only if the
    // 'cdpEnable' bit of the corresponding instance of cdtIfValid is '1'. The
    // type is bool.
    Cdtifcdpenable interface{}

    // This object specifies the name of the flow monitor associated with the
    // target interface.  This column is valid only if the 'flowMonitor' bit of
    // the corresponding instance of cdtIfValid is '1'. The type is string with
    // length: 0..255.
    Cdtifflowmonitor interface{}

    // This object specifies the interface of the source address that the target
    // interface uses when originating IPv4 packets.  The value '0' cannot be
    // written to an instance of this object. However, it serves as a convenient
    // value when the column is not valid (e.g., immediately following the
    // creation of an instance of the object).  This column is valid only if the
    // 'ipv4Unnumbered' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is interface{} with range: 0..2147483647.
    Cdtifipv4Unnumbered interface{}

    // This object specifies whether the target interface allows IPv4 subscriber
    // sessions.  This column is valid only if the 'ipv4SubEnable' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is bool.
    Cdtifipv4Subenable interface{}

    // This object specifies the Maximum Transfer Unit (MTU) size for IPv4 packets
    // sent on the target interface.  The value '0' cannot be written to an
    // instance of this object. However, it serves as a convenient value when the
    // column is not valid.  This column is valid only if the 'ipv4Mtu' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 0..None | 128..65535. Units are octets.
    Cdtifipv4Mtu interface{}

    // This object specifies the adjustment to the Maximum Segment Size (MSS) of
    // TCP SYN packets received by the target interface contained in IPv4
    // datagrams.  The value '0' cannot be written to an instance of this object.
    // However, it serves as a convenient value when the column is not valid. 
    // This column is valid only if the 'ipv4TcpMssAdjust' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 0..None | 500..1460. Units are octets.
    Cdtifipv4Tcpmssadjust interface{}

    // This object specifies whether the type of unicast RPF the system performs
    // on IPv4 packets received by the target interface.  This column is valid
    // only if the 'ipv4VerifyUniRpf' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is UnicastRpfType.
    Cdtifipv4Verifyunirpf interface{}

    // This object specifies the name (or number) of the IPv4 ACL used to
    // determine whether the system should permit/deny packets received by the
    // target interface that fail unicast RPF verification.  This column is valid
    // only if the 'ipv4VerifyUniRpfAcl' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is string with length: 0..255.
    Cdtifipv4Verifyunirpfacl interface{}

    // This object specifies the options that affect how the system performs
    // unicast RPF on IPv4 packets received by the target interface.  This column
    // is valid only if the 'ipv4VerifyUniRpfOpts' bit of the corresponding
    // instance of cdtIfValid is '1'. The type is map[string]bool.
    Cdtifipv4Verifyunirpfopts interface{}

    // This object specifies whether the system processes IPv6 packets received by
    // the target interface.  This column is valid only if the 'ipv6Enable' bit of
    // the corresponding instance of cdtIfValid is '1'. The type is bool.
    Cdtifipv6Enable interface{}

    // This object specifies whether the target interface allows IPv6 subscriber
    // sessions.  This column is valid only if the 'ipv6SubEnable' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is bool.
    Cdtifipv6Subenable interface{}

    // This object specifies the adjustment to the Maximum Segment Size (MSS) of
    // TCP SYN packets received by the target interface contained in IPv6
    // datagrams.  The value '0' cannot be written to an instance of this object.
    // However, it serves as a convenient value when the column is not valid. 
    // This column is valid only if the 'ipv6TcpMssAdjust' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 0..None | 500..1460. Units are octets.
    Cdtifipv6Tcpmssadjust interface{}

    // This object specifies whether the type of unicast RPF the system performs
    // on IPv6 packets received by the target interface.  This column is valid
    // only if the 'ipv6VerifyUniRpf' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is UnicastRpfType.
    Cdtifipv6Verifyunirpf interface{}

    // This object specifies the name (or number) of the IPv6 ACL used to
    // determine whether the system should permit/deny packets received by the
    // target interface that fail unicast RPF verification.  This column is valid
    // only if the 'ipv6VerifyUniRpfAcl' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is string with length: 0..255.
    Cdtifipv6Verifyunirpfacl interface{}

    // This object specifies the options that affect how the system performs
    // unicast RPF on IPv6 packets received by the target interface.  This column
    // is valid only if the 'ipv6VerifyUniRpfOpts' bit of the corresponding
    // instance of cdtIfValid is '1'. The type is map[string]bool.
    Cdtifipv6Verifyunirpfopts interface{}

    // This object specifies the IPv6 network number included in IPv6 router
    // advertisements sent on the target interface.  This column is valid only if
    // the 'ipv6NdPrefix' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is string.
    Cdtifipv6Ndprefix interface{}

    // This object specifies the length of the IPv6 prefix specified by the
    // corresponding instance of cdtIpv6NdPrefix.  This column is valid only if
    // the 'ipv6NdPrefix' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is interface{} with range: 0..2040.
    Cdtifipv6Ndprefixlength interface{}

    // This object specifies the interval that the system advertises the IPv6
    // prefix (i.e., the corresponding instance of cdtIfIpv6NdPrefix) as 'valid'
    // for IPv6 router advertisements sent on the target interface.  This column
    // is valid only if the 'ipv6NdValidLife' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is interface{} with range: 1..4294967295. Units
    // are seconds.
    Cdtifipv6Ndvalidlife interface{}

    // This object specifies the interval that the system advertises the IPv6
    // prefix (i.e., the corresponding instance of cdtIfIpv6NdPrefix) as
    // 'preferred' for IPv6 router advertisements sent on the target interface. 
    // This column is valid only if the 'ipv6NdPreferredLife' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 1..4294967295. Units are seconds.
    Cdtifipv6Ndpreferredlife interface{}

    // This object specifies options that affect advertisements sent on the target
    // interface:      'advertise'         This option specifies that the system
    // should advertise         the IPv6 prefix (i.e., the corresponding instance
    // of         cdtIfIpv6NdPrefix).      'onlink'         This option specifies
    // that the IPv6 prefix has been         assigned to a link.  If set to '0',
    // the system         advertises the IPv6 prefix as 'offlink'.      'router'  
    // This option indicates that the router will send the full         router
    // address and not set the 'R' bit in prefix         advertisements.     
    // 'autoConfig'         This option indicates to hosts on the local link that 
    // the specified prefix supports IPv6 auto-configuration.     
    // 'advertisementInterval'         This option specifies the advertisement
    // interval option         in router advertisements sent on the target
    // interface.      'managedConfigFlag'         This option causes the system
    // to set the 'managed         address configuration flag' in router
    // advertisements         sent on the target interface.      'otherConfigFlag'
    // This option causes the system to set the 'other stateful        
    // configuration' flag in router advertisements sent on the         target
    // interface.      'frameIpv6Prefix'         This option causes the system to
    // add the prefix in a         received RADIUS framed IPv6 prefix attribute to
    // the         target interface's neighbor discovery prefix queue and        
    // includes it in router advertisements sent on the target         interface. 
    // 'raSupress'         This option suppresses the transmission of router      
    // advertisements on the target interface.  This column is valid only if the
    // 'ipv6NdOpts' bit of the corresponding instance of cdtIfValid is '1'. The
    // type is map[string]bool.
    Cdtifipv6Ndopts interface{}

    // This object specifies the number of consecutive neighbor solitication
    // messages the system sends on the target interface while performing
    // duplicate address detection on unicast IPv6 addresses on the target
    // interface.  The value '0' disables duplicate address detection on the
    // target interface.  This column is valid only if the 'ipv6NdDadAttempts' bit
    // of the corresponding instance of cdtIfValid is '1'. The type is interface{}
    // with range: 0..600.
    Cdtifipv6Nddadattempts interface{}

    // This object specifies the interval between neighbor solicitation
    // retransmissions on the target interface.  This column is valid only if the
    // 'ipv6NdNsIntervals' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is interface{} with range: 1000..3600000. Units are milliseconds.
    Cdtifipv6Ndnsinterval interface{}

    // This object specifies the amount of time the system considers a neighbor of
    // the target interface reachable after a reachability confirmation event has
    // occurred.  The value '0' disables neighbor reachability detection on the
    // target interface.  This column is valid only if the 'ipv6NdReachable' bit
    // of the corresponding instance of cdtIfValid is '1'. The type is interface{}
    // with range: 0..4294967295. Units are milliseconds.
    Cdtifipv6Ndreachabletime interface{}

    // This object specifies the units of time for the corresponding instances of
    // cdtIfIpv6NdRaIntervalMin and cdtIfIpv6NdRaIntervalMax.  This column is
    // valid only if the 'ipv6NdRaInterval' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is Cdtifipv6Ndraintervalunits.
    Cdtifipv6Ndraintervalunits interface{}

    // This object specifies the maximum interval between IPv6 router
    // advertisements sent on the target interface.  The value '0' cannot be
    // written to an instance of this object. However, it serves as a convenient
    // value when the column is not valid.  This column is valid only if the
    // 'ipv6NdRaInterval' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is interface{} with range: 0..4294967295.
    Cdtifipv6Ndraintervalmax interface{}

    // This object specifies the minimum interval between IPv6 router
    // advertisements sent on the target interface.  The value of this column has
    // the following restrictions:  1)  This value cannot be less than 75% of the
    // value specified     for cdtIfIpv6NdRaIntervalMax.  2)  If the corresponding
    // instance of cdtIfIpv6NdRaIntervalUnits     is 'seconds', then this value
    // cannot be less than '3'.  3)  If the corresponding instance of
    // cdtIfIpv6NdRaIntervalUnits     is 'milliseconds', then this value cannot be
    // less than '30'.  If the target interface template does not specify this
    // value, then the system automatically assumes a minimum interval that is 75%
    // of the corresponding instance of cdtIfIpv6NdRaIntervalMax.  The value '0'
    // cannot be written to an instance of this object. However, it serves as a
    // convenient value when the column is not valid.  This column is valid only
    // if the 'ipv6NdRaInterval' bit of the corresponding instance of cdtIfValid
    // is '1'. The type is interface{} with range: 0..4294967295.
    Cdtifipv6Ndraintervalmin interface{}

    // This object specifies the router lifetime value in IPv6 router
    // advertisements sent on the target interface.  The value '0' specifies that
    // neighbors should not consider the router as a default router. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Cdtifipv6Ndralife interface{}

    // This object specifies the Default Router Preference (DRP) for the router on
    // the target interface.  This column is valid only if the
    // 'ipv6NdRouterPreference' bit of the corresponding instance of cdtIfValid is
    // '1'. The type is Cdtifipv6Ndrouterpreference.
    Cdtifipv6Ndrouterpreference interface{}
}

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetFilter() yfilter.YFilter { return cdtiftemplateentry.YFilter }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) SetFilter(yf yfilter.YFilter) { cdtiftemplateentry.YFilter = yf }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetGoName(yname string) string {
    if yname == "cdtTemplateName" { return "Cdttemplatename" }
    if yname == "cdtIfValid" { return "Cdtifvalid" }
    if yname == "cdtIfMtu" { return "Cdtifmtu" }
    if yname == "cdtIfCdpEnable" { return "Cdtifcdpenable" }
    if yname == "cdtIfFlowMonitor" { return "Cdtifflowmonitor" }
    if yname == "cdtIfIpv4Unnumbered" { return "Cdtifipv4Unnumbered" }
    if yname == "cdtIfIpv4SubEnable" { return "Cdtifipv4Subenable" }
    if yname == "cdtIfIpv4Mtu" { return "Cdtifipv4Mtu" }
    if yname == "cdtIfIpv4TcpMssAdjust" { return "Cdtifipv4Tcpmssadjust" }
    if yname == "cdtIfIpv4VerifyUniRpf" { return "Cdtifipv4Verifyunirpf" }
    if yname == "cdtIfIpv4VerifyUniRpfAcl" { return "Cdtifipv4Verifyunirpfacl" }
    if yname == "cdtIfIpv4VerifyUniRpfOpts" { return "Cdtifipv4Verifyunirpfopts" }
    if yname == "cdtIfIpv6Enable" { return "Cdtifipv6Enable" }
    if yname == "cdtIfIpv6SubEnable" { return "Cdtifipv6Subenable" }
    if yname == "cdtIfIpv6TcpMssAdjust" { return "Cdtifipv6Tcpmssadjust" }
    if yname == "cdtIfIpv6VerifyUniRpf" { return "Cdtifipv6Verifyunirpf" }
    if yname == "cdtIfIpv6VerifyUniRpfAcl" { return "Cdtifipv6Verifyunirpfacl" }
    if yname == "cdtIfIpv6VerifyUniRpfOpts" { return "Cdtifipv6Verifyunirpfopts" }
    if yname == "cdtIfIpv6NdPrefix" { return "Cdtifipv6Ndprefix" }
    if yname == "cdtIfIpv6NdPrefixLength" { return "Cdtifipv6Ndprefixlength" }
    if yname == "cdtIfIpv6NdValidLife" { return "Cdtifipv6Ndvalidlife" }
    if yname == "cdtIfIpv6NdPreferredLife" { return "Cdtifipv6Ndpreferredlife" }
    if yname == "cdtIfIpv6NdOpts" { return "Cdtifipv6Ndopts" }
    if yname == "cdtIfIpv6NdDadAttempts" { return "Cdtifipv6Nddadattempts" }
    if yname == "cdtIfIpv6NdNsInterval" { return "Cdtifipv6Ndnsinterval" }
    if yname == "cdtIfIpv6NdReachableTime" { return "Cdtifipv6Ndreachabletime" }
    if yname == "cdtIfIpv6NdRaIntervalUnits" { return "Cdtifipv6Ndraintervalunits" }
    if yname == "cdtIfIpv6NdRaIntervalMax" { return "Cdtifipv6Ndraintervalmax" }
    if yname == "cdtIfIpv6NdRaIntervalMin" { return "Cdtifipv6Ndraintervalmin" }
    if yname == "cdtIfIpv6NdRaLife" { return "Cdtifipv6Ndralife" }
    if yname == "cdtIfIpv6NdRouterPreference" { return "Cdtifipv6Ndrouterpreference" }
    return ""
}

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetSegmentPath() string {
    return "cdtIfTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtiftemplateentry.Cdttemplatename) + "']"
}

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateName"] = cdtiftemplateentry.Cdttemplatename
    leafs["cdtIfValid"] = cdtiftemplateentry.Cdtifvalid
    leafs["cdtIfMtu"] = cdtiftemplateentry.Cdtifmtu
    leafs["cdtIfCdpEnable"] = cdtiftemplateentry.Cdtifcdpenable
    leafs["cdtIfFlowMonitor"] = cdtiftemplateentry.Cdtifflowmonitor
    leafs["cdtIfIpv4Unnumbered"] = cdtiftemplateentry.Cdtifipv4Unnumbered
    leafs["cdtIfIpv4SubEnable"] = cdtiftemplateentry.Cdtifipv4Subenable
    leafs["cdtIfIpv4Mtu"] = cdtiftemplateentry.Cdtifipv4Mtu
    leafs["cdtIfIpv4TcpMssAdjust"] = cdtiftemplateentry.Cdtifipv4Tcpmssadjust
    leafs["cdtIfIpv4VerifyUniRpf"] = cdtiftemplateentry.Cdtifipv4Verifyunirpf
    leafs["cdtIfIpv4VerifyUniRpfAcl"] = cdtiftemplateentry.Cdtifipv4Verifyunirpfacl
    leafs["cdtIfIpv4VerifyUniRpfOpts"] = cdtiftemplateentry.Cdtifipv4Verifyunirpfopts
    leafs["cdtIfIpv6Enable"] = cdtiftemplateentry.Cdtifipv6Enable
    leafs["cdtIfIpv6SubEnable"] = cdtiftemplateentry.Cdtifipv6Subenable
    leafs["cdtIfIpv6TcpMssAdjust"] = cdtiftemplateentry.Cdtifipv6Tcpmssadjust
    leafs["cdtIfIpv6VerifyUniRpf"] = cdtiftemplateentry.Cdtifipv6Verifyunirpf
    leafs["cdtIfIpv6VerifyUniRpfAcl"] = cdtiftemplateentry.Cdtifipv6Verifyunirpfacl
    leafs["cdtIfIpv6VerifyUniRpfOpts"] = cdtiftemplateentry.Cdtifipv6Verifyunirpfopts
    leafs["cdtIfIpv6NdPrefix"] = cdtiftemplateentry.Cdtifipv6Ndprefix
    leafs["cdtIfIpv6NdPrefixLength"] = cdtiftemplateentry.Cdtifipv6Ndprefixlength
    leafs["cdtIfIpv6NdValidLife"] = cdtiftemplateentry.Cdtifipv6Ndvalidlife
    leafs["cdtIfIpv6NdPreferredLife"] = cdtiftemplateentry.Cdtifipv6Ndpreferredlife
    leafs["cdtIfIpv6NdOpts"] = cdtiftemplateentry.Cdtifipv6Ndopts
    leafs["cdtIfIpv6NdDadAttempts"] = cdtiftemplateentry.Cdtifipv6Nddadattempts
    leafs["cdtIfIpv6NdNsInterval"] = cdtiftemplateentry.Cdtifipv6Ndnsinterval
    leafs["cdtIfIpv6NdReachableTime"] = cdtiftemplateentry.Cdtifipv6Ndreachabletime
    leafs["cdtIfIpv6NdRaIntervalUnits"] = cdtiftemplateentry.Cdtifipv6Ndraintervalunits
    leafs["cdtIfIpv6NdRaIntervalMax"] = cdtiftemplateentry.Cdtifipv6Ndraintervalmax
    leafs["cdtIfIpv6NdRaIntervalMin"] = cdtiftemplateentry.Cdtifipv6Ndraintervalmin
    leafs["cdtIfIpv6NdRaLife"] = cdtiftemplateentry.Cdtifipv6Ndralife
    leafs["cdtIfIpv6NdRouterPreference"] = cdtiftemplateentry.Cdtifipv6Ndrouterpreference
    return leafs
}

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetYangName() string { return "cdtIfTemplateEntry" }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) SetParent(parent types.Entity) { cdtiftemplateentry.parent = parent }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetParent() types.Entity { return cdtiftemplateentry.parent }

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetParentYangName() string { return "cdtIfTemplateTable" }

// CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndraintervalunits represents corresponding instance of cdtIfValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndraintervalunits string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndraintervalunits_seconds CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndraintervalunits = "seconds"

    CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndraintervalunits_milliseconds CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndraintervalunits = "milliseconds"
)

// CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndrouterpreference represents the corresponding instance of cdtIfValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndrouterpreference string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndrouterpreference_high CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndrouterpreference = "high"

    CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndrouterpreference_medium CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndrouterpreference = "medium"

    CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndrouterpreference_low CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry_Cdtifipv6Ndrouterpreference = "low"
)

// CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable
// This table contains attributes relating to PPP connection
// configuration.
// 
// This table has a sparse-dependent relationship on the
// cdtTemplateTable, containing a row for each dynamic template
// having a cdtTemplateType of one of the following values:
// 
//     'derived'
//     'ppp'
type CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing attributes relating to PPP connection configuration. 
    // The system automatically creates an entry when the system or the EMS/NMS
    // creates a row in the cdtTemplateTable with a cdtTemplateType of one of the
    // following values:      'derived'     'ppp'  Likewise, the system
    // automatically destroys an entry when the system or the EMS/NMS destroys the
    // corresponding row in the cdtTemplateTable. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry.
    Cdtppptemplateentry []CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry
}

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetFilter() yfilter.YFilter { return cdtppptemplatetable.YFilter }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) SetFilter(yf yfilter.YFilter) { cdtppptemplatetable.YFilter = yf }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetGoName(yname string) string {
    if yname == "cdtPppTemplateEntry" { return "Cdtppptemplateentry" }
    return ""
}

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetSegmentPath() string {
    return "cdtPppTemplateTable"
}

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtPppTemplateEntry" {
        for _, c := range cdtppptemplatetable.Cdtppptemplateentry {
            if cdtppptemplatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry{}
        cdtppptemplatetable.Cdtppptemplateentry = append(cdtppptemplatetable.Cdtppptemplateentry, child)
        return &cdtppptemplatetable.Cdtppptemplateentry[len(cdtppptemplatetable.Cdtppptemplateentry)-1]
    }
    return nil
}

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdtppptemplatetable.Cdtppptemplateentry {
        children[cdtppptemplatetable.Cdtppptemplateentry[i].GetSegmentPath()] = &cdtppptemplatetable.Cdtppptemplateentry[i]
    }
    return children
}

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetBundleName() string { return "cisco_ios_xe" }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetYangName() string { return "cdtPppTemplateTable" }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) SetParent(parent types.Entity) { cdtppptemplatetable.parent = parent }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetParent() types.Entity { return cdtppptemplatetable.parent }

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry
// An entry containing attributes relating to PPP connection
// configuration.
// 
// The system automatically creates an entry when the system or the
// EMS/NMS creates a row in the cdtTemplateTable with a
// cdtTemplateType of one of the following values:
// 
//     'derived'
//     'ppp'
// 
// Likewise, the system automatically destroys an entry when the
// system or the EMS/NMS destroys the corresponding row in the
// cdtTemplateTable.
type CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatename
    Cdttemplatename interface{}

    // This object specifies which attributes in the dynamic template have been
    // configured to valid values.  Each bit in this bit string corresponds to a
    // column in this table.  If the bit is '0', then the value of the
    // corresponding column is not valid.  If the bit is '1', then the value of
    // the corresponding column has been configured to a valid value.  The
    // following list specifies the mappings between bits and the columns:     
    // accounting              => cdtPppAccounting     authentication          =>
    // cdtPppAuthentication     authenticationMethods   =>
    // cdtPppAuthenticationMethods     authorization           =>
    // cdtPppAuthorization     loopbackIgnore          => cdtPppLoopbackIgnore    
    // maxBadAuth              => cdtPppMaxBadAuth     maxConfigure            =>
    // cdtPppMaxConfigure     maxFailure              => cdtPppMaxFailure    
    // maxTerminate            => cdtPppMaxTerminate     timeoutAuthentication  
    // => cdtPppTimeoutAuthentication     timeoutRetry            =>
    // cdtPppTimeoutRetry     chapOpts                => cdtPppChapOpts    
    // chapHostname            => cdtPppChapHostname     chapPassword           
    // => cdtPppChapPassword     msChapV1Opts            => cdtPppMsChapV1Opts    
    // msChapV1Hostname        => cdtPppMsChapV1Hostname     msChapV1Password     
    // => cdtPppMsChapV1Password     msChapV2Opts            => cdtPppMsChapV2Opts
    // msChapV2Hostname        => cdtPppMsChapV2Hostname     msChapV2Password     
    // => cdtPppMsChapV2Password     papOpts                 => cdtPppPapOpts    
    // papSentUsername         => cdtPppPapUsername     papSentPassword         =>
    // cdtPppPapPassword     eapOpts                 => cdtPppEapOpts    
    // eapIdentity             => cdtPppEapIdentity     eapPassword             =>
    // cdtPppEapPassword     ipcpAddrOption          => cdtPppIpcpAddrOption    
    // ipcpDnsOption           => cdtPppIpcpDnsOption     ipcpDnsPrimary         
    // => cdtPppIpcpDnsPrimary     ipcpDnsSecondary        =>
    // cdtPppIpcpDnsSecondary     ipcpWinsOption          => cdtPppIpcpWinsOption 
    // ipcpWinsPrimary         => cdtPppIpcpWinsPrimary     ipcpWinsSecondary     
    // => cdtPppIpcpWinsSecondary     ipcpMaskOption          =>
    // cdtPppIpcpMaskOption     ipcpMask                => cdtPppIpcpMask    
    // peerDefIpAddrOpts       => cdtPppPeerOpts     peerDefIpAddrSrc        =>
    // cdtPppPeerDefIpAddrSrc     peerDefIpAddr           => cdtPppPeerDefIpAddr.
    // The type is map[string]bool.
    Cdtpppvalid interface{}

    // This object specifies whether the system applies accounting services to the
    // target PPP connection.  This column is valid only if the 'accounting' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is bool.
    Cdtpppaccounting interface{}

    // This object specifies authentication services applied to a target PPP
    // connection and other options affecting authentication services:      'chap'
    // This option enables the Challenge Handshake Protocol (CHAP)         on a
    // target PPP connection.      'msChap'         This option enables
    // Microsoft's CHAP on a target PPP         connection.      'msChapV2'       
    // This option enables version 2 of Microsoft's CHAP on a         target PPP
    // connection.      'pap'         This option enables Password Authentication
    // Protocol (PAP)         on a target PPP connection.      'eap'         This
    // option enables Extensible Authentication Protocol (EAP)         on a target
    // PPP connection.      'optional'         This option specifies that the
    // system accepts the connection         even if the peer of a target PPP
    // connection refuses to         accept the authentication methods the system
    // has         requested.      'callin'         This option specifies that
    // authentication should only happen         for incoming calls.     
    // 'oneTime'         This option specifies that the system accepts the
    // username         and password in the username field of authentication      
    // responses received on a target PPP connection.  This column is valid only
    // if the 'authentication' bit of the corresponding instance of cdtPppValid is
    // '1'. The type is map[string]bool.
    Cdtpppauthentication interface{}

    // This object specifies the name of a list of authentication methods used on
    // a target PPP connection.  If the template does not include this attribute,
    // then the system uses the default method list.  This column is valid only if
    // the 'authentication' bit of the corresponding instance of cdtPppValid is
    // '1'. The type is string with length: 0..255.
    Cdtpppauthenticationmethods interface{}

    // This object specifies whether the system applies authorization services to
    // a target PPP connection.  This column is valid only if the 'authorization'
    // bit of the corresponding instance of cditPppValid is '1'. The type is bool.
    Cdtpppauthorization interface{}

    // This object specifies whether the system ignores loopback on a target PPP
    // connection.  When the system ignores loopback, loopback detection is
    // disabled.  This column is valid only if the 'loopbackIgnore' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is bool.
    Cdtppploopbackignore interface{}

    // This object specifies the number of authentication failures allowed by the
    // system before a target PPP connection is reset.  The value '0' cannot be
    // written to an instance of this object. However, it serves as a convenient
    // value when the column is not valid.  This column is valid only if the
    // 'maxBadAuth' bit of the corresponding instance of cdtPppValid is '1'. The
    // type is interface{} with range: 0..4294967295.
    Cdtpppmaxbadauth interface{}

    // This object specifies the number of unacknowledged Configure-Request
    // messages a target PPP connection can send before the system abandons LCP or
    // NCP negotiations.  This column is valid only if the 'maxConfigure' bit of
    // the corresponding instance of cdtPppValid is '1'. The type is interface{}
    // with range: 1..4294967295.
    Cdtpppmaxconfigure interface{}

    // This object specifies the number of Configure-Nak messages a target PPP
    // connection can receive before the system abandons LCP or NCP negotiations. 
    // This column is valid only if the 'maxFailure' bit of the corresponding
    // instance of cdtPppValid is '1'. The type is interface{} with range:
    // 1..4294967295.
    Cdtpppmaxfailure interface{}

    // This object specifies the number of unacknowledged Terminate-Request
    // messages a target PPP connection can send before the system abandons LCP or
    // NCP negotiations.  This column is valid only if the 'maxTerminate' bit of
    // the corresponding instance of cdtPppValid is '1'. The type is interface{}
    // with range: 1..4294967295.
    Cdtpppmaxterminate interface{}

    // This objects specifies the maximum time the system will wait for a response
    // to an authentication request on a target PPP connection.  This column is
    // valid only if the 'timeoutAuthentication' bit of the corresponding instance
    // of cdtPppValid is '1'. The type is interface{} with range: 1..255. Units
    // are seconds.
    Cdtppptimeoutauthentication interface{}

    // This objects specifies the maximum time the system will wait for a response
    // to a PPP control packets on a target PPP connection.  This column is valid
    // only if the 'timeoutRetry' bit of the corresponding instance of cdtPppValid
    // is '1'. The type is interface{} with range: 1..255. Units are seconds.
    Cdtppptimeoutretry interface{}

    // This object specifies how the system processes the CHAP on a target PPP
    // connection:      'refuse'         This option specifies that the system
    // should refuse CHAP         requests from peers of a target PPP connection. 
    // 'callin'         This option specifies that the system should only refuse  
    // CHAP requests for incoming calls on a target PPP         connection.  This
    // option is only relevant if the         'refuse' option is set to '1'.     
    // 'wait'         This option delays CHAP authentication until after the      
    // peer of a target PPP connection has authenticated itself         to the
    // system.      'encrypted'         This option specifies that the value
    // specified by the         corresponding instance of cdtPppChapPassword is
    // already         encrypted.  This column is valid only if the 'chapOpts' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is
    // map[string]bool.
    Cdtpppchapopts interface{}

    // This object specifies the hostname sent in a CHAP response on a target PPP
    // connection.  If the template does not include this attribute, then the
    // system uses its assigned hostname.  This column is valid only if the
    // 'chapHostname' bit of the corresponding instance of cdtPppValid is '1'. The
    // type is string with length: 0..255.
    Cdtpppchaphostname interface{}

    // This object specifies the password used to construct a CHAP response on the
    // target PPP connection.  This column is valid only if the 'chapPassword' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is string
    // with length: 0..255.
    Cdtpppchappassword interface{}

    // This object specifies how the system processes version 1 of Microsoft CHAP
    // on a target PPP connection:      'refuse'         This option specifies
    // that the system should refuse         Microsoft CHAP (v1) requests from
    // peers of a target PPP         connection.      'callin'         This option
    // specifies that the system should only refuse         Microsoft CHAP (v1)
    // requests for incoming calls on a         target PPP connection.  This
    // option is only relevant if         the 'refuse' option is set to '1'.     
    // 'wait'         This option delays Microsoft CHAP (v1) authentication       
    // until after the peer of a target PPP connection has         authenticated
    // itself to the system.      'encrypted'         This option specifies that
    // the value specified by the         corresponding instance of
    // cdtPppMsChapV1Password is         already encrypted.  This column is valid
    // only if the 'msChapV1Opts' bit of the corresponding instance of cdtPppValid
    // is '1'. The type is map[string]bool.
    Cdtpppmschapv1Opts interface{}

    // This object specifies the hostname sent in a Microsoft CHAP (v1) response
    // on a target PPP connection.  If the template does not include this
    // attribute, then the system uses its assigned hostname.  This column is
    // valid only if the 'msChapV1Hostname' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is string with length: 0..255.
    Cdtpppmschapv1Hostname interface{}

    // This object specifies the password used to construct a Microsoft CHAP (v1)
    // response on a target PPP connection.  This column is valid only if the
    // 'msChapV1Password' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string with length: 0..255.
    Cdtpppmschapv1Password interface{}

    // This object specifies how the system processes version 2 of Microsoft CHAP
    // on a target PPP connection:      'refuse'         This option specifies
    // that the system should refuse         Microsoft CHAP (v2) requests from
    // peers of a target PPP         connection.      'callin'         This option
    // specifies that the system should only refuse         Microsoft CHAP (v2)
    // requests for incoming calls on a         target PPP connection.  This
    // option is only relevant if         the 'refuse' option is set to '1'.     
    // 'wait'         This option delays Microsoft CHAP (v2) authentication       
    // until after the peer of a target PPP connection has         authenticated
    // itself to the system.      'encrypted'         This option specifies that
    // the value specified by the         corresponding instance of
    // cdtPppMsChapV2Password is         already encrypted.  This column is valid
    // only if the 'msChapV2Opts' bit of the corresponding instance of cdtPppValid
    // is '1'. The type is map[string]bool.
    Cdtpppmschapv2Opts interface{}

    // This object specifies the hostname sent in a Microsoft CHAP (v2) response
    // on a target PPP connection.  If the template does not include this
    // attribute, then the system uses its assigned hostname.  This column is
    // valid only if the 'msChapV2Hostname' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is string with length: 0..255.
    Cdtpppmschapv2Hostname interface{}

    // This object specifies the password used to construct a Microsoft CHAP (v2)
    // response on a target PPP connection.  This column is valid only if the
    // 'msChapV2Password' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string with length: 0..255.
    Cdtpppmschapv2Password interface{}

    // This object specifies how the system processes the PAP on a target PPP
    // connection:      'refuse'         This option specifies that the system
    // should refuse PAP         requests from peers of a target PPP connection.  
    // 'encrypted'         This option specifies that the value specified by the  
    // corresponding instance of cdtPppPapSentPassword is         already
    // encrypted.  This column is valid only if the 'papOpts' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is map[string]bool.
    Cdtppppapopts interface{}

    // This object specifies the username sent in a PAP response on a target PPP
    // connection.  This column is valid only if the 'papUsername' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is string with
    // length: 0..255.
    Cdtppppapusername interface{}

    // This object specifies the username used to construct a PAP response on a
    // target PPP connection.  This column is valid only if the 'papPassword' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is string
    // with length: 0..255.
    Cdtppppappassword interface{}

    // This object specifies how the system processes the EAP on a target PPP
    // connection:      'refuse'         This option specifies that the system
    // should refuse EAP         requests from peers of a target PPP connection.  
    // 'callin'         This option specifies that the system should only refuse
    // EAP         requests for incoming calls on a target PPP connection.        
    // This option is only relevant if the 'refuse' option is         set to '1'. 
    // 'wait'         This option delays EAP authentication until after the       
    // peer of a target PPP connection has authenticated itself         to the
    // system.      'local'         This option specifies that the system should
    // locally         authenticate the peer of a target PPP connection,        
    // rather than acting as a proxy to an external AAA server.  This column is
    // valid only if the 'eapOpts' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is map[string]bool.
    Cdtpppeapopts interface{}

    // This object specifies the identity sent in an EAP response on a target PPP
    // connection.  This column is valid only if the 'eapIdentity' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is string with
    // length: 0..255.
    Cdtpppeapidentity interface{}

    // This object specifies the password used to construct an EAP response on a
    // target PPP connection.  This column is valid only if the 'eapPassword' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is string
    // with length: 0..255.
    Cdtpppeappassword interface{}

    // This object specifies the IPCP address option for a target PPP connection: 
    // 'other'         The implementation of this MIB module does not recognize   
    // the configured IPCP address option.      'accept'         The system
    // accepts any non-zero IP address from the peer         of a target PPP
    // connection.      'required'         The system disconnects the peer of a
    // target PPP         connection if it could not negotiate an IP address.     
    // 'unique'         The system disconnects the peer of a target PPP        
    // connection if the IP address is already in use.  This column is valid only
    // if the 'ipcpAddrOption' bit of the corresponding instance of cdtPppValid is
    // '1'. The type is Cdtpppipcpaddroption.
    Cdtpppipcpaddroption interface{}

    // This object specifies the IPCP DNS option for the dynamic interface:     
    // 'other'         The implementation of this MIB module does not recognize   
    // the configured DNS option.      'accept'         The system accepts any
    // non-zero DNS address form the         peer of a target PPP connection.     
    // 'request'         The system requests the DNS address from the peer of a   
    // target PPP connection.      'reject'         The system rejects the DNS
    // option from the peer of a         target PPP connection.  This column is
    // valid only if the 'ipcpDnsOption' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is Cdtpppipcpdnsoption.
    Cdtpppipcpdnsoption interface{}

    // This object specifies the IP address of the primary DNS server offered to
    // the peer of a target PPP connection.  This column is valid only if the
    // 'ipcpDnsPrimary' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string.
    Cdtpppipcpdnsprimary interface{}

    // This object specifies the IP address of the secondary DNS server offered to
    // the peer of a target PPP connection.  This column is valid only if the
    // 'ipcpDnsSecondary' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string.
    Cdtpppipcpdnssecondary interface{}

    // This object specifies the IPCP WINS option for a target PPP connection:    
    // 'other'         The implementation of this MIB module does not recognize   
    // the configured WINS option.      'accept'         The system accepts any
    // non-zero WINS address from the         peer of a target PPP connection.    
    // 'request'         The system requests the WINS address from the peer of    
    // a target PPP connection.      'reject'         The system rejects the WINS
    // option from the peer of a         target PPP connection.  This column is
    // valid only if the 'ipcpWinsOption' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is Cdtpppipcpwinsoption.
    Cdtpppipcpwinsoption interface{}

    // This object specifies the IP address of the primary WINS server offered to
    // the peer of a target PPP connection.  This column is valid only if the
    // 'ipcpWinsPrimary' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string.
    Cdtpppipcpwinsprimary interface{}

    // This object specifies the IP address of the secondary WINS server offered
    // to the peer of a target PPP connection.  This column is valid only if the
    // 'ipcpWinsSecondary' bit of the corresponding instance of cdtPppValid is
    // '1'. The type is string.
    Cdtpppipcpwinssecondary interface{}

    // This object specifies the IPCP IP subnet mask option for a target PPP
    // connection:      'other'         The implementation of this MIB module does
    // not recognize         the configured IP subnet mask option.      'request' 
    // The system requests the IP subnet mask from the peer of         a target
    // PPP connection.      'reject'         The system rejects the IP subnet mask
    // option from the         peer of a target PPP connection.  This column is
    // valid only if the 'ipcpMaskOption' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is Cdtpppipcpmaskoption.
    Cdtpppipcpmaskoption interface{}

    // This object specifies the IP address mask offered to the peer of a target
    // PPP connection.  This column is valid only if the 'ipcpMask' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is string.
    Cdtpppipcpmask interface{}

    // This object specifies options that affect how the system assigns an IP
    // address to the peer of a target PPP connection:      'ipAddrForced'        
    // This option forces the system to assign the next         available IP
    // address in the pool to the peer of a         target PPP connection.  When
    // disabled, the peer may         negotiate a specific IP address or the
    // system can assign         the peer its previously assigned IP address.     
    // 'matchAaaPools'         This option specifies that the names of the IP
    // address         pools provided by an external AAA server must appear in    
    // the corresponding list of IP address pool specified by         the
    // cdtPppPeerIpAddrPoolTable.      'backupPools'         This option specifies
    // that the corresponding names of         the IP address pools specified by
    // the         cditPppPeerIpAddrPoolTable serve as backup pools to        
    // those provided by an external AAA server.      'staticPools'         This
    // option suppresses an attempt to load pools from an         external AAA
    // server when the system encounters a missing         pool name.  This column
    // is valid only if the 'peerIpAddrOpts' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is map[string]bool.
    Cdtppppeerdefipaddropts interface{}

    // This object specifies how the system assigns an IP address to the peer of a
    // target PPP connection:      'static'         The system assigns the IP
    // address specified by the         corresponding instance of
    // cdtPppPeerDefIpAddr.      'pool'         The system allocates the first
    // available IP address from         the corresponding list of named pools
    // contained by the         cdtPppPeerIpAddrPoolTable.      'dhcp'         The
    // system acts as a DHCP proxy-client to obtain an IP         address.  This
    // column is valid only if the 'peerDefIpAddrSrc' bit of the corresponding
    // instance of cdtPppValid is '1'. The type is Cdtppppeerdefipaddrsrc.
    Cdtppppeerdefipaddrsrc interface{}

    // This object specifies the IP address the system assigns to the peer of a
    // target PPP connection.  This column is valid only if the 'peerDefIpAddr'
    // bit of the corresponding instance of cdtPppValid is '1'. The type is
    // string.
    Cdtppppeerdefipaddr interface{}
}

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetFilter() yfilter.YFilter { return cdtppptemplateentry.YFilter }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) SetFilter(yf yfilter.YFilter) { cdtppptemplateentry.YFilter = yf }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetGoName(yname string) string {
    if yname == "cdtTemplateName" { return "Cdttemplatename" }
    if yname == "cdtPppValid" { return "Cdtpppvalid" }
    if yname == "cdtPppAccounting" { return "Cdtpppaccounting" }
    if yname == "cdtPppAuthentication" { return "Cdtpppauthentication" }
    if yname == "cdtPppAuthenticationMethods" { return "Cdtpppauthenticationmethods" }
    if yname == "cdtPppAuthorization" { return "Cdtpppauthorization" }
    if yname == "cdtPppLoopbackIgnore" { return "Cdtppploopbackignore" }
    if yname == "cdtPppMaxBadAuth" { return "Cdtpppmaxbadauth" }
    if yname == "cdtPppMaxConfigure" { return "Cdtpppmaxconfigure" }
    if yname == "cdtPppMaxFailure" { return "Cdtpppmaxfailure" }
    if yname == "cdtPppMaxTerminate" { return "Cdtpppmaxterminate" }
    if yname == "cdtPppTimeoutAuthentication" { return "Cdtppptimeoutauthentication" }
    if yname == "cdtPppTimeoutRetry" { return "Cdtppptimeoutretry" }
    if yname == "cdtPppChapOpts" { return "Cdtpppchapopts" }
    if yname == "cdtPppChapHostname" { return "Cdtpppchaphostname" }
    if yname == "cdtPppChapPassword" { return "Cdtpppchappassword" }
    if yname == "cdtPppMsChapV1Opts" { return "Cdtpppmschapv1Opts" }
    if yname == "cdtPppMsChapV1Hostname" { return "Cdtpppmschapv1Hostname" }
    if yname == "cdtPppMsChapV1Password" { return "Cdtpppmschapv1Password" }
    if yname == "cdtPppMsChapV2Opts" { return "Cdtpppmschapv2Opts" }
    if yname == "cdtPppMsChapV2Hostname" { return "Cdtpppmschapv2Hostname" }
    if yname == "cdtPppMsChapV2Password" { return "Cdtpppmschapv2Password" }
    if yname == "cdtPppPapOpts" { return "Cdtppppapopts" }
    if yname == "cdtPppPapUsername" { return "Cdtppppapusername" }
    if yname == "cdtPppPapPassword" { return "Cdtppppappassword" }
    if yname == "cdtPppEapOpts" { return "Cdtpppeapopts" }
    if yname == "cdtPppEapIdentity" { return "Cdtpppeapidentity" }
    if yname == "cdtPppEapPassword" { return "Cdtpppeappassword" }
    if yname == "cdtPppIpcpAddrOption" { return "Cdtpppipcpaddroption" }
    if yname == "cdtPppIpcpDnsOption" { return "Cdtpppipcpdnsoption" }
    if yname == "cdtPppIpcpDnsPrimary" { return "Cdtpppipcpdnsprimary" }
    if yname == "cdtPppIpcpDnsSecondary" { return "Cdtpppipcpdnssecondary" }
    if yname == "cdtPppIpcpWinsOption" { return "Cdtpppipcpwinsoption" }
    if yname == "cdtPppIpcpWinsPrimary" { return "Cdtpppipcpwinsprimary" }
    if yname == "cdtPppIpcpWinsSecondary" { return "Cdtpppipcpwinssecondary" }
    if yname == "cdtPppIpcpMaskOption" { return "Cdtpppipcpmaskoption" }
    if yname == "cdtPppIpcpMask" { return "Cdtpppipcpmask" }
    if yname == "cdtPppPeerDefIpAddrOpts" { return "Cdtppppeerdefipaddropts" }
    if yname == "cdtPppPeerDefIpAddrSrc" { return "Cdtppppeerdefipaddrsrc" }
    if yname == "cdtPppPeerDefIpAddr" { return "Cdtppppeerdefipaddr" }
    return ""
}

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetSegmentPath() string {
    return "cdtPppTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtppptemplateentry.Cdttemplatename) + "']"
}

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateName"] = cdtppptemplateentry.Cdttemplatename
    leafs["cdtPppValid"] = cdtppptemplateentry.Cdtpppvalid
    leafs["cdtPppAccounting"] = cdtppptemplateentry.Cdtpppaccounting
    leafs["cdtPppAuthentication"] = cdtppptemplateentry.Cdtpppauthentication
    leafs["cdtPppAuthenticationMethods"] = cdtppptemplateentry.Cdtpppauthenticationmethods
    leafs["cdtPppAuthorization"] = cdtppptemplateentry.Cdtpppauthorization
    leafs["cdtPppLoopbackIgnore"] = cdtppptemplateentry.Cdtppploopbackignore
    leafs["cdtPppMaxBadAuth"] = cdtppptemplateentry.Cdtpppmaxbadauth
    leafs["cdtPppMaxConfigure"] = cdtppptemplateentry.Cdtpppmaxconfigure
    leafs["cdtPppMaxFailure"] = cdtppptemplateentry.Cdtpppmaxfailure
    leafs["cdtPppMaxTerminate"] = cdtppptemplateentry.Cdtpppmaxterminate
    leafs["cdtPppTimeoutAuthentication"] = cdtppptemplateentry.Cdtppptimeoutauthentication
    leafs["cdtPppTimeoutRetry"] = cdtppptemplateentry.Cdtppptimeoutretry
    leafs["cdtPppChapOpts"] = cdtppptemplateentry.Cdtpppchapopts
    leafs["cdtPppChapHostname"] = cdtppptemplateentry.Cdtpppchaphostname
    leafs["cdtPppChapPassword"] = cdtppptemplateentry.Cdtpppchappassword
    leafs["cdtPppMsChapV1Opts"] = cdtppptemplateentry.Cdtpppmschapv1Opts
    leafs["cdtPppMsChapV1Hostname"] = cdtppptemplateentry.Cdtpppmschapv1Hostname
    leafs["cdtPppMsChapV1Password"] = cdtppptemplateentry.Cdtpppmschapv1Password
    leafs["cdtPppMsChapV2Opts"] = cdtppptemplateentry.Cdtpppmschapv2Opts
    leafs["cdtPppMsChapV2Hostname"] = cdtppptemplateentry.Cdtpppmschapv2Hostname
    leafs["cdtPppMsChapV2Password"] = cdtppptemplateentry.Cdtpppmschapv2Password
    leafs["cdtPppPapOpts"] = cdtppptemplateentry.Cdtppppapopts
    leafs["cdtPppPapUsername"] = cdtppptemplateentry.Cdtppppapusername
    leafs["cdtPppPapPassword"] = cdtppptemplateentry.Cdtppppappassword
    leafs["cdtPppEapOpts"] = cdtppptemplateentry.Cdtpppeapopts
    leafs["cdtPppEapIdentity"] = cdtppptemplateentry.Cdtpppeapidentity
    leafs["cdtPppEapPassword"] = cdtppptemplateentry.Cdtpppeappassword
    leafs["cdtPppIpcpAddrOption"] = cdtppptemplateentry.Cdtpppipcpaddroption
    leafs["cdtPppIpcpDnsOption"] = cdtppptemplateentry.Cdtpppipcpdnsoption
    leafs["cdtPppIpcpDnsPrimary"] = cdtppptemplateentry.Cdtpppipcpdnsprimary
    leafs["cdtPppIpcpDnsSecondary"] = cdtppptemplateentry.Cdtpppipcpdnssecondary
    leafs["cdtPppIpcpWinsOption"] = cdtppptemplateentry.Cdtpppipcpwinsoption
    leafs["cdtPppIpcpWinsPrimary"] = cdtppptemplateentry.Cdtpppipcpwinsprimary
    leafs["cdtPppIpcpWinsSecondary"] = cdtppptemplateentry.Cdtpppipcpwinssecondary
    leafs["cdtPppIpcpMaskOption"] = cdtppptemplateentry.Cdtpppipcpmaskoption
    leafs["cdtPppIpcpMask"] = cdtppptemplateentry.Cdtpppipcpmask
    leafs["cdtPppPeerDefIpAddrOpts"] = cdtppptemplateentry.Cdtppppeerdefipaddropts
    leafs["cdtPppPeerDefIpAddrSrc"] = cdtppptemplateentry.Cdtppppeerdefipaddrsrc
    leafs["cdtPppPeerDefIpAddr"] = cdtppptemplateentry.Cdtppppeerdefipaddr
    return leafs
}

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetYangName() string { return "cdtPppTemplateEntry" }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) SetParent(parent types.Entity) { cdtppptemplateentry.parent = parent }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetParent() types.Entity { return cdtppptemplateentry.parent }

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetParentYangName() string { return "cdtPppTemplateTable" }

// CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption_other CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption = "other"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption_accept CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption = "accept"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption_required CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption = "required"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption_unique CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpaddroption = "unique"
)

// CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption_other CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption = "other"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption_accept CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption = "accept"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption_request CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption = "request"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption_reject CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpdnsoption = "reject"
)

// CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpmaskoption represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpmaskoption string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpmaskoption_other CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpmaskoption = "other"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpmaskoption_request CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpmaskoption = "request"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpmaskoption_reject CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpmaskoption = "reject"
)

// CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption_other CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption = "other"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption_accept CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption = "accept"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption_request CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption = "request"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption_reject CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtpppipcpwinsoption = "reject"
)

// CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtppppeerdefipaddrsrc represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtppppeerdefipaddrsrc string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtppppeerdefipaddrsrc_static CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtppppeerdefipaddrsrc = "static"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtppppeerdefipaddrsrc_pool CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtppppeerdefipaddrsrc = "pool"

    CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtppppeerdefipaddrsrc_dhcp CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry_Cdtppppeerdefipaddrsrc = "dhcp"
)

// CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable
// This table contains a prioritized list of named pools for each
// PPP template.  A list corresponding to a PPP template
// specifies the pools the system searches in an attempt to
// assign an IP address to the peer of a target PPP connection.
// The system searches the pools in the order of their priority.
// 
// This table has an expansion dependent relationship on the
// cdtPppTemplateTable, containing zero or more rows for each PPP
// template.
type CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry specifies a named pool in a list of pools associated with a PPP
    // template.  A PPP template can only have named pools associated with it if
    // it has a cdtPppPeerIpAddrSrc of 'pool'.  Any attempt to create an entry for
    // a PPP template that does not have a cdtPppPeerIpAddrSrc of 'pool' must
    // result in a response having an error-status of 'inconsistentValue'.  The
    // system automatically creates a corresponding entry when the system
    // associates a named pool with a PPP template through another management
    // entity (e.g., the system's local console).  The system automatically
    // destroys an entry under the following circumstances:  1)  The system or
    // EMS/NMS destroys the corresponding row in the     cdtTemplateTable.  2) 
    // The system or EMS/NMS sets the corresponding instance of    
    // cdtPppPeerIpAddrSrc with a value other than 'pool'. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry.
    Cdtppppeeripaddrpoolentry []CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry
}

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetFilter() yfilter.YFilter { return cdtppppeeripaddrpooltable.YFilter }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) SetFilter(yf yfilter.YFilter) { cdtppppeeripaddrpooltable.YFilter = yf }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetGoName(yname string) string {
    if yname == "cdtPppPeerIpAddrPoolEntry" { return "Cdtppppeeripaddrpoolentry" }
    return ""
}

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetSegmentPath() string {
    return "cdtPppPeerIpAddrPoolTable"
}

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtPppPeerIpAddrPoolEntry" {
        for _, c := range cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry {
            if cdtppppeeripaddrpooltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry{}
        cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry = append(cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry, child)
        return &cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry[len(cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry)-1]
    }
    return nil
}

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry {
        children[cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry[i].GetSegmentPath()] = &cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry[i]
    }
    return children
}

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetBundleName() string { return "cisco_ios_xe" }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetYangName() string { return "cdtPppPeerIpAddrPoolTable" }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) SetParent(parent types.Entity) { cdtppppeeripaddrpooltable.parent = parent }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetParent() types.Entity { return cdtppppeeripaddrpooltable.parent }

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry
// An entry specifies a named pool in a list of pools associated
// with a PPP template.  A PPP template can only have named
// pools associated with it if it has a cdtPppPeerIpAddrSrc of
// 'pool'.
// 
// Any attempt to create an entry for a PPP template that does not
// have a cdtPppPeerIpAddrSrc of 'pool' must result in a response
// having an error-status of 'inconsistentValue'.
// 
// The system automatically creates a corresponding entry when the
// system associates a named pool with a PPP template through
// another management entity (e.g., the system's local console).
// 
// The system automatically destroys an entry under the following
// circumstances:
// 
// 1)  The system or EMS/NMS destroys the corresponding row in the
//     cdtTemplateTable.
// 
// 2)  The system or EMS/NMS sets the corresponding instance of
//     cdtPppPeerIpAddrSrc with a value other than 'pool'.
type CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatename
    Cdttemplatename interface{}

    // This attribute is a key. This object indicates the relative priority of the
    // named pool in the list corresponding to a PPP template.  The system
    // searches pools in the order of priority, where lower values represent
    // higher priority. The type is interface{} with range: 1..4294967295.
    Cdtppppeeripaddrpoolpriority interface{}

    // This object specifies the status of the entry.  The following columns must
    // be valid before activating a subscriber access profile:      -
    // cdtPppPeerIpAddrPoolStorage     - cdtPppPeerIpAddrPoolName  However, these
    // objects specify a default value.  Thus, it is possible to use create-and-go
    // semantics without setting any additional columns.  An implementation must
    // not allow the EMS/NMS to create an entry if the corresponding instance of
    // cdtPppPeerIpAddrSrc is not 'pool'.  An implementation must allow the
    // EMS/NMS to modify any column when this column is 'active'. The type is
    // RowStatus.
    Cdtppppeeripaddrpoolstatus interface{}

    // This object specifies what happens to the name pool entry upon restart.  If
    // the corresponding instance of cdtTemplateSrc is not 'local', then this
    // column must be 'volatile'. The type is StorageType.
    Cdtppppeeripaddrpoolstorage interface{}

    // This object specifies the name of the IP address pool associated with the
    // PPP template. The type is string with length: 0..255.
    Cdtppppeeripaddrpoolname interface{}
}

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetFilter() yfilter.YFilter { return cdtppppeeripaddrpoolentry.YFilter }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) SetFilter(yf yfilter.YFilter) { cdtppppeeripaddrpoolentry.YFilter = yf }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetGoName(yname string) string {
    if yname == "cdtTemplateName" { return "Cdttemplatename" }
    if yname == "cdtPppPeerIpAddrPoolPriority" { return "Cdtppppeeripaddrpoolpriority" }
    if yname == "cdtPppPeerIpAddrPoolStatus" { return "Cdtppppeeripaddrpoolstatus" }
    if yname == "cdtPppPeerIpAddrPoolStorage" { return "Cdtppppeeripaddrpoolstorage" }
    if yname == "cdtPppPeerIpAddrPoolName" { return "Cdtppppeeripaddrpoolname" }
    return ""
}

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetSegmentPath() string {
    return "cdtPppPeerIpAddrPoolEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtppppeeripaddrpoolentry.Cdttemplatename) + "']" + "[cdtPppPeerIpAddrPoolPriority='" + fmt.Sprintf("%v", cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolpriority) + "']"
}

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateName"] = cdtppppeeripaddrpoolentry.Cdttemplatename
    leafs["cdtPppPeerIpAddrPoolPriority"] = cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolpriority
    leafs["cdtPppPeerIpAddrPoolStatus"] = cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolstatus
    leafs["cdtPppPeerIpAddrPoolStorage"] = cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolstorage
    leafs["cdtPppPeerIpAddrPoolName"] = cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolname
    return leafs
}

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetYangName() string { return "cdtPppPeerIpAddrPoolEntry" }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) SetParent(parent types.Entity) { cdtppppeeripaddrpoolentry.parent = parent }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetParent() types.Entity { return cdtppppeeripaddrpoolentry.parent }

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetParentYangName() string { return "cdtPppPeerIpAddrPoolTable" }

// CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable
// This table contains attributes relating to dynamic interfaces
// initiated on Ethernet virtual interfaces (e.g., EoMPLS) or
// automatically created VLANs.
// 
// This table has a sparse-dependent relationship on the
// cdtTemplateTable, containing a row for each dynamic template
// having a cdtTemplateType of one of the following values:
// 
//     'derived'
//     'ethernet'
type CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing attributes relating to dynamic interfaces initiated on
    // Ethernet virtual interfaces (e.g., EoMPLS) or automatically created VLANs. 
    // The system automatically creates an entry when the system or the EMS/NMS
    // creates a row in the cdtTemplateTable with a cdtTemplateType of one of the
    // following values:      'derived'     'ethernet'  Likewise, the system
    // automatically destroys an entry when the system or the EMS/NMS destroys the
    // corresponding row in the cdtTemplateTable. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry.
    Cdtethernettemplateentry []CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry
}

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetFilter() yfilter.YFilter { return cdtethernettemplatetable.YFilter }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) SetFilter(yf yfilter.YFilter) { cdtethernettemplatetable.YFilter = yf }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetGoName(yname string) string {
    if yname == "cdtEthernetTemplateEntry" { return "Cdtethernettemplateentry" }
    return ""
}

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetSegmentPath() string {
    return "cdtEthernetTemplateTable"
}

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtEthernetTemplateEntry" {
        for _, c := range cdtethernettemplatetable.Cdtethernettemplateentry {
            if cdtethernettemplatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry{}
        cdtethernettemplatetable.Cdtethernettemplateentry = append(cdtethernettemplatetable.Cdtethernettemplateentry, child)
        return &cdtethernettemplatetable.Cdtethernettemplateentry[len(cdtethernettemplatetable.Cdtethernettemplateentry)-1]
    }
    return nil
}

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdtethernettemplatetable.Cdtethernettemplateentry {
        children[cdtethernettemplatetable.Cdtethernettemplateentry[i].GetSegmentPath()] = &cdtethernettemplatetable.Cdtethernettemplateentry[i]
    }
    return children
}

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetBundleName() string { return "cisco_ios_xe" }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetYangName() string { return "cdtEthernetTemplateTable" }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) SetParent(parent types.Entity) { cdtethernettemplatetable.parent = parent }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetParent() types.Entity { return cdtethernettemplatetable.parent }

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry
// An entry containing attributes relating to dynamic interfaces
// initiated on Ethernet virtual interfaces (e.g., EoMPLS) or
// automatically created VLANs.
// 
// The system automatically creates an entry when the system or the
// EMS/NMS creates a row in the cdtTemplateTable with a
// cdtTemplateType of one of the following values:
// 
//     'derived'
//     'ethernet'
// 
// Likewise, the system automatically destroys an entry when the
// system or the EMS/NMS destroys the corresponding row in the
// cdtTemplateTable.
type CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatename
    Cdttemplatename interface{}

    // This object specifies which attributes in the dynamic template have been
    // configured to valid values.  Each bit in this bit string corresponds to a
    // column in this table.  If the bit is '0', then the value of the
    // corresponding column is not valid.  If the bit is '1', then the value of
    // the corresponding column has been configured to a valid value.  The
    // following list specifies the mappings between bits and the columns:     
    // bridgeDomain     => cdtEthernetBridgeDomain     pppoeEnable      =>
    // cdtEthernetPppoeEnable     ipv4PointToPoint => cdtEthernetIpv4PointToPoint 
    // macAddr          => cdtEthernetMacAddr. The type is map[string]bool.
    Cdtethernetvalid interface{}

    // This object specifies the name of the bridge domain... The type is string
    // with length: 0..255.
    Cdtethernetbridgedomain interface{}

    // This object specifies whether... The type is bool.
    Cdtethernetpppoeenable interface{}

    // This object specifies whether... The type is bool.
    Cdtethernetipv4Pointtopoint interface{}

    // This object specifies the... The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    Cdtethernetmacaddr interface{}
}

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetFilter() yfilter.YFilter { return cdtethernettemplateentry.YFilter }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) SetFilter(yf yfilter.YFilter) { cdtethernettemplateentry.YFilter = yf }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetGoName(yname string) string {
    if yname == "cdtTemplateName" { return "Cdttemplatename" }
    if yname == "cdtEthernetValid" { return "Cdtethernetvalid" }
    if yname == "cdtEthernetBridgeDomain" { return "Cdtethernetbridgedomain" }
    if yname == "cdtEthernetPppoeEnable" { return "Cdtethernetpppoeenable" }
    if yname == "cdtEthernetIpv4PointToPoint" { return "Cdtethernetipv4Pointtopoint" }
    if yname == "cdtEthernetMacAddr" { return "Cdtethernetmacaddr" }
    return ""
}

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetSegmentPath() string {
    return "cdtEthernetTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtethernettemplateentry.Cdttemplatename) + "']"
}

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateName"] = cdtethernettemplateentry.Cdttemplatename
    leafs["cdtEthernetValid"] = cdtethernettemplateentry.Cdtethernetvalid
    leafs["cdtEthernetBridgeDomain"] = cdtethernettemplateentry.Cdtethernetbridgedomain
    leafs["cdtEthernetPppoeEnable"] = cdtethernettemplateentry.Cdtethernetpppoeenable
    leafs["cdtEthernetIpv4PointToPoint"] = cdtethernettemplateentry.Cdtethernetipv4Pointtopoint
    leafs["cdtEthernetMacAddr"] = cdtethernettemplateentry.Cdtethernetmacaddr
    return leafs
}

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetYangName() string { return "cdtEthernetTemplateEntry" }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) SetParent(parent types.Entity) { cdtethernettemplateentry.parent = parent }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetParent() types.Entity { return cdtethernettemplateentry.parent }

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetParentYangName() string { return "cdtEthernetTemplateTable" }

// CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable
// This table contains attributes relating to a service.
// 
// This table has a sparse-dependent relationship on the
// cdtTemplateTable, containing a row for each dynamic template
// having a cdtTemplateType of one of the following values:
// 
//     'derived'
//     'service'
type CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing attributes relating to a service.  The system
    // automatically creates entry when the system or the EMS/NMS creates a row in
    // the cdtTemplateTable with a cdtTemplateType of one of the following values:
    // 'derived'     'service'  Likewise, the system automatically destroys an
    // entry when the system or the EMS/NMS destroys the corresponding row in the
    // cdtTemplateTable. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry.
    Cdtsrvtemplateentry []CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry
}

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetFilter() yfilter.YFilter { return cdtsrvtemplatetable.YFilter }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) SetFilter(yf yfilter.YFilter) { cdtsrvtemplatetable.YFilter = yf }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetGoName(yname string) string {
    if yname == "cdtSrvTemplateEntry" { return "Cdtsrvtemplateentry" }
    return ""
}

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetSegmentPath() string {
    return "cdtSrvTemplateTable"
}

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cdtSrvTemplateEntry" {
        for _, c := range cdtsrvtemplatetable.Cdtsrvtemplateentry {
            if cdtsrvtemplatetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry{}
        cdtsrvtemplatetable.Cdtsrvtemplateentry = append(cdtsrvtemplatetable.Cdtsrvtemplateentry, child)
        return &cdtsrvtemplatetable.Cdtsrvtemplateentry[len(cdtsrvtemplatetable.Cdtsrvtemplateentry)-1]
    }
    return nil
}

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdtsrvtemplatetable.Cdtsrvtemplateentry {
        children[cdtsrvtemplatetable.Cdtsrvtemplateentry[i].GetSegmentPath()] = &cdtsrvtemplatetable.Cdtsrvtemplateentry[i]
    }
    return children
}

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetBundleName() string { return "cisco_ios_xe" }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetYangName() string { return "cdtSrvTemplateTable" }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) SetParent(parent types.Entity) { cdtsrvtemplatetable.parent = parent }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetParent() types.Entity { return cdtsrvtemplatetable.parent }

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetParentYangName() string { return "CISCO-DYNAMIC-TEMPLATE-MIB" }

// CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry
// An entry containing attributes relating to a service.
// 
// The system automatically creates entry when the system or the
// EMS/NMS creates a row in the cdtTemplateTable with a
// cdtTemplateType of one of the following values:
// 
//     'derived'
//     'service'
// 
// Likewise, the system automatically destroys an entry when the
// system or the EMS/NMS destroys the corresponding row in the
// cdtTemplateTable.
type CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry_Cdttemplatename
    Cdttemplatename interface{}

    // This object specifies which attributes in the dynamic template have been
    // configured to valid values.  Each bit in this bit string corresponds to a
    // column in this table.  If the bit is '0', then the value of the
    // corresponding column is not valid.  If the bit is '1', then the value of
    // the corresponding column has been configured to a valid value.  The
    // following list specifies the mappings between bits and the columns:     
    // networkSrv     => cdtSrvNetworkSrv     vpdnGroup      => cdtSrvVpdnGroup   
    // sgSrvGroup     => cdtSrvGroup     sgSrvType      => cdtSrvSgSrvType    
    // multicast      => cdtSrvMulticast. The type is map[string]bool.
    Cdtsrvvalid interface{}

    // This object specifies the type of network service provided by the target
    // service:      'other'         The implementation of this MIB module does
    // not recognize         the configured network service.      'none'        
    // The target subscriber service does not provide a network         service to
    // subscribers sessions.      'local'         The target subscriber service
    // provides local termination         for subscriber sessions.      'vpdn'    
    // The target subscriber service provides a Virtual Private         Dialup
    // Network service for subscriber sessions.  This column is valid only if the
    // 'networkSrv' bit of the corresponding instance of cdtSrvValid is '1'. The
    // type is Cdtsrvnetworksrv.
    Cdtsrvnetworksrv interface{}

    // This object specifies the name of the VPDN group used to configure the
    // network service.  This column is valid only if the 'vpdnGroup' bit of the
    // corresponding instance of cdtSrvValid is '1'. The type is string with
    // length: 0..255.
    Cdtsrvvpdngroup interface{}

    // This object specifies the name of the service group with which the system
    // associates subscriber sessions.  A service group specifies a set of
    // services that may be active simultaneously for a given subscriber session. 
    // Typically, a service group contains a primary service and one or more
    // secondary services.  This column is valid only if the 'sgSrvGroup' bit of
    // the corresponding instance of cdtSrvValid is '1'. The type is string with
    // length: 0..255.
    Cdtsrvsgsrvgroup interface{}

    // This object specifies whether the target service specifies a
    // network-forwarding policy:      'primary'         The target service
    // specifies a network-forwarding         policy.  Primary services are
    // mutually exclusive; that         is, only one primary service can be
    // activated for any         given subscriber session.      'secondary'       
    // The target service has a dependence on the primary         service in the
    // group specified by the corresponding         instance of
    // cdtSuBSrvSgSrvGroup.  After the system         activates the primary
    // service, it activates secondary         services.  When the system
    // deactivates the primary         service, then it deactivates all the
    // secondary services         in the service group.  This column is valid only
    // if the 'sgSrvType' bit of the corresponding instance of cdtSrvValid is '1'.
    // The type is Cdtsrvsgsrvtype.
    Cdtsrvsgsrvtype interface{}

    // This objects specifies whether the system enables multicast service for
    // subscriber sessions of the target service.  This column is valid only if
    // the 'sgSrvMcastRoutingIf' bit of the corresponding instance of cdtSrvValid
    // is '1'. The type is bool.
    Cdtsrvmulticast interface{}
}

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetFilter() yfilter.YFilter { return cdtsrvtemplateentry.YFilter }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) SetFilter(yf yfilter.YFilter) { cdtsrvtemplateentry.YFilter = yf }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetGoName(yname string) string {
    if yname == "cdtTemplateName" { return "Cdttemplatename" }
    if yname == "cdtSrvValid" { return "Cdtsrvvalid" }
    if yname == "cdtSrvNetworkSrv" { return "Cdtsrvnetworksrv" }
    if yname == "cdtSrvVpdnGroup" { return "Cdtsrvvpdngroup" }
    if yname == "cdtSrvSgSrvGroup" { return "Cdtsrvsgsrvgroup" }
    if yname == "cdtSrvSgSrvType" { return "Cdtsrvsgsrvtype" }
    if yname == "cdtSrvMulticast" { return "Cdtsrvmulticast" }
    return ""
}

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetSegmentPath() string {
    return "cdtSrvTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtsrvtemplateentry.Cdttemplatename) + "']"
}

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cdtTemplateName"] = cdtsrvtemplateentry.Cdttemplatename
    leafs["cdtSrvValid"] = cdtsrvtemplateentry.Cdtsrvvalid
    leafs["cdtSrvNetworkSrv"] = cdtsrvtemplateentry.Cdtsrvnetworksrv
    leafs["cdtSrvVpdnGroup"] = cdtsrvtemplateentry.Cdtsrvvpdngroup
    leafs["cdtSrvSgSrvGroup"] = cdtsrvtemplateentry.Cdtsrvsgsrvgroup
    leafs["cdtSrvSgSrvType"] = cdtsrvtemplateentry.Cdtsrvsgsrvtype
    leafs["cdtSrvMulticast"] = cdtsrvtemplateentry.Cdtsrvmulticast
    return leafs
}

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetBundleName() string { return "cisco_ios_xe" }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetYangName() string { return "cdtSrvTemplateEntry" }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) SetParent(parent types.Entity) { cdtsrvtemplateentry.parent = parent }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetParent() types.Entity { return cdtsrvtemplateentry.parent }

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetParentYangName() string { return "cdtSrvTemplateTable" }

// CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv represents corresponding instance of cdtSrvValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv_other CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv = "other"

    CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv_none CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv = "none"

    CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv_local CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv = "local"

    CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv_vpdn CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvnetworksrv = "vpdn"
)

// CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvsgsrvtype represents corresponding instance of cdtSrvValid is '1'.
type CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvsgsrvtype string

const (
    CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvsgsrvtype_primary CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvsgsrvtype = "primary"

    CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvsgsrvtype_secondary CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry_Cdtsrvsgsrvtype = "secondary"
)

