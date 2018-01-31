// This module defines status objects for monitoring of ConfD.
package tailf_confd_monitoring

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tailf_confd_monitoring"))
    ydk.RegisterEntity("{http://tail-f.com/yang/confd-monitoring confd-state}", reflect.TypeOf(ConfdState{}))
    ydk.RegisterEntity("tailf-confd-monitoring:confd-state", reflect.TypeOf(ConfdState{}))
}

// ConfdState
type ConfdState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tail-f product version number. The type is string.
    Version interface{}

    // Indicates whether an enhanced poll() function is used. The type is bool.
    Epoll interface{}

    // The type is DaemonStatus.
    DaemonStatus interface{}

    // The type is interface{}.
    ReadOnlyMode interface{}

    // Note that if the node is in upgrade mode, it is not possible to get any
    // information from the system over NETCONF. The error-app-tag will be
    // upgrade-in-progress.  Existing CLI sessions can get system information. The
    // type is interface{}.
    UpgradeMode interface{}

    
    Smp ConfdState_Smp

    
    Ha ConfdState_Ha

    
    LoadedDataModels ConfdState_LoadedDataModels

    
    Netconf ConfdState_Netconf

    
    Cli ConfdState_Cli

    
    Webui ConfdState_Webui

    
    Rest ConfdState_Rest

    
    Snmp ConfdState_Snmp

    
    Internal ConfdState_Internal
}

func (confdState *ConfdState) GetFilter() yfilter.YFilter { return confdState.YFilter }

func (confdState *ConfdState) SetFilter(yf yfilter.YFilter) { confdState.YFilter = yf }

func (confdState *ConfdState) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "epoll" { return "Epoll" }
    if yname == "daemon-status" { return "DaemonStatus" }
    if yname == "read-only-mode" { return "ReadOnlyMode" }
    if yname == "upgrade-mode" { return "UpgradeMode" }
    if yname == "smp" { return "Smp" }
    if yname == "ha" { return "Ha" }
    if yname == "loaded-data-models" { return "LoadedDataModels" }
    if yname == "netconf" { return "Netconf" }
    if yname == "cli" { return "Cli" }
    if yname == "webui" { return "Webui" }
    if yname == "rest" { return "Rest" }
    if yname == "snmp" { return "Snmp" }
    if yname == "internal" { return "Internal" }
    return ""
}

func (confdState *ConfdState) GetSegmentPath() string {
    return "tailf-confd-monitoring:confd-state"
}

func (confdState *ConfdState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "smp" {
        return &confdState.Smp
    }
    if childYangName == "ha" {
        return &confdState.Ha
    }
    if childYangName == "loaded-data-models" {
        return &confdState.LoadedDataModels
    }
    if childYangName == "netconf" {
        return &confdState.Netconf
    }
    if childYangName == "cli" {
        return &confdState.Cli
    }
    if childYangName == "webui" {
        return &confdState.Webui
    }
    if childYangName == "rest" {
        return &confdState.Rest
    }
    if childYangName == "snmp" {
        return &confdState.Snmp
    }
    if childYangName == "internal" {
        return &confdState.Internal
    }
    return nil
}

func (confdState *ConfdState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["smp"] = &confdState.Smp
    children["ha"] = &confdState.Ha
    children["loaded-data-models"] = &confdState.LoadedDataModels
    children["netconf"] = &confdState.Netconf
    children["cli"] = &confdState.Cli
    children["webui"] = &confdState.Webui
    children["rest"] = &confdState.Rest
    children["snmp"] = &confdState.Snmp
    children["internal"] = &confdState.Internal
    return children
}

func (confdState *ConfdState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = confdState.Version
    leafs["epoll"] = confdState.Epoll
    leafs["daemon-status"] = confdState.DaemonStatus
    leafs["read-only-mode"] = confdState.ReadOnlyMode
    leafs["upgrade-mode"] = confdState.UpgradeMode
    return leafs
}

func (confdState *ConfdState) GetBundleName() string { return "cisco_ios_xe" }

func (confdState *ConfdState) GetYangName() string { return "confd-state" }

func (confdState *ConfdState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (confdState *ConfdState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (confdState *ConfdState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (confdState *ConfdState) SetParent(parent types.Entity) { confdState.parent = parent }

func (confdState *ConfdState) GetParent() types.Entity { return confdState.parent }

func (confdState *ConfdState) GetParentYangName() string { return "tailf-confd-monitoring" }

// ConfdState_Smp
// This type is a presence type.
type ConfdState_Smp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of threads used by the daemon. The type is interface{} with range:
    // 0..65535.
    NumberOfThreads interface{}
}

func (smp *ConfdState_Smp) GetFilter() yfilter.YFilter { return smp.YFilter }

func (smp *ConfdState_Smp) SetFilter(yf yfilter.YFilter) { smp.YFilter = yf }

func (smp *ConfdState_Smp) GetGoName(yname string) string {
    if yname == "number-of-threads" { return "NumberOfThreads" }
    return ""
}

func (smp *ConfdState_Smp) GetSegmentPath() string {
    return "smp"
}

func (smp *ConfdState_Smp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (smp *ConfdState_Smp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (smp *ConfdState_Smp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number-of-threads"] = smp.NumberOfThreads
    return leafs
}

func (smp *ConfdState_Smp) GetBundleName() string { return "cisco_ios_xe" }

func (smp *ConfdState_Smp) GetYangName() string { return "smp" }

func (smp *ConfdState_Smp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (smp *ConfdState_Smp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (smp *ConfdState_Smp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (smp *ConfdState_Smp) SetParent(parent types.Entity) { smp.parent = parent }

func (smp *ConfdState_Smp) GetParent() types.Entity { return smp.parent }

func (smp *ConfdState_Smp) GetParentYangName() string { return "confd-state" }

// ConfdState_Ha
// This type is a presence type.
type ConfdState_Ha struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The current HA mode of the node in the HA cluster. The type is Mode.
    Mode interface{}

    // The node identifier of this node in the HA cluster. The type is string.
    NodeId interface{}

    // The node identifier of this node's parent node. This is the HA cluster's
    // master node unless relay slaves are used. The type is string.
    MasterNodeId interface{}

    // The node identifiers of the currently connected slaves. The type is slice
    // of string.
    ConnectedSlave []interface{}

    // The node identifiers of slaves with pending acknowledgement of synchronous
    // replication. The type is slice of string.
    PendingSlave []interface{}
}

func (ha *ConfdState_Ha) GetFilter() yfilter.YFilter { return ha.YFilter }

func (ha *ConfdState_Ha) SetFilter(yf yfilter.YFilter) { ha.YFilter = yf }

func (ha *ConfdState_Ha) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "node-id" { return "NodeId" }
    if yname == "master-node-id" { return "MasterNodeId" }
    if yname == "connected-slave" { return "ConnectedSlave" }
    if yname == "pending-slave" { return "PendingSlave" }
    return ""
}

func (ha *ConfdState_Ha) GetSegmentPath() string {
    return "ha"
}

func (ha *ConfdState_Ha) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ha *ConfdState_Ha) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ha *ConfdState_Ha) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = ha.Mode
    leafs["node-id"] = ha.NodeId
    leafs["master-node-id"] = ha.MasterNodeId
    leafs["connected-slave"] = ha.ConnectedSlave
    leafs["pending-slave"] = ha.PendingSlave
    return leafs
}

func (ha *ConfdState_Ha) GetBundleName() string { return "cisco_ios_xe" }

func (ha *ConfdState_Ha) GetYangName() string { return "ha" }

func (ha *ConfdState_Ha) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ha *ConfdState_Ha) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ha *ConfdState_Ha) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ha *ConfdState_Ha) SetParent(parent types.Entity) { ha.parent = parent }

func (ha *ConfdState_Ha) GetParent() types.Entity { return ha.parent }

func (ha *ConfdState_Ha) GetParentYangName() string { return "confd-state" }

// ConfdState_Ha_Mode represents The current HA mode of the node in the HA cluster.
type ConfdState_Ha_Mode string

const (
    ConfdState_Ha_Mode_none ConfdState_Ha_Mode = "none"

    ConfdState_Ha_Mode_slave ConfdState_Ha_Mode = "slave"

    ConfdState_Ha_Mode_master ConfdState_Ha_Mode = "master"

    ConfdState_Ha_Mode_relay_slave ConfdState_Ha_Mode = "relay-slave"
)

// ConfdState_LoadedDataModels
type ConfdState_LoadedDataModels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This list contains all loaded YANG data modules.  This list is a superset
    // of the 'schema' list defined in ietf-netconf-monitoring, which only lists
    // modules exported through NETCONF. The type is slice of
    // ConfdState_LoadedDataModels_DataModel.
    DataModel []ConfdState_LoadedDataModels_DataModel
}

func (loadedDataModels *ConfdState_LoadedDataModels) GetFilter() yfilter.YFilter { return loadedDataModels.YFilter }

func (loadedDataModels *ConfdState_LoadedDataModels) SetFilter(yf yfilter.YFilter) { loadedDataModels.YFilter = yf }

func (loadedDataModels *ConfdState_LoadedDataModels) GetGoName(yname string) string {
    if yname == "data-model" { return "DataModel" }
    return ""
}

func (loadedDataModels *ConfdState_LoadedDataModels) GetSegmentPath() string {
    return "loaded-data-models"
}

func (loadedDataModels *ConfdState_LoadedDataModels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "data-model" {
        for _, c := range loadedDataModels.DataModel {
            if loadedDataModels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_LoadedDataModels_DataModel{}
        loadedDataModels.DataModel = append(loadedDataModels.DataModel, child)
        return &loadedDataModels.DataModel[len(loadedDataModels.DataModel)-1]
    }
    return nil
}

func (loadedDataModels *ConfdState_LoadedDataModels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range loadedDataModels.DataModel {
        children[loadedDataModels.DataModel[i].GetSegmentPath()] = &loadedDataModels.DataModel[i]
    }
    return children
}

func (loadedDataModels *ConfdState_LoadedDataModels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (loadedDataModels *ConfdState_LoadedDataModels) GetBundleName() string { return "cisco_ios_xe" }

func (loadedDataModels *ConfdState_LoadedDataModels) GetYangName() string { return "loaded-data-models" }

func (loadedDataModels *ConfdState_LoadedDataModels) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (loadedDataModels *ConfdState_LoadedDataModels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (loadedDataModels *ConfdState_LoadedDataModels) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (loadedDataModels *ConfdState_LoadedDataModels) SetParent(parent types.Entity) { loadedDataModels.parent = parent }

func (loadedDataModels *ConfdState_LoadedDataModels) GetParent() types.Entity { return loadedDataModels.parent }

func (loadedDataModels *ConfdState_LoadedDataModels) GetParentYangName() string { return "confd-state" }

// ConfdState_LoadedDataModels_DataModel
// This list contains all loaded YANG data modules.
// 
// This list is a superset of the 'schema' list defined in
// ietf-netconf-monitoring, which only lists modules exported
// through NETCONF.
type ConfdState_LoadedDataModels_DataModel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The YANG module name. The type is string.
    Name interface{}

    // The YANG module revision. The type is string.
    Revision interface{}

    // The YANG module namespace. The type is string.
    Namespace interface{}

    // The prefix defined in the YANG module. The type is string.
    Prefix interface{}

    // This leaf is present if the module is exported to all northbound
    // interfaces. The type is interface{}.
    ExportedToAll interface{}

    // A list of the contexts (northbound interfaces) this module is exported to.
    // The type is one of the following types: slice of  
    // :go:struct:`ConfdState_LoadedDataModels_DataModel_ExportedTo
    // <ydk/models/cisco_ios_xe/tailf_confd_monitoring/ConfdState_LoadedDataModels_DataModel_ExportedTo>`,
    // or slice of string.
    ExportedTo []interface{}
}

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetFilter() yfilter.YFilter { return dataModel.YFilter }

func (dataModel *ConfdState_LoadedDataModels_DataModel) SetFilter(yf yfilter.YFilter) { dataModel.YFilter = yf }

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "revision" { return "Revision" }
    if yname == "namespace" { return "Namespace" }
    if yname == "prefix" { return "Prefix" }
    if yname == "exported-to-all" { return "ExportedToAll" }
    if yname == "exported-to" { return "ExportedTo" }
    return ""
}

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetSegmentPath() string {
    return "data-model" + "[name='" + fmt.Sprintf("%v", dataModel.Name) + "']"
}

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = dataModel.Name
    leafs["revision"] = dataModel.Revision
    leafs["namespace"] = dataModel.Namespace
    leafs["prefix"] = dataModel.Prefix
    leafs["exported-to-all"] = dataModel.ExportedToAll
    leafs["exported-to"] = dataModel.ExportedTo
    return leafs
}

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetBundleName() string { return "cisco_ios_xe" }

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetYangName() string { return "data-model" }

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dataModel *ConfdState_LoadedDataModels_DataModel) SetParent(parent types.Entity) { dataModel.parent = parent }

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetParent() types.Entity { return dataModel.parent }

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetParentYangName() string { return "loaded-data-models" }

