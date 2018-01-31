// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-objmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   object-group: Object-group configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_objmgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_objmgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-objmgr-cfg object-group}", reflect.TypeOf(ObjectGroup{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-objmgr-cfg:object-group", reflect.TypeOf(ObjectGroup{}))
}

// EndPort represents End port
type EndPort string

const (
    // Echo (7)
    EndPort_echo EndPort = "echo"

    // Discard (9)
    EndPort_discard EndPort = "discard"

    // Daytime (13)
    EndPort_daytime EndPort = "daytime"

    // Character generator (19)
    EndPort_chargen EndPort = "chargen"

    // FTP data connections (used infrequently, 20)
    EndPort_ftp_data EndPort = "ftp-data"

    // File Transfer Protocol (21)
    EndPort_ftp EndPort = "ftp"

    // Secure Shell (22)
    EndPort_ssh EndPort = "ssh"

    // Telnet (23)
    EndPort_telnet EndPort = "telnet"

    // Simple Mail Transport Protocol (25)
    EndPort_smtp EndPort = "smtp"

    // Time (37)
    EndPort_time EndPort = "time"

    // Nicname (43)
    EndPort_nicname EndPort = "nicname"

    // TAC Access Control System (49)
    EndPort_tacacs EndPort = "tacacs"

    // Domain Name Service (53)
    EndPort_domain EndPort = "domain"

    // Gopher (70)
    EndPort_gopher EndPort = "gopher"

    // Finger (79)
    EndPort_finger EndPort = "finger"

    // World Wide Web (HTTP, 80)
    EndPort_www EndPort = "www"

    // NIC hostname server (101)
    EndPort_host_name EndPort = "host-name"

    // Post Office Protocol v2 (109)
    EndPort_pop2 EndPort = "pop2"

    // Post Office Protocol v3 (110)
    EndPort_pop3 EndPort = "pop3"

    // Sun Remote Procedure Call (111)
    EndPort_sun_rpc EndPort = "sun-rpc"

    // Ident Protocol (113)
    EndPort_ident EndPort = "ident"

    // Network News Transport Protocol (119)
    EndPort_nntp EndPort = "nntp"

    // Border Gateway Protocol (179)
    EndPort_bgp EndPort = "bgp"

    // Internet Relay Chat (194)
    EndPort_irc EndPort = "irc"

    // PIM Auto-RP (496)
    EndPort_pim_auto_rp EndPort = "pim-auto-rp"

    // Exec (rsh, 512)
    EndPort_exec EndPort = "exec"

    // Login (rlogin, 513)
    EndPort_login EndPort = "login"

    // Remote commands (rcmd, 514)
    EndPort_cmd EndPort = "cmd"

    // Printer service (515)
    EndPort_lpd EndPort = "lpd"

    // Unix-to-Unix Copy Program (540)
    EndPort_uucp EndPort = "uucp"

    // Kerberos login (543)
    EndPort_klogin EndPort = "klogin"

    // Kerberos shell (544)
    EndPort_kshell EndPort = "kshell"

    // Talk (517)
    EndPort_talk EndPort = "talk"

    // LDP session connection attempts (MPLS, 646)
    EndPort_ldp EndPort = "ldp"
)

// PortOperator represents Port operator
type PortOperator string

const (
    // Match packets on ports equal to entered port
    // number
    PortOperator_equal PortOperator = "equal"

    // Match packets on ports not equal to entered
    // port number
    PortOperator_not_equal PortOperator = "not-equal"

    // Match packets on ports greater than entered
    // port number
    PortOperator_greater_than PortOperator = "greater-than"

    // Match packets on ports less than entered port
    // number
    PortOperator_less_than PortOperator = "less-than"
)

// Port represents Port
type Port string

const (
    // Echo (7)
    Port_echo Port = "echo"

    // Discard (9)
    Port_discard Port = "discard"

    // Daytime (13)
    Port_daytime Port = "daytime"

    // Character generator (19)
    Port_chargen Port = "chargen"

    // FTP data connections (used infrequently, 20)
    Port_ftp_data Port = "ftp-data"

    // File Transfer Protocol (21)
    Port_ftp Port = "ftp"

    // Secure Shell (22)
    Port_ssh Port = "ssh"

    // Telnet (23)
    Port_telnet Port = "telnet"

    // Simple Mail Transport Protocol (25)
    Port_smtp Port = "smtp"

    // Time (37)
    Port_time Port = "time"

    // Nicname (43)
    Port_nicname Port = "nicname"

    // TAC Access Control System (49)
    Port_tacacs Port = "tacacs"

    // Domain Name Service (53)
    Port_domain Port = "domain"

    // Gopher (70)
    Port_gopher Port = "gopher"

    // Finger (79)
    Port_finger Port = "finger"

    // World Wide Web (HTTP, 80)
    Port_www Port = "www"

    // NIC hostname server (101)
    Port_host_name Port = "host-name"

    // Post Office Protocol v2 (109)
    Port_pop2 Port = "pop2"

    // Post Office Protocol v3 (110)
    Port_pop3 Port = "pop3"

    // Sun Remote Procedure Call (111)
    Port_sun_rpc Port = "sun-rpc"

    // Ident Protocol (113)
    Port_ident Port = "ident"

    // Network News Transport Protocol (119)
    Port_nntp Port = "nntp"

    // Border Gateway Protocol (179)
    Port_bgp Port = "bgp"

    // Internet Relay Chat (194)
    Port_irc Port = "irc"

    // PIM Auto-RP (496)
    Port_pim_auto_rp Port = "pim-auto-rp"

    // Exec (rsh, 512)
    Port_exec Port = "exec"

    // Login (rlogin, 513)
    Port_login Port = "login"

    // Remote commands (rcmd, 514)
    Port_cmd Port = "cmd"

    // Printer service (515)
    Port_lpd Port = "lpd"

    // Unix-to-Unix Copy Program (540)
    Port_uucp Port = "uucp"

    // Kerberos login (543)
    Port_klogin Port = "klogin"

    // Kerberos shell (544)
    Port_kshell Port = "kshell"

    // Talk (517)
    Port_talk Port = "talk"

    // LDP session connection attempts (MPLS, 646)
    Port_ldp Port = "ldp"
)

