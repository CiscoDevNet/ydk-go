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

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["input"] = types.YChild{"Input", &self.Input}
    self.EntityData.Children["output"] = types.YChild{"Output", &self.Output}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Switch_Input
type Switch_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 1..9. This attribute is mandatory.
    YSwitchNumber interface{}

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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["statck"] = types.YChild{"Statck", &input.Statck}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["_switch-number"] = types.YLeaf{"YSwitchNumber", input.YSwitchNumber}
    input.EntityData.Leafs["priority"] = types.YLeaf{"Priority", input.Priority}
    input.EntityData.Leafs["renumber"] = types.YLeaf{"Renumber", input.Renumber}
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

    statck.EntityData.Children = make(map[string]types.YChild)
    statck.EntityData.Leafs = make(map[string]types.YLeaf)
    statck.EntityData.Leafs["port"] = types.YLeaf{"Port", statck.Port}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["input"] = types.YChild{"Input", &self.Input}
    self.EntityData.Children["output"] = types.YChild{"Output", &self.Output}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Default_Input
type Default_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select an interface to configure. The type is string with pattern:
    // b'[A-Za-z]([\\w/.-]+)'.
    Interface_ interface{}
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", input.Interface_}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    clear.EntityData.Children = make(map[string]types.YChild)
    clear.EntityData.Children["input"] = types.YChild{"Input", &clear.Input}
    clear.EntityData.Children["output"] = types.YChild{"Output", &clear.Output}
    clear.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clear.EntityData)
}

// Clear_Input
type Clear_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select an interface to clear. The type is string with pattern:
    // b'[A-Za-z]([\\w/.-]+)'.
    Interface_ interface{}

    // Select an interface to clear. The type is string with pattern:
    // b'[A-Za-z]([\\w/.-]+)'.
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["flow"] = types.YChild{"Flow", &input.Flow}
    input.EntityData.Children["ip"] = types.YChild{"Ip", &input.Ip}
    input.EntityData.Children["arp-cache"] = types.YChild{"ArpCache", &input.ArpCache}
    input.EntityData.Children["aaa"] = types.YChild{"Aaa", &input.Aaa}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", input.Interface_}
    input.EntityData.Leafs["count"] = types.YLeaf{"Count", input.Count}
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

    flow.EntityData.Children = make(map[string]types.YChild)
    flow.EntityData.Children["monitor"] = types.YChild{"Monitor", &flow.Monitor}
    flow.EntityData.Children["exporter"] = types.YChild{"Exporter", &flow.Exporter}
    flow.EntityData.Leafs = make(map[string]types.YLeaf)
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

    monitor.EntityData.Children = make(map[string]types.YChild)
    monitor.EntityData.Children["cache"] = types.YChild{"Cache", &monitor.Cache}
    monitor.EntityData.Leafs = make(map[string]types.YLeaf)
    monitor.EntityData.Leafs["name"] = types.YLeaf{"Name", monitor.Name}
    monitor.EntityData.Leafs["force-export"] = types.YLeaf{"ForceExport", monitor.ForceExport}
    monitor.EntityData.Leafs["statistics"] = types.YLeaf{"Statistics", monitor.Statistics}
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

    cache.EntityData.Children = make(map[string]types.YChild)
    cache.EntityData.Leafs = make(map[string]types.YLeaf)
    cache.EntityData.Leafs["force-export"] = types.YLeaf{"ForceExport", cache.ForceExport}
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

    exporter.EntityData.Children = make(map[string]types.YChild)
    exporter.EntityData.Leafs = make(map[string]types.YLeaf)
    exporter.EntityData.Leafs["name"] = types.YLeaf{"Name", exporter.Name}
    exporter.EntityData.Leafs["statistics"] = types.YLeaf{"Statistics", exporter.Statistics}
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

    ip.EntityData.Children = make(map[string]types.YChild)
    ip.EntityData.Children["dhcp"] = types.YChild{"Dhcp", &ip.Dhcp}
    ip.EntityData.Children["ospf"] = types.YChild{"Ospf", &ip.Ospf}
    ip.EntityData.Children["bgp"] = types.YChild{"Bgp", &ip.Bgp}
    ip.EntityData.Leafs = make(map[string]types.YLeaf)
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

    dhcp.EntityData.Children = make(map[string]types.YChild)
    dhcp.EntityData.Children["binding"] = types.YChild{"Binding", &dhcp.Binding}
    dhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dhcp.EntityData)
}

