// Copyright (c) 2016 by Cisco Systems, Inc.All rights reserved.
package cisco_self_mgmt

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_self_mgmt"))
    ydk.RegisterEntity("{http://cisco.com/yang/cisco-self-mgmt netconf-yang}", reflect.TypeOf(NetconfYang{}))
    ydk.RegisterEntity("cisco-self-mgmt:netconf-yang", reflect.TypeOf(NetconfYang{}))
}

// NetconfYang
// Top level container shared by cisco-ia and cisco-odm models
type NetconfYang struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Customize the behavior of the DMI applications.
    CiscoIa NetconfYang_CiscoIa

    
    CiscoOdm NetconfYang_CiscoOdm
}

func (netconfYang *NetconfYang) GetFilter() yfilter.YFilter { return netconfYang.YFilter }

func (netconfYang *NetconfYang) SetFilter(yf yfilter.YFilter) { netconfYang.YFilter = yf }

func (netconfYang *NetconfYang) GetGoName(yname string) string {
    if yname == "cisco-ia:cisco-ia" { return "CiscoIa" }
    if yname == "cisco-odm:cisco-odm" { return "CiscoOdm" }
    return ""
}

func (netconfYang *NetconfYang) GetSegmentPath() string {
    return "cisco-self-mgmt:netconf-yang"
}

func (netconfYang *NetconfYang) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cisco-ia:cisco-ia" {
        return &netconfYang.CiscoIa
    }
    if childYangName == "cisco-odm:cisco-odm" {
        return &netconfYang.CiscoOdm
    }
    return nil
}

func (netconfYang *NetconfYang) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cisco-ia:cisco-ia"] = &netconfYang.CiscoIa
    children["cisco-odm:cisco-odm"] = &netconfYang.CiscoOdm
    return children
}

func (netconfYang *NetconfYang) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconfYang *NetconfYang) GetBundleName() string { return "cisco_ios_xe" }

func (netconfYang *NetconfYang) GetYangName() string { return "netconf-yang" }

func (netconfYang *NetconfYang) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (netconfYang *NetconfYang) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (netconfYang *NetconfYang) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (netconfYang *NetconfYang) SetParent(parent types.Entity) { netconfYang.parent = parent }

func (netconfYang *NetconfYang) GetParent() types.Entity { return netconfYang.parent }

func (netconfYang *NetconfYang) GetParentYangName() string { return "cisco-self-mgmt" }

// NetconfYang_CiscoIa
// Customize the behavior of the DMI applications
type NetconfYang_CiscoIa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enables automatic synchronization of the network element's running
    // configuration with the DMI database. The type is CiaSyncType. The default
    // value is without-defaults.
    AutoSync interface{}

    // Enables synchronization of the network element's running configuration with
    // the DMI database when DMI initializes. The type is CiaSyncType. The default
    // value is without-defaults.
    InitSync interface{}

    // When enabled, intelligent-sync monitors all  ttys for configuration changes
    // and replays  these changes to the DMI database once the tty exits
    // configuration mode.  When  disabled, the complete running-configuration is
    // used to synchronize the DMI database whenever a CLI configuration change is
    // detected. The type is bool. The default value is true.
    IntelligentSync interface{}

    // 0 = Disabled,  1 = Save input message, DMI debugs, and response,  2 = Level
    // 1 + Save "before" and "after"      running-config, 3 = Level 2 + rollback
    // file and configuration diff. The type is interface{} with range: 0..3. The
    // default value is 0.
    MessageDiagLevel interface{}

    // The maximum number of messages whose diagnostic data  in kept in persistent
    // storage. The type is interface{} with range: 0..199. The default value is
    // 30.
    MaxDiagMessagesSaved interface{}

    // Run "show ip access-list" and send to ConfD. The type is bool. The default
    // value is true.
    PostSyncAclProcess interface{}

    // Delay in ms before applying CDB change to NE. The type is interface{} with
    // range: -32768..32767. The default value is 0.
    ConfigChangeDelay interface{}

    // Process any parser output from configuration changes as a possible error.
    // The type is bool. The default value is true.
    ProcessMissingPrc interface{}

    // The community string for communication with the SNMP         agent. The
    // type is string. The default value is private.
    SnmpCommunityString interface{}

    // Preserve specified model paths in the NED model when performing a sync from
    // the  network element. The type is bool. The default value is false.
    PreservePathsEnabled interface{}

    // TTY number used by NetworkElementSynchronizer. The type is interface{} with
    // range: -32768..32767.
    NesTtynum interface{}

    // Indicates if CDB restored from NES running-config. The type is bool. The
    // default value is false.
    Restored interface{}

    // This container describes the configuration parameters for SNMP Trap to
    // NetConf notification processing.
    SnmpTrapControl NetconfYang_CiscoIa_SnmpTrapControl

    // Model paths from the NED model to preserve upon a sync from the network
    // element. These paths are not cleared from the  running data store prior to
    // the sync. These are expressed as nodes separated by  slash '/', e.g. 
    // /native/ip/access-list. The type is slice of
    // NetconfYang_CiscoIa_PreserveNedPath.
    PreserveNedPath []NetconfYang_CiscoIa_PreserveNedPath

    // Parser output from configuration  change that is informational only, not an
    // error. This is a read only list containing known informational  messages.
    // The type is slice of NetconfYang_CiscoIa_ParserMsgIgnore.
    ParserMsgIgnore []NetconfYang_CiscoIa_ParserMsgIgnore

    // Parser output from configuration  change that is informational only, not an
    // error. The type is slice of NetconfYang_CiscoIa_ConfParserMsgIgnore.
    ConfParserMsgIgnore []NetconfYang_CiscoIa_ConfParserMsgIgnore

    // IOS commands that result in other automatic configurations being applied
    // for which a complete sync is required. The type is slice of
    // NetconfYang_CiscoIa_FullSyncCli.
    FullSyncCli []NetconfYang_CiscoIa_FullSyncCli

    // A user-configurable list of IOS commands  that result in other automatic
    // configurations  being applied for which a complete sync  is required. The
    // type is slice of NetconfYang_CiscoIa_ConfFullSyncCli.
    ConfFullSyncCli []NetconfYang_CiscoIa_ConfFullSyncCli

    // Controls logging behavior of DMI applications.
    Logging NetconfYang_CiscoIa_Logging

    // Controls blocking of command lines, either  from the NE to Confd or
    // disallowing manual input from the console/vty.
    Blocking NetconfYang_CiscoIa_Blocking
}

