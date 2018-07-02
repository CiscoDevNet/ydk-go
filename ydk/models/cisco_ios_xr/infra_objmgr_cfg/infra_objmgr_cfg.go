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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port object group.
    Port ObjectGroup_Port

    // Network object group.
    Network ObjectGroup_Network
}

func (objectGroup *ObjectGroup) GetEntityData() *types.CommonEntityData {
    objectGroup.EntityData.YFilter = objectGroup.YFilter
    objectGroup.EntityData.YangName = "object-group"
    objectGroup.EntityData.BundleName = "cisco_ios_xr"
    objectGroup.EntityData.ParentYangName = "Cisco-IOS-XR-infra-objmgr-cfg"
    objectGroup.EntityData.SegmentPath = "Cisco-IOS-XR-infra-objmgr-cfg:object-group"
    objectGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectGroup.EntityData.Children = types.NewOrderedMap()
    objectGroup.EntityData.Children.Append("port", types.YChild{"Port", &objectGroup.Port})
    objectGroup.EntityData.Children.Append("network", types.YChild{"Network", &objectGroup.Network})
    objectGroup.EntityData.Leafs = types.NewOrderedMap()

    objectGroup.EntityData.YListKeys = []string {}

    return &(objectGroup.EntityData)
}

// ObjectGroup_Port
// Port object group
type ObjectGroup_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of port objects groups.
    UdfObjects ObjectGroup_Port_UdfObjects
}

func (port *ObjectGroup_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "object-group"
    port.EntityData.SegmentPath = "port"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = types.NewOrderedMap()
    port.EntityData.Children.Append("udf-objects", types.YChild{"UdfObjects", &port.UdfObjects})
    port.EntityData.Leafs = types.NewOrderedMap()

    port.EntityData.YListKeys = []string {}

    return &(port.EntityData)
}

// ObjectGroup_Port_UdfObjects
// Table of port objects groups
type ObjectGroup_Port_UdfObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port object group. The type is slice of
    // ObjectGroup_Port_UdfObjects_UdfObject.
    UdfObject []*ObjectGroup_Port_UdfObjects_UdfObject
}

func (udfObjects *ObjectGroup_Port_UdfObjects) GetEntityData() *types.CommonEntityData {
    udfObjects.EntityData.YFilter = udfObjects.YFilter
    udfObjects.EntityData.YangName = "udf-objects"
    udfObjects.EntityData.BundleName = "cisco_ios_xr"
    udfObjects.EntityData.ParentYangName = "port"
    udfObjects.EntityData.SegmentPath = "udf-objects"
    udfObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udfObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udfObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udfObjects.EntityData.Children = types.NewOrderedMap()
    udfObjects.EntityData.Children.Append("udf-object", types.YChild{"UdfObject", nil})
    for i := range udfObjects.UdfObject {
        udfObjects.EntityData.Children.Append(types.GetSegmentPath(udfObjects.UdfObject[i]), types.YChild{"UdfObject", udfObjects.UdfObject[i]})
    }
    udfObjects.EntityData.Leafs = types.NewOrderedMap()

    udfObjects.EntityData.YListKeys = []string {}

    return &(udfObjects.EntityData)
}

// ObjectGroup_Port_UdfObjects_UdfObject
// Port object group
type ObjectGroup_Port_UdfObjects_UdfObject struct {
    EntityData types.CommonEntityData
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

func (udfObject *ObjectGroup_Port_UdfObjects_UdfObject) GetEntityData() *types.CommonEntityData {
    udfObject.EntityData.YFilter = udfObject.YFilter
    udfObject.EntityData.YangName = "udf-object"
    udfObject.EntityData.BundleName = "cisco_ios_xr"
    udfObject.EntityData.ParentYangName = "udf-objects"
    udfObject.EntityData.SegmentPath = "udf-object" + types.AddKeyToken(udfObject.ObjectName, "object-name")
    udfObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udfObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udfObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udfObject.EntityData.Children = types.NewOrderedMap()
    udfObject.EntityData.Children.Append("operators", types.YChild{"Operators", &udfObject.Operators})
    udfObject.EntityData.Children.Append("nested-groups", types.YChild{"NestedGroups", &udfObject.NestedGroups})
    udfObject.EntityData.Children.Append("port-ranges", types.YChild{"PortRanges", &udfObject.PortRanges})
    udfObject.EntityData.Leafs = types.NewOrderedMap()
    udfObject.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", udfObject.ObjectName})
    udfObject.EntityData.Leafs.Append("description", types.YLeaf{"Description", udfObject.Description})

    udfObject.EntityData.YListKeys = []string {"ObjectName"}

    return &(udfObject.EntityData)
}

// ObjectGroup_Port_UdfObjects_UdfObject_Operators
// Table of port operators
type ObjectGroup_Port_UdfObjects_UdfObject_Operators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op class. The type is slice of
    // ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator.
    Operator []*ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator
}

