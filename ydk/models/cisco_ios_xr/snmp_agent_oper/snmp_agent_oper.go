// This module contains a collection of YANG definitions
// for Cisco IOS-XR snmp-agent package operational data.
// 
// This module contains definitions
// for the following management objects:
//   snmp: SNMP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// DupReqDropStatus represents Dup req drop status
type DupReqDropStatus string

const (
    // Disabled
    DupReqDropStatus_disabled DupReqDropStatus = "disabled"

    // Enabled
    DupReqDropStatus_enabled DupReqDropStatus = "enabled"
)

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
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = make(map[string]types.YChild)
    snmp.EntityData.Children["trap-servers"] = types.YChild{"TrapServers", &snmp.TrapServers}
    snmp.EntityData.Children["information"] = types.YChild{"Information", &snmp.Information}
    snmp.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &snmp.Interfaces}
    snmp.EntityData.Children["correlator"] = types.YChild{"Correlator", &snmp.Correlator}
    snmp.EntityData.Children["interface-indexes"] = types.YChild{"InterfaceIndexes", &snmp.InterfaceIndexes}
    snmp.EntityData.Children["if-indexes"] = types.YChild{"IfIndexes", &snmp.IfIndexes}
    snmp.EntityData.Children["Cisco-IOS-XR-snmp-entitymib-oper:entity-mib"] = types.YChild{"EntityMib", &snmp.EntityMib}
    snmp.EntityData.Children["Cisco-IOS-XR-snmp-ifmib-oper:interface-mib"] = types.YChild{"InterfaceMib", &snmp.InterfaceMib}
    snmp.EntityData.Children["Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib"] = types.YChild{"SensorMib", &snmp.SensorMib}
    snmp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(snmp.EntityData)
}

// Snmp_TrapServers
// List of trap hosts
type Snmp_TrapServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trap server and port to which the trap is to be sent and statistics. The
    // type is slice of Snmp_TrapServers_TrapServer.
    TrapServer []Snmp_TrapServers_TrapServer
}

func (trapServers *Snmp_TrapServers) GetEntityData() *types.CommonEntityData {
    trapServers.EntityData.YFilter = trapServers.YFilter
    trapServers.EntityData.YangName = "trap-servers"
    trapServers.EntityData.BundleName = "cisco_ios_xr"
    trapServers.EntityData.ParentYangName = "snmp"
    trapServers.EntityData.SegmentPath = "trap-servers"
    trapServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapServers.EntityData.Children = make(map[string]types.YChild)
    trapServers.EntityData.Children["trap-server"] = types.YChild{"TrapServer", nil}
    for i := range trapServers.TrapServer {
        trapServers.EntityData.Children[types.GetSegmentPath(&trapServers.TrapServer[i])] = types.YChild{"TrapServer", &trapServers.TrapServer[i]}
    }
    trapServers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trapServers.EntityData)
}

// Snmp_TrapServers_TrapServer
// Trap server and port to which the trap is to be
// sent and statistics
type Snmp_TrapServers_TrapServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    trapServer.EntityData.SegmentPath = "trap-server"
    trapServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapServer.EntityData.Children = make(map[string]types.YChild)
    trapServer.EntityData.Leafs = make(map[string]types.YLeaf)
    trapServer.EntityData.Leafs["trap-host"] = types.YLeaf{"TrapHost", trapServer.TrapHost}
    trapServer.EntityData.Leafs["port"] = types.YLeaf{"Port", trapServer.Port}
    trapServer.EntityData.Leafs["number-of-pkts-in-trap-q"] = types.YLeaf{"NumberOfPktsInTrapQ", trapServer.NumberOfPktsInTrapQ}
    trapServer.EntityData.Leafs["max-q-length-of-trap-q"] = types.YLeaf{"MaxQLengthOfTrapQ", trapServer.MaxQLengthOfTrapQ}
    trapServer.EntityData.Leafs["number-of-pkts-sent"] = types.YLeaf{"NumberOfPktsSent", trapServer.NumberOfPktsSent}
    trapServer.EntityData.Leafs["number-of-pkts-dropped"] = types.YLeaf{"NumberOfPktsDropped", trapServer.NumberOfPktsDropped}
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
    information.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    information.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    information.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    information.EntityData.Children = make(map[string]types.YChild)
    information.EntityData.Children["hosts"] = types.YChild{"Hosts", &information.Hosts}
    information.EntityData.Children["system-up-time"] = types.YChild{"SystemUpTime", &information.SystemUpTime}
    information.EntityData.Children["nms-addresses"] = types.YChild{"NmsAddresses", &information.NmsAddresses}
    information.EntityData.Children["engine-id"] = types.YChild{"EngineId", &information.EngineId}
    information.EntityData.Children["rx-queue"] = types.YChild{"RxQueue", &information.RxQueue}
    information.EntityData.Children["system-name"] = types.YChild{"SystemName", &information.SystemName}
    information.EntityData.Children["request-type-detail"] = types.YChild{"RequestTypeDetail", &information.RequestTypeDetail}
    information.EntityData.Children["duplicate-drop"] = types.YChild{"DuplicateDrop", &information.DuplicateDrop}
    information.EntityData.Children["bulk-stats-transfers"] = types.YChild{"BulkStatsTransfers", &information.BulkStatsTransfers}
    information.EntityData.Children["trap-infos"] = types.YChild{"TrapInfos", &information.TrapInfos}
    information.EntityData.Children["poll-oids"] = types.YChild{"PollOids", &information.PollOids}
    information.EntityData.Children["infom-details"] = types.YChild{"InfomDetails", &information.InfomDetails}
    information.EntityData.Children["statistics"] = types.YChild{"Statistics", &information.Statistics}
    information.EntityData.Children["incoming-queue"] = types.YChild{"IncomingQueue", &information.IncomingQueue}
    information.EntityData.Children["context-mapping"] = types.YChild{"ContextMapping", &information.ContextMapping}
    information.EntityData.Children["trap-oids"] = types.YChild{"TrapOids", &information.TrapOids}
    information.EntityData.Children["nm-spackets"] = types.YChild{"NmSpackets", &information.NmSpackets}
    information.EntityData.Children["mibs"] = types.YChild{"Mibs", &information.Mibs}
    information.EntityData.Children["serial-numbers"] = types.YChild{"SerialNumbers", &information.SerialNumbers}
    information.EntityData.Children["drop-nms-addresses"] = types.YChild{"DropNmsAddresses", &information.DropNmsAddresses}
    information.EntityData.Children["views"] = types.YChild{"Views", &information.Views}
    information.EntityData.Children["system-descr"] = types.YChild{"SystemDescr", &information.SystemDescr}
    information.EntityData.Children["tables"] = types.YChild{"Tables", &information.Tables}
    information.EntityData.Children["system-oid"] = types.YChild{"SystemOid", &information.SystemOid}
    information.EntityData.Children["trap-queue"] = types.YChild{"TrapQueue", &information.TrapQueue}
    information.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(information.EntityData)
}

// Snmp_Information_Hosts
// SNMP host information
type Snmp_Information_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP target host name. The type is slice of Snmp_Information_Hosts_Host.
    Host []Snmp_Information_Hosts_Host
}

func (hosts *Snmp_Information_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "information"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = make(map[string]types.YChild)
    hosts.EntityData.Children["host"] = types.YChild{"Host", nil}
    for i := range hosts.Host {
        hosts.EntityData.Children[types.GetSegmentPath(&hosts.Host[i])] = types.YChild{"Host", &hosts.Host[i]}
    }
    hosts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hosts.EntityData)
}

// Snmp_Information_Hosts_Host
// SNMP target host name
type Snmp_Information_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Host name ,udp-port , user, security model and level. The type is slice of
    // Snmp_Information_Hosts_Host_HostInformation.
    HostInformation []Snmp_Information_Hosts_Host_HostInformation
}

func (host *Snmp_Information_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + "[name='" + fmt.Sprintf("%v", host.Name) + "']"
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = make(map[string]types.YChild)
    host.EntityData.Children["host-information"] = types.YChild{"HostInformation", nil}
    for i := range host.HostInformation {
        host.EntityData.Children[types.GetSegmentPath(&host.HostInformation[i])] = types.YChild{"HostInformation", &host.HostInformation[i]}
    }
    host.EntityData.Leafs = make(map[string]types.YLeaf)
    host.EntityData.Leafs["name"] = types.YLeaf{"Name", host.Name}
    return &(host.EntityData)
}

// Snmp_Information_Hosts_Host_HostInformation
// Host name ,udp-port , user, security model
// and level
type Snmp_Information_Hosts_Host_HostInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP host user. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    hostInformation.EntityData.SegmentPath = "host-information" + "[user='" + fmt.Sprintf("%v", hostInformation.User) + "']"
    hostInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hostInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hostInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hostInformation.EntityData.Children = make(map[string]types.YChild)
    hostInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    hostInformation.EntityData.Leafs["user"] = types.YLeaf{"User", hostInformation.User}
    hostInformation.EntityData.Leafs["snmp-target-address-t-host"] = types.YLeaf{"SnmpTargetAddressTHost", hostInformation.SnmpTargetAddressTHost}
    hostInformation.EntityData.Leafs["snmp-target-address-port"] = types.YLeaf{"SnmpTargetAddressPort", hostInformation.SnmpTargetAddressPort}
    hostInformation.EntityData.Leafs["snmp-target-addresstype"] = types.YLeaf{"SnmpTargetAddresstype", hostInformation.SnmpTargetAddresstype}
    hostInformation.EntityData.Leafs["snmp-target-params-security-model"] = types.YLeaf{"SnmpTargetParamsSecurityModel", hostInformation.SnmpTargetParamsSecurityModel}
    hostInformation.EntityData.Leafs["snmp-target-params-security-name"] = types.YLeaf{"SnmpTargetParamsSecurityName", hostInformation.SnmpTargetParamsSecurityName}
    hostInformation.EntityData.Leafs["snmp-target-params-security-level"] = types.YLeaf{"SnmpTargetParamsSecurityLevel", hostInformation.SnmpTargetParamsSecurityLevel}
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
    systemUpTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemUpTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemUpTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemUpTime.EntityData.Children = make(map[string]types.YChild)
    systemUpTime.EntityData.Leafs = make(map[string]types.YLeaf)
    systemUpTime.EntityData.Leafs["system-up-time-edm"] = types.YLeaf{"SystemUpTimeEdm", systemUpTime.SystemUpTimeEdm}
    return &(systemUpTime.EntityData)
}

// Snmp_Information_NmsAddresses
// SNMP request type summary 
type Snmp_Information_NmsAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NMS address. The type is slice of Snmp_Information_NmsAddresses_NmsAddress.
    NmsAddress []Snmp_Information_NmsAddresses_NmsAddress
}

func (nmsAddresses *Snmp_Information_NmsAddresses) GetEntityData() *types.CommonEntityData {
    nmsAddresses.EntityData.YFilter = nmsAddresses.YFilter
    nmsAddresses.EntityData.YangName = "nms-addresses"
    nmsAddresses.EntityData.BundleName = "cisco_ios_xr"
    nmsAddresses.EntityData.ParentYangName = "information"
    nmsAddresses.EntityData.SegmentPath = "nms-addresses"
    nmsAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmsAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmsAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmsAddresses.EntityData.Children = make(map[string]types.YChild)
    nmsAddresses.EntityData.Children["nms-address"] = types.YChild{"NmsAddress", nil}
    for i := range nmsAddresses.NmsAddress {
        nmsAddresses.EntityData.Children[types.GetSegmentPath(&nmsAddresses.NmsAddress[i])] = types.YChild{"NmsAddress", &nmsAddresses.NmsAddress[i]}
    }
    nmsAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nmsAddresses.EntityData)
}

// Snmp_Information_NmsAddresses_NmsAddress
// NMS address
type Snmp_Information_NmsAddresses_NmsAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NMS address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    nmsAddress.EntityData.SegmentPath = "nms-address" + "[nms-addr='" + fmt.Sprintf("%v", nmsAddress.NmsAddr) + "']"
    nmsAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmsAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmsAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmsAddress.EntityData.Children = make(map[string]types.YChild)
    nmsAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    nmsAddress.EntityData.Leafs["nms-addr"] = types.YLeaf{"NmsAddr", nmsAddress.NmsAddr}
    nmsAddress.EntityData.Leafs["nms-address"] = types.YLeaf{"NmsAddress", nmsAddress.NmsAddress}
    nmsAddress.EntityData.Leafs["get-request-count"] = types.YLeaf{"GetRequestCount", nmsAddress.GetRequestCount}
    nmsAddress.EntityData.Leafs["getnext-request-count"] = types.YLeaf{"GetnextRequestCount", nmsAddress.GetnextRequestCount}
    nmsAddress.EntityData.Leafs["getbulk-request-count"] = types.YLeaf{"GetbulkRequestCount", nmsAddress.GetbulkRequestCount}
    nmsAddress.EntityData.Leafs["set-request-count"] = types.YLeaf{"SetRequestCount", nmsAddress.SetRequestCount}
    nmsAddress.EntityData.Leafs["test-request-count"] = types.YLeaf{"TestRequestCount", nmsAddress.TestRequestCount}
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
    engineId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    engineId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    engineId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    engineId.EntityData.Children = make(map[string]types.YChild)
    engineId.EntityData.Leafs = make(map[string]types.YLeaf)
    engineId.EntityData.Leafs["engine-id"] = types.YLeaf{"EngineId", engineId.EngineId}
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
    IncomingQ []Snmp_Information_RxQueue_IncomingQ

    // pending q. The type is slice of Snmp_Information_RxQueue_PendingQ.
    PendingQ []Snmp_Information_RxQueue_PendingQ
}

func (rxQueue *Snmp_Information_RxQueue) GetEntityData() *types.CommonEntityData {
    rxQueue.EntityData.YFilter = rxQueue.YFilter
    rxQueue.EntityData.YangName = "rx-queue"
    rxQueue.EntityData.BundleName = "cisco_ios_xr"
    rxQueue.EntityData.ParentYangName = "information"
    rxQueue.EntityData.SegmentPath = "rx-queue"
    rxQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxQueue.EntityData.Children = make(map[string]types.YChild)
    rxQueue.EntityData.Children["incoming-q"] = types.YChild{"IncomingQ", nil}
    for i := range rxQueue.IncomingQ {
        rxQueue.EntityData.Children[types.GetSegmentPath(&rxQueue.IncomingQ[i])] = types.YChild{"IncomingQ", &rxQueue.IncomingQ[i]}
    }
    rxQueue.EntityData.Children["pending-q"] = types.YChild{"PendingQ", nil}
    for i := range rxQueue.PendingQ {
        rxQueue.EntityData.Children[types.GetSegmentPath(&rxQueue.PendingQ[i])] = types.YChild{"PendingQ", &rxQueue.PendingQ[i]}
    }
    rxQueue.EntityData.Leafs = make(map[string]types.YLeaf)
    rxQueue.EntityData.Leafs["qlen"] = types.YLeaf{"Qlen", rxQueue.Qlen}
    rxQueue.EntityData.Leafs["in-min"] = types.YLeaf{"InMin", rxQueue.InMin}
    rxQueue.EntityData.Leafs["in-avg"] = types.YLeaf{"InAvg", rxQueue.InAvg}
    rxQueue.EntityData.Leafs["in-max"] = types.YLeaf{"InMax", rxQueue.InMax}
    rxQueue.EntityData.Leafs["pend-min"] = types.YLeaf{"PendMin", rxQueue.PendMin}
    rxQueue.EntityData.Leafs["pend-avg"] = types.YLeaf{"PendAvg", rxQueue.PendAvg}
    rxQueue.EntityData.Leafs["pend-max"] = types.YLeaf{"PendMax", rxQueue.PendMax}
    return &(rxQueue.EntityData)
}

