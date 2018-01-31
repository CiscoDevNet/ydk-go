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
    parent types.Entity
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

func (snmp *Snmp) GetFilter() yfilter.YFilter { return snmp.YFilter }

func (snmp *Snmp) SetFilter(yf yfilter.YFilter) { snmp.YFilter = yf }

func (snmp *Snmp) GetGoName(yname string) string {
    if yname == "trap-servers" { return "TrapServers" }
    if yname == "information" { return "Information" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "correlator" { return "Correlator" }
    if yname == "interface-indexes" { return "InterfaceIndexes" }
    if yname == "if-indexes" { return "IfIndexes" }
    if yname == "Cisco-IOS-XR-snmp-entitymib-oper:entity-mib" { return "EntityMib" }
    if yname == "Cisco-IOS-XR-snmp-ifmib-oper:interface-mib" { return "InterfaceMib" }
    if yname == "Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib" { return "SensorMib" }
    return ""
}

func (snmp *Snmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-agent-oper:snmp"
}

func (snmp *Snmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-servers" {
        return &snmp.TrapServers
    }
    if childYangName == "information" {
        return &snmp.Information
    }
    if childYangName == "interfaces" {
        return &snmp.Interfaces
    }
    if childYangName == "correlator" {
        return &snmp.Correlator
    }
    if childYangName == "interface-indexes" {
        return &snmp.InterfaceIndexes
    }
    if childYangName == "if-indexes" {
        return &snmp.IfIndexes
    }
    if childYangName == "Cisco-IOS-XR-snmp-entitymib-oper:entity-mib" {
        return &snmp.EntityMib
    }
    if childYangName == "Cisco-IOS-XR-snmp-ifmib-oper:interface-mib" {
        return &snmp.InterfaceMib
    }
    if childYangName == "Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib" {
        return &snmp.SensorMib
    }
    return nil
}

func (snmp *Snmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["trap-servers"] = &snmp.TrapServers
    children["information"] = &snmp.Information
    children["interfaces"] = &snmp.Interfaces
    children["correlator"] = &snmp.Correlator
    children["interface-indexes"] = &snmp.InterfaceIndexes
    children["if-indexes"] = &snmp.IfIndexes
    children["Cisco-IOS-XR-snmp-entitymib-oper:entity-mib"] = &snmp.EntityMib
    children["Cisco-IOS-XR-snmp-ifmib-oper:interface-mib"] = &snmp.InterfaceMib
    children["Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib"] = &snmp.SensorMib
    return children
}

func (snmp *Snmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snmp *Snmp) GetBundleName() string { return "cisco_ios_xr" }

func (snmp *Snmp) GetYangName() string { return "snmp" }

func (snmp *Snmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snmp *Snmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snmp *Snmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snmp *Snmp) SetParent(parent types.Entity) { snmp.parent = parent }

func (snmp *Snmp) GetParent() types.Entity { return snmp.parent }

func (snmp *Snmp) GetParentYangName() string { return "Cisco-IOS-XR-snmp-agent-oper" }

// Snmp_TrapServers
// List of trap hosts
type Snmp_TrapServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trap server and port to which the trap is to be sent and statistics. The
    // type is slice of Snmp_TrapServers_TrapServer.
    TrapServer []Snmp_TrapServers_TrapServer
}

func (trapServers *Snmp_TrapServers) GetFilter() yfilter.YFilter { return trapServers.YFilter }

func (trapServers *Snmp_TrapServers) SetFilter(yf yfilter.YFilter) { trapServers.YFilter = yf }

func (trapServers *Snmp_TrapServers) GetGoName(yname string) string {
    if yname == "trap-server" { return "TrapServer" }
    return ""
}

func (trapServers *Snmp_TrapServers) GetSegmentPath() string {
    return "trap-servers"
}

func (trapServers *Snmp_TrapServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-server" {
        for _, c := range trapServers.TrapServer {
            if trapServers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_TrapServers_TrapServer{}
        trapServers.TrapServer = append(trapServers.TrapServer, child)
        return &trapServers.TrapServer[len(trapServers.TrapServer)-1]
    }
    return nil
}

func (trapServers *Snmp_TrapServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapServers.TrapServer {
        children[trapServers.TrapServer[i].GetSegmentPath()] = &trapServers.TrapServer[i]
    }
    return children
}

func (trapServers *Snmp_TrapServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trapServers *Snmp_TrapServers) GetBundleName() string { return "cisco_ios_xr" }

func (trapServers *Snmp_TrapServers) GetYangName() string { return "trap-servers" }

func (trapServers *Snmp_TrapServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapServers *Snmp_TrapServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapServers *Snmp_TrapServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapServers *Snmp_TrapServers) SetParent(parent types.Entity) { trapServers.parent = parent }

func (trapServers *Snmp_TrapServers) GetParent() types.Entity { return trapServers.parent }

func (trapServers *Snmp_TrapServers) GetParentYangName() string { return "snmp" }

// Snmp_TrapServers_TrapServer
// Trap server and port to which the trap is to be
// sent and statistics
type Snmp_TrapServers_TrapServer struct {
    parent types.Entity
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

func (trapServer *Snmp_TrapServers_TrapServer) GetFilter() yfilter.YFilter { return trapServer.YFilter }

func (trapServer *Snmp_TrapServers_TrapServer) SetFilter(yf yfilter.YFilter) { trapServer.YFilter = yf }

func (trapServer *Snmp_TrapServers_TrapServer) GetGoName(yname string) string {
    if yname == "trap-host" { return "TrapHost" }
    if yname == "port" { return "Port" }
    if yname == "number-of-pkts-in-trap-q" { return "NumberOfPktsInTrapQ" }
    if yname == "max-q-length-of-trap-q" { return "MaxQLengthOfTrapQ" }
    if yname == "number-of-pkts-sent" { return "NumberOfPktsSent" }
    if yname == "number-of-pkts-dropped" { return "NumberOfPktsDropped" }
    return ""
}

func (trapServer *Snmp_TrapServers_TrapServer) GetSegmentPath() string {
    return "trap-server"
}

func (trapServer *Snmp_TrapServers_TrapServer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapServer *Snmp_TrapServers_TrapServer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapServer *Snmp_TrapServers_TrapServer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-host"] = trapServer.TrapHost
    leafs["port"] = trapServer.Port
    leafs["number-of-pkts-in-trap-q"] = trapServer.NumberOfPktsInTrapQ
    leafs["max-q-length-of-trap-q"] = trapServer.MaxQLengthOfTrapQ
    leafs["number-of-pkts-sent"] = trapServer.NumberOfPktsSent
    leafs["number-of-pkts-dropped"] = trapServer.NumberOfPktsDropped
    return leafs
}

func (trapServer *Snmp_TrapServers_TrapServer) GetBundleName() string { return "cisco_ios_xr" }

func (trapServer *Snmp_TrapServers_TrapServer) GetYangName() string { return "trap-server" }

func (trapServer *Snmp_TrapServers_TrapServer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapServer *Snmp_TrapServers_TrapServer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapServer *Snmp_TrapServers_TrapServer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapServer *Snmp_TrapServers_TrapServer) SetParent(parent types.Entity) { trapServer.parent = parent }

func (trapServer *Snmp_TrapServers_TrapServer) GetParent() types.Entity { return trapServer.parent }

func (trapServer *Snmp_TrapServers_TrapServer) GetParentYangName() string { return "trap-servers" }

// Snmp_Information
// SNMP operational information
type Snmp_Information struct {
    parent types.Entity
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

func (information *Snmp_Information) GetFilter() yfilter.YFilter { return information.YFilter }

func (information *Snmp_Information) SetFilter(yf yfilter.YFilter) { information.YFilter = yf }

func (information *Snmp_Information) GetGoName(yname string) string {
    if yname == "hosts" { return "Hosts" }
    if yname == "system-up-time" { return "SystemUpTime" }
    if yname == "nms-addresses" { return "NmsAddresses" }
    if yname == "engine-id" { return "EngineId" }
    if yname == "rx-queue" { return "RxQueue" }
    if yname == "system-name" { return "SystemName" }
    if yname == "request-type-detail" { return "RequestTypeDetail" }
    if yname == "duplicate-drop" { return "DuplicateDrop" }
    if yname == "bulk-stats-transfers" { return "BulkStatsTransfers" }
    if yname == "trap-infos" { return "TrapInfos" }
    if yname == "poll-oids" { return "PollOids" }
    if yname == "infom-details" { return "InfomDetails" }
    if yname == "statistics" { return "Statistics" }
    if yname == "incoming-queue" { return "IncomingQueue" }
    if yname == "context-mapping" { return "ContextMapping" }
    if yname == "trap-oids" { return "TrapOids" }
    if yname == "nm-spackets" { return "NmSpackets" }
    if yname == "mibs" { return "Mibs" }
    if yname == "serial-numbers" { return "SerialNumbers" }
    if yname == "drop-nms-addresses" { return "DropNmsAddresses" }
    if yname == "views" { return "Views" }
    if yname == "system-descr" { return "SystemDescr" }
    if yname == "tables" { return "Tables" }
    if yname == "system-oid" { return "SystemOid" }
    if yname == "trap-queue" { return "TrapQueue" }
    return ""
}

func (information *Snmp_Information) GetSegmentPath() string {
    return "information"
}

func (information *Snmp_Information) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hosts" {
        return &information.Hosts
    }
    if childYangName == "system-up-time" {
        return &information.SystemUpTime
    }
    if childYangName == "nms-addresses" {
        return &information.NmsAddresses
    }
    if childYangName == "engine-id" {
        return &information.EngineId
    }
    if childYangName == "rx-queue" {
        return &information.RxQueue
    }
    if childYangName == "system-name" {
        return &information.SystemName
    }
    if childYangName == "request-type-detail" {
        return &information.RequestTypeDetail
    }
    if childYangName == "duplicate-drop" {
        return &information.DuplicateDrop
    }
    if childYangName == "bulk-stats-transfers" {
        return &information.BulkStatsTransfers
    }
    if childYangName == "trap-infos" {
        return &information.TrapInfos
    }
    if childYangName == "poll-oids" {
        return &information.PollOids
    }
    if childYangName == "infom-details" {
        return &information.InfomDetails
    }
    if childYangName == "statistics" {
        return &information.Statistics
    }
    if childYangName == "incoming-queue" {
        return &information.IncomingQueue
    }
    if childYangName == "context-mapping" {
        return &information.ContextMapping
    }
    if childYangName == "trap-oids" {
        return &information.TrapOids
    }
    if childYangName == "nm-spackets" {
        return &information.NmSpackets
    }
    if childYangName == "mibs" {
        return &information.Mibs
    }
    if childYangName == "serial-numbers" {
        return &information.SerialNumbers
    }
    if childYangName == "drop-nms-addresses" {
        return &information.DropNmsAddresses
    }
    if childYangName == "views" {
        return &information.Views
    }
    if childYangName == "system-descr" {
        return &information.SystemDescr
    }
    if childYangName == "tables" {
        return &information.Tables
    }
    if childYangName == "system-oid" {
        return &information.SystemOid
    }
    if childYangName == "trap-queue" {
        return &information.TrapQueue
    }
    return nil
}

func (information *Snmp_Information) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hosts"] = &information.Hosts
    children["system-up-time"] = &information.SystemUpTime
    children["nms-addresses"] = &information.NmsAddresses
    children["engine-id"] = &information.EngineId
    children["rx-queue"] = &information.RxQueue
    children["system-name"] = &information.SystemName
    children["request-type-detail"] = &information.RequestTypeDetail
    children["duplicate-drop"] = &information.DuplicateDrop
    children["bulk-stats-transfers"] = &information.BulkStatsTransfers
    children["trap-infos"] = &information.TrapInfos
    children["poll-oids"] = &information.PollOids
    children["infom-details"] = &information.InfomDetails
    children["statistics"] = &information.Statistics
    children["incoming-queue"] = &information.IncomingQueue
    children["context-mapping"] = &information.ContextMapping
    children["trap-oids"] = &information.TrapOids
    children["nm-spackets"] = &information.NmSpackets
    children["mibs"] = &information.Mibs
    children["serial-numbers"] = &information.SerialNumbers
    children["drop-nms-addresses"] = &information.DropNmsAddresses
    children["views"] = &information.Views
    children["system-descr"] = &information.SystemDescr
    children["tables"] = &information.Tables
    children["system-oid"] = &information.SystemOid
    children["trap-queue"] = &information.TrapQueue
    return children
}

func (information *Snmp_Information) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (information *Snmp_Information) GetBundleName() string { return "cisco_ios_xr" }

func (information *Snmp_Information) GetYangName() string { return "information" }

func (information *Snmp_Information) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (information *Snmp_Information) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (information *Snmp_Information) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (information *Snmp_Information) SetParent(parent types.Entity) { information.parent = parent }

func (information *Snmp_Information) GetParent() types.Entity { return information.parent }

func (information *Snmp_Information) GetParentYangName() string { return "snmp" }

// Snmp_Information_Hosts
// SNMP host information
type Snmp_Information_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP target host name. The type is slice of Snmp_Information_Hosts_Host.
    Host []Snmp_Information_Hosts_Host
}

func (hosts *Snmp_Information_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *Snmp_Information_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *Snmp_Information_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *Snmp_Information_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *Snmp_Information_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *Snmp_Information_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *Snmp_Information_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *Snmp_Information_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *Snmp_Information_Hosts) GetYangName() string { return "hosts" }

func (hosts *Snmp_Information_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *Snmp_Information_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *Snmp_Information_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *Snmp_Information_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *Snmp_Information_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *Snmp_Information_Hosts) GetParentYangName() string { return "information" }

// Snmp_Information_Hosts_Host
// SNMP target host name
type Snmp_Information_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Host name ,udp-port , user, security model and level. The type is slice of
    // Snmp_Information_Hosts_Host_HostInformation.
    HostInformation []Snmp_Information_Hosts_Host_HostInformation
}

func (host *Snmp_Information_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *Snmp_Information_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *Snmp_Information_Hosts_Host) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "host-information" { return "HostInformation" }
    return ""
}

func (host *Snmp_Information_Hosts_Host) GetSegmentPath() string {
    return "host" + "[name='" + fmt.Sprintf("%v", host.Name) + "']"
}

func (host *Snmp_Information_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host-information" {
        for _, c := range host.HostInformation {
            if host.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Hosts_Host_HostInformation{}
        host.HostInformation = append(host.HostInformation, child)
        return &host.HostInformation[len(host.HostInformation)-1]
    }
    return nil
}

func (host *Snmp_Information_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range host.HostInformation {
        children[host.HostInformation[i].GetSegmentPath()] = &host.HostInformation[i]
    }
    return children
}

func (host *Snmp_Information_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = host.Name
    return leafs
}

func (host *Snmp_Information_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *Snmp_Information_Hosts_Host) GetYangName() string { return "host" }

func (host *Snmp_Information_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *Snmp_Information_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *Snmp_Information_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *Snmp_Information_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *Snmp_Information_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *Snmp_Information_Hosts_Host) GetParentYangName() string { return "hosts" }

// Snmp_Information_Hosts_Host_HostInformation
// Host name ,udp-port , user, security model
// and level
type Snmp_Information_Hosts_Host_HostInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetFilter() yfilter.YFilter { return hostInformation.YFilter }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) SetFilter(yf yfilter.YFilter) { hostInformation.YFilter = yf }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetGoName(yname string) string {
    if yname == "user" { return "User" }
    if yname == "snmp-target-address-t-host" { return "SnmpTargetAddressTHost" }
    if yname == "snmp-target-address-port" { return "SnmpTargetAddressPort" }
    if yname == "snmp-target-addresstype" { return "SnmpTargetAddresstype" }
    if yname == "snmp-target-params-security-model" { return "SnmpTargetParamsSecurityModel" }
    if yname == "snmp-target-params-security-name" { return "SnmpTargetParamsSecurityName" }
    if yname == "snmp-target-params-security-level" { return "SnmpTargetParamsSecurityLevel" }
    return ""
}

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetSegmentPath() string {
    return "host-information" + "[user='" + fmt.Sprintf("%v", hostInformation.User) + "']"
}

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["user"] = hostInformation.User
    leafs["snmp-target-address-t-host"] = hostInformation.SnmpTargetAddressTHost
    leafs["snmp-target-address-port"] = hostInformation.SnmpTargetAddressPort
    leafs["snmp-target-addresstype"] = hostInformation.SnmpTargetAddresstype
    leafs["snmp-target-params-security-model"] = hostInformation.SnmpTargetParamsSecurityModel
    leafs["snmp-target-params-security-name"] = hostInformation.SnmpTargetParamsSecurityName
    leafs["snmp-target-params-security-level"] = hostInformation.SnmpTargetParamsSecurityLevel
    return leafs
}

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetBundleName() string { return "cisco_ios_xr" }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetYangName() string { return "host-information" }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) SetParent(parent types.Entity) { hostInformation.parent = parent }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetParent() types.Entity { return hostInformation.parent }

func (hostInformation *Snmp_Information_Hosts_Host_HostInformation) GetParentYangName() string { return "host" }

// Snmp_Information_SystemUpTime
// System up time
type Snmp_Information_SystemUpTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // sysUpTime  1.3.6.1.2.1.1.3. The type is string.
    SystemUpTimeEdm interface{}
}

func (systemUpTime *Snmp_Information_SystemUpTime) GetFilter() yfilter.YFilter { return systemUpTime.YFilter }

func (systemUpTime *Snmp_Information_SystemUpTime) SetFilter(yf yfilter.YFilter) { systemUpTime.YFilter = yf }

func (systemUpTime *Snmp_Information_SystemUpTime) GetGoName(yname string) string {
    if yname == "system-up-time-edm" { return "SystemUpTimeEdm" }
    return ""
}

func (systemUpTime *Snmp_Information_SystemUpTime) GetSegmentPath() string {
    return "system-up-time"
}

func (systemUpTime *Snmp_Information_SystemUpTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemUpTime *Snmp_Information_SystemUpTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemUpTime *Snmp_Information_SystemUpTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-up-time-edm"] = systemUpTime.SystemUpTimeEdm
    return leafs
}

func (systemUpTime *Snmp_Information_SystemUpTime) GetBundleName() string { return "cisco_ios_xr" }

func (systemUpTime *Snmp_Information_SystemUpTime) GetYangName() string { return "system-up-time" }

func (systemUpTime *Snmp_Information_SystemUpTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemUpTime *Snmp_Information_SystemUpTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemUpTime *Snmp_Information_SystemUpTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemUpTime *Snmp_Information_SystemUpTime) SetParent(parent types.Entity) { systemUpTime.parent = parent }

func (systemUpTime *Snmp_Information_SystemUpTime) GetParent() types.Entity { return systemUpTime.parent }

func (systemUpTime *Snmp_Information_SystemUpTime) GetParentYangName() string { return "information" }

// Snmp_Information_NmsAddresses
// SNMP request type summary 
type Snmp_Information_NmsAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NMS address. The type is slice of Snmp_Information_NmsAddresses_NmsAddress.
    NmsAddress []Snmp_Information_NmsAddresses_NmsAddress
}

func (nmsAddresses *Snmp_Information_NmsAddresses) GetFilter() yfilter.YFilter { return nmsAddresses.YFilter }

func (nmsAddresses *Snmp_Information_NmsAddresses) SetFilter(yf yfilter.YFilter) { nmsAddresses.YFilter = yf }

func (nmsAddresses *Snmp_Information_NmsAddresses) GetGoName(yname string) string {
    if yname == "nms-address" { return "NmsAddress" }
    return ""
}

func (nmsAddresses *Snmp_Information_NmsAddresses) GetSegmentPath() string {
    return "nms-addresses"
}

func (nmsAddresses *Snmp_Information_NmsAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nms-address" {
        for _, c := range nmsAddresses.NmsAddress {
            if nmsAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_NmsAddresses_NmsAddress{}
        nmsAddresses.NmsAddress = append(nmsAddresses.NmsAddress, child)
        return &nmsAddresses.NmsAddress[len(nmsAddresses.NmsAddress)-1]
    }
    return nil
}

func (nmsAddresses *Snmp_Information_NmsAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nmsAddresses.NmsAddress {
        children[nmsAddresses.NmsAddress[i].GetSegmentPath()] = &nmsAddresses.NmsAddress[i]
    }
    return children
}

func (nmsAddresses *Snmp_Information_NmsAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nmsAddresses *Snmp_Information_NmsAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (nmsAddresses *Snmp_Information_NmsAddresses) GetYangName() string { return "nms-addresses" }

func (nmsAddresses *Snmp_Information_NmsAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nmsAddresses *Snmp_Information_NmsAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nmsAddresses *Snmp_Information_NmsAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nmsAddresses *Snmp_Information_NmsAddresses) SetParent(parent types.Entity) { nmsAddresses.parent = parent }

func (nmsAddresses *Snmp_Information_NmsAddresses) GetParent() types.Entity { return nmsAddresses.parent }

func (nmsAddresses *Snmp_Information_NmsAddresses) GetParentYangName() string { return "information" }

// Snmp_Information_NmsAddresses_NmsAddress
// NMS address
type Snmp_Information_NmsAddresses_NmsAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetFilter() yfilter.YFilter { return nmsAddress.YFilter }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) SetFilter(yf yfilter.YFilter) { nmsAddress.YFilter = yf }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetGoName(yname string) string {
    if yname == "nms-addr" { return "NmsAddr" }
    if yname == "nms-address" { return "NmsAddress" }
    if yname == "get-request-count" { return "GetRequestCount" }
    if yname == "getnext-request-count" { return "GetnextRequestCount" }
    if yname == "getbulk-request-count" { return "GetbulkRequestCount" }
    if yname == "set-request-count" { return "SetRequestCount" }
    if yname == "test-request-count" { return "TestRequestCount" }
    return ""
}

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetSegmentPath() string {
    return "nms-address" + "[nms-addr='" + fmt.Sprintf("%v", nmsAddress.NmsAddr) + "']"
}

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nms-addr"] = nmsAddress.NmsAddr
    leafs["nms-address"] = nmsAddress.NmsAddress
    leafs["get-request-count"] = nmsAddress.GetRequestCount
    leafs["getnext-request-count"] = nmsAddress.GetnextRequestCount
    leafs["getbulk-request-count"] = nmsAddress.GetbulkRequestCount
    leafs["set-request-count"] = nmsAddress.SetRequestCount
    leafs["test-request-count"] = nmsAddress.TestRequestCount
    return leafs
}

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetBundleName() string { return "cisco_ios_xr" }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetYangName() string { return "nms-address" }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) SetParent(parent types.Entity) { nmsAddress.parent = parent }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetParent() types.Entity { return nmsAddress.parent }

func (nmsAddress *Snmp_Information_NmsAddresses_NmsAddress) GetParentYangName() string { return "nms-addresses" }

