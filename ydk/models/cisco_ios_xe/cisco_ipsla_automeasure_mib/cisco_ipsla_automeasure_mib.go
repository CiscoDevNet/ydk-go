// This module defines the MIB for IP SLA Automation. IP SLA
// Automation consists of the following:
// 1. Use of grouping - Group is an aggregation of operations 
//    sharing the same type, for example UDP jitter type, with 
//    common characteristics like frequency, interval etc.  
//    Groups are formed by policies dictated either per customer, 
//    or by service level or any other requirements.  So, for 
//    example, there could be separate groups for customers named 
//    Customer A, Customer B etc, or service levels named 
//    Gold-service, Silver-service etc.  A single group will contain 
//    one and only one IP SLA operation definition or reference a 
//    single template.
// 2. Use of templates - It has been observed that operations can be 
//    configured quickly if the variants such as IP address and 
//    ports can be configured separately and then combined with a 
//    template  consisting of invariants.  This allows for re-use of 
//    the invariant template with various combinations of destination 
//    addresses and ports, the benefits for which are multiplied when 
//    considering groupings.
// 3. Auto operations - With this feature the software will try to 
//    automatically kickstart operations by making intelligent 
//    assumptions.  For example, if no specific operation is referenced 
//    by the group configuration then an ICMP jitter operation is 
//    assumed by default.
package cisco_ipsla_automeasure_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ipsla_automeasure_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IPSLA-AUTOMEASURE-MIB CISCO-IPSLA-AUTOMEASURE-MIB}", reflect.TypeOf(CISCOIPSLAAUTOMEASUREMIB{}))
    ydk.RegisterEntity("CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB", reflect.TypeOf(CISCOIPSLAAUTOMEASUREMIB{}))
}

