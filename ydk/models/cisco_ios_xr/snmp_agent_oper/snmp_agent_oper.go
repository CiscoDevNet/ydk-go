// This module contains a collection of YANG definitions
// for Cisco IOS-XR snmp-agent package operational data.
// 
// This module contains definitions
// for the following management objects:
//   snmp: SNMP operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package snmp_agent_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_agent_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-agent-oper snmp}", reflect.TypeOf(Snmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-agent-oper:snmp", reflect.TypeOf(Snmp{}))
}

// SnmpCorrVbindMatch represents Snmp corr vbind match
type SnmpCorrVbindMatch string

const (
    // Match regexp to varbind index
    SnmpCorrVbindMatch_index SnmpCorrVbindMatch = "index"

    // Match regexp to varbind value
    SnmpCorrVbindMatch_value SnmpCorrVbindMatch = "value"
)

// SnmpCorrRuleState represents Snmp corr rule state
type SnmpCorrRuleState string

const (
    // Rule is in Unapplied state
    SnmpCorrRuleState_rule_unapplied SnmpCorrRuleState = "rule-unapplied"

    // Rule is Applied to specified hosts
    SnmpCorrRuleState_rule_applied SnmpCorrRuleState = "rule-applied"

    // Rule is Applied to all of router
    SnmpCorrRuleState_rule_applied_all SnmpCorrRuleState = "rule-applied-all"
)

// DupReqDropStatus represents Dup req drop status
type DupReqDropStatus string

const (
    // Disabled
    DupReqDropStatus_disabled DupReqDropStatus = "disabled"

    // Enabled
    DupReqDropStatus_enabled DupReqDropStatus = "enabled"
)

// Snmp
// SNMP operational data
type Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of trap hosts.
    TrapServers Snmp_TrapServers

    // SNMP operational information.
    Information Snmp_Information

    // List of interfaces.
    Interfaces Snmp_Interfaces

    // Trap Correlator operational data.
    Correlator Snmp_Correlator

    // List of index.
    InterfaceIndexes Snmp_InterfaceIndexes

    // List of ifnames.
    IfIndexes Snmp_IfIndexes

    // SNMP entity mib.
    EntityMib Snmp_EntityMib

    // SNMP IF-MIB information.
    InterfaceMib Snmp_InterfaceMib

    // SNMP sensor MIB information.
    SensorMib Snmp_SensorMib
}

func (snmp *Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xr"
    snmp.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-agent-oper"
    snmp.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-agent-oper:snmp"
    snmp.EntityData.AbsolutePath = snmp.EntityData.SegmentPath
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Children.Append("trap-servers", types.YChild{"TrapServers", &snmp.TrapServers})
    snmp.EntityData.Children.Append("information", types.YChild{"Information", &snmp.Information})
    snmp.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &snmp.Interfaces})
    snmp.EntityData.Children.Append("correlator", types.YChild{"Correlator", &snmp.Correlator})
    snmp.EntityData.Children.Append("interface-indexes", types.YChild{"InterfaceIndexes", &snmp.InterfaceIndexes})
    snmp.EntityData.Children.Append("if-indexes", types.YChild{"IfIndexes", &snmp.IfIndexes})
    snmp.EntityData.Children.Append("Cisco-IOS-XR-snmp-entitymib-oper:entity-mib", types.YChild{"EntityMib", &snmp.EntityMib})
    snmp.EntityData.Children.Append("Cisco-IOS-XR-snmp-ifmib-oper:interface-mib", types.YChild{"InterfaceMib", &snmp.InterfaceMib})
    snmp.EntityData.Children.Append("Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib", types.YChild{"SensorMib", &snmp.SensorMib})
    snmp.EntityData.Leafs = types.NewOrderedMap()

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// Snmp_TrapServers
// List of trap hosts
type Snmp_TrapServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trap server and port to which the trap is to be sent and statistics. The
    // type is slice of Snmp_TrapServers_TrapServer.
    TrapServer []*Snmp_TrapServers_TrapServer
}

func (trapServers *Snmp_TrapServers) GetEntityData() *types.CommonEntityData {
    trapServers.EntityData.YFilter = trapServers.YFilter
    trapServers.EntityData.YangName = "trap-servers"
    trapServers.EntityData.BundleName = "cisco_ios_xr"
    trapServers.EntityData.ParentYangName = "snmp"
    trapServers.EntityData.SegmentPath = "trap-servers"
    trapServers.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + trapServers.EntityData.SegmentPath
    trapServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapServers.EntityData.Children = types.NewOrderedMap()
    trapServers.EntityData.Children.Append("trap-server", types.YChild{"TrapServer", nil})
    for i := range trapServers.TrapServer {
        types.SetYListKey(trapServers.TrapServer[i], i)
        trapServers.EntityData.Children.Append(types.GetSegmentPath(trapServers.TrapServer[i]), types.YChild{"TrapServer", trapServers.TrapServer[i]})
    }
    trapServers.EntityData.Leafs = types.NewOrderedMap()

    trapServers.EntityData.YListKeys = []string {}

    return &(trapServers.EntityData)
}

// Snmp_TrapServers_TrapServer
// Trap server and port to which the trap is to be
// sent and statistics
type Snmp_TrapServers_TrapServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trap Host. The type is string.
    TrapHost interface{}

    // Trap port. The type is interface{} with range: 0..65535.
    Port interface{}

    // No. of trap packets in trapQ. The type is interface{} with range:
    // 0..4294967295.
    NumberOfPktsInTrapQ interface{}

    // Maximum Queue length of trapQ. The type is interface{} with range:
    // 0..4294967295.
    MaxQLengthOfTrapQ interface{}

    // No. of trap packets sent. The type is interface{} with range:
    // 0..4294967295.
    NumberOfPktsSent interface{}

    // No. of trap packets dropped. The type is interface{} with range:
    // 0..4294967295.
    NumberOfPktsDropped interface{}
}

func (trapServer *Snmp_TrapServers_TrapServer) GetEntityData() *types.CommonEntityData {
    trapServer.EntityData.YFilter = trapServer.YFilter
    trapServer.EntityData.YangName = "trap-server"
    trapServer.EntityData.BundleName = "cisco_ios_xr"
    trapServer.EntityData.ParentYangName = "trap-servers"
    trapServer.EntityData.SegmentPath = "trap-server" + types.AddNoKeyToken(trapServer)
    trapServer.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/trap-servers/" + trapServer.EntityData.SegmentPath
    trapServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapServer.EntityData.Children = types.NewOrderedMap()
    trapServer.EntityData.Leafs = types.NewOrderedMap()
    trapServer.EntityData.Leafs.Append("trap-host", types.YLeaf{"TrapHost", trapServer.TrapHost})
    trapServer.EntityData.Leafs.Append("port", types.YLeaf{"Port", trapServer.Port})
    trapServer.EntityData.Leafs.Append("number-of-pkts-in-trap-q", types.YLeaf{"NumberOfPktsInTrapQ", trapServer.NumberOfPktsInTrapQ})
    trapServer.EntityData.Leafs.Append("max-q-length-of-trap-q", types.YLeaf{"MaxQLengthOfTrapQ", trapServer.MaxQLengthOfTrapQ})
    trapServer.EntityData.Leafs.Append("number-of-pkts-sent", types.YLeaf{"NumberOfPktsSent", trapServer.NumberOfPktsSent})
    trapServer.EntityData.Leafs.Append("number-of-pkts-dropped", types.YLeaf{"NumberOfPktsDropped", trapServer.NumberOfPktsDropped})

    trapServer.EntityData.YListKeys = []string {}

    return &(trapServer.EntityData)
}

// Snmp_Information
// SNMP operational information
type Snmp_Information struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP host information.
    Hosts Snmp_Information_Hosts

    // System up time.
    SystemUpTime Snmp_Information_SystemUpTime

    // SNMP request type summary .
    NmsAddresses Snmp_Information_NmsAddresses

    // SNMP engine ID.
    EngineId Snmp_Information_EngineId

    // SNMP rx queue statistics.
    RxQueue Snmp_Information_RxQueue

    // System name.
    SystemName Snmp_Information_SystemName

    // SNMP request type details .
    RequestTypeDetail Snmp_Information_RequestTypeDetail

    // Duplicate request status, count, time .
    DuplicateDrop Snmp_Information_DuplicateDrop

    // List of bulkstats transfer on the system.
    BulkStatsTransfers Snmp_Information_BulkStatsTransfers

    // SNMP trap OID.
    TrapInfos Snmp_Information_TrapInfos

    // OID list for poll PDU.
    PollOids Snmp_Information_PollOids

    // SNMP trap OID.
    InfomDetails Snmp_Information_InfomDetails

    // SNMP statistics.
    Statistics Snmp_Information_Statistics

    // Incoming queue details .
    IncomingQueue Snmp_Information_IncomingQueue

    // Context name, features name, topology name, instance.
    ContextMapping Snmp_Information_ContextMapping

    // SNMP trap OID.
    TrapOids Snmp_Information_TrapOids

    // SNMP overload statistics .
    NmSpackets Snmp_Information_NmSpackets

    // List of MIBS supported on the system.
    Mibs Snmp_Information_Mibs

    // SNMP statistics pdu .
    SerialNumbers Snmp_Information_SerialNumbers

    // NMS list for drop PDU.
    DropNmsAddresses Snmp_Information_DropNmsAddresses

    // SNMP view information.
    Views Snmp_Information_Views

    // System description.
    SystemDescr Snmp_Information_SystemDescr

    // List of table.
    Tables Snmp_Information_Tables

    // System object ID.
    SystemOid Snmp_Information_SystemOid

    // SNMP trap queue statistics.
    TrapQueue Snmp_Information_TrapQueue
}

func (information *Snmp_Information) GetEntityData() *types.CommonEntityData {
    information.EntityData.YFilter = information.YFilter
    information.EntityData.YangName = "information"
    information.EntityData.BundleName = "cisco_ios_xr"
    information.EntityData.ParentYangName = "snmp"
    information.EntityData.SegmentPath = "information"
    information.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + information.EntityData.SegmentPath
    information.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    information.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    information.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    information.EntityData.Children = types.NewOrderedMap()
    information.EntityData.Children.Append("hosts", types.YChild{"Hosts", &information.Hosts})
    information.EntityData.Children.Append("system-up-time", types.YChild{"SystemUpTime", &information.SystemUpTime})
    information.EntityData.Children.Append("nms-addresses", types.YChild{"NmsAddresses", &information.NmsAddresses})
    information.EntityData.Children.Append("engine-id", types.YChild{"EngineId", &information.EngineId})
    information.EntityData.Children.Append("rx-queue", types.YChild{"RxQueue", &information.RxQueue})
    information.EntityData.Children.Append("system-name", types.YChild{"SystemName", &information.SystemName})
    information.EntityData.Children.Append("request-type-detail", types.YChild{"RequestTypeDetail", &information.RequestTypeDetail})
    information.EntityData.Children.Append("duplicate-drop", types.YChild{"DuplicateDrop", &information.DuplicateDrop})
    information.EntityData.Children.Append("bulk-stats-transfers", types.YChild{"BulkStatsTransfers", &information.BulkStatsTransfers})
    information.EntityData.Children.Append("trap-infos", types.YChild{"TrapInfos", &information.TrapInfos})
    information.EntityData.Children.Append("poll-oids", types.YChild{"PollOids", &information.PollOids})
    information.EntityData.Children.Append("infom-details", types.YChild{"InfomDetails", &information.InfomDetails})
    information.EntityData.Children.Append("statistics", types.YChild{"Statistics", &information.Statistics})
    information.EntityData.Children.Append("incoming-queue", types.YChild{"IncomingQueue", &information.IncomingQueue})
    information.EntityData.Children.Append("context-mapping", types.YChild{"ContextMapping", &information.ContextMapping})
    information.EntityData.Children.Append("trap-oids", types.YChild{"TrapOids", &information.TrapOids})
    information.EntityData.Children.Append("nm-spackets", types.YChild{"NmSpackets", &information.NmSpackets})
    information.EntityData.Children.Append("mibs", types.YChild{"Mibs", &information.Mibs})
    information.EntityData.Children.Append("serial-numbers", types.YChild{"SerialNumbers", &information.SerialNumbers})
    information.EntityData.Children.Append("drop-nms-addresses", types.YChild{"DropNmsAddresses", &information.DropNmsAddresses})
    information.EntityData.Children.Append("views", types.YChild{"Views", &information.Views})
    information.EntityData.Children.Append("system-descr", types.YChild{"SystemDescr", &information.SystemDescr})
    information.EntityData.Children.Append("tables", types.YChild{"Tables", &information.Tables})
    information.EntityData.Children.Append("system-oid", types.YChild{"SystemOid", &information.SystemOid})
    information.EntityData.Children.Append("trap-queue", types.YChild{"TrapQueue", &information.TrapQueue})
    information.EntityData.Leafs = types.NewOrderedMap()

    information.EntityData.YListKeys = []string {}

    return &(information.EntityData)
}

// Snmp_Information_Hosts
// SNMP host information
type Snmp_Information_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP target host name. The type is slice of Snmp_Information_Hosts_Host.
    Host []*Snmp_Information_Hosts_Host
}

func (hosts *Snmp_Information_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "information"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + hosts.EntityData.SegmentPath
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// Snmp_Information_Hosts_Host
// SNMP target host name
type Snmp_Information_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Group name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Host name ,udp-port , user, security model and level. The type is slice of
    // Snmp_Information_Hosts_Host_HostInformation.
    HostInformation []*Snmp_Information_Hosts_Host_HostInformation
}

func (host *Snmp_Information_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.Name, "name")
    host.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/hosts/" + host.EntityData.SegmentPath
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Children.Append("host-information", types.YChild{"HostInformation", nil})
    for i := range host.HostInformation {
        host.EntityData.Children.Append(types.GetSegmentPath(host.HostInformation[i]), types.YChild{"HostInformation", host.HostInformation[i]})
    }
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("name", types.YLeaf{"Name", host.Name})

    host.EntityData.YListKeys = []string {"Name"}

    return &(host.EntityData)
}

// Snmp_Information_Hosts_Host_HostInformation
// Host name ,udp-port , user, security model
// and level
type Snmp_Information_Hosts_Host_HostInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. SNMP host user. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    User interface{}

    // Transport type of address. The type is string.
    SnmpTargetAddressTHost interface{}

    // Target UDP port. The type is string.
    SnmpTargetAddressPort interface{}

    // Target host type (Inform or Trap). The type is string.
    SnmpTargetAddresstype interface{}

    // Security model. The type is string.
    SnmpTargetParamsSecurityModel interface{}

    // Security name. The type is string.
    SnmpTargetParamsSecurityName interface{}

    // Security level. The type is string.
    SnmpTargetParamsSecurityLevel interface{}
}

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetEntityData() *types.CommonEntityData {
    hostInformation.EntityData.YFilter = hostInformation.YFilter
    hostInformation.EntityData.YangName = "host-information"
    hostInformation.EntityData.BundleName = "cisco_ios_xr"
    hostInformation.EntityData.ParentYangName = "host"
    hostInformation.EntityData.SegmentPath = "host-information" + types.AddKeyToken(hostInformation.User, "user")
    hostInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/hosts/host/" + hostInformation.EntityData.SegmentPath
    hostInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostInformation.EntityData.Children = types.NewOrderedMap()
    hostInformation.EntityData.Leafs = types.NewOrderedMap()
    hostInformation.EntityData.Leafs.Append("user", types.YLeaf{"User", hostInformation.User})
    hostInformation.EntityData.Leafs.Append("snmp-target-address-t-host", types.YLeaf{"SnmpTargetAddressTHost", hostInformation.SnmpTargetAddressTHost})
    hostInformation.EntityData.Leafs.Append("snmp-target-address-port", types.YLeaf{"SnmpTargetAddressPort", hostInformation.SnmpTargetAddressPort})
    hostInformation.EntityData.Leafs.Append("snmp-target-addresstype", types.YLeaf{"SnmpTargetAddresstype", hostInformation.SnmpTargetAddresstype})
    hostInformation.EntityData.Leafs.Append("snmp-target-params-security-model", types.YLeaf{"SnmpTargetParamsSecurityModel", hostInformation.SnmpTargetParamsSecurityModel})
    hostInformation.EntityData.Leafs.Append("snmp-target-params-security-name", types.YLeaf{"SnmpTargetParamsSecurityName", hostInformation.SnmpTargetParamsSecurityName})
    hostInformation.EntityData.Leafs.Append("snmp-target-params-security-level", types.YLeaf{"SnmpTargetParamsSecurityLevel", hostInformation.SnmpTargetParamsSecurityLevel})

    hostInformation.EntityData.YListKeys = []string {"User"}

    return &(hostInformation.EntityData)
}

// Snmp_Information_SystemUpTime
// System up time
type Snmp_Information_SystemUpTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sysUpTime  1.3.6.1.2.1.1.3. The type is string.
    SystemUpTimeEdm interface{}
}

func (systemUpTime *Snmp_Information_SystemUpTime) GetEntityData() *types.CommonEntityData {
    systemUpTime.EntityData.YFilter = systemUpTime.YFilter
    systemUpTime.EntityData.YangName = "system-up-time"
    systemUpTime.EntityData.BundleName = "cisco_ios_xr"
    systemUpTime.EntityData.ParentYangName = "information"
    systemUpTime.EntityData.SegmentPath = "system-up-time"
    systemUpTime.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + systemUpTime.EntityData.SegmentPath
    systemUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemUpTime.EntityData.Children = types.NewOrderedMap()
    systemUpTime.EntityData.Leafs = types.NewOrderedMap()
    systemUpTime.EntityData.Leafs.Append("system-up-time-edm", types.YLeaf{"SystemUpTimeEdm", systemUpTime.SystemUpTimeEdm})

    systemUpTime.EntityData.YListKeys = []string {}

    return &(systemUpTime.EntityData)
}

// Snmp_Information_NmsAddresses
// SNMP request type summary 
type Snmp_Information_NmsAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NMS address. The type is slice of Snmp_Information_NmsAddresses_NmsAddress.
    NmsAddress []*Snmp_Information_NmsAddresses_NmsAddress
}

func (nmsAddresses *Snmp_Information_NmsAddresses) GetEntityData() *types.CommonEntityData {
    nmsAddresses.EntityData.YFilter = nmsAddresses.YFilter
    nmsAddresses.EntityData.YangName = "nms-addresses"
    nmsAddresses.EntityData.BundleName = "cisco_ios_xr"
    nmsAddresses.EntityData.ParentYangName = "information"
    nmsAddresses.EntityData.SegmentPath = "nms-addresses"
    nmsAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + nmsAddresses.EntityData.SegmentPath
    nmsAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmsAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmsAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmsAddresses.EntityData.Children = types.NewOrderedMap()
    nmsAddresses.EntityData.Children.Append("nms-address", types.YChild{"NmsAddress", nil})
    for i := range nmsAddresses.NmsAddress {
        nmsAddresses.EntityData.Children.Append(types.GetSegmentPath(nmsAddresses.NmsAddress[i]), types.YChild{"NmsAddress", nmsAddresses.NmsAddress[i]})
    }
    nmsAddresses.EntityData.Leafs = types.NewOrderedMap()

    nmsAddresses.EntityData.YListKeys = []string {}

    return &(nmsAddresses.EntityData)
}

// Snmp_Information_NmsAddresses_NmsAddress
// NMS address
type Snmp_Information_NmsAddresses_NmsAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NMS address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    NmsAddr interface{}

    // NMS address of server. The type is string.
    NmsAddress interface{}

    // Get Request Count. The type is interface{} with range: 0..4294967295.
    GetRequestCount interface{}

    // Getnext Request Count. The type is interface{} with range: 0..4294967295.
    GetnextRequestCount interface{}

    // Getbulk Request Count. The type is interface{} with range: 0..4294967295.
    GetbulkRequestCount interface{}

    // Set Request Count. The type is interface{} with range: 0..4294967295.
    SetRequestCount interface{}

    // Test Request Count. The type is interface{} with range: 0..4294967295.
    TestRequestCount interface{}
}

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetEntityData() *types.CommonEntityData {
    nmsAddress.EntityData.YFilter = nmsAddress.YFilter
    nmsAddress.EntityData.YangName = "nms-address"
    nmsAddress.EntityData.BundleName = "cisco_ios_xr"
    nmsAddress.EntityData.ParentYangName = "nms-addresses"
    nmsAddress.EntityData.SegmentPath = "nms-address" + types.AddKeyToken(nmsAddress.NmsAddr, "nms-addr")
    nmsAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/nms-addresses/" + nmsAddress.EntityData.SegmentPath
    nmsAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmsAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmsAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmsAddress.EntityData.Children = types.NewOrderedMap()
    nmsAddress.EntityData.Leafs = types.NewOrderedMap()
    nmsAddress.EntityData.Leafs.Append("nms-addr", types.YLeaf{"NmsAddr", nmsAddress.NmsAddr})
    nmsAddress.EntityData.Leafs.Append("nms-address", types.YLeaf{"NmsAddress", nmsAddress.NmsAddress})
    nmsAddress.EntityData.Leafs.Append("get-request-count", types.YLeaf{"GetRequestCount", nmsAddress.GetRequestCount})
    nmsAddress.EntityData.Leafs.Append("getnext-request-count", types.YLeaf{"GetnextRequestCount", nmsAddress.GetnextRequestCount})
    nmsAddress.EntityData.Leafs.Append("getbulk-request-count", types.YLeaf{"GetbulkRequestCount", nmsAddress.GetbulkRequestCount})
    nmsAddress.EntityData.Leafs.Append("set-request-count", types.YLeaf{"SetRequestCount", nmsAddress.SetRequestCount})
    nmsAddress.EntityData.Leafs.Append("test-request-count", types.YLeaf{"TestRequestCount", nmsAddress.TestRequestCount})

    nmsAddress.EntityData.YListKeys = []string {"NmsAddr"}

    return &(nmsAddress.EntityData)
}

// Snmp_Information_EngineId
// SNMP engine ID
type Snmp_Information_EngineId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMPv3 engineID. The type is string.
    EngineId interface{}
}

func (engineId *Snmp_Information_EngineId) GetEntityData() *types.CommonEntityData {
    engineId.EntityData.YFilter = engineId.YFilter
    engineId.EntityData.YangName = "engine-id"
    engineId.EntityData.BundleName = "cisco_ios_xr"
    engineId.EntityData.ParentYangName = "information"
    engineId.EntityData.SegmentPath = "engine-id"
    engineId.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + engineId.EntityData.SegmentPath
    engineId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    engineId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    engineId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    engineId.EntityData.Children = types.NewOrderedMap()
    engineId.EntityData.Leafs = types.NewOrderedMap()
    engineId.EntityData.Leafs.Append("engine-id", types.YLeaf{"EngineId", engineId.EngineId})

    engineId.EntityData.YListKeys = []string {}

    return &(engineId.EntityData)
}

