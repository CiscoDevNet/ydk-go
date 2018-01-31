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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ccbpttargetattachcfg CISCOCBPTARGETMIB_Ccbpttargetattachcfg

    // This table describes the class-based policy attachments to to specific
    // targets.
    Ccbpttargettable CISCOCBPTARGETMIB_Ccbpttargettable
}

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetFilter() yfilter.YFilter { return cISCOCBPTARGETMIB.YFilter }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) SetFilter(yf yfilter.YFilter) { cISCOCBPTARGETMIB.YFilter = yf }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetGoName(yname string) string {
    if yname == "ccbptTargetAttachCfg" { return "Ccbpttargetattachcfg" }
    if yname == "ccbptTargetTable" { return "Ccbpttargettable" }
    return ""
}

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetSegmentPath() string {
    return "CISCO-CBP-TARGET-MIB:CISCO-CBP-TARGET-MIB"
}

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccbptTargetAttachCfg" {
        return &cISCOCBPTARGETMIB.Ccbpttargetattachcfg
    }
    if childYangName == "ccbptTargetTable" {
        return &cISCOCBPTARGETMIB.Ccbpttargettable
    }
    return nil
}

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ccbptTargetAttachCfg"] = &cISCOCBPTARGETMIB.Ccbpttargetattachcfg
    children["ccbptTargetTable"] = &cISCOCBPTARGETMIB.Ccbpttargettable
    return children
}

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetYangName() string { return "CISCO-CBP-TARGET-MIB" }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) SetParent(parent types.Entity) { cISCOCBPTARGETMIB.parent = parent }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetParent() types.Entity { return cISCOCBPTARGETMIB.parent }

func (cISCOCBPTARGETMIB *CISCOCBPTARGETMIB) GetParentYangName() string { return "CISCO-CBP-TARGET-MIB" }

// CISCOCBPTARGETMIB_Ccbpttargetattachcfg
type CISCOCBPTARGETMIB_Ccbpttargetattachcfg struct {
    parent types.Entity
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

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetFilter() yfilter.YFilter { return ccbpttargetattachcfg.YFilter }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) SetFilter(yf yfilter.YFilter) { ccbpttargetattachcfg.YFilter = yf }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetGoName(yname string) string {
    if yname == "ccbptPolicyIdNext" { return "Ccbptpolicyidnext" }
    if yname == "ccbptTargetTableLastChange" { return "Ccbpttargettablelastchange" }
    return ""
}

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetSegmentPath() string {
    return "ccbptTargetAttachCfg"
}

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccbptPolicyIdNext"] = ccbpttargetattachcfg.Ccbptpolicyidnext
    leafs["ccbptTargetTableLastChange"] = ccbpttargetattachcfg.Ccbpttargettablelastchange
    return leafs
}

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetBundleName() string { return "cisco_ios_xe" }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetYangName() string { return "ccbptTargetAttachCfg" }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) SetParent(parent types.Entity) { ccbpttargetattachcfg.parent = parent }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetParent() types.Entity { return ccbpttargetattachcfg.parent }

func (ccbpttargetattachcfg *CISCOCBPTARGETMIB_Ccbpttargetattachcfg) GetParentYangName() string { return "CISCO-CBP-TARGET-MIB" }

// CISCOCBPTARGETMIB_Ccbpttargettable
// This table describes the class-based policy attachments to
// to specific targets.
type CISCOCBPTARGETMIB_Ccbpttargettable struct {
    parent types.Entity
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

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetFilter() yfilter.YFilter { return ccbpttargettable.YFilter }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) SetFilter(yf yfilter.YFilter) { ccbpttargettable.YFilter = yf }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetGoName(yname string) string {
    if yname == "ccbptTargetEntry" { return "Ccbpttargetentry" }
    return ""
}

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetSegmentPath() string {
    return "ccbptTargetTable"
}

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ccbptTargetEntry" {
        for _, c := range ccbpttargettable.Ccbpttargetentry {
            if ccbpttargettable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry{}
        ccbpttargettable.Ccbpttargetentry = append(ccbpttargettable.Ccbpttargetentry, child)
        return &ccbpttargettable.Ccbpttargetentry[len(ccbpttargettable.Ccbpttargetentry)-1]
    }
    return nil
}

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ccbpttargettable.Ccbpttargetentry {
        children[ccbpttargettable.Ccbpttargetentry[i].GetSegmentPath()] = &ccbpttargettable.Ccbpttargetentry[i]
    }
    return children
}

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetBundleName() string { return "cisco_ios_xe" }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetYangName() string { return "ccbptTargetTable" }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) SetParent(parent types.Entity) { ccbpttargettable.parent = parent }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetParent() types.Entity { return ccbpttargettable.parent }

