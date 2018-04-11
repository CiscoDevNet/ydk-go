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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Customize the behavior of the DMI applications.
    CiscoIa NetconfYang_CiscoIa
}

func (netconfYang *NetconfYang) GetEntityData() *types.CommonEntityData {
    netconfYang.EntityData.YFilter = netconfYang.YFilter
    netconfYang.EntityData.YangName = "netconf-yang"
    netconfYang.EntityData.BundleName = "cisco_ios_xe"
    netconfYang.EntityData.ParentYangName = "cisco-self-mgmt"
    netconfYang.EntityData.SegmentPath = "cisco-self-mgmt:netconf-yang"
    netconfYang.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    netconfYang.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    netconfYang.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    netconfYang.EntityData.Children = make(map[string]types.YChild)
    netconfYang.EntityData.Children["cisco-ia:cisco-ia"] = types.YChild{"CiscoIa", &netconfYang.CiscoIa}
    netconfYang.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(netconfYang.EntityData)
}

// NetconfYang_CiscoIa
// Customize the behavior of the DMI applications
type NetconfYang_CiscoIa struct {
    EntityData types.CommonEntityData
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

func (ciscoIa *NetconfYang_CiscoIa) GetEntityData() *types.CommonEntityData {
    ciscoIa.EntityData.YFilter = ciscoIa.YFilter
    ciscoIa.EntityData.YangName = "cisco-ia"
    ciscoIa.EntityData.BundleName = "cisco_ios_xe"
    ciscoIa.EntityData.ParentYangName = "netconf-yang"
    ciscoIa.EntityData.SegmentPath = "cisco-ia:cisco-ia"
    ciscoIa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoIa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoIa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoIa.EntityData.Children = make(map[string]types.YChild)
    ciscoIa.EntityData.Children["snmp-trap-control"] = types.YChild{"SnmpTrapControl", &ciscoIa.SnmpTrapControl}
    ciscoIa.EntityData.Children["preserve-ned-path"] = types.YChild{"PreserveNedPath", nil}
    for i := range ciscoIa.PreserveNedPath {
        ciscoIa.EntityData.Children[types.GetSegmentPath(&ciscoIa.PreserveNedPath[i])] = types.YChild{"PreserveNedPath", &ciscoIa.PreserveNedPath[i]}
    }
    ciscoIa.EntityData.Children["parser-msg-ignore"] = types.YChild{"ParserMsgIgnore", nil}
    for i := range ciscoIa.ParserMsgIgnore {
        ciscoIa.EntityData.Children[types.GetSegmentPath(&ciscoIa.ParserMsgIgnore[i])] = types.YChild{"ParserMsgIgnore", &ciscoIa.ParserMsgIgnore[i]}
    }
    ciscoIa.EntityData.Children["conf-parser-msg-ignore"] = types.YChild{"ConfParserMsgIgnore", nil}
    for i := range ciscoIa.ConfParserMsgIgnore {
        ciscoIa.EntityData.Children[types.GetSegmentPath(&ciscoIa.ConfParserMsgIgnore[i])] = types.YChild{"ConfParserMsgIgnore", &ciscoIa.ConfParserMsgIgnore[i]}
    }
    ciscoIa.EntityData.Children["full-sync-cli"] = types.YChild{"FullSyncCli", nil}
    for i := range ciscoIa.FullSyncCli {
        ciscoIa.EntityData.Children[types.GetSegmentPath(&ciscoIa.FullSyncCli[i])] = types.YChild{"FullSyncCli", &ciscoIa.FullSyncCli[i]}
    }
    ciscoIa.EntityData.Children["conf-full-sync-cli"] = types.YChild{"ConfFullSyncCli", nil}
    for i := range ciscoIa.ConfFullSyncCli {
        ciscoIa.EntityData.Children[types.GetSegmentPath(&ciscoIa.ConfFullSyncCli[i])] = types.YChild{"ConfFullSyncCli", &ciscoIa.ConfFullSyncCli[i]}
    }
    ciscoIa.EntityData.Children["logging"] = types.YChild{"Logging", &ciscoIa.Logging}
    ciscoIa.EntityData.Children["blocking"] = types.YChild{"Blocking", &ciscoIa.Blocking}
    ciscoIa.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoIa.EntityData.Leafs["auto-sync"] = types.YLeaf{"AutoSync", ciscoIa.AutoSync}
    ciscoIa.EntityData.Leafs["init-sync"] = types.YLeaf{"InitSync", ciscoIa.InitSync}
    ciscoIa.EntityData.Leafs["intelligent-sync"] = types.YLeaf{"IntelligentSync", ciscoIa.IntelligentSync}
    ciscoIa.EntityData.Leafs["message-diag-level"] = types.YLeaf{"MessageDiagLevel", ciscoIa.MessageDiagLevel}
    ciscoIa.EntityData.Leafs["max-diag-messages-saved"] = types.YLeaf{"MaxDiagMessagesSaved", ciscoIa.MaxDiagMessagesSaved}
    ciscoIa.EntityData.Leafs["post-sync-acl-process"] = types.YLeaf{"PostSyncAclProcess", ciscoIa.PostSyncAclProcess}
    ciscoIa.EntityData.Leafs["config-change-delay"] = types.YLeaf{"ConfigChangeDelay", ciscoIa.ConfigChangeDelay}
    ciscoIa.EntityData.Leafs["process-missing-prc"] = types.YLeaf{"ProcessMissingPrc", ciscoIa.ProcessMissingPrc}
    ciscoIa.EntityData.Leafs["snmp-community-string"] = types.YLeaf{"SnmpCommunityString", ciscoIa.SnmpCommunityString}
    ciscoIa.EntityData.Leafs["preserve-paths-enabled"] = types.YLeaf{"PreservePathsEnabled", ciscoIa.PreservePathsEnabled}
    ciscoIa.EntityData.Leafs["nes-ttynum"] = types.YLeaf{"NesTtynum", ciscoIa.NesTtynum}
    ciscoIa.EntityData.Leafs["restored"] = types.YLeaf{"Restored", ciscoIa.Restored}
    return &(ciscoIa.EntityData)
}

// NetconfYang_CiscoIa_SnmpTrapControl
// This container describes the configuration parameters
// for SNMP Trap to NetConf notification processing.
type NetconfYang_CiscoIa_SnmpTrapControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf enables or disables forwarding for all SNMP traps. The type is
    // bool. The default value is true.
    GlobalForwarding interface{}

    // This list describes SNMP Traps that are  supported for automatic
    // translation to NetConf notifications. The type is slice of
    // NetconfYang_CiscoIa_SnmpTrapControl_TrapList.
    TrapList []NetconfYang_CiscoIa_SnmpTrapControl_TrapList
}