// Snmp_Information_RxQueue
// SNMP rx queue statistics
type Snmp_Information_RxQueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // qlen. The type is interface{} with range: 0..4294967295.
    Qlen interface{}

    // in min. The type is interface{} with range: 0..4294967295.
    InMin interface{}

    // in avg. The type is interface{} with range: 0..4294967295.
    InAvg interface{}

    // in max. The type is interface{} with range: 0..4294967295.
    InMax interface{}

    // pend min. The type is interface{} with range: 0..4294967295.
    PendMin interface{}

    // pend avg. The type is interface{} with range: 0..4294967295.
    PendAvg interface{}

    // pend max. The type is interface{} with range: 0..4294967295.
    PendMax interface{}

    // incoming q. The type is slice of Snmp_Information_RxQueue_IncomingQ.
    IncomingQ []*Snmp_Information_RxQueue_IncomingQ

    // pending q. The type is slice of Snmp_Information_RxQueue_PendingQ.
    PendingQ []*Snmp_Information_RxQueue_PendingQ
}

func (rxQueue *Snmp_Information_RxQueue) GetEntityData() *types.CommonEntityData {
    rxQueue.EntityData.YFilter = rxQueue.YFilter
    rxQueue.EntityData.YangName = "rx-queue"
    rxQueue.EntityData.BundleName = "cisco_ios_xr"
    rxQueue.EntityData.ParentYangName = "information"
    rxQueue.EntityData.SegmentPath = "rx-queue"
    rxQueue.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + rxQueue.EntityData.SegmentPath
    rxQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxQueue.EntityData.Children = types.NewOrderedMap()
    rxQueue.EntityData.Children.Append("incoming-q", types.YChild{"IncomingQ", nil})
    for i := range rxQueue.IncomingQ {
        types.SetYListKey(rxQueue.IncomingQ[i], i)
        rxQueue.EntityData.Children.Append(types.GetSegmentPath(rxQueue.IncomingQ[i]), types.YChild{"IncomingQ", rxQueue.IncomingQ[i]})
    }
    rxQueue.EntityData.Children.Append("pending-q", types.YChild{"PendingQ", nil})
    for i := range rxQueue.PendingQ {
        types.SetYListKey(rxQueue.PendingQ[i], i)
        rxQueue.EntityData.Children.Append(types.GetSegmentPath(rxQueue.PendingQ[i]), types.YChild{"PendingQ", rxQueue.PendingQ[i]})
    }
    rxQueue.EntityData.Leafs = types.NewOrderedMap()
    rxQueue.EntityData.Leafs.Append("qlen", types.YLeaf{"Qlen", rxQueue.Qlen})
    rxQueue.EntityData.Leafs.Append("in-min", types.YLeaf{"InMin", rxQueue.InMin})
    rxQueue.EntityData.Leafs.Append("in-avg", types.YLeaf{"InAvg", rxQueue.InAvg})
    rxQueue.EntityData.Leafs.Append("in-max", types.YLeaf{"InMax", rxQueue.InMax})
    rxQueue.EntityData.Leafs.Append("pend-min", types.YLeaf{"PendMin", rxQueue.PendMin})
    rxQueue.EntityData.Leafs.Append("pend-avg", types.YLeaf{"PendAvg", rxQueue.PendAvg})
    rxQueue.EntityData.Leafs.Append("pend-max", types.YLeaf{"PendMax", rxQueue.PendMax})

    rxQueue.EntityData.YListKeys = []string {}

    return &(rxQueue.EntityData)
}

// Snmp_Information_RxQueue_IncomingQ
// incoming q
type Snmp_Information_RxQueue_IncomingQ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // min. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // avg. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // max. The type is interface{} with range: 0..4294967295.
    Max interface{}
}

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetEntityData() *types.CommonEntityData {
    incomingQ.EntityData.YFilter = incomingQ.YFilter
    incomingQ.EntityData.YangName = "incoming-q"
    incomingQ.EntityData.BundleName = "cisco_ios_xr"
    incomingQ.EntityData.ParentYangName = "rx-queue"
    incomingQ.EntityData.SegmentPath = "incoming-q" + types.AddNoKeyToken(incomingQ)
    incomingQ.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/rx-queue/" + incomingQ.EntityData.SegmentPath
    incomingQ.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incomingQ.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incomingQ.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incomingQ.EntityData.Children = types.NewOrderedMap()
    incomingQ.EntityData.Leafs = types.NewOrderedMap()
    incomingQ.EntityData.Leafs.Append("min", types.YLeaf{"Min", incomingQ.Min})
    incomingQ.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", incomingQ.Avg})
    incomingQ.EntityData.Leafs.Append("max", types.YLeaf{"Max", incomingQ.Max})

    incomingQ.EntityData.YListKeys = []string {}

    return &(incomingQ.EntityData)
}

// Snmp_Information_RxQueue_PendingQ
// pending q
type Snmp_Information_RxQueue_PendingQ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // min. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // avg. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // max. The type is interface{} with range: 0..4294967295.
    Max interface{}
}

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetEntityData() *types.CommonEntityData {
    pendingQ.EntityData.YFilter = pendingQ.YFilter
    pendingQ.EntityData.YangName = "pending-q"
    pendingQ.EntityData.BundleName = "cisco_ios_xr"
    pendingQ.EntityData.ParentYangName = "rx-queue"
    pendingQ.EntityData.SegmentPath = "pending-q" + types.AddNoKeyToken(pendingQ)
    pendingQ.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/rx-queue/" + pendingQ.EntityData.SegmentPath
    pendingQ.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pendingQ.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pendingQ.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pendingQ.EntityData.Children = types.NewOrderedMap()
    pendingQ.EntityData.Leafs = types.NewOrderedMap()
    pendingQ.EntityData.Leafs.Append("min", types.YLeaf{"Min", pendingQ.Min})
    pendingQ.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", pendingQ.Avg})
    pendingQ.EntityData.Leafs.Append("max", types.YLeaf{"Max", pendingQ.Max})

    pendingQ.EntityData.YListKeys = []string {}

    return &(pendingQ.EntityData)
}

// Snmp_Information_SystemName
// System name
type Snmp_Information_SystemName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sysName  1.3.6.1.2.1.1.5. The type is string.
    SystemName interface{}
}

func (systemName *Snmp_Information_SystemName) GetEntityData() *types.CommonEntityData {
    systemName.EntityData.YFilter = systemName.YFilter
    systemName.EntityData.YangName = "system-name"
    systemName.EntityData.BundleName = "cisco_ios_xr"
    systemName.EntityData.ParentYangName = "information"
    systemName.EntityData.SegmentPath = "system-name"
    systemName.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + systemName.EntityData.SegmentPath
    systemName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemName.EntityData.Children = types.NewOrderedMap()
    systemName.EntityData.Leafs = types.NewOrderedMap()
    systemName.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", systemName.SystemName})

    systemName.EntityData.YListKeys = []string {}

    return &(systemName.EntityData)
}

// Snmp_Information_RequestTypeDetail
// SNMP request type details 
type Snmp_Information_RequestTypeDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // snmp request type details .
    NmsAddresses Snmp_Information_RequestTypeDetail_NmsAddresses
}

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetEntityData() *types.CommonEntityData {
    requestTypeDetail.EntityData.YFilter = requestTypeDetail.YFilter
    requestTypeDetail.EntityData.YangName = "request-type-detail"
    requestTypeDetail.EntityData.BundleName = "cisco_ios_xr"
    requestTypeDetail.EntityData.ParentYangName = "information"
    requestTypeDetail.EntityData.SegmentPath = "request-type-detail"
    requestTypeDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + requestTypeDetail.EntityData.SegmentPath
    requestTypeDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requestTypeDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requestTypeDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requestTypeDetail.EntityData.Children = types.NewOrderedMap()
    requestTypeDetail.EntityData.Children.Append("nms-addresses", types.YChild{"NmsAddresses", &requestTypeDetail.NmsAddresses})
    requestTypeDetail.EntityData.Leafs = types.NewOrderedMap()

    requestTypeDetail.EntityData.YListKeys = []string {}

    return &(requestTypeDetail.EntityData)
}

// Snmp_Information_RequestTypeDetail_NmsAddresses
// snmp request type details 
type Snmp_Information_RequestTypeDetail_NmsAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NMS address. The type is slice of
    // Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress.
    NmsAddress []*Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress
}

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetEntityData() *types.CommonEntityData {
    nmsAddresses.EntityData.YFilter = nmsAddresses.YFilter
    nmsAddresses.EntityData.YangName = "nms-addresses"
    nmsAddresses.EntityData.BundleName = "cisco_ios_xr"
    nmsAddresses.EntityData.ParentYangName = "request-type-detail"
    nmsAddresses.EntityData.SegmentPath = "nms-addresses"
    nmsAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/request-type-detail/" + nmsAddresses.EntityData.SegmentPath
    nmsAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmsAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmsAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmsAddresses.EntityData.Children = types.NewOrderedMap()
    nmsAddresses.EntityData.Children.Append("nms-address", types.YChild{"NmsAddress", nil})
    for i := range nmsAddresses.NmsAddress {
        nmsAddresses.EntityData.Children.Append(types.GetSegmentPath(nmsAddresses.NmsAddress[i]), types.YChild{"NmsAddress", nmsAddresses.NmsAddress[i]})
    }
    nmsAddresses.EntityData.Leafs = types.NewOrderedMap()

    nmsAddresses.EntityData.YListKeys = []string {}

    return &(nmsAddresses.EntityData)
}

// Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress
// NMS address
type Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NMS address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    NmsAddr interface{}

    // Total request count for each managment station or client. The type is
    // interface{} with range: 0..4294967295.
    TotalCount interface{}

    // Processing agent request count for each client for particluar managment
    // station. The type is interface{} with range: 0..4294967295.
    AgentRequestCount interface{}

    // Processing interfce request count for each client for particluar managment
    // station. The type is interface{} with range: 0..4294967295.
    InterfaceRequestCount interface{}

    // Processing entity request count for each client for particluar managment
    // station. The type is interface{} with range: 0..4294967295.
    EntityRequestCount interface{}

    // Processing route request count for each client for particluar Managment
    // station. The type is interface{} with range: 0..4294967295.
    RouteRequestCount interface{}

    // Processing infra request count for each client for particluar Managment
    // station. The type is interface{} with range: 0..4294967295.
    InfraRequestCount interface{}
}

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetEntityData() *types.CommonEntityData {
    nmsAddress.EntityData.YFilter = nmsAddress.YFilter
    nmsAddress.EntityData.YangName = "nms-address"
    nmsAddress.EntityData.BundleName = "cisco_ios_xr"
    nmsAddress.EntityData.ParentYangName = "nms-addresses"
    nmsAddress.EntityData.SegmentPath = "nms-address" + types.AddKeyToken(nmsAddress.NmsAddr, "nms-addr")
    nmsAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/request-type-detail/nms-addresses/" + nmsAddress.EntityData.SegmentPath
    nmsAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmsAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmsAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmsAddress.EntityData.Children = types.NewOrderedMap()
    nmsAddress.EntityData.Leafs = types.NewOrderedMap()
    nmsAddress.EntityData.Leafs.Append("nms-addr", types.YLeaf{"NmsAddr", nmsAddress.NmsAddr})
    nmsAddress.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", nmsAddress.TotalCount})
    nmsAddress.EntityData.Leafs.Append("agent-request-count", types.YLeaf{"AgentRequestCount", nmsAddress.AgentRequestCount})
    nmsAddress.EntityData.Leafs.Append("interface-request-count", types.YLeaf{"InterfaceRequestCount", nmsAddress.InterfaceRequestCount})
    nmsAddress.EntityData.Leafs.Append("entity-request-count", types.YLeaf{"EntityRequestCount", nmsAddress.EntityRequestCount})
    nmsAddress.EntityData.Leafs.Append("route-request-count", types.YLeaf{"RouteRequestCount", nmsAddress.RouteRequestCount})
    nmsAddress.EntityData.Leafs.Append("infra-request-count", types.YLeaf{"InfraRequestCount", nmsAddress.InfraRequestCount})

    nmsAddress.EntityData.YListKeys = []string {"NmsAddr"}

    return &(nmsAddress.EntityData)
}

// Snmp_Information_DuplicateDrop
// Duplicate request status, count, time 
type Snmp_Information_DuplicateDrop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Duplicate requests drop feature status. The type is DupReqDropStatus.
    DuplicateRequestStatus interface{}

    // Duplicate request drop feature last enable disable time (Day Mon Date
    // HH:MM:SS). The type is string.
    LastStatusChangeTime interface{}

    // Configured Duplicate Drop feature Timeout. The type is interface{} with
    // range: 0..4294967295.
    DuplicateDropConfiguredTimeout interface{}

    // Number of duplicate SNMP requests are dropped. The type is interface{} with
    // range: 0..4294967295.
    DuplicateDroppedRequests interface{}

    // Number of Retry SNMP requests are Processed. The type is interface{} with
    // range: 0..4294967295.
    RetryProcessedRequests interface{}

    // Duplicate request drop feature first  enable time (Day Mon Date HH:MM:SS).
    // The type is string.
    FirstEnableTime interface{}

    // Number of duplicate SNMP requests dropped, from the last enable time. The
    // type is interface{} with range: 0..4294967295.
    LatestDuplicateDroppedRequests interface{}

    // Number of retry SNMP requests processed, from the last enable time. The
    // type is interface{} with range: 0..4294967295.
    LatestRetryProcessedRequests interface{}

    // Duplicate request drop feature last enable time(Day Mon Date HH:MM:SS). The
    // type is string.
    DuplicateRequestLatestEnableTime interface{}

    // Number of times duplicate request drop feature is enabled. The type is
    // interface{} with range: 0..4294967295.
    DuplicateDropEnableCount interface{}

    // Number of times duplicate request drop feature is disabled. The type is
    // interface{} with range: 0..4294967295.
    DuplicateDropDisableCount interface{}
}

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetEntityData() *types.CommonEntityData {
    duplicateDrop.EntityData.YFilter = duplicateDrop.YFilter
    duplicateDrop.EntityData.YangName = "duplicate-drop"
    duplicateDrop.EntityData.BundleName = "cisco_ios_xr"
    duplicateDrop.EntityData.ParentYangName = "information"
    duplicateDrop.EntityData.SegmentPath = "duplicate-drop"
    duplicateDrop.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + duplicateDrop.EntityData.SegmentPath
    duplicateDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duplicateDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duplicateDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duplicateDrop.EntityData.Children = types.NewOrderedMap()
    duplicateDrop.EntityData.Leafs = types.NewOrderedMap()
    duplicateDrop.EntityData.Leafs.Append("duplicate-request-status", types.YLeaf{"DuplicateRequestStatus", duplicateDrop.DuplicateRequestStatus})
    duplicateDrop.EntityData.Leafs.Append("last-status-change-time", types.YLeaf{"LastStatusChangeTime", duplicateDrop.LastStatusChangeTime})
    duplicateDrop.EntityData.Leafs.Append("duplicate-drop-configured-timeout", types.YLeaf{"DuplicateDropConfiguredTimeout", duplicateDrop.DuplicateDropConfiguredTimeout})
    duplicateDrop.EntityData.Leafs.Append("duplicate-dropped-requests", types.YLeaf{"DuplicateDroppedRequests", duplicateDrop.DuplicateDroppedRequests})
    duplicateDrop.EntityData.Leafs.Append("retry-processed-requests", types.YLeaf{"RetryProcessedRequests", duplicateDrop.RetryProcessedRequests})
    duplicateDrop.EntityData.Leafs.Append("first-enable-time", types.YLeaf{"FirstEnableTime", duplicateDrop.FirstEnableTime})
    duplicateDrop.EntityData.Leafs.Append("latest-duplicate-dropped-requests", types.YLeaf{"LatestDuplicateDroppedRequests", duplicateDrop.LatestDuplicateDroppedRequests})
    duplicateDrop.EntityData.Leafs.Append("latest-retry-processed-requests", types.YLeaf{"LatestRetryProcessedRequests", duplicateDrop.LatestRetryProcessedRequests})
    duplicateDrop.EntityData.Leafs.Append("duplicate-request-latest-enable-time", types.YLeaf{"DuplicateRequestLatestEnableTime", duplicateDrop.DuplicateRequestLatestEnableTime})
    duplicateDrop.EntityData.Leafs.Append("duplicate-drop-enable-count", types.YLeaf{"DuplicateDropEnableCount", duplicateDrop.DuplicateDropEnableCount})
    duplicateDrop.EntityData.Leafs.Append("duplicate-drop-disable-count", types.YLeaf{"DuplicateDropDisableCount", duplicateDrop.DuplicateDropDisableCount})

    duplicateDrop.EntityData.YListKeys = []string {}

    return &(duplicateDrop.EntityData)
}

// Snmp_Information_BulkStatsTransfers
// List of bulkstats transfer on the system
type Snmp_Information_BulkStatsTransfers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP bulkstats transfer name. The type is slice of
    // Snmp_Information_BulkStatsTransfers_BulkStatsTransfer.
    BulkStatsTransfer []*Snmp_Information_BulkStatsTransfers_BulkStatsTransfer
}

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetEntityData() *types.CommonEntityData {
    bulkStatsTransfers.EntityData.YFilter = bulkStatsTransfers.YFilter
    bulkStatsTransfers.EntityData.YangName = "bulk-stats-transfers"
    bulkStatsTransfers.EntityData.BundleName = "cisco_ios_xr"
    bulkStatsTransfers.EntityData.ParentYangName = "information"
    bulkStatsTransfers.EntityData.SegmentPath = "bulk-stats-transfers"
    bulkStatsTransfers.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + bulkStatsTransfers.EntityData.SegmentPath
    bulkStatsTransfers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bulkStatsTransfers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bulkStatsTransfers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bulkStatsTransfers.EntityData.Children = types.NewOrderedMap()
    bulkStatsTransfers.EntityData.Children.Append("bulk-stats-transfer", types.YChild{"BulkStatsTransfer", nil})
    for i := range bulkStatsTransfers.BulkStatsTransfer {
        bulkStatsTransfers.EntityData.Children.Append(types.GetSegmentPath(bulkStatsTransfers.BulkStatsTransfer[i]), types.YChild{"BulkStatsTransfer", bulkStatsTransfers.BulkStatsTransfer[i]})
    }
    bulkStatsTransfers.EntityData.Leafs = types.NewOrderedMap()

    bulkStatsTransfers.EntityData.YListKeys = []string {}

    return &(bulkStatsTransfers.EntityData)
}

// Snmp_Information_BulkStatsTransfers_BulkStatsTransfer
// SNMP bulkstats transfer name
type Snmp_Information_BulkStatsTransfers_BulkStatsTransfer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Transfer name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TransferName interface{}

    // Name of the bulkstats transfer session. The type is string.
    TransferNameXr interface{}

    // Bulkstats transfer primary URL. The type is string.
    UrlPrimary interface{}

    // Bulkstats transfer secondary URL. The type is string.
    UrlSecondary interface{}

    // Bulkstats transfer retained file name. The type is string.
    RetainedFile interface{}

    // Bulkstats transfer retry time left in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    TimeLeft interface{}

    // Bulkstats transfer retry attempt left. The type is interface{} with range:
    // 0..4294967295.
    RetryLeft interface{}
}

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetEntityData() *types.CommonEntityData {
    bulkStatsTransfer.EntityData.YFilter = bulkStatsTransfer.YFilter
    bulkStatsTransfer.EntityData.YangName = "bulk-stats-transfer"
    bulkStatsTransfer.EntityData.BundleName = "cisco_ios_xr"
    bulkStatsTransfer.EntityData.ParentYangName = "bulk-stats-transfers"
    bulkStatsTransfer.EntityData.SegmentPath = "bulk-stats-transfer" + types.AddKeyToken(bulkStatsTransfer.TransferName, "transfer-name")
    bulkStatsTransfer.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/bulk-stats-transfers/" + bulkStatsTransfer.EntityData.SegmentPath
    bulkStatsTransfer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bulkStatsTransfer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bulkStatsTransfer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bulkStatsTransfer.EntityData.Children = types.NewOrderedMap()
    bulkStatsTransfer.EntityData.Leafs = types.NewOrderedMap()
    bulkStatsTransfer.EntityData.Leafs.Append("transfer-name", types.YLeaf{"TransferName", bulkStatsTransfer.TransferName})
    bulkStatsTransfer.EntityData.Leafs.Append("transfer-name-xr", types.YLeaf{"TransferNameXr", bulkStatsTransfer.TransferNameXr})
    bulkStatsTransfer.EntityData.Leafs.Append("url-primary", types.YLeaf{"UrlPrimary", bulkStatsTransfer.UrlPrimary})
    bulkStatsTransfer.EntityData.Leafs.Append("url-secondary", types.YLeaf{"UrlSecondary", bulkStatsTransfer.UrlSecondary})
    bulkStatsTransfer.EntityData.Leafs.Append("retained-file", types.YLeaf{"RetainedFile", bulkStatsTransfer.RetainedFile})
    bulkStatsTransfer.EntityData.Leafs.Append("time-left", types.YLeaf{"TimeLeft", bulkStatsTransfer.TimeLeft})
    bulkStatsTransfer.EntityData.Leafs.Append("retry-left", types.YLeaf{"RetryLeft", bulkStatsTransfer.RetryLeft})

    bulkStatsTransfer.EntityData.YListKeys = []string {"TransferName"}

    return &(bulkStatsTransfer.EntityData)
}

// Snmp_Information_TrapInfos
// SNMP trap OID
type Snmp_Information_TrapInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP Trap infomation like server , port and trapOID. The type is slice of
    // Snmp_Information_TrapInfos_TrapInfo.
    TrapInfo []*Snmp_Information_TrapInfos_TrapInfo
}

func (trapInfos *Snmp_Information_TrapInfos) GetEntityData() *types.CommonEntityData {
    trapInfos.EntityData.YFilter = trapInfos.YFilter
    trapInfos.EntityData.YangName = "trap-infos"
    trapInfos.EntityData.BundleName = "cisco_ios_xr"
    trapInfos.EntityData.ParentYangName = "information"
    trapInfos.EntityData.SegmentPath = "trap-infos"
    trapInfos.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + trapInfos.EntityData.SegmentPath
    trapInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapInfos.EntityData.Children = types.NewOrderedMap()
    trapInfos.EntityData.Children.Append("trap-info", types.YChild{"TrapInfo", nil})
    for i := range trapInfos.TrapInfo {
        types.SetYListKey(trapInfos.TrapInfo[i], i)
        trapInfos.EntityData.Children.Append(types.GetSegmentPath(trapInfos.TrapInfo[i]), types.YChild{"TrapInfo", trapInfos.TrapInfo[i]})
    }
    trapInfos.EntityData.Leafs = types.NewOrderedMap()

    trapInfos.EntityData.YListKeys = []string {}

    return &(trapInfos.EntityData)
}