// StartPort represents Start port
type StartPort string

const (
    // Echo (7)
    StartPort_echo StartPort = "echo"

    // Discard (9)
    StartPort_discard StartPort = "discard"

    // Daytime (13)
    StartPort_daytime StartPort = "daytime"

    // Character generator (19)
    StartPort_chargen StartPort = "chargen"

    // FTP data connections (used infrequently, 20)
    StartPort_ftp_data StartPort = "ftp-data"

    // File Transfer Protocol (21)
    StartPort_ftp StartPort = "ftp"

    // Secure Shell (22)
    StartPort_ssh StartPort = "ssh"

    // Telnet (23)
    StartPort_telnet StartPort = "telnet"

    // Simple Mail Transport Protocol (25)
    StartPort_smtp StartPort = "smtp"

    // Time (37)
    StartPort_time StartPort = "time"

    // Nicname (43)
    StartPort_nicname StartPort = "nicname"

    // TAC Access Control System (49)
    StartPort_tacacs StartPort = "tacacs"

    // Domain Name Service (53)
    StartPort_domain StartPort = "domain"

    // Gopher (70)
    StartPort_gopher StartPort = "gopher"

    // Finger (79)
    StartPort_finger StartPort = "finger"

    // World Wide Web (HTTP, 80)
    StartPort_www StartPort = "www"

    // NIC hostname server (101)
    StartPort_host_name StartPort = "host-name"

    // Post Office Protocol v2 (109)
    StartPort_pop2 StartPort = "pop2"

    // Post Office Protocol v3 (110)
    StartPort_pop3 StartPort = "pop3"

    // Sun Remote Procedure Call (111)
    StartPort_sun_rpc StartPort = "sun-rpc"

    // Ident Protocol (113)
    StartPort_ident StartPort = "ident"

    // Network News Transport Protocol (119)
    StartPort_nntp StartPort = "nntp"

    // Border Gateway Protocol (179)
    StartPort_bgp StartPort = "bgp"

    // Internet Relay Chat (194)
    StartPort_irc StartPort = "irc"

    // PIM Auto-RP (496)
    StartPort_pim_auto_rp StartPort = "pim-auto-rp"

    // Exec (rsh, 512)
    StartPort_exec StartPort = "exec"

    // Login (rlogin, 513)
    StartPort_login StartPort = "login"

    // Remote commands (rcmd, 514)
    StartPort_cmd StartPort = "cmd"

    // Printer service (515)
    StartPort_lpd StartPort = "lpd"

    // Unix-to-Unix Copy Program (540)
    StartPort_uucp StartPort = "uucp"

    // Kerberos login (543)
    StartPort_klogin StartPort = "klogin"

    // Kerberos shell (544)
    StartPort_kshell StartPort = "kshell"

    // Talk (517)
    StartPort_talk StartPort = "talk"

    // LDP session connection attempts (MPLS, 646)
    StartPort_ldp StartPort = "ldp"
)

// ObjectGroup
// Object-group configuration
type ObjectGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port object group.
    Port ObjectGroup_Port

    // Network object group.
    Network ObjectGroup_Network
}

func (objectGroup *ObjectGroup) GetFilter() yfilter.YFilter { return objectGroup.YFilter }

func (objectGroup *ObjectGroup) SetFilter(yf yfilter.YFilter) { objectGroup.YFilter = yf }

func (objectGroup *ObjectGroup) GetGoName(yname string) string {
    if yname == "port" { return "Port" }
    if yname == "network" { return "Network" }
    return ""
}

func (objectGroup *ObjectGroup) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-objmgr-cfg:object-group"
}

func (objectGroup *ObjectGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port" {
        return &objectGroup.Port
    }
    if childYangName == "network" {
        return &objectGroup.Network
    }
    return nil
}

func (objectGroup *ObjectGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["port"] = &objectGroup.Port
    children["network"] = &objectGroup.Network
    return children
}

func (objectGroup *ObjectGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (objectGroup *ObjectGroup) GetBundleName() string { return "cisco_ios_xr" }

func (objectGroup *ObjectGroup) GetYangName() string { return "object-group" }

func (objectGroup *ObjectGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (objectGroup *ObjectGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (objectGroup *ObjectGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (objectGroup *ObjectGroup) SetParent(parent types.Entity) { objectGroup.parent = parent }

func (objectGroup *ObjectGroup) GetParent() types.Entity { return objectGroup.parent }

func (objectGroup *ObjectGroup) GetParentYangName() string { return "Cisco-IOS-XR-infra-objmgr-cfg" }

// ObjectGroup_Port
// Port object group
type ObjectGroup_Port struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of port objects groups.
    UdfObjects ObjectGroup_Port_UdfObjects
}

func (port *ObjectGroup_Port) GetFilter() yfilter.YFilter { return port.YFilter }

func (port *ObjectGroup_Port) SetFilter(yf yfilter.YFilter) { port.YFilter = yf }

func (port *ObjectGroup_Port) GetGoName(yname string) string {
    if yname == "udf-objects" { return "UdfObjects" }
    return ""
}

func (port *ObjectGroup_Port) GetSegmentPath() string {
    return "port"
}

func (port *ObjectGroup_Port) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udf-objects" {
        return &port.UdfObjects
    }
    return nil
}

func (port *ObjectGroup_Port) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["udf-objects"] = &port.UdfObjects
    return children
}