// CISCOIPSLAAUTOMEASUREMIB
type CISCOIPSLAAUTOMEASUREMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table that contains IP SLA auto measure group definitions.
    Cipslaautogrouptable CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable

    // A table contains the list of destination IP addresses and ports associated
    // to the auto measure group destination name.
    Cipslaautogroupdesttable CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable

    // A table that contains reaction configurations for templates. Each
    // conceptual row in cipslaReactTable corresponds  to a reaction configured
    // for one template.  Each template can have multiple reactions and hence
    // there can be  multiple rows for a particular template. Different template
    // types can have different reactions. The reaction type is  specified as
    // cipslaReactVar based upon template type as some  reaction types are
    // applicable just for specific template types.
    Cipslareacttable CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable

    // A table of group scheduling definitions.
    Cipslaautogroupschedtable CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable
}

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetFilter() yfilter.YFilter { return cISCOIPSLAAUTOMEASUREMIB.YFilter }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) SetFilter(yf yfilter.YFilter) { cISCOIPSLAAUTOMEASUREMIB.YFilter = yf }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetGoName(yname string) string {
    if yname == "cipslaAutoGroupTable" { return "Cipslaautogrouptable" }
    if yname == "cipslaAutoGroupDestTable" { return "Cipslaautogroupdesttable" }
    if yname == "cipslaReactTable" { return "Cipslareacttable" }
    if yname == "cipslaAutoGroupSchedTable" { return "Cipslaautogroupschedtable" }
    return ""
}

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetSegmentPath() string {
    return "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB"
}

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaAutoGroupTable" {
        return &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogrouptable
    }
    if childYangName == "cipslaAutoGroupDestTable" {
        return &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogroupdesttable
    }
    if childYangName == "cipslaReactTable" {
        return &cISCOIPSLAAUTOMEASUREMIB.Cipslareacttable
    }
    if childYangName == "cipslaAutoGroupSchedTable" {
        return &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogroupschedtable
    }
    return nil
}

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cipslaAutoGroupTable"] = &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogrouptable
    children["cipslaAutoGroupDestTable"] = &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogroupdesttable
    children["cipslaReactTable"] = &cISCOIPSLAAUTOMEASUREMIB.Cipslareacttable
    children["cipslaAutoGroupSchedTable"] = &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogroupschedtable
    return children
}

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetYangName() string { return "CISCO-IPSLA-AUTOMEASURE-MIB" }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) SetParent(parent types.Entity) { cISCOIPSLAAUTOMEASUREMIB.parent = parent }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetParent() types.Entity { return cISCOIPSLAAUTOMEASUREMIB.parent }

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetParentYangName() string { return "CISCO-IPSLA-AUTOMEASURE-MIB" }

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable
// A table that contains IP SLA auto measure group definitions.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the configurations for a particular auto measure group.
    // The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry.
    Cipslaautogroupentry []CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry
}

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetFilter() yfilter.YFilter { return cipslaautogrouptable.YFilter }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) SetFilter(yf yfilter.YFilter) { cipslaautogrouptable.YFilter = yf }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetGoName(yname string) string {
    if yname == "cipslaAutoGroupEntry" { return "Cipslaautogroupentry" }
    return ""
}

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetSegmentPath() string {
    return "cipslaAutoGroupTable"
}

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaAutoGroupEntry" {
        for _, c := range cipslaautogrouptable.Cipslaautogroupentry {
            if cipslaautogrouptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry{}
        cipslaautogrouptable.Cipslaautogroupentry = append(cipslaautogrouptable.Cipslaautogroupentry, child)
        return &cipslaautogrouptable.Cipslaautogroupentry[len(cipslaautogrouptable.Cipslaautogroupentry)-1]
    }
    return nil
}

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslaautogrouptable.Cipslaautogroupentry {
        children[cipslaautogrouptable.Cipslaautogroupentry[i].GetSegmentPath()] = &cipslaautogrouptable.Cipslaautogroupentry[i]
    }
    return children
}

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetYangName() string { return "cipslaAutoGroupTable" }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) SetParent(parent types.Entity) { cipslaautogrouptable.parent = parent }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetParent() types.Entity { return cipslaautogrouptable.parent }

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetParentYangName() string { return "CISCO-IPSLA-AUTOMEASURE-MIB" }

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry
// An entry containing the configurations for a particular
// auto measure group.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A group name which is used by a management
    // application to identify the group. The type is string with length: 1..64.
    Cipslaautogroupname interface{}

    // This field is used to provide description for the group. The type is string
    // with length: 0..128.
    Cipslaautogroupdescription interface{}

    // This object refers to the cipslaAutoGroupDestName in
    // cipslaAutoGroupDestTable. If the name entered  is not present in
    // cipslaAutoGroupDestTable, then when  group is scheduled, no ip sla
    // operations will be created. The type is string with length: 0..64.
    Cipslaautogroupdestinationname interface{}

    // This object represents the destination port number for auto discovery use.
    // The type is interface{} with range: 0..65535.
    Cipslaautogroupaddestport interface{}

    // A string which is used by a management application to identify the template
    // which is associated with the group. Depends on cipslaAutoGroupOperType,
    // this object refers to cipslaIcmpEchoTmplName in cipslaIcmpEchoTmplTable, or
    // cipslaUdpEchoTmplName in cipslaUdpEchoTmplTable, or cipslaTcpConnTmplName
    // in cipslaTcpConnTmplTable, or cipslaIcmpJitterTmplName in
    // cipslaIcmpJitterTmplTable, or ciscoIpSlaUdpJitterTmplName in
    // ciscoIpSlaUdpJitterTmplTable. The type is string with length: 0..64.
    Cipslaautogroupopertemplatename interface{}

    // This object refers to the cipslaAutoGroupSchedId in
    // cipslaAutoGroupSchedTable, and is used to schedule  this group. The type is
    // string with length: 0..64.
    Cipslaautogroupschedulerid interface{}

    // When this object is set to true, QoS is enabled for this group and this
    // group is linked to policy map. The  restriction is that after QoS is
    // enabled, it can not be  disabled for this group. The type is bool.
    Cipslaautogroupqosenable interface{}

    // This object specifies the type of IP SLA operation. When operation type is
    // not ICMP jitter, then  cipslaAutoGroupOperTemplateName must be specified.
    // The type is IpSlaOperType.
    Cipslaautogroupopertype interface{}

    // When this object is set to true, destination IP address is populated
    // through auto-discovery. The type is bool.
    Cipslaautogroupdestipadenable interface{}

    // This object specifies number of measurement retries to be attempted for the
    // discovered end point after the  connection to the end point is broken. If
    // there is no  re-registration message received, the end point will be  in
    // inactive state.  When the value of cipslaAutoGroupDestIPADEnable is 
    // 'false', the value of this object has no effect. The type is interface{}
    // with range: 1..65536.
    Cipslaautogroupadmeasureretry interface{}

    // This object represents the ageout time for the discovered end point.  If
    // the end point becomes inactive for the period  of ageout time, the end
    // point will be removed from the  discovered end point list.  When the value
    // of cipslaAutoGroupDestIPADEnable is  'false', the value of this object has
    // no effect. The type is interface{} with range: 0..65536. Units are seconds.
    Cipslaautogroupaddestipageout interface{}

    // The storage type of this conceptual row. The type is StorageType.
    Cipslaautogroupstoragetype interface{}

    // The status of the conceptual group control row.  When the status is active,
    // the other writable objects may be modified unless the scheduler with name 
    // specified by cipslaAutoGroupSchedulerId is scheduled. The type is
    // RowStatus.
    Cipslaautogrouprowstatus interface{}
}

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetFilter() yfilter.YFilter { return cipslaautogroupentry.YFilter }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) SetFilter(yf yfilter.YFilter) { cipslaautogroupentry.YFilter = yf }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetGoName(yname string) string {
    if yname == "cipslaAutoGroupName" { return "Cipslaautogroupname" }
    if yname == "cipslaAutoGroupDescription" { return "Cipslaautogroupdescription" }
    if yname == "cipslaAutoGroupDestinationName" { return "Cipslaautogroupdestinationname" }
    if yname == "cipslaAutoGroupADDestPort" { return "Cipslaautogroupaddestport" }
    if yname == "cipslaAutoGroupOperTemplateName" { return "Cipslaautogroupopertemplatename" }
    if yname == "cipslaAutoGroupSchedulerId" { return "Cipslaautogroupschedulerid" }
    if yname == "cipslaAutoGroupQoSEnable" { return "Cipslaautogroupqosenable" }
    if yname == "cipslaAutoGroupOperType" { return "Cipslaautogroupopertype" }
    if yname == "cipslaAutoGroupDestIPADEnable" { return "Cipslaautogroupdestipadenable" }
    if yname == "cipslaAutoGroupADMeasureRetry" { return "Cipslaautogroupadmeasureretry" }
    if yname == "cipslaAutoGroupADDestIPAgeout" { return "Cipslaautogroupaddestipageout" }
    if yname == "cipslaAutoGroupStorageType" { return "Cipslaautogroupstoragetype" }
    if yname == "cipslaAutoGroupRowStatus" { return "Cipslaautogrouprowstatus" }
    return ""
}

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetSegmentPath() string {
    return "cipslaAutoGroupEntry" + "[cipslaAutoGroupName='" + fmt.Sprintf("%v", cipslaautogroupentry.Cipslaautogroupname) + "']"
}

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaAutoGroupName"] = cipslaautogroupentry.Cipslaautogroupname
    leafs["cipslaAutoGroupDescription"] = cipslaautogroupentry.Cipslaautogroupdescription
    leafs["cipslaAutoGroupDestinationName"] = cipslaautogroupentry.Cipslaautogroupdestinationname
    leafs["cipslaAutoGroupADDestPort"] = cipslaautogroupentry.Cipslaautogroupaddestport
    leafs["cipslaAutoGroupOperTemplateName"] = cipslaautogroupentry.Cipslaautogroupopertemplatename
    leafs["cipslaAutoGroupSchedulerId"] = cipslaautogroupentry.Cipslaautogroupschedulerid
    leafs["cipslaAutoGroupQoSEnable"] = cipslaautogroupentry.Cipslaautogroupqosenable
    leafs["cipslaAutoGroupOperType"] = cipslaautogroupentry.Cipslaautogroupopertype
    leafs["cipslaAutoGroupDestIPADEnable"] = cipslaautogroupentry.Cipslaautogroupdestipadenable
    leafs["cipslaAutoGroupADMeasureRetry"] = cipslaautogroupentry.Cipslaautogroupadmeasureretry
    leafs["cipslaAutoGroupADDestIPAgeout"] = cipslaautogroupentry.Cipslaautogroupaddestipageout
    leafs["cipslaAutoGroupStorageType"] = cipslaautogroupentry.Cipslaautogroupstoragetype
    leafs["cipslaAutoGroupRowStatus"] = cipslaautogroupentry.Cipslaautogrouprowstatus
    return leafs
}

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetYangName() string { return "cipslaAutoGroupEntry" }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) SetParent(parent types.Entity) { cipslaautogroupentry.parent = parent }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetParent() types.Entity { return cipslaautogroupentry.parent }

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetParentYangName() string { return "cipslaAutoGroupTable" }

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable
// A table contains the list of destination IP
// addresses and ports associated to the auto measure
// group destination name.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry containing the destination IP addresses and port configurations
    // associated to auto measure group destination name. The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry.
    Cipslaautogroupdestentry []CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry
}

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetFilter() yfilter.YFilter { return cipslaautogroupdesttable.YFilter }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) SetFilter(yf yfilter.YFilter) { cipslaautogroupdesttable.YFilter = yf }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetGoName(yname string) string {
    if yname == "cipslaAutoGroupDestEntry" { return "Cipslaautogroupdestentry" }
    return ""
}

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetSegmentPath() string {
    return "cipslaAutoGroupDestTable"
}

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaAutoGroupDestEntry" {
        for _, c := range cipslaautogroupdesttable.Cipslaautogroupdestentry {
            if cipslaautogroupdesttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry{}
        cipslaautogroupdesttable.Cipslaautogroupdestentry = append(cipslaautogroupdesttable.Cipslaautogroupdestentry, child)
        return &cipslaautogroupdesttable.Cipslaautogroupdestentry[len(cipslaautogroupdesttable.Cipslaautogroupdestentry)-1]
    }
    return nil
}

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslaautogroupdesttable.Cipslaautogroupdestentry {
        children[cipslaautogroupdesttable.Cipslaautogroupdestentry[i].GetSegmentPath()] = &cipslaautogroupdesttable.Cipslaautogroupdestentry[i]
    }
    return children
}

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetYangName() string { return "cipslaAutoGroupDestTable" }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) SetParent(parent types.Entity) { cipslaautogroupdesttable.parent = parent }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetParent() types.Entity { return cipslaautogroupdesttable.parent }

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetParentYangName() string { return "CISCO-IPSLA-AUTOMEASURE-MIB" }

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry
// An entry containing the destination IP addresses
// and port configurations associated to auto measure
// group destination name.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is the name for an auto measure group
    // destination. The type is string with length: 1..64.
    Cipslaautogroupdestname interface{}

    // This attribute is a key. The type of the internet address of a destination
    // for an auto measure group. The type is InetAddressType.
    Cipslaautogroupdestipaddrtype interface{}

    // This attribute is a key. The internet address of a destination for an auto
    // measure group. The type of this address is determined by the value of
    // cipslaAutoGroupDestIpAddrType. The type is string with length: 0..255.
    Cipslaautogroupdestipaddr interface{}

    // This attribute is a key. This object represents the destination port
    // number. For ICMP echo and ICMP jitter, the suggested value is  '0'. The
    // type is interface{} with range: 0..65535.
    Cipslaautogroupdestport interface{}

    // The storage type of this conceptual row.  By default the entry will be
    // saved into non-volatile memory. The type is StorageType.
    Cipslaautogroupdeststoragetype interface{}

    // The status of the conceptual destination table control row. No other
    // objects in this row need to be set before this object can become active. 
    // During 'destroy', when cipslaAutoGroupDestIpAddr is specified  as '0.0.0.0'
    // and cipslaAutoGroupDestPort is specified as '0',  then all the rows with
    // same cipslaAutoGroupDestName will be  deleted. The type is RowStatus.
    Cipslaautogroupdestrowstatus interface{}
}

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetFilter() yfilter.YFilter { return cipslaautogroupdestentry.YFilter }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) SetFilter(yf yfilter.YFilter) { cipslaautogroupdestentry.YFilter = yf }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetGoName(yname string) string {
    if yname == "cipslaAutoGroupDestName" { return "Cipslaautogroupdestname" }
    if yname == "cipslaAutoGroupDestIpAddrType" { return "Cipslaautogroupdestipaddrtype" }
    if yname == "cipslaAutoGroupDestIpAddr" { return "Cipslaautogroupdestipaddr" }
    if yname == "cipslaAutoGroupDestPort" { return "Cipslaautogroupdestport" }
    if yname == "cipslaAutoGroupDestStorageType" { return "Cipslaautogroupdeststoragetype" }
    if yname == "cipslaAutoGroupDestRowStatus" { return "Cipslaautogroupdestrowstatus" }
    return ""
}

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetSegmentPath() string {
    return "cipslaAutoGroupDestEntry" + "[cipslaAutoGroupDestName='" + fmt.Sprintf("%v", cipslaautogroupdestentry.Cipslaautogroupdestname) + "']" + "[cipslaAutoGroupDestIpAddrType='" + fmt.Sprintf("%v", cipslaautogroupdestentry.Cipslaautogroupdestipaddrtype) + "']" + "[cipslaAutoGroupDestIpAddr='" + fmt.Sprintf("%v", cipslaautogroupdestentry.Cipslaautogroupdestipaddr) + "']" + "[cipslaAutoGroupDestPort='" + fmt.Sprintf("%v", cipslaautogroupdestentry.Cipslaautogroupdestport) + "']"
}

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaAutoGroupDestName"] = cipslaautogroupdestentry.Cipslaautogroupdestname
    leafs["cipslaAutoGroupDestIpAddrType"] = cipslaautogroupdestentry.Cipslaautogroupdestipaddrtype
    leafs["cipslaAutoGroupDestIpAddr"] = cipslaautogroupdestentry.Cipslaautogroupdestipaddr
    leafs["cipslaAutoGroupDestPort"] = cipslaautogroupdestentry.Cipslaautogroupdestport
    leafs["cipslaAutoGroupDestStorageType"] = cipslaautogroupdestentry.Cipslaautogroupdeststoragetype
    leafs["cipslaAutoGroupDestRowStatus"] = cipslaautogroupdestentry.Cipslaautogroupdestrowstatus
    return leafs
}

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetYangName() string { return "cipslaAutoGroupDestEntry" }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) SetParent(parent types.Entity) { cipslaautogroupdestentry.parent = parent }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetParent() types.Entity { return cipslaautogroupdestentry.parent }

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetParentYangName() string { return "cipslaAutoGroupDestTable" }

// CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable
// A table that contains reaction configurations for templates.
// Each conceptual row in cipslaReactTable corresponds 
// to a reaction configured for one template.
// 
// Each template can have multiple reactions and hence there can be 
// multiple rows for a particular template. Different template
// types can have different reactions. The reaction type is 
// specified as cipslaReactVar based upon template type as some 
// reaction types are applicable just for specific template types.
type CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual reaction configuration
    // control row. The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry.
    Cipslareactentry []CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry
}

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetFilter() yfilter.YFilter { return cipslareacttable.YFilter }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) SetFilter(yf yfilter.YFilter) { cipslareacttable.YFilter = yf }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetGoName(yname string) string {
    if yname == "cipslaReactEntry" { return "Cipslareactentry" }
    return ""
}

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetSegmentPath() string {
    return "cipslaReactTable"
}

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaReactEntry" {
        for _, c := range cipslareacttable.Cipslareactentry {
            if cipslareacttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry{}
        cipslareacttable.Cipslareactentry = append(cipslareacttable.Cipslareactentry, child)
        return &cipslareacttable.Cipslareactentry[len(cipslareacttable.Cipslareactentry)-1]
    }
    return nil
}

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslareacttable.Cipslareactentry {
        children[cipslareacttable.Cipslareactentry[i].GetSegmentPath()] = &cipslareacttable.Cipslareactentry[i]
    }
    return children
}

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetYangName() string { return "cipslaReactTable" }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) SetParent(parent types.Entity) { cipslareacttable.parent = parent }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetParent() types.Entity { return cipslareacttable.parent }

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetParentYangName() string { return "CISCO-IPSLA-AUTOMEASURE-MIB" }

// CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry
// A base list of objects that define a conceptual reaction
// configuration control row.
type CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is IpSlaOperType. Refers to
    // cisco_ipsla_automeasure_mib.CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry_Cipslaautogroupopertype
    Cipslaautogroupopertype interface{}

    // This attribute is a key. This object along with cipslaAutoGroupOperType and
    // cipslaAutoGroupOperTemplateName identifies a particular
    // reaction-configuration for one IP SLA  template.  This number is persistent
    // across reboots. The type is interface{} with range: 1..2147483647.
    Cipslareactconfigindex interface{}

    // This attribute is a key. The type is string with length: 0..64. Refers to
    // cisco_ipsla_automeasure_mib.CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry_Cipslaautogroupopertemplatename
    Cipslaautogroupopertemplatename interface{}

    // This object specifies the type of reaction configured for an IP SLA
    // template. Default value is 'rtt' for ICMP echo, UDP echo and TCP connect.
    // Default value is 'jitterAvg' for UDP jitter and  ICMP jitter.  The reaction
    // types 'rtt', 'timeout', 'connectionLoss' and 'verifyError' can be
    // configured for all template types.  The reaction types 'jitterSDAvg',
    // 'jitterDSAvg', 'jitterAvg',  'packetLateArrival', 'packetOutOfSequence', 
    // 'maxOfPositiveSD', 'maxOfNegativeSD', 'maxOfPositiveDS' 'maxOfNegativeDS',
    // 'mos' and 'icpif' can be configured for  UDP jitter and ICMP jitter types
    // only.  The reaction types 'packetLossDS', 'packetLossSD' and  'packetMIA'
    // can be configured for UDP jitter type only.  The reaction types
    // 'successivePacketLoss', 'maxOfLatencyDS',  'maxOfLatencySD',
    // 'latencyDSAvg', 'latencySDAvg' and  'packetLoss' can be configured for ICMP
    // jitter type only. The type is IpSlaReactVar.
    Cipslareactvar interface{}

    // This object specifies the conditions under which a notification ( trap ) is
    // sent. The rttMonReactOccurred object defined in rttMonReactTable in
    // CISCO-RTTMON-MIB will change accordingly:  never(1)       -
    // rttMonReactOccurred is never set  immediate(2)   - rttMonReactOccurred is
    // set to 'true' when the                  value of parameter for which
    // reaction is                  configured ( e.g rtt, jitterAvg, packetLossSD,
    // mos etc ) violates the threshold.                  Conversely,
    // rttMonReactOccurred is set to 'false'                  when the parameter (
    // e.g rtt, jitterAvg,                  packetLossSD, mos etc ) is below the
    // threshold                  limits.  consecutive(3) - rttMonReactOccurred is
    // set to true when the value                  of parameter for which reaction
    // is configured                  ( e.g rtt, jitterAvg, packetLossSD, mos etc
    // )                  violates the threshold for configured consecutive       
    // times.                  Conversely, rttMonReactOccurred is set to false    
    // when the value of parameter ( e.g rtt, jitterAvg                 
    // packetLossSD, mos etc ) is below the threshold                  limits for
    // the same number of consecutive                  operations.  xOfy(4)       
    // - rttMonReactOccurred is set to true when x                  ( as specified
    // by cipslaReactThresholdCountX )                  out of the last y ( as
    // specified by                  cipslaReacthresholdCountY ) times the value
    // of                  parameter for which the reaction is configured         
    // ( e.g rtt, jitterAvg, packetLossSD, mos etc )                  violates the
    // threshold.                  Conversely, it is set to false when x, out of
    // the                  last y times the value of parameter                  (
    // e.g rtt, jitterAvg, packetLossSD, mos ) is                  below the
    // threshold limits.                  NOTE: If x > y, this will never         
    // generate a reaction.  average(5)    - rttMonReactOccurred is set to true
    // when the                 average ( cipslaReactThresholdCountX times )      
    // value of parameter for which reaction is                  configured ( e.g
    // rtt, jitterAvg, packetLossSD,                 mos etc ) violates the
    // threshold condition.                 Conversely, it is set to false when
    // the                 average value of parameter ( e.g rtt, jitterAvg,       
    // packetLossSD, mos etc ) is below the threshold                 limits.  If
    // this value is changed by a management station, rttMonReactOccurred is set
    // to false, but no reaction is generated if the prior value of
    // rttMonReactOccurred was true. The type is Cipslareactthresholdtype.
    Cipslareactthresholdtype interface{}

    // Specifies what type, if any, of reaction to generate if one of the watched
    // (reaction-configuration ) conditions is satisfied:  none(1)               
    // - no reaction is generated notificationOnly(2)    - a notification is
    // generated. The type is Cipslareactactiontype.
    Cipslareactactiontype interface{}

    // This object defines the higher threshold limit. If the value ( e.g rtt,
    // jitterAvg, packetLossSD etc ) rises above this limit and if the condition
    // specified in cipslaReactThresholdType is satisfied, a notification is 
    // generated.  Default value of cipslaReactThresholdRising for    'rtt' is
    // 5000    'jitterAvg' is 100.    'jitterSDAvg' is 100.    'jitterDSAvg' 100. 
    // 'packetLossSD' is 10000.    'packetLossDS' is 10000.    'mos' is 500.   
    // 'icpif' is 93.    'packetMIA' is 10000.    'packetLateArrival' is 10000.   
    // 'packetOutOfSequence' is 10000.    'maxOfPositiveSD' is 10000.   
    // 'maxOfNegativeSD' is 10000.    'maxOfPositiveDS' is 10000.   
    // 'maxOfNegativeDS' is 10000.    'successivePacketLoss' is 1000.   
    // 'maxOfLatencyDS' is 5000.    'maxOfLatencySD' is 5000.    'latencyDSAvg' is
    // 5000.    'latencySDAvg' is 5000.    'packetLoss' is 10000.  This object is
    // not applicable if the cipslaReactVar is 'timeout', 'connectionLoss' or
    // 'verifyError'. For 'timeout', 'connectionLoss' and 'verifyError' default
    // value of  cipslaReactThresholdRising will be 0. The type is interface{}
    // with range: 0..4294967295.
    Cipslareactthresholdrising interface{}

    // This object defines a lower threshold limit. If the value ( e.g rtt,
    // jitterAvg, packetLossSD etc ) falls below this limit and if the condition
    // specified in cipslaReactThresholdType is satisfied, a notification  is
    // generated. This object value can not bigger than cipslaReactThresholdRising
    // value.  Default value of cipslaReactThresholdFalling    'rtt' is 3000   
    // 'jitterAvg' is 100.    'jitterSDAvg' is 100.    'jitterDSAvg' 100.   
    // 'packetLossSD' is 10000.    'packetLossDS' is 10000.    'mos' is 500.   
    // 'icpif' is 93.    'packetMIA' is 10000.    'packetLateArrival' is 10000.   
    // 'packetOutOfSequence' is 10000.    'maxOfPositiveSD' is 10000.   
    // 'maxOfNegativeSD' is 10000.    'maxOfPositiveDS' is 10000.   
    // 'maxOfNegativeDS' is 10000.    'successivePacketLoss' is 1000.   
    // 'maxOfLatencyDS' is 3000.    'maxOfLatencySD' is 3000.    'latencyDSAvg' is
    // 3000.    'latencySDAvg' is 3000.    'packetLoss' is 10000.  This object is
    // not applicable if the cipslaReactVar is 'timeout', 'connectionLoss' or
    // 'verifyError'. For 'timeout', 'connectionLoss' and 'verifyError', default
    // value of cipslaReactThresholdFalling will be 0. The type is interface{}
    // with range: 0..4294967295.
    Cipslareactthresholdfalling interface{}

    // If cipslaReactThresholdType value is 'xOfy', this object defines the 'x'
    // value.  If cipslaReactThresholdType value is 'consecutive' this object
    // defines the number of consecutive occurrences that needs threshold
    // violation before setting  cipslaReactOccurred as true.  If
    // cipslaReactThresholdType value is 'average' this object defines the number
    // of samples that needs be considered for calculating average.  This object
    // has no meaning if cipslaReactThresholdType has value of 'never' and
    // 'immediate'. The type is interface{} with range: 1..16.
    Cipslareactthresholdcountx interface{}

    // This object defines the 'y' value of the xOfy condition if
    // cipslaReactThresholdType is 'xOfy'. The default for the  'y' value is 5. 
    // For other values of cipslaReactThresholdType, this object is not
    // applicable. The type is interface{} with range: 1..16.
    Cipslareactthresholdcounty interface{}

    // The storage type of this conceptual row.  By default the entry will be
    // saved into non-volatile memory. The type is StorageType.
    Cipslareactstoragetype interface{}

    // This objects indicates the status of the conceptual Reaction Control Row.  
    // When this object moves to active state, the conceptual row  is monitored
    // and notifications are generated when threshold  violation takes place.  In
    // order for this object to become active cipslaReactVar must be defined. All
    // other objects assume default values.  When the  status is active, the
    // following objects in that row can be  modified.  cipslaReactThresholdType, 
    // cipslaReactActionType,  cipslaReactThresholdRising, 
    // cipslaReactThresholdFalling,  cipslaReactThresholdCountX, 
    // cipslaReactThresholdCountY,  cipslaReactStorageType  This object can be set
    // to 'destroy' from any value at any time. When this object is set to
    // 'destroy' no reaction configuration would exist. The reaction configuration
    // for the template is  removed. The type is RowStatus.
    Cipslareactrowstatus interface{}
}

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetFilter() yfilter.YFilter { return cipslareactentry.YFilter }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) SetFilter(yf yfilter.YFilter) { cipslareactentry.YFilter = yf }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetGoName(yname string) string {
    if yname == "cipslaAutoGroupOperType" { return "Cipslaautogroupopertype" }
    if yname == "cipslaReactConfigIndex" { return "Cipslareactconfigindex" }
    if yname == "cipslaAutoGroupOperTemplateName" { return "Cipslaautogroupopertemplatename" }
    if yname == "cipslaReactVar" { return "Cipslareactvar" }
    if yname == "cipslaReactThresholdType" { return "Cipslareactthresholdtype" }
    if yname == "cipslaReactActionType" { return "Cipslareactactiontype" }
    if yname == "cipslaReactThresholdRising" { return "Cipslareactthresholdrising" }
    if yname == "cipslaReactThresholdFalling" { return "Cipslareactthresholdfalling" }
    if yname == "cipslaReactThresholdCountX" { return "Cipslareactthresholdcountx" }
    if yname == "cipslaReactThresholdCountY" { return "Cipslareactthresholdcounty" }
    if yname == "cipslaReactStorageType" { return "Cipslareactstoragetype" }
    if yname == "cipslaReactRowStatus" { return "Cipslareactrowstatus" }
    return ""
}

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetSegmentPath() string {
    return "cipslaReactEntry" + "[cipslaAutoGroupOperType='" + fmt.Sprintf("%v", cipslareactentry.Cipslaautogroupopertype) + "']" + "[cipslaReactConfigIndex='" + fmt.Sprintf("%v", cipslareactentry.Cipslareactconfigindex) + "']" + "[cipslaAutoGroupOperTemplateName='" + fmt.Sprintf("%v", cipslareactentry.Cipslaautogroupopertemplatename) + "']"
}

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaAutoGroupOperType"] = cipslareactentry.Cipslaautogroupopertype
    leafs["cipslaReactConfigIndex"] = cipslareactentry.Cipslareactconfigindex
    leafs["cipslaAutoGroupOperTemplateName"] = cipslareactentry.Cipslaautogroupopertemplatename
    leafs["cipslaReactVar"] = cipslareactentry.Cipslareactvar
    leafs["cipslaReactThresholdType"] = cipslareactentry.Cipslareactthresholdtype
    leafs["cipslaReactActionType"] = cipslareactentry.Cipslareactactiontype
    leafs["cipslaReactThresholdRising"] = cipslareactentry.Cipslareactthresholdrising
    leafs["cipslaReactThresholdFalling"] = cipslareactentry.Cipslareactthresholdfalling
    leafs["cipslaReactThresholdCountX"] = cipslareactentry.Cipslareactthresholdcountx
    leafs["cipslaReactThresholdCountY"] = cipslareactentry.Cipslareactthresholdcounty
    leafs["cipslaReactStorageType"] = cipslareactentry.Cipslareactstoragetype
    leafs["cipslaReactRowStatus"] = cipslareactentry.Cipslareactrowstatus
    return leafs
}

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetYangName() string { return "cipslaReactEntry" }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) SetParent(parent types.Entity) { cipslareactentry.parent = parent }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetParent() types.Entity { return cipslareactentry.parent }

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetParentYangName() string { return "cipslaReactTable" }

// CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactactiontype represents notificationOnly(2)    - a notification is generated
type CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactactiontype string

const (
    CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactactiontype_none CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactactiontype = "none"

    CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactactiontype_notificationOnly CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactactiontype = "notificationOnly"
)

// CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype represents rttMonReactOccurred was true.
type CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype string

const (
    CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype_never CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype = "never"

    CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype_immediate CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype = "immediate"

    CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype_consecutive CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype = "consecutive"

    CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype_xOfy CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype = "xOfy"

    CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype_average CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry_Cipslareactthresholdtype = "average"
)

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable
// A table of group scheduling definitions.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for group scheduling.
    // The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry.
    Cipslaautogroupschedentry []CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry
}

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetFilter() yfilter.YFilter { return cipslaautogroupschedtable.YFilter }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) SetFilter(yf yfilter.YFilter) { cipslaautogroupschedtable.YFilter = yf }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetGoName(yname string) string {
    if yname == "cipslaAutoGroupSchedEntry" { return "Cipslaautogroupschedentry" }
    return ""
}

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetSegmentPath() string {
    return "cipslaAutoGroupSchedTable"
}

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipslaAutoGroupSchedEntry" {
        for _, c := range cipslaautogroupschedtable.Cipslaautogroupschedentry {
            if cipslaautogroupschedtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry{}
        cipslaautogroupschedtable.Cipslaautogroupschedentry = append(cipslaautogroupschedtable.Cipslaautogroupschedentry, child)
        return &cipslaautogroupschedtable.Cipslaautogroupschedentry[len(cipslaautogroupschedtable.Cipslaautogroupschedentry)-1]
    }
    return nil
}

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipslaautogroupschedtable.Cipslaautogroupschedentry {
        children[cipslaautogroupschedtable.Cipslaautogroupschedentry[i].GetSegmentPath()] = &cipslaautogroupschedtable.Cipslaautogroupschedentry[i]
    }
    return children
}

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetYangName() string { return "cipslaAutoGroupSchedTable" }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) SetParent(parent types.Entity) { cipslaautogroupschedtable.parent = parent }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetParent() types.Entity { return cipslaautogroupschedtable.parent }

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetParentYangName() string { return "CISCO-IPSLA-AUTOMEASURE-MIB" }

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry
// A list of objects that define specific configuration for
// group scheduling.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This string uniquely identifies a row in the
    // cipslaAutoGroupSchedTable. The type is string with length: 1..64.
    Cipslaautogroupschedid interface{}

    // Specifies the time duration between initiating two IP SLA operations
    // generated via the auto measure group. The type is interface{} with range:
    // 100..99000. Units are seconds.
    Cipslaautogroupschedperiod interface{}

    // Specifies the duration between initiating each RTT operation for one IP SLA
    // operation generated via the auto  measure group.  The value of this object
    // is only effective when both cipslaAutoGroupSchedMaxInterval and 
    // cipslaAutoGroupSchedMinInterval have zero values. The type is interface{}
    // with range: 1..604800. Units are seconds.
    Cipslaautogroupschedinterval interface{}

    // This object specifies the life of all the operations that are getting group
    // scheduled. This value will be placed into cipslaAutoGroupSchedRttLife
    // object when this conceptual control row becomes 'active'.  The value
    // 2147483647 has a special meaning. When this object is set to 2147483647,
    // the rttMonCtrlOperRttLife object for all the operations will not decrement,
    // and thus the life time of the  operation will never end. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    Cipslaautogroupschedlife interface{}

    // This object specifies the ageout value of the operations that are getting
    // group scheduled. This value will be placed into
    // rttMonSchedAdminConceptRowAgeout object for each of the operations in the
    // group when this conceptual control row becomes  'active'.  When this value
    // is set to zero, the operations will never ageout. The type is interface{}
    // with range: 0..2073600. Units are seconds.
    Cipslaautogroupschedageout interface{}

    // Specifies the max duration between initiating each RTT operation for one IP
    // SLA operation in the group. The type is interface{} with range: 0..604800.
    // Units are seconds.
    Cipslaautogroupschedmaxinterval interface{}

    // Specifies the min duration between initiating each RTT operation for one IP
    // SLA operation in the group.  The value of this object should be lower than
    // the value of cipslaAutoGroupSchedMaxInterval. The type is interface{} with
    // range: 0..604800. Units are seconds.
    Cipslaautogroupschedmininterval interface{}

    // This is the time in seconds after which the operations of the associated
    // groups  will take transition to active state. When set to the value other
    // than '1' (pending), then all  objects in this row cannot be modified. The
    // type is interface{} with range: 0..604800. Units are seconds.
    Cipslaautogroupschedstarttime interface{}

    // The storage type of this conceptual row.  By default the entry will be
    // saved into non-volatile memory. The type is StorageType.
    Cipslaautogroupschedstoragetype interface{}

    // The status of the conceptual group schedule control row.  When the status
    // is active and the value of  cipslaAutoGroupSchedStartTime is '1', the other
    // writable  objects may be modified.  This object can be set to 'destroy'
    // from any value at any time. When this object is set to 'destroy' it will
    // stop all the  operations which had been group scheduled by it earlier, 
    // before destroying the group schedule control row. The type is RowStatus.
    Cipslaautogroupschedrowstatus interface{}
}

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetFilter() yfilter.YFilter { return cipslaautogroupschedentry.YFilter }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) SetFilter(yf yfilter.YFilter) { cipslaautogroupschedentry.YFilter = yf }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetGoName(yname string) string {
    if yname == "cipslaAutoGroupSchedId" { return "Cipslaautogroupschedid" }
    if yname == "cipslaAutoGroupSchedPeriod" { return "Cipslaautogroupschedperiod" }
    if yname == "cipslaAutoGroupSchedInterval" { return "Cipslaautogroupschedinterval" }
    if yname == "cipslaAutoGroupSchedLife" { return "Cipslaautogroupschedlife" }
    if yname == "cipslaAutoGroupSchedAgeout" { return "Cipslaautogroupschedageout" }
    if yname == "cipslaAutoGroupSchedMaxInterval" { return "Cipslaautogroupschedmaxinterval" }
    if yname == "cipslaAutoGroupSchedMinInterval" { return "Cipslaautogroupschedmininterval" }
    if yname == "cipslaAutoGroupSchedStartTime" { return "Cipslaautogroupschedstarttime" }
    if yname == "cipslaAutoGroupSchedStorageType" { return "Cipslaautogroupschedstoragetype" }
    if yname == "cipslaAutoGroupSchedRowStatus" { return "Cipslaautogroupschedrowstatus" }
    return ""
}

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetSegmentPath() string {
    return "cipslaAutoGroupSchedEntry" + "[cipslaAutoGroupSchedId='" + fmt.Sprintf("%v", cipslaautogroupschedentry.Cipslaautogroupschedid) + "']"
}

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipslaAutoGroupSchedId"] = cipslaautogroupschedentry.Cipslaautogroupschedid
    leafs["cipslaAutoGroupSchedPeriod"] = cipslaautogroupschedentry.Cipslaautogroupschedperiod
    leafs["cipslaAutoGroupSchedInterval"] = cipslaautogroupschedentry.Cipslaautogroupschedinterval
    leafs["cipslaAutoGroupSchedLife"] = cipslaautogroupschedentry.Cipslaautogroupschedlife
    leafs["cipslaAutoGroupSchedAgeout"] = cipslaautogroupschedentry.Cipslaautogroupschedageout
    leafs["cipslaAutoGroupSchedMaxInterval"] = cipslaautogroupschedentry.Cipslaautogroupschedmaxinterval
    leafs["cipslaAutoGroupSchedMinInterval"] = cipslaautogroupschedentry.Cipslaautogroupschedmininterval
    leafs["cipslaAutoGroupSchedStartTime"] = cipslaautogroupschedentry.Cipslaautogroupschedstarttime
    leafs["cipslaAutoGroupSchedStorageType"] = cipslaautogroupschedentry.Cipslaautogroupschedstoragetype
    leafs["cipslaAutoGroupSchedRowStatus"] = cipslaautogroupschedentry.Cipslaautogroupschedrowstatus
    return leafs
}

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetYangName() string { return "cipslaAutoGroupSchedEntry" }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) SetParent(parent types.Entity) { cipslaautogroupschedentry.parent = parent }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetParent() types.Entity { return cipslaautogroupschedentry.parent }

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetParentYangName() string { return "cipslaAutoGroupSchedTable" }

