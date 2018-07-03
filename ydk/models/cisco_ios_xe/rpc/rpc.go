// NED RPC YANG module for IOS
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package rpc

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rpc"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc switch}", reflect.TypeOf(Switch{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:switch", reflect.TypeOf(Switch{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc default}", reflect.TypeOf(Default{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:default", reflect.TypeOf(Default{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc clear}", reflect.TypeOf(Clear{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:clear", reflect.TypeOf(Clear{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc release}", reflect.TypeOf(Release{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:release", reflect.TypeOf(Release{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc reload}", reflect.TypeOf(Reload{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:reload", reflect.TypeOf(Reload{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc cellular}", reflect.TypeOf(Cellular{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:cellular", reflect.TypeOf(Cellular{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc license}", reflect.TypeOf(License{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:license", reflect.TypeOf(License{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc service}", reflect.TypeOf(Service{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:service", reflect.TypeOf(Service{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc virtual-service}", reflect.TypeOf(VirtualService{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:virtual-service", reflect.TypeOf(VirtualService{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc copy}", reflect.TypeOf(Copy{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:copy", reflect.TypeOf(Copy{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc delete}", reflect.TypeOf(Delete{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:delete", reflect.TypeOf(Delete{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc app-hosting}", reflect.TypeOf(AppHosting{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:app-hosting", reflect.TypeOf(AppHosting{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc guestshell}", reflect.TypeOf(Guestshell{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:guestshell", reflect.TypeOf(Guestshell{}))
}

// Switch
type Switch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Switch_Input

    
    Output Switch_Output
}

func (self *Switch) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "switch"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    self.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:switch"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("input", types.YChild{"Input", &self.Input})
    self.EntityData.Children.Append("output", types.YChild{"Output", &self.Output})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Switch_Input
type Switch_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 1..9. This attribute is mandatory.
    _SwitchNumber interface{}

    // <1-15>  Switch Priority. The type is interface{} with range: 1..15.
    Priority interface{}

    // <1-9>  New number of the Switch. The type is interface{} with range: 1..9.
    Renumber interface{}

    
    Statck Switch_Input_Statck
}

func (input *Switch_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "switch"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("statck", types.YChild{"Statck", &input.Statck})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("_switch-number", types.YLeaf{"_SwitchNumber", input._SwitchNumber})
    input.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", input.Priority})
    input.EntityData.Leafs.Append("renumber", types.YLeaf{"Renumber", input.Renumber})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Switch_Input_Statck
type Switch_Input_Statck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // <1-2>  Stack port number to enable/disable. The type is interface{} with
    // range: 1..2.
    Port interface{}
}

func (statck *Switch_Input_Statck) GetEntityData() *types.CommonEntityData {
    statck.EntityData.YFilter = statck.YFilter
    statck.EntityData.YangName = "statck"
    statck.EntityData.BundleName = "cisco_ios_xe"
    statck.EntityData.ParentYangName = "input"
    statck.EntityData.SegmentPath = "statck"
    statck.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    statck.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    statck.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    statck.EntityData.Children = types.NewOrderedMap()
    statck.EntityData.Leafs = types.NewOrderedMap()
    statck.EntityData.Leafs.Append("port", types.YLeaf{"Port", statck.Port})

    statck.EntityData.YListKeys = []string {}

    return &(statck.EntityData)
}

// Switch_Output
type Switch_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Switch_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "switch"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Default
// Set a command to its defaults
type Default struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Default_Input

    
    Output Default_Output
}

func (self *Default) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "default"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    self.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:default"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("input", types.YChild{"Input", &self.Input})
    self.EntityData.Children.Append("output", types.YChild{"Output", &self.Output})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Default_Input
type Default_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select an interface to configure. The type is string with pattern:
    // [A-Za-z]([\w/.-]+).
    Interface interface{}
}

func (input *Default_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "default"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", input.Interface})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Default_Output
type Default_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Default_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "default"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Clear
// Clear
type Clear struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Clear_Input

    
    Output Clear_Output
}

func (clear *Clear) GetEntityData() *types.CommonEntityData {
    clear.EntityData.YFilter = clear.YFilter
    clear.EntityData.YangName = "clear"
    clear.EntityData.BundleName = "cisco_ios_xe"
    clear.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    clear.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:clear"
    clear.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clear.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clear.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clear.EntityData.Children = types.NewOrderedMap()
    clear.EntityData.Children.Append("input", types.YChild{"Input", &clear.Input})
    clear.EntityData.Children.Append("output", types.YChild{"Output", &clear.Output})
    clear.EntityData.Leafs = types.NewOrderedMap()

    clear.EntityData.YListKeys = []string {}

    return &(clear.EntityData)
}

// Clear_Input
type Clear_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select an interface to clear. The type is string with pattern:
    // [A-Za-z]([\w/.-]+).
    Interface interface{}

    // Select an interface to clear. The type is string with pattern:
    // [A-Za-z]([\w/.-]+).
    Count interface{}

    // Flow monitor clear commands.
    Flow Clear_Input_Flow

    
    Ip Clear_Input_Ip

    // Clear the entire ARP cache.
    ArpCache Clear_Input_ArpCache

    // Clear AAA values.
    Aaa Clear_Input_Aaa
}

func (input *Clear_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "clear"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("flow", types.YChild{"Flow", &input.Flow})
    input.EntityData.Children.Append("ip", types.YChild{"Ip", &input.Ip})
    input.EntityData.Children.Append("arp-cache", types.YChild{"ArpCache", &input.ArpCache})
    input.EntityData.Children.Append("aaa", types.YChild{"Aaa", &input.Aaa})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", input.Interface})
    input.EntityData.Leafs.Append("count", types.YLeaf{"Count", input.Count})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Clear_Input_Flow
// Flow monitor clear commands
type Clear_Input_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Monitor Clear_Input_Flow_Monitor

    
    Exporter Clear_Input_Flow_Exporter
}

func (flow *Clear_Input_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xe"
    flow.EntityData.ParentYangName = "input"
    flow.EntityData.SegmentPath = "flow"
    flow.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flow.EntityData.Children = types.NewOrderedMap()
    flow.EntityData.Children.Append("monitor", types.YChild{"Monitor", &flow.Monitor})
    flow.EntityData.Children.Append("exporter", types.YChild{"Exporter", &flow.Exporter})
    flow.EntityData.Leafs = types.NewOrderedMap()

    flow.EntityData.YListKeys = []string {}

    return &(flow.EntityData)
}

// Clear_Input_Flow_Monitor
type Clear_Input_Flow_Monitor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string. This attribute is mandatory.
    Name interface{}

    // The type is interface{}.
    ForceExport interface{}

    // The type is interface{}.
    Statistics interface{}

    
    Cache Clear_Input_Flow_Monitor_Cache
}

func (monitor *Clear_Input_Flow_Monitor) GetEntityData() *types.CommonEntityData {
    monitor.EntityData.YFilter = monitor.YFilter
    monitor.EntityData.YangName = "monitor"
    monitor.EntityData.BundleName = "cisco_ios_xe"
    monitor.EntityData.ParentYangName = "flow"
    monitor.EntityData.SegmentPath = "monitor"
    monitor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    monitor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    monitor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    monitor.EntityData.Children = types.NewOrderedMap()
    monitor.EntityData.Children.Append("cache", types.YChild{"Cache", &monitor.Cache})
    monitor.EntityData.Leafs = types.NewOrderedMap()
    monitor.EntityData.Leafs.Append("name", types.YLeaf{"Name", monitor.Name})
    monitor.EntityData.Leafs.Append("force-export", types.YLeaf{"ForceExport", monitor.ForceExport})
    monitor.EntityData.Leafs.Append("statistics", types.YLeaf{"Statistics", monitor.Statistics})

    monitor.EntityData.YListKeys = []string {}

    return &(monitor.EntityData)
}

// Clear_Input_Flow_Monitor_Cache
type Clear_Input_Flow_Monitor_Cache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    ForceExport interface{}
}

func (cache *Clear_Input_Flow_Monitor_Cache) GetEntityData() *types.CommonEntityData {
    cache.EntityData.YFilter = cache.YFilter
    cache.EntityData.YangName = "cache"
    cache.EntityData.BundleName = "cisco_ios_xe"
    cache.EntityData.ParentYangName = "monitor"
    cache.EntityData.SegmentPath = "cache"
    cache.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cache.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cache.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cache.EntityData.Children = types.NewOrderedMap()
    cache.EntityData.Leafs = types.NewOrderedMap()
    cache.EntityData.Leafs.Append("force-export", types.YLeaf{"ForceExport", cache.ForceExport})

    cache.EntityData.YListKeys = []string {}

    return &(cache.EntityData)
}

// Clear_Input_Flow_Exporter
type Clear_Input_Flow_Exporter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Name interface{}

    // The type is interface{}.
    Statistics interface{}
}

func (exporter *Clear_Input_Flow_Exporter) GetEntityData() *types.CommonEntityData {
    exporter.EntityData.YFilter = exporter.YFilter
    exporter.EntityData.YangName = "exporter"
    exporter.EntityData.BundleName = "cisco_ios_xe"
    exporter.EntityData.ParentYangName = "flow"
    exporter.EntityData.SegmentPath = "exporter"
    exporter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    exporter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    exporter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    exporter.EntityData.Children = types.NewOrderedMap()
    exporter.EntityData.Leafs = types.NewOrderedMap()
    exporter.EntityData.Leafs.Append("name", types.YLeaf{"Name", exporter.Name})
    exporter.EntityData.Leafs.Append("statistics", types.YLeaf{"Statistics", exporter.Statistics})

    exporter.EntityData.YListKeys = []string {}

    return &(exporter.EntityData)
}

// Clear_Input_Ip
type Clear_Input_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delete items from the DHCP database.
    Dhcp Clear_Input_Ip_Dhcp

    
    Ospf Clear_Input_Ip_Ospf

    // Clear BGP connections.
    Bgp Clear_Input_Ip_Bgp
}

func (ip *Clear_Input_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xe"
    ip.EntityData.ParentYangName = "input"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Children.Append("dhcp", types.YChild{"Dhcp", &ip.Dhcp})
    ip.EntityData.Children.Append("ospf", types.YChild{"Ospf", &ip.Ospf})
    ip.EntityData.Children.Append("bgp", types.YChild{"Bgp", &ip.Bgp})
    ip.EntityData.Leafs = types.NewOrderedMap()

    ip.EntityData.YListKeys = []string {}

    return &(ip.EntityData)
}

// Clear_Input_Ip_Dhcp
// Delete items from the DHCP database
type Clear_Input_Ip_Dhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP address bindings.
    Binding Clear_Input_Ip_Dhcp_Binding
}

func (dhcp *Clear_Input_Ip_Dhcp) GetEntityData() *types.CommonEntityData {
    dhcp.EntityData.YFilter = dhcp.YFilter
    dhcp.EntityData.YangName = "dhcp"
    dhcp.EntityData.BundleName = "cisco_ios_xe"
    dhcp.EntityData.ParentYangName = "ip"
    dhcp.EntityData.SegmentPath = "dhcp"
    dhcp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcp.EntityData.Children = types.NewOrderedMap()
    dhcp.EntityData.Children.Append("binding", types.YChild{"Binding", &dhcp.Binding})
    dhcp.EntityData.Leafs = types.NewOrderedMap()

    dhcp.EntityData.YListKeys = []string {}

    return &(dhcp.EntityData)
}

// Clear_Input_Ip_Dhcp_Binding
// DHCP address bindings
type Clear_Input_Ip_Dhcp_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP vrf bindings. The type is string.
    Vrf interface{}

    // Clear all automatic bindings. The type is string with pattern: [*]. This
    // attribute is mandatory.
    _All interface{}
}

func (binding *Clear_Input_Ip_Dhcp_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xe"
    binding.EntityData.ParentYangName = "dhcp"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", binding.Vrf})
    binding.EntityData.Leafs.Append("_all", types.YLeaf{"_All", binding._All})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// Clear_Input_Ip_Ospf
type Clear_Input_Ip_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process ID number. The type is interface{} with range: 0..65535.
    _Id interface{}

    // Reset OSPF process. The type is interface{}.
    Process interface{}
}

func (ospf *Clear_Input_Ip_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xe"
    ospf.EntityData.ParentYangName = "ip"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Leafs = types.NewOrderedMap()
    ospf.EntityData.Leafs.Append("_id", types.YLeaf{"_Id", ospf._Id})
    ospf.EntityData.Leafs.Append("process", types.YLeaf{"Process", ospf.Process})

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Clear_Input_Ip_Bgp
// Clear BGP connections
type Clear_Input_Ip_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP neighbor address to clear. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    _PeerAddress interface{}

    // Select a VPN Routing/Forwarding instance. The type is string.
    Vrf interface{}
}

func (bgp *Clear_Input_Ip_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xe"
    bgp.EntityData.ParentYangName = "ip"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("_peer-address", types.YLeaf{"_PeerAddress", bgp._PeerAddress})
    bgp.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", bgp.Vrf})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Clear_Input_ArpCache
// Clear the entire ARP cache
// This type is a presence type.
type Clear_Input_ArpCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Clear entries for a VPN Routing/Forwarding instance. The type is string.
    Vrf interface{}

    // Clear the entire ARP cache on the interface. The type is string with
    // pattern: [A-Za-z]([\w/.-]+).
    Interface interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    _Ip interface{}
}

func (arpCache *Clear_Input_ArpCache) GetEntityData() *types.CommonEntityData {
    arpCache.EntityData.YFilter = arpCache.YFilter
    arpCache.EntityData.YangName = "arp-cache"
    arpCache.EntityData.BundleName = "cisco_ios_xe"
    arpCache.EntityData.ParentYangName = "input"
    arpCache.EntityData.SegmentPath = "arp-cache"
    arpCache.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    arpCache.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    arpCache.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    arpCache.EntityData.Children = types.NewOrderedMap()
    arpCache.EntityData.Leafs = types.NewOrderedMap()
    arpCache.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", arpCache.Vrf})
    arpCache.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", arpCache.Interface})
    arpCache.EntityData.Leafs.Append("_ip", types.YLeaf{"_Ip", arpCache._Ip})

    arpCache.EntityData.YListKeys = []string {}

    return &(arpCache.EntityData)
}

// Clear_Input_Aaa
// Clear AAA values
type Clear_Input_Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear AAA local method options.
    Local Clear_Input_Aaa_Local
}

func (aaa *Clear_Input_Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xe"
    aaa.EntityData.ParentYangName = "input"
    aaa.EntityData.SegmentPath = "aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Children.Append("local", types.YChild{"Local", &aaa.Local})
    aaa.EntityData.Leafs = types.NewOrderedMap()

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Clear_Input_Aaa_Local
// Clear AAA local method options
type Clear_Input_Aaa_Local struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear local AAA users.
    User Clear_Input_Aaa_Local_User
}

func (local *Clear_Input_Aaa_Local) GetEntityData() *types.CommonEntityData {
    local.EntityData.YFilter = local.YFilter
    local.EntityData.YangName = "local"
    local.EntityData.BundleName = "cisco_ios_xe"
    local.EntityData.ParentYangName = "aaa"
    local.EntityData.SegmentPath = "local"
    local.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    local.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    local.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    local.EntityData.Children = types.NewOrderedMap()
    local.EntityData.Children.Append("user", types.YChild{"User", &local.User})
    local.EntityData.Leafs = types.NewOrderedMap()

    local.EntityData.YListKeys = []string {}

    return &(local.EntityData)
}

// Clear_Input_Aaa_Local_User
// Clear local AAA users
type Clear_Input_Aaa_Local_User struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear locked out local AAA users.
    Lockout Clear_Input_Aaa_Local_User_Lockout
}

func (user *Clear_Input_Aaa_Local_User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xe"
    user.EntityData.ParentYangName = "local"
    user.EntityData.SegmentPath = "user"
    user.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    user.EntityData.Children = types.NewOrderedMap()
    user.EntityData.Children.Append("lockout", types.YChild{"Lockout", &user.Lockout})
    user.EntityData.Leafs = types.NewOrderedMap()

    user.EntityData.YListKeys = []string {}

    return &(user.EntityData)
}

// Clear_Input_Aaa_Local_User_Lockout
// Clear locked out local AAA users
type Clear_Input_Aaa_Local_User_Lockout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Username of the locked-user. The type is string.
    Username interface{}
}

func (lockout *Clear_Input_Aaa_Local_User_Lockout) GetEntityData() *types.CommonEntityData {
    lockout.EntityData.YFilter = lockout.YFilter
    lockout.EntityData.YangName = "lockout"
    lockout.EntityData.BundleName = "cisco_ios_xe"
    lockout.EntityData.ParentYangName = "user"
    lockout.EntityData.SegmentPath = "lockout"
    lockout.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lockout.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lockout.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lockout.EntityData.Children = types.NewOrderedMap()
    lockout.EntityData.Leafs = types.NewOrderedMap()
    lockout.EntityData.Leafs.Append("username", types.YLeaf{"Username", lockout.Username})

    lockout.EntityData.YListKeys = []string {}

    return &(lockout.EntityData)
}

// Clear_Output
type Clear_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Clear_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "clear"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Release
// Release a resource
type Release struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Release_Input

    
    Output Release_Output
}

func (release *Release) GetEntityData() *types.CommonEntityData {
    release.EntityData.YFilter = release.YFilter
    release.EntityData.YangName = "release"
    release.EntityData.BundleName = "cisco_ios_xe"
    release.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    release.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:release"
    release.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    release.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    release.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    release.EntityData.Children = types.NewOrderedMap()
    release.EntityData.Children.Append("input", types.YChild{"Input", &release.Input})
    release.EntityData.Children.Append("output", types.YChild{"Output", &release.Output})
    release.EntityData.Leafs = types.NewOrderedMap()

    release.EntityData.YListKeys = []string {}

    return &(release.EntityData)
}

// Release_Input
type Release_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Release a dhcp lease (an interface). The type is string.
    Dhcp interface{}
}

func (input *Release_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "release"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("dhcp", types.YLeaf{"Dhcp", input.Dhcp})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Release_Output
type Release_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Release_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "release"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Reload
// Halt and perform a cold restart
type Reload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Reload_Input

    
    Output Reload_Output
}

func (reload *Reload) GetEntityData() *types.CommonEntityData {
    reload.EntityData.YFilter = reload.YFilter
    reload.EntityData.YangName = "reload"
    reload.EntityData.BundleName = "cisco_ios_xe"
    reload.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    reload.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:reload"
    reload.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    reload.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    reload.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    reload.EntityData.Children = types.NewOrderedMap()
    reload.EntityData.Children.Append("input", types.YChild{"Input", &reload.Input})
    reload.EntityData.Children.Append("output", types.YChild{"Output", &reload.Output})
    reload.EntityData.Leafs = types.NewOrderedMap()

    reload.EntityData.YListKeys = []string {}

    return &(reload.EntityData)
}

// Reload_Input
type Reload_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Force a restart even if there is unsaved config. The type is bool.
    Force interface{}

    // Reload reason. The type is string.
    Reason interface{}
}

func (input *Reload_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "reload"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("force", types.YLeaf{"Force", input.Force})
    input.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", input.Reason})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Reload_Output
type Reload_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Reload_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "reload"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Cellular
type Cellular struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Cellular_Input

    
    Output Cellular_Output
}

func (cellular *Cellular) GetEntityData() *types.CommonEntityData {
    cellular.EntityData.YFilter = cellular.YFilter
    cellular.EntityData.YangName = "cellular"
    cellular.EntityData.BundleName = "cisco_ios_xe"
    cellular.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    cellular.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:cellular"
    cellular.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cellular.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cellular.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cellular.EntityData.Children = types.NewOrderedMap()
    cellular.EntityData.Children.Append("input", types.YChild{"Input", &cellular.Input})
    cellular.EntityData.Children.Append("output", types.YChild{"Output", &cellular.Output})
    cellular.EntityData.Leafs = types.NewOrderedMap()

    cellular.EntityData.YListKeys = []string {}

    return &(cellular.EntityData)
}

// Cellular_Input
type Cellular_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    _IfName interface{}

    
    Lte Cellular_Input_Lte
}

func (input *Cellular_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "cellular"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("lte", types.YChild{"Lte", &input.Lte})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("_if-name", types.YLeaf{"_IfName", input._IfName})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Cellular_Input_Lte
type Cellular_Input_Lte struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Technology Cellular_Input_Lte_Technology

    
    Profile Cellular_Input_Lte_Profile
}

func (lte *Cellular_Input_Lte) GetEntityData() *types.CommonEntityData {
    lte.EntityData.YFilter = lte.YFilter
    lte.EntityData.YangName = "lte"
    lte.EntityData.BundleName = "cisco_ios_xe"
    lte.EntityData.ParentYangName = "input"
    lte.EntityData.SegmentPath = "lte"
    lte.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lte.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lte.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lte.EntityData.Children = types.NewOrderedMap()
    lte.EntityData.Children.Append("technology", types.YChild{"Technology", &lte.Technology})
    lte.EntityData.Children.Append("profile", types.YChild{"Profile", &lte.Profile})
    lte.EntityData.Leafs = types.NewOrderedMap()

    lte.EntityData.YListKeys = []string {}

    return &(lte.EntityData)
}

// Cellular_Input_Lte_Technology
type Cellular_Input_Lte_Technology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Lte interface{}

    // The type is interface{}.
    Auto interface{}

    // The type is interface{}.
    Umts interface{}
}

func (technology *Cellular_Input_Lte_Technology) GetEntityData() *types.CommonEntityData {
    technology.EntityData.YFilter = technology.YFilter
    technology.EntityData.YangName = "technology"
    technology.EntityData.BundleName = "cisco_ios_xe"
    technology.EntityData.ParentYangName = "lte"
    technology.EntityData.SegmentPath = "technology"
    technology.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    technology.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    technology.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    technology.EntityData.Children = types.NewOrderedMap()
    technology.EntityData.Leafs = types.NewOrderedMap()
    technology.EntityData.Leafs.Append("lte", types.YLeaf{"Lte", technology.Lte})
    technology.EntityData.Leafs.Append("auto", types.YLeaf{"Auto", technology.Auto})
    technology.EntityData.Leafs.Append("umts", types.YLeaf{"Umts", technology.Umts})

    technology.EntityData.YListKeys = []string {}

    return &(technology.EntityData)
}

// Cellular_Input_Lte_Profile
type Cellular_Input_Lte_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delete a Profile.
    Delete Cellular_Input_Lte_Profile_Delete

    // Create a Profile.
    Create Cellular_Input_Lte_Profile_Create
}

func (profile *Cellular_Input_Lte_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xe"
    profile.EntityData.ParentYangName = "lte"
    profile.EntityData.SegmentPath = "profile"
    profile.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("delete", types.YChild{"Delete", &profile.Delete})
    profile.EntityData.Children.Append("create", types.YChild{"Create", &profile.Create})
    profile.EntityData.Leafs = types.NewOrderedMap()

    profile.EntityData.YListKeys = []string {}

    return &(profile.EntityData)
}

// Cellular_Input_Lte_Profile_Delete
// Delete a Profile
type Cellular_Input_Lte_Profile_Delete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile ID. The type is interface{} with range: 1..16.
    _ProfileId interface{}
}

func (self *Cellular_Input_Lte_Profile_Delete) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "delete"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "profile"
    self.EntityData.SegmentPath = "delete"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("_profile-id", types.YLeaf{"_ProfileId", self._ProfileId})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Cellular_Input_Lte_Profile_Create
// Create a Profile
type Cellular_Input_Lte_Profile_Create struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile ID. The type is interface{} with range: 1..16.
    _ProfileId interface{}

    // Profile name. The type is string.
    _ProfName interface{}

    // The type is interface{}.
    None interface{}

    // The type is interface{}.
    Chap interface{}

    // The type is interface{}.
    Pap interface{}

    // The type is interface{}.
    PapChap interface{}

    // Username for authentication. The type is string.
    _UserName interface{}

    // Password for authentication. The type is string.
    _Password interface{}

    // The type is interface{}.
    Ipv4 interface{}

    // The type is interface{}.
    Ipv6 interface{}

    // The type is interface{}.
    Ipv4v6 interface{}
}

func (create *Cellular_Input_Lte_Profile_Create) GetEntityData() *types.CommonEntityData {
    create.EntityData.YFilter = create.YFilter
    create.EntityData.YangName = "create"
    create.EntityData.BundleName = "cisco_ios_xe"
    create.EntityData.ParentYangName = "profile"
    create.EntityData.SegmentPath = "create"
    create.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    create.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    create.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    create.EntityData.Children = types.NewOrderedMap()
    create.EntityData.Leafs = types.NewOrderedMap()
    create.EntityData.Leafs.Append("_profile-id", types.YLeaf{"_ProfileId", create._ProfileId})
    create.EntityData.Leafs.Append("_prof_name", types.YLeaf{"_ProfName", create._ProfName})
    create.EntityData.Leafs.Append("none", types.YLeaf{"None", create.None})
    create.EntityData.Leafs.Append("chap", types.YLeaf{"Chap", create.Chap})
    create.EntityData.Leafs.Append("pap", types.YLeaf{"Pap", create.Pap})
    create.EntityData.Leafs.Append("pap_chap", types.YLeaf{"PapChap", create.PapChap})
    create.EntityData.Leafs.Append("_user_name", types.YLeaf{"_UserName", create._UserName})
    create.EntityData.Leafs.Append("_password", types.YLeaf{"_Password", create._Password})
    create.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", create.Ipv4})
    create.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", create.Ipv6})
    create.EntityData.Leafs.Append("ipv4v6", types.YLeaf{"Ipv4v6", create.Ipv4v6})

    create.EntityData.YListKeys = []string {}

    return &(create.EntityData)
}

// Cellular_Output
type Cellular_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Cellular_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "cellular"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// License
type License struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input License_Input

    
    Output License_Output
}

func (license *License) GetEntityData() *types.CommonEntityData {
    license.EntityData.YFilter = license.YFilter
    license.EntityData.YangName = "license"
    license.EntityData.BundleName = "cisco_ios_xe"
    license.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    license.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:license"
    license.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    license.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    license.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    license.EntityData.Children = types.NewOrderedMap()
    license.EntityData.Children.Append("input", types.YChild{"Input", &license.Input})
    license.EntityData.Children.Append("output", types.YChild{"Output", &license.Output})
    license.EntityData.Leafs = types.NewOrderedMap()

    license.EntityData.YListKeys = []string {}

    return &(license.EntityData)
}

// License_Input
type License_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Smart License_Input_Smart
}

func (input *License_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "license"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("smart", types.YChild{"Smart", &input.Smart})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// License_Input_Smart
type License_Input_Smart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Deregister interface{}

    
    Register License_Input_Smart_Register

    
    Renew License_Input_Smart_Renew
}

func (smart *License_Input_Smart) GetEntityData() *types.CommonEntityData {
    smart.EntityData.YFilter = smart.YFilter
    smart.EntityData.YangName = "smart"
    smart.EntityData.BundleName = "cisco_ios_xe"
    smart.EntityData.ParentYangName = "input"
    smart.EntityData.SegmentPath = "smart"
    smart.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    smart.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    smart.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    smart.EntityData.Children = types.NewOrderedMap()
    smart.EntityData.Children.Append("register", types.YChild{"Register", &smart.Register})
    smart.EntityData.Children.Append("renew", types.YChild{"Renew", &smart.Renew})
    smart.EntityData.Leafs = types.NewOrderedMap()
    smart.EntityData.Leafs.Append("deregister", types.YLeaf{"Deregister", smart.Deregister})

    smart.EntityData.YListKeys = []string {}

    return &(smart.EntityData)
}

// License_Input_Smart_Register
type License_Input_Smart_Register struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string. This attribute is mandatory.
    Idtoken interface{}

    // Forcefully register. The type is interface{}.
    Force interface{}
}