// Snmp_Information_RxQueue_IncomingQ
// incoming q
type Snmp_Information_RxQueue_IncomingQ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    incomingQ.EntityData.SegmentPath = "incoming-q"
    incomingQ.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incomingQ.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incomingQ.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incomingQ.EntityData.Children = make(map[string]types.YChild)
    incomingQ.EntityData.Leafs = make(map[string]types.YLeaf)
    incomingQ.EntityData.Leafs["min"] = types.YLeaf{"Min", incomingQ.Min}
    incomingQ.EntityData.Leafs["avg"] = types.YLeaf{"Avg", incomingQ.Avg}
    incomingQ.EntityData.Leafs["max"] = types.YLeaf{"Max", incomingQ.Max}
    return &(incomingQ.EntityData)
}

// Snmp_Information_RxQueue_PendingQ
// pending q
type Snmp_Information_RxQueue_PendingQ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    pendingQ.EntityData.SegmentPath = "pending-q"
    pendingQ.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pendingQ.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pendingQ.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pendingQ.EntityData.Children = make(map[string]types.YChild)
    pendingQ.EntityData.Leafs = make(map[string]types.YLeaf)
    pendingQ.EntityData.Leafs["min"] = types.YLeaf{"Min", pendingQ.Min}
    pendingQ.EntityData.Leafs["avg"] = types.YLeaf{"Avg", pendingQ.Avg}
    pendingQ.EntityData.Leafs["max"] = types.YLeaf{"Max", pendingQ.Max}
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
    systemName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemName.EntityData.Children = make(map[string]types.YChild)
    systemName.EntityData.Leafs = make(map[string]types.YLeaf)
    systemName.EntityData.Leafs["system-name"] = types.YLeaf{"SystemName", systemName.SystemName}
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
    requestTypeDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requestTypeDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requestTypeDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requestTypeDetail.EntityData.Children = make(map[string]types.YChild)
    requestTypeDetail.EntityData.Children["nms-addresses"] = types.YChild{"NmsAddresses", &requestTypeDetail.NmsAddresses}
    requestTypeDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(requestTypeDetail.EntityData)
}

// Snmp_Information_RequestTypeDetail_NmsAddresses
// snmp request type details 
type Snmp_Information_RequestTypeDetail_NmsAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NMS address. The type is slice of
    // Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress.
    NmsAddress []Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress
}

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetEntityData() *types.CommonEntityData {
    nmsAddresses.EntityData.YFilter = nmsAddresses.YFilter
    nmsAddresses.EntityData.YangName = "nms-addresses"
    nmsAddresses.EntityData.BundleName = "cisco_ios_xr"
    nmsAddresses.EntityData.ParentYangName = "request-type-detail"
    nmsAddresses.EntityData.SegmentPath = "nms-addresses"
    nmsAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmsAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmsAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmsAddresses.EntityData.Children = make(map[string]types.YChild)
    nmsAddresses.EntityData.Children["nms-address"] = types.YChild{"NmsAddress", nil}
    for i := range nmsAddresses.NmsAddress {
        nmsAddresses.EntityData.Children[types.GetSegmentPath(&nmsAddresses.NmsAddress[i])] = types.YChild{"NmsAddress", &nmsAddresses.NmsAddress[i]}
    }
    nmsAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nmsAddresses.EntityData)
}

// Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress
// NMS address
type Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NMS address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    nmsAddress.EntityData.SegmentPath = "nms-address" + "[nms-addr='" + fmt.Sprintf("%v", nmsAddress.NmsAddr) + "']"
    nmsAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmsAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmsAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmsAddress.EntityData.Children = make(map[string]types.YChild)
    nmsAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    nmsAddress.EntityData.Leafs["nms-addr"] = types.YLeaf{"NmsAddr", nmsAddress.NmsAddr}
    nmsAddress.EntityData.Leafs["total-count"] = types.YLeaf{"TotalCount", nmsAddress.TotalCount}
    nmsAddress.EntityData.Leafs["agent-request-count"] = types.YLeaf{"AgentRequestCount", nmsAddress.AgentRequestCount}
    nmsAddress.EntityData.Leafs["interface-request-count"] = types.YLeaf{"InterfaceRequestCount", nmsAddress.InterfaceRequestCount}
    nmsAddress.EntityData.Leafs["entity-request-count"] = types.YLeaf{"EntityRequestCount", nmsAddress.EntityRequestCount}
    nmsAddress.EntityData.Leafs["route-request-count"] = types.YLeaf{"RouteRequestCount", nmsAddress.RouteRequestCount}
    nmsAddress.EntityData.Leafs["infra-request-count"] = types.YLeaf{"InfraRequestCount", nmsAddress.InfraRequestCount}
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
    duplicateDrop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duplicateDrop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duplicateDrop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duplicateDrop.EntityData.Children = make(map[string]types.YChild)
    duplicateDrop.EntityData.Leafs = make(map[string]types.YLeaf)
    duplicateDrop.EntityData.Leafs["duplicate-request-status"] = types.YLeaf{"DuplicateRequestStatus", duplicateDrop.DuplicateRequestStatus}
    duplicateDrop.EntityData.Leafs["last-status-change-time"] = types.YLeaf{"LastStatusChangeTime", duplicateDrop.LastStatusChangeTime}
    duplicateDrop.EntityData.Leafs["duplicate-drop-configured-timeout"] = types.YLeaf{"DuplicateDropConfiguredTimeout", duplicateDrop.DuplicateDropConfiguredTimeout}
    duplicateDrop.EntityData.Leafs["duplicate-dropped-requests"] = types.YLeaf{"DuplicateDroppedRequests", duplicateDrop.DuplicateDroppedRequests}
    duplicateDrop.EntityData.Leafs["retry-processed-requests"] = types.YLeaf{"RetryProcessedRequests", duplicateDrop.RetryProcessedRequests}
    duplicateDrop.EntityData.Leafs["first-enable-time"] = types.YLeaf{"FirstEnableTime", duplicateDrop.FirstEnableTime}
    duplicateDrop.EntityData.Leafs["latest-duplicate-dropped-requests"] = types.YLeaf{"LatestDuplicateDroppedRequests", duplicateDrop.LatestDuplicateDroppedRequests}
    duplicateDrop.EntityData.Leafs["latest-retry-processed-requests"] = types.YLeaf{"LatestRetryProcessedRequests", duplicateDrop.LatestRetryProcessedRequests}
    duplicateDrop.EntityData.Leafs["duplicate-request-latest-enable-time"] = types.YLeaf{"DuplicateRequestLatestEnableTime", duplicateDrop.DuplicateRequestLatestEnableTime}
    duplicateDrop.EntityData.Leafs["duplicate-drop-enable-count"] = types.YLeaf{"DuplicateDropEnableCount", duplicateDrop.DuplicateDropEnableCount}
    duplicateDrop.EntityData.Leafs["duplicate-drop-disable-count"] = types.YLeaf{"DuplicateDropDisableCount", duplicateDrop.DuplicateDropDisableCount}
    return &(duplicateDrop.EntityData)
}

// Snmp_Information_BulkStatsTransfers
// List of bulkstats transfer on the system
type Snmp_Information_BulkStatsTransfers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP bulkstats transfer name. The type is slice of
    // Snmp_Information_BulkStatsTransfers_BulkStatsTransfer.
    BulkStatsTransfer []Snmp_Information_BulkStatsTransfers_BulkStatsTransfer
}

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetEntityData() *types.CommonEntityData {
    bulkStatsTransfers.EntityData.YFilter = bulkStatsTransfers.YFilter
    bulkStatsTransfers.EntityData.YangName = "bulk-stats-transfers"
    bulkStatsTransfers.EntityData.BundleName = "cisco_ios_xr"
    bulkStatsTransfers.EntityData.ParentYangName = "information"
    bulkStatsTransfers.EntityData.SegmentPath = "bulk-stats-transfers"
    bulkStatsTransfers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bulkStatsTransfers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bulkStatsTransfers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bulkStatsTransfers.EntityData.Children = make(map[string]types.YChild)
    bulkStatsTransfers.EntityData.Children["bulk-stats-transfer"] = types.YChild{"BulkStatsTransfer", nil}
    for i := range bulkStatsTransfers.BulkStatsTransfer {
        bulkStatsTransfers.EntityData.Children[types.GetSegmentPath(&bulkStatsTransfers.BulkStatsTransfer[i])] = types.YChild{"BulkStatsTransfer", &bulkStatsTransfers.BulkStatsTransfer[i]}
    }
    bulkStatsTransfers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bulkStatsTransfers.EntityData)
}

// Snmp_Information_BulkStatsTransfers_BulkStatsTransfer
// SNMP bulkstats transfer name
type Snmp_Information_BulkStatsTransfers_BulkStatsTransfer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Transfer name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    bulkStatsTransfer.EntityData.SegmentPath = "bulk-stats-transfer" + "[transfer-name='" + fmt.Sprintf("%v", bulkStatsTransfer.TransferName) + "']"
    bulkStatsTransfer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bulkStatsTransfer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bulkStatsTransfer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bulkStatsTransfer.EntityData.Children = make(map[string]types.YChild)
    bulkStatsTransfer.EntityData.Leafs = make(map[string]types.YLeaf)
    bulkStatsTransfer.EntityData.Leafs["transfer-name"] = types.YLeaf{"TransferName", bulkStatsTransfer.TransferName}
    bulkStatsTransfer.EntityData.Leafs["transfer-name-xr"] = types.YLeaf{"TransferNameXr", bulkStatsTransfer.TransferNameXr}
    bulkStatsTransfer.EntityData.Leafs["url-primary"] = types.YLeaf{"UrlPrimary", bulkStatsTransfer.UrlPrimary}
    bulkStatsTransfer.EntityData.Leafs["url-secondary"] = types.YLeaf{"UrlSecondary", bulkStatsTransfer.UrlSecondary}
    bulkStatsTransfer.EntityData.Leafs["retained-file"] = types.YLeaf{"RetainedFile", bulkStatsTransfer.RetainedFile}
    bulkStatsTransfer.EntityData.Leafs["time-left"] = types.YLeaf{"TimeLeft", bulkStatsTransfer.TimeLeft}
    bulkStatsTransfer.EntityData.Leafs["retry-left"] = types.YLeaf{"RetryLeft", bulkStatsTransfer.RetryLeft}
    return &(bulkStatsTransfer.EntityData)
}

// Snmp_Information_TrapInfos
// SNMP trap OID
type Snmp_Information_TrapInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP Trap infomation like server , port and trapOID. The type is slice of
    // Snmp_Information_TrapInfos_TrapInfo.
    TrapInfo []Snmp_Information_TrapInfos_TrapInfo
}

func (trapInfos *Snmp_Information_TrapInfos) GetEntityData() *types.CommonEntityData {
    trapInfos.EntityData.YFilter = trapInfos.YFilter
    trapInfos.EntityData.YangName = "trap-infos"
    trapInfos.EntityData.BundleName = "cisco_ios_xr"
    trapInfos.EntityData.ParentYangName = "information"
    trapInfos.EntityData.SegmentPath = "trap-infos"
    trapInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapInfos.EntityData.Children = make(map[string]types.YChild)
    trapInfos.EntityData.Children["trap-info"] = types.YChild{"TrapInfo", nil}
    for i := range trapInfos.TrapInfo {
        trapInfos.EntityData.Children[types.GetSegmentPath(&trapInfos.TrapInfo[i])] = types.YChild{"TrapInfo", &trapInfos.TrapInfo[i]}
    }
    trapInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trapInfos.EntityData)
}

// Snmp_Information_TrapInfos_TrapInfo
// SNMP Trap infomation like server , port and
// trapOID
type Snmp_Information_TrapInfos_TrapInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    TrapOiDinfo []Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo
}

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetEntityData() *types.CommonEntityData {
    trapInfo.EntityData.YFilter = trapInfo.YFilter
    trapInfo.EntityData.YangName = "trap-info"
    trapInfo.EntityData.BundleName = "cisco_ios_xr"
    trapInfo.EntityData.ParentYangName = "trap-infos"
    trapInfo.EntityData.SegmentPath = "trap-info"
    trapInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapInfo.EntityData.Children = make(map[string]types.YChild)
    trapInfo.EntityData.Children["trap-oi-dinfo"] = types.YChild{"TrapOiDinfo", nil}
    for i := range trapInfo.TrapOiDinfo {
        trapInfo.EntityData.Children[types.GetSegmentPath(&trapInfo.TrapOiDinfo[i])] = types.YChild{"TrapOiDinfo", &trapInfo.TrapOiDinfo[i]}
    }
    trapInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trapInfo.EntityData.Leafs["trap-host"] = types.YLeaf{"TrapHost", trapInfo.TrapHost}
    trapInfo.EntityData.Leafs["port"] = types.YLeaf{"Port", trapInfo.Port}
    trapInfo.EntityData.Leafs["host"] = types.YLeaf{"Host", trapInfo.Host}
    trapInfo.EntityData.Leafs["port-xr"] = types.YLeaf{"PortXr", trapInfo.PortXr}
    trapInfo.EntityData.Leafs["trap-oid-count"] = types.YLeaf{"TrapOidCount", trapInfo.TrapOidCount}
    return &(trapInfo.EntityData)
}

// Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo
// Per trap OID statistics
type Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    trapOiDinfo.EntityData.SegmentPath = "trap-oi-dinfo"
    trapOiDinfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapOiDinfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapOiDinfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapOiDinfo.EntityData.Children = make(map[string]types.YChild)
    trapOiDinfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trapOiDinfo.EntityData.Leafs["trap-oid"] = types.YLeaf{"TrapOid", trapOiDinfo.TrapOid}
    trapOiDinfo.EntityData.Leafs["count"] = types.YLeaf{"Count", trapOiDinfo.Count}
    trapOiDinfo.EntityData.Leafs["drop-count"] = types.YLeaf{"DropCount", trapOiDinfo.DropCount}
    trapOiDinfo.EntityData.Leafs["retry-count"] = types.YLeaf{"RetryCount", trapOiDinfo.RetryCount}
    trapOiDinfo.EntityData.Leafs["lastsent-time"] = types.YLeaf{"LastsentTime", trapOiDinfo.LastsentTime}
    trapOiDinfo.EntityData.Leafs["lasrdrop-time"] = types.YLeaf{"LasrdropTime", trapOiDinfo.LasrdropTime}
    return &(trapOiDinfo.EntityData)
}