// Snmp_Information_EngineId
// SNMP engine ID
type Snmp_Information_EngineId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMPv3 engineID. The type is string.
    EngineId interface{}
}

func (engineId *Snmp_Information_EngineId) GetFilter() yfilter.YFilter { return engineId.YFilter }

func (engineId *Snmp_Information_EngineId) SetFilter(yf yfilter.YFilter) { engineId.YFilter = yf }

func (engineId *Snmp_Information_EngineId) GetGoName(yname string) string {
    if yname == "engine-id" { return "EngineId" }
    return ""
}

func (engineId *Snmp_Information_EngineId) GetSegmentPath() string {
    return "engine-id"
}

func (engineId *Snmp_Information_EngineId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (engineId *Snmp_Information_EngineId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (engineId *Snmp_Information_EngineId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["engine-id"] = engineId.EngineId
    return leafs
}

func (engineId *Snmp_Information_EngineId) GetBundleName() string { return "cisco_ios_xr" }

func (engineId *Snmp_Information_EngineId) GetYangName() string { return "engine-id" }

func (engineId *Snmp_Information_EngineId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (engineId *Snmp_Information_EngineId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (engineId *Snmp_Information_EngineId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (engineId *Snmp_Information_EngineId) SetParent(parent types.Entity) { engineId.parent = parent }

func (engineId *Snmp_Information_EngineId) GetParent() types.Entity { return engineId.parent }

func (engineId *Snmp_Information_EngineId) GetParentYangName() string { return "information" }

// Snmp_Information_RxQueue
// SNMP rx queue statistics
type Snmp_Information_RxQueue struct {
    parent types.Entity
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

func (rxQueue *Snmp_Information_RxQueue) GetFilter() yfilter.YFilter { return rxQueue.YFilter }

func (rxQueue *Snmp_Information_RxQueue) SetFilter(yf yfilter.YFilter) { rxQueue.YFilter = yf }

func (rxQueue *Snmp_Information_RxQueue) GetGoName(yname string) string {
    if yname == "qlen" { return "Qlen" }
    if yname == "in-min" { return "InMin" }
    if yname == "in-avg" { return "InAvg" }
    if yname == "in-max" { return "InMax" }
    if yname == "pend-min" { return "PendMin" }
    if yname == "pend-avg" { return "PendAvg" }
    if yname == "pend-max" { return "PendMax" }
    if yname == "incoming-q" { return "IncomingQ" }
    if yname == "pending-q" { return "PendingQ" }
    return ""
}

func (rxQueue *Snmp_Information_RxQueue) GetSegmentPath() string {
    return "rx-queue"
}

func (rxQueue *Snmp_Information_RxQueue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "incoming-q" {
        for _, c := range rxQueue.IncomingQ {
            if rxQueue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_RxQueue_IncomingQ{}
        rxQueue.IncomingQ = append(rxQueue.IncomingQ, child)
        return &rxQueue.IncomingQ[len(rxQueue.IncomingQ)-1]
    }
    if childYangName == "pending-q" {
        for _, c := range rxQueue.PendingQ {
            if rxQueue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_RxQueue_PendingQ{}
        rxQueue.PendingQ = append(rxQueue.PendingQ, child)
        return &rxQueue.PendingQ[len(rxQueue.PendingQ)-1]
    }
    return nil
}

func (rxQueue *Snmp_Information_RxQueue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rxQueue.IncomingQ {
        children[rxQueue.IncomingQ[i].GetSegmentPath()] = &rxQueue.IncomingQ[i]
    }
    for i := range rxQueue.PendingQ {
        children[rxQueue.PendingQ[i].GetSegmentPath()] = &rxQueue.PendingQ[i]
    }
    return children
}

func (rxQueue *Snmp_Information_RxQueue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["qlen"] = rxQueue.Qlen
    leafs["in-min"] = rxQueue.InMin
    leafs["in-avg"] = rxQueue.InAvg
    leafs["in-max"] = rxQueue.InMax
    leafs["pend-min"] = rxQueue.PendMin
    leafs["pend-avg"] = rxQueue.PendAvg
    leafs["pend-max"] = rxQueue.PendMax
    return leafs
}

func (rxQueue *Snmp_Information_RxQueue) GetBundleName() string { return "cisco_ios_xr" }

func (rxQueue *Snmp_Information_RxQueue) GetYangName() string { return "rx-queue" }

func (rxQueue *Snmp_Information_RxQueue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxQueue *Snmp_Information_RxQueue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxQueue *Snmp_Information_RxQueue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxQueue *Snmp_Information_RxQueue) SetParent(parent types.Entity) { rxQueue.parent = parent }

func (rxQueue *Snmp_Information_RxQueue) GetParent() types.Entity { return rxQueue.parent }

func (rxQueue *Snmp_Information_RxQueue) GetParentYangName() string { return "information" }

// Snmp_Information_RxQueue_IncomingQ
// incoming q
type Snmp_Information_RxQueue_IncomingQ struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // min. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // avg. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // max. The type is interface{} with range: 0..4294967295.
    Max interface{}
}

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetFilter() yfilter.YFilter { return incomingQ.YFilter }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) SetFilter(yf yfilter.YFilter) { incomingQ.YFilter = yf }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "avg" { return "Avg" }
    if yname == "max" { return "Max" }
    return ""
}

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetSegmentPath() string {
    return "incoming-q"
}

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = incomingQ.Min
    leafs["avg"] = incomingQ.Avg
    leafs["max"] = incomingQ.Max
    return leafs
}

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetBundleName() string { return "cisco_ios_xr" }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetYangName() string { return "incoming-q" }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) SetParent(parent types.Entity) { incomingQ.parent = parent }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetParent() types.Entity { return incomingQ.parent }

func (incomingQ *Snmp_Information_RxQueue_IncomingQ) GetParentYangName() string { return "rx-queue" }

// Snmp_Information_RxQueue_PendingQ
// pending q
type Snmp_Information_RxQueue_PendingQ struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // min. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // avg. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // max. The type is interface{} with range: 0..4294967295.
    Max interface{}
}

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetFilter() yfilter.YFilter { return pendingQ.YFilter }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) SetFilter(yf yfilter.YFilter) { pendingQ.YFilter = yf }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "avg" { return "Avg" }
    if yname == "max" { return "Max" }
    return ""
}

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetSegmentPath() string {
    return "pending-q"
}

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = pendingQ.Min
    leafs["avg"] = pendingQ.Avg
    leafs["max"] = pendingQ.Max
    return leafs
}

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetBundleName() string { return "cisco_ios_xr" }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetYangName() string { return "pending-q" }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) SetParent(parent types.Entity) { pendingQ.parent = parent }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetParent() types.Entity { return pendingQ.parent }

func (pendingQ *Snmp_Information_RxQueue_PendingQ) GetParentYangName() string { return "rx-queue" }

// Snmp_Information_SystemName
// System name
type Snmp_Information_SystemName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // sysName  1.3.6.1.2.1.1.5. The type is string.
    SystemName interface{}
}

func (systemName *Snmp_Information_SystemName) GetFilter() yfilter.YFilter { return systemName.YFilter }

func (systemName *Snmp_Information_SystemName) SetFilter(yf yfilter.YFilter) { systemName.YFilter = yf }

func (systemName *Snmp_Information_SystemName) GetGoName(yname string) string {
    if yname == "system-name" { return "SystemName" }
    return ""
}

func (systemName *Snmp_Information_SystemName) GetSegmentPath() string {
    return "system-name"
}

func (systemName *Snmp_Information_SystemName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemName *Snmp_Information_SystemName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemName *Snmp_Information_SystemName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-name"] = systemName.SystemName
    return leafs
}

func (systemName *Snmp_Information_SystemName) GetBundleName() string { return "cisco_ios_xr" }

func (systemName *Snmp_Information_SystemName) GetYangName() string { return "system-name" }

func (systemName *Snmp_Information_SystemName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemName *Snmp_Information_SystemName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemName *Snmp_Information_SystemName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemName *Snmp_Information_SystemName) SetParent(parent types.Entity) { systemName.parent = parent }

func (systemName *Snmp_Information_SystemName) GetParent() types.Entity { return systemName.parent }

func (systemName *Snmp_Information_SystemName) GetParentYangName() string { return "information" }

// Snmp_Information_RequestTypeDetail
// SNMP request type details 
type Snmp_Information_RequestTypeDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // snmp request type details .
    NmsAddresses Snmp_Information_RequestTypeDetail_NmsAddresses
}

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetFilter() yfilter.YFilter { return requestTypeDetail.YFilter }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) SetFilter(yf yfilter.YFilter) { requestTypeDetail.YFilter = yf }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetGoName(yname string) string {
    if yname == "nms-addresses" { return "NmsAddresses" }
    return ""
}

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetSegmentPath() string {
    return "request-type-detail"
}

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nms-addresses" {
        return &requestTypeDetail.NmsAddresses
    }
    return nil
}

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nms-addresses"] = &requestTypeDetail.NmsAddresses
    return children
}

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetBundleName() string { return "cisco_ios_xr" }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetYangName() string { return "request-type-detail" }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) SetParent(parent types.Entity) { requestTypeDetail.parent = parent }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetParent() types.Entity { return requestTypeDetail.parent }

func (requestTypeDetail *Snmp_Information_RequestTypeDetail) GetParentYangName() string { return "information" }

// Snmp_Information_RequestTypeDetail_NmsAddresses
// snmp request type details 
type Snmp_Information_RequestTypeDetail_NmsAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NMS address. The type is slice of
    // Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress.
    NmsAddress []Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress
}

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetFilter() yfilter.YFilter { return nmsAddresses.YFilter }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) SetFilter(yf yfilter.YFilter) { nmsAddresses.YFilter = yf }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetGoName(yname string) string {
    if yname == "nms-address" { return "NmsAddress" }
    return ""
}

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetSegmentPath() string {
    return "nms-addresses"
}

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nms-address" {
        for _, c := range nmsAddresses.NmsAddress {
            if nmsAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress{}
        nmsAddresses.NmsAddress = append(nmsAddresses.NmsAddress, child)
        return &nmsAddresses.NmsAddress[len(nmsAddresses.NmsAddress)-1]
    }
    return nil
}

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nmsAddresses.NmsAddress {
        children[nmsAddresses.NmsAddress[i].GetSegmentPath()] = &nmsAddresses.NmsAddress[i]
    }
    return children
}

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetYangName() string { return "nms-addresses" }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) SetParent(parent types.Entity) { nmsAddresses.parent = parent }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetParent() types.Entity { return nmsAddresses.parent }

func (nmsAddresses *Snmp_Information_RequestTypeDetail_NmsAddresses) GetParentYangName() string { return "request-type-detail" }

// Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress
// NMS address
type Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetFilter() yfilter.YFilter { return nmsAddress.YFilter }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) SetFilter(yf yfilter.YFilter) { nmsAddress.YFilter = yf }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetGoName(yname string) string {
    if yname == "nms-addr" { return "NmsAddr" }
    if yname == "total-count" { return "TotalCount" }
    if yname == "agent-request-count" { return "AgentRequestCount" }
    if yname == "interface-request-count" { return "InterfaceRequestCount" }
    if yname == "entity-request-count" { return "EntityRequestCount" }
    if yname == "route-request-count" { return "RouteRequestCount" }
    if yname == "infra-request-count" { return "InfraRequestCount" }
    return ""
}

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetSegmentPath() string {
    return "nms-address" + "[nms-addr='" + fmt.Sprintf("%v", nmsAddress.NmsAddr) + "']"
}

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nms-addr"] = nmsAddress.NmsAddr
    leafs["total-count"] = nmsAddress.TotalCount
    leafs["agent-request-count"] = nmsAddress.AgentRequestCount
    leafs["interface-request-count"] = nmsAddress.InterfaceRequestCount
    leafs["entity-request-count"] = nmsAddress.EntityRequestCount
    leafs["route-request-count"] = nmsAddress.RouteRequestCount
    leafs["infra-request-count"] = nmsAddress.InfraRequestCount
    return leafs
}

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetBundleName() string { return "cisco_ios_xr" }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetYangName() string { return "nms-address" }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) SetParent(parent types.Entity) { nmsAddress.parent = parent }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetParent() types.Entity { return nmsAddress.parent }

func (nmsAddress *Snmp_Information_RequestTypeDetail_NmsAddresses_NmsAddress) GetParentYangName() string { return "nms-addresses" }

// Snmp_Information_DuplicateDrop
// Duplicate request status, count, time 
type Snmp_Information_DuplicateDrop struct {
    parent types.Entity
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

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetFilter() yfilter.YFilter { return duplicateDrop.YFilter }

func (duplicateDrop *Snmp_Information_DuplicateDrop) SetFilter(yf yfilter.YFilter) { duplicateDrop.YFilter = yf }

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetGoName(yname string) string {
    if yname == "duplicate-request-status" { return "DuplicateRequestStatus" }
    if yname == "last-status-change-time" { return "LastStatusChangeTime" }
    if yname == "duplicate-drop-configured-timeout" { return "DuplicateDropConfiguredTimeout" }
    if yname == "duplicate-dropped-requests" { return "DuplicateDroppedRequests" }
    if yname == "retry-processed-requests" { return "RetryProcessedRequests" }
    if yname == "first-enable-time" { return "FirstEnableTime" }
    if yname == "latest-duplicate-dropped-requests" { return "LatestDuplicateDroppedRequests" }
    if yname == "latest-retry-processed-requests" { return "LatestRetryProcessedRequests" }
    if yname == "duplicate-request-latest-enable-time" { return "DuplicateRequestLatestEnableTime" }
    if yname == "duplicate-drop-enable-count" { return "DuplicateDropEnableCount" }
    if yname == "duplicate-drop-disable-count" { return "DuplicateDropDisableCount" }
    return ""
}

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetSegmentPath() string {
    return "duplicate-drop"
}

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["duplicate-request-status"] = duplicateDrop.DuplicateRequestStatus
    leafs["last-status-change-time"] = duplicateDrop.LastStatusChangeTime
    leafs["duplicate-drop-configured-timeout"] = duplicateDrop.DuplicateDropConfiguredTimeout
    leafs["duplicate-dropped-requests"] = duplicateDrop.DuplicateDroppedRequests
    leafs["retry-processed-requests"] = duplicateDrop.RetryProcessedRequests
    leafs["first-enable-time"] = duplicateDrop.FirstEnableTime
    leafs["latest-duplicate-dropped-requests"] = duplicateDrop.LatestDuplicateDroppedRequests
    leafs["latest-retry-processed-requests"] = duplicateDrop.LatestRetryProcessedRequests
    leafs["duplicate-request-latest-enable-time"] = duplicateDrop.DuplicateRequestLatestEnableTime
    leafs["duplicate-drop-enable-count"] = duplicateDrop.DuplicateDropEnableCount
    leafs["duplicate-drop-disable-count"] = duplicateDrop.DuplicateDropDisableCount
    return leafs
}

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetBundleName() string { return "cisco_ios_xr" }

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetYangName() string { return "duplicate-drop" }

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (duplicateDrop *Snmp_Information_DuplicateDrop) SetParent(parent types.Entity) { duplicateDrop.parent = parent }

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetParent() types.Entity { return duplicateDrop.parent }

func (duplicateDrop *Snmp_Information_DuplicateDrop) GetParentYangName() string { return "information" }

// Snmp_Information_BulkStatsTransfers
// List of bulkstats transfer on the system
type Snmp_Information_BulkStatsTransfers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP bulkstats transfer name. The type is slice of
    // Snmp_Information_BulkStatsTransfers_BulkStatsTransfer.
    BulkStatsTransfer []Snmp_Information_BulkStatsTransfers_BulkStatsTransfer
}

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetFilter() yfilter.YFilter { return bulkStatsTransfers.YFilter }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) SetFilter(yf yfilter.YFilter) { bulkStatsTransfers.YFilter = yf }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetGoName(yname string) string {
    if yname == "bulk-stats-transfer" { return "BulkStatsTransfer" }
    return ""
}

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetSegmentPath() string {
    return "bulk-stats-transfers"
}

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bulk-stats-transfer" {
        for _, c := range bulkStatsTransfers.BulkStatsTransfer {
            if bulkStatsTransfers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_BulkStatsTransfers_BulkStatsTransfer{}
        bulkStatsTransfers.BulkStatsTransfer = append(bulkStatsTransfers.BulkStatsTransfer, child)
        return &bulkStatsTransfers.BulkStatsTransfer[len(bulkStatsTransfers.BulkStatsTransfer)-1]
    }
    return nil
}

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bulkStatsTransfers.BulkStatsTransfer {
        children[bulkStatsTransfers.BulkStatsTransfer[i].GetSegmentPath()] = &bulkStatsTransfers.BulkStatsTransfer[i]
    }
    return children
}

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetBundleName() string { return "cisco_ios_xr" }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetYangName() string { return "bulk-stats-transfers" }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) SetParent(parent types.Entity) { bulkStatsTransfers.parent = parent }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetParent() types.Entity { return bulkStatsTransfers.parent }

func (bulkStatsTransfers *Snmp_Information_BulkStatsTransfers) GetParentYangName() string { return "information" }

// Snmp_Information_BulkStatsTransfers_BulkStatsTransfer
// SNMP bulkstats transfer name
type Snmp_Information_BulkStatsTransfers_BulkStatsTransfer struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetFilter() yfilter.YFilter { return bulkStatsTransfer.YFilter }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) SetFilter(yf yfilter.YFilter) { bulkStatsTransfer.YFilter = yf }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetGoName(yname string) string {
    if yname == "transfer-name" { return "TransferName" }
    if yname == "transfer-name-xr" { return "TransferNameXr" }
    if yname == "url-primary" { return "UrlPrimary" }
    if yname == "url-secondary" { return "UrlSecondary" }
    if yname == "retained-file" { return "RetainedFile" }
    if yname == "time-left" { return "TimeLeft" }
    if yname == "retry-left" { return "RetryLeft" }
    return ""
}

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetSegmentPath() string {
    return "bulk-stats-transfer" + "[transfer-name='" + fmt.Sprintf("%v", bulkStatsTransfer.TransferName) + "']"
}

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transfer-name"] = bulkStatsTransfer.TransferName
    leafs["transfer-name-xr"] = bulkStatsTransfer.TransferNameXr
    leafs["url-primary"] = bulkStatsTransfer.UrlPrimary
    leafs["url-secondary"] = bulkStatsTransfer.UrlSecondary
    leafs["retained-file"] = bulkStatsTransfer.RetainedFile
    leafs["time-left"] = bulkStatsTransfer.TimeLeft
    leafs["retry-left"] = bulkStatsTransfer.RetryLeft
    return leafs
}

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetBundleName() string { return "cisco_ios_xr" }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetYangName() string { return "bulk-stats-transfer" }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) SetParent(parent types.Entity) { bulkStatsTransfer.parent = parent }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetParent() types.Entity { return bulkStatsTransfer.parent }

func (bulkStatsTransfer *Snmp_Information_BulkStatsTransfers_BulkStatsTransfer) GetParentYangName() string { return "bulk-stats-transfers" }

// Snmp_Information_TrapInfos
// SNMP trap OID
type Snmp_Information_TrapInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP Trap infomation like server , port and trapOID. The type is slice of
    // Snmp_Information_TrapInfos_TrapInfo.
    TrapInfo []Snmp_Information_TrapInfos_TrapInfo
}

func (trapInfos *Snmp_Information_TrapInfos) GetFilter() yfilter.YFilter { return trapInfos.YFilter }

func (trapInfos *Snmp_Information_TrapInfos) SetFilter(yf yfilter.YFilter) { trapInfos.YFilter = yf }

func (trapInfos *Snmp_Information_TrapInfos) GetGoName(yname string) string {
    if yname == "trap-info" { return "TrapInfo" }
    return ""
}

func (trapInfos *Snmp_Information_TrapInfos) GetSegmentPath() string {
    return "trap-infos"
}

func (trapInfos *Snmp_Information_TrapInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-info" {
        for _, c := range trapInfos.TrapInfo {
            if trapInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_TrapInfos_TrapInfo{}
        trapInfos.TrapInfo = append(trapInfos.TrapInfo, child)
        return &trapInfos.TrapInfo[len(trapInfos.TrapInfo)-1]
    }
    return nil
}

func (trapInfos *Snmp_Information_TrapInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapInfos.TrapInfo {
        children[trapInfos.TrapInfo[i].GetSegmentPath()] = &trapInfos.TrapInfo[i]
    }
    return children
}

func (trapInfos *Snmp_Information_TrapInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trapInfos *Snmp_Information_TrapInfos) GetBundleName() string { return "cisco_ios_xr" }

func (trapInfos *Snmp_Information_TrapInfos) GetYangName() string { return "trap-infos" }

func (trapInfos *Snmp_Information_TrapInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapInfos *Snmp_Information_TrapInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapInfos *Snmp_Information_TrapInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapInfos *Snmp_Information_TrapInfos) SetParent(parent types.Entity) { trapInfos.parent = parent }

func (trapInfos *Snmp_Information_TrapInfos) GetParent() types.Entity { return trapInfos.parent }

func (trapInfos *Snmp_Information_TrapInfos) GetParentYangName() string { return "information" }

// Snmp_Information_TrapInfos_TrapInfo
// SNMP Trap infomation like server , port and
// trapOID
type Snmp_Information_TrapInfos_TrapInfo struct {
    parent types.Entity
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

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetFilter() yfilter.YFilter { return trapInfo.YFilter }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) SetFilter(yf yfilter.YFilter) { trapInfo.YFilter = yf }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetGoName(yname string) string {
    if yname == "trap-host" { return "TrapHost" }
    if yname == "port" { return "Port" }
    if yname == "host" { return "Host" }
    if yname == "port-xr" { return "PortXr" }
    if yname == "trap-oid-count" { return "TrapOidCount" }
    if yname == "trap-oi-dinfo" { return "TrapOiDinfo" }
    return ""
}

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetSegmentPath() string {
    return "trap-info"
}

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-oi-dinfo" {
        for _, c := range trapInfo.TrapOiDinfo {
            if trapInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo{}
        trapInfo.TrapOiDinfo = append(trapInfo.TrapOiDinfo, child)
        return &trapInfo.TrapOiDinfo[len(trapInfo.TrapOiDinfo)-1]
    }
    return nil
}

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapInfo.TrapOiDinfo {
        children[trapInfo.TrapOiDinfo[i].GetSegmentPath()] = &trapInfo.TrapOiDinfo[i]
    }
    return children
}

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-host"] = trapInfo.TrapHost
    leafs["port"] = trapInfo.Port
    leafs["host"] = trapInfo.Host
    leafs["port-xr"] = trapInfo.PortXr
    leafs["trap-oid-count"] = trapInfo.TrapOidCount
    return leafs
}

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetYangName() string { return "trap-info" }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) SetParent(parent types.Entity) { trapInfo.parent = parent }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetParent() types.Entity { return trapInfo.parent }