func (register *License_Input_Smart_Register) GetEntityData() *types.CommonEntityData {
    register.EntityData.YFilter = register.YFilter
    register.EntityData.YangName = "register"
    register.EntityData.BundleName = "cisco_ios_xe"
    register.EntityData.ParentYangName = "smart"
    register.EntityData.SegmentPath = "register"
    register.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    register.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    register.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    register.EntityData.Children = types.NewOrderedMap()
    register.EntityData.Leafs = types.NewOrderedMap()
    register.EntityData.Leafs.Append("idtoken", types.YLeaf{"Idtoken", register.Idtoken})
    register.EntityData.Leafs.Append("force", types.YLeaf{"Force", register.Force})

    register.EntityData.YListKeys = []string {}

    return &(register.EntityData)
}

// License_Input_Smart_Renew
type License_Input_Smart_Renew struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    ID interface{}

    // The type is interface{}.
    Auth interface{}
}

func (renew *License_Input_Smart_Renew) GetEntityData() *types.CommonEntityData {
    renew.EntityData.YFilter = renew.YFilter
    renew.EntityData.YangName = "renew"
    renew.EntityData.BundleName = "cisco_ios_xe"
    renew.EntityData.ParentYangName = "smart"
    renew.EntityData.SegmentPath = "renew"
    renew.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    renew.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    renew.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    renew.EntityData.Children = types.NewOrderedMap()
    renew.EntityData.Leafs = types.NewOrderedMap()
    renew.EntityData.Leafs.Append("ID", types.YLeaf{"ID", renew.ID})
    renew.EntityData.Leafs.Append("auth", types.YLeaf{"Auth", renew.Auth})

    renew.EntityData.YListKeys = []string {}

    return &(renew.EntityData)
}