func (ciscoIa *NetconfYang_CiscoIa) GetFilter() yfilter.YFilter { return ciscoIa.YFilter }

func (ciscoIa *NetconfYang_CiscoIa) SetFilter(yf yfilter.YFilter) { ciscoIa.YFilter = yf }

func (ciscoIa *NetconfYang_CiscoIa) GetGoName(yname string) string {
    if yname == "auto-sync" { return "AutoSync" }
    if yname == "init-sync" { return "InitSync" }
    if yname == "intelligent-sync" { return "IntelligentSync" }
    if yname == "message-diag-level" { return "MessageDiagLevel" }
    if yname == "max-diag-messages-saved" { return "MaxDiagMessagesSaved" }
    if yname == "post-sync-acl-process" { return "PostSyncAclProcess" }
    if yname == "config-change-delay" { return "ConfigChangeDelay" }
    if yname == "process-missing-prc" { return "ProcessMissingPrc" }
    if yname == "snmp-community-string" { return "SnmpCommunityString" }
    if yname == "preserve-paths-enabled" { return "PreservePathsEnabled" }
    if yname == "nes-ttynum" { return "NesTtynum" }
    if yname == "restored" { return "Restored" }
    if yname == "snmp-trap-control" { return "SnmpTrapControl" }
    if yname == "preserve-ned-path" { return "PreserveNedPath" }
    if yname == "parser-msg-ignore" { return "ParserMsgIgnore" }
    if yname == "conf-parser-msg-ignore" { return "ConfParserMsgIgnore" }
    if yname == "full-sync-cli" { return "FullSyncCli" }
    if yname == "conf-full-sync-cli" { return "ConfFullSyncCli" }
    if yname == "logging" { return "Logging" }
    if yname == "blocking" { return "Blocking" }
    return ""
}

func (ciscoIa *NetconfYang_CiscoIa) GetSegmentPath() string {
    return "cisco-ia:cisco-ia"
}

