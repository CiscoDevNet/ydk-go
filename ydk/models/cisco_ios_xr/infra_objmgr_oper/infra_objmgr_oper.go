// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-objmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   object-group: Object-group operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_objmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_objmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-objmgr-oper object-group}", reflect.TypeOf(ObjectGroup{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-objmgr-oper:object-group", reflect.TypeOf(ObjectGroup{}))
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
// Object-group operational data
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
    objectGroup.EntityData.ParentYangName = "Cisco-IOS-XR-infra-objmgr-oper"
    objectGroup.EntityData.SegmentPath = "Cisco-IOS-XR-infra-objmgr-oper:object-group"
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

    // Table of Object.
    Objects ObjectGroup_Port_Objects
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
    port.EntityData.Children.Append("objects", types.YChild{"Objects", &port.Objects})
    port.EntityData.Leafs = types.NewOrderedMap()

    port.EntityData.YListKeys = []string {}

    return &(port.EntityData)
}

// ObjectGroup_Port_Objects
// Table of Object
type ObjectGroup_Port_Objects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port object group. The type is slice of ObjectGroup_Port_Objects_Object.
    Object []*ObjectGroup_Port_Objects_Object
}

func (objects *ObjectGroup_Port_Objects) GetEntityData() *types.CommonEntityData {
    objects.EntityData.YFilter = objects.YFilter
    objects.EntityData.YangName = "objects"
    objects.EntityData.BundleName = "cisco_ios_xr"
    objects.EntityData.ParentYangName = "port"
    objects.EntityData.SegmentPath = "objects"
    objects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objects.EntityData.Children = types.NewOrderedMap()
    objects.EntityData.Children.Append("object", types.YChild{"Object", nil})
    for i := range objects.Object {
        objects.EntityData.Children.Append(types.GetSegmentPath(objects.Object[i]), types.YChild{"Object", objects.Object[i]})
    }
    objects.EntityData.Leafs = types.NewOrderedMap()

    objects.EntityData.YListKeys = []string {}

    return &(objects.EntityData)
}

// ObjectGroup_Port_Objects_Object
// Port object group
type ObjectGroup_Port_Objects_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port object group name. The type is string with
    // length: 1..64.
    ObjectName interface{}

    // Table of NestedGroup.
    NestedGroups ObjectGroup_Port_Objects_Object_NestedGroups

    // Table of Operator.
    Operators ObjectGroup_Port_Objects_Object_Operators

    // Table of PortRange.
    PortRanges ObjectGroup_Port_Objects_Object_PortRanges

    // Table of ParentGroup.
    ParentGroups ObjectGroup_Port_Objects_Object_ParentGroups
}

func (object *ObjectGroup_Port_Objects_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "objects"
    object.EntityData.SegmentPath = "object" + types.AddKeyToken(object.ObjectName, "object-name")
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = types.NewOrderedMap()
    object.EntityData.Children.Append("nested-groups", types.YChild{"NestedGroups", &object.NestedGroups})
    object.EntityData.Children.Append("operators", types.YChild{"Operators", &object.Operators})
    object.EntityData.Children.Append("port-ranges", types.YChild{"PortRanges", &object.PortRanges})
    object.EntityData.Children.Append("parent-groups", types.YChild{"ParentGroups", &object.ParentGroups})
    object.EntityData.Leafs = types.NewOrderedMap()
    object.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", object.ObjectName})

    object.EntityData.YListKeys = []string {"ObjectName"}

    return &(object.EntityData)
}

// ObjectGroup_Port_Objects_Object_NestedGroups
// Table of NestedGroup
type ObjectGroup_Port_Objects_Object_NestedGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nested object group. The type is slice of
    // ObjectGroup_Port_Objects_Object_NestedGroups_NestedGroup.
    NestedGroup []*ObjectGroup_Port_Objects_Object_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Port_Objects_Object_NestedGroups) GetEntityData() *types.CommonEntityData {
    nestedGroups.EntityData.YFilter = nestedGroups.YFilter
    nestedGroups.EntityData.YangName = "nested-groups"
    nestedGroups.EntityData.BundleName = "cisco_ios_xr"
    nestedGroups.EntityData.ParentYangName = "object"
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

// ObjectGroup_Port_Objects_Object_NestedGroups_NestedGroup
// nested object group
type ObjectGroup_Port_Objects_Object_NestedGroups_NestedGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nested object group. The type is string with
    // length: 1..64.
    NestedGroupName interface{}

    // Nested group. The type is string.
    NestedGroupNameXr interface{}
}

