// This MIB module defines the managed objects for
// representing targets which have class-based policy  
// mappings.  A target can be any logical interface 
// to which a class-based policy is able to be associated.
package cisco_cbp_target_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_cbp_target_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-CBP-TARGET-MIB CISCO-CBP-TARGET-MIB}", reflect.TypeOf(CISCOCBPTARGETMIB{}))
    ydk.RegisterEntity("CISCO-CBP-TARGET-MIB:CISCO-CBP-TARGET-MIB", reflect.TypeOf(CISCOCBPTARGETMIB{}))
}

// CISCOCBPTARGETMIB
type CISCOCBPTARGETMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CcbptTargetAttachCfg CISCOCBPTARGETMIB_CcbptTargetAttachCfg

    // This table describes the class-based policy attachments to to specific
    // targets.
    CcbptTargetTable CISCOCBPTARGETMIB_CcbptTargetTable
}

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetEntityData() *types.CommonEntityData {
    cISCOCBPTARGETMIB.EntityData.YFilter = cISCOCBPTARGETMIB.YFilter
    cISCOCBPTARGETMIB.EntityData.YangName = "CISCO-CBP-TARGET-MIB"
    cISCOCBPTARGETMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCBPTARGETMIB.EntityData.ParentYangName = "CISCO-CBP-TARGET-MIB"
    cISCOCBPTARGETMIB.EntityData.SegmentPath = "CISCO-CBP-TARGET-MIB:CISCO-CBP-TARGET-MIB"
    cISCOCBPTARGETMIB.EntityData.AbsolutePath = cISCOCBPTARGETMIB.EntityData.SegmentPath
    cISCOCBPTARGETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCBPTARGETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCBPTARGETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCBPTARGETMIB.EntityData.Children = types.NewOrderedMap()
    cISCOCBPTARGETMIB.EntityData.Children.Append("ccbptTargetAttachCfg", types.YChild{"CcbptTargetAttachCfg", &cISCOCBPTARGETMIB.CcbptTargetAttachCfg})
    cISCOCBPTARGETMIB.EntityData.Children.Append("ccbptTargetTable", types.YChild{"CcbptTargetTable", &cISCOCBPTARGETMIB.CcbptTargetTable})
    cISCOCBPTARGETMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOCBPTARGETMIB.EntityData.YListKeys = []string {}

    return &(cISCOCBPTARGETMIB.EntityData)
}

// CISCOCBPTARGETMIB_CcbptTargetAttachCfg
type CISCOCBPTARGETMIB_CcbptTargetAttachCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the next available value of  ccbptPolicyId that can
    // be used to create a new conceptual row in the ccbptTargetTable.  If no
    // available identifier exists, then this object will have the value '0'. The
    // type is interface{} with range: 0..4294967295.
    CcbptPolicyIdNext interface{}

    // The value of sysUpTime at the time of the last change to an entry in the
    // ccbptTargetTable. The type is interface{} with range: 0..4294967295.
    CcbptTargetTableLastChange interface{}
}

func (ccbptTargetAttachCfg *CISCOCBPTARGETMIB_CcbptTargetAttachCfg) GetEntityData() *types.CommonEntityData {
    ccbptTargetAttachCfg.EntityData.YFilter = ccbptTargetAttachCfg.YFilter
    ccbptTargetAttachCfg.EntityData.YangName = "ccbptTargetAttachCfg"
    ccbptTargetAttachCfg.EntityData.BundleName = "cisco_ios_xe"
    ccbptTargetAttachCfg.EntityData.ParentYangName = "CISCO-CBP-TARGET-MIB"
    ccbptTargetAttachCfg.EntityData.SegmentPath = "ccbptTargetAttachCfg"
    ccbptTargetAttachCfg.EntityData.AbsolutePath = "CISCO-CBP-TARGET-MIB:CISCO-CBP-TARGET-MIB/" + ccbptTargetAttachCfg.EntityData.SegmentPath
    ccbptTargetAttachCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccbptTargetAttachCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccbptTargetAttachCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccbptTargetAttachCfg.EntityData.Children = types.NewOrderedMap()
    ccbptTargetAttachCfg.EntityData.Leafs = types.NewOrderedMap()
    ccbptTargetAttachCfg.EntityData.Leafs.Append("ccbptPolicyIdNext", types.YLeaf{"CcbptPolicyIdNext", ccbptTargetAttachCfg.CcbptPolicyIdNext})
    ccbptTargetAttachCfg.EntityData.Leafs.Append("ccbptTargetTableLastChange", types.YLeaf{"CcbptTargetTableLastChange", ccbptTargetAttachCfg.CcbptTargetTableLastChange})

    ccbptTargetAttachCfg.EntityData.YListKeys = []string {}

    return &(ccbptTargetAttachCfg.EntityData)
}