// Snmp_Information_PollOids
// OID list for poll PDU
type Snmp_Information_PollOids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PDU drop info for OID. The type is slice of
    // Snmp_Information_PollOids_PollOid.
    PollOid []Snmp_Information_PollOids_PollOid
}

func (pollOids *Snmp_Information_PollOids) GetEntityData() *types.CommonEntityData {
    pollOids.EntityData.YFilter = pollOids.YFilter
    pollOids.EntityData.YangName = "poll-oids"
    pollOids.EntityData.BundleName = "cisco_ios_xr"
    pollOids.EntityData.ParentYangName = "information"
    pollOids.EntityData.SegmentPath = "poll-oids"
    pollOids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pollOids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pollOids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pollOids.EntityData.Children = make(map[string]types.YChild)
    pollOids.EntityData.Children["poll-oid"] = types.YChild{"PollOid", nil}
    for i := range pollOids.PollOid {
        pollOids.EntityData.Children[types.GetSegmentPath(&pollOids.PollOid[i])] = types.YChild{"PollOid", &pollOids.PollOid[i]}
    }
    pollOids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pollOids.EntityData)
}

// Snmp_Information_PollOids_PollOid
// PDU drop info for OID
type Snmp_Information_PollOids_PollOid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Object ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    pollOid.EntityData.SegmentPath = "poll-oid" + "[object-id='" + fmt.Sprintf("%v", pollOid.ObjectId) + "']"
    pollOid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pollOid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pollOid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pollOid.EntityData.Children = make(map[string]types.YChild)
    pollOid.EntityData.Leafs = make(map[string]types.YLeaf)
    pollOid.EntityData.Leafs["object-id"] = types.YLeaf{"ObjectId", pollOid.ObjectId}
    pollOid.EntityData.Leafs["nms-count"] = types.YLeaf{"NmsCount", pollOid.NmsCount}
    pollOid.EntityData.Leafs["nms"] = types.YLeaf{"Nms", pollOid.Nms}
    pollOid.EntityData.Leafs["request-count"] = types.YLeaf{"RequestCount", pollOid.RequestCount}
    return &(pollOid.EntityData)
}

// Snmp_Information_InfomDetails
// SNMP trap OID
type Snmp_Information_InfomDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP Trap infomation like server , port and trapOID. The type is slice of
    // Snmp_Information_InfomDetails_InfomDetail.
    InfomDetail []Snmp_Information_InfomDetails_InfomDetail
}

func (infomDetails *Snmp_Information_InfomDetails) GetEntityData() *types.CommonEntityData {
    infomDetails.EntityData.YFilter = infomDetails.YFilter
    infomDetails.EntityData.YangName = "infom-details"
    infomDetails.EntityData.BundleName = "cisco_ios_xr"
    infomDetails.EntityData.ParentYangName = "information"
    infomDetails.EntityData.SegmentPath = "infom-details"
    infomDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infomDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infomDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infomDetails.EntityData.Children = make(map[string]types.YChild)
    infomDetails.EntityData.Children["infom-detail"] = types.YChild{"InfomDetail", nil}
    for i := range infomDetails.InfomDetail {
        infomDetails.EntityData.Children[types.GetSegmentPath(&infomDetails.InfomDetail[i])] = types.YChild{"InfomDetail", &infomDetails.InfomDetail[i]}
    }
    infomDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(infomDetails.EntityData)
}

// Snmp_Information_InfomDetails_InfomDetail
// SNMP Trap infomation like server , port and
// trapOID
type Snmp_Information_InfomDetails_InfomDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    TrapOiDinfo []Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo
}

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetEntityData() *types.CommonEntityData {
    infomDetail.EntityData.YFilter = infomDetail.YFilter
    infomDetail.EntityData.YangName = "infom-detail"
    infomDetail.EntityData.BundleName = "cisco_ios_xr"
    infomDetail.EntityData.ParentYangName = "infom-details"
    infomDetail.EntityData.SegmentPath = "infom-detail"
    infomDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    infomDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    infomDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    infomDetail.EntityData.Children = make(map[string]types.YChild)
    infomDetail.EntityData.Children["trap-oi-dinfo"] = types.YChild{"TrapOiDinfo", nil}
    for i := range infomDetail.TrapOiDinfo {
        infomDetail.EntityData.Children[types.GetSegmentPath(&infomDetail.TrapOiDinfo[i])] = types.YChild{"TrapOiDinfo", &infomDetail.TrapOiDinfo[i]}
    }
    infomDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    infomDetail.EntityData.Leafs["trap-host"] = types.YLeaf{"TrapHost", infomDetail.TrapHost}
    infomDetail.EntityData.Leafs["port"] = types.YLeaf{"Port", infomDetail.Port}
    infomDetail.EntityData.Leafs["host"] = types.YLeaf{"Host", infomDetail.Host}
    infomDetail.EntityData.Leafs["port-xr"] = types.YLeaf{"PortXr", infomDetail.PortXr}
    infomDetail.EntityData.Leafs["trap-oid-count"] = types.YLeaf{"TrapOidCount", infomDetail.TrapOidCount}
    return &(infomDetail.EntityData)
}

// Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo
// Per trap OID statistics
type Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    trapOiDinfo.EntityData.SegmentPath = "trap-oi-dinfo"
    trapOiDinfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapOiDinfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapOiDinfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapOiDinfo.EntityData.Children = make(map[string]types.YChild)
    trapOiDinfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trapOiDinfo.EntityData.Leafs["trap-oid"] = types.YLeaf{"TrapOid", trapOiDinfo.TrapOid}
    trapOiDinfo.EntityData.Leafs["count"] = types.YLeaf{"Count", trapOiDinfo.Count}
    trapOiDinfo.EntityData.Leafs["drop-count"] = types.YLeaf{"DropCount", trapOiDinfo.DropCount}
    trapOiDinfo.EntityData.Leafs["retry-count"] = types.YLeaf{"RetryCount", trapOiDinfo.RetryCount}
    trapOiDinfo.EntityData.Leafs["lastsent-time"] = types.YLeaf{"LastsentTime", trapOiDinfo.LastsentTime}
    trapOiDinfo.EntityData.Leafs["lasrdrop-time"] = types.YLeaf{"LasrdropTime", trapOiDinfo.LasrdropTime}
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
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["packets-received"] = types.YLeaf{"PacketsReceived", statistics.PacketsReceived}
    statistics.EntityData.Leafs["bad-versions-received"] = types.YLeaf{"BadVersionsReceived", statistics.BadVersionsReceived}
    statistics.EntityData.Leafs["bad-community-names-received"] = types.YLeaf{"BadCommunityNamesReceived", statistics.BadCommunityNamesReceived}
    statistics.EntityData.Leafs["bad-community-uses-received"] = types.YLeaf{"BadCommunityUsesReceived", statistics.BadCommunityUsesReceived}
    statistics.EntityData.Leafs["asn-parse-errors-received"] = types.YLeaf{"AsnParseErrorsReceived", statistics.AsnParseErrorsReceived}
    statistics.EntityData.Leafs["silent-drop-count"] = types.YLeaf{"SilentDropCount", statistics.SilentDropCount}
    statistics.EntityData.Leafs["proxy-drop-count"] = types.YLeaf{"ProxyDropCount", statistics.ProxyDropCount}
    statistics.EntityData.Leafs["too-big-packet-received"] = types.YLeaf{"TooBigPacketReceived", statistics.TooBigPacketReceived}
    statistics.EntityData.Leafs["max-packet-size"] = types.YLeaf{"MaxPacketSize", statistics.MaxPacketSize}
    statistics.EntityData.Leafs["no-such-names-received"] = types.YLeaf{"NoSuchNamesReceived", statistics.NoSuchNamesReceived}
    statistics.EntityData.Leafs["bad-values-received"] = types.YLeaf{"BadValuesReceived", statistics.BadValuesReceived}
    statistics.EntityData.Leafs["read-only-received"] = types.YLeaf{"ReadOnlyReceived", statistics.ReadOnlyReceived}
    statistics.EntityData.Leafs["total-general-errors"] = types.YLeaf{"TotalGeneralErrors", statistics.TotalGeneralErrors}
    statistics.EntityData.Leafs["total-requested-variables"] = types.YLeaf{"TotalRequestedVariables", statistics.TotalRequestedVariables}
    statistics.EntityData.Leafs["total-set-variables-received"] = types.YLeaf{"TotalSetVariablesReceived", statistics.TotalSetVariablesReceived}
    statistics.EntityData.Leafs["get-requests-received"] = types.YLeaf{"GetRequestsReceived", statistics.GetRequestsReceived}
    statistics.EntityData.Leafs["get-next-requests-received"] = types.YLeaf{"GetNextRequestsReceived", statistics.GetNextRequestsReceived}
    statistics.EntityData.Leafs["set-requests-received"] = types.YLeaf{"SetRequestsReceived", statistics.SetRequestsReceived}
    statistics.EntityData.Leafs["get-responses-received"] = types.YLeaf{"GetResponsesReceived", statistics.GetResponsesReceived}
    statistics.EntityData.Leafs["traps-received"] = types.YLeaf{"TrapsReceived", statistics.TrapsReceived}
    statistics.EntityData.Leafs["total-packets-sent"] = types.YLeaf{"TotalPacketsSent", statistics.TotalPacketsSent}
    statistics.EntityData.Leafs["too-big-packets-sent"] = types.YLeaf{"TooBigPacketsSent", statistics.TooBigPacketsSent}
    statistics.EntityData.Leafs["no-such-names-sent"] = types.YLeaf{"NoSuchNamesSent", statistics.NoSuchNamesSent}
    statistics.EntityData.Leafs["bad-values-sent"] = types.YLeaf{"BadValuesSent", statistics.BadValuesSent}
    statistics.EntityData.Leafs["general-errors-sent"] = types.YLeaf{"GeneralErrorsSent", statistics.GeneralErrorsSent}
    statistics.EntityData.Leafs["get-requests-sent"] = types.YLeaf{"GetRequestsSent", statistics.GetRequestsSent}
    statistics.EntityData.Leafs["get-next-request-sent"] = types.YLeaf{"GetNextRequestSent", statistics.GetNextRequestSent}
    statistics.EntityData.Leafs["set-requests-sent"] = types.YLeaf{"SetRequestsSent", statistics.SetRequestsSent}
    statistics.EntityData.Leafs["get-responses-sent"] = types.YLeaf{"GetResponsesSent", statistics.GetResponsesSent}
    statistics.EntityData.Leafs["traps-sent"] = types.YLeaf{"TrapsSent", statistics.TrapsSent}
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
    InqEntry []Snmp_Information_IncomingQueue_InqEntry
}

func (incomingQueue *Snmp_Information_IncomingQueue) GetEntityData() *types.CommonEntityData {
    incomingQueue.EntityData.YFilter = incomingQueue.YFilter
    incomingQueue.EntityData.YangName = "incoming-queue"
    incomingQueue.EntityData.BundleName = "cisco_ios_xr"
    incomingQueue.EntityData.ParentYangName = "information"
    incomingQueue.EntityData.SegmentPath = "incoming-queue"
    incomingQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incomingQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incomingQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incomingQueue.EntityData.Children = make(map[string]types.YChild)
    incomingQueue.EntityData.Children["inq-entry"] = types.YChild{"InqEntry", nil}
    for i := range incomingQueue.InqEntry {
        incomingQueue.EntityData.Children[types.GetSegmentPath(&incomingQueue.InqEntry[i])] = types.YChild{"InqEntry", &incomingQueue.InqEntry[i]}
    }
    incomingQueue.EntityData.Leafs = make(map[string]types.YLeaf)
    incomingQueue.EntityData.Leafs["queue-count"] = types.YLeaf{"QueueCount", incomingQueue.QueueCount}
    return &(incomingQueue.EntityData)
}

// Snmp_Information_IncomingQueue_InqEntry
// Each Entry Details.
type Snmp_Information_IncomingQueue_InqEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    inqEntry.EntityData.SegmentPath = "inq-entry"
    inqEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inqEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inqEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inqEntry.EntityData.Children = make(map[string]types.YChild)
    inqEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    inqEntry.EntityData.Leafs["address-of-queue"] = types.YLeaf{"AddressOfQueue", inqEntry.AddressOfQueue}
    inqEntry.EntityData.Leafs["request-count"] = types.YLeaf{"RequestCount", inqEntry.RequestCount}
    inqEntry.EntityData.Leafs["processed-request-count"] = types.YLeaf{"ProcessedRequestCount", inqEntry.ProcessedRequestCount}
    inqEntry.EntityData.Leafs["last-access-time"] = types.YLeaf{"LastAccessTime", inqEntry.LastAccessTime}
    inqEntry.EntityData.Leafs["priority"] = types.YLeaf{"Priority", inqEntry.Priority}
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
    ContexMapping []Snmp_Information_ContextMapping_ContexMapping
}

func (contextMapping *Snmp_Information_ContextMapping) GetEntityData() *types.CommonEntityData {
    contextMapping.EntityData.YFilter = contextMapping.YFilter
    contextMapping.EntityData.YangName = "context-mapping"
    contextMapping.EntityData.BundleName = "cisco_ios_xr"
    contextMapping.EntityData.ParentYangName = "information"
    contextMapping.EntityData.SegmentPath = "context-mapping"
    contextMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextMapping.EntityData.Children = make(map[string]types.YChild)
    contextMapping.EntityData.Children["contex-mapping"] = types.YChild{"ContexMapping", nil}
    for i := range contextMapping.ContexMapping {
        contextMapping.EntityData.Children[types.GetSegmentPath(&contextMapping.ContexMapping[i])] = types.YChild{"ContexMapping", &contextMapping.ContexMapping[i]}
    }
    contextMapping.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(contextMapping.EntityData)
}

// Snmp_Information_ContextMapping_ContexMapping
// Context Mapping
type Snmp_Information_ContextMapping_ContexMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    contexMapping.EntityData.SegmentPath = "contex-mapping"
    contexMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contexMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contexMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contexMapping.EntityData.Children = make(map[string]types.YChild)
    contexMapping.EntityData.Leafs = make(map[string]types.YLeaf)
    contexMapping.EntityData.Leafs["context"] = types.YLeaf{"Context", contexMapping.Context}
    contexMapping.EntityData.Leafs["feature-name"] = types.YLeaf{"FeatureName", contexMapping.FeatureName}
    contexMapping.EntityData.Leafs["instance"] = types.YLeaf{"Instance", contexMapping.Instance}
    contexMapping.EntityData.Leafs["topology"] = types.YLeaf{"Topology", contexMapping.Topology}
    contexMapping.EntityData.Leafs["feature"] = types.YLeaf{"Feature", contexMapping.Feature}
    return &(contexMapping.EntityData)
}