func (nestedGroup *ObjectGroup_Port_Objects_Object_NestedGroups_NestedGroup) GetEntityData() *types.CommonEntityData {
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
    nestedGroup.EntityData.Leafs.Append("nested-group-name-xr", types.YLeaf{"NestedGroupNameXr", nestedGroup.NestedGroupNameXr})

    nestedGroup.EntityData.YListKeys = []string {"NestedGroupName"}

    return &(nestedGroup.EntityData)
}

// ObjectGroup_Port_Objects_Object_Operators
// Table of Operator
type ObjectGroup_Port_Objects_Object_Operators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // op class. The type is slice of
    // ObjectGroup_Port_Objects_Object_Operators_Operator.
    Operator []*ObjectGroup_Port_Objects_Object_Operators_Operator
}

func (operators *ObjectGroup_Port_Objects_Object_Operators) GetEntityData() *types.CommonEntityData {
    operators.EntityData.YFilter = operators.YFilter
    operators.EntityData.YangName = "operators"
    operators.EntityData.BundleName = "cisco_ios_xr"
    operators.EntityData.ParentYangName = "object"
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

// ObjectGroup_Port_Objects_Object_Operators_Operator
// op class
type ObjectGroup_Port_Objects_Object_Operators_Operator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // operation for ports. The type is PortOperator.
    OperatorType interface{}

    // Port number. The type is one of the following types: enumeration Port, or
    // int with range: 0..65535.
    Port interface{}

    // Operator. The type is interface{} with range: 0..4294967295.
    OperatorTypeXr interface{}

    // Port. The type is interface{} with range: 0..4294967295.
    PortXr interface{}
}

func (operator *ObjectGroup_Port_Objects_Object_Operators_Operator) GetEntityData() *types.CommonEntityData {
    operator.EntityData.YFilter = operator.YFilter
    operator.EntityData.YangName = "operator"
    operator.EntityData.BundleName = "cisco_ios_xr"
    operator.EntityData.ParentYangName = "operators"
    operator.EntityData.SegmentPath = "operator"
    operator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operator.EntityData.Children = types.NewOrderedMap()
    operator.EntityData.Leafs = types.NewOrderedMap()
    operator.EntityData.Leafs.Append("operator-type", types.YLeaf{"OperatorType", operator.OperatorType})
    operator.EntityData.Leafs.Append("port", types.YLeaf{"Port", operator.Port})
    operator.EntityData.Leafs.Append("operator-type-xr", types.YLeaf{"OperatorTypeXr", operator.OperatorTypeXr})
    operator.EntityData.Leafs.Append("port-xr", types.YLeaf{"PortXr", operator.PortXr})

    operator.EntityData.YListKeys = []string {}

    return &(operator.EntityData)
}

// ObjectGroup_Port_Objects_Object_PortRanges
// Table of PortRange
type ObjectGroup_Port_Objects_Object_PortRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match only packets on a given port range. The type is slice of
    // ObjectGroup_Port_Objects_Object_PortRanges_PortRange.
    PortRange []*ObjectGroup_Port_Objects_Object_PortRanges_PortRange
}

func (portRanges *ObjectGroup_Port_Objects_Object_PortRanges) GetEntityData() *types.CommonEntityData {
    portRanges.EntityData.YFilter = portRanges.YFilter
    portRanges.EntityData.YangName = "port-ranges"
    portRanges.EntityData.BundleName = "cisco_ios_xr"
    portRanges.EntityData.ParentYangName = "object"
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

// ObjectGroup_Port_Objects_Object_PortRanges_PortRange
// Match only packets on a given port range
type ObjectGroup_Port_Objects_Object_PortRanges_PortRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start port number. The type is one of the following types: enumeration
    // StartPort, or int with range: 0..65535.
    StartPort interface{}

    // End port number. The type is one of the following types: enumeration
    // EndPort, or int with range: 0..65535.
    EndPort interface{}

    // Port start address. The type is interface{} with range: 0..4294967295.
    StartPortXr interface{}

    // Port end address. The type is interface{} with range: 0..4294967295.
    EndPortXr interface{}
}