// Clear_Input_Ip_Dhcp_Binding
// DHCP address bindings
type Clear_Input_Ip_Dhcp_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP vrf bindings. The type is string.
    Vrf interface{}

    // Clear all automatic bindings. The type is string with pattern: b'[*]'. This
    // attribute is mandatory.
    YAll interface{}
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

    binding.EntityData.Children = make(map[string]types.YChild)
    binding.EntityData.Leafs = make(map[string]types.YLeaf)
    binding.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", binding.Vrf}
    binding.EntityData.Leafs["_all"] = types.YLeaf{"YAll", binding.YAll}
    return &(binding.EntityData)
}

// Clear_Input_Ip_Ospf
type Clear_Input_Ip_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process ID number. The type is interface{} with range: 0..65535.
    YId interface{}

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

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["_id"] = types.YLeaf{"YId", ospf.YId}
    ospf.EntityData.Leafs["process"] = types.YLeaf{"Process", ospf.Process}
    return &(ospf.EntityData)
}

// Clear_Input_Ip_Bgp
// Clear BGP connections
type Clear_Input_Ip_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP neighbor address to clear. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    YPeerAddress interface{}

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

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["_peer-address"] = types.YLeaf{"YPeerAddress", bgp.YPeerAddress}
    bgp.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", bgp.Vrf}
    return &(bgp.EntityData)
}

// Clear_Input_ArpCache
// Clear the entire ARP cache
// This type is a presence type.
type Clear_Input_ArpCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear entries for a VPN Routing/Forwarding instance. The type is string.
    Vrf interface{}

    // Clear the entire ARP cache on the interface. The type is string with
    // pattern: b'[A-Za-z]([\\w/.-]+)'.
    Interface_ interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    YIp interface{}
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

    arpCache.EntityData.Children = make(map[string]types.YChild)
    arpCache.EntityData.Leafs = make(map[string]types.YLeaf)
    arpCache.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", arpCache.Vrf}
    arpCache.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", arpCache.Interface_}
    arpCache.EntityData.Leafs["_ip"] = types.YLeaf{"YIp", arpCache.YIp}
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

    aaa.EntityData.Children = make(map[string]types.YChild)
    aaa.EntityData.Children["local"] = types.YChild{"Local", &aaa.Local}
    aaa.EntityData.Leafs = make(map[string]types.YLeaf)
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

    local.EntityData.Children = make(map[string]types.YChild)
    local.EntityData.Children["user"] = types.YChild{"User", &local.User}
    local.EntityData.Leafs = make(map[string]types.YLeaf)
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

    user.EntityData.Children = make(map[string]types.YChild)
    user.EntityData.Children["lockout"] = types.YChild{"Lockout", &user.Lockout}
    user.EntityData.Leafs = make(map[string]types.YLeaf)
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

    lockout.EntityData.Children = make(map[string]types.YChild)
    lockout.EntityData.Leafs = make(map[string]types.YLeaf)
    lockout.EntityData.Leafs["username"] = types.YLeaf{"Username", lockout.Username}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    release.EntityData.Children = make(map[string]types.YChild)
    release.EntityData.Children["input"] = types.YChild{"Input", &release.Input}
    release.EntityData.Children["output"] = types.YChild{"Output", &release.Output}
    release.EntityData.Leafs = make(map[string]types.YLeaf)
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["dhcp"] = types.YLeaf{"Dhcp", input.Dhcp}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    reload.EntityData.Children = make(map[string]types.YChild)
    reload.EntityData.Children["input"] = types.YChild{"Input", &reload.Input}
    reload.EntityData.Children["output"] = types.YChild{"Output", &reload.Output}
    reload.EntityData.Leafs = make(map[string]types.YLeaf)
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["force"] = types.YLeaf{"Force", input.Force}
    input.EntityData.Leafs["reason"] = types.YLeaf{"Reason", input.Reason}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    cellular.EntityData.Children = make(map[string]types.YChild)
    cellular.EntityData.Children["input"] = types.YChild{"Input", &cellular.Input}
    cellular.EntityData.Children["output"] = types.YChild{"Output", &cellular.Output}
    cellular.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cellular.EntityData)
}

