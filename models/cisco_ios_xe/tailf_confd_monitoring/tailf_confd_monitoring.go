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
    EntityData types.CommonEntityData
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

func (confdState *ConfdState) GetEntityData() *types.CommonEntityData {
    confdState.EntityData.YFilter = confdState.YFilter
    confdState.EntityData.YangName = "confd-state"
    confdState.EntityData.BundleName = "cisco_ios_xe"
    confdState.EntityData.ParentYangName = "tailf-confd-monitoring"
    confdState.EntityData.SegmentPath = "tailf-confd-monitoring:confd-state"
    confdState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    confdState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    confdState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    confdState.EntityData.Children = types.NewOrderedMap()
    confdState.EntityData.Children.Append("smp", types.YChild{"Smp", &confdState.Smp})
    confdState.EntityData.Children.Append("ha", types.YChild{"Ha", &confdState.Ha})
    confdState.EntityData.Children.Append("loaded-data-models", types.YChild{"LoadedDataModels", &confdState.LoadedDataModels})
    confdState.EntityData.Children.Append("netconf", types.YChild{"Netconf", &confdState.Netconf})
    confdState.EntityData.Children.Append("cli", types.YChild{"Cli", &confdState.Cli})
    confdState.EntityData.Children.Append("webui", types.YChild{"Webui", &confdState.Webui})
    confdState.EntityData.Children.Append("rest", types.YChild{"Rest", &confdState.Rest})
    confdState.EntityData.Children.Append("snmp", types.YChild{"Snmp", &confdState.Snmp})
    confdState.EntityData.Children.Append("internal", types.YChild{"Internal", &confdState.Internal})
    confdState.EntityData.Leafs = types.NewOrderedMap()
    confdState.EntityData.Leafs.Append("version", types.YLeaf{"Version", confdState.Version})
    confdState.EntityData.Leafs.Append("epoll", types.YLeaf{"Epoll", confdState.Epoll})
    confdState.EntityData.Leafs.Append("daemon-status", types.YLeaf{"DaemonStatus", confdState.DaemonStatus})
    confdState.EntityData.Leafs.Append("read-only-mode", types.YLeaf{"ReadOnlyMode", confdState.ReadOnlyMode})
    confdState.EntityData.Leafs.Append("upgrade-mode", types.YLeaf{"UpgradeMode", confdState.UpgradeMode})

    confdState.EntityData.YListKeys = []string {}

    return &(confdState.EntityData)
}

// ConfdState_Smp
// This type is a presence type.
type ConfdState_Smp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Number of threads used by the daemon. The type is interface{} with range:
    // 0..65535.
    NumberOfThreads interface{}
}

func (smp *ConfdState_Smp) GetEntityData() *types.CommonEntityData {
    smp.EntityData.YFilter = smp.YFilter
    smp.EntityData.YangName = "smp"
    smp.EntityData.BundleName = "cisco_ios_xe"
    smp.EntityData.ParentYangName = "confd-state"
    smp.EntityData.SegmentPath = "smp"
    smp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    smp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    smp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    smp.EntityData.Children = types.NewOrderedMap()
    smp.EntityData.Leafs = types.NewOrderedMap()
    smp.EntityData.Leafs.Append("number-of-threads", types.YLeaf{"NumberOfThreads", smp.NumberOfThreads})

    smp.EntityData.YListKeys = []string {}

    return &(smp.EntityData)
}

// ConfdState_Ha
// This type is a presence type.
type ConfdState_Ha struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (ha *ConfdState_Ha) GetEntityData() *types.CommonEntityData {
    ha.EntityData.YFilter = ha.YFilter
    ha.EntityData.YangName = "ha"
    ha.EntityData.BundleName = "cisco_ios_xe"
    ha.EntityData.ParentYangName = "confd-state"
    ha.EntityData.SegmentPath = "ha"
    ha.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ha.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ha.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ha.EntityData.Children = types.NewOrderedMap()
    ha.EntityData.Leafs = types.NewOrderedMap()
    ha.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", ha.Mode})
    ha.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ha.NodeId})
    ha.EntityData.Leafs.Append("master-node-id", types.YLeaf{"MasterNodeId", ha.MasterNodeId})
    ha.EntityData.Leafs.Append("connected-slave", types.YLeaf{"ConnectedSlave", ha.ConnectedSlave})
    ha.EntityData.Leafs.Append("pending-slave", types.YLeaf{"PendingSlave", ha.PendingSlave})

    ha.EntityData.YListKeys = []string {}

    return &(ha.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This list contains all loaded YANG data modules.  This list is a superset
    // of the 'schema' list defined in ietf-netconf-monitoring, which only lists
    // modules exported through NETCONF. The type is slice of
    // ConfdState_LoadedDataModels_DataModel.
    DataModel []*ConfdState_LoadedDataModels_DataModel
}

func (loadedDataModels *ConfdState_LoadedDataModels) GetEntityData() *types.CommonEntityData {
    loadedDataModels.EntityData.YFilter = loadedDataModels.YFilter
    loadedDataModels.EntityData.YangName = "loaded-data-models"
    loadedDataModels.EntityData.BundleName = "cisco_ios_xe"
    loadedDataModels.EntityData.ParentYangName = "confd-state"
    loadedDataModels.EntityData.SegmentPath = "loaded-data-models"
    loadedDataModels.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    loadedDataModels.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    loadedDataModels.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    loadedDataModels.EntityData.Children = types.NewOrderedMap()
    loadedDataModels.EntityData.Children.Append("data-model", types.YChild{"DataModel", nil})
    for i := range loadedDataModels.DataModel {
        loadedDataModels.EntityData.Children.Append(types.GetSegmentPath(loadedDataModels.DataModel[i]), types.YChild{"DataModel", loadedDataModels.DataModel[i]})
    }
    loadedDataModels.EntityData.Leafs = types.NewOrderedMap()

    loadedDataModels.EntityData.YListKeys = []string {}

    return &(loadedDataModels.EntityData)
}

// ConfdState_LoadedDataModels_DataModel
// This list contains all loaded YANG data modules.
// 
// This list is a superset of the 'schema' list defined in
// ietf-netconf-monitoring, which only lists modules exported
// through NETCONF.
type ConfdState_LoadedDataModels_DataModel struct {
    EntityData types.CommonEntityData
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

func (dataModel *ConfdState_LoadedDataModels_DataModel) GetEntityData() *types.CommonEntityData {
    dataModel.EntityData.YFilter = dataModel.YFilter
    dataModel.EntityData.YangName = "data-model"
    dataModel.EntityData.BundleName = "cisco_ios_xe"
    dataModel.EntityData.ParentYangName = "loaded-data-models"
    dataModel.EntityData.SegmentPath = "data-model" + types.AddKeyToken(dataModel.Name, "name")
    dataModel.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dataModel.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dataModel.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dataModel.EntityData.Children = types.NewOrderedMap()
    dataModel.EntityData.Leafs = types.NewOrderedMap()
    dataModel.EntityData.Leafs.Append("name", types.YLeaf{"Name", dataModel.Name})
    dataModel.EntityData.Leafs.Append("revision", types.YLeaf{"Revision", dataModel.Revision})
    dataModel.EntityData.Leafs.Append("namespace", types.YLeaf{"Namespace", dataModel.Namespace})
    dataModel.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", dataModel.Prefix})
    dataModel.EntityData.Leafs.Append("exported-to-all", types.YLeaf{"ExportedToAll", dataModel.ExportedToAll})
    dataModel.EntityData.Leafs.Append("exported-to", types.YLeaf{"ExportedTo", dataModel.ExportedTo})

    dataModel.EntityData.YListKeys = []string {"Name"}

    return &(dataModel.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The transport addresses the NETCONF server is listening on.  Note that
    // other mechanisms can be put in front of the TCP addresses below, e.g., an
    // OpenSSH server.  Such mechanisms are not known to the daemon and not listed
    // here.
    Listen ConfdState_Netconf_Listen
}

func (netconf *ConfdState_Netconf) GetEntityData() *types.CommonEntityData {
    netconf.EntityData.YFilter = netconf.YFilter
    netconf.EntityData.YangName = "netconf"
    netconf.EntityData.BundleName = "cisco_ios_xe"
    netconf.EntityData.ParentYangName = "confd-state"
    netconf.EntityData.SegmentPath = "netconf"
    netconf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    netconf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    netconf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    netconf.EntityData.Children = types.NewOrderedMap()
    netconf.EntityData.Children.Append("listen", types.YChild{"Listen", &netconf.Listen})
    netconf.EntityData.Leafs = types.NewOrderedMap()

    netconf.EntityData.YListKeys = []string {}

    return &(netconf.EntityData)
}

// ConfdState_Netconf_Listen
// The transport addresses the NETCONF server is listening on.
// 
// Note that other mechanisms can be put in front of the TCP
// addresses below, e.g., an OpenSSH server.  Such mechanisms
// are not known to the daemon and not listed here.
type ConfdState_Netconf_Listen struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Netconf_Listen_Tcp.
    Tcp []*ConfdState_Netconf_Listen_Tcp

    // The type is slice of ConfdState_Netconf_Listen_Ssh.
    Ssh []*ConfdState_Netconf_Listen_Ssh
}

func (listen *ConfdState_Netconf_Listen) GetEntityData() *types.CommonEntityData {
    listen.EntityData.YFilter = listen.YFilter
    listen.EntityData.YangName = "listen"
    listen.EntityData.BundleName = "cisco_ios_xe"
    listen.EntityData.ParentYangName = "netconf"
    listen.EntityData.SegmentPath = "listen"
    listen.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    listen.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    listen.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    listen.EntityData.Children = types.NewOrderedMap()
    listen.EntityData.Children.Append("tcp", types.YChild{"Tcp", nil})
    for i := range listen.Tcp {
        listen.EntityData.Children.Append(types.GetSegmentPath(listen.Tcp[i]), types.YChild{"Tcp", listen.Tcp[i]})
    }
    listen.EntityData.Children.Append("ssh", types.YChild{"Ssh", nil})
    for i := range listen.Ssh {
        listen.EntityData.Children.Append(types.GetSegmentPath(listen.Ssh[i]), types.YChild{"Ssh", listen.Ssh[i]})
    }
    listen.EntityData.Leafs = types.NewOrderedMap()

    listen.EntityData.YListKeys = []string {}

    return &(listen.EntityData)
}

// ConfdState_Netconf_Listen_Tcp
type ConfdState_Netconf_Listen_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (tcp *ConfdState_Netconf_Listen_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xe"
    tcp.EntityData.ParentYangName = "listen"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcp.EntityData.Children = types.NewOrderedMap()
    tcp.EntityData.Leafs = types.NewOrderedMap()
    tcp.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", tcp.Ip})
    tcp.EntityData.Leafs.Append("port", types.YLeaf{"Port", tcp.Port})

    tcp.EntityData.YListKeys = []string {}

    return &(tcp.EntityData)
}

// ConfdState_Netconf_Listen_Ssh
type ConfdState_Netconf_Listen_Ssh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (ssh *ConfdState_Netconf_Listen_Ssh) GetEntityData() *types.CommonEntityData {
    ssh.EntityData.YFilter = ssh.YFilter
    ssh.EntityData.YangName = "ssh"
    ssh.EntityData.BundleName = "cisco_ios_xe"
    ssh.EntityData.ParentYangName = "listen"
    ssh.EntityData.SegmentPath = "ssh"
    ssh.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ssh.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ssh.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ssh.EntityData.Children = types.NewOrderedMap()
    ssh.EntityData.Leafs = types.NewOrderedMap()
    ssh.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", ssh.Ip})
    ssh.EntityData.Leafs.Append("port", types.YLeaf{"Port", ssh.Port})

    ssh.EntityData.YListKeys = []string {}

    return &(ssh.EntityData)
}

// ConfdState_Cli
// This type is a presence type.
type ConfdState_Cli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The transport addresses the CLI server is listening on.  In addition to the
    // SSH addresses listen below, the CLI can always be invoked through the
    // daemons IPC port.  Note that other mechanisms can be put in front of the
    // IPC port, e.g., an OpenSSH server.  Such mechanisms are not known to the
    // daemon and not listed here.
    Listen ConfdState_Cli_Listen
}