// License_Output
type License_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *License_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "license"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Service
// SD-AVC service management
type Service struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Service_Input

    
    Output Service_Output
}

func (service *Service) GetEntityData() *types.CommonEntityData {
    service.EntityData.YFilter = service.YFilter
    service.EntityData.YangName = "service"
    service.EntityData.BundleName = "cisco_ios_xe"
    service.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    service.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:service"
    service.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    service.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    service.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    service.EntityData.Children = types.NewOrderedMap()
    service.EntityData.Children.Append("input", types.YChild{"Input", &service.Input})
    service.EntityData.Children.Append("output", types.YChild{"Output", &service.Output})
    service.EntityData.Leafs = types.NewOrderedMap()

    service.EntityData.YListKeys = []string {}

    return &(service.EntityData)
}

// Service_Input
type Service_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    SdAvc Service_Input_SdAvc
}

func (input *Service_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "service"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("sd-avc", types.YChild{"SdAvc", &input.SdAvc})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Service_Input_SdAvc
type Service_Input_SdAvc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Activate interface{}

    // The type is interface{}.
    Connect interface{}

    // The type is interface{}.
    Help interface{}

    // The type is interface{}.
    Deactivate interface{}

    // The type is interface{}.
    Status interface{}

    // The type is interface{}.
    Unconfigure interface{}

    // The type is interface{}.
    Uninstall interface{}

    
    Configure Service_Input_SdAvc_Configure

    
    Install Service_Input_SdAvc_Install

    
    Upgrade Service_Input_SdAvc_Upgrade
}