func (portRange *ObjectGroup_Port_Objects_Object_PortRanges_PortRange) GetEntityData() *types.CommonEntityData {
    portRange.EntityData.YFilter = portRange.YFilter
    portRange.EntityData.YangName = "port-range"
    portRange.EntityData.BundleName = "cisco_ios_xr"
    portRange.EntityData.ParentYangName = "port-ranges"
    portRange.EntityData.SegmentPath = "port-range"
    portRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portRange.EntityData.Children = types.NewOrderedMap()
    portRange.EntityData.Leafs = types.NewOrderedMap()
    portRange.EntityData.Leafs.Append("start-port", types.YLeaf{"StartPort", portRange.StartPort})
    portRange.EntityData.Leafs.Append("end-port", types.YLeaf{"EndPort", portRange.EndPort})
    portRange.EntityData.Leafs.Append("start-port-xr", types.YLeaf{"StartPortXr", portRange.StartPortXr})
    portRange.EntityData.Leafs.Append("end-port-xr", types.YLeaf{"EndPortXr", portRange.EndPortXr})

    portRange.EntityData.YListKeys = []string {}

    return &(portRange.EntityData)
}

// ObjectGroup_Port_Objects_Object_ParentGroups
// Table of ParentGroup
type ObjectGroup_Port_Objects_Object_ParentGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Parent object group. The type is slice of
    // ObjectGroup_Port_Objects_Object_ParentGroups_ParentGroup.
    ParentGroup []*ObjectGroup_Port_Objects_Object_ParentGroups_ParentGroup
}

func (parentGroups *ObjectGroup_Port_Objects_Object_ParentGroups) GetEntityData() *types.CommonEntityData {
    parentGroups.EntityData.YFilter = parentGroups.YFilter
    parentGroups.EntityData.YangName = "parent-groups"
    parentGroups.EntityData.BundleName = "cisco_ios_xr"
    parentGroups.EntityData.ParentYangName = "object"
    parentGroups.EntityData.SegmentPath = "parent-groups"
    parentGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentGroups.EntityData.Children = types.NewOrderedMap()
    parentGroups.EntityData.Children.Append("parent-group", types.YChild{"ParentGroup", nil})
    for i := range parentGroups.ParentGroup {
        parentGroups.EntityData.Children.Append(types.GetSegmentPath(parentGroups.ParentGroup[i]), types.YChild{"ParentGroup", parentGroups.ParentGroup[i]})
    }
    parentGroups.EntityData.Leafs = types.NewOrderedMap()

    parentGroups.EntityData.YListKeys = []string {}

    return &(parentGroups.EntityData)
}

// ObjectGroup_Port_Objects_Object_ParentGroups_ParentGroup
// Parent object group
type ObjectGroup_Port_Objects_Object_ParentGroups_ParentGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nested object group. The type is string with
    // length: 1..64.
    ParentGroupName interface{}

    // Parent node. The type is string.
    ParentName interface{}
}

func (parentGroup *ObjectGroup_Port_Objects_Object_ParentGroups_ParentGroup) GetEntityData() *types.CommonEntityData {
    parentGroup.EntityData.YFilter = parentGroup.YFilter
    parentGroup.EntityData.YangName = "parent-group"
    parentGroup.EntityData.BundleName = "cisco_ios_xr"
    parentGroup.EntityData.ParentYangName = "parent-groups"
    parentGroup.EntityData.SegmentPath = "parent-group" + types.AddKeyToken(parentGroup.ParentGroupName, "parent-group-name")
    parentGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentGroup.EntityData.Children = types.NewOrderedMap()
    parentGroup.EntityData.Leafs = types.NewOrderedMap()
    parentGroup.EntityData.Leafs.Append("parent-group-name", types.YLeaf{"ParentGroupName", parentGroup.ParentGroupName})
    parentGroup.EntityData.Leafs.Append("parent-name", types.YLeaf{"ParentName", parentGroup.ParentName})

    parentGroup.EntityData.YListKeys = []string {"ParentGroupName"}

    return &(parentGroup.EntityData)
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

    // Table of Object.
    Objects ObjectGroup_Network_Ipv6_Objects
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
    ipv6.EntityData.Children.Append("objects", types.YChild{"Objects", &ipv6.Objects})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// ObjectGroup_Network_Ipv6_Objects
// Table of Object
type ObjectGroup_Network_Ipv6_Objects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 object group. The type is slice of
    // ObjectGroup_Network_Ipv6_Objects_Object.
    Object []*ObjectGroup_Network_Ipv6_Objects_Object
}