// Cellular_Input
type Cellular_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    YIfName interface{}

    
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["lte"] = types.YChild{"Lte", &input.Lte}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["_if-name"] = types.YLeaf{"YIfName", input.YIfName}
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

    lte.EntityData.Children = make(map[string]types.YChild)
    lte.EntityData.Children["technology"] = types.YChild{"Technology", &lte.Technology}
    lte.EntityData.Children["profile"] = types.YChild{"Profile", &lte.Profile}
    lte.EntityData.Leafs = make(map[string]types.YLeaf)
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

    technology.EntityData.Children = make(map[string]types.YChild)
    technology.EntityData.Leafs = make(map[string]types.YLeaf)
    technology.EntityData.Leafs["lte"] = types.YLeaf{"Lte", technology.Lte}
    technology.EntityData.Leafs["auto"] = types.YLeaf{"Auto", technology.Auto}
    technology.EntityData.Leafs["umts"] = types.YLeaf{"Umts", technology.Umts}
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

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["delete"] = types.YChild{"Delete", &profile.Delete}
    profile.EntityData.Children["create"] = types.YChild{"Create", &profile.Create}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profile.EntityData)
}

// Cellular_Input_Lte_Profile_Delete
// Delete a Profile
type Cellular_Input_Lte_Profile_Delete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile ID. The type is interface{} with range: 1..16.
    YProfileId interface{}
}

func (delete *Cellular_Input_Lte_Profile_Delete) GetEntityData() *types.CommonEntityData {
    delete.EntityData.YFilter = delete.YFilter
    delete.EntityData.YangName = "delete"
    delete.EntityData.BundleName = "cisco_ios_xe"
    delete.EntityData.ParentYangName = "profile"
    delete.EntityData.SegmentPath = "delete"
    delete.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    delete.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    delete.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    delete.EntityData.Children = make(map[string]types.YChild)
    delete.EntityData.Leafs = make(map[string]types.YLeaf)
    delete.EntityData.Leafs["_profile-id"] = types.YLeaf{"YProfileId", delete.YProfileId}
    return &(delete.EntityData)
}

// Cellular_Input_Lte_Profile_Create
// Create a Profile
type Cellular_Input_Lte_Profile_Create struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile ID. The type is interface{} with range: 1..16.
    YProfileId interface{}

    // Profile name. The type is string.
    YProfName interface{}

    // The type is interface{}.
    None interface{}

    // The type is interface{}.
    Chap interface{}

    // The type is interface{}.
    Pap interface{}

    // The type is interface{}.
    PapChap interface{}

    // Username for authentication. The type is string.
    YUserName interface{}

    // Password for authentication. The type is string.
    YPassword interface{}

    // The type is interface{}.
    Ipv4 interface{}

    // The type is interface{}.
    Ipv6 interface{}

    // The type is interface{}.
    Ipv4V6 interface{}
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

    create.EntityData.Children = make(map[string]types.YChild)
    create.EntityData.Leafs = make(map[string]types.YLeaf)
    create.EntityData.Leafs["_profile-id"] = types.YLeaf{"YProfileId", create.YProfileId}
    create.EntityData.Leafs["_prof_name"] = types.YLeaf{"YProfName", create.YProfName}
    create.EntityData.Leafs["none"] = types.YLeaf{"None", create.None}
    create.EntityData.Leafs["chap"] = types.YLeaf{"Chap", create.Chap}
    create.EntityData.Leafs["pap"] = types.YLeaf{"Pap", create.Pap}
    create.EntityData.Leafs["pap_chap"] = types.YLeaf{"PapChap", create.PapChap}
    create.EntityData.Leafs["_user_name"] = types.YLeaf{"YUserName", create.YUserName}
    create.EntityData.Leafs["_password"] = types.YLeaf{"YPassword", create.YPassword}
    create.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", create.Ipv4}
    create.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", create.Ipv6}
    create.EntityData.Leafs["ipv4v6"] = types.YLeaf{"Ipv4V6", create.Ipv4V6}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    license.EntityData.Children = make(map[string]types.YChild)
    license.EntityData.Children["input"] = types.YChild{"Input", &license.Input}
    license.EntityData.Children["output"] = types.YChild{"Output", &license.Output}
    license.EntityData.Leafs = make(map[string]types.YLeaf)
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["smart"] = types.YChild{"Smart", &input.Smart}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
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

    smart.EntityData.Children = make(map[string]types.YChild)
    smart.EntityData.Children["register"] = types.YChild{"Register", &smart.Register}
    smart.EntityData.Children["renew"] = types.YChild{"Renew", &smart.Renew}
    smart.EntityData.Leafs = make(map[string]types.YLeaf)
    smart.EntityData.Leafs["deregister"] = types.YLeaf{"Deregister", smart.Deregister}
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

    register.EntityData.Children = make(map[string]types.YChild)
    register.EntityData.Leafs = make(map[string]types.YLeaf)
    register.EntityData.Leafs["idtoken"] = types.YLeaf{"Idtoken", register.Idtoken}
    register.EntityData.Leafs["force"] = types.YLeaf{"Force", register.Force}
    return &(register.EntityData)
}