func (ccbpttargettable *CISCOCBPTARGETMIB_Ccbpttargettable) GetParentYangName() string { return "CISCO-CBP-TARGET-MIB" }

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
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Ccbptpolicymap interface{}

    // Refers to the first accessible object in the policy  instance table used to
    // manage policy instance information  for policy-maps of this
    // ccbptPolicySourceType.  Specific MIB tables are not mentioned here as the
    // intent of this mapping is to allow for different implementations to  refer
    // to their supported class-based policy definition table without requiring
    // support of a specific MIB module. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Ccbptpolicyinstance interface{}

    // The value of sysUpTime for the last time that the corresponding
    // ccbptTargetStatus instance transitioned to the 'active' state.  . The type
    // is interface{} with range: 0..4294967295.
    Ccbptpolicyattachtime interface{}
}

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetFilter() yfilter.YFilter { return ccbpttargetentry.YFilter }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) SetFilter(yf yfilter.YFilter) { ccbpttargetentry.YFilter = yf }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetGoName(yname string) string {
    if yname == "ccbptTargetType" { return "Ccbpttargettype" }
    if yname == "ccbptTargetId" { return "Ccbpttargetid" }
    if yname == "ccbptTargetDir" { return "Ccbpttargetdir" }
    if yname == "ccbptPolicySourceType" { return "Ccbptpolicysourcetype" }
    if yname == "ccbptPolicyId" { return "Ccbptpolicyid" }
    if yname == "ccbptTargetStatus" { return "Ccbpttargetstatus" }
    if yname == "ccbptTargetStorageType" { return "Ccbpttargetstoragetype" }
    if yname == "ccbptPolicyMap" { return "Ccbptpolicymap" }
    if yname == "ccbptPolicyInstance" { return "Ccbptpolicyinstance" }
    if yname == "ccbptPolicyAttachTime" { return "Ccbptpolicyattachtime" }
    return ""
}

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetSegmentPath() string {
    return "ccbptTargetEntry" + "[ccbptTargetType='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbpttargettype) + "']" + "[ccbptTargetId='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbpttargetid) + "']" + "[ccbptTargetDir='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbpttargetdir) + "']" + "[ccbptPolicySourceType='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbptpolicysourcetype) + "']" + "[ccbptPolicyId='" + fmt.Sprintf("%v", ccbpttargetentry.Ccbptpolicyid) + "']"
}

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccbptTargetType"] = ccbpttargetentry.Ccbpttargettype
    leafs["ccbptTargetId"] = ccbpttargetentry.Ccbpttargetid
    leafs["ccbptTargetDir"] = ccbpttargetentry.Ccbpttargetdir
    leafs["ccbptPolicySourceType"] = ccbpttargetentry.Ccbptpolicysourcetype
    leafs["ccbptPolicyId"] = ccbpttargetentry.Ccbptpolicyid
    leafs["ccbptTargetStatus"] = ccbpttargetentry.Ccbpttargetstatus
    leafs["ccbptTargetStorageType"] = ccbpttargetentry.Ccbpttargetstoragetype
    leafs["ccbptPolicyMap"] = ccbpttargetentry.Ccbptpolicymap
    leafs["ccbptPolicyInstance"] = ccbpttargetentry.Ccbptpolicyinstance
    leafs["ccbptPolicyAttachTime"] = ccbpttargetentry.Ccbptpolicyattachtime
    return leafs
}

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetBundleName() string { return "cisco_ios_xe" }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetYangName() string { return "ccbptTargetEntry" }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) SetParent(parent types.Entity) { ccbpttargetentry.parent = parent }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetParent() types.Entity { return ccbpttargetentry.parent }

func (ccbpttargetentry *CISCOCBPTARGETMIB_Ccbpttargettable_Ccbpttargetentry) GetParentYangName() string { return "ccbptTargetTable" }

