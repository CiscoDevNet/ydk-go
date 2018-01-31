// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-l2tun package configuration.
// 
// This module contains definitions
// for the following management objects:
//   l2tp: L2TPv3 class used for L2VPNs
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tunnel_l2tun_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_l2tun_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-l2tun-cfg l2tp}", reflect.TypeOf(L2Tp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-l2tun-cfg:l2tp", reflect.TypeOf(L2Tp{}))
}

// L2tpDigestHashMethod represents L2tp digest hash method
type L2tpDigestHashMethod string

const (
    // MD5
    L2tpDigestHashMethod_md5 L2tpDigestHashMethod = "md5"

    // SHA1
    L2tpDigestHashMethod_sha1 L2tpDigestHashMethod = "sha1"
)

// L2tpHashMethod represents L2tp hash method
type L2tpHashMethod string

const (
    // MD5
    L2tpHashMethod_md5 L2tpHashMethod = "md5"

    // SHA1
    L2tpHashMethod_sha1 L2tpHashMethod = "sha1"

    // None
    L2tpHashMethod_none L2tpHashMethod = "none"
)

// L2Tp
// L2TPv3 class used for L2VPNs
type L2Tp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of classes.
    Classes L2Tp_Classes
}

func (l2Tp *L2Tp) GetFilter() yfilter.YFilter { return l2Tp.YFilter }

func (l2Tp *L2Tp) SetFilter(yf yfilter.YFilter) { l2Tp.YFilter = yf }

func (l2Tp *L2Tp) GetGoName(yname string) string {
    if yname == "classes" { return "Classes" }
    return ""
}

func (l2Tp *L2Tp) GetSegmentPath() string {
    return "Cisco-IOS-XR-tunnel-l2tun-cfg:l2tp"
}

func (l2Tp *L2Tp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "classes" {
        return &l2Tp.Classes
    }
    return nil
}

func (l2Tp *L2Tp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["classes"] = &l2Tp.Classes
    return children
}

func (l2Tp *L2Tp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (l2Tp *L2Tp) GetBundleName() string { return "cisco_ios_xr" }

func (l2Tp *L2Tp) GetYangName() string { return "l2tp" }

func (l2Tp *L2Tp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Tp *L2Tp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Tp *L2Tp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Tp *L2Tp) SetParent(parent types.Entity) { l2Tp.parent = parent }

func (l2Tp *L2Tp) GetParent() types.Entity { return l2Tp.parent }

func (l2Tp *L2Tp) GetParentYangName() string { return "Cisco-IOS-XR-tunnel-l2tun-cfg" }

// L2Tp_Classes
// List of classes
type L2Tp_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a specific class. The type is slice of
    // L2Tp_Classes_Class.
    Class []L2Tp_Classes_Class
}

func (classes *L2Tp_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *L2Tp_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *L2Tp_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *L2Tp_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *L2Tp_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tp_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *L2Tp_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *L2Tp_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *L2Tp_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *L2Tp_Classes) GetYangName() string { return "classes" }

func (classes *L2Tp_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *L2Tp_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *L2Tp_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *L2Tp_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *L2Tp_Classes) GetParent() types.Entity { return classes.parent }

func (classes *L2Tp_Classes) GetParentYangName() string { return "l2tp" }