// Snmp_Information_TrapInfos_TrapInfo
// SNMP Trap infomation like server , port and
// trapOID
type Snmp_Information_TrapInfos_TrapInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trap Host. The type is string.
    TrapHost interface{}

    // Trap port. The type is interface{} with range: 0..65535.
    Port interface{}

    // NMS/Host address. The type is string.
    Host interface{}

    // udp port number. The type is interface{} with range: 0..65535.
    PortXr interface{}

    // Total number of OID's sent. The type is interface{} with range:
    // 0..4294967295.
    TrapOidCount interface{}

    // Per trap OID statistics. The type is slice of
    // Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo.
    TrapOiDinfo []*Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo
}

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetEntityData() *types.CommonEntityData {
    trapInfo.EntityData.YFilter = trapInfo.YFilter
    trapInfo.EntityData.YangName = "trap-info"
    trapInfo.EntityData.BundleName = "cisco_ios_xr"
    trapInfo.EntityData.ParentYangName = "trap-infos"
    trapInfo.EntityData.SegmentPath = "trap-info" + types.AddNoKeyToken(trapInfo)
    trapInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/trap-infos/" + trapInfo.EntityData.SegmentPath
    trapInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapInfo.EntityData.Children = types.NewOrderedMap()
    trapInfo.EntityData.Children.Append("trap-oi-dinfo", types.YChild{"TrapOiDinfo", nil})
    for i := range trapInfo.TrapOiDinfo {
        types.SetYListKey(trapInfo.TrapOiDinfo[i], i)
        trapInfo.EntityData.Children.Append(types.GetSegmentPath(trapInfo.TrapOiDinfo[i]), types.YChild{"TrapOiDinfo", trapInfo.TrapOiDinfo[i]})
    }
    trapInfo.EntityData.Leafs = types.NewOrderedMap()
    trapInfo.EntityData.Leafs.Append("trap-host", types.YLeaf{"TrapHost", trapInfo.TrapHost})
    trapInfo.EntityData.Leafs.Append("port", types.YLeaf{"Port", trapInfo.Port})
    trapInfo.EntityData.Leafs.Append("host", types.YLeaf{"Host", trapInfo.Host})
    trapInfo.EntityData.Leafs.Append("port-xr", types.YLeaf{"PortXr", trapInfo.PortXr})
    trapInfo.EntityData.Leafs.Append("trap-oid-count", types.YLeaf{"TrapOidCount", trapInfo.TrapOidCount})

    trapInfo.EntityData.YListKeys = []string {}

    return &(trapInfo.EntityData)
}

// Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo
// Per trap OID statistics
type Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // TRAP OID. The type is string.
    TrapOid interface{}

    // Number of traps sent. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Number of Traps Dropped. The type is interface{} with range: 0..4294967295.
    DropCount interface{}

    // Num of times retry. The type is interface{} with range: 0..4294967295.
    RetryCount interface{}

    // Timestamp of latest successfully sent. The type is string.
    LastsentTime interface{}

    // Timestamp of latest droped. The type is string.
    LasrdropTime interface{}
}

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetEntityData() *types.CommonEntityData {
    trapOiDinfo.EntityData.YFilter = trapOiDinfo.YFilter
    trapOiDinfo.EntityData.YangName = "trap-oi-dinfo"
    trapOiDinfo.EntityData.BundleName = "cisco_ios_xr"
    trapOiDinfo.EntityData.ParentYangName = "trap-info"
    trapOiDinfo.EntityData.SegmentPath = "trap-oi-dinfo" + types.AddNoKeyToken(trapOiDinfo)
    trapOiDinfo.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/trap-infos/trap-info/" + trapOiDinfo.EntityData.SegmentPath
    trapOiDinfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapOiDinfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapOiDinfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapOiDinfo.EntityData.Children = types.NewOrderedMap()
    trapOiDinfo.EntityData.Leafs = types.NewOrderedMap()
    trapOiDinfo.EntityData.Leafs.Append("trap-oid", types.YLeaf{"TrapOid", trapOiDinfo.TrapOid})
    trapOiDinfo.EntityData.Leafs.Append("count", types.YLeaf{"Count", trapOiDinfo.Count})
    trapOiDinfo.EntityData.Leafs.Append("drop-count", types.YLeaf{"DropCount", trapOiDinfo.DropCount})
    trapOiDinfo.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", trapOiDinfo.RetryCount})
    trapOiDinfo.EntityData.Leafs.Append("lastsent-time", types.YLeaf{"LastsentTime", trapOiDinfo.LastsentTime})
    trapOiDinfo.EntityData.Leafs.Append("lasrdrop-time", types.YLeaf{"LasrdropTime", trapOiDinfo.LasrdropTime})

    trapOiDinfo.EntityData.YListKeys = []string {}

    return &(trapOiDinfo.EntityData)
}

// Snmp_Information_PollOids
// OID list for poll PDU
type Snmp_Information_PollOids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PDU drop info for OID. The type is slice of
    // Snmp_Information_PollOids_PollOid.
    PollOid []*Snmp_Information_PollOids_PollOid
}

func (pollOids *Snmp_Information_PollOids) GetEntityData() *types.CommonEntityData {
    pollOids.EntityData.YFilter = pollOids.YFilter
    pollOids.EntityData.YangName = "poll-oids"
    pollOids.EntityData.BundleName = "cisco_ios_xr"
    pollOids.EntityData.ParentYangName = "information"
    pollOids.EntityData.SegmentPath = "poll-oids"
    pollOids.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + pollOids.EntityData.SegmentPath
    pollOids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pollOids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pollOids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pollOids.EntityData.Children = types.NewOrderedMap()
    pollOids.EntityData.Children.Append("poll-oid", types.YChild{"PollOid", nil})
    for i := range pollOids.PollOid {
        pollOids.EntityData.Children.Append(types.GetSegmentPath(pollOids.PollOid[i]), types.YChild{"PollOid", pollOids.PollOid[i]})
    }
    pollOids.EntityData.Leafs = types.NewOrderedMap()

    pollOids.EntityData.YListKeys = []string {}

    return &(pollOids.EntityData)
}

// Snmp_Information_PollOids_PollOid
// PDU drop info for OID
type Snmp_Information_PollOids_PollOid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ObjectId interface{}

    // Managment station count. The type is interface{} with range: 0..4294967295.
    NmsCount interface{}

    // Network Managment station ipadress. The type is slice of string.
    Nms []interface{}

    // OID request count for each Managment station . The type is slice of
    // interface{} with range: 0..4294967295.
    RequestCount []interface{}
}

func (pollOid *Snmp_Information_PollOids_PollOid) GetEntityData() *types.CommonEntityData {
    pollOid.EntityData.YFilter = pollOid.YFilter
    pollOid.EntityData.YangName = "poll-oid"
    pollOid.EntityData.BundleName = "cisco_ios_xr"
    pollOid.EntityData.ParentYangName = "poll-oids"
    pollOid.EntityData.SegmentPath = "poll-oid" + types.AddKeyToken(pollOid.ObjectId, "object-id")
    pollOid.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/poll-oids/" + pollOid.EntityData.SegmentPath
    pollOid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pollOid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pollOid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pollOid.EntityData.Children = types.NewOrderedMap()
    pollOid.EntityData.Leafs = types.NewOrderedMap()
    pollOid.EntityData.Leafs.Append("object-id", types.YLeaf{"ObjectId", pollOid.ObjectId})
    pollOid.EntityData.Leafs.Append("nms-count", types.YLeaf{"NmsCount", pollOid.NmsCount})
    pollOid.EntityData.Leafs.Append("nms", types.YLeaf{"Nms", pollOid.Nms})
    pollOid.EntityData.Leafs.Append("request-count", types.YLeaf{"RequestCount", pollOid.RequestCount})

    pollOid.EntityData.YListKeys = []string {"ObjectId"}

    return &(pollOid.EntityData)
}

// Snmp_Information_InfomDetails
// SNMP trap OID
type Snmp_Information_InfomDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP Trap infomation like server , port and trapOID. The type is slice of
    // Snmp_Information_InfomDetails_InfomDetail.
    InfomDetail []*Snmp_Information_InfomDetails_InfomDetail
}

func (infomDetails *Snmp_Information_InfomDetails) GetEntityData() *types.CommonEntityData {
    infomDetails.EntityData.YFilter = infomDetails.YFilter
    infomDetails.EntityData.YangName = "infom-details"
    infomDetails.EntityData.BundleName = "cisco_ios_xr"
    infomDetails.EntityData.ParentYangName = "information"
    infomDetails.EntityData.SegmentPath = "infom-details"
    infomDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + infomDetails.EntityData.SegmentPath
    infomDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infomDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infomDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infomDetails.EntityData.Children = types.NewOrderedMap()
    infomDetails.EntityData.Children.Append("infom-detail", types.YChild{"InfomDetail", nil})
    for i := range infomDetails.InfomDetail {
        types.SetYListKey(infomDetails.InfomDetail[i], i)
        infomDetails.EntityData.Children.Append(types.GetSegmentPath(infomDetails.InfomDetail[i]), types.YChild{"InfomDetail", infomDetails.InfomDetail[i]})
    }
    infomDetails.EntityData.Leafs = types.NewOrderedMap()

    infomDetails.EntityData.YListKeys = []string {}

    return &(infomDetails.EntityData)
}

// Snmp_Information_InfomDetails_InfomDetail
// SNMP Trap infomation like server , port and
// trapOID
type Snmp_Information_InfomDetails_InfomDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trap Host. The type is string.
    TrapHost interface{}

    // Trap port. The type is interface{} with range: 0..65535.
    Port interface{}

    // NMS/Host address. The type is string.
    Host interface{}

    // udp port number. The type is interface{} with range: 0..65535.
    PortXr interface{}

    // Total number of OID's sent. The type is interface{} with range:
    // 0..4294967295.
    TrapOidCount interface{}

    // Per trap OID statistics. The type is slice of
    // Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo.
    TrapOiDinfo []*Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo
}

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetEntityData() *types.CommonEntityData {
    infomDetail.EntityData.YFilter = infomDetail.YFilter
    infomDetail.EntityData.YangName = "infom-detail"
    infomDetail.EntityData.BundleName = "cisco_ios_xr"
    infomDetail.EntityData.ParentYangName = "infom-details"
    infomDetail.EntityData.SegmentPath = "infom-detail" + types.AddNoKeyToken(infomDetail)
    infomDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/infom-details/" + infomDetail.EntityData.SegmentPath
    infomDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infomDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infomDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infomDetail.EntityData.Children = types.NewOrderedMap()
    infomDetail.EntityData.Children.Append("trap-oi-dinfo", types.YChild{"TrapOiDinfo", nil})
    for i := range infomDetail.TrapOiDinfo {
        types.SetYListKey(infomDetail.TrapOiDinfo[i], i)
        infomDetail.EntityData.Children.Append(types.GetSegmentPath(infomDetail.TrapOiDinfo[i]), types.YChild{"TrapOiDinfo", infomDetail.TrapOiDinfo[i]})
    }
    infomDetail.EntityData.Leafs = types.NewOrderedMap()
    infomDetail.EntityData.Leafs.Append("trap-host", types.YLeaf{"TrapHost", infomDetail.TrapHost})
    infomDetail.EntityData.Leafs.Append("port", types.YLeaf{"Port", infomDetail.Port})
    infomDetail.EntityData.Leafs.Append("host", types.YLeaf{"Host", infomDetail.Host})
    infomDetail.EntityData.Leafs.Append("port-xr", types.YLeaf{"PortXr", infomDetail.PortXr})
    infomDetail.EntityData.Leafs.Append("trap-oid-count", types.YLeaf{"TrapOidCount", infomDetail.TrapOidCount})

    infomDetail.EntityData.YListKeys = []string {}

    return &(infomDetail.EntityData)
}

// Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo
// Per trap OID statistics
type Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // TRAP OID. The type is string.
    TrapOid interface{}

    // Number of traps sent. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Number of Traps Dropped. The type is interface{} with range: 0..4294967295.
    DropCount interface{}

    // Num of times retry. The type is interface{} with range: 0..4294967295.
    RetryCount interface{}

    // Timestamp of latest successfully sent. The type is string.
    LastsentTime interface{}

    // Timestamp of latest droped. The type is string.
    LasrdropTime interface{}
}

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetEntityData() *types.CommonEntityData {
    trapOiDinfo.EntityData.YFilter = trapOiDinfo.YFilter
    trapOiDinfo.EntityData.YangName = "trap-oi-dinfo"
    trapOiDinfo.EntityData.BundleName = "cisco_ios_xr"
    trapOiDinfo.EntityData.ParentYangName = "infom-detail"
    trapOiDinfo.EntityData.SegmentPath = "trap-oi-dinfo" + types.AddNoKeyToken(trapOiDinfo)
    trapOiDinfo.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/infom-details/infom-detail/" + trapOiDinfo.EntityData.SegmentPath
    trapOiDinfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapOiDinfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapOiDinfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapOiDinfo.EntityData.Children = types.NewOrderedMap()
    trapOiDinfo.EntityData.Leafs = types.NewOrderedMap()
    trapOiDinfo.EntityData.Leafs.Append("trap-oid", types.YLeaf{"TrapOid", trapOiDinfo.TrapOid})
    trapOiDinfo.EntityData.Leafs.Append("count", types.YLeaf{"Count", trapOiDinfo.Count})
    trapOiDinfo.EntityData.Leafs.Append("drop-count", types.YLeaf{"DropCount", trapOiDinfo.DropCount})
    trapOiDinfo.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", trapOiDinfo.RetryCount})
    trapOiDinfo.EntityData.Leafs.Append("lastsent-time", types.YLeaf{"LastsentTime", trapOiDinfo.LastsentTime})
    trapOiDinfo.EntityData.Leafs.Append("lasrdrop-time", types.YLeaf{"LasrdropTime", trapOiDinfo.LasrdropTime})

    trapOiDinfo.EntityData.YListKeys = []string {}

    return &(trapOiDinfo.EntityData)
}

// Snmp_Information_Statistics
// SNMP statistics
type Snmp_Information_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // snmpInPkts. The type is interface{} with range: 0..4294967295.
    PacketsReceived interface{}

    // snmpInBadVersions. The type is interface{} with range: 0..4294967295.
    BadVersionsReceived interface{}

    // snmpInBadCommunityNames. The type is interface{} with range: 0..4294967295.
    BadCommunityNamesReceived interface{}

    // snmpInBadCommunityUses. The type is interface{} with range: 0..4294967295.
    BadCommunityUsesReceived interface{}

    // snmpInASNParseErrs. The type is interface{} with range: 0..4294967295.
    AsnParseErrorsReceived interface{}

    // snmpSilentDrops. The type is interface{} with range: 0..4294967295.
    SilentDropCount interface{}

    // snmpProxyDrops. The type is interface{} with range: 0..4294967295.
    ProxyDropCount interface{}

    // snmpInTooBigs. The type is interface{} with range: 0..4294967295.
    TooBigPacketReceived interface{}

    // snmp maximum packet size. The type is interface{} with range:
    // 0..4294967295.
    MaxPacketSize interface{}

    // snmpInNoSuchNames. The type is interface{} with range: 0..4294967295.
    NoSuchNamesReceived interface{}

    // snmpInBadValues. The type is interface{} with range: 0..4294967295.
    BadValuesReceived interface{}

    // snmpInReadOnlys. The type is interface{} with range: 0..4294967295.
    ReadOnlyReceived interface{}

    // snmpInGenErrs. The type is interface{} with range: 0..4294967295.
    TotalGeneralErrors interface{}

    // snmpInTotalReqVars. The type is interface{} with range: 0..4294967295.
    TotalRequestedVariables interface{}

    // snmpInTotalSetVars. The type is interface{} with range: 0..4294967295.
    TotalSetVariablesReceived interface{}

    // snmpInGetRequests. The type is interface{} with range: 0..4294967295.
    GetRequestsReceived interface{}

    // snmpInGetNexts. The type is interface{} with range: 0..4294967295.
    GetNextRequestsReceived interface{}

    // snmpInSetRequests. The type is interface{} with range: 0..4294967295.
    SetRequestsReceived interface{}

    // snmpInGetResponses. The type is interface{} with range: 0..4294967295.
    GetResponsesReceived interface{}

    // snmpInTraps. The type is interface{} with range: 0..4294967295.
    TrapsReceived interface{}

    // snmpOutPkts. The type is interface{} with range: 0..4294967295.
    TotalPacketsSent interface{}

    // snmpOutTooBigs. The type is interface{} with range: 0..4294967295.
    TooBigPacketsSent interface{}

    // snmpOutNoSuchNames. The type is interface{} with range: 0..4294967295.
    NoSuchNamesSent interface{}

    // snmpOutBadValues. The type is interface{} with range: 0..4294967295.
    BadValuesSent interface{}

    // snmpOutGenErrs. The type is interface{} with range: 0..4294967295.
    GeneralErrorsSent interface{}

    // snmpOutGetRequests. The type is interface{} with range: 0..4294967295.
    GetRequestsSent interface{}

    // snmpOutGetNexts. The type is interface{} with range: 0..4294967295.
    GetNextRequestSent interface{}

    // snmpOutSetRequests. The type is interface{} with range: 0..4294967295.
    SetRequestsSent interface{}

    // snmpOutGetResponses. The type is interface{} with range: 0..4294967295.
    GetResponsesSent interface{}

    // snmpOutTraps. The type is interface{} with range: 0..4294967295.
    TrapsSent interface{}
}

func (statistics *Snmp_Information_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "information"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", statistics.PacketsReceived})
    statistics.EntityData.Leafs.Append("bad-versions-received", types.YLeaf{"BadVersionsReceived", statistics.BadVersionsReceived})
    statistics.EntityData.Leafs.Append("bad-community-names-received", types.YLeaf{"BadCommunityNamesReceived", statistics.BadCommunityNamesReceived})
    statistics.EntityData.Leafs.Append("bad-community-uses-received", types.YLeaf{"BadCommunityUsesReceived", statistics.BadCommunityUsesReceived})
    statistics.EntityData.Leafs.Append("asn-parse-errors-received", types.YLeaf{"AsnParseErrorsReceived", statistics.AsnParseErrorsReceived})
    statistics.EntityData.Leafs.Append("silent-drop-count", types.YLeaf{"SilentDropCount", statistics.SilentDropCount})
    statistics.EntityData.Leafs.Append("proxy-drop-count", types.YLeaf{"ProxyDropCount", statistics.ProxyDropCount})
    statistics.EntityData.Leafs.Append("too-big-packet-received", types.YLeaf{"TooBigPacketReceived", statistics.TooBigPacketReceived})
    statistics.EntityData.Leafs.Append("max-packet-size", types.YLeaf{"MaxPacketSize", statistics.MaxPacketSize})
    statistics.EntityData.Leafs.Append("no-such-names-received", types.YLeaf{"NoSuchNamesReceived", statistics.NoSuchNamesReceived})
    statistics.EntityData.Leafs.Append("bad-values-received", types.YLeaf{"BadValuesReceived", statistics.BadValuesReceived})
    statistics.EntityData.Leafs.Append("read-only-received", types.YLeaf{"ReadOnlyReceived", statistics.ReadOnlyReceived})
    statistics.EntityData.Leafs.Append("total-general-errors", types.YLeaf{"TotalGeneralErrors", statistics.TotalGeneralErrors})
    statistics.EntityData.Leafs.Append("total-requested-variables", types.YLeaf{"TotalRequestedVariables", statistics.TotalRequestedVariables})
    statistics.EntityData.Leafs.Append("total-set-variables-received", types.YLeaf{"TotalSetVariablesReceived", statistics.TotalSetVariablesReceived})
    statistics.EntityData.Leafs.Append("get-requests-received", types.YLeaf{"GetRequestsReceived", statistics.GetRequestsReceived})
    statistics.EntityData.Leafs.Append("get-next-requests-received", types.YLeaf{"GetNextRequestsReceived", statistics.GetNextRequestsReceived})
    statistics.EntityData.Leafs.Append("set-requests-received", types.YLeaf{"SetRequestsReceived", statistics.SetRequestsReceived})
    statistics.EntityData.Leafs.Append("get-responses-received", types.YLeaf{"GetResponsesReceived", statistics.GetResponsesReceived})
    statistics.EntityData.Leafs.Append("traps-received", types.YLeaf{"TrapsReceived", statistics.TrapsReceived})
    statistics.EntityData.Leafs.Append("total-packets-sent", types.YLeaf{"TotalPacketsSent", statistics.TotalPacketsSent})
    statistics.EntityData.Leafs.Append("too-big-packets-sent", types.YLeaf{"TooBigPacketsSent", statistics.TooBigPacketsSent})
    statistics.EntityData.Leafs.Append("no-such-names-sent", types.YLeaf{"NoSuchNamesSent", statistics.NoSuchNamesSent})
    statistics.EntityData.Leafs.Append("bad-values-sent", types.YLeaf{"BadValuesSent", statistics.BadValuesSent})
    statistics.EntityData.Leafs.Append("general-errors-sent", types.YLeaf{"GeneralErrorsSent", statistics.GeneralErrorsSent})
    statistics.EntityData.Leafs.Append("get-requests-sent", types.YLeaf{"GetRequestsSent", statistics.GetRequestsSent})
    statistics.EntityData.Leafs.Append("get-next-request-sent", types.YLeaf{"GetNextRequestSent", statistics.GetNextRequestSent})
    statistics.EntityData.Leafs.Append("set-requests-sent", types.YLeaf{"SetRequestsSent", statistics.SetRequestsSent})
    statistics.EntityData.Leafs.Append("get-responses-sent", types.YLeaf{"GetResponsesSent", statistics.GetResponsesSent})
    statistics.EntityData.Leafs.Append("traps-sent", types.YLeaf{"TrapsSent", statistics.TrapsSent})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Snmp_Information_IncomingQueue
// Incoming queue details 
type Snmp_Information_IncomingQueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of NMS Queues Exist. The type is interface{} with range:
    // 0..4294967295.
    QueueCount interface{}

    // Each Entry Details. The type is slice of
    // Snmp_Information_IncomingQueue_InqEntry.
    InqEntry []*Snmp_Information_IncomingQueue_InqEntry
}

func (incomingQueue *Snmp_Information_IncomingQueue) GetEntityData() *types.CommonEntityData {
    incomingQueue.EntityData.YFilter = incomingQueue.YFilter
    incomingQueue.EntityData.YangName = "incoming-queue"
    incomingQueue.EntityData.BundleName = "cisco_ios_xr"
    incomingQueue.EntityData.ParentYangName = "information"
    incomingQueue.EntityData.SegmentPath = "incoming-queue"
    incomingQueue.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + incomingQueue.EntityData.SegmentPath
    incomingQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incomingQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incomingQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incomingQueue.EntityData.Children = types.NewOrderedMap()
    incomingQueue.EntityData.Children.Append("inq-entry", types.YChild{"InqEntry", nil})
    for i := range incomingQueue.InqEntry {
        types.SetYListKey(incomingQueue.InqEntry[i], i)
        incomingQueue.EntityData.Children.Append(types.GetSegmentPath(incomingQueue.InqEntry[i]), types.YChild{"InqEntry", incomingQueue.InqEntry[i]})
    }
    incomingQueue.EntityData.Leafs = types.NewOrderedMap()
    incomingQueue.EntityData.Leafs.Append("queue-count", types.YLeaf{"QueueCount", incomingQueue.QueueCount})

    incomingQueue.EntityData.YListKeys = []string {}

    return &(incomingQueue.EntityData)
}