func (trapInfo *Snmp_Information_TrapInfos_TrapInfo) GetParentYangName() string { return "trap-infos" }

// Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo
// Per trap OID statistics
type Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo struct {
    parent types.Entity
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

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetFilter() yfilter.YFilter { return trapOiDinfo.YFilter }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) SetFilter(yf yfilter.YFilter) { trapOiDinfo.YFilter = yf }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetGoName(yname string) string {
    if yname == "trap-oid" { return "TrapOid" }
    if yname == "count" { return "Count" }
    if yname == "drop-count" { return "DropCount" }
    if yname == "retry-count" { return "RetryCount" }
    if yname == "lastsent-time" { return "LastsentTime" }
    if yname == "lasrdrop-time" { return "LasrdropTime" }
    return ""
}

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetSegmentPath() string {
    return "trap-oi-dinfo"
}

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-oid"] = trapOiDinfo.TrapOid
    leafs["count"] = trapOiDinfo.Count
    leafs["drop-count"] = trapOiDinfo.DropCount
    leafs["retry-count"] = trapOiDinfo.RetryCount
    leafs["lastsent-time"] = trapOiDinfo.LastsentTime
    leafs["lasrdrop-time"] = trapOiDinfo.LasrdropTime
    return leafs
}

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetBundleName() string { return "cisco_ios_xr" }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetYangName() string { return "trap-oi-dinfo" }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) SetParent(parent types.Entity) { trapOiDinfo.parent = parent }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetParent() types.Entity { return trapOiDinfo.parent }

func (trapOiDinfo *Snmp_Information_TrapInfos_TrapInfo_TrapOiDinfo) GetParentYangName() string { return "trap-info" }

// Snmp_Information_PollOids
// OID list for poll PDU
type Snmp_Information_PollOids struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PDU drop info for OID. The type is slice of
    // Snmp_Information_PollOids_PollOid.
    PollOid []Snmp_Information_PollOids_PollOid
}

func (pollOids *Snmp_Information_PollOids) GetFilter() yfilter.YFilter { return pollOids.YFilter }

func (pollOids *Snmp_Information_PollOids) SetFilter(yf yfilter.YFilter) { pollOids.YFilter = yf }

func (pollOids *Snmp_Information_PollOids) GetGoName(yname string) string {
    if yname == "poll-oid" { return "PollOid" }
    return ""
}

func (pollOids *Snmp_Information_PollOids) GetSegmentPath() string {
    return "poll-oids"
}

func (pollOids *Snmp_Information_PollOids) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "poll-oid" {
        for _, c := range pollOids.PollOid {
            if pollOids.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_PollOids_PollOid{}
        pollOids.PollOid = append(pollOids.PollOid, child)
        return &pollOids.PollOid[len(pollOids.PollOid)-1]
    }
    return nil
}

func (pollOids *Snmp_Information_PollOids) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pollOids.PollOid {
        children[pollOids.PollOid[i].GetSegmentPath()] = &pollOids.PollOid[i]
    }
    return children
}

func (pollOids *Snmp_Information_PollOids) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pollOids *Snmp_Information_PollOids) GetBundleName() string { return "cisco_ios_xr" }

func (pollOids *Snmp_Information_PollOids) GetYangName() string { return "poll-oids" }

func (pollOids *Snmp_Information_PollOids) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pollOids *Snmp_Information_PollOids) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pollOids *Snmp_Information_PollOids) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pollOids *Snmp_Information_PollOids) SetParent(parent types.Entity) { pollOids.parent = parent }

func (pollOids *Snmp_Information_PollOids) GetParent() types.Entity { return pollOids.parent }

func (pollOids *Snmp_Information_PollOids) GetParentYangName() string { return "information" }

// Snmp_Information_PollOids_PollOid
// PDU drop info for OID
type Snmp_Information_PollOids_PollOid struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (pollOid *Snmp_Information_PollOids_PollOid) GetFilter() yfilter.YFilter { return pollOid.YFilter }

func (pollOid *Snmp_Information_PollOids_PollOid) SetFilter(yf yfilter.YFilter) { pollOid.YFilter = yf }

func (pollOid *Snmp_Information_PollOids_PollOid) GetGoName(yname string) string {
    if yname == "object-id" { return "ObjectId" }
    if yname == "nms-count" { return "NmsCount" }
    if yname == "nms" { return "Nms" }
    if yname == "request-count" { return "RequestCount" }
    return ""
}

func (pollOid *Snmp_Information_PollOids_PollOid) GetSegmentPath() string {
    return "poll-oid" + "[object-id='" + fmt.Sprintf("%v", pollOid.ObjectId) + "']"
}

func (pollOid *Snmp_Information_PollOids_PollOid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pollOid *Snmp_Information_PollOids_PollOid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pollOid *Snmp_Information_PollOids_PollOid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-id"] = pollOid.ObjectId
    leafs["nms-count"] = pollOid.NmsCount
    leafs["nms"] = pollOid.Nms
    leafs["request-count"] = pollOid.RequestCount
    return leafs
}

func (pollOid *Snmp_Information_PollOids_PollOid) GetBundleName() string { return "cisco_ios_xr" }

func (pollOid *Snmp_Information_PollOids_PollOid) GetYangName() string { return "poll-oid" }

func (pollOid *Snmp_Information_PollOids_PollOid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pollOid *Snmp_Information_PollOids_PollOid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pollOid *Snmp_Information_PollOids_PollOid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pollOid *Snmp_Information_PollOids_PollOid) SetParent(parent types.Entity) { pollOid.parent = parent }

func (pollOid *Snmp_Information_PollOids_PollOid) GetParent() types.Entity { return pollOid.parent }

func (pollOid *Snmp_Information_PollOids_PollOid) GetParentYangName() string { return "poll-oids" }

// Snmp_Information_InfomDetails
// SNMP trap OID
type Snmp_Information_InfomDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP Trap infomation like server , port and trapOID. The type is slice of
    // Snmp_Information_InfomDetails_InfomDetail.
    InfomDetail []Snmp_Information_InfomDetails_InfomDetail
}

func (infomDetails *Snmp_Information_InfomDetails) GetFilter() yfilter.YFilter { return infomDetails.YFilter }

func (infomDetails *Snmp_Information_InfomDetails) SetFilter(yf yfilter.YFilter) { infomDetails.YFilter = yf }

func (infomDetails *Snmp_Information_InfomDetails) GetGoName(yname string) string {
    if yname == "infom-detail" { return "InfomDetail" }
    return ""
}

func (infomDetails *Snmp_Information_InfomDetails) GetSegmentPath() string {
    return "infom-details"
}

func (infomDetails *Snmp_Information_InfomDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "infom-detail" {
        for _, c := range infomDetails.InfomDetail {
            if infomDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_InfomDetails_InfomDetail{}
        infomDetails.InfomDetail = append(infomDetails.InfomDetail, child)
        return &infomDetails.InfomDetail[len(infomDetails.InfomDetail)-1]
    }
    return nil
}

func (infomDetails *Snmp_Information_InfomDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range infomDetails.InfomDetail {
        children[infomDetails.InfomDetail[i].GetSegmentPath()] = &infomDetails.InfomDetail[i]
    }
    return children
}

func (infomDetails *Snmp_Information_InfomDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (infomDetails *Snmp_Information_InfomDetails) GetBundleName() string { return "cisco_ios_xr" }

func (infomDetails *Snmp_Information_InfomDetails) GetYangName() string { return "infom-details" }

func (infomDetails *Snmp_Information_InfomDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infomDetails *Snmp_Information_InfomDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infomDetails *Snmp_Information_InfomDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infomDetails *Snmp_Information_InfomDetails) SetParent(parent types.Entity) { infomDetails.parent = parent }

func (infomDetails *Snmp_Information_InfomDetails) GetParent() types.Entity { return infomDetails.parent }

func (infomDetails *Snmp_Information_InfomDetails) GetParentYangName() string { return "information" }

// Snmp_Information_InfomDetails_InfomDetail
// SNMP Trap infomation like server , port and
// trapOID
type Snmp_Information_InfomDetails_InfomDetail struct {
    parent types.Entity
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

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetFilter() yfilter.YFilter { return infomDetail.YFilter }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) SetFilter(yf yfilter.YFilter) { infomDetail.YFilter = yf }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetGoName(yname string) string {
    if yname == "trap-host" { return "TrapHost" }
    if yname == "port" { return "Port" }
    if yname == "host" { return "Host" }
    if yname == "port-xr" { return "PortXr" }
    if yname == "trap-oid-count" { return "TrapOidCount" }
    if yname == "trap-oi-dinfo" { return "TrapOiDinfo" }
    return ""
}

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetSegmentPath() string {
    return "infom-detail"
}

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-oi-dinfo" {
        for _, c := range infomDetail.TrapOiDinfo {
            if infomDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo{}
        infomDetail.TrapOiDinfo = append(infomDetail.TrapOiDinfo, child)
        return &infomDetail.TrapOiDinfo[len(infomDetail.TrapOiDinfo)-1]
    }
    return nil
}

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range infomDetail.TrapOiDinfo {
        children[infomDetail.TrapOiDinfo[i].GetSegmentPath()] = &infomDetail.TrapOiDinfo[i]
    }
    return children
}

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-host"] = infomDetail.TrapHost
    leafs["port"] = infomDetail.Port
    leafs["host"] = infomDetail.Host
    leafs["port-xr"] = infomDetail.PortXr
    leafs["trap-oid-count"] = infomDetail.TrapOidCount
    return leafs
}

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetBundleName() string { return "cisco_ios_xr" }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetYangName() string { return "infom-detail" }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) SetParent(parent types.Entity) { infomDetail.parent = parent }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetParent() types.Entity { return infomDetail.parent }

func (infomDetail *Snmp_Information_InfomDetails_InfomDetail) GetParentYangName() string { return "infom-details" }

// Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo
// Per trap OID statistics
type Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo struct {
    parent types.Entity
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

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetFilter() yfilter.YFilter { return trapOiDinfo.YFilter }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) SetFilter(yf yfilter.YFilter) { trapOiDinfo.YFilter = yf }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetGoName(yname string) string {
    if yname == "trap-oid" { return "TrapOid" }
    if yname == "count" { return "Count" }
    if yname == "drop-count" { return "DropCount" }
    if yname == "retry-count" { return "RetryCount" }
    if yname == "lastsent-time" { return "LastsentTime" }
    if yname == "lasrdrop-time" { return "LasrdropTime" }
    return ""
}

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetSegmentPath() string {
    return "trap-oi-dinfo"
}

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-oid"] = trapOiDinfo.TrapOid
    leafs["count"] = trapOiDinfo.Count
    leafs["drop-count"] = trapOiDinfo.DropCount
    leafs["retry-count"] = trapOiDinfo.RetryCount
    leafs["lastsent-time"] = trapOiDinfo.LastsentTime
    leafs["lasrdrop-time"] = trapOiDinfo.LasrdropTime
    return leafs
}

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetBundleName() string { return "cisco_ios_xr" }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetYangName() string { return "trap-oi-dinfo" }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) SetParent(parent types.Entity) { trapOiDinfo.parent = parent }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetParent() types.Entity { return trapOiDinfo.parent }

func (trapOiDinfo *Snmp_Information_InfomDetails_InfomDetail_TrapOiDinfo) GetParentYangName() string { return "infom-detail" }

// Snmp_Information_Statistics
// SNMP statistics
type Snmp_Information_Statistics struct {
    parent types.Entity
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

func (statistics *Snmp_Information_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Snmp_Information_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Snmp_Information_Statistics) GetGoName(yname string) string {
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "bad-versions-received" { return "BadVersionsReceived" }
    if yname == "bad-community-names-received" { return "BadCommunityNamesReceived" }
    if yname == "bad-community-uses-received" { return "BadCommunityUsesReceived" }
    if yname == "asn-parse-errors-received" { return "AsnParseErrorsReceived" }
    if yname == "silent-drop-count" { return "SilentDropCount" }
    if yname == "proxy-drop-count" { return "ProxyDropCount" }
    if yname == "too-big-packet-received" { return "TooBigPacketReceived" }
    if yname == "max-packet-size" { return "MaxPacketSize" }
    if yname == "no-such-names-received" { return "NoSuchNamesReceived" }
    if yname == "bad-values-received" { return "BadValuesReceived" }
    if yname == "read-only-received" { return "ReadOnlyReceived" }
    if yname == "total-general-errors" { return "TotalGeneralErrors" }
    if yname == "total-requested-variables" { return "TotalRequestedVariables" }
    if yname == "total-set-variables-received" { return "TotalSetVariablesReceived" }
    if yname == "get-requests-received" { return "GetRequestsReceived" }
    if yname == "get-next-requests-received" { return "GetNextRequestsReceived" }
    if yname == "set-requests-received" { return "SetRequestsReceived" }
    if yname == "get-responses-received" { return "GetResponsesReceived" }
    if yname == "traps-received" { return "TrapsReceived" }
    if yname == "total-packets-sent" { return "TotalPacketsSent" }
    if yname == "too-big-packets-sent" { return "TooBigPacketsSent" }
    if yname == "no-such-names-sent" { return "NoSuchNamesSent" }
    if yname == "bad-values-sent" { return "BadValuesSent" }
    if yname == "general-errors-sent" { return "GeneralErrorsSent" }
    if yname == "get-requests-sent" { return "GetRequestsSent" }
    if yname == "get-next-request-sent" { return "GetNextRequestSent" }
    if yname == "set-requests-sent" { return "SetRequestsSent" }
    if yname == "get-responses-sent" { return "GetResponsesSent" }
    if yname == "traps-sent" { return "TrapsSent" }
    return ""
}

func (statistics *Snmp_Information_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Snmp_Information_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *Snmp_Information_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *Snmp_Information_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packets-received"] = statistics.PacketsReceived
    leafs["bad-versions-received"] = statistics.BadVersionsReceived
    leafs["bad-community-names-received"] = statistics.BadCommunityNamesReceived
    leafs["bad-community-uses-received"] = statistics.BadCommunityUsesReceived
    leafs["asn-parse-errors-received"] = statistics.AsnParseErrorsReceived
    leafs["silent-drop-count"] = statistics.SilentDropCount
    leafs["proxy-drop-count"] = statistics.ProxyDropCount
    leafs["too-big-packet-received"] = statistics.TooBigPacketReceived
    leafs["max-packet-size"] = statistics.MaxPacketSize
    leafs["no-such-names-received"] = statistics.NoSuchNamesReceived
    leafs["bad-values-received"] = statistics.BadValuesReceived
    leafs["read-only-received"] = statistics.ReadOnlyReceived
    leafs["total-general-errors"] = statistics.TotalGeneralErrors
    leafs["total-requested-variables"] = statistics.TotalRequestedVariables
    leafs["total-set-variables-received"] = statistics.TotalSetVariablesReceived
    leafs["get-requests-received"] = statistics.GetRequestsReceived
    leafs["get-next-requests-received"] = statistics.GetNextRequestsReceived
    leafs["set-requests-received"] = statistics.SetRequestsReceived
    leafs["get-responses-received"] = statistics.GetResponsesReceived
    leafs["traps-received"] = statistics.TrapsReceived
    leafs["total-packets-sent"] = statistics.TotalPacketsSent
    leafs["too-big-packets-sent"] = statistics.TooBigPacketsSent
    leafs["no-such-names-sent"] = statistics.NoSuchNamesSent
    leafs["bad-values-sent"] = statistics.BadValuesSent
    leafs["general-errors-sent"] = statistics.GeneralErrorsSent
    leafs["get-requests-sent"] = statistics.GetRequestsSent
    leafs["get-next-request-sent"] = statistics.GetNextRequestSent
    leafs["set-requests-sent"] = statistics.SetRequestsSent
    leafs["get-responses-sent"] = statistics.GetResponsesSent
    leafs["traps-sent"] = statistics.TrapsSent
    return leafs
}

func (statistics *Snmp_Information_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Snmp_Information_Statistics) GetYangName() string { return "statistics" }

func (statistics *Snmp_Information_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Snmp_Information_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Snmp_Information_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Snmp_Information_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Snmp_Information_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Snmp_Information_Statistics) GetParentYangName() string { return "information" }

// Snmp_Information_IncomingQueue
// Incoming queue details 
type Snmp_Information_IncomingQueue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of NMS Queues Exist. The type is interface{} with range:
    // 0..4294967295.
    QueueCount interface{}

    // Each Entry Details. The type is slice of
    // Snmp_Information_IncomingQueue_InqEntry.
    InqEntry []Snmp_Information_IncomingQueue_InqEntry
}

func (incomingQueue *Snmp_Information_IncomingQueue) GetFilter() yfilter.YFilter { return incomingQueue.YFilter }

func (incomingQueue *Snmp_Information_IncomingQueue) SetFilter(yf yfilter.YFilter) { incomingQueue.YFilter = yf }

func (incomingQueue *Snmp_Information_IncomingQueue) GetGoName(yname string) string {
    if yname == "queue-count" { return "QueueCount" }
    if yname == "inq-entry" { return "InqEntry" }
    return ""
}

func (incomingQueue *Snmp_Information_IncomingQueue) GetSegmentPath() string {
    return "incoming-queue"
}

func (incomingQueue *Snmp_Information_IncomingQueue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inq-entry" {
        for _, c := range incomingQueue.InqEntry {
            if incomingQueue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_IncomingQueue_InqEntry{}
        incomingQueue.InqEntry = append(incomingQueue.InqEntry, child)
        return &incomingQueue.InqEntry[len(incomingQueue.InqEntry)-1]
    }
    return nil
}

func (incomingQueue *Snmp_Information_IncomingQueue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range incomingQueue.InqEntry {
        children[incomingQueue.InqEntry[i].GetSegmentPath()] = &incomingQueue.InqEntry[i]
    }
    return children
}

func (incomingQueue *Snmp_Information_IncomingQueue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-count"] = incomingQueue.QueueCount
    return leafs
}

func (incomingQueue *Snmp_Information_IncomingQueue) GetBundleName() string { return "cisco_ios_xr" }

func (incomingQueue *Snmp_Information_IncomingQueue) GetYangName() string { return "incoming-queue" }

func (incomingQueue *Snmp_Information_IncomingQueue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incomingQueue *Snmp_Information_IncomingQueue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incomingQueue *Snmp_Information_IncomingQueue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incomingQueue *Snmp_Information_IncomingQueue) SetParent(parent types.Entity) { incomingQueue.parent = parent }

func (incomingQueue *Snmp_Information_IncomingQueue) GetParent() types.Entity { return incomingQueue.parent }

func (incomingQueue *Snmp_Information_IncomingQueue) GetParentYangName() string { return "information" }

// Snmp_Information_IncomingQueue_InqEntry
// Each Entry Details.
type Snmp_Information_IncomingQueue_InqEntry struct {
    parent types.Entity
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

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetFilter() yfilter.YFilter { return inqEntry.YFilter }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) SetFilter(yf yfilter.YFilter) { inqEntry.YFilter = yf }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetGoName(yname string) string {
    if yname == "address-of-queue" { return "AddressOfQueue" }
    if yname == "request-count" { return "RequestCount" }
    if yname == "processed-request-count" { return "ProcessedRequestCount" }
    if yname == "last-access-time" { return "LastAccessTime" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetSegmentPath() string {
    return "inq-entry"
}

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-of-queue"] = inqEntry.AddressOfQueue
    leafs["request-count"] = inqEntry.RequestCount
    leafs["processed-request-count"] = inqEntry.ProcessedRequestCount
    leafs["last-access-time"] = inqEntry.LastAccessTime
    leafs["priority"] = inqEntry.Priority
    return leafs
}

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetBundleName() string { return "cisco_ios_xr" }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetYangName() string { return "inq-entry" }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) SetParent(parent types.Entity) { inqEntry.parent = parent }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetParent() types.Entity { return inqEntry.parent }

func (inqEntry *Snmp_Information_IncomingQueue_InqEntry) GetParentYangName() string { return "incoming-queue" }

// Snmp_Information_ContextMapping
// Context name, features name, topology name,
// instance
type Snmp_Information_ContextMapping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Context Mapping. The type is slice of
    // Snmp_Information_ContextMapping_ContexMapping.
    ContexMapping []Snmp_Information_ContextMapping_ContexMapping
}

func (contextMapping *Snmp_Information_ContextMapping) GetFilter() yfilter.YFilter { return contextMapping.YFilter }

func (contextMapping *Snmp_Information_ContextMapping) SetFilter(yf yfilter.YFilter) { contextMapping.YFilter = yf }

func (contextMapping *Snmp_Information_ContextMapping) GetGoName(yname string) string {
    if yname == "contex-mapping" { return "ContexMapping" }
    return ""
}

func (contextMapping *Snmp_Information_ContextMapping) GetSegmentPath() string {
    return "context-mapping"
}

func (contextMapping *Snmp_Information_ContextMapping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "contex-mapping" {
        for _, c := range contextMapping.ContexMapping {
            if contextMapping.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_ContextMapping_ContexMapping{}
        contextMapping.ContexMapping = append(contextMapping.ContexMapping, child)
        return &contextMapping.ContexMapping[len(contextMapping.ContexMapping)-1]
    }
    return nil
}

func (contextMapping *Snmp_Information_ContextMapping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextMapping.ContexMapping {
        children[contextMapping.ContexMapping[i].GetSegmentPath()] = &contextMapping.ContexMapping[i]
    }
    return children
}

func (contextMapping *Snmp_Information_ContextMapping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contextMapping *Snmp_Information_ContextMapping) GetBundleName() string { return "cisco_ios_xr" }

func (contextMapping *Snmp_Information_ContextMapping) GetYangName() string { return "context-mapping" }

func (contextMapping *Snmp_Information_ContextMapping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextMapping *Snmp_Information_ContextMapping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextMapping *Snmp_Information_ContextMapping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextMapping *Snmp_Information_ContextMapping) SetParent(parent types.Entity) { contextMapping.parent = parent }

func (contextMapping *Snmp_Information_ContextMapping) GetParent() types.Entity { return contextMapping.parent }

func (contextMapping *Snmp_Information_ContextMapping) GetParentYangName() string { return "information" }

// Snmp_Information_ContextMapping_ContexMapping
// Context Mapping
type Snmp_Information_ContextMapping_ContexMapping struct {
    parent types.Entity
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

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetFilter() yfilter.YFilter { return contexMapping.YFilter }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) SetFilter(yf yfilter.YFilter) { contexMapping.YFilter = yf }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetGoName(yname string) string {
    if yname == "context" { return "Context" }
    if yname == "feature-name" { return "FeatureName" }
    if yname == "instance" { return "Instance" }
    if yname == "topology" { return "Topology" }
    if yname == "feature" { return "Feature" }
    return ""
}

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetSegmentPath() string {
    return "contex-mapping"
}

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context"] = contexMapping.Context
    leafs["feature-name"] = contexMapping.FeatureName
    leafs["instance"] = contexMapping.Instance
    leafs["topology"] = contexMapping.Topology
    leafs["feature"] = contexMapping.Feature
    return leafs
}

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetBundleName() string { return "cisco_ios_xr" }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetYangName() string { return "contex-mapping" }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) SetParent(parent types.Entity) { contexMapping.parent = parent }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetParent() types.Entity { return contexMapping.parent }