func (snmpTrapControl *NetconfYang_CiscoIa_SnmpTrapControl) GetEntityData() *types.CommonEntityData {
    snmpTrapControl.EntityData.YFilter = snmpTrapControl.YFilter
    snmpTrapControl.EntityData.YangName = "snmp-trap-control"
    snmpTrapControl.EntityData.BundleName = "cisco_ios_xe"
    snmpTrapControl.EntityData.ParentYangName = "cisco-ia"
    snmpTrapControl.EntityData.SegmentPath = "snmp-trap-control"
    snmpTrapControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpTrapControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpTrapControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpTrapControl.EntityData.Children = make(map[string]types.YChild)
    snmpTrapControl.EntityData.Children["trap-list"] = types.YChild{"TrapList", nil}
    for i := range snmpTrapControl.TrapList {
        snmpTrapControl.EntityData.Children[types.GetSegmentPath(&snmpTrapControl.TrapList[i])] = types.YChild{"TrapList", &snmpTrapControl.TrapList[i]}
    }
    snmpTrapControl.EntityData.Leafs = make(map[string]types.YLeaf)
    snmpTrapControl.EntityData.Leafs["global-forwarding"] = types.YLeaf{"GlobalForwarding", snmpTrapControl.GlobalForwarding}
    return &(snmpTrapControl.EntityData)
}

// NetconfYang_CiscoIa_SnmpTrapControl_TrapList
// This list describes SNMP Traps that are 
// supported for automatic translation to NetConf
// notifications.
type NetconfYang_CiscoIa_SnmpTrapControl_TrapList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf contains the OID for the  SNMP trap to
    // be forwarded. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    TrapOid interface{}

    // This leaf contains the name of the SNMP trap to be  forwarded. The type is
    // string.
    Description interface{}

    // This leaf enables or disables forwarding for this SNMP trap. The type is
    // bool. The default value is true.
    Forward interface{}
}