func (operators *ObjectGroup_Port_UdfObjects_UdfObject_Operators) GetEntityData() *types.CommonEntityData {
    operators.EntityData.YFilter = operators.YFilter
    operators.EntityData.YangName = "operators"
    operators.EntityData.BundleName = "cisco_ios_xr"
    operators.EntityData.ParentYangName = "udf-object"
    operators.EntityData.SegmentPath = "operators"
    operators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operators.EntityData.Children = types.NewOrderedMap()
    operators.EntityData.Children.Append("operator", types.YChild{"Operator", nil})
    for i := range operators.Operator {
        operators.EntityData.Children.Append(types.GetSegmentPath(operators.Operator[i]), types.YChild{"Operator", operators.Operator[i]})
    }
    operators.EntityData.Leafs = types.NewOrderedMap()

    operators.EntityData.YListKeys = []string {}

    return &(operators.EntityData)
}

// ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator
// op class
type ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. operation for ports. The type is PortOperator.
    OperatorType interface{}

    // This attribute is a key. Port number. The type is one of the following
    // types: enumeration Port, or int with range: 0..65535.
    Port interface{}
}

func (operator *ObjectGroup_Port_UdfObjects_UdfObject_Operators_Operator) GetEntityData() *types.CommonEntityData {
    operator.EntityData.YFilter = operator.YFilter
    operator.EntityData.YangName = "operator"
    operator.EntityData.BundleName = "cisco_ios_xr"
    operator.EntityData.ParentYangName = "operators"
    operator.EntityData.SegmentPath = "operator" + types.AddKeyToken(operator.OperatorType, "operator-type") + types.AddKeyToken(operator.Port, "port")
    operator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operator.EntityData.Children = types.NewOrderedMap()
    operator.EntityData.Leafs = types.NewOrderedMap()
    operator.EntityData.Leafs.Append("operator-type", types.YLeaf{"OperatorType", operator.OperatorType})
    operator.EntityData.Leafs.Append("port", types.YLeaf{"Port", operator.Port})

    operator.EntityData.YListKeys = []string {"OperatorType", "Port"}

    return &(operator.EntityData)
}

// ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups
// Table of nested port object groups
type ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nested object group. The type is slice of
    // ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup.
    NestedGroup []*ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups) GetEntityData() *types.CommonEntityData {
    nestedGroups.EntityData.YFilter = nestedGroups.YFilter
    nestedGroups.EntityData.YangName = "nested-groups"
    nestedGroups.EntityData.BundleName = "cisco_ios_xr"
    nestedGroups.EntityData.ParentYangName = "udf-object"
    nestedGroups.EntityData.SegmentPath = "nested-groups"
    nestedGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nestedGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nestedGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nestedGroups.EntityData.Children = types.NewOrderedMap()
    nestedGroups.EntityData.Children.Append("nested-group", types.YChild{"NestedGroup", nil})
    for i := range nestedGroups.NestedGroup {
        nestedGroups.EntityData.Children.Append(types.GetSegmentPath(nestedGroups.NestedGroup[i]), types.YChild{"NestedGroup", nestedGroups.NestedGroup[i]})
    }
    nestedGroups.EntityData.Leafs = types.NewOrderedMap()

    nestedGroups.EntityData.YListKeys = []string {}

    return &(nestedGroups.EntityData)
}

// ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup
// nested object group
type ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of a nested object group. The type is string
    // with length: 1..64.
    NestedGroupName interface{}
}