func (objects *ObjectGroup_Network_Ipv6_Objects) GetEntityData() *types.CommonEntityData {
    objects.EntityData.YFilter = objects.YFilter
    objects.EntityData.YangName = "objects"
    objects.EntityData.BundleName = "cisco_ios_xr"
    objects.EntityData.ParentYangName = "ipv6"
    objects.EntityData.SegmentPath = "objects"
    objects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objects.EntityData.Children = types.NewOrderedMap()
    objects.EntityData.Children.Append("object", types.YChild{"Object", nil})
    for i := range objects.Object {
        objects.EntityData.Children.Append(types.GetSegmentPath(objects.Object[i]), types.YChild{"Object", objects.Object[i]})
    }
    objects.EntityData.Leafs = types.NewOrderedMap()

    objects.EntityData.YListKeys = []string {}

    return &(objects.EntityData)
}

// ObjectGroup_Network_Ipv6_Objects_Object
// IPv6 object group
type ObjectGroup_Network_Ipv6_Objects_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 object group name - maximum 64 characters.
    // The type is string with length: 1..64.
    ObjectName interface{}

    // Table of NestedGroup.
    NestedGroups ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups

    // Table of Address.
    Addresses ObjectGroup_Network_Ipv6_Objects_Object_Addresses

    // Table of AddressRange.
    AddressRanges ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges

    // Table of parent object group.
    ParentGroups ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups

    // Table of Host.
    Hosts ObjectGroup_Network_Ipv6_Objects_Object_Hosts
}

func (object *ObjectGroup_Network_Ipv6_Objects_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "objects"
    object.EntityData.SegmentPath = "object" + types.AddKeyToken(object.ObjectName, "object-name")
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = types.NewOrderedMap()
    object.EntityData.Children.Append("nested-groups", types.YChild{"NestedGroups", &object.NestedGroups})
    object.EntityData.Children.Append("addresses", types.YChild{"Addresses", &object.Addresses})
    object.EntityData.Children.Append("address-ranges", types.YChild{"AddressRanges", &object.AddressRanges})
    object.EntityData.Children.Append("parent-groups", types.YChild{"ParentGroups", &object.ParentGroups})
    object.EntityData.Children.Append("hosts", types.YChild{"Hosts", &object.Hosts})
    object.EntityData.Leafs = types.NewOrderedMap()
    object.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", object.ObjectName})

    object.EntityData.YListKeys = []string {"ObjectName"}

    return &(object.EntityData)
}

// ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups
// Table of NestedGroup
type ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nested object group. The type is slice of
    // ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups_NestedGroup.
    NestedGroup []*ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups) GetEntityData() *types.CommonEntityData {
    nestedGroups.EntityData.YFilter = nestedGroups.YFilter
    nestedGroups.EntityData.YangName = "nested-groups"
    nestedGroups.EntityData.BundleName = "cisco_ios_xr"
    nestedGroups.EntityData.ParentYangName = "object"
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

// ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups_NestedGroup
// nested object group
type ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups_NestedGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Enter the name of a nested object group. The type
    // is string with length: 1..64.
    NestedGroupName interface{}

    // Nested group. The type is string.
    NestedGroupNameXr interface{}
}

func (nestedGroup *ObjectGroup_Network_Ipv6_Objects_Object_NestedGroups_NestedGroup) GetEntityData() *types.CommonEntityData {
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
    nestedGroup.EntityData.Leafs.Append("nested-group-name-xr", types.YLeaf{"NestedGroupNameXr", nestedGroup.NestedGroupNameXr})

    nestedGroup.EntityData.YListKeys = []string {"NestedGroupName"}

    return &(nestedGroup.EntityData)
}

// ObjectGroup_Network_Ipv6_Objects_Object_Addresses
// Table of Address
type ObjectGroup_Network_Ipv6_Objects_Object_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // ObjectGroup_Network_Ipv6_Objects_Object_Addresses_Address.
    Address []*ObjectGroup_Network_Ipv6_Objects_Object_Addresses_Address
}

