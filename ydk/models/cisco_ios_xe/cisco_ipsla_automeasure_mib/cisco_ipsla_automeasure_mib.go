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
    CipslaAutoGroupTable CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable

    // A table contains the list of destination IP addresses and ports associated
    // to the auto measure group destination name.
    CipslaAutoGroupDestTable CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable

    // A table that contains reaction configurations for templates. Each
    // conceptual row in cipslaReactTable corresponds  to a reaction configured
    // for one template.  Each template can have multiple reactions and hence
    // there can be  multiple rows for a particular template. Different template
    // types can have different reactions. The reaction type is  specified as
    // cipslaReactVar based upon template type as some  reaction types are
    // applicable just for specific template types.
    CipslaReactTable CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable

    // A table of group scheduling definitions.
    CipslaAutoGroupSchedTable CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable
}

func (cISCOIPSLAAUTOMEASUREMIB *CISCOIPSLAAUTOMEASUREMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSLAAUTOMEASUREMIB.EntityData.YFilter = cISCOIPSLAAUTOMEASUREMIB.YFilter
    cISCOIPSLAAUTOMEASUREMIB.EntityData.YangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cISCOIPSLAAUTOMEASUREMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSLAAUTOMEASUREMIB.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cISCOIPSLAAUTOMEASUREMIB.EntityData.SegmentPath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB"
    cISCOIPSLAAUTOMEASUREMIB.EntityData.AbsolutePath = cISCOIPSLAAUTOMEASUREMIB.EntityData.SegmentPath
    cISCOIPSLAAUTOMEASUREMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSLAAUTOMEASUREMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSLAAUTOMEASUREMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children.Append("cipslaAutoGroupTable", types.YChild{"CipslaAutoGroupTable", &cISCOIPSLAAUTOMEASUREMIB.CipslaAutoGroupTable})
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children.Append("cipslaAutoGroupDestTable", types.YChild{"CipslaAutoGroupDestTable", &cISCOIPSLAAUTOMEASUREMIB.CipslaAutoGroupDestTable})
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children.Append("cipslaReactTable", types.YChild{"CipslaReactTable", &cISCOIPSLAAUTOMEASUREMIB.CipslaReactTable})
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Children.Append("cipslaAutoGroupSchedTable", types.YChild{"CipslaAutoGroupSchedTable", &cISCOIPSLAAUTOMEASUREMIB.CipslaAutoGroupSchedTable})
    cISCOIPSLAAUTOMEASUREMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPSLAAUTOMEASUREMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPSLAAUTOMEASUREMIB.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable
// A table that contains IP SLA auto measure group definitions.
type CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the configurations for a particular auto measure group.
    // The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable_CipslaAutoGroupEntry.
    CipslaAutoGroupEntry []*CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable_CipslaAutoGroupEntry
}