// ConfdState_LoadedDataModels_DataModel_ExportedTo represents is exported to.
type ConfdState_LoadedDataModels_DataModel_ExportedTo string

const (
    ConfdState_LoadedDataModels_DataModel_ExportedTo_netconf ConfdState_LoadedDataModels_DataModel_ExportedTo = "netconf"

    ConfdState_LoadedDataModels_DataModel_ExportedTo_cli ConfdState_LoadedDataModels_DataModel_ExportedTo = "cli"

    ConfdState_LoadedDataModels_DataModel_ExportedTo_webui ConfdState_LoadedDataModels_DataModel_ExportedTo = "webui"

    ConfdState_LoadedDataModels_DataModel_ExportedTo_rest ConfdState_LoadedDataModels_DataModel_ExportedTo = "rest"

    ConfdState_LoadedDataModels_DataModel_ExportedTo_snmp ConfdState_LoadedDataModels_DataModel_ExportedTo = "snmp"
)

// ConfdState_Netconf
// This type is a presence type.
type ConfdState_Netconf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The transport addresses the NETCONF server is listening on.  Note that
    // other mechanisms can be put in front of the TCP addresses below, e.g., an
    // OpenSSH server.  Such mechanisms are not known to the daemon and not listed
    // here.
    Listen ConfdState_Netconf_Listen
}

func (netconf *ConfdState_Netconf) GetFilter() yfilter.YFilter { return netconf.YFilter }

func (netconf *ConfdState_Netconf) SetFilter(yf yfilter.YFilter) { netconf.YFilter = yf }

func (netconf *ConfdState_Netconf) GetGoName(yname string) string {
    if yname == "listen" { return "Listen" }
    return ""
}

func (netconf *ConfdState_Netconf) GetSegmentPath() string {
    return "netconf"
}

func (netconf *ConfdState_Netconf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "listen" {
        return &netconf.Listen
    }
    return nil
}

func (netconf *ConfdState_Netconf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["listen"] = &netconf.Listen
    return children
}

func (netconf *ConfdState_Netconf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconf *ConfdState_Netconf) GetBundleName() string { return "cisco_ios_xe" }

func (netconf *ConfdState_Netconf) GetYangName() string { return "netconf" }

func (netconf *ConfdState_Netconf) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (netconf *ConfdState_Netconf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (netconf *ConfdState_Netconf) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (netconf *ConfdState_Netconf) SetParent(parent types.Entity) { netconf.parent = parent }

func (netconf *ConfdState_Netconf) GetParent() types.Entity { return netconf.parent }

func (netconf *ConfdState_Netconf) GetParentYangName() string { return "confd-state" }

// ConfdState_Netconf_Listen
// The transport addresses the NETCONF server is listening on.
// 
// Note that other mechanisms can be put in front of the TCP
// addresses below, e.g., an OpenSSH server.  Such mechanisms
// are not known to the daemon and not listed here.
type ConfdState_Netconf_Listen struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Netconf_Listen_Tcp.
    Tcp []ConfdState_Netconf_Listen_Tcp

    // The type is slice of ConfdState_Netconf_Listen_Ssh.
    Ssh []ConfdState_Netconf_Listen_Ssh
}

func (listen *ConfdState_Netconf_Listen) GetFilter() yfilter.YFilter { return listen.YFilter }

func (listen *ConfdState_Netconf_Listen) SetFilter(yf yfilter.YFilter) { listen.YFilter = yf }

func (listen *ConfdState_Netconf_Listen) GetGoName(yname string) string {
    if yname == "tcp" { return "Tcp" }
    if yname == "ssh" { return "Ssh" }
    return ""
}

func (listen *ConfdState_Netconf_Listen) GetSegmentPath() string {
    return "listen"
}

func (listen *ConfdState_Netconf_Listen) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcp" {
        for _, c := range listen.Tcp {
            if listen.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Netconf_Listen_Tcp{}
        listen.Tcp = append(listen.Tcp, child)
        return &listen.Tcp[len(listen.Tcp)-1]
    }
    if childYangName == "ssh" {
        for _, c := range listen.Ssh {
            if listen.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Netconf_Listen_Ssh{}
        listen.Ssh = append(listen.Ssh, child)
        return &listen.Ssh[len(listen.Ssh)-1]
    }
    return nil
}

func (listen *ConfdState_Netconf_Listen) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range listen.Tcp {
        children[listen.Tcp[i].GetSegmentPath()] = &listen.Tcp[i]
    }
    for i := range listen.Ssh {
        children[listen.Ssh[i].GetSegmentPath()] = &listen.Ssh[i]
    }
    return children
}

func (listen *ConfdState_Netconf_Listen) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (listen *ConfdState_Netconf_Listen) GetBundleName() string { return "cisco_ios_xe" }

func (listen *ConfdState_Netconf_Listen) GetYangName() string { return "listen" }

func (listen *ConfdState_Netconf_Listen) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (listen *ConfdState_Netconf_Listen) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (listen *ConfdState_Netconf_Listen) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (listen *ConfdState_Netconf_Listen) SetParent(parent types.Entity) { listen.parent = parent }

func (listen *ConfdState_Netconf_Listen) GetParent() types.Entity { return listen.parent }

func (listen *ConfdState_Netconf_Listen) GetParentYangName() string { return "netconf" }

// ConfdState_Netconf_Listen_Tcp
type ConfdState_Netconf_Listen_Tcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (tcp *ConfdState_Netconf_Listen_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *ConfdState_Netconf_Listen_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *ConfdState_Netconf_Listen_Tcp) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "port" { return "Port" }
    return ""
}

func (tcp *ConfdState_Netconf_Listen_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *ConfdState_Netconf_Listen_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *ConfdState_Netconf_Listen_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *ConfdState_Netconf_Listen_Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = tcp.Ip
    leafs["port"] = tcp.Port
    return leafs
}

func (tcp *ConfdState_Netconf_Listen_Tcp) GetBundleName() string { return "cisco_ios_xe" }

func (tcp *ConfdState_Netconf_Listen_Tcp) GetYangName() string { return "tcp" }

func (tcp *ConfdState_Netconf_Listen_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcp *ConfdState_Netconf_Listen_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcp *ConfdState_Netconf_Listen_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcp *ConfdState_Netconf_Listen_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *ConfdState_Netconf_Listen_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *ConfdState_Netconf_Listen_Tcp) GetParentYangName() string { return "listen" }

// ConfdState_Netconf_Listen_Ssh
type ConfdState_Netconf_Listen_Ssh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (ssh *ConfdState_Netconf_Listen_Ssh) GetFilter() yfilter.YFilter { return ssh.YFilter }

func (ssh *ConfdState_Netconf_Listen_Ssh) SetFilter(yf yfilter.YFilter) { ssh.YFilter = yf }

func (ssh *ConfdState_Netconf_Listen_Ssh) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "port" { return "Port" }
    return ""
}

func (ssh *ConfdState_Netconf_Listen_Ssh) GetSegmentPath() string {
    return "ssh"
}

func (ssh *ConfdState_Netconf_Listen_Ssh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssh *ConfdState_Netconf_Listen_Ssh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssh *ConfdState_Netconf_Listen_Ssh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = ssh.Ip
    leafs["port"] = ssh.Port
    return leafs
}

func (ssh *ConfdState_Netconf_Listen_Ssh) GetBundleName() string { return "cisco_ios_xe" }

func (ssh *ConfdState_Netconf_Listen_Ssh) GetYangName() string { return "ssh" }

func (ssh *ConfdState_Netconf_Listen_Ssh) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ssh *ConfdState_Netconf_Listen_Ssh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ssh *ConfdState_Netconf_Listen_Ssh) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ssh *ConfdState_Netconf_Listen_Ssh) SetParent(parent types.Entity) { ssh.parent = parent }

func (ssh *ConfdState_Netconf_Listen_Ssh) GetParent() types.Entity { return ssh.parent }

func (ssh *ConfdState_Netconf_Listen_Ssh) GetParentYangName() string { return "listen" }

// ConfdState_Cli
// This type is a presence type.
type ConfdState_Cli struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The transport addresses the CLI server is listening on.  In addition to the
    // SSH addresses listen below, the CLI can always be invoked through the
    // daemons IPC port.  Note that other mechanisms can be put in front of the
    // IPC port, e.g., an OpenSSH server.  Such mechanisms are not known to the
    // daemon and not listed here.
    Listen ConfdState_Cli_Listen
}

func (cli *ConfdState_Cli) GetFilter() yfilter.YFilter { return cli.YFilter }

func (cli *ConfdState_Cli) SetFilter(yf yfilter.YFilter) { cli.YFilter = yf }

func (cli *ConfdState_Cli) GetGoName(yname string) string {
    if yname == "listen" { return "Listen" }
    return ""
}

func (cli *ConfdState_Cli) GetSegmentPath() string {
    return "cli"
}

func (cli *ConfdState_Cli) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "listen" {
        return &cli.Listen
    }
    return nil
}

func (cli *ConfdState_Cli) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["listen"] = &cli.Listen
    return children
}

func (cli *ConfdState_Cli) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cli *ConfdState_Cli) GetBundleName() string { return "cisco_ios_xe" }

func (cli *ConfdState_Cli) GetYangName() string { return "cli" }

func (cli *ConfdState_Cli) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cli *ConfdState_Cli) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cli *ConfdState_Cli) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cli *ConfdState_Cli) SetParent(parent types.Entity) { cli.parent = parent }

func (cli *ConfdState_Cli) GetParent() types.Entity { return cli.parent }

func (cli *ConfdState_Cli) GetParentYangName() string { return "confd-state" }

// ConfdState_Cli_Listen
// The transport addresses the CLI server is listening on.
// 
// In addition to the SSH addresses listen below, the CLI can
// always be invoked through the daemons IPC port.
// 
// Note that other mechanisms can be put in front of the IPC
// port, e.g., an OpenSSH server.  Such mechanisms are not
// known to the daemon and not listed here.
type ConfdState_Cli_Listen struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Cli_Listen_Ssh.
    Ssh []ConfdState_Cli_Listen_Ssh
}

func (listen *ConfdState_Cli_Listen) GetFilter() yfilter.YFilter { return listen.YFilter }

func (listen *ConfdState_Cli_Listen) SetFilter(yf yfilter.YFilter) { listen.YFilter = yf }

func (listen *ConfdState_Cli_Listen) GetGoName(yname string) string {
    if yname == "ssh" { return "Ssh" }
    return ""
}

func (listen *ConfdState_Cli_Listen) GetSegmentPath() string {
    return "listen"
}

func (listen *ConfdState_Cli_Listen) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ssh" {
        for _, c := range listen.Ssh {
            if listen.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Cli_Listen_Ssh{}
        listen.Ssh = append(listen.Ssh, child)
        return &listen.Ssh[len(listen.Ssh)-1]
    }
    return nil
}

func (listen *ConfdState_Cli_Listen) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range listen.Ssh {
        children[listen.Ssh[i].GetSegmentPath()] = &listen.Ssh[i]
    }
    return children
}

func (listen *ConfdState_Cli_Listen) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (listen *ConfdState_Cli_Listen) GetBundleName() string { return "cisco_ios_xe" }

func (listen *ConfdState_Cli_Listen) GetYangName() string { return "listen" }

func (listen *ConfdState_Cli_Listen) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (listen *ConfdState_Cli_Listen) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (listen *ConfdState_Cli_Listen) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (listen *ConfdState_Cli_Listen) SetParent(parent types.Entity) { listen.parent = parent }

func (listen *ConfdState_Cli_Listen) GetParent() types.Entity { return listen.parent }

func (listen *ConfdState_Cli_Listen) GetParentYangName() string { return "cli" }

// ConfdState_Cli_Listen_Ssh
type ConfdState_Cli_Listen_Ssh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (ssh *ConfdState_Cli_Listen_Ssh) GetFilter() yfilter.YFilter { return ssh.YFilter }

func (ssh *ConfdState_Cli_Listen_Ssh) SetFilter(yf yfilter.YFilter) { ssh.YFilter = yf }

func (ssh *ConfdState_Cli_Listen_Ssh) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "port" { return "Port" }
    return ""
}

func (ssh *ConfdState_Cli_Listen_Ssh) GetSegmentPath() string {
    return "ssh"
}

func (ssh *ConfdState_Cli_Listen_Ssh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssh *ConfdState_Cli_Listen_Ssh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssh *ConfdState_Cli_Listen_Ssh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = ssh.Ip
    leafs["port"] = ssh.Port
    return leafs
}