func (ciscoIa *NetconfYang_CiscoIa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snmp-trap-control" {
        return &ciscoIa.SnmpTrapControl
    }
    if childYangName == "preserve-ned-path" {
        for _, c := range ciscoIa.PreserveNedPath {
            if ciscoIa.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoIa_PreserveNedPath{}
        ciscoIa.PreserveNedPath = append(ciscoIa.PreserveNedPath, child)
        return &ciscoIa.PreserveNedPath[len(ciscoIa.PreserveNedPath)-1]
    }
    if childYangName == "parser-msg-ignore" {
        for _, c := range ciscoIa.ParserMsgIgnore {
            if ciscoIa.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoIa_ParserMsgIgnore{}
        ciscoIa.ParserMsgIgnore = append(ciscoIa.ParserMsgIgnore, child)
        return &ciscoIa.ParserMsgIgnore[len(ciscoIa.ParserMsgIgnore)-1]
    }
    if childYangName == "conf-parser-msg-ignore" {
        for _, c := range ciscoIa.ConfParserMsgIgnore {
            if ciscoIa.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoIa_ConfParserMsgIgnore{}
        ciscoIa.ConfParserMsgIgnore = append(ciscoIa.ConfParserMsgIgnore, child)
        return &ciscoIa.ConfParserMsgIgnore[len(ciscoIa.ConfParserMsgIgnore)-1]
    }
    if childYangName == "full-sync-cli" {
        for _, c := range ciscoIa.FullSyncCli {
            if ciscoIa.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoIa_FullSyncCli{}
        ciscoIa.FullSyncCli = append(ciscoIa.FullSyncCli, child)
        return &ciscoIa.FullSyncCli[len(ciscoIa.FullSyncCli)-1]
    }
    if childYangName == "conf-full-sync-cli" {
        for _, c := range ciscoIa.ConfFullSyncCli {
            if ciscoIa.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoIa_ConfFullSyncCli{}
        ciscoIa.ConfFullSyncCli = append(ciscoIa.ConfFullSyncCli, child)
        return &ciscoIa.ConfFullSyncCli[len(ciscoIa.ConfFullSyncCli)-1]
    }
    if childYangName == "logging" {
        return &ciscoIa.Logging
    }
    if childYangName == "blocking" {
        return &ciscoIa.Blocking
    }
    return nil
}

func (ciscoIa *NetconfYang_CiscoIa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snmp-trap-control"] = &ciscoIa.SnmpTrapControl
    for i := range ciscoIa.PreserveNedPath {
        children[ciscoIa.PreserveNedPath[i].GetSegmentPath()] = &ciscoIa.PreserveNedPath[i]
    }
    for i := range ciscoIa.ParserMsgIgnore {
        children[ciscoIa.ParserMsgIgnore[i].GetSegmentPath()] = &ciscoIa.ParserMsgIgnore[i]
    }
    for i := range ciscoIa.ConfParserMsgIgnore {
        children[ciscoIa.ConfParserMsgIgnore[i].GetSegmentPath()] = &ciscoIa.ConfParserMsgIgnore[i]
    }
    for i := range ciscoIa.FullSyncCli {
        children[ciscoIa.FullSyncCli[i].GetSegmentPath()] = &ciscoIa.FullSyncCli[i]
    }
    for i := range ciscoIa.ConfFullSyncCli {
        children[ciscoIa.ConfFullSyncCli[i].GetSegmentPath()] = &ciscoIa.ConfFullSyncCli[i]
    }
    children["logging"] = &ciscoIa.Logging
    children["blocking"] = &ciscoIa.Blocking
    return children
}

func (ciscoIa *NetconfYang_CiscoIa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto-sync"] = ciscoIa.AutoSync
    leafs["init-sync"] = ciscoIa.InitSync
    leafs["intelligent-sync"] = ciscoIa.IntelligentSync
    leafs["message-diag-level"] = ciscoIa.MessageDiagLevel
    leafs["max-diag-messages-saved"] = ciscoIa.MaxDiagMessagesSaved
    leafs["post-sync-acl-process"] = ciscoIa.PostSyncAclProcess
    leafs["config-change-delay"] = ciscoIa.ConfigChangeDelay
    leafs["process-missing-prc"] = ciscoIa.ProcessMissingPrc
    leafs["snmp-community-string"] = ciscoIa.SnmpCommunityString
    leafs["preserve-paths-enabled"] = ciscoIa.PreservePathsEnabled
    leafs["nes-ttynum"] = ciscoIa.NesTtynum
    leafs["restored"] = ciscoIa.Restored
    return leafs
}

func (ciscoIa *NetconfYang_CiscoIa) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoIa *NetconfYang_CiscoIa) GetYangName() string { return "cisco-ia" }

func (ciscoIa *NetconfYang_CiscoIa) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoIa *NetconfYang_CiscoIa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoIa *NetconfYang_CiscoIa) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoIa *NetconfYang_CiscoIa) SetParent(parent types.Entity) { ciscoIa.parent = parent }

func (ciscoIa *NetconfYang_CiscoIa) GetParent() types.Entity { return ciscoIa.parent }

func (ciscoIa *NetconfYang_CiscoIa) GetParentYangName() string { return "netconf-yang" }

// NetconfYang_CiscoIa_SnmpTrapControl
// This container describes the configuration parameters
// for SNMP Trap to NetConf notification processing.
type NetconfYang_CiscoIa_SnmpTrapControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf enables or disables forwarding for all SNMP traps. The type is
    // bool. The default value is true.
    GlobalForwarding interface{}

    // This list describes SNMP Traps that are  supported for automatic
    // translation to NetConf notifications. The type is slice of
    // NetconfYang_CiscoIa_SnmpTrapControl_TrapList.
    TrapList []NetconfYang_CiscoIa_SnmpTrapControl_TrapList
}

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetFilter() yfilter.YFilter { return snmpTrapControl.YFilter }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) SetFilter(yf yfilter.YFilter) { snmpTrapControl.YFilter = yf }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetGoName(yname string) string {
    if yname == "global-forwarding" { return "GlobalForwarding" }
    if yname == "trap-list" { return "TrapList" }
    return ""
}

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetSegmentPath() string {
    return "snmp-trap-control"
}

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-list" {
        for _, c := range snmpTrapControl.TrapList {
            if snmpTrapControl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoIa_SnmpTrapControl_TrapList{}
        snmpTrapControl.TrapList = append(snmpTrapControl.TrapList, child)
        return &snmpTrapControl.TrapList[len(snmpTrapControl.TrapList)-1]
    }
    return nil
}

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range snmpTrapControl.TrapList {
        children[snmpTrapControl.TrapList[i].GetSegmentPath()] = &snmpTrapControl.TrapList[i]
    }
    return children
}

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global-forwarding"] = snmpTrapControl.GlobalForwarding
    return leafs
}

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetBundleName() string { return "cisco_ios_xe" }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetYangName() string { return "snmp-trap-control" }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) SetParent(parent types.Entity) { snmpTrapControl.parent = parent }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetParent() types.Entity { return snmpTrapControl.parent }

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetParentYangName() string { return "cisco-ia" }