// Snmp_Information_IncomingQueue_InqEntry
// Each Entry Details.
type Snmp_Information_IncomingQueue_InqEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Address of NMS Q. The type is string.
    AddressOfQueue interface{}

    // Request Count of Each Queue. The type is interface{} with range:
    // 0..4294967295.
    RequestCount interface{}

    // Processed request Count. The type is interface{} with range: 0..4294967295.
    ProcessedRequestCount interface{}

    // Last Access time of Each Queue. The type is string.
    LastAccessTime interface{}

    // Priority of Each Queue. The type is interface{} with range: 0..255.
    Priority interface{}
}

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetEntityData() *types.CommonEntityData {
    inqEntry.EntityData.YFilter = inqEntry.YFilter
    inqEntry.EntityData.YangName = "inq-entry"
    inqEntry.EntityData.BundleName = "cisco_ios_xr"
    inqEntry.EntityData.ParentYangName = "incoming-queue"
    inqEntry.EntityData.SegmentPath = "inq-entry" + types.AddNoKeyToken(inqEntry)
    inqEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/incoming-queue/" + inqEntry.EntityData.SegmentPath
    inqEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inqEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inqEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inqEntry.EntityData.Children = types.NewOrderedMap()
    inqEntry.EntityData.Leafs = types.NewOrderedMap()
    inqEntry.EntityData.Leafs.Append("address-of-queue", types.YLeaf{"AddressOfQueue", inqEntry.AddressOfQueue})
    inqEntry.EntityData.Leafs.Append("request-count", types.YLeaf{"RequestCount", inqEntry.RequestCount})
    inqEntry.EntityData.Leafs.Append("processed-request-count", types.YLeaf{"ProcessedRequestCount", inqEntry.ProcessedRequestCount})
    inqEntry.EntityData.Leafs.Append("last-access-time", types.YLeaf{"LastAccessTime", inqEntry.LastAccessTime})
    inqEntry.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", inqEntry.Priority})

    inqEntry.EntityData.YListKeys = []string {}

    return &(inqEntry.EntityData)
}

// Snmp_Information_ContextMapping
// Context name, features name, topology name,
// instance
type Snmp_Information_ContextMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context Mapping. The type is slice of
    // Snmp_Information_ContextMapping_ContexMapping.
    ContexMapping []*Snmp_Information_ContextMapping_ContexMapping
}

func (contextMapping *Snmp_Information_ContextMapping) GetEntityData() *types.CommonEntityData {
    contextMapping.EntityData.YFilter = contextMapping.YFilter
    contextMapping.EntityData.YangName = "context-mapping"
    contextMapping.EntityData.BundleName = "cisco_ios_xr"
    contextMapping.EntityData.ParentYangName = "information"
    contextMapping.EntityData.SegmentPath = "context-mapping"
    contextMapping.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + contextMapping.EntityData.SegmentPath
    contextMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextMapping.EntityData.Children = types.NewOrderedMap()
    contextMapping.EntityData.Children.Append("contex-mapping", types.YChild{"ContexMapping", nil})
    for i := range contextMapping.ContexMapping {
        types.SetYListKey(contextMapping.ContexMapping[i], i)
        contextMapping.EntityData.Children.Append(types.GetSegmentPath(contextMapping.ContexMapping[i]), types.YChild{"ContexMapping", contextMapping.ContexMapping[i]})
    }
    contextMapping.EntityData.Leafs = types.NewOrderedMap()

    contextMapping.EntityData.YListKeys = []string {}

    return &(contextMapping.EntityData)
}

// Snmp_Information_ContextMapping_ContexMapping
// Context Mapping
type Snmp_Information_ContextMapping_ContexMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Context name. The type is string.
    Context interface{}

    // Feature name. The type is string.
    FeatureName interface{}

    // Instance name. The type is string.
    Instance interface{}

    // Topology name. The type is string.
    Topology interface{}

    // Feature. The type is string.
    Feature interface{}
}

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetEntityData() *types.CommonEntityData {
    contexMapping.EntityData.YFilter = contexMapping.YFilter
    contexMapping.EntityData.YangName = "contex-mapping"
    contexMapping.EntityData.BundleName = "cisco_ios_xr"
    contexMapping.EntityData.ParentYangName = "context-mapping"
    contexMapping.EntityData.SegmentPath = "contex-mapping" + types.AddNoKeyToken(contexMapping)
    contexMapping.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/context-mapping/" + contexMapping.EntityData.SegmentPath
    contexMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contexMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contexMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contexMapping.EntityData.Children = types.NewOrderedMap()
    contexMapping.EntityData.Leafs = types.NewOrderedMap()
    contexMapping.EntityData.Leafs.Append("context", types.YLeaf{"Context", contexMapping.Context})
    contexMapping.EntityData.Leafs.Append("feature-name", types.YLeaf{"FeatureName", contexMapping.FeatureName})
    contexMapping.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", contexMapping.Instance})
    contexMapping.EntityData.Leafs.Append("topology", types.YLeaf{"Topology", contexMapping.Topology})
    contexMapping.EntityData.Leafs.Append("feature", types.YLeaf{"Feature", contexMapping.Feature})

    contexMapping.EntityData.YListKeys = []string {}

    return &(contexMapping.EntityData)
}

// Snmp_Information_TrapOids
// SNMP trap OID
type Snmp_Information_TrapOids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP trap . The type is slice of Snmp_Information_TrapOids_TrapOid.
    TrapOid []*Snmp_Information_TrapOids_TrapOid
}

func (trapOids *Snmp_Information_TrapOids) GetEntityData() *types.CommonEntityData {
    trapOids.EntityData.YFilter = trapOids.YFilter
    trapOids.EntityData.YangName = "trap-oids"
    trapOids.EntityData.BundleName = "cisco_ios_xr"
    trapOids.EntityData.ParentYangName = "information"
    trapOids.EntityData.SegmentPath = "trap-oids"
    trapOids.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + trapOids.EntityData.SegmentPath
    trapOids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapOids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapOids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapOids.EntityData.Children = types.NewOrderedMap()
    trapOids.EntityData.Children.Append("trap-oid", types.YChild{"TrapOid", nil})
    for i := range trapOids.TrapOid {
        trapOids.EntityData.Children.Append(types.GetSegmentPath(trapOids.TrapOid[i]), types.YChild{"TrapOid", trapOids.TrapOid[i]})
    }
    trapOids.EntityData.Leafs = types.NewOrderedMap()

    trapOids.EntityData.YListKeys = []string {}

    return &(trapOids.EntityData)
}

// Snmp_Information_TrapOids_TrapOid
// SNMP trap 
type Snmp_Information_TrapOids_TrapOid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Trap object ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TrapOid interface{}

    // Total number of OID's sent. The type is interface{} with range:
    // 0..4294967295.
    TrapOidCount interface{}

    // TRAP OID. The type is string.
    TrapOidXr interface{}
}

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetEntityData() *types.CommonEntityData {
    trapOid.EntityData.YFilter = trapOid.YFilter
    trapOid.EntityData.YangName = "trap-oid"
    trapOid.EntityData.BundleName = "cisco_ios_xr"
    trapOid.EntityData.ParentYangName = "trap-oids"
    trapOid.EntityData.SegmentPath = "trap-oid" + types.AddKeyToken(trapOid.TrapOid, "trap-oid")
    trapOid.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/trap-oids/" + trapOid.EntityData.SegmentPath
    trapOid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapOid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapOid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapOid.EntityData.Children = types.NewOrderedMap()
    trapOid.EntityData.Leafs = types.NewOrderedMap()
    trapOid.EntityData.Leafs.Append("trap-oid", types.YLeaf{"TrapOid", trapOid.TrapOid})
    trapOid.EntityData.Leafs.Append("trap-oid-count", types.YLeaf{"TrapOidCount", trapOid.TrapOidCount})
    trapOid.EntityData.Leafs.Append("trap-oid-xr", types.YLeaf{"TrapOidXr", trapOid.TrapOidXr})

    trapOid.EntityData.YListKeys = []string {"TrapOid"}

    return &(trapOid.EntityData)
}

// Snmp_Information_NmSpackets
// SNMP overload statistics 
type Snmp_Information_NmSpackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NMS packet drop count. The type is slice of
    // Snmp_Information_NmSpackets_NmSpacket.
    NmSpacket []*Snmp_Information_NmSpackets_NmSpacket
}

func (nmSpackets *Snmp_Information_NmSpackets) GetEntityData() *types.CommonEntityData {
    nmSpackets.EntityData.YFilter = nmSpackets.YFilter
    nmSpackets.EntityData.YangName = "nm-spackets"
    nmSpackets.EntityData.BundleName = "cisco_ios_xr"
    nmSpackets.EntityData.ParentYangName = "information"
    nmSpackets.EntityData.SegmentPath = "nm-spackets"
    nmSpackets.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + nmSpackets.EntityData.SegmentPath
    nmSpackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmSpackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmSpackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmSpackets.EntityData.Children = types.NewOrderedMap()
    nmSpackets.EntityData.Children.Append("nm-spacket", types.YChild{"NmSpacket", nil})
    for i := range nmSpackets.NmSpacket {
        nmSpackets.EntityData.Children.Append(types.GetSegmentPath(nmSpackets.NmSpacket[i]), types.YChild{"NmSpacket", nmSpackets.NmSpacket[i]})
    }
    nmSpackets.EntityData.Leafs = types.NewOrderedMap()

    nmSpackets.EntityData.YListKeys = []string {}

    return &(nmSpackets.EntityData)
}

// Snmp_Information_NmSpackets_NmSpacket
// NMS packet drop count
type Snmp_Information_NmSpackets_NmSpacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NMS packet drop count. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Packetcount interface{}

    // Number of packets which are currently enqueued within the NMS queues. The
    // type is interface{} with range: 0..4294967295.
    NumberOfNmsqPktsDropped interface{}

    // Number of packets dropped. The type is interface{} with range:
    // 0..4294967295.
    NumberOfPktsDropped interface{}

    // Time of overload contol begin. The type is string.
    OverloadStartTime interface{}

    // Time of overload contol End. The type is string.
    OverloadEndTime interface{}
}

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetEntityData() *types.CommonEntityData {
    nmSpacket.EntityData.YFilter = nmSpacket.YFilter
    nmSpacket.EntityData.YangName = "nm-spacket"
    nmSpacket.EntityData.BundleName = "cisco_ios_xr"
    nmSpacket.EntityData.ParentYangName = "nm-spackets"
    nmSpacket.EntityData.SegmentPath = "nm-spacket" + types.AddKeyToken(nmSpacket.Packetcount, "packetcount")
    nmSpacket.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/nm-spackets/" + nmSpacket.EntityData.SegmentPath
    nmSpacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmSpacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmSpacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmSpacket.EntityData.Children = types.NewOrderedMap()
    nmSpacket.EntityData.Leafs = types.NewOrderedMap()
    nmSpacket.EntityData.Leafs.Append("packetcount", types.YLeaf{"Packetcount", nmSpacket.Packetcount})
    nmSpacket.EntityData.Leafs.Append("number-of-nmsq-pkts-dropped", types.YLeaf{"NumberOfNmsqPktsDropped", nmSpacket.NumberOfNmsqPktsDropped})
    nmSpacket.EntityData.Leafs.Append("number-of-pkts-dropped", types.YLeaf{"NumberOfPktsDropped", nmSpacket.NumberOfPktsDropped})
    nmSpacket.EntityData.Leafs.Append("overload-start-time", types.YLeaf{"OverloadStartTime", nmSpacket.OverloadStartTime})
    nmSpacket.EntityData.Leafs.Append("overload-end-time", types.YLeaf{"OverloadEndTime", nmSpacket.OverloadEndTime})

    nmSpacket.EntityData.YListKeys = []string {"Packetcount"}

    return &(nmSpacket.EntityData)
}

// Snmp_Information_Mibs
// List of MIBS supported on the system
type Snmp_Information_Mibs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP MIB Name. The type is slice of Snmp_Information_Mibs_Mib.
    Mib []*Snmp_Information_Mibs_Mib
}

func (mibs *Snmp_Information_Mibs) GetEntityData() *types.CommonEntityData {
    mibs.EntityData.YFilter = mibs.YFilter
    mibs.EntityData.YangName = "mibs"
    mibs.EntityData.BundleName = "cisco_ios_xr"
    mibs.EntityData.ParentYangName = "information"
    mibs.EntityData.SegmentPath = "mibs"
    mibs.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + mibs.EntityData.SegmentPath
    mibs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mibs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mibs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mibs.EntityData.Children = types.NewOrderedMap()
    mibs.EntityData.Children.Append("mib", types.YChild{"Mib", nil})
    for i := range mibs.Mib {
        mibs.EntityData.Children.Append(types.GetSegmentPath(mibs.Mib[i]), types.YChild{"Mib", mibs.Mib[i]})
    }
    mibs.EntityData.Leafs = types.NewOrderedMap()

    mibs.EntityData.YListKeys = []string {}

    return &(mibs.EntityData)
}

// Snmp_Information_Mibs_Mib
// SNMP MIB Name
type Snmp_Information_Mibs_Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MIB Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // List of OIDs per MIB.
    Oids Snmp_Information_Mibs_Mib_Oids

    // MIB state and information.
    MibInformation Snmp_Information_Mibs_Mib_MibInformation
}

func (mib *Snmp_Information_Mibs_Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "mibs"
    mib.EntityData.SegmentPath = "mib" + types.AddKeyToken(mib.Name, "name")
    mib.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/mibs/" + mib.EntityData.SegmentPath
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = types.NewOrderedMap()
    mib.EntityData.Children.Append("oids", types.YChild{"Oids", &mib.Oids})
    mib.EntityData.Children.Append("mib-information", types.YChild{"MibInformation", &mib.MibInformation})
    mib.EntityData.Leafs = types.NewOrderedMap()
    mib.EntityData.Leafs.Append("name", types.YLeaf{"Name", mib.Name})

    mib.EntityData.YListKeys = []string {"Name"}

    return &(mib.EntityData)
}

// Snmp_Information_Mibs_Mib_Oids
// List of OIDs per MIB
type Snmp_Information_Mibs_Mib_Oids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object identifiers of a mib. The type is slice of
    // Snmp_Information_Mibs_Mib_Oids_Oid.
    Oid []*Snmp_Information_Mibs_Mib_Oids_Oid
}

func (oids *Snmp_Information_Mibs_Mib_Oids) GetEntityData() *types.CommonEntityData {
    oids.EntityData.YFilter = oids.YFilter
    oids.EntityData.YangName = "oids"
    oids.EntityData.BundleName = "cisco_ios_xr"
    oids.EntityData.ParentYangName = "mib"
    oids.EntityData.SegmentPath = "oids"
    oids.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/mibs/mib/" + oids.EntityData.SegmentPath
    oids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oids.EntityData.Children = types.NewOrderedMap()
    oids.EntityData.Children.Append("oid", types.YChild{"Oid", nil})
    for i := range oids.Oid {
        oids.EntityData.Children.Append(types.GetSegmentPath(oids.Oid[i]), types.YChild{"Oid", oids.Oid[i]})
    }
    oids.EntityData.Leafs = types.NewOrderedMap()

    oids.EntityData.YListKeys = []string {}

    return &(oids.EntityData)
}

// Snmp_Information_Mibs_Mib_Oids_Oid
// Object identifiers of a mib
type Snmp_Information_Mibs_Mib_Oids_Oid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object Identifier. The type is string.
    Oid interface{}

    // MIB OID Name. The type is string. This attribute is mandatory.
    OidName interface{}
}

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetEntityData() *types.CommonEntityData {
    oid.EntityData.YFilter = oid.YFilter
    oid.EntityData.YangName = "oid"
    oid.EntityData.BundleName = "cisco_ios_xr"
    oid.EntityData.ParentYangName = "oids"
    oid.EntityData.SegmentPath = "oid" + types.AddKeyToken(oid.Oid, "oid")
    oid.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/mibs/mib/oids/" + oid.EntityData.SegmentPath
    oid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oid.EntityData.Children = types.NewOrderedMap()
    oid.EntityData.Leafs = types.NewOrderedMap()
    oid.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", oid.Oid})
    oid.EntityData.Leafs.Append("oid-name", types.YLeaf{"OidName", oid.OidName})

    oid.EntityData.YListKeys = []string {"Oid"}

    return &(oid.EntityData)
}

// Snmp_Information_Mibs_Mib_MibInformation
// MIB state and information
type Snmp_Information_Mibs_Mib_MibInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the MIB module. The type is string.
    MibName interface{}

    // MIB DLL filename, non-DLL MIBs will have no value. The type is string.
    DllName interface{}

    // MIB config filename, non-DLL MIBs will have no value. The type is string.
    MibConfigFilename interface{}

    // TRUE if MIB DLL is currently loaded, will always be TRUE for non-DLL MIBs.
    // The type is bool.
    IsMibLoaded interface{}

    // DLL capabilities. The type is interface{} with range: 0..4294967295.
    DllCapabilities interface{}

    // List of trapstring configured. The type is string.
    TrapStrings interface{}

    // TRUE is mib is in phase 1 timeout. The type is bool.
    Timeout interface{}

    // Load time. The type is interface{} with range: 0..4294967295.
    LoadTime interface{}
}

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetEntityData() *types.CommonEntityData {
    mibInformation.EntityData.YFilter = mibInformation.YFilter
    mibInformation.EntityData.YangName = "mib-information"
    mibInformation.EntityData.BundleName = "cisco_ios_xr"
    mibInformation.EntityData.ParentYangName = "mib"
    mibInformation.EntityData.SegmentPath = "mib-information"
    mibInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/mibs/mib/" + mibInformation.EntityData.SegmentPath
    mibInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mibInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mibInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mibInformation.EntityData.Children = types.NewOrderedMap()
    mibInformation.EntityData.Leafs = types.NewOrderedMap()
    mibInformation.EntityData.Leafs.Append("mib-name", types.YLeaf{"MibName", mibInformation.MibName})
    mibInformation.EntityData.Leafs.Append("dll-name", types.YLeaf{"DllName", mibInformation.DllName})
    mibInformation.EntityData.Leafs.Append("mib-config-filename", types.YLeaf{"MibConfigFilename", mibInformation.MibConfigFilename})
    mibInformation.EntityData.Leafs.Append("is-mib-loaded", types.YLeaf{"IsMibLoaded", mibInformation.IsMibLoaded})
    mibInformation.EntityData.Leafs.Append("dll-capabilities", types.YLeaf{"DllCapabilities", mibInformation.DllCapabilities})
    mibInformation.EntityData.Leafs.Append("trap-strings", types.YLeaf{"TrapStrings", mibInformation.TrapStrings})
    mibInformation.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", mibInformation.Timeout})
    mibInformation.EntityData.Leafs.Append("load-time", types.YLeaf{"LoadTime", mibInformation.LoadTime})

    mibInformation.EntityData.YListKeys = []string {}

    return &(mibInformation.EntityData)
}

// Snmp_Information_SerialNumbers
// SNMP statistics pdu 
type Snmp_Information_SerialNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Serial number. The type is slice of
    // Snmp_Information_SerialNumbers_SerialNumber.
    SerialNumber []*Snmp_Information_SerialNumbers_SerialNumber
}

func (serialNumbers *Snmp_Information_SerialNumbers) GetEntityData() *types.CommonEntityData {
    serialNumbers.EntityData.YFilter = serialNumbers.YFilter
    serialNumbers.EntityData.YangName = "serial-numbers"
    serialNumbers.EntityData.BundleName = "cisco_ios_xr"
    serialNumbers.EntityData.ParentYangName = "information"
    serialNumbers.EntityData.SegmentPath = "serial-numbers"
    serialNumbers.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + serialNumbers.EntityData.SegmentPath
    serialNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serialNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serialNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serialNumbers.EntityData.Children = types.NewOrderedMap()
    serialNumbers.EntityData.Children.Append("serial-number", types.YChild{"SerialNumber", nil})
    for i := range serialNumbers.SerialNumber {
        types.SetYListKey(serialNumbers.SerialNumber[i], i)
        serialNumbers.EntityData.Children.Append(types.GetSegmentPath(serialNumbers.SerialNumber[i]), types.YChild{"SerialNumber", serialNumbers.SerialNumber[i]})
    }
    serialNumbers.EntityData.Leafs = types.NewOrderedMap()

    serialNumbers.EntityData.YListKeys = []string {}

    return &(serialNumbers.EntityData)
}

// Snmp_Information_SerialNumbers_SerialNumber
// Serial number
type Snmp_Information_SerialNumbers_SerialNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Serial number. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Number interface{}

    // Request ID. The type is interface{} with range: 0..4294967295.
    ReqId interface{}

    // Port. The type is interface{} with range: 0..65535.
    Port interface{}

    // NMS address Rx PDU. The type is string.
    Nms interface{}

    // SNMP request id per PDU. The type is interface{} with range: 0..4294967295.
    RequestId interface{}

    // NMS port number. The type is interface{} with range: 0..65535.
    PortXr interface{}

    // PDU type. The type is interface{} with range: 0..65535.
    PduType interface{}

    // Is reques dropped due to error. The type is interface{} with range:
    // 0..65535.
    ErrorStatus interface{}

    // Serial number per PDU processing. The type is interface{} with range:
    // 0..4294967295.
    SerialNum interface{}

    // Request inserted into input queue. The type is string.
    InputQ interface{}

    // De-queue the request from the input queue. The type is interface{} with
    // range: 0..4294967295.
    OutputQ interface{}

    // Enqueue the request into pending queue. The type is interface{} with range:
    // 0..4294967295.
    PendingQ interface{}

    // Response sent. The type is interface{} with range: 0..4294967295.
    ResponseOut interface{}
}

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetEntityData() *types.CommonEntityData {
    serialNumber.EntityData.YFilter = serialNumber.YFilter
    serialNumber.EntityData.YangName = "serial-number"
    serialNumber.EntityData.BundleName = "cisco_ios_xr"
    serialNumber.EntityData.ParentYangName = "serial-numbers"
    serialNumber.EntityData.SegmentPath = "serial-number" + types.AddNoKeyToken(serialNumber)
    serialNumber.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/serial-numbers/" + serialNumber.EntityData.SegmentPath
    serialNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serialNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serialNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serialNumber.EntityData.Children = types.NewOrderedMap()
    serialNumber.EntityData.Leafs = types.NewOrderedMap()
    serialNumber.EntityData.Leafs.Append("number", types.YLeaf{"Number", serialNumber.Number})
    serialNumber.EntityData.Leafs.Append("req-id", types.YLeaf{"ReqId", serialNumber.ReqId})
    serialNumber.EntityData.Leafs.Append("port", types.YLeaf{"Port", serialNumber.Port})
    serialNumber.EntityData.Leafs.Append("nms", types.YLeaf{"Nms", serialNumber.Nms})
    serialNumber.EntityData.Leafs.Append("request-id", types.YLeaf{"RequestId", serialNumber.RequestId})
    serialNumber.EntityData.Leafs.Append("port-xr", types.YLeaf{"PortXr", serialNumber.PortXr})
    serialNumber.EntityData.Leafs.Append("pdu-type", types.YLeaf{"PduType", serialNumber.PduType})
    serialNumber.EntityData.Leafs.Append("error-status", types.YLeaf{"ErrorStatus", serialNumber.ErrorStatus})
    serialNumber.EntityData.Leafs.Append("serial-num", types.YLeaf{"SerialNum", serialNumber.SerialNum})
    serialNumber.EntityData.Leafs.Append("input-q", types.YLeaf{"InputQ", serialNumber.InputQ})
    serialNumber.EntityData.Leafs.Append("output-q", types.YLeaf{"OutputQ", serialNumber.OutputQ})
    serialNumber.EntityData.Leafs.Append("pending-q", types.YLeaf{"PendingQ", serialNumber.PendingQ})
    serialNumber.EntityData.Leafs.Append("response-out", types.YLeaf{"ResponseOut", serialNumber.ResponseOut})

    serialNumber.EntityData.YListKeys = []string {}

    return &(serialNumber.EntityData)
}