func (nestedGroup *ObjectGroup_Port_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetEntityData() *types.CommonEntityData {
    nestedGroup.EntityData.YFilter = nestedGroup.YFilter
    nestedGroup.EntityData.YangName = "nested-group"
    nestedGroup.EntityData.BundleName = "cisco_ios_xr"
    nestedGroup.EntityData.ParentYangName = "nested-groups"
    nestedGroup.EntityData.SegmentPath = "nested-group" + types.AddKeyToken(nestedGroup.NestedGroupName, "nested-group-name")
    nestedGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nestedGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nestedGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nestedGroup.EntityData.Children = types.NewOrderedMap()
    nestedGroup.EntityData.Leafs = types.NewOrderedMap()
    nestedGroup.EntityData.Leafs.Append("nested-group-name", types.YLeaf{"NestedGroupName", nestedGroup.NestedGroupName})

    nestedGroup.EntityData.YListKeys = []string {"NestedGroupName"}

    return &(nestedGroup.EntityData)
}

// ObjectGroup_Port_UdfObjects_UdfObject_PortRanges
// Table of port range addresses
type ObjectGroup_Port_UdfObjects_UdfObject_PortRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match only packets on a given port range. The type is slice of
    // ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange.
    PortRange []*ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange
}

func (portRanges *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges) GetEntityData() *types.CommonEntityData {
    portRanges.EntityData.YFilter = portRanges.YFilter
    portRanges.EntityData.YangName = "port-ranges"
    portRanges.EntityData.BundleName = "cisco_ios_xr"
    portRanges.EntityData.ParentYangName = "udf-object"
    portRanges.EntityData.SegmentPath = "port-ranges"
    portRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portRanges.EntityData.Children = types.NewOrderedMap()
    portRanges.EntityData.Children.Append("port-range", types.YChild{"PortRange", nil})
    for i := range portRanges.PortRange {
        portRanges.EntityData.Children.Append(types.GetSegmentPath(portRanges.PortRange[i]), types.YChild{"PortRange", portRanges.PortRange[i]})
    }
    portRanges.EntityData.Leafs = types.NewOrderedMap()

    portRanges.EntityData.YListKeys = []string {}

    return &(portRanges.EntityData)
}

// ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange
// Match only packets on a given port range
type ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port number. The type is one of the following
    // types: enumeration StartPort, or int with range: 0..65535.
    StartPort interface{}

    // This attribute is a key. Port number. The type is one of the following
    // types: enumeration EndPort, or int with range: 0..65535.
    EndPort interface{}
}

func (portRange *ObjectGroup_Port_UdfObjects_UdfObject_PortRanges_PortRange) GetEntityData() *types.CommonEntityData {
    portRange.EntityData.YFilter = portRange.YFilter
    portRange.EntityData.YangName = "port-range"
    portRange.EntityData.BundleName = "cisco_ios_xr"
    portRange.EntityData.ParentYangName = "port-ranges"
    portRange.EntityData.SegmentPath = "port-range" + types.AddKeyToken(portRange.StartPort, "start-port") + types.AddKeyToken(portRange.EndPort, "end-port")
    portRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portRange.EntityData.Children = types.NewOrderedMap()
    portRange.EntityData.Leafs = types.NewOrderedMap()
    portRange.EntityData.Leafs.Append("start-port", types.YLeaf{"StartPort", portRange.StartPort})
    portRange.EntityData.Leafs.Append("end-port", types.YLeaf{"EndPort", portRange.EndPort})

    portRange.EntityData.YListKeys = []string {"StartPort", "EndPort"}

    return &(portRange.EntityData)
}

// ObjectGroup_Network
// Network object group
type ObjectGroup_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 object group.
    Ipv6 ObjectGroup_Network_Ipv6

    // IPv4 object group.
    Ipv4 ObjectGroup_Network_Ipv4
}

func (network *ObjectGroup_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xr"
    network.EntityData.ParentYangName = "object-group"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &network.Ipv6})
    network.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &network.Ipv4})
    network.EntityData.Leafs = types.NewOrderedMap()

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// ObjectGroup_Network_Ipv6
// IPv6 object group
type ObjectGroup_Network_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of ipv6 object groups.
    UdfObjects ObjectGroup_Network_Ipv6_UdfObjects
}