func (cipslaAutoGroupTable *CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable) GetEntityData() *types.CommonEntityData {
    cipslaAutoGroupTable.EntityData.YFilter = cipslaAutoGroupTable.YFilter
    cipslaAutoGroupTable.EntityData.YangName = "cipslaAutoGroupTable"
    cipslaAutoGroupTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaAutoGroupTable.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cipslaAutoGroupTable.EntityData.SegmentPath = "cipslaAutoGroupTable"
    cipslaAutoGroupTable.EntityData.AbsolutePath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB/" + cipslaAutoGroupTable.EntityData.SegmentPath
    cipslaAutoGroupTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaAutoGroupTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaAutoGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaAutoGroupTable.EntityData.Children = types.NewOrderedMap()
    cipslaAutoGroupTable.EntityData.Children.Append("cipslaAutoGroupEntry", types.YChild{"CipslaAutoGroupEntry", nil})
    for i := range cipslaAutoGroupTable.CipslaAutoGroupEntry {
        cipslaAutoGroupTable.EntityData.Children.Append(types.GetSegmentPath(cipslaAutoGroupTable.CipslaAutoGroupEntry[i]), types.YChild{"CipslaAutoGroupEntry", cipslaAutoGroupTable.CipslaAutoGroupEntry[i]})
    }
    cipslaAutoGroupTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaAutoGroupTable.EntityData.YListKeys = []string {}

    return &(cipslaAutoGroupTable.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable_CipslaAutoGroupEntry
// An entry containing the configurations for a particular
// auto measure group.
type CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable_CipslaAutoGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A group name which is used by a management
    // application to identify the group. The type is string with length: 1..64.
    CipslaAutoGroupName interface{}

    // This field is used to provide description for the group. The type is string
    // with length: 0..128.
    CipslaAutoGroupDescription interface{}

    // This object refers to the cipslaAutoGroupDestName in
    // cipslaAutoGroupDestTable. If the name entered  is not present in
    // cipslaAutoGroupDestTable, then when  group is scheduled, no ip sla
    // operations will be created. The type is string with length: 0..64.
    CipslaAutoGroupDestinationName interface{}

    // This object represents the destination port number for auto discovery use.
    // The type is interface{} with range: 0..65535.
    CipslaAutoGroupADDestPort interface{}

    // A string which is used by a management application to identify the template
    // which is associated with the group. Depends on cipslaAutoGroupOperType,
    // this object refers to cipslaIcmpEchoTmplName in cipslaIcmpEchoTmplTable, or
    // cipslaUdpEchoTmplName in cipslaUdpEchoTmplTable, or cipslaTcpConnTmplName
    // in cipslaTcpConnTmplTable, or cipslaIcmpJitterTmplName in
    // cipslaIcmpJitterTmplTable, or ciscoIpSlaUdpJitterTmplName in
    // ciscoIpSlaUdpJitterTmplTable. The type is string with length: 0..64.
    CipslaAutoGroupOperTemplateName interface{}

    // This object refers to the cipslaAutoGroupSchedId in
    // cipslaAutoGroupSchedTable, and is used to schedule  this group. The type is
    // string with length: 0..64.
    CipslaAutoGroupSchedulerId interface{}

    // When this object is set to true, QoS is enabled for this group and this
    // group is linked to policy map. The  restriction is that after QoS is
    // enabled, it can not be  disabled for this group. The type is bool.
    CipslaAutoGroupQoSEnable interface{}

    // This object specifies the type of IP SLA operation. When operation type is
    // not ICMP jitter, then  cipslaAutoGroupOperTemplateName must be specified.
    // The type is IpSlaOperType.
    CipslaAutoGroupOperType interface{}

    // When this object is set to true, destination IP address is populated
    // through auto-discovery. The type is bool.
    CipslaAutoGroupDestIPADEnable interface{}

    // This object specifies number of measurement retries to be attempted for the
    // discovered end point after the  connection to the end point is broken. If
    // there is no  re-registration message received, the end point will be  in
    // inactive state.  When the value of cipslaAutoGroupDestIPADEnable is 
    // 'false', the value of this object has no effect. The type is interface{}
    // with range: 1..65536.
    CipslaAutoGroupADMeasureRetry interface{}

    // This object represents the ageout time for the discovered end point.  If
    // the end point becomes inactive for the period  of ageout time, the end
    // point will be removed from the  discovered end point list.  When the value
    // of cipslaAutoGroupDestIPADEnable is  'false', the value of this object has
    // no effect. The type is interface{} with range: 0..65536. Units are seconds.
    CipslaAutoGroupADDestIPAgeout interface{}

    // The storage type of this conceptual row. The type is StorageType.
    CipslaAutoGroupStorageType interface{}

    // The status of the conceptual group control row.  When the status is active,
    // the other writable objects may be modified unless the scheduler with name 
    // specified by cipslaAutoGroupSchedulerId is scheduled. The type is
    // RowStatus.
    CipslaAutoGroupRowStatus interface{}
}

func (cipslaAutoGroupEntry *CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable_CipslaAutoGroupEntry) GetEntityData() *types.CommonEntityData {
    cipslaAutoGroupEntry.EntityData.YFilter = cipslaAutoGroupEntry.YFilter
    cipslaAutoGroupEntry.EntityData.YangName = "cipslaAutoGroupEntry"
    cipslaAutoGroupEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaAutoGroupEntry.EntityData.ParentYangName = "cipslaAutoGroupTable"
    cipslaAutoGroupEntry.EntityData.SegmentPath = "cipslaAutoGroupEntry" + types.AddKeyToken(cipslaAutoGroupEntry.CipslaAutoGroupName, "cipslaAutoGroupName")
    cipslaAutoGroupEntry.EntityData.AbsolutePath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB/cipslaAutoGroupTable/" + cipslaAutoGroupEntry.EntityData.SegmentPath
    cipslaAutoGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaAutoGroupEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaAutoGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaAutoGroupEntry.EntityData.Children = types.NewOrderedMap()
    cipslaAutoGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupName", types.YLeaf{"CipslaAutoGroupName", cipslaAutoGroupEntry.CipslaAutoGroupName})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupDescription", types.YLeaf{"CipslaAutoGroupDescription", cipslaAutoGroupEntry.CipslaAutoGroupDescription})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupDestinationName", types.YLeaf{"CipslaAutoGroupDestinationName", cipslaAutoGroupEntry.CipslaAutoGroupDestinationName})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupADDestPort", types.YLeaf{"CipslaAutoGroupADDestPort", cipslaAutoGroupEntry.CipslaAutoGroupADDestPort})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupOperTemplateName", types.YLeaf{"CipslaAutoGroupOperTemplateName", cipslaAutoGroupEntry.CipslaAutoGroupOperTemplateName})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedulerId", types.YLeaf{"CipslaAutoGroupSchedulerId", cipslaAutoGroupEntry.CipslaAutoGroupSchedulerId})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupQoSEnable", types.YLeaf{"CipslaAutoGroupQoSEnable", cipslaAutoGroupEntry.CipslaAutoGroupQoSEnable})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupOperType", types.YLeaf{"CipslaAutoGroupOperType", cipslaAutoGroupEntry.CipslaAutoGroupOperType})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupDestIPADEnable", types.YLeaf{"CipslaAutoGroupDestIPADEnable", cipslaAutoGroupEntry.CipslaAutoGroupDestIPADEnable})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupADMeasureRetry", types.YLeaf{"CipslaAutoGroupADMeasureRetry", cipslaAutoGroupEntry.CipslaAutoGroupADMeasureRetry})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupADDestIPAgeout", types.YLeaf{"CipslaAutoGroupADDestIPAgeout", cipslaAutoGroupEntry.CipslaAutoGroupADDestIPAgeout})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupStorageType", types.YLeaf{"CipslaAutoGroupStorageType", cipslaAutoGroupEntry.CipslaAutoGroupStorageType})
    cipslaAutoGroupEntry.EntityData.Leafs.Append("cipslaAutoGroupRowStatus", types.YLeaf{"CipslaAutoGroupRowStatus", cipslaAutoGroupEntry.CipslaAutoGroupRowStatus})

    cipslaAutoGroupEntry.EntityData.YListKeys = []string {"CipslaAutoGroupName"}

    return &(cipslaAutoGroupEntry.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable
// A table contains the list of destination IP
// addresses and ports associated to the auto measure
// group destination name.
type CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry containing the destination IP addresses and port configurations
    // associated to auto measure group destination name. The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable_CipslaAutoGroupDestEntry.
    CipslaAutoGroupDestEntry []*CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable_CipslaAutoGroupDestEntry
}