// Snmp_Information_DropNmsAddresses
// NMS list for drop PDU
type Snmp_Information_DropNmsAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PDU drop info for NMS. The type is slice of
    // Snmp_Information_DropNmsAddresses_DropNmsAddress.
    DropNmsAddress []*Snmp_Information_DropNmsAddresses_DropNmsAddress
}

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetEntityData() *types.CommonEntityData {
    dropNmsAddresses.EntityData.YFilter = dropNmsAddresses.YFilter
    dropNmsAddresses.EntityData.YangName = "drop-nms-addresses"
    dropNmsAddresses.EntityData.BundleName = "cisco_ios_xr"
    dropNmsAddresses.EntityData.ParentYangName = "information"
    dropNmsAddresses.EntityData.SegmentPath = "drop-nms-addresses"
    dropNmsAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + dropNmsAddresses.EntityData.SegmentPath
    dropNmsAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropNmsAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropNmsAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropNmsAddresses.EntityData.Children = types.NewOrderedMap()
    dropNmsAddresses.EntityData.Children.Append("drop-nms-address", types.YChild{"DropNmsAddress", nil})
    for i := range dropNmsAddresses.DropNmsAddress {
        dropNmsAddresses.EntityData.Children.Append(types.GetSegmentPath(dropNmsAddresses.DropNmsAddress[i]), types.YChild{"DropNmsAddress", dropNmsAddresses.DropNmsAddress[i]})
    }
    dropNmsAddresses.EntityData.Leafs = types.NewOrderedMap()

    dropNmsAddresses.EntityData.YListKeys = []string {}

    return &(dropNmsAddresses.EntityData)
}

// Snmp_Information_DropNmsAddresses_DropNmsAddress
// PDU drop info for NMS
type Snmp_Information_DropNmsAddresses_DropNmsAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NMS address. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    NmsAddr interface{}

    // NMS address of server. The type is string.
    NmsAddress interface{}

    // Drop Count at Incoming Q. The type is interface{} with range:
    // 0..4294967295.
    IncomingQCount interface{}

    // Drop Count at Incoming Q after threshold limit. The type is interface{}
    // with range: 0..4294967295.
    ThresholdIncomingQCount interface{}

    // Drop Count with Encode errors. The type is interface{} with range:
    // 0..4294967295.
    EncodeCount interface{}

    // Duplicate request drop count. The type is interface{} with range:
    // 0..4294967295.
    DuplicateCount interface{}

    // Drop Count at snmp Stack. The type is interface{} with range:
    // 0..4294967295.
    StackCount interface{}

    // drop count with AIPC Buffer Full. The type is interface{} with range:
    // 0..4294967295.
    AipcCount interface{}

    // Drop Count with overload notification. The type is interface{} with range:
    // 0..4294967295.
    OverloadCount interface{}

    // Drop count with timeout. The type is interface{} with range: 0..4294967295.
    TimeoutCount interface{}

    // drop with Internal Errors. The type is interface{} with range:
    // 0..4294967295.
    InternalCount interface{}
}

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetEntityData() *types.CommonEntityData {
    dropNmsAddress.EntityData.YFilter = dropNmsAddress.YFilter
    dropNmsAddress.EntityData.YangName = "drop-nms-address"
    dropNmsAddress.EntityData.BundleName = "cisco_ios_xr"
    dropNmsAddress.EntityData.ParentYangName = "drop-nms-addresses"
    dropNmsAddress.EntityData.SegmentPath = "drop-nms-address" + types.AddKeyToken(dropNmsAddress.NmsAddr, "nms-addr")
    dropNmsAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/drop-nms-addresses/" + dropNmsAddress.EntityData.SegmentPath
    dropNmsAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropNmsAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropNmsAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropNmsAddress.EntityData.Children = types.NewOrderedMap()
    dropNmsAddress.EntityData.Leafs = types.NewOrderedMap()
    dropNmsAddress.EntityData.Leafs.Append("nms-addr", types.YLeaf{"NmsAddr", dropNmsAddress.NmsAddr})
    dropNmsAddress.EntityData.Leafs.Append("nms-address", types.YLeaf{"NmsAddress", dropNmsAddress.NmsAddress})
    dropNmsAddress.EntityData.Leafs.Append("incoming-q-count", types.YLeaf{"IncomingQCount", dropNmsAddress.IncomingQCount})
    dropNmsAddress.EntityData.Leafs.Append("threshold-incoming-q-count", types.YLeaf{"ThresholdIncomingQCount", dropNmsAddress.ThresholdIncomingQCount})
    dropNmsAddress.EntityData.Leafs.Append("encode-count", types.YLeaf{"EncodeCount", dropNmsAddress.EncodeCount})
    dropNmsAddress.EntityData.Leafs.Append("duplicate-count", types.YLeaf{"DuplicateCount", dropNmsAddress.DuplicateCount})
    dropNmsAddress.EntityData.Leafs.Append("stack-count", types.YLeaf{"StackCount", dropNmsAddress.StackCount})
    dropNmsAddress.EntityData.Leafs.Append("aipc-count", types.YLeaf{"AipcCount", dropNmsAddress.AipcCount})
    dropNmsAddress.EntityData.Leafs.Append("overload-count", types.YLeaf{"OverloadCount", dropNmsAddress.OverloadCount})
    dropNmsAddress.EntityData.Leafs.Append("timeout-count", types.YLeaf{"TimeoutCount", dropNmsAddress.TimeoutCount})
    dropNmsAddress.EntityData.Leafs.Append("internal-count", types.YLeaf{"InternalCount", dropNmsAddress.InternalCount})

    dropNmsAddress.EntityData.YListKeys = []string {"NmsAddr"}

    return &(dropNmsAddress.EntityData)
}

// Snmp_Information_Views
// SNMP view information
type Snmp_Information_Views struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP target view name. The type is slice of Snmp_Information_Views_View.
    View []*Snmp_Information_Views_View
}

func (views *Snmp_Information_Views) GetEntityData() *types.CommonEntityData {
    views.EntityData.YFilter = views.YFilter
    views.EntityData.YangName = "views"
    views.EntityData.BundleName = "cisco_ios_xr"
    views.EntityData.ParentYangName = "information"
    views.EntityData.SegmentPath = "views"
    views.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + views.EntityData.SegmentPath
    views.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    views.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    views.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    views.EntityData.Children = types.NewOrderedMap()
    views.EntityData.Children.Append("view", types.YChild{"View", nil})
    for i := range views.View {
        views.EntityData.Children.Append(types.GetSegmentPath(views.View[i]), types.YChild{"View", views.View[i]})
    }
    views.EntityData.Leafs = types.NewOrderedMap()

    views.EntityData.YListKeys = []string {}

    return &(views.EntityData)
}

// Snmp_Information_Views_View
// SNMP target view name
type Snmp_Information_Views_View struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. View name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // View name ,familytype, storagetype and status. The type is slice of
    // Snmp_Information_Views_View_ViewInformation.
    ViewInformation []*Snmp_Information_Views_View_ViewInformation
}

func (view *Snmp_Information_Views_View) GetEntityData() *types.CommonEntityData {
    view.EntityData.YFilter = view.YFilter
    view.EntityData.YangName = "view"
    view.EntityData.BundleName = "cisco_ios_xr"
    view.EntityData.ParentYangName = "views"
    view.EntityData.SegmentPath = "view" + types.AddKeyToken(view.Name, "name")
    view.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/views/" + view.EntityData.SegmentPath
    view.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    view.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    view.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    view.EntityData.Children = types.NewOrderedMap()
    view.EntityData.Children.Append("view-information", types.YChild{"ViewInformation", nil})
    for i := range view.ViewInformation {
        view.EntityData.Children.Append(types.GetSegmentPath(view.ViewInformation[i]), types.YChild{"ViewInformation", view.ViewInformation[i]})
    }
    view.EntityData.Leafs = types.NewOrderedMap()
    view.EntityData.Leafs.Append("name", types.YLeaf{"Name", view.Name})

    view.EntityData.YListKeys = []string {"Name"}

    return &(view.EntityData)
}

// Snmp_Information_Views_View_ViewInformation
// View name ,familytype, storagetype and status
type Snmp_Information_Views_View_ViewInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. SNMP view OID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ObjectId interface{}

    // Include or exclude. The type is string.
    SnmpViewFamilyType interface{}

    // Storage type. The type is string.
    SnmpViewFamilyStorageType interface{}

    // Status of this entry. The type is string.
    SnmpViewFamilyStatus interface{}
}

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetEntityData() *types.CommonEntityData {
    viewInformation.EntityData.YFilter = viewInformation.YFilter
    viewInformation.EntityData.YangName = "view-information"
    viewInformation.EntityData.BundleName = "cisco_ios_xr"
    viewInformation.EntityData.ParentYangName = "view"
    viewInformation.EntityData.SegmentPath = "view-information" + types.AddKeyToken(viewInformation.ObjectId, "object-id")
    viewInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/views/view/" + viewInformation.EntityData.SegmentPath
    viewInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    viewInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    viewInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    viewInformation.EntityData.Children = types.NewOrderedMap()
    viewInformation.EntityData.Leafs = types.NewOrderedMap()
    viewInformation.EntityData.Leafs.Append("object-id", types.YLeaf{"ObjectId", viewInformation.ObjectId})
    viewInformation.EntityData.Leafs.Append("snmp-view-family-type", types.YLeaf{"SnmpViewFamilyType", viewInformation.SnmpViewFamilyType})
    viewInformation.EntityData.Leafs.Append("snmp-view-family-storage-type", types.YLeaf{"SnmpViewFamilyStorageType", viewInformation.SnmpViewFamilyStorageType})
    viewInformation.EntityData.Leafs.Append("snmp-view-family-status", types.YLeaf{"SnmpViewFamilyStatus", viewInformation.SnmpViewFamilyStatus})

    viewInformation.EntityData.YListKeys = []string {"ObjectId"}

    return &(viewInformation.EntityData)
}

// Snmp_Information_SystemDescr
// System description
type Snmp_Information_SystemDescr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sysDescr  1.3.6.1.2.1.1.1. The type is string.
    SysDescr interface{}
}

func (systemDescr *Snmp_Information_SystemDescr) GetEntityData() *types.CommonEntityData {
    systemDescr.EntityData.YFilter = systemDescr.YFilter
    systemDescr.EntityData.YangName = "system-descr"
    systemDescr.EntityData.BundleName = "cisco_ios_xr"
    systemDescr.EntityData.ParentYangName = "information"
    systemDescr.EntityData.SegmentPath = "system-descr"
    systemDescr.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + systemDescr.EntityData.SegmentPath
    systemDescr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemDescr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemDescr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemDescr.EntityData.Children = types.NewOrderedMap()
    systemDescr.EntityData.Leafs = types.NewOrderedMap()
    systemDescr.EntityData.Leafs.Append("sys-descr", types.YLeaf{"SysDescr", systemDescr.SysDescr})

    systemDescr.EntityData.YListKeys = []string {}

    return &(systemDescr.EntityData)
}

// Snmp_Information_Tables
// List of table
type Snmp_Information_Tables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of vacmAccessTable.
    Groups Snmp_Information_Tables_Groups

    // List of User.
    UserEngineIds Snmp_Information_Tables_UserEngineIds
}

func (tables *Snmp_Information_Tables) GetEntityData() *types.CommonEntityData {
    tables.EntityData.YFilter = tables.YFilter
    tables.EntityData.YangName = "tables"
    tables.EntityData.BundleName = "cisco_ios_xr"
    tables.EntityData.ParentYangName = "information"
    tables.EntityData.SegmentPath = "tables"
    tables.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + tables.EntityData.SegmentPath
    tables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tables.EntityData.Children = types.NewOrderedMap()
    tables.EntityData.Children.Append("groups", types.YChild{"Groups", &tables.Groups})
    tables.EntityData.Children.Append("user-engine-ids", types.YChild{"UserEngineIds", &tables.UserEngineIds})
    tables.EntityData.Leafs = types.NewOrderedMap()

    tables.EntityData.YListKeys = []string {}

    return &(tables.EntityData)
}

// Snmp_Information_Tables_Groups
// List of vacmAccessTable
type Snmp_Information_Tables_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP group name. The type is slice of Snmp_Information_Tables_Groups_Group.
    Group []*Snmp_Information_Tables_Groups_Group
}

func (groups *Snmp_Information_Tables_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "tables"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/tables/" + groups.EntityData.SegmentPath
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Snmp_Information_Tables_Groups_Group
// SNMP group name
type Snmp_Information_Tables_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Group Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Group Model.
    GroupInformations Snmp_Information_Tables_Groups_Group_GroupInformations
}

func (group *Snmp_Information_Tables_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.Name, "name")
    group.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/tables/groups/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("group-informations", types.YChild{"GroupInformations", &group.GroupInformations})
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("name", types.YLeaf{"Name", group.Name})

    group.EntityData.YListKeys = []string {"Name"}

    return &(group.EntityData)
}

// Snmp_Information_Tables_Groups_Group_GroupInformations
// Group Model
type Snmp_Information_Tables_Groups_Group_GroupInformations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group name ,status  and information. The type is slice of
    // Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation.
    GroupInformation []*Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation
}

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetEntityData() *types.CommonEntityData {
    groupInformations.EntityData.YFilter = groupInformations.YFilter
    groupInformations.EntityData.YangName = "group-informations"
    groupInformations.EntityData.BundleName = "cisco_ios_xr"
    groupInformations.EntityData.ParentYangName = "group"
    groupInformations.EntityData.SegmentPath = "group-informations"
    groupInformations.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/tables/groups/group/" + groupInformations.EntityData.SegmentPath
    groupInformations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupInformations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupInformations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupInformations.EntityData.Children = types.NewOrderedMap()
    groupInformations.EntityData.Children.Append("group-information", types.YChild{"GroupInformation", nil})
    for i := range groupInformations.GroupInformation {
        types.SetYListKey(groupInformations.GroupInformation[i], i)
        groupInformations.EntityData.Children.Append(types.GetSegmentPath(groupInformations.GroupInformation[i]), types.YChild{"GroupInformation", groupInformations.GroupInformation[i]})
    }
    groupInformations.EntityData.Leafs = types.NewOrderedMap()

    groupInformations.EntityData.YListKeys = []string {}

    return &(groupInformations.EntityData)
}

// Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation
// Group name ,status  and information
type Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Model number. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Modelnumber interface{}

    // Level. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Level interface{}

    // Read view name. The type is string.
    VacmAccessReadViewName interface{}

    // Write view name. The type is string.
    VacmAccessWriteViewName interface{}

    // Notify view name. The type is string.
    VacmAccessNotifyViewName interface{}

    // Status of this view configuration. The type is interface{} with range:
    // 0..4294967295.
    VacmAccessStatus interface{}
}

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetEntityData() *types.CommonEntityData {
    groupInformation.EntityData.YFilter = groupInformation.YFilter
    groupInformation.EntityData.YangName = "group-information"
    groupInformation.EntityData.BundleName = "cisco_ios_xr"
    groupInformation.EntityData.ParentYangName = "group-informations"
    groupInformation.EntityData.SegmentPath = "group-information" + types.AddNoKeyToken(groupInformation)
    groupInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/tables/groups/group/group-informations/" + groupInformation.EntityData.SegmentPath
    groupInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupInformation.EntityData.Children = types.NewOrderedMap()
    groupInformation.EntityData.Leafs = types.NewOrderedMap()
    groupInformation.EntityData.Leafs.Append("modelnumber", types.YLeaf{"Modelnumber", groupInformation.Modelnumber})
    groupInformation.EntityData.Leafs.Append("level", types.YLeaf{"Level", groupInformation.Level})
    groupInformation.EntityData.Leafs.Append("vacm-access-read-view-name", types.YLeaf{"VacmAccessReadViewName", groupInformation.VacmAccessReadViewName})
    groupInformation.EntityData.Leafs.Append("vacm-access-write-view-name", types.YLeaf{"VacmAccessWriteViewName", groupInformation.VacmAccessWriteViewName})
    groupInformation.EntityData.Leafs.Append("vacm-access-notify-view-name", types.YLeaf{"VacmAccessNotifyViewName", groupInformation.VacmAccessNotifyViewName})
    groupInformation.EntityData.Leafs.Append("vacm-access-status", types.YLeaf{"VacmAccessStatus", groupInformation.VacmAccessStatus})

    groupInformation.EntityData.YListKeys = []string {}

    return &(groupInformation.EntityData)
}

// Snmp_Information_Tables_UserEngineIds
// List of User
type Snmp_Information_Tables_UserEngineIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP engineId. The type is slice of
    // Snmp_Information_Tables_UserEngineIds_UserEngineId.
    UserEngineId []*Snmp_Information_Tables_UserEngineIds_UserEngineId
}

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetEntityData() *types.CommonEntityData {
    userEngineIds.EntityData.YFilter = userEngineIds.YFilter
    userEngineIds.EntityData.YangName = "user-engine-ids"
    userEngineIds.EntityData.BundleName = "cisco_ios_xr"
    userEngineIds.EntityData.ParentYangName = "tables"
    userEngineIds.EntityData.SegmentPath = "user-engine-ids"
    userEngineIds.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/tables/" + userEngineIds.EntityData.SegmentPath
    userEngineIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userEngineIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userEngineIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userEngineIds.EntityData.Children = types.NewOrderedMap()
    userEngineIds.EntityData.Children.Append("user-engine-id", types.YChild{"UserEngineId", nil})
    for i := range userEngineIds.UserEngineId {
        userEngineIds.EntityData.Children.Append(types.GetSegmentPath(userEngineIds.UserEngineId[i]), types.YChild{"UserEngineId", userEngineIds.UserEngineId[i]})
    }
    userEngineIds.EntityData.Leafs = types.NewOrderedMap()

    userEngineIds.EntityData.YListKeys = []string {}

    return &(userEngineIds.EntityData)
}

// Snmp_Information_Tables_UserEngineIds_UserEngineId
// SNMP engineId
type Snmp_Information_Tables_UserEngineIds_UserEngineId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. SNMP Engine ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    EngineId interface{}

    // User name ,storage type ,status . The type is slice of
    // Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName.
    UserName []*Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName
}

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetEntityData() *types.CommonEntityData {
    userEngineId.EntityData.YFilter = userEngineId.YFilter
    userEngineId.EntityData.YangName = "user-engine-id"
    userEngineId.EntityData.BundleName = "cisco_ios_xr"
    userEngineId.EntityData.ParentYangName = "user-engine-ids"
    userEngineId.EntityData.SegmentPath = "user-engine-id" + types.AddKeyToken(userEngineId.EngineId, "engine-id")
    userEngineId.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/tables/user-engine-ids/" + userEngineId.EntityData.SegmentPath
    userEngineId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userEngineId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userEngineId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userEngineId.EntityData.Children = types.NewOrderedMap()
    userEngineId.EntityData.Children.Append("user-name", types.YChild{"UserName", nil})
    for i := range userEngineId.UserName {
        userEngineId.EntityData.Children.Append(types.GetSegmentPath(userEngineId.UserName[i]), types.YChild{"UserName", userEngineId.UserName[i]})
    }
    userEngineId.EntityData.Leafs = types.NewOrderedMap()
    userEngineId.EntityData.Leafs.Append("engine-id", types.YLeaf{"EngineId", userEngineId.EngineId})

    userEngineId.EntityData.YListKeys = []string {"EngineId"}

    return &(userEngineId.EntityData)
}

// Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName
// User name ,storage type ,status 
type Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. User name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    UserName interface{}

    // Storage type. The type is interface{} with range: 0..4294967295.
    UsmUserStorageType interface{}

    // Status of this user. The type is interface{} with range: 0..4294967295.
    UsmUserStatus interface{}
}

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetEntityData() *types.CommonEntityData {
    userName.EntityData.YFilter = userName.YFilter
    userName.EntityData.YangName = "user-name"
    userName.EntityData.BundleName = "cisco_ios_xr"
    userName.EntityData.ParentYangName = "user-engine-id"
    userName.EntityData.SegmentPath = "user-name" + types.AddKeyToken(userName.UserName, "user-name")
    userName.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/tables/user-engine-ids/user-engine-id/" + userName.EntityData.SegmentPath
    userName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userName.EntityData.Children = types.NewOrderedMap()
    userName.EntityData.Leafs = types.NewOrderedMap()
    userName.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", userName.UserName})
    userName.EntityData.Leafs.Append("usm-user-storage-type", types.YLeaf{"UsmUserStorageType", userName.UsmUserStorageType})
    userName.EntityData.Leafs.Append("usm-user-status", types.YLeaf{"UsmUserStatus", userName.UsmUserStatus})

    userName.EntityData.YListKeys = []string {"UserName"}

    return &(userName.EntityData)
}

// Snmp_Information_SystemOid
// System object ID
type Snmp_Information_SystemOid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sysObjID  1.3.6.1.2.1.1.2. The type is string.
    SysObjId interface{}
}

func (systemOid *Snmp_Information_SystemOid) GetEntityData() *types.CommonEntityData {
    systemOid.EntityData.YFilter = systemOid.YFilter
    systemOid.EntityData.YangName = "system-oid"
    systemOid.EntityData.BundleName = "cisco_ios_xr"
    systemOid.EntityData.ParentYangName = "information"
    systemOid.EntityData.SegmentPath = "system-oid"
    systemOid.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + systemOid.EntityData.SegmentPath
    systemOid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemOid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemOid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemOid.EntityData.Children = types.NewOrderedMap()
    systemOid.EntityData.Leafs = types.NewOrderedMap()
    systemOid.EntityData.Leafs.Append("sys-obj-id", types.YLeaf{"SysObjId", systemOid.SysObjId})

    systemOid.EntityData.YListKeys = []string {}

    return &(systemOid.EntityData)
}

// Snmp_Information_TrapQueue
// SNMP trap queue statistics
type Snmp_Information_TrapQueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // trap min. The type is interface{} with range: 0..4294967295.
    TrapMin interface{}

    // trap avg. The type is interface{} with range: 0..4294967295.
    TrapAvg interface{}

    // trap max. The type is interface{} with range: 0..4294967295.
    TrapMax interface{}

    // trap q. The type is slice of Snmp_Information_TrapQueue_TrapQ.
    TrapQ []*Snmp_Information_TrapQueue_TrapQ
}