// License_Input_Smart_Renew
type License_Input_Smart_Renew struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Id interface{}

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

    renew.EntityData.Children = make(map[string]types.YChild)
    renew.EntityData.Leafs = make(map[string]types.YLeaf)
    renew.EntityData.Leafs["ID"] = types.YLeaf{"Id", renew.Id}
    renew.EntityData.Leafs["auth"] = types.YLeaf{"Auth", renew.Auth}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    service.EntityData.Children = make(map[string]types.YChild)
    service.EntityData.Children["input"] = types.YChild{"Input", &service.Input}
    service.EntityData.Children["output"] = types.YChild{"Output", &service.Output}
    service.EntityData.Leafs = make(map[string]types.YLeaf)
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["sd-avc"] = types.YChild{"SdAvc", &input.SdAvc}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
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

    sdAvc.EntityData.Children = make(map[string]types.YChild)
    sdAvc.EntityData.Children["configure"] = types.YChild{"Configure", &sdAvc.Configure}
    sdAvc.EntityData.Children["install"] = types.YChild{"Install", &sdAvc.Install}
    sdAvc.EntityData.Children["upgrade"] = types.YChild{"Upgrade", &sdAvc.Upgrade}
    sdAvc.EntityData.Leafs = make(map[string]types.YLeaf)
    sdAvc.EntityData.Leafs["activate"] = types.YLeaf{"Activate", sdAvc.Activate}
    sdAvc.EntityData.Leafs["connect"] = types.YLeaf{"Connect", sdAvc.Connect}
    sdAvc.EntityData.Leafs["help"] = types.YLeaf{"Help", sdAvc.Help}
    sdAvc.EntityData.Leafs["deactivate"] = types.YLeaf{"Deactivate", sdAvc.Deactivate}
    sdAvc.EntityData.Leafs["status"] = types.YLeaf{"Status", sdAvc.Status}
    sdAvc.EntityData.Leafs["unconfigure"] = types.YLeaf{"Unconfigure", sdAvc.Unconfigure}
    sdAvc.EntityData.Leafs["uninstall"] = types.YLeaf{"Uninstall", sdAvc.Uninstall}
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

    configure.EntityData.Children = make(map[string]types.YChild)
    configure.EntityData.Children["gateway"] = types.YChild{"Gateway", &configure.Gateway}
    configure.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configure.EntityData)
}

// Service_Input_SdAvc_Configure_Gateway
type Service_Input_SdAvc_Configure_Gateway struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Interface_ interface{}

    // The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

    gateway.EntityData.Children = make(map[string]types.YChild)
    gateway.EntityData.Leafs = make(map[string]types.YLeaf)
    gateway.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", gateway.Interface_}
    gateway.EntityData.Leafs["service-ip"] = types.YLeaf{"ServiceIp", gateway.ServiceIp}
    gateway.EntityData.Leafs["activate"] = types.YLeaf{"Activate", gateway.Activate}
    return &(gateway.EntityData)
}