func (ssh *ConfdState_Cli_Listen_Ssh) GetBundleName() string { return "cisco_ios_xe" }

func (ssh *ConfdState_Cli_Listen_Ssh) GetYangName() string { return "ssh" }

func (ssh *ConfdState_Cli_Listen_Ssh) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ssh *ConfdState_Cli_Listen_Ssh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ssh *ConfdState_Cli_Listen_Ssh) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ssh *ConfdState_Cli_Listen_Ssh) SetParent(parent types.Entity) { ssh.parent = parent }

func (ssh *ConfdState_Cli_Listen_Ssh) GetParent() types.Entity { return ssh.parent }

func (ssh *ConfdState_Cli_Listen_Ssh) GetParentYangName() string { return "listen" }

// ConfdState_Webui
// This type is a presence type.
type ConfdState_Webui struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The transport addresses the WebUI server is listening on.
    Listen ConfdState_Webui_Listen
}

func (webui *ConfdState_Webui) GetFilter() yfilter.YFilter { return webui.YFilter }

func (webui *ConfdState_Webui) SetFilter(yf yfilter.YFilter) { webui.YFilter = yf }

func (webui *ConfdState_Webui) GetGoName(yname string) string {
    if yname == "listen" { return "Listen" }
    return ""
}

func (webui *ConfdState_Webui) GetSegmentPath() string {
    return "webui"
}

func (webui *ConfdState_Webui) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "listen" {
        return &webui.Listen
    }
    return nil
}

func (webui *ConfdState_Webui) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["listen"] = &webui.Listen
    return children
}

func (webui *ConfdState_Webui) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (webui *ConfdState_Webui) GetBundleName() string { return "cisco_ios_xe" }

func (webui *ConfdState_Webui) GetYangName() string { return "webui" }

func (webui *ConfdState_Webui) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (webui *ConfdState_Webui) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (webui *ConfdState_Webui) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (webui *ConfdState_Webui) SetParent(parent types.Entity) { webui.parent = parent }

func (webui *ConfdState_Webui) GetParent() types.Entity { return webui.parent }

func (webui *ConfdState_Webui) GetParentYangName() string { return "confd-state" }

// ConfdState_Webui_Listen
// The transport addresses the WebUI server is listening on.
type ConfdState_Webui_Listen struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Webui_Listen_Tcp.
    Tcp []ConfdState_Webui_Listen_Tcp

    // The type is slice of ConfdState_Webui_Listen_Ssl.
    Ssl []ConfdState_Webui_Listen_Ssl
}

func (listen *ConfdState_Webui_Listen) GetFilter() yfilter.YFilter { return listen.YFilter }

func (listen *ConfdState_Webui_Listen) SetFilter(yf yfilter.YFilter) { listen.YFilter = yf }

func (listen *ConfdState_Webui_Listen) GetGoName(yname string) string {
    if yname == "tcp" { return "Tcp" }
    if yname == "ssl" { return "Ssl" }
    return ""
}

func (listen *ConfdState_Webui_Listen) GetSegmentPath() string {
    return "listen"
}

func (listen *ConfdState_Webui_Listen) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcp" {
        for _, c := range listen.Tcp {
            if listen.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Webui_Listen_Tcp{}
        listen.Tcp = append(listen.Tcp, child)
        return &listen.Tcp[len(listen.Tcp)-1]
    }
    if childYangName == "ssl" {
        for _, c := range listen.Ssl {
            if listen.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Webui_Listen_Ssl{}
        listen.Ssl = append(listen.Ssl, child)
        return &listen.Ssl[len(listen.Ssl)-1]
    }
    return nil
}

func (listen *ConfdState_Webui_Listen) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range listen.Tcp {
        children[listen.Tcp[i].GetSegmentPath()] = &listen.Tcp[i]
    }
    for i := range listen.Ssl {
        children[listen.Ssl[i].GetSegmentPath()] = &listen.Ssl[i]
    }
    return children
}

func (listen *ConfdState_Webui_Listen) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (listen *ConfdState_Webui_Listen) GetBundleName() string { return "cisco_ios_xe" }

func (listen *ConfdState_Webui_Listen) GetYangName() string { return "listen" }

func (listen *ConfdState_Webui_Listen) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (listen *ConfdState_Webui_Listen) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (listen *ConfdState_Webui_Listen) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (listen *ConfdState_Webui_Listen) SetParent(parent types.Entity) { listen.parent = parent }

func (listen *ConfdState_Webui_Listen) GetParent() types.Entity { return listen.parent }

func (listen *ConfdState_Webui_Listen) GetParentYangName() string { return "webui" }

// ConfdState_Webui_Listen_Tcp
type ConfdState_Webui_Listen_Tcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (tcp *ConfdState_Webui_Listen_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *ConfdState_Webui_Listen_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *ConfdState_Webui_Listen_Tcp) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "port" { return "Port" }
    return ""
}

func (tcp *ConfdState_Webui_Listen_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *ConfdState_Webui_Listen_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *ConfdState_Webui_Listen_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *ConfdState_Webui_Listen_Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = tcp.Ip
    leafs["port"] = tcp.Port
    return leafs
}

func (tcp *ConfdState_Webui_Listen_Tcp) GetBundleName() string { return "cisco_ios_xe" }

func (tcp *ConfdState_Webui_Listen_Tcp) GetYangName() string { return "tcp" }

func (tcp *ConfdState_Webui_Listen_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcp *ConfdState_Webui_Listen_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcp *ConfdState_Webui_Listen_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcp *ConfdState_Webui_Listen_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *ConfdState_Webui_Listen_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *ConfdState_Webui_Listen_Tcp) GetParentYangName() string { return "listen" }

// ConfdState_Webui_Listen_Ssl
type ConfdState_Webui_Listen_Ssl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (ssl *ConfdState_Webui_Listen_Ssl) GetFilter() yfilter.YFilter { return ssl.YFilter }

func (ssl *ConfdState_Webui_Listen_Ssl) SetFilter(yf yfilter.YFilter) { ssl.YFilter = yf }

func (ssl *ConfdState_Webui_Listen_Ssl) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "port" { return "Port" }
    return ""
}

func (ssl *ConfdState_Webui_Listen_Ssl) GetSegmentPath() string {
    return "ssl"
}

func (ssl *ConfdState_Webui_Listen_Ssl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssl *ConfdState_Webui_Listen_Ssl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssl *ConfdState_Webui_Listen_Ssl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = ssl.Ip
    leafs["port"] = ssl.Port
    return leafs
}

func (ssl *ConfdState_Webui_Listen_Ssl) GetBundleName() string { return "cisco_ios_xe" }

func (ssl *ConfdState_Webui_Listen_Ssl) GetYangName() string { return "ssl" }

func (ssl *ConfdState_Webui_Listen_Ssl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ssl *ConfdState_Webui_Listen_Ssl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ssl *ConfdState_Webui_Listen_Ssl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ssl *ConfdState_Webui_Listen_Ssl) SetParent(parent types.Entity) { ssl.parent = parent }

func (ssl *ConfdState_Webui_Listen_Ssl) GetParent() types.Entity { return ssl.parent }

func (ssl *ConfdState_Webui_Listen_Ssl) GetParentYangName() string { return "listen" }

// ConfdState_Rest
// This type is a presence type.
type ConfdState_Rest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The transport addresses the REST server is listening on.
    Listen ConfdState_Rest_Listen
}

func (rest *ConfdState_Rest) GetFilter() yfilter.YFilter { return rest.YFilter }

func (rest *ConfdState_Rest) SetFilter(yf yfilter.YFilter) { rest.YFilter = yf }

func (rest *ConfdState_Rest) GetGoName(yname string) string {
    if yname == "listen" { return "Listen" }
    return ""
}

func (rest *ConfdState_Rest) GetSegmentPath() string {
    return "rest"
}

func (rest *ConfdState_Rest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "listen" {
        return &rest.Listen
    }
    return nil
}

func (rest *ConfdState_Rest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["listen"] = &rest.Listen
    return children
}

func (rest *ConfdState_Rest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rest *ConfdState_Rest) GetBundleName() string { return "cisco_ios_xe" }

func (rest *ConfdState_Rest) GetYangName() string { return "rest" }

func (rest *ConfdState_Rest) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rest *ConfdState_Rest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rest *ConfdState_Rest) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rest *ConfdState_Rest) SetParent(parent types.Entity) { rest.parent = parent }

func (rest *ConfdState_Rest) GetParent() types.Entity { return rest.parent }

func (rest *ConfdState_Rest) GetParentYangName() string { return "confd-state" }

// ConfdState_Rest_Listen
// The transport addresses the REST server is listening on.
type ConfdState_Rest_Listen struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Rest_Listen_Tcp.
    Tcp []ConfdState_Rest_Listen_Tcp

    // The type is slice of ConfdState_Rest_Listen_Ssl.
    Ssl []ConfdState_Rest_Listen_Ssl
}

func (listen *ConfdState_Rest_Listen) GetFilter() yfilter.YFilter { return listen.YFilter }

func (listen *ConfdState_Rest_Listen) SetFilter(yf yfilter.YFilter) { listen.YFilter = yf }

func (listen *ConfdState_Rest_Listen) GetGoName(yname string) string {
    if yname == "tcp" { return "Tcp" }
    if yname == "ssl" { return "Ssl" }
    return ""
}

func (listen *ConfdState_Rest_Listen) GetSegmentPath() string {
    return "listen"
}

func (listen *ConfdState_Rest_Listen) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tcp" {
        for _, c := range listen.Tcp {
            if listen.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Rest_Listen_Tcp{}
        listen.Tcp = append(listen.Tcp, child)
        return &listen.Tcp[len(listen.Tcp)-1]
    }
    if childYangName == "ssl" {
        for _, c := range listen.Ssl {
            if listen.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Rest_Listen_Ssl{}
        listen.Ssl = append(listen.Ssl, child)
        return &listen.Ssl[len(listen.Ssl)-1]
    }
    return nil
}

func (listen *ConfdState_Rest_Listen) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range listen.Tcp {
        children[listen.Tcp[i].GetSegmentPath()] = &listen.Tcp[i]
    }
    for i := range listen.Ssl {
        children[listen.Ssl[i].GetSegmentPath()] = &listen.Ssl[i]
    }
    return children
}

func (listen *ConfdState_Rest_Listen) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (listen *ConfdState_Rest_Listen) GetBundleName() string { return "cisco_ios_xe" }

func (listen *ConfdState_Rest_Listen) GetYangName() string { return "listen" }

func (listen *ConfdState_Rest_Listen) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (listen *ConfdState_Rest_Listen) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (listen *ConfdState_Rest_Listen) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (listen *ConfdState_Rest_Listen) SetParent(parent types.Entity) { listen.parent = parent }

func (listen *ConfdState_Rest_Listen) GetParent() types.Entity { return listen.parent }

func (listen *ConfdState_Rest_Listen) GetParentYangName() string { return "rest" }

// ConfdState_Rest_Listen_Tcp
type ConfdState_Rest_Listen_Tcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (tcp *ConfdState_Rest_Listen_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *ConfdState_Rest_Listen_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *ConfdState_Rest_Listen_Tcp) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "port" { return "Port" }
    return ""
}

func (tcp *ConfdState_Rest_Listen_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *ConfdState_Rest_Listen_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *ConfdState_Rest_Listen_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *ConfdState_Rest_Listen_Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = tcp.Ip
    leafs["port"] = tcp.Port
    return leafs
}

func (tcp *ConfdState_Rest_Listen_Tcp) GetBundleName() string { return "cisco_ios_xe" }

func (tcp *ConfdState_Rest_Listen_Tcp) GetYangName() string { return "tcp" }

func (tcp *ConfdState_Rest_Listen_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcp *ConfdState_Rest_Listen_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcp *ConfdState_Rest_Listen_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcp *ConfdState_Rest_Listen_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *ConfdState_Rest_Listen_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *ConfdState_Rest_Listen_Tcp) GetParentYangName() string { return "listen" }

// ConfdState_Rest_Listen_Ssl
type ConfdState_Rest_Listen_Ssl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (ssl *ConfdState_Rest_Listen_Ssl) GetFilter() yfilter.YFilter { return ssl.YFilter }

func (ssl *ConfdState_Rest_Listen_Ssl) SetFilter(yf yfilter.YFilter) { ssl.YFilter = yf }

func (ssl *ConfdState_Rest_Listen_Ssl) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "port" { return "Port" }
    return ""
}

func (ssl *ConfdState_Rest_Listen_Ssl) GetSegmentPath() string {
    return "ssl"
}