func (port *ObjectGroup_Port) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (port *ObjectGroup_Port) GetBundleName() string { return "cisco_ios_xr" }

func (port *ObjectGroup_Port) GetYangName() string { return "port" }

func (port *ObjectGroup_Port) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (port *ObjectGroup_Port) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (port *ObjectGroup_Port) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (port *ObjectGroup_Port) SetParent(parent types.Entity) { port.parent = parent }

func (port *ObjectGroup_Port) GetParent() types.Entity { return port.parent }

func (port *ObjectGroup_Port) GetParentYangName() string { return "object-group" }

// ObjectGroup_Port_UdfObjects
// Table of port objects groups
type ObjectGroup_Port_UdfObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Port object group. The type is slice of
    // ObjectGroup_Port_UdfObjects_UdfObject.
    UdfObject []ObjectGroup_Port_UdfObjects_UdfObject
}

func (udfObjects *ObjectGroup_Port_UdfObjects) GetFilter() yfilter.YFilter { return udfObjects.YFilter }

func (udfObjects *ObjectGroup_Port_UdfObjects) SetFilter(yf yfilter.YFilter) { udfObjects.YFilter = yf }

func (udfObjects *ObjectGroup_Port_UdfObjects) GetGoName(yname string) string {
    if yname == "udf-object" { return "UdfObject" }
    return ""
}

func (udfObjects *ObjectGroup_Port_UdfObjects) GetSegmentPath() string {
    return "udf-objects"
}

func (udfObjects *ObjectGroup_Port_UdfObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udf-object" {
        for _, c := range udfObjects.UdfObject {
            if udfObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Port_UdfObjects_UdfObject{}
        udfObjects.UdfObject = append(udfObjects.UdfObject, child)
        return &udfObjects.UdfObject[len(udfObjects.UdfObject)-1]
    }
    return nil
}

func (udfObjects *ObjectGroup_Port_UdfObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range udfObjects.UdfObject {
        children[udfObjects.UdfObject[i].GetSegmentPath()] = &udfObjects.UdfObject[i]
    }
    return children
}

func (udfObjects *ObjectGroup_Port_UdfObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (udfObjects *ObjectGroup_Port_UdfObjects) GetBundleName() string { return "cisco_ios_xr" }

func (udfObjects *ObjectGroup_Port_UdfObjects) GetYangName() string { return "udf-objects" }

func (udfObjects *ObjectGroup_Port_UdfObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udfObjects *ObjectGroup_Port_UdfObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udfObjects *ObjectGroup_Port_UdfObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udfObjects *ObjectGroup_Port_UdfObjects) SetParent(parent types.Entity) { udfObjects.parent = parent }

func (udfObjects *ObjectGroup_Port_UdfObjects) GetParent() types.Entity { return udfObjects.parent }

func (udfObjects *ObjectGroup_Port_UdfObjects) GetParentYangName() string { return "port" }

// ObjectGroup_Port_UdfObjects_UdfObject
// Port object group
type ObjectGroup_Port_UdfObjects_UdfObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port object group name - maximum 64 characters.
    // The type is string with length: 1..64.
    ObjectName interface{}

    // Up to 100 characters describing this object. The type is string with
    // length: 1..100.
    Description interface{}

    // Table of port operators.
    Operators ObjectGroup_Port_UdfObjects_UdfObject_Operators

    // Table of nested port object groups.
    NestedGroups ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups

    // Table of port range addresses.
    PortRanges ObjectGroup_Port_UdfObjects_UdfObject_PortRanges
}

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetFilter() yfilter.YFilter { return udfObject.YFilter }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) SetFilter(yf yfilter.YFilter) { udfObject.YFilter = yf }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "description" { return "Description" }
    if yname == "operators" { return "Operators" }
    if yname == "nested-groups" { return "NestedGroups" }
    if yname == "port-ranges" { return "PortRanges" }
    return ""
}

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetSegmentPath() string {
    return "udf-object" + "[object-name='" + fmt.Sprintf("%v", udfObject.ObjectName) + "']"
}

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operators" {
        return &udfObject.Operators
    }
    if childYangName == "nested-groups" {
        return &udfObject.NestedGroups
    }
    if childYangName == "port-ranges" {
        return &udfObject.PortRanges
    }
    return nil
}

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["operators"] = &udfObject.Operators
    children["nested-groups"] = &udfObject.NestedGroups
    children["port-ranges"] = &udfObject.PortRanges
    return children
}

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = udfObject.ObjectName
    leafs["description"] = udfObject.Description
    return leafs
}

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetBundleName() string { return "cisco_ios_xr" }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetYangName() string { return "udf-object" }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) SetParent(parent types.Entity) { udfObject.parent = parent }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetParent() types.Entity { return udfObject.parent }

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetParentYangName() string { return "udf-objects" }