func (cli *ConfdState_Cli) GetEntityData() *types.CommonEntityData {
    cli.EntityData.YFilter = cli.YFilter
    cli.EntityData.YangName = "cli"
    cli.EntityData.BundleName = "cisco_ios_xe"
    cli.EntityData.ParentYangName = "confd-state"
    cli.EntityData.SegmentPath = "cli"
    cli.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cli.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cli.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cli.EntityData.Children = types.NewOrderedMap()
    cli.EntityData.Children.Append("listen", types.YChild{"Listen", &cli.Listen})
    cli.EntityData.Leafs = types.NewOrderedMap()

    cli.EntityData.YListKeys = []string {}

    return &(cli.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Cli_Listen_Ssh.
    Ssh []*ConfdState_Cli_Listen_Ssh
}

func (listen *ConfdState_Cli_Listen) GetEntityData() *types.CommonEntityData {
    listen.EntityData.YFilter = listen.YFilter
    listen.EntityData.YangName = "listen"
    listen.EntityData.BundleName = "cisco_ios_xe"
    listen.EntityData.ParentYangName = "cli"
    listen.EntityData.SegmentPath = "listen"
    listen.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    listen.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    listen.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    listen.EntityData.Children = types.NewOrderedMap()
    listen.EntityData.Children.Append("ssh", types.YChild{"Ssh", nil})
    for i := range listen.Ssh {
        listen.EntityData.Children.Append(types.GetSegmentPath(listen.Ssh[i]), types.YChild{"Ssh", listen.Ssh[i]})
    }
    listen.EntityData.Leafs = types.NewOrderedMap()

    listen.EntityData.YListKeys = []string {}

    return &(listen.EntityData)
}

// ConfdState_Cli_Listen_Ssh
type ConfdState_Cli_Listen_Ssh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (ssh *ConfdState_Cli_Listen_Ssh) GetEntityData() *types.CommonEntityData {
    ssh.EntityData.YFilter = ssh.YFilter
    ssh.EntityData.YangName = "ssh"
    ssh.EntityData.BundleName = "cisco_ios_xe"
    ssh.EntityData.ParentYangName = "listen"
    ssh.EntityData.SegmentPath = "ssh"
    ssh.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ssh.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ssh.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ssh.EntityData.Children = types.NewOrderedMap()
    ssh.EntityData.Leafs = types.NewOrderedMap()
    ssh.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", ssh.Ip})
    ssh.EntityData.Leafs.Append("port", types.YLeaf{"Port", ssh.Port})

    ssh.EntityData.YListKeys = []string {}

    return &(ssh.EntityData)
}

// ConfdState_Webui
// This type is a presence type.
type ConfdState_Webui struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The transport addresses the WebUI server is listening on.
    Listen ConfdState_Webui_Listen
}

func (webui *ConfdState_Webui) GetEntityData() *types.CommonEntityData {
    webui.EntityData.YFilter = webui.YFilter
    webui.EntityData.YangName = "webui"
    webui.EntityData.BundleName = "cisco_ios_xe"
    webui.EntityData.ParentYangName = "confd-state"
    webui.EntityData.SegmentPath = "webui"
    webui.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    webui.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    webui.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    webui.EntityData.Children = types.NewOrderedMap()
    webui.EntityData.Children.Append("listen", types.YChild{"Listen", &webui.Listen})
    webui.EntityData.Leafs = types.NewOrderedMap()

    webui.EntityData.YListKeys = []string {}

    return &(webui.EntityData)
}

// ConfdState_Webui_Listen
// The transport addresses the WebUI server is listening on.
type ConfdState_Webui_Listen struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Webui_Listen_Tcp.
    Tcp []*ConfdState_Webui_Listen_Tcp

    // The type is slice of ConfdState_Webui_Listen_Ssl.
    Ssl []*ConfdState_Webui_Listen_Ssl
}

func (listen *ConfdState_Webui_Listen) GetEntityData() *types.CommonEntityData {
    listen.EntityData.YFilter = listen.YFilter
    listen.EntityData.YangName = "listen"
    listen.EntityData.BundleName = "cisco_ios_xe"
    listen.EntityData.ParentYangName = "webui"
    listen.EntityData.SegmentPath = "listen"
    listen.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    listen.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    listen.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    listen.EntityData.Children = types.NewOrderedMap()
    listen.EntityData.Children.Append("tcp", types.YChild{"Tcp", nil})
    for i := range listen.Tcp {
        listen.EntityData.Children.Append(types.GetSegmentPath(listen.Tcp[i]), types.YChild{"Tcp", listen.Tcp[i]})
    }
    listen.EntityData.Children.Append("ssl", types.YChild{"Ssl", nil})
    for i := range listen.Ssl {
        listen.EntityData.Children.Append(types.GetSegmentPath(listen.Ssl[i]), types.YChild{"Ssl", listen.Ssl[i]})
    }
    listen.EntityData.Leafs = types.NewOrderedMap()

    listen.EntityData.YListKeys = []string {}

    return &(listen.EntityData)
}

// ConfdState_Webui_Listen_Tcp
type ConfdState_Webui_Listen_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (tcp *ConfdState_Webui_Listen_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xe"
    tcp.EntityData.ParentYangName = "listen"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcp.EntityData.Children = types.NewOrderedMap()
    tcp.EntityData.Leafs = types.NewOrderedMap()
    tcp.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", tcp.Ip})
    tcp.EntityData.Leafs.Append("port", types.YLeaf{"Port", tcp.Port})

    tcp.EntityData.YListKeys = []string {}

    return &(tcp.EntityData)
}

// ConfdState_Webui_Listen_Ssl
type ConfdState_Webui_Listen_Ssl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (ssl *ConfdState_Webui_Listen_Ssl) GetEntityData() *types.CommonEntityData {
    ssl.EntityData.YFilter = ssl.YFilter
    ssl.EntityData.YangName = "ssl"
    ssl.EntityData.BundleName = "cisco_ios_xe"
    ssl.EntityData.ParentYangName = "listen"
    ssl.EntityData.SegmentPath = "ssl"
    ssl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ssl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ssl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ssl.EntityData.Children = types.NewOrderedMap()
    ssl.EntityData.Leafs = types.NewOrderedMap()
    ssl.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", ssl.Ip})
    ssl.EntityData.Leafs.Append("port", types.YLeaf{"Port", ssl.Port})

    ssl.EntityData.YListKeys = []string {}

    return &(ssl.EntityData)
}