// NetconfYang_CiscoIa_SnmpTrapControl_TrapList
// This list describes SNMP Traps that are 
// supported for automatic translation to NetConf
// notifications.
type NetconfYang_CiscoIa_SnmpTrapControl_TrapList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf contains the OID for the  SNMP trap to
    // be forwarded. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    TrapOid interface{}

    // This leaf contains the name of the SNMP trap to be  forwarded. The type is
    // string.
    Description interface{}

    // This leaf enables or disables forwarding for this SNMP trap. The type is
    // bool. The default value is true.
    Forward interface{}
}

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetFilter() yfilter.YFilter { return trapList.YFilter }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) SetFilter(yf yfilter.YFilter) { trapList.YFilter = yf }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetGoName(yname string) string {
    if yname == "trap-oid" { return "TrapOid" }
    if yname == "description" { return "Description" }
    if yname == "forward" { return "Forward" }
    return ""
}

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetSegmentPath() string {
    return "trap-list" + "[trap-oid='" + fmt.Sprintf("%v", trapList.TrapOid) + "']"
}

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trap-oid"] = trapList.TrapOid
    leafs["description"] = trapList.Description
    leafs["forward"] = trapList.Forward
    return leafs
}

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetBundleName() string { return "cisco_ios_xe" }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetYangName() string { return "trap-list" }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) SetParent(parent types.Entity) { trapList.parent = parent }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetParent() types.Entity { return trapList.parent }

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetParentYangName() string { return "snmp-trap-control" }

// NetconfYang_CiscoIa_PreserveNedPath
// Model paths from the NED model to preserve
// upon a sync from the network element.
// These paths are not cleared from the 
// running data store prior to the sync.
// These are expressed as nodes separated by 
// slash '/', e.g.  /native/ip/access-list
type NetconfYang_CiscoIa_PreserveNedPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An XPath from the NED model. The type is string
    // with length: 1..1024.
    Xpath interface{}
}

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetFilter() yfilter.YFilter { return preserveNedPath.YFilter }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) SetFilter(yf yfilter.YFilter) { preserveNedPath.YFilter = yf }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetGoName(yname string) string {
    if yname == "xpath" { return "Xpath" }
    return ""
}

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetSegmentPath() string {
    return "preserve-ned-path" + "[xpath='" + fmt.Sprintf("%v", preserveNedPath.Xpath) + "']"
}

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["xpath"] = preserveNedPath.Xpath
    return leafs
}

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetBundleName() string { return "cisco_ios_xe" }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetYangName() string { return "preserve-ned-path" }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) SetParent(parent types.Entity) { preserveNedPath.parent = parent }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetParent() types.Entity { return preserveNedPath.parent }

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetParentYangName() string { return "cisco-ia" }

// NetconfYang_CiscoIa_ParserMsgIgnore
// Parser output from configuration 
// change that is informational only,
// not an error. This is a read only
// list containing known informational 
// messages
type NetconfYang_CiscoIa_ParserMsgIgnore struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression to match parser output to be
    // ignored. The type is string with length: 1..255.
    Message interface{}
}

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetFilter() yfilter.YFilter { return parserMsgIgnore.YFilter }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) SetFilter(yf yfilter.YFilter) { parserMsgIgnore.YFilter = yf }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetGoName(yname string) string {
    if yname == "message" { return "Message" }
    return ""
}

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetSegmentPath() string {
    return "parser-msg-ignore" + "[message='" + fmt.Sprintf("%v", parserMsgIgnore.Message) + "']"
}

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["message"] = parserMsgIgnore.Message
    return leafs
}

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetBundleName() string { return "cisco_ios_xe" }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetYangName() string { return "parser-msg-ignore" }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) SetParent(parent types.Entity) { parserMsgIgnore.parent = parent }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetParent() types.Entity { return parserMsgIgnore.parent }

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetParentYangName() string { return "cisco-ia" }