func (contexMapping *Snmp_Information_ContextMapping_ContexMapping) GetParentYangName() string { return "context-mapping" }

// Snmp_Information_TrapOids
// SNMP trap OID
type Snmp_Information_TrapOids struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP trap . The type is slice of Snmp_Information_TrapOids_TrapOid.
    TrapOid []Snmp_Information_TrapOids_TrapOid
}

func (trapOids *Snmp_Information_TrapOids) GetFilter() yfilter.YFilter { return trapOids.YFilter }

func (trapOids *Snmp_Information_TrapOids) SetFilter(yf yfilter.YFilter) { trapOids.YFilter = yf }

func (trapOids *Snmp_Information_TrapOids) GetGoName(yname string) string {
    if yname == "trap-oid" { return "TrapOid" }
    return ""
}

func (trapOids *Snmp_Information_TrapOids) GetSegmentPath() string {
    return "trap-oids"
}

func (trapOids *Snmp_Information_TrapOids) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-oid" {
        for _, c := range trapOids.TrapOid {
            if trapOids.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_TrapOids_TrapOid{}
        trapOids.TrapOid = append(trapOids.TrapOid, child)
        return &trapOids.TrapOid[len(trapOids.TrapOid)-1]
    }
    return nil
}

func (trapOids *Snmp_Information_TrapOids) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapOids.TrapOid {
        children[trapOids.TrapOid[i].GetSegmentPath()] = &trapOids.TrapOid[i]
    }
    return children
}

func (trapOids *Snmp_Information_TrapOids) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trapOids *Snmp_Information_TrapOids) GetBundleName() string { return "cisco_ios_xr" }

func (trapOids *Snmp_Information_TrapOids) GetYangName() string { return "trap-oids" }

func (trapOids *Snmp_Information_TrapOids) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapOids *Snmp_Information_TrapOids) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapOids *Snmp_Information_TrapOids) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapOids *Snmp_Information_TrapOids) SetParent(parent types.Entity) { trapOids.parent = parent }

func (trapOids *Snmp_Information_TrapOids) GetParent() types.Entity { return trapOids.parent }

func (trapOids *Snmp_Information_TrapOids) GetParentYangName() string { return "information" }

// Snmp_Information_TrapOids_TrapOid
// SNMP trap 
type Snmp_Information_TrapOids_TrapOid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Trap object ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TrapOid interface{}

    // Total number of OID's sent. The type is interface{} with range:
    // 0..4294967295.
    TrapOidCount interface{}

    // TRAP OID. The type is string.
    TrapOidXr interface{}
}

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetFilter() yfilter.YFilter { return trapOid.YFilter }

func (trapOid *Snmp_Information_TrapOids_TrapOid) SetFilter(yf yfilter.YFilter) { trapOid.YFilter = yf }

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetGoName(yname string) string {
    if yname == "trap-oid" { return "TrapOid" }
    if yname == "trap-oid-count" { return "TrapOidCount" }
    if yname == "trap-oid-xr" { return "TrapOidXr" }
    return ""
}

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetSegmentPath() string {
    return "trap-oid" + "[trap-oid='" + fmt.Sprintf("%v", trapOid.TrapOid) + "']"
}

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-oid"] = trapOid.TrapOid
    leafs["trap-oid-count"] = trapOid.TrapOidCount
    leafs["trap-oid-xr"] = trapOid.TrapOidXr
    return leafs
}

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetBundleName() string { return "cisco_ios_xr" }

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetYangName() string { return "trap-oid" }

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapOid *Snmp_Information_TrapOids_TrapOid) SetParent(parent types.Entity) { trapOid.parent = parent }

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetParent() types.Entity { return trapOid.parent }

func (trapOid *Snmp_Information_TrapOids_TrapOid) GetParentYangName() string { return "trap-oids" }

// Snmp_Information_NmSpackets
// SNMP overload statistics 
type Snmp_Information_NmSpackets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NMS packet drop count. The type is slice of
    // Snmp_Information_NmSpackets_NmSpacket.
    NmSpacket []Snmp_Information_NmSpackets_NmSpacket
}

func (nmSpackets *Snmp_Information_NmSpackets) GetFilter() yfilter.YFilter { return nmSpackets.YFilter }

func (nmSpackets *Snmp_Information_NmSpackets) SetFilter(yf yfilter.YFilter) { nmSpackets.YFilter = yf }

func (nmSpackets *Snmp_Information_NmSpackets) GetGoName(yname string) string {
    if yname == "nm-spacket" { return "NmSpacket" }
    return ""
}

func (nmSpackets *Snmp_Information_NmSpackets) GetSegmentPath() string {
    return "nm-spackets"
}

func (nmSpackets *Snmp_Information_NmSpackets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nm-spacket" {
        for _, c := range nmSpackets.NmSpacket {
            if nmSpackets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_NmSpackets_NmSpacket{}
        nmSpackets.NmSpacket = append(nmSpackets.NmSpacket, child)
        return &nmSpackets.NmSpacket[len(nmSpackets.NmSpacket)-1]
    }
    return nil
}

func (nmSpackets *Snmp_Information_NmSpackets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nmSpackets.NmSpacket {
        children[nmSpackets.NmSpacket[i].GetSegmentPath()] = &nmSpackets.NmSpacket[i]
    }
    return children
}

func (nmSpackets *Snmp_Information_NmSpackets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nmSpackets *Snmp_Information_NmSpackets) GetBundleName() string { return "cisco_ios_xr" }

func (nmSpackets *Snmp_Information_NmSpackets) GetYangName() string { return "nm-spackets" }

func (nmSpackets *Snmp_Information_NmSpackets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nmSpackets *Snmp_Information_NmSpackets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nmSpackets *Snmp_Information_NmSpackets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nmSpackets *Snmp_Information_NmSpackets) SetParent(parent types.Entity) { nmSpackets.parent = parent }

func (nmSpackets *Snmp_Information_NmSpackets) GetParent() types.Entity { return nmSpackets.parent }

func (nmSpackets *Snmp_Information_NmSpackets) GetParentYangName() string { return "information" }

// Snmp_Information_NmSpackets_NmSpacket
// NMS packet drop count
type Snmp_Information_NmSpackets_NmSpacket struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetFilter() yfilter.YFilter { return nmSpacket.YFilter }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) SetFilter(yf yfilter.YFilter) { nmSpacket.YFilter = yf }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetGoName(yname string) string {
    if yname == "packetcount" { return "Packetcount" }
    if yname == "number-of-nmsq-pkts-dropped" { return "NumberOfNmsqPktsDropped" }
    if yname == "number-of-pkts-dropped" { return "NumberOfPktsDropped" }
    if yname == "overload-start-time" { return "OverloadStartTime" }
    if yname == "overload-end-time" { return "OverloadEndTime" }
    return ""
}

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetSegmentPath() string {
    return "nm-spacket" + "[packetcount='" + fmt.Sprintf("%v", nmSpacket.Packetcount) + "']"
}

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packetcount"] = nmSpacket.Packetcount
    leafs["number-of-nmsq-pkts-dropped"] = nmSpacket.NumberOfNmsqPktsDropped
    leafs["number-of-pkts-dropped"] = nmSpacket.NumberOfPktsDropped
    leafs["overload-start-time"] = nmSpacket.OverloadStartTime
    leafs["overload-end-time"] = nmSpacket.OverloadEndTime
    return leafs
}

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetBundleName() string { return "cisco_ios_xr" }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetYangName() string { return "nm-spacket" }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) SetParent(parent types.Entity) { nmSpacket.parent = parent }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetParent() types.Entity { return nmSpacket.parent }

func (nmSpacket *Snmp_Information_NmSpackets_NmSpacket) GetParentYangName() string { return "nm-spackets" }

// Snmp_Information_Mibs
// List of MIBS supported on the system
type Snmp_Information_Mibs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP MIB Name. The type is slice of Snmp_Information_Mibs_Mib.
    Mib []Snmp_Information_Mibs_Mib
}

func (mibs *Snmp_Information_Mibs) GetFilter() yfilter.YFilter { return mibs.YFilter }

func (mibs *Snmp_Information_Mibs) SetFilter(yf yfilter.YFilter) { mibs.YFilter = yf }

func (mibs *Snmp_Information_Mibs) GetGoName(yname string) string {
    if yname == "mib" { return "Mib" }
    return ""
}

func (mibs *Snmp_Information_Mibs) GetSegmentPath() string {
    return "mibs"
}

func (mibs *Snmp_Information_Mibs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mib" {
        for _, c := range mibs.Mib {
            if mibs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Mibs_Mib{}
        mibs.Mib = append(mibs.Mib, child)
        return &mibs.Mib[len(mibs.Mib)-1]
    }
    return nil
}

func (mibs *Snmp_Information_Mibs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mibs.Mib {
        children[mibs.Mib[i].GetSegmentPath()] = &mibs.Mib[i]
    }
    return children
}

func (mibs *Snmp_Information_Mibs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mibs *Snmp_Information_Mibs) GetBundleName() string { return "cisco_ios_xr" }

func (mibs *Snmp_Information_Mibs) GetYangName() string { return "mibs" }

func (mibs *Snmp_Information_Mibs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mibs *Snmp_Information_Mibs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mibs *Snmp_Information_Mibs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mibs *Snmp_Information_Mibs) SetParent(parent types.Entity) { mibs.parent = parent }

func (mibs *Snmp_Information_Mibs) GetParent() types.Entity { return mibs.parent }

func (mibs *Snmp_Information_Mibs) GetParentYangName() string { return "information" }

// Snmp_Information_Mibs_Mib
// SNMP MIB Name
type Snmp_Information_Mibs_Mib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. MIB Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // List of OIDs per MIB.
    Oids Snmp_Information_Mibs_Mib_Oids

    // MIB state and information.
    MibInformation Snmp_Information_Mibs_Mib_MibInformation
}

func (mib *Snmp_Information_Mibs_Mib) GetFilter() yfilter.YFilter { return mib.YFilter }

func (mib *Snmp_Information_Mibs_Mib) SetFilter(yf yfilter.YFilter) { mib.YFilter = yf }

func (mib *Snmp_Information_Mibs_Mib) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "oids" { return "Oids" }
    if yname == "mib-information" { return "MibInformation" }
    return ""
}

func (mib *Snmp_Information_Mibs_Mib) GetSegmentPath() string {
    return "mib" + "[name='" + fmt.Sprintf("%v", mib.Name) + "']"
}

func (mib *Snmp_Information_Mibs_Mib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oids" {
        return &mib.Oids
    }
    if childYangName == "mib-information" {
        return &mib.MibInformation
    }
    return nil
}

func (mib *Snmp_Information_Mibs_Mib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["oids"] = &mib.Oids
    children["mib-information"] = &mib.MibInformation
    return children
}

func (mib *Snmp_Information_Mibs_Mib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = mib.Name
    return leafs
}

func (mib *Snmp_Information_Mibs_Mib) GetBundleName() string { return "cisco_ios_xr" }

func (mib *Snmp_Information_Mibs_Mib) GetYangName() string { return "mib" }

func (mib *Snmp_Information_Mibs_Mib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mib *Snmp_Information_Mibs_Mib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mib *Snmp_Information_Mibs_Mib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mib *Snmp_Information_Mibs_Mib) SetParent(parent types.Entity) { mib.parent = parent }

func (mib *Snmp_Information_Mibs_Mib) GetParent() types.Entity { return mib.parent }

func (mib *Snmp_Information_Mibs_Mib) GetParentYangName() string { return "mibs" }

// Snmp_Information_Mibs_Mib_Oids
// List of OIDs per MIB
type Snmp_Information_Mibs_Mib_Oids struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object identifiers of a mib. The type is slice of
    // Snmp_Information_Mibs_Mib_Oids_Oid.
    Oid []Snmp_Information_Mibs_Mib_Oids_Oid
}

func (oids *Snmp_Information_Mibs_Mib_Oids) GetFilter() yfilter.YFilter { return oids.YFilter }

func (oids *Snmp_Information_Mibs_Mib_Oids) SetFilter(yf yfilter.YFilter) { oids.YFilter = yf }

func (oids *Snmp_Information_Mibs_Mib_Oids) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    return ""
}

func (oids *Snmp_Information_Mibs_Mib_Oids) GetSegmentPath() string {
    return "oids"
}

func (oids *Snmp_Information_Mibs_Mib_Oids) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oid" {
        for _, c := range oids.Oid {
            if oids.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Mibs_Mib_Oids_Oid{}
        oids.Oid = append(oids.Oid, child)
        return &oids.Oid[len(oids.Oid)-1]
    }
    return nil
}

func (oids *Snmp_Information_Mibs_Mib_Oids) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range oids.Oid {
        children[oids.Oid[i].GetSegmentPath()] = &oids.Oid[i]
    }
    return children
}

func (oids *Snmp_Information_Mibs_Mib_Oids) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oids *Snmp_Information_Mibs_Mib_Oids) GetBundleName() string { return "cisco_ios_xr" }

func (oids *Snmp_Information_Mibs_Mib_Oids) GetYangName() string { return "oids" }

func (oids *Snmp_Information_Mibs_Mib_Oids) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oids *Snmp_Information_Mibs_Mib_Oids) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oids *Snmp_Information_Mibs_Mib_Oids) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oids *Snmp_Information_Mibs_Mib_Oids) SetParent(parent types.Entity) { oids.parent = parent }

func (oids *Snmp_Information_Mibs_Mib_Oids) GetParent() types.Entity { return oids.parent }

func (oids *Snmp_Information_Mibs_Mib_Oids) GetParentYangName() string { return "mib" }

// Snmp_Information_Mibs_Mib_Oids_Oid
// Object identifiers of a mib
type Snmp_Information_Mibs_Mib_Oids_Oid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object Identifier. The type is string.
    Oid interface{}

    // MIB OID Name. The type is string. This attribute is mandatory.
    OidName interface{}
}

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetFilter() yfilter.YFilter { return oid.YFilter }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) SetFilter(yf yfilter.YFilter) { oid.YFilter = yf }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "oid-name" { return "OidName" }
    return ""
}

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetSegmentPath() string {
    return "oid" + "[oid='" + fmt.Sprintf("%v", oid.Oid) + "']"
}

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = oid.Oid
    leafs["oid-name"] = oid.OidName
    return leafs
}

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetBundleName() string { return "cisco_ios_xr" }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetYangName() string { return "oid" }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) SetParent(parent types.Entity) { oid.parent = parent }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetParent() types.Entity { return oid.parent }

func (oid *Snmp_Information_Mibs_Mib_Oids_Oid) GetParentYangName() string { return "oids" }

// Snmp_Information_Mibs_Mib_MibInformation
// MIB state and information
type Snmp_Information_Mibs_Mib_MibInformation struct {
    parent types.Entity
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

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetFilter() yfilter.YFilter { return mibInformation.YFilter }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) SetFilter(yf yfilter.YFilter) { mibInformation.YFilter = yf }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetGoName(yname string) string {
    if yname == "mib-name" { return "MibName" }
    if yname == "dll-name" { return "DllName" }
    if yname == "mib-config-filename" { return "MibConfigFilename" }
    if yname == "is-mib-loaded" { return "IsMibLoaded" }
    if yname == "dll-capabilities" { return "DllCapabilities" }
    if yname == "trap-strings" { return "TrapStrings" }
    if yname == "timeout" { return "Timeout" }
    if yname == "load-time" { return "LoadTime" }
    return ""
}

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetSegmentPath() string {
    return "mib-information"
}

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mib-name"] = mibInformation.MibName
    leafs["dll-name"] = mibInformation.DllName
    leafs["mib-config-filename"] = mibInformation.MibConfigFilename
    leafs["is-mib-loaded"] = mibInformation.IsMibLoaded
    leafs["dll-capabilities"] = mibInformation.DllCapabilities
    leafs["trap-strings"] = mibInformation.TrapStrings
    leafs["timeout"] = mibInformation.Timeout
    leafs["load-time"] = mibInformation.LoadTime
    return leafs
}

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetBundleName() string { return "cisco_ios_xr" }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetYangName() string { return "mib-information" }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) SetParent(parent types.Entity) { mibInformation.parent = parent }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetParent() types.Entity { return mibInformation.parent }

func (mibInformation *Snmp_Information_Mibs_Mib_MibInformation) GetParentYangName() string { return "mib" }

// Snmp_Information_SerialNumbers
// SNMP statistics pdu 
type Snmp_Information_SerialNumbers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Serial number. The type is slice of
    // Snmp_Information_SerialNumbers_SerialNumber.
    SerialNumber []Snmp_Information_SerialNumbers_SerialNumber
}

func (serialNumbers *Snmp_Information_SerialNumbers) GetFilter() yfilter.YFilter { return serialNumbers.YFilter }

func (serialNumbers *Snmp_Information_SerialNumbers) SetFilter(yf yfilter.YFilter) { serialNumbers.YFilter = yf }

func (serialNumbers *Snmp_Information_SerialNumbers) GetGoName(yname string) string {
    if yname == "serial-number" { return "SerialNumber" }
    return ""
}

func (serialNumbers *Snmp_Information_SerialNumbers) GetSegmentPath() string {
    return "serial-numbers"
}

func (serialNumbers *Snmp_Information_SerialNumbers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "serial-number" {
        for _, c := range serialNumbers.SerialNumber {
            if serialNumbers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_SerialNumbers_SerialNumber{}
        serialNumbers.SerialNumber = append(serialNumbers.SerialNumber, child)
        return &serialNumbers.SerialNumber[len(serialNumbers.SerialNumber)-1]
    }
    return nil
}

func (serialNumbers *Snmp_Information_SerialNumbers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serialNumbers.SerialNumber {
        children[serialNumbers.SerialNumber[i].GetSegmentPath()] = &serialNumbers.SerialNumber[i]
    }
    return children
}

func (serialNumbers *Snmp_Information_SerialNumbers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serialNumbers *Snmp_Information_SerialNumbers) GetBundleName() string { return "cisco_ios_xr" }

func (serialNumbers *Snmp_Information_SerialNumbers) GetYangName() string { return "serial-numbers" }

func (serialNumbers *Snmp_Information_SerialNumbers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serialNumbers *Snmp_Information_SerialNumbers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serialNumbers *Snmp_Information_SerialNumbers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serialNumbers *Snmp_Information_SerialNumbers) SetParent(parent types.Entity) { serialNumbers.parent = parent }

func (serialNumbers *Snmp_Information_SerialNumbers) GetParent() types.Entity { return serialNumbers.parent }

func (serialNumbers *Snmp_Information_SerialNumbers) GetParentYangName() string { return "information" }

// Snmp_Information_SerialNumbers_SerialNumber
// Serial number
type Snmp_Information_SerialNumbers_SerialNumber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Serial number. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetFilter() yfilter.YFilter { return serialNumber.YFilter }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) SetFilter(yf yfilter.YFilter) { serialNumber.YFilter = yf }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "req-id" { return "ReqId" }
    if yname == "port" { return "Port" }
    if yname == "nms" { return "Nms" }
    if yname == "request-id" { return "RequestId" }
    if yname == "port-xr" { return "PortXr" }
    if yname == "pdu-type" { return "PduType" }
    if yname == "error-status" { return "ErrorStatus" }
    if yname == "serial-num" { return "SerialNum" }
    if yname == "input-q" { return "InputQ" }
    if yname == "output-q" { return "OutputQ" }
    if yname == "pending-q" { return "PendingQ" }
    if yname == "response-out" { return "ResponseOut" }
    return ""
}

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetSegmentPath() string {
    return "serial-number"
}

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = serialNumber.Number
    leafs["req-id"] = serialNumber.ReqId
    leafs["port"] = serialNumber.Port
    leafs["nms"] = serialNumber.Nms
    leafs["request-id"] = serialNumber.RequestId
    leafs["port-xr"] = serialNumber.PortXr
    leafs["pdu-type"] = serialNumber.PduType
    leafs["error-status"] = serialNumber.ErrorStatus
    leafs["serial-num"] = serialNumber.SerialNum
    leafs["input-q"] = serialNumber.InputQ
    leafs["output-q"] = serialNumber.OutputQ
    leafs["pending-q"] = serialNumber.PendingQ
    leafs["response-out"] = serialNumber.ResponseOut
    return leafs
}

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetBundleName() string { return "cisco_ios_xr" }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetYangName() string { return "serial-number" }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) SetParent(parent types.Entity) { serialNumber.parent = parent }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetParent() types.Entity { return serialNumber.parent }

func (serialNumber *Snmp_Information_SerialNumbers_SerialNumber) GetParentYangName() string { return "serial-numbers" }

// Snmp_Information_DropNmsAddresses
// NMS list for drop PDU
type Snmp_Information_DropNmsAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PDU drop info for NMS. The type is slice of
    // Snmp_Information_DropNmsAddresses_DropNmsAddress.
    DropNmsAddress []Snmp_Information_DropNmsAddresses_DropNmsAddress
}

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetFilter() yfilter.YFilter { return dropNmsAddresses.YFilter }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) SetFilter(yf yfilter.YFilter) { dropNmsAddresses.YFilter = yf }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetGoName(yname string) string {
    if yname == "drop-nms-address" { return "DropNmsAddress" }
    return ""
}

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetSegmentPath() string {
    return "drop-nms-addresses"
}

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "drop-nms-address" {
        for _, c := range dropNmsAddresses.DropNmsAddress {
            if dropNmsAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_DropNmsAddresses_DropNmsAddress{}
        dropNmsAddresses.DropNmsAddress = append(dropNmsAddresses.DropNmsAddress, child)
        return &dropNmsAddresses.DropNmsAddress[len(dropNmsAddresses.DropNmsAddress)-1]
    }
    return nil
}

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dropNmsAddresses.DropNmsAddress {
        children[dropNmsAddresses.DropNmsAddress[i].GetSegmentPath()] = &dropNmsAddresses.DropNmsAddress[i]
    }
    return children
}

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetYangName() string { return "drop-nms-addresses" }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) SetParent(parent types.Entity) { dropNmsAddresses.parent = parent }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetParent() types.Entity { return dropNmsAddresses.parent }

func (dropNmsAddresses *Snmp_Information_DropNmsAddresses) GetParentYangName() string { return "information" }