// Snmp_Information_TrapOids
// SNMP trap OID
type Snmp_Information_TrapOids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP trap . The type is slice of Snmp_Information_TrapOids_TrapOid.
    TrapOid []Snmp_Information_TrapOids_TrapOid
}

func (trapOids *Snmp_Information_TrapOids) GetEntityData() *types.CommonEntityData {
    trapOids.EntityData.YFilter = trapOids.YFilter
    trapOids.EntityData.YangName = "trap-oids"
    trapOids.EntityData.BundleName = "cisco_ios_xr"
    trapOids.EntityData.ParentYangName = "information"
    trapOids.EntityData.SegmentPath = "trap-oids"
    trapOids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapOids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapOids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapOids.EntityData.Children = make(map[string]types.YChild)
    trapOids.EntityData.Children["trap-oid"] = types.YChild{"TrapOid", nil}
    for i := range trapOids.TrapOid {
        trapOids.EntityData.Children[types.GetSegmentPath(&trapOids.TrapOid[i])] = types.YChild{"TrapOid", &trapOids.TrapOid[i]}
    }
    trapOids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trapOids.EntityData)
}

// Snmp_Information_TrapOids_TrapOid
// SNMP trap 
type Snmp_Information_TrapOids_TrapOid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Trap object ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    trapOid.EntityData.SegmentPath = "trap-oid" + "[trap-oid='" + fmt.Sprintf("%v", trapOid.TrapOid) + "']"
    trapOid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapOid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapOid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapOid.EntityData.Children = make(map[string]types.YChild)
    trapOid.EntityData.Leafs = make(map[string]types.YLeaf)
    trapOid.EntityData.Leafs["trap-oid"] = types.YLeaf{"TrapOid", trapOid.TrapOid}
    trapOid.EntityData.Leafs["trap-oid-count"] = types.YLeaf{"TrapOidCount", trapOid.TrapOidCount}
    trapOid.EntityData.Leafs["trap-oid-xr"] = types.YLeaf{"TrapOidXr", trapOid.TrapOidXr}
    return &(trapOid.EntityData)
}

// Snmp_Information_NmSpackets
// SNMP overload statistics 
type Snmp_Information_NmSpackets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NMS packet drop count. The type is slice of
    // Snmp_Information_NmSpackets_NmSpacket.
    NmSpacket []Snmp_Information_NmSpackets_NmSpacket
}

func (nmSpackets *Snmp_Information_NmSpackets) GetEntityData() *types.CommonEntityData {
    nmSpackets.EntityData.YFilter = nmSpackets.YFilter
    nmSpackets.EntityData.YangName = "nm-spackets"
    nmSpackets.EntityData.BundleName = "cisco_ios_xr"
    nmSpackets.EntityData.ParentYangName = "information"
    nmSpackets.EntityData.SegmentPath = "nm-spackets"
    nmSpackets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmSpackets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmSpackets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmSpackets.EntityData.Children = make(map[string]types.YChild)
    nmSpackets.EntityData.Children["nm-spacket"] = types.YChild{"NmSpacket", nil}
    for i := range nmSpackets.NmSpacket {
        nmSpackets.EntityData.Children[types.GetSegmentPath(&nmSpackets.NmSpacket[i])] = types.YChild{"NmSpacket", &nmSpackets.NmSpacket[i]}
    }
    nmSpackets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nmSpackets.EntityData)
}

// Snmp_Information_NmSpackets_NmSpacket
// NMS packet drop count
type Snmp_Information_NmSpackets_NmSpacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NMS packet drop count. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    nmSpacket.EntityData.SegmentPath = "nm-spacket" + "[packetcount='" + fmt.Sprintf("%v", nmSpacket.Packetcount) + "']"
    nmSpacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nmSpacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nmSpacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nmSpacket.EntityData.Children = make(map[string]types.YChild)
    nmSpacket.EntityData.Leafs = make(map[string]types.YLeaf)
    nmSpacket.EntityData.Leafs["packetcount"] = types.YLeaf{"Packetcount", nmSpacket.Packetcount}
    nmSpacket.EntityData.Leafs["number-of-nmsq-pkts-dropped"] = types.YLeaf{"NumberOfNmsqPktsDropped", nmSpacket.NumberOfNmsqPktsDropped}
    nmSpacket.EntityData.Leafs["number-of-pkts-dropped"] = types.YLeaf{"NumberOfPktsDropped", nmSpacket.NumberOfPktsDropped}
    nmSpacket.EntityData.Leafs["overload-start-time"] = types.YLeaf{"OverloadStartTime", nmSpacket.OverloadStartTime}
    nmSpacket.EntityData.Leafs["overload-end-time"] = types.YLeaf{"OverloadEndTime", nmSpacket.OverloadEndTime}
    return &(nmSpacket.EntityData)
}

// Snmp_Information_Mibs
// List of MIBS supported on the system
type Snmp_Information_Mibs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP MIB Name. The type is slice of Snmp_Information_Mibs_Mib.
    Mib []Snmp_Information_Mibs_Mib
}

func (mibs *Snmp_Information_Mibs) GetEntityData() *types.CommonEntityData {
    mibs.EntityData.YFilter = mibs.YFilter
    mibs.EntityData.YangName = "mibs"
    mibs.EntityData.BundleName = "cisco_ios_xr"
    mibs.EntityData.ParentYangName = "information"
    mibs.EntityData.SegmentPath = "mibs"
    mibs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mibs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mibs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mibs.EntityData.Children = make(map[string]types.YChild)
    mibs.EntityData.Children["mib"] = types.YChild{"Mib", nil}
    for i := range mibs.Mib {
        mibs.EntityData.Children[types.GetSegmentPath(&mibs.Mib[i])] = types.YChild{"Mib", &mibs.Mib[i]}
    }
    mibs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mibs.EntityData)
}

// Snmp_Information_Mibs_Mib
// SNMP MIB Name
type Snmp_Information_Mibs_Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MIB Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    mib.EntityData.SegmentPath = "mib" + "[name='" + fmt.Sprintf("%v", mib.Name) + "']"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = make(map[string]types.YChild)
    mib.EntityData.Children["oids"] = types.YChild{"Oids", &mib.Oids}
    mib.EntityData.Children["mib-information"] = types.YChild{"MibInformation", &mib.MibInformation}
    mib.EntityData.Leafs = make(map[string]types.YLeaf)
    mib.EntityData.Leafs["name"] = types.YLeaf{"Name", mib.Name}
    return &(mib.EntityData)
}

// Snmp_Information_Mibs_Mib_Oids
// List of OIDs per MIB
type Snmp_Information_Mibs_Mib_Oids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object identifiers of a mib. The type is slice of
    // Snmp_Information_Mibs_Mib_Oids_Oid.
    Oid []Snmp_Information_Mibs_Mib_Oids_Oid
}

func (oids *Snmp_Information_Mibs_Mib_Oids) GetEntityData() *types.CommonEntityData {
    oids.EntityData.YFilter = oids.YFilter
    oids.EntityData.YangName = "oids"
    oids.EntityData.BundleName = "cisco_ios_xr"
    oids.EntityData.ParentYangName = "mib"
    oids.EntityData.SegmentPath = "oids"
    oids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oids.EntityData.Children = make(map[string]types.YChild)
    oids.EntityData.Children["oid"] = types.YChild{"Oid", nil}
    for i := range oids.Oid {
        oids.EntityData.Children[types.GetSegmentPath(&oids.Oid[i])] = types.YChild{"Oid", &oids.Oid[i]}
    }
    oids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oids.EntityData)
}

// Snmp_Information_Mibs_Mib_Oids_Oid
// Object identifiers of a mib
type Snmp_Information_Mibs_Mib_Oids_Oid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    oid.EntityData.SegmentPath = "oid" + "[oid='" + fmt.Sprintf("%v", oid.Oid) + "']"
    oid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oid.EntityData.Children = make(map[string]types.YChild)
    oid.EntityData.Leafs = make(map[string]types.YLeaf)
    oid.EntityData.Leafs["oid"] = types.YLeaf{"Oid", oid.Oid}
    oid.EntityData.Leafs["oid-name"] = types.YLeaf{"OidName", oid.OidName}
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
    mibInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mibInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mibInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mibInformation.EntityData.Children = make(map[string]types.YChild)
    mibInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    mibInformation.EntityData.Leafs["mib-name"] = types.YLeaf{"MibName", mibInformation.MibName}
    mibInformation.EntityData.Leafs["dll-name"] = types.YLeaf{"DllName", mibInformation.DllName}
    mibInformation.EntityData.Leafs["mib-config-filename"] = types.YLeaf{"MibConfigFilename", mibInformation.MibConfigFilename}
    mibInformation.EntityData.Leafs["is-mib-loaded"] = types.YLeaf{"IsMibLoaded", mibInformation.IsMibLoaded}
    mibInformation.EntityData.Leafs["dll-capabilities"] = types.YLeaf{"DllCapabilities", mibInformation.DllCapabilities}
    mibInformation.EntityData.Leafs["trap-strings"] = types.YLeaf{"TrapStrings", mibInformation.TrapStrings}
    mibInformation.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", mibInformation.Timeout}
    mibInformation.EntityData.Leafs["load-time"] = types.YLeaf{"LoadTime", mibInformation.LoadTime}
    return &(mibInformation.EntityData)
}

// Snmp_Information_SerialNumbers
// SNMP statistics pdu 
type Snmp_Information_SerialNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Serial number. The type is slice of
    // Snmp_Information_SerialNumbers_SerialNumber.
    SerialNumber []Snmp_Information_SerialNumbers_SerialNumber
}

func (serialNumbers *Snmp_Information_SerialNumbers) GetEntityData() *types.CommonEntityData {
    serialNumbers.EntityData.YFilter = serialNumbers.YFilter
    serialNumbers.EntityData.YangName = "serial-numbers"
    serialNumbers.EntityData.BundleName = "cisco_ios_xr"
    serialNumbers.EntityData.ParentYangName = "information"
    serialNumbers.EntityData.SegmentPath = "serial-numbers"
    serialNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serialNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serialNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serialNumbers.EntityData.Children = make(map[string]types.YChild)
    serialNumbers.EntityData.Children["serial-number"] = types.YChild{"SerialNumber", nil}
    for i := range serialNumbers.SerialNumber {
        serialNumbers.EntityData.Children[types.GetSegmentPath(&serialNumbers.SerialNumber[i])] = types.YChild{"SerialNumber", &serialNumbers.SerialNumber[i]}
    }
    serialNumbers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serialNumbers.EntityData)
}

// Snmp_Information_SerialNumbers_SerialNumber
// Serial number
type Snmp_Information_SerialNumbers_SerialNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Serial number. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Number interface{}

    // Request ID. The type is interface{} with range: -2147483648..2147483647.
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
    serialNumber.EntityData.SegmentPath = "serial-number"
    serialNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serialNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serialNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serialNumber.EntityData.Children = make(map[string]types.YChild)
    serialNumber.EntityData.Leafs = make(map[string]types.YLeaf)
    serialNumber.EntityData.Leafs["number"] = types.YLeaf{"Number", serialNumber.Number}
    serialNumber.EntityData.Leafs["req-id"] = types.YLeaf{"ReqId", serialNumber.ReqId}
    serialNumber.EntityData.Leafs["port"] = types.YLeaf{"Port", serialNumber.Port}
    serialNumber.EntityData.Leafs["nms"] = types.YLeaf{"Nms", serialNumber.Nms}
    serialNumber.EntityData.Leafs["request-id"] = types.YLeaf{"RequestId", serialNumber.RequestId}
    serialNumber.EntityData.Leafs["port-xr"] = types.YLeaf{"PortXr", serialNumber.PortXr}
    serialNumber.EntityData.Leafs["pdu-type"] = types.YLeaf{"PduType", serialNumber.PduType}
    serialNumber.EntityData.Leafs["error-status"] = types.YLeaf{"ErrorStatus", serialNumber.ErrorStatus}
    serialNumber.EntityData.Leafs["serial-num"] = types.YLeaf{"SerialNum", serialNumber.SerialNum}
    serialNumber.EntityData.Leafs["input-q"] = types.YLeaf{"InputQ", serialNumber.InputQ}
    serialNumber.EntityData.Leafs["output-q"] = types.YLeaf{"OutputQ", serialNumber.OutputQ}
    serialNumber.EntityData.Leafs["pending-q"] = types.YLeaf{"PendingQ", serialNumber.PendingQ}
    serialNumber.EntityData.Leafs["response-out"] = types.YLeaf{"ResponseOut", serialNumber.ResponseOut}
    return &(serialNumber.EntityData)
}

// Snmp_Information_DropNmsAddresses
// NMS list for drop PDU
type Snmp_Information_DropNmsAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PDU drop info for NMS. The type is slice of
    // Snmp_Information_DropNmsAddresses_DropNmsAddress.
    DropNmsAddress []Snmp_Information_DropNmsAddresses_DropNmsAddress
}

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetEntityData() *types.CommonEntityData {
    dropNmsAddresses.EntityData.YFilter = dropNmsAddresses.YFilter
    dropNmsAddresses.EntityData.YangName = "drop-nms-addresses"
    dropNmsAddresses.EntityData.BundleName = "cisco_ios_xr"
    dropNmsAddresses.EntityData.ParentYangName = "information"
    dropNmsAddresses.EntityData.SegmentPath = "drop-nms-addresses"
    dropNmsAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropNmsAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropNmsAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropNmsAddresses.EntityData.Children = make(map[string]types.YChild)
    dropNmsAddresses.EntityData.Children["drop-nms-address"] = types.YChild{"DropNmsAddress", nil}
    for i := range dropNmsAddresses.DropNmsAddress {
        dropNmsAddresses.EntityData.Children[types.GetSegmentPath(&dropNmsAddresses.DropNmsAddress[i])] = types.YChild{"DropNmsAddress", &dropNmsAddresses.DropNmsAddress[i]}
    }
    dropNmsAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dropNmsAddresses.EntityData)
}