// Service_Input_SdAvc_Install
type Service_Input_SdAvc_Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Package_ interface{}
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

    install.EntityData.Children = make(map[string]types.YChild)
    install.EntityData.Leafs = make(map[string]types.YLeaf)
    install.EntityData.Leafs["package"] = types.YLeaf{"Package_", install.Package_}
    return &(install.EntityData)
}

// Service_Input_SdAvc_Upgrade
type Service_Input_SdAvc_Upgrade struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Package_ interface{}
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

    upgrade.EntityData.Children = make(map[string]types.YChild)
    upgrade.EntityData.Leafs = make(map[string]types.YLeaf)
    upgrade.EntityData.Leafs["package"] = types.YLeaf{"Package_", upgrade.Package_}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    virtualService.EntityData.Children = make(map[string]types.YChild)
    virtualService.EntityData.Children["input"] = types.YChild{"Input", &virtualService.Input}
    virtualService.EntityData.Children["output"] = types.YChild{"Output", &virtualService.Output}
    virtualService.EntityData.Leafs = make(map[string]types.YLeaf)
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["install"] = types.YChild{"Install", &input.Install}
    input.EntityData.Children["uninstall"] = types.YChild{"Uninstall", &input.Uninstall}
    input.EntityData.Children["upgrade"] = types.YChild{"Upgrade", &input.Upgrade}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(input.EntityData)
}

// VirtualService_Input_Install
type VirtualService_Input_Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Name interface{}

    // The type is string.
    Package_ interface{}

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

    install.EntityData.Children = make(map[string]types.YChild)
    install.EntityData.Leafs = make(map[string]types.YLeaf)
    install.EntityData.Leafs["name"] = types.YLeaf{"Name", install.Name}
    install.EntityData.Leafs["package"] = types.YLeaf{"Package_", install.Package_}
    install.EntityData.Leafs["media"] = types.YLeaf{"Media", install.Media}
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

    uninstall.EntityData.Children = make(map[string]types.YChild)
    uninstall.EntityData.Leafs = make(map[string]types.YLeaf)
    uninstall.EntityData.Leafs["name"] = types.YLeaf{"Name", uninstall.Name}
    return &(uninstall.EntityData)
}

// VirtualService_Input_Upgrade
type VirtualService_Input_Upgrade struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Name interface{}

    // The type is string.
    Package_ interface{}
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

    upgrade.EntityData.Children = make(map[string]types.YChild)
    upgrade.EntityData.Leafs = make(map[string]types.YLeaf)
    upgrade.EntityData.Leafs["name"] = types.YLeaf{"Name", upgrade.Name}
    upgrade.EntityData.Leafs["package"] = types.YLeaf{"Package_", upgrade.Package_}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

func (copy *Copy) GetEntityData() *types.CommonEntityData {
    copy.EntityData.YFilter = copy.YFilter
    copy.EntityData.YangName = "copy"
    copy.EntityData.BundleName = "cisco_ios_xe"
    copy.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    copy.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:copy"
    copy.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    copy.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    copy.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    copy.EntityData.Children = make(map[string]types.YChild)
    copy.EntityData.Children["input"] = types.YChild{"Input", &copy.Input}
    copy.EntityData.Children["output"] = types.YChild{"Output", &copy.Output}
    copy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(copy.EntityData)
}

// Copy_Input
type Copy_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // b'((((bootflash:)|(harddisk:)|(flash:)|(nvram:)|(ftp:)|(http:)|(https:)|(scp:)|(tftp:)).*)|((running-config)|(startup-config)))'.
    // This attribute is mandatory.
    YSource interface{}

    // The type is string with pattern:
    // b'((((bootflash:)|(harddisk:)|(flash:)|(nvram:)|(ftp:)|(http:)|(https:)|(scp:)|(tftp:)).*)|((running-config)|(startup-config)))'.
    // This attribute is mandatory.
    YDestination interface{}
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["_source"] = types.YLeaf{"YSource", input.YSource}
    input.EntityData.Leafs["_destination"] = types.YLeaf{"YDestination", input.YDestination}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