// Snmp_Information_DropNmsAddresses_DropNmsAddress
// PDU drop info for NMS
type Snmp_Information_DropNmsAddresses_DropNmsAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetFilter() yfilter.YFilter { return dropNmsAddress.YFilter }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) SetFilter(yf yfilter.YFilter) { dropNmsAddress.YFilter = yf }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetGoName(yname string) string {
    if yname == "nms-addr" { return "NmsAddr" }
    if yname == "nms-address" { return "NmsAddress" }
    if yname == "incoming-q-count" { return "IncomingQCount" }
    if yname == "threshold-incoming-q-count" { return "ThresholdIncomingQCount" }
    if yname == "encode-count" { return "EncodeCount" }
    if yname == "duplicate-count" { return "DuplicateCount" }
    if yname == "stack-count" { return "StackCount" }
    if yname == "aipc-count" { return "AipcCount" }
    if yname == "overload-count" { return "OverloadCount" }
    if yname == "timeout-count" { return "TimeoutCount" }
    if yname == "internal-count" { return "InternalCount" }
    return ""
}

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetSegmentPath() string {
    return "drop-nms-address" + "[nms-addr='" + fmt.Sprintf("%v", dropNmsAddress.NmsAddr) + "']"
}

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nms-addr"] = dropNmsAddress.NmsAddr
    leafs["nms-address"] = dropNmsAddress.NmsAddress
    leafs["incoming-q-count"] = dropNmsAddress.IncomingQCount
    leafs["threshold-incoming-q-count"] = dropNmsAddress.ThresholdIncomingQCount
    leafs["encode-count"] = dropNmsAddress.EncodeCount
    leafs["duplicate-count"] = dropNmsAddress.DuplicateCount
    leafs["stack-count"] = dropNmsAddress.StackCount
    leafs["aipc-count"] = dropNmsAddress.AipcCount
    leafs["overload-count"] = dropNmsAddress.OverloadCount
    leafs["timeout-count"] = dropNmsAddress.TimeoutCount
    leafs["internal-count"] = dropNmsAddress.InternalCount
    return leafs
}

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetBundleName() string { return "cisco_ios_xr" }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetYangName() string { return "drop-nms-address" }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) SetParent(parent types.Entity) { dropNmsAddress.parent = parent }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetParent() types.Entity { return dropNmsAddress.parent }

func (dropNmsAddress *Snmp_Information_DropNmsAddresses_DropNmsAddress) GetParentYangName() string { return "drop-nms-addresses" }

// Snmp_Information_Views
// SNMP view information
type Snmp_Information_Views struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP target view name. The type is slice of Snmp_Information_Views_View.
    View []Snmp_Information_Views_View
}

func (views *Snmp_Information_Views) GetFilter() yfilter.YFilter { return views.YFilter }

func (views *Snmp_Information_Views) SetFilter(yf yfilter.YFilter) { views.YFilter = yf }

func (views *Snmp_Information_Views) GetGoName(yname string) string {
    if yname == "view" { return "View" }
    return ""
}

func (views *Snmp_Information_Views) GetSegmentPath() string {
    return "views"
}

func (views *Snmp_Information_Views) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "view" {
        for _, c := range views.View {
            if views.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Views_View{}
        views.View = append(views.View, child)
        return &views.View[len(views.View)-1]
    }
    return nil
}

func (views *Snmp_Information_Views) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range views.View {
        children[views.View[i].GetSegmentPath()] = &views.View[i]
    }
    return children
}

func (views *Snmp_Information_Views) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (views *Snmp_Information_Views) GetBundleName() string { return "cisco_ios_xr" }

func (views *Snmp_Information_Views) GetYangName() string { return "views" }

func (views *Snmp_Information_Views) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (views *Snmp_Information_Views) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (views *Snmp_Information_Views) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (views *Snmp_Information_Views) SetParent(parent types.Entity) { views.parent = parent }

func (views *Snmp_Information_Views) GetParent() types.Entity { return views.parent }

func (views *Snmp_Information_Views) GetParentYangName() string { return "information" }

// Snmp_Information_Views_View
// SNMP target view name
type Snmp_Information_Views_View struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. View name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // View name ,familytype, storagetype and status. The type is slice of
    // Snmp_Information_Views_View_ViewInformation.
    ViewInformation []Snmp_Information_Views_View_ViewInformation
}

func (view *Snmp_Information_Views_View) GetFilter() yfilter.YFilter { return view.YFilter }

func (view *Snmp_Information_Views_View) SetFilter(yf yfilter.YFilter) { view.YFilter = yf }

func (view *Snmp_Information_Views_View) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "view-information" { return "ViewInformation" }
    return ""
}

func (view *Snmp_Information_Views_View) GetSegmentPath() string {
    return "view" + "[name='" + fmt.Sprintf("%v", view.Name) + "']"
}

func (view *Snmp_Information_Views_View) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "view-information" {
        for _, c := range view.ViewInformation {
            if view.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Views_View_ViewInformation{}
        view.ViewInformation = append(view.ViewInformation, child)
        return &view.ViewInformation[len(view.ViewInformation)-1]
    }
    return nil
}

func (view *Snmp_Information_Views_View) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range view.ViewInformation {
        children[view.ViewInformation[i].GetSegmentPath()] = &view.ViewInformation[i]
    }
    return children
}

func (view *Snmp_Information_Views_View) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = view.Name
    return leafs
}

func (view *Snmp_Information_Views_View) GetBundleName() string { return "cisco_ios_xr" }

func (view *Snmp_Information_Views_View) GetYangName() string { return "view" }

func (view *Snmp_Information_Views_View) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (view *Snmp_Information_Views_View) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (view *Snmp_Information_Views_View) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (view *Snmp_Information_Views_View) SetParent(parent types.Entity) { view.parent = parent }

func (view *Snmp_Information_Views_View) GetParent() types.Entity { return view.parent }

func (view *Snmp_Information_Views_View) GetParentYangName() string { return "views" }

// Snmp_Information_Views_View_ViewInformation
// View name ,familytype, storagetype and status
type Snmp_Information_Views_View_ViewInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetFilter() yfilter.YFilter { return viewInformation.YFilter }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) SetFilter(yf yfilter.YFilter) { viewInformation.YFilter = yf }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetGoName(yname string) string {
    if yname == "object-id" { return "ObjectId" }
    if yname == "snmp-view-family-type" { return "SnmpViewFamilyType" }
    if yname == "snmp-view-family-storage-type" { return "SnmpViewFamilyStorageType" }
    if yname == "snmp-view-family-status" { return "SnmpViewFamilyStatus" }
    return ""
}

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetSegmentPath() string {
    return "view-information" + "[object-id='" + fmt.Sprintf("%v", viewInformation.ObjectId) + "']"
}

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-id"] = viewInformation.ObjectId
    leafs["snmp-view-family-type"] = viewInformation.SnmpViewFamilyType
    leafs["snmp-view-family-storage-type"] = viewInformation.SnmpViewFamilyStorageType
    leafs["snmp-view-family-status"] = viewInformation.SnmpViewFamilyStatus
    return leafs
}

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetBundleName() string { return "cisco_ios_xr" }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetYangName() string { return "view-information" }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) SetParent(parent types.Entity) { viewInformation.parent = parent }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetParent() types.Entity { return viewInformation.parent }

func (viewInformation *Snmp_Information_Views_View_ViewInformation) GetParentYangName() string { return "view" }

// Snmp_Information_SystemDescr
// System description
type Snmp_Information_SystemDescr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // sysDescr  1.3.6.1.2.1.1.1. The type is string.
    SysDescr interface{}
}

func (systemDescr *Snmp_Information_SystemDescr) GetFilter() yfilter.YFilter { return systemDescr.YFilter }

func (systemDescr *Snmp_Information_SystemDescr) SetFilter(yf yfilter.YFilter) { systemDescr.YFilter = yf }

func (systemDescr *Snmp_Information_SystemDescr) GetGoName(yname string) string {
    if yname == "sys-descr" { return "SysDescr" }
    return ""
}

func (systemDescr *Snmp_Information_SystemDescr) GetSegmentPath() string {
    return "system-descr"
}

func (systemDescr *Snmp_Information_SystemDescr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemDescr *Snmp_Information_SystemDescr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemDescr *Snmp_Information_SystemDescr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sys-descr"] = systemDescr.SysDescr
    return leafs
}

func (systemDescr *Snmp_Information_SystemDescr) GetBundleName() string { return "cisco_ios_xr" }

func (systemDescr *Snmp_Information_SystemDescr) GetYangName() string { return "system-descr" }

func (systemDescr *Snmp_Information_SystemDescr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemDescr *Snmp_Information_SystemDescr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemDescr *Snmp_Information_SystemDescr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemDescr *Snmp_Information_SystemDescr) SetParent(parent types.Entity) { systemDescr.parent = parent }

func (systemDescr *Snmp_Information_SystemDescr) GetParent() types.Entity { return systemDescr.parent }

func (systemDescr *Snmp_Information_SystemDescr) GetParentYangName() string { return "information" }

// Snmp_Information_Tables
// List of table
type Snmp_Information_Tables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of vacmAccessTable.
    Groups Snmp_Information_Tables_Groups

    // List of User.
    UserEngineIds Snmp_Information_Tables_UserEngineIds
}

func (tables *Snmp_Information_Tables) GetFilter() yfilter.YFilter { return tables.YFilter }

func (tables *Snmp_Information_Tables) SetFilter(yf yfilter.YFilter) { tables.YFilter = yf }

func (tables *Snmp_Information_Tables) GetGoName(yname string) string {
    if yname == "groups" { return "Groups" }
    if yname == "user-engine-ids" { return "UserEngineIds" }
    return ""
}

func (tables *Snmp_Information_Tables) GetSegmentPath() string {
    return "tables"
}

func (tables *Snmp_Information_Tables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &tables.Groups
    }
    if childYangName == "user-engine-ids" {
        return &tables.UserEngineIds
    }
    return nil
}

func (tables *Snmp_Information_Tables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &tables.Groups
    children["user-engine-ids"] = &tables.UserEngineIds
    return children
}

func (tables *Snmp_Information_Tables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tables *Snmp_Information_Tables) GetBundleName() string { return "cisco_ios_xr" }

func (tables *Snmp_Information_Tables) GetYangName() string { return "tables" }

func (tables *Snmp_Information_Tables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tables *Snmp_Information_Tables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tables *Snmp_Information_Tables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tables *Snmp_Information_Tables) SetParent(parent types.Entity) { tables.parent = parent }

func (tables *Snmp_Information_Tables) GetParent() types.Entity { return tables.parent }

func (tables *Snmp_Information_Tables) GetParentYangName() string { return "information" }

// Snmp_Information_Tables_Groups
// List of vacmAccessTable
type Snmp_Information_Tables_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP group name. The type is slice of Snmp_Information_Tables_Groups_Group.
    Group []Snmp_Information_Tables_Groups_Group
}

func (groups *Snmp_Information_Tables_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Snmp_Information_Tables_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Snmp_Information_Tables_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Snmp_Information_Tables_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Snmp_Information_Tables_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Tables_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Snmp_Information_Tables_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Snmp_Information_Tables_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Snmp_Information_Tables_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Snmp_Information_Tables_Groups) GetYangName() string { return "groups" }

func (groups *Snmp_Information_Tables_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Snmp_Information_Tables_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Snmp_Information_Tables_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Snmp_Information_Tables_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Snmp_Information_Tables_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Snmp_Information_Tables_Groups) GetParentYangName() string { return "tables" }

// Snmp_Information_Tables_Groups_Group
// SNMP group name
type Snmp_Information_Tables_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Group Model.
    GroupInformations Snmp_Information_Tables_Groups_Group_GroupInformations
}

func (group *Snmp_Information_Tables_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Snmp_Information_Tables_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Snmp_Information_Tables_Groups_Group) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "group-informations" { return "GroupInformations" }
    return ""
}

func (group *Snmp_Information_Tables_Groups_Group) GetSegmentPath() string {
    return "group" + "[name='" + fmt.Sprintf("%v", group.Name) + "']"
}

func (group *Snmp_Information_Tables_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-informations" {
        return &group.GroupInformations
    }
    return nil
}

func (group *Snmp_Information_Tables_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["group-informations"] = &group.GroupInformations
    return children
}

func (group *Snmp_Information_Tables_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = group.Name
    return leafs
}

func (group *Snmp_Information_Tables_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Snmp_Information_Tables_Groups_Group) GetYangName() string { return "group" }

func (group *Snmp_Information_Tables_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Snmp_Information_Tables_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Snmp_Information_Tables_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Snmp_Information_Tables_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Snmp_Information_Tables_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Snmp_Information_Tables_Groups_Group) GetParentYangName() string { return "groups" }

// Snmp_Information_Tables_Groups_Group_GroupInformations
// Group Model
type Snmp_Information_Tables_Groups_Group_GroupInformations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group name ,status  and information. The type is slice of
    // Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation.
    GroupInformation []Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation
}

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetFilter() yfilter.YFilter { return groupInformations.YFilter }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) SetFilter(yf yfilter.YFilter) { groupInformations.YFilter = yf }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetGoName(yname string) string {
    if yname == "group-information" { return "GroupInformation" }
    return ""
}

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetSegmentPath() string {
    return "group-informations"
}

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-information" {
        for _, c := range groupInformations.GroupInformation {
            if groupInformations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation{}
        groupInformations.GroupInformation = append(groupInformations.GroupInformation, child)
        return &groupInformations.GroupInformation[len(groupInformations.GroupInformation)-1]
    }
    return nil
}

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groupInformations.GroupInformation {
        children[groupInformations.GroupInformation[i].GetSegmentPath()] = &groupInformations.GroupInformation[i]
    }
    return children
}

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetBundleName() string { return "cisco_ios_xr" }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetYangName() string { return "group-informations" }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) SetParent(parent types.Entity) { groupInformations.parent = parent }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetParent() types.Entity { return groupInformations.parent }

func (groupInformations *Snmp_Information_Tables_Groups_Group_GroupInformations) GetParentYangName() string { return "group" }

// Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation
// Group name ,status  and information
type Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetFilter() yfilter.YFilter { return groupInformation.YFilter }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) SetFilter(yf yfilter.YFilter) { groupInformation.YFilter = yf }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetGoName(yname string) string {
    if yname == "modelnumber" { return "Modelnumber" }
    if yname == "level" { return "Level" }
    if yname == "vacm-access-read-view-name" { return "VacmAccessReadViewName" }
    if yname == "vacm-access-write-view-name" { return "VacmAccessWriteViewName" }
    if yname == "vacm-access-notify-view-name" { return "VacmAccessNotifyViewName" }
    if yname == "vacm-access-status" { return "VacmAccessStatus" }
    return ""
}

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetSegmentPath() string {
    return "group-information"
}

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["modelnumber"] = groupInformation.Modelnumber
    leafs["level"] = groupInformation.Level
    leafs["vacm-access-read-view-name"] = groupInformation.VacmAccessReadViewName
    leafs["vacm-access-write-view-name"] = groupInformation.VacmAccessWriteViewName
    leafs["vacm-access-notify-view-name"] = groupInformation.VacmAccessNotifyViewName
    leafs["vacm-access-status"] = groupInformation.VacmAccessStatus
    return leafs
}

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetBundleName() string { return "cisco_ios_xr" }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetYangName() string { return "group-information" }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) SetParent(parent types.Entity) { groupInformation.parent = parent }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetParent() types.Entity { return groupInformation.parent }

func (groupInformation *Snmp_Information_Tables_Groups_Group_GroupInformations_GroupInformation) GetParentYangName() string { return "group-informations" }

// Snmp_Information_Tables_UserEngineIds
// List of User
type Snmp_Information_Tables_UserEngineIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP engineId. The type is slice of
    // Snmp_Information_Tables_UserEngineIds_UserEngineId.
    UserEngineId []Snmp_Information_Tables_UserEngineIds_UserEngineId
}

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetFilter() yfilter.YFilter { return userEngineIds.YFilter }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) SetFilter(yf yfilter.YFilter) { userEngineIds.YFilter = yf }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetGoName(yname string) string {
    if yname == "user-engine-id" { return "UserEngineId" }
    return ""
}

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetSegmentPath() string {
    return "user-engine-ids"
}

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "user-engine-id" {
        for _, c := range userEngineIds.UserEngineId {
            if userEngineIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Tables_UserEngineIds_UserEngineId{}
        userEngineIds.UserEngineId = append(userEngineIds.UserEngineId, child)
        return &userEngineIds.UserEngineId[len(userEngineIds.UserEngineId)-1]
    }
    return nil
}

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range userEngineIds.UserEngineId {
        children[userEngineIds.UserEngineId[i].GetSegmentPath()] = &userEngineIds.UserEngineId[i]
    }
    return children
}

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetBundleName() string { return "cisco_ios_xr" }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetYangName() string { return "user-engine-ids" }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) SetParent(parent types.Entity) { userEngineIds.parent = parent }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetParent() types.Entity { return userEngineIds.parent }

func (userEngineIds *Snmp_Information_Tables_UserEngineIds) GetParentYangName() string { return "tables" }

// Snmp_Information_Tables_UserEngineIds_UserEngineId
// SNMP engineId
type Snmp_Information_Tables_UserEngineIds_UserEngineId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP Engine ID. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    EngineId interface{}

    // User name ,storage type ,status . The type is slice of
    // Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName.
    UserName []Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName
}

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetFilter() yfilter.YFilter { return userEngineId.YFilter }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) SetFilter(yf yfilter.YFilter) { userEngineId.YFilter = yf }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetGoName(yname string) string {
    if yname == "engine-id" { return "EngineId" }
    if yname == "user-name" { return "UserName" }
    return ""
}

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetSegmentPath() string {
    return "user-engine-id" + "[engine-id='" + fmt.Sprintf("%v", userEngineId.EngineId) + "']"
}

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "user-name" {
        for _, c := range userEngineId.UserName {
            if userEngineId.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName{}
        userEngineId.UserName = append(userEngineId.UserName, child)
        return &userEngineId.UserName[len(userEngineId.UserName)-1]
    }
    return nil
}

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range userEngineId.UserName {
        children[userEngineId.UserName[i].GetSegmentPath()] = &userEngineId.UserName[i]
    }
    return children
}

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["engine-id"] = userEngineId.EngineId
    return leafs
}

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetBundleName() string { return "cisco_ios_xr" }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetYangName() string { return "user-engine-id" }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) SetParent(parent types.Entity) { userEngineId.parent = parent }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetParent() types.Entity { return userEngineId.parent }

func (userEngineId *Snmp_Information_Tables_UserEngineIds_UserEngineId) GetParentYangName() string { return "user-engine-ids" }

// Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName
// User name ,storage type ,status 
type Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. User name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    UserName interface{}

    // Storage type. The type is interface{} with range: 0..4294967295.
    UsmUserStorageType interface{}

    // Status of this user. The type is interface{} with range: 0..4294967295.
    UsmUserStatus interface{}
}

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetFilter() yfilter.YFilter { return userName.YFilter }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) SetFilter(yf yfilter.YFilter) { userName.YFilter = yf }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetGoName(yname string) string {
    if yname == "user-name" { return "UserName" }
    if yname == "usm-user-storage-type" { return "UsmUserStorageType" }
    if yname == "usm-user-status" { return "UsmUserStatus" }
    return ""
}

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetSegmentPath() string {
    return "user-name" + "[user-name='" + fmt.Sprintf("%v", userName.UserName) + "']"
}

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["user-name"] = userName.UserName
    leafs["usm-user-storage-type"] = userName.UsmUserStorageType
    leafs["usm-user-status"] = userName.UsmUserStatus
    return leafs
}

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetBundleName() string { return "cisco_ios_xr" }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetYangName() string { return "user-name" }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) SetParent(parent types.Entity) { userName.parent = parent }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetParent() types.Entity { return userName.parent }

func (userName *Snmp_Information_Tables_UserEngineIds_UserEngineId_UserName) GetParentYangName() string { return "user-engine-id" }

// Snmp_Information_SystemOid
// System object ID
type Snmp_Information_SystemOid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // sysObjID  1.3.6.1.2.1.1.2. The type is string.
    SysObjId interface{}
}

func (systemOid *Snmp_Information_SystemOid) GetFilter() yfilter.YFilter { return systemOid.YFilter }

func (systemOid *Snmp_Information_SystemOid) SetFilter(yf yfilter.YFilter) { systemOid.YFilter = yf }

func (systemOid *Snmp_Information_SystemOid) GetGoName(yname string) string {
    if yname == "sys-obj-id" { return "SysObjId" }
    return ""
}

func (systemOid *Snmp_Information_SystemOid) GetSegmentPath() string {
    return "system-oid"
}

func (systemOid *Snmp_Information_SystemOid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemOid *Snmp_Information_SystemOid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemOid *Snmp_Information_SystemOid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sys-obj-id"] = systemOid.SysObjId
    return leafs
}

func (systemOid *Snmp_Information_SystemOid) GetBundleName() string { return "cisco_ios_xr" }

func (systemOid *Snmp_Information_SystemOid) GetYangName() string { return "system-oid" }

func (systemOid *Snmp_Information_SystemOid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemOid *Snmp_Information_SystemOid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemOid *Snmp_Information_SystemOid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemOid *Snmp_Information_SystemOid) SetParent(parent types.Entity) { systemOid.parent = parent }

func (systemOid *Snmp_Information_SystemOid) GetParent() types.Entity { return systemOid.parent }

func (systemOid *Snmp_Information_SystemOid) GetParentYangName() string { return "information" }