// ConfdState_Rest
// This type is a presence type.
type ConfdState_Rest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The transport addresses the REST server is listening on.
    Listen ConfdState_Rest_Listen
}

func (rest *ConfdState_Rest) GetEntityData() *types.CommonEntityData {
    rest.EntityData.YFilter = rest.YFilter
    rest.EntityData.YangName = "rest"
    rest.EntityData.BundleName = "cisco_ios_xe"
    rest.EntityData.ParentYangName = "confd-state"
    rest.EntityData.SegmentPath = "rest"
    rest.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rest.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rest.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rest.EntityData.Children = types.NewOrderedMap()
    rest.EntityData.Children.Append("listen", types.YChild{"Listen", &rest.Listen})
    rest.EntityData.Leafs = types.NewOrderedMap()

    rest.EntityData.YListKeys = []string {}

    return &(rest.EntityData)
}

// ConfdState_Rest_Listen
// The transport addresses the REST server is listening on.
type ConfdState_Rest_Listen struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Rest_Listen_Tcp.
    Tcp []*ConfdState_Rest_Listen_Tcp

    // The type is slice of ConfdState_Rest_Listen_Ssl.
    Ssl []*ConfdState_Rest_Listen_Ssl
}

func (listen *ConfdState_Rest_Listen) GetEntityData() *types.CommonEntityData {
    listen.EntityData.YFilter = listen.YFilter
    listen.EntityData.YangName = "listen"
    listen.EntityData.BundleName = "cisco_ios_xe"
    listen.EntityData.ParentYangName = "rest"
    listen.EntityData.SegmentPath = "listen"
    listen.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    listen.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    listen.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    listen.EntityData.Children = types.NewOrderedMap()
    listen.EntityData.Children.Append("tcp", types.YChild{"Tcp", nil})
    for i := range listen.Tcp {
        listen.EntityData.Children.Append(types.GetSegmentPath(listen.Tcp[i]), types.YChild{"Tcp", listen.Tcp[i]})
    }
    listen.EntityData.Children.Append("ssl", types.YChild{"Ssl", nil})
    for i := range listen.Ssl {
        listen.EntityData.Children.Append(types.GetSegmentPath(listen.Ssl[i]), types.YChild{"Ssl", listen.Ssl[i]})
    }
    listen.EntityData.Leafs = types.NewOrderedMap()

    listen.EntityData.YListKeys = []string {}

    return &(listen.EntityData)
}

// ConfdState_Rest_Listen_Tcp
type ConfdState_Rest_Listen_Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (tcp *ConfdState_Rest_Listen_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xe"
    tcp.EntityData.ParentYangName = "listen"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcp.EntityData.Children = types.NewOrderedMap()
    tcp.EntityData.Leafs = types.NewOrderedMap()
    tcp.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", tcp.Ip})
    tcp.EntityData.Leafs.Append("port", types.YLeaf{"Port", tcp.Port})

    tcp.EntityData.YListKeys = []string {}

    return &(tcp.EntityData)
}

// ConfdState_Rest_Listen_Ssl
type ConfdState_Rest_Listen_Ssl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (ssl *ConfdState_Rest_Listen_Ssl) GetEntityData() *types.CommonEntityData {
    ssl.EntityData.YFilter = ssl.YFilter
    ssl.EntityData.YangName = "ssl"
    ssl.EntityData.BundleName = "cisco_ios_xe"
    ssl.EntityData.ParentYangName = "listen"
    ssl.EntityData.SegmentPath = "ssl"
    ssl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ssl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ssl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ssl.EntityData.Children = types.NewOrderedMap()
    ssl.EntityData.Leafs = types.NewOrderedMap()
    ssl.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", ssl.Ip})
    ssl.EntityData.Leafs.Append("port", types.YLeaf{"Port", ssl.Port})

    ssl.EntityData.YListKeys = []string {}

    return &(ssl.EntityData)
}

// ConfdState_Snmp
// This type is a presence type.
type ConfdState_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (snmp *ConfdState_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xe"
    snmp.EntityData.ParentYangName = "confd-state"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Children.Append("listen", types.YChild{"Listen", &snmp.Listen})
    snmp.EntityData.Children.Append("version", types.YChild{"Version", &snmp.Version})
    snmp.EntityData.Leafs = types.NewOrderedMap()
    snmp.EntityData.Leafs.Append("mib", types.YLeaf{"Mib", snmp.Mib})
    snmp.EntityData.Leafs.Append("engine-id", types.YLeaf{"EngineId", snmp.EngineId})

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// ConfdState_Snmp_Listen
// The transport addresses the SNMP agent is listening on.
type ConfdState_Snmp_Listen struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Snmp_Listen_Udp.
    Udp []*ConfdState_Snmp_Listen_Udp
}

func (listen *ConfdState_Snmp_Listen) GetEntityData() *types.CommonEntityData {
    listen.EntityData.YFilter = listen.YFilter
    listen.EntityData.YangName = "listen"
    listen.EntityData.BundleName = "cisco_ios_xe"
    listen.EntityData.ParentYangName = "snmp"
    listen.EntityData.SegmentPath = "listen"
    listen.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    listen.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    listen.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    listen.EntityData.Children = types.NewOrderedMap()
    listen.EntityData.Children.Append("udp", types.YChild{"Udp", nil})
    for i := range listen.Udp {
        listen.EntityData.Children.Append(types.GetSegmentPath(listen.Udp[i]), types.YChild{"Udp", listen.Udp[i]})
    }
    listen.EntityData.Leafs = types.NewOrderedMap()

    listen.EntityData.YListKeys = []string {}

    return &(listen.EntityData)
}

// ConfdState_Snmp_Listen_Udp
type ConfdState_Snmp_Listen_Udp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // The type is interface{} with range: 0..65535.
    Port interface{}
}

func (udp *ConfdState_Snmp_Listen_Udp) GetEntityData() *types.CommonEntityData {
    udp.EntityData.YFilter = udp.YFilter
    udp.EntityData.YangName = "udp"
    udp.EntityData.BundleName = "cisco_ios_xe"
    udp.EntityData.ParentYangName = "listen"
    udp.EntityData.SegmentPath = "udp"
    udp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udp.EntityData.Children = types.NewOrderedMap()
    udp.EntityData.Leafs = types.NewOrderedMap()
    udp.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", udp.Ip})
    udp.EntityData.Leafs.Append("port", types.YLeaf{"Port", udp.Port})

    udp.EntityData.YListKeys = []string {}

    return &(udp.EntityData)
}

// ConfdState_Snmp_Version
// SNMP version used by the engine.
type ConfdState_Snmp_Version struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    V1 interface{}

    // The type is interface{}.
    V2c interface{}

    // The type is interface{}.
    V3 interface{}
}

func (version *ConfdState_Snmp_Version) GetEntityData() *types.CommonEntityData {
    version.EntityData.YFilter = version.YFilter
    version.EntityData.YangName = "version"
    version.EntityData.BundleName = "cisco_ios_xe"
    version.EntityData.ParentYangName = "snmp"
    version.EntityData.SegmentPath = "version"
    version.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    version.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    version.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    version.EntityData.Children = types.NewOrderedMap()
    version.EntityData.Leafs = types.NewOrderedMap()
    version.EntityData.Leafs.Append("v1", types.YLeaf{"V1", version.V1})
    version.EntityData.Leafs.Append("v2c", types.YLeaf{"V2c", version.V2c})
    version.EntityData.Leafs.Append("v3", types.YLeaf{"V3", version.V3})

    version.EntityData.YListKeys = []string {}

    return &(version.EntityData)
}

// ConfdState_Internal
type ConfdState_Internal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Callpoints ConfdState_Internal_Callpoints

    
    Cdb ConfdState_Internal_Cdb
}

func (internal *ConfdState_Internal) GetEntityData() *types.CommonEntityData {
    internal.EntityData.YFilter = internal.YFilter
    internal.EntityData.YangName = "internal"
    internal.EntityData.BundleName = "cisco_ios_xe"
    internal.EntityData.ParentYangName = "confd-state"
    internal.EntityData.SegmentPath = "internal"
    internal.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    internal.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    internal.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    internal.EntityData.Children = types.NewOrderedMap()
    internal.EntityData.Children.Append("callpoints", types.YChild{"Callpoints", &internal.Callpoints})
    internal.EntityData.Children.Append("cdb", types.YChild{"Cdb", &internal.Cdb})
    internal.EntityData.Leafs = types.NewOrderedMap()

    internal.EntityData.YListKeys = []string {}

    return &(internal.EntityData)
}