// Snmp_Information_DropNmsAddresses_DropNmsAddress
// PDU drop info for NMS
type Snmp_Information_DropNmsAddresses_DropNmsAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. NMS address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    dropNmsAddress.EntityData.SegmentPath = "drop-nms-address" + "[nms-addr='" + fmt.Sprintf("%v", dropNmsAddress.NmsAddr) + "']"
    dropNmsAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropNmsAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropNmsAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropNmsAddress.EntityData.Children = make(map[string]types.YChild)
    dropNmsAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    dropNmsAddress.EntityData.Leafs["nms-addr"] = types.YLeaf{"NmsAddr", dropNmsAddress.NmsAddr}
    dropNmsAddress.EntityData.Leafs["nms-address"] = types.YLeaf{"NmsAddress", dropNmsAddress.NmsAddress}
    dropNmsAddress.EntityData.Leafs["incoming-q-count"] = types.YLeaf{"IncomingQCount", dropNmsAddress.IncomingQCount}
    dropNmsAddress.EntityData.Leafs["threshold-incoming-q-count"] = types.YLeaf{"ThresholdIncomingQCount", dropNmsAddress.ThresholdIncomingQCount}
    dropNmsAddress.EntityData.Leafs["encode-count"] = types.YLeaf{"EncodeCount", dropNmsAddress.EncodeCount}
    dropNmsAddress.EntityData.Leafs["duplicate-count"] = types.YLeaf{"DuplicateCount", dropNmsAddress.DuplicateCount}
    dropNmsAddress.EntityData.Leafs["stack-count"] = types.YLeaf{"StackCount", dropNmsAddress.StackCount}
    dropNmsAddress.EntityData.Leafs["aipc-count"] = types.YLeaf{"AipcCount", dropNmsAddress.AipcCount}
    dropNmsAddress.EntityData.Leafs["overload-count"] = types.YLeaf{"OverloadCount", dropNmsAddress.OverloadCount}
    dropNmsAddress.EntityData.Leafs["timeout-count"] = types.YLeaf{"TimeoutCount", dropNmsAddress.TimeoutCount}
    dropNmsAddress.EntityData.Leafs["internal-count"] = types.YLeaf{"InternalCount", dropNmsAddress.InternalCount}
    return &(dropNmsAddress.EntityData)
}

// Snmp_Information_Views
// SNMP view information
type Snmp_Information_Views struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP target view name. The type is slice of Snmp_Information_Views_View.
    View []Snmp_Information_Views_View
}

func (views *Snmp_Information_Views) GetEntityData() *types.CommonEntityData {
    views.EntityData.YFilter = views.YFilter
    views.EntityData.YangName = "views"
    views.EntityData.BundleName = "cisco_ios_xr"
    views.EntityData.ParentYangName = "information"
    views.EntityData.SegmentPath = "views"
    views.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    views.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    views.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    views.EntityData.Children = make(map[string]types.YChild)
    views.EntityData.Children["view"] = types.YChild{"View", nil}
    for i := range views.View {
        views.EntityData.Children[types.GetSegmentPath(&views.View[i])] = types.YChild{"View", &views.View[i]}
    }
    views.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(views.EntityData)
}

// Snmp_Information_Views_View
// SNMP target view name
type Snmp_Information_Views_View struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. View name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // View name ,familytype, storagetype and status. The type is slice of
    // Snmp_Information_Views_View_ViewInformation.
    ViewInformation []Snmp_Information_Views_View_ViewInformation
}

func (view *Snmp_Information_Views_View) GetEntityData() *types.CommonEntityData {
    view.EntityData.YFilter = view.YFilter
    view.EntityData.YangName = "view"
    view.EntityData.BundleName = "cisco_ios_xr"
    view.EntityData.ParentYangName = "views"
    view.EntityData.SegmentPath = "view" + "[name='" + fmt.Sprintf("%v", view.Name) + "']"
    view.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    view.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    view.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    view.EntityData.Children = make(map[string]types.YChild)
    view.EntityData.Children["view-information"] = types.YChild{"ViewInformation", nil}
    for i := range view.ViewInformation {
        view.EntityData.Children[types.GetSegmentPath(&view.ViewInformation[i])] = types.YChild{"ViewInformation", &view.ViewInformation[i]}
    }
    view.EntityData.Leafs = make(map[string]types.YLeaf)
    view.EntityData.Leafs["name"] = types.YLeaf{"Name", view.Name}
    return &(view.EntityData)
}

// Snmp_Information_Views_View_ViewInformation
// View name ,familytype, storagetype and status
type Snmp_Information_Views_View_ViewInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP view OID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    viewInformation.EntityData.SegmentPath = "view-information" + "[object-id='" + fmt.Sprintf("%v", viewInformation.ObjectId) + "']"
    viewInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    viewInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    viewInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    viewInformation.EntityData.Children = make(map[string]types.YChild)
    viewInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    viewInformation.EntityData.Leafs["object-id"] = types.YLeaf{"ObjectId", viewInformation.ObjectId}
    viewInformation.EntityData.Leafs["snmp-view-family-type"] = types.YLeaf{"SnmpViewFamilyType", viewInformation.SnmpViewFamilyType}
    viewInformation.EntityData.Leafs["snmp-view-family-storage-type"] = types.YLeaf{"SnmpViewFamilyStorageType", viewInformation.SnmpViewFamilyStorageType}
    viewInformation.EntityData.Leafs["snmp-view-family-status"] = types.YLeaf{"SnmpViewFamilyStatus", viewInformation.SnmpViewFamilyStatus}
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
    systemDescr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemDescr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemDescr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemDescr.EntityData.Children = make(map[string]types.YChild)
    systemDescr.EntityData.Leafs = make(map[string]types.YLeaf)
    systemDescr.EntityData.Leafs["sys-descr"] = types.YLeaf{"SysDescr", systemDescr.SysDescr}
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
    tables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tables.EntityData.Children = make(map[string]types.YChild)
    tables.EntityData.Children["groups"] = types.YChild{"Groups", &tables.Groups}
    tables.EntityData.Children["user-engine-ids"] = types.YChild{"UserEngineIds", &tables.UserEngineIds}
    tables.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tables.EntityData)
}

// Snmp_Information_Tables_Groups
// List of vacmAccessTable
type Snmp_Information_Tables_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP group name. The type is slice of Snmp_Information_Tables_Groups_Group.
    Group []Snmp_Information_Tables_Groups_Group
}

func (groups *Snmp_Information_Tables_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "tables"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// Snmp_Information_Tables_Groups_Group
// SNMP group name
type Snmp_Information_Tables_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // Group Model.
    GroupInformations Snmp_Information_Tables_Groups_Group_GroupInformations
}

func (group *Snmp_Information_Tables_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[name='" + fmt.Sprintf("%v", group.Name) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Children["group-informations"] = types.YChild{"GroupInformations", &group.GroupInformations}
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["name"] = types.YLeaf{"Name", group.Name}
    return &(group.EntityData)
}

// Snmp_Information_Tables_Groups_Group_GroupInformations
// Group Model
type Snmp_Information_Tables_Groups_Group_GroupInformations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group name ,status  and information. The type is slice of
    // Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation.
    GroupInformation []Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation
}

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetEntityData() *types.CommonEntityData {
    groupInformations.EntityData.YFilter = groupInformations.YFilter
    groupInformations.EntityData.YangName = "group-informations"
    groupInformations.EntityData.BundleName = "cisco_ios_xr"
    groupInformations.EntityData.ParentYangName = "group"
    groupInformations.EntityData.SegmentPath = "group-informations"
    groupInformations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupInformations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupInformations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupInformations.EntityData.Children = make(map[string]types.YChild)
    groupInformations.EntityData.Children["group-information"] = types.YChild{"GroupInformation", nil}
    for i := range groupInformations.GroupInformation {
        groupInformations.EntityData.Children[types.GetSegmentPath(&groupInformations.GroupInformation[i])] = types.YChild{"GroupInformation", &groupInformations.GroupInformation[i]}
    }
    groupInformations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groupInformations.EntityData)
}

// Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation
// Group name ,status  and information
type Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Model number. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Modelnumber interface{}

    // Level. The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    groupInformation.EntityData.SegmentPath = "group-information"
    groupInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupInformation.EntityData.Children = make(map[string]types.YChild)
    groupInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    groupInformation.EntityData.Leafs["modelnumber"] = types.YLeaf{"Modelnumber", groupInformation.Modelnumber}
    groupInformation.EntityData.Leafs["level"] = types.YLeaf{"Level", groupInformation.Level}
    groupInformation.EntityData.Leafs["vacm-access-read-view-name"] = types.YLeaf{"VacmAccessReadViewName", groupInformation.VacmAccessReadViewName}
    groupInformation.EntityData.Leafs["vacm-access-write-view-name"] = types.YLeaf{"VacmAccessWriteViewName", groupInformation.VacmAccessWriteViewName}
    groupInformation.EntityData.Leafs["vacm-access-notify-view-name"] = types.YLeaf{"VacmAccessNotifyViewName", groupInformation.VacmAccessNotifyViewName}
    groupInformation.EntityData.Leafs["vacm-access-status"] = types.YLeaf{"VacmAccessStatus", groupInformation.VacmAccessStatus}
    return &(groupInformation.EntityData)
}

// Snmp_Information_Tables_UserEngineIds
// List of User
type Snmp_Information_Tables_UserEngineIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP engineId. The type is slice of
    // Snmp_Information_Tables_UserEngineIds_UserEngineId.
    UserEngineId []Snmp_Information_Tables_UserEngineIds_UserEngineId
}

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetEntityData() *types.CommonEntityData {
    userEngineIds.EntityData.YFilter = userEngineIds.YFilter
    userEngineIds.EntityData.YangName = "user-engine-ids"
    userEngineIds.EntityData.BundleName = "cisco_ios_xr"
    userEngineIds.EntityData.ParentYangName = "tables"
    userEngineIds.EntityData.SegmentPath = "user-engine-ids"
    userEngineIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userEngineIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userEngineIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userEngineIds.EntityData.Children = make(map[string]types.YChild)
    userEngineIds.EntityData.Children["user-engine-id"] = types.YChild{"UserEngineId", nil}
    for i := range userEngineIds.UserEngineId {
        userEngineIds.EntityData.Children[types.GetSegmentPath(&userEngineIds.UserEngineId[i])] = types.YChild{"UserEngineId", &userEngineIds.UserEngineId[i]}
    }
    userEngineIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(userEngineIds.EntityData)
}

// Snmp_Information_Tables_UserEngineIds_UserEngineId
// SNMP engineId
type Snmp_Information_Tables_UserEngineIds_UserEngineId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP Engine ID. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    EngineId interface{}

    // User name ,storage type ,status . The type is slice of
    // Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName.
    UserName []Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName
}

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetEntityData() *types.CommonEntityData {
    userEngineId.EntityData.YFilter = userEngineId.YFilter
    userEngineId.EntityData.YangName = "user-engine-id"
    userEngineId.EntityData.BundleName = "cisco_ios_xr"
    userEngineId.EntityData.ParentYangName = "user-engine-ids"
    userEngineId.EntityData.SegmentPath = "user-engine-id" + "[engine-id='" + fmt.Sprintf("%v", userEngineId.EngineId) + "']"
    userEngineId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userEngineId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userEngineId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userEngineId.EntityData.Children = make(map[string]types.YChild)
    userEngineId.EntityData.Children["user-name"] = types.YChild{"UserName", nil}
    for i := range userEngineId.UserName {
        userEngineId.EntityData.Children[types.GetSegmentPath(&userEngineId.UserName[i])] = types.YChild{"UserName", &userEngineId.UserName[i]}
    }
    userEngineId.EntityData.Leafs = make(map[string]types.YLeaf)
    userEngineId.EntityData.Leafs["engine-id"] = types.YLeaf{"EngineId", userEngineId.EngineId}
    return &(userEngineId.EntityData)
}

// Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName
// User name ,storage type ,status 
type Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. User name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    userName.EntityData.SegmentPath = "user-name" + "[user-name='" + fmt.Sprintf("%v", userName.UserName) + "']"
    userName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userName.EntityData.Children = make(map[string]types.YChild)
    userName.EntityData.Leafs = make(map[string]types.YLeaf)
    userName.EntityData.Leafs["user-name"] = types.YLeaf{"UserName", userName.UserName}
    userName.EntityData.Leafs["usm-user-storage-type"] = types.YLeaf{"UsmUserStorageType", userName.UsmUserStorageType}
    userName.EntityData.Leafs["usm-user-status"] = types.YLeaf{"UsmUserStatus", userName.UsmUserStatus}
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
    systemOid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemOid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemOid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemOid.EntityData.Children = make(map[string]types.YChild)
    systemOid.EntityData.Leafs = make(map[string]types.YLeaf)
    systemOid.EntityData.Leafs["sys-obj-id"] = types.YLeaf{"SysObjId", systemOid.SysObjId}
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
    TrapQ []Snmp_Information_TrapQueue_TrapQ
}

func (trapQueue *Snmp_Information_TrapQueue) GetEntityData() *types.CommonEntityData {
    trapQueue.EntityData.YFilter = trapQueue.YFilter
    trapQueue.EntityData.YangName = "trap-queue"
    trapQueue.EntityData.BundleName = "cisco_ios_xr"
    trapQueue.EntityData.ParentYangName = "information"
    trapQueue.EntityData.SegmentPath = "trap-queue"
    trapQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapQueue.EntityData.Children = make(map[string]types.YChild)
    trapQueue.EntityData.Children["trap-q"] = types.YChild{"TrapQ", nil}
    for i := range trapQueue.TrapQ {
        trapQueue.EntityData.Children[types.GetSegmentPath(&trapQueue.TrapQ[i])] = types.YChild{"TrapQ", &trapQueue.TrapQ[i]}
    }
    trapQueue.EntityData.Leafs = make(map[string]types.YLeaf)
    trapQueue.EntityData.Leafs["trap-min"] = types.YLeaf{"TrapMin", trapQueue.TrapMin}
    trapQueue.EntityData.Leafs["trap-avg"] = types.YLeaf{"TrapAvg", trapQueue.TrapAvg}
    trapQueue.EntityData.Leafs["trap-max"] = types.YLeaf{"TrapMax", trapQueue.TrapMax}
    return &(trapQueue.EntityData)
}

// Snmp_Information_TrapQueue_TrapQ
// trap q
type Snmp_Information_TrapQueue_TrapQ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    trapQ.EntityData.SegmentPath = "trap-q"
    trapQ.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapQ.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapQ.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapQ.EntityData.Children = make(map[string]types.YChild)
    trapQ.EntityData.Leafs = make(map[string]types.YLeaf)
    trapQ.EntityData.Leafs["min"] = types.YLeaf{"Min", trapQ.Min}
    trapQ.EntityData.Leafs["avg"] = types.YLeaf{"Avg", trapQ.Avg}
    trapQ.EntityData.Leafs["max"] = types.YLeaf{"Max", trapQ.Max}
    return &(trapQ.EntityData)
}

// Snmp_Interfaces
// List of interfaces
type Snmp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is slice of Snmp_Interfaces_Interface_.
    Interface_ []Snmp_Interfaces_Interface
}

func (interfaces *Snmp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "snmp"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Snmp_Interfaces_Interface
// Interface Name
type Snmp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // Interface Index as used by MIB tables. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    InterfaceIndex interface{}
}