func (sdAvc *Service_Input_SdAvc) GetEntityData() *types.CommonEntityData {
    sdAvc.EntityData.YFilter = sdAvc.YFilter
    sdAvc.EntityData.YangName = "sd-avc"
    sdAvc.EntityData.BundleName = "cisco_ios_xe"
    sdAvc.EntityData.ParentYangName = "input"
    sdAvc.EntityData.SegmentPath = "sd-avc"
    sdAvc.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sdAvc.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sdAvc.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sdAvc.EntityData.Children = types.NewOrderedMap()
    sdAvc.EntityData.Children.Append("configure", types.YChild{"Configure", &sdAvc.Configure})
    sdAvc.EntityData.Children.Append("install", types.YChild{"Install", &sdAvc.Install})
    sdAvc.EntityData.Children.Append("upgrade", types.YChild{"Upgrade", &sdAvc.Upgrade})
    sdAvc.EntityData.Leafs = types.NewOrderedMap()
    sdAvc.EntityData.Leafs.Append("activate", types.YLeaf{"Activate", sdAvc.Activate})
    sdAvc.EntityData.Leafs.Append("connect", types.YLeaf{"Connect", sdAvc.Connect})
    sdAvc.EntityData.Leafs.Append("help", types.YLeaf{"Help", sdAvc.Help})
    sdAvc.EntityData.Leafs.Append("deactivate", types.YLeaf{"Deactivate", sdAvc.Deactivate})
    sdAvc.EntityData.Leafs.Append("status", types.YLeaf{"Status", sdAvc.Status})
    sdAvc.EntityData.Leafs.Append("unconfigure", types.YLeaf{"Unconfigure", sdAvc.Unconfigure})
    sdAvc.EntityData.Leafs.Append("uninstall", types.YLeaf{"Uninstall", sdAvc.Uninstall})

    sdAvc.EntityData.YListKeys = []string {}

    return &(sdAvc.EntityData)
}

