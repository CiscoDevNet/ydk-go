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
    EntityData types.CommonEntityData
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

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSLAAUTOMEASUREMIB.EntityData.YFilter = cISCOIPSLAAUTOMEASUREMIB.YFilter
    cISCOIPSLAAUTOMEASUREMIB.EntityData.YangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cISCOIPSLAAUTOMEASUREMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSLAAUTOMEASUREMIB.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cISCOIPSLAAUTOMEASUREMIB.EntityData.SegmentPath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB"
    cISCOIPSLAAUTOMEASUREMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSLAAUTOMEASUREMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSLAAUTOMEASUREMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children["cipslaAutoGroupTable"] = types.YChild{"Cipslaautogrouptable", &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogrouptable}
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children["cipslaAutoGroupDestTable"] = types.YChild{"Cipslaautogroupdesttable", &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogroupdesttable}
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children["cipslaReactTable"] = types.YChild{"Cipslareacttable", &cISCOIPSLAAUTOMEASUREMIB.Cipslareacttable}
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children["cipslaAutoGroupSchedTable"] = types.YChild{"Cipslaautogroupschedtable", &cISCOIPSLAAUTOMEASUREMIB.Cipslaautogroupschedtable}
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPSLAAUTOMEASUREMIB.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable
// A table that contains IP SLA auto measure group definitions.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the configurations for a particular auto measure group.
    // The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry.
    Cipslaautogroupentry []CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry
}

func (cipslaautogrouptable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable) GetEntityData() *types.CommonEntityData {
    cipslaautogrouptable.EntityData.YFilter = cipslaautogrouptable.YFilter
    cipslaautogrouptable.EntityData.YangName = "cipslaAutoGroupTable"
    cipslaautogrouptable.EntityData.BundleName = "cisco_ios_xe"
    cipslaautogrouptable.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cipslaautogrouptable.EntityData.SegmentPath = "cipslaAutoGroupTable"
    cipslaautogrouptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaautogrouptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaautogrouptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaautogrouptable.EntityData.Children = make(map[string]types.YChild)
    cipslaautogrouptable.EntityData.Children["cipslaAutoGroupEntry"] = types.YChild{"Cipslaautogroupentry", nil}
    for i := range cipslaautogrouptable.Cipslaautogroupentry {
        cipslaautogrouptable.EntityData.Children[types.GetSegmentPath(&cipslaautogrouptable.Cipslaautogroupentry[i])] = types.YChild{"Cipslaautogroupentry", &cipslaautogrouptable.Cipslaautogroupentry[i]}
    }
    cipslaautogrouptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslaautogrouptable.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry
// An entry containing the configurations for a particular
// auto measure group.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry struct {
    EntityData types.CommonEntityData
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

func (cipslaautogroupentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogrouptable_Cipslaautogroupentry) GetEntityData() *types.CommonEntityData {
    cipslaautogroupentry.EntityData.YFilter = cipslaautogroupentry.YFilter
    cipslaautogroupentry.EntityData.YangName = "cipslaAutoGroupEntry"
    cipslaautogroupentry.EntityData.BundleName = "cisco_ios_xe"
    cipslaautogroupentry.EntityData.ParentYangName = "cipslaAutoGroupTable"
    cipslaautogroupentry.EntityData.SegmentPath = "cipslaAutoGroupEntry" + "[cipslaAutoGroupName='" + fmt.Sprintf("%v", cipslaautogroupentry.Cipslaautogroupname) + "']"
    cipslaautogroupentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaautogroupentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaautogroupentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaautogroupentry.EntityData.Children = make(map[string]types.YChild)
    cipslaautogroupentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupName"] = types.YLeaf{"Cipslaautogroupname", cipslaautogroupentry.Cipslaautogroupname}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupDescription"] = types.YLeaf{"Cipslaautogroupdescription", cipslaautogroupentry.Cipslaautogroupdescription}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupDestinationName"] = types.YLeaf{"Cipslaautogroupdestinationname", cipslaautogroupentry.Cipslaautogroupdestinationname}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupADDestPort"] = types.YLeaf{"Cipslaautogroupaddestport", cipslaautogroupentry.Cipslaautogroupaddestport}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupOperTemplateName"] = types.YLeaf{"Cipslaautogroupopertemplatename", cipslaautogroupentry.Cipslaautogroupopertemplatename}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupSchedulerId"] = types.YLeaf{"Cipslaautogroupschedulerid", cipslaautogroupentry.Cipslaautogroupschedulerid}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupQoSEnable"] = types.YLeaf{"Cipslaautogroupqosenable", cipslaautogroupentry.Cipslaautogroupqosenable}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupOperType"] = types.YLeaf{"Cipslaautogroupopertype", cipslaautogroupentry.Cipslaautogroupopertype}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupDestIPADEnable"] = types.YLeaf{"Cipslaautogroupdestipadenable", cipslaautogroupentry.Cipslaautogroupdestipadenable}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupADMeasureRetry"] = types.YLeaf{"Cipslaautogroupadmeasureretry", cipslaautogroupentry.Cipslaautogroupadmeasureretry}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupADDestIPAgeout"] = types.YLeaf{"Cipslaautogroupaddestipageout", cipslaautogroupentry.Cipslaautogroupaddestipageout}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupStorageType"] = types.YLeaf{"Cipslaautogroupstoragetype", cipslaautogroupentry.Cipslaautogroupstoragetype}
    cipslaautogroupentry.EntityData.Leafs["cipslaAutoGroupRowStatus"] = types.YLeaf{"Cipslaautogrouprowstatus", cipslaautogroupentry.Cipslaautogrouprowstatus}
    return &(cipslaautogroupentry.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable
// A table contains the list of destination IP
// addresses and ports associated to the auto measure
// group destination name.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the destination IP addresses and port configurations
    // associated to auto measure group destination name. The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry.
    Cipslaautogroupdestentry []CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry
}