// NetconfYang_CiscoIa_ConfParserMsgIgnore
// Parser output from configuration 
// change that is informational only,
// not an error
type NetconfYang_CiscoIa_ConfParserMsgIgnore struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression to match parser output to be
    // ignored. The type is string with length: 1..255.
    Message interface{}
}

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetFilter() yfilter.YFilter { return confParserMsgIgnore.YFilter }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) SetFilter(yf yfilter.YFilter) { confParserMsgIgnore.YFilter = yf }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetGoName(yname string) string {
    if yname == "message" { return "Message" }
    return ""
}

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetSegmentPath() string {
    return "conf-parser-msg-ignore" + "[message='" + fmt.Sprintf("%v", confParserMsgIgnore.Message) + "']"
}

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["message"] = confParserMsgIgnore.Message
    return leafs
}

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetBundleName() string { return "cisco_ios_xe" }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetYangName() string { return "conf-parser-msg-ignore" }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) SetParent(parent types.Entity) { confParserMsgIgnore.parent = parent }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetParent() types.Entity { return confParserMsgIgnore.parent }

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetParentYangName() string { return "cisco-ia" }

// NetconfYang_CiscoIa_FullSyncCli
// IOS commands that result in other
// automatic configurations being applied
// for which a complete sync is required
type NetconfYang_CiscoIa_FullSyncCli struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression matching command lines which
    // cause other automatic configuration changes. The type is string with
    // length: 1..255.
    Command interface{}
}

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetFilter() yfilter.YFilter { return fullSyncCli.YFilter }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) SetFilter(yf yfilter.YFilter) { fullSyncCli.YFilter = yf }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetGoName(yname string) string {
    if yname == "command" { return "Command" }
    return ""
}

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetSegmentPath() string {
    return "full-sync-cli" + "[command='" + fmt.Sprintf("%v", fullSyncCli.Command) + "']"
}

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["command"] = fullSyncCli.Command
    return leafs
}

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetBundleName() string { return "cisco_ios_xe" }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetYangName() string { return "full-sync-cli" }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) SetParent(parent types.Entity) { fullSyncCli.parent = parent }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetParent() types.Entity { return fullSyncCli.parent }

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetParentYangName() string { return "cisco-ia" }

// NetconfYang_CiscoIa_ConfFullSyncCli
// A user-configurable list of IOS commands 
// that result in other automatic configurations 
// being applied for which a complete sync 
// is required
type NetconfYang_CiscoIa_ConfFullSyncCli struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression matching command lines which
    // cause other automatic configuration changes. The type is string with
    // length: 1..255.
    Command interface{}
}

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetFilter() yfilter.YFilter { return confFullSyncCli.YFilter }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) SetFilter(yf yfilter.YFilter) { confFullSyncCli.YFilter = yf }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetGoName(yname string) string {
    if yname == "command" { return "Command" }
    return ""
}

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetSegmentPath() string {
    return "conf-full-sync-cli" + "[command='" + fmt.Sprintf("%v", confFullSyncCli.Command) + "']"
}

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["command"] = confFullSyncCli.Command
    return leafs
}

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetBundleName() string { return "cisco_ios_xe" }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetYangName() string { return "conf-full-sync-cli" }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) SetParent(parent types.Entity) { confFullSyncCli.parent = parent }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetParent() types.Entity { return confFullSyncCli.parent }

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetParentYangName() string { return "cisco-ia" }

// NetconfYang_CiscoIa_Logging
// Controls logging behavior of DMI applications
type NetconfYang_CiscoIa_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Logging level for Confd. The type is SyslogSeverity. The default value is
    // error.
    ConfdLogLevel interface{}

    // Logging level for CiaAuthDaemon. The type is CiaLogLevel. The default value
    // is error.
    CiaauthdLogLevel interface{}

    // Logging level for Network Element Synchronizer. The type is CiaLogLevel.
    // The default value is error.
    NesLogLevel interface{}

    // Logging level for ONEP. The type is OnepLogLevel. The default value is
    // error.
    OnepLogLevel interface{}

    // Logging level for  Operational Data Manager. The type is CiaLogLevel. The
    // default value is error.
    OdmLogLevel interface{}

    // Logging level for Sync-From Daemon. The type is CiaLogLevel. The default
    // value is error.
    SyncLogLevel interface{}
}