// Service_Input_SdAvc_Configure
type Service_Input_SdAvc_Configure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Gateway Service_Input_SdAvc_Configure_Gateway
}

func (configure *Service_Input_SdAvc_Configure) GetEntityData() *types.CommonEntityData {
    configure.EntityData.YFilter = configure.YFilter
    configure.EntityData.YangName = "configure"
    configure.EntityData.BundleName = "cisco_ios_xe"
    configure.EntityData.ParentYangName = "sd-avc"
    configure.EntityData.SegmentPath = "configure"
    configure.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    configure.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    configure.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    configure.EntityData.Children = types.NewOrderedMap()
    configure.EntityData.Children.Append("gateway", types.YChild{"Gateway", &configure.Gateway})
    configure.EntityData.Leafs = types.NewOrderedMap()

    configure.EntityData.YListKeys = []string {}

    return &(configure.EntityData)
}

// Service_Input_SdAvc_Configure_Gateway
type Service_Input_SdAvc_Configure_Gateway struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Interface interface{}

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServiceIp interface{}

    // The type is interface{}.
    Activate interface{}
}

func (gateway *Service_Input_SdAvc_Configure_Gateway) GetEntityData() *types.CommonEntityData {
    gateway.EntityData.YFilter = gateway.YFilter
    gateway.EntityData.YangName = "gateway"
    gateway.EntityData.BundleName = "cisco_ios_xe"
    gateway.EntityData.ParentYangName = "configure"
    gateway.EntityData.SegmentPath = "gateway"
    gateway.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    gateway.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    gateway.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    gateway.EntityData.Children = types.NewOrderedMap()
    gateway.EntityData.Leafs = types.NewOrderedMap()
    gateway.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", gateway.Interface})
    gateway.EntityData.Leafs.Append("service-ip", types.YLeaf{"ServiceIp", gateway.ServiceIp})
    gateway.EntityData.Leafs.Append("activate", types.YLeaf{"Activate", gateway.Activate})

    gateway.EntityData.YListKeys = []string {}

    return &(gateway.EntityData)
}