// CISCOCBPTARGETMIB_CcbptTargetTable
// This table describes the class-based policy attachments to
// to specific targets.
type CISCOCBPTARGETMIB_CcbptTargetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry describes a class-based policy attachment to a  particular
    // target.    The ccbptTargetType uniquely identifies the type of target in
    // the attachment.  Additionally, the ccbptTargetId uniquely identifies the
    // target in attachment and is of the format indicated by the ccbptTargetType.
    // The ccbptTargetDir  identifies the direction, relative to the
    // ccbptTargetId,  to which the policy is attached.  The ccbptPolicySourceType
    // identifies the source-type of the policy attached.  The  ccbptPolicyId
    // uniquely identifies the policy within the scope of ccbptTargetType,
    // ccbptTargetId, ccbptTargetDir, and  ccbptPolicySourceType.  A class-based
    // policy attachment to a target can be created  through other network
    // management interfaces (e.g., the local console), in which case the SNMP
    // entity will automatically  create an entry in this table.  A class-based
    // policy attachment to a target can be destroyed through other network
    // management interfaces, in which case the SNMP entity will automatically
    // destroy the corresponding entry in this table.  A class-based policy
    // attachment to a target can be created, destroyed, and modified through the
    // SNMP using  ccbptTargetStatus using the semantics described by the 
    // RowStatus Textual Convention.  However, when creating a new class-based
    // policy attachment to a target, the value of ccbptPolicyIdNext should be
    // used to identify the new policy within the scope of the target type,
    // identifier, direction, and policy-source type. The type is slice of
    // CISCOCBPTARGETMIB_CcbptTargetTable_CcbptTargetEntry.
    CcbptTargetEntry []*CISCOCBPTARGETMIB_CcbptTargetTable_CcbptTargetEntry
}

func (ccbptTargetTable *CISCOCBPTARGETMIB_CcbptTargetTable) GetEntityData() *types.CommonEntityData {
    ccbptTargetTable.EntityData.YFilter = ccbptTargetTable.YFilter
    ccbptTargetTable.EntityData.YangName = "ccbptTargetTable"
    ccbptTargetTable.EntityData.BundleName = "cisco_ios_xe"
    ccbptTargetTable.EntityData.ParentYangName = "CISCO-CBP-TARGET-MIB"
    ccbptTargetTable.EntityData.SegmentPath = "ccbptTargetTable"
    ccbptTargetTable.EntityData.AbsolutePath = "CISCO-CBP-TARGET-MIB:CISCO-CBP-TARGET-MIB/" + ccbptTargetTable.EntityData.SegmentPath
    ccbptTargetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccbptTargetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccbptTargetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccbptTargetTable.EntityData.Children = types.NewOrderedMap()
    ccbptTargetTable.EntityData.Children.Append("ccbptTargetEntry", types.YChild{"CcbptTargetEntry", nil})
    for i := range ccbptTargetTable.CcbptTargetEntry {
        ccbptTargetTable.EntityData.Children.Append(types.GetSegmentPath(ccbptTargetTable.CcbptTargetEntry[i]), types.YChild{"CcbptTargetEntry", ccbptTargetTable.CcbptTargetEntry[i]})
    }
    ccbptTargetTable.EntityData.Leafs = types.NewOrderedMap()

    ccbptTargetTable.EntityData.YListKeys = []string {}

    return &(ccbptTargetTable.EntityData)
}

