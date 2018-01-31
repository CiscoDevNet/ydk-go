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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc reload}", reflect.TypeOf(Reload{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:reload", reflect.TypeOf(Reload{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-rpc license}", reflect.TypeOf(License{}))
    ydk.RegisterEntity("Cisco-IOS-XE-rpc:license", reflect.TypeOf(License{}))
}

// Switch
type Switch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Switch_Input

    
    Output Switch_Output
}

func (self *Switch) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Switch) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Switch) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (self *Switch) GetSegmentPath() string {
    return "Cisco-IOS-XE-rpc:switch"
}

func (self *Switch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &self.Input
    }
    if childYangName == "output" {
        return &self.Output
    }
    return nil
}

func (self *Switch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &self.Input
    children["output"] = &self.Output
    return children
}

func (self *Switch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *Switch) GetBundleName() string { return "cisco_ios_xe" }

func (self *Switch) GetYangName() string { return "switch" }

func (self *Switch) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *Switch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *Switch) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *Switch) SetParent(parent types.Entity) { self.parent = parent }

func (self *Switch) GetParent() types.Entity { return self.parent }

func (self *Switch) GetParentYangName() string { return "Cisco-IOS-XE-rpc" }

// Switch_Input
type Switch_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{} with range: 1..9. This attribute is mandatory.
    SwitchNumber interface{}

    // <1-15>  Switch Priority. The type is interface{} with range: 1..15.
    Priority interface{}

    // <1-9>  New number of the Switch. The type is interface{} with range: 1..9.
    Renumber interface{}

    
    Statck Switch_Input_Statck
}

func (input *Switch_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Switch_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Switch_Input) GetGoName(yname string) string {
    if yname == "switch-number" { return "SwitchNumber" }
    if yname == "priority" { return "Priority" }
    if yname == "renumber" { return "Renumber" }
    if yname == "statck" { return "Statck" }
    return ""
}

func (input *Switch_Input) GetSegmentPath() string {
    return "input"
}

func (input *Switch_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statck" {
        return &input.Statck
    }
    return nil
}

func (input *Switch_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statck"] = &input.Statck
    return children
}

func (input *Switch_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["switch-number"] = input.SwitchNumber
    leafs["priority"] = input.Priority
    leafs["renumber"] = input.Renumber
    return leafs
}

func (input *Switch_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *Switch_Input) GetYangName() string { return "input" }

func (input *Switch_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *Switch_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *Switch_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *Switch_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Switch_Input) GetParent() types.Entity { return input.parent }

func (input *Switch_Input) GetParentYangName() string { return "switch" }

// Switch_Input_Statck
type Switch_Input_Statck struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // <1-2>  Stack port number to enable/disable. The type is interface{} with
    // range: 1..2.
    Port interface{}
}

func (statck *Switch_Input_Statck) GetFilter() yfilter.YFilter { return statck.YFilter }

func (statck *Switch_Input_Statck) SetFilter(yf yfilter.YFilter) { statck.YFilter = yf }

func (statck *Switch_Input_Statck) GetGoName(yname string) string {
    if yname == "port" { return "Port" }
    return ""
}

func (statck *Switch_Input_Statck) GetSegmentPath() string {
    return "statck"
}

func (statck *Switch_Input_Statck) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statck *Switch_Input_Statck) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statck *Switch_Input_Statck) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port"] = statck.Port
    return leafs
}

func (statck *Switch_Input_Statck) GetBundleName() string { return "cisco_ios_xe" }

func (statck *Switch_Input_Statck) GetYangName() string { return "statck" }

func (statck *Switch_Input_Statck) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (statck *Switch_Input_Statck) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (statck *Switch_Input_Statck) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (statck *Switch_Input_Statck) SetParent(parent types.Entity) { statck.parent = parent }

func (statck *Switch_Input_Statck) GetParent() types.Entity { return statck.parent }

func (statck *Switch_Input_Statck) GetParentYangName() string { return "input" }

// Switch_Output
type Switch_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Switch_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Switch_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Switch_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *Switch_Output) GetSegmentPath() string {
    return "output"
}

func (output *Switch_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Switch_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Switch_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *Switch_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *Switch_Output) GetYangName() string { return "output" }

func (output *Switch_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *Switch_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *Switch_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *Switch_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Switch_Output) GetParent() types.Entity { return output.parent }

func (output *Switch_Output) GetParentYangName() string { return "switch" }

// Default
// Set a command to its defaults
type Default struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Default_Input

    
    Output Default_Output
}

func (self *Default) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Default) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Default) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (self *Default) GetSegmentPath() string {
    return "Cisco-IOS-XE-rpc:default"
}

func (self *Default) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &self.Input
    }
    if childYangName == "output" {
        return &self.Output
    }
    return nil
}