func (logging *NetconfYang_CiscoIa_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *NetconfYang_CiscoIa_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *NetconfYang_CiscoIa_Logging) GetGoName(yname string) string {
    if yname == "confd-log-level" { return "ConfdLogLevel" }
    if yname == "ciaauthd-log-level" { return "CiaauthdLogLevel" }
    if yname == "nes-log-level" { return "NesLogLevel" }
    if yname == "onep-log-level" { return "OnepLogLevel" }
    if yname == "odm-log-level" { return "OdmLogLevel" }
    if yname == "sync-log-level" { return "SyncLogLevel" }
    return ""
}

func (logging *NetconfYang_CiscoIa_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *NetconfYang_CiscoIa_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logging *NetconfYang_CiscoIa_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logging *NetconfYang_CiscoIa_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["confd-log-level"] = logging.ConfdLogLevel
    leafs["ciaauthd-log-level"] = logging.CiaauthdLogLevel
    leafs["nes-log-level"] = logging.NesLogLevel
    leafs["onep-log-level"] = logging.OnepLogLevel
    leafs["odm-log-level"] = logging.OdmLogLevel
    leafs["sync-log-level"] = logging.SyncLogLevel
    return leafs
}

func (logging *NetconfYang_CiscoIa_Logging) GetBundleName() string { return "cisco_ios_xe" }

func (logging *NetconfYang_CiscoIa_Logging) GetYangName() string { return "logging" }

func (logging *NetconfYang_CiscoIa_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (logging *NetconfYang_CiscoIa_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (logging *NetconfYang_CiscoIa_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (logging *NetconfYang_CiscoIa_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *NetconfYang_CiscoIa_Logging) GetParent() types.Entity { return logging.parent }

func (logging *NetconfYang_CiscoIa_Logging) GetParentYangName() string { return "cisco-ia" }

// NetconfYang_CiscoIa_Blocking
// Controls blocking of command lines, either 
// from the NE to Confd or disallowing
// manual input from the console/vty
type NetconfYang_CiscoIa_Blocking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enables blocking of designated command lines via the  network element's
    // CLI. The type is bool. The default value is false.
    CliBlockingEnabled interface{}

    // Enables blocking of designated command lines via the  network element's
    // CLI. The type is bool. The default value is true.
    ConfdCfgBlockingEnabled interface{}

    // Command line pattern to disallow via the network element's CLI. The type is
    // slice of NetconfYang_CiscoIa_Blocking_NetworkElementCommand.
    NetworkElementCommand []NetconfYang_CiscoIa_Blocking_NetworkElementCommand

    // Command line pattern to omit syncing to Confd CDB. The type is slice of
    // NetconfYang_CiscoIa_Blocking_ConfdCfgCommand.
    ConfdCfgCommand []NetconfYang_CiscoIa_Blocking_ConfdCfgCommand
}

func (blocking *NetconfYang_CiscoIa_Blocking) GetFilter() yfilter.YFilter { return blocking.YFilter }

func (blocking *NetconfYang_CiscoIa_Blocking) SetFilter(yf yfilter.YFilter) { blocking.YFilter = yf }

func (blocking *NetconfYang_CiscoIa_Blocking) GetGoName(yname string) string {
    if yname == "cli-blocking-enabled" { return "CliBlockingEnabled" }
    if yname == "confd-cfg-blocking-enabled" { return "ConfdCfgBlockingEnabled" }
    if yname == "network-element-command" { return "NetworkElementCommand" }
    if yname == "confd-cfg-command" { return "ConfdCfgCommand" }
    return ""
}

func (blocking *NetconfYang_CiscoIa_Blocking) GetSegmentPath() string {
    return "blocking"
}

func (blocking *NetconfYang_CiscoIa_Blocking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-element-command" {
        for _, c := range blocking.NetworkElementCommand {
            if blocking.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoIa_Blocking_NetworkElementCommand{}
        blocking.NetworkElementCommand = append(blocking.NetworkElementCommand, child)
        return &blocking.NetworkElementCommand[len(blocking.NetworkElementCommand)-1]
    }
    if childYangName == "confd-cfg-command" {
        for _, c := range blocking.ConfdCfgCommand {
            if blocking.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoIa_Blocking_ConfdCfgCommand{}
        blocking.ConfdCfgCommand = append(blocking.ConfdCfgCommand, child)
        return &blocking.ConfdCfgCommand[len(blocking.ConfdCfgCommand)-1]
    }
    return nil
}

func (blocking *NetconfYang_CiscoIa_Blocking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range blocking.NetworkElementCommand {
        children[blocking.NetworkElementCommand[i].GetSegmentPath()] = &blocking.NetworkElementCommand[i]
    }
    for i := range blocking.ConfdCfgCommand {
        children[blocking.ConfdCfgCommand[i].GetSegmentPath()] = &blocking.ConfdCfgCommand[i]
    }
    return children
}

func (blocking *NetconfYang_CiscoIa_Blocking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cli-blocking-enabled"] = blocking.CliBlockingEnabled
    leafs["confd-cfg-blocking-enabled"] = blocking.ConfdCfgBlockingEnabled
    return leafs
}

func (blocking *NetconfYang_CiscoIa_Blocking) GetBundleName() string { return "cisco_ios_xe" }

func (blocking *NetconfYang_CiscoIa_Blocking) GetYangName() string { return "blocking" }

func (blocking *NetconfYang_CiscoIa_Blocking) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (blocking *NetconfYang_CiscoIa_Blocking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (blocking *NetconfYang_CiscoIa_Blocking) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (blocking *NetconfYang_CiscoIa_Blocking) SetParent(parent types.Entity) { blocking.parent = parent }

func (blocking *NetconfYang_CiscoIa_Blocking) GetParent() types.Entity { return blocking.parent }

func (blocking *NetconfYang_CiscoIa_Blocking) GetParentYangName() string { return "cisco-ia" }

// NetconfYang_CiscoIa_Blocking_NetworkElementCommand
// Command line pattern to disallow via the
// network element's CLI
type NetconfYang_CiscoIa_Blocking_NetworkElementCommand struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression matching command lines which
    // should be blocked from entry via console/vty. The type is string with
    // length: 1..255.
    Command interface{}
}

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetFilter() yfilter.YFilter { return networkElementCommand.YFilter }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) SetFilter(yf yfilter.YFilter) { networkElementCommand.YFilter = yf }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetGoName(yname string) string {
    if yname == "command" { return "Command" }
    return ""
}

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetSegmentPath() string {
    return "network-element-command" + "[command='" + fmt.Sprintf("%v", networkElementCommand.Command) + "']"
}

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["command"] = networkElementCommand.Command
    return leafs
}

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetBundleName() string { return "cisco_ios_xe" }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetYangName() string { return "network-element-command" }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) SetParent(parent types.Entity) { networkElementCommand.parent = parent }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetParent() types.Entity { return networkElementCommand.parent }

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetParentYangName() string { return "blocking" }