// Service_Input_SdAvc_Install
type Service_Input_SdAvc_Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Package interface{}
}

func (install *Service_Input_SdAvc_Install) GetEntityData() *types.CommonEntityData {
    install.EntityData.YFilter = install.YFilter
    install.EntityData.YangName = "install"
    install.EntityData.BundleName = "cisco_ios_xe"
    install.EntityData.ParentYangName = "sd-avc"
    install.EntityData.SegmentPath = "install"
    install.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    install.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    install.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    install.EntityData.Children = types.NewOrderedMap()
    install.EntityData.Leafs = types.NewOrderedMap()
    install.EntityData.Leafs.Append("package", types.YLeaf{"Package", install.Package})

    install.EntityData.YListKeys = []string {}

    return &(install.EntityData)
}

// Service_Input_SdAvc_Upgrade
type Service_Input_SdAvc_Upgrade struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Package interface{}
}

func (upgrade *Service_Input_SdAvc_Upgrade) GetEntityData() *types.CommonEntityData {
    upgrade.EntityData.YFilter = upgrade.YFilter
    upgrade.EntityData.YangName = "upgrade"
    upgrade.EntityData.BundleName = "cisco_ios_xe"
    upgrade.EntityData.ParentYangName = "sd-avc"
    upgrade.EntityData.SegmentPath = "upgrade"
    upgrade.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    upgrade.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    upgrade.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    upgrade.EntityData.Children = types.NewOrderedMap()
    upgrade.EntityData.Leafs = types.NewOrderedMap()
    upgrade.EntityData.Leafs.Append("package", types.YLeaf{"Package", upgrade.Package})

    upgrade.EntityData.YListKeys = []string {}

    return &(upgrade.EntityData)
}

// Service_Output
type Service_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Service_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "service"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// VirtualService
// Virtual-service management
type VirtualService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input VirtualService_Input

    
    Output VirtualService_Output
}

func (virtualService *VirtualService) GetEntityData() *types.CommonEntityData {
    virtualService.EntityData.YFilter = virtualService.YFilter
    virtualService.EntityData.YangName = "virtual-service"
    virtualService.EntityData.BundleName = "cisco_ios_xe"
    virtualService.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    virtualService.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:virtual-service"
    virtualService.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    virtualService.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    virtualService.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    virtualService.EntityData.Children = types.NewOrderedMap()
    virtualService.EntityData.Children.Append("input", types.YChild{"Input", &virtualService.Input})
    virtualService.EntityData.Children.Append("output", types.YChild{"Output", &virtualService.Output})
    virtualService.EntityData.Leafs = types.NewOrderedMap()

    virtualService.EntityData.YListKeys = []string {}

    return &(virtualService.EntityData)
}

// VirtualService_Input
type VirtualService_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Install VirtualService_Input_Install

    
    Uninstall VirtualService_Input_Uninstall

    
    Upgrade VirtualService_Input_Upgrade
}