func (addresses *ObjectGroup_Network_Ipv6_Objects_Object_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "object"
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

// ObjectGroup_Network_Ipv6_Objects_Object_Addresses_Address
// IPv6 address
type ObjectGroup_Network_Ipv6_Objects_Object_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 prefix x:x::x/y. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix of the IP Address. The type is interface{} with range: 0..128.
    PrefixLength interface{}

    // IPv4 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixXr interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLengthXr interface{}
}

func (address *ObjectGroup_Network_Ipv6_Objects_Object_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", address.Prefix})
    address.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", address.PrefixLength})
    address.EntityData.Leafs.Append("prefix-xr", types.YLeaf{"PrefixXr", address.PrefixXr})
    address.EntityData.Leafs.Append("prefix-length-xr", types.YLeaf{"PrefixLengthXr", address.PrefixLengthXr})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges
// Table of AddressRange
type ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Range of host addresses. The type is slice of
    // ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges_AddressRange.
    AddressRange []*ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges_AddressRange
}

func (addressRanges *ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges) GetEntityData() *types.CommonEntityData {
    addressRanges.EntityData.YFilter = addressRanges.YFilter
    addressRanges.EntityData.YangName = "address-ranges"
    addressRanges.EntityData.BundleName = "cisco_ios_xr"
    addressRanges.EntityData.ParentYangName = "object"
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

// ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges_AddressRange
// Range of host addresses
type ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges_AddressRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    EndAddress interface{}

    // Range start address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddressXr interface{}

    // Range end address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    EndAddressXr interface{}
}

func (addressRange *ObjectGroup_Network_Ipv6_Objects_Object_AddressRanges_AddressRange) GetEntityData() *types.CommonEntityData {
    addressRange.EntityData.YFilter = addressRange.YFilter
    addressRange.EntityData.YangName = "address-range"
    addressRange.EntityData.BundleName = "cisco_ios_xr"
    addressRange.EntityData.ParentYangName = "address-ranges"
    addressRange.EntityData.SegmentPath = "address-range"
    addressRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRange.EntityData.Children = types.NewOrderedMap()
    addressRange.EntityData.Leafs = types.NewOrderedMap()
    addressRange.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", addressRange.StartAddress})
    addressRange.EntityData.Leafs.Append("end-address", types.YLeaf{"EndAddress", addressRange.EndAddress})
    addressRange.EntityData.Leafs.Append("start-address-xr", types.YLeaf{"StartAddressXr", addressRange.StartAddressXr})
    addressRange.EntityData.Leafs.Append("end-address-xr", types.YLeaf{"EndAddressXr", addressRange.EndAddressXr})

    addressRange.EntityData.YListKeys = []string {}

    return &(addressRange.EntityData)
}

// ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups
// Table of parent object group
type ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Parent object group. The type is slice of
    // ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups_ParentGroup.
    ParentGroup []*ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups_ParentGroup
}

func (parentGroups *ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups) GetEntityData() *types.CommonEntityData {
    parentGroups.EntityData.YFilter = parentGroups.YFilter
    parentGroups.EntityData.YangName = "parent-groups"
    parentGroups.EntityData.BundleName = "cisco_ios_xr"
    parentGroups.EntityData.ParentYangName = "object"
    parentGroups.EntityData.SegmentPath = "parent-groups"
    parentGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentGroups.EntityData.Children = types.NewOrderedMap()
    parentGroups.EntityData.Children.Append("parent-group", types.YChild{"ParentGroup", nil})
    for i := range parentGroups.ParentGroup {
        parentGroups.EntityData.Children.Append(types.GetSegmentPath(parentGroups.ParentGroup[i]), types.YChild{"ParentGroup", parentGroups.ParentGroup[i]})
    }
    parentGroups.EntityData.Leafs = types.NewOrderedMap()

    parentGroups.EntityData.YListKeys = []string {}

    return &(parentGroups.EntityData)
}

// ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups_ParentGroup
// Parent object group
type ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups_ParentGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nested object group. The type is string with
    // length: 1..64.
    ParentGroupName interface{}

    // Parent node. The type is string.
    ParentName interface{}
}