func (ssl *ConfdState_Rest_Listen_Ssl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssl *ConfdState_Rest_Listen_Ssl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssl *ConfdState_Rest_Listen_Ssl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = ssl.Ip
    leafs["port"] = ssl.Port
    return leafs
}

func (ssl *ConfdState_Rest_Listen_Ssl) GetBundleName() string { return "cisco_ios_xe" }

func (ssl *ConfdState_Rest_Listen_Ssl) GetYangName() string { return "ssl" }

func (ssl *ConfdState_Rest_Listen_Ssl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ssl *ConfdState_Rest_Listen_Ssl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ssl *ConfdState_Rest_Listen_Ssl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ssl *ConfdState_Rest_Listen_Ssl) SetParent(parent types.Entity) { ssl.parent = parent }

func (ssl *ConfdState_Rest_Listen_Ssl) GetParent() types.Entity { return ssl.parent }

func (ssl *ConfdState_Rest_Listen_Ssl) GetParentYangName() string { return "listen" }

// ConfdState_Snmp
// This type is a presence type.
type ConfdState_Snmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MIBs loaded by the SNMP agent. The type is slice of string.
    Mib []interface{}

    // The local Engine ID specified as a list of colon-specified hexadecimal
    // octets e.g. '4F:4C:41:71'. The type is string with pattern:
    // ([0-9a-fA-F]){2}(:([0-9a-fA-F]){2}){4,31}.
    EngineId interface{}

    // The transport addresses the SNMP agent is listening on.
    Listen ConfdState_Snmp_Listen

    // SNMP version used by the engine.
    Version ConfdState_Snmp_Version
}

func (snmp *ConfdState_Snmp) GetFilter() yfilter.YFilter { return snmp.YFilter }

func (snmp *ConfdState_Snmp) SetFilter(yf yfilter.YFilter) { snmp.YFilter = yf }

func (snmp *ConfdState_Snmp) GetGoName(yname string) string {
    if yname == "mib" { return "Mib" }
    if yname == "engine-id" { return "EngineId" }
    if yname == "listen" { return "Listen" }
    if yname == "version" { return "Version" }
    return ""
}

func (snmp *ConfdState_Snmp) GetSegmentPath() string {
    return "snmp"
}

func (snmp *ConfdState_Snmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "listen" {
        return &snmp.Listen
    }
    if childYangName == "version" {
        return &snmp.Version
    }
    return nil
}

func (snmp *ConfdState_Snmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["listen"] = &snmp.Listen
    children["version"] = &snmp.Version
    return children
}

func (snmp *ConfdState_Snmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mib"] = snmp.Mib
    leafs["engine-id"] = snmp.EngineId
    return leafs
}

func (snmp *ConfdState_Snmp) GetBundleName() string { return "cisco_ios_xe" }

func (snmp *ConfdState_Snmp) GetYangName() string { return "snmp" }

func (snmp *ConfdState_Snmp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmp *ConfdState_Snmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmp *ConfdState_Snmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmp *ConfdState_Snmp) SetParent(parent types.Entity) { snmp.parent = parent }

func (snmp *ConfdState_Snmp) GetParent() types.Entity { return snmp.parent }

func (snmp *ConfdState_Snmp) GetParentYangName() string { return "confd-state" }

// ConfdState_Snmp_Listen
// The transport addresses the SNMP agent is listening on.
type ConfdState_Snmp_Listen struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Snmp_Listen_Udp.
    Udp []ConfdState_Snmp_Listen_Udp
}

func (listen *ConfdState_Snmp_Listen) GetFilter() yfilter.YFilter { return listen.YFilter }

func (listen *ConfdState_Snmp_Listen) SetFilter(yf yfilter.YFilter) { listen.YFilter = yf }

func (listen *ConfdState_Snmp_Listen) GetGoName(yname string) string {
    if yname == "udp" { return "Udp" }
    return ""
}

func (listen *ConfdState_Snmp_Listen) GetSegmentPath() string {
    return "listen"
}

func (listen *ConfdState_Snmp_Listen) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udp" {
        for _, c := range listen.Udp {
            if listen.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Snmp_Listen_Udp{}
        listen.Udp = append(listen.Udp, child)
        return &listen.Udp[len(listen.Udp)-1]
    }
    return nil
}

func (listen *ConfdState_Snmp_Listen) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range listen.Udp {
        children[listen.Udp[i].GetSegmentPath()] = &listen.Udp[i]
    }
    return children
}

func (listen *ConfdState_Snmp_Listen) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (listen *ConfdState_Snmp_Listen) GetBundleName() string { return "cisco_ios_xe" }

func (listen *ConfdState_Snmp_Listen) GetYangName() string { return "listen" }

func (listen *ConfdState_Snmp_Listen) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (listen *ConfdState_Snmp_Listen) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (listen *ConfdState_Snmp_Listen) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (listen *ConfdState_Snmp_Listen) SetParent(parent types.Entity) { listen.parent = parent }

func (listen *ConfdState_Snmp_Listen) GetParent() types.Entity { return listen.parent }

func (listen *ConfdState_Snmp_Listen) GetParentYangName() string { return "snmp" }

// ConfdState_Snmp_Listen_Udp
type ConfdState_Snmp_Listen_Udp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (udp *ConfdState_Snmp_Listen_Udp) GetFilter() yfilter.YFilter { return udp.YFilter }

func (udp *ConfdState_Snmp_Listen_Udp) SetFilter(yf yfilter.YFilter) { udp.YFilter = yf }

func (udp *ConfdState_Snmp_Listen_Udp) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "port" { return "Port" }
    return ""
}

func (udp *ConfdState_Snmp_Listen_Udp) GetSegmentPath() string {
    return "udp"
}

func (udp *ConfdState_Snmp_Listen_Udp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udp *ConfdState_Snmp_Listen_Udp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udp *ConfdState_Snmp_Listen_Udp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = udp.Ip
    leafs["port"] = udp.Port
    return leafs
}

func (udp *ConfdState_Snmp_Listen_Udp) GetBundleName() string { return "cisco_ios_xe" }

func (udp *ConfdState_Snmp_Listen_Udp) GetYangName() string { return "udp" }

func (udp *ConfdState_Snmp_Listen_Udp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (udp *ConfdState_Snmp_Listen_Udp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (udp *ConfdState_Snmp_Listen_Udp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (udp *ConfdState_Snmp_Listen_Udp) SetParent(parent types.Entity) { udp.parent = parent }

func (udp *ConfdState_Snmp_Listen_Udp) GetParent() types.Entity { return udp.parent }

func (udp *ConfdState_Snmp_Listen_Udp) GetParentYangName() string { return "listen" }

// ConfdState_Snmp_Version
// SNMP version used by the engine.
type ConfdState_Snmp_Version struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{}.
    V1 interface{}

    // The type is interface{}.
    V2C interface{}

    // The type is interface{}.
    V3 interface{}
}

func (version *ConfdState_Snmp_Version) GetFilter() yfilter.YFilter { return version.YFilter }

func (version *ConfdState_Snmp_Version) SetFilter(yf yfilter.YFilter) { version.YFilter = yf }

func (version *ConfdState_Snmp_Version) GetGoName(yname string) string {
    if yname == "v1" { return "V1" }
    if yname == "v2c" { return "V2C" }
    if yname == "v3" { return "V3" }
    return ""
}

func (version *ConfdState_Snmp_Version) GetSegmentPath() string {
    return "version"
}

func (version *ConfdState_Snmp_Version) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (version *ConfdState_Snmp_Version) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (version *ConfdState_Snmp_Version) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["v1"] = version.V1
    leafs["v2c"] = version.V2C
    leafs["v3"] = version.V3
    return leafs
}

func (version *ConfdState_Snmp_Version) GetBundleName() string { return "cisco_ios_xe" }

func (version *ConfdState_Snmp_Version) GetYangName() string { return "version" }

func (version *ConfdState_Snmp_Version) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (version *ConfdState_Snmp_Version) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (version *ConfdState_Snmp_Version) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (version *ConfdState_Snmp_Version) SetParent(parent types.Entity) { version.parent = parent }

func (version *ConfdState_Snmp_Version) GetParent() types.Entity { return version.parent }

func (version *ConfdState_Snmp_Version) GetParentYangName() string { return "snmp" }

// ConfdState_Internal
type ConfdState_Internal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Callpoints ConfdState_Internal_Callpoints

    
    Cdb ConfdState_Internal_Cdb
}

func (internal *ConfdState_Internal) GetFilter() yfilter.YFilter { return internal.YFilter }

func (internal *ConfdState_Internal) SetFilter(yf yfilter.YFilter) { internal.YFilter = yf }

func (internal *ConfdState_Internal) GetGoName(yname string) string {
    if yname == "callpoints" { return "Callpoints" }
    if yname == "cdb" { return "Cdb" }
    return ""
}

func (internal *ConfdState_Internal) GetSegmentPath() string {
    return "internal"
}

func (internal *ConfdState_Internal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "callpoints" {
        return &internal.Callpoints
    }
    if childYangName == "cdb" {
        return &internal.Cdb
    }
    return nil
}

func (internal *ConfdState_Internal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["callpoints"] = &internal.Callpoints
    children["cdb"] = &internal.Cdb
    return children
}