// Snmp_Information_TrapQueue
// SNMP trap queue statistics
type Snmp_Information_TrapQueue struct {
    parent types.Entity
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

func (trapQueue *Snmp_Information_TrapQueue) GetFilter() yfilter.YFilter { return trapQueue.YFilter }

func (trapQueue *Snmp_Information_TrapQueue) SetFilter(yf yfilter.YFilter) { trapQueue.YFilter = yf }

func (trapQueue *Snmp_Information_TrapQueue) GetGoName(yname string) string {
    if yname == "trap-min" { return "TrapMin" }
    if yname == "trap-avg" { return "TrapAvg" }
    if yname == "trap-max" { return "TrapMax" }
    if yname == "trap-q" { return "TrapQ" }
    return ""
}

func (trapQueue *Snmp_Information_TrapQueue) GetSegmentPath() string {
    return "trap-queue"
}

func (trapQueue *Snmp_Information_TrapQueue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-q" {
        for _, c := range trapQueue.TrapQ {
            if trapQueue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Information_TrapQueue_TrapQ{}
        trapQueue.TrapQ = append(trapQueue.TrapQ, child)
        return &trapQueue.TrapQ[len(trapQueue.TrapQ)-1]
    }
    return nil
}

func (trapQueue *Snmp_Information_TrapQueue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapQueue.TrapQ {
        children[trapQueue.TrapQ[i].GetSegmentPath()] = &trapQueue.TrapQ[i]
    }
    return children
}

func (trapQueue *Snmp_Information_TrapQueue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-min"] = trapQueue.TrapMin
    leafs["trap-avg"] = trapQueue.TrapAvg
    leafs["trap-max"] = trapQueue.TrapMax
    return leafs
}

func (trapQueue *Snmp_Information_TrapQueue) GetBundleName() string { return "cisco_ios_xr" }

func (trapQueue *Snmp_Information_TrapQueue) GetYangName() string { return "trap-queue" }

func (trapQueue *Snmp_Information_TrapQueue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapQueue *Snmp_Information_TrapQueue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapQueue *Snmp_Information_TrapQueue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapQueue *Snmp_Information_TrapQueue) SetParent(parent types.Entity) { trapQueue.parent = parent }

func (trapQueue *Snmp_Information_TrapQueue) GetParent() types.Entity { return trapQueue.parent }

func (trapQueue *Snmp_Information_TrapQueue) GetParentYangName() string { return "information" }

// Snmp_Information_TrapQueue_TrapQ
// trap q
type Snmp_Information_TrapQueue_TrapQ struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // min. The type is interface{} with range: 0..4294967295.
    Min interface{}

    // avg. The type is interface{} with range: 0..4294967295.
    Avg interface{}

    // max. The type is interface{} with range: 0..4294967295.
    Max interface{}
}

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetFilter() yfilter.YFilter { return trapQ.YFilter }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) SetFilter(yf yfilter.YFilter) { trapQ.YFilter = yf }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetGoName(yname string) string {
    if yname == "min" { return "Min" }
    if yname == "avg" { return "Avg" }
    if yname == "max" { return "Max" }
    return ""
}

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetSegmentPath() string {
    return "trap-q"
}

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min"] = trapQ.Min
    leafs["avg"] = trapQ.Avg
    leafs["max"] = trapQ.Max
    return leafs
}

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetBundleName() string { return "cisco_ios_xr" }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetYangName() string { return "trap-q" }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) SetParent(parent types.Entity) { trapQ.parent = parent }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetParent() types.Entity { return trapQ.parent }

func (trapQ *Snmp_Information_TrapQueue_TrapQ) GetParentYangName() string { return "trap-queue" }

// Snmp_Interfaces
// List of interfaces
type Snmp_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is slice of Snmp_Interfaces_Interface.
    Interface []Snmp_Interfaces_Interface
}

func (interfaces *Snmp_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Snmp_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Snmp_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Snmp_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Snmp_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Snmp_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Snmp_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Snmp_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Snmp_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Snmp_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Snmp_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Snmp_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Snmp_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Snmp_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Snmp_Interfaces) GetParentYangName() string { return "snmp" }

// Snmp_Interfaces_Interface
// Interface Name
type Snmp_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // Interface Index as used by MIB tables. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    InterfaceIndex interface{}
}

func (self *Snmp_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Snmp_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Snmp_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "interface-index" { return "InterfaceIndex" }
    return ""
}

func (self *Snmp_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Snmp_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Snmp_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Snmp_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    leafs["interface-index"] = self.InterfaceIndex
    return leafs
}

func (self *Snmp_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Snmp_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Snmp_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Snmp_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Snmp_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Snmp_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Snmp_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Snmp_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Snmp_Correlator
// Trap Correlator operational data
type Snmp_Correlator struct {
    parent types.Entity
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

func (correlator *Snmp_Correlator) GetFilter() yfilter.YFilter { return correlator.YFilter }

func (correlator *Snmp_Correlator) SetFilter(yf yfilter.YFilter) { correlator.YFilter = yf }

func (correlator *Snmp_Correlator) GetGoName(yname string) string {
    if yname == "rule-details" { return "RuleDetails" }
    if yname == "buffer-status" { return "BufferStatus" }
    if yname == "rule-set-details" { return "RuleSetDetails" }
    if yname == "traps" { return "Traps" }
    return ""
}

func (correlator *Snmp_Correlator) GetSegmentPath() string {
    return "correlator"
}

func (correlator *Snmp_Correlator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-details" {
        return &correlator.RuleDetails
    }
    if childYangName == "buffer-status" {
        return &correlator.BufferStatus
    }
    if childYangName == "rule-set-details" {
        return &correlator.RuleSetDetails
    }
    if childYangName == "traps" {
        return &correlator.Traps
    }
    return nil
}

func (correlator *Snmp_Correlator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rule-details"] = &correlator.RuleDetails
    children["buffer-status"] = &correlator.BufferStatus
    children["rule-set-details"] = &correlator.RuleSetDetails
    children["traps"] = &correlator.Traps
    return children
}

func (correlator *Snmp_Correlator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (correlator *Snmp_Correlator) GetBundleName() string { return "cisco_ios_xr" }

func (correlator *Snmp_Correlator) GetYangName() string { return "correlator" }

func (correlator *Snmp_Correlator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (correlator *Snmp_Correlator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (correlator *Snmp_Correlator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (correlator *Snmp_Correlator) SetParent(parent types.Entity) { correlator.parent = parent }

func (correlator *Snmp_Correlator) GetParent() types.Entity { return correlator.parent }

func (correlator *Snmp_Correlator) GetParentYangName() string { return "snmp" }

// Snmp_Correlator_RuleDetails
// Table that contains the database of correlation
// rule details
type Snmp_Correlator_RuleDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details of one of the correlation rules. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail.
    RuleDetail []Snmp_Correlator_RuleDetails_RuleDetail
}

func (ruleDetails *Snmp_Correlator_RuleDetails) GetFilter() yfilter.YFilter { return ruleDetails.YFilter }

func (ruleDetails *Snmp_Correlator_RuleDetails) SetFilter(yf yfilter.YFilter) { ruleDetails.YFilter = yf }

func (ruleDetails *Snmp_Correlator_RuleDetails) GetGoName(yname string) string {
    if yname == "rule-detail" { return "RuleDetail" }
    return ""
}

func (ruleDetails *Snmp_Correlator_RuleDetails) GetSegmentPath() string {
    return "rule-details"
}

func (ruleDetails *Snmp_Correlator_RuleDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-detail" {
        for _, c := range ruleDetails.RuleDetail {
            if ruleDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleDetails_RuleDetail{}
        ruleDetails.RuleDetail = append(ruleDetails.RuleDetail, child)
        return &ruleDetails.RuleDetail[len(ruleDetails.RuleDetail)-1]
    }
    return nil
}

func (ruleDetails *Snmp_Correlator_RuleDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleDetails.RuleDetail {
        children[ruleDetails.RuleDetail[i].GetSegmentPath()] = &ruleDetails.RuleDetail[i]
    }
    return children
}

func (ruleDetails *Snmp_Correlator_RuleDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleDetails *Snmp_Correlator_RuleDetails) GetBundleName() string { return "cisco_ios_xr" }

func (ruleDetails *Snmp_Correlator_RuleDetails) GetYangName() string { return "rule-details" }

func (ruleDetails *Snmp_Correlator_RuleDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleDetails *Snmp_Correlator_RuleDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleDetails *Snmp_Correlator_RuleDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleDetails *Snmp_Correlator_RuleDetails) SetParent(parent types.Entity) { ruleDetails.parent = parent }

func (ruleDetails *Snmp_Correlator_RuleDetails) GetParent() types.Entity { return ruleDetails.parent }

func (ruleDetails *Snmp_Correlator_RuleDetails) GetParentYangName() string { return "correlator" }

// Snmp_Correlator_RuleDetails_RuleDetail
// Details of one of the correlation rules
type Snmp_Correlator_RuleDetails_RuleDetail struct {
    parent types.Entity
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

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetFilter() yfilter.YFilter { return ruleDetail.YFilter }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) SetFilter(yf yfilter.YFilter) { ruleDetail.YFilter = yf }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "timeout" { return "Timeout" }
    if yname == "rule-summary" { return "RuleSummary" }
    if yname == "root-cause" { return "RootCause" }
    if yname == "non-rootcaus" { return "NonRootcaus" }
    if yname == "apply-host" { return "ApplyHost" }
    return ""
}

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetSegmentPath() string {
    return "rule-detail" + "[rule-name='" + fmt.Sprintf("%v", ruleDetail.RuleName) + "']"
}

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-summary" {
        return &ruleDetail.RuleSummary
    }
    if childYangName == "root-cause" {
        return &ruleDetail.RootCause
    }
    if childYangName == "non-rootcaus" {
        for _, c := range ruleDetail.NonRootcaus {
            if ruleDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus{}
        ruleDetail.NonRootcaus = append(ruleDetail.NonRootcaus, child)
        return &ruleDetail.NonRootcaus[len(ruleDetail.NonRootcaus)-1]
    }
    if childYangName == "apply-host" {
        for _, c := range ruleDetail.ApplyHost {
            if ruleDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost{}
        ruleDetail.ApplyHost = append(ruleDetail.ApplyHost, child)
        return &ruleDetail.ApplyHost[len(ruleDetail.ApplyHost)-1]
    }
    return nil
}

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rule-summary"] = &ruleDetail.RuleSummary
    children["root-cause"] = &ruleDetail.RootCause
    for i := range ruleDetail.NonRootcaus {
        children[ruleDetail.NonRootcaus[i].GetSegmentPath()] = &ruleDetail.NonRootcaus[i]
    }
    for i := range ruleDetail.ApplyHost {
        children[ruleDetail.ApplyHost[i].GetSegmentPath()] = &ruleDetail.ApplyHost[i]
    }
    return children
}

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = ruleDetail.RuleName
    leafs["timeout"] = ruleDetail.Timeout
    return leafs
}

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetBundleName() string { return "cisco_ios_xr" }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetYangName() string { return "rule-detail" }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) SetParent(parent types.Entity) { ruleDetail.parent = parent }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetParent() types.Entity { return ruleDetail.parent }

func (ruleDetail *Snmp_Correlator_RuleDetails_RuleDetail) GetParentYangName() string { return "rule-details" }

// Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary
// Rule summary, name, etc
type Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary struct {
    parent types.Entity
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

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetFilter() yfilter.YFilter { return ruleSummary.YFilter }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) SetFilter(yf yfilter.YFilter) { ruleSummary.YFilter = yf }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "rule-state" { return "RuleState" }
    if yname == "buffered-traps-count" { return "BufferedTrapsCount" }
    return ""
}

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetSegmentPath() string {
    return "rule-summary"
}

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = ruleSummary.RuleName
    leafs["rule-state"] = ruleSummary.RuleState
    leafs["buffered-traps-count"] = ruleSummary.BufferedTrapsCount
    return leafs
}

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetYangName() string { return "rule-summary" }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) SetParent(parent types.Entity) { ruleSummary.parent = parent }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetParent() types.Entity { return ruleSummary.parent }

func (ruleSummary *Snmp_Correlator_RuleDetails_RuleDetail_RuleSummary) GetParentYangName() string { return "rule-detail" }

// Snmp_Correlator_RuleDetails_RuleDetail_RootCause
// OID/VarBinds defining the rootcause match
// conditions.
type Snmp_Correlator_RuleDetails_RuleDetail_RootCause struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OID of the trap. The type is string.
    Oid interface{}

    // VarBinds of the trap. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind.
    VarBind []Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind
}

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetFilter() yfilter.YFilter { return rootCause.YFilter }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) SetFilter(yf yfilter.YFilter) { rootCause.YFilter = yf }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "var-bind" { return "VarBind" }
    return ""
}

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetSegmentPath() string {
    return "root-cause"
}

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "var-bind" {
        for _, c := range rootCause.VarBind {
            if rootCause.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind{}
        rootCause.VarBind = append(rootCause.VarBind, child)
        return &rootCause.VarBind[len(rootCause.VarBind)-1]
    }
    return nil
}

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rootCause.VarBind {
        children[rootCause.VarBind[i].GetSegmentPath()] = &rootCause.VarBind[i]
    }
    return children
}

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = rootCause.Oid
    return leafs
}

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetBundleName() string { return "cisco_ios_xr" }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetYangName() string { return "root-cause" }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) SetParent(parent types.Entity) { rootCause.parent = parent }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetParent() types.Entity { return rootCause.parent }

func (rootCause *Snmp_Correlator_RuleDetails_RuleDetail_RootCause) GetParentYangName() string { return "rule-detail" }

// Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind
// VarBinds of the trap
type Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OID of the varbind. The type is string.
    Oid interface{}

    // Varbind match type. The type is SnmpCorrVbindMatch.
    MatchType interface{}

    // Regular expression to match. The type is string.
    RegExp interface{}
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetFilter() yfilter.YFilter { return varBind.YFilter }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) SetFilter(yf yfilter.YFilter) { varBind.YFilter = yf }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "match-type" { return "MatchType" }
    if yname == "reg-exp" { return "RegExp" }
    return ""
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetSegmentPath() string {
    return "var-bind"
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = varBind.Oid
    leafs["match-type"] = varBind.MatchType
    leafs["reg-exp"] = varBind.RegExp
    return leafs
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetBundleName() string { return "cisco_ios_xr" }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetYangName() string { return "var-bind" }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) SetParent(parent types.Entity) { varBind.parent = parent }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetParent() types.Entity { return varBind.parent }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_RootCause_VarBind) GetParentYangName() string { return "root-cause" }

// Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus
// OIDs/VarBinds defining the nonrootcause match
// conditions.
type Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OID of the trap. The type is string.
    Oid interface{}

    // VarBinds of the trap. The type is slice of
    // Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind.
    VarBind []Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind
}

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetFilter() yfilter.YFilter { return nonRootcaus.YFilter }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) SetFilter(yf yfilter.YFilter) { nonRootcaus.YFilter = yf }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "var-bind" { return "VarBind" }
    return ""
}

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetSegmentPath() string {
    return "non-rootcaus"
}

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "var-bind" {
        for _, c := range nonRootcaus.VarBind {
            if nonRootcaus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind{}
        nonRootcaus.VarBind = append(nonRootcaus.VarBind, child)
        return &nonRootcaus.VarBind[len(nonRootcaus.VarBind)-1]
    }
    return nil
}

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nonRootcaus.VarBind {
        children[nonRootcaus.VarBind[i].GetSegmentPath()] = &nonRootcaus.VarBind[i]
    }
    return children
}

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = nonRootcaus.Oid
    return leafs
}

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetBundleName() string { return "cisco_ios_xr" }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetYangName() string { return "non-rootcaus" }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) SetParent(parent types.Entity) { nonRootcaus.parent = parent }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetParent() types.Entity { return nonRootcaus.parent }

func (nonRootcaus *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus) GetParentYangName() string { return "rule-detail" }

// Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind
// VarBinds of the trap
type Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OID of the varbind. The type is string.
    Oid interface{}

    // Varbind match type. The type is SnmpCorrVbindMatch.
    MatchType interface{}

    // Regular expression to match. The type is string.
    RegExp interface{}
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetFilter() yfilter.YFilter { return varBind.YFilter }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) SetFilter(yf yfilter.YFilter) { varBind.YFilter = yf }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "match-type" { return "MatchType" }
    if yname == "reg-exp" { return "RegExp" }
    return ""
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetSegmentPath() string {
    return "var-bind"
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = varBind.Oid
    leafs["match-type"] = varBind.MatchType
    leafs["reg-exp"] = varBind.RegExp
    return leafs
}

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetBundleName() string { return "cisco_ios_xr" }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetYangName() string { return "var-bind" }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) SetParent(parent types.Entity) { varBind.parent = parent }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetParent() types.Entity { return varBind.parent }

func (varBind *Snmp_Correlator_RuleDetails_RuleDetail_NonRootcaus_VarBind) GetParentYangName() string { return "non-rootcaus" }

// Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost
// Hosts (IP/port) to which the rule is applied
type Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of the host. The type is string.
    IpAddress interface{}

    // Port of the host. The type is interface{} with range: 0..65535.
    Port interface{}
}

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetFilter() yfilter.YFilter { return applyHost.YFilter }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) SetFilter(yf yfilter.YFilter) { applyHost.YFilter = yf }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "port" { return "Port" }
    return ""
}

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetSegmentPath() string {
    return "apply-host"
}

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = applyHost.IpAddress
    leafs["port"] = applyHost.Port
    return leafs
}

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetBundleName() string { return "cisco_ios_xr" }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetYangName() string { return "apply-host" }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) SetParent(parent types.Entity) { applyHost.parent = parent }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetParent() types.Entity { return applyHost.parent }

func (applyHost *Snmp_Correlator_RuleDetails_RuleDetail_ApplyHost) GetParentYangName() string { return "rule-detail" }

// Snmp_Correlator_BufferStatus
// Describes buffer utilization and parameters
// configured
type Snmp_Correlator_BufferStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current buffer usage. The type is interface{} with range: 0..4294967295.
    CurrentSize interface{}

    // Configured buffer size. The type is interface{} with range: 0..4294967295.
    ConfiguredSize interface{}
}

func (bufferStatus *Snmp_Correlator_BufferStatus) GetFilter() yfilter.YFilter { return bufferStatus.YFilter }

func (bufferStatus *Snmp_Correlator_BufferStatus) SetFilter(yf yfilter.YFilter) { bufferStatus.YFilter = yf }

func (bufferStatus *Snmp_Correlator_BufferStatus) GetGoName(yname string) string {
    if yname == "current-size" { return "CurrentSize" }
    if yname == "configured-size" { return "ConfiguredSize" }
    return ""
}

func (bufferStatus *Snmp_Correlator_BufferStatus) GetSegmentPath() string {
    return "buffer-status"
}

func (bufferStatus *Snmp_Correlator_BufferStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bufferStatus *Snmp_Correlator_BufferStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bufferStatus *Snmp_Correlator_BufferStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-size"] = bufferStatus.CurrentSize
    leafs["configured-size"] = bufferStatus.ConfiguredSize
    return leafs
}

func (bufferStatus *Snmp_Correlator_BufferStatus) GetBundleName() string { return "cisco_ios_xr" }

func (bufferStatus *Snmp_Correlator_BufferStatus) GetYangName() string { return "buffer-status" }

func (bufferStatus *Snmp_Correlator_BufferStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bufferStatus *Snmp_Correlator_BufferStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bufferStatus *Snmp_Correlator_BufferStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bufferStatus *Snmp_Correlator_BufferStatus) SetParent(parent types.Entity) { bufferStatus.parent = parent }

func (bufferStatus *Snmp_Correlator_BufferStatus) GetParent() types.Entity { return bufferStatus.parent }

func (bufferStatus *Snmp_Correlator_BufferStatus) GetParentYangName() string { return "correlator" }

// Snmp_Correlator_RuleSetDetails
// Table that contains the ruleset detail info
type Snmp_Correlator_RuleSetDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detail of one of the correlation rulesets. The type is slice of
    // Snmp_Correlator_RuleSetDetails_RuleSetDetail.
    RuleSetDetail []Snmp_Correlator_RuleSetDetails_RuleSetDetail
}

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetFilter() yfilter.YFilter { return ruleSetDetails.YFilter }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) SetFilter(yf yfilter.YFilter) { ruleSetDetails.YFilter = yf }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetGoName(yname string) string {
    if yname == "rule-set-detail" { return "RuleSetDetail" }
    return ""
}

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetSegmentPath() string {
    return "rule-set-details"
}

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-set-detail" {
        for _, c := range ruleSetDetails.RuleSetDetail {
            if ruleSetDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleSetDetails_RuleSetDetail{}
        ruleSetDetails.RuleSetDetail = append(ruleSetDetails.RuleSetDetail, child)
        return &ruleSetDetails.RuleSetDetail[len(ruleSetDetails.RuleSetDetail)-1]
    }
    return nil
}

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSetDetails.RuleSetDetail {
        children[ruleSetDetails.RuleSetDetail[i].GetSegmentPath()] = &ruleSetDetails.RuleSetDetail[i]
    }
    return children
}

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetYangName() string { return "rule-set-details" }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) SetParent(parent types.Entity) { ruleSetDetails.parent = parent }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetParent() types.Entity { return ruleSetDetails.parent }

func (ruleSetDetails *Snmp_Correlator_RuleSetDetails) GetParentYangName() string { return "correlator" }

// Snmp_Correlator_RuleSetDetails_RuleSetDetail
// Detail of one of the correlation rulesets
type Snmp_Correlator_RuleSetDetails_RuleSetDetail struct {
    parent types.Entity
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

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetFilter() yfilter.YFilter { return ruleSetDetail.YFilter }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) SetFilter(yf yfilter.YFilter) { ruleSetDetail.YFilter = yf }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetGoName(yname string) string {
    if yname == "rule-set-name" { return "RuleSetName" }
    if yname == "rule-set-name-xr" { return "RuleSetNameXr" }
    if yname == "rules" { return "Rules" }
    return ""
}

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetSegmentPath() string {
    return "rule-set-detail" + "[rule-set-name='" + fmt.Sprintf("%v", ruleSetDetail.RuleSetName) + "']"
}

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rules" {
        for _, c := range ruleSetDetail.Rules {
            if ruleSetDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules{}
        ruleSetDetail.Rules = append(ruleSetDetail.Rules, child)
        return &ruleSetDetail.Rules[len(ruleSetDetail.Rules)-1]
    }
    return nil
}

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSetDetail.Rules {
        children[ruleSetDetail.Rules[i].GetSegmentPath()] = &ruleSetDetail.Rules[i]
    }
    return children
}

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-set-name"] = ruleSetDetail.RuleSetName
    leafs["rule-set-name-xr"] = ruleSetDetail.RuleSetNameXr
    return leafs
}

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetYangName() string { return "rule-set-detail" }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) SetParent(parent types.Entity) { ruleSetDetail.parent = parent }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetParent() types.Entity { return ruleSetDetail.parent }

func (ruleSetDetail *Snmp_Correlator_RuleSetDetails_RuleSetDetail) GetParentYangName() string { return "rule-set-details" }

// Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules
// Rules contained in a ruleset
type Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules struct {
    parent types.Entity
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

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetFilter() yfilter.YFilter { return rules.YFilter }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) SetFilter(yf yfilter.YFilter) { rules.YFilter = yf }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "rule-state" { return "RuleState" }
    if yname == "buffered-traps-count" { return "BufferedTrapsCount" }
    return ""
}

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetSegmentPath() string {
    return "rules"
}

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = rules.RuleName
    leafs["rule-state"] = rules.RuleState
    leafs["buffered-traps-count"] = rules.BufferedTrapsCount
    return leafs
}

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetBundleName() string { return "cisco_ios_xr" }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetYangName() string { return "rules" }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) SetParent(parent types.Entity) { rules.parent = parent }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetParent() types.Entity { return rules.parent }

func (rules *Snmp_Correlator_RuleSetDetails_RuleSetDetail_Rules) GetParentYangName() string { return "rule-set-detail" }

