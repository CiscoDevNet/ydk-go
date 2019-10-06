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
    CdtTemplateTable CISCODYNAMICTEMPLATEMIB_CdtTemplateTable

    // This table contains a list of targets associated with one or more dynamic
    // templates.
    CdtTemplateTargetTable CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable

    // This table contains a list of templates associated with each target.  This
    // table has an expansion dependent relationship on the
    // cdtTemplateTargetTable, containing zero or more rows for each target.
    CdtTemplateAssociationTable CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable

    // This table contains a list of targets using each dynamic template.  This
    // table has an expansion dependent relationship on the cdtTemplateTable,
    // containing zero or more rows for each dynamic template.
    CdtTemplateUsageTable CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable

    // This table contains attributes relating to all dynamic templates.  Observe
    // that the type of dynamic templates containing these attributes may change
    // with the addition of new dynamic template types.  This table has a
    // sparse-dependent relationship on the cdtTemplateTable, containing a row for
    // each dynamic template having a cdtTemplateType of one of the following
    // values:      'derived'     'ppp'     'ethernet'     'ipSubscriber'    
    // 'service'.
    CdtTemplateCommonTable CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable

    // This table contains attributes relating to interface configuration.  This
    // table has a sparse-dependent relationship on the cdtTemplateTable,
    // containing a row for each dynamic template having a cdtTemplateType of one
    // of the following values:      'derived'     'ppp'     'ethernet'    
    // 'ipSubscriber'.
    CdtIfTemplateTable CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable

    // This table contains attributes relating to PPP connection configuration. 
    // This table has a sparse-dependent relationship on the cdtTemplateTable,
    // containing a row for each dynamic template having a cdtTemplateType of one
    // of the following values:      'derived'     'ppp'.
    CdtPppTemplateTable CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable

    // This table contains a prioritized list of named pools for each PPP
    // template.  A list corresponding to a PPP template specifies the pools the
    // system searches in an attempt to assign an IP address to the peer of a
    // target PPP connection. The system searches the pools in the order of their
    // priority.  This table has an expansion dependent relationship on the
    // cdtPppTemplateTable, containing zero or more rows for each PPP template.
    CdtPppPeerIpAddrPoolTable CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable

    // This table contains attributes relating to dynamic interfaces initiated on
    // Ethernet virtual interfaces (e.g., EoMPLS) or automatically created VLANs. 
    // This table has a sparse-dependent relationship on the cdtTemplateTable,
    // containing a row for each dynamic template having a cdtTemplateType of one
    // of the following values:      'derived'     'ethernet'.
    CdtEthernetTemplateTable CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable

    // This table contains attributes relating to a service.  This table has a
    // sparse-dependent relationship on the cdtTemplateTable, containing a row for
    // each dynamic template having a cdtTemplateType of one of the following
    // values:      'derived'     'service'.
    CdtSrvTemplateTable CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable
}