func (trapQueue *Snmp_Information_TrapQueue) GetEntityData() *types.CommonEntityData {
    trapQueue.EntityData.YFilter = trapQueue.YFilter
    trapQueue.EntityData.YangName = "trap-queue"
    trapQueue.EntityData.BundleName = "cisco_ios_xr"
    trapQueue.EntityData.ParentYangName = "information"
    trapQueue.EntityData.SegmentPath = "trap-queue"
    trapQueue.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/" + trapQueue.EntityData.SegmentPath
    trapQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapQueue.EntityData.Children = types.NewOrderedMap()
    trapQueue.EntityData.Children.Append("trap-q", types.YChild{"TrapQ", nil})
    for i := range trapQueue.TrapQ {
        types.SetYListKey(trapQueue.TrapQ[i], i)
        trapQueue.EntityData.Children.Append(types.GetSegmentPath(trapQueue.TrapQ[i]), types.YChild{"TrapQ", trapQueue.TrapQ[i]})
    }
    trapQueue.EntityData.Leafs = types.NewOrderedMap()
    trapQueue.EntityData.Leafs.Append("trap-min", types.YLeaf{"TrapMin", trapQueue.TrapMin})
    trapQueue.EntityData.Leafs.Append("trap-avg", types.YLeaf{"TrapAvg", trapQueue.TrapAvg})
    trapQueue.EntityData.Leafs.Append("trap-max", types.YLeaf{"TrapMax", trapQueue.TrapMax})

    trapQueue.EntityData.YListKeys = []string {}

    return &(trapQueue.EntityData)
}

// Snmp_Information_TrapQueue_TrapQ
// trap q
type Snmp_Information_TrapQueue_TrapQ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // min. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // avg. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // max. The type is interface{} with range: 0..4294967295.
    Max interface{}
}

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetEntityData() *types.CommonEntityData {
    trapQ.EntityData.YFilter = trapQ.YFilter
    trapQ.EntityData.YangName = "trap-q"
    trapQ.EntityData.BundleName = "cisco_ios_xr"
    trapQ.EntityData.ParentYangName = "trap-queue"
    trapQ.EntityData.SegmentPath = "trap-q" + types.AddNoKeyToken(trapQ)
    trapQ.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/information/trap-queue/" + trapQ.EntityData.SegmentPath
    trapQ.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapQ.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapQ.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapQ.EntityData.Children = types.NewOrderedMap()
    trapQ.EntityData.Leafs = types.NewOrderedMap()
    trapQ.EntityData.Leafs.Append("min", types.YLeaf{"Min", trapQ.Min})
    trapQ.EntityData.Leafs.Append("avg", types.YLeaf{"Avg", trapQ.Avg})
    trapQ.EntityData.Leafs.Append("max", types.YLeaf{"Max", trapQ.Max})

    trapQ.EntityData.YListKeys = []string {}

    return &(trapQ.EntityData)
}

// Snmp_Interfaces
// List of interfaces
type Snmp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is slice of Snmp_Interfaces_Interface.
    Interface []*Snmp_Interfaces_Interface
}

func (interfaces *Snmp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "snmp"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Snmp_Interfaces_Interface
// Interface Name
type Snmp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Name interface{}

    // Interface Index as used by MIB tables. The type is interface{} with range:
    // 0..4294967295. This attribute is mandatory.
    InterfaceIndex interface{}
}

func (self *Snmp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Name, "name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})
    self.EntityData.Leafs.Append("interface-index", types.YLeaf{"InterfaceIndex", self.InterfaceIndex})

    self.EntityData.YListKeys = []string {"Name"}

    return &(self.EntityData)
}

// Snmp_Correlator
// Trap Correlator operational data
type Snmp_Correlator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table that contains the database of correlation rule details.
    RuleDetails Snmp_Correlator_RuleDetails

    // Describes buffer utilization and parameters configured.
    BufferStatus Snmp_Correlator_BufferStatus

    // Table that contains the ruleset detail info.
    RuleSetDetails Snmp_Correlator_RuleSetDetails

    // Correlated traps Table.
    Traps Snmp_Correlator_Traps
}

func (correlator *Snmp_Correlator) GetEntityData() *types.CommonEntityData {
    correlator.EntityData.YFilter = correlator.YFilter
    correlator.EntityData.YangName = "correlator"
    correlator.EntityData.BundleName = "cisco_ios_xr"
    correlator.EntityData.ParentYangName = "snmp"
    correlator.EntityData.SegmentPath = "correlator"
    correlator.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + correlator.EntityData.SegmentPath
    correlator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    correlator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    correlator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    correlator.EntityData.Children = types.NewOrderedMap()
    correlator.EntityData.Children.Append("rule-details", types.YChild{"RuleDetails", &correlator.RuleDetails})
    correlator.EntityData.Children.Append("buffer-status", types.YChild{"BufferStatus", &correlator.BufferStatus})
    correlator.EntityData.Children.Append("rule-set-details", types.YChild{"RuleSetDetails", &correlator.RuleSetDetails})
    correlator.EntityData.Children.Append("traps", types.YChild{"Traps", &correlator.Traps})
    correlator.EntityData.Leafs = types.NewOrderedMap()

    correlator.EntityData.YListKeys = []string {}

    return &(correlator.EntityData)
}

// Snmp_Correlator_RuleDetails
// Table that contains the database of correlation
// rule details
type Snmp_Correlator_RuleDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Details of one of the correlation rules. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail.
    RuleDetail []*Snmp_Correlator_RuleDetails_RuleDetail
}

func (ruleDetails *Snmp_Correlator_RuleDetails) GetEntityData() *types.CommonEntityData {
    ruleDetails.EntityData.YFilter = ruleDetails.YFilter
    ruleDetails.EntityData.YangName = "rule-details"
    ruleDetails.EntityData.BundleName = "cisco_ios_xr"
    ruleDetails.EntityData.ParentYangName = "correlator"
    ruleDetails.EntityData.SegmentPath = "rule-details"
    ruleDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/" + ruleDetails.EntityData.SegmentPath
    ruleDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleDetails.EntityData.Children = types.NewOrderedMap()
    ruleDetails.EntityData.Children.Append("rule-detail", types.YChild{"RuleDetail", nil})
    for i := range ruleDetails.RuleDetail {
        ruleDetails.EntityData.Children.Append(types.GetSegmentPath(ruleDetails.RuleDetail[i]), types.YChild{"RuleDetail", ruleDetails.RuleDetail[i]})
    }
    ruleDetails.EntityData.Leafs = types.NewOrderedMap()

    ruleDetails.EntityData.YListKeys = []string {}

    return &(ruleDetails.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail
// Details of one of the correlation rules
type Snmp_Correlator_RuleDetails_RuleDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Correlation Rule Name. The type is string with
    // length: 1..32.
    RuleName interface{}

    // Time window (in ms) for which root/all messages are kept in correlater
    // before sending them to hosts. The type is interface{} with range:
    // 0..4294967295.
    Timeout interface{}

    // Rule summary, name, etc.
    RuleSummary Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary

    // OID/VarBinds defining the rootcause match conditions.
    RootCause Snmp_Correlator_RuleDetails_RuleDetail_RootCause

    // OIDs/VarBinds defining the nonrootcause match conditions. The type is slice
    // of Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus.
    NonRootcaus []*Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus

    // Hosts (IP/port) to which the rule is applied. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost.
    ApplyHost []*Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost
}

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetEntityData() *types.CommonEntityData {
    ruleDetail.EntityData.YFilter = ruleDetail.YFilter
    ruleDetail.EntityData.YangName = "rule-detail"
    ruleDetail.EntityData.BundleName = "cisco_ios_xr"
    ruleDetail.EntityData.ParentYangName = "rule-details"
    ruleDetail.EntityData.SegmentPath = "rule-detail" + types.AddKeyToken(ruleDetail.RuleName, "rule-name")
    ruleDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-details/" + ruleDetail.EntityData.SegmentPath
    ruleDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleDetail.EntityData.Children = types.NewOrderedMap()
    ruleDetail.EntityData.Children.Append("rule-summary", types.YChild{"RuleSummary", &ruleDetail.RuleSummary})
    ruleDetail.EntityData.Children.Append("root-cause", types.YChild{"RootCause", &ruleDetail.RootCause})
    ruleDetail.EntityData.Children.Append("non-rootcaus", types.YChild{"NonRootcaus", nil})
    for i := range ruleDetail.NonRootcaus {
        types.SetYListKey(ruleDetail.NonRootcaus[i], i)
        ruleDetail.EntityData.Children.Append(types.GetSegmentPath(ruleDetail.NonRootcaus[i]), types.YChild{"NonRootcaus", ruleDetail.NonRootcaus[i]})
    }
    ruleDetail.EntityData.Children.Append("apply-host", types.YChild{"ApplyHost", nil})
    for i := range ruleDetail.ApplyHost {
        types.SetYListKey(ruleDetail.ApplyHost[i], i)
        ruleDetail.EntityData.Children.Append(types.GetSegmentPath(ruleDetail.ApplyHost[i]), types.YChild{"ApplyHost", ruleDetail.ApplyHost[i]})
    }
    ruleDetail.EntityData.Leafs = types.NewOrderedMap()
    ruleDetail.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", ruleDetail.RuleName})
    ruleDetail.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", ruleDetail.Timeout})

    ruleDetail.EntityData.YListKeys = []string {"RuleName"}

    return &(ruleDetail.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary
// Rule summary, name, etc
type Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Correlation Rule Name. The type is string.
    RuleName interface{}

    // Applied state of the rule It could be not applied, applied or applied to
    // all. The type is SnmpCorrRuleState.
    RuleState interface{}

    // Number of buffered traps correlated to this rule. The type is interface{}
    // with range: 0..4294967295.
    BufferedTrapsCount interface{}
}

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetEntityData() *types.CommonEntityData {
    ruleSummary.EntityData.YFilter = ruleSummary.YFilter
    ruleSummary.EntityData.YangName = "rule-summary"
    ruleSummary.EntityData.BundleName = "cisco_ios_xr"
    ruleSummary.EntityData.ParentYangName = "rule-detail"
    ruleSummary.EntityData.SegmentPath = "rule-summary"
    ruleSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-details/rule-detail/" + ruleSummary.EntityData.SegmentPath
    ruleSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSummary.EntityData.Children = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", ruleSummary.RuleName})
    ruleSummary.EntityData.Leafs.Append("rule-state", types.YLeaf{"RuleState", ruleSummary.RuleState})
    ruleSummary.EntityData.Leafs.Append("buffered-traps-count", types.YLeaf{"BufferedTrapsCount", ruleSummary.BufferedTrapsCount})

    ruleSummary.EntityData.YListKeys = []string {}

    return &(ruleSummary.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_RootCause
// OID/VarBinds defining the rootcause match
// conditions.
type Snmp_Correlator_RuleDetails_RuleDetail_RootCause struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OID of the trap. The type is string.
    Oid interface{}

    // VarBinds of the trap. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind.
    VarBind []*Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind
}

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetEntityData() *types.CommonEntityData {
    rootCause.EntityData.YFilter = rootCause.YFilter
    rootCause.EntityData.YangName = "root-cause"
    rootCause.EntityData.BundleName = "cisco_ios_xr"
    rootCause.EntityData.ParentYangName = "rule-detail"
    rootCause.EntityData.SegmentPath = "root-cause"
    rootCause.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-details/rule-detail/" + rootCause.EntityData.SegmentPath
    rootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rootCause.EntityData.Children = types.NewOrderedMap()
    rootCause.EntityData.Children.Append("var-bind", types.YChild{"VarBind", nil})
    for i := range rootCause.VarBind {
        types.SetYListKey(rootCause.VarBind[i], i)
        rootCause.EntityData.Children.Append(types.GetSegmentPath(rootCause.VarBind[i]), types.YChild{"VarBind", rootCause.VarBind[i]})
    }
    rootCause.EntityData.Leafs = types.NewOrderedMap()
    rootCause.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", rootCause.Oid})

    rootCause.EntityData.YListKeys = []string {}

    return &(rootCause.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind
// VarBinds of the trap
type Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // OID of the varbind. The type is string.
    Oid interface{}

    // Varbind match type. The type is SnmpCorrVbindMatch.
    MatchType interface{}

    // Regular expression to match. The type is string.
    RegExp interface{}
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetEntityData() *types.CommonEntityData {
    varBind.EntityData.YFilter = varBind.YFilter
    varBind.EntityData.YangName = "var-bind"
    varBind.EntityData.BundleName = "cisco_ios_xr"
    varBind.EntityData.ParentYangName = "root-cause"
    varBind.EntityData.SegmentPath = "var-bind" + types.AddNoKeyToken(varBind)
    varBind.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-details/rule-detail/root-cause/" + varBind.EntityData.SegmentPath
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = types.NewOrderedMap()
    varBind.EntityData.Leafs = types.NewOrderedMap()
    varBind.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", varBind.Oid})
    varBind.EntityData.Leafs.Append("match-type", types.YLeaf{"MatchType", varBind.MatchType})
    varBind.EntityData.Leafs.Append("reg-exp", types.YLeaf{"RegExp", varBind.RegExp})

    varBind.EntityData.YListKeys = []string {}

    return &(varBind.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus
// OIDs/VarBinds defining the nonrootcause match
// conditions.
type Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // OID of the trap. The type is string.
    Oid interface{}

    // VarBinds of the trap. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind.
    VarBind []*Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind
}

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetEntityData() *types.CommonEntityData {
    nonRootcaus.EntityData.YFilter = nonRootcaus.YFilter
    nonRootcaus.EntityData.YangName = "non-rootcaus"
    nonRootcaus.EntityData.BundleName = "cisco_ios_xr"
    nonRootcaus.EntityData.ParentYangName = "rule-detail"
    nonRootcaus.EntityData.SegmentPath = "non-rootcaus" + types.AddNoKeyToken(nonRootcaus)
    nonRootcaus.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-details/rule-detail/" + nonRootcaus.EntityData.SegmentPath
    nonRootcaus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootcaus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootcaus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootcaus.EntityData.Children = types.NewOrderedMap()
    nonRootcaus.EntityData.Children.Append("var-bind", types.YChild{"VarBind", nil})
    for i := range nonRootcaus.VarBind {
        types.SetYListKey(nonRootcaus.VarBind[i], i)
        nonRootcaus.EntityData.Children.Append(types.GetSegmentPath(nonRootcaus.VarBind[i]), types.YChild{"VarBind", nonRootcaus.VarBind[i]})
    }
    nonRootcaus.EntityData.Leafs = types.NewOrderedMap()
    nonRootcaus.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", nonRootcaus.Oid})

    nonRootcaus.EntityData.YListKeys = []string {}

    return &(nonRootcaus.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind
// VarBinds of the trap
type Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // OID of the varbind. The type is string.
    Oid interface{}

    // Varbind match type. The type is SnmpCorrVbindMatch.
    MatchType interface{}

    // Regular expression to match. The type is string.
    RegExp interface{}
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetEntityData() *types.CommonEntityData {
    varBind.EntityData.YFilter = varBind.YFilter
    varBind.EntityData.YangName = "var-bind"
    varBind.EntityData.BundleName = "cisco_ios_xr"
    varBind.EntityData.ParentYangName = "non-rootcaus"
    varBind.EntityData.SegmentPath = "var-bind" + types.AddNoKeyToken(varBind)
    varBind.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-details/rule-detail/non-rootcaus/" + varBind.EntityData.SegmentPath
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = types.NewOrderedMap()
    varBind.EntityData.Leafs = types.NewOrderedMap()
    varBind.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", varBind.Oid})
    varBind.EntityData.Leafs.Append("match-type", types.YLeaf{"MatchType", varBind.MatchType})
    varBind.EntityData.Leafs.Append("reg-exp", types.YLeaf{"RegExp", varBind.RegExp})

    varBind.EntityData.YListKeys = []string {}

    return &(varBind.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost
// Hosts (IP/port) to which the rule is applied
type Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IP address of the host. The type is string.
    IpAddress interface{}

    // Port of the host. The type is interface{} with range: 0..65535.
    Port interface{}
}

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetEntityData() *types.CommonEntityData {
    applyHost.EntityData.YFilter = applyHost.YFilter
    applyHost.EntityData.YangName = "apply-host"
    applyHost.EntityData.BundleName = "cisco_ios_xr"
    applyHost.EntityData.ParentYangName = "rule-detail"
    applyHost.EntityData.SegmentPath = "apply-host" + types.AddNoKeyToken(applyHost)
    applyHost.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-details/rule-detail/" + applyHost.EntityData.SegmentPath
    applyHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    applyHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    applyHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    applyHost.EntityData.Children = types.NewOrderedMap()
    applyHost.EntityData.Leafs = types.NewOrderedMap()
    applyHost.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", applyHost.IpAddress})
    applyHost.EntityData.Leafs.Append("port", types.YLeaf{"Port", applyHost.Port})

    applyHost.EntityData.YListKeys = []string {}

    return &(applyHost.EntityData)
}

// Snmp_Correlator_BufferStatus
// Describes buffer utilization and parameters
// configured
type Snmp_Correlator_BufferStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current buffer usage. The type is interface{} with range: 0..4294967295.
    CurrentSize interface{}

    // Configured buffer size. The type is interface{} with range: 0..4294967295.
    ConfiguredSize interface{}
}

func (bufferStatus *Snmp_Correlator_BufferStatus) GetEntityData() *types.CommonEntityData {
    bufferStatus.EntityData.YFilter = bufferStatus.YFilter
    bufferStatus.EntityData.YangName = "buffer-status"
    bufferStatus.EntityData.BundleName = "cisco_ios_xr"
    bufferStatus.EntityData.ParentYangName = "correlator"
    bufferStatus.EntityData.SegmentPath = "buffer-status"
    bufferStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/" + bufferStatus.EntityData.SegmentPath
    bufferStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bufferStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bufferStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bufferStatus.EntityData.Children = types.NewOrderedMap()
    bufferStatus.EntityData.Leafs = types.NewOrderedMap()
    bufferStatus.EntityData.Leafs.Append("current-size", types.YLeaf{"CurrentSize", bufferStatus.CurrentSize})
    bufferStatus.EntityData.Leafs.Append("configured-size", types.YLeaf{"ConfiguredSize", bufferStatus.ConfiguredSize})

    bufferStatus.EntityData.YListKeys = []string {}

    return &(bufferStatus.EntityData)
}

// Snmp_Correlator_RuleSetDetails
// Table that contains the ruleset detail info
type Snmp_Correlator_RuleSetDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail of one of the correlation rulesets. The type is slice of
    // Snmp_Correlator_RuleSetDetails_RuleSetDetail.
    RuleSetDetail []*Snmp_Correlator_RuleSetDetails_RuleSetDetail
}

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetEntityData() *types.CommonEntityData {
    ruleSetDetails.EntityData.YFilter = ruleSetDetails.YFilter
    ruleSetDetails.EntityData.YangName = "rule-set-details"
    ruleSetDetails.EntityData.BundleName = "cisco_ios_xr"
    ruleSetDetails.EntityData.ParentYangName = "correlator"
    ruleSetDetails.EntityData.SegmentPath = "rule-set-details"
    ruleSetDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/" + ruleSetDetails.EntityData.SegmentPath
    ruleSetDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSetDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSetDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSetDetails.EntityData.Children = types.NewOrderedMap()
    ruleSetDetails.EntityData.Children.Append("rule-set-detail", types.YChild{"RuleSetDetail", nil})
    for i := range ruleSetDetails.RuleSetDetail {
        ruleSetDetails.EntityData.Children.Append(types.GetSegmentPath(ruleSetDetails.RuleSetDetail[i]), types.YChild{"RuleSetDetail", ruleSetDetails.RuleSetDetail[i]})
    }
    ruleSetDetails.EntityData.Leafs = types.NewOrderedMap()

    ruleSetDetails.EntityData.YListKeys = []string {}

    return &(ruleSetDetails.EntityData)
}

// Snmp_Correlator_RuleSetDetails_RuleSetDetail
// Detail of one of the correlation rulesets
type Snmp_Correlator_RuleSetDetails_RuleSetDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Ruleset Name. The type is string with length:
    // 1..32.
    RuleSetName interface{}

    // Ruleset Name. The type is string.
    RuleSetNameXr interface{}

    // Rules contained in a ruleset. The type is slice of
    // Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules.
    Rules []*Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules
}

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetEntityData() *types.CommonEntityData {
    ruleSetDetail.EntityData.YFilter = ruleSetDetail.YFilter
    ruleSetDetail.EntityData.YangName = "rule-set-detail"
    ruleSetDetail.EntityData.BundleName = "cisco_ios_xr"
    ruleSetDetail.EntityData.ParentYangName = "rule-set-details"
    ruleSetDetail.EntityData.SegmentPath = "rule-set-detail" + types.AddKeyToken(ruleSetDetail.RuleSetName, "rule-set-name")
    ruleSetDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-set-details/" + ruleSetDetail.EntityData.SegmentPath
    ruleSetDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSetDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSetDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSetDetail.EntityData.Children = types.NewOrderedMap()
    ruleSetDetail.EntityData.Children.Append("rules", types.YChild{"Rules", nil})
    for i := range ruleSetDetail.Rules {
        types.SetYListKey(ruleSetDetail.Rules[i], i)
        ruleSetDetail.EntityData.Children.Append(types.GetSegmentPath(ruleSetDetail.Rules[i]), types.YChild{"Rules", ruleSetDetail.Rules[i]})
    }
    ruleSetDetail.EntityData.Leafs = types.NewOrderedMap()
    ruleSetDetail.EntityData.Leafs.Append("rule-set-name", types.YLeaf{"RuleSetName", ruleSetDetail.RuleSetName})
    ruleSetDetail.EntityData.Leafs.Append("rule-set-name-xr", types.YLeaf{"RuleSetNameXr", ruleSetDetail.RuleSetNameXr})

    ruleSetDetail.EntityData.YListKeys = []string {"RuleSetName"}

    return &(ruleSetDetail.EntityData)
}

// Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules
// Rules contained in a ruleset
type Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Correlation Rule Name. The type is string.
    RuleName interface{}

    // Applied state of the rule It could be not applied, applied or applied to
    // all. The type is SnmpCorrRuleState.
    RuleState interface{}

    // Number of buffered traps correlated to this rule. The type is interface{}
    // with range: 0..4294967295.
    BufferedTrapsCount interface{}
}

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "rule-set-detail"
    rules.EntityData.SegmentPath = "rules" + types.AddNoKeyToken(rules)
    rules.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/rule-set-details/rule-set-detail/" + rules.EntityData.SegmentPath
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = types.NewOrderedMap()
    rules.EntityData.Leafs = types.NewOrderedMap()
    rules.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", rules.RuleName})
    rules.EntityData.Leafs.Append("rule-state", types.YLeaf{"RuleState", rules.RuleState})
    rules.EntityData.Leafs.Append("buffered-traps-count", types.YLeaf{"BufferedTrapsCount", rules.BufferedTrapsCount})

    rules.EntityData.YListKeys = []string {}

    return &(rules.EntityData)
}

// Snmp_Correlator_Traps
// Correlated traps Table
type Snmp_Correlator_Traps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One of the correlated traps. The type is slice of
    // Snmp_Correlator_Traps_Trap.
    Trap []*Snmp_Correlator_Traps_Trap
}

func (traps *Snmp_Correlator_Traps) GetEntityData() *types.CommonEntityData {
    traps.EntityData.YFilter = traps.YFilter
    traps.EntityData.YangName = "traps"
    traps.EntityData.BundleName = "cisco_ios_xr"
    traps.EntityData.ParentYangName = "correlator"
    traps.EntityData.SegmentPath = "traps"
    traps.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/" + traps.EntityData.SegmentPath
    traps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traps.EntityData.Children = types.NewOrderedMap()
    traps.EntityData.Children.Append("trap", types.YChild{"Trap", nil})
    for i := range traps.Trap {
        traps.EntityData.Children.Append(types.GetSegmentPath(traps.Trap[i]), types.YChild{"Trap", traps.Trap[i]})
    }
    traps.EntityData.Leafs = types.NewOrderedMap()

    traps.EntityData.YListKeys = []string {}

    return &(traps.EntityData)
}

// Snmp_Correlator_Traps_Trap
// One of the correlated traps
type Snmp_Correlator_Traps_Trap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Entry ID. The type is interface{} with range:
    // 0..4294967295.
    EntryId interface{}

    // Correlation ID. The type is interface{} with range: 0..4294967295.
    CorrelationId interface{}

    // True if this is the rootcause. The type is bool.
    IsRootCause interface{}

    // Correlation rule name. The type is string.
    RuleName interface{}

    // Correlated trap information.
    TrapInfo Snmp_Correlator_Traps_Trap_TrapInfo
}

func (trap *Snmp_Correlator_Traps_Trap) GetEntityData() *types.CommonEntityData {
    trap.EntityData.YFilter = trap.YFilter
    trap.EntityData.YangName = "trap"
    trap.EntityData.BundleName = "cisco_ios_xr"
    trap.EntityData.ParentYangName = "traps"
    trap.EntityData.SegmentPath = "trap" + types.AddKeyToken(trap.EntryId, "entry-id")
    trap.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/traps/" + trap.EntityData.SegmentPath
    trap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trap.EntityData.Children = types.NewOrderedMap()
    trap.EntityData.Children.Append("trap-info", types.YChild{"TrapInfo", &trap.TrapInfo})
    trap.EntityData.Leafs = types.NewOrderedMap()
    trap.EntityData.Leafs.Append("entry-id", types.YLeaf{"EntryId", trap.EntryId})
    trap.EntityData.Leafs.Append("correlation-id", types.YLeaf{"CorrelationId", trap.CorrelationId})
    trap.EntityData.Leafs.Append("is-root-cause", types.YLeaf{"IsRootCause", trap.IsRootCause})
    trap.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", trap.RuleName})

    trap.EntityData.YListKeys = []string {"EntryId"}

    return &(trap.EntityData)
}

// Snmp_Correlator_Traps_Trap_TrapInfo
// Correlated trap information
type Snmp_Correlator_Traps_Trap_TrapInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object ID. The type is string.
    Oid interface{}

    // Number of hsecs elapsed since snmpd was started. The type is interface{}
    // with range: 0..4294967295. Units are second.
    RelativeTimestamp interface{}

    // Time when the trap was generated. It is expressed in number of milliseconds
    // since 00:00 :00 UTC, January 1, 1970. The type is interface{} with range:
    // 0..18446744073709551615. Units are millisecond.
    Timestamp interface{}

    // VarBinds on the trap. The type is slice of
    // Snmp_Correlator_Traps_Trap_TrapInfo_VarBind.
    VarBind []*Snmp_Correlator_Traps_Trap_TrapInfo_VarBind
}

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetEntityData() *types.CommonEntityData {
    trapInfo.EntityData.YFilter = trapInfo.YFilter
    trapInfo.EntityData.YangName = "trap-info"
    trapInfo.EntityData.BundleName = "cisco_ios_xr"
    trapInfo.EntityData.ParentYangName = "trap"
    trapInfo.EntityData.SegmentPath = "trap-info"
    trapInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/traps/trap/" + trapInfo.EntityData.SegmentPath
    trapInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapInfo.EntityData.Children = types.NewOrderedMap()
    trapInfo.EntityData.Children.Append("var-bind", types.YChild{"VarBind", nil})
    for i := range trapInfo.VarBind {
        types.SetYListKey(trapInfo.VarBind[i], i)
        trapInfo.EntityData.Children.Append(types.GetSegmentPath(trapInfo.VarBind[i]), types.YChild{"VarBind", trapInfo.VarBind[i]})
    }
    trapInfo.EntityData.Leafs = types.NewOrderedMap()
    trapInfo.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", trapInfo.Oid})
    trapInfo.EntityData.Leafs.Append("relative-timestamp", types.YLeaf{"RelativeTimestamp", trapInfo.RelativeTimestamp})
    trapInfo.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", trapInfo.Timestamp})

    trapInfo.EntityData.YListKeys = []string {}

    return &(trapInfo.EntityData)
}

// Snmp_Correlator_Traps_Trap_TrapInfo_VarBind
// VarBinds on the trap
type Snmp_Correlator_Traps_Trap_TrapInfo_VarBind struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // OID of the varbind. The type is string.
    Oid interface{}

    // Value of the varbind. The type is string.
    Value interface{}
}

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetEntityData() *types.CommonEntityData {
    varBind.EntityData.YFilter = varBind.YFilter
    varBind.EntityData.YangName = "var-bind"
    varBind.EntityData.BundleName = "cisco_ios_xr"
    varBind.EntityData.ParentYangName = "trap-info"
    varBind.EntityData.SegmentPath = "var-bind" + types.AddNoKeyToken(varBind)
    varBind.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/correlator/traps/trap/trap-info/" + varBind.EntityData.SegmentPath
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = types.NewOrderedMap()
    varBind.EntityData.Leafs = types.NewOrderedMap()
    varBind.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", varBind.Oid})
    varBind.EntityData.Leafs.Append("value", types.YLeaf{"Value", varBind.Value})

    varBind.EntityData.YListKeys = []string {}

    return &(varBind.EntityData)
}