// ObjectGroup_Port_UdfObjects_UdfObject_Operators
// Table of port operators
type ObjectGroup_Port_UdfObjects_UdfObject_Operators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // op class. The type is slice of
    // ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator.
    Operator []ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator
}

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetFilter() yfilter.YFilter { return operators.YFilter }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) SetFilter(yf yfilter.YFilter) { operators.YFilter = yf }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetGoName(yname string) string {
    if yname == "operator" { return "Operator" }
    return ""
}

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetSegmentPath() string {
    return "operators"
}

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "operator" {
        for _, c := range operators.Operator {
            if operators.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator{}
        operators.Operator = append(operators.Operator, child)
        return &operators.Operator[len(operators.Operator)-1]
    }
    return nil
}

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range operators.Operator {
        children[operators.Operator[i].GetSegmentPath()] = &operators.Operator[i]
    }
    return children
}

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetBundleName() string { return "cisco_ios_xr" }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetYangName() string { return "operators" }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) SetParent(parent types.Entity) { operators.parent = parent }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetParent() types.Entity { return operators.parent }

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator
// op class
type ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. operation for ports. The type is PortOperator.
    OperatorType interface{}

    // This attribute is a key. Port number. The type is one of the following
    // types: enumeration Port, or int with range: 0..65535.
    Port interface{}
}

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetFilter() yfilter.YFilter { return operator.YFilter }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) SetFilter(yf yfilter.YFilter) { operator.YFilter = yf }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetGoName(yname string) string {
    if yname == "operator-type" { return "OperatorType" }
    if yname == "port" { return "Port" }
    return ""
}

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetSegmentPath() string {
    return "operator" + "[operator-type='" + fmt.Sprintf("%v", operator.OperatorType) + "']" + "[port='" + fmt.Sprintf("%v", operator.Port) + "']"
}

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["operator-type"] = operator.OperatorType
    leafs["port"] = operator.Port
    return leafs
}

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetBundleName() string { return "cisco_ios_xr" }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetYangName() string { return "operator" }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) SetParent(parent types.Entity) { operator.parent = parent }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetParent() types.Entity { return operator.parent }

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetParentYangName() string { return "operators" }

// ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups
// Table of nested port object groups
type ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // nested object group. The type is slice of
    // ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup.
    NestedGroup []ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetFilter() yfilter.YFilter { return nestedGroups.YFilter }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) SetFilter(yf yfilter.YFilter) { nestedGroups.YFilter = yf }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetGoName(yname string) string {
    if yname == "nested-group" { return "NestedGroup" }
    return ""
}

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetSegmentPath() string {
    return "nested-groups"
}

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nested-group" {
        for _, c := range nestedGroups.NestedGroup {
            if nestedGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup{}
        nestedGroups.NestedGroup = append(nestedGroups.NestedGroup, child)
        return &nestedGroups.NestedGroup[len(nestedGroups.NestedGroup)-1]
    }
    return nil
}

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nestedGroups.NestedGroup {
        children[nestedGroups.NestedGroup[i].GetSegmentPath()] = &nestedGroups.NestedGroup[i]
    }
    return children
}

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetBundleName() string { return "cisco_ios_xr" }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetYangName() string { return "nested-groups" }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) SetParent(parent types.Entity) { nestedGroups.parent = parent }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetParent() types.Entity { return nestedGroups.parent }

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup
// nested object group
type ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of a nested object group. The type is string
    // with length: 1..64.
    NestedGroupName interface{}
}

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetFilter() yfilter.YFilter { return nestedGroup.YFilter }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) SetFilter(yf yfilter.YFilter) { nestedGroup.YFilter = yf }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetGoName(yname string) string {
    if yname == "nested-group-name" { return "NestedGroupName" }
    return ""
}

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetSegmentPath() string {
    return "nested-group" + "[nested-group-name='" + fmt.Sprintf("%v", nestedGroup.NestedGroupName) + "']"
}

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nested-group-name"] = nestedGroup.NestedGroupName
    return leafs
}

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetBundleName() string { return "cisco_ios_xr" }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetYangName() string { return "nested-group" }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) SetParent(parent types.Entity) { nestedGroup.parent = parent }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetParent() types.Entity { return nestedGroup.parent }

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetParentYangName() string { return "nested-groups" }

// ObjectGroup_Port_UdfObjects_UdfObject_PortRanges
// Table of port range addresses
type ObjectGroup_Port_UdfObjects_UdfObject_PortRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match only packets on a given port range. The type is slice of
    // ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange.
    PortRange []ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange
}

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetFilter() yfilter.YFilter { return portRanges.YFilter }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) SetFilter(yf yfilter.YFilter) { portRanges.YFilter = yf }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetGoName(yname string) string {
    if yname == "port-range" { return "PortRange" }
    return ""
}

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetSegmentPath() string {
    return "port-ranges"
}

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "port-range" {
        for _, c := range portRanges.PortRange {
            if portRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange{}
        portRanges.PortRange = append(portRanges.PortRange, child)
        return &portRanges.PortRange[len(portRanges.PortRange)-1]
    }
    return nil
}

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range portRanges.PortRange {
        children[portRanges.PortRange[i].GetSegmentPath()] = &portRanges.PortRange[i]
    }
    return children
}

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetBundleName() string { return "cisco_ios_xr" }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetYangName() string { return "port-ranges" }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) SetParent(parent types.Entity) { portRanges.parent = parent }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetParent() types.Entity { return portRanges.parent }

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange
// Match only packets on a given port range
type ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port number. The type is one of the following
    // types: enumeration StartPort, or int with range: 0..65535.
    StartPort interface{}

    // This attribute is a key. Port number. The type is one of the following
    // types: enumeration EndPort, or int with range: 0..65535.
    EndPort interface{}
}

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetFilter() yfilter.YFilter { return portRange.YFilter }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) SetFilter(yf yfilter.YFilter) { portRange.YFilter = yf }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetGoName(yname string) string {
    if yname == "start-port" { return "StartPort" }
    if yname == "end-port" { return "EndPort" }
    return ""
}

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetSegmentPath() string {
    return "port-range" + "[start-port='" + fmt.Sprintf("%v", portRange.StartPort) + "']" + "[end-port='" + fmt.Sprintf("%v", portRange.EndPort) + "']"
}

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-port"] = portRange.StartPort
    leafs["end-port"] = portRange.EndPort
    return leafs
}

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetBundleName() string { return "cisco_ios_xr" }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetYangName() string { return "port-range" }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) SetParent(parent types.Entity) { portRange.parent = parent }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetParent() types.Entity { return portRange.parent }

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetParentYangName() string { return "port-ranges" }