// NetconfYang_CiscoIa_Blocking_ConfdCfgCommand
// Command line pattern to omit syncing to Confd CDB
type NetconfYang_CiscoIa_Blocking_ConfdCfgCommand struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression matching command lines which
    // should be blocked from being sent to Confd from the network element. The
    // type is string.
    Command interface{}
}

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetFilter() yfilter.YFilter { return confdCfgCommand.YFilter }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) SetFilter(yf yfilter.YFilter) { confdCfgCommand.YFilter = yf }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetGoName(yname string) string {
    if yname == "command" { return "Command" }
    return ""
}

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetSegmentPath() string {
    return "confd-cfg-command" + "[command='" + fmt.Sprintf("%v", confdCfgCommand.Command) + "']"
}

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["command"] = confdCfgCommand.Command
    return leafs
}

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetBundleName() string { return "cisco_ios_xe" }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetYangName() string { return "confd-cfg-command" }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) SetParent(parent types.Entity) { confdCfgCommand.parent = parent }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetParent() types.Entity { return confdCfgCommand.parent }

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetParentYangName() string { return "blocking" }

// NetconfYang_CiscoOdm
type NetconfYang_CiscoOdm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is bool. The default value is false.
    PollingEnable interface{}

    // The type is interface{} with range: 500..4294967295. The default value is
    // 30000.
    OnDemandDefaultTime interface{}

    // The type is bool. The default value is false.
    OnDemandEnable interface{}

    // The type is slice of NetconfYang_CiscoOdm_Actions.
    Actions []NetconfYang_CiscoOdm_Actions
}

func (ciscoOdm *NetconfYang_CiscoOdm) GetFilter() yfilter.YFilter { return ciscoOdm.YFilter }

func (ciscoOdm *NetconfYang_CiscoOdm) SetFilter(yf yfilter.YFilter) { ciscoOdm.YFilter = yf }

func (ciscoOdm *NetconfYang_CiscoOdm) GetGoName(yname string) string {
    if yname == "polling-enable" { return "PollingEnable" }
    if yname == "on-demand-default-time" { return "OnDemandDefaultTime" }
    if yname == "on-demand-enable" { return "OnDemandEnable" }
    if yname == "actions" { return "Actions" }
    return ""
}

func (ciscoOdm *NetconfYang_CiscoOdm) GetSegmentPath() string {
    return "cisco-odm:cisco-odm"
}