func (self *Default) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &self.Input
    children["output"] = &self.Output
    return children
}

func (self *Default) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *Default) GetBundleName() string { return "cisco_ios_xe" }

func (self *Default) GetYangName() string { return "default" }

func (self *Default) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *Default) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *Default) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *Default) SetParent(parent types.Entity) { self.parent = parent }

func (self *Default) GetParent() types.Entity { return self.parent }

func (self *Default) GetParentYangName() string { return "Cisco-IOS-XE-rpc" }

// Default_Input
type Default_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Select an interface to configure. The type is string.
    Interface interface{}
}

func (input *Default_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Default_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Default_Input) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (input *Default_Input) GetSegmentPath() string {
    return "input"
}

func (input *Default_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Default_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Default_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = input.Interface
    return leafs
}

func (input *Default_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *Default_Input) GetYangName() string { return "input" }

func (input *Default_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *Default_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *Default_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *Default_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Default_Input) GetParent() types.Entity { return input.parent }

func (input *Default_Input) GetParentYangName() string { return "default" }

// Default_Output
type Default_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Default_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Default_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Default_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *Default_Output) GetSegmentPath() string {
    return "output"
}

func (output *Default_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Default_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Default_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *Default_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *Default_Output) GetYangName() string { return "output" }

func (output *Default_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *Default_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *Default_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *Default_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Default_Output) GetParent() types.Entity { return output.parent }

func (output *Default_Output) GetParentYangName() string { return "default" }

// Reload
// Halt and perform a cold restart
type Reload struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Output Reload_Output
}

func (reload *Reload) GetFilter() yfilter.YFilter { return reload.YFilter }

func (reload *Reload) SetFilter(yf yfilter.YFilter) { reload.YFilter = yf }

func (reload *Reload) GetGoName(yname string) string {
    if yname == "output" { return "Output" }
    return ""
}

func (reload *Reload) GetSegmentPath() string {
    return "Cisco-IOS-XE-rpc:reload"
}

func (reload *Reload) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output" {
        return &reload.Output
    }
    return nil
}

func (reload *Reload) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output"] = &reload.Output
    return children
}

func (reload *Reload) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (reload *Reload) GetBundleName() string { return "cisco_ios_xe" }

func (reload *Reload) GetYangName() string { return "reload" }

func (reload *Reload) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (reload *Reload) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (reload *Reload) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (reload *Reload) SetParent(parent types.Entity) { reload.parent = parent }

func (reload *Reload) GetParent() types.Entity { return reload.parent }

func (reload *Reload) GetParentYangName() string { return "Cisco-IOS-XE-rpc" }

// Reload_Output
type Reload_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *Reload_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Reload_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Reload_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *Reload_Output) GetSegmentPath() string {
    return "output"
}

func (output *Reload_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Reload_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Reload_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *Reload_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *Reload_Output) GetYangName() string { return "output" }

func (output *Reload_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *Reload_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *Reload_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *Reload_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Reload_Output) GetParent() types.Entity { return output.parent }

func (output *Reload_Output) GetParentYangName() string { return "reload" }

// License
type License struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input License_Input

    
    Output License_Output
}

func (license *License) GetFilter() yfilter.YFilter { return license.YFilter }

func (license *License) SetFilter(yf yfilter.YFilter) { license.YFilter = yf }

func (license *License) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (license *License) GetSegmentPath() string {
    return "Cisco-IOS-XE-rpc:license"
}

func (license *License) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &license.Input
    }
    if childYangName == "output" {
        return &license.Output
    }
    return nil
}

func (license *License) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &license.Input
    children["output"] = &license.Output
    return children
}

func (license *License) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (license *License) GetBundleName() string { return "cisco_ios_xe" }

func (license *License) GetYangName() string { return "license" }

func (license *License) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (license *License) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (license *License) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (license *License) SetParent(parent types.Entity) { license.parent = parent }

func (license *License) GetParent() types.Entity { return license.parent }

func (license *License) GetParentYangName() string { return "Cisco-IOS-XE-rpc" }

// License_Input
type License_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Smart License_Input_Smart
}

func (input *License_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *License_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *License_Input) GetGoName(yname string) string {
    if yname == "smart" { return "Smart" }
    return ""
}

func (input *License_Input) GetSegmentPath() string {
    return "input"
}

func (input *License_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "smart" {
        return &input.Smart
    }
    return nil
}

func (input *License_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["smart"] = &input.Smart
    return children
}

func (input *License_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *License_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *License_Input) GetYangName() string { return "input" }

func (input *License_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *License_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *License_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *License_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *License_Input) GetParent() types.Entity { return input.parent }

func (input *License_Input) GetParentYangName() string { return "license" }

// License_Input_Smart
type License_Input_Smart struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{}.
    Deregister interface{}

    
    Register License_Input_Smart_Register

    
    Renew License_Input_Smart_Renew
}