func (cISCODYNAMICTEMPLATEMIB *CISCODYNAMICTEMPLATEMIB) GetEntityData() *types.CommonEntityData {
    cISCODYNAMICTEMPLATEMIB.EntityData.YFilter = cISCODYNAMICTEMPLATEMIB.YFilter
    cISCODYNAMICTEMPLATEMIB.EntityData.YangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cISCODYNAMICTEMPLATEMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCODYNAMICTEMPLATEMIB.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cISCODYNAMICTEMPLATEMIB.EntityData.SegmentPath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB"
    cISCODYNAMICTEMPLATEMIB.EntityData.AbsolutePath = cISCODYNAMICTEMPLATEMIB.EntityData.SegmentPath
    cISCODYNAMICTEMPLATEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCODYNAMICTEMPLATEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCODYNAMICTEMPLATEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCODYNAMICTEMPLATEMIB.EntityData.Children = types.NewOrderedMap()
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtTemplateTable", types.YChild{"CdtTemplateTable", &cISCODYNAMICTEMPLATEMIB.CdtTemplateTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtTemplateTargetTable", types.YChild{"CdtTemplateTargetTable", &cISCODYNAMICTEMPLATEMIB.CdtTemplateTargetTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtTemplateAssociationTable", types.YChild{"CdtTemplateAssociationTable", &cISCODYNAMICTEMPLATEMIB.CdtTemplateAssociationTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtTemplateUsageTable", types.YChild{"CdtTemplateUsageTable", &cISCODYNAMICTEMPLATEMIB.CdtTemplateUsageTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtTemplateCommonTable", types.YChild{"CdtTemplateCommonTable", &cISCODYNAMICTEMPLATEMIB.CdtTemplateCommonTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtIfTemplateTable", types.YChild{"CdtIfTemplateTable", &cISCODYNAMICTEMPLATEMIB.CdtIfTemplateTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtPppTemplateTable", types.YChild{"CdtPppTemplateTable", &cISCODYNAMICTEMPLATEMIB.CdtPppTemplateTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtPppPeerIpAddrPoolTable", types.YChild{"CdtPppPeerIpAddrPoolTable", &cISCODYNAMICTEMPLATEMIB.CdtPppPeerIpAddrPoolTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtEthernetTemplateTable", types.YChild{"CdtEthernetTemplateTable", &cISCODYNAMICTEMPLATEMIB.CdtEthernetTemplateTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Children.Append("cdtSrvTemplateTable", types.YChild{"CdtSrvTemplateTable", &cISCODYNAMICTEMPLATEMIB.CdtSrvTemplateTable})
    cISCODYNAMICTEMPLATEMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCODYNAMICTEMPLATEMIB.EntityData.YListKeys = []string {}

    return &(cISCODYNAMICTEMPLATEMIB.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateTable
// This table lists the dynamic templates maintained by the
// system, including those that have been locally-configured on the
// system and those pushed to the system by external policy
// servers.
type CISCODYNAMICTEMPLATEMIB_CdtTemplateTable struct {
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
    // CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry.
    CdtTemplateEntry []*CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry
}

func (cdtTemplateTable *CISCODYNAMICTEMPLATEMIB_CdtTemplateTable) GetEntityData() *types.CommonEntityData {
    cdtTemplateTable.EntityData.YFilter = cdtTemplateTable.YFilter
    cdtTemplateTable.EntityData.YangName = "cdtTemplateTable"
    cdtTemplateTable.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtTemplateTable.EntityData.SegmentPath = "cdtTemplateTable"
    cdtTemplateTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtTemplateTable.EntityData.SegmentPath
    cdtTemplateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateTable.EntityData.Children = types.NewOrderedMap()
    cdtTemplateTable.EntityData.Children.Append("cdtTemplateEntry", types.YChild{"CdtTemplateEntry", nil})
    for i := range cdtTemplateTable.CdtTemplateEntry {
        cdtTemplateTable.EntityData.Children.Append(types.GetSegmentPath(cdtTemplateTable.CdtTemplateEntry[i]), types.YChild{"CdtTemplateEntry", cdtTemplateTable.CdtTemplateEntry[i]})
    }
    cdtTemplateTable.EntityData.Leafs = types.NewOrderedMap()

    cdtTemplateTable.EntityData.YListKeys = []string {}

    return &(cdtTemplateTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object indicates a string-value that uniquely
    // identifies the dynamic template.  If the corresponding instance of
    // cdtTemplateSrc is not 'local', then the system automatically generates the
    // name identifying the dynamic template. The type is string with length:
    // 1..64.
    CdtTemplateName interface{}

    // This object specifies the status of the dynamic template.  The following
    // columns must be valid before activating a dynamic template:      -
    // cdtTemplateStorage     - cdtTemplateType  However, these objects specify a
    // default value.  Thus, it is possible to use create-and-go semantics without
    // setting any additional columns.  An implementation must allow the EMS/NMS
    // to modify any column when this column is 'active', including columns
    // defined in tables that have a one-to-one or sparse dependent relationship
    // on this table. The type is RowStatus.
    CdtTemplateStatus interface{}

    // This object specifies what happens to the dynamic template upon restart. 
    // If the corresponding instance of cdtTemplateSrc is not 'local', then this
    // column must be 'volatile'. The type is StorageType.
    CdtTemplateStorage interface{}

    // This object indicates the types of dynamic template. The type is
    // DynamicTemplateType.
    CdtTemplateType interface{}

    // This object specifies the source of the dynamic template:  'other'     The
    // implementation of the MIB module does not recognize the     source of the
    // dynamic template.  'derived'     The system created the set of attributes
    // from one or     more dynamic templates.  'local'     The dynamic template
    // was locally configured through a     management entity, such as the local
    // console or a SNMP     entity.  'aaaUserProfile'     The dynamic template
    // originated from a user profile     pushed from an external policy server. 
    // 'aaaServiceProfile'     The dynamic template originated from a service
    // profile     pushed from an external policy server. The type is
    // CdtTemplateSrc.
    CdtTemplateSrc interface{}

    // This object specifies the number of targets using a dynamic template. The
    // type is interface{} with range: 0..4294967295.
    CdtTemplateUsageCount interface{}
}

func (cdtTemplateEntry *CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry) GetEntityData() *types.CommonEntityData {
    cdtTemplateEntry.EntityData.YFilter = cdtTemplateEntry.YFilter
    cdtTemplateEntry.EntityData.YangName = "cdtTemplateEntry"
    cdtTemplateEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateEntry.EntityData.ParentYangName = "cdtTemplateTable"
    cdtTemplateEntry.EntityData.SegmentPath = "cdtTemplateEntry" + types.AddKeyToken(cdtTemplateEntry.CdtTemplateName, "cdtTemplateName")
    cdtTemplateEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtTemplateTable/" + cdtTemplateEntry.EntityData.SegmentPath
    cdtTemplateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateEntry.EntityData.Children = types.NewOrderedMap()
    cdtTemplateEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtTemplateEntry.EntityData.Leafs.Append("cdtTemplateName", types.YLeaf{"CdtTemplateName", cdtTemplateEntry.CdtTemplateName})
    cdtTemplateEntry.EntityData.Leafs.Append("cdtTemplateStatus", types.YLeaf{"CdtTemplateStatus", cdtTemplateEntry.CdtTemplateStatus})
    cdtTemplateEntry.EntityData.Leafs.Append("cdtTemplateStorage", types.YLeaf{"CdtTemplateStorage", cdtTemplateEntry.CdtTemplateStorage})
    cdtTemplateEntry.EntityData.Leafs.Append("cdtTemplateType", types.YLeaf{"CdtTemplateType", cdtTemplateEntry.CdtTemplateType})
    cdtTemplateEntry.EntityData.Leafs.Append("cdtTemplateSrc", types.YLeaf{"CdtTemplateSrc", cdtTemplateEntry.CdtTemplateSrc})
    cdtTemplateEntry.EntityData.Leafs.Append("cdtTemplateUsageCount", types.YLeaf{"CdtTemplateUsageCount", cdtTemplateEntry.CdtTemplateUsageCount})

    cdtTemplateEntry.EntityData.YListKeys = []string {"CdtTemplateName"}

    return &(cdtTemplateEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc represents     pushed from an external policy server.
type CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc string

const (
    CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc_other CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc = "other"

    CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc_derived CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc = "derived"

    CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc_local CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc = "local"

    CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc_aaaUserProfile CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc = "aaaUserProfile"

    CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc_aaaServiceProfile CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateSrc = "aaaServiceProfile"
)

// CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable
// This table contains a list of targets associated with
// one or more dynamic templates.
type CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry describes a target associated with one or more dynamic templates. 
    // The system automatically creates an entry when it associates a dynamic
    // template to a target.  Likewise, the system automatically destroys an entry
    // when a target no longer has any associated dynamic templates. The type is
    // slice of
    // CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable_CdtTemplateTargetEntry.
    CdtTemplateTargetEntry []*CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable_CdtTemplateTargetEntry
}

func (cdtTemplateTargetTable *CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable) GetEntityData() *types.CommonEntityData {
    cdtTemplateTargetTable.EntityData.YFilter = cdtTemplateTargetTable.YFilter
    cdtTemplateTargetTable.EntityData.YangName = "cdtTemplateTargetTable"
    cdtTemplateTargetTable.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateTargetTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtTemplateTargetTable.EntityData.SegmentPath = "cdtTemplateTargetTable"
    cdtTemplateTargetTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtTemplateTargetTable.EntityData.SegmentPath
    cdtTemplateTargetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateTargetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateTargetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateTargetTable.EntityData.Children = types.NewOrderedMap()
    cdtTemplateTargetTable.EntityData.Children.Append("cdtTemplateTargetEntry", types.YChild{"CdtTemplateTargetEntry", nil})
    for i := range cdtTemplateTargetTable.CdtTemplateTargetEntry {
        cdtTemplateTargetTable.EntityData.Children.Append(types.GetSegmentPath(cdtTemplateTargetTable.CdtTemplateTargetEntry[i]), types.YChild{"CdtTemplateTargetEntry", cdtTemplateTargetTable.CdtTemplateTargetEntry[i]})
    }
    cdtTemplateTargetTable.EntityData.Leafs = types.NewOrderedMap()

    cdtTemplateTargetTable.EntityData.YListKeys = []string {}

    return &(cdtTemplateTargetTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable_CdtTemplateTargetEntry
// An entry describes a target associated with one or more
// dynamic templates.
// 
// The system automatically creates an entry when it associates a
// dynamic template to a target.  Likewise, the system
// automatically destroys an entry when a target no longer has any
// associated dynamic templates.
type CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable_CdtTemplateTargetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object indicates the type of target. The type
    // is DynamicTemplateTargetType.
    CdtTemplateTargetType interface{}

    // This attribute is a key. This object uniquely identifies the target within
    // the scope of its type. The type is string with length: 1..20.
    CdtTemplateTargetId interface{}

    // This object specifies the status of the dynamic template target.  The
    // following columns must be valid before activating a subscriber access
    // profile:      - cdtTemplateTargetStorage  However, these objects specify a
    // default value.  Thus, it is possible to use create-and-go semantics without
    // setting any additional columns.  An implementation must allow the EMS/NMS
    // to modify any column when this column is 'active', including columns
    // defined in tables that have a one-to-one or sparse dependent relationship
    // on this table. The type is RowStatus.
    CdtTemplateTargetStatus interface{}

    // This object specifies what happens to the dynamic template target upon
    // restart. The type is StorageType.
    CdtTemplateTargetStorage interface{}
}

func (cdtTemplateTargetEntry *CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable_CdtTemplateTargetEntry) GetEntityData() *types.CommonEntityData {
    cdtTemplateTargetEntry.EntityData.YFilter = cdtTemplateTargetEntry.YFilter
    cdtTemplateTargetEntry.EntityData.YangName = "cdtTemplateTargetEntry"
    cdtTemplateTargetEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateTargetEntry.EntityData.ParentYangName = "cdtTemplateTargetTable"
    cdtTemplateTargetEntry.EntityData.SegmentPath = "cdtTemplateTargetEntry" + types.AddKeyToken(cdtTemplateTargetEntry.CdtTemplateTargetType, "cdtTemplateTargetType") + types.AddKeyToken(cdtTemplateTargetEntry.CdtTemplateTargetId, "cdtTemplateTargetId")
    cdtTemplateTargetEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtTemplateTargetTable/" + cdtTemplateTargetEntry.EntityData.SegmentPath
    cdtTemplateTargetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateTargetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateTargetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateTargetEntry.EntityData.Children = types.NewOrderedMap()
    cdtTemplateTargetEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtTemplateTargetEntry.EntityData.Leafs.Append("cdtTemplateTargetType", types.YLeaf{"CdtTemplateTargetType", cdtTemplateTargetEntry.CdtTemplateTargetType})
    cdtTemplateTargetEntry.EntityData.Leafs.Append("cdtTemplateTargetId", types.YLeaf{"CdtTemplateTargetId", cdtTemplateTargetEntry.CdtTemplateTargetId})
    cdtTemplateTargetEntry.EntityData.Leafs.Append("cdtTemplateTargetStatus", types.YLeaf{"CdtTemplateTargetStatus", cdtTemplateTargetEntry.CdtTemplateTargetStatus})
    cdtTemplateTargetEntry.EntityData.Leafs.Append("cdtTemplateTargetStorage", types.YLeaf{"CdtTemplateTargetStorage", cdtTemplateTargetEntry.CdtTemplateTargetStorage})

    cdtTemplateTargetEntry.EntityData.YListKeys = []string {"CdtTemplateTargetType", "CdtTemplateTargetId"}

    return &(cdtTemplateTargetEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable
// This table contains a list of templates associated with each
// target.
// 
// This table has an expansion dependent relationship on the
// cdtTemplateTargetTable, containing zero or more rows for each
// target.
type CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry indicates an association of a dynamic template with a target.  The
    // system automatically creates an entry when it associates a dynamic template
    // to a target.  The system also creates an entry when it derives the
    // configuration that it applies to a target.  The system automatically
    // destroys an entry under the following circumstances:  1)  The target
    // indicated by the entry no longer exists.  2)  The association between the
    // target and the dynamic template     no longer exists. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable_CdtTemplateAssociationEntry.
    CdtTemplateAssociationEntry []*CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable_CdtTemplateAssociationEntry
}

func (cdtTemplateAssociationTable *CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable) GetEntityData() *types.CommonEntityData {
    cdtTemplateAssociationTable.EntityData.YFilter = cdtTemplateAssociationTable.YFilter
    cdtTemplateAssociationTable.EntityData.YangName = "cdtTemplateAssociationTable"
    cdtTemplateAssociationTable.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateAssociationTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtTemplateAssociationTable.EntityData.SegmentPath = "cdtTemplateAssociationTable"
    cdtTemplateAssociationTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtTemplateAssociationTable.EntityData.SegmentPath
    cdtTemplateAssociationTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateAssociationTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateAssociationTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateAssociationTable.EntityData.Children = types.NewOrderedMap()
    cdtTemplateAssociationTable.EntityData.Children.Append("cdtTemplateAssociationEntry", types.YChild{"CdtTemplateAssociationEntry", nil})
    for i := range cdtTemplateAssociationTable.CdtTemplateAssociationEntry {
        cdtTemplateAssociationTable.EntityData.Children.Append(types.GetSegmentPath(cdtTemplateAssociationTable.CdtTemplateAssociationEntry[i]), types.YChild{"CdtTemplateAssociationEntry", cdtTemplateAssociationTable.CdtTemplateAssociationEntry[i]})
    }
    cdtTemplateAssociationTable.EntityData.Leafs = types.NewOrderedMap()

    cdtTemplateAssociationTable.EntityData.YListKeys = []string {}

    return &(cdtTemplateAssociationTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable_CdtTemplateAssociationEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable_CdtTemplateAssociationEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is DynamicTemplateTargetType. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable_CdtTemplateTargetEntry_CdtTemplateTargetType
    CdtTemplateTargetType interface{}

    // This attribute is a key. The type is string with length: 1..20. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTargetTable_CdtTemplateTargetEntry_CdtTemplateTargetId
    CdtTemplateTargetId interface{}

    // This attribute is a key. This object indicates the name of the template
    // associated with the target. The type is string with length: 1..64.
    CdtTemplateAssociationName interface{}

    // This object indicates the relative precedence of the associated dynamic
    // template.  Lower values have higher precedence than higher values. The type
    // is interface{} with range: 0..4294967295.
    CdtTemplateAssociationPrecedence interface{}
}

func (cdtTemplateAssociationEntry *CISCODYNAMICTEMPLATEMIB_CdtTemplateAssociationTable_CdtTemplateAssociationEntry) GetEntityData() *types.CommonEntityData {
    cdtTemplateAssociationEntry.EntityData.YFilter = cdtTemplateAssociationEntry.YFilter
    cdtTemplateAssociationEntry.EntityData.YangName = "cdtTemplateAssociationEntry"
    cdtTemplateAssociationEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateAssociationEntry.EntityData.ParentYangName = "cdtTemplateAssociationTable"
    cdtTemplateAssociationEntry.EntityData.SegmentPath = "cdtTemplateAssociationEntry" + types.AddKeyToken(cdtTemplateAssociationEntry.CdtTemplateTargetType, "cdtTemplateTargetType") + types.AddKeyToken(cdtTemplateAssociationEntry.CdtTemplateTargetId, "cdtTemplateTargetId") + types.AddKeyToken(cdtTemplateAssociationEntry.CdtTemplateAssociationName, "cdtTemplateAssociationName")
    cdtTemplateAssociationEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtTemplateAssociationTable/" + cdtTemplateAssociationEntry.EntityData.SegmentPath
    cdtTemplateAssociationEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateAssociationEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateAssociationEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateAssociationEntry.EntityData.Children = types.NewOrderedMap()
    cdtTemplateAssociationEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtTemplateAssociationEntry.EntityData.Leafs.Append("cdtTemplateTargetType", types.YLeaf{"CdtTemplateTargetType", cdtTemplateAssociationEntry.CdtTemplateTargetType})
    cdtTemplateAssociationEntry.EntityData.Leafs.Append("cdtTemplateTargetId", types.YLeaf{"CdtTemplateTargetId", cdtTemplateAssociationEntry.CdtTemplateTargetId})
    cdtTemplateAssociationEntry.EntityData.Leafs.Append("cdtTemplateAssociationName", types.YLeaf{"CdtTemplateAssociationName", cdtTemplateAssociationEntry.CdtTemplateAssociationName})
    cdtTemplateAssociationEntry.EntityData.Leafs.Append("cdtTemplateAssociationPrecedence", types.YLeaf{"CdtTemplateAssociationPrecedence", cdtTemplateAssociationEntry.CdtTemplateAssociationPrecedence})

    cdtTemplateAssociationEntry.EntityData.YListKeys = []string {"CdtTemplateTargetType", "CdtTemplateTargetId", "CdtTemplateAssociationName"}

    return &(cdtTemplateAssociationEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable
// This table contains a list of targets using each dynamic
// template.
// 
// This table has an expansion dependent relationship on the
// cdtTemplateTable, containing zero or more rows for each
// dynamic template.
type CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry indicates a target using the dynamic template.  The system
    // automatically creates an entry when it associates a dynamic template to a
    // target.  The system also creates an entry when it derives the configuration
    // that it applies to a target.  The system automatically destroys an entry
    // under the following circumstances:  1)  The target indicated by the entry
    // no longer exists.  2)  The association between the target and the dynamic
    // template     no longer exists. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable_CdtTemplateUsageEntry.
    CdtTemplateUsageEntry []*CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable_CdtTemplateUsageEntry
}

func (cdtTemplateUsageTable *CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable) GetEntityData() *types.CommonEntityData {
    cdtTemplateUsageTable.EntityData.YFilter = cdtTemplateUsageTable.YFilter
    cdtTemplateUsageTable.EntityData.YangName = "cdtTemplateUsageTable"
    cdtTemplateUsageTable.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateUsageTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtTemplateUsageTable.EntityData.SegmentPath = "cdtTemplateUsageTable"
    cdtTemplateUsageTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtTemplateUsageTable.EntityData.SegmentPath
    cdtTemplateUsageTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateUsageTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateUsageTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateUsageTable.EntityData.Children = types.NewOrderedMap()
    cdtTemplateUsageTable.EntityData.Children.Append("cdtTemplateUsageEntry", types.YChild{"CdtTemplateUsageEntry", nil})
    for i := range cdtTemplateUsageTable.CdtTemplateUsageEntry {
        cdtTemplateUsageTable.EntityData.Children.Append(types.GetSegmentPath(cdtTemplateUsageTable.CdtTemplateUsageEntry[i]), types.YChild{"CdtTemplateUsageEntry", cdtTemplateUsageTable.CdtTemplateUsageEntry[i]})
    }
    cdtTemplateUsageTable.EntityData.Leafs = types.NewOrderedMap()

    cdtTemplateUsageTable.EntityData.YListKeys = []string {}

    return &(cdtTemplateUsageTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable_CdtTemplateUsageEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable_CdtTemplateUsageEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateName
    CdtTemplateName interface{}

    // This attribute is a key. This object indicates the type of target using the
    // dynamic template. The type is DynamicTemplateTargetType.
    CdtTemplateUsageTargetType interface{}

    // This attribute is a key. This object indicates the name of the target using
    // the dynamic template. The type is string with length: 1..20.
    CdtTemplateUsageTargetId interface{}
}

func (cdtTemplateUsageEntry *CISCODYNAMICTEMPLATEMIB_CdtTemplateUsageTable_CdtTemplateUsageEntry) GetEntityData() *types.CommonEntityData {
    cdtTemplateUsageEntry.EntityData.YFilter = cdtTemplateUsageEntry.YFilter
    cdtTemplateUsageEntry.EntityData.YangName = "cdtTemplateUsageEntry"
    cdtTemplateUsageEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateUsageEntry.EntityData.ParentYangName = "cdtTemplateUsageTable"
    cdtTemplateUsageEntry.EntityData.SegmentPath = "cdtTemplateUsageEntry" + types.AddKeyToken(cdtTemplateUsageEntry.CdtTemplateName, "cdtTemplateName") + types.AddKeyToken(cdtTemplateUsageEntry.CdtTemplateUsageTargetType, "cdtTemplateUsageTargetType") + types.AddKeyToken(cdtTemplateUsageEntry.CdtTemplateUsageTargetId, "cdtTemplateUsageTargetId")
    cdtTemplateUsageEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtTemplateUsageTable/" + cdtTemplateUsageEntry.EntityData.SegmentPath
    cdtTemplateUsageEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateUsageEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateUsageEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateUsageEntry.EntityData.Children = types.NewOrderedMap()
    cdtTemplateUsageEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtTemplateUsageEntry.EntityData.Leafs.Append("cdtTemplateName", types.YLeaf{"CdtTemplateName", cdtTemplateUsageEntry.CdtTemplateName})
    cdtTemplateUsageEntry.EntityData.Leafs.Append("cdtTemplateUsageTargetType", types.YLeaf{"CdtTemplateUsageTargetType", cdtTemplateUsageEntry.CdtTemplateUsageTargetType})
    cdtTemplateUsageEntry.EntityData.Leafs.Append("cdtTemplateUsageTargetId", types.YLeaf{"CdtTemplateUsageTargetId", cdtTemplateUsageEntry.CdtTemplateUsageTargetId})

    cdtTemplateUsageEntry.EntityData.YListKeys = []string {"CdtTemplateName", "CdtTemplateUsageTargetType", "CdtTemplateUsageTargetId"}

    return &(cdtTemplateUsageEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable
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
type CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing attributes relating to any target.  The system
    // automatically creates an entry when the system or the EMS/NMS creates a row
    // in the cdtTemplateTable with a cdtTemplateType of on the following values: 
    // 'derived'     'ppp'     'ethernet'     'ipSubscriber'     'service' 
    // Likewise, the system automatically destroys an entry when the system or the
    // EMS/NMS destroys the corresponding row in the cdtTemplateTable. The type is
    // slice of
    // CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable_CdtTemplateCommonEntry.
    CdtTemplateCommonEntry []*CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable_CdtTemplateCommonEntry
}

func (cdtTemplateCommonTable *CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable) GetEntityData() *types.CommonEntityData {
    cdtTemplateCommonTable.EntityData.YFilter = cdtTemplateCommonTable.YFilter
    cdtTemplateCommonTable.EntityData.YangName = "cdtTemplateCommonTable"
    cdtTemplateCommonTable.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateCommonTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtTemplateCommonTable.EntityData.SegmentPath = "cdtTemplateCommonTable"
    cdtTemplateCommonTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtTemplateCommonTable.EntityData.SegmentPath
    cdtTemplateCommonTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateCommonTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateCommonTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateCommonTable.EntityData.Children = types.NewOrderedMap()
    cdtTemplateCommonTable.EntityData.Children.Append("cdtTemplateCommonEntry", types.YChild{"CdtTemplateCommonEntry", nil})
    for i := range cdtTemplateCommonTable.CdtTemplateCommonEntry {
        cdtTemplateCommonTable.EntityData.Children.Append(types.GetSegmentPath(cdtTemplateCommonTable.CdtTemplateCommonEntry[i]), types.YChild{"CdtTemplateCommonEntry", cdtTemplateCommonTable.CdtTemplateCommonEntry[i]})
    }
    cdtTemplateCommonTable.EntityData.Leafs = types.NewOrderedMap()

    cdtTemplateCommonTable.EntityData.YListKeys = []string {}

    return &(cdtTemplateCommonTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable_CdtTemplateCommonEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable_CdtTemplateCommonEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateName
    CdtTemplateName interface{}

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
    CdtCommonValid interface{}

    // This object specifies a human-readable description for the dynamic
    // template.  This column is valid only if the 'descr' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is string with
    // length: 0..255.
    CdtCommonDescr interface{}

    // This object specifies the interval that the system sends keepalive messages
    // to a target.  This column is valid only if the 'keepaliveInterval' bit of
    // the corresponding instance of cdtCommonValid is '1'. The type is
    // interface{} with range: 1..4294967295. Units are seconds.
    CdtCommonKeepaliveInt interface{}

    // This object specifies the number of times the system will resend a
    // keepalive message without a response before it transitions a target to an
    // operationally down state.  This column is valid only if the
    // 'keepaliveRetries' bit of the corresponding instance of cdtCommonValid is
    // '1'. The type is interface{} with range: 1..255. Units are retries.
    CdtCommonKeepaliveRetries interface{}

    // This object specifies the name of the VRF with which a target has an
    // association.  This column is valid only if the 'vrf' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is string with
    // length: 0..32.
    CdtCommonVrf interface{}

    // This object specifies the name of the IP address pool the system will use
    // to assign an IP address to a peer of a target.  This column is valid only
    // if the 'addrPool' bit of the corresponding instance of cdtCommonValid is
    // '1'. The type is string with length: 0..255.
    CdtCommonAddrPool interface{}

    // This object specifies the name (or number) of the IPv4 ACL applied to a
    // target.  This column is valid only if the 'ipv4AccessGroup' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is string with
    // length: 0..255.
    CdtCommonIpv4AccessGroup interface{}

    // This object specifies whether a target generates ICMPv4 unreachable
    // messages.  This column is valid only if the 'ipv4Unreachables' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is bool.
    CdtCommonIpv4Unreachables interface{}

    // This object specifies the name (or number) of the IPv4 ACL applied to a
    // target.  This column is valid only if the 'ipv6AccessGroup' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is string with
    // length: 0..255.
    CdtCommonIpv6AccessGroup interface{}

    // This object specifies whether a target generates ICMPv6 unreachable
    // messages.  This column is valid only if the 'ipv6Unreachables' bit of the
    // corresponding instance of cdtCommonValid is '1'. The type is bool.
    CdtCommonIpv6Unreachables interface{}

    // This object specifies the name of the subscriber control policy applied to
    // a target.  The system should assume that the cbpPolicyMapType (defined by
    // the CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtControlSubscriber
    // (defined by the CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the
    // 'srvSubControl' bit of the corresponding instance of cdtCommonValid is '1'.
    // The type is string.
    CdtCommonSrvSubControl interface{}

    // This object specifies the name of the traffic redirect policy applied to a
    // target.  The system should assume that the cbpPolicyMapType (defined by the
    // CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtTrafficRedirect (defined by
    // the CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the
    // 'srvRedirect' bit of the corresponding instance of cdtCommonValid is '1'.
    // The type is string.
    CdtCommonSrvRedirect interface{}

    // This object specifies the name of the traffic accounting policy applied to
    // a target.  The system should assume that the cbpPolicyMapType (defined by
    // the CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtTrafficAccounting
    // (defined by the CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the
    // 'srvAcct' bit of the corresponding instance of cdtCommonValid is '1'. The
    // type is string.
    CdtCommonSrvAcct interface{}

    // This object specifies the name of the traffic QoS policy applied to a
    // target.  The system should assume that the cbpPolicyMapType (defined by the
    // CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtQos (defined by the
    // CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the 'srvQos' bit of
    // the corresponding instance of cdtCommonValid is '1'. The type is string.
    CdtCommonSrvQos interface{}

    // This object specifies the name of the NetFlow policy applied to a target. 
    // The system should assume that the cbpPolicyMapType (defined by the
    // CISCO-CBP-BASE-CFG-MIB) of the policy is cbpPmtNetflow (defined by the
    // CISCO-CBP-TYPE-OID-MIB).  This column is valid only if the 'srvNetflow' bit
    // of the corresponding instance of cdtCommonValid is '1'. The type is string.
    CdtCommonSrvNetflow interface{}
}

func (cdtTemplateCommonEntry *CISCODYNAMICTEMPLATEMIB_CdtTemplateCommonTable_CdtTemplateCommonEntry) GetEntityData() *types.CommonEntityData {
    cdtTemplateCommonEntry.EntityData.YFilter = cdtTemplateCommonEntry.YFilter
    cdtTemplateCommonEntry.EntityData.YangName = "cdtTemplateCommonEntry"
    cdtTemplateCommonEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtTemplateCommonEntry.EntityData.ParentYangName = "cdtTemplateCommonTable"
    cdtTemplateCommonEntry.EntityData.SegmentPath = "cdtTemplateCommonEntry" + types.AddKeyToken(cdtTemplateCommonEntry.CdtTemplateName, "cdtTemplateName")
    cdtTemplateCommonEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtTemplateCommonTable/" + cdtTemplateCommonEntry.EntityData.SegmentPath
    cdtTemplateCommonEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtTemplateCommonEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtTemplateCommonEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtTemplateCommonEntry.EntityData.Children = types.NewOrderedMap()
    cdtTemplateCommonEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtTemplateName", types.YLeaf{"CdtTemplateName", cdtTemplateCommonEntry.CdtTemplateName})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonValid", types.YLeaf{"CdtCommonValid", cdtTemplateCommonEntry.CdtCommonValid})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonDescr", types.YLeaf{"CdtCommonDescr", cdtTemplateCommonEntry.CdtCommonDescr})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonKeepaliveInt", types.YLeaf{"CdtCommonKeepaliveInt", cdtTemplateCommonEntry.CdtCommonKeepaliveInt})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonKeepaliveRetries", types.YLeaf{"CdtCommonKeepaliveRetries", cdtTemplateCommonEntry.CdtCommonKeepaliveRetries})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonVrf", types.YLeaf{"CdtCommonVrf", cdtTemplateCommonEntry.CdtCommonVrf})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonAddrPool", types.YLeaf{"CdtCommonAddrPool", cdtTemplateCommonEntry.CdtCommonAddrPool})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonIpv4AccessGroup", types.YLeaf{"CdtCommonIpv4AccessGroup", cdtTemplateCommonEntry.CdtCommonIpv4AccessGroup})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonIpv4Unreachables", types.YLeaf{"CdtCommonIpv4Unreachables", cdtTemplateCommonEntry.CdtCommonIpv4Unreachables})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonIpv6AccessGroup", types.YLeaf{"CdtCommonIpv6AccessGroup", cdtTemplateCommonEntry.CdtCommonIpv6AccessGroup})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonIpv6Unreachables", types.YLeaf{"CdtCommonIpv6Unreachables", cdtTemplateCommonEntry.CdtCommonIpv6Unreachables})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonSrvSubControl", types.YLeaf{"CdtCommonSrvSubControl", cdtTemplateCommonEntry.CdtCommonSrvSubControl})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonSrvRedirect", types.YLeaf{"CdtCommonSrvRedirect", cdtTemplateCommonEntry.CdtCommonSrvRedirect})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonSrvAcct", types.YLeaf{"CdtCommonSrvAcct", cdtTemplateCommonEntry.CdtCommonSrvAcct})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonSrvQos", types.YLeaf{"CdtCommonSrvQos", cdtTemplateCommonEntry.CdtCommonSrvQos})
    cdtTemplateCommonEntry.EntityData.Leafs.Append("cdtCommonSrvNetflow", types.YLeaf{"CdtCommonSrvNetflow", cdtTemplateCommonEntry.CdtCommonSrvNetflow})

    cdtTemplateCommonEntry.EntityData.YListKeys = []string {"CdtTemplateName"}

    return &(cdtTemplateCommonEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable
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
type CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing attributes relating to interface configuration.  The
    // system automatically creates an entry when the system or the EMS/NMS
    // creates a row in the cdtTemplateTable with a cdtTemplateType of one of the
    // following values:      'derived'     'ppp'     'ethernet'    
    // 'ipSubscriber'  Likewise, the system automatically destroys an entry when
    // the system or the EMS/NMS destroys the corresponding row in the
    // cdtTemplateTable. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry.
    CdtIfTemplateEntry []*CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry
}

func (cdtIfTemplateTable *CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable) GetEntityData() *types.CommonEntityData {
    cdtIfTemplateTable.EntityData.YFilter = cdtIfTemplateTable.YFilter
    cdtIfTemplateTable.EntityData.YangName = "cdtIfTemplateTable"
    cdtIfTemplateTable.EntityData.BundleName = "cisco_ios_xe"
    cdtIfTemplateTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtIfTemplateTable.EntityData.SegmentPath = "cdtIfTemplateTable"
    cdtIfTemplateTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtIfTemplateTable.EntityData.SegmentPath
    cdtIfTemplateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtIfTemplateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtIfTemplateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtIfTemplateTable.EntityData.Children = types.NewOrderedMap()
    cdtIfTemplateTable.EntityData.Children.Append("cdtIfTemplateEntry", types.YChild{"CdtIfTemplateEntry", nil})
    for i := range cdtIfTemplateTable.CdtIfTemplateEntry {
        cdtIfTemplateTable.EntityData.Children.Append(types.GetSegmentPath(cdtIfTemplateTable.CdtIfTemplateEntry[i]), types.YChild{"CdtIfTemplateEntry", cdtIfTemplateTable.CdtIfTemplateEntry[i]})
    }
    cdtIfTemplateTable.EntityData.Leafs = types.NewOrderedMap()

    cdtIfTemplateTable.EntityData.YListKeys = []string {}

    return &(cdtIfTemplateTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateName
    CdtTemplateName interface{}

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
    CdtIfValid interface{}

    // This object specifies the Maximum Transfer Unit (MTU) size for all packets
    // sent on the target interface.  The value '0' cannot be written to an
    // instance of this object. However, it serves as a convenient value when the
    // column is not valid.  This column is valid only if the 'mtu' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 0..0 | 64..65535. Units are octets.
    CdtIfMtu interface{}

    // This object specifies whether the target interface participates in the
    // Cisco Discovery Protocol (CDP).  This column is valid only if the
    // 'cdpEnable' bit of the corresponding instance of cdtIfValid is '1'. The
    // type is bool.
    CdtIfCdpEnable interface{}

    // This object specifies the name of the flow monitor associated with the
    // target interface.  This column is valid only if the 'flowMonitor' bit of
    // the corresponding instance of cdtIfValid is '1'. The type is string with
    // length: 0..255.
    CdtIfFlowMonitor interface{}

    // This object specifies the interface of the source address that the target
    // interface uses when originating IPv4 packets.  The value '0' cannot be
    // written to an instance of this object. However, it serves as a convenient
    // value when the column is not valid (e.g., immediately following the
    // creation of an instance of the object).  This column is valid only if the
    // 'ipv4Unnumbered' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is interface{} with range: 0..2147483647.
    CdtIfIpv4Unnumbered interface{}

    // This object specifies whether the target interface allows IPv4 subscriber
    // sessions.  This column is valid only if the 'ipv4SubEnable' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is bool.
    CdtIfIpv4SubEnable interface{}

    // This object specifies the Maximum Transfer Unit (MTU) size for IPv4 packets
    // sent on the target interface.  The value '0' cannot be written to an
    // instance of this object. However, it serves as a convenient value when the
    // column is not valid.  This column is valid only if the 'ipv4Mtu' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 0..0 | 128..65535. Units are octets.
    CdtIfIpv4Mtu interface{}

    // This object specifies the adjustment to the Maximum Segment Size (MSS) of
    // TCP SYN packets received by the target interface contained in IPv4
    // datagrams.  The value '0' cannot be written to an instance of this object.
    // However, it serves as a convenient value when the column is not valid. 
    // This column is valid only if the 'ipv4TcpMssAdjust' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 0..0 | 500..1460. Units are octets.
    CdtIfIpv4TcpMssAdjust interface{}

    // This object specifies whether the type of unicast RPF the system performs
    // on IPv4 packets received by the target interface.  This column is valid
    // only if the 'ipv4VerifyUniRpf' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is UnicastRpfType.
    CdtIfIpv4VerifyUniRpf interface{}

    // This object specifies the name (or number) of the IPv4 ACL used to
    // determine whether the system should permit/deny packets received by the
    // target interface that fail unicast RPF verification.  This column is valid
    // only if the 'ipv4VerifyUniRpfAcl' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is string with length: 0..255.
    CdtIfIpv4VerifyUniRpfAcl interface{}

    // This object specifies the options that affect how the system performs
    // unicast RPF on IPv4 packets received by the target interface.  This column
    // is valid only if the 'ipv4VerifyUniRpfOpts' bit of the corresponding
    // instance of cdtIfValid is '1'. The type is map[string]bool.
    CdtIfIpv4VerifyUniRpfOpts interface{}

    // This object specifies whether the system processes IPv6 packets received by
    // the target interface.  This column is valid only if the 'ipv6Enable' bit of
    // the corresponding instance of cdtIfValid is '1'. The type is bool.
    CdtIfIpv6Enable interface{}

    // This object specifies whether the target interface allows IPv6 subscriber
    // sessions.  This column is valid only if the 'ipv6SubEnable' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is bool.
    CdtIfIpv6SubEnable interface{}

    // This object specifies the adjustment to the Maximum Segment Size (MSS) of
    // TCP SYN packets received by the target interface contained in IPv6
    // datagrams.  The value '0' cannot be written to an instance of this object.
    // However, it serves as a convenient value when the column is not valid. 
    // This column is valid only if the 'ipv6TcpMssAdjust' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 0..0 | 500..1460. Units are octets.
    CdtIfIpv6TcpMssAdjust interface{}

    // This object specifies whether the type of unicast RPF the system performs
    // on IPv6 packets received by the target interface.  This column is valid
    // only if the 'ipv6VerifyUniRpf' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is UnicastRpfType.
    CdtIfIpv6VerifyUniRpf interface{}

    // This object specifies the name (or number) of the IPv6 ACL used to
    // determine whether the system should permit/deny packets received by the
    // target interface that fail unicast RPF verification.  This column is valid
    // only if the 'ipv6VerifyUniRpfAcl' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is string with length: 0..255.
    CdtIfIpv6VerifyUniRpfAcl interface{}

    // This object specifies the options that affect how the system performs
    // unicast RPF on IPv6 packets received by the target interface.  This column
    // is valid only if the 'ipv6VerifyUniRpfOpts' bit of the corresponding
    // instance of cdtIfValid is '1'. The type is map[string]bool.
    CdtIfIpv6VerifyUniRpfOpts interface{}

    // This object specifies the IPv6 network number included in IPv6 router
    // advertisements sent on the target interface.  This column is valid only if
    // the 'ipv6NdPrefix' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is string.
    CdtIfIpv6NdPrefix interface{}

    // This object specifies the length of the IPv6 prefix specified by the
    // corresponding instance of cdtIpv6NdPrefix.  This column is valid only if
    // the 'ipv6NdPrefix' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is interface{} with range: 0..2040.
    CdtIfIpv6NdPrefixLength interface{}

    // This object specifies the interval that the system advertises the IPv6
    // prefix (i.e., the corresponding instance of cdtIfIpv6NdPrefix) as 'valid'
    // for IPv6 router advertisements sent on the target interface.  This column
    // is valid only if the 'ipv6NdValidLife' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is interface{} with range: 1..4294967295. Units
    // are seconds.
    CdtIfIpv6NdValidLife interface{}

    // This object specifies the interval that the system advertises the IPv6
    // prefix (i.e., the corresponding instance of cdtIfIpv6NdPrefix) as
    // 'preferred' for IPv6 router advertisements sent on the target interface. 
    // This column is valid only if the 'ipv6NdPreferredLife' bit of the
    // corresponding instance of cdtIfValid is '1'. The type is interface{} with
    // range: 1..4294967295. Units are seconds.
    CdtIfIpv6NdPreferredLife interface{}

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
    CdtIfIpv6NdOpts interface{}

    // This object specifies the number of consecutive neighbor solitication
    // messages the system sends on the target interface while performing
    // duplicate address detection on unicast IPv6 addresses on the target
    // interface.  The value '0' disables duplicate address detection on the
    // target interface.  This column is valid only if the 'ipv6NdDadAttempts' bit
    // of the corresponding instance of cdtIfValid is '1'. The type is interface{}
    // with range: 0..600.
    CdtIfIpv6NdDadAttempts interface{}

    // This object specifies the interval between neighbor solicitation
    // retransmissions on the target interface.  This column is valid only if the
    // 'ipv6NdNsIntervals' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is interface{} with range: 1000..3600000. Units are milliseconds.
    CdtIfIpv6NdNsInterval interface{}

    // This object specifies the amount of time the system considers a neighbor of
    // the target interface reachable after a reachability confirmation event has
    // occurred.  The value '0' disables neighbor reachability detection on the
    // target interface.  This column is valid only if the 'ipv6NdReachable' bit
    // of the corresponding instance of cdtIfValid is '1'. The type is interface{}
    // with range: 0..4294967295. Units are milliseconds.
    CdtIfIpv6NdReachableTime interface{}

    // This object specifies the units of time for the corresponding instances of
    // cdtIfIpv6NdRaIntervalMin and cdtIfIpv6NdRaIntervalMax.  This column is
    // valid only if the 'ipv6NdRaInterval' bit of the corresponding instance of
    // cdtIfValid is '1'. The type is CdtIfIpv6NdRaIntervalUnits.
    CdtIfIpv6NdRaIntervalUnits interface{}

    // This object specifies the maximum interval between IPv6 router
    // advertisements sent on the target interface.  The value '0' cannot be
    // written to an instance of this object. However, it serves as a convenient
    // value when the column is not valid.  This column is valid only if the
    // 'ipv6NdRaInterval' bit of the corresponding instance of cdtIfValid is '1'.
    // The type is interface{} with range: 0..4294967295.
    CdtIfIpv6NdRaIntervalMax interface{}

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
    CdtIfIpv6NdRaIntervalMin interface{}

    // This object specifies the router lifetime value in IPv6 router
    // advertisements sent on the target interface.  The value '0' specifies that
    // neighbors should not consider the router as a default router. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    CdtIfIpv6NdRaLife interface{}

    // This object specifies the Default Router Preference (DRP) for the router on
    // the target interface.  This column is valid only if the
    // 'ipv6NdRouterPreference' bit of the corresponding instance of cdtIfValid is
    // '1'. The type is CdtIfIpv6NdRouterPreference.
    CdtIfIpv6NdRouterPreference interface{}
}

func (cdtIfTemplateEntry *CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry) GetEntityData() *types.CommonEntityData {
    cdtIfTemplateEntry.EntityData.YFilter = cdtIfTemplateEntry.YFilter
    cdtIfTemplateEntry.EntityData.YangName = "cdtIfTemplateEntry"
    cdtIfTemplateEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtIfTemplateEntry.EntityData.ParentYangName = "cdtIfTemplateTable"
    cdtIfTemplateEntry.EntityData.SegmentPath = "cdtIfTemplateEntry" + types.AddKeyToken(cdtIfTemplateEntry.CdtTemplateName, "cdtTemplateName")
    cdtIfTemplateEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtIfTemplateTable/" + cdtIfTemplateEntry.EntityData.SegmentPath
    cdtIfTemplateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtIfTemplateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtIfTemplateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtIfTemplateEntry.EntityData.Children = types.NewOrderedMap()
    cdtIfTemplateEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtTemplateName", types.YLeaf{"CdtTemplateName", cdtIfTemplateEntry.CdtTemplateName})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfValid", types.YLeaf{"CdtIfValid", cdtIfTemplateEntry.CdtIfValid})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfMtu", types.YLeaf{"CdtIfMtu", cdtIfTemplateEntry.CdtIfMtu})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfCdpEnable", types.YLeaf{"CdtIfCdpEnable", cdtIfTemplateEntry.CdtIfCdpEnable})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfFlowMonitor", types.YLeaf{"CdtIfFlowMonitor", cdtIfTemplateEntry.CdtIfFlowMonitor})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv4Unnumbered", types.YLeaf{"CdtIfIpv4Unnumbered", cdtIfTemplateEntry.CdtIfIpv4Unnumbered})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv4SubEnable", types.YLeaf{"CdtIfIpv4SubEnable", cdtIfTemplateEntry.CdtIfIpv4SubEnable})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv4Mtu", types.YLeaf{"CdtIfIpv4Mtu", cdtIfTemplateEntry.CdtIfIpv4Mtu})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv4TcpMssAdjust", types.YLeaf{"CdtIfIpv4TcpMssAdjust", cdtIfTemplateEntry.CdtIfIpv4TcpMssAdjust})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv4VerifyUniRpf", types.YLeaf{"CdtIfIpv4VerifyUniRpf", cdtIfTemplateEntry.CdtIfIpv4VerifyUniRpf})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv4VerifyUniRpfAcl", types.YLeaf{"CdtIfIpv4VerifyUniRpfAcl", cdtIfTemplateEntry.CdtIfIpv4VerifyUniRpfAcl})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv4VerifyUniRpfOpts", types.YLeaf{"CdtIfIpv4VerifyUniRpfOpts", cdtIfTemplateEntry.CdtIfIpv4VerifyUniRpfOpts})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6Enable", types.YLeaf{"CdtIfIpv6Enable", cdtIfTemplateEntry.CdtIfIpv6Enable})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6SubEnable", types.YLeaf{"CdtIfIpv6SubEnable", cdtIfTemplateEntry.CdtIfIpv6SubEnable})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6TcpMssAdjust", types.YLeaf{"CdtIfIpv6TcpMssAdjust", cdtIfTemplateEntry.CdtIfIpv6TcpMssAdjust})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6VerifyUniRpf", types.YLeaf{"CdtIfIpv6VerifyUniRpf", cdtIfTemplateEntry.CdtIfIpv6VerifyUniRpf})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6VerifyUniRpfAcl", types.YLeaf{"CdtIfIpv6VerifyUniRpfAcl", cdtIfTemplateEntry.CdtIfIpv6VerifyUniRpfAcl})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6VerifyUniRpfOpts", types.YLeaf{"CdtIfIpv6VerifyUniRpfOpts", cdtIfTemplateEntry.CdtIfIpv6VerifyUniRpfOpts})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdPrefix", types.YLeaf{"CdtIfIpv6NdPrefix", cdtIfTemplateEntry.CdtIfIpv6NdPrefix})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdPrefixLength", types.YLeaf{"CdtIfIpv6NdPrefixLength", cdtIfTemplateEntry.CdtIfIpv6NdPrefixLength})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdValidLife", types.YLeaf{"CdtIfIpv6NdValidLife", cdtIfTemplateEntry.CdtIfIpv6NdValidLife})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdPreferredLife", types.YLeaf{"CdtIfIpv6NdPreferredLife", cdtIfTemplateEntry.CdtIfIpv6NdPreferredLife})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdOpts", types.YLeaf{"CdtIfIpv6NdOpts", cdtIfTemplateEntry.CdtIfIpv6NdOpts})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdDadAttempts", types.YLeaf{"CdtIfIpv6NdDadAttempts", cdtIfTemplateEntry.CdtIfIpv6NdDadAttempts})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdNsInterval", types.YLeaf{"CdtIfIpv6NdNsInterval", cdtIfTemplateEntry.CdtIfIpv6NdNsInterval})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdReachableTime", types.YLeaf{"CdtIfIpv6NdReachableTime", cdtIfTemplateEntry.CdtIfIpv6NdReachableTime})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdRaIntervalUnits", types.YLeaf{"CdtIfIpv6NdRaIntervalUnits", cdtIfTemplateEntry.CdtIfIpv6NdRaIntervalUnits})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdRaIntervalMax", types.YLeaf{"CdtIfIpv6NdRaIntervalMax", cdtIfTemplateEntry.CdtIfIpv6NdRaIntervalMax})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdRaIntervalMin", types.YLeaf{"CdtIfIpv6NdRaIntervalMin", cdtIfTemplateEntry.CdtIfIpv6NdRaIntervalMin})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdRaLife", types.YLeaf{"CdtIfIpv6NdRaLife", cdtIfTemplateEntry.CdtIfIpv6NdRaLife})
    cdtIfTemplateEntry.EntityData.Leafs.Append("cdtIfIpv6NdRouterPreference", types.YLeaf{"CdtIfIpv6NdRouterPreference", cdtIfTemplateEntry.CdtIfIpv6NdRouterPreference})

    cdtIfTemplateEntry.EntityData.YListKeys = []string {"CdtTemplateName"}

    return &(cdtIfTemplateEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRaIntervalUnits represents corresponding instance of cdtIfValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRaIntervalUnits string

const (
    CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRaIntervalUnits_seconds CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRaIntervalUnits = "seconds"

    CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRaIntervalUnits_milliseconds CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRaIntervalUnits = "milliseconds"
)

// CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRouterPreference represents the corresponding instance of cdtIfValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRouterPreference string

const (
    CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRouterPreference_high CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRouterPreference = "high"

    CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRouterPreference_medium CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRouterPreference = "medium"

    CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRouterPreference_low CISCODYNAMICTEMPLATEMIB_CdtIfTemplateTable_CdtIfTemplateEntry_CdtIfIpv6NdRouterPreference = "low"
)

// CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable
// This table contains attributes relating to PPP connection
// configuration.
// 
// This table has a sparse-dependent relationship on the
// cdtTemplateTable, containing a row for each dynamic template
// having a cdtTemplateType of one of the following values:
// 
//     'derived'
//     'ppp'
type CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing attributes relating to PPP connection configuration. 
    // The system automatically creates an entry when the system or the EMS/NMS
    // creates a row in the cdtTemplateTable with a cdtTemplateType of one of the
    // following values:      'derived'     'ppp'  Likewise, the system
    // automatically destroys an entry when the system or the EMS/NMS destroys the
    // corresponding row in the cdtTemplateTable. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry.
    CdtPppTemplateEntry []*CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry
}