func (input *VirtualService_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "virtual-service"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("install", types.YChild{"Install", &input.Install})
    input.EntityData.Children.Append("uninstall", types.YChild{"Uninstall", &input.Uninstall})
    input.EntityData.Children.Append("upgrade", types.YChild{"Upgrade", &input.Upgrade})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// VirtualService_Input_Install
type VirtualService_Input_Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Name interface{}

    // The type is string.
    Package interface{}

    // The type is string.
    Media interface{}
}

func (install *VirtualService_Input_Install) GetEntityData() *types.CommonEntityData {
    install.EntityData.YFilter = install.YFilter
    install.EntityData.YangName = "install"
    install.EntityData.BundleName = "cisco_ios_xe"
    install.EntityData.ParentYangName = "input"
    install.EntityData.SegmentPath = "install"
    install.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    install.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    install.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    install.EntityData.Children = types.NewOrderedMap()
    install.EntityData.Leafs = types.NewOrderedMap()
    install.EntityData.Leafs.Append("name", types.YLeaf{"Name", install.Name})
    install.EntityData.Leafs.Append("package", types.YLeaf{"Package", install.Package})
    install.EntityData.Leafs.Append("media", types.YLeaf{"Media", install.Media})

    install.EntityData.YListKeys = []string {}

    return &(install.EntityData)
}

// VirtualService_Input_Uninstall
type VirtualService_Input_Uninstall struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Name interface{}
}

func (uninstall *VirtualService_Input_Uninstall) GetEntityData() *types.CommonEntityData {
    uninstall.EntityData.YFilter = uninstall.YFilter
    uninstall.EntityData.YangName = "uninstall"
    uninstall.EntityData.BundleName = "cisco_ios_xe"
    uninstall.EntityData.ParentYangName = "input"
    uninstall.EntityData.SegmentPath = "uninstall"
    uninstall.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    uninstall.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    uninstall.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    uninstall.EntityData.Children = types.NewOrderedMap()
    uninstall.EntityData.Leafs = types.NewOrderedMap()
    uninstall.EntityData.Leafs.Append("name", types.YLeaf{"Name", uninstall.Name})

    uninstall.EntityData.YListKeys = []string {}

    return &(uninstall.EntityData)
}

// VirtualService_Input_Upgrade
type VirtualService_Input_Upgrade struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Name interface{}

    // The type is string.
    Package interface{}
}

func (upgrade *VirtualService_Input_Upgrade) GetEntityData() *types.CommonEntityData {
    upgrade.EntityData.YFilter = upgrade.YFilter
    upgrade.EntityData.YangName = "upgrade"
    upgrade.EntityData.BundleName = "cisco_ios_xe"
    upgrade.EntityData.ParentYangName = "input"
    upgrade.EntityData.SegmentPath = "upgrade"
    upgrade.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    upgrade.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    upgrade.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    upgrade.EntityData.Children = types.NewOrderedMap()
    upgrade.EntityData.Leafs = types.NewOrderedMap()
    upgrade.EntityData.Leafs.Append("name", types.YLeaf{"Name", upgrade.Name})
    upgrade.EntityData.Leafs.Append("package", types.YLeaf{"Package", upgrade.Package})

    upgrade.EntityData.YListKeys = []string {}

    return &(upgrade.EntityData)
}

// VirtualService_Output
type VirtualService_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *VirtualService_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "virtual-service"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Copy
// Copy from one file to another
type Copy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Copy_Input

    
    Output Copy_Output
}

func (self *Copy) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "copy"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    self.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:copy"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("input", types.YChild{"Input", &self.Input})
    self.EntityData.Children.Append("output", types.YChild{"Output", &self.Output})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Copy_Input
type Copy_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // ((((bootflash:)|(harddisk:)|(flash:)|(nvram:)|(ftp:)|(http:)|(https:)|(scp:)|(tftp:)).*)|((running-config)|(startup-config))).
    // This attribute is mandatory.
    _Source interface{}

    // The type is string with pattern:
    // ((((bootflash:)|(harddisk:)|(flash:)|(nvram:)|(ftp:)|(http:)|(https:)|(scp:)|(tftp:)).*)|((running-config)|(startup-config))).
    // This attribute is mandatory.
    _Destination interface{}
}

func (input *Copy_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "copy"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("_source", types.YLeaf{"_Source", input._Source})
    input.EntityData.Leafs.Append("_destination", types.YLeaf{"_Destination", input._Destination})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Copy_Output
type Copy_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Copy_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "copy"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Delete
// Delete a file
type Delete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Delete_Input

    
    Output Delete_Output
}

func (self *Delete) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "delete"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    self.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:delete"
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("input", types.YChild{"Input", &self.Input})
    self.EntityData.Children.Append("output", types.YChild{"Output", &self.Output})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Delete_Input
type Delete_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // (((bootflash:)|(harddisk:)|(flash:)|(nvram:)).*). This attribute is
    // mandatory.
    _Filename interface{}
}

func (input *Delete_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "delete"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("_filename", types.YLeaf{"_Filename", input._Filename})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Delete_Output
type Delete_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Delete_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "delete"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// AppHosting
// App-hosting management
type AppHosting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input AppHosting_Input

    
    Output AppHosting_Output
}

func (appHosting *AppHosting) GetEntityData() *types.CommonEntityData {
    appHosting.EntityData.YFilter = appHosting.YFilter
    appHosting.EntityData.YangName = "app-hosting"
    appHosting.EntityData.BundleName = "cisco_ios_xe"
    appHosting.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    appHosting.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:app-hosting"
    appHosting.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    appHosting.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    appHosting.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    appHosting.EntityData.Children = types.NewOrderedMap()
    appHosting.EntityData.Children.Append("input", types.YChild{"Input", &appHosting.Input})
    appHosting.EntityData.Children.Append("output", types.YChild{"Output", &appHosting.Output})
    appHosting.EntityData.Leafs = types.NewOrderedMap()

    appHosting.EntityData.YListKeys = []string {}

    return &(appHosting.EntityData)
}

// AppHosting_Input
type AppHosting_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Install AppHosting_Input_Install

    
    Uninstall AppHosting_Input_Uninstall

    
    Activate AppHosting_Input_Activate

    
    Deactivate AppHosting_Input_Deactivate

    
    Start AppHosting_Input_Start

    
    Stop AppHosting_Input_Stop
}