// L2Tp_Classes_Class
// Configuration for a specific class
type L2Tp_Classes_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the class name. Regexp:
    // ^[a-z0-9A-Z][-_.a-z0-9A-Z]*$. The type is string with length: 1..31.
    ClassName interface{}

    // Local hostname for control connection authentication. The type is string.
    HostName interface{}

    // Specify to hide AVPs in outgoing control messages. The type is interface{}.
    Hidden interface{}

    // Specify interval (in seconds). The type is interface{} with range: 0..1000.
    // Units are second.
    HelloInterval interface{}

    // Time permitted to set up a control connection. The type is interface{} with
    // range: 60..6000. Units are second.
    TimeoutSetup interface{}

    // Receive window size for the control connection. The type is interface{}
    // with range: 1..16384. Units are byte.
    ReceiveWindow interface{}

    // Congestion control enabled. The type is interface{}.
    CongestionControl interface{}

    // Timeout value for no-user in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TimeoutNoUser interface{}

    // Authenticate the L2TP control connection. The type is interface{} with
    // range: -2147483648..2147483647.
    Authentication interface{}

    // Enable L2TPv3 class used for L2VPNs. The type is interface{}.
    Enable interface{}

    // Specify the password for control channel authentication. The type is string
    // with pattern: (!.+)|([^!].+).
    Password interface{}

    // Security check.
    Security L2Tp_Classes_Class_Security

    // Control message retransmission parameters.
    Retransmit L2Tp_Classes_Class_Retransmit

    // l2TP tunnel.
    Tunnel L2Tp_Classes_Class_Tunnel

    // Message digest authentication for the L2TP control connection.
    Digest L2Tp_Classes_Class_Digest

    // IP TOS value.
    Ip L2Tp_Classes_Class_Ip
}

func (class *L2Tp_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *L2Tp_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *L2Tp_Classes_Class) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "host-name" { return "HostName" }
    if yname == "hidden" { return "Hidden" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "timeout-setup" { return "TimeoutSetup" }
    if yname == "receive-window" { return "ReceiveWindow" }
    if yname == "congestion-control" { return "CongestionControl" }
    if yname == "timeout-no-user" { return "TimeoutNoUser" }
    if yname == "authentication" { return "Authentication" }
    if yname == "enable" { return "Enable" }
    if yname == "password" { return "Password" }
    if yname == "security" { return "Security" }
    if yname == "retransmit" { return "Retransmit" }
    if yname == "tunnel" { return "Tunnel" }
    if yname == "digest" { return "Digest" }
    if yname == "ip" { return "Ip" }
    return ""
}

func (class *L2Tp_Classes_Class) GetSegmentPath() string {
    return "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
}

func (class *L2Tp_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "security" {
        return &class.Security
    }
    if childYangName == "retransmit" {
        return &class.Retransmit
    }
    if childYangName == "tunnel" {
        return &class.Tunnel
    }
    if childYangName == "digest" {
        return &class.Digest
    }
    if childYangName == "ip" {
        return &class.Ip
    }
    return nil
}

func (class *L2Tp_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["security"] = &class.Security
    children["retransmit"] = &class.Retransmit
    children["tunnel"] = &class.Tunnel
    children["digest"] = &class.Digest
    children["ip"] = &class.Ip
    return children
}

func (class *L2Tp_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = class.ClassName
    leafs["host-name"] = class.HostName
    leafs["hidden"] = class.Hidden
    leafs["hello-interval"] = class.HelloInterval
    leafs["timeout-setup"] = class.TimeoutSetup
    leafs["receive-window"] = class.ReceiveWindow
    leafs["congestion-control"] = class.CongestionControl
    leafs["timeout-no-user"] = class.TimeoutNoUser
    leafs["authentication"] = class.Authentication
    leafs["enable"] = class.Enable
    leafs["password"] = class.Password
    return leafs
}

func (class *L2Tp_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *L2Tp_Classes_Class) GetYangName() string { return "class" }

func (class *L2Tp_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *L2Tp_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *L2Tp_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *L2Tp_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *L2Tp_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *L2Tp_Classes_Class) GetParentYangName() string { return "classes" }

// L2Tp_Classes_Class_Security
// Security check
type L2Tp_Classes_Class_Security struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Security check for IP.
    Ip L2Tp_Classes_Class_Security_Ip
}

func (security *L2Tp_Classes_Class_Security) GetFilter() yfilter.YFilter { return security.YFilter }

func (security *L2Tp_Classes_Class_Security) SetFilter(yf yfilter.YFilter) { security.YFilter = yf }

func (security *L2Tp_Classes_Class_Security) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    return ""
}

func (security *L2Tp_Classes_Class_Security) GetSegmentPath() string {
    return "security"
}