func (ciscoOdm *NetconfYang_CiscoOdm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "actions" {
        for _, c := range ciscoOdm.Actions {
            if ciscoOdm.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := NetconfYang_CiscoOdm_Actions{}
        ciscoOdm.Actions = append(ciscoOdm.Actions, child)
        return &ciscoOdm.Actions[len(ciscoOdm.Actions)-1]
    }
    return nil
}

func (ciscoOdm *NetconfYang_CiscoOdm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoOdm.Actions {
        children[ciscoOdm.Actions[i].GetSegmentPath()] = &ciscoOdm.Actions[i]
    }
    return children
}

func (ciscoOdm *NetconfYang_CiscoOdm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["polling-enable"] = ciscoOdm.PollingEnable
    leafs["on-demand-default-time"] = ciscoOdm.OnDemandDefaultTime
    leafs["on-demand-enable"] = ciscoOdm.OnDemandEnable
    return leafs
}

func (ciscoOdm *NetconfYang_CiscoOdm) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoOdm *NetconfYang_CiscoOdm) GetYangName() string { return "cisco-odm" }

func (ciscoOdm *NetconfYang_CiscoOdm) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoOdm *NetconfYang_CiscoOdm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoOdm *NetconfYang_CiscoOdm) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoOdm *NetconfYang_CiscoOdm) SetParent(parent types.Entity) { ciscoOdm.parent = parent }

func (ciscoOdm *NetconfYang_CiscoOdm) GetParent() types.Entity { return ciscoOdm.parent }

func (ciscoOdm *NetconfYang_CiscoOdm) GetParentYangName() string { return "netconf-yang" }

// NetconfYang_CiscoOdm_Actions
type NetconfYang_CiscoOdm_Actions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is one of the following:
    // FlowMonitorVirtualServiceDiffservBridgeDomainMPLSStaticBindingEthernetCFMStatsMPLSForwardingTableBGPIPRoutePlatformSoftwareMPLSLDPNeighborsBFDNeighborsOSPF.
    ActionName interface{}

    // The type is interface{} with range: 1..4294967295. The default value is
    // 120000.
    PollingInterval interface{}

    // The type is Mode. The default value is poll.
    Mode interface{}

    // The type is string.
    CdbXpath interface{}
}

func (actions *NetconfYang_CiscoOdm_Actions) GetFilter() yfilter.YFilter { return actions.YFilter }

func (actions *NetconfYang_CiscoOdm_Actions) SetFilter(yf yfilter.YFilter) { actions.YFilter = yf }

func (actions *NetconfYang_CiscoOdm_Actions) GetGoName(yname string) string {
    if yname == "action-name" { return "ActionName" }
    if yname == "polling-interval" { return "PollingInterval" }
    if yname == "mode" { return "Mode" }
    if yname == "cdb-xpath" { return "CdbXpath" }
    return ""
}

func (actions *NetconfYang_CiscoOdm_Actions) GetSegmentPath() string {
    return "actions" + "[action-name='" + fmt.Sprintf("%v", actions.ActionName) + "']"
}

func (actions *NetconfYang_CiscoOdm_Actions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (actions *NetconfYang_CiscoOdm_Actions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (actions *NetconfYang_CiscoOdm_Actions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-name"] = actions.ActionName
    leafs["polling-interval"] = actions.PollingInterval
    leafs["mode"] = actions.Mode
    leafs["cdb-xpath"] = actions.CdbXpath
    return leafs
}

func (actions *NetconfYang_CiscoOdm_Actions) GetBundleName() string { return "cisco_ios_xe" }

func (actions *NetconfYang_CiscoOdm_Actions) GetYangName() string { return "actions" }

func (actions *NetconfYang_CiscoOdm_Actions) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (actions *NetconfYang_CiscoOdm_Actions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (actions *NetconfYang_CiscoOdm_Actions) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (actions *NetconfYang_CiscoOdm_Actions) SetParent(parent types.Entity) { actions.parent = parent }

func (actions *NetconfYang_CiscoOdm_Actions) GetParent() types.Entity { return actions.parent }

func (actions *NetconfYang_CiscoOdm_Actions) GetParentYangName() string { return "cisco-odm" }

// NetconfYang_CiscoOdm_Actions_Mode
type NetconfYang_CiscoOdm_Actions_Mode string

const (
    NetconfYang_CiscoOdm_Actions_Mode_poll NetconfYang_CiscoOdm_Actions_Mode = "poll"

    NetconfYang_CiscoOdm_Actions_Mode_on_demand NetconfYang_CiscoOdm_Actions_Mode = "on-demand"

    NetconfYang_CiscoOdm_Actions_Mode_none NetconfYang_CiscoOdm_Actions_Mode = "none"
)