func (cipslaAutoGroupDestTable *CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable) GetEntityData() *types.CommonEntityData {
    cipslaAutoGroupDestTable.EntityData.YFilter = cipslaAutoGroupDestTable.YFilter
    cipslaAutoGroupDestTable.EntityData.YangName = "cipslaAutoGroupDestTable"
    cipslaAutoGroupDestTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaAutoGroupDestTable.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cipslaAutoGroupDestTable.EntityData.SegmentPath = "cipslaAutoGroupDestTable"
    cipslaAutoGroupDestTable.EntityData.AbsolutePath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB/" + cipslaAutoGroupDestTable.EntityData.SegmentPath
    cipslaAutoGroupDestTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaAutoGroupDestTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaAutoGroupDestTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaAutoGroupDestTable.EntityData.Children = types.NewOrderedMap()
    cipslaAutoGroupDestTable.EntityData.Children.Append("cipslaAutoGroupDestEntry", types.YChild{"CipslaAutoGroupDestEntry", nil})
    for i := range cipslaAutoGroupDestTable.CipslaAutoGroupDestEntry {
        cipslaAutoGroupDestTable.EntityData.Children.Append(types.GetSegmentPath(cipslaAutoGroupDestTable.CipslaAutoGroupDestEntry[i]), types.YChild{"CipslaAutoGroupDestEntry", cipslaAutoGroupDestTable.CipslaAutoGroupDestEntry[i]})
    }
    cipslaAutoGroupDestTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaAutoGroupDestTable.EntityData.YListKeys = []string {}

    return &(cipslaAutoGroupDestTable.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable_CipslaAutoGroupDestEntry
// An entry containing the destination IP addresses
// and port configurations associated to auto measure
// group destination name.
type CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable_CipslaAutoGroupDestEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is the name for an auto measure group
    // destination. The type is string with length: 1..64.
    CipslaAutoGroupDestName interface{}

    // This attribute is a key. The type of the internet address of a destination
    // for an auto measure group. The type is InetAddressType.
    CipslaAutoGroupDestIpAddrType interface{}

    // This attribute is a key. The internet address of a destination for an auto
    // measure group. The type of this address is determined by the value of
    // cipslaAutoGroupDestIpAddrType. The type is string with length: 0..255.
    CipslaAutoGroupDestIpAddr interface{}

    // This attribute is a key. This object represents the destination port
    // number. For ICMP echo and ICMP jitter, the suggested value is  '0'. The
    // type is interface{} with range: 0..65535.
    CipslaAutoGroupDestPort interface{}

    // The storage type of this conceptual row.  By default the entry will be
    // saved into non-volatile memory. The type is StorageType.
    CipslaAutoGroupDestStorageType interface{}

    // The status of the conceptual destination table control row. No other
    // objects in this row need to be set before this object can become active. 
    // During 'destroy', when cipslaAutoGroupDestIpAddr is specified  as '0.0.0.0'
    // and cipslaAutoGroupDestPort is specified as '0',  then all the rows with
    // same cipslaAutoGroupDestName will be  deleted. The type is RowStatus.
    CipslaAutoGroupDestRowStatus interface{}
}

func (cipslaAutoGroupDestEntry *CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupDestTable_CipslaAutoGroupDestEntry) GetEntityData() *types.CommonEntityData {
    cipslaAutoGroupDestEntry.EntityData.YFilter = cipslaAutoGroupDestEntry.YFilter
    cipslaAutoGroupDestEntry.EntityData.YangName = "cipslaAutoGroupDestEntry"
    cipslaAutoGroupDestEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaAutoGroupDestEntry.EntityData.ParentYangName = "cipslaAutoGroupDestTable"
    cipslaAutoGroupDestEntry.EntityData.SegmentPath = "cipslaAutoGroupDestEntry" + types.AddKeyToken(cipslaAutoGroupDestEntry.CipslaAutoGroupDestName, "cipslaAutoGroupDestName") + types.AddKeyToken(cipslaAutoGroupDestEntry.CipslaAutoGroupDestIpAddrType, "cipslaAutoGroupDestIpAddrType") + types.AddKeyToken(cipslaAutoGroupDestEntry.CipslaAutoGroupDestIpAddr, "cipslaAutoGroupDestIpAddr") + types.AddKeyToken(cipslaAutoGroupDestEntry.CipslaAutoGroupDestPort, "cipslaAutoGroupDestPort")
    cipslaAutoGroupDestEntry.EntityData.AbsolutePath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB/cipslaAutoGroupDestTable/" + cipslaAutoGroupDestEntry.EntityData.SegmentPath
    cipslaAutoGroupDestEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaAutoGroupDestEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaAutoGroupDestEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaAutoGroupDestEntry.EntityData.Children = types.NewOrderedMap()
    cipslaAutoGroupDestEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaAutoGroupDestEntry.EntityData.Leafs.Append("cipslaAutoGroupDestName", types.YLeaf{"CipslaAutoGroupDestName", cipslaAutoGroupDestEntry.CipslaAutoGroupDestName})
    cipslaAutoGroupDestEntry.EntityData.Leafs.Append("cipslaAutoGroupDestIpAddrType", types.YLeaf{"CipslaAutoGroupDestIpAddrType", cipslaAutoGroupDestEntry.CipslaAutoGroupDestIpAddrType})
    cipslaAutoGroupDestEntry.EntityData.Leafs.Append("cipslaAutoGroupDestIpAddr", types.YLeaf{"CipslaAutoGroupDestIpAddr", cipslaAutoGroupDestEntry.CipslaAutoGroupDestIpAddr})
    cipslaAutoGroupDestEntry.EntityData.Leafs.Append("cipslaAutoGroupDestPort", types.YLeaf{"CipslaAutoGroupDestPort", cipslaAutoGroupDestEntry.CipslaAutoGroupDestPort})
    cipslaAutoGroupDestEntry.EntityData.Leafs.Append("cipslaAutoGroupDestStorageType", types.YLeaf{"CipslaAutoGroupDestStorageType", cipslaAutoGroupDestEntry.CipslaAutoGroupDestStorageType})
    cipslaAutoGroupDestEntry.EntityData.Leafs.Append("cipslaAutoGroupDestRowStatus", types.YLeaf{"CipslaAutoGroupDestRowStatus", cipslaAutoGroupDestEntry.CipslaAutoGroupDestRowStatus})

    cipslaAutoGroupDestEntry.EntityData.YListKeys = []string {"CipslaAutoGroupDestName", "CipslaAutoGroupDestIpAddrType", "CipslaAutoGroupDestIpAddr", "CipslaAutoGroupDestPort"}

    return &(cipslaAutoGroupDestEntry.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable
// A table that contains reaction configurations for templates.
// Each conceptual row in cipslaReactTable corresponds 
// to a reaction configured for one template.
// 
// Each template can have multiple reactions and hence there can be 
// multiple rows for a particular template. Different template
// types can have different reactions. The reaction type is 
// specified as cipslaReactVar based upon template type as some 
// reaction types are applicable just for specific template types.
type CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A base list of objects that define a conceptual reaction configuration
    // control row. The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry.
    CipslaReactEntry []*CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry
}

func (cipslaReactTable *CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable) GetEntityData() *types.CommonEntityData {
    cipslaReactTable.EntityData.YFilter = cipslaReactTable.YFilter
    cipslaReactTable.EntityData.YangName = "cipslaReactTable"
    cipslaReactTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaReactTable.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cipslaReactTable.EntityData.SegmentPath = "cipslaReactTable"
    cipslaReactTable.EntityData.AbsolutePath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB/" + cipslaReactTable.EntityData.SegmentPath
    cipslaReactTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaReactTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaReactTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaReactTable.EntityData.Children = types.NewOrderedMap()
    cipslaReactTable.EntityData.Children.Append("cipslaReactEntry", types.YChild{"CipslaReactEntry", nil})
    for i := range cipslaReactTable.CipslaReactEntry {
        cipslaReactTable.EntityData.Children.Append(types.GetSegmentPath(cipslaReactTable.CipslaReactEntry[i]), types.YChild{"CipslaReactEntry", cipslaReactTable.CipslaReactEntry[i]})
    }
    cipslaReactTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaReactTable.EntityData.YListKeys = []string {}

    return &(cipslaReactTable.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry
// A base list of objects that define a conceptual reaction
// configuration control row.
type CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is IpSlaOperType. Refers to
    // cisco_ipsla_automeasure_mib.CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable_CipslaAutoGroupEntry_CipslaAutoGroupOperType
    CipslaAutoGroupOperType interface{}

    // This attribute is a key. This object along with cipslaAutoGroupOperType and
    // cipslaAutoGroupOperTemplateName identifies a particular
    // reaction-configuration for one IP SLA  template.  This number is persistent
    // across reboots. The type is interface{} with range: 1..2147483647.
    CipslaReactConfigIndex interface{}

    // This attribute is a key. The type is string with length: 0..64. Refers to
    // cisco_ipsla_automeasure_mib.CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupTable_CipslaAutoGroupEntry_CipslaAutoGroupOperTemplateName
    CipslaAutoGroupOperTemplateName interface{}

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
    CipslaReactVar interface{}

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
    // rttMonReactOccurred was true. The type is CipslaReactThresholdType.
    CipslaReactThresholdType interface{}

    // Specifies what type, if any, of reaction to generate if one of the watched
    // (reaction-configuration ) conditions is satisfied:  none(1)               
    // - no reaction is generated notificationOnly(2)    - a notification is
    // generated. The type is CipslaReactActionType.
    CipslaReactActionType interface{}

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
    CipslaReactThresholdRising interface{}

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
    CipslaReactThresholdFalling interface{}

    // If cipslaReactThresholdType value is 'xOfy', this object defines the 'x'
    // value.  If cipslaReactThresholdType value is 'consecutive' this object
    // defines the number of consecutive occurrences that needs threshold
    // violation before setting  cipslaReactOccurred as true.  If
    // cipslaReactThresholdType value is 'average' this object defines the number
    // of samples that needs be considered for calculating average.  This object
    // has no meaning if cipslaReactThresholdType has value of 'never' and
    // 'immediate'. The type is interface{} with range: 1..16.
    CipslaReactThresholdCountX interface{}

    // This object defines the 'y' value of the xOfy condition if
    // cipslaReactThresholdType is 'xOfy'. The default for the  'y' value is 5. 
    // For other values of cipslaReactThresholdType, this object is not
    // applicable. The type is interface{} with range: 1..16.
    CipslaReactThresholdCountY interface{}

    // The storage type of this conceptual row.  By default the entry will be
    // saved into non-volatile memory. The type is StorageType.
    CipslaReactStorageType interface{}

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
    CipslaReactRowStatus interface{}
}

func (cipslaReactEntry *CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry) GetEntityData() *types.CommonEntityData {
    cipslaReactEntry.EntityData.YFilter = cipslaReactEntry.YFilter
    cipslaReactEntry.EntityData.YangName = "cipslaReactEntry"
    cipslaReactEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaReactEntry.EntityData.ParentYangName = "cipslaReactTable"
    cipslaReactEntry.EntityData.SegmentPath = "cipslaReactEntry" + types.AddKeyToken(cipslaReactEntry.CipslaAutoGroupOperType, "cipslaAutoGroupOperType") + types.AddKeyToken(cipslaReactEntry.CipslaReactConfigIndex, "cipslaReactConfigIndex") + types.AddKeyToken(cipslaReactEntry.CipslaAutoGroupOperTemplateName, "cipslaAutoGroupOperTemplateName")
    cipslaReactEntry.EntityData.AbsolutePath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB/cipslaReactTable/" + cipslaReactEntry.EntityData.SegmentPath
    cipslaReactEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaReactEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaReactEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaReactEntry.EntityData.Children = types.NewOrderedMap()
    cipslaReactEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaReactEntry.EntityData.Leafs.Append("cipslaAutoGroupOperType", types.YLeaf{"CipslaAutoGroupOperType", cipslaReactEntry.CipslaAutoGroupOperType})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactConfigIndex", types.YLeaf{"CipslaReactConfigIndex", cipslaReactEntry.CipslaReactConfigIndex})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaAutoGroupOperTemplateName", types.YLeaf{"CipslaAutoGroupOperTemplateName", cipslaReactEntry.CipslaAutoGroupOperTemplateName})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactVar", types.YLeaf{"CipslaReactVar", cipslaReactEntry.CipslaReactVar})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactThresholdType", types.YLeaf{"CipslaReactThresholdType", cipslaReactEntry.CipslaReactThresholdType})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactActionType", types.YLeaf{"CipslaReactActionType", cipslaReactEntry.CipslaReactActionType})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactThresholdRising", types.YLeaf{"CipslaReactThresholdRising", cipslaReactEntry.CipslaReactThresholdRising})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactThresholdFalling", types.YLeaf{"CipslaReactThresholdFalling", cipslaReactEntry.CipslaReactThresholdFalling})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactThresholdCountX", types.YLeaf{"CipslaReactThresholdCountX", cipslaReactEntry.CipslaReactThresholdCountX})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactThresholdCountY", types.YLeaf{"CipslaReactThresholdCountY", cipslaReactEntry.CipslaReactThresholdCountY})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactStorageType", types.YLeaf{"CipslaReactStorageType", cipslaReactEntry.CipslaReactStorageType})
    cipslaReactEntry.EntityData.Leafs.Append("cipslaReactRowStatus", types.YLeaf{"CipslaReactRowStatus", cipslaReactEntry.CipslaReactRowStatus})

    cipslaReactEntry.EntityData.YListKeys = []string {"CipslaAutoGroupOperType", "CipslaReactConfigIndex", "CipslaAutoGroupOperTemplateName"}

    return &(cipslaReactEntry.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactActionType represents notificationOnly(2)    - a notification is generated
type CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactActionType string

const (
    CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactActionType_none CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactActionType = "none"

    CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactActionType_notificationOnly CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactActionType = "notificationOnly"
)

// CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType represents rttMonReactOccurred was true.
type CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType string

const (
    CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType_never CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType = "never"

    CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType_immediate CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType = "immediate"

    CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType_consecutive CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType = "consecutive"

    CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType_xOfy CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType = "xOfy"

    CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType_average CISCOIPSLAAUTOMEASUREMIB_CipslaReactTable_CipslaReactEntry_CipslaReactThresholdType = "average"
)

// CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable
// A table of group scheduling definitions.
type CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of objects that define specific configuration for group scheduling.
    // The type is slice of
    // CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable_CipslaAutoGroupSchedEntry.
    CipslaAutoGroupSchedEntry []*CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable_CipslaAutoGroupSchedEntry
}