func (security *L2Tp_Classes_Class_Security) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ip" {
        return &security.Ip
    }
    return nil
}

func (security *L2Tp_Classes_Class_Security) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ip"] = &security.Ip
    return children
}

func (security *L2Tp_Classes_Class_Security) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (security *L2Tp_Classes_Class_Security) GetBundleName() string { return "cisco_ios_xr" }

func (security *L2Tp_Classes_Class_Security) GetYangName() string { return "security" }

func (security *L2Tp_Classes_Class_Security) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (security *L2Tp_Classes_Class_Security) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (security *L2Tp_Classes_Class_Security) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (security *L2Tp_Classes_Class_Security) SetParent(parent types.Entity) { security.parent = parent }

func (security *L2Tp_Classes_Class_Security) GetParent() types.Entity { return security.parent }

func (security *L2Tp_Classes_Class_Security) GetParentYangName() string { return "class" }

// L2Tp_Classes_Class_Security_Ip
// Security check for IP
type L2Tp_Classes_Class_Security_Ip struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable IP address check for L2TP packets. The type is interface{}.
    AddressCheck interface{}
}

func (ip *L2Tp_Classes_Class_Security_Ip) GetFilter() yfilter.YFilter { return ip.YFilter }

func (ip *L2Tp_Classes_Class_Security_Ip) SetFilter(yf yfilter.YFilter) { ip.YFilter = yf }

func (ip *L2Tp_Classes_Class_Security_Ip) GetGoName(yname string) string {
    if yname == "address-check" { return "AddressCheck" }
    return ""
}

func (ip *L2Tp_Classes_Class_Security_Ip) GetSegmentPath() string {
    return "ip"
}

func (ip *L2Tp_Classes_Class_Security_Ip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ip *L2Tp_Classes_Class_Security_Ip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ip *L2Tp_Classes_Class_Security_Ip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-check"] = ip.AddressCheck
    return leafs
}

func (ip *L2Tp_Classes_Class_Security_Ip) GetBundleName() string { return "cisco_ios_xr" }

func (ip *L2Tp_Classes_Class_Security_Ip) GetYangName() string { return "ip" }

func (ip *L2Tp_Classes_Class_Security_Ip) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ip *L2Tp_Classes_Class_Security_Ip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ip *L2Tp_Classes_Class_Security_Ip) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ip *L2Tp_Classes_Class_Security_Ip) SetParent(parent types.Entity) { ip.parent = parent }

func (ip *L2Tp_Classes_Class_Security_Ip) GetParent() types.Entity { return ip.parent }

func (ip *L2Tp_Classes_Class_Security_Ip) GetParentYangName() string { return "security" }

// L2Tp_Classes_Class_Retransmit
// Control message retransmission parameters
type L2Tp_Classes_Class_Retransmit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify retransmit retry count. The type is interface{} with range:
    // 5..1000.
    Retry interface{}

    // Set retries and timeouts for initial.
    Initial L2Tp_Classes_Class_Retransmit_Initial

    // Set timeout value range.
    Timeout L2Tp_Classes_Class_Retransmit_Timeout
}

func (retransmit *L2Tp_Classes_Class_Retransmit) GetFilter() yfilter.YFilter { return retransmit.YFilter }

func (retransmit *L2Tp_Classes_Class_Retransmit) SetFilter(yf yfilter.YFilter) { retransmit.YFilter = yf }