func (internal *ConfdState_Internal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (internal *ConfdState_Internal) GetBundleName() string { return "cisco_ios_xe" }

func (internal *ConfdState_Internal) GetYangName() string { return "internal" }

func (internal *ConfdState_Internal) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (internal *ConfdState_Internal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (internal *ConfdState_Internal) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (internal *ConfdState_Internal) SetParent(parent types.Entity) { internal.parent = parent }

func (internal *ConfdState_Internal) GetParent() types.Entity { return internal.parent }

func (internal *ConfdState_Internal) GetParentYangName() string { return "confd-state" }

// ConfdState_Internal_Callpoints
type ConfdState_Internal_Callpoints struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Internal_Callpoints_Callpoint.
    Callpoint []ConfdState_Internal_Callpoints_Callpoint

    // The type is slice of ConfdState_Internal_Callpoints_Validationpoint.
    Validationpoint []ConfdState_Internal_Callpoints_Validationpoint

    // The type is slice of ConfdState_Internal_Callpoints_Actionpoint.
    Actionpoint []ConfdState_Internal_Callpoints_Actionpoint

    // The type is slice of ConfdState_Internal_Callpoints_SnmpInformCallback.
    SnmpInformCallback []ConfdState_Internal_Callpoints_SnmpInformCallback

    // The type is slice of
    // ConfdState_Internal_Callpoints_SnmpNotificationSubscription.
    SnmpNotificationSubscription []ConfdState_Internal_Callpoints_SnmpNotificationSubscription

    // The type is slice of
    // ConfdState_Internal_Callpoints_ErrorFormattingCallback.
    ErrorFormattingCallback []ConfdState_Internal_Callpoints_ErrorFormattingCallback

    // The type is slice of ConfdState_Internal_Callpoints_Typepoint.
    Typepoint []ConfdState_Internal_Callpoints_Typepoint

    // The type is slice of
    // ConfdState_Internal_Callpoints_NotificationStreamReplay.
    NotificationStreamReplay []ConfdState_Internal_Callpoints_NotificationStreamReplay

    
    AuthenticationCallback ConfdState_Internal_Callpoints_AuthenticationCallback

    
    AuthorizationCallbacks ConfdState_Internal_Callpoints_AuthorizationCallbacks
}

func (callpoints *ConfdState_Internal_Callpoints) GetFilter() yfilter.YFilter { return callpoints.YFilter }

func (callpoints *ConfdState_Internal_Callpoints) SetFilter(yf yfilter.YFilter) { callpoints.YFilter = yf }

func (callpoints *ConfdState_Internal_Callpoints) GetGoName(yname string) string {
    if yname == "callpoint" { return "Callpoint" }
    if yname == "validationpoint" { return "Validationpoint" }
    if yname == "actionpoint" { return "Actionpoint" }
    if yname == "snmp-inform-callback" { return "SnmpInformCallback" }
    if yname == "snmp-notification-subscription" { return "SnmpNotificationSubscription" }
    if yname == "error-formatting-callback" { return "ErrorFormattingCallback" }
    if yname == "typepoint" { return "Typepoint" }
    if yname == "notification-stream-replay" { return "NotificationStreamReplay" }
    if yname == "authentication-callback" { return "AuthenticationCallback" }
    if yname == "authorization-callbacks" { return "AuthorizationCallbacks" }
    return ""
}

func (callpoints *ConfdState_Internal_Callpoints) GetSegmentPath() string {
    return "callpoints"
}

func (callpoints *ConfdState_Internal_Callpoints) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "callpoint" {
        for _, c := range callpoints.Callpoint {
            if callpoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_Callpoint{}
        callpoints.Callpoint = append(callpoints.Callpoint, child)
        return &callpoints.Callpoint[len(callpoints.Callpoint)-1]
    }
    if childYangName == "validationpoint" {
        for _, c := range callpoints.Validationpoint {
            if callpoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_Validationpoint{}
        callpoints.Validationpoint = append(callpoints.Validationpoint, child)
        return &callpoints.Validationpoint[len(callpoints.Validationpoint)-1]
    }
    if childYangName == "actionpoint" {
        for _, c := range callpoints.Actionpoint {
            if callpoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_Actionpoint{}
        callpoints.Actionpoint = append(callpoints.Actionpoint, child)
        return &callpoints.Actionpoint[len(callpoints.Actionpoint)-1]
    }
    if childYangName == "snmp-inform-callback" {
        for _, c := range callpoints.SnmpInformCallback {
            if callpoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_SnmpInformCallback{}
        callpoints.SnmpInformCallback = append(callpoints.SnmpInformCallback, child)
        return &callpoints.SnmpInformCallback[len(callpoints.SnmpInformCallback)-1]
    }
    if childYangName == "snmp-notification-subscription" {
        for _, c := range callpoints.SnmpNotificationSubscription {
            if callpoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_SnmpNotificationSubscription{}
        callpoints.SnmpNotificationSubscription = append(callpoints.SnmpNotificationSubscription, child)
        return &callpoints.SnmpNotificationSubscription[len(callpoints.SnmpNotificationSubscription)-1]
    }
    if childYangName == "error-formatting-callback" {
        for _, c := range callpoints.ErrorFormattingCallback {
            if callpoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_ErrorFormattingCallback{}
        callpoints.ErrorFormattingCallback = append(callpoints.ErrorFormattingCallback, child)
        return &callpoints.ErrorFormattingCallback[len(callpoints.ErrorFormattingCallback)-1]
    }
    if childYangName == "typepoint" {
        for _, c := range callpoints.Typepoint {
            if callpoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_Typepoint{}
        callpoints.Typepoint = append(callpoints.Typepoint, child)
        return &callpoints.Typepoint[len(callpoints.Typepoint)-1]
    }
    if childYangName == "notification-stream-replay" {
        for _, c := range callpoints.NotificationStreamReplay {
            if callpoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_NotificationStreamReplay{}
        callpoints.NotificationStreamReplay = append(callpoints.NotificationStreamReplay, child)
        return &callpoints.NotificationStreamReplay[len(callpoints.NotificationStreamReplay)-1]
    }
    if childYangName == "authentication-callback" {
        return &callpoints.AuthenticationCallback
    }
    if childYangName == "authorization-callbacks" {
        return &callpoints.AuthorizationCallbacks
    }
    return nil
}

func (callpoints *ConfdState_Internal_Callpoints) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range callpoints.Callpoint {
        children[callpoints.Callpoint[i].GetSegmentPath()] = &callpoints.Callpoint[i]
    }
    for i := range callpoints.Validationpoint {
        children[callpoints.Validationpoint[i].GetSegmentPath()] = &callpoints.Validationpoint[i]
    }
    for i := range callpoints.Actionpoint {
        children[callpoints.Actionpoint[i].GetSegmentPath()] = &callpoints.Actionpoint[i]
    }
    for i := range callpoints.SnmpInformCallback {
        children[callpoints.SnmpInformCallback[i].GetSegmentPath()] = &callpoints.SnmpInformCallback[i]
    }
    for i := range callpoints.SnmpNotificationSubscription {
        children[callpoints.SnmpNotificationSubscription[i].GetSegmentPath()] = &callpoints.SnmpNotificationSubscription[i]
    }
    for i := range callpoints.ErrorFormattingCallback {
        children[callpoints.ErrorFormattingCallback[i].GetSegmentPath()] = &callpoints.ErrorFormattingCallback[i]
    }
    for i := range callpoints.Typepoint {
        children[callpoints.Typepoint[i].GetSegmentPath()] = &callpoints.Typepoint[i]
    }
    for i := range callpoints.NotificationStreamReplay {
        children[callpoints.NotificationStreamReplay[i].GetSegmentPath()] = &callpoints.NotificationStreamReplay[i]
    }
    children["authentication-callback"] = &callpoints.AuthenticationCallback
    children["authorization-callbacks"] = &callpoints.AuthorizationCallbacks
    return children
}

func (callpoints *ConfdState_Internal_Callpoints) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (callpoints *ConfdState_Internal_Callpoints) GetBundleName() string { return "cisco_ios_xe" }

func (callpoints *ConfdState_Internal_Callpoints) GetYangName() string { return "callpoints" }

func (callpoints *ConfdState_Internal_Callpoints) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (callpoints *ConfdState_Internal_Callpoints) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (callpoints *ConfdState_Internal_Callpoints) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (callpoints *ConfdState_Internal_Callpoints) SetParent(parent types.Entity) { callpoints.parent = parent }

func (callpoints *ConfdState_Internal_Callpoints) GetParent() types.Entity { return callpoints.parent }

func (callpoints *ConfdState_Internal_Callpoints) GetParentYangName() string { return "internal" }

// ConfdState_Internal_Callpoints_Callpoint
type ConfdState_Internal_Callpoints_Callpoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Callpoint id. The type is string.
    Id interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_Callpoint_Daemon

    // The type is slice of ConfdState_Internal_Callpoints_Callpoint_Range.
    Range []ConfdState_Internal_Callpoints_Callpoint_Range
}

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetFilter() yfilter.YFilter { return callpoint.YFilter }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) SetFilter(yf yfilter.YFilter) { callpoint.YFilter = yf }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetSegmentPath() string {
    return "callpoint" + "[id='" + fmt.Sprintf("%v", callpoint.Id) + "']"
}

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &callpoint.Daemon
    }
    if childYangName == "range" {
        for _, c := range callpoint.Range {
            if callpoint.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_Callpoint_Range{}
        callpoint.Range = append(callpoint.Range, child)
        return &callpoint.Range[len(callpoint.Range)-1]
    }
    return nil
}

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &callpoint.Daemon
    for i := range callpoint.Range {
        children[callpoint.Range[i].GetSegmentPath()] = &callpoint.Range[i]
    }
    return children
}

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = callpoint.Id
    leafs["path"] = callpoint.Path
    leafs["file"] = callpoint.File
    leafs["error"] = callpoint.Error
    return leafs
}

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetBundleName() string { return "cisco_ios_xe" }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetYangName() string { return "callpoint" }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) SetParent(parent types.Entity) { callpoint.parent = parent }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetParent() types.Entity { return callpoint.parent }

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_Callpoint_Daemon
type ConfdState_Internal_Callpoints_Callpoint_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetParentYangName() string { return "callpoint" }

// ConfdState_Internal_Callpoints_Callpoint_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_Callpoint_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Callpoint_Daemon_Error_PENDING ConfdState_Internal_Callpoints_Callpoint_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_Callpoint_Range
type ConfdState_Internal_Callpoints_Callpoint_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_Callpoint_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetParentYangName() string { return "callpoint" }

// ConfdState_Internal_Callpoints_Callpoint_Range_Daemon
type ConfdState_Internal_Callpoints_Callpoint_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_Callpoint_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_Callpoint_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Callpoint_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_Callpoint_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_Callpoint_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_Callpoint_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_Callpoint_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_Callpoint_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_Callpoint_Error_UNKNOWN ConfdState_Internal_Callpoints_Callpoint_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_Validationpoint
type ConfdState_Internal_Callpoints_Validationpoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Callpoint id. The type is string.
    Id interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_Validationpoint_Daemon

    // The type is slice of ConfdState_Internal_Callpoints_Validationpoint_Range.
    Range []ConfdState_Internal_Callpoints_Validationpoint_Range
}

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetFilter() yfilter.YFilter { return validationpoint.YFilter }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) SetFilter(yf yfilter.YFilter) { validationpoint.YFilter = yf }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetSegmentPath() string {
    return "validationpoint" + "[id='" + fmt.Sprintf("%v", validationpoint.Id) + "']"
}

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &validationpoint.Daemon
    }
    if childYangName == "range" {
        for _, c := range validationpoint.Range {
            if validationpoint.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_Validationpoint_Range{}
        validationpoint.Range = append(validationpoint.Range, child)
        return &validationpoint.Range[len(validationpoint.Range)-1]
    }
    return nil
}

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &validationpoint.Daemon
    for i := range validationpoint.Range {
        children[validationpoint.Range[i].GetSegmentPath()] = &validationpoint.Range[i]
    }
    return children
}

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = validationpoint.Id
    leafs["path"] = validationpoint.Path
    leafs["file"] = validationpoint.File
    leafs["error"] = validationpoint.Error
    return leafs
}

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetBundleName() string { return "cisco_ios_xe" }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetYangName() string { return "validationpoint" }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) SetParent(parent types.Entity) { validationpoint.parent = parent }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetParent() types.Entity { return validationpoint.parent }

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_Validationpoint_Daemon
type ConfdState_Internal_Callpoints_Validationpoint_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetParentYangName() string { return "validationpoint" }

// ConfdState_Internal_Callpoints_Validationpoint_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_Validationpoint_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Validationpoint_Daemon_Error_PENDING ConfdState_Internal_Callpoints_Validationpoint_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_Validationpoint_Range
type ConfdState_Internal_Callpoints_Validationpoint_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetParentYangName() string { return "validationpoint" }

// ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon
type ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_Validationpoint_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_Validationpoint_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_Validationpoint_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_Validationpoint_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_Validationpoint_Error_UNKNOWN ConfdState_Internal_Callpoints_Validationpoint_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_Actionpoint
type ConfdState_Internal_Callpoints_Actionpoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Callpoint id. The type is string.
    Id interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_Actionpoint_Daemon

    // The type is slice of ConfdState_Internal_Callpoints_Actionpoint_Range.
    Range []ConfdState_Internal_Callpoints_Actionpoint_Range
}

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetFilter() yfilter.YFilter { return actionpoint.YFilter }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) SetFilter(yf yfilter.YFilter) { actionpoint.YFilter = yf }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetSegmentPath() string {
    return "actionpoint" + "[id='" + fmt.Sprintf("%v", actionpoint.Id) + "']"
}

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &actionpoint.Daemon
    }
    if childYangName == "range" {
        for _, c := range actionpoint.Range {
            if actionpoint.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_Actionpoint_Range{}
        actionpoint.Range = append(actionpoint.Range, child)
        return &actionpoint.Range[len(actionpoint.Range)-1]
    }
    return nil
}

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &actionpoint.Daemon
    for i := range actionpoint.Range {
        children[actionpoint.Range[i].GetSegmentPath()] = &actionpoint.Range[i]
    }
    return children
}

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = actionpoint.Id
    leafs["path"] = actionpoint.Path
    leafs["file"] = actionpoint.File
    leafs["error"] = actionpoint.Error
    return leafs
}

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetBundleName() string { return "cisco_ios_xe" }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetYangName() string { return "actionpoint" }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) SetParent(parent types.Entity) { actionpoint.parent = parent }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetParent() types.Entity { return actionpoint.parent }

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_Actionpoint_Daemon
type ConfdState_Internal_Callpoints_Actionpoint_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetParentYangName() string { return "actionpoint" }

// ConfdState_Internal_Callpoints_Actionpoint_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_Actionpoint_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Actionpoint_Daemon_Error_PENDING ConfdState_Internal_Callpoints_Actionpoint_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_Actionpoint_Range
type ConfdState_Internal_Callpoints_Actionpoint_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetParentYangName() string { return "actionpoint" }

// ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon
type ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_Actionpoint_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_Actionpoint_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_Actionpoint_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_Actionpoint_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_Actionpoint_Error_UNKNOWN ConfdState_Internal_Callpoints_Actionpoint_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_SnmpInformCallback
type ConfdState_Internal_Callpoints_SnmpInformCallback struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Callpoint id. The type is string.
    Id interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_SnmpInformCallback_Range.
    Range []ConfdState_Internal_Callpoints_SnmpInformCallback_Range
}

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetFilter() yfilter.YFilter { return snmpInformCallback.YFilter }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) SetFilter(yf yfilter.YFilter) { snmpInformCallback.YFilter = yf }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetSegmentPath() string {
    return "snmp-inform-callback" + "[id='" + fmt.Sprintf("%v", snmpInformCallback.Id) + "']"
}

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &snmpInformCallback.Daemon
    }
    if childYangName == "range" {
        for _, c := range snmpInformCallback.Range {
            if snmpInformCallback.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_SnmpInformCallback_Range{}
        snmpInformCallback.Range = append(snmpInformCallback.Range, child)
        return &snmpInformCallback.Range[len(snmpInformCallback.Range)-1]
    }
    return nil
}

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &snmpInformCallback.Daemon
    for i := range snmpInformCallback.Range {
        children[snmpInformCallback.Range[i].GetSegmentPath()] = &snmpInformCallback.Range[i]
    }
    return children
}

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = snmpInformCallback.Id
    leafs["path"] = snmpInformCallback.Path
    leafs["file"] = snmpInformCallback.File
    leafs["error"] = snmpInformCallback.Error
    return leafs
}

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetBundleName() string { return "cisco_ios_xe" }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetYangName() string { return "snmp-inform-callback" }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) SetParent(parent types.Entity) { snmpInformCallback.parent = parent }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetParent() types.Entity { return snmpInformCallback.parent }

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon
type ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetParentYangName() string { return "snmp-inform-callback" }

// ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon_Error_PENDING ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_SnmpInformCallback_Range
type ConfdState_Internal_Callpoints_SnmpInformCallback_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetParentYangName() string { return "snmp-inform-callback" }

// ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon
type ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_SnmpInformCallback_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_SnmpInformCallback_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_SnmpInformCallback_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_SnmpInformCallback_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_SnmpInformCallback_Error_UNKNOWN ConfdState_Internal_Callpoints_SnmpInformCallback_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Callpoint id. The type is string.
    Id interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range.
    Range []ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range
}

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetFilter() yfilter.YFilter { return snmpNotificationSubscription.YFilter }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) SetFilter(yf yfilter.YFilter) { snmpNotificationSubscription.YFilter = yf }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetSegmentPath() string {
    return "snmp-notification-subscription" + "[id='" + fmt.Sprintf("%v", snmpNotificationSubscription.Id) + "']"
}

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &snmpNotificationSubscription.Daemon
    }
    if childYangName == "range" {
        for _, c := range snmpNotificationSubscription.Range {
            if snmpNotificationSubscription.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range{}
        snmpNotificationSubscription.Range = append(snmpNotificationSubscription.Range, child)
        return &snmpNotificationSubscription.Range[len(snmpNotificationSubscription.Range)-1]
    }
    return nil
}

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &snmpNotificationSubscription.Daemon
    for i := range snmpNotificationSubscription.Range {
        children[snmpNotificationSubscription.Range[i].GetSegmentPath()] = &snmpNotificationSubscription.Range[i]
    }
    return children
}

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = snmpNotificationSubscription.Id
    leafs["path"] = snmpNotificationSubscription.Path
    leafs["file"] = snmpNotificationSubscription.File
    leafs["error"] = snmpNotificationSubscription.Error
    return leafs
}

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetBundleName() string { return "cisco_ios_xe" }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetYangName() string { return "snmp-notification-subscription" }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) SetParent(parent types.Entity) { snmpNotificationSubscription.parent = parent }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetParent() types.Entity { return snmpNotificationSubscription.parent }

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetParentYangName() string { return "snmp-notification-subscription" }

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon_Error_PENDING ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetParentYangName() string { return "snmp-notification-subscription" }

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error_UNKNOWN ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_ErrorFormattingCallback
type ConfdState_Internal_Callpoints_ErrorFormattingCallback struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Callpoint id. The type is string.
    Id interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range.
    Range []ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range
}

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetFilter() yfilter.YFilter { return errorFormattingCallback.YFilter }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) SetFilter(yf yfilter.YFilter) { errorFormattingCallback.YFilter = yf }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetSegmentPath() string {
    return "error-formatting-callback" + "[id='" + fmt.Sprintf("%v", errorFormattingCallback.Id) + "']"
}

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &errorFormattingCallback.Daemon
    }
    if childYangName == "range" {
        for _, c := range errorFormattingCallback.Range {
            if errorFormattingCallback.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range{}
        errorFormattingCallback.Range = append(errorFormattingCallback.Range, child)
        return &errorFormattingCallback.Range[len(errorFormattingCallback.Range)-1]
    }
    return nil
}

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &errorFormattingCallback.Daemon
    for i := range errorFormattingCallback.Range {
        children[errorFormattingCallback.Range[i].GetSegmentPath()] = &errorFormattingCallback.Range[i]
    }
    return children
}

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = errorFormattingCallback.Id
    leafs["path"] = errorFormattingCallback.Path
    leafs["file"] = errorFormattingCallback.File
    leafs["error"] = errorFormattingCallback.Error
    return leafs
}

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetBundleName() string { return "cisco_ios_xe" }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetYangName() string { return "error-formatting-callback" }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) SetParent(parent types.Entity) { errorFormattingCallback.parent = parent }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetParent() types.Entity { return errorFormattingCallback.parent }

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetParentYangName() string { return "error-formatting-callback" }

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon_Error_PENDING ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetParentYangName() string { return "error-formatting-callback" }

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error_UNKNOWN ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_Typepoint
type ConfdState_Internal_Callpoints_Typepoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Callpoint id. The type is string.
    Id interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_Typepoint_Daemon

    // The type is slice of ConfdState_Internal_Callpoints_Typepoint_Range.
    Range []ConfdState_Internal_Callpoints_Typepoint_Range
}

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetFilter() yfilter.YFilter { return typepoint.YFilter }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) SetFilter(yf yfilter.YFilter) { typepoint.YFilter = yf }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetSegmentPath() string {
    return "typepoint" + "[id='" + fmt.Sprintf("%v", typepoint.Id) + "']"
}

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &typepoint.Daemon
    }
    if childYangName == "range" {
        for _, c := range typepoint.Range {
            if typepoint.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_Typepoint_Range{}
        typepoint.Range = append(typepoint.Range, child)
        return &typepoint.Range[len(typepoint.Range)-1]
    }
    return nil
}

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &typepoint.Daemon
    for i := range typepoint.Range {
        children[typepoint.Range[i].GetSegmentPath()] = &typepoint.Range[i]
    }
    return children
}

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = typepoint.Id
    leafs["path"] = typepoint.Path
    leafs["file"] = typepoint.File
    leafs["error"] = typepoint.Error
    return leafs
}

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetBundleName() string { return "cisco_ios_xe" }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetYangName() string { return "typepoint" }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) SetParent(parent types.Entity) { typepoint.parent = parent }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetParent() types.Entity { return typepoint.parent }

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_Typepoint_Daemon
type ConfdState_Internal_Callpoints_Typepoint_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetParentYangName() string { return "typepoint" }

// ConfdState_Internal_Callpoints_Typepoint_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_Typepoint_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Typepoint_Daemon_Error_PENDING ConfdState_Internal_Callpoints_Typepoint_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_Typepoint_Range
type ConfdState_Internal_Callpoints_Typepoint_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_Typepoint_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetParentYangName() string { return "typepoint" }

// ConfdState_Internal_Callpoints_Typepoint_Range_Daemon
type ConfdState_Internal_Callpoints_Typepoint_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_Typepoint_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_Typepoint_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Typepoint_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_Typepoint_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_Typepoint_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_Typepoint_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_Typepoint_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_Typepoint_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_Typepoint_Error_UNKNOWN ConfdState_Internal_Callpoints_Typepoint_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_NotificationStreamReplay
type ConfdState_Internal_Callpoints_NotificationStreamReplay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the notification stream. The type is
    // string.
    Name interface{}

    // The type is ReplaySupport.
    ReplaySupport interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_NotificationStreamReplay_Range.
    Range []ConfdState_Internal_Callpoints_NotificationStreamReplay_Range
}

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetFilter() yfilter.YFilter { return notificationStreamReplay.YFilter }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) SetFilter(yf yfilter.YFilter) { notificationStreamReplay.YFilter = yf }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "replay-support" { return "ReplaySupport" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetSegmentPath() string {
    return "notification-stream-replay" + "[name='" + fmt.Sprintf("%v", notificationStreamReplay.Name) + "']"
}

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &notificationStreamReplay.Daemon
    }
    if childYangName == "range" {
        for _, c := range notificationStreamReplay.Range {
            if notificationStreamReplay.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_NotificationStreamReplay_Range{}
        notificationStreamReplay.Range = append(notificationStreamReplay.Range, child)
        return &notificationStreamReplay.Range[len(notificationStreamReplay.Range)-1]
    }
    return nil
}

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &notificationStreamReplay.Daemon
    for i := range notificationStreamReplay.Range {
        children[notificationStreamReplay.Range[i].GetSegmentPath()] = &notificationStreamReplay.Range[i]
    }
    return children
}

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = notificationStreamReplay.Name
    leafs["replay-support"] = notificationStreamReplay.ReplaySupport
    leafs["path"] = notificationStreamReplay.Path
    leafs["file"] = notificationStreamReplay.File
    leafs["error"] = notificationStreamReplay.Error
    return leafs
}

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetBundleName() string { return "cisco_ios_xe" }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetYangName() string { return "notification-stream-replay" }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) SetParent(parent types.Entity) { notificationStreamReplay.parent = parent }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetParent() types.Entity { return notificationStreamReplay.parent }

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetParentYangName() string { return "notification-stream-replay" }

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon_Error_PENDING ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Range
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetParentYangName() string { return "notification-stream-replay" }

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_NotificationStreamReplay_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_NotificationStreamReplay_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_NotificationStreamReplay_Error_UNKNOWN ConfdState_Internal_Callpoints_NotificationStreamReplay_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_NotificationStreamReplay_ReplaySupport
type ConfdState_Internal_Callpoints_NotificationStreamReplay_ReplaySupport string

const (
    ConfdState_Internal_Callpoints_NotificationStreamReplay_ReplaySupport_none ConfdState_Internal_Callpoints_NotificationStreamReplay_ReplaySupport = "none"

    ConfdState_Internal_Callpoints_NotificationStreamReplay_ReplaySupport_builtin ConfdState_Internal_Callpoints_NotificationStreamReplay_ReplaySupport = "builtin"

    ConfdState_Internal_Callpoints_NotificationStreamReplay_ReplaySupport_external ConfdState_Internal_Callpoints_NotificationStreamReplay_ReplaySupport = "external"
)

// ConfdState_Internal_Callpoints_AuthenticationCallback
// This type is a presence type.
type ConfdState_Internal_Callpoints_AuthenticationCallback struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is bool.
    Enabled interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_AuthenticationCallback_Range.
    Range []ConfdState_Internal_Callpoints_AuthenticationCallback_Range
}

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetFilter() yfilter.YFilter { return authenticationCallback.YFilter }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) SetFilter(yf yfilter.YFilter) { authenticationCallback.YFilter = yf }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetSegmentPath() string {
    return "authentication-callback"
}

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &authenticationCallback.Daemon
    }
    if childYangName == "range" {
        for _, c := range authenticationCallback.Range {
            if authenticationCallback.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_AuthenticationCallback_Range{}
        authenticationCallback.Range = append(authenticationCallback.Range, child)
        return &authenticationCallback.Range[len(authenticationCallback.Range)-1]
    }
    return nil
}

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &authenticationCallback.Daemon
    for i := range authenticationCallback.Range {
        children[authenticationCallback.Range[i].GetSegmentPath()] = &authenticationCallback.Range[i]
    }
    return children
}

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = authenticationCallback.Enabled
    leafs["path"] = authenticationCallback.Path
    leafs["file"] = authenticationCallback.File
    leafs["error"] = authenticationCallback.Error
    return leafs
}

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetBundleName() string { return "cisco_ios_xe" }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetYangName() string { return "authentication-callback" }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) SetParent(parent types.Entity) { authenticationCallback.parent = parent }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetParent() types.Entity { return authenticationCallback.parent }

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon
type ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetParentYangName() string { return "authentication-callback" }

// ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon_Error_PENDING ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_AuthenticationCallback_Range
type ConfdState_Internal_Callpoints_AuthenticationCallback_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetParentYangName() string { return "authentication-callback" }

// ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon
type ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_AuthenticationCallback_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_AuthenticationCallback_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_AuthenticationCallback_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_AuthenticationCallback_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_AuthenticationCallback_Error_UNKNOWN ConfdState_Internal_Callpoints_AuthenticationCallback_Error = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_AuthorizationCallbacks
// This type is a presence type.
type ConfdState_Internal_Callpoints_AuthorizationCallbacks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is bool.
    Enabled interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range.
    Range []ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range
}

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetFilter() yfilter.YFilter { return authorizationCallbacks.YFilter }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) SetFilter(yf yfilter.YFilter) { authorizationCallbacks.YFilter = yf }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetGoName(yname string) string {
    if yname == "enabled" { return "Enabled" }
    if yname == "path" { return "Path" }
    if yname == "file" { return "File" }
    if yname == "error" { return "Error" }
    if yname == "daemon" { return "Daemon" }
    if yname == "range" { return "Range" }
    return ""
}

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetSegmentPath() string {
    return "authorization-callbacks"
}

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &authorizationCallbacks.Daemon
    }
    if childYangName == "range" {
        for _, c := range authorizationCallbacks.Range {
            if authorizationCallbacks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range{}
        authorizationCallbacks.Range = append(authorizationCallbacks.Range, child)
        return &authorizationCallbacks.Range[len(authorizationCallbacks.Range)-1]
    }
    return nil
}

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &authorizationCallbacks.Daemon
    for i := range authorizationCallbacks.Range {
        children[authorizationCallbacks.Range[i].GetSegmentPath()] = &authorizationCallbacks.Range[i]
    }
    return children
}

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enabled"] = authorizationCallbacks.Enabled
    leafs["path"] = authorizationCallbacks.Path
    leafs["file"] = authorizationCallbacks.File
    leafs["error"] = authorizationCallbacks.Error
    return leafs
}

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetBundleName() string { return "cisco_ios_xe" }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetYangName() string { return "authorization-callbacks" }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) SetParent(parent types.Entity) { authorizationCallbacks.parent = parent }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetParent() types.Entity { return authorizationCallbacks.parent }

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetParentYangName() string { return "callpoints" }

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetParentYangName() string { return "authorization-callbacks" }

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon_Error_PENDING ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The space-separated set of keys that defines the lower endpoint of the
    // range for a non-default registration. The type is string.
    Lower interface{}

    // The space-separated set of keys that defines the upper endpoint of the
    // range for a non-default registration. The type is string.
    Upper interface{}

    // If this leaf exists, this is a default registration - in this case 'lower'
    // and 'upper' do not exist. The type is interface{}.
    Default interface{}

    
    Daemon ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon
}

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetGoName(yname string) string {
    if yname == "lower" { return "Lower" }
    if yname == "upper" { return "Upper" }
    if yname == "default" { return "Default" }
    if yname == "daemon" { return "Daemon" }
    return ""
}

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetSegmentPath() string {
    return "range"
}

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "daemon" {
        return &self.Daemon
    }
    return nil
}

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["daemon"] = &self.Daemon
    return children
}

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower"] = self.Lower
    leafs["upper"] = self.Upper
    leafs["default"] = self.Default
    return leafs
}

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetBundleName() string { return "cisco_ios_xe" }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetYangName() string { return "range" }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetParent() types.Entity { return self.parent }

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetParentYangName() string { return "authorization-callbacks" }

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetFilter() yfilter.YFilter { return daemon.YFilter }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) SetFilter(yf yfilter.YFilter) { daemon.YFilter = yf }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "error" { return "Error" }
    return ""
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetSegmentPath() string {
    return "daemon"
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = daemon.Id
    leafs["name"] = daemon.Name
    leafs["error"] = daemon.Error
    return leafs
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetBundleName() string { return "cisco_ios_xe" }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetYangName() string { return "daemon" }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) SetParent(parent types.Entity) { daemon.parent = parent }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetParent() types.Entity { return daemon.parent }

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetParentYangName() string { return "range" }

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon_Error represents with the daemon registration.
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon_Error string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon_Error_PENDING ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon_Error = "PENDING"
)

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error represents with the callpoint registration.
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error_NOT_REGISTERED ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error_UNKNOWN ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error = "UNKNOWN"
)

// ConfdState_Internal_Cdb
type ConfdState_Internal_Cdb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Internal_Cdb_Datastore.
    Datastore []ConfdState_Internal_Cdb_Datastore

    // The type is slice of ConfdState_Internal_Cdb_Client.
    Client []ConfdState_Internal_Cdb_Client
}

func (cdb *ConfdState_Internal_Cdb) GetFilter() yfilter.YFilter { return cdb.YFilter }

func (cdb *ConfdState_Internal_Cdb) SetFilter(yf yfilter.YFilter) { cdb.YFilter = yf }

func (cdb *ConfdState_Internal_Cdb) GetGoName(yname string) string {
    if yname == "datastore" { return "Datastore" }
    if yname == "client" { return "Client" }
    return ""
}

func (cdb *ConfdState_Internal_Cdb) GetSegmentPath() string {
    return "cdb"
}

func (cdb *ConfdState_Internal_Cdb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "datastore" {
        for _, c := range cdb.Datastore {
            if cdb.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Cdb_Datastore{}
        cdb.Datastore = append(cdb.Datastore, child)
        return &cdb.Datastore[len(cdb.Datastore)-1]
    }
    if childYangName == "client" {
        for _, c := range cdb.Client {
            if cdb.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Cdb_Client{}
        cdb.Client = append(cdb.Client, child)
        return &cdb.Client[len(cdb.Client)-1]
    }
    return nil
}

func (cdb *ConfdState_Internal_Cdb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cdb.Datastore {
        children[cdb.Datastore[i].GetSegmentPath()] = &cdb.Datastore[i]
    }
    for i := range cdb.Client {
        children[cdb.Client[i].GetSegmentPath()] = &cdb.Client[i]
    }
    return children
}

func (cdb *ConfdState_Internal_Cdb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cdb *ConfdState_Internal_Cdb) GetBundleName() string { return "cisco_ios_xe" }

func (cdb *ConfdState_Internal_Cdb) GetYangName() string { return "cdb" }

func (cdb *ConfdState_Internal_Cdb) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cdb *ConfdState_Internal_Cdb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cdb *ConfdState_Internal_Cdb) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cdb *ConfdState_Internal_Cdb) SetParent(parent types.Entity) { cdb.parent = parent }

func (cdb *ConfdState_Internal_Cdb) GetParent() types.Entity { return cdb.parent }

func (cdb *ConfdState_Internal_Cdb) GetParentYangName() string { return "internal" }

// ConfdState_Internal_Cdb_Datastore
type ConfdState_Internal_Cdb_Datastore struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is DatastoreName.
    Name interface{}

    // The id of the last committed transaction for the 'running' datastore, or
    // the last update for the 'operational' datastore. For the 'operational'
    // datastore, it is only present when HA is enabled. The type is string.
    TransactionId interface{}

    // Number of pending write requests for the 'operational' datastore on a HA
    // slave that is in the process of syncronizing with the master. The type is
    // interface{} with range: 0..4294967295.
    WriteQueue interface{}

    // The pathname of the file that is used for persistent storage for the
    // datastore. Not present for 'running' when 'startup' exists. The type is
    // string.
    Filename interface{}

    // The size of the file that is used for persistent storage for the datastore.
    // Only present if 'filename' is present. The type is interface{} with range:
    // 0..18446744073709551615.
    DiskSize interface{}

    // The size of the in-memory representation of the datastore. The type is
    // interface{} with range: 0..18446744073709551615.
    RamSize interface{}

    // The number of read locks on the datastore. Not present for the
    // 'operational' data store. The type is interface{} with range:
    // 0..4294967295.
    ReadLocks interface{}

    // Indicates whether a write lock is in effect for the datastore. Not present
    // for the 'operational' datastore. The type is bool.
    WriteLockSet interface{}

    // Indicates whether a subscription lock is in effect for the 'operational'
    // datastore. The type is bool.
    SubscriptionLockSet interface{}

    // Indicates whether synchronous replication from HA master to HA slave is in
    // progress for the datastore. Not present for the 'operational' datastore.
    // The type is bool.
    WaitingForReplicationSync interface{}

    // Information pertaining to subscription notifications that have been
    // delivered to, but not yet acknowledged by, subscribing clients. Not present
    // for the 'startup' datastore.
    PendingSubscriptionSync ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync

    // Queues of notifications that have been generated but not yet delivered to
    // subscribing clients. Not present for the 'startup' datastore. The type is
    // slice of ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue.
    PendingNotificationQueue []ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue
}

func (datastore *ConfdState_Internal_Cdb_Datastore) GetFilter() yfilter.YFilter { return datastore.YFilter }

func (datastore *ConfdState_Internal_Cdb_Datastore) SetFilter(yf yfilter.YFilter) { datastore.YFilter = yf }

func (datastore *ConfdState_Internal_Cdb_Datastore) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "transaction-id" { return "TransactionId" }
    if yname == "write-queue" { return "WriteQueue" }
    if yname == "filename" { return "Filename" }
    if yname == "disk-size" { return "DiskSize" }
    if yname == "ram-size" { return "RamSize" }
    if yname == "read-locks" { return "ReadLocks" }
    if yname == "write-lock-set" { return "WriteLockSet" }
    if yname == "subscription-lock-set" { return "SubscriptionLockSet" }
    if yname == "waiting-for-replication-sync" { return "WaitingForReplicationSync" }
    if yname == "pending-subscription-sync" { return "PendingSubscriptionSync" }
    if yname == "pending-notification-queue" { return "PendingNotificationQueue" }
    return ""
}

func (datastore *ConfdState_Internal_Cdb_Datastore) GetSegmentPath() string {
    return "datastore" + "[name='" + fmt.Sprintf("%v", datastore.Name) + "']"
}

func (datastore *ConfdState_Internal_Cdb_Datastore) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pending-subscription-sync" {
        return &datastore.PendingSubscriptionSync
    }
    if childYangName == "pending-notification-queue" {
        for _, c := range datastore.PendingNotificationQueue {
            if datastore.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue{}
        datastore.PendingNotificationQueue = append(datastore.PendingNotificationQueue, child)
        return &datastore.PendingNotificationQueue[len(datastore.PendingNotificationQueue)-1]
    }
    return nil
}

func (datastore *ConfdState_Internal_Cdb_Datastore) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pending-subscription-sync"] = &datastore.PendingSubscriptionSync
    for i := range datastore.PendingNotificationQueue {
        children[datastore.PendingNotificationQueue[i].GetSegmentPath()] = &datastore.PendingNotificationQueue[i]
    }
    return children
}

func (datastore *ConfdState_Internal_Cdb_Datastore) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = datastore.Name
    leafs["transaction-id"] = datastore.TransactionId
    leafs["write-queue"] = datastore.WriteQueue
    leafs["filename"] = datastore.Filename
    leafs["disk-size"] = datastore.DiskSize
    leafs["ram-size"] = datastore.RamSize
    leafs["read-locks"] = datastore.ReadLocks
    leafs["write-lock-set"] = datastore.WriteLockSet
    leafs["subscription-lock-set"] = datastore.SubscriptionLockSet
    leafs["waiting-for-replication-sync"] = datastore.WaitingForReplicationSync
    return leafs
}

func (datastore *ConfdState_Internal_Cdb_Datastore) GetBundleName() string { return "cisco_ios_xe" }

func (datastore *ConfdState_Internal_Cdb_Datastore) GetYangName() string { return "datastore" }

func (datastore *ConfdState_Internal_Cdb_Datastore) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (datastore *ConfdState_Internal_Cdb_Datastore) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (datastore *ConfdState_Internal_Cdb_Datastore) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (datastore *ConfdState_Internal_Cdb_Datastore) SetParent(parent types.Entity) { datastore.parent = parent }

func (datastore *ConfdState_Internal_Cdb_Datastore) GetParent() types.Entity { return datastore.parent }

func (datastore *ConfdState_Internal_Cdb_Datastore) GetParentYangName() string { return "cdb" }

// ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync
// Information pertaining to subscription notifications that have
// been delivered to, but not yet acknowledged by, subscribing
// clients. Not present for the 'startup' datastore.
// This type is a presence type.
type ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The priority of the subscriptions that generated the notifications that are
    // waiting for acknowledgement. Not present for the 'operational' datastore.
    // The type is interface{} with range: -2147483648..2147483647.
    Priority interface{}

    // The remaining time in seconds until subscribing clients that have not
    // acknowledged their notifications are considered unresponsive and will be
    // disconnected. See /confdConfig/cdb/clientTimeout in the confd.conf(5)
    // manual page. The value 'infinity' means that no timeout has been configured
    // in confd.conf. The type is one of the following types: int with range:
    // 0..18446744073709551615, or enumeration
    // ConfdState.Internal.Cdb.Datastore.PendingSubscriptionSync.TimeRemaining.
    TimeRemaining interface{}

    // The type is slice of
    // ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification.
    Notification []ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification
}

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetFilter() yfilter.YFilter { return pendingSubscriptionSync.YFilter }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) SetFilter(yf yfilter.YFilter) { pendingSubscriptionSync.YFilter = yf }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetGoName(yname string) string {
    if yname == "priority" { return "Priority" }
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "notification" { return "Notification" }
    return ""
}

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetSegmentPath() string {
    return "pending-subscription-sync"
}

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "notification" {
        for _, c := range pendingSubscriptionSync.Notification {
            if pendingSubscriptionSync.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification{}
        pendingSubscriptionSync.Notification = append(pendingSubscriptionSync.Notification, child)
        return &pendingSubscriptionSync.Notification[len(pendingSubscriptionSync.Notification)-1]
    }
    return nil
}

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pendingSubscriptionSync.Notification {
        children[pendingSubscriptionSync.Notification[i].GetSegmentPath()] = &pendingSubscriptionSync.Notification[i]
    }
    return children
}

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority"] = pendingSubscriptionSync.Priority
    leafs["time-remaining"] = pendingSubscriptionSync.TimeRemaining
    return leafs
}

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetBundleName() string { return "cisco_ios_xe" }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetYangName() string { return "pending-subscription-sync" }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) SetParent(parent types.Entity) { pendingSubscriptionSync.parent = parent }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetParent() types.Entity { return pendingSubscriptionSync.parent }

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetParentYangName() string { return "datastore" }

// ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification
type ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the client that is the recipient of the notification. The type
    // is string.
    ClientName interface{}

    // The subscription identifiers for the subscriptions that generated the
    // notification. The type is slice of interface{} with range: 0..4294967295.
    SubscriptionIds []interface{}
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetFilter() yfilter.YFilter { return notification.YFilter }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) SetFilter(yf yfilter.YFilter) { notification.YFilter = yf }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetGoName(yname string) string {
    if yname == "client-name" { return "ClientName" }
    if yname == "subscription-ids" { return "SubscriptionIds" }
    return ""
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetSegmentPath() string {
    return "notification"
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-name"] = notification.ClientName
    leafs["subscription-ids"] = notification.SubscriptionIds
    return leafs
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetBundleName() string { return "cisco_ios_xe" }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetYangName() string { return "notification" }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) SetParent(parent types.Entity) { notification.parent = parent }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetParent() types.Entity { return notification.parent }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetParentYangName() string { return "pending-subscription-sync" }

// ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_TimeRemaining represents configured in confd.conf.
type ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_TimeRemaining string

const (
    ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_TimeRemaining_infinity ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_TimeRemaining = "infinity"
)

// ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue
// Queues of notifications that have been generated but not
// yet delivered to subscribing clients. Not present for the
// 'startup' datastore.
type ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of
    // ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification.
    Notification []ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification
}

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetFilter() yfilter.YFilter { return pendingNotificationQueue.YFilter }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) SetFilter(yf yfilter.YFilter) { pendingNotificationQueue.YFilter = yf }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetGoName(yname string) string {
    if yname == "notification" { return "Notification" }
    return ""
}

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetSegmentPath() string {
    return "pending-notification-queue"
}

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "notification" {
        for _, c := range pendingNotificationQueue.Notification {
            if pendingNotificationQueue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification{}
        pendingNotificationQueue.Notification = append(pendingNotificationQueue.Notification, child)
        return &pendingNotificationQueue.Notification[len(pendingNotificationQueue.Notification)-1]
    }
    return nil
}

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pendingNotificationQueue.Notification {
        children[pendingNotificationQueue.Notification[i].GetSegmentPath()] = &pendingNotificationQueue.Notification[i]
    }
    return children
}

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetBundleName() string { return "cisco_ios_xe" }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetYangName() string { return "pending-notification-queue" }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) SetParent(parent types.Entity) { pendingNotificationQueue.parent = parent }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetParent() types.Entity { return pendingNotificationQueue.parent }

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetParentYangName() string { return "datastore" }

// ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification
type ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The priority of the subscriptions that generated the notification. Not
    // present for the the 'operational' datastore. The type is interface{} with
    // range: -2147483648..2147483647.
    Priority interface{}

    // The name of the client that is the recipient of the notification. The type
    // is string.
    ClientName interface{}

    // The subscription identifiers for the subscriptions that generated the
    // notification. The type is slice of interface{} with range: 0..4294967295.
    SubscriptionIds []interface{}
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetFilter() yfilter.YFilter { return notification.YFilter }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) SetFilter(yf yfilter.YFilter) { notification.YFilter = yf }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetGoName(yname string) string {
    if yname == "priority" { return "Priority" }
    if yname == "client-name" { return "ClientName" }
    if yname == "subscription-ids" { return "SubscriptionIds" }
    return ""
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetSegmentPath() string {
    return "notification"
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["priority"] = notification.Priority
    leafs["client-name"] = notification.ClientName
    leafs["subscription-ids"] = notification.SubscriptionIds
    return leafs
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetBundleName() string { return "cisco_ios_xe" }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetYangName() string { return "notification" }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) SetParent(parent types.Entity) { notification.parent = parent }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetParent() types.Entity { return notification.parent }

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetParentYangName() string { return "pending-notification-queue" }

// ConfdState_Internal_Cdb_Client
type ConfdState_Internal_Cdb_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The client name. The type is string.
    Name interface{}

    // Additional information about the client obtained at connect time. If
    // present, it consists of client PID and socket file descriptor number. The
    // type is string.
    Info interface{}

    // The type of client: 'inactive' is a client connection without any specific
    // state. 'client' means that the client has an active session towards a
    // datastore. A 'subscriber' has made one or more subscriptions. 'waiting'
    // means that the client is waiting for completion of a call such as
    // cdb_wait_start(). The type is Type.
    Type interface{}

    // The name of the datastore when 'type' = 'client'. The value
    // 'pre_commit_running' is the 'pseudo' datastore representing 'running'
    // before a commit. The type is one of the following types: enumeration
    // ConfdState.Internal.DatastoreName, or enumeration
    // ConfdState.Internal.Cdb.Client.Datastore.
    Datastore interface{}

    // Set when 'type' = 'client' and the client has requested or acquired a lock
    // on the datastore. The 'pending-read' and 'pending-subscription' values
    // indicate that the client has requested but not yet acquired the
    // corresponding lock. A 'read' lock is never taken for the 'operational'
    // datastore, a 'subscription' lock is never taken for any other datastore
    // than 'operational'. The type is Lock.
    Lock interface{}

    // The type is slice of ConfdState_Internal_Cdb_Client_Subscription.
    Subscription []ConfdState_Internal_Cdb_Client_Subscription
}

func (client *ConfdState_Internal_Cdb_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *ConfdState_Internal_Cdb_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *ConfdState_Internal_Cdb_Client) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "info" { return "Info" }
    if yname == "type" { return "Type" }
    if yname == "datastore" { return "Datastore" }
    if yname == "lock" { return "Lock" }
    if yname == "subscription" { return "Subscription" }
    return ""
}

func (client *ConfdState_Internal_Cdb_Client) GetSegmentPath() string {
    return "client"
}

func (client *ConfdState_Internal_Cdb_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscription" {
        for _, c := range client.Subscription {
            if client.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ConfdState_Internal_Cdb_Client_Subscription{}
        client.Subscription = append(client.Subscription, child)
        return &client.Subscription[len(client.Subscription)-1]
    }
    return nil
}

func (client *ConfdState_Internal_Cdb_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range client.Subscription {
        children[client.Subscription[i].GetSegmentPath()] = &client.Subscription[i]
    }
    return children
}

func (client *ConfdState_Internal_Cdb_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = client.Name
    leafs["info"] = client.Info
    leafs["type"] = client.Type
    leafs["datastore"] = client.Datastore
    leafs["lock"] = client.Lock
    return leafs
}

func (client *ConfdState_Internal_Cdb_Client) GetBundleName() string { return "cisco_ios_xe" }

func (client *ConfdState_Internal_Cdb_Client) GetYangName() string { return "client" }

func (client *ConfdState_Internal_Cdb_Client) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (client *ConfdState_Internal_Cdb_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (client *ConfdState_Internal_Cdb_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (client *ConfdState_Internal_Cdb_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *ConfdState_Internal_Cdb_Client) GetParent() types.Entity { return client.parent }

func (client *ConfdState_Internal_Cdb_Client) GetParentYangName() string { return "cdb" }

// ConfdState_Internal_Cdb_Client_Subscription
type ConfdState_Internal_Cdb_Client_Subscription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the datastore for the subscription - only 'running' and
    // 'operational' can have subscriptions. The type is DatastoreName.
    Datastore interface{}

    // Present if this is a 'twophase' subscription, i.e. notifications will be
    // delivered at 'prepare' in addition to 'commit'. Only for the 'running'
    // datastore. The type is interface{}.
    Twophase interface{}

    // The priority of the subscription. The type is interface{} with range:
    // -2147483648..2147483647.
    Priority interface{}

    // The subscription identifier for the subscription. The type is interface{}
    // with range: 0..4294967295.
    Id interface{}

    // The path that the subscription pertains to. The type is string.
    Path interface{}

    // If this leaf exists, there is a problem  with the subscription. The type is
    // Error.
    Error interface{}
}

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetFilter() yfilter.YFilter { return subscription.YFilter }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) SetFilter(yf yfilter.YFilter) { subscription.YFilter = yf }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetGoName(yname string) string {
    if yname == "datastore" { return "Datastore" }
    if yname == "twophase" { return "Twophase" }
    if yname == "priority" { return "Priority" }
    if yname == "id" { return "Id" }
    if yname == "path" { return "Path" }
    if yname == "error" { return "Error" }
    return ""
}

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetSegmentPath() string {
    return "subscription"
}

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["datastore"] = subscription.Datastore
    leafs["twophase"] = subscription.Twophase
    leafs["priority"] = subscription.Priority
    leafs["id"] = subscription.Id
    leafs["path"] = subscription.Path
    leafs["error"] = subscription.Error
    return leafs
}

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetBundleName() string { return "cisco_ios_xe" }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetYangName() string { return "subscription" }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) SetParent(parent types.Entity) { subscription.parent = parent }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetParent() types.Entity { return subscription.parent }

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetParentYangName() string { return "client" }

// ConfdState_Internal_Cdb_Client_Subscription_Error represents  with the subscription.
type ConfdState_Internal_Cdb_Client_Subscription_Error string

const (
    // This value means that the subscribing client has not
    // completed the subscription (with cdb_subscribe_done()).
    ConfdState_Internal_Cdb_Client_Subscription_Error_PENDING ConfdState_Internal_Cdb_Client_Subscription_Error = "PENDING"
)

// ConfdState_Internal_Cdb_Client_Datastore represents 'running' before a commit.
type ConfdState_Internal_Cdb_Client_Datastore string

const (
    ConfdState_Internal_Cdb_Client_Datastore_pre_commit_running ConfdState_Internal_Cdb_Client_Datastore = "pre_commit_running"
)

// ConfdState_Internal_Cdb_Client_Lock represents than 'operational'.
type ConfdState_Internal_Cdb_Client_Lock string

const (
    ConfdState_Internal_Cdb_Client_Lock_read ConfdState_Internal_Cdb_Client_Lock = "read"

    ConfdState_Internal_Cdb_Client_Lock_subscription ConfdState_Internal_Cdb_Client_Lock = "subscription"

    ConfdState_Internal_Cdb_Client_Lock_pending_read ConfdState_Internal_Cdb_Client_Lock = "pending-read"

    ConfdState_Internal_Cdb_Client_Lock_pending_subscription ConfdState_Internal_Cdb_Client_Lock = "pending-subscription"
)

// ConfdState_Internal_Cdb_Client_Type represents waiting for completion of a call such as cdb_wait_start().
type ConfdState_Internal_Cdb_Client_Type string

const (
    ConfdState_Internal_Cdb_Client_Type_inactive ConfdState_Internal_Cdb_Client_Type = "inactive"

    ConfdState_Internal_Cdb_Client_Type_client ConfdState_Internal_Cdb_Client_Type = "client"

    ConfdState_Internal_Cdb_Client_Type_subscriber ConfdState_Internal_Cdb_Client_Type = "subscriber"

    ConfdState_Internal_Cdb_Client_Type_waiting ConfdState_Internal_Cdb_Client_Type = "waiting"
)

// ConfdState_Internal_DatastoreName represents Name of one of the datastores implemented by CDB.
type ConfdState_Internal_DatastoreName string

const (
    ConfdState_Internal_DatastoreName_running ConfdState_Internal_DatastoreName = "running"

    ConfdState_Internal_DatastoreName_startup ConfdState_Internal_DatastoreName = "startup"

    ConfdState_Internal_DatastoreName_operational ConfdState_Internal_DatastoreName = "operational"
)

// ConfdState_DaemonStatus
type ConfdState_DaemonStatus string

const (
    // The daemon is starting up.
    ConfdState_DaemonStatus_starting ConfdState_DaemonStatus = "starting"

    // The daemon is running in start phase 0.
    ConfdState_DaemonStatus_phase0 ConfdState_DaemonStatus = "phase0"

    // The daemon is running in start phase 1.
    ConfdState_DaemonStatus_phase1 ConfdState_DaemonStatus = "phase1"

    // The daemon is started.
    ConfdState_DaemonStatus_started ConfdState_DaemonStatus = "started"

    // The daemon is stopping.
    ConfdState_DaemonStatus_stopping ConfdState_DaemonStatus = "stopping"
)