func (ipv6 *ObjectGroup_Network_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "network"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("udf-objects", types.YChild{"UdfObjects", &ipv6.UdfObjects})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects
// Table of ipv6 object groups
type ObjectGroup_Network_Ipv6_UdfObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 object group. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject.
    UdfObject []*ObjectGroup_Network_Ipv6_UdfObjects_UdfObject
}

func (udfObjects *ObjectGroup_Network_Ipv6_UdfObjects) GetEntityData() *types.CommonEntityData {
    udfObjects.EntityData.YFilter = udfObjects.YFilter
    udfObjects.EntityData.YangName = "udf-objects"
    udfObjects.EntityData.BundleName = "cisco_ios_xr"
    udfObjects.EntityData.ParentYangName = "ipv6"
    udfObjects.EntityData.SegmentPath = "udf-objects"
    udfObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udfObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udfObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udfObjects.EntityData.Children = types.NewOrderedMap()
    udfObjects.EntityData.Children.Append("udf-object", types.YChild{"UdfObject", nil})
    for i := range udfObjects.UdfObject {
        udfObjects.EntityData.Children.Append(types.GetSegmentPath(udfObjects.UdfObject[i]), types.YChild{"UdfObject", udfObjects.UdfObject[i]})
    }
    udfObjects.EntityData.Leafs = types.NewOrderedMap()

    udfObjects.EntityData.YListKeys = []string {}

    return &(udfObjects.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject
// IPv6 object group
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject struct {
    EntityData types.CommonEntityData
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

func (udfObject *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject) GetEntityData() *types.CommonEntityData {
    udfObject.EntityData.YFilter = udfObject.YFilter
    udfObject.EntityData.YangName = "udf-object"
    udfObject.EntityData.BundleName = "cisco_ios_xr"
    udfObject.EntityData.ParentYangName = "udf-objects"
    udfObject.EntityData.SegmentPath = "udf-object" + types.AddKeyToken(udfObject.ObjectName, "object-name")
    udfObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udfObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udfObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udfObject.EntityData.Children = types.NewOrderedMap()
    udfObject.EntityData.Children.Append("nested-groups", types.YChild{"NestedGroups", &udfObject.NestedGroups})
    udfObject.EntityData.Children.Append("address-ranges", types.YChild{"AddressRanges", &udfObject.AddressRanges})
    udfObject.EntityData.Children.Append("addresses", types.YChild{"Addresses", &udfObject.Addresses})
    udfObject.EntityData.Children.Append("hosts", types.YChild{"Hosts", &udfObject.Hosts})
    udfObject.EntityData.Leafs = types.NewOrderedMap()
    udfObject.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", udfObject.ObjectName})
    udfObject.EntityData.Leafs.Append("description", types.YLeaf{"Description", udfObject.Description})

    udfObject.EntityData.YListKeys = []string {"ObjectName"}

    return &(udfObject.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups
// Table of nested ipv6 object groups
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nested object group. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup.
    NestedGroup []*ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups) GetEntityData() *types.CommonEntityData {
    nestedGroups.EntityData.YFilter = nestedGroups.YFilter
    nestedGroups.EntityData.YangName = "nested-groups"
    nestedGroups.EntityData.BundleName = "cisco_ios_xr"
    nestedGroups.EntityData.ParentYangName = "udf-object"
    nestedGroups.EntityData.SegmentPath = "nested-groups"
    nestedGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nestedGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nestedGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nestedGroups.EntityData.Children = types.NewOrderedMap()
    nestedGroups.EntityData.Children.Append("nested-group", types.YChild{"NestedGroup", nil})
    for i := range nestedGroups.NestedGroup {
        nestedGroups.EntityData.Children.Append(types.GetSegmentPath(nestedGroups.NestedGroup[i]), types.YChild{"NestedGroup", nestedGroups.NestedGroup[i]})
    }
    nestedGroups.EntityData.Leafs = types.NewOrderedMap()

    nestedGroups.EntityData.YListKeys = []string {}

    return &(nestedGroups.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup
// nested object group
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Enter the name of a nested object group. The type
    // is string with length: 1..64.
    NestedGroupName interface{}
}

func (nestedGroup *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetEntityData() *types.CommonEntityData {
    nestedGroup.EntityData.YFilter = nestedGroup.YFilter
    nestedGroup.EntityData.YangName = "nested-group"
    nestedGroup.EntityData.BundleName = "cisco_ios_xr"
    nestedGroup.EntityData.ParentYangName = "nested-groups"
    nestedGroup.EntityData.SegmentPath = "nested-group" + types.AddKeyToken(nestedGroup.NestedGroupName, "nested-group-name")
    nestedGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nestedGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nestedGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nestedGroup.EntityData.Children = types.NewOrderedMap()
    nestedGroup.EntityData.Leafs = types.NewOrderedMap()
    nestedGroup.EntityData.Leafs.Append("nested-group-name", types.YLeaf{"NestedGroupName", nestedGroup.NestedGroupName})

    nestedGroup.EntityData.YListKeys = []string {"NestedGroupName"}

    return &(nestedGroup.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges
// Table of ipv6 address ranges
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Range of host addresses. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange.
    AddressRange []*ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange
}

func (addressRanges *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges) GetEntityData() *types.CommonEntityData {
    addressRanges.EntityData.YFilter = addressRanges.YFilter
    addressRanges.EntityData.YangName = "address-ranges"
    addressRanges.EntityData.BundleName = "cisco_ios_xr"
    addressRanges.EntityData.ParentYangName = "udf-object"
    addressRanges.EntityData.SegmentPath = "address-ranges"
    addressRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRanges.EntityData.Children = types.NewOrderedMap()
    addressRanges.EntityData.Children.Append("address-range", types.YChild{"AddressRange", nil})
    for i := range addressRanges.AddressRange {
        addressRanges.EntityData.Children.Append(types.GetSegmentPath(addressRanges.AddressRange[i]), types.YChild{"AddressRange", addressRanges.AddressRange[i]})
    }
    addressRanges.EntityData.Leafs = types.NewOrderedMap()

    addressRanges.EntityData.YListKeys = []string {}

    return &(addressRanges.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange
// Range of host addresses
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange struct {
    EntityData types.CommonEntityData
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

func (addressRange *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_AddressRanges_AddressRange) GetEntityData() *types.CommonEntityData {
    addressRange.EntityData.YFilter = addressRange.YFilter
    addressRange.EntityData.YangName = "address-range"
    addressRange.EntityData.BundleName = "cisco_ios_xr"
    addressRange.EntityData.ParentYangName = "address-ranges"
    addressRange.EntityData.SegmentPath = "address-range" + types.AddKeyToken(addressRange.StartAddress, "start-address") + types.AddKeyToken(addressRange.EndAddress, "end-address")
    addressRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRange.EntityData.Children = types.NewOrderedMap()
    addressRange.EntityData.Leafs = types.NewOrderedMap()
    addressRange.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", addressRange.StartAddress})
    addressRange.EntityData.Leafs.Append("end-address", types.YLeaf{"EndAddress", addressRange.EndAddress})

    addressRange.EntityData.YListKeys = []string {"StartAddress", "EndAddress"}

    return &(addressRange.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses
// Table of ipv6 addresses
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address.
    Address []*ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address
}

func (addresses *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "udf-object"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range addresses.Address {
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.Address[i]), types.YChild{"Address", addresses.Address[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address
// IPv6 address
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address struct {
    EntityData types.CommonEntityData
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

func (address *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + types.AddKeyToken(address.Prefix, "prefix") + types.AddKeyToken(address.PrefixLength, "prefix-length")
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", address.Prefix})
    address.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", address.PrefixLength})

    address.EntityData.YListKeys = []string {"Prefix", "PrefixLength"}

    return &(address.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts
// Table of ipv6 host addresses
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single host address. The type is slice of
    // ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host.
    Host []*ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host
}

func (hosts *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "udf-object"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host
// A single host address
type ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. host ipv6 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HostAddress interface{}
}

func (host *ObjectGroup_Network_Ipv6_UdfObjects_UdfObject_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.HostAddress, "host-address")
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("host-address", types.YLeaf{"HostAddress", host.HostAddress})

    host.EntityData.YListKeys = []string {"HostAddress"}

    return &(host.EntityData)
}

// ObjectGroup_Network_Ipv4
// IPv4 object group
type ObjectGroup_Network_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of ipv4 object groups.
    UdfObjects ObjectGroup_Network_Ipv4_UdfObjects
}

func (ipv4 *ObjectGroup_Network_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "network"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("udf-objects", types.YChild{"UdfObjects", &ipv4.UdfObjects})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects
// Table of ipv4 object groups
type ObjectGroup_Network_Ipv4_UdfObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 object group. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject.
    UdfObject []*ObjectGroup_Network_Ipv4_UdfObjects_UdfObject
}

func (udfObjects *ObjectGroup_Network_Ipv4_UdfObjects) GetEntityData() *types.CommonEntityData {
    udfObjects.EntityData.YFilter = udfObjects.YFilter
    udfObjects.EntityData.YangName = "udf-objects"
    udfObjects.EntityData.BundleName = "cisco_ios_xr"
    udfObjects.EntityData.ParentYangName = "ipv4"
    udfObjects.EntityData.SegmentPath = "udf-objects"
    udfObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udfObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udfObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udfObjects.EntityData.Children = types.NewOrderedMap()
    udfObjects.EntityData.Children.Append("udf-object", types.YChild{"UdfObject", nil})
    for i := range udfObjects.UdfObject {
        udfObjects.EntityData.Children.Append(types.GetSegmentPath(udfObjects.UdfObject[i]), types.YChild{"UdfObject", udfObjects.UdfObject[i]})
    }
    udfObjects.EntityData.Leafs = types.NewOrderedMap()

    udfObjects.EntityData.YListKeys = []string {}

    return &(udfObjects.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject
// IPv4 object group
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject struct {
    EntityData types.CommonEntityData
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

func (udfObject *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject) GetEntityData() *types.CommonEntityData {
    udfObject.EntityData.YFilter = udfObject.YFilter
    udfObject.EntityData.YangName = "udf-object"
    udfObject.EntityData.BundleName = "cisco_ios_xr"
    udfObject.EntityData.ParentYangName = "udf-objects"
    udfObject.EntityData.SegmentPath = "udf-object" + types.AddKeyToken(udfObject.ObjectName, "object-name")
    udfObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udfObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udfObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udfObject.EntityData.Children = types.NewOrderedMap()
    udfObject.EntityData.Children.Append("nested-groups", types.YChild{"NestedGroups", &udfObject.NestedGroups})
    udfObject.EntityData.Children.Append("address-ranges", types.YChild{"AddressRanges", &udfObject.AddressRanges})
    udfObject.EntityData.Children.Append("addresses", types.YChild{"Addresses", &udfObject.Addresses})
    udfObject.EntityData.Children.Append("hosts", types.YChild{"Hosts", &udfObject.Hosts})
    udfObject.EntityData.Leafs = types.NewOrderedMap()
    udfObject.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", udfObject.ObjectName})
    udfObject.EntityData.Leafs.Append("description", types.YLeaf{"Description", udfObject.Description})

    udfObject.EntityData.YListKeys = []string {"ObjectName"}

    return &(udfObject.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups
// Table of nested ipv4 object groups
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nested object group. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup.
    NestedGroup []*ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups) GetEntityData() *types.CommonEntityData {
    nestedGroups.EntityData.YFilter = nestedGroups.YFilter
    nestedGroups.EntityData.YangName = "nested-groups"
    nestedGroups.EntityData.BundleName = "cisco_ios_xr"
    nestedGroups.EntityData.ParentYangName = "udf-object"
    nestedGroups.EntityData.SegmentPath = "nested-groups"
    nestedGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nestedGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nestedGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nestedGroups.EntityData.Children = types.NewOrderedMap()
    nestedGroups.EntityData.Children.Append("nested-group", types.YChild{"NestedGroup", nil})
    for i := range nestedGroups.NestedGroup {
        nestedGroups.EntityData.Children.Append(types.GetSegmentPath(nestedGroups.NestedGroup[i]), types.YChild{"NestedGroup", nestedGroups.NestedGroup[i]})
    }
    nestedGroups.EntityData.Leafs = types.NewOrderedMap()

    nestedGroups.EntityData.YListKeys = []string {}

    return &(nestedGroups.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup
// Nested object group
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nested object group. The type is string with
    // length: 1..64.
    NestedGroupName interface{}
}

func (nestedGroup *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_NestedGroups_NestedGroup) GetEntityData() *types.CommonEntityData {
    nestedGroup.EntityData.YFilter = nestedGroup.YFilter
    nestedGroup.EntityData.YangName = "nested-group"
    nestedGroup.EntityData.BundleName = "cisco_ios_xr"
    nestedGroup.EntityData.ParentYangName = "nested-groups"
    nestedGroup.EntityData.SegmentPath = "nested-group" + types.AddKeyToken(nestedGroup.NestedGroupName, "nested-group-name")
    nestedGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nestedGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nestedGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nestedGroup.EntityData.Children = types.NewOrderedMap()
    nestedGroup.EntityData.Leafs = types.NewOrderedMap()
    nestedGroup.EntityData.Leafs.Append("nested-group-name", types.YLeaf{"NestedGroupName", nestedGroup.NestedGroupName})

    nestedGroup.EntityData.YListKeys = []string {"NestedGroupName"}

    return &(nestedGroup.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges
// Table of ipv4 host address ranges
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Range of host addresses. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange.
    AddressRange []*ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange
}

func (addressRanges *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges) GetEntityData() *types.CommonEntityData {
    addressRanges.EntityData.YFilter = addressRanges.YFilter
    addressRanges.EntityData.YangName = "address-ranges"
    addressRanges.EntityData.BundleName = "cisco_ios_xr"
    addressRanges.EntityData.ParentYangName = "udf-object"
    addressRanges.EntityData.SegmentPath = "address-ranges"
    addressRanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRanges.EntityData.Children = types.NewOrderedMap()
    addressRanges.EntityData.Children.Append("address-range", types.YChild{"AddressRange", nil})
    for i := range addressRanges.AddressRange {
        addressRanges.EntityData.Children.Append(types.GetSegmentPath(addressRanges.AddressRange[i]), types.YChild{"AddressRange", addressRanges.AddressRange[i]})
    }
    addressRanges.EntityData.Leafs = types.NewOrderedMap()

    addressRanges.EntityData.YListKeys = []string {}

    return &(addressRanges.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange
// Range of host addresses
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange struct {
    EntityData types.CommonEntityData
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

func (addressRange *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_AddressRanges_AddressRange) GetEntityData() *types.CommonEntityData {
    addressRange.EntityData.YFilter = addressRange.YFilter
    addressRange.EntityData.YangName = "address-range"
    addressRange.EntityData.BundleName = "cisco_ios_xr"
    addressRange.EntityData.ParentYangName = "address-ranges"
    addressRange.EntityData.SegmentPath = "address-range" + types.AddKeyToken(addressRange.StartAddress, "start-address") + types.AddKeyToken(addressRange.EndAddress, "end-address")
    addressRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRange.EntityData.Children = types.NewOrderedMap()
    addressRange.EntityData.Leafs = types.NewOrderedMap()
    addressRange.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", addressRange.StartAddress})
    addressRange.EntityData.Leafs.Append("end-address", types.YLeaf{"EndAddress", addressRange.EndAddress})

    addressRange.EntityData.YListKeys = []string {"StartAddress", "EndAddress"}

    return &(addressRange.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses
// Table of addresses
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address.
    Address []*ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address
}

func (addresses *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "udf-object"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range addresses.Address {
        addresses.EntityData.Children.Append(types.GetSegmentPath(addresses.Address[i]), types.YChild{"Address", addresses.Address[i]})
    }
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address
// IPv4 address
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address struct {
    EntityData types.CommonEntityData
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

func (address *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address" + types.AddKeyToken(address.Prefix, "prefix") + types.AddKeyToken(address.PrefixLength, "prefix-length")
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", address.Prefix})
    address.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", address.PrefixLength})

    address.EntityData.YListKeys = []string {"Prefix", "PrefixLength"}

    return &(address.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts
// Table of host addresses
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single host address. The type is slice of
    // ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host.
    Host []*ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host
}

func (hosts *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "udf-object"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host
// A single host address
type ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Host ipv4 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HostAddress interface{}
}

func (host *ObjectGroup_Network_Ipv4_UdfObjects_UdfObject_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.HostAddress, "host-address")
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("host-address", types.YLeaf{"HostAddress", host.HostAddress})

    host.EntityData.YListKeys = []string {"HostAddress"}

    return &(host.EntityData)
}