func (retransmit *L2Tp_Classes_Class_Retransmit) GetGoName(yname string) string {
    if yname == "retry" { return "Retry" }
    if yname == "initial" { return "Initial" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (retransmit *L2Tp_Classes_Class_Retransmit) GetSegmentPath() string {
    return "retransmit"
}

func (retransmit *L2Tp_Classes_Class_Retransmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "initial" {
        return &retransmit.Initial
    }
    if childYangName == "timeout" {
        return &retransmit.Timeout
    }
    return nil
}

func (retransmit *L2Tp_Classes_Class_Retransmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["initial"] = &retransmit.Initial
    children["timeout"] = &retransmit.Timeout
    return children
}

func (retransmit *L2Tp_Classes_Class_Retransmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["retry"] = retransmit.Retry
    return leafs
}

func (retransmit *L2Tp_Classes_Class_Retransmit) GetBundleName() string { return "cisco_ios_xr" }

func (retransmit *L2Tp_Classes_Class_Retransmit) GetYangName() string { return "retransmit" }

func (retransmit *L2Tp_Classes_Class_Retransmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (retransmit *L2Tp_Classes_Class_Retransmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (retransmit *L2Tp_Classes_Class_Retransmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (retransmit *L2Tp_Classes_Class_Retransmit) SetParent(parent types.Entity) { retransmit.parent = parent }

func (retransmit *L2Tp_Classes_Class_Retransmit) GetParent() types.Entity { return retransmit.parent }

func (retransmit *L2Tp_Classes_Class_Retransmit) GetParentYangName() string { return "class" }

// L2Tp_Classes_Class_Retransmit_Initial
// Set retries and timeouts for initial
type L2Tp_Classes_Class_Retransmit_Initial struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the retry number. The type is interface{} with range: 1..1000.
    Retry interface{}

    // Set timeout value range.
    Timeout L2Tp_Classes_Class_Retransmit_Initial_Timeout
}

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetFilter() yfilter.YFilter { return initial.YFilter }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) SetFilter(yf yfilter.YFilter) { initial.YFilter = yf }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetGoName(yname string) string {
    if yname == "retry" { return "Retry" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetSegmentPath() string {
    return "initial"
}

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "timeout" {
        return &initial.Timeout
    }
    return nil
}

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["timeout"] = &initial.Timeout
    return children
}

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["retry"] = initial.Retry
    return leafs
}

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetBundleName() string { return "cisco_ios_xr" }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetYangName() string { return "initial" }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) SetParent(parent types.Entity) { initial.parent = parent }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetParent() types.Entity { return initial.parent }

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetParentYangName() string { return "retransmit" }

// L2Tp_Classes_Class_Retransmit_Initial_Timeout
// Set timeout value range
type L2Tp_Classes_Class_Retransmit_Initial_Timeout struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify minimum timeout. The type is interface{} with range: 1..8.
    Minimum interface{}

    // Specify maximum timeout. The type is interface{} with range: 1..8.
    Maximum interface{}
}

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetFilter() yfilter.YFilter { return timeout.YFilter }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) SetFilter(yf yfilter.YFilter) { timeout.YFilter = yf }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetGoName(yname string) string {
    if yname == "minimum" { return "Minimum" }
    if yname == "maximum" { return "Maximum" }
    return ""
}

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetSegmentPath() string {
    return "timeout"
}

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minimum"] = timeout.Minimum
    leafs["maximum"] = timeout.Maximum
    return leafs
}

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetBundleName() string { return "cisco_ios_xr" }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetYangName() string { return "timeout" }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) SetParent(parent types.Entity) { timeout.parent = parent }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetParent() types.Entity { return timeout.parent }

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetParentYangName() string { return "initial" }

// L2Tp_Classes_Class_Retransmit_Timeout
// Set timeout value range
type L2Tp_Classes_Class_Retransmit_Timeout struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify minimum timeout. The type is interface{} with range: 1..8.
    Minimum interface{}

    // Specify maximum timeout. The type is interface{} with range: 1..8.
    Maximum interface{}
}

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetFilter() yfilter.YFilter { return timeout.YFilter }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) SetFilter(yf yfilter.YFilter) { timeout.YFilter = yf }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetGoName(yname string) string {
    if yname == "minimum" { return "Minimum" }
    if yname == "maximum" { return "Maximum" }
    return ""
}

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetSegmentPath() string {
    return "timeout"
}

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minimum"] = timeout.Minimum
    leafs["maximum"] = timeout.Maximum
    return leafs
}

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetBundleName() string { return "cisco_ios_xr" }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetYangName() string { return "timeout" }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) SetParent(parent types.Entity) { timeout.parent = parent }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetParent() types.Entity { return timeout.parent }

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetParentYangName() string { return "retransmit" }