// ConfdState_Internal_Callpoints
type ConfdState_Internal_Callpoints struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Internal_Callpoints_Callpoint.
    Callpoint []*ConfdState_Internal_Callpoints_Callpoint

    // The type is slice of ConfdState_Internal_Callpoints_Validationpoint.
    Validationpoint []*ConfdState_Internal_Callpoints_Validationpoint

    // The type is slice of ConfdState_Internal_Callpoints_Actionpoint.
    Actionpoint []*ConfdState_Internal_Callpoints_Actionpoint

    // The type is slice of ConfdState_Internal_Callpoints_SnmpInformCallback.
    SnmpInformCallback []*ConfdState_Internal_Callpoints_SnmpInformCallback

    // The type is slice of
    // ConfdState_Internal_Callpoints_SnmpNotificationSubscription.
    SnmpNotificationSubscription []*ConfdState_Internal_Callpoints_SnmpNotificationSubscription

    // The type is slice of
    // ConfdState_Internal_Callpoints_ErrorFormattingCallback.
    ErrorFormattingCallback []*ConfdState_Internal_Callpoints_ErrorFormattingCallback

    // The type is slice of ConfdState_Internal_Callpoints_Typepoint.
    Typepoint []*ConfdState_Internal_Callpoints_Typepoint

    // The type is slice of
    // ConfdState_Internal_Callpoints_NotificationStreamReplay.
    NotificationStreamReplay []*ConfdState_Internal_Callpoints_NotificationStreamReplay

    
    AuthenticationCallback ConfdState_Internal_Callpoints_AuthenticationCallback

    
    AuthorizationCallbacks ConfdState_Internal_Callpoints_AuthorizationCallbacks
}

func (callpoints *ConfdState_Internal_Callpoints) GetEntityData() *types.CommonEntityData {
    callpoints.EntityData.YFilter = callpoints.YFilter
    callpoints.EntityData.YangName = "callpoints"
    callpoints.EntityData.BundleName = "cisco_ios_xe"
    callpoints.EntityData.ParentYangName = "internal"
    callpoints.EntityData.SegmentPath = "callpoints"
    callpoints.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callpoints.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callpoints.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callpoints.EntityData.Children = types.NewOrderedMap()
    callpoints.EntityData.Children.Append("callpoint", types.YChild{"Callpoint", nil})
    for i := range callpoints.Callpoint {
        callpoints.EntityData.Children.Append(types.GetSegmentPath(callpoints.Callpoint[i]), types.YChild{"Callpoint", callpoints.Callpoint[i]})
    }
    callpoints.EntityData.Children.Append("validationpoint", types.YChild{"Validationpoint", nil})
    for i := range callpoints.Validationpoint {
        callpoints.EntityData.Children.Append(types.GetSegmentPath(callpoints.Validationpoint[i]), types.YChild{"Validationpoint", callpoints.Validationpoint[i]})
    }
    callpoints.EntityData.Children.Append("actionpoint", types.YChild{"Actionpoint", nil})
    for i := range callpoints.Actionpoint {
        callpoints.EntityData.Children.Append(types.GetSegmentPath(callpoints.Actionpoint[i]), types.YChild{"Actionpoint", callpoints.Actionpoint[i]})
    }
    callpoints.EntityData.Children.Append("snmp-inform-callback", types.YChild{"SnmpInformCallback", nil})
    for i := range callpoints.SnmpInformCallback {
        callpoints.EntityData.Children.Append(types.GetSegmentPath(callpoints.SnmpInformCallback[i]), types.YChild{"SnmpInformCallback", callpoints.SnmpInformCallback[i]})
    }
    callpoints.EntityData.Children.Append("snmp-notification-subscription", types.YChild{"SnmpNotificationSubscription", nil})
    for i := range callpoints.SnmpNotificationSubscription {
        callpoints.EntityData.Children.Append(types.GetSegmentPath(callpoints.SnmpNotificationSubscription[i]), types.YChild{"SnmpNotificationSubscription", callpoints.SnmpNotificationSubscription[i]})
    }
    callpoints.EntityData.Children.Append("error-formatting-callback", types.YChild{"ErrorFormattingCallback", nil})
    for i := range callpoints.ErrorFormattingCallback {
        callpoints.EntityData.Children.Append(types.GetSegmentPath(callpoints.ErrorFormattingCallback[i]), types.YChild{"ErrorFormattingCallback", callpoints.ErrorFormattingCallback[i]})
    }
    callpoints.EntityData.Children.Append("typepoint", types.YChild{"Typepoint", nil})
    for i := range callpoints.Typepoint {
        callpoints.EntityData.Children.Append(types.GetSegmentPath(callpoints.Typepoint[i]), types.YChild{"Typepoint", callpoints.Typepoint[i]})
    }
    callpoints.EntityData.Children.Append("notification-stream-replay", types.YChild{"NotificationStreamReplay", nil})
    for i := range callpoints.NotificationStreamReplay {
        callpoints.EntityData.Children.Append(types.GetSegmentPath(callpoints.NotificationStreamReplay[i]), types.YChild{"NotificationStreamReplay", callpoints.NotificationStreamReplay[i]})
    }
    callpoints.EntityData.Children.Append("authentication-callback", types.YChild{"AuthenticationCallback", &callpoints.AuthenticationCallback})
    callpoints.EntityData.Children.Append("authorization-callbacks", types.YChild{"AuthorizationCallbacks", &callpoints.AuthorizationCallbacks})
    callpoints.EntityData.Leafs = types.NewOrderedMap()

    callpoints.EntityData.YListKeys = []string {}

    return &(callpoints.EntityData)
}

// ConfdState_Internal_Callpoints_Callpoint
type ConfdState_Internal_Callpoints_Callpoint struct {
    EntityData types.CommonEntityData
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
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_Callpoint_Daemon

    // The type is slice of ConfdState_Internal_Callpoints_Callpoint_Range.
    Range []*ConfdState_Internal_Callpoints_Callpoint_Range
}

func (callpoint *ConfdState_Internal_Callpoints_Callpoint) GetEntityData() *types.CommonEntityData {
    callpoint.EntityData.YFilter = callpoint.YFilter
    callpoint.EntityData.YangName = "callpoint"
    callpoint.EntityData.BundleName = "cisco_ios_xe"
    callpoint.EntityData.ParentYangName = "callpoints"
    callpoint.EntityData.SegmentPath = "callpoint" + types.AddKeyToken(callpoint.Id, "id")
    callpoint.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    callpoint.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    callpoint.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    callpoint.EntityData.Children = types.NewOrderedMap()
    callpoint.EntityData.Children.Append("daemon", types.YChild{"Daemon", &callpoint.Daemon})
    callpoint.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range callpoint.Range {
        callpoint.EntityData.Children.Append(types.GetSegmentPath(callpoint.Range[i]), types.YChild{"Range", callpoint.Range[i]})
    }
    callpoint.EntityData.Leafs = types.NewOrderedMap()
    callpoint.EntityData.Leafs.Append("id", types.YLeaf{"Id", callpoint.Id})
    callpoint.EntityData.Leafs.Append("path", types.YLeaf{"Path", callpoint.Path})
    callpoint.EntityData.Leafs.Append("file", types.YLeaf{"File", callpoint.File})
    callpoint.EntityData.Leafs.Append("error", types.YLeaf{"Error", callpoint.Error})

    callpoint.EntityData.YListKeys = []string {"Id"}

    return &(callpoint.EntityData)
}

// ConfdState_Internal_Callpoints_Callpoint_Daemon
type ConfdState_Internal_Callpoints_Callpoint_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "callpoint"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_Callpoint_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_Callpoint_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Callpoint_Daemon_Error__PENDING ConfdState_Internal_Callpoints_Callpoint_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_Callpoint_Range
type ConfdState_Internal_Callpoints_Callpoint_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_Callpoint_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "callpoint"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_Callpoint_Range_Daemon
type ConfdState_Internal_Callpoints_Callpoint_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Callpoint_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_Callpoint_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_Callpoint_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Callpoint_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_Callpoint_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_Callpoint_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_Callpoint_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_Callpoint_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_Callpoint_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_Callpoint_Error__UNKNOWN ConfdState_Internal_Callpoints_Callpoint_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_Validationpoint
type ConfdState_Internal_Callpoints_Validationpoint struct {
    EntityData types.CommonEntityData
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
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_Validationpoint_Daemon

    // The type is slice of ConfdState_Internal_Callpoints_Validationpoint_Range.
    Range []*ConfdState_Internal_Callpoints_Validationpoint_Range
}

func (validationpoint *ConfdState_Internal_Callpoints_Validationpoint) GetEntityData() *types.CommonEntityData {
    validationpoint.EntityData.YFilter = validationpoint.YFilter
    validationpoint.EntityData.YangName = "validationpoint"
    validationpoint.EntityData.BundleName = "cisco_ios_xe"
    validationpoint.EntityData.ParentYangName = "callpoints"
    validationpoint.EntityData.SegmentPath = "validationpoint" + types.AddKeyToken(validationpoint.Id, "id")
    validationpoint.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    validationpoint.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    validationpoint.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    validationpoint.EntityData.Children = types.NewOrderedMap()
    validationpoint.EntityData.Children.Append("daemon", types.YChild{"Daemon", &validationpoint.Daemon})
    validationpoint.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range validationpoint.Range {
        validationpoint.EntityData.Children.Append(types.GetSegmentPath(validationpoint.Range[i]), types.YChild{"Range", validationpoint.Range[i]})
    }
    validationpoint.EntityData.Leafs = types.NewOrderedMap()
    validationpoint.EntityData.Leafs.Append("id", types.YLeaf{"Id", validationpoint.Id})
    validationpoint.EntityData.Leafs.Append("path", types.YLeaf{"Path", validationpoint.Path})
    validationpoint.EntityData.Leafs.Append("file", types.YLeaf{"File", validationpoint.File})
    validationpoint.EntityData.Leafs.Append("error", types.YLeaf{"Error", validationpoint.Error})

    validationpoint.EntityData.YListKeys = []string {"Id"}

    return &(validationpoint.EntityData)
}