// ObjectGroup_Network
// Network object group
type ObjectGroup_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 object group.
    Ipv6 ObjectGroup_Network_Ipv6

    // IPv4 object group.
    Ipv4 ObjectGroup_Network_Ipv4
}

func (network *ObjectGroup_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *ObjectGroup_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *ObjectGroup_Network) GetGoName(yname string) string {
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (network *ObjectGroup_Network) GetSegmentPath() string {
    return "network"
}

func (network *ObjectGroup_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &network.Ipv6
    }
    if childYangName == "ipv4" {
        return &network.Ipv4
    }
    return nil
}

func (network *ObjectGroup_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &network.Ipv6
    children["ipv4"] = &network.Ipv4
    return children
}

func (network *ObjectGroup_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (network *ObjectGroup_Network) GetBundleName() string { return "cisco_ios_xr" }

func (network *ObjectGroup_Network) GetYangName() string { return "network" }

func (network *ObjectGroup_Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (network *ObjectGroup_Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (network *ObjectGroup_Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (network *ObjectGroup_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *ObjectGroup_Network) GetParent() types.Entity { return network.parent }

func (network *ObjectGroup_Network) GetParentYangName() string { return "object-group" }

// ObjectGroup_Network_Ipv6
// IPv6 object group
type ObjectGroup_Network_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of ipv6 object groups.
    UdfObjects ObjectGroup_Network_Ipv6_UdfObjects
}

func (ipv6 *ObjectGroup_Network_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *ObjectGroup_Network_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *ObjectGroup_Network_Ipv6) GetGoName(yname string) string {
    if yname == "udf-objects" { return "UdfObjects" }
    return ""
}

func (ipv6 *ObjectGroup_Network_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *ObjectGroup_Network_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udf-objects" {
        return &ipv6.UdfObjects
    }
    return nil
}

func (ipv6 *ObjectGroup_Network_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["udf-objects"] = &ipv6.UdfObjects
    return children
}

func (ipv6 *ObjectGroup_Network_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *ObjectGroup_Network_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *ObjectGroup_Network_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *ObjectGroup_Network_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *ObjectGroup_Network_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *ObjectGroup_Network_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *ObjectGroup_Network_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *ObjectGroup_Network_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *ObjectGroup_Network_Ipv6) GetParentYangName() string { return "network" }

// ObjectGroup_Network_Ipv6_UdfObjects
// Table of ipv6 object groups
type ObjectGroup_Network_Ipv6_UdfObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 object group. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject.
    UdfObject []ObjectGroup_Network_Ipv6_UdfObjects_UdfObject
}

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetFilter() yfilter.YFilter { return udfObjects.YFilter }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) SetFilter(yf yfilter.YFilter) { udfObjects.YFilter = yf }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetGoName(yname string) string {
    if yname == "udf-object" { return "UdfObject" }
    return ""
}

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetSegmentPath() string {
    return "udf-objects"
}

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udf-object" {
        for _, c := range udfObjects.UdfObject {
            if udfObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv6_UdfObjects_UdfObject{}
        udfObjects.UdfObject = append(udfObjects.UdfObject, child)
        return &udfObjects.UdfObject[len(udfObjects.UdfObject)-1]
    }
    return nil
}

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range udfObjects.UdfObject {
        children[udfObjects.UdfObject[i].GetSegmentPath()] = &udfObjects.UdfObject[i]
    }
    return children
}

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetBundleName() string { return "cisco_ios_xr" }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetYangName() string { return "udf-objects" }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) SetParent(parent types.Entity) { udfObjects.parent = parent }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetParent() types.Entity { return udfObjects.parent }

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetParentYangName() string { return "ipv6" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject
// IPv6 object group
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 object group name - maximum 64 characters.
    // The type is string with length: 1..64.
    ObjectName interface{}

    // Up to 100 characters describing this object. The type is string with
    // length: 1..100.
    Description interface{}

    // Table of nested ipv6 object groups.
    NestedGroups ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups

    // Table of ipv6 address ranges.
    AddressRanges ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges

    // Table of ipv6 addresses.
    Addresses ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses

    // Table of ipv6 host addresses.
    Hosts ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts
}

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetFilter() yfilter.YFilter { return udfObject.YFilter }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) SetFilter(yf yfilter.YFilter) { udfObject.YFilter = yf }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "description" { return "Description" }
    if yname == "nested-groups" { return "NestedGroups" }
    if yname == "address-ranges" { return "AddressRanges" }
    if yname == "addresses" { return "Addresses" }
    if yname == "hosts" { return "Hosts" }
    return ""
}

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetSegmentPath() string {
    return "udf-object" + "[object-name='" + fmt.Sprintf("%v", udfObject.ObjectName) + "']"
}

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nested-groups" {
        return &udfObject.NestedGroups
    }
    if childYangName == "address-ranges" {
        return &udfObject.AddressRanges
    }
    if childYangName == "addresses" {
        return &udfObject.Addresses
    }
    if childYangName == "hosts" {
        return &udfObject.Hosts
    }
    return nil
}

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nested-groups"] = &udfObject.NestedGroups
    children["address-ranges"] = &udfObject.AddressRanges
    children["addresses"] = &udfObject.Addresses
    children["hosts"] = &udfObject.Hosts
    return children
}

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = udfObject.ObjectName
    leafs["description"] = udfObject.Description
    return leafs
}

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetBundleName() string { return "cisco_ios_xr" }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetYangName() string { return "udf-object" }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) SetParent(parent types.Entity) { udfObject.parent = parent }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetParent() types.Entity { return udfObject.parent }

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetParentYangName() string { return "udf-objects" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups
// Table of nested ipv6 object groups
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // nested object group. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup.
    NestedGroup []ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetFilter() yfilter.YFilter { return nestedGroups.YFilter }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) SetFilter(yf yfilter.YFilter) { nestedGroups.YFilter = yf }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetGoName(yname string) string {
    if yname == "nested-group" { return "NestedGroup" }
    return ""
}

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetSegmentPath() string {
    return "nested-groups"
}

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nested-group" {
        for _, c := range nestedGroups.NestedGroup {
            if nestedGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup{}
        nestedGroups.NestedGroup = append(nestedGroups.NestedGroup, child)
        return &nestedGroups.NestedGroup[len(nestedGroups.NestedGroup)-1]
    }
    return nil
}

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nestedGroups.NestedGroup {
        children[nestedGroups.NestedGroup[i].GetSegmentPath()] = &nestedGroups.NestedGroup[i]
    }
    return children
}

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetBundleName() string { return "cisco_ios_xr" }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetYangName() string { return "nested-groups" }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) SetParent(parent types.Entity) { nestedGroups.parent = parent }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetParent() types.Entity { return nestedGroups.parent }

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup
// nested object group
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Enter the name of a nested object group. The type
    // is string with length: 1..64.
    NestedGroupName interface{}
}

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetFilter() yfilter.YFilter { return nestedGroup.YFilter }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) SetFilter(yf yfilter.YFilter) { nestedGroup.YFilter = yf }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetGoName(yname string) string {
    if yname == "nested-group-name" { return "NestedGroupName" }
    return ""
}

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetSegmentPath() string {
    return "nested-group" + "[nested-group-name='" + fmt.Sprintf("%v", nestedGroup.NestedGroupName) + "']"
}

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nested-group-name"] = nestedGroup.NestedGroupName
    return leafs
}

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetBundleName() string { return "cisco_ios_xr" }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetYangName() string { return "nested-group" }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) SetParent(parent types.Entity) { nestedGroup.parent = parent }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetParent() types.Entity { return nestedGroup.parent }

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetParentYangName() string { return "nested-groups" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges
// Table of ipv6 address ranges
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Range of host addresses. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange.
    AddressRange []ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange
}

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetFilter() yfilter.YFilter { return addressRanges.YFilter }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) SetFilter(yf yfilter.YFilter) { addressRanges.YFilter = yf }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetGoName(yname string) string {
    if yname == "address-range" { return "AddressRange" }
    return ""
}

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetSegmentPath() string {
    return "address-ranges"
}

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-range" {
        for _, c := range addressRanges.AddressRange {
            if addressRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange{}
        addressRanges.AddressRange = append(addressRanges.AddressRange, child)
        return &addressRanges.AddressRange[len(addressRanges.AddressRange)-1]
    }
    return nil
}

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addressRanges.AddressRange {
        children[addressRanges.AddressRange[i].GetSegmentPath()] = &addressRanges.AddressRange[i]
    }
    return children
}

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetBundleName() string { return "cisco_ios_xr" }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetYangName() string { return "address-ranges" }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) SetParent(parent types.Entity) { addressRanges.parent = parent }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetParent() types.Entity { return addressRanges.parent }

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange
// Range of host addresses
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // This attribute is a key. IPv6 address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    EndAddress interface{}
}

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetFilter() yfilter.YFilter { return addressRange.YFilter }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) SetFilter(yf yfilter.YFilter) { addressRange.YFilter = yf }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "end-address" { return "EndAddress" }
    return ""
}

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetSegmentPath() string {
    return "address-range" + "[start-address='" + fmt.Sprintf("%v", addressRange.StartAddress) + "']" + "[end-address='" + fmt.Sprintf("%v", addressRange.EndAddress) + "']"
}

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = addressRange.StartAddress
    leafs["end-address"] = addressRange.EndAddress
    return leafs
}

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetBundleName() string { return "cisco_ios_xr" }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetYangName() string { return "address-range" }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) SetParent(parent types.Entity) { addressRange.parent = parent }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetParent() types.Entity { return addressRange.parent }

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetParentYangName() string { return "address-ranges" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses
// Table of ipv6 addresses
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address.
    Address []ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address
}

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range addresses.Address {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address{}
        addresses.Address = append(addresses.Address, child)
        return &addresses.Address[len(addresses.Address)-1]
    }
    return nil
}

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.Address {
        children[addresses.Address[i].GetSegmentPath()] = &addresses.Address[i]
    }
    return children
}

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetYangName() string { return "addresses" }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address
// IPv6 address
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 prefix x:x::x/y. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: 0..128.
    PrefixLength interface{}
}

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetSegmentPath() string {
    return "address" + "[prefix='" + fmt.Sprintf("%v", address.Prefix) + "']" + "[prefix-length='" + fmt.Sprintf("%v", address.PrefixLength) + "']"
}

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = address.Prefix
    leafs["prefix-length"] = address.PrefixLength
    return leafs
}

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetYangName() string { return "address" }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetParent() types.Entity { return address.parent }

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetParentYangName() string { return "addresses" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts
// Table of ipv6 host addresses
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single host address. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host.
    Host []ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host
}

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetYangName() string { return "hosts" }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host
// A single host address
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. host ipv6 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HostAddress interface{}
}

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetGoName(yname string) string {
    if yname == "host-address" { return "HostAddress" }
    return ""
}

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetSegmentPath() string {
    return "host" + "[host-address='" + fmt.Sprintf("%v", host.HostAddress) + "']"
}

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-address"] = host.HostAddress
    return leafs
}

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetYangName() string { return "host" }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetParentYangName() string { return "hosts" }