func (parentGroup *ObjectGroup_Network_Ipv6_Objects_Object_ParentGroups_ParentGroup) GetEntityData() *types.CommonEntityData {
    parentGroup.EntityData.YFilter = parentGroup.YFilter
    parentGroup.EntityData.YangName = "parent-group"
    parentGroup.EntityData.BundleName = "cisco_ios_xr"
    parentGroup.EntityData.ParentYangName = "parent-groups"
    parentGroup.EntityData.SegmentPath = "parent-group" + types.AddKeyToken(parentGroup.ParentGroupName, "parent-group-name")
    parentGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentGroup.EntityData.Children = types.NewOrderedMap()
    parentGroup.EntityData.Leafs = types.NewOrderedMap()
    parentGroup.EntityData.Leafs.Append("parent-group-name", types.YLeaf{"ParentGroupName", parentGroup.ParentGroupName})
    parentGroup.EntityData.Leafs.Append("parent-name", types.YLeaf{"ParentName", parentGroup.ParentName})

    parentGroup.EntityData.YListKeys = []string {"ParentGroupName"}

    return &(parentGroup.EntityData)
}

// ObjectGroup_Network_Ipv6_Objects_Object_Hosts
// Table of Host
type ObjectGroup_Network_Ipv6_Objects_Object_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single host address. The type is slice of
    // ObjectGroup_Network_Ipv6_Objects_Object_Hosts_Host.
    Host []*ObjectGroup_Network_Ipv6_Objects_Object_Hosts_Host
}

func (hosts *ObjectGroup_Network_Ipv6_Objects_Object_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "object"
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

// ObjectGroup_Network_Ipv6_Objects_Object_Hosts_Host
// A single host address
type ObjectGroup_Network_Ipv6_Objects_Object_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. host ipv6 address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HostAddress interface{}

    // Host address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HostAddressXr interface{}
}

func (host *ObjectGroup_Network_Ipv6_Objects_Object_Hosts_Host) GetEntityData() *types.CommonEntityData {
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
    host.EntityData.Leafs.Append("host-address-xr", types.YLeaf{"HostAddressXr", host.HostAddressXr})

    host.EntityData.YListKeys = []string {"HostAddress"}

    return &(host.EntityData)
}

// ObjectGroup_Network_Ipv4
// IPv4 object group
type ObjectGroup_Network_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Object.
    Objects ObjectGroup_Network_Ipv4_Objects
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
    ipv4.EntityData.Children.Append("objects", types.YChild{"Objects", &ipv4.Objects})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// ObjectGroup_Network_Ipv4_Objects
// Table of Object
type ObjectGroup_Network_Ipv4_Objects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 object group. The type is slice of
    // ObjectGroup_Network_Ipv4_Objects_Object.
    Object []*ObjectGroup_Network_Ipv4_Objects_Object
}

func (objects *ObjectGroup_Network_Ipv4_Objects) GetEntityData() *types.CommonEntityData {
    objects.EntityData.YFilter = objects.YFilter
    objects.EntityData.YangName = "objects"
    objects.EntityData.BundleName = "cisco_ios_xr"
    objects.EntityData.ParentYangName = "ipv4"
    objects.EntityData.SegmentPath = "objects"
    objects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objects.EntityData.Children = types.NewOrderedMap()
    objects.EntityData.Children.Append("object", types.YChild{"Object", nil})
    for i := range objects.Object {
        objects.EntityData.Children.Append(types.GetSegmentPath(objects.Object[i]), types.YChild{"Object", objects.Object[i]})
    }
    objects.EntityData.Leafs = types.NewOrderedMap()

    objects.EntityData.YListKeys = []string {}

    return &(objects.EntityData)
}

// ObjectGroup_Network_Ipv4_Objects_Object
// IPv4 object group
type ObjectGroup_Network_Ipv4_Objects_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 object group name - maximum 64 characters.
    // The type is string with length: 1..64.
    ObjectName interface{}

    // Table of NestedGroup.
    NestedGroups ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups

    // Table of Address.
    Addresses ObjectGroup_Network_Ipv4_Objects_Object_Addresses

    // Table of AddressRange.
    AddressRanges ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges

    // Table of parent object group.
    ParentGroups ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups

    // Table of Host.
    Hosts ObjectGroup_Network_Ipv4_Objects_Object_Hosts
}