// L2Tp_Classes_Class_Tunnel
// l2TP tunnel
type L2Tp_Classes_Class_Tunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel accounting. The type is string.
    Accounting interface{}
}

func (tunnel *L2Tp_Classes_Class_Tunnel) GetFilter() yfilter.YFilter { return tunnel.YFilter }

func (tunnel *L2Tp_Classes_Class_Tunnel) SetFilter(yf yfilter.YFilter) { tunnel.YFilter = yf }

func (tunnel *L2Tp_Classes_Class_Tunnel) GetGoName(yname string) string {
    if yname == "accounting" { return "Accounting" }
    return ""
}

func (tunnel *L2Tp_Classes_Class_Tunnel) GetSegmentPath() string {
    return "tunnel"
}

func (tunnel *L2Tp_Classes_Class_Tunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnel *L2Tp_Classes_Class_Tunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnel *L2Tp_Classes_Class_Tunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accounting"] = tunnel.Accounting
    return leafs
}

func (tunnel *L2Tp_Classes_Class_Tunnel) GetBundleName() string { return "cisco_ios_xr" }

func (tunnel *L2Tp_Classes_Class_Tunnel) GetYangName() string { return "tunnel" }

func (tunnel *L2Tp_Classes_Class_Tunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnel *L2Tp_Classes_Class_Tunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnel *L2Tp_Classes_Class_Tunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnel *L2Tp_Classes_Class_Tunnel) SetParent(parent types.Entity) { tunnel.parent = parent }

func (tunnel *L2Tp_Classes_Class_Tunnel) GetParent() types.Entity { return tunnel.parent }

func (tunnel *L2Tp_Classes_Class_Tunnel) GetParentYangName() string { return "class" }

// L2Tp_Classes_Class_Digest
// Message digest authentication for the L2TP
// control connection
type L2Tp_Classes_Class_Digest struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify hash method. The type is L2tpDigestHashMethod.
    Hash interface{}

    // Disable digest checking. The type is interface{}.
    CheckDisable interface{}

    // Set shared secret for message digest.
    Secrets L2Tp_Classes_Class_Digest_Secrets
}

func (digest *L2Tp_Classes_Class_Digest) GetFilter() yfilter.YFilter { return digest.YFilter }

func (digest *L2Tp_Classes_Class_Digest) SetFilter(yf yfilter.YFilter) { digest.YFilter = yf }

func (digest *L2Tp_Classes_Class_Digest) GetGoName(yname string) string {
    if yname == "hash" { return "Hash" }
    if yname == "check-disable" { return "CheckDisable" }
    if yname == "secrets" { return "Secrets" }
    return ""
}

func (digest *L2Tp_Classes_Class_Digest) GetSegmentPath() string {
    return "digest"
}

func (digest *L2Tp_Classes_Class_Digest) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secrets" {
        return &digest.Secrets
    }
    return nil
}

func (digest *L2Tp_Classes_Class_Digest) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["secrets"] = &digest.Secrets
    return children
}

func (digest *L2Tp_Classes_Class_Digest) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hash"] = digest.Hash
    leafs["check-disable"] = digest.CheckDisable
    return leafs
}

func (digest *L2Tp_Classes_Class_Digest) GetBundleName() string { return "cisco_ios_xr" }

func (digest *L2Tp_Classes_Class_Digest) GetYangName() string { return "digest" }

func (digest *L2Tp_Classes_Class_Digest) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (digest *L2Tp_Classes_Class_Digest) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (digest *L2Tp_Classes_Class_Digest) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (digest *L2Tp_Classes_Class_Digest) SetParent(parent types.Entity) { digest.parent = parent }

func (digest *L2Tp_Classes_Class_Digest) GetParent() types.Entity { return digest.parent }

func (digest *L2Tp_Classes_Class_Digest) GetParentYangName() string { return "class" }