// ConfdState_Internal_Callpoints_Validationpoint_Daemon
type ConfdState_Internal_Callpoints_Validationpoint_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "validationpoint"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_Validationpoint_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_Validationpoint_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Validationpoint_Daemon_Error__PENDING ConfdState_Internal_Callpoints_Validationpoint_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_Validationpoint_Range
type ConfdState_Internal_Callpoints_Validationpoint_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_Validationpoint_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "validationpoint"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon
type ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_Validationpoint_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_Validationpoint_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_Validationpoint_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_Validationpoint_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_Validationpoint_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_Validationpoint_Error__UNKNOWN ConfdState_Internal_Callpoints_Validationpoint_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_Actionpoint
type ConfdState_Internal_Callpoints_Actionpoint struct {
    EntityData types.CommonEntityData
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
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_Actionpoint_Daemon

    // The type is slice of ConfdState_Internal_Callpoints_Actionpoint_Range.
    Range []*ConfdState_Internal_Callpoints_Actionpoint_Range
}

func (actionpoint *ConfdState_Internal_Callpoints_Actionpoint) GetEntityData() *types.CommonEntityData {
    actionpoint.EntityData.YFilter = actionpoint.YFilter
    actionpoint.EntityData.YangName = "actionpoint"
    actionpoint.EntityData.BundleName = "cisco_ios_xe"
    actionpoint.EntityData.ParentYangName = "callpoints"
    actionpoint.EntityData.SegmentPath = "actionpoint" + types.AddKeyToken(actionpoint.Id, "id")
    actionpoint.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    actionpoint.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    actionpoint.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    actionpoint.EntityData.Children = types.NewOrderedMap()
    actionpoint.EntityData.Children.Append("daemon", types.YChild{"Daemon", &actionpoint.Daemon})
    actionpoint.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range actionpoint.Range {
        actionpoint.EntityData.Children.Append(types.GetSegmentPath(actionpoint.Range[i]), types.YChild{"Range", actionpoint.Range[i]})
    }
    actionpoint.EntityData.Leafs = types.NewOrderedMap()
    actionpoint.EntityData.Leafs.Append("id", types.YLeaf{"Id", actionpoint.Id})
    actionpoint.EntityData.Leafs.Append("path", types.YLeaf{"Path", actionpoint.Path})
    actionpoint.EntityData.Leafs.Append("file", types.YLeaf{"File", actionpoint.File})
    actionpoint.EntityData.Leafs.Append("error", types.YLeaf{"Error", actionpoint.Error})

    actionpoint.EntityData.YListKeys = []string {"Id"}

    return &(actionpoint.EntityData)
}

// ConfdState_Internal_Callpoints_Actionpoint_Daemon
type ConfdState_Internal_Callpoints_Actionpoint_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "actionpoint"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_Actionpoint_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_Actionpoint_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Actionpoint_Daemon_Error__PENDING ConfdState_Internal_Callpoints_Actionpoint_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_Actionpoint_Range
type ConfdState_Internal_Callpoints_Actionpoint_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_Actionpoint_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "actionpoint"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon
type ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_Actionpoint_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_Actionpoint_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_Actionpoint_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_Actionpoint_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_Actionpoint_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_Actionpoint_Error__UNKNOWN ConfdState_Internal_Callpoints_Actionpoint_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_SnmpInformCallback
type ConfdState_Internal_Callpoints_SnmpInformCallback struct {
    EntityData types.CommonEntityData
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
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_SnmpInformCallback_Range.
    Range []*ConfdState_Internal_Callpoints_SnmpInformCallback_Range
}

func (snmpInformCallback *ConfdState_Internal_Callpoints_SnmpInformCallback) GetEntityData() *types.CommonEntityData {
    snmpInformCallback.EntityData.YFilter = snmpInformCallback.YFilter
    snmpInformCallback.EntityData.YangName = "snmp-inform-callback"
    snmpInformCallback.EntityData.BundleName = "cisco_ios_xe"
    snmpInformCallback.EntityData.ParentYangName = "callpoints"
    snmpInformCallback.EntityData.SegmentPath = "snmp-inform-callback" + types.AddKeyToken(snmpInformCallback.Id, "id")
    snmpInformCallback.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpInformCallback.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpInformCallback.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpInformCallback.EntityData.Children = types.NewOrderedMap()
    snmpInformCallback.EntityData.Children.Append("daemon", types.YChild{"Daemon", &snmpInformCallback.Daemon})
    snmpInformCallback.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range snmpInformCallback.Range {
        snmpInformCallback.EntityData.Children.Append(types.GetSegmentPath(snmpInformCallback.Range[i]), types.YChild{"Range", snmpInformCallback.Range[i]})
    }
    snmpInformCallback.EntityData.Leafs = types.NewOrderedMap()
    snmpInformCallback.EntityData.Leafs.Append("id", types.YLeaf{"Id", snmpInformCallback.Id})
    snmpInformCallback.EntityData.Leafs.Append("path", types.YLeaf{"Path", snmpInformCallback.Path})
    snmpInformCallback.EntityData.Leafs.Append("file", types.YLeaf{"File", snmpInformCallback.File})
    snmpInformCallback.EntityData.Leafs.Append("error", types.YLeaf{"Error", snmpInformCallback.Error})

    snmpInformCallback.EntityData.YListKeys = []string {"Id"}

    return &(snmpInformCallback.EntityData)
}

// ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon
type ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "snmp-inform-callback"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon_Error__PENDING ConfdState_Internal_Callpoints_SnmpInformCallback_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_SnmpInformCallback_Range
type ConfdState_Internal_Callpoints_SnmpInformCallback_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_SnmpInformCallback_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "snmp-inform-callback"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon
type ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_SnmpInformCallback_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_SnmpInformCallback_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_SnmpInformCallback_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_SnmpInformCallback_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_SnmpInformCallback_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_SnmpInformCallback_Error__UNKNOWN ConfdState_Internal_Callpoints_SnmpInformCallback_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription struct {
    EntityData types.CommonEntityData
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
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range.
    Range []*ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range
}

func (snmpNotificationSubscription *ConfdState_Internal_Callpoints_SnmpNotificationSubscription) GetEntityData() *types.CommonEntityData {
    snmpNotificationSubscription.EntityData.YFilter = snmpNotificationSubscription.YFilter
    snmpNotificationSubscription.EntityData.YangName = "snmp-notification-subscription"
    snmpNotificationSubscription.EntityData.BundleName = "cisco_ios_xe"
    snmpNotificationSubscription.EntityData.ParentYangName = "callpoints"
    snmpNotificationSubscription.EntityData.SegmentPath = "snmp-notification-subscription" + types.AddKeyToken(snmpNotificationSubscription.Id, "id")
    snmpNotificationSubscription.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    snmpNotificationSubscription.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    snmpNotificationSubscription.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    snmpNotificationSubscription.EntityData.Children = types.NewOrderedMap()
    snmpNotificationSubscription.EntityData.Children.Append("daemon", types.YChild{"Daemon", &snmpNotificationSubscription.Daemon})
    snmpNotificationSubscription.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range snmpNotificationSubscription.Range {
        snmpNotificationSubscription.EntityData.Children.Append(types.GetSegmentPath(snmpNotificationSubscription.Range[i]), types.YChild{"Range", snmpNotificationSubscription.Range[i]})
    }
    snmpNotificationSubscription.EntityData.Leafs = types.NewOrderedMap()
    snmpNotificationSubscription.EntityData.Leafs.Append("id", types.YLeaf{"Id", snmpNotificationSubscription.Id})
    snmpNotificationSubscription.EntityData.Leafs.Append("path", types.YLeaf{"Path", snmpNotificationSubscription.Path})
    snmpNotificationSubscription.EntityData.Leafs.Append("file", types.YLeaf{"File", snmpNotificationSubscription.File})
    snmpNotificationSubscription.EntityData.Leafs.Append("error", types.YLeaf{"Error", snmpNotificationSubscription.Error})

    snmpNotificationSubscription.EntityData.YListKeys = []string {"Id"}

    return &(snmpNotificationSubscription.EntityData)
}

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "snmp-notification-subscription"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon_Error__PENDING ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "snmp-notification-subscription"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error__UNKNOWN ConfdState_Internal_Callpoints_SnmpNotificationSubscription_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_ErrorFormattingCallback
type ConfdState_Internal_Callpoints_ErrorFormattingCallback struct {
    EntityData types.CommonEntityData
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
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range.
    Range []*ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range
}