func (smart *License_Input_Smart) GetFilter() yfilter.YFilter { return smart.YFilter }

func (smart *License_Input_Smart) SetFilter(yf yfilter.YFilter) { smart.YFilter = yf }

func (smart *License_Input_Smart) GetGoName(yname string) string {
    if yname == "deregister" { return "Deregister" }
    if yname == "register" { return "Register" }
    if yname == "renew" { return "Renew" }
    return ""
}

func (smart *License_Input_Smart) GetSegmentPath() string {
    return "smart"
}

func (smart *License_Input_Smart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "register" {
        return &smart.Register
    }
    if childYangName == "renew" {
        return &smart.Renew
    }
    return nil
}

func (smart *License_Input_Smart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["register"] = &smart.Register
    children["renew"] = &smart.Renew
    return children
}

func (smart *License_Input_Smart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["deregister"] = smart.Deregister
    return leafs
}

func (smart *License_Input_Smart) GetBundleName() string { return "cisco_ios_xe" }

func (smart *License_Input_Smart) GetYangName() string { return "smart" }

func (smart *License_Input_Smart) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (smart *License_Input_Smart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (smart *License_Input_Smart) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (smart *License_Input_Smart) SetParent(parent types.Entity) { smart.parent = parent }

func (smart *License_Input_Smart) GetParent() types.Entity { return smart.parent }

func (smart *License_Input_Smart) GetParentYangName() string { return "input" }

// License_Input_Smart_Register
type License_Input_Smart_Register struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{}.
    Idtoken interface{}
}

func (register *License_Input_Smart_Register) GetFilter() yfilter.YFilter { return register.YFilter }

func (register *License_Input_Smart_Register) SetFilter(yf yfilter.YFilter) { register.YFilter = yf }

func (register *License_Input_Smart_Register) GetGoName(yname string) string {
    if yname == "idtoken" { return "Idtoken" }
    return ""
}

func (register *License_Input_Smart_Register) GetSegmentPath() string {
    return "register"
}

func (register *License_Input_Smart_Register) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (register *License_Input_Smart_Register) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (register *License_Input_Smart_Register) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["idtoken"] = register.Idtoken
    return leafs
}

func (register *License_Input_Smart_Register) GetBundleName() string { return "cisco_ios_xe" }

func (register *License_Input_Smart_Register) GetYangName() string { return "register" }

func (register *License_Input_Smart_Register) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (register *License_Input_Smart_Register) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (register *License_Input_Smart_Register) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (register *License_Input_Smart_Register) SetParent(parent types.Entity) { register.parent = parent }

func (register *License_Input_Smart_Register) GetParent() types.Entity { return register.parent }

func (register *License_Input_Smart_Register) GetParentYangName() string { return "smart" }

// License_Input_Smart_Renew
type License_Input_Smart_Renew struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{}.
    Id interface{}

    // The type is interface{}.
    Auth interface{}
}

func (renew *License_Input_Smart_Renew) GetFilter() yfilter.YFilter { return renew.YFilter }

func (renew *License_Input_Smart_Renew) SetFilter(yf yfilter.YFilter) { renew.YFilter = yf }

func (renew *License_Input_Smart_Renew) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "auth" { return "Auth" }
    return ""
}

func (renew *License_Input_Smart_Renew) GetSegmentPath() string {
    return "renew"
}

func (renew *License_Input_Smart_Renew) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (renew *License_Input_Smart_Renew) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (renew *License_Input_Smart_Renew) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = renew.Id
    leafs["auth"] = renew.Auth
    return leafs
}

func (renew *License_Input_Smart_Renew) GetBundleName() string { return "cisco_ios_xe" }

func (renew *License_Input_Smart_Renew) GetYangName() string { return "renew" }

func (renew *License_Input_Smart_Renew) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (renew *License_Input_Smart_Renew) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (renew *License_Input_Smart_Renew) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (renew *License_Input_Smart_Renew) SetParent(parent types.Entity) { renew.parent = parent }

func (renew *License_Input_Smart_Renew) GetParent() types.Entity { return renew.parent }

func (renew *License_Input_Smart_Renew) GetParentYangName() string { return "smart" }

// License_Output
type License_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output returned by the network element. The type is string.
    Result interface{}
}

func (output *License_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *License_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *License_Output) GetGoName(yname string) string {
    if yname == "result" { return "Result" }
    return ""
}

func (output *License_Output) GetSegmentPath() string {
    return "output"
}

func (output *License_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *License_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *License_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["result"] = output.Result
    return leafs
}

func (output *License_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *License_Output) GetYangName() string { return "output" }

func (output *License_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *License_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *License_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *License_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *License_Output) GetParent() types.Entity { return output.parent }

func (output *License_Output) GetParentYangName() string { return "license" }

