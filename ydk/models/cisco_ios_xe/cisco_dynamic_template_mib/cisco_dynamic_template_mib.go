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
    EntityData types.CommonEntityData
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

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetEntityData() *types.CommonEntityData {
    cISCODYNAMICTEMPLATEMIB.EntityData.YFilter = cISCODYNAMICTEMPLATEMIB.YFilter
    cISCODYNAMICTEMPLATEMIB.EntityData.YangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cISCODYNAMICTEMPLATEMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCODYNAMICTEMPLATEMIB.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cISCODYNAMICTEMPLATEMIB.EntityData.SegmentPath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB"
    cISCODYNAMICTEMPLATEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCODYNAMICTEMPLATEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCODYNAMICTEMPLATEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCODYNAMICTEMPLATEMIB.EntityData.Children = make(map[string]types.YChild)
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtTemplateTable"] = types.YChild{"Cdttemplatetable", &cISCODYNAMICTEMPLATEMIB.Cdttemplatetable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtTemplateTargetTable"] = types.YChild{"Cdttemplatetargettable", &cISCODYNAMICTEMPLATEMIB.Cdttemplatetargettable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtTemplateAssociationTable"] = types.YChild{"Cdttemplateassociationtable", &cISCODYNAMICTEMPLATEMIB.Cdttemplateassociationtable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtTemplateUsageTable"] = types.YChild{"Cdttemplateusagetable", &cISCODYNAMICTEMPLATEMIB.Cdttemplateusagetable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtTemplateCommonTable"] = types.YChild{"Cdttemplatecommontable", &cISCODYNAMICTEMPLATEMIB.Cdttemplatecommontable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtIfTemplateTable"] = types.YChild{"Cdtiftemplatetable", &cISCODYNAMICTEMPLATEMIB.Cdtiftemplatetable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtPppTemplateTable"] = types.YChild{"Cdtppptemplatetable", &cISCODYNAMICTEMPLATEMIB.Cdtppptemplatetable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtPppPeerIpAddrPoolTable"] = types.YChild{"Cdtppppeeripaddrpooltable", &cISCODYNAMICTEMPLATEMIB.Cdtppppeeripaddrpooltable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtEthernetTemplateTable"] = types.YChild{"Cdtethernettemplatetable", &cISCODYNAMICTEMPLATEMIB.Cdtethernettemplatetable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Children["cdtSrvTemplateTable"] = types.YChild{"Cdtsrvtemplatetable", &cISCODYNAMICTEMPLATEMIB.Cdtsrvtemplatetable}
    cISCODYNAMICTEMPLATEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCODYNAMICTEMPLATEMIB.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_Cdttemplatetable
// This table lists the dynamic templates maintained by the
// system, including those that have been locally-configured on the
// system and those pushed to the system by external policy
// servers.
type CISCODYNAMICTEMPLATEMIB_Cdttemplatetable struct {
    EntityData types.CommonEntityData
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

func (cdttemplatetable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable) GetEntityData() *types.CommonEntityData {
    cdttemplatetable.EntityData.YFilter = cdttemplatetable.YFilter
    cdttemplatetable.EntityData.YangName = "cdtTemplateTable"
    cdttemplatetable.EntityData.BundleName = "cisco_ios_xe"
    cdttemplatetable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdttemplatetable.EntityData.SegmentPath = "cdtTemplateTable"
    cdttemplatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplatetable.EntityData.Children = make(map[string]types.YChild)
    cdttemplatetable.EntityData.Children["cdtTemplateEntry"] = types.YChild{"Cdttemplateentry", nil}
    for i := range cdttemplatetable.Cdttemplateentry {
        cdttemplatetable.EntityData.Children[types.GetSegmentPath(&cdttemplatetable.Cdttemplateentry[i])] = types.YChild{"Cdttemplateentry", &cdttemplatetable.Cdttemplateentry[i]}
    }
    cdttemplatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdttemplatetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdttemplateentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetable_Cdttemplateentry) GetEntityData() *types.CommonEntityData {
    cdttemplateentry.EntityData.YFilter = cdttemplateentry.YFilter
    cdttemplateentry.EntityData.YangName = "cdtTemplateEntry"
    cdttemplateentry.EntityData.BundleName = "cisco_ios_xe"
    cdttemplateentry.EntityData.ParentYangName = "cdtTemplateTable"
    cdttemplateentry.EntityData.SegmentPath = "cdtTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdttemplateentry.Cdttemplatename) + "']"
    cdttemplateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplateentry.EntityData.Children = make(map[string]types.YChild)
    cdttemplateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdttemplateentry.EntityData.Leafs["cdtTemplateName"] = types.YLeaf{"Cdttemplatename", cdttemplateentry.Cdttemplatename}
    cdttemplateentry.EntityData.Leafs["cdtTemplateStatus"] = types.YLeaf{"Cdttemplatestatus", cdttemplateentry.Cdttemplatestatus}
    cdttemplateentry.EntityData.Leafs["cdtTemplateStorage"] = types.YLeaf{"Cdttemplatestorage", cdttemplateentry.Cdttemplatestorage}
    cdttemplateentry.EntityData.Leafs["cdtTemplateType"] = types.YLeaf{"Cdttemplatetype", cdttemplateentry.Cdttemplatetype}
    cdttemplateentry.EntityData.Leafs["cdtTemplateSrc"] = types.YLeaf{"Cdttemplatesrc", cdttemplateentry.Cdttemplatesrc}
    cdttemplateentry.EntityData.Leafs["cdtTemplateUsageCount"] = types.YLeaf{"Cdttemplateusagecount", cdttemplateentry.Cdttemplateusagecount}
    return &(cdttemplateentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry describes a target associated with one or more dynamic templates. 
    // The system automatically creates an entry when it associates a dynamic
    // template to a target.  Likewise, the system automatically destroys an entry
    // when a target no longer has any associated dynamic templates. The type is
    // slice of
    // CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry.
    Cdttemplatetargetentry []CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry
}

func (cdttemplatetargettable *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable) GetEntityData() *types.CommonEntityData {
    cdttemplatetargettable.EntityData.YFilter = cdttemplatetargettable.YFilter
    cdttemplatetargettable.EntityData.YangName = "cdtTemplateTargetTable"
    cdttemplatetargettable.EntityData.BundleName = "cisco_ios_xe"
    cdttemplatetargettable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdttemplatetargettable.EntityData.SegmentPath = "cdtTemplateTargetTable"
    cdttemplatetargettable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplatetargettable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplatetargettable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplatetargettable.EntityData.Children = make(map[string]types.YChild)
    cdttemplatetargettable.EntityData.Children["cdtTemplateTargetEntry"] = types.YChild{"Cdttemplatetargetentry", nil}
    for i := range cdttemplatetargettable.Cdttemplatetargetentry {
        cdttemplatetargettable.EntityData.Children[types.GetSegmentPath(&cdttemplatetargettable.Cdttemplatetargetentry[i])] = types.YChild{"Cdttemplatetargetentry", &cdttemplatetargettable.Cdttemplatetargetentry[i]}
    }
    cdttemplatetargettable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdttemplatetargettable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry
// An entry describes a target associated with one or more
// dynamic templates.
// 
// The system automatically creates an entry when it associates a
// dynamic template to a target.  Likewise, the system
// automatically destroys an entry when a target no longer has any
// associated dynamic templates.
type CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry struct {
    EntityData types.CommonEntityData
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

func (cdttemplatetargetentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatetargettable_Cdttemplatetargetentry) GetEntityData() *types.CommonEntityData {
    cdttemplatetargetentry.EntityData.YFilter = cdttemplatetargetentry.YFilter
    cdttemplatetargetentry.EntityData.YangName = "cdtTemplateTargetEntry"
    cdttemplatetargetentry.EntityData.BundleName = "cisco_ios_xe"
    cdttemplatetargetentry.EntityData.ParentYangName = "cdtTemplateTargetTable"
    cdttemplatetargetentry.EntityData.SegmentPath = "cdtTemplateTargetEntry" + "[cdtTemplateTargetType='" + fmt.Sprintf("%v", cdttemplatetargetentry.Cdttemplatetargettype) + "']" + "[cdtTemplateTargetId='" + fmt.Sprintf("%v", cdttemplatetargetentry.Cdttemplatetargetid) + "']"
    cdttemplatetargetentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplatetargetentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplatetargetentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplatetargetentry.EntityData.Children = make(map[string]types.YChild)
    cdttemplatetargetentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdttemplatetargetentry.EntityData.Leafs["cdtTemplateTargetType"] = types.YLeaf{"Cdttemplatetargettype", cdttemplatetargetentry.Cdttemplatetargettype}
    cdttemplatetargetentry.EntityData.Leafs["cdtTemplateTargetId"] = types.YLeaf{"Cdttemplatetargetid", cdttemplatetargetentry.Cdttemplatetargetid}
    cdttemplatetargetentry.EntityData.Leafs["cdtTemplateTargetStatus"] = types.YLeaf{"Cdttemplatetargetstatus", cdttemplatetargetentry.Cdttemplatetargetstatus}
    cdttemplatetargetentry.EntityData.Leafs["cdtTemplateTargetStorage"] = types.YLeaf{"Cdttemplatetargetstorage", cdttemplatetargetentry.Cdttemplatetargetstorage}
    return &(cdttemplatetargetentry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable
// This table contains a list of templates associated with each
// target.
// 
// This table has an expansion dependent relationship on the
// cdtTemplateTargetTable, containing zero or more rows for each
// target.
type CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable struct {
    EntityData types.CommonEntityData
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

func (cdttemplateassociationtable *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable) GetEntityData() *types.CommonEntityData {
    cdttemplateassociationtable.EntityData.YFilter = cdttemplateassociationtable.YFilter
    cdttemplateassociationtable.EntityData.YangName = "cdtTemplateAssociationTable"
    cdttemplateassociationtable.EntityData.BundleName = "cisco_ios_xe"
    cdttemplateassociationtable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdttemplateassociationtable.EntityData.SegmentPath = "cdtTemplateAssociationTable"
    cdttemplateassociationtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplateassociationtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplateassociationtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplateassociationtable.EntityData.Children = make(map[string]types.YChild)
    cdttemplateassociationtable.EntityData.Children["cdtTemplateAssociationEntry"] = types.YChild{"Cdttemplateassociationentry", nil}
    for i := range cdttemplateassociationtable.Cdttemplateassociationentry {
        cdttemplateassociationtable.EntityData.Children[types.GetSegmentPath(&cdttemplateassociationtable.Cdttemplateassociationentry[i])] = types.YChild{"Cdttemplateassociationentry", &cdttemplateassociationtable.Cdttemplateassociationentry[i]}
    }
    cdttemplateassociationtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdttemplateassociationtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdttemplateassociationentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateassociationtable_Cdttemplateassociationentry) GetEntityData() *types.CommonEntityData {
    cdttemplateassociationentry.EntityData.YFilter = cdttemplateassociationentry.YFilter
    cdttemplateassociationentry.EntityData.YangName = "cdtTemplateAssociationEntry"
    cdttemplateassociationentry.EntityData.BundleName = "cisco_ios_xe"
    cdttemplateassociationentry.EntityData.ParentYangName = "cdtTemplateAssociationTable"
    cdttemplateassociationentry.EntityData.SegmentPath = "cdtTemplateAssociationEntry" + "[cdtTemplateTargetType='" + fmt.Sprintf("%v", cdttemplateassociationentry.Cdttemplatetargettype) + "']" + "[cdtTemplateTargetId='" + fmt.Sprintf("%v", cdttemplateassociationentry.Cdttemplatetargetid) + "']" + "[cdtTemplateAssociationName='" + fmt.Sprintf("%v", cdttemplateassociationentry.Cdttemplateassociationname) + "']"
    cdttemplateassociationentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplateassociationentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplateassociationentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplateassociationentry.EntityData.Children = make(map[string]types.YChild)
    cdttemplateassociationentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdttemplateassociationentry.EntityData.Leafs["cdtTemplateTargetType"] = types.YLeaf{"Cdttemplatetargettype", cdttemplateassociationentry.Cdttemplatetargettype}
    cdttemplateassociationentry.EntityData.Leafs["cdtTemplateTargetId"] = types.YLeaf{"Cdttemplatetargetid", cdttemplateassociationentry.Cdttemplatetargetid}
    cdttemplateassociationentry.EntityData.Leafs["cdtTemplateAssociationName"] = types.YLeaf{"Cdttemplateassociationname", cdttemplateassociationentry.Cdttemplateassociationname}
    cdttemplateassociationentry.EntityData.Leafs["cdtTemplateAssociationPrecedence"] = types.YLeaf{"Cdttemplateassociationprecedence", cdttemplateassociationentry.Cdttemplateassociationprecedence}
    return &(cdttemplateassociationentry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable
// This table contains a list of targets using each dynamic
// template.
// 
// This table has an expansion dependent relationship on the
// cdtTemplateTable, containing zero or more rows for each
// dynamic template.
type CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable struct {
    EntityData types.CommonEntityData
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

func (cdttemplateusagetable *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable) GetEntityData() *types.CommonEntityData {
    cdttemplateusagetable.EntityData.YFilter = cdttemplateusagetable.YFilter
    cdttemplateusagetable.EntityData.YangName = "cdtTemplateUsageTable"
    cdttemplateusagetable.EntityData.BundleName = "cisco_ios_xe"
    cdttemplateusagetable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdttemplateusagetable.EntityData.SegmentPath = "cdtTemplateUsageTable"
    cdttemplateusagetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplateusagetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplateusagetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplateusagetable.EntityData.Children = make(map[string]types.YChild)
    cdttemplateusagetable.EntityData.Children["cdtTemplateUsageEntry"] = types.YChild{"Cdttemplateusageentry", nil}
    for i := range cdttemplateusagetable.Cdttemplateusageentry {
        cdttemplateusagetable.EntityData.Children[types.GetSegmentPath(&cdttemplateusagetable.Cdttemplateusageentry[i])] = types.YChild{"Cdttemplateusageentry", &cdttemplateusagetable.Cdttemplateusageentry[i]}
    }
    cdttemplateusagetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdttemplateusagetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdttemplateusageentry *CISCODYNAMICTEMPLATEMIB_Cdttemplateusagetable_Cdttemplateusageentry) GetEntityData() *types.CommonEntityData {
    cdttemplateusageentry.EntityData.YFilter = cdttemplateusageentry.YFilter
    cdttemplateusageentry.EntityData.YangName = "cdtTemplateUsageEntry"
    cdttemplateusageentry.EntityData.BundleName = "cisco_ios_xe"
    cdttemplateusageentry.EntityData.ParentYangName = "cdtTemplateUsageTable"
    cdttemplateusageentry.EntityData.SegmentPath = "cdtTemplateUsageEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdttemplateusageentry.Cdttemplatename) + "']" + "[cdtTemplateUsageTargetType='" + fmt.Sprintf("%v", cdttemplateusageentry.Cdttemplateusagetargettype) + "']" + "[cdtTemplateUsageTargetId='" + fmt.Sprintf("%v", cdttemplateusageentry.Cdttemplateusagetargetid) + "']"
    cdttemplateusageentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplateusageentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplateusageentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplateusageentry.EntityData.Children = make(map[string]types.YChild)
    cdttemplateusageentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdttemplateusageentry.EntityData.Leafs["cdtTemplateName"] = types.YLeaf{"Cdttemplatename", cdttemplateusageentry.Cdttemplatename}
    cdttemplateusageentry.EntityData.Leafs["cdtTemplateUsageTargetType"] = types.YLeaf{"Cdttemplateusagetargettype", cdttemplateusageentry.Cdttemplateusagetargettype}
    cdttemplateusageentry.EntityData.Leafs["cdtTemplateUsageTargetId"] = types.YLeaf{"Cdttemplateusagetargetid", cdttemplateusageentry.Cdttemplateusagetargetid}
    return &(cdttemplateusageentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdttemplatecommontable *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable) GetEntityData() *types.CommonEntityData {
    cdttemplatecommontable.EntityData.YFilter = cdttemplatecommontable.YFilter
    cdttemplatecommontable.EntityData.YangName = "cdtTemplateCommonTable"
    cdttemplatecommontable.EntityData.BundleName = "cisco_ios_xe"
    cdttemplatecommontable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdttemplatecommontable.EntityData.SegmentPath = "cdtTemplateCommonTable"
    cdttemplatecommontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplatecommontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplatecommontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplatecommontable.EntityData.Children = make(map[string]types.YChild)
    cdttemplatecommontable.EntityData.Children["cdtTemplateCommonEntry"] = types.YChild{"Cdttemplatecommonentry", nil}
    for i := range cdttemplatecommontable.Cdttemplatecommonentry {
        cdttemplatecommontable.EntityData.Children[types.GetSegmentPath(&cdttemplatecommontable.Cdttemplatecommonentry[i])] = types.YChild{"Cdttemplatecommonentry", &cdttemplatecommontable.Cdttemplatecommonentry[i]}
    }
    cdttemplatecommontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdttemplatecommontable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdttemplatecommonentry *CISCODYNAMICTEMPLATEMIB_Cdttemplatecommontable_Cdttemplatecommonentry) GetEntityData() *types.CommonEntityData {
    cdttemplatecommonentry.EntityData.YFilter = cdttemplatecommonentry.YFilter
    cdttemplatecommonentry.EntityData.YangName = "cdtTemplateCommonEntry"
    cdttemplatecommonentry.EntityData.BundleName = "cisco_ios_xe"
    cdttemplatecommonentry.EntityData.ParentYangName = "cdtTemplateCommonTable"
    cdttemplatecommonentry.EntityData.SegmentPath = "cdtTemplateCommonEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdttemplatecommonentry.Cdttemplatename) + "']"
    cdttemplatecommonentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdttemplatecommonentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdttemplatecommonentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdttemplatecommonentry.EntityData.Children = make(map[string]types.YChild)
    cdttemplatecommonentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdttemplatecommonentry.EntityData.Leafs["cdtTemplateName"] = types.YLeaf{"Cdttemplatename", cdttemplatecommonentry.Cdttemplatename}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonValid"] = types.YLeaf{"Cdtcommonvalid", cdttemplatecommonentry.Cdtcommonvalid}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonDescr"] = types.YLeaf{"Cdtcommondescr", cdttemplatecommonentry.Cdtcommondescr}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonKeepaliveInt"] = types.YLeaf{"Cdtcommonkeepaliveint", cdttemplatecommonentry.Cdtcommonkeepaliveint}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonKeepaliveRetries"] = types.YLeaf{"Cdtcommonkeepaliveretries", cdttemplatecommonentry.Cdtcommonkeepaliveretries}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonVrf"] = types.YLeaf{"Cdtcommonvrf", cdttemplatecommonentry.Cdtcommonvrf}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonAddrPool"] = types.YLeaf{"Cdtcommonaddrpool", cdttemplatecommonentry.Cdtcommonaddrpool}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonIpv4AccessGroup"] = types.YLeaf{"Cdtcommonipv4Accessgroup", cdttemplatecommonentry.Cdtcommonipv4Accessgroup}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonIpv4Unreachables"] = types.YLeaf{"Cdtcommonipv4Unreachables", cdttemplatecommonentry.Cdtcommonipv4Unreachables}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonIpv6AccessGroup"] = types.YLeaf{"Cdtcommonipv6Accessgroup", cdttemplatecommonentry.Cdtcommonipv6Accessgroup}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonIpv6Unreachables"] = types.YLeaf{"Cdtcommonipv6Unreachables", cdttemplatecommonentry.Cdtcommonipv6Unreachables}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonSrvSubControl"] = types.YLeaf{"Cdtcommonsrvsubcontrol", cdttemplatecommonentry.Cdtcommonsrvsubcontrol}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonSrvRedirect"] = types.YLeaf{"Cdtcommonsrvredirect", cdttemplatecommonentry.Cdtcommonsrvredirect}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonSrvAcct"] = types.YLeaf{"Cdtcommonsrvacct", cdttemplatecommonentry.Cdtcommonsrvacct}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonSrvQos"] = types.YLeaf{"Cdtcommonsrvqos", cdttemplatecommonentry.Cdtcommonsrvqos}
    cdttemplatecommonentry.EntityData.Leafs["cdtCommonSrvNetflow"] = types.YLeaf{"Cdtcommonsrvnetflow", cdttemplatecommonentry.Cdtcommonsrvnetflow}
    return &(cdttemplatecommonentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtiftemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable) GetEntityData() *types.CommonEntityData {
    cdtiftemplatetable.EntityData.YFilter = cdtiftemplatetable.YFilter
    cdtiftemplatetable.EntityData.YangName = "cdtIfTemplateTable"
    cdtiftemplatetable.EntityData.BundleName = "cisco_ios_xe"
    cdtiftemplatetable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtiftemplatetable.EntityData.SegmentPath = "cdtIfTemplateTable"
    cdtiftemplatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtiftemplatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtiftemplatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtiftemplatetable.EntityData.Children = make(map[string]types.YChild)
    cdtiftemplatetable.EntityData.Children["cdtIfTemplateEntry"] = types.YChild{"Cdtiftemplateentry", nil}
    for i := range cdtiftemplatetable.Cdtiftemplateentry {
        cdtiftemplatetable.EntityData.Children[types.GetSegmentPath(&cdtiftemplatetable.Cdtiftemplateentry[i])] = types.YChild{"Cdtiftemplateentry", &cdtiftemplatetable.Cdtiftemplateentry[i]}
    }
    cdtiftemplatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdtiftemplatetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtiftemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtiftemplatetable_Cdtiftemplateentry) GetEntityData() *types.CommonEntityData {
    cdtiftemplateentry.EntityData.YFilter = cdtiftemplateentry.YFilter
    cdtiftemplateentry.EntityData.YangName = "cdtIfTemplateEntry"
    cdtiftemplateentry.EntityData.BundleName = "cisco_ios_xe"
    cdtiftemplateentry.EntityData.ParentYangName = "cdtIfTemplateTable"
    cdtiftemplateentry.EntityData.SegmentPath = "cdtIfTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtiftemplateentry.Cdttemplatename) + "']"
    cdtiftemplateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtiftemplateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtiftemplateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtiftemplateentry.EntityData.Children = make(map[string]types.YChild)
    cdtiftemplateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdtiftemplateentry.EntityData.Leafs["cdtTemplateName"] = types.YLeaf{"Cdttemplatename", cdtiftemplateentry.Cdttemplatename}
    cdtiftemplateentry.EntityData.Leafs["cdtIfValid"] = types.YLeaf{"Cdtifvalid", cdtiftemplateentry.Cdtifvalid}
    cdtiftemplateentry.EntityData.Leafs["cdtIfMtu"] = types.YLeaf{"Cdtifmtu", cdtiftemplateentry.Cdtifmtu}
    cdtiftemplateentry.EntityData.Leafs["cdtIfCdpEnable"] = types.YLeaf{"Cdtifcdpenable", cdtiftemplateentry.Cdtifcdpenable}
    cdtiftemplateentry.EntityData.Leafs["cdtIfFlowMonitor"] = types.YLeaf{"Cdtifflowmonitor", cdtiftemplateentry.Cdtifflowmonitor}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv4Unnumbered"] = types.YLeaf{"Cdtifipv4Unnumbered", cdtiftemplateentry.Cdtifipv4Unnumbered}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv4SubEnable"] = types.YLeaf{"Cdtifipv4Subenable", cdtiftemplateentry.Cdtifipv4Subenable}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv4Mtu"] = types.YLeaf{"Cdtifipv4Mtu", cdtiftemplateentry.Cdtifipv4Mtu}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv4TcpMssAdjust"] = types.YLeaf{"Cdtifipv4Tcpmssadjust", cdtiftemplateentry.Cdtifipv4Tcpmssadjust}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv4VerifyUniRpf"] = types.YLeaf{"Cdtifipv4Verifyunirpf", cdtiftemplateentry.Cdtifipv4Verifyunirpf}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv4VerifyUniRpfAcl"] = types.YLeaf{"Cdtifipv4Verifyunirpfacl", cdtiftemplateentry.Cdtifipv4Verifyunirpfacl}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv4VerifyUniRpfOpts"] = types.YLeaf{"Cdtifipv4Verifyunirpfopts", cdtiftemplateentry.Cdtifipv4Verifyunirpfopts}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6Enable"] = types.YLeaf{"Cdtifipv6Enable", cdtiftemplateentry.Cdtifipv6Enable}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6SubEnable"] = types.YLeaf{"Cdtifipv6Subenable", cdtiftemplateentry.Cdtifipv6Subenable}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6TcpMssAdjust"] = types.YLeaf{"Cdtifipv6Tcpmssadjust", cdtiftemplateentry.Cdtifipv6Tcpmssadjust}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6VerifyUniRpf"] = types.YLeaf{"Cdtifipv6Verifyunirpf", cdtiftemplateentry.Cdtifipv6Verifyunirpf}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6VerifyUniRpfAcl"] = types.YLeaf{"Cdtifipv6Verifyunirpfacl", cdtiftemplateentry.Cdtifipv6Verifyunirpfacl}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6VerifyUniRpfOpts"] = types.YLeaf{"Cdtifipv6Verifyunirpfopts", cdtiftemplateentry.Cdtifipv6Verifyunirpfopts}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdPrefix"] = types.YLeaf{"Cdtifipv6Ndprefix", cdtiftemplateentry.Cdtifipv6Ndprefix}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdPrefixLength"] = types.YLeaf{"Cdtifipv6Ndprefixlength", cdtiftemplateentry.Cdtifipv6Ndprefixlength}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdValidLife"] = types.YLeaf{"Cdtifipv6Ndvalidlife", cdtiftemplateentry.Cdtifipv6Ndvalidlife}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdPreferredLife"] = types.YLeaf{"Cdtifipv6Ndpreferredlife", cdtiftemplateentry.Cdtifipv6Ndpreferredlife}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdOpts"] = types.YLeaf{"Cdtifipv6Ndopts", cdtiftemplateentry.Cdtifipv6Ndopts}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdDadAttempts"] = types.YLeaf{"Cdtifipv6Nddadattempts", cdtiftemplateentry.Cdtifipv6Nddadattempts}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdNsInterval"] = types.YLeaf{"Cdtifipv6Ndnsinterval", cdtiftemplateentry.Cdtifipv6Ndnsinterval}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdReachableTime"] = types.YLeaf{"Cdtifipv6Ndreachabletime", cdtiftemplateentry.Cdtifipv6Ndreachabletime}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdRaIntervalUnits"] = types.YLeaf{"Cdtifipv6Ndraintervalunits", cdtiftemplateentry.Cdtifipv6Ndraintervalunits}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdRaIntervalMax"] = types.YLeaf{"Cdtifipv6Ndraintervalmax", cdtiftemplateentry.Cdtifipv6Ndraintervalmax}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdRaIntervalMin"] = types.YLeaf{"Cdtifipv6Ndraintervalmin", cdtiftemplateentry.Cdtifipv6Ndraintervalmin}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdRaLife"] = types.YLeaf{"Cdtifipv6Ndralife", cdtiftemplateentry.Cdtifipv6Ndralife}
    cdtiftemplateentry.EntityData.Leafs["cdtIfIpv6NdRouterPreference"] = types.YLeaf{"Cdtifipv6Ndrouterpreference", cdtiftemplateentry.Cdtifipv6Ndrouterpreference}
    return &(cdtiftemplateentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtppptemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable) GetEntityData() *types.CommonEntityData {
    cdtppptemplatetable.EntityData.YFilter = cdtppptemplatetable.YFilter
    cdtppptemplatetable.EntityData.YangName = "cdtPppTemplateTable"
    cdtppptemplatetable.EntityData.BundleName = "cisco_ios_xe"
    cdtppptemplatetable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtppptemplatetable.EntityData.SegmentPath = "cdtPppTemplateTable"
    cdtppptemplatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtppptemplatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtppptemplatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtppptemplatetable.EntityData.Children = make(map[string]types.YChild)
    cdtppptemplatetable.EntityData.Children["cdtPppTemplateEntry"] = types.YChild{"Cdtppptemplateentry", nil}
    for i := range cdtppptemplatetable.Cdtppptemplateentry {
        cdtppptemplatetable.EntityData.Children[types.GetSegmentPath(&cdtppptemplatetable.Cdtppptemplateentry[i])] = types.YChild{"Cdtppptemplateentry", &cdtppptemplatetable.Cdtppptemplateentry[i]}
    }
    cdtppptemplatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdtppptemplatetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtppptemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtppptemplatetable_Cdtppptemplateentry) GetEntityData() *types.CommonEntityData {
    cdtppptemplateentry.EntityData.YFilter = cdtppptemplateentry.YFilter
    cdtppptemplateentry.EntityData.YangName = "cdtPppTemplateEntry"
    cdtppptemplateentry.EntityData.BundleName = "cisco_ios_xe"
    cdtppptemplateentry.EntityData.ParentYangName = "cdtPppTemplateTable"
    cdtppptemplateentry.EntityData.SegmentPath = "cdtPppTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtppptemplateentry.Cdttemplatename) + "']"
    cdtppptemplateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtppptemplateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtppptemplateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtppptemplateentry.EntityData.Children = make(map[string]types.YChild)
    cdtppptemplateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdtppptemplateentry.EntityData.Leafs["cdtTemplateName"] = types.YLeaf{"Cdttemplatename", cdtppptemplateentry.Cdttemplatename}
    cdtppptemplateentry.EntityData.Leafs["cdtPppValid"] = types.YLeaf{"Cdtpppvalid", cdtppptemplateentry.Cdtpppvalid}
    cdtppptemplateentry.EntityData.Leafs["cdtPppAccounting"] = types.YLeaf{"Cdtpppaccounting", cdtppptemplateentry.Cdtpppaccounting}
    cdtppptemplateentry.EntityData.Leafs["cdtPppAuthentication"] = types.YLeaf{"Cdtpppauthentication", cdtppptemplateentry.Cdtpppauthentication}
    cdtppptemplateentry.EntityData.Leafs["cdtPppAuthenticationMethods"] = types.YLeaf{"Cdtpppauthenticationmethods", cdtppptemplateentry.Cdtpppauthenticationmethods}
    cdtppptemplateentry.EntityData.Leafs["cdtPppAuthorization"] = types.YLeaf{"Cdtpppauthorization", cdtppptemplateentry.Cdtpppauthorization}
    cdtppptemplateentry.EntityData.Leafs["cdtPppLoopbackIgnore"] = types.YLeaf{"Cdtppploopbackignore", cdtppptemplateentry.Cdtppploopbackignore}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMaxBadAuth"] = types.YLeaf{"Cdtpppmaxbadauth", cdtppptemplateentry.Cdtpppmaxbadauth}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMaxConfigure"] = types.YLeaf{"Cdtpppmaxconfigure", cdtppptemplateentry.Cdtpppmaxconfigure}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMaxFailure"] = types.YLeaf{"Cdtpppmaxfailure", cdtppptemplateentry.Cdtpppmaxfailure}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMaxTerminate"] = types.YLeaf{"Cdtpppmaxterminate", cdtppptemplateentry.Cdtpppmaxterminate}
    cdtppptemplateentry.EntityData.Leafs["cdtPppTimeoutAuthentication"] = types.YLeaf{"Cdtppptimeoutauthentication", cdtppptemplateentry.Cdtppptimeoutauthentication}
    cdtppptemplateentry.EntityData.Leafs["cdtPppTimeoutRetry"] = types.YLeaf{"Cdtppptimeoutretry", cdtppptemplateentry.Cdtppptimeoutretry}
    cdtppptemplateentry.EntityData.Leafs["cdtPppChapOpts"] = types.YLeaf{"Cdtpppchapopts", cdtppptemplateentry.Cdtpppchapopts}
    cdtppptemplateentry.EntityData.Leafs["cdtPppChapHostname"] = types.YLeaf{"Cdtpppchaphostname", cdtppptemplateentry.Cdtpppchaphostname}
    cdtppptemplateentry.EntityData.Leafs["cdtPppChapPassword"] = types.YLeaf{"Cdtpppchappassword", cdtppptemplateentry.Cdtpppchappassword}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMsChapV1Opts"] = types.YLeaf{"Cdtpppmschapv1Opts", cdtppptemplateentry.Cdtpppmschapv1Opts}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMsChapV1Hostname"] = types.YLeaf{"Cdtpppmschapv1Hostname", cdtppptemplateentry.Cdtpppmschapv1Hostname}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMsChapV1Password"] = types.YLeaf{"Cdtpppmschapv1Password", cdtppptemplateentry.Cdtpppmschapv1Password}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMsChapV2Opts"] = types.YLeaf{"Cdtpppmschapv2Opts", cdtppptemplateentry.Cdtpppmschapv2Opts}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMsChapV2Hostname"] = types.YLeaf{"Cdtpppmschapv2Hostname", cdtppptemplateentry.Cdtpppmschapv2Hostname}
    cdtppptemplateentry.EntityData.Leafs["cdtPppMsChapV2Password"] = types.YLeaf{"Cdtpppmschapv2Password", cdtppptemplateentry.Cdtpppmschapv2Password}
    cdtppptemplateentry.EntityData.Leafs["cdtPppPapOpts"] = types.YLeaf{"Cdtppppapopts", cdtppptemplateentry.Cdtppppapopts}
    cdtppptemplateentry.EntityData.Leafs["cdtPppPapUsername"] = types.YLeaf{"Cdtppppapusername", cdtppptemplateentry.Cdtppppapusername}
    cdtppptemplateentry.EntityData.Leafs["cdtPppPapPassword"] = types.YLeaf{"Cdtppppappassword", cdtppptemplateentry.Cdtppppappassword}
    cdtppptemplateentry.EntityData.Leafs["cdtPppEapOpts"] = types.YLeaf{"Cdtpppeapopts", cdtppptemplateentry.Cdtpppeapopts}
    cdtppptemplateentry.EntityData.Leafs["cdtPppEapIdentity"] = types.YLeaf{"Cdtpppeapidentity", cdtppptemplateentry.Cdtpppeapidentity}
    cdtppptemplateentry.EntityData.Leafs["cdtPppEapPassword"] = types.YLeaf{"Cdtpppeappassword", cdtppptemplateentry.Cdtpppeappassword}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpAddrOption"] = types.YLeaf{"Cdtpppipcpaddroption", cdtppptemplateentry.Cdtpppipcpaddroption}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpDnsOption"] = types.YLeaf{"Cdtpppipcpdnsoption", cdtppptemplateentry.Cdtpppipcpdnsoption}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpDnsPrimary"] = types.YLeaf{"Cdtpppipcpdnsprimary", cdtppptemplateentry.Cdtpppipcpdnsprimary}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpDnsSecondary"] = types.YLeaf{"Cdtpppipcpdnssecondary", cdtppptemplateentry.Cdtpppipcpdnssecondary}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpWinsOption"] = types.YLeaf{"Cdtpppipcpwinsoption", cdtppptemplateentry.Cdtpppipcpwinsoption}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpWinsPrimary"] = types.YLeaf{"Cdtpppipcpwinsprimary", cdtppptemplateentry.Cdtpppipcpwinsprimary}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpWinsSecondary"] = types.YLeaf{"Cdtpppipcpwinssecondary", cdtppptemplateentry.Cdtpppipcpwinssecondary}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpMaskOption"] = types.YLeaf{"Cdtpppipcpmaskoption", cdtppptemplateentry.Cdtpppipcpmaskoption}
    cdtppptemplateentry.EntityData.Leafs["cdtPppIpcpMask"] = types.YLeaf{"Cdtpppipcpmask", cdtppptemplateentry.Cdtpppipcpmask}
    cdtppptemplateentry.EntityData.Leafs["cdtPppPeerDefIpAddrOpts"] = types.YLeaf{"Cdtppppeerdefipaddropts", cdtppptemplateentry.Cdtppppeerdefipaddropts}
    cdtppptemplateentry.EntityData.Leafs["cdtPppPeerDefIpAddrSrc"] = types.YLeaf{"Cdtppppeerdefipaddrsrc", cdtppptemplateentry.Cdtppppeerdefipaddrsrc}
    cdtppptemplateentry.EntityData.Leafs["cdtPppPeerDefIpAddr"] = types.YLeaf{"Cdtppppeerdefipaddr", cdtppptemplateentry.Cdtppppeerdefipaddr}
    return &(cdtppptemplateentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtppppeeripaddrpooltable *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable) GetEntityData() *types.CommonEntityData {
    cdtppppeeripaddrpooltable.EntityData.YFilter = cdtppppeeripaddrpooltable.YFilter
    cdtppppeeripaddrpooltable.EntityData.YangName = "cdtPppPeerIpAddrPoolTable"
    cdtppppeeripaddrpooltable.EntityData.BundleName = "cisco_ios_xe"
    cdtppppeeripaddrpooltable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtppppeeripaddrpooltable.EntityData.SegmentPath = "cdtPppPeerIpAddrPoolTable"
    cdtppppeeripaddrpooltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtppppeeripaddrpooltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtppppeeripaddrpooltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtppppeeripaddrpooltable.EntityData.Children = make(map[string]types.YChild)
    cdtppppeeripaddrpooltable.EntityData.Children["cdtPppPeerIpAddrPoolEntry"] = types.YChild{"Cdtppppeeripaddrpoolentry", nil}
    for i := range cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry {
        cdtppppeeripaddrpooltable.EntityData.Children[types.GetSegmentPath(&cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry[i])] = types.YChild{"Cdtppppeeripaddrpoolentry", &cdtppppeeripaddrpooltable.Cdtppppeeripaddrpoolentry[i]}
    }
    cdtppppeeripaddrpooltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdtppppeeripaddrpooltable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtppppeeripaddrpoolentry *CISCODYNAMICTEMPLATEMIB_Cdtppppeeripaddrpooltable_Cdtppppeeripaddrpoolentry) GetEntityData() *types.CommonEntityData {
    cdtppppeeripaddrpoolentry.EntityData.YFilter = cdtppppeeripaddrpoolentry.YFilter
    cdtppppeeripaddrpoolentry.EntityData.YangName = "cdtPppPeerIpAddrPoolEntry"
    cdtppppeeripaddrpoolentry.EntityData.BundleName = "cisco_ios_xe"
    cdtppppeeripaddrpoolentry.EntityData.ParentYangName = "cdtPppPeerIpAddrPoolTable"
    cdtppppeeripaddrpoolentry.EntityData.SegmentPath = "cdtPppPeerIpAddrPoolEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtppppeeripaddrpoolentry.Cdttemplatename) + "']" + "[cdtPppPeerIpAddrPoolPriority='" + fmt.Sprintf("%v", cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolpriority) + "']"
    cdtppppeeripaddrpoolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtppppeeripaddrpoolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtppppeeripaddrpoolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtppppeeripaddrpoolentry.EntityData.Children = make(map[string]types.YChild)
    cdtppppeeripaddrpoolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdtppppeeripaddrpoolentry.EntityData.Leafs["cdtTemplateName"] = types.YLeaf{"Cdttemplatename", cdtppppeeripaddrpoolentry.Cdttemplatename}
    cdtppppeeripaddrpoolentry.EntityData.Leafs["cdtPppPeerIpAddrPoolPriority"] = types.YLeaf{"Cdtppppeeripaddrpoolpriority", cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolpriority}
    cdtppppeeripaddrpoolentry.EntityData.Leafs["cdtPppPeerIpAddrPoolStatus"] = types.YLeaf{"Cdtppppeeripaddrpoolstatus", cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolstatus}
    cdtppppeeripaddrpoolentry.EntityData.Leafs["cdtPppPeerIpAddrPoolStorage"] = types.YLeaf{"Cdtppppeeripaddrpoolstorage", cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolstorage}
    cdtppppeeripaddrpoolentry.EntityData.Leafs["cdtPppPeerIpAddrPoolName"] = types.YLeaf{"Cdtppppeeripaddrpoolname", cdtppppeeripaddrpoolentry.Cdtppppeeripaddrpoolname}
    return &(cdtppppeeripaddrpoolentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtethernettemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable) GetEntityData() *types.CommonEntityData {
    cdtethernettemplatetable.EntityData.YFilter = cdtethernettemplatetable.YFilter
    cdtethernettemplatetable.EntityData.YangName = "cdtEthernetTemplateTable"
    cdtethernettemplatetable.EntityData.BundleName = "cisco_ios_xe"
    cdtethernettemplatetable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtethernettemplatetable.EntityData.SegmentPath = "cdtEthernetTemplateTable"
    cdtethernettemplatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtethernettemplatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtethernettemplatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtethernettemplatetable.EntityData.Children = make(map[string]types.YChild)
    cdtethernettemplatetable.EntityData.Children["cdtEthernetTemplateEntry"] = types.YChild{"Cdtethernettemplateentry", nil}
    for i := range cdtethernettemplatetable.Cdtethernettemplateentry {
        cdtethernettemplatetable.EntityData.Children[types.GetSegmentPath(&cdtethernettemplatetable.Cdtethernettemplateentry[i])] = types.YChild{"Cdtethernettemplateentry", &cdtethernettemplatetable.Cdtethernettemplateentry[i]}
    }
    cdtethernettemplatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdtethernettemplatetable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Cdtethernetmacaddr interface{}
}

func (cdtethernettemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtethernettemplatetable_Cdtethernettemplateentry) GetEntityData() *types.CommonEntityData {
    cdtethernettemplateentry.EntityData.YFilter = cdtethernettemplateentry.YFilter
    cdtethernettemplateentry.EntityData.YangName = "cdtEthernetTemplateEntry"
    cdtethernettemplateentry.EntityData.BundleName = "cisco_ios_xe"
    cdtethernettemplateentry.EntityData.ParentYangName = "cdtEthernetTemplateTable"
    cdtethernettemplateentry.EntityData.SegmentPath = "cdtEthernetTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtethernettemplateentry.Cdttemplatename) + "']"
    cdtethernettemplateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtethernettemplateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtethernettemplateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtethernettemplateentry.EntityData.Children = make(map[string]types.YChild)
    cdtethernettemplateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdtethernettemplateentry.EntityData.Leafs["cdtTemplateName"] = types.YLeaf{"Cdttemplatename", cdtethernettemplateentry.Cdttemplatename}
    cdtethernettemplateentry.EntityData.Leafs["cdtEthernetValid"] = types.YLeaf{"Cdtethernetvalid", cdtethernettemplateentry.Cdtethernetvalid}
    cdtethernettemplateentry.EntityData.Leafs["cdtEthernetBridgeDomain"] = types.YLeaf{"Cdtethernetbridgedomain", cdtethernettemplateentry.Cdtethernetbridgedomain}
    cdtethernettemplateentry.EntityData.Leafs["cdtEthernetPppoeEnable"] = types.YLeaf{"Cdtethernetpppoeenable", cdtethernettemplateentry.Cdtethernetpppoeenable}
    cdtethernettemplateentry.EntityData.Leafs["cdtEthernetIpv4PointToPoint"] = types.YLeaf{"Cdtethernetipv4Pointtopoint", cdtethernettemplateentry.Cdtethernetipv4Pointtopoint}
    cdtethernettemplateentry.EntityData.Leafs["cdtEthernetMacAddr"] = types.YLeaf{"Cdtethernetmacaddr", cdtethernettemplateentry.Cdtethernetmacaddr}
    return &(cdtethernettemplateentry.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtsrvtemplatetable *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable) GetEntityData() *types.CommonEntityData {
    cdtsrvtemplatetable.EntityData.YFilter = cdtsrvtemplatetable.YFilter
    cdtsrvtemplatetable.EntityData.YangName = "cdtSrvTemplateTable"
    cdtsrvtemplatetable.EntityData.BundleName = "cisco_ios_xe"
    cdtsrvtemplatetable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtsrvtemplatetable.EntityData.SegmentPath = "cdtSrvTemplateTable"
    cdtsrvtemplatetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtsrvtemplatetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtsrvtemplatetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtsrvtemplatetable.EntityData.Children = make(map[string]types.YChild)
    cdtsrvtemplatetable.EntityData.Children["cdtSrvTemplateEntry"] = types.YChild{"Cdtsrvtemplateentry", nil}
    for i := range cdtsrvtemplatetable.Cdtsrvtemplateentry {
        cdtsrvtemplatetable.EntityData.Children[types.GetSegmentPath(&cdtsrvtemplatetable.Cdtsrvtemplateentry[i])] = types.YChild{"Cdtsrvtemplateentry", &cdtsrvtemplatetable.Cdtsrvtemplateentry[i]}
    }
    cdtsrvtemplatetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cdtsrvtemplatetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cdtsrvtemplateentry *CISCODYNAMICTEMPLATEMIB_Cdtsrvtemplatetable_Cdtsrvtemplateentry) GetEntityData() *types.CommonEntityData {
    cdtsrvtemplateentry.EntityData.YFilter = cdtsrvtemplateentry.YFilter
    cdtsrvtemplateentry.EntityData.YangName = "cdtSrvTemplateEntry"
    cdtsrvtemplateentry.EntityData.BundleName = "cisco_ios_xe"
    cdtsrvtemplateentry.EntityData.ParentYangName = "cdtSrvTemplateTable"
    cdtsrvtemplateentry.EntityData.SegmentPath = "cdtSrvTemplateEntry" + "[cdtTemplateName='" + fmt.Sprintf("%v", cdtsrvtemplateentry.Cdttemplatename) + "']"
    cdtsrvtemplateentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtsrvtemplateentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtsrvtemplateentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtsrvtemplateentry.EntityData.Children = make(map[string]types.YChild)
    cdtsrvtemplateentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cdtsrvtemplateentry.EntityData.Leafs["cdtTemplateName"] = types.YLeaf{"Cdttemplatename", cdtsrvtemplateentry.Cdttemplatename}
    cdtsrvtemplateentry.EntityData.Leafs["cdtSrvValid"] = types.YLeaf{"Cdtsrvvalid", cdtsrvtemplateentry.Cdtsrvvalid}
    cdtsrvtemplateentry.EntityData.Leafs["cdtSrvNetworkSrv"] = types.YLeaf{"Cdtsrvnetworksrv", cdtsrvtemplateentry.Cdtsrvnetworksrv}
    cdtsrvtemplateentry.EntityData.Leafs["cdtSrvVpdnGroup"] = types.YLeaf{"Cdtsrvvpdngroup", cdtsrvtemplateentry.Cdtsrvvpdngroup}
    cdtsrvtemplateentry.EntityData.Leafs["cdtSrvSgSrvGroup"] = types.YLeaf{"Cdtsrvsgsrvgroup", cdtsrvtemplateentry.Cdtsrvsgsrvgroup}
    cdtsrvtemplateentry.EntityData.Leafs["cdtSrvSgSrvType"] = types.YLeaf{"Cdtsrvsgsrvtype", cdtsrvtemplateentry.Cdtsrvsgsrvtype}
    cdtsrvtemplateentry.EntityData.Leafs["cdtSrvMulticast"] = types.YLeaf{"Cdtsrvmulticast", cdtsrvtemplateentry.Cdtsrvmulticast}
    return &(cdtsrvtemplateentry.EntityData)
}

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