func (self *Snmp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    self.EntityData.Leafs["interface-index"] = types.YLeaf{"InterfaceIndex", self.InterfaceIndex}
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
    correlator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    correlator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    correlator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    correlator.EntityData.Children = make(map[string]types.YChild)
    correlator.EntityData.Children["rule-details"] = types.YChild{"RuleDetails", &correlator.RuleDetails}
    correlator.EntityData.Children["buffer-status"] = types.YChild{"BufferStatus", &correlator.BufferStatus}
    correlator.EntityData.Children["rule-set-details"] = types.YChild{"RuleSetDetails", &correlator.RuleSetDetails}
    correlator.EntityData.Children["traps"] = types.YChild{"Traps", &correlator.Traps}
    correlator.EntityData.Leafs = make(map[string]types.YLeaf)
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
    RuleDetail []Snmp_Correlator_RuleDetails_RuleDetail
}

func (ruleDetails *Snmp_Correlator_RuleDetails) GetEntityData() *types.CommonEntityData {
    ruleDetails.EntityData.YFilter = ruleDetails.YFilter
    ruleDetails.EntityData.YangName = "rule-details"
    ruleDetails.EntityData.BundleName = "cisco_ios_xr"
    ruleDetails.EntityData.ParentYangName = "correlator"
    ruleDetails.EntityData.SegmentPath = "rule-details"
    ruleDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleDetails.EntityData.Children = make(map[string]types.YChild)
    ruleDetails.EntityData.Children["rule-detail"] = types.YChild{"RuleDetail", nil}
    for i := range ruleDetails.RuleDetail {
        ruleDetails.EntityData.Children[types.GetSegmentPath(&ruleDetails.RuleDetail[i])] = types.YChild{"RuleDetail", &ruleDetails.RuleDetail[i]}
    }
    ruleDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ruleDetails.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail
// Details of one of the correlation rules
type Snmp_Correlator_RuleDetails_RuleDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    NonRootcaus []Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus

    // Hosts (IP/port) to which the rule is applied. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost.
    ApplyHost []Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost
}

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetEntityData() *types.CommonEntityData {
    ruleDetail.EntityData.YFilter = ruleDetail.YFilter
    ruleDetail.EntityData.YangName = "rule-detail"
    ruleDetail.EntityData.BundleName = "cisco_ios_xr"
    ruleDetail.EntityData.ParentYangName = "rule-details"
    ruleDetail.EntityData.SegmentPath = "rule-detail" + "[rule-name='" + fmt.Sprintf("%v", ruleDetail.RuleName) + "']"
    ruleDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleDetail.EntityData.Children = make(map[string]types.YChild)
    ruleDetail.EntityData.Children["rule-summary"] = types.YChild{"RuleSummary", &ruleDetail.RuleSummary}
    ruleDetail.EntityData.Children["root-cause"] = types.YChild{"RootCause", &ruleDetail.RootCause}
    ruleDetail.EntityData.Children["non-rootcaus"] = types.YChild{"NonRootcaus", nil}
    for i := range ruleDetail.NonRootcaus {
        ruleDetail.EntityData.Children[types.GetSegmentPath(&ruleDetail.NonRootcaus[i])] = types.YChild{"NonRootcaus", &ruleDetail.NonRootcaus[i]}
    }
    ruleDetail.EntityData.Children["apply-host"] = types.YChild{"ApplyHost", nil}
    for i := range ruleDetail.ApplyHost {
        ruleDetail.EntityData.Children[types.GetSegmentPath(&ruleDetail.ApplyHost[i])] = types.YChild{"ApplyHost", &ruleDetail.ApplyHost[i]}
    }
    ruleDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    ruleDetail.EntityData.Leafs["rule-name"] = types.YLeaf{"RuleName", ruleDetail.RuleName}
    ruleDetail.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", ruleDetail.Timeout}
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
    ruleSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSummary.EntityData.Children = make(map[string]types.YChild)
    ruleSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    ruleSummary.EntityData.Leafs["rule-name"] = types.YLeaf{"RuleName", ruleSummary.RuleName}
    ruleSummary.EntityData.Leafs["rule-state"] = types.YLeaf{"RuleState", ruleSummary.RuleState}
    ruleSummary.EntityData.Leafs["buffered-traps-count"] = types.YLeaf{"BufferedTrapsCount", ruleSummary.BufferedTrapsCount}
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
    VarBind []Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind
}

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetEntityData() *types.CommonEntityData {
    rootCause.EntityData.YFilter = rootCause.YFilter
    rootCause.EntityData.YangName = "root-cause"
    rootCause.EntityData.BundleName = "cisco_ios_xr"
    rootCause.EntityData.ParentYangName = "rule-detail"
    rootCause.EntityData.SegmentPath = "root-cause"
    rootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rootCause.EntityData.Children = make(map[string]types.YChild)
    rootCause.EntityData.Children["var-bind"] = types.YChild{"VarBind", nil}
    for i := range rootCause.VarBind {
        rootCause.EntityData.Children[types.GetSegmentPath(&rootCause.VarBind[i])] = types.YChild{"VarBind", &rootCause.VarBind[i]}
    }
    rootCause.EntityData.Leafs = make(map[string]types.YLeaf)
    rootCause.EntityData.Leafs["oid"] = types.YLeaf{"Oid", rootCause.Oid}
    return &(rootCause.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind
// VarBinds of the trap
type Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    varBind.EntityData.SegmentPath = "var-bind"
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = make(map[string]types.YChild)
    varBind.EntityData.Leafs = make(map[string]types.YLeaf)
    varBind.EntityData.Leafs["oid"] = types.YLeaf{"Oid", varBind.Oid}
    varBind.EntityData.Leafs["match-type"] = types.YLeaf{"MatchType", varBind.MatchType}
    varBind.EntityData.Leafs["reg-exp"] = types.YLeaf{"RegExp", varBind.RegExp}
    return &(varBind.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus
// OIDs/VarBinds defining the nonrootcause match
// conditions.
type Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OID of the trap. The type is string.
    Oid interface{}

    // VarBinds of the trap. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind.
    VarBind []Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind
}

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetEntityData() *types.CommonEntityData {
    nonRootcaus.EntityData.YFilter = nonRootcaus.YFilter
    nonRootcaus.EntityData.YangName = "non-rootcaus"
    nonRootcaus.EntityData.BundleName = "cisco_ios_xr"
    nonRootcaus.EntityData.ParentYangName = "rule-detail"
    nonRootcaus.EntityData.SegmentPath = "non-rootcaus"
    nonRootcaus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootcaus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootcaus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootcaus.EntityData.Children = make(map[string]types.YChild)
    nonRootcaus.EntityData.Children["var-bind"] = types.YChild{"VarBind", nil}
    for i := range nonRootcaus.VarBind {
        nonRootcaus.EntityData.Children[types.GetSegmentPath(&nonRootcaus.VarBind[i])] = types.YChild{"VarBind", &nonRootcaus.VarBind[i]}
    }
    nonRootcaus.EntityData.Leafs = make(map[string]types.YLeaf)
    nonRootcaus.EntityData.Leafs["oid"] = types.YLeaf{"Oid", nonRootcaus.Oid}
    return &(nonRootcaus.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind
// VarBinds of the trap
type Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    varBind.EntityData.SegmentPath = "var-bind"
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = make(map[string]types.YChild)
    varBind.EntityData.Leafs = make(map[string]types.YLeaf)
    varBind.EntityData.Leafs["oid"] = types.YLeaf{"Oid", varBind.Oid}
    varBind.EntityData.Leafs["match-type"] = types.YLeaf{"MatchType", varBind.MatchType}
    varBind.EntityData.Leafs["reg-exp"] = types.YLeaf{"RegExp", varBind.RegExp}
    return &(varBind.EntityData)
}

// Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost
// Hosts (IP/port) to which the rule is applied
type Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    applyHost.EntityData.SegmentPath = "apply-host"
    applyHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    applyHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    applyHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    applyHost.EntityData.Children = make(map[string]types.YChild)
    applyHost.EntityData.Leafs = make(map[string]types.YLeaf)
    applyHost.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", applyHost.IpAddress}
    applyHost.EntityData.Leafs["port"] = types.YLeaf{"Port", applyHost.Port}
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
    bufferStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bufferStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bufferStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bufferStatus.EntityData.Children = make(map[string]types.YChild)
    bufferStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    bufferStatus.EntityData.Leafs["current-size"] = types.YLeaf{"CurrentSize", bufferStatus.CurrentSize}
    bufferStatus.EntityData.Leafs["configured-size"] = types.YLeaf{"ConfiguredSize", bufferStatus.ConfiguredSize}
    return &(bufferStatus.EntityData)
}

// Snmp_Correlator_RuleSetDetails
// Table that contains the ruleset detail info
type Snmp_Correlator_RuleSetDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail of one of the correlation rulesets. The type is slice of
    // Snmp_Correlator_RuleSetDetails_RuleSetDetail.
    RuleSetDetail []Snmp_Correlator_RuleSetDetails_RuleSetDetail
}

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetEntityData() *types.CommonEntityData {
    ruleSetDetails.EntityData.YFilter = ruleSetDetails.YFilter
    ruleSetDetails.EntityData.YangName = "rule-set-details"
    ruleSetDetails.EntityData.BundleName = "cisco_ios_xr"
    ruleSetDetails.EntityData.ParentYangName = "correlator"
    ruleSetDetails.EntityData.SegmentPath = "rule-set-details"
    ruleSetDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSetDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSetDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSetDetails.EntityData.Children = make(map[string]types.YChild)
    ruleSetDetails.EntityData.Children["rule-set-detail"] = types.YChild{"RuleSetDetail", nil}
    for i := range ruleSetDetails.RuleSetDetail {
        ruleSetDetails.EntityData.Children[types.GetSegmentPath(&ruleSetDetails.RuleSetDetail[i])] = types.YChild{"RuleSetDetail", &ruleSetDetails.RuleSetDetail[i]}
    }
    ruleSetDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ruleSetDetails.EntityData)
}

// Snmp_Correlator_RuleSetDetails_RuleSetDetail
// Detail of one of the correlation rulesets
type Snmp_Correlator_RuleSetDetails_RuleSetDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ruleset Name. The type is string with length:
    // 1..32.
    RuleSetName interface{}

    // Ruleset Name. The type is string.
    RuleSetNameXr interface{}

    // Rules contained in a ruleset. The type is slice of
    // Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules.
    Rules []Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules
}

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetEntityData() *types.CommonEntityData {
    ruleSetDetail.EntityData.YFilter = ruleSetDetail.YFilter
    ruleSetDetail.EntityData.YangName = "rule-set-detail"
    ruleSetDetail.EntityData.BundleName = "cisco_ios_xr"
    ruleSetDetail.EntityData.ParentYangName = "rule-set-details"
    ruleSetDetail.EntityData.SegmentPath = "rule-set-detail" + "[rule-set-name='" + fmt.Sprintf("%v", ruleSetDetail.RuleSetName) + "']"
    ruleSetDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSetDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSetDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSetDetail.EntityData.Children = make(map[string]types.YChild)
    ruleSetDetail.EntityData.Children["rules"] = types.YChild{"Rules", nil}
    for i := range ruleSetDetail.Rules {
        ruleSetDetail.EntityData.Children[types.GetSegmentPath(&ruleSetDetail.Rules[i])] = types.YChild{"Rules", &ruleSetDetail.Rules[i]}
    }
    ruleSetDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    ruleSetDetail.EntityData.Leafs["rule-set-name"] = types.YLeaf{"RuleSetName", ruleSetDetail.RuleSetName}
    ruleSetDetail.EntityData.Leafs["rule-set-name-xr"] = types.YLeaf{"RuleSetNameXr", ruleSetDetail.RuleSetNameXr}
    return &(ruleSetDetail.EntityData)
}

// Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules
// Rules contained in a ruleset
type Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules struct {
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

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "rule-set-detail"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = make(map[string]types.YChild)
    rules.EntityData.Leafs = make(map[string]types.YLeaf)
    rules.EntityData.Leafs["rule-name"] = types.YLeaf{"RuleName", rules.RuleName}
    rules.EntityData.Leafs["rule-state"] = types.YLeaf{"RuleState", rules.RuleState}
    rules.EntityData.Leafs["buffered-traps-count"] = types.YLeaf{"BufferedTrapsCount", rules.BufferedTrapsCount}
    return &(rules.EntityData)
}

// Snmp_Correlator_Traps
// Correlated traps Table
type Snmp_Correlator_Traps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One of the correlated traps. The type is slice of
    // Snmp_Correlator_Traps_Trap.
    Trap []Snmp_Correlator_Traps_Trap
}

func (traps *Snmp_Correlator_Traps) GetEntityData() *types.CommonEntityData {
    traps.EntityData.YFilter = traps.YFilter
    traps.EntityData.YangName = "traps"
    traps.EntityData.BundleName = "cisco_ios_xr"
    traps.EntityData.ParentYangName = "correlator"
    traps.EntityData.SegmentPath = "traps"
    traps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traps.EntityData.Children = make(map[string]types.YChild)
    traps.EntityData.Children["trap"] = types.YChild{"Trap", nil}
    for i := range traps.Trap {
        traps.EntityData.Children[types.GetSegmentPath(&traps.Trap[i])] = types.YChild{"Trap", &traps.Trap[i]}
    }
    traps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(traps.EntityData)
}

// Snmp_Correlator_Traps_Trap
// One of the correlated traps
type Snmp_Correlator_Traps_Trap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Entry ID. The type is interface{} with range:
    // -2147483648..2147483647.
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
    trap.EntityData.SegmentPath = "trap" + "[entry-id='" + fmt.Sprintf("%v", trap.EntryId) + "']"
    trap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trap.EntityData.Children = make(map[string]types.YChild)
    trap.EntityData.Children["trap-info"] = types.YChild{"TrapInfo", &trap.TrapInfo}
    trap.EntityData.Leafs = make(map[string]types.YLeaf)
    trap.EntityData.Leafs["entry-id"] = types.YLeaf{"EntryId", trap.EntryId}
    trap.EntityData.Leafs["correlation-id"] = types.YLeaf{"CorrelationId", trap.CorrelationId}
    trap.EntityData.Leafs["is-root-cause"] = types.YLeaf{"IsRootCause", trap.IsRootCause}
    trap.EntityData.Leafs["rule-name"] = types.YLeaf{"RuleName", trap.RuleName}
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
    VarBind []Snmp_Correlator_Traps_Trap_TrapInfo_VarBind
}

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetEntityData() *types.CommonEntityData {
    trapInfo.EntityData.YFilter = trapInfo.YFilter
    trapInfo.EntityData.YangName = "trap-info"
    trapInfo.EntityData.BundleName = "cisco_ios_xr"
    trapInfo.EntityData.ParentYangName = "trap"
    trapInfo.EntityData.SegmentPath = "trap-info"
    trapInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapInfo.EntityData.Children = make(map[string]types.YChild)
    trapInfo.EntityData.Children["var-bind"] = types.YChild{"VarBind", nil}
    for i := range trapInfo.VarBind {
        trapInfo.EntityData.Children[types.GetSegmentPath(&trapInfo.VarBind[i])] = types.YChild{"VarBind", &trapInfo.VarBind[i]}
    }
    trapInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    trapInfo.EntityData.Leafs["oid"] = types.YLeaf{"Oid", trapInfo.Oid}
    trapInfo.EntityData.Leafs["relative-timestamp"] = types.YLeaf{"RelativeTimestamp", trapInfo.RelativeTimestamp}
    trapInfo.EntityData.Leafs["timestamp"] = types.YLeaf{"Timestamp", trapInfo.Timestamp}
    return &(trapInfo.EntityData)
}

