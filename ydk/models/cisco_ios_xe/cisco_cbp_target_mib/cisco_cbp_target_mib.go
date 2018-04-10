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

    
    Ccbpttargetattachcfg CISCOCBPTARGETMIB_Ccbpttargetattachcfg

    // This table describes the class-based policy attachments to to specific
    // targets.
    Ccbpttargettable CISCOCBPTARGETMIB_Ccbpttargettable
}

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetEntityData() *types.CommonEntityData {
    cISCOCBPTARGETMIB.EntityData.YFilter = cISCOCBPTARGETMIB.YFilter
    cISCOCBPTARGETMIB.EntityData.YangName = "CISCO-CBP-TARGET-MIB"
    cISCOCBPTARGETMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOCBPTARGETMIB.EntityData.ParentYangName = "CISCO-CBP-TARGET-MIB"
    cISCOCBPTARGETMIB.EntityData.SegmentPath = "CISCO-CBP-TARGET-MIB:CISCO-CBP-TARGET-MIB"
    cISCOCBPTARGETMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOCBPTARGETMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOCBPTARGETMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOCBPTARGETMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOCBPTARGETMIB.EntityData.Children["ccbptTargetAttachCfg"] = types.YChild{"Ccbpttargetattachcfg", &cISCOCBPTARGETMIB.Ccbpttargetattachcfg}
    cISCOCBPTARGETMIB.EntityData.Children["ccbptTargetTable"] = types.YChild{"Ccbpttargettable", &cISCOCBPTARGETMIB.Ccbpttargettable}
    cISCOCBPTARGETMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOCBPTARGETMIB.EntityData)
}

// CISCOCBPTARGETMIB_Ccbpttargetattachcfg
type CISCOCBPTARGETMIB_Ccbpttargetattachcfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates the next available value of  ccbptPolicyId that can
    // be used to create a new conceptual row in the ccbptTargetTable.  If no
    // available identifier exists, then this object will have the value '0'. The
    // type is interface{} with range: 0..4294967295.
    Ccbptpolicyidnext interface{}

    // The value of sysUpTime at the time of the last change to an entry in the
    // ccbptTargetTable. The type is interface{} with range: 0..4294967295.
    Ccbpttargettablelastchange interface{}
}

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetEntityData() *types.CommonEntityData {
    ccbpttargetattachcfg.EntityData.YFilter = ccbpttargetattachcfg.YFilter
    ccbpttargetattachcfg.EntityData.YangName = "ccbptTargetAttachCfg"
    ccbpttargetattachcfg.EntityData.BundleName = "cisco_ios_xe"
    ccbpttargetattachcfg.EntityData.ParentYangName = "CISCO-CBP-TARGET-MIB"
    ccbpttargetattachcfg.EntityData.SegmentPath = "ccbptTargetAttachCfg"
    ccbpttargetattachcfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccbpttargetattachcfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccbpttargetattachcfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccbpttargetattachcfg.EntityData.Children = make(map[string]types.YChild)
    ccbpttargetattachcfg.EntityData.Leafs = make(map[string]types.YLeaf)
    ccbpttargetattachcfg.EntityData.Leafs["ccbptPolicyIdNext"] = types.YLeaf{"Ccbptpolicyidnext", ccbpttargetattachcfg.Ccbptpolicyidnext}
    ccbpttargetattachcfg.EntityData.Leafs["ccbptTargetTableLastChange"] = types.YLeaf{"Ccbpttargettablelastchange", ccbpttargetattachcfg.Ccbpttargettablelastchange}
    return &(ccbpttargetattachcfg.EntityData)
}

// CISCOCBPTARGETMIB_Ccbpttargettable
// This table describes the class-based policy attachments to
// to specific targets.
type CISCOCBPTARGETMIB_Ccbpttargettable struct {
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
    // CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry.
    Ccbpttargetentry []CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry
}

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetEntityData() *types.CommonEntityData {
    ccbpttargettable.EntityData.YFilter = ccbpttargettable.YFilter
    ccbpttargettable.EntityData.YangName = "ccbptTargetTable"
    ccbpttargettable.EntityData.BundleName = "cisco_ios_xe"
    ccbpttargettable.EntityData.ParentYangName = "CISCO-CBP-TARGET-MIB"
    ccbpttargettable.EntityData.SegmentPath = "ccbptTargetTable"
    ccbpttargettable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccbpttargettable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccbpttargettable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccbpttargettable.EntityData.Children = make(map[string]types.YChild)
    ccbpttargettable.EntityData.Children["ccbptTargetEntry"] = types.YChild{"Ccbpttargetentry", nil}
    for i := range ccbpttargettable.Ccbpttargetentry {
        ccbpttargettable.EntityData.Children[types.GetSegmentPath(&ccbpttargettable.Ccbpttargetentry[i])] = types.YChild{"Ccbpttargetentry", &ccbpttargettable.Ccbpttargetentry[i]}
    }
    ccbpttargettable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ccbpttargettable.EntityData)
}

// CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry
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
type CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of target for this class-based policy
    // attachment. This object identifies the format of the ccbptTargetId for this
    // entry. The type is CcbptTargetType.
    Ccbpttargettype interface{}

    // This attribute is a key. The target identifier for this class-based policy
    // attachment. For decoding the ccbptTargetId refer to the ccbptTargetType
    // object and the CcbptTargetType description. The type is string with length:
    // 0..64.
    Ccbpttargetid interface{}

    // This attribute is a key. The direction relative to the ccbptTargetId for
    // this class based policy attachment.  . The type is CcbptTargetDirection.
    Ccbpttargetdir interface{}

    // This attribute is a key. The source-type of the class-based policy for this
    // target. The source-type refers to the source of the class-based policy
    // definition.  The intent of this object is to allow implementations to
    // distinguish between different MIBs defining policy-maps. . The type is
    // CcbptPolicySourceType.
    Ccbptpolicysourcetype interface{}

    // This attribute is a key. Unique identifier of this class-based policy
    // instance. The type is interface{} with range: 1..4294967295.
    Ccbptpolicyid interface{}

    // The status of the policy attachment to this target.  The value for the
    // corresponding instance of each of the  following objects must be valid
    // before the attachment  can be activated:     -ccbptTargetStorageType    
    // -ccbptPolicyMap  Observe that no corresponding instance of any object in 
    // this table can be modified when the value of this object is 'active'. The
    // type is RowStatus.
    Ccbpttargetstatus interface{}

    // This object indicates how the device stores the data  contained by the
    // conceptual row.  If an instance of this object has the value 'permanent',
    // then this MIB definition does not require the SNMP entity to allow the
    // instance of any object in the corresponding conceptual row to be writable
    // through the SNMP. The type is StorageType.
    Ccbpttargetstoragetype interface{}

    // Refers to the first accessible object in the policy-map definition table
    // used to manage policy-map information for policy-maps for the corresponding
    // ccbptPolicySourceType.  Specific MIB tables are not mentioned here as the
    // intent of this mapping is to allow for different implementations to  refer
    // to their supported class-based policy definition table without requiring
    // support of a specific MIB module. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Ccbptpolicymap interface{}

    // Refers to the first accessible object in the policy  instance table used to
    // manage policy instance information  for policy-maps of this
    // ccbptPolicySourceType.  Specific MIB tables are not mentioned here as the
    // intent of this mapping is to allow for different implementations to  refer
    // to their supported class-based policy definition table without requiring
    // support of a specific MIB module. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Ccbptpolicyinstance interface{}

    // The value of sysUpTime for the last time that the corresponding
    // ccbptTargetStatus instance transitioned to the 'active' state.  . The type
    // is interface{} with range: 0..4294967295.
    Ccbptpolicyattachtime interface{}
}

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetEntityData() *types.CommonEntityData {
    ccbpttargetentry.EntityData.YFilter = ccbpttargetentry.YFilter
    ccbpttargetentry.EntityData.YangName = "ccbptTargetEntry"
    ccbpttargetentry.EntityData.BundleName = "cisco_ios_xe"
    ccbpttargetentry.EntityData.ParentYangName = "ccbptTargetTable"
    ccbpttargetentry.EntityData.SegmentPath = "ccbptTargetEntry" + "[ccbptTargetType='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbpttargettype) + "']" + "[ccbptTargetId='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbpttargetid) + "']" + "[ccbptTargetDir='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbpttargetdir) + "']" + "[ccbptPolicySourceType='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbptpolicysourcetype) + "']" + "[ccbptPolicyId='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbptpolicyid) + "']"
    ccbpttargetentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ccbpttargetentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ccbpttargetentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ccbpttargetentry.EntityData.Children = make(map[string]types.YChild)
    ccbpttargetentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ccbpttargetentry.EntityData.Leafs["ccbptTargetType"] = types.YLeaf{"Ccbpttargettype", ccbpttargetentry.Ccbpttargettype}
    ccbpttargetentry.EntityData.Leafs["ccbptTargetId"] = types.YLeaf{"Ccbpttargetid", ccbpttargetentry.Ccbpttargetid}
    ccbpttargetentry.EntityData.Leafs["ccbptTargetDir"] = types.YLeaf{"Ccbpttargetdir", ccbpttargetentry.Ccbpttargetdir}
    ccbpttargetentry.EntityData.Leafs["ccbptPolicySourceType"] = types.YLeaf{"Ccbptpolicysourcetype", ccbpttargetentry.Ccbptpolicysourcetype}
    ccbpttargetentry.EntityData.Leafs["ccbptPolicyId"] = types.YLeaf{"Ccbptpolicyid", ccbpttargetentry.Ccbptpolicyid}
    ccbpttargetentry.EntityData.Leafs["ccbptTargetStatus"] = types.YLeaf{"Ccbpttargetstatus", ccbpttargetentry.Ccbpttargetstatus}
    ccbpttargetentry.EntityData.Leafs["ccbptTargetStorageType"] = types.YLeaf{"Ccbpttargetstoragetype", ccbpttargetentry.Ccbpttargetstoragetype}
    ccbpttargetentry.EntityData.Leafs["ccbptPolicyMap"] = types.YLeaf{"Ccbptpolicymap", ccbpttargetentry.Ccbptpolicymap}
    ccbpttargetentry.EntityData.Leafs["ccbptPolicyInstance"] = types.YLeaf{"Ccbptpolicyinstance", ccbpttargetentry.Ccbptpolicyinstance}
    ccbpttargetentry.EntityData.Leafs["ccbptPolicyAttachTime"] = types.YLeaf{"Ccbptpolicyattachtime", ccbpttargetentry.Ccbptpolicyattachtime}
    return &(ccbpttargetentry.EntityData)
}