// ObjectGroup_Network_Ipv4
// IPv4 object group
type ObjectGroup_Network_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of ipv4 object groups.
    UdfObjects ObjectGroup_Network_Ipv4_UdfObjects
}

func (ipv4 *ObjectGroup_Network_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *ObjectGroup_Network_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *ObjectGroup_Network_Ipv4) GetGoName(yname string) string {
    if yname == "udf-objects" { return "UdfObjects" }
    return ""
}

func (ipv4 *ObjectGroup_Network_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *ObjectGroup_Network_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udf-objects" {
        return &ipv4.UdfObjects
    }
    return nil
}

func (ipv4 *ObjectGroup_Network_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["udf-objects"] = &ipv4.UdfObjects
    return children
}

func (ipv4 *ObjectGroup_Network_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *ObjectGroup_Network_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *ObjectGroup_Network_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *ObjectGroup_Network_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *ObjectGroup_Network_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *ObjectGroup_Network_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *ObjectGroup_Network_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *ObjectGroup_Network_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *ObjectGroup_Network_Ipv4) GetParentYangName() string { return "network" }

// ObjectGroup_Network_Ipv4_UdfObjects
// Table of ipv4 object groups
type ObjectGroup_Network_Ipv4_UdfObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 object group. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject.
    UdfObject []ObjectGroup_Network_Ipv4_UdfObjects_UdfObject
}

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetFilter() yfilter.YFilter { return udfObjects.YFilter }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) SetFilter(yf yfilter.YFilter) { udfObjects.YFilter = yf }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetGoName(yname string) string {
    if yname == "udf-object" { return "UdfObject" }
    return ""
}

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetSegmentPath() string {
    return "udf-objects"
}

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "udf-object" {
        for _, c := range udfObjects.UdfObject {
            if udfObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv4_UdfObjects_UdfObject{}
        udfObjects.UdfObject = append(udfObjects.UdfObject, child)
        return &udfObjects.UdfObject[len(udfObjects.UdfObject)-1]
    }
    return nil
}

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range udfObjects.UdfObject {
        children[udfObjects.UdfObject[i].GetSegmentPath()] = &udfObjects.UdfObject[i]
    }
    return children
}

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetBundleName() string { return "cisco_ios_xr" }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetYangName() string { return "udf-objects" }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) SetParent(parent types.Entity) { udfObjects.parent = parent }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetParent() types.Entity { return udfObjects.parent }

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetParentYangName() string { return "ipv4" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject
// IPv4 object group
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 object group name - maximum 64 characters.
    // The type is string with length: 1..64.
    ObjectName interface{}

    // Up to 100 characters describing this object. The type is string with
    // length: 1..100.
    Description interface{}

    // Table of nested ipv4 object groups.
    NestedGroups ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups

    // Table of ipv4 host address ranges.
    AddressRanges ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges

    // Table of addresses.
    Addresses ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses

    // Table of host addresses.
    Hosts ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts
}

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetFilter() yfilter.YFilter { return udfObject.YFilter }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) SetFilter(yf yfilter.YFilter) { udfObject.YFilter = yf }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "description" { return "Description" }
    if yname == "nested-groups" { return "NestedGroups" }
    if yname == "address-ranges" { return "AddressRanges" }
    if yname == "addresses" { return "Addresses" }
    if yname == "hosts" { return "Hosts" }
    return ""
}

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetSegmentPath() string {
    return "udf-object" + "[object-name='" + fmt.Sprintf("%v", udfObject.ObjectName) + "']"
}

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nested-groups" {
        return &udfObject.NestedGroups
    }
    if childYangName == "address-ranges" {
        return &udfObject.AddressRanges
    }
    if childYangName == "addresses" {
        return &udfObject.Addresses
    }
    if childYangName == "hosts" {
        return &udfObject.Hosts
    }
    return nil
}

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nested-groups"] = &udfObject.NestedGroups
    children["address-ranges"] = &udfObject.AddressRanges
    children["addresses"] = &udfObject.Addresses
    children["hosts"] = &udfObject.Hosts
    return children
}

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = udfObject.ObjectName
    leafs["description"] = udfObject.Description
    return leafs
}

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetBundleName() string { return "cisco_ios_xr" }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetYangName() string { return "udf-object" }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) SetParent(parent types.Entity) { udfObject.parent = parent }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetParent() types.Entity { return udfObject.parent }

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetParentYangName() string { return "udf-objects" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups
// Table of nested ipv4 object groups
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Nested object group. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup.
    NestedGroup []ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetFilter() yfilter.YFilter { return nestedGroups.YFilter }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) SetFilter(yf yfilter.YFilter) { nestedGroups.YFilter = yf }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetGoName(yname string) string {
    if yname == "nested-group" { return "NestedGroup" }
    return ""
}

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetSegmentPath() string {
    return "nested-groups"
}

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nested-group" {
        for _, c := range nestedGroups.NestedGroup {
            if nestedGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup{}
        nestedGroups.NestedGroup = append(nestedGroups.NestedGroup, child)
        return &nestedGroups.NestedGroup[len(nestedGroups.NestedGroup)-1]
    }
    return nil
}

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nestedGroups.NestedGroup {
        children[nestedGroups.NestedGroup[i].GetSegmentPath()] = &nestedGroups.NestedGroup[i]
    }
    return children
}

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetBundleName() string { return "cisco_ios_xr" }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetYangName() string { return "nested-groups" }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) SetParent(parent types.Entity) { nestedGroups.parent = parent }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetParent() types.Entity { return nestedGroups.parent }

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup
// Nested object group
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Nested object group. The type is string with
    // length: 1..64.
    NestedGroupName interface{}
}

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetFilter() yfilter.YFilter { return nestedGroup.YFilter }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) SetFilter(yf yfilter.YFilter) { nestedGroup.YFilter = yf }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetGoName(yname string) string {
    if yname == "nested-group-name" { return "NestedGroupName" }
    return ""
}

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetSegmentPath() string {
    return "nested-group" + "[nested-group-name='" + fmt.Sprintf("%v", nestedGroup.NestedGroupName) + "']"
}

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nested-group-name"] = nestedGroup.NestedGroupName
    return leafs
}

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetBundleName() string { return "cisco_ios_xr" }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetYangName() string { return "nested-group" }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) SetParent(parent types.Entity) { nestedGroup.parent = parent }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetParent() types.Entity { return nestedGroup.parent }

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetParentYangName() string { return "nested-groups" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges
// Table of ipv4 host address ranges
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Range of host addresses. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange.
    AddressRange []ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange
}

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetFilter() yfilter.YFilter { return addressRanges.YFilter }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) SetFilter(yf yfilter.YFilter) { addressRanges.YFilter = yf }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetGoName(yname string) string {
    if yname == "address-range" { return "AddressRange" }
    return ""
}

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetSegmentPath() string {
    return "address-ranges"
}

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-range" {
        for _, c := range addressRanges.AddressRange {
            if addressRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange{}
        addressRanges.AddressRange = append(addressRanges.AddressRange, child)
        return &addressRanges.AddressRange[len(addressRanges.AddressRange)-1]
    }
    return nil
}

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addressRanges.AddressRange {
        children[addressRanges.AddressRange[i].GetSegmentPath()] = &addressRanges.AddressRange[i]
    }
    return children
}

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetBundleName() string { return "cisco_ios_xr" }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetYangName() string { return "address-ranges" }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) SetParent(parent types.Entity) { addressRanges.parent = parent }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetParent() types.Entity { return addressRanges.parent }

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange
// Range of host addresses
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // This attribute is a key. IPv4 address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    EndAddress interface{}
}

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetFilter() yfilter.YFilter { return addressRange.YFilter }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) SetFilter(yf yfilter.YFilter) { addressRange.YFilter = yf }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "end-address" { return "EndAddress" }
    return ""
}

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetSegmentPath() string {
    return "address-range" + "[start-address='" + fmt.Sprintf("%v", addressRange.StartAddress) + "']" + "[end-address='" + fmt.Sprintf("%v", addressRange.EndAddress) + "']"
}

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = addressRange.StartAddress
    leafs["end-address"] = addressRange.EndAddress
    return leafs
}

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetBundleName() string { return "cisco_ios_xr" }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetYangName() string { return "address-range" }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) SetParent(parent types.Entity) { addressRange.parent = parent }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetParent() types.Entity { return addressRange.parent }

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetParentYangName() string { return "address-ranges" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses
// Table of addresses
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address.
    Address []ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address
}

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        for _, c := range addresses.Address {
            if addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address{}
        addresses.Address = append(addresses.Address, child)
        return &addresses.Address[len(addresses.Address)-1]
    }
    return nil
}

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range addresses.Address {
        children[addresses.Address[i].GetSegmentPath()] = &addresses.Address[i]
    }
    return children
}

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetYangName() string { return "addresses" }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address
// IPv4 address
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address/prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // This attribute is a key. Prefix of the IP Address. The type is interface{}
    // with range: 0..32.
    PrefixLength interface{}
}

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetSegmentPath() string {
    return "address" + "[prefix='" + fmt.Sprintf("%v", address.Prefix) + "']" + "[prefix-length='" + fmt.Sprintf("%v", address.PrefixLength) + "']"
}

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = address.Prefix
    leafs["prefix-length"] = address.PrefixLength
    return leafs
}

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetYangName() string { return "address" }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetParent() types.Entity { return address.parent }

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetParentYangName() string { return "addresses" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts
// Table of host addresses
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A single host address. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host.
    Host []ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host
}

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetYangName() string { return "hosts" }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetParentYangName() string { return "udf-object" }

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host
// A single host address
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Host ipv4 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HostAddress interface{}
}

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetGoName(yname string) string {
    if yname == "host-address" { return "HostAddress" }
    return ""
}

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetSegmentPath() string {
    return "host" + "[host-address='" + fmt.Sprintf("%v", host.HostAddress) + "']"
}

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-address"] = host.HostAddress
    return leafs
}

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetYangName() string { return "host" }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetParentYangName() string { return "hosts" }