func (cipslaAutoGroupSchedTable *CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable) GetEntityData() *types.CommonEntityData {
    cipslaAutoGroupSchedTable.EntityData.YFilter = cipslaAutoGroupSchedTable.YFilter
    cipslaAutoGroupSchedTable.EntityData.YangName = "cipslaAutoGroupSchedTable"
    cipslaAutoGroupSchedTable.EntityData.BundleName = "cisco_ios_xe"
    cipslaAutoGroupSchedTable.EntityData.ParentYangName = "CISCO-IPSLA-AUTOMEASURE-MIB"
    cipslaAutoGroupSchedTable.EntityData.SegmentPath = "cipslaAutoGroupSchedTable"
    cipslaAutoGroupSchedTable.EntityData.AbsolutePath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB/" + cipslaAutoGroupSchedTable.EntityData.SegmentPath
    cipslaAutoGroupSchedTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaAutoGroupSchedTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaAutoGroupSchedTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaAutoGroupSchedTable.EntityData.Children = types.NewOrderedMap()
    cipslaAutoGroupSchedTable.EntityData.Children.Append("cipslaAutoGroupSchedEntry", types.YChild{"CipslaAutoGroupSchedEntry", nil})
    for i := range cipslaAutoGroupSchedTable.CipslaAutoGroupSchedEntry {
        cipslaAutoGroupSchedTable.EntityData.Children.Append(types.GetSegmentPath(cipslaAutoGroupSchedTable.CipslaAutoGroupSchedEntry[i]), types.YChild{"CipslaAutoGroupSchedEntry", cipslaAutoGroupSchedTable.CipslaAutoGroupSchedEntry[i]})
    }
    cipslaAutoGroupSchedTable.EntityData.Leafs = types.NewOrderedMap()

    cipslaAutoGroupSchedTable.EntityData.YListKeys = []string {}

    return &(cipslaAutoGroupSchedTable.EntityData)
}

// CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable_CipslaAutoGroupSchedEntry
// A list of objects that define specific configuration for
// group scheduling.
type CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable_CipslaAutoGroupSchedEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This string uniquely identifies a row in the
    // cipslaAutoGroupSchedTable. The type is string with length: 1..64.
    CipslaAutoGroupSchedId interface{}

    // Specifies the time duration between initiating two IP SLA operations
    // generated via the auto measure group. The type is interface{} with range:
    // 100..99000. Units are seconds.
    CipslaAutoGroupSchedPeriod interface{}

    // Specifies the duration between initiating each RTT operation for one IP SLA
    // operation generated via the auto  measure group.  The value of this object
    // is only effective when both cipslaAutoGroupSchedMaxInterval and 
    // cipslaAutoGroupSchedMinInterval have zero values. The type is interface{}
    // with range: 1..604800. Units are seconds.
    CipslaAutoGroupSchedInterval interface{}

    // This object specifies the life of all the operations that are getting group
    // scheduled. This value will be placed into cipslaAutoGroupSchedRttLife
    // object when this conceptual control row becomes 'active'.  The value
    // 2147483647 has a special meaning. When this object is set to 2147483647,
    // the rttMonCtrlOperRttLife object for all the operations will not decrement,
    // and thus the life time of the  operation will never end. The type is
    // interface{} with range: 0..2147483647. Units are seconds.
    CipslaAutoGroupSchedLife interface{}

    // This object specifies the ageout value of the operations that are getting
    // group scheduled. This value will be placed into
    // rttMonSchedAdminConceptRowAgeout object for each of the operations in the
    // group when this conceptual control row becomes  'active'.  When this value
    // is set to zero, the operations will never ageout. The type is interface{}
    // with range: 0..2073600. Units are seconds.
    CipslaAutoGroupSchedAgeout interface{}

    // Specifies the max duration between initiating each RTT operation for one IP
    // SLA operation in the group. The type is interface{} with range: 0..604800.
    // Units are seconds.
    CipslaAutoGroupSchedMaxInterval interface{}

    // Specifies the min duration between initiating each RTT operation for one IP
    // SLA operation in the group.  The value of this object should be lower than
    // the value of cipslaAutoGroupSchedMaxInterval. The type is interface{} with
    // range: 0..604800. Units are seconds.
    CipslaAutoGroupSchedMinInterval interface{}

    // This is the time in seconds after which the operations of the associated
    // groups  will take transition to active state. When set to the value other
    // than '1' (pending), then all  objects in this row cannot be modified. The
    // type is interface{} with range: 0..604800. Units are seconds.
    CipslaAutoGroupSchedStartTime interface{}

    // The storage type of this conceptual row.  By default the entry will be
    // saved into non-volatile memory. The type is StorageType.
    CipslaAutoGroupSchedStorageType interface{}

    // The status of the conceptual group schedule control row.  When the status
    // is active and the value of  cipslaAutoGroupSchedStartTime is '1', the other
    // writable  objects may be modified.  This object can be set to 'destroy'
    // from any value at any time. When this object is set to 'destroy' it will
    // stop all the  operations which had been group scheduled by it earlier, 
    // before destroying the group schedule control row. The type is RowStatus.
    CipslaAutoGroupSchedRowStatus interface{}
}