// CISCOCBPTARGETMIB_CcbptTargetTable_CcbptTargetEntry
// Each entry describes a class-based policy attachment to a 
// particular target. 
//  
// The ccbptTargetType uniquely identifies the type of target
// in the attachment.  Additionally, the ccbptTargetId uniquely
// identifies the target in attachment and is of the format
// indicated by the ccbptTargetType.  The ccbptTargetDir 
// identifies the direction, relative to the ccbptTargetId, 
// to which the policy is attached.  The ccbptPolicySourceType
// identifies the source-type of the policy attached.  The 
// ccbptPolicyId uniquely identifies the policy within the scope
// of ccbptTargetType, ccbptTargetId, ccbptTargetDir, and 
// ccbptPolicySourceType.
// 
// A class-based policy attachment to a target can be created 
// through other network management interfaces (e.g., the local
// console), in which case the SNMP entity will automatically 
// create an entry in this table.
// 
// A class-based policy attachment to a target can be destroyed
// through other network management interfaces, in which case
// the SNMP entity will automatically destroy the corresponding
// entry in this table.
// 
// A class-based policy attachment to a target can be created,
// destroyed, and modified through the SNMP using 
// ccbptTargetStatus using the semantics described by the 
// RowStatus Textual Convention.  However, when creating a new
// class-based policy attachment to a target, the value of
// ccbptPolicyIdNext should be used to identify the new policy
// within the scope of the target type, identifier, direction,
// and policy-source type.
type CISCOCBPTARGETMIB_CcbptTargetTable_CcbptTargetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type of target for this class-based policy
    // attachment. This object identifies the format of the ccbptTargetId for this
    // entry. The type is CcbptTargetType.
    CcbptTargetType interface{}

    // This attribute is a key. The target identifier for this class-based policy
    // attachment. For decoding the ccbptTargetId refer to the ccbptTargetType
    // object and the CcbptTargetType description. The type is string with length:
    // 0..64.
    CcbptTargetId interface{}

    // This attribute is a key. The direction relative to the ccbptTargetId for
    // this class based policy attachment.  . The type is CcbptTargetDirection.
    CcbptTargetDir interface{}

    // This attribute is a key. The source-type of the class-based policy for this
    // target. The source-type refers to the source of the class-based policy
    // definition.  The intent of this object is to allow implementations to
    // distinguish between different MIBs defining policy-maps. . The type is
    // CcbptPolicySourceType.
    CcbptPolicySourceType interface{}

    // This attribute is a key. Unique identifier of this class-based policy
    // instance. The type is interface{} with range: 1..4294967295.
    CcbptPolicyId interface{}

    // The status of the policy attachment to this target.  The value for the
    // corresponding instance of each of the  following objects must be valid
    // before the attachment  can be activated:     -ccbptTargetStorageType    
    // -ccbptPolicyMap  Observe that no corresponding instance of any object in 
    // this table can be modified when the value of this object is 'active'. The
    // type is RowStatus.
    CcbptTargetStatus interface{}

    // This object indicates how the device stores the data  contained by the
    // conceptual row.  If an instance of this object has the value 'permanent',
    // then this MIB definition does not require the SNMP entity to allow the
    // instance of any object in the corresponding conceptual row to be writable
    // through the SNMP. The type is StorageType.
    CcbptTargetStorageType interface{}

    // Refers to the first accessible object in the policy-map definition table
    // used to manage policy-map information for policy-maps for the corresponding
    // ccbptPolicySourceType.  Specific MIB tables are not mentioned here as the
    // intent of this mapping is to allow for different implementations to  refer
    // to their supported class-based policy definition table without requiring
    // support of a specific MIB module. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    CcbptPolicyMap interface{}

    // Refers to the first accessible object in the policy  instance table used to
    // manage policy instance information  for policy-maps of this
    // ccbptPolicySourceType.  Specific MIB tables are not mentioned here as the
    // intent of this mapping is to allow for different implementations to  refer
    // to their supported class-based policy definition table without requiring
    // support of a specific MIB module. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    CcbptPolicyInstance interface{}

    // The value of sysUpTime for the last time that the corresponding
    // ccbptTargetStatus instance transitioned to the 'active' state.  . The type
    // is interface{} with range: 0..4294967295.
    CcbptPolicyAttachTime interface{}
}