// Snmp_Correlator_Traps
// Correlated traps Table
type Snmp_Correlator_Traps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One of the correlated traps. The type is slice of
    // Snmp_Correlator_Traps_Trap.
    Trap []Snmp_Correlator_Traps_Trap
}

func (traps *Snmp_Correlator_Traps) GetFilter() yfilter.YFilter { return traps.YFilter }

func (traps *Snmp_Correlator_Traps) SetFilter(yf yfilter.YFilter) { traps.YFilter = yf }

func (traps *Snmp_Correlator_Traps) GetGoName(yname string) string {
    if yname == "trap" { return "Trap" }
    return ""
}

func (traps *Snmp_Correlator_Traps) GetSegmentPath() string {
    return "traps"
}

func (traps *Snmp_Correlator_Traps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap" {
        for _, c := range traps.Trap {
            if traps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_Traps_Trap{}
        traps.Trap = append(traps.Trap, child)
        return &traps.Trap[len(traps.Trap)-1]
    }
    return nil
}

func (traps *Snmp_Correlator_Traps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range traps.Trap {
        children[traps.Trap[i].GetSegmentPath()] = &traps.Trap[i]
    }
    return children
}

func (traps *Snmp_Correlator_Traps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (traps *Snmp_Correlator_Traps) GetBundleName() string { return "cisco_ios_xr" }

func (traps *Snmp_Correlator_Traps) GetYangName() string { return "traps" }

func (traps *Snmp_Correlator_Traps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traps *Snmp_Correlator_Traps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traps *Snmp_Correlator_Traps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traps *Snmp_Correlator_Traps) SetParent(parent types.Entity) { traps.parent = parent }

func (traps *Snmp_Correlator_Traps) GetParent() types.Entity { return traps.parent }

func (traps *Snmp_Correlator_Traps) GetParentYangName() string { return "correlator" }

// Snmp_Correlator_Traps_Trap
// One of the correlated traps
type Snmp_Correlator_Traps_Trap struct {
    parent types.Entity
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

func (trap *Snmp_Correlator_Traps_Trap) GetFilter() yfilter.YFilter { return trap.YFilter }

func (trap *Snmp_Correlator_Traps_Trap) SetFilter(yf yfilter.YFilter) { trap.YFilter = yf }

func (trap *Snmp_Correlator_Traps_Trap) GetGoName(yname string) string {
    if yname == "entry-id" { return "EntryId" }
    if yname == "correlation-id" { return "CorrelationId" }
    if yname == "is-root-cause" { return "IsRootCause" }
    if yname == "rule-name" { return "RuleName" }
    if yname == "trap-info" { return "TrapInfo" }
    return ""
}

func (trap *Snmp_Correlator_Traps_Trap) GetSegmentPath() string {
    return "trap" + "[entry-id='" + fmt.Sprintf("%v", trap.EntryId) + "']"
}

func (trap *Snmp_Correlator_Traps_Trap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-info" {
        return &trap.TrapInfo
    }
    return nil
}

func (trap *Snmp_Correlator_Traps_Trap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["trap-info"] = &trap.TrapInfo
    return children
}

func (trap *Snmp_Correlator_Traps_Trap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entry-id"] = trap.EntryId
    leafs["correlation-id"] = trap.CorrelationId
    leafs["is-root-cause"] = trap.IsRootCause
    leafs["rule-name"] = trap.RuleName
    return leafs
}

func (trap *Snmp_Correlator_Traps_Trap) GetBundleName() string { return "cisco_ios_xr" }

func (trap *Snmp_Correlator_Traps_Trap) GetYangName() string { return "trap" }

func (trap *Snmp_Correlator_Traps_Trap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trap *Snmp_Correlator_Traps_Trap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trap *Snmp_Correlator_Traps_Trap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trap *Snmp_Correlator_Traps_Trap) SetParent(parent types.Entity) { trap.parent = parent }

func (trap *Snmp_Correlator_Traps_Trap) GetParent() types.Entity { return trap.parent }

func (trap *Snmp_Correlator_Traps_Trap) GetParentYangName() string { return "traps" }

// Snmp_Correlator_Traps_Trap_TrapInfo
// Correlated trap information
type Snmp_Correlator_Traps_Trap_TrapInfo struct {
    parent types.Entity
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

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetFilter() yfilter.YFilter { return trapInfo.YFilter }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) SetFilter(yf yfilter.YFilter) { trapInfo.YFilter = yf }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "relative-timestamp" { return "RelativeTimestamp" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "var-bind" { return "VarBind" }
    return ""
}

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetSegmentPath() string {
    return "trap-info"
}

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "var-bind" {
        for _, c := range trapInfo.VarBind {
            if trapInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_Traps_Trap_TrapInfo_VarBind{}
        trapInfo.VarBind = append(trapInfo.VarBind, child)
        return &trapInfo.VarBind[len(trapInfo.VarBind)-1]
    }
    return nil
}

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapInfo.VarBind {
        children[trapInfo.VarBind[i].GetSegmentPath()] = &trapInfo.VarBind[i]
    }
    return children
}

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = trapInfo.Oid
    leafs["relative-timestamp"] = trapInfo.RelativeTimestamp
    leafs["timestamp"] = trapInfo.Timestamp
    return leafs
}

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetBundleName() string { return "cisco_ios_xr" }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetYangName() string { return "trap-info" }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) SetParent(parent types.Entity) { trapInfo.parent = parent }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetParent() types.Entity { return trapInfo.parent }

func (trapInfo *Snmp_Correlator_Traps_Trap_TrapInfo) GetParentYangName() string { return "trap" }

// Snmp_Correlator_Traps_Trap_TrapInfo_VarBind
// VarBinds on the trap
type Snmp_Correlator_Traps_Trap_TrapInfo_VarBind struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OID of the varbind. The type is string.
    Oid interface{}

    // Value of the varbind. The type is string.
    Value interface{}
}

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetFilter() yfilter.YFilter { return varBind.YFilter }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) SetFilter(yf yfilter.YFilter) { varBind.YFilter = yf }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "value" { return "Value" }
    return ""
}

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetSegmentPath() string {
    return "var-bind"
}

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = varBind.Oid
    leafs["value"] = varBind.Value
    return leafs
}

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetBundleName() string { return "cisco_ios_xr" }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetYangName() string { return "var-bind" }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) SetParent(parent types.Entity) { varBind.parent = parent }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetParent() types.Entity { return varBind.parent }

func (varBind *Snmp_Correlator_Traps_Trap_TrapInfo_VarBind) GetParentYangName() string { return "trap-info" }

// Snmp_InterfaceIndexes
// List of index
type Snmp_InterfaceIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Index. The type is slice of Snmp_InterfaceIndexes_InterfaceIndex.
    InterfaceIndex []Snmp_InterfaceIndexes_InterfaceIndex
}

func (interfaceIndexes *Snmp_InterfaceIndexes) GetFilter() yfilter.YFilter { return interfaceIndexes.YFilter }

func (interfaceIndexes *Snmp_InterfaceIndexes) SetFilter(yf yfilter.YFilter) { interfaceIndexes.YFilter = yf }

func (interfaceIndexes *Snmp_InterfaceIndexes) GetGoName(yname string) string {
    if yname == "interface-index" { return "InterfaceIndex" }
    return ""
}

func (interfaceIndexes *Snmp_InterfaceIndexes) GetSegmentPath() string {
    return "interface-indexes"
}

func (interfaceIndexes *Snmp_InterfaceIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-index" {
        for _, c := range interfaceIndexes.InterfaceIndex {
            if interfaceIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_InterfaceIndexes_InterfaceIndex{}
        interfaceIndexes.InterfaceIndex = append(interfaceIndexes.InterfaceIndex, child)
        return &interfaceIndexes.InterfaceIndex[len(interfaceIndexes.InterfaceIndex)-1]
    }
    return nil
}

func (interfaceIndexes *Snmp_InterfaceIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceIndexes.InterfaceIndex {
        children[interfaceIndexes.InterfaceIndex[i].GetSegmentPath()] = &interfaceIndexes.InterfaceIndex[i]
    }
    return children
}

func (interfaceIndexes *Snmp_InterfaceIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceIndexes *Snmp_InterfaceIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceIndexes *Snmp_InterfaceIndexes) GetYangName() string { return "interface-indexes" }

func (interfaceIndexes *Snmp_InterfaceIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceIndexes *Snmp_InterfaceIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceIndexes *Snmp_InterfaceIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceIndexes *Snmp_InterfaceIndexes) SetParent(parent types.Entity) { interfaceIndexes.parent = parent }

func (interfaceIndexes *Snmp_InterfaceIndexes) GetParent() types.Entity { return interfaceIndexes.parent }

func (interfaceIndexes *Snmp_InterfaceIndexes) GetParentYangName() string { return "snmp" }

// Snmp_InterfaceIndexes_InterfaceIndex
// Interface Index
type Snmp_InterfaceIndexes_InterfaceIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Index as used by MIB tables. The type is
    // interface{} with range: -2147483648..2147483647.
    InterfaceIndex interface{}

    // Interface Name. The type is string. This attribute is mandatory.
    InterfaceName interface{}
}

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetFilter() yfilter.YFilter { return interfaceIndex.YFilter }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) SetFilter(yf yfilter.YFilter) { interfaceIndex.YFilter = yf }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetGoName(yname string) string {
    if yname == "interface-index" { return "InterfaceIndex" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetSegmentPath() string {
    return "interface-index" + "[interface-index='" + fmt.Sprintf("%v", interfaceIndex.InterfaceIndex) + "']"
}

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-index"] = interfaceIndex.InterfaceIndex
    leafs["interface-name"] = interfaceIndex.InterfaceName
    return leafs
}

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetYangName() string { return "interface-index" }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) SetParent(parent types.Entity) { interfaceIndex.parent = parent }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetParent() types.Entity { return interfaceIndex.parent }

func (interfaceIndex *Snmp_InterfaceIndexes_InterfaceIndex) GetParentYangName() string { return "interface-indexes" }

// Snmp_IfIndexes
// List of ifnames
type Snmp_IfIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Index. The type is slice of Snmp_IfIndexes_IfIndex.
    IfIndex []Snmp_IfIndexes_IfIndex
}

func (ifIndexes *Snmp_IfIndexes) GetFilter() yfilter.YFilter { return ifIndexes.YFilter }

func (ifIndexes *Snmp_IfIndexes) SetFilter(yf yfilter.YFilter) { ifIndexes.YFilter = yf }

func (ifIndexes *Snmp_IfIndexes) GetGoName(yname string) string {
    if yname == "if-index" { return "IfIndex" }
    return ""
}

func (ifIndexes *Snmp_IfIndexes) GetSegmentPath() string {
    return "if-indexes"
}

func (ifIndexes *Snmp_IfIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "if-index" {
        for _, c := range ifIndexes.IfIndex {
            if ifIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_IfIndexes_IfIndex{}
        ifIndexes.IfIndex = append(ifIndexes.IfIndex, child)
        return &ifIndexes.IfIndex[len(ifIndexes.IfIndex)-1]
    }
    return nil
}

func (ifIndexes *Snmp_IfIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ifIndexes.IfIndex {
        children[ifIndexes.IfIndex[i].GetSegmentPath()] = &ifIndexes.IfIndex[i]
    }
    return children
}

func (ifIndexes *Snmp_IfIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ifIndexes *Snmp_IfIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (ifIndexes *Snmp_IfIndexes) GetYangName() string { return "if-indexes" }

func (ifIndexes *Snmp_IfIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifIndexes *Snmp_IfIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifIndexes *Snmp_IfIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifIndexes *Snmp_IfIndexes) SetParent(parent types.Entity) { ifIndexes.parent = parent }

func (ifIndexes *Snmp_IfIndexes) GetParent() types.Entity { return ifIndexes.parent }

func (ifIndexes *Snmp_IfIndexes) GetParentYangName() string { return "snmp" }

// Snmp_IfIndexes_IfIndex
// Interface Index
type Snmp_IfIndexes_IfIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Index as used by MIB tables. The type is
    // interface{} with range: -2147483648..2147483647.
    InterfaceIndex interface{}

    // Interface Name. The type is string.
    InterfaceName interface{}
}

func (ifIndex *Snmp_IfIndexes_IfIndex) GetFilter() yfilter.YFilter { return ifIndex.YFilter }

func (ifIndex *Snmp_IfIndexes_IfIndex) SetFilter(yf yfilter.YFilter) { ifIndex.YFilter = yf }

func (ifIndex *Snmp_IfIndexes_IfIndex) GetGoName(yname string) string {
    if yname == "interface-index" { return "InterfaceIndex" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (ifIndex *Snmp_IfIndexes_IfIndex) GetSegmentPath() string {
    return "if-index" + "[interface-index='" + fmt.Sprintf("%v", ifIndex.InterfaceIndex) + "']"
}

func (ifIndex *Snmp_IfIndexes_IfIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ifIndex *Snmp_IfIndexes_IfIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ifIndex *Snmp_IfIndexes_IfIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-index"] = ifIndex.InterfaceIndex
    leafs["interface-name"] = ifIndex.InterfaceName
    return leafs
}

func (ifIndex *Snmp_IfIndexes_IfIndex) GetBundleName() string { return "cisco_ios_xr" }

func (ifIndex *Snmp_IfIndexes_IfIndex) GetYangName() string { return "if-index" }

func (ifIndex *Snmp_IfIndexes_IfIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifIndex *Snmp_IfIndexes_IfIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifIndex *Snmp_IfIndexes_IfIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifIndex *Snmp_IfIndexes_IfIndex) SetParent(parent types.Entity) { ifIndex.parent = parent }

func (ifIndex *Snmp_IfIndexes_IfIndex) GetParent() types.Entity { return ifIndex.parent }

func (ifIndex *Snmp_IfIndexes_IfIndex) GetParentYangName() string { return "if-indexes" }

// Snmp_EntityMib
// SNMP entity mib
type Snmp_EntityMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP entity mib.
    EntityPhysicalIndexes Snmp_EntityMib_EntityPhysicalIndexes
}

func (entityMib *Snmp_EntityMib) GetFilter() yfilter.YFilter { return entityMib.YFilter }

func (entityMib *Snmp_EntityMib) SetFilter(yf yfilter.YFilter) { entityMib.YFilter = yf }

func (entityMib *Snmp_EntityMib) GetGoName(yname string) string {
    if yname == "entity-physical-indexes" { return "EntityPhysicalIndexes" }
    return ""
}

func (entityMib *Snmp_EntityMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-entitymib-oper:entity-mib"
}

func (entityMib *Snmp_EntityMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entity-physical-indexes" {
        return &entityMib.EntityPhysicalIndexes
    }
    return nil
}

func (entityMib *Snmp_EntityMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["entity-physical-indexes"] = &entityMib.EntityPhysicalIndexes
    return children
}

func (entityMib *Snmp_EntityMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entityMib *Snmp_EntityMib) GetBundleName() string { return "cisco_ios_xr" }

func (entityMib *Snmp_EntityMib) GetYangName() string { return "entity-mib" }

func (entityMib *Snmp_EntityMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityMib *Snmp_EntityMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityMib *Snmp_EntityMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityMib *Snmp_EntityMib) SetParent(parent types.Entity) { entityMib.parent = parent }

func (entityMib *Snmp_EntityMib) GetParent() types.Entity { return entityMib.parent }

func (entityMib *Snmp_EntityMib) GetParentYangName() string { return "snmp" }

// Snmp_EntityMib_EntityPhysicalIndexes
// SNMP entity mib
type Snmp_EntityMib_EntityPhysicalIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP entPhysical index number. The type is slice of
    // Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex.
    EntityPhysicalIndex []Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex
}

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetFilter() yfilter.YFilter { return entityPhysicalIndexes.YFilter }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) SetFilter(yf yfilter.YFilter) { entityPhysicalIndexes.YFilter = yf }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetGoName(yname string) string {
    if yname == "entity-physical-index" { return "EntityPhysicalIndex" }
    return ""
}

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetSegmentPath() string {
    return "entity-physical-indexes"
}

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "entity-physical-index" {
        for _, c := range entityPhysicalIndexes.EntityPhysicalIndex {
            if entityPhysicalIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex{}
        entityPhysicalIndexes.EntityPhysicalIndex = append(entityPhysicalIndexes.EntityPhysicalIndex, child)
        return &entityPhysicalIndexes.EntityPhysicalIndex[len(entityPhysicalIndexes.EntityPhysicalIndex)-1]
    }
    return nil
}

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entityPhysicalIndexes.EntityPhysicalIndex {
        children[entityPhysicalIndexes.EntityPhysicalIndex[i].GetSegmentPath()] = &entityPhysicalIndexes.EntityPhysicalIndex[i]
    }
    return children
}

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetYangName() string { return "entity-physical-indexes" }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) SetParent(parent types.Entity) { entityPhysicalIndexes.parent = parent }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetParent() types.Entity { return entityPhysicalIndexes.parent }

func (entityPhysicalIndexes *Snmp_EntityMib_EntityPhysicalIndexes) GetParentYangName() string { return "entity-mib" }

// Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex
// SNMP entPhysical index number
type Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetFilter() yfilter.YFilter { return entityPhysicalIndex.YFilter }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) SetFilter(yf yfilter.YFilter) { entityPhysicalIndex.YFilter = yf }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetGoName(yname string) string {
    if yname == "entity-phynum" { return "EntityPhynum" }
    if yname == "physical-index" { return "PhysicalIndex" }
    if yname == "ent-physical-name" { return "EntPhysicalName" }
    if yname == "location" { return "Location" }
    if yname == "ent-physical-descr" { return "EntPhysicalDescr" }
    if yname == "ent-physical-firmware-rev" { return "EntPhysicalFirmwareRev" }
    if yname == "ent-physical-hardware-rev" { return "EntPhysicalHardwareRev" }
    if yname == "ent-physical-modelname" { return "EntPhysicalModelname" }
    if yname == "ent-physical-serial-num" { return "EntPhysicalSerialNum" }
    if yname == "ent-physical-software-rev" { return "EntPhysicalSoftwareRev" }
    if yname == "ent-physical-mfg-name" { return "EntPhysicalMfgName" }
    return ""
}

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetSegmentPath() string {
    return "entity-physical-index" + "[entity-phynum='" + fmt.Sprintf("%v", entityPhysicalIndex.EntityPhynum) + "']"
}

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entity-phynum"] = entityPhysicalIndex.EntityPhynum
    leafs["physical-index"] = entityPhysicalIndex.PhysicalIndex
    leafs["ent-physical-name"] = entityPhysicalIndex.EntPhysicalName
    leafs["location"] = entityPhysicalIndex.Location
    leafs["ent-physical-descr"] = entityPhysicalIndex.EntPhysicalDescr
    leafs["ent-physical-firmware-rev"] = entityPhysicalIndex.EntPhysicalFirmwareRev
    leafs["ent-physical-hardware-rev"] = entityPhysicalIndex.EntPhysicalHardwareRev
    leafs["ent-physical-modelname"] = entityPhysicalIndex.EntPhysicalModelname
    leafs["ent-physical-serial-num"] = entityPhysicalIndex.EntPhysicalSerialNum
    leafs["ent-physical-software-rev"] = entityPhysicalIndex.EntPhysicalSoftwareRev
    leafs["ent-physical-mfg-name"] = entityPhysicalIndex.EntPhysicalMfgName
    return leafs
}

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetBundleName() string { return "cisco_ios_xr" }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetYangName() string { return "entity-physical-index" }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) SetParent(parent types.Entity) { entityPhysicalIndex.parent = parent }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetParent() types.Entity { return entityPhysicalIndex.parent }

func (entityPhysicalIndex *Snmp_EntityMib_EntityPhysicalIndexes_EntityPhysicalIndex) GetParentYangName() string { return "entity-physical-indexes" }

// Snmp_InterfaceMib
// SNMP IF-MIB information
type Snmp_InterfaceMib struct {
    parent types.Entity
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

func (interfaceMib *Snmp_InterfaceMib) GetFilter() yfilter.YFilter { return interfaceMib.YFilter }

func (interfaceMib *Snmp_InterfaceMib) SetFilter(yf yfilter.YFilter) { interfaceMib.YFilter = yf }

func (interfaceMib *Snmp_InterfaceMib) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "interface-connectors" { return "InterfaceConnectors" }
    if yname == "interface-aliases" { return "InterfaceAliases" }
    if yname == "notification-interfaces" { return "NotificationInterfaces" }
    if yname == "interface-stack-statuses" { return "InterfaceStackStatuses" }
    return ""
}

func (interfaceMib *Snmp_InterfaceMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-ifmib-oper:interface-mib"
}

func (interfaceMib *Snmp_InterfaceMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &interfaceMib.Interfaces
    }
    if childYangName == "interface-connectors" {
        return &interfaceMib.InterfaceConnectors
    }
    if childYangName == "interface-aliases" {
        return &interfaceMib.InterfaceAliases
    }
    if childYangName == "notification-interfaces" {
        return &interfaceMib.NotificationInterfaces
    }
    if childYangName == "interface-stack-statuses" {
        return &interfaceMib.InterfaceStackStatuses
    }
    return nil
}

func (interfaceMib *Snmp_InterfaceMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &interfaceMib.Interfaces
    children["interface-connectors"] = &interfaceMib.InterfaceConnectors
    children["interface-aliases"] = &interfaceMib.InterfaceAliases
    children["notification-interfaces"] = &interfaceMib.NotificationInterfaces
    children["interface-stack-statuses"] = &interfaceMib.InterfaceStackStatuses
    return children
}

func (interfaceMib *Snmp_InterfaceMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceMib *Snmp_InterfaceMib) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceMib *Snmp_InterfaceMib) GetYangName() string { return "interface-mib" }

func (interfaceMib *Snmp_InterfaceMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceMib *Snmp_InterfaceMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceMib *Snmp_InterfaceMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceMib *Snmp_InterfaceMib) SetParent(parent types.Entity) { interfaceMib.parent = parent }

func (interfaceMib *Snmp_InterfaceMib) GetParent() types.Entity { return interfaceMib.parent }

func (interfaceMib *Snmp_InterfaceMib) GetParentYangName() string { return "snmp" }

// Snmp_InterfaceMib_Interfaces
// Interfaces ifIndex information
type Snmp_InterfaceMib_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ifIndex for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_Interfaces_Interface.
    Interface []Snmp_InterfaceMib_Interfaces_Interface
}

func (interfaces *Snmp_InterfaceMib_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Snmp_InterfaceMib_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Snmp_InterfaceMib_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Snmp_InterfaceMib_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Snmp_InterfaceMib_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_InterfaceMib_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Snmp_InterfaceMib_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Snmp_InterfaceMib_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Snmp_InterfaceMib_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Snmp_InterfaceMib_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Snmp_InterfaceMib_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Snmp_InterfaceMib_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Snmp_InterfaceMib_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Snmp_InterfaceMib_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Snmp_InterfaceMib_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Snmp_InterfaceMib_Interfaces) GetParentYangName() string { return "interface-mib" }