// Snmp_Correlator_Traps_Trap_TrapInfo_VarBind
// VarBinds on the trap
type Snmp_Correlator_Traps_Trap_TrapInfo_VarBind struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    varBind.EntityData.SegmentPath = "var-bind"
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = make(map[string]types.YChild)
    varBind.EntityData.Leafs = make(map[string]types.YLeaf)
    varBind.EntityData.Leafs["oid"] = types.YLeaf{"Oid", varBind.Oid}
    varBind.EntityData.Leafs["value"] = types.YLeaf{"Value", varBind.Value}
    return &(varBind.EntityData)
}

// Snmp_InterfaceIndexes
// List of index
type Snmp_InterfaceIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Index. The type is slice of Snmp_InterfaceIndexes_InterfaceIndex.
    InterfaceIndex []Snmp_InterfaceIndexes_InterfaceIndex
}

func (interfaceIndexes *Snmp_InterfaceIndexes) GetEntityData() *types.CommonEntityData {
    interfaceIndexes.EntityData.YFilter = interfaceIndexes.YFilter
    interfaceIndexes.EntityData.YangName = "interface-indexes"
    interfaceIndexes.EntityData.BundleName = "cisco_ios_xr"
    interfaceIndexes.EntityData.ParentYangName = "snmp"
    interfaceIndexes.EntityData.SegmentPath = "interface-indexes"
    interfaceIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIndexes.EntityData.Children = make(map[string]types.YChild)
    interfaceIndexes.EntityData.Children["interface-index"] = types.YChild{"InterfaceIndex", nil}
    for i := range interfaceIndexes.InterfaceIndex {
        interfaceIndexes.EntityData.Children[types.GetSegmentPath(&interfaceIndexes.InterfaceIndex[i])] = types.YChild{"InterfaceIndex", &interfaceIndexes.InterfaceIndex[i]}
    }
    interfaceIndexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceIndexes.EntityData)
}

// Snmp_InterfaceIndexes_InterfaceIndex
// Interface Index
type Snmp_InterfaceIndexes_InterfaceIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Index as used by MIB tables. The type is
    // interface{} with range: -2147483648..2147483647.
    InterfaceIndex interface{}

    // Interface Name. The type is string. This attribute is mandatory.
    InterfaceName interface{}
}

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetEntityData() *types.CommonEntityData {
    interfaceIndex.EntityData.YFilter = interfaceIndex.YFilter
    interfaceIndex.EntityData.YangName = "interface-index"
    interfaceIndex.EntityData.BundleName = "cisco_ios_xr"
    interfaceIndex.EntityData.ParentYangName = "interface-indexes"
    interfaceIndex.EntityData.SegmentPath = "interface-index" + "[interface-index='" + fmt.Sprintf("%v", interfaceIndex.InterfaceIndex) + "']"
    interfaceIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceIndex.EntityData.Children = make(map[string]types.YChild)
    interfaceIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceIndex.EntityData.Leafs["interface-index"] = types.YLeaf{"InterfaceIndex", interfaceIndex.InterfaceIndex}
    interfaceIndex.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceIndex.InterfaceName}
    return &(interfaceIndex.EntityData)
}

// Snmp_IfIndexes
// List of ifnames
type Snmp_IfIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Index. The type is slice of Snmp_IfIndexes_IfIndex.
    IfIndex []Snmp_IfIndexes_IfIndex
}

func (ifIndexes *Snmp_IfIndexes) GetEntityData() *types.CommonEntityData {
    ifIndexes.EntityData.YFilter = ifIndexes.YFilter
    ifIndexes.EntityData.YangName = "if-indexes"
    ifIndexes.EntityData.BundleName = "cisco_ios_xr"
    ifIndexes.EntityData.ParentYangName = "snmp"
    ifIndexes.EntityData.SegmentPath = "if-indexes"
    ifIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifIndexes.EntityData.Children = make(map[string]types.YChild)
    ifIndexes.EntityData.Children["if-index"] = types.YChild{"IfIndex", nil}
    for i := range ifIndexes.IfIndex {
        ifIndexes.EntityData.Children[types.GetSegmentPath(&ifIndexes.IfIndex[i])] = types.YChild{"IfIndex", &ifIndexes.IfIndex[i]}
    }
    ifIndexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ifIndexes.EntityData)
}

// Snmp_IfIndexes_IfIndex
// Interface Index
type Snmp_IfIndexes_IfIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Index as used by MIB tables. The type is
    // interface{} with range: -2147483648..2147483647.
    InterfaceIndex interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (ifIndex *Snmp_IfIndexes_IfIndex) GetEntityData() *types.CommonEntityData {
    ifIndex.EntityData.YFilter = ifIndex.YFilter
    ifIndex.EntityData.YangName = "if-index"
    ifIndex.EntityData.BundleName = "cisco_ios_xr"
    ifIndex.EntityData.ParentYangName = "if-indexes"
    ifIndex.EntityData.SegmentPath = "if-index" + "[interface-index='" + fmt.Sprintf("%v", ifIndex.InterfaceIndex) + "']"
    ifIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifIndex.EntityData.Children = make(map[string]types.YChild)
    ifIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    ifIndex.EntityData.Leafs["interface-index"] = types.YLeaf{"InterfaceIndex", ifIndex.InterfaceIndex}
    ifIndex.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", ifIndex.InterfaceName}
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
    entityMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityMib.EntityData.Children = make(map[string]types.YChild)
    entityMib.EntityData.Children["entity-physical-indexes"] = types.YChild{"EntityPhysicalIndexes", &entityMib.EntityPhysicalIndexes}
    entityMib.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entityMib.EntityData)
}

// Snmp_EntityMib_EntityPhysicalIndexes
// SNMP entity mib
type Snmp_EntityMib_EntityPhysicalIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP entPhysical index number. The type is slice of
    // Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex.
    EntityPhysicalIndex []Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex
}

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetEntityData() *types.CommonEntityData {
    entityPhysicalIndexes.EntityData.YFilter = entityPhysicalIndexes.YFilter
    entityPhysicalIndexes.EntityData.YangName = "entity-physical-indexes"
    entityPhysicalIndexes.EntityData.BundleName = "cisco_ios_xr"
    entityPhysicalIndexes.EntityData.ParentYangName = "entity-mib"
    entityPhysicalIndexes.EntityData.SegmentPath = "entity-physical-indexes"
    entityPhysicalIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityPhysicalIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityPhysicalIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityPhysicalIndexes.EntityData.Children = make(map[string]types.YChild)
    entityPhysicalIndexes.EntityData.Children["entity-physical-index"] = types.YChild{"EntityPhysicalIndex", nil}
    for i := range entityPhysicalIndexes.EntityPhysicalIndex {
        entityPhysicalIndexes.EntityData.Children[types.GetSegmentPath(&entityPhysicalIndexes.EntityPhysicalIndex[i])] = types.YChild{"EntityPhysicalIndex", &entityPhysicalIndexes.EntityPhysicalIndex[i]}
    }
    entityPhysicalIndexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entityPhysicalIndexes.EntityData)
}

// Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex
// SNMP entPhysical index number
type Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Entity physical index. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    entityPhysicalIndex.EntityData.SegmentPath = "entity-physical-index" + "[entity-phynum='" + fmt.Sprintf("%v", entityPhysicalIndex.EntityPhynum) + "']"
    entityPhysicalIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityPhysicalIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityPhysicalIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityPhysicalIndex.EntityData.Children = make(map[string]types.YChild)
    entityPhysicalIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    entityPhysicalIndex.EntityData.Leafs["entity-phynum"] = types.YLeaf{"EntityPhynum", entityPhysicalIndex.EntityPhynum}
    entityPhysicalIndex.EntityData.Leafs["physical-index"] = types.YLeaf{"PhysicalIndex", entityPhysicalIndex.PhysicalIndex}
    entityPhysicalIndex.EntityData.Leafs["ent-physical-name"] = types.YLeaf{"EntPhysicalName", entityPhysicalIndex.EntPhysicalName}
    entityPhysicalIndex.EntityData.Leafs["location"] = types.YLeaf{"Location", entityPhysicalIndex.Location}
    entityPhysicalIndex.EntityData.Leafs["ent-physical-descr"] = types.YLeaf{"EntPhysicalDescr", entityPhysicalIndex.EntPhysicalDescr}
    entityPhysicalIndex.EntityData.Leafs["ent-physical-firmware-rev"] = types.YLeaf{"EntPhysicalFirmwareRev", entityPhysicalIndex.EntPhysicalFirmwareRev}
    entityPhysicalIndex.EntityData.Leafs["ent-physical-hardware-rev"] = types.YLeaf{"EntPhysicalHardwareRev", entityPhysicalIndex.EntPhysicalHardwareRev}
    entityPhysicalIndex.EntityData.Leafs["ent-physical-modelname"] = types.YLeaf{"EntPhysicalModelname", entityPhysicalIndex.EntPhysicalModelname}
    entityPhysicalIndex.EntityData.Leafs["ent-physical-serial-num"] = types.YLeaf{"EntPhysicalSerialNum", entityPhysicalIndex.EntPhysicalSerialNum}
    entityPhysicalIndex.EntityData.Leafs["ent-physical-software-rev"] = types.YLeaf{"EntPhysicalSoftwareRev", entityPhysicalIndex.EntPhysicalSoftwareRev}
    entityPhysicalIndex.EntityData.Leafs["ent-physical-mfg-name"] = types.YLeaf{"EntPhysicalMfgName", entityPhysicalIndex.EntPhysicalMfgName}
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
    interfaceMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMib.EntityData.Children = make(map[string]types.YChild)
    interfaceMib.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &interfaceMib.Interfaces}
    interfaceMib.EntityData.Children["interface-connectors"] = types.YChild{"InterfaceConnectors", &interfaceMib.InterfaceConnectors}
    interfaceMib.EntityData.Children["interface-aliases"] = types.YChild{"InterfaceAliases", &interfaceMib.InterfaceAliases}
    interfaceMib.EntityData.Children["notification-interfaces"] = types.YChild{"NotificationInterfaces", &interfaceMib.NotificationInterfaces}
    interfaceMib.EntityData.Children["interface-stack-statuses"] = types.YChild{"InterfaceStackStatuses", &interfaceMib.InterfaceStackStatuses}
    interfaceMib.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceMib.EntityData)
}

// Snmp_InterfaceMib_Interfaces
// Interfaces ifIndex information
type Snmp_InterfaceMib_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ifIndex for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_Interfaces_Interface_.
    Interface_ []Snmp_InterfaceMib_Interfaces_Interface
}

func (interfaces *Snmp_InterfaceMib_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-mib"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Snmp_InterfaceMib_Interfaces_Interface
// ifIndex for a specific Interface Name
type Snmp_InterfaceMib_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface Index. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["if-index"] = types.YLeaf{"IfIndex", self.IfIndex}
    return &(self.EntityData)
}

// Snmp_InterfaceMib_InterfaceConnectors
// Interfaces ifConnectorPresent information
type Snmp_InterfaceMib_InterfaceConnectors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ifConnectorPresent for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector.
    InterfaceConnector []Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector
}

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetEntityData() *types.CommonEntityData {
    interfaceConnectors.EntityData.YFilter = interfaceConnectors.YFilter
    interfaceConnectors.EntityData.YangName = "interface-connectors"
    interfaceConnectors.EntityData.BundleName = "cisco_ios_xr"
    interfaceConnectors.EntityData.ParentYangName = "interface-mib"
    interfaceConnectors.EntityData.SegmentPath = "interface-connectors"
    interfaceConnectors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConnectors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConnectors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConnectors.EntityData.Children = make(map[string]types.YChild)
    interfaceConnectors.EntityData.Children["interface-connector"] = types.YChild{"InterfaceConnector", nil}
    for i := range interfaceConnectors.InterfaceConnector {
        interfaceConnectors.EntityData.Children[types.GetSegmentPath(&interfaceConnectors.InterfaceConnector[i])] = types.YChild{"InterfaceConnector", &interfaceConnectors.InterfaceConnector[i]}
    }
    interfaceConnectors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceConnectors.EntityData)
}

// Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector
// ifConnectorPresent for a specific Interface
// Name
type Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface ifConnector. The type is string.
    IfConnectorPresent interface{}
}

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetEntityData() *types.CommonEntityData {
    interfaceConnector.EntityData.YFilter = interfaceConnector.YFilter
    interfaceConnector.EntityData.YangName = "interface-connector"
    interfaceConnector.EntityData.BundleName = "cisco_ios_xr"
    interfaceConnector.EntityData.ParentYangName = "interface-connectors"
    interfaceConnector.EntityData.SegmentPath = "interface-connector" + "[interface-name='" + fmt.Sprintf("%v", interfaceConnector.InterfaceName) + "']"
    interfaceConnector.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceConnector.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceConnector.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceConnector.EntityData.Children = make(map[string]types.YChild)
    interfaceConnector.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceConnector.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceConnector.InterfaceName}
    interfaceConnector.EntityData.Leafs["if-connector-present"] = types.YLeaf{"IfConnectorPresent", interfaceConnector.IfConnectorPresent}
    return &(interfaceConnector.EntityData)
}

// Snmp_InterfaceMib_InterfaceAliases
// Interfaces ifAlias information
type Snmp_InterfaceMib_InterfaceAliases struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ifAlias for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias.
    InterfaceAlias []Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias
}

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetEntityData() *types.CommonEntityData {
    interfaceAliases.EntityData.YFilter = interfaceAliases.YFilter
    interfaceAliases.EntityData.YangName = "interface-aliases"
    interfaceAliases.EntityData.BundleName = "cisco_ios_xr"
    interfaceAliases.EntityData.ParentYangName = "interface-mib"
    interfaceAliases.EntityData.SegmentPath = "interface-aliases"
    interfaceAliases.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAliases.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAliases.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAliases.EntityData.Children = make(map[string]types.YChild)
    interfaceAliases.EntityData.Children["interface-alias"] = types.YChild{"InterfaceAlias", nil}
    for i := range interfaceAliases.InterfaceAlias {
        interfaceAliases.EntityData.Children[types.GetSegmentPath(&interfaceAliases.InterfaceAlias[i])] = types.YChild{"InterfaceAlias", &interfaceAliases.InterfaceAlias[i]}
    }
    interfaceAliases.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceAliases.EntityData)
}

// Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias
// ifAlias for a specific Interface Name
type Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Interface ifAlias. The type is string.
    IfAlias interface{}
}

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetEntityData() *types.CommonEntityData {
    interfaceAlias.EntityData.YFilter = interfaceAlias.YFilter
    interfaceAlias.EntityData.YangName = "interface-alias"
    interfaceAlias.EntityData.BundleName = "cisco_ios_xr"
    interfaceAlias.EntityData.ParentYangName = "interface-aliases"
    interfaceAlias.EntityData.SegmentPath = "interface-alias" + "[interface-name='" + fmt.Sprintf("%v", interfaceAlias.InterfaceName) + "']"
    interfaceAlias.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAlias.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAlias.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAlias.EntityData.Children = make(map[string]types.YChild)
    interfaceAlias.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceAlias.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceAlias.InterfaceName}
    interfaceAlias.EntityData.Leafs["if-alias"] = types.YLeaf{"IfAlias", interfaceAlias.IfAlias}
    return &(interfaceAlias.EntityData)
}

// Snmp_InterfaceMib_NotificationInterfaces
// Interfaces Notification information
type Snmp_InterfaceMib_NotificationInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Notification for specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface.
    NotificationInterface []Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface
}

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetEntityData() *types.CommonEntityData {
    notificationInterfaces.EntityData.YFilter = notificationInterfaces.YFilter
    notificationInterfaces.EntityData.YangName = "notification-interfaces"
    notificationInterfaces.EntityData.BundleName = "cisco_ios_xr"
    notificationInterfaces.EntityData.ParentYangName = "interface-mib"
    notificationInterfaces.EntityData.SegmentPath = "notification-interfaces"
    notificationInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationInterfaces.EntityData.Children = make(map[string]types.YChild)
    notificationInterfaces.EntityData.Children["notification-interface"] = types.YChild{"NotificationInterface", nil}
    for i := range notificationInterfaces.NotificationInterface {
        notificationInterfaces.EntityData.Children[types.GetSegmentPath(&notificationInterfaces.NotificationInterface[i])] = types.YChild{"NotificationInterface", &notificationInterfaces.NotificationInterface[i]}
    }
    notificationInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(notificationInterfaces.EntityData)
}

// Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface
// Notification for specific Interface Name
type Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // LinkUpDown notification status. The type is LinkUpDownStatus.
    LinkUpDownNotifStatus interface{}
}

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetEntityData() *types.CommonEntityData {
    notificationInterface.EntityData.YFilter = notificationInterface.YFilter
    notificationInterface.EntityData.YangName = "notification-interface"
    notificationInterface.EntityData.BundleName = "cisco_ios_xr"
    notificationInterface.EntityData.ParentYangName = "notification-interfaces"
    notificationInterface.EntityData.SegmentPath = "notification-interface" + "[interface-name='" + fmt.Sprintf("%v", notificationInterface.InterfaceName) + "']"
    notificationInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationInterface.EntityData.Children = make(map[string]types.YChild)
    notificationInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    notificationInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", notificationInterface.InterfaceName}
    notificationInterface.EntityData.Leafs["link-up-down-notif-status"] = types.YLeaf{"LinkUpDownNotifStatus", notificationInterface.LinkUpDownNotifStatus}
    return &(notificationInterface.EntityData)
}

// Snmp_InterfaceMib_InterfaceStackStatuses
// Interfaces ifstackstatus information
type Snmp_InterfaceMib_InterfaceStackStatuses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ifstatus for a pair of Interface. The type is slice of
    // Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus.
    InterfaceStackStatus []Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus
}

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetEntityData() *types.CommonEntityData {
    interfaceStackStatuses.EntityData.YFilter = interfaceStackStatuses.YFilter
    interfaceStackStatuses.EntityData.YangName = "interface-stack-statuses"
    interfaceStackStatuses.EntityData.BundleName = "cisco_ios_xr"
    interfaceStackStatuses.EntityData.ParentYangName = "interface-mib"
    interfaceStackStatuses.EntityData.SegmentPath = "interface-stack-statuses"
    interfaceStackStatuses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStackStatuses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStackStatuses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStackStatuses.EntityData.Children = make(map[string]types.YChild)
    interfaceStackStatuses.EntityData.Children["interface-stack-status"] = types.YChild{"InterfaceStackStatus", nil}
    for i := range interfaceStackStatuses.InterfaceStackStatus {
        interfaceStackStatuses.EntityData.Children[types.GetSegmentPath(&interfaceStackStatuses.InterfaceStackStatus[i])] = types.YChild{"InterfaceStackStatus", &interfaceStackStatuses.InterfaceStackStatus[i]}
    }
    interfaceStackStatuses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceStackStatuses.EntityData)
}

// Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus
// ifstatus for a pair of Interface
type Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. StackHigherLayer.StackLowerLayer. The type is
    // string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    interfaceStackStatus.EntityData.SegmentPath = "interface-stack-status" + "[interface-stack-status='" + fmt.Sprintf("%v", interfaceStackStatus.InterfaceStackStatus) + "']"
    interfaceStackStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceStackStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceStackStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceStackStatus.EntityData.Children = make(map[string]types.YChild)
    interfaceStackStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceStackStatus.EntityData.Leafs["interface-stack-status"] = types.YLeaf{"InterfaceStackStatus", interfaceStackStatus.InterfaceStackStatus}
    interfaceStackStatus.EntityData.Leafs["if-stack-higher-layer"] = types.YLeaf{"IfStackHigherLayer", interfaceStackStatus.IfStackHigherLayer}
    interfaceStackStatus.EntityData.Leafs["if-stack-lower-layer"] = types.YLeaf{"IfStackLowerLayer", interfaceStackStatus.IfStackLowerLayer}
    interfaceStackStatus.EntityData.Leafs["if-stack-status"] = types.YLeaf{"IfStackStatus", interfaceStackStatus.IfStackStatus}
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
    sensorMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensorMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensorMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensorMib.EntityData.Children = make(map[string]types.YChild)
    sensorMib.EntityData.Children["physical-indexes"] = types.YChild{"PhysicalIndexes", &sensorMib.PhysicalIndexes}
    sensorMib.EntityData.Children["ent-phy-indexes"] = types.YChild{"EntPhyIndexes", &sensorMib.EntPhyIndexes}
    sensorMib.EntityData.Leafs = make(map[string]types.YLeaf)
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
    PhysicalIndex []Snmp_SensorMib_PhysicalIndexes_PhysicalIndex
}

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetEntityData() *types.CommonEntityData {
    physicalIndexes.EntityData.YFilter = physicalIndexes.YFilter
    physicalIndexes.EntityData.YangName = "physical-indexes"
    physicalIndexes.EntityData.BundleName = "cisco_ios_xr"
    physicalIndexes.EntityData.ParentYangName = "sensor-mib"
    physicalIndexes.EntityData.SegmentPath = "physical-indexes"
    physicalIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    physicalIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    physicalIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    physicalIndexes.EntityData.Children = make(map[string]types.YChild)
    physicalIndexes.EntityData.Children["physical-index"] = types.YChild{"PhysicalIndex", nil}
    for i := range physicalIndexes.PhysicalIndex {
        physicalIndexes.EntityData.Children[types.GetSegmentPath(&physicalIndexes.PhysicalIndex[i])] = types.YChild{"PhysicalIndex", &physicalIndexes.PhysicalIndex[i]}
    }
    physicalIndexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(physicalIndexes.EntityData)
}

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex
// Threshold value for physical index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Physical index. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Index interface{}

    // List of threshold index.
    ThresholdIndexes Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes
}

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetEntityData() *types.CommonEntityData {
    physicalIndex.EntityData.YFilter = physicalIndex.YFilter
    physicalIndex.EntityData.YangName = "physical-index"
    physicalIndex.EntityData.BundleName = "cisco_ios_xr"
    physicalIndex.EntityData.ParentYangName = "physical-indexes"
    physicalIndex.EntityData.SegmentPath = "physical-index" + "[index='" + fmt.Sprintf("%v", physicalIndex.Index) + "']"
    physicalIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    physicalIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    physicalIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    physicalIndex.EntityData.Children = make(map[string]types.YChild)
    physicalIndex.EntityData.Children["threshold-indexes"] = types.YChild{"ThresholdIndexes", &physicalIndex.ThresholdIndexes}
    physicalIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    physicalIndex.EntityData.Leafs["index"] = types.YLeaf{"Index", physicalIndex.Index}
    return &(physicalIndex.EntityData)
}

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes
// List of threshold index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold value for threshold index. The type is slice of
    // Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex.
    ThresholdIndex []Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex
}

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetEntityData() *types.CommonEntityData {
    thresholdIndexes.EntityData.YFilter = thresholdIndexes.YFilter
    thresholdIndexes.EntityData.YangName = "threshold-indexes"
    thresholdIndexes.EntityData.BundleName = "cisco_ios_xr"
    thresholdIndexes.EntityData.ParentYangName = "physical-index"
    thresholdIndexes.EntityData.SegmentPath = "threshold-indexes"
    thresholdIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdIndexes.EntityData.Children = make(map[string]types.YChild)
    thresholdIndexes.EntityData.Children["threshold-index"] = types.YChild{"ThresholdIndex", nil}
    for i := range thresholdIndexes.ThresholdIndex {
        thresholdIndexes.EntityData.Children[types.GetSegmentPath(&thresholdIndexes.ThresholdIndex[i])] = types.YChild{"ThresholdIndex", &thresholdIndexes.ThresholdIndex[i]}
    }
    thresholdIndexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(thresholdIndexes.EntityData)
}

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex
// Threshold value for threshold index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Physical Index. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PhyIndex interface{}

    // Threshold index. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    thresholdIndex.EntityData.SegmentPath = "threshold-index"
    thresholdIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thresholdIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thresholdIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thresholdIndex.EntityData.Children = make(map[string]types.YChild)
    thresholdIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholdIndex.EntityData.Leafs["phy-index"] = types.YLeaf{"PhyIndex", thresholdIndex.PhyIndex}
    thresholdIndex.EntityData.Leafs["thre-index"] = types.YLeaf{"ThreIndex", thresholdIndex.ThreIndex}
    thresholdIndex.EntityData.Leafs["threshold-severity"] = types.YLeaf{"ThresholdSeverity", thresholdIndex.ThresholdSeverity}
    thresholdIndex.EntityData.Leafs["threshold-relation"] = types.YLeaf{"ThresholdRelation", thresholdIndex.ThresholdRelation}
    thresholdIndex.EntityData.Leafs["threshold-value"] = types.YLeaf{"ThresholdValue", thresholdIndex.ThresholdValue}
    thresholdIndex.EntityData.Leafs["threshold-evaluation"] = types.YLeaf{"ThresholdEvaluation", thresholdIndex.ThresholdEvaluation}
    thresholdIndex.EntityData.Leafs["threshold-notification-enabled"] = types.YLeaf{"ThresholdNotificationEnabled", thresholdIndex.ThresholdNotificationEnabled}
    return &(thresholdIndex.EntityData)
}

// Snmp_SensorMib_EntPhyIndexes
// List of physical index 
type Snmp_SensorMib_EntPhyIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sensor value for physical index. The type is slice of
    // Snmp_SensorMib_EntPhyIndexes_EntPhyIndex.
    EntPhyIndex []Snmp_SensorMib_EntPhyIndexes_EntPhyIndex
}

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetEntityData() *types.CommonEntityData {
    entPhyIndexes.EntityData.YFilter = entPhyIndexes.YFilter
    entPhyIndexes.EntityData.YangName = "ent-phy-indexes"
    entPhyIndexes.EntityData.BundleName = "cisco_ios_xr"
    entPhyIndexes.EntityData.ParentYangName = "sensor-mib"
    entPhyIndexes.EntityData.SegmentPath = "ent-phy-indexes"
    entPhyIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entPhyIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entPhyIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entPhyIndexes.EntityData.Children = make(map[string]types.YChild)
    entPhyIndexes.EntityData.Children["ent-phy-index"] = types.YChild{"EntPhyIndex", nil}
    for i := range entPhyIndexes.EntPhyIndex {
        entPhyIndexes.EntityData.Children[types.GetSegmentPath(&entPhyIndexes.EntPhyIndex[i])] = types.YChild{"EntPhyIndex", &entPhyIndexes.EntPhyIndex[i]}
    }
    entPhyIndexes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(entPhyIndexes.EntityData)
}

// Snmp_SensorMib_EntPhyIndexes_EntPhyIndex
// Sensor value for physical index
type Snmp_SensorMib_EntPhyIndexes_EntPhyIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Physical index. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    entPhyIndex.EntityData.SegmentPath = "ent-phy-index" + "[index='" + fmt.Sprintf("%v", entPhyIndex.Index) + "']"
    entPhyIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entPhyIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entPhyIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entPhyIndex.EntityData.Children = make(map[string]types.YChild)
    entPhyIndex.EntityData.Leafs = make(map[string]types.YLeaf)
    entPhyIndex.EntityData.Leafs["index"] = types.YLeaf{"Index", entPhyIndex.Index}
    entPhyIndex.EntityData.Leafs["field-validity-bitmap"] = types.YLeaf{"FieldValidityBitmap", entPhyIndex.FieldValidityBitmap}
    entPhyIndex.EntityData.Leafs["device-description"] = types.YLeaf{"DeviceDescription", entPhyIndex.DeviceDescription}
    entPhyIndex.EntityData.Leafs["units"] = types.YLeaf{"Units", entPhyIndex.Units}
    entPhyIndex.EntityData.Leafs["device-id"] = types.YLeaf{"DeviceId", entPhyIndex.DeviceId}
    entPhyIndex.EntityData.Leafs["value"] = types.YLeaf{"Value", entPhyIndex.Value}
    entPhyIndex.EntityData.Leafs["alarm-type"] = types.YLeaf{"AlarmType", entPhyIndex.AlarmType}
    entPhyIndex.EntityData.Leafs["data-type"] = types.YLeaf{"DataType", entPhyIndex.DataType}
    entPhyIndex.EntityData.Leafs["scale"] = types.YLeaf{"Scale", entPhyIndex.Scale}
    entPhyIndex.EntityData.Leafs["precision"] = types.YLeaf{"Precision", entPhyIndex.Precision}
    entPhyIndex.EntityData.Leafs["status"] = types.YLeaf{"Status", entPhyIndex.Status}
    entPhyIndex.EntityData.Leafs["age-time-stamp"] = types.YLeaf{"AgeTimeStamp", entPhyIndex.AgeTimeStamp}
    entPhyIndex.EntityData.Leafs["update-rate"] = types.YLeaf{"UpdateRate", entPhyIndex.UpdateRate}
    entPhyIndex.EntityData.Leafs["measured-entity"] = types.YLeaf{"MeasuredEntity", entPhyIndex.MeasuredEntity}
    return &(entPhyIndex.EntityData)
}