func (ccbptTargetEntry *CISCOCBPTARGETMIB_CcbptTargetTable_CcbptTargetEntry) GetEntityData() *types.CommonEntityData {
    ccbptTargetEntry.EntityData.YFilter = ccbptTargetEntry.YFilter
    ccbptTargetEntry.EntityData.YangName = "ccbptTargetEntry"
    ccbptTargetEntry.EntityData.BundleName = "cisco_ios_xe"
    ccbptTargetEntry.EntityData.ParentYangName = "ccbptTargetTable"
    ccbptTargetEntry.EntityData.SegmentPath = "ccbptTargetEntry" + types.AddKeyToken(ccbptTargetEntry.CcbptTargetType, "ccbptTargetType") + types.AddKeyToken(ccbptTargetEntry.CcbptTargetId, "ccbptTargetId") + types.AddKeyToken(ccbptTargetEntry.CcbptTargetDir, "ccbptTargetDir") + types.AddKeyToken(ccbptTargetEntry.CcbptPolicySourceType, "ccbptPolicySourceType") + types.AddKeyToken(ccbptTargetEntry.CcbptPolicyId, "ccbptPolicyId")
    ccbptTargetEntry.EntityData.AbsolutePath = "CISCO-CBP-TARGET-MIB:CISCO-CBP-TARGET-MIB/ccbptTargetTable/" + ccbptTargetEntry.EntityData.SegmentPath
    ccbptTargetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccbptTargetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccbptTargetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccbptTargetEntry.EntityData.Children = types.NewOrderedMap()
    ccbptTargetEntry.EntityData.Leafs = types.NewOrderedMap()
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptTargetType", types.YLeaf{"CcbptTargetType", ccbptTargetEntry.CcbptTargetType})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptTargetId", types.YLeaf{"CcbptTargetId", ccbptTargetEntry.CcbptTargetId})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptTargetDir", types.YLeaf{"CcbptTargetDir", ccbptTargetEntry.CcbptTargetDir})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptPolicySourceType", types.YLeaf{"CcbptPolicySourceType", ccbptTargetEntry.CcbptPolicySourceType})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptPolicyId", types.YLeaf{"CcbptPolicyId", ccbptTargetEntry.CcbptPolicyId})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptTargetStatus", types.YLeaf{"CcbptTargetStatus", ccbptTargetEntry.CcbptTargetStatus})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptTargetStorageType", types.YLeaf{"CcbptTargetStorageType", ccbptTargetEntry.CcbptTargetStorageType})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptPolicyMap", types.YLeaf{"CcbptPolicyMap", ccbptTargetEntry.CcbptPolicyMap})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptPolicyInstance", types.YLeaf{"CcbptPolicyInstance", ccbptTargetEntry.CcbptPolicyInstance})
    ccbptTargetEntry.EntityData.Leafs.Append("ccbptPolicyAttachTime", types.YLeaf{"CcbptPolicyAttachTime", ccbptTargetEntry.CcbptPolicyAttachTime})

    ccbptTargetEntry.EntityData.YListKeys = []string {"CcbptTargetType", "CcbptTargetId", "CcbptTargetDir", "CcbptPolicySourceType", "CcbptPolicyId"}

    return &(ccbptTargetEntry.EntityData)
}