// Snmp_InterfaceMib_Interfaces_Interface
// ifIndex for a specific Interface Name
type Snmp_InterfaceMib_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface Index. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}
}

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Snmp_InterfaceMib_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "if-index" { return "IfIndex" }
    return ""
}

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["if-index"] = self.IfIndex
    return leafs
}

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Snmp_InterfaceMib_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Snmp_InterfaceMib_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Snmp_InterfaceMib_InterfaceConnectors
// Interfaces ifConnectorPresent information
type Snmp_InterfaceMib_InterfaceConnectors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ifConnectorPresent for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector.
    InterfaceConnector []Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector
}

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetFilter() yfilter.YFilter { return interfaceConnectors.YFilter }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) SetFilter(yf yfilter.YFilter) { interfaceConnectors.YFilter = yf }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetGoName(yname string) string {
    if yname == "interface-connector" { return "InterfaceConnector" }
    return ""
}

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetSegmentPath() string {
    return "interface-connectors"
}

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-connector" {
        for _, c := range interfaceConnectors.InterfaceConnector {
            if interfaceConnectors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector{}
        interfaceConnectors.InterfaceConnector = append(interfaceConnectors.InterfaceConnector, child)
        return &interfaceConnectors.InterfaceConnector[len(interfaceConnectors.InterfaceConnector)-1]
    }
    return nil
}

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceConnectors.InterfaceConnector {
        children[interfaceConnectors.InterfaceConnector[i].GetSegmentPath()] = &interfaceConnectors.InterfaceConnector[i]
    }
    return children
}

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetYangName() string { return "interface-connectors" }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) SetParent(parent types.Entity) { interfaceConnectors.parent = parent }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetParent() types.Entity { return interfaceConnectors.parent }

func (interfaceConnectors *Snmp_InterfaceMib_InterfaceConnectors) GetParentYangName() string { return "interface-mib" }

// Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector
// ifConnectorPresent for a specific Interface
// Name
type Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface ifConnector. The type is string.
    IfConnectorPresent interface{}
}

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetFilter() yfilter.YFilter { return interfaceConnector.YFilter }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) SetFilter(yf yfilter.YFilter) { interfaceConnector.YFilter = yf }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "if-connector-present" { return "IfConnectorPresent" }
    return ""
}

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetSegmentPath() string {
    return "interface-connector" + "[interface-name='" + fmt.Sprintf("%v", interfaceConnector.InterfaceName) + "']"
}

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceConnector.InterfaceName
    leafs["if-connector-present"] = interfaceConnector.IfConnectorPresent
    return leafs
}

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetYangName() string { return "interface-connector" }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) SetParent(parent types.Entity) { interfaceConnector.parent = parent }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetParent() types.Entity { return interfaceConnector.parent }

func (interfaceConnector *Snmp_InterfaceMib_InterfaceConnectors_InterfaceConnector) GetParentYangName() string { return "interface-connectors" }

// Snmp_InterfaceMib_InterfaceAliases
// Interfaces ifAlias information
type Snmp_InterfaceMib_InterfaceAliases struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ifAlias for a specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias.
    InterfaceAlias []Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias
}

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetFilter() yfilter.YFilter { return interfaceAliases.YFilter }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) SetFilter(yf yfilter.YFilter) { interfaceAliases.YFilter = yf }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetGoName(yname string) string {
    if yname == "interface-alias" { return "InterfaceAlias" }
    return ""
}

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetSegmentPath() string {
    return "interface-aliases"
}

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-alias" {
        for _, c := range interfaceAliases.InterfaceAlias {
            if interfaceAliases.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias{}
        interfaceAliases.InterfaceAlias = append(interfaceAliases.InterfaceAlias, child)
        return &interfaceAliases.InterfaceAlias[len(interfaceAliases.InterfaceAlias)-1]
    }
    return nil
}

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceAliases.InterfaceAlias {
        children[interfaceAliases.InterfaceAlias[i].GetSegmentPath()] = &interfaceAliases.InterfaceAlias[i]
    }
    return children
}

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetYangName() string { return "interface-aliases" }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) SetParent(parent types.Entity) { interfaceAliases.parent = parent }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetParent() types.Entity { return interfaceAliases.parent }

func (interfaceAliases *Snmp_InterfaceMib_InterfaceAliases) GetParentYangName() string { return "interface-mib" }

// Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias
// ifAlias for a specific Interface Name
type Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface ifAlias. The type is string.
    IfAlias interface{}
}

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetFilter() yfilter.YFilter { return interfaceAlias.YFilter }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) SetFilter(yf yfilter.YFilter) { interfaceAlias.YFilter = yf }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "if-alias" { return "IfAlias" }
    return ""
}

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetSegmentPath() string {
    return "interface-alias" + "[interface-name='" + fmt.Sprintf("%v", interfaceAlias.InterfaceName) + "']"
}

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceAlias.InterfaceName
    leafs["if-alias"] = interfaceAlias.IfAlias
    return leafs
}

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetYangName() string { return "interface-alias" }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) SetParent(parent types.Entity) { interfaceAlias.parent = parent }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetParent() types.Entity { return interfaceAlias.parent }

func (interfaceAlias *Snmp_InterfaceMib_InterfaceAliases_InterfaceAlias) GetParentYangName() string { return "interface-aliases" }

// Snmp_InterfaceMib_NotificationInterfaces
// Interfaces Notification information
type Snmp_InterfaceMib_NotificationInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Notification for specific Interface Name. The type is slice of
    // Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface.
    NotificationInterface []Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface
}

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetFilter() yfilter.YFilter { return notificationInterfaces.YFilter }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) SetFilter(yf yfilter.YFilter) { notificationInterfaces.YFilter = yf }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetGoName(yname string) string {
    if yname == "notification-interface" { return "NotificationInterface" }
    return ""
}

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetSegmentPath() string {
    return "notification-interfaces"
}

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "notification-interface" {
        for _, c := range notificationInterfaces.NotificationInterface {
            if notificationInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface{}
        notificationInterfaces.NotificationInterface = append(notificationInterfaces.NotificationInterface, child)
        return &notificationInterfaces.NotificationInterface[len(notificationInterfaces.NotificationInterface)-1]
    }
    return nil
}

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range notificationInterfaces.NotificationInterface {
        children[notificationInterfaces.NotificationInterface[i].GetSegmentPath()] = &notificationInterfaces.NotificationInterface[i]
    }
    return children
}

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetYangName() string { return "notification-interfaces" }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) SetParent(parent types.Entity) { notificationInterfaces.parent = parent }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetParent() types.Entity { return notificationInterfaces.parent }

func (notificationInterfaces *Snmp_InterfaceMib_NotificationInterfaces) GetParentYangName() string { return "interface-mib" }

// Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface
// Notification for specific Interface Name
type Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // LinkUpDown notification status. The type is LinkUpDownStatus.
    LinkUpDownNotifStatus interface{}
}

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetFilter() yfilter.YFilter { return notificationInterface.YFilter }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) SetFilter(yf yfilter.YFilter) { notificationInterface.YFilter = yf }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "link-up-down-notif-status" { return "LinkUpDownNotifStatus" }
    return ""
}

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetSegmentPath() string {
    return "notification-interface" + "[interface-name='" + fmt.Sprintf("%v", notificationInterface.InterfaceName) + "']"
}

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = notificationInterface.InterfaceName
    leafs["link-up-down-notif-status"] = notificationInterface.LinkUpDownNotifStatus
    return leafs
}

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetBundleName() string { return "cisco_ios_xr" }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetYangName() string { return "notification-interface" }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) SetParent(parent types.Entity) { notificationInterface.parent = parent }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetParent() types.Entity { return notificationInterface.parent }

func (notificationInterface *Snmp_InterfaceMib_NotificationInterfaces_NotificationInterface) GetParentYangName() string { return "notification-interfaces" }

// Snmp_InterfaceMib_InterfaceStackStatuses
// Interfaces ifstackstatus information
type Snmp_InterfaceMib_InterfaceStackStatuses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ifstatus for a pair of Interface. The type is slice of
    // Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus.
    InterfaceStackStatus []Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus
}

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetFilter() yfilter.YFilter { return interfaceStackStatuses.YFilter }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) SetFilter(yf yfilter.YFilter) { interfaceStackStatuses.YFilter = yf }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetGoName(yname string) string {
    if yname == "interface-stack-status" { return "InterfaceStackStatus" }
    return ""
}

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetSegmentPath() string {
    return "interface-stack-statuses"
}

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-stack-status" {
        for _, c := range interfaceStackStatuses.InterfaceStackStatus {
            if interfaceStackStatuses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus{}
        interfaceStackStatuses.InterfaceStackStatus = append(interfaceStackStatuses.InterfaceStackStatus, child)
        return &interfaceStackStatuses.InterfaceStackStatus[len(interfaceStackStatuses.InterfaceStackStatus)-1]
    }
    return nil
}

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceStackStatuses.InterfaceStackStatus {
        children[interfaceStackStatuses.InterfaceStackStatus[i].GetSegmentPath()] = &interfaceStackStatuses.InterfaceStackStatus[i]
    }
    return children
}

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetYangName() string { return "interface-stack-statuses" }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) SetParent(parent types.Entity) { interfaceStackStatuses.parent = parent }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetParent() types.Entity { return interfaceStackStatuses.parent }

func (interfaceStackStatuses *Snmp_InterfaceMib_InterfaceStackStatuses) GetParentYangName() string { return "interface-mib" }

// Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus
// ifstatus for a pair of Interface
type Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetFilter() yfilter.YFilter { return interfaceStackStatus.YFilter }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) SetFilter(yf yfilter.YFilter) { interfaceStackStatus.YFilter = yf }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetGoName(yname string) string {
    if yname == "interface-stack-status" { return "InterfaceStackStatus" }
    if yname == "if-stack-higher-layer" { return "IfStackHigherLayer" }
    if yname == "if-stack-lower-layer" { return "IfStackLowerLayer" }
    if yname == "if-stack-status" { return "IfStackStatus" }
    return ""
}

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetSegmentPath() string {
    return "interface-stack-status" + "[interface-stack-status='" + fmt.Sprintf("%v", interfaceStackStatus.InterfaceStackStatus) + "']"
}

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-stack-status"] = interfaceStackStatus.InterfaceStackStatus
    leafs["if-stack-higher-layer"] = interfaceStackStatus.IfStackHigherLayer
    leafs["if-stack-lower-layer"] = interfaceStackStatus.IfStackLowerLayer
    leafs["if-stack-status"] = interfaceStackStatus.IfStackStatus
    return leafs
}

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetYangName() string { return "interface-stack-status" }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) SetParent(parent types.Entity) { interfaceStackStatus.parent = parent }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetParent() types.Entity { return interfaceStackStatus.parent }

func (interfaceStackStatus *Snmp_InterfaceMib_InterfaceStackStatuses_InterfaceStackStatus) GetParentYangName() string { return "interface-stack-statuses" }

// Snmp_SensorMib
// SNMP sensor MIB information
type Snmp_SensorMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of physical index table for threshold value.
    PhysicalIndexes Snmp_SensorMib_PhysicalIndexes

    // List of physical index .
    EntPhyIndexes Snmp_SensorMib_EntPhyIndexes
}

func (sensorMib *Snmp_SensorMib) GetFilter() yfilter.YFilter { return sensorMib.YFilter }

func (sensorMib *Snmp_SensorMib) SetFilter(yf yfilter.YFilter) { sensorMib.YFilter = yf }

func (sensorMib *Snmp_SensorMib) GetGoName(yname string) string {
    if yname == "physical-indexes" { return "PhysicalIndexes" }
    if yname == "ent-phy-indexes" { return "EntPhyIndexes" }
    return ""
}

func (sensorMib *Snmp_SensorMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-sensormib-oper:sensor-mib"
}

func (sensorMib *Snmp_SensorMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "physical-indexes" {
        return &sensorMib.PhysicalIndexes
    }
    if childYangName == "ent-phy-indexes" {
        return &sensorMib.EntPhyIndexes
    }
    return nil
}

func (sensorMib *Snmp_SensorMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["physical-indexes"] = &sensorMib.PhysicalIndexes
    children["ent-phy-indexes"] = &sensorMib.EntPhyIndexes
    return children
}

func (sensorMib *Snmp_SensorMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sensorMib *Snmp_SensorMib) GetBundleName() string { return "cisco_ios_xr" }

func (sensorMib *Snmp_SensorMib) GetYangName() string { return "sensor-mib" }

func (sensorMib *Snmp_SensorMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensorMib *Snmp_SensorMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensorMib *Snmp_SensorMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensorMib *Snmp_SensorMib) SetParent(parent types.Entity) { sensorMib.parent = parent }

func (sensorMib *Snmp_SensorMib) GetParent() types.Entity { return sensorMib.parent }

func (sensorMib *Snmp_SensorMib) GetParentYangName() string { return "snmp" }

// Snmp_SensorMib_PhysicalIndexes
// List of physical index table for threshold
// value
type Snmp_SensorMib_PhysicalIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold value for physical index. The type is slice of
    // Snmp_SensorMib_PhysicalIndexes_PhysicalIndex.
    PhysicalIndex []Snmp_SensorMib_PhysicalIndexes_PhysicalIndex
}

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetFilter() yfilter.YFilter { return physicalIndexes.YFilter }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) SetFilter(yf yfilter.YFilter) { physicalIndexes.YFilter = yf }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetGoName(yname string) string {
    if yname == "physical-index" { return "PhysicalIndex" }
    return ""
}

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetSegmentPath() string {
    return "physical-indexes"
}

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "physical-index" {
        for _, c := range physicalIndexes.PhysicalIndex {
            if physicalIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_SensorMib_PhysicalIndexes_PhysicalIndex{}
        physicalIndexes.PhysicalIndex = append(physicalIndexes.PhysicalIndex, child)
        return &physicalIndexes.PhysicalIndex[len(physicalIndexes.PhysicalIndex)-1]
    }
    return nil
}

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range physicalIndexes.PhysicalIndex {
        children[physicalIndexes.PhysicalIndex[i].GetSegmentPath()] = &physicalIndexes.PhysicalIndex[i]
    }
    return children
}

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetYangName() string { return "physical-indexes" }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) SetParent(parent types.Entity) { physicalIndexes.parent = parent }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetParent() types.Entity { return physicalIndexes.parent }

func (physicalIndexes *Snmp_SensorMib_PhysicalIndexes) GetParentYangName() string { return "sensor-mib" }

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex
// Threshold value for physical index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Physical index. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Index interface{}

    // List of threshold index.
    ThresholdIndexes Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes
}

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetFilter() yfilter.YFilter { return physicalIndex.YFilter }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) SetFilter(yf yfilter.YFilter) { physicalIndex.YFilter = yf }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "threshold-indexes" { return "ThresholdIndexes" }
    return ""
}

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetSegmentPath() string {
    return "physical-index" + "[index='" + fmt.Sprintf("%v", physicalIndex.Index) + "']"
}

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-indexes" {
        return &physicalIndex.ThresholdIndexes
    }
    return nil
}

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold-indexes"] = &physicalIndex.ThresholdIndexes
    return children
}

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = physicalIndex.Index
    return leafs
}

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetBundleName() string { return "cisco_ios_xr" }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetYangName() string { return "physical-index" }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) SetParent(parent types.Entity) { physicalIndex.parent = parent }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetParent() types.Entity { return physicalIndex.parent }

func (physicalIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex) GetParentYangName() string { return "physical-indexes" }

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes
// List of threshold index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Threshold value for threshold index. The type is slice of
    // Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex.
    ThresholdIndex []Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex
}

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetFilter() yfilter.YFilter { return thresholdIndexes.YFilter }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) SetFilter(yf yfilter.YFilter) { thresholdIndexes.YFilter = yf }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetGoName(yname string) string {
    if yname == "threshold-index" { return "ThresholdIndex" }
    return ""
}

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetSegmentPath() string {
    return "threshold-indexes"
}

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold-index" {
        for _, c := range thresholdIndexes.ThresholdIndex {
            if thresholdIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex{}
        thresholdIndexes.ThresholdIndex = append(thresholdIndexes.ThresholdIndex, child)
        return &thresholdIndexes.ThresholdIndex[len(thresholdIndexes.ThresholdIndex)-1]
    }
    return nil
}

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range thresholdIndexes.ThresholdIndex {
        children[thresholdIndexes.ThresholdIndex[i].GetSegmentPath()] = &thresholdIndexes.ThresholdIndex[i]
    }
    return children
}

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetYangName() string { return "threshold-indexes" }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) SetParent(parent types.Entity) { thresholdIndexes.parent = parent }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetParent() types.Entity { return thresholdIndexes.parent }

func (thresholdIndexes *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes) GetParentYangName() string { return "physical-index" }

// Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex
// Threshold value for threshold index
type Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetFilter() yfilter.YFilter { return thresholdIndex.YFilter }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) SetFilter(yf yfilter.YFilter) { thresholdIndex.YFilter = yf }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetGoName(yname string) string {
    if yname == "phy-index" { return "PhyIndex" }
    if yname == "thre-index" { return "ThreIndex" }
    if yname == "threshold-severity" { return "ThresholdSeverity" }
    if yname == "threshold-relation" { return "ThresholdRelation" }
    if yname == "threshold-value" { return "ThresholdValue" }
    if yname == "threshold-evaluation" { return "ThresholdEvaluation" }
    if yname == "threshold-notification-enabled" { return "ThresholdNotificationEnabled" }
    return ""
}

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetSegmentPath() string {
    return "threshold-index"
}

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["phy-index"] = thresholdIndex.PhyIndex
    leafs["thre-index"] = thresholdIndex.ThreIndex
    leafs["threshold-severity"] = thresholdIndex.ThresholdSeverity
    leafs["threshold-relation"] = thresholdIndex.ThresholdRelation
    leafs["threshold-value"] = thresholdIndex.ThresholdValue
    leafs["threshold-evaluation"] = thresholdIndex.ThresholdEvaluation
    leafs["threshold-notification-enabled"] = thresholdIndex.ThresholdNotificationEnabled
    return leafs
}

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetBundleName() string { return "cisco_ios_xr" }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetYangName() string { return "threshold-index" }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) SetParent(parent types.Entity) { thresholdIndex.parent = parent }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetParent() types.Entity { return thresholdIndex.parent }

func (thresholdIndex *Snmp_SensorMib_PhysicalIndexes_PhysicalIndex_ThresholdIndexes_ThresholdIndex) GetParentYangName() string { return "threshold-indexes" }

// Snmp_SensorMib_EntPhyIndexes
// List of physical index 
type Snmp_SensorMib_EntPhyIndexes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sensor value for physical index. The type is slice of
    // Snmp_SensorMib_EntPhyIndexes_EntPhyIndex.
    EntPhyIndex []Snmp_SensorMib_EntPhyIndexes_EntPhyIndex
}

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetFilter() yfilter.YFilter { return entPhyIndexes.YFilter }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) SetFilter(yf yfilter.YFilter) { entPhyIndexes.YFilter = yf }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetGoName(yname string) string {
    if yname == "ent-phy-index" { return "EntPhyIndex" }
    return ""
}

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetSegmentPath() string {
    return "ent-phy-indexes"
}

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ent-phy-index" {
        for _, c := range entPhyIndexes.EntPhyIndex {
            if entPhyIndexes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_SensorMib_EntPhyIndexes_EntPhyIndex{}
        entPhyIndexes.EntPhyIndex = append(entPhyIndexes.EntPhyIndex, child)
        return &entPhyIndexes.EntPhyIndex[len(entPhyIndexes.EntPhyIndex)-1]
    }
    return nil
}

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range entPhyIndexes.EntPhyIndex {
        children[entPhyIndexes.EntPhyIndex[i].GetSegmentPath()] = &entPhyIndexes.EntPhyIndex[i]
    }
    return children
}

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetBundleName() string { return "cisco_ios_xr" }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetYangName() string { return "ent-phy-indexes" }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) SetParent(parent types.Entity) { entPhyIndexes.parent = parent }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetParent() types.Entity { return entPhyIndexes.parent }

func (entPhyIndexes *Snmp_SensorMib_EntPhyIndexes) GetParentYangName() string { return "sensor-mib" }

// Snmp_SensorMib_EntPhyIndexes_EntPhyIndex
// Sensor value for physical index
type Snmp_SensorMib_EntPhyIndexes_EntPhyIndex struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetFilter() yfilter.YFilter { return entPhyIndex.YFilter }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) SetFilter(yf yfilter.YFilter) { entPhyIndex.YFilter = yf }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "field-validity-bitmap" { return "FieldValidityBitmap" }
    if yname == "device-description" { return "DeviceDescription" }
    if yname == "units" { return "Units" }
    if yname == "device-id" { return "DeviceId" }
    if yname == "value" { return "Value" }
    if yname == "alarm-type" { return "AlarmType" }
    if yname == "data-type" { return "DataType" }
    if yname == "scale" { return "Scale" }
    if yname == "precision" { return "Precision" }
    if yname == "status" { return "Status" }
    if yname == "age-time-stamp" { return "AgeTimeStamp" }
    if yname == "update-rate" { return "UpdateRate" }
    if yname == "measured-entity" { return "MeasuredEntity" }
    return ""
}

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetSegmentPath() string {
    return "ent-phy-index" + "[index='" + fmt.Sprintf("%v", entPhyIndex.Index) + "']"
}

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = entPhyIndex.Index
    leafs["field-validity-bitmap"] = entPhyIndex.FieldValidityBitmap
    leafs["device-description"] = entPhyIndex.DeviceDescription
    leafs["units"] = entPhyIndex.Units
    leafs["device-id"] = entPhyIndex.DeviceId
    leafs["value"] = entPhyIndex.Value
    leafs["alarm-type"] = entPhyIndex.AlarmType
    leafs["data-type"] = entPhyIndex.DataType
    leafs["scale"] = entPhyIndex.Scale
    leafs["precision"] = entPhyIndex.Precision
    leafs["status"] = entPhyIndex.Status
    leafs["age-time-stamp"] = entPhyIndex.AgeTimeStamp
    leafs["update-rate"] = entPhyIndex.UpdateRate
    leafs["measured-entity"] = entPhyIndex.MeasuredEntity
    return leafs
}

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetBundleName() string { return "cisco_ios_xr" }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetYangName() string { return "ent-phy-index" }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) SetParent(parent types.Entity) { entPhyIndex.parent = parent }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetParent() types.Entity { return entPhyIndex.parent }

func (entPhyIndex *Snmp_SensorMib_EntPhyIndexes_EntPhyIndex) GetParentYangName() string { return "ent-phy-indexes" }