func (cdtPppTemplateTable *CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable) GetEntityData() *types.CommonEntityData {
    cdtPppTemplateTable.EntityData.YFilter = cdtPppTemplateTable.YFilter
    cdtPppTemplateTable.EntityData.YangName = "cdtPppTemplateTable"
    cdtPppTemplateTable.EntityData.BundleName = "cisco_ios_xe"
    cdtPppTemplateTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtPppTemplateTable.EntityData.SegmentPath = "cdtPppTemplateTable"
    cdtPppTemplateTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtPppTemplateTable.EntityData.SegmentPath
    cdtPppTemplateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtPppTemplateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtPppTemplateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtPppTemplateTable.EntityData.Children = types.NewOrderedMap()
    cdtPppTemplateTable.EntityData.Children.Append("cdtPppTemplateEntry", types.YChild{"CdtPppTemplateEntry", nil})
    for i := range cdtPppTemplateTable.CdtPppTemplateEntry {
        cdtPppTemplateTable.EntityData.Children.Append(types.GetSegmentPath(cdtPppTemplateTable.CdtPppTemplateEntry[i]), types.YChild{"CdtPppTemplateEntry", cdtPppTemplateTable.CdtPppTemplateEntry[i]})
    }
    cdtPppTemplateTable.EntityData.Leafs = types.NewOrderedMap()

    cdtPppTemplateTable.EntityData.YListKeys = []string {}

    return &(cdtPppTemplateTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateName
    CdtTemplateName interface{}

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
    CdtPppValid interface{}

    // This object specifies whether the system applies accounting services to the
    // target PPP connection.  This column is valid only if the 'accounting' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is bool.
    CdtPppAccounting interface{}

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
    CdtPppAuthentication interface{}

    // This object specifies the name of a list of authentication methods used on
    // a target PPP connection.  If the template does not include this attribute,
    // then the system uses the default method list.  This column is valid only if
    // the 'authentication' bit of the corresponding instance of cdtPppValid is
    // '1'. The type is string with length: 0..255.
    CdtPppAuthenticationMethods interface{}

    // This object specifies whether the system applies authorization services to
    // a target PPP connection.  This column is valid only if the 'authorization'
    // bit of the corresponding instance of cditPppValid is '1'. The type is bool.
    CdtPppAuthorization interface{}

    // This object specifies whether the system ignores loopback on a target PPP
    // connection.  When the system ignores loopback, loopback detection is
    // disabled.  This column is valid only if the 'loopbackIgnore' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is bool.
    CdtPppLoopbackIgnore interface{}

    // This object specifies the number of authentication failures allowed by the
    // system before a target PPP connection is reset.  The value '0' cannot be
    // written to an instance of this object. However, it serves as a convenient
    // value when the column is not valid.  This column is valid only if the
    // 'maxBadAuth' bit of the corresponding instance of cdtPppValid is '1'. The
    // type is interface{} with range: 0..4294967295.
    CdtPppMaxBadAuth interface{}

    // This object specifies the number of unacknowledged Configure-Request
    // messages a target PPP connection can send before the system abandons LCP or
    // NCP negotiations.  This column is valid only if the 'maxConfigure' bit of
    // the corresponding instance of cdtPppValid is '1'. The type is interface{}
    // with range: 1..4294967295.
    CdtPppMaxConfigure interface{}

    // This object specifies the number of Configure-Nak messages a target PPP
    // connection can receive before the system abandons LCP or NCP negotiations. 
    // This column is valid only if the 'maxFailure' bit of the corresponding
    // instance of cdtPppValid is '1'. The type is interface{} with range:
    // 1..4294967295.
    CdtPppMaxFailure interface{}

    // This object specifies the number of unacknowledged Terminate-Request
    // messages a target PPP connection can send before the system abandons LCP or
    // NCP negotiations.  This column is valid only if the 'maxTerminate' bit of
    // the corresponding instance of cdtPppValid is '1'. The type is interface{}
    // with range: 1..4294967295.
    CdtPppMaxTerminate interface{}

    // This objects specifies the maximum time the system will wait for a response
    // to an authentication request on a target PPP connection.  This column is
    // valid only if the 'timeoutAuthentication' bit of the corresponding instance
    // of cdtPppValid is '1'. The type is interface{} with range: 1..255. Units
    // are seconds.
    CdtPppTimeoutAuthentication interface{}

    // This objects specifies the maximum time the system will wait for a response
    // to a PPP control packets on a target PPP connection.  This column is valid
    // only if the 'timeoutRetry' bit of the corresponding instance of cdtPppValid
    // is '1'. The type is interface{} with range: 1..255. Units are seconds.
    CdtPppTimeoutRetry interface{}

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
    CdtPppChapOpts interface{}

    // This object specifies the hostname sent in a CHAP response on a target PPP
    // connection.  If the template does not include this attribute, then the
    // system uses its assigned hostname.  This column is valid only if the
    // 'chapHostname' bit of the corresponding instance of cdtPppValid is '1'. The
    // type is string with length: 0..255.
    CdtPppChapHostname interface{}

    // This object specifies the password used to construct a CHAP response on the
    // target PPP connection.  This column is valid only if the 'chapPassword' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is string
    // with length: 0..255.
    CdtPppChapPassword interface{}

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
    CdtPppMsChapV1Opts interface{}

    // This object specifies the hostname sent in a Microsoft CHAP (v1) response
    // on a target PPP connection.  If the template does not include this
    // attribute, then the system uses its assigned hostname.  This column is
    // valid only if the 'msChapV1Hostname' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is string with length: 0..255.
    CdtPppMsChapV1Hostname interface{}

    // This object specifies the password used to construct a Microsoft CHAP (v1)
    // response on a target PPP connection.  This column is valid only if the
    // 'msChapV1Password' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string with length: 0..255.
    CdtPppMsChapV1Password interface{}

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
    CdtPppMsChapV2Opts interface{}

    // This object specifies the hostname sent in a Microsoft CHAP (v2) response
    // on a target PPP connection.  If the template does not include this
    // attribute, then the system uses its assigned hostname.  This column is
    // valid only if the 'msChapV2Hostname' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is string with length: 0..255.
    CdtPppMsChapV2Hostname interface{}

    // This object specifies the password used to construct a Microsoft CHAP (v2)
    // response on a target PPP connection.  This column is valid only if the
    // 'msChapV2Password' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string with length: 0..255.
    CdtPppMsChapV2Password interface{}

    // This object specifies how the system processes the PAP on a target PPP
    // connection:      'refuse'         This option specifies that the system
    // should refuse PAP         requests from peers of a target PPP connection.  
    // 'encrypted'         This option specifies that the value specified by the  
    // corresponding instance of cdtPppPapSentPassword is         already
    // encrypted.  This column is valid only if the 'papOpts' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is map[string]bool.
    CdtPppPapOpts interface{}

    // This object specifies the username sent in a PAP response on a target PPP
    // connection.  This column is valid only if the 'papUsername' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is string with
    // length: 0..255.
    CdtPppPapUsername interface{}

    // This object specifies the username used to construct a PAP response on a
    // target PPP connection.  This column is valid only if the 'papPassword' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is string
    // with length: 0..255.
    CdtPppPapPassword interface{}

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
    CdtPppEapOpts interface{}

    // This object specifies the identity sent in an EAP response on a target PPP
    // connection.  This column is valid only if the 'eapIdentity' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is string with
    // length: 0..255.
    CdtPppEapIdentity interface{}

    // This object specifies the password used to construct an EAP response on a
    // target PPP connection.  This column is valid only if the 'eapPassword' bit
    // of the corresponding instance of cdtPppValid is '1'. The type is string
    // with length: 0..255.
    CdtPppEapPassword interface{}

    // This object specifies the IPCP address option for a target PPP connection: 
    // 'other'         The implementation of this MIB module does not recognize   
    // the configured IPCP address option.      'accept'         The system
    // accepts any non-zero IP address from the peer         of a target PPP
    // connection.      'required'         The system disconnects the peer of a
    // target PPP         connection if it could not negotiate an IP address.     
    // 'unique'         The system disconnects the peer of a target PPP        
    // connection if the IP address is already in use.  This column is valid only
    // if the 'ipcpAddrOption' bit of the corresponding instance of cdtPppValid is
    // '1'. The type is CdtPppIpcpAddrOption.
    CdtPppIpcpAddrOption interface{}

    // This object specifies the IPCP DNS option for the dynamic interface:     
    // 'other'         The implementation of this MIB module does not recognize   
    // the configured DNS option.      'accept'         The system accepts any
    // non-zero DNS address form the         peer of a target PPP connection.     
    // 'request'         The system requests the DNS address from the peer of a   
    // target PPP connection.      'reject'         The system rejects the DNS
    // option from the peer of a         target PPP connection.  This column is
    // valid only if the 'ipcpDnsOption' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is CdtPppIpcpDnsOption.
    CdtPppIpcpDnsOption interface{}

    // This object specifies the IP address of the primary DNS server offered to
    // the peer of a target PPP connection.  This column is valid only if the
    // 'ipcpDnsPrimary' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string.
    CdtPppIpcpDnsPrimary interface{}

    // This object specifies the IP address of the secondary DNS server offered to
    // the peer of a target PPP connection.  This column is valid only if the
    // 'ipcpDnsSecondary' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string.
    CdtPppIpcpDnsSecondary interface{}

    // This object specifies the IPCP WINS option for a target PPP connection:    
    // 'other'         The implementation of this MIB module does not recognize   
    // the configured WINS option.      'accept'         The system accepts any
    // non-zero WINS address from the         peer of a target PPP connection.    
    // 'request'         The system requests the WINS address from the peer of    
    // a target PPP connection.      'reject'         The system rejects the WINS
    // option from the peer of a         target PPP connection.  This column is
    // valid only if the 'ipcpWinsOption' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is CdtPppIpcpWinsOption.
    CdtPppIpcpWinsOption interface{}

    // This object specifies the IP address of the primary WINS server offered to
    // the peer of a target PPP connection.  This column is valid only if the
    // 'ipcpWinsPrimary' bit of the corresponding instance of cdtPppValid is '1'.
    // The type is string.
    CdtPppIpcpWinsPrimary interface{}

    // This object specifies the IP address of the secondary WINS server offered
    // to the peer of a target PPP connection.  This column is valid only if the
    // 'ipcpWinsSecondary' bit of the corresponding instance of cdtPppValid is
    // '1'. The type is string.
    CdtPppIpcpWinsSecondary interface{}

    // This object specifies the IPCP IP subnet mask option for a target PPP
    // connection:      'other'         The implementation of this MIB module does
    // not recognize         the configured IP subnet mask option.      'request' 
    // The system requests the IP subnet mask from the peer of         a target
    // PPP connection.      'reject'         The system rejects the IP subnet mask
    // option from the         peer of a target PPP connection.  This column is
    // valid only if the 'ipcpMaskOption' bit of the corresponding instance of
    // cdtPppValid is '1'. The type is CdtPppIpcpMaskOption.
    CdtPppIpcpMaskOption interface{}

    // This object specifies the IP address mask offered to the peer of a target
    // PPP connection.  This column is valid only if the 'ipcpMask' bit of the
    // corresponding instance of cdtPppValid is '1'. The type is string.
    CdtPppIpcpMask interface{}

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
    CdtPppPeerDefIpAddrOpts interface{}

    // This object specifies how the system assigns an IP address to the peer of a
    // target PPP connection:      'static'         The system assigns the IP
    // address specified by the         corresponding instance of
    // cdtPppPeerDefIpAddr.      'pool'         The system allocates the first
    // available IP address from         the corresponding list of named pools
    // contained by the         cdtPppPeerIpAddrPoolTable.      'dhcp'         The
    // system acts as a DHCP proxy-client to obtain an IP         address.  This
    // column is valid only if the 'peerDefIpAddrSrc' bit of the corresponding
    // instance of cdtPppValid is '1'. The type is CdtPppPeerDefIpAddrSrc.
    CdtPppPeerDefIpAddrSrc interface{}

    // This object specifies the IP address the system assigns to the peer of a
    // target PPP connection.  This column is valid only if the 'peerDefIpAddr'
    // bit of the corresponding instance of cdtPppValid is '1'. The type is
    // string.
    CdtPppPeerDefIpAddr interface{}
}

func (cdtPppTemplateEntry *CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry) GetEntityData() *types.CommonEntityData {
    cdtPppTemplateEntry.EntityData.YFilter = cdtPppTemplateEntry.YFilter
    cdtPppTemplateEntry.EntityData.YangName = "cdtPppTemplateEntry"
    cdtPppTemplateEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtPppTemplateEntry.EntityData.ParentYangName = "cdtPppTemplateTable"
    cdtPppTemplateEntry.EntityData.SegmentPath = "cdtPppTemplateEntry" + types.AddKeyToken(cdtPppTemplateEntry.CdtTemplateName, "cdtTemplateName")
    cdtPppTemplateEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtPppTemplateTable/" + cdtPppTemplateEntry.EntityData.SegmentPath
    cdtPppTemplateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtPppTemplateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtPppTemplateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtPppTemplateEntry.EntityData.Children = types.NewOrderedMap()
    cdtPppTemplateEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtTemplateName", types.YLeaf{"CdtTemplateName", cdtPppTemplateEntry.CdtTemplateName})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppValid", types.YLeaf{"CdtPppValid", cdtPppTemplateEntry.CdtPppValid})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppAccounting", types.YLeaf{"CdtPppAccounting", cdtPppTemplateEntry.CdtPppAccounting})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppAuthentication", types.YLeaf{"CdtPppAuthentication", cdtPppTemplateEntry.CdtPppAuthentication})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppAuthenticationMethods", types.YLeaf{"CdtPppAuthenticationMethods", cdtPppTemplateEntry.CdtPppAuthenticationMethods})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppAuthorization", types.YLeaf{"CdtPppAuthorization", cdtPppTemplateEntry.CdtPppAuthorization})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppLoopbackIgnore", types.YLeaf{"CdtPppLoopbackIgnore", cdtPppTemplateEntry.CdtPppLoopbackIgnore})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMaxBadAuth", types.YLeaf{"CdtPppMaxBadAuth", cdtPppTemplateEntry.CdtPppMaxBadAuth})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMaxConfigure", types.YLeaf{"CdtPppMaxConfigure", cdtPppTemplateEntry.CdtPppMaxConfigure})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMaxFailure", types.YLeaf{"CdtPppMaxFailure", cdtPppTemplateEntry.CdtPppMaxFailure})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMaxTerminate", types.YLeaf{"CdtPppMaxTerminate", cdtPppTemplateEntry.CdtPppMaxTerminate})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppTimeoutAuthentication", types.YLeaf{"CdtPppTimeoutAuthentication", cdtPppTemplateEntry.CdtPppTimeoutAuthentication})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppTimeoutRetry", types.YLeaf{"CdtPppTimeoutRetry", cdtPppTemplateEntry.CdtPppTimeoutRetry})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppChapOpts", types.YLeaf{"CdtPppChapOpts", cdtPppTemplateEntry.CdtPppChapOpts})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppChapHostname", types.YLeaf{"CdtPppChapHostname", cdtPppTemplateEntry.CdtPppChapHostname})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppChapPassword", types.YLeaf{"CdtPppChapPassword", cdtPppTemplateEntry.CdtPppChapPassword})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMsChapV1Opts", types.YLeaf{"CdtPppMsChapV1Opts", cdtPppTemplateEntry.CdtPppMsChapV1Opts})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMsChapV1Hostname", types.YLeaf{"CdtPppMsChapV1Hostname", cdtPppTemplateEntry.CdtPppMsChapV1Hostname})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMsChapV1Password", types.YLeaf{"CdtPppMsChapV1Password", cdtPppTemplateEntry.CdtPppMsChapV1Password})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMsChapV2Opts", types.YLeaf{"CdtPppMsChapV2Opts", cdtPppTemplateEntry.CdtPppMsChapV2Opts})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMsChapV2Hostname", types.YLeaf{"CdtPppMsChapV2Hostname", cdtPppTemplateEntry.CdtPppMsChapV2Hostname})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppMsChapV2Password", types.YLeaf{"CdtPppMsChapV2Password", cdtPppTemplateEntry.CdtPppMsChapV2Password})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppPapOpts", types.YLeaf{"CdtPppPapOpts", cdtPppTemplateEntry.CdtPppPapOpts})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppPapUsername", types.YLeaf{"CdtPppPapUsername", cdtPppTemplateEntry.CdtPppPapUsername})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppPapPassword", types.YLeaf{"CdtPppPapPassword", cdtPppTemplateEntry.CdtPppPapPassword})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppEapOpts", types.YLeaf{"CdtPppEapOpts", cdtPppTemplateEntry.CdtPppEapOpts})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppEapIdentity", types.YLeaf{"CdtPppEapIdentity", cdtPppTemplateEntry.CdtPppEapIdentity})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppEapPassword", types.YLeaf{"CdtPppEapPassword", cdtPppTemplateEntry.CdtPppEapPassword})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpAddrOption", types.YLeaf{"CdtPppIpcpAddrOption", cdtPppTemplateEntry.CdtPppIpcpAddrOption})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpDnsOption", types.YLeaf{"CdtPppIpcpDnsOption", cdtPppTemplateEntry.CdtPppIpcpDnsOption})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpDnsPrimary", types.YLeaf{"CdtPppIpcpDnsPrimary", cdtPppTemplateEntry.CdtPppIpcpDnsPrimary})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpDnsSecondary", types.YLeaf{"CdtPppIpcpDnsSecondary", cdtPppTemplateEntry.CdtPppIpcpDnsSecondary})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpWinsOption", types.YLeaf{"CdtPppIpcpWinsOption", cdtPppTemplateEntry.CdtPppIpcpWinsOption})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpWinsPrimary", types.YLeaf{"CdtPppIpcpWinsPrimary", cdtPppTemplateEntry.CdtPppIpcpWinsPrimary})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpWinsSecondary", types.YLeaf{"CdtPppIpcpWinsSecondary", cdtPppTemplateEntry.CdtPppIpcpWinsSecondary})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpMaskOption", types.YLeaf{"CdtPppIpcpMaskOption", cdtPppTemplateEntry.CdtPppIpcpMaskOption})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppIpcpMask", types.YLeaf{"CdtPppIpcpMask", cdtPppTemplateEntry.CdtPppIpcpMask})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppPeerDefIpAddrOpts", types.YLeaf{"CdtPppPeerDefIpAddrOpts", cdtPppTemplateEntry.CdtPppPeerDefIpAddrOpts})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppPeerDefIpAddrSrc", types.YLeaf{"CdtPppPeerDefIpAddrSrc", cdtPppTemplateEntry.CdtPppPeerDefIpAddrSrc})
    cdtPppTemplateEntry.EntityData.Leafs.Append("cdtPppPeerDefIpAddr", types.YLeaf{"CdtPppPeerDefIpAddr", cdtPppTemplateEntry.CdtPppPeerDefIpAddr})

    cdtPppTemplateEntry.EntityData.YListKeys = []string {"CdtTemplateName"}

    return &(cdtPppTemplateEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption string

const (
    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption_other CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption = "other"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption_accept CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption = "accept"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption_required CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption = "required"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption_unique CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpAddrOption = "unique"
)

// CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption string

const (
    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption_other CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption = "other"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption_accept CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption = "accept"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption_request CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption = "request"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption_reject CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpDnsOption = "reject"
)

// CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpMaskOption represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpMaskOption string

const (
    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpMaskOption_other CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpMaskOption = "other"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpMaskOption_request CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpMaskOption = "request"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpMaskOption_reject CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpMaskOption = "reject"
)

// CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption string

const (
    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption_other CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption = "other"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption_accept CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption = "accept"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption_request CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption = "request"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption_reject CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppIpcpWinsOption = "reject"
)

// CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppPeerDefIpAddrSrc represents corresponding instance of cdtPppValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppPeerDefIpAddrSrc string

const (
    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppPeerDefIpAddrSrc_static CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppPeerDefIpAddrSrc = "static"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppPeerDefIpAddrSrc_pool CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppPeerDefIpAddrSrc = "pool"

    CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppPeerDefIpAddrSrc_dhcp CISCODYNAMICTEMPLATEMIB_CdtPppTemplateTable_CdtPppTemplateEntry_CdtPppPeerDefIpAddrSrc = "dhcp"
)

// CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable
// This table contains a prioritized list of named pools for each
// PPP template.  A list corresponding to a PPP template
// specifies the pools the system searches in an attempt to
// assign an IP address to the peer of a target PPP connection.
// The system searches the pools in the order of their priority.
// 
// This table has an expansion dependent relationship on the
// cdtPppTemplateTable, containing zero or more rows for each PPP
// template.
type CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable struct {
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
    // CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable_CdtPppPeerIpAddrPoolEntry.
    CdtPppPeerIpAddrPoolEntry []*CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable_CdtPppPeerIpAddrPoolEntry
}

func (cdtPppPeerIpAddrPoolTable *CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable) GetEntityData() *types.CommonEntityData {
    cdtPppPeerIpAddrPoolTable.EntityData.YFilter = cdtPppPeerIpAddrPoolTable.YFilter
    cdtPppPeerIpAddrPoolTable.EntityData.YangName = "cdtPppPeerIpAddrPoolTable"
    cdtPppPeerIpAddrPoolTable.EntityData.BundleName = "cisco_ios_xe"
    cdtPppPeerIpAddrPoolTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtPppPeerIpAddrPoolTable.EntityData.SegmentPath = "cdtPppPeerIpAddrPoolTable"
    cdtPppPeerIpAddrPoolTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtPppPeerIpAddrPoolTable.EntityData.SegmentPath
    cdtPppPeerIpAddrPoolTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtPppPeerIpAddrPoolTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtPppPeerIpAddrPoolTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtPppPeerIpAddrPoolTable.EntityData.Children = types.NewOrderedMap()
    cdtPppPeerIpAddrPoolTable.EntityData.Children.Append("cdtPppPeerIpAddrPoolEntry", types.YChild{"CdtPppPeerIpAddrPoolEntry", nil})
    for i := range cdtPppPeerIpAddrPoolTable.CdtPppPeerIpAddrPoolEntry {
        cdtPppPeerIpAddrPoolTable.EntityData.Children.Append(types.GetSegmentPath(cdtPppPeerIpAddrPoolTable.CdtPppPeerIpAddrPoolEntry[i]), types.YChild{"CdtPppPeerIpAddrPoolEntry", cdtPppPeerIpAddrPoolTable.CdtPppPeerIpAddrPoolEntry[i]})
    }
    cdtPppPeerIpAddrPoolTable.EntityData.Leafs = types.NewOrderedMap()

    cdtPppPeerIpAddrPoolTable.EntityData.YListKeys = []string {}

    return &(cdtPppPeerIpAddrPoolTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable_CdtPppPeerIpAddrPoolEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable_CdtPppPeerIpAddrPoolEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateName
    CdtTemplateName interface{}

    // This attribute is a key. This object indicates the relative priority of the
    // named pool in the list corresponding to a PPP template.  The system
    // searches pools in the order of priority, where lower values represent
    // higher priority. The type is interface{} with range: 1..4294967295.
    CdtPppPeerIpAddrPoolPriority interface{}

    // This object specifies the status of the entry.  The following columns must
    // be valid before activating a subscriber access profile:      -
    // cdtPppPeerIpAddrPoolStorage     - cdtPppPeerIpAddrPoolName  However, these
    // objects specify a default value.  Thus, it is possible to use create-and-go
    // semantics without setting any additional columns.  An implementation must
    // not allow the EMS/NMS to create an entry if the corresponding instance of
    // cdtPppPeerIpAddrSrc is not 'pool'.  An implementation must allow the
    // EMS/NMS to modify any column when this column is 'active'. The type is
    // RowStatus.
    CdtPppPeerIpAddrPoolStatus interface{}

    // This object specifies what happens to the name pool entry upon restart.  If
    // the corresponding instance of cdtTemplateSrc is not 'local', then this
    // column must be 'volatile'. The type is StorageType.
    CdtPppPeerIpAddrPoolStorage interface{}

    // This object specifies the name of the IP address pool associated with the
    // PPP template. The type is string with length: 0..255.
    CdtPppPeerIpAddrPoolName interface{}
}

func (cdtPppPeerIpAddrPoolEntry *CISCODYNAMICTEMPLATEMIB_CdtPppPeerIpAddrPoolTable_CdtPppPeerIpAddrPoolEntry) GetEntityData() *types.CommonEntityData {
    cdtPppPeerIpAddrPoolEntry.EntityData.YFilter = cdtPppPeerIpAddrPoolEntry.YFilter
    cdtPppPeerIpAddrPoolEntry.EntityData.YangName = "cdtPppPeerIpAddrPoolEntry"
    cdtPppPeerIpAddrPoolEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtPppPeerIpAddrPoolEntry.EntityData.ParentYangName = "cdtPppPeerIpAddrPoolTable"
    cdtPppPeerIpAddrPoolEntry.EntityData.SegmentPath = "cdtPppPeerIpAddrPoolEntry" + types.AddKeyToken(cdtPppPeerIpAddrPoolEntry.CdtTemplateName, "cdtTemplateName") + types.AddKeyToken(cdtPppPeerIpAddrPoolEntry.CdtPppPeerIpAddrPoolPriority, "cdtPppPeerIpAddrPoolPriority")
    cdtPppPeerIpAddrPoolEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtPppPeerIpAddrPoolTable/" + cdtPppPeerIpAddrPoolEntry.EntityData.SegmentPath
    cdtPppPeerIpAddrPoolEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtPppPeerIpAddrPoolEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtPppPeerIpAddrPoolEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtPppPeerIpAddrPoolEntry.EntityData.Children = types.NewOrderedMap()
    cdtPppPeerIpAddrPoolEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtPppPeerIpAddrPoolEntry.EntityData.Leafs.Append("cdtTemplateName", types.YLeaf{"CdtTemplateName", cdtPppPeerIpAddrPoolEntry.CdtTemplateName})
    cdtPppPeerIpAddrPoolEntry.EntityData.Leafs.Append("cdtPppPeerIpAddrPoolPriority", types.YLeaf{"CdtPppPeerIpAddrPoolPriority", cdtPppPeerIpAddrPoolEntry.CdtPppPeerIpAddrPoolPriority})
    cdtPppPeerIpAddrPoolEntry.EntityData.Leafs.Append("cdtPppPeerIpAddrPoolStatus", types.YLeaf{"CdtPppPeerIpAddrPoolStatus", cdtPppPeerIpAddrPoolEntry.CdtPppPeerIpAddrPoolStatus})
    cdtPppPeerIpAddrPoolEntry.EntityData.Leafs.Append("cdtPppPeerIpAddrPoolStorage", types.YLeaf{"CdtPppPeerIpAddrPoolStorage", cdtPppPeerIpAddrPoolEntry.CdtPppPeerIpAddrPoolStorage})
    cdtPppPeerIpAddrPoolEntry.EntityData.Leafs.Append("cdtPppPeerIpAddrPoolName", types.YLeaf{"CdtPppPeerIpAddrPoolName", cdtPppPeerIpAddrPoolEntry.CdtPppPeerIpAddrPoolName})

    cdtPppPeerIpAddrPoolEntry.EntityData.YListKeys = []string {"CdtTemplateName", "CdtPppPeerIpAddrPoolPriority"}

    return &(cdtPppPeerIpAddrPoolEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable
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
type CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing attributes relating to dynamic interfaces initiated on
    // Ethernet virtual interfaces (e.g., EoMPLS) or automatically created VLANs. 
    // The system automatically creates an entry when the system or the EMS/NMS
    // creates a row in the cdtTemplateTable with a cdtTemplateType of one of the
    // following values:      'derived'     'ethernet'  Likewise, the system
    // automatically destroys an entry when the system or the EMS/NMS destroys the
    // corresponding row in the cdtTemplateTable. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable_CdtEthernetTemplateEntry.
    CdtEthernetTemplateEntry []*CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable_CdtEthernetTemplateEntry
}

func (cdtEthernetTemplateTable *CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable) GetEntityData() *types.CommonEntityData {
    cdtEthernetTemplateTable.EntityData.YFilter = cdtEthernetTemplateTable.YFilter
    cdtEthernetTemplateTable.EntityData.YangName = "cdtEthernetTemplateTable"
    cdtEthernetTemplateTable.EntityData.BundleName = "cisco_ios_xe"
    cdtEthernetTemplateTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtEthernetTemplateTable.EntityData.SegmentPath = "cdtEthernetTemplateTable"
    cdtEthernetTemplateTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtEthernetTemplateTable.EntityData.SegmentPath
    cdtEthernetTemplateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtEthernetTemplateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtEthernetTemplateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtEthernetTemplateTable.EntityData.Children = types.NewOrderedMap()
    cdtEthernetTemplateTable.EntityData.Children.Append("cdtEthernetTemplateEntry", types.YChild{"CdtEthernetTemplateEntry", nil})
    for i := range cdtEthernetTemplateTable.CdtEthernetTemplateEntry {
        cdtEthernetTemplateTable.EntityData.Children.Append(types.GetSegmentPath(cdtEthernetTemplateTable.CdtEthernetTemplateEntry[i]), types.YChild{"CdtEthernetTemplateEntry", cdtEthernetTemplateTable.CdtEthernetTemplateEntry[i]})
    }
    cdtEthernetTemplateTable.EntityData.Leafs = types.NewOrderedMap()

    cdtEthernetTemplateTable.EntityData.YListKeys = []string {}

    return &(cdtEthernetTemplateTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable_CdtEthernetTemplateEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable_CdtEthernetTemplateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateName
    CdtTemplateName interface{}

    // This object specifies which attributes in the dynamic template have been
    // configured to valid values.  Each bit in this bit string corresponds to a
    // column in this table.  If the bit is '0', then the value of the
    // corresponding column is not valid.  If the bit is '1', then the value of
    // the corresponding column has been configured to a valid value.  The
    // following list specifies the mappings between bits and the columns:     
    // bridgeDomain     => cdtEthernetBridgeDomain     pppoeEnable      =>
    // cdtEthernetPppoeEnable     ipv4PointToPoint => cdtEthernetIpv4PointToPoint 
    // macAddr          => cdtEthernetMacAddr. The type is map[string]bool.
    CdtEthernetValid interface{}

    // This object specifies the name of the bridge domain... The type is string
    // with length: 0..255.
    CdtEthernetBridgeDomain interface{}

    // This object specifies whether... The type is bool.
    CdtEthernetPppoeEnable interface{}

    // This object specifies whether... The type is bool.
    CdtEthernetIpv4PointToPoint interface{}

    // This object specifies the... The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    CdtEthernetMacAddr interface{}
}

func (cdtEthernetTemplateEntry *CISCODYNAMICTEMPLATEMIB_CdtEthernetTemplateTable_CdtEthernetTemplateEntry) GetEntityData() *types.CommonEntityData {
    cdtEthernetTemplateEntry.EntityData.YFilter = cdtEthernetTemplateEntry.YFilter
    cdtEthernetTemplateEntry.EntityData.YangName = "cdtEthernetTemplateEntry"
    cdtEthernetTemplateEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtEthernetTemplateEntry.EntityData.ParentYangName = "cdtEthernetTemplateTable"
    cdtEthernetTemplateEntry.EntityData.SegmentPath = "cdtEthernetTemplateEntry" + types.AddKeyToken(cdtEthernetTemplateEntry.CdtTemplateName, "cdtTemplateName")
    cdtEthernetTemplateEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtEthernetTemplateTable/" + cdtEthernetTemplateEntry.EntityData.SegmentPath
    cdtEthernetTemplateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtEthernetTemplateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtEthernetTemplateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtEthernetTemplateEntry.EntityData.Children = types.NewOrderedMap()
    cdtEthernetTemplateEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtEthernetTemplateEntry.EntityData.Leafs.Append("cdtTemplateName", types.YLeaf{"CdtTemplateName", cdtEthernetTemplateEntry.CdtTemplateName})
    cdtEthernetTemplateEntry.EntityData.Leafs.Append("cdtEthernetValid", types.YLeaf{"CdtEthernetValid", cdtEthernetTemplateEntry.CdtEthernetValid})
    cdtEthernetTemplateEntry.EntityData.Leafs.Append("cdtEthernetBridgeDomain", types.YLeaf{"CdtEthernetBridgeDomain", cdtEthernetTemplateEntry.CdtEthernetBridgeDomain})
    cdtEthernetTemplateEntry.EntityData.Leafs.Append("cdtEthernetPppoeEnable", types.YLeaf{"CdtEthernetPppoeEnable", cdtEthernetTemplateEntry.CdtEthernetPppoeEnable})
    cdtEthernetTemplateEntry.EntityData.Leafs.Append("cdtEthernetIpv4PointToPoint", types.YLeaf{"CdtEthernetIpv4PointToPoint", cdtEthernetTemplateEntry.CdtEthernetIpv4PointToPoint})
    cdtEthernetTemplateEntry.EntityData.Leafs.Append("cdtEthernetMacAddr", types.YLeaf{"CdtEthernetMacAddr", cdtEthernetTemplateEntry.CdtEthernetMacAddr})

    cdtEthernetTemplateEntry.EntityData.YListKeys = []string {"CdtTemplateName"}

    return &(cdtEthernetTemplateEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable
// This table contains attributes relating to a service.
// 
// This table has a sparse-dependent relationship on the
// cdtTemplateTable, containing a row for each dynamic template
// having a cdtTemplateType of one of the following values:
// 
//     'derived'
//     'service'
type CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing attributes relating to a service.  The system
    // automatically creates entry when the system or the EMS/NMS creates a row in
    // the cdtTemplateTable with a cdtTemplateType of one of the following values:
    // 'derived'     'service'  Likewise, the system automatically destroys an
    // entry when the system or the EMS/NMS destroys the corresponding row in the
    // cdtTemplateTable. The type is slice of
    // CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry.
    CdtSrvTemplateEntry []*CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry
}

func (cdtSrvTemplateTable *CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable) GetEntityData() *types.CommonEntityData {
    cdtSrvTemplateTable.EntityData.YFilter = cdtSrvTemplateTable.YFilter
    cdtSrvTemplateTable.EntityData.YangName = "cdtSrvTemplateTable"
    cdtSrvTemplateTable.EntityData.BundleName = "cisco_ios_xe"
    cdtSrvTemplateTable.EntityData.ParentYangName = "CISCO-DYNAMIC-TEMPLATE-MIB"
    cdtSrvTemplateTable.EntityData.SegmentPath = "cdtSrvTemplateTable"
    cdtSrvTemplateTable.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/" + cdtSrvTemplateTable.EntityData.SegmentPath
    cdtSrvTemplateTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtSrvTemplateTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtSrvTemplateTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtSrvTemplateTable.EntityData.Children = types.NewOrderedMap()
    cdtSrvTemplateTable.EntityData.Children.Append("cdtSrvTemplateEntry", types.YChild{"CdtSrvTemplateEntry", nil})
    for i := range cdtSrvTemplateTable.CdtSrvTemplateEntry {
        cdtSrvTemplateTable.EntityData.Children.Append(types.GetSegmentPath(cdtSrvTemplateTable.CdtSrvTemplateEntry[i]), types.YChild{"CdtSrvTemplateEntry", cdtSrvTemplateTable.CdtSrvTemplateEntry[i]})
    }
    cdtSrvTemplateTable.EntityData.Leafs = types.NewOrderedMap()

    cdtSrvTemplateTable.EntityData.YListKeys = []string {}

    return &(cdtSrvTemplateTable.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry
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
type CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with length: 1..64. Refers to
    // cisco_dynamic_template_mib.CISCODYNAMICTEMPLATEMIB_CdtTemplateTable_CdtTemplateEntry_CdtTemplateName
    CdtTemplateName interface{}

    // This object specifies which attributes in the dynamic template have been
    // configured to valid values.  Each bit in this bit string corresponds to a
    // column in this table.  If the bit is '0', then the value of the
    // corresponding column is not valid.  If the bit is '1', then the value of
    // the corresponding column has been configured to a valid value.  The
    // following list specifies the mappings between bits and the columns:     
    // networkSrv     => cdtSrvNetworkSrv     vpdnGroup      => cdtSrvVpdnGroup   
    // sgSrvGroup     => cdtSrvGroup     sgSrvType      => cdtSrvSgSrvType    
    // multicast      => cdtSrvMulticast. The type is map[string]bool.
    CdtSrvValid interface{}

    // This object specifies the type of network service provided by the target
    // service:      'other'         The implementation of this MIB module does
    // not recognize         the configured network service.      'none'        
    // The target subscriber service does not provide a network         service to
    // subscribers sessions.      'local'         The target subscriber service
    // provides local termination         for subscriber sessions.      'vpdn'    
    // The target subscriber service provides a Virtual Private         Dialup
    // Network service for subscriber sessions.  This column is valid only if the
    // 'networkSrv' bit of the corresponding instance of cdtSrvValid is '1'. The
    // type is CdtSrvNetworkSrv.
    CdtSrvNetworkSrv interface{}

    // This object specifies the name of the VPDN group used to configure the
    // network service.  This column is valid only if the 'vpdnGroup' bit of the
    // corresponding instance of cdtSrvValid is '1'. The type is string with
    // length: 0..255.
    CdtSrvVpdnGroup interface{}

    // This object specifies the name of the service group with which the system
    // associates subscriber sessions.  A service group specifies a set of
    // services that may be active simultaneously for a given subscriber session. 
    // Typically, a service group contains a primary service and one or more
    // secondary services.  This column is valid only if the 'sgSrvGroup' bit of
    // the corresponding instance of cdtSrvValid is '1'. The type is string with
    // length: 0..255.
    CdtSrvSgSrvGroup interface{}

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
    // The type is CdtSrvSgSrvType.
    CdtSrvSgSrvType interface{}

    // This objects specifies whether the system enables multicast service for
    // subscriber sessions of the target service.  This column is valid only if
    // the 'sgSrvMcastRoutingIf' bit of the corresponding instance of cdtSrvValid
    // is '1'. The type is bool.
    CdtSrvMulticast interface{}
}

func (cdtSrvTemplateEntry *CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry) GetEntityData() *types.CommonEntityData {
    cdtSrvTemplateEntry.EntityData.YFilter = cdtSrvTemplateEntry.YFilter
    cdtSrvTemplateEntry.EntityData.YangName = "cdtSrvTemplateEntry"
    cdtSrvTemplateEntry.EntityData.BundleName = "cisco_ios_xe"
    cdtSrvTemplateEntry.EntityData.ParentYangName = "cdtSrvTemplateTable"
    cdtSrvTemplateEntry.EntityData.SegmentPath = "cdtSrvTemplateEntry" + types.AddKeyToken(cdtSrvTemplateEntry.CdtTemplateName, "cdtTemplateName")
    cdtSrvTemplateEntry.EntityData.AbsolutePath = "CISCO-DYNAMIC-TEMPLATE-MIB:CISCO-DYNAMIC-TEMPLATE-MIB/cdtSrvTemplateTable/" + cdtSrvTemplateEntry.EntityData.SegmentPath
    cdtSrvTemplateEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdtSrvTemplateEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdtSrvTemplateEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdtSrvTemplateEntry.EntityData.Children = types.NewOrderedMap()
    cdtSrvTemplateEntry.EntityData.Leafs = types.NewOrderedMap()
    cdtSrvTemplateEntry.EntityData.Leafs.Append("cdtTemplateName", types.YLeaf{"CdtTemplateName", cdtSrvTemplateEntry.CdtTemplateName})
    cdtSrvTemplateEntry.EntityData.Leafs.Append("cdtSrvValid", types.YLeaf{"CdtSrvValid", cdtSrvTemplateEntry.CdtSrvValid})
    cdtSrvTemplateEntry.EntityData.Leafs.Append("cdtSrvNetworkSrv", types.YLeaf{"CdtSrvNetworkSrv", cdtSrvTemplateEntry.CdtSrvNetworkSrv})
    cdtSrvTemplateEntry.EntityData.Leafs.Append("cdtSrvVpdnGroup", types.YLeaf{"CdtSrvVpdnGroup", cdtSrvTemplateEntry.CdtSrvVpdnGroup})
    cdtSrvTemplateEntry.EntityData.Leafs.Append("cdtSrvSgSrvGroup", types.YLeaf{"CdtSrvSgSrvGroup", cdtSrvTemplateEntry.CdtSrvSgSrvGroup})
    cdtSrvTemplateEntry.EntityData.Leafs.Append("cdtSrvSgSrvType", types.YLeaf{"CdtSrvSgSrvType", cdtSrvTemplateEntry.CdtSrvSgSrvType})
    cdtSrvTemplateEntry.EntityData.Leafs.Append("cdtSrvMulticast", types.YLeaf{"CdtSrvMulticast", cdtSrvTemplateEntry.CdtSrvMulticast})

    cdtSrvTemplateEntry.EntityData.YListKeys = []string {"CdtTemplateName"}

    return &(cdtSrvTemplateEntry.EntityData)
}

// CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv represents corresponding instance of cdtSrvValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv string

const (
    CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv_other CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv = "other"

    CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv_none CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv = "none"

    CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv_local CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv = "local"

    CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv_vpdn CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvNetworkSrv = "vpdn"
)

// CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvSgSrvType represents corresponding instance of cdtSrvValid is '1'.
type CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvSgSrvType string

const (
    CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvSgSrvType_primary CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvSgSrvType = "primary"

    CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvSgSrvType_secondary CISCODYNAMICTEMPLATEMIB_CdtSrvTemplateTable_CdtSrvTemplateEntry_CdtSrvSgSrvType = "secondary"
)