// L2Tp_Classes_Class_Digest_Secrets
// Set shared secret for message digest
type L2Tp_Classes_Class_Digest_Secrets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The encrypted user secret and hash method. The type is slice of
    // L2Tp_Classes_Class_Digest_Secrets_Secret.
    Secret []L2Tp_Classes_Class_Digest_Secrets_Secret
}

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetFilter() yfilter.YFilter { return secrets.YFilter }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) SetFilter(yf yfilter.YFilter) { secrets.YFilter = yf }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetGoName(yname string) string {
    if yname == "secret" { return "Secret" }
    return ""
}

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetSegmentPath() string {
    return "secrets"
}

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secret" {
        for _, c := range secrets.Secret {
            if secrets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := L2Tp_Classes_Class_Digest_Secrets_Secret{}
        secrets.Secret = append(secrets.Secret, child)
        return &secrets.Secret[len(secrets.Secret)-1]
    }
    return nil
}

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range secrets.Secret {
        children[secrets.Secret[i].GetSegmentPath()] = &secrets.Secret[i]
    }
    return children
}

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetBundleName() string { return "cisco_ios_xr" }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetYangName() string { return "secrets" }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) SetParent(parent types.Entity) { secrets.parent = parent }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetParent() types.Entity { return secrets.parent }

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetParentYangName() string { return "digest" }

// L2Tp_Classes_Class_Digest_Secrets_Secret
// The encrypted user secret and hash method
type L2Tp_Classes_Class_Digest_Secrets_Secret struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the encrypted user secret. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    SecretName interface{}

    // Specify hash method. The type is L2tpHashMethod. This attribute is
    // mandatory.
    Hash interface{}
}

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetFilter() yfilter.YFilter { return secret.YFilter }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) SetFilter(yf yfilter.YFilter) { secret.YFilter = yf }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetGoName(yname string) string {
    if yname == "secret-name" { return "SecretName" }
    if yname == "hash" { return "Hash" }
    return ""
}

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetSegmentPath() string {
    return "secret" + "[secret-name='" + fmt.Sprintf("%v", secret.SecretName) + "']"
}

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["secret-name"] = secret.SecretName
    leafs["hash"] = secret.Hash
    return leafs
}

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetBundleName() string { return "cisco_ios_xr" }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetYangName() string { return "secret" }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) SetParent(parent types.Entity) { secret.parent = parent }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetParent() types.Entity { return secret.parent }

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetParentYangName() string { return "secrets" }

// L2Tp_Classes_Class_Ip
// IP TOS value
type L2Tp_Classes_Class_Ip struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP TOS value (decimal). The type is interface{} with range: 0..255.
    Tos interface{}
}

func (ip *L2Tp_Classes_Class_Ip) GetFilter() yfilter.YFilter { return ip.YFilter }

func (ip *L2Tp_Classes_Class_Ip) SetFilter(yf yfilter.YFilter) { ip.YFilter = yf }

func (ip *L2Tp_Classes_Class_Ip) GetGoName(yname string) string {
    if yname == "tos" { return "Tos" }
    return ""
}

func (ip *L2Tp_Classes_Class_Ip) GetSegmentPath() string {
    return "ip"
}

func (ip *L2Tp_Classes_Class_Ip) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ip *L2Tp_Classes_Class_Ip) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ip *L2Tp_Classes_Class_Ip) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tos"] = ip.Tos
    return leafs
}

func (ip *L2Tp_Classes_Class_Ip) GetBundleName() string { return "cisco_ios_xr" }

func (ip *L2Tp_Classes_Class_Ip) GetYangName() string { return "ip" }

func (ip *L2Tp_Classes_Class_Ip) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ip *L2Tp_Classes_Class_Ip) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ip *L2Tp_Classes_Class_Ip) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ip *L2Tp_Classes_Class_Ip) SetParent(parent types.Entity) { ip.parent = parent }

func (ip *L2Tp_Classes_Class_Ip) GetParent() types.Entity { return ip.parent }

func (ip *L2Tp_Classes_Class_Ip) GetParentYangName() string { return "class" }