func (object *ObjectGroup_Network_Ipv4_Objects_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "objects"
    object.EntityData.SegmentPath = "object" + types.AddKeyToken(object.ObjectName, "object-name")
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = types.NewOrderedMap()
    object.EntityData.Children.Append("nested-groups", types.YChild{"NestedGroups", &object.NestedGroups})
    object.EntityData.Children.Append("addresses", types.YChild{"Addresses", &object.Addresses})
    object.EntityData.Children.Append("address-ranges", types.YChild{"AddressRanges", &object.AddressRanges})
    object.EntityData.Children.Append("parent-groups", types.YChild{"ParentGroups", &object.ParentGroups})
    object.EntityData.Children.Append("hosts", types.YChild{"Hosts", &object.Hosts})
    object.EntityData.Leafs = types.NewOrderedMap()
    object.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", object.ObjectName})

    object.EntityData.YListKeys = []string {"ObjectName"}

    return &(object.EntityData)
}

// ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups
// Table of NestedGroup
type ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nested object group. The type is slice of
    // ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups_NestedGroup.
    NestedGroup []*ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups_NestedGroup
}

func (nestedGroups *ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups) GetEntityData() *types.CommonEntityData {
    nestedGroups.EntityData.YFilter = nestedGroups.YFilter
    nestedGroups.EntityData.YangName = "nested-groups"
    nestedGroups.EntityData.BundleName = "cisco_ios_xr"
    nestedGroups.EntityData.ParentYangName = "object"
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

// ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups_NestedGroup
// Nested object group
type ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups_NestedGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nested object group. The type is string with
    // length: 1..64.
    NestedGroupName interface{}

    // Nested group. The type is string.
    NestedGroupNameXr interface{}
}

func (nestedGroup *ObjectGroup_Network_Ipv4_Objects_Object_NestedGroups_NestedGroup) GetEntityData() *types.CommonEntityData {
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
    nestedGroup.EntityData.Leafs.Append("nested-group-name-xr", types.YLeaf{"NestedGroupNameXr", nestedGroup.NestedGroupNameXr})

    nestedGroup.EntityData.YListKeys = []string {"NestedGroupName"}

    return &(nestedGroup.EntityData)
}

// ObjectGroup_Network_Ipv4_Objects_Object_Addresses
// Table of Address
type ObjectGroup_Network_Ipv4_Objects_Object_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is slice of
    // ObjectGroup_Network_Ipv4_Objects_Object_Addresses_Address.
    Address []*ObjectGroup_Network_Ipv4_Objects_Object_Addresses_Address
}

func (addresses *ObjectGroup_Network_Ipv4_Objects_Object_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "object"
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

// ObjectGroup_Network_Ipv4_Objects_Object_Addresses_Address
// IPv4 address
type ObjectGroup_Network_Ipv4_Objects_Object_Addresses_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address/prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix of the IP Address. The type is interface{} with range: 0..32.
    PrefixLength interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrefixXr interface{}

    // Prefix length. The type is interface{} with range: 0..4294967295.
    PrefixLengthXr interface{}
}

func (address *ObjectGroup_Network_Ipv4_Objects_Object_Addresses_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "addresses"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", address.Prefix})
    address.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", address.PrefixLength})
    address.EntityData.Leafs.Append("prefix-xr", types.YLeaf{"PrefixXr", address.PrefixXr})
    address.EntityData.Leafs.Append("prefix-length-xr", types.YLeaf{"PrefixLengthXr", address.PrefixLengthXr})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges
// Table of AddressRange
type ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Range of host addresses. The type is slice of
    // ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges_AddressRange.
    AddressRange []*ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges_AddressRange
}

func (addressRanges *ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges) GetEntityData() *types.CommonEntityData {
    addressRanges.EntityData.YFilter = addressRanges.YFilter
    addressRanges.EntityData.YangName = "address-ranges"
    addressRanges.EntityData.BundleName = "cisco_ios_xr"
    addressRanges.EntityData.ParentYangName = "object"
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

// ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges_AddressRange
// Range of host addresses
type ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges_AddressRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    EndAddress interface{}

    // Range start address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddressXr interface{}

    // Range end address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    EndAddressXr interface{}
}