func (errorFormattingCallback *ConfdState_Internal_Callpoints_ErrorFormattingCallback) GetEntityData() *types.CommonEntityData {
    errorFormattingCallback.EntityData.YFilter = errorFormattingCallback.YFilter
    errorFormattingCallback.EntityData.YangName = "error-formatting-callback"
    errorFormattingCallback.EntityData.BundleName = "cisco_ios_xe"
    errorFormattingCallback.EntityData.ParentYangName = "callpoints"
    errorFormattingCallback.EntityData.SegmentPath = "error-formatting-callback" + types.AddKeyToken(errorFormattingCallback.Id, "id")
    errorFormattingCallback.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    errorFormattingCallback.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    errorFormattingCallback.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    errorFormattingCallback.EntityData.Children = types.NewOrderedMap()
    errorFormattingCallback.EntityData.Children.Append("daemon", types.YChild{"Daemon", &errorFormattingCallback.Daemon})
    errorFormattingCallback.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range errorFormattingCallback.Range {
        errorFormattingCallback.EntityData.Children.Append(types.GetSegmentPath(errorFormattingCallback.Range[i]), types.YChild{"Range", errorFormattingCallback.Range[i]})
    }
    errorFormattingCallback.EntityData.Leafs = types.NewOrderedMap()
    errorFormattingCallback.EntityData.Leafs.Append("id", types.YLeaf{"Id", errorFormattingCallback.Id})
    errorFormattingCallback.EntityData.Leafs.Append("path", types.YLeaf{"Path", errorFormattingCallback.Path})
    errorFormattingCallback.EntityData.Leafs.Append("file", types.YLeaf{"File", errorFormattingCallback.File})
    errorFormattingCallback.EntityData.Leafs.Append("error", types.YLeaf{"Error", errorFormattingCallback.Error})

    errorFormattingCallback.EntityData.YListKeys = []string {"Id"}

    return &(errorFormattingCallback.EntityData)
}

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "error-formatting-callback"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon_Error__PENDING ConfdState_Internal_Callpoints_ErrorFormattingCallback_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "error-formatting-callback"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_ErrorFormattingCallback_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error__UNKNOWN ConfdState_Internal_Callpoints_ErrorFormattingCallback_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_Typepoint
type ConfdState_Internal_Callpoints_Typepoint struct {
    EntityData types.CommonEntityData
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
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_Typepoint_Daemon

    // The type is slice of ConfdState_Internal_Callpoints_Typepoint_Range.
    Range []*ConfdState_Internal_Callpoints_Typepoint_Range
}

func (typepoint *ConfdState_Internal_Callpoints_Typepoint) GetEntityData() *types.CommonEntityData {
    typepoint.EntityData.YFilter = typepoint.YFilter
    typepoint.EntityData.YangName = "typepoint"
    typepoint.EntityData.BundleName = "cisco_ios_xe"
    typepoint.EntityData.ParentYangName = "callpoints"
    typepoint.EntityData.SegmentPath = "typepoint" + types.AddKeyToken(typepoint.Id, "id")
    typepoint.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    typepoint.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    typepoint.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    typepoint.EntityData.Children = types.NewOrderedMap()
    typepoint.EntityData.Children.Append("daemon", types.YChild{"Daemon", &typepoint.Daemon})
    typepoint.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range typepoint.Range {
        typepoint.EntityData.Children.Append(types.GetSegmentPath(typepoint.Range[i]), types.YChild{"Range", typepoint.Range[i]})
    }
    typepoint.EntityData.Leafs = types.NewOrderedMap()
    typepoint.EntityData.Leafs.Append("id", types.YLeaf{"Id", typepoint.Id})
    typepoint.EntityData.Leafs.Append("path", types.YLeaf{"Path", typepoint.Path})
    typepoint.EntityData.Leafs.Append("file", types.YLeaf{"File", typepoint.File})
    typepoint.EntityData.Leafs.Append("error", types.YLeaf{"Error", typepoint.Error})

    typepoint.EntityData.YListKeys = []string {"Id"}

    return &(typepoint.EntityData)
}

// ConfdState_Internal_Callpoints_Typepoint_Daemon
type ConfdState_Internal_Callpoints_Typepoint_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "typepoint"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_Typepoint_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_Typepoint_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Typepoint_Daemon_Error__PENDING ConfdState_Internal_Callpoints_Typepoint_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_Typepoint_Range
type ConfdState_Internal_Callpoints_Typepoint_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_Typepoint_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "typepoint"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_Typepoint_Range_Daemon
type ConfdState_Internal_Callpoints_Typepoint_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_Typepoint_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_Typepoint_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_Typepoint_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_Typepoint_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_Typepoint_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_Typepoint_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_Typepoint_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_Typepoint_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_Typepoint_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_Typepoint_Error__UNKNOWN ConfdState_Internal_Callpoints_Typepoint_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_NotificationStreamReplay
type ConfdState_Internal_Callpoints_NotificationStreamReplay struct {
    EntityData types.CommonEntityData
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
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_NotificationStreamReplay_Range.
    Range []*ConfdState_Internal_Callpoints_NotificationStreamReplay_Range
}

func (notificationStreamReplay *ConfdState_Internal_Callpoints_NotificationStreamReplay) GetEntityData() *types.CommonEntityData {
    notificationStreamReplay.EntityData.YFilter = notificationStreamReplay.YFilter
    notificationStreamReplay.EntityData.YangName = "notification-stream-replay"
    notificationStreamReplay.EntityData.BundleName = "cisco_ios_xe"
    notificationStreamReplay.EntityData.ParentYangName = "callpoints"
    notificationStreamReplay.EntityData.SegmentPath = "notification-stream-replay" + types.AddKeyToken(notificationStreamReplay.Name, "name")
    notificationStreamReplay.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    notificationStreamReplay.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    notificationStreamReplay.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    notificationStreamReplay.EntityData.Children = types.NewOrderedMap()
    notificationStreamReplay.EntityData.Children.Append("daemon", types.YChild{"Daemon", &notificationStreamReplay.Daemon})
    notificationStreamReplay.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range notificationStreamReplay.Range {
        notificationStreamReplay.EntityData.Children.Append(types.GetSegmentPath(notificationStreamReplay.Range[i]), types.YChild{"Range", notificationStreamReplay.Range[i]})
    }
    notificationStreamReplay.EntityData.Leafs = types.NewOrderedMap()
    notificationStreamReplay.EntityData.Leafs.Append("name", types.YLeaf{"Name", notificationStreamReplay.Name})
    notificationStreamReplay.EntityData.Leafs.Append("replay-support", types.YLeaf{"ReplaySupport", notificationStreamReplay.ReplaySupport})
    notificationStreamReplay.EntityData.Leafs.Append("path", types.YLeaf{"Path", notificationStreamReplay.Path})
    notificationStreamReplay.EntityData.Leafs.Append("file", types.YLeaf{"File", notificationStreamReplay.File})
    notificationStreamReplay.EntityData.Leafs.Append("error", types.YLeaf{"Error", notificationStreamReplay.Error})

    notificationStreamReplay.EntityData.YListKeys = []string {"Name"}

    return &(notificationStreamReplay.EntityData)
}

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "notification-stream-replay"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon_Error__PENDING ConfdState_Internal_Callpoints_NotificationStreamReplay_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Range
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "notification-stream-replay"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_NotificationStreamReplay_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_NotificationStreamReplay_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_NotificationStreamReplay_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_NotificationStreamReplay_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_NotificationStreamReplay_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_NotificationStreamReplay_Error__UNKNOWN ConfdState_Internal_Callpoints_NotificationStreamReplay_Error_ = "UNKNOWN"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The type is bool.
    Enabled interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_AuthenticationCallback_Range.
    Range []*ConfdState_Internal_Callpoints_AuthenticationCallback_Range
}

func (authenticationCallback *ConfdState_Internal_Callpoints_AuthenticationCallback) GetEntityData() *types.CommonEntityData {
    authenticationCallback.EntityData.YFilter = authenticationCallback.YFilter
    authenticationCallback.EntityData.YangName = "authentication-callback"
    authenticationCallback.EntityData.BundleName = "cisco_ios_xe"
    authenticationCallback.EntityData.ParentYangName = "callpoints"
    authenticationCallback.EntityData.SegmentPath = "authentication-callback"
    authenticationCallback.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authenticationCallback.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authenticationCallback.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authenticationCallback.EntityData.Children = types.NewOrderedMap()
    authenticationCallback.EntityData.Children.Append("daemon", types.YChild{"Daemon", &authenticationCallback.Daemon})
    authenticationCallback.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range authenticationCallback.Range {
        authenticationCallback.EntityData.Children.Append(types.GetSegmentPath(authenticationCallback.Range[i]), types.YChild{"Range", authenticationCallback.Range[i]})
    }
    authenticationCallback.EntityData.Leafs = types.NewOrderedMap()
    authenticationCallback.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", authenticationCallback.Enabled})
    authenticationCallback.EntityData.Leafs.Append("path", types.YLeaf{"Path", authenticationCallback.Path})
    authenticationCallback.EntityData.Leafs.Append("file", types.YLeaf{"File", authenticationCallback.File})
    authenticationCallback.EntityData.Leafs.Append("error", types.YLeaf{"Error", authenticationCallback.Error})

    authenticationCallback.EntityData.YListKeys = []string {}

    return &(authenticationCallback.EntityData)
}

// ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon
type ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "authentication-callback"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon_Error__PENDING ConfdState_Internal_Callpoints_AuthenticationCallback_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_AuthenticationCallback_Range
type ConfdState_Internal_Callpoints_AuthenticationCallback_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_AuthenticationCallback_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "authentication-callback"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon
type ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_AuthenticationCallback_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_AuthenticationCallback_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_AuthenticationCallback_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_AuthenticationCallback_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_AuthenticationCallback_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_AuthenticationCallback_Error__UNKNOWN ConfdState_Internal_Callpoints_AuthenticationCallback_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Callpoints_AuthorizationCallbacks
// This type is a presence type.
type ConfdState_Internal_Callpoints_AuthorizationCallbacks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The type is bool.
    Enabled interface{}

    // The path of the list that a range registration pertains to. The type is
    // string.
    Path interface{}

    // The pathname of the shared object implementing the type for a typepoint.
    // The type is string.
    File interface{}

    // If this leaf exists, there is a problem with the callpoint registration.
    // The type is Error_.
    Error interface{}

    
    Daemon ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon

    // The type is slice of
    // ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range.
    Range []*ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range
}

func (authorizationCallbacks *ConfdState_Internal_Callpoints_AuthorizationCallbacks) GetEntityData() *types.CommonEntityData {
    authorizationCallbacks.EntityData.YFilter = authorizationCallbacks.YFilter
    authorizationCallbacks.EntityData.YangName = "authorization-callbacks"
    authorizationCallbacks.EntityData.BundleName = "cisco_ios_xe"
    authorizationCallbacks.EntityData.ParentYangName = "callpoints"
    authorizationCallbacks.EntityData.SegmentPath = "authorization-callbacks"
    authorizationCallbacks.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorizationCallbacks.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorizationCallbacks.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorizationCallbacks.EntityData.Children = types.NewOrderedMap()
    authorizationCallbacks.EntityData.Children.Append("daemon", types.YChild{"Daemon", &authorizationCallbacks.Daemon})
    authorizationCallbacks.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range authorizationCallbacks.Range {
        authorizationCallbacks.EntityData.Children.Append(types.GetSegmentPath(authorizationCallbacks.Range[i]), types.YChild{"Range", authorizationCallbacks.Range[i]})
    }
    authorizationCallbacks.EntityData.Leafs = types.NewOrderedMap()
    authorizationCallbacks.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", authorizationCallbacks.Enabled})
    authorizationCallbacks.EntityData.Leafs.Append("path", types.YLeaf{"Path", authorizationCallbacks.Path})
    authorizationCallbacks.EntityData.Leafs.Append("file", types.YLeaf{"File", authorizationCallbacks.File})
    authorizationCallbacks.EntityData.Leafs.Append("error", types.YLeaf{"Error", authorizationCallbacks.Error})

    authorizationCallbacks.EntityData.YListKeys = []string {}

    return &(authorizationCallbacks.EntityData)
}

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "authorization-callbacks"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon_Error__PENDING ConfdState_Internal_Callpoints_AuthorizationCallbacks_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range struct {
    EntityData types.CommonEntityData
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

func (self *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "authorization-callbacks"
    self.EntityData.SegmentPath = "range"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("daemon", types.YChild{"Daemon", &self.Daemon})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("lower", types.YLeaf{"Lower", self.Lower})
    self.EntityData.Leafs.Append("upper", types.YLeaf{"Upper", self.Upper})
    self.EntityData.Leafs.Append("default", types.YLeaf{"Default", self.Default})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The numerical id assigned to the application daemon that has registered for
    // a callpoint. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The name of the application daemon that has registered for a callpoint. The
    // type is string.
    Name interface{}

    // If this leaf exists, there is a problem with the daemon registration. The
    // type is Error_.
    Error interface{}
}

func (daemon *ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon) GetEntityData() *types.CommonEntityData {
    daemon.EntityData.YFilter = daemon.YFilter
    daemon.EntityData.YangName = "daemon"
    daemon.EntityData.BundleName = "cisco_ios_xe"
    daemon.EntityData.ParentYangName = "range"
    daemon.EntityData.SegmentPath = "daemon"
    daemon.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daemon.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daemon.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daemon.EntityData.Children = types.NewOrderedMap()
    daemon.EntityData.Leafs = types.NewOrderedMap()
    daemon.EntityData.Leafs.Append("id", types.YLeaf{"Id", daemon.Id})
    daemon.EntityData.Leafs.Append("name", types.YLeaf{"Name", daemon.Name})
    daemon.EntityData.Leafs.Append("error", types.YLeaf{"Error", daemon.Error})

    daemon.EntityData.YListKeys = []string {}

    return &(daemon.EntityData)
}

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon_Error_ represents with the daemon registration.
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon_Error_ string

const (
    // This value means that the application daemon has not
    // completed the registration (with confd_register_done()).
    ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon_Error__PENDING ConfdState_Internal_Callpoints_AuthorizationCallbacks_Range_Daemon_Error_ = "PENDING"
)

// ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error_ represents with the callpoint registration.
type ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error_ string

const (
    // This value means that there is no registration
    // for the callpoint.
    ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error__NOT_REGISTERED ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error_ = "NOT-REGISTERED"

    // This value means that the callpoint does not exist,
    // but there is a registration for it.
    ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error__UNKNOWN ConfdState_Internal_Callpoints_AuthorizationCallbacks_Error_ = "UNKNOWN"
)

// ConfdState_Internal_Cdb
type ConfdState_Internal_Cdb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of ConfdState_Internal_Cdb_Datastore.
    Datastore []*ConfdState_Internal_Cdb_Datastore

    // The type is slice of ConfdState_Internal_Cdb_Client.
    Client []*ConfdState_Internal_Cdb_Client
}

func (cdb *ConfdState_Internal_Cdb) GetEntityData() *types.CommonEntityData {
    cdb.EntityData.YFilter = cdb.YFilter
    cdb.EntityData.YangName = "cdb"
    cdb.EntityData.BundleName = "cisco_ios_xe"
    cdb.EntityData.ParentYangName = "internal"
    cdb.EntityData.SegmentPath = "cdb"
    cdb.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cdb.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cdb.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cdb.EntityData.Children = types.NewOrderedMap()
    cdb.EntityData.Children.Append("datastore", types.YChild{"Datastore", nil})
    for i := range cdb.Datastore {
        cdb.EntityData.Children.Append(types.GetSegmentPath(cdb.Datastore[i]), types.YChild{"Datastore", cdb.Datastore[i]})
    }
    cdb.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range cdb.Client {
        cdb.EntityData.Children.Append(types.GetSegmentPath(cdb.Client[i]), types.YChild{"Client", cdb.Client[i]})
    }
    cdb.EntityData.Leafs = types.NewOrderedMap()

    cdb.EntityData.YListKeys = []string {}

    return &(cdb.EntityData)
}

// ConfdState_Internal_Cdb_Datastore
type ConfdState_Internal_Cdb_Datastore struct {
    EntityData types.CommonEntityData
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
    PendingNotificationQueue []*ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue
}

func (datastore *ConfdState_Internal_Cdb_Datastore) GetEntityData() *types.CommonEntityData {
    datastore.EntityData.YFilter = datastore.YFilter
    datastore.EntityData.YangName = "datastore"
    datastore.EntityData.BundleName = "cisco_ios_xe"
    datastore.EntityData.ParentYangName = "cdb"
    datastore.EntityData.SegmentPath = "datastore" + types.AddKeyToken(datastore.Name, "name")
    datastore.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    datastore.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    datastore.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    datastore.EntityData.Children = types.NewOrderedMap()
    datastore.EntityData.Children.Append("pending-subscription-sync", types.YChild{"PendingSubscriptionSync", &datastore.PendingSubscriptionSync})
    datastore.EntityData.Children.Append("pending-notification-queue", types.YChild{"PendingNotificationQueue", nil})
    for i := range datastore.PendingNotificationQueue {
        datastore.EntityData.Children.Append(types.GetSegmentPath(datastore.PendingNotificationQueue[i]), types.YChild{"PendingNotificationQueue", datastore.PendingNotificationQueue[i]})
    }
    datastore.EntityData.Leafs = types.NewOrderedMap()
    datastore.EntityData.Leafs.Append("name", types.YLeaf{"Name", datastore.Name})
    datastore.EntityData.Leafs.Append("transaction-id", types.YLeaf{"TransactionId", datastore.TransactionId})
    datastore.EntityData.Leafs.Append("write-queue", types.YLeaf{"WriteQueue", datastore.WriteQueue})
    datastore.EntityData.Leafs.Append("filename", types.YLeaf{"Filename", datastore.Filename})
    datastore.EntityData.Leafs.Append("disk-size", types.YLeaf{"DiskSize", datastore.DiskSize})
    datastore.EntityData.Leafs.Append("ram-size", types.YLeaf{"RamSize", datastore.RamSize})
    datastore.EntityData.Leafs.Append("read-locks", types.YLeaf{"ReadLocks", datastore.ReadLocks})
    datastore.EntityData.Leafs.Append("write-lock-set", types.YLeaf{"WriteLockSet", datastore.WriteLockSet})
    datastore.EntityData.Leafs.Append("subscription-lock-set", types.YLeaf{"SubscriptionLockSet", datastore.SubscriptionLockSet})
    datastore.EntityData.Leafs.Append("waiting-for-replication-sync", types.YLeaf{"WaitingForReplicationSync", datastore.WaitingForReplicationSync})

    datastore.EntityData.YListKeys = []string {"Name"}

    return &(datastore.EntityData)
}

// ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync
// Information pertaining to subscription notifications that have
// been delivered to, but not yet acknowledged by, subscribing
// clients. Not present for the 'startup' datastore.
// This type is a presence type.
type ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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
    Notification []*ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification
}

func (pendingSubscriptionSync *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync) GetEntityData() *types.CommonEntityData {
    pendingSubscriptionSync.EntityData.YFilter = pendingSubscriptionSync.YFilter
    pendingSubscriptionSync.EntityData.YangName = "pending-subscription-sync"
    pendingSubscriptionSync.EntityData.BundleName = "cisco_ios_xe"
    pendingSubscriptionSync.EntityData.ParentYangName = "datastore"
    pendingSubscriptionSync.EntityData.SegmentPath = "pending-subscription-sync"
    pendingSubscriptionSync.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pendingSubscriptionSync.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pendingSubscriptionSync.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pendingSubscriptionSync.EntityData.Children = types.NewOrderedMap()
    pendingSubscriptionSync.EntityData.Children.Append("notification", types.YChild{"Notification", nil})
    for i := range pendingSubscriptionSync.Notification {
        pendingSubscriptionSync.EntityData.Children.Append(types.GetSegmentPath(pendingSubscriptionSync.Notification[i]), types.YChild{"Notification", pendingSubscriptionSync.Notification[i]})
    }
    pendingSubscriptionSync.EntityData.Leafs = types.NewOrderedMap()
    pendingSubscriptionSync.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", pendingSubscriptionSync.Priority})
    pendingSubscriptionSync.EntityData.Leafs.Append("time-remaining", types.YLeaf{"TimeRemaining", pendingSubscriptionSync.TimeRemaining})

    pendingSubscriptionSync.EntityData.YListKeys = []string {}

    return &(pendingSubscriptionSync.EntityData)
}

// ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification
type ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the client that is the recipient of the notification. The type
    // is string.
    ClientName interface{}

    // The subscription identifiers for the subscriptions that generated the
    // notification. The type is slice of interface{} with range: 0..4294967295.
    SubscriptionIds []interface{}
}

func (notification *ConfdState_Internal_Cdb_Datastore_PendingSubscriptionSync_Notification) GetEntityData() *types.CommonEntityData {
    notification.EntityData.YFilter = notification.YFilter
    notification.EntityData.YangName = "notification"
    notification.EntityData.BundleName = "cisco_ios_xe"
    notification.EntityData.ParentYangName = "pending-subscription-sync"
    notification.EntityData.SegmentPath = "notification"
    notification.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    notification.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    notification.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    notification.EntityData.Children = types.NewOrderedMap()
    notification.EntityData.Leafs = types.NewOrderedMap()
    notification.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", notification.ClientName})
    notification.EntityData.Leafs.Append("subscription-ids", types.YLeaf{"SubscriptionIds", notification.SubscriptionIds})

    notification.EntityData.YListKeys = []string {}

    return &(notification.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification.
    Notification []*ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification
}

func (pendingNotificationQueue *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue) GetEntityData() *types.CommonEntityData {
    pendingNotificationQueue.EntityData.YFilter = pendingNotificationQueue.YFilter
    pendingNotificationQueue.EntityData.YangName = "pending-notification-queue"
    pendingNotificationQueue.EntityData.BundleName = "cisco_ios_xe"
    pendingNotificationQueue.EntityData.ParentYangName = "datastore"
    pendingNotificationQueue.EntityData.SegmentPath = "pending-notification-queue"
    pendingNotificationQueue.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pendingNotificationQueue.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pendingNotificationQueue.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pendingNotificationQueue.EntityData.Children = types.NewOrderedMap()
    pendingNotificationQueue.EntityData.Children.Append("notification", types.YChild{"Notification", nil})
    for i := range pendingNotificationQueue.Notification {
        pendingNotificationQueue.EntityData.Children.Append(types.GetSegmentPath(pendingNotificationQueue.Notification[i]), types.YChild{"Notification", pendingNotificationQueue.Notification[i]})
    }
    pendingNotificationQueue.EntityData.Leafs = types.NewOrderedMap()

    pendingNotificationQueue.EntityData.YListKeys = []string {}

    return &(pendingNotificationQueue.EntityData)
}

// ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification
type ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification struct {
    EntityData types.CommonEntityData
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

func (notification *ConfdState_Internal_Cdb_Datastore_PendingNotificationQueue_Notification) GetEntityData() *types.CommonEntityData {
    notification.EntityData.YFilter = notification.YFilter
    notification.EntityData.YangName = "notification"
    notification.EntityData.BundleName = "cisco_ios_xe"
    notification.EntityData.ParentYangName = "pending-notification-queue"
    notification.EntityData.SegmentPath = "notification"
    notification.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    notification.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    notification.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    notification.EntityData.Children = types.NewOrderedMap()
    notification.EntityData.Leafs = types.NewOrderedMap()
    notification.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", notification.Priority})
    notification.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", notification.ClientName})
    notification.EntityData.Leafs.Append("subscription-ids", types.YLeaf{"SubscriptionIds", notification.SubscriptionIds})

    notification.EntityData.YListKeys = []string {}

    return &(notification.EntityData)
}

// ConfdState_Internal_Cdb_Client
type ConfdState_Internal_Cdb_Client struct {
    EntityData types.CommonEntityData
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
    // cdb_wait_start(). The type is Type_.
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
    Subscription []*ConfdState_Internal_Cdb_Client_Subscription
}

func (client *ConfdState_Internal_Cdb_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xe"
    client.EntityData.ParentYangName = "cdb"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("subscription", types.YChild{"Subscription", nil})
    for i := range client.Subscription {
        client.EntityData.Children.Append(types.GetSegmentPath(client.Subscription[i]), types.YChild{"Subscription", client.Subscription[i]})
    }
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("name", types.YLeaf{"Name", client.Name})
    client.EntityData.Leafs.Append("info", types.YLeaf{"Info", client.Info})
    client.EntityData.Leafs.Append("type", types.YLeaf{"Type", client.Type})
    client.EntityData.Leafs.Append("datastore", types.YLeaf{"Datastore", client.Datastore})
    client.EntityData.Leafs.Append("lock", types.YLeaf{"Lock", client.Lock})

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// ConfdState_Internal_Cdb_Client_Subscription
type ConfdState_Internal_Cdb_Client_Subscription struct {
    EntityData types.CommonEntityData
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
    // Error_.
    Error interface{}
}

func (subscription *ConfdState_Internal_Cdb_Client_Subscription) GetEntityData() *types.CommonEntityData {
    subscription.EntityData.YFilter = subscription.YFilter
    subscription.EntityData.YangName = "subscription"
    subscription.EntityData.BundleName = "cisco_ios_xe"
    subscription.EntityData.ParentYangName = "client"
    subscription.EntityData.SegmentPath = "subscription"
    subscription.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    subscription.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    subscription.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    subscription.EntityData.Children = types.NewOrderedMap()
    subscription.EntityData.Leafs = types.NewOrderedMap()
    subscription.EntityData.Leafs.Append("datastore", types.YLeaf{"Datastore", subscription.Datastore})
    subscription.EntityData.Leafs.Append("twophase", types.YLeaf{"Twophase", subscription.Twophase})
    subscription.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", subscription.Priority})
    subscription.EntityData.Leafs.Append("id", types.YLeaf{"Id", subscription.Id})
    subscription.EntityData.Leafs.Append("path", types.YLeaf{"Path", subscription.Path})
    subscription.EntityData.Leafs.Append("error", types.YLeaf{"Error", subscription.Error})

    subscription.EntityData.YListKeys = []string {}

    return &(subscription.EntityData)
}

// ConfdState_Internal_Cdb_Client_Subscription_Error_ represents  with the subscription.
type ConfdState_Internal_Cdb_Client_Subscription_Error_ string

const (
    // This value means that the subscribing client has not
    // completed the subscription (with cdb_subscribe_done()).
    ConfdState_Internal_Cdb_Client_Subscription_Error__PENDING ConfdState_Internal_Cdb_Client_Subscription_Error_ = "PENDING"
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

// ConfdState_Internal_Cdb_Client_Type_ represents waiting for completion of a call such as cdb_wait_start().
type ConfdState_Internal_Cdb_Client_Type_ string

const (
    ConfdState_Internal_Cdb_Client_Type__inactive ConfdState_Internal_Cdb_Client_Type_ = "inactive"

    ConfdState_Internal_Cdb_Client_Type__client ConfdState_Internal_Cdb_Client_Type_ = "client"

    ConfdState_Internal_Cdb_Client_Type__subscriber ConfdState_Internal_Cdb_Client_Type_ = "subscriber"

    ConfdState_Internal_Cdb_Client_Type__waiting ConfdState_Internal_Cdb_Client_Type_ = "waiting"
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