// Snmp_InterfaceIndexes
// List of index
type Snmp_InterfaceIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Index. The type is slice of Snmp_InterfaceIndexes_InterfaceIndex.
    InterfaceIndex []*Snmp_InterfaceIndexes_InterfaceIndex
}

func (interfaceIndexes *Snmp_InterfaceIndexes) GetEntityData() *types.CommonEntityData {
    interfaceIndexes.EntityData.YFilter = interfaceIndexes.YFilter
    interfaceIndexes.EntityData.YangName = "interface-indexes"
    interfaceIndexes.EntityData.BundleName = "cisco_ios_xr"
    interfaceIndexes.EntityData.ParentYangName = "snmp"
    interfaceIndexes.EntityData.SegmentPath = "interface-indexes"
    interfaceIndexes.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + interfaceIndexes.EntityData.SegmentPath
    interfaceIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIndexes.EntityData.Children = types.NewOrderedMap()
    interfaceIndexes.EntityData.Children.Append("interface-index", types.YChild{"InterfaceIndex", nil})
    for i := range interfaceIndexes.InterfaceIndex {
        interfaceIndexes.EntityData.Children.Append(types.GetSegmentPath(interfaceIndexes.InterfaceIndex[i]), types.YChild{"InterfaceIndex", interfaceIndexes.InterfaceIndex[i]})
    }
    interfaceIndexes.EntityData.Leafs = types.NewOrderedMap()

    interfaceIndexes.EntityData.YListKeys = []string {}

    return &(interfaceIndexes.EntityData)
}

// Snmp_InterfaceIndexes_InterfaceIndex
// Interface Index
type Snmp_InterfaceIndexes_InterfaceIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Index as used by MIB tables. The type is
    // interface{} with range: 0..4294967295.
    InterfaceIndex interface{}

    // Interface Name. The type is string. This attribute is mandatory.
    InterfaceName interface{}
}

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetEntityData() *types.CommonEntityData {
    interfaceIndex.EntityData.YFilter = interfaceIndex.YFilter
    interfaceIndex.EntityData.YangName = "interface-index"
    interfaceIndex.EntityData.BundleName = "cisco_ios_xr"
    interfaceIndex.EntityData.ParentYangName = "interface-indexes"
    interfaceIndex.EntityData.SegmentPath = "interface-index" + types.AddKeyToken(interfaceIndex.InterfaceIndex, "interface-index")
    interfaceIndex.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/interface-indexes/" + interfaceIndex.EntityData.SegmentPath
    interfaceIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIndex.EntityData.Children = types.NewOrderedMap()
    interfaceIndex.EntityData.Leafs = types.NewOrderedMap()
    interfaceIndex.EntityData.Leafs.Append("interface-index", types.YLeaf{"InterfaceIndex", interfaceIndex.InterfaceIndex})
    interfaceIndex.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceIndex.InterfaceName})

    interfaceIndex.EntityData.YListKeys = []string {"InterfaceIndex"}

    return &(interfaceIndex.EntityData)
}

// Snmp_IfIndexes
// List of ifnames
type Snmp_IfIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Index. The type is slice of Snmp_IfIndexes_IfIndex.
    IfIndex []*Snmp_IfIndexes_IfIndex
}

func (ifIndexes *Snmp_IfIndexes) GetEntityData() *types.CommonEntityData {
    ifIndexes.EntityData.YFilter = ifIndexes.YFilter
    ifIndexes.EntityData.YangName = "if-indexes"
    ifIndexes.EntityData.BundleName = "cisco_ios_xr"
    ifIndexes.EntityData.ParentYangName = "snmp"
    ifIndexes.EntityData.SegmentPath = "if-indexes"
    ifIndexes.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + ifIndexes.EntityData.SegmentPath
    ifIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifIndexes.EntityData.Children = types.NewOrderedMap()
    ifIndexes.EntityData.Children.Append("if-index", types.YChild{"IfIndex", nil})
    for i := range ifIndexes.IfIndex {
        ifIndexes.EntityData.Children.Append(types.GetSegmentPath(ifIndexes.IfIndex[i]), types.YChild{"IfIndex", ifIndexes.IfIndex[i]})
    }
    ifIndexes.EntityData.Leafs = types.NewOrderedMap()

    ifIndexes.EntityData.YListKeys = []string {}

    return &(ifIndexes.EntityData)
}

// Snmp_IfIndexes_IfIndex
// Interface Index
type Snmp_IfIndexes_IfIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Index as used by MIB tables. The type is
    // interface{} with range: 0..4294967295.
    InterfaceIndex interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (ifIndex *Snmp_IfIndexes_IfIndex) GetEntityData() *types.CommonEntityData {
    ifIndex.EntityData.YFilter = ifIndex.YFilter
    ifIndex.EntityData.YangName = "if-index"
    ifIndex.EntityData.BundleName = "cisco_ios_xr"
    ifIndex.EntityData.ParentYangName = "if-indexes"
    ifIndex.EntityData.SegmentPath = "if-index" + types.AddKeyToken(ifIndex.InterfaceIndex, "interface-index")
    ifIndex.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/if-indexes/" + ifIndex.EntityData.SegmentPath
    ifIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifIndex.EntityData.Children = types.NewOrderedMap()
    ifIndex.EntityData.Leafs = types.NewOrderedMap()
    ifIndex.EntityData.Leafs.Append("interface-index", types.YLeaf{"InterfaceIndex", ifIndex.InterfaceIndex})
    ifIndex.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ifIndex.InterfaceName})

    ifIndex.EntityData.YListKeys = []string {"InterfaceIndex"}

    return &(ifIndex.EntityData)
}

// Snmp_EntityMib
// SNMP entity mib
type Snmp_EntityMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP entity mib.
    EntityPhysicalIndexes Snmp_EntityMib_EntityPhysicalIndexes
}

func (entityMib *Snmp_EntityMib) GetEntityData() *types.CommonEntityData {
    entityMib.EntityData.YFilter = entityMib.YFilter
    entityMib.EntityData.YangName = "entity-mib"
    entityMib.EntityData.BundleName = "cisco_ios_xr"
    entityMib.EntityData.ParentYangName = "snmp"
    entityMib.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-entitymib-oper:entity-mib"
    entityMib.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + entityMib.EntityData.SegmentPath
    entityMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityMib.EntityData.Children = types.NewOrderedMap()
    entityMib.EntityData.Children.Append("entity-physical-indexes", types.YChild{"EntityPhysicalIndexes", &entityMib.EntityPhysicalIndexes})
    entityMib.EntityData.Leafs = types.NewOrderedMap()

    entityMib.EntityData.YListKeys = []string {}

    return &(entityMib.EntityData)
}

// Snmp_EntityMib_EntityPhysicalIndexes
// SNMP entity mib
type Snmp_EntityMib_EntityPhysicalIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP entPhysical index number. The type is slice of
    // Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex.
    EntityPhysicalIndex []*Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex
}

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetEntityData() *types.CommonEntityData {
    entityPhysicalIndexes.EntityData.YFilter = entityPhysicalIndexes.YFilter
    entityPhysicalIndexes.EntityData.YangName = "entity-physical-indexes"
    entityPhysicalIndexes.EntityData.BundleName = "cisco_ios_xr"
    entityPhysicalIndexes.EntityData.ParentYangName = "entity-mib"
    entityPhysicalIndexes.EntityData.SegmentPath = "entity-physical-indexes"
    entityPhysicalIndexes.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-entitymib-oper:entity-mib/" + entityPhysicalIndexes.EntityData.SegmentPath
    entityPhysicalIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityPhysicalIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityPhysicalIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityPhysicalIndexes.EntityData.Children = types.NewOrderedMap()
    entityPhysicalIndexes.EntityData.Children.Append("entity-physical-index", types.YChild{"EntityPhysicalIndex", nil})
    for i := range entityPhysicalIndexes.EntityPhysicalIndex {
        entityPhysicalIndexes.EntityData.Children.Append(types.GetSegmentPath(entityPhysicalIndexes.EntityPhysicalIndex[i]), types.YChild{"EntityPhysicalIndex", entityPhysicalIndexes.EntityPhysicalIndex[i]})
    }
    entityPhysicalIndexes.EntityData.Leafs = types.NewOrderedMap()

    entityPhysicalIndexes.EntityData.YListKeys = []string {}

    return &(entityPhysicalIndexes.EntityData)
}

// Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex
// SNMP entPhysical index number
type Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Entity physical index. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    EntityPhynum interface{}

    // entPhysicalIndex. The type is interface{} with range: 0..4294967295.
    PhysicalIndex interface{}

    // entPhysicalName. The type is string.
    EntPhysicalName interface{}

    // invmgr EDM path. The type is string.
    Location interface{}

    // EntPhysicalDescription. The type is string.
    EntPhysicalDescr interface{}

    // entphysicalFirmwareRev. The type is string.
    EntPhysicalFirmwareRev interface{}

    // entphysicalHardwareRev. The type is string.
    EntPhysicalHardwareRev interface{}

    // entphysicalModelName. The type is string.
    EntPhysicalModelname interface{}

    // entphysicalSerialNum. The type is string.
    EntPhysicalSerialNum interface{}

    // entphysicalSoftwareRev. The type is string.
    EntPhysicalSoftwareRev interface{}

    // entphysicalMfgName. The type is string.
    EntPhysicalMfgName interface{}
}

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetEntityData() *types.CommonEntityData {
    entityPhysicalIndex.EntityData.YFilter = entityPhysicalIndex.YFilter
    entityPhysicalIndex.EntityData.YangName = "entity-physical-index"
    entityPhysicalIndex.EntityData.BundleName = "cisco_ios_xr"
    entityPhysicalIndex.EntityData.ParentYangName = "entity-physical-indexes"
    entityPhysicalIndex.EntityData.SegmentPath = "entity-physical-index" + types.AddKeyToken(entityPhysicalIndex.EntityPhynum, "entity-phynum")
    entityPhysicalIndex.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-entitymib-oper:entity-mib/entity-physical-indexes/" + entityPhysicalIndex.EntityData.SegmentPath
    entityPhysicalIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityPhysicalIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityPhysicalIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityPhysicalIndex.EntityData.Children = types.NewOrderedMap()
    entityPhysicalIndex.EntityData.Leafs = types.NewOrderedMap()
    entityPhysicalIndex.EntityData.Leafs.Append("entity-phynum", types.YLeaf{"EntityPhynum", entityPhysicalIndex.EntityPhynum})
    entityPhysicalIndex.EntityData.Leafs.Append("physical-index", types.YLeaf{"PhysicalIndex", entityPhysicalIndex.PhysicalIndex})
    entityPhysicalIndex.EntityData.Leafs.Append("ent-physical-name", types.YLeaf{"EntPhysicalName", entityPhysicalIndex.EntPhysicalName})
    entityPhysicalIndex.EntityData.Leafs.Append("location", types.YLeaf{"Location", entityPhysicalIndex.Location})
    entityPhysicalIndex.EntityData.Leafs.Append("ent-physical-descr", types.YLeaf{"EntPhysicalDescr", entityPhysicalIndex.EntPhysicalDescr})
    entityPhysicalIndex.EntityData.Leafs.Append("ent-physical-firmware-rev", types.YLeaf{"EntPhysicalFirmwareRev", entityPhysicalIndex.EntPhysicalFirmwareRev})
    entityPhysicalIndex.EntityData.Leafs.Append("ent-physical-hardware-rev", types.YLeaf{"EntPhysicalHardwareRev", entityPhysicalIndex.EntPhysicalHardwareRev})
    entityPhysicalIndex.EntityData.Leafs.Append("ent-physical-modelname", types.YLeaf{"EntPhysicalModelname", entityPhysicalIndex.EntPhysicalModelname})
    entityPhysicalIndex.EntityData.Leafs.Append("ent-physical-serial-num", types.YLeaf{"EntPhysicalSerialNum", entityPhysicalIndex.EntPhysicalSerialNum})
    entityPhysicalIndex.EntityData.Leafs.Append("ent-physical-software-rev", types.YLeaf{"EntPhysicalSoftwareRev", entityPhysicalIndex.EntPhysicalSoftwareRev})
    entityPhysicalIndex.EntityData.Leafs.Append("ent-physical-mfg-name", types.YLeaf{"EntPhysicalMfgName", entityPhysicalIndex.EntPhysicalMfgName})

    entityPhysicalIndex.EntityData.YListKeys = []string {"EntityPhynum"}

    return &(entityPhysicalIndex.EntityData)
}

// Snmp_InterfaceMib
// SNMP IF-MIB information
type Snmp_InterfaceMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interfaces ifIndex information.
    Interfaces Snmp_InterfaceMib_Interfaces

    // Interfaces ifConnectorPresent information.
    InterfaceConnectors Snmp_InterfaceMib_InterfaceConnectors

    // Interfaces ifAlias information.
    InterfaceAliases Snmp_InterfaceMib_InterfaceAliases

    // Interfaces Notification information.
    NotificationInterfaces Snmp_InterfaceMib_NotificationInterfaces

    // Interfaces ifstackstatus information.
    InterfaceStackStatuses Snmp_InterfaceMib_InterfaceStackStatuses
}

func (interfaceMib *Snmp_InterfaceMib) GetEntityData() *types.CommonEntityData {
    interfaceMib.EntityData.YFilter = interfaceMib.YFilter
    interfaceMib.EntityData.YangName = "interface-mib"
    interfaceMib.EntityData.BundleName = "cisco_ios_xr"
    interfaceMib.EntityData.ParentYangName = "snmp"
    interfaceMib.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-ifmib-oper:interface-mib"
    interfaceMib.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + interfaceMib.EntityData.SegmentPath
    interfaceMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMib.EntityData.Children = types.NewOrderedMap()
    interfaceMib.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceMib.Interfaces})
    interfaceMib.EntityData.Children.Append("interface-connectors", types.YChild{"InterfaceConnectors", &interfaceMib.InterfaceConnectors})
    interfaceMib.EntityData.Children.Append("interface-aliases", types.YChild{"InterfaceAliases", &interfaceMib.InterfaceAliases})
    interfaceMib.EntityData.Children.Append("notification-interfaces", types.YChild{"NotificationInterfaces", &interfaceMib.NotificationInterfaces})
    interfaceMib.EntityData.Children.Append("interface-stack-statuses", types.YChild{"InterfaceStackStatuses", &interfaceMib.InterfaceStackStatuses})
    interfaceMib.EntityData.Leafs = types.NewOrderedMap()

    interfaceMib.EntityData.YListKeys = []string {}

    return &(interfaceMib.EntityData)
}

// Snmp_InterfaceMib_Interfaces
// Interfaces ifIndex information
type Snmp_InterfaceMib_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ifIndex for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_Interfaces_Interface.
    Interface []*Snmp_InterfaceMib_Interfaces_Interface
}

func (interfaces *Snmp_InterfaceMib_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-mib"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Snmp_InterfaceMib_Interfaces_Interface
// ifIndex for a specific Interface Name
type Snmp_InterfaceMib_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface Index. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", self.IfIndex})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Snmp_InterfaceMib_InterfaceConnectors
// Interfaces ifConnectorPresent information
type Snmp_InterfaceMib_InterfaceConnectors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ifConnectorPresent for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector.
    InterfaceConnector []*Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector
}

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetEntityData() *types.CommonEntityData {
    interfaceConnectors.EntityData.YFilter = interfaceConnectors.YFilter
    interfaceConnectors.EntityData.YangName = "interface-connectors"
    interfaceConnectors.EntityData.BundleName = "cisco_ios_xr"
    interfaceConnectors.EntityData.ParentYangName = "interface-mib"
    interfaceConnectors.EntityData.SegmentPath = "interface-connectors"
    interfaceConnectors.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/" + interfaceConnectors.EntityData.SegmentPath
    interfaceConnectors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConnectors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConnectors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConnectors.EntityData.Children = types.NewOrderedMap()
    interfaceConnectors.EntityData.Children.Append("interface-connector", types.YChild{"InterfaceConnector", nil})
    for i := range interfaceConnectors.InterfaceConnector {
        interfaceConnectors.EntityData.Children.Append(types.GetSegmentPath(interfaceConnectors.InterfaceConnector[i]), types.YChild{"InterfaceConnector", interfaceConnectors.InterfaceConnector[i]})
    }
    interfaceConnectors.EntityData.Leafs = types.NewOrderedMap()

    interfaceConnectors.EntityData.YListKeys = []string {}

    return &(interfaceConnectors.EntityData)
}

// Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector
// ifConnectorPresent for a specific Interface
// Name
type Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface ifConnector. The type is string.
    IfConnectorPresent interface{}
}

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetEntityData() *types.CommonEntityData {
    interfaceConnector.EntityData.YFilter = interfaceConnector.YFilter
    interfaceConnector.EntityData.YangName = "interface-connector"
    interfaceConnector.EntityData.BundleName = "cisco_ios_xr"
    interfaceConnector.EntityData.ParentYangName = "interface-connectors"
    interfaceConnector.EntityData.SegmentPath = "interface-connector" + types.AddKeyToken(interfaceConnector.InterfaceName, "interface-name")
    interfaceConnector.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/interface-connectors/" + interfaceConnector.EntityData.SegmentPath
    interfaceConnector.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConnector.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConnector.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConnector.EntityData.Children = types.NewOrderedMap()
    interfaceConnector.EntityData.Leafs = types.NewOrderedMap()
    interfaceConnector.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceConnector.InterfaceName})
    interfaceConnector.EntityData.Leafs.Append("if-connector-present", types.YLeaf{"IfConnectorPresent", interfaceConnector.IfConnectorPresent})

    interfaceConnector.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceConnector.EntityData)
}