func (delete *Delete) GetEntityData() *types.CommonEntityData {
    delete.EntityData.YFilter = delete.YFilter
    delete.EntityData.YangName = "delete"
    delete.EntityData.BundleName = "cisco_ios_xe"
    delete.EntityData.ParentYangName = "Cisco-IOS-XE-rpc"
    delete.EntityData.SegmentPath = "Cisco-IOS-XE-rpc:delete"
    delete.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    delete.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    delete.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    delete.EntityData.Children = make(map[string]types.YChild)
    delete.EntityData.Children["input"] = types.YChild{"Input", &delete.Input}
    delete.EntityData.Children["output"] = types.YChild{"Output", &delete.Output}
    delete.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(delete.EntityData)
}

// Delete_Input
type Delete_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string with pattern:
    // b'(((bootflash:)|(harddisk:)|(flash:)|(nvram:)).*)'. This attribute is
    // mandatory.
    YFilename interface{}
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["_filename"] = types.YLeaf{"YFilename", input.YFilename}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    appHosting.EntityData.Children = make(map[string]types.YChild)
    appHosting.EntityData.Children["input"] = types.YChild{"Input", &appHosting.Input}
    appHosting.EntityData.Children["output"] = types.YChild{"Output", &appHosting.Output}
    appHosting.EntityData.Leafs = make(map[string]types.YLeaf)
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["install"] = types.YChild{"Install", &input.Install}
    input.EntityData.Children["uninstall"] = types.YChild{"Uninstall", &input.Uninstall}
    input.EntityData.Children["activate"] = types.YChild{"Activate", &input.Activate}
    input.EntityData.Children["deactivate"] = types.YChild{"Deactivate", &input.Deactivate}
    input.EntityData.Children["start"] = types.YChild{"Start", &input.Start}
    input.EntityData.Children["stop"] = types.YChild{"Stop", &input.Stop}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(input.EntityData)
}

// AppHosting_Input_Install
type AppHosting_Input_Install struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Appid interface{}

    // The type is string.
    Package_ interface{}
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

    install.EntityData.Children = make(map[string]types.YChild)
    install.EntityData.Leafs = make(map[string]types.YLeaf)
    install.EntityData.Leafs["appid"] = types.YLeaf{"Appid", install.Appid}
    install.EntityData.Leafs["package"] = types.YLeaf{"Package_", install.Package_}
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

    uninstall.EntityData.Children = make(map[string]types.YChild)
    uninstall.EntityData.Leafs = make(map[string]types.YLeaf)
    uninstall.EntityData.Leafs["appid"] = types.YLeaf{"Appid", uninstall.Appid}
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

    activate.EntityData.Children = make(map[string]types.YChild)
    activate.EntityData.Leafs = make(map[string]types.YLeaf)
    activate.EntityData.Leafs["appid"] = types.YLeaf{"Appid", activate.Appid}
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

    deactivate.EntityData.Children = make(map[string]types.YChild)
    deactivate.EntityData.Leafs = make(map[string]types.YLeaf)
    deactivate.EntityData.Leafs["appid"] = types.YLeaf{"Appid", deactivate.Appid}
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

    start.EntityData.Children = make(map[string]types.YChild)
    start.EntityData.Leafs = make(map[string]types.YLeaf)
    start.EntityData.Leafs["appid"] = types.YLeaf{"Appid", start.Appid}
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

    stop.EntityData.Children = make(map[string]types.YChild)
    stop.EntityData.Leafs = make(map[string]types.YLeaf)
    stop.EntityData.Leafs["appid"] = types.YLeaf{"Appid", stop.Appid}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
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

    guestshell.EntityData.Children = make(map[string]types.YChild)
    guestshell.EntityData.Children["input"] = types.YChild{"Input", &guestshell.Input}
    guestshell.EntityData.Children["output"] = types.YChild{"Output", &guestshell.Output}
    guestshell.EntityData.Leafs = make(map[string]types.YLeaf)
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

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["destroy"] = types.YLeaf{"Destroy", input.Destroy}
    input.EntityData.Leafs["disable"] = types.YLeaf{"Disable", input.Disable}
    input.EntityData.Leafs["enable"] = types.YLeaf{"Enable", input.Enable}
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

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["result"] = types.YLeaf{"Result", output.Result}
    return &(output.EntityData)
}