func (addressRange *ObjectGroup_Network_Ipv4_Objects_Object_AddressRanges_AddressRange) GetEntityData() *types.CommonEntityData {
    addressRange.EntityData.YFilter = addressRange.YFilter
    addressRange.EntityData.YangName = "address-range"
    addressRange.EntityData.BundleName = "cisco_ios_xr"
    addressRange.EntityData.ParentYangName = "address-ranges"
    addressRange.EntityData.SegmentPath = "address-range"
    addressRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressRange.EntityData.Children = types.NewOrderedMap()
    addressRange.EntityData.Leafs = types.NewOrderedMap()
    addressRange.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", addressRange.StartAddress})
    addressRange.EntityData.Leafs.Append("end-address", types.YLeaf{"EndAddress", addressRange.EndAddress})
    addressRange.EntityData.Leafs.Append("start-address-xr", types.YLeaf{"StartAddressXr", addressRange.StartAddressXr})
    addressRange.EntityData.Leafs.Append("end-address-xr", types.YLeaf{"EndAddressXr", addressRange.EndAddressXr})

    addressRange.EntityData.YListKeys = []string {}

    return &(addressRange.EntityData)
}

// ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups
// Table of parent object group
type ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Parent object group. The type is slice of
    // ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups_ParentGroup.
    ParentGroup []*ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups_ParentGroup
}

func (parentGroups *ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups) GetEntityData() *types.CommonEntityData {
    parentGroups.EntityData.YFilter = parentGroups.YFilter
    parentGroups.EntityData.YangName = "parent-groups"
    parentGroups.EntityData.BundleName = "cisco_ios_xr"
    parentGroups.EntityData.ParentYangName = "object"
    parentGroups.EntityData.SegmentPath = "parent-groups"
    parentGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentGroups.EntityData.Children = types.NewOrderedMap()
    parentGroups.EntityData.Children.Append("parent-group", types.YChild{"ParentGroup", nil})
    for i := range parentGroups.ParentGroup {
        parentGroups.EntityData.Children.Append(types.GetSegmentPath(parentGroups.ParentGroup[i]), types.YChild{"ParentGroup", parentGroups.ParentGroup[i]})
    }
    parentGroups.EntityData.Leafs = types.NewOrderedMap()

    parentGroups.EntityData.YListKeys = []string {}

    return &(parentGroups.EntityData)
}

// ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups_ParentGroup
// Parent object group
type ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups_ParentGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Nested object group. The type is string with
    // length: 1..64.
    ParentGroupName interface{}

    // Parent node. The type is string.
    ParentName interface{}
}

func (parentGroup *ObjectGroup_Network_Ipv4_Objects_Object_ParentGroups_ParentGroup) GetEntityData() *types.CommonEntityData {
    parentGroup.EntityData.YFilter = parentGroup.YFilter
    parentGroup.EntityData.YangName = "parent-group"
    parentGroup.EntityData.BundleName = "cisco_ios_xr"
    parentGroup.EntityData.ParentYangName = "parent-groups"
    parentGroup.EntityData.SegmentPath = "parent-group" + types.AddKeyToken(parentGroup.ParentGroupName, "parent-group-name")
    parentGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parentGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parentGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parentGroup.EntityData.Children = types.NewOrderedMap()
    parentGroup.EntityData.Leafs = types.NewOrderedMap()
    parentGroup.EntityData.Leafs.Append("parent-group-name", types.YLeaf{"ParentGroupName", parentGroup.ParentGroupName})
    parentGroup.EntityData.Leafs.Append("parent-name", types.YLeaf{"ParentName", parentGroup.ParentName})

    parentGroup.EntityData.YListKeys = []string {"ParentGroupName"}

    return &(parentGroup.EntityData)
}

// ObjectGroup_Network_Ipv4_Objects_Object_Hosts
// Table of Host
type ObjectGroup_Network_Ipv4_Objects_Object_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A single host address. The type is slice of
    // ObjectGroup_Network_Ipv4_Objects_Object_Hosts_Host.
    Host []*ObjectGroup_Network_Ipv4_Objects_Object_Hosts_Host
}

func (hosts *ObjectGroup_Network_Ipv4_Objects_Object_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "object"
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

// ObjectGroup_Network_Ipv4_Objects_Object_Hosts_Host
// A single host address
type ObjectGroup_Network_Ipv4_Objects_Object_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Host ipv4 address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    HostAddress interface{}

    // Host address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    HostAddressXr interface{}
}

func (host *ObjectGroup_Network_Ipv4_Objects_Object_Hosts_Host) GetEntityData() *types.CommonEntityData {
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
    host.EntityData.Leafs.Append("host-address-xr", types.YLeaf{"HostAddressXr", host.HostAddressXr})

    host.EntityData.YListKeys = []string {"HostAddress"}

    return &(host.EntityData)
}