// Snmp_InterfaceMib_InterfaceAliases
// Interfaces ifAlias information
type Snmp_InterfaceMib_InterfaceAliases struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ifAlias for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias.
    InterfaceAlias []*Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias
}

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetEntityData() *types.CommonEntityData {
    interfaceAliases.EntityData.YFilter = interfaceAliases.YFilter
    interfaceAliases.EntityData.YangName = "interface-aliases"
    interfaceAliases.EntityData.BundleName = "cisco_ios_xr"
    interfaceAliases.EntityData.ParentYangName = "interface-mib"
    interfaceAliases.EntityData.SegmentPath = "interface-aliases"
    interfaceAliases.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/" + interfaceAliases.EntityData.SegmentPath
    interfaceAliases.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAliases.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAliases.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAliases.EntityData.Children = types.NewOrderedMap()
    interfaceAliases.EntityData.Children.Append("interface-alias", types.YChild{"InterfaceAlias", nil})
    for i := range interfaceAliases.InterfaceAlias {
        interfaceAliases.EntityData.Children.Append(types.GetSegmentPath(interfaceAliases.InterfaceAlias[i]), types.YChild{"InterfaceAlias", interfaceAliases.InterfaceAlias[i]})
    }
    interfaceAliases.EntityData.Leafs = types.NewOrderedMap()

    interfaceAliases.EntityData.YListKeys = []string {}

    return &(interfaceAliases.EntityData)
}

// Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias
// ifAlias for a specific Interface Name
type Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Interface ifAlias. The type is string.
    IfAlias interface{}
}

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetEntityData() *types.CommonEntityData {
    interfaceAlias.EntityData.YFilter = interfaceAlias.YFilter
    interfaceAlias.EntityData.YangName = "interface-alias"
    interfaceAlias.EntityData.BundleName = "cisco_ios_xr"
    interfaceAlias.EntityData.ParentYangName = "interface-aliases"
    interfaceAlias.EntityData.SegmentPath = "interface-alias" + types.AddKeyToken(interfaceAlias.InterfaceName, "interface-name")
    interfaceAlias.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/interface-aliases/" + interfaceAlias.EntityData.SegmentPath
    interfaceAlias.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAlias.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAlias.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAlias.EntityData.Children = types.NewOrderedMap()
    interfaceAlias.EntityData.Leafs = types.NewOrderedMap()
    interfaceAlias.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceAlias.InterfaceName})
    interfaceAlias.EntityData.Leafs.Append("if-alias", types.YLeaf{"IfAlias", interfaceAlias.IfAlias})

    interfaceAlias.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceAlias.EntityData)
}

// Snmp_InterfaceMib_NotificationInterfaces
// Interfaces Notification information
type Snmp_InterfaceMib_NotificationInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Notification for specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface.
    NotificationInterface []*Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface
}

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetEntityData() *types.CommonEntityData {
    notificationInterfaces.EntityData.YFilter = notificationInterfaces.YFilter
    notificationInterfaces.EntityData.YangName = "notification-interfaces"
    notificationInterfaces.EntityData.BundleName = "cisco_ios_xr"
    notificationInterfaces.EntityData.ParentYangName = "interface-mib"
    notificationInterfaces.EntityData.SegmentPath = "notification-interfaces"
    notificationInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/" + notificationInterfaces.EntityData.SegmentPath
    notificationInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationInterfaces.EntityData.Children = types.NewOrderedMap()
    notificationInterfaces.EntityData.Children.Append("notification-interface", types.YChild{"NotificationInterface", nil})
    for i := range notificationInterfaces.NotificationInterface {
        notificationInterfaces.EntityData.Children.Append(types.GetSegmentPath(notificationInterfaces.NotificationInterface[i]), types.YChild{"NotificationInterface", notificationInterfaces.NotificationInterface[i]})
    }
    notificationInterfaces.EntityData.Leafs = types.NewOrderedMap()

    notificationInterfaces.EntityData.YListKeys = []string {}

    return &(notificationInterfaces.EntityData)
}

// Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface
// Notification for specific Interface Name
type Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // LinkUpDown notification status. The type is LinkUpDownStatus.
    LinkUpDownNotifStatus interface{}
}

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetEntityData() *types.CommonEntityData {
    notificationInterface.EntityData.YFilter = notificationInterface.YFilter
    notificationInterface.EntityData.YangName = "notification-interface"
    notificationInterface.EntityData.BundleName = "cisco_ios_xr"
    notificationInterface.EntityData.ParentYangName = "notification-interfaces"
    notificationInterface.EntityData.SegmentPath = "notification-interface" + types.AddKeyToken(notificationInterface.InterfaceName, "interface-name")
    notificationInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/notification-interfaces/" + notificationInterface.EntityData.SegmentPath
    notificationInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationInterface.EntityData.Children = types.NewOrderedMap()
    notificationInterface.EntityData.Leafs = types.NewOrderedMap()
    notificationInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", notificationInterface.InterfaceName})
    notificationInterface.EntityData.Leafs.Append("link-up-down-notif-status", types.YLeaf{"LinkUpDownNotifStatus", notificationInterface.LinkUpDownNotifStatus})

    notificationInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(notificationInterface.EntityData)
}

// Snmp_InterfaceMib_InterfaceStackStatuses
// Interfaces ifstackstatus information
type Snmp_InterfaceMib_InterfaceStackStatuses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ifstatus for a pair of Interface. The type is slice of
    // Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus.
    InterfaceStackStatus []*Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus
}

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetEntityData() *types.CommonEntityData {
    interfaceStackStatuses.EntityData.YFilter = interfaceStackStatuses.YFilter
    interfaceStackStatuses.EntityData.YangName = "interface-stack-statuses"
    interfaceStackStatuses.EntityData.BundleName = "cisco_ios_xr"
    interfaceStackStatuses.EntityData.ParentYangName = "interface-mib"
    interfaceStackStatuses.EntityData.SegmentPath = "interface-stack-statuses"
    interfaceStackStatuses.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/" + interfaceStackStatuses.EntityData.SegmentPath
    interfaceStackStatuses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStackStatuses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStackStatuses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStackStatuses.EntityData.Children = types.NewOrderedMap()
    interfaceStackStatuses.EntityData.Children.Append("interface-stack-status", types.YChild{"InterfaceStackStatus", nil})
    for i := range interfaceStackStatuses.InterfaceStackStatus {
        interfaceStackStatuses.EntityData.Children.Append(types.GetSegmentPath(interfaceStackStatuses.InterfaceStackStatus[i]), types.YChild{"InterfaceStackStatus", interfaceStackStatuses.InterfaceStackStatus[i]})
    }
    interfaceStackStatuses.EntityData.Leafs = types.NewOrderedMap()

    interfaceStackStatuses.EntityData.YListKeys = []string {}

    return &(interfaceStackStatuses.EntityData)
}

// Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus
// ifstatus for a pair of Interface
type Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. StackHigherLayer.StackLowerLayer. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceStackStatus interface{}

    // Higher Layer Index. The type is string.
    IfStackHigherLayer interface{}

    // Lowyer Layer Index. The type is string.
    IfStackLowerLayer interface{}

    // Interface ifStackStaus info. The type is string.
    IfStackStatus interface{}
}

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetEntityData() *types.CommonEntityData {
    interfaceStackStatus.EntityData.YFilter = interfaceStackStatus.YFilter
    interfaceStackStatus.EntityData.YangName = "interface-stack-status"
    interfaceStackStatus.EntityData.BundleName = "cisco_ios_xr"
    interfaceStackStatus.EntityData.ParentYangName = "interface-stack-statuses"
    interfaceStackStatus.EntityData.SegmentPath = "interface-stack-status" + types.AddKeyToken(interfaceStackStatus.InterfaceStackStatus, "interface-stack-status")
    interfaceStackStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-ifmib-oper:interface-mib/interface-stack-statuses/" + interfaceStackStatus.EntityData.SegmentPath
    interfaceStackStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStackStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStackStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStackStatus.EntityData.Children = types.NewOrderedMap()
    interfaceStackStatus.EntityData.Leafs = types.NewOrderedMap()
    interfaceStackStatus.EntityData.Leafs.Append("interface-stack-status", types.YLeaf{"InterfaceStackStatus", interfaceStackStatus.InterfaceStackStatus})
    interfaceStackStatus.EntityData.Leafs.Append("if-stack-higher-layer", types.YLeaf{"IfStackHigherLayer", interfaceStackStatus.IfStackHigherLayer})
    interfaceStackStatus.EntityData.Leafs.Append("if-stack-lower-layer", types.YLeaf{"IfStackLowerLayer", interfaceStackStatus.IfStackLowerLayer})
    interfaceStackStatus.EntityData.Leafs.Append("if-stack-status", types.YLeaf{"IfStackStatus", interfaceStackStatus.IfStackStatus})

    interfaceStackStatus.EntityData.YListKeys = []string {"InterfaceStackStatus"}

    return &(interfaceStackStatus.EntityData)
}

// Snmp_SensorMib
// SNMP sensor MIB information
type Snmp_SensorMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of physical index table for threshold value.
    PhysicalIndexes Snmp_SensorMib_PhysicalIndexes

    // List of physical index .
    EntPhyIndexes Snmp_SensorMib_EntPhyIndexes
}

func (sensorMib *Snmp_SensorMib) GetEntityData() *types.CommonEntityData {
    sensorMib.EntityData.YFilter = sensorMib.YFilter
    sensorMib.EntityData.YangName = "sensor-mib"
    sensorMib.EntityData.BundleName = "cisco_ios_xr"
    sensorMib.EntityData.ParentYangName = "snmp"
    sensorMib.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib"
    sensorMib.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/" + sensorMib.EntityData.SegmentPath
    sensorMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorMib.EntityData.Children = types.NewOrderedMap()
    sensorMib.EntityData.Children.Append("physical-indexes", types.YChild{"PhysicalIndexes", &sensorMib.PhysicalIndexes})
    sensorMib.EntityData.Children.Append("ent-phy-indexes", types.YChild{"EntPhyIndexes", &sensorMib.EntPhyIndexes})
    sensorMib.EntityData.Leafs = types.NewOrderedMap()

    sensorMib.EntityData.YListKeys = []string {}

    return &(sensorMib.EntityData)
}

// Snmp_SensorMib_PhysicalIndexes
// List of physical index table for threshold
// value
type Snmp_SensorMib_PhysicalIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold value for physical index. The type is slice of
    // Snmp_SensorMib_PhysicalIndexes_PhysicalIndex.
    PhysicalIndex []*Snmp_SensorMib_PhysicalIndexes_PhysicalIndex
}

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetEntityData() *types.CommonEntityData {
    physicalIndexes.EntityData.YFilter = physicalIndexes.YFilter
    physicalIndexes.EntityData.YangName = "physical-indexes"
    physicalIndexes.EntityData.BundleName = "cisco_ios_xr"
    physicalIndexes.EntityData.ParentYangName = "sensor-mib"
    physicalIndexes.EntityData.SegmentPath = "physical-indexes"
    physicalIndexes.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib/" + physicalIndexes.EntityData.SegmentPath
    physicalIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    physicalIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    physicalIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    physicalIndexes.EntityData.Children = types.NewOrderedMap()
    physicalIndexes.EntityData.Children.Append("physical-index", types.YChild{"PhysicalIndex", nil})
    for i := range physicalIndexes.PhysicalIndex {
        physicalIndexes.EntityData.Children.Append(types.GetSegmentPath(physicalIndexes.PhysicalIndex[i]), types.YChild{"PhysicalIndex", physicalIndexes.PhysicalIndex[i]})
    }
    physicalIndexes.EntityData.Leafs = types.NewOrderedMap()

    physicalIndexes.EntityData.YListKeys = []string {}

    return &(physicalIndexes.EntityData)
}

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex
// Threshold value for physical index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Physical index. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Index interface{}

    // List of threshold index.
    ThresholdIndexes Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes
}

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetEntityData() *types.CommonEntityData {
    physicalIndex.EntityData.YFilter = physicalIndex.YFilter
    physicalIndex.EntityData.YangName = "physical-index"
    physicalIndex.EntityData.BundleName = "cisco_ios_xr"
    physicalIndex.EntityData.ParentYangName = "physical-indexes"
    physicalIndex.EntityData.SegmentPath = "physical-index" + types.AddKeyToken(physicalIndex.Index, "index")
    physicalIndex.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib/physical-indexes/" + physicalIndex.EntityData.SegmentPath
    physicalIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    physicalIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    physicalIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    physicalIndex.EntityData.Children = types.NewOrderedMap()
    physicalIndex.EntityData.Children.Append("threshold-indexes", types.YChild{"ThresholdIndexes", &physicalIndex.ThresholdIndexes})
    physicalIndex.EntityData.Leafs = types.NewOrderedMap()
    physicalIndex.EntityData.Leafs.Append("index", types.YLeaf{"Index", physicalIndex.Index})

    physicalIndex.EntityData.YListKeys = []string {"Index"}

    return &(physicalIndex.EntityData)
}

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes
// List of threshold index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold value for threshold index. The type is slice of
    // Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex.
    ThresholdIndex []*Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex
}

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetEntityData() *types.CommonEntityData {
    thresholdIndexes.EntityData.YFilter = thresholdIndexes.YFilter
    thresholdIndexes.EntityData.YangName = "threshold-indexes"
    thresholdIndexes.EntityData.BundleName = "cisco_ios_xr"
    thresholdIndexes.EntityData.ParentYangName = "physical-index"
    thresholdIndexes.EntityData.SegmentPath = "threshold-indexes"
    thresholdIndexes.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib/physical-indexes/physical-index/" + thresholdIndexes.EntityData.SegmentPath
    thresholdIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdIndexes.EntityData.Children = types.NewOrderedMap()
    thresholdIndexes.EntityData.Children.Append("threshold-index", types.YChild{"ThresholdIndex", nil})
    for i := range thresholdIndexes.ThresholdIndex {
        types.SetYListKey(thresholdIndexes.ThresholdIndex[i], i)
        thresholdIndexes.EntityData.Children.Append(types.GetSegmentPath(thresholdIndexes.ThresholdIndex[i]), types.YChild{"ThresholdIndex", thresholdIndexes.ThresholdIndex[i]})
    }
    thresholdIndexes.EntityData.Leafs = types.NewOrderedMap()

    thresholdIndexes.EntityData.YListKeys = []string {}

    return &(thresholdIndexes.EntityData)
}

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex
// Threshold value for threshold index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Physical Index. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PhyIndex interface{}

    // Threshold index. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ThreIndex interface{}

    // Indicates minor, major, critical severities. The type is interface{} with
    // range: 0..4294967295.
    ThresholdSeverity interface{}

    // Indicates relation between sensor value and threshold. The type is
    // interface{} with range: 0..4294967295.
    ThresholdRelation interface{}

    // Value of the configured threshold. The type is interface{} with range:
    // 0..4294967295.
    ThresholdValue interface{}

    // Indicates the result of the most recent evaluation of the thresholD. The
    // type is bool.
    ThresholdEvaluation interface{}

    // Indicates whether or not a notification should result, in case of threshold
    // violation. The type is bool.
    ThresholdNotificationEnabled interface{}
}

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetEntityData() *types.CommonEntityData {
    thresholdIndex.EntityData.YFilter = thresholdIndex.YFilter
    thresholdIndex.EntityData.YangName = "threshold-index"
    thresholdIndex.EntityData.BundleName = "cisco_ios_xr"
    thresholdIndex.EntityData.ParentYangName = "threshold-indexes"
    thresholdIndex.EntityData.SegmentPath = "threshold-index" + types.AddNoKeyToken(thresholdIndex)
    thresholdIndex.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib/physical-indexes/physical-index/threshold-indexes/" + thresholdIndex.EntityData.SegmentPath
    thresholdIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdIndex.EntityData.Children = types.NewOrderedMap()
    thresholdIndex.EntityData.Leafs = types.NewOrderedMap()
    thresholdIndex.EntityData.Leafs.Append("phy-index", types.YLeaf{"PhyIndex", thresholdIndex.PhyIndex})
    thresholdIndex.EntityData.Leafs.Append("thre-index", types.YLeaf{"ThreIndex", thresholdIndex.ThreIndex})
    thresholdIndex.EntityData.Leafs.Append("threshold-severity", types.YLeaf{"ThresholdSeverity", thresholdIndex.ThresholdSeverity})
    thresholdIndex.EntityData.Leafs.Append("threshold-relation", types.YLeaf{"ThresholdRelation", thresholdIndex.ThresholdRelation})
    thresholdIndex.EntityData.Leafs.Append("threshold-value", types.YLeaf{"ThresholdValue", thresholdIndex.ThresholdValue})
    thresholdIndex.EntityData.Leafs.Append("threshold-evaluation", types.YLeaf{"ThresholdEvaluation", thresholdIndex.ThresholdEvaluation})
    thresholdIndex.EntityData.Leafs.Append("threshold-notification-enabled", types.YLeaf{"ThresholdNotificationEnabled", thresholdIndex.ThresholdNotificationEnabled})

    thresholdIndex.EntityData.YListKeys = []string {}

    return &(thresholdIndex.EntityData)
}

// Snmp_SensorMib_EntPhyIndexes
// List of physical index 
type Snmp_SensorMib_EntPhyIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor value for physical index. The type is slice of
    // Snmp_SensorMib_EntPhyIndexes_EntPhyIndex.
    EntPhyIndex []*Snmp_SensorMib_EntPhyIndexes_EntPhyIndex
}

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetEntityData() *types.CommonEntityData {
    entPhyIndexes.EntityData.YFilter = entPhyIndexes.YFilter
    entPhyIndexes.EntityData.YangName = "ent-phy-indexes"
    entPhyIndexes.EntityData.BundleName = "cisco_ios_xr"
    entPhyIndexes.EntityData.ParentYangName = "sensor-mib"
    entPhyIndexes.EntityData.SegmentPath = "ent-phy-indexes"
    entPhyIndexes.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib/" + entPhyIndexes.EntityData.SegmentPath
    entPhyIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entPhyIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entPhyIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entPhyIndexes.EntityData.Children = types.NewOrderedMap()
    entPhyIndexes.EntityData.Children.Append("ent-phy-index", types.YChild{"EntPhyIndex", nil})
    for i := range entPhyIndexes.EntPhyIndex {
        entPhyIndexes.EntityData.Children.Append(types.GetSegmentPath(entPhyIndexes.EntPhyIndex[i]), types.YChild{"EntPhyIndex", entPhyIndexes.EntPhyIndex[i]})
    }
    entPhyIndexes.EntityData.Leafs = types.NewOrderedMap()

    entPhyIndexes.EntityData.YListKeys = []string {}

    return &(entPhyIndexes.EntityData)
}

// Snmp_SensorMib_EntPhyIndexes_EntPhyIndex
// Sensor value for physical index
type Snmp_SensorMib_EntPhyIndexes_EntPhyIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Physical index. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Index interface{}

    // Sensor valid bitmap. The type is interface{} with range: 0..4294967295.
    FieldValidityBitmap interface{}

    // Device Name. The type is string with length: 0..64.
    DeviceDescription interface{}

    // Units of variable being read. The type is string with length: 0..64.
    Units interface{}

    // Identifier for this device. The type is interface{} with range:
    // 0..4294967295.
    DeviceId interface{}

    // Current reading of sensor. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Indicates threshold violation. The type is interface{} with range:
    // 0..4294967295.
    AlarmType interface{}

    // Sensor data type enums. The type is interface{} with range: 0..4294967295.
    DataType interface{}

    // Sensor scale enums. The type is interface{} with range: 0..4294967295.
    Scale interface{}

    // Sensor precision range. The type is interface{} with range: 0..4294967295.
    Precision interface{}

    // Sensor operation state enums. The type is interface{} with range:
    // 0..4294967295.
    Status interface{}

    // Age of the sensor value; set to the current time if directly access the
    // value from sensor. The type is interface{} with range: 0..4294967295.
    AgeTimeStamp interface{}

    // Sensor value update rate;set to 0 if sensor value is updated and evaluated
    // immediately. The type is interface{} with range: 0..4294967295.
    UpdateRate interface{}

    // physical entity for which the sensor is taking measurements. The type is
    // interface{} with range: 0..4294967295.
    MeasuredEntity interface{}
}

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetEntityData() *types.CommonEntityData {
    entPhyIndex.EntityData.YFilter = entPhyIndex.YFilter
    entPhyIndex.EntityData.YangName = "ent-phy-index"
    entPhyIndex.EntityData.BundleName = "cisco_ios_xr"
    entPhyIndex.EntityData.ParentYangName = "ent-phy-indexes"
    entPhyIndex.EntityData.SegmentPath = "ent-phy-index" + types.AddKeyToken(entPhyIndex.Index, "index")
    entPhyIndex.EntityData.AbsolutePath = "Cisco-IOS-XR-snmp-agent-oper:snmp/Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib/ent-phy-indexes/" + entPhyIndex.EntityData.SegmentPath
    entPhyIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entPhyIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entPhyIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entPhyIndex.EntityData.Children = types.NewOrderedMap()
    entPhyIndex.EntityData.Leafs = types.NewOrderedMap()
    entPhyIndex.EntityData.Leafs.Append("index", types.YLeaf{"Index", entPhyIndex.Index})
    entPhyIndex.EntityData.Leafs.Append("field-validity-bitmap", types.YLeaf{"FieldValidityBitmap", entPhyIndex.FieldValidityBitmap})
    entPhyIndex.EntityData.Leafs.Append("device-description", types.YLeaf{"DeviceDescription", entPhyIndex.DeviceDescription})
    entPhyIndex.EntityData.Leafs.Append("units", types.YLeaf{"Units", entPhyIndex.Units})
    entPhyIndex.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", entPhyIndex.DeviceId})
    entPhyIndex.EntityData.Leafs.Append("value", types.YLeaf{"Value", entPhyIndex.Value})
    entPhyIndex.EntityData.Leafs.Append("alarm-type", types.YLeaf{"AlarmType", entPhyIndex.AlarmType})
    entPhyIndex.EntityData.Leafs.Append("data-type", types.YLeaf{"DataType", entPhyIndex.DataType})
    entPhyIndex.EntityData.Leafs.Append("scale", types.YLeaf{"Scale", entPhyIndex.Scale})
    entPhyIndex.EntityData.Leafs.Append("precision", types.YLeaf{"Precision", entPhyIndex.Precision})
    entPhyIndex.EntityData.Leafs.Append("status", types.YLeaf{"Status", entPhyIndex.Status})
    entPhyIndex.EntityData.Leafs.Append("age-time-stamp", types.YLeaf{"AgeTimeStamp", entPhyIndex.AgeTimeStamp})
    entPhyIndex.EntityData.Leafs.Append("update-rate", types.YLeaf{"UpdateRate", entPhyIndex.UpdateRate})
    entPhyIndex.EntityData.Leafs.Append("measured-entity", types.YLeaf{"MeasuredEntity", entPhyIndex.MeasuredEntity})

    entPhyIndex.EntityData.YListKeys = []string {"Index"}

    return &(entPhyIndex.EntityData)
}