func (trapList *NetconfYang_CiscoIa_SnmpTrapControl_TrapList) GetEntityData() *types.CommonEntityData {
    trapList.EntityData.YFilter = trapList.YFilter
    trapList.EntityData.YangName = "trap-list"
    trapList.EntityData.BundleName = "cisco_ios_xe"
    trapList.EntityData.ParentYangName = "snmp-trap-control"
    trapList.EntityData.SegmentPath = "trap-list" + "[trap-oid='" + fmt.Sprintf("%v", trapList.TrapOid) + "']"
    trapList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    trapList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    trapList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    trapList.EntityData.Children = make(map[string]types.YChild)
    trapList.EntityData.Leafs = make(map[string]types.YLeaf)
    trapList.EntityData.Leafs["trap-oid"] = types.YLeaf{"TrapOid", trapList.TrapOid}
    trapList.EntityData.Leafs["description"] = types.YLeaf{"Description", trapList.Description}
    trapList.EntityData.Leafs["forward"] = types.YLeaf{"Forward", trapList.Forward}
    return &(trapList.EntityData)
}

// NetconfYang_CiscoIa_PreserveNedPath
// Model paths from the NED model to preserve
// upon a sync from the network element.
// These paths are not cleared from the 
// running data store prior to the sync.
// These are expressed as nodes separated by 
// slash '/', e.g.  /native/ip/access-list
type NetconfYang_CiscoIa_PreserveNedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An XPath from the NED model. The type is string
    // with length: 1..1024.
    Xpath interface{}
}

func (preserveNedPath *NetconfYang_CiscoIa_PreserveNedPath) GetEntityData() *types.CommonEntityData {
    preserveNedPath.EntityData.YFilter = preserveNedPath.YFilter
    preserveNedPath.EntityData.YangName = "preserve-ned-path"
    preserveNedPath.EntityData.BundleName = "cisco_ios_xe"
    preserveNedPath.EntityData.ParentYangName = "cisco-ia"
    preserveNedPath.EntityData.SegmentPath = "preserve-ned-path" + "[xpath='" + fmt.Sprintf("%v", preserveNedPath.Xpath) + "']"
    preserveNedPath.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    preserveNedPath.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    preserveNedPath.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    preserveNedPath.EntityData.Children = make(map[string]types.YChild)
    preserveNedPath.EntityData.Leafs = make(map[string]types.YLeaf)
    preserveNedPath.EntityData.Leafs["xpath"] = types.YLeaf{"Xpath", preserveNedPath.Xpath}
    return &(preserveNedPath.EntityData)
}

// NetconfYang_CiscoIa_ParserMsgIgnore
// Parser output from configuration 
// change that is informational only,
// not an error. This is a read only
// list containing known informational 
// messages
type NetconfYang_CiscoIa_ParserMsgIgnore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression to match parser output to be
    // ignored. The type is string with length: 1..255.
    Message interface{}
}

func (parserMsgIgnore *NetconfYang_CiscoIa_ParserMsgIgnore) GetEntityData() *types.CommonEntityData {
    parserMsgIgnore.EntityData.YFilter = parserMsgIgnore.YFilter
    parserMsgIgnore.EntityData.YangName = "parser-msg-ignore"
    parserMsgIgnore.EntityData.BundleName = "cisco_ios_xe"
    parserMsgIgnore.EntityData.ParentYangName = "cisco-ia"
    parserMsgIgnore.EntityData.SegmentPath = "parser-msg-ignore" + "[message='" + fmt.Sprintf("%v", parserMsgIgnore.Message) + "']"
    parserMsgIgnore.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    parserMsgIgnore.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    parserMsgIgnore.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    parserMsgIgnore.EntityData.Children = make(map[string]types.YChild)
    parserMsgIgnore.EntityData.Leafs = make(map[string]types.YLeaf)
    parserMsgIgnore.EntityData.Leafs["message"] = types.YLeaf{"Message", parserMsgIgnore.Message}
    return &(parserMsgIgnore.EntityData)
}

// NetconfYang_CiscoIa_ConfParserMsgIgnore
// Parser output from configuration 
// change that is informational only,
// not an error
type NetconfYang_CiscoIa_ConfParserMsgIgnore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression to match parser output to be
    // ignored. The type is string with length: 1..255.
    Message interface{}
}

func (confParserMsgIgnore *NetconfYang_CiscoIa_ConfParserMsgIgnore) GetEntityData() *types.CommonEntityData {
    confParserMsgIgnore.EntityData.YFilter = confParserMsgIgnore.YFilter
    confParserMsgIgnore.EntityData.YangName = "conf-parser-msg-ignore"
    confParserMsgIgnore.EntityData.BundleName = "cisco_ios_xe"
    confParserMsgIgnore.EntityData.ParentYangName = "cisco-ia"
    confParserMsgIgnore.EntityData.SegmentPath = "conf-parser-msg-ignore" + "[message='" + fmt.Sprintf("%v", confParserMsgIgnore.Message) + "']"
    confParserMsgIgnore.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    confParserMsgIgnore.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    confParserMsgIgnore.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    confParserMsgIgnore.EntityData.Children = make(map[string]types.YChild)
    confParserMsgIgnore.EntityData.Leafs = make(map[string]types.YLeaf)
    confParserMsgIgnore.EntityData.Leafs["message"] = types.YLeaf{"Message", confParserMsgIgnore.Message}
    return &(confParserMsgIgnore.EntityData)
}