func (cipslaAutoGroupSchedEntry *CISCOIPSLAAUTOMEASUREMIB_CipslaAutoGroupSchedTable_CipslaAutoGroupSchedEntry) GetEntityData() *types.CommonEntityData {
    cipslaAutoGroupSchedEntry.EntityData.YFilter = cipslaAutoGroupSchedEntry.YFilter
    cipslaAutoGroupSchedEntry.EntityData.YangName = "cipslaAutoGroupSchedEntry"
    cipslaAutoGroupSchedEntry.EntityData.BundleName = "cisco_ios_xe"
    cipslaAutoGroupSchedEntry.EntityData.ParentYangName = "cipslaAutoGroupSchedTable"
    cipslaAutoGroupSchedEntry.EntityData.SegmentPath = "cipslaAutoGroupSchedEntry" + types.AddKeyToken(cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedId, "cipslaAutoGroupSchedId")
    cipslaAutoGroupSchedEntry.EntityData.AbsolutePath = "CISCO-IPSLA-AUTOMEASURE-MIB:CISCO-IPSLA-AUTOMEASURE-MIB/cipslaAutoGroupSchedTable/" + cipslaAutoGroupSchedEntry.EntityData.SegmentPath
    cipslaAutoGroupSchedEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipslaAutoGroupSchedEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipslaAutoGroupSchedEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipslaAutoGroupSchedEntry.EntityData.Children = types.NewOrderedMap()
    cipslaAutoGroupSchedEntry.EntityData.Leafs = types.NewOrderedMap()
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedId", types.YLeaf{"CipslaAutoGroupSchedId", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedId})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedPeriod", types.YLeaf{"CipslaAutoGroupSchedPeriod", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedPeriod})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedInterval", types.YLeaf{"CipslaAutoGroupSchedInterval", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedInterval})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedLife", types.YLeaf{"CipslaAutoGroupSchedLife", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedLife})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedAgeout", types.YLeaf{"CipslaAutoGroupSchedAgeout", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedAgeout})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedMaxInterval", types.YLeaf{"CipslaAutoGroupSchedMaxInterval", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedMaxInterval})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedMinInterval", types.YLeaf{"CipslaAutoGroupSchedMinInterval", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedMinInterval})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedStartTime", types.YLeaf{"CipslaAutoGroupSchedStartTime", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedStartTime})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedStorageType", types.YLeaf{"CipslaAutoGroupSchedStorageType", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedStorageType})
    cipslaAutoGroupSchedEntry.EntityData.Leafs.Append("cipslaAutoGroupSchedRowStatus", types.YLeaf{"CipslaAutoGroupSchedRowStatus", cipslaAutoGroupSchedEntry.CipslaAutoGroupSchedRowStatus})

    cipslaAutoGroupSchedEntry.EntityData.YListKeys = []string {"CipslaAutoGroupSchedId"}

    return &(cipslaAutoGroupSchedEntry.EntityData)
}