func (cipslaautogroupdesttable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable) GetEntityData() *types.CommonEntityData {
    cipslaautogroupdesttable.EntityData.YFilter = cipslaautogroupdesttable.YFilter
    cipslaautogroupdesttable.EntityData.YangName = "cipslaAutoGroupDestTable"
    cipslaautogroupdesttable.EntityData.BundleName = "cisco_ios_xe"
    cipslaautogroupdesttable.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cipslaautogroupdesttable.EntityData.SegmentPath = "cipslaAutoGroupDestTable"
    cipslaautogroupdesttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaautogroupdesttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaautogroupdesttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaautogroupdesttable.EntityData.Children = make(map[string]types.YChild)
    cipslaautogroupdesttable.EntityData.Children["cipslaAutoGroupDestEntry"] = types.YChild{"Cipslaautogroupdestentry", nil}
    for i := range cipslaautogroupdesttable.Cipslaautogroupdestentry {
        cipslaautogroupdesttable.EntityData.Children[types.GetSegmentPath(&cipslaautogroupdesttable.Cipslaautogroupdestentry[i])] = types.YChild{"Cipslaautogroupdestentry", &cipslaautogroupdesttable.Cipslaautogroupdestentry[i]}
    }
    cipslaautogroupdesttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslaautogroupdesttable.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry
// An entry containing the destination IP addresses
// and port configurations associated to auto measure
// group destination name.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry struct {
    EntityData types.CommonEntityData
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

func (cipslaautogroupdestentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupdesttable_Cipslaautogroupdestentry) GetEntityData() *types.CommonEntityData {
    cipslaautogroupdestentry.EntityData.YFilter = cipslaautogroupdestentry.YFilter
    cipslaautogroupdestentry.EntityData.YangName = "cipslaAutoGroupDestEntry"
    cipslaautogroupdestentry.EntityData.BundleName = "cisco_ios_xe"
    cipslaautogroupdestentry.EntityData.ParentYangName = "cipslaAutoGroupDestTable"
    cipslaautogroupdestentry.EntityData.SegmentPath = "cipslaAutoGroupDestEntry" + "[cipslaAutoGroupDestName='" + fmt.Sprintf("%v", cipslaautogroupdestentry.Cipslaautogroupdestname) + "']" + "[cipslaAutoGroupDestIpAddrType='" + fmt.Sprintf("%v", cipslaautogroupdestentry.Cipslaautogroupdestipaddrtype) + "']" + "[cipslaAutoGroupDestIpAddr='" + fmt.Sprintf("%v", cipslaautogroupdestentry.Cipslaautogroupdestipaddr) + "']" + "[cipslaAutoGroupDestPort='" + fmt.Sprintf("%v", cipslaautogroupdestentry.Cipslaautogroupdestport) + "']"
    cipslaautogroupdestentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaautogroupdestentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaautogroupdestentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaautogroupdestentry.EntityData.Children = make(map[string]types.YChild)
    cipslaautogroupdestentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslaautogroupdestentry.EntityData.Leafs["cipslaAutoGroupDestName"] = types.YLeaf{"Cipslaautogroupdestname", cipslaautogroupdestentry.Cipslaautogroupdestname}
    cipslaautogroupdestentry.EntityData.Leafs["cipslaAutoGroupDestIpAddrType"] = types.YLeaf{"Cipslaautogroupdestipaddrtype", cipslaautogroupdestentry.Cipslaautogroupdestipaddrtype}
    cipslaautogroupdestentry.EntityData.Leafs["cipslaAutoGroupDestIpAddr"] = types.YLeaf{"Cipslaautogroupdestipaddr", cipslaautogroupdestentry.Cipslaautogroupdestipaddr}
    cipslaautogroupdestentry.EntityData.Leafs["cipslaAutoGroupDestPort"] = types.YLeaf{"Cipslaautogroupdestport", cipslaautogroupdestentry.Cipslaautogroupdestport}
    cipslaautogroupdestentry.EntityData.Leafs["cipslaAutoGroupDestStorageType"] = types.YLeaf{"Cipslaautogroupdeststoragetype", cipslaautogroupdestentry.Cipslaautogroupdeststoragetype}
    cipslaautogroupdestentry.EntityData.Leafs["cipslaAutoGroupDestRowStatus"] = types.YLeaf{"Cipslaautogroupdestrowstatus", cipslaautogroupdestentry.Cipslaautogroupdestrowstatus}
    return &(cipslaautogroupdestentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual reaction configuration
    // control row. The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry.
    Cipslareactentry []CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry
}

func (cipslareacttable *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable) GetEntityData() *types.CommonEntityData {
    cipslareacttable.EntityData.YFilter = cipslareacttable.YFilter
    cipslareacttable.EntityData.YangName = "cipslaReactTable"
    cipslareacttable.EntityData.BundleName = "cisco_ios_xe"
    cipslareacttable.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cipslareacttable.EntityData.SegmentPath = "cipslaReactTable"
    cipslareacttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslareacttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslareacttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslareacttable.EntityData.Children = make(map[string]types.YChild)
    cipslareacttable.EntityData.Children["cipslaReactEntry"] = types.YChild{"Cipslareactentry", nil}
    for i := range cipslareacttable.Cipslareactentry {
        cipslareacttable.EntityData.Children[types.GetSegmentPath(&cipslareacttable.Cipslareactentry[i])] = types.YChild{"Cipslareactentry", &cipslareacttable.Cipslareactentry[i]}
    }
    cipslareacttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslareacttable.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry
// A base list of objects that define a conceptual reaction
// configuration control row.
type CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry struct {
    EntityData types.CommonEntityData
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

func (cipslareactentry *CISCOIPSLAAUTOMEASUREMIB_Cipslareacttable_Cipslareactentry) GetEntityData() *types.CommonEntityData {
    cipslareactentry.EntityData.YFilter = cipslareactentry.YFilter
    cipslareactentry.EntityData.YangName = "cipslaReactEntry"
    cipslareactentry.EntityData.BundleName = "cisco_ios_xe"
    cipslareactentry.EntityData.ParentYangName = "cipslaReactTable"
    cipslareactentry.EntityData.SegmentPath = "cipslaReactEntry" + "[cipslaAutoGroupOperType='" + fmt.Sprintf("%v", cipslareactentry.Cipslaautogroupopertype) + "']" + "[cipslaReactConfigIndex='" + fmt.Sprintf("%v", cipslareactentry.Cipslareactconfigindex) + "']" + "[cipslaAutoGroupOperTemplateName='" + fmt.Sprintf("%v", cipslareactentry.Cipslaautogroupopertemplatename) + "']"
    cipslareactentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslareactentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslareactentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslareactentry.EntityData.Children = make(map[string]types.YChild)
    cipslareactentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslareactentry.EntityData.Leafs["cipslaAutoGroupOperType"] = types.YLeaf{"Cipslaautogroupopertype", cipslareactentry.Cipslaautogroupopertype}
    cipslareactentry.EntityData.Leafs["cipslaReactConfigIndex"] = types.YLeaf{"Cipslareactconfigindex", cipslareactentry.Cipslareactconfigindex}
    cipslareactentry.EntityData.Leafs["cipslaAutoGroupOperTemplateName"] = types.YLeaf{"Cipslaautogroupopertemplatename", cipslareactentry.Cipslaautogroupopertemplatename}
    cipslareactentry.EntityData.Leafs["cipslaReactVar"] = types.YLeaf{"Cipslareactvar", cipslareactentry.Cipslareactvar}
    cipslareactentry.EntityData.Leafs["cipslaReactThresholdType"] = types.YLeaf{"Cipslareactthresholdtype", cipslareactentry.Cipslareactthresholdtype}
    cipslareactentry.EntityData.Leafs["cipslaReactActionType"] = types.YLeaf{"Cipslareactactiontype", cipslareactentry.Cipslareactactiontype}
    cipslareactentry.EntityData.Leafs["cipslaReactThresholdRising"] = types.YLeaf{"Cipslareactthresholdrising", cipslareactentry.Cipslareactthresholdrising}
    cipslareactentry.EntityData.Leafs["cipslaReactThresholdFalling"] = types.YLeaf{"Cipslareactthresholdfalling", cipslareactentry.Cipslareactthresholdfalling}
    cipslareactentry.EntityData.Leafs["cipslaReactThresholdCountX"] = types.YLeaf{"Cipslareactthresholdcountx", cipslareactentry.Cipslareactthresholdcountx}
    cipslareactentry.EntityData.Leafs["cipslaReactThresholdCountY"] = types.YLeaf{"Cipslareactthresholdcounty", cipslareactentry.Cipslareactthresholdcounty}
    cipslareactentry.EntityData.Leafs["cipslaReactStorageType"] = types.YLeaf{"Cipslareactstoragetype", cipslareactentry.Cipslareactstoragetype}
    cipslareactentry.EntityData.Leafs["cipslaReactRowStatus"] = types.YLeaf{"Cipslareactrowstatus", cipslareactentry.Cipslareactrowstatus}
    return &(cipslareactentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for group scheduling.
    // The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry.
    Cipslaautogroupschedentry []CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry
}

func (cipslaautogroupschedtable *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable) GetEntityData() *types.CommonEntityData {
    cipslaautogroupschedtable.EntityData.YFilter = cipslaautogroupschedtable.YFilter
    cipslaautogroupschedtable.EntityData.YangName = "cipslaAutoGroupSchedTable"
    cipslaautogroupschedtable.EntityData.BundleName = "cisco_ios_xe"
    cipslaautogroupschedtable.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cipslaautogroupschedtable.EntityData.SegmentPath = "cipslaAutoGroupSchedTable"
    cipslaautogroupschedtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaautogroupschedtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaautogroupschedtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaautogroupschedtable.EntityData.Children = make(map[string]types.YChild)
    cipslaautogroupschedtable.EntityData.Children["cipslaAutoGroupSchedEntry"] = types.YChild{"Cipslaautogroupschedentry", nil}
    for i := range cipslaautogroupschedtable.Cipslaautogroupschedentry {
        cipslaautogroupschedtable.EntityData.Children[types.GetSegmentPath(&cipslaautogroupschedtable.Cipslaautogroupschedentry[i])] = types.YChild{"Cipslaautogroupschedentry", &cipslaautogroupschedtable.Cipslaautogroupschedentry[i]}
    }
    cipslaautogroupschedtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipslaautogroupschedtable.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry
// A list of objects that define specific configuration for
// group scheduling.
type CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry struct {
    EntityData types.CommonEntityData
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

func (cipslaautogroupschedentry *CISCOIPSLAAUTOMEASUREMIB_Cipslaautogroupschedtable_Cipslaautogroupschedentry) GetEntityData() *types.CommonEntityData {
    cipslaautogroupschedentry.EntityData.YFilter = cipslaautogroupschedentry.YFilter
    cipslaautogroupschedentry.EntityData.YangName = "cipslaAutoGroupSchedEntry"
    cipslaautogroupschedentry.EntityData.BundleName = "cisco_ios_xe"
    cipslaautogroupschedentry.EntityData.ParentYangName = "cipslaAutoGroupSchedTable"
    cipslaautogroupschedentry.EntityData.SegmentPath = "cipslaAutoGroupSchedEntry" + "[cipslaAutoGroupSchedId='" + fmt.Sprintf("%v", cipslaautogroupschedentry.Cipslaautogroupschedid) + "']"
    cipslaautogroupschedentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaautogroupschedentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaautogroupschedentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaautogroupschedentry.EntityData.Children = make(map[string]types.YChild)
    cipslaautogroupschedentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedId"] = types.YLeaf{"Cipslaautogroupschedid", cipslaautogroupschedentry.Cipslaautogroupschedid}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedPeriod"] = types.YLeaf{"Cipslaautogroupschedperiod", cipslaautogroupschedentry.Cipslaautogroupschedperiod}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedInterval"] = types.YLeaf{"Cipslaautogroupschedinterval", cipslaautogroupschedentry.Cipslaautogroupschedinterval}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedLife"] = types.YLeaf{"Cipslaautogroupschedlife", cipslaautogroupschedentry.Cipslaautogroupschedlife}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedAgeout"] = types.YLeaf{"Cipslaautogroupschedageout", cipslaautogroupschedentry.Cipslaautogroupschedageout}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedMaxInterval"] = types.YLeaf{"Cipslaautogroupschedmaxinterval", cipslaautogroupschedentry.Cipslaautogroupschedmaxinterval}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedMinInterval"] = types.YLeaf{"Cipslaautogroupschedmininterval", cipslaautogroupschedentry.Cipslaautogroupschedmininterval}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedStartTime"] = types.YLeaf{"Cipslaautogroupschedstarttime", cipslaautogroupschedentry.Cipslaautogroupschedstarttime}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedStorageType"] = types.YLeaf{"Cipslaautogroupschedstoragetype", cipslaautogroupschedentry.Cipslaautogroupschedstoragetype}
    cipslaautogroupschedentry.EntityData.Leafs["cipslaAutoGroupSchedRowStatus"] = types.YLeaf{"Cipslaautogroupschedrowstatus", cipslaautogroupschedentry.Cipslaautogroupschedrowstatus}
    return &(cipslaautogroupschedentry.EntityData)
}