// NetconfYang_CiscoIa_FullSyncCli
// IOS commands that result in other
// automatic configurations being applied
// for which a complete sync is required
type NetconfYang_CiscoIa_FullSyncCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression matching command lines which
    // cause other automatic configuration changes. The type is string with
    // length: 1..255.
    Command interface{}
}

func (fullSyncCli *NetconfYang_CiscoIa_FullSyncCli) GetEntityData() *types.CommonEntityData {
    fullSyncCli.EntityData.YFilter = fullSyncCli.YFilter
    fullSyncCli.EntityData.YangName = "full-sync-cli"
    fullSyncCli.EntityData.BundleName = "cisco_ios_xe"
    fullSyncCli.EntityData.ParentYangName = "cisco-ia"
    fullSyncCli.EntityData.SegmentPath = "full-sync-cli" + "[command='" + fmt.Sprintf("%v", fullSyncCli.Command) + "']"
    fullSyncCli.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fullSyncCli.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fullSyncCli.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fullSyncCli.EntityData.Children = make(map[string]types.YChild)
    fullSyncCli.EntityData.Leafs = make(map[string]types.YLeaf)
    fullSyncCli.EntityData.Leafs["command"] = types.YLeaf{"Command", fullSyncCli.Command}
    return &(fullSyncCli.EntityData)
}

// NetconfYang_CiscoIa_ConfFullSyncCli
// A user-configurable list of IOS commands 
// that result in other automatic configurations 
// being applied for which a complete sync 
// is required
type NetconfYang_CiscoIa_ConfFullSyncCli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression matching command lines which
    // cause other automatic configuration changes. The type is string with
    // length: 1..255.
    Command interface{}
}

func (confFullSyncCli *NetconfYang_CiscoIa_ConfFullSyncCli) GetEntityData() *types.CommonEntityData {
    confFullSyncCli.EntityData.YFilter = confFullSyncCli.YFilter
    confFullSyncCli.EntityData.YangName = "conf-full-sync-cli"
    confFullSyncCli.EntityData.BundleName = "cisco_ios_xe"
    confFullSyncCli.EntityData.ParentYangName = "cisco-ia"
    confFullSyncCli.EntityData.SegmentPath = "conf-full-sync-cli" + "[command='" + fmt.Sprintf("%v", confFullSyncCli.Command) + "']"
    confFullSyncCli.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    confFullSyncCli.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    confFullSyncCli.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    confFullSyncCli.EntityData.Children = make(map[string]types.YChild)
    confFullSyncCli.EntityData.Leafs = make(map[string]types.YLeaf)
    confFullSyncCli.EntityData.Leafs["command"] = types.YLeaf{"Command", confFullSyncCli.Command}
    return &(confFullSyncCli.EntityData)
}