func (input *AppHosting_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "app-hosting"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("install", types.YChild{"Install", &input.Install})
    input.EntityData.Children.Append("uninstall", types.YChild{"Uninstall", &input.Uninstall})
    input.EntityData.Children.Append("activate", types.YChild{"Activate", &input.Activate})
    input.EntityData.Children.Append("deactivate", types.YChild{"Deactivate", &input.Deactivate})
    input.EntityData.Children.Append("start", types.YChild{"Start", &input.Start})
    input.EntityData.Children.Append("stop", types.YChild{"Stop", &input.Stop})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// AppHosting_Input_Install
type AppHosting_Input_Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Appid interface{}

    // The type is string.
    Package interface{}
}

func (install *AppHosting_Input_Install) GetEntityData() *types.CommonEntityData {
    install.EntityData.YFilter = install.YFilter
    install.EntityData.YangName = "install"
    install.EntityData.BundleName = "cisco_ios_xe"
    install.EntityData.ParentYangName = "input"
    install.EntityData.SegmentPath = "install"
    install.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    install.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    install.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    install.EntityData.Children = types.NewOrderedMap()
    install.EntityData.Leafs = types.NewOrderedMap()
    install.EntityData.Leafs.Append("appid", types.YLeaf{"Appid", install.Appid})
    install.EntityData.Leafs.Append("package", types.YLeaf{"Package", install.Package})

    install.EntityData.YListKeys = []string {}

    return &(install.EntityData)
}

// AppHosting_Input_Uninstall
type AppHosting_Input_Uninstall struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Appid interface{}
}

func (uninstall *AppHosting_Input_Uninstall) GetEntityData() *types.CommonEntityData {
    uninstall.EntityData.YFilter = uninstall.YFilter
    uninstall.EntityData.YangName = "uninstall"
    uninstall.EntityData.BundleName = "cisco_ios_xe"
    uninstall.EntityData.ParentYangName = "input"
    uninstall.EntityData.SegmentPath = "uninstall"
    uninstall.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    uninstall.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    uninstall.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    uninstall.EntityData.Children = types.NewOrderedMap()
    uninstall.EntityData.Leafs = types.NewOrderedMap()
    uninstall.EntityData.Leafs.Append("appid", types.YLeaf{"Appid", uninstall.Appid})

    uninstall.EntityData.YListKeys = []string {}

    return &(uninstall.EntityData)
}

// AppHosting_Input_Activate
type AppHosting_Input_Activate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Appid interface{}
}

func (activate *AppHosting_Input_Activate) GetEntityData() *types.CommonEntityData {
    activate.EntityData.YFilter = activate.YFilter
    activate.EntityData.YangName = "activate"
    activate.EntityData.BundleName = "cisco_ios_xe"
    activate.EntityData.ParentYangName = "input"
    activate.EntityData.SegmentPath = "activate"
    activate.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    activate.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    activate.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    activate.EntityData.Children = types.NewOrderedMap()
    activate.EntityData.Leafs = types.NewOrderedMap()
    activate.EntityData.Leafs.Append("appid", types.YLeaf{"Appid", activate.Appid})

    activate.EntityData.YListKeys = []string {}

    return &(activate.EntityData)
}

// AppHosting_Input_Deactivate
type AppHosting_Input_Deactivate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Appid interface{}
}

func (deactivate *AppHosting_Input_Deactivate) GetEntityData() *types.CommonEntityData {
    deactivate.EntityData.YFilter = deactivate.YFilter
    deactivate.EntityData.YangName = "deactivate"
    deactivate.EntityData.BundleName = "cisco_ios_xe"
    deactivate.EntityData.ParentYangName = "input"
    deactivate.EntityData.SegmentPath = "deactivate"
    deactivate.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deactivate.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deactivate.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deactivate.EntityData.Children = types.NewOrderedMap()
    deactivate.EntityData.Leafs = types.NewOrderedMap()
    deactivate.EntityData.Leafs.Append("appid", types.YLeaf{"Appid", deactivate.Appid})

    deactivate.EntityData.YListKeys = []string {}

    return &(deactivate.EntityData)
}

// AppHosting_Input_Start
type AppHosting_Input_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Appid interface{}
}

func (start *AppHosting_Input_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xe"
    start.EntityData.ParentYangName = "input"
    start.EntityData.SegmentPath = "start"
    start.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    start.EntityData.Children = types.NewOrderedMap()
    start.EntityData.Leafs = types.NewOrderedMap()
    start.EntityData.Leafs.Append("appid", types.YLeaf{"Appid", start.Appid})

    start.EntityData.YListKeys = []string {}

    return &(start.EntityData)
}

// AppHosting_Input_Stop
type AppHosting_Input_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Appid interface{}
}

func (stop *AppHosting_Input_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xe"
    stop.EntityData.ParentYangName = "input"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stop.EntityData.Children = types.NewOrderedMap()
    stop.EntityData.Leafs = types.NewOrderedMap()
    stop.EntityData.Leafs.Append("appid", types.YLeaf{"Appid", stop.Appid})

    stop.EntityData.YListKeys = []string {}

    return &(stop.EntityData)
}

// AppHosting_Output
type AppHosting_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *AppHosting_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "app-hosting"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Guestshell
// guestshell managementshell
type Guestshell struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Guestshell_Input

    
    Output Guestshell_Output
}

func (guestshell *Guestshell) GetEntityData() *types.CommonEntityData {
    guestshell.EntityData.YFilter = guestshell.YFilter
    guestshell.EntityData.YangName = "guestshell"
    guestshell.EntityData.BundleName = "cisco_ios_xe"
    guestshell.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    guestshell.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:guestshell"
    guestshell.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    guestshell.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    guestshell.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    guestshell.EntityData.Children = types.NewOrderedMap()
    guestshell.EntityData.Children.Append("input", types.YChild{"Input", &guestshell.Input})
    guestshell.EntityData.Children.Append("output", types.YChild{"Output", &guestshell.Output})
    guestshell.EntityData.Leafs = types.NewOrderedMap()

    guestshell.EntityData.YListKeys = []string {}

    return &(guestshell.EntityData)
}

// Guestshell_Input
type Guestshell_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destroy guestshell. The type is interface{}.
    Destroy interface{}

    // Disable guestshell. The type is interface{}.
    Disable interface{}

    // Enable guestshell. The type is interface{}.
    Enable interface{}
}

func (input *Guestshell_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "guestshell"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("destroy", types.YLeaf{"Destroy", input.Destroy})
    input.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", input.Disable})
    input.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", input.Enable})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Guestshell_Output
type Guestshell_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Guestshell_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "guestshell"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("result", types.YLeaf{"Result", output.Result})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