// NetconfYang_CiscoIa_Logging
// Controls logging behavior of DMI applications
type NetconfYang_CiscoIa_Logging struct {
    EntityData types.CommonEntityData
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

func (logging *NetconfYang_CiscoIa_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xe"
    logging.EntityData.ParentYangName = "cisco-ia"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
    logging.EntityData.Leafs["confd-log-level"] = types.YLeaf{"ConfdLogLevel", logging.ConfdLogLevel}
    logging.EntityData.Leafs["ciaauthd-log-level"] = types.YLeaf{"CiaauthdLogLevel", logging.CiaauthdLogLevel}
    logging.EntityData.Leafs["nes-log-level"] = types.YLeaf{"NesLogLevel", logging.NesLogLevel}
    logging.EntityData.Leafs["onep-log-level"] = types.YLeaf{"OnepLogLevel", logging.OnepLogLevel}
    logging.EntityData.Leafs["odm-log-level"] = types.YLeaf{"OdmLogLevel", logging.OdmLogLevel}
    logging.EntityData.Leafs["sync-log-level"] = types.YLeaf{"SyncLogLevel", logging.SyncLogLevel}
    return &(logging.EntityData)
}

// NetconfYang_CiscoIa_Blocking
// Controls blocking of command lines, either 
// from the NE to Confd or disallowing
// manual input from the console/vty
type NetconfYang_CiscoIa_Blocking struct {
    EntityData types.CommonEntityData
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

func (blocking *NetconfYang_CiscoIa_Blocking) GetEntityData() *types.CommonEntityData {
    blocking.EntityData.YFilter = blocking.YFilter
    blocking.EntityData.YangName = "blocking"
    blocking.EntityData.BundleName = "cisco_ios_xe"
    blocking.EntityData.ParentYangName = "cisco-ia"
    blocking.EntityData.SegmentPath = "blocking"
    blocking.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    blocking.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    blocking.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    blocking.EntityData.Children = make(map[string]types.YChild)
    blocking.EntityData.Children["network-element-command"] = types.YChild{"NetworkElementCommand", nil}
    for i := range blocking.NetworkElementCommand {
        blocking.EntityData.Children[types.GetSegmentPath(&blocking.NetworkElementCommand[i])] = types.YChild{"NetworkElementCommand", &blocking.NetworkElementCommand[i]}
    }
    blocking.EntityData.Children["confd-cfg-command"] = types.YChild{"ConfdCfgCommand", nil}
    for i := range blocking.ConfdCfgCommand {
        blocking.EntityData.Children[types.GetSegmentPath(&blocking.ConfdCfgCommand[i])] = types.YChild{"ConfdCfgCommand", &blocking.ConfdCfgCommand[i]}
    }
    blocking.EntityData.Leafs = make(map[string]types.YLeaf)
    blocking.EntityData.Leafs["cli-blocking-enabled"] = types.YLeaf{"CliBlockingEnabled", blocking.CliBlockingEnabled}
    blocking.EntityData.Leafs["confd-cfg-blocking-enabled"] = types.YLeaf{"ConfdCfgBlockingEnabled", blocking.ConfdCfgBlockingEnabled}
    return &(blocking.EntityData)
}

// NetconfYang_CiscoIa_Blocking_NetworkElementCommand
// Command line pattern to disallow via the
// network element's CLI
type NetconfYang_CiscoIa_Blocking_NetworkElementCommand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression matching command lines which
    // should be blocked from entry via console/vty. The type is string with
    // length: 1..255.
    Command interface{}
}

func (networkElementCommand *NetconfYang_CiscoIa_Blocking_NetworkElementCommand) GetEntityData() *types.CommonEntityData {
    networkElementCommand.EntityData.YFilter = networkElementCommand.YFilter
    networkElementCommand.EntityData.YangName = "network-element-command"
    networkElementCommand.EntityData.BundleName = "cisco_ios_xe"
    networkElementCommand.EntityData.ParentYangName = "blocking"
    networkElementCommand.EntityData.SegmentPath = "network-element-command" + "[command='" + fmt.Sprintf("%v", networkElementCommand.Command) + "']"
    networkElementCommand.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkElementCommand.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkElementCommand.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkElementCommand.EntityData.Children = make(map[string]types.YChild)
    networkElementCommand.EntityData.Leafs = make(map[string]types.YLeaf)
    networkElementCommand.EntityData.Leafs["command"] = types.YLeaf{"Command", networkElementCommand.Command}
    return &(networkElementCommand.EntityData)
}

// NetconfYang_CiscoIa_Blocking_ConfdCfgCommand
// Command line pattern to omit syncing to Confd CDB
type NetconfYang_CiscoIa_Blocking_ConfdCfgCommand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A regular expression matching command lines which
    // should be blocked from being sent to Confd from the network element. The
    // type is string.
    Command interface{}
}

func (confdCfgCommand *NetconfYang_CiscoIa_Blocking_ConfdCfgCommand) GetEntityData() *types.CommonEntityData {
    confdCfgCommand.EntityData.YFilter = confdCfgCommand.YFilter
    confdCfgCommand.EntityData.YangName = "confd-cfg-command"
    confdCfgCommand.EntityData.BundleName = "cisco_ios_xe"
    confdCfgCommand.EntityData.ParentYangName = "blocking"
    confdCfgCommand.EntityData.SegmentPath = "confd-cfg-command" + "[command='" + fmt.Sprintf("%v", confdCfgCommand.Command) + "']"
    confdCfgCommand.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    confdCfgCommand.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    confdCfgCommand.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    confdCfgCommand.EntityData.Children = make(map[string]types.YChild)
    confdCfgCommand.EntityData.Leafs = make(map[string]types.YLeaf)
    confdCfgCommand.EntityData.Leafs["command"] = types.YLeaf{"Command", confdCfgCommand.Command}
    return &(confdCfgCommand.EntityData)
}

